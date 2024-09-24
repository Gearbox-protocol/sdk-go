package priceFetcher

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type TokensStore struct {
	mu     *sync.RWMutex
	tokens map[common.Address]*schemas.Token
	client core.ClientI
}

func NewTokensStore(client core.ClientI) *TokensStore {
	return &TokensStore{
		tokens: map[common.Address]*schemas.Token{},
		mu:     &sync.RWMutex{},
		client: client,
	}
}

func (mdl *TokensStore) Update(tokens []*schemas.Token) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	for _, token := range tokens {
		mdl.tokens[common.HexToAddress(token.Address)] = token
	}
}

func (mdl TokensStore) GetTokens() (tokens []*schemas.Token) {
	for _, entry := range mdl.tokens {
		tokens = append(tokens, entry)
	}
	return
}

func (mdl TokensStore) GetToken(addr string) *schemas.Token {
	return mdl.getToken(common.HexToAddress(addr))
}

func (mdl TokensStore) getToken(tokenAddr common.Address) *schemas.Token {
	mdl.mu.RLock()
	_, ok := mdl.tokens[tokenAddr]
	mdl.mu.RUnlock()
	if !ok {
		token, err := schemas.NewToken(tokenAddr.Hex(), mdl.client)
		if err != nil {
			log.Fatalf("Err(%s) for token: %s", err, token.Address)
		}
		mdl.mu.Lock()
		mdl.tokens[tokenAddr] = token
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

// filter already fetched tokens
func (mdl TokensStore) getNotPresentAddrs(addrs []common.Address) (ans []common.Address) {
	for _, addr := range addrs {
		if mdl.tokens[addr] == nil {
			ans = append(ans, addr)
		}
	}
	return
}
// if already has data on that token, doesn't fetch again
func (mdl TokensStore) GetDecimalsForList(addrs []common.Address) {
	mdl.mu.Lock()
	defer mdl.mu.Unlock()
	addrs = mdl.getNotPresentAddrs(addrs)
	//
	tokenABI := core.GetAbi("Token")
	// create calls
	calls := make([]multicall.Multicall2Call, 0, len(addrs))
	decimalsData, err := tokenABI.Pack("decimals")
	log.CheckFatal(err)
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
	// get results and process them
	results := core.MakeMultiCall(mdl.client, 0, false, calls, 20)
	itr := core.NewMulticallResultIterator(results)
	for _, addr := range addrs {
		result := itr.Next()
		decimals, ok := core.MulticallAnsBigInt(result)
		result = itr.Next()
		var sym string
		if result.Success {
			if values, err := tokenABI.Unpack("symbol", result.ReturnData); err == nil {
				if len(values) > 0 {
					sym = values[0].(string)
				}
			} else {
				sym, _ = schemas.SymbolFnReturnsBytes(mdl.client, addr)
			}
		}
		if sym == "" {
			log.Fatal("can't get symbol for token", addr.Hex())
		}
		if ok {
			mdl.tokens[addr] = &schemas.Token{
				Decimals: int8(decimals.Int64()),
				Symbol:   sym,
				Address:  addr.Hex(),
			}
		}
	}
}
