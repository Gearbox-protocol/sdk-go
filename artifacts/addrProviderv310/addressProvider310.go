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

// Addrv310MetaData contains all meta data concerning the Addrv310 contract.
var Addrv310MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMarketConfigurator\",\"inputs\":[{\"name\":\"_marketConfigurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addresses\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSavedContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structContractValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestAddressOrRevert\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMarketConfigurator\",\"inputs\":[{\"name\":\"riskCurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestVersions\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketConfiguratorByCreditManager\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketConfiguratorByPool\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketConfigurators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerCreditManager\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerPool\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeMarketConfigurator\",\"inputs\":[{\"name\":\"_marketConfigurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddMarketConfigurator\",\"inputs\":[{\"name\":\"marketConfigurator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveMarketConfigurator\",\"inputs\":[{\"name\":\"marketConfigurator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressNotFoundException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CantRemoveMarketConfiguratorWithExistingPoolsException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketConfiguratorNotFoundException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketConfiguratorsOnlyException\",\"inputs\":[]}]",
}

// Addrv310ABI is the input ABI used to generate the binding from.
// Deprecated: Use Addrv310MetaData.ABI instead.
var Addrv310ABI = Addrv310MetaData.ABI

// Addrv310 is an auto generated Go binding around an Ethereum contract.
type Addrv310 struct {
	Addrv310Caller     // Read-only binding to the contract
	Addrv310Transactor // Write-only binding to the contract
	Addrv310Filterer   // Log filterer for contract events
}

