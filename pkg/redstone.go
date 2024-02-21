package pkg

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/routerv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type RedStoneMgr struct {
	chainId        int64
	lastResponses  *core.MutexDS[common.Address, *RSPriceOnDemand]
	redStoneTokens map[common.Address]core.Symbol
	lastUpdatedAt  int64
	updating       *atomic.Bool
}

type RedStoneMgrI interface {
	GetRedStoneForAccount(tokensNeeded []common.Address, balances core.DBBalanceFormat) (ans []dataCompressorv3.PriceOnDemand, prices map[string]float64)
	GetPriceOnDemandCalls(cf common.Address, tokensNeeded []common.Address, balances core.DBBalanceFormat) []routerv3.MultiCall
	Update(blockNum int64, blockPerUpdateRound int64)
	IsRedStoneToken(token string) bool
	GetPrice(token string) *big.Int
}

func NewRedStoneMgr(client core.ClientI) RedStoneMgrI {
	chainId := core.GetChainId(client)
	//
	redStoneTokens := map[common.Address]core.Symbol{}
	pfs := core.GetRedStonePFByChainId(chainId)
	symToAddr := core.GetSymToAddrByChainId(chainId)
	for sym := range pfs.Mains {
		address := symToAddr.Tokens[string(sym)]
		redStoneTokens[address] = sym
	}
	return &RedStoneMgr{
		chainId:        chainId,
		lastResponses:  core.NewMutexDS[common.Address, *RSPriceOnDemand](),
		redStoneTokens: redStoneTokens,
		updating:       &atomic.Bool{},
	}
}

type RSPriceOnDemand struct {
	CallData  string  `json:"callData"`
	Timestamp int64   `json:"ts"`
	Price     float64 `json:"price"`
	PriceBI   *big.Int
	// MetaData  string `json:"metaData"`
	ans dcv3.PriceOnDemand
}

func (r *RSPriceOnDemand) convert(token common.Address) {
	priceOnDemand := dcv3.PriceOnDemand{}
	priceOnDemand.Token = token
	var err error
	if len(r.CallData) >= 2 && r.CallData[:2] == "0x" {
		r.CallData = r.CallData[2:]
	}
	priceOnDemand.CallData, err = hex.DecodeString(r.CallData) // due to 0x
	if err != nil {
		log.Warnf("failed to decode callData for token %s: %s", token, r.CallData)
		return
	}
	r.PriceBI = utils.FloatDecimalsTo64(r.Price, 8)
	r.ans = priceOnDemand
}

func (r *RedStoneMgr) IsRedStoneToken(token string) bool {
	// TODO
	_, ok := r.redStoneTokens[common.HexToAddress(token)]
	return ok
}

func (r *RedStoneMgr) Update(blockNum int64, blockPerUpdateRound int64) {
	if !r.updating.CompareAndSwap(false, true) {
		return
	}
	defer func() {
		r.updating.Store(false)
	}()
	//
	if blockNum-r.lastUpdatedAt < blockPerUpdateRound {
		return
	}
	r.lastUpdatedAt = blockNum
	//
	pfs := core.GetRedStonePFByChainId(r.chainId)
	for token, sym := range r.redStoneTokens {
		details := pfs.Mains[sym]
		resp := getRSPriceOnDemand(sym, details)
		{
			resp := resp[sym]
			resp.convert(token)
			r.lastResponses.Set(token, resp)
		}
	}
}

func (r *RedStoneMgr) GetPrice(token string) *big.Int {
	details := r.lastResponses.Get(common.HexToAddress(token))
	if details == nil {
		return nil
	}
	return details.PriceBI
}

func (r *RedStoneMgr) GetRedStoneForAccount(tokensNeeded []common.Address, balances core.DBBalanceFormat) (ans []dcv3.PriceOnDemand, prices map[string]float64) {
	pfs := core.GetRedStonePFByChainId(r.chainId)
	addrToSym := core.GetTokenToSymbolByChainId(r.chainId)
	//
	prices = map[string]float64{}
	for _, token := range tokensNeeded {
		if balances[token.Hex()].IsEnabled && balances[token.Hex()].HasBalanceMoreThanOne() {
			sym := addrToSym[token]
			details := pfs.Mains[sym]
			//
			if sym == "rETH" {
				details = core.RedStonePF{
					DataServiceId:    "redstone-primary-prod",
					DataId:           "rETH",
					SignersThreshold: 5,
				}
			}
			lastResp := r.lastResponses.Get(token)
			if lastResp == nil || lastResp.Timestamp < time.Now().Unix()-details.StalenessPeriod/2 {
				// log.Info(sym, details)
				resp := getRSPriceOnDemand(sym, details)
				resp[sym].convert(token)
				r.lastResponses.Set(token, resp[sym]) // prod/aave/1
			}

			{
				lastResp := r.lastResponses.Get(token)
				if lastResp != nil {
					ans = append(ans, lastResp.ans)
					prices[token.Hex()] = lastResp.Price
				}
			}
		}
	}
	return
}

func (r *RedStoneMgr) GetPriceOnDemandCalls(cf common.Address, tokensNeeded []common.Address, balances core.DBBalanceFormat) (calls []routerv3.MultiCall) {
	abi := core.GetAbi("CreditFacadev3Multicall")
	prices, _ := r.GetRedStoneForAccount(tokensNeeded, balances)
	for _, price := range prices {
		data, err := abi.Pack("onDemandPriceUpdate", price.Token, false, price.CallData)
		log.CheckFatal(err)
		calls = append(calls, routerv3.MultiCall{
			Target:   cf,
			CallData: data,
		})
	}
	return
}

func getRSPriceOnDemand(sym core.Symbol, details core.RedStonePF) map[core.Symbol]*RSPriceOnDemand {
	// prod/aave/1
	// log.Info(details.DataServiceId, details.SignersThreshold, details.DataId)
	url := fmt.Sprintf("https://testnet.gearbox.foundation/redstone/%s/%d?dataFeeds=%s", details.DataServiceId, details.SignersThreshold, details.DataId)
	res, err := http.Get(url)
	log.Debug("Getting priceOnDemand", url)
	if err != nil || res.StatusCode/100 != 2 {
		respStr := ""
		if res.Body != nil {
			respStr = ReadBuffer(res.Body)
		}
		log.Warnf("For sym %s , redStone failed with  err(%v) and resp(%s)", sym, err, respStr)
		return nil
	}
	data := map[core.Symbol]*RSPriceOnDemand{}
	if err = utils.ReadJsonReaderAndSetInterface(res.Body, &data); err != nil {
		log.Warnf("For sym %s , redStone response parsing failed with  err(%v)", sym, err)
		return nil
	}
	return data
}
func ReadBuffer(r io.ReadCloser) string {
	b := bytes.NewBuffer(nil)
	b.ReadFrom(r)
	return b.String()
}
