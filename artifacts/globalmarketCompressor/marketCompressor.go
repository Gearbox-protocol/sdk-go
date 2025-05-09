// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalmarketCompressor

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

// AdapterState is an auto generated low-level Go binding around an user-defined struct.
type AdapterState struct {
	BaseParams     BaseParams
	TargetContract common.Address
}

// BaseParams is an auto generated low-level Go binding around an user-defined struct.
type BaseParams struct {
	Addr             common.Address
	Version          *big.Int
	ContractType     [32]byte
	SerializedParams []byte
}

// BaseState is an auto generated low-level Go binding around an user-defined struct.
type BaseState struct {
	BaseParams BaseParams
}

// CollateralToken is an auto generated low-level Go binding around an user-defined struct.
type CollateralToken struct {
	Token                common.Address
	LiquidationThreshold uint16
}

// CreditFacadeState is an auto generated low-level Go binding around an user-defined struct.
type CreditFacadeState struct {
	BaseParams                BaseParams
	DegenNFT                  common.Address
	BotList                   common.Address
	Expirable                 bool
	ExpirationDate            *big.Int
	MaxDebtPerBlockMultiplier uint8
	MinDebt                   *big.Int
	MaxDebt                   *big.Int
	ForbiddenTokensMask       *big.Int
	IsPaused                  bool
}

// CreditManagerDebtParams is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerDebtParams struct {
	CreditManager common.Address
	Borrowed      *big.Int
	Limit         *big.Int
	Available     *big.Int
}

// CreditManagerState is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerState struct {
	BaseParams                 BaseParams
	Name                       string
	AccountFactory             common.Address
	Underlying                 common.Address
	Pool                       common.Address
	CreditFacade               common.Address
	CreditConfigurator         common.Address
	MaxEnabledTokens           uint8
	CollateralTokens           []CollateralToken
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}

// CreditSuiteData is an auto generated low-level Go binding around an user-defined struct.
type CreditSuiteData struct {
	CreditFacade       CreditFacadeState
	CreditManager      CreditManagerState
	CreditConfigurator BaseState
	AccountFactory     BaseState
	Adapters           []AdapterState
}

// MarketData is an auto generated low-level Go binding around an user-defined struct.
type MarketData struct {
	Acl                  common.Address
	ContractsRegister    common.Address
	Treasury             common.Address
	Pool                 PoolState
	QuotaKeeper          QuotaKeeperState
	InterestRateModel    BaseState
	RateKeeper           RateKeeperState
	PriceOracle          PriceOracleState
	LossPolicy           BaseState
	Tokens               []TokenData
	CreditManagers       []CreditSuiteData
	Configurator         common.Address
	PausableAdmins       []common.Address
	UnpausableAdmins     []common.Address
	EmergencyLiquidators []common.Address
}

// MarketFilter is an auto generated low-level Go binding around an user-defined struct.
type MarketFilter struct {
	Configurators []common.Address
	Pools         []common.Address
	Underlying    common.Address
}

// PoolState is an auto generated low-level Go binding around an user-defined struct.
type PoolState struct {
	BaseParams              BaseParams
	Symbol                  string
	Name                    string
	Decimals                uint8
	TotalSupply             *big.Int
	QuotaKeeper             common.Address
	InterestRateModel       common.Address
	Underlying              common.Address
	AvailableLiquidity      *big.Int
	ExpectedLiquidity       *big.Int
	BaseInterestIndex       *big.Int
	BaseInterestRate        *big.Int
	DieselRate              *big.Int
	SupplyRate              *big.Int
	WithdrawFee             *big.Int
	TotalBorrowed           *big.Int
	TotalDebtLimit          *big.Int
	CreditManagerDebtParams []CreditManagerDebtParams
	BaseInterestIndexLU     *big.Int
	ExpectedLiquidityLU     *big.Int
	QuotaRevenue            *big.Int
	LastBaseInterestUpdate  *big.Int
	LastQuotaRevenueUpdate  *big.Int
	IsPaused                bool
}

// PriceFeedAnswer is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedAnswer struct {
	Price     *big.Int
	UpdatedAt *big.Int
	Success   bool
}

