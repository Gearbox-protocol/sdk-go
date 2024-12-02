package priceFetcher

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/calc"
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
	ArbBaseTokens []string `json:"arbBaseTokens"`
	OptBaseTokens []string `json:"optBaseTokens"`
	//
	crvLen             int
	generatedCalls     []multicall.Multicall2Call
	resolveArbToensToo bool
	//
	symToAddr  *core.SymTOAddrStore
	allSyms    []string
	inchOracle common.Address
	client     core.ClientI
	decimals   DecimalStoreI
	//
	extraurls *URLsAndResolve
}

func (o *OneInchOracle) Reset(net log.NETWORK) (all []string) {
	{
		baseReset := []string{}
		for _, baseToenSym := range o.BaseTokens {
			if _, ok := o.symToAddr.Tokens[baseToenSym]; ok {
				baseReset = append(baseReset, baseToenSym)
			}
		}
		o.BaseTokens = baseReset
		all = append(all, baseReset...)
	}
	if net != log.MAINNET {
		o.ConstToken = nil
	}
	{
		yearnReset := []YearnSpotPriceCalc{}
		for _, yearnToken := range o.YearnTokens {
			symToToken := o.symToAddr.Tokens[yearnToken.Token]
			if symToToken != core.NULL_ADDR {
				yearnReset = append(yearnReset, yearnToken)
				all = append(all, yearnToken.Token)
			}
		}
		o.YearnTokens = yearnReset
	}
	{
		crvReset := []CrvSpotPriceCalc{}
		for _, crvToken := range o.CrvTokens {
			symToToken := o.symToAddr.Tokens[crvToken.Token]
			if symToToken != core.NULL_ADDR {
				crvReset = append(crvReset, crvToken)
				all = append(all, crvToken.Token)
			}
		}
		o.CrvTokens = crvReset
	}
	return all
}

var MAINNET_GMX = common.HexToAddress("0x00eee00eee00eee00eee00eee00eee00eee00eee")
var MAINNET_OP = common.HexToAddress("0x00fff00fff00fff00fff00fff00fff00fff00fff")

func (details OneInchOracle) getAllTokens() (addrs []common.Address) {
	addrs = make([]common.Address, 0, len(details.symToAddr.Tokens))
	for sym, addr := range details.symToAddr.Tokens {
		if !utils.Contains([]common.Address{MAINNET_GMX, MAINNET_OP}, addr) && utils.Contains(details.allSyms, sym) {
			addrs = append(addrs, addr)
		}
	}
	return
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
	IsMaker    bool   `json:"isMaker"`
	Underlying string `json:"underlying"`
}

type DecimalStoreI interface {
	GetDecimals(tokenAddr common.Address) int8
	GetDecimalsForList([]common.Address)
}

func JsonnetStringInchConfig(syms []core.Symbol) string {
	var subPhrase []string
	for _, sym := range syms {
		subPhrase = append(subPhrase, fmt.Sprintf(`'%s'`, sym))
	}
	return fmt.Sprintf(`{baseTokens: [%s]}`, strings.Join(subPhrase, ","))
}

type URLsAndResolve struct {
	Resolve       bool   `json:"resolve"`
	Urls          string `json:"urls"`
	arbclient     core.ClientI
	mainclient    core.ClientI
	optclient     core.ClientI
	lastMainBlock int64
	result        map[string]*core.BigInt
}

// if resolve is set, set all 3 clients.
// else client for given network
func (d *URLsAndResolve) resolve(client core.ClientI) {
	if d.Resolve {
		d.mainclient = GetNetworkClient(d.Urls, log.MAINNET)
		d.arbclient = GetNetworkClient(d.Urls, log.ARBITRUM)
		d.optclient = GetNetworkClient(d.Urls, log.OPTIMISM)
	} else {
		switch log.GetBaseNet(core.GetChainId(client)) {
		case log.MAINNET:
			d.mainclient = client
		case log.OPTIMISM:
			d.optclient = client
		case log.ARBITRUM:
			d.arbclient = client
		}
	}
}

