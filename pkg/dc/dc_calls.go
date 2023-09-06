package dc

import (
	"math/big"

	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/common"
)

type DCCalls struct {
	Pools    map[string]PoolCallData
	CMs      map[string]CMCallData
	Accounts map[string]CreditAccountCallData
}

func NewDCCalls() *DCCalls {
	return &DCCalls{
		Pools:    make(map[string]PoolCallData),
		CMs:      make(map[string]CMCallData),
		Accounts: make(map[string]CreditAccountCallData),
	}
}

type TestTokenBalance struct {
	Token     string       `json:"token"`
	Balance   *core.BigInt `json:"balance"`
	IsAllowed bool         `json:"isAllowed"`
	IsEnabled bool         `json:"isEnabled"`
}

type PoolCallData struct {
	Addr                  common.Address `json:"address"`
	Underlying            common.Address `json:"underlyingToken"`
	DieselToken           common.Address `json:"dieselToken"`
	LinearCumulativeIndex *big.Int       `json:"linearCumulativeIndex"`
	AvailableLiquidity    *big.Int       `json:"availableLiquidity"`
	ExpectedLiquidity     *big.Int       `json:"expectedLiquidity"`
	TotalBorrowed         *big.Int       `json:"totalBorrowed"`
	// TotalDebtLimit          *big.Int
	CreditManagerDebtParams []dcv3.CreditManagerDebtParams `json:"creditManagerDebtParams,omitempty"`
	TotalAssets             *big.Int                       `json:"-"`
	TotalSupply             *big.Int                       `json:"-"`
	SupplyRate              *big.Int                       `json:"depositAPY"`
	BaseInterestRate        *big.Int                       `json:"baseBorrowRate"` // DC_CHANGED
	DieselRateRAY           *big.Int                       `json:"dieselRate"`
	WithdrawFee             *big.Int                       `json:"withdrawFee"`
	CumulativeIndexRAY      *big.Int                       `json:"cumulativeIndex"`
	// BaseInterestIndexLU     *big.Int
	Version *big.Int         `json:"version"`
	Quotas  []dcv3.QuotaInfo `json:"quotas,omitempty"`
}

type CMCallData struct {
	Addr common.Address
	// CfVersion                  *big.Int
	// CreditFacade               common.Address
	// CreditConfigurator         common.Address
	Underlying common.Address
	// Pool                       common.Address
	// TotalDebt                  *big.Int
	// TotalDebtLimit             *big.Int
	BaseBorrowRate *big.Int `json:"baseBorrowRate"` // DC_CHANGED
	MinDebt        *big.Int `json:"minAmount"`
	MaxDebt        *big.Int `json:"maxAmount"`
	// AvailableToBorrow          *big.Int
	// CollateralTokens           []common.Address
	Adapters []dcv3.ContractAdapter `json:"adapters"`
	// LiquidationThresholds      []*big.Int
	// IsDegenMode                bool
	// DegenNFT                   common.Address
	// ForbiddenTokenMask         *big.Int
	// MaxEnabledTokensLength     uint8
	// FeeInterest                uint16
	// FeeLiquidation             uint16
	// LiquidationDiscount        uint16
	// FeeLiquidationExpired      uint16
	// LiquidationDiscountExpired uint16
	Quotas []dcv3.QuotaInfo `json:"quotas,omitempty"`
	// IsPaused bool
}

type CreditAccountCallData struct {
	Addr                       common.Address `json:"address"`
	Borrower                   common.Address `json:"borrower"`
	CreditManager              common.Address `json:"creditManager"`
	Underlying                 common.Address `json:"underlyingToken"`
	BorrowedAmount             *big.Int       `json:"borrowedAmount"` // DC_CHANGED
	BorrowedAmountPlusInterest *big.Int       `json:"borrowedAmountPlusInterest"`
	// TODO:"totalBorrowedWithInterst" is removed
	CumulativeIndexAtOpen   *big.Int `json:"cumulativeIndexAtOpen"`
	CumulativeQuotaInterest *big.Int `json:"cumulativeQuotaInterest,omitempty"`
	AccruedInterest         *big.Int `json:"accruedInterest,omitempty"`
	AccruedFees             *big.Int `json:"accruedFees,omitempty"`
	TotalValue              *big.Int `json:"totalValue"`

	RepayAmountv1v2 *big.Int `json:"repayAmount,omitempty"`
	// TotalValueUSD
	// TwvUSD
	// EnabledTokensMask
	HealthFactor   *big.Int `json:"healthFactor"`
	BaseBorrowRate *big.Int `json:"baseBorrowRate"` // DC_CHANGED
	// AggregatedBorrowRate
	Since    uint64                 `json:"since"`
	Balances []TokenBalanceCallData `json:"balances"`
	Version  core.VersionType       `json:"version"`
}

type TokenBalanceCallData struct {
	Token       common.Address `json:"token"`
	Balance     *big.Int       `json:"balance"`
	IsForbidden bool           `json:"isForbidden"`
	IsEnabled   bool           `json:"isEnabled"`
	IsQuoted    bool           `json:"isQuoted"`
	Quota       *big.Int       `json:"quota"`
	QuotaRate   uint16         `json:"quotaRate"`
}