// PriceFeedMapEntry is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedMapEntry struct {
	Token           common.Address
	Reserve         bool
	PriceFeed       common.Address
	StalenessPeriod uint32
}

// PriceFeedTreeNode is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedTreeNode struct {
	BaseParams                 BaseParams
	Description                string
	Decimals                   uint8
	SkipCheck                  bool
	Updatable                  bool
	UnderlyingFeeds            []common.Address
	UnderlyingStalenessPeriods []uint32
	Answer                     PriceFeedAnswer
}

// PriceOracleState is an auto generated low-level Go binding around an user-defined struct.
type PriceOracleState struct {
	BaseParams    BaseParams
	PriceFeedMap  []PriceFeedMapEntry
	PriceFeedTree []PriceFeedTreeNode
}

// QuotaKeeperState is an auto generated low-level Go binding around an user-defined struct.
type QuotaKeeperState struct {
	BaseParams          BaseParams
	RateKeeper          common.Address
	Quotas              []QuotaTokenParams
	CreditManagers      []common.Address
	LastQuotaRateUpdate *big.Int
}

// QuotaTokenParams is an auto generated low-level Go binding around an user-defined struct.
type QuotaTokenParams struct {
	Token             common.Address
	Rate              uint16
	CumulativeIndexLU *big.Int
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
}

// Rate is an auto generated low-level Go binding around an user-defined struct.
type Rate struct {
	Token common.Address
	Rate  uint16
}

// RateKeeperState is an auto generated low-level Go binding around an user-defined struct.
type RateKeeperState struct {
	BaseParams BaseParams
	Rates      []Rate
}

// TokenData is an auto generated low-level Go binding around an user-defined struct.
type TokenData struct {
	Addr     common.Address
	Symbol   string
	Name     string
	Decimals uint8
}

