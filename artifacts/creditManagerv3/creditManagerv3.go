// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditManagerv3

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

// CollateralDebtData is an auto generated low-level Go binding around an user-defined struct.
type CollateralDebtData struct {
	Debt                      *big.Int
	CumulativeIndexNow        *big.Int
	CumulativeIndexLastUpdate *big.Int
	CumulativeQuotaInterest   *big.Int
	AccruedInterest           *big.Int
	AccruedFees               *big.Int
	TotalDebtUSD              *big.Int
	TotalValue                *big.Int
	TotalValueUSD             *big.Int
	TwvUSD                    *big.Int
	EnabledTokensMask         *big.Int
	QuotedTokensMask          *big.Int
	QuotedTokens              []common.Address
	PoolQuotaKeeper           common.Address
}

// RevocationPair is an auto generated low-level Go binding around an user-defined struct.
type RevocationPair struct {
	Spender common.Address
	Token   common.Address
}

// CreditManagerv3MetaData contains all meta data concerning the CreditManagerv3 contract.
var CreditManagerv3MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addressProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accountFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"adapterToContract\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addCollateral\",\"inputs\":[{\"name\":\"payer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveCreditAccount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveToken\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calcDebtAndCollateral\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"task\",\"type\":\"uint8\",\"internalType\":\"enumCollateralCalcTask\"}],\"outputs\":[{\"name\":\"cdd\",\"type\":\"tuple\",\"internalType\":\"structCollateralDebtData\",\"components\":[{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotedTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotedTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_poolQuotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closeCreditAccount\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralTokenByMask\",\"inputs\":[{\"name\":\"tokenMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralTokensCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractToAdapter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditAccountInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"quotaFees\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flags\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"lastDebtUpdate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditAccounts\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditAccounts\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditAccountsLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditConfigurator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditFacade\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enabledTokensMaskOf\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"externalCall\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fees\",\"inputs\":[],\"outputs\":[{\"name\":\"_feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flagsOf\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fullCollateralCheck\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralHints\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"useSafePrices\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"enabledTokensMaskAfter\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getActiveCreditAccountOrRevert\",\"inputs\":[],\"outputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBorrowerOrRevert\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenByMask\",\"inputs\":[{\"name\":\"tokenMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenMaskOrRevert\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidatable\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidateCreditAccount\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralDebtData\",\"type\":\"tuple\",\"internalType\":\"structCollateralDebtData\",\"components\":[{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotedTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotedTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_poolQuotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isExpired\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"remainingFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"loss\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationThresholds\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"lt\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ltParams\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"ltInitial\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ltFinal\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"timestampRampStart\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"rampDuration\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manageDebt\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumManageDebtAction\"}],\"outputs\":[{\"name\":\"newDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxEnabledTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"openCreditAccount\",\"inputs\":[{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolQuotaKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quotedTokensMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeAdapterAllowances\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revocations\",\"type\":\"tuple[]\",\"internalType\":\"structRevocationPair[]\",\"components\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setActiveCreditAccount\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCollateralTokenData\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ltInitial\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ltFinal\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"timestampRampStart\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"rampDuration\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractAllowance\",\"inputs\":[{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCreditConfigurator\",\"inputs\":[{\"name\":\"_creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCreditFacade\",\"inputs\":[{\"name\":\"_creditFacade\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"_feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFlagFor\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxEnabledTokens\",\"inputs\":[{\"name\":\"_maxEnabledTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceOracle\",\"inputs\":[{\"name\":\"_priceOracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setQuotedMask\",\"inputs\":[{\"name\":\"_quotedTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateQuota\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotaChange\",\"type\":\"int96\",\"internalType\":\"int96\"},{\"name\":\"minQuota\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"maxQuota\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SetCreditConfigurator\",\"inputs\":[{\"name\":\"newConfigurator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActiveCreditAccountNotSetException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ActiveCreditAccountOverridenException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AllowanceFailedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotAdapterException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CloseAccountWithNonZeroDebtException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CreditAccountDoesNotExistException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DebtToZeroWithActiveQuotasException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DebtUpdatedTwiceInOneBlockException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectParameterException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRemainingFundsException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughCollateralException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeTransferFromFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TargetContractNotAllowedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAlreadyAddedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotAllowedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooManyEnabledTokensException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooManyTokensException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UpdateQuotaOnZeroDebtAccountException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// CreditManagerv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditManagerv3MetaData.ABI instead.
var CreditManagerv3ABI = CreditManagerv3MetaData.ABI

// CreditManagerv3 is an auto generated Go binding around an Ethereum contract.
type CreditManagerv3 struct {
	CreditManagerv3Caller     // Read-only binding to the contract
	CreditManagerv3Transactor // Write-only binding to the contract
	CreditManagerv3Filterer   // Log filterer for contract events
}

// CreditManagerv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditManagerv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditManagerv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditManagerv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditManagerv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditManagerv3Session struct {
	Contract     *CreditManagerv3  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditManagerv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditManagerv3CallerSession struct {
	Contract *CreditManagerv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CreditManagerv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditManagerv3TransactorSession struct {
	Contract     *CreditManagerv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CreditManagerv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditManagerv3Raw struct {
	Contract *CreditManagerv3 // Generic contract binding to access the raw methods on
}

// CreditManagerv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditManagerv3CallerRaw struct {
	Contract *CreditManagerv3Caller // Generic read-only contract binding to access the raw methods on
}

// CreditManagerv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditManagerv3TransactorRaw struct {
	Contract *CreditManagerv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditManagerv3 creates a new instance of CreditManagerv3, bound to a specific deployed contract.
func NewCreditManagerv3(address common.Address, backend bind.ContractBackend) (*CreditManagerv3, error) {
	contract, err := bindCreditManagerv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv3{CreditManagerv3Caller: CreditManagerv3Caller{contract: contract}, CreditManagerv3Transactor: CreditManagerv3Transactor{contract: contract}, CreditManagerv3Filterer: CreditManagerv3Filterer{contract: contract}}, nil
}

// NewCreditManagerv3Caller creates a new read-only instance of CreditManagerv3, bound to a specific deployed contract.
func NewCreditManagerv3Caller(address common.Address, caller bind.ContractCaller) (*CreditManagerv3Caller, error) {
	contract, err := bindCreditManagerv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv3Caller{contract: contract}, nil
}

// NewCreditManagerv3Transactor creates a new write-only instance of CreditManagerv3, bound to a specific deployed contract.
func NewCreditManagerv3Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditManagerv3Transactor, error) {
	contract, err := bindCreditManagerv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv3Transactor{contract: contract}, nil
}

// NewCreditManagerv3Filterer creates a new log filterer instance of CreditManagerv3, bound to a specific deployed contract.
func NewCreditManagerv3Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditManagerv3Filterer, error) {
	contract, err := bindCreditManagerv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv3Filterer{contract: contract}, nil
}

