// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearnv2Adapter

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

// Yearnv2AdapterMetaData contains all meta data concerning the Yearnv2Adapter contract.
var Yearnv2AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Yearnv2AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use Yearnv2AdapterMetaData.ABI instead.
var Yearnv2AdapterABI = Yearnv2AdapterMetaData.ABI

// Yearnv2Adapter is an auto generated Go binding around an Ethereum contract.
type Yearnv2Adapter struct {
	Yearnv2AdapterCaller     // Read-only binding to the contract
	Yearnv2AdapterTransactor // Write-only binding to the contract
	Yearnv2AdapterFilterer   // Log filterer for contract events
}

// Yearnv2AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Yearnv2AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yearnv2AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Yearnv2AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yearnv2AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Yearnv2AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Yearnv2AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Yearnv2AdapterSession struct {
	Contract     *Yearnv2Adapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Yearnv2AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Yearnv2AdapterCallerSession struct {
	Contract *Yearnv2AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Yearnv2AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Yearnv2AdapterTransactorSession struct {
	Contract     *Yearnv2AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Yearnv2AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Yearnv2AdapterRaw struct {
	Contract *Yearnv2Adapter // Generic contract binding to access the raw methods on
}

// Yearnv2AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Yearnv2AdapterCallerRaw struct {
	Contract *Yearnv2AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// Yearnv2AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Yearnv2AdapterTransactorRaw struct {
	Contract *Yearnv2AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnv2Adapter creates a new instance of Yearnv2Adapter, bound to a specific deployed contract.
func NewYearnv2Adapter(address common.Address, backend bind.ContractBackend) (*Yearnv2Adapter, error) {
	contract, err := bindYearnv2Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yearnv2Adapter{Yearnv2AdapterCaller: Yearnv2AdapterCaller{contract: contract}, Yearnv2AdapterTransactor: Yearnv2AdapterTransactor{contract: contract}, Yearnv2AdapterFilterer: Yearnv2AdapterFilterer{contract: contract}}, nil
}

// NewYearnv2AdapterCaller creates a new read-only instance of Yearnv2Adapter, bound to a specific deployed contract.
func NewYearnv2AdapterCaller(address common.Address, caller bind.ContractCaller) (*Yearnv2AdapterCaller, error) {
	contract, err := bindYearnv2Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Yearnv2AdapterCaller{contract: contract}, nil
}

// NewYearnv2AdapterTransactor creates a new write-only instance of Yearnv2Adapter, bound to a specific deployed contract.
func NewYearnv2AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*Yearnv2AdapterTransactor, error) {
	contract, err := bindYearnv2Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Yearnv2AdapterTransactor{contract: contract}, nil
}

// NewYearnv2AdapterFilterer creates a new log filterer instance of Yearnv2Adapter, bound to a specific deployed contract.
func NewYearnv2AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*Yearnv2AdapterFilterer, error) {
	contract, err := bindYearnv2Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Yearnv2AdapterFilterer{contract: contract}, nil
}

