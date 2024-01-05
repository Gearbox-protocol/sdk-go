package schemas_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type TradingPriceObj struct {
	SessionId              string                `gorm:"column:id;primaryKey"`
	BorrowedAmountBI       *core.BigInt          `gorm:"column:borrowed_amount_bi"`
	Balances               *core.DBBalanceFormat `gorm:"column:balances"`
	UnderlyingToken        string                `gorm:"column:underlying_token"`
	CollateralInUnderlying float64               `gorm:"column:collateral_underlying"`
	DebtBI                 *core.BigInt          `gorm:"column:cal_borrowed_amt_with_interest_bi"`
	RemainingFunds         *core.BigInt          `gorm:"column:remaining_funds"`
	ClosedAt               int64                 `gorm:"column:closed_at"`
}

func (w TradingPriceObj) GetBalances() core.DBBalanceFormat {
	return *w.Balances
}

func (w TradingPriceObj) GetBorrowedAmount() *big.Int {
	return w.BorrowedAmountBI.Convert()
}
func (w TradingPriceObj) GetDebt() *big.Int {
	if w.DebtBI == nil {
		return big.NewInt(0)
	}
	return w.DebtBI.Convert()
}

func (w TradingPriceObj) GetCollateralInUnderlying() interface{} {
	return w.CollateralInUnderlying
}

func (w TradingPriceObj) GetUnderlyingToken() string {
	return w.UnderlyingToken
}

func (w TradingPriceObj) GetRemainingFunds() *core.BigInt {
	return w.RemainingFunds
}
