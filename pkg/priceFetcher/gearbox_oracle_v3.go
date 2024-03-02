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

type GearboxOraclev3 struct {
	GearboxOracle
	tokenToReserve            map[string]reserveUsage
	compositefeedToPriceFeed0 map[common.Address]common.Address
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
		compositefeedToPriceFeed0: map[common.Address]common.Address{},
	}
	return po
}

func (pOracle *GearboxOraclev3) addFeedToType(feed common.Address) {
	typeData, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, "3fd0875f", feed, 0, []byte{}) // priceFeedType
	if err == nil {
		if int(new(big.Int).SetBytes(typeData).Int64()) == core.V3_COMPOSITE_ORACLE {
			priceFeed0, err := core.CallFuncWithExtraBytes(pOracle.Node.Client, "385aee1b", feed, 0, []byte{})
			if err == nil {
				pOracle.compositefeedToPriceFeed0[feed] = common.BytesToAddress(priceFeed0)
			}
		}
	}
}

func (pOracle GearboxOraclev3) GetPriceFeed0(compfeed common.Address) common.Address {
	return pOracle.compositefeedToPriceFeed0[compfeed]
}

func (pOracle *GearboxOraclev3) OnLog(txLog types.Log) bool {
	// chainId := core.GetChainId(pOracle.Node.Client)
	// addrtosym := core.GetTokenToSymbolByChainId(chainId)
	switch txLog.Topics[0] {
	case pOracle.topics[0]:
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addFeedToType(feed)
		pOracle.tokenToFeed[token] = feed
		return true
	case pOracle.topics[1]: // set reserve
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.addFeedToType(feed)
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
