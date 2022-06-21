package core

import (
	"os"

	"github.com/Gearbox-protocol/sdk-go/configs"
	"github.com/google/go-jsonnet"
)

type rf struct {
}

func (rf) ReadFile(file string) ([]byte, error) {
	return os.ReadFile(file)
}

type readFileI interface {
	ReadFile(file string) ([]byte, error)
}

func getJsonnet(file string, importedFiles map[string]string, reader readFileI) (string, error) {
	vm := jsonnet.MakeVM()
	imports := map[string]jsonnet.Contents{}
	for k, imporedFile := range importedFiles {
		content, err := reader.ReadFile(imporedFile)
		if err != nil {
			return "", err
		}
		imports[k] = jsonnet.MakeContents(string(content))
	}
	vm.Importer(&jsonnet.MemoryImporter{
		Data: imports,
	})
	if data, err := reader.ReadFile(file); err != nil {
		return "", err
	} else if jsonStr, err := vm.EvaluateAnonymousSnippet("sample.jsonnet", string(data)); err == nil {
		return jsonStr, nil
	} else {
		return "", err
	}
}

func GetJsonnet(file string, importedFiles map[string]string) (string, error) {
	return getJsonnet(file, importedFiles, rf{})
}
func GetEmbeddedJsonnet(file string, importedFiles map[string]string) (string, error) {
	return getJsonnet(file, importedFiles, configs.Configs)
}
