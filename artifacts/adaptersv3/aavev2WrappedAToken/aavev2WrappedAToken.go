// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev2WrappedAToken

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

// Aavev2WrappedATokenMetaData contains all meta data concerning the Aavev2WrappedAToken contract.
var Aavev2WrappedATokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_waToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aTokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositDiff\",\"inputs\":[{\"name\":\"leftoverAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositDiffUnderlying\",\"inputs\":[{\"name\":\"leftoverAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositUnderlying\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"waTokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDiff\",\"inputs\":[{\"name\":\"leftoverShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDiffUnderlying\",\"inputs\":[{\"name\":\"leftoverShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawUnderlying\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Aavev2WrappedATokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Aavev2WrappedATokenMetaData.ABI instead.
var Aavev2WrappedATokenABI = Aavev2WrappedATokenMetaData.ABI

// Aavev2WrappedAToken is an auto generated Go binding around an Ethereum contract.
type Aavev2WrappedAToken struct {
	Aavev2WrappedATokenCaller     // Read-only binding to the contract
	Aavev2WrappedATokenTransactor // Write-only binding to the contract
	Aavev2WrappedATokenFilterer   // Log filterer for contract events
}

// Aavev2WrappedATokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Aavev2WrappedATokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2WrappedATokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Aavev2WrappedATokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2WrappedATokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Aavev2WrappedATokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aavev2WrappedATokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Aavev2WrappedATokenSession struct {
	Contract     *Aavev2WrappedAToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Aavev2WrappedATokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Aavev2WrappedATokenCallerSession struct {
	Contract *Aavev2WrappedATokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Aavev2WrappedATokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Aavev2WrappedATokenTransactorSession struct {
	Contract     *Aavev2WrappedATokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Aavev2WrappedATokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Aavev2WrappedATokenRaw struct {
	Contract *Aavev2WrappedAToken // Generic contract binding to access the raw methods on
}

// Aavev2WrappedATokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Aavev2WrappedATokenCallerRaw struct {
	Contract *Aavev2WrappedATokenCaller // Generic read-only contract binding to access the raw methods on
}

// Aavev2WrappedATokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Aavev2WrappedATokenTransactorRaw struct {
	Contract *Aavev2WrappedATokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavev2WrappedAToken creates a new instance of Aavev2WrappedAToken, bound to a specific deployed contract.
func NewAavev2WrappedAToken(address common.Address, backend bind.ContractBackend) (*Aavev2WrappedAToken, error) {
	contract, err := bindAavev2WrappedAToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aavev2WrappedAToken{Aavev2WrappedATokenCaller: Aavev2WrappedATokenCaller{contract: contract}, Aavev2WrappedATokenTransactor: Aavev2WrappedATokenTransactor{contract: contract}, Aavev2WrappedATokenFilterer: Aavev2WrappedATokenFilterer{contract: contract}}, nil
}

// NewAavev2WrappedATokenCaller creates a new read-only instance of Aavev2WrappedAToken, bound to a specific deployed contract.
func NewAavev2WrappedATokenCaller(address common.Address, caller bind.ContractCaller) (*Aavev2WrappedATokenCaller, error) {
	contract, err := bindAavev2WrappedAToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Aavev2WrappedATokenCaller{contract: contract}, nil
}

// NewAavev2WrappedATokenTransactor creates a new write-only instance of Aavev2WrappedAToken, bound to a specific deployed contract.
func NewAavev2WrappedATokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Aavev2WrappedATokenTransactor, error) {
	contract, err := bindAavev2WrappedAToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Aavev2WrappedATokenTransactor{contract: contract}, nil
}

// NewAavev2WrappedATokenFilterer creates a new log filterer instance of Aavev2WrappedAToken, bound to a specific deployed contract.
func NewAavev2WrappedATokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Aavev2WrappedATokenFilterer, error) {
	contract, err := bindAavev2WrappedAToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Aavev2WrappedATokenFilterer{contract: contract}, nil
}

