// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolCompressor

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BaseParams is an auto generated low-level Go binding around an user-defined struct.
type BaseParams struct {
	Addr             common.Address
	Version          *big.Int
	ContractType     [32]byte
	SerializedParams []byte
}

// BaseState is an auto generated low-level Go binding around an user-defined struct.
type BaseState struct {
	BaseParams BaseParams
}

// CreditManagerDebtParams is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerDebtParams struct {
	CreditManager common.Address
	Borrowed      *big.Int
	Limit         *big.Int
	Available     *big.Int
}

// PoolQuotaKeeperState is an auto generated low-level Go binding around an user-defined struct.
type PoolQuotaKeeperState struct {
	BaseParams          BaseParams
	RateKeeper          common.Address
	Quotas              []QuotaTokenParams
	CreditManagers      []common.Address
	LastQuotaRateUpdate *big.Int
}

// PoolState is an auto generated low-level Go binding around an user-defined struct.
type PoolState struct {
	BaseParams              BaseParams
	Symbol                  string
	Name                    string
	Decimals                uint8
	TotalSupply             *big.Int
	PoolQuotaKeeper         common.Address
	InterestRateModel       common.Address
	Underlying              common.Address
	AvailableLiquidity      *big.Int
	ExpectedLiquidity       *big.Int
	BaseInterestIndex       *big.Int
	BaseInterestRate        *big.Int
	DieselRate              *big.Int
	SupplyRate              *big.Int
	WithdrawFee             *big.Int
	TotalBorrowed           *big.Int
	TotalDebtLimit          *big.Int
	CreditManagerDebtParams []CreditManagerDebtParams
	BaseInterestIndexLU     *big.Int
	ExpectedLiquidityLU     *big.Int
	QuotaRevenue            *big.Int
	LastBaseInterestUpdate  *big.Int
	LastQuotaRevenueUpdate  *big.Int
	IsPaused                bool
}

// QuotaTokenParams is an auto generated low-level Go binding around an user-defined struct.
type QuotaTokenParams struct {
	Token             common.Address
	Rate              uint16
	CumulativeIndexLU *big.Int
	QuotaIncreaseFee  uint16
	TotalQuoted       *big.Int
	Limit             *big.Int
	IsActive          bool
}

// Rate is an auto generated low-level Go binding around an user-defined struct.
type Rate struct {
	Token common.Address
	Rate  uint16
}

// RateKeeperState is an auto generated low-level Go binding around an user-defined struct.
type RateKeeperState struct {
	BaseParams BaseParams
	Rates      []Rate
}

// PoolCompressorMetaData contains all meta data concerning the PoolCompressor contract.
var PoolCompressorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"}],\"name\":\"getInterestRateModelState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lossPolicy\",\"type\":\"address\"}],\"name\":\"getLossPolicyState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quotaKeeper\",\"type\":\"address\"}],\"name\":\"getPoolQuotaKeeperState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"rateKeeper\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"uint192\",\"name\":\"cumulativeIndexLU\",\"type\":\"uint192\"},{\"internalType\":\"uint16\",\"name\":\"quotaIncreaseFee\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"totalQuoted\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"limit\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structQuotaTokenParams[]\",\"name\":\"quotas\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"creditManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint40\",\"name\":\"lastQuotaRateUpdate\",\"type\":\"uint40\"}],\"internalType\":\"structPoolQuotaKeeperState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"poolQuotaKeeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"internalType\":\"structCreditManagerDebtParams[]\",\"name\":\"creditManagerDebtParams\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"baseInterestIndexLU\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidityLU\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quotaRevenue\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastBaseInterestUpdate\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"lastQuotaRevenueUpdate\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structPoolState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rateKeeper\",\"type\":\"address\"}],\"name\":\"getRateKeeperState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"internalType\":\"structRate[]\",\"name\":\"rates\",\"type\":\"tuple[]\"}],\"internalType\":\"structRateKeeperState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolCompressorMetaData.ABI instead.
var PoolCompressorABI = PoolCompressorMetaData.ABI

// PoolCompressor is an auto generated Go binding around an Ethereum contract.
type PoolCompressor struct {
	PoolCompressorCaller     // Read-only binding to the contract
	PoolCompressorTransactor // Write-only binding to the contract
	PoolCompressorFilterer   // Log filterer for contract events
}

// PoolCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolCompressorSession struct {
	Contract     *PoolCompressor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCompressorCallerSession struct {
	Contract *PoolCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PoolCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolCompressorTransactorSession struct {
	Contract     *PoolCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PoolCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolCompressorRaw struct {
	Contract *PoolCompressor // Generic contract binding to access the raw methods on
}

// PoolCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCompressorCallerRaw struct {
	Contract *PoolCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// PoolCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolCompressorTransactorRaw struct {
	Contract *PoolCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolCompressor creates a new instance of PoolCompressor, bound to a specific deployed contract.
func NewPoolCompressor(address common.Address, backend bind.ContractBackend) (*PoolCompressor, error) {
	contract, err := bindPoolCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolCompressor{PoolCompressorCaller: PoolCompressorCaller{contract: contract}, PoolCompressorTransactor: PoolCompressorTransactor{contract: contract}, PoolCompressorFilterer: PoolCompressorFilterer{contract: contract}}, nil
}

// NewPoolCompressorCaller creates a new read-only instance of PoolCompressor, bound to a specific deployed contract.
func NewPoolCompressorCaller(address common.Address, caller bind.ContractCaller) (*PoolCompressorCaller, error) {
	contract, err := bindPoolCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCompressorCaller{contract: contract}, nil
}

// NewPoolCompressorTransactor creates a new write-only instance of PoolCompressor, bound to a specific deployed contract.
func NewPoolCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolCompressorTransactor, error) {
	contract, err := bindPoolCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCompressorTransactor{contract: contract}, nil
}

// NewPoolCompressorFilterer creates a new log filterer instance of PoolCompressor, bound to a specific deployed contract.
func NewPoolCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolCompressorFilterer, error) {
	contract, err := bindPoolCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolCompressorFilterer{contract: contract}, nil
}

// bindPoolCompressor binds a generic wrapper to an already deployed contract.
func bindPoolCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolCompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolCompressor *PoolCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolCompressor.Contract.PoolCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolCompressor *PoolCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolCompressor.Contract.PoolCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolCompressor *PoolCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolCompressor.Contract.PoolCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolCompressor *PoolCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolCompressor *PoolCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolCompressor *PoolCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolCompressor.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PoolCompressor *PoolCompressorCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PoolCompressor *PoolCompressorSession) ContractType() ([32]byte, error) {
	return _PoolCompressor.Contract.ContractType(&_PoolCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PoolCompressor *PoolCompressorCallerSession) ContractType() ([32]byte, error) {
	return _PoolCompressor.Contract.ContractType(&_PoolCompressor.CallOpts)
}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorCaller) GetInterestRateModelState(opts *bind.CallOpts, interestRateModel common.Address) (BaseState, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "getInterestRateModelState", interestRateModel)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorSession) GetInterestRateModelState(interestRateModel common.Address) (BaseState, error) {
	return _PoolCompressor.Contract.GetInterestRateModelState(&_PoolCompressor.CallOpts, interestRateModel)
}

// GetInterestRateModelState is a free data retrieval call binding the contract method 0xc6da210b.
//
// Solidity: function getInterestRateModelState(address interestRateModel) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorCallerSession) GetInterestRateModelState(interestRateModel common.Address) (BaseState, error) {
	return _PoolCompressor.Contract.GetInterestRateModelState(&_PoolCompressor.CallOpts, interestRateModel)
}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorCaller) GetLossPolicyState(opts *bind.CallOpts, lossPolicy common.Address) (BaseState, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "getLossPolicyState", lossPolicy)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorSession) GetLossPolicyState(lossPolicy common.Address) (BaseState, error) {
	return _PoolCompressor.Contract.GetLossPolicyState(&_PoolCompressor.CallOpts, lossPolicy)
}

// GetLossPolicyState is a free data retrieval call binding the contract method 0x6776c3e5.
//
// Solidity: function getLossPolicyState(address lossPolicy) view returns(((address,uint256,bytes32,bytes)))
func (_PoolCompressor *PoolCompressorCallerSession) GetLossPolicyState(lossPolicy common.Address) (BaseState, error) {
	return _PoolCompressor.Contract.GetLossPolicyState(&_PoolCompressor.CallOpts, lossPolicy)
}

