package core

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var versionABI string = "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// if version is not set it is 1 else get from contract
func FetchVersion(addr string, blockNum int64, client ClientI) VersionType {
	var opts *bind.CallOpts
	if blockNum != 0 {
		opts = &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
	}
	contract := bind.NewBoundContract(common.HexToAddress(addr), *GetAbi("Version"), client, client, client)
	var out []interface{}
	err := contract.Call(opts, &out, "version")
	if err != nil {
		return NewVersion(1)
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	return NewVersion(int16(out0.Int64()))
}

type VersionType int16

func NewVersion(v int16) VersionType {
	if v == 1 || v == 2 || v == 210 {
		return VersionType(v)
	}
	log.Fatal("version not supported")
	panic("")
}

func (v VersionType) Decimals() int8 {
	switch v {
	case 1:
		return 18 // eth decimals
	case 2, 210:
		return 8 // USD decimals
	default:
		log.Fatal("version not supported")
		panic("")
	}
}

func (v VersionType) IsGBv1() bool {
	return v == 1
}

func (v VersionType) IsPriceInUSD() bool {
	return v != 1
}

// used for not setting session balance from dc for account closure in v2 or above
func (v VersionType) IsGBv2orAbove() bool {
	return v >= 2
}

func (v VersionType) MoreThan(cmpAgainst VersionType) bool {
	return floatOrd(v) > floatOrd(cmpAgainst)
}
func floatOrd(v VersionType) float64 {
	if v < 10 {
		return float64(v)
	} else if v == 210 {
		return float64(v) / 100
	}
	panic("")
}
func FetchVersionOptimized(addr common.Address, blockNum int64, client ClientI) VersionType {
	_version, err := CallFuncWithExtraBytes(client, "54fd4d50", addr, blockNum, nil)
	if err != nil {
		return 1
	}
	version := new(big.Int).SetBytes(_version).Int64()
	if version == 0 {
		version = 1
	}
	return VersionType(version)
}
