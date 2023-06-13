package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var versionABI string = "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func FetchVersion(addr string, blockNum int64, client ClientI) int16 {
	var opts *bind.CallOpts
	if blockNum != 0 {
		opts = &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
	}
	contract := bind.NewBoundContract(common.HexToAddress(addr), *GetAbi("Version"), client, client, client)
	var out []interface{}
	err := contract.Call(opts, &out, "version")
	if err != nil {
		return 1
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return int16(out0.Int64())
}
func FetchVersionOptimized(addr common.Address, blockNum int64, client ClientI) int16 {
	version, err := CallFuncWithExtraBytes(client, "54fd4d50", addr, blockNum, nil)
	if err != nil {
		return 1
	}
	return int16(new(big.Int).SetBytes(version).Int64())
}
