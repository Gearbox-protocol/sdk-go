// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gaugev3

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

// Gaugev3MetaData contains all meta data concerning the Gaugev3 contract.
var Gaugev3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gearStaking\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addQuotaToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeQuotaMaxRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeQuotaMinRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochFrozen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochLastUpdate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRates\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"rates\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenAdded\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quotaRateParams\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"minRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalVotesLpSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"totalVotesCaSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFrozenEpoch\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unvote\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userTokenVotes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"votesLpSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"votesCaSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddQuotaToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minRate\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFrozenEpoch\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetQuotaTokenParams\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"minRate\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unvote\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"lpSide\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateEpoch\",\"inputs\":[{\"name\":\"epochNow\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Vote\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"lpSide\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressIsNotContractException\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerNotConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotControllerOrConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotVoterException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectParameterException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientVotesException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotAllowedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Gaugev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Gaugev3MetaData.ABI instead.
var Gaugev3ABI = Gaugev3MetaData.ABI

// Gaugev3 is an auto generated Go binding around an Ethereum contract.
type Gaugev3 struct {
	Gaugev3Caller     // Read-only binding to the contract
	Gaugev3Transactor // Write-only binding to the contract
	Gaugev3Filterer   // Log filterer for contract events
}

