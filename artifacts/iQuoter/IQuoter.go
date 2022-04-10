// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iQuoter

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

// IQuoterABI is the input ABI used to generate the binding from.
const IQuoterABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IQuoter is an auto generated Go binding around an Ethereum contract.
type IQuoter struct {
	IQuoterCaller     // Read-only binding to the contract
	IQuoterTransactor // Write-only binding to the contract
	IQuoterFilterer   // Log filterer for contract events
}

// IQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IQuoterSession struct {
	Contract     *IQuoter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IQuoterCallerSession struct {
	Contract *IQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IQuoterTransactorSession struct {
	Contract     *IQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IQuoterRaw struct {
	Contract *IQuoter // Generic contract binding to access the raw methods on
}

// IQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IQuoterCallerRaw struct {
	Contract *IQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// IQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IQuoterTransactorRaw struct {
	Contract *IQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIQuoter creates a new instance of IQuoter, bound to a specific deployed contract.
func NewIQuoter(address common.Address, backend bind.ContractBackend) (*IQuoter, error) {
	contract, err := bindIQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IQuoter{IQuoterCaller: IQuoterCaller{contract: contract}, IQuoterTransactor: IQuoterTransactor{contract: contract}, IQuoterFilterer: IQuoterFilterer{contract: contract}}, nil
}

// NewIQuoterCaller creates a new read-only instance of IQuoter, bound to a specific deployed contract.
func NewIQuoterCaller(address common.Address, caller bind.ContractCaller) (*IQuoterCaller, error) {
	contract, err := bindIQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IQuoterCaller{contract: contract}, nil
}

// NewIQuoterTransactor creates a new write-only instance of IQuoter, bound to a specific deployed contract.
func NewIQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*IQuoterTransactor, error) {
	contract, err := bindIQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IQuoterTransactor{contract: contract}, nil
}

// NewIQuoterFilterer creates a new log filterer instance of IQuoter, bound to a specific deployed contract.
func NewIQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*IQuoterFilterer, error) {
	contract, err := bindIQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IQuoterFilterer{contract: contract}, nil
}

// bindIQuoter binds a generic wrapper to an already deployed contract.
func bindIQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IQuoterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuoter *IQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuoter.Contract.IQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuoter *IQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuoter.Contract.IQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuoter *IQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuoter.Contract.IQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IQuoter *IQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IQuoter *IQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IQuoter *IQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IQuoter.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_IQuoter *IQuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_IQuoter *IQuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuoter.Contract.QuoteExactInput(&_IQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_IQuoter *IQuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _IQuoter.Contract.QuoteExactInput(&_IQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)

func (_IQuoter *IQuoterCaller) QuoteExactInputSingle(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IQuoter.contract.Call(opts, &out, "quoteExactInputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}
func (_IQuoter *IQuoterCaller) QuoteExactOutputSingle(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IQuoter.contract.Call(opts, &out, "quoteExactOutputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return out0, err
}

