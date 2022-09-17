package utils

import (
	"math"
	"math/big"
)

// https://docs.uniswap.org/sdk/guides/fetching-prices#understanding-sqrtprice
// sqrt[price*2^192/10^decimalsT0]
// price is t1/t0
func GetSqrtX96FromPrice(price *big.Int, decimals int8) *big.Int {
	normalizeFactor := new(big.Int).Exp(big.NewInt(2), big.NewInt(96*2), nil)
	nom := new(big.Int).Mul(price, normalizeFactor)
	return new(big.Int).Sqrt(GetInt64(nom, decimals))
}

// https://docs.uniswap.org/sdk/guides/fetching-prices#understanding-sqrtprice
// (tick^2*decimalsT0/2^192)
func GetSqrtX96ToPrice(slot0 *big.Int, decimals int8) *big.Int {
	normalizeFactor := new(big.Int).Exp(big.NewInt(2), big.NewInt(96*2), nil)
	sqSlot0 := new(big.Int).Mul(slot0, slot0)
	return new(big.Int).Quo(GetInt64(sqSlot0, -1*decimals), normalizeFactor)
}

/////////////
/////////////
/////////////

// sqrt(1.001^tick)*2**96
// tick to sqrtx96
//
// sqrtx96 to tick
// log1.0001(sqrtx96^2/2**192)
func GetTickFromSqrtX96(sqrtX96 *big.Int) (tick *big.Int) {
	x96Float, _ := new(big.Float).SetPrec(96).SetInt(sqrtX96).Float64()
	nom := math.Log2(x96Float)*2 - 192
	//
	bigNom := new(big.Float).SetPrec(96).SetFloat64(nom)
	bigDenom := new(big.Float).SetPrec(96).SetFloat64(math.Log2(1.0001))
	tick, _ = new(big.Float).SetPrec(96).Quo(bigNom, bigDenom).Int(nil)
	//
	// tick, _ = new(big.Int)(nom / math.Log2(1.0001)).Int(nil)
	return
}

func powFloat(a *big.Int) *big.Float {
	f := big.NewFloat(1.0001).SetPrec(96)
	ans := big.NewFloat(1).SetPrec(96)
	absA := new(big.Int).Abs(a)
	for i := 0; i < absA.BitLen(); i++ {
		if absA.Bit(i) == 1 {
			ans = new(big.Float).Mul(f, ans)
		}
		f = new(big.Float).SetPrec(96).Mul(f, f)
	}
	if absA == a {
		return ans
	}
	return new(big.Float).Quo(big.NewFloat(1), ans)
}

// sqrt(1.0001^tick)*2**96
// tick to sqrtx96
func GetTickToSqrtX96(tick *big.Int) *big.Int {
	pow, _ := IntToFloat(tick).Float64()
	tickExp := new(big.Float).SetPrec(96).SetFloat64(math.Pow(1.0001, pow))
	//
	floatingPoint := IntToFloat(new(big.Int).Exp(big.NewInt(2), big.NewInt(192), nil))
	ansSq := new(big.Float).SetPrec(96).Mul(tickExp, floatingPoint)
	ans := new(big.Float).SetPrec(96).Sqrt(ansSq)
	sqrtX96, _ := ans.Int(nil)
	return sqrtX96
}
