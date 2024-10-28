package dc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccountCompressor"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolCompressor"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// quotaValues poolCompressor.PoolQuotaKeeperState
// quotaValues can be another param for getting quotas
func getPoolDatav310(pool poolCompressor.PoolState) PoolCallData {
	return PoolCallData{
		Addr:               pool.BaseParams.Addr,
		Underlying:         pool.Underlying,
		DieselToken:        pool.BaseParams.Addr,
		AvailableLiquidity: (*core.BigInt)(pool.AvailableLiquidity),
		ExpectedLiquidity:  (*core.BigInt)(pool.ExpectedLiquidity),
		TotalBorrowed:      (*core.BigInt)(pool.TotalBorrowed),
		// : pool.//,
		// CreditManagerDebtParams: pool.CreditManagerDebtParams,
		TotalAssets:      (*core.BigInt)(pool.TotalAssets),
		TotalSupply:      (*core.BigInt)(pool.TotalSupply),
		SupplyRate:       (*core.BigInt)(pool.SupplyRate),
		BaseInterestRate: (*core.BigInt)(pool.BaseInterestRate),
		DieselRateRAY:    (*core.BigInt)(pool.DieselRate),
		WithdrawFee:      (*core.BigInt)(pool.WithdrawFee),
		// assuming cumulativeIndexRay is depcreated
		// CumulativeIndexRAY is same as BaseInterestIndexLU
		CumulativeIndexRAY: (*core.BigInt)(pool.BaseInterestIndexLU),
		// : values.//,
		Version: (*core.BigInt)(big.NewInt(300)),
		// Quotas:  quotaRates(quotaValues),
	}
}

func getCMDatav310(values poolCompressor.CreditManagerData) CMCallData {
	return CMCallData{
		Addr:       values.CreditManager.BaseParams.Addr,
		Underlying: values.CreditManager.Underlying,
		// BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
		//
		MinDebt: (*core.BigInt)(values.CreditFacade.MinDebt),
		MaxDebt: (*core.BigInt)(values.CreditFacade.MaxDebt),
		//
		Adapters: func() (ans []dcv3.ContractAdapter) {
			for _, entry := range values.Adapters {
				ans = append(ans, dcv3.ContractAdapter{
					TargetContract: entry.TargetContract,
					Adapter:        entry.BaseParams.Addr,
				})
			}
			return
		}(),
		// Quotas:   values.Quotas,
	}
}

func quotaRates(quotaValues poolCompressor.PoolQuotaKeeperState) (ans []dcv3.QuotaInfo) {
	for _, quota := range quotaValues.Quotas {
		ans = append(ans, dcv3.QuotaInfo{
			Token:            quota.Token,
			Rate:             quota.Rate,
			QuotaIncreaseFee: quota.QuotaIncreaseFee,
			TotalQuoted:      quota.TotalQuoted,
			Limit:            quota.Limit,
			IsActive:         quota.IsActive,
		})
	}
	return
}

type CreditAccountv310 struct {
	creditAccountCompressor.CreditAccountData
	CumulativeIndexLastUpdate *big.Int
	CumulativeQuotaInterest   *big.Int
	QuotaCumIndexMap          map[string]*big.Int
}

func GetCreditAccountv310(values CreditAccountv310) CreditAccountCallData {
	if values.CumulativeIndexLastUpdate == nil || values.CumulativeQuotaInterest == nil {
		log.Fatal("other details not set.")
	}
	data := CreditAccountCallData{
		CumulativeIndexAtOpen:   (*core.BigInt)(values.CumulativeIndexLastUpdate),
		CumulativeQuotaInterest: (*core.BigInt)(values.CumulativeQuotaInterest),
		CreditAccountInner:      GetCreditAccountv310Inner(values.CreditAccountData),
	}
	data.Balances = convertv310ToBalance(values.Tokens, values)
	return data
}
func GetCreditAccountv310Inner(values creditAccountCompressor.CreditAccountData) CreditAccountInner {
	return CreditAccountInner{
		Addr:           values.CreditAccount,
		CreditFacade:   values.CreditFacade,
		Borrower:       values.Owner,
		CreditManager:  values.CreditManager,
		Underlying:     values.Underlying,
		BorrowedAmount: (*core.BigInt)(values.Debt),
		Debt:           (*core.BigInt)(utils.BigIntAdd3(values.Debt, values.AccruedFees, values.AccruedInterest)), // DC_CHANGED

		//
		QuotaFeeCalc: QuotaFeeCalc{
			AccruedInterest: (*core.BigInt)(values.AccruedInterest),
			AccruedFees:     (*core.BigInt)(values.AccruedFees),
			Version:         core.NewVersion(300),
		},
		//
		TotalValue:   (*core.BigInt)(values.TotalValue),
		HealthFactor: (*core.BigInt)(big.NewInt(int64(values.HealthFactor))),
		//
		// BaseBorrowRate: (*core.BigInt)(values.BaseBorrowRate),
		// Since:    values.Since,
		Balances: Convertv310BalWithoutQuotaIndex(values.Tokens),
	}
}

func convertv310ToBalance(balances []creditAccountCompressor.TokenInfo, values CreditAccountv310) (dcv2Balances []core.TokenBalanceCallData) {
	ans := Convertv310BalWithoutQuotaIndex(balances)
	for ind, entry := range ans {
		entry.QuotaIndexLU = (*core.BigInt)(values.QuotaCumIndexMap[entry.Token])
		ans[ind] = entry
	}
	return ans
}
func Convertv310BalWithoutQuotaIndex(balances []creditAccountCompressor.TokenInfo) (dcv2Balances []core.TokenBalanceCallData) {
	for ind, balance := range balances {
		dcv2Balances = append(dcv2Balances, core.TokenBalanceCallData{
			Token: balance.Token.Hex(),
			DBTokenBalance: core.DBTokenBalance{
				BI:          (*core.BigInt)(balance.Balance),
				IsForbidden: false, // is set on credit manager
				IsEnabled:   true,  // is used by credit account
				IsQuoted:    balance.Quota != nil && balance.Quota.Cmp(new(big.Int)) > 0,
				Quota:       (*core.BigInt)(balance.Quota),
				// QuotaRate:    balance.QuotaRate,
				Ind: ind,
			},
		})
	}
	return
}
