package core

import (
	"math"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

var _forkBlock int64

func urlIsAnvil(url string) bool {
	return strings.Contains(url, "anvil.gearbox.foundation") || strings.Contains(url, "localhost")
}

// if url is anvil and the request to get fork block is successful, it returns the fork block number
func GetForkBlock(url string) (block int64) {
	if _forkBlock != 0 {
		return _forkBlock
	}
	if urlIsAnvil(url) { // fork or anvil
		body := utils.GetJsonRPCRequestBody("anvil_nodeInfo")
		data, err := utils.JsonRPCMakeRequest(url, body)
		if err != nil {
			return
		}
		cfg := data.(map[string]interface{})["forkConfig"].(map[string]interface{})
		if anvilUrl := cfg["forkUrl"].(string); urlIsAnvil(anvilUrl) {
			log.Info("checking", anvilUrl)
			return GetForkBlock(anvilUrl)
		} else {
			_forkBlock = int64(cfg["forkBlockNumber"].(float64))
			log.Info("Anvil fork block", _forkBlock, "from", url)
			return _forkBlock
		}
	}
	return math.MaxInt64
}
