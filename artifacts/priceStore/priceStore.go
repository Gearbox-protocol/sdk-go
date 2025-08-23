// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceStore

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

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	Target   common.Address
	CallData []byte
}

// ConnectedPriceFeed is an auto generated low-level Go binding around an user-defined struct.
type ConnectedPriceFeed struct {
	Token      common.Address
	PriceFeeds []common.Address
}

// PriceFeedInfo is an auto generated low-level Go binding around an user-defined struct.
type PriceFeedInfo struct {
	Name            string
	StalenessPeriod uint32
	PriceFeedType   [32]byte
	Version         *big.Int
}

// PriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type PriceUpdate struct {
	PriceFeed common.Address
	Data      []byte
}

// PriceStoreMetaData contains all meta data concerning the PriceStore contract.
var PriceStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AddressIsNotContractException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotOwnerException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"ForbiddenConfigurationMethodException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectPriceFeedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsAlreadyAddedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsAlreadyAllowedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsNotAllowedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsNotKnownException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsNotOwnedByStore\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"PriceFeedIsNotUpdatableException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StalePriceException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"AddPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"AddUpdatablePriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"AllowPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"ForbidPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"RemovePriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"SetStalenessPeriod\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"addPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"allowPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bytecodeRepository\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"configurePriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"forbidPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"getAllowanceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKnownPriceFeeds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKnownTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceFeeds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"getStalenessPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPriceFeedsMap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"priceFeeds\",\"type\":\"address[]\"}],\"internalType\":\"structConnectedPriceFeed[]\",\"name\":\"connectedPriceFeeds\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpdatablePriceFeeds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"isAllowedPriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"isKnownPriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isKnownToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"priceFeedInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"priceFeedType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structPriceFeedInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"removePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessPeriod\",\"type\":\"uint32\"}],\"name\":\"setStalenessPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structPriceUpdate[]\",\"name\":\"updates\",\"type\":\"tuple[]\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PriceStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceStoreMetaData.ABI instead.
var PriceStoreABI = PriceStoreMetaData.ABI

// PriceStore is an auto generated Go binding around an Ethereum contract.
type PriceStore struct {
	PriceStoreCaller     // Read-only binding to the contract
	PriceStoreTransactor // Write-only binding to the contract
	PriceStoreFilterer   // Log filterer for contract events
}

// PriceStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceStoreSession struct {
	Contract     *PriceStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceStoreCallerSession struct {
	Contract *PriceStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PriceStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceStoreTransactorSession struct {
	Contract     *PriceStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PriceStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceStoreRaw struct {
	Contract *PriceStore // Generic contract binding to access the raw methods on
}

// PriceStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceStoreCallerRaw struct {
	Contract *PriceStoreCaller // Generic read-only contract binding to access the raw methods on
}

// PriceStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceStoreTransactorRaw struct {
	Contract *PriceStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceStore creates a new instance of PriceStore, bound to a specific deployed contract.
func NewPriceStore(address common.Address, backend bind.ContractBackend) (*PriceStore, error) {
	contract, err := bindPriceStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceStore{PriceStoreCaller: PriceStoreCaller{contract: contract}, PriceStoreTransactor: PriceStoreTransactor{contract: contract}, PriceStoreFilterer: PriceStoreFilterer{contract: contract}}, nil
}

// NewPriceStoreCaller creates a new read-only instance of PriceStore, bound to a specific deployed contract.
func NewPriceStoreCaller(address common.Address, caller bind.ContractCaller) (*PriceStoreCaller, error) {
	contract, err := bindPriceStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceStoreCaller{contract: contract}, nil
}

// NewPriceStoreTransactor creates a new write-only instance of PriceStore, bound to a specific deployed contract.
func NewPriceStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceStoreTransactor, error) {
	contract, err := bindPriceStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceStoreTransactor{contract: contract}, nil
}

// NewPriceStoreFilterer creates a new log filterer instance of PriceStore, bound to a specific deployed contract.
func NewPriceStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceStoreFilterer, error) {
	contract, err := bindPriceStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceStoreFilterer{contract: contract}, nil
}

