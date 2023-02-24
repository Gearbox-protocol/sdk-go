package schemas

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
)

type LatestRounData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

type UniPoolPrices struct {
	ID             int64   `gorm:"primaryKey;column:id;autoIncrement:true"`
	PriceV2        float64 `gorm:"column:uniswapv2_price"`
	TwapV3         float64 `gorm:"column:uniswapv3_price"`
	PriceV3        float64 `gorm:"column:uniswapv3_twap"`
	PriceV2Success bool    `gorm:"-"`
	TwapV3Success  bool    `gorm:"-"`
	PriceV3Success bool    `gorm:"-"`
	BlockNum       int64   `gorm:"column:block_num"`
	Token          string  `gorm:"column:token"`
}

func (UniPoolPrices) TableName() string {
	return "uniswap_pool_prices"
}

type UniswapPools struct {
	V2     string `gorm:"column:pool_v2"`
	V3_100 string `gorm:"column:pool_v3_100"`
	V3_500 string `gorm:"column:pool_v3_500"`
	V3_3k  string `gorm:"column:pool_v3_3k"`
	V3_10k string `gorm:"column:pool_v3_10k"`

	Token   string `gorm:"column:token;primaryKey"`
	Updated bool   `gorm:"-"`
}

type FeeAndUniv3Pool struct {
	Fee  int16
	Pool string
}

func (p UniswapPools) GetPools() (pools []FeeAndUniv3Pool) {
	nullAddr := core.NULL_ADDR.Hex()
	fees := []int16{100, 500, 3000, 10000}
	for ind, pool := range []string{p.V3_100, p.V3_500, p.V3_3k, p.V3_10k} {
		if pool != nullAddr {
			pools = append(pools, FeeAndUniv3Pool{Fee: fees[ind], Pool: pool})
		}
	}
	return pools
}

type TokenSyncDetails struct {
	Decimals int8
	LastSync int64
}

// sort event balances by block number/log id
type SortedUniPoolPrices []*UniPoolPrices

func (ts SortedUniPoolPrices) Len() int {
	return len(ts)
}
func (ts SortedUniPoolPrices) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts SortedUniPoolPrices) Less(i, j int) bool {
	return ts[i].BlockNum < ts[j].BlockNum
}
