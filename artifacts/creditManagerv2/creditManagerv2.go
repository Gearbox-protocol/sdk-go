// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManagerv2

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

// CreditManagerv2MetaData contains all meta data concerning the CreditManagerv2 contract.
var CreditManagerv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdaptersOrCreditFacadeOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowanceFailedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditConfiguratorOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditFacadeOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HasNoOpenedAccountException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughCollateralException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyLockException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetContractNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyAddedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyEnabledTokensException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyTokensException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressOrUserAlreadyHasAccountException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"ExecuteOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newConfigurator\",\"type\":\"address\"}],\"name\":\"NewConfigurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_accountFactory\",\"outputs\":[{\"internalType\":\"contractIAccountFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acl\",\"outputs\":[{\"internalType\":\"contractIACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"adapterToContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"addEmergencyLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumClosureAction\",\"name\":\"closureActionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountWithInterest\",\"type\":\"uint256\"}],\"name\":\"calcClosePayments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToPool\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"calcCreditAccountAccruedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountWithInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountWithInterestAndFees\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canLiquidateWhilePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"changeContractAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"checkAndEnableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"checkAndOptimizeEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"enumClosureAction\",\"name\":\"closureActionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"skipTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"convertWETH\",\"type\":\"bool\"}],\"name\":\"closeCreditAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"collateralTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenMask\",\"type\":\"uint256\"}],\"name\":\"collateralTokensByMask\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralTokensCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractToAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeDropAtFastCheckRAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"disableToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabledTokensMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeOrder\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceInBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOutBefore\",\"type\":\"uint256\"}],\"name\":\"fastCollateralCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forbiddenTokenMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"fullCollateralCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountOrRevert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"liquidationThresholds\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"lt\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"increase\",\"type\":\"bool\"}],\"name\":\"manageDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newBorrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllowedEnabledTokenLength\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"openCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"contractIPriceOracleV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"}],\"name\":\"removeEmergencyLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditConfigurator\",\"type\":\"address\"}],\"name\":\"setConfigurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_forbidMask\",\"type\":\"uint256\"}],\"name\":\"setForbidMask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"name\":\"setLiquidationThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newMaxEnabledTokens\",\"type\":\"uint8\"}],\"name\":\"setMaxEnabledTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_liquidationDiscountExpired\",\"type\":\"uint16\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenMasksMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferAccountOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"universalAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditFacade\",\"type\":\"address\"}],\"name\":\"upgradeCreditFacade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceOracle\",\"type\":\"address\"}],\"name\":\"upgradePriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditManagerv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditManagerv2MetaData.ABI instead.
var CreditManagerv2ABI = CreditManagerv2MetaData.ABI

// CreditManagerv2 is an auto generated Go binding around an Ethereum contract.
type CreditManagerv2 struct {
	CreditManagerv2Caller     // Read-only binding to the contract
	CreditManagerv2Transactor // Write-only binding to the contract
	CreditManagerv2Filterer   // Log filterer for contract events
}

// CreditManagerv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerv2Session struct {
	Contract     *CreditManagerv2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditManagerv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerv2CallerSession struct {
	Contract *CreditManagerv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CreditManagerv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerv2TransactorSession struct {
	Contract     *CreditManagerv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CreditManagerv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerv2Raw struct {
	Contract *CreditManagerv2 // Generic contract binding to access the raw methods on
}

// CreditManagerv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerv2CallerRaw struct {
	Contract *CreditManagerv2Caller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerv2TransactorRaw struct {
	Contract *CreditManagerv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManagerv2 creates a new instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2(address common.Address, backend bind.ContractBackend) (*CreditManagerv2, error) {
	contract, err := bindCreditManagerv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2{CreditManagerv2Caller: CreditManagerv2Caller{contract: contract}, CreditManagerv2Transactor: CreditManagerv2Transactor{contract: contract}, CreditManagerv2Filterer: CreditManagerv2Filterer{contract: contract}}, nil
}

// NewCreditManagerv2Caller creates a new read-only instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Caller(address common.Address, caller bind.ContractCaller) (*CreditManagerv2Caller, error) {
	contract, err := bindCreditManagerv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Caller{contract: contract}, nil
}

// NewCreditManagerv2Transactor creates a new write-only instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerv2Transactor, error) {
	contract, err := bindCreditManagerv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Transactor{contract: contract}, nil
}

// NewCreditManagerv2Filterer creates a new log filterer instance of CreditManagerv2, bound to a specific deployed contract.
func NewCreditManagerv2Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerv2Filterer, error) {
	contract, err := bindCreditManagerv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2Filterer{contract: contract}, nil
}

// bindCreditManagerv2 binds a generic wrapper to an already deployed contract.
func bindCreditManagerv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditManagerv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv2 *CreditManagerv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv2.Contract.CreditManagerv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv2 *CreditManagerv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CreditManagerv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv2 *CreditManagerv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CreditManagerv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv2 *CreditManagerv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv2 *CreditManagerv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv2 *CreditManagerv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.contract.Transact(opts, method, params...)
}

