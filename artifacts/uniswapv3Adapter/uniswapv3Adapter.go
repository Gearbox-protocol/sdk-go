// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3Adapter

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IUniswapV3AdapterExactAllInputParams is an auto generated low-level Go binding around an user-defined struct.
type IUniswapV3AdapterExactAllInputParams struct {
	Path       []byte
	Deadline   *big.Int
	RateMinRAY *big.Int
}

// IUniswapV3AdapterExactAllInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IUniswapV3AdapterExactAllInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Deadline          *big.Int
	RateMinRAY        *big.Int
	SqrtPriceLimitX96 *big.Int
}

// Uniswapv3AdapterMetaData contains all meta data concerning the Uniswapv3Adapter contract.
var Uniswapv3AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectPathLengthException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"internalType\":\"structIUniswapV3Adapter.ExactAllInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactAllInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIUniswapV3Adapter.ExactAllInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactAllInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv3AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3AdapterMetaData.ABI instead.
var Uniswapv3AdapterABI = Uniswapv3AdapterMetaData.ABI

// Uniswapv3Adapter is an auto generated Go binding around an Ethereum contract.
type Uniswapv3Adapter struct {
	Uniswapv3AdapterCaller     // Read-only binding to the contract
	Uniswapv3AdapterTransactor // Write-only binding to the contract
	Uniswapv3AdapterFilterer   // Log filterer for contract events
}

// Uniswapv3AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3AdapterSession struct {
	Contract     *Uniswapv3Adapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv3AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3AdapterCallerSession struct {
	Contract *Uniswapv3AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Uniswapv3AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3AdapterTransactorSession struct {
	Contract     *Uniswapv3AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Uniswapv3AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3AdapterRaw struct {
	Contract *Uniswapv3Adapter // Generic contract binding to access the raw methods on
}

// Uniswapv3AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3AdapterCallerRaw struct {
	Contract *Uniswapv3AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3AdapterTransactorRaw struct {
	Contract *Uniswapv3AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3Adapter creates a new instance of Uniswapv3Adapter, bound to a specific deployed contract.
func NewUniswapv3Adapter(address common.Address, backend bind.ContractBackend) (*Uniswapv3Adapter, error) {
	contract, err := bindUniswapv3Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3Adapter{Uniswapv3AdapterCaller: Uniswapv3AdapterCaller{contract: contract}, Uniswapv3AdapterTransactor: Uniswapv3AdapterTransactor{contract: contract}, Uniswapv3AdapterFilterer: Uniswapv3AdapterFilterer{contract: contract}}, nil
}

// NewUniswapv3AdapterCaller creates a new read-only instance of Uniswapv3Adapter, bound to a specific deployed contract.
func NewUniswapv3AdapterCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv3AdapterCaller, error) {
	contract, err := bindUniswapv3Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3AdapterCaller{contract: contract}, nil
}

// NewUniswapv3AdapterTransactor creates a new write-only instance of Uniswapv3Adapter, bound to a specific deployed contract.
func NewUniswapv3AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3AdapterTransactor, error) {
	contract, err := bindUniswapv3Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3AdapterTransactor{contract: contract}, nil
}

// NewUniswapv3AdapterFilterer creates a new log filterer instance of Uniswapv3Adapter, bound to a specific deployed contract.
func NewUniswapv3AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3AdapterFilterer, error) {
	contract, err := bindUniswapv3Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3AdapterFilterer{contract: contract}, nil
}

// bindUniswapv3Adapter binds a generic wrapper to an already deployed contract.
func bindUniswapv3Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv3AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3Adapter *Uniswapv3AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3Adapter.Contract.Uniswapv3AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3Adapter *Uniswapv3AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.Uniswapv3AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3Adapter *Uniswapv3AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.Uniswapv3AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv3Adapter *Uniswapv3AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Uniswapv3Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) GearboxAdapterType() (uint8, error) {
	return _Uniswapv3Adapter.Contract.GearboxAdapterType(&_Uniswapv3Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _Uniswapv3Adapter.Contract.GearboxAdapterType(&_Uniswapv3Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv3Adapter *Uniswapv3AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Uniswapv3Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _Uniswapv3Adapter.Contract.GearboxAdapterVersion(&_Uniswapv3Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Uniswapv3Adapter.Contract.GearboxAdapterVersion(&_Uniswapv3Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3Adapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) CreditFacade() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.CreditFacade(&_Uniswapv3Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerSession) CreditFacade() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.CreditFacade(&_Uniswapv3Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) CreditManager() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.CreditManager(&_Uniswapv3Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerSession) CreditManager() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.CreditManager(&_Uniswapv3Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) TargetContract() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.TargetContract(&_Uniswapv3Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv3Adapter *Uniswapv3AdapterCallerSession) TargetContract() (common.Address, error) {
	return _Uniswapv3Adapter.Contract.TargetContract(&_Uniswapv3Adapter.CallOpts)
}

// ExactAllInput is a paid mutator transaction binding the contract method 0xf4f18d90.
//
// Solidity: function exactAllInput((bytes,uint256,uint256) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactAllInput(opts *bind.TransactOpts, params IUniswapV3AdapterExactAllInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactAllInput", params)
}

// ExactAllInput is a paid mutator transaction binding the contract method 0xf4f18d90.
//
// Solidity: function exactAllInput((bytes,uint256,uint256) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactAllInput(params IUniswapV3AdapterExactAllInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactAllInput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactAllInput is a paid mutator transaction binding the contract method 0xf4f18d90.
//
// Solidity: function exactAllInput((bytes,uint256,uint256) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactAllInput(params IUniswapV3AdapterExactAllInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactAllInput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactAllInputSingle is a paid mutator transaction binding the contract method 0xc7fbf4de.
//
// Solidity: function exactAllInputSingle((address,address,uint24,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactAllInputSingle(opts *bind.TransactOpts, params IUniswapV3AdapterExactAllInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactAllInputSingle", params)
}

// ExactAllInputSingle is a paid mutator transaction binding the contract method 0xc7fbf4de.
//
// Solidity: function exactAllInputSingle((address,address,uint24,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactAllInputSingle(params IUniswapV3AdapterExactAllInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactAllInputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactAllInputSingle is a paid mutator transaction binding the contract method 0xc7fbf4de.
//
// Solidity: function exactAllInputSingle((address,address,uint24,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactAllInputSingle(params IUniswapV3AdapterExactAllInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactAllInputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactInput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactInput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactInputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactInputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactOutput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactOutput(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactOutputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_Uniswapv3Adapter *Uniswapv3AdapterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _Uniswapv3Adapter.Contract.ExactOutputSingle(&_Uniswapv3Adapter.TransactOpts, params)
}
