package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"math/big"
)

func (CreditManagerState) TableName() string {
	return "credit_managers"
}

type CreditManagerState struct {
	CreditManagerData
	Address           string            `gorm:"primaryKey" json:"address"`
	IsWETH            bool              `gorm:"is_weth"`
	PoolAddress       string            `gorm:"column:pool_address" json:"pool"`
	UnderlyingToken   string            `gorm:"column:underlying_token" json:"underlyingToken"`
	MaxLeverageFactor int64             `gorm:"column:max_leverage"`
	MinAmount         *core.BigInt      `gorm:"column:min_amount"`
	MaxAmount         *core.BigInt      `gorm:"column:max_amount"`
	FeeInterest       int64             `gorm:"column:fee_interest"`
	Sessions          map[string]string `gorm:"-"`
}

type CreditManagerData struct {
	BorrowRateBI            *core.BigInt `gorm:"column:borrow_rate_bi" `
	BorrowRate              float64      `gorm:"column:borrow_rate"`
	AvailableLiquidityBI    *core.BigInt `gorm:"column:available_liquidity_bi"`
	AvailableLiquidity      float64      `gorm:"column:available_liquidity"`
	OpenedAccountsCount     int          `gorm:"column:opened_accounts_count"`
	TotalOpenedAccounts     int          `gorm:"column:total_opened_accounts"`
	TotalClosedAccounts     int          `gorm:"column:total_closed_accounts"`
	TotalRepaidAccounts     int          `gorm:"column:total_repaid_accounts"`
	TotalLiquidatedAccounts int          `gorm:"column:total_liquidated_accounts"`
	UniqueUsers             int          `gorm:"column:unique_users"`
	TotalBorrowed           float64      `gorm:"column:total_borrowed"`
	TotalBorrowedBI         *core.BigInt `gorm:"column:total_borrowed_bi"`
	CumulativeBorrowed      float64      `gorm:"column:cumulative_borrowed"`
	CumulativeBorrowedBI    *core.BigInt `gorm:"column:cumulative_borrowed_bi"`
	TotalRepaid             float64      `gorm:"column:total_repaid"`
	TotalRepaidBI           *core.BigInt `gorm:"column:total_repaid_bi"`
	TotalProfit             float64      `gorm:"column:total_profit"`
	TotalProfitBI           *core.BigInt `gorm:"column:total_profit_bi"`
	TotalLosses             float64      `gorm:"column:total_losses"`
	TotalLossesBI           *core.BigInt `gorm:"column:total_losses_bi"`
}

type Parameters struct {
	BlockNum            int64        `gorm:"column:block_num;primaryKey"`
	CreditManager       string       `gorm:"column:credit_manager;primaryKey"`
	MinAmount           *core.BigInt `gorm:"column:min_amount"`
	MaxAmount           *core.BigInt `gorm:"column:max_amount"`
	MaxLeverage         *core.BigInt `gorm:"column:max_leverage"`
	FeeInterest         *core.BigInt `gorm:"column:fee_interest"`
	FeeLiquidation      *core.BigInt `gorm:"column:fee_liquidation"`
	LiquidationDiscount *core.BigInt `gorm:"column:liq_discount"`
}

func NewParameters() *Parameters {
	return &Parameters{
		MinAmount:           (*core.BigInt)(big.NewInt(0)),
		MaxAmount:           (*core.BigInt)(big.NewInt(0)),
		MaxLeverage:         (*core.BigInt)(big.NewInt(0)),
		FeeInterest:         (*core.BigInt)(big.NewInt(0)),
		FeeLiquidation:      (*core.BigInt)(big.NewInt(0)),
		LiquidationDiscount: (*core.BigInt)(big.NewInt(0)),
	}
}

func (old *Parameters) Diff(new *Parameters) *core.Json {
	obj := core.Json{}
	obj["minAmount"] = []*core.BigInt{old.MinAmount, new.MinAmount}
	obj["maxAmount"] = []*core.BigInt{old.MaxAmount, new.MaxAmount}
	obj["maxLeverage"] = []*core.BigInt{old.MaxLeverage, new.MaxLeverage}
	obj["feeInterest"] = []*core.BigInt{old.FeeInterest, new.FeeInterest}
	obj["feeLiquidation"] = []*core.BigInt{old.FeeLiquidation, new.FeeLiquidation}
	obj["LiquidationDiscount"] = []*core.BigInt{old.LiquidationDiscount, new.LiquidationDiscount}
	return &obj
}

type FastCheckParams struct {
	BlockNum        int64        `gorm:"column:block_num;primaryKey"`
	CreditManager   string       `gorm:"column:credit_manager;primaryKey"`
	ChiThreshold    *core.BigInt `gorm:"column:chi_threshold"`
	HFCheckInterval *core.BigInt `gorm:"column:hf_checkinterval"`
}

func NewFastCheckParams() *FastCheckParams {
	return &FastCheckParams{
		ChiThreshold:    (*core.BigInt)(big.NewInt(0)),
		HFCheckInterval: (*core.BigInt)(big.NewInt(0)),
	}
}
func (old *FastCheckParams) Diff(new *FastCheckParams) *core.Json {
	obj := core.Json{}
	obj["chiThreshold"] = []*core.BigInt{old.ChiThreshold, new.ChiThreshold}
	obj["fastDelay"] = []*core.BigInt{old.HFCheckInterval, new.HFCheckInterval}
	return &obj
}

func (FastCheckParams) TableName() string {
	return "fast_check_params"
}

type CreditManagerUpdate struct {
	*CreditManagerData
	Address string `gorm:"primaryKey"`
}

func (CreditManagerUpdate) TableName() string {
	return "credit_managers"
}

type CreditManagerStat struct {
	*CreditManagerData
	ID       int64  `gorm:"primaryKey"`
	BlockNum int64  `gorm:"column:block_num"`
	Address  string `gorm:"column:credit_manager"`
}

type PnlOnRepay struct {
	Loss           *big.Int
	Profit         *big.Int
	BorrowedAmount *big.Int
}

type TransferAccountAllowed struct {
	From        string `gorm:"column:sender"`
	To          string `gorm:"column:receiver"`
	Allowed     bool   `gorm:"column:allowed"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	LogId       int64  `gorm:"column:log_id;primaryKey"`
}

type Protocol struct {
	Id            string `gorm:"primaryKey;column:id;autoincrement:true" json:"-"`
	Protocol      string `gorm:"column:protocol" json:"protocol"`
	Adapter       string `gorm:"column:adapter" json:"adapter"`
	BlockNumber   int64  `gorm:"column:block_num" json:"blockNum"`
	CreditManager string `gorm:"column:credit_manager" json:"creditManager"`
}

func (Protocol) TableName() string {
	return "allowed_protocols"
}
