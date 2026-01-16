// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balanceGateway

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

// BalanceGatewayMetaData contains all meta data concerning the BalanceGateway contract.
var BalanceGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_balancerV3Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permit2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"exactAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minBptAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"wethIsEth\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"addLiquidityUnbalanced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptAmountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balancerV3Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactBptAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"wethIsEth\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"removeLiquiditySingleTokenExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"wethIsEth\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"swapSingleTokenExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BalanceGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceGatewayMetaData.ABI instead.
var BalanceGatewayABI = BalanceGatewayMetaData.ABI

// BalanceGateway is an auto generated Go binding around an Ethereum contract.
type BalanceGateway struct {
	BalanceGatewayCaller     // Read-only binding to the contract
	BalanceGatewayTransactor // Write-only binding to the contract
	BalanceGatewayFilterer   // Log filterer for contract events
}

// BalanceGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceGatewaySession struct {
	Contract     *BalanceGateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceGatewayCallerSession struct {
	Contract *BalanceGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalanceGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceGatewayTransactorSession struct {
	Contract     *BalanceGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalanceGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceGatewayRaw struct {
	Contract *BalanceGateway // Generic contract binding to access the raw methods on
}

// BalanceGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceGatewayCallerRaw struct {
	Contract *BalanceGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceGatewayTransactorRaw struct {
	Contract *BalanceGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceGateway creates a new instance of BalanceGateway, bound to a specific deployed contract.
func NewBalanceGateway(address common.Address, backend bind.ContractBackend) (*BalanceGateway, error) {
	contract, err := bindBalanceGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceGateway{BalanceGatewayCaller: BalanceGatewayCaller{contract: contract}, BalanceGatewayTransactor: BalanceGatewayTransactor{contract: contract}, BalanceGatewayFilterer: BalanceGatewayFilterer{contract: contract}}, nil
}

// NewBalanceGatewayCaller creates a new read-only instance of BalanceGateway, bound to a specific deployed contract.
func NewBalanceGatewayCaller(address common.Address, caller bind.ContractCaller) (*BalanceGatewayCaller, error) {
	contract, err := bindBalanceGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceGatewayCaller{contract: contract}, nil
}

// NewBalanceGatewayTransactor creates a new write-only instance of BalanceGateway, bound to a specific deployed contract.
func NewBalanceGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceGatewayTransactor, error) {
	contract, err := bindBalanceGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceGatewayTransactor{contract: contract}, nil
}

// NewBalanceGatewayFilterer creates a new log filterer instance of BalanceGateway, bound to a specific deployed contract.
func NewBalanceGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceGatewayFilterer, error) {
	contract, err := bindBalanceGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceGatewayFilterer{contract: contract}, nil
}

// bindBalanceGateway binds a generic wrapper to an already deployed contract.
func bindBalanceGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceGateway *BalanceGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceGateway.Contract.BalanceGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceGateway *BalanceGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceGateway.Contract.BalanceGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceGateway *BalanceGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceGateway.Contract.BalanceGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceGateway *BalanceGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceGateway *BalanceGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceGateway *BalanceGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceGateway.Contract.contract.Transact(opts, method, params...)
}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_BalanceGateway *BalanceGatewayCaller) BalancerV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalanceGateway.contract.Call(opts, &out, "balancerV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_BalanceGateway *BalanceGatewaySession) BalancerV3Router() (common.Address, error) {
	return _BalanceGateway.Contract.BalancerV3Router(&_BalanceGateway.CallOpts)
}

