// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inchFarmingPool

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FarmAccountingInfo is an auto generated low-level Go binding around an user-defined struct.
type FarmAccountingInfo struct {
	Finished *big.Int
	Duration uint32
	Reward   *big.Int
	Balance  *big.Int
}

// InchFarmingPoolMetaData contains all meta data concerning the InchFarmingPool contract.
var InchFarmingPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"stakingToken_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBalanceExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDistributor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameStakingAndRewardsTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDistributorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardsTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStakingTokenAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"DistributorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"finished\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint184\",\"name\":\"reward\",\"type\":\"uint184\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structFarmAccounting.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"farmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor_\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"startFarming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopFarming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InchFarmingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use InchFarmingPoolMetaData.ABI instead.
var InchFarmingPoolABI = InchFarmingPoolMetaData.ABI

// InchFarmingPool is an auto generated Go binding around an Ethereum contract.
type InchFarmingPool struct {
	InchFarmingPoolCaller     // Read-only binding to the contract
	InchFarmingPoolTransactor // Write-only binding to the contract
	InchFarmingPoolFilterer   // Log filterer for contract events
}

// InchFarmingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type InchFarmingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InchFarmingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InchFarmingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InchFarmingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InchFarmingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InchFarmingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InchFarmingPoolSession struct {
	Contract     *InchFarmingPool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InchFarmingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InchFarmingPoolCallerSession struct {
	Contract *InchFarmingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InchFarmingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InchFarmingPoolTransactorSession struct {
	Contract     *InchFarmingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InchFarmingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type InchFarmingPoolRaw struct {
	Contract *InchFarmingPool // Generic contract binding to access the raw methods on
}

// InchFarmingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InchFarmingPoolCallerRaw struct {
	Contract *InchFarmingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// InchFarmingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InchFarmingPoolTransactorRaw struct {
	Contract *InchFarmingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInchFarmingPool creates a new instance of InchFarmingPool, bound to a specific deployed contract.
func NewInchFarmingPool(address common.Address, backend bind.ContractBackend) (*InchFarmingPool, error) {
	contract, err := bindInchFarmingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPool{InchFarmingPoolCaller: InchFarmingPoolCaller{contract: contract}, InchFarmingPoolTransactor: InchFarmingPoolTransactor{contract: contract}, InchFarmingPoolFilterer: InchFarmingPoolFilterer{contract: contract}}, nil
}

// NewInchFarmingPoolCaller creates a new read-only instance of InchFarmingPool, bound to a specific deployed contract.
func NewInchFarmingPoolCaller(address common.Address, caller bind.ContractCaller) (*InchFarmingPoolCaller, error) {
	contract, err := bindInchFarmingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolCaller{contract: contract}, nil
}

// NewInchFarmingPoolTransactor creates a new write-only instance of InchFarmingPool, bound to a specific deployed contract.
func NewInchFarmingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*InchFarmingPoolTransactor, error) {
	contract, err := bindInchFarmingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolTransactor{contract: contract}, nil
}

// NewInchFarmingPoolFilterer creates a new log filterer instance of InchFarmingPool, bound to a specific deployed contract.
func NewInchFarmingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*InchFarmingPoolFilterer, error) {
	contract, err := bindInchFarmingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolFilterer{contract: contract}, nil
}

