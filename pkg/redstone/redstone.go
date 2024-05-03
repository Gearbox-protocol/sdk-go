package redstone

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/routerv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type redStoneDS struct {
	data map[string]core.RedStonePF
}

func (r redStoneDS) Get(token string, composite bool) core.RedStonePF {
	if composite {
		return r.data["Composite"+token]
	}
	return r.data[token]
}

type RedStoneMgr struct {
	chainId int64
	//
	lastPods       *core.MutexDS[string, *RSPriceOnDemand] // for liquidator, third-eye , gearbox
	prices         *core.MutexDS[string, *big.Int]
	redStoneTokens redStoneDS
	//
}

type TokenAndFeedType struct {
	Token  common.Address
	PFType int
}

func (obj TokenAndFeedType) IsComposite() bool {
	return obj.PFType == core.V3_BACKEND_COMPOSITE_REDSTONE_ORACLE
}

type RedStoneMgrI interface {
	//
	GetPrice(ts int64, token string, composite bool) *big.Int
	GetPodSign(ts int64, tokensNeeded []TokenAndFeedType, balances core.DBBalanceFormat) (ans []dcv3.PriceOnDemand)
}

func NewRedStoneMgr(client core.ClientI) RedStoneMgrI {
	chainId := core.GetChainId(client)
	//
	redStoneTokens := map[string]core.RedStonePF{}
	symToAddr := core.GetSymToAddrByChainId(chainId)
	{ // redstone
		pfs := core.GetRedStonePFByChainId(chainId)
		for sym, details := range pfs {
			address := symToAddr.Tokens[string(sym)].Hex()
			redStoneTokens[address] = details
		}
	}
	{ // composite redstone
		pfs := core.GetCompositeRedStonePFByChainId(chainId)
		for sym, details := range pfs {
			address := symToAddr.Tokens[string(sym)].Hex()
			redStoneTokens["Composite"+address] = details
		}
	}
	//
	return &RedStoneMgr{
		chainId:  chainId,
		lastPods: core.NewMutexDS[string, *RSPriceOnDemand](),
		prices:   core.NewMutexDS[string, *big.Int](),
		redStoneTokens: redStoneDS{
			data: redStoneTokens,
		},
	}
}

// token is address
func (r *RedStoneMgr) GetPrice(ts int64, token string, composite bool) *big.Int {
	key := fmt.Sprintf("%d-%s", ts, token)
	if price := r.prices.Get(key); price != nil {
		return price
	}
	price, fromWhere := r.getHistoricPrice(ts, token, composite)
	r.prices.Set(key, price)
	log.Infof("RedStone price at %d for %s from %s is %d", ts, r.redStoneTokens.Get(token, composite).DataId, fromWhere, price)
	return price
}
func (r *RedStoneMgr) getHistoricPrice(ts int64, token string, composite bool) (*big.Int, string) {

	price := r.getAPIPrice(ts, token, composite)
	if price.Cmp(new(big.Int)) == 0 {
		details := r.redStoneTokens.Get(token, composite)
		ans := getHistoricPodSign(ts, details)[details.DataId]
		if ans == nil {
			log.Warnf("Failed to get podSign for token %s at timestamp %d", token, ts)
			return new(big.Int), "pod"
		} else {
			log.Info(ans.Timestamp)
			return ans.convert(common.HexToAddress(token)).PriceBI, "pod"
		}
	}
	return price, "api"
}

func tenthMillSec(ts int64) int64 {
	return (ts / 10) * 10000
}

