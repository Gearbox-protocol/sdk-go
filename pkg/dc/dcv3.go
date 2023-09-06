package dc

import (
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
)

func getPoolDatav3(values dcv3.PoolData) PoolCallData {
	return PoolCallData{
		Addr:                  values.Addr,
		Underlying:            values.Underlying,
		DieselToken:           values.DieselToken,
		LinearCumulativeIndex: values.LinearCumulativeIndex,
		AvailableLiquidity:    values.AvailableLiquidity,
		ExpectedLiquidity:     values.ExpectedLiquidity,
		TotalBorrowed:         values.TotalBorrowed,
		// : values.//,
		CreditManagerDebtParams: values.CreditManagerDebtParams,
		TotalAssets:             values.TotalAssets,
		TotalSupply:             values.TotalSupply,
		SupplyRate:              values.SupplyRate,
		BaseInterestRate:        values.BaseInterestRate,
		DieselRateRAY:           values.DieselRateRAY,
		WithdrawFee:             values.WithdrawFee,
		CumulativeIndexRAY:      values.CumulativeIndexRAY,
		// : values.//,
		Version: values.Version,
		Quotas:  values.Quotas,
	}
}
func getCMDatav3(values dcv3.CreditManagerData) CMCallData {
	return CMCallData{
		Addr:           values.Addr,
		Underlying:     values.Underlying,
		BaseBorrowRate: values.BaseBorrowRate,
		//
		MinDebt: values.MinDebt,
		MaxDebt: values.MaxDebt,
		//
		Adapters: values.Adapters,
		Quotas:   values.Quotas,
	}
}
func getAccountDatav3(values dcv3.CreditAccountData) CreditAccountCallData {
	return CreditAccountCallData{
		Addr:                       values.Addr,
		Borrower:                   values.Borrower,
		CreditManager:              values.CreditManager,
		Underlying:                 values.Underlying,
		BorrowedAmount:             values.Debt,
		BorrowedAmountPlusInterest: nil, // DC_CHANGED

		CumulativeIndexAtOpen:   values.CumulativeIndexLastUpdate,
		CumulativeQuotaInterest: values.CumulativeQuotaInterest,
		TotalValue:              values.TotalValue,
		HealthFactor:            values.HealthFactor,
		BaseBorrowRate:          values.BaseBorrowRate,
		Since:                   values.Since,
		Balances:                convertv3ToBalance(values.Balances),
		Version:                 3,
	}
}

func convertv3ToBalance(balances []dcv3.TokenBalance) (dcv2Balances []TokenBalanceCallData) {
	for _, balance := range balances {
		dcv2Balances = append(dcv2Balances, TokenBalanceCallData{
			Token:       balance.Token,
			Balance:     balance.Balance,
			IsForbidden: balance.IsForbidden, // is set on credit manager
			IsEnabled:   balance.IsEnabled,   // is used by credit account
			IsQuoted:    balance.IsQuoted,
			Quota:       balance.Quota,
			QuotaRate:   balance.QuotaRate,
		})
	}
	return
}
