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
		// TODO: NEWNETWORK

	case 56, 7883: // bnb
		return NetworkUI{
			ExplorerUrl: "https://bscscan.com/",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 42793, 7884: // etherlink
		return NetworkUI{
			ExplorerUrl: "https://explorer.etherlink.com/",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 1135, 7885: // lisk
		return NetworkUI{
			ExplorerUrl: "https://blockscout.lisk.com/",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 43111, 7886: // hemi
		return NetworkUI{
			ExplorerUrl: "https://explorer.hemi.xyz/",
			ChartUrl:    "https://charts.gearbox.fi",
		}
	case 9745, 7887: // hemi
		return NetworkUI{
			ExplorerUrl: "https://plasmascan.to/",
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
var ARBITRUM NETWORK = "ARBITRUM"
var ANVIL NETWORK = "ANVIL"
var ARBTEST NETWORK = "ARBTEST"

// on cloudamqp for optimism
var OPTIMISM NETWORK = "OPTIMISM"
var OPTTEST NETWORK = "OPTTEST"
var SONIC NETWORK = "SONIC" // server with arbitrum
var SONICTEST NETWORK = "SONICTEST"
var BNB NETWORK = "BNB" // server with arbitrum
var BNBTEST NETWORK = "BNBTEST"

// on common
var ETHERLINK NETWORK = "ETHERLINK"
var ETHERLINKTEST NETWORK = "ETHERLINKTEST"
var LISK NETWORK = "LISK"
var HEMIBTC NETWORK = "HEMIBTC"
var LISKTEST NETWORK = "LISKTEST"
var HEMIBTCTEST NETWORK = "HEMIBTCTEST"

var PLASMA NETWORK = "PLASMA"
var PLASMATEST NETWORK = "PLASMATEST"

var MONAD NETWORK = "MONAD"
var MONADTEST NETWORK = "MONADTEST"

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
	// TODO: NEWNETWORK

	7883: {BNBTEST, 56},
	7884: {ETHERLINKTEST, 42793},
	7885: {LISKTEST, 1135}, // lisk testnet
	7886: {HEMIBTCTEST, 43111},
	7887: {PLASMATEST, 9745},
	7888: {MONADTEST, 143},
	// etherlink testnet
}
var Basenet = map[int64]struct {
	net  NETWORK
	test int64
}{
	1:     {MAINNET, 7878},
	42161: {ARBITRUM, 7880},
	10:    {OPTIMISM, 7879},
	146:   {SONIC, 7882},
	// TODO: NEWNETWORK

	56:    {BNB, 7883},
	42793: {ETHERLINK, 7884},
	1135:  {LISK, 7885}, // lisk mainnet
	43111: {HEMIBTC, 7886},
	9745:  {PLASMA, 7887},
	143:   {MONAD, 7888},
}
var BasenetToDB = map[int64]string{
	1:     "v310",
	42161: "gearbox_arb",
	10:    "gearbox_opt",
	146:   "gearbox_sonic",
	// TODO: NEWNETWORK

	56:    "gearbox_bnb",
	42793: "gearbox_etherlink",
	1135:  "gearbox_lisk",
	43111: "gearbox_hemibtc",
	9745:  "gearbox_plasma",
	143:   "gearbox_monad",
}

func GetNetworkName(chainId int64) (name NETWORK) {
	if name, ok := testnet[chainId]; ok {
		return name.net
	}
	if name, ok := Basenet[chainId]; ok {
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
	if name, ok := Basenet[chainId]; ok {
		return name.net
	}
	Fatal("network not found", chainId)
	return ""
}
func GetTestNet(chainId int64) NETWORK {
	if name, ok := Basenet[chainId]; ok {
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
	for id, net := range Basenet {
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
