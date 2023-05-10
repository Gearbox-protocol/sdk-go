package priceFetcher

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type OneInchOracle struct {
	ConstToken    []ConstTokenPriceCalc `json:"constTokens"`
	YearnTokens   []YearnSpotPriceCalc  `json:"yearnTokens"`
	BaseTokens    []string              `json:"baseTokens"`
	CrvTokens     []CrvSpotPriceCalc    `json:"crvTokens"`
	CopyPricesFor map[string]string     `json:"copyPricesFor"`
	//
	crvLen         int
	generatedCalls []multicall.Multicall2Call
	//
	symToAddr  *core.SymTOAddrStore
	inchOracle common.Address
	client     core.ClientI
	decimals   DecimalStoreI
}

type CrvSpotPriceCalc struct {
	Token            string   `json:"token"`
	Pool             string   `json:"pool"`
	UnderlyingTokens []string `json:"underlyingTokens"`
}
type ConstTokenPriceCalc struct {
	Token string `json:"token"`
	Price int64  `json:"prices"`
}
type YearnSpotPriceCalc struct {
	Token      string `json:"token"`
	Underlying string `json:"underlying"`
}

type DecimalStoreI interface {
	GetDecimals(tokenAddr common.Address) int8
}

func New1InchOracle(client core.ClientI, chainId int64, inchOracle common.Address, tStore DecimalStoreI) *OneInchOracle {
	calc := &OneInchOracle{}
	// get 1inch jsonnet
	data, err := core.GetEmbeddedJsonnet("1inch_price_calc_details.jsonnet", core.JsonnetImports{})
	log.CheckFatal(err)
	utils.SetJson([]byte(data), calc)
	//
	calc.symToAddr = core.GetSymToAddrByChainId(chainId)
	calc.inchOracle = inchOracle
	calc.client = client
	calc.setCrvCallLen()

	// decimal store
	calc.decimals = tStore
	return calc
}

func (calc *OneInchOracle) setCrvCallLen() {
	var crvLen int
	for _, details := range calc.CrvTokens {
		crvLen += len(details.UnderlyingTokens) + 1 // for balance of underlying, total supply of crv pool token
	}
	calc.crvLen = crvLen
	log.Info(calc)
}

func (calc *OneInchOracle) String() string {
	var totalLen int
	totalLen += len(calc.ConstToken)
	totalLen += len(calc.BaseTokens)
	totalLen += len(calc.YearnTokens)
	totalLen += calc.crvLen
	totalLen += len(calc.CopyPricesFor)
	//
	return fmt.Sprintf("Const :%d, Base: %d, Yearn: %d, Crv: %d, Cvx: %d,Total: %d",
		len(calc.ConstToken),
		len(calc.BaseTokens),
		len(calc.YearnTokens),
		calc.crvLen,
		len(calc.CopyPricesFor),
		totalLen,
	)
}

func (calc *OneInchOracle) GetCalls() []multicall.Multicall2Call {
	if calc.generatedCalls == nil {
		calls := []multicall.Multicall2Call{}
		calls = append(calls, calc.GetBaseCalls()...)
		calls = append(calls, calc.GetCrvCalls()...) // yvcurve-steth - yearn token is dependent on curve
		calls = append(calls, calc.GetYearnCalls()...)
		calc.generatedCalls = calls
	}
	return calc.generatedCalls
}

