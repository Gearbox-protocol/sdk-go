// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditConfigurator

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

// CollateralToken is an auto generated low-level Go binding around an user-defined struct.
type CollateralToken struct {
	Token                common.Address
	LiquidationThreshold uint16
}

// CreditManagerOpts is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerOpts struct {
	MinBorrowedAmount *big.Int
	MaxBorrowedAmount *big.Int
	CollateralTokens  []CollateralToken
	DegenNFT          common.Address
}

// CreditConfiguratorMetaData contains all meta data concerning the CreditConfigurator contract.
var CreditConfiguratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCreditManager\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"contractCreditFacade\",\"name\":\"_creditFacade\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"minBorrowedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxBorrowedAmount\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"}],\"internalType\":\"structCreditManagerOpts\",\"name\":\"opts\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdapterUsedTwiceException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChiThresholdMoreOneException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditManagerOrFacadeUsedAsAllowContractsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DegenNFTnotSetException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FastCheckNotCoverCollateralDropException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectFeesException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectLimitsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectLiquidationThresholdException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SetLTForUnderlyingException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ContractAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"}],\"name\":\"ContractForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCreditConfigurator\",\"type\":\"address\"}],\"name\":\"CreditConfiguratorUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCreditFacade\",\"type\":\"address\"}],\"name\":\"CreditFacadeUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"DegenModeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationPremium\",\"type\":\"uint16\"}],\"name\":\"FeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"IncreaseDebtModeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"LimitPerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBorrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxBorrowedAmount\",\"type\":\"uint256\"}],\"name\":\"LimitsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"PriceOracleUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenForbidden\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidityThreshold\",\"type\":\"uint16\"}],\"name\":\"TokenLiquidationThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acl\",\"outputs\":[{\"internalType\":\"contractACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"addCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractIAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"allowedContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedContractsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"contractCreditFacade\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractCreditManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"forbidContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"forbidToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_liquidationPremium\",\"type\":\"uint16\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_mode\",\"type\":\"bool\"}],\"name\":\"setIncreaseDebtForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newLimit\",\"type\":\"uint128\"}],\"name\":\"setLimitPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_minBorrowedAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_maxBorrowedAmount\",\"type\":\"uint128\"}],\"name\":\"setLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"setLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditConfigurator\",\"type\":\"address\"}],\"name\":\"upgradeCreditConfigurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditFacade\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"migrateLimits\",\"type\":\"bool\"}],\"name\":\"upgradeCreditFacade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradePriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditConfiguratorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditConfiguratorMetaData.ABI instead.
var CreditConfiguratorABI = CreditConfiguratorMetaData.ABI

// CreditConfigurator is an auto generated Go binding around an Ethereum contract.
type CreditConfigurator struct {
	CreditConfiguratorCaller     // Read-only binding to the contract
	CreditConfiguratorTransactor // Write-only binding to the contract
	CreditConfiguratorFilterer   // Log filterer for contract events
}

// CreditConfiguratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditConfiguratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditConfiguratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditConfiguratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditConfiguratorSession struct {
	Contract     *CreditConfigurator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CreditConfiguratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditConfiguratorCallerSession struct {
	Contract *CreditConfiguratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CreditConfiguratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditConfiguratorTransactorSession struct {
	Contract     *CreditConfiguratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CreditConfiguratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditConfiguratorRaw struct {
	Contract *CreditConfigurator // Generic contract binding to access the raw methods on
}

// CreditConfiguratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditConfiguratorCallerRaw struct {
	Contract *CreditConfiguratorCaller // Generic read-only contract binding to access the raw methods on
}

// CreditConfiguratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditConfiguratorTransactorRaw struct {
	Contract *CreditConfiguratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditConfigurator creates a new instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfigurator(address common.Address, backend bind.ContractBackend) (*CreditConfigurator, error) {
	contract, err := bindCreditConfigurator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator{CreditConfiguratorCaller: CreditConfiguratorCaller{contract: contract}, CreditConfiguratorTransactor: CreditConfiguratorTransactor{contract: contract}, CreditConfiguratorFilterer: CreditConfiguratorFilterer{contract: contract}}, nil
}

// NewCreditConfiguratorCaller creates a new read-only instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorCaller(address common.Address, caller bind.ContractCaller) (*CreditConfiguratorCaller, error) {
	contract, err := bindCreditConfigurator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCaller{contract: contract}, nil
}

// NewCreditConfiguratorTransactor creates a new write-only instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditConfiguratorTransactor, error) {
	contract, err := bindCreditConfigurator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTransactor{contract: contract}, nil
}

// NewCreditConfiguratorFilterer creates a new log filterer instance of CreditConfigurator, bound to a specific deployed contract.
func NewCreditConfiguratorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditConfiguratorFilterer, error) {
	contract, err := bindCreditConfigurator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorFilterer{contract: contract}, nil
}

