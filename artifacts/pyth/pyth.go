// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pyth

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

// PythMetaData contains all meta data concerning the Pyth contract.
var PythMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_priceFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_pyth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxConfToPriceRatio\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"descriptionTicker\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ConfToPriceRatioTooHighException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectExpectedPublishTimestampException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceDecimalsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceTimestampTooFarAheadException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceTimestampTooFarBehindException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxConfToPriceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeedId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pyth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serialize\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipPriceCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PythABI is the input ABI used to generate the binding from.
// Deprecated: Use PythMetaData.ABI instead.
var PythABI = PythMetaData.ABI

// Pyth is an auto generated Go binding around an Ethereum contract.
type Pyth struct {
	PythCaller     // Read-only binding to the contract
	PythTransactor // Write-only binding to the contract
	PythFilterer   // Log filterer for contract events
}

// PythCaller is an auto generated read-only Go binding around an Ethereum contract.
type PythCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PythTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PythFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PythSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PythSession struct {
	Contract     *Pyth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PythCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PythCallerSession struct {
	Contract *PythCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PythTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PythTransactorSession struct {
	Contract     *PythTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PythRaw is an auto generated low-level Go binding around an Ethereum contract.
type PythRaw struct {
	Contract *Pyth // Generic contract binding to access the raw methods on
}

// PythCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PythCallerRaw struct {
	Contract *PythCaller // Generic read-only contract binding to access the raw methods on
}

// PythTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PythTransactorRaw struct {
	Contract *PythTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPyth creates a new instance of Pyth, bound to a specific deployed contract.
func NewPyth(address common.Address, backend bind.ContractBackend) (*Pyth, error) {
	contract, err := bindPyth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pyth{PythCaller: PythCaller{contract: contract}, PythTransactor: PythTransactor{contract: contract}, PythFilterer: PythFilterer{contract: contract}}, nil
}

// NewPythCaller creates a new read-only instance of Pyth, bound to a specific deployed contract.
func NewPythCaller(address common.Address, caller bind.ContractCaller) (*PythCaller, error) {
	contract, err := bindPyth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PythCaller{contract: contract}, nil
}

// NewPythTransactor creates a new write-only instance of Pyth, bound to a specific deployed contract.
func NewPythTransactor(address common.Address, transactor bind.ContractTransactor) (*PythTransactor, error) {
	contract, err := bindPyth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PythTransactor{contract: contract}, nil
}

// NewPythFilterer creates a new log filterer instance of Pyth, bound to a specific deployed contract.
func NewPythFilterer(address common.Address, filterer bind.ContractFilterer) (*PythFilterer, error) {
	contract, err := bindPyth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PythFilterer{contract: contract}, nil
}

// bindPyth binds a generic wrapper to an already deployed contract.
func bindPyth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PythMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pyth *PythRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pyth.Contract.PythCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pyth *PythRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pyth.Contract.PythTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pyth *PythRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pyth.Contract.PythTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pyth *PythCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pyth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pyth *PythTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pyth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pyth *PythTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pyth.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Pyth *PythCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Pyth *PythSession) ContractType() ([32]byte, error) {
	return _Pyth.Contract.ContractType(&_Pyth.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Pyth *PythCallerSession) ContractType() ([32]byte, error) {
	return _Pyth.Contract.ContractType(&_Pyth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pyth *PythCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pyth *PythSession) Decimals() (uint8, error) {
	return _Pyth.Contract.Decimals(&_Pyth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pyth *PythCallerSession) Decimals() (uint8, error) {
	return _Pyth.Contract.Decimals(&_Pyth.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pyth *PythCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pyth *PythSession) Description() (string, error) {
	return _Pyth.Contract.Description(&_Pyth.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Pyth *PythCallerSession) Description() (string, error) {
	return _Pyth.Contract.Description(&_Pyth.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Pyth *PythCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Pyth *PythSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pyth.Contract.LatestRoundData(&_Pyth.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Pyth *PythCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Pyth.Contract.LatestRoundData(&_Pyth.CallOpts)
}

// MaxConfToPriceRatio is a free data retrieval call binding the contract method 0x81b9dd0a.
//
// Solidity: function maxConfToPriceRatio() view returns(uint256)
func (_Pyth *PythCaller) MaxConfToPriceRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "maxConfToPriceRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxConfToPriceRatio is a free data retrieval call binding the contract method 0x81b9dd0a.
//
// Solidity: function maxConfToPriceRatio() view returns(uint256)
func (_Pyth *PythSession) MaxConfToPriceRatio() (*big.Int, error) {
	return _Pyth.Contract.MaxConfToPriceRatio(&_Pyth.CallOpts)
}

// MaxConfToPriceRatio is a free data retrieval call binding the contract method 0x81b9dd0a.
//
// Solidity: function maxConfToPriceRatio() view returns(uint256)
func (_Pyth *PythCallerSession) MaxConfToPriceRatio() (*big.Int, error) {
	return _Pyth.Contract.MaxConfToPriceRatio(&_Pyth.CallOpts)
}

// PriceFeedId is a free data retrieval call binding the contract method 0x1999bb9e.
//
// Solidity: function priceFeedId() view returns(bytes32)
func (_Pyth *PythCaller) PriceFeedId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "priceFeedId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceFeedId is a free data retrieval call binding the contract method 0x1999bb9e.
//
// Solidity: function priceFeedId() view returns(bytes32)
func (_Pyth *PythSession) PriceFeedId() ([32]byte, error) {
	return _Pyth.Contract.PriceFeedId(&_Pyth.CallOpts)
}

// PriceFeedId is a free data retrieval call binding the contract method 0x1999bb9e.
//
// Solidity: function priceFeedId() view returns(bytes32)
func (_Pyth *PythCallerSession) PriceFeedId() ([32]byte, error) {
	return _Pyth.Contract.PriceFeedId(&_Pyth.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Pyth *PythCaller) Pyth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "pyth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Pyth *PythSession) Pyth() (common.Address, error) {
	return _Pyth.Contract.Pyth(&_Pyth.CallOpts)
}

// Pyth is a free data retrieval call binding the contract method 0xf98d06f0.
//
// Solidity: function pyth() view returns(address)
func (_Pyth *PythCallerSession) Pyth() (common.Address, error) {
	return _Pyth.Contract.Pyth(&_Pyth.CallOpts)
}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_Pyth *PythCaller) Serialize(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "serialize")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_Pyth *PythSession) Serialize() ([]byte, error) {
	return _Pyth.Contract.Serialize(&_Pyth.CallOpts)
}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_Pyth *PythCallerSession) Serialize() ([]byte, error) {
	return _Pyth.Contract.Serialize(&_Pyth.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Pyth *PythCaller) SkipPriceCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "skipPriceCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Pyth *PythSession) SkipPriceCheck() (bool, error) {
	return _Pyth.Contract.SkipPriceCheck(&_Pyth.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Pyth *PythCallerSession) SkipPriceCheck() (bool, error) {
	return _Pyth.Contract.SkipPriceCheck(&_Pyth.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pyth *PythCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pyth *PythSession) Token() (common.Address, error) {
	return _Pyth.Contract.Token(&_Pyth.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pyth *PythCallerSession) Token() (common.Address, error) {
	return _Pyth.Contract.Token(&_Pyth.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Pyth *PythCaller) Updatable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "updatable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Pyth *PythSession) Updatable() (bool, error) {
	return _Pyth.Contract.Updatable(&_Pyth.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Pyth *PythCallerSession) Updatable() (bool, error) {
	return _Pyth.Contract.Updatable(&_Pyth.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Pyth *PythCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pyth.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Pyth *PythSession) Version() (*big.Int, error) {
	return _Pyth.Contract.Version(&_Pyth.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Pyth *PythCallerSession) Version() (*big.Int, error) {
	return _Pyth.Contract.Version(&_Pyth.CallOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Pyth *PythTransactor) UpdatePrice(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Pyth.contract.Transact(opts, "updatePrice", data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Pyth *PythSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _Pyth.Contract.UpdatePrice(&_Pyth.TransactOpts, data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Pyth *PythTransactorSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _Pyth.Contract.UpdatePrice(&_Pyth.TransactOpts, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pyth *PythTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pyth.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pyth *PythSession) Receive() (*types.Transaction, error) {
	return _Pyth.Contract.Receive(&_Pyth.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pyth *PythTransactorSession) Receive() (*types.Transaction, error) {
	return _Pyth.Contract.Receive(&_Pyth.TransactOpts)
}

// PythUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the Pyth contract.
type PythUpdatePriceIterator struct {
	Event *PythUpdatePrice // Event containing the contract specifics and raw log

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
func (it *PythUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PythUpdatePrice)
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
		it.Event = new(PythUpdatePrice)
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
func (it *PythUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PythUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PythUpdatePrice represents a UpdatePrice event raised by the Pyth contract.
type PythUpdatePrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Pyth *PythFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*PythUpdatePriceIterator, error) {

	logs, sub, err := _Pyth.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &PythUpdatePriceIterator{contract: _Pyth.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Pyth *PythFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *PythUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _Pyth.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PythUpdatePrice)
				if err := _Pyth.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Pyth *PythFilterer) ParseUpdatePrice(log types.Log) (*PythUpdatePrice, error) {
	event := new(PythUpdatePrice)
	if err := _Pyth.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
