// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalAdapter

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
)

// RevocationPair is an auto generated low-level Go binding around an user-defined struct.
type RevocationPair struct {
	Spender common.Address
	Token   common.Address
}

// UniversalAdapterMetaData contains all meta data concerning the UniversalAdapter contract.
var UniversalAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"UnexpectedCreditAccountException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structRevocationPair[]\",\"name\":\"revocations\",\"type\":\"tuple[]\"}],\"name\":\"revokeAdapterAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structRevocationPair[]\",\"name\":\"revocations\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"expectedCreditAccount\",\"type\":\"address\"}],\"name\":\"revokeAdapterAllowances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniversalAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalAdapterMetaData.ABI instead.
var UniversalAdapterABI = UniversalAdapterMetaData.ABI

// UniversalAdapter is an auto generated Go binding around an Ethereum contract.
type UniversalAdapter struct {
	UniversalAdapterCaller     // Read-only binding to the contract
	UniversalAdapterTransactor // Write-only binding to the contract
	UniversalAdapterFilterer   // Log filterer for contract events
}

// UniversalAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalAdapterSession struct {
	Contract     *UniversalAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniversalAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalAdapterCallerSession struct {
	Contract *UniversalAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UniversalAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalAdapterTransactorSession struct {
	Contract     *UniversalAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UniversalAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalAdapterRaw struct {
	Contract *UniversalAdapter // Generic contract binding to access the raw methods on
}

// UniversalAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalAdapterCallerRaw struct {
	Contract *UniversalAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalAdapterTransactorRaw struct {
	Contract *UniversalAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalAdapter creates a new instance of UniversalAdapter, bound to a specific deployed contract.
func NewUniversalAdapter(address common.Address, backend bind.ContractBackend) (*UniversalAdapter, error) {
	contract, err := bindUniversalAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalAdapter{UniversalAdapterCaller: UniversalAdapterCaller{contract: contract}, UniversalAdapterTransactor: UniversalAdapterTransactor{contract: contract}, UniversalAdapterFilterer: UniversalAdapterFilterer{contract: contract}}, nil
}

// NewUniversalAdapterCaller creates a new read-only instance of UniversalAdapter, bound to a specific deployed contract.
func NewUniversalAdapterCaller(address common.Address, caller bind.ContractCaller) (*UniversalAdapterCaller, error) {
	contract, err := bindUniversalAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalAdapterCaller{contract: contract}, nil
}

// NewUniversalAdapterTransactor creates a new write-only instance of UniversalAdapter, bound to a specific deployed contract.
func NewUniversalAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalAdapterTransactor, error) {
	contract, err := bindUniversalAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalAdapterTransactor{contract: contract}, nil
}

// NewUniversalAdapterFilterer creates a new log filterer instance of UniversalAdapter, bound to a specific deployed contract.
func NewUniversalAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalAdapterFilterer, error) {
	contract, err := bindUniversalAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalAdapterFilterer{contract: contract}, nil
}

// bindUniversalAdapter binds a generic wrapper to an already deployed contract.
func bindUniversalAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniversalAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalAdapter *UniversalAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalAdapter.Contract.UniversalAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalAdapter *UniversalAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.UniversalAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalAdapter *UniversalAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.UniversalAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalAdapter *UniversalAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalAdapter *UniversalAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalAdapter *UniversalAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_UniversalAdapter *UniversalAdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UniversalAdapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_UniversalAdapter *UniversalAdapterSession) GearboxAdapterType() (uint8, error) {
	return _UniversalAdapter.Contract.GearboxAdapterType(&_UniversalAdapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_UniversalAdapter *UniversalAdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _UniversalAdapter.Contract.GearboxAdapterType(&_UniversalAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_UniversalAdapter *UniversalAdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _UniversalAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_UniversalAdapter *UniversalAdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _UniversalAdapter.Contract.GearboxAdapterVersion(&_UniversalAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_UniversalAdapter *UniversalAdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _UniversalAdapter.Contract.GearboxAdapterVersion(&_UniversalAdapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_UniversalAdapter *UniversalAdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalAdapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_UniversalAdapter *UniversalAdapterSession) CreditFacade() (common.Address, error) {
	return _UniversalAdapter.Contract.CreditFacade(&_UniversalAdapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_UniversalAdapter *UniversalAdapterCallerSession) CreditFacade() (common.Address, error) {
	return _UniversalAdapter.Contract.CreditFacade(&_UniversalAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniversalAdapter *UniversalAdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalAdapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniversalAdapter *UniversalAdapterSession) CreditManager() (common.Address, error) {
	return _UniversalAdapter.Contract.CreditManager(&_UniversalAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_UniversalAdapter *UniversalAdapterCallerSession) CreditManager() (common.Address, error) {
	return _UniversalAdapter.Contract.CreditManager(&_UniversalAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_UniversalAdapter *UniversalAdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalAdapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_UniversalAdapter *UniversalAdapterSession) TargetContract() (common.Address, error) {
	return _UniversalAdapter.Contract.TargetContract(&_UniversalAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_UniversalAdapter *UniversalAdapterCallerSession) TargetContract() (common.Address, error) {
	return _UniversalAdapter.Contract.TargetContract(&_UniversalAdapter.CallOpts)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x565a820d.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations) returns()
func (_UniversalAdapter *UniversalAdapterTransactor) RevokeAdapterAllowances(opts *bind.TransactOpts, revocations []RevocationPair) (*types.Transaction, error) {
	return _UniversalAdapter.contract.Transact(opts, "revokeAdapterAllowances", revocations)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x565a820d.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations) returns()
func (_UniversalAdapter *UniversalAdapterSession) RevokeAdapterAllowances(revocations []RevocationPair) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.RevokeAdapterAllowances(&_UniversalAdapter.TransactOpts, revocations)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x565a820d.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations) returns()
func (_UniversalAdapter *UniversalAdapterTransactorSession) RevokeAdapterAllowances(revocations []RevocationPair) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.RevokeAdapterAllowances(&_UniversalAdapter.TransactOpts, revocations)
}

// RevokeAdapterAllowances0 is a paid mutator transaction binding the contract method 0xe6d6b153.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations, address expectedCreditAccount) returns()
func (_UniversalAdapter *UniversalAdapterTransactor) RevokeAdapterAllowances0(opts *bind.TransactOpts, revocations []RevocationPair, expectedCreditAccount common.Address) (*types.Transaction, error) {
	return _UniversalAdapter.contract.Transact(opts, "revokeAdapterAllowances0", revocations, expectedCreditAccount)
}

// RevokeAdapterAllowances0 is a paid mutator transaction binding the contract method 0xe6d6b153.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations, address expectedCreditAccount) returns()
func (_UniversalAdapter *UniversalAdapterSession) RevokeAdapterAllowances0(revocations []RevocationPair, expectedCreditAccount common.Address) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.RevokeAdapterAllowances0(&_UniversalAdapter.TransactOpts, revocations, expectedCreditAccount)
}

// RevokeAdapterAllowances0 is a paid mutator transaction binding the contract method 0xe6d6b153.
//
// Solidity: function revokeAdapterAllowances((address,address)[] revocations, address expectedCreditAccount) returns()
func (_UniversalAdapter *UniversalAdapterTransactorSession) RevokeAdapterAllowances0(revocations []RevocationPair, expectedCreditAccount common.Address) (*types.Transaction, error) {
	return _UniversalAdapter.Contract.RevokeAdapterAllowances0(&_UniversalAdapter.TransactOpts, revocations, expectedCreditAccount)
}