// Addrv310Caller is an auto generated read-only Go binding around an Ethereum contract.
type Addrv310Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Addrv310Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Addrv310Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Addrv310Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Addrv310Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Addrv310Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Addrv310Session struct {
	Contract     *Addrv310         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Addrv310CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Addrv310CallerSession struct {
	Contract *Addrv310Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Addrv310TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Addrv310TransactorSession struct {
	Contract     *Addrv310Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Addrv310Raw is an auto generated low-level Go binding around an Ethereum contract.
type Addrv310Raw struct {
	Contract *Addrv310 // Generic contract binding to access the raw methods on
}

// Addrv310CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Addrv310CallerRaw struct {
	Contract *Addrv310Caller // Generic read-only contract binding to access the raw methods on
}

// Addrv310TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Addrv310TransactorRaw struct {
	Contract *Addrv310Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrv310 creates a new instance of Addrv310, bound to a specific deployed contract.
func NewAddrv310(address common.Address, backend bind.ContractBackend) (*Addrv310, error) {
	contract, err := bindAddrv310(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Addrv310{Addrv310Caller: Addrv310Caller{contract: contract}, Addrv310Transactor: Addrv310Transactor{contract: contract}, Addrv310Filterer: Addrv310Filterer{contract: contract}}, nil
}

// NewAddrv310Caller creates a new read-only instance of Addrv310, bound to a specific deployed contract.
func NewAddrv310Caller(address common.Address, caller bind.ContractCaller) (*Addrv310Caller, error) {
	contract, err := bindAddrv310(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Addrv310Caller{contract: contract}, nil
}

// NewAddrv310Transactor creates a new write-only instance of Addrv310, bound to a specific deployed contract.
func NewAddrv310Transactor(address common.Address, transactor bind.ContractTransactor) (*Addrv310Transactor, error) {
	contract, err := bindAddrv310(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Addrv310Transactor{contract: contract}, nil
}

// NewAddrv310Filterer creates a new log filterer instance of Addrv310, bound to a specific deployed contract.
func NewAddrv310Filterer(address common.Address, filterer bind.ContractFilterer) (*Addrv310Filterer, error) {
	contract, err := bindAddrv310(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Addrv310Filterer{contract: contract}, nil
}

// bindAddrv310 binds a generic wrapper to an already deployed contract.
func bindAddrv310(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Addrv310MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addrv310 *Addrv310Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addrv310.Contract.Addrv310Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addrv310 *Addrv310Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addrv310.Contract.Addrv310Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addrv310 *Addrv310Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addrv310.Contract.Addrv310Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addrv310 *Addrv310CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addrv310.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addrv310 *Addrv310TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addrv310.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addrv310 *Addrv310TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addrv310.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_Addrv310 *Addrv310Caller) Addresses(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "addresses", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_Addrv310 *Addrv310Session) Addresses(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _Addrv310.Contract.Addresses(&_Addrv310.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0xace61afd.
//
// Solidity: function addresses(string , uint256 ) view returns(address)
func (_Addrv310 *Addrv310CallerSession) Addresses(arg0 string, arg1 *big.Int) (common.Address, error) {
	return _Addrv310.Contract.Addresses(&_Addrv310.CallOpts, arg0, arg1)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Addrv310 *Addrv310Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Addrv310 *Addrv310Session) ContractType() ([32]byte, error) {
	return _Addrv310.Contract.ContractType(&_Addrv310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Addrv310 *Addrv310CallerSession) ContractType() ([32]byte, error) {
	return _Addrv310.Contract.ContractType(&_Addrv310.CallOpts)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310Caller) GetAddressOrRevert(opts *bind.CallOpts, key string, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "getAddressOrRevert", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310Session) GetAddressOrRevert(key string, _version *big.Int) (common.Address, error) {
	return _Addrv310.Contract.GetAddressOrRevert(&_Addrv310.CallOpts, key, _version)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x0badbf3a.
//
// Solidity: function getAddressOrRevert(string key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310CallerSession) GetAddressOrRevert(key string, _version *big.Int) (common.Address, error) {
	return _Addrv310.Contract.GetAddressOrRevert(&_Addrv310.CallOpts, key, _version)
}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310Caller) GetAddressOrRevert0(opts *bind.CallOpts, key [32]byte, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "getAddressOrRevert0", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310Session) GetAddressOrRevert0(key [32]byte, _version *big.Int) (common.Address, error) {
	return _Addrv310.Contract.GetAddressOrRevert0(&_Addrv310.CallOpts, key, _version)
}

// GetAddressOrRevert0 is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_Addrv310 *Addrv310CallerSession) GetAddressOrRevert0(key [32]byte, _version *big.Int) (common.Address, error) {
	return _Addrv310.Contract.GetAddressOrRevert0(&_Addrv310.CallOpts, key, _version)
}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_Addrv310 *Addrv310Caller) GetAllSavedContracts(opts *bind.CallOpts) ([]ContractValue, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "getAllSavedContracts")

	if err != nil {
		return *new([]ContractValue), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractValue)).(*[]ContractValue)

	return out0, err

}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_Addrv310 *Addrv310Session) GetAllSavedContracts() ([]ContractValue, error) {
	return _Addrv310.Contract.GetAllSavedContracts(&_Addrv310.CallOpts)
}

// GetAllSavedContracts is a free data retrieval call binding the contract method 0xaad64b3a.
//
// Solidity: function getAllSavedContracts() view returns((string,address,uint256)[])
func (_Addrv310 *Addrv310CallerSession) GetAllSavedContracts() ([]ContractValue, error) {
	return _Addrv310.Contract.GetAllSavedContracts(&_Addrv310.CallOpts)
}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_Addrv310 *Addrv310Caller) GetLatestAddressOrRevert(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "getLatestAddressOrRevert", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_Addrv310 *Addrv310Session) GetLatestAddressOrRevert(_key [32]byte) (common.Address, error) {
	return _Addrv310.Contract.GetLatestAddressOrRevert(&_Addrv310.CallOpts, _key)
}

// GetLatestAddressOrRevert is a free data retrieval call binding the contract method 0x3b346687.
//
// Solidity: function getLatestAddressOrRevert(bytes32 _key) view returns(address result)
func (_Addrv310 *Addrv310CallerSession) GetLatestAddressOrRevert(_key [32]byte) (common.Address, error) {
	return _Addrv310.Contract.GetLatestAddressOrRevert(&_Addrv310.CallOpts, _key)
}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_Addrv310 *Addrv310Caller) GetLatestAddressOrRevert0(opts *bind.CallOpts, key string) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "getLatestAddressOrRevert0", key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_Addrv310 *Addrv310Session) GetLatestAddressOrRevert0(key string) (common.Address, error) {
	return _Addrv310.Contract.GetLatestAddressOrRevert0(&_Addrv310.CallOpts, key)
}

