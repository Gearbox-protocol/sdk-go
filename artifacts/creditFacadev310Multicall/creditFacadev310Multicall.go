// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFacadev310Multicall

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

// BalanceDelta is an auto generated low-level Go binding around an user-defined struct.
type BalanceDelta struct {
	Token  common.Address
	Amount *big.Int
}

// PriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type PriceUpdate struct {
	PriceFeed common.Address
	Data      []byte
}

// CreditFacadev310MulticallMetaData contains all meta data concerning the CreditFacadev310Multicall contract.
var CreditFacadev310MulticallMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addCollateral\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCollateralWithPermit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compareBalances\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseDebt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseDebt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onDemandPriceUpdates\",\"inputs\":[{\"name\":\"updates\",\"type\":\"tuple[]\",\"internalType\":\"structPriceUpdate[]\",\"components\":[{\"name\":\"priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBotPermissions\",\"inputs\":[{\"name\":\"bot\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissions\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFullCheckParams\",\"inputs\":[{\"name\":\"collateralHints\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"storeExpectedBalances\",\"inputs\":[{\"name\":\"balanceDeltas\",\"type\":\"tuple[]\",\"internalType\":\"structBalanceDelta[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"int256\",\"internalType\":\"int256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateQuota\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quotaChange\",\"type\":\"int96\",\"internalType\":\"int96\"},{\"name\":\"minQuota\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawCollateral\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// CreditFacadev310MulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditFacadev310MulticallMetaData.ABI instead.
var CreditFacadev310MulticallABI = CreditFacadev310MulticallMetaData.ABI

// CreditFacadev310Multicall is an auto generated Go binding around an Ethereum contract.
type CreditFacadev310Multicall struct {
	CreditFacadev310MulticallCaller     // Read-only binding to the contract
	CreditFacadev310MulticallTransactor // Write-only binding to the contract
	CreditFacadev310MulticallFilterer   // Log filterer for contract events
}

// CreditFacadev310MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFacadev310MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev310MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFacadev310MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev310MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFacadev310MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev310MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFacadev310MulticallSession struct {
	Contract     *CreditFacadev310Multicall // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CreditFacadev310MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFacadev310MulticallCallerSession struct {
	Contract *CreditFacadev310MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CreditFacadev310MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFacadev310MulticallTransactorSession struct {
	Contract     *CreditFacadev310MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CreditFacadev310MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFacadev310MulticallRaw struct {
	Contract *CreditFacadev310Multicall // Generic contract binding to access the raw methods on
}

// CreditFacadev310MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFacadev310MulticallCallerRaw struct {
	Contract *CreditFacadev310MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// CreditFacadev310MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFacadev310MulticallTransactorRaw struct {
	Contract *CreditFacadev310MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFacadev310Multicall creates a new instance of CreditFacadev310Multicall, bound to a specific deployed contract.
func NewCreditFacadev310Multicall(address common.Address, backend bind.ContractBackend) (*CreditFacadev310Multicall, error) {
	contract, err := bindCreditFacadev310Multicall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev310Multicall{CreditFacadev310MulticallCaller: CreditFacadev310MulticallCaller{contract: contract}, CreditFacadev310MulticallTransactor: CreditFacadev310MulticallTransactor{contract: contract}, CreditFacadev310MulticallFilterer: CreditFacadev310MulticallFilterer{contract: contract}}, nil
}

// NewCreditFacadev310MulticallCaller creates a new read-only instance of CreditFacadev310Multicall, bound to a specific deployed contract.
func NewCreditFacadev310MulticallCaller(address common.Address, caller bind.ContractCaller) (*CreditFacadev310MulticallCaller, error) {
	contract, err := bindCreditFacadev310Multicall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev310MulticallCaller{contract: contract}, nil
}

// NewCreditFacadev310MulticallTransactor creates a new write-only instance of CreditFacadev310Multicall, bound to a specific deployed contract.
func NewCreditFacadev310MulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditFacadev310MulticallTransactor, error) {
	contract, err := bindCreditFacadev310Multicall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev310MulticallTransactor{contract: contract}, nil
}

// NewCreditFacadev310MulticallFilterer creates a new log filterer instance of CreditFacadev310Multicall, bound to a specific deployed contract.
func NewCreditFacadev310MulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFacadev310MulticallFilterer, error) {
	contract, err := bindCreditFacadev310Multicall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev310MulticallFilterer{contract: contract}, nil
}