// BalancerV3Router is a free data retrieval call binding the contract method 0x7f4c2f10.
//
// Solidity: function balancerV3Router() view returns(address)
func (_BalanceGateway *BalanceGatewayCallerSession) BalancerV3Router() (common.Address, error) {
	return _BalanceGateway.Contract.BalancerV3Router(&_BalanceGateway.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_BalanceGateway *BalanceGatewayCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalanceGateway.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_BalanceGateway *BalanceGatewaySession) ContractType() ([32]byte, error) {
	return _BalanceGateway.Contract.ContractType(&_BalanceGateway.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_BalanceGateway *BalanceGatewayCallerSession) ContractType() ([32]byte, error) {
	return _BalanceGateway.Contract.ContractType(&_BalanceGateway.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_BalanceGateway *BalanceGatewayCaller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalanceGateway.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_BalanceGateway *BalanceGatewaySession) Permit2() (common.Address, error) {
	return _BalanceGateway.Contract.Permit2(&_BalanceGateway.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_BalanceGateway *BalanceGatewayCallerSession) Permit2() (common.Address, error) {
	return _BalanceGateway.Contract.Permit2(&_BalanceGateway.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BalanceGateway *BalanceGatewayCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalanceGateway.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BalanceGateway *BalanceGatewaySession) Version() (*big.Int, error) {
	return _BalanceGateway.Contract.Version(&_BalanceGateway.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BalanceGateway *BalanceGatewayCallerSession) Version() (*big.Int, error) {
	return _BalanceGateway.Contract.Version(&_BalanceGateway.CallOpts)
}

// AddLiquidityUnbalanced is a paid mutator transaction binding the contract method 0xc08bc851.
//
// Solidity: function addLiquidityUnbalanced(address pool, uint256[] exactAmountsIn, uint256 minBptAmountOut, bool wethIsEth, bytes userData) returns(uint256 bptAmountOut)
func (_BalanceGateway *BalanceGatewayTransactor) AddLiquidityUnbalanced(opts *bind.TransactOpts, pool common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.contract.Transact(opts, "addLiquidityUnbalanced", pool, exactAmountsIn, minBptAmountOut, wethIsEth, userData)
}

// AddLiquidityUnbalanced is a paid mutator transaction binding the contract method 0xc08bc851.
//
// Solidity: function addLiquidityUnbalanced(address pool, uint256[] exactAmountsIn, uint256 minBptAmountOut, bool wethIsEth, bytes userData) returns(uint256 bptAmountOut)
func (_BalanceGateway *BalanceGatewaySession) AddLiquidityUnbalanced(pool common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.AddLiquidityUnbalanced(&_BalanceGateway.TransactOpts, pool, exactAmountsIn, minBptAmountOut, wethIsEth, userData)
}

// AddLiquidityUnbalanced is a paid mutator transaction binding the contract method 0xc08bc851.
//
// Solidity: function addLiquidityUnbalanced(address pool, uint256[] exactAmountsIn, uint256 minBptAmountOut, bool wethIsEth, bytes userData) returns(uint256 bptAmountOut)
func (_BalanceGateway *BalanceGatewayTransactorSession) AddLiquidityUnbalanced(pool common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.AddLiquidityUnbalanced(&_BalanceGateway.TransactOpts, pool, exactAmountsIn, minBptAmountOut, wethIsEth, userData)
}

// RemoveLiquiditySingleTokenExactIn is a paid mutator transaction binding the contract method 0xecb2182c.
//
// Solidity: function removeLiquiditySingleTokenExactIn(address pool, uint256 exactBptAmountIn, address tokenOut, uint256 minAmountOut, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewayTransactor) RemoveLiquiditySingleTokenExactIn(opts *bind.TransactOpts, pool common.Address, exactBptAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.contract.Transact(opts, "removeLiquiditySingleTokenExactIn", pool, exactBptAmountIn, tokenOut, minAmountOut, wethIsEth, userData)
}

// RemoveLiquiditySingleTokenExactIn is a paid mutator transaction binding the contract method 0xecb2182c.
//
// Solidity: function removeLiquiditySingleTokenExactIn(address pool, uint256 exactBptAmountIn, address tokenOut, uint256 minAmountOut, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewaySession) RemoveLiquiditySingleTokenExactIn(pool common.Address, exactBptAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.RemoveLiquiditySingleTokenExactIn(&_BalanceGateway.TransactOpts, pool, exactBptAmountIn, tokenOut, minAmountOut, wethIsEth, userData)
}

// RemoveLiquiditySingleTokenExactIn is a paid mutator transaction binding the contract method 0xecb2182c.
//
// Solidity: function removeLiquiditySingleTokenExactIn(address pool, uint256 exactBptAmountIn, address tokenOut, uint256 minAmountOut, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewayTransactorSession) RemoveLiquiditySingleTokenExactIn(pool common.Address, exactBptAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.RemoveLiquiditySingleTokenExactIn(&_BalanceGateway.TransactOpts, pool, exactBptAmountIn, tokenOut, minAmountOut, wethIsEth, userData)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewayTransactor) SwapSingleTokenExactIn(opts *bind.TransactOpts, pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.contract.Transact(opts, "swapSingleTokenExactIn", pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewaySession) SwapSingleTokenExactIn(pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.SwapSingleTokenExactIn(&_BalanceGateway.TransactOpts, pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}

// SwapSingleTokenExactIn is a paid mutator transaction binding the contract method 0x750283bc.
//
// Solidity: function swapSingleTokenExactIn(address pool, address tokenIn, address tokenOut, uint256 exactAmountIn, uint256 minAmountOut, uint256 deadline, bool wethIsEth, bytes userData) returns(uint256 amountOut)
func (_BalanceGateway *BalanceGatewayTransactorSession) SwapSingleTokenExactIn(pool common.Address, tokenIn common.Address, tokenOut common.Address, exactAmountIn *big.Int, minAmountOut *big.Int, deadline *big.Int, wethIsEth bool, userData []byte) (*types.Transaction, error) {
	return _BalanceGateway.Contract.SwapSingleTokenExactIn(&_BalanceGateway.TransactOpts, pool, tokenIn, tokenOut, exactAmountIn, minAmountOut, deadline, wethIsEth, userData)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceGateway *BalanceGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceGateway *BalanceGatewaySession) Receive() (*types.Transaction, error) {
	return _BalanceGateway.Contract.Receive(&_BalanceGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BalanceGateway *BalanceGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _BalanceGateway.Contract.Receive(&_BalanceGateway.TransactOpts)
}
