package core

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ClientI interface {
	bind.ContractBackend
	ChainID(ctx context.Context) (*big.Int, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// for kovan arbitrage sync bot
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	// pendingNonceAt can lead to issues where the next tx will have (max pending tx nonce+1)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	//
	// HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error)
	// TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	// TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)
	// SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	// SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	// NetworkID(ctx context.Context) (*big.Int, error)
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	// CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	// FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	// PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	// PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	// PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// PendingTransactionCount(ctx context.Context) (uint, error)
	// CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	// SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// SendTransaction(ctx context.Context, tx *types.Transaction) error
}

// only use for calls that return address(20 bytes) or (amount 32 bytes) basically datatype uint256
func CallFuncGetSingleValue(client ClientI, sigStr string, to common.Address, blockNum int64, extra []byte) ([]byte, error) {
	bytes, err := CallFuncGetAllData(client, sigStr, to, blockNum, extra)
	if len(bytes) > 32 {
		bytes = bytes[:32]
	}
	return bytes, err
}
func CallFuncGetAllData(client ClientI, sigStr string, to common.Address, blockNum int64, extra []byte) ([]byte, error) {
	data, err := hex.DecodeString(sigStr) // enabledTokens
	log.CheckFatal(err)
	data = append(data, extra...)
	//
	msg := ethereum.CallMsg{
		To:   &to,
		Data: data,
	}
	var blockBI *big.Int
	if blockNum != 0 {
		blockBI = big.NewInt(blockNum)
	}
	bytes, err := client.CallContract(context.TODO(), msg, blockBI)
	return bytes, err
}

func GetAddress(client ClientI, field string, version int64) (common.Address, error) {
	providerAddr := GetAddressProvider(GetChainId(client), NewVersion(300))
	abi := GetAbi("AddressProviderv310")
	data, err := abi.Pack("addresses", field, big.NewInt(version))
	log.CheckFatal(err)
	results := MakeMultiCall(client, 0, false, []multicall.Multicall2Call{{
		Target:   common.HexToAddress(providerAddr),
		CallData: data,
	}})
	value, ok := MulticallAnsAddress(results[0])
	if !ok {
		err = fmt.Errorf("can't get %s from addrProvider", field)
	}
	return value, err
	// fbytes := []byte(field)
	// slot := make([]byte, 32)
	// //
	// copy(slot[:32], fbytes)
	// hash := common.BytesToHash(big.NewInt(version).Bytes())
	// slot = append(slot, hash[:]...)
	// //
	// addr, err := CallFuncGetSingleValue(client, "b76b70d5", common.HexToAddress(providerAddr), 0, slot) // addresses
	// return common.BytesToAddress(addr), err

}

func GetChainId(client ClientI) int64 {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	return chainId.Int64()
}
func GetBaseChainId(client ClientI) int64 {
	chainId, err := client.(interface {
		BaseChainID(ctx context.Context) (*big.Int, error)
	}).BaseChainID(context.TODO())
	log.CheckFatal(err)
	return chainId.Int64()
}
