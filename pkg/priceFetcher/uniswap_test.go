package priceFetcher

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

func TestSlot0(t *testing.T) {
	price := univ3SlotToPriceInBase(utils.StringToInt("1240331789286057424084068704635"), true, 8)
	log.Info(price)
}
