package core

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

type FlashbotsTransactor struct {
	// flashbots
	RPC    *flashbotsrpc.FlashbotsRPC
	sigKey *ecdsa.PrivateKey
}

func NewFlashbotsTransactor(flashbotRelay string) *FlashbotsTransactor {
	rpc := flashbotsrpc.New(flashbotRelay)

	var sigKey, _ = crypto.GenerateKey()
	return &FlashbotsTransactor{
		// RPC:    rpc,
		RPC:    rpc,
		sigKey: sigKey,
	}
}

type BundleInfo struct {
	BlockNumber int64
	Stats       flashbotsrpc.FlashbotsGetBundleStatsResponse
	SendBundle  flashbotsrpc.FlashbotsSendBundleResponse
}

func (p *FlashbotsTransactor) Send(blockNum int64, txs []*types.Transaction) error {
	var txStrs []string
	for _, tx := range txs {
		b, err := tx.MarshalBinary()
		log.CheckFatal(err)
		txStr := hexutil.Encode(b)
		txStrs = append(txStrs, txStr)
	}
	blockNumStr := fmt.Sprintf("0x%x", blockNum)
	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         txStrs,
		BlockNumber: blockNumStr,
	}
	sendBundleResult, err := p.RPC.FlashbotsSendBundle(p.sigKey, sendBundleArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			log.Error(blockNum, err.Error())
		} else {
			log.Errorf("error(blockNum: %d): %+v\n", blockNum, err)
		}
		return err
	}

	statsResult, err := p.RPC.FlashbotsGetBundleStats(p.sigKey, flashbotsrpc.FlashbotsGetBundleStatsParam{
		BlockNumber: blockNumStr,
		BundleHash:  sendBundleResult.BundleHash,
	})
	if err != nil {
		return err
	}
	bInfo := BundleInfo{
		SendBundle:  sendBundleResult,
		Stats:       statsResult,
		BlockNumber: blockNum,
	}
	log.Info(utils.ToJson(bInfo))
	log.AMQPMsgf("Sucessfully paused contracts via flashbot relay at %d bundle(%s)", blockNum, sendBundleResult.BundleHash)
	return nil
}