// MarketCompressorMetaData contains all meta data concerning the MarketCompressor contract.
var MarketCompressorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"addressProvider_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInterestRateModelState\",\"inputs\":[{\"name\":\"interestRateModel\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLossPolicyState\",\"inputs\":[{\"name\":\"lossPolicy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketData\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structMarketData\",\"components\":[{\"name\":\"acl\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractsRegister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"tuple\",\"internalType\":\"structPoolState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interestRateModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidityLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRevenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"lastQuotaRevenueUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"quotaKeeper\",\"type\":\"tuple\",\"internalType\":\"structQuotaKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rateKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaTokenParams[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"lastQuotaRateUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"name\":\"interestRateModel\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"rateKeeper\",\"type\":\"tuple\",\"internalType\":\"structRateKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rates\",\"type\":\"tuple[]\",\"internalType\":\"structRate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"name\":\"priceOracle\",\"type\":\"tuple\",\"internalType\":\"structPriceOracleState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"priceFeedMap\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedMapEntry[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceFeedTree\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedTreeNode[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"skipCheck\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"updatable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"underlyingFeeds\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlyingStalenessPeriods\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"answer\",\"type\":\"tuple\",\"internalType\":\"structPriceFeedAnswer\",\"components\":[{\"name\":\"price\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"lossPolicy\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenData[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"creditManagers\",\"type\":\"tuple[]\",\"internalType\":\"structCreditSuiteData[]\",\"components\":[{\"name\":\"creditFacade\",\"type\":\"tuple\",\"internalType\":\"structCreditFacadeState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"degenNFT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"botList\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManager\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"accountFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxEnabledTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"collateralTokens\",\"type\":\"tuple[]\",\"internalType\":\"structCollateralToken[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"creditConfigurator\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"accountFactory\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"adapters\",\"type\":\"tuple[]\",\"internalType\":\"structAdapterState[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"configurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"unpausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"emergencyLiquidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketData\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structMarketData\",\"components\":[{\"name\":\"acl\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractsRegister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"tuple\",\"internalType\":\"structPoolState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interestRateModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidityLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRevenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"lastQuotaRevenueUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"quotaKeeper\",\"type\":\"tuple\",\"internalType\":\"structQuotaKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rateKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaTokenParams[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"lastQuotaRateUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"name\":\"interestRateModel\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"rateKeeper\",\"type\":\"tuple\",\"internalType\":\"structRateKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rates\",\"type\":\"tuple[]\",\"internalType\":\"structRate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"name\":\"priceOracle\",\"type\":\"tuple\",\"internalType\":\"structPriceOracleState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"priceFeedMap\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedMapEntry[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceFeedTree\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedTreeNode[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"skipCheck\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"updatable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"underlyingFeeds\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlyingStalenessPeriods\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"answer\",\"type\":\"tuple\",\"internalType\":\"structPriceFeedAnswer\",\"components\":[{\"name\":\"price\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"lossPolicy\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenData[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"creditManagers\",\"type\":\"tuple[]\",\"internalType\":\"structCreditSuiteData[]\",\"components\":[{\"name\":\"creditFacade\",\"type\":\"tuple\",\"internalType\":\"structCreditFacadeState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"degenNFT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"botList\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManager\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"accountFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxEnabledTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"collateralTokens\",\"type\":\"tuple[]\",\"internalType\":\"structCollateralToken[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"creditConfigurator\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"accountFactory\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"adapters\",\"type\":\"tuple[]\",\"internalType\":\"structAdapterState[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"configurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"unpausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"emergencyLiquidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarkets\",\"inputs\":[{\"name\":\"filter\",\"type\":\"tuple\",\"internalType\":\"structMarketFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple[]\",\"internalType\":\"structMarketData[]\",\"components\":[{\"name\":\"acl\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractsRegister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"tuple\",\"internalType\":\"structPoolState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interestRateModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidityLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRevenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"lastQuotaRevenueUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"quotaKeeper\",\"type\":\"tuple\",\"internalType\":\"structQuotaKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rateKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaTokenParams[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"lastQuotaRateUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]},{\"name\":\"interestRateModel\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"rateKeeper\",\"type\":\"tuple\",\"internalType\":\"structRateKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rates\",\"type\":\"tuple[]\",\"internalType\":\"structRate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"name\":\"priceOracle\",\"type\":\"tuple\",\"internalType\":\"structPriceOracleState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"priceFeedMap\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedMapEntry[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stalenessPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"priceFeedTree\",\"type\":\"tuple[]\",\"internalType\":\"structPriceFeedTreeNode[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"skipCheck\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"updatable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"underlyingFeeds\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlyingStalenessPeriods\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"answer\",\"type\":\"tuple\",\"internalType\":\"structPriceFeedAnswer\",\"components\":[{\"name\":\"price\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}]},{\"name\":\"lossPolicy\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenData[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"name\":\"creditManagers\",\"type\":\"tuple[]\",\"internalType\":\"structCreditSuiteData[]\",\"components\":[{\"name\":\"creditFacade\",\"type\":\"tuple\",\"internalType\":\"structCreditFacadeState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"degenNFT\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"botList\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDebt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManager\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"accountFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditConfigurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxEnabledTokens\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"collateralTokens\",\"type\":\"tuple[]\",\"internalType\":\"structCollateralToken[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"feeInterest\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscount\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"name\":\"creditConfigurator\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"accountFactory\",\"type\":\"tuple\",\"internalType\":\"structBaseState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"adapters\",\"type\":\"tuple[]\",\"internalType\":\"structAdapterState[]\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"targetContract\",\"type\":\"address\",\"internalType\":\"address\"}]}]},{\"name\":\"configurator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"unpausableAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"emergencyLiquidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolState\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structPoolState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interestRateModel\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseInterestRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dieselRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBorrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\",\"internalType\":\"structCreditManagerDebtParams[]\",\"components\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"borrowed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedLiquidityLU\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quotaRevenue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"lastQuotaRevenueUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"isPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuotaKeeperState\",\"inputs\":[{\"name\":\"quotaKeeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structQuotaKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rateKeeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotas\",\"type\":\"tuple[]\",\"internalType\":\"structQuotaTokenParams[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\",\"internalType\":\"uint192\"},{\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"totalQuoted\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"limit\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"lastQuotaRateUpdate\",\"type\":\"uint40\",\"internalType\":\"uint40\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateKeeperState\",\"inputs\":[{\"name\":\"rateKeeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structRateKeeperState\",\"components\":[{\"name\":\"baseParams\",\"type\":\"tuple\",\"internalType\":\"structBaseParams\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"serializedParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"rates\",\"type\":\"tuple[]\",\"internalType\":\"structRate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rate\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// MarketCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketCompressorMetaData.ABI instead.
var MarketCompressorABI = MarketCompressorMetaData.ABI

// MarketCompressor is an auto generated Go binding around an Ethereum contract.
type MarketCompressor struct {
	MarketCompressorCaller     // Read-only binding to the contract
	MarketCompressorTransactor // Write-only binding to the contract
	MarketCompressorFilterer   // Log filterer for contract events
}

// MarketCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketCompressorSession struct {
	Contract     *MarketCompressor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCompressorCallerSession struct {
	Contract *MarketCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MarketCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketCompressorTransactorSession struct {
	Contract     *MarketCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MarketCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketCompressorRaw struct {
	Contract *MarketCompressor // Generic contract binding to access the raw methods on
}

// MarketCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCompressorCallerRaw struct {
	Contract *MarketCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// MarketCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketCompressorTransactorRaw struct {
	Contract *MarketCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketCompressor creates a new instance of MarketCompressor, bound to a specific deployed contract.
func NewMarketCompressor(address common.Address, backend bind.ContractBackend) (*MarketCompressor, error) {
	contract, err := bindMarketCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketCompressor{MarketCompressorCaller: MarketCompressorCaller{contract: contract}, MarketCompressorTransactor: MarketCompressorTransactor{contract: contract}, MarketCompressorFilterer: MarketCompressorFilterer{contract: contract}}, nil
}

// NewMarketCompressorCaller creates a new read-only instance of MarketCompressor, bound to a specific deployed contract.
func NewMarketCompressorCaller(address common.Address, caller bind.ContractCaller) (*MarketCompressorCaller, error) {
	contract, err := bindMarketCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCompressorCaller{contract: contract}, nil
}

// NewMarketCompressorTransactor creates a new write-only instance of MarketCompressor, bound to a specific deployed contract.
func NewMarketCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketCompressorTransactor, error) {
	contract, err := bindMarketCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCompressorTransactor{contract: contract}, nil
}

// NewMarketCompressorFilterer creates a new log filterer instance of MarketCompressor, bound to a specific deployed contract.
func NewMarketCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketCompressorFilterer, error) {
	contract, err := bindMarketCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketCompressorFilterer{contract: contract}, nil
}

// bindMarketCompressor binds a generic wrapper to an already deployed contract.
func bindMarketCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketCompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCompressor *MarketCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCompressor.Contract.MarketCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCompressor *MarketCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCompressor.Contract.MarketCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCompressor *MarketCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCompressor.Contract.MarketCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketCompressor *MarketCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketCompressor *MarketCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketCompressor *MarketCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketCompressor.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MarketCompressor *MarketCompressorCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MarketCompressor *MarketCompressorSession) AddressProvider() (common.Address, error) {
	return _MarketCompressor.Contract.AddressProvider(&_MarketCompressor.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_MarketCompressor *MarketCompressorCallerSession) AddressProvider() (common.Address, error) {
	return _MarketCompressor.Contract.AddressProvider(&_MarketCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_MarketCompressor *MarketCompressorCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_MarketCompressor *MarketCompressorSession) ContractType() ([32]byte, error) {
	return _MarketCompressor.Contract.ContractType(&_MarketCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_MarketCompressor *MarketCompressorCallerSession) ContractType() ([32]byte, error) {
	return _MarketCompressor.Contract.ContractType(&_MarketCompressor.CallOpts)
}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorCaller) GetInterestRateModelState(opts *bind.CallOpts, interestRateModel common.Address) (BaseState, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getInterestRateModelState", interestRateModel)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorSession) GetInterestRateModelState(interestRateModel common.Address) (BaseState, error) {
	return _MarketCompressor.Contract.GetInterestRateModelState(&_MarketCompressor.CallOpts, interestRateModel)
}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorCallerSession) GetInterestRateModelState(interestRateModel common.Address) (BaseState, error) {
	return _MarketCompressor.Contract.GetInterestRateModelState(&_MarketCompressor.CallOpts, interestRateModel)
}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorCaller) GetLossPolicyState(opts *bind.CallOpts, lossPolicy common.Address) (BaseState, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getLossPolicyState", lossPolicy)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorSession) GetLossPolicyState(lossPolicy common.Address) (BaseState, error) {
	return _MarketCompressor.Contract.GetLossPolicyState(&_MarketCompressor.CallOpts, lossPolicy)
}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_MarketCompressor *MarketCompressorCallerSession) GetLossPolicyState(lossPolicy common.Address) (BaseState, error) {
	return _MarketCompressor.Contract.GetLossPolicyState(&_MarketCompressor.CallOpts, lossPolicy)
}

// GetMarketData is a free data retrieval call binding the contract method 0x34a2b766.
//
// Solidity: function getMarketData(address pool, address configurator) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorCaller) GetMarketData(opts *bind.CallOpts, pool common.Address, configurator common.Address) (MarketData, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getMarketData", pool, configurator)

	if err != nil {
		return *new(MarketData), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketData)).(*MarketData)

	return out0, err

}

// GetMarketData is a free data retrieval call binding the contract method 0x34a2b766.
//
// Solidity: function getMarketData(address pool, address configurator) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorSession) GetMarketData(pool common.Address, configurator common.Address) (MarketData, error) {
	return _MarketCompressor.Contract.GetMarketData(&_MarketCompressor.CallOpts, pool, configurator)
}

// GetMarketData is a free data retrieval call binding the contract method 0x34a2b766.
//
// Solidity: function getMarketData(address pool, address configurator) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorCallerSession) GetMarketData(pool common.Address, configurator common.Address) (MarketData, error) {
	return _MarketCompressor.Contract.GetMarketData(&_MarketCompressor.CallOpts, pool, configurator)
}

