package calc

import (
	"fmt"
	"log"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
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
	GetBalances() core.DBBalanceFormat
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

var _addrToSym map[common.Address]core.Symbol

func loadSymToAddrStore(chainId int64) {
	if _addrToSym == nil {
		_addrToSym = core.GetAddrToSymbolByChainId(chainId)
	}
}

func CalcCurrentPriceBySession(session EntryPriceI, chainId int64, store tokenI) {
	cBal := new(big.Int).Add(
		toBigInt(session.GetCollateralInUnderlying(), store.GetToken(session.GetUnderlyingToken()).Decimals),
		session.GetBorrowedAmount(),
	)
	tradingToken, quoteToken, currentPrice, err := CalcCurrentPrice(
		session.GetBalances(),
		session.GetUnderlyingToken(),
		cBal,
		chainId,
		store,
	)
	if err != nil {
		return
	}
	session.SetCurrentPrice(tradingToken, quoteToken, currentPrice)
}
func CalcCurrentPrice(bal core.DBBalanceFormat,
	cToken string, cBal *big.Int, chainId int64, store tokenI) (string, string, float64, error) {
	loadSymToAddrStore(chainId)
	//
	//
	if _, ok := bal[cToken]; ok {
		cBal = new(big.Int).Sub(cBal, bal[cToken].BI.Convert())
	}
	//
	return getCurrentPrice(bal, store, cToken, cBal)
}

func getCurrentPrice(bal core.DBBalanceFormat, store tokenI, cToken string, cBal *big.Int) (string, string, float64, error) {
	otherToken, otherAmount, ok := singleEntryPriceToken(bal, cToken)
	if !ok {
		return "", "", 0, fmt.Errorf("balances has 2 or more token")
	}

	tradingToken, baseToken := tradingAndBase(otherToken, cToken)
	var tradingAmount, baseAmount *big.Int
	if tradingToken == otherToken {
		tradingAmount = otherAmount
		baseAmount = cBal
	} else {
		tradingAmount = cBal
		baseAmount = otherAmount
	}

	currentPrice := utils.GetFloat64Decimal(
		new(big.Int).Quo(
			utils.GetInt64(baseAmount, -store.GetToken(tradingToken).Decimals), // cBal is usdc
			tradingAmount), // bBal is tradingToken
		store.GetToken(baseToken).Decimals)
	//
	return tradingToken, baseToken, currentPrice, nil
}

// trading priority is higher than base
func tradingAndBase(a, b string) (trading, base string) {
	if priority(a) > priority(b) {
		return a, b
	} else {
		return b, a
	}
}

func priority(addr string) int {
	switch _addrToSym[common.HexToAddress(addr)] {
	case "USDC":
		return 0
	case "DAI":
		return 1
	case "BTC":
		return 2
	case "ETH":
		return 3
	default:
		return 100
	}
}

// returns false if there are more than 1 token
func singleEntryPriceToken(bal core.DBBalanceFormat, underlyingToken string) (ansToken string, ansamt *big.Int, ansOk bool) {
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
