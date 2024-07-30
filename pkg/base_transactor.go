package pkg

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/howeyc/gopass"
)

type BaseTransactor struct {
	Topts      *bind.TransactOpts
	Client     core.ClientI
	timeoutSec int
}

func NewBaseTransactor(addr, privateKey string, client core.ClientI, timeoutSec int) *BaseTransactor {
	//
	chainId := core.GetChainId(client)
	//
	if strings.HasPrefix(privateKey, "enc:") {
		fmt.Printf("Enter password:")
		password, err := gopass.GetPasswd()
		log.CheckFatal(err)
		privateKey = utils.Decrypt(strings.Split(privateKey, ":")[1], password)
	}
	wallet := core.GetWallet(privateKey)
	if addr != "" && common.HexToAddress(addr) != wallet.Address {
		log.Fatal("Wrong prv key with addr", wallet.Address)
	}
	topts, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, big.NewInt(chainId))
	log.CheckFatal(err)
	//
	return &BaseTransactor{
		Topts:      topts,
		Client:     client,
		timeoutSec: timeoutSec,
	}
}

func (p *BaseTransactor) WaitForTx(job string, tx *types.Transaction) (*types.Receipt, error) {
	ctx, cancel := utils.GetTimeoutCtx(p.timeoutSec)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, p.Client, tx)
	if err != nil {
		return receipt, fmt.Errorf("tx failed for %s(%s): %s",
			job, utils.ToJson(core.ToDynamicTx(tx)), err.Error())
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return receipt, fmt.Errorf("tx not successful for %s(%s): %+v",
			job, utils.ToJson(core.ToDynamicTx(tx)), receipt)
	}
	ethUsed, gasUsed := p.getEthUsed(receipt)
	ts := func() int64 {
		ts, err := p.Client.HeaderByNumber(context.TODO(), receipt.BlockNumber)
		if err != nil {
			return 0
		}
		return int64(ts.Time)

	}()
	log.AMQPMsgf("%s TxHash: %s/tx/%s used eth %f  and gas used is %d. Ts: %s", job, log.NetworkUIUrl(core.GetChainId(p.Client)).ExplorerUrl, receipt.TxHash.Hex(),
		utils.GetFloat64Decimal(ethUsed, 18), gasUsed, time.Unix(ts, 0))
	return receipt, nil
}

func (p *BaseTransactor) getEthUsed(receipt *types.Receipt) (*big.Int, uint64) {
	node := Node{Client: p.Client}
	header, err := p.Client.HeaderByNumber(context.TODO(), big.NewInt(receipt.BlockNumber.Int64()))
	log.CheckFatal(err)
	return node.EthUsed(receipt.TxHash, header.BaseFee), receipt.GasUsed
}
