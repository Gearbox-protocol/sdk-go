// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditConfiguratorv310

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

// CreditConfiguratorv310MetaData contains all meta data concerning the CreditConfiguratorv310 contract.
var CreditConfiguratorv310MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCollateralToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addEmergencyLiquidator\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowAdapter\",\"inputs\":[{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowedAdapters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"controller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditFacade\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forbidAdapter\",\"inputs\":[{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forbidBorrowing\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forbidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rampLiquidationThreshold\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThresholdFinal\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rampStart\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"rampDuration\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeEmergencyLiquidator\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCreditFacade\",\"inputs\":[{\"name\":\"newCreditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"migrateParams\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDebtLimits\",\"inputs\":[{\"name\":\"newMinDebt\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"newMaxDebt\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExpirationDate\",\"inputs\":[{\"name\":\"newExpirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationPremium\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationPremiumExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidationThreshold\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLossLiquidator\",\"inputs\":[{\"name\":\"newLossLiquidator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxDebtPerBlockMultiplier\",\"inputs\":[{\"name\":\"newMaxDebtLimitPerBlockMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceOracle\",\"inputs\":[{\"name\":\"newPriceOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeCreditConfigurator\",\"inputs\":[{\"name\":\"newCreditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddCollateralToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddEmergencyLiquidator\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowAdapter\",\"inputs\":[{\"name\":\"targetContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreditConfiguratorUpgraded\",\"inputs\":[{\"name\":\"creditConfigurator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForbidAdapter\",\"inputs\":[{\"name\":\"targetContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForbidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewController\",\"inputs\":[{\"name\":\"newController\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveEmergencyLiquidator\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ScheduleTokenLiquidationThresholdRamp\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"liquidationThresholdInitial\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationThresholdFinal\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"timestampRampStart\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"timestampRampEnd\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetBorrowingLimits\",\"inputs\":[{\"name\":\"minDebt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetCreditFacade\",\"inputs\":[{\"name\":\"creditFacade\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetExpirationDate\",\"inputs\":[{\"name\":\"expirationDate\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetLossLiquidator\",\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMaxDebtPerBlockMultiplier\",\"inputs\":[{\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPriceOracle\",\"inputs\":[{\"name\":\"priceOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTokenLiquidationThreshold\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateFees\",\"inputs\":[{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationPremium\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"liquidationPremiumExpired\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdapterIsNotRegisteredException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressIsNotContractException\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerNotConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotControllerOrConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotPausableAdminException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncompatibleContractException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectAdaptersSetException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectExpirationDateException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectLimitsException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectLiquidationThresholdException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectParameterException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectPriceException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectTokenContractException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceFeedDoesNotExistException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TargetContractNotAllowedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenIsNotQuotedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotAllowedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// CreditConfiguratorv310ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditConfiguratorv310MetaData.ABI instead.
var CreditConfiguratorv310ABI = CreditConfiguratorv310MetaData.ABI

// CreditConfiguratorv310 is an auto generated Go binding around an Ethereum contract.
type CreditConfiguratorv310 struct {
	CreditConfiguratorv310Caller     // Read-only binding to the contract
	CreditConfiguratorv310Transactor // Write-only binding to the contract
	CreditConfiguratorv310Filterer   // Log filterer for contract events
}

