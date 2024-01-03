package utils

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"golang.org/x/exp/constraints"
)

// maths
func GetExpFloat(decimals int8) *big.Float {
	if decimals < 0 {
		panic(fmt.Sprintf("GetExpFloat received pow:%d", decimals))
	}
	bigIntDecimal := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(decimals)), big.NewInt(0))
	return new(big.Float).SetInt(bigIntDecimal)
}

func GetExpInt(decimals int8) *big.Int {
	if decimals < 0 {
		panic(fmt.Sprintf("GetInt received pow:%d", decimals))
	}
	return new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(decimals)), big.NewInt(0))
}

func IntToFloat(amt *big.Int) *big.Float {
	return new(big.Float).SetInt(amt)
}

// 100 percent is 10000
func PercentMul(a, b *big.Int) *big.Int {
	ans := new(big.Int).Mul(a, b)
	return new(big.Int).Quo(ans, big.NewInt(10000))
}
func PercentMulByUInt16(a *big.Int, percent uint16) *big.Int {
	return PercentMul(a, big.NewInt(int64(percent)))
}

func GetFloat64Decimal(_num interface{}, decimals int8) float64 {
	type convertI interface {
		Convert() *big.Int
	}
	var bigInt *big.Int
	switch v := _num.(type) {
	case *big.Int:
		bigInt = v
	default:
		coreBigInt, ok := _num.(convertI)
		if !ok {
			log.Fatal("GetFloat64Decimal received unknown type")
		}
		bigInt = coreBigInt.Convert()
	}
	floatBorrowedAmount, _ := GetFloat64(bigInt, decimals).Float64()
	return floatBorrowedAmount
}

func GetInt64(num *big.Int, decimals int8) *big.Int {
	if decimals > 0 {
		return new(big.Int).Quo(
			num,
			GetExpInt(decimals))
	} else {
		return new(big.Int).Mul(
			num,
			GetExpInt(-decimals))
	}
}

func StringToInt(v string) *big.Int {
	value, ok := new(big.Int).SetString(v, 10)
	if !ok {
		panic("Parsing string to big.int failed")
	}
	return value
}

func GetFloat64(num *big.Int, decimals int8) *big.Float {
	if decimals > 0 {
		return new(big.Float).Quo(
			IntToFloat(num),
			GetExpFloat(decimals))
	} else {
		return new(big.Float).Mul(
			IntToFloat(num),
			GetExpFloat(-decimals))
	}
}

func AlmostSameBigInt(a, b *big.Int, noOFZeroIndiff int8) bool {
	// diff should be less than 100
	return new(big.Int).Sub(a, b).CmpAbs(GetExpInt(noOFZeroIndiff)) <= 0
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// others
// number of decimal digits to compare
func GetPrecision(symbol string) int8 {
	switch symbol {
	case "USDC":
		return 0
	case "DAI":
		return 0
	case "WBTC":
		return 5
	case "LINK":
		return 2
	case "SNX":
		return 3
	}
	return 0
}

func absInt64(a int64) int64 {
	if a > 0 {
		return a
	}
	return -1 * a
}

func absFloat64(a float64) float64 {
	if a > 0 {
		return a
	}
	return -1 * a
}

func NotNilBigInt(a *big.Int) *big.Int {
	if a == nil {
		return big.NewInt(0)
	}
	return a
}

func IntDiffMoreThanFraction(oldValue, newValue int64, diff float64) bool {
	return FloatDiffMoreThanFraction(float64(oldValue), float64(newValue), diff)
}
func FloatDiffMoreThanFraction(oldValue, newValue float64, diff float64) bool {
	return absFloat64(float64(newValue-oldValue))/float64(oldValue) > diff
}

func ifZeroReturnOneBigInt(a *big.Float) *big.Float {
	if a.Cmp(new(big.Float)) == 0 {
		// log.InfoStackN(4, c"var is zero", a)
		return big.NewFloat(1)
	} else {
		return a
	}
}

func DiffMoreThanFraction(oldValue, newValue *big.Int, diff *big.Float) bool {
	oldFloat := new(big.Float).SetInt(oldValue)
	newFloat := new(big.Float).SetInt(newValue)
	value := new(big.Float).Quo(new(big.Float).Sub(newFloat, oldFloat), ifZeroReturnOneBigInt(oldFloat))
	return new(big.Float).Abs(value).Cmp(diff) > 0
}

func BytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}

func BigIntAdd3(a, b, c *big.Int) *big.Int {
	return new(big.Int).Add(a, new(big.Int).Add(b, c))
}
