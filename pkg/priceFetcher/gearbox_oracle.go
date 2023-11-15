package priceFetcher

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"

	// "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	// "github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type GearboxOracleI interface {
	//
	GetPriceTokenTill(blockNum int64)
	GetAddress() common.Address
	GetVersion() core.VersionType
	GetTokenToFeed() map[string]common.Address
	OnLog(txLog types.Log) bool
	//
	GetCalls() []multicall.Multicall2Call
	GetPrices(results []multicall.Multicall2Result, blockNum int64) map[string]*big.Int
}

type GearboxOracle struct {
	Address     common.Address
	TokenToFeed map[string]common.Address
	Node        *pkg.Node
	version     core.VersionType
	topics      []common.Hash
	//
	tokens []string
}

func NewGearboxOracle(addr common.Address, version core.VersionType, client core.ClientI) GearboxOracleI {
	po := &GearboxOracle{
		Address:     addr,
		TokenToFeed: map[string]common.Address{},
		Node: &pkg.Node{
			Client: client,
		},
		version: version,
		topics:  []common.Hash{core.Topic("NewPriceFeed(address,address)")},
	}
	return po
}

func (pOracle GearboxOracle) GetAddress() common.Address {
	return pOracle.Address
}

func (pOracle *GearboxOracle) GetVersion() core.VersionType {
	return pOracle.version
}

// gets all the prie feeds add events
func (pOracle *GearboxOracle) GetPriceTokenTill(blockNum int64) {
	txLogs, err := pOracle.Node.GetLogs(0, blockNum,
		[]common.Address{pOracle.Address},
		[][]common.Hash{pOracle.topics})
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		pOracle.OnLog(txLog)
	}
}

func (pOracle *GearboxOracle) OnLog(txLog types.Log) bool {
	switch txLog.Topics[0] {
	case pOracle.topics[0]:
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
			if err != nil {
				log.Fatalf("for token %s, err: %s", pOracle.tokens[i], err)
			}
			// latestData := abi.ConvertType(value[0], new(LatestData)).(*LatestData)
			answer := value[1].(*big.Int)
			prices[pOracle.tokens[i]] = answer
		}
	}
	return prices
}
