// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditConfiguratorv3

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

// CollateralToken is an auto generated low-level Go binding around an user-defined struct.
type CollateralToken struct {
	Token                common.Address
	LiquidationThreshold uint16
}

// CreditManagerOpts is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerOpts struct {
	MinDebt          *big.Int
	MaxDebt          *big.Int
	CollateralTokens []CollateralToken
	DegenNFT         common.Address
	Expirable        bool
	Name             string
}

// CreditConfiguratorv3MetaData contains all meta data concerning the CreditConfiguratorv3 contract.
var CreditConfiguratorv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCreditManagerV3\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"contractCreditFacadeV3\",\"name\":\"_creditFacade\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"minDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxDebt\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structCreditManagerOpts\",\"name\":\"opts\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdapterIsNotRegisteredException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotControllerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnpausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectExpirationDateException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectLimitsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectLiquidationThresholdException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenContractException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetContractNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsNotQuotedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AddCollateralToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"AddEmergencyLiquidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"AllowAdapter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AllowToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"}],\"name\":\"CreditConfiguratorUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"ForbidAdapter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ForbidToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"QuoteToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"RemoveEmergencyLiquidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ResetCumulativeLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationThresholdInitial\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationThresholdFinal\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestampRampStart\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestampRampEnd\",\"type\":\"uint40\"}],\"name\":\"ScheduleTokenLiquidationThresholdRamp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"}],\"name\":\"SetBorrowingLimits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"}],\"name\":\"SetBotList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"}],\"name\":\"SetCreditFacade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"}],\"name\":\"SetExpirationDate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"maxCumulativeLoss\",\"type\":\"uint128\"}],\"name\":\"SetMaxCumulativeLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"}],\"name\":\"SetMaxDebtPerBlockMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"}],\"name\":\"SetMaxEnabledTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"SetPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"SetTokenLiquidationThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationPremium\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationPremiumExpired\",\"type\":\"uint16\"}],\"name\":\"UpdateFees\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"addCollateralToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"addEmergencyLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"allowAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowedAdapters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyLiquidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"forbidAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forbidBorrowing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"forbidToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"makeTokenQuoted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThresholdFinal\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"rampStart\",\"type\":\"uint40\"},{\"internalType\":\"uint24\",\"name\":\"rampDuration\",\"type\":\"uint24\"}],\"name\":\"rampLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"removeEmergencyLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetCumulativeLoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"setBotList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCreditFacade\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"migrateParams\",\"type\":\"bool\"}],\"name\":\"setCreditFacade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"newExpirationDate\",\"type\":\"uint40\"}],\"name\":\"setExpirationDate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationPremium\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationPremiumExpired\",\"type\":\"uint16\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"setLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newMaxCumulativeLoss\",\"type\":\"uint128\"}],\"name\":\"setMaxCumulativeLoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxDebt\",\"type\":\"uint128\"}],\"name\":\"setMaxDebtLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newMaxDebtLimitPerBlockMultiplier\",\"type\":\"uint8\"}],\"name\":\"setMaxDebtPerBlockMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newMaxEnabledTokens\",\"type\":\"uint8\"}],\"name\":\"setMaxEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"minDebt\",\"type\":\"uint128\"}],\"name\":\"setMinDebtLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVersion\",\"type\":\"uint256\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCreditConfigurator\",\"type\":\"address\"}],\"name\":\"upgradeCreditConfigurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditConfiguratorv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditConfiguratorv3MetaData.ABI instead.
var CreditConfiguratorv3ABI = CreditConfiguratorv3MetaData.ABI

// CreditConfiguratorv3 is an auto generated Go binding around an Ethereum contract.
type CreditConfiguratorv3 struct {
	CreditConfiguratorv3Caller     // Read-only binding to the contract
	CreditConfiguratorv3Transactor // Write-only binding to the contract
	CreditConfiguratorv3Filterer   // Log filterer for contract events
}

// CreditConfiguratorv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditConfiguratorv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditConfiguratorv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditConfiguratorv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfiguratorv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditConfiguratorv3Session struct {
	Contract     *CreditConfiguratorv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CreditConfiguratorv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditConfiguratorv3CallerSession struct {
	Contract *CreditConfiguratorv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CreditConfiguratorv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditConfiguratorv3TransactorSession struct {
	Contract     *CreditConfiguratorv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CreditConfiguratorv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditConfiguratorv3Raw struct {
	Contract *CreditConfiguratorv3 // Generic contract binding to access the raw methods on
}

// CreditConfiguratorv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditConfiguratorv3CallerRaw struct {
	Contract *CreditConfiguratorv3Caller // Generic read-only contract binding to access the raw methods on
}

// CreditConfiguratorv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditConfiguratorv3TransactorRaw struct {
	Contract *CreditConfiguratorv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditConfiguratorv3 creates a new instance of CreditConfiguratorv3, bound to a specific deployed contract.
func NewCreditConfiguratorv3(address common.Address, backend bind.ContractBackend) (*CreditConfiguratorv3, error) {
	contract, err := bindCreditConfiguratorv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3{CreditConfiguratorv3Caller: CreditConfiguratorv3Caller{contract: contract}, CreditConfiguratorv3Transactor: CreditConfiguratorv3Transactor{contract: contract}, CreditConfiguratorv3Filterer: CreditConfiguratorv3Filterer{contract: contract}}, nil
}

// NewCreditConfiguratorv3Caller creates a new read-only instance of CreditConfiguratorv3, bound to a specific deployed contract.
func NewCreditConfiguratorv3Caller(address common.Address, caller bind.ContractCaller) (*CreditConfiguratorv3Caller, error) {
	contract, err := bindCreditConfiguratorv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3Caller{contract: contract}, nil
}

// NewCreditConfiguratorv3Transactor creates a new write-only instance of CreditConfiguratorv3, bound to a specific deployed contract.
func NewCreditConfiguratorv3Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditConfiguratorv3Transactor, error) {
	contract, err := bindCreditConfiguratorv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3Transactor{contract: contract}, nil
}

// NewCreditConfiguratorv3Filterer creates a new log filterer instance of CreditConfiguratorv3, bound to a specific deployed contract.
func NewCreditConfiguratorv3Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditConfiguratorv3Filterer, error) {
	contract, err := bindCreditConfiguratorv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3Filterer{contract: contract}, nil
}

// bindCreditConfiguratorv3 binds a generic wrapper to an already deployed contract.
func bindCreditConfiguratorv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditConfiguratorv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfiguratorv3 *CreditConfiguratorv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfiguratorv3.Contract.CreditConfiguratorv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfiguratorv3 *CreditConfiguratorv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.CreditConfiguratorv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfiguratorv3 *CreditConfiguratorv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.CreditConfiguratorv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfiguratorv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Acl() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Acl(&_CreditConfiguratorv3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) Acl() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Acl(&_CreditConfiguratorv3.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AddressProvider() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.AddressProvider(&_CreditConfiguratorv3.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) AddressProvider() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.AddressProvider(&_CreditConfiguratorv3.CallOpts)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) AllowedAdapters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "allowedAdapters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AllowedAdapters() ([]common.Address, error) {
	return _CreditConfiguratorv3.Contract.AllowedAdapters(&_CreditConfiguratorv3.CallOpts)
}

// AllowedAdapters is a free data retrieval call binding the contract method 0x1c42130e.
//
// Solidity: function allowedAdapters() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) AllowedAdapters() ([]common.Address, error) {
	return _CreditConfiguratorv3.Contract.AllowedAdapters(&_CreditConfiguratorv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Controller() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Controller(&_CreditConfiguratorv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) Controller() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Controller(&_CreditConfiguratorv3.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) CreditFacade() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.CreditFacade(&_CreditConfiguratorv3.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) CreditFacade() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.CreditFacade(&_CreditConfiguratorv3.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) CreditManager() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.CreditManager(&_CreditConfiguratorv3.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) CreditManager() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.CreditManager(&_CreditConfiguratorv3.CallOpts)
}