// bindCreditManagerv3 binds a generic wrapper to an already deployed contract.
func bindCreditManagerv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditManagerv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv3 *CreditManagerv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv3.Contract.CreditManagerv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv3 *CreditManagerv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.CreditManagerv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv3 *CreditManagerv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.CreditManagerv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditManagerv3 *CreditManagerv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditManagerv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditManagerv3 *CreditManagerv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditManagerv3 *CreditManagerv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.contract.Transact(opts, method, params...)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) AccountFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "accountFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) AccountFactory() (common.Address, error) {
	return _CreditManagerv3.Contract.AccountFactory(&_CreditManagerv3.CallOpts)
}

// AccountFactory is a free data retrieval call binding the contract method 0x687cd9c1.
//
// Solidity: function accountFactory() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) AccountFactory() (common.Address, error) {
	return _CreditManagerv3.Contract.AccountFactory(&_CreditManagerv3.CallOpts)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) AdapterToContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "adapterToContract", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.AdapterToContract(&_CreditManagerv3.CallOpts, arg0)
}

// AdapterToContract is a free data retrieval call binding the contract method 0xff687543.
//
// Solidity: function adapterToContract(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) AdapterToContract(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.AdapterToContract(&_CreditManagerv3.CallOpts, arg0)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) AddressProvider() (common.Address, error) {
	return _CreditManagerv3.Contract.AddressProvider(&_CreditManagerv3.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) AddressProvider() (common.Address, error) {
	return _CreditManagerv3.Contract.AddressProvider(&_CreditManagerv3.CallOpts)
}

// CalcDebtAndCollateral is a free data retrieval call binding the contract method 0x0d334ca6.
//
// Solidity: function calcDebtAndCollateral(address creditAccount, uint8 task) view returns((uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) cdd)
func (_CreditManagerv3 *CreditManagerv3Caller) CalcDebtAndCollateral(opts *bind.CallOpts, creditAccount common.Address, task uint8) (CollateralDebtData, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "calcDebtAndCollateral", creditAccount, task)

	if err != nil {
		return *new(CollateralDebtData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollateralDebtData)).(*CollateralDebtData)

	return out0, err

}

// CalcDebtAndCollateral is a free data retrieval call binding the contract method 0x0d334ca6.
//
// Solidity: function calcDebtAndCollateral(address creditAccount, uint8 task) view returns((uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) cdd)
func (_CreditManagerv3 *CreditManagerv3Session) CalcDebtAndCollateral(creditAccount common.Address, task uint8) (CollateralDebtData, error) {
	return _CreditManagerv3.Contract.CalcDebtAndCollateral(&_CreditManagerv3.CallOpts, creditAccount, task)
}

// CalcDebtAndCollateral is a free data retrieval call binding the contract method 0x0d334ca6.
//
// Solidity: function calcDebtAndCollateral(address creditAccount, uint8 task) view returns((uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) cdd)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CalcDebtAndCollateral(creditAccount common.Address, task uint8) (CollateralDebtData, error) {
	return _CreditManagerv3.Contract.CalcDebtAndCollateral(&_CreditManagerv3.CallOpts, creditAccount, task)
}

// CollateralTokenByMask is a free data retrieval call binding the contract method 0x52c5fe11.
//
// Solidity: function collateralTokenByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv3 *CreditManagerv3Caller) CollateralTokenByMask(opts *bind.CallOpts, tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "collateralTokenByMask", tokenMask)

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

// CollateralTokenByMask is a free data retrieval call binding the contract method 0x52c5fe11.
//
// Solidity: function collateralTokenByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv3 *CreditManagerv3Session) CollateralTokenByMask(tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv3.Contract.CollateralTokenByMask(&_CreditManagerv3.CallOpts, tokenMask)
}

// CollateralTokenByMask is a free data retrieval call binding the contract method 0x52c5fe11.
//
// Solidity: function collateralTokenByMask(uint256 tokenMask) view returns(address token, uint16 liquidationThreshold)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CollateralTokenByMask(tokenMask *big.Int) (struct {
	Token                common.Address
	LiquidationThreshold uint16
}, error) {
	return _CreditManagerv3.Contract.CollateralTokenByMask(&_CreditManagerv3.CallOpts, tokenMask)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3Caller) CollateralTokensCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "collateralTokensCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3Session) CollateralTokensCount() (uint8, error) {
	return _CreditManagerv3.Contract.CollateralTokensCount(&_CreditManagerv3.CallOpts)
}

// CollateralTokensCount is a free data retrieval call binding the contract method 0x458936f5.
//
// Solidity: function collateralTokensCount() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CollateralTokensCount() (uint8, error) {
	return _CreditManagerv3.Contract.CollateralTokensCount(&_CreditManagerv3.CallOpts)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) ContractToAdapter(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "contractToAdapter", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.ContractToAdapter(&_CreditManagerv3.CallOpts, arg0)
}

// ContractToAdapter is a free data retrieval call binding the contract method 0xfdd57645.
//
// Solidity: function contractToAdapter(address ) view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) ContractToAdapter(arg0 common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.ContractToAdapter(&_CreditManagerv3.CallOpts, arg0)
}