// GetMarketData0 is a free data retrieval call binding the contract method 0xa30c302d.
//
// Solidity: function getMarketData(address pool) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorCaller) GetMarketData0(opts *bind.CallOpts, pool common.Address) (MarketData, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getMarketData0", pool)

	if err != nil {
		return *new(MarketData), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketData)).(*MarketData)

	return out0, err

}

// GetMarketData0 is a free data retrieval call binding the contract method 0xa30c302d.
//
// Solidity: function getMarketData(address pool) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorSession) GetMarketData0(pool common.Address) (MarketData, error) {
	return _MarketCompressor.Contract.GetMarketData0(&_MarketCompressor.CallOpts, pool)
}

// GetMarketData0 is a free data retrieval call binding the contract method 0xa30c302d.
//
// Solidity: function getMarketData(address pool) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[]) result)
func (_MarketCompressor *MarketCompressorCallerSession) GetMarketData0(pool common.Address) (MarketData, error) {
	return _MarketCompressor.Contract.GetMarketData0(&_MarketCompressor.CallOpts, pool)
}

// GetMarkets is a free data retrieval call binding the contract method 0x0736779d.
//
// Solidity: function getMarkets((address[],address[],address) filter) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[])[] result)
func (_MarketCompressor *MarketCompressorCaller) GetMarkets(opts *bind.CallOpts, filter MarketFilter) ([]MarketData, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getMarkets", filter)

	if err != nil {
		return *new([]MarketData), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketData)).(*[]MarketData)

	return out0, err

}

