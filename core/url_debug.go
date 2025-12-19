package core

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

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
	client := &http.Client{
		// Transport: LoggingRoundTripper{Proxied: http.DefaultTransport},
	}
	resp, err := client.Get(url)
	//
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
	debug.DecodeError = utils.ReadJsonReaderAndSetInterface(resp.Body, data)
	return debug
}

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Dump the request
	dumpReq, err := httputil.DumpRequestOut(req, true) // 'true' includes the body
	if err != nil {
		fmt.Printf("ERROR dumping request: %v\n", err)
	} else {
		fmt.Printf("REQUEST:\n%s\n", dumpReq)
	}

	// Send the request
	res, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil, err
	}

	// Dump the response
	dumpResp, err := httputil.DumpResponse(res, true) // 'true' includes the body
	if err != nil {
		fmt.Printf("ERROR dumping response: %v\n", err)
	} else {
		// Note: Dumping the response body here consumes the body stream.
		// To allow the original http.Client to use the body later, we need to handle it.
		fmt.Printf("RESULT:\n%s\n", dumpResp)
	}

	return res, nil
}