// CreditAccountInfo is a free data retrieval call binding the contract method 0x3c5bc3b2.
//
// Solidity: function creditAccountInfo(address ) view returns(uint256 debt, uint256 cumulativeIndexLastUpdate, uint128 cumulativeQuotaInterest, uint128 quotaFees, uint256 enabledTokensMask, uint16 flags, uint64 lastDebtUpdate, address borrower)
func (_CreditManagerv3 *CreditManagerv3Caller) CreditAccountInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Debt                      *big.Int
	CumulativeIndexLastUpdate *big.Int
	CumulativeQuotaInterest   *big.Int
	QuotaFees                 *big.Int
	EnabledTokensMask         *big.Int
	Flags                     uint16
	LastDebtUpdate            uint64
	Borrower                  common.Address
}, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditAccountInfo", arg0)

	outstruct := new(struct {
		Debt                      *big.Int
		CumulativeIndexLastUpdate *big.Int
		CumulativeQuotaInterest   *big.Int
		QuotaFees                 *big.Int
		EnabledTokensMask         *big.Int
		Flags                     uint16
		LastDebtUpdate            uint64
		Borrower                  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CumulativeIndexLastUpdate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CumulativeQuotaInterest = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.QuotaFees = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EnabledTokensMask = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Flags = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.LastDebtUpdate = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.Borrower = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CreditAccountInfo is a free data retrieval call binding the contract method 0x3c5bc3b2.
//
// Solidity: function creditAccountInfo(address ) view returns(uint256 debt, uint256 cumulativeIndexLastUpdate, uint128 cumulativeQuotaInterest, uint128 quotaFees, uint256 enabledTokensMask, uint16 flags, uint64 lastDebtUpdate, address borrower)
func (_CreditManagerv3 *CreditManagerv3Session) CreditAccountInfo(arg0 common.Address) (struct {
	Debt                      *big.Int
	CumulativeIndexLastUpdate *big.Int
	CumulativeQuotaInterest   *big.Int
	QuotaFees                 *big.Int
	EnabledTokensMask         *big.Int
	Flags                     uint16
	LastDebtUpdate            uint64
	Borrower                  common.Address
}, error) {
	return _CreditManagerv3.Contract.CreditAccountInfo(&_CreditManagerv3.CallOpts, arg0)
}

// CreditAccountInfo is a free data retrieval call binding the contract method 0x3c5bc3b2.
//
// Solidity: function creditAccountInfo(address ) view returns(uint256 debt, uint256 cumulativeIndexLastUpdate, uint128 cumulativeQuotaInterest, uint128 quotaFees, uint256 enabledTokensMask, uint16 flags, uint64 lastDebtUpdate, address borrower)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditAccountInfo(arg0 common.Address) (struct {
	Debt                      *big.Int
	CumulativeIndexLastUpdate *big.Int
	CumulativeQuotaInterest   *big.Int
	QuotaFees                 *big.Int
	EnabledTokensMask         *big.Int
	Flags                     uint16
	LastDebtUpdate            uint64
	Borrower                  common.Address
}, error) {
	return _CreditManagerv3.Contract.CreditAccountInfo(&_CreditManagerv3.CallOpts, arg0)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x2c9db6f1.
//
// Solidity: function creditAccounts(uint256 offset, uint256 limit) view returns(address[] result)
func (_CreditManagerv3 *CreditManagerv3Caller) CreditAccounts(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditAccounts", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CreditAccounts is a free data retrieval call binding the contract method 0x2c9db6f1.
//
// Solidity: function creditAccounts(uint256 offset, uint256 limit) view returns(address[] result)
func (_CreditManagerv3 *CreditManagerv3Session) CreditAccounts(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _CreditManagerv3.Contract.CreditAccounts(&_CreditManagerv3.CallOpts, offset, limit)
}

// CreditAccounts is a free data retrieval call binding the contract method 0x2c9db6f1.
//
// Solidity: function creditAccounts(uint256 offset, uint256 limit) view returns(address[] result)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditAccounts(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _CreditManagerv3.Contract.CreditAccounts(&_CreditManagerv3.CallOpts, offset, limit)
}

// CreditAccounts0 is a free data retrieval call binding the contract method 0x741f3e3c.
//
// Solidity: function creditAccounts() view returns(address[])
func (_CreditManagerv3 *CreditManagerv3Caller) CreditAccounts0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditAccounts0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CreditAccounts0 is a free data retrieval call binding the contract method 0x741f3e3c.
//
// Solidity: function creditAccounts() view returns(address[])
func (_CreditManagerv3 *CreditManagerv3Session) CreditAccounts0() ([]common.Address, error) {
	return _CreditManagerv3.Contract.CreditAccounts0(&_CreditManagerv3.CallOpts)
}

// CreditAccounts0 is a free data retrieval call binding the contract method 0x741f3e3c.
//
// Solidity: function creditAccounts() view returns(address[])
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditAccounts0() ([]common.Address, error) {
	return _CreditManagerv3.Contract.CreditAccounts0(&_CreditManagerv3.CallOpts)
}

// CreditAccountsLen is a free data retrieval call binding the contract method 0xf13d0fc6.
//
// Solidity: function creditAccountsLen() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Caller) CreditAccountsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditAccountsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditAccountsLen is a free data retrieval call binding the contract method 0xf13d0fc6.
//
// Solidity: function creditAccountsLen() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Session) CreditAccountsLen() (*big.Int, error) {
	return _CreditManagerv3.Contract.CreditAccountsLen(&_CreditManagerv3.CallOpts)
}

// CreditAccountsLen is a free data retrieval call binding the contract method 0xf13d0fc6.
//
// Solidity: function creditAccountsLen() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditAccountsLen() (*big.Int, error) {
	return _CreditManagerv3.Contract.CreditAccountsLen(&_CreditManagerv3.CallOpts)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) CreditConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv3.Contract.CreditConfigurator(&_CreditManagerv3.CallOpts)
}

// CreditConfigurator is a free data retrieval call binding the contract method 0xf9aa028a.
//
// Solidity: function creditConfigurator() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditConfigurator() (common.Address, error) {
	return _CreditManagerv3.Contract.CreditConfigurator(&_CreditManagerv3.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) CreditFacade() (common.Address, error) {
	return _CreditManagerv3.Contract.CreditFacade(&_CreditManagerv3.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) CreditFacade() (common.Address, error) {
	return _CreditManagerv3.Contract.CreditFacade(&_CreditManagerv3.CallOpts)
}

// EnabledTokensMaskOf is a free data retrieval call binding the contract method 0xf9f0ca66.
//
// Solidity: function enabledTokensMaskOf(address creditAccount) view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Caller) EnabledTokensMaskOf(opts *bind.CallOpts, creditAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "enabledTokensMaskOf", creditAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledTokensMaskOf is a free data retrieval call binding the contract method 0xf9f0ca66.
//
// Solidity: function enabledTokensMaskOf(address creditAccount) view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Session) EnabledTokensMaskOf(creditAccount common.Address) (*big.Int, error) {
	return _CreditManagerv3.Contract.EnabledTokensMaskOf(&_CreditManagerv3.CallOpts, creditAccount)
}

// EnabledTokensMaskOf is a free data retrieval call binding the contract method 0xf9f0ca66.
//
// Solidity: function enabledTokensMaskOf(address creditAccount) view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3CallerSession) EnabledTokensMaskOf(creditAccount common.Address) (*big.Int, error) {
	return _CreditManagerv3.Contract.EnabledTokensMaskOf(&_CreditManagerv3.CallOpts, creditAccount)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired)
func (_CreditManagerv3 *CreditManagerv3Caller) Fees(opts *bind.CallOpts) (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "fees")

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
// Solidity: function fees() view returns(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired)
func (_CreditManagerv3 *CreditManagerv3Session) Fees() (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	return _CreditManagerv3.Contract.Fees(&_CreditManagerv3.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired)
func (_CreditManagerv3 *CreditManagerv3CallerSession) Fees() (struct {
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}, error) {
	return _CreditManagerv3.Contract.Fees(&_CreditManagerv3.CallOpts)
}

// FlagsOf is a free data retrieval call binding the contract method 0x845104de.
//
// Solidity: function flagsOf(address creditAccount) view returns(uint16)
func (_CreditManagerv3 *CreditManagerv3Caller) FlagsOf(opts *bind.CallOpts, creditAccount common.Address) (uint16, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "flagsOf", creditAccount)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FlagsOf is a free data retrieval call binding the contract method 0x845104de.
//
// Solidity: function flagsOf(address creditAccount) view returns(uint16)
func (_CreditManagerv3 *CreditManagerv3Session) FlagsOf(creditAccount common.Address) (uint16, error) {
	return _CreditManagerv3.Contract.FlagsOf(&_CreditManagerv3.CallOpts, creditAccount)
}

// FlagsOf is a free data retrieval call binding the contract method 0x845104de.
//
// Solidity: function flagsOf(address creditAccount) view returns(uint16)
func (_CreditManagerv3 *CreditManagerv3CallerSession) FlagsOf(creditAccount common.Address) (uint16, error) {
	return _CreditManagerv3.Contract.FlagsOf(&_CreditManagerv3.CallOpts, creditAccount)
}

// GetActiveCreditAccountOrRevert is a free data retrieval call binding the contract method 0x34878f54.
//
// Solidity: function getActiveCreditAccountOrRevert() view returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3Caller) GetActiveCreditAccountOrRevert(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "getActiveCreditAccountOrRevert")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetActiveCreditAccountOrRevert is a free data retrieval call binding the contract method 0x34878f54.
//
// Solidity: function getActiveCreditAccountOrRevert() view returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3Session) GetActiveCreditAccountOrRevert() (common.Address, error) {
	return _CreditManagerv3.Contract.GetActiveCreditAccountOrRevert(&_CreditManagerv3.CallOpts)
}

