package dc

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
)

func getPoolDatav2(blockNum int64, data dcv2.PoolData) PoolCallData {
	latestFormat := PoolCallData{
		Addr:                  data.Addr,
		Underlying:            data.Underlying,
		DieselToken:           data.DieselToken,
		LinearCumulativeIndex: (*core.BigInt)(data.LinearCumulativeIndex),
		AvailableLiquidity:    (*core.BigInt)(data.AvailableLiquidity),
		ExpectedLiquidity:     (*core.BigInt)(data.ExpectedLiquidity),
		TotalBorrowed:         (*core.BigInt)(data.TotalBorrowed),
		// TotalDebtLimit: ,
		// CreditManagerDebtParams: ,
		TotalAssets:        (*core.BigInt)(data.ExpectedLiquidity), // D_BY_US _expectedLiquidityLU + _calcBaseInterestAccrued() + _calcQuotaRevenueAccrued()
		TotalSupply:        (*core.BigInt)(new(big.Int).Add(data.AvailableLiquidity, data.TotalBorrowed)),
		SupplyRate:         (*core.BigInt)(data.DepositAPYRAY),
		BaseInterestRate:   (*core.BigInt)(data.BorrowAPYRAY),
		DieselRateRAY:      (*core.BigInt)(data.DieselRateRAY),
		WithdrawFee:        (*core.BigInt)(data.WithdrawFee),
		CumulativeIndexRAY: (*core.BigInt)(data.CumulativeIndexRAY),
		// BaseInterestIndexLU: data.CumulativeIndexRAY,
		Version: (*core.BigInt)(big.NewInt(2)),
		// Quotas:              nil,
		//
		//
		// IsWETH:                 data.IsWETH,
		// TimestampLU:        data.TimestampLU,
		// ExpectedLiquidityLimit
	}
	//
	return latestFormat
}

func getCMDatav2(values dcv2.CreditManagerData) CMCallData {
	return CMCallData{
		Addr:           values.Addr,
		Underlying:     values.Underlying,
		BaseBorrowRate: (*core.BigInt)(values.BorrowRate),
		//
		MinDebt: (*core.BigInt)(values.MinAmount),
		MaxDebt: (*core.BigInt)(values.MaxAmount),
		//
		Adapters: func() (ans []dcv3.ContractAdapter) {
			for _, adapter := range values.Adapters {
				ans = append(ans, dcv3.ContractAdapter{
					TargetContract: adapter.AllowedContract,
					Adapter:        adapter.Adapter,
				})
			}
			return
		}(),
		Quotas: nil,
	}
}
func getAccountDatav2(values dcv2.CreditAccountData) CreditAccountCallData {
	return CreditAccountCallData{
		Addr:                       values.Addr,
		Borrower:                   values.Borrower,
		CreditManager:              values.CreditManager,
		Underlying:                 values.Underlying,
		BorrowedAmount:             (*core.BigInt)(values.BorrowedAmount),
		BorrowedAmountPlusInterest: (*core.BigInt)(values.BorrowedAmountPlusInterest),
		// (*core.BigInt)(values.BorrowedAmountPlusInterest),
		CumulativeIndexAtOpen:   (*core.BigInt)(values.CumulativeIndexAtOpen),
		CumulativeQuotaInterest: new(core.BigInt), // D_BY_US
		//
		AccruedInterest: (*core.BigInt)(new(big.Int)),
		AccruedFees:     (*core.BigInt)(new(big.Int)),
		//
		TotalValue:     (*core.BigInt)(values.TotalValue),
		HealthFactor:   (*core.BigInt)(values.HealthFactor),
		BaseBorrowRate: (*core.BigInt)(values.BorrowRate),
		Since:          uint64(values.Since.Int64()),
		Balances:       Convertv2ToBalance(values.Balances),
		Version:        core.NewVersion(int16(values.Version)),
		//
		RepayAmountv1v2: (*core.BigInt)(values.RepayAmount),
	}
}
func Convertv2ToBalance(balances []dcv2.TokenBalance) (dcv2Balances []core.TokenBalanceCallData) {
	for ind, balance := range balances {
		dcv2Balances = append(dcv2Balances, core.TokenBalanceCallData{
			Token: balance.Token.Hex(),
			DBTokenBalance: core.DBTokenBalance{
				BI:          (*core.BigInt)(balance.Balance),
				IsForbidden: !balance.IsAllowed, // is set on credit manager
				IsEnabled:   balance.IsEnabled,  // is used by credit account
				IsQuoted:    false,
				Quota:       new(core.BigInt),
				QuotaRate:   0,
				Ind:         ind,
			},
		})
	}
	return
}
