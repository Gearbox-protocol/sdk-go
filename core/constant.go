package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const LogFilterLenError = "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response."
const QueryMoreThan10000Error = "query returned more than 10000 results"
const LogFilterQueryTimeout = "Query timeout exceeded. Consider reducing your block range."
const NoderealFilterLogError = "exceed maximum block range:"
const NoOfBlocksPerMin int64 = 5
const NoOfBlocksPerHr int64 = NoOfBlocksPerMin * 60

var WETHPrice, USDCPrice *big.Int

func init() {
	WETHPrice, _ = new(big.Int).SetString("1000000000000000000", 10)
	USDCPrice, _ = new(big.Int).SetString("100000000", 10)
}

type NetworkUI struct {
	ExplorerUrl string
	ChartUrl    string
}

var NULL_ADDR = common.Address{}

func NetworkUIUrl(chainId int64) NetworkUI {
	switch chainId {
	case 1:
		return NetworkUI{
			ExplorerUrl: "https://etherscan.io",
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
func (net NetworkUI) ExplorerAddrUrl(addr string) string {
	return fmt.Sprintf("%s/address/%s", net.ExplorerUrl, addr)
}
func (net NetworkUI) ExplorerHashUrl(txHash string) string {
	return fmt.Sprintf("%s/tx/%s", net.ExplorerUrl, txHash)
}

var MAX_BIG_INT = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
