package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type PriceDS struct {
	EntryPrice   float64 `json:"entryPrice,omitempty"`
	TradingToken string  `json:"tradingToken,omitempty"`
}

func (session *PriceDS) SetEntryPrice(tradingToken string, price float64) {
	session.TradingToken = tradingToken
	session.EntryPrice = price
}

type EntryPriceI interface {
	GetBalances() map[string]core.BalanceType
	GetBorrowedAmount() *big.Int
	GetCollateral() *core.JsonBigIntMap
	GetUnderlyingToken() string
	SetEntryPrice(tradingToken string, amt float64)
}

type tokenI interface {
	GetToken(token string) *schemas.Token
}

func CalcEntryPrice(session EntryPriceI, usdc string, store tokenI) {
	if session.GetCollateral() == nil {
		return
	}
	collateral := *session.GetCollateral()
	balances := session.GetBalances()
	// etry price is valid then
	if len(collateral) == 1 {
		isLong := collateral[usdc] != nil
		bToken, bBal, ok := singleEntryPriceToken(balances) // balance token, returns false if there are more than 1 token
		if !ok {
			return
		}
		cToken, cBal := singleEntryPriceCToken(collateral) // col token
		if cToken == session.GetUnderlyingToken() {
			cBal = new(big.Int).Add(cBal, session.GetBorrowedAmount())
		}
		if isLong {
			tradingToken := bToken
			entryPrice := utils.GetFloat64Decimal(
				new(big.Int).Quo(
					utils.GetInt64(cBal, -store.GetToken(bToken).Decimals),
					bBal),
				store.GetToken(cToken).Decimals)
			session.SetEntryPrice(tradingToken, entryPrice)
		} else {
			tradingToken := cToken
			entryPrice := utils.GetFloat64Decimal(
				new(big.Int).Quo(
					utils.GetInt64(bBal, -store.GetToken(cToken).Decimals),
					cBal,
				),
				store.GetToken(bToken).Decimals)
			session.SetEntryPrice(tradingToken, entryPrice)
		}
	}
}

func singleEntryPriceCToken(bal core.JsonBigIntMap) (string, *big.Int) {
	for token, amt := range bal {
		return token, amt.Convert()
	}
	return "", nil
}

// returns false if there are more than 1 token
func singleEntryPriceToken(bal map[string]core.BalanceType) (ansToken string, ansamt *big.Int, ansOk bool) {
	tokens := 0
	for token, details := range bal {
		if details.IsEnabled && details.HasBalanceMoreThanOne() {
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