// AccountFactory is a free data retrieval call binding the contract method 0xdb7ceb80.
//
// Solidity: function _accountFactory() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) AccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "_accountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountFactory is a free data retrieval call binding the contract method 0xdb7ceb80.
//
// Solidity: function _accountFactory() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) AccountFactory() (common.Address, error) {
	return _CreditManagerv2.Contract.AccountFactory(&_CreditManagerv2.CallOpts)
}

// AccountFactory is a free data retrieval call binding the contract method 0xdb7ceb80.
//
// Solidity: function _accountFactory() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) AccountFactory() (common.Address, error) {
	return _CreditManagerv2.Contract.AccountFactory(&_CreditManagerv2.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "_acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) Acl() (common.Address, error) {
	return _CreditManagerv2.Contract.Acl(&_CreditManagerv2.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Acl() (common.Address, error) {
	return _CreditManagerv2.Contract.Acl(&_CreditManagerv2.CallOpts)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) AdapterToContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "adapterToContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.AdapterToContract(&_CreditManagerv2.CallOpts, arg0)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.AdapterToContract(&_CreditManagerv2.CallOpts, arg0)
}

// CalcClosePayments is a free data retrieval call binding the contract method 0x5063524a.
//
// Solidity: function calcClosePayments(uint256 totalValue, uint8 closureActionType, uint256 borrowedAmount, uint256 borrowedAmountWithInterest) view returns(uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManagerv2 *CreditManagerv2Caller) CalcClosePayments(opts *bind.CallOpts, totalValue *big.Int, closureActionType uint8, borrowedAmount *big.Int, borrowedAmountWithInterest *big.Int) (struct {
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "calcClosePayments", totalValue, closureActionType, borrowedAmount, borrowedAmountWithInterest)

	outstruct := new(struct {
		AmountToPool   *big.Int
		RemainingFunds *big.Int
		Profit         *big.Int
		Loss           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountToPool = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RemainingFunds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Profit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Loss = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalcClosePayments is a free data retrieval call binding the contract method 0x5063524a.
//
// Solidity: function calcClosePayments(uint256 totalValue, uint8 closureActionType, uint256 borrowedAmount, uint256 borrowedAmountWithInterest) view returns(uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManagerv2 *CreditManagerv2Session) CalcClosePayments(totalValue *big.Int, closureActionType uint8, borrowedAmount *big.Int, borrowedAmountWithInterest *big.Int) (struct {
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcClosePayments(&_CreditManagerv2.CallOpts, totalValue, closureActionType, borrowedAmount, borrowedAmountWithInterest)
}

// CalcClosePayments is a free data retrieval call binding the contract method 0x5063524a.
//
// Solidity: function calcClosePayments(uint256 totalValue, uint8 closureActionType, uint256 borrowedAmount, uint256 borrowedAmountWithInterest) view returns(uint256 amountToPool, uint256 remainingFunds, uint256 profit, uint256 loss)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CalcClosePayments(totalValue *big.Int, closureActionType uint8, borrowedAmount *big.Int, borrowedAmountWithInterest *big.Int) (struct {
	AmountToPool   *big.Int
	RemainingFunds *big.Int
	Profit         *big.Int
	Loss           *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcClosePayments(&_CreditManagerv2.CallOpts, totalValue, closureActionType, borrowedAmount, borrowedAmountWithInterest)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest, uint256 borrowedAmountWithInterestAndFees)
func (_CreditManagerv2 *CreditManagerv2Caller) CalcCreditAccountAccruedInterest(opts *bind.CallOpts, creditAccount common.Address) (struct {
	BorrowedAmount                    *big.Int
	BorrowedAmountWithInterest        *big.Int
	BorrowedAmountWithInterestAndFees *big.Int
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "calcCreditAccountAccruedInterest", creditAccount)

	outstruct := new(struct {
		BorrowedAmount                    *big.Int
		BorrowedAmountWithInterest        *big.Int
		BorrowedAmountWithInterestAndFees *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BorrowedAmountWithInterest = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BorrowedAmountWithInterestAndFees = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest, uint256 borrowedAmountWithInterestAndFees)
func (_CreditManagerv2 *CreditManagerv2Session) CalcCreditAccountAccruedInterest(creditAccount common.Address) (struct {
	BorrowedAmount                    *big.Int
	BorrowedAmountWithInterest        *big.Int
	BorrowedAmountWithInterestAndFees *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcCreditAccountAccruedInterest(&_CreditManagerv2.CallOpts, creditAccount)
}

// CalcCreditAccountAccruedInterest is a free data retrieval call binding the contract method 0x3192195c.
//
// Solidity: function calcCreditAccountAccruedInterest(address creditAccount) view returns(uint256 borrowedAmount, uint256 borrowedAmountWithInterest, uint256 borrowedAmountWithInterestAndFees)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CalcCreditAccountAccruedInterest(creditAccount common.Address) (struct {
	BorrowedAmount                    *big.Int
	BorrowedAmountWithInterest        *big.Int
	BorrowedAmountWithInterestAndFees *big.Int
}, error) {
	return _CreditManagerv2.Contract.CalcCreditAccountAccruedInterest(&_CreditManagerv2.CallOpts, creditAccount)
}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Caller) CanLiquidateWhilePaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "canLiquidateWhilePaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Session) CanLiquidateWhilePaused(arg0 common.Address) (bool, error) {
	return _CreditManagerv2.Contract.CanLiquidateWhilePaused(&_CreditManagerv2.CallOpts, arg0)
}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CanLiquidateWhilePaused(arg0 common.Address) (bool, error) {
	return _CreditManagerv2.Contract.CanLiquidateWhilePaused(&_CreditManagerv2.CallOpts, arg0)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 id) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2Caller) CollateralTokens(opts *bind.CallOpts, id *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "collateralTokens", id)

	outstruct := new(struct {
		Token                common.Address
		LiquidationThreshold uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 id) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2Session) CollateralTokens(id *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv2.Contract.CollateralTokens(&_CreditManagerv2.CallOpts, id)
}

