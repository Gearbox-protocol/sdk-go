package calc

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type TokenDetailsForCalcI interface {
	GetToken(token string) *schemas.Token
	GetPriceOnBlock(cm, underlyingToken string, version core.VersionType, blockNums ...int64) *big.Int
	GetLiqThreshold(ts uint64, cm, token string) *big.Int
}

type DebtDetails struct {
	total          *big.Int
	interest       *big.Int
	borrowedAmount *big.Int
}

func (d DebtDetails) String() string {
	return fmt.Sprintf("total:%s, interest:%s, borrowedAmount:%s", d.total, d.interest, d.borrowedAmount)
}

func NewDebtDetails(total, interest, borrowedAmount *big.Int) *DebtDetails {
	return &DebtDetails{
		total:          total,
		interest:       interest,
		borrowedAmount: borrowedAmount,
	}
}

func (d DebtDetails) Total() *big.Int {
	return d.total
}
func (d DebtDetails) Interest() *big.Int {
	return d.interest
}
func (d DebtDetails) BorrowedAmount() *big.Int {
	return d.borrowedAmount
}
func (d DebtDetails) BorrowedAmountWithInterest() *big.Int {
	return new(big.Int).Add(d.borrowedAmount, d.interest)
}

type AccountForCalcI interface {
	GetAddr() string
	GetCM() string
	GetBalances() core.DBBalanceFormat
	GetBorrowedAmount() *big.Int
	GetCumulativeIndex() *big.Int
	GetVersion() core.VersionType
	GetQuotaCumInterestAndFees() (*big.Int, *big.Int)
}

type Calculator struct {
	Store TokenDetailsForCalcI
}

func (c Calculator) CalcAccountFields(ts uint64, blockNum int64,
	poolDetails PoolForCalcI,
	account AccountForCalcI, feeInterest uint16, failure bool,
) (calHF, calTotalValue, calThresholdValue *big.Int, debtDetails *DebtDetails, profile string) {
	version := account.GetVersion()
	if version.Eq(300) {
		return c.CalcAccountFieldsv3(version, ts, blockNum, poolDetails, account, feeInterest, failure)
	}

	// logic for v1 and v2
	underlyingToken := poolDetails.GetUnderlying()
	//
	calThresholdValueInUSD := new(big.Int)
	calTotalValueInUSD := new(big.Int)
	// calculate the debt parameters
	for token, balance := range account.GetBalances() {
		if balance.IsEnabled && balance.HasBalanceMoreThanOne() {
			//
			tokenValueInUSD := c.convertToUSD(account.GetCM(), balance.BI.Convert(), token, version, blockNum)
			tokenThresholdValueInUSD := new(big.Int).Mul(tokenValueInUSD, c.Store.GetLiqThreshold(ts, account.GetCM(), token))
			//
			calThresholdValueInUSD = new(big.Int).Add(calThresholdValueInUSD, tokenThresholdValueInUSD)
			calTotalValueInUSD = new(big.Int).Add(calTotalValueInUSD, tokenValueInUSD)
			profile += fmt.Sprintf("token:%s price: %s, tv:%s tvw: %s", token, c.Store.GetPriceOnBlock(account.GetCM(), token, version, blockNum), tokenValueInUSD, tokenThresholdValueInUSD)
		}
	}

	calBorrowWithInterest := new(big.Int).Quo(
		new(big.Int).Mul(poolDetails.GetCumIndexNow(), account.GetBorrowedAmount()),
		account.GetCumulativeIndex())
	debtDetails = &DebtDetails{
		borrowedAmount: account.GetBorrowedAmount(),
		total:          calBorrowWithInterest,
		interest:       new(big.Int).Sub(calBorrowWithInterest, account.GetBorrowedAmount()),
	}
	//
	calTotalValue = c.convertFromUSD(account.GetCM(), calTotalValueInUSD, underlyingToken, version, blockNum)
	calThresholdValue = utils.GetInt64(
		c.convertFromUSD(
			account.GetCM(),
			calThresholdValueInUSD,
			underlyingToken, version, blockNum,
		),
		4, // div by the per factor
	)
	//
	if version.Eq(1) {
		calHF = new(big.Int).Quo(
			utils.GetInt64(calThresholdValue, -4),
			debtDetails.Total(),
		)
	} else if version.Eq(2) {
		//https://github.com/Gearbox-protocol/core-v2/blob/da38b329f0c59e4a3dcedc993192bbc849d981f5/contracts/credit/CreditFacade.sol#L1206
		fees := utils.PercentMul(debtDetails.Interest(), big.NewInt(int64(feeInterest)))
		debtDetails.total = new(big.Int).Add(debtDetails.total, fees)
		calHF = new(big.Int).Quo(
			utils.GetInt64(calThresholdValue, -4),
			debtDetails.Total(),
		)
	}
	return
}

func (c Calculator) convertToUSD(cm string, amount *big.Int, token string, version core.VersionType, blockNum int64) *big.Int {
	tokenDecimals := c.Store.GetToken(token).Decimals
	tokenPrice := c.Store.GetPriceOnBlock(cm, token, version, blockNum)
	tokenValueInUSD := utils.GetInt64(new(big.Int).Mul(amount, tokenPrice), tokenDecimals)
	return tokenValueInUSD
}
func (c Calculator) convertFromUSD(cm string, amount *big.Int, underlyingToken string, version core.VersionType, blockNum int64) *big.Int {
	underlyingDecimals := c.Store.GetToken(underlyingToken).Decimals
	underlyingPrice := c.Store.GetPriceOnBlock(cm, underlyingToken, version, blockNum)
	value := new(big.Int).Quo(utils.GetInt64(amount, -1*underlyingDecimals), underlyingPrice)
	return value
}
