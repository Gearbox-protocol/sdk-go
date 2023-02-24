package priceFetcher

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"

	// "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	// "github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type GearboxOracleI interface {
	//
	GetPriceTokenTill(blockNum int64)
	GetAddress() common.Address
	GetVersion() int16
	GetTokenToFeed() map[string]common.Address
	OnLog(txLog types.Log) bool
	//
	GetCalls() []multicall.Multicall2Call
	GetPrices(results []multicall.Multicall2Result, blockNum int64) map[string]*big.Int
}

type GearboxOracle struct {
	Address     string
	TokenToFeed map[string]common.Address
	Node        *core.Node
	version     int16
	//
	tokens []string
}

func NewGearboxOracle(addr string, version int16, client core.ClientI) GearboxOracleI {
	po := &GearboxOracle{
		Address:     addr,
		TokenToFeed: map[string]common.Address{},
		Node: &core.Node{
			Client: client,
		},
		version: version,
	}
	return po
}

func (pOracle GearboxOracle) GetAddress() common.Address {
	return common.HexToAddress(pOracle.Address)
}

func (pOracle *GearboxOracle) GetVersion() int16 {
	return pOracle.version
}

// gets all the prie feeds add events
func (pOracle *GearboxOracle) GetPriceTokenTill(blockNum int64) {
	txLogs, err := pOracle.Node.GetLogs(0, blockNum,
		[]common.Address{common.HexToAddress(pOracle.Address)},
		[][]common.Hash{{core.Topic("NewPriceFeed(address,address)")}})
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		pOracle.OnLog(txLog)
	}
}

func (pOracle *GearboxOracle) OnLog(txLog types.Log) bool {
	switch txLog.Topics[0] {
	case core.Topic("NewPriceFeed(address,address)"):
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.TokenToFeed[token] = feed
		return true
	}
	return false
}

func (pOracle *GearboxOracle) GetTokenToFeed() map[string]common.Address {
	return pOracle.TokenToFeed
}

func (pOracle *GearboxOracle) GetCalls() []multicall.Multicall2Call {
	poABI := core.GetAbi("YearnPriceFeed")
	data, err := poABI.Pack("latestRoundData")
	log.CheckFatal(err)
	//
	calls := make([]multicall.Multicall2Call, 0, len(pOracle.TokenToFeed))
	tokens := make([]string, 0, len(pOracle.TokenToFeed))
	for token, feed := range pOracle.TokenToFeed {
		tokens = append(tokens, token)
		calls = append(calls, multicall.Multicall2Call{
			Target:   feed,
			CallData: data,
		})
	}
	pOracle.tokens = tokens
	return calls
}

func (pOracle *GearboxOracle) GetPrices(results []multicall.Multicall2Result, _ int64) map[string]*big.Int {
	poABI := core.GetAbi("YearnPriceFeed")
	prices := map[string]*big.Int{}
	for i, entry := range results {
		if entry.Success {
			value, err := poABI.Unpack("latestRoundData", entry.ReturnData)
			log.CheckFatal(err)
			// latestData := abi.ConvertType(value[0], new(LatestData)).(*LatestData)
			answer := value[1].(*big.Int)
			prices[pOracle.tokens[i]] = answer
		}
	}
	return prices
}
