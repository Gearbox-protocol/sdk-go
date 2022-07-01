// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lidoKovan

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LidoKovanABI is the input ABI used to generate the binding from.
const LidoKovanABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"syncer_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getLastCompletedReportDelta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"postTotalPooledEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preTotalPooledEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeElapsed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postTotalPooledEtherSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preTotalPooledEtherSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_postTotalPooledEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_preTotalPooledEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeElapsed\",\"type\":\"uint256\"}],\"name\":\"syncOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncer\",\"outputs\":[{\"internalType\":\"contractSyncer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeElapsedSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LidoKovan is an auto generated Go binding around an Ethereum contract.
type LidoKovan struct {
	LidoKovanCaller     // Read-only binding to the contract
	LidoKovanTransactor // Write-only binding to the contract
	LidoKovanFilterer   // Log filterer for contract events
}

// LidoKovanCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoKovanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoKovanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoKovanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoKovanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoKovanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoKovanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoKovanSession struct {
	Contract     *LidoKovan        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoKovanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoKovanCallerSession struct {
	Contract *LidoKovanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// LidoKovanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoKovanTransactorSession struct {
	Contract     *LidoKovanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LidoKovanRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoKovanRaw struct {
	Contract *LidoKovan // Generic contract binding to access the raw methods on
}

// LidoKovanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoKovanCallerRaw struct {
	Contract *LidoKovanCaller // Generic read-only contract binding to access the raw methods on
}

// LidoKovanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoKovanTransactorRaw struct {
	Contract *LidoKovanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidoKovan creates a new instance of LidoKovan, bound to a specific deployed contract.
func NewLidoKovan(address common.Address, backend bind.ContractBackend) (*LidoKovan, error) {
	contract, err := bindLidoKovan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LidoKovan{LidoKovanCaller: LidoKovanCaller{contract: contract}, LidoKovanTransactor: LidoKovanTransactor{contract: contract}, LidoKovanFilterer: LidoKovanFilterer{contract: contract}}, nil
}

// NewLidoKovanCaller creates a new read-only instance of LidoKovan, bound to a specific deployed contract.
func NewLidoKovanCaller(address common.Address, caller bind.ContractCaller) (*LidoKovanCaller, error) {
	contract, err := bindLidoKovan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoKovanCaller{contract: contract}, nil
}

// NewLidoKovanTransactor creates a new write-only instance of LidoKovan, bound to a specific deployed contract.
func NewLidoKovanTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoKovanTransactor, error) {
	contract, err := bindLidoKovan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoKovanTransactor{contract: contract}, nil
}

// NewLidoKovanFilterer creates a new log filterer instance of LidoKovan, bound to a specific deployed contract.
func NewLidoKovanFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoKovanFilterer, error) {
	contract, err := bindLidoKovan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoKovanFilterer{contract: contract}, nil
}

// bindLidoKovan binds a generic wrapper to an already deployed contract.
func bindLidoKovan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LidoKovanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoKovan *LidoKovanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoKovan.Contract.LidoKovanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoKovan *LidoKovanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoKovan.Contract.LidoKovanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoKovan *LidoKovanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoKovan.Contract.LidoKovanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoKovan *LidoKovanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoKovan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoKovan *LidoKovanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoKovan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoKovan *LidoKovanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoKovan.Contract.contract.Transact(opts, method, params...)
}

// GetLastCompletedReportDelta is a free data retrieval call binding the contract method 0x534649c4.
//
// Solidity: function getLastCompletedReportDelta() view returns(uint256 postTotalPooledEther, uint256 preTotalPooledEther, uint256 timeElapsed)
func (_LidoKovan *LidoKovanCaller) GetLastCompletedReportDelta(opts *bind.CallOpts) (struct {
	PostTotalPooledEther *big.Int
	PreTotalPooledEther  *big.Int
	TimeElapsed          *big.Int
}, error) {
	var out []interface{}
	err := _LidoKovan.contract.Call(opts, &out, "getLastCompletedReportDelta")

	outstruct := new(struct {
		PostTotalPooledEther *big.Int
		PreTotalPooledEther  *big.Int
		TimeElapsed          *big.Int
	})

	outstruct.PostTotalPooledEther = out[0].(*big.Int)
	outstruct.PreTotalPooledEther = out[1].(*big.Int)
	outstruct.TimeElapsed = out[2].(*big.Int)

	return *outstruct, err

}

// GetLastCompletedReportDelta is a free data retrieval call binding the contract method 0x534649c4.
//
// Solidity: function getLastCompletedReportDelta() view returns(uint256 postTotalPooledEther, uint256 preTotalPooledEther, uint256 timeElapsed)
func (_LidoKovan *LidoKovanSession) GetLastCompletedReportDelta() (struct {
	PostTotalPooledEther *big.Int
	PreTotalPooledEther  *big.Int
	TimeElapsed          *big.Int
}, error) {
	return _LidoKovan.Contract.GetLastCompletedReportDelta(&_LidoKovan.CallOpts)
}

