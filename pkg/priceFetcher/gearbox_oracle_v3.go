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
				core.Topic("SetPriceFeed(address,address,uint32,bool"),         // main v310
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
	FeedToken     common.Address
	SignThreshold int
	DataId        string
}

func (info FeedInfo) GetRedstonePF() core.RedStonePF {
	return core.RedStonePF{
		Type:             info.Type,
		DataServiceId:    "redstone-primary-prod",
		DataId:           info.DataId,
		SignersThreshold: info.SignThreshold,
		UnderlyingToken:  info.FeedToken,
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
func (pOracle *GearboxOraclev3) GetPF01AndFeedType(token common.Address, feed common.Address, blockNum int64, client core.ClientI) {
	if val, ok := pOracle.feedToInfo[feed]; ok {
		if val.Type == core.V3_REDSTONE_ORACLE && val.FeedToken != token && token != core.NULL_ADDR { // same redstone feed can be added for multiple tokens, use the latest one.
			val.FeedToken = token
		}
		return
	}
	fn := func(_feed common.Address, sig string) common.Address {
		priceFeed0, err := core.CallFuncWithExtraBytes(client, sig, _feed, 0, []byte{}) // priceFeedType
		if err != nil {
			log.Fatal(_feed, sig, err)
		}
		return common.BytesToAddress(priceFeed0)
	}
	typeData, err := core.CallFuncWithExtraBytes(client, "3fd0875f", feed, 0, []byte{}) // priceFeedType
	if err == nil {
		pfType := int(new(big.Int).SetBytes(typeData).Int64())
		obj := &FeedInfo{
			typeAndBlock: typeAndBlock{
				Type:     pfType,
				BlockNum: blockNum,
				Feed:     feed,
			},
		}
		if pfType == core.V3_COMPOSITE_ORACLE {
			//
			pf0 := fn(feed, "385aee1b") // priceFeed0
			pf1 := fn(feed, "ab0ca0e1") // priceFeed1
			obj.PF0 = pf0
			obj.PF1 = pf1
			obj.DecimalsPF0 = func() int8 {
				decimals, err := core.CallFuncWithExtraBytes(client, "313ce567", pf0, 0, []byte{}) // decimals
				log.CheckFatal(err)
				return int8(new(big.Int).SetBytes(decimals).Int64())
			}()
			//
			pf0Type, err := core.CallFuncWithExtraBytes(client, "3fd0875f", pf0, 0, []byte{})
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
					if obj.FeedToken == core.NULL_ADDR {
						log.Fatalf("pf0(%s) for composite %s dones't have known ticker token", pf0, feed)
					}
				}
			}
		} else if pfType == core.V3_PENDLE_PT_TWAP_ORACLE {
			obj.PF0 = fn(feed, "741bef1a") // priceFeed
			pOracle.GetPF01AndFeedType(core.NULL_ADDR, obj.PF0, blockNum, client)
			if utils.Contains([]int{core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE, core.V3_REDSTONE_ORACLE}, pOracle.GetFeedInfo(obj.PF0).Type) {
				obj.Type = core.V3_PULL_UNDERLYING_ORACLE
			}
		} else if pfType == core.V3_ERC4626_VAULT_ORACLE { // for stkUSDS on mainnet.
			obj.PF0 = fn(feed, "741bef1a")  // priceFeed
			lpToken := fn(feed, "5fcbd285") // lpToken
			pOracle.GetPF01AndFeedType(lpToken, obj.PF0, blockNum, client)
			if utils.Contains([]int{core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE, core.V3_REDSTONE_ORACLE}, pOracle.GetFeedInfo(obj.PF0).Type) {
				obj.Type = core.V3_PULL_UNDERLYING_ORACLE
			}
		} else if pfType == core.V3_REDSTONE_ORACLE { // onChainToken is not directly used as the onChainToken returned for sUSDS (redstoneToken feed) is DAI,but should be sUSDS. bcz for DAI the feed in priceorcle is not the sUSDS redstone feed.
			onChainToken, signThreshold, dataId := RedstoneDetails(feed, pOracle.Node.Client)
			if token == core.NULL_ADDR {
				token = onChainToken
			}
			//
			obj.FeedToken = token
			obj.SignThreshold = signThreshold
			obj.DataId = dataId
			//
		}
		pOracle.feedToInfo[feed] = obj
	} else {
		obj := &FeedInfo{
			typeAndBlock: typeAndBlock{
				Type:     core.V3_CHAINLINK_ORACLE,
				BlockNum: blockNum,
				Feed:     feed,
			},
		}
		pOracle.feedToInfo[feed] = obj
	}
}

