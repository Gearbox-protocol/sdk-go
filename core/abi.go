package core

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/aCL"
	"github.com/Gearbox-protocol/sdk-go/artifacts/aCLTrait"
	"github.com/Gearbox-protocol/sdk-go/artifacts/accountFactory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/addressProvider"
	"github.com/Gearbox-protocol/sdk-go/artifacts/contractsRegister"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccount"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	mainnetDC "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dieselToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/eRC20"
	"github.com/Gearbox-protocol/sdk-go/artifacts/gearToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidoKovan"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidoMock"
	"github.com/Gearbox-protocol/sdk-go/artifacts/linearInterestRateModel"
	"github.com/Gearbox-protocol/sdk-go/artifacts/mainnetLido"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/nonFungiblePosManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/rewardPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/tokenMock"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/wETHGateway"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetAbi(contractName string) *abi.ABI {
	abiMap := map[string]string{
		"Version": versionABI,
		"Pauser":  pauserABI,
	}
	if data := abiMap[contractName]; data != "" {
		return getABI(data)
	}
	metadataMap := map[string]*bind.MetaData{
		////////////////
		// common
		////////////////
		// Configuration
		"ACL":              aCL.ACLMetaData,
		"AddressProvider":  addressProvider.AddressProviderMetaData,
		"ACLTrait":         aCLTrait.ACLTraitMetaData,
		"ContractRegister": contractsRegister.ContractsRegisterMetaData,

		// Core
		"AccountFactory": accountFactory.AccountFactoryMetaData,
		"CreditAccount":  creditAccount.CreditAccountMetaData,
		"WETHGateway":    wETHGateway.WETHGatewayMetaData,

		// Oracle
		"YearnPriceFeed":        &bind.MetaData{ABI: yearnPriceFeed.YearnPriceFeedABI},
		"PriceFeed":             &bind.MetaData{ABI: priceFeed.PriceFeedABI},
		"DataCompressorMainnet": &bind.MetaData{ABI: mainnetDC.DataCompressorABI},
		"DataCompressorV2":      &bind.MetaData{ABI: dataCompressorv2.DataCompressorv2ABI},

		// tokens
		"DieselToken": dieselToken.DieselTokenMetaData,
		"GearToken":   gearToken.GearTokenMetaData,
		"TokenMock":   tokenMock.TokenMockMetaData,
		"Token":       eRC20.ERC20MetaData,

		//////////////
		// v1
		//////////////
		"PriceOracle":             &bind.MetaData{ABI: priceOracle.PriceOracleABI},
		"CreditManager":           &bind.MetaData{ABI: creditManager.CreditManagerABI},
		"CreditManagerv2":         &bind.MetaData{ABI: creditManagerv2.CreditManagerv2ABI},
		"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
		"CreditFilter":            &bind.MetaData{ABI: creditFilter.CreditFilterABI},
		"Pool":                    poolService.PoolServiceMetaData,

		//////////////
		// v2
		//////////////
		"PriceOraclev2":      &bind.MetaData{ABI: priceOraclev2.PriceOraclev2ABI},
		"CreditConfigurator": &bind.MetaData{ABI: creditConfigurator.CreditConfiguratorABI},
		"CreditFacade":       &bind.MetaData{ABI: creditFacade.CreditFacadeABI},
		// multicall
		"MultiCall": &bind.MetaData{ABI: multicall.MulticallABI},
		//
		"Uniswapv3Pool":    &bind.MetaData{ABI: uniswapv3Pool.Uniswapv3PoolABI},
		"Uniswapv3Router":  &bind.MetaData{ABI: uniswapv3Router.Uniswapv3RouterABI},
		"Uniswapv3Factory": &bind.MetaData{ABI: uniswapv3Factory.Uniswapv3FactoryABI},
		//
		"Uniswapv2Pool":    &bind.MetaData{ABI: uniswapv2Pool.Uniswapv2PoolABI},
		"Uniswapv2Router":  &bind.MetaData{ABI: uniswapv2Router.Uniswapv2RouterABI},
		"Uniswapv2Factory": &bind.MetaData{ABI: uniswapv2Factory.Uniswapv2FactoryABI},
		//
		"LidoMock":    &bind.MetaData{ABI: lidoMock.LidoMockABI},
		"MainnetLido": &bind.MetaData{ABI: mainnetLido.MainnetLidoABI},
		"LidoKovan":   &bind.MetaData{ABI: lidoKovan.LidoKovanABI},
		//
		"YearnAdapter": &bind.MetaData{ABI: yearnAdapter.YearnAdapterABI},
		// convex for liquidator calculation
		"RewardPool": &bind.MetaData{ABI: rewardPool.RewardPoolABI},
		// v3 pos for synctron
		"NonFungiblePosManager": &bind.MetaData{ABI: nonFungiblePosManager.NonFungiblePosManagerABI},
	}
	abiStr, ok := metadataMap[contractName]
	if !ok {
		panic("ABI for %s doesn't exists")
	}

	abi, err := abiStr.GetAbi()
	if err != nil {
		log.Infof("Cant get ABI for %s", contractName)
		log.Fatal(err)
	}

	return abi
}

func getABI(data string) *abi.ABI {
	abi, err := abi.JSON(strings.NewReader(data))
	log.CheckFatal(err)
	return &abi
}
