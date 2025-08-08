package priceFetcher

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/redstone"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	redstonemgr "github.com/Gearbox-protocol/sdk-go/pkg/redstone"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type typeAndBlock struct {
	Type     int
	BlockNum int64
	Feed     common.Address
}

type compositeFeedDetails struct {
	PF0      common.Address
	PF1      common.Address
	Decimals int8
}
type GearboxOraclev3 struct {
	GearboxOracle
	tokenToType  map[common.Address]map[bool][]typeAndBlock
	feedToInfo   map[common.Address]*FeedInfo
	feedToTicker map[common.Address]common.Address // feedToTicker
	mgr          redstonemgr.RedStoneMgrI
}

type ReserveUsage struct {
	Feed common.Address
	Use  bool
}

func (pOracle GearboxOraclev3) HasReserveFeed(token string) bool {
	data := pOracle.tokenToType[common.HexToAddress(token)][true]
	return len(data) > 0
}

// only redstone oracles
func (pOracle GearboxOraclev3) GetPullOracles() (ans []core.RedStonePF) {
	for _, feed := range pOracle.feedToInfo {
		if feed.Type == core.V3_REDSTONE_ORACLE {
			ans = append(ans, feed.GetRedstonePF())
		}
	}
	return
}

func (pOracle GearboxOraclev3) GetFeedAndType(token string, reserve bool) (typeAndBlock, error) {
	data := pOracle.tokenToType[common.HexToAddress(token)][reserve]
	if len(data) == 0 {
		return typeAndBlock{}, fmt.Errorf("GetMainFeed: token %s has no main feed", token)
	}
	return data[len(data)-1], nil
}
func NewGearboxOraclev3(addr schemas.PriceOracleT, version core.VersionType, client core.ClientI) GearboxOracleI {
	po := &GearboxOraclev3{
		GearboxOracle: GearboxOracle{
			Address:     addr,
			tokenToFeed: map[string]common.Address{},
			Node: &pkg.Node{
				Client: client,
			},
			topics: []common.Hash{
				core.Topic("SetPriceFeed(address,address,uint32,bool,bool)"),   // main v3
				core.Topic("SetPriceFeed(address,address,uint32,bool)"),        // main v310
				core.Topic("SetReservePriceFeed(address,address,uint32,bool)"), // reserve
				core.Topic("SetReservePriceFeedStatus(address,bool)"),          // change
			},
			version: version,
		},
		tokenToType:  map[common.Address]map[bool][]typeAndBlock{},
		feedToInfo:   map[common.Address]*FeedInfo{},
		mgr:          redstonemgr.NewRedStoneMgr(client),
		feedToTicker: map[common.Address]common.Address{},
	}
	return po
}

type FeedInfo struct {
	typeAndBlock
	//
	PF0           common.Address
	PF1           common.Address
	DecimalsPF0   int8
	FeedToken     common.Address // for redstone
	SignThreshold int            // for redstone
	DataId        string         // for redstone
}

func (info FeedInfo) GetRedstonePF() core.RedStonePF {
	feed := info.Feed
	if info.PF0 != core.NULL_ADDR {
		feed = info.PF0
	}
	return core.RedStonePF{
		Type:             info.Type,
		DataServiceId:    "redstone-primary-prod",
		DataId:           info.DataId,
		SignersThreshold: info.SignThreshold,
		UnderlyingToken:  info.FeedToken,
		Feed:             feed,
	}
}