// GetActiveCreditAccountOrRevert is a free data retrieval call binding the contract method 0x34878f54.
//
// Solidity: function getActiveCreditAccountOrRevert() view returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3CallerSession) GetActiveCreditAccountOrRevert() (common.Address, error) {
	return _CreditManagerv3.Contract.GetActiveCreditAccountOrRevert(&_CreditManagerv3.CallOpts)
}

// GetBorrowerOrRevert is a free data retrieval call binding the contract method 0xc53afb1e.
//
// Solidity: function getBorrowerOrRevert(address creditAccount) view returns(address borrower)
func (_CreditManagerv3 *CreditManagerv3Caller) GetBorrowerOrRevert(opts *bind.CallOpts, creditAccount common.Address) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "getBorrowerOrRevert", creditAccount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowerOrRevert is a free data retrieval call binding the contract method 0xc53afb1e.
//
// Solidity: function getBorrowerOrRevert(address creditAccount) view returns(address borrower)
func (_CreditManagerv3 *CreditManagerv3Session) GetBorrowerOrRevert(creditAccount common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.GetBorrowerOrRevert(&_CreditManagerv3.CallOpts, creditAccount)
}

// GetBorrowerOrRevert is a free data retrieval call binding the contract method 0xc53afb1e.
//
// Solidity: function getBorrowerOrRevert(address creditAccount) view returns(address borrower)
func (_CreditManagerv3 *CreditManagerv3CallerSession) GetBorrowerOrRevert(creditAccount common.Address) (common.Address, error) {
	return _CreditManagerv3.Contract.GetBorrowerOrRevert(&_CreditManagerv3.CallOpts, creditAccount)
}

// GetTokenByMask is a free data retrieval call binding the contract method 0x4fc0e3a8.
//
// Solidity: function getTokenByMask(uint256 tokenMask) view returns(address token)
func (_CreditManagerv3 *CreditManagerv3Caller) GetTokenByMask(opts *bind.CallOpts, tokenMask *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "getTokenByMask", tokenMask)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenByMask is a free data retrieval call binding the contract method 0x4fc0e3a8.
//
// Solidity: function getTokenByMask(uint256 tokenMask) view returns(address token)
func (_CreditManagerv3 *CreditManagerv3Session) GetTokenByMask(tokenMask *big.Int) (common.Address, error) {
	return _CreditManagerv3.Contract.GetTokenByMask(&_CreditManagerv3.CallOpts, tokenMask)
}

// GetTokenByMask is a free data retrieval call binding the contract method 0x4fc0e3a8.
//
// Solidity: function getTokenByMask(uint256 tokenMask) view returns(address token)
func (_CreditManagerv3 *CreditManagerv3CallerSession) GetTokenByMask(tokenMask *big.Int) (common.Address, error) {
	return _CreditManagerv3.Contract.GetTokenByMask(&_CreditManagerv3.CallOpts, tokenMask)
}

