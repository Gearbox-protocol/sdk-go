package core

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/exp/rand"
)

var _check bool

func check() []string {

	url := utils.GetEnvOrDefault("ETHERSCAN_PROXY_URL", "https://testnet.gearbox.foundation/etherscan/proxy")
	urls := []string{url}
	if _check {
		return urls
	}
	x, err := http.Get(urls[0] + "/health")
	if err != nil {
		return nil
	}
	if x.StatusCode != 200 {
		// log.Warnf("etherscan url %s is not working, status code: %d", urls[0], x.StatusCode)
		return nil
	}
	_check = true
	log.Info("Using etherscan proxy url", urls[0])
	return urls
}
func getEtherscanUrl(chainId int64) (urls []string, proxy bool) {
	if urls := check(); len(urls) > 0 {
		return urls, true
	}

	etherscanAPI := utils.GetEnvOrDefault("ETHERSCAN_API_KEY", "")
	if etherscanAPI == "" {
		log.Fatalf("ETHERSCAN_API_KEY can't be empty", log.GetNetworkName(chainId))
	}
	network := log.GetBaseNet(chainId)
	chainId = log.GetNetworkToChainId(network)
	urls = []string{}
	for _, api := range strings.Split(etherscanAPI, ",") {
		url := "https://api.etherscan.io/v2/api?chainid=%d&apikey=%s"
		url = fmt.Sprintf(url, chainId, api)
		urls = append(urls, url)
	}
	return urls, false
}
func getEtherscanTsUrl(chainId int64, ts int64) []string {
	urls, proxy := getEtherscanUrl(chainId)
	for i, url := range urls {
		if proxy {
			urls = append(urls, fmt.Sprintf(url+"/timestamp?chainId=%d&timestamp=%d", chainId, ts))
		} else {
			urls[i] = fmt.Sprintf(url+"&"+"module=block&action=getblocknobytime&timestamp=%d&closest=before", ts)
		}
	}
	return urls
}
func getEtherscanLogUrl(chainId int64, addr common.Address, fromBlock int64) []string {
	urls, proxy := getEtherscanUrl(chainId)
	for i, url := range urls {
		if proxy {
			urls[i] = fmt.Sprintf(url+"/logs?chainId=%d&address=%s&fromBlock=%d", chainId, addr.Hex(), fromBlock)
			continue
		}
		if fromBlock == 0 {
			urls[i] = fmt.Sprintf(url+"&"+"module=logs&action=getLogs&address=%s", addr.Hex())
		} else {
			urls[i] = fmt.Sprintf(url+"&"+"module=logs&action=getLogs&address=%s&fromBlock=%d", addr.Hex(), fromBlock)
		}
	}
	return urls
}

// dont use outside sdk-go, chainid should be of main network, not testnet
func getEtherscanBlockNum(chainId int64, ts int64) (int64, error) {
	urls := getEtherscanTsUrl(chainId, ts)
	result, err := etherscanResult(urls)
	if err != nil {
		return 0, err
	}
	blockNum, err := strconv.ParseInt(result.(string), 10, 64)
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}

var firstLogBlockCache = map[common.Address]int64{}

func GetEtherscanFirstLog(chainId int64, addr common.Address) (int64, error) {
	if block, ok := firstLogBlockCache[addr]; ok {
		return block, nil
	}
	block, err := getEtherscanFirstLog(chainId, addr)
	if err != nil {
		return 0, err
	}
	firstLogBlockCache[addr] = block
	return block, nil
}
func getEtherscanFirstLog(chainId int64, addr common.Address) (int64, error) {
	urls := getEtherscanLogUrl(chainId, addr, 0)
	result, err := etherscanResult(urls)
	if err != nil {
		return 0, err
	}
	if arr, ok := result.([]interface{}); ok && len(arr) > 0 {
		if entry, ok := arr[0].(map[string]interface{}); ok {
			block := entry["blockNumber"]
			blockNum, err := strconv.ParseInt(block.(string)[2:], 16, 64)
			if err != nil {
				return 0, err
			}
			return blockNum, nil
		}
	}
	log.Debugf("getEtherscanFirstLog: %s, result: %v", addr, result)
	return 0, fmt.Errorf("failed to get first log block num for %s", addr)
}

