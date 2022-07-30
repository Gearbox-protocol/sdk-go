package test

import "github.com/Gearbox-protocol/sdk-go/core"

type DCCalls struct {
	Pools    map[string]TestPoolCallData
	CMs      map[string]TestCMCallData
	Accounts map[string]TestAccountCallData
}

func NewDCCalls() *DCCalls {
	return &DCCalls{
		Pools:    make(map[string]TestPoolCallData),
		CMs:      make(map[string]TestCMCallData),
		Accounts: make(map[string]TestAccountCallData),
	}
}

type TestTokenBalance struct {
	Token     string       `json:"token"`
	Balance   *core.BigInt `json:"balance"`
	IsAllowed bool         `json:"isAllowed"`
}

type TestAccountCallData struct {
	Addr                       string             `json:"address"`
	Borrower                   string             `json:"borrower"`
	InUse                      bool               `json:"inUse"`
	CreditManager              string             `json:"creditManager"`
	UnderlyingToken            string             `json:"underlyingToken"`
	BorrowedAmountPlusInterest *core.BigInt       `json:"borrowAmountPlusInterest"`
	TotalValue                 *core.BigInt       `json:"totalValue"`
	HealthFactor               *core.BigInt       `json:"healthFactor"`
	BorrowRate                 *core.BigInt       `json:"borrowRate"`
	Balances                   []TestTokenBalance `json:"balances"`
	RepayAmount                *core.BigInt       `json:"repayAmount"`
	LiquidationAmount          *core.BigInt       `json:"liquidationAmount"`
	CanBeClosed                bool               `json:"canBeClosed"`
	BorrowedAmount             *core.BigInt       `json:"borrowedAmount"`
	CumulativeIndexAtOpen      *core.BigInt       `json:"cumulativeIndexAtOpen"`
	Since                      *core.BigInt       `json:"since"`
}

type TestPoolCallData struct {
	Addr                   string       `json:"address"`
	IsWETH                 bool         `json:"isWETH"`
	UnderlyingToken        string       `json:"underlyingToken"`
	DieselToken            string       `json:"dieselToken"`
	LinearCumulativeIndex  *core.BigInt `json:"linearCumulativeIndex"`
	AvailableLiquidity     *core.BigInt `json:"availableLiquidity"`
	ExpectedLiquidity      *core.BigInt `json:"expectedLiquidity"`
	ExpectedLiquidityLimit *core.BigInt `json:"expectedLiquidityLimit"`
	TotalBorrowed          *core.BigInt `json:"totalBorrowed"`
	DepositAPYRAY          *core.BigInt `json:"depositAPY"`
	BorrowAPYRAY           *core.BigInt `json:"borrowAPY"`
	DieselRateRAY          *core.BigInt `json:"dieselRate"`
	WithdrawFee            *core.BigInt `json:"withdrawFee"`
	CumulativeIndexRAY     *core.BigInt `json:"cumulativeIndex"`
}

type TestCMCallData struct {
	Addr               string       `json:"address"`
	HasAccount         bool         `json:"hasAddress"`
	UnderlyingToken    string       `json:"underlyingToken"`
	IsWETH             bool         `json:"isWETH"`
	CanBorrow          bool         `json:"canBorrow"`
	BorrowRate         *core.BigInt `json:"borrowRate"`
	MinAmount          *core.BigInt `json:"minAmount"`
	MaxAmount          *core.BigInt `json:"maxAmount"`
	MaxLeverageFactor  *core.BigInt `json:"maxLeverageFactor"`
	AvailableLiquidity *core.BigInt `json:"availableLiquidity"`
}