// bindCreditFacadev310Multicall binds a generic wrapper to an already deployed contract.
func bindCreditFacadev310Multicall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditFacadev310MulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadev310Multicall.Contract.CreditFacadev310MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.CreditFacadev310MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.CreditFacadev310MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadev310Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.contract.Transact(opts, method, params...)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(address token, uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) AddCollateral(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "addCollateral", token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(address token, uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) AddCollateral(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.AddCollateral(&_CreditFacadev310Multicall.TransactOpts, token, amount)
}

// AddCollateral is a paid mutator transaction binding the contract method 0x6d75b9ee.
//
// Solidity: function addCollateral(address token, uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) AddCollateral(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.AddCollateral(&_CreditFacadev310Multicall.TransactOpts, token, amount)
}

// AddCollateralWithPermit is a paid mutator transaction binding the contract method 0x438f8fcc.
//
// Solidity: function addCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) AddCollateralWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "addCollateralWithPermit", token, amount, deadline, v, r, s)
}

// AddCollateralWithPermit is a paid mutator transaction binding the contract method 0x438f8fcc.
//
// Solidity: function addCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) AddCollateralWithPermit(token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.AddCollateralWithPermit(&_CreditFacadev310Multicall.TransactOpts, token, amount, deadline, v, r, s)
}

// AddCollateralWithPermit is a paid mutator transaction binding the contract method 0x438f8fcc.
//
// Solidity: function addCollateralWithPermit(address token, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) AddCollateralWithPermit(token common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.AddCollateralWithPermit(&_CreditFacadev310Multicall.TransactOpts, token, amount, deadline, v, r, s)
}

// CompareBalances is a paid mutator transaction binding the contract method 0xf42aeb00.
//
// Solidity: function compareBalances() returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) CompareBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "compareBalances")
}

// CompareBalances is a paid mutator transaction binding the contract method 0xf42aeb00.
//
// Solidity: function compareBalances() returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) CompareBalances() (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.CompareBalances(&_CreditFacadev310Multicall.TransactOpts)
}

// CompareBalances is a paid mutator transaction binding the contract method 0xf42aeb00.
//
// Solidity: function compareBalances() returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) CompareBalances() (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.CompareBalances(&_CreditFacadev310Multicall.TransactOpts)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) DecreaseDebt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "decreaseDebt", amount)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) DecreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.DecreaseDebt(&_CreditFacadev310Multicall.TransactOpts, amount)
}

// DecreaseDebt is a paid mutator transaction binding the contract method 0x2a7ba1f7.
//
// Solidity: function decreaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) DecreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.DecreaseDebt(&_CreditFacadev310Multicall.TransactOpts, amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) IncreaseDebt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "increaseDebt", amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) IncreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.IncreaseDebt(&_CreditFacadev310Multicall.TransactOpts, amount)
}

// IncreaseDebt is a paid mutator transaction binding the contract method 0x2b7c7b11.
//
// Solidity: function increaseDebt(uint256 amount) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) IncreaseDebt(amount *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.IncreaseDebt(&_CreditFacadev310Multicall.TransactOpts, amount)
}

// OnDemandPriceUpdates is a paid mutator transaction binding the contract method 0x28b83c48.
//
// Solidity: function onDemandPriceUpdates((address,bytes)[] updates) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) OnDemandPriceUpdates(opts *bind.TransactOpts, updates []PriceUpdate) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "onDemandPriceUpdates", updates)
}

// OnDemandPriceUpdates is a paid mutator transaction binding the contract method 0x28b83c48.
//
// Solidity: function onDemandPriceUpdates((address,bytes)[] updates) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) OnDemandPriceUpdates(updates []PriceUpdate) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.OnDemandPriceUpdates(&_CreditFacadev310Multicall.TransactOpts, updates)
}

// OnDemandPriceUpdates is a paid mutator transaction binding the contract method 0x28b83c48.
//
// Solidity: function onDemandPriceUpdates((address,bytes)[] updates) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) OnDemandPriceUpdates(updates []PriceUpdate) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.OnDemandPriceUpdates(&_CreditFacadev310Multicall.TransactOpts, updates)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0x19043543.
//
// Solidity: function setBotPermissions(address bot, uint192 permissions) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) SetBotPermissions(opts *bind.TransactOpts, bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "setBotPermissions", bot, permissions)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0x19043543.
//
// Solidity: function setBotPermissions(address bot, uint192 permissions) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) SetBotPermissions(bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.SetBotPermissions(&_CreditFacadev310Multicall.TransactOpts, bot, permissions)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0x19043543.
//
// Solidity: function setBotPermissions(address bot, uint192 permissions) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) SetBotPermissions(bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.SetBotPermissions(&_CreditFacadev310Multicall.TransactOpts, bot, permissions)
}

