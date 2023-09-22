package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type PoolForCalcI interface {
	GetPoolQuotaDetails(underlyingToken string) map[string]*schemas_v3.QuotaDetails
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

func (c Calculator) CalcAccountFieldsv3(ts uint64, blockNum int64, poolCumIndexNow *big.Int, poolQuotaDetails PoolForCalcI, session AccountForCalcI, feeInterest uint16) (calHF, calBorrowWithInterestAndFees, calTotalValue, calThresholdValue, calBorrowWithInterest *big.Int) {
	version := core.NewVersion(3)
	_, accruedInterest, totalDebt := c.getDebt(ts, poolCumIndexNow, poolQuotaDetails, session, feeInterest)

	underlying := session.GetUnderlying()
	// quotaedTokens := poolQuotaDetails.GetPoolQuotaDetails(underlying)
	accountQuotas := session.GetQuotas()
	//
	totalValueInUSD, tvwValueInUSD := big.NewInt(0), big.NewInt(0)
	for token, balance := range session.GetBalances() {
		if balance.IsEnabled && balance.HasBalanceMoreThanOne() {
			//
			var quotaInUSD *big.Int
			if quota := accountQuotas[token]; quota != nil {
				quotaInUSD = c.convertToUSD(quota.Quota.Convert(), token, version, blockNum)
			} else {
				quotaInUSD = utils.GetExpInt(84)
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
	// calBorrowWithInterestAndFees := totalDebt
	calBorrowWithInterest = new(big.Int).Sub(session.GetBorrowedAmount(), accruedInterest)
	calHF = new(big.Int).Quo(utils.GetInt64(calThresholdValue, -4), totalDebt)

	return
}

func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) == -1 {
		return a
	} else {
		return b
	}
}

func (c Calculator) getDebt(ts uint64, poolCumIndexNow *big.Int, poolQuotaDetails PoolForCalcI, session AccountForCalcI, feeInterest uint16) (*big.Int, *big.Int, *big.Int) {
	borrowedAmount := session.GetBorrowedAmount()
	baseInterestSinceUpdate := GetbaseInterest(poolCumIndexNow, session)

	cumQuotaInterest, quotaFees := session.GetQuotaCumInterestAndFees()
	extraQuotaInterest := calcExtraQuotaInterest(ts, poolQuotaDetails, session)

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

	// totalDebt DebtWithInterstAndFees
	totalDebt := new(big.Int).Add(new(big.Int).Add(borrowedAmount, accruedInterest), accruedFees)

	return accruedFees, accruedInterest, totalDebt
}

func calcExtraQuotaInterest(ts uint64, poolQuotaDetails PoolForCalcI, session AccountForCalcI) *big.Int {
	accountQuotas := session.GetQuotas()
	poolQuotas := poolQuotaDetails.GetPoolQuotaDetails(session.GetUnderlying())
	totalQuotedInterest := big.NewInt(0)
	for _, quota := range accountQuotas {
		poolQuota := poolQuotas[quota.Token]
		interest := quota.CalcAccruedQuotaInterest(ts, poolQuota)

		//
		totalQuotedInterest = new(big.Int).Add(totalQuotedInterest, interest)
	}

	return totalQuotedInterest
}
