package schemas

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

const (
	Active = iota
	Closed
	Repaid
	// status type for v2 liquidation is set after the call from tenderly is fetched
	// status for v3 liquidation is set in the onLiquidateCreditAccount via expirationDate and based on the paused state of the credit manager
	Liquidated
	LiquidateExpired
	LiquidatePaused
	//
)

func AccountStatusToStr(id int) (string, error) {
	switch id {
	case 0:
		return "Active", nil
	case 1:
		return "Closed", nil
	case 2:
		return "Repaid", nil
	case 3:
		return "Liquidated", nil
	case 4:
		return "LiquidateExpired", nil
	case 5:
		return "LiquidatePaused", nil
	}
	return "", fmt.Errorf("unknown status id: %d", id)
}
func AccountStatusStrToId(status string) (int, error) {
	switch strings.ToLower(status) {
	case "active":
		return 0, nil
	case "closed":
		return 1, nil
	case "repaid":
		return 2, nil
	case "liquidated":
		return 3, nil
	case "liquidateexpired":
		return 4, nil
	case "liquidatepaused":
		return 5, nil
	}
	return -1, fmt.Errorf("unknown status type: %s", status)
}

func AllowedAccountStatus(status int) bool {
	return status >= Active && status <= LiquidatePaused
}

func IsStatusLiquidated(status int) bool {
	return status == Liquidated || status == LiquidateExpired || status == LiquidatePaused
}

type (
	// TODO: delete remainingfunds/ health_factor/ profit_percent / profit not required
	CreditSession struct {
		ID             string                `gorm:"primaryKey" json:"sessionId"`
		Status         int                   `json:"status"`
		TeritaryStatus *core.Json            `gorm:"column:teritary_status"`
		Borrower       string                `json:"borrower"`
		CreditManager  string                `json:"creditManager"`
		Account        string                `json:"account"`
		Since          int64                 `json:"since"`
		ClosedAt       int64                 `json:"closedAt"`
		InitialAmount  *core.BigInt          `json:"initialAmount"`
		BorrowedAmount *core.BigInt          `json:"borrowedAmount"`
		Balances       *core.DBBalanceFormat `gorm:"column:balances"`
		Collateral     *core.JsonBigIntMap   `gorm:"column:collateral" json:"-"`
		// these two values are set and used for setting values in css
		CollateralInUSD        float64          `gorm:"<-:false;column:collateral_usd"`
		CollateralInUnderlying float64          `gorm:"<-:false;column:collateral_underlying"`
		IsDirty                bool             `gorm:"-"`
		Liquidator             string           `gorm:"liquidator"`
		Version                core.VersionType `gorm:"version"`

		// v1 liquidte and close use remainingFunds for debt calculation
		// v2 liquidate also uses it for calculation
		// it shouldn't be directly used in API
		RemainingFunds *core.BigInt `gorm:"column:remaining_funds"`
		// for v2 close accounts.
		CloseTransfers *core.JsonFloatMap `gorm:"column:close_transfers" json:"-"`
		//
		EntryPrice float64 `gorm:"column:entry_price" json:"entryPrice,omitempty"`
		ClosePrice float64 `gorm:"column:close_price" json:"closePrice,omitempty"`
	}

	CreditAccountData struct {
		Address                    string
		Borrower                   string
		InUse                      bool
		CreditManager              string
		Kind                       string
		UnderlyingToken            string
		BorrowedAmountPlusInterest *big.Int
		TotalValue                 *big.Int
		HealthFactor               *big.Int
		BorrowRate                 *big.Int
	}

	CreditAccountDataExtended struct {
		CreditAccountData
		RepayAmount           *big.Int
		LiquidationAmount     *big.Int
		BorrowedAmount        *big.Int
		СumulativeIndexAtOpen *big.Int
		Since                 int64
	}
	CreditSessionSnapshot struct {
		ID                      int64        `gorm:"primaryKey;autoincrement:true" json:"-"`
		BlockNum                int64        `gorm:"column:block_num" json:"blockNum"`
		SessionId               string       `gorm:"column:session_id" json:"sessionId"`
		BorrowedAmountBI        *core.BigInt `gorm:"column:borrowed_amount_bi" json:"borrowedAmountBI"`
		BorrowedAmount          float64      `gorm:"column:borrowed_amount" json:"borrowedAmount"`
		ExtraQuotaAPY           float64      `gorm:"column:extra_quota_apy" json:"extraQuotaAPY"`
		TotalValueBI            *core.BigInt `gorm:"column:total_value_bi" json:"totalValueBI"`
		TotalValue              float64      `gorm:"column:total_value" json:"totalValue"`
		CumulativeQuotaInterest *core.BigInt `gorm:"column:cum_quota_interest" json:"cumQuotaInterest,omitempty"`
		QuotaFees               *core.BigInt `gorm:"column:quota_fees" json:"quotaFees,omitempty"`
		// enabled can be false but amount is always non -zero
		Balances               *core.DBBalanceFormat `gorm:"column:balances" json:"balance"`
		Borrower               string                `gorm:"column:borrower" json:"borrower"`
		CollateralInUSD        float64               `gorm:"column:collateral_usd" json:"collateralInUSD"`
		CollateralInUnderlying float64               `gorm:"column:collateral_underlying" json:"collateralInUnderlying"`
		СumulativeIndexAtOpen  *core.BigInt          `gorm:"column:cumulative_index" json:"cumulativeIndexAtOpen"`
		HealthFactor           *core.BigInt          `gorm:"column:health_factor" json:"healthFactor"`
		CM                     string                `json:"-" gorm:"-"`
	}
	CreditSessionUpdate struct {
		SessionId        string       `gorm:"column:id;primaryKey"`
		BorrowedAmountBI *core.BigInt `gorm:"column:borrowed_amount_bi"`
		TotalValueBI     *core.BigInt `gorm:"column:total_value_bi"`
		Borrower         string       `gorm:"column:borrower"`
		HealthFactor     *core.BigInt `gorm:"column:health_factor"`
	}
)

