package schemas

import (
	"github.com/Gearbox-protocol/sdk-go/core"
)

type AccountOperation struct {
	// Input string
	ID int64 `gorm:"primaryKey;autoIncrement:true" json:"-"`
	// ordering data
	TxHash      string `gorm:"column:tx_hash" json:"txHash"`
	BlockNumber int64  `gorm:"column:block_num" json:"blockNum"`
	LogId       uint   `gorm:"column:log_id" json:"logId"`
	// owner/account data
	Borrower  string `json:"borrower"`
	SessionId string `gorm:"column:session_id" json:"sessionId"`
	// application
	Dapp string `gorm:"column:dapp" json:"dapp"`
	// call/events data
	AdapterCall bool            `gorm:"column:adapter_call" json:"adapterCall"`
	Action      string          `gorm:"column:action" json:"action"`
	Args        *core.Json      `gorm:"column:args" json:"args"`
	Transfers   *core.Transfers `gorm:"column:transfers" json:"transfers"`
	// Note that MainAction is pointer to int64
	// this avoids a foreign key violation, else  0 will be inserted for this column, which is foreign key violation.
	MainAction *int64              `gorm:"column:main_action"`
	MultiCall  []*AccountOperation `gorm:"foreignkey:MainAction" json:"multicalls"`
}

func (AccountOperation) TableName() string {
	return "account_operations"
}

const (
	EventType = iota
	CallType
)

type AccountOperationState struct {
	ID               int64                 `gorm:"primaryKey;autoincrement:true" json:"id"`
	BlockNum         int64                 `gorm:"column:block_num"`
	LogId            int64                 `gorm:"column:log_id"`
	SessionId        string                `gorm:"column:session_id"`
	BorrowedAmountBI *core.BigInt          `gorm:"column:borrowed_amount_bi"`
	BorrowedAmount   float64               `gorm:"column:borrowed_amount"`
	Balances         *core.DBBalanceFormat `gorm:"column:balances"`
}

func (AccountOperationState) TableName() string {
	return "account_operation_states"
}
