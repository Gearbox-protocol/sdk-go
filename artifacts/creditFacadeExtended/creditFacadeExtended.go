// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFacadeExtended

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

// Balance is an auto generated low-level Go binding around an user-defined struct.
type Balance struct {
	Token   common.Address
	Balance *big.Int
}

// CreditFacadeExtendedMetaData contains all meta data concerning the CreditFacadeExtended contract.
var CreditFacadeExtendedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"expected\",\"type\":\"tuple[]\"}],\"name\":\"revertIfReceivedLessThan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreditFacadeExtendedABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditFacadeExtendedMetaData.ABI instead.
var CreditFacadeExtendedABI = CreditFacadeExtendedMetaData.ABI

// CreditFacadeExtended is an auto generated Go binding around an Ethereum contract.
type CreditFacadeExtended struct {
	CreditFacadeExtendedCaller     // Read-only binding to the contract
	CreditFacadeExtendedTransactor // Write-only binding to the contract
	CreditFacadeExtendedFilterer   // Log filterer for contract events
}

// CreditFacadeExtendedCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFacadeExtendedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeExtendedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFacadeExtendedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeExtendedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFacadeExtendedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadeExtendedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFacadeExtendedSession struct {
	Contract     *CreditFacadeExtended // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CreditFacadeExtendedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFacadeExtendedCallerSession struct {
	Contract *CreditFacadeExtendedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CreditFacadeExtendedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFacadeExtendedTransactorSession struct {
	Contract     *CreditFacadeExtendedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CreditFacadeExtendedRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFacadeExtendedRaw struct {
	Contract *CreditFacadeExtended // Generic contract binding to access the raw methods on
}

// CreditFacadeExtendedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFacadeExtendedCallerRaw struct {
	Contract *CreditFacadeExtendedCaller // Generic read-only contract binding to access the raw methods on
}

// CreditFacadeExtendedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFacadeExtendedTransactorRaw struct {
	Contract *CreditFacadeExtendedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFacadeExtended creates a new instance of CreditFacadeExtended, bound to a specific deployed contract.
func NewCreditFacadeExtended(address common.Address, backend bind.ContractBackend) (*CreditFacadeExtended, error) {
	contract, err := bindCreditFacadeExtended(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeExtended{CreditFacadeExtendedCaller: CreditFacadeExtendedCaller{contract: contract}, CreditFacadeExtendedTransactor: CreditFacadeExtendedTransactor{contract: contract}, CreditFacadeExtendedFilterer: CreditFacadeExtendedFilterer{contract: contract}}, nil
}

// NewCreditFacadeExtendedCaller creates a new read-only instance of CreditFacadeExtended, bound to a specific deployed contract.
func NewCreditFacadeExtendedCaller(address common.Address, caller bind.ContractCaller) (*CreditFacadeExtendedCaller, error) {
	contract, err := bindCreditFacadeExtended(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeExtendedCaller{contract: contract}, nil
}

// NewCreditFacadeExtendedTransactor creates a new write-only instance of CreditFacadeExtended, bound to a specific deployed contract.
func NewCreditFacadeExtendedTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditFacadeExtendedTransactor, error) {
	contract, err := bindCreditFacadeExtended(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeExtendedTransactor{contract: contract}, nil
}

// NewCreditFacadeExtendedFilterer creates a new log filterer instance of CreditFacadeExtended, bound to a specific deployed contract.
func NewCreditFacadeExtendedFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFacadeExtendedFilterer, error) {
	contract, err := bindCreditFacadeExtended(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFacadeExtendedFilterer{contract: contract}, nil
}

// bindCreditFacadeExtended binds a generic wrapper to an already deployed contract.
func bindCreditFacadeExtended(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditFacadeExtendedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadeExtended *CreditFacadeExtendedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadeExtended.Contract.CreditFacadeExtendedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadeExtended *CreditFacadeExtendedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.CreditFacadeExtendedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadeExtended *CreditFacadeExtendedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.CreditFacadeExtendedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadeExtended *CreditFacadeExtendedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadeExtended.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadeExtended *CreditFacadeExtendedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadeExtended *CreditFacadeExtendedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.contract.Transact(opts, method, params...)
}

// RevertIfReceivedLessThan is a paid mutator transaction binding the contract method 0x81314b59.
//
// Solidity: function revertIfReceivedLessThan((address,uint256)[] expected) returns()
func (_CreditFacadeExtended *CreditFacadeExtendedTransactor) RevertIfReceivedLessThan(opts *bind.TransactOpts, expected []Balance) (*types.Transaction, error) {
	return _CreditFacadeExtended.contract.Transact(opts, "revertIfReceivedLessThan", expected)
}

// RevertIfReceivedLessThan is a paid mutator transaction binding the contract method 0x81314b59.
//
// Solidity: function revertIfReceivedLessThan((address,uint256)[] expected) returns()
func (_CreditFacadeExtended *CreditFacadeExtendedSession) RevertIfReceivedLessThan(expected []Balance) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.RevertIfReceivedLessThan(&_CreditFacadeExtended.TransactOpts, expected)
}

// RevertIfReceivedLessThan is a paid mutator transaction binding the contract method 0x81314b59.
//
// Solidity: function revertIfReceivedLessThan((address,uint256)[] expected) returns()
func (_CreditFacadeExtended *CreditFacadeExtendedTransactorSession) RevertIfReceivedLessThan(expected []Balance) (*types.Transaction, error) {
	return _CreditFacadeExtended.Contract.RevertIfReceivedLessThan(&_CreditFacadeExtended.TransactOpts, expected)
}