// EmergencyLiquidators is a free data retrieval call binding the contract method 0xc752d247.
//
// Solidity: function emergencyLiquidators() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) EmergencyLiquidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "emergencyLiquidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EmergencyLiquidators is a free data retrieval call binding the contract method 0xc752d247.
//
// Solidity: function emergencyLiquidators() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) EmergencyLiquidators() ([]common.Address, error) {
	return _CreditConfiguratorv3.Contract.EmergencyLiquidators(&_CreditConfiguratorv3.CallOpts)
}

// EmergencyLiquidators is a free data retrieval call binding the contract method 0xc752d247.
//
// Solidity: function emergencyLiquidators() view returns(address[])
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) EmergencyLiquidators() ([]common.Address, error) {
	return _CreditConfiguratorv3.Contract.EmergencyLiquidators(&_CreditConfiguratorv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Paused() (bool, error) {
	return _CreditConfiguratorv3.Contract.Paused(&_CreditConfiguratorv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) Paused() (bool, error) {
	return _CreditConfiguratorv3.Contract.Paused(&_CreditConfiguratorv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Underlying() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Underlying(&_CreditConfiguratorv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) Underlying() (common.Address, error) {
	return _CreditConfiguratorv3.Contract.Underlying(&_CreditConfiguratorv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditConfiguratorv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Version() (*big.Int, error) {
	return _CreditConfiguratorv3.Contract.Version(&_CreditConfiguratorv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditConfiguratorv3 *CreditConfiguratorv3CallerSession) Version() (*big.Int, error) {
	return _CreditConfiguratorv3.Contract.Version(&_CreditConfiguratorv3.CallOpts)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) AddCollateralToken(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "addCollateralToken", token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AddCollateralToken(&_CreditConfiguratorv3.TransactOpts, token, liquidationThreshold)
}

// AddCollateralToken is a paid mutator transaction binding the contract method 0x3e7c88d6.
//
// Solidity: function addCollateralToken(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) AddCollateralToken(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AddCollateralToken(&_CreditConfiguratorv3.TransactOpts, token, liquidationThreshold)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) AddEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "addEmergencyLiquidator", liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AddEmergencyLiquidator(&_CreditConfiguratorv3.TransactOpts, liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AddEmergencyLiquidator(&_CreditConfiguratorv3.TransactOpts, liquidator)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) AllowAdapter(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "allowAdapter", adapter)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AllowAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AllowAdapter(&_CreditConfiguratorv3.TransactOpts, adapter)
}

// AllowAdapter is a paid mutator transaction binding the contract method 0xeffa5d6e.
//
// Solidity: function allowAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) AllowAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AllowAdapter(&_CreditConfiguratorv3.TransactOpts, adapter)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) AllowToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "allowToken", token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AllowToken(&_CreditConfiguratorv3.TransactOpts, token)
}

// AllowToken is a paid mutator transaction binding the contract method 0xb53472ef.
//
// Solidity: function allowToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) AllowToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.AllowToken(&_CreditConfiguratorv3.TransactOpts, token)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) ForbidAdapter(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "forbidAdapter", adapter)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) ForbidAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidAdapter(&_CreditConfiguratorv3.TransactOpts, adapter)
}

// ForbidAdapter is a paid mutator transaction binding the contract method 0x1495c7d2.
//
// Solidity: function forbidAdapter(address adapter) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) ForbidAdapter(adapter common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidAdapter(&_CreditConfiguratorv3.TransactOpts, adapter)
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) ForbidBorrowing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "forbidBorrowing")
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) ForbidBorrowing() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidBorrowing(&_CreditConfiguratorv3.TransactOpts)
}

// ForbidBorrowing is a paid mutator transaction binding the contract method 0xbee1babf.
//
// Solidity: function forbidBorrowing() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) ForbidBorrowing() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidBorrowing(&_CreditConfiguratorv3.TransactOpts)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) ForbidToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "forbidToken", token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidToken(&_CreditConfiguratorv3.TransactOpts, token)
}

// ForbidToken is a paid mutator transaction binding the contract method 0x24147708.
//
// Solidity: function forbidToken(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) ForbidToken(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ForbidToken(&_CreditConfiguratorv3.TransactOpts, token)
}

// MakeTokenQuoted is a paid mutator transaction binding the contract method 0x4e48e9c7.
//
// Solidity: function makeTokenQuoted(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) MakeTokenQuoted(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "makeTokenQuoted", token)
}

// MakeTokenQuoted is a paid mutator transaction binding the contract method 0x4e48e9c7.
//
// Solidity: function makeTokenQuoted(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) MakeTokenQuoted(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.MakeTokenQuoted(&_CreditConfiguratorv3.TransactOpts, token)
}

// MakeTokenQuoted is a paid mutator transaction binding the contract method 0x4e48e9c7.
//
// Solidity: function makeTokenQuoted(address token) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) MakeTokenQuoted(token common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.MakeTokenQuoted(&_CreditConfiguratorv3.TransactOpts, token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Pause() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.Pause(&_CreditConfiguratorv3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) Pause() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.Pause(&_CreditConfiguratorv3.TransactOpts)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) RampLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "rampLiquidationThreshold", token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) RampLiquidationThreshold(token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.RampLiquidationThreshold(&_CreditConfiguratorv3.TransactOpts, token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RampLiquidationThreshold is a paid mutator transaction binding the contract method 0x3d2ff001.
//
// Solidity: function rampLiquidationThreshold(address token, uint16 liquidationThresholdFinal, uint40 rampStart, uint24 rampDuration) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) RampLiquidationThreshold(token common.Address, liquidationThresholdFinal uint16, rampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.RampLiquidationThreshold(&_CreditConfiguratorv3.TransactOpts, token, liquidationThresholdFinal, rampStart, rampDuration)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) RemoveEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "removeEmergencyLiquidator", liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.RemoveEmergencyLiquidator(&_CreditConfiguratorv3.TransactOpts, liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.RemoveEmergencyLiquidator(&_CreditConfiguratorv3.TransactOpts, liquidator)
}

// ResetCumulativeLoss is a paid mutator transaction binding the contract method 0xa706efc4.
//
// Solidity: function resetCumulativeLoss() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) ResetCumulativeLoss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "resetCumulativeLoss")
}

// ResetCumulativeLoss is a paid mutator transaction binding the contract method 0xa706efc4.
//
// Solidity: function resetCumulativeLoss() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) ResetCumulativeLoss() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ResetCumulativeLoss(&_CreditConfiguratorv3.TransactOpts)
}

// ResetCumulativeLoss is a paid mutator transaction binding the contract method 0xa706efc4.
//
// Solidity: function resetCumulativeLoss() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) ResetCumulativeLoss() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.ResetCumulativeLoss(&_CreditConfiguratorv3.TransactOpts)
}

// SetBotList is a paid mutator transaction binding the contract method 0xc1b9366f.
//
// Solidity: function setBotList(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetBotList(opts *bind.TransactOpts, newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setBotList", newVersion)
}

// SetBotList is a paid mutator transaction binding the contract method 0xc1b9366f.
//
// Solidity: function setBotList(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetBotList(newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetBotList(&_CreditConfiguratorv3.TransactOpts, newVersion)
}

