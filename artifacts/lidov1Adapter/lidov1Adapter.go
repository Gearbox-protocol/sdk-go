// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lidov1Adapter

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

// Lidov1AdapterMetaData contains all meta data concerning the Lidov1Adapter contract.
var Lidov1AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lidoGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsOverException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"NewLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acl\",\"outputs\":[{\"internalType\":\"contractIACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"setLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Lidov1AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use Lidov1AdapterMetaData.ABI instead.
var Lidov1AdapterABI = Lidov1AdapterMetaData.ABI

// Lidov1Adapter is an auto generated Go binding around an Ethereum contract.
type Lidov1Adapter struct {
	Lidov1AdapterCaller     // Read-only binding to the contract
	Lidov1AdapterTransactor // Write-only binding to the contract
	Lidov1AdapterFilterer   // Log filterer for contract events
}

// Lidov1AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Lidov1AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Lidov1AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Lidov1AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Lidov1AdapterSession struct {
	Contract     *Lidov1Adapter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Lidov1AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Lidov1AdapterCallerSession struct {
	Contract *Lidov1AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Lidov1AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Lidov1AdapterTransactorSession struct {
	Contract     *Lidov1AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Lidov1AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Lidov1AdapterRaw struct {
	Contract *Lidov1Adapter // Generic contract binding to access the raw methods on
}

// Lidov1AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Lidov1AdapterCallerRaw struct {
	Contract *Lidov1AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// Lidov1AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Lidov1AdapterTransactorRaw struct {
	Contract *Lidov1AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidov1Adapter creates a new instance of Lidov1Adapter, bound to a specific deployed contract.
func NewLidov1Adapter(address common.Address, backend bind.ContractBackend) (*Lidov1Adapter, error) {
	contract, err := bindLidov1Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lidov1Adapter{Lidov1AdapterCaller: Lidov1AdapterCaller{contract: contract}, Lidov1AdapterTransactor: Lidov1AdapterTransactor{contract: contract}, Lidov1AdapterFilterer: Lidov1AdapterFilterer{contract: contract}}, nil
}

// NewLidov1AdapterCaller creates a new read-only instance of Lidov1Adapter, bound to a specific deployed contract.
func NewLidov1AdapterCaller(address common.Address, caller bind.ContractCaller) (*Lidov1AdapterCaller, error) {
	contract, err := bindLidov1Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterCaller{contract: contract}, nil
}

// NewLidov1AdapterTransactor creates a new write-only instance of Lidov1Adapter, bound to a specific deployed contract.
func NewLidov1AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*Lidov1AdapterTransactor, error) {
	contract, err := bindLidov1Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterTransactor{contract: contract}, nil
}

// NewLidov1AdapterFilterer creates a new log filterer instance of Lidov1Adapter, bound to a specific deployed contract.
func NewLidov1AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*Lidov1AdapterFilterer, error) {
	contract, err := bindLidov1Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterFilterer{contract: contract}, nil
}

