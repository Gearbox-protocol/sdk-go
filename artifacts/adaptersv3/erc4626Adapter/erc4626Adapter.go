// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc4626Adapter

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

// Erc4626AdapterMetaData contains all meta data concerning the Erc4626Adapter contract.
var Erc4626AdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assetMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Erc4626AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc4626AdapterMetaData.ABI instead.
var Erc4626AdapterABI = Erc4626AdapterMetaData.ABI

// Erc4626Adapter is an auto generated Go binding around an Ethereum contract.
type Erc4626Adapter struct {
	Erc4626AdapterCaller     // Read-only binding to the contract
	Erc4626AdapterTransactor // Write-only binding to the contract
	Erc4626AdapterFilterer   // Log filterer for contract events
}

// Erc4626AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc4626AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc4626AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc4626AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc4626AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc4626AdapterSession struct {
	Contract     *Erc4626Adapter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc4626AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc4626AdapterCallerSession struct {
	Contract *Erc4626AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc4626AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc4626AdapterTransactorSession struct {
	Contract     *Erc4626AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc4626AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc4626AdapterRaw struct {
	Contract *Erc4626Adapter // Generic contract binding to access the raw methods on
}

// Erc4626AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc4626AdapterCallerRaw struct {
	Contract *Erc4626AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// Erc4626AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc4626AdapterTransactorRaw struct {
	Contract *Erc4626AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc4626Adapter creates a new instance of Erc4626Adapter, bound to a specific deployed contract.
func NewErc4626Adapter(address common.Address, backend bind.ContractBackend) (*Erc4626Adapter, error) {
	contract, err := bindErc4626Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc4626Adapter{Erc4626AdapterCaller: Erc4626AdapterCaller{contract: contract}, Erc4626AdapterTransactor: Erc4626AdapterTransactor{contract: contract}, Erc4626AdapterFilterer: Erc4626AdapterFilterer{contract: contract}}, nil
}

// NewErc4626AdapterCaller creates a new read-only instance of Erc4626Adapter, bound to a specific deployed contract.
func NewErc4626AdapterCaller(address common.Address, caller bind.ContractCaller) (*Erc4626AdapterCaller, error) {
	contract, err := bindErc4626Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4626AdapterCaller{contract: contract}, nil
}

// NewErc4626AdapterTransactor creates a new write-only instance of Erc4626Adapter, bound to a specific deployed contract.
func NewErc4626AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc4626AdapterTransactor, error) {
	contract, err := bindErc4626Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc4626AdapterTransactor{contract: contract}, nil
}

// NewErc4626AdapterFilterer creates a new log filterer instance of Erc4626Adapter, bound to a specific deployed contract.
func NewErc4626AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc4626AdapterFilterer, error) {
	contract, err := bindErc4626Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc4626AdapterFilterer{contract: contract}, nil
}

