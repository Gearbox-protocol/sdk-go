package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

type TokenBalanceCallData struct {
	Token string `json:"token"`
	DBTokenBalance
}

// there is balanceType without float decimal value and
// CoreIntBalance with float decimal value
// DBFOrmatBalance is map[string]CoreIntBalance
// there is method on DbFormatBalance ToBalanceType() map[string]BalanceType
// type BalanceType struct {
// 	// not important field is not used for any calculation
// 	IsForbidden bool    `json:"isForbidden"`
// 	IsEnabled   bool    `json:"isEnabled"` // based on mask
// 	BI          *BigInt `json:"BI"`
// 	Ind         int     `json:"-"`
// 	// not used in liquidator
// }

// func (b BalanceType) HasBalanceMoreThanOne() bool {
// 	return b.BI != nil && b.BI.Convert().Cmp(big.NewInt(1)) > 0
// }

// filters isEnabled
func ConvertToDBBalanceFormat(dcv2Balances []TokenBalanceCallData) DBBalanceFormat {
	m := DBBalanceFormat{}
	for _, entry := range dcv2Balances {
		if entry.IsEnabled {
			m[entry.Token] = entry.DBTokenBalance
		}
	}
	return m
}

type DBTokenBalance struct {
	// not important field is not used for any calculation
	BI *BigInt `json:"BI"`
	// is forbidden on credit manager
	IsForbidden bool `json:"isForbidden"`
	// is enabled on account
	IsEnabled bool    `json:"isEnabled"`
	Ind       int     `json:"ind"`
	IsQuoted  bool    `json:"isQuoted,omitempty"`
	Quota     *BigInt `json:"quota,omitempty"`
	// TODO add last updated index
	QuotaIndexLU *BigInt `json:"quotaIndexLU,omitempty"`
	QuotaRate    uint16  `json:"quotaRate,omitempty"`
	F            float64 `json:"F"`
	//
} // @name DBTokenBalance

func (b DBTokenBalance) HasBalanceMoreThanOne() bool {
	return b.BI != nil && b.BI.Convert().Cmp(big.NewInt(1)) > 0
}

// uint128(uint256(quoted) * (cumulativeIndexNow - cumulativeIndexLU) / RAY);

type DBBalanceFormat map[string]DBTokenBalance // @name DBBalanceFormat

// doesn't filter on isEnabled
//
//	func (j DBBalanceFormat) ToBalanceType() map[string]BalanceType {
//		m := map[string]BalanceType{}
//		for token, bal := range j {
//			m[token] = BalanceType{
//				Ind:         bal.Ind,
//				BI:          bal.BI,
//				IsForbidden: bal.IsForbidden,
//				IsEnabled:   bal.IsEnabled,
//			}
//		}
//		return m
//	}

func (j DBBalanceFormat) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *DBBalanceFormat) Scan(value interface{}) error {
	switch t := value.(type) {
	case []byte, string:
		dataBytes, ok := t.([]byte)
		if !ok {
			dataBytes = []byte((t.(string)))
		}
		out := DBBalanceFormat{}
		err := json.Unmarshal(dataBytes, &out)
		if err != nil {
			return err
		}
		*z = out
	default:
		return fmt.Errorf("could not scan type %T", t)
	}
	return nil
}

// deprecated
// func (j DBBalanceFormat) ValueInUnderlying(underlyingToken string, uDecimals int8, prices JsonFloatMap) *big.Int {
// 	var total float64
// 	priceOfUnderlying := prices[underlyingToken]
// 	for token, bal := range j {
// 		tokenPrice := prices[token]
// 		value := (bal.F * tokenPrice) / priceOfUnderlying
// 		total += value
// 	}
// 	valueInFloat := new(big.Float).Mul(big.NewFloat(total), utils.GetExpFloat(uDecimals))
// 	remainingFunds, _ := valueInFloat.Int(nil)
// 	return remainingFunds
// }
