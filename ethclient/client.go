package ethclient

import (
	"context"
	"math/big"
	"math/rand"
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
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	clients []*MutextedClient
	chainId int64
}

func (c *Client) SetChainId(id int64) {
	c.chainId = id
}

type MutextedClient struct {
	mu            *sync.Mutex
	client        *ethclient.Client
	_lockedTillTs int64
}

func (mc *MutextedClient) addSleepForSecs(sec int64) {
	atomic.SwapInt64(&mc._lockedTillTs, time.Now().Add(time.Second*time.Duration(sec)).Unix())
}

func (mc MutextedClient) available(curTs int64) bool {
	if mc._lockedTillTs > curTs {
		return false
	}
	return mc.mu.TryLock()
}

func (mc MutextedClient) wait() {
	mc.mu.Lock()
	sleepFor := time.Unix(mc._lockedTillTs, 0).Sub(time.Now())
	if sleepFor > 0 {
		time.Sleep(sleepFor)
	}
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
	}
	for i, url := range urls {
		c.clients[i] = NewMutextedClient(url)
	}
	return c, nil
}

func sleepFor429Error(msg string) int64 {
	re2, _ := regexp.Compile(`backoff_seconds":(\d+)`)
	matches := re2.FindStringSubmatch(msg)
	var sleepFor int64 = 20
	if len(matches) == 2 {
		secs, err := strconv.ParseInt(matches[1], 10, 64)
		log.CheckFatal(err)
		sleepFor = secs
	}
	return sleepFor
}

func errorHandler(err error, mc *MutextedClient) bool {
	if err != nil {
		// fmt.Println("retrying")
		if err.Error() == "execution aborted (timeout = 10s)" {
			log.Error("sleeping due to execution aborted (timeout = 10s)")
			mc.addSleepForSecs(2)
		} else if strings.HasPrefix(err.Error(), "429") { // too many request on infura
			mc.addSleepForSecs(sleepFor429Error(err.Error()))
		} else if strings.Contains(err.Error(), "504") { // Gateway Timeout server error
			// channel/connection is not open
			mc.addSleepForSecs(30)
		} else if strings.Contains(err.Error(), "your node is running with state pruning") {
			// this error occurs in definder state engine, when trying to get multicall data for latest blocknum
			log.Info("sleeping for 15 secs")
			mc.addSleepForSecs(15)
			//This request is not supported because your node is running with state pruning. Run with --pruning=archive.
		} else if err.Error() == "header not found" { // makemulticall failed with this error in definder
			mc.addSleepForSecs(15)
		} else if strings.Contains(err.Error(), "EVM error FatalExternalError") { // anvil error
			log.Verbose("Trying on anvil error")
			mc.addSleepForSecs(3)
		} else if strings.Contains(err.Error(), "project ID does not have access to archive state") {
			log.Fatal(err)
		} else if strings.Contains(err.Error(), "Your app has exceeded its compute units per second capacity") { // too many request alchemy
			log.Fatal(err)
		} else {
			return false
		}
	}
	return true
}

// copied from ethclient go-ethereum repo

func (rc *Client) Close() {
	mc := rc.getClient()
	mc.client.Close()
}

func getDataViaRetry[T any](wrapperClient *Client, mc *MutextedClient, getData func(c *ethclient.Client) (T, error)) (T, error) {
	data, err := getData(mc.client)
	if err != nil && errorHandler(err, mc) {
		mc.mu.Unlock()
		mc = wrapperClient.getClient()
		data, err = getData(mc.client)
	}
	mc.mu.Unlock()
	return data, err
}

func (rc Client) getClient() *MutextedClient {
	curTs := time.Now().Unix()
	l := len(rc.clients)
	//
	start := rand.Intn(l)
	for i := 0; i < l; i++ {
		ind := (i + start) % l
		muclient := rc.clients[ind]
		if muclient.available(curTs) {
			// fmt.Println(ind)
			return muclient
		}
	}

	muclient := rc.clients[start]
	muclient.wait()
	return muclient
}
func (rc *Client) ChainID(ctx context.Context) (*big.Int, error) {
	// cache
	id := atomic.LoadInt64(&(rc.chainId))
	if id != 0 {
		return big.NewInt(id), nil
	}
	// locks
	mc := rc.getClient()
	v, err := getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.ChainID(ctx) })
	//
	if v != nil {
		atomic.SwapInt64(&(rc.chainId), v.Int64())
	}
	return v, err
}

