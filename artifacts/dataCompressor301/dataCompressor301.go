// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataCompressor301

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
	CmName                    string
	CreditFacade              common.Address
	Underlying                common.Address
	Debt                      *big.Int
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
	Underlying   common.Address
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

// DataCompressor301MetaData contains all meta data concerning the DataCompressor301 contract.
var DataCompressor301MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_addressProvider\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractsRegister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccountData\",\"inputs\":[{\"name\":\"_creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceOnDemand[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountData\",\"components\":[{\"name\":\"isSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cmName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balances\",\"type\":\"tuple[]\",\"internalType\":\"structTokenBalance[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isForbidden\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isQuoted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"since\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeBots\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCreditAccountsByBorrower\",\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceOnDemand[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"isSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cmName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balances\",\"type\":\"tuple[]\",\"internalType\":\"structTokenBalance[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isForbidden\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isQuoted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"since\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeBots\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCreditAccountsByCreditManager\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceOnDemand[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"isSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cmName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balances\",\"type\":\"tuple[]\",\"internalType\":\"structTokenBalance[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isForbidden\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isQuoted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"since\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeBots\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCreditManagerData\",\"inputs\":[{\"name\":\"_cm\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerData\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableToBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adapters\",\"type\":\"tuple[]\",\"internalType\":\"structContractAdapter[]\",\"components\":[{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"isDegenMode\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"degenNFT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"lirm\",\"type\":\"tuple\",\"internalType\":\"structLinearModel\",\"components\":[{\"name\":\"interestModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"U_1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"U_2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_base\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope3\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditManagersV3List\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerData[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableToBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adapters\",\"type\":\"tuple[]\",\"internalType\":\"structContractAdapter[]\",\"components\":[{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"adapter\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"isDegenMode\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"degenNFT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"lirm\",\"type\":\"tuple\",\"internalType\":\"structLinearModel\",\"components\":[{\"name\":\"interestModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"U_1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"U_2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_base\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope3\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGaugesV3Data\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structGaugeInfo[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentEpoch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"epochFrozen\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quotaParams\",\"type\":\"tuple[]\",\"internalType\":\"structGaugeQuotaParams[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalVotesLpSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"totalVotesCaSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stakerVotesLpSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"stakerVotesCaSide\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLiquidatableCreditAccounts\",\"inputs\":[{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceOnDemand[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"isSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeedsNeeded\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cmName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeIndexLastUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cumulativeQuotaInterest\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatedBorrowRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balances\",\"type\":\"tuple[]\",\"internalType\":\"structTokenBalance[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isForbidden\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isQuoted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaCumulativeIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"since\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cfVersion\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"activeBots\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPoolData\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structPoolData\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dieselToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableToBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate_RAY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolQuotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gauge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"zappers\",\"type\":\"tuple[]\",\"internalType\":\"structZapperInfo[]\",\"components\":[{\"name\":\"zapper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"lirm\",\"type\":\"tuple\",\"internalType\":\"structLinearModel\",\"components\":[{\"name\":\"interestModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"U_1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"U_2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_base\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope3\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolsV3List\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structPoolData[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dieselToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"availableToBorrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate_RAY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"poolQuotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gauge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"zappers\",\"type\":\"tuple[]\",\"internalType\":\"structZapperInfo[]\",\"components\":[{\"name\":\"zapper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"lirm\",\"type\":\"tuple\",\"internalType\":\"structLinearModel\",\"components\":[{\"name\":\"interestModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"U_1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"U_2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_base\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope1\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope2\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"R_slope3\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isBorrowingMoreU2Forbidden\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zapperRegister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIZapperRegister\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"CreditManagerIsNotV3Exception\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceFeedDoesNotExistException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RegisteredCreditManagerOnlyException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RegisteredPoolOnlyException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// DataCompressor301ABI is the input ABI used to generate the binding from.
// Deprecated: Use DataCompressor301MetaData.ABI instead.
var DataCompressor301ABI = DataCompressor301MetaData.ABI

// DataCompressor301 is an auto generated Go binding around an Ethereum contract.
type DataCompressor301 struct {
	DataCompressor301Caller     // Read-only binding to the contract
	DataCompressor301Transactor // Write-only binding to the contract
	DataCompressor301Filterer   // Log filterer for contract events
}

// DataCompressor301Caller is an auto generated read-only Go binding around an Ethereum contract.
type DataCompressor301Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressor301Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DataCompressor301Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressor301Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataCompressor301Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressor301Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataCompressor301Session struct {
	Contract     *DataCompressor301 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DataCompressor301CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCompressor301CallerSession struct {
	Contract *DataCompressor301Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DataCompressor301TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataCompressor301TransactorSession struct {
	Contract     *DataCompressor301Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DataCompressor301Raw is an auto generated low-level Go binding around an Ethereum contract.
type DataCompressor301Raw struct {
	Contract *DataCompressor301 // Generic contract binding to access the raw methods on
}

// DataCompressor301CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCompressor301CallerRaw struct {
	Contract *DataCompressor301Caller // Generic read-only contract binding to access the raw methods on
}

// DataCompressor301TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataCompressor301TransactorRaw struct {
	Contract *DataCompressor301Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDataCompressor301 creates a new instance of DataCompressor301, bound to a specific deployed contract.
func NewDataCompressor301(address common.Address, backend bind.ContractBackend) (*DataCompressor301, error) {
	contract, err := bindDataCompressor301(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataCompressor301{DataCompressor301Caller: DataCompressor301Caller{contract: contract}, DataCompressor301Transactor: DataCompressor301Transactor{contract: contract}, DataCompressor301Filterer: DataCompressor301Filterer{contract: contract}}, nil
}

// NewDataCompressor301Caller creates a new read-only instance of DataCompressor301, bound to a specific deployed contract.
func NewDataCompressor301Caller(address common.Address, caller bind.ContractCaller) (*DataCompressor301Caller, error) {
	contract, err := bindDataCompressor301(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressor301Caller{contract: contract}, nil
}

// NewDataCompressor301Transactor creates a new write-only instance of DataCompressor301, bound to a specific deployed contract.
func NewDataCompressor301Transactor(address common.Address, transactor bind.ContractTransactor) (*DataCompressor301Transactor, error) {
	contract, err := bindDataCompressor301(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressor301Transactor{contract: contract}, nil
}

// NewDataCompressor301Filterer creates a new log filterer instance of DataCompressor301, bound to a specific deployed contract.
func NewDataCompressor301Filterer(address common.Address, filterer bind.ContractFilterer) (*DataCompressor301Filterer, error) {
	contract, err := bindDataCompressor301(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataCompressor301Filterer{contract: contract}, nil
}

// bindDataCompressor301 binds a generic wrapper to an already deployed contract.
func bindDataCompressor301(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataCompressor301MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressor301 *DataCompressor301Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressor301.Contract.DataCompressor301Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressor301 *DataCompressor301Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressor301.Contract.DataCompressor301Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressor301 *DataCompressor301Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressor301.Contract.DataCompressor301Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressor301 *DataCompressor301CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressor301.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressor301 *DataCompressor301TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressor301.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressor301 *DataCompressor301TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressor301.Contract.contract.Transact(opts, method, params...)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301Caller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301Session) ContractsRegister() (common.Address, error) {
	return _DataCompressor301.Contract.ContractsRegister(&_DataCompressor301.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301CallerSession) ContractsRegister() (common.Address, error) {
	return _DataCompressor301.Contract.ContractsRegister(&_DataCompressor301.CallOpts)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301Caller) GetCreditManagerData(opts *bind.CallOpts, _cm common.Address) (CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "getCreditManagerData", _cm)

	if err != nil {
		return *new(CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditManagerData)).(*CreditManagerData)

	return out0, err

}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301Session) GetCreditManagerData(_cm common.Address) (CreditManagerData, error) {
	return _DataCompressor301.Contract.GetCreditManagerData(&_DataCompressor301.CallOpts, _cm)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _cm) view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301CallerSession) GetCreditManagerData(_cm common.Address) (CreditManagerData, error) {
	return _DataCompressor301.Contract.GetCreditManagerData(&_DataCompressor301.CallOpts, _cm)
}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301Caller) GetCreditManagersV3List(opts *bind.CallOpts) ([]CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "getCreditManagersV3List")

	if err != nil {
		return *new([]CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditManagerData)).(*[]CreditManagerData)

	return out0, err

}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301Session) GetCreditManagersV3List() ([]CreditManagerData, error) {
	return _DataCompressor301.Contract.GetCreditManagersV3List(&_DataCompressor301.CallOpts)
}

// GetCreditManagersV3List is a free data retrieval call binding the contract method 0xc7fd2b45.
//
// Solidity: function getCreditManagersV3List() view returns((address,string,uint256,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],bool,address,uint256,uint8,uint16,uint16,uint16,uint16,uint16,(address,uint16,uint16,uint96,uint96,bool)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301CallerSession) GetCreditManagersV3List() ([]CreditManagerData, error) {
	return _DataCompressor301.Contract.GetCreditManagersV3List(&_DataCompressor301.CallOpts)
}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,address,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressor301 *DataCompressor301Caller) GetGaugesV3Data(opts *bind.CallOpts, staker common.Address) ([]GaugeInfo, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "getGaugesV3Data", staker)

	if err != nil {
		return *new([]GaugeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]GaugeInfo)).(*[]GaugeInfo)

	return out0, err

}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,address,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressor301 *DataCompressor301Session) GetGaugesV3Data(staker common.Address) ([]GaugeInfo, error) {
	return _DataCompressor301.Contract.GetGaugesV3Data(&_DataCompressor301.CallOpts, staker)
}

