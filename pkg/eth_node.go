package pkg

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/addrProviderv310"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
)

type Node struct {
	Client  core.ClientI
	chainId int64
}

func (lf Node) GetLogs(fromBlock, toBlock int64, addrs []common.Address, topics [][]common.Hash, etherscanOnly ...bool) ([]types.Log, error) {
	if fromBlock == 0 {
		var minBlock int64 = math.MaxInt64
		baseChainId := core.GetBaseChainId(lf.Client)
		for _, addr := range addrs {
			block, err := core.GetEtherscanFirstLog(baseChainId, addr)
			if err != nil {
				var newBlock int64 = 1
				if lf.GetLatestBlockNumber()-100_000 > 0 {
					newBlock = lf.GetLatestBlockNumber() - 100_000
				}
				// log.Warnf("GetLogs: GetEtherscanFirstLog for %s error: %s. Set to latest-10k %d", addr.Hex(), err, newBlock)
				block = newBlock
			}
			minBlock = utils.Min(minBlock, block)
		}
		fromBlock = minBlock
		// log.Info("GetLogs: fromBlock is 0, set to", fromBlock, addrs)
		//
		if len(etherscanOnly) > 0 && etherscanOnly[0] {
			baseChainId := core.GetBaseChainId(lf.Client)
			logs, err := core.GetEtherscanLogs(baseChainId, addrs, toBlock, topics)
			log.Info("GetLogs: logs using etherscan for single addr with no topic for", fromBlock, toBlock, addrs, len(logs), " fetched")
			if !(len(logs) == 0 && err == nil) { // when there are no logs, and no error, this means check on rpc for logs
				return logs, err
			}
		}
	}
	logs, err := lf.getLogs(fromBlock, toBlock, addrs, topics)
	log.Debugf("GetLogs: fromBlock %d, toBlock %d from rpc. %d", fromBlock, toBlock, len(logs))
	return logs, err
}

func (lf Node) getLogs(fromBlock, toBlock int64, addrs []common.Address, topics [][]common.Hash) ([]types.Log, error) {
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

// contract key hash to contract name
func initv310ContractHashMap(client core.ClientI, addressProvider common.Address) map[common.Hash]string {
	con, err := addrProviderv310.NewAddrProviderv310(addressProvider, client)
	log.CheckFatal(err)
	contracts, err := con.GetAllEntries(nil) // hash and contract address
	log.CheckFatal(err)

	node := Node{Client: client}
	txLogs, err := node.GetLogs(0, node.GetLatestBlockNumber(), []common.Address{addressProvider}, [][]common.Hash{})
	log.CheckFatal(err)
	log.Info("Init gearbox addresses: count", len(txLogs))
	return createMap(contracts, txLogs)
}

func createMap(contracts []addrProviderv310.AddressProviderEntry, txLogs []types.Log) map[common.Hash]string {
	var addrv310, _ = addrProviderv310.NewAddrProviderv310(core.NULL_ADDR, nil)
	addrToContractHash := map[common.Address]common.Hash{}
	for _, txLog := range txLogs {
		if core.Topic("SetAddress(bytes32,uint256,address)") == txLog.Topics[0] {
			event, err := addrv310.ParseSetAddress(txLog)
			log.CheckFatal(err)
			addrToContractHash[event.Value] = event.Key
		}
	}
	log.Infof("getAllContractsMap: %d, getAllContracts: %d", len(addrToContractHash), len(contracts))
	//

	hashToContractName := map[common.Hash]string{
		common.HexToHash("0xd833762afa5298dc6d804f571e7b934fd54a0f7b926c90ff79912e7379744b31"): "ADDRESS_PROVIDER",
	}
	for _, contract := range contracts {
		hash := addrToContractHash[contract.Value]
		// contractName := contract.Key // string
		contractName := strings.Trim(string(contract.Key[:]), "\x00") // byte32
		hashToContractName[hash] = contractName
	}
	return hashToContractName
}
