// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolQuotaKeeperv3

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

// PoolQuotaKeeperv3MetaData contains all meta data concerning the PoolQuotaKeeperv3 contract.
var PoolQuotaKeeperv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotControllerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotCreditManagerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotGaugeException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnpausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleCreditManagerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuotaIsOutOfBoundsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredCreditManagerOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyAddedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotQuotedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"AddCreditManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AddQuotaToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGauge\",\"type\":\"address\"}],\"name\":\"SetGauge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"name\":\"SetQuotaIncreaseFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"}],\"name\":\"SetTokenLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int96\",\"name\":\"quotaChange\",\"type\":\"int96\"}],\"name\":\"UpdateQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"UpdateTokenQuotaRate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"accrueQuotaInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"addCreditManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addQuotaToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"cumulativeIndex\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gauge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getQuota\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"quota\",\"type\":\"uint96\"},{\"internalType\":\"uint192\",\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getQuotaAndOutstandingInterest\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"quoted\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"outstandingInterest\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getQuotaRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenQuotaParams\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint192\",\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isQuotedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastQuotaRateUpdate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolQuotaRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quotaRevenue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quotedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"setLimitsToZero\",\"type\":\"bool\"}],\"name\":\"removeQuotas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"setGauge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"}],\"name\":\"setTokenLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"}],\"name\":\"setTokenQuotaIncreaseFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"int96\",\"name\":\"requestedChange\",\"type\":\"int96\"},{\"internalType\":\"uint96\",\"name\":\"minQuota\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"maxQuota\",\"type\":\"uint96\"}],\"name\":\"updateQuota\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"caQuotaInterestChange\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fees\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"enableToken\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"disableToken\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolQuotaKeeperv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolQuotaKeeperv3MetaData.ABI instead.
var PoolQuotaKeeperv3ABI = PoolQuotaKeeperv3MetaData.ABI

// PoolQuotaKeeperv3 is an auto generated Go binding around an Ethereum contract.
type PoolQuotaKeeperv3 struct {
	PoolQuotaKeeperv3Caller     // Read-only binding to the contract
	PoolQuotaKeeperv3Transactor // Write-only binding to the contract
	PoolQuotaKeeperv3Filterer   // Log filterer for contract events
}

// PoolQuotaKeeperv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolQuotaKeeperv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolQuotaKeeperv3Session struct {
	Contract     *PoolQuotaKeeperv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PoolQuotaKeeperv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolQuotaKeeperv3CallerSession struct {
	Contract *PoolQuotaKeeperv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PoolQuotaKeeperv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolQuotaKeeperv3TransactorSession struct {
	Contract     *PoolQuotaKeeperv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PoolQuotaKeeperv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Raw struct {
	Contract *PoolQuotaKeeperv3 // Generic contract binding to access the raw methods on
}

// PoolQuotaKeeperv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3CallerRaw struct {
	Contract *PoolQuotaKeeperv3Caller // Generic read-only contract binding to access the raw methods on
}

// PoolQuotaKeeperv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3TransactorRaw struct {
	Contract *PoolQuotaKeeperv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolQuotaKeeperv3 creates a new instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3(address common.Address, backend bind.ContractBackend) (*PoolQuotaKeeperv3, error) {
	contract, err := bindPoolQuotaKeeperv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3{PoolQuotaKeeperv3Caller: PoolQuotaKeeperv3Caller{contract: contract}, PoolQuotaKeeperv3Transactor: PoolQuotaKeeperv3Transactor{contract: contract}, PoolQuotaKeeperv3Filterer: PoolQuotaKeeperv3Filterer{contract: contract}}, nil
}

// NewPoolQuotaKeeperv3Caller creates a new read-only instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Caller(address common.Address, caller bind.ContractCaller) (*PoolQuotaKeeperv3Caller, error) {
	contract, err := bindPoolQuotaKeeperv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Caller{contract: contract}, nil
}

// NewPoolQuotaKeeperv3Transactor creates a new write-only instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Transactor(address common.Address, transactor bind.ContractTransactor) (*PoolQuotaKeeperv3Transactor, error) {
	contract, err := bindPoolQuotaKeeperv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Transactor{contract: contract}, nil
}

// NewPoolQuotaKeeperv3Filterer creates a new log filterer instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Filterer(address common.Address, filterer bind.ContractFilterer) (*PoolQuotaKeeperv3Filterer, error) {
	contract, err := bindPoolQuotaKeeperv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Filterer{contract: contract}, nil
}

// bindPoolQuotaKeeperv3 binds a generic wrapper to an already deployed contract.
func bindPoolQuotaKeeperv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolQuotaKeeperv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolQuotaKeeperv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Acl() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Acl(&_PoolQuotaKeeperv3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Acl() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Acl(&_PoolQuotaKeeperv3.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) ContractsRegister() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.ContractsRegister(&_PoolQuotaKeeperv3.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) ContractsRegister() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.ContractsRegister(&_PoolQuotaKeeperv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Controller() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Controller(&_PoolQuotaKeeperv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Controller() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Controller(&_PoolQuotaKeeperv3.CallOpts)
}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) CreditManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "creditManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) CreditManagers() ([]common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.CreditManagers(&_PoolQuotaKeeperv3.CallOpts)
}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) CreditManagers() ([]common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.CreditManagers(&_PoolQuotaKeeperv3.CallOpts)
}

