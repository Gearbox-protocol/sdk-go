package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type BalanceType struct {
	IsAllowed bool     `json:"isAllowed"`
	IsEnabled bool     `json:"isEnabled"` // based on mask
	BI        *big.Int `json:"BI"`
	Ind       int      `json:"-"`
	// not used in liquidator
}

func (b BalanceType) HasBalanceMoreThanOne() bool {
	return b.BI != nil && b.BI.Cmp(big.NewInt(1)) > 0
}

func ConvertToBalanceType(dcv2Balances []dataCompressorv2.TokenBalance) map[string]BalanceType {
	m := map[string]BalanceType{}
	for ind, entry := range dcv2Balances {
		if entry.IsAllowed && entry.IsEnabled {
			m[entry.Token.Hex()] = BalanceType{
				IsAllowed: entry.IsAllowed,
				IsEnabled: entry.IsEnabled,
				BI:        entry.Balance,
				Ind:       ind,
			}
		}
	}
	return m
}

type CoreIntBalance struct {
	IsAllowed bool    `json:"isAllowed"`
	IsEnabled bool    `json:"isEnabled"`
	BI        *BigInt `json:"BI"`
	F         float64 `json:"F"`
	Ind       int     `json:"ind"`
}

func (b CoreIntBalance) HasBalanceMoreThanOne() bool {
	return b.BI != nil && b.BI.Convert().Cmp(big.NewInt(1)) > 0
}

type DBBalanceFormat map[string]CoreIntBalance

func (j DBBalanceFormat) ToBalanceType() map[string]BalanceType {
	m := map[string]BalanceType{}
	for token, bal := range j {
		m[token] = BalanceType{
			Ind:       bal.Ind,
			BI:        bal.BI.Convert(),
			IsAllowed: bal.IsAllowed,
			IsEnabled: bal.IsEnabled,
		}
	}
	return m
}
func (j DBBalanceFormat) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *DBBalanceFormat) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte:
		out := DBBalanceFormat{}
		err := json.Unmarshal(value.([]byte), &out)
		if err != nil {
			return err
		}
		*z = out
	default:
		return fmt.Errorf("could not scan type %T", t)
	}
	return nil
}

func (j DBBalanceFormat) ValueInUnderlying(underlyingToken string, uDecimals int8, prices JsonFloatMap) *big.Int {
	var total float64
	priceOfUnderlying := prices[underlyingToken]
	for token, bal := range j {
		tokenPrice := prices[token]
		value := (bal.F * tokenPrice) / priceOfUnderlying
		total += value
	}
	valueInFloat := new(big.Float).Mul(big.NewFloat(total), utils.GetExpFloat(uDecimals))
	remainingFunds, _ := valueInFloat.Int(nil)
	return remainingFunds
}