// GetTokenMaskOrRevert is a free data retrieval call binding the contract method 0xd5c2f486.
//
// Solidity: function getTokenMaskOrRevert(address token) view returns(uint256 tokenMask)
func (_CreditManagerv3 *CreditManagerv3Caller) GetTokenMaskOrRevert(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "getTokenMaskOrRevert", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMaskOrRevert is a free data retrieval call binding the contract method 0xd5c2f486.
//
// Solidity: function getTokenMaskOrRevert(address token) view returns(uint256 tokenMask)
func (_CreditManagerv3 *CreditManagerv3Session) GetTokenMaskOrRevert(token common.Address) (*big.Int, error) {
	return _CreditManagerv3.Contract.GetTokenMaskOrRevert(&_CreditManagerv3.CallOpts, token)
}

// GetTokenMaskOrRevert is a free data retrieval call binding the contract method 0xd5c2f486.
//
// Solidity: function getTokenMaskOrRevert(address token) view returns(uint256 tokenMask)
func (_CreditManagerv3 *CreditManagerv3CallerSession) GetTokenMaskOrRevert(token common.Address) (*big.Int, error) {
	return _CreditManagerv3.Contract.GetTokenMaskOrRevert(&_CreditManagerv3.CallOpts, token)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x8340e24d.
//
// Solidity: function isLiquidatable(address creditAccount, uint16 minHealthFactor) view returns(bool)
func (_CreditManagerv3 *CreditManagerv3Caller) IsLiquidatable(opts *bind.CallOpts, creditAccount common.Address, minHealthFactor uint16) (bool, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "isLiquidatable", creditAccount, minHealthFactor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidatable is a free data retrieval call binding the contract method 0x8340e24d.
//
// Solidity: function isLiquidatable(address creditAccount, uint16 minHealthFactor) view returns(bool)
func (_CreditManagerv3 *CreditManagerv3Session) IsLiquidatable(creditAccount common.Address, minHealthFactor uint16) (bool, error) {
	return _CreditManagerv3.Contract.IsLiquidatable(&_CreditManagerv3.CallOpts, creditAccount, minHealthFactor)
}

// IsLiquidatable is a free data retrieval call binding the contract method 0x8340e24d.
//
// Solidity: function isLiquidatable(address creditAccount, uint16 minHealthFactor) view returns(bool)
func (_CreditManagerv3 *CreditManagerv3CallerSession) IsLiquidatable(creditAccount common.Address, minHealthFactor uint16) (bool, error) {
	return _CreditManagerv3.Contract.IsLiquidatable(&_CreditManagerv3.CallOpts, creditAccount, minHealthFactor)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv3 *CreditManagerv3Caller) LiquidationThresholds(opts *bind.CallOpts, token common.Address) (uint16, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "liquidationThresholds", token)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv3 *CreditManagerv3Session) LiquidationThresholds(token common.Address) (uint16, error) {
	return _CreditManagerv3.Contract.LiquidationThresholds(&_CreditManagerv3.CallOpts, token)
}

// LiquidationThresholds is a free data retrieval call binding the contract method 0x78327438.
//
// Solidity: function liquidationThresholds(address token) view returns(uint16 lt)
func (_CreditManagerv3 *CreditManagerv3CallerSession) LiquidationThresholds(token common.Address) (uint16, error) {
	return _CreditManagerv3.Contract.LiquidationThresholds(&_CreditManagerv3.CallOpts, token)
}

// LtParams is a free data retrieval call binding the contract method 0x3201de4c.
//
// Solidity: function ltParams(address token) view returns(uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration)
func (_CreditManagerv3 *CreditManagerv3Caller) LtParams(opts *bind.CallOpts, token common.Address) (struct {
	LtInitial          uint16
	LtFinal            uint16
	TimestampRampStart *big.Int
	RampDuration       *big.Int
}, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "ltParams", token)

	outstruct := new(struct {
		LtInitial          uint16
		LtFinal            uint16
		TimestampRampStart *big.Int
		RampDuration       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LtInitial = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.LtFinal = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.TimestampRampStart = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RampDuration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LtParams is a free data retrieval call binding the contract method 0x3201de4c.
//
// Solidity: function ltParams(address token) view returns(uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration)
func (_CreditManagerv3 *CreditManagerv3Session) LtParams(token common.Address) (struct {
	LtInitial          uint16
	LtFinal            uint16
	TimestampRampStart *big.Int
	RampDuration       *big.Int
}, error) {
	return _CreditManagerv3.Contract.LtParams(&_CreditManagerv3.CallOpts, token)
}

// LtParams is a free data retrieval call binding the contract method 0x3201de4c.
//
// Solidity: function ltParams(address token) view returns(uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration)
func (_CreditManagerv3 *CreditManagerv3CallerSession) LtParams(token common.Address) (struct {
	LtInitial          uint16
	LtFinal            uint16
	TimestampRampStart *big.Int
	RampDuration       *big.Int
}, error) {
	return _CreditManagerv3.Contract.LtParams(&_CreditManagerv3.CallOpts, token)
}

// MaxEnabledTokens is a free data retrieval call binding the contract method 0x2c521c06.
//
// Solidity: function maxEnabledTokens() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3Caller) MaxEnabledTokens(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "maxEnabledTokens")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxEnabledTokens is a free data retrieval call binding the contract method 0x2c521c06.
//
// Solidity: function maxEnabledTokens() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3Session) MaxEnabledTokens() (uint8, error) {
	return _CreditManagerv3.Contract.MaxEnabledTokens(&_CreditManagerv3.CallOpts)
}

// MaxEnabledTokens is a free data retrieval call binding the contract method 0x2c521c06.
//
// Solidity: function maxEnabledTokens() view returns(uint8)
func (_CreditManagerv3 *CreditManagerv3CallerSession) MaxEnabledTokens() (uint8, error) {
	return _CreditManagerv3.Contract.MaxEnabledTokens(&_CreditManagerv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreditManagerv3 *CreditManagerv3Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreditManagerv3 *CreditManagerv3Session) Name() (string, error) {
	return _CreditManagerv3.Contract.Name(&_CreditManagerv3.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CreditManagerv3 *CreditManagerv3CallerSession) Name() (string, error) {
	return _CreditManagerv3.Contract.Name(&_CreditManagerv3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) Pool() (common.Address, error) {
	return _CreditManagerv3.Contract.Pool(&_CreditManagerv3.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) Pool() (common.Address, error) {
	return _CreditManagerv3.Contract.Pool(&_CreditManagerv3.CallOpts)
}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) PoolQuotaKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "poolQuotaKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) PoolQuotaKeeper() (common.Address, error) {
	return _CreditManagerv3.Contract.PoolQuotaKeeper(&_CreditManagerv3.CallOpts)
}

// PoolQuotaKeeper is a free data retrieval call binding the contract method 0xbe8da14b.
//
// Solidity: function poolQuotaKeeper() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) PoolQuotaKeeper() (common.Address, error) {
	return _CreditManagerv3.Contract.PoolQuotaKeeper(&_CreditManagerv3.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) PriceOracle() (common.Address, error) {
	return _CreditManagerv3.Contract.PriceOracle(&_CreditManagerv3.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) PriceOracle() (common.Address, error) {
	return _CreditManagerv3.Contract.PriceOracle(&_CreditManagerv3.CallOpts)
}

// QuotedTokensMask is a free data retrieval call binding the contract method 0xe0f83824.
//
// Solidity: function quotedTokensMask() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Caller) QuotedTokensMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "quotedTokensMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotedTokensMask is a free data retrieval call binding the contract method 0xe0f83824.
//
// Solidity: function quotedTokensMask() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Session) QuotedTokensMask() (*big.Int, error) {
	return _CreditManagerv3.Contract.QuotedTokensMask(&_CreditManagerv3.CallOpts)
}

// QuotedTokensMask is a free data retrieval call binding the contract method 0xe0f83824.
//
// Solidity: function quotedTokensMask() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3CallerSession) QuotedTokensMask() (*big.Int, error) {
	return _CreditManagerv3.Contract.QuotedTokensMask(&_CreditManagerv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv3 *CreditManagerv3Session) Underlying() (common.Address, error) {
	return _CreditManagerv3.Contract.Underlying(&_CreditManagerv3.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CreditManagerv3 *CreditManagerv3CallerSession) Underlying() (common.Address, error) {
	return _CreditManagerv3.Contract.Underlying(&_CreditManagerv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditManagerv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3Session) Version() (*big.Int, error) {
	return _CreditManagerv3.Contract.Version(&_CreditManagerv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditManagerv3 *CreditManagerv3CallerSession) Version() (*big.Int, error) {
	return _CreditManagerv3.Contract.Version(&_CreditManagerv3.CallOpts)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns(uint256 tokensToEnable)
func (_CreditManagerv3 *CreditManagerv3Transactor) AddCollateral(opts *bind.TransactOpts, payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "addCollateral", payer, creditAccount, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns(uint256 tokensToEnable)
func (_CreditManagerv3 *CreditManagerv3Session) AddCollateral(payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.AddCollateral(&_CreditManagerv3.TransactOpts, payer, creditAccount, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x830aa745.
//
// Solidity: function addCollateral(address payer, address creditAccount, address token, uint256 amount) returns(uint256 tokensToEnable)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) AddCollateral(payer common.Address, creditAccount common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.AddCollateral(&_CreditManagerv3.TransactOpts, payer, creditAccount, token, amount)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv3 *CreditManagerv3Session) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.AddToken(&_CreditManagerv3.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.AddToken(&_CreditManagerv3.TransactOpts, token)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0xfa30b30f.
//
// Solidity: function approveCreditAccount(address token, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) ApproveCreditAccount(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "approveCreditAccount", token, amount)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0xfa30b30f.
//
// Solidity: function approveCreditAccount(address token, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3Session) ApproveCreditAccount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ApproveCreditAccount(&_CreditManagerv3.TransactOpts, token, amount)
}

// ApproveCreditAccount is a paid mutator transaction binding the contract method 0xfa30b30f.
//
// Solidity: function approveCreditAccount(address token, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) ApproveCreditAccount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ApproveCreditAccount(&_CreditManagerv3.TransactOpts, token, amount)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x12ef080d.
//
// Solidity: function approveToken(address creditAccount, address token, address spender, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) ApproveToken(opts *bind.TransactOpts, creditAccount common.Address, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "approveToken", creditAccount, token, spender, amount)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x12ef080d.
//
// Solidity: function approveToken(address creditAccount, address token, address spender, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3Session) ApproveToken(creditAccount common.Address, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ApproveToken(&_CreditManagerv3.TransactOpts, creditAccount, token, spender, amount)
}

// ApproveToken is a paid mutator transaction binding the contract method 0x12ef080d.
//
// Solidity: function approveToken(address creditAccount, address token, address spender, uint256 amount) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) ApproveToken(creditAccount common.Address, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ApproveToken(&_CreditManagerv3.TransactOpts, creditAccount, token, spender, amount)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x7687f670.
//
// Solidity: function closeCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) CloseCreditAccount(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "closeCreditAccount", creditAccount)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x7687f670.
//
// Solidity: function closeCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3Session) CloseCreditAccount(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.CloseCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x7687f670.
//
// Solidity: function closeCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) CloseCreditAccount(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.CloseCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes data) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3Transactor) Execute(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "execute", data)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes data) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3Session) Execute(data []byte) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.Execute(&_CreditManagerv3.TransactOpts, data)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes data) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) Execute(data []byte) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.Execute(&_CreditManagerv3.TransactOpts, data)
}

// ExternalCall is a paid mutator transaction binding the contract method 0xeb23d33e.
//
// Solidity: function externalCall(address creditAccount, address target, bytes callData) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3Transactor) ExternalCall(opts *bind.TransactOpts, creditAccount common.Address, target common.Address, callData []byte) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "externalCall", creditAccount, target, callData)
}

// ExternalCall is a paid mutator transaction binding the contract method 0xeb23d33e.
//
// Solidity: function externalCall(address creditAccount, address target, bytes callData) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3Session) ExternalCall(creditAccount common.Address, target common.Address, callData []byte) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ExternalCall(&_CreditManagerv3.TransactOpts, creditAccount, target, callData)
}

// ExternalCall is a paid mutator transaction binding the contract method 0xeb23d33e.
//
// Solidity: function externalCall(address creditAccount, address target, bytes callData) returns(bytes result)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) ExternalCall(creditAccount common.Address, target common.Address, callData []byte) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ExternalCall(&_CreditManagerv3.TransactOpts, creditAccount, target, callData)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x3d7e5dc4.
//
// Solidity: function fullCollateralCheck(address creditAccount, uint256 enabledTokensMask, uint256[] collateralHints, uint16 minHealthFactor, bool useSafePrices) returns(uint256 enabledTokensMaskAfter)
func (_CreditManagerv3 *CreditManagerv3Transactor) FullCollateralCheck(opts *bind.TransactOpts, creditAccount common.Address, enabledTokensMask *big.Int, collateralHints []*big.Int, minHealthFactor uint16, useSafePrices bool) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "fullCollateralCheck", creditAccount, enabledTokensMask, collateralHints, minHealthFactor, useSafePrices)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x3d7e5dc4.
//
// Solidity: function fullCollateralCheck(address creditAccount, uint256 enabledTokensMask, uint256[] collateralHints, uint16 minHealthFactor, bool useSafePrices) returns(uint256 enabledTokensMaskAfter)
func (_CreditManagerv3 *CreditManagerv3Session) FullCollateralCheck(creditAccount common.Address, enabledTokensMask *big.Int, collateralHints []*big.Int, minHealthFactor uint16, useSafePrices bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.FullCollateralCheck(&_CreditManagerv3.TransactOpts, creditAccount, enabledTokensMask, collateralHints, minHealthFactor, useSafePrices)
}

// FullCollateralCheck is a paid mutator transaction binding the contract method 0x3d7e5dc4.
//
// Solidity: function fullCollateralCheck(address creditAccount, uint256 enabledTokensMask, uint256[] collateralHints, uint16 minHealthFactor, bool useSafePrices) returns(uint256 enabledTokensMaskAfter)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) FullCollateralCheck(creditAccount common.Address, enabledTokensMask *big.Int, collateralHints []*big.Int, minHealthFactor uint16, useSafePrices bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.FullCollateralCheck(&_CreditManagerv3.TransactOpts, creditAccount, enabledTokensMask, collateralHints, minHealthFactor, useSafePrices)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe2f1490f.
//
// Solidity: function liquidateCreditAccount(address creditAccount, (uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) collateralDebtData, address to, bool isExpired) returns(uint256 remainingFunds, uint256 loss)
func (_CreditManagerv3 *CreditManagerv3Transactor) LiquidateCreditAccount(opts *bind.TransactOpts, creditAccount common.Address, collateralDebtData CollateralDebtData, to common.Address, isExpired bool) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "liquidateCreditAccount", creditAccount, collateralDebtData, to, isExpired)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe2f1490f.
//
// Solidity: function liquidateCreditAccount(address creditAccount, (uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) collateralDebtData, address to, bool isExpired) returns(uint256 remainingFunds, uint256 loss)
func (_CreditManagerv3 *CreditManagerv3Session) LiquidateCreditAccount(creditAccount common.Address, collateralDebtData CollateralDebtData, to common.Address, isExpired bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.LiquidateCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount, collateralDebtData, to, isExpired)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe2f1490f.
//
// Solidity: function liquidateCreditAccount(address creditAccount, (uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address[],address) collateralDebtData, address to, bool isExpired) returns(uint256 remainingFunds, uint256 loss)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) LiquidateCreditAccount(creditAccount common.Address, collateralDebtData CollateralDebtData, to common.Address, isExpired bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.LiquidateCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount, collateralDebtData, to, isExpired)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x432017e8.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, uint256 enabledTokensMask, uint8 action) returns(uint256 newDebt, uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Transactor) ManageDebt(opts *bind.TransactOpts, creditAccount common.Address, amount *big.Int, enabledTokensMask *big.Int, action uint8) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "manageDebt", creditAccount, amount, enabledTokensMask, action)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x432017e8.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, uint256 enabledTokensMask, uint8 action) returns(uint256 newDebt, uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Session) ManageDebt(creditAccount common.Address, amount *big.Int, enabledTokensMask *big.Int, action uint8) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ManageDebt(&_CreditManagerv3.TransactOpts, creditAccount, amount, enabledTokensMask, action)
}

// ManageDebt is a paid mutator transaction binding the contract method 0x432017e8.
//
// Solidity: function manageDebt(address creditAccount, uint256 amount, uint256 enabledTokensMask, uint8 action) returns(uint256 newDebt, uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) ManageDebt(creditAccount common.Address, amount *big.Int, enabledTokensMask *big.Int, action uint8) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.ManageDebt(&_CreditManagerv3.TransactOpts, creditAccount, amount, enabledTokensMask, action)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x43fe7bbe.
//
// Solidity: function openCreditAccount(address onBehalfOf) returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3Transactor) OpenCreditAccount(opts *bind.TransactOpts, onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "openCreditAccount", onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x43fe7bbe.
//
// Solidity: function openCreditAccount(address onBehalfOf) returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3Session) OpenCreditAccount(onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.OpenCreditAccount(&_CreditManagerv3.TransactOpts, onBehalfOf)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x43fe7bbe.
//
// Solidity: function openCreditAccount(address onBehalfOf) returns(address creditAccount)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) OpenCreditAccount(onBehalfOf common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.OpenCreditAccount(&_CreditManagerv3.TransactOpts, onBehalfOf)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x3031b01a.
//
// Solidity: function revokeAdapterAllowances(address creditAccount, (address,address)[] revocations) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) RevokeAdapterAllowances(opts *bind.TransactOpts, creditAccount common.Address, revocations []RevocationPair) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "revokeAdapterAllowances", creditAccount, revocations)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x3031b01a.
//
// Solidity: function revokeAdapterAllowances(address creditAccount, (address,address)[] revocations) returns()
func (_CreditManagerv3 *CreditManagerv3Session) RevokeAdapterAllowances(creditAccount common.Address, revocations []RevocationPair) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.RevokeAdapterAllowances(&_CreditManagerv3.TransactOpts, creditAccount, revocations)
}

// RevokeAdapterAllowances is a paid mutator transaction binding the contract method 0x3031b01a.
//
// Solidity: function revokeAdapterAllowances(address creditAccount, (address,address)[] revocations) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) RevokeAdapterAllowances(creditAccount common.Address, revocations []RevocationPair) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.RevokeAdapterAllowances(&_CreditManagerv3.TransactOpts, creditAccount, revocations)
}

// SetActiveCreditAccount is a paid mutator transaction binding the contract method 0xe09357a4.
//
// Solidity: function setActiveCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetActiveCreditAccount(opts *bind.TransactOpts, creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setActiveCreditAccount", creditAccount)
}

// SetActiveCreditAccount is a paid mutator transaction binding the contract method 0xe09357a4.
//
// Solidity: function setActiveCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetActiveCreditAccount(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetActiveCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount)
}

