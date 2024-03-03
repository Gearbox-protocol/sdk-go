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
	tokenToReserve            map[string]reserveUsage
	compositefeedToPriceFeed0 map[common.Address]common.Address
	tokenToType               map[common.Address][]typeAndBlock
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
		tokenToReserve:            map[string]reserveUsage{},
		tokenToType:               map[common.Address][]typeAndBlock{},
		compositefeedToPriceFeed0: map[common.Address]common.Address{},
	}
	return po
}

func (pOracle *GearboxOraclev3) addtokenToType(blockNum int64, feed common.Address) {
	typeData, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, "3fd0875f", feed, blockNum, []byte{}) // priceFeedType
	if err == nil {
		pOracle.tokenToType[feed] = append(pOracle.tokenToType[feed], typeAndBlock{
			int(new(big.Int).SetBytes(typeData).Int64()),
			blockNum,
		})
	}
}

func (pOracle GearboxOraclev3) GetPriceFeed0(compfeed common.Address) common.Address {
	return pOracle.compositefeedToPriceFeed0[compfeed]
}

func (pOracle GearboxOraclev3) getTypeAndBlock(token common.Address, blockNum ...int64) typeAndBlock {
	l := len(pOracle.tokenToType[token])
	if len(blockNum) > 0 {
		for i := l - 1; i >= 0; i-- {
			if pOracle.tokenToType[token][i].BlockNum <= blockNum[0] {
				return pOracle.tokenToType[token][i]
			}
		}
	}
	return pOracle.tokenToType[token][l-1]
}

func (pOracle GearboxOraclev3) IsRedStoneToken(token common.Address, blockNum ...int64) bool {
	data := pOracle.getTypeAndBlock(token, blockNum...)
	return data.Type == core.V3_REDSTONE_ORACLE
}

func (pOracle *GearboxOraclev3) OnLog(txLog types.Log) bool {
	// chainId := core.GetChainId(pOracle.Node.Client)
	// addrtosym := core.GetTokenToSymbolByChainId(chainId)
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case pOracle.topics[0]:
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed)
		pOracle.tokenToFeed[token] = feed
		return true
	case pOracle.topics[1]: // set reserve
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addtokenToType(blockNum, feed)
		pOracle.tokenToReserve[token] = reserveUsage{feed: feed, use: pOracle.tokenToReserve[token].use}
	case pOracle.topics[2]:
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