// CreditConfiguratorv310Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditConfiguratorv310Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv310Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditConfiguratorv310Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv310Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditConfiguratorv310Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv310Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditConfiguratorv310Session struct {
	Contract     *CreditConfiguratorv310 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CreditConfiguratorv310CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditConfiguratorv310CallerSession struct {
	Contract *CreditConfiguratorv310Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CreditConfiguratorv310TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditConfiguratorv310TransactorSession struct {
	Contract     *CreditConfiguratorv310Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CreditConfiguratorv310Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditConfiguratorv310Raw struct {
	Contract *CreditConfiguratorv310 // Generic contract binding to access the raw methods on
}

// CreditConfiguratorv310CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditConfiguratorv310CallerRaw struct {
	Contract *CreditConfiguratorv310Caller // Generic read-only contract binding to access the raw methods on
}

// CreditConfiguratorv310TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditConfiguratorv310TransactorRaw struct {
	Contract *CreditConfiguratorv310Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditConfiguratorv310 creates a new instance of CreditConfiguratorv310, bound to a specific deployed contract.
func NewCreditConfiguratorv310(address common.Address, backend bind.ContractBackend) (*CreditConfiguratorv310, error) {
	contract, err := bindCreditConfiguratorv310(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310{CreditConfiguratorv310Caller: CreditConfiguratorv310Caller{contract: contract}, CreditConfiguratorv310Transactor: CreditConfiguratorv310Transactor{contract: contract}, CreditConfiguratorv310Filterer: CreditConfiguratorv310Filterer{contract: contract}}, nil
}

// NewCreditConfiguratorv310Caller creates a new read-only instance of CreditConfiguratorv310, bound to a specific deployed contract.
func NewCreditConfiguratorv310Caller(address common.Address, caller bind.ContractCaller) (*CreditConfiguratorv310Caller, error) {
	contract, err := bindCreditConfiguratorv310(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310Caller{contract: contract}, nil
}

// NewCreditConfiguratorv310Transactor creates a new write-only instance of CreditConfiguratorv310, bound to a specific deployed contract.
func NewCreditConfiguratorv310Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditConfiguratorv310Transactor, error) {
	contract, err := bindCreditConfiguratorv310(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310Transactor{contract: contract}, nil
}

// NewCreditConfiguratorv310Filterer creates a new log filterer instance of CreditConfiguratorv310, bound to a specific deployed contract.
func NewCreditConfiguratorv310Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditConfiguratorv310Filterer, error) {
	contract, err := bindCreditConfiguratorv310(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310Filterer{contract: contract}, nil
}

// bindCreditConfiguratorv310 binds a generic wrapper to an already deployed contract.
func bindCreditConfiguratorv310(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditConfiguratorv310MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfiguratorv310 *CreditConfiguratorv310Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfiguratorv310.Contract.CreditConfiguratorv310Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfiguratorv310 *CreditConfiguratorv310Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.CreditConfiguratorv310Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfiguratorv310 *CreditConfiguratorv310Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.CreditConfiguratorv310Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfiguratorv310.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) Acl() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Acl(&_CreditConfiguratorv310.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) Acl() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Acl(&_CreditConfiguratorv310.CallOpts)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) AllowedAdapters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "allowedAdapters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) AllowedAdapters() ([]common.Address, error) {
	return _CreditConfiguratorv310.Contract.AllowedAdapters(&_CreditConfiguratorv310.CallOpts)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) AllowedAdapters() ([]common.Address, error) {
	return _CreditConfiguratorv310.Contract.AllowedAdapters(&_CreditConfiguratorv310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) ContractType() ([32]byte, error) {
	return _CreditConfiguratorv310.Contract.ContractType(&_CreditConfiguratorv310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) ContractType() ([32]byte, error) {
	return _CreditConfiguratorv310.Contract.ContractType(&_CreditConfiguratorv310.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) Controller() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Controller(&_CreditConfiguratorv310.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) Controller() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Controller(&_CreditConfiguratorv310.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) CreditFacade() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.CreditFacade(&_CreditConfiguratorv310.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) CreditFacade() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.CreditFacade(&_CreditConfiguratorv310.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) CreditManager() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.CreditManager(&_CreditConfiguratorv310.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) CreditManager() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.CreditManager(&_CreditConfiguratorv310.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) Underlying() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Underlying(&_CreditConfiguratorv310.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) Underlying() (common.Address, error) {
	return _CreditConfiguratorv310.Contract.Underlying(&_CreditConfiguratorv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfiguratorv310.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) Version() (*big.Int, error) {
	return _CreditConfiguratorv310.Contract.Version(&_CreditConfiguratorv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv310 *CreditConfiguratorv310CallerSession) Version() (*big.Int, error) {
	return _CreditConfiguratorv310.Contract.Version(&_CreditConfiguratorv310.CallOpts)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) AddCollateralToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "addCollateralToken", token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AddCollateralToken(&_CreditConfiguratorv310.TransactOpts, token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AddCollateralToken(&_CreditConfiguratorv310.TransactOpts, token, liquidationThreshold)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) AddEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "addEmergencyLiquidator", liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AddEmergencyLiquidator(&_CreditConfiguratorv310.TransactOpts, liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AddEmergencyLiquidator(&_CreditConfiguratorv310.TransactOpts, liquidator)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) AllowAdapter(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "allowAdapter", adapter)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) AllowAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AllowAdapter(&_CreditConfiguratorv310.TransactOpts, adapter)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) AllowAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AllowAdapter(&_CreditConfiguratorv310.TransactOpts, adapter)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) AllowToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "allowToken", token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AllowToken(&_CreditConfiguratorv310.TransactOpts, token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.AllowToken(&_CreditConfiguratorv310.TransactOpts, token)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) ForbidAdapter(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "forbidAdapter", adapter)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) ForbidAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidAdapter(&_CreditConfiguratorv310.TransactOpts, adapter)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) ForbidAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidAdapter(&_CreditConfiguratorv310.TransactOpts, adapter)
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) ForbidBorrowing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "forbidBorrowing")
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) ForbidBorrowing() (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidBorrowing(&_CreditConfiguratorv310.TransactOpts)
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) ForbidBorrowing() (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidBorrowing(&_CreditConfiguratorv310.TransactOpts)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidToken(&_CreditConfiguratorv310.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.ForbidToken(&_CreditConfiguratorv310.TransactOpts, token)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) RampLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "rampLiquidationThreshold", token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) RampLiquidationThreshold(token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.RampLiquidationThreshold(&_CreditConfiguratorv310.TransactOpts, token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) RampLiquidationThreshold(token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.RampLiquidationThreshold(&_CreditConfiguratorv310.TransactOpts, token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) RemoveEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "removeEmergencyLiquidator", liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.RemoveEmergencyLiquidator(&_CreditConfiguratorv310.TransactOpts, liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.RemoveEmergencyLiquidator(&_CreditConfiguratorv310.TransactOpts, liquidator)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetController(&_CreditConfiguratorv310.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetController(&_CreditConfiguratorv310.TransactOpts, newController)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetCreditFacade(opts *bind.TransactOpts, newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setCreditFacade", newCreditFacade, migrateParams)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetCreditFacade(newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetCreditFacade(&_CreditConfiguratorv310.TransactOpts, newCreditFacade, migrateParams)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetCreditFacade(newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetCreditFacade(&_CreditConfiguratorv310.TransactOpts, newCreditFacade, migrateParams)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0xef697683.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetDebtLimits(opts *bind.TransactOpts, newMinDebt *big.Int, newMaxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setDebtLimits", newMinDebt, newMaxDebt)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0xef697683.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetDebtLimits(newMinDebt *big.Int, newMaxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetDebtLimits(&_CreditConfiguratorv310.TransactOpts, newMinDebt, newMaxDebt)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0xef697683.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetDebtLimits(newMinDebt *big.Int, newMaxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetDebtLimits(&_CreditConfiguratorv310.TransactOpts, newMinDebt, newMaxDebt)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetExpirationDate(opts *bind.TransactOpts, newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setExpirationDate", newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetExpirationDate(&_CreditConfiguratorv310.TransactOpts, newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetExpirationDate(&_CreditConfiguratorv310.TransactOpts, newExpirationDate)
}

// SetFees is a paid mutator transaction binding the contract method 0x42f7723f.
//
// Solidity: function setFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetFees(opts *bind.TransactOpts, feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setFees", feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0x42f7723f.
//
// Solidity: function setFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetFees(feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetFees(&_CreditConfiguratorv310.TransactOpts, feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0x42f7723f.
//
// Solidity: function setFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetFees(feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetFees(&_CreditConfiguratorv310.TransactOpts, feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetLiquidationThreshold(&_CreditConfiguratorv310.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetLiquidationThreshold(&_CreditConfiguratorv310.TransactOpts, token, liquidationThreshold)
}

// SetLossLiquidator is a paid mutator transaction binding the contract method 0xf3003d8f.
//
// Solidity: function setLossLiquidator(address newLossLiquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetLossLiquidator(opts *bind.TransactOpts, newLossLiquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setLossLiquidator", newLossLiquidator)
}

// SetLossLiquidator is a paid mutator transaction binding the contract method 0xf3003d8f.
//
// Solidity: function setLossLiquidator(address newLossLiquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetLossLiquidator(newLossLiquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetLossLiquidator(&_CreditConfiguratorv310.TransactOpts, newLossLiquidator)
}

// SetLossLiquidator is a paid mutator transaction binding the contract method 0xf3003d8f.
//
// Solidity: function setLossLiquidator(address newLossLiquidator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetLossLiquidator(newLossLiquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetLossLiquidator(&_CreditConfiguratorv310.TransactOpts, newLossLiquidator)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetMaxDebtPerBlockMultiplier(opts *bind.TransactOpts, newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setMaxDebtPerBlockMultiplier", newMaxDebtLimitPerBlockMultiplier)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetMaxDebtPerBlockMultiplier(newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetMaxDebtPerBlockMultiplier(&_CreditConfiguratorv310.TransactOpts, newMaxDebtLimitPerBlockMultiplier)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetMaxDebtPerBlockMultiplier(newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetMaxDebtPerBlockMultiplier(&_CreditConfiguratorv310.TransactOpts, newMaxDebtLimitPerBlockMultiplier)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) SetPriceOracle(opts *bind.TransactOpts, newPriceOracle common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "setPriceOracle", newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetPriceOracle(&_CreditConfiguratorv310.TransactOpts, newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.SetPriceOracle(&_CreditConfiguratorv310.TransactOpts, newPriceOracle)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Transactor) UpgradeCreditConfigurator(opts *bind.TransactOpts, newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.contract.Transact(opts, "upgradeCreditConfigurator", newCreditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310Session) UpgradeCreditConfigurator(newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.UpgradeCreditConfigurator(&_CreditConfiguratorv310.TransactOpts, newCreditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv310 *CreditConfiguratorv310TransactorSession) UpgradeCreditConfigurator(newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv310.Contract.UpgradeCreditConfigurator(&_CreditConfiguratorv310.TransactOpts, newCreditConfigurator)
}

// CreditConfiguratorv310AddCollateralTokenIterator is returned from FilterAddCollateralToken and is used to iterate over the raw logs and unpacked data for AddCollateralToken events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AddCollateralTokenIterator struct {
	Event *CreditConfiguratorv310AddCollateralToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310AddCollateralTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310AddCollateralToken)
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
		it.Event = new(CreditConfiguratorv310AddCollateralToken)
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
func (it *CreditConfiguratorv310AddCollateralTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310AddCollateralTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310AddCollateralToken represents a AddCollateralToken event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AddCollateralToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddCollateralToken is a free log retrieval operation binding the contract event 0x7c3f95f8569977586927f95930461a261e2121e326fcb513242f9e5c8b8ea6dc.
//
// Solidity: event AddCollateralToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterAddCollateralToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv310AddCollateralTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "AddCollateralToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310AddCollateralTokenIterator{contract: _CreditConfiguratorv310.contract, event: "AddCollateralToken", logs: logs, sub: sub}, nil
}

// WatchAddCollateralToken is a free log subscription operation binding the contract event 0x7c3f95f8569977586927f95930461a261e2121e326fcb513242f9e5c8b8ea6dc.
//
// Solidity: event AddCollateralToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchAddCollateralToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310AddCollateralToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "AddCollateralToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310AddCollateralToken)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AddCollateralToken", log); err != nil {
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

// ParseAddCollateralToken is a log parse operation binding the contract event 0x7c3f95f8569977586927f95930461a261e2121e326fcb513242f9e5c8b8ea6dc.
//
// Solidity: event AddCollateralToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseAddCollateralToken(log types.Log) (*CreditConfiguratorv310AddCollateralToken, error) {
	event := new(CreditConfiguratorv310AddCollateralToken)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AddCollateralToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310AddEmergencyLiquidatorIterator is returned from FilterAddEmergencyLiquidator and is used to iterate over the raw logs and unpacked data for AddEmergencyLiquidator events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AddEmergencyLiquidatorIterator struct {
	Event *CreditConfiguratorv310AddEmergencyLiquidator // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310AddEmergencyLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310AddEmergencyLiquidator)
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
		it.Event = new(CreditConfiguratorv310AddEmergencyLiquidator)
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
func (it *CreditConfiguratorv310AddEmergencyLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310AddEmergencyLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310AddEmergencyLiquidator represents a AddEmergencyLiquidator event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AddEmergencyLiquidator struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddEmergencyLiquidator is a free log retrieval operation binding the contract event 0x35b5318c4163fcef2999d30de8d1af689327f68fa51a148804fa6ed8f5f40ff4.
//
// Solidity: event AddEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterAddEmergencyLiquidator(opts *bind.FilterOpts, liquidator []common.Address) (*CreditConfiguratorv310AddEmergencyLiquidatorIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "AddEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310AddEmergencyLiquidatorIterator{contract: _CreditConfiguratorv310.contract, event: "AddEmergencyLiquidator", logs: logs, sub: sub}, nil
}

// WatchAddEmergencyLiquidator is a free log subscription operation binding the contract event 0x35b5318c4163fcef2999d30de8d1af689327f68fa51a148804fa6ed8f5f40ff4.
//
// Solidity: event AddEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchAddEmergencyLiquidator(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310AddEmergencyLiquidator, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "AddEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310AddEmergencyLiquidator)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AddEmergencyLiquidator", log); err != nil {
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

// ParseAddEmergencyLiquidator is a log parse operation binding the contract event 0x35b5318c4163fcef2999d30de8d1af689327f68fa51a148804fa6ed8f5f40ff4.
//
// Solidity: event AddEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseAddEmergencyLiquidator(log types.Log) (*CreditConfiguratorv310AddEmergencyLiquidator, error) {
	event := new(CreditConfiguratorv310AddEmergencyLiquidator)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AddEmergencyLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310AllowAdapterIterator is returned from FilterAllowAdapter and is used to iterate over the raw logs and unpacked data for AllowAdapter events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AllowAdapterIterator struct {
	Event *CreditConfiguratorv310AllowAdapter // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310AllowAdapterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310AllowAdapter)
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
		it.Event = new(CreditConfiguratorv310AllowAdapter)
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
func (it *CreditConfiguratorv310AllowAdapterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310AllowAdapterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310AllowAdapter represents a AllowAdapter event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AllowAdapter struct {
	TargetContract common.Address
	Adapter        common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowAdapter is a free log retrieval operation binding the contract event 0x0bc09e53304ef58ff3ff8295411d9171c75ee4af48277db5fc605ab12e056bee.
//
// Solidity: event AllowAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterAllowAdapter(opts *bind.FilterOpts, targetContract []common.Address, adapter []common.Address) (*CreditConfiguratorv310AllowAdapterIterator, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "AllowAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310AllowAdapterIterator{contract: _CreditConfiguratorv310.contract, event: "AllowAdapter", logs: logs, sub: sub}, nil
}

// WatchAllowAdapter is a free log subscription operation binding the contract event 0x0bc09e53304ef58ff3ff8295411d9171c75ee4af48277db5fc605ab12e056bee.
//
// Solidity: event AllowAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchAllowAdapter(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310AllowAdapter, targetContract []common.Address, adapter []common.Address) (event.Subscription, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "AllowAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310AllowAdapter)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AllowAdapter", log); err != nil {
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

// ParseAllowAdapter is a log parse operation binding the contract event 0x0bc09e53304ef58ff3ff8295411d9171c75ee4af48277db5fc605ab12e056bee.
//
// Solidity: event AllowAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseAllowAdapter(log types.Log) (*CreditConfiguratorv310AllowAdapter, error) {
	event := new(CreditConfiguratorv310AllowAdapter)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AllowAdapter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310AllowTokenIterator is returned from FilterAllowToken and is used to iterate over the raw logs and unpacked data for AllowToken events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AllowTokenIterator struct {
	Event *CreditConfiguratorv310AllowToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310AllowTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310AllowToken)
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
		it.Event = new(CreditConfiguratorv310AllowToken)
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
func (it *CreditConfiguratorv310AllowTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310AllowTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310AllowToken represents a AllowToken event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310AllowToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAllowToken is a free log retrieval operation binding the contract event 0x14009112f2dcb15cad32dab6bf972d6d85286e4ae1178f27323ffe25359459e6.
//
// Solidity: event AllowToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterAllowToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv310AllowTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "AllowToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310AllowTokenIterator{contract: _CreditConfiguratorv310.contract, event: "AllowToken", logs: logs, sub: sub}, nil
}

// WatchAllowToken is a free log subscription operation binding the contract event 0x14009112f2dcb15cad32dab6bf972d6d85286e4ae1178f27323ffe25359459e6.
//
// Solidity: event AllowToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchAllowToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310AllowToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "AllowToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310AllowToken)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AllowToken", log); err != nil {
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

// ParseAllowToken is a log parse operation binding the contract event 0x14009112f2dcb15cad32dab6bf972d6d85286e4ae1178f27323ffe25359459e6.
//
// Solidity: event AllowToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseAllowToken(log types.Log) (*CreditConfiguratorv310AllowToken, error) {
	event := new(CreditConfiguratorv310AllowToken)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "AllowToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310CreditConfiguratorUpgradedIterator is returned from FilterCreditConfiguratorUpgraded and is used to iterate over the raw logs and unpacked data for CreditConfiguratorUpgraded events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310CreditConfiguratorUpgradedIterator struct {
	Event *CreditConfiguratorv310CreditConfiguratorUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310CreditConfiguratorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310CreditConfiguratorUpgraded)
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
		it.Event = new(CreditConfiguratorv310CreditConfiguratorUpgraded)
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
func (it *CreditConfiguratorv310CreditConfiguratorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310CreditConfiguratorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310CreditConfiguratorUpgraded represents a CreditConfiguratorUpgraded event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310CreditConfiguratorUpgraded struct {
	CreditConfigurator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreditConfiguratorUpgraded is a free log retrieval operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed creditConfigurator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterCreditConfiguratorUpgraded(opts *bind.FilterOpts, creditConfigurator []common.Address) (*CreditConfiguratorv310CreditConfiguratorUpgradedIterator, error) {

	var creditConfiguratorRule []interface{}
	for _, creditConfiguratorItem := range creditConfigurator {
		creditConfiguratorRule = append(creditConfiguratorRule, creditConfiguratorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "CreditConfiguratorUpgraded", creditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310CreditConfiguratorUpgradedIterator{contract: _CreditConfiguratorv310.contract, event: "CreditConfiguratorUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditConfiguratorUpgraded is a free log subscription operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed creditConfigurator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchCreditConfiguratorUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310CreditConfiguratorUpgraded, creditConfigurator []common.Address) (event.Subscription, error) {

	var creditConfiguratorRule []interface{}
	for _, creditConfiguratorItem := range creditConfigurator {
		creditConfiguratorRule = append(creditConfiguratorRule, creditConfiguratorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "CreditConfiguratorUpgraded", creditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310CreditConfiguratorUpgraded)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
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

// ParseCreditConfiguratorUpgraded is a log parse operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed creditConfigurator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseCreditConfiguratorUpgraded(log types.Log) (*CreditConfiguratorv310CreditConfiguratorUpgraded, error) {
	event := new(CreditConfiguratorv310CreditConfiguratorUpgraded)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310ForbidAdapterIterator is returned from FilterForbidAdapter and is used to iterate over the raw logs and unpacked data for ForbidAdapter events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ForbidAdapterIterator struct {
	Event *CreditConfiguratorv310ForbidAdapter // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310ForbidAdapterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310ForbidAdapter)
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
		it.Event = new(CreditConfiguratorv310ForbidAdapter)
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
func (it *CreditConfiguratorv310ForbidAdapterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310ForbidAdapterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310ForbidAdapter represents a ForbidAdapter event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ForbidAdapter struct {
	TargetContract common.Address
	Adapter        common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterForbidAdapter is a free log retrieval operation binding the contract event 0x3f688c7b4a117ceec70e927a9ed68836d3da0224eee121f856fc87ad5baa2a80.
//
// Solidity: event ForbidAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterForbidAdapter(opts *bind.FilterOpts, targetContract []common.Address, adapter []common.Address) (*CreditConfiguratorv310ForbidAdapterIterator, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "ForbidAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310ForbidAdapterIterator{contract: _CreditConfiguratorv310.contract, event: "ForbidAdapter", logs: logs, sub: sub}, nil
}

// WatchForbidAdapter is a free log subscription operation binding the contract event 0x3f688c7b4a117ceec70e927a9ed68836d3da0224eee121f856fc87ad5baa2a80.
//
// Solidity: event ForbidAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchForbidAdapter(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310ForbidAdapter, targetContract []common.Address, adapter []common.Address) (event.Subscription, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "ForbidAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310ForbidAdapter)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ForbidAdapter", log); err != nil {
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

// ParseForbidAdapter is a log parse operation binding the contract event 0x3f688c7b4a117ceec70e927a9ed68836d3da0224eee121f856fc87ad5baa2a80.
//
// Solidity: event ForbidAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseForbidAdapter(log types.Log) (*CreditConfiguratorv310ForbidAdapter, error) {
	event := new(CreditConfiguratorv310ForbidAdapter)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ForbidAdapter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310ForbidTokenIterator is returned from FilterForbidToken and is used to iterate over the raw logs and unpacked data for ForbidToken events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ForbidTokenIterator struct {
	Event *CreditConfiguratorv310ForbidToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310ForbidTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310ForbidToken)
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
		it.Event = new(CreditConfiguratorv310ForbidToken)
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
func (it *CreditConfiguratorv310ForbidTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310ForbidTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310ForbidToken represents a ForbidToken event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ForbidToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterForbidToken is a free log retrieval operation binding the contract event 0x9d65afef45c30b784a1e4621dbcbb194ebb6aabe16c9a4abce9ab1775a962b76.
//
// Solidity: event ForbidToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterForbidToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv310ForbidTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "ForbidToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310ForbidTokenIterator{contract: _CreditConfiguratorv310.contract, event: "ForbidToken", logs: logs, sub: sub}, nil
}

// WatchForbidToken is a free log subscription operation binding the contract event 0x9d65afef45c30b784a1e4621dbcbb194ebb6aabe16c9a4abce9ab1775a962b76.
//
// Solidity: event ForbidToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchForbidToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310ForbidToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "ForbidToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310ForbidToken)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ForbidToken", log); err != nil {
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

// ParseForbidToken is a log parse operation binding the contract event 0x9d65afef45c30b784a1e4621dbcbb194ebb6aabe16c9a4abce9ab1775a962b76.
//
// Solidity: event ForbidToken(address indexed token)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseForbidToken(log types.Log) (*CreditConfiguratorv310ForbidToken, error) {
	event := new(CreditConfiguratorv310ForbidToken)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ForbidToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310NewControllerIterator struct {
	Event *CreditConfiguratorv310NewController // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310NewController)
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
		it.Event = new(CreditConfiguratorv310NewController)
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
func (it *CreditConfiguratorv310NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310NewController represents a NewController event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*CreditConfiguratorv310NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310NewControllerIterator{contract: _CreditConfiguratorv310.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310NewController)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "NewController", log); err != nil {
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
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseNewController(log types.Log) (*CreditConfiguratorv310NewController, error) {
	event := new(CreditConfiguratorv310NewController)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310RemoveEmergencyLiquidatorIterator is returned from FilterRemoveEmergencyLiquidator and is used to iterate over the raw logs and unpacked data for RemoveEmergencyLiquidator events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310RemoveEmergencyLiquidatorIterator struct {
	Event *CreditConfiguratorv310RemoveEmergencyLiquidator // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310RemoveEmergencyLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310RemoveEmergencyLiquidator)
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
		it.Event = new(CreditConfiguratorv310RemoveEmergencyLiquidator)
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
func (it *CreditConfiguratorv310RemoveEmergencyLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310RemoveEmergencyLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310RemoveEmergencyLiquidator represents a RemoveEmergencyLiquidator event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310RemoveEmergencyLiquidator struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveEmergencyLiquidator is a free log retrieval operation binding the contract event 0xc03fe683aa5f2a3776871ebf04508ced24c0335e0d19abd72b6a0d1950e1e255.
//
// Solidity: event RemoveEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterRemoveEmergencyLiquidator(opts *bind.FilterOpts, liquidator []common.Address) (*CreditConfiguratorv310RemoveEmergencyLiquidatorIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "RemoveEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310RemoveEmergencyLiquidatorIterator{contract: _CreditConfiguratorv310.contract, event: "RemoveEmergencyLiquidator", logs: logs, sub: sub}, nil
}

// WatchRemoveEmergencyLiquidator is a free log subscription operation binding the contract event 0xc03fe683aa5f2a3776871ebf04508ced24c0335e0d19abd72b6a0d1950e1e255.
//
// Solidity: event RemoveEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchRemoveEmergencyLiquidator(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310RemoveEmergencyLiquidator, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "RemoveEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310RemoveEmergencyLiquidator)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "RemoveEmergencyLiquidator", log); err != nil {
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

// ParseRemoveEmergencyLiquidator is a log parse operation binding the contract event 0xc03fe683aa5f2a3776871ebf04508ced24c0335e0d19abd72b6a0d1950e1e255.
//
// Solidity: event RemoveEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseRemoveEmergencyLiquidator(log types.Log) (*CreditConfiguratorv310RemoveEmergencyLiquidator, error) {
	event := new(CreditConfiguratorv310RemoveEmergencyLiquidator)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "RemoveEmergencyLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator is returned from FilterScheduleTokenLiquidationThresholdRamp and is used to iterate over the raw logs and unpacked data for ScheduleTokenLiquidationThresholdRamp events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator struct {
	Event *CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp)
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
		it.Event = new(CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp)
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
func (it *CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp represents a ScheduleTokenLiquidationThresholdRamp event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp struct {
	Token                       common.Address
	LiquidationThresholdInitial uint16
	LiquidationThresholdFinal   uint16
	TimestampRampStart          *big.Int
	TimestampRampEnd            *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterScheduleTokenLiquidationThresholdRamp is a free log retrieval operation binding the contract event 0xa8193c198aab4146e3640f414ba8473918c6d028f45b27fb08b185a16c15ce23.
//
// Solidity: event ScheduleTokenLiquidationThresholdRamp(address indexed token, uint16 liquidationThresholdInitial, uint16 liquidationThresholdFinal, uint40 timestampRampStart, uint40 timestampRampEnd)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterScheduleTokenLiquidationThresholdRamp(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "ScheduleTokenLiquidationThresholdRamp", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310ScheduleTokenLiquidationThresholdRampIterator{contract: _CreditConfiguratorv310.contract, event: "ScheduleTokenLiquidationThresholdRamp", logs: logs, sub: sub}, nil
}

// WatchScheduleTokenLiquidationThresholdRamp is a free log subscription operation binding the contract event 0xa8193c198aab4146e3640f414ba8473918c6d028f45b27fb08b185a16c15ce23.
//
// Solidity: event ScheduleTokenLiquidationThresholdRamp(address indexed token, uint16 liquidationThresholdInitial, uint16 liquidationThresholdFinal, uint40 timestampRampStart, uint40 timestampRampEnd)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchScheduleTokenLiquidationThresholdRamp(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "ScheduleTokenLiquidationThresholdRamp", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ScheduleTokenLiquidationThresholdRamp", log); err != nil {
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

// ParseScheduleTokenLiquidationThresholdRamp is a log parse operation binding the contract event 0xa8193c198aab4146e3640f414ba8473918c6d028f45b27fb08b185a16c15ce23.
//
// Solidity: event ScheduleTokenLiquidationThresholdRamp(address indexed token, uint16 liquidationThresholdInitial, uint16 liquidationThresholdFinal, uint40 timestampRampStart, uint40 timestampRampEnd)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseScheduleTokenLiquidationThresholdRamp(log types.Log) (*CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp, error) {
	event := new(CreditConfiguratorv310ScheduleTokenLiquidationThresholdRamp)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "ScheduleTokenLiquidationThresholdRamp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetBorrowingLimitsIterator is returned from FilterSetBorrowingLimits and is used to iterate over the raw logs and unpacked data for SetBorrowingLimits events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetBorrowingLimitsIterator struct {
	Event *CreditConfiguratorv310SetBorrowingLimits // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetBorrowingLimitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetBorrowingLimits)
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
		it.Event = new(CreditConfiguratorv310SetBorrowingLimits)
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
func (it *CreditConfiguratorv310SetBorrowingLimitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetBorrowingLimitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetBorrowingLimits represents a SetBorrowingLimits event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetBorrowingLimits struct {
	MinDebt *big.Int
	MaxDebt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowingLimits is a free log retrieval operation binding the contract event 0xb2cc80ffa4c2f75731dbb99fcd29cccd7829c55d4cd5d6a884506b1435d6d1f3.
//
// Solidity: event SetBorrowingLimits(uint256 minDebt, uint256 maxDebt)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetBorrowingLimits(opts *bind.FilterOpts) (*CreditConfiguratorv310SetBorrowingLimitsIterator, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetBorrowingLimits")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetBorrowingLimitsIterator{contract: _CreditConfiguratorv310.contract, event: "SetBorrowingLimits", logs: logs, sub: sub}, nil
}

// WatchSetBorrowingLimits is a free log subscription operation binding the contract event 0xb2cc80ffa4c2f75731dbb99fcd29cccd7829c55d4cd5d6a884506b1435d6d1f3.
//
// Solidity: event SetBorrowingLimits(uint256 minDebt, uint256 maxDebt)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetBorrowingLimits(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetBorrowingLimits) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetBorrowingLimits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetBorrowingLimits)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetBorrowingLimits", log); err != nil {
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

// ParseSetBorrowingLimits is a log parse operation binding the contract event 0xb2cc80ffa4c2f75731dbb99fcd29cccd7829c55d4cd5d6a884506b1435d6d1f3.
//
// Solidity: event SetBorrowingLimits(uint256 minDebt, uint256 maxDebt)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetBorrowingLimits(log types.Log) (*CreditConfiguratorv310SetBorrowingLimits, error) {
	event := new(CreditConfiguratorv310SetBorrowingLimits)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetBorrowingLimits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetCreditFacadeIterator is returned from FilterSetCreditFacade and is used to iterate over the raw logs and unpacked data for SetCreditFacade events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetCreditFacadeIterator struct {
	Event *CreditConfiguratorv310SetCreditFacade // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetCreditFacadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetCreditFacade)
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
		it.Event = new(CreditConfiguratorv310SetCreditFacade)
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
func (it *CreditConfiguratorv310SetCreditFacadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetCreditFacadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetCreditFacade represents a SetCreditFacade event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetCreditFacade struct {
	CreditFacade common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetCreditFacade is a free log retrieval operation binding the contract event 0x1cd439329e916b95ce297eb699326f2799c8de28be6bba10f28db1d9067778f1.
//
// Solidity: event SetCreditFacade(address indexed creditFacade)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetCreditFacade(opts *bind.FilterOpts, creditFacade []common.Address) (*CreditConfiguratorv310SetCreditFacadeIterator, error) {

	var creditFacadeRule []interface{}
	for _, creditFacadeItem := range creditFacade {
		creditFacadeRule = append(creditFacadeRule, creditFacadeItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetCreditFacade", creditFacadeRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetCreditFacadeIterator{contract: _CreditConfiguratorv310.contract, event: "SetCreditFacade", logs: logs, sub: sub}, nil
}

// WatchSetCreditFacade is a free log subscription operation binding the contract event 0x1cd439329e916b95ce297eb699326f2799c8de28be6bba10f28db1d9067778f1.
//
// Solidity: event SetCreditFacade(address indexed creditFacade)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetCreditFacade(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetCreditFacade, creditFacade []common.Address) (event.Subscription, error) {

	var creditFacadeRule []interface{}
	for _, creditFacadeItem := range creditFacade {
		creditFacadeRule = append(creditFacadeRule, creditFacadeItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetCreditFacade", creditFacadeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetCreditFacade)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetCreditFacade", log); err != nil {
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

// ParseSetCreditFacade is a log parse operation binding the contract event 0x1cd439329e916b95ce297eb699326f2799c8de28be6bba10f28db1d9067778f1.
//
// Solidity: event SetCreditFacade(address indexed creditFacade)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetCreditFacade(log types.Log) (*CreditConfiguratorv310SetCreditFacade, error) {
	event := new(CreditConfiguratorv310SetCreditFacade)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetCreditFacade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetExpirationDateIterator is returned from FilterSetExpirationDate and is used to iterate over the raw logs and unpacked data for SetExpirationDate events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetExpirationDateIterator struct {
	Event *CreditConfiguratorv310SetExpirationDate // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetExpirationDateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetExpirationDate)
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
		it.Event = new(CreditConfiguratorv310SetExpirationDate)
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
func (it *CreditConfiguratorv310SetExpirationDateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetExpirationDateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetExpirationDate represents a SetExpirationDate event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetExpirationDate struct {
	ExpirationDate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetExpirationDate is a free log retrieval operation binding the contract event 0xb019cf1dc4b3caa72aa4723abcc271a2bb3138bee0a89cd911fb8980b0c93d56.
//
// Solidity: event SetExpirationDate(uint40 expirationDate)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetExpirationDate(opts *bind.FilterOpts) (*CreditConfiguratorv310SetExpirationDateIterator, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetExpirationDate")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetExpirationDateIterator{contract: _CreditConfiguratorv310.contract, event: "SetExpirationDate", logs: logs, sub: sub}, nil
}

// WatchSetExpirationDate is a free log subscription operation binding the contract event 0xb019cf1dc4b3caa72aa4723abcc271a2bb3138bee0a89cd911fb8980b0c93d56.
//
// Solidity: event SetExpirationDate(uint40 expirationDate)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetExpirationDate(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetExpirationDate) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetExpirationDate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetExpirationDate)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetExpirationDate", log); err != nil {
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

// ParseSetExpirationDate is a log parse operation binding the contract event 0xb019cf1dc4b3caa72aa4723abcc271a2bb3138bee0a89cd911fb8980b0c93d56.
//
// Solidity: event SetExpirationDate(uint40 expirationDate)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetExpirationDate(log types.Log) (*CreditConfiguratorv310SetExpirationDate, error) {
	event := new(CreditConfiguratorv310SetExpirationDate)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetExpirationDate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetLossLiquidatorIterator is returned from FilterSetLossLiquidator and is used to iterate over the raw logs and unpacked data for SetLossLiquidator events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetLossLiquidatorIterator struct {
	Event *CreditConfiguratorv310SetLossLiquidator // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetLossLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetLossLiquidator)
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
		it.Event = new(CreditConfiguratorv310SetLossLiquidator)
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
func (it *CreditConfiguratorv310SetLossLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetLossLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetLossLiquidator represents a SetLossLiquidator event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetLossLiquidator struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLossLiquidator is a free log retrieval operation binding the contract event 0x6cfb8aa80a47f60adc07603cbb475f1a47984baf8dd6fae5cc68cf24b9ef693c.
//
// Solidity: event SetLossLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetLossLiquidator(opts *bind.FilterOpts, liquidator []common.Address) (*CreditConfiguratorv310SetLossLiquidatorIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetLossLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetLossLiquidatorIterator{contract: _CreditConfiguratorv310.contract, event: "SetLossLiquidator", logs: logs, sub: sub}, nil
}

// WatchSetLossLiquidator is a free log subscription operation binding the contract event 0x6cfb8aa80a47f60adc07603cbb475f1a47984baf8dd6fae5cc68cf24b9ef693c.
//
// Solidity: event SetLossLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetLossLiquidator(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetLossLiquidator, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetLossLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetLossLiquidator)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetLossLiquidator", log); err != nil {
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

// ParseSetLossLiquidator is a log parse operation binding the contract event 0x6cfb8aa80a47f60adc07603cbb475f1a47984baf8dd6fae5cc68cf24b9ef693c.
//
// Solidity: event SetLossLiquidator(address indexed liquidator)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetLossLiquidator(log types.Log) (*CreditConfiguratorv310SetLossLiquidator, error) {
	event := new(CreditConfiguratorv310SetLossLiquidator)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetLossLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator is returned from FilterSetMaxDebtPerBlockMultiplier and is used to iterate over the raw logs and unpacked data for SetMaxDebtPerBlockMultiplier events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator struct {
	Event *CreditConfiguratorv310SetMaxDebtPerBlockMultiplier // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetMaxDebtPerBlockMultiplier)
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
		it.Event = new(CreditConfiguratorv310SetMaxDebtPerBlockMultiplier)
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
func (it *CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetMaxDebtPerBlockMultiplier represents a SetMaxDebtPerBlockMultiplier event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetMaxDebtPerBlockMultiplier struct {
	MaxDebtPerBlockMultiplier uint8
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetMaxDebtPerBlockMultiplier is a free log retrieval operation binding the contract event 0xaebbd82c9dcdcd553331f5850bbdf5add33bf8fce5c7c76e2c9e7912ad5f1564.
//
// Solidity: event SetMaxDebtPerBlockMultiplier(uint8 maxDebtPerBlockMultiplier)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetMaxDebtPerBlockMultiplier(opts *bind.FilterOpts) (*CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetMaxDebtPerBlockMultiplier")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetMaxDebtPerBlockMultiplierIterator{contract: _CreditConfiguratorv310.contract, event: "SetMaxDebtPerBlockMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetMaxDebtPerBlockMultiplier is a free log subscription operation binding the contract event 0xaebbd82c9dcdcd553331f5850bbdf5add33bf8fce5c7c76e2c9e7912ad5f1564.
//
// Solidity: event SetMaxDebtPerBlockMultiplier(uint8 maxDebtPerBlockMultiplier)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetMaxDebtPerBlockMultiplier(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetMaxDebtPerBlockMultiplier) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetMaxDebtPerBlockMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetMaxDebtPerBlockMultiplier)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetMaxDebtPerBlockMultiplier", log); err != nil {
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

// ParseSetMaxDebtPerBlockMultiplier is a log parse operation binding the contract event 0xaebbd82c9dcdcd553331f5850bbdf5add33bf8fce5c7c76e2c9e7912ad5f1564.
//
// Solidity: event SetMaxDebtPerBlockMultiplier(uint8 maxDebtPerBlockMultiplier)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetMaxDebtPerBlockMultiplier(log types.Log) (*CreditConfiguratorv310SetMaxDebtPerBlockMultiplier, error) {
	event := new(CreditConfiguratorv310SetMaxDebtPerBlockMultiplier)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetMaxDebtPerBlockMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetPriceOracleIterator is returned from FilterSetPriceOracle and is used to iterate over the raw logs and unpacked data for SetPriceOracle events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetPriceOracleIterator struct {
	Event *CreditConfiguratorv310SetPriceOracle // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetPriceOracle)
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
		it.Event = new(CreditConfiguratorv310SetPriceOracle)
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
func (it *CreditConfiguratorv310SetPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetPriceOracle represents a SetPriceOracle event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetPriceOracle struct {
	PriceOracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetPriceOracle is a free log retrieval operation binding the contract event 0x88a686e0e341d9099f2f990c3aa759a86822142a67579064b43ded9354a25662.
//
// Solidity: event SetPriceOracle(address indexed priceOracle)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetPriceOracle(opts *bind.FilterOpts, priceOracle []common.Address) (*CreditConfiguratorv310SetPriceOracleIterator, error) {

	var priceOracleRule []interface{}
	for _, priceOracleItem := range priceOracle {
		priceOracleRule = append(priceOracleRule, priceOracleItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetPriceOracle", priceOracleRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetPriceOracleIterator{contract: _CreditConfiguratorv310.contract, event: "SetPriceOracle", logs: logs, sub: sub}, nil
}

// WatchSetPriceOracle is a free log subscription operation binding the contract event 0x88a686e0e341d9099f2f990c3aa759a86822142a67579064b43ded9354a25662.
//
// Solidity: event SetPriceOracle(address indexed priceOracle)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetPriceOracle(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetPriceOracle, priceOracle []common.Address) (event.Subscription, error) {

	var priceOracleRule []interface{}
	for _, priceOracleItem := range priceOracle {
		priceOracleRule = append(priceOracleRule, priceOracleItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetPriceOracle", priceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetPriceOracle)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetPriceOracle", log); err != nil {
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

// ParseSetPriceOracle is a log parse operation binding the contract event 0x88a686e0e341d9099f2f990c3aa759a86822142a67579064b43ded9354a25662.
//
// Solidity: event SetPriceOracle(address indexed priceOracle)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetPriceOracle(log types.Log) (*CreditConfiguratorv310SetPriceOracle, error) {
	event := new(CreditConfiguratorv310SetPriceOracle)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310SetTokenLiquidationThresholdIterator is returned from FilterSetTokenLiquidationThreshold and is used to iterate over the raw logs and unpacked data for SetTokenLiquidationThreshold events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetTokenLiquidationThresholdIterator struct {
	Event *CreditConfiguratorv310SetTokenLiquidationThreshold // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310SetTokenLiquidationThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310SetTokenLiquidationThreshold)
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
		it.Event = new(CreditConfiguratorv310SetTokenLiquidationThreshold)
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
func (it *CreditConfiguratorv310SetTokenLiquidationThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310SetTokenLiquidationThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310SetTokenLiquidationThreshold represents a SetTokenLiquidationThreshold event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310SetTokenLiquidationThreshold struct {
	Token                common.Address
	LiquidationThreshold uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTokenLiquidationThreshold is a free log retrieval operation binding the contract event 0xda5e841a0cb137f4a60661969e409f01ef7627723a4a929414e4f69b5475ee8c.
//
// Solidity: event SetTokenLiquidationThreshold(address indexed token, uint16 liquidationThreshold)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterSetTokenLiquidationThreshold(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv310SetTokenLiquidationThresholdIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "SetTokenLiquidationThreshold", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310SetTokenLiquidationThresholdIterator{contract: _CreditConfiguratorv310.contract, event: "SetTokenLiquidationThreshold", logs: logs, sub: sub}, nil
}

// WatchSetTokenLiquidationThreshold is a free log subscription operation binding the contract event 0xda5e841a0cb137f4a60661969e409f01ef7627723a4a929414e4f69b5475ee8c.
//
// Solidity: event SetTokenLiquidationThreshold(address indexed token, uint16 liquidationThreshold)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchSetTokenLiquidationThreshold(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310SetTokenLiquidationThreshold, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "SetTokenLiquidationThreshold", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310SetTokenLiquidationThreshold)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetTokenLiquidationThreshold", log); err != nil {
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

// ParseSetTokenLiquidationThreshold is a log parse operation binding the contract event 0xda5e841a0cb137f4a60661969e409f01ef7627723a4a929414e4f69b5475ee8c.
//
// Solidity: event SetTokenLiquidationThreshold(address indexed token, uint16 liquidationThreshold)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseSetTokenLiquidationThreshold(log types.Log) (*CreditConfiguratorv310SetTokenLiquidationThreshold, error) {
	event := new(CreditConfiguratorv310SetTokenLiquidationThreshold)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "SetTokenLiquidationThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv310UpdateFeesIterator is returned from FilterUpdateFees and is used to iterate over the raw logs and unpacked data for UpdateFees events raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310UpdateFeesIterator struct {
	Event *CreditConfiguratorv310UpdateFees // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv310UpdateFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv310UpdateFees)
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
		it.Event = new(CreditConfiguratorv310UpdateFees)
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
func (it *CreditConfiguratorv310UpdateFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv310UpdateFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv310UpdateFees represents a UpdateFees event raised by the CreditConfiguratorv310 contract.
type CreditConfiguratorv310UpdateFees struct {
	FeeLiquidation            uint16
	LiquidationPremium        uint16
	FeeLiquidationExpired     uint16
	LiquidationPremiumExpired uint16
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterUpdateFees is a free log retrieval operation binding the contract event 0x2d179a43c34a4a80c102e61bb259930222752df8bdfc749e5f2fd6ef9dab971c.
//
// Solidity: event UpdateFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) FilterUpdateFees(opts *bind.FilterOpts) (*CreditConfiguratorv310UpdateFeesIterator, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.FilterLogs(opts, "UpdateFees")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv310UpdateFeesIterator{contract: _CreditConfiguratorv310.contract, event: "UpdateFees", logs: logs, sub: sub}, nil
}

// WatchUpdateFees is a free log subscription operation binding the contract event 0x2d179a43c34a4a80c102e61bb259930222752df8bdfc749e5f2fd6ef9dab971c.
//
// Solidity: event UpdateFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) WatchUpdateFees(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv310UpdateFees) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv310.contract.WatchLogs(opts, "UpdateFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv310UpdateFees)
				if err := _CreditConfiguratorv310.contract.UnpackLog(event, "UpdateFees", log); err != nil {
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

// ParseUpdateFees is a log parse operation binding the contract event 0x2d179a43c34a4a80c102e61bb259930222752df8bdfc749e5f2fd6ef9dab971c.
//
// Solidity: event UpdateFees(uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv310 *CreditConfiguratorv310Filterer) ParseUpdateFees(log types.Log) (*CreditConfiguratorv310UpdateFees, error) {
	event := new(CreditConfiguratorv310UpdateFees)
	if err := _CreditConfiguratorv310.contract.UnpackLog(event, "UpdateFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
