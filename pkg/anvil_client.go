package pkg

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
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

func (anvil *AnvilClient) SendAsImpersonator(impAccount common.Address, tx *types.Transaction, nonce ...int64) common.Hash {
	// defer func() {
	// 	log.CheckFatal(anvil.StopImpersonateAccount(impAccount.Hex()))
	// }()
	err := anvil.ImpersonateAccount(impAccount.Hex())
	log.CheckFatal(err)
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
		return "0x00"
	}
	return hex.EncodeToString(b)
}

func (anvil *AnvilClient) SendTransaction(from common.Address, tx *types.Transaction, nonceArr ...int64) common.Hash {
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
		log.Fatal("from", from, "err:", err)
	}
	return common.HexToHash(result.(string))
}

func (anvil *AnvilClient) SetEthBalance(to common.Address, value *big.Int) {
	body := utils.GetJsonRPCRequestBody("anvil_setBalance", to.Hex(), fmt.Sprintf("%x", value))
	_, err := utils.JsonRPCMakeRequest(anvil.url, body)
	log.CheckFatal(err)
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