// bindErc4626Adapter binds a generic wrapper to an already deployed contract.
func bindErc4626Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc4626AdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4626Adapter *Erc4626AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4626Adapter.Contract.Erc4626AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4626Adapter *Erc4626AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Erc4626AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4626Adapter *Erc4626AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Erc4626AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc4626Adapter *Erc4626AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc4626Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc4626Adapter *Erc4626AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc4626Adapter *Erc4626AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Erc4626Adapter *Erc4626AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Erc4626Adapter *Erc4626AdapterSession) GearboxAdapterType() (uint8, error) {
	return _Erc4626Adapter.Contract.GearboxAdapterType(&_Erc4626Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _Erc4626Adapter.Contract.GearboxAdapterType(&_Erc4626Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Erc4626Adapter *Erc4626AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Erc4626Adapter *Erc4626AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _Erc4626Adapter.Contract.GearboxAdapterVersion(&_Erc4626Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Erc4626Adapter.Contract.GearboxAdapterVersion(&_Erc4626Adapter.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterSession) Acl() (common.Address, error) {
	return _Erc4626Adapter.Contract.Acl(&_Erc4626Adapter.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) Acl() (common.Address, error) {
	return _Erc4626Adapter.Contract.Acl(&_Erc4626Adapter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterSession) AddressProvider() (common.Address, error) {
	return _Erc4626Adapter.Contract.AddressProvider(&_Erc4626Adapter.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) AddressProvider() (common.Address, error) {
	return _Erc4626Adapter.Contract.AddressProvider(&_Erc4626Adapter.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterSession) Asset() (common.Address, error) {
	return _Erc4626Adapter.Contract.Asset(&_Erc4626Adapter.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) Asset() (common.Address, error) {
	return _Erc4626Adapter.Contract.Asset(&_Erc4626Adapter.CallOpts)
}

// AssetMask is a free data retrieval call binding the contract method 0xd823dcd5.
//
// Solidity: function assetMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterCaller) AssetMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "assetMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetMask is a free data retrieval call binding the contract method 0xd823dcd5.
//
// Solidity: function assetMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterSession) AssetMask() (*big.Int, error) {
	return _Erc4626Adapter.Contract.AssetMask(&_Erc4626Adapter.CallOpts)
}

// AssetMask is a free data retrieval call binding the contract method 0xd823dcd5.
//
// Solidity: function assetMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) AssetMask() (*big.Int, error) {
	return _Erc4626Adapter.Contract.AssetMask(&_Erc4626Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterSession) CreditManager() (common.Address, error) {
	return _Erc4626Adapter.Contract.CreditManager(&_Erc4626Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) CreditManager() (common.Address, error) {
	return _Erc4626Adapter.Contract.CreditManager(&_Erc4626Adapter.CallOpts)
}

// SharesMask is a free data retrieval call binding the contract method 0x1a0a59a1.
//
// Solidity: function sharesMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterCaller) SharesMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "sharesMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesMask is a free data retrieval call binding the contract method 0x1a0a59a1.
//
// Solidity: function sharesMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterSession) SharesMask() (*big.Int, error) {
	return _Erc4626Adapter.Contract.SharesMask(&_Erc4626Adapter.CallOpts)
}

// SharesMask is a free data retrieval call binding the contract method 0x1a0a59a1.
//
// Solidity: function sharesMask() view returns(uint256)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) SharesMask() (*big.Int, error) {
	return _Erc4626Adapter.Contract.SharesMask(&_Erc4626Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc4626Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterSession) TargetContract() (common.Address, error) {
	return _Erc4626Adapter.Contract.TargetContract(&_Erc4626Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Erc4626Adapter *Erc4626AdapterCallerSession) TargetContract() (common.Address, error) {
	return _Erc4626Adapter.Contract.TargetContract(&_Erc4626Adapter.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "deposit", assets, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) Deposit(assets *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Deposit(&_Erc4626Adapter.TransactOpts, assets, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) Deposit(assets *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Deposit(&_Erc4626Adapter.TransactOpts, assets, arg1)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) DepositDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "depositDiff", leftoverAmount)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) DepositDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.DepositDiff(&_Erc4626Adapter.TransactOpts, leftoverAmount)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) DepositDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.DepositDiff(&_Erc4626Adapter.TransactOpts, leftoverAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "mint", shares, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) Mint(shares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Mint(&_Erc4626Adapter.TransactOpts, shares, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) Mint(shares *big.Int, arg1 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Mint(&_Erc4626Adapter.TransactOpts, shares, arg1)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "redeem", shares, arg1, arg2)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) Redeem(shares *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Redeem(&_Erc4626Adapter.TransactOpts, shares, arg1, arg2)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) Redeem(shares *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Redeem(&_Erc4626Adapter.TransactOpts, shares, arg1, arg2)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) RedeemDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "redeemDiff", leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.RedeemDiff(&_Erc4626Adapter.TransactOpts, leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.RedeemDiff(&_Erc4626Adapter.TransactOpts, leftoverAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.contract.Transact(opts, "withdraw", assets, arg1, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterSession) Withdraw(assets *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Withdraw(&_Erc4626Adapter.TransactOpts, assets, arg1, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address , address ) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Erc4626Adapter *Erc4626AdapterTransactorSession) Withdraw(assets *big.Int, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Erc4626Adapter.Contract.Withdraw(&_Erc4626Adapter.TransactOpts, assets, arg1, arg2)
}
