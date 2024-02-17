package schemas

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
)

func TestPriceFeedClone(t *testing.T) {
	a := &PriceFeed{
		BlockNumber:     23,
		Feed:            "0x1111111111111111111111111111111111111111",
		RoundId:         24,
		MergedPFVersion: MergedPFVersion(V2PF),
		PriceBI:         (*core.BigInt)(big.NewInt(12345)),
		Price:           .12345,
	}
	b := a.Clone()
	if *a != *b {
		t.Errorf("%+v's cloned copy is differet %+v", a, b)
	}
}
