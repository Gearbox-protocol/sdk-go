package dc

import (
	"math/big"

	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
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
	LinearCumulativeIndex *core.BigInt   `json:"linearCumulativeIndex"`
	AvailableLiquidity    *core.BigInt   `json:"availableLiquidity"`
	ExpectedLiquidity     *core.BigInt   `json:"expectedLiquidity"`
	TotalBorrowed         *core.BigInt   `json:"totalBorrowed"`
	// TotalDebtLimit          *core.BigInt
	CreditManagerDebtParams []dcv3.CreditManagerDebtParams `json:"creditManagerDebtParams,omitempty"`
	TotalAssets             *core.BigInt                   `json:"-"`
	TotalSupply             *core.BigInt                   `json:"-"`
	SupplyRate              *core.BigInt                   `json:"depositAPY"`
	BaseInterestRate        *core.BigInt                   `json:"baseBorrowRate"` // DC_CHANGED
	DieselRateRAY           *core.BigInt                   `json:"dieselRate"`
	WithdrawFee             *core.BigInt                   `json:"withdrawFee"`
	CumulativeIndexRAY      *core.BigInt                   `json:"cumulativeIndex"`
	// BaseInterestIndexLU     *core.BigInt
	Version *core.BigInt     `json:"version"`
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
	BaseBorrowRate *core.BigInt `json:"baseBorrowRate"` // DC_CHANGED
	MinDebt        *core.BigInt `json:"minAmount"`
	MaxDebt        *core.BigInt `json:"maxAmount"`
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
	BorrowedAmount             *core.BigInt   `json:"borrowedAmount"` // DC_CHANGED
	BorrowedAmountPlusInterest *core.BigInt   `json:"borrowedAmountPlusInterest"`
	// TODO:"totalBorrowedWithInterst" is removed
	CumulativeIndexAtOpen   *core.BigInt `json:"cumulativeIndexAtOpen"`
	CumulativeQuotaInterest *core.BigInt `json:"cumulativeQuotaInterest,omitempty"`
	QuotaFeeCalc
	TotalValue *core.BigInt `json:"totalValue"`

	RepayAmountv1v2 *core.BigInt `json:"repayAmount,omitempty"`
	// TotalValueUSD
	// TwvUSD
	// EnabledTokensMask
	HealthFactor   *core.BigInt `json:"healthFactor"`
	BaseBorrowRate *core.BigInt `json:"baseBorrowRate"` // DC_CHANGED
	// AggregatedBorrowRate
	Since    uint64                      `json:"since"`
	Balances []core.TokenBalanceCallData `json:"balances"`
}

type QuotaFeeCalc struct {
	AccruedInterest *core.BigInt     `json:"accruedInterest,omitempty"`
	AccruedFees     *core.BigInt     `json:"accruedFees,omitempty"`
	Version         core.VersionType `json:"version"`
}

func (data QuotaFeeCalc) GetQuotaFees(feeInterest uint16) *big.Int {
	if !data.Version.Eq(3) {
		return new(big.Int)
	}
	interestFees := new(big.Int).Quo(
		new(big.Int).Mul(
			data.AccruedInterest.Convert(), big.NewInt(int64(feeInterest)),
		),
		utils.GetExpInt(4),
	)
	// quota fees
	return new(big.Int).Sub(data.AccruedFees.Convert(), interestFees)
}
