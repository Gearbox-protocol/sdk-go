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
		_addrToSym = core.GetTokenToSymbolByChainId(chainId)
	}
	return _addrToSym
}

func TradingAndBaseTokens(chainId int64, bal core.DBBalanceFormat, underlying string) (tradingToken, baseToken string) {
	otherToken, _, ok := otherTokenAndItsBalance(bal, underlying)
	if !ok {
		return "", ""
	}
	syms := loadSymToAddrStore(chainId)
	tradingToken, baseToken, ok  = tradingAndBase(syms, otherToken, underlying)
	if !ok {
		return "", ""
	}
	return
}

// trading priority is higher than base
func tradingAndBase(m map[common.Address]core.Symbol, a, b string) (trading, base string, ok bool) {
	asym := m[common.HexToAddress(a)]
	if asym  == "WETH" || asym == "WBTC" {
		asym = asym[1:]
	}
	bsym := m[common.HexToAddress(b)]
	if bsym == "WETH" || bsym == "WBTC" {
		bsym = bsym[1:]
	}
	if core.Priority(asym) > core.Priority(bsym) {
		return a, b, core.TradingPairs[core.TradingPair{Trading: asym, Base: bsym}]
	} else {
		return b, a, core.TradingPairs[core.TradingPair{Trading: bsym, Base: asym}]
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
