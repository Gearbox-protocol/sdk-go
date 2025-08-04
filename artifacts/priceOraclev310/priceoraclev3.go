// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceoraclev310

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

// PriceFeedParams is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedParams struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	TokenDecimals   uint8
}

// Priceoraclev310MetaData contains all meta data concerning the Priceoraclev310 contract.
var Priceoraclev310MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_acl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedDoesNotExistException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StalePriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"}],\"name\":\"SetPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"}],\"name\":\"SetReservePriceFeed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertFromUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertToUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getSafePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeedParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structPriceFeedParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reservePriceFeedParams\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structPriceFeedParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reservePriceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeConvertToUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"setReservePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Priceoraclev310ABI is the input ABI used to generate the binding from.
// Deprecated: Use Priceoraclev310MetaData.ABI instead.
var Priceoraclev310ABI = Priceoraclev310MetaData.ABI

// Priceoraclev310 is an auto generated Go binding around an Ethereum contract.
type Priceoraclev310 struct {
	Priceoraclev310Caller     // Read-only binding to the contract
	Priceoraclev310Transactor // Write-only binding to the contract
	Priceoraclev310Filterer   // Log filterer for contract events
}

// Priceoraclev310Caller is an auto generated read-only Go binding around an Ethereum contract.
type Priceoraclev310Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Priceoraclev310Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Priceoraclev310Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Priceoraclev310Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Priceoraclev310Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Priceoraclev310Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Priceoraclev310Session struct {
	Contract     *Priceoraclev310  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Priceoraclev310CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Priceoraclev310CallerSession struct {
	Contract *Priceoraclev310Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Priceoraclev310TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Priceoraclev310TransactorSession struct {
	Contract     *Priceoraclev310Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Priceoraclev310Raw is an auto generated low-level Go binding around an Ethereum contract.
type Priceoraclev310Raw struct {
	Contract *Priceoraclev310 // Generic contract binding to access the raw methods on
}

// Priceoraclev310CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Priceoraclev310CallerRaw struct {
	Contract *Priceoraclev310Caller // Generic read-only contract binding to access the raw methods on
}

// Priceoraclev310TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Priceoraclev310TransactorRaw struct {
	Contract *Priceoraclev310Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceoraclev310 creates a new instance of Priceoraclev310, bound to a specific deployed contract.
func NewPriceoraclev310(address common.Address, backend bind.ContractBackend) (*Priceoraclev310, error) {
	contract, err := bindPriceoraclev310(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310{Priceoraclev310Caller: Priceoraclev310Caller{contract: contract}, Priceoraclev310Transactor: Priceoraclev310Transactor{contract: contract}, Priceoraclev310Filterer: Priceoraclev310Filterer{contract: contract}}, nil
}

// NewPriceoraclev310Caller creates a new read-only instance of Priceoraclev310, bound to a specific deployed contract.
func NewPriceoraclev310Caller(address common.Address, caller bind.ContractCaller) (*Priceoraclev310Caller, error) {
	contract, err := bindPriceoraclev310(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310Caller{contract: contract}, nil
}

// NewPriceoraclev310Transactor creates a new write-only instance of Priceoraclev310, bound to a specific deployed contract.
func NewPriceoraclev310Transactor(address common.Address, transactor bind.ContractTransactor) (*Priceoraclev310Transactor, error) {
	contract, err := bindPriceoraclev310(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310Transactor{contract: contract}, nil
}

// NewPriceoraclev310Filterer creates a new log filterer instance of Priceoraclev310, bound to a specific deployed contract.
func NewPriceoraclev310Filterer(address common.Address, filterer bind.ContractFilterer) (*Priceoraclev310Filterer, error) {
	contract, err := bindPriceoraclev310(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310Filterer{contract: contract}, nil
}

// bindPriceoraclev310 binds a generic wrapper to an already deployed contract.
func bindPriceoraclev310(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Priceoraclev310MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Priceoraclev310 *Priceoraclev310Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Priceoraclev310.Contract.Priceoraclev310Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Priceoraclev310 *Priceoraclev310Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.Priceoraclev310Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Priceoraclev310 *Priceoraclev310Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.Priceoraclev310Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Priceoraclev310 *Priceoraclev310CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Priceoraclev310.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Priceoraclev310 *Priceoraclev310TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Priceoraclev310 *Priceoraclev310TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Priceoraclev310 *Priceoraclev310Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Priceoraclev310 *Priceoraclev310Session) Acl() (common.Address, error) {
	return _Priceoraclev310.Contract.Acl(&_Priceoraclev310.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Priceoraclev310 *Priceoraclev310CallerSession) Acl() (common.Address, error) {
	return _Priceoraclev310.Contract.Acl(&_Priceoraclev310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Priceoraclev310 *Priceoraclev310Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Priceoraclev310 *Priceoraclev310Session) ContractType() ([32]byte, error) {
	return _Priceoraclev310.Contract.ContractType(&_Priceoraclev310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Priceoraclev310 *Priceoraclev310CallerSession) ContractType() ([32]byte, error) {
	return _Priceoraclev310.Contract.ContractType(&_Priceoraclev310.CallOpts)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Caller) Convert(opts *bind.CallOpts, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "convert", amount, tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Session) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.Convert(&_Priceoraclev310.CallOpts, amount, tokenFrom, tokenTo)
}

// Convert is a free data retrieval call binding the contract method 0xb66102df.
//
// Solidity: function convert(uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310CallerSession) Convert(amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.Convert(&_Priceoraclev310.CallOpts, amount, tokenFrom, tokenTo)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Caller) ConvertFromUSD(opts *bind.CallOpts, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "convertFromUSD", amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Session) ConvertFromUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.ConvertFromUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x7afb0104.
//
// Solidity: function convertFromUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310CallerSession) ConvertFromUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.ConvertFromUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Caller) ConvertToUSD(opts *bind.CallOpts, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "convertToUSD", amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Session) ConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.ConvertToUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0xf9a65030.
//
// Solidity: function convertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310CallerSession) ConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.ConvertToUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Caller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Session) GetPrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetPrice(&_Priceoraclev310.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310CallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetPrice(&_Priceoraclev310.CallOpts, token)
}

// GetReservePrice is a free data retrieval call binding the contract method 0x1321acd8.
//
// Solidity: function getReservePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Caller) GetReservePrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "getReservePrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservePrice is a free data retrieval call binding the contract method 0x1321acd8.
//
// Solidity: function getReservePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Session) GetReservePrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetReservePrice(&_Priceoraclev310.CallOpts, token)
}

// GetReservePrice is a free data retrieval call binding the contract method 0x1321acd8.
//
// Solidity: function getReservePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310CallerSession) GetReservePrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetReservePrice(&_Priceoraclev310.CallOpts, token)
}

// GetSafePrice is a free data retrieval call binding the contract method 0x3c75f448.
//
// Solidity: function getSafePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Caller) GetSafePrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "getSafePrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSafePrice is a free data retrieval call binding the contract method 0x3c75f448.
//
// Solidity: function getSafePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310Session) GetSafePrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetSafePrice(&_Priceoraclev310.CallOpts, token)
}

// GetSafePrice is a free data retrieval call binding the contract method 0x3c75f448.
//
// Solidity: function getSafePrice(address token) view returns(uint256 price)
func (_Priceoraclev310 *Priceoraclev310CallerSession) GetSafePrice(token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.GetSafePrice(&_Priceoraclev310.CallOpts, token)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Priceoraclev310 *Priceoraclev310Caller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Priceoraclev310 *Priceoraclev310Session) GetTokens() ([]common.Address, error) {
	return _Priceoraclev310.Contract.GetTokens(&_Priceoraclev310.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Priceoraclev310 *Priceoraclev310CallerSession) GetTokens() ([]common.Address, error) {
	return _Priceoraclev310.Contract.GetTokens(&_Priceoraclev310.CallOpts)
}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310Caller) PriceFeedParams(opts *bind.CallOpts, token common.Address) (PriceFeedParams, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "priceFeedParams", token)

	if err != nil {
		return *new(PriceFeedParams), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceFeedParams)).(*PriceFeedParams)

	return out0, err

}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310Session) PriceFeedParams(token common.Address) (PriceFeedParams, error) {
	return _Priceoraclev310.Contract.PriceFeedParams(&_Priceoraclev310.CallOpts, token)
}

// PriceFeedParams is a free data retrieval call binding the contract method 0xccdd1ce3.
//
// Solidity: function priceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310CallerSession) PriceFeedParams(token common.Address) (PriceFeedParams, error) {
	return _Priceoraclev310.Contract.PriceFeedParams(&_Priceoraclev310.CallOpts, token)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310Caller) PriceFeeds(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "priceFeeds", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310Session) PriceFeeds(token common.Address) (common.Address, error) {
	return _Priceoraclev310.Contract.PriceFeeds(&_Priceoraclev310.CallOpts, token)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310CallerSession) PriceFeeds(token common.Address) (common.Address, error) {
	return _Priceoraclev310.Contract.PriceFeeds(&_Priceoraclev310.CallOpts, token)
}

