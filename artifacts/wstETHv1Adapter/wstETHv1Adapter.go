// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wstETHv1Adapter

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

// WstETHv1AdapterMetaData contains all meta data concerning the WstETHv1Adapter contract.
var WstETHv1AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wstETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenIsNotAddedToCreditManagerException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"getStETHByWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"getWstETHByStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stEthPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerStEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unwrapAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WstETHv1AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use WstETHv1AdapterMetaData.ABI instead.
var WstETHv1AdapterABI = WstETHv1AdapterMetaData.ABI

// WstETHv1Adapter is an auto generated Go binding around an Ethereum contract.
type WstETHv1Adapter struct {
	WstETHv1AdapterCaller     // Read-only binding to the contract
	WstETHv1AdapterTransactor // Write-only binding to the contract
	WstETHv1AdapterFilterer   // Log filterer for contract events
}

// WstETHv1AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type WstETHv1AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WstETHv1AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WstETHv1AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WstETHv1AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WstETHv1AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WstETHv1AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WstETHv1AdapterSession struct {
	Contract     *WstETHv1Adapter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WstETHv1AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WstETHv1AdapterCallerSession struct {
	Contract *WstETHv1AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WstETHv1AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WstETHv1AdapterTransactorSession struct {
	Contract     *WstETHv1AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WstETHv1AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type WstETHv1AdapterRaw struct {
	Contract *WstETHv1Adapter // Generic contract binding to access the raw methods on
}

// WstETHv1AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WstETHv1AdapterCallerRaw struct {
	Contract *WstETHv1AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// WstETHv1AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WstETHv1AdapterTransactorRaw struct {
	Contract *WstETHv1AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWstETHv1Adapter creates a new instance of WstETHv1Adapter, bound to a specific deployed contract.
func NewWstETHv1Adapter(address common.Address, backend bind.ContractBackend) (*WstETHv1Adapter, error) {
	contract, err := bindWstETHv1Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WstETHv1Adapter{WstETHv1AdapterCaller: WstETHv1AdapterCaller{contract: contract}, WstETHv1AdapterTransactor: WstETHv1AdapterTransactor{contract: contract}, WstETHv1AdapterFilterer: WstETHv1AdapterFilterer{contract: contract}}, nil
}

// NewWstETHv1AdapterCaller creates a new read-only instance of WstETHv1Adapter, bound to a specific deployed contract.
func NewWstETHv1AdapterCaller(address common.Address, caller bind.ContractCaller) (*WstETHv1AdapterCaller, error) {
	contract, err := bindWstETHv1Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WstETHv1AdapterCaller{contract: contract}, nil
}

// NewWstETHv1AdapterTransactor creates a new write-only instance of WstETHv1Adapter, bound to a specific deployed contract.
func NewWstETHv1AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*WstETHv1AdapterTransactor, error) {
	contract, err := bindWstETHv1Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WstETHv1AdapterTransactor{contract: contract}, nil
}

// NewWstETHv1AdapterFilterer creates a new log filterer instance of WstETHv1Adapter, bound to a specific deployed contract.
func NewWstETHv1AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*WstETHv1AdapterFilterer, error) {
	contract, err := bindWstETHv1Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WstETHv1AdapterFilterer{contract: contract}, nil
}

