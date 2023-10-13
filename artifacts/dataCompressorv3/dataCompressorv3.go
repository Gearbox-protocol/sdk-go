// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataCompressorv3

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

// ContractAdapter is an auto generated low-level Go binding around an user-defined struct.
type ContractAdapter struct {
	TargetContract common.Address
	Adapter        common.Address
}

// CreditAccountData is an auto generated low-level Go binding around an user-defined struct.
type CreditAccountData struct {
	IsSuccessful              bool
	PriceFeedsNeeded          []common.Address
	Addr                      common.Address
	Borrower                  common.Address
	CreditManager             common.Address
	CmDescription             string
	CreditFacade              common.Address
	Underlying                common.Address
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
	HealthFactor              *big.Int
	BaseBorrowRate            *big.Int
	AggregatedBorrowRate      *big.Int
	Balances                  []TokenBalance
	Since                     uint64
	CfVersion                 *big.Int
	ExpirationDate            *big.Int
	ActiveBots                []common.Address
	MaxApprovedBots           *big.Int
	SchedultedWithdrawals     [2]ScheduledWithdrawal
}

// CreditManagerData is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerData struct {
	Addr                       common.Address
	Name                       string
	CfVersion                  *big.Int
	CreditFacade               common.Address
	CreditConfigurator         common.Address
	Underlying                 common.Address
	Pool                       common.Address
	TotalDebt                  *big.Int
	TotalDebtLimit             *big.Int
	BaseBorrowRate             *big.Int
	MinDebt                    *big.Int
	MaxDebt                    *big.Int
	AvailableToBorrow          *big.Int
	CollateralTokens           []common.Address
	Adapters                   []ContractAdapter
	LiquidationThresholds      []*big.Int
	IsDegenMode                bool
	DegenNFT                   common.Address
	ForbiddenTokenMask         *big.Int
	MaxEnabledTokensLength     uint8
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
	Quotas                     []QuotaInfo
	Lirm                       LinearModel
	IsPaused                   bool
}

// CreditManagerDebtParams is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerDebtParams struct {
	CreditManager     common.Address
	Borrowed          *big.Int
	Limit             *big.Int
	AvailableToBorrow *big.Int
}

// GaugeInfo is an auto generated low-level Go binding around an user-defined struct.
type GaugeInfo struct {
	Addr         common.Address
	Pool         common.Address
	Symbol       string
	Name         string
	CurrentEpoch uint16
	EpochFrozen  bool
	QuotaParams  []GaugeQuotaParams
}

// GaugeQuotaParams is an auto generated low-level Go binding around an user-defined struct.
type GaugeQuotaParams struct {
	Token             common.Address
	MinRate           uint16
	MaxRate           uint16
	TotalVotesLpSide  *big.Int
	TotalVotesCaSide  *big.Int
	Rate              uint16
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
	StakerVotesLpSide *big.Int
	StakerVotesCaSide *big.Int
}

// LinearModel is an auto generated low-level Go binding around an user-defined struct.
type LinearModel struct {
	InterestModel              common.Address
	Version                    *big.Int
	U1                         uint16
	U2                         uint16
	RBase                      uint16
	RSlope1                    uint16
	RSlope2                    uint16
	RSlope3                    uint16
	IsBorrowingMoreU2Forbidden bool
}

// PoolData is an auto generated low-level Go binding around an user-defined struct.
type PoolData struct {
	Addr                    common.Address
	Underlying              common.Address
	DieselToken             common.Address
	Symbol                  string
	Name                    string
	BaseInterestIndex       *big.Int
	AvailableLiquidity      *big.Int
	ExpectedLiquidity       *big.Int
	TotalBorrowed           *big.Int
	TotalDebtLimit          *big.Int
	CreditManagerDebtParams []CreditManagerDebtParams
	TotalAssets             *big.Int
	TotalSupply             *big.Int
	SupplyRate              *big.Int
	BaseInterestRate        *big.Int
	DieselRateRAY           *big.Int
	WithdrawFee             *big.Int
	LastBaseInterestUpdate  *big.Int
	BaseInterestIndexLU     *big.Int
	Version                 *big.Int
	PoolQuotaKeeper         common.Address
	Gauge                   common.Address
	Quotas                  []QuotaInfo
	Zappers                 []ZapperInfo
	Lirm                    LinearModel
	IsPaused                bool
}

// PriceOnDemand is an auto generated low-level Go binding around an user-defined struct.
type PriceOnDemand struct {
	Token    common.Address
	CallData []byte
}

// QuotaInfo is an auto generated low-level Go binding around an user-defined struct.
type QuotaInfo struct {
	Token            common.Address
	Rate             uint16
	QuotaIncreaseFee uint16
	TotalQuoted      *big.Int
	Limit            *big.Int
	IsActive         bool
}

// ScheduledWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type ScheduledWithdrawal struct {
	TokenIndex uint8
	Maturity   *big.Int
	Token      common.Address
	Amount     *big.Int
}

// TokenBalance is an auto generated low-level Go binding around an user-defined struct.
type TokenBalance struct {
	Token                  common.Address
	Balance                *big.Int
	IsForbidden            bool
	IsEnabled              bool
	IsQuoted               bool
	Quota                  *big.Int
	QuotaRate              uint16
	QuotaCumulativeIndexLU *big.Int
}

// ZapperInfo is an auto generated low-level Go binding around an user-defined struct.
type ZapperInfo struct {
	Zapper   common.Address
	TokenIn  common.Address
	TokenOut common.Address
}

// DataCompressorv3MetaData contains all meta data concerning the DataCompressorv3 contract.
var DataCompressorv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CreditManagerIsNotV3Exception\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedDoesNotExistException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredCreditManagerOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegisteredPoolOnlyException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOnDemand[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"getCreditAccountData\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSuccessful\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cmDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accruedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValueUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twvUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForbidden\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isQuoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"quotaRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"since\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"address[]\",\"name\":\"activeBots\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxApprovedBots\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"maturity\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structScheduledWithdrawal[2]\",\"name\":\"schedultedWithdrawals\",\"type\":\"tuple[2]\"}],\"internalType\":\"structCreditAccountData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOnDemand[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"getCreditAccountsByBorrower\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSuccessful\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cmDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accruedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValueUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twvUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForbidden\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isQuoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"quotaRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"since\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"address[]\",\"name\":\"activeBots\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxApprovedBots\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"maturity\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structScheduledWithdrawal[2]\",\"name\":\"schedultedWithdrawals\",\"type\":\"tuple[2]\"}],\"internalType\":\"structCreditAccountData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOnDemand[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"getCreditAccountsByCreditManager\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSuccessful\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cmDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accruedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValueUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twvUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForbidden\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isQuoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"quotaRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"since\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"address[]\",\"name\":\"activeBots\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxApprovedBots\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"maturity\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structScheduledWithdrawal[2]\",\"name\":\"schedultedWithdrawals\",\"type\":\"tuple[2]\"}],\"internalType\":\"structCreditAccountData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cm\",\"type\":\"address\"}],\"name\":\"getCreditManagerData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableToBorrow\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"collateralTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isDegenMode\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structQuotaInfo[]\",\"name\":\"quotas\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"U_1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"U_2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_base\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope3\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\"}],\"internalType\":\"structLinearModel\",\"name\":\"lirm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditManagerData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditManagersV3List\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableToBorrow\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"collateralTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isDegenMode\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structQuotaInfo[]\",\"name\":\"quotas\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"U_1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"U_2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_base\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope3\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\"}],\"internalType\":\"structLinearModel\",\"name\":\"lirm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditManagerData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getGaugesV3Data\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"currentEpoch\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"epochFrozen\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"minRate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxRate\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalVotesLpSide\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"totalVotesCaSide\",\"type\":\"uint96\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint96\",\"name\":\"stakerVotesLpSide\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"stakerVotesCaSide\",\"type\":\"uint96\"}],\"internalType\":\"structGaugeQuotaParams[]\",\"name\":\"quotaParams\",\"type\":\"tuple[]\"}],\"internalType\":\"structGaugeInfo[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structPriceOnDemand[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"getLiquidatableCreditAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isSuccessful\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cmDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexNow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"accruedInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValueUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twvUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isForbidden\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isQuoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"quotaRate\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"since\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"cfVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"address[]\",\"name\":\"activeBots\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxApprovedBots\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"maturity\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structScheduledWithdrawal[2]\",\"name\":\"schedultedWithdrawals\",\"type\":\"tuple[2]\"}],\"internalType\":\"structCreditAccountData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableToBorrow\",\"type\":\"uint256\"}],\"internalType\":\"structCreditManagerDebtParams[]\",\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolQuotaKeeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structQuotaInfo[]\",\"name\":\"quotas\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"zapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"internalType\":\"structZapperInfo[]\",\"name\":\"zappers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"U_1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"U_2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_base\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope3\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\"}],\"internalType\":\"structLinearModel\",\"name\":\"lirm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structPoolData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolsV3List\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableToBorrow\",\"type\":\"uint256\"}],\"internalType\":\"structCreditManagerDebtParams[]\",\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolQuotaKeeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structQuotaInfo[]\",\"name\":\"quotas\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"zapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"internalType\":\"structZapperInfo[]\",\"name\":\"zappers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestModel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"U_1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"U_2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_base\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"R_slope3\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\"}],\"internalType\":\"structLinearModel\",\"name\":\"lirm\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structPoolData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapperRegister\",\"outputs\":[{\"internalType\":\"contractIZapperRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataCompressorv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use DataCompressorv3MetaData.ABI instead.
var DataCompressorv3ABI = DataCompressorv3MetaData.ABI

// DataCompressorv3 is an auto generated Go binding around an Ethereum contract.
type DataCompressorv3 struct {
	DataCompressorv3Caller     // Read-only binding to the contract
	DataCompressorv3Transactor // Write-only binding to the contract
	DataCompressorv3Filterer   // Log filterer for contract events
}

// DataCompressorv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type DataCompressorv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DataCompressorv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataCompressorv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataCompressorv3Session struct {
	Contract     *DataCompressorv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCompressorv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCompressorv3CallerSession struct {
	Contract *DataCompressorv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DataCompressorv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataCompressorv3TransactorSession struct {
	Contract     *DataCompressorv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DataCompressorv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type DataCompressorv3Raw struct {
	Contract *DataCompressorv3 // Generic contract binding to access the raw methods on
}

// DataCompressorv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCompressorv3CallerRaw struct {
	Contract *DataCompressorv3Caller // Generic read-only contract binding to access the raw methods on
}

// DataCompressorv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataCompressorv3TransactorRaw struct {
	Contract *DataCompressorv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDataCompressorv3 creates a new instance of DataCompressorv3, bound to a specific deployed contract.
func NewDataCompressorv3(address common.Address, backend bind.ContractBackend) (*DataCompressorv3, error) {
	contract, err := bindDataCompressorv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv3{DataCompressorv3Caller: DataCompressorv3Caller{contract: contract}, DataCompressorv3Transactor: DataCompressorv3Transactor{contract: contract}, DataCompressorv3Filterer: DataCompressorv3Filterer{contract: contract}}, nil
}

// NewDataCompressorv3Caller creates a new read-only instance of DataCompressorv3, bound to a specific deployed contract.
func NewDataCompressorv3Caller(address common.Address, caller bind.ContractCaller) (*DataCompressorv3Caller, error) {
	contract, err := bindDataCompressorv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv3Caller{contract: contract}, nil
}

// NewDataCompressorv3Transactor creates a new write-only instance of DataCompressorv3, bound to a specific deployed contract.
func NewDataCompressorv3Transactor(address common.Address, transactor bind.ContractTransactor) (*DataCompressorv3Transactor, error) {
	contract, err := bindDataCompressorv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv3Transactor{contract: contract}, nil
}

// NewDataCompressorv3Filterer creates a new log filterer instance of DataCompressorv3, bound to a specific deployed contract.
func NewDataCompressorv3Filterer(address common.Address, filterer bind.ContractFilterer) (*DataCompressorv3Filterer, error) {
	contract, err := bindDataCompressorv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv3Filterer{contract: contract}, nil
}

// bindDataCompressorv3 binds a generic wrapper to an already deployed contract.
func bindDataCompressorv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataCompressorv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressorv3 *DataCompressorv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressorv3.Contract.DataCompressorv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressorv3 *DataCompressorv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.DataCompressorv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressorv3 *DataCompressorv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.DataCompressorv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressorv3 *DataCompressorv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressorv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressorv3 *DataCompressorv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressorv3 *DataCompressorv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.contract.Transact(opts, method, params...)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3Caller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3Session) ContractsRegister() (common.Address, error) {
	return _DataCompressorv3.Contract.ContractsRegister(&_DataCompressorv3.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3CallerSession) ContractsRegister() (common.Address, error) {
	return _DataCompressorv3.Contract.ContractsRegister(&_DataCompressorv3.CallOpts)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3Caller) GetCreditManagerData(opts *bind.CallOpts, _cm common.Address) (CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getCreditManagerData", _cm)

	if err != nil {
		return *new(CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditManagerData)).(*CreditManagerData)

	return out0, err

}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3Session) GetCreditManagerData(_cm common.Address) (CreditManagerData, error) {
	return _DataCompressorv3.Contract.GetCreditManagerData(&_DataCompressorv3.CallOpts, _cm)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3CallerSession) GetCreditManagerData(_cm common.Address) (CreditManagerData, error) {
	return _DataCompressorv3.Contract.GetCreditManagerData(&_DataCompressorv3.CallOpts, _cm)
}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3Caller) GetCreditManagersV3List(opts *bind.CallOpts) ([]CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getCreditManagersV3List")

	if err != nil {
		return *new([]CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditManagerData)).(*[]CreditManagerData)

	return out0, err

}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetCreditManagersV3List() ([]CreditManagerData, error) {
	return _DataCompressorv3.Contract.GetCreditManagersV3List(&_DataCompressorv3.CallOpts)
}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3CallerSession) GetCreditManagersV3List() ([]CreditManagerData, error) {
	return _DataCompressorv3.Contract.GetCreditManagersV3List(&_DataCompressorv3.CallOpts)
}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressorv3 *DataCompressorv3Caller) GetGaugesV3Data(opts *bind.CallOpts, staker common.Address) ([]GaugeInfo, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getGaugesV3Data", staker)

	if err != nil {
		return *new([]GaugeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]GaugeInfo)).(*[]GaugeInfo)

	return out0, err

}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetGaugesV3Data(staker common.Address) ([]GaugeInfo, error) {
	return _DataCompressorv3.Contract.GetGaugesV3Data(&_DataCompressorv3.CallOpts, staker)
}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressorv3 *DataCompressorv3CallerSession) GetGaugesV3Data(staker common.Address) ([]GaugeInfo, error) {
	return _DataCompressorv3.Contract.GetGaugesV3Data(&_DataCompressorv3.CallOpts, staker)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3Caller) GetPoolData(opts *bind.CallOpts, _pool common.Address) (PoolData, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getPoolData", _pool)

	if err != nil {
		return *new(PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolData)).(*PoolData)

	return out0, err

}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3Session) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressorv3.Contract.GetPoolData(&_DataCompressorv3.CallOpts, _pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressorv3 *DataCompressorv3CallerSession) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressorv3.Contract.GetPoolData(&_DataCompressorv3.CallOpts, _pool)
}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3Caller) GetPoolsV3List(opts *bind.CallOpts) ([]PoolData, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getPoolsV3List")

	if err != nil {
		return *new([]PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolData)).(*[]PoolData)

	return out0, err

}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetPoolsV3List() ([]PoolData, error) {
	return _DataCompressorv3.Contract.GetPoolsV3List(&_DataCompressorv3.CallOpts)
}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressorv3 *DataCompressorv3CallerSession) GetPoolsV3List() ([]PoolData, error) {
	return _DataCompressorv3.Contract.GetPoolsV3List(&_DataCompressorv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv3 *DataCompressorv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv3 *DataCompressorv3Session) Version() (*big.Int, error) {
	return _DataCompressorv3.Contract.Version(&_DataCompressorv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv3 *DataCompressorv3CallerSession) Version() (*big.Int, error) {
	return _DataCompressorv3.Contract.Version(&_DataCompressorv3.CallOpts)
}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3Caller) ZapperRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "zapperRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3Session) ZapperRegister() (common.Address, error) {
	return _DataCompressorv3.Contract.ZapperRegister(&_DataCompressorv3.CallOpts)
}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressorv3 *DataCompressorv3CallerSession) ZapperRegister() (common.Address, error) {
	return _DataCompressorv3.Contract.ZapperRegister(&_DataCompressorv3.CallOpts)
}



// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Transactor) GetCreditAccountsByBorrower(opts *bind.TransactOpts, borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.contract.Transact(opts, "getCreditAccountsByBorrower", borrower, priceUpdates)
}

// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetCreditAccountsByBorrower(borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetCreditAccountsByBorrower(&_DataCompressorv3.TransactOpts, borrower, priceUpdates)
}

// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3TransactorSession) GetCreditAccountsByBorrower(borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetCreditAccountsByBorrower(&_DataCompressorv3.TransactOpts, borrower, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Transactor) GetCreditAccountsByCreditManager(opts *bind.TransactOpts, creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.contract.Transact(opts, "getCreditAccountsByCreditManager", creditManager, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetCreditAccountsByCreditManager(creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetCreditAccountsByCreditManager(&_DataCompressorv3.TransactOpts, creditManager, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3TransactorSession) GetCreditAccountsByCreditManager(creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetCreditAccountsByCreditManager(&_DataCompressorv3.TransactOpts, creditManager, priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Transactor) GetLiquidatableCreditAccounts(opts *bind.TransactOpts, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.contract.Transact(opts, "getLiquidatableCreditAccounts", priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3Session) GetLiquidatableCreditAccounts(priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetLiquidatableCreditAccounts(&_DataCompressorv3.TransactOpts, priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[],uint256,(uint8,uint40,address,uint256)[2])[] result)
func (_DataCompressorv3 *DataCompressorv3TransactorSession) GetLiquidatableCreditAccounts(priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressorv3.Contract.GetLiquidatableCreditAccounts(&_DataCompressorv3.TransactOpts, priceUpdates)
}