// ReservePriceFeedParams is a free data retrieval call binding the contract method 0x20a481e1.
//
// Solidity: function reservePriceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310Caller) ReservePriceFeedParams(opts *bind.CallOpts, token common.Address) (PriceFeedParams, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "reservePriceFeedParams", token)

	if err != nil {
		return *new(PriceFeedParams), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceFeedParams)).(*PriceFeedParams)

	return out0, err

}

// ReservePriceFeedParams is a free data retrieval call binding the contract method 0x20a481e1.
//
// Solidity: function reservePriceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310Session) ReservePriceFeedParams(token common.Address) (PriceFeedParams, error) {
	return _Priceoraclev310.Contract.ReservePriceFeedParams(&_Priceoraclev310.CallOpts, token)
}

// ReservePriceFeedParams is a free data retrieval call binding the contract method 0x20a481e1.
//
// Solidity: function reservePriceFeedParams(address token) view returns((address,uint32,bool,uint8))
func (_Priceoraclev310 *Priceoraclev310CallerSession) ReservePriceFeedParams(token common.Address) (PriceFeedParams, error) {
	return _Priceoraclev310.Contract.ReservePriceFeedParams(&_Priceoraclev310.CallOpts, token)
}

// ReservePriceFeeds is a free data retrieval call binding the contract method 0x7c70dd51.
//
// Solidity: function reservePriceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310Caller) ReservePriceFeeds(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "reservePriceFeeds", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReservePriceFeeds is a free data retrieval call binding the contract method 0x7c70dd51.
//
// Solidity: function reservePriceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310Session) ReservePriceFeeds(token common.Address) (common.Address, error) {
	return _Priceoraclev310.Contract.ReservePriceFeeds(&_Priceoraclev310.CallOpts, token)
}