// bindCreditConfigurator binds a generic wrapper to an already deployed contract.
func bindCreditConfigurator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditConfiguratorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator *CreditConfiguratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator.Contract.CreditConfiguratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator *CreditConfiguratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.CreditConfiguratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator *CreditConfiguratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.CreditConfiguratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator *CreditConfiguratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator *CreditConfiguratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator *CreditConfiguratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "_acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) Acl() (common.Address, error) {
	return _CreditConfigurator.Contract.Acl(&_CreditConfigurator.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Acl() (common.Address, error) {
	return _CreditConfigurator.Contract.Acl(&_CreditConfigurator.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) AddressProvider() (common.Address, error) {
	return _CreditConfigurator.Contract.AddressProvider(&_CreditConfigurator.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AddressProvider() (common.Address, error) {
	return _CreditConfigurator.Contract.AddressProvider(&_CreditConfigurator.CallOpts)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) AllowedContracts(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "allowedContracts", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditConfigurator.Contract.AllowedContracts(&_CreditConfigurator.CallOpts, i)
}

// AllowedContracts is a free data retrieval call binding the contract method 0x5094cb4f.
//
// Solidity: function allowedContracts(uint256 i) view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AllowedContracts(i *big.Int) (common.Address, error) {
	return _CreditConfigurator.Contract.AllowedContracts(&_CreditConfigurator.CallOpts, i)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCaller) AllowedContractsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "allowedContractsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditConfigurator.Contract.AllowedContractsCount(&_CreditConfigurator.CallOpts)
}

// AllowedContractsCount is a free data retrieval call binding the contract method 0x50e036ff.
//
// Solidity: function allowedContractsCount() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCallerSession) AllowedContractsCount() (*big.Int, error) {
	return _CreditConfigurator.Contract.AllowedContractsCount(&_CreditConfigurator.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) CreditFacade() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditFacade(&_CreditConfigurator.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) CreditFacade() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditFacade(&_CreditConfigurator.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) CreditManager() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditManager(&_CreditConfigurator.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) CreditManager() (common.Address, error) {
	return _CreditConfigurator.Contract.CreditManager(&_CreditConfigurator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorSession) Paused() (bool, error) {
	return _CreditConfigurator.Contract.Paused(&_CreditConfigurator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Paused() (bool, error) {
	return _CreditConfigurator.Contract.Paused(&_CreditConfigurator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfigurator *CreditConfiguratorSession) Underlying() (common.Address, error) {
	return _CreditConfigurator.Contract.Underlying(&_CreditConfigurator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Underlying() (common.Address, error) {
	return _CreditConfigurator.Contract.Underlying(&_CreditConfigurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfigurator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorSession) Version() (*big.Int, error) {
	return _CreditConfigurator.Contract.Version(&_CreditConfigurator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfigurator *CreditConfiguratorCallerSession) Version() (*big.Int, error) {
	return _CreditConfigurator.Contract.Version(&_CreditConfigurator.CallOpts)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AddCollateralToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "addCollateralToken", token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AddCollateralToken(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AddCollateralToken(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AllowContract(opts *bind.TransactOpts, targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "allowContract", targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowContract(&_CreditConfigurator.TransactOpts, targetContract, adapter)
}

// AllowContract is a paid mutator transaction binding the contract method 0x7bccacee.
//
// Solidity: function allowContract(address targetContract, address adapter) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AllowContract(targetContract common.Address, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowContract(&_CreditConfigurator.TransactOpts, targetContract, adapter)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) AllowToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "allowToken", token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowToken(&_CreditConfigurator.TransactOpts, token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.AllowToken(&_CreditConfigurator.TransactOpts, token)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) ForbidContract(opts *bind.TransactOpts, targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "forbidContract", targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidContract(&_CreditConfigurator.TransactOpts, targetContract)
}

// ForbidContract is a paid mutator transaction binding the contract method 0x52438e54.
//
// Solidity: function forbidContract(address targetContract) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) ForbidContract(targetContract common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidContract(&_CreditConfigurator.TransactOpts, targetContract)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidToken(&_CreditConfigurator.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.ForbidToken(&_CreditConfigurator.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorSession) Pause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Pause(&_CreditConfigurator.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Pause(&_CreditConfigurator.TransactOpts)
}

// SetFees is a paid mutator transaction binding the contract method 0x534c6bcb.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetFees(opts *bind.TransactOpts, _feeInterest uint16, _feeLiquidation uint16, _liquidationPremium uint16) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setFees", _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetFees is a paid mutator transaction binding the contract method 0x534c6bcb.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetFees(_feeInterest uint16, _feeLiquidation uint16, _liquidationPremium uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFees(&_CreditConfigurator.TransactOpts, _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetFees is a paid mutator transaction binding the contract method 0x534c6bcb.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationPremium) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetFees(_feeInterest uint16, _feeLiquidation uint16, _liquidationPremium uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetFees(&_CreditConfigurator.TransactOpts, _feeInterest, _feeLiquidation, _liquidationPremium)
}

// SetIncreaseDebtForbidden is a paid mutator transaction binding the contract method 0xffd9b907.
//
// Solidity: function setIncreaseDebtForbidden(bool _mode) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetIncreaseDebtForbidden(opts *bind.TransactOpts, _mode bool) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setIncreaseDebtForbidden", _mode)
}

// SetIncreaseDebtForbidden is a paid mutator transaction binding the contract method 0xffd9b907.
//
// Solidity: function setIncreaseDebtForbidden(bool _mode) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetIncreaseDebtForbidden(_mode bool) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetIncreaseDebtForbidden(&_CreditConfigurator.TransactOpts, _mode)
}

// SetIncreaseDebtForbidden is a paid mutator transaction binding the contract method 0xffd9b907.
//
// Solidity: function setIncreaseDebtForbidden(bool _mode) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetIncreaseDebtForbidden(_mode bool) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetIncreaseDebtForbidden(&_CreditConfigurator.TransactOpts, _mode)
}

// SetLimitPerBlock is a paid mutator transaction binding the contract method 0x9c55a054.
//
// Solidity: function setLimitPerBlock(uint128 newLimit) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetLimitPerBlock(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setLimitPerBlock", newLimit)
}

// SetLimitPerBlock is a paid mutator transaction binding the contract method 0x9c55a054.
//
// Solidity: function setLimitPerBlock(uint128 newLimit) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetLimitPerBlock(newLimit *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimitPerBlock(&_CreditConfigurator.TransactOpts, newLimit)
}

// SetLimitPerBlock is a paid mutator transaction binding the contract method 0x9c55a054.
//
// Solidity: function setLimitPerBlock(uint128 newLimit) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetLimitPerBlock(newLimit *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimitPerBlock(&_CreditConfigurator.TransactOpts, newLimit)
}

// SetLimits is a paid mutator transaction binding the contract method 0x9f1c5069.
//
// Solidity: function setLimits(uint128 _minBorrowedAmount, uint128 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetLimits(opts *bind.TransactOpts, _minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setLimits", _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLimits is a paid mutator transaction binding the contract method 0x9f1c5069.
//
// Solidity: function setLimits(uint128 _minBorrowedAmount, uint128 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetLimits(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimits(&_CreditConfigurator.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLimits is a paid mutator transaction binding the contract method 0x9f1c5069.
//
// Solidity: function setLimits(uint128 _minBorrowedAmount, uint128 _maxBorrowedAmount) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetLimits(_minBorrowedAmount *big.Int, _maxBorrowedAmount *big.Int) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLimits(&_CreditConfigurator.TransactOpts, _minBorrowedAmount, _maxBorrowedAmount)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLiquidationThreshold(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.SetLiquidationThreshold(&_CreditConfigurator.TransactOpts, token, liquidationThreshold)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorSession) Unpause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Unpause(&_CreditConfigurator.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.Unpause(&_CreditConfigurator.TransactOpts)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradeCreditConfigurator(opts *bind.TransactOpts, _creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradeCreditConfigurator", _creditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradeCreditConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditConfigurator(&_CreditConfigurator.TransactOpts, _creditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address _creditConfigurator) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradeCreditConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditConfigurator(&_CreditConfigurator.TransactOpts, _creditConfigurator)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x526a41e8.
//
// Solidity: function upgradeCreditFacade(address _creditFacade, bool migrateLimits) returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradeCreditFacade(opts *bind.TransactOpts, _creditFacade common.Address, migrateLimits bool) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradeCreditFacade", _creditFacade, migrateLimits)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x526a41e8.
//
// Solidity: function upgradeCreditFacade(address _creditFacade, bool migrateLimits) returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradeCreditFacade(_creditFacade common.Address, migrateLimits bool) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditFacade(&_CreditConfigurator.TransactOpts, _creditFacade, migrateLimits)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x526a41e8.
//
// Solidity: function upgradeCreditFacade(address _creditFacade, bool migrateLimits) returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradeCreditFacade(_creditFacade common.Address, migrateLimits bool) (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradeCreditFacade(&_CreditConfigurator.TransactOpts, _creditFacade, migrateLimits)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorTransactor) UpgradePriceOracle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator.contract.Transact(opts, "upgradePriceOracle")
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradePriceOracle(&_CreditConfigurator.TransactOpts)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xf0527ac6.
//
// Solidity: function upgradePriceOracle() returns()
func (_CreditConfigurator *CreditConfiguratorTransactorSession) UpgradePriceOracle() (*types.Transaction, error) {
	return _CreditConfigurator.Contract.UpgradePriceOracle(&_CreditConfigurator.TransactOpts)
}

// CreditConfiguratorContractAllowedIterator is returned from FilterContractAllowed and is used to iterate over the raw logs and unpacked data for ContractAllowed events raised by the CreditConfigurator contract.
type CreditConfiguratorContractAllowedIterator struct {
	Event *CreditConfiguratorContractAllowed // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorContractAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorContractAllowed)
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
		it.Event = new(CreditConfiguratorContractAllowed)
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
func (it *CreditConfiguratorContractAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorContractAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorContractAllowed represents a ContractAllowed event raised by the CreditConfigurator contract.
type CreditConfiguratorContractAllowed struct {
	Protocol common.Address
	Adapter  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAllowed is a free log retrieval operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterContractAllowed(opts *bind.FilterOpts, protocol []common.Address, adapter []common.Address) (*CreditConfiguratorContractAllowedIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorContractAllowedIterator{contract: _CreditConfigurator.contract, event: "ContractAllowed", logs: logs, sub: sub}, nil
}

// WatchContractAllowed is a free log subscription operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchContractAllowed(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorContractAllowed, protocol []common.Address, adapter []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "ContractAllowed", protocolRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorContractAllowed)
				if err := _CreditConfigurator.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
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

// ParseContractAllowed is a log parse operation binding the contract event 0x4bcbefaef68b99503d502f5a6abe7bca2b183ab8ac55457013c77d084ebd1305.
//
// Solidity: event ContractAllowed(address indexed protocol, address indexed adapter)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseContractAllowed(log types.Log) (*CreditConfiguratorContractAllowed, error) {
	event := new(CreditConfiguratorContractAllowed)
	if err := _CreditConfigurator.contract.UnpackLog(event, "ContractAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorContractForbiddenIterator is returned from FilterContractForbidden and is used to iterate over the raw logs and unpacked data for ContractForbidden events raised by the CreditConfigurator contract.
type CreditConfiguratorContractForbiddenIterator struct {
	Event *CreditConfiguratorContractForbidden // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorContractForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorContractForbidden)
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
		it.Event = new(CreditConfiguratorContractForbidden)
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
func (it *CreditConfiguratorContractForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorContractForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorContractForbidden represents a ContractForbidden event raised by the CreditConfigurator contract.
type CreditConfiguratorContractForbidden struct {
	Protocol common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractForbidden is a free log retrieval operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterContractForbidden(opts *bind.FilterOpts, protocol []common.Address) (*CreditConfiguratorContractForbiddenIterator, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorContractForbiddenIterator{contract: _CreditConfigurator.contract, event: "ContractForbidden", logs: logs, sub: sub}, nil
}

// WatchContractForbidden is a free log subscription operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchContractForbidden(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorContractForbidden, protocol []common.Address) (event.Subscription, error) {

	var protocolRule []interface{}
	for _, protocolItem := range protocol {
		protocolRule = append(protocolRule, protocolItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "ContractForbidden", protocolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorContractForbidden)
				if err := _CreditConfigurator.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
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

// ParseContractForbidden is a log parse operation binding the contract event 0xab9f405bf0c19b97f65a7031634db41569cd2f0e0376a610a1e977f9ab22b58f.
//
// Solidity: event ContractForbidden(address indexed protocol)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseContractForbidden(log types.Log) (*CreditConfiguratorContractForbidden, error) {
	event := new(CreditConfiguratorContractForbidden)
	if err := _CreditConfigurator.contract.UnpackLog(event, "ContractForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorCreditConfiguratorUpgradedIterator is returned from FilterCreditConfiguratorUpgraded and is used to iterate over the raw logs and unpacked data for CreditConfiguratorUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorCreditConfiguratorUpgradedIterator struct {
	Event *CreditConfiguratorCreditConfiguratorUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorCreditConfiguratorUpgraded)
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
		it.Event = new(CreditConfiguratorCreditConfiguratorUpgraded)
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
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorCreditConfiguratorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorCreditConfiguratorUpgraded represents a CreditConfiguratorUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorCreditConfiguratorUpgraded struct {
	NewCreditConfigurator common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreditConfiguratorUpgraded is a free log retrieval operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterCreditConfiguratorUpgraded(opts *bind.FilterOpts, newCreditConfigurator []common.Address) (*CreditConfiguratorCreditConfiguratorUpgradedIterator, error) {

	var newCreditConfiguratorRule []interface{}
	for _, newCreditConfiguratorItem := range newCreditConfigurator {
		newCreditConfiguratorRule = append(newCreditConfiguratorRule, newCreditConfiguratorItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "CreditConfiguratorUpgraded", newCreditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCreditConfiguratorUpgradedIterator{contract: _CreditConfigurator.contract, event: "CreditConfiguratorUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditConfiguratorUpgraded is a free log subscription operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchCreditConfiguratorUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorCreditConfiguratorUpgraded, newCreditConfigurator []common.Address) (event.Subscription, error) {

	var newCreditConfiguratorRule []interface{}
	for _, newCreditConfiguratorItem := range newCreditConfigurator {
		newCreditConfiguratorRule = append(newCreditConfiguratorRule, newCreditConfiguratorItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "CreditConfiguratorUpgraded", newCreditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorCreditConfiguratorUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
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
// Solidity: event CreditConfiguratorUpgraded(address indexed newCreditConfigurator)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseCreditConfiguratorUpgraded(log types.Log) (*CreditConfiguratorCreditConfiguratorUpgraded, error) {
	event := new(CreditConfiguratorCreditConfiguratorUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorCreditFacadeUpgradedIterator is returned from FilterCreditFacadeUpgraded and is used to iterate over the raw logs and unpacked data for CreditFacadeUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorCreditFacadeUpgradedIterator struct {
	Event *CreditConfiguratorCreditFacadeUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorCreditFacadeUpgraded)
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
		it.Event = new(CreditConfiguratorCreditFacadeUpgraded)
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
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorCreditFacadeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorCreditFacadeUpgraded represents a CreditFacadeUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorCreditFacadeUpgraded struct {
	NewCreditFacade common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreditFacadeUpgraded is a free log retrieval operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterCreditFacadeUpgraded(opts *bind.FilterOpts, newCreditFacade []common.Address) (*CreditConfiguratorCreditFacadeUpgradedIterator, error) {

	var newCreditFacadeRule []interface{}
	for _, newCreditFacadeItem := range newCreditFacade {
		newCreditFacadeRule = append(newCreditFacadeRule, newCreditFacadeItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "CreditFacadeUpgraded", newCreditFacadeRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorCreditFacadeUpgradedIterator{contract: _CreditConfigurator.contract, event: "CreditFacadeUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditFacadeUpgraded is a free log subscription operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchCreditFacadeUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorCreditFacadeUpgraded, newCreditFacade []common.Address) (event.Subscription, error) {

	var newCreditFacadeRule []interface{}
	for _, newCreditFacadeItem := range newCreditFacade {
		newCreditFacadeRule = append(newCreditFacadeRule, newCreditFacadeItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "CreditFacadeUpgraded", newCreditFacadeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorCreditFacadeUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "CreditFacadeUpgraded", log); err != nil {
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

// ParseCreditFacadeUpgraded is a log parse operation binding the contract event 0xa8b21f72cb83bce808df32dc2330217d744a1c22f3e9e44e4b11bbf049d37d9d.
//
// Solidity: event CreditFacadeUpgraded(address indexed newCreditFacade)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseCreditFacadeUpgraded(log types.Log) (*CreditConfiguratorCreditFacadeUpgraded, error) {
	event := new(CreditConfiguratorCreditFacadeUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "CreditFacadeUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorDegenModeUpdatedIterator is returned from FilterDegenModeUpdated and is used to iterate over the raw logs and unpacked data for DegenModeUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorDegenModeUpdatedIterator struct {
	Event *CreditConfiguratorDegenModeUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorDegenModeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorDegenModeUpdated)
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
		it.Event = new(CreditConfiguratorDegenModeUpdated)
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
func (it *CreditConfiguratorDegenModeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorDegenModeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorDegenModeUpdated represents a DegenModeUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorDegenModeUpdated struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDegenModeUpdated is a free log retrieval operation binding the contract event 0x58790fa9acf6d2945872a1c9a13aa9983af92345ab3c7b0c41afe45f02ac27fa.
//
// Solidity: event DegenModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterDegenModeUpdated(opts *bind.FilterOpts) (*CreditConfiguratorDegenModeUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "DegenModeUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorDegenModeUpdatedIterator{contract: _CreditConfigurator.contract, event: "DegenModeUpdated", logs: logs, sub: sub}, nil
}

// WatchDegenModeUpdated is a free log subscription operation binding the contract event 0x58790fa9acf6d2945872a1c9a13aa9983af92345ab3c7b0c41afe45f02ac27fa.
//
// Solidity: event DegenModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchDegenModeUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorDegenModeUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "DegenModeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorDegenModeUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "DegenModeUpdated", log); err != nil {
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

// ParseDegenModeUpdated is a log parse operation binding the contract event 0x58790fa9acf6d2945872a1c9a13aa9983af92345ab3c7b0c41afe45f02ac27fa.
//
// Solidity: event DegenModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseDegenModeUpdated(log types.Log) (*CreditConfiguratorDegenModeUpdated, error) {
	event := new(CreditConfiguratorDegenModeUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "DegenModeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorFeesUpdatedIterator is returned from FilterFeesUpdated and is used to iterate over the raw logs and unpacked data for FeesUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorFeesUpdatedIterator struct {
	Event *CreditConfiguratorFeesUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorFeesUpdated)
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
		it.Event = new(CreditConfiguratorFeesUpdated)
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
func (it *CreditConfiguratorFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorFeesUpdated represents a FeesUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorFeesUpdated struct {
	FeeInterest        uint16
	FeeLiquidation     uint16
	LiquidationPremium uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeesUpdated is a free log retrieval operation binding the contract event 0xb3ef341b591e573ddca7176a74bb92c8e453cce6d6885fcd6a544c2385d3811f.
//
// Solidity: event FeesUpdated(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterFeesUpdated(opts *bind.FilterOpts) (*CreditConfiguratorFeesUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorFeesUpdatedIterator{contract: _CreditConfigurator.contract, event: "FeesUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesUpdated is a free log subscription operation binding the contract event 0xb3ef341b591e573ddca7176a74bb92c8e453cce6d6885fcd6a544c2385d3811f.
//
// Solidity: event FeesUpdated(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchFeesUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorFeesUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
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

// ParseFeesUpdated is a log parse operation binding the contract event 0xb3ef341b591e573ddca7176a74bb92c8e453cce6d6885fcd6a544c2385d3811f.
//
// Solidity: event FeesUpdated(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseFeesUpdated(log types.Log) (*CreditConfiguratorFeesUpdated, error) {
	event := new(CreditConfiguratorFeesUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorIncreaseDebtModeUpdatedIterator is returned from FilterIncreaseDebtModeUpdated and is used to iterate over the raw logs and unpacked data for IncreaseDebtModeUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorIncreaseDebtModeUpdatedIterator struct {
	Event *CreditConfiguratorIncreaseDebtModeUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorIncreaseDebtModeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorIncreaseDebtModeUpdated)
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
		it.Event = new(CreditConfiguratorIncreaseDebtModeUpdated)
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
func (it *CreditConfiguratorIncreaseDebtModeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorIncreaseDebtModeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorIncreaseDebtModeUpdated represents a IncreaseDebtModeUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorIncreaseDebtModeUpdated struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIncreaseDebtModeUpdated is a free log retrieval operation binding the contract event 0xf97ba903470ff3262b4a68a938fe81def401bac1bd742169d2e8d5c4ed2185d3.
//
// Solidity: event IncreaseDebtModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterIncreaseDebtModeUpdated(opts *bind.FilterOpts) (*CreditConfiguratorIncreaseDebtModeUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "IncreaseDebtModeUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorIncreaseDebtModeUpdatedIterator{contract: _CreditConfigurator.contract, event: "IncreaseDebtModeUpdated", logs: logs, sub: sub}, nil
}

// WatchIncreaseDebtModeUpdated is a free log subscription operation binding the contract event 0xf97ba903470ff3262b4a68a938fe81def401bac1bd742169d2e8d5c4ed2185d3.
//
// Solidity: event IncreaseDebtModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchIncreaseDebtModeUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorIncreaseDebtModeUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "IncreaseDebtModeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorIncreaseDebtModeUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "IncreaseDebtModeUpdated", log); err != nil {
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

// ParseIncreaseDebtModeUpdated is a log parse operation binding the contract event 0xf97ba903470ff3262b4a68a938fe81def401bac1bd742169d2e8d5c4ed2185d3.
//
// Solidity: event IncreaseDebtModeUpdated(bool arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseIncreaseDebtModeUpdated(log types.Log) (*CreditConfiguratorIncreaseDebtModeUpdated, error) {
	event := new(CreditConfiguratorIncreaseDebtModeUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "IncreaseDebtModeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorLimitPerBlockUpdatedIterator is returned from FilterLimitPerBlockUpdated and is used to iterate over the raw logs and unpacked data for LimitPerBlockUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorLimitPerBlockUpdatedIterator struct {
	Event *CreditConfiguratorLimitPerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorLimitPerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorLimitPerBlockUpdated)
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
		it.Event = new(CreditConfiguratorLimitPerBlockUpdated)
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
func (it *CreditConfiguratorLimitPerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorLimitPerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorLimitPerBlockUpdated represents a LimitPerBlockUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorLimitPerBlockUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLimitPerBlockUpdated is a free log retrieval operation binding the contract event 0xa030eec20ae970821387224ec3bb15b4bd22b89bfb86dc4d2b3827aa3d99e48d.
//
// Solidity: event LimitPerBlockUpdated(uint128 arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterLimitPerBlockUpdated(opts *bind.FilterOpts) (*CreditConfiguratorLimitPerBlockUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "LimitPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorLimitPerBlockUpdatedIterator{contract: _CreditConfigurator.contract, event: "LimitPerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchLimitPerBlockUpdated is a free log subscription operation binding the contract event 0xa030eec20ae970821387224ec3bb15b4bd22b89bfb86dc4d2b3827aa3d99e48d.
//
// Solidity: event LimitPerBlockUpdated(uint128 arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchLimitPerBlockUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorLimitPerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "LimitPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorLimitPerBlockUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "LimitPerBlockUpdated", log); err != nil {
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

// ParseLimitPerBlockUpdated is a log parse operation binding the contract event 0xa030eec20ae970821387224ec3bb15b4bd22b89bfb86dc4d2b3827aa3d99e48d.
//
// Solidity: event LimitPerBlockUpdated(uint128 arg0)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseLimitPerBlockUpdated(log types.Log) (*CreditConfiguratorLimitPerBlockUpdated, error) {
	event := new(CreditConfiguratorLimitPerBlockUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "LimitPerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorLimitsUpdatedIterator is returned from FilterLimitsUpdated and is used to iterate over the raw logs and unpacked data for LimitsUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorLimitsUpdatedIterator struct {
	Event *CreditConfiguratorLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorLimitsUpdated)
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
		it.Event = new(CreditConfiguratorLimitsUpdated)
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
func (it *CreditConfiguratorLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorLimitsUpdated represents a LimitsUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorLimitsUpdated struct {
	MinBorrowedAmount *big.Int
	MaxBorrowedAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLimitsUpdated is a free log retrieval operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterLimitsUpdated(opts *bind.FilterOpts) (*CreditConfiguratorLimitsUpdatedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "LimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorLimitsUpdatedIterator{contract: _CreditConfigurator.contract, event: "LimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchLimitsUpdated is a free log subscription operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchLimitsUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "LimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorLimitsUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "LimitsUpdated", log); err != nil {
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

// ParseLimitsUpdated is a log parse operation binding the contract event 0x4d4981437d0211f9e6843eb024d9ada1fa3a99514d4343d4aece106dd11524bb.
//
// Solidity: event LimitsUpdated(uint256 minBorrowedAmount, uint256 maxBorrowedAmount)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseLimitsUpdated(log types.Log) (*CreditConfiguratorLimitsUpdated, error) {
	event := new(CreditConfiguratorLimitsUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "LimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditConfigurator contract.
type CreditConfiguratorPausedIterator struct {
	Event *CreditConfiguratorPaused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorPaused)
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
		it.Event = new(CreditConfiguratorPaused)
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
func (it *CreditConfiguratorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorPaused represents a Paused event raised by the CreditConfigurator contract.
type CreditConfiguratorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditConfiguratorPausedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorPausedIterator{contract: _CreditConfigurator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorPaused) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorPaused)
				if err := _CreditConfigurator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditConfigurator *CreditConfiguratorFilterer) ParsePaused(log types.Log) (*CreditConfiguratorPaused, error) {
	event := new(CreditConfiguratorPaused)
	if err := _CreditConfigurator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorPriceOracleUpgradedIterator is returned from FilterPriceOracleUpgraded and is used to iterate over the raw logs and unpacked data for PriceOracleUpgraded events raised by the CreditConfigurator contract.
type CreditConfiguratorPriceOracleUpgradedIterator struct {
	Event *CreditConfiguratorPriceOracleUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorPriceOracleUpgraded)
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
		it.Event = new(CreditConfiguratorPriceOracleUpgraded)
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
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorPriceOracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorPriceOracleUpgraded represents a PriceOracleUpgraded event raised by the CreditConfigurator contract.
type CreditConfiguratorPriceOracleUpgraded struct {
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpgraded is a free log retrieval operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterPriceOracleUpgraded(opts *bind.FilterOpts, newPriceOracle []common.Address) (*CreditConfiguratorPriceOracleUpgradedIterator, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "PriceOracleUpgraded", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorPriceOracleUpgradedIterator{contract: _CreditConfigurator.contract, event: "PriceOracleUpgraded", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpgraded is a free log subscription operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchPriceOracleUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorPriceOracleUpgraded, newPriceOracle []common.Address) (event.Subscription, error) {

	var newPriceOracleRule []interface{}
	for _, newPriceOracleItem := range newPriceOracle {
		newPriceOracleRule = append(newPriceOracleRule, newPriceOracleItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "PriceOracleUpgraded", newPriceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorPriceOracleUpgraded)
				if err := _CreditConfigurator.contract.UnpackLog(event, "PriceOracleUpgraded", log); err != nil {
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

// ParsePriceOracleUpgraded is a log parse operation binding the contract event 0x3f82447be465b0b5a4a9e54c74d5f6ae73f2e9537f2cc1590a340524703d0961.
//
// Solidity: event PriceOracleUpgraded(address indexed newPriceOracle)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParsePriceOracleUpgraded(log types.Log) (*CreditConfiguratorPriceOracleUpgraded, error) {
	event := new(CreditConfiguratorPriceOracleUpgraded)
	if err := _CreditConfigurator.contract.UnpackLog(event, "PriceOracleUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorTokenAllowedIterator is returned from FilterTokenAllowed and is used to iterate over the raw logs and unpacked data for TokenAllowed events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenAllowedIterator struct {
	Event *CreditConfiguratorTokenAllowed // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenAllowed)
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
		it.Event = new(CreditConfiguratorTokenAllowed)
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
func (it *CreditConfiguratorTokenAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenAllowed represents a TokenAllowed event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenAllowed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAllowed is a free log retrieval operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenAllowed(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenAllowedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenAllowedIterator{contract: _CreditConfigurator.contract, event: "TokenAllowed", logs: logs, sub: sub}, nil
}

// WatchTokenAllowed is a free log subscription operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenAllowed(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenAllowed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenAllowed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenAllowed)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
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

// ParseTokenAllowed is a log parse operation binding the contract event 0xbeceb48aeaa805aeae57be163cca6249077a18734e408a85aa74e875c4373809.
//
// Solidity: event TokenAllowed(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenAllowed(log types.Log) (*CreditConfiguratorTokenAllowed, error) {
	event := new(CreditConfiguratorTokenAllowed)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorTokenForbiddenIterator is returned from FilterTokenForbidden and is used to iterate over the raw logs and unpacked data for TokenForbidden events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenForbiddenIterator struct {
	Event *CreditConfiguratorTokenForbidden // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenForbiddenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenForbidden)
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
		it.Event = new(CreditConfiguratorTokenForbidden)
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
func (it *CreditConfiguratorTokenForbiddenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenForbiddenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenForbidden represents a TokenForbidden event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenForbidden struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenForbidden is a free log retrieval operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenForbidden(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenForbiddenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenForbiddenIterator{contract: _CreditConfigurator.contract, event: "TokenForbidden", logs: logs, sub: sub}, nil
}

// WatchTokenForbidden is a free log subscription operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenForbidden(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenForbidden, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenForbidden", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenForbidden)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
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

// ParseTokenForbidden is a log parse operation binding the contract event 0xf17b849746e74d7186170c9553d4bbf60b4f8bb1ed81fe50c099b934fb078f05.
//
// Solidity: event TokenForbidden(address indexed token)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenForbidden(log types.Log) (*CreditConfiguratorTokenForbidden, error) {
	event := new(CreditConfiguratorTokenForbidden)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenForbidden", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorTokenLiquidationThresholdUpdatedIterator is returned from FilterTokenLiquidationThresholdUpdated and is used to iterate over the raw logs and unpacked data for TokenLiquidationThresholdUpdated events raised by the CreditConfigurator contract.
type CreditConfiguratorTokenLiquidationThresholdUpdatedIterator struct {
	Event *CreditConfiguratorTokenLiquidationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorTokenLiquidationThresholdUpdated)
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
		it.Event = new(CreditConfiguratorTokenLiquidationThresholdUpdated)
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
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorTokenLiquidationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorTokenLiquidationThresholdUpdated represents a TokenLiquidationThresholdUpdated event raised by the CreditConfigurator contract.
type CreditConfiguratorTokenLiquidationThresholdUpdated struct {
	Token              common.Address
	LiquidityThreshold uint16
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenLiquidationThresholdUpdated is a free log retrieval operation binding the contract event 0x94525be3f877fb3bb3d260299de56a31f86aaacb06ccba90e4fbedcb693ac650.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint16 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterTokenLiquidationThresholdUpdated(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorTokenLiquidationThresholdUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "TokenLiquidationThresholdUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorTokenLiquidationThresholdUpdatedIterator{contract: _CreditConfigurator.contract, event: "TokenLiquidationThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenLiquidationThresholdUpdated is a free log subscription operation binding the contract event 0x94525be3f877fb3bb3d260299de56a31f86aaacb06ccba90e4fbedcb693ac650.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint16 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchTokenLiquidationThresholdUpdated(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorTokenLiquidationThresholdUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "TokenLiquidationThresholdUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorTokenLiquidationThresholdUpdated)
				if err := _CreditConfigurator.contract.UnpackLog(event, "TokenLiquidationThresholdUpdated", log); err != nil {
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

// ParseTokenLiquidationThresholdUpdated is a log parse operation binding the contract event 0x94525be3f877fb3bb3d260299de56a31f86aaacb06ccba90e4fbedcb693ac650.
//
// Solidity: event TokenLiquidationThresholdUpdated(address indexed token, uint16 liquidityThreshold)
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseTokenLiquidationThresholdUpdated(log types.Log) (*CreditConfiguratorTokenLiquidationThresholdUpdated, error) {
	event := new(CreditConfiguratorTokenLiquidationThresholdUpdated)
	if err := _CreditConfigurator.contract.UnpackLog(event, "TokenLiquidationThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditConfigurator contract.
type CreditConfiguratorUnpausedIterator struct {
	Event *CreditConfiguratorUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorUnpaused)
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
		it.Event = new(CreditConfiguratorUnpaused)
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
func (it *CreditConfiguratorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorUnpaused represents a Unpaused event raised by the CreditConfigurator contract.
type CreditConfiguratorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditConfiguratorUnpausedIterator, error) {

	logs, sub, err := _CreditConfigurator.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorUnpausedIterator{contract: _CreditConfigurator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfigurator *CreditConfiguratorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorUnpaused)
				if err := _CreditConfigurator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditConfigurator *CreditConfiguratorFilterer) ParseUnpaused(log types.Log) (*CreditConfiguratorUnpaused, error) {
	event := new(CreditConfiguratorUnpaused)
	if err := _CreditConfigurator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
