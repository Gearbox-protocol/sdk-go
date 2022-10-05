package calc

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type account struct {
	CM              string
	BorrowedAmount  *big.Int
	CumulativeIndex *big.Int
	Balances        []core.BalanceType
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
func (a account) GetBalances() []core.BalanceType {
	return a.Balances
}

type store struct {
	Prices        map[int16]map[string]*core.BigInt
	LiqThresholds map[string]map[string]*big.Int `json:"LT"`
	Tokens        map[string]*schemas.Token
}

func (s store) GetPrices(token string, version int16, blockNums ...int64) *big.Int {
	return s.Prices[version][token].Convert()
}
func (s store) GetToken(token string) *schemas.Token {
	return s.Tokens[token]
}
func (s store) GetLiqThreshold(cm, token string) *big.Int {
	return s.LiqThresholds[cm][token]
}

type CalcFieldsParams struct {
	BlockNum        int64
	Version         int16
	CumIndexOfPool  *core.BigInt
	UnderlyingToken common.Address
	//
	////
	store
	Account account
}

func TestCalcFields(t *testing.T) {
	//
	input := CalcFieldsParams{}
	utils.ReadJsonAndSetInterface("../inputs/calc_account_fields.json", &input)
	// //

	calHF, calBorrowAmountPLusInterest, calTotalValue, calThresholdValue := Calculator{Store: input.store}.CalcAccountFields(
		input.Version,
		0,
		input.Account,
		input.CumIndexOfPool.Convert(),
		input.UnderlyingToken.Hex(),
		0,
	)
	if calHF.Cmp(utils.StringToInt("13225")) != 0 {
		t.Fatalf("calculated HF(%d) is wrong", calHF)
	}
	if calBorrowAmountPLusInterest.Cmp(utils.StringToInt("8319356395")) != 0 {
		t.Fatalf("calculated borrowedamount + interest(%d) is wrong", calBorrowAmountPLusInterest)
	}
	if calTotalValue.Cmp(utils.StringToInt("12944049438")) != 0 {
		t.Fatalf("calculated totalvalue(%d) is wrong", calTotalValue)
	}
	if calThresholdValue.Cmp(utils.StringToInt("11002442022")) != 0 {
		t.Fatalf("calculated thresholdvalue(%d) is wrong", calThresholdValue)
	}
}
