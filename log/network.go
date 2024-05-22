package log

import (
	"fmt"
	"log"
	"strings"
)

type NetworkUI struct {
	ExplorerUrl string
	ChartUrl    string
}

func (net NetworkUI) ExplorerAddrUrl(addr string) string {
	return fmt.Sprintf("%s/address/%s", net.ExplorerUrl, addr)
}
func (net NetworkUI) ExplorerHashUrl(txHash string) string {
	return fmt.Sprintf("%s/tx/%s", net.ExplorerUrl, txHash)
}

func NetworkUIUrl(chainId int64) NetworkUI {
	switch chainId {
	case 1, 7878, 1337:
		return NetworkUI{
			ExplorerUrl: "https://etherscan.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42161, 7880:
		return NetworkUI{
			ExplorerUrl: "https://arbitrum.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42: // kovan
		return NetworkUI{
			ExplorerUrl: "https://kovan.etherscan.io",
			ChartUrl:    "https://charts.kovan.gearbox.fi",
		}
	case 10, 7879: // optimism
		return NetworkUI{
			ExplorerUrl: "https://optimistic.etherscan.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 5: // goerli
		return NetworkUI{
			ExplorerUrl: "https://goerli.etherscan.io",
			ChartUrl:    "https://charts.goerli.gearbox.fi",
		}
	}
	return NetworkUI{}
}

type NETWORK string

var KOVAN NETWORK = "KOVAN"
var GOERLI NETWORK = "GOERLI"
var MAINNET NETWORK = "MAINNET"
var TEST NETWORK = "TEST"
var ANVIL NETWORK = "ANVIL"
var ARBITRUM NETWORK = "ARBITRUM"
var ARBTEST NETWORK = "ARBTEST"
var OPTIMISM NETWORK = "OPTIMISM"
var OPTTEST NETWORK = "OPTTEST"

func GetNetworkName(chainId int64) (name NETWORK) {
	switch chainId {
	case 42:
		name = KOVAN
	case 5:
		name = GOERLI
	case 1:
		name = MAINNET
	case 1337:
		name = TEST
	case 7878:
		name = ANVIL
	case 42161:
		name = ARBITRUM
	case 7880:
		name = ARBTEST
	case 10:
		name = OPTIMISM
	case 7879:
		name = OPTTEST
	}
	return
}

func GetBaseNet(chainId int64) NETWORK {
	net := GetNetworkName(chainId)
	if net == ANVIL {
		net = MAINNET
	} else if net == ARBTEST {
		net = ARBITRUM
	} else if net == TEST {
		net = MAINNET
	} else if net == OPTTEST {
		net = OPTIMISM
	}
	return net
}
func GetNetworkToChainId(net NETWORK) int64 {
	switch net {
	case MAINNET:
		return 1
	case KOVAN:
		return 42
	case GOERLI:
		return 5
	case ARBITRUM:
		return 42161
	case OPTIMISM:
		return 10
	default:
		log.Fatal("network to chainid not found", net)
	}
	return 0
}
func GetConfigFile(chainId int64) string {
	return strings.ToLower(string(GetBaseNet(chainId))) + ".jsonnet"
}
