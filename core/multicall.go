package core

import (
	"context"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// make multicall
func MakeMultiCall(client ClientI, blockNum int64, successRequired bool, calls []multicall.Multicall2Call, params ...int) []multicall.Multicall2Result {
	contract := getMultiCallContract(client)
	opts := &bind.CallOpts{}
	if blockNum != 0 {
		opts.BlockNumber = big.NewInt(blockNum)
	}
	defaultSize := 20
	if params != nil {
		defaultSize = params[0]
		if defaultSize == 0 {
			log.Fatal("can't make multicall with batch size of 0")
		}
	}
	//

	callsLen := len(calls)
	jobsCount := int(math.Ceil(float64(callsLen) / float64(defaultSize)))
	sch := NewMulticallScheduler(jobsCount)

	callsInd := 0
	for callsInd < callsLen {
		end := callsInd + defaultSize
		if end > callsLen {
			end = callsLen
		}

		jobCalls := calls[callsInd:end]
		sch.AddJob(func() []multicall.Multicall2Result {
			tmpResult, err := contract.TryAggregate(opts, successRequired, jobCalls)
			if err != nil {
				if strings.Contains(err.Error(), "OutOfGas") || // alchemy
					strings.Contains(err.Error(), "524: A timeout occurred") || // anvil
					strings.Contains(err.Error(), "we can't execute this request") { // ankr
					tmpResult = MakeMultiCall(client, blockNum, successRequired, jobCalls, defaultSize/2)
				} else {
					log.Fatal(err)
				}
			}
			return tmpResult
		})
		callsInd += defaultSize
	}
	return sch.GetResult()
}

// / multicall contract addr
func getMultiCallAddr(chainId int64) string {
	if log.GetBaseNet(chainId) == "ARBITRUM" {
		return "0x842eC2c7D803033Edf55E478F461FC547Bc54EB2"
	}
	return "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"
}

func getMultiCallContract(client ClientI) *multicall.Multicall {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	addr := common.HexToAddress(getMultiCallAddr(chainId.Int64()))
	contract, err := multicall.NewMulticall(addr, client)
	log.CheckFatal(err)
	return contract
}

// / scheduler
type MulticallScheduler struct {
	num      int
	ch       chan MulticallJob
	curJobId int
	wg       *sync.WaitGroup
}

func NewMulticallScheduler(num int) MulticallScheduler {
	return MulticallScheduler{num: num, ch: make(chan MulticallJob, num), wg: &sync.WaitGroup{}}
}

type MulticallJob struct {
	jobId  int
	result []multicall.Multicall2Result
}

func (sch *MulticallScheduler) AddJob(fn func() []multicall.Multicall2Result) {
	sch.curJobId++
	jobId := sch.curJobId
	sch.wg.Add(1)
	go func() {
		results := fn()
		sch.ch <- MulticallJob{jobId: jobId, result: results}
		sch.wg.Done()
	}()
}

func (sch MulticallScheduler) GetResult() (ans []multicall.Multicall2Result) {
	sch.wg.Wait()
	close(sch.ch)
	jobs := map[int][]multicall.Multicall2Result{}
	for job := range sch.ch {
		jobs[job.jobId] = job.result
	}
	for i := 1; i <= sch.num; i++ {
		ans = append(ans, jobs[i]...)
	}
	return
}

// utils
func MulticallAnsBigInt(result multicall.Multicall2Result) (*big.Int, bool) {
	if result.Success && len(result.ReturnData) >= 32 {
		return new(big.Int).SetBytes(result.ReturnData[:32]), true
	}
	return big.NewInt(0), false
}
func MulticallAnsAddress(result multicall.Multicall2Result) (common.Address, bool) {
	if result.Success && len(result.ReturnData) >= 32 {
		return common.BytesToAddress(result.ReturnData[:32]), true
	}
	return NULL_ADDR, false
}

// get next multicall result
type MulticallResultIterator struct {
	results []multicall.Multicall2Result
	ind     int
}

func NewMulticallResultIterator(results []multicall.Multicall2Result) MulticallResultIterator {
	return MulticallResultIterator{
		results: results,
	}
}

func (c *MulticallResultIterator) Next() multicall.Multicall2Result {
	if c.ind == len(c.results) {
		log.Fatal("ind exceeded len of results")
	}
	ans := c.results[c.ind]
	c.ind++
	return ans
}
