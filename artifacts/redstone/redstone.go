// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redstone

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

// RedstoneMetaData contains all meta data concerning the Redstone contract.
var RedstoneMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_dataFeedId\",\"type\":\"bytes32\"},{\"internalType\":\"address[10]\",\"name\":\"_signers\",\"type\":\"address[10]\"},{\"internalType\":\"uint8\",\"name\":\"signersThreshold\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CalldataMustHaveValidPayload\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CalldataOverOrUnderFlow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanNotPickMedianOfEmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataPackageTimestampIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataPackageTimestampMustNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataPackageTimestampsMustBeEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSignersException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EachSignerMustProvideTheSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyCalldataPointersArr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetDataServiceIdNotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSignersThresholdException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectUnsignedMetadataSize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedSignersCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignersCount\",\"type\":\"uint256\"}],\"name\":\"InsufficientNumberOfUniqueSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldataPointer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSignersException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedstonePayloadMustHaveAtLeastOneDataPackage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedstonePayloadTimestampIncorrect\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receivedSigner\",\"type\":\"address\"}],\"name\":\"SignerNotAuthorised\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"aggregateValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataFeedId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extractTimestampsAndAssertAllAreEqual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"extractedTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"getAuthorisedSignerIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDataServiceId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniqueSignersThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPayloadTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeedType\",\"outputs\":[{\"internalType\":\"enumPriceFeedType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress4\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress5\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress6\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress7\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress8\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerAddress9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skipPriceCheck\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedTimestampMilliseconds\",\"type\":\"uint256\"}],\"name\":\"validateTimestamp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RedstoneABI is the input ABI used to generate the binding from.
// Deprecated: Use RedstoneMetaData.ABI instead.
var RedstoneABI = RedstoneMetaData.ABI

// Redstone is an auto generated Go binding around an Ethereum contract.
type Redstone struct {
	RedstoneCaller     // Read-only binding to the contract
	RedstoneTransactor // Write-only binding to the contract
	RedstoneFilterer   // Log filterer for contract events
}

