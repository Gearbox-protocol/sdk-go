package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

var _addrToSym map[common.Address]core.Symbol

func loadSymToAddrStore(chainId int64) map[common.Address]core.Symbol {
	if _addrToSym == nil {
		_addrToSym = core.GetAddrToSymbolByChainId(chainId)
	}
	return _addrToSym
}

func TradingAndBaseTokens(chainId int64, bal core.DBBalanceFormat, underlying string) (tradingToken, baseToken string) {
	otherToken, _, ok := otherTokenAndItsBalance(bal, underlying)
	if !ok {
		return "", ""
	}

	return tradingAndBase(loadSymToAddrStore(chainId), otherToken, underlying)
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
	case "USDC", "yvUSDC":
		return 0
	case "DAI", "sDAI", "yvDAI":
		return 1
	case "WBTC", "yvWBTC":
		return 2
	case "WETH", "yvWETH", "stETH":
		return 3
	default:
		return 100
	}
}

// returns false if there are more than 1 token
func otherTokenAndItsBalance(bal core.DBBalanceFormat, underlyingToken string) (ansToken string, ansamt *big.Int, ansOk bool) {
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