// SetBotList is a paid mutator transaction binding the contract method 0xc1b9366f.
//
// Solidity: function setBotList(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetBotList(newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetBotList(&_CreditConfiguratorv3.TransactOpts, newVersion)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetController(&_CreditConfiguratorv3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetController(&_CreditConfiguratorv3.TransactOpts, newController)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetCreditFacade(opts *bind.TransactOpts, newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setCreditFacade", newCreditFacade, migrateParams)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetCreditFacade(newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetCreditFacade(&_CreditConfiguratorv3.TransactOpts, newCreditFacade, migrateParams)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0x28afc97c.
//
// Solidity: function setCreditFacade(address newCreditFacade, bool migrateParams) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetCreditFacade(newCreditFacade common.Address, migrateParams bool) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetCreditFacade(&_CreditConfiguratorv3.TransactOpts, newCreditFacade, migrateParams)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetExpirationDate(opts *bind.TransactOpts, newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setExpirationDate", newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetExpirationDate(&_CreditConfiguratorv3.TransactOpts, newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetExpirationDate(&_CreditConfiguratorv3.TransactOpts, newExpirationDate)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetFees(opts *bind.TransactOpts, feeInterest uint16, feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setFees", feeInterest, feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetFees(feeInterest uint16, feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetFees(&_CreditConfiguratorv3.TransactOpts, feeInterest, feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetFees(feeInterest uint16, feeLiquidation uint16, liquidationPremium uint16, feeLiquidationExpired uint16, liquidationPremiumExpired uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetFees(&_CreditConfiguratorv3.TransactOpts, feeInterest, feeLiquidation, liquidationPremium, feeLiquidationExpired, liquidationPremiumExpired)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetLiquidationThreshold(&_CreditConfiguratorv3.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetLiquidationThreshold(&_CreditConfiguratorv3.TransactOpts, token, liquidationThreshold)
}

// SetMaxCumulativeLoss is a paid mutator transaction binding the contract method 0x08c25f8f.
//
// Solidity: function setMaxCumulativeLoss(uint128 newMaxCumulativeLoss) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetMaxCumulativeLoss(opts *bind.TransactOpts, newMaxCumulativeLoss *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setMaxCumulativeLoss", newMaxCumulativeLoss)
}

// SetMaxCumulativeLoss is a paid mutator transaction binding the contract method 0x08c25f8f.
//
// Solidity: function setMaxCumulativeLoss(uint128 newMaxCumulativeLoss) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetMaxCumulativeLoss(newMaxCumulativeLoss *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxCumulativeLoss(&_CreditConfiguratorv3.TransactOpts, newMaxCumulativeLoss)
}

// SetMaxCumulativeLoss is a paid mutator transaction binding the contract method 0x08c25f8f.
//
// Solidity: function setMaxCumulativeLoss(uint128 newMaxCumulativeLoss) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetMaxCumulativeLoss(newMaxCumulativeLoss *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxCumulativeLoss(&_CreditConfiguratorv3.TransactOpts, newMaxCumulativeLoss)
}

// SetMaxDebtLimit is a paid mutator transaction binding the contract method 0x8c83d0dd.
//
// Solidity: function setMaxDebtLimit(uint128 maxDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetMaxDebtLimit(opts *bind.TransactOpts, maxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setMaxDebtLimit", maxDebt)
}

// SetMaxDebtLimit is a paid mutator transaction binding the contract method 0x8c83d0dd.
//
// Solidity: function setMaxDebtLimit(uint128 maxDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetMaxDebtLimit(maxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxDebtLimit(&_CreditConfiguratorv3.TransactOpts, maxDebt)
}

// SetMaxDebtLimit is a paid mutator transaction binding the contract method 0x8c83d0dd.
//
// Solidity: function setMaxDebtLimit(uint128 maxDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetMaxDebtLimit(maxDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxDebtLimit(&_CreditConfiguratorv3.TransactOpts, maxDebt)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetMaxDebtPerBlockMultiplier(opts *bind.TransactOpts, newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setMaxDebtPerBlockMultiplier", newMaxDebtLimitPerBlockMultiplier)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetMaxDebtPerBlockMultiplier(newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxDebtPerBlockMultiplier(&_CreditConfiguratorv3.TransactOpts, newMaxDebtLimitPerBlockMultiplier)
}

// SetMaxDebtPerBlockMultiplier is a paid mutator transaction binding the contract method 0xb954d809.
//
// Solidity: function setMaxDebtPerBlockMultiplier(uint8 newMaxDebtLimitPerBlockMultiplier) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetMaxDebtPerBlockMultiplier(newMaxDebtLimitPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxDebtPerBlockMultiplier(&_CreditConfiguratorv3.TransactOpts, newMaxDebtLimitPerBlockMultiplier)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetMaxEnabledTokens(opts *bind.TransactOpts, newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setMaxEnabledTokens", newMaxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetMaxEnabledTokens(newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxEnabledTokens(&_CreditConfiguratorv3.TransactOpts, newMaxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetMaxEnabledTokens(newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMaxEnabledTokens(&_CreditConfiguratorv3.TransactOpts, newMaxEnabledTokens)
}

// SetMinDebtLimit is a paid mutator transaction binding the contract method 0x98acdd77.
//
// Solidity: function setMinDebtLimit(uint128 minDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetMinDebtLimit(opts *bind.TransactOpts, minDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setMinDebtLimit", minDebt)
}

// SetMinDebtLimit is a paid mutator transaction binding the contract method 0x98acdd77.
//
// Solidity: function setMinDebtLimit(uint128 minDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetMinDebtLimit(minDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMinDebtLimit(&_CreditConfiguratorv3.TransactOpts, minDebt)
}

// SetMinDebtLimit is a paid mutator transaction binding the contract method 0x98acdd77.
//
// Solidity: function setMinDebtLimit(uint128 minDebt) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetMinDebtLimit(minDebt *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetMinDebtLimit(&_CreditConfiguratorv3.TransactOpts, minDebt)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x625c513b.
//
// Solidity: function setPriceOracle(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) SetPriceOracle(opts *bind.TransactOpts, newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "setPriceOracle", newVersion)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x625c513b.
//
// Solidity: function setPriceOracle(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) SetPriceOracle(newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetPriceOracle(&_CreditConfiguratorv3.TransactOpts, newVersion)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x625c513b.
//
// Solidity: function setPriceOracle(uint256 newVersion) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) SetPriceOracle(newVersion *big.Int) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.SetPriceOracle(&_CreditConfiguratorv3.TransactOpts, newVersion)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) Unpause() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.Unpause(&_CreditConfiguratorv3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.Unpause(&_CreditConfiguratorv3.TransactOpts)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Transactor) UpgradeCreditConfigurator(opts *bind.TransactOpts, newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.contract.Transact(opts, "upgradeCreditConfigurator", newCreditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Session) UpgradeCreditConfigurator(newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.UpgradeCreditConfigurator(&_CreditConfiguratorv3.TransactOpts, newCreditConfigurator)
}

// UpgradeCreditConfigurator is a paid mutator transaction binding the contract method 0x456e0742.
//
// Solidity: function upgradeCreditConfigurator(address newCreditConfigurator) returns()
func (_CreditConfiguratorv3 *CreditConfiguratorv3TransactorSession) UpgradeCreditConfigurator(newCreditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditConfiguratorv3.Contract.UpgradeCreditConfigurator(&_CreditConfiguratorv3.TransactOpts, newCreditConfigurator)
}

// CreditConfiguratorv3AddCollateralTokenIterator is returned from FilterAddCollateralToken and is used to iterate over the raw logs and unpacked data for AddCollateralToken events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AddCollateralTokenIterator struct {
	Event *CreditConfiguratorv3AddCollateralToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3AddCollateralTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3AddCollateralToken)
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
		it.Event = new(CreditConfiguratorv3AddCollateralToken)
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
func (it *CreditConfiguratorv3AddCollateralTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3AddCollateralTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3AddCollateralToken represents a AddCollateralToken event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AddCollateralToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddCollateralToken is a free log retrieval operation binding the contract event 0x7c3f95f8569977586927f95930461a261e2121e326fcb513242f9e5c8b8ea6dc.
//
// Solidity: event AddCollateralToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterAddCollateralToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3AddCollateralTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "AddCollateralToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3AddCollateralTokenIterator{contract: _CreditConfiguratorv3.contract, event: "AddCollateralToken", logs: logs, sub: sub}, nil
}

