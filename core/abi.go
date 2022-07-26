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
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/rewardPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/tokenMock"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/wETHGateway"
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
		"PriceOracle":           &bind.MetaData{ABI: priceOracle.PriceOracleABI},
		"YearnPriceFeed":        &bind.MetaData{ABI: yearnPriceFeed.YearnPriceFeedABI},
		"PriceFeed":             &bind.MetaData{ABI: priceFeed.PriceFeedABI},
		"DataCompressorMainnet": &bind.MetaData{ABI: mainnetDC.DataCompressorABI},
		"DataCompressorV2":      &bind.MetaData{ABI: dataCompressorv2.DataCompressorV2ABI},

		// Pool
		"CreditManager":           &bind.MetaData{ABI: creditManager.CreditManagerABI},
		"CreditManagerv2":         &bind.MetaData{ABI: creditManagerv2.CreditManagerv2ABI},
		"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
		"CreditFilter":            &bind.MetaData{ABI: creditFilter.CreditFilterABI},
		"LidoMock":                &bind.MetaData{ABI: lidoMock.LidoMockABI},
		"MainnetLido":             &bind.MetaData{ABI: mainnetLido.MainnetLidoABI},
		"Pool":                    poolService.PoolServiceMetaData,

		// GetUnderlyingToken
		"DieselToken":        dieselToken.DieselTokenMetaData,
		"GearToken":          gearToken.GearTokenMetaData,
		"TokenMock":          tokenMock.TokenMockMetaData,
		"RewardPool":         &bind.MetaData{ABI: rewardPool.RewardPoolABI},
		"Token":              eRC20.ERC20MetaData,
		"Uniswapv2Pool":      &bind.MetaData{ABI: uniswapv2Pool.Uniswapv2PoolABI},
		"Uniswapv3Pool":      &bind.MetaData{ABI: uniswapv3Pool.Uniswapv3PoolABI},
		"Uniswapv2Router":    &bind.MetaData{ABI: uniswapv2Router.Uniswapv2RouterABI},
		"Uniswapv2Factory":   &bind.MetaData{ABI: uniswapv2Factory.Uniswapv2FactoryABI},
		"PriceOraclev2":      &bind.MetaData{ABI: priceOraclev2.PriceOraclev2ABI},
		"CreditConfigurator": &bind.MetaData{ABI: creditConfigurator.CreditConfiguratorABI},
		"CreditFacade":       &bind.MetaData{ABI: creditFacade.CreditFacadeABI},
		"MultiCall":          &bind.MetaData{ABI: multicall.MulticallABI},
		"LidoKovan":          &bind.MetaData{ABI: lidoKovan.LidoKovanABI},
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
