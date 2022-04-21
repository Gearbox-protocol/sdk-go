// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package convexManager

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

// ConvexManagerABI is the input ABI used to generate the binding from.
const ConvexManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"syncer_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crv_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curveLpToken\",\"type\":\"address\"}],\"name\":\"addBasePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_basePool\",\"type\":\"address\"}],\"name\":\"addExtraPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"booster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"syncCVXSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queuedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_historicalRewards\",\"type\":\"uint256\"}],\"name\":\"syncPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConvexManager is an auto generated Go binding around an Ethereum contract.
type ConvexManager struct {
	ConvexManagerCaller     // Read-only binding to the contract
	ConvexManagerTransactor // Write-only binding to the contract
	ConvexManagerFilterer   // Log filterer for contract events
}

// ConvexManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvexManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvexManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvexManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvexManagerSession struct {
	Contract     *ConvexManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvexManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvexManagerCallerSession struct {
	Contract *ConvexManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConvexManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvexManagerTransactorSession struct {
	Contract     *ConvexManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConvexManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvexManagerRaw struct {
	Contract *ConvexManager // Generic contract binding to access the raw methods on
}

// ConvexManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvexManagerCallerRaw struct {
	Contract *ConvexManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ConvexManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvexManagerTransactorRaw struct {
	Contract *ConvexManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvexManager creates a new instance of ConvexManager, bound to a specific deployed contract.
func NewConvexManager(address common.Address, backend bind.ContractBackend) (*ConvexManager, error) {
	contract, err := bindConvexManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvexManager{ConvexManagerCaller: ConvexManagerCaller{contract: contract}, ConvexManagerTransactor: ConvexManagerTransactor{contract: contract}, ConvexManagerFilterer: ConvexManagerFilterer{contract: contract}}, nil
}

// NewConvexManagerCaller creates a new read-only instance of ConvexManager, bound to a specific deployed contract.
func NewConvexManagerCaller(address common.Address, caller bind.ContractCaller) (*ConvexManagerCaller, error) {
	contract, err := bindConvexManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexManagerCaller{contract: contract}, nil
}

// NewConvexManagerTransactor creates a new write-only instance of ConvexManager, bound to a specific deployed contract.
func NewConvexManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvexManagerTransactor, error) {
	contract, err := bindConvexManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexManagerTransactor{contract: contract}, nil
}

// NewConvexManagerFilterer creates a new log filterer instance of ConvexManager, bound to a specific deployed contract.
func NewConvexManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvexManagerFilterer, error) {
	contract, err := bindConvexManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvexManagerFilterer{contract: contract}, nil
}

// bindConvexManager binds a generic wrapper to an already deployed contract.
func bindConvexManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvexManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexManager *ConvexManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexManager.Contract.ConvexManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexManager *ConvexManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexManager.Contract.ConvexManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexManager *ConvexManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexManager.Contract.ConvexManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexManager *ConvexManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexManager *ConvexManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexManager *ConvexManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexManager.Contract.contract.Transact(opts, method, params...)
}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_ConvexManager *ConvexManagerCaller) Booster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "booster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_ConvexManager *ConvexManagerSession) Booster() (common.Address, error) {
	return _ConvexManager.Contract.Booster(&_ConvexManager.CallOpts)
}

