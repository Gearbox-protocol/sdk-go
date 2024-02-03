package pkg

import (
	"context"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

func WsFetchBlockNumFrom(_ctx context.Context, syncedTill int64, wsProvider string, fn func(_ int64, _ uint64), dontGetTs ...bool) {
	//
	wsClient, err := ethclient.Dial(wsProvider)
	log.CheckFatal(err)
	headers := make(chan *types.Header)
	sub := event.ResubscribeErr(500*time.Millisecond, func(ctx context.Context, err error) (event.Subscription, error) {
		log.Debugf("resubscribing due to %v", err)
		return wsClient.SubscribeNewHead(ctx, headers)
	})
	for {
		select {
		case <-_ctx.Done():
			return
		case err := <-sub.Err():
			log.Info(err)
		case header := <-headers:
			latestBlockNum := header.Number.Int64()
			log.Info("Block from ws", latestBlockNum)

			// solves 3 cases
			// - syncs all the block till latest blockNum on restart of application
			// - syncs all blocks till latest blockNum when there is an error in websocket and there is resubscription of ws
			// - handles rollback and ignores all the block until `syncedTill`
			if syncedTill != 0 {
				for nextBlock := syncedTill + 1; nextBlock <= latestBlockNum; nextBlock++ {
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
					fn(nextBlock, ts)
					syncedTill = nextBlock
				}
			} else {
				syncedTill = latestBlockNum
			}
		}
	}
}
