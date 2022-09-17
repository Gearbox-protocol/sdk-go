// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nonFungiblePosManager

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

// INonfungiblePositionManagerMintParams is an auto generated low-level Go binding around an user-defined struct.
type INonfungiblePositionManagerMintParams struct {
	Token0         common.Address
	Token1         common.Address
	Fee            *big.Int
	TickLower      *big.Int
	TickUpper      *big.Int
	Amount0Desired *big.Int
	Amount1Desired *big.Int
	Amount0Min     *big.Int
	Amount1Min     *big.Int
	Recipient      common.Address
	Deadline       *big.Int
}

// NonFungiblePosManagerMetaData contains all meta data concerning the NonFungiblePosManager contract.
var NonFungiblePosManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createAndInitializePoolIfNecessary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structINonfungiblePositionManager.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NonFungiblePosManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NonFungiblePosManagerMetaData.ABI instead.
var NonFungiblePosManagerABI = NonFungiblePosManagerMetaData.ABI

// NonFungiblePosManager is an auto generated Go binding around an Ethereum contract.
type NonFungiblePosManager struct {
	NonFungiblePosManagerCaller     // Read-only binding to the contract
	NonFungiblePosManagerTransactor // Write-only binding to the contract
	NonFungiblePosManagerFilterer   // Log filterer for contract events
}

// NonFungiblePosManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonFungiblePosManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonFungiblePosManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonFungiblePosManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonFungiblePosManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonFungiblePosManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonFungiblePosManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonFungiblePosManagerSession struct {
	Contract     *NonFungiblePosManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NonFungiblePosManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonFungiblePosManagerCallerSession struct {
	Contract *NonFungiblePosManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// NonFungiblePosManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonFungiblePosManagerTransactorSession struct {
	Contract     *NonFungiblePosManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// NonFungiblePosManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonFungiblePosManagerRaw struct {
	Contract *NonFungiblePosManager // Generic contract binding to access the raw methods on
}

// NonFungiblePosManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonFungiblePosManagerCallerRaw struct {
	Contract *NonFungiblePosManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NonFungiblePosManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonFungiblePosManagerTransactorRaw struct {
	Contract *NonFungiblePosManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonFungiblePosManager creates a new instance of NonFungiblePosManager, bound to a specific deployed contract.
func NewNonFungiblePosManager(address common.Address, backend bind.ContractBackend) (*NonFungiblePosManager, error) {
	contract, err := bindNonFungiblePosManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonFungiblePosManager{NonFungiblePosManagerCaller: NonFungiblePosManagerCaller{contract: contract}, NonFungiblePosManagerTransactor: NonFungiblePosManagerTransactor{contract: contract}, NonFungiblePosManagerFilterer: NonFungiblePosManagerFilterer{contract: contract}}, nil
}

// NewNonFungiblePosManagerCaller creates a new read-only instance of NonFungiblePosManager, bound to a specific deployed contract.
func NewNonFungiblePosManagerCaller(address common.Address, caller bind.ContractCaller) (*NonFungiblePosManagerCaller, error) {
	contract, err := bindNonFungiblePosManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonFungiblePosManagerCaller{contract: contract}, nil
}

// NewNonFungiblePosManagerTransactor creates a new write-only instance of NonFungiblePosManager, bound to a specific deployed contract.
func NewNonFungiblePosManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NonFungiblePosManagerTransactor, error) {
	contract, err := bindNonFungiblePosManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonFungiblePosManagerTransactor{contract: contract}, nil
}

// NewNonFungiblePosManagerFilterer creates a new log filterer instance of NonFungiblePosManager, bound to a specific deployed contract.
func NewNonFungiblePosManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NonFungiblePosManagerFilterer, error) {
	contract, err := bindNonFungiblePosManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonFungiblePosManagerFilterer{contract: contract}, nil
}

// bindNonFungiblePosManager binds a generic wrapper to an already deployed contract.
func bindNonFungiblePosManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NonFungiblePosManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonFungiblePosManager *NonFungiblePosManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonFungiblePosManager.Contract.NonFungiblePosManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonFungiblePosManager *NonFungiblePosManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.NonFungiblePosManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonFungiblePosManager *NonFungiblePosManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.NonFungiblePosManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonFungiblePosManager *NonFungiblePosManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonFungiblePosManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonFungiblePosManager *NonFungiblePosManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonFungiblePosManager *NonFungiblePosManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.contract.Transact(opts, method, params...)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactor) CreateAndInitializePoolIfNecessary(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _NonFungiblePosManager.contract.Transact(opts, "createAndInitializePoolIfNecessary", token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_NonFungiblePosManager *NonFungiblePosManagerSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.CreateAndInitializePoolIfNecessary(&_NonFungiblePosManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// CreateAndInitializePoolIfNecessary is a paid mutator transaction binding the contract method 0x13ead562.
//
// Solidity: function createAndInitializePoolIfNecessary(address token0, address token1, uint24 fee, uint160 sqrtPriceX96) payable returns(address pool)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactorSession) CreateAndInitializePoolIfNecessary(token0 common.Address, token1 common.Address, fee *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.CreateAndInitializePoolIfNecessary(&_NonFungiblePosManager.TransactOpts, token0, token1, fee, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactor) Mint(opts *bind.TransactOpts, params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonFungiblePosManager.contract.Transact(opts, "mint", params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonFungiblePosManager *NonFungiblePosManagerSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.Mint(&_NonFungiblePosManager.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x88316456.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,uint256,uint256,address,uint256) params) payable returns(uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactorSession) Mint(params INonfungiblePositionManagerMintParams) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.Mint(&_NonFungiblePosManager.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _NonFungiblePosManager.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_NonFungiblePosManager *NonFungiblePosManagerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.Multicall(&_NonFungiblePosManager.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_NonFungiblePosManager *NonFungiblePosManagerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _NonFungiblePosManager.Contract.Multicall(&_NonFungiblePosManager.TransactOpts, data)
}
