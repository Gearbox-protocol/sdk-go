package schemas

import (
	"time"

	"github.com/Gearbox-protocol/sdk-go/core"
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

var s = map[string]int64{
	"0xcF64698AFF7E5f27A11dff868AF228653ba53be0": 13810899,  // mainnet
	"0x7d04ecdb892ae074f03b5d0aba03796f90f3f2af": 184650310, // arbitrum
	"0x3761ca4bfacfcffc1b8034e69f19116dd6756726": 118410666, // optimism
	"0x4b27b296273B72d7c7bfee1ACE93DC081467C41B": 9779380,   //s onic
	"0xF7f0a609BfAb9a0A98786951ef10e5FE26cC1E38": 48761804,  // bnb
}

func NewContract(address, contractName string, discoveredAt int64, client core.ClientI) *Contract {

	con := &Contract{
		ContractName: contractName,
		DiscoveredAt: discoveredAt,
		Address:      address,
		Client:       client,
	}
	if discoveredAt == -1 {
		discoveredAt = s[address]
	}
	if discoveredAt == 0 {
		log.Fatal("discoveredAt is not set", address)
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

	// log.Debugf("Discovering first log of: %s\n", s.Address)
	// lastBlock, err := c.Client.BlockNumber(context.Background())
	// if err != nil {
	// 	log.Fatal("Cant get last block at discovery " + err.Error())
	// }
	if utils.GetEnvOrDefault("ETHERSCAN_API_KEY", "") != "" {
		block, err := core.GetEtherscanFirstLog(core.GetBaseChainId(c.Client), common.HexToAddress(c.Address))
		if err != nil {
			log.Warnf("DiscoverFirstLog: GetEtherscanFirstLog for %s error: %s. Set to discoveredAt %d", c.Address, err, discoveredAt)
			block = utils.Max(discoveredAt-100_000, 1)
		}
		return block
	}

	FirstLogAt, err := c.findFirstLogBound(utils.Max(discoveredAt-100_000, 1), discoveredAt)
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
		if core.EthLogErrorCheck(err, c.Client) {
			middle := (fromBlock + toBlock) / 2

			log.Debugf("FirstLog %d %d %d", fromBlock, middle-1, toBlock)
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