// bindPriceStore binds a generic wrapper to an already deployed contract.
func bindPriceStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceStore *PriceStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceStore.Contract.PriceStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceStore *PriceStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceStore.Contract.PriceStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceStore *PriceStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceStore.Contract.PriceStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceStore *PriceStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceStore *PriceStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceStore *PriceStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceStore.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PriceStore *PriceStoreCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PriceStore *PriceStoreSession) AddressProvider() (common.Address, error) {
	return _PriceStore.Contract.AddressProvider(&_PriceStore.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_PriceStore *PriceStoreCallerSession) AddressProvider() (common.Address, error) {
	return _PriceStore.Contract.AddressProvider(&_PriceStore.CallOpts)
}

// BytecodeRepository is a free data retrieval call binding the contract method 0x60e93cfd.
//
// Solidity: function bytecodeRepository() view returns(address)
func (_PriceStore *PriceStoreCaller) BytecodeRepository(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "bytecodeRepository")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BytecodeRepository is a free data retrieval call binding the contract method 0x60e93cfd.
//
// Solidity: function bytecodeRepository() view returns(address)
func (_PriceStore *PriceStoreSession) BytecodeRepository() (common.Address, error) {
	return _PriceStore.Contract.BytecodeRepository(&_PriceStore.CallOpts)
}

// BytecodeRepository is a free data retrieval call binding the contract method 0x60e93cfd.
//
// Solidity: function bytecodeRepository() view returns(address)
func (_PriceStore *PriceStoreCallerSession) BytecodeRepository() (common.Address, error) {
	return _PriceStore.Contract.BytecodeRepository(&_PriceStore.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PriceStore *PriceStoreCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PriceStore *PriceStoreSession) ContractType() ([32]byte, error) {
	return _PriceStore.Contract.ContractType(&_PriceStore.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_PriceStore *PriceStoreCallerSession) ContractType() ([32]byte, error) {
	return _PriceStore.Contract.ContractType(&_PriceStore.CallOpts)
}

// GetAllowanceTimestamp is a free data retrieval call binding the contract method 0xba0f681e.
//
// Solidity: function getAllowanceTimestamp(address token, address priceFeed) view returns(uint256)
func (_PriceStore *PriceStoreCaller) GetAllowanceTimestamp(opts *bind.CallOpts, token common.Address, priceFeed common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getAllowanceTimestamp", token, priceFeed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllowanceTimestamp is a free data retrieval call binding the contract method 0xba0f681e.
//
// Solidity: function getAllowanceTimestamp(address token, address priceFeed) view returns(uint256)
func (_PriceStore *PriceStoreSession) GetAllowanceTimestamp(token common.Address, priceFeed common.Address) (*big.Int, error) {
	return _PriceStore.Contract.GetAllowanceTimestamp(&_PriceStore.CallOpts, token, priceFeed)
}

// GetAllowanceTimestamp is a free data retrieval call binding the contract method 0xba0f681e.
//
// Solidity: function getAllowanceTimestamp(address token, address priceFeed) view returns(uint256)
func (_PriceStore *PriceStoreCallerSession) GetAllowanceTimestamp(token common.Address, priceFeed common.Address) (*big.Int, error) {
	return _PriceStore.Contract.GetAllowanceTimestamp(&_PriceStore.CallOpts, token, priceFeed)
}

// GetKnownPriceFeeds is a free data retrieval call binding the contract method 0x614eb733.
//
// Solidity: function getKnownPriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreCaller) GetKnownPriceFeeds(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getKnownPriceFeeds")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetKnownPriceFeeds is a free data retrieval call binding the contract method 0x614eb733.
//
// Solidity: function getKnownPriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreSession) GetKnownPriceFeeds() ([]common.Address, error) {
	return _PriceStore.Contract.GetKnownPriceFeeds(&_PriceStore.CallOpts)
}

// GetKnownPriceFeeds is a free data retrieval call binding the contract method 0x614eb733.
//
// Solidity: function getKnownPriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreCallerSession) GetKnownPriceFeeds() ([]common.Address, error) {
	return _PriceStore.Contract.GetKnownPriceFeeds(&_PriceStore.CallOpts)
}

// GetKnownTokens is a free data retrieval call binding the contract method 0x4026ebd6.
//
// Solidity: function getKnownTokens() view returns(address[])
func (_PriceStore *PriceStoreCaller) GetKnownTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getKnownTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetKnownTokens is a free data retrieval call binding the contract method 0x4026ebd6.
//
// Solidity: function getKnownTokens() view returns(address[])
func (_PriceStore *PriceStoreSession) GetKnownTokens() ([]common.Address, error) {
	return _PriceStore.Contract.GetKnownTokens(&_PriceStore.CallOpts)
}

// GetKnownTokens is a free data retrieval call binding the contract method 0x4026ebd6.
//
// Solidity: function getKnownTokens() view returns(address[])
func (_PriceStore *PriceStoreCallerSession) GetKnownTokens() ([]common.Address, error) {
	return _PriceStore.Contract.GetKnownTokens(&_PriceStore.CallOpts)
}

// GetPriceFeeds is a free data retrieval call binding the contract method 0x7cc6ce52.
//
// Solidity: function getPriceFeeds(address token) view returns(address[])
func (_PriceStore *PriceStoreCaller) GetPriceFeeds(opts *bind.CallOpts, token common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getPriceFeeds", token)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPriceFeeds is a free data retrieval call binding the contract method 0x7cc6ce52.
//
// Solidity: function getPriceFeeds(address token) view returns(address[])
func (_PriceStore *PriceStoreSession) GetPriceFeeds(token common.Address) ([]common.Address, error) {
	return _PriceStore.Contract.GetPriceFeeds(&_PriceStore.CallOpts, token)
}

// GetPriceFeeds is a free data retrieval call binding the contract method 0x7cc6ce52.
//
// Solidity: function getPriceFeeds(address token) view returns(address[])
func (_PriceStore *PriceStoreCallerSession) GetPriceFeeds(token common.Address) ([]common.Address, error) {
	return _PriceStore.Contract.GetPriceFeeds(&_PriceStore.CallOpts, token)
}

// GetStalenessPeriod is a free data retrieval call binding the contract method 0xce37a61e.
//
// Solidity: function getStalenessPeriod(address priceFeed) view returns(uint32)
func (_PriceStore *PriceStoreCaller) GetStalenessPeriod(opts *bind.CallOpts, priceFeed common.Address) (uint32, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getStalenessPeriod", priceFeed)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStalenessPeriod is a free data retrieval call binding the contract method 0xce37a61e.
//
// Solidity: function getStalenessPeriod(address priceFeed) view returns(uint32)
func (_PriceStore *PriceStoreSession) GetStalenessPeriod(priceFeed common.Address) (uint32, error) {
	return _PriceStore.Contract.GetStalenessPeriod(&_PriceStore.CallOpts, priceFeed)
}

// GetStalenessPeriod is a free data retrieval call binding the contract method 0xce37a61e.
//
// Solidity: function getStalenessPeriod(address priceFeed) view returns(uint32)
func (_PriceStore *PriceStoreCallerSession) GetStalenessPeriod(priceFeed common.Address) (uint32, error) {
	return _PriceStore.Contract.GetStalenessPeriod(&_PriceStore.CallOpts, priceFeed)
}

// GetTokenPriceFeedsMap is a free data retrieval call binding the contract method 0x161455ad.
//
// Solidity: function getTokenPriceFeedsMap() view returns((address,address[])[] connectedPriceFeeds)
func (_PriceStore *PriceStoreCaller) GetTokenPriceFeedsMap(opts *bind.CallOpts) ([]ConnectedPriceFeed, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getTokenPriceFeedsMap")

	if err != nil {
		return *new([]ConnectedPriceFeed), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConnectedPriceFeed)).(*[]ConnectedPriceFeed)

	return out0, err

}

// GetTokenPriceFeedsMap is a free data retrieval call binding the contract method 0x161455ad.
//
// Solidity: function getTokenPriceFeedsMap() view returns((address,address[])[] connectedPriceFeeds)
func (_PriceStore *PriceStoreSession) GetTokenPriceFeedsMap() ([]ConnectedPriceFeed, error) {
	return _PriceStore.Contract.GetTokenPriceFeedsMap(&_PriceStore.CallOpts)
}

// GetTokenPriceFeedsMap is a free data retrieval call binding the contract method 0x161455ad.
//
// Solidity: function getTokenPriceFeedsMap() view returns((address,address[])[] connectedPriceFeeds)
func (_PriceStore *PriceStoreCallerSession) GetTokenPriceFeedsMap() ([]ConnectedPriceFeed, error) {
	return _PriceStore.Contract.GetTokenPriceFeedsMap(&_PriceStore.CallOpts)
}

// GetUpdatablePriceFeeds is a free data retrieval call binding the contract method 0x26ec6c30.
//
// Solidity: function getUpdatablePriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreCaller) GetUpdatablePriceFeeds(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "getUpdatablePriceFeeds")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUpdatablePriceFeeds is a free data retrieval call binding the contract method 0x26ec6c30.
//
// Solidity: function getUpdatablePriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreSession) GetUpdatablePriceFeeds() ([]common.Address, error) {
	return _PriceStore.Contract.GetUpdatablePriceFeeds(&_PriceStore.CallOpts)
}

// GetUpdatablePriceFeeds is a free data retrieval call binding the contract method 0x26ec6c30.
//
// Solidity: function getUpdatablePriceFeeds() view returns(address[])
func (_PriceStore *PriceStoreCallerSession) GetUpdatablePriceFeeds() ([]common.Address, error) {
	return _PriceStore.Contract.GetUpdatablePriceFeeds(&_PriceStore.CallOpts)
}

// IsAllowedPriceFeed is a free data retrieval call binding the contract method 0x30d1f9fe.
//
// Solidity: function isAllowedPriceFeed(address token, address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreCaller) IsAllowedPriceFeed(opts *bind.CallOpts, token common.Address, priceFeed common.Address) (bool, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "isAllowedPriceFeed", token, priceFeed)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedPriceFeed is a free data retrieval call binding the contract method 0x30d1f9fe.
//
// Solidity: function isAllowedPriceFeed(address token, address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreSession) IsAllowedPriceFeed(token common.Address, priceFeed common.Address) (bool, error) {
	return _PriceStore.Contract.IsAllowedPriceFeed(&_PriceStore.CallOpts, token, priceFeed)
}

// IsAllowedPriceFeed is a free data retrieval call binding the contract method 0x30d1f9fe.
//
// Solidity: function isAllowedPriceFeed(address token, address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreCallerSession) IsAllowedPriceFeed(token common.Address, priceFeed common.Address) (bool, error) {
	return _PriceStore.Contract.IsAllowedPriceFeed(&_PriceStore.CallOpts, token, priceFeed)
}

