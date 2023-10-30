package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type PoolForCalcI interface {
	GetPoolQuotaDetails() map[string]*schemas_v3.QuotaDetails
	GetCumIndexNow() *big.Int
	GetUnderlying() string
}

func GetbaseInterest(poolCumIndexNow *big.Int, session AccountForCalcI) *big.Int {
	borrowedAmountWithInterest := new(big.Int).Quo(
		new(big.Int).Mul(poolCumIndexNow, session.GetBorrowedAmount()),
		session.GetCumulativeIndex(),
	)
	return new(big.Int).Sub(
		borrowedAmountWithInterest,
		session.GetBorrowedAmount(),
	)
}

// GENERIC_PARAMS
// - whenever there is updateQuota, cumulative quota intereest and fees are stored in creditManager and cumulativeQuotaIndex for account is updated.
// cumulative quota interest and quotafees increase on every updateQuota and decrase on decrease debt.

func (c Calculator) CalcAccountFieldsv3(version core.VersionType, ts uint64, blockNum int64, poolDetails PoolForCalcI, session AccountForCalcI, feeInterest uint16) (calHF, calTotalValue, calThresholdValue *big.Int, debtDetails *DebtDetails) {
	debtDetails = c.getDebtDetails(version, ts, blockNum, poolDetails, session, feeInterest)

	underlying := poolDetails.GetUnderlying()
	// quotaedTokens := poolQuotaDetails.GetPoolQuotaDetails(underlying)
	//
	totalValueInUSD, tvwValueInUSD := big.NewInt(0), big.NewInt(0)
	for token, balance := range session.GetBalances() {
		if balance.IsEnabled && balance.HasBalanceMoreThanOne() {
			//
			var quotaInUSD *big.Int
			if balance.IsQuoted {
				quota := new(core.BigInt)
				if balance.Quota != nil {
					quota = balance.Quota
				}
				quotaInUSD = c.convertToUSD(quota.Convert(), underlying, version, blockNum)
			} else {
				quotaInUSD = utils.GetExpInt(90)
			}
			//
			tokenValueInUSD := c.convertToUSD(balance.BI.Convert(), token, version, blockNum)
			tokenTwvValueInUSD := new(big.Int).Quo(
				new(big.Int).Mul(
					minBigInt(tokenValueInUSD, quotaInUSD),              // quotaed value
					c.Store.GetLiqThreshold(ts, session.GetCM(), token), // lt
				),
				utils.GetExpInt(4),
			)

			// sum
			totalValueInUSD = new(big.Int).Add(totalValueInUSD, tokenValueInUSD)
			tvwValueInUSD = new(big.Int).Add(tvwValueInUSD, tokenTwvValueInUSD)

		}
	}
	//
	calTotalValue = c.convertFromUSD(totalValueInUSD, underlying, version, blockNum)
	calThresholdValue = c.convertFromUSD(tvwValueInUSD, underlying, version, blockNum)

	calHF = new(big.Int).Quo(utils.GetInt64(calThresholdValue, -4), debtDetails.Total())
	return
}

func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) == -1 {
		return a
	} else {
		return b
	}
}

func (c Calculator) getDebtDetails(version core.VersionType, ts uint64, blockNum int64, poolDetails PoolForCalcI, session AccountForCalcI, feeInterest uint16) *DebtDetails {
	borrowedAmount := session.GetBorrowedAmount()
	baseInterestSinceUpdate := GetbaseInterest(poolDetails.GetCumIndexNow(), session)

	cumQuotaInterest, quotaFees := session.GetQuotaCumInterestAndFees()
	extraQuotaInterest := calcExtraQuotaInterest(version, ts, blockNum, poolDetails, session)

	// total interest
	totalNewInterest := new(big.Int).Add(baseInterestSinceUpdate, extraQuotaInterest)
	//
	accruedFees := func() *big.Int {
		feeForInterest := new(big.Int).Quo(
			new(big.Int).Mul(totalNewInterest, big.NewInt(int64(feeInterest))),
			utils.GetExpInt(4),
		)
		return new(big.Int).Add(feeForInterest, quotaFees)
	}()

	accruedInterest := new(big.Int).Add(cumQuotaInterest, totalNewInterest)

	return &DebtDetails{
		total:          utils.BigIntAdd3(accruedFees, accruedInterest, borrowedAmount),
		interest:       accruedInterest,
		borrowedAmount: borrowedAmount,
	}
}

func calcExtraQuotaInterest(version core.VersionType, ts uint64, blockNum int64, poolDetails PoolForCalcI, session AccountForCalcI) *big.Int {
	balances := session.GetBalances()
	poolQuotas := poolDetails.GetPoolQuotaDetails()
	totalQuotedInterest := big.NewInt(0)
	//
	for tokenAddr, balance := range balances {
		if balance.IsQuoted && balance.Quota.Convert().Cmp(new(big.Int)) > 0 {
			poolQuota := poolQuotas[tokenAddr]
			interest := schemas_v3.CalcAccruedQuotaInterest(ts, balance, poolQuota)
			totalQuotedInterest = new(big.Int).Add(totalQuotedInterest, interest)
		}
		//
	}

	return totalQuotedInterest
}
