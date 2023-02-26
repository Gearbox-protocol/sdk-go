package ethclient

import (
	"context"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/semaphore"
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	urls       []string
	indexInUse []bool
	clients    []*ethclient.Client
	sem        *semaphore.Weighted
	chainId    int64
}

func (c *Client) SetChainId(id int64) {
	c.chainId = id
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	urls := strings.Split(rawurl, ",")
	var l int64 = int64(len(urls))
	c := &Client{
		urls:       urls,
		indexInUse: make([]bool, l),
		sem:        semaphore.NewWeighted(l),
	}
	var err error
	var client *ethclient.Client
	for _, url := range urls {
		client, err = ethclient.Dial(url)
		if err != nil {
			return c, err
		}
		c.clients = append(c.clients, client)
	}
	return c, err
}

func (rc *Client) setInUseAndGetClient() int {
	for i, v := range rc.indexInUse {
		if !v {
			rc.indexInUse[i] = true
			return i
		}
	}
	return -1
}

func (rc *Client) returnIndex(index int) {
	rc.indexInUse[index] = false
}

func SleepFor429Error(msg string) {
	re2, _ := regexp.Compile(`backoff_seconds":(\d+)`)
	matches := re2.FindStringSubmatch(msg)
	sleepFor := 20 * time.Second
	if len(matches) == 2 {
		secs, err := strconv.ParseInt(matches[1], 10, 64)
		log.CheckFatal(err)
		sleepFor = time.Second * time.Duration(secs*2)
	}
	log.Verbosef("sleeping for %s due to 429 error.", sleepFor)
	time.Sleep(sleepFor)
}

func (rc *Client) errorHandler(err error) bool {
	if err != nil {
		if err.Error() == "execution aborted (timeout = 10s)" {
			log.Error("sleeping due to execution aborted (timeout = 10s)")
			time.Sleep(2 * time.Second)
		} else if strings.HasPrefix(err.Error(), "429") { // too many request on infura
			SleepFor429Error(err.Error())
		} else if strings.Contains(err.Error(), "Your app has exceeded its compute units per second capacity") { // too many request alchemy
			time.Sleep(20 * time.Second)
			panic(err.Error())
		} else if strings.Contains(err.Error(), "504") { // Gateway Timeout server error
			// channel/connection is not open
			time.Sleep(30 * time.Second)
		} else if strings.Contains(err.Error(), "your node is running with state pruning") {
			// this error occurs in definder state engine, when trying to get multicall data for latest blocknum
			log.Info("sleeping for 15 secs")
			//This request is not supported because your node is running with state pruning. Run with --pruning=archive.
			time.Sleep(15 * time.Second)
		} else if err.Error() == "header not found" { // makemulticall failed with this error in definder
			time.Sleep(15 * time.Second)
		} else if strings.Contains(err.Error(), "project ID does not have access to archive state") {
			log.Fatal(err)
		} else if strings.Contains(err.Error(), "EVM error FatalExternalError") { // anvil error
			log.Verbose("Trying on anvil error")
			time.Sleep(3 * time.Second)
		}
	} else {
		return false
	}
	return true
}

// copied from ethclient go-ethereum repo

func (rc *Client) Close() {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.returnIndex(ind)
	defer rc.sem.Release(1)
	client.Close()
}

func (rc *Client) ChainID(ctx context.Context) (*big.Int, error) {
	// cache
	id := atomic.LoadInt64(&(rc.chainId))
	if id != 0 {
		return big.NewInt(id), nil
	}
	// locks
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	// getting the value
	v, err := client.ChainID(ctx)
	if rc.errorHandler(err) {
		v, err = client.ChainID(ctx)
	}
	if v != nil {
		atomic.SwapInt64(&(rc.chainId), v.Int64())
	}
	return v, err
}

func (rc *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.BlockByHash(ctx, hash)
	if rc.errorHandler(err) {
		v, err = client.BlockByHash(ctx, hash)
	}
	return v, err
}

func (rc *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.BlockByNumber(ctx, number)
	if rc.errorHandler(err) {
		v, err = client.BlockByNumber(ctx, number)
	}
	return v, err
}

func (rc *Client) BlockNumber(ctx context.Context) (uint64, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.BlockNumber(ctx)
	if rc.errorHandler(err) {
		v, err = client.BlockNumber(ctx)
	}
	return v, err
}

func (rc *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.HeaderByHash(ctx, hash)
	if rc.errorHandler(err) {
		v, err = client.HeaderByHash(ctx, hash)
	}
	return v, err
}

