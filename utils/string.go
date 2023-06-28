package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/uniswapv2Router"
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

func ToJsonBytes(obj interface{}) []byte {
	str, err := json.Marshal(obj)
	log.CheckFatal(err)
	return str
}
func ToJson(obj interface{}) string {
	return string(ToJsonBytes(obj))
}

func Method() {
	if a, err := abi.JSON(strings.NewReader(uniswapv2Router.Uniswapv2RouterABI)); err == nil {
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
func GetEnvOrPanic(field string) string {
	ans := os.Getenv(field)
	if ans == "" {
		log.Fatalf("env %s not present", field)
	}
	return ans
}

func getField(envVarDS reflect.StructField) (string, string) {
	envField := envVarDS.Tag.Get("env")
	defaultValue := envVarDS.Tag.Get("default")
	value := strings.Replace(GetEnvOrDefault(envField, defaultValue), "\\n", "\n", -1)
	if envField != "" && // if the field doesn't have env tag, it's not an env var
		envVarDS.Type.Kind() != reflect.String && defaultValue == "" {
		log.Fatalf("Default value can't be missing if type golang isn't string. Env: %s, Type: %s", envField, envVarDS.Type.Kind())
	}
	return envField, value
}
func ReadFromEnv(val interface{}) {
	rv := reflect.ValueOf(val).Elem()
	num := rv.NumField()
	for i := 0; i < num; i++ {
		envVarDS := rv.Type().Field(i)
		envField, value := getField(envVarDS)
		if envField == "" {
			continue
		}
		switch envVarDS.Type.Kind() {
		case reflect.String:
			rv.Field(i).SetString(value)
		case reflect.Int64, reflect.Int32, reflect.Int:
			num, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Fatalf("Err(%v) while getting env %s", err, envField)
			}
			rv.Field(i).SetInt(num)
		case reflect.Bool:
			value = strings.ToLower(value)
			if value == "1" || value == "true" {
				rv.Field(i).SetBool(true)
			} else if value == "0" || value == "false" {
				rv.Field(i).SetBool(false)
			} else {
				log.Fatalf("Wrong value (%s) gotten for %s bool", value, envField)
			}
		case reflect.Float64:
			f, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Fatalf("Err(%s) while getting env float %s", err, envField)
			}
			rv.Field(i).SetFloat(f)
		}
	}
}

type Errors []error

func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}
	var s string
	for i, err := range e {
		s += fmt.Sprintf("Err %d: %s\n", i, err.Error())
	}
	return s
}
