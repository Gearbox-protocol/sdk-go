// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceOraclev3

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

// PriceOraclev3MetaData contains all meta data concerning the PriceOraclev3 contract.
var PriceOraclev3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotControllerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnpausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedDoesNotExistException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StalePriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"}],\"name\":\"SetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"}],\"name\":\"SetReservePriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"SetReservePriceFeedStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertFromUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertToUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reserve\",\"type\":\"bool\"}],\"name\":\"getPriceRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeedParams\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"reserve\",\"type\":\"bool\"}],\"name\":\"priceFeedsRaw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"setReservePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setReservePriceFeedStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceOraclev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOraclev3MetaData.ABI instead.
var PriceOraclev3ABI = PriceOraclev3MetaData.ABI

// PriceOraclev3 is an auto generated Go binding around an Ethereum contract.
type PriceOraclev3 struct {
	PriceOraclev3Caller     // Read-only binding to the contract
	PriceOraclev3Transactor // Write-only binding to the contract
	PriceOraclev3Filterer   // Log filterer for contract events
}

// PriceOraclev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOraclev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOraclev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOraclev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOraclev3Session struct {
	Contract     *PriceOraclev3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOraclev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOraclev3CallerSession struct {
	Contract *PriceOraclev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriceOraclev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOraclev3TransactorSession struct {
	Contract     *PriceOraclev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriceOraclev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOraclev3Raw struct {
	Contract *PriceOraclev3 // Generic contract binding to access the raw methods on
}

// PriceOraclev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOraclev3CallerRaw struct {
	Contract *PriceOraclev3Caller // Generic read-only contract binding to access the raw methods on
}

// PriceOraclev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOraclev3TransactorRaw struct {
	Contract *PriceOraclev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOraclev3 creates a new instance of PriceOraclev3, bound to a specific deployed contract.
func NewPriceOraclev3(address common.Address, backend bind.ContractBackend) (*PriceOraclev3, error) {
	contract, err := bindPriceOraclev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3{PriceOraclev3Caller: PriceOraclev3Caller{contract: contract}, PriceOraclev3Transactor: PriceOraclev3Transactor{contract: contract}, PriceOraclev3Filterer: PriceOraclev3Filterer{contract: contract}}, nil
}

// NewPriceOraclev3Caller creates a new read-only instance of PriceOraclev3, bound to a specific deployed contract.
func NewPriceOraclev3Caller(address common.Address, caller bind.ContractCaller) (*PriceOraclev3Caller, error) {
	contract, err := bindPriceOraclev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3Caller{contract: contract}, nil
}

// NewPriceOraclev3Transactor creates a new write-only instance of PriceOraclev3, bound to a specific deployed contract.
func NewPriceOraclev3Transactor(address common.Address, transactor bind.ContractTransactor) (*PriceOraclev3Transactor, error) {
	contract, err := bindPriceOraclev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3Transactor{contract: contract}, nil
}

// NewPriceOraclev3Filterer creates a new log filterer instance of PriceOraclev3, bound to a specific deployed contract.
func NewPriceOraclev3Filterer(address common.Address, filterer bind.ContractFilterer) (*PriceOraclev3Filterer, error) {
	contract, err := bindPriceOraclev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3Filterer{contract: contract}, nil
}

// bindPriceOraclev3 binds a generic wrapper to an already deployed contract.
func bindPriceOraclev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceOraclev3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOraclev3 *PriceOraclev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOraclev3.Contract.PriceOraclev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOraclev3 *PriceOraclev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.PriceOraclev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOraclev3 *PriceOraclev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.PriceOraclev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOraclev3 *PriceOraclev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOraclev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOraclev3 *PriceOraclev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOraclev3 *PriceOraclev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PriceOraclev3 *PriceOraclev3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PriceOraclev3 *PriceOraclev3Session) Acl() (common.Address, error) {
	return _PriceOraclev3.Contract.Acl(&_PriceOraclev3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PriceOraclev3 *PriceOraclev3CallerSession) Acl() (common.Address, error) {
	return _PriceOraclev3.Contract.Acl(&_PriceOraclev3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PriceOraclev3 *PriceOraclev3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PriceOraclev3 *PriceOraclev3Session) Controller() (common.Address, error) {
	return _PriceOraclev3.Contract.Controller(&_PriceOraclev3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PriceOraclev3 *PriceOraclev3CallerSession) Controller() (common.Address, error) {
	return _PriceOraclev3.Contract.Controller(&_PriceOraclev3.CallOpts)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Caller) Convert(opts *bind.CallOpts, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "convert", amount, tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Session) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.Convert(&_PriceOraclev3.CallOpts, amount, tokenFrom, tokenTo)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3CallerSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.Convert(&_PriceOraclev3.CallOpts, amount, tokenFrom, tokenTo)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Caller) ConvertFromUSD(opts *bind.CallOpts, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "convertFromUSD", amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Session) ConvertFromUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.ConvertFromUSD(&_PriceOraclev3.CallOpts, amount, token)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3CallerSession) ConvertFromUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.ConvertFromUSD(&_PriceOraclev3.CallOpts, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Caller) ConvertToUSD(opts *bind.CallOpts, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "convertToUSD", amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Session) ConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.ConvertToUSD(&_PriceOraclev3.CallOpts, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3CallerSession) ConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.ConvertToUSD(&_PriceOraclev3.CallOpts, amount, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3Caller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3Session) GetPrice(token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.GetPrice(&_PriceOraclev3.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3CallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _PriceOraclev3.Contract.GetPrice(&_PriceOraclev3.CallOpts, token)
}

// GetPriceRaw is a free data retrieval call binding the contract method 0x8f8a8aba.
//
// Solidity: function getPriceRaw(address token, bool reserve) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3Caller) GetPriceRaw(opts *bind.CallOpts, token common.Address, reserve bool) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "getPriceRaw", token, reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceRaw is a free data retrieval call binding the contract method 0x8f8a8aba.
//
// Solidity: function getPriceRaw(address token, bool reserve) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3Session) GetPriceRaw(token common.Address, reserve bool) (*big.Int, error) {
	return _PriceOraclev3.Contract.GetPriceRaw(&_PriceOraclev3.CallOpts, token, reserve)
}

// GetPriceRaw is a free data retrieval call binding the contract method 0x8f8a8aba.
//
// Solidity: function getPriceRaw(address token, bool reserve) view returns(uint256 price)
func (_PriceOraclev3 *PriceOraclev3CallerSession) GetPriceRaw(token common.Address, reserve bool) (*big.Int, error) {
	return _PriceOraclev3.Contract.GetPriceRaw(&_PriceOraclev3.CallOpts, token, reserve)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev3 *PriceOraclev3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev3 *PriceOraclev3Session) Paused() (bool, error) {
	return _PriceOraclev3.Contract.Paused(&_PriceOraclev3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev3 *PriceOraclev3CallerSession) Paused() (bool, error) {
	return _PriceOraclev3.Contract.Paused(&_PriceOraclev3.CallOpts)
}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns(address priceFeed, uint32 stalenessPeriod, bool skipCheck, uint8 decimals)
func (_PriceOraclev3 *PriceOraclev3Caller) PriceFeedParams(opts *bind.CallOpts, token common.Address) (struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Decimals        uint8
}, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "priceFeedParams", token)

	outstruct := new(struct {
		PriceFeed       common.Address
		StalenessPeriod uint32
		SkipCheck       bool
		Decimals        uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceFeed = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StalenessPeriod = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.SkipCheck = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns(address priceFeed, uint32 stalenessPeriod, bool skipCheck, uint8 decimals)
func (_PriceOraclev3 *PriceOraclev3Session) PriceFeedParams(token common.Address) (struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Decimals        uint8
}, error) {
	return _PriceOraclev3.Contract.PriceFeedParams(&_PriceOraclev3.CallOpts, token)
}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns(address priceFeed, uint32 stalenessPeriod, bool skipCheck, uint8 decimals)
func (_PriceOraclev3 *PriceOraclev3CallerSession) PriceFeedParams(token common.Address) (struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Decimals        uint8
}, error) {
	return _PriceOraclev3.Contract.PriceFeedParams(&_PriceOraclev3.CallOpts, token)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3Caller) PriceFeeds(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "priceFeeds", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3Session) PriceFeeds(token common.Address) (common.Address, error) {
	return _PriceOraclev3.Contract.PriceFeeds(&_PriceOraclev3.CallOpts, token)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3CallerSession) PriceFeeds(token common.Address) (common.Address, error) {
	return _PriceOraclev3.Contract.PriceFeeds(&_PriceOraclev3.CallOpts, token)
}

// PriceFeedsRaw is a free data retrieval call binding the contract method 0xff299845.
//
// Solidity: function priceFeedsRaw(address token, bool reserve) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3Caller) PriceFeedsRaw(opts *bind.CallOpts, token common.Address, reserve bool) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "priceFeedsRaw", token, reserve)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeedsRaw is a free data retrieval call binding the contract method 0xff299845.
//
// Solidity: function priceFeedsRaw(address token, bool reserve) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3Session) PriceFeedsRaw(token common.Address, reserve bool) (common.Address, error) {
	return _PriceOraclev3.Contract.PriceFeedsRaw(&_PriceOraclev3.CallOpts, token, reserve)
}