// SetActiveCreditAccount is a paid mutator transaction binding the contract method 0xe09357a4.
//
// Solidity: function setActiveCreditAccount(address creditAccount) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetActiveCreditAccount(creditAccount common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetActiveCreditAccount(&_CreditManagerv3.TransactOpts, creditAccount)
}

// SetCollateralTokenData is a paid mutator transaction binding the contract method 0x2f232138.
//
// Solidity: function setCollateralTokenData(address token, uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetCollateralTokenData(opts *bind.TransactOpts, token common.Address, ltInitial uint16, ltFinal uint16, timestampRampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setCollateralTokenData", token, ltInitial, ltFinal, timestampRampStart, rampDuration)
}

// SetCollateralTokenData is a paid mutator transaction binding the contract method 0x2f232138.
//
// Solidity: function setCollateralTokenData(address token, uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetCollateralTokenData(token common.Address, ltInitial uint16, ltFinal uint16, timestampRampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCollateralTokenData(&_CreditManagerv3.TransactOpts, token, ltInitial, ltFinal, timestampRampStart, rampDuration)
}

// SetCollateralTokenData is a paid mutator transaction binding the contract method 0x2f232138.
//
// Solidity: function setCollateralTokenData(address token, uint16 ltInitial, uint16 ltFinal, uint40 timestampRampStart, uint24 rampDuration) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetCollateralTokenData(token common.Address, ltInitial uint16, ltFinal uint16, timestampRampStart *big.Int, rampDuration *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCollateralTokenData(&_CreditManagerv3.TransactOpts, token, ltInitial, ltFinal, timestampRampStart, rampDuration)
}

