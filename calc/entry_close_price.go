package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type TradingPriceI interface {
	GetBalances() core.DBBalanceFormat
	GetBorrowedAmount() *big.Int
	GetDebt() *big.Int
	GetCollateralInUnderlying() interface{}
	GetUnderlyingToken() string
	SetCurrentPrice(tradingToken, quoteToken string, amt float64)
	GetRemaingFunds() *core.BigInt
}

type tokenI interface {
	GetToken(token string) *schemas.Token
}

func toBigInt(collateral interface{}, decimals int8) *big.Int {
	switch v := collateral.(type) {
	case float64:
		ans, _ := new(big.Float).Mul(big.NewFloat(v), utils.GetExpFloat(decimals)).Int(nil)
		return ans
	case *big.Int:
		return v
	case *core.BigInt:
		return v.Convert()
	default:
		log.Fatal("can't covert this type", v)
		return nil
	}
}

func calculatePrice(store tokenI, tradingToken, baseToken string, tradingAmount, baseAmount *big.Int) float64 {
	currentPrice := utils.GetFloat64Decimal(
		new(big.Int).Quo(
			utils.GetInt64(baseAmount, -store.GetToken(tradingToken).Decimals), // cBal is usdc
			tradingAmount), // bBal is tradingToken
		store.GetToken(baseToken).Decimals)

	return currentPrice
}

func calcTradingPrice(chainId int64, store tokenI, session TradingPriceI, cBal *big.Int) float64 {
	underlyingToken := session.GetUnderlyingToken()
	bal := session.GetBalances()
	if _, ok := bal[underlyingToken]; ok {
		cBal = new(big.Int).Sub(cBal, bal[underlyingToken].BI.Convert())
	}
	// other Token bal
	otherToken, otherAmount, ok := otherTokenAndItsBalance(bal, underlyingToken)
	if !ok {
		return 0
	}
	tradingToken, baseToken := TradingAndBaseTokens(chainId, bal, underlyingToken)
	if tradingToken == "" {
		return 0
	}

	var tradingAmount, baseAmount *big.Int
	if tradingToken == otherToken {
		tradingAmount = otherAmount
		baseAmount = cBal
	} else {
		tradingAmount = cBal
		baseAmount = otherAmount
	}

	return calculatePrice(store, tradingToken, baseToken, tradingAmount, baseAmount)
}

// calcEntryPrice
func CalcEntryPriceBySession(chainId int64, store tokenI, session TradingPriceI) float64 {
	// collteral + borrowedAmount - underlyingTokenBal = cBal
	cBal := new(big.Int).Add(
		toBigInt(session.GetCollateralInUnderlying(), store.GetToken(session.GetUnderlyingToken()).Decimals),
		session.GetBorrowedAmount(),
	)
	return calcTradingPrice(chainId, store, session, cBal)
}

// closePrice
func getRemaingFunds(session TradingPriceI) *big.Int {
	remainingFunds := session.GetRemaingFunds()
	if remainingFunds == nil {
		return big.NewInt(0)
	}
	return remainingFunds.Convert()
}
func CalcClosePriceBySession(chainId int64, store tokenI, session TradingPriceI) float64 {
	// collteral + borrowedAmount - underlyingTokenBal = cBal
	cBal := new(big.Int).Add(
		session.GetDebt(),
		getRemaingFunds(session),
	)
	return calcTradingPrice(chainId, store, session, cBal)
}