func RedstoneDetails(feed common.Address, client core.ClientI) (feedToken common.Address, signThreshold int, dataId string) {
	contract, err := redstone.NewRedstone(feed, client)
	log.CheckFatal(err)
	feedToken, err = contract.Token(nil)
	log.CheckFatal(err)
	_signThreshold, err := contract.GetUniqueSignersThreshold(nil)
	log.CheckFatal(err)
	signThreshold = int(_signThreshold)
	//
	dataIdBytes, err := contract.DataFeedId(nil)
	log.CheckFatal(err)
	dataId = func() string {
		var s []byte
		for _, b := range dataIdBytes {
			if b != 0 {
				s = append(s, b)
			} else {
				break
			}
		}
		return string(s)
	}()
	return
}
func (pOracle *GearboxOraclev3) GetPF01AndFeedType(feed common.Address, blockNum int64, client core.ClientI, pythtoken common.Address) {
	// if val, ok := pOracle.feedToInfo[feed]; ok {
	// 	if val.Type == core.V3_REDSTONE_ORACLE && val.FeedToken != token && token != core.NULL_ADDR { // same redstone feed can be added for multiple tokens, use the latest one.
	// 		val.FeedToken = token
	// 	}
	// 	return
	// }
	fn := func(_feed common.Address, sig string) common.Address {
		priceFeed0, err := core.CallFuncGetSingleValue(client, sig, _feed, 0, []byte{}) // priceFeedType
		if err != nil {
			log.Fatal(_feed, sig, err)
		}
		return common.BytesToAddress(priceFeed0)
	}
	//
	pfType, err := core.GetGearboxPfType(client, feed.Hex(), "") // token is for logging purposes only.
	log.CheckFatal(err)
	obj := &FeedInfo{
		typeAndBlock: typeAndBlock{
			Type:     int(pfType),
			BlockNum: blockNum,
			Feed:     feed,
		},
	}
	// typeData, err := core.CallFuncGetSingleValue(client, "3fd0875f", feed, 0, []byte{}) // priceFeedType
	if pfType == core.V3_COMPOSITE_ORACLE {
		//
		pf0 := fn(feed, "385aee1b") // priceFeed0
		pf1 := fn(feed, "ab0ca0e1") // priceFeed1
		obj.PF0 = pf0
		obj.PF1 = pf1
		obj.DecimalsPF0 = func() int8 {
			decimals, err := core.CallFuncGetSingleValue(client, "313ce567", pf0, 0, []byte{}) // decimals
			log.CheckFatal(err)
			return int8(new(big.Int).SetBytes(decimals).Int64())
		}()
		//
		pf0Type, err := core.CallFuncGetSingleValue(client, "3fd0875f", pf0, 0, []byte{})
		if err == nil {
			if new(big.Int).SetBytes(pf0Type).Int64() == core.V3_REDSTONE_ORACLE {
				_, signThreshold, dataId := RedstoneDetails(pf0, pOracle.Node.Client)
				//
				obj.SignThreshold = signThreshold
				obj.DataId = dataId
				obj.Type = core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE
				obj.FeedToken = pOracle.feedToTicker[pf0]
				if pf0.Hex() == "0x14497e822B70554537dB9950126461C23dC4f237" {
					obj.FeedToken = common.HexToAddress("0x07299E4E806e4253727084c0493fFDf6fB2dBa3D")
				}
				//
				// SONIC_TEST
				if obj.FeedToken == core.NULL_ADDR && core.GetBaseChainId(pOracle.Node.Client) == 146 {
					obj.FeedToken = common.HexToAddress("0x56a5b6267d6b8de8ade88455b9342787e49e2f1a") // stS ticker token on sonic
				}
				if obj.FeedToken == core.NULL_ADDR {
					log.Info(utils.ToJson(pOracle.feedToTicker))
					log.Fatalf("pf0(%s) for composite %s dones't have known ticker token. priceoracle %s", pf0, feed, pOracle.Address)
				}
			}
		}
	} else if pfType == core.V3_PENDLE_PT_TWAP_ORACLE {
		obj.PF0 = fn(feed, "741bef1a") // priceFeed
		pOracle.GetPF01AndFeedType(obj.PF0, blockNum, client, core.NULL_ADDR)
		if utils.Contains([]int{core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE, core.V3_REDSTONE_ORACLE}, pOracle.GetFeedInfo(obj.PF0).Type) {
			obj.Type = core.V3_PULL_UNDERLYING_ORACLE
		}
	} else if pfType == core.V3_ERC4626_VAULT_ORACLE { // for stkUSDS on mainnet.
		obj.PF0 = fn(feed, "741bef1a") // priceFeed
		// lpToken := fn(feed, "5fcbd285") // lpToken
		pOracle.GetPF01AndFeedType(obj.PF0, blockNum, client, core.NULL_ADDR)
		if utils.Contains([]int{core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE, core.V3_REDSTONE_ORACLE}, pOracle.GetFeedInfo(obj.PF0).Type) {
			obj.Type = core.V3_PULL_UNDERLYING_ORACLE
		}
	} else if pfType == core.V3_REDSTONE_ORACLE { // onChainToken is not directly used as the onChainToken returned for sUSDS (redstoneToken feed) is DAI,but should be sUSDS. bcz for DAI the feed in priceorcle is not the sUSDS redstone feed.
		onChainToken, signThreshold, dataId := RedstoneDetails(feed, pOracle.Node.Client)
		// if token == core.NULL_ADDR {
		// 	token = onChainToken
		// }
		//
		obj.FeedToken = onChainToken
		obj.SignThreshold = signThreshold
		obj.DataId = dataId
		//
	} else if pfType == core.V3_PYTH_ORACLE { // external oracle
		data, err := core.CallFuncGetSingleValue(client, "1999bb9e", feed, blockNum, nil) // dataId
		log.CheckFatal(err)
		obj.DataId = common.BytesToHash(data).Hex()
		obj.FeedToken = pythtoken
	} else if pfType == core.V3_CHAINLINK_ORACLE { // chainlink oracle
		obj = &FeedInfo{
			typeAndBlock: typeAndBlock{
				Type:     core.V3_CHAINLINK_ORACLE,
				BlockNum: blockNum,
				Feed:     feed,
			},
		}
		// } else {
		// 	log.Info(pfType)
		// 	return
	}
	if obj.DataId == "beraETH_FUNDAMENTAL" {
		obj.DataId = "beraSTONE_FUNDAMENTAL"
	}
	pOracle.feedToInfo[feed] = obj
}

