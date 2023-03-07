package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
)

type JsonRPCRequestBody struct {
	JsonRPC string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type JsonRPCResponseBody struct {
	JsonRPC string      `json:"jsonrpc"`
	Error   interface{} `json:"error"`
	Result  interface{} `json:"result"`
}

func GetJsonRPCRequestBody(method string, params ...interface{}) *JsonRPCRequestBody {
	return &JsonRPCRequestBody{
		JsonRPC: "2.0",
		Id:      1,
		Method:  method,
		Params:  params,
	}
}

func JsonRPCMakeRequest(url string, body *JsonRPCRequestBody) (interface{}, error) {
	bodyBytes, err := json.Marshal(body)
	log.CheckFatal(err)
	reader := bytes.NewReader(bodyBytes)
	// make request
	req, err := http.NewRequest(http.MethodPost, url, reader)
	log.CheckFatal(err)
	req.Header["Content-Type"] = []string{"application/json"}
	// get response
	response, err := (&http.Client{}).Do(req)
	log.CheckFatal(err)
	respBytes, err := ioutil.ReadAll(response.Body)
	log.CheckFatal(err)
	//
	data := JsonRPCResponseBody{}
	json.Unmarshal(respBytes, &data)
	if data.Error != nil {
		return nil, fmt.Errorf("%+v", data.Error)
	} else {
		return data.Result, nil
	}
}
