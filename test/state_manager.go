package test

import (
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

// /
type TestState struct {
	OracleMgr OracleMgr `json:"oracles"`
	// sig to target to data
	OtherCallsGlobal OtherCalls `json:"otherCalls"`
}

func (state *TestState) Merge(other TestState) {
	state.OracleMgr.Merge(other.OracleMgr)
	if state.OtherCallsGlobal == nil {
		state.OtherCallsGlobal = make(OtherCalls)
	}
	for sig, calls := range other.OtherCallsGlobal {
		if state.OtherCallsGlobal[sig] == nil {
			state.OtherCallsGlobal[sig] = calls
		} else {
			x := state.OtherCallsGlobal[sig]
			x.Merge(calls)
		}
	}
}

// /
type OtherCalls map[string]SigCalls

// /
type SigCalls map[common.Address]string

func (calls *SigCalls) Merge(other SigCalls) {
	for target, data := range other {
		(*calls)[target] = data
	}
}

// //
type StateManager struct {
	OracleMgr OracleMgr
	// block sig to target to data
	OtherCalls map[string]map[common.Address][]*NumAndData
}
type NumAndData struct {
	BlockNum int64
	Data     string
}

func NewStateManager() *StateManager {
	return &StateManager{
		OtherCalls: make(map[string]map[common.Address][]*NumAndData),
	}
}

func (state *StateManager) addState(s *TestState, calls map[int64]OtherCalls) {
	for blockNum, cs := range calls {
		state.addForBlockNum(blockNum, cs)
	}
	//
	state.OracleMgr = s.OracleMgr
	state.addForBlockNum(0, s.OtherCallsGlobal)

	for _, data := range state.OtherCalls {
		for _, list := range data {
			sort.Slice(list, func(i, j int) bool { return list[i].BlockNum < list[j].BlockNum })
		}
	}
}

func (state *StateManager) addForBlockNum(blockNum int64, cs OtherCalls) {
	for sig, dataWithTarget := range cs {
		for target, data := range dataWithTarget {
			if state.OtherCalls[sig] == nil {
				state.OtherCalls[sig] = make(map[common.Address][]*NumAndData)
			}
			state.OtherCalls[sig][target] = append(state.OtherCalls[sig][target], &NumAndData{
				BlockNum: blockNum,
				Data:     data,
			})
		}
	}
}

func (state *StateManager) GetOtherCall(blockNum int64, sig string, target common.Address) string {
	// log.Info(utils.ToJson(state.OtherCallsByBlock))
	// }
	if state.OtherCalls[sig] == nil || state.OtherCalls[sig][target] == nil {
		return ""
	}
	entries := state.OtherCalls[sig][target]
	ind := sort.Search(len(entries), func(i int) bool {
		return entries[i].BlockNum > blockNum
	})
	if ind != 0 {
		return entries[ind-1].Data
	}
	return ""
}
