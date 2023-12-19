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
	Exchanges map[string]common.Address `json:"exchanges"`
	Tokens    map[string]common.Address `json:"tokens"`
}

func GetSymToAddrStore(fileName string) *SymTOAddrStore {
	data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
	log.CheckFatal(err)
	store := &SymTOAddrStore{}
	err = json.Unmarshal([]byte(data), store)
	log.CheckFatal(err)
	return store
}

func GetAddrToSymbol(fileName string) map[common.Address]Symbol {
	store := GetSymToAddrStore(fileName)
	addrToName := map[common.Address]Symbol{}
	for name, token := range store.Tokens {
		addrToName[token] = Symbol(name)
	}
	for name, exchg := range store.Exchanges {
		addrToName[exchg] = Symbol(name)
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

func GetAddrToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	if chainId == 1337 || chainId == 7878 {
		chainId = 1
	}
	fileName := strings.ToLower(log.GetNetworkName(chainId)) + ".jsonnet"
	return GetAddrToSymbol(fileName)
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

func GetTokenGroups(fileName string) *TokenGroup {
	data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
	log.CheckFatal(err)
	store := &tokenGroupWrapper{}
	err = json.Unmarshal([]byte(data), store)
	log.CheckFatal(err)

	//
	obj := &TokenGroup{}
	{
		symToAddr := GetSymToAddrStore(fileName)
		for k, v := range store.Groups.CurvePools {
			obj.CurvePools[symToAddr.Tokens[string(k)].Hex()] = v
		}
		for k, v := range store.Groups.BalancerTokens {
			obj.BalancerTokens[symToAddr.Tokens[string(k)].Hex()] = v
		}
		for k, v := range store.Groups.ConvexCurveTokens {
			obj.ConvexCurveTokens[symToAddr.Tokens[string(k)].Hex()] = symToAddr.Tokens[string(v)].Hex()
		}
		for k, v := range store.Groups.YearnCurveTokens {
			obj.YearnCurveTokens[symToAddr.Tokens[string(k)].Hex()] = symToAddr.Tokens[string(v)].Hex()
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
