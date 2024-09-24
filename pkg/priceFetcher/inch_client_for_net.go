package priceFetcher

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type prefixs struct {
	infura  string
	alchemy string
}

func rpcRrefixForDifferentProviders(network log.NETWORK) prefixs {
	var infura, alchemy string
	switch network {
	case log.ARBITRUM:
		infura = "https://arbitrum-mainnet.infura.io"
		alchemy = "arb-mainnet"
	case log.OPTIMISM:
		infura = "https://optimism-mainnet.infura.io"
		alchemy = "opt-mainnet"
	case log.MAINNET:
		infura = "https://mainnet.infura.io"
		alchemy = "eth-mainnet"
	}
	if infura == "" || alchemy == "" {
		log.Fatal("network not supported", network)
	}
	return prefixs{
		infura:  infura,
		alchemy: alchemy,
	}
}

func GetNetworkClient(urls string, toNetwork log.NETWORK) core.ClientI {
	newUrls := []string{}
	toPrefixes := rpcRrefixForDifferentProviders(toNetwork)
	for _, url := range strings.Split(urls, ",") {
		// already valid
		if strings.Contains(url, toPrefixes.alchemy) || strings.Contains(url, toPrefixes.infura) {
			newUrls = append(newUrls, url)
			continue
		}
		// getNewClient
		for _, fromNetwork := range []log.NETWORK{log.ARBITRUM, log.MAINNET, log.OPTIMISM} {
			fromPrefixes := rpcRrefixForDifferentProviders(fromNetwork)
			if strings.Contains(url, fromPrefixes.alchemy) {
				url = strings.Replace(url, fromPrefixes.alchemy, toPrefixes.alchemy, 1)
				newUrls = append(newUrls, url)
				break
			} else if strings.Contains(url, fromPrefixes.infura) {
				url = strings.Replace(url, fromPrefixes.infura, toPrefixes.infura, 1)
				newUrls = append(newUrls, url)
				break
			}
		}
	}
	return _getClient(strings.Join(newUrls, ","))
}

func _getClient(url string) core.ClientI {
	if url == "" {
		return nil
	}
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil
	}
	return ethclient
}

// arbitrum
// arbitrum
// arbitrum
func (calc OneInchOracle) arbForMainnet(mainnetTs uint64, prices map[string]*core.BigInt) {
	if calc.extraurls.arbclient != nil {
		defer utils.Elapsed("arbitrum price fetch")()
		calls := calc.GetArbBaseCalls()
		arbblock := pkg.GetBlockNum(mainnetTs, 42161)
		if arbblock == 0 {
			return
		}
		results := core.MakeMultiCall(calc.extraurls.arbclient, arbblock, false, calls)
		calc.processSeparateBaseResults(results, prices, calc.ArbBaseTokens)
	}
}

func (calc OneInchOracle) GetArbBaseCalls() (calls []multicall.Multicall2Call) {
	symToAddr := core.GetSymToAddrByChainId(42161)
	pfABI := core.GetAbi("1InchOracle")
	for _, token := range calc.ArbBaseTokens {
		tokenAddr := symToAddr.Tokens[token]
		data, err := pfABI.Pack("getRate", tokenAddr, symToAddr.Tokens["USDC"], false)
		if err != nil {
			log.Fatal(err)
		}
		calls = append(calls, multicall.Multicall2Call{
			Target:   get1InchAddress(42161),
			CallData: data,
		})
	}
	return
}

// optimism
// optimism
// optimism
func (calc OneInchOracle) optForMainnet(mainnetTs uint64, prices map[string]*core.BigInt) {
	if calc.extraurls.optclient != nil {
		defer utils.Elapsed("optimism price fetch")()
		calls := calc.GetOptBaseCalls()
		optblock := pkg.GetBlockNum(mainnetTs, 10)
		if optblock == 0 {
			return
		}
		results := core.MakeMultiCall(calc.extraurls.optclient, optblock, false, calls)
		calc.processSeparateBaseResults(results, prices, calc.OptBaseTokens)
	}
}

func (calc OneInchOracle) GetOptBaseCalls() (calls []multicall.Multicall2Call) {
	symToAddr := core.GetSymToAddrByChainId(10)
	pfABI := core.GetAbi("1InchOracle")
	for _, token := range calc.OptBaseTokens {
		tokenAddr := symToAddr.Tokens[token]
		data, err := pfABI.Pack("getRate", tokenAddr, symToAddr.Tokens["USDC"], false)
		if err != nil {
			log.Fatal(err)
		}
		calls = append(calls, multicall.Multicall2Call{
			Target:   get1InchAddress(10),
			CallData: data,
		})
	}
	return
}
