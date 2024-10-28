package core

import (
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/aCL"
	"github.com/Gearbox-protocol/sdk-go/artifacts/aCLTrait"
	"github.com/Gearbox-protocol/sdk-go/artifacts/accountFactory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/adaptersv3/savingMaker"
	"github.com/Gearbox-protocol/sdk-go/artifacts/addressProvider"
	"github.com/Gearbox-protocol/sdk-go/artifacts/contractsRegister"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccount"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccountCompressor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfiguratorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dieselToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/eRC20"
	"github.com/Gearbox-protocol/sdk-go/artifacts/gearToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/inchFarmingPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolQuotaKeeperv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/redstone"
	"github.com/Gearbox-protocol/sdk-go/artifacts/updatePriceFeed"
	"github.com/Gearbox-protocol/sdk-go/log"

	// v1
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFilter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManager"
	mainnetDC "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/tokenMock"

	// v2
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditConfigurator210"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacadeExtended"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"

	// adapter
	"github.com/Gearbox-protocol/sdk-go/artifacts/addrProviderv310"
	"github.com/Gearbox-protocol/sdk-go/artifacts/baseRewardPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidoKovan"
	"github.com/Gearbox-protocol/sdk-go/artifacts/lidoMock"
	"github.com/Gearbox-protocol/sdk-go/artifacts/linearInterestRateModel"
	"github.com/Gearbox-protocol/sdk-go/artifacts/mainnetLido"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/nonFungiblePosManager"
	"github.com/Gearbox-protocol/sdk-go/artifacts/tokenDistributor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapConnectorChecker"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/wETHGateway"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnAdapter"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"

	// v3
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacadev3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/creditManagerv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolv3"

	//
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var inchOracleABI = "[{\"inputs\":[{\"internalType\":\"contract IERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contract IERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var curveBalanceABI = "[{\"name\":\"balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"i\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2943}]"

var creditFacadev3MulticallABI = "[{\"type\":\"function\",\"name\":\"onDemandPriceUpdate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]"

func GetAbi(contractName string) *abi.ABI {
	abiMap := map[string]string{
		"Version":                 versionABI,
		"Pauser":                  pauserABI,
		"1InchOracle":             inchOracleABI,
		"curveBalance":            curveBalanceABI,
		"CreditFacadev3Multicall": creditFacadev3MulticallABI,
	}
	if data := abiMap[contractName]; data != "" {
		return getABI(data)
	}
	fn := func(abiM AbiMap, contractName string) *abi.ABI {
		abiStr, ok := abiM[contractName]
		if !ok {
			return nil
		}
		if abi, err := abiStr.GetAbi(); err != nil {
			log.Fatalf("Cant get ABI for %s: %s", contractName, err)
		} else {
			return abi
		}
		return nil
	}

	if abi := fn(v1Map, contractName); abi != nil {
		return abi
	}
	if abi := fn(v2Map, contractName); abi != nil {
		return abi
	}
	if abi := fn(v3Map, contractName); abi != nil {
		return abi
	}
	if abi := fn(v310Map, contractName); abi != nil {
		return abi
	}
	if abi := fn(adapterMap, contractName); abi != nil {
		return abi
	}
	log.Fatalf("Abi for %s not found", contractName)
	return nil
}

type AbiMap map[string]*bind.MetaData

var v1Map = AbiMap{
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
	"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
	"CreditFilter":            {ABI: creditFilter.CreditFilterABI},
	"Pool":                    poolService.PoolServiceMetaData,
}
var v2Map = AbiMap{
	"DataCompressorV2": {ABI: dataCompressorv2.DataCompressorv2ABI},

	"CreditManagerv2":       {ABI: creditManagerv2.CreditManagerv2ABI},
	"PriceOraclev2":         {ABI: priceOraclev2.PriceOraclev2ABI},
	"CreditConfigurator":    {ABI: creditConfigurator.CreditConfiguratorABI},
	"CreditConfigurator210": {ABI: creditConfigurator210.CreditConfigurator210ABI},
	"CreditFacade":          {ABI: creditFacade.CreditFacadeABI},
	"CreditFacadeExtended":  {ABI: creditFacadeExtended.CreditFacadeExtendedABI},
	// multicall
	"MultiCall": {ABI: multicall.MulticallABI},
}
var v3Map = AbiMap{
	"CreditManagerv3":      {ABI: creditManagerv3.CreditManagerv3ABI},
	"CreditFacadev3":       {ABI: creditFacadev3.CreditFacadev3ABI},
	"CreditConfiguratorv3": {ABI: creditConfiguratorv3.CreditConfiguratorv3ABI},
	"PoolQuotaKeeper":      {ABI: poolQuotaKeeperv3.PoolQuotaKeeperv3ABI},
	"Poolv3":               {ABI: poolv3.Poolv3ABI},
	"DataCompressorv3":     {ABI: dataCompressorv3.DataCompressorv3ABI},
	"InchFarming":          {ABI: inchFarmingPool.InchFarmingPoolABI},
	"PriceOraclev3":        {ABI: priceOraclev3.PriceOraclev3ABI},
	"UpdatePriceFeed":      {ABI: updatePriceFeed.UpdatePriceFeedABI},
	"RedStone":             {ABI: redstone.RedstoneABI},
}
var v310Map = AbiMap{
	"CreditAccountCompressor": {ABI: creditAccountCompressor.CreditAccountCompressorABI},
	"AddressProviderv310":     {ABI: addrProviderv310.Addrv310ABI},
}
var adapterMap = AbiMap{
	//
	"Uniswapv3Pool":    {ABI: uniswapv3Pool.Uniswapv3PoolABI},
	"Uniswapv3Router":  {ABI: uniswapv3Router.Uniswapv3RouterABI},
	"Uniswapv3Factory": {ABI: uniswapv3Factory.Uniswapv3FactoryABI},
	//
	"Uniswapv2Pool":    {ABI: uniswapv2Pool.Uniswapv2PoolABI},
	"Uniswapv2Router":  {ABI: uniswapv2Router.Uniswapv2RouterABI},
	"Uniswapv2Factory": {ABI: uniswapv2Factory.Uniswapv2FactoryABI},
	//
	"LidoMock":                {ABI: lidoMock.LidoMockABI},
	"MainnetLido":             {ABI: mainnetLido.MainnetLidoABI},
	"LidoKovan":               {ABI: lidoKovan.LidoKovanABI},
	"UniswapConnectorChecker": {ABI: uniswapConnectorChecker.UniswapConnectorCheckerABI},
	//
	"CurvePool": {ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"},
	//
	"YearnAdapter": {ABI: yearnAdapter.YearnAdapterABI},
	// convex for liquidator calculation
	"BaseRewardPool": {ABI: baseRewardPool.BaseRewardPoolABI},
	// v3 pos for synctron
	"NonFungiblePosManager": {ABI: nonFungiblePosManager.NonFungiblePosManagerABI},
	"TokenDistributor":      {ABI: tokenDistributor.TokenDistributorABI},
	//
	//
	"MakerDAI": {ABI: savingMaker.SavingMakerABI},
}

func getABI(data string) *abi.ABI {
	abi, err := abi.JSON(strings.NewReader(data))
	log.CheckFatal(err)
	return &abi
}