// SetContractAllowance is a paid mutator transaction binding the contract method 0x80213c74.
//
// Solidity: function setContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetContractAllowance(opts *bind.TransactOpts, adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setContractAllowance", adapter, targetContract)
}

// SetContractAllowance is a paid mutator transaction binding the contract method 0x80213c74.
//
// Solidity: function setContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetContractAllowance(&_CreditManagerv3.TransactOpts, adapter, targetContract)
}

// SetContractAllowance is a paid mutator transaction binding the contract method 0x80213c74.
//
// Solidity: function setContractAllowance(address adapter, address targetContract) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetContractAllowance(adapter common.Address, targetContract common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetContractAllowance(&_CreditManagerv3.TransactOpts, adapter, targetContract)
}

// SetCreditConfigurator is a paid mutator transaction binding the contract method 0xf30ba499.
//
// Solidity: function setCreditConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetCreditConfigurator(opts *bind.TransactOpts, _creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setCreditConfigurator", _creditConfigurator)
}

// SetCreditConfigurator is a paid mutator transaction binding the contract method 0xf30ba499.
//
// Solidity: function setCreditConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetCreditConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCreditConfigurator(&_CreditManagerv3.TransactOpts, _creditConfigurator)
}

// SetCreditConfigurator is a paid mutator transaction binding the contract method 0xf30ba499.
//
// Solidity: function setCreditConfigurator(address _creditConfigurator) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetCreditConfigurator(_creditConfigurator common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCreditConfigurator(&_CreditManagerv3.TransactOpts, _creditConfigurator)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0xacb0e845.
//
// Solidity: function setCreditFacade(address _creditFacade) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetCreditFacade(opts *bind.TransactOpts, _creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setCreditFacade", _creditFacade)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0xacb0e845.
//
// Solidity: function setCreditFacade(address _creditFacade) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCreditFacade(&_CreditManagerv3.TransactOpts, _creditFacade)
}

// SetCreditFacade is a paid mutator transaction binding the contract method 0xacb0e845.
//
// Solidity: function setCreditFacade(address _creditFacade) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetCreditFacade(_creditFacade common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetCreditFacade(&_CreditManagerv3.TransactOpts, _creditFacade)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetFees(opts *bind.TransactOpts, _feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setFees", _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetFees(_feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetFees(&_CreditManagerv3.TransactOpts, _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// SetFees is a paid mutator transaction binding the contract method 0xf206d32a.
//
// Solidity: function setFees(uint16 _feeInterest, uint16 _feeLiquidation, uint16 _liquidationDiscount, uint16 _feeLiquidationExpired, uint16 _liquidationDiscountExpired) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetFees(_feeInterest uint16, _feeLiquidation uint16, _liquidationDiscount uint16, _feeLiquidationExpired uint16, _liquidationDiscountExpired uint16) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetFees(&_CreditManagerv3.TransactOpts, _feeInterest, _feeLiquidation, _liquidationDiscount, _feeLiquidationExpired, _liquidationDiscountExpired)
}

// SetFlagFor is a paid mutator transaction binding the contract method 0xc544cf7d.
//
// Solidity: function setFlagFor(address creditAccount, uint16 flag, bool value) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetFlagFor(opts *bind.TransactOpts, creditAccount common.Address, flag uint16, value bool) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setFlagFor", creditAccount, flag, value)
}

// SetFlagFor is a paid mutator transaction binding the contract method 0xc544cf7d.
//
// Solidity: function setFlagFor(address creditAccount, uint16 flag, bool value) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetFlagFor(creditAccount common.Address, flag uint16, value bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetFlagFor(&_CreditManagerv3.TransactOpts, creditAccount, flag, value)
}

// SetFlagFor is a paid mutator transaction binding the contract method 0xc544cf7d.
//
// Solidity: function setFlagFor(address creditAccount, uint16 flag, bool value) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetFlagFor(creditAccount common.Address, flag uint16, value bool) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetFlagFor(&_CreditManagerv3.TransactOpts, creditAccount, flag, value)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 _maxEnabledTokens) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetMaxEnabledTokens(opts *bind.TransactOpts, _maxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setMaxEnabledTokens", _maxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 _maxEnabledTokens) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetMaxEnabledTokens(_maxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetMaxEnabledTokens(&_CreditManagerv3.TransactOpts, _maxEnabledTokens)
}

