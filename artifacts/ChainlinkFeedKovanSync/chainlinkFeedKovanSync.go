// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainlinkFeedKovanSync

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChainlinkFeedKovanSyncABI is the input ABI used to generate the binding from.
const ChainlinkFeedKovanSyncABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_syncer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_mainnetOracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainnetOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"phaseAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phaseId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncer\",\"outputs\":[{\"internalType\":\"contractSyncer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"updateParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ChainlinkFeedKovanSync is an auto generated Go binding around an Ethereum contract.
type ChainlinkFeedKovanSync struct {
	ChainlinkFeedKovanSyncCaller     // Read-only binding to the contract
	ChainlinkFeedKovanSyncTransactor // Write-only binding to the contract
	ChainlinkFeedKovanSyncFilterer   // Log filterer for contract events
}

// ChainlinkFeedKovanSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkFeedKovanSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedKovanSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkFeedKovanSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedKovanSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkFeedKovanSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkFeedKovanSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkFeedKovanSyncSession struct {
	Contract     *ChainlinkFeedKovanSync // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChainlinkFeedKovanSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkFeedKovanSyncCallerSession struct {
	Contract *ChainlinkFeedKovanSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ChainlinkFeedKovanSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkFeedKovanSyncTransactorSession struct {
	Contract     *ChainlinkFeedKovanSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ChainlinkFeedKovanSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkFeedKovanSyncRaw struct {
	Contract *ChainlinkFeedKovanSync // Generic contract binding to access the raw methods on
}

// ChainlinkFeedKovanSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkFeedKovanSyncCallerRaw struct {
	Contract *ChainlinkFeedKovanSyncCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkFeedKovanSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkFeedKovanSyncTransactorRaw struct {
	Contract *ChainlinkFeedKovanSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkFeedKovanSync creates a new instance of ChainlinkFeedKovanSync, bound to a specific deployed contract.
func NewChainlinkFeedKovanSync(address common.Address, backend bind.ContractBackend) (*ChainlinkFeedKovanSync, error) {
	contract, err := bindChainlinkFeedKovanSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedKovanSync{ChainlinkFeedKovanSyncCaller: ChainlinkFeedKovanSyncCaller{contract: contract}, ChainlinkFeedKovanSyncTransactor: ChainlinkFeedKovanSyncTransactor{contract: contract}, ChainlinkFeedKovanSyncFilterer: ChainlinkFeedKovanSyncFilterer{contract: contract}}, nil
}

// NewChainlinkFeedKovanSyncCaller creates a new read-only instance of ChainlinkFeedKovanSync, bound to a specific deployed contract.
func NewChainlinkFeedKovanSyncCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkFeedKovanSyncCaller, error) {
	contract, err := bindChainlinkFeedKovanSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedKovanSyncCaller{contract: contract}, nil
}

// NewChainlinkFeedKovanSyncTransactor creates a new write-only instance of ChainlinkFeedKovanSync, bound to a specific deployed contract.
func NewChainlinkFeedKovanSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkFeedKovanSyncTransactor, error) {
	contract, err := bindChainlinkFeedKovanSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedKovanSyncTransactor{contract: contract}, nil
}

// NewChainlinkFeedKovanSyncFilterer creates a new log filterer instance of ChainlinkFeedKovanSync, bound to a specific deployed contract.
func NewChainlinkFeedKovanSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkFeedKovanSyncFilterer, error) {
	contract, err := bindChainlinkFeedKovanSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedKovanSyncFilterer{contract: contract}, nil
}

// bindChainlinkFeedKovanSync binds a generic wrapper to an already deployed contract.
func bindChainlinkFeedKovanSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkFeedKovanSyncABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkFeedKovanSync.Contract.ChainlinkFeedKovanSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.ChainlinkFeedKovanSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.ChainlinkFeedKovanSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkFeedKovanSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) Decimals() (uint8, error) {
	return _ChainlinkFeedKovanSync.Contract.Decimals(&_ChainlinkFeedKovanSync.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) Decimals() (uint8, error) {
	return _ChainlinkFeedKovanSync.Contract.Decimals(&_ChainlinkFeedKovanSync.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) Description() (string, error) {
	return _ChainlinkFeedKovanSync.Contract.Description(&_ChainlinkFeedKovanSync.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() pure returns(string)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) Description() (string, error) {
	return _ChainlinkFeedKovanSync.Contract.Description(&_ChainlinkFeedKovanSync.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) GetRoundData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "getRoundData", arg0)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})

	outstruct.RoundId = out[0].(*big.Int)
	outstruct.Answer = out[1].(*big.Int)
	outstruct.StartedAt = out[2].(*big.Int)
	outstruct.UpdatedAt = out[3].(*big.Int)
	outstruct.AnsweredInRound = out[4].(*big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) GetRoundData(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkFeedKovanSync.Contract.GetRoundData(&_ChainlinkFeedKovanSync.CallOpts, arg0)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 ) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) GetRoundData(arg0 *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkFeedKovanSync.Contract.GetRoundData(&_ChainlinkFeedKovanSync.CallOpts, arg0)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})

	outstruct.RoundId = out[0].(*big.Int)
	outstruct.Answer = out[1].(*big.Int)
	outstruct.StartedAt = out[2].(*big.Int)
	outstruct.UpdatedAt = out[3].(*big.Int)
	outstruct.AnsweredInRound = out[4].(*big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkFeedKovanSync.Contract.LatestRoundData(&_ChainlinkFeedKovanSync.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _ChainlinkFeedKovanSync.Contract.LatestRoundData(&_ChainlinkFeedKovanSync.CallOpts)
}

// MainnetOracle is a free data retrieval call binding the contract method 0x785c9ad4.
//
// Solidity: function mainnetOracle() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) MainnetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "mainnetOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainnetOracle is a free data retrieval call binding the contract method 0x785c9ad4.
//
// Solidity: function mainnetOracle() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) MainnetOracle() (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.MainnetOracle(&_ChainlinkFeedKovanSync.CallOpts)
}

