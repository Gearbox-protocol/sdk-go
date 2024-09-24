package test

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
)

type TestMask struct {
	Mask    *core.BigInt `json:"mask"`
	Account string       `json:"account"`
}

type TestCall struct {
	Pools    []dc.PoolCallData          `json:"pools"`
	CMs      []dc.CMCallData            `json:"cms"`
	Accounts []dc.CreditAccountCallData `json:"accounts"`
	//
	Masks      []TestMask `json:"masks"`
	OtherCalls OtherCalls `json:"others"`
}

func (c *TestCall) Process() {
}
