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

func (p *FlashbotsTransactor) Send(blockNum int64, txs []*types.Transaction) error {
	var txStrs []string
	for _, tx := range txs {
		b, err := tx.MarshalBinary()
		log.CheckFatal(err)
		txStr := hexutil.Encode(b)
		txStrs = append(txStrs, txStr)
	}
	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         txStrs,
		BlockNumber: fmt.Sprintf("0x%x", blockNum),
	}
	result, err := p.RPC.FlashbotsSendBundle(p.sigKey, sendBundleArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			log.Error(blockNum, err.Error())
		} else {
			log.Errorf("error(blockNum: %d): %+v\n", blockNum, err)
		}
		return err
	}
	log.Msgf("Sucessfully paused contracts at %d tx: ", blockNum, utils.ToJson(result))
	return nil
}
