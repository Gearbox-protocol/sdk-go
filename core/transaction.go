package core

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	Transaction struct {
		Hash             string   `json:"hash"`
		Nonce            int64    `json:"nonce"`
		TransactionIndex int64    `json:"transaction_index"`
		From             string   `json:"from_addr"`
		To               string   `json:"to_addr"` // nil means contract creation
		Value            *big.Int `json:"value"`
		Gas              int64    `json:"gas_limit"`
		GasPrice         *big.Int `json:"gas_price"`
		Data             []byte   `json:"data"`
		BlockNum         int64    `json:"block_num"`
		Timestamp        int64    `json:"timestamp"`
	}
)

func GetWallet(pk string) keystore.Key {
	privateKey, err := crypto.HexToECDSA(pk)
	log.CheckFatal(err)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return keystore.Key{
		Address:    fromAddress,
		PrivateKey: privateKey,
	}
}

func ToDynamicTx(tx *types.Transaction) *types.DynamicFeeTx {
	if tx == nil {
		return nil
	}
	var gasFeeCap, gasTipCap *big.Int
	switch tx.Type() {
	case types.LegacyTxType:
	case types.DynamicFeeTxType:
		gasTipCap = tx.GasTipCap()
		gasFeeCap = tx.GasFeeCap()
	}
	return &types.DynamicFeeTx{
		Nonce:      tx.Nonce(),
		GasTipCap:  gasTipCap,
		GasFeeCap:  gasFeeCap,
		Gas:        tx.Gas(),
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}
}