// IsKnownPriceFeed is a free data retrieval call binding the contract method 0xc7a22fe0.
//
// Solidity: function isKnownPriceFeed(address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreCaller) IsKnownPriceFeed(opts *bind.CallOpts, priceFeed common.Address) (bool, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "isKnownPriceFeed", priceFeed)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownPriceFeed is a free data retrieval call binding the contract method 0xc7a22fe0.
//
// Solidity: function isKnownPriceFeed(address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreSession) IsKnownPriceFeed(priceFeed common.Address) (bool, error) {
	return _PriceStore.Contract.IsKnownPriceFeed(&_PriceStore.CallOpts, priceFeed)
}

// IsKnownPriceFeed is a free data retrieval call binding the contract method 0xc7a22fe0.
//
// Solidity: function isKnownPriceFeed(address priceFeed) view returns(bool)
func (_PriceStore *PriceStoreCallerSession) IsKnownPriceFeed(priceFeed common.Address) (bool, error) {
	return _PriceStore.Contract.IsKnownPriceFeed(&_PriceStore.CallOpts, priceFeed)
}

// IsKnownToken is a free data retrieval call binding the contract method 0xe647e5cc.
//
// Solidity: function isKnownToken(address token) view returns(bool)
func (_PriceStore *PriceStoreCaller) IsKnownToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "isKnownToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKnownToken is a free data retrieval call binding the contract method 0xe647e5cc.
//
// Solidity: function isKnownToken(address token) view returns(bool)
func (_PriceStore *PriceStoreSession) IsKnownToken(token common.Address) (bool, error) {
	return _PriceStore.Contract.IsKnownToken(&_PriceStore.CallOpts, token)
}

