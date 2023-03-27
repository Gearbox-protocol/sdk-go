// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenDistributor

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
)

// TokenDistributorMetaData contains all meta data concerning the TokenDistributor contract.
var TokenDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"contributorsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"}],\"name\":\"contributorVestingContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenDistributorMetaData.ABI instead.
var TokenDistributorABI = TokenDistributorMetaData.ABI

// TokenDistributor is an auto generated Go binding around an Ethereum contract.
type TokenDistributor struct {
	TokenDistributorCaller     // Read-only binding to the contract
	TokenDistributorTransactor // Write-only binding to the contract
	TokenDistributorFilterer   // Log filterer for contract events
}

// TokenDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenDistributorSession struct {
	Contract     *TokenDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenDistributorCallerSession struct {
	Contract *TokenDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TokenDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenDistributorTransactorSession struct {
	Contract     *TokenDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenDistributorRaw struct {
	Contract *TokenDistributor // Generic contract binding to access the raw methods on
}

// TokenDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenDistributorCallerRaw struct {
	Contract *TokenDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TokenDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenDistributorTransactorRaw struct {
	Contract *TokenDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenDistributor creates a new instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributor(address common.Address, backend bind.ContractBackend) (*TokenDistributor, error) {
	contract, err := bindTokenDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenDistributor{TokenDistributorCaller: TokenDistributorCaller{contract: contract}, TokenDistributorTransactor: TokenDistributorTransactor{contract: contract}, TokenDistributorFilterer: TokenDistributorFilterer{contract: contract}}, nil
}

// NewTokenDistributorCaller creates a new read-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorCaller(address common.Address, caller bind.ContractCaller) (*TokenDistributorCaller, error) {
	contract, err := bindTokenDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorCaller{contract: contract}, nil
}

// NewTokenDistributorTransactor creates a new write-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenDistributorTransactor, error) {
	contract, err := bindTokenDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorTransactor{contract: contract}, nil
}

// NewTokenDistributorFilterer creates a new log filterer instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenDistributorFilterer, error) {
	contract, err := bindTokenDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorFilterer{contract: contract}, nil
}

// bindTokenDistributor binds a generic wrapper to an already deployed contract.
func bindTokenDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.TokenDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transact(opts, method, params...)
}

// ContributorVestingContracts is a free data retrieval call binding the contract method 0x48c1de87.
//
// Solidity: function contributorVestingContracts(address contributor) view returns(address[])
func (_TokenDistributor *TokenDistributorCaller) ContributorVestingContracts(opts *bind.CallOpts, contributor common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenDistributor.contract.Call(opts, &out, "contributorVestingContracts", contributor)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ContributorVestingContracts is a free data retrieval call binding the contract method 0x48c1de87.
//
// Solidity: function contributorVestingContracts(address contributor) view returns(address[])
func (_TokenDistributor *TokenDistributorSession) ContributorVestingContracts(contributor common.Address) ([]common.Address, error) {
	return _TokenDistributor.Contract.ContributorVestingContracts(&_TokenDistributor.CallOpts, contributor)
}

// ContributorVestingContracts is a free data retrieval call binding the contract method 0x48c1de87.
//
// Solidity: function contributorVestingContracts(address contributor) view returns(address[])
func (_TokenDistributor *TokenDistributorCallerSession) ContributorVestingContracts(contributor common.Address) ([]common.Address, error) {
	return _TokenDistributor.Contract.ContributorVestingContracts(&_TokenDistributor.CallOpts, contributor)
}

// ContributorsList is a free data retrieval call binding the contract method 0x45f02a5b.
//
// Solidity: function contributorsList() view returns(address[])
func (_TokenDistributor *TokenDistributorCaller) ContributorsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenDistributor.contract.Call(opts, &out, "contributorsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ContributorsList is a free data retrieval call binding the contract method 0x45f02a5b.
//
// Solidity: function contributorsList() view returns(address[])
func (_TokenDistributor *TokenDistributorSession) ContributorsList() ([]common.Address, error) {
	return _TokenDistributor.Contract.ContributorsList(&_TokenDistributor.CallOpts)
}

// ContributorsList is a free data retrieval call binding the contract method 0x45f02a5b.
//
// Solidity: function contributorsList() view returns(address[])
func (_TokenDistributor *TokenDistributorCallerSession) ContributorsList() ([]common.Address, error) {
	return _TokenDistributor.Contract.ContributorsList(&_TokenDistributor.CallOpts)
}
