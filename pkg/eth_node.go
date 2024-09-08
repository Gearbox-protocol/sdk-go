package pkg

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

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
	if err != nil && toBlock-fromBlock > 1 {
		if core.EthLogErrorCheck(err, lf.Client) {
			middle := (fromBlock + toBlock) / 2
			if middle < fromBlock {
				return nil, fmt.Errorf("middle > fromBlock %d, %d, %s", middle, fromBlock, err)
			}
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
	if tx.Type() == 2 { // types/t_dynamic_fee.go:110(effectiveGasPrice)
		if baseFee == nil {
			return tx.GasFeeCap()
		}
		tip := new(big.Int).Sub(tx.GasFeeCap(), baseFee)
		if tip.Cmp(tx.GasTipCap()) > 0 {
			tip.Set(tx.GasTipCap())
		}
		return tip.Add(tip, baseFee)
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

func (lf Node) onGetLogsError(err error, queryFrom, queryTill int64, hexAddrs []common.Address, treasuryAddrTopic []common.Hash) ([]types.Log, error) {
	if !strings.Contains(err.Error(), "exceed max topics") {
		return nil, err
	}
	batchSize := 1000
	ans := []types.Log{}

	topicsForRequest := []common.Hash{}
	for len(treasuryAddrTopic) > 0 {
		splitPoint := min(batchSize, len(treasuryAddrTopic))
		topicsForRequest, treasuryAddrTopic = treasuryAddrTopic[:splitPoint], treasuryAddrTopic[splitPoint:]
		firstLogs, err := lf.GetLogsForTransfer(queryFrom, queryTill, hexAddrs, topicsForRequest)
		if err != nil {
			return firstLogs, log.WrapErrWithLine(err)
		}
		ans = append(ans, firstLogs...)
	}
	return ans, nil
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
		return lf.onGetLogsError(err, queryFrom, queryTill, hexAddrs, treasuryAddrTopic) // in small batches, another code in this function is ignored
	}

	// from other address to treasury
	{
		newLogs, err := lf.GetLogs(queryFrom, queryTill, hexAddrs, append(topics, otherAddrTopic, treasuryAddrTopic))
		if err != nil {
			return lf.onGetLogsError(err, queryFrom, queryTill, hexAddrs, treasuryAddrTopic) // in small batches, another code in this function is ignored
		}
		logs = append(logs, newLogs...)

	}
	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i].BlockNumber < logs[j].BlockNumber ||
			(logs[i].BlockNumber == logs[j].BlockNumber && logs[i].Index < logs[j].Index)
	})
	return logs, nil
}

func getEtherscanUrl(etherscanAPI string, chainId int64, ts int64) string {
	url := "https://%s/api?module=block&action=getblocknobytime&timestamp=%d&closest=before&apikey=%s"
	var suffix string
	switch log.GetBaseNet(chainId) {
	case log.MAINNET:
		suffix = "api.etherscan.io"
	case log.ARBITRUM:
		suffix = "api.arbiscan.io"
	case log.OPTIMISM:
		suffix = "api-optimistic.etherscan.io"
	}
	url = fmt.Sprintf(url, suffix, ts, etherscanAPI)
	return url
}
// dont use outside sdk-go, chainid should be of main network, not testnet
func getEtherscanBlockNum(chainId int64, ts int64) (int64, error) {
	etherscanAPI := utils.GetEnvOrDefault(fmt.Sprintf("%s_API_KEY", log.GetNetworkName(chainId)), "")
	if etherscanAPI == "" {
		log.Fatalf("%s_API_KEY can't be empty", log.GetNetworkName(chainId))
	}

	url := getEtherscanUrl(etherscanAPI, chainId, ts)
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
// dont use outside sdk-go, chainid should be of main network, not testnet
func getMoralisBlockNum(chainId int64, ts int64) (int64, error) {
	moralis := utils.GetEnvOrDefault("MORALIS_API_KEY", "")
	if moralis == "" {
		return 0, fmt.Errorf("MORALIS_API_KEY not set")
	}
	//
	var chain string
	switch log.GetBaseNet(chainId) {
	case log.MAINNET:
		chain = "eth"
	case log.ARBITRUM:
		chain = "arbitrum"
	case log.OPTIMISM:
		chain = "optimism"
	}
	url := "https://deep-index.moralis.io/api/v2.2/dateToBlock?chain=%s&date=%d"
	url = fmt.Sprintf(url, chain, ts)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("X-API-Key", moralis)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	//s
	type respBody struct {
		Block int64 `json:"block"`
	}
	msg := &respBody{}
	utils.ReadJsonReaderAndSetInterface(resp.Body, msg)
	return msg.Block, nil
}


func GetBlockNum(ts uint64, chainId int64) int64 {
	network := log.GetBaseNet(chainId)
	chainId = log.GetNetworkToChainId(network)
	if ts == 0 {
		log.Fatalf("ts can't be 0 for %d", chainId)
	}
	//
	var errEtherScan error
	{
		for i := 0; i < 2; i++ {
			blockNum, _err := getEtherscanBlockNum(chainId, int64(ts))
			if _err == nil {
				return blockNum
			}
			time.Sleep(5 * time.Second)
			errEtherScan = _err
		}
	}
	var moralisErr error
	{
		blockNum, err := getMoralisBlockNum(chainId, int64(ts))
		if err == nil {
			return blockNum
		}
		moralisErr = err
	}
	log.Warn( "for ts", ts, chainId, "blockNum is 0", errEtherScan, moralisErr)
	return 0
}