func (pOracle *GearboxOraclev3) addtokenToType(blockNum int64, feed common.Address, token common.Address, reserve bool) {
	if pOracle.tokenToType[token] == nil {
		pOracle.tokenToType[token] = map[bool][]typeAndBlock{}
	}
	pOracle.GetPF01AndFeedType(token, feed, blockNum, pOracle.Node.Client)
	//
	info := pOracle.feedToInfo[feed]
	pOracle.tokenToType[token][reserve] = append(pOracle.tokenToType[token][reserve], info.typeAndBlock)
}

// gets all the prie feeds add events
func (pOracle GearboxOraclev3) GetPriceTokenTill(blockNum int64) {
	txLogs, err := pOracle.Node.GetLogs(0, blockNum,
		[]common.Address{common.HexToAddress(string(pOracle.Address))},
		[][]common.Hash{pOracle.topics})
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
	data, err := poABI.Pack("latestRoundData")
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
		if utils.Contains([]int{core.V3_PULL_UNDERLYING_ORACLE, core.V3_REDSTONE_ORACLE, core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE}, feedInfo.Type) {
			if feedInfo.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE {
				continue
			}
			if core.V3_PULL_UNDERLYING_ORACLE == feedInfo.Type {
				feedInfo = pOracle.GetFeedInfo(feedInfo.PF0)
			}
			ans := pOracle.mgr.GetPodSignWithRedstoneToken(ts, core.RedStonePF{
				Type:             15,
				DataServiceId:    "redstone-primary-prod",
				DataId:           feedInfo.DataId,
				SignersThreshold: feedInfo.SignThreshold,
				UnderlyingToken:  feedInfo.FeedToken,
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
		}
		tokens = append(tokens, token)
		calls = append(calls, multicall.Multicall2Call{
			Target:   feed,
			CallData: data,
		})
	}
	_ = data
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
			log.Info(i, "failed token", pOracle.tokens[i], "failed feed", pOracle.tokenToFeed[pOracle.tokens[i]])
		}
	}
	pOracle.AddCompsite(ts, prices)
	return prices
}

func (pOracle *GearboxOraclev3) AddCompsite(ts int64, prices map[string]*big.Int) {
	for token, feed := range pOracle.tokenToFeed {
		info := pOracle.GetFeedInfo(feed)
		if info.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE {
			price := pOracle.mgr.GetPrice(ts, info.GetRedstonePF())
			weth, wbtc := func() (string, string) {
				tokens := core.GetSymToAddrByChainId(core.GetChainId(pOracle.Node.Client))
				return tokens.Tokens["WETH"].Hex(), tokens.Tokens["WBTC"].Hex()
			}()
			if pOracle.GetFeedForETHBTC(weth) == info.PF1 {
				wethPrice := prices[weth]
				prices[token] = utils.GetInt64(new(big.Int).Mul(price, wethPrice), info.DecimalsPF0)
			} else if pOracle.GetFeedForETHBTC(wbtc) == info.PF1 {
				wbtcPrice := prices[wbtc]
				prices[token] = utils.GetInt64(new(big.Int).Mul(price, wbtcPrice), info.DecimalsPF0)
			} else {
				log.Warn("composite redstone price feed 1 is not weth/wbtc", token, utils.ToJson(info))
				prices[token] = new(big.Int)
			}
		}
	}
	return
}

func (pOracle *GearboxOraclev3) GetFeedForETHBTC(wbtc string) common.Address {
	feed := pOracle.GetFeed(wbtc)
	info := pOracle.GetFeedInfo(feed)
	if info.Type == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE || info.Type == core.V3_COMPOSITE_ORACLE {
		return info.PF1
	}
	return feed
}
