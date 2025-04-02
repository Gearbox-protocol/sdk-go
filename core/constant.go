package core

import (
	"math/big"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

const LogFilterLenError = "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response."
const QueryMoreThan10000Error = "query returned more than 10000 results"
const LogFilterQueryTimeout = "Query timeout exceeded. Consider reducing your block range."
const NoderealFilterLogError = "exceed maximum block range:"
const AnkrRangeError = "block range is too wide"
const AnvilManagerError = "cannot_be_a_base"
const AclhemyExceedError = "Your app has exceeded its compute units per second capacity"
const InfuraError = "query returned more than 113 results"
const SECONDS_PER_YEAR = 86400 * 365

func BlockPer(c int64, d time.Duration) int64 {
	if d == time.Hour {
		return blockPerMin(c) * 60
	} else if d == time.Minute {
		return blockPerMin(c)
	}
	log.Fatalf("unsupported duration %v", d)
	return 0
}

func blockPerMin(chainId int64) int64 {
	net := log.GetBaseNet(chainId)
	switch net {
	case log.MAINNET:
		return 5
	case log.ARBITRUM:
		return 4 * 60
	case log.SONIC:
		return 2 * 60
	case log.OPTIMISM:
		return 4 * 60
	default:
		log.Fatalf("unsupported chainId %d", chainId)
	}
	return 0
}

func EthLogErrorCheck(err error, client ClientI) bool {
	if err != nil {
		if strings.Contains(err.Error(), QueryMoreThan10000Error) ||
			strings.Contains(err.Error(), NoderealFilterLogError) ||
			strings.Contains(err.Error(), AnvilManagerError) ||
			strings.Contains(err.Error(), AnkrRangeError) ||
			strings.Contains(err.Error(), AclhemyExceedError) ||
			strings.Contains(err.Error(), "exceed max topics") || // for anvil
			strings.Contains(err.Error(), LogFilterLenError) ||
			strings.Contains(err.Error(), InfuraError) ||
			(strings.Contains(err.Error(), "we can't execute this request") && GetChainId(client) == 42161) || // for arbitrum get logs for account Manager
			// failure: we can't execute this request range  192549019 192549555 tokenAddrs 32 accountHashes 6
			strings.Contains(err.Error(), LogFilterQueryTimeout) {
			return true
		}
	}
	return false
}

var WETHPrice, USDCPrice *big.Int

func init() {
	WETHPrice, _ = new(big.Int).SetString("1000000000000000000", 10)
	USDCPrice, _ = new(big.Int).SetString("100000000", 10)
}

var NULL_ADDR = common.Address{}

var MAX_BIG_INT = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

// for ray
const RAY_DECIMALS int8 = 27

var RAY *big.Int = utils.GetExpInt(RAY_DECIMALS)

// if 300 return all versions
func GetAddressProvider(chainId int64, version VersionType) string {
	addr := utils.GetEnvOrDefault("ADDRESS_PROVIDER", "")
	//
	var s string
	switch log.GetBaseNet(chainId) {
	case log.MAINNET:
		s = "0xcF64698AFF7E5f27A11dff868AF228653ba53be0,0x9ea7b04da02a5373317d745c1571c84aad03321d,0xBaB2014Dd88223E168bA06911c06df638311a097" // without v310
	case log.ARBITRUM:
		s = "0x7d04ecdb892ae074f03b5d0aba03796f90f3f2af"
	case log.OPTIMISM:
		s = "0x3761ca4bfacfcffc1b8034e69f19116dd6756726"
	case log.SONIC:
		s = "0x4b27b296273B72d7c7bfee1ACE93DC081467C41B"
	}
	if addr != "" {
		s += "," + addr
	}
	if version == NewVersion(300) {
		splits := strings.Split(s, ",")
		return splits[len(splits)-1]
	}
	return s
}

var WAD = utils.GetExpInt(18)
