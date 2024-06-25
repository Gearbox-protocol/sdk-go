package ethclient

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
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
	_, err := utils.JsonRPCMakeRequest(url, body)
	if err != nil {
		return getChainIdFromRPC(url)
	}

	var flagChainId *big.Int = new(big.Int)
	for _, netId := range []int64{1, 42161, 10} {
		addrs := core.GetSymToAddrByChainId(netId)
		usdc := addrs.Tokens["USDC"]
		client, err := ethclient.Dial(url)
		log.CheckFatal(err)
		_, err = core.CallFuncWithExtraBytes(client, "95d89b41", usdc, 0, nil)
		if err == nil {
			flagChainId = big.NewInt(netId)
			break
		}
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return flagChainId, nil, err
	}
	testChainId, err := client.ChainID(context.Background())
	return flagChainId, testChainId, err
}
