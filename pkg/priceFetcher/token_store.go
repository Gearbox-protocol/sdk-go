package priceFetcher

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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

func (mdl TokensStore) getToken(tokenAddr common.Address) schemas.Token {
	mdl.mu.RLock()
	_, ok := mdl.tokens[tokenAddr]
	mdl.mu.RUnlock()
	if !ok {
		token, err := schemas.NewToken(tokenAddr.Hex(), mdl.client)
		if err != nil {
			log.Fatalf("Err(%s) for token: %s", err, token.Address)
		}
		mdl.mu.Lock()
		mdl.tokens[tokenAddr] = *token
		mdl.mu.Unlock()
	}
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	return mdl.tokens[tokenAddr]
}

func (mdl TokensStore) GetDecimals(tokenAddr common.Address) int8 {
	return mdl.getToken(tokenAddr).Decimals
}

func (mdl TokensStore) GetSymbol(tokenAddr common.Address) core.Symbol {
	return core.Symbol(mdl.getToken(tokenAddr).Symbol)
}

func (mdl TokensStore) Exists(token common.Address) bool {
	mdl.mu.RLock()
	defer mdl.mu.RUnlock()
	_, ok := mdl.tokens[token]
	return ok
}

func (mdl TokensStore) GetDecimalsForList(addrs []common.Address) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	calls := make([]multicall.Multicall2Call, 0, len(addrs))
	decimalsData, err := core.GetAbi("Token").Pack("decimals")
	symbolData, err := core.GetAbi("Token").Pack("symbol")
	log.CheckFatal(err)
	for _, addr := range addrs {
		calls = append(calls, multicall.Multicall2Call{
			Target:   addr,
			CallData: decimalsData,
		}, multicall.Multicall2Call{
			Target:   addr,
			CallData: symbolData,
		})
	}
	results := core.MakeMultiCall(mdl.client, 0, false, calls, 20)
	itr := core.NewMulticallResultIterator(results)
	for _, addr := range addrs {
		result := itr.Next()
		decimals, ok := core.MulticallAnsBigInt(result)
		result = itr.Next()
		var sym string
		if result.Success {
			sym = string(result.ReturnData[utils.Min(64, len(result.ReturnData)):])
		}
		if ok {
			mdl.tokens[addr] = schemas.Token{
				Decimals: int8(decimals.Int64()),
				Symbol:   sym,
				Address:  addr.Hex(),
			}
		}
	}
}