// CumulativeIndex is a free data retrieval call binding the contract method 0xe08a03db.
//
// Solidity: function cumulativeIndex(address token) view returns(uint192)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) CumulativeIndex(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "cumulativeIndex", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeIndex is a free data retrieval call binding the contract method 0xe08a03db.
//
// Solidity: function cumulativeIndex(address token) view returns(uint192)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) CumulativeIndex(token common.Address) (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.CumulativeIndex(&_PoolQuotaKeeperv3.CallOpts, token)
}

// CumulativeIndex is a free data retrieval call binding the contract method 0xe08a03db.
//
// Solidity: function cumulativeIndex(address token) view returns(uint192)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) CumulativeIndex(token common.Address) (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.CumulativeIndex(&_PoolQuotaKeeperv3.CallOpts, token)
}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Gauge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "gauge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Gauge() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Gauge(&_PoolQuotaKeeperv3.CallOpts)
}

// Gauge is a free data retrieval call binding the contract method 0xa6f19c84.
//
// Solidity: function gauge() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Gauge() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Gauge(&_PoolQuotaKeeperv3.CallOpts)
}

// GetQuota is a free data retrieval call binding the contract method 0x26d6a2f4.
//
// Solidity: function getQuota(address creditAccount, address token) view returns(uint96 quota, uint192 cumulativeIndexLU)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) GetQuota(opts *bind.CallOpts, creditAccount common.Address, token common.Address) (struct {
	Quota             *big.Int
	CumulativeIndexLU *big.Int
}, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "getQuota", creditAccount, token)

	outstruct := new(struct {
		Quota             *big.Int
		CumulativeIndexLU *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quota = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CumulativeIndexLU = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetQuota is a free data retrieval call binding the contract method 0x26d6a2f4.
//
// Solidity: function getQuota(address creditAccount, address token) view returns(uint96 quota, uint192 cumulativeIndexLU)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) GetQuota(creditAccount common.Address, token common.Address) (struct {
	Quota             *big.Int
	CumulativeIndexLU *big.Int
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuota(&_PoolQuotaKeeperv3.CallOpts, creditAccount, token)
}

// GetQuota is a free data retrieval call binding the contract method 0x26d6a2f4.
//
// Solidity: function getQuota(address creditAccount, address token) view returns(uint96 quota, uint192 cumulativeIndexLU)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) GetQuota(creditAccount common.Address, token common.Address) (struct {
	Quota             *big.Int
	CumulativeIndexLU *big.Int
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuota(&_PoolQuotaKeeperv3.CallOpts, creditAccount, token)
}

// GetQuotaAndOutstandingInterest is a free data retrieval call binding the contract method 0xf3ef1813.
//
// Solidity: function getQuotaAndOutstandingInterest(address creditAccount, address token) view returns(uint96 quoted, uint128 outstandingInterest)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) GetQuotaAndOutstandingInterest(opts *bind.CallOpts, creditAccount common.Address, token common.Address) (struct {
	Quoted              *big.Int
	OutstandingInterest *big.Int
}, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "getQuotaAndOutstandingInterest", creditAccount, token)

	outstruct := new(struct {
		Quoted              *big.Int
		OutstandingInterest *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quoted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OutstandingInterest = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetQuotaAndOutstandingInterest is a free data retrieval call binding the contract method 0xf3ef1813.
//
// Solidity: function getQuotaAndOutstandingInterest(address creditAccount, address token) view returns(uint96 quoted, uint128 outstandingInterest)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) GetQuotaAndOutstandingInterest(creditAccount common.Address, token common.Address) (struct {
	Quoted              *big.Int
	OutstandingInterest *big.Int
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuotaAndOutstandingInterest(&_PoolQuotaKeeperv3.CallOpts, creditAccount, token)
}

// GetQuotaAndOutstandingInterest is a free data retrieval call binding the contract method 0xf3ef1813.
//
// Solidity: function getQuotaAndOutstandingInterest(address creditAccount, address token) view returns(uint96 quoted, uint128 outstandingInterest)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) GetQuotaAndOutstandingInterest(creditAccount common.Address, token common.Address) (struct {
	Quoted              *big.Int
	OutstandingInterest *big.Int
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuotaAndOutstandingInterest(&_PoolQuotaKeeperv3.CallOpts, creditAccount, token)
}

// GetQuotaRate is a free data retrieval call binding the contract method 0x0ab3640f.
//
// Solidity: function getQuotaRate(address token) view returns(uint16)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) GetQuotaRate(opts *bind.CallOpts, token common.Address) (uint16, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "getQuotaRate", token)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetQuotaRate is a free data retrieval call binding the contract method 0x0ab3640f.
//
// Solidity: function getQuotaRate(address token) view returns(uint16)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) GetQuotaRate(token common.Address) (uint16, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuotaRate(&_PoolQuotaKeeperv3.CallOpts, token)
}

