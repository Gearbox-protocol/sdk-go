package priceFetcher

import (
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func powFloat(a *big.Int) *big.Float {
	f := big.NewFloat(1.0001)
	ans := big.NewFloat(1)
	absA := new(big.Int).Abs(a)
	for i := 0; i < absA.BitLen(); i++ {
		if absA.Bit(i) == 1 {
			ans = new(big.Float).Mul(f, ans)
		}
		f = new(big.Float).Mul(f, f)
	}
	if absA == a {
		return ans
	}
	return new(big.Float).Quo(big.NewFloat(1), ans)
}

// INSTRUCTIONS: TO FIX
//
// - UniTokenLastSync is not updated, as AddLastSyncForUniToken is commented
// - AddUniPools is not getting called as in chainlinkPriceFeed removed link to repo method.AddUniPoolsForToken
//
type UNIFetcher struct {
	UniPoolByToken map[string]*schemas.UniswapPools
	baseToken      string
	//
	decimalsStore *TokensStore
	tokens        []string
}

func NewPriceOnUNIFetcher(decimalsStore *TokensStore, baseToken string) *UNIFetcher {
	return &UNIFetcher{
		UniPoolByToken: map[string]*schemas.UniswapPools{},
		decimalsStore:  decimalsStore,
		baseToken:      baseToken,
	}
}

func (mdl *UNIFetcher) AddUniPools(uniswapPools *schemas.UniswapPools) {
	if mdl.UniPoolByToken[uniswapPools.Token] == nil {
		mdl.UniPoolByToken[uniswapPools.Token] = uniswapPools
	}
}

func (mdl *UNIFetcher) GetCalls() (calls []multicall.Multicall2Call, tokens []string) {
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	for token, pools := range mdl.UniPoolByToken {
		uniswapv2Price, err := v2ABI.Pack("getReserves")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V2),
			CallData: uniswapv2Price,
		})
		uniswapv3Price, err := v3ABI.Pack("slot0")
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Price,
		})
		uniswapv3Twap, err := v3ABI.Pack("observe", []uint32{0, 600})
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   common.HexToAddress(pools.V3),
			CallData: uniswapv3Twap,
		})
		tokens = append(tokens, token)
	}
	mdl.tokens = tokens
	return
}

type TokenPrice struct {
	Token    string  `gorm:"column:token;primaryKey"`
	BlockNum string  `gorm:"column:block_num;primaryKey"`
	Price    float64 `gorm:"column:price"`
}

func (mdl *UNIFetcher) GetPrices(results []multicall.Multicall2Result, _ int64) (univ2, univ3, twapV3 map[string]float64) {
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	baseDecimals := mdl.decimalsStore.GetDecimals(common.HexToAddress(mdl.baseToken))

	univ2, univ3, twapV3 = map[string]float64{}, map[string]float64{}, map[string]float64{}
	//
	for ind, entry := range results {
		tokenInd := ind / 3
		callInd := ind - tokenInd*3
		token := mdl.tokens[tokenInd]
		tokenDecimals := mdl.decimalsStore.GetDecimals(common.HexToAddress(token))
		// ignore if failed
		if !entry.Success {
			continue
		}
		switch callInd {
		case 0: // uniswap v2 price
			value, err := v2ABI.Unpack("getReserves", entry.ReturnData)
			log.CheckFatal(err)
			r0 := value[0].(*big.Int)
			r1 := value[1].(*big.Int)
			v2Price := priceInBaseToken(token, mdl.baseToken, tokenDecimals, r0, r1)
			//
			univ2[token] = utils.GetFloat64Decimal(v2Price, baseDecimals)
		case 1: // univ3 price
			value, err := v3ABI.Unpack("slot0", entry.ReturnData)
			log.CheckFatal(err)
			//https://docs.uniswap.org/sdk/guides/fetching-prices#understanding-sqrtprice
			// [(slot0**2 *Token0decimals)/2**192], divide by token for getting the float price in WETH
			//
			v3Price := univ3SlotToPriceInBase(value[0].(*big.Int), areSorted(token, mdl.baseToken), tokenDecimals)
			//
			univ3[token] = utils.GetFloat64Decimal(v3Price, baseDecimals)
		case 2: // v3 Oracle price
			value, err := v3ABI.Unpack("observe", entry.ReturnData)
			log.CheckFatal(err)
			ticks := value[0].([]*big.Int)
			// https://medium.com/blockchain-development-notes/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5
			// (t1-t0)/interval
			tickDiff := new(big.Int).Quo(new(big.Int).Sub(ticks[1], ticks[0]), big.NewInt(600))
			sqrtPrice := powFloat(tickDiff)
			decimal := 18 - tokenDecimals
			if decimal != 0 {
				sqrtPrice = new(big.Float).Mul(utils.GetExpFloat(decimal), sqrtPrice)
				sqrtPrice = new(big.Float).Quo(big.NewFloat(1), sqrtPrice)
			}
			twapV3Price, _ := sqrtPrice.Float64()
			// // if sorted use resiprocal
			// if tokenDetails.Symbol == "YFI" {
			// 	prices.TwapV3 = 1 / prices.TwapV3
			// }
			twapV3[token] = twapV3Price
		}
	}
	return
}

/////////////////////////
// UNI pools related methods
/////////////////////////

func areSorted(token, weth string) bool {
	return strings.Compare(strings.ToLower(token), strings.ToLower(weth)) == -1
}

// r1*x/(r0+x)
func priceInBaseToken(token, baseToken string, tokenDecimals int8, r0, r1 *big.Int) *big.Int {
	if !areSorted(token, baseToken) {
		r1, r0 = r0, r1
	}
	amountIn := utils.GetExpInt(tokenDecimals)
	nom := new(big.Int).Mul(r1, amountIn)
	denom := new(big.Int).Add(r0, amountIn)
	return new(big.Int).Quo(nom, denom)
}

func squareIt(a *big.Int) *big.Int {
	return new(big.Int).Mul(a, a)
}

// uni v3 slot to price in base
// returns price in usdc or weth, base is usdc/weth
// if base is token1: [(slot0**2 *Token0decimals)/2**192]
// if base is token0: [(2**192 *Token1decimals)/slot0**2]
func univ3SlotToPriceInBase(slot0 *big.Int, baseIsToken1 bool, decimals int8) (price *big.Int) {
	normalizeFactor := new(big.Int).Exp(big.NewInt(2), big.NewInt(96*2), nil)
	//
	sqSlot0 := squareIt(slot0)
	if baseIsToken1 {
		price = utils.GetInt64(sqSlot0, -1*decimals)
		price = new(big.Int).Quo(price, normalizeFactor)
	} else {
		price = utils.GetInt64(normalizeFactor, -1*decimals)
		price = new(big.Int).Quo(price, sqSlot0)
	}
	return price
}
