// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addressProviderv3

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

// AddressProviderv3MetaData contains all meta data concerning the AddressProviderv3 contract.
var AddressProviderv3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_acl\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addresses\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getACL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddressOrRevert\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractsRegister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDataCompressor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGearToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLeveragedActions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPriceOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTreasuryContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWETHGateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWethToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"saveVersion\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SetAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressNotFoundException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotConfiguratorException\",\"inputs\":[]}]",
}

// AddressProviderv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressProviderv3MetaData.ABI instead.
var AddressProviderv3ABI = AddressProviderv3MetaData.ABI

// AddressProviderv3 is an auto generated Go binding around an Ethereum contract.
type AddressProviderv3 struct {
	AddressProviderv3Caller     // Read-only binding to the contract
	AddressProviderv3Transactor // Write-only binding to the contract
	AddressProviderv3Filterer   // Log filterer for contract events
}

// AddressProviderv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AddressProviderv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressProviderv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressProviderv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressProviderv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressProviderv3Session struct {
	Contract     *AddressProviderv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressProviderv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressProviderv3CallerSession struct {
	Contract *AddressProviderv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AddressProviderv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressProviderv3TransactorSession struct {
	Contract     *AddressProviderv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AddressProviderv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AddressProviderv3Raw struct {
	Contract *AddressProviderv3 // Generic contract binding to access the raw methods on
}

// AddressProviderv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressProviderv3CallerRaw struct {
	Contract *AddressProviderv3Caller // Generic read-only contract binding to access the raw methods on
}

// AddressProviderv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressProviderv3TransactorRaw struct {
	Contract *AddressProviderv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressProviderv3 creates a new instance of AddressProviderv3, bound to a specific deployed contract.
func NewAddressProviderv3(address common.Address, backend bind.ContractBackend) (*AddressProviderv3, error) {
	contract, err := bindAddressProviderv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressProviderv3{AddressProviderv3Caller: AddressProviderv3Caller{contract: contract}, AddressProviderv3Transactor: AddressProviderv3Transactor{contract: contract}, AddressProviderv3Filterer: AddressProviderv3Filterer{contract: contract}}, nil
}

// NewAddressProviderv3Caller creates a new read-only instance of AddressProviderv3, bound to a specific deployed contract.
func NewAddressProviderv3Caller(address common.Address, caller bind.ContractCaller) (*AddressProviderv3Caller, error) {
	contract, err := bindAddressProviderv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderv3Caller{contract: contract}, nil
}

// NewAddressProviderv3Transactor creates a new write-only instance of AddressProviderv3, bound to a specific deployed contract.
func NewAddressProviderv3Transactor(address common.Address, transactor bind.ContractTransactor) (*AddressProviderv3Transactor, error) {
	contract, err := bindAddressProviderv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressProviderv3Transactor{contract: contract}, nil
}

// NewAddressProviderv3Filterer creates a new log filterer instance of AddressProviderv3, bound to a specific deployed contract.
func NewAddressProviderv3Filterer(address common.Address, filterer bind.ContractFilterer) (*AddressProviderv3Filterer, error) {
	contract, err := bindAddressProviderv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressProviderv3Filterer{contract: contract}, nil
}

// bindAddressProviderv3 binds a generic wrapper to an already deployed contract.
func bindAddressProviderv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressProviderv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProviderv3 *AddressProviderv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProviderv3.Contract.AddressProviderv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProviderv3 *AddressProviderv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.AddressProviderv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProviderv3 *AddressProviderv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.AddressProviderv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressProviderv3 *AddressProviderv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressProviderv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressProviderv3 *AddressProviderv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressProviderv3 *AddressProviderv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xb76b70d5.
//
// Solidity: function addresses(bytes32 , uint256 ) view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) Addresses(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "addresses", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xb76b70d5.
//
// Solidity: function addresses(bytes32 , uint256 ) view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) Addresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AddressProviderv3.Contract.Addresses(&_AddressProviderv3.CallOpts, arg0, arg1)
}

// Addresses is a free data retrieval call binding the contract method 0xb76b70d5.
//
// Solidity: function addresses(bytes32 , uint256 ) view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) Addresses(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _AddressProviderv3.Contract.Addresses(&_AddressProviderv3.CallOpts, arg0, arg1)
}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetACL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getACL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetACL() (common.Address, error) {
	return _AddressProviderv3.Contract.GetACL(&_AddressProviderv3.CallOpts)
}

