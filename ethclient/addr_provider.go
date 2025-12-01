package ethclient

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
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

func GetBaseAndCurrentChainId(url string) (*big.Int, *big.Int, error) {
	body := utils.GetJsonRPCRequestBody("anvil_nodeInfo")
	_, err := utils.JsonRPCMakeRequest(url, body)
	if err != nil {
		return getChainIdFromRPC(url)
	}

	var baseChainId *big.Int = new(big.Int)
	// TODO: NEWNETWORK
	for netId := range log.Basenet { // fitst should be 146, as USDC_e address is not returning error
		var usdc common.Address
		if netId == 1135 || netId == 43111 || netId == 5031 { // hemi and lisk
			usdc = core.GetToken(netId, "USDC.e")
		} else if netId == 146 { // sonic
			usdc = core.GetToken(netId, "USDC_e")
		} else if netId == 9745 { // plasma
			usdc = core.GetToken(netId, "USDT0")
		} else {
			usdc = core.GetToken(netId, "USDC") // for optimism, arbitrum, etherum, monad, bnb, and etherlink
		}
		client, err := ethclient.Dial(url)
		log.CheckFatal(err)
		x, err := core.CallFuncGetSingleValue(client, "95d89b41", usdc, 0, nil) // symbol
		if err == nil && len(x) > 0 {
			baseChainId = big.NewInt(netId)
			break
		}
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return baseChainId, nil, err
	}
	currentChainId, err := client.ChainID(context.Background())
	return baseChainId, currentChainId, err
}
