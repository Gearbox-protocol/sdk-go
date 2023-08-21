package schemas_v3

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type QuotaDetails struct {
	PoolQuotaKeeper string `gorm:"column:pool_quota_keeper;primaryKey"`
	BlockNum        int64  `gorm:"column:block_num;primaryKey"`
	Token           string `gorm:"column:token;primaryKey"`
	//
	Timestamp     uint64       `gorm:"column:timestamp"`
	Pool          string       `gorm:"column:pool"`
	CumQuotaIndex *core.BigInt `gorm:"column:cum_quota_index"`
	//
	Limit       *core.BigInt `gorm:"column:limit"`
	IncreaseFee uint16       `gorm:"column:increase_fee"`
	Rate        uint16       `gorm:"column:rate"`

	//
	IsDirty bool `gorm:"-"`
}

func (old QuotaDetails) Copy() *QuotaDetails {
	return &QuotaDetails{
		PoolQuotaKeeper: old.PoolQuotaKeeper,
		BlockNum:        old.BlockNum,
		Token:           old.Token,
		//
		Timestamp:     old.Timestamp,
		Pool:          old.Pool,
		CumQuotaIndex: old.CumQuotaIndex,
		//
		Limit:       old.Limit,
		IncreaseFee: old.IncreaseFee,
		Rate:        old.Rate,
	}
}

func (QuotaDetails) TableName() string {
	return "quota_details"
}

type AccountQuotaInfo struct {
	SessionId       string `gorm:"column:session_id;primaryKey"`
	BlockNum        int64  `gorm:"column:block_num;primaryKey"`
	Token           string `gorm:"column:token;primaryKey"`
	PoolQuotaKeeper string `gorm:"column:pool_quota_keeper"`
	//
	Timestamp uint64 `gorm:"column:timestamp"`
	//
	QuotaIndex *core.BigInt `gorm:"column:quota_index"`
	Quota      *core.BigInt `gorm:"column:quota"`
	Fees       *core.BigInt `gorm:"column:fees"`
	Interest   *core.BigInt `gorm:"column:interest"`
}

func (d AccountQuotaInfo) IsDisabled() bool {
	return d.Quota == nil || d.Quota.Convert().Cmp(big.NewInt(2)) < 0
}

func (old AccountQuotaInfo) Copy() *AccountQuotaInfo {
	return &AccountQuotaInfo{
		SessionId:       old.SessionId,
		BlockNum:        old.BlockNum,
		Token:           old.Token,
		PoolQuotaKeeper: old.PoolQuotaKeeper,
		//
		Timestamp: old.Timestamp,
		//
		QuotaIndex: old.QuotaIndex,
		Quota:      old.Quota,
		Fees:       old.Fees,
		Interest:   old.Interest,
	}
}

func NotNilAccountQuotaInfo() *AccountQuotaInfo {
	return &AccountQuotaInfo{
		QuotaIndex: core.NewBigInt(nil),
		Quota:      core.NewBigInt(nil),
		Fees:       core.NewBigInt(nil),
		Interest:   core.NewBigInt(nil),
	}
}

func (AccountQuotaInfo) TableName() string {
	return "account_quota_info"
}
