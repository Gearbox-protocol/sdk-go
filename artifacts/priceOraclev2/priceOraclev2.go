// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceOraclev2

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

// PriceFeedConfig is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedConfig struct {
	Token     common.Address
	PriceFeed common.Address
}

// PriceOraclev2MetaData contains all meta data concerning the PriceOraclev2 contract.
var PriceOraclev2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"internalType\":\"structPriceFeedConfig[]\",\"name\":\"defaults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainPriceStaleException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedRequiresAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceOracleNotExistsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPriceException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"NewPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acl\",\"outputs\":[{\"internalType\":\"contractACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"addPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"convert\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertFromUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"convertToUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountFrom\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTo\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"}],\"name\":\"fastCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralTo\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeeds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"priceFeedsWithFlags\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"dependsOnAddress\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"skipCheck\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceOraclev2ABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceOraclev2MetaData.ABI instead.
var PriceOraclev2ABI = PriceOraclev2MetaData.ABI

// PriceOraclev2 is an auto generated Go binding around an Ethereum contract.
type PriceOraclev2 struct {
	PriceOraclev2Caller     // Read-only binding to the contract
	PriceOraclev2Transactor // Write-only binding to the contract
	PriceOraclev2Filterer   // Log filterer for contract events
}

// PriceOraclev2Caller is an auto generated read-only Go binding around an Ethereum contract.
type PriceOraclev2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceOraclev2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceOraclev2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceOraclev2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceOraclev2Session struct {
	Contract     *PriceOraclev2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceOraclev2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceOraclev2CallerSession struct {
	Contract *PriceOraclev2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriceOraclev2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceOraclev2TransactorSession struct {
	Contract     *PriceOraclev2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriceOraclev2Raw is an auto generated low-level Go binding around an Ethereum contract.
type PriceOraclev2Raw struct {
	Contract *PriceOraclev2 // Generic contract binding to access the raw methods on
}

// PriceOraclev2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceOraclev2CallerRaw struct {
	Contract *PriceOraclev2Caller // Generic read-only contract binding to access the raw methods on
}

// PriceOraclev2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceOraclev2TransactorRaw struct {
	Contract *PriceOraclev2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceOraclev2 creates a new instance of PriceOraclev2, bound to a specific deployed contract.
func NewPriceOraclev2(address common.Address, backend bind.ContractBackend) (*PriceOraclev2, error) {
	contract, err := bindPriceOraclev2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2{PriceOraclev2Caller: PriceOraclev2Caller{contract: contract}, PriceOraclev2Transactor: PriceOraclev2Transactor{contract: contract}, PriceOraclev2Filterer: PriceOraclev2Filterer{contract: contract}}, nil
}

// NewPriceOraclev2Caller creates a new read-only instance of PriceOraclev2, bound to a specific deployed contract.
func NewPriceOraclev2Caller(address common.Address, caller bind.ContractCaller) (*PriceOraclev2Caller, error) {
	contract, err := bindPriceOraclev2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2Caller{contract: contract}, nil
}

// NewPriceOraclev2Transactor creates a new write-only instance of PriceOraclev2, bound to a specific deployed contract.
func NewPriceOraclev2Transactor(address common.Address, transactor bind.ContractTransactor) (*PriceOraclev2Transactor, error) {
	contract, err := bindPriceOraclev2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2Transactor{contract: contract}, nil
}

// NewPriceOraclev2Filterer creates a new log filterer instance of PriceOraclev2, bound to a specific deployed contract.
func NewPriceOraclev2Filterer(address common.Address, filterer bind.ContractFilterer) (*PriceOraclev2Filterer, error) {
	contract, err := bindPriceOraclev2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2Filterer{contract: contract}, nil
}

// bindPriceOraclev2 binds a generic wrapper to an already deployed contract.
func bindPriceOraclev2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceOraclev2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOraclev2 *PriceOraclev2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOraclev2.Contract.PriceOraclev2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOraclev2 *PriceOraclev2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.PriceOraclev2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOraclev2 *PriceOraclev2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.PriceOraclev2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceOraclev2 *PriceOraclev2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceOraclev2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceOraclev2 *PriceOraclev2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceOraclev2 *PriceOraclev2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_PriceOraclev2 *PriceOraclev2Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "_acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_PriceOraclev2 *PriceOraclev2Session) Acl() (common.Address, error) {
	return _PriceOraclev2.Contract.Acl(&_PriceOraclev2.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_PriceOraclev2 *PriceOraclev2CallerSession) Acl() (common.Address, error) {
	return _PriceOraclev2.Contract.Acl(&_PriceOraclev2.CallOpts)
}

// Convert is a free data retrieval call binding the contract method 0xd6d19b27.
//
// Solidity: function convert(address creditAccount, uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) Convert(opts *bind.CallOpts, creditAccount common.Address, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "convert", creditAccount, amount, tokenFrom, tokenTo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Convert is a free data retrieval call binding the contract method 0xd6d19b27.
//
// Solidity: function convert(address creditAccount, uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) Convert(creditAccount common.Address, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.Convert(&_PriceOraclev2.CallOpts, creditAccount, amount, tokenFrom, tokenTo)
}

// Convert is a free data retrieval call binding the contract method 0xd6d19b27.
//
// Solidity: function convert(address creditAccount, uint256 amount, address tokenFrom, address tokenTo) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) Convert(creditAccount common.Address, amount *big.Int, tokenFrom common.Address, tokenTo common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.Convert(&_PriceOraclev2.CallOpts, creditAccount, amount, tokenFrom, tokenTo)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x902003df.
//
// Solidity: function convertFromUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) ConvertFromUSD(opts *bind.CallOpts, creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "convertFromUSD", creditAccount, amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x902003df.
//
// Solidity: function convertFromUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) ConvertFromUSD(creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.ConvertFromUSD(&_PriceOraclev2.CallOpts, creditAccount, amount, token)
}

