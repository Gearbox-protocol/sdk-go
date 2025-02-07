package dataCompressorv3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (_DataCompressorv3 *DataCompressorv3Caller) GetCreditAccountData(opts *bind.CallOpts, _cm common.Address, priceUpdates []PriceOnDemand) (CreditAccountData, error) {
	var out []interface{}
	err := _DataCompressorv3.contract.Call(opts, &out, "getCreditAccountData", _cm, priceUpdates)

	if err != nil {
		return *new(CreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditAccountData)).(*CreditAccountData)

	return out0, err

}