// supports only infura and alchemy
func New1InchOracle(client core.ClientI, tStore DecimalStoreI, details URLsAndResolve, dataStrings ...string) *OneInchOracle {
	calc := &OneInchOracle{}
	// get 1inch jsonnet
	data := func() string {
		if dataStrings == nil {
			data, err := core.GetEmbeddedJsonnet("1inch_price_calc_details.jsonnet", core.JsonnetImports{})
			log.CheckFatal(err)
			return data
		} else {
			data, err := core.GetJsonFromJsonnetData(dataStrings[0], core.JsonnetImports{})
			log.CheckFatal(err)
			return data
		}
	}()
	utils.SetJson([]byte(data), calc)

	chainId := core.GetChainId(client)
	calc.symToAddr = core.GetSymToAddrByChainId(chainId)
	calc.allSyms = calc.Reset(log.GetBaseNet(chainId))
	details.resolve(client)
	calc.extraurls = &details

	//
	// delete ETH as 0xee address can't be used for getting symbol in token_store.go
	delete(calc.symToAddr.Tokens, "ETH")
	//
	calc.inchOracle = get1InchAddress(chainId)
	calc.client = client
	calc.setCrvCallLen()

	// decimal store
	calc.decimals = tStore
	return calc
}

func get1InchAddress(chainId int64) common.Address {
	net := log.GetBaseNet(chainId)
	switch net {
	case log.MAINNET:
		return common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
	case log.ARBITRUM, log.OPTIMISM:
		return common.HexToAddress("0x0AdDd25a91563696D8567Df78D5A01C9a991F9B8")
	}
	log.Fatal("Can't get the inch oracle for", chainId)
	return core.NULL_ADDR
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
		switch log.GetBaseNet(core.GetChainId(calc.client)) {
		case log.ARBITRUM: // ARB_LOGIC
			calls = append(calls, calc.GetArbBaseCalls()...)
		case log.OPTIMISM:
			calls = append(calls, calc.GetOptBaseCalls()...)
		}
		calls = append(calls, calc.GetCrvCalls()...) // yvcurve-steth - yearn token is dependent on curve
		calls = append(calls, calc.GetYearnCalls()...)
		calc.generatedCalls = calls
	}
	return calc.generatedCalls
}

// yearn dependent on curve and base
// curve dependent on base
// const and base doesn't dependent on any token.
// for mainnet fetch all 3network tokens.
// for optimism, mainnet+optimism
// for arbitrum, mainnet+optimism
func (calc OneInchOracle) GetPrices(results []multicall.Multicall2Result, blockNumber int64, ts uint64) map[string]*core.BigInt {
	defer utils.Elapsed(fmt.Sprintf("GetPrices for block: %d", blockNumber))()
	if len(results) != len(calc.generatedCalls) {
		log.Fatalf("call len %d, result len %d", len(calc.generatedCalls), len(results))
	}
	calc.decimals.GetDecimalsForList(calc.getAllTokens())
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
	// ARB_LOGIC
	switch log.GetBaseNet(core.GetChainId(calc.client)) {
	case log.ARBITRUM:
		till = len(calc.ArbBaseTokens)
		arbResults, _results := results[:till], results[till:]
		calc.processSeparateBaseResults(arbResults, prices, calc.ArbBaseTokens)
		results = _results
	case log.OPTIMISM:
		till = len(calc.OptBaseTokens)
		optResults, _results := results[:till], results[till:]
		calc.processSeparateBaseResults(optResults, prices, calc.OptBaseTokens)
		results = _results
	case log.MAINNET:
		// ARB_LOGIC
		calc.arbForMainnet(ts, prices)
		calc.optForMainnet(ts, prices)
	}
	//
	till = calc.crvLen
	crvResults, results := results[:till], results[till:]
	calc.processCrvResults(crvResults, prices, blockNumber)
	//
	till = len(calc.YearnTokens)
	yearnResults, _ := results[:till], results[till:]
	calc.processYearnResults(yearnResults, prices)
	//
	for tokenSym, fromTokenSym := range calc.CopyPricesFor {
		fromToken := calc.symToAddr.Tokens[fromTokenSym]
		token := calc.symToAddr.Tokens[tokenSym]
		if fromToken != core.NULL_ADDR && token != core.NULL_ADDR {
			prices[token.Hex()] = core.NewBigInt(prices[fromToken.Hex()])
		}
	}

	// resolve arb tokens
	// get 1inch token from the 1inch spot contract use 1inch api
	calc.addGearPrice(prices)
	return prices
}

