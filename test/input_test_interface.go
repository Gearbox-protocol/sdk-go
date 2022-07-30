package test

import (
	"encoding/json"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type TestInputI interface {
	Merge(b TestInputI)
}

func LoadTestFiles(filePaths []string, files []TestInputI, t *testing.T) (mergedFile TestInputI, addrMap core.AddressMap) {
	addrMap = core.AddressMap{}
	for i, filePath := range filePaths {
		// get test input file
		addAddressSetJson(filePath, files[i], addrMap, t)
		if mergedFile == nil {
			mergedFile = files[i]
		} else {
			mergedFile.Merge(files[i])
		}
	}
	return
}

func CreateDefaultTestInput(filePaths []string, t *testing.T) (*TestInput, core.AddressMap) {
	filesI := make([]TestInputI, len(filePaths))
	for i := 0; i < len(filePaths); i++ {
		filesI[i] = NewTestInput()
	}
	outputFile, addrMap := LoadTestFiles(filePaths, filesI, t)
	mergedFile := outputFile.(*TestInput)
	return mergedFile, addrMap
}

func addAddressSetJson(filePath string, obj TestInputI, addressMap core.AddressMap, t *testing.T) {
	var mock core.Json = utils.ReadJson(filePath)
	mock.ParseAddress(t, addressMap)
	// log.Info(utils.ToJson(mock))
	b, err := json.Marshal(mock)
	if err != nil {
		t.Error(err)
	}
	utils.SetJson(b, obj)
}