// GetGaugesV3Data is a free data retrieval call binding the contract method 0x0ddedaaa.
//
// Solidity: function getGaugesV3Data(address staker) view returns((address,address,string,string,address,uint16,bool,(address,uint16,uint16,uint96,uint96,uint16,uint16,uint96,uint96,bool,uint96,uint96)[])[] result)
func (_DataCompressor301 *DataCompressor301CallerSession) GetGaugesV3Data(staker common.Address) ([]GaugeInfo, error) {
	return _DataCompressor301.Contract.GetGaugesV3Data(&_DataCompressor301.CallOpts, staker)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301Caller) GetPoolData(opts *bind.CallOpts, _pool common.Address) (PoolData, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "getPoolData", _pool)

	if err != nil {
		return *new(PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolData)).(*PoolData)

	return out0, err

}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301Session) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressor301.Contract.GetPoolData(&_DataCompressor301.CallOpts, _pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool) result)
func (_DataCompressor301 *DataCompressor301CallerSession) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressor301.Contract.GetPoolData(&_DataCompressor301.CallOpts, _pool)
}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301Caller) GetPoolsV3List(opts *bind.CallOpts) ([]PoolData, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "getPoolsV3List")

	if err != nil {
		return *new([]PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolData)).(*[]PoolData)

	return out0, err

}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301Session) GetPoolsV3List() ([]PoolData, error) {
	return _DataCompressor301.Contract.GetPoolsV3List(&_DataCompressor301.CallOpts)
}

