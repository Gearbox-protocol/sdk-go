// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testAdapter

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

// TestAdapterABI is the input ABI used to generate the binding from.
const TestAdapterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"hack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestAdapter is an auto generated Go binding around an Ethereum contract.
type TestAdapter struct {
	TestAdapterCaller     // Read-only binding to the contract
	TestAdapterTransactor // Write-only binding to the contract
	TestAdapterFilterer   // Log filterer for contract events
}

// TestAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAdapterSession struct {
	Contract     *TestAdapter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAdapterCallerSession struct {
	Contract *TestAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAdapterTransactorSession struct {
	Contract     *TestAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestAdapterRaw struct {
	Contract *TestAdapter // Generic contract binding to access the raw methods on
}

// TestAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAdapterCallerRaw struct {
	Contract *TestAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// TestAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAdapterTransactorRaw struct {
	Contract *TestAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAdapter creates a new instance of TestAdapter, bound to a specific deployed contract.
func NewTestAdapter(address common.Address, backend bind.ContractBackend) (*TestAdapter, error) {
	contract, err := bindTestAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAdapter{TestAdapterCaller: TestAdapterCaller{contract: contract}, TestAdapterTransactor: TestAdapterTransactor{contract: contract}, TestAdapterFilterer: TestAdapterFilterer{contract: contract}}, nil
}

// NewTestAdapterCaller creates a new read-only instance of TestAdapter, bound to a specific deployed contract.
func NewTestAdapterCaller(address common.Address, caller bind.ContractCaller) (*TestAdapterCaller, error) {
	contract, err := bindTestAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAdapterCaller{contract: contract}, nil
}

// NewTestAdapterTransactor creates a new write-only instance of TestAdapter, bound to a specific deployed contract.
func NewTestAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestAdapterTransactor, error) {
	contract, err := bindTestAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAdapterTransactor{contract: contract}, nil
}

// NewTestAdapterFilterer creates a new log filterer instance of TestAdapter, bound to a specific deployed contract.
func NewTestAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestAdapterFilterer, error) {
	contract, err := bindTestAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAdapterFilterer{contract: contract}, nil
}

// bindTestAdapter binds a generic wrapper to an already deployed contract.
func bindTestAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAdapter *TestAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAdapter.Contract.TestAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAdapter *TestAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAdapter.Contract.TestAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAdapter *TestAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAdapter.Contract.TestAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAdapter *TestAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAdapter *TestAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAdapter *TestAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAdapter.Contract.contract.Transact(opts, method, params...)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_TestAdapter *TestAdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAdapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_TestAdapter *TestAdapterSession) CreditManager() (common.Address, error) {
	return _TestAdapter.Contract.CreditManager(&_TestAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_TestAdapter *TestAdapterCallerSession) CreditManager() (common.Address, error) {
	return _TestAdapter.Contract.CreditManager(&_TestAdapter.CallOpts)
}

// Hack is a paid mutator transaction binding the contract method 0x6c4c174f.
//
// Solidity: function hack(address token) returns()
func (_TestAdapter *TestAdapterTransactor) Hack(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TestAdapter.contract.Transact(opts, "hack", token)
}

// Hack is a paid mutator transaction binding the contract method 0x6c4c174f.
//
// Solidity: function hack(address token) returns()
func (_TestAdapter *TestAdapterSession) Hack(token common.Address) (*types.Transaction, error) {
	return _TestAdapter.Contract.Hack(&_TestAdapter.TransactOpts, token)
}

// Hack is a paid mutator transaction binding the contract method 0x6c4c174f.
//
// Solidity: function hack(address token) returns()
func (_TestAdapter *TestAdapterTransactorSession) Hack(token common.Address) (*types.Transaction, error) {
	return _TestAdapter.Contract.Hack(&_TestAdapter.TransactOpts, token)
}

// TransferMoney is a paid mutator transaction binding the contract method 0xe7d0aae9.
//
// Solidity: function transferMoney(address token, address to) returns()
func (_TestAdapter *TestAdapterTransactor) TransferMoney(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestAdapter.contract.Transact(opts, "transferMoney", token, to)
}

// TransferMoney is a paid mutator transaction binding the contract method 0xe7d0aae9.
//
// Solidity: function transferMoney(address token, address to) returns()
func (_TestAdapter *TestAdapterSession) TransferMoney(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestAdapter.Contract.TransferMoney(&_TestAdapter.TransactOpts, token, to)
}

// TransferMoney is a paid mutator transaction binding the contract method 0xe7d0aae9.
//
// Solidity: function transferMoney(address token, address to) returns()
func (_TestAdapter *TestAdapterTransactorSession) TransferMoney(token common.Address, to common.Address) (*types.Transaction, error) {
	return _TestAdapter.Contract.TransferMoney(&_TestAdapter.TransactOpts, token, to)
}