// ReservePriceFeeds is a free data retrieval call binding the contract method 0x7c70dd51.
//
// Solidity: function reservePriceFeeds(address token) view returns(address)
func (_Priceoraclev310 *Priceoraclev310CallerSession) ReservePriceFeeds(token common.Address) (common.Address, error) {
	return _Priceoraclev310.Contract.ReservePriceFeeds(&_Priceoraclev310.CallOpts, token)
}

// SafeConvertToUSD is a free data retrieval call binding the contract method 0x553552fe.
//
// Solidity: function safeConvertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Caller) SafeConvertToUSD(opts *bind.CallOpts, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "safeConvertToUSD", amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SafeConvertToUSD is a free data retrieval call binding the contract method 0x553552fe.
//
// Solidity: function safeConvertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Session) SafeConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.SafeConvertToUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// SafeConvertToUSD is a free data retrieval call binding the contract method 0x553552fe.
//
// Solidity: function safeConvertToUSD(uint256 amount, address token) view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310CallerSession) SafeConvertToUSD(amount *big.Int, token common.Address) (*big.Int, error) {
	return _Priceoraclev310.Contract.SafeConvertToUSD(&_Priceoraclev310.CallOpts, amount, token)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Priceoraclev310.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310Session) Version() (*big.Int, error) {
	return _Priceoraclev310.Contract.Version(&_Priceoraclev310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Priceoraclev310 *Priceoraclev310CallerSession) Version() (*big.Int, error) {
	return _Priceoraclev310.Contract.Version(&_Priceoraclev310.CallOpts)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310Transactor) SetPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.contract.Transact(opts, "setPriceFeed", token, priceFeed, stalenessPeriod)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310Session) SetPriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.SetPriceFeed(&_Priceoraclev310.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x0214433b.
//
// Solidity: function setPriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310TransactorSession) SetPriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.SetPriceFeed(&_Priceoraclev310.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310Transactor) SetReservePriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.contract.Transact(opts, "setReservePriceFeed", token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310Session) SetReservePriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.SetReservePriceFeed(&_Priceoraclev310.TransactOpts, token, priceFeed, stalenessPeriod)
}

// SetReservePriceFeed is a paid mutator transaction binding the contract method 0x0b51dadf.
//
// Solidity: function setReservePriceFeed(address token, address priceFeed, uint32 stalenessPeriod) returns()
func (_Priceoraclev310 *Priceoraclev310TransactorSession) SetReservePriceFeed(token common.Address, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _Priceoraclev310.Contract.SetReservePriceFeed(&_Priceoraclev310.TransactOpts, token, priceFeed, stalenessPeriod)
}