// GetQuotaRate is a free data retrieval call binding the contract method 0x0ab3640f.
//
// Solidity: function getQuotaRate(address token) view returns(uint16)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) GetQuotaRate(token common.Address) (uint16, error) {
	return _PoolQuotaKeeperv3.Contract.GetQuotaRate(&_PoolQuotaKeeperv3.CallOpts, token)
}

// GetTokenQuotaParams is a free data retrieval call binding the contract method 0xbd42a06f.
//
// Solidity: function getTokenQuotaParams(address token) view returns(uint16 rate, uint192 cumulativeIndexLU, uint16 quotaIncreaseFee, uint96 totalQuoted, uint96 limit, bool isActive)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) GetTokenQuotaParams(opts *bind.CallOpts, token common.Address) (struct {
	Rate              uint16
	CumulativeIndexLU *big.Int
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
}, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "getTokenQuotaParams", token)

	outstruct := new(struct {
		Rate              uint16
		CumulativeIndexLU *big.Int
		QuotaIncreaseFee  uint16
		TotalQuoted       *big.Int
		Limit             *big.Int
		IsActive          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.CumulativeIndexLU = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuotaIncreaseFee = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TotalQuoted = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Limit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetTokenQuotaParams is a free data retrieval call binding the contract method 0xbd42a06f.
//
// Solidity: function getTokenQuotaParams(address token) view returns(uint16 rate, uint192 cumulativeIndexLU, uint16 quotaIncreaseFee, uint96 totalQuoted, uint96 limit, bool isActive)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) GetTokenQuotaParams(token common.Address) (struct {
	Rate              uint16
	CumulativeIndexLU *big.Int
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetTokenQuotaParams(&_PoolQuotaKeeperv3.CallOpts, token)
}

// GetTokenQuotaParams is a free data retrieval call binding the contract method 0xbd42a06f.
//
// Solidity: function getTokenQuotaParams(address token) view returns(uint16 rate, uint192 cumulativeIndexLU, uint16 quotaIncreaseFee, uint96 totalQuoted, uint96 limit, bool isActive)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) GetTokenQuotaParams(token common.Address) (struct {
	Rate              uint16
	CumulativeIndexLU *big.Int
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
}, error) {
	return _PoolQuotaKeeperv3.Contract.GetTokenQuotaParams(&_PoolQuotaKeeperv3.CallOpts, token)
}

// IsQuotedToken is a free data retrieval call binding the contract method 0xd9b94b06.
//
// Solidity: function isQuotedToken(address token) view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) IsQuotedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "isQuotedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuotedToken is a free data retrieval call binding the contract method 0xd9b94b06.
//
// Solidity: function isQuotedToken(address token) view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) IsQuotedToken(token common.Address) (bool, error) {
	return _PoolQuotaKeeperv3.Contract.IsQuotedToken(&_PoolQuotaKeeperv3.CallOpts, token)
}

// IsQuotedToken is a free data retrieval call binding the contract method 0xd9b94b06.
//
// Solidity: function isQuotedToken(address token) view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) IsQuotedToken(token common.Address) (bool, error) {
	return _PoolQuotaKeeperv3.Contract.IsQuotedToken(&_PoolQuotaKeeperv3.CallOpts, token)
}

// LastQuotaRateUpdate is a free data retrieval call binding the contract method 0x112024ff.
//
// Solidity: function lastQuotaRateUpdate() view returns(uint40)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) LastQuotaRateUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "lastQuotaRateUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastQuotaRateUpdate is a free data retrieval call binding the contract method 0x112024ff.
//
// Solidity: function lastQuotaRateUpdate() view returns(uint40)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) LastQuotaRateUpdate() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.LastQuotaRateUpdate(&_PoolQuotaKeeperv3.CallOpts)
}

// LastQuotaRateUpdate is a free data retrieval call binding the contract method 0x112024ff.
//
// Solidity: function lastQuotaRateUpdate() view returns(uint40)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) LastQuotaRateUpdate() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.LastQuotaRateUpdate(&_PoolQuotaKeeperv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Paused() (bool, error) {
	return _PoolQuotaKeeperv3.Contract.Paused(&_PoolQuotaKeeperv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Paused() (bool, error) {
	return _PoolQuotaKeeperv3.Contract.Paused(&_PoolQuotaKeeperv3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Pool() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Pool(&_PoolQuotaKeeperv3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Pool() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Pool(&_PoolQuotaKeeperv3.CallOpts)
}

// PoolQuotaRevenue is a free data retrieval call binding the contract method 0xeb9860a3.
//
// Solidity: function poolQuotaRevenue() view returns(uint256 quotaRevenue)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) PoolQuotaRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "poolQuotaRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolQuotaRevenue is a free data retrieval call binding the contract method 0xeb9860a3.
//
// Solidity: function poolQuotaRevenue() view returns(uint256 quotaRevenue)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) PoolQuotaRevenue() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaRevenue(&_PoolQuotaKeeperv3.CallOpts)
}

// PoolQuotaRevenue is a free data retrieval call binding the contract method 0xeb9860a3.
//
// Solidity: function poolQuotaRevenue() view returns(uint256 quotaRevenue)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) PoolQuotaRevenue() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaRevenue(&_PoolQuotaKeeperv3.CallOpts)
}

// QuotedTokens is a free data retrieval call binding the contract method 0x58279237.
//
// Solidity: function quotedTokens() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) QuotedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "quotedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// QuotedTokens is a free data retrieval call binding the contract method 0x58279237.
//
// Solidity: function quotedTokens() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) QuotedTokens() ([]common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.QuotedTokens(&_PoolQuotaKeeperv3.CallOpts)
}

