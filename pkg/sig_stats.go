package pkg

import (
	"fmt"
	"io"
	"math/big"
	"reflect"
	"sort"
	"time"

	"github.com/Gearbox-protocol/sdk-go/utils"
)

type SigDetails struct {
	days map[int64]*SigDetailsByDay
}

func (x SigDetails) GetDayStats() (ans []SigDetailsByDay) {
	keys := reflect.ValueOf(x.days).MapKeys()
	keysOrder := func(i, j int) bool { return (keys[i].Interface().(int64)) < keys[j].Interface().(int64) }
	sort.Slice(keys, keysOrder)

	// process map in key-sorted order
	for _, key := range keys {
		dateTs := key.Interface().(int64)
		details := x.days[key.Interface().(int64)]
		ans = append(ans, SigDetailsByDay{
			day:     utils.TimeToDate(time.Unix(dateTs, 0)),
			Count:   details.Count,
			EthUsed: new(big.Int).Quo(details.EthUsed, big.NewInt(details.Count)),
		})
	}
	return
}

type SigDetailsByDay struct {
	EthUsed *big.Int
	Count   int64
	day     string
}

func NewSigStatsHandler(sigsToMethodName map[string]string) *SigStatsHandler {
	return &SigStatsHandler{
		sigsToMethodName: sigsToMethodName,
		stats:            make(map[string]*SigDetails, len(sigsToMethodName)),
	}
}
func (f *SigStatsHandler) Add(sig string, tx EtherScanCallInput) {
	if _, ok := f.stats[sig]; !ok {
		f.stats[sig] = &SigDetails{
			days: make(map[int64]*SigDetailsByDay),
		}
	}
	ts := utils.TimeToDateEndTs(time.Unix(tx.TimeStamp, 0))
	f.stats[sig].days[ts].Count++
	f.stats[sig].days[ts].EthUsed = new(big.Int).Add(
		utils.NotNilBigInt(f.stats[sig].days[ts].EthUsed),
		new(big.Int).Mul(utils.StringToInt(tx.GasPrice), utils.StringToInt(tx.GasUsed)),
	)
}
func (f *SigStatsHandler) Print(steam io.Writer) {
	for sig, details := range f.stats {
		fmt.Fprintln(steam, "##############")
		fmt.Fprintln(steam, "For method", f.sigsToMethodName[sig])
		fmt.Fprintln(steam, "##############")
		for _, stat := range details.GetDayStats() {
			fmt.Fprintln(steam, stat.day, stat.Count, stat.EthUsed)
		}
	}
}

//
type SigStatsHandler struct {
	stats            map[string]*SigDetails
	sigsToMethodName map[string]string
}

func (f SigStatsHandler) AllowedSig(sig string) bool {
	_, ok := f.sigsToMethodName[sig]
	return ok
}