// WatchAddCollateralToken is a free log subscription operation binding the contract event 0x7c3f95f8569977586927f95930461a261e2121e326fcb513242f9e5c8b8ea6dc.
//
// Solidity: event AddCollateralToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchAddCollateralToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3AddCollateralToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "AddCollateralToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3AddCollateralToken)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AddCollateralToken", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseAddCollateralToken(log types.Log) (*CreditConfiguratorv3AddCollateralToken, error) {
	event := new(CreditConfiguratorv3AddCollateralToken)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AddCollateralToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3AddEmergencyLiquidatorIterator is returned from FilterAddEmergencyLiquidator and is used to iterate over the raw logs and unpacked data for AddEmergencyLiquidator events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AddEmergencyLiquidatorIterator struct {
	Event *CreditConfiguratorv3AddEmergencyLiquidator // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3AddEmergencyLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3AddEmergencyLiquidator)
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
		it.Event = new(CreditConfiguratorv3AddEmergencyLiquidator)
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
func (it *CreditConfiguratorv3AddEmergencyLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3AddEmergencyLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3AddEmergencyLiquidator represents a AddEmergencyLiquidator event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AddEmergencyLiquidator struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddEmergencyLiquidator is a free log retrieval operation binding the contract event 0x35b5318c4163fcef2999d30de8d1af689327f68fa51a148804fa6ed8f5f40ff4.
//
// Solidity: event AddEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterAddEmergencyLiquidator(opts *bind.FilterOpts, liquidator []common.Address) (*CreditConfiguratorv3AddEmergencyLiquidatorIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "AddEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3AddEmergencyLiquidatorIterator{contract: _CreditConfiguratorv3.contract, event: "AddEmergencyLiquidator", logs: logs, sub: sub}, nil
}

// WatchAddEmergencyLiquidator is a free log subscription operation binding the contract event 0x35b5318c4163fcef2999d30de8d1af689327f68fa51a148804fa6ed8f5f40ff4.
//
// Solidity: event AddEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchAddEmergencyLiquidator(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3AddEmergencyLiquidator, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "AddEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3AddEmergencyLiquidator)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AddEmergencyLiquidator", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseAddEmergencyLiquidator(log types.Log) (*CreditConfiguratorv3AddEmergencyLiquidator, error) {
	event := new(CreditConfiguratorv3AddEmergencyLiquidator)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AddEmergencyLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3AllowAdapterIterator is returned from FilterAllowAdapter and is used to iterate over the raw logs and unpacked data for AllowAdapter events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AllowAdapterIterator struct {
	Event *CreditConfiguratorv3AllowAdapter // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3AllowAdapterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3AllowAdapter)
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
		it.Event = new(CreditConfiguratorv3AllowAdapter)
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
func (it *CreditConfiguratorv3AllowAdapterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3AllowAdapterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3AllowAdapter represents a AllowAdapter event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AllowAdapter struct {
	TargetContract common.Address
	Adapter        common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllowAdapter is a free log retrieval operation binding the contract event 0x0bc09e53304ef58ff3ff8295411d9171c75ee4af48277db5fc605ab12e056bee.
//
// Solidity: event AllowAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterAllowAdapter(opts *bind.FilterOpts, targetContract []common.Address, adapter []common.Address) (*CreditConfiguratorv3AllowAdapterIterator, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "AllowAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3AllowAdapterIterator{contract: _CreditConfiguratorv3.contract, event: "AllowAdapter", logs: logs, sub: sub}, nil
}

// WatchAllowAdapter is a free log subscription operation binding the contract event 0x0bc09e53304ef58ff3ff8295411d9171c75ee4af48277db5fc605ab12e056bee.
//
// Solidity: event AllowAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchAllowAdapter(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3AllowAdapter, targetContract []common.Address, adapter []common.Address) (event.Subscription, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "AllowAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3AllowAdapter)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AllowAdapter", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseAllowAdapter(log types.Log) (*CreditConfiguratorv3AllowAdapter, error) {
	event := new(CreditConfiguratorv3AllowAdapter)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AllowAdapter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3AllowTokenIterator is returned from FilterAllowToken and is used to iterate over the raw logs and unpacked data for AllowToken events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AllowTokenIterator struct {
	Event *CreditConfiguratorv3AllowToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3AllowTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3AllowToken)
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
		it.Event = new(CreditConfiguratorv3AllowToken)
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
func (it *CreditConfiguratorv3AllowTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3AllowTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3AllowToken represents a AllowToken event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3AllowToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAllowToken is a free log retrieval operation binding the contract event 0x14009112f2dcb15cad32dab6bf972d6d85286e4ae1178f27323ffe25359459e6.
//
// Solidity: event AllowToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterAllowToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3AllowTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "AllowToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3AllowTokenIterator{contract: _CreditConfiguratorv3.contract, event: "AllowToken", logs: logs, sub: sub}, nil
}

// WatchAllowToken is a free log subscription operation binding the contract event 0x14009112f2dcb15cad32dab6bf972d6d85286e4ae1178f27323ffe25359459e6.
//
// Solidity: event AllowToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchAllowToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3AllowToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "AllowToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3AllowToken)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AllowToken", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseAllowToken(log types.Log) (*CreditConfiguratorv3AllowToken, error) {
	event := new(CreditConfiguratorv3AllowToken)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "AllowToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3CreditConfiguratorUpgradedIterator is returned from FilterCreditConfiguratorUpgraded and is used to iterate over the raw logs and unpacked data for CreditConfiguratorUpgraded events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3CreditConfiguratorUpgradedIterator struct {
	Event *CreditConfiguratorv3CreditConfiguratorUpgraded // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3CreditConfiguratorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3CreditConfiguratorUpgraded)
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
		it.Event = new(CreditConfiguratorv3CreditConfiguratorUpgraded)
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
func (it *CreditConfiguratorv3CreditConfiguratorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3CreditConfiguratorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3CreditConfiguratorUpgraded represents a CreditConfiguratorUpgraded event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3CreditConfiguratorUpgraded struct {
	CreditConfigurator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreditConfiguratorUpgraded is a free log retrieval operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed creditConfigurator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterCreditConfiguratorUpgraded(opts *bind.FilterOpts, creditConfigurator []common.Address) (*CreditConfiguratorv3CreditConfiguratorUpgradedIterator, error) {

	var creditConfiguratorRule []interface{}
	for _, creditConfiguratorItem := range creditConfigurator {
		creditConfiguratorRule = append(creditConfiguratorRule, creditConfiguratorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "CreditConfiguratorUpgraded", creditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3CreditConfiguratorUpgradedIterator{contract: _CreditConfiguratorv3.contract, event: "CreditConfiguratorUpgraded", logs: logs, sub: sub}, nil
}

// WatchCreditConfiguratorUpgraded is a free log subscription operation binding the contract event 0x5a0b7d0f9c24b39256e112a0584b4c5ce38d8f1dee2e7c56f15b852604cdc886.
//
// Solidity: event CreditConfiguratorUpgraded(address indexed creditConfigurator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchCreditConfiguratorUpgraded(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3CreditConfiguratorUpgraded, creditConfigurator []common.Address) (event.Subscription, error) {

	var creditConfiguratorRule []interface{}
	for _, creditConfiguratorItem := range creditConfigurator {
		creditConfiguratorRule = append(creditConfiguratorRule, creditConfiguratorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "CreditConfiguratorUpgraded", creditConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3CreditConfiguratorUpgraded)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseCreditConfiguratorUpgraded(log types.Log) (*CreditConfiguratorv3CreditConfiguratorUpgraded, error) {
	event := new(CreditConfiguratorv3CreditConfiguratorUpgraded)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "CreditConfiguratorUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3ForbidAdapterIterator is returned from FilterForbidAdapter and is used to iterate over the raw logs and unpacked data for ForbidAdapter events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ForbidAdapterIterator struct {
	Event *CreditConfiguratorv3ForbidAdapter // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3ForbidAdapterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3ForbidAdapter)
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
		it.Event = new(CreditConfiguratorv3ForbidAdapter)
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
func (it *CreditConfiguratorv3ForbidAdapterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3ForbidAdapterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3ForbidAdapter represents a ForbidAdapter event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ForbidAdapter struct {
	TargetContract common.Address
	Adapter        common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterForbidAdapter is a free log retrieval operation binding the contract event 0x3f688c7b4a117ceec70e927a9ed68836d3da0224eee121f856fc87ad5baa2a80.
//
// Solidity: event ForbidAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterForbidAdapter(opts *bind.FilterOpts, targetContract []common.Address, adapter []common.Address) (*CreditConfiguratorv3ForbidAdapterIterator, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "ForbidAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3ForbidAdapterIterator{contract: _CreditConfiguratorv3.contract, event: "ForbidAdapter", logs: logs, sub: sub}, nil
}

// WatchForbidAdapter is a free log subscription operation binding the contract event 0x3f688c7b4a117ceec70e927a9ed68836d3da0224eee121f856fc87ad5baa2a80.
//
// Solidity: event ForbidAdapter(address indexed targetContract, address indexed adapter)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchForbidAdapter(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3ForbidAdapter, targetContract []common.Address, adapter []common.Address) (event.Subscription, error) {

	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}
	var adapterRule []interface{}
	for _, adapterItem := range adapter {
		adapterRule = append(adapterRule, adapterItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "ForbidAdapter", targetContractRule, adapterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3ForbidAdapter)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ForbidAdapter", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseForbidAdapter(log types.Log) (*CreditConfiguratorv3ForbidAdapter, error) {
	event := new(CreditConfiguratorv3ForbidAdapter)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ForbidAdapter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3ForbidTokenIterator is returned from FilterForbidToken and is used to iterate over the raw logs and unpacked data for ForbidToken events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ForbidTokenIterator struct {
	Event *CreditConfiguratorv3ForbidToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3ForbidTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3ForbidToken)
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
		it.Event = new(CreditConfiguratorv3ForbidToken)
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
func (it *CreditConfiguratorv3ForbidTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3ForbidTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3ForbidToken represents a ForbidToken event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ForbidToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterForbidToken is a free log retrieval operation binding the contract event 0x9d65afef45c30b784a1e4621dbcbb194ebb6aabe16c9a4abce9ab1775a962b76.
//
// Solidity: event ForbidToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterForbidToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3ForbidTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "ForbidToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3ForbidTokenIterator{contract: _CreditConfiguratorv3.contract, event: "ForbidToken", logs: logs, sub: sub}, nil
}

// WatchForbidToken is a free log subscription operation binding the contract event 0x9d65afef45c30b784a1e4621dbcbb194ebb6aabe16c9a4abce9ab1775a962b76.
//
// Solidity: event ForbidToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchForbidToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3ForbidToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "ForbidToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3ForbidToken)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ForbidToken", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseForbidToken(log types.Log) (*CreditConfiguratorv3ForbidToken, error) {
	event := new(CreditConfiguratorv3ForbidToken)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ForbidToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3NewControllerIterator struct {
	Event *CreditConfiguratorv3NewController // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3NewController)
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
		it.Event = new(CreditConfiguratorv3NewController)
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
func (it *CreditConfiguratorv3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3NewController represents a NewController event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*CreditConfiguratorv3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3NewControllerIterator{contract: _CreditConfiguratorv3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3NewController)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "NewController", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseNewController(log types.Log) (*CreditConfiguratorv3NewController, error) {
	event := new(CreditConfiguratorv3NewController)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3PausedIterator struct {
	Event *CreditConfiguratorv3Paused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3Paused)
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
		it.Event = new(CreditConfiguratorv3Paused)
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
func (it *CreditConfiguratorv3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3Paused represents a Paused event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterPaused(opts *bind.FilterOpts) (*CreditConfiguratorv3PausedIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3PausedIterator{contract: _CreditConfiguratorv3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3Paused) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3Paused)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParsePaused(log types.Log) (*CreditConfiguratorv3Paused, error) {
	event := new(CreditConfiguratorv3Paused)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3QuoteTokenIterator is returned from FilterQuoteToken and is used to iterate over the raw logs and unpacked data for QuoteToken events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3QuoteTokenIterator struct {
	Event *CreditConfiguratorv3QuoteToken // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3QuoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3QuoteToken)
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
		it.Event = new(CreditConfiguratorv3QuoteToken)
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
func (it *CreditConfiguratorv3QuoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3QuoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3QuoteToken represents a QuoteToken event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3QuoteToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterQuoteToken is a free log retrieval operation binding the contract event 0xff7d56250177b6941b86ddf2db6637adfc6d47f46540feec1bc0fd5f1326858b.
//
// Solidity: event QuoteToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterQuoteToken(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3QuoteTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "QuoteToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3QuoteTokenIterator{contract: _CreditConfiguratorv3.contract, event: "QuoteToken", logs: logs, sub: sub}, nil
}

// WatchQuoteToken is a free log subscription operation binding the contract event 0xff7d56250177b6941b86ddf2db6637adfc6d47f46540feec1bc0fd5f1326858b.
//
// Solidity: event QuoteToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchQuoteToken(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3QuoteToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "QuoteToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3QuoteToken)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "QuoteToken", log); err != nil {
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

// ParseQuoteToken is a log parse operation binding the contract event 0xff7d56250177b6941b86ddf2db6637adfc6d47f46540feec1bc0fd5f1326858b.
//
// Solidity: event QuoteToken(address indexed token)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseQuoteToken(log types.Log) (*CreditConfiguratorv3QuoteToken, error) {
	event := new(CreditConfiguratorv3QuoteToken)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "QuoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3RemoveEmergencyLiquidatorIterator is returned from FilterRemoveEmergencyLiquidator and is used to iterate over the raw logs and unpacked data for RemoveEmergencyLiquidator events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3RemoveEmergencyLiquidatorIterator struct {
	Event *CreditConfiguratorv3RemoveEmergencyLiquidator // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3RemoveEmergencyLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3RemoveEmergencyLiquidator)
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
		it.Event = new(CreditConfiguratorv3RemoveEmergencyLiquidator)
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
func (it *CreditConfiguratorv3RemoveEmergencyLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3RemoveEmergencyLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3RemoveEmergencyLiquidator represents a RemoveEmergencyLiquidator event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3RemoveEmergencyLiquidator struct {
	Liquidator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveEmergencyLiquidator is a free log retrieval operation binding the contract event 0xc03fe683aa5f2a3776871ebf04508ced24c0335e0d19abd72b6a0d1950e1e255.
//
// Solidity: event RemoveEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterRemoveEmergencyLiquidator(opts *bind.FilterOpts, liquidator []common.Address) (*CreditConfiguratorv3RemoveEmergencyLiquidatorIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "RemoveEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3RemoveEmergencyLiquidatorIterator{contract: _CreditConfiguratorv3.contract, event: "RemoveEmergencyLiquidator", logs: logs, sub: sub}, nil
}

// WatchRemoveEmergencyLiquidator is a free log subscription operation binding the contract event 0xc03fe683aa5f2a3776871ebf04508ced24c0335e0d19abd72b6a0d1950e1e255.
//
// Solidity: event RemoveEmergencyLiquidator(address indexed liquidator)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchRemoveEmergencyLiquidator(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3RemoveEmergencyLiquidator, liquidator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "RemoveEmergencyLiquidator", liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3RemoveEmergencyLiquidator)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "RemoveEmergencyLiquidator", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseRemoveEmergencyLiquidator(log types.Log) (*CreditConfiguratorv3RemoveEmergencyLiquidator, error) {
	event := new(CreditConfiguratorv3RemoveEmergencyLiquidator)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "RemoveEmergencyLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3ResetCumulativeLossIterator is returned from FilterResetCumulativeLoss and is used to iterate over the raw logs and unpacked data for ResetCumulativeLoss events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ResetCumulativeLossIterator struct {
	Event *CreditConfiguratorv3ResetCumulativeLoss // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3ResetCumulativeLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3ResetCumulativeLoss)
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
		it.Event = new(CreditConfiguratorv3ResetCumulativeLoss)
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
func (it *CreditConfiguratorv3ResetCumulativeLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3ResetCumulativeLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3ResetCumulativeLoss represents a ResetCumulativeLoss event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ResetCumulativeLoss struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResetCumulativeLoss is a free log retrieval operation binding the contract event 0x45013665e7af2da20f8bbc745e6760467c3c18d820f7052ad772158ce360d92d.
//
// Solidity: event ResetCumulativeLoss()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterResetCumulativeLoss(opts *bind.FilterOpts) (*CreditConfiguratorv3ResetCumulativeLossIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "ResetCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3ResetCumulativeLossIterator{contract: _CreditConfiguratorv3.contract, event: "ResetCumulativeLoss", logs: logs, sub: sub}, nil
}

// WatchResetCumulativeLoss is a free log subscription operation binding the contract event 0x45013665e7af2da20f8bbc745e6760467c3c18d820f7052ad772158ce360d92d.
//
// Solidity: event ResetCumulativeLoss()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchResetCumulativeLoss(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3ResetCumulativeLoss) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "ResetCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3ResetCumulativeLoss)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ResetCumulativeLoss", log); err != nil {
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

// ParseResetCumulativeLoss is a log parse operation binding the contract event 0x45013665e7af2da20f8bbc745e6760467c3c18d820f7052ad772158ce360d92d.
//
// Solidity: event ResetCumulativeLoss()
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseResetCumulativeLoss(log types.Log) (*CreditConfiguratorv3ResetCumulativeLoss, error) {
	event := new(CreditConfiguratorv3ResetCumulativeLoss)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ResetCumulativeLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator is returned from FilterScheduleTokenLiquidationThresholdRamp and is used to iterate over the raw logs and unpacked data for ScheduleTokenLiquidationThresholdRamp events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator struct {
	Event *CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp)
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
		it.Event = new(CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp)
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
func (it *CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp represents a ScheduleTokenLiquidationThresholdRamp event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp struct {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterScheduleTokenLiquidationThresholdRamp(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "ScheduleTokenLiquidationThresholdRamp", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3ScheduleTokenLiquidationThresholdRampIterator{contract: _CreditConfiguratorv3.contract, event: "ScheduleTokenLiquidationThresholdRamp", logs: logs, sub: sub}, nil
}

// WatchScheduleTokenLiquidationThresholdRamp is a free log subscription operation binding the contract event 0xa8193c198aab4146e3640f414ba8473918c6d028f45b27fb08b185a16c15ce23.
//
// Solidity: event ScheduleTokenLiquidationThresholdRamp(address indexed token, uint16 liquidationThresholdInitial, uint16 liquidationThresholdFinal, uint40 timestampRampStart, uint40 timestampRampEnd)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchScheduleTokenLiquidationThresholdRamp(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "ScheduleTokenLiquidationThresholdRamp", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ScheduleTokenLiquidationThresholdRamp", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseScheduleTokenLiquidationThresholdRamp(log types.Log) (*CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp, error) {
	event := new(CreditConfiguratorv3ScheduleTokenLiquidationThresholdRamp)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "ScheduleTokenLiquidationThresholdRamp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetBorrowingLimitsIterator is returned from FilterSetBorrowingLimits and is used to iterate over the raw logs and unpacked data for SetBorrowingLimits events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetBorrowingLimitsIterator struct {
	Event *CreditConfiguratorv3SetBorrowingLimits // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetBorrowingLimitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetBorrowingLimits)
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
		it.Event = new(CreditConfiguratorv3SetBorrowingLimits)
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
func (it *CreditConfiguratorv3SetBorrowingLimitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetBorrowingLimitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetBorrowingLimits represents a SetBorrowingLimits event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetBorrowingLimits struct {
	MinDebt *big.Int
	MaxDebt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBorrowingLimits is a free log retrieval operation binding the contract event 0xb2cc80ffa4c2f75731dbb99fcd29cccd7829c55d4cd5d6a884506b1435d6d1f3.
//
// Solidity: event SetBorrowingLimits(uint256 minDebt, uint256 maxDebt)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetBorrowingLimits(opts *bind.FilterOpts) (*CreditConfiguratorv3SetBorrowingLimitsIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetBorrowingLimits")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetBorrowingLimitsIterator{contract: _CreditConfiguratorv3.contract, event: "SetBorrowingLimits", logs: logs, sub: sub}, nil
}

// WatchSetBorrowingLimits is a free log subscription operation binding the contract event 0xb2cc80ffa4c2f75731dbb99fcd29cccd7829c55d4cd5d6a884506b1435d6d1f3.
//
// Solidity: event SetBorrowingLimits(uint256 minDebt, uint256 maxDebt)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetBorrowingLimits(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetBorrowingLimits) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetBorrowingLimits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetBorrowingLimits)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetBorrowingLimits", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetBorrowingLimits(log types.Log) (*CreditConfiguratorv3SetBorrowingLimits, error) {
	event := new(CreditConfiguratorv3SetBorrowingLimits)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetBorrowingLimits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetBotListIterator is returned from FilterSetBotList and is used to iterate over the raw logs and unpacked data for SetBotList events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetBotListIterator struct {
	Event *CreditConfiguratorv3SetBotList // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetBotListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetBotList)
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
		it.Event = new(CreditConfiguratorv3SetBotList)
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
func (it *CreditConfiguratorv3SetBotListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetBotListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetBotList represents a SetBotList event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetBotList struct {
	BotList common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBotList is a free log retrieval operation binding the contract event 0x93c2c5d02f328eae23f02b70bcaf858ad014f03c74e0a10cb57124d45ea3e542.
//
// Solidity: event SetBotList(address indexed botList)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetBotList(opts *bind.FilterOpts, botList []common.Address) (*CreditConfiguratorv3SetBotListIterator, error) {

	var botListRule []interface{}
	for _, botListItem := range botList {
		botListRule = append(botListRule, botListItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetBotList", botListRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetBotListIterator{contract: _CreditConfiguratorv3.contract, event: "SetBotList", logs: logs, sub: sub}, nil
}

// WatchSetBotList is a free log subscription operation binding the contract event 0x93c2c5d02f328eae23f02b70bcaf858ad014f03c74e0a10cb57124d45ea3e542.
//
// Solidity: event SetBotList(address indexed botList)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetBotList(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetBotList, botList []common.Address) (event.Subscription, error) {

	var botListRule []interface{}
	for _, botListItem := range botList {
		botListRule = append(botListRule, botListItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetBotList", botListRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetBotList)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetBotList", log); err != nil {
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

// ParseSetBotList is a log parse operation binding the contract event 0x93c2c5d02f328eae23f02b70bcaf858ad014f03c74e0a10cb57124d45ea3e542.
//
// Solidity: event SetBotList(address indexed botList)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetBotList(log types.Log) (*CreditConfiguratorv3SetBotList, error) {
	event := new(CreditConfiguratorv3SetBotList)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetBotList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetCreditFacadeIterator is returned from FilterSetCreditFacade and is used to iterate over the raw logs and unpacked data for SetCreditFacade events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetCreditFacadeIterator struct {
	Event *CreditConfiguratorv3SetCreditFacade // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetCreditFacadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetCreditFacade)
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
		it.Event = new(CreditConfiguratorv3SetCreditFacade)
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
func (it *CreditConfiguratorv3SetCreditFacadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetCreditFacadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetCreditFacade represents a SetCreditFacade event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetCreditFacade struct {
	CreditFacade common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetCreditFacade is a free log retrieval operation binding the contract event 0x1cd439329e916b95ce297eb699326f2799c8de28be6bba10f28db1d9067778f1.
//
// Solidity: event SetCreditFacade(address indexed creditFacade)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetCreditFacade(opts *bind.FilterOpts, creditFacade []common.Address) (*CreditConfiguratorv3SetCreditFacadeIterator, error) {

	var creditFacadeRule []interface{}
	for _, creditFacadeItem := range creditFacade {
		creditFacadeRule = append(creditFacadeRule, creditFacadeItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetCreditFacade", creditFacadeRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetCreditFacadeIterator{contract: _CreditConfiguratorv3.contract, event: "SetCreditFacade", logs: logs, sub: sub}, nil
}

// WatchSetCreditFacade is a free log subscription operation binding the contract event 0x1cd439329e916b95ce297eb699326f2799c8de28be6bba10f28db1d9067778f1.
//
// Solidity: event SetCreditFacade(address indexed creditFacade)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetCreditFacade(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetCreditFacade, creditFacade []common.Address) (event.Subscription, error) {

	var creditFacadeRule []interface{}
	for _, creditFacadeItem := range creditFacade {
		creditFacadeRule = append(creditFacadeRule, creditFacadeItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetCreditFacade", creditFacadeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetCreditFacade)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetCreditFacade", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetCreditFacade(log types.Log) (*CreditConfiguratorv3SetCreditFacade, error) {
	event := new(CreditConfiguratorv3SetCreditFacade)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetCreditFacade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetExpirationDateIterator is returned from FilterSetExpirationDate and is used to iterate over the raw logs and unpacked data for SetExpirationDate events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetExpirationDateIterator struct {
	Event *CreditConfiguratorv3SetExpirationDate // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetExpirationDateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetExpirationDate)
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
		it.Event = new(CreditConfiguratorv3SetExpirationDate)
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
func (it *CreditConfiguratorv3SetExpirationDateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetExpirationDateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetExpirationDate represents a SetExpirationDate event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetExpirationDate struct {
	ExpirationDate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetExpirationDate is a free log retrieval operation binding the contract event 0xb019cf1dc4b3caa72aa4723abcc271a2bb3138bee0a89cd911fb8980b0c93d56.
//
// Solidity: event SetExpirationDate(uint40 expirationDate)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetExpirationDate(opts *bind.FilterOpts) (*CreditConfiguratorv3SetExpirationDateIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetExpirationDate")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetExpirationDateIterator{contract: _CreditConfiguratorv3.contract, event: "SetExpirationDate", logs: logs, sub: sub}, nil
}

// WatchSetExpirationDate is a free log subscription operation binding the contract event 0xb019cf1dc4b3caa72aa4723abcc271a2bb3138bee0a89cd911fb8980b0c93d56.
//
// Solidity: event SetExpirationDate(uint40 expirationDate)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetExpirationDate(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetExpirationDate) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetExpirationDate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetExpirationDate)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetExpirationDate", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetExpirationDate(log types.Log) (*CreditConfiguratorv3SetExpirationDate, error) {
	event := new(CreditConfiguratorv3SetExpirationDate)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetExpirationDate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetMaxCumulativeLossIterator is returned from FilterSetMaxCumulativeLoss and is used to iterate over the raw logs and unpacked data for SetMaxCumulativeLoss events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxCumulativeLossIterator struct {
	Event *CreditConfiguratorv3SetMaxCumulativeLoss // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetMaxCumulativeLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetMaxCumulativeLoss)
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
		it.Event = new(CreditConfiguratorv3SetMaxCumulativeLoss)
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
func (it *CreditConfiguratorv3SetMaxCumulativeLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetMaxCumulativeLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetMaxCumulativeLoss represents a SetMaxCumulativeLoss event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxCumulativeLoss struct {
	MaxCumulativeLoss *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMaxCumulativeLoss is a free log retrieval operation binding the contract event 0x349a31f3899f92873d644a2bb70787ec009398cb92d694f8420a9f03c7e3b0b1.
//
// Solidity: event SetMaxCumulativeLoss(uint128 maxCumulativeLoss)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetMaxCumulativeLoss(opts *bind.FilterOpts) (*CreditConfiguratorv3SetMaxCumulativeLossIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetMaxCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetMaxCumulativeLossIterator{contract: _CreditConfiguratorv3.contract, event: "SetMaxCumulativeLoss", logs: logs, sub: sub}, nil
}

// WatchSetMaxCumulativeLoss is a free log subscription operation binding the contract event 0x349a31f3899f92873d644a2bb70787ec009398cb92d694f8420a9f03c7e3b0b1.
//
// Solidity: event SetMaxCumulativeLoss(uint128 maxCumulativeLoss)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetMaxCumulativeLoss(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetMaxCumulativeLoss) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetMaxCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetMaxCumulativeLoss)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxCumulativeLoss", log); err != nil {
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

// ParseSetMaxCumulativeLoss is a log parse operation binding the contract event 0x349a31f3899f92873d644a2bb70787ec009398cb92d694f8420a9f03c7e3b0b1.
//
// Solidity: event SetMaxCumulativeLoss(uint128 maxCumulativeLoss)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetMaxCumulativeLoss(log types.Log) (*CreditConfiguratorv3SetMaxCumulativeLoss, error) {
	event := new(CreditConfiguratorv3SetMaxCumulativeLoss)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxCumulativeLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator is returned from FilterSetMaxDebtPerBlockMultiplier and is used to iterate over the raw logs and unpacked data for SetMaxDebtPerBlockMultiplier events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator struct {
	Event *CreditConfiguratorv3SetMaxDebtPerBlockMultiplier // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetMaxDebtPerBlockMultiplier)
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
		it.Event = new(CreditConfiguratorv3SetMaxDebtPerBlockMultiplier)
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
func (it *CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetMaxDebtPerBlockMultiplier represents a SetMaxDebtPerBlockMultiplier event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxDebtPerBlockMultiplier struct {
	MaxDebtPerBlockMultiplier uint8
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetMaxDebtPerBlockMultiplier is a free log retrieval operation binding the contract event 0xaebbd82c9dcdcd553331f5850bbdf5add33bf8fce5c7c76e2c9e7912ad5f1564.
//
// Solidity: event SetMaxDebtPerBlockMultiplier(uint8 maxDebtPerBlockMultiplier)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetMaxDebtPerBlockMultiplier(opts *bind.FilterOpts) (*CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetMaxDebtPerBlockMultiplier")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetMaxDebtPerBlockMultiplierIterator{contract: _CreditConfiguratorv3.contract, event: "SetMaxDebtPerBlockMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetMaxDebtPerBlockMultiplier is a free log subscription operation binding the contract event 0xaebbd82c9dcdcd553331f5850bbdf5add33bf8fce5c7c76e2c9e7912ad5f1564.
//
// Solidity: event SetMaxDebtPerBlockMultiplier(uint8 maxDebtPerBlockMultiplier)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetMaxDebtPerBlockMultiplier(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetMaxDebtPerBlockMultiplier) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetMaxDebtPerBlockMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetMaxDebtPerBlockMultiplier)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxDebtPerBlockMultiplier", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetMaxDebtPerBlockMultiplier(log types.Log) (*CreditConfiguratorv3SetMaxDebtPerBlockMultiplier, error) {
	event := new(CreditConfiguratorv3SetMaxDebtPerBlockMultiplier)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxDebtPerBlockMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetMaxEnabledTokensIterator is returned from FilterSetMaxEnabledTokens and is used to iterate over the raw logs and unpacked data for SetMaxEnabledTokens events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxEnabledTokensIterator struct {
	Event *CreditConfiguratorv3SetMaxEnabledTokens // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetMaxEnabledTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetMaxEnabledTokens)
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
		it.Event = new(CreditConfiguratorv3SetMaxEnabledTokens)
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
func (it *CreditConfiguratorv3SetMaxEnabledTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetMaxEnabledTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetMaxEnabledTokens represents a SetMaxEnabledTokens event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetMaxEnabledTokens struct {
	MaxEnabledTokens uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetMaxEnabledTokens is a free log retrieval operation binding the contract event 0x289338cf948e424769e26fe06f36f4b1e62b60837ad92f16d81f61607c89b698.
//
// Solidity: event SetMaxEnabledTokens(uint8 maxEnabledTokens)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetMaxEnabledTokens(opts *bind.FilterOpts) (*CreditConfiguratorv3SetMaxEnabledTokensIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetMaxEnabledTokens")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetMaxEnabledTokensIterator{contract: _CreditConfiguratorv3.contract, event: "SetMaxEnabledTokens", logs: logs, sub: sub}, nil
}

// WatchSetMaxEnabledTokens is a free log subscription operation binding the contract event 0x289338cf948e424769e26fe06f36f4b1e62b60837ad92f16d81f61607c89b698.
//
// Solidity: event SetMaxEnabledTokens(uint8 maxEnabledTokens)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetMaxEnabledTokens(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetMaxEnabledTokens) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetMaxEnabledTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetMaxEnabledTokens)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxEnabledTokens", log); err != nil {
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

// ParseSetMaxEnabledTokens is a log parse operation binding the contract event 0x289338cf948e424769e26fe06f36f4b1e62b60837ad92f16d81f61607c89b698.
//
// Solidity: event SetMaxEnabledTokens(uint8 maxEnabledTokens)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetMaxEnabledTokens(log types.Log) (*CreditConfiguratorv3SetMaxEnabledTokens, error) {
	event := new(CreditConfiguratorv3SetMaxEnabledTokens)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetMaxEnabledTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetPriceOracleIterator is returned from FilterSetPriceOracle and is used to iterate over the raw logs and unpacked data for SetPriceOracle events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetPriceOracleIterator struct {
	Event *CreditConfiguratorv3SetPriceOracle // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetPriceOracle)
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
		it.Event = new(CreditConfiguratorv3SetPriceOracle)
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
func (it *CreditConfiguratorv3SetPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetPriceOracle represents a SetPriceOracle event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetPriceOracle struct {
	PriceOracle common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetPriceOracle is a free log retrieval operation binding the contract event 0x88a686e0e341d9099f2f990c3aa759a86822142a67579064b43ded9354a25662.
//
// Solidity: event SetPriceOracle(address indexed priceOracle)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetPriceOracle(opts *bind.FilterOpts, priceOracle []common.Address) (*CreditConfiguratorv3SetPriceOracleIterator, error) {

	var priceOracleRule []interface{}
	for _, priceOracleItem := range priceOracle {
		priceOracleRule = append(priceOracleRule, priceOracleItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetPriceOracle", priceOracleRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetPriceOracleIterator{contract: _CreditConfiguratorv3.contract, event: "SetPriceOracle", logs: logs, sub: sub}, nil
}

// WatchSetPriceOracle is a free log subscription operation binding the contract event 0x88a686e0e341d9099f2f990c3aa759a86822142a67579064b43ded9354a25662.
//
// Solidity: event SetPriceOracle(address indexed priceOracle)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetPriceOracle(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetPriceOracle, priceOracle []common.Address) (event.Subscription, error) {

	var priceOracleRule []interface{}
	for _, priceOracleItem := range priceOracle {
		priceOracleRule = append(priceOracleRule, priceOracleItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetPriceOracle", priceOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetPriceOracle)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetPriceOracle", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetPriceOracle(log types.Log) (*CreditConfiguratorv3SetPriceOracle, error) {
	event := new(CreditConfiguratorv3SetPriceOracle)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3SetTokenLiquidationThresholdIterator is returned from FilterSetTokenLiquidationThreshold and is used to iterate over the raw logs and unpacked data for SetTokenLiquidationThreshold events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetTokenLiquidationThresholdIterator struct {
	Event *CreditConfiguratorv3SetTokenLiquidationThreshold // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3SetTokenLiquidationThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3SetTokenLiquidationThreshold)
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
		it.Event = new(CreditConfiguratorv3SetTokenLiquidationThreshold)
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
func (it *CreditConfiguratorv3SetTokenLiquidationThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3SetTokenLiquidationThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3SetTokenLiquidationThreshold represents a SetTokenLiquidationThreshold event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3SetTokenLiquidationThreshold struct {
	Token                common.Address
	LiquidationThreshold uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTokenLiquidationThreshold is a free log retrieval operation binding the contract event 0xda5e841a0cb137f4a60661969e409f01ef7627723a4a929414e4f69b5475ee8c.
//
// Solidity: event SetTokenLiquidationThreshold(address indexed token, uint16 liquidationThreshold)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterSetTokenLiquidationThreshold(opts *bind.FilterOpts, token []common.Address) (*CreditConfiguratorv3SetTokenLiquidationThresholdIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "SetTokenLiquidationThreshold", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3SetTokenLiquidationThresholdIterator{contract: _CreditConfiguratorv3.contract, event: "SetTokenLiquidationThreshold", logs: logs, sub: sub}, nil
}

// WatchSetTokenLiquidationThreshold is a free log subscription operation binding the contract event 0xda5e841a0cb137f4a60661969e409f01ef7627723a4a929414e4f69b5475ee8c.
//
// Solidity: event SetTokenLiquidationThreshold(address indexed token, uint16 liquidationThreshold)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchSetTokenLiquidationThreshold(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3SetTokenLiquidationThreshold, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "SetTokenLiquidationThreshold", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3SetTokenLiquidationThreshold)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetTokenLiquidationThreshold", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseSetTokenLiquidationThreshold(log types.Log) (*CreditConfiguratorv3SetTokenLiquidationThreshold, error) {
	event := new(CreditConfiguratorv3SetTokenLiquidationThreshold)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "SetTokenLiquidationThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3UnpausedIterator struct {
	Event *CreditConfiguratorv3Unpaused // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3Unpaused)
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
		it.Event = new(CreditConfiguratorv3Unpaused)
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
func (it *CreditConfiguratorv3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3Unpaused represents a Unpaused event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditConfiguratorv3UnpausedIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3UnpausedIterator{contract: _CreditConfiguratorv3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3Unpaused) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3Unpaused)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseUnpaused(log types.Log) (*CreditConfiguratorv3Unpaused, error) {
	event := new(CreditConfiguratorv3Unpaused)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfiguratorv3UpdateFeesIterator is returned from FilterUpdateFees and is used to iterate over the raw logs and unpacked data for UpdateFees events raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3UpdateFeesIterator struct {
	Event *CreditConfiguratorv3UpdateFees // Event containing the contract specifics and raw log

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
func (it *CreditConfiguratorv3UpdateFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfiguratorv3UpdateFees)
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
		it.Event = new(CreditConfiguratorv3UpdateFees)
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
func (it *CreditConfiguratorv3UpdateFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfiguratorv3UpdateFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfiguratorv3UpdateFees represents a UpdateFees event raised by the CreditConfiguratorv3 contract.
type CreditConfiguratorv3UpdateFees struct {
	FeeInterest               uint16
	FeeLiquidation            uint16
	LiquidationPremium        uint16
	FeeLiquidationExpired     uint16
	LiquidationPremiumExpired uint16
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterUpdateFees is a free log retrieval operation binding the contract event 0x214e595f9b6bdb12147befdaaea117fe5a00b2e9217e2e646923f6092798e7de.
//
// Solidity: event UpdateFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) FilterUpdateFees(opts *bind.FilterOpts) (*CreditConfiguratorv3UpdateFeesIterator, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.FilterLogs(opts, "UpdateFees")
	if err != nil {
		return nil, err
	}
	return &CreditConfiguratorv3UpdateFeesIterator{contract: _CreditConfiguratorv3.contract, event: "UpdateFees", logs: logs, sub: sub}, nil
}

// WatchUpdateFees is a free log subscription operation binding the contract event 0x214e595f9b6bdb12147befdaaea117fe5a00b2e9217e2e646923f6092798e7de.
//
// Solidity: event UpdateFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) WatchUpdateFees(opts *bind.WatchOpts, sink chan<- *CreditConfiguratorv3UpdateFees) (event.Subscription, error) {

	logs, sub, err := _CreditConfiguratorv3.contract.WatchLogs(opts, "UpdateFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfiguratorv3UpdateFees)
				if err := _CreditConfiguratorv3.contract.UnpackLog(event, "UpdateFees", log); err != nil {
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

// ParseUpdateFees is a log parse operation binding the contract event 0x214e595f9b6bdb12147befdaaea117fe5a00b2e9217e2e646923f6092798e7de.
//
// Solidity: event UpdateFees(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationPremium, uint16 feeLiquidationExpired, uint16 liquidationPremiumExpired)
func (_CreditConfiguratorv3 *CreditConfiguratorv3Filterer) ParseUpdateFees(log types.Log) (*CreditConfiguratorv3UpdateFees, error) {
	event := new(CreditConfiguratorv3UpdateFees)
	if err := _CreditConfiguratorv3.contract.UnpackLog(event, "UpdateFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