func (pOracle *GearboxOraclev3) addtokenToType(blockNum int64, feed common.Address, token common.Address, reserve bool) {
	if pOracle.tokenToType[token] == nil {
		pOracle.tokenToType[token] = map[bool][]typeAndBlock{}
	}
	// pOracle.GetPF01AndFeedType(token, feed, blockNum, pOracle.Node.Client)
	pOracle.GetPF01AndFeedType(feed, blockNum, pOracle.Node.Client, token)
	//
	info := pOracle.feedToInfo[feed]
	if pOracle.tokenToType[token] == nil {
		pOracle.tokenToType[token] = map[bool][]typeAndBlock{}
	}
	pOracle.tokenToType[token][reserve] = append(pOracle.tokenToType[token][reserve], info.typeAndBlock)
}

// gets all the prie feeds add events
func (pOracle GearboxOraclev3) GetPriceTokenTill(blockNum int64) {
	txLogs, err := pOracle.Node.GetLogs(0, blockNum,
		[]common.Address{common.HexToAddress(string(pOracle.Address))},
		[][]common.Hash{pOracle.topics}, true)
	log.CheckFatal(err)
	// feedToTicker
	for _, txLog := range txLogs {
		if pOracle.topics[0] == txLog.Topics[0] || pOracle.topics[1] == txLog.Topics[0] {
			token := common.HexToAddress(txLog.Topics[1].Hex())
			feed := common.HexToAddress(txLog.Topics[2].Hex())
			pOracle.feedToTicker[feed] = token
		}
	}
	//
	for _, txLog := range txLogs {
		pOracle.OnLog(txLog)
	}
}

func (pOracle GearboxOraclev3) getTypeAndBlock(token common.Address, blockNum ...int64) typeAndBlock {
	// reserve := pOracle.tokenToReserve[token.Hex()].Use
	reverse := false
	typeAndBlocks := pOracle.tokenToType[token][reverse]
	l := len(typeAndBlocks)
	if l == 0 {
		// log.Warnf("getTypeAndBlock: token %s has no typeAndBlocks: %+v", token.Hex(), blockNum)
		return typeAndBlock{}
	}
	if len(blockNum) > 0 && blockNum[0] != 0 {
		for i := l - 1; i >= 0; i-- {
			if typeAndBlocks[i].BlockNum <= blockNum[0] {
				return typeAndBlocks[i]
			}
		}
	}
	return typeAndBlocks[l-1]
}

