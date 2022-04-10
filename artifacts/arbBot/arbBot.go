// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbBot

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

// ArbBotABI is the input ABI used to generate the binding from.
const ArbBotABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wethToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"allowExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"allowRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"checkUniV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"faucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"forbidExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"updatePriceV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v3Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"v3ethToToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v3Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"v3tokenToEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ArbBot is an auto generated Go binding around an Ethereum contract.
type ArbBot struct {
	ArbBotCaller     // Read-only binding to the contract
	ArbBotTransactor // Write-only binding to the contract
	ArbBotFilterer   // Log filterer for contract events
}

// ArbBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbBotSession struct {
	Contract     *ArbBot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbBotCallerSession struct {
	Contract *ArbBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbBotTransactorSession struct {
	Contract     *ArbBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbBotRaw struct {
	Contract *ArbBot // Generic contract binding to access the raw methods on
}

// ArbBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbBotCallerRaw struct {
	Contract *ArbBotCaller // Generic read-only contract binding to access the raw methods on
}

// ArbBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbBotTransactorRaw struct {
	Contract *ArbBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbBot creates a new instance of ArbBot, bound to a specific deployed contract.
func NewArbBot(address common.Address, backend bind.ContractBackend) (*ArbBot, error) {
	contract, err := bindArbBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbBot{ArbBotCaller: ArbBotCaller{contract: contract}, ArbBotTransactor: ArbBotTransactor{contract: contract}, ArbBotFilterer: ArbBotFilterer{contract: contract}}, nil
}

// NewArbBotCaller creates a new read-only instance of ArbBot, bound to a specific deployed contract.
func NewArbBotCaller(address common.Address, caller bind.ContractCaller) (*ArbBotCaller, error) {
	contract, err := bindArbBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBotCaller{contract: contract}, nil
}

// NewArbBotTransactor creates a new write-only instance of ArbBot, bound to a specific deployed contract.
func NewArbBotTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbBotTransactor, error) {
	contract, err := bindArbBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbBotTransactor{contract: contract}, nil
}

// NewArbBotFilterer creates a new log filterer instance of ArbBot, bound to a specific deployed contract.
func NewArbBotFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbBotFilterer, error) {
	contract, err := bindArbBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbBotFilterer{contract: contract}, nil
}

// bindArbBot binds a generic wrapper to an already deployed contract.
func bindArbBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbBotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBot *ArbBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbBot.Contract.ArbBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBot *ArbBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBot.Contract.ArbBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBot *ArbBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBot.Contract.ArbBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbBot *ArbBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbBot *ArbBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbBot *ArbBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbBot.Contract.contract.Transact(opts, method, params...)
}