func (CreditSessionUpdate) TableName() string {
	return "credit_sessions"
}

func (ses CreditSession) secStatusMap() map[int64]int {
	ans := map[int64]int{}
	if ses.TeritaryStatus != nil {
		secStatus := utils.ListOfInt64List((*ses.TeritaryStatus)["secStatus"])
		for _, l := range secStatus {
			ans[l[0]] = int(l[1])
		}
	}
	return ans
}

func (ses CreditSession) StatusAt(blockNum int64) int {
	if ses.ClosedAt == blockNum && ses.Status != Active {
		return ses.Status
	}
	if ses.TeritaryStatus == nil {
		return Active
	}
	secStatus := ses.secStatusMap()
	if secStatus[blockNum] != 0 {
		return secStatus[blockNum]
	}
	return Active
}

// in 10**27
func QuotaBorrowRate(balances core.DBBalanceFormat, totalValue *core.BigInt) float64 {
	if totalValue.Convert().Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	//
	total := new(big.Float)
	for _, balance := range balances {
		if balance.IsQuoted && balance.Quota != nil {
			total = new(big.Float).Add(
				total,
				new(big.Float).Mul(
					utils.GetFloat64(balance.Quota.Convert(), 0),
					big.NewFloat(float64(balance.QuotaRate)),
				),
			)
		}
	}
	totalValueF := utils.GetFloat64(totalValue.Convert(), 0)
	val, _ := new(big.Float).Quo(total, totalValueF).Float64()
	return val
}

func QuotaBorrowRAY(extraQuotaAPY float64) *big.Int {
	f := new(big.Float).Mul(big.NewFloat(extraQuotaAPY), utils.GetExpFloat(23))
	val, _ := f.Int(nil)
	return val
}