// QuotedTokens is a free data retrieval call binding the contract method 0x58279237.
//
// Solidity: function quotedTokens() view returns(address[])
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) QuotedTokens() ([]common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.QuotedTokens(&_PoolQuotaKeeperv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Underlying() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Underlying(&_PoolQuotaKeeperv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Underlying() (common.Address, error) {
	return _PoolQuotaKeeperv3.Contract.Underlying(&_PoolQuotaKeeperv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolQuotaKeeperv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Version() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.Version(&_PoolQuotaKeeperv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerSession) Version() (*big.Int, error) {
	return _PoolQuotaKeeperv3.Contract.Version(&_PoolQuotaKeeperv3.CallOpts)
}

// AccrueQuotaInterest is a paid mutator transaction binding the contract method 0x0db1b8ca.
//
// Solidity: function accrueQuotaInterest(address creditAccount, address[] tokens) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) AccrueQuotaInterest(opts *bind.TransactOpts, creditAccount common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "accrueQuotaInterest", creditAccount, tokens)
}

// AccrueQuotaInterest is a paid mutator transaction binding the contract method 0x0db1b8ca.
//
// Solidity: function accrueQuotaInterest(address creditAccount, address[] tokens) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) AccrueQuotaInterest(creditAccount common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AccrueQuotaInterest(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, tokens)
}

// AccrueQuotaInterest is a paid mutator transaction binding the contract method 0x0db1b8ca.
//
// Solidity: function accrueQuotaInterest(address creditAccount, address[] tokens) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) AccrueQuotaInterest(creditAccount common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AccrueQuotaInterest(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, tokens)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address _creditManager) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) AddCreditManager(opts *bind.TransactOpts, _creditManager common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "addCreditManager", _creditManager)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address _creditManager) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) AddCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AddCreditManager(&_PoolQuotaKeeperv3.TransactOpts, _creditManager)
}

// AddCreditManager is a paid mutator transaction binding the contract method 0xe26b2f63.
//
// Solidity: function addCreditManager(address _creditManager) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) AddCreditManager(_creditManager common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AddCreditManager(&_PoolQuotaKeeperv3.TransactOpts, _creditManager)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x364395ee.
//
// Solidity: function addQuotaToken(address token) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) AddQuotaToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "addQuotaToken", token)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x364395ee.
//
// Solidity: function addQuotaToken(address token) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) AddQuotaToken(token common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AddQuotaToken(&_PoolQuotaKeeperv3.TransactOpts, token)
}

// AddQuotaToken is a paid mutator transaction binding the contract method 0x364395ee.
//
// Solidity: function addQuotaToken(address token) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) AddQuotaToken(token common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.AddQuotaToken(&_PoolQuotaKeeperv3.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Pause() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.Pause(&_PoolQuotaKeeperv3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) Pause() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.Pause(&_PoolQuotaKeeperv3.TransactOpts)
}

// RemoveQuotas is a paid mutator transaction binding the contract method 0xfcde5ddc.
//
// Solidity: function removeQuotas(address creditAccount, address[] tokens, bool setLimitsToZero) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) RemoveQuotas(opts *bind.TransactOpts, creditAccount common.Address, tokens []common.Address, setLimitsToZero bool) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "removeQuotas", creditAccount, tokens, setLimitsToZero)
}

// RemoveQuotas is a paid mutator transaction binding the contract method 0xfcde5ddc.
//
// Solidity: function removeQuotas(address creditAccount, address[] tokens, bool setLimitsToZero) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) RemoveQuotas(creditAccount common.Address, tokens []common.Address, setLimitsToZero bool) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.RemoveQuotas(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, tokens, setLimitsToZero)
}

// RemoveQuotas is a paid mutator transaction binding the contract method 0xfcde5ddc.
//
// Solidity: function removeQuotas(address creditAccount, address[] tokens, bool setLimitsToZero) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) RemoveQuotas(creditAccount common.Address, tokens []common.Address, setLimitsToZero bool) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.RemoveQuotas(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, tokens, setLimitsToZero)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetController(&_PoolQuotaKeeperv3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetController(&_PoolQuotaKeeperv3.TransactOpts, newController)
}

// SetGauge is a paid mutator transaction binding the contract method 0x55a68ed3.
//
// Solidity: function setGauge(address _gauge) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) SetGauge(opts *bind.TransactOpts, _gauge common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "setGauge", _gauge)
}