// CollateralTokens is a free data retrieval call binding the contract method 0x172c48c7.
//
// Solidity: function collateralTokens(uint256 id) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CollateralTokens(id *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv2.Contract.CollateralTokens(&_CreditManagerv2.CallOpts, id)
}

// CollateralTokensByMask is a free data retrieval call binding the contract method 0xe75538c7.
//
// Solidity: function collateralTokensByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2Caller) CollateralTokensByMask(opts *bind.CallOpts, tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "collateralTokensByMask", tokenMask)

	outstruct := new(struct {
		Token                common.Address
		LiquidationThreshold uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// CollateralTokensByMask is a free data retrieval call binding the contract method 0xe75538c7.
//
// Solidity: function collateralTokensByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2Session) CollateralTokensByMask(tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv2.Contract.CollateralTokensByMask(&_CreditManagerv2.CallOpts, tokenMask)
}

// CollateralTokensByMask is a free data retrieval call binding the contract method 0xe75538c7.
//
// Solidity: function collateralTokensByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CollateralTokensByMask(tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv2.Contract.CollateralTokensByMask(&_CreditManagerv2.CallOpts, tokenMask)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) CollateralTokensCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "collateralTokensCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) CollateralTokensCount() (*big.Int, error) {
	return _CreditManagerv2.Contract.CollateralTokensCount(&_CreditManagerv2.CallOpts)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CollateralTokensCount() (*big.Int, error) {
	return _CreditManagerv2.Contract.CollateralTokensCount(&_CreditManagerv2.CallOpts)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) ContractToAdapter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "contractToAdapter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.ContractToAdapter(&_CreditManagerv2.CallOpts, arg0)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.ContractToAdapter(&_CreditManagerv2.CallOpts, arg0)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditAccounts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.CreditAccounts(&_CreditManagerv2.CallOpts, arg0)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x055ee9b5.
//
// Solidity: function creditAccounts(address ) view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditAccounts(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.CreditAccounts(&_CreditManagerv2.CallOpts, arg0)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditConfigurator(&_CreditManagerv2.CallOpts)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditConfigurator(&_CreditManagerv2.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) CreditFacade() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditFacade(&_CreditManagerv2.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CreditFacade() (common.Address, error) {
	return _CreditManagerv2.Contract.CreditFacade(&_CreditManagerv2.CallOpts)
}

// CumulativeDropAtFastCheckRAY is a free data retrieval call binding the contract method 0x3e8297ca.
//
// Solidity: function cumulativeDropAtFastCheckRAY(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) CumulativeDropAtFastCheckRAY(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "cumulativeDropAtFastCheckRAY", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeDropAtFastCheckRAY is a free data retrieval call binding the contract method 0x3e8297ca.
//
// Solidity: function cumulativeDropAtFastCheckRAY(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) CumulativeDropAtFastCheckRAY(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.CumulativeDropAtFastCheckRAY(&_CreditManagerv2.CallOpts, arg0)
}

// CumulativeDropAtFastCheckRAY is a free data retrieval call binding the contract method 0x3e8297ca.
//
// Solidity: function cumulativeDropAtFastCheckRAY(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) CumulativeDropAtFastCheckRAY(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.CumulativeDropAtFastCheckRAY(&_CreditManagerv2.CallOpts, arg0)
}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) EnabledTokensMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "enabledTokensMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) EnabledTokensMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.EnabledTokensMap(&_CreditManagerv2.CallOpts, arg0)
}

// EnabledTokensMap is a free data retrieval call binding the contract method 0x8991b2f1.
//
// Solidity: function enabledTokensMap(address ) view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) EnabledTokensMap(arg0 common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.EnabledTokensMap(&_CreditManagerv2.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationDiscount, uint16 feeLiquidationExpired, uint16 liquidationDiscountExpired)
func (_CreditManagerv2 *CreditManagerv2Caller) Fees(opts *bind.CallOpts) (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "fees")

	outstruct := new(struct {
		FeeInterest                uint16
		FeeLiquidation             uint16
		LiquidationDiscount        uint16
		FeeLiquidationExpired      uint16
		LiquidationDiscountExpired uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeInterest = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.FeeLiquidation = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.LiquidationDiscount = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.FeeLiquidationExpired = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.LiquidationDiscountExpired = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationDiscount, uint16 feeLiquidationExpired, uint16 liquidationDiscountExpired)
func (_CreditManagerv2 *CreditManagerv2Session) Fees() (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	return _CreditManagerv2.Contract.Fees(&_CreditManagerv2.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint16 feeInterest, uint16 feeLiquidation, uint16 liquidationDiscount, uint16 feeLiquidationExpired, uint16 liquidationDiscountExpired)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Fees() (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	return _CreditManagerv2.Contract.Fees(&_CreditManagerv2.CallOpts)
}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) ForbiddenTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "forbiddenTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) ForbiddenTokenMask() (*big.Int, error) {
	return _CreditManagerv2.Contract.ForbiddenTokenMask(&_CreditManagerv2.CallOpts)
}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) ForbiddenTokenMask() (*big.Int, error) {
	return _CreditManagerv2.Contract.ForbiddenTokenMask(&_CreditManagerv2.CallOpts)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2Caller) GetCreditAccountOrRevert(opts *bind.CallOpts, borrower common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "getCreditAccountOrRevert", borrower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2Session) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.GetCreditAccountOrRevert(&_CreditManagerv2.CallOpts, borrower)
}

