package dc

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
)

func IsWETH(client core.ClientI, underlying common.Address) bool {
	chainId := core.GetChainId(client)
	expectedWETHAddr := core.GetToken(chainId, "WETH")
	return underlying == expectedWETHAddr
}
