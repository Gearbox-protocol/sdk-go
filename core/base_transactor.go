package core

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type BaseTransactor struct {
	Topts   *bind.TransactOpts
	Client  ClientI
	ChainId int64
}

func NewBaseTransactor(privateKey string, client ClientI) *BaseTransactor {
	//
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	//
	wallet := GetWallet(privateKey)
	topts, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(chainId.Int64()))
	log.CheckFatal(err)
	//
	return &BaseTransactor{
		Topts:   topts,
		Client:  client,
		ChainId: chainId.Int64(),
	}
}

func (p *BaseTransactor) WaitForTx(job string, tx *types.Transaction) error {
	ctx, cancel := utils.GetTimeoutCtx(60)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, p.Client, tx)
	if err != nil {
		return fmt.Errorf("tx failed for %s: %s", job, err.Error())
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("tx not successful for %s %+v", job, receipt)
	}
	log.Msgf("%s TxHash: %s/tx/%s.", job, NetworkUIUrl(uint(p.ChainId)).ExplorerUrl, receipt.TxHash.Hex())
	return nil
}
