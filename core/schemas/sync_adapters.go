package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
)

type SyncAdapterSchema struct {
	*Contract
	LastSync         int64            `gorm:"column:last_sync" json:"lastSync"`
	Details          core.Json        `gorm:"column:details" json:"details"`
	Error            string           `gorm:"column:error" json:"error"`
	BlockToDisableOn int64            `gorm:"column:disabled_at" json:"disabled_at"`
	V                core.VersionType `gorm:"column:version" json:"version"`
}

func (SyncAdapterSchema) TableName() string {
	return "sync_adapters"
}