// IsKnownToken is a free data retrieval call binding the contract method 0xe647e5cc.
//
// Solidity: function isKnownToken(address token) view returns(bool)
func (_PriceStore *PriceStoreCallerSession) IsKnownToken(token common.Address) (bool, error) {
	return _PriceStore.Contract.IsKnownToken(&_PriceStore.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceStore *PriceStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceStore *PriceStoreSession) Owner() (common.Address, error) {
	return _PriceStore.Contract.Owner(&_PriceStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceStore *PriceStoreCallerSession) Owner() (common.Address, error) {
	return _PriceStore.Contract.Owner(&_PriceStore.CallOpts)
}

// PriceFeedInfo is a free data retrieval call binding the contract method 0xa9ff6c0f.
//
// Solidity: function priceFeedInfo(address priceFeed) view returns((string,uint32,bytes32,uint256))
func (_PriceStore *PriceStoreCaller) PriceFeedInfo(opts *bind.CallOpts, priceFeed common.Address) (PriceFeedInfo, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "priceFeedInfo", priceFeed)

	if err != nil {
		return *new(PriceFeedInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceFeedInfo)).(*PriceFeedInfo)

	return out0, err

}

// PriceFeedInfo is a free data retrieval call binding the contract method 0xa9ff6c0f.
//
// Solidity: function priceFeedInfo(address priceFeed) view returns((string,uint32,bytes32,uint256))
func (_PriceStore *PriceStoreSession) PriceFeedInfo(priceFeed common.Address) (PriceFeedInfo, error) {
	return _PriceStore.Contract.PriceFeedInfo(&_PriceStore.CallOpts, priceFeed)
}

// PriceFeedInfo is a free data retrieval call binding the contract method 0xa9ff6c0f.
//
// Solidity: function priceFeedInfo(address priceFeed) view returns((string,uint32,bytes32,uint256))
func (_PriceStore *PriceStoreCallerSession) PriceFeedInfo(priceFeed common.Address) (PriceFeedInfo, error) {
	return _PriceStore.Contract.PriceFeedInfo(&_PriceStore.CallOpts, priceFeed)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceStore *PriceStoreCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceStore *PriceStoreSession) Version() (*big.Int, error) {
	return _PriceStore.Contract.Version(&_PriceStore.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_PriceStore *PriceStoreCallerSession) Version() (*big.Int, error) {
	return _PriceStore.Contract.Version(&_PriceStore.CallOpts)
}

// ZeroPriceFeed is a free data retrieval call binding the contract method 0x89e402f1.
//
// Solidity: function zeroPriceFeed() view returns(address)
func (_PriceStore *PriceStoreCaller) ZeroPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceStore.contract.Call(opts, &out, "zeroPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZeroPriceFeed is a free data retrieval call binding the contract method 0x89e402f1.
//
// Solidity: function zeroPriceFeed() view returns(address)
func (_PriceStore *PriceStoreSession) ZeroPriceFeed() (common.Address, error) {
	return _PriceStore.Contract.ZeroPriceFeed(&_PriceStore.CallOpts)
}

// ZeroPriceFeed is a free data retrieval call binding the contract method 0x89e402f1.
//
// Solidity: function zeroPriceFeed() view returns(address)
func (_PriceStore *PriceStoreCallerSession) ZeroPriceFeed() (common.Address, error) {
	return _PriceStore.Contract.ZeroPriceFeed(&_PriceStore.CallOpts)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0x16f5a392.
//
// Solidity: function addPriceFeed(address priceFeed, uint32 stalenessPeriod, string name) returns()
func (_PriceStore *PriceStoreTransactor) AddPriceFeed(opts *bind.TransactOpts, priceFeed common.Address, stalenessPeriod uint32, name string) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "addPriceFeed", priceFeed, stalenessPeriod, name)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0x16f5a392.
//
// Solidity: function addPriceFeed(address priceFeed, uint32 stalenessPeriod, string name) returns()
func (_PriceStore *PriceStoreSession) AddPriceFeed(priceFeed common.Address, stalenessPeriod uint32, name string) (*types.Transaction, error) {
	return _PriceStore.Contract.AddPriceFeed(&_PriceStore.TransactOpts, priceFeed, stalenessPeriod, name)
}

// AddPriceFeed is a paid mutator transaction binding the contract method 0x16f5a392.
//
// Solidity: function addPriceFeed(address priceFeed, uint32 stalenessPeriod, string name) returns()
func (_PriceStore *PriceStoreTransactorSession) AddPriceFeed(priceFeed common.Address, stalenessPeriod uint32, name string) (*types.Transaction, error) {
	return _PriceStore.Contract.AddPriceFeed(&_PriceStore.TransactOpts, priceFeed, stalenessPeriod, name)
}

// AllowPriceFeed is a paid mutator transaction binding the contract method 0x45af8396.
//
// Solidity: function allowPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreTransactor) AllowPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "allowPriceFeed", token, priceFeed)
}

// AllowPriceFeed is a paid mutator transaction binding the contract method 0x45af8396.
//
// Solidity: function allowPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreSession) AllowPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.AllowPriceFeed(&_PriceStore.TransactOpts, token, priceFeed)
}

// AllowPriceFeed is a paid mutator transaction binding the contract method 0x45af8396.
//
// Solidity: function allowPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreTransactorSession) AllowPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.AllowPriceFeed(&_PriceStore.TransactOpts, token, priceFeed)
}

