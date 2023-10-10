package dc

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

// only called for v1 , doesn't have logic for EnabledTokensMap(Gearboxv2)
func GetMask(client core.ClientI, blockNum int64, cfAddr, accountAddr common.Address) *big.Int {
	extra := make([]byte, 12)
	extra = append(extra, accountAddr.Bytes()...)
	returnData, err := core.CallFuncWithExtraBytes(client, "b451cecc", cfAddr, blockNum, extra) // enabledTokens
	log.CheckFatal(err)
	return new(big.Int).SetBytes(returnData)

}

func getPoolDataV1(data mainnet.DataTypesPoolData) PoolCallData {
	return PoolCallData{
		Addr:        data.Addr,
		Underlying:  data.UnderlyingToken,
		DieselToken: data.DieselToken,
		//
		LinearCumulativeIndex: (*core.BigInt)(data.LinearCumulativeIndex),
		AvailableLiquidity:    (*core.BigInt)(data.AvailableLiquidity),
		ExpectedLiquidity:     (*core.BigInt)(data.ExpectedLiquidity),
		TotalBorrowed:         (*core.BigInt)(data.TotalBorrowed),
		//
		TotalAssets: (*core.BigInt)(data.ExpectedLiquidity),                                        // D_BY_US _expectedLiquidityLU + _calcBaseInterestAccrued() + _calcQuotaRevenueAccrued()
		TotalSupply: (*core.BigInt)(new(big.Int).Add(data.AvailableLiquidity, data.TotalBorrowed)), // D_BY_US

		//
		SupplyRate:       (*core.BigInt)(data.DepositAPYRAY),
		BaseInterestRate: (*core.BigInt)(data.BorrowAPYRAY),
		DieselRateRAY:    (*core.BigInt)(data.DieselRateRAY),
		WithdrawFee:      (*core.BigInt)(data.WithdrawFee),
		//
		CumulativeIndexRAY: (*core.BigInt)(data.CumulativeIndexRAY),
		// BaseInterestIndexLU: data.CumulativeIndexRAY, // D_BY_US
		Version: (*core.BigInt)(big.NewInt(1)),
		//
		// TotalDebtLimit:          nil,
		CreditManagerDebtParams: nil,
		Quotas:                  nil,
		//
		// IsWETH:                 data.IsWETH,
		// ExpectedLiquidityLimit: data.ExpectedLiquidityLimit,
		// TimestampLU:            data.TimestampLU,
	}
}

func getCMDatav1(data mainnet.DataTypesCreditManagerData) CMCallData {
	return CMCallData{
		Addr: data.Addr,
		// 		CfVersion:
		// CreditFacade:
		// CreditConfigurator:
		Underlying: data.UnderlyingToken,
		// Pool:
		BaseBorrowRate: (*core.BigInt)(data.BorrowRate),
		// TotalDebt
		// TotalDebtLimit
		MinDebt: (*core.BigInt)(data.MinAmount),
		MaxDebt: (*core.BigInt)(data.MaxAmount),
		// AvailableToBorrow:
		// CollateralTokens: data.AllowedTokens,
		Adapters: nil,
		// LiquidationThresholds
		// IsDegenMode
		// DegenNFT
		// ForbiddenTokenMask
		// MaxEnabledTokensLength
		// FeeInterest
		// FeeLiquidation
		// LiquidationDiscount
		// FeeLiquidationExpired
		// LiquidationDiscountExpired
		Quotas: nil,
		// not found
		// IsPaused
		// IsWETH:             data.IsWETH,
		// CanBorrow:          data.CanBorrow,
		// BorrowRate:         data.BorrowRate,
		// MaxLeverageFactor:  data.MaxLeverageFactor,
		// AvailableLiquidity: data.AvailableLiquidity,
		// CollateralTokens:   data.AllowedTokens,
	}
}

func getCreditAccountDatav1(client core.ClientI, cfAddrv1 common.Address, blockNum int64, data mainnet.DataTypesCreditAccountDataExtended) (CreditAccountCallData, error) {
	latestFormat := CreditAccountCallData{
		// 	IsSuccessful              bool
		// PriceFeedsNeeded          []common.Address
		Addr:          data.Addr,
		Borrower:      data.Borrower,
		CreditManager: data.CreditManager,
		Underlying:    data.UnderlyingToken,
		// CreditFacade
		// Underlying
		BorrowedAmount:             (*core.BigInt)(data.BorrowedAmount),
		BorrowedAmountPlusInterest: (*core.BigInt)(data.BorrowedAmountPlusInterest),
		//:                  data.BorrowedAmountPlusInterest,
		CumulativeIndexAtOpen: (*core.BigInt)(data.CumulativeIndexAtOpen),
		// CumulativeIndexLastUpdate: nil,
		CumulativeQuotaInterest: new(core.BigInt), // D_BY_US
		//
		QuotaFeeCalc: QuotaFeeCalc{
			AccruedInterest: (*core.BigInt)(new(big.Int)),
			AccruedFees:     (*core.BigInt)(new(big.Int)),
			Version:         core.NewVersion(1),
		},
		// TotalDebtUSD
		TotalValue: (*core.BigInt)(data.TotalValue),
		// TotalValueUSD
		// TwvUSD
		// EnabledTokensMask
		HealthFactor:   (*core.BigInt)(data.HealthFactor),
		BaseBorrowRate: (*core.BigInt)(data.BorrowRate),
		// AggregatedBorrowRate
		Since: uint64(data.Since.Int64()),
		// ExpirationDate
		// ActiveBot
		// MaxApprovedBots
		// SchedultedWithdrawals

		RepayAmountv1v2: (*core.BigInt)(data.RepayAmount),
		// LiquidationAmount: data.LiquidationAmount,
		// CanBeClosed:       data.CanBeClosed,
		// BorrowedAmount:    data.BorrowedAmount,
	}
	//
	mask := GetMask(client, blockNum, cfAddrv1, latestFormat.Addr)
	latestFormat.Balances = convertv1ToBalance(data.Balances, mask)
	//
	return latestFormat, nil
}

func convertv1ToBalance(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (dcv2Balances []core.TokenBalanceCallData) {
	maskInBits := fmt.Sprintf("%b", mask)
	maskLen := len(maskInBits)
	for i, balance := range balances {
		var isEnabled bool
		if maskLen > i {
			isEnabled = maskInBits[maskLen-i-1] == '1'
		}
		dcv2Balances = append(dcv2Balances, core.TokenBalanceCallData{
			Token: balance.Token.Hex(),
			DBTokenBalance: core.DBTokenBalance{
				BI:          (*core.BigInt)(balance.Balance),
				IsForbidden: !balance.IsAllowed, // is set on credit manager
				IsEnabled:   isEnabled,          // is used by credit account
				IsQuoted:    false,
				Quota:       new(core.BigInt),
				QuotaRate:   0,
				Ind:         i,
			},
		})
	}
	return
}