// yearn dependent on curve and base
// curve dependent on base
// const and base doesn't dependent on any token.
func (calc OneInchOracle) GetPrices(results []multicall.Multicall2Result, blockNumber int64) map[string]*core.BigInt {
	if len(results) != len(calc.generatedCalls) {
		log.Fatalf("call len %d, result len %d", len(calc.generatedCalls), len(results))
	}
	//
	prices := map[string]*core.BigInt{}
	//
	for _, constTokenDetails := range calc.ConstToken {
		tokenAddr := calc.symToAddr.Tokens[constTokenDetails.Token]
		prices[tokenAddr.Hex()] = (*core.BigInt)(big.NewInt(constTokenDetails.Price))
	}
	//
	till := len(calc.BaseTokens)
	baseResults, results := results[:till], results[till:]
	calc.processBaseResults(baseResults, prices)
	//
	till = calc.crvLen
	crvResults, results := results[:till], results[till:]
	calc.processCrvResults(crvResults, prices, blockNumber)
	//
	till = len(calc.YearnTokens)
	yearnResults, results := results[:till], results[till:]
	calc.processYearnResults(yearnResults, prices)
	//
	for tokenSym, fromTokenSym := range calc.CopyPricesFor {
		fromToken := calc.symToAddr.Tokens[fromTokenSym]
		token := calc.symToAddr.Tokens[tokenSym]
		prices[token.Hex()] = core.NewBigInt(prices[fromToken.Hex()])
	}
	// if the gearPrice is not available from the 1inch spot contract use 1inch api
	gearPrice := prices[calc.symToAddr.Tokens["GEAR"].Hex()]
	if gearPrice != nil && gearPrice.Convert().Cmp(new(big.Int)) == 0 { // only set gearPrice it if already set , but is zero
		price := calc.getPriceForAPI("GEAR")
		prices[calc.symToAddr.Tokens["GEAR"].Hex()] = (*core.BigInt)(price)
	}
	return prices
}

// get the price from 1inch api for `token to usdc quote`
func (calc OneInchOracle) getPriceForAPI(tokenSym string) *big.Int {
	token := calc.symToAddr.Tokens[tokenSym]
	decimals := calc.decimals.GetDecimals(token)
	url := fmt.Sprintf("https://api.1inch.io/v5.0/1/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s",
		token.Hex(),
		calc.symToAddr.Tokens["USDC"].Hex(),
		utils.GetExpInt(decimals+2), // 100 token units for price in 10^8 for usdc
	)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
	}
	val := struct {
		Error      string `json:"error"`
		StatusCode int    `json:"statusCode"`
		ToTokenAmt string `json:"toTokenAmount"`
	}{}
	if resp.StatusCode/100 != 2 { // 2xx
		return new(big.Int)
	}
	utils.ReadJsonReaderAndSetInterface(resp.Body, &val)
	if val.StatusCode != 0 {
		fmt.Println("Error", val.Error)
		return new(big.Int)
	}
	return utils.StringToInt(val.ToTokenAmt)
}

func (calc OneInchOracle) GetBaseCalls() (calls []multicall.Multicall2Call) {
	pfABI := core.GetAbi("1InchOracle")
	for _, token := range calc.BaseTokens {
		tokenAddr := calc.symToAddr.Tokens[token]
		data, err := pfABI.Pack("getRate", tokenAddr, calc.symToAddr.Tokens["USDC"], false)
		if err != nil {
			log.Fatal(err)
		}
		calls = append(calls, multicall.Multicall2Call{
			Target:   calc.inchOracle,
			CallData: data,
		})
	}
	return
}

func (calc OneInchOracle) processBaseResults(results []multicall.Multicall2Result, prices map[string]*core.BigInt) {
	for ind, entry := range results {
		tokenAddr := calc.symToAddr.Tokens[calc.BaseTokens[ind]]
		if entry.Success {
			price := new(big.Int).SetBytes(entry.ReturnData)
			// for usdt = 18-6-2 = 10
			// for wbtc = 18-8-2 = 8
			// for gusd = 18-2-2= 14
			normalizeDecimal := 18 - calc.decimals.GetDecimals(tokenAddr) - 2
			price = utils.GetInt64(price, normalizeDecimal)
			prices[tokenAddr.Hex()] = (*core.BigInt)(price)
		} else if calc.BaseTokens[ind] == "USDC" {
			prices[tokenAddr.Hex()] = (*core.BigInt)(big.NewInt(1000_000_00))
		}
	}
}

func (calc OneInchOracle) GetYearnCalls() (calls []multicall.Multicall2Call) {
	data, err := hex.DecodeString("99530b06") // pricepershare
	if err != nil {
		log.Fatal(err)
	}
	for _, details := range calc.YearnTokens {
		tokenAddr := calc.symToAddr.Tokens[details.Token]
		calls = append(calls, multicall.Multicall2Call{
			Target:   tokenAddr,
			CallData: data,
		})
	}
	return
}

