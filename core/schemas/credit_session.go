package schemas

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
)

const (
	Active = iota
	Closed
	Repaid
	Liquidated
	LiquidateExpired
	LiquidatePaused
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
		CollateralInUSD        float64 `gorm:"<-:false;column:collateral_usd"`
		CollateralInUnderlying float64 `gorm:"<-:false;column:collateral_underlying"`
		IsDirty                bool    `gorm:"-"`
		Liquidator             string  `gorm:"liquidator"`
		Version                int16   `gorm:"version"`

		// v1 liquidte and close use remainingFunds for debt calculation
		// v2 liquidate also uses it for calculation
		// it shouldn't be directly used in API
		RemainingFunds *core.BigInt `gorm:"column:remaining_funds"`
		// for v2 close accounts.
		CloseTransfers *core.JsonFloatMap `gorm:"column:close_transfers" json:"-"`
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
		ID               int64        `gorm:"primaryKey;autoincrement:true" json:"-"`
		BlockNum         int64        `gorm:"column:block_num" json:"blockNum"`
		SessionId        string       `gorm:"column:session_id" json:"sessionId"`
		BorrowedAmountBI *core.BigInt `gorm:"column:borrowed_amount_bi" json:"borrowedAmountBI"`
		BorrowedAmount   float64      `gorm:"column:borrowed_amount" json:"borrowedAmount"`
		TotalValueBI     *core.BigInt `gorm:"column:total_value_bi" json:"totalValueBI"`
		TotalValue       float64      `gorm:"column:total_value" json:"totalValue"`
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
