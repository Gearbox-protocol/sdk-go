package calc

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type session struct {
	balances       map[string]core.BalanceType
	borrowedAmount *big.Int
	collateral     core.JsonBigIntMap
	underlying     string
	//
	tradingToken string
	quoteToken   string
	entryPrice   float64
}

func (s session) GetBalances() map[string]core.BalanceType {
	return s.balances
}
func (s session) GetBorrowedAmount() *big.Int {
	return s.borrowedAmount
}
func (s session) GetCollateral() *core.JsonBigIntMap {
	return &s.collateral
}
func (s session) GetUnderlyingToken() string {
	return s.underlying
}
func (s *session) SetEntryPrice(tradingToken, quoteToken string, entryPrice float64) {
	s.tradingToken = tradingToken
	s.quoteToken = quoteToken
	s.entryPrice = entryPrice
}

type dStore struct {
	tokens map[string]*schemas.Token
}

func (d dStore) GetToken(token string) *schemas.Token {
	return d.tokens[token]
}
func TestEntryPriceForLong(t *testing.T) {
	usdc := utils.RandomAddr()
	tradingToken := utils.RandomAddr()
	sess := &session{
		balances:       map[string]core.BalanceType{tradingToken: {IsEnabled: true, BI: (*core.BigInt)(utils.GetExpInt(18))}},
		borrowedAmount: new(big.Int).Mul(big.NewInt(3), utils.GetExpInt(6+3)),
		collateral:     core.JsonBigIntMap{usdc: (*core.BigInt)(utils.GetExpInt(6 + 3))},
		underlying:     usdc,
	}
	CalcEntryPrice(sess, []string{usdc}, dStore{
		tokens: map[string]*schemas.Token{
			usdc:         {Decimals: 6},
			tradingToken: {Decimals: 18},
		},
	})

	if sess.tradingToken != tradingToken {
		t.Fatal("trading token is not set", sess.tradingToken)
	}
	if sess.entryPrice != 4000 {
		t.Fatal("entry price is wrong", sess.entryPrice)
	}
}

func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < i; j++ {
			_ = 1 + 1
		}
	}
}
