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
		LinearCumulativeIndex: data.LinearCumulativeIndex,
		AvailableLiquidity:    data.AvailableLiquidity,
		ExpectedLiquidity:     data.ExpectedLiquidity,
		TotalBorrowed:         data.TotalBorrowed,
		// TotalDebtLimit: ,
		// CreditManagerDebtParams: ,
		TotalAssets:        data.ExpectedLiquidity, // D_BY_US _expectedLiquidityLU + _calcBaseInterestAccrued() + _calcQuotaRevenueAccrued()
		TotalSupply:        new(big.Int).Add(data.AvailableLiquidity, data.TotalBorrowed),
		SupplyRate:         data.DepositAPYRAY,
		BaseInterestRate:   data.BorrowAPYRAY,
		DieselRateRAY:      data.DieselRateRAY,
		WithdrawFee:        data.WithdrawFee,
		CumulativeIndexRAY: data.CumulativeIndexRAY,
		// BaseInterestIndexLU: data.CumulativeIndexRAY,
		Version: big.NewInt(2),
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
		BaseBorrowRate: values.BorrowRate,
		//
		MinDebt: values.MinAmount,
		MaxDebt: values.MaxAmount,
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
		BorrowedAmount:             values.BorrowedAmount,
		BorrowedAmountPlusInterest: values.BorrowedAmountPlusInterest,
		// values.BorrowedAmountPlusInterest,
		CumulativeIndexAtOpen:   values.CumulativeIndexAtOpen,
		CumulativeQuotaInterest: new(big.Int), // D_BY_US
		TotalValue:              values.TotalValue,
		HealthFactor:            values.HealthFactor,
		BaseBorrowRate:          values.BorrowRate,
		Since:                   uint64(values.Since.Int64()),
		Balances:                convertv2ToBalance(values.Balances),
		Version:                 core.NewVersion(int16(values.Version)),
		//
		RepayAmountv1v2: values.RepayAmount,
	}
}
func convertv2ToBalance(balances []dcv2.TokenBalance) (dcv2Balances []TokenBalanceCallData) {
	for _, balance := range balances {
		dcv2Balances = append(dcv2Balances, TokenBalanceCallData{
			Token:       balance.Token,
			Balance:     balance.Balance,
			IsForbidden: balance.IsAllowed, // is set on credit manager
			IsEnabled:   balance.IsEnabled, // is used by credit account
			IsQuoted:    false,
			Quota:       new(big.Int),
			QuotaRate:   0,
		})
	}
	return
}