// GetCreditAccountOrRevert is a free data retrieval call binding the contract method 0xe958b704.
//
// Solidity: function getCreditAccountOrRevert(address borrower) view returns(address result)
func (_CreditManagerv2 *CreditManagerv2CallerSession) GetCreditAccountOrRevert(borrower common.Address) (common.Address, error) {
	return _CreditManagerv2.Contract.GetCreditAccountOrRevert(&_CreditManagerv2.CallOpts, borrower)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv2 *CreditManagerv2Caller) LiquidationThresholds(opts *bind.CallOpts, token common.Address) (uint16, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "liquidationThresholds", token)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv2 *CreditManagerv2Session) LiquidationThresholds(token common.Address) (uint16, error) {
	return _CreditManagerv2.Contract.LiquidationThresholds(&_CreditManagerv2.CallOpts, token)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv2 *CreditManagerv2CallerSession) LiquidationThresholds(token common.Address) (uint16, error) {
	return _CreditManagerv2.Contract.LiquidationThresholds(&_CreditManagerv2.CallOpts, token)
}

// MaxAllowedEnabledTokenLength is a free data retrieval call binding the contract method 0x8345f26e.
//
// Solidity: function maxAllowedEnabledTokenLength() view returns(uint8)
func (_CreditManagerv2 *CreditManagerv2Caller) MaxAllowedEnabledTokenLength(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "maxAllowedEnabledTokenLength")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxAllowedEnabledTokenLength is a free data retrieval call binding the contract method 0x8345f26e.
//
// Solidity: function maxAllowedEnabledTokenLength() view returns(uint8)
func (_CreditManagerv2 *CreditManagerv2Session) MaxAllowedEnabledTokenLength() (uint8, error) {
	return _CreditManagerv2.Contract.MaxAllowedEnabledTokenLength(&_CreditManagerv2.CallOpts)
}