func (wrapClient *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	mc := wrapClient.getClient()
	return getDataViaRetry(wrapClient, mc, func(c *ethclient.Client) (*types.Block, error) { return c.BlockByHash(ctx, hash) })
}

func (rc *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*types.Block, error) { return c.BlockByNumber(ctx, number) })
}

func (rc *Client) BlockNumber(ctx context.Context) (uint64, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint64, error) { return c.BlockNumber(ctx) })
}

func (rc *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*types.Header, error) { return c.HeaderByHash(ctx, hash) })
}

func (rc *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*types.Header, error) { return c.HeaderByNumber(ctx, number) })
}

// TransactionByHash returns the transaction with the given hash.
func (rc *Client) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	mc := rc.getClient()
	//
	tx, pending, err := mc.client.TransactionByHash(ctx, hash)
	if errorHandler(err, mc) {
		mc.mu.Unlock()
		mc = rc.getClient()
		tx, pending, err = mc.client.TransactionByHash(ctx, hash)
	}
	mc.mu.Unlock()
	return tx, pending, err
}

func (rc *Client) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (common.Address, error) {
		return c.TransactionSender(ctx, tx, block, index)
	})
}

// TransactionCount returns the total number of transactions in the given block.
func (rc *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint, error) { return c.TransactionCount(ctx, blockHash) })
}

// TransactionInBlock returns a single transaction at index in the given block.
func (rc *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*types.Transaction, error) {
		return c.TransactionInBlock(ctx, blockHash, index)
	})
}

func (rc *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*types.Receipt, error) { return c.TransactionReceipt(ctx, txHash) })

}

func (rc *Client) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*ethereum.SyncProgress, error) { return c.SyncProgress(ctx) })

}

func (rc *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (ethereum.Subscription, error) { return c.SubscribeNewHead(ctx, ch) })
}

func (rc *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.NetworkID(ctx) })
}

func (rc *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.BalanceAt(ctx, account, blockNumber) })
}

func (rc *Client) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.StorageAt(ctx, account, key, blockNumber) })
}

func (rc *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.CodeAt(ctx, account, blockNumber) })
}

func (rc *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint64, error) { return c.NonceAt(ctx, account, blockNumber) })
}

func (rc *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]types.Log, error) { return c.FilterLogs(ctx, q) })
}

func (rc *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (ethereum.Subscription, error) { return c.SubscribeFilterLogs(ctx, q, ch) })
}

func (rc *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.PendingBalanceAt(ctx, account) })
}

func (rc *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.PendingStorageAt(ctx, account, key) })
}

func (rc *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.PendingCodeAt(ctx, account) })
}

func (rc *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint64, error) { return c.PendingNonceAt(ctx, account) })
}

func (rc *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint, error) { return c.PendingTransactionCount(ctx) })
}

func (rc *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.CallContract(ctx, msg, blockNumber) })
}

func (rc *Client) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) ([]byte, error) { return c.PendingCallContract(ctx, msg) })
}

func (rc *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.SuggestGasPrice(ctx) })
}

func (rc *Client) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (*big.Int, error) { return c.SuggestGasTipCap(ctx) })
}

func (rc *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	mc := rc.getClient()
	return getDataViaRetry(rc, mc, func(c *ethclient.Client) (uint64, error) { return c.EstimateGas(ctx, msg) })
}

func (rc *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	mc := rc.getClient()
	err := mc.client.SendTransaction(ctx, tx)
	if errorHandler(err, mc) {
		mc.mu.Unlock()
		mc = rc.getClient()
		err = mc.client.SendTransaction(ctx, tx)
	}
	return err
}