// GetACL is a free data retrieval call binding the contract method 0x08737695.
//
// Solidity: function getACL() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetACL() (common.Address, error) {
	return _AddressProviderv3.Contract.GetACL(&_AddressProviderv3.CallOpts)
}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetAccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getAccountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetAccountFactory() (common.Address, error) {
	return _AddressProviderv3.Contract.GetAccountFactory(&_AddressProviderv3.CallOpts)
}

// GetAccountFactory is a free data retrieval call binding the contract method 0x9068a868.
//
// Solidity: function getAccountFactory() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetAccountFactory() (common.Address, error) {
	return _AddressProviderv3.Contract.GetAccountFactory(&_AddressProviderv3.CallOpts)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddressProviderv3 *AddressProviderv3Caller) GetAddressOrRevert(opts *bind.CallOpts, key [32]byte, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getAddressOrRevert", key, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddressProviderv3 *AddressProviderv3Session) GetAddressOrRevert(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddressProviderv3.Contract.GetAddressOrRevert(&_AddressProviderv3.CallOpts, key, _version)
}

// GetAddressOrRevert is a free data retrieval call binding the contract method 0x57b5a1c6.
//
// Solidity: function getAddressOrRevert(bytes32 key, uint256 _version) view returns(address result)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetAddressOrRevert(key [32]byte, _version *big.Int) (common.Address, error) {
	return _AddressProviderv3.Contract.GetAddressOrRevert(&_AddressProviderv3.CallOpts, key, _version)
}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getContractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetContractsRegister() (common.Address, error) {
	return _AddressProviderv3.Contract.GetContractsRegister(&_AddressProviderv3.CallOpts)
}

// GetContractsRegister is a free data retrieval call binding the contract method 0xc513c9bb.
//
// Solidity: function getContractsRegister() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetContractsRegister() (common.Address, error) {
	return _AddressProviderv3.Contract.GetContractsRegister(&_AddressProviderv3.CallOpts)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetDataCompressor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getDataCompressor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetDataCompressor() (common.Address, error) {
	return _AddressProviderv3.Contract.GetDataCompressor(&_AddressProviderv3.CallOpts)
}

