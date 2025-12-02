package core

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

// https://github.com/Gearbox-protocol/integrations-v2/tree/faa9cfd4921c62165782dcdc196ff5a0c0e6075d/contracts/oracles
// https://github.com/Gearbox-protocol/oracles-v3/tree/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles
const (
	V3_CHAINLINK_ORACLE = iota
	V3_YEARN_ORACLE
	V3_CURVE_2LP_ORACLE
	V3_CURVE_3LP_ORACLE
	V3_CURVE_4LP_ORACLE
	V3_ZERO_ORACLE
	V3_WSTETH_ORACLE
	V3_BOUNDED_ORACLE
	V3_COMPOSITE_ORACLE
	V3_WRAPPED_AAVE_V2_ORACLE
	V3_COMPOUND_V2_ORACLE
	V3_BALANCER_STABLE_LP_ORACLE
	V3_BALANCER_WEIGHTED_LP_ORACLE
	V3_CURVE_CRYPTO_ORACLE
	V3_THE_SAME_AS
	V3_REDSTONE_ORACLE
	V3_ERC4626_VAULT_ORACLE
	V3_NETWORK_DEPENDENT
	V3_CURVE_USD_ORACLE
	V3_PYTH_ORACLE
	V3_MELLOW_LRT_ORACLE
	V3_PENDLE_PT_TWAP_ORACLE

	V3_EXTERNAL

	//
	V3_BACKEND_COMPOSITE_REDSTONE_ORACLE = 100
	V3_BACKEND_GENERAL_ORACLE            = 101
	V3_PULL_UNDERLYING_ORACLE            = 102
)

func GetContractTypeToPFType(x string) int64 {
	switch x {
	case "PRICE_FEED::BALANCER_STABLE":
		return V3_BALANCER_STABLE_LP_ORACLE
	case "PRICE_FEED::BALANCER_WEIGHTED":
		return V3_BALANCER_WEIGHTED_LP_ORACLE
	case "PRICE_FEED::BOUNDED":
		return V3_BOUNDED_ORACLE
	case "PRICE_FEED::COMPOSITE":
		return V3_COMPOSITE_ORACLE
	case "PRICE_FEED::CURVE_CRYPTO":
		return V3_CURVE_CRYPTO_ORACLE
	case "PRICE_FEED::CURVE_STABLE":
		return V3_CURVE_3LP_ORACLE //
	case "PRICE_FEED::CURVE_USD":
		return V3_CURVE_USD_ORACLE
	case "PRICE_FEED::ERC4626":
		return V3_ERC4626_VAULT_ORACLE
	case "PRICE_FEED::EXTERNAL":
		return V3_EXTERNAL
	case "PRICE_FEED::MELLOW_LRT":
		return V3_MELLOW_LRT_ORACLE
	case "PRICE_FEED::PENDLE_PT_TWAP":
		return V3_PENDLE_PT_TWAP_ORACLE
	case "PRICE_FEED::PYTH":
		return V3_PYTH_ORACLE
	case "PRICE_FEED::REDSTONE":
		return V3_REDSTONE_ORACLE
	case "PRICE_FEED::WSTETH":
		return V3_WSTETH_ORACLE
	case "PRICE_FEED::YEARN":
		return V3_YEARN_ORACLE
	case "PRICE_FEED::ZERO":
		return V3_ZERO_ORACLE
	case "PRICE_FEED::CONSTANT":
		return V3_YEARN_ORACLE
	}
	log.Fatal(x)
	return 10000
}

