package pkg

import (
	"context"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"

	// "github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

func getHeaders(wsProviders []string) (chan *types.Header, event.Subscription, core.ClientI) {
	for ind, str := range wsProviders {
		wsClient, err := ethclient.Dial(str)
		log.CheckFatal(err)
		headers := make(chan *types.Header)
		sub := event.ResubscribeErr(500*time.Millisecond, func(ctx context.Context, err error) (event.Subscription, error) {
			log.Debugf("resubscribing due to %v", err)
			return wsClient.SubscribeNewHead(ctx, headers)
		})
		timer := time.NewTimer(time.Minute)
		select {
		case <-headers:
			return headers, sub, wsClient
		case <-timer.C:
			log.Warnf("timeout on %d wsprovider's header can't get anything for 1min, trying with next wsprovider", ind)
		case err := <-sub.Err():
			log.Info(err)
		}
	}
	return nil, nil, nil
}
func WsFetchBlockNumFrom(_ctx context.Context, syncedTill int64, wsProviders []string, fn func(fromBlock int64, toBlock int64, tots uint64), interval int64, dontGetTs ...bool) {
	//
	headers, sub, wsClient := getHeaders(wsProviders)
	for {
		select {
		case <-_ctx.Done():
			return
		case err := <-sub.Err():
			log.Info(err)
		case header := <-headers:
			latestBlockNum := header.Number.Int64()
			log.Debug("Block from ws", latestBlockNum)

			// solves 3 cases
			// - syncs all the block till latest blockNum on restart of application
			// - syncs all blocks till latest blockNum when there is an error in websocket and there is resubscription of ws
			// - handles rollback and ignores all the block until `syncedTill`
			if syncedTill != 0 {
				for nextBlock := syncedTill + interval; nextBlock <= latestBlockNum; nextBlock += interval {
					var ts uint64
					if !(len(dontGetTs) != 0 && dontGetTs[0]) {
						if latestBlockNum == nextBlock { // for latest block, we get timestamp from ws subscrption
							ts = header.Time
						} else {
							header, err := wsClient.HeaderByNumber(context.Background(), big.NewInt(nextBlock))
							log.CheckFatal(err)
							ts = header.Time
						}
					}
					fn(nextBlock-interval+1, nextBlock, ts)
					syncedTill = nextBlock
				}
			} else {
				syncedTill = latestBlockNum
			}
		}
	}
}