// bindAavev2WrappedAToken binds a generic wrapper to an already deployed contract.
func bindAavev2WrappedAToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Aavev2WrappedATokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aavev2WrappedAToken.Contract.Aavev2WrappedATokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Aavev2WrappedATokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Aavev2WrappedATokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aavev2WrappedAToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) GearboxAdapterType() (uint8, error) {
	return _Aavev2WrappedAToken.Contract.GearboxAdapterType(&_Aavev2WrappedAToken.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) GearboxAdapterType() (uint8, error) {
	return _Aavev2WrappedAToken.Contract.GearboxAdapterType(&_Aavev2WrappedAToken.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) GearboxAdapterVersion() (uint16, error) {
	return _Aavev2WrappedAToken.Contract.GearboxAdapterVersion(&_Aavev2WrappedAToken.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Aavev2WrappedAToken.Contract.GearboxAdapterVersion(&_Aavev2WrappedAToken.CallOpts)
}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) AToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "aToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) AToken() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.AToken(&_Aavev2WrappedAToken.CallOpts)
}

// AToken is a free data retrieval call binding the contract method 0xa0c1f15e.
//
// Solidity: function aToken() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) AToken() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.AToken(&_Aavev2WrappedAToken.CallOpts)
}

// ATokenMask is a free data retrieval call binding the contract method 0x57c92566.
//
// Solidity: function aTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) ATokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "aTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATokenMask is a free data retrieval call binding the contract method 0x57c92566.
//
// Solidity: function aTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) ATokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.ATokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// ATokenMask is a free data retrieval call binding the contract method 0x57c92566.
//
// Solidity: function aTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) ATokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.ATokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) Acl() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.Acl(&_Aavev2WrappedAToken.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) Acl() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.Acl(&_Aavev2WrappedAToken.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) AddressProvider() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.AddressProvider(&_Aavev2WrappedAToken.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) AddressProvider() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.AddressProvider(&_Aavev2WrappedAToken.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) CreditManager() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.CreditManager(&_Aavev2WrappedAToken.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) CreditManager() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.CreditManager(&_Aavev2WrappedAToken.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) TargetContract() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.TargetContract(&_Aavev2WrappedAToken.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) TargetContract() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.TargetContract(&_Aavev2WrappedAToken.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) TokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "tokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) TokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.TokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) TokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.TokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) Underlying() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.Underlying(&_Aavev2WrappedAToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) Underlying() (common.Address, error) {
	return _Aavev2WrappedAToken.Contract.Underlying(&_Aavev2WrappedAToken.CallOpts)
}

// WaTokenMask is a free data retrieval call binding the contract method 0x2b4e9ecb.
//
// Solidity: function waTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCaller) WaTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aavev2WrappedAToken.contract.Call(opts, &out, "waTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WaTokenMask is a free data retrieval call binding the contract method 0x2b4e9ecb.
//
// Solidity: function waTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) WaTokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.WaTokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// WaTokenMask is a free data retrieval call binding the contract method 0x2b4e9ecb.
//
// Solidity: function waTokenMask() view returns(uint256)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenCallerSession) WaTokenMask() (*big.Int, error) {
	return _Aavev2WrappedAToken.Contract.WaTokenMask(&_Aavev2WrappedAToken.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "deposit", assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Deposit(&_Aavev2WrappedAToken.TransactOpts, assets)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) Deposit(assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Deposit(&_Aavev2WrappedAToken.TransactOpts, assets)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) DepositDiff(opts *bind.TransactOpts, leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "depositDiff", leftoverAssets)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) DepositDiff(leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositDiff(&_Aavev2WrappedAToken.TransactOpts, leftoverAssets)
}

// DepositDiff is a paid mutator transaction binding the contract method 0x1f4f702e.
//
// Solidity: function depositDiff(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) DepositDiff(leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositDiff(&_Aavev2WrappedAToken.TransactOpts, leftoverAssets)
}

// DepositDiffUnderlying is a paid mutator transaction binding the contract method 0x4f959d63.
//
// Solidity: function depositDiffUnderlying(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) DepositDiffUnderlying(opts *bind.TransactOpts, leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "depositDiffUnderlying", leftoverAssets)
}