// AllowedRouters is a free data retrieval call binding the contract method 0xc646aee2.
//
// Solidity: function allowedRouters(address ) view returns(bool)
func (_ArbBot *ArbBotCaller) AllowedRouters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ArbBot.contract.Call(opts, &out, "allowedRouters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedRouters is a free data retrieval call binding the contract method 0xc646aee2.
//
// Solidity: function allowedRouters(address ) view returns(bool)
func (_ArbBot *ArbBotSession) AllowedRouters(arg0 common.Address) (bool, error) {
	return _ArbBot.Contract.AllowedRouters(&_ArbBot.CallOpts, arg0)
}

// AllowedRouters is a free data retrieval call binding the contract method 0xc646aee2.
//
// Solidity: function allowedRouters(address ) view returns(bool)
func (_ArbBot *ArbBotCallerSession) AllowedRouters(arg0 common.Address) (bool, error) {
	return _ArbBot.Contract.AllowedRouters(&_ArbBot.CallOpts, arg0)
}

// CheckUniV2 is a free data retrieval call binding the contract method 0x26df5565.
//
// Solidity: function checkUniV2(address _router, address tokenA, address tokenB) view returns(uint256, uint256, uint256)
func (_ArbBot *ArbBotCaller) CheckUniV2(opts *bind.CallOpts, _router common.Address, tokenA common.Address, tokenB common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ArbBot.contract.Call(opts, &out, "checkUniV2", _router, tokenA, tokenB)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CheckUniV2 is a free data retrieval call binding the contract method 0x26df5565.
//
// Solidity: function checkUniV2(address _router, address tokenA, address tokenB) view returns(uint256, uint256, uint256)
func (_ArbBot *ArbBotSession) CheckUniV2(_router common.Address, tokenA common.Address, tokenB common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ArbBot.Contract.CheckUniV2(&_ArbBot.CallOpts, _router, tokenA, tokenB)
}

// CheckUniV2 is a free data retrieval call binding the contract method 0x26df5565.
//
// Solidity: function checkUniV2(address _router, address tokenA, address tokenB) view returns(uint256, uint256, uint256)
func (_ArbBot *ArbBotCallerSession) CheckUniV2(_router common.Address, tokenA common.Address, tokenB common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ArbBot.Contract.CheckUniV2(&_ArbBot.CallOpts, _router, tokenA, tokenB)
}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ArbBot *ArbBotCaller) Executors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ArbBot.contract.Call(opts, &out, "executors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ArbBot *ArbBotSession) Executors(arg0 common.Address) (bool, error) {
	return _ArbBot.Contract.Executors(&_ArbBot.CallOpts, arg0)
}

// Executors is a free data retrieval call binding the contract method 0x9ac2a011.
//
// Solidity: function executors(address ) view returns(bool)
func (_ArbBot *ArbBotCallerSession) Executors(arg0 common.Address) (bool, error) {
	return _ArbBot.Contract.Executors(&_ArbBot.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbBot *ArbBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbBot *ArbBotSession) Owner() (common.Address, error) {
	return _ArbBot.Contract.Owner(&_ArbBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbBot *ArbBotCallerSession) Owner() (common.Address, error) {
	return _ArbBot.Contract.Owner(&_ArbBot.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ArbBot *ArbBotCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbBot.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ArbBot *ArbBotSession) PriceOracle() (common.Address, error) {
	return _ArbBot.Contract.PriceOracle(&_ArbBot.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_ArbBot *ArbBotCallerSession) PriceOracle() (common.Address, error) {
	return _ArbBot.Contract.PriceOracle(&_ArbBot.CallOpts)
}

// AllowExecutor is a paid mutator transaction binding the contract method 0xb1b05f2a.
//
// Solidity: function allowExecutor(address _executor) returns()
func (_ArbBot *ArbBotTransactor) AllowExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "allowExecutor", _executor)
}

// AllowExecutor is a paid mutator transaction binding the contract method 0xb1b05f2a.
//
// Solidity: function allowExecutor(address _executor) returns()
func (_ArbBot *ArbBotSession) AllowExecutor(_executor common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.AllowExecutor(&_ArbBot.TransactOpts, _executor)
}

// AllowExecutor is a paid mutator transaction binding the contract method 0xb1b05f2a.
//
// Solidity: function allowExecutor(address _executor) returns()
func (_ArbBot *ArbBotTransactorSession) AllowExecutor(_executor common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.AllowExecutor(&_ArbBot.TransactOpts, _executor)
}

// AllowRouter is a paid mutator transaction binding the contract method 0xfa1fd0fb.
//
// Solidity: function allowRouter(address router) returns()
func (_ArbBot *ArbBotTransactor) AllowRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "allowRouter", router)
}

// AllowRouter is a paid mutator transaction binding the contract method 0xfa1fd0fb.
//
// Solidity: function allowRouter(address router) returns()
func (_ArbBot *ArbBotSession) AllowRouter(router common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.AllowRouter(&_ArbBot.TransactOpts, router)
}

// AllowRouter is a paid mutator transaction binding the contract method 0xfa1fd0fb.
//
// Solidity: function allowRouter(address router) returns()
func (_ArbBot *ArbBotTransactorSession) AllowRouter(router common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.AllowRouter(&_ArbBot.TransactOpts, router)
}

// Faucet is a paid mutator transaction binding the contract method 0xcdb5afdd.
//
// Solidity: function faucet(address token, address to, uint256 amount) returns()
func (_ArbBot *ArbBotTransactor) Faucet(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "faucet", token, to, amount)
}

// Faucet is a paid mutator transaction binding the contract method 0xcdb5afdd.
//
// Solidity: function faucet(address token, address to, uint256 amount) returns()
func (_ArbBot *ArbBotSession) Faucet(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.Faucet(&_ArbBot.TransactOpts, token, to, amount)
}

// Faucet is a paid mutator transaction binding the contract method 0xcdb5afdd.
//
// Solidity: function faucet(address token, address to, uint256 amount) returns()
func (_ArbBot *ArbBotTransactorSession) Faucet(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.Faucet(&_ArbBot.TransactOpts, token, to, amount)
}

// ForbidExecutor is a paid mutator transaction binding the contract method 0x6e9d5987.
//
// Solidity: function forbidExecutor(address _executor) returns()
func (_ArbBot *ArbBotTransactor) ForbidExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "forbidExecutor", _executor)
}

// ForbidExecutor is a paid mutator transaction binding the contract method 0x6e9d5987.
//
// Solidity: function forbidExecutor(address _executor) returns()
func (_ArbBot *ArbBotSession) ForbidExecutor(_executor common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.ForbidExecutor(&_ArbBot.TransactOpts, _executor)
}

