package schemas

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/artifacts/eRC20"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address  string       `gorm:"primaryKey;column:address" json:"addr"`
	Symbol   string       `gorm:"column:symbol" json:"symbol"`
	Decimals int8         `gorm:"column:decimals" json:"decimals"`
	client   core.ClientI `gorm:"-" json:"-"`
} // @name TokenData

func (Token) TableName() string {
	return "tokens"
}

func NewToken(addr string, client core.ClientI) (*Token, error) {
	token := &Token{
		Address: addr,
		client:  client,
	}
	err := token.init()
	return token, err
}

func (t *Token) init() error {
	tokenAddr := common.HexToAddress(t.Address)
	contract, err := eRC20.NewERC20(tokenAddr, t.client)
	if err != nil {
		return err
	}
	if symbol, err := contract.Symbol(&bind.CallOpts{}); err != nil {
		if err2 := func() error {
			symbolBytes, err := core.CallFuncWithExtraBytes(t.client, "95d89b41", tokenAddr, 0, nil)
			if err != nil {
				return err
			}
			t.Symbol = func() string {
				for ind, ele := range symbolBytes {
					if ele == 0 {
						symbolBytes = symbolBytes[:ind]
						break
					}
				}
				return string(symbolBytes)
			}()
			return nil
		}(); err2 != nil { // if we can't get bytes data return error
			return err
		}
	} else {
		t.Symbol = symbol
	}
	if decimals, err := contract.Decimals(&bind.CallOpts{}); err != nil {
		return err
	} else {
		t.Decimals = int8(decimals)
	}
	return nil
}

type AllowedToken struct {
	BlockNumber        int64        `gorm:"column:block_num;primaryKey"`
	CreditManager      string       `gorm:"column:credit_manager;primaryKey"`
	Configurator       string       `gorm:"column:configurator"`
	Token              string       `gorm:"column:token;primaryKey"`
	LiquidityThreshold *core.BigInt `gorm:"column:liquiditythreshold"`
	DisableBlock       int64        `gorm:"column:disable_block"`
}

func (AllowedToken) TableName() string {
	return "allowed_tokens"
}

type TokenTransfer struct {
	BlockNum      int64        `gorm:"column:block_num;primaryKey"`
	LogID         uint         `gorm:"column:log_id;primaryKey"`
	TxHash        string       `gorm:"column:tx_hash"`
	Token         string       `gorm:"column:token"`
	From          string       `gorm:"column:source"`
	To            string       `gorm:"column:destination"`
	Amount        *core.BigInt `gorm:"column:amount"`
	IsFromAccount bool         `gorm:"column:isfrom_account"`
	IsToAccount   bool         `gorm:"column:isto_account"`
}

func (tt *TokenTransfer) String() string {
	return fmt.Sprintf("DirecTokenTransfer detected at %d from %s to %s for token(%s) amount(%s)",
		tt.BlockNum, tt.From, tt.To, tt.Token, tt.Amount)
}

type TokenTransferList []*TokenTransfer

func (ts TokenTransferList) Len() int {
	return len(ts)
}
func (ts TokenTransferList) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts TokenTransferList) Less(i, j int) bool {
	return ts[i].LogID < ts[j].LogID
}
func (a *TokenTransfer) Equal(b *TokenTransfer) bool {
	return a.BlockNum == b.BlockNum &&
		a.LogID == b.LogID &&
		a.TxHash == b.TxHash &&
		a.Token == b.Token &&
		a.From == b.From &&
		a.To == b.To &&
		a.Amount == b.Amount &&
		a.IsFromAccount == b.IsFromAccount &&
		a.IsToAccount == b.IsToAccount
}

func (TokenTransfer) TableName() string {
	return "no_session_transfers"
}

type TokenPrice struct {
	Token    core.Symbol `gorm:"column:token;primaryKey" json:"symbol"`
	Ts       uint64      `gorm:"column:ts;primaryKey" json:"ts"`
	Price    float64     `gorm:"column:price" json:"price"`
	Category string      `gorm:"column:category;primaryKey" json:"-"`
}

func (TokenPrice) TableName() string {
	return "prices"
}
