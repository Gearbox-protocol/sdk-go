package core

import (
	"os"

	"github.com/Gearbox-protocol/sdk-go/configs"
	"github.com/google/go-jsonnet"
)

type JsonnetImports struct {
	Embedded map[string]string
	Files    map[string]string
}

func GetJsonFromJsonnetData(mainFileData string, importedFiles JsonnetImports) (string, error) {
	vm := jsonnet.MakeVM()
	imports := map[string]jsonnet.Contents{}
	if err := getJsonnetFileImports(importedFiles.Files, imports); err != nil {
		return "", err
	}
	if err := getEmbeddedJsonnetImports(importedFiles.Embedded, imports); err != nil {
		return "", err
	}
	//
	vm.Importer(&jsonnet.MemoryImporter{
		Data: imports,
	})
	//
	if jsonStr, err := vm.EvaluateAnonymousSnippet("sample.jsonnet", mainFileData); err == nil {
		return jsonStr, nil
	} else {
		return "", err
	}
}

func getJsonnetFileImports(importedFiles map[string]string, imports map[string]jsonnet.Contents) error {
	for k, importedFile := range importedFiles {
		content, err := os.ReadFile(importedFile)
		if err != nil {
			return err
		}
		imports[k] = jsonnet.MakeContents(string(content))
	}
	return nil
}

func getEmbeddedJsonnetImports(importedFiles map[string]string, imports map[string]jsonnet.Contents) error {
	embeddedFiles := configs.Configs
	for k, importedFile := range importedFiles {
		content, err := embeddedFiles.ReadFile(importedFile)
		if err != nil {
			return err
		}
		imports[k] = jsonnet.MakeContents(string(content))
	}
	return nil
}

func GetJsonnetFile(mainFile string, importedFiles JsonnetImports) (string, error) {
	if mainFileData, err := os.ReadFile(mainFile); err != nil {
		return "", err
	} else {
		return GetJsonFromJsonnetData(string(mainFileData), importedFiles)
	}
}
func GetEmbeddedJsonnet(mainFile string, importedFiles JsonnetImports) (string, error) {
	if mainFileData, err := configs.Configs.ReadFile(mainFile); err != nil {
		return "", err
	} else {
		return GetJsonFromJsonnetData(string(mainFileData), importedFiles)
	}
}
