package schemas_v3

import "math/big"

type TokenLTRamp struct {
	BlockNum      int64  `gorm:"column:block_num;primaryKey"`
	CreditManager string `gorm:"column:credit_manager;primaryKey"`
	Token         string `gorm:"column:token;primaryKey"`
	LtInitial     uint16 `gorm:"column:lt_initial"`
	LtFinal       uint16 `gorm:"column:lt_final"`
	RampStart     int64  `gorm:"column:ramp_start"`
	RampEnd       int64  `gorm:"column:ramp_end"`
}

func (TokenLTRamp) TableName() string {
	return "token_ltramp"
}

func (details TokenLTRamp) GetLTForTs(ts int64) *big.Int {
	return big.NewInt(int64(details.getLTForTs(ts)))
}
func (details TokenLTRamp) getLTForTs(ts int64) uint16 {
	if details.RampStart > ts {
		return details.LtInitial
	}
	if details.RampEnd < ts {
		return details.LtFinal
	}
	num := int64(details.LtFinal)*(ts-details.RampStart) + int64(details.LtInitial)*(details.RampEnd-ts)
	return uint16(num / (details.RampEnd - details.RampStart))
}
