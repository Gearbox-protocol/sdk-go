package core

import (
	"fmt"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/utils"
)

// update with 1inch trigger config too.
var tradingSymbols = []Symbol{
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
	// on arbitrum
	"PENDLE",
}

var arbTradingTokens = []Symbol{"GMX", "ARB"}
var opTradingTokens = []Symbol{"WLD",
	"SNX",
	"OP"}

var baseSymbols []Symbol = []Symbol{"DAI", "ETH", "BTC", "USDC"}

var farmedBaseSymbols []Symbol = []Symbol{"sDAI", "yvUSDC", "yvDAI", "stETH", "yvWBTC", "yvWETH",
	// ARBITRUM farmed
	"wstETH",
	"rETH",
	"cbETH"}

// trading pair
var TradingPairs map[TradingPair]bool

var TokenToTradingPairs map[Symbol][]TradingPair

var upperCaseToRealCase map[Symbol]Symbol

// add trading and base to pair
func updateTokenToPairs(pair TradingPair, trading, base Symbol) {
	TokenToTradingPairs[trading] = append(TokenToTradingPairs[trading], pair)
	TokenToTradingPairs[base] = append(TokenToTradingPairs[base], pair)
}

func TradingSymNoW(sym Symbol) Symbol {
	if sym == "WBTC" || sym == "WETH" {
		sym = sym[1:]
	}
	return sym
}

// generate all the pairs - lie ETHUSDC, ARByvWETH
func init() {
	TradingPairs = map[TradingPair]bool{}
	TokenToTradingPairs = map[Symbol][]TradingPair{}
	// trading and base
	for _, trading := range append(tradingSymbols, append(arbTradingTokens, opTradingTokens...)...) {
		for _, base := range baseSymbols {
			pair := NewTradingPair(trading, base)
			//
			TradingPairs[pair] = true
			updateTokenToPairs(pair, trading, base)
		}
	}
	// all base pairs
	for _, trading := range append(baseSymbols, farmedBaseSymbols...) {
		for _, base := range baseSymbols {
			tP := Priority(trading)
			bP := Priority(base)
			if tP == bP {

			} else if tP > bP {
				pair := NewTradingPair(trading, base)
				TradingPairs[pair] = true
				updateTokenToPairs(pair, trading, base)
			} else {
				pair := NewTradingPair(base, trading)
				TradingPairs[pair] = true
				updateTokenToPairs(pair, trading, base)
			}
		}
	}
	// STETH to stETH, STETH for matching name in the api request , stETH for config and internal details
	upperCaseToRealCase = map[Symbol]Symbol{}
	for _, sym := range AllTradingSymbolForDBWithW() {
		sym = TradingSymNoW(sym)
		upperCaseToRealCase[Symbol(strings.ToUpper(string(sym)))] = sym
	}
}
func GetArbTradingTokens() []Symbol {
	return arbTradingTokens
}
func GetOptTradingTokens() []Symbol {
	return opTradingTokens
}

func AllTradingSymbolForDBWithW() (ans []Symbol) {
	set := map[Symbol]bool{}
	for _, symbols := range [][]Symbol{baseSymbols, arbTradingTokens, opTradingTokens, farmedBaseSymbols, tradingSymbols} {
		for _, sym := range symbols {
			set[sym] = true
		}
	}
	delete(set, "USDC")
	ans = make([]Symbol, 0, len(set))
	for token := range set {
		if utils.Contains([]string{"ETH", "BTC"}, string(token)) {
			token = Symbol("W" + string(token))
		}
		ans = append(ans, token)
	}
	return
}

func Priority(sym Symbol) int {
	switch sym {
	case "USDC", "yvUSDC", "DAI", "sDAI", "yvDAI", "USDC.e", "USDC_e":
		return 0
	case "WBTC", "yvWBTC", "BTC":
		return 2
	case "WETH", "yvWETH", "stETH", "ETH", "cbETH", "wstETH", "rETH":
		return 3
	default:
		return 100
	}
}

func GetPriceScale(p TradingPair) int64 {
	if utils.Contains([]Symbol{"BTC", "yvWBTC"}, p.Base) {
		return 10_000
	} else if utils.Contains([]Symbol{"WETH", "yvWETH", "stETH", "ETH", "cbETH", "wstETH", "rETH"}, p.Base) {
		return 1000
	} else {
		return 100
	}
}

//

type TradingPair struct {
	Trading Symbol
	Base    Symbol
}

func (p TradingPair) String() string {
	return string(p.Trading) + string(p.Base)
}

func NewTradingPair[T ~string, X ~string](trading X, base T) TradingPair {
	return TradingPair{
		Trading: Symbol(trading),
		Base:    Symbol(base),
	}
}

func (z TradingPair) MarshalJSON() ([]byte, error) {
	return []byte("\"" + z.String() + "\""), nil
}

func (z *TradingPair) UnmarshalJSON(b []byte) error {
	str := strings.ToUpper(strings.Trim(string(b), "\""))
	for _, baseRealCase := range append(farmedBaseSymbols, baseSymbols...) { // farmed before normal base
		baseUpperCase := strings.ToUpper(string(baseRealCase))
		if strings.HasSuffix(str, baseUpperCase) {
			splitInd := len(str) - len(baseUpperCase)
			//
			tradingToken := Symbol(str[:splitInd])
			z.Trading = upperCaseToRealCase[tradingToken]
			z.Base = baseRealCase
			return nil
		}
	}
	return fmt.Errorf("can't unmarshal TradingPair with input: %s", string(b))
}

// https://stackoverflow.com/questions/55335296/problem-with-marshal-unmarshal-when-key-of-map-is-a-struct for map
func (z TradingPair) MarshalText() (text []byte, err error) {
	return z.MarshalJSON()
}

func (s *TradingPair) UnmarshalText(text []byte) error {
	return s.UnmarshalJSON(text)
}

// base and yearn tokens
// for 1inch config
func InchConfigBaseTokensForTrading() (ans []Symbol) {
	set := map[Symbol]bool{}
	for _, symbols := range [][]Symbol{baseSymbols, tradingSymbols} {
		for _, sym := range symbols {
			set[sym] = true
		}
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

func YearnTradingTokens() string {
	return `
	yearnTokens: [
		{
		  token: 'yvDAI',
		  underlying: 'DAI',
		},
		{
		  token: 'yvUSDC',
		  underlying: 'USDC',
		},
		{
		  token: 'yvWETH',
		  underlying: 'WETH',
		},
		{
		  token: 'yvWBTC',
		  underlying: 'WBTC',
		},
		{
		  token: 'sDAI',
		  isMaker: true,
		  underlying: 'DAI',
		},
	  ]`
}