// RedstoneCaller is an auto generated read-only Go binding around an Ethereum contract.
type RedstoneCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedstoneTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RedstoneTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedstoneFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RedstoneFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RedstoneSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RedstoneSession struct {
	Contract     *Redstone         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RedstoneCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RedstoneCallerSession struct {
	Contract *RedstoneCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RedstoneTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RedstoneTransactorSession struct {
	Contract     *RedstoneTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RedstoneRaw is an auto generated low-level Go binding around an Ethereum contract.
type RedstoneRaw struct {
	Contract *Redstone // Generic contract binding to access the raw methods on
}

// RedstoneCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RedstoneCallerRaw struct {
	Contract *RedstoneCaller // Generic read-only contract binding to access the raw methods on
}

// RedstoneTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RedstoneTransactorRaw struct {
	Contract *RedstoneTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRedstone creates a new instance of Redstone, bound to a specific deployed contract.
func NewRedstone(address common.Address, backend bind.ContractBackend) (*Redstone, error) {
	contract, err := bindRedstone(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Redstone{RedstoneCaller: RedstoneCaller{contract: contract}, RedstoneTransactor: RedstoneTransactor{contract: contract}, RedstoneFilterer: RedstoneFilterer{contract: contract}}, nil
}

// NewRedstoneCaller creates a new read-only instance of Redstone, bound to a specific deployed contract.
func NewRedstoneCaller(address common.Address, caller bind.ContractCaller) (*RedstoneCaller, error) {
	contract, err := bindRedstone(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RedstoneCaller{contract: contract}, nil
}

// NewRedstoneTransactor creates a new write-only instance of Redstone, bound to a specific deployed contract.
func NewRedstoneTransactor(address common.Address, transactor bind.ContractTransactor) (*RedstoneTransactor, error) {
	contract, err := bindRedstone(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RedstoneTransactor{contract: contract}, nil
}

// NewRedstoneFilterer creates a new log filterer instance of Redstone, bound to a specific deployed contract.
func NewRedstoneFilterer(address common.Address, filterer bind.ContractFilterer) (*RedstoneFilterer, error) {
	contract, err := bindRedstone(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RedstoneFilterer{contract: contract}, nil
}

// bindRedstone binds a generic wrapper to an already deployed contract.
func bindRedstone(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RedstoneMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redstone *RedstoneRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redstone.Contract.RedstoneCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redstone *RedstoneRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redstone.Contract.RedstoneTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redstone *RedstoneRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redstone.Contract.RedstoneTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Redstone *RedstoneCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Redstone.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Redstone *RedstoneTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Redstone.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Redstone *RedstoneTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Redstone.Contract.contract.Transact(opts, method, params...)
}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_Redstone *RedstoneCaller) AggregateValues(opts *bind.CallOpts, values []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "aggregateValues", values)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_Redstone *RedstoneSession) AggregateValues(values []*big.Int) (*big.Int, error) {
	return _Redstone.Contract.AggregateValues(&_Redstone.CallOpts, values)
}

// AggregateValues is a free data retrieval call binding the contract method 0xb24ebfcc.
//
// Solidity: function aggregateValues(uint256[] values) view returns(uint256)
func (_Redstone *RedstoneCallerSession) AggregateValues(values []*big.Int) (*big.Int, error) {
	return _Redstone.Contract.AggregateValues(&_Redstone.CallOpts, values)
}

// DataFeedId is a free data retrieval call binding the contract method 0x370c826b.
//
// Solidity: function dataFeedId() view returns(bytes32)
func (_Redstone *RedstoneCaller) DataFeedId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "dataFeedId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataFeedId is a free data retrieval call binding the contract method 0x370c826b.
//
// Solidity: function dataFeedId() view returns(bytes32)
func (_Redstone *RedstoneSession) DataFeedId() ([32]byte, error) {
	return _Redstone.Contract.DataFeedId(&_Redstone.CallOpts)
}

// DataFeedId is a free data retrieval call binding the contract method 0x370c826b.
//
// Solidity: function dataFeedId() view returns(bytes32)
func (_Redstone *RedstoneCallerSession) DataFeedId() ([32]byte, error) {
	return _Redstone.Contract.DataFeedId(&_Redstone.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Redstone *RedstoneCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Redstone *RedstoneSession) Decimals() (uint8, error) {
	return _Redstone.Contract.Decimals(&_Redstone.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Redstone *RedstoneCallerSession) Decimals() (uint8, error) {
	return _Redstone.Contract.Decimals(&_Redstone.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Redstone *RedstoneCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Redstone *RedstoneSession) Description() (string, error) {
	return _Redstone.Contract.Description(&_Redstone.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Redstone *RedstoneCallerSession) Description() (string, error) {
	return _Redstone.Contract.Description(&_Redstone.CallOpts)
}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_Redstone *RedstoneCaller) ExtractTimestampsAndAssertAllAreEqual(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "extractTimestampsAndAssertAllAreEqual")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_Redstone *RedstoneSession) ExtractTimestampsAndAssertAllAreEqual() (*big.Int, error) {
	return _Redstone.Contract.ExtractTimestampsAndAssertAllAreEqual(&_Redstone.CallOpts)
}

// ExtractTimestampsAndAssertAllAreEqual is a free data retrieval call binding the contract method 0x55a547d5.
//
// Solidity: function extractTimestampsAndAssertAllAreEqual() pure returns(uint256 extractedTimestamp)
func (_Redstone *RedstoneCallerSession) ExtractTimestampsAndAssertAllAreEqual() (*big.Int, error) {
	return _Redstone.Contract.ExtractTimestampsAndAssertAllAreEqual(&_Redstone.CallOpts)
}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_Redstone *RedstoneCaller) GetAuthorisedSignerIndex(opts *bind.CallOpts, signerAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "getAuthorisedSignerIndex", signerAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_Redstone *RedstoneSession) GetAuthorisedSignerIndex(signerAddress common.Address) (uint8, error) {
	return _Redstone.Contract.GetAuthorisedSignerIndex(&_Redstone.CallOpts, signerAddress)
}

// GetAuthorisedSignerIndex is a free data retrieval call binding the contract method 0x3ce142f5.
//
// Solidity: function getAuthorisedSignerIndex(address signerAddress) view returns(uint8)
func (_Redstone *RedstoneCallerSession) GetAuthorisedSignerIndex(signerAddress common.Address) (uint8, error) {
	return _Redstone.Contract.GetAuthorisedSignerIndex(&_Redstone.CallOpts, signerAddress)
}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_Redstone *RedstoneCaller) GetDataServiceId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "getDataServiceId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_Redstone *RedstoneSession) GetDataServiceId() (string, error) {
	return _Redstone.Contract.GetDataServiceId(&_Redstone.CallOpts)
}

// GetDataServiceId is a free data retrieval call binding the contract method 0xc274583a.
//
// Solidity: function getDataServiceId() view returns(string)
func (_Redstone *RedstoneCallerSession) GetDataServiceId() (string, error) {
	return _Redstone.Contract.GetDataServiceId(&_Redstone.CallOpts)
}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_Redstone *RedstoneCaller) GetUniqueSignersThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "getUniqueSignersThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_Redstone *RedstoneSession) GetUniqueSignersThreshold() (uint8, error) {
	return _Redstone.Contract.GetUniqueSignersThreshold(&_Redstone.CallOpts)
}

// GetUniqueSignersThreshold is a free data retrieval call binding the contract method 0xf90c4924.
//
// Solidity: function getUniqueSignersThreshold() view returns(uint8)
func (_Redstone *RedstoneCallerSession) GetUniqueSignersThreshold() (uint8, error) {
	return _Redstone.Contract.GetUniqueSignersThreshold(&_Redstone.CallOpts)
}

// LastPayloadTimestamp is a free data retrieval call binding the contract method 0xbdbe4204.
//
// Solidity: function lastPayloadTimestamp() view returns(uint40)
func (_Redstone *RedstoneCaller) LastPayloadTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "lastPayloadTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPayloadTimestamp is a free data retrieval call binding the contract method 0xbdbe4204.
//
// Solidity: function lastPayloadTimestamp() view returns(uint40)
func (_Redstone *RedstoneSession) LastPayloadTimestamp() (*big.Int, error) {
	return _Redstone.Contract.LastPayloadTimestamp(&_Redstone.CallOpts)
}

// LastPayloadTimestamp is a free data retrieval call binding the contract method 0xbdbe4204.
//
// Solidity: function lastPayloadTimestamp() view returns(uint40)
func (_Redstone *RedstoneCallerSession) LastPayloadTimestamp() (*big.Int, error) {
	return _Redstone.Contract.LastPayloadTimestamp(&_Redstone.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint128)
func (_Redstone *RedstoneCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "lastPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint128)
func (_Redstone *RedstoneSession) LastPrice() (*big.Int, error) {
	return _Redstone.Contract.LastPrice(&_Redstone.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint128)
func (_Redstone *RedstoneCallerSession) LastPrice() (*big.Int, error) {
	return _Redstone.Contract.LastPrice(&_Redstone.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Redstone *RedstoneCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Redstone *RedstoneSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Redstone.Contract.LatestRoundData(&_Redstone.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256, uint256, uint256, uint80)
func (_Redstone *RedstoneCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Redstone.Contract.LatestRoundData(&_Redstone.CallOpts)
}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_Redstone *RedstoneCaller) PriceFeedType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "priceFeedType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_Redstone *RedstoneSession) PriceFeedType() (uint8, error) {
	return _Redstone.Contract.PriceFeedType(&_Redstone.CallOpts)
}

// PriceFeedType is a free data retrieval call binding the contract method 0x3fd0875f.
//
// Solidity: function priceFeedType() view returns(uint8)
func (_Redstone *RedstoneCallerSession) PriceFeedType() (uint8, error) {
	return _Redstone.Contract.PriceFeedType(&_Redstone.CallOpts)
}

// SignerAddress0 is a free data retrieval call binding the contract method 0x816f444a.
//
// Solidity: function signerAddress0() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress0 is a free data retrieval call binding the contract method 0x816f444a.
//
// Solidity: function signerAddress0() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress0() (common.Address, error) {
	return _Redstone.Contract.SignerAddress0(&_Redstone.CallOpts)
}

// SignerAddress0 is a free data retrieval call binding the contract method 0x816f444a.
//
// Solidity: function signerAddress0() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress0() (common.Address, error) {
	return _Redstone.Contract.SignerAddress0(&_Redstone.CallOpts)
}

// SignerAddress1 is a free data retrieval call binding the contract method 0x5decfe37.
//
// Solidity: function signerAddress1() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress1 is a free data retrieval call binding the contract method 0x5decfe37.
//
// Solidity: function signerAddress1() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress1() (common.Address, error) {
	return _Redstone.Contract.SignerAddress1(&_Redstone.CallOpts)
}

// SignerAddress1 is a free data retrieval call binding the contract method 0x5decfe37.
//
// Solidity: function signerAddress1() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress1() (common.Address, error) {
	return _Redstone.Contract.SignerAddress1(&_Redstone.CallOpts)
}

