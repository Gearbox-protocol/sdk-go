package core

import (
	"testing"
)

type pythBody struct {
	Binary struct {
		Data []string `json:"data"`
	} `json:"binary"`
	Parsed []struct {
		Price struct {
			Price       string `json:"price"`
			Expo        int8   `json:"expo"`
			PublishTime int64  `json:"publish_time"`
		} `json:"price"`
		Id string `json:"id"`
	} `json:"parsed"`
}

func TestUrlDebug(t *testing.T) {
	data := pythBody{}
	url := "http://84.247.174.187:41001/proxy/any?id=0x9554b381f1e1d12f5e2baa6a9c802114284df21ced6c90490c565d1334e72c94&timestamp=1765926742"
	err := GetUrlWithDebug(url, &data)
	t.Log(err.IsError(), err)
}
