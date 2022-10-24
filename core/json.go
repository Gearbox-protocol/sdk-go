package core

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

type Json map[string]interface{}

func (j *Json) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *Json) Scan(value interface{}) error {
	out := map[string]interface{}{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = Json(out)
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = Json(out)
		return err
	default:
		return fmt.Errorf("could not scan type %T", t)
	}
}

//
type JsonBigIntMap map[string]*BigInt

func (j *JsonBigIntMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (z *JsonBigIntMap) Scan(value interface{}) error {
	out := JsonBigIntMap{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = out
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = out
		return err
	default:
		return fmt.Errorf("could not scan type %T", t)
	}
}

type JsonFloatMap map[string]float64

func (z *JsonFloatMap) Value() (driver.Value, error) {
	return json.Marshal(z)
}

func (z JsonFloatMap) Copy() *JsonFloatMap {
	obj := JsonFloatMap{}
	for token, amount := range z {
		obj[token] = amount
	}
	return &obj
}

func (z JsonFloatMap) ValueInUSD(prices JsonFloatMap) (valueInUSD float64) {
	for token, amount := range z {
		valueInUSD += amount * prices[token]
	}
	return
}

func (z *JsonFloatMap) Scan(value interface{}) error {
	out := JsonFloatMap{}
	switch t := value.(type) {
	case string:
		err := json.Unmarshal([]byte(value.(string)), &out)
		*z = out
		return err
	case []byte:
		err := json.Unmarshal(value.([]byte), &out)
		*z = out
		return err
	default:
		return fmt.Errorf("could not scan type %T", t)
	}
}

func (addrs AddressMap) checkIfAddress(v string) string {
	if strings.HasPrefix(v, "#") {
		key := strings.Trim(v, "#")
		if addrs[key] == "" {
			addrs[key] = utils.RandomAddr()
		}
		return addrs[key]
	} else if strings.HasPrefix(v, "!#") {
		key := strings.Trim(v, "!#")
		if addrs[key] == "" {
			addrs[key] = utils.RandomHash()
		}
		return addrs[key]
	} else {
		return v
	}
}

type AddressMap map[string]string

func (addrs AddressMap) checkInterface(data interface{}, t *testing.T) interface{} {
	switch value := data.(type) {
	case string:
		return addrs.checkIfAddress(value)
	case []interface{}:
		for i, entry := range value {
			value[i] = addrs.checkInterface(entry, t)
		}
		return value
	case map[string]interface{}:
		for key, entry := range value {
			newKey := addrs.checkIfAddress(key)
			if newKey != key {
				delete(value, key)
				key = newKey
			}
			value[key] = addrs.checkInterface(entry, t)
		}
		return value
	default:
		return data
	}
}

func (z *Json) ParseAddress(t *testing.T, addrs AddressMap) {
	for key, data := range *z {
		newKey := addrs.checkIfAddress(key)
		if newKey != key {
			delete(*z, key)
			key = newKey
		}
		(*z)[key] = addrs.checkInterface(data, t)
	}
}

func (z *Json) UnmarshalJSON(b []byte) (err error) {
	tmpMap := map[string]interface{}{}
	err = json.Unmarshal(b, &tmpMap)
	if err == nil {
		(*z) = tmpMap
	}
	return
}

func (addrs AddressMap) ReplaceWithVariable(key string, data interface{}) interface{} {
	if key == "sessionId" {
		value, ok := data.(string)
		if !ok {
			log.Error("string parsing failed")
		}
		splits := strings.Split(value, "_")
		splits[0] = addrs[splits[0]]
		return strings.Join(splits, "_")
	}
	switch value := data.(type) {
	case string:
		answer := addrs[value]
		if answer != "" {
			return answer
		}
	case []interface{}:
		for i, entry := range value {
			value[i] = addrs.ReplaceWithVariable(key, entry)
		}
		return value
	case map[string]interface{}:
		obj := Json(value)
		obj.ReplaceWithVariable(addrs)
		return obj
	default:
		return data
	}
	return data
}

func (z *Json) ReplaceWithVariable(addrToVariable AddressMap) {
	for key, value := range *z {
		if value == nil {
			delete(*z, key)
			continue
		}
		if addrToVariable[key] != "" {
			delete(*z, key)
			variable := addrToVariable[key]
			(*z)[variable] = addrToVariable.ReplaceWithVariable(key, value)
		} else {
			(*z)[key] = addrToVariable.ReplaceWithVariable(key, value)
		}
	}
}

func (z *Json) CheckSumAddress() {
	for key, value := range *z {
		(*z)[key] = fixAddress(value)
	}
}

func fixAddress(data interface{}) interface{} {
	switch value := data.(type) {
	case common.Address:
		return value.Hex()
	case string:
		if len(value) == 42 && value[:2] == "0x" {
			log.Info(value)
			return common.HexToAddress(value).Hex()
		}
	case []interface{}:
		for i, entry := range value {
			value[i] = fixAddress(entry)
		}
		return value
	case map[string]interface{}:
		value, ok := data.(map[string]interface{})
		if !ok {
			log.Error("map[string]interface{} parsing failed")
		}
		for key, entry := range value {
			value[key] = fixAddress(entry)
		}
		return value
	}
	return data
}
