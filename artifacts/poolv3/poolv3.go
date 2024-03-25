// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolv3

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

// Poolv3MetaData contains all meta data concerning the Poolv3 contract.
var Poolv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"supportsQuotas_\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotControllerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotCreditManagerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPoolQuotaKeeperException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnpausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditManagerCantBorrowException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleCreditManagerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatiblePoolQuotaKeeperException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuotasNotSupportedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredCreditManagerOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"AddCreditManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"IncurUncoveredLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"SetCreditManagerDebtLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"SetInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPoolQuotaKeeper\",\"type\":\"address\"}],\"name\":\"SetPoolQuotaKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"SetTotalDebtLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseInterestIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseInterestIndexLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcLinearCumulative_RAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"creditManagerBorrowable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowable\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"creditManagerBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"creditManagerDebtLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManagers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"depositWithReferral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expectedLiquidityLU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBaseInterestUpdate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastQuotaRevenueUpdate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"lendCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"mintWithReferral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolQuotaKeeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quotaRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repaidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"}],\"name\":\"repayCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setCreditManagerDebtLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolQuotaKeeper\",\"type\":\"address\"}],\"name\":\"setPoolQuotaKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuotaRevenue\",\"type\":\"uint256\"}],\"name\":\"setQuotaRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setTotalDebtLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newWithdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supportsQuotas\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebtLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"quotaRevenueDelta\",\"type\":\"int256\"}],\"name\":\"updateQuotaRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Poolv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Poolv3MetaData.ABI instead.
var Poolv3ABI = Poolv3MetaData.ABI

// Poolv3 is an auto generated Go binding around an Ethereum contract.
type Poolv3 struct {
	Poolv3Caller     // Read-only binding to the contract
	Poolv3Transactor // Write-only binding to the contract
	Poolv3Filterer   // Log filterer for contract events
}

// Poolv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Poolv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Poolv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Poolv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Poolv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Poolv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Poolv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Poolv3Session struct {
	Contract     *Poolv3           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Poolv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Poolv3CallerSession struct {
	Contract *Poolv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Poolv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Poolv3TransactorSession struct {
	Contract     *Poolv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Poolv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Poolv3Raw struct {
	Contract *Poolv3 // Generic contract binding to access the raw methods on
}

// Poolv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Poolv3CallerRaw struct {
	Contract *Poolv3Caller // Generic read-only contract binding to access the raw methods on
}

// Poolv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Poolv3TransactorRaw struct {
	Contract *Poolv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolv3 creates a new instance of Poolv3, bound to a specific deployed contract.
func NewPoolv3(address common.Address, backend bind.ContractBackend) (*Poolv3, error) {
	contract, err := bindPoolv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolv3{Poolv3Caller: Poolv3Caller{contract: contract}, Poolv3Transactor: Poolv3Transactor{contract: contract}, Poolv3Filterer: Poolv3Filterer{contract: contract}}, nil
}

// NewPoolv3Caller creates a new read-only instance of Poolv3, bound to a specific deployed contract.
func NewPoolv3Caller(address common.Address, caller bind.ContractCaller) (*Poolv3Caller, error) {
	contract, err := bindPoolv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Poolv3Caller{contract: contract}, nil
}

// NewPoolv3Transactor creates a new write-only instance of Poolv3, bound to a specific deployed contract.
func NewPoolv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Poolv3Transactor, error) {
	contract, err := bindPoolv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Poolv3Transactor{contract: contract}, nil
}

// NewPoolv3Filterer creates a new log filterer instance of Poolv3, bound to a specific deployed contract.
func NewPoolv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Poolv3Filterer, error) {
	contract, err := bindPoolv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Poolv3Filterer{contract: contract}, nil
}