// bindInchFarmingPool binds a generic wrapper to an already deployed contract.
func bindInchFarmingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InchFarmingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InchFarmingPool *InchFarmingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InchFarmingPool.Contract.InchFarmingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InchFarmingPool *InchFarmingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.InchFarmingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InchFarmingPool *InchFarmingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.InchFarmingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InchFarmingPool *InchFarmingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InchFarmingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InchFarmingPool *InchFarmingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InchFarmingPool *InchFarmingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.Allowance(&_InchFarmingPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.Allowance(&_InchFarmingPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.BalanceOf(&_InchFarmingPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.BalanceOf(&_InchFarmingPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InchFarmingPool *InchFarmingPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InchFarmingPool *InchFarmingPoolSession) Decimals() (uint8, error) {
	return _InchFarmingPool.Contract.Decimals(&_InchFarmingPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Decimals() (uint8, error) {
	return _InchFarmingPool.Contract.Decimals(&_InchFarmingPool.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_InchFarmingPool *InchFarmingPoolSession) Distributor() (common.Address, error) {
	return _InchFarmingPool.Contract.Distributor(&_InchFarmingPool.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Distributor() (common.Address, error) {
	return _InchFarmingPool.Contract.Distributor(&_InchFarmingPool.CallOpts)
}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_InchFarmingPool *InchFarmingPoolCaller) FarmInfo(opts *bind.CallOpts) (FarmAccountingInfo, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "farmInfo")

	if err != nil {
		return *new(FarmAccountingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmAccountingInfo)).(*FarmAccountingInfo)

	return out0, err

}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_InchFarmingPool *InchFarmingPoolSession) FarmInfo() (FarmAccountingInfo, error) {
	return _InchFarmingPool.Contract.FarmInfo(&_InchFarmingPool.CallOpts)
}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_InchFarmingPool *InchFarmingPoolCallerSession) FarmInfo() (FarmAccountingInfo, error) {
	return _InchFarmingPool.Contract.FarmInfo(&_InchFarmingPool.CallOpts)
}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCaller) Farmed(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "farmed", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolSession) Farmed(account common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.Farmed(&_InchFarmingPool.CallOpts, account)
}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Farmed(account common.Address) (*big.Int, error) {
	return _InchFarmingPool.Contract.Farmed(&_InchFarmingPool.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InchFarmingPool *InchFarmingPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InchFarmingPool *InchFarmingPoolSession) Name() (string, error) {
	return _InchFarmingPool.Contract.Name(&_InchFarmingPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Name() (string, error) {
	return _InchFarmingPool.Contract.Name(&_InchFarmingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InchFarmingPool *InchFarmingPoolSession) Owner() (common.Address, error) {
	return _InchFarmingPool.Contract.Owner(&_InchFarmingPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Owner() (common.Address, error) {
	return _InchFarmingPool.Contract.Owner(&_InchFarmingPool.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolSession) RewardsToken() (common.Address, error) {
	return _InchFarmingPool.Contract.RewardsToken(&_InchFarmingPool.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCallerSession) RewardsToken() (common.Address, error) {
	return _InchFarmingPool.Contract.RewardsToken(&_InchFarmingPool.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolSession) StakingToken() (common.Address, error) {
	return _InchFarmingPool.Contract.StakingToken(&_InchFarmingPool.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InchFarmingPool *InchFarmingPoolCallerSession) StakingToken() (common.Address, error) {
	return _InchFarmingPool.Contract.StakingToken(&_InchFarmingPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InchFarmingPool *InchFarmingPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InchFarmingPool *InchFarmingPoolSession) Symbol() (string, error) {
	return _InchFarmingPool.Contract.Symbol(&_InchFarmingPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InchFarmingPool *InchFarmingPoolCallerSession) Symbol() (string, error) {
	return _InchFarmingPool.Contract.Symbol(&_InchFarmingPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InchFarmingPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolSession) TotalSupply() (*big.Int, error) {
	return _InchFarmingPool.Contract.TotalSupply(&_InchFarmingPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InchFarmingPool *InchFarmingPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _InchFarmingPool.Contract.TotalSupply(&_InchFarmingPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Approve(&_InchFarmingPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Approve(&_InchFarmingPool.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InchFarmingPool *InchFarmingPoolSession) Claim() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Claim(&_InchFarmingPool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Claim() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Claim(&_InchFarmingPool.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.DecreaseAllowance(&_InchFarmingPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.DecreaseAllowance(&_InchFarmingPool.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Deposit(&_InchFarmingPool.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Deposit(&_InchFarmingPool.TransactOpts, amount)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InchFarmingPool *InchFarmingPoolSession) Exit() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Exit(&_InchFarmingPool.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Exit() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Exit(&_InchFarmingPool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.IncreaseAllowance(&_InchFarmingPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.IncreaseAllowance(&_InchFarmingPool.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InchFarmingPool *InchFarmingPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.RenounceOwnership(&_InchFarmingPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.RenounceOwnership(&_InchFarmingPool.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.RescueFunds(&_InchFarmingPool.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.RescueFunds(&_InchFarmingPool.TransactOpts, token, amount)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) SetDistributor(opts *bind.TransactOpts, distributor_ common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "setDistributor", distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_InchFarmingPool *InchFarmingPoolSession) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.SetDistributor(&_InchFarmingPool.TransactOpts, distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.SetDistributor(&_InchFarmingPool.TransactOpts, distributor_)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) StartFarming(opts *bind.TransactOpts, amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "startFarming", amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_InchFarmingPool *InchFarmingPoolSession) StartFarming(amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.StartFarming(&_InchFarmingPool.TransactOpts, amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) StartFarming(amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.StartFarming(&_InchFarmingPool.TransactOpts, amount, period)
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) StopFarming(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "stopFarming")
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_InchFarmingPool *InchFarmingPoolSession) StopFarming() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.StopFarming(&_InchFarmingPool.TransactOpts)
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) StopFarming() (*types.Transaction, error) {
	return _InchFarmingPool.Contract.StopFarming(&_InchFarmingPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Transfer(&_InchFarmingPool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Transfer(&_InchFarmingPool.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.TransferFrom(&_InchFarmingPool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_InchFarmingPool *InchFarmingPoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.TransferFrom(&_InchFarmingPool.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InchFarmingPool *InchFarmingPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.TransferOwnership(&_InchFarmingPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.TransferOwnership(&_InchFarmingPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Withdraw(&_InchFarmingPool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InchFarmingPool *InchFarmingPoolTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InchFarmingPool.Contract.Withdraw(&_InchFarmingPool.TransactOpts, amount)
}

// InchFarmingPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the InchFarmingPool contract.
type InchFarmingPoolApprovalIterator struct {
	Event *InchFarmingPoolApproval // Event containing the contract specifics and raw log

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
func (it *InchFarmingPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InchFarmingPoolApproval)
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
		it.Event = new(InchFarmingPoolApproval)
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
func (it *InchFarmingPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InchFarmingPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InchFarmingPoolApproval represents a Approval event raised by the InchFarmingPool contract.
type InchFarmingPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*InchFarmingPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InchFarmingPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolApprovalIterator{contract: _InchFarmingPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *InchFarmingPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InchFarmingPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InchFarmingPoolApproval)
				if err := _InchFarmingPool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) ParseApproval(log types.Log) (*InchFarmingPoolApproval, error) {
	event := new(InchFarmingPoolApproval)
	if err := _InchFarmingPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InchFarmingPoolDistributorChangedIterator is returned from FilterDistributorChanged and is used to iterate over the raw logs and unpacked data for DistributorChanged events raised by the InchFarmingPool contract.
type InchFarmingPoolDistributorChangedIterator struct {
	Event *InchFarmingPoolDistributorChanged // Event containing the contract specifics and raw log

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
func (it *InchFarmingPoolDistributorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InchFarmingPoolDistributorChanged)
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
		it.Event = new(InchFarmingPoolDistributorChanged)
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
func (it *InchFarmingPoolDistributorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InchFarmingPoolDistributorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InchFarmingPoolDistributorChanged represents a DistributorChanged event raised by the InchFarmingPool contract.
type InchFarmingPoolDistributorChanged struct {
	NewDistributor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributorChanged is a free log retrieval operation binding the contract event 0xe37acc13f5ed9d0cc83c2842e093fe5a494d5b8fb5b1db06356b327081832f52.
//
// Solidity: event DistributorChanged(address newDistributor)
func (_InchFarmingPool *InchFarmingPoolFilterer) FilterDistributorChanged(opts *bind.FilterOpts) (*InchFarmingPoolDistributorChangedIterator, error) {

	logs, sub, err := _InchFarmingPool.contract.FilterLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolDistributorChangedIterator{contract: _InchFarmingPool.contract, event: "DistributorChanged", logs: logs, sub: sub}, nil
}

// WatchDistributorChanged is a free log subscription operation binding the contract event 0xe37acc13f5ed9d0cc83c2842e093fe5a494d5b8fb5b1db06356b327081832f52.
//
// Solidity: event DistributorChanged(address newDistributor)
func (_InchFarmingPool *InchFarmingPoolFilterer) WatchDistributorChanged(opts *bind.WatchOpts, sink chan<- *InchFarmingPoolDistributorChanged) (event.Subscription, error) {

	logs, sub, err := _InchFarmingPool.contract.WatchLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InchFarmingPoolDistributorChanged)
				if err := _InchFarmingPool.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
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

// ParseDistributorChanged is a log parse operation binding the contract event 0xe37acc13f5ed9d0cc83c2842e093fe5a494d5b8fb5b1db06356b327081832f52.
//
// Solidity: event DistributorChanged(address newDistributor)
func (_InchFarmingPool *InchFarmingPoolFilterer) ParseDistributorChanged(log types.Log) (*InchFarmingPoolDistributorChanged, error) {
	event := new(InchFarmingPoolDistributorChanged)
	if err := _InchFarmingPool.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InchFarmingPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InchFarmingPool contract.
type InchFarmingPoolOwnershipTransferredIterator struct {
	Event *InchFarmingPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InchFarmingPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InchFarmingPoolOwnershipTransferred)
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
		it.Event = new(InchFarmingPoolOwnershipTransferred)
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
func (it *InchFarmingPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InchFarmingPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InchFarmingPoolOwnershipTransferred represents a OwnershipTransferred event raised by the InchFarmingPool contract.
type InchFarmingPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InchFarmingPool *InchFarmingPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InchFarmingPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InchFarmingPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolOwnershipTransferredIterator{contract: _InchFarmingPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InchFarmingPool *InchFarmingPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InchFarmingPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InchFarmingPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InchFarmingPoolOwnershipTransferred)
				if err := _InchFarmingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_InchFarmingPool *InchFarmingPoolFilterer) ParseOwnershipTransferred(log types.Log) (*InchFarmingPoolOwnershipTransferred, error) {
	event := new(InchFarmingPoolOwnershipTransferred)
	if err := _InchFarmingPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InchFarmingPoolRewardUpdatedIterator is returned from FilterRewardUpdated and is used to iterate over the raw logs and unpacked data for RewardUpdated events raised by the InchFarmingPool contract.
type InchFarmingPoolRewardUpdatedIterator struct {
	Event *InchFarmingPoolRewardUpdated // Event containing the contract specifics and raw log

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
func (it *InchFarmingPoolRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InchFarmingPoolRewardUpdated)
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
		it.Event = new(InchFarmingPoolRewardUpdated)
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
func (it *InchFarmingPoolRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InchFarmingPoolRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InchFarmingPoolRewardUpdated represents a RewardUpdated event raised by the InchFarmingPool contract.
type InchFarmingPoolRewardUpdated struct {
	Reward   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdated is a free log retrieval operation binding the contract event 0x3efe2b1ade87153c913a322f09a35c930d7fd699770b8d7cdd294e1debf6e9e4.
//
// Solidity: event RewardUpdated(uint256 reward, uint256 duration)
func (_InchFarmingPool *InchFarmingPoolFilterer) FilterRewardUpdated(opts *bind.FilterOpts) (*InchFarmingPoolRewardUpdatedIterator, error) {

	logs, sub, err := _InchFarmingPool.contract.FilterLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolRewardUpdatedIterator{contract: _InchFarmingPool.contract, event: "RewardUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardUpdated is a free log subscription operation binding the contract event 0x3efe2b1ade87153c913a322f09a35c930d7fd699770b8d7cdd294e1debf6e9e4.
//
// Solidity: event RewardUpdated(uint256 reward, uint256 duration)
func (_InchFarmingPool *InchFarmingPoolFilterer) WatchRewardUpdated(opts *bind.WatchOpts, sink chan<- *InchFarmingPoolRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _InchFarmingPool.contract.WatchLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InchFarmingPoolRewardUpdated)
				if err := _InchFarmingPool.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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

// ParseRewardUpdated is a log parse operation binding the contract event 0x3efe2b1ade87153c913a322f09a35c930d7fd699770b8d7cdd294e1debf6e9e4.
//
// Solidity: event RewardUpdated(uint256 reward, uint256 duration)
func (_InchFarmingPool *InchFarmingPoolFilterer) ParseRewardUpdated(log types.Log) (*InchFarmingPoolRewardUpdated, error) {
	event := new(InchFarmingPoolRewardUpdated)
	if err := _InchFarmingPool.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InchFarmingPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the InchFarmingPool contract.
type InchFarmingPoolTransferIterator struct {
	Event *InchFarmingPoolTransfer // Event containing the contract specifics and raw log

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
func (it *InchFarmingPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InchFarmingPoolTransfer)
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
		it.Event = new(InchFarmingPoolTransfer)
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
func (it *InchFarmingPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InchFarmingPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InchFarmingPoolTransfer represents a Transfer event raised by the InchFarmingPool contract.
type InchFarmingPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*InchFarmingPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InchFarmingPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InchFarmingPoolTransferIterator{contract: _InchFarmingPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *InchFarmingPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InchFarmingPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InchFarmingPoolTransfer)
				if err := _InchFarmingPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InchFarmingPool *InchFarmingPoolFilterer) ParseTransfer(log types.Log) (*InchFarmingPoolTransfer, error) {
	event := new(InchFarmingPoolTransfer)
	if err := _InchFarmingPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
