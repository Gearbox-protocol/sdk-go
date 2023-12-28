package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type PriceDS struct {
	CurrentPrice float64 `json:"closePrice"`
	TradingToken string  `json:"tradingToken,omitempty"`
	BaseToken    string  `json:"quoteToken,omitempty"`
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

func loadSymToAddrStore(chainId int64) map[common.Address]core.Symbol {
	if _addrToSym == nil {
		_addrToSym = core.GetAddrToSymbolByChainId(chainId)
	}
	return _addrToSym
}

func TradingAndBaseTokens(chainId int64, bal core.DBBalanceFormat, cToken string) (tradingToken, baseToken string) {
	otherToken, _, ok := singleEntryPriceToken(bal, cToken)
	if !ok {
		return "", ""
	}

	return tradingAndBase(loadSymToAddrStore(chainId), otherToken, cToken)
}

// trading priority is higher than base
func tradingAndBase(m map[common.Address]core.Symbol, a, b string) (trading, base string) {
	if priority(m, a) > priority(m, b) {
		return a, b
	} else {
		return b, a
	}
}

func priority(m map[common.Address]core.Symbol, addr string) int {
	switch m[common.HexToAddress(addr)] {
	case "USDC":
		return 0
	case "DAI":
		return 1
	case "WBTC":
		return 2
	case "WETH":
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
			!utils.Contains( // can't be underlying token or null address
				[]string{
					underlyingToken,
					core.NULL_ADDR.Hex(),
				},
				token,
			) { // ignore underlyingToken
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