// bindWstETHv1Adapter binds a generic wrapper to an already deployed contract.
func bindWstETHv1Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WstETHv1AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WstETHv1Adapter *WstETHv1AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WstETHv1Adapter.Contract.WstETHv1AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WstETHv1Adapter *WstETHv1AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.WstETHv1AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WstETHv1Adapter *WstETHv1AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.WstETHv1AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WstETHv1Adapter *WstETHv1AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WstETHv1Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WstETHv1Adapter *WstETHv1AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WstETHv1Adapter *WstETHv1AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterSession) GearboxAdapterType() (uint8, error) {
	return _WstETHv1Adapter.Contract.GearboxAdapterType(&_WstETHv1Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _WstETHv1Adapter.Contract.GearboxAdapterType(&_WstETHv1Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_WstETHv1Adapter *WstETHv1AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _WstETHv1Adapter.Contract.GearboxAdapterVersion(&_WstETHv1Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _WstETHv1Adapter.Contract.GearboxAdapterVersion(&_WstETHv1Adapter.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.Allowance(&_WstETHv1Adapter.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.Allowance(&_WstETHv1Adapter.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.BalanceOf(&_WstETHv1Adapter.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.BalanceOf(&_WstETHv1Adapter.CallOpts, _account)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterSession) CreditFacade() (common.Address, error) {
	return _WstETHv1Adapter.Contract.CreditFacade(&_WstETHv1Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) CreditFacade() (common.Address, error) {
	return _WstETHv1Adapter.Contract.CreditFacade(&_WstETHv1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterSession) CreditManager() (common.Address, error) {
	return _WstETHv1Adapter.Contract.CreditManager(&_WstETHv1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) CreditManager() (common.Address, error) {
	return _WstETHv1Adapter.Contract.CreditManager(&_WstETHv1Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Decimals() (uint8, error) {
	return _WstETHv1Adapter.Contract.Decimals(&_WstETHv1Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) Decimals() (uint8, error) {
	return _WstETHv1Adapter.Contract.Decimals(&_WstETHv1Adapter.CallOpts)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) GetStETHByWstETH(opts *bind.CallOpts, _wstETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "getStETHByWstETH", _wstETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.GetStETHByWstETH(&_WstETHv1Adapter.CallOpts, _wstETHAmount)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.GetStETHByWstETH(&_WstETHv1Adapter.CallOpts, _wstETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) GetWstETHByStETH(opts *bind.CallOpts, _stETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "getWstETHByStETH", _stETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.GetWstETHByStETH(&_WstETHv1Adapter.CallOpts, _stETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _WstETHv1Adapter.Contract.GetWstETHByStETH(&_WstETHv1Adapter.CallOpts, _stETHAmount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Name() (string, error) {
	return _WstETHv1Adapter.Contract.Name(&_WstETHv1Adapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) Name() (string, error) {
	return _WstETHv1Adapter.Contract.Name(&_WstETHv1Adapter.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterSession) StETH() (common.Address, error) {
	return _WstETHv1Adapter.Contract.StETH(&_WstETHv1Adapter.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) StETH() (common.Address, error) {
	return _WstETHv1Adapter.Contract.StETH(&_WstETHv1Adapter.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) StEthPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "stEthPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) StEthPerToken() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.StEthPerToken(&_WstETHv1Adapter.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) StEthPerToken() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.StEthPerToken(&_WstETHv1Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Symbol() (string, error) {
	return _WstETHv1Adapter.Contract.Symbol(&_WstETHv1Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) Symbol() (string, error) {
	return _WstETHv1Adapter.Contract.Symbol(&_WstETHv1Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterSession) TargetContract() (common.Address, error) {
	return _WstETHv1Adapter.Contract.TargetContract(&_WstETHv1Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) TargetContract() (common.Address, error) {
	return _WstETHv1Adapter.Contract.TargetContract(&_WstETHv1Adapter.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) TokensPerStEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "tokensPerStEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) TokensPerStEth() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.TokensPerStEth(&_WstETHv1Adapter.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) TokensPerStEth() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.TokensPerStEth(&_WstETHv1Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WstETHv1Adapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) TotalSupply() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.TotalSupply(&_WstETHv1Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _WstETHv1Adapter.Contract.TotalSupply(&_WstETHv1Adapter.CallOpts)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactor) Unwrap(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.contract.Transact(opts, "unwrap", _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.Unwrap(&_WstETHv1Adapter.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactorSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.Unwrap(&_WstETHv1Adapter.TransactOpts, _wstETHAmount)
}

// UnwrapAll is a paid mutator transaction binding the contract method 0x4982e3b7.
//
// Solidity: function unwrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactor) UnwrapAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WstETHv1Adapter.contract.Transact(opts, "unwrapAll")
}

// UnwrapAll is a paid mutator transaction binding the contract method 0x4982e3b7.
//
// Solidity: function unwrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) UnwrapAll() (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.UnwrapAll(&_WstETHv1Adapter.TransactOpts)
}

// UnwrapAll is a paid mutator transaction binding the contract method 0x4982e3b7.
//
// Solidity: function unwrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactorSession) UnwrapAll() (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.UnwrapAll(&_WstETHv1Adapter.TransactOpts)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactor) Wrap(opts *bind.TransactOpts, _stETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.contract.Transact(opts, "wrap", _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.Wrap(&_WstETHv1Adapter.TransactOpts, _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactorSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.Wrap(&_WstETHv1Adapter.TransactOpts, _stETHAmount)
}

// WrapAll is a paid mutator transaction binding the contract method 0x4c84c1c8.
//
// Solidity: function wrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactor) WrapAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WstETHv1Adapter.contract.Transact(opts, "wrapAll")
}

// WrapAll is a paid mutator transaction binding the contract method 0x4c84c1c8.
//
// Solidity: function wrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterSession) WrapAll() (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.WrapAll(&_WstETHv1Adapter.TransactOpts)
}

// WrapAll is a paid mutator transaction binding the contract method 0x4c84c1c8.
//
// Solidity: function wrapAll() returns(uint256)
func (_WstETHv1Adapter *WstETHv1AdapterTransactorSession) WrapAll() (*types.Transaction, error) {
	return _WstETHv1Adapter.Contract.WrapAll(&_WstETHv1Adapter.TransactOpts)
}