// ConvertFromUSD is a free data retrieval call binding the contract method 0x902003df.
//
// Solidity: function convertFromUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) ConvertFromUSD(creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.ConvertFromUSD(&_PriceOraclev2.CallOpts, creditAccount, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0x5b461005.
//
// Solidity: function convertToUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) ConvertToUSD(opts *bind.CallOpts, creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "convertToUSD", creditAccount, amount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToUSD is a free data retrieval call binding the contract method 0x5b461005.
//
// Solidity: function convertToUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) ConvertToUSD(creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.ConvertToUSD(&_PriceOraclev2.CallOpts, creditAccount, amount, token)
}

// ConvertToUSD is a free data retrieval call binding the contract method 0x5b461005.
//
// Solidity: function convertToUSD(address creditAccount, uint256 amount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) ConvertToUSD(creditAccount common.Address, amount *big.Int, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.ConvertToUSD(&_PriceOraclev2.CallOpts, creditAccount, amount, token)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address ) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) Decimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "decimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address ) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) Decimals(arg0 common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.Decimals(&_PriceOraclev2.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0xd449a832.
//
// Solidity: function decimals(address ) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) Decimals(arg0 common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.Decimals(&_PriceOraclev2.CallOpts, arg0)
}

// FastCheck is a free data retrieval call binding the contract method 0x5cecbd0e.
//
// Solidity: function fastCheck(uint256 amountFrom, address tokenFrom, uint256 amountTo, address tokenTo) view returns(uint256 collateralFrom, uint256 collateralTo)
func (_PriceOraclev2 *PriceOraclev2Caller) FastCheck(opts *bind.CallOpts, amountFrom *big.Int, tokenFrom common.Address, amountTo *big.Int, tokenTo common.Address) (struct {
	CollateralFrom *big.Int
	CollateralTo   *big.Int
}, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "fastCheck", amountFrom, tokenFrom, amountTo, tokenTo)

	outstruct := new(struct {
		CollateralFrom *big.Int
		CollateralTo   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralFrom = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralTo = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FastCheck is a free data retrieval call binding the contract method 0x5cecbd0e.
//
// Solidity: function fastCheck(uint256 amountFrom, address tokenFrom, uint256 amountTo, address tokenTo) view returns(uint256 collateralFrom, uint256 collateralTo)
func (_PriceOraclev2 *PriceOraclev2Session) FastCheck(amountFrom *big.Int, tokenFrom common.Address, amountTo *big.Int, tokenTo common.Address) (struct {
	CollateralFrom *big.Int
	CollateralTo   *big.Int
}, error) {
	return _PriceOraclev2.Contract.FastCheck(&_PriceOraclev2.CallOpts, amountFrom, tokenFrom, amountTo, tokenTo)
}

// FastCheck is a free data retrieval call binding the contract method 0x5cecbd0e.
//
// Solidity: function fastCheck(uint256 amountFrom, address tokenFrom, uint256 amountTo, address tokenTo) view returns(uint256 collateralFrom, uint256 collateralTo)
func (_PriceOraclev2 *PriceOraclev2CallerSession) FastCheck(amountFrom *big.Int, tokenFrom common.Address, amountTo *big.Int, tokenTo common.Address) (struct {
	CollateralFrom *big.Int
	CollateralTo   *big.Int
}, error) {
	return _PriceOraclev2.Contract.FastCheck(&_PriceOraclev2.CallOpts, amountFrom, tokenFrom, amountTo, tokenTo)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address creditAccount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) GetPrice(opts *bind.CallOpts, creditAccount common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "getPrice", creditAccount, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address creditAccount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) GetPrice(creditAccount common.Address, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.GetPrice(&_PriceOraclev2.CallOpts, creditAccount, token)
}

// GetPrice is a free data retrieval call binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address creditAccount, address token) view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) GetPrice(creditAccount common.Address, token common.Address) (*big.Int, error) {
	return _PriceOraclev2.Contract.GetPrice(&_PriceOraclev2.CallOpts, creditAccount, token)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev2 *PriceOraclev2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev2 *PriceOraclev2Session) Paused() (bool, error) {
	return _PriceOraclev2.Contract.Paused(&_PriceOraclev2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PriceOraclev2 *PriceOraclev2CallerSession) Paused() (bool, error) {
	return _PriceOraclev2.Contract.Paused(&_PriceOraclev2.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev2 *PriceOraclev2Caller) PriceFeeds(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "priceFeeds", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev2 *PriceOraclev2Session) PriceFeeds(token common.Address) (common.Address, error) {
	return _PriceOraclev2.Contract.PriceFeeds(&_PriceOraclev2.CallOpts, token)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address token) view returns(address priceFeed)
func (_PriceOraclev2 *PriceOraclev2CallerSession) PriceFeeds(token common.Address) (common.Address, error) {
	return _PriceOraclev2.Contract.PriceFeeds(&_PriceOraclev2.CallOpts, token)
}

// PriceFeedsWithFlags is a free data retrieval call binding the contract method 0xf190e5fa.
//
// Solidity: function priceFeedsWithFlags(address token) view returns(address priceFeed, bool dependsOnAddress, bool skipCheck)
func (_PriceOraclev2 *PriceOraclev2Caller) PriceFeedsWithFlags(opts *bind.CallOpts, token common.Address) (struct {
	PriceFeed        common.Address
	DependsOnAddress bool
	SkipCheck        bool
}, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "priceFeedsWithFlags", token)

	outstruct := new(struct {
		PriceFeed        common.Address
		DependsOnAddress bool
		SkipCheck        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceFeed = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DependsOnAddress = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.SkipCheck = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// PriceFeedsWithFlags is a free data retrieval call binding the contract method 0xf190e5fa.
//
// Solidity: function priceFeedsWithFlags(address token) view returns(address priceFeed, bool dependsOnAddress, bool skipCheck)
func (_PriceOraclev2 *PriceOraclev2Session) PriceFeedsWithFlags(token common.Address) (struct {
	PriceFeed        common.Address
	DependsOnAddress bool
	SkipCheck        bool
}, error) {
	return _PriceOraclev2.Contract.PriceFeedsWithFlags(&_PriceOraclev2.CallOpts, token)
}

// PriceFeedsWithFlags is a free data retrieval call binding the contract method 0xf190e5fa.
//
// Solidity: function priceFeedsWithFlags(address token) view returns(address priceFeed, bool dependsOnAddress, bool skipCheck)
func (_PriceOraclev2 *PriceOraclev2CallerSession) PriceFeedsWithFlags(token common.Address) (struct {
	PriceFeed        common.Address
	DependsOnAddress bool
	SkipCheck        bool
}, error) {
	return _PriceOraclev2.Contract.PriceFeedsWithFlags(&_PriceOraclev2.CallOpts, token)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceOraclev2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2Session) Version() (*big.Int, error) {
	return _PriceOraclev2.Contract.Version(&_PriceOraclev2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceOraclev2 *PriceOraclev2CallerSession) Version() (*big.Int, error) {
	return _PriceOraclev2.Contract.Version(&_PriceOraclev2.CallOpts)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOraclev2 *PriceOraclev2Transactor) AddPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOraclev2.contract.Transact(opts, "addPriceFeed", token, priceFeed)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOraclev2 *PriceOraclev2Session) AddPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.AddPriceFeed(&_PriceOraclev2.TransactOpts, token, priceFeed)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0xe8a97a3e.
//
// Solidity: function addPriceFeed(address token, address priceFeed) returns()
func (_PriceOraclev2 *PriceOraclev2TransactorSession) AddPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceOraclev2.Contract.AddPriceFeed(&_PriceOraclev2.TransactOpts, token, priceFeed)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev2 *PriceOraclev2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev2 *PriceOraclev2Session) Pause() (*types.Transaction, error) {
	return _PriceOraclev2.Contract.Pause(&_PriceOraclev2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PriceOraclev2 *PriceOraclev2TransactorSession) Pause() (*types.Transaction, error) {
	return _PriceOraclev2.Contract.Pause(&_PriceOraclev2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev2 *PriceOraclev2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceOraclev2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev2 *PriceOraclev2Session) Unpause() (*types.Transaction, error) {
	return _PriceOraclev2.Contract.Unpause(&_PriceOraclev2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PriceOraclev2 *PriceOraclev2TransactorSession) Unpause() (*types.Transaction, error) {
	return _PriceOraclev2.Contract.Unpause(&_PriceOraclev2.TransactOpts)
}

// PriceOraclev2NewPriceFeedIterator is returned from FilterNewPriceFeed and is used to iterate over the raw logs and unpacked data for NewPriceFeed events raised by the PriceOraclev2 contract.
type PriceOraclev2NewPriceFeedIterator struct {
	Event *PriceOraclev2NewPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceOraclev2NewPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev2NewPriceFeed)
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
		it.Event = new(PriceOraclev2NewPriceFeed)
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
func (it *PriceOraclev2NewPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev2NewPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev2NewPriceFeed represents a NewPriceFeed event raised by the PriceOraclev2 contract.
type PriceOraclev2NewPriceFeed struct {
	Token     common.Address
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewPriceFeed is a free log retrieval operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOraclev2 *PriceOraclev2Filterer) FilterNewPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceOraclev2NewPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev2.contract.FilterLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2NewPriceFeedIterator{contract: _PriceOraclev2.contract, event: "NewPriceFeed", logs: logs, sub: sub}, nil
}

// WatchNewPriceFeed is a free log subscription operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOraclev2 *PriceOraclev2Filterer) WatchNewPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceOraclev2NewPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceOraclev2.contract.WatchLogs(opts, "NewPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev2NewPriceFeed)
				if err := _PriceOraclev2.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
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

// ParseNewPriceFeed is a log parse operation binding the contract event 0xe263805b03657ab13064915d0723c5ce14981547e7cba5283f66b9e5d81f6e6e.
//
// Solidity: event NewPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceOraclev2 *PriceOraclev2Filterer) ParseNewPriceFeed(log types.Log) (*PriceOraclev2NewPriceFeed, error) {
	event := new(PriceOraclev2NewPriceFeed)
	if err := _PriceOraclev2.contract.UnpackLog(event, "NewPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PriceOraclev2 contract.
type PriceOraclev2PausedIterator struct {
	Event *PriceOraclev2Paused // Event containing the contract specifics and raw log

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
func (it *PriceOraclev2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev2Paused)
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
		it.Event = new(PriceOraclev2Paused)
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
func (it *PriceOraclev2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev2Paused represents a Paused event raised by the PriceOraclev2 contract.
type PriceOraclev2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOraclev2 *PriceOraclev2Filterer) FilterPaused(opts *bind.FilterOpts) (*PriceOraclev2PausedIterator, error) {

	logs, sub, err := _PriceOraclev2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2PausedIterator{contract: _PriceOraclev2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PriceOraclev2 *PriceOraclev2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PriceOraclev2Paused) (event.Subscription, error) {

	logs, sub, err := _PriceOraclev2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev2Paused)
				if err := _PriceOraclev2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PriceOraclev2 *PriceOraclev2Filterer) ParsePaused(log types.Log) (*PriceOraclev2Paused, error) {
	event := new(PriceOraclev2Paused)
	if err := _PriceOraclev2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceOraclev2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PriceOraclev2 contract.
type PriceOraclev2UnpausedIterator struct {
	Event *PriceOraclev2Unpaused // Event containing the contract specifics and raw log

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
func (it *PriceOraclev2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceOraclev2Unpaused)
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
		it.Event = new(PriceOraclev2Unpaused)
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
func (it *PriceOraclev2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceOraclev2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceOraclev2Unpaused represents a Unpaused event raised by the PriceOraclev2 contract.
type PriceOraclev2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOraclev2 *PriceOraclev2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PriceOraclev2UnpausedIterator, error) {

	logs, sub, err := _PriceOraclev2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PriceOraclev2UnpausedIterator{contract: _PriceOraclev2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PriceOraclev2 *PriceOraclev2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PriceOraclev2Unpaused) (event.Subscription, error) {

	logs, sub, err := _PriceOraclev2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceOraclev2Unpaused)
				if err := _PriceOraclev2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PriceOraclev2 *PriceOraclev2Filterer) ParseUnpaused(log types.Log) (*PriceOraclev2Unpaused, error) {
	event := new(PriceOraclev2Unpaused)
	if err := _PriceOraclev2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