// SetMaxEnabledTokens is a paid mutator transaction binding the contract method 0xdc2b21c1.
//
// Solidity: function setMaxEnabledTokens(uint8 _maxEnabledTokens) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetMaxEnabledTokens(_maxEnabledTokens uint8) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetMaxEnabledTokens(&_CreditManagerv3.TransactOpts, _maxEnabledTokens)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetPriceOracle(opts *bind.TransactOpts, _priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setPriceOracle", _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetPriceOracle(&_CreditManagerv3.TransactOpts, _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetPriceOracle(&_CreditManagerv3.TransactOpts, _priceOracle)
}

// SetQuotedMask is a paid mutator transaction binding the contract method 0x86dfa536.
//
// Solidity: function setQuotedMask(uint256 _quotedTokensMask) returns()
func (_CreditManagerv3 *CreditManagerv3Transactor) SetQuotedMask(opts *bind.TransactOpts, _quotedTokensMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "setQuotedMask", _quotedTokensMask)
}

// SetQuotedMask is a paid mutator transaction binding the contract method 0x86dfa536.
//
// Solidity: function setQuotedMask(uint256 _quotedTokensMask) returns()
func (_CreditManagerv3 *CreditManagerv3Session) SetQuotedMask(_quotedTokensMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetQuotedMask(&_CreditManagerv3.TransactOpts, _quotedTokensMask)
}

// SetQuotedMask is a paid mutator transaction binding the contract method 0x86dfa536.
//
// Solidity: function setQuotedMask(uint256 _quotedTokensMask) returns()
func (_CreditManagerv3 *CreditManagerv3TransactorSession) SetQuotedMask(_quotedTokensMask *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.SetQuotedMask(&_CreditManagerv3.TransactOpts, _quotedTokensMask)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 quotaChange, uint96 minQuota, uint96 maxQuota) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Transactor) UpdateQuota(opts *bind.TransactOpts, creditAccount common.Address, token common.Address, quotaChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "updateQuota", creditAccount, token, quotaChange, minQuota, maxQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 quotaChange, uint96 minQuota, uint96 maxQuota) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Session) UpdateQuota(creditAccount common.Address, token common.Address, quotaChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.UpdateQuota(&_CreditManagerv3.TransactOpts, creditAccount, token, quotaChange, minQuota, maxQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x604ca15f.
//
// Solidity: function updateQuota(address creditAccount, address token, int96 quotaChange, uint96 minQuota, uint96 maxQuota) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) UpdateQuota(creditAccount common.Address, token common.Address, quotaChange *big.Int, minQuota *big.Int, maxQuota *big.Int) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.UpdateQuota(&_CreditManagerv3.TransactOpts, creditAccount, token, quotaChange, minQuota, maxQuota)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x5a6f45de.
//
// Solidity: function withdrawCollateral(address creditAccount, address token, uint256 amount, address to) returns(uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Transactor) WithdrawCollateral(opts *bind.TransactOpts, creditAccount common.Address, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.contract.Transact(opts, "withdrawCollateral", creditAccount, token, amount, to)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x5a6f45de.
//
// Solidity: function withdrawCollateral(address creditAccount, address token, uint256 amount, address to) returns(uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3Session) WithdrawCollateral(creditAccount common.Address, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.WithdrawCollateral(&_CreditManagerv3.TransactOpts, creditAccount, token, amount, to)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x5a6f45de.
//
// Solidity: function withdrawCollateral(address creditAccount, address token, uint256 amount, address to) returns(uint256 tokensToDisable)
func (_CreditManagerv3 *CreditManagerv3TransactorSession) WithdrawCollateral(creditAccount common.Address, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditManagerv3.Contract.WithdrawCollateral(&_CreditManagerv3.TransactOpts, creditAccount, token, amount, to)
}

// CreditManagerv3SetCreditConfiguratorIterator is returned from FilterSetCreditConfigurator and is used to iterate over the raw logs and unpacked data for SetCreditConfigurator events raised by the CreditManagerv3 contract.
type CreditManagerv3SetCreditConfiguratorIterator struct {
	Event *CreditManagerv3SetCreditConfigurator // Event containing the contract specifics and raw log

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
func (it *CreditManagerv3SetCreditConfiguratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditManagerv3SetCreditConfigurator)
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
		it.Event = new(CreditManagerv3SetCreditConfigurator)
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
func (it *CreditManagerv3SetCreditConfiguratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditManagerv3SetCreditConfiguratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditManagerv3SetCreditConfigurator represents a SetCreditConfigurator event raised by the CreditManagerv3 contract.
type CreditManagerv3SetCreditConfigurator struct {
	NewConfigurator common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetCreditConfigurator is a free log retrieval operation binding the contract event 0xd87efcee33ed285df83ed2ffd66f67c15e0ecf17eb1f1705adae3ae2f1778da0.
//
// Solidity: event SetCreditConfigurator(address indexed newConfigurator)
func (_CreditManagerv3 *CreditManagerv3Filterer) FilterSetCreditConfigurator(opts *bind.FilterOpts, newConfigurator []common.Address) (*CreditManagerv3SetCreditConfiguratorIterator, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv3.contract.FilterLogs(opts, "SetCreditConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return &CreditManagerv3SetCreditConfiguratorIterator{contract: _CreditManagerv3.contract, event: "SetCreditConfigurator", logs: logs, sub: sub}, nil
}

// WatchSetCreditConfigurator is a free log subscription operation binding the contract event 0xd87efcee33ed285df83ed2ffd66f67c15e0ecf17eb1f1705adae3ae2f1778da0.
//
// Solidity: event SetCreditConfigurator(address indexed newConfigurator)
func (_CreditManagerv3 *CreditManagerv3Filterer) WatchSetCreditConfigurator(opts *bind.WatchOpts, sink chan<- *CreditManagerv3SetCreditConfigurator, newConfigurator []common.Address) (event.Subscription, error) {

	var newConfiguratorRule []interface{}
	for _, newConfiguratorItem := range newConfigurator {
		newConfiguratorRule = append(newConfiguratorRule, newConfiguratorItem)
	}

	logs, sub, err := _CreditManagerv3.contract.WatchLogs(opts, "SetCreditConfigurator", newConfiguratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditManagerv3SetCreditConfigurator)
				if err := _CreditManagerv3.contract.UnpackLog(event, "SetCreditConfigurator", log); err != nil {
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

// ParseSetCreditConfigurator is a log parse operation binding the contract event 0xd87efcee33ed285df83ed2ffd66f67c15e0ecf17eb1f1705adae3ae2f1778da0.
//
// Solidity: event SetCreditConfigurator(address indexed newConfigurator)
func (_CreditManagerv3 *CreditManagerv3Filterer) ParseSetCreditConfigurator(log types.Log) (*CreditManagerv3SetCreditConfigurator, error) {
	event := new(CreditManagerv3SetCreditConfigurator)
	if err := _CreditManagerv3.contract.UnpackLog(event, "SetCreditConfigurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
