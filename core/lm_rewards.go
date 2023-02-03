package core

import (
	"math/big"
	"sort"
)

var poolRewardv2 map[string]*big.Int = map[string]*big.Int{
	"dDAI":    big.NewInt(2283),
	"dUSDC":   big.NewInt(2283),
	"dWETH":   big.NewInt(3196),
	"dWBTC":   big.NewInt(913),
	"dwstETH": big.NewInt(1636),
}

var poolRewardv2GIP30 map[string]*big.Int = map[string]*big.Int{
	"dDAI":    big.NewInt(2283),
	"dUSDC":   big.NewInt(3101),
	"dWETH":   big.NewInt(4014),
	"dWBTC":   big.NewInt(457),
	"dwstETH": big.NewInt(0),
}

type PoolRewardSnapshot struct {
	RewardPerBlock map[string]*big.Int
	Block          int64
}

type PoolRewardSnapshots []PoolRewardSnapshot

var MainnetPoolRewards = PoolRewardSnapshots{{poolRewardv2, 15820000}, {poolRewardv2GIP30, 15977000}}
var GoerliPoolRewards = PoolRewardSnapshots{{poolRewardv2, 7694030}}

func GetRewardPerToken(chainId int64, from, to int64) []PoolRewardSnapshot {
	switch chainId {
	case 1:
		return MainnetPoolRewards.getSnapshotsInRange(from, to)
	case 5:
		return GoerliPoolRewards.getSnapshotsInRange(from, to)
	default:
		return nil
	}
}

func (changes PoolRewardSnapshots) smallOrEqualBlock(from int64) int {
	return sort.Search(len(changes), func(ind int) bool { return changes[ind].Block > from })
}
func (changes PoolRewardSnapshots) getSnapshotsInRange(from, to int64) []PoolRewardSnapshot {
	fromRewardBlock := changes.smallOrEqualBlock(from)
	toRewardBlock := changes.smallOrEqualBlock(to)
	if fromRewardBlock != 0 {
		fromRewardBlock -= 1
	}
	if toRewardBlock == len(changes) {
		toRewardBlock -= 1
	}
	return changes[fromRewardBlock:toRewardBlock]
}
