package calc

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func TestCalcCloseAmount(t *testing.T) {
	amountToPool, remainingFunds, profit, loss := CalCloseAmount(
		&schemas.Parameters{
			FeeInterest:         100,
			FeeLiquidation:      200,
			LiquidationDiscount: 9500,
		},
		2,
		utils.StringToInt("2003445408514560318103"),
		schemas.Liquidated,
		utils.StringToInt("2000036753743938235"),
		utils.StringToInt("1999999990917566710"),
	)
	expectedAmounToPool := utils.StringToInt("42068945291663408312")
	expectedRemainingFunds := utils.StringToInt("1861204192797168893884")
	expectedProfit := utils.StringToInt("40068908537919470077")

	if expectedAmounToPool.Cmp(amountToPool) != 0 {
		t.Fatal("amountToPool is different: ", amountToPool)
	}
	if expectedRemainingFunds.Cmp(remainingFunds) != 0 {
		t.Fatal("remainingFunds is different: ", remainingFunds)
	}
	if expectedProfit.Cmp(profit) != 0 {
		t.Fatal("profit is different: ", profit)
	}
	if loss.Int64() != 0 {
		t.Fatal("loss is different: ", loss)
	}
}
