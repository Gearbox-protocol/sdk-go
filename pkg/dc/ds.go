package dc

import (
	"reflect"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func GetPoolDataFromDCCall(data interface{}) (PoolCallData, error) {
	switch values := data.(type) {
	case dcv3.PoolData:
		return getPoolDatav3(values), nil
	case dcv2.PoolData:
		return getPoolDatav2(0, values), nil
	case mainnet.DataTypesPoolData:
		return getPoolDataV1(values), nil
	default:
		log.Fatalf("Can't process %s in poolData", reflect.TypeOf(data))
	}
	return PoolCallData{}, nil
}
func GetCMDataFromDCCall(data interface{}) (CMCallData, error) {
	switch values := data.(type) {
	case dcv3.CreditManagerData:
		return getCMDatav3(values), nil
	case dcv2.CreditManagerData:
		return getCMDatav2(values), nil
	case mainnet.DataTypesCreditManagerData:
		return getCMDatav1(values), nil
	default:
		log.Fatalf("Can't process %s in managerData", reflect.TypeOf(data))
	}
	return CMCallData{}, nil
}
func GetAccountDataFromDCCall(client core.ClientI, cfAddrv1 common.Address, blockNum int64, data interface{}) (CreditAccountCallData, error) {
	switch values := data.(type) {
	case dcv3.CreditAccountData:
		return getAccountDatav3(values), nil
	case dcv2.CreditAccountData:
		return getAccountDatav2(values), nil
	case mainnet.DataTypesCreditAccountDataExtended:
		return getCreditAccountDatav1(client, cfAddrv1, blockNum, values)
	default:
		log.Fatalf("Can't process %s in accountData", reflect.TypeOf(data))
	}
	return CreditAccountCallData{}, nil
}
