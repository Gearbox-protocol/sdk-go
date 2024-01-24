package dc

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
)

var _symbol map[common.Address]core.Symbol

func IsWETH(client core.ClientI, underlying common.Address) bool {
	if _symbol == nil {
		chainId := core.GetChainId(client)
		_symbol = core.GetTokenToSymbolByChainId(chainId)
	}
	return _symbol[underlying] == "WETH"
}
