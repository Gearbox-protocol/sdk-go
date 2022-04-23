// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tradingBot

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

// TradingBotABI is the input ABI used to generate the binding from.
const TradingBotABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_syncer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"approveToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncer\",\"outputs\":[{\"internalType\":\"contractSyncer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TradingBot is an auto generated Go binding around an Ethereum contract.
type TradingBot struct {
	TradingBotCaller     // Read-only binding to the contract
	TradingBotTransactor // Write-only binding to the contract
	TradingBotFilterer   // Log filterer for contract events
}

// TradingBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingBotSession struct {
	Contract     *TradingBot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingBotCallerSession struct {
	Contract *TradingBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradingBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingBotTransactorSession struct {
	Contract     *TradingBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradingBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingBotRaw struct {
	Contract *TradingBot // Generic contract binding to access the raw methods on
}

// TradingBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingBotCallerRaw struct {
	Contract *TradingBotCaller // Generic read-only contract binding to access the raw methods on
}

// TradingBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingBotTransactorRaw struct {
	Contract *TradingBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingBot creates a new instance of TradingBot, bound to a specific deployed contract.
func NewTradingBot(address common.Address, backend bind.ContractBackend) (*TradingBot, error) {
	contract, err := bindTradingBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradingBot{TradingBotCaller: TradingBotCaller{contract: contract}, TradingBotTransactor: TradingBotTransactor{contract: contract}, TradingBotFilterer: TradingBotFilterer{contract: contract}}, nil
}

// NewTradingBotCaller creates a new read-only instance of TradingBot, bound to a specific deployed contract.
func NewTradingBotCaller(address common.Address, caller bind.ContractCaller) (*TradingBotCaller, error) {
	contract, err := bindTradingBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingBotCaller{contract: contract}, nil
}

// NewTradingBotTransactor creates a new write-only instance of TradingBot, bound to a specific deployed contract.
func NewTradingBotTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingBotTransactor, error) {
	contract, err := bindTradingBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingBotTransactor{contract: contract}, nil
}

// NewTradingBotFilterer creates a new log filterer instance of TradingBot, bound to a specific deployed contract.
func NewTradingBotFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingBotFilterer, error) {
	contract, err := bindTradingBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingBotFilterer{contract: contract}, nil
}

// bindTradingBot binds a generic wrapper to an already deployed contract.
func bindTradingBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingBotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingBot *TradingBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingBot.Contract.TradingBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingBot *TradingBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBot.Contract.TradingBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingBot *TradingBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingBot.Contract.TradingBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingBot *TradingBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingBot *TradingBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingBot *TradingBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingBot.Contract.contract.Transact(opts, method, params...)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_TradingBot *TradingBotCaller) Syncer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingBot.contract.Call(opts, &out, "syncer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_TradingBot *TradingBotSession) Syncer() (common.Address, error) {
	return _TradingBot.Contract.Syncer(&_TradingBot.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_TradingBot *TradingBotCallerSession) Syncer() (common.Address, error) {
	return _TradingBot.Contract.Syncer(&_TradingBot.CallOpts)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address targetContract) returns()
func (_TradingBot *TradingBotTransactor) ApproveToken(opts *bind.TransactOpts, token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _TradingBot.contract.Transact(opts, "approveToken", token, targetContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address targetContract) returns()
func (_TradingBot *TradingBotSession) ApproveToken(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _TradingBot.Contract.ApproveToken(&_TradingBot.TransactOpts, token, targetContract)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x03105b04.
//
// Solidity: function approveToken(address token, address targetContract) returns()
func (_TradingBot *TradingBotTransactorSession) ApproveToken(token common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _TradingBot.Contract.ApproveToken(&_TradingBot.TransactOpts, token, targetContract)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_TradingBot *TradingBotTransactor) Execute(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _TradingBot.contract.Transact(opts, "execute", destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_TradingBot *TradingBotSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _TradingBot.Contract.Execute(&_TradingBot.TransactOpts, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address destination, bytes data) returns(bytes)
func (_TradingBot *TradingBotTransactorSession) Execute(destination common.Address, data []byte) (*types.Transaction, error) {
	return _TradingBot.Contract.Execute(&_TradingBot.TransactOpts, destination, data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_TradingBot *TradingBotTransactor) SafeTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradingBot.contract.Transact(opts, "safeTransfer", token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_TradingBot *TradingBotSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradingBot.Contract.SafeTransfer(&_TradingBot.TransactOpts, token, to, amount)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0xd1660f99.
//
// Solidity: function safeTransfer(address token, address to, uint256 amount) returns()
func (_TradingBot *TradingBotTransactorSession) SafeTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TradingBot.Contract.SafeTransfer(&_TradingBot.TransactOpts, token, to, amount)
}
