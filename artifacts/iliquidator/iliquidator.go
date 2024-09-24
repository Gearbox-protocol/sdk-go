// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iliquidator

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

// IPartialLiquidationBotV3PriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type IPartialLiquidationBotV3PriceUpdate struct {
	Token   common.Address
	Reserve bool
	Data    []byte
}

// LiquidationResult is an auto generated low-level Go binding around an user-defined struct.
type LiquidationResult struct {
	Calls     []MultiCall
	Profit    *big.Int
	AmountIn  *big.Int
	AmountOut *big.Int
}

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// IliquidatorMetaData contains all meta data concerning the Iliquidator contract.
var IliquidatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cmToCA\",\"inputs\":[{\"name\":\"cm\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOptimalLiquidation\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hfOptimal\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structIPartialLiquidationBotV3.PriceUpdate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"repaidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flashLoanAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isOptimalRepayable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"partialLiquidateAndConvert\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flashLoanAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structIPartialLiquidationBotV3.PriceUpdate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"conversionCalls\",\"type\":\"tuple[]\",\"internalType\":\"structMultiCall[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"partialLiquidationBot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewPartialLiquidation\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flashLoanAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"priceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structIPartialLiquidationBotV3.PriceUpdate[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reserve\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"connectors\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"slippage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"res\",\"type\":\"tuple\",\"internalType\":\"structLiquidationResult\",\"components\":[{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structMultiCall[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"profit\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerCM\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPartialLiquidationBot\",\"inputs\":[{\"name\":\"newPLB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IliquidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IliquidatorMetaData.ABI instead.
var IliquidatorABI = IliquidatorMetaData.ABI

// Iliquidator is an auto generated Go binding around an Ethereum contract.
type Iliquidator struct {
	IliquidatorCaller     // Read-only binding to the contract
	IliquidatorTransactor // Write-only binding to the contract
	IliquidatorFilterer   // Log filterer for contract events
}

// IliquidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IliquidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IliquidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IliquidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IliquidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IliquidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IliquidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IliquidatorSession struct {
	Contract     *Iliquidator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IliquidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IliquidatorCallerSession struct {
	Contract *IliquidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IliquidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IliquidatorTransactorSession struct {
	Contract     *IliquidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IliquidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IliquidatorRaw struct {
	Contract *Iliquidator // Generic contract binding to access the raw methods on
}

// IliquidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IliquidatorCallerRaw struct {
	Contract *IliquidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IliquidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IliquidatorTransactorRaw struct {
	Contract *IliquidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIliquidator creates a new instance of Iliquidator, bound to a specific deployed contract.
func NewIliquidator(address common.Address, backend bind.ContractBackend) (*Iliquidator, error) {
	contract, err := bindIliquidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iliquidator{IliquidatorCaller: IliquidatorCaller{contract: contract}, IliquidatorTransactor: IliquidatorTransactor{contract: contract}, IliquidatorFilterer: IliquidatorFilterer{contract: contract}}, nil
}

// NewIliquidatorCaller creates a new read-only instance of Iliquidator, bound to a specific deployed contract.
func NewIliquidatorCaller(address common.Address, caller bind.ContractCaller) (*IliquidatorCaller, error) {
	contract, err := bindIliquidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IliquidatorCaller{contract: contract}, nil
}

// NewIliquidatorTransactor creates a new write-only instance of Iliquidator, bound to a specific deployed contract.
func NewIliquidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IliquidatorTransactor, error) {
	contract, err := bindIliquidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IliquidatorTransactor{contract: contract}, nil
}

// NewIliquidatorFilterer creates a new log filterer instance of Iliquidator, bound to a specific deployed contract.
func NewIliquidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IliquidatorFilterer, error) {
	contract, err := bindIliquidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IliquidatorFilterer{contract: contract}, nil
}

// bindIliquidator binds a generic wrapper to an already deployed contract.
func bindIliquidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IliquidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iliquidator *IliquidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iliquidator.Contract.IliquidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iliquidator *IliquidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iliquidator.Contract.IliquidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iliquidator *IliquidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iliquidator.Contract.IliquidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iliquidator *IliquidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iliquidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iliquidator *IliquidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iliquidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iliquidator *IliquidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iliquidator.Contract.contract.Transact(opts, method, params...)
}

