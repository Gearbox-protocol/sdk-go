package schemas

import (
	"math"
	"strings"
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"context"
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
)

const MaxUint = ^int64(0)

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

var s = map[int64]map[string]int64{
	1:     {"0xcF64698AFF7E5f27A11dff868AF228653ba53be0": 13810899},  // mainnet
	42161: {"0x7d04ecdb892ae074f03b5d0aba03796f90f3f2af": 184650310}, // arbitrum
	10:    {"0x3761ca4bfacfcffc1b8034e69f19116dd6756726": 118410666}, // optimism
	146:   {"0x4b27b296273B72d7c7bfee1ACE93DC081467C41B": 9779380},   // sonic
	42793: {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 16672969},  // etherlink
	56:    {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 48761804},  // bnb
	1135:  {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 18369936},  // lisk
	43111: {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 2133492},   // lisk
	9745: {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 670918,
		"0x274785849595F3f9669c44FD53EbCAec4f7ED03C": 2999860,
		"0x53E4e9b8766969c43895839CC9c673bb6bC8Ac97": 4276921,
		//
		"0xF8509f177b4d942b0268A25b5AA847864aCd8284": 4272207,
		"0xB74760FD26400030620027DD29D19d74D514700e": 4272207,
		"0xCC817FEe72F683a21B0ACe215889382cA24E6e67": 4032124,
		"0xBf4B23909014b7dd04e22cF3010b6e8F139fA2d3": 4180800,
		"0x052d92a9766E968342c5315DEEaC055C6f308ac2": 4277342,
		"0x089704AE81E15185Bc523999Bbb1cAb0d536fC2F": 4276977,
		"0xC6D4845f178A0F1B8d3d2298E35489FccbDf7000": 4276977,
		"0x02439e41047274C17A3385790433A6Aaf4Efa8E9": 4276921,
		"0xE70486DaAdAa38C54d876672508f3A73202cDF17": 4277342,
		"0x19eb08b867C752Dc4686Ea0FD26af1C072eD62f9": 4279728,
		"0x6377989B9Ef6fd485E28a6Fd66966f0CC8109F9d": 4288350,
		"0xd6B6c5897D8A4824573dEb1F3190186124d880B7": 4279728,
		"0x01e897f689Ec2Ab56E76AA2DD6860537810f86aF": 4288350,
		"0x9C64Da71EAc61C15Bc76F5c92b3451abE337022f": 4288752,
		"0xB202C10312722AB4F02F8cBaFF10Cc47Fc082F6A": 4288913,
		"0x11307a61853baBcF5De7F74E23fDf6645Ad45A2F": 4288913,
		"0xA6D16367333A40F6B9779C5EaD0d9EaA6759Ab9D": 4288752,
		"0xF0d022Af7616f7B861d1c767A543a5b5e1DC18EF": 4272207,
		"0xac1594ac2DCa4c78b231482C25F141BD1eCA637f": 4276921,
		"0xF9B9Ad7FE06fd74C39c79cafedeb21Aa0980391A": 4364688,
		"0xB327fF096Ff13F83a7ce839Af0F88540c8d952c9": 4364688,
		"0xC56EA16EA06B0a6A7b3B03B2f48751e549bE40fD": 4364798,
		//
		"0xC0fE79990D6b6372EE4700bC8FaC03bd0DB5FFcf": 2999860}, // "0xB1d8397bCb77018AE2EA47e210F86A7C03673414": 2160145,
	// "0x61F7f5875eC741Ed7321E7CDc70C7662C75c5a06": 2160145,
	// plasma
	143: {"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 34650265},
}
var deploymentdates = map[int64][]int64{9745: {670918, 1820915, 2999860, 4032124}}

func NewContract(address, contractName string, discoveredAt int64, client core.ClientI) *Contract {

	con := &Contract{
		ContractName: contractName,
		DiscoveredAt: discoveredAt,
		Address:      address,
		Client:       client,
	}
	if discoveredAt == -1 {
		if core.GetBaseChainId(client) == 42793 {
			if "0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38" == address {
				discoveredAt = s[core.GetBaseChainId(client)][address]
			} else {
				discoveredAt = 22699432
			}
		} else {
			discoveredAt = s[core.GetBaseChainId(client)][address]
		}
	}
	con.FirstLogAt = con.DiscoverFirstLog(discoveredAt)
	if con.FirstLogAt == 0 && core.GetChainId(client) != 1337 { //don't updateif testnet
		con.FirstLogAt = discoveredAt
	}
	con.DiscoveredAt = discoveredAt
	return con
}

// setter
func (c *Contract) GetAbi() {
	c.ABI = core.GetAbi(c.ContractName)
}
func (c *Contract) SetAddress(addr string) {
	c.Address = addr
}

// Getter

func (c *Contract) GetAddress() string {
	if c.Address == "" {
		log.Fatal("Adapter address is not set")
	}
	return c.Address
}

func (c *Contract) GetName() string {
	if c.ContractName == "" {
		log.Fatal("Contract name is not set")
	}
	return c.ContractName
}
func (c *Contract) IsDisabled() bool {
	return c.Disabled
}
func (c *Contract) GetDiscoveredAt() int64 {
	return c.DiscoveredAt
}

// Extras
func (c *Contract) DiscoverFirstLog(discoveredAt int64) int64 {
	if firstlog, ok := s[core.GetBaseChainId(c.Client)][c.Address]; ok {
		return firstlog
	}
	// log.Debugf("Discovering first log of: %s\n", s.Address)
	// lastBlock, err := c.Client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal("Cant get last block at discovery " + err.Error())
	// }
	// if core.GetBaseChainId(c.Client) == 42793 && discoveredAt == 0 && c.ContractName == "ContractRegister" {
	// 	discoveredAt = core.GetLatestBlockNumber(c.Client)

	// }
	network := log.GetNetworkName(core.GetBaseChainId(c.Client))
	if utils.GetEnvOrDefault("ETHERSCAN_API_KEY", "") == "" && discoveredAt == 0 &&
		!utils.Contains([]log.NETWORK{log.ETHERLINK, log.LISK, log.HEMIBTC, log.PLASMA, log.ETHERLINK, log.MONAD}, network) { // on etherlink there is no etherscan api
		log.Fatal("discoveredAt is not set and can't use etherscan for getting etherscan", c.Address)
	}
	if utils.GetEnvOrDefault("ETHERSCAN_API_KEY", "") != "" {
		block, err := core.GetEtherscanFirstLog(core.GetBaseChainId(c.Client), common.HexToAddress(c.Address))
		if err != nil {
			if strings.Contains(err.Error(), "failed to get first log block num") {
				forkBlock := core.GetForkBlock(c.Client.(*ethclient.Client).GetUrl()) // update fork block
				if forkBlock < math.MaxInt64 {
					FirstLogAt, err := c.findFirstLogBound(forkBlock, core.GetLatestBlockNumber(c.Client))
					if err != nil {
						log.Fatal(c.Address, err.Error())
					}
					return FirstLogAt
				}
			} else {
				log.Fatal(c.Address, "logs from etherscan, but still failed", err.Error())
			}
		}
		return block
	}

	// FirstLogAt, err := c.findFirstLogBound(utils.Max(discoveredAt-100_000, 1), discoveredAt)
	latestBlock := core.GetLatestBlockNumber(c.Client)
	var end int64 = utils.Min(discoveredAt+100_000, latestBlock)
	if discoveredAt == 0 {
		end = latestBlock
	}
	var start = discoveredAt
	// get the deployment date , currently only for plasma
	for _, b := range deploymentdates[core.GetBaseChainId(c.Client)] {
		if discoveredAt > b {
			start = b
		}
	}
	if core.GetBaseChainId(c.Client) == 42793 {
		if c.Address != "0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38" {
			start = utils.Max(22699432, start)
		}
		if c.Address == "0x20766A7BE771301b097B28d5288AA62Ee28bd641" {
			return 22699432
		}
		if c.Address == "0x653e62A9Ef0e869F91Dc3D627B479592aA02eA75" {
			return 22699432
		}
	}
	FirstLogAt, err := c.findFirstLogBound(start, end)
	if err != nil {
		log.Fatal(c.Address, err.Error(), start, end)
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
		if core.EthLogErrorCheck(err, c.Client) {
			middle := (fromBlock + toBlock) / 2

			log.Infof("FirstLog for %s from %d, mid %d, to %d", c.Address, fromBlock, middle-1, toBlock)
			foundLow, err := c.findFirstLogBound(fromBlock, middle-1)
			if err != nil {
				return 0, err
			}
			if foundLow != 0 {
				return foundLow, nil
			}

			foundHigh, err := c.findFirstLogBound(middle, toBlock)
			return foundHigh, err
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	//
	logs, err := c.Client.FilterLogs(ctx, query)
	if err != nil {
		if core.EthLogErrorCheck(err, c.Client) {
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
	jsonData.QuoteBigInt()
	return c.ABI.Events[eventName].Sig, &jsonData
}
