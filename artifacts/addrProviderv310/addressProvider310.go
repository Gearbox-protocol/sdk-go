// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addrProviderv310

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

// AddressProviderEntry is an auto generated low-level Go binding around an user-defined struct.
type AddressProviderEntry struct {
	Key   [32]byte
	Ver   *big.Int
	Value common.Address
}

// AddrProviderv310MetaData contains all meta data concerning the AddrProviderv310 contract.
var AddrProviderv310MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllEntries\",\"inputs\":[],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structAddressProviderEntry[]\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeys\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestMinorVersion\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"majorVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPatchVersion\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minorVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestVersion\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVersions\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SetAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressNotFoundException\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerException\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidVersionException\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ver\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"VersionNotFoundException\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// AddrProviderv310ABI is the input ABI used to generate the binding from.
// Deprecated: Use AddrProviderv310MetaData.ABI instead.
var AddrProviderv310ABI = AddrProviderv310MetaData.ABI

// AddrProviderv310 is an auto generated Go binding around an Ethereum contract.
type AddrProviderv310 struct {
	AddrProviderv310Caller     // Read-only binding to the contract
	AddrProviderv310Transactor // Write-only binding to the contract
	AddrProviderv310Filterer   // Log filterer for contract events
}

// AddrProviderv310Caller is an auto generated read-only Go binding around an Ethereum contract.
type AddrProviderv310Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrProviderv310Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AddrProviderv310Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrProviderv310Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddrProviderv310Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrProviderv310Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddrProviderv310Session struct {
	Contract     *AddrProviderv310 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddrProviderv310CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddrProviderv310CallerSession struct {
	Contract *AddrProviderv310Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AddrProviderv310TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddrProviderv310TransactorSession struct {
	Contract     *AddrProviderv310Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AddrProviderv310Raw is an auto generated low-level Go binding around an Ethereum contract.
type AddrProviderv310Raw struct {
	Contract *AddrProviderv310 // Generic contract binding to access the raw methods on
}

// AddrProviderv310CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddrProviderv310CallerRaw struct {
	Contract *AddrProviderv310Caller // Generic read-only contract binding to access the raw methods on
}

// AddrProviderv310TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddrProviderv310TransactorRaw struct {
	Contract *AddrProviderv310Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrProviderv310 creates a new instance of AddrProviderv310, bound to a specific deployed contract.
func NewAddrProviderv310(address common.Address, backend bind.ContractBackend) (*AddrProviderv310, error) {
	contract, err := bindAddrProviderv310(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310{AddrProviderv310Caller: AddrProviderv310Caller{contract: contract}, AddrProviderv310Transactor: AddrProviderv310Transactor{contract: contract}, AddrProviderv310Filterer: AddrProviderv310Filterer{contract: contract}}, nil
}

// NewAddrProviderv310Caller creates a new read-only instance of AddrProviderv310, bound to a specific deployed contract.
func NewAddrProviderv310Caller(address common.Address, caller bind.ContractCaller) (*AddrProviderv310Caller, error) {
	contract, err := bindAddrProviderv310(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310Caller{contract: contract}, nil
}

// NewAddrProviderv310Transactor creates a new write-only instance of AddrProviderv310, bound to a specific deployed contract.
func NewAddrProviderv310Transactor(address common.Address, transactor bind.ContractTransactor) (*AddrProviderv310Transactor, error) {
	contract, err := bindAddrProviderv310(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310Transactor{contract: contract}, nil
}

// NewAddrProviderv310Filterer creates a new log filterer instance of AddrProviderv310, bound to a specific deployed contract.
func NewAddrProviderv310Filterer(address common.Address, filterer bind.ContractFilterer) (*AddrProviderv310Filterer, error) {
	contract, err := bindAddrProviderv310(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310Filterer{contract: contract}, nil
}

// bindAddrProviderv310 binds a generic wrapper to an already deployed contract.
func bindAddrProviderv310(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddrProviderv310MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrProviderv310 *AddrProviderv310Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddrProviderv310.Contract.AddrProviderv310Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrProviderv310 *AddrProviderv310Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.AddrProviderv310Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrProviderv310 *AddrProviderv310Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.AddrProviderv310Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrProviderv310 *AddrProviderv310CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddrProviderv310.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrProviderv310 *AddrProviderv310TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrProviderv310 *AddrProviderv310TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AddrProviderv310 *AddrProviderv310Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AddrProviderv310 *AddrProviderv310Session) ContractType() ([32]byte, error) {
	return _AddrProviderv310.Contract.ContractType(&_AddrProviderv310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AddrProviderv310 *AddrProviderv310CallerSession) ContractType() ([32]byte, error) {
	return _AddrProviderv310.Contract.ContractType(&_AddrProviderv310.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0xbbd6dd6b.
//
// Solidity: function getAddress(bytes32 key, uint256 ver) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAddress(opts *bind.CallOpts, key [32]byte, ver *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAddress", key, ver)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbbd6dd6b.
//
// Solidity: function getAddress(bytes32 key, uint256 ver) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Session) GetAddress(key [32]byte, ver *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddress(&_AddrProviderv310.CallOpts, key, ver)
}

// GetAddress is a free data retrieval call binding the contract method 0xbbd6dd6b.
//
// Solidity: function getAddress(bytes32 key, uint256 ver) view returns(address)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAddress(key [32]byte, ver *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddress(&_AddrProviderv310.CallOpts, key, ver)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 ver) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAddressOrRevert(opts *bind.CallOpts, key [32]byte, ver *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAddressOrRevert", key, ver)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 ver) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Session) GetAddressOrRevert(key [32]byte, ver *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert(&_AddrProviderv310.CallOpts, key, ver)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 ver) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAddressOrRevert(key [32]byte, ver *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert(&_AddrProviderv310.CallOpts, key, ver)
}

// GetAllEntries is a free data retrieval call binding the contract method 0x4a0364e8.
//
// Solidity: function getAllEntries() view returns((bytes32,uint256,address)[] entries)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAllEntries(opts *bind.CallOpts) ([]AddressProviderEntry, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAllEntries")

	if err != nil {
		return *new([]AddressProviderEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]AddressProviderEntry)).(*[]AddressProviderEntry)

	return out0, err

}

// GetAllEntries is a free data retrieval call binding the contract method 0x4a0364e8.
//
// Solidity: function getAllEntries() view returns((bytes32,uint256,address)[] entries)
func (_AddrProviderv310 *AddrProviderv310Session) GetAllEntries() ([]AddressProviderEntry, error) {
	return _AddrProviderv310.Contract.GetAllEntries(&_AddrProviderv310.CallOpts)
}

// GetAllEntries is a free data retrieval call binding the contract method 0x4a0364e8.
//
// Solidity: function getAllEntries() view returns((bytes32,uint256,address)[] entries)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAllEntries() ([]AddressProviderEntry, error) {
	return _AddrProviderv310.Contract.GetAllEntries(&_AddrProviderv310.CallOpts)
}

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(bytes32[])
func (_AddrProviderv310 *AddrProviderv310Caller) GetKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(bytes32[])
func (_AddrProviderv310 *AddrProviderv310Session) GetKeys() ([][32]byte, error) {
	return _AddrProviderv310.Contract.GetKeys(&_AddrProviderv310.CallOpts)
}

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(bytes32[])
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetKeys() ([][32]byte, error) {
	return _AddrProviderv310.Contract.GetKeys(&_AddrProviderv310.CallOpts)
}

// GetLatestMinorVersion is a free data retrieval call binding the contract method 0x170ecb17.
//
// Solidity: function getLatestMinorVersion(bytes32 key, uint256 majorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Caller) GetLatestMinorVersion(opts *bind.CallOpts, key [32]byte, majorVersion *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getLatestMinorVersion", key, majorVersion)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestMinorVersion is a free data retrieval call binding the contract method 0x170ecb17.
//
// Solidity: function getLatestMinorVersion(bytes32 key, uint256 majorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Session) GetLatestMinorVersion(key [32]byte, majorVersion *big.Int) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestMinorVersion(&_AddrProviderv310.CallOpts, key, majorVersion)
}

// GetLatestMinorVersion is a free data retrieval call binding the contract method 0x170ecb17.
//
// Solidity: function getLatestMinorVersion(bytes32 key, uint256 majorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetLatestMinorVersion(key [32]byte, majorVersion *big.Int) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestMinorVersion(&_AddrProviderv310.CallOpts, key, majorVersion)
}

// GetLatestPatchVersion is a free data retrieval call binding the contract method 0x4ceff70e.
//
// Solidity: function getLatestPatchVersion(bytes32 key, uint256 minorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Caller) GetLatestPatchVersion(opts *bind.CallOpts, key [32]byte, minorVersion *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getLatestPatchVersion", key, minorVersion)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPatchVersion is a free data retrieval call binding the contract method 0x4ceff70e.
//
// Solidity: function getLatestPatchVersion(bytes32 key, uint256 minorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Session) GetLatestPatchVersion(key [32]byte, minorVersion *big.Int) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestPatchVersion(&_AddrProviderv310.CallOpts, key, minorVersion)
}

// GetLatestPatchVersion is a free data retrieval call binding the contract method 0x4ceff70e.
//
// Solidity: function getLatestPatchVersion(bytes32 key, uint256 minorVersion) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetLatestPatchVersion(key [32]byte, minorVersion *big.Int) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestPatchVersion(&_AddrProviderv310.CallOpts, key, minorVersion)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0xdd3b014c.
//
// Solidity: function getLatestVersion(bytes32 key) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Caller) GetLatestVersion(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getLatestVersion", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestVersion is a free data retrieval call binding the contract method 0xdd3b014c.
//
// Solidity: function getLatestVersion(bytes32 key) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310Session) GetLatestVersion(key [32]byte) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestVersion(&_AddrProviderv310.CallOpts, key)
}

