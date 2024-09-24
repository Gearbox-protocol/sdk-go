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
	if absA.Cmp(a) == 0 {
		return ans
	}
	return new(big.Float).Quo(big.NewFloat(1), ans)
}

// INSTRUCTIONS: TO FIX
//
// - UniTokenLastSync is not updated, as AddLastSyncForUniToken is commented
// - AddUniPools is not getting called as in chainlinkPriceFeed removed link to repo method.AddUniPoolsForToken
type UNIFetcher struct {
	UniPoolByToken map[string]*schemas.UniswapPools
	baseToken      string
	baseDecimals   int8
	//
	decimalsStore *TokensStore
	tokens        []string
}

func NewUNIFetcher(client core.ClientI, decimalsStore *TokensStore, baseToken string) *UNIFetcher {
	decimals := core.GetDecimals(client, common.HexToAddress(baseToken), 0)
	return &UNIFetcher{
		UniPoolByToken: map[string]*schemas.UniswapPools{},
		decimalsStore:  decimalsStore,
		baseToken:      baseToken,
		baseDecimals:   decimals,
	}
}

func (mdl *UNIFetcher) AddUniPools(uniswapPools *schemas.UniswapPools) {
	if mdl.UniPoolByToken[uniswapPools.Token] == nil {
		mdl.UniPoolByToken[uniswapPools.Token] = uniswapPools
	}
}

func (mdl UNIFetcher) GetUniPools() (pools []*schemas.UniswapPools) {
	for _, entry := range mdl.UniPoolByToken {
		pools = append(pools, entry)
	}
	return
}

func (mdl *UNIFetcher) GetCalls() (calls []multicall.Multicall2Call) {
	v2ABI := core.GetAbi("Uniswapv2Pool")
	v3ABI := core.GetAbi("Uniswapv3Pool")
	tokens := []string{}
	//
	uniswapv2Price, err := v2ABI.Pack("getReserves")
	log.CheckFatal(err)
	uniswapv3Price, err := v3ABI.Pack("slot0")
	log.CheckFatal(err)
	for token, pools := range mdl.UniPoolByToken {
		if pools.V2 != core.NULL_ADDR.Hex() {
			calls = append(calls, multicall.Multicall2Call{
				Target:   common.HexToAddress(pools.V2),
				CallData: uniswapv2Price,
			})
		}
		for _, feeAndPool := range pools.GetPools() {
			poolAddr := feeAndPool.Pool
			calls = append(calls, multicall.Multicall2Call{
				Target:   common.HexToAddress(poolAddr),
				CallData: uniswapv3Price,
			})
			uniswapv3Twap, err := v3ABI.Pack("observe", []uint32{0, 600})
			log.CheckFatal(err)
			calls = append(calls, multicall.Multicall2Call{
				Target:   common.HexToAddress(poolAddr),
				CallData: uniswapv3Twap,
			})
		}
		tokens = append(tokens, token)
	}
	mdl.tokens = tokens
	return
}

func (mdl UNIFetcher) setuniv2(entry multicall.Multicall2Result, univ2 map[string]float64, token string, tokenDecimals int8) {
	if !entry.Success || len(entry.ReturnData) == 0 {
		return
	}
	v2ABI := core.GetAbi("Uniswapv2Pool")
	value, err := v2ABI.Unpack("getReserves", entry.ReturnData)
	log.CheckFatal(err)
	r0 := value[0].(*big.Int)
	r1 := value[1].(*big.Int)
	v2Price := priceInBaseToken(token, mdl.baseToken, tokenDecimals, r0, r1)
	//
	univ2[token] = utils.GetFloat64Decimal(v2Price, mdl.baseDecimals)
}