// bindPoolv3 binds a generic wrapper to an already deployed contract.
func bindPoolv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Poolv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolv3 *Poolv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolv3.Contract.Poolv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolv3 *Poolv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolv3.Contract.Poolv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolv3 *Poolv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolv3.Contract.Poolv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolv3 *Poolv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolv3 *Poolv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolv3 *Poolv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolv3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Poolv3 *Poolv3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Poolv3 *Poolv3Session) Acl() (common.Address, error) {
	return _Poolv3.Contract.Acl(&_Poolv3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Poolv3 *Poolv3CallerSession) Acl() (common.Address, error) {
	return _Poolv3.Contract.Acl(&_Poolv3.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Poolv3 *Poolv3Caller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Poolv3 *Poolv3Session) AddressProvider() (common.Address, error) {
	return _Poolv3.Contract.AddressProvider(&_Poolv3.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Poolv3 *Poolv3CallerSession) AddressProvider() (common.Address, error) {
	return _Poolv3.Contract.AddressProvider(&_Poolv3.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Poolv3 *Poolv3Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Poolv3 *Poolv3Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Poolv3.Contract.Allowance(&_Poolv3.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Poolv3.Contract.Allowance(&_Poolv3.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Poolv3 *Poolv3Caller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Poolv3 *Poolv3Session) Asset() (common.Address, error) {
	return _Poolv3.Contract.Asset(&_Poolv3.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Poolv3 *Poolv3CallerSession) Asset() (common.Address, error) {
	return _Poolv3.Contract.Asset(&_Poolv3.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3Caller) AvailableLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "availableLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3Session) AvailableLiquidity() (*big.Int, error) {
	return _Poolv3.Contract.AvailableLiquidity(&_Poolv3.CallOpts)
}

// AvailableLiquidity is a free data retrieval call binding the contract method 0x74375359.
//
// Solidity: function availableLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) AvailableLiquidity() (*big.Int, error) {
	return _Poolv3.Contract.AvailableLiquidity(&_Poolv3.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Poolv3 *Poolv3Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Poolv3 *Poolv3Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Poolv3.Contract.BalanceOf(&_Poolv3.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Poolv3.Contract.BalanceOf(&_Poolv3.CallOpts, account)
}

// BaseInterestIndex is a free data retrieval call binding the contract method 0xfaaba9e2.
//
// Solidity: function baseInterestIndex() view returns(uint256)
func (_Poolv3 *Poolv3Caller) BaseInterestIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "baseInterestIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseInterestIndex is a free data retrieval call binding the contract method 0xfaaba9e2.
//
// Solidity: function baseInterestIndex() view returns(uint256)
func (_Poolv3 *Poolv3Session) BaseInterestIndex() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestIndex(&_Poolv3.CallOpts)
}

// BaseInterestIndex is a free data retrieval call binding the contract method 0xfaaba9e2.
//
// Solidity: function baseInterestIndex() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) BaseInterestIndex() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestIndex(&_Poolv3.CallOpts)
}

// BaseInterestIndexLU is a free data retrieval call binding the contract method 0xa74d4910.
//
// Solidity: function baseInterestIndexLU() view returns(uint256)
func (_Poolv3 *Poolv3Caller) BaseInterestIndexLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "baseInterestIndexLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseInterestIndexLU is a free data retrieval call binding the contract method 0xa74d4910.
//
// Solidity: function baseInterestIndexLU() view returns(uint256)
func (_Poolv3 *Poolv3Session) BaseInterestIndexLU() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestIndexLU(&_Poolv3.CallOpts)
}

// BaseInterestIndexLU is a free data retrieval call binding the contract method 0xa74d4910.
//
// Solidity: function baseInterestIndexLU() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) BaseInterestIndexLU() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestIndexLU(&_Poolv3.CallOpts)
}

// BaseInterestRate is a free data retrieval call binding the contract method 0xafd92762.
//
// Solidity: function baseInterestRate() view returns(uint256)
func (_Poolv3 *Poolv3Caller) BaseInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "baseInterestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseInterestRate is a free data retrieval call binding the contract method 0xafd92762.
//
// Solidity: function baseInterestRate() view returns(uint256)
func (_Poolv3 *Poolv3Session) BaseInterestRate() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestRate(&_Poolv3.CallOpts)
}

// BaseInterestRate is a free data retrieval call binding the contract method 0xafd92762.
//
// Solidity: function baseInterestRate() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) BaseInterestRate() (*big.Int, error) {
	return _Poolv3.Contract.BaseInterestRate(&_Poolv3.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_Poolv3 *Poolv3Caller) CalcLinearCumulativeRAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "calcLinearCumulative_RAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_Poolv3 *Poolv3Session) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _Poolv3.Contract.CalcLinearCumulativeRAY(&_Poolv3.CallOpts)
}

// CalcLinearCumulativeRAY is a free data retrieval call binding the contract method 0x0fce70fb.
//
// Solidity: function calcLinearCumulative_RAY() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) CalcLinearCumulativeRAY() (*big.Int, error) {
	return _Poolv3.Contract.CalcLinearCumulativeRAY(&_Poolv3.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_Poolv3 *Poolv3Caller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_Poolv3 *Poolv3Session) ContractsRegister() (common.Address, error) {
	return _Poolv3.Contract.ContractsRegister(&_Poolv3.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_Poolv3 *Poolv3CallerSession) ContractsRegister() (common.Address, error) {
	return _Poolv3.Contract.ContractsRegister(&_Poolv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Poolv3 *Poolv3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Poolv3 *Poolv3Session) Controller() (common.Address, error) {
	return _Poolv3.Contract.Controller(&_Poolv3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Poolv3 *Poolv3CallerSession) Controller() (common.Address, error) {
	return _Poolv3.Contract.Controller(&_Poolv3.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Poolv3 *Poolv3Caller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Poolv3 *Poolv3Session) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.ConvertToAssets(&_Poolv3.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256 assets)
func (_Poolv3 *Poolv3CallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.ConvertToAssets(&_Poolv3.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3Caller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3Session) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.ConvertToShares(&_Poolv3.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3CallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.ConvertToShares(&_Poolv3.CallOpts, assets)
}

// CreditManagerBorrowable is a free data retrieval call binding the contract method 0x136a6833.
//
// Solidity: function creditManagerBorrowable(address creditManager) view returns(uint256 borrowable)
func (_Poolv3 *Poolv3Caller) CreditManagerBorrowable(opts *bind.CallOpts, creditManager common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "creditManagerBorrowable", creditManager)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagerBorrowable is a free data retrieval call binding the contract method 0x136a6833.
//
// Solidity: function creditManagerBorrowable(address creditManager) view returns(uint256 borrowable)
func (_Poolv3 *Poolv3Session) CreditManagerBorrowable(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerBorrowable(&_Poolv3.CallOpts, creditManager)
}

// CreditManagerBorrowable is a free data retrieval call binding the contract method 0x136a6833.
//
// Solidity: function creditManagerBorrowable(address creditManager) view returns(uint256 borrowable)
func (_Poolv3 *Poolv3CallerSession) CreditManagerBorrowable(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerBorrowable(&_Poolv3.CallOpts, creditManager)
}

// CreditManagerBorrowed is a free data retrieval call binding the contract method 0x7a99c017.
//
// Solidity: function creditManagerBorrowed(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3Caller) CreditManagerBorrowed(opts *bind.CallOpts, creditManager common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "creditManagerBorrowed", creditManager)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagerBorrowed is a free data retrieval call binding the contract method 0x7a99c017.
//
// Solidity: function creditManagerBorrowed(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3Session) CreditManagerBorrowed(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerBorrowed(&_Poolv3.CallOpts, creditManager)
}

// CreditManagerBorrowed is a free data retrieval call binding the contract method 0x7a99c017.
//
// Solidity: function creditManagerBorrowed(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) CreditManagerBorrowed(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerBorrowed(&_Poolv3.CallOpts, creditManager)
}

// CreditManagerDebtLimit is a free data retrieval call binding the contract method 0xb0df2c66.
//
// Solidity: function creditManagerDebtLimit(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3Caller) CreditManagerDebtLimit(opts *bind.CallOpts, creditManager common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "creditManagerDebtLimit", creditManager)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditManagerDebtLimit is a free data retrieval call binding the contract method 0xb0df2c66.
//
// Solidity: function creditManagerDebtLimit(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3Session) CreditManagerDebtLimit(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerDebtLimit(&_Poolv3.CallOpts, creditManager)
}

// CreditManagerDebtLimit is a free data retrieval call binding the contract method 0xb0df2c66.
//
// Solidity: function creditManagerDebtLimit(address creditManager) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) CreditManagerDebtLimit(creditManager common.Address) (*big.Int, error) {
	return _Poolv3.Contract.CreditManagerDebtLimit(&_Poolv3.CallOpts, creditManager)
}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_Poolv3 *Poolv3Caller) CreditManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "creditManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_Poolv3 *Poolv3Session) CreditManagers() ([]common.Address, error) {
	return _Poolv3.Contract.CreditManagers(&_Poolv3.CallOpts)
}

// CreditManagers is a free data retrieval call binding the contract method 0xdac54431.
//
// Solidity: function creditManagers() view returns(address[])
func (_Poolv3 *Poolv3CallerSession) CreditManagers() ([]common.Address, error) {
	return _Poolv3.Contract.CreditManagers(&_Poolv3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolv3 *Poolv3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolv3 *Poolv3Session) Decimals() (uint8, error) {
	return _Poolv3.Contract.Decimals(&_Poolv3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Poolv3 *Poolv3CallerSession) Decimals() (uint8, error) {
	return _Poolv3.Contract.Decimals(&_Poolv3.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3Caller) ExpectedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "expectedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3Session) ExpectedLiquidity() (*big.Int, error) {
	return _Poolv3.Contract.ExpectedLiquidity(&_Poolv3.CallOpts)
}

// ExpectedLiquidity is a free data retrieval call binding the contract method 0xfe14112d.
//
// Solidity: function expectedLiquidity() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) ExpectedLiquidity() (*big.Int, error) {
	return _Poolv3.Contract.ExpectedLiquidity(&_Poolv3.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0xc8c98662.
//
// Solidity: function expectedLiquidityLU() view returns(uint256)
func (_Poolv3 *Poolv3Caller) ExpectedLiquidityLU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "expectedLiquidityLU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0xc8c98662.
//
// Solidity: function expectedLiquidityLU() view returns(uint256)
func (_Poolv3 *Poolv3Session) ExpectedLiquidityLU() (*big.Int, error) {
	return _Poolv3.Contract.ExpectedLiquidityLU(&_Poolv3.CallOpts)
}

// ExpectedLiquidityLU is a free data retrieval call binding the contract method 0xc8c98662.
//
// Solidity: function expectedLiquidityLU() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) ExpectedLiquidityLU() (*big.Int, error) {
	return _Poolv3.Contract.ExpectedLiquidityLU(&_Poolv3.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Poolv3 *Poolv3Caller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Poolv3 *Poolv3Session) InterestRateModel() (common.Address, error) {
	return _Poolv3.Contract.InterestRateModel(&_Poolv3.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Poolv3 *Poolv3CallerSession) InterestRateModel() (common.Address, error) {
	return _Poolv3.Contract.InterestRateModel(&_Poolv3.CallOpts)
}

// LastBaseInterestUpdate is a free data retrieval call binding the contract method 0x6b88245b.
//
// Solidity: function lastBaseInterestUpdate() view returns(uint40)
func (_Poolv3 *Poolv3Caller) LastBaseInterestUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "lastBaseInterestUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBaseInterestUpdate is a free data retrieval call binding the contract method 0x6b88245b.
//
// Solidity: function lastBaseInterestUpdate() view returns(uint40)
func (_Poolv3 *Poolv3Session) LastBaseInterestUpdate() (*big.Int, error) {
	return _Poolv3.Contract.LastBaseInterestUpdate(&_Poolv3.CallOpts)
}

// LastBaseInterestUpdate is a free data retrieval call binding the contract method 0x6b88245b.
//
// Solidity: function lastBaseInterestUpdate() view returns(uint40)
func (_Poolv3 *Poolv3CallerSession) LastBaseInterestUpdate() (*big.Int, error) {
	return _Poolv3.Contract.LastBaseInterestUpdate(&_Poolv3.CallOpts)
}

// LastQuotaRevenueUpdate is a free data retrieval call binding the contract method 0x88ae7842.
//
// Solidity: function lastQuotaRevenueUpdate() view returns(uint40)
func (_Poolv3 *Poolv3Caller) LastQuotaRevenueUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "lastQuotaRevenueUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastQuotaRevenueUpdate is a free data retrieval call binding the contract method 0x88ae7842.
//
// Solidity: function lastQuotaRevenueUpdate() view returns(uint40)
func (_Poolv3 *Poolv3Session) LastQuotaRevenueUpdate() (*big.Int, error) {
	return _Poolv3.Contract.LastQuotaRevenueUpdate(&_Poolv3.CallOpts)
}

// LastQuotaRevenueUpdate is a free data retrieval call binding the contract method 0x88ae7842.
//
// Solidity: function lastQuotaRevenueUpdate() view returns(uint40)
func (_Poolv3 *Poolv3CallerSession) LastQuotaRevenueUpdate() (*big.Int, error) {
	return _Poolv3.Contract.LastQuotaRevenueUpdate(&_Poolv3.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Poolv3 *Poolv3Caller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Poolv3 *Poolv3Session) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxDeposit(&_Poolv3.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxDeposit(&_Poolv3.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Poolv3 *Poolv3Caller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Poolv3 *Poolv3Session) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxMint(&_Poolv3.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxMint(&_Poolv3.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Poolv3 *Poolv3Caller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Poolv3 *Poolv3Session) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxRedeem(&_Poolv3.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxRedeem(&_Poolv3.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Poolv3 *Poolv3Caller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Poolv3 *Poolv3Session) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxWithdraw(&_Poolv3.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Poolv3.Contract.MaxWithdraw(&_Poolv3.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolv3 *Poolv3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolv3 *Poolv3Session) Name() (string, error) {
	return _Poolv3.Contract.Name(&_Poolv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Poolv3 *Poolv3CallerSession) Name() (string, error) {
	return _Poolv3.Contract.Name(&_Poolv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolv3 *Poolv3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolv3 *Poolv3Session) Paused() (bool, error) {
	return _Poolv3.Contract.Paused(&_Poolv3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolv3 *Poolv3CallerSession) Paused() (bool, error) {
	return _Poolv3.Contract.Paused(&_Poolv3.CallOpts)
}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_Poolv3 *Poolv3Caller) PoolQuotaKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "poolQuotaKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_Poolv3 *Poolv3Session) PoolQuotaKeeper() (common.Address, error) {
	return _Poolv3.Contract.PoolQuotaKeeper(&_Poolv3.CallOpts)
}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_Poolv3 *Poolv3CallerSession) PoolQuotaKeeper() (common.Address, error) {
	return _Poolv3.Contract.PoolQuotaKeeper(&_Poolv3.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3Caller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3Session) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewDeposit(&_Poolv3.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256 shares)
func (_Poolv3 *Poolv3CallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewDeposit(&_Poolv3.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3Caller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3Session) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewMint(&_Poolv3.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewMint(&_Poolv3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3Caller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3Session) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewRedeem(&_Poolv3.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewRedeem(&_Poolv3.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Poolv3 *Poolv3Caller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Poolv3 *Poolv3Session) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewWithdraw(&_Poolv3.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Poolv3.Contract.PreviewWithdraw(&_Poolv3.CallOpts, assets)
}

// QuotaRevenue is a free data retrieval call binding the contract method 0x5a6952e4.
//
// Solidity: function quotaRevenue() view returns(uint256)
func (_Poolv3 *Poolv3Caller) QuotaRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "quotaRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaRevenue is a free data retrieval call binding the contract method 0x5a6952e4.
//
// Solidity: function quotaRevenue() view returns(uint256)
func (_Poolv3 *Poolv3Session) QuotaRevenue() (*big.Int, error) {
	return _Poolv3.Contract.QuotaRevenue(&_Poolv3.CallOpts)
}

// QuotaRevenue is a free data retrieval call binding the contract method 0x5a6952e4.
//
// Solidity: function quotaRevenue() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) QuotaRevenue() (*big.Int, error) {
	return _Poolv3.Contract.QuotaRevenue(&_Poolv3.CallOpts)
}

// SupplyRate is a free data retrieval call binding the contract method 0xad2961a3.
//
// Solidity: function supplyRate() view returns(uint256)
func (_Poolv3 *Poolv3Caller) SupplyRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "supplyRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRate is a free data retrieval call binding the contract method 0xad2961a3.
//
// Solidity: function supplyRate() view returns(uint256)
func (_Poolv3 *Poolv3Session) SupplyRate() (*big.Int, error) {
	return _Poolv3.Contract.SupplyRate(&_Poolv3.CallOpts)
}

// SupplyRate is a free data retrieval call binding the contract method 0xad2961a3.
//
// Solidity: function supplyRate() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) SupplyRate() (*big.Int, error) {
	return _Poolv3.Contract.SupplyRate(&_Poolv3.CallOpts)
}

// SupportsQuotas is a free data retrieval call binding the contract method 0x9fd8a10c.
//
// Solidity: function supportsQuotas() view returns(bool)
func (_Poolv3 *Poolv3Caller) SupportsQuotas(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "supportsQuotas")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsQuotas is a free data retrieval call binding the contract method 0x9fd8a10c.
//
// Solidity: function supportsQuotas() view returns(bool)
func (_Poolv3 *Poolv3Session) SupportsQuotas() (bool, error) {
	return _Poolv3.Contract.SupportsQuotas(&_Poolv3.CallOpts)
}

// SupportsQuotas is a free data retrieval call binding the contract method 0x9fd8a10c.
//
// Solidity: function supportsQuotas() view returns(bool)
func (_Poolv3 *Poolv3CallerSession) SupportsQuotas() (bool, error) {
	return _Poolv3.Contract.SupportsQuotas(&_Poolv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolv3 *Poolv3Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolv3 *Poolv3Session) Symbol() (string, error) {
	return _Poolv3.Contract.Symbol(&_Poolv3.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Poolv3 *Poolv3CallerSession) Symbol() (string, error) {
	return _Poolv3.Contract.Symbol(&_Poolv3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Poolv3 *Poolv3Caller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Poolv3 *Poolv3Session) TotalAssets() (*big.Int, error) {
	return _Poolv3.Contract.TotalAssets(&_Poolv3.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 assets)
func (_Poolv3 *Poolv3CallerSession) TotalAssets() (*big.Int, error) {
	return _Poolv3.Contract.TotalAssets(&_Poolv3.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_Poolv3 *Poolv3Caller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_Poolv3 *Poolv3Session) TotalBorrowed() (*big.Int, error) {
	return _Poolv3.Contract.TotalBorrowed(&_Poolv3.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) TotalBorrowed() (*big.Int, error) {
	return _Poolv3.Contract.TotalBorrowed(&_Poolv3.CallOpts)
}

// TotalDebtLimit is a free data retrieval call binding the contract method 0x183ace90.
//
// Solidity: function totalDebtLimit() view returns(uint256)
func (_Poolv3 *Poolv3Caller) TotalDebtLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "totalDebtLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebtLimit is a free data retrieval call binding the contract method 0x183ace90.
//
// Solidity: function totalDebtLimit() view returns(uint256)
func (_Poolv3 *Poolv3Session) TotalDebtLimit() (*big.Int, error) {
	return _Poolv3.Contract.TotalDebtLimit(&_Poolv3.CallOpts)
}

// TotalDebtLimit is a free data retrieval call binding the contract method 0x183ace90.
//
// Solidity: function totalDebtLimit() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) TotalDebtLimit() (*big.Int, error) {
	return _Poolv3.Contract.TotalDebtLimit(&_Poolv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolv3 *Poolv3Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolv3 *Poolv3Session) TotalSupply() (*big.Int, error) {
	return _Poolv3.Contract.TotalSupply(&_Poolv3.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) TotalSupply() (*big.Int, error) {
	return _Poolv3.Contract.TotalSupply(&_Poolv3.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Poolv3 *Poolv3Caller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Poolv3 *Poolv3Session) Treasury() (common.Address, error) {
	return _Poolv3.Contract.Treasury(&_Poolv3.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Poolv3 *Poolv3CallerSession) Treasury() (common.Address, error) {
	return _Poolv3.Contract.Treasury(&_Poolv3.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Poolv3 *Poolv3Caller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Poolv3 *Poolv3Session) UnderlyingToken() (common.Address, error) {
	return _Poolv3.Contract.UnderlyingToken(&_Poolv3.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Poolv3 *Poolv3CallerSession) UnderlyingToken() (common.Address, error) {
	return _Poolv3.Contract.UnderlyingToken(&_Poolv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Poolv3 *Poolv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Poolv3 *Poolv3Session) Version() (*big.Int, error) {
	return _Poolv3.Contract.Version(&_Poolv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Poolv3 *Poolv3CallerSession) Version() (*big.Int, error) {
	return _Poolv3.Contract.Version(&_Poolv3.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint16)
func (_Poolv3 *Poolv3Caller) WithdrawFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Poolv3.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint16)
func (_Poolv3 *Poolv3Session) WithdrawFee() (uint16, error) {
	return _Poolv3.Contract.WithdrawFee(&_Poolv3.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint16)
func (_Poolv3 *Poolv3CallerSession) WithdrawFee() (uint16, error) {
	return _Poolv3.Contract.WithdrawFee(&_Poolv3.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.Approve(&_Poolv3.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.Approve(&_Poolv3.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Poolv3 *Poolv3Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Poolv3 *Poolv3Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.DecreaseAllowance(&_Poolv3.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Poolv3 *Poolv3TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.DecreaseAllowance(&_Poolv3.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Poolv3 *Poolv3Transactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Poolv3 *Poolv3Session) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Deposit(&_Poolv3.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Poolv3 *Poolv3TransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Deposit(&_Poolv3.TransactOpts, assets, receiver)
}

// DepositWithReferral is a paid mutator transaction binding the contract method 0x6e37c0b3.
//
// Solidity: function depositWithReferral(uint256 assets, address receiver, uint16 referralCode) returns(uint256 shares)
func (_Poolv3 *Poolv3Transactor) DepositWithReferral(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "depositWithReferral", assets, receiver, referralCode)
}

// DepositWithReferral is a paid mutator transaction binding the contract method 0x6e37c0b3.
//
// Solidity: function depositWithReferral(uint256 assets, address receiver, uint16 referralCode) returns(uint256 shares)
func (_Poolv3 *Poolv3Session) DepositWithReferral(assets *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.Contract.DepositWithReferral(&_Poolv3.TransactOpts, assets, receiver, referralCode)
}

// DepositWithReferral is a paid mutator transaction binding the contract method 0x6e37c0b3.
//
// Solidity: function depositWithReferral(uint256 assets, address receiver, uint16 referralCode) returns(uint256 shares)
func (_Poolv3 *Poolv3TransactorSession) DepositWithReferral(assets *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.Contract.DepositWithReferral(&_Poolv3.TransactOpts, assets, receiver, referralCode)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Poolv3 *Poolv3Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Poolv3 *Poolv3Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.IncreaseAllowance(&_Poolv3.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Poolv3 *Poolv3TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.IncreaseAllowance(&_Poolv3.TransactOpts, spender, addedValue)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_Poolv3 *Poolv3Transactor) LendCreditAccount(opts *bind.TransactOpts, borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "lendCreditAccount", borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_Poolv3 *Poolv3Session) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.LendCreditAccount(&_Poolv3.TransactOpts, borrowedAmount, creditAccount)
}

// LendCreditAccount is a paid mutator transaction binding the contract method 0xbf28068b.
//
// Solidity: function lendCreditAccount(uint256 borrowedAmount, address creditAccount) returns()
func (_Poolv3 *Poolv3TransactorSession) LendCreditAccount(borrowedAmount *big.Int, creditAccount common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.LendCreditAccount(&_Poolv3.TransactOpts, borrowedAmount, creditAccount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Poolv3 *Poolv3Transactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Poolv3 *Poolv3Session) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Mint(&_Poolv3.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Poolv3 *Poolv3TransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Mint(&_Poolv3.TransactOpts, shares, receiver)
}

// MintWithReferral is a paid mutator transaction binding the contract method 0x3413048d.
//
// Solidity: function mintWithReferral(uint256 shares, address receiver, uint16 referralCode) returns(uint256 assets)
func (_Poolv3 *Poolv3Transactor) MintWithReferral(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "mintWithReferral", shares, receiver, referralCode)
}

// MintWithReferral is a paid mutator transaction binding the contract method 0x3413048d.
//
// Solidity: function mintWithReferral(uint256 shares, address receiver, uint16 referralCode) returns(uint256 assets)
func (_Poolv3 *Poolv3Session) MintWithReferral(shares *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.Contract.MintWithReferral(&_Poolv3.TransactOpts, shares, receiver, referralCode)
}

// MintWithReferral is a paid mutator transaction binding the contract method 0x3413048d.
//
// Solidity: function mintWithReferral(uint256 shares, address receiver, uint16 referralCode) returns(uint256 assets)
func (_Poolv3 *Poolv3TransactorSession) MintWithReferral(shares *big.Int, receiver common.Address, referralCode uint16) (*types.Transaction, error) {
	return _Poolv3.Contract.MintWithReferral(&_Poolv3.TransactOpts, shares, receiver, referralCode)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolv3 *Poolv3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolv3 *Poolv3Session) Pause() (*types.Transaction, error) {
	return _Poolv3.Contract.Pause(&_Poolv3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolv3 *Poolv3TransactorSession) Pause() (*types.Transaction, error) {
	return _Poolv3.Contract.Pause(&_Poolv3.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Poolv3 *Poolv3Transactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Poolv3 *Poolv3Session) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Redeem(&_Poolv3.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Poolv3 *Poolv3TransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Redeem(&_Poolv3.TransactOpts, shares, receiver, owner)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 repaidAmount, uint256 profit, uint256 loss) returns()
func (_Poolv3 *Poolv3Transactor) RepayCreditAccount(opts *bind.TransactOpts, repaidAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "repayCreditAccount", repaidAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 repaidAmount, uint256 profit, uint256 loss) returns()
func (_Poolv3 *Poolv3Session) RepayCreditAccount(repaidAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.RepayCreditAccount(&_Poolv3.TransactOpts, repaidAmount, profit, loss)
}

// RepayCreditAccount is a paid mutator transaction binding the contract method 0xca9505e4.
//
// Solidity: function repayCreditAccount(uint256 repaidAmount, uint256 profit, uint256 loss) returns()
func (_Poolv3 *Poolv3TransactorSession) RepayCreditAccount(repaidAmount *big.Int, profit *big.Int, loss *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.RepayCreditAccount(&_Poolv3.TransactOpts, repaidAmount, profit, loss)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Poolv3 *Poolv3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Poolv3 *Poolv3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetController(&_Poolv3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_Poolv3 *Poolv3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetController(&_Poolv3.TransactOpts, newController)
}

// SetCreditManagerDebtLimit is a paid mutator transaction binding the contract method 0x79e4e3a9.
//
// Solidity: function setCreditManagerDebtLimit(address creditManager, uint256 newLimit) returns()
func (_Poolv3 *Poolv3Transactor) SetCreditManagerDebtLimit(opts *bind.TransactOpts, creditManager common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setCreditManagerDebtLimit", creditManager, newLimit)
}

// SetCreditManagerDebtLimit is a paid mutator transaction binding the contract method 0x79e4e3a9.
//
// Solidity: function setCreditManagerDebtLimit(address creditManager, uint256 newLimit) returns()
func (_Poolv3 *Poolv3Session) SetCreditManagerDebtLimit(creditManager common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetCreditManagerDebtLimit(&_Poolv3.TransactOpts, creditManager, newLimit)
}

// SetCreditManagerDebtLimit is a paid mutator transaction binding the contract method 0x79e4e3a9.
//
// Solidity: function setCreditManagerDebtLimit(address creditManager, uint256 newLimit) returns()
func (_Poolv3 *Poolv3TransactorSession) SetCreditManagerDebtLimit(creditManager common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetCreditManagerDebtLimit(&_Poolv3.TransactOpts, creditManager, newLimit)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newInterestRateModel) returns()
func (_Poolv3 *Poolv3Transactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newInterestRateModel) returns()
func (_Poolv3 *Poolv3Session) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetInterestRateModel(&_Poolv3.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newInterestRateModel) returns()
func (_Poolv3 *Poolv3TransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetInterestRateModel(&_Poolv3.TransactOpts, newInterestRateModel)
}

// SetPoolQuotaKeeper is a paid mutator transaction binding the contract method 0x1ab7c7d7.
//
// Solidity: function setPoolQuotaKeeper(address newPoolQuotaKeeper) returns()
func (_Poolv3 *Poolv3Transactor) SetPoolQuotaKeeper(opts *bind.TransactOpts, newPoolQuotaKeeper common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setPoolQuotaKeeper", newPoolQuotaKeeper)
}

// SetPoolQuotaKeeper is a paid mutator transaction binding the contract method 0x1ab7c7d7.
//
// Solidity: function setPoolQuotaKeeper(address newPoolQuotaKeeper) returns()
func (_Poolv3 *Poolv3Session) SetPoolQuotaKeeper(newPoolQuotaKeeper common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetPoolQuotaKeeper(&_Poolv3.TransactOpts, newPoolQuotaKeeper)
}

// SetPoolQuotaKeeper is a paid mutator transaction binding the contract method 0x1ab7c7d7.
//
// Solidity: function setPoolQuotaKeeper(address newPoolQuotaKeeper) returns()
func (_Poolv3 *Poolv3TransactorSession) SetPoolQuotaKeeper(newPoolQuotaKeeper common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.SetPoolQuotaKeeper(&_Poolv3.TransactOpts, newPoolQuotaKeeper)
}

// SetQuotaRevenue is a paid mutator transaction binding the contract method 0x275df3ad.
//
// Solidity: function setQuotaRevenue(uint256 newQuotaRevenue) returns()
func (_Poolv3 *Poolv3Transactor) SetQuotaRevenue(opts *bind.TransactOpts, newQuotaRevenue *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setQuotaRevenue", newQuotaRevenue)
}

// SetQuotaRevenue is a paid mutator transaction binding the contract method 0x275df3ad.
//
// Solidity: function setQuotaRevenue(uint256 newQuotaRevenue) returns()
func (_Poolv3 *Poolv3Session) SetQuotaRevenue(newQuotaRevenue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetQuotaRevenue(&_Poolv3.TransactOpts, newQuotaRevenue)
}

// SetQuotaRevenue is a paid mutator transaction binding the contract method 0x275df3ad.
//
// Solidity: function setQuotaRevenue(uint256 newQuotaRevenue) returns()
func (_Poolv3 *Poolv3TransactorSession) SetQuotaRevenue(newQuotaRevenue *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetQuotaRevenue(&_Poolv3.TransactOpts, newQuotaRevenue)
}

// SetTotalDebtLimit is a paid mutator transaction binding the contract method 0x871d7268.
//
// Solidity: function setTotalDebtLimit(uint256 newLimit) returns()
func (_Poolv3 *Poolv3Transactor) SetTotalDebtLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setTotalDebtLimit", newLimit)
}

// SetTotalDebtLimit is a paid mutator transaction binding the contract method 0x871d7268.
//
// Solidity: function setTotalDebtLimit(uint256 newLimit) returns()
func (_Poolv3 *Poolv3Session) SetTotalDebtLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetTotalDebtLimit(&_Poolv3.TransactOpts, newLimit)
}

// SetTotalDebtLimit is a paid mutator transaction binding the contract method 0x871d7268.
//
// Solidity: function setTotalDebtLimit(uint256 newLimit) returns()
func (_Poolv3 *Poolv3TransactorSession) SetTotalDebtLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetTotalDebtLimit(&_Poolv3.TransactOpts, newLimit)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 newWithdrawFee) returns()
func (_Poolv3 *Poolv3Transactor) SetWithdrawFee(opts *bind.TransactOpts, newWithdrawFee *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "setWithdrawFee", newWithdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 newWithdrawFee) returns()
func (_Poolv3 *Poolv3Session) SetWithdrawFee(newWithdrawFee *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetWithdrawFee(&_Poolv3.TransactOpts, newWithdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 newWithdrawFee) returns()
func (_Poolv3 *Poolv3TransactorSession) SetWithdrawFee(newWithdrawFee *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.SetWithdrawFee(&_Poolv3.TransactOpts, newWithdrawFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.Transfer(&_Poolv3.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.Transfer(&_Poolv3.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.TransferFrom(&_Poolv3.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Poolv3 *Poolv3TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.TransferFrom(&_Poolv3.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolv3 *Poolv3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolv3 *Poolv3Session) Unpause() (*types.Transaction, error) {
	return _Poolv3.Contract.Unpause(&_Poolv3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolv3 *Poolv3TransactorSession) Unpause() (*types.Transaction, error) {
	return _Poolv3.Contract.Unpause(&_Poolv3.TransactOpts)
}

// UpdateQuotaRevenue is a paid mutator transaction binding the contract method 0xd6458eea.
//
// Solidity: function updateQuotaRevenue(int256 quotaRevenueDelta) returns()
func (_Poolv3 *Poolv3Transactor) UpdateQuotaRevenue(opts *bind.TransactOpts, quotaRevenueDelta *big.Int) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "updateQuotaRevenue", quotaRevenueDelta)
}

// UpdateQuotaRevenue is a paid mutator transaction binding the contract method 0xd6458eea.
//
// Solidity: function updateQuotaRevenue(int256 quotaRevenueDelta) returns()
func (_Poolv3 *Poolv3Session) UpdateQuotaRevenue(quotaRevenueDelta *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.UpdateQuotaRevenue(&_Poolv3.TransactOpts, quotaRevenueDelta)
}

// UpdateQuotaRevenue is a paid mutator transaction binding the contract method 0xd6458eea.
//
// Solidity: function updateQuotaRevenue(int256 quotaRevenueDelta) returns()
func (_Poolv3 *Poolv3TransactorSession) UpdateQuotaRevenue(quotaRevenueDelta *big.Int) (*types.Transaction, error) {
	return _Poolv3.Contract.UpdateQuotaRevenue(&_Poolv3.TransactOpts, quotaRevenueDelta)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Poolv3 *Poolv3Transactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Poolv3 *Poolv3Session) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Withdraw(&_Poolv3.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Poolv3 *Poolv3TransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Poolv3.Contract.Withdraw(&_Poolv3.TransactOpts, assets, receiver, owner)
}

// Poolv3AddCreditManagerIterator is returned from FilterAddCreditManager and is used to iterate over the raw logs and unpacked data for AddCreditManager events raised by the Poolv3 contract.
type Poolv3AddCreditManagerIterator struct {
	Event *Poolv3AddCreditManager // Event containing the contract specifics and raw log

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
func (it *Poolv3AddCreditManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3AddCreditManager)
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
		it.Event = new(Poolv3AddCreditManager)
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
func (it *Poolv3AddCreditManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3AddCreditManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3AddCreditManager represents a AddCreditManager event raised by the Poolv3 contract.
type Poolv3AddCreditManager struct {
	CreditManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddCreditManager is a free log retrieval operation binding the contract event 0xbca7ba46bb626fab79d5a673d0d8293df21968a25350c4d71433f98600618f5f.
//
// Solidity: event AddCreditManager(address indexed creditManager)
func (_Poolv3 *Poolv3Filterer) FilterAddCreditManager(opts *bind.FilterOpts, creditManager []common.Address) (*Poolv3AddCreditManagerIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "AddCreditManager", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3AddCreditManagerIterator{contract: _Poolv3.contract, event: "AddCreditManager", logs: logs, sub: sub}, nil
}

// WatchAddCreditManager is a free log subscription operation binding the contract event 0xbca7ba46bb626fab79d5a673d0d8293df21968a25350c4d71433f98600618f5f.
//
// Solidity: event AddCreditManager(address indexed creditManager)
func (_Poolv3 *Poolv3Filterer) WatchAddCreditManager(opts *bind.WatchOpts, sink chan<- *Poolv3AddCreditManager, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "AddCreditManager", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3AddCreditManager)
				if err := _Poolv3.contract.UnpackLog(event, "AddCreditManager", log); err != nil {
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
func (_Poolv3 *Poolv3Filterer) ParseAddCreditManager(log types.Log) (*Poolv3AddCreditManager, error) {
	event := new(Poolv3AddCreditManager)
	if err := _Poolv3.contract.UnpackLog(event, "AddCreditManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Poolv3 contract.
type Poolv3ApprovalIterator struct {
	Event *Poolv3Approval // Event containing the contract specifics and raw log

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
func (it *Poolv3ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Approval)
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
		it.Event = new(Poolv3Approval)
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
func (it *Poolv3ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Approval represents a Approval event raised by the Poolv3 contract.
type Poolv3Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Poolv3 *Poolv3Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Poolv3ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3ApprovalIterator{contract: _Poolv3.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Poolv3 *Poolv3Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Poolv3Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Approval)
				if err := _Poolv3.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Poolv3 *Poolv3Filterer) ParseApproval(log types.Log) (*Poolv3Approval, error) {
	event := new(Poolv3Approval)
	if err := _Poolv3.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Poolv3 contract.
type Poolv3BorrowIterator struct {
	Event *Poolv3Borrow // Event containing the contract specifics and raw log

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
func (it *Poolv3BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Borrow)
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
		it.Event = new(Poolv3Borrow)
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
func (it *Poolv3BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Borrow represents a Borrow event raised by the Poolv3 contract.
type Poolv3Borrow struct {
	CreditManager common.Address
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_Poolv3 *Poolv3Filterer) FilterBorrow(opts *bind.FilterOpts, creditManager []common.Address, creditAccount []common.Address) (*Poolv3BorrowIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3BorrowIterator{contract: _Poolv3.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_Poolv3 *Poolv3Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Poolv3Borrow, creditManager []common.Address, creditAccount []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}
	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Borrow", creditManagerRule, creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Borrow)
				if err := _Poolv3.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x312a5e5e1079f5dda4e95dbbd0b908b291fd5b992ef22073643ab691572c5b52.
//
// Solidity: event Borrow(address indexed creditManager, address indexed creditAccount, uint256 amount)
func (_Poolv3 *Poolv3Filterer) ParseBorrow(log types.Log) (*Poolv3Borrow, error) {
	event := new(Poolv3Borrow)
	if err := _Poolv3.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Poolv3 contract.
type Poolv3DepositIterator struct {
	Event *Poolv3Deposit // Event containing the contract specifics and raw log

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
func (it *Poolv3DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Deposit)
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
		it.Event = new(Poolv3Deposit)
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
func (it *Poolv3DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Deposit represents a Deposit event raised by the Poolv3 contract.
type Poolv3Deposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*Poolv3DepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3DepositIterator{contract: _Poolv3.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Poolv3Deposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Deposit)
				if err := _Poolv3.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) ParseDeposit(log types.Log) (*Poolv3Deposit, error) {
	event := new(Poolv3Deposit)
	if err := _Poolv3.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3IncurUncoveredLossIterator is returned from FilterIncurUncoveredLoss and is used to iterate over the raw logs and unpacked data for IncurUncoveredLoss events raised by the Poolv3 contract.
type Poolv3IncurUncoveredLossIterator struct {
	Event *Poolv3IncurUncoveredLoss // Event containing the contract specifics and raw log

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
func (it *Poolv3IncurUncoveredLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3IncurUncoveredLoss)
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
		it.Event = new(Poolv3IncurUncoveredLoss)
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
func (it *Poolv3IncurUncoveredLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3IncurUncoveredLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3IncurUncoveredLoss represents a IncurUncoveredLoss event raised by the Poolv3 contract.
type Poolv3IncurUncoveredLoss struct {
	CreditManager common.Address
	Loss          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterIncurUncoveredLoss is a free log retrieval operation binding the contract event 0x33fc1787be707f18e553b02263e12d2fa6d2d40733535382066fd1d77e32c595.
//
// Solidity: event IncurUncoveredLoss(address indexed creditManager, uint256 loss)
func (_Poolv3 *Poolv3Filterer) FilterIncurUncoveredLoss(opts *bind.FilterOpts, creditManager []common.Address) (*Poolv3IncurUncoveredLossIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "IncurUncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3IncurUncoveredLossIterator{contract: _Poolv3.contract, event: "IncurUncoveredLoss", logs: logs, sub: sub}, nil
}

// WatchIncurUncoveredLoss is a free log subscription operation binding the contract event 0x33fc1787be707f18e553b02263e12d2fa6d2d40733535382066fd1d77e32c595.
//
// Solidity: event IncurUncoveredLoss(address indexed creditManager, uint256 loss)
func (_Poolv3 *Poolv3Filterer) WatchIncurUncoveredLoss(opts *bind.WatchOpts, sink chan<- *Poolv3IncurUncoveredLoss, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "IncurUncoveredLoss", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3IncurUncoveredLoss)
				if err := _Poolv3.contract.UnpackLog(event, "IncurUncoveredLoss", log); err != nil {
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

// ParseIncurUncoveredLoss is a log parse operation binding the contract event 0x33fc1787be707f18e553b02263e12d2fa6d2d40733535382066fd1d77e32c595.
//
// Solidity: event IncurUncoveredLoss(address indexed creditManager, uint256 loss)
func (_Poolv3 *Poolv3Filterer) ParseIncurUncoveredLoss(log types.Log) (*Poolv3IncurUncoveredLoss, error) {
	event := new(Poolv3IncurUncoveredLoss)
	if err := _Poolv3.contract.UnpackLog(event, "IncurUncoveredLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the Poolv3 contract.
type Poolv3NewControllerIterator struct {
	Event *Poolv3NewController // Event containing the contract specifics and raw log

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
func (it *Poolv3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3NewController)
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
		it.Event = new(Poolv3NewController)
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
func (it *Poolv3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3NewController represents a NewController event raised by the Poolv3 contract.
type Poolv3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_Poolv3 *Poolv3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*Poolv3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3NewControllerIterator{contract: _Poolv3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_Poolv3 *Poolv3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *Poolv3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3NewController)
				if err := _Poolv3.contract.UnpackLog(event, "NewController", log); err != nil {
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
func (_Poolv3 *Poolv3Filterer) ParseNewController(log types.Log) (*Poolv3NewController, error) {
	event := new(Poolv3NewController)
	if err := _Poolv3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Poolv3 contract.
type Poolv3PausedIterator struct {
	Event *Poolv3Paused // Event containing the contract specifics and raw log

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
func (it *Poolv3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Paused)
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
		it.Event = new(Poolv3Paused)
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
func (it *Poolv3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Paused represents a Paused event raised by the Poolv3 contract.
type Poolv3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Poolv3 *Poolv3Filterer) FilterPaused(opts *bind.FilterOpts) (*Poolv3PausedIterator, error) {

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Poolv3PausedIterator{contract: _Poolv3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Poolv3 *Poolv3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Poolv3Paused) (event.Subscription, error) {

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Paused)
				if err := _Poolv3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Poolv3 *Poolv3Filterer) ParsePaused(log types.Log) (*Poolv3Paused, error) {
	event := new(Poolv3Paused)
	if err := _Poolv3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3ReferIterator is returned from FilterRefer and is used to iterate over the raw logs and unpacked data for Refer events raised by the Poolv3 contract.
type Poolv3ReferIterator struct {
	Event *Poolv3Refer // Event containing the contract specifics and raw log

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
func (it *Poolv3ReferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Refer)
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
		it.Event = new(Poolv3Refer)
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
func (it *Poolv3ReferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3ReferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Refer represents a Refer event raised by the Poolv3 contract.
type Poolv3Refer struct {
	OnBehalfOf   common.Address
	ReferralCode *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefer is a free log retrieval operation binding the contract event 0xd01c12ea61a25b0a57aa9b86b06dacf8f140567dd44ec9db66ef7955f6a956d2.
//
// Solidity: event Refer(address indexed onBehalfOf, uint256 indexed referralCode, uint256 amount)
func (_Poolv3 *Poolv3Filterer) FilterRefer(opts *bind.FilterOpts, onBehalfOf []common.Address, referralCode []*big.Int) (*Poolv3ReferIterator, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Refer", onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3ReferIterator{contract: _Poolv3.contract, event: "Refer", logs: logs, sub: sub}, nil
}

// WatchRefer is a free log subscription operation binding the contract event 0xd01c12ea61a25b0a57aa9b86b06dacf8f140567dd44ec9db66ef7955f6a956d2.
//
// Solidity: event Refer(address indexed onBehalfOf, uint256 indexed referralCode, uint256 amount)
func (_Poolv3 *Poolv3Filterer) WatchRefer(opts *bind.WatchOpts, sink chan<- *Poolv3Refer, onBehalfOf []common.Address, referralCode []*big.Int) (event.Subscription, error) {

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Refer", onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Refer)
				if err := _Poolv3.contract.UnpackLog(event, "Refer", log); err != nil {
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

// ParseRefer is a log parse operation binding the contract event 0xd01c12ea61a25b0a57aa9b86b06dacf8f140567dd44ec9db66ef7955f6a956d2.
//
// Solidity: event Refer(address indexed onBehalfOf, uint256 indexed referralCode, uint256 amount)
func (_Poolv3 *Poolv3Filterer) ParseRefer(log types.Log) (*Poolv3Refer, error) {
	event := new(Poolv3Refer)
	if err := _Poolv3.contract.UnpackLog(event, "Refer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3RepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Poolv3 contract.
type Poolv3RepayIterator struct {
	Event *Poolv3Repay // Event containing the contract specifics and raw log

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
func (it *Poolv3RepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Repay)
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
		it.Event = new(Poolv3Repay)
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
func (it *Poolv3RepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3RepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Repay represents a Repay event raised by the Poolv3 contract.
type Poolv3Repay struct {
	CreditManager  common.Address
	BorrowedAmount *big.Int
	Profit         *big.Int
	Loss           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_Poolv3 *Poolv3Filterer) FilterRepay(opts *bind.FilterOpts, creditManager []common.Address) (*Poolv3RepayIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3RepayIterator{contract: _Poolv3.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_Poolv3 *Poolv3Filterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *Poolv3Repay, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Repay", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Repay)
				if err := _Poolv3.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x2fe77b1c99aca6b022b8efc6e3e8dd1b48b30748709339b65c50ef3263443e09.
//
// Solidity: event Repay(address indexed creditManager, uint256 borrowedAmount, uint256 profit, uint256 loss)
func (_Poolv3 *Poolv3Filterer) ParseRepay(log types.Log) (*Poolv3Repay, error) {
	event := new(Poolv3Repay)
	if err := _Poolv3.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3SetCreditManagerDebtLimitIterator is returned from FilterSetCreditManagerDebtLimit and is used to iterate over the raw logs and unpacked data for SetCreditManagerDebtLimit events raised by the Poolv3 contract.
type Poolv3SetCreditManagerDebtLimitIterator struct {
	Event *Poolv3SetCreditManagerDebtLimit // Event containing the contract specifics and raw log

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
func (it *Poolv3SetCreditManagerDebtLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3SetCreditManagerDebtLimit)
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
		it.Event = new(Poolv3SetCreditManagerDebtLimit)
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
func (it *Poolv3SetCreditManagerDebtLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3SetCreditManagerDebtLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3SetCreditManagerDebtLimit represents a SetCreditManagerDebtLimit event raised by the Poolv3 contract.
type Poolv3SetCreditManagerDebtLimit struct {
	CreditManager common.Address
	NewLimit      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetCreditManagerDebtLimit is a free log retrieval operation binding the contract event 0xce20e043afe93acdab0352023688eb8da23cdfd33d80471cce1e6c9239662bcd.
//
// Solidity: event SetCreditManagerDebtLimit(address indexed creditManager, uint256 newLimit)
func (_Poolv3 *Poolv3Filterer) FilterSetCreditManagerDebtLimit(opts *bind.FilterOpts, creditManager []common.Address) (*Poolv3SetCreditManagerDebtLimitIterator, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "SetCreditManagerDebtLimit", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3SetCreditManagerDebtLimitIterator{contract: _Poolv3.contract, event: "SetCreditManagerDebtLimit", logs: logs, sub: sub}, nil
}

// WatchSetCreditManagerDebtLimit is a free log subscription operation binding the contract event 0xce20e043afe93acdab0352023688eb8da23cdfd33d80471cce1e6c9239662bcd.
//
// Solidity: event SetCreditManagerDebtLimit(address indexed creditManager, uint256 newLimit)
func (_Poolv3 *Poolv3Filterer) WatchSetCreditManagerDebtLimit(opts *bind.WatchOpts, sink chan<- *Poolv3SetCreditManagerDebtLimit, creditManager []common.Address) (event.Subscription, error) {

	var creditManagerRule []interface{}
	for _, creditManagerItem := range creditManager {
		creditManagerRule = append(creditManagerRule, creditManagerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "SetCreditManagerDebtLimit", creditManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3SetCreditManagerDebtLimit)
				if err := _Poolv3.contract.UnpackLog(event, "SetCreditManagerDebtLimit", log); err != nil {
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

// ParseSetCreditManagerDebtLimit is a log parse operation binding the contract event 0xce20e043afe93acdab0352023688eb8da23cdfd33d80471cce1e6c9239662bcd.
//
// Solidity: event SetCreditManagerDebtLimit(address indexed creditManager, uint256 newLimit)
func (_Poolv3 *Poolv3Filterer) ParseSetCreditManagerDebtLimit(log types.Log) (*Poolv3SetCreditManagerDebtLimit, error) {
	event := new(Poolv3SetCreditManagerDebtLimit)
	if err := _Poolv3.contract.UnpackLog(event, "SetCreditManagerDebtLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3SetInterestRateModelIterator is returned from FilterSetInterestRateModel and is used to iterate over the raw logs and unpacked data for SetInterestRateModel events raised by the Poolv3 contract.
type Poolv3SetInterestRateModelIterator struct {
	Event *Poolv3SetInterestRateModel // Event containing the contract specifics and raw log

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
func (it *Poolv3SetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3SetInterestRateModel)
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
		it.Event = new(Poolv3SetInterestRateModel)
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
func (it *Poolv3SetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3SetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3SetInterestRateModel represents a SetInterestRateModel event raised by the Poolv3 contract.
type Poolv3SetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetInterestRateModel is a free log retrieval operation binding the contract event 0x60d671e95013fc5fd0cf35d947791aa49209ad86fccf748e0b126f3f9f0a83ba.
//
// Solidity: event SetInterestRateModel(address indexed newInterestRateModel)
func (_Poolv3 *Poolv3Filterer) FilterSetInterestRateModel(opts *bind.FilterOpts, newInterestRateModel []common.Address) (*Poolv3SetInterestRateModelIterator, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "SetInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3SetInterestRateModelIterator{contract: _Poolv3.contract, event: "SetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchSetInterestRateModel is a free log subscription operation binding the contract event 0x60d671e95013fc5fd0cf35d947791aa49209ad86fccf748e0b126f3f9f0a83ba.
//
// Solidity: event SetInterestRateModel(address indexed newInterestRateModel)
func (_Poolv3 *Poolv3Filterer) WatchSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *Poolv3SetInterestRateModel, newInterestRateModel []common.Address) (event.Subscription, error) {

	var newInterestRateModelRule []interface{}
	for _, newInterestRateModelItem := range newInterestRateModel {
		newInterestRateModelRule = append(newInterestRateModelRule, newInterestRateModelItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "SetInterestRateModel", newInterestRateModelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3SetInterestRateModel)
				if err := _Poolv3.contract.UnpackLog(event, "SetInterestRateModel", log); err != nil {
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

// ParseSetInterestRateModel is a log parse operation binding the contract event 0x60d671e95013fc5fd0cf35d947791aa49209ad86fccf748e0b126f3f9f0a83ba.
//
// Solidity: event SetInterestRateModel(address indexed newInterestRateModel)
func (_Poolv3 *Poolv3Filterer) ParseSetInterestRateModel(log types.Log) (*Poolv3SetInterestRateModel, error) {
	event := new(Poolv3SetInterestRateModel)
	if err := _Poolv3.contract.UnpackLog(event, "SetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3SetPoolQuotaKeeperIterator is returned from FilterSetPoolQuotaKeeper and is used to iterate over the raw logs and unpacked data for SetPoolQuotaKeeper events raised by the Poolv3 contract.
type Poolv3SetPoolQuotaKeeperIterator struct {
	Event *Poolv3SetPoolQuotaKeeper // Event containing the contract specifics and raw log

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
func (it *Poolv3SetPoolQuotaKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3SetPoolQuotaKeeper)
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
		it.Event = new(Poolv3SetPoolQuotaKeeper)
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
func (it *Poolv3SetPoolQuotaKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3SetPoolQuotaKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3SetPoolQuotaKeeper represents a SetPoolQuotaKeeper event raised by the Poolv3 contract.
type Poolv3SetPoolQuotaKeeper struct {
	NewPoolQuotaKeeper common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetPoolQuotaKeeper is a free log retrieval operation binding the contract event 0x553438de7e02bc6929ef4f6c3653130beca086dd506f1aa2785b58e6a13c3264.
//
// Solidity: event SetPoolQuotaKeeper(address indexed newPoolQuotaKeeper)
func (_Poolv3 *Poolv3Filterer) FilterSetPoolQuotaKeeper(opts *bind.FilterOpts, newPoolQuotaKeeper []common.Address) (*Poolv3SetPoolQuotaKeeperIterator, error) {

	var newPoolQuotaKeeperRule []interface{}
	for _, newPoolQuotaKeeperItem := range newPoolQuotaKeeper {
		newPoolQuotaKeeperRule = append(newPoolQuotaKeeperRule, newPoolQuotaKeeperItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "SetPoolQuotaKeeper", newPoolQuotaKeeperRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3SetPoolQuotaKeeperIterator{contract: _Poolv3.contract, event: "SetPoolQuotaKeeper", logs: logs, sub: sub}, nil
}

// WatchSetPoolQuotaKeeper is a free log subscription operation binding the contract event 0x553438de7e02bc6929ef4f6c3653130beca086dd506f1aa2785b58e6a13c3264.
//
// Solidity: event SetPoolQuotaKeeper(address indexed newPoolQuotaKeeper)
func (_Poolv3 *Poolv3Filterer) WatchSetPoolQuotaKeeper(opts *bind.WatchOpts, sink chan<- *Poolv3SetPoolQuotaKeeper, newPoolQuotaKeeper []common.Address) (event.Subscription, error) {

	var newPoolQuotaKeeperRule []interface{}
	for _, newPoolQuotaKeeperItem := range newPoolQuotaKeeper {
		newPoolQuotaKeeperRule = append(newPoolQuotaKeeperRule, newPoolQuotaKeeperItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "SetPoolQuotaKeeper", newPoolQuotaKeeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3SetPoolQuotaKeeper)
				if err := _Poolv3.contract.UnpackLog(event, "SetPoolQuotaKeeper", log); err != nil {
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

// ParseSetPoolQuotaKeeper is a log parse operation binding the contract event 0x553438de7e02bc6929ef4f6c3653130beca086dd506f1aa2785b58e6a13c3264.
//
// Solidity: event SetPoolQuotaKeeper(address indexed newPoolQuotaKeeper)
func (_Poolv3 *Poolv3Filterer) ParseSetPoolQuotaKeeper(log types.Log) (*Poolv3SetPoolQuotaKeeper, error) {
	event := new(Poolv3SetPoolQuotaKeeper)
	if err := _Poolv3.contract.UnpackLog(event, "SetPoolQuotaKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3SetTotalDebtLimitIterator is returned from FilterSetTotalDebtLimit and is used to iterate over the raw logs and unpacked data for SetTotalDebtLimit events raised by the Poolv3 contract.
type Poolv3SetTotalDebtLimitIterator struct {
	Event *Poolv3SetTotalDebtLimit // Event containing the contract specifics and raw log

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
func (it *Poolv3SetTotalDebtLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3SetTotalDebtLimit)
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
		it.Event = new(Poolv3SetTotalDebtLimit)
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
func (it *Poolv3SetTotalDebtLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3SetTotalDebtLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3SetTotalDebtLimit represents a SetTotalDebtLimit event raised by the Poolv3 contract.
type Poolv3SetTotalDebtLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetTotalDebtLimit is a free log retrieval operation binding the contract event 0x9154a5b15c38625466fe66233214f14f17fd994f819818caf08017b94d0787ba.
//
// Solidity: event SetTotalDebtLimit(uint256 limit)
func (_Poolv3 *Poolv3Filterer) FilterSetTotalDebtLimit(opts *bind.FilterOpts) (*Poolv3SetTotalDebtLimitIterator, error) {

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "SetTotalDebtLimit")
	if err != nil {
		return nil, err
	}
	return &Poolv3SetTotalDebtLimitIterator{contract: _Poolv3.contract, event: "SetTotalDebtLimit", logs: logs, sub: sub}, nil
}

// WatchSetTotalDebtLimit is a free log subscription operation binding the contract event 0x9154a5b15c38625466fe66233214f14f17fd994f819818caf08017b94d0787ba.
//
// Solidity: event SetTotalDebtLimit(uint256 limit)
func (_Poolv3 *Poolv3Filterer) WatchSetTotalDebtLimit(opts *bind.WatchOpts, sink chan<- *Poolv3SetTotalDebtLimit) (event.Subscription, error) {

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "SetTotalDebtLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3SetTotalDebtLimit)
				if err := _Poolv3.contract.UnpackLog(event, "SetTotalDebtLimit", log); err != nil {
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

// ParseSetTotalDebtLimit is a log parse operation binding the contract event 0x9154a5b15c38625466fe66233214f14f17fd994f819818caf08017b94d0787ba.
//
// Solidity: event SetTotalDebtLimit(uint256 limit)
func (_Poolv3 *Poolv3Filterer) ParseSetTotalDebtLimit(log types.Log) (*Poolv3SetTotalDebtLimit, error) {
	event := new(Poolv3SetTotalDebtLimit)
	if err := _Poolv3.contract.UnpackLog(event, "SetTotalDebtLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3SetWithdrawFeeIterator is returned from FilterSetWithdrawFee and is used to iterate over the raw logs and unpacked data for SetWithdrawFee events raised by the Poolv3 contract.
type Poolv3SetWithdrawFeeIterator struct {
	Event *Poolv3SetWithdrawFee // Event containing the contract specifics and raw log

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
func (it *Poolv3SetWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3SetWithdrawFee)
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
		it.Event = new(Poolv3SetWithdrawFee)
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
func (it *Poolv3SetWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3SetWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3SetWithdrawFee represents a SetWithdrawFee event raised by the Poolv3 contract.
type Poolv3SetWithdrawFee struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawFee is a free log retrieval operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 fee)
func (_Poolv3 *Poolv3Filterer) FilterSetWithdrawFee(opts *bind.FilterOpts) (*Poolv3SetWithdrawFeeIterator, error) {

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "SetWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &Poolv3SetWithdrawFeeIterator{contract: _Poolv3.contract, event: "SetWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawFee is a free log subscription operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 fee)
func (_Poolv3 *Poolv3Filterer) WatchSetWithdrawFee(opts *bind.WatchOpts, sink chan<- *Poolv3SetWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "SetWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3SetWithdrawFee)
				if err := _Poolv3.contract.UnpackLog(event, "SetWithdrawFee", log); err != nil {
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

// ParseSetWithdrawFee is a log parse operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 fee)
func (_Poolv3 *Poolv3Filterer) ParseSetWithdrawFee(log types.Log) (*Poolv3SetWithdrawFee, error) {
	event := new(Poolv3SetWithdrawFee)
	if err := _Poolv3.contract.UnpackLog(event, "SetWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Poolv3 contract.
type Poolv3TransferIterator struct {
	Event *Poolv3Transfer // Event containing the contract specifics and raw log

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
func (it *Poolv3TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Transfer)
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
		it.Event = new(Poolv3Transfer)
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
func (it *Poolv3TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Transfer represents a Transfer event raised by the Poolv3 contract.
type Poolv3Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Poolv3 *Poolv3Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Poolv3TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3TransferIterator{contract: _Poolv3.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Poolv3 *Poolv3Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Poolv3Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Transfer)
				if err := _Poolv3.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Poolv3 *Poolv3Filterer) ParseTransfer(log types.Log) (*Poolv3Transfer, error) {
	event := new(Poolv3Transfer)
	if err := _Poolv3.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Poolv3 contract.
type Poolv3UnpausedIterator struct {
	Event *Poolv3Unpaused // Event containing the contract specifics and raw log

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
func (it *Poolv3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Unpaused)
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
		it.Event = new(Poolv3Unpaused)
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
func (it *Poolv3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Unpaused represents a Unpaused event raised by the Poolv3 contract.
type Poolv3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Poolv3 *Poolv3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Poolv3UnpausedIterator, error) {

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Poolv3UnpausedIterator{contract: _Poolv3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Poolv3 *Poolv3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Poolv3Unpaused) (event.Subscription, error) {

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Unpaused)
				if err := _Poolv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Poolv3 *Poolv3Filterer) ParseUnpaused(log types.Log) (*Poolv3Unpaused, error) {
	event := new(Poolv3Unpaused)
	if err := _Poolv3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Poolv3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Poolv3 contract.
type Poolv3WithdrawIterator struct {
	Event *Poolv3Withdraw // Event containing the contract specifics and raw log

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
func (it *Poolv3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Poolv3Withdraw)
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
		it.Event = new(Poolv3Withdraw)
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
func (it *Poolv3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Poolv3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Poolv3Withdraw represents a Withdraw event raised by the Poolv3 contract.
type Poolv3Withdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*Poolv3WithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Poolv3.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &Poolv3WithdrawIterator{contract: _Poolv3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Poolv3Withdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Poolv3.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Poolv3Withdraw)
				if err := _Poolv3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Poolv3 *Poolv3Filterer) ParseWithdraw(log types.Log) (*Poolv3Withdraw, error) {
	event := new(Poolv3Withdraw)
	if err := _Poolv3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