// SignerAddress2 is a free data retrieval call binding the contract method 0x8dba0538.
//
// Solidity: function signerAddress2() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress2 is a free data retrieval call binding the contract method 0x8dba0538.
//
// Solidity: function signerAddress2() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress2() (common.Address, error) {
	return _Redstone.Contract.SignerAddress2(&_Redstone.CallOpts)
}

// SignerAddress2 is a free data retrieval call binding the contract method 0x8dba0538.
//
// Solidity: function signerAddress2() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress2() (common.Address, error) {
	return _Redstone.Contract.SignerAddress2(&_Redstone.CallOpts)
}

// SignerAddress3 is a free data retrieval call binding the contract method 0x09c032ba.
//
// Solidity: function signerAddress3() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress3 is a free data retrieval call binding the contract method 0x09c032ba.
//
// Solidity: function signerAddress3() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress3() (common.Address, error) {
	return _Redstone.Contract.SignerAddress3(&_Redstone.CallOpts)
}

// SignerAddress3 is a free data retrieval call binding the contract method 0x09c032ba.
//
// Solidity: function signerAddress3() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress3() (common.Address, error) {
	return _Redstone.Contract.SignerAddress3(&_Redstone.CallOpts)
}

// SignerAddress4 is a free data retrieval call binding the contract method 0x7ed0185f.
//
// Solidity: function signerAddress4() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress4 is a free data retrieval call binding the contract method 0x7ed0185f.
//
// Solidity: function signerAddress4() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress4() (common.Address, error) {
	return _Redstone.Contract.SignerAddress4(&_Redstone.CallOpts)
}

