// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2Adapter

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

// Uniswapv2AdapterMetaData contains all meta data concerning the Uniswapv2Adapter contract.
var Uniswapv2AdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotImplementedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAllTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Uniswapv2AdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv2AdapterMetaData.ABI instead.
var Uniswapv2AdapterABI = Uniswapv2AdapterMetaData.ABI

// Uniswapv2Adapter is an auto generated Go binding around an Ethereum contract.
type Uniswapv2Adapter struct {
	Uniswapv2AdapterCaller     // Read-only binding to the contract
	Uniswapv2AdapterTransactor // Write-only binding to the contract
	Uniswapv2AdapterFilterer   // Log filterer for contract events
}

// Uniswapv2AdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2AdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2AdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2AdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2AdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2AdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2AdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2AdapterSession struct {
	Contract     *Uniswapv2Adapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2AdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2AdapterCallerSession struct {
	Contract *Uniswapv2AdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Uniswapv2AdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2AdapterTransactorSession struct {
	Contract     *Uniswapv2AdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Uniswapv2AdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2AdapterRaw struct {
	Contract *Uniswapv2Adapter // Generic contract binding to access the raw methods on
}

// Uniswapv2AdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2AdapterCallerRaw struct {
	Contract *Uniswapv2AdapterCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2AdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2AdapterTransactorRaw struct {
	Contract *Uniswapv2AdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2Adapter creates a new instance of Uniswapv2Adapter, bound to a specific deployed contract.
func NewUniswapv2Adapter(address common.Address, backend bind.ContractBackend) (*Uniswapv2Adapter, error) {
	contract, err := bindUniswapv2Adapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2Adapter{Uniswapv2AdapterCaller: Uniswapv2AdapterCaller{contract: contract}, Uniswapv2AdapterTransactor: Uniswapv2AdapterTransactor{contract: contract}, Uniswapv2AdapterFilterer: Uniswapv2AdapterFilterer{contract: contract}}, nil
}

// NewUniswapv2AdapterCaller creates a new read-only instance of Uniswapv2Adapter, bound to a specific deployed contract.
func NewUniswapv2AdapterCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2AdapterCaller, error) {
	contract, err := bindUniswapv2Adapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2AdapterCaller{contract: contract}, nil
}

// NewUniswapv2AdapterTransactor creates a new write-only instance of Uniswapv2Adapter, bound to a specific deployed contract.
func NewUniswapv2AdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2AdapterTransactor, error) {
	contract, err := bindUniswapv2Adapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2AdapterTransactor{contract: contract}, nil
}

// NewUniswapv2AdapterFilterer creates a new log filterer instance of Uniswapv2Adapter, bound to a specific deployed contract.
func NewUniswapv2AdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2AdapterFilterer, error) {
	contract, err := bindUniswapv2Adapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2AdapterFilterer{contract: contract}, nil
}

