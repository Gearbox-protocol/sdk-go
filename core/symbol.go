package core

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type Symbol string

type SymTOAddrStore struct {
	Exchanges    map[string]common.Address `json:"exchanges"`
	Tokens       map[string]common.Address `json:"tokens"`
	FarmingPools map[string]common.Address `json:"farmingPools"`
}

func (s *SymTOAddrStore) getTokenAddr(sym Symbol) string {
	if _, ok := s.Tokens[string(sym)]; !ok {
		log.Fatal("can't get token", sym)
	}
	return s.Tokens[string(sym)].Hex()
}

func GetSymToAddrStore(fileName string) *SymTOAddrStore {
	data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
	log.CheckFatal(err)
	store := &SymTOAddrStore{}
	err = json.Unmarshal([]byte(data), store)
	log.CheckFatal(err)
	return store
}

func getAddrToSymbol(fileName string, opts map[string]bool) map[common.Address]Symbol {
	store := GetSymToAddrStore(fileName)
	addrToName := map[common.Address]Symbol{}
	if opts["tokens"] {
		for name, token := range store.Tokens {
			addrToName[token] = Symbol(name)
		}
	}
	if opts["exchanges"] {
		for name, exchg := range store.Exchanges {
			addrToName[exchg] = Symbol(name)
		}
	}
	if opts["farmingPools"] {
		for name, exchg := range store.FarmingPools {
			addrToName[exchg] = Symbol(name)
		}
	}
	return addrToName
}

func GetSymToAddrByChainId(chainId int64) *SymTOAddrStore {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return GetSymToAddrStore(fileName)
}

func GetFarmingPoolsToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return getAddrToSymbol(fileName, map[string]bool{"farmingPools": true})
}

func GetTokenToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return getAddrToSymbol(fileName, map[string]bool{"tokens": true})
}
func GetExchangeToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return getAddrToSymbol(fileName, map[string]bool{"exchanges": true})
}

func GetDecimals(client ClientI, addr common.Address, blockNum int64) int8 {
	decimals, err := CallFuncWithExtraBytes(client, "313ce567", addr, blockNum, nil) // decimals
	if err != nil {
		log.Fatalf("Can't get decimals for addr(%s) : %s", addr, err)
	}
	return int8(new(big.Int).SetBytes(decimals).Int64())
}

type PriceSource string

const (
	SOURCE_CHAINLINK PriceSource = "chainlink"
	SOURCE_SPOT      PriceSource = "spot"
)

type TokenGroup struct {
	CurvePools        map[string]int64  `json:"curvePools"`
	BalancerTokens    map[string]int64  `json:"balancerTokens"`
	YearnCurveTokens  map[string]string `json:"yearnCurveTokens"`
	ConvexCurveTokens map[string]string `json:"convexCurveTokens"`
}

type tokenGroupWrapper struct {
	Groups struct {
		CurvePools        map[Symbol]int64  `json:"curvePools"`
		BalancerTokens    map[Symbol]int64  `json:"balancerTokens"`
		YearnCurveTokens  map[Symbol]Symbol `json:"yearnCurveTokens"`
		ConvexCurveTokens map[Symbol]Symbol `json:"convexCurveTokens"`
	} `json:"groups"`
}

func newTokenGroup() *TokenGroup {
	return &TokenGroup{
		CurvePools:        map[string]int64{},
		BalancerTokens:    map[string]int64{},
		YearnCurveTokens:  map[string]string{},
		ConvexCurveTokens: map[string]string{},
	}
}
func GetTokenGroups(fileName string) *TokenGroup {
	data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
	log.CheckFatal(err)
	store := &tokenGroupWrapper{}
	err = json.Unmarshal([]byte(data), store)
	log.CheckFatal(err)

	//
	obj := newTokenGroup()
	{
		symToAddr := GetSymToAddrStore(fileName)
		for k, v := range store.Groups.CurvePools {
			obj.CurvePools[symToAddr.getTokenAddr(k)] = v
		}
		for k, v := range store.Groups.BalancerTokens {
			obj.BalancerTokens[symToAddr.getTokenAddr(k)] = v
		}
		for k, v := range store.Groups.ConvexCurveTokens {
			obj.ConvexCurveTokens[symToAddr.getTokenAddr(k)] = symToAddr.getTokenAddr(v)
		}
		for k, v := range store.Groups.YearnCurveTokens {
			obj.YearnCurveTokens[symToAddr.getTokenAddr(k)] = symToAddr.getTokenAddr(v)
		}
	}
	return obj
}

func GetTokenGroupsByChainId(chainId int64) *TokenGroup {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return GetTokenGroups(fileName)
}
