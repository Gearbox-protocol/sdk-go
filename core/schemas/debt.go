package schemas

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type CommonDebtFields struct {
	BlockNumber int64 `gorm:"column:block_num" json:"blockNum"`
	//
	CalHealthFactor           *core.BigInt `gorm:"column:cal_health_factor" json:"calHealthFactor"`
	CalTotalValueBI           *core.BigInt `gorm:"column:cal_total_value_bi" json:"calTotalValue"`
	CalDebtBI                 *core.BigInt `gorm:"column:cal_borrowed_amt_with_interest_bi" json:"calBorrowedAmountPlusInterest"`
	CalBorrowedWithInterestBI *core.BigInt `gorm:"-" json:"-"`
	CalThresholdValueBI       *core.BigInt `gorm:"column:cal_threshold_value_bi" json:"calThresholdValue"`
	// these usd fields uses price at block collateral was updated, so collateralInUSD doesn't change if there is no change in collateral but change in price of collateral assets. As a result, profitInUSD too might not change.
	ProfitInUSD            float64 `gorm:"column:profit_usd" json:"profitUSD"`
	CollateralInUSD        float64 `gorm:"column:collateral_usd" json:"collateralUSD"`
	CollateralInUnderlying float64 `gorm:"column:collateral_underlying" json:"collateralUnderlying"`
	ProfitInUnderlying     float64 `gorm:"column:profit_underlying" json:"profitUnderlying"`
}
type Debt struct {
	CommonDebtFields
	Id        int64  `gorm:"primaryKey;column:id" json:"-"`
	SessionId string `gorm:"column:session_id" json:"sessionId"`
	// HealthFactor                 *core.BigInt `gorm:"-" json:"-"`
	// TotalValueBI                 *core.BigInt `gorm:"-" json:"-"`
	// BorrowedAmountPlusInterestBI *core.BigInt `gorm:"-" json:"-"`
	//
	RepayAmountBI  *core.BigInt `gorm:"-" json:"-"`
	AmountToPoolBI *core.BigInt `gorm:"-" json:"-"`
	// field not present in current_debts
	TotalValueInUSD float64 `gorm:"column:total_value_usd" json:"totalValueInUSD"`
	FarmingValUSD   float64 `gorm:"-" json:"-"`
}

type CurrentDebt struct {
	CommonDebtFields
	SessionId string `gorm:"column:session_id;primaryKey" json:"sessionId"`
	//
	CalTotalValue     float64 `gorm:"column:cal_total_value" json:"calTotalValue"`
	CalDebt           float64 `gorm:"column:cal_borrowed_amt_with_interest" json:"calBorrowedAmountPlusInterest"`
	CalThresholdValue float64 `gorm:"column:cal_threshold_value" json:"calThresholdValue"`
	//
	AmountToPool    float64      `gorm:"-" json:"amountToPool"`
	RepayAmountBI   *core.BigInt `gorm:"column:repay_amount_bi" json:"-"`
	RepayAmount     float64      `gorm:"column:repay_amount" json:"repayAmount"`
	TotalValueInUSD float64      `gorm:"column:total_value_usd" json:"totalValueInUSD"`
	TFIndex         float64      `gorm:"column:tf_index" json:"-"`
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