// DepositDiffUnderlying is a paid mutator transaction binding the contract method 0x4f959d63.
//
// Solidity: function depositDiffUnderlying(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) DepositDiffUnderlying(leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositDiffUnderlying(&_Aavev2WrappedAToken.TransactOpts, leftoverAssets)
}

// DepositDiffUnderlying is a paid mutator transaction binding the contract method 0x4f959d63.
//
// Solidity: function depositDiffUnderlying(uint256 leftoverAssets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) DepositDiffUnderlying(leftoverAssets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositDiffUnderlying(&_Aavev2WrappedAToken.TransactOpts, leftoverAssets)
}

// DepositUnderlying is a paid mutator transaction binding the contract method 0xb9f5be41.
//
// Solidity: function depositUnderlying(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) DepositUnderlying(opts *bind.TransactOpts, assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "depositUnderlying", assets)
}

// DepositUnderlying is a paid mutator transaction binding the contract method 0xb9f5be41.
//
// Solidity: function depositUnderlying(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) DepositUnderlying(assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositUnderlying(&_Aavev2WrappedAToken.TransactOpts, assets)
}

// DepositUnderlying is a paid mutator transaction binding the contract method 0xb9f5be41.
//
// Solidity: function depositUnderlying(uint256 assets) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) DepositUnderlying(assets *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.DepositUnderlying(&_Aavev2WrappedAToken.TransactOpts, assets)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "withdraw", shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) Withdraw(shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Withdraw(&_Aavev2WrappedAToken.TransactOpts, shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) Withdraw(shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.Withdraw(&_Aavev2WrappedAToken.TransactOpts, shares)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x67150887.
//
// Solidity: function withdrawDiff(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) WithdrawDiff(opts *bind.TransactOpts, leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "withdrawDiff", leftoverShares)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x67150887.
//
// Solidity: function withdrawDiff(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) WithdrawDiff(leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawDiff(&_Aavev2WrappedAToken.TransactOpts, leftoverShares)
}

// WithdrawDiff is a paid mutator transaction binding the contract method 0x67150887.
//
// Solidity: function withdrawDiff(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) WithdrawDiff(leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawDiff(&_Aavev2WrappedAToken.TransactOpts, leftoverShares)
}

// WithdrawDiffUnderlying is a paid mutator transaction binding the contract method 0xc1b0b22b.
//
// Solidity: function withdrawDiffUnderlying(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) WithdrawDiffUnderlying(opts *bind.TransactOpts, leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "withdrawDiffUnderlying", leftoverShares)
}

// WithdrawDiffUnderlying is a paid mutator transaction binding the contract method 0xc1b0b22b.
//
// Solidity: function withdrawDiffUnderlying(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) WithdrawDiffUnderlying(leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawDiffUnderlying(&_Aavev2WrappedAToken.TransactOpts, leftoverShares)
}

// WithdrawDiffUnderlying is a paid mutator transaction binding the contract method 0xc1b0b22b.
//
// Solidity: function withdrawDiffUnderlying(uint256 leftoverShares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) WithdrawDiffUnderlying(leftoverShares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawDiffUnderlying(&_Aavev2WrappedAToken.TransactOpts, leftoverShares)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0x1071a290.
//
// Solidity: function withdrawUnderlying(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactor) WithdrawUnderlying(opts *bind.TransactOpts, shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.contract.Transact(opts, "withdrawUnderlying", shares)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0x1071a290.
//
// Solidity: function withdrawUnderlying(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenSession) WithdrawUnderlying(shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawUnderlying(&_Aavev2WrappedAToken.TransactOpts, shares)
}

// WithdrawUnderlying is a paid mutator transaction binding the contract method 0x1071a290.
//
// Solidity: function withdrawUnderlying(uint256 shares) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Aavev2WrappedAToken *Aavev2WrappedATokenTransactorSession) WithdrawUnderlying(shares *big.Int) (*types.Transaction, error) {
	return _Aavev2WrappedAToken.Contract.WithdrawUnderlying(&_Aavev2WrappedAToken.TransactOpts, shares)
}
