package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AnvilClient struct {
	client http.Client
	url    string
}

func NewAnvilClient(url string) *AnvilClient {
	return &AnvilClient{
		client: http.Client{},
		url:    url,
	}
}

func getRequestBody(method string, params ...interface{}) *RequestBody {
	return &RequestBody{
		JsonRPC: "2.0",
		Id:      1,
		Method:  method,
		Params:  params,
	}
}

func (anvil *AnvilClient) ImpersonateAccount(account string) error {
	body := getRequestBody("anvil_impersonateAccount", account)
	_, err := anvil.makeRequest(body)
	return err
}

func (anvil *AnvilClient) makeRequest(body *RequestBody) (interface{}, error) {
	bodyBytes, err := json.Marshal(body)
	log.CheckFatal(err)
	reader := bytes.NewReader(bodyBytes)
	// make request
	req, err := http.NewRequest(http.MethodPost, anvil.url, reader)
	log.CheckFatal(err)
	req.Header["Content-Type"] = []string{"application/json"}
	// get response
	response, err := anvil.client.Do(req)
	log.CheckFatal(err)
	respBytes, err := ioutil.ReadAll(response.Body)
	log.CheckFatal(err)
	//
	data := ResponseBody{}
	json.Unmarshal(respBytes, &data)
	if data.Error != nil {
		return nil, fmt.Errorf("%+v", data.Error)
	} else {
		return data.Result, nil
	}
}

type ResponseBody struct {
	JsonRPC string      `json:"jsonrpc"`
	Error   interface{} `json:"error"`
	Result  interface{} `json:"result"`
}

type anvilTransaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	GasPrice        string `json:"gasPrice"`
	Gas             string `json:"gas"`
	Value           string `json:"value,omitempty"`
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	TransactionType string `json:"type"`
}

func bigIntToString(n int64) string {
	b := big.NewInt(n).Bytes()
	if len(b) == 0 {
		return "0x00"
	}
	return hex.EncodeToString(b)
}

func (anvil *AnvilClient) SendTransaction(from common.Address, tx *types.Transaction) common.Hash {
	body := getRequestBody("eth_sendTransaction", anvilTransaction{
		From:            from.Hex(),
		To:              tx.To().Hex(),
		GasPrice:        bigIntToString(tx.GasPrice().Int64()),
		Gas:             bigIntToString(int64(tx.Gas())),
		Data:            hex.EncodeToString(tx.Data()),
		Nonce:           bigIntToString(int64(tx.Nonce())),
		Value:           fmt.Sprintf("%x", tx.Value()),
		TransactionType: bigIntToString(int64(tx.Type())),
	})
	result, err := anvil.makeRequest(body)
	if err != nil {
		log.Fatal(err, "from", from)
	}
	return common.HexToHash(result.(string))
}

func (anvil *AnvilClient) TakeSnapshot() int64 {
	body := getRequestBody("evm_snapshot")
	result, err := anvil.makeRequest(body)
	log.CheckFatal(err)
	id, err := strconv.ParseInt(result.(string), 16, 64)
	log.CheckFatal(err)
	return id
}
func (anvil *AnvilClient) RevertSnapshot(id int64) {
	body := getRequestBody("evm_snapshot", id)
	result, err := anvil.makeRequest(body)
	if err != nil {
		log.Fatal("Snapshot revert to %d failed with %s", id, err)
	}
	if !result.(bool) {
		log.Fatal("revert %d to failed", id)
	}
}

type RequestBody struct {
	JsonRPC string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
