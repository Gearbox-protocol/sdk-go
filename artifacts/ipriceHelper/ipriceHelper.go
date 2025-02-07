// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipriceHelper

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

// PriceOnDemand is an auto generated low-level Go binding around an user-defined struct.
type PriceOnDemand struct {
	Token    common.Address
	CallData []byte
}

// TokenPriceInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenPriceInfo struct {
	Token                common.Address
	Balance              *big.Int
	BalanceInUnderlying  *big.Int
	LiquidationThreshold *big.Int
}

// IpriceHelperMetaData contains all meta data concerning the IpriceHelper contract.
var IpriceHelperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"previewTokens\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceOnDemand[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPriceInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balanceInUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"}]",
}

// IpriceHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use IpriceHelperMetaData.ABI instead.
var IpriceHelperABI = IpriceHelperMetaData.ABI

// IpriceHelper is an auto generated Go binding around an Ethereum contract.
type IpriceHelper struct {
	IpriceHelperCaller     // Read-only binding to the contract
	IpriceHelperTransactor // Write-only binding to the contract
	IpriceHelperFilterer   // Log filterer for contract events
}

// IpriceHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpriceHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpriceHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpriceHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpriceHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpriceHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpriceHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpriceHelperSession struct {
	Contract     *IpriceHelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpriceHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpriceHelperCallerSession struct {
	Contract *IpriceHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IpriceHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpriceHelperTransactorSession struct {
	Contract     *IpriceHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IpriceHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpriceHelperRaw struct {
	Contract *IpriceHelper // Generic contract binding to access the raw methods on
}

// IpriceHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpriceHelperCallerRaw struct {
	Contract *IpriceHelperCaller // Generic read-only contract binding to access the raw methods on
}

// IpriceHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpriceHelperTransactorRaw struct {
	Contract *IpriceHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpriceHelper creates a new instance of IpriceHelper, bound to a specific deployed contract.
func NewIpriceHelper(address common.Address, backend bind.ContractBackend) (*IpriceHelper, error) {
	contract, err := bindIpriceHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IpriceHelper{IpriceHelperCaller: IpriceHelperCaller{contract: contract}, IpriceHelperTransactor: IpriceHelperTransactor{contract: contract}, IpriceHelperFilterer: IpriceHelperFilterer{contract: contract}}, nil
}

// NewIpriceHelperCaller creates a new read-only instance of IpriceHelper, bound to a specific deployed contract.
func NewIpriceHelperCaller(address common.Address, caller bind.ContractCaller) (*IpriceHelperCaller, error) {
	contract, err := bindIpriceHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpriceHelperCaller{contract: contract}, nil
}

// NewIpriceHelperTransactor creates a new write-only instance of IpriceHelper, bound to a specific deployed contract.
func NewIpriceHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*IpriceHelperTransactor, error) {
	contract, err := bindIpriceHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpriceHelperTransactor{contract: contract}, nil
}

// NewIpriceHelperFilterer creates a new log filterer instance of IpriceHelper, bound to a specific deployed contract.
func NewIpriceHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*IpriceHelperFilterer, error) {
	contract, err := bindIpriceHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpriceHelperFilterer{contract: contract}, nil
}

// bindIpriceHelper binds a generic wrapper to an already deployed contract.
func bindIpriceHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpriceHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IpriceHelper *IpriceHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IpriceHelper.Contract.IpriceHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IpriceHelper *IpriceHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IpriceHelper.Contract.IpriceHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IpriceHelper *IpriceHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IpriceHelper.Contract.IpriceHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IpriceHelper *IpriceHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IpriceHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IpriceHelper *IpriceHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IpriceHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IpriceHelper *IpriceHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IpriceHelper.Contract.contract.Transact(opts, method, params...)
}

// PreviewTokens is a paid mutator transaction binding the contract method 0x324d1c8e.
//
// Solidity: function previewTokens(address creditAccount, (address,bytes)[] priceUpdates) returns((address,uint256,uint256,uint256)[] results)
func (_IpriceHelper *IpriceHelperTransactor) PreviewTokens(opts *bind.TransactOpts, creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _IpriceHelper.contract.Transact(opts, "previewTokens", creditAccount, priceUpdates)
}

// PreviewTokens is a paid mutator transaction binding the contract method 0x324d1c8e.
//
// Solidity: function previewTokens(address creditAccount, (address,bytes)[] priceUpdates) returns((address,uint256,uint256,uint256)[] results)
func (_IpriceHelper *IpriceHelperSession) PreviewTokens(creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _IpriceHelper.Contract.PreviewTokens(&_IpriceHelper.TransactOpts, creditAccount, priceUpdates)
}

// PreviewTokens is a paid mutator transaction binding the contract method 0x324d1c8e.
//
// Solidity: function previewTokens(address creditAccount, (address,bytes)[] priceUpdates) returns((address,uint256,uint256,uint256)[] results)
func (_IpriceHelper *IpriceHelperTransactorSession) PreviewTokens(creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _IpriceHelper.Contract.PreviewTokens(&_IpriceHelper.TransactOpts, creditAccount, priceUpdates)
}