// GetPoolsV3List is a free data retrieval call binding the contract method 0xa0f068df.
//
// Solidity: function getPoolsV3List() view returns((address,address,address,string,string,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,(address,uint16,uint16,uint96,uint96,bool)[],(address,address,address)[],(address,uint256,uint16,uint16,uint16,uint16,uint16,uint16,bool),bool)[] result)
func (_DataCompressor301 *DataCompressor301CallerSession) GetPoolsV3List() ([]PoolData, error) {
	return _DataCompressor301.Contract.GetPoolsV3List(&_DataCompressor301.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressor301 *DataCompressor301Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressor301 *DataCompressor301Session) Version() (*big.Int, error) {
	return _DataCompressor301.Contract.Version(&_DataCompressor301.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressor301 *DataCompressor301CallerSession) Version() (*big.Int, error) {
	return _DataCompressor301.Contract.Version(&_DataCompressor301.CallOpts)
}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301Caller) ZapperRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressor301.contract.Call(opts, &out, "zapperRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301Session) ZapperRegister() (common.Address, error) {
	return _DataCompressor301.Contract.ZapperRegister(&_DataCompressor301.CallOpts)
}

// ZapperRegister is a free data retrieval call binding the contract method 0x90e6ffff.
//
// Solidity: function zapperRegister() view returns(address)
func (_DataCompressor301 *DataCompressor301CallerSession) ZapperRegister() (common.Address, error) {
	return _DataCompressor301.Contract.ZapperRegister(&_DataCompressor301.CallOpts)
}

// GetCreditAccountData is a paid mutator transaction binding the contract method 0xe50b35ec.
//
// Solidity: function getCreditAccountData(address _creditAccount, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[]) result)
func (_DataCompressor301 *DataCompressor301Transactor) GetCreditAccountData(opts *bind.TransactOpts, _creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.contract.Transact(opts, "getCreditAccountData", _creditAccount, priceUpdates)
}

// GetCreditAccountData is a paid mutator transaction binding the contract method 0xe50b35ec.
//
// Solidity: function getCreditAccountData(address _creditAccount, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[]) result)
func (_DataCompressor301 *DataCompressor301Session) GetCreditAccountData(_creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountData(&_DataCompressor301.TransactOpts, _creditAccount, priceUpdates)
}

// GetCreditAccountData is a paid mutator transaction binding the contract method 0xe50b35ec.
//
// Solidity: function getCreditAccountData(address _creditAccount, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[]) result)
func (_DataCompressor301 *DataCompressor301TransactorSession) GetCreditAccountData(_creditAccount common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountData(&_DataCompressor301.TransactOpts, _creditAccount, priceUpdates)
}

// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Transactor) GetCreditAccountsByBorrower(opts *bind.TransactOpts, borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.contract.Transact(opts, "getCreditAccountsByBorrower", borrower, priceUpdates)
}

// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Session) GetCreditAccountsByBorrower(borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountsByBorrower(&_DataCompressor301.TransactOpts, borrower, priceUpdates)
}

// GetCreditAccountsByBorrower is a paid mutator transaction binding the contract method 0xcccecfd3.
//
// Solidity: function getCreditAccountsByBorrower(address borrower, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301TransactorSession) GetCreditAccountsByBorrower(borrower common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountsByBorrower(&_DataCompressor301.TransactOpts, borrower, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Transactor) GetCreditAccountsByCreditManager(opts *bind.TransactOpts, creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.contract.Transact(opts, "getCreditAccountsByCreditManager", creditManager, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Session) GetCreditAccountsByCreditManager(creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountsByCreditManager(&_DataCompressor301.TransactOpts, creditManager, priceUpdates)
}

// GetCreditAccountsByCreditManager is a paid mutator transaction binding the contract method 0xc432fd20.
//
// Solidity: function getCreditAccountsByCreditManager(address creditManager, (address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301TransactorSession) GetCreditAccountsByCreditManager(creditManager common.Address, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetCreditAccountsByCreditManager(&_DataCompressor301.TransactOpts, creditManager, priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Transactor) GetLiquidatableCreditAccounts(opts *bind.TransactOpts, priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.contract.Transact(opts, "getLiquidatableCreditAccounts", priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301Session) GetLiquidatableCreditAccounts(priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetLiquidatableCreditAccounts(&_DataCompressor301.TransactOpts, priceUpdates)
}

// GetLiquidatableCreditAccounts is a paid mutator transaction binding the contract method 0x33eb9df6.
//
// Solidity: function getLiquidatableCreditAccounts((address,bytes)[] priceUpdates) returns((bool,address[],address,address,address,string,address,address,uint256,uint256,uint128,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool,bool,uint256,uint16,uint256)[],uint64,uint256,uint40,address[])[] result)
func (_DataCompressor301 *DataCompressor301TransactorSession) GetLiquidatableCreditAccounts(priceUpdates []PriceOnDemand) (*types.Transaction, error) {
	return _DataCompressor301.Contract.GetLiquidatableCreditAccounts(&_DataCompressor301.TransactOpts, priceUpdates)
}
