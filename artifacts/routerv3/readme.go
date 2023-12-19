package routerv3

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (_Routerv3 *Routerv3Caller) FindBestClosePath(opts *bind.CallOpts, creditAccount common.Address, expectedBalances []Balance, leftoverBalances []Balance, connectors []common.Address, slippage *big.Int, pathOptions []PathOption, loops *big.Int, force bool) (struct {
	Result            RouterResult
	GasPriceTargetRAY *big.Int
}, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "findBestClosePath", creditAccount, expectedBalances, leftoverBalances, connectors, slippage, pathOptions, loops, force)

	outstruct := new(struct {
		Result            RouterResult
		GasPriceTargetRAY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(RouterResult)).(*RouterResult)
	outstruct.GasPriceTargetRAY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}