// SetFullCheckParams is a paid mutator transaction binding the contract method 0x0768bbfe.
//
// Solidity: function setFullCheckParams(uint256[] collateralHints, uint16 minHealthFactor) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) SetFullCheckParams(opts *bind.TransactOpts, collateralHints []*big.Int, minHealthFactor uint16) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "setFullCheckParams", collateralHints, minHealthFactor)
}

// SetFullCheckParams is a paid mutator transaction binding the contract method 0x0768bbfe.
//
// Solidity: function setFullCheckParams(uint256[] collateralHints, uint16 minHealthFactor) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) SetFullCheckParams(collateralHints []*big.Int, minHealthFactor uint16) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.SetFullCheckParams(&_CreditFacadev310Multicall.TransactOpts, collateralHints, minHealthFactor)
}

// SetFullCheckParams is a paid mutator transaction binding the contract method 0x0768bbfe.
//
// Solidity: function setFullCheckParams(uint256[] collateralHints, uint16 minHealthFactor) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) SetFullCheckParams(collateralHints []*big.Int, minHealthFactor uint16) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.SetFullCheckParams(&_CreditFacadev310Multicall.TransactOpts, collateralHints, minHealthFactor)
}

// StoreExpectedBalances is a paid mutator transaction binding the contract method 0x2f2ca49b.
//
// Solidity: function storeExpectedBalances((address,int256)[] balanceDeltas) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) StoreExpectedBalances(opts *bind.TransactOpts, balanceDeltas []BalanceDelta) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "storeExpectedBalances", balanceDeltas)
}

// StoreExpectedBalances is a paid mutator transaction binding the contract method 0x2f2ca49b.
//
// Solidity: function storeExpectedBalances((address,int256)[] balanceDeltas) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) StoreExpectedBalances(balanceDeltas []BalanceDelta) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.StoreExpectedBalances(&_CreditFacadev310Multicall.TransactOpts, balanceDeltas)
}

// StoreExpectedBalances is a paid mutator transaction binding the contract method 0x2f2ca49b.
//
// Solidity: function storeExpectedBalances((address,int256)[] balanceDeltas) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) StoreExpectedBalances(balanceDeltas []BalanceDelta) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.StoreExpectedBalances(&_CreditFacadev310Multicall.TransactOpts, balanceDeltas)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x712c10ad.
//
// Solidity: function updateQuota(address token, int96 quotaChange, uint96 minQuota) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) UpdateQuota(opts *bind.TransactOpts, token common.Address, quotaChange *big.Int, minQuota *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "updateQuota", token, quotaChange, minQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x712c10ad.
//
// Solidity: function updateQuota(address token, int96 quotaChange, uint96 minQuota) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) UpdateQuota(token common.Address, quotaChange *big.Int, minQuota *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.UpdateQuota(&_CreditFacadev310Multicall.TransactOpts, token, quotaChange, minQuota)
}

// UpdateQuota is a paid mutator transaction binding the contract method 0x712c10ad.
//
// Solidity: function updateQuota(address token, int96 quotaChange, uint96 minQuota) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) UpdateQuota(token common.Address, quotaChange *big.Int, minQuota *big.Int) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.UpdateQuota(&_CreditFacadev310Multicall.TransactOpts, token, quotaChange, minQuota)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x1f1088a0.
//
// Solidity: function withdrawCollateral(address token, uint256 amount, address to) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactor) WithdrawCollateral(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.contract.Transact(opts, "withdrawCollateral", token, amount, to)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x1f1088a0.
//
// Solidity: function withdrawCollateral(address token, uint256 amount, address to) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallSession) WithdrawCollateral(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.WithdrawCollateral(&_CreditFacadev310Multicall.TransactOpts, token, amount, to)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x1f1088a0.
//
// Solidity: function withdrawCollateral(address token, uint256 amount, address to) returns()
func (_CreditFacadev310Multicall *CreditFacadev310MulticallTransactorSession) WithdrawCollateral(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _CreditFacadev310Multicall.Contract.WithdrawCollateral(&_CreditFacadev310Multicall.TransactOpts, token, amount, to)
}
