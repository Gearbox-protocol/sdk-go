package test

import "github.com/Gearbox-protocol/sdk-go/core"

type TestMask struct {
	Mask    *core.BigInt `json:"mask"`
	Account string       `json:"account"`
}

type TestCall struct {
	Pools    []TestPoolCallData    `json:"pools"`
	CMs      []TestCMCallData      `json:"cms"`
	Accounts []TestAccountCallData `json:"accounts"`
	//
	Masks      []TestMask `json:"masks"`
	OtherCalls OtherCalls `json:"others"`
}

func (c *TestCall) Process() {
	return
}