func (calc OneInchOracle) processYearnResults(results []multicall.Multicall2Result, prices map[string]*core.BigInt) {
	for ind, entry := range results {
		if entry.Success {
			pricePerShare := new(big.Int).SetBytes(entry.ReturnData[:32])
			tokenAddr := calc.symToAddr.Tokens[calc.YearnTokens[ind].Token]
			underlyingAddr := calc.symToAddr.Tokens[calc.YearnTokens[ind].Underlying]
			price := new(big.Int).Mul(pricePerShare, prices[underlyingAddr.Hex()].Convert())
			prices[tokenAddr.Hex()] = (*core.BigInt)(utils.GetInt64(price,
				calc.decimals.GetDecimals(tokenAddr))) // div by 10**d due to pricepershare decimals
		}
	}
}

func (calc OneInchOracle) GetCrvCalls() (calls []multicall.Multicall2Call) {
	erc20ABI := core.GetAbi("Token")
	supplyData, err := erc20ABI.Pack("totalSupply")
	if err != nil {
		log.Fatal("Crv pool method's totalSupply encoding failed: ", err)
	}
	//
	for _, details := range calc.CrvTokens {
		//
		poolAddr := calc.symToAddr.Exchanges[details.Pool]
		for _, underlyingSym := range details.UnderlyingTokens {
			underlyingAddr := calc.symToAddr.Tokens[underlyingSym]
			//
			data, err := erc20ABI.Pack("balanceOf", poolAddr)
			if err != nil {
				log.Fatal(err)
			}
			calls = append(calls, multicall.Multicall2Call{
				Target:   underlyingAddr,
				CallData: data,
			})
		}
		//
		tokenAddr := calc.symToAddr.Tokens[details.Token]
		calls = append(calls, multicall.Multicall2Call{
			Target:   tokenAddr,
			CallData: supplyData,
		})
	}
	return
}

// (balanceOfUnderlying*priceUnderlying/10**DecimalUnderlying)*10**mainTokenDecimals/mainTokentotalSupply
func (calc OneInchOracle) processCrvResults(results []multicall.Multicall2Result,
	prices map[string]*core.BigInt, blockNumber int64) {
	ind := 0
	for _, details := range calc.CrvTokens {
		tokenAddr := calc.symToAddr.Tokens[details.Token]
		//
		totalValue := new(big.Int)
		for _, underlyingSym := range details.UnderlyingTokens {
			underlyingAddr := calc.symToAddr.Tokens[underlyingSym]
			var balanceOfPool *big.Int
			if underlyingSym == "WETH" {
				blockBigInt := big.NewInt(blockNumber)
				if blockNumber == 0 {
					blockBigInt = nil
				}
				// get eth balance of crv pool
				balance, err := calc.client.BalanceAt(
					context.TODO(),
					calc.symToAddr.Exchanges[details.Pool],
					blockBigInt,
				)
				log.CheckFatal(err)
				balanceOfPool = balance
			} else {
				entry := results[ind]
				if entry.Success {
					balanceOfPool = new(big.Int).SetBytes(entry.ReturnData)
				} else {
					log.Fatalf("crv get result failed for %s", utils.ToJson(details))
				}
			}
			//
			underlyingValue := utils.GetInt64(
				new(big.Int).Mul(prices[underlyingAddr.Hex()].Convert(), balanceOfPool),
				calc.decimals.GetDecimals(underlyingAddr),
			)

			totalValue = new(big.Int).Add(totalValue, underlyingValue)
			ind++
		}
		//
		totalSupplyData := results[ind]
		if totalSupplyData.Success {
			totalSupply := new(big.Int).SetBytes(totalSupplyData.ReturnData[:32])
			totalValue := utils.GetInt64(totalValue, -1*calc.decimals.GetDecimals(tokenAddr))
			//
			price := new(big.Int).Quo(totalValue, totalSupply)
			prices[tokenAddr.Hex()] = (*core.BigInt)(price)
		}
		ind++
	}
}
