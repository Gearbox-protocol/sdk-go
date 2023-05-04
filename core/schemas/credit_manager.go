package schemas

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

func (CreditManagerState) TableName() string {
	return "credit_managers"
}

type CreditManagerState struct {
	CreditManagerData
	Address           string            `gorm:"primaryKey" json:"address"`
	IsWETH            bool              `gorm:"column:is_weth"`
	PoolAddress       string            `gorm:"column:pool_address" json:"pool"`
	UnderlyingToken   string            `gorm:"column:underlying_token" json:"underlyingToken"`
	MaxLeverageFactor int64             `gorm:"column:max_leverage"`
	MinAmount         *core.BigInt      `gorm:"column:min_amount"`
	MaxAmount         *core.BigInt      `gorm:"column:max_amount"`
	Sessions          map[string]string `gorm:"-" json:"-"`
	Paused            bool              `gorm:"column:paused"`
	Version           core.VersionType  `gorm:"column:_version"`
}

type CreditManagerData struct {
	AvailableLiquidityBI    *core.BigInt `gorm:"column:available_liquidity_bi"`
	AvailableLiquidity      float64      `gorm:"column:available_liquidity"`
	OpenedAccountsCount     int          `gorm:"column:opened_accounts_count"`
	TotalOpenedAccounts     int          `gorm:"column:total_opened_accounts"`
	TotalClosedAccounts     int          `gorm:"column:total_closed_accounts"`
	TotalRepaidAccounts     int          `gorm:"column:total_repaid_accounts"`
	TotalLiquidatedAccounts int          `gorm:"column:total_liquidated_accounts"`
	UniqueUsers             int          `gorm:"column:unique_users"`
	TotalBorrowed           float64      `gorm:"column:total_borrowed"`
	// increase -> in open creditaccount (CMStatsOnOpenAccount), in increaseBorrowAmount v1/v2
	// decreased -> when pool repay is emitted, subtract borrowAmount in event(this doesn't include fees and interest)
	TotalBorrowedBI      *core.BigInt `gorm:"column:total_borrowed_bi"`
	CumulativeBorrowed   float64      `gorm:"column:cumulative_borrowed"`
	CumulativeBorrowedBI *core.BigInt `gorm:"column:cumulative_borrowed_bi"`
	TotalRepaid          float64      `gorm:"column:total_repaid"`
	TotalRepaidBI        *core.BigInt `gorm:"column:total_repaid_bi"`
	TotalProfit          float64      `gorm:"column:total_profit"`
	TotalProfitBI        *core.BigInt `gorm:"column:total_profit_bi"`
	TotalLosses          float64      `gorm:"column:total_losses"`
	TotalLossesBI        *core.BigInt `gorm:"column:total_losses_bi"`
}

type Parameters struct {
	BlockNum                   int64        `gorm:"column:block_num;primaryKey"`
	CreditManager              string       `gorm:"column:credit_manager;primaryKey"`
	MinAmount                  *core.BigInt `gorm:"column:min_amount"`
	MaxAmount                  *core.BigInt `gorm:"column:max_amount"`
	MaxLeverage                *core.BigInt `gorm:"column:max_leverage"`
	FeeInterest                uint16       `gorm:"column:fee_interest"`
	FeeLiquidation             uint16       `gorm:"column:fee_liquidation"`
	LiquidationDiscount        uint16       `gorm:"column:liq_discount"`
	FeeLiquidationExpired      uint16       `gorm:"column:fee_liquidation_expired"`
	LiquidationDiscountExpired uint16       `gorm:"column:liq_discount_expired"`
}

func NewParameters() *Parameters {
	return &Parameters{
		MinAmount:   (*core.BigInt)(big.NewInt(0)),
		MaxAmount:   (*core.BigInt)(big.NewInt(0)),
		MaxLeverage: (*core.BigInt)(big.NewInt(0)),
	}
}

func (old *Parameters) Diffv2(new *Parameters) *core.Json {
	obj := old.Diffv1(new)
	(*obj)["feeLiquidationExpired"] = []uint16{old.FeeLiquidationExpired, new.FeeLiquidationExpired}
	(*obj)["liquidationDiscountExpired"] = []uint16{old.LiquidationDiscountExpired, new.LiquidationDiscountExpired}
	return obj
}

func (old *Parameters) Diffv1(new *Parameters) *core.Json {
	obj := core.Json{}
	obj["minAmount"] = []*core.BigInt{old.MinAmount, new.MinAmount}
	obj["maxAmount"] = []*core.BigInt{old.MaxAmount, new.MaxAmount}
	obj["maxLeverage"] = []*core.BigInt{old.MaxLeverage, new.MaxLeverage}
	obj["feeInterest"] = []uint16{old.FeeInterest, new.FeeInterest}
	obj["feeLiquidation"] = []uint16{old.FeeLiquidation, new.FeeLiquidation}
	obj["liquidationDiscount"] = []uint16{old.LiquidationDiscount, new.LiquidationDiscount}
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
	BlockNum int64
	// loss in underlying token
	Loss *big.Int
	// profit in underlying token
	Profit *big.Int
	// borrowed amount is in underlying token, repaid borrow amount
	BorrowedAmount *big.Int
}

type TransferAccountAllowed struct {
	From        string `gorm:"column:sender"`
	To          string `gorm:"column:receiver"`
	Allowed     bool   `gorm:"column:allowed"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey"`
	LogId       int64  `gorm:"column:log_id;primaryKey"`
}

func (TransferAccountAllowed) TableName() string {
	return "transfer_account_allowed"
}

type Protocol struct {
	Id            string `gorm:"primaryKey;column:id;autoincrement:true" json:"-"`
	Protocol      string `gorm:"column:protocol" json:"protocol"`
	Adapter       string `gorm:"column:adapter" json:"adapter"`
	BlockNumber   int64  `gorm:"column:block_num" json:"blockNum"`
	CreditManager string `gorm:"column:credit_manager" json:"creditManager"`
	Configurator  string `gorm:"column:configurator" json:"configurator"`
}

func (Protocol) TableName() string {
	return "allowed_protocols"
}