// GetLastCompletedReportDelta is a free data retrieval call binding the contract method 0x534649c4.
//
// Solidity: function getLastCompletedReportDelta() view returns(uint256 postTotalPooledEther, uint256 preTotalPooledEther, uint256 timeElapsed)
func (_LidoKovan *LidoKovanCallerSession) GetLastCompletedReportDelta() (struct {
	PostTotalPooledEther *big.Int
	PreTotalPooledEther  *big.Int
	TimeElapsed          *big.Int
}, error) {
	return _LidoKovan.Contract.GetLastCompletedReportDelta(&_LidoKovan.CallOpts)
}

// PostTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x05225966.
//
// Solidity: function postTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCaller) PostTotalPooledEtherSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoKovan.contract.Call(opts, &out, "postTotalPooledEtherSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x05225966.
//
// Solidity: function postTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanSession) PostTotalPooledEtherSynced() (*big.Int, error) {
	return _LidoKovan.Contract.PostTotalPooledEtherSynced(&_LidoKovan.CallOpts)
}

// PostTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x05225966.
//
// Solidity: function postTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCallerSession) PostTotalPooledEtherSynced() (*big.Int, error) {
	return _LidoKovan.Contract.PostTotalPooledEtherSynced(&_LidoKovan.CallOpts)
}

// PreTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x30833b57.
//
// Solidity: function preTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCaller) PreTotalPooledEtherSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoKovan.contract.Call(opts, &out, "preTotalPooledEtherSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x30833b57.
//
// Solidity: function preTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanSession) PreTotalPooledEtherSynced() (*big.Int, error) {
	return _LidoKovan.Contract.PreTotalPooledEtherSynced(&_LidoKovan.CallOpts)
}

// PreTotalPooledEtherSynced is a free data retrieval call binding the contract method 0x30833b57.
//
// Solidity: function preTotalPooledEtherSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCallerSession) PreTotalPooledEtherSynced() (*big.Int, error) {
	return _LidoKovan.Contract.PreTotalPooledEtherSynced(&_LidoKovan.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoKovan *LidoKovanCaller) Syncer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoKovan.contract.Call(opts, &out, "syncer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoKovan *LidoKovanSession) Syncer() (common.Address, error) {
	return _LidoKovan.Contract.Syncer(&_LidoKovan.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoKovan *LidoKovanCallerSession) Syncer() (common.Address, error) {
	return _LidoKovan.Contract.Syncer(&_LidoKovan.CallOpts)
}

// TimeElapsedSynced is a free data retrieval call binding the contract method 0x60ed5e10.
//
// Solidity: function timeElapsedSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCaller) TimeElapsedSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoKovan.contract.Call(opts, &out, "timeElapsedSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeElapsedSynced is a free data retrieval call binding the contract method 0x60ed5e10.
//
// Solidity: function timeElapsedSynced() view returns(uint256)
func (_LidoKovan *LidoKovanSession) TimeElapsedSynced() (*big.Int, error) {
	return _LidoKovan.Contract.TimeElapsedSynced(&_LidoKovan.CallOpts)
}

// TimeElapsedSynced is a free data retrieval call binding the contract method 0x60ed5e10.
//
// Solidity: function timeElapsedSynced() view returns(uint256)
func (_LidoKovan *LidoKovanCallerSession) TimeElapsedSynced() (*big.Int, error) {
	return _LidoKovan.Contract.TimeElapsedSynced(&_LidoKovan.CallOpts)
}

// SyncOracle is a paid mutator transaction binding the contract method 0x5d0c0ba7.
//
// Solidity: function syncOracle(uint256 _postTotalPooledEther, uint256 _preTotalPooledEther, uint256 _timeElapsed) returns()
func (_LidoKovan *LidoKovanTransactor) SyncOracle(opts *bind.TransactOpts, _postTotalPooledEther *big.Int, _preTotalPooledEther *big.Int, _timeElapsed *big.Int) (*types.Transaction, error) {
	return _LidoKovan.contract.Transact(opts, "syncOracle", _postTotalPooledEther, _preTotalPooledEther, _timeElapsed)
}

// SyncOracle is a paid mutator transaction binding the contract method 0x5d0c0ba7.
//
// Solidity: function syncOracle(uint256 _postTotalPooledEther, uint256 _preTotalPooledEther, uint256 _timeElapsed) returns()
func (_LidoKovan *LidoKovanSession) SyncOracle(_postTotalPooledEther *big.Int, _preTotalPooledEther *big.Int, _timeElapsed *big.Int) (*types.Transaction, error) {
	return _LidoKovan.Contract.SyncOracle(&_LidoKovan.TransactOpts, _postTotalPooledEther, _preTotalPooledEther, _timeElapsed)
}

// SyncOracle is a paid mutator transaction binding the contract method 0x5d0c0ba7.
//
// Solidity: function syncOracle(uint256 _postTotalPooledEther, uint256 _preTotalPooledEther, uint256 _timeElapsed) returns()
func (_LidoKovan *LidoKovanTransactorSession) SyncOracle(_postTotalPooledEther *big.Int, _preTotalPooledEther *big.Int, _timeElapsed *big.Int) (*types.Transaction, error) {
	return _LidoKovan.Contract.SyncOracle(&_LidoKovan.TransactOpts, _postTotalPooledEther, _preTotalPooledEther, _timeElapsed)
}
