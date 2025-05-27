package core

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

type Symbol string

type RedStonePF struct {
	Type             int            `json:"type"`
	DataServiceId    string         `json:"dataServiceId"`
	DataId           string         `json:"dataId"`
	SignersThreshold int            `json:"signersThreshold"`
	UnderlyingToken  common.Address `json:"token"`
	Feed             common.Address `json:"feed"`
}

type SymTOAddrStore struct {
	Exchanges         map[string]common.Address `json:"exchanges"`
	Names             map[string]string         `json:"names"`
	Tokens            map[string]common.Address `json:"tokens"`
	RedStone          map[Symbol]RedStonePF     `json:"redStone"`
	CompositeRedStone map[Symbol]RedStonePF     `json:"compositeRedStone"`
}

func (s *SymTOAddrStore) getTokenAddr(sym Symbol) (string, bool) {
	if _, ok := s.Tokens[string(sym)]; !ok {
		// log.Fatal("can't get token", sym)
		return "", false
	}
	return s.Tokens[string(sym)].Hex(), true
}

var _globalCopy = map[string]*SymTOAddrStore{}

func getSymToAddrStore(fileName string) *SymTOAddrStore {
	if _globalCopy[fileName] == nil {
		data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
		if err == nil {
			store := &SymTOAddrStore{}
			err = json.Unmarshal([]byte(data), store)
			log.CheckFatal(err)
			//
			_globalCopy[fileName] = store
		} else {
			_globalCopy[fileName] = &SymTOAddrStore{}
		}
	}
	return _globalCopy[fileName]
}

// func GetRedStonePFByChainId(chainId int64) map[Symbol]RedStonePF {
// 	fileName := log.GetConfigFile(chainId)
// 	data := getSymToAddrStore(fileName)
// 	return data.RedStone
// }
// func GetCompositeRedStonePFByChainId(chainId int64) map[Symbol]RedStonePF {
// 	fileName := log.GetConfigFile(chainId)
// 	data := getSymToAddrStore(fileName)
// 	return data.CompositeRedStone
// }

// func getAddrToSymbol(fileName string, opts map[string]bool) map[common.Address]Symbol {
// 	store := getSymToAddrStore(fileName)
// 	addrToName := map[common.Address]Symbol{}
// 	if opts["tokens"] {
// 		for name, token := range store.Tokens {
// 			addrToName[token] = Symbol(name)
// 		}
// 	}
// 	if opts["exchanges"] {
// 		for name, exchg := range store.Exchanges {
// 			addrToName[exchg] = Symbol(name)
// 		}
// 	}
// 	return addrToName
// }

func GetSymToAddr(chainId int64) *SymTOAddrStore {
	fileName := log.GetConfigFile(chainId)
	return getSymToAddrStore(fileName)
}

const (
	NORMAL = iota
	WRAPPED_NATIVE
	NATIVE
)

func GetType(chainId int64, sym string) int64 {
	sym = strings.ToLower(sym)
	switch log.GetBaseNet(chainId) {
	case log.SONIC:
		if sym == "s" {
			return NATIVE
		} else if sym == "ws" {
			return WRAPPED_NATIVE
		}
	case log.BNB:
		if sym == "bnb" {
			return NATIVE
		} else if sym == "wbnb" {
			return WRAPPED_NATIVE
		}
	case log.MAINNET:
		if sym == "eth" {
			return NATIVE
		} else if sym == "weth" {
			return WRAPPED_NATIVE
		}
	}
	return NORMAL
}

// func GetSymToNameByChainId(chainId int64, sym string) string {
// 	if GetType(chainId, sym) == NATIVE {
// 		return sym
// 	}
// 	fileName := log.GetConfigFile(chainId)
// 	return getSymToAddrStore(fileName).Names[sym]
// }