// PriceFeedsRaw is a free data retrieval call binding the contract method 0xff299845.
//
// Solidity: function priceFeedsRaw(address token, bool reserve) view returns(address priceFeed)
func (_PriceOraclev3 *PriceOraclev3CallerSession) PriceFeedsRaw(token common.Address, reserve bool) (common.Address, error) {
	return _PriceOraclev3.Contract.PriceFeedsRaw(&_PriceOraclev3.CallOpts, token, reserve)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3Session) Version() (*big.Int, error) {
	return _PriceOraclev3.Contract.Version(&_PriceOraclev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev3 *PriceOraclev3CallerSession) Version() (*big.Int, error) {
	return _PriceOraclev3.Contract.Version(&_PriceOraclev3.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev3 *PriceOraclev3Session) Pause() (*types.Transaction, error) {
	return _PriceOraclev3.Contract.Pause(&_PriceOraclev3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) Pause() (*types.Transaction, error) {
	return _PriceOraclev3.Contract.Pause(&_PriceOraclev3.TransactOpts)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PriceOraclev3 *PriceOraclev3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetController(&_PriceOraclev3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetController(&_PriceOraclev3.TransactOpts, newController)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) SetPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "setPriceFeed", token, priceFeed, stalenessPeriod)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3Session) SetPriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetPriceFeed(&_PriceOraclev3.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) SetPriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetPriceFeed(&_PriceOraclev3.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) SetReservePriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "setReservePriceFeed", token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3Session) SetReservePriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetReservePriceFeed(&_PriceOraclev3.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) SetReservePriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetReservePriceFeed(&_PriceOraclev3.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeedStatus is a paid mutator transaction binding the contract method 0x1b6dfcc0.
//
// Solidity: function setReservePriceFeedStatus(address token, bool active) returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) SetReservePriceFeedStatus(opts *bind.TransactOpts, token common.Address, active bool) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "setReservePriceFeedStatus", token, active)
}

// SetReservePriceFeedStatus is a paid mutator transaction binding the contract method 0x1b6dfcc0.
//
// Solidity: function setReservePriceFeedStatus(address token, bool active) returns()
func (_PriceOraclev3 *PriceOraclev3Session) SetReservePriceFeedStatus(token common.Address, active bool) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetReservePriceFeedStatus(&_PriceOraclev3.TransactOpts, token, active)
}

