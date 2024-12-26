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

// ContractValue is an auto generated low-level Go binding around an user-defined struct.
type ContractValue struct {
	Key     string
	Value   common.Address
	Version *big.Int
}

// AddrProviderv310MetaData contains all meta data concerning the AddrProviderv310 contract.
var AddrProviderv310MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addresses\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSavedContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structContractValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestAddressOrRevert\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestVersions\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressNotFoundException\",\"inputs\":[]}]",
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

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Caller) Addresses(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "addresses", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Session) Addresses(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.Addresses(&_AddrProviderv310.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_AddrProviderv310 *AddrProviderv310CallerSession) Addresses(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.Addresses(&_AddrProviderv310.CallOpts, arg0, arg1)
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
// Solidity: function getAddress(bytes32 key, uint256 _version) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAddress(opts *bind.CallOpts, key [32]byte, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAddress", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbbd6dd6b.
//
// Solidity: function getAddress(bytes32 key, uint256 _version) view returns(address)
func (_AddrProviderv310 *AddrProviderv310Session) GetAddress(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddress(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAddress is a free data retrieval call binding the contract method 0xbbd6dd6b.
//
// Solidity: function getAddress(bytes32 key, uint256 _version) view returns(address)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAddress(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddress(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAddressOrRevert(opts *bind.CallOpts, key string, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAddressOrRevert", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Session) GetAddressOrRevert(key string, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAddressOrRevert(key string, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Caller) GetAddressOrRevert0(opts *bind.CallOpts, key [32]byte, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAddressOrRevert0", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Session) GetAddressOrRevert0(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert0(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAddressOrRevert0(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddrProviderv310.Contract.GetAddressOrRevert0(&_AddrProviderv310.CallOpts, key, _version)
}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_AddrProviderv310 *AddrProviderv310Caller) GetAllSavedContracts(opts *bind.CallOpts) ([]ContractValue, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getAllSavedContracts")

	if err != nil {
		return *new([]ContractValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractValue)).(*[]ContractValue)

	return out0, err

}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_AddrProviderv310 *AddrProviderv310Session) GetAllSavedContracts() ([]ContractValue, error) {
	return _AddrProviderv310.Contract.GetAllSavedContracts(&_AddrProviderv310.CallOpts)
}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetAllSavedContracts() ([]ContractValue, error) {
	return _AddrProviderv310.Contract.GetAllSavedContracts(&_AddrProviderv310.CallOpts)
}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Caller) GetLatestAddressOrRevert(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getLatestAddressOrRevert", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Session) GetLatestAddressOrRevert(_key [32]byte) (common.Address, error) {
	return _AddrProviderv310.Contract.GetLatestAddressOrRevert(&_AddrProviderv310.CallOpts, _key)
}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetLatestAddressOrRevert(_key [32]byte) (common.Address, error) {
	return _AddrProviderv310.Contract.GetLatestAddressOrRevert(&_AddrProviderv310.CallOpts, _key)
}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Caller) GetLatestAddressOrRevert0(opts *bind.CallOpts, key string) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "getLatestAddressOrRevert0", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310Session) GetLatestAddressOrRevert0(key string) (common.Address, error) {
	return _AddrProviderv310.Contract.GetLatestAddressOrRevert0(&_AddrProviderv310.CallOpts, key)
}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_AddrProviderv310 *AddrProviderv310CallerSession) GetLatestAddressOrRevert0(key string) (common.Address, error) {
	return _AddrProviderv310.Contract.GetLatestAddressOrRevert0(&_AddrProviderv310.CallOpts, key)
}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310Caller) LatestVersions(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "latestVersions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310Session) LatestVersions(arg0 string) (*big.Int, error) {
	return _AddrProviderv310.Contract.LatestVersions(&_AddrProviderv310.CallOpts, arg0)
}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_AddrProviderv310 *AddrProviderv310CallerSession) LatestVersions(arg0 string) (*big.Int, error) {
	return _AddrProviderv310.Contract.LatestVersions(&_AddrProviderv310.CallOpts, arg0)
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

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddrProviderv310.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310Session) PendingOwner() (common.Address, error) {
	return _AddrProviderv310.Contract.PendingOwner(&_AddrProviderv310.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AddrProviderv310 *AddrProviderv310CallerSession) PendingOwner() (common.Address, error) {
	return _AddrProviderv310.Contract.PendingOwner(&_AddrProviderv310.CallOpts)
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

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310Session) AcceptOwnership() (*types.Transaction, error) {
	return _AddrProviderv310.Contract.AcceptOwnership(&_AddrProviderv310.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AddrProviderv310.Contract.AcceptOwnership(&_AddrProviderv310.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310Session) RenounceOwnership() (*types.Transaction, error) {
	return _AddrProviderv310.Contract.RenounceOwnership(&_AddrProviderv310.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddrProviderv310.Contract.RenounceOwnership(&_AddrProviderv310.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) SetAddress(opts *bind.TransactOpts, key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "setAddress", key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Session) SetAddress(key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress(&_AddrProviderv310.TransactOpts, key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) SetAddress(key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress(&_AddrProviderv310.TransactOpts, key, value, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) SetAddress0(opts *bind.TransactOpts, addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "setAddress0", addr, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310Session) SetAddress0(addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress0(&_AddrProviderv310.TransactOpts, addr, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) SetAddress0(addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.SetAddress0(&_AddrProviderv310.TransactOpts, addr, saveVersion)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddrProviderv310 *AddrProviderv310Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddrProviderv310.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddrProviderv310 *AddrProviderv310Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.TransferOwnership(&_AddrProviderv310.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddrProviderv310 *AddrProviderv310TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddrProviderv310.Contract.TransferOwnership(&_AddrProviderv310.TransactOpts, newOwner)
}

// AddrProviderv310OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the AddrProviderv310 contract.
type AddrProviderv310OwnershipTransferStartedIterator struct {
	Event *AddrProviderv310OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *AddrProviderv310OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrProviderv310OwnershipTransferStarted)
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
		it.Event = new(AddrProviderv310OwnershipTransferStarted)
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
func (it *AddrProviderv310OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrProviderv310OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrProviderv310OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the AddrProviderv310 contract.
type AddrProviderv310OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddrProviderv310OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddrProviderv310.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310OwnershipTransferStartedIterator{contract: _AddrProviderv310.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *AddrProviderv310OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddrProviderv310.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrProviderv310OwnershipTransferStarted)
				if err := _AddrProviderv310.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) ParseOwnershipTransferStarted(log types.Log) (*AddrProviderv310OwnershipTransferStarted, error) {
	event := new(AddrProviderv310OwnershipTransferStarted)
	if err := _AddrProviderv310.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddrProviderv310OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddrProviderv310 contract.
type AddrProviderv310OwnershipTransferredIterator struct {
	Event *AddrProviderv310OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddrProviderv310OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddrProviderv310OwnershipTransferred)
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
		it.Event = new(AddrProviderv310OwnershipTransferred)
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
func (it *AddrProviderv310OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddrProviderv310OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddrProviderv310OwnershipTransferred represents a OwnershipTransferred event raised by the AddrProviderv310 contract.
type AddrProviderv310OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddrProviderv310OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddrProviderv310.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310OwnershipTransferredIterator{contract: _AddrProviderv310.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddrProviderv310OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddrProviderv310.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddrProviderv310OwnershipTransferred)
				if err := _AddrProviderv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddrProviderv310 *AddrProviderv310Filterer) ParseOwnershipTransferred(log types.Log) (*AddrProviderv310OwnershipTransferred, error) {
	event := new(AddrProviderv310OwnershipTransferred)
	if err := _AddrProviderv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Key     common.Hash
	Version *big.Int
	Value   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAddress is a free log retrieval operation binding the contract event 0x36b2968106d776973f4f92aac227fdf92a080898e36dd7314a22124679c050dd.
//
// Solidity: event SetAddress(string indexed key, uint256 indexed version, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) FilterSetAddress(opts *bind.FilterOpts, key []string, version []*big.Int, value []common.Address) (*AddrProviderv310SetAddressIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _AddrProviderv310.contract.FilterLogs(opts, "SetAddress", keyRule, versionRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &AddrProviderv310SetAddressIterator{contract: _AddrProviderv310.contract, event: "SetAddress", logs: logs, sub: sub}, nil
}

// WatchSetAddress is a free log subscription operation binding the contract event 0x36b2968106d776973f4f92aac227fdf92a080898e36dd7314a22124679c050dd.
//
// Solidity: event SetAddress(string indexed key, uint256 indexed version, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) WatchSetAddress(opts *bind.WatchOpts, sink chan<- *AddrProviderv310SetAddress, key []string, version []*big.Int, value []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _AddrProviderv310.contract.WatchLogs(opts, "SetAddress", keyRule, versionRule, valueRule)
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

// ParseSetAddress is a log parse operation binding the contract event 0x36b2968106d776973f4f92aac227fdf92a080898e36dd7314a22124679c050dd.
//
// Solidity: event SetAddress(string indexed key, uint256 indexed version, address indexed value)
func (_AddrProviderv310 *AddrProviderv310Filterer) ParseSetAddress(log types.Log) (*AddrProviderv310SetAddress, error) {
	event := new(AddrProviderv310SetAddress)
	if err := _AddrProviderv310.contract.UnpackLog(event, "SetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