func (pOracle GearboxOraclev3) GetPFType(token common.Address, blockNum ...int64) int {
	data := pOracle.getTypeAndBlock(token, blockNum...)
	return data.Type
}
func (pOracle GearboxOraclev3) GetFeedForBlock(token common.Address, blockNum int64) common.Address {
	data := pOracle.getTypeAndBlock(token, blockNum)
	return data.Feed
}

func (pOracle *GearboxOraclev3) OnLog(txLog types.Log) bool {
	// chainId := core.GetChainId(pOracle.Node.Client)
	// addrtosym := core.GetTokenToSymbolByChainId(chainId)
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case pOracle.topics[0], pOracle.topics[1]:
		token := common.HexToAddress(txLog.Topics[1].Hex())
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed, token, false)
		pOracle.tokenToFeed[token.Hex()] = feed
		return true
	case pOracle.topics[2]: // set reserve
		token := common.HexToAddress(txLog.Topics[1].Hex())
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed, token, true)
	case pOracle.topics[3]: // change from reserve to main feed, vice verse
		// token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		// pOracle.tokenToReserve[token] = ReserveUsage{Feed: pOracle.tokenToReserve[token].Feed, Use: txLog.Topics[2][:][63] == 1}
	}
	return false
}

func (pOracle GearboxOraclev3) GetFeed(token string) common.Address {
	// if pOracle.tokenToReserve[token].Use { // set reserve
	// 	return pOracle.tokenToReserve[token].Feed
	// }
	return pOracle.tokenToFeed[token]
}

func (pOracle *GearboxOraclev3) GetFeedInfo(feed common.Address) *FeedInfo {
	return pOracle.feedToInfo[feed]
}

func (pOracle *GearboxOraclev3) GetCalls(ts int64) []multicall.Multicall2Call {
	poABI := core.GetAbi("YearnPriceFeed")
	latestRunDataBytes, err := poABI.Pack("latestRoundData")
	log.CheckFatal(err)
	//
	calls := make([]multicall.Multicall2Call, 0, len(pOracle.tokenToFeed))
	pushPrices := []multicall.Multicall2Call{}
	//
	tokens := make([]string, 0, len(pOracle.tokenToFeed))
	updateABI := core.GetAbi("UpdatePriceFeed")
	for token, feed := range pOracle.tokenToFeed {
		//
		feedInfo := pOracle.GetFeedInfo(feed)
		//
		switch feedInfo.Type {
		case core.V3_PULL_UNDERLYING_ORACLE, core.V3_REDSTONE_ORACLE:
			if core.V3_PULL_UNDERLYING_ORACLE == feedInfo.Type {
				feedInfo = pOracle.GetFeedInfo(feedInfo.PF0)
			}
			ans := pOracle.mgr.GetPodSignWithRedstoneToken(ts, core.RedStonePF{
				Type:             15,
				DataServiceId:    "redstone-primary-prod",
				DataId:           feedInfo.DataId,
				SignersThreshold: feedInfo.SignThreshold,
				UnderlyingToken:  feedInfo.FeedToken,
				Feed:             feedInfo.Feed,
			})
			data, err := updateABI.Pack("updatePrice", ans.CallData)
			log.CheckFatal(err)
			//
			target := feedInfo.Feed
			// if feedInfo.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE {
			// 	continue
			// }
			//
			pushPrices = append(pushPrices, multicall.Multicall2Call{
				Target:   target,
				CallData: data,
			})
		case core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE: // The belief here is that for these composite oracles, the underlying will be some redstone oracle which will already be registered in Gearbox and patched through other data.
			continue
		case core.V3_PYTH_ORACLE:
			dataObj, err := pkg.GetPythPrice(feedInfo.DataId, ts)
			log.CheckFatal(err)
			data, err := updateABI.Pack("updatePrice", dataObj.Data)
			log.CheckFatal(err)
			pushPrices = append(pushPrices, multicall.Multicall2Call{
				Target:   feed,
				CallData: data,
			})
		}
		tokens = append(tokens, token)
		calls = append(calls, multicall.Multicall2Call{
			Target:   feed,
			CallData: latestRunDataBytes,
		})
	}
	pOracle.tokens = tokens
	return append(pushPrices, calls...)
	// return calls
}

