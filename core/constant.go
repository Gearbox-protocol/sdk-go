package core

import (
	"math/big"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

const LogFilterLenError = "Log response size exceeded. You can make eth_getLogs requests with up to" // previously it was 2k now 10k
const QueryMoreThan10000Error = "query returned more than 10000 results"
const LogFilterQueryTimeout = "Query timeout exceeded. Consider reducing your block range."
const NoderealFilterLogError = "exceed maximum block range:"
const AnkrRangeError = "block range is too wide"
const AnvilManagerError = "cannot_be_a_base"
const AclhemyExceedError = "Your app has exceeded its compute units per second capacity"
const Anvil10kError = "You can make eth_getLogs requests with up to a 10000 block range"
const InfuraError = "query returned more than 113 results"
const SECONDS_PER_YEAR = 86400 * 365

var DrpcFreeTierError = "ranges over 10000 blocks are not supported on freetier"
var RequestTimeoutDrpcFreeTierError = "Request timeout on the free tier"
var DrpcError = "query exceeds max block range 100000"
var BlockRangeChainStackError = "Block range limit exceeded."

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
	case log.BNB:
		return 30 // per min
	default:
		// TODO: NEWNETWORK
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
			strings.Contains(err.Error(), Anvil10kError) ||
			strings.Contains(err.Error(), AclhemyExceedError) ||
			strings.Contains(err.Error(), "exceed max topics") || // for anvil
			strings.Contains(err.Error(), LogFilterLenError) ||
			strings.Contains(err.Error(), DrpcError) ||
			strings.Contains(err.Error(), DrpcFreeTierError) ||
			strings.Contains(err.Error(), RequestTimeoutDrpcFreeTierError) ||
			strings.Contains(err.Error(), InfuraError) ||
			strings.Contains(err.Error(), BlockRangeChainStackError) ||
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
	addrProds := GetAddressProviderDS(chainId)
	//
	if version != NewVersion(1) {
		log.Fatal("Address provider version is not supported, use 0 for all versions or 300 for latest")
	}
	return addrProds.All()
}

var WAD = utils.GetExpInt(18)

type addrProviderV struct {
	Address string `json:"address"`
	Version int64  `json:"version"`
}

type AddrProviderV struct {
	x []addrProviderV
}

func (a AddrProviderV) check() {
	if len(a.x) == 0 {
		log.Fatal("Address provider is empty")
	}
}

// func (x AddrProviderV) First() string {
// 	x.check()
// 	return x.x[0].Address
// }

//	func (x AddrProviderV) First() string {
//		x.check()
//		return x.x[0].Address
//	}
func (x AddrProviderV) last() common.Address {
	x.check()
	if len(x.x) == 0 {
		log.Fatal("Address provider is empty")
	}
	return common.HexToAddress(x.x[len(x.x)-1].Address)
}

func (x AddrProviderV) MoreThanEq(vt int64) (ans []common.Address) {
	if vt != 300 && vt != 310 {
		log.Fatal("Address provider version is not supported, use 300 for latest or 310 for latest with more features")
	}
	x.check()
	if len(x.x) == 0 {
		log.Fatal("Address provider is empty")
	}
	// ans = []common.Address{}
	for _, v := range x.x {
		if v.Version >= vt {
			ans = append(ans, common.HexToAddress(v.Address))
		}
	}
	return ans
}

// Last returns the last address in the provider list, which is usually the latest version.
func (x AddrProviderV) Liquidators() (ans []common.Address) {
	// return []common.Address{x.last()}

	return x.MoreThanEq(310)
}

func (x AddrProviderV) All() string {
	x.check()
	var sb strings.Builder
	for i, v := range x.x {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(v.Address)
	}
	return sb.String()
}

func GetAddressProviderDS(chainId int64) AddrProviderV {
	var addrProviders []addrProviderV
	switch log.GetBaseNet(chainId) {
	case log.MAINNET:
		addrProviders = []addrProviderV{
			{Address: "0xcF64698AFF7E5f27A11dff868AF228653ba53be0", Version: 1},
			{Address: "0x9ea7b04da02a5373317d745c1571c84aad03321d", Version: 300},
			{Address: "0xBaB2014Dd88223E168bA06911c06df638311a097", Version: 310},
			{Address: "0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38", Version: 310}, // for wsteth 0xc168343C791D56
		}
	case log.ARBITRUM:
		addrProviders = []addrProviderV{
			{Address: "0x7d04ecdb892ae074f03b5d0aba03796f90f3f2af", Version: 300},
			{Address: "0xBaB2014Dd88223E168bA06911c06df638311a097", Version: 310},
		}
	case log.OPTIMISM:
		addrProviders = []addrProviderV{
			{Address: "0x3761ca4bfacfcffc1b8034e69f19116dd6756726", Version: 300},
			{Address: "0xBaB2014Dd88223E168bA06911c06df638311a097", Version: 310},
		}
	case log.SONIC:
		addrProviders = []addrProviderV{
			{Address: "0x4b27b296273B72d7c7bfee1ACE93DC081467C41B", Version: 300},
			{Address: "0xBaB2014Dd88223E168bA06911c06df638311a097", Version: 310},
		}
	case log.BNB:
		addrProviders = []addrProviderV{
			{Address: "0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38", Version: 310},
		}
	}
	if addr := utils.GetEnvOrDefault("ADDRESS_PROVIDER", ""); addr != "" {
		addrProviders = append(addrProviders, addrProviderV{Address: addr, Version: 310})
	}
	return AddrProviderV{addrProviders}
}

func GetMarketConfigurators(chainId int64) []common.Address {
	var markets []common.Address
	if market := utils.GetEnvOrDefault("MARKET_CONFIGURATORS", ""); market != "" {
		for _, addr := range strings.Split(market, ",") {
			markets = append(markets, common.HexToAddress(addr))
		}
		// return markets
	}
	switch log.GetBaseNet(chainId) { // check if supported
	case log.MAINNET:
		markets = append(markets, []common.Address{
			common.HexToAddress("0x354fe9f450F60b8547f88BE042E4A45b46128a06"),
			common.HexToAddress("0x4d427D418342d8CE89a7634c3a402851978B680A"), // 30)0
			common.HexToAddress("0x3b56538833fc02f4f0e75609390f26ded0c32e42"), // 310 mc for tbtc
			common.HexToAddress("0xc168343C791D56dD1Da4b4B8B0cc1C1EC1A16E6B"), // for wseth
		}...)
	case log.ARBITRUM:
		markets = append(markets, []common.Address{
			common.HexToAddress("0x01023850b360b88de0d0f84015bbba1eba57fe7e"),
		}...)
	case log.OPTIMISM:
		markets = append(markets, []common.Address{
			common.HexToAddress("0x2a15969CE5320868eb609680751cF8896DD92De5"),
		}...)
	case log.SONIC:
		markets = append(markets, []common.Address{
			common.HexToAddress("0x8FFDd1F1433674516f83645a768E8900A2A5D076"),
		}...)
	case log.BNB:
		markets = append(markets, []common.Address{
			common.HexToAddress("0x19037a281025b83fa37e3264b77af523ff87a3a4"),
		}...)
	}
	// log.Fatal("Market configurators not supported for chainId", chainId)
	return markets
}