func GetGearboxPfType(client ClientI, oracle string, token string) (int64, error) {
	// check if v300 oracle with pricefeedtype
	data, err := CallFuncGetSingleValue(client, "3fd0875f", common.HexToAddress(oracle), 0, nil) // priceFeedType
	var pfType int64
	if err != nil {
		// check if v310 oracle with contractType
		data, err := CallFuncGetSingleValue(client, "cb2ef6f7", common.HexToAddress(oracle), 0, nil) // contractType
		if err == nil {
			pfName := strings.Trim(string(data), "\x00") // contractType
			pfType = GetContractTypeToPFType(pfName)
		} else {
			// check if chainlink oracle with phaseId
			pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracle), client)
			log.CheckFatal(err)
			_, err = pfContract.PhaseId(nil) // only on chainlink
			if err == nil {
				return V3_CHAINLINK_ORACLE, nil // chainlink oracle
			}
			// check if outside redstone oracle with description
			con, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), client)
			log.CheckFatal(err)
			if description, err := con.Description(nil); err != nil {
				return 0, log.WrapErrWithLine(fmt.Errorf("%s %s priceFeedType failed: %s", oracle, token, err))
			} else {
				description = strings.ToLower(string(description))
				if strings.Contains(description, "redstone") || // the oracles that don't have priceFeedType method,
					GetBaseChainId(client) == 43111 ||
					strings.Contains(strings.ToLower(description), "aegis oracle") || // like "Aegis Oracle - chainlinkOracle", // 0x600f888a50c66bC39Ac83523EBdbba20Da984173 on ethereum.
					strings.Contains(strings.ToLower(description), "aggregatorv3 interface") || // like "AggregatorV3 interface", // https://etherscan.io/address/0xf9C7c25FE58AAA494EE7ff1f6Cf0b70d7C7ce88c#readProxyContract
					strings.Contains(strings.ToLower(description), "chainlink compatible adapter") || // https://etherscan.io/address/0xA4fd428191e263453BC802cFAD71A31DDDf74895#readContract
					// (strings.ToLower(description[:3]) == "mre" && GetBaseChainId(client) == 42793) || // like mre oracle on etherelink https://explorer.etherlink.com/address/0xc5A45B22d7E9e1c1E2524bF53B42e89376B787dF?tab=contract proxy by https://explorer.etherlink.com/address/0xc5A45B22d7E9e1c1E2524bF53B42e89376B787dF
					oracle == "0xc5A45B22d7E9e1c1E2524bF53B42e89376B787dF" || // on etherlink
					oracle == "0x1989329b72C1C81E5460481671298A5a046f3B8E" || // on etherlink  mre7yield/usd
					oracle == "0x2da05F177485264D432878D4A17d722bc64Db0EF" || // 0x2da05F177485264D432878D4A17d722bc64Db0EF on ethereum.
					oracle == "0x31D211312D9cF5A67436517C324504ebd5BD50a0" || //  on etherlink. // https://explorer.etherlink.com/address/0x31D211312D9cF5A67436517C324504ebd5BD50a0
					oracle == "0x51d947B18f546696c31d9a1c81B55d84e6d8e959" || //  on plasma. https://plasmascan.to/address/0x51d947B18f546696c31d9a1c81B55d84e6d8e959/contract/9745/readProxyContract
					oracle == "0x5D4266f4DD721c1cD8367FEb23E4940d17C83C93" || //  on somnia. https://explorer.somnia.network/address/0x5D4266f4DD721c1cD8367FEb23E4940d17C83C93?tab=read_write_contract
					oracle == "0x1f5f46B0DABEf8806a1f33772522ED683Ba64E27" || //  on somnia. https://explorer.somnia.network/address/0x046EDe9564A72571df6F5e44d0405360c0f4dCab?tab=read_write_contract
					strings.Contains(description, "hemiBTC/USD") { // the oracles that don't have priceFeedType method,
					if GetBaseChainId(client) == 43111 {
						log.Warn("The oracle on HemiBTC. has this description and is an external oracle.", description)
					}
					// // are outside redstne oracle and in control of redstone team to update regularly so can be treated as curve pf
					return V3_EXTERNAL, nil
				}
				log.Fatal(oracle, token, "priceFeedType failed: ", description, err)
			}
		}
	} else {
		pfType = new(big.Int).SetBytes(data).Int64()
	}
	return pfType, nil
}