// bindUniswapv2Adapter binds a generic wrapper to an already deployed contract.
func bindUniswapv2Adapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv2AdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2Adapter *Uniswapv2AdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2Adapter.Contract.Uniswapv2AdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2Adapter *Uniswapv2AdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.Uniswapv2AdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2Adapter *Uniswapv2AdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.Uniswapv2AdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2Adapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) WETH() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.WETH(&_Uniswapv2Adapter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) WETH() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.WETH(&_Uniswapv2Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GearboxAdapterType() (uint8, error) {
	return _Uniswapv2Adapter.Contract.GearboxAdapterType(&_Uniswapv2Adapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _Uniswapv2Adapter.Contract.GearboxAdapterType(&_Uniswapv2Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _Uniswapv2Adapter.Contract.GearboxAdapterVersion(&_Uniswapv2Adapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Uniswapv2Adapter.Contract.GearboxAdapterVersion(&_Uniswapv2Adapter.CallOpts)
}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) AddLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "addLiquidity", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) AddLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.AddLiquidity(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// AddLiquidity is a free data retrieval call binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address , address , uint256 , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) AddLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 *big.Int, arg6 common.Address, arg7 *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.AddLiquidity(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) CreditFacade() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.CreditFacade(&_Uniswapv2Adapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) CreditFacade() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.CreditFacade(&_Uniswapv2Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) CreditManager() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.CreditManager(&_Uniswapv2Adapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) CreditManager() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.CreditManager(&_Uniswapv2Adapter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) Factory() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.Factory(&_Uniswapv2Adapter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) Factory() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.Factory(&_Uniswapv2Adapter.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountIn(&_Uniswapv2Adapter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountIn)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountIn(&_Uniswapv2Adapter.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountOut(&_Uniswapv2Adapter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) view returns(uint256 amountOut)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountOut(&_Uniswapv2Adapter.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountsIn(&_Uniswapv2Adapter.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountsIn(&_Uniswapv2Adapter.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountsOut(&_Uniswapv2Adapter.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.GetAmountsOut(&_Uniswapv2Adapter.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.Quote(&_Uniswapv2Adapter.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) view returns(uint256 amountB)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.Quote(&_Uniswapv2Adapter.CallOpts, amountA, reserveA, reserveB)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidity(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidity", arg0, arg1, arg2, arg3, arg4, arg5, arg6)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidity(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RemoveLiquidity is a free data retrieval call binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address , address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidity(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidity(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidityETH(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidityETH", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETH is a free data retrieval call binding the contract method 0x02751cec.
//
// Solidity: function removeLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidityETHSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidityETHWithPermit(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidityETHWithPermit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidityETHWithPermit(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHWithPermit(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermit is a free data retrieval call binding the contract method 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidityETHWithPermit(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHWithPermit(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int, arg6 bool, arg7 uint8, arg8 [32]byte, arg9 [32]byte) (*big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) RemoveLiquidityWithPermit(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "removeLiquidityWithPermit", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) RemoveLiquidityWithPermit(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityWithPermit(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// RemoveLiquidityWithPermit is a free data retrieval call binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address , address , uint256 , uint256 , uint256 , address , uint256 , bool , uint8 , bytes32 , bytes32 ) pure returns(uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) RemoveLiquidityWithPermit(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, arg5 common.Address, arg6 *big.Int, arg7 bool, arg8 uint8, arg9 [32]byte, arg10 [32]byte) (*big.Int, *big.Int, error) {
	return _Uniswapv2Adapter.Contract.RemoveLiquidityWithPermit(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) SwapExactTokensForETH(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "swapExactTokensForETH", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactTokensForETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETH is a free data retrieval call binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) SwapExactTokensForETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "swapExactTokensForETHSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "swapExactTokensForTokensSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a free data retrieval call binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 , uint256 , address[] , address , uint256 ) pure returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) error {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) SwapTokensForExactETH(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "swapTokensForExactETH", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapTokensForExactETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.SwapTokensForExactETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapTokensForExactETH is a free data retrieval call binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 , uint256 , address[] , address , uint256 ) pure returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) SwapTokensForExactETH(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) ([]*big.Int, error) {
	return _Uniswapv2Adapter.Contract.SwapTokensForExactETH(&_Uniswapv2Adapter.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2Adapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) TargetContract() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.TargetContract(&_Uniswapv2Adapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Uniswapv2Adapter *Uniswapv2AdapterCallerSession) TargetContract() (common.Address, error) {
	return _Uniswapv2Adapter.Contract.TargetContract(&_Uniswapv2Adapter.CallOpts)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) AddLiquidityETH(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "addLiquidityETH", arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) AddLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.AddLiquidityETH(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xf305d719.
//
// Solidity: function addLiquidityETH(address , uint256 , uint256 , uint256 , address , uint256 ) payable returns(uint256, uint256, uint256)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) AddLiquidityETH(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address, arg5 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.AddLiquidityETH(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SwapAllTokensForTokens is a paid mutator transaction binding the contract method 0xbdbeaa31.
//
// Solidity: function swapAllTokensForTokens(uint256 rateMinRAY, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapAllTokensForTokens(opts *bind.TransactOpts, rateMinRAY *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapAllTokensForTokens", rateMinRAY, path, deadline)
}

// SwapAllTokensForTokens is a paid mutator transaction binding the contract method 0xbdbeaa31.
//
// Solidity: function swapAllTokensForTokens(uint256 rateMinRAY, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapAllTokensForTokens(rateMinRAY *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapAllTokensForTokens(&_Uniswapv2Adapter.TransactOpts, rateMinRAY, path, deadline)
}

// SwapAllTokensForTokens is a paid mutator transaction binding the contract method 0xbdbeaa31.
//
// Solidity: function swapAllTokensForTokens(uint256 rateMinRAY, address[] path, uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapAllTokensForTokens(rateMinRAY *big.Int, path []common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapAllTokensForTokens(&_Uniswapv2Adapter.TransactOpts, rateMinRAY, path, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapETHForExactTokens", arg0, arg1, arg2, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapETHForExactTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapETHForExactTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapETHForExactTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapETHForExactTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapExactETHForTokens", arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactETHForTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactETHForTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 , address[] , address , uint256 ) payable returns(uint256[])
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapExactETHForTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactETHForTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 , address[] , address , uint256 ) payable returns()
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(arg0 *big.Int, arg1 []common.Address, arg2 common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_Uniswapv2Adapter.TransactOpts, arg0, arg1, arg2, arg3)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, arg3, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForTokens(&_Uniswapv2Adapter.TransactOpts, amountIn, amountOutMin, path, arg3, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapExactTokensForTokens(&_Uniswapv2Adapter.TransactOpts, amountIn, amountOutMin, path, arg3, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, arg3, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapTokensForExactTokens(&_Uniswapv2Adapter.TransactOpts, amountOut, amountInMax, path, arg3, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address , uint256 deadline) returns(uint256[] amounts)
func (_Uniswapv2Adapter *Uniswapv2AdapterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, arg3 common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswapv2Adapter.Contract.SwapTokensForExactTokens(&_Uniswapv2Adapter.TransactOpts, amountOut, amountInMax, path, arg3, deadline)
}
