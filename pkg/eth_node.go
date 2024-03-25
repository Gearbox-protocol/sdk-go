package pkg

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Node struct {
	Client  core.ClientI
	chainId int64
}

func (lf Node) GetLogs(fromBlock, toBlock int64, addrs []common.Address, topics [][]common.Hash) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(fromBlock),
		ToBlock:   new(big.Int).SetInt64(toBlock),
		Addresses: addrs, //[]common.Address{common.HexToAddress(addr)},
		Topics:    topics,
	}
	var logs []types.Log
	var err error
	logs, err = lf.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if strings.Contains(err.Error(), core.QueryMoreThan10000Error) ||
			strings.Contains(err.Error(), core.NoderealFilterLogError) ||
			strings.Contains(err.Error(), core.AnkrRangeError) ||
			strings.Contains(err.Error(), core.LogFilterLenError) ||
			(strings.Contains(err.Error(), "we can't execute this request") && core.GetChainId(lf.Client) == 42161) || // for arbitrum get logs for account Manager
			// failure: we can't execute this request range  192549019 192549555 tokenAddrs 32 accountHashes 6
			err.Error() == core.LogFilterQueryTimeout {
			middle := (fromBlock + toBlock) / 2
			bottomHalfLogs, err := lf.GetLogs(fromBlock, middle-1, addrs, topics)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, bottomHalfLogs...)

			topHalfLogs, err := lf.GetLogs(middle, toBlock, addrs, topics)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, topHalfLogs...)
			return logs, nil
		}
	}
	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i].BlockNumber < logs[j].BlockNumber ||
			(logs[i].BlockNumber == logs[j].BlockNumber && logs[i].Index < logs[j].Index)
	})
	return logs, err
}

func (lf *Node) GetLatestBlockNumber() int64 {
	lf.setChainId()
	latestBlockNum, err := lf.Client.BlockNumber(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	blockNumToReturn := int64(latestBlockNum)
	return blockNumToReturn
}

func (lf *Node) GetLatestFinalizedBlock(skipBlocks int64) int64 {
	blockNum := lf.GetLatestBlockNumber()
	// skip 2 blocks ~30 sec latest block might reorder
	if lf.chainId != 1337 {
		blockNum -= skipBlocks
	}
	log.Info("Last finalized block", blockNum)
	return blockNum
}

func (lf *Node) setChainId() {
	if lf.chainId == 0 {
		chainId, err := lf.Client.ChainID(context.TODO())
		log.CheckFatal(err)
		lf.chainId = chainId.Int64()
	}
}

func (lf Node) GetHeader(blockNum int64) *types.Header {
	b, err := lf.Client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	log.CheckFatal(err)
	return b.Header()
}

func (lf Node) GasPrice(txHash common.Hash, baseFee *big.Int) *big.Int {
	tx, pending, err := lf.Client.TransactionByHash(context.TODO(), txHash)
	log.CheckFatal(err)
	if pending {
		log.Fatalf("Tx is pending, something not right %s", txHash.Hex())
	}
	if tx.Type() == 2 {
		return new(big.Int).Add(tx.GasTipCap(), baseFee)
	} else {
		return tx.GasPrice()
	}
}

func (lf Node) EthUsed(txHash common.Hash, baseFee *big.Int) *big.Int {
	receipt := lf.GetReceipt(txHash)
	gasUsed := big.NewInt(int64(receipt.GasUsed))
	return new(big.Int).Mul(lf.GasPrice(txHash, baseFee), gasUsed)
}

func (lf Node) GetReceipt(txHash common.Hash) *types.Receipt {
	receipt, err := lf.Client.TransactionReceipt(context.TODO(), txHash)
	log.CheckFatal(err)
	return receipt
}

func (lf Node) GetLogsForTransfer(queryFrom, queryTill int64, hexAddrs []common.Address, treasuryAddrTopic []common.Hash) ([]types.Log, error) {
	topics := [][]common.Hash{
		{
			core.Topic("Transfer(address,address,uint256)"),
		},
	}
	otherAddrTopic := []common.Hash{}
	// from treasury to other address
	logs, err := lf.GetLogs(queryFrom, queryTill, hexAddrs, append(topics, treasuryAddrTopic, otherAddrTopic))
	if err != nil {
		return logs, err
	}

	// from other address to treasury
	newLogs, err := lf.GetLogs(queryFrom, queryTill, hexAddrs, append(topics, otherAddrTopic, treasuryAddrTopic))
	if err != nil {
		return logs, err
	}
	return append(newLogs, logs...), nil
}

func GetBlockNumForTs(etherscanAPI string, chainId int64, ts int64) (int64, error) {
	url := "https://%s/api?module=block&action=getblocknobytime&timestamp=%d&closest=before&apikey=%s"
	var suffix string
	switch log.GetBaseNet(chainId) {
	case "MAINNET":
		suffix = "api.etherscan.io"
	case "ARBITRUM":
		suffix = "api.arbiscan.io"
	}
	url = fmt.Sprintf(url, suffix, ts, etherscanAPI)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	type respBody struct {
		Status string `json:"status"`
		Result string `json:"result"`
	}
	msg := &respBody{}
	utils.ReadJsonReaderAndSetInterface(resp.Body, msg)
	if msg.Status != "1" {
		return 0, fmt.Errorf("failed to get block num")
	}
	blockNum, err := strconv.ParseInt(msg.Result, 10, 64)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}