// MaxAllowedEnabledTokenLength is a free data retrieval call binding the contract method 0x8345f26e.
//
// Solidity: function maxAllowedEnabledTokenLength() view returns(uint8)
func (_CreditManagerv2 *CreditManagerv2CallerSession) MaxAllowedEnabledTokenLength() (uint8, error) {
	return _CreditManagerv2.Contract.MaxAllowedEnabledTokenLength(&_CreditManagerv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2Session) Paused() (bool, error) {
	return _CreditManagerv2.Contract.Paused(&_CreditManagerv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Paused() (bool, error) {
	return _CreditManagerv2.Contract.Paused(&_CreditManagerv2.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) Pool() (common.Address, error) {
	return _CreditManagerv2.Contract.Pool(&_CreditManagerv2.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Pool() (common.Address, error) {
	return _CreditManagerv2.Contract.Pool(&_CreditManagerv2.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) PoolService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "poolService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) PoolService() (common.Address, error) {
	return _CreditManagerv2.Contract.PoolService(&_CreditManagerv2.CallOpts)
}

// PoolService is a free data retrieval call binding the contract method 0x570a7af2.
//
// Solidity: function poolService() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) PoolService() (common.Address, error) {
	return _CreditManagerv2.Contract.PoolService(&_CreditManagerv2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) PriceOracle() (common.Address, error) {
	return _CreditManagerv2.Contract.PriceOracle(&_CreditManagerv2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) PriceOracle() (common.Address, error) {
	return _CreditManagerv2.Contract.PriceOracle(&_CreditManagerv2.CallOpts)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address token) view returns(uint256 mask)
func (_CreditManagerv2 *CreditManagerv2Caller) TokenMasksMap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "tokenMasksMap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address token) view returns(uint256 mask)
func (_CreditManagerv2 *CreditManagerv2Session) TokenMasksMap(token common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.TokenMasksMap(&_CreditManagerv2.CallOpts, token)
}

// TokenMasksMap is a free data retrieval call binding the contract method 0xf67c5bd0.
//
// Solidity: function tokenMasksMap(address token) view returns(uint256 mask)
func (_CreditManagerv2 *CreditManagerv2CallerSession) TokenMasksMap(token common.Address) (*big.Int, error) {
	return _CreditManagerv2.Contract.TokenMasksMap(&_CreditManagerv2.CallOpts, token)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) Underlying() (common.Address, error) {
	return _CreditManagerv2.Contract.Underlying(&_CreditManagerv2.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Underlying() (common.Address, error) {
	return _CreditManagerv2.Contract.Underlying(&_CreditManagerv2.CallOpts)
}

// UniversalAdapter is a free data retrieval call binding the contract method 0xfe47cde7.
//
// Solidity: function universalAdapter() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) UniversalAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "universalAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniversalAdapter is a free data retrieval call binding the contract method 0xfe47cde7.
//
// Solidity: function universalAdapter() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) UniversalAdapter() (common.Address, error) {
	return _CreditManagerv2.Contract.UniversalAdapter(&_CreditManagerv2.CallOpts)
}

// UniversalAdapter is a free data retrieval call binding the contract method 0xfe47cde7.
//
// Solidity: function universalAdapter() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) UniversalAdapter() (common.Address, error) {
	return _CreditManagerv2.Contract.UniversalAdapter(&_CreditManagerv2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2Session) Version() (*big.Int, error) {
	return _CreditManagerv2.Contract.Version(&_CreditManagerv2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv2 *CreditManagerv2CallerSession) Version() (*big.Int, error) {
	return _CreditManagerv2.Contract.Version(&_CreditManagerv2.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) WethAddress() (common.Address, error) {
	return _CreditManagerv2.Contract.WethAddress(&_CreditManagerv2.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) WethAddress() (common.Address, error) {
	return _CreditManagerv2.Contract.WethAddress(&_CreditManagerv2.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Caller) WethGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv2.contract.Call(opts, &out, "wethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) WethGateway() (common.Address, error) {
	return _CreditManagerv2.Contract.WethGateway(&_CreditManagerv2.CallOpts)
}

// WethGateway is a free data retrieval call binding the contract method 0xc5e10eef.
//
// Solidity: function wethGateway() view returns(address)
func (_CreditManagerv2 *CreditManagerv2CallerSession) WethGateway() (common.Address, error) {
	return _CreditManagerv2.Contract.WethGateway(&_CreditManagerv2.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) AddCollateral(opts *bind.TransactOpts, payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "addCollateral", payer, creditAccount, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) AddCollateral(payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddCollateral(&_CreditManagerv2.TransactOpts, payer, creditAccount, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) AddCollateral(payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddCollateral(&_CreditManagerv2.TransactOpts, payer, creditAccount, token, amount)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) AddEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "addEmergencyLiquidator", liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2Session) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddEmergencyLiquidator(&_CreditManagerv2.TransactOpts, liquidator)
}

// AddEmergencyLiquidator is a paid mutator transaction binding the contract method 0x84edaa42.
//
// Solidity: function addEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) AddEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddEmergencyLiquidator(&_CreditManagerv2.TransactOpts, liquidator)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2Session) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddToken(&_CreditManagerv2.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.AddToken(&_CreditManagerv2.TransactOpts, token)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x46fb371d.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) ApproveCreditAccount(opts *bind.TransactOpts, borrower common.Address, targetContract common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "approveCreditAccount", borrower, targetContract, token, amount)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x46fb371d.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) ApproveCreditAccount(borrower common.Address, targetContract common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ApproveCreditAccount(&_CreditManagerv2.TransactOpts, borrower, targetContract, token, amount)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0x46fb371d.
//
// Solidity: function approveCreditAccount(address borrower, address targetContract, address token, uint256 amount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ApproveCreditAccount(borrower common.Address, targetContract common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ApproveCreditAccount(&_CreditManagerv2.TransactOpts, borrower, targetContract, token, amount)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) ChangeContractAllowance(opts *bind.TransactOpts, adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "changeContractAllowance", adapter, targetContract)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2Session) ChangeContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ChangeContractAllowance(&_CreditManagerv2.TransactOpts, adapter, targetContract)
}

// ChangeContractAllowance is a paid mutator transaction binding the contract method 0x6e98e5e4.
//
// Solidity: function changeContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ChangeContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ChangeContractAllowance(&_CreditManagerv2.TransactOpts, adapter, targetContract)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) CheckAndEnableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "checkAndEnableToken", creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Session) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndEnableToken(&_CreditManagerv2.TransactOpts, creditAccount, token)
}

// CheckAndEnableToken is a paid mutator transaction binding the contract method 0x51e3f160.
//
// Solidity: function checkAndEnableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) CheckAndEnableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndEnableToken(&_CreditManagerv2.TransactOpts, creditAccount, token)
}

// CheckAndOptimizeEnabledTokens is a paid mutator transaction binding the contract method 0x29df0b93.
//
// Solidity: function checkAndOptimizeEnabledTokens(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) CheckAndOptimizeEnabledTokens(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "checkAndOptimizeEnabledTokens", creditAccount)
}

// CheckAndOptimizeEnabledTokens is a paid mutator transaction binding the contract method 0x29df0b93.
//
// Solidity: function checkAndOptimizeEnabledTokens(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) CheckAndOptimizeEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndOptimizeEnabledTokens(&_CreditManagerv2.TransactOpts, creditAccount)
}

// CheckAndOptimizeEnabledTokens is a paid mutator transaction binding the contract method 0x29df0b93.
//
// Solidity: function checkAndOptimizeEnabledTokens(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) CheckAndOptimizeEnabledTokens(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CheckAndOptimizeEnabledTokens(&_CreditManagerv2.TransactOpts, creditAccount)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x2362a2d8.
//
// Solidity: function closeCreditAccount(address borrower, uint8 closureActionType, uint256 totalValue, address payer, address to, uint256 skipTokenMask, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2Transactor) CloseCreditAccount(opts *bind.TransactOpts, borrower common.Address, closureActionType uint8, totalValue *big.Int, payer common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "closeCreditAccount", borrower, closureActionType, totalValue, payer, to, skipTokenMask, convertWETH)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x2362a2d8.
//
// Solidity: function closeCreditAccount(address borrower, uint8 closureActionType, uint256 totalValue, address payer, address to, uint256 skipTokenMask, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2Session) CloseCreditAccount(borrower common.Address, closureActionType uint8, totalValue *big.Int, payer common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CloseCreditAccount(&_CreditManagerv2.TransactOpts, borrower, closureActionType, totalValue, payer, to, skipTokenMask, convertWETH)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x2362a2d8.
//
// Solidity: function closeCreditAccount(address borrower, uint8 closureActionType, uint256 totalValue, address payer, address to, uint256 skipTokenMask, bool convertWETH) returns(uint256 remainingFunds)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) CloseCreditAccount(borrower common.Address, closureActionType uint8, totalValue *big.Int, payer common.Address, to common.Address, skipTokenMask *big.Int, convertWETH bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.CloseCreditAccount(&_CreditManagerv2.TransactOpts, borrower, closureActionType, totalValue, payer, to, skipTokenMask, convertWETH)
}

// DisableToken is a paid mutator transaction binding the contract method 0x0d8f9cee.
//
// Solidity: function disableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) DisableToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "disableToken", creditAccount, token)
}

// DisableToken is a paid mutator transaction binding the contract method 0x0d8f9cee.
//
// Solidity: function disableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2Session) DisableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.DisableToken(&_CreditManagerv2.TransactOpts, creditAccount, token)
}

// DisableToken is a paid mutator transaction binding the contract method 0x0d8f9cee.
//
// Solidity: function disableToken(address creditAccount, address token) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) DisableToken(creditAccount common.Address, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.DisableToken(&_CreditManagerv2.TransactOpts, creditAccount, token)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2Transactor) ExecuteOrder(opts *bind.TransactOpts, borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "executeOrder", borrower, targetContract, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2Session) ExecuteOrder(borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ExecuteOrder(&_CreditManagerv2.TransactOpts, borrower, targetContract, data)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0x6ce4074a.
//
// Solidity: function executeOrder(address borrower, address targetContract, bytes data) returns(bytes)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ExecuteOrder(borrower common.Address, targetContract common.Address, data []byte) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ExecuteOrder(&_CreditManagerv2.TransactOpts, borrower, targetContract, data)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) FastCollateralCheck(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "fastCollateralCheck", creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2Session) FastCollateralCheck(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FastCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FastCollateralCheck is a paid mutator transaction binding the contract method 0x654a9eda.
//
// Solidity: function fastCollateralCheck(address creditAccount, address tokenIn, address tokenOut, uint256 balanceInBefore, uint256 balanceOutBefore) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) FastCollateralCheck(creditAccount common.Address, tokenIn common.Address, tokenOut common.Address, balanceInBefore *big.Int, balanceOutBefore *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FastCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount, tokenIn, tokenOut, balanceInBefore, balanceOutBefore)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) FullCollateralCheck(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "fullCollateralCheck", creditAccount)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2Session) FullCollateralCheck(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FullCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x95373018.
//
// Solidity: function fullCollateralCheck(address creditAccount) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) FullCollateralCheck(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.FullCollateralCheck(&_CreditManagerv2.TransactOpts, creditAccount)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2Transactor) ManageDebt(opts *bind.TransactOpts, creditAccount common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "manageDebt", creditAccount, amount, increase)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2Session) ManageDebt(creditAccount common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ManageDebt(&_CreditManagerv2.TransactOpts, creditAccount, amount, increase)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x94cf073a.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, bool increase) returns(uint256 newBorrowedAmount)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) ManageDebt(creditAccount common.Address, amount *big.Int, increase bool) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.ManageDebt(&_CreditManagerv2.TransactOpts, creditAccount, amount, increase)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2Transactor) OpenCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "openCreditAccount", borrowedAmount, onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2Session) OpenCreditAccount(borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.OpenCreditAccount(&_CreditManagerv2.TransactOpts, borrowedAmount, onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x8fe3f93f.
//
// Solidity: function openCreditAccount(uint256 borrowedAmount, address onBehalfOf) returns(address)
func (_CreditManagerv2 *CreditManagerv2TransactorSession) OpenCreditAccount(borrowedAmount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.OpenCreditAccount(&_CreditManagerv2.TransactOpts, borrowedAmount, onBehalfOf)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2Session) Pause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Pause(&_CreditManagerv2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) Pause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Pause(&_CreditManagerv2.TransactOpts)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) RemoveEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "removeEmergencyLiquidator", liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2Session) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.RemoveEmergencyLiquidator(&_CreditManagerv2.TransactOpts, liquidator)
}

// RemoveEmergencyLiquidator is a paid mutator transaction binding the contract method 0xa460e104.
//
// Solidity: function removeEmergencyLiquidator(address liquidator) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) RemoveEmergencyLiquidator(liquidator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.RemoveEmergencyLiquidator(&_CreditManagerv2.TransactOpts, liquidator)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetConfigurator(opts *bind.TransactOpts, _creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setConfigurator", _creditConfigurator)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetConfigurator(&_CreditManagerv2.TransactOpts, _creditConfigurator)
}