// Priceoraclev310SetPriceFeedIterator is returned from FilterSetPriceFeed and is used to iterate over the raw logs and unpacked data for SetPriceFeed events raised by the Priceoraclev310 contract.
type Priceoraclev310SetPriceFeedIterator struct {
	Event *Priceoraclev310SetPriceFeed // Event containing the contract specifics and raw log

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
func (it *Priceoraclev310SetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Priceoraclev310SetPriceFeed)
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
		it.Event = new(Priceoraclev310SetPriceFeed)
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
func (it *Priceoraclev310SetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Priceoraclev310SetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Priceoraclev310SetPriceFeed represents a SetPriceFeed event raised by the Priceoraclev310 contract.
type Priceoraclev310SetPriceFeed struct {
	Token           common.Address
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetPriceFeed is a free log retrieval operation binding the contract event 0x54bafe4dd078667ca4498828ef773452f21b7c2b2f71c734bd085cab2a0f4f4b.
//
// Solidity: event SetPriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_Priceoraclev310 *Priceoraclev310Filterer) FilterSetPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*Priceoraclev310SetPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _Priceoraclev310.contract.FilterLogs(opts, "SetPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310SetPriceFeedIterator{contract: _Priceoraclev310.contract, event: "SetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetPriceFeed is a free log subscription operation binding the contract event 0x54bafe4dd078667ca4498828ef773452f21b7c2b2f71c734bd085cab2a0f4f4b.
//
// Solidity: event SetPriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_Priceoraclev310 *Priceoraclev310Filterer) WatchSetPriceFeed(opts *bind.WatchOpts, sink chan<- *Priceoraclev310SetPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _Priceoraclev310.contract.WatchLogs(opts, "SetPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Priceoraclev310SetPriceFeed)
				if err := _Priceoraclev310.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
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
func (_Priceoraclev310 *Priceoraclev310Filterer) ParseSetPriceFeed(log types.Log) (*Priceoraclev310SetPriceFeed, error) {
	event := new(Priceoraclev310SetPriceFeed)
	if err := _Priceoraclev310.contract.UnpackLog(event, "SetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Priceoraclev310SetReservePriceFeedIterator is returned from FilterSetReservePriceFeed and is used to iterate over the raw logs and unpacked data for SetReservePriceFeed events raised by the Priceoraclev310 contract.
type Priceoraclev310SetReservePriceFeedIterator struct {
	Event *Priceoraclev310SetReservePriceFeed // Event containing the contract specifics and raw log

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
func (it *Priceoraclev310SetReservePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Priceoraclev310SetReservePriceFeed)
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
		it.Event = new(Priceoraclev310SetReservePriceFeed)
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
func (it *Priceoraclev310SetReservePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Priceoraclev310SetReservePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Priceoraclev310SetReservePriceFeed represents a SetReservePriceFeed event raised by the Priceoraclev310 contract.
type Priceoraclev310SetReservePriceFeed struct {
	Token           common.Address
	PriceFeed       common.Address
	StalenessPeriod uint32
	SkipCheck       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReservePriceFeed is a free log retrieval operation binding the contract event 0x0622a3fd57e320160fa23aaad98e12139a7922016875937c646e76237005c01c.
//
// Solidity: event SetReservePriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_Priceoraclev310 *Priceoraclev310Filterer) FilterSetReservePriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*Priceoraclev310SetReservePriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _Priceoraclev310.contract.FilterLogs(opts, "SetReservePriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &Priceoraclev310SetReservePriceFeedIterator{contract: _Priceoraclev310.contract, event: "SetReservePriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetReservePriceFeed is a free log subscription operation binding the contract event 0x0622a3fd57e320160fa23aaad98e12139a7922016875937c646e76237005c01c.
//
// Solidity: event SetReservePriceFeed(address indexed token, address indexed priceFeed, uint32 stalenessPeriod, bool skipCheck)
func (_Priceoraclev310 *Priceoraclev310Filterer) WatchSetReservePriceFeed(opts *bind.WatchOpts, sink chan<- *Priceoraclev310SetReservePriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _Priceoraclev310.contract.WatchLogs(opts, "SetReservePriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Priceoraclev310SetReservePriceFeed)
				if err := _Priceoraclev310.contract.UnpackLog(event, "SetReservePriceFeed", log); err != nil {
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
func (_Priceoraclev310 *Priceoraclev310Filterer) ParseSetReservePriceFeed(log types.Log) (*Priceoraclev310SetReservePriceFeed, error) {
	event := new(Priceoraclev310SetReservePriceFeed)
	if err := _Priceoraclev310.contract.UnpackLog(event, "SetReservePriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
