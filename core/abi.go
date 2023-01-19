package core

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/aCL"
	"github.com/Gearbox-protocol/sdk-go/artifacts/aCLTrait"
	"github.com/Gearbox-protocol/sdk-go/artifacts/accountFactory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/addressProvider"
	"github.com/Gearbox-protocol/sdk-go/artifacts/baseRewardPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/contractsRegister"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccount"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacadeExtended"
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

var inchOracleABI = "[{\"inputs\":[{\"internalType\":\"contract IERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contract IERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func GetAbi(contractName string) *abi.ABI {
	abiMap := map[string]string{
		"Version":     versionABI,
		"Pauser":      pauserABI,
		"1InchOracle": inchOracleABI,
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
		"YearnPriceFeed":        {ABI: yearnPriceFeed.YearnPriceFeedABI},
		"PriceFeed":             {ABI: priceFeed.PriceFeedABI},
		"DataCompressorMainnet": {ABI: mainnetDC.DataCompressorABI},
		"DataCompressorV2":      {ABI: dataCompressorv2.DataCompressorv2ABI},

		// tokens
		"DieselToken": dieselToken.DieselTokenMetaData,
		"GearToken":   gearToken.GearTokenMetaData,
		"TokenMock":   tokenMock.TokenMockMetaData,
		"Token":       eRC20.ERC20MetaData,

		//////////////
		// v1
		//////////////
		"PriceOracle":             {ABI: priceOracle.PriceOracleABI},
		"CreditManager":           {ABI: creditManager.CreditManagerABI},
		"CreditManagerv2":         {ABI: creditManagerv2.CreditManagerv2ABI},
		"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
		"CreditFilter":            {ABI: creditFilter.CreditFilterABI},
		"Pool":                    poolService.PoolServiceMetaData,

		//////////////
		// v2
		//////////////
		"PriceOraclev2":        {ABI: priceOraclev2.PriceOraclev2ABI},
		"CreditConfigurator":   {ABI: creditConfigurator.CreditConfiguratorABI},
		"CreditFacade":         {ABI: creditFacade.CreditFacadeABI},
		"CreditFacadeExtended": {ABI: creditFacadeExtended.CreditFacadeExtendedABI},
		// multicall
		"MultiCall": {ABI: multicall.MulticallABI},
		//
		"Uniswapv3Pool":    {ABI: uniswapv3Pool.Uniswapv3PoolABI},
		"Uniswapv3Router":  {ABI: uniswapv3Router.Uniswapv3RouterABI},
		"Uniswapv3Factory": {ABI: uniswapv3Factory.Uniswapv3FactoryABI},
		//
		"Uniswapv2Pool":    {ABI: uniswapv2Pool.Uniswapv2PoolABI},
		"Uniswapv2Router":  {ABI: uniswapv2Router.Uniswapv2RouterABI},
		"Uniswapv2Factory": {ABI: uniswapv2Factory.Uniswapv2FactoryABI},
		//
		"LidoMock":    {ABI: lidoMock.LidoMockABI},
		"MainnetLido": {ABI: mainnetLido.MainnetLidoABI},
		"LidoKovan":   {ABI: lidoKovan.LidoKovanABI},
		//
		"YearnAdapter": {ABI: yearnAdapter.YearnAdapterABI},
		// convex for liquidator calculation
		"BaseRewardPool": {ABI: baseRewardPool.BaseRewardPoolABI},
		// v3 pos for synctron
		"NonFungiblePosManager": {ABI: nonFungiblePosManager.NonFungiblePosManagerABI},
	}
	abiStr, ok := metadataMap[contractName]
	if !ok {
		panic("ABI  doesn't exists for " + contractName)
	}

	abi, err := abiStr.GetAbi()
	if err != nil {
		log.Infof("Cant get ABI for %s" + contractName)
		log.Fatal(err)
	}

	return abi
}

func getABI(data string) *abi.ABI {
	abi, err := abi.JSON(strings.NewReader(data))
	log.CheckFatal(err)
	return &abi
}