// bindLidov1Adapter binds a generic wrapper to an already deployed contract.
func bindLidov1Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Lidov1AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lidov1Adapter *Lidov1AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lidov1Adapter.Contract.Lidov1AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lidov1Adapter *Lidov1AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Lidov1AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lidov1Adapter *Lidov1AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Lidov1AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lidov1Adapter *Lidov1AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lidov1Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lidov1Adapter *Lidov1AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lidov1Adapter *Lidov1AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "_acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) Acl() (common.Address, error) {
	return _Lidov1Adapter.Contract.Acl(&_Lidov1Adapter.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Acl() (common.Address, error) {
	return _Lidov1Adapter.Contract.Acl(&_Lidov1Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterSession) GearboxAdapterType() (uint8, error) {
	return _Lidov1Adapter.Contract.GearboxAdapterType(&_Lidov1Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _Lidov1Adapter.Contract.GearboxAdapterType(&_Lidov1Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _Lidov1Adapter.Contract.GearboxAdapterVersion(&_Lidov1Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Lidov1Adapter.Contract.GearboxAdapterVersion(&_Lidov1Adapter.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.Allowance(&_Lidov1Adapter.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.Allowance(&_Lidov1Adapter.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.BalanceOf(&_Lidov1Adapter.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.BalanceOf(&_Lidov1Adapter.CallOpts, _account)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) CreditFacade() (common.Address, error) {
	return _Lidov1Adapter.Contract.CreditFacade(&_Lidov1Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) CreditFacade() (common.Address, error) {
	return _Lidov1Adapter.Contract.CreditFacade(&_Lidov1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) CreditManager() (common.Address, error) {
	return _Lidov1Adapter.Contract.CreditManager(&_Lidov1Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) CreditManager() (common.Address, error) {
	return _Lidov1Adapter.Contract.CreditManager(&_Lidov1Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterSession) Decimals() (uint8, error) {
	return _Lidov1Adapter.Contract.Decimals(&_Lidov1Adapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Decimals() (uint8, error) {
	return _Lidov1Adapter.Contract.Decimals(&_Lidov1Adapter.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterSession) GetFee() (uint16, error) {
	return _Lidov1Adapter.Contract.GetFee(&_Lidov1Adapter.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GetFee() (uint16, error) {
	return _Lidov1Adapter.Contract.GetFee(&_Lidov1Adapter.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetPooledEthByShares(&_Lidov1Adapter.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetPooledEthByShares(&_Lidov1Adapter.CallOpts, _sharesAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetSharesByPooledEth(&_Lidov1Adapter.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetSharesByPooledEth(&_Lidov1Adapter.CallOpts, _ethAmount)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetTotalPooledEther(&_Lidov1Adapter.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetTotalPooledEther(&_Lidov1Adapter.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) GetTotalShares() (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetTotalShares(&_Lidov1Adapter.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) GetTotalShares() (*big.Int, error) {
	return _Lidov1Adapter.Contract.GetTotalShares(&_Lidov1Adapter.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) Limit() (*big.Int, error) {
	return _Lidov1Adapter.Contract.Limit(&_Lidov1Adapter.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Limit() (*big.Int, error) {
	return _Lidov1Adapter.Contract.Limit(&_Lidov1Adapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterSession) Name() (string, error) {
	return _Lidov1Adapter.Contract.Name(&_Lidov1Adapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Name() (string, error) {
	return _Lidov1Adapter.Contract.Name(&_Lidov1Adapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lidov1Adapter *Lidov1AdapterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lidov1Adapter *Lidov1AdapterSession) Paused() (bool, error) {
	return _Lidov1Adapter.Contract.Paused(&_Lidov1Adapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Paused() (bool, error) {
	return _Lidov1Adapter.Contract.Paused(&_Lidov1Adapter.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.SharesOf(&_Lidov1Adapter.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Adapter.Contract.SharesOf(&_Lidov1Adapter.CallOpts, _account)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) StETH() (common.Address, error) {
	return _Lidov1Adapter.Contract.StETH(&_Lidov1Adapter.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) StETH() (common.Address, error) {
	return _Lidov1Adapter.Contract.StETH(&_Lidov1Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterSession) Symbol() (string, error) {
	return _Lidov1Adapter.Contract.Symbol(&_Lidov1Adapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Symbol() (string, error) {
	return _Lidov1Adapter.Contract.Symbol(&_Lidov1Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) TargetContract() (common.Address, error) {
	return _Lidov1Adapter.Contract.TargetContract(&_Lidov1Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) TargetContract() (common.Address, error) {
	return _Lidov1Adapter.Contract.TargetContract(&_Lidov1Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterSession) TotalSupply() (*big.Int, error) {
	return _Lidov1Adapter.Contract.TotalSupply(&_Lidov1Adapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _Lidov1Adapter.Contract.TotalSupply(&_Lidov1Adapter.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) Treasury() (common.Address, error) {
	return _Lidov1Adapter.Contract.Treasury(&_Lidov1Adapter.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Treasury() (common.Address, error) {
	return _Lidov1Adapter.Contract.Treasury(&_Lidov1Adapter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Adapter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterSession) Weth() (common.Address, error) {
	return _Lidov1Adapter.Contract.Weth(&_Lidov1Adapter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Adapter *Lidov1AdapterCallerSession) Weth() (common.Address, error) {
	return _Lidov1Adapter.Contract.Weth(&_Lidov1Adapter.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lidov1Adapter *Lidov1AdapterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Adapter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lidov1Adapter *Lidov1AdapterSession) Pause() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Pause(&_Lidov1Adapter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lidov1Adapter *Lidov1AdapterTransactorSession) Pause() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Pause(&_Lidov1Adapter.TransactOpts)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_Lidov1Adapter *Lidov1AdapterTransactor) SetLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.contract.Transact(opts, "setLimit", _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_Lidov1Adapter *Lidov1AdapterSession) SetLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.SetLimit(&_Lidov1Adapter.TransactOpts, _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0x27ea6f2b.
//
// Solidity: function setLimit(uint256 _limit) returns()
func (_Lidov1Adapter *Lidov1AdapterTransactorSession) SetLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.SetLimit(&_Lidov1Adapter.TransactOpts, _limit)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 amount) returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterTransactor) Submit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.contract.Transact(opts, "submit", amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 amount) returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterSession) Submit(amount *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Submit(&_Lidov1Adapter.TransactOpts, amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 amount) returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterTransactorSession) Submit(amount *big.Int) (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Submit(&_Lidov1Adapter.TransactOpts, amount)
}

// SubmitAll is a paid mutator transaction binding the contract method 0x30bebac9.
//
// Solidity: function submitAll() returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterTransactor) SubmitAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Adapter.contract.Transact(opts, "submitAll")
}

// SubmitAll is a paid mutator transaction binding the contract method 0x30bebac9.
//
// Solidity: function submitAll() returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterSession) SubmitAll() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.SubmitAll(&_Lidov1Adapter.TransactOpts)
}

// SubmitAll is a paid mutator transaction binding the contract method 0x30bebac9.
//
// Solidity: function submitAll() returns(uint256 result)
func (_Lidov1Adapter *Lidov1AdapterTransactorSession) SubmitAll() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.SubmitAll(&_Lidov1Adapter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lidov1Adapter *Lidov1AdapterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Adapter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lidov1Adapter *Lidov1AdapterSession) Unpause() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Unpause(&_Lidov1Adapter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lidov1Adapter *Lidov1AdapterTransactorSession) Unpause() (*types.Transaction, error) {
	return _Lidov1Adapter.Contract.Unpause(&_Lidov1Adapter.TransactOpts)
}

// Lidov1AdapterNewLimitIterator is returned from FilterNewLimit and is used to iterate over the raw logs and unpacked data for NewLimit events raised by the Lidov1Adapter contract.
type Lidov1AdapterNewLimitIterator struct {
	Event *Lidov1AdapterNewLimit // Event containing the contract specifics and raw log

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
func (it *Lidov1AdapterNewLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lidov1AdapterNewLimit)
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
		it.Event = new(Lidov1AdapterNewLimit)
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
func (it *Lidov1AdapterNewLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lidov1AdapterNewLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lidov1AdapterNewLimit represents a NewLimit event raised by the Lidov1Adapter contract.
type Lidov1AdapterNewLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewLimit is a free log retrieval operation binding the contract event 0xe1e1c8251499b303aefb01cf84a5ce22a95911c20ce2f3f5ae670441a6353d82.
//
// Solidity: event NewLimit(uint256 _limit)
func (_Lidov1Adapter *Lidov1AdapterFilterer) FilterNewLimit(opts *bind.FilterOpts) (*Lidov1AdapterNewLimitIterator, error) {

	logs, sub, err := _Lidov1Adapter.contract.FilterLogs(opts, "NewLimit")
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterNewLimitIterator{contract: _Lidov1Adapter.contract, event: "NewLimit", logs: logs, sub: sub}, nil
}

// WatchNewLimit is a free log subscription operation binding the contract event 0xe1e1c8251499b303aefb01cf84a5ce22a95911c20ce2f3f5ae670441a6353d82.
//
// Solidity: event NewLimit(uint256 _limit)
func (_Lidov1Adapter *Lidov1AdapterFilterer) WatchNewLimit(opts *bind.WatchOpts, sink chan<- *Lidov1AdapterNewLimit) (event.Subscription, error) {

	logs, sub, err := _Lidov1Adapter.contract.WatchLogs(opts, "NewLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lidov1AdapterNewLimit)
				if err := _Lidov1Adapter.contract.UnpackLog(event, "NewLimit", log); err != nil {
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

// ParseNewLimit is a log parse operation binding the contract event 0xe1e1c8251499b303aefb01cf84a5ce22a95911c20ce2f3f5ae670441a6353d82.
//
// Solidity: event NewLimit(uint256 _limit)
func (_Lidov1Adapter *Lidov1AdapterFilterer) ParseNewLimit(log types.Log) (*Lidov1AdapterNewLimit, error) {
	event := new(Lidov1AdapterNewLimit)
	if err := _Lidov1Adapter.contract.UnpackLog(event, "NewLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lidov1AdapterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lidov1Adapter contract.
type Lidov1AdapterPausedIterator struct {
	Event *Lidov1AdapterPaused // Event containing the contract specifics and raw log

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
func (it *Lidov1AdapterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lidov1AdapterPaused)
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
		it.Event = new(Lidov1AdapterPaused)
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
func (it *Lidov1AdapterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lidov1AdapterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lidov1AdapterPaused represents a Paused event raised by the Lidov1Adapter contract.
type Lidov1AdapterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lidov1Adapter *Lidov1AdapterFilterer) FilterPaused(opts *bind.FilterOpts) (*Lidov1AdapterPausedIterator, error) {

	logs, sub, err := _Lidov1Adapter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterPausedIterator{contract: _Lidov1Adapter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lidov1Adapter *Lidov1AdapterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Lidov1AdapterPaused) (event.Subscription, error) {

	logs, sub, err := _Lidov1Adapter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lidov1AdapterPaused)
				if err := _Lidov1Adapter.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Lidov1Adapter *Lidov1AdapterFilterer) ParsePaused(log types.Log) (*Lidov1AdapterPaused, error) {
	event := new(Lidov1AdapterPaused)
	if err := _Lidov1Adapter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lidov1AdapterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lidov1Adapter contract.
type Lidov1AdapterUnpausedIterator struct {
	Event *Lidov1AdapterUnpaused // Event containing the contract specifics and raw log

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
func (it *Lidov1AdapterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lidov1AdapterUnpaused)
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
		it.Event = new(Lidov1AdapterUnpaused)
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
func (it *Lidov1AdapterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lidov1AdapterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lidov1AdapterUnpaused represents a Unpaused event raised by the Lidov1Adapter contract.
type Lidov1AdapterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lidov1Adapter *Lidov1AdapterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Lidov1AdapterUnpausedIterator, error) {

	logs, sub, err := _Lidov1Adapter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Lidov1AdapterUnpausedIterator{contract: _Lidov1Adapter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lidov1Adapter *Lidov1AdapterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Lidov1AdapterUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lidov1Adapter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lidov1AdapterUnpaused)
				if err := _Lidov1Adapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Lidov1Adapter *Lidov1AdapterFilterer) ParseUnpaused(log types.Log) (*Lidov1AdapterUnpaused, error) {
	event := new(Lidov1AdapterUnpaused)
	if err := _Lidov1Adapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