// MainnetOracle is a free data retrieval call binding the contract method 0x785c9ad4.
//
// Solidity: function mainnetOracle() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) MainnetOracle() (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.MainnetOracle(&_ChainlinkFeedKovanSync.CallOpts)
}

// PhaseAggregator is a free data retrieval call binding the contract method 0xd6bcd745.
//
// Solidity: function phaseAggregator(uint16 ) view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) PhaseAggregator(opts *bind.CallOpts, arg0 uint16) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "phaseAggregator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PhaseAggregator is a free data retrieval call binding the contract method 0xd6bcd745.
//
// Solidity: function phaseAggregator(uint16 ) view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) PhaseAggregator(arg0 uint16) (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.PhaseAggregator(&_ChainlinkFeedKovanSync.CallOpts, arg0)
}

// PhaseAggregator is a free data retrieval call binding the contract method 0xd6bcd745.
//
// Solidity: function phaseAggregator(uint16 ) view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) PhaseAggregator(arg0 uint16) (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.PhaseAggregator(&_ChainlinkFeedKovanSync.CallOpts, arg0)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) PhaseId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "phaseId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) PhaseId() (uint16, error) {
	return _ChainlinkFeedKovanSync.Contract.PhaseId(&_ChainlinkFeedKovanSync.CallOpts)
}

// PhaseId is a free data retrieval call binding the contract method 0x58303b10.
//
// Solidity: function phaseId() view returns(uint16)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) PhaseId() (uint16, error) {
	return _ChainlinkFeedKovanSync.Contract.PhaseId(&_ChainlinkFeedKovanSync.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) Syncer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "syncer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) Syncer() (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.Syncer(&_ChainlinkFeedKovanSync.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) Syncer() (common.Address, error) {
	return _ChainlinkFeedKovanSync.Contract.Syncer(&_ChainlinkFeedKovanSync.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkFeedKovanSync.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) Version() (*big.Int, error) {
	return _ChainlinkFeedKovanSync.Contract.Version(&_ChainlinkFeedKovanSync.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncCallerSession) Version() (*big.Int, error) {
	return _ChainlinkFeedKovanSync.Contract.Version(&_ChainlinkFeedKovanSync.CallOpts)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x1dd9e364.
//
// Solidity: function updateParams(uint80 roundId, int256 answer, uint256 timestamp) returns()
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncTransactor) UpdateParams(opts *bind.TransactOpts, roundId *big.Int, answer *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.contract.Transact(opts, "updateParams", roundId, answer, timestamp)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x1dd9e364.
//
// Solidity: function updateParams(uint80 roundId, int256 answer, uint256 timestamp) returns()
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncSession) UpdateParams(roundId *big.Int, answer *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.UpdateParams(&_ChainlinkFeedKovanSync.TransactOpts, roundId, answer, timestamp)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x1dd9e364.
//
// Solidity: function updateParams(uint80 roundId, int256 answer, uint256 timestamp) returns()
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncTransactorSession) UpdateParams(roundId *big.Int, answer *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _ChainlinkFeedKovanSync.Contract.UpdateParams(&_ChainlinkFeedKovanSync.TransactOpts, roundId, answer, timestamp)
}

// ChainlinkFeedKovanSyncAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the ChainlinkFeedKovanSync contract.
type ChainlinkFeedKovanSyncAnswerUpdatedIterator struct {
	Event *ChainlinkFeedKovanSyncAnswerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChainlinkFeedKovanSyncAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkFeedKovanSyncAnswerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChainlinkFeedKovanSyncAnswerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChainlinkFeedKovanSyncAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkFeedKovanSyncAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkFeedKovanSyncAnswerUpdated represents a AnswerUpdated event raised by the ChainlinkFeedKovanSync contract.
type ChainlinkFeedKovanSyncAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ChainlinkFeedKovanSyncAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkFeedKovanSync.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkFeedKovanSyncAnswerUpdatedIterator{contract: _ChainlinkFeedKovanSync.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ChainlinkFeedKovanSyncAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ChainlinkFeedKovanSync.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkFeedKovanSyncAnswerUpdated)
				if err := _ChainlinkFeedKovanSync.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_ChainlinkFeedKovanSync *ChainlinkFeedKovanSyncFilterer) ParseAnswerUpdated(log types.Log) (*ChainlinkFeedKovanSyncAnswerUpdated, error) {
	event := new(ChainlinkFeedKovanSyncAnswerUpdated)
	if err := _ChainlinkFeedKovanSync.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
