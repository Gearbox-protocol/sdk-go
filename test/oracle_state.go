package test

import (
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

// oracle to blockNum and feed
type OracleMgr map[common.Address][]OracleState

type OracleState struct {
	Oracle   string `json:"oracle"`
	BlockNum int64  `json:"block"`
	Feed     string `json:"feed"`
}

func (s *OracleMgr) GetIndex(oracle common.Address, blockNum int64) (index int) {
	feeds := (*s)[oracle]
	for i := 0; i < len(feeds); i++ {
		if feeds[i].BlockNum <= blockNum {
			index = i
		} else {
			break
		}
	}
	return index
}
func (s *OracleMgr) Merge(other OracleMgr) {
	for oracle, feeds := range other {
		(*s)[oracle] = append((*s)[oracle], feeds...)
		x := (*s)[oracle]
		sort.Slice(x, func(i, j int) bool { return x[i].BlockNum < x[j].BlockNum })
	}
	for _, feeds := range *s {
		sort.Slice(feeds, func(i, j int) bool { return feeds[i].BlockNum < feeds[j].BlockNum })
	}
}
func (s OracleMgr) GetState(oracle common.Address, index int) OracleState {
	return s[oracle][index]
}
