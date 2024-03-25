// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev2LendingPool

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

// Aavev2LendingPoolMetaData contains all meta data concerning the Aavev2LendingPool contract.
var Aavev2LendingPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lendingPool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositDiff\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDiff\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Aavev2LendingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use Aavev2LendingPoolMetaData.ABI instead.
var Aavev2LendingPoolABI = Aavev2LendingPoolMetaData.ABI

// Aavev2LendingPool is an auto generated Go binding around an Ethereum contract.
type Aavev2LendingPool struct {
	Aavev2LendingPoolCaller     // Read-only binding to the contract
	Aavev2LendingPoolTransactor // Write-only binding to the contract
	Aavev2LendingPoolFilterer   // Log filterer for contract events
}

// Aavev2LendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type Aavev2LendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2LendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Aavev2LendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2LendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Aavev2LendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2LendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Aavev2LendingPoolSession struct {
	Contract     *Aavev2LendingPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Aavev2LendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Aavev2LendingPoolCallerSession struct {
	Contract *Aavev2LendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Aavev2LendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Aavev2LendingPoolTransactorSession struct {
	Contract     *Aavev2LendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Aavev2LendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type Aavev2LendingPoolRaw struct {
	Contract *Aavev2LendingPool // Generic contract binding to access the raw methods on
}

// Aavev2LendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Aavev2LendingPoolCallerRaw struct {
	Contract *Aavev2LendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// Aavev2LendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Aavev2LendingPoolTransactorRaw struct {
	Contract *Aavev2LendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavev2LendingPool creates a new instance of Aavev2LendingPool, bound to a specific deployed contract.
func NewAavev2LendingPool(address common.Address, backend bind.ContractBackend) (*Aavev2LendingPool, error) {
	contract, err := bindAavev2LendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aavev2LendingPool{Aavev2LendingPoolCaller: Aavev2LendingPoolCaller{contract: contract}, Aavev2LendingPoolTransactor: Aavev2LendingPoolTransactor{contract: contract}, Aavev2LendingPoolFilterer: Aavev2LendingPoolFilterer{contract: contract}}, nil
}

// NewAavev2LendingPoolCaller creates a new read-only instance of Aavev2LendingPool, bound to a specific deployed contract.
func NewAavev2LendingPoolCaller(address common.Address, caller bind.ContractCaller) (*Aavev2LendingPoolCaller, error) {
	contract, err := bindAavev2LendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Aavev2LendingPoolCaller{contract: contract}, nil
}

// NewAavev2LendingPoolTransactor creates a new write-only instance of Aavev2LendingPool, bound to a specific deployed contract.
func NewAavev2LendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*Aavev2LendingPoolTransactor, error) {
	contract, err := bindAavev2LendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Aavev2LendingPoolTransactor{contract: contract}, nil
}

// NewAavev2LendingPoolFilterer creates a new log filterer instance of Aavev2LendingPool, bound to a specific deployed contract.
func NewAavev2LendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*Aavev2LendingPoolFilterer, error) {
	contract, err := bindAavev2LendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Aavev2LendingPoolFilterer{contract: contract}, nil
}

// bindAavev2LendingPool binds a generic wrapper to an already deployed contract.
func bindAavev2LendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Aavev2LendingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aavev2LendingPool *Aavev2LendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aavev2LendingPool.Contract.Aavev2LendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aavev2LendingPool *Aavev2LendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Aavev2LendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aavev2LendingPool *Aavev2LendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Aavev2LendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aavev2LendingPool *Aavev2LendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aavev2LendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) GearboxAdapterType() (uint8, error) {
	return _Aavev2LendingPool.Contract.GearboxAdapterType(&_Aavev2LendingPool.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) GearboxAdapterType() (uint8, error) {
	return _Aavev2LendingPool.Contract.GearboxAdapterType(&_Aavev2LendingPool.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) GearboxAdapterVersion() (uint16, error) {
	return _Aavev2LendingPool.Contract.GearboxAdapterVersion(&_Aavev2LendingPool.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Aavev2LendingPool.Contract.GearboxAdapterVersion(&_Aavev2LendingPool.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) Acl() (common.Address, error) {
	return _Aavev2LendingPool.Contract.Acl(&_Aavev2LendingPool.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) Acl() (common.Address, error) {
	return _Aavev2LendingPool.Contract.Acl(&_Aavev2LendingPool.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) AddressProvider() (common.Address, error) {
	return _Aavev2LendingPool.Contract.AddressProvider(&_Aavev2LendingPool.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) AddressProvider() (common.Address, error) {
	return _Aavev2LendingPool.Contract.AddressProvider(&_Aavev2LendingPool.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) CreditManager() (common.Address, error) {
	return _Aavev2LendingPool.Contract.CreditManager(&_Aavev2LendingPool.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) CreditManager() (common.Address, error) {
	return _Aavev2LendingPool.Contract.CreditManager(&_Aavev2LendingPool.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2LendingPool.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) TargetContract() (common.Address, error) {
	return _Aavev2LendingPool.Contract.TargetContract(&_Aavev2LendingPool.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2LendingPool *Aavev2LendingPoolCallerSession) TargetContract() (common.Address, error) {
	return _Aavev2LendingPool.Contract.TargetContract(&_Aavev2LendingPool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address , uint16 ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, arg2 common.Address, arg3 uint16) (*types.Transaction, error) {
	return _Aavev2LendingPool.contract.Transact(opts, "deposit", asset, amount, arg2, arg3)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address , uint16 ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) Deposit(asset common.Address, amount *big.Int, arg2 common.Address, arg3 uint16) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Deposit(&_Aavev2LendingPool.TransactOpts, asset, amount, arg2, arg3)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address , uint16 ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorSession) Deposit(asset common.Address, amount *big.Int, arg2 common.Address, arg3 uint16) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Deposit(&_Aavev2LendingPool.TransactOpts, asset, amount, arg2, arg3)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x055181cc.
//
// Solidity: function depositDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactor) DepositDiff(opts *bind.TransactOpts, asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.contract.Transact(opts, "depositDiff", asset, leftoverAmount)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x055181cc.
//
// Solidity: function depositDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) DepositDiff(asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.DepositDiff(&_Aavev2LendingPool.TransactOpts, asset, leftoverAmount)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x055181cc.
//
// Solidity: function depositDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorSession) DepositDiff(asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.DepositDiff(&_Aavev2LendingPool.TransactOpts, asset, leftoverAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Aavev2LendingPool.contract.Transact(opts, "withdraw", asset, amount, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) Withdraw(asset common.Address, amount *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Withdraw(&_Aavev2LendingPool.TransactOpts, asset, amount, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, arg2 common.Address) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.Withdraw(&_Aavev2LendingPool.TransactOpts, asset, amount, arg2)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x7bdd45b1.
//
// Solidity: function withdrawDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactor) WithdrawDiff(opts *bind.TransactOpts, asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.contract.Transact(opts, "withdrawDiff", asset, leftoverAmount)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x7bdd45b1.
//
// Solidity: function withdrawDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolSession) WithdrawDiff(asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.WithdrawDiff(&_Aavev2LendingPool.TransactOpts, asset, leftoverAmount)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x7bdd45b1.
//
// Solidity: function withdrawDiff(address asset, uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2LendingPool *Aavev2LendingPoolTransactorSession) WithdrawDiff(asset common.Address, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Aavev2LendingPool.Contract.WithdrawDiff(&_Aavev2LendingPool.TransactOpts, asset, leftoverAmount)
}
