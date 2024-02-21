package pkg

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
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
	chainId       int64
	lastResponses *core.MutexDS[common.Address, *RSPriceOnDemand]
}

type RedStoneMgrI interface {
	GetRedStoneForAccount(tokensNeeded []common.Address, balances core.DBBalanceFormat) (ans []dataCompressorv3.PriceOnDemand, prices map[string]float64)
	GetPriceOnDemandCalls(cf common.Address, tokensNeeded []common.Address, balances core.DBBalanceFormat) []routerv3.MultiCall
	Update()
}

func NewRedStoneMgr(client core.ClientI) RedStoneMgrI {
	return &RedStoneMgr{
		chainId:       core.GetChainId(client),
		lastResponses: core.NewMutexDS[common.Address, *RSPriceOnDemand](),
	}
}

type RSPriceOnDemand struct {
	CallData  string  `json:"callData"`
	Timestamp int64   `json:"ts"`
	Price     float64 `json:"price"`
	// MetaData  string `json:"metaData"`
	ans dcv3.PriceOnDemand
}

func (r *RSPriceOnDemand) convert(token common.Address) {
	price := dcv3.PriceOnDemand{}
	price.Token = token
	var err error
	if len(r.CallData) >= 2 && r.CallData[:2] == "0x" {
		r.CallData = r.CallData[2:]
	}
	price.CallData, err = hex.DecodeString(r.CallData) // due to 0x
	if err != nil {
		log.Warnf("failed to decode callData for token %s: %s", token, r.CallData)
		return
	}
	r.ans = price
}

func (r *RedStoneMgr) Update() {
	pfs := core.GetRedStonePFByChainId(r.chainId)
	symToAddr := core.GetSymToAddrByChainId(r.chainId)
	for sym, details := range pfs.Mains {
		// log.Info(sym, details)
		resp := getRSPriceOnDemand(sym, details)
		for sym, resp := range resp {
			token := symToAddr.Tokens[string(sym)]
			resp.convert(token)
			r.lastResponses.Set(token, resp)
		}
	}
}

func (r *RedStoneMgr) GetPrice(token string) (float64, bool) {
	details := r.lastResponses.Get(common.HexToAddress(token))
	if details == nil {
		return 0, false
	}
	return details.Price, true
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
	log.Info("Getting priceOnDemand", url)
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
