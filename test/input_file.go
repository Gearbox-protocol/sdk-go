package test

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/core/types"
)

type TestInput struct {
	Blocks    map[int64]*BlockInput `json:"blocks"`
	States    TestState             `json:"states"`
	MockFiles map[string]string     `json:"mocks"`
}

func NewTestInput() *TestInput {
	return &TestInput{
		Blocks: map[int64]*BlockInput{},
	}
}

type BlockInput struct {
	Events []TestEvent `json:"events"`
	Calls  TestCall    `json:"calls"`
}

func (file *TestInput) Merge(other TestInputI) {
	otherFile := other.(*TestInput)
	//
	if file.Blocks == nil {
		file.Blocks = make(map[int64]*BlockInput)
	}
	for blockNum, block := range otherFile.Blocks {
		if file.Blocks[blockNum] != nil {
			log.Fatal("can't merge blocks for testinput, block found in both files", blockNum, utils.ToJson(file), utils.ToJson(other))
		}
		file.Blocks[blockNum] = block
	}
	//
	if file.MockFiles == nil {
		file.MockFiles = make(map[string]string)
	}
	for mockType, mockPath := range otherFile.MockFiles {
		file.MockFiles[mockType] = mockPath
	}
	//
	file.States.Merge(otherFile.States)
}

type TokenFile struct {
	Tokens []*schemas.Token `json:"tokens"`
}

func (*TokenFile) Merge(TestInputI) {

}

func (file *TestInput) AddToClient(client *TestClient, addrMap core.AddressMap) {
	// load tokens
	if tokenFilePath := file.MockFiles["tokens"]; tokenFilePath != "" {
		obj := &TokenFile{}
		addAddressSetJson(tokenFilePath, obj, addrMap, nil)
		for _, tokenObj := range obj.Tokens {
			switch tokenObj.Symbol {
			case "USDC":
				client.SetUSDC(tokenObj.Address)
			case "WETH":
				client.SetWETH(tokenObj.Address)
			}
			client.AddToken(tokenObj)
		}
	}
	// blockNum  -> funcsig -> call response data
	otherCalls := map[int64]OtherCalls{}

	// addr to contract type
	addrToType := map[string]string{}
	for name, addr := range addrMap {
		addrToType[addr] = strings.Split(name, "_")[0]
	}
	// load events in client
	events := map[int64]map[string][]types.Log{}
	accountMask := map[int64]map[string]*big.Int{}

	for blockNum, block := range file.Blocks {
		otherCalls[blockNum] = block.Calls.OtherCalls
		if events[blockNum] == nil {
			events[blockNum] = make(map[string][]types.Log)
		}
		for ind, event := range block.Events {
			txLog := event.Process(addrToType[event.Address])
			txLog.Index = uint(ind)
			txLog.BlockNumber = uint64(blockNum)
			events[blockNum][event.Address] = append(events[blockNum][event.Address], txLog)
		}
		// get masks
		for _, maskDetails := range block.Calls.Masks {
			if accountMask[blockNum] == nil {
				accountMask[blockNum] = make(map[string]*big.Int)
			}
			accountMask[blockNum][maskDetails.Account] = maskDetails.Mask.Convert()

		}
	}
	client.SetMasks(accountMask)
	client.SetState(&file.States, otherCalls)
	client.SetEvents(events)
}

// for matching state with the expected output
func ReplaceWithVariable(obj interface{}, addrMap core.AddressMap) core.Json {
	bytes, err := json.Marshal(obj)
	log.CheckFatal(err)
	addrToVariable := core.AddressMap{}
	// TODO: FIX FOR HASH
	for variable, addr := range addrMap {
		addrToVariable[addr] = "#" + variable
		addrToVariable[strings.ToLower(addr)] = "#" + variable
	}
	outputJson := core.Json{}
	err = json.Unmarshal(bytes, &outputJson)
	log.CheckFatal(err)
	outputJson.ReplaceWithVariable(addrToVariable)
	return outputJson
}
