package core

import (
	"encoding/json"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type Symbol string

type RedStonePF struct {
	Type             int      `json:"type"`
	DataServiceId    string   `json:"dataServiceId"`
	DataId           string   `json:"dataId"`
	StalenessPeriod  int64    `json:"stalenessPeriod"`
	Signers          []string `json:"signers"`
	SignersThreshold int      `json:"signersThreshold"`
}
type RedStone struct {
	Mains    map[Symbol]RedStonePF `json:"mains"`
	Reserves map[Symbol]RedStonePF `json:"reserves"`
}
type SymTOAddrStore struct {
	Exchanges    map[string]common.Address `json:"exchanges"`
	Tokens       map[string]common.Address `json:"tokens"`
	FarmingPools map[string]common.Address `json:"farmingPools"`
	RedStone     RedStone                  `json:"redStone"`
}

func (s *SymTOAddrStore) getTokenAddr(sym Symbol) (string, bool) {
	if _, ok := s.Tokens[string(sym)]; !ok {
		// log.Fatal("can't get token", sym)
		return "", false
	}
	return s.Tokens[string(sym)].Hex(), true
}

var _globalCopy = map[string]*SymTOAddrStore{}

func getSymToAddrStore(fileName string) *SymTOAddrStore {
	if _globalCopy[fileName] == nil {
		data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
		log.CheckFatal(err)
		store := &SymTOAddrStore{}
		err = json.Unmarshal([]byte(data), store)
		log.CheckFatal(err)
		//
		_globalCopy[fileName] = store
	}
	return _globalCopy[fileName]
}

func GetRedStonePFByChainId(chainId int64) RedStone {
	fileName := log.GetConfigFile(chainId)
	data := getSymToAddrStore(fileName)
	return data.RedStone
}

func getAddrToSymbol(fileName string, opts map[string]bool) map[common.Address]Symbol {
	store := getSymToAddrStore(fileName)
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
	fileName := log.GetConfigFile(chainId)
	return getSymToAddrStore(fileName)
}

func GetFarmingPoolsToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	fileName := log.GetConfigFile(chainId)
	return getAddrToSymbol(fileName, map[string]bool{"farmingPools": true})
}

func GetTokenToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	fileName := log.GetConfigFile(chainId)
	return getAddrToSymbol(fileName, map[string]bool{"tokens": true})
}
func GetExchangeToSymbolByChainId(chainId int64) map[common.Address]Symbol {
	fileName := log.GetConfigFile(chainId)
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
		symToAddr := getSymToAddrStore(fileName)
		for k, v := range store.Groups.CurvePools {
			k, ok := symToAddr.getTokenAddr(k)
			if ok {
				obj.CurvePools[k] = v
			}
		}
		for k, v := range store.Groups.BalancerTokens {
			k, ok := symToAddr.getTokenAddr(k)
			if ok {
				obj.BalancerTokens[k] = v
			}
		}
		for k, v := range store.Groups.ConvexCurveTokens {
			k, ok := symToAddr.getTokenAddr(k)
			v, ok2 := symToAddr.getTokenAddr(v)
			if ok && ok2 {
				obj.ConvexCurveTokens[k] = v
			}
		}
		for k, v := range store.Groups.YearnCurveTokens {
			k, ok := symToAddr.getTokenAddr(k)
			v, ok2 := symToAddr.getTokenAddr(v)
			if ok && ok2 {
				obj.YearnCurveTokens[k] = v
			}
		}
	}
	return obj
}

func GetTokenGroupsByChainId(chainId int64) *TokenGroup {
	fileName := log.GetConfigFile(chainId)
	return GetTokenGroups(fileName)
}