// Gaugev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Gaugev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gaugev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Gaugev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gaugev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Gaugev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Gaugev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Gaugev3Session struct {
	Contract     *Gaugev3          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Gaugev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Gaugev3CallerSession struct {
	Contract *Gaugev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Gaugev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Gaugev3TransactorSession struct {
	Contract     *Gaugev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Gaugev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Gaugev3Raw struct {
	Contract *Gaugev3 // Generic contract binding to access the raw methods on
}

// Gaugev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Gaugev3CallerRaw struct {
	Contract *Gaugev3Caller // Generic read-only contract binding to access the raw methods on
}

// Gaugev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Gaugev3TransactorRaw struct {
	Contract *Gaugev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGaugev3 creates a new instance of Gaugev3, bound to a specific deployed contract.
func NewGaugev3(address common.Address, backend bind.ContractBackend) (*Gaugev3, error) {
	contract, err := bindGaugev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gaugev3{Gaugev3Caller: Gaugev3Caller{contract: contract}, Gaugev3Transactor: Gaugev3Transactor{contract: contract}, Gaugev3Filterer: Gaugev3Filterer{contract: contract}}, nil
}

// NewGaugev3Caller creates a new read-only instance of Gaugev3, bound to a specific deployed contract.
func NewGaugev3Caller(address common.Address, caller bind.ContractCaller) (*Gaugev3Caller, error) {
	contract, err := bindGaugev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Gaugev3Caller{contract: contract}, nil
}

// NewGaugev3Transactor creates a new write-only instance of Gaugev3, bound to a specific deployed contract.
func NewGaugev3Transactor(address common.Address, transactor bind.ContractTransactor) (*Gaugev3Transactor, error) {
	contract, err := bindGaugev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Gaugev3Transactor{contract: contract}, nil
}

// NewGaugev3Filterer creates a new log filterer instance of Gaugev3, bound to a specific deployed contract.
func NewGaugev3Filterer(address common.Address, filterer bind.ContractFilterer) (*Gaugev3Filterer, error) {
	contract, err := bindGaugev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Gaugev3Filterer{contract: contract}, nil
}

// bindGaugev3 binds a generic wrapper to an already deployed contract.
func bindGaugev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Gaugev3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gaugev3 *Gaugev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gaugev3.Contract.Gaugev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gaugev3 *Gaugev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gaugev3.Contract.Gaugev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gaugev3 *Gaugev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gaugev3.Contract.Gaugev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gaugev3 *Gaugev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gaugev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gaugev3 *Gaugev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gaugev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gaugev3 *Gaugev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gaugev3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Gaugev3 *Gaugev3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Gaugev3 *Gaugev3Session) Acl() (common.Address, error) {
	return _Gaugev3.Contract.Acl(&_Gaugev3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Gaugev3 *Gaugev3CallerSession) Acl() (common.Address, error) {
	return _Gaugev3.Contract.Acl(&_Gaugev3.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Gaugev3 *Gaugev3Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Gaugev3 *Gaugev3Session) ContractType() ([32]byte, error) {
	return _Gaugev3.Contract.ContractType(&_Gaugev3.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Gaugev3 *Gaugev3CallerSession) ContractType() ([32]byte, error) {
	return _Gaugev3.Contract.ContractType(&_Gaugev3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gaugev3 *Gaugev3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gaugev3 *Gaugev3Session) Controller() (common.Address, error) {
	return _Gaugev3.Contract.Controller(&_Gaugev3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Gaugev3 *Gaugev3CallerSession) Controller() (common.Address, error) {
	return _Gaugev3.Contract.Controller(&_Gaugev3.CallOpts)
}

// EpochFrozen is a free data retrieval call binding the contract method 0xf9534828.
//
// Solidity: function epochFrozen() view returns(bool)
func (_Gaugev3 *Gaugev3Caller) EpochFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "epochFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochFrozen is a free data retrieval call binding the contract method 0xf9534828.
//
// Solidity: function epochFrozen() view returns(bool)
func (_Gaugev3 *Gaugev3Session) EpochFrozen() (bool, error) {
	return _Gaugev3.Contract.EpochFrozen(&_Gaugev3.CallOpts)
}

// EpochFrozen is a free data retrieval call binding the contract method 0xf9534828.
//
// Solidity: function epochFrozen() view returns(bool)
func (_Gaugev3 *Gaugev3CallerSession) EpochFrozen() (bool, error) {
	return _Gaugev3.Contract.EpochFrozen(&_Gaugev3.CallOpts)
}

// EpochLastUpdate is a free data retrieval call binding the contract method 0xfb832c71.
//
// Solidity: function epochLastUpdate() view returns(uint16)
func (_Gaugev3 *Gaugev3Caller) EpochLastUpdate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "epochLastUpdate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// EpochLastUpdate is a free data retrieval call binding the contract method 0xfb832c71.
//
// Solidity: function epochLastUpdate() view returns(uint16)
func (_Gaugev3 *Gaugev3Session) EpochLastUpdate() (uint16, error) {
	return _Gaugev3.Contract.EpochLastUpdate(&_Gaugev3.CallOpts)
}

// EpochLastUpdate is a free data retrieval call binding the contract method 0xfb832c71.
//
// Solidity: function epochLastUpdate() view returns(uint16)
func (_Gaugev3 *Gaugev3CallerSession) EpochLastUpdate() (uint16, error) {
	return _Gaugev3.Contract.EpochLastUpdate(&_Gaugev3.CallOpts)
}

// GetRates is a free data retrieval call binding the contract method 0x67bd79a2.
//
// Solidity: function getRates(address[] tokens) view returns(uint16[] rates)
func (_Gaugev3 *Gaugev3Caller) GetRates(opts *bind.CallOpts, tokens []common.Address) ([]uint16, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "getRates", tokens)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetRates is a free data retrieval call binding the contract method 0x67bd79a2.
//
// Solidity: function getRates(address[] tokens) view returns(uint16[] rates)
func (_Gaugev3 *Gaugev3Session) GetRates(tokens []common.Address) ([]uint16, error) {
	return _Gaugev3.Contract.GetRates(&_Gaugev3.CallOpts, tokens)
}

// GetRates is a free data retrieval call binding the contract method 0x67bd79a2.
//
// Solidity: function getRates(address[] tokens) view returns(uint16[] rates)
func (_Gaugev3 *Gaugev3CallerSession) GetRates(tokens []common.Address) ([]uint16, error) {
	return _Gaugev3.Contract.GetRates(&_Gaugev3.CallOpts, tokens)
}

// IsTokenAdded is a free data retrieval call binding the contract method 0xa36532b2.
//
// Solidity: function isTokenAdded(address token) view returns(bool)
func (_Gaugev3 *Gaugev3Caller) IsTokenAdded(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "isTokenAdded", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAdded is a free data retrieval call binding the contract method 0xa36532b2.
//
// Solidity: function isTokenAdded(address token) view returns(bool)
func (_Gaugev3 *Gaugev3Session) IsTokenAdded(token common.Address) (bool, error) {
	return _Gaugev3.Contract.IsTokenAdded(&_Gaugev3.CallOpts, token)
}

// IsTokenAdded is a free data retrieval call binding the contract method 0xa36532b2.
//
// Solidity: function isTokenAdded(address token) view returns(bool)
func (_Gaugev3 *Gaugev3CallerSession) IsTokenAdded(token common.Address) (bool, error) {
	return _Gaugev3.Contract.IsTokenAdded(&_Gaugev3.CallOpts, token)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Gaugev3 *Gaugev3Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Gaugev3 *Gaugev3Session) Pool() (common.Address, error) {
	return _Gaugev3.Contract.Pool(&_Gaugev3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Gaugev3 *Gaugev3CallerSession) Pool() (common.Address, error) {
	return _Gaugev3.Contract.Pool(&_Gaugev3.CallOpts)
}

// QuotaRateParams is a free data retrieval call binding the contract method 0x0b647622.
//
// Solidity: function quotaRateParams(address ) view returns(uint16 minRate, uint16 maxRate, uint96 totalVotesLpSide, uint96 totalVotesCaSide)
func (_Gaugev3 *Gaugev3Caller) QuotaRateParams(opts *bind.CallOpts, arg0 common.Address) (struct {
	MinRate          uint16
	MaxRate          uint16
	TotalVotesLpSide *big.Int
	TotalVotesCaSide *big.Int
}, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "quotaRateParams", arg0)

	outstruct := new(struct {
		MinRate          uint16
		MaxRate          uint16
		TotalVotesLpSide *big.Int
		TotalVotesCaSide *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinRate = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxRate = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.TotalVotesLpSide = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalVotesCaSide = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuotaRateParams is a free data retrieval call binding the contract method 0x0b647622.
//
// Solidity: function quotaRateParams(address ) view returns(uint16 minRate, uint16 maxRate, uint96 totalVotesLpSide, uint96 totalVotesCaSide)
func (_Gaugev3 *Gaugev3Session) QuotaRateParams(arg0 common.Address) (struct {
	MinRate          uint16
	MaxRate          uint16
	TotalVotesLpSide *big.Int
	TotalVotesCaSide *big.Int
}, error) {
	return _Gaugev3.Contract.QuotaRateParams(&_Gaugev3.CallOpts, arg0)
}

// QuotaRateParams is a free data retrieval call binding the contract method 0x0b647622.
//
// Solidity: function quotaRateParams(address ) view returns(uint16 minRate, uint16 maxRate, uint96 totalVotesLpSide, uint96 totalVotesCaSide)
func (_Gaugev3 *Gaugev3CallerSession) QuotaRateParams(arg0 common.Address) (struct {
	MinRate          uint16
	MaxRate          uint16
	TotalVotesLpSide *big.Int
	TotalVotesCaSide *big.Int
}, error) {
	return _Gaugev3.Contract.QuotaRateParams(&_Gaugev3.CallOpts, arg0)
}

// UserTokenVotes is a free data retrieval call binding the contract method 0xcc10ad86.
//
// Solidity: function userTokenVotes(address , address ) view returns(uint96 votesLpSide, uint96 votesCaSide)
func (_Gaugev3 *Gaugev3Caller) UserTokenVotes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	VotesLpSide *big.Int
	VotesCaSide *big.Int
}, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "userTokenVotes", arg0, arg1)

	outstruct := new(struct {
		VotesLpSide *big.Int
		VotesCaSide *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VotesLpSide = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VotesCaSide = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserTokenVotes is a free data retrieval call binding the contract method 0xcc10ad86.
//
// Solidity: function userTokenVotes(address , address ) view returns(uint96 votesLpSide, uint96 votesCaSide)
func (_Gaugev3 *Gaugev3Session) UserTokenVotes(arg0 common.Address, arg1 common.Address) (struct {
	VotesLpSide *big.Int
	VotesCaSide *big.Int
}, error) {
	return _Gaugev3.Contract.UserTokenVotes(&_Gaugev3.CallOpts, arg0, arg1)
}

// UserTokenVotes is a free data retrieval call binding the contract method 0xcc10ad86.
//
// Solidity: function userTokenVotes(address , address ) view returns(uint96 votesLpSide, uint96 votesCaSide)
func (_Gaugev3 *Gaugev3CallerSession) UserTokenVotes(arg0 common.Address, arg1 common.Address) (struct {
	VotesLpSide *big.Int
	VotesCaSide *big.Int
}, error) {
	return _Gaugev3.Contract.UserTokenVotes(&_Gaugev3.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Gaugev3 *Gaugev3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Gaugev3 *Gaugev3Session) Version() (*big.Int, error) {
	return _Gaugev3.Contract.Version(&_Gaugev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Gaugev3 *Gaugev3CallerSession) Version() (*big.Int, error) {
	return _Gaugev3.Contract.Version(&_Gaugev3.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gaugev3 *Gaugev3Caller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gaugev3.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gaugev3 *Gaugev3Session) Voter() (common.Address, error) {
	return _Gaugev3.Contract.Voter(&_Gaugev3.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Gaugev3 *Gaugev3CallerSession) Voter() (common.Address, error) {
	return _Gaugev3.Contract.Voter(&_Gaugev3.CallOpts)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x359cd5bf.
//
// Solidity: function addQuotaToken(address token, uint16 minRate, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3Transactor) AddQuotaToken(opts *bind.TransactOpts, token common.Address, minRate uint16, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "addQuotaToken", token, minRate, maxRate)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x359cd5bf.
//
// Solidity: function addQuotaToken(address token, uint16 minRate, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3Session) AddQuotaToken(token common.Address, minRate uint16, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.AddQuotaToken(&_Gaugev3.TransactOpts, token, minRate, maxRate)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x359cd5bf.
//
// Solidity: function addQuotaToken(address token, uint16 minRate, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3TransactorSession) AddQuotaToken(token common.Address, minRate uint16, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.AddQuotaToken(&_Gaugev3.TransactOpts, token, minRate, maxRate)
}

// ChangeQuotaMaxRate is a paid mutator transaction binding the contract method 0x32f8e50c.
//
// Solidity: function changeQuotaMaxRate(address token, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3Transactor) ChangeQuotaMaxRate(opts *bind.TransactOpts, token common.Address, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "changeQuotaMaxRate", token, maxRate)
}

// ChangeQuotaMaxRate is a paid mutator transaction binding the contract method 0x32f8e50c.
//
// Solidity: function changeQuotaMaxRate(address token, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3Session) ChangeQuotaMaxRate(token common.Address, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.ChangeQuotaMaxRate(&_Gaugev3.TransactOpts, token, maxRate)
}

// ChangeQuotaMaxRate is a paid mutator transaction binding the contract method 0x32f8e50c.
//
// Solidity: function changeQuotaMaxRate(address token, uint16 maxRate) returns()
func (_Gaugev3 *Gaugev3TransactorSession) ChangeQuotaMaxRate(token common.Address, maxRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.ChangeQuotaMaxRate(&_Gaugev3.TransactOpts, token, maxRate)
}

// ChangeQuotaMinRate is a paid mutator transaction binding the contract method 0xb2863529.
//
// Solidity: function changeQuotaMinRate(address token, uint16 minRate) returns()
func (_Gaugev3 *Gaugev3Transactor) ChangeQuotaMinRate(opts *bind.TransactOpts, token common.Address, minRate uint16) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "changeQuotaMinRate", token, minRate)
}

// ChangeQuotaMinRate is a paid mutator transaction binding the contract method 0xb2863529.
//
// Solidity: function changeQuotaMinRate(address token, uint16 minRate) returns()
func (_Gaugev3 *Gaugev3Session) ChangeQuotaMinRate(token common.Address, minRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.ChangeQuotaMinRate(&_Gaugev3.TransactOpts, token, minRate)
}

// ChangeQuotaMinRate is a paid mutator transaction binding the contract method 0xb2863529.
//
// Solidity: function changeQuotaMinRate(address token, uint16 minRate) returns()
func (_Gaugev3 *Gaugev3TransactorSession) ChangeQuotaMinRate(token common.Address, minRate uint16) (*types.Transaction, error) {
	return _Gaugev3.Contract.ChangeQuotaMinRate(&_Gaugev3.TransactOpts, token, minRate)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Gaugev3 *Gaugev3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Gaugev3 *Gaugev3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _Gaugev3.Contract.SetController(&_Gaugev3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Gaugev3 *Gaugev3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _Gaugev3.Contract.SetController(&_Gaugev3.TransactOpts, newController)
}

// SetFrozenEpoch is a paid mutator transaction binding the contract method 0xed519474.
//
// Solidity: function setFrozenEpoch(bool status) returns()
func (_Gaugev3 *Gaugev3Transactor) SetFrozenEpoch(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "setFrozenEpoch", status)
}

// SetFrozenEpoch is a paid mutator transaction binding the contract method 0xed519474.
//
// Solidity: function setFrozenEpoch(bool status) returns()
func (_Gaugev3 *Gaugev3Session) SetFrozenEpoch(status bool) (*types.Transaction, error) {
	return _Gaugev3.Contract.SetFrozenEpoch(&_Gaugev3.TransactOpts, status)
}

// SetFrozenEpoch is a paid mutator transaction binding the contract method 0xed519474.
//
// Solidity: function setFrozenEpoch(bool status) returns()
func (_Gaugev3 *Gaugev3TransactorSession) SetFrozenEpoch(status bool) (*types.Transaction, error) {
	return _Gaugev3.Contract.SetFrozenEpoch(&_Gaugev3.TransactOpts, status)
}

// Unvote is a paid mutator transaction binding the contract method 0x102418f3.
//
// Solidity: function unvote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3Transactor) Unvote(opts *bind.TransactOpts, user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "unvote", user, votes, extraData)
}

// Unvote is a paid mutator transaction binding the contract method 0x102418f3.
//
// Solidity: function unvote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3Session) Unvote(user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.Contract.Unvote(&_Gaugev3.TransactOpts, user, votes, extraData)
}

// Unvote is a paid mutator transaction binding the contract method 0x102418f3.
//
// Solidity: function unvote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3TransactorSession) Unvote(user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.Contract.Unvote(&_Gaugev3.TransactOpts, user, votes, extraData)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_Gaugev3 *Gaugev3Transactor) UpdateEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "updateEpoch")
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_Gaugev3 *Gaugev3Session) UpdateEpoch() (*types.Transaction, error) {
	return _Gaugev3.Contract.UpdateEpoch(&_Gaugev3.TransactOpts)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0x36f4fb02.
//
// Solidity: function updateEpoch() returns()
func (_Gaugev3 *Gaugev3TransactorSession) UpdateEpoch() (*types.Transaction, error) {
	return _Gaugev3.Contract.UpdateEpoch(&_Gaugev3.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x3c4f98ac.
//
// Solidity: function vote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3Transactor) Vote(opts *bind.TransactOpts, user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.contract.Transact(opts, "vote", user, votes, extraData)
}

// Vote is a paid mutator transaction binding the contract method 0x3c4f98ac.
//
// Solidity: function vote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3Session) Vote(user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.Contract.Vote(&_Gaugev3.TransactOpts, user, votes, extraData)
}

// Vote is a paid mutator transaction binding the contract method 0x3c4f98ac.
//
// Solidity: function vote(address user, uint96 votes, bytes extraData) returns()
func (_Gaugev3 *Gaugev3TransactorSession) Vote(user common.Address, votes *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Gaugev3.Contract.Vote(&_Gaugev3.TransactOpts, user, votes, extraData)
}

// Gaugev3AddQuotaTokenIterator is returned from FilterAddQuotaToken and is used to iterate over the raw logs and unpacked data for AddQuotaToken events raised by the Gaugev3 contract.
type Gaugev3AddQuotaTokenIterator struct {
	Event *Gaugev3AddQuotaToken // Event containing the contract specifics and raw log

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
func (it *Gaugev3AddQuotaTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3AddQuotaToken)
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
		it.Event = new(Gaugev3AddQuotaToken)
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
func (it *Gaugev3AddQuotaTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3AddQuotaTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3AddQuotaToken represents a AddQuotaToken event raised by the Gaugev3 contract.
type Gaugev3AddQuotaToken struct {
	Token   common.Address
	MinRate uint16
	MaxRate uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddQuotaToken is a free log retrieval operation binding the contract event 0x26ed208f71237782f945612371d1e179300898e965c647fd1d23c97bdc973812.
//
// Solidity: event AddQuotaToken(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) FilterAddQuotaToken(opts *bind.FilterOpts, token []common.Address) (*Gaugev3AddQuotaTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "AddQuotaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &Gaugev3AddQuotaTokenIterator{contract: _Gaugev3.contract, event: "AddQuotaToken", logs: logs, sub: sub}, nil
}

// WatchAddQuotaToken is a free log subscription operation binding the contract event 0x26ed208f71237782f945612371d1e179300898e965c647fd1d23c97bdc973812.
//
// Solidity: event AddQuotaToken(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) WatchAddQuotaToken(opts *bind.WatchOpts, sink chan<- *Gaugev3AddQuotaToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "AddQuotaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3AddQuotaToken)
				if err := _Gaugev3.contract.UnpackLog(event, "AddQuotaToken", log); err != nil {
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

// ParseAddQuotaToken is a log parse operation binding the contract event 0x26ed208f71237782f945612371d1e179300898e965c647fd1d23c97bdc973812.
//
// Solidity: event AddQuotaToken(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) ParseAddQuotaToken(log types.Log) (*Gaugev3AddQuotaToken, error) {
	event := new(Gaugev3AddQuotaToken)
	if err := _Gaugev3.contract.UnpackLog(event, "AddQuotaToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the Gaugev3 contract.
type Gaugev3NewControllerIterator struct {
	Event *Gaugev3NewController // Event containing the contract specifics and raw log

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
func (it *Gaugev3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3NewController)
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
		it.Event = new(Gaugev3NewController)
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
func (it *Gaugev3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3NewController represents a NewController event raised by the Gaugev3 contract.
type Gaugev3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_Gaugev3 *Gaugev3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*Gaugev3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &Gaugev3NewControllerIterator{contract: _Gaugev3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_Gaugev3 *Gaugev3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *Gaugev3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3NewController)
				if err := _Gaugev3.contract.UnpackLog(event, "NewController", log); err != nil {
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

// ParseNewController is a log parse operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_Gaugev3 *Gaugev3Filterer) ParseNewController(log types.Log) (*Gaugev3NewController, error) {
	event := new(Gaugev3NewController)
	if err := _Gaugev3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3SetFrozenEpochIterator is returned from FilterSetFrozenEpoch and is used to iterate over the raw logs and unpacked data for SetFrozenEpoch events raised by the Gaugev3 contract.
type Gaugev3SetFrozenEpochIterator struct {
	Event *Gaugev3SetFrozenEpoch // Event containing the contract specifics and raw log

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
func (it *Gaugev3SetFrozenEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3SetFrozenEpoch)
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
		it.Event = new(Gaugev3SetFrozenEpoch)
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
func (it *Gaugev3SetFrozenEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3SetFrozenEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3SetFrozenEpoch represents a SetFrozenEpoch event raised by the Gaugev3 contract.
type Gaugev3SetFrozenEpoch struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFrozenEpoch is a free log retrieval operation binding the contract event 0xa60fade018f4d462a5daa5f95377789f4577d1e71865ab6f46e65c8bcfd91d5b.
//
// Solidity: event SetFrozenEpoch(bool status)
func (_Gaugev3 *Gaugev3Filterer) FilterSetFrozenEpoch(opts *bind.FilterOpts) (*Gaugev3SetFrozenEpochIterator, error) {

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "SetFrozenEpoch")
	if err != nil {
		return nil, err
	}
	return &Gaugev3SetFrozenEpochIterator{contract: _Gaugev3.contract, event: "SetFrozenEpoch", logs: logs, sub: sub}, nil
}

// WatchSetFrozenEpoch is a free log subscription operation binding the contract event 0xa60fade018f4d462a5daa5f95377789f4577d1e71865ab6f46e65c8bcfd91d5b.
//
// Solidity: event SetFrozenEpoch(bool status)
func (_Gaugev3 *Gaugev3Filterer) WatchSetFrozenEpoch(opts *bind.WatchOpts, sink chan<- *Gaugev3SetFrozenEpoch) (event.Subscription, error) {

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "SetFrozenEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3SetFrozenEpoch)
				if err := _Gaugev3.contract.UnpackLog(event, "SetFrozenEpoch", log); err != nil {
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

// ParseSetFrozenEpoch is a log parse operation binding the contract event 0xa60fade018f4d462a5daa5f95377789f4577d1e71865ab6f46e65c8bcfd91d5b.
//
// Solidity: event SetFrozenEpoch(bool status)
func (_Gaugev3 *Gaugev3Filterer) ParseSetFrozenEpoch(log types.Log) (*Gaugev3SetFrozenEpoch, error) {
	event := new(Gaugev3SetFrozenEpoch)
	if err := _Gaugev3.contract.UnpackLog(event, "SetFrozenEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3SetQuotaTokenParamsIterator is returned from FilterSetQuotaTokenParams and is used to iterate over the raw logs and unpacked data for SetQuotaTokenParams events raised by the Gaugev3 contract.
type Gaugev3SetQuotaTokenParamsIterator struct {
	Event *Gaugev3SetQuotaTokenParams // Event containing the contract specifics and raw log

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
func (it *Gaugev3SetQuotaTokenParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3SetQuotaTokenParams)
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
		it.Event = new(Gaugev3SetQuotaTokenParams)
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
func (it *Gaugev3SetQuotaTokenParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3SetQuotaTokenParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3SetQuotaTokenParams represents a SetQuotaTokenParams event raised by the Gaugev3 contract.
type Gaugev3SetQuotaTokenParams struct {
	Token   common.Address
	MinRate uint16
	MaxRate uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetQuotaTokenParams is a free log retrieval operation binding the contract event 0xce0a212664f433711432d2fcd73ba6b7395bc67a540150eecf21c37b71c26b00.
//
// Solidity: event SetQuotaTokenParams(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) FilterSetQuotaTokenParams(opts *bind.FilterOpts, token []common.Address) (*Gaugev3SetQuotaTokenParamsIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "SetQuotaTokenParams", tokenRule)
	if err != nil {
		return nil, err
	}
	return &Gaugev3SetQuotaTokenParamsIterator{contract: _Gaugev3.contract, event: "SetQuotaTokenParams", logs: logs, sub: sub}, nil
}

// WatchSetQuotaTokenParams is a free log subscription operation binding the contract event 0xce0a212664f433711432d2fcd73ba6b7395bc67a540150eecf21c37b71c26b00.
//
// Solidity: event SetQuotaTokenParams(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) WatchSetQuotaTokenParams(opts *bind.WatchOpts, sink chan<- *Gaugev3SetQuotaTokenParams, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "SetQuotaTokenParams", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3SetQuotaTokenParams)
				if err := _Gaugev3.contract.UnpackLog(event, "SetQuotaTokenParams", log); err != nil {
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

// ParseSetQuotaTokenParams is a log parse operation binding the contract event 0xce0a212664f433711432d2fcd73ba6b7395bc67a540150eecf21c37b71c26b00.
//
// Solidity: event SetQuotaTokenParams(address indexed token, uint16 minRate, uint16 maxRate)
func (_Gaugev3 *Gaugev3Filterer) ParseSetQuotaTokenParams(log types.Log) (*Gaugev3SetQuotaTokenParams, error) {
	event := new(Gaugev3SetQuotaTokenParams)
	if err := _Gaugev3.contract.UnpackLog(event, "SetQuotaTokenParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3UnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the Gaugev3 contract.
type Gaugev3UnvoteIterator struct {
	Event *Gaugev3Unvote // Event containing the contract specifics and raw log

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
func (it *Gaugev3UnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3Unvote)
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
		it.Event = new(Gaugev3Unvote)
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
func (it *Gaugev3UnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3UnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3Unvote represents a Unvote event raised by the Gaugev3 contract.
type Gaugev3Unvote struct {
	User   common.Address
	Token  common.Address
	Votes  *big.Int
	LpSide bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xb4cfba462215396ee513d049a336a7a90798f99a4c100a605f6f5b76c9f47d21.
//
// Solidity: event Unvote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) FilterUnvote(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*Gaugev3UnvoteIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "Unvote", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Gaugev3UnvoteIterator{contract: _Gaugev3.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xb4cfba462215396ee513d049a336a7a90798f99a4c100a605f6f5b76c9f47d21.
//
// Solidity: event Unvote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *Gaugev3Unvote, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "Unvote", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3Unvote)
				if err := _Gaugev3.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// ParseUnvote is a log parse operation binding the contract event 0xb4cfba462215396ee513d049a336a7a90798f99a4c100a605f6f5b76c9f47d21.
//
// Solidity: event Unvote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) ParseUnvote(log types.Log) (*Gaugev3Unvote, error) {
	event := new(Gaugev3Unvote)
	if err := _Gaugev3.contract.UnpackLog(event, "Unvote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3UpdateEpochIterator is returned from FilterUpdateEpoch and is used to iterate over the raw logs and unpacked data for UpdateEpoch events raised by the Gaugev3 contract.
type Gaugev3UpdateEpochIterator struct {
	Event *Gaugev3UpdateEpoch // Event containing the contract specifics and raw log

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
func (it *Gaugev3UpdateEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3UpdateEpoch)
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
		it.Event = new(Gaugev3UpdateEpoch)
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
func (it *Gaugev3UpdateEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3UpdateEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3UpdateEpoch represents a UpdateEpoch event raised by the Gaugev3 contract.
type Gaugev3UpdateEpoch struct {
	EpochNow uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateEpoch is a free log retrieval operation binding the contract event 0x44301d5732d11f39569dd7cfde533dacce079da5aa37171ab8a9cbf098818b52.
//
// Solidity: event UpdateEpoch(uint16 epochNow)
func (_Gaugev3 *Gaugev3Filterer) FilterUpdateEpoch(opts *bind.FilterOpts) (*Gaugev3UpdateEpochIterator, error) {

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "UpdateEpoch")
	if err != nil {
		return nil, err
	}
	return &Gaugev3UpdateEpochIterator{contract: _Gaugev3.contract, event: "UpdateEpoch", logs: logs, sub: sub}, nil
}

// WatchUpdateEpoch is a free log subscription operation binding the contract event 0x44301d5732d11f39569dd7cfde533dacce079da5aa37171ab8a9cbf098818b52.
//
// Solidity: event UpdateEpoch(uint16 epochNow)
func (_Gaugev3 *Gaugev3Filterer) WatchUpdateEpoch(opts *bind.WatchOpts, sink chan<- *Gaugev3UpdateEpoch) (event.Subscription, error) {

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "UpdateEpoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3UpdateEpoch)
				if err := _Gaugev3.contract.UnpackLog(event, "UpdateEpoch", log); err != nil {
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

// ParseUpdateEpoch is a log parse operation binding the contract event 0x44301d5732d11f39569dd7cfde533dacce079da5aa37171ab8a9cbf098818b52.
//
// Solidity: event UpdateEpoch(uint16 epochNow)
func (_Gaugev3 *Gaugev3Filterer) ParseUpdateEpoch(log types.Log) (*Gaugev3UpdateEpoch, error) {
	event := new(Gaugev3UpdateEpoch)
	if err := _Gaugev3.contract.UnpackLog(event, "UpdateEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Gaugev3VoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the Gaugev3 contract.
type Gaugev3VoteIterator struct {
	Event *Gaugev3Vote // Event containing the contract specifics and raw log

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
func (it *Gaugev3VoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Gaugev3Vote)
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
		it.Event = new(Gaugev3Vote)
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
func (it *Gaugev3VoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Gaugev3VoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Gaugev3Vote represents a Vote event raised by the Gaugev3 contract.
type Gaugev3Vote struct {
	User   common.Address
	Token  common.Address
	Votes  *big.Int
	LpSide bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x720dd6f175f68c73fb470a1d57e611c555e7f963fba76f6b9fa7f75daa59d176.
//
// Solidity: event Vote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) FilterVote(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*Gaugev3VoteIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.FilterLogs(opts, "Vote", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &Gaugev3VoteIterator{contract: _Gaugev3.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x720dd6f175f68c73fb470a1d57e611c555e7f963fba76f6b9fa7f75daa59d176.
//
// Solidity: event Vote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) WatchVote(opts *bind.WatchOpts, sink chan<- *Gaugev3Vote, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Gaugev3.contract.WatchLogs(opts, "Vote", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Gaugev3Vote)
				if err := _Gaugev3.contract.UnpackLog(event, "Vote", log); err != nil {
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

// ParseVote is a log parse operation binding the contract event 0x720dd6f175f68c73fb470a1d57e611c555e7f963fba76f6b9fa7f75daa59d176.
//
// Solidity: event Vote(address indexed user, address indexed token, uint96 votes, bool lpSide)
func (_Gaugev3 *Gaugev3Filterer) ParseVote(log types.Log) (*Gaugev3Vote, error) {
	event := new(Gaugev3Vote)
	if err := _Gaugev3.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