// ForbidExecutor is a paid mutator transaction binding the contract method 0x6e9d5987.
//
// Solidity: function forbidExecutor(address _executor) returns()
func (_ArbBot *ArbBotTransactorSession) ForbidExecutor(_executor common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.ForbidExecutor(&_ArbBot.TransactOpts, _executor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBot *ArbBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBot *ArbBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbBot.Contract.RenounceOwnership(&_ArbBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbBot *ArbBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbBot.Contract.RenounceOwnership(&_ArbBot.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBot *ArbBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBot *ArbBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.TransferOwnership(&_ArbBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbBot *ArbBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.TransferOwnership(&_ArbBot.TransactOpts, newOwner)
}

// UpdatePriceV2 is a paid mutator transaction binding the contract method 0x0cca81f1.
//
// Solidity: function updatePriceV2(address _router, address tokenA, address tokenB) returns()
func (_ArbBot *ArbBotTransactor) UpdatePriceV2(opts *bind.TransactOpts, _router common.Address, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "updatePriceV2", _router, tokenA, tokenB)
}

// UpdatePriceV2 is a paid mutator transaction binding the contract method 0x0cca81f1.
//
// Solidity: function updatePriceV2(address _router, address tokenA, address tokenB) returns()
func (_ArbBot *ArbBotSession) UpdatePriceV2(_router common.Address, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.UpdatePriceV2(&_ArbBot.TransactOpts, _router, tokenA, tokenB)
}

// UpdatePriceV2 is a paid mutator transaction binding the contract method 0x0cca81f1.
//
// Solidity: function updatePriceV2(address _router, address tokenA, address tokenB) returns()
func (_ArbBot *ArbBotTransactorSession) UpdatePriceV2(_router common.Address, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _ArbBot.Contract.UpdatePriceV2(&_ArbBot.TransactOpts, _router, tokenA, tokenB)
}

// V3ethToToken is a paid mutator transaction binding the contract method 0xd0616f97.
//
// Solidity: function v3ethToToken(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotTransactor) V3ethToToken(opts *bind.TransactOpts, v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "v3ethToToken", v3Router, v2Router, token, multiplier)
}

// V3ethToToken is a paid mutator transaction binding the contract method 0xd0616f97.
//
// Solidity: function v3ethToToken(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotSession) V3ethToToken(v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.V3ethToToken(&_ArbBot.TransactOpts, v3Router, v2Router, token, multiplier)
}

// V3ethToToken is a paid mutator transaction binding the contract method 0xd0616f97.
//
// Solidity: function v3ethToToken(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotTransactorSession) V3ethToToken(v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.V3ethToToken(&_ArbBot.TransactOpts, v3Router, v2Router, token, multiplier)
}

// V3tokenToEth is a paid mutator transaction binding the contract method 0x0decdb49.
//
// Solidity: function v3tokenToEth(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotTransactor) V3tokenToEth(opts *bind.TransactOpts, v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.contract.Transact(opts, "v3tokenToEth", v3Router, v2Router, token, multiplier)
}

// V3tokenToEth is a paid mutator transaction binding the contract method 0x0decdb49.
//
// Solidity: function v3tokenToEth(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotSession) V3tokenToEth(v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.V3tokenToEth(&_ArbBot.TransactOpts, v3Router, v2Router, token, multiplier)
}

// V3tokenToEth is a paid mutator transaction binding the contract method 0x0decdb49.
//
// Solidity: function v3tokenToEth(address v3Router, address v2Router, address token, uint256 multiplier) returns()
func (_ArbBot *ArbBotTransactorSession) V3tokenToEth(v3Router common.Address, v2Router common.Address, token common.Address, multiplier *big.Int) (*types.Transaction, error) {
	return _ArbBot.Contract.V3tokenToEth(&_ArbBot.TransactOpts, v3Router, v2Router, token, multiplier)
}

// ArbBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArbBot contract.
type ArbBotOwnershipTransferredIterator struct {
	Event *ArbBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbBotOwnershipTransferred)
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
		it.Event = new(ArbBotOwnershipTransferred)
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
func (it *ArbBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbBotOwnershipTransferred represents a OwnershipTransferred event raised by the ArbBot contract.
type ArbBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbBot *ArbBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbBotOwnershipTransferredIterator{contract: _ArbBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbBot *ArbBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbBotOwnershipTransferred)
				if err := _ArbBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ArbBot *ArbBotFilterer) ParseOwnershipTransferred(log types.Log) (*ArbBotOwnershipTransferred, error) {
	event := new(ArbBotOwnershipTransferred)
	if err := _ArbBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