// SetConfigurator is a paid mutator transaction binding the contract method 0x9f5f86ae.
//
// Solidity: function setConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetConfigurator(&_CreditManagerv2.TransactOpts, _creditConfigurator)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetForbidMask(opts *bind.TransactOpts, _forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setForbidMask", _forbidMask)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetForbidMask(_forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetForbidMask(&_CreditManagerv2.TransactOpts, _forbidMask)
}

// SetForbidMask is a paid mutator transaction binding the contract method 0xa366f496.
//
// Solidity: function setForbidMask(uint256 _forbidMask) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetForbidMask(_forbidMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetForbidMask(&_CreditManagerv2.TransactOpts, _forbidMask)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetLiquidationThreshold(opts *bind.TransactOpts, token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setLiquidationThreshold", token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetLiquidationThreshold(&_CreditManagerv2.TransactOpts, token, liquidationThreshold)
}

// SetLiquidationThreshold is a paid mutator transaction binding the contract method 0xa70bc542.
//
// Solidity: function setLiquidationThreshold(address token, uint16 liquidationThreshold) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetLiquidationThreshold(token common.Address, liquidationThreshold uint16) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetLiquidationThreshold(&_CreditManagerv2.TransactOpts, token, liquidationThreshold)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetMaxEnabledTokens(opts *bind.TransactOpts, newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setMaxEnabledTokens", newMaxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetMaxEnabledTokens(newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetMaxEnabledTokens(&_CreditManagerv2.TransactOpts, newMaxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 newMaxEnabledTokens) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetMaxEnabledTokens(newMaxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetMaxEnabledTokens(&_CreditManagerv2.TransactOpts, newMaxEnabledTokens)
}

