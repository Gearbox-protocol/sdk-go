package schemas

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type Debt struct {
	Id                              int64        `gorm:"primaryKey;column:id" json:"-"`
	BlockNumber                     int64        `gorm:"column:block_num" json:"blockNum"`
	SessionId                       string       `gorm:"column:session_id" json:"sessionId"`
	HealthFactor                    *core.BigInt `gorm:"-" json:"-"`
	TotalValueBI                    *core.BigInt `gorm:"-" json:"-"`
	BorrowedAmountPlusInterestBI    *core.BigInt `gorm:"-" json:"-"`
	CalHealthFactor                 *core.BigInt `gorm:"column:cal_health_factor" json:"calHealthFactor"`
	CalTotalValueBI                 *core.BigInt `gorm:"column:cal_total_value_bi" json:"calTotalValue"`
	CalBorrowedAmountPlusInterestBI *core.BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi" json:"calBorrowedAmountWithInterest"`
	CalThresholdValueBI             *core.BigInt `gorm:"column:cal_threshold_value_bi" json:"calThresholdValue"`
	RepayAmountBI                   *core.BigInt `gorm:"-" json:"-"`
	AmountToPoolBI                  *core.BigInt `gorm:"-" json:"-"`
	ProfitInUSD                     float64      `gorm:"column:profit_usd" json:"profitUSD"`
	CollateralInUSD                 float64      `gorm:"column:collateral_usd" json:"collateralUSD"`
	CollateralInUnderlying          float64      `gorm:"column:collateral_underlying" json:"collateralUnderlying"`
	ProfitInUnderlying              float64      `gorm:"column:profit_underlying" json:"profitUnderlying"`
	// field not present in current_debts
	TotalValueInUSD float64 `gorm:"column:total_value_usd" json:"totalValueInUSD"`
}

type CurrentDebt struct {
	SessionId                       string       `gorm:"column:session_id;primaryKey" json:"sessionId"`
	BlockNumber                     int64        `gorm:"column:block_num" json:"blockNum"`
	CalHealthFactor                 *core.BigInt `gorm:"column:cal_health_factor" json:"calHealthFactor"`
	CalTotalValue                   float64      `gorm:"column:cal_total_value" json:"calTotalValue"`
	CalTotalValueBI                 *core.BigInt `gorm:"column:cal_total_value_bi" json:"-"`
	CalBorrowedAmountPlusInterest   float64      `gorm:"column:cal_borrowed_amt_with_interest" json:"calBorrowedAmountPlusInterest"`
	CalBorrowedAmountPlusInterestBI *core.BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi" json:"-"`
	CalThresholdValue               float64      `gorm:"column:cal_threshold_value" json:"calThresholdValue"`
	CalThresholdValueBI             *core.BigInt `gorm:"column:cal_threshold_value_bi" json:"-"`
	AmountToPool                    float64      `gorm:"-" json:"amountToPool"`
	RepayAmountBI                   *core.BigInt `gorm:"column:repay_amount_bi" json:"-"`
	RepayAmount                     float64      `gorm:"column:repay_amount" json:"repayAmount"`
	ProfitInUSD                     float64      `gorm:"column:profit_usd" json:"profitUSD"`
	ProfitInUnderlying              float64      `gorm:"column:profit_underlying" json:"profitUnderlying"`
	CollateralInUSD                 float64      `gorm:"column:collateral_usd" json:"collateralUSD"`
	CollateralInUnderlying          float64      `gorm:"column:collateral_underlying" json:"collateralUnderlying"`
}

func (CurrentDebt) TableName() string {
	return "current_debts"
}

type DebtSync struct {
	LastCalculatedAt int64 `gorm:"column:last_calculated_at"`
	FieldSet         bool  `gorm:"column:field_set;primaryKey"`
}

func (DebtSync) TableName() string {
	return "debt_sync"
}

type TokenDetails struct {
	Price             *big.Int
	Decimals          int8
	TokenLiqThreshold *core.BigInt `json:"tokenLiqThreshold"`
	Symbol            string       `json:"symbol"`
	Version           int16        `json:"version"`
}

type ProfileTable struct {
	Profile string `gorm:"column:profile"`
}

func (ProfileTable) TableName() string {
	return "profiles"
}

type LiquidableAccount struct {
	SessionId            string `gorm:"primaryKey;column:session_id"`
	BlockNum             int64  `gorm:"column:block_num"`
	NotifiedIfLiquidable bool   `gorm:"column:notified_if_liquidable"`
	Updated              bool   `gorm:"-"`
}
