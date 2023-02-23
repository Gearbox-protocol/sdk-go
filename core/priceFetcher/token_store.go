package priceFetcher

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type TokensStore struct {
	mu       *sync.RWMutex
	decimals map[common.Address]schemas.Token
}

func NewTokensStore() *TokensStore {
	return &TokensStore{
		decimals: map[common.Address]schemas.Token{},
		mu:       &sync.RWMutex{},
	}
}

func (mdl *TokensStore) Update(tokens []schemas.Token) {
	store := map[common.Address]schemas.Token{}
	for _, token := range tokens {
		store[common.HexToAddress(token.Address)] = token
	}
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	mdl.decimals = store
}

func (mdl TokensStore) GetDecimals(token common.Address) int8 {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	v, ok := mdl.decimals[token]
	if !ok {
		log.Fatal("Decimal not found for token ", token)
	}
	return v.Decimals
}

func (mdl TokensStore) GetSymbol(token common.Address) core.Symbol {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	v, ok := mdl.decimals[token]
	if !ok {
		log.Fatal("Decimal not found for token ", token)
	}
	return core.Symbol(v.Symbol)
}
