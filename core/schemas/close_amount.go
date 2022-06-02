package schemas

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/utils"
)

func CalCloseAmount(params *Parameters, version int16, totalValue *big.Int, isLiquidated bool, borrowedAmountWithInterest, borrowedAmount *big.Int) (amountToPool, remainingFunds, profit, loss *big.Int) {
	switch version {
	case 1:
		return calCloseAmountV1(params, totalValue, isLiquidated, borrowedAmountWithInterest, borrowedAmount)
	case 2:
		amountToPool, remainingFunds, profit, loss = calCloseAmountV2(params, totalValue, isLiquidated, borrowedAmountWithInterest, borrowedAmount)
	}
	return
}
func calCloseAmountV1(params *Parameters, totalValue *big.Int, isLiquidated bool, borrowedAmountWithInterest, borrowedAmount *big.Int) (amountToPool, remainingFunds, profit, loss *big.Int) {
	loss = big.NewInt(0)
	profit = big.NewInt(0)
	remainingFunds = new(big.Int)
	var totalFunds *big.Int
	if isLiquidated {
		totalFunds = utils.PercentMul(totalValue, params.LiquidationDiscount.Convert())
	} else {
		totalFunds = totalValue
	}
	// borrow amt is greater than total funds
	if totalFunds.Cmp(borrowedAmountWithInterest) < 0 {
		amountToPool = new(big.Int).Sub(totalFunds, big.NewInt(1))
		loss = new(big.Int).Sub(borrowedAmountWithInterest, amountToPool)
	} else {
		if isLiquidated {
			amountToPool = utils.PercentMul(totalFunds, params.FeeLiquidation.Convert())
			amountToPool = new(big.Int).Add(borrowedAmountWithInterest, amountToPool)
		} else {
			interestAmt := new(big.Int).Sub(borrowedAmountWithInterest, borrowedAmount)
			fee := utils.PercentMul(interestAmt, params.FeeInterest.Convert())
			amountToPool = new(big.Int).Add(borrowedAmountWithInterest, fee)
		}

		if totalFunds.Cmp(amountToPool) <= 0 {
			amountToPool = new(big.Int).Sub(totalFunds, big.NewInt(1))
		} else {
			remainingFunds = new(big.Int).Sub(totalFunds, amountToPool)
			// remainingFunds = new(big.Int).Sub(new(big.Int).Sub(totalFunds, amountToPool), big.NewInt(1))
		}
		profit = new(big.Int).Sub(amountToPool, borrowedAmountWithInterest)
	}
	return
}

func calCloseAmountV2(params *Parameters, totalValue *big.Int, isLiquidated bool, borrowedAmountWithInterest, borrowedAmount *big.Int) (amountToPool, remainingFunds, profit, loss *big.Int) {
	loss = big.NewInt(0)
	profit = big.NewInt(0)
	remainingFunds = new(big.Int)

	amountToPool = utils.PercentMul(
		new(big.Int).Sub(borrowedAmountWithInterest, borrowedAmount),
		params.FeeInterest.Convert(),
	)
	amountToPool = new(big.Int).Add(amountToPool, borrowedAmountWithInterest)

	if isLiquidated {
		totalFunds := utils.PercentMul(totalValue, params.LiquidationDiscount.Convert())
		liquidationFeeToPool := utils.PercentMul(totalValue, params.FeeLiquidation.Convert())
		amountToPool = new(big.Int).Add(amountToPool, liquidationFeeToPool)
		if totalFunds.Cmp(amountToPool) > 0 {
			remainingFunds = new(big.Int).Sub(totalFunds, new(big.Int).Add(amountToPool, big.NewInt(1)))
		} else {
			amountToPool = totalFunds
		}

		if totalFunds.Cmp(borrowedAmountWithInterest) >= 0 {
			profit = new(big.Int).Sub(amountToPool, borrowedAmountWithInterest)
		} else {
			loss = new(big.Int).Sub(borrowedAmountWithInterest, amountToPool)
		}
	} else {
		profit = new(big.Int).Sub(amountToPool, borrowedAmountWithInterest)
	}
	return
}
