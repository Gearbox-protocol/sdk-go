package routerv310

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (_Routerv3 *Routerv310Caller) RouteManyToOne(opts *bind.CallOpts, creditAccount common.Address, target common.Address, slippage *big.Int, tData []TokenData) (RouterResult, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "routeManyToOne", creditAccount, target, slippage, tData)

	outstruct := RouterResult{}
	if err != nil {
		return outstruct, err
	}

	outstruct = *abi.ConvertType(out[0], new(RouterResult)).(*RouterResult)

	return outstruct, err
}