func (calc OneInchOracle) addGearPrice(prices map[string]*core.BigInt) {
	gearToken := calc.symToAddr.Tokens["GEAR"]
	if gearToken != core.NULL_ADDR {
		price := calc.getGearPriceFromCurve(prices[calc.symToAddr.Tokens["WETH"].Hex()].Convert())
		prices[gearToken.Hex()] = (*core.BigInt)(price)
	}
}

func (calc OneInchOracle) USDC() string {
	return calc.symToAddr.Tokens["USDC"].Hex()
}

// get the price from 1inch api for `token to usdc quote`
func (calc OneInchOracle) getPriceForAPI(tokenSym string) *big.Int {
	token := calc.symToAddr.Tokens[tokenSym]
	decimals := calc.decimals.GetDecimals(token)
	url := fmt.Sprintf("https://api.1inch.io/v5.0/1/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s",
		token.Hex(),
		calc.USDC(),
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

func (calc OneInchOracle) getGearPriceFromCurve(ethPriceInUSD *big.Int) *big.Int {
	// TODO
	pool := common.HexToAddress("0x0E9B5B092caD6F1c5E6bc7f89Ffe1abb5c95F1C2") // for mainnet
	d0Data, err := core.GetAbi("curveBalance").Pack("balances", big.NewInt(0))
	log.CheckFatal(err)
	d1Data, err := core.GetAbi("curveBalance").Pack("balances", big.NewInt(1))
	log.CheckFatal(err)
	results := core.MakeMultiCall(calc.client, 0, false, []multicall.Multicall2Call{{
		Target:   pool,
		CallData: d0Data,
	}, {
		Target:   pool,
		CallData: d1Data,
	}})
	if len(results[0].ReturnData) == 0 || len(results[1].ReturnData) == 0 || ethPriceInUSD.Cmp(new(big.Int)) == 0 {
		return new(big.Int)
	}
	gearBalance, ok := core.MulticallAnsBigInt(results[0])
	if !ok {
		log.Info("Can't get gear balance.")
		return new(big.Int)
	}
	ethBalance, ok := core.MulticallAnsBigInt(results[1])
	if !ok {
		log.Info("Can't get eth balance.")
		return new(big.Int)
	}
	return new(big.Int).Quo(
		new(big.Int).Mul(ethBalance, ethPriceInUSD),
		gearBalance,
	)
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
			if log.GetBaseNet(core.GetChainId(calc.client)) != log.MAINNET {
				prices[calc.symToAddr.Tokens["USDC_e"].Hex()] = (*core.BigInt)(big.NewInt(1000_000_00))
			}
		}
	}
}

func (calc OneInchOracle) processSeparateBaseResults(results []multicall.Multicall2Result, prices map[string]*core.BigInt, tokens []string) {
	for ind, entry := range results {
		tokenAddr := calc.symToAddr.Tokens[tokens[ind]] // this should be the token for the client network not neccessarily the arb network
		if entry.Success {
			price := new(big.Int).SetBytes(entry.ReturnData)
			// for usdt = 18-6-2 = 10
			// for wbtc = 18-8-2 = 8
			// for gusd = 18-2-2= 14
			var decimals int8 = 18
			if !utils.Contains([]common.Address{MAINNET_GMX, MAINNET_OP}, tokenAddr) {
				decimals = calc.decimals.GetDecimals(tokenAddr)
			}
			normalizeDecimal := 18 - decimals - 2
			price = utils.GetInt64(price, normalizeDecimal)
			prices[tokenAddr.Hex()] = (*core.BigInt)(price)
		}
	}
}