// CmToCA is a free data retrieval call binding the contract method 0xecb6c7cf.
//
// Solidity: function cmToCA(address cm) view returns(address)
func (_Iliquidator *IliquidatorCaller) CmToCA(opts *bind.CallOpts, cm common.Address) (common.Address, error) {
	var out []interface{}
	err := _Iliquidator.contract.Call(opts, &out, "cmToCA", cm)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CmToCA is a free data retrieval call binding the contract method 0xecb6c7cf.
//
// Solidity: function cmToCA(address cm) view returns(address)
func (_Iliquidator *IliquidatorSession) CmToCA(cm common.Address) (common.Address, error) {
	return _Iliquidator.Contract.CmToCA(&_Iliquidator.CallOpts, cm)
}

// CmToCA is a free data retrieval call binding the contract method 0xecb6c7cf.
//
// Solidity: function cmToCA(address cm) view returns(address)
func (_Iliquidator *IliquidatorCallerSession) CmToCA(cm common.Address) (common.Address, error) {
	return _Iliquidator.Contract.CmToCA(&_Iliquidator.CallOpts, cm)
}

// PartialLiquidationBot is a free data retrieval call binding the contract method 0x789ea176.
//
// Solidity: function partialLiquidationBot() view returns(address)
func (_Iliquidator *IliquidatorCaller) PartialLiquidationBot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iliquidator.contract.Call(opts, &out, "partialLiquidationBot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PartialLiquidationBot is a free data retrieval call binding the contract method 0x789ea176.
//
// Solidity: function partialLiquidationBot() view returns(address)
func (_Iliquidator *IliquidatorSession) PartialLiquidationBot() (common.Address, error) {
	return _Iliquidator.Contract.PartialLiquidationBot(&_Iliquidator.CallOpts)
}

// PartialLiquidationBot is a free data retrieval call binding the contract method 0x789ea176.
//
// Solidity: function partialLiquidationBot() view returns(address)
func (_Iliquidator *IliquidatorCallerSession) PartialLiquidationBot() (common.Address, error) {
	return _Iliquidator.Contract.PartialLiquidationBot(&_Iliquidator.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Iliquidator *IliquidatorCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iliquidator.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Iliquidator *IliquidatorSession) Router() (common.Address, error) {
	return _Iliquidator.Contract.Router(&_Iliquidator.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Iliquidator *IliquidatorCallerSession) Router() (common.Address, error) {
	return _Iliquidator.Contract.Router(&_Iliquidator.CallOpts)
}

// GetOptimalLiquidation is a paid mutator transaction binding the contract method 0xc9a523da.
//
// Solidity: function getOptimalLiquidation(address creditAccount, uint256 hfOptimal, (address,bool,bytes)[] priceUpdates) returns(address tokenOut, uint256 optimalAmount, uint256 repaidAmount, uint256 flashLoanAmount, bool isOptimalRepayable)
func (_Iliquidator *IliquidatorTransactor) GetOptimalLiquidation(opts *bind.TransactOpts, creditAccount common.Address, hfOptimal *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "getOptimalLiquidation", creditAccount, hfOptimal, priceUpdates)
}

// GetOptimalLiquidation is a paid mutator transaction binding the contract method 0xc9a523da.
//
// Solidity: function getOptimalLiquidation(address creditAccount, uint256 hfOptimal, (address,bool,bytes)[] priceUpdates) returns(address tokenOut, uint256 optimalAmount, uint256 repaidAmount, uint256 flashLoanAmount, bool isOptimalRepayable)
func (_Iliquidator *IliquidatorSession) GetOptimalLiquidation(creditAccount common.Address, hfOptimal *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate) (*types.Transaction, error) {
	return _Iliquidator.Contract.GetOptimalLiquidation(&_Iliquidator.TransactOpts, creditAccount, hfOptimal, priceUpdates)
}

// GetOptimalLiquidation is a paid mutator transaction binding the contract method 0xc9a523da.
//
// Solidity: function getOptimalLiquidation(address creditAccount, uint256 hfOptimal, (address,bool,bytes)[] priceUpdates) returns(address tokenOut, uint256 optimalAmount, uint256 repaidAmount, uint256 flashLoanAmount, bool isOptimalRepayable)
func (_Iliquidator *IliquidatorTransactorSession) GetOptimalLiquidation(creditAccount common.Address, hfOptimal *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate) (*types.Transaction, error) {
	return _Iliquidator.Contract.GetOptimalLiquidation(&_Iliquidator.TransactOpts, creditAccount, hfOptimal, priceUpdates)
}

// PartialLiquidateAndConvert is a paid mutator transaction binding the contract method 0xff024f4d.
//
// Solidity: function partialLiquidateAndConvert(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, (address,bytes)[] conversionCalls) returns()
func (_Iliquidator *IliquidatorTransactor) PartialLiquidateAndConvert(opts *bind.TransactOpts, creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, conversionCalls []MultiCall) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "partialLiquidateAndConvert", creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, conversionCalls)
}

// PartialLiquidateAndConvert is a paid mutator transaction binding the contract method 0xff024f4d.
//
// Solidity: function partialLiquidateAndConvert(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, (address,bytes)[] conversionCalls) returns()
func (_Iliquidator *IliquidatorSession) PartialLiquidateAndConvert(creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, conversionCalls []MultiCall) (*types.Transaction, error) {
	return _Iliquidator.Contract.PartialLiquidateAndConvert(&_Iliquidator.TransactOpts, creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, conversionCalls)
}

// PartialLiquidateAndConvert is a paid mutator transaction binding the contract method 0xff024f4d.
//
// Solidity: function partialLiquidateAndConvert(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, (address,bytes)[] conversionCalls) returns()
func (_Iliquidator *IliquidatorTransactorSession) PartialLiquidateAndConvert(creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, conversionCalls []MultiCall) (*types.Transaction, error) {
	return _Iliquidator.Contract.PartialLiquidateAndConvert(&_Iliquidator.TransactOpts, creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, conversionCalls)
}

// PreviewPartialLiquidation is a paid mutator transaction binding the contract method 0x8c3b7362.
//
// Solidity: function previewPartialLiquidation(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, address[] connectors, uint256 slippage) returns(((address,bytes)[],int256,uint256,uint256) res)
func (_Iliquidator *IliquidatorTransactor) PreviewPartialLiquidation(opts *bind.TransactOpts, creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "previewPartialLiquidation", creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, connectors, slippage)
}

