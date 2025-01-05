package dc

import (
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func getPoolDatav301(values dcv3.PoolData) PoolCallData {
	return PoolCallData{
		Addr:               values.Addr,
		Underlying:         values.Underlying,
		DieselToken:        values.DieselToken,
		AvailableLiquidity: (*core.BigInt)(values.AvailableLiquidity),
		ExpectedLiquidity:  (*core.BigInt)(values.ExpectedLiquidity),
		TotalBorrowed:      (*core.BigInt)(values.TotalBorrowed),
		// : values.//,
		// CreditManagerDebtParams: values.CreditManagerDebtParams,
		// TotalAssets:      (*core.BigInt)(values.TotalAssets),
		TotalSupply:      (*core.BigInt)(values.TotalSupply),
		SupplyRate:       (*core.BigInt)(values.SupplyRate),
		BaseInterestRate: (*core.BigInt)(values.BaseInterestRate),
		DieselRateRAY:    (*core.BigInt)(values.DieselRateRAY),
		WithdrawFee:      (*core.BigInt)(values.WithdrawFee),
		// assuming cumulativeIndexRay is depcreated
		// CumulativeIndexRAY is same as BaseInterestIndexLU
		CumulativeIndexRAY: (*core.BigInt)(values.BaseInterestIndexLU),
		// : values.//,
		Version: (*core.BigInt)(values.Version),
		// Quotas:  values.Quotas,
	}
}
func getCMDatav301(values dcv3.CreditManagerData) CMCallData {
	return CMCallData{
		Addr: values.Addr,
		// Underlying: values.Underlying,
		// BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
		//
		MinDebt: (*core.BigInt)(values.MinDebt),
		MaxDebt: (*core.BigInt)(values.MaxDebt),
		//
		Adapters: values.Adapters,
		// Quotas:   values.Quotas,
	}
}
func getAccountDatav301(values dcv3.CreditAccountData) CreditAccountCallData {
	return CreditAccountCallData{
		CumulativeIndexAtOpen:   (*core.BigInt)(values.CumulativeIndexLastUpdate),
		CumulativeQuotaInterest: (*core.BigInt)(values.CumulativeQuotaInterest),
		CreditAccountInner: CreditAccountInner{
			Addr:           values.Addr,
			Borrower:       values.Borrower,
			CreditFacade:   values.CreditFacade,
			CreditManager:  values.CreditManager,
			Underlying:     values.Underlying,
			BorrowedAmount: (*core.BigInt)(values.Debt),
			Debt:           (*core.BigInt)(utils.BigIntAdd3(values.Debt, values.AccruedFees, values.AccruedInterest)), // DC_CHANGED

			//
			QuotaFeeCalc: QuotaFeeCalc{
				AccruedInterest: (*core.BigInt)(values.AccruedInterest),
				AccruedFees:     (*core.BigInt)(values.AccruedFees),
				Version:         core.NewVersion(int16(values.CfVersion.Int64())),
			},
			//
			TotalValue:   (*core.BigInt)(values.TotalValue),
			HealthFactor: (*core.BigInt)(values.HealthFactor),
			//
			// BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
			// Since:    values.Since,
			Balances: Convertv3ToBalance(values.Balances),
		},
	}
}

func Convertv301ToBalance(balances []dcv3.TokenBalance) (dcv2Balances []core.TokenBalanceCallData) {
	for ind, balance := range balances {
		dcv2Balances = append(dcv2Balances, core.TokenBalanceCallData{
			Token: balance.Token.Hex(),
			DBTokenBalance: core.DBTokenBalance{
				BI:          (*core.BigInt)(balance.Balance),
				IsForbidden: balance.IsForbidden, // is set on credit manager
				IsEnabled:   balance.IsEnabled,   // is used by credit account
				IsQuoted:    balance.IsQuoted,
				Quota:       (*core.BigInt)(balance.Quota),
				// QuotaRate:    balance.QuotaRate,
				QuotaIndexLU: (*core.BigInt)(balance.QuotaCumulativeIndexLU),
				Ind:          ind,
			},
		})
	}
	return
}
