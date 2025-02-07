// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compositeOraclev3

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
}

// CompositeOraclev3MetaData contains all meta data concerning the CompositeOraclev3 contract.
var CompositeOraclev3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"internalType\":\"structPriceFeedParams[2]\",\"name\":\"priceFeeds\",\"type\":\"tuple[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StalePriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeedType\",\"outputs\":[{\"internalType\":\"enumPriceFeedType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipCheck1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipPriceCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stalenessPeriod0\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stalenessPeriod1\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetFeedScale\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CompositeOraclev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use CompositeOraclev3MetaData.ABI instead.
var CompositeOraclev3ABI = CompositeOraclev3MetaData.ABI

// CompositeOraclev3 is an auto generated Go binding around an Ethereum contract.
type CompositeOraclev3 struct {
	CompositeOraclev3Caller     // Read-only binding to the contract
	CompositeOraclev3Transactor // Write-only binding to the contract
	CompositeOraclev3Filterer   // Log filterer for contract events
}

// CompositeOraclev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type CompositeOraclev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeOraclev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CompositeOraclev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeOraclev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompositeOraclev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompositeOraclev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompositeOraclev3Session struct {
	Contract     *CompositeOraclev3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CompositeOraclev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompositeOraclev3CallerSession struct {
	Contract *CompositeOraclev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CompositeOraclev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompositeOraclev3TransactorSession struct {
	Contract     *CompositeOraclev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CompositeOraclev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type CompositeOraclev3Raw struct {
	Contract *CompositeOraclev3 // Generic contract binding to access the raw methods on
}

// CompositeOraclev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompositeOraclev3CallerRaw struct {
	Contract *CompositeOraclev3Caller // Generic read-only contract binding to access the raw methods on
}

// CompositeOraclev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompositeOraclev3TransactorRaw struct {
	Contract *CompositeOraclev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCompositeOraclev3 creates a new instance of CompositeOraclev3, bound to a specific deployed contract.
func NewCompositeOraclev3(address common.Address, backend bind.ContractBackend) (*CompositeOraclev3, error) {
	contract, err := bindCompositeOraclev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompositeOraclev3{CompositeOraclev3Caller: CompositeOraclev3Caller{contract: contract}, CompositeOraclev3Transactor: CompositeOraclev3Transactor{contract: contract}, CompositeOraclev3Filterer: CompositeOraclev3Filterer{contract: contract}}, nil
}

// NewCompositeOraclev3Caller creates a new read-only instance of CompositeOraclev3, bound to a specific deployed contract.
func NewCompositeOraclev3Caller(address common.Address, caller bind.ContractCaller) (*CompositeOraclev3Caller, error) {
	contract, err := bindCompositeOraclev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompositeOraclev3Caller{contract: contract}, nil
}

// NewCompositeOraclev3Transactor creates a new write-only instance of CompositeOraclev3, bound to a specific deployed contract.
func NewCompositeOraclev3Transactor(address common.Address, transactor bind.ContractTransactor) (*CompositeOraclev3Transactor, error) {
	contract, err := bindCompositeOraclev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompositeOraclev3Transactor{contract: contract}, nil
}

// NewCompositeOraclev3Filterer creates a new log filterer instance of CompositeOraclev3, bound to a specific deployed contract.
func NewCompositeOraclev3Filterer(address common.Address, filterer bind.ContractFilterer) (*CompositeOraclev3Filterer, error) {
	contract, err := bindCompositeOraclev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompositeOraclev3Filterer{contract: contract}, nil
}

// bindCompositeOraclev3 binds a generic wrapper to an already deployed contract.
func bindCompositeOraclev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompositeOraclev3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompositeOraclev3 *CompositeOraclev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompositeOraclev3.Contract.CompositeOraclev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompositeOraclev3 *CompositeOraclev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompositeOraclev3.Contract.CompositeOraclev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompositeOraclev3 *CompositeOraclev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompositeOraclev3.Contract.CompositeOraclev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompositeOraclev3 *CompositeOraclev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompositeOraclev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompositeOraclev3 *CompositeOraclev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompositeOraclev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompositeOraclev3 *CompositeOraclev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompositeOraclev3.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3Session) Decimals() (uint8, error) {
	return _CompositeOraclev3.Contract.Decimals(&_CompositeOraclev3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) Decimals() (uint8, error) {
	return _CompositeOraclev3.Contract.Decimals(&_CompositeOraclev3.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_CompositeOraclev3 *CompositeOraclev3Caller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_CompositeOraclev3 *CompositeOraclev3Session) Description() (string, error) {
	return _CompositeOraclev3.Contract.Description(&_CompositeOraclev3.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) Description() (string, error) {
	return _CompositeOraclev3.Contract.Description(&_CompositeOraclev3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256, uint80)
func (_CompositeOraclev3 *CompositeOraclev3Caller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "latestRoundData")

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
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256, uint80)
func (_CompositeOraclev3 *CompositeOraclev3Session) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CompositeOraclev3.Contract.LatestRoundData(&_CompositeOraclev3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256, uint80)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CompositeOraclev3.Contract.LatestRoundData(&_CompositeOraclev3.CallOpts)
}

// PriceFeed0 is a free data retrieval call binding the contract method 0x385aee1b.
//
// Solidity: function priceFeed0() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3Caller) PriceFeed0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "priceFeed0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed0 is a free data retrieval call binding the contract method 0x385aee1b.
//
// Solidity: function priceFeed0() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3Session) PriceFeed0() (common.Address, error) {
	return _CompositeOraclev3.Contract.PriceFeed0(&_CompositeOraclev3.CallOpts)
}

// PriceFeed0 is a free data retrieval call binding the contract method 0x385aee1b.
//
// Solidity: function priceFeed0() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) PriceFeed0() (common.Address, error) {
	return _CompositeOraclev3.Contract.PriceFeed0(&_CompositeOraclev3.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3Caller) PriceFeed1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "priceFeed1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3Session) PriceFeed1() (common.Address, error) {
	return _CompositeOraclev3.Contract.PriceFeed1(&_CompositeOraclev3.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() view returns(address)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) PriceFeed1() (common.Address, error) {
	return _CompositeOraclev3.Contract.PriceFeed1(&_CompositeOraclev3.CallOpts)
}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3Caller) PriceFeedType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "priceFeedType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3Session) PriceFeedType() (uint8, error) {
	return _CompositeOraclev3.Contract.PriceFeedType(&_CompositeOraclev3.CallOpts)
}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) PriceFeedType() (uint8, error) {
	return _CompositeOraclev3.Contract.PriceFeedType(&_CompositeOraclev3.CallOpts)
}