// SignerAddress4 is a free data retrieval call binding the contract method 0x7ed0185f.
//
// Solidity: function signerAddress4() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress4() (common.Address, error) {
	return _Redstone.Contract.SignerAddress4(&_Redstone.CallOpts)
}

// SignerAddress5 is a free data retrieval call binding the contract method 0xdcdacf89.
//
// Solidity: function signerAddress5() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress5(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress5")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress5 is a free data retrieval call binding the contract method 0xdcdacf89.
//
// Solidity: function signerAddress5() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress5() (common.Address, error) {
	return _Redstone.Contract.SignerAddress5(&_Redstone.CallOpts)
}

// SignerAddress5 is a free data retrieval call binding the contract method 0xdcdacf89.
//
// Solidity: function signerAddress5() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress5() (common.Address, error) {
	return _Redstone.Contract.SignerAddress5(&_Redstone.CallOpts)
}

// SignerAddress6 is a free data retrieval call binding the contract method 0x9e0153d4.
//
// Solidity: function signerAddress6() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress6(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress6")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress6 is a free data retrieval call binding the contract method 0x9e0153d4.
//
// Solidity: function signerAddress6() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress6() (common.Address, error) {
	return _Redstone.Contract.SignerAddress6(&_Redstone.CallOpts)
}

// SignerAddress6 is a free data retrieval call binding the contract method 0x9e0153d4.
//
// Solidity: function signerAddress6() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress6() (common.Address, error) {
	return _Redstone.Contract.SignerAddress6(&_Redstone.CallOpts)
}

// SignerAddress7 is a free data retrieval call binding the contract method 0x1ac23b3d.
//
// Solidity: function signerAddress7() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress7(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress7")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress7 is a free data retrieval call binding the contract method 0x1ac23b3d.
//
// Solidity: function signerAddress7() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress7() (common.Address, error) {
	return _Redstone.Contract.SignerAddress7(&_Redstone.CallOpts)
}

// SignerAddress7 is a free data retrieval call binding the contract method 0x1ac23b3d.
//
// Solidity: function signerAddress7() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress7() (common.Address, error) {
	return _Redstone.Contract.SignerAddress7(&_Redstone.CallOpts)
}

// SignerAddress8 is a free data retrieval call binding the contract method 0x4ee1a1e6.
//
// Solidity: function signerAddress8() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress8(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress8")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress8 is a free data retrieval call binding the contract method 0x4ee1a1e6.
//
// Solidity: function signerAddress8() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress8() (common.Address, error) {
	return _Redstone.Contract.SignerAddress8(&_Redstone.CallOpts)
}

// SignerAddress8 is a free data retrieval call binding the contract method 0x4ee1a1e6.
//
// Solidity: function signerAddress8() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress8() (common.Address, error) {
	return _Redstone.Contract.SignerAddress8(&_Redstone.CallOpts)
}

