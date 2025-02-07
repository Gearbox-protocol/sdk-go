package multicall3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (con *Multicall3Caller) TryAggregate(opts *bind.CallOpts, success bool, calls []Multicall3Call) ([]Multicall3Result, error) {
	var out []interface{}
	err := con.contract.Call(opts, &out, "tryAggregate", success, calls)

	if err != nil {
		return *new([]Multicall3Result), err
	}

	out0 := *abi.ConvertType(out[0], new([]Multicall3Result)).(*[]Multicall3Result)

	return out0, err

}
