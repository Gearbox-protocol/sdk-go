package priceFetcher

import "github.com/Gearbox-protocol/sdk-go/artifacts/multicall"

type TokenCalls struct {
	calls  []multicall.Multicall2Call
	tokens []string
}
