package test

import (
	"math/big"
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/ethereum/go-ethereum/core/types"
)

func (t *TestClient) SetUSDC(addr string) {
	t.USDCAddr = addr
}
func (t *TestClient) SetWETH(addr string) {
	t.WETHAddr = addr
}
func (t *TestClient) SetState(state *TestState, calls map[int64]OtherCalls) {
	t.state.addState(state, calls)
}

func (t *TestClient) AddToken(token *schemas.Token) {
	t.token[token.Address] = token
}
func (t TestClient) GetToken() map[string]*schemas.Token {
	return t.token
}

// blocknum => event address => txlogs
func (t *TestClient) SetEvents(obj map[int64]map[string][]types.Log) {
	if t.events == nil {
		t.events = map[int64]map[string][]types.Log{}
	}
	for blockNum, logs := range obj {
		t.events[blockNum] = logs
	}
	blockNums := []int64{}
	for blockNum := range t.events {
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	t.blockNums = blockNums
}

// token => block => prices
func (t *TestClient) SetPrices(obj map[string]map[int64]*big.Int) {
	if t.prices == nil {
		t.prices = map[string]map[int64]*big.Int{}
	}
	for token, block := range obj {
		if t.prices[token] == nil {
			t.prices[token] = map[int64]*big.Int{}
		}
		for blockNum, price := range block {
			t.prices[token][blockNum] = price
		}
	}
}

// block => account => mask
func (t *TestClient) SetMasks(masks map[int64]map[string]*big.Int) {
	if t.masks == nil {
		t.masks = map[int64]map[string]*big.Int{}
	}
	for blockNum, mask := range masks {
		t.masks[blockNum] = mask
	}
}