func (rc *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.HeaderByNumber(ctx, number)
	if rc.errorHandler(err) {
		v, err = client.HeaderByNumber(ctx, number)
	}
	return v, err
}

// TransactionByHash returns the transaction with the given hash.
func (rc *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	a, b, err := client.TransactionByHash(ctx, hash)
	if rc.errorHandler(err) {
		a, b, err = client.TransactionByHash(ctx, hash)
	}
	return a, b, err
}

func (rc *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.TransactionSender(ctx, tx, block, index)
	if rc.errorHandler(err) {
		v, err = client.TransactionSender(ctx, tx, block, index)
	}
	return v, err
}

// TransactionCount returns the total number of transactions in the given block.
func (rc *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.TransactionCount(ctx, blockHash)
	if rc.errorHandler(err) {
		v, err = client.TransactionCount(ctx, blockHash)
	}
	return v, err
}

// TransactionInBlock returns a single transaction at index in the given block.
func (rc *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.TransactionInBlock(ctx, blockHash, index)
	if rc.errorHandler(err) {
		v, err = client.TransactionInBlock(ctx, blockHash, index)
	}
	return v, err
}

func (rc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.TransactionReceipt(ctx, txHash)
	if rc.errorHandler(err) {
		v, err = client.TransactionReceipt(ctx, txHash)
	}
	return v, err
}

func (rc *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.SyncProgress(ctx)
	if rc.errorHandler(err) {
		v, err = client.SyncProgress(ctx)
	}
	return v, err
}

func (rc *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.SubscribeNewHead(ctx, ch)
	if rc.errorHandler(err) {
		v, err = client.SubscribeNewHead(ctx, ch)
	}
	return v, err
}

func (rc *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.NetworkID(ctx)
	if rc.errorHandler(err) {
		v, err = client.NetworkID(ctx)
	}
	return v, err
}

func (rc *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.BalanceAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = client.BalanceAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.StorageAt(ctx, account, key, blockNumber)
	if rc.errorHandler(err) {
		v, err = client.StorageAt(ctx, account, key, blockNumber)
	}
	return v, err
}

func (rc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.CodeAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = client.CodeAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.NonceAt(ctx, account, blockNumber)
	if rc.errorHandler(err) {
		v, err = client.NonceAt(ctx, account, blockNumber)
	}
	return v, err
}

func (rc *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.FilterLogs(ctx, q)
	if rc.errorHandler(err) {
		v, err = client.FilterLogs(ctx, q)
	}
	return v, err
}

func (rc *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.SubscribeFilterLogs(ctx, q, ch)
	if rc.errorHandler(err) {
		v, err = client.SubscribeFilterLogs(ctx, q, ch)
	}
	return v, err
}

func (rc *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingBalanceAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = client.PendingBalanceAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingStorageAt(ctx, account, key)
	if rc.errorHandler(err) {
		v, err = client.PendingStorageAt(ctx, account, key)
	}
	return v, err
}

func (rc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingCodeAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = client.PendingCodeAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingNonceAt(ctx, account)
	if rc.errorHandler(err) {
		v, err = client.PendingNonceAt(ctx, account)
	}
	return v, err
}

func (rc *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingTransactionCount(ctx)
	if rc.errorHandler(err) {
		v, err = client.PendingTransactionCount(ctx)
	}
	return v, err
}

func (rc *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.CallContract(ctx, msg, blockNumber)
	if rc.errorHandler(err) {
		v, err = client.CallContract(ctx, msg, blockNumber)
	}
	return v, err
}

func (rc *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.PendingCallContract(ctx, msg)
	if rc.errorHandler(err) {
		v, err = client.PendingCallContract(ctx, msg)
	}
	return v, err
}

func (rc *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.SuggestGasPrice(ctx)
	if rc.errorHandler(err) {
		v, err = client.SuggestGasPrice(ctx)
	}
	return v, err
}

func (rc *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.SuggestGasTipCap(ctx)
	if rc.errorHandler(err) {
		v, err = client.SuggestGasTipCap(ctx)
	}
	return v, err
}

func (rc *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	v, err := client.EstimateGas(ctx, msg)
	if rc.errorHandler(err) {
		v, err = client.EstimateGas(ctx, msg)
	}
	return v, err
}

func (rc *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	rc.sem.Acquire(context.TODO(), 1)
	ind := rc.setInUseAndGetClient()
	client := rc.clients[ind]
	defer rc.sem.Release(1)
	defer rc.returnIndex(ind)
	//
	err := client.SendTransaction(ctx, tx)
	if rc.errorHandler(err) {
		err = client.SendTransaction(ctx, tx)
	}
	return err
}
