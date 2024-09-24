package core

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type Multicall3 struct {
	*multicall3.Multicall3
}

func (con Multicall3) TryAggregate(opts *bind.CallOpts, successRequired bool, _calls []multicall.Multicall2Call) ([]multicall.Multicall2Result, error) {
	calls := []multicall3.Multicall3Call{}
	for _, call := range _calls {
		calls = append(calls, multicall3.Multicall3Call{
			Target:   call.Target,
			CallData: call.CallData,
		})
	}
	_results, err := con.Multicall3.TryAggregate(opts, successRequired, calls)
	if err != nil {
		return nil, err
	}
	results := []multicall.Multicall2Result{}
	for _, entry := range _results {
		results = append(results, multicall.Multicall2Result{
			Success:    entry.Success,
			ReturnData: entry.ReturnData,
		})
	}
	return results, err
}
