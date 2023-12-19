package core

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

var TRADING_SYMBOLS = []Symbol{
	"MKR",
	"UNI",
	"LINK",
	"LDO",
	"RPL",
	"CRV",
	"APE",
	"CVX",
	"FXS",
	// "BLUR",
	"ETH",
	"BTC",
	//
	"stETH",
}

var BASE_SYMBOLS []Symbol = []Symbol{"DAI", "ETH", "BTC"}

var COMMON_TRADING_SYMBOLS []Symbol = []Symbol{"ETH", "BTC"}

func NotCommonBase() (ans []Symbol) {
	set := map[Symbol]bool{}
	for _, common := range COMMON_TRADING_SYMBOLS {
		set[common] = true
	}
	for _, base := range BASE_SYMBOLS {
		if !set[base] {
			ans = append(ans, base)
		}
	}
	return
}
func AllTradingSymbolForDBWithW() (ans []Symbol) {
	set := map[Symbol]bool{}
	for _, symbol := range TRADING_SYMBOLS {
		set[symbol] = true
	}
	for _, symbol := range BASE_SYMBOLS {
		set[symbol] = true
	}
	ans = make([]Symbol, 0, len(set))
	for token := range set {
		if utils.Contains([]string{"ETH", "BTC"}, string(token)) {
			token = Symbol("W" + string(token))
		}
		ans = append(ans, token)
	}
	return
}

type TradingPair Symbol

func NewTradingPair[T ~string, X ~string](trading X, base T) TradingPair {
	return TradingPair(string(trading) + string(base))
}

func (pair TradingPair) Tokens() (token, base Symbol) {
	if strings.HasSuffix(string(pair), "USDC") {
		return Symbol(pair[:len(pair)-len("USDC")]), Symbol("USDC")
	}
	for _, base := range BASE_SYMBOLS {
		if strings.HasSuffix(string(pair), string(base)) {
			return Symbol(pair[:len(pair)-len(base)]), base
		}
	}

	log.Fatalf("can't find the base token for pair %s", pair)
	return "", ""
}

var notAllowedPairs = []TradingPair{
	"ETHBTC",
	"DAIBTC",
	"DAIETH",
}

func PairIsAllowed(tradingSym, baseSym Symbol) bool {
	if tradingSym == baseSym {
		return false
	}
	pair := NewTradingPair(tradingSym, baseSym)
	if utils.Contains(notAllowedPairs, pair) {
		return false
	}
	return true
}