// func GetTokenToSymbolByChainId(chainId int64) map[common.Address]Symbol {
// 	fileName := log.GetConfigFile(chainId)
// 	return getAddrToSymbol(fileName, map[string]bool{"tokens": true})
// }
// func GetExchangeToSymbolByChainId(chainId int64) map[common.Address]Symbol {
// 	fileName := log.GetConfigFile(chainId)
// 	return getAddrToSymbol(fileName, map[string]bool{"exchanges": true})
// }

func GetToken(chainId int64, token Symbol) common.Address {
	tokens := map[log.NETWORK]map[Symbol]string{
		log.MAINNET: {
			// "WETH":  common.HexToAddress("0xC02aaA39b223FE8D0A0e6C6324aE7E56dB8f03d8"),
			// "WBTC":  common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"),
			// "USDC":  common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			"GEAR": "0xBa3335588D9403515223F109EdC4eB7269a9Ab5D",
			// "stETH": common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84"),
			//
			"ETH":   "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			"WETH":  "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			"DAI":   "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			"USDC":  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			"USDe":  "0x4c9EDD5852cd905f086C759E8383e09bff1E68B3",
			"FRAX":  "0x853d955aCEf822Db058eb8505911ED77F175b99e",
			"GHO":   "0x40D16FC0246aD3160Ccc09B8D0D3A2cD28aE6C2f",
			"rETH":  "0xae78736Cd615f374D3085123A210448E74Fc6393",
			"ezETH": "0xbf5495Efe5DB9ce00f80364C8B423567e58d2110",
			"weETH": "0xCd5fE23C85820F7B72D0926FC9b05b43E359b7ee",
			// "wstETH":  "0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0",
			"stETH":   "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84",
			"WBTC":    "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
			"USDS":    "0xdC035D45d973E3EC169d2276DDab16f1e407384F",
			"eBTC":    "0x657e8C867D8B37dCC18fA4Caead9C45EB088C642",
			"LBTC":    "0x8236a87084f8B84306f72007F36F2618A5634494",
			"solvBTC": "0xd9D920AA40f578ab794426F5C90F6C731D159DEf",
			"pumpBTC": "0xF469fBD2abcd6B9de8E169d128226C0Fc90a012e",
			//
			"GEARBOX_WETH_POOL":   "0xB03670c20F87f2169A7c4eBE35746007e9575901",
			"WETH_GATEWAY":        "0x4F952c4C5415B2609899AbDC2F8F352F600d14D6",
			"WSTETH_GATEWAY":      "0x5a97e3E43dCBFe620ccF7865739075f92E93F5E4",
			"GEARBOX_WSTETH_POOL": "0xB8cf3Ed326bB0E51454361Fb37E9E8df6DC5C286",
			"wstETH":              "0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0",
			// liquidator
			"DOLA":                 "0x865377367054516e17014CcdED1e7d814EDC9ce4",
			"AAVE_V3_LENDING_POOL": "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
			"waEthLidowstETH":      "0x775F661b0bD1739349b9A2A3EF60be277c5d2D29",
		},
		log.ARBITRUM: {
			// "WETH": common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1"),
			"WBTC": "0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f",
			// "USDC": common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
			//
			"WETH":   "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1",
			"DAI":    "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1",
			"USDC":   "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			"USDT":   "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
			"rETH":   "0xEC70Dcb4A1EFa46b8F2D97C310C9c4790ba5ffA8",
			"USDC_e": "0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8",
			"wstETH": "0x5979D7b546E38E414F7E9822514be443A4800529",
		},
		log.SONIC: {
			"WBTC": "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c",
			// "s": common.HexToAddress("0x50c42dEAcD8Fc9773493ED674b675bE577f2634b"),
			//
			"WETH":   "0x50c42dEAcD8Fc9773493ED674b675bE577f2634b",
			"wS":     "0x039e2fB66102314Ce7b64Ce5Ce3E5183bc94aD38",
			"S":      "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			"USDC_e": "0x29219dd400f2Bf60E5a23d13Be72B486D4038894",
		},
		log.OPTIMISM: {
			// "WETH": common.HexToAddress("0x4200000000000000000000000000000000000006"),
			// "USDC": common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
			"WBTC": "0x68f180fcCe6836688e9084f035309E29Bf0A2095",
			//
			"WETH":   "0x4200000000000000000000000000000000000006",
			"USDC":   "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			"USDT":   "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58",
			"USDC_e": "0x7F5c764cBc14f9669B88837ca1490cCa17c31607",
			"wstETH": "0x1F32b1c2345538c0c6f582fCB022739c4A194Ebb",
		},
		log.BNB: {
			"WBTC": "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c",
			"WETH": "0x2170Ed0880ac9A755fd29B2688956BD959F933F8",
			"WBNB": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
			"BNB":  "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			"USDC": "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d",
		},
	}
	network := log.GetBaseNet(chainId)
	if _, ok := tokens[network]; !ok {
		log.Fatalf("unsupported network %s", network)
	}
	if _, ok := tokens[network][token]; !ok {
		log.Fatalf("unsupported token %s %s", token, network)
	}
	return common.HexToAddress(tokens[network][token])
}

