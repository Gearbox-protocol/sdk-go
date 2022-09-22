package test

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *TestEvent) Process(contractName string) types.Log {
	topic0 := core.Topic(c.Topics[0])
	// for data
	data, err := c.ParseData([]string{contractName}, topic0)
	if err != nil {
		log.Fatal(c.Topics[0], err, contractName)
	}
	// for topics
	c.Topics[0] = topic0.Hex()
	var topics []common.Hash
	for _, value := range c.Topics {
		splits := strings.Split(value, ":")
		var newTopic string
		if len(splits) == 1 {
			newTopic = value
		} else {
			switch splits[0] {
			case "bigint":
				arg, ok := new(big.Int).SetString(splits[1], 10)
				if !ok {
					log.Fatalf("bigint parsing failed for %s", value)
				}
				newTopic = fmt.Sprintf("%x", arg)
			}
		}
		topics = append(topics, common.HexToHash(newTopic))
	}
	return types.Log{
		Data:    data,
		Topics:  topics,
		Address: common.HexToAddress(c.Address),
		TxHash:  common.HexToHash(c.TxHash),
	}
}

func (c *TestEvent) ParseData(contractName []string, topic0 common.Hash) ([]byte, error) {
	if len(c.Data) == 0 {
		return []byte{}, nil
	}
	if contractName[0] == "ACL" {
		contractName = append(contractName, "ACLTrait")
	}
	var event *abi.Event
	var err error
	for _, name := range contractName {
		abi := core.GetAbi(name)
		event, err = abi.EventByID(topic0)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatal(err, topic0)
	}
	var args []interface{}
	for _, entry := range c.Data {
		args = append(args, dataToArg(entry))
	}
	return event.Inputs.NonIndexed().Pack(args...)
}

func dataToArg(entry string) (arg interface{}) {
	splits := strings.Split(entry, ":")
	if len(splits) == 2 {
		var ok bool
		switch splits[0] {
		case "bigint":
			arg, ok = new(big.Int).SetString(splits[1], 10)
			if !ok {
				log.Fatalf("bigint parsing failed for %s", entry)
			}
		case "uint16":
			value, err := strconv.ParseUint(splits[1], 10, 16)
			log.CheckFatal(err)
			arg = uint16(value)
		case "uint8":
			value, err := strconv.ParseUint(splits[1], 10, 16)
			log.CheckFatal(err)
			arg = uint8(value)
		case "bool":
			if splits[1] == "1" {
				arg = true
			} else {
				arg = false
			}
		}
	} else {
		arg = common.HexToAddress(entry)
	}
	return arg
}

type TestEvent struct {
	Address string   `json:"address"`
	Data    []string `json:"data"`
	Topics  []string `json:"topics"`
	TxHash  string   `json:"txHash"`
}
