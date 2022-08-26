package pauser

import (
	"context"
	"fmt"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/pauseMulticall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Pauser struct {
	isDone bool
	mu     *sync.Mutex
	*core.BaseTransactor
	flashBots *core.FlashbotsTransactor
	//
	multicall *pauseMulticall.PauseMulticall
}

func NewPauser(transactor *core.BaseTransactor, privateKey, flashbotsRelay, pauserMCAddr string, client core.ClientI) *Pauser {
	c := &Pauser{
		BaseTransactor: core.NewBaseTransactor(privateKey, client),
		mu:             &sync.Mutex{},
		flashBots:      core.NewFlashbotsTransactor(flashbotsRelay),
	}
	// don't send tx , while creating it
	c.Topts.NoSend = true
	//
	multicall, err := pauseMulticall.NewPauseMulticall(common.HexToAddress(pauserMCAddr), c.Client)
	log.CheckFatal(err)
	c.multicall = multicall
	return c
}

func (p *Pauser) getIsDone() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.isDone
}

func (p *Pauser) setIsDone(v bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.isDone = v
}

func (p *Pauser) Pause(blockNum int64, contractsToPause []string) {
	log.Info("Pausing for", blockNum)
	if p.getIsDone() {
		return
	}
	// flashbotsrpc
	n := 4
	ch := make(chan error, n)
	for i := 0; i < n; i++ {
		go p.pauseMulticallFlashbots(ch, blockNum+int64(i))
	}
	// rpc
	ch2 := make(chan error, 1)
	go p.pauseMulticallRPC(ch2)
	//

	for i := 0; i < n; i++ {
		select {
		case err := <-ch:
			if err == nil {
				p.setIsDone(true)
				return
			} else {
				log.Error(err)
			}
		case err := <-ch2:
			if err == nil {
				p.setIsDone(true)
				return
			} else {
				log.Error(err)
			}
		}
	}
}

func (p *Pauser) pauseMulticallFlashbots(ch chan<- error, blockNum int64) {
	tx, err := p.multicall.PauseAllContracts(p.Topts)
	if err != nil {
		ch <- err
	}

	if err := p.flashBots.Send(blockNum, []*types.Transaction{tx}); err != nil {
		ch <- err
	}
	ch <- nil
}
func (p *Pauser) pauseMulticallRPC(ch chan error) {
	tx, err := p.multicall.PauseAllContracts(p.Topts)
	if err != nil {
		ch <- fmt.Errorf("tx building failed pausing: %s", err.Error())
	}
	if err := p.Client.SendTransaction(context.TODO(), tx); err != nil {
		ch <- err
	}
	_, err = p.WaitForTx("Pausing contracts", tx)
	ch <- err
}