func (pOracle *GearboxOraclev3) GetPrices(ts int64, results []multicall.Multicall2Result, _ int64) map[string]*big.Int {
	defer utils.Elapsed("getprice gearbox oracle")()
	poABI := core.GetAbi("YearnPriceFeed")
	prices := map[string]*big.Int{}
	log.Info("total calls", len(results), "push price", len(results)-len(pOracle.tokens))
	for i, entry := range results[len(results)-len(pOracle.tokens):] {
		if entry.Success {
			value, err := poABI.Unpack("latestRoundData", entry.ReturnData)
			if err != nil {
				log.Fatalf("for token %s, feed: %s err: %s", pOracle.tokens[i], pOracle.tokenToFeed[pOracle.tokens[i]], err)
			}
			// latestData := abi.ConvertType(value[0], new(LatestData)).(*LatestData)
			answer := value[1].(*big.Int)
			prices[pOracle.tokens[i]] = answer
		} else {
			feed := pOracle.tokenToFeed[pOracle.tokens[i]]
			info := pOracle.feedToInfo[feed]
			if info.PF0 != core.NULL_ADDR {
				priceData, err := core.CallFuncGetSingleValue(pOracle.Node.Client, "50d25bcd", info.PF0, 0, []byte{}) // priceFeed0
				if err == nil {
					prices[pOracle.tokens[i]] = new(big.Int).SetBytes(priceData)
				} else {
					log.Info("failed to get price for", pOracle.tokens[i], "feed's pf0", info.PF0, "err", err)
				}
			}
		}
	}
	pOracle.AddCompsite(ts, prices)
	return prices
}

func (pOracle *GearboxOraclev3) AddCompsite(ts int64, prices map[string]*big.Int) {
	chainId := core.GetChainId(pOracle.Node.Client)
	// tokens := core.GetSymToAddrByChainId(core.GetChainId(pOracle.Node.Client))
	weth, wbtc := func() (common.Address, common.Address) {
		a, _ := core.GetTokenWithError(chainId, "WETH")
		b, _ := core.GetTokenWithError(chainId, "WBTC")
		return a, b
	}()
	for token, feed := range pOracle.tokenToFeed {
		info := pOracle.GetFeedInfo(feed)
		if info.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE {
			price := pOracle.mgr.GetPrice(ts, info.GetRedstonePF())
			if pOracle.GetFeedForETHBTC("WETH", weth) == info.PF1 {
				wethPrice := prices[weth.Hex()]
				prices[token] = utils.GetInt64(new(big.Int).Mul(price, wethPrice), info.DecimalsPF0)
			} else if pOracle.GetFeedForETHBTC("WBTC", wbtc) == info.PF1 {
				wbtcPrice := prices[wbtc.Hex()]
				if wbtcPrice == nil {
					if utils.GetEnvOrDefault("OPTIMISTIC_LIQUIDATION", "") == "1" {
						prices[token] = new(big.Int)
					}
				} else {
					prices[token] = utils.GetInt64(new(big.Int).Mul(price, wbtcPrice), info.DecimalsPF0)
				}
			} else {
				switch log.GetBaseNet(core.GetChainId(pOracle.Node.Client)) {
				case log.SONIC:
					if wS := core.GetToken(chainId, "wS"); pOracle.GetFeedForETHBTC("wS", wS) == info.PF1 {
						wSPrice := prices[wS.Hex()]
						prices[token] = utils.GetInt64(new(big.Int).Mul(price, wSPrice), info.DecimalsPF0)
					}
				case log.ETHERLINK:
					if WXTZ := core.GetToken(chainId, "WXTZ"); pOracle.GetFeedForETHBTC("WXTZ", WXTZ) == info.PF1 {
						WXTZPrice := prices[WXTZ.Hex()]
						prices[token] = utils.GetInt64(new(big.Int).Mul(price, WXTZPrice), info.DecimalsPF0)
					}
				case log.BNB:
					if wbnb := core.GetToken(chainId, "WBNB"); pOracle.GetFeedForETHBTC("WBNB", wbnb) == info.PF1 {
						wbnbPrice := prices[wbnb.Hex()]
						prices[token] = utils.GetInt64(new(big.Int).Mul(price, wbnbPrice), info.DecimalsPF0)
					}
				default:
					log.Warn("composite redstone price feed 1 is not wS/weth/wbtc", token, utils.ToJson(info))
					prices[token] = new(big.Int)
				}
			}
		}
	}
}