// SetReservePriceFeedStatus is a paid mutator transaction binding the contract method 0x1b6dfcc0.
//
// Solidity: function setReservePriceFeedStatus(address token, bool active) returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) SetReservePriceFeedStatus(token common.Address, active bool) (*types.Transaction, error) {
	return _PriceOraclev3.Contract.SetReservePriceFeedStatus(&_PriceOraclev3.TransactOpts, token, active)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev3 *PriceOraclev3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev3 *PriceOraclev3Session) Unpause() (*types.Transaction, error) {
	return _PriceOraclev3.Contract.Unpause(&_PriceOraclev3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev3 *PriceOraclev3TransactorSession) Unpause() (*types.Transaction, error) {
	return _PriceOraclev3.Contract.Unpause(&_PriceOraclev3.TransactOpts)
}

// PriceOraclev3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the PriceOraclev3 contract.
type PriceOraclev3NewControllerIterator struct {
	Event *PriceOraclev3NewController // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3NewController)
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
		it.Event = new(PriceOraclev3NewController)
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
func (it *PriceOraclev3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3NewController represents a NewController event raised by the PriceOraclev3 contract.
type PriceOraclev3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*PriceOraclev3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3NewControllerIterator{contract: _PriceOraclev3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *PriceOraclev3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3NewController)
				if err := _PriceOraclev3.contract.UnpackLog(event, "NewController", log); err != nil {
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

// ParseNewController is a log parse operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParseNewController(log types.Log) (*PriceOraclev3NewController, error) {
	event := new(PriceOraclev3NewController)
	if err := _PriceOraclev3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceOraclev3 contract.
type PriceOraclev3PausedIterator struct {
	Event *PriceOraclev3Paused // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3Paused)
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
		it.Event = new(PriceOraclev3Paused)
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
func (it *PriceOraclev3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3Paused represents a Paused event raised by the PriceOraclev3 contract.
type PriceOraclev3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterPaused(opts *bind.FilterOpts) (*PriceOraclev3PausedIterator, error) {

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3PausedIterator{contract: _PriceOraclev3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceOraclev3Paused) (event.Subscription, error) {

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3Paused)
				if err := _PriceOraclev3.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParsePaused(log types.Log) (*PriceOraclev3Paused, error) {
	event := new(PriceOraclev3Paused)
	if err := _PriceOraclev3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev3SetPriceFeedIterator is returned from FilterSetPriceFeed and is used to iterate over the raw logs and unpacked data for SetPriceFeed events raised by the PriceOraclev3 contract.
type PriceOraclev3SetPriceFeedIterator struct {
	Event *PriceOraclev3SetPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3SetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3SetPriceFeed)
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
		it.Event = new(PriceOraclev3SetPriceFeed)
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
func (it *PriceOraclev3SetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3SetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3SetPriceFeed represents a SetPriceFeed event raised by the PriceOraclev3 contract.
type PriceOraclev3SetPriceFeed struct {
	Token           common.Address
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeed is a free log retrieval operation binding the contract event 0x54bafe4dd078667ca4498828ef773452f21b7c2b2f71c734bd085cab2a0f4f4b.
//
// Solidity: event SetPriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterSetPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceOraclev3SetPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "SetPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3SetPriceFeedIterator{contract: _PriceOraclev3.contract, event: "SetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeed is a free log subscription operation binding the contract event 0x54bafe4dd078667ca4498828ef773452f21b7c2b2f71c734bd085cab2a0f4f4b.
//
// Solidity: event SetPriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchSetPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceOraclev3SetPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "SetPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3SetPriceFeed)
				if err := _PriceOraclev3.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
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

// ParseSetPriceFeed is a log parse operation binding the contract event 0x54bafe4dd078667ca4498828ef773452f21b7c2b2f71c734bd085cab2a0f4f4b.
//
// Solidity: event SetPriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParseSetPriceFeed(log types.Log) (*PriceOraclev3SetPriceFeed, error) {
	event := new(PriceOraclev3SetPriceFeed)
	if err := _PriceOraclev3.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev3SetReservePriceFeedIterator is returned from FilterSetReservePriceFeed and is used to iterate over the raw logs and unpacked data for SetReservePriceFeed events raised by the PriceOraclev3 contract.
type PriceOraclev3SetReservePriceFeedIterator struct {
	Event *PriceOraclev3SetReservePriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3SetReservePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3SetReservePriceFeed)
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
		it.Event = new(PriceOraclev3SetReservePriceFeed)
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
func (it *PriceOraclev3SetReservePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3SetReservePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3SetReservePriceFeed represents a SetReservePriceFeed event raised by the PriceOraclev3 contract.
type PriceOraclev3SetReservePriceFeed struct {
	Token           common.Address
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePriceFeed is a free log retrieval operation binding the contract event 0x0622a3fd57e320160fa23aaad98e12139a7922016875937c646e76237005c01c.
//
// Solidity: event SetReservePriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterSetReservePriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceOraclev3SetReservePriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "SetReservePriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3SetReservePriceFeedIterator{contract: _PriceOraclev3.contract, event: "SetReservePriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetReservePriceFeed is a free log subscription operation binding the contract event 0x0622a3fd57e320160fa23aaad98e12139a7922016875937c646e76237005c01c.
//
// Solidity: event SetReservePriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchSetReservePriceFeed(opts *bind.WatchOpts, sink chan<- *PriceOraclev3SetReservePriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "SetReservePriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3SetReservePriceFeed)
				if err := _PriceOraclev3.contract.UnpackLog(event, "SetReservePriceFeed", log); err != nil {
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

// ParseSetReservePriceFeed is a log parse operation binding the contract event 0x0622a3fd57e320160fa23aaad98e12139a7922016875937c646e76237005c01c.
//
// Solidity: event SetReservePriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParseSetReservePriceFeed(log types.Log) (*PriceOraclev3SetReservePriceFeed, error) {
	event := new(PriceOraclev3SetReservePriceFeed)
	if err := _PriceOraclev3.contract.UnpackLog(event, "SetReservePriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev3SetReservePriceFeedStatusIterator is returned from FilterSetReservePriceFeedStatus and is used to iterate over the raw logs and unpacked data for SetReservePriceFeedStatus events raised by the PriceOraclev3 contract.
type PriceOraclev3SetReservePriceFeedStatusIterator struct {
	Event *PriceOraclev3SetReservePriceFeedStatus // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3SetReservePriceFeedStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3SetReservePriceFeedStatus)
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
		it.Event = new(PriceOraclev3SetReservePriceFeedStatus)
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
func (it *PriceOraclev3SetReservePriceFeedStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3SetReservePriceFeedStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3SetReservePriceFeedStatus represents a SetReservePriceFeedStatus event raised by the PriceOraclev3 contract.
type PriceOraclev3SetReservePriceFeedStatus struct {
	Token  common.Address
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetReservePriceFeedStatus is a free log retrieval operation binding the contract event 0x8065e881e49e3c3c04dadba3c10e1abdd738d4aa4a7807b49dc5d4600ec62885.
//
// Solidity: event SetReservePriceFeedStatus(address indexed token, bool active)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterSetReservePriceFeedStatus(opts *bind.FilterOpts, token []common.Address) (*PriceOraclev3SetReservePriceFeedStatusIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "SetReservePriceFeedStatus", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3SetReservePriceFeedStatusIterator{contract: _PriceOraclev3.contract, event: "SetReservePriceFeedStatus", logs: logs, sub: sub}, nil
}

// WatchSetReservePriceFeedStatus is a free log subscription operation binding the contract event 0x8065e881e49e3c3c04dadba3c10e1abdd738d4aa4a7807b49dc5d4600ec62885.
//
// Solidity: event SetReservePriceFeedStatus(address indexed token, bool active)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchSetReservePriceFeedStatus(opts *bind.WatchOpts, sink chan<- *PriceOraclev3SetReservePriceFeedStatus, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "SetReservePriceFeedStatus", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3SetReservePriceFeedStatus)
				if err := _PriceOraclev3.contract.UnpackLog(event, "SetReservePriceFeedStatus", log); err != nil {
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

// ParseSetReservePriceFeedStatus is a log parse operation binding the contract event 0x8065e881e49e3c3c04dadba3c10e1abdd738d4aa4a7807b49dc5d4600ec62885.
//
// Solidity: event SetReservePriceFeedStatus(address indexed token, bool active)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParseSetReservePriceFeedStatus(log types.Log) (*PriceOraclev3SetReservePriceFeedStatus, error) {
	event := new(PriceOraclev3SetReservePriceFeedStatus)
	if err := _PriceOraclev3.contract.UnpackLog(event, "SetReservePriceFeedStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceOraclev3 contract.
type PriceOraclev3UnpausedIterator struct {
	Event *PriceOraclev3Unpaused // Event containing the contract specifics and raw log

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
func (it *PriceOraclev3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev3Unpaused)
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
		it.Event = new(PriceOraclev3Unpaused)
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
func (it *PriceOraclev3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev3Unpaused represents a Unpaused event raised by the PriceOraclev3 contract.
type PriceOraclev3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceOraclev3UnpausedIterator, error) {

	logs, sub, err := _PriceOraclev3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceOraclev3UnpausedIterator{contract: _PriceOraclev3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceOraclev3Unpaused) (event.Subscription, error) {

	logs, sub, err := _PriceOraclev3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev3Unpaused)
				if err := _PriceOraclev3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOraclev3 *PriceOraclev3Filterer) ParseUnpaused(log types.Log) (*PriceOraclev3Unpaused, error) {
	event := new(PriceOraclev3Unpaused)
	if err := _PriceOraclev3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
