package calc

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type session struct {
	balances       core.DBBalanceFormat
	borrowedAmount *big.Int
	collateral     core.JsonBigIntMap
	underlying     string
	//
	tradingToken string
	quoteToken   string
	currentPrice float64
}

func (s session) GetBalances() core.DBBalanceFormat {
	return s.balances
}
func (s session) GetBorrowedAmount() *big.Int {
	return s.borrowedAmount
}
func (s session) GetCollateralInUnderlying() interface{} {
	return s.collateral[s.underlying]
}
func (s session) GetUnderlyingToken() string {
	return s.underlying
}
func (s *session) SetCurrentPrice(tradingToken, quoteToken string, currentPrice float64) {
	s.tradingToken = tradingToken
	s.quoteToken = quoteToken
	s.currentPrice = currentPrice
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
		balances:       map[string]core.DBTokenBalance{tradingToken: {IsEnabled: true, BI: (*core.BigInt)(utils.GetExpInt(18))}},
		borrowedAmount: new(big.Int).Mul(big.NewInt(3), utils.GetExpInt(6+3)),
		collateral:     core.JsonBigIntMap{usdc: (*core.BigInt)(utils.GetExpInt(6 + 3))},
		underlying:     usdc,
	}
	CalcCurrentPrice(sess, []string{usdc}, dStore{
		tokens: map[string]*schemas.Token{
			usdc:         {Decimals: 6},
			tradingToken: {Decimals: 18},
		},
	})

	if sess.tradingToken != tradingToken {
		t.Fatal("trading token is not set", sess.tradingToken)
	}
	if sess.currentPrice != 4000 {
		t.Fatal("entry price is wrong", sess.currentPrice)
	}
}

func BenchmarkFooer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < i; j++ {
			_ = 1 + 1
		}
	}
}