// SetGauge is a paid mutator transaction binding the contract method 0x55a68ed3.
//
// Solidity: function setGauge(address _gauge) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) SetGauge(_gauge common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetGauge(&_PoolQuotaKeeperv3.TransactOpts, _gauge)
}

// SetGauge is a paid mutator transaction binding the contract method 0x55a68ed3.
//
// Solidity: function setGauge(address _gauge) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) SetGauge(_gauge common.Address) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetGauge(&_PoolQuotaKeeperv3.TransactOpts, _gauge)
}

// SetTokenLimit is a paid mutator transaction binding the contract method 0xb26453f6.
//
// Solidity: function setTokenLimit(address token, uint96 limit) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) SetTokenLimit(opts *bind.TransactOpts, token common.Address, limit *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "setTokenLimit", token, limit)
}

// SetTokenLimit is a paid mutator transaction binding the contract method 0xb26453f6.
//
// Solidity: function setTokenLimit(address token, uint96 limit) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) SetTokenLimit(token common.Address, limit *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetTokenLimit(&_PoolQuotaKeeperv3.TransactOpts, token, limit)
}

// SetTokenLimit is a paid mutator transaction binding the contract method 0xb26453f6.
//
// Solidity: function setTokenLimit(address token, uint96 limit) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) SetTokenLimit(token common.Address, limit *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetTokenLimit(&_PoolQuotaKeeperv3.TransactOpts, token, limit)
}

// SetTokenQuotaIncreaseFee is a paid mutator transaction binding the contract method 0x099b9bd7.
//
// Solidity: function setTokenQuotaIncreaseFee(address token, uint16 fee) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) SetTokenQuotaIncreaseFee(opts *bind.TransactOpts, token common.Address, fee uint16) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "setTokenQuotaIncreaseFee", token, fee)
}

// SetTokenQuotaIncreaseFee is a paid mutator transaction binding the contract method 0x099b9bd7.
//
// Solidity: function setTokenQuotaIncreaseFee(address token, uint16 fee) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) SetTokenQuotaIncreaseFee(token common.Address, fee uint16) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetTokenQuotaIncreaseFee(&_PoolQuotaKeeperv3.TransactOpts, token, fee)
}

// SetTokenQuotaIncreaseFee is a paid mutator transaction binding the contract method 0x099b9bd7.
//
// Solidity: function setTokenQuotaIncreaseFee(address token, uint16 fee) returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) SetTokenQuotaIncreaseFee(token common.Address, fee uint16) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.SetTokenQuotaIncreaseFee(&_PoolQuotaKeeperv3.TransactOpts, token, fee)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) Unpause() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.Unpause(&_PoolQuotaKeeperv3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) Unpause() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.Unpause(&_PoolQuotaKeeperv3.TransactOpts)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 requestedChange, uint96 minQuota, uint96 maxQuota) returns(uint128 caQuotaInterestChange, uint128 fees, bool enableToken, bool disableToken)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) UpdateQuota(opts *bind.TransactOpts, creditAccount common.Address, token common.Address, requestedChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "updateQuota", creditAccount, token, requestedChange, minQuota, maxQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 requestedChange, uint96 minQuota, uint96 maxQuota) returns(uint128 caQuotaInterestChange, uint128 fees, bool enableToken, bool disableToken)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) UpdateQuota(creditAccount common.Address, token common.Address, requestedChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.UpdateQuota(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, token, requestedChange, minQuota, maxQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 requestedChange, uint96 minQuota, uint96 maxQuota) returns(uint128 caQuotaInterestChange, uint128 fees, bool enableToken, bool disableToken)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) UpdateQuota(creditAccount common.Address, token common.Address, requestedChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.UpdateQuota(&_PoolQuotaKeeperv3.TransactOpts, creditAccount, token, requestedChange, minQuota, maxQuota)
}

// UpdateRates is a paid mutator transaction binding the contract method 0x3c3821f4.
//
// Solidity: function updateRates() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Transactor) UpdateRates(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.contract.Transact(opts, "updateRates")
}

// UpdateRates is a paid mutator transaction binding the contract method 0x3c3821f4.
//
// Solidity: function updateRates() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Session) UpdateRates() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.UpdateRates(&_PoolQuotaKeeperv3.TransactOpts)
}

// UpdateRates is a paid mutator transaction binding the contract method 0x3c3821f4.
//
// Solidity: function updateRates() returns()
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorSession) UpdateRates() (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.UpdateRates(&_PoolQuotaKeeperv3.TransactOpts)
}

// PoolQuotaKeeperv3AddCreditManagerIterator is returned from FilterAddCreditManager and is used to iterate over the raw logs and unpacked data for AddCreditManager events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3AddCreditManagerIterator struct {
	Event *PoolQuotaKeeperv3AddCreditManager // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3AddCreditManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3AddCreditManager)
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
		it.Event = new(PoolQuotaKeeperv3AddCreditManager)
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
func (it *PoolQuotaKeeperv3AddCreditManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3AddCreditManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3AddCreditManager represents a AddCreditManager event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3AddCreditManager struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddCreditManager is a free log retrieval operation binding the contract event 0xbca7ba46bb626fab79d5a673d0d8293df21968a25350c4d71433f98600618f5f.
//
// Solidity: event AddCreditManager(address indexed creditManager)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterAddCreditManager(opts *bind.FilterOpts, creditManager []common.Address) (*PoolQuotaKeeperv3AddCreditManagerIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "AddCreditManager", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3AddCreditManagerIterator{contract: _PoolQuotaKeeperv3.contract, event: "AddCreditManager", logs: logs, sub: sub}, nil
}

