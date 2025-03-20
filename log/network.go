package log

import (
	"fmt"
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
	case 146, 7882: // goerli
		return NetworkUI{
			ExplorerUrl: "https://sonicscan.org",
			ChartUrl:    "https://charts.gearbox.fi",
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
var SONIC NETWORK = "SONIC"
var SONICTEST NETWORK = "SONICTEST"

var testnet = map[int64]struct {
	net  NETWORK
	base int64
}{
	42:   {KOVAN, 1},
	5:    {GOERLI, 1},
	1337: {TEST, 1},
	7878: {ANVIL, 1},
	7880: {ARBTEST, 42161},
	7879: {OPTTEST, 10},
	7882: {SONICTEST, 146},
}
var basenet = map[int64]struct {
	net  NETWORK
	test int64
}{
	1:     {MAINNET, 7878},
	42161: {ARBITRUM, 7880},
	10:    {OPTIMISM, 7879},
	146:   {SONIC, 7882},
}

func GetNetworkName(chainId int64) (name NETWORK) {
	if name, ok := testnet[chainId]; ok {
		return name.net
	}
	if name, ok := basenet[chainId]; ok {
		return name.net
	}
	Fatal("network not found", chainId)
	return
}

func GetBaseNet(chainId int64) NETWORK {
	if name, ok := testnet[chainId]; ok {
		chainId = name.base
	}
	// get for base
	if name, ok := basenet[chainId]; ok {
		return name.net
	}
	Fatal("network not found", chainId)
	return ""
}
func GetTestNet(chainId int64) NETWORK {
	if name, ok := basenet[chainId]; ok {
		chainId = name.test
	}
	// get for base
	if name, ok := testnet[chainId]; ok {
		return name.net
	}
	Fatal("network not found", chainId)
	return ""
}
func GetNetworkToChainId(netname NETWORK) int64 {
	for id, net := range testnet {
		if net.net == netname {
			return id
		}
	}
	for id, net := range basenet {
		if net.net == netname {
			return id
		}
	}
	Fatal("network not found", netname)
	return 0
}
func GetConfigFile(chainId int64) string {
	return strings.ToLower(string(GetBaseNet(chainId))) + ".jsonnet"
}