// SetParams is a paid mutator transaction binding the contract method 0x944ac59f.
//
// Solidity: function setParams(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) SetParams(opts *bind.TransactOpts, _feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "setParams", _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// SetParams is a paid mutator transaction binding the contract method 0x944ac59f.
//
// Solidity: function setParams(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv2 *CreditManagerv2Session) SetParams(_feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetParams(&_CreditManagerv2.TransactOpts, _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// SetParams is a paid mutator transaction binding the contract method 0x944ac59f.
//
// Solidity: function setParams(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) SetParams(_feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.SetParams(&_CreditManagerv2.TransactOpts, _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) TransferAccountOwnership(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "transferAccountOwnership", from, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2Session) TransferAccountOwnership(from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.TransferAccountOwnership(&_CreditManagerv2.TransactOpts, from, to)
}

// TransferAccountOwnership is a paid mutator transaction binding the contract method 0xe1998cf9.
//
// Solidity: function transferAccountOwnership(address from, address to) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) TransferAccountOwnership(from common.Address, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.TransferAccountOwnership(&_CreditManagerv2.TransactOpts, from, to)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2Session) Unpause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Unpause(&_CreditManagerv2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditManagerv2.Contract.Unpause(&_CreditManagerv2.TransactOpts)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) UpgradeCreditFacade(opts *bind.TransactOpts, _creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "upgradeCreditFacade", _creditFacade)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditManagerv2 *CreditManagerv2Session) UpgradeCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradeCreditFacade(&_CreditManagerv2.TransactOpts, _creditFacade)
}

// UpgradeCreditFacade is a paid mutator transaction binding the contract method 0x693ce7f5.
//
// Solidity: function upgradeCreditFacade(address _creditFacade) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) UpgradeCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradeCreditFacade(&_CreditManagerv2.TransactOpts, _creditFacade)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xdc9e0faa.
//
// Solidity: function upgradePriceOracle(address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2Transactor) UpgradePriceOracle(opts *bind.TransactOpts, _priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.contract.Transact(opts, "upgradePriceOracle", _priceOracle)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xdc9e0faa.
//
// Solidity: function upgradePriceOracle(address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2Session) UpgradePriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradePriceOracle(&_CreditManagerv2.TransactOpts, _priceOracle)
}

// UpgradePriceOracle is a paid mutator transaction binding the contract method 0xdc9e0faa.
//
// Solidity: function upgradePriceOracle(address _priceOracle) returns()
func (_CreditManagerv2 *CreditManagerv2TransactorSession) UpgradePriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv2.Contract.UpgradePriceOracle(&_CreditManagerv2.TransactOpts, _priceOracle)
}

