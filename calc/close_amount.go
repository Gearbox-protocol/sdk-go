package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func CalCloseAmount(params *schemas.Parameters, version core.VersionType, totalValue *big.Int, closureStatus int, debtDetails *DebtDetails) (amountToPool, remainingFunds, profit, loss *big.Int) {
	if version.Eq(1) {
		return calCloseAmountV1(params, totalValue, schemas.IsStatusLiquidated(closureStatus), debtDetails)
	} else if version.Eq(2) || version.Eq(300) {
		amountToPool, remainingFunds, profit, loss = calCloseAmountV2(params, totalValue, closureStatus, debtDetails)
	}
	return
}

func calCloseAmountV1(params *schemas.Parameters, totalValue *big.Int, isLiquidated bool, debtDetails *DebtDetails) (amountToPool, remainingFunds, profit, loss *big.Int) {
	loss = big.NewInt(0)
	profit = big.NewInt(0)
	remainingFunds = new(big.Int)
	var totalFunds *big.Int
	if isLiquidated {
		totalFunds = utils.PercentMulByUInt16(totalValue, params.LiquidationDiscount)
	} else {
		totalFunds = totalValue
	}
	borrowedAmountWithInterest := debtDetails.BorrowedAmountWithInterest()
	// borrow amt is greater than total funds
	if totalFunds.Cmp(borrowedAmountWithInterest) < 0 {
		amountToPool = new(big.Int).Sub(totalFunds, big.NewInt(1))
		loss = new(big.Int).Sub(borrowedAmountWithInterest, amountToPool)
	} else {
		if isLiquidated {
			amountToPool = utils.PercentMulByUInt16(totalFunds, params.FeeLiquidation)
			amountToPool = new(big.Int).Add(borrowedAmountWithInterest, amountToPool)
		} else {
			interestAmt := debtDetails.Interest()
			fee := utils.PercentMulByUInt16(interestAmt, params.FeeInterest)
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

// https://github.com/Gearbox-protocol/core-v2/blob/da38b329f0c59e4a3dcedc993192bbc849d981f5/contracts/credit/CreditManager.sol#L1238
func calCloseAmountV2(params *schemas.Parameters, totalValue *big.Int, closureStatus int, debtDetails *DebtDetails) (amountToPool, remainingFunds, profit, loss *big.Int) {
	loss = big.NewInt(0)
	profit = big.NewInt(0)
	remainingFunds = new(big.Int)

	amountToPool = debtDetails.Total()

	if schemas.IsStatusLiquidated(closureStatus) {
		var totalFunds *big.Int
		switch closureStatus {
		case schemas.Liquidated:
			totalFunds = utils.PercentMulByUInt16(totalValue, params.LiquidationDiscount)
			liquidationFeeToPool := utils.PercentMulByUInt16(totalValue, params.FeeLiquidation)
			amountToPool = new(big.Int).Add(amountToPool, liquidationFeeToPool)
		case schemas.LiquidateExpired:
			totalFunds = utils.PercentMulByUInt16(totalValue, params.LiquidationDiscountExpired)
			liquidationFeeToPool := utils.PercentMulByUInt16(totalValue, params.FeeLiquidationExpired)
			amountToPool = new(big.Int).Add(amountToPool, liquidationFeeToPool)
		case schemas.LiquidatePaused:
			totalFunds = utils.PercentMulByUInt16(totalValue, params.EmergencyLiqDiscount)
			// here liquidationFee is calculated by multiple totalFunds not totalValue.
			// https://github.com/Gearbox-protocol/core-v2/blob/2f01dcaa2512a4f51157bacce45544c51e5033b3/contracts/credit/CreditFacade.sol#L545-L548
			liquidationFeeToPool := utils.PercentMulByUInt16(totalFunds, params.FeeLiquidation)
			amountToPool = new(big.Int).Add(amountToPool, liquidationFeeToPool)
		}
		//
		if totalFunds.Cmp(amountToPool) > 0 {
			remainingFunds = new(big.Int).Sub(totalFunds, new(big.Int).Add(amountToPool, big.NewInt(1)))
		} else {
			amountToPool = totalFunds
		}

		borrowedAmountWithInterest := debtDetails.BorrowedAmountWithInterest()
		if totalFunds.Cmp(borrowedAmountWithInterest) >= 0 {
			profit = new(big.Int).Sub(amountToPool, borrowedAmountWithInterest)
		} else {
			loss = new(big.Int).Sub(borrowedAmountWithInterest, amountToPool)
		}
	} else {
		profit = new(big.Int).Sub(amountToPool, debtDetails.BorrowedAmountWithInterest())
		// reminaingFunds is totalValue - debt
		remainingFunds = new(big.Int).Sub(totalValue, amountToPool)
	}
	return
}