// SkipCheck1 is a free data retrieval call binding the contract method 0x7ff361ec.
//
// Solidity: function skipCheck1() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3Caller) SkipCheck1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "skipCheck1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipCheck1 is a free data retrieval call binding the contract method 0x7ff361ec.
//
// Solidity: function skipCheck1() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3Session) SkipCheck1() (bool, error) {
	return _CompositeOraclev3.Contract.SkipCheck1(&_CompositeOraclev3.CallOpts)
}

// SkipCheck1 is a free data retrieval call binding the contract method 0x7ff361ec.
//
// Solidity: function skipCheck1() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) SkipCheck1() (bool, error) {
	return _CompositeOraclev3.Contract.SkipCheck1(&_CompositeOraclev3.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3Caller) SkipPriceCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "skipPriceCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3Session) SkipPriceCheck() (bool, error) {
	return _CompositeOraclev3.Contract.SkipPriceCheck(&_CompositeOraclev3.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) SkipPriceCheck() (bool, error) {
	return _CompositeOraclev3.Contract.SkipPriceCheck(&_CompositeOraclev3.CallOpts)
}

// StalenessPeriod0 is a free data retrieval call binding the contract method 0x178793e8.
//
// Solidity: function stalenessPeriod0() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3Caller) StalenessPeriod0(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "stalenessPeriod0")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// StalenessPeriod0 is a free data retrieval call binding the contract method 0x178793e8.
//
// Solidity: function stalenessPeriod0() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3Session) StalenessPeriod0() (uint32, error) {
	return _CompositeOraclev3.Contract.StalenessPeriod0(&_CompositeOraclev3.CallOpts)
}

// StalenessPeriod0 is a free data retrieval call binding the contract method 0x178793e8.
//
// Solidity: function stalenessPeriod0() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) StalenessPeriod0() (uint32, error) {
	return _CompositeOraclev3.Contract.StalenessPeriod0(&_CompositeOraclev3.CallOpts)
}

// StalenessPeriod1 is a free data retrieval call binding the contract method 0x3e777fd2.
//
// Solidity: function stalenessPeriod1() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3Caller) StalenessPeriod1(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "stalenessPeriod1")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// StalenessPeriod1 is a free data retrieval call binding the contract method 0x3e777fd2.
//
// Solidity: function stalenessPeriod1() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3Session) StalenessPeriod1() (uint32, error) {
	return _CompositeOraclev3.Contract.StalenessPeriod1(&_CompositeOraclev3.CallOpts)
}

// StalenessPeriod1 is a free data retrieval call binding the contract method 0x3e777fd2.
//
// Solidity: function stalenessPeriod1() view returns(uint32)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) StalenessPeriod1() (uint32, error) {
	return _CompositeOraclev3.Contract.StalenessPeriod1(&_CompositeOraclev3.CallOpts)
}

// TargetFeedScale is a free data retrieval call binding the contract method 0x0ff8bdd2.
//
// Solidity: function targetFeedScale() view returns(int256)
func (_CompositeOraclev3 *CompositeOraclev3Caller) TargetFeedScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "targetFeedScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetFeedScale is a free data retrieval call binding the contract method 0x0ff8bdd2.
//
// Solidity: function targetFeedScale() view returns(int256)
func (_CompositeOraclev3 *CompositeOraclev3Session) TargetFeedScale() (*big.Int, error) {
	return _CompositeOraclev3.Contract.TargetFeedScale(&_CompositeOraclev3.CallOpts)
}

// TargetFeedScale is a free data retrieval call binding the contract method 0x0ff8bdd2.
//
// Solidity: function targetFeedScale() view returns(int256)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) TargetFeedScale() (*big.Int, error) {
	return _CompositeOraclev3.Contract.TargetFeedScale(&_CompositeOraclev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CompositeOraclev3 *CompositeOraclev3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompositeOraclev3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CompositeOraclev3 *CompositeOraclev3Session) Version() (*big.Int, error) {
	return _CompositeOraclev3.Contract.Version(&_CompositeOraclev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CompositeOraclev3 *CompositeOraclev3CallerSession) Version() (*big.Int, error) {
	return _CompositeOraclev3.Contract.Version(&_CompositeOraclev3.CallOpts)
}