func GetDecimals(client ClientI, addr common.Address, blockNum int64) int8 {
	decimals, err := CallFuncGetSingleValue(client, "313ce567", addr, blockNum, nil) // decimals
	if err != nil {
		log.Fatalf("Can't get decimals for addr(%s) : %s", addr, err)
	}
	return int8(new(big.Int).SetBytes(decimals).Int64())
}

type PriceSource string

const (
	SOURCE_SPOT    PriceSource = "spot"
	SOURCE_GEARBOX PriceSource = "gearbox"
)

type TokenGroup struct {
	CurvePools        map[string]int64  `json:"curvePools"`
	BalancerTokens    map[string]int64  `json:"balancerTokens"`
	YearnCurveTokens  map[string]string `json:"yearnCurveTokens"`
	ConvexCurveTokens map[string]string `json:"convexCurveTokens"`
}

type tokenGroupWrapper struct {
	Groups struct {
		CurvePools        map[Symbol]int64  `json:"curvePools"`
		BalancerTokens    map[Symbol]int64  `json:"balancerTokens"`
		YearnCurveTokens  map[Symbol]Symbol `json:"yearnCurveTokens"`
		ConvexCurveTokens map[Symbol]Symbol `json:"convexCurveTokens"`
	} `json:"groups"`
}

func newTokenGroup() *TokenGroup {
	return &TokenGroup{
		CurvePools:        map[string]int64{},
		BalancerTokens:    map[string]int64{},
		YearnCurveTokens:  map[string]string{},
		ConvexCurveTokens: map[string]string{},
	}
}

func getTokenGroups(fileName string) *TokenGroup {
	data, err := GetEmbeddedJsonnet(fileName, JsonnetImports{})
	log.CheckFatal(err)
	store := &tokenGroupWrapper{}
	err = json.Unmarshal([]byte(data), store)
	log.CheckFatal(err)

	//
	obj := newTokenGroup()
	{
		symToAddr := getSymToAddrStore(fileName)
		for k, v := range store.Groups.CurvePools {
			k, ok := symToAddr.getTokenAddr(k)
			if ok {
				obj.CurvePools[k] = v
			}
		}
		for k, v := range store.Groups.BalancerTokens {
			k, ok := symToAddr.getTokenAddr(k)
			if ok {
				obj.BalancerTokens[k] = v
			}
		}
		for k, v := range store.Groups.ConvexCurveTokens {
			k, ok := symToAddr.getTokenAddr(k)
			v, ok2 := symToAddr.getTokenAddr(v)
			if ok && ok2 {
				obj.ConvexCurveTokens[k] = v
			}
		}
		for k, v := range store.Groups.YearnCurveTokens {
			k, ok := symToAddr.getTokenAddr(k)
			v, ok2 := symToAddr.getTokenAddr(v)
			if ok && ok2 {
				obj.YearnCurveTokens[k] = v
			}
		}
	}
	return obj
}

func GetTokenGroups(chainId int64) *TokenGroup {
	fileName := log.GetConfigFile(chainId)
	return getTokenGroups(fileName)
}