// https://api.docs.redstone.finance/methods/gethistoricalprice
func (r *RedStoneMgr) getAPIPrice(ts int64, token string, composite bool) *big.Int {
	details := r.redStoneTokens.Get(token, composite)
	url := fmt.Sprintf("https://api.redstone.finance/prices?symbol=%s&provider=redstone&toTimestamp=%d&limit=1", details.DataId, tenthMillSec(ts))
	res, err := http.Get(url)
	if err != nil || res.StatusCode/100 != 2 {
		return new(big.Int)
	}
	// from historic api
	parsedResp := []struct {
		Value float64 `json:"value"`
	}{}
	reader, body := ReadBuffer(res.Body)
	err = utils.ReadJsonReaderAndSetInterface(reader, &parsedResp)
	if err != nil {
		log.Warnf("Failed to get historic price for token %s at timestamp %d. resbody: %s", token, ts, body)
		return new(big.Int)
	}
	if len(parsedResp) == 0 {
		log.Info("empty response from redstone api", url, token)
	}
	return utils.FloatDecimalsTo64(parsedResp[0].Value, 8)
}

func (r *RedStoneMgr) GetPodSign(ts int64, tokensNeeded []TokenAndFeedType, balances core.DBBalanceFormat) (ans []dcv3.PriceOnDemand) {
	fromWhere := ""
	if time.Now().Unix()-ts > 30 { // https://github.com/Gearbox-protocol/oracles-v3/blob/main/contracts/oracles/updatable/RedstonePriceFeed.sol#L196-L203
		// can't be ahead more than 60 seconds
		ans, fromWhere = r.getHistoricPodSign(ts, tokensNeeded, balances)
	} else {
		ans, fromWhere = r.getLatestPodSign(tokensNeeded, balances)
	}
	log.Infof("RedStone podSign at %d from %s", ts, fromWhere)
	return
}

func (r *RedStoneMgr) getLatestPodSign(tokensNeeded []TokenAndFeedType, balances core.DBBalanceFormat) (ans []dcv3.PriceOnDemand, fromWhere string) {
	//
	fromWhere = "latest"
	for _, token := range tokensNeeded {
		if balances[token.Token.Hex()].IsEnabled && balances[token.Token.Hex()].HasBalanceMoreThanOne() {
			details := r.redStoneTokens.Get(token.Token.Hex(), token.IsComposite())
			// lat response not set
			resp := getLatestPodSign(details)
			lastResp := resp[details.DataId].convert(token.Token) // prod/aave/1
			if lastResp != nil {
				ans = append(ans, lastResp.pod)
				fromWhere = fmt.Sprintf("latest-%d", lastResp.Timestamp)
			}
		}
	}
	return ans, fromWhere
}

func (r *RedStoneMgr) getHistoricPodSign(ts int64, tokensNeeded []TokenAndFeedType, balances core.DBBalanceFormat) (ans []dcv3.PriceOnDemand, fromWhere string) {
	//
	fromWhere = "historic"
	for _, token := range tokensNeeded {
		if balances[token.Token.Hex()].IsEnabled && balances[token.Token.Hex()].HasBalanceMoreThanOne() {
			key := fmt.Sprintf("%d-%s", ts, token.Token)
			details := r.redStoneTokens.Get(token.Token.Hex(), token.IsComposite())
			//
			lastResp := r.lastPods.Get(key)
			if lastResp == nil { // back , ahead 30 secs
				resp := getHistoricPodSign(ts, details)
				r.lastPods.Set(key, resp[details.DataId].convert(token.Token)) // prod/aave/1
				lastResp = resp[details.DataId].convert(token.Token)
			} else {
				fromWhere = "historic-stored"
			}

			{
				if lastResp != nil {
					ans = append(ans, lastResp.pod)
				}
			}
		}
	}
	return
}

func GetPriceOnDemandCalls(cf common.Address, pods []dcv3.PriceOnDemand) (calls []routerv3.MultiCall) {
	abi := core.GetAbi("CreditFacadev3Multicall")
	for _, pod := range pods {
		data, err := abi.Pack("onDemandPriceUpdate", pod.Token, false, pod.CallData)
		log.CheckFatal(err)
		calls = append(calls, routerv3.MultiCall{
			Target:   cf,
			CallData: data,
		})
	}
	return
}

func ReadBuffer(r io.ReadCloser) (*bytes.Buffer, string) {
	b := bytes.NewBuffer(nil)
	b.ReadFrom(r)
	s := b.String()
	return bytes.NewBuffer([]byte(s)), s
}
