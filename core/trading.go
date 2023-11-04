package core

var TRADING_SYMBOLS = []string{
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
	"WETH",
	"WBTC",
	//
	"stETH",
}

const OTHER_TRADING_BASE_SYMBOL = "DAI"

func TradingSymNoW[T ~string](sym T) T {
	if sym[0] == 'W' {
		sym = sym[1:]
	}
	return sym
}
