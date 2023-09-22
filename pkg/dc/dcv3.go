package dc

import (
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
)

func getPoolDatav3(values dcv3.PoolData) PoolCallData {
	return PoolCallData{
		Addr:                  values.Addr,
		Underlying:            values.Underlying,
		DieselToken:           values.DieselToken,
		LinearCumulativeIndex: (*core.BigInt)(values.LinearCumulativeIndex),
		AvailableLiquidity:    (*core.BigInt)(values.AvailableLiquidity),
		ExpectedLiquidity:     (*core.BigInt)(values.ExpectedLiquidity),
		TotalBorrowed:         (*core.BigInt)(values.TotalBorrowed),
		// : values.//,
		CreditManagerDebtParams: values.CreditManagerDebtParams,
		TotalAssets:             (*core.BigInt)(values.TotalAssets),
		TotalSupply:             (*core.BigInt)(values.TotalSupply),
		SupplyRate:              (*core.BigInt)(values.SupplyRate),
		BaseInterestRate:        (*core.BigInt)(values.BaseInterestRate),
		DieselRateRAY:           (*core.BigInt)(values.DieselRateRAY),
		WithdrawFee:             (*core.BigInt)(values.WithdrawFee),
		CumulativeIndexRAY:      (*core.BigInt)(values.CumulativeIndexRAY),
		// : values.//,
		Version: (*core.BigInt)(values.Version),
		Quotas:  values.Quotas,
	}
}
func getCMDatav3(values dcv3.CreditManagerData) CMCallData {
	return CMCallData{
		Addr:           values.Addr,
		Underlying:     values.Underlying,
		BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
		//
		MinDebt: (*core.BigInt)(values.MinDebt),
		MaxDebt: (*core.BigInt)(values.MaxDebt),
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
		BorrowedAmount:             (*core.BigInt)(values.Debt),
		BorrowedAmountPlusInterest: nil, // DC_CHANGED

		CumulativeIndexAtOpen:   (*core.BigInt)(values.CumulativeIndexLastUpdate),
		CumulativeQuotaInterest: (*core.BigInt)(values.CumulativeQuotaInterest),
		//
		AccruedInterest: (*core.BigInt)(values.AccruedInterest),
		AccruedFees:     (*core.BigInt)(values.AccruedFees),
		//
		TotalValue:     (*core.BigInt)(values.TotalValue),
		HealthFactor:   (*core.BigInt)(values.HealthFactor),
		BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
		Since:          values.Since,
		Balances:       convertv3ToBalance(values.Balances),
		Version:        core.NewVersion(3),
	}
}

func convertv3ToBalance(balances []dcv3.TokenBalance) (dcv2Balances []core.TokenBalanceCallData) {
	for ind, balance := range balances {
		dcv2Balances = append(dcv2Balances, core.TokenBalanceCallData{
			Token: balance.Token.Hex(),
			DBTokenBalance: core.DBTokenBalance{
				BI:          (*core.BigInt)(balance.Balance),
				IsForbidden: balance.IsForbidden, // is set on credit manager
				IsEnabled:   balance.IsEnabled,   // is used by credit account
				IsQuoted:    balance.IsQuoted,
				Quota:       (*core.BigInt)(balance.Quota),
				QuotaRate:   balance.QuotaRate,
				Ind:         ind,
			},
		})
	}
	return
}
