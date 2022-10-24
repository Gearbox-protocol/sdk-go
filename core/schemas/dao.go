package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
)

type DAOOperation struct {
	LogID       uint       `gorm:"column:log_id;primaryKey"`
	TxHash      string     `gorm:"column:tx_hash"`
	BlockNumber int64      `gorm:"column:block_num;primaryKey"`
	Contract    string     `gorm:"column:contract"`
	Type        uint       `gorm:"column:type"`
	Args        *core.Json `gorm:"column:args"`
}

type TreasuryTransfer struct {
	Amount              *core.BigInt `gorm:"column:amount"`
	Token               string       `gorm:"column:token"`
	LogID               uint         `gorm:"column:log_id;primaryKey"`
	BlockNum            int64        `gorm:"column:block_num;primaryKey"`
	OperationalTransfer bool         `gorm:"column:operational_transfer"`
}

type BlockDate struct {
	// Date      string `gorm:"column:date"`
	BlockNum  int64 `gorm:"column:block_num"`
	Timestamp int64 `gorm:"column:timestamp"`
}

type TreasurySnapshot struct {
	BlockNum              int64              `gorm:"column:block_num"`
	Date                  string             `gorm:"column:date_str"`
	PricesInUSD           *core.JsonFloatMap `gorm:"column:prices_in_usd"`
	Balances              *core.JsonFloatMap `gorm:"column:balances"`
	ValueInUSD            float64            `gorm:"column:value_in_usd"`
	OperationalBalances   *core.JsonFloatMap `gorm:"column:operational_balances"`
	OperationalValueInUSD float64            `gorm:"column:operational_value_in_usd"`
}

type TreasurySnapshotModel2 struct {
	BlockNum              int64              `gorm:"column:block_num;primaryKey"`
	Date                  string             `gorm:"column:date_str"`
	PricesInUSD           *core.JsonFloatMap `gorm:"column:prices_in_usd"`
	Balances              *core.JsonFloatMap `gorm:"column:balances"`
	ValueInUSD            float64            `gorm:"column:value_in_usd"`
	OperationalBalances   *core.JsonFloatMap `gorm:"column:operational_balances"`
	OperationalValueInUSD float64            `gorm:"column:operational_value_in_usd"`
}

func (TreasurySnapshotModel2) TableName() string {
	return "treasury_snapshots"
}

const (
	// credit filter
	TokenAllowed = iota
	TokenForbidden
	ContractAllowed
	ContractForbidden
	NewFastCheckParameters
	TransferPluginAllowed
	PriceOracleUpdated
	// pools
	NewInterestRateModel
	NewCreditManagerConnected
	NewExpectedLiquidityLimit
	BorrowForbidden
	NewWithdrawFee
	// price oracle
	NewPriceFeed
	// account factory
	TakeForever
	// acl
	PausableAdminAdded
	PausableAdminRemoved
	UnpausableAdminAdded
	UnpausableAdminRemoved
	OwnershipTransferred
	// for every contract
	Paused
	UnPaused
	EventNewParameters
	///////////
	// v2 events
	///////////
	TokenAllowedV2
	LimitsUpdated
	FeesUpdated
	CreditFacadeUpgraded
	NewConfigurator
	LTUpdated
	//
	IncreaseDebtForbiddenModeChanged
	ExpirationDateUpdated
	MaxEnabledTokensUpdated
	LimitPerBlockUpdated
	AddedToUpgradeable
	RemovedFromUpgradeable
	EmergencyLiquidatorAdded
	EmergencyLiquidatorRemoved
)
