package priceFetcher

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type TokensStore struct {
	mu     *sync.RWMutex
	tokens map[common.Address]schemas.Token
	client core.ClientI
}

func NewTokensStore(client core.ClientI) *TokensStore {
	return &TokensStore{
		tokens: map[common.Address]schemas.Token{},
		mu:     &sync.RWMutex{},
		client: client,
	}
}

func (mdl *TokensStore) Update(tokens []schemas.Token) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	for _, token := range tokens {
		mdl.tokens[common.HexToAddress(token.Address)] = token
	}
}

func (mdl TokensStore) GetTokens() (tokens []schemas.Token) {
	for _, entry := range mdl.tokens {
		tokens = append(tokens, entry)
	}
	return
}

func (mdl TokensStore) GetDecimals(tokenAddr common.Address) int8 {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	_, ok := mdl.tokens[tokenAddr]
	if !ok {
		token, err := schemas.NewToken(tokenAddr.Hex(), mdl.client)
		if err != nil {
			log.Fatalf("Err(%s) for token: %s", err, token.Address)
		}
		mdl.tokens[tokenAddr] = *token
	}
	return mdl.tokens[tokenAddr].Decimals
}

func (mdl TokensStore) GetSymbol(token common.Address) core.Symbol {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	v, ok := mdl.tokens[token]
	if !ok {
		log.Fatal("Decimal not found for token ", token)
	}
	return core.Symbol(v.Symbol)
}

func (mdl TokensStore) Exists(token common.Address) bool {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	_, ok := mdl.tokens[token]
	return ok
}
