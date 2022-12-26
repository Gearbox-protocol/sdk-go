package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditFacade"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ToJson(obj interface{}) string {
	str, err := json.Marshal(obj)
	log.CheckFatal(err)
	return string(str)
}

func Method() {
	if a, err := abi.JSON(strings.NewReader(creditFacade.CreditFacadeABI)); err == nil {
		for name, method := range a.Methods {
			log.Info(name, method.RawName, method.Name, method.Sig, hex.EncodeToString(method.ID))
		}
	}
}

// for testing and json file read methods
func RandomAddr() string {
	return common.HexToAddress(random(20)).Hex()
}
func RandomHash() string {
	return "0x" + random(32)
}
func random(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func ChecksumAddr(addr string) string {
	return common.HexToAddress(addr).Hex()
}

func ConvertToListOfString(list interface{}) (accountAddrs []string) {
	switch accountList := list.(type) {
	case []interface{}:
		// accountList, ok := list.([]interface{})
		// if !ok {
		// 	log.Fatal("parsing accounts list for token transfer failed")
		// }
		for _, account := range accountList {
			accountAddr, ok := account.(string)
			if !ok {
				log.Fatalf("parsing single account for token transfer failed %v", account)
			}
			accountAddrs = append(accountAddrs, accountAddr)
		}
	case []string:
		// accountList, ok := list.([]string)
		// if !ok {
		// 	log.Fatal("parsing accounts list for token transfer failed")
		// }
		accountAddrs = accountList
	default:
		log.Fatal("parsing accounts list for token transfer failed")
	}
	return
}

func GetEnvOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ReadFromEnv(val interface{}) {
	rv := reflect.ValueOf(val).Elem()
	num := rv.NumField()
	for i := 0; i < num; i++ {
		envValue := rv.Type().Field(i).Tag.Get("env")
		defaultValue := rv.Type().Field(i).Tag.Get("default")
		if envValue != "" {
			value := strings.Replace(GetEnvOrDefault(envValue, defaultValue), "\\n", "\n", -1)
			rv.Field(i).SetString(value)
		}
	}
}