func (pOracle *GearboxOraclev3) GetFeedForETHBTC(sym string, token common.Address) common.Address {
	// token := tokens[sym]
	if token == core.NULL_ADDR { // for sonic
		log.Debug("GetFeedForETHBTC feed not found for ", token, sym)
		return core.NULL_ADDR
	}
	feed := pOracle.GetFeed(token.Hex())
	info := pOracle.GetFeedInfo(feed)
	if info == nil {
		log.Debug("GetFeedForETHBTC feed not found for ", token, sym)
		return core.NULL_ADDR
	}
	if info.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE || info.Type == core.V3_COMPOSITE_ORACLE {
		return info.PF1
	}
	return feed
}

func (pOracle GearboxOraclev3) AddTokens(tokens []common.Address) {
	for _, token := range tokens {
		if len(pOracle.tokenToType[token][false]) != 0 {
			continue
		}
		tokenhash := common.BytesToHash(token[:])
		zero := common.HexToHash("0x0")
		_v3Main, err := core.CallFuncGetSingleValue(pOracle.Node.Client, "ff299845", pOracle.Address.Hex(), 0, append(tokenhash[:], zero[:]...)) // pricefeedraw
		v3MainAddr := common.BytesToAddress(_v3Main)
		if err != nil || v3MainAddr == core.NULL_ADDR {
			pOracle.v310AddToken(token)
		} else { // main and reserve
			pOracle.tokenToFeed[token.Hex()] = v3MainAddr
			// log.Info("set main", token)
			pOracle.addtokenToType(0, v3MainAddr, token, false) // feed
			one := common.HexToHash("0x1")
			_v3Reserve, err := core.CallFuncGetSingleValue(pOracle.Node.Client, "ff299845", pOracle.Address.Hex(), 0, append(tokenhash[:], one[:]...)) // pricefeedraw
			v3ReserveAddr := common.BytesToAddress(_v3Reserve)
			// log.Info("set reserve", token, err, v3ReserveAddr.Hex())
			if err == nil && v3ReserveAddr != core.NULL_ADDR {
				pOracle.addtokenToType(0, v3ReserveAddr, token, true)
			}
		}
	}
	// log.Infof("%v", pOracle.tokenToType)
}

func (pOracle GearboxOraclev3) v310AddToken(token common.Address) {
	hash := common.BytesToHash(token[:])
	// 9dcb511a
	v3Main, err := core.CallFuncGetSingleValue(pOracle.Node.Client, "9dcb511a", pOracle.Address.Hex(), 0, hash[:]) // pricefeedraw
	v3MainAddr := common.BytesToAddress(v3Main)
	if v3MainAddr != core.NULL_ADDR && err == nil {
		pOracle.feedToTicker[v3MainAddr] = token // v3MainAddr is the feed address // same is the ticker token
		pOracle.tokenToFeed[token.Hex()] = v3MainAddr
		pOracle.addtokenToType(0, v3MainAddr, token, false) // feed
	} else {
		return
	}
	//
	v3MReverse, err := core.CallFuncGetSingleValue(pOracle.Node.Client, "7c70dd51", pOracle.Address.Hex(), 0, hash[:]) // pricefeedraw reserve
	v3MReverseAddr := common.BytesToAddress(v3MReverse)
	if v3MReverseAddr != core.NULL_ADDR && err == nil {
		pOracle.addtokenToType(0, v3MReverseAddr, token, true) // feed
	}
}

func (pOracle *GearboxOraclev3) LoadFeedToTicker(_feedToTicker map[string]string) {
	feedToTicker := map[common.Address]common.Address{}
	for feed, ticker := range _feedToTicker {
		feedToTicker[common.HexToAddress(feed)] = common.HexToAddress(ticker)
	}
	pOracle.feedToTicker = feedToTicker
}
