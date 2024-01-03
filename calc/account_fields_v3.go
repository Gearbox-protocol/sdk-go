package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
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
		getCumIndexOfAccount(session),
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
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalf("err: %s blockNum:%d ts:%d", err, blockNum, ts)
		}
	}()
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
			tokenTwvValueInUSD := minBigInt(
				new(big.Int).Quo(
					new(big.Int).Mul(
						tokenValueInUSD,
						c.Store.GetLiqThreshold(ts, session.GetCM(), token), // lt
					),
					utils.GetExpInt(4)),
				quotaInUSD,
			) // quoted value

			// sum
			totalValueInUSD = new(big.Int).Add(totalValueInUSD, tokenValueInUSD)
			tvwValueInUSD = new(big.Int).Add(tvwValueInUSD, tokenTwvValueInUSD)

		}
	}
	//
	calTotalValue = c.convertFromUSD(totalValueInUSD, underlying, version, blockNum)
	calThresholdValue = c.convertFromUSD(tvwValueInUSD, underlying, version, blockNum)
	if debtDetails.borrowedAmount.Cmp(big.NewInt(0)) == 0 {
		calHF = big.NewInt(65535)
	} else {
		calHF = new(big.Int).Quo(utils.GetInt64(calThresholdValue, -4), debtDetails.Total())
	}
	return
}

func minBigInt(a, b *big.Int) *big.Int {
	if a.Cmp(b) == -1 {
		return a
	} else {
		return b
	}
}

func getCumIndexOfAccount(session AccountForCalcI) *big.Int {
	index := session.GetCumulativeIndex()
	if index.Cmp(big.NewInt(0)) == 0 {
		return core.RAY
	}
	return index
}
func (c Calculator) getDebtDetails(version core.VersionType, ts uint64, blockNum int64, poolDetails PoolForCalcI, session AccountForCalcI, feeInterest uint16) *DebtDetails {
	defer func() {
		err := recover()
		if err != nil {
			log.Info(session.GetBorrowedAmount(), session.GetAddr(), getCumIndexOfAccount(session), poolDetails.GetCumIndexNow())
			log.Info(poolDetails.GetPoolQuotaDetails(), utils.ToJson(session.GetBalances()))
			log.Fatalf("err: %s blockNum:%d ts:%d", err, blockNum, ts)
		}
	}()
	borrowedAmount := session.GetBorrowedAmount()
	baseInterestSinceUpdate := GetbaseInterest(poolDetails.GetCumIndexNow(), session)

	cumQuotaInterest, quotaFees := session.GetQuotaCumInterestAndFees()
	extraQuotaInterest := calcExtraQuotaInterest(version, ts, blockNum, poolDetails, session)

	// total interest = base interest + extra quota interest + cumQuotaInterest
	accruedInterest := new(big.Int).Add(baseInterestSinceUpdate, extraQuotaInterest)
	accruedInterest = new(big.Int).Add(cumQuotaInterest, accruedInterest)
	//
	accruedFees := func() *big.Int {
		feeForInterest := new(big.Int).Quo(
			new(big.Int).Mul(accruedInterest, big.NewInt(int64(feeInterest))),
			utils.GetExpInt(4),
		)
		return new(big.Int).Add(feeForInterest, quotaFees)
	}()

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
