package priceFetcher

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"

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
	GetTokens() []string
	GetTopics() []common.Hash
	OnLog(txLog types.Log) bool
	//
	GetCalls(int64) []multicall.Multicall2Call
	GetPrices(ts int64, results []multicall.Multicall2Result, blockNum int64) map[string]*big.Int
	GetFeed(token string) common.Address
	GetFeedInfo(feed common.Address) *FeedInfo
	//
	//
	GetPFType(token common.Address, blockNum ...int64) int
	GetFeedForBlock(token common.Address, blockNum int64) common.Address
	HasReserveFeed(token string) bool
	GetFeedAndType(token string, reserve bool) (typeAndBlock, error)
}

type GearboxOracle struct {
	Address     common.Address
	tokenToFeed map[string]common.Address
	Node        *pkg.Node
	version     core.VersionType
	topics      []common.Hash
	//
	tokens []string
	//
}

func (GearboxOracle) GetFeedInfo(feed common.Address) *FeedInfo {
	return nil
}

func NewGearboxOracle(addr common.Address, version core.VersionType, client core.ClientI) GearboxOracleI {
	po := &GearboxOracle{
		Address:     addr,
		tokenToFeed: map[string]common.Address{},
		Node: &pkg.Node{
			Client: client,
		},
		version: version,
		topics:  []common.Hash{core.Topic("NewPriceFeed(address,address)")},
	}
	return po
}

func (pOracle GearboxOracle) HasReserveFeed(token string) bool {
	return false
}

func (pOracle GearboxOracle) GetTopics() []common.Hash {
	return pOracle.topics
}

func (pOracle GearboxOracle) GetPFType(token common.Address, blockNum ...int64) int {
	return core.V3_BACKEND_GENERAL_ORACLE
}

func (pOracle GearboxOracle) GetAddress() common.Address {
	return pOracle.Address
}

// overridden in gearbox_oracle_v3.go
func (pOracle GearboxOracle) GetFeed(token string) common.Address {
	return pOracle.tokenToFeed[token]
}
func (pOracle GearboxOracle) GetFeedAndType(token string, reserve bool) (typeAndBlock, error) {
	return typeAndBlock{}, nil
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

// overridden in gearbox_oracle_v3.go
func (pOracle *GearboxOracle) OnLog(txLog types.Log) bool {
	switch txLog.Topics[0] {
	case pOracle.topics[0]:
		token := common.HexToAddress(txLog.Topics[1].Hex()).Hex()
		feed := common.HexToAddress(txLog.Topics[2].Hex())
		pOracle.tokenToFeed[token] = feed
		return true
	}
	return false
}

func (pOracle GearboxOracle) GetFeedForBlock(token common.Address, blockNum int64) common.Address {
	return core.NULL_ADDR
}

func (pOracle *GearboxOracle) GetTokens() []string {
	tokens := make([]string, 0, len(pOracle.tokenToFeed))
	for token := range pOracle.tokenToFeed {
		tokens = append(tokens, token)
	}
	return tokens
}

// overridden in gearbox_oracle_v3.go
func (pOracle *GearboxOracle) GetCalls(_ int64) []multicall.Multicall2Call {
	poABI := core.GetAbi("YearnPriceFeed")
	data, err := poABI.Pack("latestRoundData")
	log.CheckFatal(err)
	//
	calls := make([]multicall.Multicall2Call, 0, len(pOracle.tokenToFeed))
	tokens := make([]string, 0, len(pOracle.tokenToFeed))
	for token, feed := range pOracle.tokenToFeed {
		tokens = append(tokens, token)
		calls = append(calls, multicall.Multicall2Call{
			Target:   feed,
			CallData: data,
		})
	}
	pOracle.tokens = tokens
	return calls
}

func (pOracle *GearboxOracle) GetPrices(ts int64, results []multicall.Multicall2Result, _ int64) map[string]*big.Int {
	defer utils.Elapsed("getprice gearbox oracle")()
	poABI := core.GetAbi("YearnPriceFeed")
	prices := map[string]*big.Int{}
	log.Info("total calls", len(results), "push price", len(results)-len(pOracle.tokens))
	log.Info(results[0])
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
	return prices
}
