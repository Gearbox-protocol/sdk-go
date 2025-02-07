package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
)

type PoolState struct {
	Address                string           `gorm:"primaryKey" json:"address"`
	UnderlyingToken        string           `gorm:"column:underlying_token" json:"underlyingToken"`
	DieselToken            string           `gorm:"column:diesel_token" json:"dieselToken"`
	IsWETH                 bool             `gorm:"column:is_weth"`
	ExpectedLiquidityLimit *core.BigInt     `gorm:"column:expected_liq_limit"`
	WithdrawFee            *core.BigInt     `gorm:"column:withdraw_fee"`
	BaseBorrowAPYBI        *core.BigInt     `gorm:"column:base_borrow_apy_bi"`
	DepositAPYBI           *core.BigInt     `gorm:"column:deposit_apy_bi"`
	Version                core.VersionType `gorm:"column:_version"`
	//
	Name string `gorm:"column:name"`
}

func (PoolState) TableName() string {
	return "pools"
}

type UTokenAndPool struct {
	Pool    string
	UToken  string
	Version core.VersionType
}
type PoolStat struct {
	ID                  int64        `json:"-"`
	BlockNum            int64        `gorm:"column:block_num;primaryKey" json:"blockNum"`
	Address             string       `gorm:"column:pool;primaryKey" json:"pool"`
	UniqueUsers         int          `gorm:"column:unique_users" json:"uniqueUsers"`
	DepositAPY          float64      `gorm:"column:deposit_apy" json:"depositAPY"`
	DepositAPYBI        *core.BigInt `gorm:"column:deposit_apy_bi" json:"depositAPYBI"`
	BaseBorrowAPY       float64      `gorm:"column:base_borrow_apy" json:"baseBorrowAPY"`
	BaseBorrowAPYBI     *core.BigInt `gorm:"column:base_borrow_apy_bi" json:"baseBorrowAPYBI"`
	ExpectedLiquidity   float64      `gorm:"column:expected_liquidity" json:"expectedLiquidity"`
	ExpectedLiquidityBI *core.BigInt `gorm:"column:expected_liquidity_bi" json:"expectedLiquidityBI"`
	// ExpectedLiquidityLimitBI *core.BigInt `gorm:"column:expected_liquidity_limit_bi" json:"expectedLiquidityLimitBI"`
	AvailableLiquidity   float64      `gorm:"column:available_liquidity" json:"availableLiquidity"`
	AvailableLiquidityBI *core.BigInt `gorm:"column:available_liquidity_bi" json:"availableLiquidityBI"`
	TotalBorrowed        float64      `gorm:"column:total_borrowed" json:"totalBorrowed"`
	TotalBorrowedBI      *core.BigInt `gorm:"column:total_borrowed_bi" json:"totalBorrowedBI"`
	TotalProfit          float64      `gorm:"column:total_profit" json:"totalProfit"`
	TotalProfitBI        *core.BigInt `gorm:"column:total_profit_bi" json:"totalProfitBI"`
	TotalLosses          float64      `gorm:"column:total_losses" json:"totalLosses"`
	TotalLossesBI        *core.BigInt `gorm:"column:total_losses_bi" json:"totalLossesBI"`
	DieselRate           float64      `gorm:"column:diesel_rate" json:"dieselRate"`
	DieselRateBI         *core.BigInt `gorm:"column:diesel_rate_bi" json:"dieselRateBI"`
	WithdrawFee          int          `gorm:"column:withdraw_fee" json:"withdrawFee"`
	CumulativeIndexRAY   *core.BigInt `gorm:"column:cumulative_index_ray" json:"cumulativeIndexRAY"`
}

type PoolInterestData struct {
	BaseBorrowAPYBI      *core.BigInt `gorm:"column:base_borrow_apy_bi"`
	CumulativeIndexRAY   *core.BigInt `gorm:"column:cumulative_index_ray"`
	AvailableLiquidityBI *core.BigInt `gorm:"column:available_liquidity_bi"`
	BlockNum             int64        `gorm:"column:block_num"`
	Address              string       `gorm:"column:pool"`
	Timestamp            uint64       `gorm:"column:timestamp"`
}

type TvlSnapshots struct {
	BlockNum           int64   `gorm:"column:block_num;primaryKey"`
	AvailableLiquidity float64 `gorm:"column:available_liquidity"`
	CATotalValue       float64 `gorm:"column:ca_total_value"`
}

func (TvlSnapshots) TableName() string {
	return "tvl_snapshots"
}

type PoolLedger struct {
	Id          int64  `json:"-"`
	BlockNumber int64  `gorm:"column:block_num;primaryKey" json:"blockNum"`
	Pool        string `gorm:"column:pool;primaryKey" json:"pool"`
	// executor
	Executor string `gorm:"column:executor" json:"executor,omitempty"`
	// receiver
	Receiver string `gorm:"column:receiver" json:"receiver,omitempty"`
	// owner
	User string `gorm:"column:user_address" json:"user"`
	//
	TxHash    string `gorm:"column:tx_hash" json:"txHash"`
	SessionId string `gorm:"column:session_id" json:"sessionId"`
	LogId     uint   `gorm:"column:log_id;primaryKey" json:"logId"`
	Event     string `gorm:"column:event" json:"event"`
	// underlying token
	AmountBI *core.BigInt `gorm:"column:amount_bi" json:"-"`
	Amount   float64      `gorm:"column:amount" json:"amount"`
	// diesel token
	Shares   float64      `gorm:"column:shares" json:"shares,omitempty"`
	SharesBI *core.BigInt `gorm:"column:shares_bi" json:"-"`
}

func (PoolLedger) TableName() string {
	return "pool_ledger"
}