// GetMarkets is a free data retrieval call binding the contract method 0x0736779d.
//
// Solidity: function getMarkets((address[],address[],address) filter) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[])[] result)
func (_MarketCompressor *MarketCompressorSession) GetMarkets(filter MarketFilter) ([]MarketData, error) {
	return _MarketCompressor.Contract.GetMarkets(&_MarketCompressor.CallOpts, filter)
}

// GetMarkets is a free data retrieval call binding the contract method 0x0736779d.
//
// Solidity: function getMarkets((address[],address[],address) filter) view returns((address,address,address,((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool),((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),(address,uint16)[]),((address,uint256,bytes32,bytes),(address,bool,address,uint32)[],((address,uint256,bytes32,bytes),string,uint8,bool,bool,address[],uint32[],(int256,uint256,bool))[]),((address,uint256,bytes32,bytes)),(address,string,string,uint8)[],(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[],address,address[],address[],address[])[] result)
func (_MarketCompressor *MarketCompressorCallerSession) GetMarkets(filter MarketFilter) ([]MarketData, error) {
	return _MarketCompressor.Contract.GetMarkets(&_MarketCompressor.CallOpts, filter)
}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_MarketCompressor *MarketCompressorCaller) GetPoolState(opts *bind.CallOpts, pool common.Address) (PoolState, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getPoolState", pool)

	if err != nil {
		return *new(PoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolState)).(*PoolState)

	return out0, err

}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_MarketCompressor *MarketCompressorSession) GetPoolState(pool common.Address) (PoolState, error) {
	return _MarketCompressor.Contract.GetPoolState(&_MarketCompressor.CallOpts, pool)
}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_MarketCompressor *MarketCompressorCallerSession) GetPoolState(pool common.Address) (PoolState, error) {
	return _MarketCompressor.Contract.GetPoolState(&_MarketCompressor.CallOpts, pool)
}