func (calc OneInchOracle) GetYearnCalls() (calls []multicall.Multicall2Call) {
	data, err := hex.DecodeString("99530b06") // pricepershare
	log.CheckFatal(err)
	makerABI := core.GetAbi("MakerDAI")
	makerData, err := makerABI.Pack("convertToAssets", utils.GetExpInt(18))
	log.CheckFatal(err)
	for _, details := range calc.YearnTokens {
		tokenAddr := calc.symToAddr.Tokens[details.Token]
		if details.IsMaker {
			calls = append(calls, multicall.Multicall2Call{
				Target:   tokenAddr,
				CallData: makerData,
			})
		} else {
			calls = append(calls, multicall.Multicall2Call{
				Target:   tokenAddr,
				CallData: data,
			})
		}
	}
	return
}

func (calc OneInchOracle) processYearnResults(results []multicall.Multicall2Result, prices map[string]*core.BigInt) {
	for ind, entry := range results {
		if entry.Success {
			tokenAddr := calc.symToAddr.Tokens[calc.YearnTokens[ind].Token]
			pricePerShare := new(big.Int).SetBytes(entry.ReturnData[:32])
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

func (o *OneInchOracle) GetCurrentPriceAtBlockNum(blockNum int64, bal core.DBBalanceFormat, underlying string) (float64, error) {
	tradingToken, baseToken := calc.TradingAndBaseTokens(core.GetChainId(o.client), bal, underlying)
	if tradingToken == "" {
		return 0, nil
	}
	//
	prices := map[string]*core.BigInt{}
	//
	{ // getting price via multicall
		inchOracle := get1InchAddress(core.GetChainId(o.client))
		tokens := []common.Address{common.HexToAddress(tradingToken), common.HexToAddress(baseToken)}
		usdc := o.symToAddr.Tokens["USDC"]
		//
		pfABI := core.GetAbi("1InchOracle")
		calls := []multicall.Multicall2Call{}
		for _, token := range tokens {
			data, err := pfABI.Pack("getRate", token, usdc, false)
			if err != nil {
				return 0, log.WrapErrWithLine(err)
			}
			calls = append(calls, multicall.Multicall2Call{
				Target:   inchOracle,
				CallData: data,
			})
		}
		results := core.MakeMultiCall(o.client, blockNum, false, calls)
		for ind, entry := range results {
			tokenAddr := tokens[ind]
			if entry.Success {
				price := new(big.Int).SetBytes(entry.ReturnData)
				// for usdt = 18-6-2 = 10
				// for wbtc = 18-8-2 = 8
				// for gusd = 18-2-2= 14
				normalizeDecimal := 18 - o.decimals.GetDecimals(tokenAddr) - 2
				price = utils.GetInt64(price, normalizeDecimal)
				prices[tokenAddr.Hex()] = (*core.BigInt)(price)
			} else if usdc == tokenAddr {
				prices[tokenAddr.Hex()] = (*core.BigInt)(big.NewInt(1000_000_00))
			}
		}
	}
	return GetTradingPriceFrom1Inch(prices, tradingToken, baseToken), nil
}

func GetTradingPriceFrom1Inch(prices map[string]*core.BigInt, tradingToken, baseToken string) float64 {
	getPrice := func(token string) *big.Int {
		p := prices[token]
		if p == nil {
			return new(big.Int)
		}
		return p.Convert()
	}
	tradingPrice := utils.GetFloat64Decimal(getPrice(tradingToken), 8)
	basePrice := utils.GetFloat64Decimal(getPrice(baseToken), 8)
	if tradingPrice == 0 || basePrice == 0 {
		return 0
	}
	return tradingPrice / basePrice
}