// SignerAddress9 is a free data retrieval call binding the contract method 0xbc48da9e.
//
// Solidity: function signerAddress9() view returns(address)
func (_Redstone *RedstoneCaller) SignerAddress9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "signerAddress9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerAddress9 is a free data retrieval call binding the contract method 0xbc48da9e.
//
// Solidity: function signerAddress9() view returns(address)
func (_Redstone *RedstoneSession) SignerAddress9() (common.Address, error) {
	return _Redstone.Contract.SignerAddress9(&_Redstone.CallOpts)
}

// SignerAddress9 is a free data retrieval call binding the contract method 0xbc48da9e.
//
// Solidity: function signerAddress9() view returns(address)
func (_Redstone *RedstoneCallerSession) SignerAddress9() (common.Address, error) {
	return _Redstone.Contract.SignerAddress9(&_Redstone.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Redstone *RedstoneCaller) SkipPriceCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "skipPriceCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Redstone *RedstoneSession) SkipPriceCheck() (bool, error) {
	return _Redstone.Contract.SkipPriceCheck(&_Redstone.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_Redstone *RedstoneCallerSession) SkipPriceCheck() (bool, error) {
	return _Redstone.Contract.SkipPriceCheck(&_Redstone.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Redstone *RedstoneCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Redstone *RedstoneSession) Token() (common.Address, error) {
	return _Redstone.Contract.Token(&_Redstone.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Redstone *RedstoneCallerSession) Token() (common.Address, error) {
	return _Redstone.Contract.Token(&_Redstone.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Redstone *RedstoneCaller) Updatable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "updatable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Redstone *RedstoneSession) Updatable() (bool, error) {
	return _Redstone.Contract.Updatable(&_Redstone.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_Redstone *RedstoneCallerSession) Updatable() (bool, error) {
	return _Redstone.Contract.Updatable(&_Redstone.CallOpts)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_Redstone *RedstoneCaller) ValidateTimestamp(opts *bind.CallOpts, receivedTimestampMilliseconds *big.Int) error {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "validateTimestamp", receivedTimestampMilliseconds)

	if err != nil {
		return err
	}

	return err

}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_Redstone *RedstoneSession) ValidateTimestamp(receivedTimestampMilliseconds *big.Int) error {
	return _Redstone.Contract.ValidateTimestamp(&_Redstone.CallOpts, receivedTimestampMilliseconds)
}

// ValidateTimestamp is a free data retrieval call binding the contract method 0xf50b2efe.
//
// Solidity: function validateTimestamp(uint256 receivedTimestampMilliseconds) view returns()
func (_Redstone *RedstoneCallerSession) ValidateTimestamp(receivedTimestampMilliseconds *big.Int) error {
	return _Redstone.Contract.ValidateTimestamp(&_Redstone.CallOpts, receivedTimestampMilliseconds)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Redstone *RedstoneCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Redstone.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Redstone *RedstoneSession) Version() (*big.Int, error) {
	return _Redstone.Contract.Version(&_Redstone.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Redstone *RedstoneCallerSession) Version() (*big.Int, error) {
	return _Redstone.Contract.Version(&_Redstone.CallOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Redstone *RedstoneTransactor) UpdatePrice(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Redstone.contract.Transact(opts, "updatePrice", data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Redstone *RedstoneSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _Redstone.Contract.UpdatePrice(&_Redstone.TransactOpts, data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_Redstone *RedstoneTransactorSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _Redstone.Contract.UpdatePrice(&_Redstone.TransactOpts, data)
}

// RedstoneUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the Redstone contract.
type RedstoneUpdatePriceIterator struct {
	Event *RedstoneUpdatePrice // Event containing the contract specifics and raw log

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
func (it *RedstoneUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RedstoneUpdatePrice)
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
		it.Event = new(RedstoneUpdatePrice)
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
func (it *RedstoneUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RedstoneUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RedstoneUpdatePrice represents a UpdatePrice event raised by the Redstone contract.
type RedstoneUpdatePrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Redstone *RedstoneFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*RedstoneUpdatePriceIterator, error) {

	logs, sub, err := _Redstone.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &RedstoneUpdatePriceIterator{contract: _Redstone.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Redstone *RedstoneFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *RedstoneUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _Redstone.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RedstoneUpdatePrice)
				if err := _Redstone.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_Redstone *RedstoneFilterer) ParseUpdatePrice(log types.Log) (*RedstoneUpdatePrice, error) {
	event := new(RedstoneUpdatePrice)
	if err := _Redstone.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
