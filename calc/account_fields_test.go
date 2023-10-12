package calc

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

// ///////////
// ///////////
type account struct {
	CM              string
	BorrowedAmount  *big.Int
	CumulativeIndex *big.Int
	Balances        []dataCompressorv2.TokenBalance
	Version         int16
}

func (a account) GetCM() string {
	return a.CM
}

func (a account) GetBorrowedAmount() *big.Int {
	return a.BorrowedAmount
}

func (a account) GetCumulativeIndex() *big.Int {
	return a.CumulativeIndex
}
func (a account) GetBalances() core.DBBalanceFormat {
	return core.ConvertToDBBalanceFormat(dc.Convertv2ToBalance(a.Balances))
}

func (a account) GetQuotaCumInterestAndFees() (*big.Int, *big.Int) {
	return new(big.Int), new(big.Int)
}
func (a account) GetVersion() core.VersionType {
	return core.NewVersion(a.Version)
}

// ///////////
// ///////////
type store struct {
	Prices        map[core.VersionType]map[string]*core.BigInt
	LiqThresholds map[string]map[string]*big.Int `json:"LT"`
	Tokens        map[string]*schemas.Token
}

func (s store) GetPrices(token string, version core.VersionType, blockNums ...int64) *big.Int {
	return s.Prices[version][token].Convert()
}
func (s store) GetToken(token string) *schemas.Token {
	return s.Tokens[token]
}
func (s store) GetLiqThreshold(_ uint64, cm, token string) *big.Int {
	return s.LiqThresholds[cm][token]
}

// ///////////
// ///////////
type PoolDetails struct {
	CumIndexOfPool  *core.BigInt
	UnderlyingToken common.Address
}

func (details PoolDetails) GetPoolQuotaDetails() map[string]*schemas_v3.QuotaDetails {
	return nil
}
func (details PoolDetails) GetCumIndexNow() *big.Int {
	return details.CumIndexOfPool.Convert()
}
func (details PoolDetails) GetUnderlying() string {
	return details.UnderlyingToken.Hex()
}

type CalcFieldsParams struct {
	BlockNum int64

	FeeInterest uint16
	//
	////
	store
	PoolDetails PoolDetails
	Account     account
}

// for v1
func TestCalcFields(t *testing.T) {
	log.SetTestLogging(t)
	//
	input := CalcFieldsParams{}
	utils.ReadJsonAndSetInterface("../inputs/calc_account_fields.json", &input)
	// //

	calHF, calDebt, calTotalValue, calThresholdValue, _ := Calculator{Store: input.store}.CalcAccountFields(
		0,
		0,
		input.PoolDetails,
		input.Account,
		input.FeeInterest,
	)
	if calHF.Cmp(utils.StringToInt("13225")) != 0 {
		t.Fatalf("calculated HF(%d) is wrong", calHF)
	}
	if calDebt.Cmp(utils.StringToInt("8319356395")) != 0 {
		t.Fatalf("calculated borrowedamount + interest(%d) is wrong", calDebt)
	}
	if calTotalValue.Cmp(utils.StringToInt("12944049438")) != 0 {
		t.Fatalf("calculated totalvalue(%d) is wrong", calTotalValue)
	}
	if calThresholdValue.Cmp(utils.StringToInt("11002442022")) != 0 {
		t.Fatalf("calculated thresholdvalue(%d) is wrong", calThresholdValue)
	}
}

// for v2
func TestCalcFieldsWithFeeInterest(t *testing.T) {
	log.SetTestLogging(t)
	//
	input := CalcFieldsParams{}
	utils.ReadJsonAndSetInterface("../inputs/calc_account_fields_v2.json", &input)
	// //

	calHF, calDebt, calTotalValue, calThresholdValue, _ := Calculator{Store: input.store}.CalcAccountFields(
		0,
		0,
		//
		input.PoolDetails,
		//
		input.Account,
		input.FeeInterest,
	)
	if calHF.Cmp(utils.StringToInt("2725195")) != 0 {
		t.Fatalf("calculated HF(%d) is wrong", calHF)
	}
	if calDebt.Cmp(utils.StringToInt("5000008908519365762")) != 0 {
		t.Fatalf("calculated borrowedamount + interest(%d) is wrong", calDebt)
	}
	if calTotalValue.Cmp(utils.StringToInt("1651636475399519415042")) != 0 {
		t.Fatalf("calculated totalvalue(%d) is wrong", calTotalValue)
	}
	if calThresholdValue.Cmp(utils.StringToInt("1362600092204603517410")) != 0 {
		t.Fatalf("calculated thresholdvalue(%d) is wrong", calThresholdValue)
	}
}
