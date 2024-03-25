package ethclient

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

type R struct {
	Result struct {
		ForkConfig struct {
			ForkUrl string `json:"forkUrl"`
		} `json:"forkConfig"`
	} `json:"result"`
}

func getChainIdFromRPC(url string) (*big.Int, *big.Int, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, log.WrapErrWithLine(err)
	}
	chainId, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, nil, log.WrapErrWithLine(err)
	}
	return chainId, chainId, nil
}

func forkUrl(resp interface{}) string {
	defer func() {
		if err := recover(); err != nil {
			log.Info("error panic", err)
		}
	}()
	return resp.(map[string]interface{})["forkConfig"].(map[string]interface{})["forkUrl"].(string)
}

func GetFlagAndTestChainId(url string) (*big.Int, *big.Int, error) {
	body := utils.GetJsonRPCRequestBody("anvil_nodeInfo")
	resp, err := utils.JsonRPCMakeRequest(url, body)
	if err != nil {
		return getChainIdFromRPC(url)
	}

	forkUrl := forkUrl(resp)
	if forkUrl == "" {
		log.Fatalf("forkurl not found for %s", url)
	}

	_, test, err := GetFlagAndTestChainId(forkUrl)
	if err != nil {
		return nil, nil, log.WrapErrWithLine(err)
	}
	switch test.Int64() {
	case 1:
		return test, big.NewInt(7878), nil
	case 42161:
		return test, big.NewInt(7880), nil
	case 10:
		return test, big.NewInt(7879), nil
	default:
		return test, test, nil
	}
}
