package calc

import (
	"fmt"
	"log"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type PriceDS struct {
	CurrentPrice float64 `json:"closePrice"`
	TradingToken string  `json:"tradingToken,omitempty"`
	QuoteToken   string  `json:"quoteToken,omitempty"`
}

func (session *PriceDS) SetCurrentPrice(tradingToken, quoteToken string, price float64) {
	session.TradingToken = tradingToken
	session.QuoteToken = quoteToken
	session.CurrentPrice = price
}

type EntryPriceI interface {
	GetBalances() map[string]core.BalanceType
	GetBorrowedAmount() *big.Int
	GetCollateralInUnderlying() interface{}
	GetUnderlyingToken() string
	SetCurrentPrice(tradingToken, quoteToken string, amt float64)
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
func CalcCurrentPrice(session EntryPriceI, quoteTokens []string, store tokenI) {
	cToken, cBal := session.GetUnderlyingToken(), new(big.Int).Add(
		toBigInt(session.GetCollateralInUnderlying(), store.GetToken(session.GetUnderlyingToken()).Decimals), // col token,
		session.GetBorrowedAmount())
	isLong := utils.Contains(quoteTokens, session.GetUnderlyingToken())
	tradingToken, quoteToken, currentPrice, err := CalcPriceBasedOnBalanceAndCollateral(isLong, store, session.GetBalances(), cToken, cBal)
	if err != nil {
		return
	}
	session.SetCurrentPrice(tradingToken, quoteToken, currentPrice)
}
func CalcPriceBasedOnBalanceAndCollateral(isLong bool, store tokenI,
	balances map[string]core.BalanceType, cToken string, cBal *big.Int) (string, string, float64, error) {
	bToken, bBal, ok := singleEntryPriceToken(balances, cToken) // balance token, returns false if there are more than 1 token
	if !ok {
		return "", "", 0, fmt.Errorf("balances has 2 or more token")
	}
	if isLong {
		tradingToken := bToken
		currentPrice := utils.GetFloat64Decimal(
			new(big.Int).Quo(
				utils.GetInt64(cBal, -store.GetToken(bToken).Decimals), // cBal is usdc
				bBal), // bBal is tradingToken
			store.GetToken(cToken).Decimals)
		return tradingToken, cToken, currentPrice, nil
	} else {
		tradingToken := cToken
		currentPrice := utils.GetFloat64Decimal(
			new(big.Int).Quo(
				utils.GetInt64(bBal, -store.GetToken(cToken).Decimals), // bBal is usdc
				cBal,
			),
			store.GetToken(bToken).Decimals)
		return tradingToken, bToken, currentPrice, nil
	}
}

// returns false if there are more than 1 token
func singleEntryPriceToken(bal map[string]core.BalanceType, underlyingToken string) (ansToken string, ansamt *big.Int, ansOk bool) {
	tokens := 0
	for token, details := range bal {
		if details.IsEnabled && details.HasBalanceMoreThanOne() && // ignore tokens which are disabled or have zero balances
			underlyingToken != token { // ignore underlyingToken
			tokens++
			if tokens == 2 {
				return "", nil, false
			}
			ansToken = token
			ansamt = details.BI.Convert()
			ansOk = true
		}
	}
	return
}
