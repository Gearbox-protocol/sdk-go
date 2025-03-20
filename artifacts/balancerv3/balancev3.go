// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerv3

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

// Balancerv3MetaData contains all meta data concerning the Balancerv3 contract.
var Balancerv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_balancerV3Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permit2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"balancerV3Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"wethIsEth\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"swapSingleTokenExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Balancerv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Balancerv3MetaData.ABI instead.
var Balancerv3ABI = Balancerv3MetaData.ABI

// Balancerv3 is an auto generated Go binding around an Ethereum contract.
type Balancerv3 struct {
	Balancerv3Caller     // Read-only binding to the contract
	Balancerv3Transactor // Write-only binding to the contract
	Balancerv3Filterer   // Log filterer for contract events
}

// Balancerv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Balancerv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Balancerv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Balancerv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Balancerv3Session struct {
	Contract     *Balancerv3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Balancerv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Balancerv3CallerSession struct {
	Contract *Balancerv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Balancerv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Balancerv3TransactorSession struct {
	Contract     *Balancerv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Balancerv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Balancerv3Raw struct {
	Contract *Balancerv3 // Generic contract binding to access the raw methods on
}

// Balancerv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Balancerv3CallerRaw struct {
	Contract *Balancerv3Caller // Generic read-only contract binding to access the raw methods on
}

// Balancerv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Balancerv3TransactorRaw struct {
	Contract *Balancerv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerv3 creates a new instance of Balancerv3, bound to a specific deployed contract.
func NewBalancerv3(address common.Address, backend bind.ContractBackend) (*Balancerv3, error) {
	contract, err := bindBalancerv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerv3{Balancerv3Caller: Balancerv3Caller{contract: contract}, Balancerv3Transactor: Balancerv3Transactor{contract: contract}, Balancerv3Filterer: Balancerv3Filterer{contract: contract}}, nil
}

// NewBalancerv3Caller creates a new read-only instance of Balancerv3, bound to a specific deployed contract.
func NewBalancerv3Caller(address common.Address, caller bind.ContractCaller) (*Balancerv3Caller, error) {
	contract, err := bindBalancerv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv3Caller{contract: contract}, nil
}

// NewBalancerv3Transactor creates a new write-only instance of Balancerv3, bound to a specific deployed contract.
func NewBalancerv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Balancerv3Transactor, error) {
	contract, err := bindBalancerv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv3Transactor{contract: contract}, nil
}

// NewBalancerv3Filterer creates a new log filterer instance of Balancerv3, bound to a specific deployed contract.
func NewBalancerv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Balancerv3Filterer, error) {
	contract, err := bindBalancerv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Balancerv3Filterer{contract: contract}, nil
}

// bindBalancerv3 binds a generic wrapper to an already deployed contract.
func bindBalancerv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Balancerv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv3 *Balancerv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv3.Contract.Balancerv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv3 *Balancerv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv3.Contract.Balancerv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv3 *Balancerv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv3.Contract.Balancerv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv3 *Balancerv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv3 *Balancerv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv3 *Balancerv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv3.Contract.contract.Transact(opts, method, params...)
}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_Balancerv3 *Balancerv3Caller) BalancerV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv3.contract.Call(opts, &out, "balancerV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_Balancerv3 *Balancerv3Session) BalancerV3Router() (common.Address, error) {
	return _Balancerv3.Contract.BalancerV3Router(&_Balancerv3.CallOpts)
}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_Balancerv3 *Balancerv3CallerSession) BalancerV3Router() (common.Address, error) {
	return _Balancerv3.Contract.BalancerV3Router(&_Balancerv3.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Balancerv3 *Balancerv3Caller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv3.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Balancerv3 *Balancerv3Session) Permit2() (common.Address, error) {
	return _Balancerv3.Contract.Permit2(&_Balancerv3.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_Balancerv3 *Balancerv3CallerSession) Permit2() (common.Address, error) {
	return _Balancerv3.Contract.Permit2(&_Balancerv3.CallOpts)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_Balancerv3 *Balancerv3Transactor) SwapSingleTokenExactIn(opts *bind.TransactOpts, pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _Balancerv3.contract.Transact(opts, "swapSingleTokenExactIn", pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_Balancerv3 *Balancerv3Session) SwapSingleTokenExactIn(pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _Balancerv3.Contract.SwapSingleTokenExactIn(&_Balancerv3.TransactOpts, pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_Balancerv3 *Balancerv3TransactorSession) SwapSingleTokenExactIn(pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _Balancerv3.Contract.SwapSingleTokenExactIn(&_Balancerv3.TransactOpts, pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}
