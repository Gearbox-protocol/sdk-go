package core

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type UrlDebug struct {
	ApiError     error
	url          string
	StatusCode   int
	JsonParseErr error
	Body         string
	DecodeError  error
}

func (d UrlDebug) String() (s string) {
	if d.ApiError != nil {
		s += fmt.Sprintf("ApiError: %v\n", d.ApiError)
		return
	}
	s += fmt.Sprintf("StatusCode: %d\n", d.StatusCode)
	if d.JsonParseErr != nil {
		s += fmt.Sprintf("JsonParseErr: %v\n", d.JsonParseErr)
		return
	} else {
		s += fmt.Sprintf("Body: %s\n. url: %s", d.Body, d.url)
	}
	if d.DecodeError != nil {
		s += fmt.Sprintf("DecodeError: %v\n", d.DecodeError)
		return
	}
	return
}

func (d UrlDebug) IsError() bool {
	return d.StatusCode/100 != 2 || d.JsonParseErr != nil || d.ApiError != nil || d.DecodeError != nil
}

func GetUrlWithDebug(url string, data interface{}) UrlDebug {
	log.Info(url)
	resp, err := http.Get(url)
	debug := UrlDebug{ApiError: err, url: url}
	if err != nil {
		return debug
	}
	//
	debug.StatusCode = resp.StatusCode
	if resp.StatusCode/100 != 2 {
		defer resp.Body.Close()
		bs, jsonParseErr := io.ReadAll(resp.Body)
		debug.JsonParseErr = jsonParseErr
		if jsonParseErr == nil {
			debug.Body = string(bs)
		}
	}
	log.Info(utils.ToJson(debug))
	debug.DecodeError = utils.ReadJsonReaderAndSetInterface(resp.Body, data)
	return debug
}