func (mdl UNIFetcher) setuniv3(entry multicall.Multicall2Result, univ3 map[string]float64, token string, tokenDecimals int8) {
	if !entry.Success || len(entry.ReturnData) == 0 {
		return
	}
	v3ABI := core.GetAbi("Uniswapv3Pool")
	value, err := v3ABI.Unpack("slot0", entry.ReturnData)
	if err != nil {
		log.Fatalf("for token %s, %v err: %s", token, entry.ReturnData, err)
	}
	//https://docs.uniswap.org/sdk/guides/fetching-prices#understanding-sqrtprice
	// [(slot0**2 *Token0decimals)/2**192], divide by token for getting the float price in WETH
	//
	v3Price := univ3SlotToPriceInBase(value[0].(*big.Int), areSorted(token, mdl.baseToken), tokenDecimals)
	//
	univ3[token] = utils.GetFloat64Decimal(v3Price, mdl.baseDecimals)
}

// https://stackoverflow.com/questions/74555451/uniswap-v3-what-does-price-mean-at-a-given-tick
// pow = tick_diff/tick_spacing
// num=1.0001**pow is ratio of token1/token0 min unit
// if sorted, divdecimals = baseDecimals-tokenDecimals else tokenDecimals-baseDecimals
// (num/token1Decimals) * token0Decimals
// if sorted is ans else 1/ans
func (mdl UNIFetcher) settwapv3(entry multicall.Multicall2Result, twapV3 map[string]float64, token string, tokenDecimals int8) {
	if !entry.Success || len(entry.ReturnData) == 0 {
		return
	}
	v3ABI := core.GetAbi("Uniswapv3Pool")
	value, err := v3ABI.Unpack("observe", entry.ReturnData)
	log.CheckFatal(err)
	ticks := value[0].([]*big.Int)
	// https://medium.com/blockchain-development-notes/a-guide-on-uniswap-v3-twap-oracle-2aa74a4a97c5
	// (t1-t0)/interval
	tickDiff := new(big.Int).Quo(new(big.Int).Sub(ticks[0], ticks[1]), big.NewInt(600))
	sqrtPrice := powFloat(tickDiff)
	//
	divDecimals := mdl.baseDecimals - tokenDecimals
	sorted := areSorted(token, mdl.baseToken)
	//
	if !sorted {
		divDecimals = divDecimals * -1
	}
	//
	if divDecimals > 0 {
		sqrtPrice = new(big.Float).Quo(sqrtPrice, utils.GetExpFloat(divDecimals))
	} else {
		sqrtPrice = new(big.Float).Mul(sqrtPrice, utils.GetExpFloat(divDecimals*-1))
	}
	//
	if !sorted {
		sqrtPrice = new(big.Float).Quo(big.NewFloat(1), sqrtPrice)
	}
	twapV3Price, _ := sqrtPrice.Float64()
	twapV3[token] = twapV3Price
}

func (mdl *UNIFetcher) GetPrices(results []multicall.Multicall2Result, _ int64) (univ2 map[string]float64, univ3, twapv3 map[int16]map[string]float64) {
	univ2 = map[string]float64{}
	univ3, twapv3 = map[int16]map[string]float64{}, map[int16]map[string]float64{}
	//
	var resultInd int
	for _, token := range mdl.tokens {
		tokenDecimals := mdl.decimalsStore.GetDecimals(common.HexToAddress(token))
		//
		if mdl.UniPoolByToken[token].V2 != core.NULL_ADDR.Hex() {
			mdl.setuniv2(results[resultInd], univ2, token, tokenDecimals) // uniswap v2 price
			resultInd++
		}

		for _, feeAndPool := range mdl.UniPoolByToken[token].GetPools() {
			fee := feeAndPool.Fee
			//
			if univ3[fee] == nil {
				univ3[fee] = map[string]float64{}
			}
			mdl.setuniv3(results[resultInd], univ3[fee], token, tokenDecimals) // uniswap v3 price
			resultInd++
			//
			if twapv3[fee] == nil {
				twapv3[fee] = map[string]float64{}
			}
			mdl.settwapv3(results[resultInd], twapv3[fee], token, tokenDecimals) // uniswap v3 twap price
			resultInd++
		}
	}
	return
}

/////////////////////////
// UNI pools related methods
/////////////////////////

func areSorted(token, base string) bool {
	return strings.Compare(strings.ToLower(token), strings.ToLower(base)) == -1
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