// range is [fromBlock, toBlock] // 0, toBLock
func GetEtherscanLogs(chainId int64, addrs []common.Address, toBlock int64, topics [][]common.Hash) ([]types.Log, error) {
	if len(topics) > 1 {
		log.Fatal("can't work with multiple topics")
	}
	allLogs := []types.Log{}
	for _, addr := range addrs {
		unfilteredtxlogs, err := getEtherscanLogs(chainId, addr, toBlock) // no topics
		if err != nil {
			return nil, err
		}
		if len(topics) == 0 {
			allLogs = append(allLogs, unfilteredtxlogs...)
		} else {
			for _, txlog := range unfilteredtxlogs {
				if len(txlog.Topics) == 0 {
					log.Info(txlog.TxHash, txlog.BlockNumber, "has no topics, skipping")
				}
				if utils.Contains(topics[0], txlog.Topics[0]) {
					allLogs = append(allLogs, txlog)
				}
			}
		}
	}
	// sort them
	sort.SliceStable(allLogs, func(i, j int) bool {
		return allLogs[i].BlockNumber < allLogs[j].BlockNumber ||
			(allLogs[i].BlockNumber == allLogs[j].BlockNumber && allLogs[i].Index < allLogs[j].Index)
	})
	return allLogs, nil
}
func getEtherscanLogs(chainId int64, addr common.Address, toBlock int64) ([]types.Log, error) {
	var fromBlock int64 = 0
	logs := []types.Log{}
	ind := 0
	for {
		url := getEtherscanLogUrl(chainId, addr, fromBlock)
		result, err := etherscanResult(url, addr)
		if err != nil {
			return nil, err
		}
		for ind := range result.([]interface{}) {
			obj := result.([]interface{})[ind].(map[string]interface{})
			if x := obj["transactionIndex"]; x.(string) == "0x" {
				result.([]interface{})[ind].(map[string]interface{})["transactionIndex"] = "0x0"
			}
			if x := obj["logIndex"]; x.(string) == "0x" {
				result.([]interface{})[ind].(map[string]interface{})["logIndex"] = "0x0"
			}
		}
		data := []types.Log{}
		err = utils.ReadJsonReaderAndSetInterface(bytes.NewBuffer(utils.ToJsonBytes(result)), &data)
		log.CheckFatal(err)
		logs = append(logs, data...)
		ind++
		// return for toBlock found, range is [fromBlock, toBlock]
		if ind, found := sort.Find(len(logs), func(i int) int {
			return int(toBlock+1) - int(logs[i].BlockNumber) // cmp is <=0 , true for block bigger  or equal than toBlock+1
		}); found {
			return logs[:ind], nil
		}
		if len(data) != 1000 {
			return logs, nil
		} else {
			last := data[len(data)-1].BlockNumber
			for logs[len(logs)-1].BlockNumber == last { // edge case when there 1000 logs with same block number, there will be index issue.
				logs = logs[:len(logs)-1] // remove last element if it has the same block number
			}
			fromBlock = int64(last)
		}
	}
}

func checkEtherscanToRepeat(err error) bool {
	if err == nil {
		return false
	}
	msgs := []string{
		"Max calls per sec rate limit reached",
		"failed to read etherscan response",
		"timeout or server too busy",
		"connection reset by peer",
	}
	for _, msg := range msgs {
		if strings.Contains(err.Error(), msg) {
			log.Info("Debug: retrying due to", err)
			time.Sleep(20 * time.Second) // wait for 20 seconds before retrying
			return true
		}
	}
	return false
}
func etherscanResult(url []string, addr ...common.Address) (interface{}, error) {
	for i := 0; i < 10; i++ {
		result, err := etherscanResultInner(url, addr...)
		if checkEtherscanToRepeat(err) {
			continue
		}
		return result, err
	}
	return nil, fmt.Errorf("failed to get etherscan result after 10 attempts for %v", addr)
}
func etherscanResultInner(urls []string, addr ...common.Address) (interface{}, error) {
	// range over urls, try to get the result
	url := urls[rand.Intn(len(urls))]

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("etherscan request failed: %s, status code: %d", url, resp.StatusCode)
	}
	type respBody struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Result  interface{} `json:"result"`
	}
	msg := &respBody{}
	str, buffer := utils.StringAndBuffer(resp.Body)
	// err = utils.ReadJsonReaderAndSetInterface(buffer, msg)

	err = utils.ReadJsonReaderAndSetInterface(buffer, msg)
	if err != nil {
		return 0, fmt.Errorf("etherscan json parsing failed: %w, %s", err, str)
	}
	if msg.Status != "1" {
		if msg.Message == "No records found" && fmt.Sprintf("%v", msg.Result) == "[]" {
			// no logs found, this is ok
			return msg.Result, nil
		}
		return 0, fmt.Errorf("%v Response failed: %s, status: %s.Result: %v", addr, msg.Message, msg.Status, msg.Result)
	}
	return msg.Result, nil

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
	log.Warn("for ts", ts, chainId, "blockNum is 0", errEtherScan, moralisErr)
	return 0
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
