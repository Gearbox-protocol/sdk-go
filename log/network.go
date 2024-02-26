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
	case 1, 7878:
		return NetworkUI{
			ExplorerUrl: "https://etherscan.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42161:
		return NetworkUI{
			ExplorerUrl: "https://arbitrum.io",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42: // kovan
		return NetworkUI{
			ExplorerUrl: "https://kovan.etherscan.io",
			ChartUrl:    "https://charts.kovan.gearbox.fi",
		}
	case 5: // goerli
		return NetworkUI{
			ExplorerUrl: "https://goerli.etherscan.io",
			ChartUrl:    "https://charts.goerli.gearbox.fi",
		}
	}
	return NetworkUI{}
}

func GetNetworkName(chainId int64) (name string) {
	switch chainId {
	case 42:
		name = "KOVAN"
	case 5:
		name = "GOERLI"
	case 1:
		name = "MAINNET"
	case 1337:
		name = "TEST"
	case 7878:
		name = "ANVIL"
	case 42161:
		name = "ARBITRUM"
	case 43114:
		name = "ANVIL_ARBITRUM"
	}
	return
}

func GetBaseNet(chainId int64) string {
	net := GetNetworkName(chainId)
	if net == "ANVIL" {
		net = "MAINNET"
	} else if net == "ANVIL_ARBITRUM" {
		net = "ARBITRUM"
	}
	return net
}
func GetConfigFile(chainId int64) string {
	net := GetNetworkName(chainId)
	if net == "ANVIL" {
		net = "MAINNET"
	} else if net == "ANVIL_ARBITRUM" {
		net = "ARBITRUM"
	}
	return strings.ToLower(net) + ".jsonnet"
}