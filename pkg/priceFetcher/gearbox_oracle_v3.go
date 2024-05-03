package priceFetcher

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type typeAndBlock struct {
	Type     int
	BlockNum int64
}
type GearboxOraclev3 struct {
	GearboxOracle
	tokenToReserve                 map[string]reserveUsage
	compositefeedToUnderlyingFeeds map[string]common.Address
	tokenToType                    map[common.Address]map[bool][]typeAndBlock
}

type reserveUsage struct {
	feed common.Address
	use  bool
}

func NewGearboxOraclev3(addr common.Address, version core.VersionType, client core.ClientI) GearboxOracleI {
	po := &GearboxOraclev3{
		GearboxOracle: GearboxOracle{
			Address:     addr,
			tokenToFeed: map[string]common.Address{},
			Node: &pkg.Node{
				Client: client,
			},
			topics: []common.Hash{
				core.Topic("SetPriceFeed(address,address,uint32,bool,bool)"),   // main
				core.Topic("SetReservePriceFeed(address,address,uint32,bool)"), // reserve
				core.Topic("SetReservePriceFeedStatus(address,bool)"),          // change
			},
			version: version,
		},
		tokenToReserve:                 map[string]reserveUsage{},
		tokenToType:                    map[common.Address]map[bool][]typeAndBlock{},
		compositefeedToUnderlyingFeeds: map[string]common.Address{},
	}
	return po
}

func (pOracle *GearboxOraclev3) addtokenToType(blockNum int64, feed common.Address, token common.Address, reserve bool) {
	typeData, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, "3fd0875f", feed, blockNum, []byte{}) // priceFeedType
	if err == nil {
		if pOracle.tokenToType[token] == nil {
			pOracle.tokenToType[token] = map[bool][]typeAndBlock{}
		}
		pfType := int(new(big.Int).SetBytes(typeData).Int64())
		obj := typeAndBlock{
			Type:     pfType,
			BlockNum: blockNum,
		}
		if pfType == core.V3_COMPOSITE_ORACLE {
			fn := func(sig string) common.Address {
				priceFeed0, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, sig, feed, blockNum, []byte{}) // priceFeedType
				log.CheckFatal(err)
				return common.BytesToAddress(priceFeed0)
			}
			//
			pf0 := fn("385aee1b")
			pf1 := fn("ab0ca0e1")
			pOracle.compositefeedToUnderlyingFeeds[feed.Hex()+"pf0"] = pf0
			pOracle.compositefeedToUnderlyingFeeds[feed.Hex()+"pf1"] = pf1
			//
			pf0Type, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, "3fd0875f", pf0, blockNum, []byte{})
			if err == nil {
				if new(big.Int).SetBytes(pf0Type).Int64() == core.V3_REDSTONE_ORACLE {
					obj.Type = core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE
				}
			}
		}
		pOracle.tokenToType[token][reserve] = append(pOracle.tokenToType[token][reserve], obj)
	}
}

func (pOracle GearboxOraclev3) GetUnderlyingPF(compfeed common.Address, pf0 bool) common.Address {
	if pf0 {
		return pOracle.compositefeedToUnderlyingFeeds[compfeed.Hex()+"pf0"]
	} else {
		return pOracle.compositefeedToUnderlyingFeeds[compfeed.Hex()+"pf1"]
	}
}

// gets all the prie feeds add events
func (pOracle GearboxOraclev3) GetPriceTokenTill(blockNum int64) {
	txLogs, err := pOracle.Node.GetLogs(0, blockNum,
		[]common.Address{pOracle.Address},
		[][]common.Hash{pOracle.topics})
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		pOracle.OnLog(txLog)
	}
}

func (pOracle GearboxOraclev3) getTypeAndBlock(token common.Address, blockNum ...int64) typeAndBlock {
	reserve := pOracle.tokenToReserve[token.Hex()].use
	typeAndBlocks := pOracle.tokenToType[token][reserve]
	l := len(typeAndBlocks)
	if len(blockNum) > 0 {
		for i := l - 1; i >= 0; i-- {
			if typeAndBlocks[i].BlockNum <= blockNum[0] {
				return typeAndBlocks[i]
			}
		}
	}
	if l == 0 {
		// log.Warnf("getTypeAndBlock: token %s has no typeAndBlocks: %+v", token.Hex(), blockNum)
		return typeAndBlock{}
	}
	return typeAndBlocks[l-1]
}

func (pOracle GearboxOraclev3) GetPFType(token common.Address, blockNum ...int64) int {
	data := pOracle.getTypeAndBlock(token, blockNum...)
	return data.Type
}

func (pOracle *GearboxOraclev3) OnLog(txLog types.Log) bool {
	// chainId := core.GetChainId(pOracle.Node.Client)
	// addrtosym := core.GetTokenToSymbolByChainId(chainId)
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case pOracle.topics[0]:
		token := common.HexToAddress(txLog.Topics[1].Hex())
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed, token, false)
		pOracle.tokenToFeed[token.Hex()] = feed
		return true
	case pOracle.topics[1]: // set reserve
		token := common.HexToAddress(txLog.Topics[1].Hex())
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed, token, true)
		pOracle.tokenToReserve[token.Hex()] = reserveUsage{feed: feed, use: pOracle.tokenToReserve[token.Hex()].use}
	case pOracle.topics[2]: // change from reserve to main feed, vice verse
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		pOracle.tokenToReserve[token] = reserveUsage{feed: pOracle.tokenToReserve[token].feed, use: txLog.Topics[2][:][63] == 1}
	}
	return false
}

func (pOracle GearboxOraclev3) GetFeed(token string) common.Address {
	if pOracle.tokenToReserve[token].use { // set reserve
		return pOracle.tokenToReserve[token].feed
	}
	return pOracle.tokenToFeed[token]
}

func (pOracle *GearboxOraclev3) GetCalls() []multicall.Multicall2Call {
	poABI := core.GetAbi("YearnPriceFeed")
	data, err := poABI.Pack("latestRoundData")
	log.CheckFatal(err)
	//
	calls := make([]multicall.Multicall2Call, 0, len(pOracle.tokenToFeed))
	tokens := make([]string, 0, len(pOracle.tokenToFeed))
	for token, feed := range pOracle.tokenToFeed {
		if pOracle.tokenToReserve[token].use { // set reserve
			feed = pOracle.tokenToReserve[token].feed
		}
		tokens = append(tokens, token)
		calls = append(calls, multicall.Multicall2Call{
			Target:   feed,
			CallData: data,
		})
	}
	pOracle.tokens = tokens
	return calls
}