// GetPoolQuotaKeeperState is a free data retrieval call binding the contract method 0x1927805d.
//
// Solidity: function getPoolQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_PoolCompressor *PoolCompressorCaller) GetPoolQuotaKeeperState(opts *bind.CallOpts, quotaKeeper common.Address) (PoolQuotaKeeperState, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "getPoolQuotaKeeperState", quotaKeeper)

	if err != nil {
		return *new(PoolQuotaKeeperState), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolQuotaKeeperState)).(*PoolQuotaKeeperState)

	return out0, err

}

// GetPoolQuotaKeeperState is a free data retrieval call binding the contract method 0x1927805d.
//
// Solidity: function getPoolQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_PoolCompressor *PoolCompressorSession) GetPoolQuotaKeeperState(quotaKeeper common.Address) (PoolQuotaKeeperState, error) {
	return _PoolCompressor.Contract.GetPoolQuotaKeeperState(&_PoolCompressor.CallOpts, quotaKeeper)
}

// GetPoolQuotaKeeperState is a free data retrieval call binding the contract method 0x1927805d.
//
// Solidity: function getPoolQuotaKeeperState(address quotaKeeper) view returns(((address,uint256,bytes32,bytes),address,(address,uint16,uint192,uint16,uint96,uint96,bool)[],address[],uint40) result)
func (_PoolCompressor *PoolCompressorCallerSession) GetPoolQuotaKeeperState(quotaKeeper common.Address) (PoolQuotaKeeperState, error) {
	return _PoolCompressor.Contract.GetPoolQuotaKeeperState(&_PoolCompressor.CallOpts, quotaKeeper)
}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_PoolCompressor *PoolCompressorCaller) GetPoolState(opts *bind.CallOpts, pool common.Address) (PoolState, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "getPoolState", pool)

	if err != nil {
		return *new(PoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolState)).(*PoolState)

	return out0, err

}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_PoolCompressor *PoolCompressorSession) GetPoolState(pool common.Address) (PoolState, error) {
	return _PoolCompressor.Contract.GetPoolState(&_PoolCompressor.CallOpts, pool)
}

// GetPoolState is a free data retrieval call binding the contract method 0xec04205b.
//
// Solidity: function getPoolState(address pool) view returns(((address,uint256,bytes32,bytes),string,string,uint8,uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)[],uint256,uint256,uint256,uint40,uint40,bool) result)
func (_PoolCompressor *PoolCompressorCallerSession) GetPoolState(pool common.Address) (PoolState, error) {
	return _PoolCompressor.Contract.GetPoolState(&_PoolCompressor.CallOpts, pool)
}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_PoolCompressor *PoolCompressorCaller) GetRateKeeperState(opts *bind.CallOpts, rateKeeper common.Address) (RateKeeperState, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "getRateKeeperState", rateKeeper)

	if err != nil {
		return *new(RateKeeperState), err
	}

	out0 := *abi.ConvertType(out[0], new(RateKeeperState)).(*RateKeeperState)

	return out0, err

}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_PoolCompressor *PoolCompressorSession) GetRateKeeperState(rateKeeper common.Address) (RateKeeperState, error) {
	return _PoolCompressor.Contract.GetRateKeeperState(&_PoolCompressor.CallOpts, rateKeeper)
}

// GetRateKeeperState is a free data retrieval call binding the contract method 0x3c77fd5d.
//
// Solidity: function getRateKeeperState(address rateKeeper) view returns(((address,uint256,bytes32,bytes),(address,uint16)[]) result)
func (_PoolCompressor *PoolCompressorCallerSession) GetRateKeeperState(rateKeeper common.Address) (RateKeeperState, error) {
	return _PoolCompressor.Contract.GetRateKeeperState(&_PoolCompressor.CallOpts, rateKeeper)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolCompressor *PoolCompressorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolCompressor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolCompressor *PoolCompressorSession) Version() (*big.Int, error) {
	return _PoolCompressor.Contract.Version(&_PoolCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PoolCompressor *PoolCompressorCallerSession) Version() (*big.Int, error) {
	return _PoolCompressor.Contract.Version(&_PoolCompressor.CallOpts)
}