// GetLatestAddressOrRevert0 is a free data retrieval call binding the contract method 0xd631b136.
//
// Solidity: function getLatestAddressOrRevert(string key) view returns(address result)
func (_Addrv310 *Addrv310CallerSession) GetLatestAddressOrRevert0(key string) (common.Address, error) {
	return _Addrv310.Contract.GetLatestAddressOrRevert0(&_Addrv310.CallOpts, key)
}

// IsMarketConfigurator is a free data retrieval call binding the contract method 0x7b731af8.
//
// Solidity: function isMarketConfigurator(address riskCurator) view returns(bool)
func (_Addrv310 *Addrv310Caller) IsMarketConfigurator(opts *bind.CallOpts, riskCurator common.Address) (bool, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "isMarketConfigurator", riskCurator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMarketConfigurator is a free data retrieval call binding the contract method 0x7b731af8.
//
// Solidity: function isMarketConfigurator(address riskCurator) view returns(bool)
func (_Addrv310 *Addrv310Session) IsMarketConfigurator(riskCurator common.Address) (bool, error) {
	return _Addrv310.Contract.IsMarketConfigurator(&_Addrv310.CallOpts, riskCurator)
}

// IsMarketConfigurator is a free data retrieval call binding the contract method 0x7b731af8.
//
// Solidity: function isMarketConfigurator(address riskCurator) view returns(bool)
func (_Addrv310 *Addrv310CallerSession) IsMarketConfigurator(riskCurator common.Address) (bool, error) {
	return _Addrv310.Contract.IsMarketConfigurator(&_Addrv310.CallOpts, riskCurator)
}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_Addrv310 *Addrv310Caller) LatestVersions(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "latestVersions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_Addrv310 *Addrv310Session) LatestVersions(arg0 string) (*big.Int, error) {
	return _Addrv310.Contract.LatestVersions(&_Addrv310.CallOpts, arg0)
}

// LatestVersions is a free data retrieval call binding the contract method 0xb6a45d62.
//
// Solidity: function latestVersions(string ) view returns(uint256)
func (_Addrv310 *Addrv310CallerSession) LatestVersions(arg0 string) (*big.Int, error) {
	return _Addrv310.Contract.LatestVersions(&_Addrv310.CallOpts, arg0)
}

// MarketConfiguratorByCreditManager is a free data retrieval call binding the contract method 0xfb1ab826.
//
// Solidity: function marketConfiguratorByCreditManager(address creditManager) view returns(address)
func (_Addrv310 *Addrv310Caller) MarketConfiguratorByCreditManager(opts *bind.CallOpts, creditManager common.Address) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "marketConfiguratorByCreditManager", creditManager)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketConfiguratorByCreditManager is a free data retrieval call binding the contract method 0xfb1ab826.
//
// Solidity: function marketConfiguratorByCreditManager(address creditManager) view returns(address)
func (_Addrv310 *Addrv310Session) MarketConfiguratorByCreditManager(creditManager common.Address) (common.Address, error) {
	return _Addrv310.Contract.MarketConfiguratorByCreditManager(&_Addrv310.CallOpts, creditManager)
}

