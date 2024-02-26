package core

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

const LogFilterLenError = "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response."
const QueryMoreThan10000Error = "query returned more than 10000 results"
const LogFilterQueryTimeout = "Query timeout exceeded. Consider reducing your block range."
const NoderealFilterLogError = "exceed maximum block range:"
const NoOfBlocksPerMin int64 = 5
const NoOfBlocksPerHr int64 = NoOfBlocksPerMin * 60
const SECONDS_PER_YEAR = 86400 * 365

var WETHPrice, USDCPrice *big.Int

func init() {
	WETHPrice, _ = new(big.Int).SetString("1000000000000000000", 10)
	USDCPrice, _ = new(big.Int).SetString("100000000", 10)
}

var NULL_ADDR = common.Address{}

var MAX_BIG_INT = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))

// for ray
const RAY_DECIMALS int8 = 27

var RAY *big.Int = utils.GetExpInt(RAY_DECIMALS)