// Booster is a free data retrieval call binding the contract method 0xc6def076.
//
// Solidity: function booster() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) Booster() (common.Address, error) {
	return _ConvexManager.Contract.Booster(&_ConvexManager.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexManager *ConvexManagerCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexManager *ConvexManagerSession) Crv() (common.Address, error) {
	return _ConvexManager.Contract.Crv(&_ConvexManager.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) Crv() (common.Address, error) {
	return _ConvexManager.Contract.Crv(&_ConvexManager.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexManager *ConvexManagerCaller) Cvx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "cvx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexManager *ConvexManagerSession) Cvx() (common.Address, error) {
	return _ConvexManager.Contract.Cvx(&_ConvexManager.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) Cvx() (common.Address, error) {
	return _ConvexManager.Contract.Cvx(&_ConvexManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvexManager *ConvexManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvexManager *ConvexManagerSession) Owner() (common.Address, error) {
	return _ConvexManager.Contract.Owner(&_ConvexManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) Owner() (common.Address, error) {
	return _ConvexManager.Contract.Owner(&_ConvexManager.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ConvexManager *ConvexManagerCaller) PoolFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "poolFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ConvexManager *ConvexManagerSession) PoolFactory() (common.Address, error) {
	return _ConvexManager.Contract.PoolFactory(&_ConvexManager.CallOpts)
}

// PoolFactory is a free data retrieval call binding the contract method 0x4219dc40.
//
// Solidity: function poolFactory() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) PoolFactory() (common.Address, error) {
	return _ConvexManager.Contract.PoolFactory(&_ConvexManager.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ConvexManager *ConvexManagerCaller) Syncer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "syncer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ConvexManager *ConvexManagerSession) Syncer() (common.Address, error) {
	return _ConvexManager.Contract.Syncer(&_ConvexManager.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) Syncer() (common.Address, error) {
	return _ConvexManager.Contract.Syncer(&_ConvexManager.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_ConvexManager *ConvexManagerCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexManager.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_ConvexManager *ConvexManagerSession) TokenFactory() (common.Address, error) {
	return _ConvexManager.Contract.TokenFactory(&_ConvexManager.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_ConvexManager *ConvexManagerCallerSession) TokenFactory() (common.Address, error) {
	return _ConvexManager.Contract.TokenFactory(&_ConvexManager.CallOpts)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x4d115d66.
//
// Solidity: function addBasePool(address _curveLpToken) returns(address)
func (_ConvexManager *ConvexManagerTransactor) AddBasePool(opts *bind.TransactOpts, _curveLpToken common.Address) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "addBasePool", _curveLpToken)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x4d115d66.
//
// Solidity: function addBasePool(address _curveLpToken) returns(address)
func (_ConvexManager *ConvexManagerSession) AddBasePool(_curveLpToken common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.AddBasePool(&_ConvexManager.TransactOpts, _curveLpToken)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x4d115d66.
//
// Solidity: function addBasePool(address _curveLpToken) returns(address)
func (_ConvexManager *ConvexManagerTransactorSession) AddBasePool(_curveLpToken common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.AddBasePool(&_ConvexManager.TransactOpts, _curveLpToken)
}

// AddExtraPool is a paid mutator transaction binding the contract method 0x3031bc92.
//
// Solidity: function addExtraPool(address _rewardToken, address _basePool) returns(address)
func (_ConvexManager *ConvexManagerTransactor) AddExtraPool(opts *bind.TransactOpts, _rewardToken common.Address, _basePool common.Address) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "addExtraPool", _rewardToken, _basePool)
}

// AddExtraPool is a paid mutator transaction binding the contract method 0x3031bc92.
//
// Solidity: function addExtraPool(address _rewardToken, address _basePool) returns(address)
func (_ConvexManager *ConvexManagerSession) AddExtraPool(_rewardToken common.Address, _basePool common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.AddExtraPool(&_ConvexManager.TransactOpts, _rewardToken, _basePool)
}

// AddExtraPool is a paid mutator transaction binding the contract method 0x3031bc92.
//
// Solidity: function addExtraPool(address _rewardToken, address _basePool) returns(address)
func (_ConvexManager *ConvexManagerTransactorSession) AddExtraPool(_rewardToken common.Address, _basePool common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.AddExtraPool(&_ConvexManager.TransactOpts, _rewardToken, _basePool)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConvexManager *ConvexManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConvexManager *ConvexManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConvexManager.Contract.RenounceOwnership(&_ConvexManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConvexManager *ConvexManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConvexManager.Contract.RenounceOwnership(&_ConvexManager.TransactOpts)
}

// SyncCVXSupply is a paid mutator transaction binding the contract method 0xfd4038f5.
//
// Solidity: function syncCVXSupply(uint256 supply) returns()
func (_ConvexManager *ConvexManagerTransactor) SyncCVXSupply(opts *bind.TransactOpts, supply *big.Int) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "syncCVXSupply", supply)
}

// SyncCVXSupply is a paid mutator transaction binding the contract method 0xfd4038f5.
//
// Solidity: function syncCVXSupply(uint256 supply) returns()
func (_ConvexManager *ConvexManagerSession) SyncCVXSupply(supply *big.Int) (*types.Transaction, error) {
	return _ConvexManager.Contract.SyncCVXSupply(&_ConvexManager.TransactOpts, supply)
}

// SyncCVXSupply is a paid mutator transaction binding the contract method 0xfd4038f5.
//
// Solidity: function syncCVXSupply(uint256 supply) returns()
func (_ConvexManager *ConvexManagerTransactorSession) SyncCVXSupply(supply *big.Int) (*types.Transaction, error) {
	return _ConvexManager.Contract.SyncCVXSupply(&_ConvexManager.TransactOpts, supply)
}

// SyncPool is a paid mutator transaction binding the contract method 0x07c4e33a.
//
// Solidity: function syncPool(address pool, uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_ConvexManager *ConvexManagerTransactor) SyncPool(opts *bind.TransactOpts, pool common.Address, _periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "syncPool", pool, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// SyncPool is a paid mutator transaction binding the contract method 0x07c4e33a.
//
// Solidity: function syncPool(address pool, uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_ConvexManager *ConvexManagerSession) SyncPool(pool common.Address, _periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _ConvexManager.Contract.SyncPool(&_ConvexManager.TransactOpts, pool, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// SyncPool is a paid mutator transaction binding the contract method 0x07c4e33a.
//
// Solidity: function syncPool(address pool, uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_ConvexManager *ConvexManagerTransactorSession) SyncPool(pool common.Address, _periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _ConvexManager.Contract.SyncPool(&_ConvexManager.TransactOpts, pool, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConvexManager *ConvexManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConvexManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConvexManager *ConvexManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.TransferOwnership(&_ConvexManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConvexManager *ConvexManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConvexManager.Contract.TransferOwnership(&_ConvexManager.TransactOpts, newOwner)
}

// ConvexManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConvexManager contract.
type ConvexManagerOwnershipTransferredIterator struct {
	Event *ConvexManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvexManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvexManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvexManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ConvexManager contract.
type ConvexManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConvexManager *ConvexManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConvexManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConvexManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConvexManagerOwnershipTransferredIterator{contract: _ConvexManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConvexManager *ConvexManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConvexManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConvexManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexManagerOwnershipTransferred)
				if err := _ConvexManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConvexManager *ConvexManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ConvexManagerOwnershipTransferred, error) {
	event := new(ConvexManagerOwnershipTransferred)
	if err := _ConvexManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
