package schemas

import (
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
	dc "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dieselToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/eRC20"
	"github.com/Gearbox-protocol/sdk-go/artifacts/gearToken"
	"github.com/Gearbox-protocol/sdk-go/artifacts/linearInterestRateModel"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolService"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/rewardPool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/tokenMock"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv3Pool"
	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Factory"
	"github.com/Gearbox-protocol/sdk-go/artifacts/wETHGateway"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
)

const MaxUint = ^int64(0)
var VersionABI   abi.ABI
type Contract struct {
	DiscoveredAt int64        `gorm:"column:discovered_at" json:"discoveredAt"`
	FirstLogAt   int64        `gorm:"column:firstlog_at" json:"firstLogAt"`
	Address      string       `gorm:"primaryKey;column:address" json:"address"`
	Disabled     bool         `gorm:"column:disabled" json:"disabled"`
	ContractName string       `gorm:"column:type" json:"type"`
	Client       core.ClientI `gorm:"-" json:"-"`
	ABI          *abi.ABI     `gorm:"-" json:"-"`
	// VersionABI   abi.ABI      `gorm:"-" json:"-"`
}

func (c *Contract) Disable() {
	c.Disabled = true
}
func init() {
	abiStr := "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	versionABI, err := abi.JSON(strings.NewReader(abiStr))
	log.CheckFatal(err)
	VersionABI = versionABI
}
func NewContract(address, contractName string, discoveredAt int64, client core.ClientI) *Contract {

	con := &Contract{
		ContractName: contractName,
		DiscoveredAt: discoveredAt,
		Address:      address,
		Client:       client,
	}
	con.FirstLogAt = con.DiscoverFirstLog()
	// for address provider discoveredAt is -1
	if discoveredAt == -1 {
		con.DiscoveredAt = con.FirstLogAt
	} else {
		con.DiscoveredAt = discoveredAt
	}

	return con
}

func GetAbi(contractName string) *abi.ABI {

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
		"PriceOracle":    &bind.MetaData{ABI: priceOracle.PriceOracleABI},
		"YearnPriceFeed": &bind.MetaData{ABI: yearnPriceFeed.YearnPriceFeedABI},
		"PriceFeed": &bind.MetaData{ABI: priceFeed.PriceFeedABI},
		"DataCompressorV1": &bind.MetaData{ABI: dc.DataCompressorABI},

		// Pool
		"CreditManager":           &bind.MetaData{ABI: creditManager.CreditManagerABI},
		"CreditManagerV2":           &bind.MetaData{ABI: creditManagerv2.CreditManagerv2ABI},
		"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
		"CreditFilter":            &bind.MetaData{ABI: creditFilter.CreditFilterABI},
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
		"Uniswapv2Factory":    &bind.MetaData{ABI: uniswapv2Factory.Uniswapv2FactoryABI},
		"CreditConfigurator": &bind.MetaData{ABI: creditConfigurator.CreditConfiguratorABI},
		"CreditFacade":       &bind.MetaData{ABI: creditFacade.CreditFacadeABI},
		"MultiCall":          &bind.MetaData{ABI: multicall.MulticallABI},
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

// setter
func (c *Contract) GetAbi() {
	c.ABI = GetAbi(c.ContractName)
}
func (c *Contract) SetAddress(addr string) {
	c.Address = addr
}

// Getter

func (c *Contract) GetAddress() string {
	return c.Address
}

func (c *Contract) GetName() string {
	return c.ContractName
}
func (c *Contract) IsDisabled() bool {
	return c.Disabled
}
func (c *Contract) GetFirstLog() int64 {
	return c.FirstLogAt
}
func (c *Contract) GetDiscoveredAt() int64 {
	return c.DiscoveredAt
}

// Extras

func (c *Contract) DiscoverFirstLog() int64 {

	// log.Debugf("Discovering first log of: %s\n", s.Address)
	lastBlock, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Cant get last block at discovery " + err.Error())
	}

	FirstLogAt, err := c.findFirstLogBound(1, int64(lastBlock))
	if err != nil {
		log.Fatal(c.Address, err.Error())
	}

	return FirstLogAt
}

