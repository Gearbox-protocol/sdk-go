package pkg

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

var pythUrl = "https://hermes.pyth.network"

type PythData struct {
	Price       *core.BigInt
	F           float64
	Data        []byte
	d           string
	Id          string
	PublishTime int64
}

func GetPythPrice(ids string, ts ...int64) (*PythData, error) {
	if len(ts) > 0 {
		if ts[0] >= 1761836699 && ids == "0xaf2a90410d6fdcaad452c6081fb7ec26cd6bffbe18b6dbc894f310e6e49585ee" {
			x := big.NewInt(0)
			return &PythData{
				Price:       (*core.BigInt)(x),
				F:           0,
				Data:        []byte{},
				Id:          ids,
				PublishTime: 0,
			}, nil
		}
		return getHistoric(ids, ts[0])
	}
	return getLatest(ids)
}

type pythBody struct {
	Binary struct {
		Data []string `json:"data"`
	} `json:"binary"`
	Parsed []struct {
		Price struct {
			Price       *core.BigInt `json:"price"`
			Expo        int8         `json:"expo"`
			PublishTime int64        `json:"publish_time"`
		} `json:"price"`
		Id string `json:"id"`
	} `json:"parsed"`
}

func (p pythBody) Convert(id string) (*PythData, error) {
	parsed := p.Parsed[0]
	first := common.BytesToHash(big.NewInt(parsed.Price.PublishTime).Bytes()).Hex()
	all := []string{
		first[2:],
		"0000000000000000000000000000000000000000000000000000000000000040",
		"0000000000000000000000000000000000000000000000000000000000000001",
		"0000000000000000000000000000000000000000000000000000000000000020",
		common.BytesToHash(big.NewInt(int64(len(p.Binary.Data[0])) / 2).Bytes()).Hex()[2:],
		p.Binary.Data[0],
		zeros(len(p.Binary.Data[0]) % 64),
	}
	d := strings.Join(all, "")
	data, err := hex.DecodeString(p.Binary.Data[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode price data: %s", err)
	}
	return &PythData{
		Price:       parsed.Price.Price,
		F:           utils.GetFloat64Decimal(parsed.Price.Price.Convert(), -1*parsed.Price.Expo),
		Data:        data,
		Id:          id,
		d:           d,
		PublishTime: parsed.Price.PublishTime,
	}, nil
}
func zeros(n int) string {
	if n == 0 {
		return ""
	}
	x := ""
	for i := 0; i < (64 - n); i++ {
		x += "0"
	}
	return x
}

func getLatest(ids string) (*PythData, error) {
	url := fmt.Sprintf("%s/v2/updates/price/latest?ids[]=%s", pythUrl, ids)
	data := pythBody{}
	err := core.GetUrlWithDebug(url, &data)
	if err.IsError() {
		return nil, fmt.Errorf("failed to get latest price from Pyth: %s", err)
	}
	return data.Convert(ids)
}

func getHistoric(ids string, ts int64) (*PythData, error) {
	url := fmt.Sprintf("%s/v2/updates/price/%d?ids[]=%s", pythUrl, ts, ids)
	data := pythBody{}
	err := core.GetUrlWithDebug(url, &data)
	if err.IsError() {
		if err.StatusCode == 404 && utils.GetEnvOrDefault("OPTIMISTIC_LIQUIDATION", "") == "1" {
			url := fmt.Sprintf("http://84.247.174.187:41001/proxy/any?id=%s&timestamp=%d", ids, ts) // dns resolution failing?
			log.Info("Using etherscan proxy for pyth", url)
			err = core.GetUrlWithDebug(url, &data)
		}
		if err.IsError() { // message rewritten in proxy is also error.
			return nil, fmt.Errorf("failed to get latest price from Pyth: %s. %s", err, url)
		}
	}
	return data.Convert(ids)
}
