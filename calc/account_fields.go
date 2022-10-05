package calc

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type TokenDetailsForCalcI interface {
	GetToken(token string) *schemas.Token
	GetPrices(underlyingToken string, version int16, blockNums ...int64) *big.Int
	GetLiqThreshold(cm, token string) *big.Int
}

type AccountForCalcI interface {
	GetCM() string
	GetBalances() []core.BalanceType
	GetBorrowedAmount() *big.Int
	GetCumulativeIndex() *big.Int
}

type Calculator struct {
	Store TokenDetailsForCalcI
}

func (c Calculator) CalcAccountFields(version int16, blockNum int64,
	account AccountForCalcI, curCumIndex *big.Int, underlyingToken string, feeInterest uint16,
) (calHF, calBorrowWithInterest, calTotalValue, calThresholdValue *big.Int) {

	calThresholdValueInUSD := new(big.Int)
	calTotalValueInUSD := new(big.Int)
	// calculate the debt parameters
	for _, balance := range account.GetBalances() {
		if balance.IsAllowed && balance.IsEnabled && balance.BI.Convert().Cmp(big.NewInt(1)) > 0 {
			//
			tokenValueInUSD := c.convertToUSD(balance.BI.Convert(), balance.Token, version, blockNum)
			tokenThresholdValueInUSD := new(big.Int).Mul(tokenValueInUSD, c.Store.GetLiqThreshold(account.GetCM(), balance.Token))
			//
			calThresholdValueInUSD = new(big.Int).Add(calThresholdValueInUSD, tokenThresholdValueInUSD)
			calTotalValueInUSD = new(big.Int).Add(calTotalValueInUSD, tokenValueInUSD)
		}
	}
	calBorrowWithInterest = new(big.Int).Quo(
		new(big.Int).Mul(curCumIndex, account.GetBorrowedAmount()),
		account.GetCumulativeIndex())
	//
	calTotalValue = c.convertFromUSD(calTotalValueInUSD, underlyingToken, version, blockNum)
	calThresholdValue = utils.GetInt64(
		c.convertFromUSD(
			calThresholdValueInUSD,
			underlyingToken, version, blockNum,
		),
		4, // div by the per factor
	)
	//
	if version == 1 {
		calHF = new(big.Int).Quo(
			utils.GetInt64(calThresholdValue, -4),
			calBorrowWithInterest,
		)
	} else if version == 2 {
		interest := new(big.Int).Sub(calBorrowWithInterest, account.GetBorrowedAmount())
		fees := utils.PercentMul(interest, big.NewInt(int64(feeInterest)))
		calBorrowWithInterestAndFees := new(big.Int).Add(calBorrowWithInterest, fees)
		calHF = new(big.Int).Quo(
			utils.GetInt64(calThresholdValue, -4),
			calBorrowWithInterestAndFees,
		)
	}
	return
}

func (c Calculator) convertToUSD(amount *big.Int, token string, version int16, blockNum int64) *big.Int {
	tokenDecimals := c.Store.GetToken(token).Decimals
	tokenPrice := c.Store.GetPrices(token, version, blockNum)
	tokenValueInUSD := utils.GetInt64(new(big.Int).Mul(amount, tokenPrice), tokenDecimals)
	return tokenValueInUSD
}
func (c Calculator) convertFromUSD(amount *big.Int, underlyingToken string, version int16, blockNum int64) *big.Int {
	underlyingDecimals := c.Store.GetToken(underlyingToken).Decimals
	underlyingPrice := c.Store.GetPrices(underlyingToken, version, blockNum)
	value := new(big.Int).Quo(utils.GetInt64(amount, -1*underlyingDecimals), underlyingPrice)
	return value
}