// bindYearnv2Adapter binds a generic wrapper to an already deployed contract.
func bindYearnv2Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Yearnv2AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yearnv2Adapter *Yearnv2AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yearnv2Adapter.Contract.Yearnv2AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yearnv2Adapter *Yearnv2AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Yearnv2AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yearnv2Adapter *Yearnv2AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Yearnv2AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yearnv2Adapter *Yearnv2AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yearnv2Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yearnv2Adapter *Yearnv2AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yearnv2Adapter *Yearnv2AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterSession) GearboxAdapterType() (uint8, error) {
	return _Yearnv2Adapter.Contract.GearboxAdapterType(&_Yearnv2Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _Yearnv2Adapter.Contract.GearboxAdapterType(&_Yearnv2Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Yearnv2Adapter *Yearnv2AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _Yearnv2Adapter.Contract.GearboxAdapterVersion(&_Yearnv2Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Yearnv2Adapter.Contract.GearboxAdapterVersion(&_Yearnv2Adapter.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yearnv2Adapter.Contract.Allowance(&_Yearnv2Adapter.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yearnv2Adapter.Contract.Allowance(&_Yearnv2Adapter.CallOpts, owner, spender)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Approve(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "approve", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Approve(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.Approve(&_Yearnv2Adapter.CallOpts, arg0, arg1)
}

// Approve is a free data retrieval call binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Approve(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.Approve(&_Yearnv2Adapter.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Yearnv2Adapter.Contract.BalanceOf(&_Yearnv2Adapter.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Yearnv2Adapter.Contract.BalanceOf(&_Yearnv2Adapter.CallOpts, account)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterSession) CreditFacade() (common.Address, error) {
	return _Yearnv2Adapter.Contract.CreditFacade(&_Yearnv2Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) CreditFacade() (common.Address, error) {
	return _Yearnv2Adapter.Contract.CreditFacade(&_Yearnv2Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterSession) CreditManager() (common.Address, error) {
	return _Yearnv2Adapter.Contract.CreditManager(&_Yearnv2Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) CreditManager() (common.Address, error) {
	return _Yearnv2Adapter.Contract.CreditManager(&_Yearnv2Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Decimals() (uint8, error) {
	return _Yearnv2Adapter.Contract.Decimals(&_Yearnv2Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Decimals() (uint8, error) {
	return _Yearnv2Adapter.Contract.Decimals(&_Yearnv2Adapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Name() (string, error) {
	return _Yearnv2Adapter.Contract.Name(&_Yearnv2Adapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Name() (string, error) {
	return _Yearnv2Adapter.Contract.Name(&_Yearnv2Adapter.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) PricePerShare() (*big.Int, error) {
	return _Yearnv2Adapter.Contract.PricePerShare(&_Yearnv2Adapter.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) PricePerShare() (*big.Int, error) {
	return _Yearnv2Adapter.Contract.PricePerShare(&_Yearnv2Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Symbol() (string, error) {
	return _Yearnv2Adapter.Contract.Symbol(&_Yearnv2Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Symbol() (string, error) {
	return _Yearnv2Adapter.Contract.Symbol(&_Yearnv2Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterSession) TargetContract() (common.Address, error) {
	return _Yearnv2Adapter.Contract.TargetContract(&_Yearnv2Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) TargetContract() (common.Address, error) {
	return _Yearnv2Adapter.Contract.TargetContract(&_Yearnv2Adapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Token() (common.Address, error) {
	return _Yearnv2Adapter.Contract.Token(&_Yearnv2Adapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Token() (common.Address, error) {
	return _Yearnv2Adapter.Contract.Token(&_Yearnv2Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) TotalSupply() (*big.Int, error) {
	return _Yearnv2Adapter.Contract.TotalSupply(&_Yearnv2Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _Yearnv2Adapter.Contract.TotalSupply(&_Yearnv2Adapter.CallOpts)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) Transfer(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "transfer", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Transfer(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.Transfer(&_Yearnv2Adapter.CallOpts, arg0, arg1)
}

// Transfer is a free data retrieval call binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) Transfer(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.Transfer(&_Yearnv2Adapter.CallOpts, arg0, arg1)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCaller) TransferFrom(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _Yearnv2Adapter.contract.Call(opts, &out, "transferFrom", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.TransferFrom(&_Yearnv2Adapter.CallOpts, arg0, arg1, arg2)
}

// TransferFrom is a free data retrieval call binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) pure returns(bool)
func (_Yearnv2Adapter *Yearnv2AdapterCallerSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _Yearnv2Adapter.Contract.TransferFrom(&_Yearnv2Adapter.CallOpts, arg0, arg1, arg2)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "deposit", amount, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Deposit(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit(&_Yearnv2Adapter.TransactOpts, amount, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Deposit(amount *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit(&_Yearnv2Adapter.TransactOpts, amount, arg1)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Deposit0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "deposit0", amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit0(&_Yearnv2Adapter.TransactOpts, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Deposit0(amount *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit0(&_Yearnv2Adapter.TransactOpts, amount)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256 shares)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256 shares)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Deposit1() (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit1(&_Yearnv2Adapter.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256 shares)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Deposit1() (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Deposit1(&_Yearnv2Adapter.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Withdraw(opts *bind.TransactOpts, maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "withdraw", maxShares, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Withdraw(maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw(&_Yearnv2Adapter.TransactOpts, maxShares, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 maxShares, address ) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Withdraw(maxShares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw(&_Yearnv2Adapter.TransactOpts, maxShares, arg1)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Withdraw0(opts *bind.TransactOpts, maxShares *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "withdraw0", maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw0(&_Yearnv2Adapter.TransactOpts, maxShares)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 maxShares) returns(uint256)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Withdraw0(maxShares *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw0(&_Yearnv2Adapter.TransactOpts, maxShares)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Withdraw1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "withdraw1")
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Withdraw1() (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw1(&_Yearnv2Adapter.TransactOpts)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Withdraw1() (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw1(&_Yearnv2Adapter.TransactOpts)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterTransactor) Withdraw2(opts *bind.TransactOpts, maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.contract.Transact(opts, "withdraw2", maxShares, arg1, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterSession) Withdraw2(maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw2(&_Yearnv2Adapter.TransactOpts, maxShares, arg1, maxLoss)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xe63697c8.
//
// Solidity: function withdraw(uint256 maxShares, address , uint256 maxLoss) returns(uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterTransactorSession) Withdraw2(maxShares *big.Int, arg1 common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _Yearnv2Adapter.Contract.Withdraw2(&_Yearnv2Adapter.TransactOpts, maxShares, arg1, maxLoss)
}

// Yearnv2AdapterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yearnv2Adapter contract.
type Yearnv2AdapterApprovalIterator struct {
	Event *Yearnv2AdapterApproval // Event containing the contract specifics and raw log

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
func (it *Yearnv2AdapterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yearnv2AdapterApproval)
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
		it.Event = new(Yearnv2AdapterApproval)
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
func (it *Yearnv2AdapterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yearnv2AdapterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yearnv2AdapterApproval represents a Approval event raised by the Yearnv2Adapter contract.
type Yearnv2AdapterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Yearnv2AdapterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yearnv2Adapter.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Yearnv2AdapterApprovalIterator{contract: _Yearnv2Adapter.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Yearnv2AdapterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yearnv2Adapter.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yearnv2AdapterApproval)
				if err := _Yearnv2Adapter.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) ParseApproval(log types.Log) (*Yearnv2AdapterApproval, error) {
	event := new(Yearnv2AdapterApproval)
	if err := _Yearnv2Adapter.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Yearnv2AdapterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yearnv2Adapter contract.
type Yearnv2AdapterTransferIterator struct {
	Event *Yearnv2AdapterTransfer // Event containing the contract specifics and raw log

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
func (it *Yearnv2AdapterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Yearnv2AdapterTransfer)
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
		it.Event = new(Yearnv2AdapterTransfer)
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
func (it *Yearnv2AdapterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Yearnv2AdapterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Yearnv2AdapterTransfer represents a Transfer event raised by the Yearnv2Adapter contract.
type Yearnv2AdapterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Yearnv2AdapterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yearnv2Adapter.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Yearnv2AdapterTransferIterator{contract: _Yearnv2Adapter.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Yearnv2AdapterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yearnv2Adapter.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Yearnv2AdapterTransfer)
				if err := _Yearnv2Adapter.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yearnv2Adapter *Yearnv2AdapterFilterer) ParseTransfer(log types.Log) (*Yearnv2AdapterTransfer, error) {
	event := new(Yearnv2AdapterTransfer)
	if err := _Yearnv2Adapter.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