// GetQuotaKeeperState is a free data retrieval call binding the contract method 0x2e834de9.
//
// Solidity: function getQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_MarketCompressor *MarketCompressorCaller) GetQuotaKeeperState(opts *bind.CallOpts, quotaKeeper common.Address) (QuotaKeeperState, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getQuotaKeeperState", quotaKeeper)

	if err != nil {
		return *new(QuotaKeeperState), err
	}

	out0 := *abi.ConvertType(out[0], new(QuotaKeeperState)).(*QuotaKeeperState)

	return out0, err

}

// GetQuotaKeeperState is a free data retrieval call binding the contract method 0x2e834de9.
//
// Solidity: function getQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_MarketCompressor *MarketCompressorSession) GetQuotaKeeperState(quotaKeeper common.Address) (QuotaKeeperState, error) {
	return _MarketCompressor.Contract.GetQuotaKeeperState(&_MarketCompressor.CallOpts, quotaKeeper)
}

// GetQuotaKeeperState is a free data retrieval call binding the contract method 0x2e834de9.
//
// Solidity: function getQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_MarketCompressor *MarketCompressorCallerSession) GetQuotaKeeperState(quotaKeeper common.Address) (QuotaKeeperState, error) {
	return _MarketCompressor.Contract.GetQuotaKeeperState(&_MarketCompressor.CallOpts, quotaKeeper)
}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_MarketCompressor *MarketCompressorCaller) GetRateKeeperState(opts *bind.CallOpts, rateKeeper common.Address) (RateKeeperState, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "getRateKeeperState", rateKeeper)

	if err != nil {
		return *new(RateKeeperState), err
	}

	out0 := *abi.ConvertType(out[0], new(RateKeeperState)).(*RateKeeperState)

	return out0, err

}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_MarketCompressor *MarketCompressorSession) GetRateKeeperState(rateKeeper common.Address) (RateKeeperState, error) {
	return _MarketCompressor.Contract.GetRateKeeperState(&_MarketCompressor.CallOpts, rateKeeper)
}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_MarketCompressor *MarketCompressorCallerSession) GetRateKeeperState(rateKeeper common.Address) (RateKeeperState, error) {
	return _MarketCompressor.Contract.GetRateKeeperState(&_MarketCompressor.CallOpts, rateKeeper)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MarketCompressor *MarketCompressorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MarketCompressor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MarketCompressor *MarketCompressorSession) Version() (*big.Int, error) {
	return _MarketCompressor.Contract.Version(&_MarketCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_MarketCompressor *MarketCompressorCallerSession) Version() (*big.Int, error) {
	return _MarketCompressor.Contract.Version(&_MarketCompressor.CallOpts)
}