// WatchAddCreditManager is a free log subscription operation binding the contract event 0xbca7ba46bb626fab79d5a673d0d8293df21968a25350c4d71433f98600618f5f.
//
// Solidity: event AddCreditManager(address indexed creditManager)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchAddCreditManager(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3AddCreditManager, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "AddCreditManager", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3AddCreditManager)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "AddCreditManager", log); err != nil {
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

// ParseAddCreditManager is a log parse operation binding the contract event 0xbca7ba46bb626fab79d5a673d0d8293df21968a25350c4d71433f98600618f5f.
//
// Solidity: event AddCreditManager(address indexed creditManager)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseAddCreditManager(log types.Log) (*PoolQuotaKeeperv3AddCreditManager, error) {
	event := new(PoolQuotaKeeperv3AddCreditManager)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "AddCreditManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3AddQuotaTokenIterator is returned from FilterAddQuotaToken and is used to iterate over the raw logs and unpacked data for AddQuotaToken events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3AddQuotaTokenIterator struct {
	Event *PoolQuotaKeeperv3AddQuotaToken // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3AddQuotaTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3AddQuotaToken)
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
		it.Event = new(PoolQuotaKeeperv3AddQuotaToken)
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
func (it *PoolQuotaKeeperv3AddQuotaTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3AddQuotaTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3AddQuotaToken represents a AddQuotaToken event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3AddQuotaToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddQuotaToken is a free log retrieval operation binding the contract event 0x7401ff10219be3dd6d26cc491114a8ae5a0e13ac3af651aae1286afad365947d.
//
// Solidity: event AddQuotaToken(address indexed token)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterAddQuotaToken(opts *bind.FilterOpts, token []common.Address) (*PoolQuotaKeeperv3AddQuotaTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "AddQuotaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3AddQuotaTokenIterator{contract: _PoolQuotaKeeperv3.contract, event: "AddQuotaToken", logs: logs, sub: sub}, nil
}

// WatchAddQuotaToken is a free log subscription operation binding the contract event 0x7401ff10219be3dd6d26cc491114a8ae5a0e13ac3af651aae1286afad365947d.
//
// Solidity: event AddQuotaToken(address indexed token)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchAddQuotaToken(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3AddQuotaToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "AddQuotaToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3AddQuotaToken)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "AddQuotaToken", log); err != nil {
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