func (c *Contract) findFirstLogBound(fromBlock, toBlock int64) (int64, error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(c.Address),
		},
		Topics: [][]common.Hash{},
	}

	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == core.QueryMoreThan10000Error ||
			strings.Contains(err.Error(), core.LogFilterLenError) {
			middle := (fromBlock + toBlock) / 2

			log.Verbosef("FirstLog %d %d %d", fromBlock, middle-1, toBlock)
			foundLow, err := c.findFirstLogBound(fromBlock, middle-1)
			if err != nil && err.Error() != "no events found" {
				return 0, err
			}

			foundHigh, err := c.findFirstLogBound(middle, toBlock)
			if err != nil && err.Error() != "no events found" && err.Error() != "Cant find any events" {
				return 0, err
			}

			if foundLow == 0 && foundHigh == 0 {
				return 0, fmt.Errorf("No events was found for the contract")
			}

			if foundLow == 0 {
				return foundHigh, nil
			}

			return foundLow, nil

		}
		return 0, err
	}

	FirstLogAt := int64(0)

	for _, vLog := range logs {
		block := int64(vLog.BlockNumber)
		if block < FirstLogAt || FirstLogAt == 0 {
			FirstLogAt = block
		}
	}

	if FirstLogAt == MaxUint {
		return 0, fmt.Errorf("no events found")
	}

	return FirstLogAt, nil
}

func (c *Contract) FindLastLogBound(fromBlock, toBlock int64, topics []common.Hash) (int64, error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(c.Address),
		},
		Topics: [][]common.Hash{
			topics,
		},
	}
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == core.QueryMoreThan10000Error ||
			strings.Contains(err.Error(), core.LogFilterLenError) {
			middle := (fromBlock + toBlock) / 2
			foundHigh, err := c.FindLastLogBound(middle, toBlock, topics)
			if err != nil {
				return 0, err
			}
			if foundHigh != 0 {
				return foundHigh, nil
			}
			foundLow, err := c.FindLastLogBound(fromBlock, middle-1, topics)
			if err != nil {
				return 0, err
			}
			if foundLow != 0 {
				return foundLow, nil
			}
		}
		return 0, err
	}

	logLen := len(logs)
	if logLen > 0 {
		return int64(logs[logLen-1].BlockNumber), nil
	} else {
		return 0, nil
	}
}

func (c *Contract) UnpackLogIntoMap(out map[string]interface{}, event string, txLog types.Log) error {
	if txLog.Topics[0] != c.ABI.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(txLog.Data) > 0 {
		if err := c.ABI.UnpackIntoMap(out, event, txLog.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopicsIntoMap(out, indexed, txLog.Topics[1:])
}

func (c *Contract) ParseEvent(eventName string, txLog *types.Log) (string, *core.Json) {
	data := map[string]interface{}{}
	if eventName == "TransferAccount" && len(txLog.Data) > 0 {
		data = map[string]interface{}{
			"oldOwner": common.BytesToAddress(txLog.Data[:32]).Hex(),
			"newOwner": common.BytesToAddress(txLog.Data[32:]).Hex(),
		}
	} else {
		if err := c.UnpackLogIntoMap(data, eventName, *txLog); err != nil {
			log.Fatal(err)
		}
	}
	// add order
	var argNames []interface{}
	for _, input := range c.ABI.Events[eventName].Inputs {
		argNames = append(argNames, input.Name)
	}
	data["_order"] = argNames
	jsonData := core.Json(data)
	jsonData.CheckSumAddress()
	return c.ABI.Events[eventName].Sig, &jsonData
}

func FetchVersion(addr string, blockNum int64, client core.ClientI) int16 {
	var opts *bind.CallOpts
	if blockNum != 0 {
		opts = &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
	}
	contract := bind.NewBoundContract(common.HexToAddress(addr), VersionABI, client, client, client)
	var out []interface{}
	err := contract.Call(opts, &out, "version")
	if err != nil {
		return 1
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return int16(out0.Int64())
}
