package calc

import "math/big"

type CollateralDebtData struct {
	BorrowedAmount *big.Int
	LastCumIndex   *big.Int
}

// GENERIC_PARAMS
func calcDebtAndCollateral(poolCumIndexNow *big.Int, task string, supportsQuotas bool) {
	if supportsQuotas {

	}
}