// ConfigurePriceFeeds is a paid mutator transaction binding the contract method 0x71346b42.
//
// Solidity: function configurePriceFeeds((address,bytes)[] calls) returns()
func (_PriceStore *PriceStoreTransactor) ConfigurePriceFeeds(opts *bind.TransactOpts, calls []Call) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "configurePriceFeeds", calls)
}

// ConfigurePriceFeeds is a paid mutator transaction binding the contract method 0x71346b42.
//
// Solidity: function configurePriceFeeds((address,bytes)[] calls) returns()
func (_PriceStore *PriceStoreSession) ConfigurePriceFeeds(calls []Call) (*types.Transaction, error) {
	return _PriceStore.Contract.ConfigurePriceFeeds(&_PriceStore.TransactOpts, calls)
}

// ConfigurePriceFeeds is a paid mutator transaction binding the contract method 0x71346b42.
//
// Solidity: function configurePriceFeeds((address,bytes)[] calls) returns()
func (_PriceStore *PriceStoreTransactorSession) ConfigurePriceFeeds(calls []Call) (*types.Transaction, error) {
	return _PriceStore.Contract.ConfigurePriceFeeds(&_PriceStore.TransactOpts, calls)
}

// ForbidPriceFeed is a paid mutator transaction binding the contract method 0x6f898fb4.
//
// Solidity: function forbidPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreTransactor) ForbidPriceFeed(opts *bind.TransactOpts, token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "forbidPriceFeed", token, priceFeed)
}

// ForbidPriceFeed is a paid mutator transaction binding the contract method 0x6f898fb4.
//
// Solidity: function forbidPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreSession) ForbidPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.ForbidPriceFeed(&_PriceStore.TransactOpts, token, priceFeed)
}

// ForbidPriceFeed is a paid mutator transaction binding the contract method 0x6f898fb4.
//
// Solidity: function forbidPriceFeed(address token, address priceFeed) returns()
func (_PriceStore *PriceStoreTransactorSession) ForbidPriceFeed(token common.Address, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.ForbidPriceFeed(&_PriceStore.TransactOpts, token, priceFeed)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xfceb0024.
//
// Solidity: function removePriceFeed(address priceFeed) returns()
func (_PriceStore *PriceStoreTransactor) RemovePriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "removePriceFeed", priceFeed)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xfceb0024.
//
// Solidity: function removePriceFeed(address priceFeed) returns()
func (_PriceStore *PriceStoreSession) RemovePriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.RemovePriceFeed(&_PriceStore.TransactOpts, priceFeed)
}

// RemovePriceFeed is a paid mutator transaction binding the contract method 0xfceb0024.
//
// Solidity: function removePriceFeed(address priceFeed) returns()
func (_PriceStore *PriceStoreTransactorSession) RemovePriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _PriceStore.Contract.RemovePriceFeed(&_PriceStore.TransactOpts, priceFeed)
}

// SetStalenessPeriod is a paid mutator transaction binding the contract method 0x149d9191.
//
// Solidity: function setStalenessPeriod(address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceStore *PriceStoreTransactor) SetStalenessPeriod(opts *bind.TransactOpts, priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "setStalenessPeriod", priceFeed, stalenessPeriod)
}

// SetStalenessPeriod is a paid mutator transaction binding the contract method 0x149d9191.
//
// Solidity: function setStalenessPeriod(address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceStore *PriceStoreSession) SetStalenessPeriod(priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceStore.Contract.SetStalenessPeriod(&_PriceStore.TransactOpts, priceFeed, stalenessPeriod)
}