// ParseAddQuotaToken is a log parse operation binding the contract event 0x7401ff10219be3dd6d26cc491114a8ae5a0e13ac3af651aae1286afad365947d.
//
// Solidity: event AddQuotaToken(address indexed token)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseAddQuotaToken(log types.Log) (*PoolQuotaKeeperv3AddQuotaToken, error) {
	event := new(PoolQuotaKeeperv3AddQuotaToken)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "AddQuotaToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3NewControllerIterator struct {
	Event *PoolQuotaKeeperv3NewController // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3NewController)
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
		it.Event = new(PoolQuotaKeeperv3NewController)
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
func (it *PoolQuotaKeeperv3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3NewController represents a NewController event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*PoolQuotaKeeperv3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3NewControllerIterator{contract: _PoolQuotaKeeperv3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3NewController)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "NewController", log); err != nil {
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
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseNewController(log types.Log) (*PoolQuotaKeeperv3NewController, error) {
	event := new(PoolQuotaKeeperv3NewController)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3PausedIterator struct {
	Event *PoolQuotaKeeperv3Paused // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3Paused)
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
		it.Event = new(PoolQuotaKeeperv3Paused)
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
func (it *PoolQuotaKeeperv3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3Paused represents a Paused event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterPaused(opts *bind.FilterOpts) (*PoolQuotaKeeperv3PausedIterator, error) {

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3PausedIterator{contract: _PoolQuotaKeeperv3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3Paused) (event.Subscription, error) {

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3Paused)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParsePaused(log types.Log) (*PoolQuotaKeeperv3Paused, error) {
	event := new(PoolQuotaKeeperv3Paused)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3SetGaugeIterator is returned from FilterSetGauge and is used to iterate over the raw logs and unpacked data for SetGauge events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetGaugeIterator struct {
	Event *PoolQuotaKeeperv3SetGauge // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3SetGaugeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3SetGauge)
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
		it.Event = new(PoolQuotaKeeperv3SetGauge)
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
func (it *PoolQuotaKeeperv3SetGaugeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3SetGaugeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3SetGauge represents a SetGauge event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetGauge struct {
	NewGauge common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGauge is a free log retrieval operation binding the contract event 0x17228b08e4c958112a0827a6d8dc8475dba58dd068a3d400800a606794db02a6.
//
// Solidity: event SetGauge(address indexed newGauge)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterSetGauge(opts *bind.FilterOpts, newGauge []common.Address) (*PoolQuotaKeeperv3SetGaugeIterator, error) {

	var newGaugeRule []interface{}
	for _, newGaugeItem := range newGauge {
		newGaugeRule = append(newGaugeRule, newGaugeItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "SetGauge", newGaugeRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3SetGaugeIterator{contract: _PoolQuotaKeeperv3.contract, event: "SetGauge", logs: logs, sub: sub}, nil
}

// WatchSetGauge is a free log subscription operation binding the contract event 0x17228b08e4c958112a0827a6d8dc8475dba58dd068a3d400800a606794db02a6.
//
// Solidity: event SetGauge(address indexed newGauge)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchSetGauge(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3SetGauge, newGauge []common.Address) (event.Subscription, error) {

	var newGaugeRule []interface{}
	for _, newGaugeItem := range newGauge {
		newGaugeRule = append(newGaugeRule, newGaugeItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "SetGauge", newGaugeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3SetGauge)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetGauge", log); err != nil {
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

// ParseSetGauge is a log parse operation binding the contract event 0x17228b08e4c958112a0827a6d8dc8475dba58dd068a3d400800a606794db02a6.
//
// Solidity: event SetGauge(address indexed newGauge)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseSetGauge(log types.Log) (*PoolQuotaKeeperv3SetGauge, error) {
	event := new(PoolQuotaKeeperv3SetGauge)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetGauge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator is returned from FilterSetQuotaIncreaseFee and is used to iterate over the raw logs and unpacked data for SetQuotaIncreaseFee events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator struct {
	Event *PoolQuotaKeeperv3SetQuotaIncreaseFee // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3SetQuotaIncreaseFee)
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
		it.Event = new(PoolQuotaKeeperv3SetQuotaIncreaseFee)
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
func (it *PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3SetQuotaIncreaseFee represents a SetQuotaIncreaseFee event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetQuotaIncreaseFee struct {
	Token common.Address
	Fee   uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetQuotaIncreaseFee is a free log retrieval operation binding the contract event 0x1f985277936e1ecc9dd715575b48f1c6f18902eeb1a1b3a32779122296e64a66.
//
// Solidity: event SetQuotaIncreaseFee(address indexed token, uint16 fee)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterSetQuotaIncreaseFee(opts *bind.FilterOpts, token []common.Address) (*PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "SetQuotaIncreaseFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3SetQuotaIncreaseFeeIterator{contract: _PoolQuotaKeeperv3.contract, event: "SetQuotaIncreaseFee", logs: logs, sub: sub}, nil
}

// WatchSetQuotaIncreaseFee is a free log subscription operation binding the contract event 0x1f985277936e1ecc9dd715575b48f1c6f18902eeb1a1b3a32779122296e64a66.
//
// Solidity: event SetQuotaIncreaseFee(address indexed token, uint16 fee)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchSetQuotaIncreaseFee(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3SetQuotaIncreaseFee, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "SetQuotaIncreaseFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3SetQuotaIncreaseFee)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetQuotaIncreaseFee", log); err != nil {
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

// ParseSetQuotaIncreaseFee is a log parse operation binding the contract event 0x1f985277936e1ecc9dd715575b48f1c6f18902eeb1a1b3a32779122296e64a66.
//
// Solidity: event SetQuotaIncreaseFee(address indexed token, uint16 fee)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseSetQuotaIncreaseFee(log types.Log) (*PoolQuotaKeeperv3SetQuotaIncreaseFee, error) {
	event := new(PoolQuotaKeeperv3SetQuotaIncreaseFee)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetQuotaIncreaseFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3SetTokenLimitIterator is returned from FilterSetTokenLimit and is used to iterate over the raw logs and unpacked data for SetTokenLimit events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetTokenLimitIterator struct {
	Event *PoolQuotaKeeperv3SetTokenLimit // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3SetTokenLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3SetTokenLimit)
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
		it.Event = new(PoolQuotaKeeperv3SetTokenLimit)
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
func (it *PoolQuotaKeeperv3SetTokenLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3SetTokenLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3SetTokenLimit represents a SetTokenLimit event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3SetTokenLimit struct {
	Token common.Address
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetTokenLimit is a free log retrieval operation binding the contract event 0x86089ad7ab4cb6d03a20ccb3176599b628f4a4b80ceacf88369108bf10ffa1c9.
//
// Solidity: event SetTokenLimit(address indexed token, uint96 limit)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterSetTokenLimit(opts *bind.FilterOpts, token []common.Address) (*PoolQuotaKeeperv3SetTokenLimitIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "SetTokenLimit", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3SetTokenLimitIterator{contract: _PoolQuotaKeeperv3.contract, event: "SetTokenLimit", logs: logs, sub: sub}, nil
}

// WatchSetTokenLimit is a free log subscription operation binding the contract event 0x86089ad7ab4cb6d03a20ccb3176599b628f4a4b80ceacf88369108bf10ffa1c9.
//
// Solidity: event SetTokenLimit(address indexed token, uint96 limit)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchSetTokenLimit(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3SetTokenLimit, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "SetTokenLimit", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3SetTokenLimit)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetTokenLimit", log); err != nil {
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

// ParseSetTokenLimit is a log parse operation binding the contract event 0x86089ad7ab4cb6d03a20ccb3176599b628f4a4b80ceacf88369108bf10ffa1c9.
//
// Solidity: event SetTokenLimit(address indexed token, uint96 limit)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseSetTokenLimit(log types.Log) (*PoolQuotaKeeperv3SetTokenLimit, error) {
	event := new(PoolQuotaKeeperv3SetTokenLimit)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "SetTokenLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UnpausedIterator struct {
	Event *PoolQuotaKeeperv3Unpaused // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3Unpaused)
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
		it.Event = new(PoolQuotaKeeperv3Unpaused)
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
func (it *PoolQuotaKeeperv3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3Unpaused represents a Unpaused event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolQuotaKeeperv3UnpausedIterator, error) {

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3UnpausedIterator{contract: _PoolQuotaKeeperv3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3Unpaused) (event.Subscription, error) {

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3Unpaused)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseUnpaused(log types.Log) (*PoolQuotaKeeperv3Unpaused, error) {
	event := new(PoolQuotaKeeperv3Unpaused)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3UpdateQuotaIterator is returned from FilterUpdateQuota and is used to iterate over the raw logs and unpacked data for UpdateQuota events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateQuotaIterator struct {
	Event *PoolQuotaKeeperv3UpdateQuota // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3UpdateQuota)
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
		it.Event = new(PoolQuotaKeeperv3UpdateQuota)
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
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3UpdateQuota represents a UpdateQuota event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateQuota struct {
	CreditAccount common.Address
	Token         common.Address
	QuotaChange   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuota is a free log retrieval operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterUpdateQuota(opts *bind.FilterOpts, creditAccount []common.Address, token []common.Address) (*PoolQuotaKeeperv3UpdateQuotaIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "UpdateQuota", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3UpdateQuotaIterator{contract: _PoolQuotaKeeperv3.contract, event: "UpdateQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateQuota is a free log subscription operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchUpdateQuota(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3UpdateQuota, creditAccount []common.Address, token []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "UpdateQuota", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3UpdateQuota)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateQuota", log); err != nil {
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

// ParseUpdateQuota is a log parse operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseUpdateQuota(log types.Log) (*PoolQuotaKeeperv3UpdateQuota, error) {
	event := new(PoolQuotaKeeperv3UpdateQuota)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateQuota", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolQuotaKeeperv3UpdateTokenQuotaRateIterator is returned from FilterUpdateTokenQuotaRate and is used to iterate over the raw logs and unpacked data for UpdateTokenQuotaRate events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateTokenQuotaRateIterator struct {
	Event *PoolQuotaKeeperv3UpdateTokenQuotaRate // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3UpdateTokenQuotaRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3UpdateTokenQuotaRate)
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
		it.Event = new(PoolQuotaKeeperv3UpdateTokenQuotaRate)
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
func (it *PoolQuotaKeeperv3UpdateTokenQuotaRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3UpdateTokenQuotaRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3UpdateTokenQuotaRate represents a UpdateTokenQuotaRate event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateTokenQuotaRate struct {
	Token common.Address
	Rate  uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenQuotaRate is a free log retrieval operation binding the contract event 0xfb19913ea8fcd2e3d22d200707473d031876b05d1ecb42173e73292ed910ac85.
//
// Solidity: event UpdateTokenQuotaRate(address indexed token, uint16 rate)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterUpdateTokenQuotaRate(opts *bind.FilterOpts, token []common.Address) (*PoolQuotaKeeperv3UpdateTokenQuotaRateIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "UpdateTokenQuotaRate", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3UpdateTokenQuotaRateIterator{contract: _PoolQuotaKeeperv3.contract, event: "UpdateTokenQuotaRate", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenQuotaRate is a free log subscription operation binding the contract event 0xfb19913ea8fcd2e3d22d200707473d031876b05d1ecb42173e73292ed910ac85.
//
// Solidity: event UpdateTokenQuotaRate(address indexed token, uint16 rate)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchUpdateTokenQuotaRate(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3UpdateTokenQuotaRate, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "UpdateTokenQuotaRate", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3UpdateTokenQuotaRate)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateTokenQuotaRate", log); err != nil {
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

// ParseUpdateTokenQuotaRate is a log parse operation binding the contract event 0xfb19913ea8fcd2e3d22d200707473d031876b05d1ecb42173e73292ed910ac85.
//
// Solidity: event UpdateTokenQuotaRate(address indexed token, uint16 rate)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseUpdateTokenQuotaRate(log types.Log) (*PoolQuotaKeeperv3UpdateTokenQuotaRate, error) {
	event := new(PoolQuotaKeeperv3UpdateTokenQuotaRate)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateTokenQuotaRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
