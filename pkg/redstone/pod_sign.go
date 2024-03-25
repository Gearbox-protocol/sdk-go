package redstone

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"

	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type RSPriceOnDemandObj struct {
	CallData  string  `json:"callData"`
	Timestamp int64   `json:"ts"`
	Price     float64 `json:"price"`
}
type RSPriceOnDemand struct {
	RSPriceOnDemandObj
	PriceBI *big.Int // can't be null
	// MetaData  string `json:"metaData"`
	pod dcv3.PriceOnDemand
}

// sets pod and priceBI
func (r *RSPriceOnDemandObj) convert(token common.Address) *RSPriceOnDemand {
	priceOnDemand := dcv3.PriceOnDemand{}
	priceOnDemand.Token = token
	var err error
	if len(r.CallData) >= 2 && r.CallData[:2] == "0x" {
		r.CallData = r.CallData[2:]
	}
	priceOnDemand.CallData, err = hex.DecodeString(r.CallData) // due to 0x
	if err != nil {
		log.Warnf("failed to decode callData for token %s: %s", token, r.CallData)
		return nil
	}
	return &RSPriceOnDemand{
		RSPriceOnDemandObj: *r,
		PriceBI:            utils.FloatDecimalsTo64(r.Price, 8),
		pod:                priceOnDemand,
	}
}

func getLatestPodSign(details core.RedStonePF) map[string]*RSPriceOnDemandObj {
	// prod/aave/1
	url := fmt.Sprintf("https://testnet.gearbox.foundation/redstone/%s/%d?dataFeeds=%s", details.DataServiceId, details.SignersThreshold, details.DataId)
	return getpodSign(url, "latest-"+details.DataId)
}
func getHistoricPodSign(timestamp int64, details core.RedStonePF) map[string]*RSPriceOnDemandObj {
	// prod/aave/1
	timestamp = tenthMillSec(timestamp) // due to node js
	url := fmt.Sprintf("https://testnet.gearbox.foundation/redstone/%s/%d/%d?dataFeeds=%s", details.DataServiceId, details.SignersThreshold, timestamp, details.DataId)
	return getpodSign(url, "historic-"+details.DataId)
}

func getpodSign(url string, dataId string) map[string]*RSPriceOnDemandObj {
	res, err := http.Get(url)
	log.Debug("Getting priceOnDemand", url)
	if err != nil || res.StatusCode/100 != 2 {
		respStr := ""
		if res.Body != nil {
			_, respStr = ReadBuffer(res.Body)
		}
		log.Warnf("For dataId %s , redStone failed with err(%v) and resp(%s)", dataId, err, respStr)
		return nil
	}
	data := map[string]*RSPriceOnDemandObj{}
	if err = utils.ReadJsonReaderAndSetInterface(res.Body, &data); err != nil {
		log.Warnf("For dataId %s , redStone response parsing failed with  err(%v)", dataId, err)
		return nil
	}
	return data
}
