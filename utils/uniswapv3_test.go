package utils

import (
	"math/big"
	"testing"
)

// token0 == usdc
// token1 wbtc
// wbtc .0009999 per usdc
func TestSqrtX96FromPrice(t *testing.T) {
	slot0 := GetSqrtX96FromPrice(big.NewInt(99990), 6)
	if slot0.String() != "25052892098943670850177508112" {
		t.Fatal("Wrong sqrtX96 for price 99990 with usdc as token0", slot0)
	}
}
func TestTickFromSqrtX96(t *testing.T) {
	tick := GetTickFromSqrtX96(StringToInt("25052894984021797146183221489"))
	if tick.String() != "-23027" {
		t.Fatal("Wrong tick for sqrtx96 with price[99990]", tick)
	}
}
func TestTickFromPrice(t *testing.T) {
	slot0 := GetSqrtX96FromPrice(big.NewInt(99990), 6)
	tick := GetTickFromSqrtX96(slot0)
	if tick.String() != "-23028" {
		t.Fatal("Wrong tick for price[99990] with usdc as token0", tick)
	}
}

func TestTickToPrice(t *testing.T) {
	slot0 := GetTickToSqrtX96(big.NewInt(-23028))
	if slot0.String() != "25052894984031011900217414762" {
		t.Fatal("Wrong slot0 for tick[-23028] ", slot0)
	}

	price := GetSqrtX96ToPrice(slot0, 6)
	if price.String() != "99990" {
		t.Fatal("Wrong price for tick[-23028] with usdc as token0", price)
	}
}

// func TestValue(t *testing.T) {
// 	slot0 := GetTickToSqrtX96(big.NewInt(-138162))
// 	t.Log(slot0)
// 	price := GetSqrtX96ToPrice(slot0, 6)
// 	t.Fatal(price)
// }

// func TestReverse(t *testing.T) {
// 	slot0 := GetSqrtX96FromPrice(big.NewInt(1), 6)
// 	tick := GetTickFromSqrtX96(slot0)
// 	t.Fatal(tick)
// }
