package pkg

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AnvilClient struct {
	client http.Client
	url    string
}

func NewAnvilClient(url string) *AnvilClient {
	return &AnvilClient{
		client: http.Client{},
		url:    url,
	}
}

func (anvil *AnvilClient) ImpersonateAccount(account string) error {
	body := utils.GetJsonRPCRequestBody("anvil_impersonateAccount", account)
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	return err
}
func (anvil *AnvilClient) StopImpersonateAccount(account string) error {
	body := utils.GetJsonRPCRequestBody("anvil_stopImpersonatingAccount", account)
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	return err
}

func (anvil *AnvilClient) SendAsImpersonator(impAccount common.Address, tx *types.Transaction, nonce ...int64) (common.Hash, error) {
	// defer func() {
	// 	log.CheckFatal(anvil.StopImpersonateAccount(impAccount.Hex()))
	// }()
	err := anvil.ImpersonateAccount(impAccount.Hex())
	if err != nil {
		return common.Hash{}, log.WrapErrWithLine(err)
	}
	if len(nonce) > 0 {
		return anvil.SendTransaction(impAccount, tx, nonce[0])
	} else {
		return anvil.SendTransaction(impAccount, tx)
	}
}

type anvilTransaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	GasPrice        string `json:"gasPrice,omitempty"`
	MaxPriority     string `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerGas    string `json:"maxFeePerGas,omitempty"`
	Gas             string `json:"gas"`
	Value           string `json:"value,omitempty"`
	Data            string `json:"data"`
	Nonce           string `json:"nonce"`
	TransactionType string `json:"type"`
}

func bigIntToString(n int64) string {
	b := big.NewInt(n).Bytes()
	if len(b) == 0 {
		return "0000000000000000"
	}
	return hex.EncodeToString(b)
}

func (anvil *AnvilClient) SendTransaction(from common.Address, tx *types.Transaction, nonceArr ...int64) (common.Hash, error) {
	var anvilTx anvilTransaction
	nonce := func() (nonce int64) {
		if nonceArr == nil {
			nonce = int64(tx.Nonce())
		} else {
			nonce = nonceArr[0]
		}
		return
	}()

	if tx.Type() <= 1 { // 0,1
		anvilTx = anvilTransaction{
			From:            from.Hex(),
			To:              tx.To().Hex(),
			GasPrice:        "0x" + hex.EncodeToString(tx.GasTipCap().Bytes()),
			Gas:             "0x" + bigIntToString(int64(tx.Gas()*2)),
			Data:            hex.EncodeToString(tx.Data()),
			Nonce:           "0x" + bigIntToString(nonce),
			Value:           "0x" + fmt.Sprintf("%x", tx.Value()),
			TransactionType: bigIntToString(int64(tx.Type())),
		}
	} else {
		anvilTx = anvilTransaction{
			From:            from.Hex(),
			To:              tx.To().Hex(),
			MaxPriority:     hex.EncodeToString(tx.GasTipCap().Bytes()),
			MaxFeePerGas:    hex.EncodeToString(tx.GasFeeCap().Bytes()),
			Gas:             bigIntToString(int64(tx.Gas())),
			Data:            hex.EncodeToString(tx.Data()),
			Nonce:           bigIntToString(nonce),
			Value:           fmt.Sprintf("%x", tx.Value()),
			TransactionType: bigIntToString(int64(tx.Type())),
		}
	}
	body := utils.GetJsonRPCRequestBody("eth_sendTransaction", anvilTx)
	result, err := utils.JsonRPCMakeRequest(anvil.url, body)
	if err != nil {
		log.Debug(utils.ToJson(body))
		return common.Hash{}, log.WrapErrWithLine(err)
	}
	return common.HexToHash(result.(string)), nil
}

func (anvil *AnvilClient) SetEthBalance(to common.Address, value *big.Int) error {
	body := utils.GetJsonRPCRequestBody("anvil_setBalance", to.Hex(), fmt.Sprintf("0x%x", value))
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	return err
}
func (anvil *AnvilClient) EvmMine() {
	body := utils.GetJsonRPCRequestBody("evm_mine")
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	log.CheckFatal(err)
}
func (anvil *AnvilClient) TakeSnapshot() string {
	body := utils.GetJsonRPCRequestBody("evm_snapshot")
	result, err := utils.JsonRPCMakeRequest(anvil.url, body)
	log.CheckFatal(err)
	return result.(string)
}
func (anvil *AnvilClient) RevertSnapshot(id string) {
	log.Info("Reverting snapshot", id)
	body := utils.GetJsonRPCRequestBody("evm_revert", id)
	result, err := utils.JsonRPCMakeRequest(anvil.url, body)
	if err != nil {
		log.Fatalf("Snapshot revert to %d failed with %s", id, err)
	}
	if !result.(bool) {
		log.Fatalf("revert %s to failed", id)
	}
}
func (anvil *AnvilClient) NextTs(ts int64) error {
	log.Info("Setting next block anvil ts", ts)
	body := utils.GetJsonRPCRequestBody("evm_setNextBlockTimestamp", ts)
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	if err != nil {
		return fmt.Errorf("set next block ts to %d failed with %s", ts, err)
	}
	return nil
}

// utils
func GetImpersonationTopts(impersonateAddr common.Address, gasPrice *big.Int, gasLimit uint64) *bind.TransactOpts {
	topts := &bind.TransactOpts{
		From:     impersonateAddr,
		Context:  context.Background(),
		Signer:   func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) { return tx, nil },
		NoSend:   true,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
	}
	return topts
}
