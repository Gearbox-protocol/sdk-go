package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
)

type GearBalance struct {
	Balance *core.BigInt `gorm:"column:balance"`
	Updated bool         `gorm:"-"`
	User    string       `gorm:"column:user_address;primaryKey"`
}

func (GearBalance) TableName() string {
	return "gear_balances"
}
