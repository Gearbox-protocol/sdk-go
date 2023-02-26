package ethclient

import (
	"context"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	clients []*MutextedClient
	sem     *semaphore.Weighted
	chainId int64
}

func (c *Client) SetChainId(id int64) {
	c.chainId = id
}

type MutextedClient struct {
	mu     *sync.Mutex
	client *ethclient.Client
}

func NewMutextedClient(url string) *MutextedClient {
	client, err := ethclient.Dial(url)
	log.CheckFatal(err)
	return &MutextedClient{
		client: client,
		mu:     &sync.Mutex{},
	}
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	urls := strings.Split(rawurl, ",")
	l := int64(len(urls))
	c := &Client{
		clients: make([]*MutextedClient, l),
		sem:     semaphore.NewWeighted(l),
	}
	for i, url := range urls {
		c.clients[i] = NewMutextedClient(url)
	}
	return c, nil
}

func sleepFor429Error(msg string) {
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

func errorHandler(err error) bool {
	if err != nil {
		if err.Error() == "execution aborted (timeout = 10s)" {
			log.Error("sleeping due to execution aborted (timeout = 10s)")
			time.Sleep(2 * time.Second)
		} else if strings.HasPrefix(err.Error(), "429") { // too many request on infura
			sleepFor429Error(err.Error())
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
		} else {
			return false
		}
	}
	return true
}

// copied from ethclient go-ethereum repo

func (rc *Client) Close() {
	client, unlock := rc.getClient()
	defer unlock()
	client.Close()
}

func getDataViaRetry[T any](getData func() (T, error)) (T, error) {
	data, err := getData()
	if err != nil && errorHandler(err) {
		data, err = getData()
	}
	return data, err
}

func (rc Client) getClient() (*ethclient.Client, func()) {
	rc.sem.Acquire(context.TODO(), 1)
	for _, muclient := range rc.clients {
		if muclient.mu.TryLock() {
			mu := muclient.mu
			return muclient.client, func() {
				mu.Unlock()
				rc.sem.Release(1)
			}
		}
	}
	log.Fatal("no client available")
	return nil, nil
}
func (rc *Client) ChainID(ctx context.Context) (*big.Int, error) {
	// cache
	id := atomic.LoadInt64(&(rc.chainId))
	if id != 0 {
		return big.NewInt(id), nil
	}
	// locks
	client, unlock := rc.getClient()
	defer unlock()
	v, err := getDataViaRetry(func() (*big.Int, error) { return client.ChainID(ctx) })
	//
	if v != nil {
		atomic.SwapInt64(&(rc.chainId), v.Int64())
	}
	return v, err
}

func (rc *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Block, error) { return client.BlockByHash(ctx, hash) })
}

func (rc *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Block, error) { return client.BlockByNumber(ctx, number) })
}

func (rc *Client) BlockNumber(ctx context.Context) (uint64, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint64, error) { return client.BlockNumber(ctx) })
}

func (rc *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Header, error) { return client.HeaderByHash(ctx, hash) })
}

func (rc *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Header, error) { return client.HeaderByNumber(ctx, number) })
}

// TransactionByHash returns the transaction with the given hash.
func (rc *Client) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	client, unlock := rc.getClient()
	defer unlock()
	//
	tx, pending, err := client.TransactionByHash(ctx, hash)
	if errorHandler(err) {
		tx, pending, err = client.TransactionByHash(ctx, hash)
	}
	return tx, pending, err
}

func (rc *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (common.Address, error) { return client.TransactionSender(ctx, tx, block, index) })
}

// TransactionCount returns the total number of transactions in the given block.
func (rc *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint, error) { return client.TransactionCount(ctx, blockHash) })
}

// TransactionInBlock returns a single transaction at index in the given block.
func (rc *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Transaction, error) { return client.TransactionInBlock(ctx, blockHash, index) })
}

func (rc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*types.Receipt, error) { return client.TransactionReceipt(ctx, txHash) })

}

func (rc *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*ethereum.SyncProgress, error) { return client.SyncProgress(ctx) })

}

func (rc *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (ethereum.Subscription, error) { return client.SubscribeNewHead(ctx, ch) })
}

func (rc *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*big.Int, error) { return client.NetworkID(ctx) })
}

func (rc *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*big.Int, error) { return client.BalanceAt(ctx, account, blockNumber) })
}

func (rc *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.StorageAt(ctx, account, key, blockNumber) })
}

func (rc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.CodeAt(ctx, account, blockNumber) })
}

func (rc *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint64, error) { return client.NonceAt(ctx, account, blockNumber) })
}

func (rc *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]types.Log, error) { return client.FilterLogs(ctx, q) })
}

func (rc *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (ethereum.Subscription, error) { return client.SubscribeFilterLogs(ctx, q, ch) })
}

func (rc *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*big.Int, error) { return client.PendingBalanceAt(ctx, account) })
}

func (rc *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.PendingStorageAt(ctx, account, key) })
}

func (rc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.PendingCodeAt(ctx, account) })
}

func (rc *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint64, error) { return client.PendingNonceAt(ctx, account) })
}

func (rc *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint, error) { return client.PendingTransactionCount(ctx) })
}

func (rc *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.CallContract(ctx, msg, blockNumber) })
}

func (rc *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() ([]byte, error) { return client.PendingCallContract(ctx, msg) })
}

func (rc *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*big.Int, error) { return client.SuggestGasPrice(ctx) })
}

func (rc *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (*big.Int, error) { return client.SuggestGasTipCap(ctx) })
}

func (rc *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	client, unlock := rc.getClient()
	defer unlock()
	return getDataViaRetry(func() (uint64, error) { return client.EstimateGas(ctx, msg) })
}

func (rc *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	client, unlock := rc.getClient()
	defer unlock()
	err := client.SendTransaction(ctx, tx)
	if errorHandler(err) {
		err = client.SendTransaction(ctx, tx)
	}
	return err
}