// PreviewPartialLiquidation is a paid mutator transaction binding the contract method 0x8c3b7362.
//
// Solidity: function previewPartialLiquidation(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, address[] connectors, uint256 slippage) returns(((address,bytes)[],int256,uint256,uint256) res)
func (_Iliquidator *IliquidatorSession) PreviewPartialLiquidation(creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Iliquidator.Contract.PreviewPartialLiquidation(&_Iliquidator.TransactOpts, creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, connectors, slippage)
}

// PreviewPartialLiquidation is a paid mutator transaction binding the contract method 0x8c3b7362.
//
// Solidity: function previewPartialLiquidation(address creditManager, address creditAccount, address assetOut, uint256 amountOut, uint256 flashLoanAmount, (address,bool,bytes)[] priceUpdates, address[] connectors, uint256 slippage) returns(((address,bytes)[],int256,uint256,uint256) res)
func (_Iliquidator *IliquidatorTransactorSession) PreviewPartialLiquidation(creditManager common.Address, creditAccount common.Address, assetOut common.Address, amountOut *big.Int, flashLoanAmount *big.Int, priceUpdates []IPartialLiquidationBotV3PriceUpdate, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Iliquidator.Contract.PreviewPartialLiquidation(&_Iliquidator.TransactOpts, creditManager, creditAccount, assetOut, amountOut, flashLoanAmount, priceUpdates, connectors, slippage)
}

// RegisterCM is a paid mutator transaction binding the contract method 0xb60beeb8.
//
// Solidity: function registerCM(address creditManager) returns()
func (_Iliquidator *IliquidatorTransactor) RegisterCM(opts *bind.TransactOpts, creditManager common.Address) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "registerCM", creditManager)
}

// RegisterCM is a paid mutator transaction binding the contract method 0xb60beeb8.
//
// Solidity: function registerCM(address creditManager) returns()
func (_Iliquidator *IliquidatorSession) RegisterCM(creditManager common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.RegisterCM(&_Iliquidator.TransactOpts, creditManager)
}

// RegisterCM is a paid mutator transaction binding the contract method 0xb60beeb8.
//
// Solidity: function registerCM(address creditManager) returns()
func (_Iliquidator *IliquidatorTransactorSession) RegisterCM(creditManager common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.RegisterCM(&_Iliquidator.TransactOpts, creditManager)
}

// SetPartialLiquidationBot is a paid mutator transaction binding the contract method 0xb58f18c8.
//
// Solidity: function setPartialLiquidationBot(address newPLB) returns()
func (_Iliquidator *IliquidatorTransactor) SetPartialLiquidationBot(opts *bind.TransactOpts, newPLB common.Address) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "setPartialLiquidationBot", newPLB)
}

// SetPartialLiquidationBot is a paid mutator transaction binding the contract method 0xb58f18c8.
//
// Solidity: function setPartialLiquidationBot(address newPLB) returns()
func (_Iliquidator *IliquidatorSession) SetPartialLiquidationBot(newPLB common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.SetPartialLiquidationBot(&_Iliquidator.TransactOpts, newPLB)
}

// SetPartialLiquidationBot is a paid mutator transaction binding the contract method 0xb58f18c8.
//
// Solidity: function setPartialLiquidationBot(address newPLB) returns()
func (_Iliquidator *IliquidatorTransactorSession) SetPartialLiquidationBot(newPLB common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.SetPartialLiquidationBot(&_Iliquidator.TransactOpts, newPLB)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_Iliquidator *IliquidatorTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "setRouter", newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_Iliquidator *IliquidatorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.SetRouter(&_Iliquidator.TransactOpts, newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_Iliquidator *IliquidatorTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.SetRouter(&_Iliquidator.TransactOpts, newRouter)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address to) returns()
func (_Iliquidator *IliquidatorTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Iliquidator.contract.Transact(opts, "withdrawToken", token, amount, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address to) returns()
func (_Iliquidator *IliquidatorSession) WithdrawToken(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.WithdrawToken(&_Iliquidator.TransactOpts, token, amount, to)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address to) returns()
func (_Iliquidator *IliquidatorTransactorSession) WithdrawToken(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Iliquidator.Contract.WithdrawToken(&_Iliquidator.TransactOpts, token, amount, to)
}
