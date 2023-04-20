package pkg

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type EtherScan struct {
	url        string
	client     http.Client
	ApiKey     string
	retryTimes int
}

func NewEtherScan(url string) *EtherScan {
	return &EtherScan{
		url:        url,
		ApiKey:     utils.GetEnvOrPanic("ETHERSCAN_API_KEY"),
		retryTimes: 3,
	}
}
func (e EtherScan) makeRequest(q url.Values) (response *http.Response, reqReq error) {
	req, err := http.NewRequest(http.MethodGet, e.url, nil)
	log.CheckFatal(err)
	req.URL.RawQuery = q.Encode()
	for i := 0; i < e.retryTimes; i++ {
		response, err = e.client.Do(req)
		if err != nil {
			return nil, err
		}
		if response.StatusCode != 200 {
			if response.StatusCode != 502 { // on bad gateway retry
				return nil, fmt.Errorf("status code: %d %s", response.StatusCode, utils.ReadJsonReader(response.Body))
			}
		} else {
			return response, nil
		}
		log.Info("retrying", i)
	}
	return nil, fmt.Errorf("status code: %d %s", response.StatusCode, utils.ReadJsonReader(response.Body))
}

func min[T int | int64](a, b T) T {
	if a > b {
		return b
	}
	return a
}

//
type statsI interface {
	Add(sig string, tx EtherScanCallInput)
	AllowedSig(sig string) bool
}

func (e *EtherScan) GetTxs(contract common.Address, endBlock, datapoints int, statsHandler statsI, ch chan EtherScanCallInput) (txs []EtherScanCallInput) {
	totalEntries := 0
	for {
		query := e.getQuery(contract, 0, endBlock, 1, min(datapoints, 10000))
		response, err := e.makeRequest(query)
		msg := &EtherscanResp{}
		utils.ReadJsonReaderAndSetInterface(response.Body, msg)
		for _, tx := range msg.Result {
			if tx.IsError == "0" && tx.Input != "0x" {
				sig := tx.Input[2:10]
				if statsHandler.AllowedSig(sig) {
					ch <- tx
					statsHandler.Add(sig, tx)
					totalEntries++
				}
			}
		}
		if len(msg.Result) == 0 {
			close(ch)
			return
		}
		blockNum, err := strconv.ParseInt(LastArrElem(msg.Result).BlockNumber, 10, 32)
		log.CheckFatal(err)
		endBlock = int(blockNum) - 1
		log.Verbosef("Fetched(%d) %s matching txs, new endBlock: %d", totalEntries, endBlock)
		if totalEntries > datapoints {
			close(ch)
			return
		}
	}
}

func LastArrElem[T any](x []T) T {
	if len(x) == 0 {
		panic("empty array")
	}
	return x[len(x)-1]
}

func (e EtherScan) getQuery(contract common.Address, start, end, page, offset int) url.Values {
	q := url.Values{}
	q.Add("module", "account")
	q.Add("action", "txlist")
	q.Add("address", contract.Hex())
	q.Add("startblock", fmt.Sprintf("%d", start))
	q.Add("endblock", fmt.Sprintf("%d", end))
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("offset", fmt.Sprintf("%d", offset))
	q.Add("sort", "desc")
	q.Add("apikey", e.ApiKey)
	return q
}

type EtherScanCallInput struct {
	Input       string `json:"input"`
	GasUsed     string `json:"gasUsed"`
	Hash        string `json:"hash"`
	IsError     string `json:"isError"`
	BlockNumber string `json:"blockNumber"`
	GasPrice    string `json:"gasPrice"`
	TimeStamp   int64  `json:"timeStamp,string,omitempty"`
}
type EtherscanResp struct {
	Result []EtherScanCallInput `json:"result"`
}