// GetDataCompressor is a free data retrieval call binding the contract method 0x060678c2.
//
// Solidity: function getDataCompressor() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetDataCompressor() (common.Address, error) {
	return _AddressProviderv3.Contract.GetDataCompressor(&_AddressProviderv3.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetGearToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getGearToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetGearToken() (common.Address, error) {
	return _AddressProviderv3.Contract.GetGearToken(&_AddressProviderv3.CallOpts)
}

// GetGearToken is a free data retrieval call binding the contract method 0xaffd9243.
//
// Solidity: function getGearToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetGearToken() (common.Address, error) {
	return _AddressProviderv3.Contract.GetGearToken(&_AddressProviderv3.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetLeveragedActions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getLeveragedActions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetLeveragedActions() (common.Address, error) {
	return _AddressProviderv3.Contract.GetLeveragedActions(&_AddressProviderv3.CallOpts)
}

// GetLeveragedActions is a free data retrieval call binding the contract method 0x44b88563.
//
// Solidity: function getLeveragedActions() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetLeveragedActions() (common.Address, error) {
	return _AddressProviderv3.Contract.GetLeveragedActions(&_AddressProviderv3.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetPriceOracle() (common.Address, error) {
	return _AddressProviderv3.Contract.GetPriceOracle(&_AddressProviderv3.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetPriceOracle() (common.Address, error) {
	return _AddressProviderv3.Contract.GetPriceOracle(&_AddressProviderv3.CallOpts)
}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetTreasuryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getTreasuryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetTreasuryContract() (common.Address, error) {
	return _AddressProviderv3.Contract.GetTreasuryContract(&_AddressProviderv3.CallOpts)
}

// GetTreasuryContract is a free data retrieval call binding the contract method 0x26c74fc3.
//
// Solidity: function getTreasuryContract() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetTreasuryContract() (common.Address, error) {
	return _AddressProviderv3.Contract.GetTreasuryContract(&_AddressProviderv3.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetWETHGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getWETHGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetWETHGateway() (common.Address, error) {
	return _AddressProviderv3.Contract.GetWETHGateway(&_AddressProviderv3.CallOpts)
}

// GetWETHGateway is a free data retrieval call binding the contract method 0x77532ed9.
//
// Solidity: function getWETHGateway() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetWETHGateway() (common.Address, error) {
	return _AddressProviderv3.Contract.GetWETHGateway(&_AddressProviderv3.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Caller) GetWethToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "getWethToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3Session) GetWethToken() (common.Address, error) {
	return _AddressProviderv3.Contract.GetWethToken(&_AddressProviderv3.CallOpts)
}

// GetWethToken is a free data retrieval call binding the contract method 0x4c252f91.
//
// Solidity: function getWethToken() view returns(address)
func (_AddressProviderv3 *AddressProviderv3CallerSession) GetWethToken() (common.Address, error) {
	return _AddressProviderv3.Contract.GetWethToken(&_AddressProviderv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddressProviderv3 *AddressProviderv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressProviderv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddressProviderv3 *AddressProviderv3Session) Version() (*big.Int, error) {
	return _AddressProviderv3.Contract.Version(&_AddressProviderv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AddressProviderv3 *AddressProviderv3CallerSession) Version() (*big.Int, error) {
	return _AddressProviderv3.Contract.Version(&_AddressProviderv3.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddressProviderv3 *AddressProviderv3Transactor) SetAddress(opts *bind.TransactOpts, key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddressProviderv3.contract.Transact(opts, "setAddress", key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddressProviderv3 *AddressProviderv3Session) SetAddress(key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.SetAddress(&_AddressProviderv3.TransactOpts, key, value, saveVersion)
}

// SetAddress is a paid mutator transaction binding the contract method 0xbe99a980.
//
// Solidity: function setAddress(bytes32 key, address value, bool saveVersion) returns()
func (_AddressProviderv3 *AddressProviderv3TransactorSession) SetAddress(key [32]byte, value common.Address, saveVersion bool) (*types.Transaction, error) {
	return _AddressProviderv3.Contract.SetAddress(&_AddressProviderv3.TransactOpts, key, value, saveVersion)
}

// AddressProviderv3SetAddressIterator is returned from FilterSetAddress and is used to iterate over the raw logs and unpacked data for SetAddress events raised by the AddressProviderv3 contract.
type AddressProviderv3SetAddressIterator struct {
	Event *AddressProviderv3SetAddress // Event containing the contract specifics and raw log

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
func (it *AddressProviderv3SetAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressProviderv3SetAddress)
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
		it.Event = new(AddressProviderv3SetAddress)
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
func (it *AddressProviderv3SetAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressProviderv3SetAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressProviderv3SetAddress represents a SetAddress event raised by the AddressProviderv3 contract.
type AddressProviderv3SetAddress struct {
	Key     [32]byte
	Value   common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAddress is a free log retrieval operation binding the contract event 0xb0c728d7569a96de0c070aa765819f6e665acbc3d4fa293440dd65c8c3e8b5ff.
//
// Solidity: event SetAddress(bytes32 indexed key, address indexed value, uint256 indexed version)
func (_AddressProviderv3 *AddressProviderv3Filterer) FilterSetAddress(opts *bind.FilterOpts, key [][32]byte, value []common.Address, version []*big.Int) (*AddressProviderv3SetAddressIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _AddressProviderv3.contract.FilterLogs(opts, "SetAddress", keyRule, valueRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &AddressProviderv3SetAddressIterator{contract: _AddressProviderv3.contract, event: "SetAddress", logs: logs, sub: sub}, nil
}

// WatchSetAddress is a free log subscription operation binding the contract event 0xb0c728d7569a96de0c070aa765819f6e665acbc3d4fa293440dd65c8c3e8b5ff.
//
// Solidity: event SetAddress(bytes32 indexed key, address indexed value, uint256 indexed version)
func (_AddressProviderv3 *AddressProviderv3Filterer) WatchSetAddress(opts *bind.WatchOpts, sink chan<- *AddressProviderv3SetAddress, key [][32]byte, value []common.Address, version []*big.Int) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _AddressProviderv3.contract.WatchLogs(opts, "SetAddress", keyRule, valueRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressProviderv3SetAddress)
				if err := _AddressProviderv3.contract.UnpackLog(event, "SetAddress", log); err != nil {
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

// ParseSetAddress is a log parse operation binding the contract event 0xb0c728d7569a96de0c070aa765819f6e665acbc3d4fa293440dd65c8c3e8b5ff.
//
// Solidity: event SetAddress(bytes32 indexed key, address indexed value, uint256 indexed version)
func (_AddressProviderv3 *AddressProviderv3Filterer) ParseSetAddress(log types.Log) (*AddressProviderv3SetAddress, error) {
	event := new(AddressProviderv3SetAddress)
	if err := _AddressProviderv3.contract.UnpackLog(event, "SetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
