package core

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

func (anvil *AnvilClient) SendAsImpersonator(impAccount common.Address, tx *types.Transaction) common.Hash {
	err := anvil.ImpersonateAccount(impAccount.Hex())
	log.CheckFatal(err)
	return anvil.SendTransaction(impAccount, tx)
}

type anvilTransaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	GasPrice        string `json:"gasPrice"`
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

func (anvil *AnvilClient) SendTransaction(from common.Address, tx *types.Transaction) common.Hash {
	body := utils.GetJsonRPCRequestBody("eth_sendTransaction", anvilTransaction{
		From:            from.Hex(),
		To:              tx.To().Hex(),
		GasPrice:        bigIntToString(tx.GasPrice().Int64()),
		Gas:             bigIntToString(int64(tx.Gas())),
		Data:            hex.EncodeToString(tx.Data()),
		Nonce:           bigIntToString(int64(tx.Nonce())),
		Value:           fmt.Sprintf("%x", tx.Value()),
		TransactionType: bigIntToString(int64(tx.Type())),
	})
	result, err := utils.JsonRPCMakeRequest(anvil.url, body)
	if err != nil {
		log.Fatal(err, "from", from)
	}
	return common.HexToHash(result.(string))
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
