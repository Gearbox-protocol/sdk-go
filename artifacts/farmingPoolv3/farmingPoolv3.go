// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package farmingPoolv3

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

// FarmingPoolv3MetaData contains all meta data concerning the FarmingPoolv3 contract.
var FarmingPoolv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"stakingToken_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBalanceExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameDistributor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameStakingAndRewardsTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDistributorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardsTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStakingTokenAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDistributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"DistributorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"RewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"finished\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"duration\",\"type\":\"uint32\"},{\"internalType\":\"uint184\",\"name\":\"reward\",\"type\":\"uint184\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structFarmAccounting.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"farmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"distributor_\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"startFarming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopFarming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FarmingPoolv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmingPoolv3MetaData.ABI instead.
var FarmingPoolv3ABI = FarmingPoolv3MetaData.ABI

// FarmingPoolv3 is an auto generated Go binding around an Ethereum contract.
type FarmingPoolv3 struct {
	FarmingPoolv3Caller     // Read-only binding to the contract
	FarmingPoolv3Transactor // Write-only binding to the contract
	FarmingPoolv3Filterer   // Log filterer for contract events
}

// FarmingPoolv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type FarmingPoolv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingPoolv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmingPoolv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingPoolv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmingPoolv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingPoolv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmingPoolv3Session struct {
	Contract     *FarmingPoolv3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmingPoolv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmingPoolv3CallerSession struct {
	Contract *FarmingPoolv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FarmingPoolv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmingPoolv3TransactorSession struct {
	Contract     *FarmingPoolv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FarmingPoolv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type FarmingPoolv3Raw struct {
	Contract *FarmingPoolv3 // Generic contract binding to access the raw methods on
}

// FarmingPoolv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmingPoolv3CallerRaw struct {
	Contract *FarmingPoolv3Caller // Generic read-only contract binding to access the raw methods on
}

// FarmingPoolv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmingPoolv3TransactorRaw struct {
	Contract *FarmingPoolv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFarmingPoolv3 creates a new instance of FarmingPoolv3, bound to a specific deployed contract.
func NewFarmingPoolv3(address common.Address, backend bind.ContractBackend) (*FarmingPoolv3, error) {
	contract, err := bindFarmingPoolv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3{FarmingPoolv3Caller: FarmingPoolv3Caller{contract: contract}, FarmingPoolv3Transactor: FarmingPoolv3Transactor{contract: contract}, FarmingPoolv3Filterer: FarmingPoolv3Filterer{contract: contract}}, nil
}

// NewFarmingPoolv3Caller creates a new read-only instance of FarmingPoolv3, bound to a specific deployed contract.
func NewFarmingPoolv3Caller(address common.Address, caller bind.ContractCaller) (*FarmingPoolv3Caller, error) {
	contract, err := bindFarmingPoolv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3Caller{contract: contract}, nil
}

// NewFarmingPoolv3Transactor creates a new write-only instance of FarmingPoolv3, bound to a specific deployed contract.
func NewFarmingPoolv3Transactor(address common.Address, transactor bind.ContractTransactor) (*FarmingPoolv3Transactor, error) {
	contract, err := bindFarmingPoolv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3Transactor{contract: contract}, nil
}

// NewFarmingPoolv3Filterer creates a new log filterer instance of FarmingPoolv3, bound to a specific deployed contract.
func NewFarmingPoolv3Filterer(address common.Address, filterer bind.ContractFilterer) (*FarmingPoolv3Filterer, error) {
	contract, err := bindFarmingPoolv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3Filterer{contract: contract}, nil
}

// bindFarmingPoolv3 binds a generic wrapper to an already deployed contract.
func bindFarmingPoolv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarmingPoolv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmingPoolv3 *FarmingPoolv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmingPoolv3.Contract.FarmingPoolv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmingPoolv3 *FarmingPoolv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.FarmingPoolv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmingPoolv3 *FarmingPoolv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.FarmingPoolv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmingPoolv3 *FarmingPoolv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmingPoolv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmingPoolv3 *FarmingPoolv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmingPoolv3 *FarmingPoolv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.Allowance(&_FarmingPoolv3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.Allowance(&_FarmingPoolv3.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.BalanceOf(&_FarmingPoolv3.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.BalanceOf(&_FarmingPoolv3.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FarmingPoolv3 *FarmingPoolv3Session) Decimals() (uint8, error) {
	return _FarmingPoolv3.Contract.Decimals(&_FarmingPoolv3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Decimals() (uint8, error) {
	return _FarmingPoolv3.Contract.Decimals(&_FarmingPoolv3.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Session) Distributor() (common.Address, error) {
	return _FarmingPoolv3.Contract.Distributor(&_FarmingPoolv3.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Distributor() (common.Address, error) {
	return _FarmingPoolv3.Contract.Distributor(&_FarmingPoolv3.CallOpts)
}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_FarmingPoolv3 *FarmingPoolv3Caller) FarmInfo(opts *bind.CallOpts) (FarmAccountingInfo, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "farmInfo")

	if err != nil {
		return *new(FarmAccountingInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmAccountingInfo)).(*FarmAccountingInfo)

	return out0, err

}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_FarmingPoolv3 *FarmingPoolv3Session) FarmInfo() (FarmAccountingInfo, error) {
	return _FarmingPoolv3.Contract.FarmInfo(&_FarmingPoolv3.CallOpts)
}

// FarmInfo is a free data retrieval call binding the contract method 0x1d49d66c.
//
// Solidity: function farmInfo() view returns((uint40,uint32,uint184,uint256))
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) FarmInfo() (FarmAccountingInfo, error) {
	return _FarmingPoolv3.Contract.FarmInfo(&_FarmingPoolv3.CallOpts)
}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Farmed(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "farmed", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Session) Farmed(account common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.Farmed(&_FarmingPoolv3.CallOpts, account)
}

// Farmed is a free data retrieval call binding the contract method 0x4216f972.
//
// Solidity: function farmed(address account) view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Farmed(account common.Address) (*big.Int, error) {
	return _FarmingPoolv3.Contract.Farmed(&_FarmingPoolv3.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3Session) Name() (string, error) {
	return _FarmingPoolv3.Contract.Name(&_FarmingPoolv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Name() (string, error) {
	return _FarmingPoolv3.Contract.Name(&_FarmingPoolv3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Session) Owner() (common.Address, error) {
	return _FarmingPoolv3.Contract.Owner(&_FarmingPoolv3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Owner() (common.Address, error) {
	return _FarmingPoolv3.Contract.Owner(&_FarmingPoolv3.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Caller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Session) RewardsToken() (common.Address, error) {
	return _FarmingPoolv3.Contract.RewardsToken(&_FarmingPoolv3.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) RewardsToken() (common.Address, error) {
	return _FarmingPoolv3.Contract.RewardsToken(&_FarmingPoolv3.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Caller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3Session) StakingToken() (common.Address, error) {
	return _FarmingPoolv3.Contract.StakingToken(&_FarmingPoolv3.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) StakingToken() (common.Address, error) {
	return _FarmingPoolv3.Contract.StakingToken(&_FarmingPoolv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3Session) Symbol() (string, error) {
	return _FarmingPoolv3.Contract.Symbol(&_FarmingPoolv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) Symbol() (string, error) {
	return _FarmingPoolv3.Contract.Symbol(&_FarmingPoolv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FarmingPoolv3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3Session) TotalSupply() (*big.Int, error) {
	return _FarmingPoolv3.Contract.TotalSupply(&_FarmingPoolv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FarmingPoolv3 *FarmingPoolv3CallerSession) TotalSupply() (*big.Int, error) {
	return _FarmingPoolv3.Contract.TotalSupply(&_FarmingPoolv3.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Approve(&_FarmingPoolv3.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Approve(&_FarmingPoolv3.TransactOpts, spender, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) Claim() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Claim(&_FarmingPoolv3.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Claim() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Claim(&_FarmingPoolv3.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.DecreaseAllowance(&_FarmingPoolv3.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.DecreaseAllowance(&_FarmingPoolv3.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Deposit(&_FarmingPoolv3.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Deposit(&_FarmingPoolv3.TransactOpts, amount)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) Exit() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Exit(&_FarmingPoolv3.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Exit() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Exit(&_FarmingPoolv3.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.IncreaseAllowance(&_FarmingPoolv3.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.IncreaseAllowance(&_FarmingPoolv3.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) RenounceOwnership() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.RenounceOwnership(&_FarmingPoolv3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.RenounceOwnership(&_FarmingPoolv3.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.RescueFunds(&_FarmingPoolv3.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.RescueFunds(&_FarmingPoolv3.TransactOpts, token, amount)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) SetDistributor(opts *bind.TransactOpts, distributor_ common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "setDistributor", distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.SetDistributor(&_FarmingPoolv3.TransactOpts, distributor_)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address distributor_) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) SetDistributor(distributor_ common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.SetDistributor(&_FarmingPoolv3.TransactOpts, distributor_)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) StartFarming(opts *bind.TransactOpts, amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "startFarming", amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) StartFarming(amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.StartFarming(&_FarmingPoolv3.TransactOpts, amount, period)
}

// StartFarming is a paid mutator transaction binding the contract method 0x70261222.
//
// Solidity: function startFarming(uint256 amount, uint256 period) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) StartFarming(amount *big.Int, period *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.StartFarming(&_FarmingPoolv3.TransactOpts, amount, period)
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) StopFarming(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "stopFarming")
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) StopFarming() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.StopFarming(&_FarmingPoolv3.TransactOpts)
}

// StopFarming is a paid mutator transaction binding the contract method 0x1bfa4c04.
//
// Solidity: function stopFarming() returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) StopFarming() (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.StopFarming(&_FarmingPoolv3.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Transfer(&_FarmingPoolv3.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Transfer(&_FarmingPoolv3.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.TransferFrom(&_FarmingPoolv3.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.TransferFrom(&_FarmingPoolv3.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.TransferOwnership(&_FarmingPoolv3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.TransferOwnership(&_FarmingPoolv3.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Withdraw(&_FarmingPoolv3.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_FarmingPoolv3 *FarmingPoolv3TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _FarmingPoolv3.Contract.Withdraw(&_FarmingPoolv3.TransactOpts, amount)
}

// FarmingPoolv3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FarmingPoolv3 contract.
type FarmingPoolv3ApprovalIterator struct {
	Event *FarmingPoolv3Approval // Event containing the contract specifics and raw log

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
func (it *FarmingPoolv3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPoolv3Approval)
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
		it.Event = new(FarmingPoolv3Approval)
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
func (it *FarmingPoolv3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPoolv3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPoolv3Approval represents a Approval event raised by the FarmingPoolv3 contract.
type FarmingPoolv3Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FarmingPoolv3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3ApprovalIterator{contract: _FarmingPoolv3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FarmingPoolv3Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPoolv3Approval)
				if err := _FarmingPoolv3.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FarmingPoolv3 *FarmingPoolv3Filterer) ParseApproval(log types.Log) (*FarmingPoolv3Approval, error) {
	event := new(FarmingPoolv3Approval)
	if err := _FarmingPoolv3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPoolv3DistributorChangedIterator is returned from FilterDistributorChanged and is used to iterate over the raw logs and unpacked data for DistributorChanged events raised by the FarmingPoolv3 contract.
type FarmingPoolv3DistributorChangedIterator struct {
	Event *FarmingPoolv3DistributorChanged // Event containing the contract specifics and raw log

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
func (it *FarmingPoolv3DistributorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPoolv3DistributorChanged)
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
		it.Event = new(FarmingPoolv3DistributorChanged)
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
func (it *FarmingPoolv3DistributorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPoolv3DistributorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPoolv3DistributorChanged represents a DistributorChanged event raised by the FarmingPoolv3 contract.
type FarmingPoolv3DistributorChanged struct {
	OldDistributor common.Address
	NewDistributor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDistributorChanged is a free log retrieval operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) FilterDistributorChanged(opts *bind.FilterOpts) (*FarmingPoolv3DistributorChangedIterator, error) {

	logs, sub, err := _FarmingPoolv3.contract.FilterLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3DistributorChangedIterator{contract: _FarmingPoolv3.contract, event: "DistributorChanged", logs: logs, sub: sub}, nil
}

// WatchDistributorChanged is a free log subscription operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) WatchDistributorChanged(opts *bind.WatchOpts, sink chan<- *FarmingPoolv3DistributorChanged) (event.Subscription, error) {

	logs, sub, err := _FarmingPoolv3.contract.WatchLogs(opts, "DistributorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPoolv3DistributorChanged)
				if err := _FarmingPoolv3.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
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

// ParseDistributorChanged is a log parse operation binding the contract event 0xa9f739537fc57540bed0a44e33e27baa63290d865cc15f0f16cf17d38c998a4d.
//
// Solidity: event DistributorChanged(address oldDistributor, address newDistributor)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) ParseDistributorChanged(log types.Log) (*FarmingPoolv3DistributorChanged, error) {
	event := new(FarmingPoolv3DistributorChanged)
	if err := _FarmingPoolv3.contract.UnpackLog(event, "DistributorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPoolv3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FarmingPoolv3 contract.
type FarmingPoolv3OwnershipTransferredIterator struct {
	Event *FarmingPoolv3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarmingPoolv3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPoolv3OwnershipTransferred)
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
		it.Event = new(FarmingPoolv3OwnershipTransferred)
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
func (it *FarmingPoolv3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPoolv3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPoolv3OwnershipTransferred represents a OwnershipTransferred event raised by the FarmingPoolv3 contract.
type FarmingPoolv3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarmingPoolv3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3OwnershipTransferredIterator{contract: _FarmingPoolv3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarmingPoolv3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPoolv3OwnershipTransferred)
				if err := _FarmingPoolv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FarmingPoolv3 *FarmingPoolv3Filterer) ParseOwnershipTransferred(log types.Log) (*FarmingPoolv3OwnershipTransferred, error) {
	event := new(FarmingPoolv3OwnershipTransferred)
	if err := _FarmingPoolv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPoolv3RewardUpdatedIterator is returned from FilterRewardUpdated and is used to iterate over the raw logs and unpacked data for RewardUpdated events raised by the FarmingPoolv3 contract.
type FarmingPoolv3RewardUpdatedIterator struct {
	Event *FarmingPoolv3RewardUpdated // Event containing the contract specifics and raw log

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
func (it *FarmingPoolv3RewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPoolv3RewardUpdated)
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
		it.Event = new(FarmingPoolv3RewardUpdated)
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
func (it *FarmingPoolv3RewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPoolv3RewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPoolv3RewardUpdated represents a RewardUpdated event raised by the FarmingPoolv3 contract.
type FarmingPoolv3RewardUpdated struct {
	Reward   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardUpdated is a free log retrieval operation binding the contract event 0x3efe2b1ade87153c913a322f09a35c930d7fd699770b8d7cdd294e1debf6e9e4.
//
// Solidity: event RewardUpdated(uint256 reward, uint256 duration)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) FilterRewardUpdated(opts *bind.FilterOpts) (*FarmingPoolv3RewardUpdatedIterator, error) {

	logs, sub, err := _FarmingPoolv3.contract.FilterLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3RewardUpdatedIterator{contract: _FarmingPoolv3.contract, event: "RewardUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardUpdated is a free log subscription operation binding the contract event 0x3efe2b1ade87153c913a322f09a35c930d7fd699770b8d7cdd294e1debf6e9e4.
//
// Solidity: event RewardUpdated(uint256 reward, uint256 duration)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) WatchRewardUpdated(opts *bind.WatchOpts, sink chan<- *FarmingPoolv3RewardUpdated) (event.Subscription, error) {

	logs, sub, err := _FarmingPoolv3.contract.WatchLogs(opts, "RewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPoolv3RewardUpdated)
				if err := _FarmingPoolv3.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
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
func (_FarmingPoolv3 *FarmingPoolv3Filterer) ParseRewardUpdated(log types.Log) (*FarmingPoolv3RewardUpdated, error) {
	event := new(FarmingPoolv3RewardUpdated)
	if err := _FarmingPoolv3.contract.UnpackLog(event, "RewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPoolv3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FarmingPoolv3 contract.
type FarmingPoolv3TransferIterator struct {
	Event *FarmingPoolv3Transfer // Event containing the contract specifics and raw log

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
func (it *FarmingPoolv3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPoolv3Transfer)
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
		it.Event = new(FarmingPoolv3Transfer)
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
func (it *FarmingPoolv3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPoolv3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPoolv3Transfer represents a Transfer event raised by the FarmingPoolv3 contract.
type FarmingPoolv3Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FarmingPoolv3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FarmingPoolv3TransferIterator{contract: _FarmingPoolv3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FarmingPoolv3 *FarmingPoolv3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FarmingPoolv3Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FarmingPoolv3.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPoolv3Transfer)
				if err := _FarmingPoolv3.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FarmingPoolv3 *FarmingPoolv3Filterer) ParseTransfer(log types.Log) (*FarmingPoolv3Transfer, error) {
	event := new(FarmingPoolv3Transfer)
	if err := _FarmingPoolv3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
