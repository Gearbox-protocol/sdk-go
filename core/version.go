package core

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var versionABI string = "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// if version is not set it is 1 else get from contract
func FetchVersion(addr string, blockNum int64, client ClientI) VersionType {
	if common.HexToAddress(addr).Hex() == "0x39E6C2E1757ae4354087266E2C3EA9aC4257C1eb" { // bcz https://optimistic.etherscan.io/address/0x39E6C2E1757ae4354087266E2C3EA9aC4257C1eb#readContract has version as string
		return NewVersion(1)
	}
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

type VersionType struct {
	v int16
}

func NewVersion(v int16) VersionType {
	if v == 1 {
		return VersionType{v: 1}
	} else if v == 300 || v == 301 || v == 302 {
		return VersionType{v: 300}
	} else if v == 210 || v == 220 || v == 2 {
		return VersionType{v: 2}
	} else if v == 10_000 { // for testing
		return VersionType{v: 10000}
	}
	log.Fatal("version not supported", v)
	panic("")
}

func (v VersionType) IsGBv1() bool {
	return v.v == 1
}

func (v VersionType) IsPriceInUSD() bool {
	return v.v != 1
}

func (v VersionType) MoreThan(cmpAgainst VersionType) bool {
	return v.v > cmpAgainst.v
}
func (v VersionType) LessThan(cmpAgainst VersionType) bool {
	return v.v < cmpAgainst.v
}
func (v VersionType) MoreThanEq(cmpAgainst VersionType) bool {
	return v.v >= cmpAgainst.v
}

func (z VersionType) Value() (driver.Value, error) {
	return fmt.Sprintf("%d", z.v), nil
}
func (z *VersionType) Scan(value interface{}) error {
	if value == nil {
		z = nil
	}
	switch t := value.(type) {
	case string:
		v, err := strconv.ParseInt(t, 10, 32)
		log.CheckFatal(err)
		*z = NewVersion(int16(v))
	case int:
		*z = NewVersion(int16(t))
	case int64:
		*z = NewVersion(int16(t))
	default:
		return fmt.Errorf("could not scan type %T into versionType", t)
	}

	return nil
}

func (z VersionType) Decimals() int8{
	if z.Eq(1) {
		return 18
	} else if z.MoreThan(NewVersion(1)) {
		return 8
	}
	log.Fatal("")
	return 0
}

func (z VersionType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(z.v))), nil
}

func (z *VersionType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")
	v, err := strconv.Atoi(str)
	if err != nil {
		return fmt.Errorf("can unmarshal versiontype %s", err)
	}

	z.v = int16(v)
	return nil
}

// https://stackoverflow.com/questions/55335296/problem-with-marshal-unmarshal-when-key-of-map-is-a-struct for map
func (z VersionType) MarshalText() (text []byte, err error) {
	return z.MarshalJSON()
}

func (s *VersionType) UnmarshalText(text []byte) error {
	return s.UnmarshalJSON(text)
}

func (v VersionType) Eq(in int16) bool {
	return v.v == in
}
func FetchVersionOptimized(addr common.Address, blockNum int64, client ClientI) VersionType {
	_version, err := CallFuncWithExtraBytes(client, "54fd4d50", addr, blockNum, nil)
	if err != nil {
		return NewVersion(1)
	}
	version := new(big.Int).SetBytes(_version).Int64()
	if version == 0 {
		version = 1
	}
	return NewVersion(int16(version))
}
