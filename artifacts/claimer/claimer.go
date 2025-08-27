// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package claimer

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

// ClaimerMetaData contains all meta data concerning the Claimer contract.
var ClaimerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"multiVault\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"subvaultIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"indices\",\"type\":\"uint256[][]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAssets\",\"type\":\"uint256\"}],\"name\":\"multiAcceptAndClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClaimerABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimerMetaData.ABI instead.
var ClaimerABI = ClaimerMetaData.ABI

// Claimer is an auto generated Go binding around an Ethereum contract.
type Claimer struct {
	ClaimerCaller     // Read-only binding to the contract
	ClaimerTransactor // Write-only binding to the contract
	ClaimerFilterer   // Log filterer for contract events
}

// ClaimerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimerSession struct {
	Contract     *Claimer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimerCallerSession struct {
	Contract *ClaimerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ClaimerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimerTransactorSession struct {
	Contract     *ClaimerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ClaimerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimerRaw struct {
	Contract *Claimer // Generic contract binding to access the raw methods on
}

// ClaimerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimerCallerRaw struct {
	Contract *ClaimerCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimerTransactorRaw struct {
	Contract *ClaimerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimer creates a new instance of Claimer, bound to a specific deployed contract.
func NewClaimer(address common.Address, backend bind.ContractBackend) (*Claimer, error) {
	contract, err := bindClaimer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claimer{ClaimerCaller: ClaimerCaller{contract: contract}, ClaimerTransactor: ClaimerTransactor{contract: contract}, ClaimerFilterer: ClaimerFilterer{contract: contract}}, nil
}

// NewClaimerCaller creates a new read-only instance of Claimer, bound to a specific deployed contract.
func NewClaimerCaller(address common.Address, caller bind.ContractCaller) (*ClaimerCaller, error) {
	contract, err := bindClaimer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimerCaller{contract: contract}, nil
}

// NewClaimerTransactor creates a new write-only instance of Claimer, bound to a specific deployed contract.
func NewClaimerTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimerTransactor, error) {
	contract, err := bindClaimer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimerTransactor{contract: contract}, nil
}

// NewClaimerFilterer creates a new log filterer instance of Claimer, bound to a specific deployed contract.
func NewClaimerFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimerFilterer, error) {
	contract, err := bindClaimer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimerFilterer{contract: contract}, nil
}

// bindClaimer binds a generic wrapper to an already deployed contract.
func bindClaimer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimer *ClaimerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claimer.Contract.ClaimerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimer *ClaimerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimer.Contract.ClaimerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimer *ClaimerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimer.Contract.ClaimerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimer *ClaimerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claimer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimer *ClaimerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimer *ClaimerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimer.Contract.contract.Transact(opts, method, params...)
}

// MultiAcceptAndClaim is a paid mutator transaction binding the contract method 0x82ef9d32.
//
// Solidity: function multiAcceptAndClaim(address multiVault, uint256[] subvaultIndices, uint256[][] indices, address recipient, uint256 maxAssets) returns(uint256 assets)
func (_Claimer *ClaimerTransactor) MultiAcceptAndClaim(opts *bind.TransactOpts, multiVault common.Address, subvaultIndices []*big.Int, indices [][]*big.Int, recipient common.Address, maxAssets *big.Int) (*types.Transaction, error) {
	return _Claimer.contract.Transact(opts, "multiAcceptAndClaim", multiVault, subvaultIndices, indices, recipient, maxAssets)
}

// MultiAcceptAndClaim is a paid mutator transaction binding the contract method 0x82ef9d32.
//
// Solidity: function multiAcceptAndClaim(address multiVault, uint256[] subvaultIndices, uint256[][] indices, address recipient, uint256 maxAssets) returns(uint256 assets)
func (_Claimer *ClaimerSession) MultiAcceptAndClaim(multiVault common.Address, subvaultIndices []*big.Int, indices [][]*big.Int, recipient common.Address, maxAssets *big.Int) (*types.Transaction, error) {
	return _Claimer.Contract.MultiAcceptAndClaim(&_Claimer.TransactOpts, multiVault, subvaultIndices, indices, recipient, maxAssets)
}

// MultiAcceptAndClaim is a paid mutator transaction binding the contract method 0x82ef9d32.
//
// Solidity: function multiAcceptAndClaim(address multiVault, uint256[] subvaultIndices, uint256[][] indices, address recipient, uint256 maxAssets) returns(uint256 assets)
func (_Claimer *ClaimerTransactorSession) MultiAcceptAndClaim(multiVault common.Address, subvaultIndices []*big.Int, indices [][]*big.Int, recipient common.Address, maxAssets *big.Int) (*types.Transaction, error) {
	return _Claimer.Contract.MultiAcceptAndClaim(&_Claimer.TransactOpts, multiVault, subvaultIndices, indices, recipient, maxAssets)
}
