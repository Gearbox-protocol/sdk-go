package core

import (
	"encoding/json"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type Symbol string

type SymTOAddrStore struct {
	Exchanges map[string]common.Address `json:"exchanges"`
	Tokens    map[string]common.Address `json:"tokens"`
}

func GetSymToAddrStore(fileName string) *SymTOAddrStore {
	data, err := GetEmbeddedJsonnet(fileName, nil)
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