// SetStalenessPeriod is a paid mutator transaction binding the contract method 0x149d9191.
//
// Solidity: function setStalenessPeriod(address priceFeed, uint32 stalenessPeriod) returns()
func (_PriceStore *PriceStoreTransactorSession) SetStalenessPeriod(priceFeed common.Address, stalenessPeriod uint32) (*types.Transaction, error) {
	return _PriceStore.Contract.SetStalenessPeriod(&_PriceStore.TransactOpts, priceFeed, stalenessPeriod)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7199e2c9.
//
// Solidity: function updatePrices((address,bytes)[] updates) returns()
func (_PriceStore *PriceStoreTransactor) UpdatePrices(opts *bind.TransactOpts, updates []PriceUpdate) (*types.Transaction, error) {
	return _PriceStore.contract.Transact(opts, "updatePrices", updates)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7199e2c9.
//
// Solidity: function updatePrices((address,bytes)[] updates) returns()
func (_PriceStore *PriceStoreSession) UpdatePrices(updates []PriceUpdate) (*types.Transaction, error) {
	return _PriceStore.Contract.UpdatePrices(&_PriceStore.TransactOpts, updates)
}

// UpdatePrices is a paid mutator transaction binding the contract method 0x7199e2c9.
//
// Solidity: function updatePrices((address,bytes)[] updates) returns()
func (_PriceStore *PriceStoreTransactorSession) UpdatePrices(updates []PriceUpdate) (*types.Transaction, error) {
	return _PriceStore.Contract.UpdatePrices(&_PriceStore.TransactOpts, updates)
}

// PriceStoreAddPriceFeedIterator is returned from FilterAddPriceFeed and is used to iterate over the raw logs and unpacked data for AddPriceFeed events raised by the PriceStore contract.
type PriceStoreAddPriceFeedIterator struct {
	Event *PriceStoreAddPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceStoreAddPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreAddPriceFeed)
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
		it.Event = new(PriceStoreAddPriceFeed)
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
func (it *PriceStoreAddPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreAddPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreAddPriceFeed represents a AddPriceFeed event raised by the PriceStore contract.
type PriceStoreAddPriceFeed struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	Name            string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddPriceFeed is a free log retrieval operation binding the contract event 0x18c05f0511be33b4b3636da69473fd2c8b44012e11eabd7729a54611b6085932.
//
// Solidity: event AddPriceFeed(address indexed priceFeed, uint32 stalenessPeriod, string name)
func (_PriceStore *PriceStoreFilterer) FilterAddPriceFeed(opts *bind.FilterOpts, priceFeed []common.Address) (*PriceStoreAddPriceFeedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "AddPriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreAddPriceFeedIterator{contract: _PriceStore.contract, event: "AddPriceFeed", logs: logs, sub: sub}, nil
}

// WatchAddPriceFeed is a free log subscription operation binding the contract event 0x18c05f0511be33b4b3636da69473fd2c8b44012e11eabd7729a54611b6085932.
//
// Solidity: event AddPriceFeed(address indexed priceFeed, uint32 stalenessPeriod, string name)
func (_PriceStore *PriceStoreFilterer) WatchAddPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceStoreAddPriceFeed, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "AddPriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreAddPriceFeed)
				if err := _PriceStore.contract.UnpackLog(event, "AddPriceFeed", log); err != nil {
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

// ParseAddPriceFeed is a log parse operation binding the contract event 0x18c05f0511be33b4b3636da69473fd2c8b44012e11eabd7729a54611b6085932.
//
// Solidity: event AddPriceFeed(address indexed priceFeed, uint32 stalenessPeriod, string name)
func (_PriceStore *PriceStoreFilterer) ParseAddPriceFeed(log types.Log) (*PriceStoreAddPriceFeed, error) {
	event := new(PriceStoreAddPriceFeed)
	if err := _PriceStore.contract.UnpackLog(event, "AddPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceStoreAddUpdatablePriceFeedIterator is returned from FilterAddUpdatablePriceFeed and is used to iterate over the raw logs and unpacked data for AddUpdatablePriceFeed events raised by the PriceStore contract.
type PriceStoreAddUpdatablePriceFeedIterator struct {
	Event *PriceStoreAddUpdatablePriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceStoreAddUpdatablePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreAddUpdatablePriceFeed)
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
		it.Event = new(PriceStoreAddUpdatablePriceFeed)
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
func (it *PriceStoreAddUpdatablePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreAddUpdatablePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreAddUpdatablePriceFeed represents a AddUpdatablePriceFeed event raised by the PriceStore contract.
type PriceStoreAddUpdatablePriceFeed struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddUpdatablePriceFeed is a free log retrieval operation binding the contract event 0x126881a22591f9c505916524e1e627cee9847bd538cc439d374e156d965fd5ec.
//
// Solidity: event AddUpdatablePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) FilterAddUpdatablePriceFeed(opts *bind.FilterOpts, priceFeed []common.Address) (*PriceStoreAddUpdatablePriceFeedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "AddUpdatablePriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreAddUpdatablePriceFeedIterator{contract: _PriceStore.contract, event: "AddUpdatablePriceFeed", logs: logs, sub: sub}, nil
}

// WatchAddUpdatablePriceFeed is a free log subscription operation binding the contract event 0x126881a22591f9c505916524e1e627cee9847bd538cc439d374e156d965fd5ec.
//
// Solidity: event AddUpdatablePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) WatchAddUpdatablePriceFeed(opts *bind.WatchOpts, sink chan<- *PriceStoreAddUpdatablePriceFeed, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "AddUpdatablePriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreAddUpdatablePriceFeed)
				if err := _PriceStore.contract.UnpackLog(event, "AddUpdatablePriceFeed", log); err != nil {
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

// ParseAddUpdatablePriceFeed is a log parse operation binding the contract event 0x126881a22591f9c505916524e1e627cee9847bd538cc439d374e156d965fd5ec.
//
// Solidity: event AddUpdatablePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) ParseAddUpdatablePriceFeed(log types.Log) (*PriceStoreAddUpdatablePriceFeed, error) {
	event := new(PriceStoreAddUpdatablePriceFeed)
	if err := _PriceStore.contract.UnpackLog(event, "AddUpdatablePriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceStoreAllowPriceFeedIterator is returned from FilterAllowPriceFeed and is used to iterate over the raw logs and unpacked data for AllowPriceFeed events raised by the PriceStore contract.
type PriceStoreAllowPriceFeedIterator struct {
	Event *PriceStoreAllowPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceStoreAllowPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreAllowPriceFeed)
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
		it.Event = new(PriceStoreAllowPriceFeed)
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
func (it *PriceStoreAllowPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreAllowPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreAllowPriceFeed represents a AllowPriceFeed event raised by the PriceStore contract.
type PriceStoreAllowPriceFeed struct {
	Token     common.Address
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowPriceFeed is a free log retrieval operation binding the contract event 0x73720c9031fae32be0c4f58f3b3934399946cfa29d2b13019ead89d904aa950e.
//
// Solidity: event AllowPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) FilterAllowPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceStoreAllowPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "AllowPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreAllowPriceFeedIterator{contract: _PriceStore.contract, event: "AllowPriceFeed", logs: logs, sub: sub}, nil
}

// WatchAllowPriceFeed is a free log subscription operation binding the contract event 0x73720c9031fae32be0c4f58f3b3934399946cfa29d2b13019ead89d904aa950e.
//
// Solidity: event AllowPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) WatchAllowPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceStoreAllowPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "AllowPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreAllowPriceFeed)
				if err := _PriceStore.contract.UnpackLog(event, "AllowPriceFeed", log); err != nil {
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

// ParseAllowPriceFeed is a log parse operation binding the contract event 0x73720c9031fae32be0c4f58f3b3934399946cfa29d2b13019ead89d904aa950e.
//
// Solidity: event AllowPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) ParseAllowPriceFeed(log types.Log) (*PriceStoreAllowPriceFeed, error) {
	event := new(PriceStoreAllowPriceFeed)
	if err := _PriceStore.contract.UnpackLog(event, "AllowPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceStoreForbidPriceFeedIterator is returned from FilterForbidPriceFeed and is used to iterate over the raw logs and unpacked data for ForbidPriceFeed events raised by the PriceStore contract.
type PriceStoreForbidPriceFeedIterator struct {
	Event *PriceStoreForbidPriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceStoreForbidPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreForbidPriceFeed)
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
		it.Event = new(PriceStoreForbidPriceFeed)
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
func (it *PriceStoreForbidPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreForbidPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreForbidPriceFeed represents a ForbidPriceFeed event raised by the PriceStore contract.
type PriceStoreForbidPriceFeed struct {
	Token     common.Address
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForbidPriceFeed is a free log retrieval operation binding the contract event 0x6b88c312dae29edcabbccaba2c987c9b3f61480dd7ed03bd7f932f717dccc54f.
//
// Solidity: event ForbidPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) FilterForbidPriceFeed(opts *bind.FilterOpts, token []common.Address, priceFeed []common.Address) (*PriceStoreForbidPriceFeedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "ForbidPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreForbidPriceFeedIterator{contract: _PriceStore.contract, event: "ForbidPriceFeed", logs: logs, sub: sub}, nil
}

// WatchForbidPriceFeed is a free log subscription operation binding the contract event 0x6b88c312dae29edcabbccaba2c987c9b3f61480dd7ed03bd7f932f717dccc54f.
//
// Solidity: event ForbidPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) WatchForbidPriceFeed(opts *bind.WatchOpts, sink chan<- *PriceStoreForbidPriceFeed, token []common.Address, priceFeed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "ForbidPriceFeed", tokenRule, priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreForbidPriceFeed)
				if err := _PriceStore.contract.UnpackLog(event, "ForbidPriceFeed", log); err != nil {
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

// ParseForbidPriceFeed is a log parse operation binding the contract event 0x6b88c312dae29edcabbccaba2c987c9b3f61480dd7ed03bd7f932f717dccc54f.
//
// Solidity: event ForbidPriceFeed(address indexed token, address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) ParseForbidPriceFeed(log types.Log) (*PriceStoreForbidPriceFeed, error) {
	event := new(PriceStoreForbidPriceFeed)
	if err := _PriceStore.contract.UnpackLog(event, "ForbidPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceStoreRemovePriceFeedIterator is returned from FilterRemovePriceFeed and is used to iterate over the raw logs and unpacked data for RemovePriceFeed events raised by the PriceStore contract.
type PriceStoreRemovePriceFeedIterator struct {
	Event *PriceStoreRemovePriceFeed // Event containing the contract specifics and raw log

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
func (it *PriceStoreRemovePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreRemovePriceFeed)
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
		it.Event = new(PriceStoreRemovePriceFeed)
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
func (it *PriceStoreRemovePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreRemovePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreRemovePriceFeed represents a RemovePriceFeed event raised by the PriceStore contract.
type PriceStoreRemovePriceFeed struct {
	PriceFeed common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovePriceFeed is a free log retrieval operation binding the contract event 0x0d34c8e3483f57054e96fce18e1b58c102738aca38e401f22ce095d181dc7dba.
//
// Solidity: event RemovePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) FilterRemovePriceFeed(opts *bind.FilterOpts, priceFeed []common.Address) (*PriceStoreRemovePriceFeedIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "RemovePriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreRemovePriceFeedIterator{contract: _PriceStore.contract, event: "RemovePriceFeed", logs: logs, sub: sub}, nil
}

// WatchRemovePriceFeed is a free log subscription operation binding the contract event 0x0d34c8e3483f57054e96fce18e1b58c102738aca38e401f22ce095d181dc7dba.
//
// Solidity: event RemovePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) WatchRemovePriceFeed(opts *bind.WatchOpts, sink chan<- *PriceStoreRemovePriceFeed, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "RemovePriceFeed", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreRemovePriceFeed)
				if err := _PriceStore.contract.UnpackLog(event, "RemovePriceFeed", log); err != nil {
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

// ParseRemovePriceFeed is a log parse operation binding the contract event 0x0d34c8e3483f57054e96fce18e1b58c102738aca38e401f22ce095d181dc7dba.
//
// Solidity: event RemovePriceFeed(address indexed priceFeed)
func (_PriceStore *PriceStoreFilterer) ParseRemovePriceFeed(log types.Log) (*PriceStoreRemovePriceFeed, error) {
	event := new(PriceStoreRemovePriceFeed)
	if err := _PriceStore.contract.UnpackLog(event, "RemovePriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceStoreSetStalenessPeriodIterator is returned from FilterSetStalenessPeriod and is used to iterate over the raw logs and unpacked data for SetStalenessPeriod events raised by the PriceStore contract.
type PriceStoreSetStalenessPeriodIterator struct {
	Event *PriceStoreSetStalenessPeriod // Event containing the contract specifics and raw log

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
func (it *PriceStoreSetStalenessPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceStoreSetStalenessPeriod)
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
		it.Event = new(PriceStoreSetStalenessPeriod)
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
func (it *PriceStoreSetStalenessPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceStoreSetStalenessPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceStoreSetStalenessPeriod represents a SetStalenessPeriod event raised by the PriceStore contract.
type PriceStoreSetStalenessPeriod struct {
	PriceFeed       common.Address
	StalenessPeriod uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetStalenessPeriod is a free log retrieval operation binding the contract event 0xd85fed169ea0f7852e9ec4becfec49eca5fe689c65e95486e97ab1f4f02623b9.
//
// Solidity: event SetStalenessPeriod(address indexed priceFeed, uint32 stalenessPeriod)
func (_PriceStore *PriceStoreFilterer) FilterSetStalenessPeriod(opts *bind.FilterOpts, priceFeed []common.Address) (*PriceStoreSetStalenessPeriodIterator, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.FilterLogs(opts, "SetStalenessPeriod", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return &PriceStoreSetStalenessPeriodIterator{contract: _PriceStore.contract, event: "SetStalenessPeriod", logs: logs, sub: sub}, nil
}

// WatchSetStalenessPeriod is a free log subscription operation binding the contract event 0xd85fed169ea0f7852e9ec4becfec49eca5fe689c65e95486e97ab1f4f02623b9.
//
// Solidity: event SetStalenessPeriod(address indexed priceFeed, uint32 stalenessPeriod)
func (_PriceStore *PriceStoreFilterer) WatchSetStalenessPeriod(opts *bind.WatchOpts, sink chan<- *PriceStoreSetStalenessPeriod, priceFeed []common.Address) (event.Subscription, error) {

	var priceFeedRule []interface{}
	for _, priceFeedItem := range priceFeed {
		priceFeedRule = append(priceFeedRule, priceFeedItem)
	}

	logs, sub, err := _PriceStore.contract.WatchLogs(opts, "SetStalenessPeriod", priceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceStoreSetStalenessPeriod)
				if err := _PriceStore.contract.UnpackLog(event, "SetStalenessPeriod", log); err != nil {
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

// ParseSetStalenessPeriod is a log parse operation binding the contract event 0xd85fed169ea0f7852e9ec4becfec49eca5fe689c65e95486e97ab1f4f02623b9.
//
// Solidity: event SetStalenessPeriod(address indexed priceFeed, uint32 stalenessPeriod)
func (_PriceStore *PriceStoreFilterer) ParseSetStalenessPeriod(log types.Log) (*PriceStoreSetStalenessPeriod, error) {
	event := new(PriceStoreSetStalenessPeriod)
	if err := _PriceStore.contract.UnpackLog(event, "SetStalenessPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