// MarketConfiguratorByCreditManager is a free data retrieval call binding the contract method 0xfb1ab826.
//
// Solidity: function marketConfiguratorByCreditManager(address creditManager) view returns(address)
func (_Addrv310 *Addrv310CallerSession) MarketConfiguratorByCreditManager(creditManager common.Address) (common.Address, error) {
	return _Addrv310.Contract.MarketConfiguratorByCreditManager(&_Addrv310.CallOpts, creditManager)
}

// MarketConfiguratorByPool is a free data retrieval call binding the contract method 0x9e0735c0.
//
// Solidity: function marketConfiguratorByPool(address ) view returns(address)
func (_Addrv310 *Addrv310Caller) MarketConfiguratorByPool(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "marketConfiguratorByPool", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketConfiguratorByPool is a free data retrieval call binding the contract method 0x9e0735c0.
//
// Solidity: function marketConfiguratorByPool(address ) view returns(address)
func (_Addrv310 *Addrv310Session) MarketConfiguratorByPool(arg0 common.Address) (common.Address, error) {
	return _Addrv310.Contract.MarketConfiguratorByPool(&_Addrv310.CallOpts, arg0)
}

// MarketConfiguratorByPool is a free data retrieval call binding the contract method 0x9e0735c0.
//
// Solidity: function marketConfiguratorByPool(address ) view returns(address)
func (_Addrv310 *Addrv310CallerSession) MarketConfiguratorByPool(arg0 common.Address) (common.Address, error) {
	return _Addrv310.Contract.MarketConfiguratorByPool(&_Addrv310.CallOpts, arg0)
}

// MarketConfigurators is a free data retrieval call binding the contract method 0xb0bc88a3.
//
// Solidity: function marketConfigurators() view returns(address[])
func (_Addrv310 *Addrv310Caller) MarketConfigurators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "marketConfigurators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// MarketConfigurators is a free data retrieval call binding the contract method 0xb0bc88a3.
//
// Solidity: function marketConfigurators() view returns(address[])
func (_Addrv310 *Addrv310Session) MarketConfigurators() ([]common.Address, error) {
	return _Addrv310.Contract.MarketConfigurators(&_Addrv310.CallOpts)
}

// MarketConfigurators is a free data retrieval call binding the contract method 0xb0bc88a3.
//
// Solidity: function marketConfigurators() view returns(address[])
func (_Addrv310 *Addrv310CallerSession) MarketConfigurators() ([]common.Address, error) {
	return _Addrv310.Contract.MarketConfigurators(&_Addrv310.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Addrv310 *Addrv310Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Addrv310 *Addrv310Session) Owner() (common.Address, error) {
	return _Addrv310.Contract.Owner(&_Addrv310.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Addrv310 *Addrv310CallerSession) Owner() (common.Address, error) {
	return _Addrv310.Contract.Owner(&_Addrv310.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Addrv310 *Addrv310Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Addrv310 *Addrv310Session) PendingOwner() (common.Address, error) {
	return _Addrv310.Contract.PendingOwner(&_Addrv310.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Addrv310 *Addrv310CallerSession) PendingOwner() (common.Address, error) {
	return _Addrv310.Contract.PendingOwner(&_Addrv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Addrv310 *Addrv310Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Addrv310.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Addrv310 *Addrv310Session) Version() (*big.Int, error) {
	return _Addrv310.Contract.Version(&_Addrv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Addrv310 *Addrv310CallerSession) Version() (*big.Int, error) {
	return _Addrv310.Contract.Version(&_Addrv310.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Addrv310 *Addrv310Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Addrv310 *Addrv310Session) AcceptOwnership() (*types.Transaction, error) {
	return _Addrv310.Contract.AcceptOwnership(&_Addrv310.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Addrv310 *Addrv310TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Addrv310.Contract.AcceptOwnership(&_Addrv310.TransactOpts)
}

// AddMarketConfigurator is a paid mutator transaction binding the contract method 0xcee2840e.
//
// Solidity: function addMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310Transactor) AddMarketConfigurator(opts *bind.TransactOpts, _marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "addMarketConfigurator", _marketConfigurator)
}

// AddMarketConfigurator is a paid mutator transaction binding the contract method 0xcee2840e.
//
// Solidity: function addMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310Session) AddMarketConfigurator(_marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.AddMarketConfigurator(&_Addrv310.TransactOpts, _marketConfigurator)
}

// AddMarketConfigurator is a paid mutator transaction binding the contract method 0xcee2840e.
//
// Solidity: function addMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310TransactorSession) AddMarketConfigurator(_marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.AddMarketConfigurator(&_Addrv310.TransactOpts, _marketConfigurator)
}

// RegisterCreditManager is a paid mutator transaction binding the contract method 0xaae93578.
//
// Solidity: function registerCreditManager(address creditManager) returns()
func (_Addrv310 *Addrv310Transactor) RegisterCreditManager(opts *bind.TransactOpts, creditManager common.Address) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "registerCreditManager", creditManager)
}

// RegisterCreditManager is a paid mutator transaction binding the contract method 0xaae93578.
//
// Solidity: function registerCreditManager(address creditManager) returns()
func (_Addrv310 *Addrv310Session) RegisterCreditManager(creditManager common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RegisterCreditManager(&_Addrv310.TransactOpts, creditManager)
}

// RegisterCreditManager is a paid mutator transaction binding the contract method 0xaae93578.
//
// Solidity: function registerCreditManager(address creditManager) returns()
func (_Addrv310 *Addrv310TransactorSession) RegisterCreditManager(creditManager common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RegisterCreditManager(&_Addrv310.TransactOpts, creditManager)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address pool) returns()
func (_Addrv310 *Addrv310Transactor) RegisterPool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "registerPool", pool)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address pool) returns()
func (_Addrv310 *Addrv310Session) RegisterPool(pool common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RegisterPool(&_Addrv310.TransactOpts, pool)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xabd90846.
//
// Solidity: function registerPool(address pool) returns()
func (_Addrv310 *Addrv310TransactorSession) RegisterPool(pool common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RegisterPool(&_Addrv310.TransactOpts, pool)
}

// RemoveMarketConfigurator is a paid mutator transaction binding the contract method 0xf2c99320.
//
// Solidity: function removeMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310Transactor) RemoveMarketConfigurator(opts *bind.TransactOpts, _marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "removeMarketConfigurator", _marketConfigurator)
}

// RemoveMarketConfigurator is a paid mutator transaction binding the contract method 0xf2c99320.
//
// Solidity: function removeMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310Session) RemoveMarketConfigurator(_marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RemoveMarketConfigurator(&_Addrv310.TransactOpts, _marketConfigurator)
}

// RemoveMarketConfigurator is a paid mutator transaction binding the contract method 0xf2c99320.
//
// Solidity: function removeMarketConfigurator(address _marketConfigurator) returns()
func (_Addrv310 *Addrv310TransactorSession) RemoveMarketConfigurator(_marketConfigurator common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.RemoveMarketConfigurator(&_Addrv310.TransactOpts, _marketConfigurator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Addrv310 *Addrv310Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Addrv310 *Addrv310Session) RenounceOwnership() (*types.Transaction, error) {
	return _Addrv310.Contract.RenounceOwnership(&_Addrv310.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Addrv310 *Addrv310TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Addrv310.Contract.RenounceOwnership(&_Addrv310.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_Addrv310 *Addrv310Transactor) SetAddress(opts *bind.TransactOpts, key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "setAddress", key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_Addrv310 *Addrv310Session) SetAddress(key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.Contract.SetAddress(&_Addrv310.TransactOpts, key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0x1870ad14.
//
// Solidity: function setAddress(string key, address value, bool saveVersion) returns()
func (_Addrv310 *Addrv310TransactorSession) SetAddress(key string, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.Contract.SetAddress(&_Addrv310.TransactOpts, key, value, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_Addrv310 *Addrv310Transactor) SetAddress0(opts *bind.TransactOpts, addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "setAddress0", addr, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_Addrv310 *Addrv310Session) SetAddress0(addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.Contract.SetAddress0(&_Addrv310.TransactOpts, addr, saveVersion)
}

// SetAddress0 is a paid mutator transaction binding the contract method 0x608b326f.
//
// Solidity: function setAddress(address addr, bool saveVersion) returns()
func (_Addrv310 *Addrv310TransactorSession) SetAddress0(addr common.Address, saveVersion bool) (*types.Transaction, error) {
	return _Addrv310.Contract.SetAddress0(&_Addrv310.TransactOpts, addr, saveVersion)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Addrv310 *Addrv310Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Addrv310.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Addrv310 *Addrv310Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.TransferOwnership(&_Addrv310.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Addrv310 *Addrv310TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Addrv310.Contract.TransferOwnership(&_Addrv310.TransactOpts, newOwner)
}

// Addrv310AddMarketConfiguratorIterator is returned from FilterAddMarketConfigurator and is used to iterate over the raw logs and unpacked data for AddMarketConfigurator events raised by the Addrv310 contract.
type Addrv310AddMarketConfiguratorIterator struct {
	Event *Addrv310AddMarketConfigurator // Event containing the contract specifics and raw log

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
func (it *Addrv310AddMarketConfiguratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Addrv310AddMarketConfigurator)
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
		it.Event = new(Addrv310AddMarketConfigurator)
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
func (it *Addrv310AddMarketConfiguratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Addrv310AddMarketConfiguratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Addrv310AddMarketConfigurator represents a AddMarketConfigurator event raised by the Addrv310 contract.
type Addrv310AddMarketConfigurator struct {
	MarketConfigurator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAddMarketConfigurator is a free log retrieval operation binding the contract event 0x1c7ee50fcfe7ff2e2c77e141fc72b116234385abf6000bb265a77963eba43c51.
//
// Solidity: event AddMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) FilterAddMarketConfigurator(opts *bind.FilterOpts, marketConfigurator []common.Address) (*Addrv310AddMarketConfiguratorIterator, error) {

	var marketConfiguratorRule []interface{}
	for _, marketConfiguratorItem := range marketConfigurator {
		marketConfiguratorRule = append(marketConfiguratorRule, marketConfiguratorItem)
	}

	logs, sub, err := _Addrv310.contract.FilterLogs(opts, "AddMarketConfigurator", marketConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &Addrv310AddMarketConfiguratorIterator{contract: _Addrv310.contract, event: "AddMarketConfigurator", logs: logs, sub: sub}, nil
}

// WatchAddMarketConfigurator is a free log subscription operation binding the contract event 0x1c7ee50fcfe7ff2e2c77e141fc72b116234385abf6000bb265a77963eba43c51.
//
// Solidity: event AddMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) WatchAddMarketConfigurator(opts *bind.WatchOpts, sink chan<- *Addrv310AddMarketConfigurator, marketConfigurator []common.Address) (event.Subscription, error) {

	var marketConfiguratorRule []interface{}
	for _, marketConfiguratorItem := range marketConfigurator {
		marketConfiguratorRule = append(marketConfiguratorRule, marketConfiguratorItem)
	}

	logs, sub, err := _Addrv310.contract.WatchLogs(opts, "AddMarketConfigurator", marketConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Addrv310AddMarketConfigurator)
				if err := _Addrv310.contract.UnpackLog(event, "AddMarketConfigurator", log); err != nil {
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

// ParseAddMarketConfigurator is a log parse operation binding the contract event 0x1c7ee50fcfe7ff2e2c77e141fc72b116234385abf6000bb265a77963eba43c51.
//
// Solidity: event AddMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) ParseAddMarketConfigurator(log types.Log) (*Addrv310AddMarketConfigurator, error) {
	event := new(Addrv310AddMarketConfigurator)
	if err := _Addrv310.contract.UnpackLog(event, "AddMarketConfigurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Addrv310OwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Addrv310 contract.
type Addrv310OwnershipTransferStartedIterator struct {
	Event *Addrv310OwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *Addrv310OwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Addrv310OwnershipTransferStarted)
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
		it.Event = new(Addrv310OwnershipTransferStarted)
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
func (it *Addrv310OwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Addrv310OwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Addrv310OwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Addrv310 contract.
type Addrv310OwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Addrv310 *Addrv310Filterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Addrv310OwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Addrv310.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Addrv310OwnershipTransferStartedIterator{contract: _Addrv310.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Addrv310 *Addrv310Filterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *Addrv310OwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Addrv310.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Addrv310OwnershipTransferStarted)
				if err := _Addrv310.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Addrv310 *Addrv310Filterer) ParseOwnershipTransferStarted(log types.Log) (*Addrv310OwnershipTransferStarted, error) {
	event := new(Addrv310OwnershipTransferStarted)
	if err := _Addrv310.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Addrv310OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Addrv310 contract.
type Addrv310OwnershipTransferredIterator struct {
	Event *Addrv310OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Addrv310OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Addrv310OwnershipTransferred)
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
		it.Event = new(Addrv310OwnershipTransferred)
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
func (it *Addrv310OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Addrv310OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Addrv310OwnershipTransferred represents a OwnershipTransferred event raised by the Addrv310 contract.
type Addrv310OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Addrv310 *Addrv310Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Addrv310OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Addrv310.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Addrv310OwnershipTransferredIterator{contract: _Addrv310.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Addrv310 *Addrv310Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Addrv310OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Addrv310.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Addrv310OwnershipTransferred)
				if err := _Addrv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Addrv310 *Addrv310Filterer) ParseOwnershipTransferred(log types.Log) (*Addrv310OwnershipTransferred, error) {
	event := new(Addrv310OwnershipTransferred)
	if err := _Addrv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Addrv310RemoveMarketConfiguratorIterator is returned from FilterRemoveMarketConfigurator and is used to iterate over the raw logs and unpacked data for RemoveMarketConfigurator events raised by the Addrv310 contract.
type Addrv310RemoveMarketConfiguratorIterator struct {
	Event *Addrv310RemoveMarketConfigurator // Event containing the contract specifics and raw log

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
func (it *Addrv310RemoveMarketConfiguratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Addrv310RemoveMarketConfigurator)
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
		it.Event = new(Addrv310RemoveMarketConfigurator)
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
func (it *Addrv310RemoveMarketConfiguratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Addrv310RemoveMarketConfiguratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Addrv310RemoveMarketConfigurator represents a RemoveMarketConfigurator event raised by the Addrv310 contract.
type Addrv310RemoveMarketConfigurator struct {
	MarketConfigurator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRemoveMarketConfigurator is a free log retrieval operation binding the contract event 0x6ce83e979d7e7ad30a9e4e20c214dcce48669a327d103ee3075adb93999b6f46.
//
// Solidity: event RemoveMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) FilterRemoveMarketConfigurator(opts *bind.FilterOpts, marketConfigurator []common.Address) (*Addrv310RemoveMarketConfiguratorIterator, error) {

	var marketConfiguratorRule []interface{}
	for _, marketConfiguratorItem := range marketConfigurator {
		marketConfiguratorRule = append(marketConfiguratorRule, marketConfiguratorItem)
	}

	logs, sub, err := _Addrv310.contract.FilterLogs(opts, "RemoveMarketConfigurator", marketConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &Addrv310RemoveMarketConfiguratorIterator{contract: _Addrv310.contract, event: "RemoveMarketConfigurator", logs: logs, sub: sub}, nil
}

// WatchRemoveMarketConfigurator is a free log subscription operation binding the contract event 0x6ce83e979d7e7ad30a9e4e20c214dcce48669a327d103ee3075adb93999b6f46.
//
// Solidity: event RemoveMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) WatchRemoveMarketConfigurator(opts *bind.WatchOpts, sink chan<- *Addrv310RemoveMarketConfigurator, marketConfigurator []common.Address) (event.Subscription, error) {

	var marketConfiguratorRule []interface{}
	for _, marketConfiguratorItem := range marketConfigurator {
		marketConfiguratorRule = append(marketConfiguratorRule, marketConfiguratorItem)
	}

	logs, sub, err := _Addrv310.contract.WatchLogs(opts, "RemoveMarketConfigurator", marketConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Addrv310RemoveMarketConfigurator)
				if err := _Addrv310.contract.UnpackLog(event, "RemoveMarketConfigurator", log); err != nil {
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

// ParseRemoveMarketConfigurator is a log parse operation binding the contract event 0x6ce83e979d7e7ad30a9e4e20c214dcce48669a327d103ee3075adb93999b6f46.
//
// Solidity: event RemoveMarketConfigurator(address indexed marketConfigurator)
func (_Addrv310 *Addrv310Filterer) ParseRemoveMarketConfigurator(log types.Log) (*Addrv310RemoveMarketConfigurator, error) {
	event := new(Addrv310RemoveMarketConfigurator)
	if err := _Addrv310.contract.UnpackLog(event, "RemoveMarketConfigurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Addrv310SetAddressIterator is returned from FilterSetAddress and is used to iterate over the raw logs and unpacked data for SetAddress events raised by the Addrv310 contract.
type Addrv310SetAddressIterator struct {
	Event *Addrv310SetAddress // Event containing the contract specifics and raw log

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
func (it *Addrv310SetAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Addrv310SetAddress)
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
		it.Event = new(Addrv310SetAddress)
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
func (it *Addrv310SetAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Addrv310SetAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Addrv310SetAddress represents a SetAddress event raised by the Addrv310 contract.
type Addrv310SetAddress struct {
	Key     string
	Value   common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAddress is a free log retrieval operation binding the contract event 0xdd66229bbe4615795978f9b416fa3c1b1d1f201f367612696a1c628b4debf435.
//
// Solidity: event SetAddress(string key, address indexed value, uint256 version)
func (_Addrv310 *Addrv310Filterer) FilterSetAddress(opts *bind.FilterOpts, value []common.Address) (*Addrv310SetAddressIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Addrv310.contract.FilterLogs(opts, "SetAddress", valueRule)
	if err != nil {
		return nil, err
	}
	return &Addrv310SetAddressIterator{contract: _Addrv310.contract, event: "SetAddress", logs: logs, sub: sub}, nil
}

// WatchSetAddress is a free log subscription operation binding the contract event 0xdd66229bbe4615795978f9b416fa3c1b1d1f201f367612696a1c628b4debf435.
//
// Solidity: event SetAddress(string key, address indexed value, uint256 version)
func (_Addrv310 *Addrv310Filterer) WatchSetAddress(opts *bind.WatchOpts, sink chan<- *Addrv310SetAddress, value []common.Address) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Addrv310.contract.WatchLogs(opts, "SetAddress", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Addrv310SetAddress)
				if err := _Addrv310.contract.UnpackLog(event, "SetAddress", log); err != nil {
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

// ParseSetAddress is a log parse operation binding the contract event 0xdd66229bbe4615795978f9b416fa3c1b1d1f201f367612696a1c628b4debf435.
//
// Solidity: event SetAddress(string key, address indexed value, uint256 version)
func (_Addrv310 *Addrv310Filterer) ParseSetAddress(log types.Log) (*Addrv310SetAddress, error) {
	event := new(Addrv310SetAddress)
	if err := _Addrv310.contract.UnpackLog(event, "SetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
