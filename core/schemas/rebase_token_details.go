package schemas

import (
	"encoding/json"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type RebaseTokenDetails struct {
	Kernel            common.Address `json:"kernel"`
	DepositValidators int64
	DepositBalance    *core.BigInt
	CLValidators      int64
	CLBalance         *core.BigInt
	RebaseDetailsForDB
}

type RebaseDetailsForDB struct {
	TotalShares *core.BigInt `gorm:"column:total_shares"`
	BlockNum    int64        `gorm:"column:blockNum" json:"-"`
	TotalETH    *core.BigInt `gorm:"column:total_eth"`
}

func (RebaseDetailsForDB) TableName() string {
	return "rebase_details"
}
func (mdl RebaseTokenDetails) GetDataForDB(blockNum int64) *RebaseDetailsForDB {
	return &RebaseDetailsForDB{
		BlockNum:    blockNum,
		TotalShares: mdl.TotalShares,
		TotalETH:    mdl.TotalETH,
	}
}

func (calc RebaseTokenDetails) GetPostTotalEther() *big.Int {
	// transient  validator eth balance
	transient := new(big.Int).Mul(big.NewInt((calc.DepositValidators-calc.CLValidators)*32), utils.GetExpInt(18))
	//
	return new(big.Int).Add(
		new(big.Int).Add(calc.DepositBalance.Convert(), calc.CLBalance.Convert()),
		transient,
	)
}

func (mdl *RebaseTokenDetails) Serialize() core.Json {
	ans := core.Json{}
	data, err := json.Marshal(mdl)
	log.CheckFatal(err)
	err = json.Unmarshal(data, &ans)
	log.CheckFatal(err)
	return ans
}
func (mdl *RebaseTokenDetails) Unserialize(input core.Json) {
	data, err := json.Marshal(input)
	log.CheckFatal(err)
	err = json.Unmarshal(data, mdl)
	log.CheckFatal(err)
}

func AdjustRebaseToken(balances map[string]core.BalanceType, stETH string, rebaseDetails *RebaseDetailsForDB) {
	if rebaseDetails == nil {
		return
	}
	if balances[stETH].IsEnabled {
		oldValues := balances[stETH]
		// (usershares*eth)/totalshares
		amt := new(big.Int).Quo(
			new(big.Int).Mul(
				rebaseDetails.TotalETH.Convert(),
				balances[core.NULL_ADDR.Hex()].BI.Convert(),
			),
			rebaseDetails.TotalShares.Convert(),
		)
		oldValues.BI = (*core.BigInt)(amt)
		balances[stETH] = oldValues
	}
}