// GetLatestVersion is a free data retrieval call binding the contract method 0xdd3b014c.
//
// Solidity: function getLatestVersion(bytes32 key) view returns(uint256 ver)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetLatestVersion(key [32]byte) (*big.Int, error) {
	return _AddrProviderv310.Contract.GetLatestVersion(&_AddrProviderv310.CallOpts, key)
}

// GetVersions is a free data retrieval call binding the contract method 0x9738418c.
//
// Solidity: function getVersions(bytes32 key) view returns(uint256[])
func (_AddrProviderv310 *AddrProviderv310Caller) GetVersions(opts *bind.CallOpts, key [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getVersions", key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVersions is a free data retrieval call binding the contract method 0x9738418c.
//
// Solidity: function getVersions(bytes32 key) view returns(uint256[])
func (_AddrProviderv310 *AddrProviderv310Session) GetVersions(key [32]byte) ([]*big.Int, error) {
	return _AddrProviderv310.Contract.GetVersions(&_AddrProviderv310.CallOpts, key)
}

// GetVersions is a free data retrieval call binding the contract method 0x9738418c.
//
// Solidity: function getVersions(bytes32 key) view returns(uint256[])
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetVersions(key [32]byte) ([]*big.Int, error) {
	return _AddrProviderv310.Contract.GetVersions(&_AddrProviderv310.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310Session) Owner() (common.Address, error) {
	return _AddrProviderv310.Contract.Owner(&_AddrProviderv310.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310CallerSession) Owner() (common.Address, error) {
	return _AddrProviderv310.Contract.Owner(&_AddrProviderv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310Session) Version() (*big.Int, error) {
	return _AddrProviderv310.Contract.Version(&_AddrProviderv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310CallerSession) Version() (*big.Int, error) {
	return _AddrProviderv310.Contract.Version(&_AddrProviderv310.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) SetAddress(opts *bind.TransactOpts, key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "setAddress", key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Session) SetAddress(key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress(&_AddrProviderv310.TransactOpts, key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) SetAddress(key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress(&_AddrProviderv310.TransactOpts, key, value, saveVersion)
}

// AddrProviderv310SetAddressIterator is returned from FilterSetAddress and is used to iterate over the raw logs and unpacked data for SetAddress events raised by the AddrProviderv310 contract.
type AddrProviderv310SetAddressIterator struct {
	Event *AddrProviderv310SetAddress // Event containing the contract specifics and raw log

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
func (it *AddrProviderv310SetAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrProviderv310SetAddress)
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
		it.Event = new(AddrProviderv310SetAddress)
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
func (it *AddrProviderv310SetAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrProviderv310SetAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrProviderv310SetAddress represents a SetAddress event raised by the AddrProviderv310 contract.
type AddrProviderv310SetAddress struct {
	Key   [32]byte
	Ver   *big.Int
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAddress is a free log retrieval operation binding the contract event 0x426978cf4f253ea2b73ddec55a29967461d22a644b2798d388c33857ede15802.
//
// Solidity: event SetAddress(bytes32 indexed key, uint256 indexed ver, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) FilterSetAddress(opts *bind.FilterOpts, key [][32]byte, ver []*big.Int, value []common.Address) (*AddrProviderv310SetAddressIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var verRule []interface{}
	for _, verItem := range ver {
		verRule = append(verRule, verItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _AddrProviderv310.contract.FilterLogs(opts, "SetAddress", keyRule, verRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310SetAddressIterator{contract: _AddrProviderv310.contract, event: "SetAddress", logs: logs, sub: sub}, nil
}

// WatchSetAddress is a free log subscription operation binding the contract event 0x426978cf4f253ea2b73ddec55a29967461d22a644b2798d388c33857ede15802.
//
// Solidity: event SetAddress(bytes32 indexed key, uint256 indexed ver, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) WatchSetAddress(opts *bind.WatchOpts, sink chan<- *AddrProviderv310SetAddress, key [][32]byte, ver []*big.Int, value []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var verRule []interface{}
	for _, verItem := range ver {
		verRule = append(verRule, verItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _AddrProviderv310.contract.WatchLogs(opts, "SetAddress", keyRule, verRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrProviderv310SetAddress)
				if err := _AddrProviderv310.contract.UnpackLog(event, "SetAddress", log); err != nil {
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

// ParseSetAddress is a log parse operation binding the contract event 0x426978cf4f253ea2b73ddec55a29967461d22a644b2798d388c33857ede15802.
//
// Solidity: event SetAddress(bytes32 indexed key, uint256 indexed ver, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) ParseSetAddress(log types.Log) (*AddrProviderv310SetAddress, error) {
	event := new(AddrProviderv310SetAddress)
	if err := _AddrProviderv310.contract.UnpackLog(event, "SetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