// CreditManagerv2ExecuteOrderIterator is returned from FilterExecuteOrder and is used to iterate over the raw logs and unpacked data for ExecuteOrder events raised by the CreditManagerv2 contract.
type CreditManagerv2ExecuteOrderIterator struct {
	Event *CreditManagerv2ExecuteOrder // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2ExecuteOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2ExecuteOrder)
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
		it.Event = new(CreditManagerv2ExecuteOrder)
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
func (it *CreditManagerv2ExecuteOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2ExecuteOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2ExecuteOrder represents a ExecuteOrder event raised by the CreditManagerv2 contract.
type CreditManagerv2ExecuteOrder struct {
	Borrower common.Address
	Target   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecuteOrder is a free log retrieval operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterExecuteOrder(opts *bind.FilterOpts, borrower []common.Address, target []common.Address) (*CreditManagerv2ExecuteOrderIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2ExecuteOrderIterator{contract: _CreditManagerv2.contract, event: "ExecuteOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteOrder is a free log subscription operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchExecuteOrder(opts *bind.WatchOpts, sink chan<- *CreditManagerv2ExecuteOrder, borrower []common.Address, target []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "ExecuteOrder", borrowerRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2ExecuteOrder)
				if err := _CreditManagerv2.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
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

// ParseExecuteOrder is a log parse operation binding the contract event 0xaed1eb34af6acd8c1e3911fb2ebb875a66324b03957886bd002227b17f52ab03.
//
// Solidity: event ExecuteOrder(address indexed borrower, address indexed target)
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseExecuteOrder(log types.Log) (*CreditManagerv2ExecuteOrder, error) {
	event := new(CreditManagerv2ExecuteOrder)
	if err := _CreditManagerv2.contract.UnpackLog(event, "ExecuteOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditManagerv2NewConfiguratorIterator is returned from FilterNewConfigurator and is used to iterate over the raw logs and unpacked data for NewConfigurator events raised by the CreditManagerv2 contract.
type CreditManagerv2NewConfiguratorIterator struct {
	Event *CreditManagerv2NewConfigurator // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2NewConfiguratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2NewConfigurator)
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
		it.Event = new(CreditManagerv2NewConfigurator)
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
func (it *CreditManagerv2NewConfiguratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2NewConfiguratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2NewConfigurator represents a NewConfigurator event raised by the CreditManagerv2 contract.
type CreditManagerv2NewConfigurator struct {
	NewConfigurator common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewConfigurator is a free log retrieval operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterNewConfigurator(opts *bind.FilterOpts, newConfigurator []common.Address) (*CreditManagerv2NewConfiguratorIterator, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "NewConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2NewConfiguratorIterator{contract: _CreditManagerv2.contract, event: "NewConfigurator", logs: logs, sub: sub}, nil
}

// WatchNewConfigurator is a free log subscription operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchNewConfigurator(opts *bind.WatchOpts, sink chan<- *CreditManagerv2NewConfigurator, newConfigurator []common.Address) (event.Subscription, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "NewConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2NewConfigurator)
				if err := _CreditManagerv2.contract.UnpackLog(event, "NewConfigurator", log); err != nil {
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

// ParseNewConfigurator is a log parse operation binding the contract event 0xf62005acebe9b616aefb5f248b48f5e89f28437b27d1eebc0b2d911209f297af.
//
// Solidity: event NewConfigurator(address indexed newConfigurator)
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseNewConfigurator(log types.Log) (*CreditManagerv2NewConfigurator, error) {
	event := new(CreditManagerv2NewConfigurator)
	if err := _CreditManagerv2.contract.UnpackLog(event, "NewConfigurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditManagerv2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditManagerv2 contract.
type CreditManagerv2PausedIterator struct {
	Event *CreditManagerv2Paused // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2Paused)
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
		it.Event = new(CreditManagerv2Paused)
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
func (it *CreditManagerv2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2Paused represents a Paused event raised by the CreditManagerv2 contract.
type CreditManagerv2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterPaused(opts *bind.FilterOpts) (*CreditManagerv2PausedIterator, error) {

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2PausedIterator{contract: _CreditManagerv2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditManagerv2Paused) (event.Subscription, error) {

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2Paused)
				if err := _CreditManagerv2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditManagerv2 *CreditManagerv2Filterer) ParsePaused(log types.Log) (*CreditManagerv2Paused, error) {
	event := new(CreditManagerv2Paused)
	if err := _CreditManagerv2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditManagerv2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditManagerv2 contract.
type CreditManagerv2UnpausedIterator struct {
	Event *CreditManagerv2Unpaused // Event containing the contract specifics and raw log

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
func (it *CreditManagerv2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv2Unpaused)
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
		it.Event = new(CreditManagerv2Unpaused)
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
func (it *CreditManagerv2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv2Unpaused represents a Unpaused event raised by the CreditManagerv2 contract.
type CreditManagerv2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditManagerv2UnpausedIterator, error) {

	logs, sub, err := _CreditManagerv2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditManagerv2UnpausedIterator{contract: _CreditManagerv2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditManagerv2 *CreditManagerv2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditManagerv2Unpaused) (event.Subscription, error) {

	logs, sub, err := _CreditManagerv2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv2Unpaused)
				if err := _CreditManagerv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditManagerv2 *CreditManagerv2Filterer) ParseUnpaused(log types.Log) (*CreditManagerv2Unpaused, error) {
	event := new(CreditManagerv2Unpaused)
	if err := _CreditManagerv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
