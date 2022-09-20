// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lidov1Gateway

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

// Lidov1GatewayMetaData contains all meta data concerning the Lidov1Gateway contract.
var Lidov1GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"contractIstETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Lidov1GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Lidov1GatewayMetaData.ABI instead.
var Lidov1GatewayABI = Lidov1GatewayMetaData.ABI

// Lidov1Gateway is an auto generated Go binding around an Ethereum contract.
type Lidov1Gateway struct {
	Lidov1GatewayCaller     // Read-only binding to the contract
	Lidov1GatewayTransactor // Write-only binding to the contract
	Lidov1GatewayFilterer   // Log filterer for contract events
}

// Lidov1GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Lidov1GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Lidov1GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Lidov1GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lidov1GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Lidov1GatewaySession struct {
	Contract     *Lidov1Gateway    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Lidov1GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Lidov1GatewayCallerSession struct {
	Contract *Lidov1GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Lidov1GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Lidov1GatewayTransactorSession struct {
	Contract     *Lidov1GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Lidov1GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Lidov1GatewayRaw struct {
	Contract *Lidov1Gateway // Generic contract binding to access the raw methods on
}

// Lidov1GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Lidov1GatewayCallerRaw struct {
	Contract *Lidov1GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Lidov1GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Lidov1GatewayTransactorRaw struct {
	Contract *Lidov1GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidov1Gateway creates a new instance of Lidov1Gateway, bound to a specific deployed contract.
func NewLidov1Gateway(address common.Address, backend bind.ContractBackend) (*Lidov1Gateway, error) {
	contract, err := bindLidov1Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lidov1Gateway{Lidov1GatewayCaller: Lidov1GatewayCaller{contract: contract}, Lidov1GatewayTransactor: Lidov1GatewayTransactor{contract: contract}, Lidov1GatewayFilterer: Lidov1GatewayFilterer{contract: contract}}, nil
}

// NewLidov1GatewayCaller creates a new read-only instance of Lidov1Gateway, bound to a specific deployed contract.
func NewLidov1GatewayCaller(address common.Address, caller bind.ContractCaller) (*Lidov1GatewayCaller, error) {
	contract, err := bindLidov1Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Lidov1GatewayCaller{contract: contract}, nil
}

// NewLidov1GatewayTransactor creates a new write-only instance of Lidov1Gateway, bound to a specific deployed contract.
func NewLidov1GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Lidov1GatewayTransactor, error) {
	contract, err := bindLidov1Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Lidov1GatewayTransactor{contract: contract}, nil
}

// NewLidov1GatewayFilterer creates a new log filterer instance of Lidov1Gateway, bound to a specific deployed contract.
func NewLidov1GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Lidov1GatewayFilterer, error) {
	contract, err := bindLidov1Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Lidov1GatewayFilterer{contract: contract}, nil
}

// bindLidov1Gateway binds a generic wrapper to an already deployed contract.
func bindLidov1Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Lidov1GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lidov1Gateway *Lidov1GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lidov1Gateway.Contract.Lidov1GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lidov1Gateway *Lidov1GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Lidov1GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lidov1Gateway *Lidov1GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Lidov1GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lidov1Gateway *Lidov1GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lidov1Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lidov1Gateway *Lidov1GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lidov1Gateway *Lidov1GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.Allowance(&_Lidov1Gateway.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.Allowance(&_Lidov1Gateway.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.BalanceOf(&_Lidov1Gateway.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.BalanceOf(&_Lidov1Gateway.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Gateway *Lidov1GatewayCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Gateway *Lidov1GatewaySession) Decimals() (uint8, error) {
	return _Lidov1Gateway.Contract.Decimals(&_Lidov1Gateway.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) Decimals() (uint8, error) {
	return _Lidov1Gateway.Contract.Decimals(&_Lidov1Gateway.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Gateway *Lidov1GatewayCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Gateway *Lidov1GatewaySession) GetFee() (uint16, error) {
	return _Lidov1Gateway.Contract.GetFee(&_Lidov1Gateway.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) GetFee() (uint16, error) {
	return _Lidov1Gateway.Contract.GetFee(&_Lidov1Gateway.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetPooledEthByShares(&_Lidov1Gateway.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetPooledEthByShares(&_Lidov1Gateway.CallOpts, _sharesAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetSharesByPooledEth(&_Lidov1Gateway.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetSharesByPooledEth(&_Lidov1Gateway.CallOpts, _ethAmount)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) GetTotalPooledEther() (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetTotalPooledEther(&_Lidov1Gateway.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetTotalPooledEther(&_Lidov1Gateway.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) GetTotalShares() (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetTotalShares(&_Lidov1Gateway.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) GetTotalShares() (*big.Int, error) {
	return _Lidov1Gateway.Contract.GetTotalShares(&_Lidov1Gateway.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Gateway *Lidov1GatewayCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Gateway *Lidov1GatewaySession) Name() (string, error) {
	return _Lidov1Gateway.Contract.Name(&_Lidov1Gateway.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) Name() (string, error) {
	return _Lidov1Gateway.Contract.Name(&_Lidov1Gateway.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.SharesOf(&_Lidov1Gateway.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lidov1Gateway.Contract.SharesOf(&_Lidov1Gateway.CallOpts, _account)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Gateway *Lidov1GatewayCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Gateway *Lidov1GatewaySession) StETH() (common.Address, error) {
	return _Lidov1Gateway.Contract.StETH(&_Lidov1Gateway.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) StETH() (common.Address, error) {
	return _Lidov1Gateway.Contract.StETH(&_Lidov1Gateway.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Gateway *Lidov1GatewayCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Gateway *Lidov1GatewaySession) Symbol() (string, error) {
	return _Lidov1Gateway.Contract.Symbol(&_Lidov1Gateway.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) Symbol() (string, error) {
	return _Lidov1Gateway.Contract.Symbol(&_Lidov1Gateway.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewaySession) TotalSupply() (*big.Int, error) {
	return _Lidov1Gateway.Contract.TotalSupply(&_Lidov1Gateway.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) TotalSupply() (*big.Int, error) {
	return _Lidov1Gateway.Contract.TotalSupply(&_Lidov1Gateway.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Gateway *Lidov1GatewayCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lidov1Gateway.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Gateway *Lidov1GatewaySession) Weth() (common.Address, error) {
	return _Lidov1Gateway.Contract.Weth(&_Lidov1Gateway.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Lidov1Gateway *Lidov1GatewayCallerSession) Weth() (common.Address, error) {
	return _Lidov1Gateway.Contract.Weth(&_Lidov1Gateway.CallOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 amount, address _referral) returns(uint256 value)
func (_Lidov1Gateway *Lidov1GatewayTransactor) Submit(opts *bind.TransactOpts, amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _Lidov1Gateway.contract.Transact(opts, "submit", amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 amount, address _referral) returns(uint256 value)
func (_Lidov1Gateway *Lidov1GatewaySession) Submit(amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Submit(&_Lidov1Gateway.TransactOpts, amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 amount, address _referral) returns(uint256 value)
func (_Lidov1Gateway *Lidov1GatewayTransactorSession) Submit(amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Submit(&_Lidov1Gateway.TransactOpts, amount, _referral)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lidov1Gateway *Lidov1GatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lidov1Gateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lidov1Gateway *Lidov1GatewaySession) Receive() (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Receive(&_Lidov1Gateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lidov1Gateway *Lidov1GatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _Lidov1Gateway.Contract.Receive(&_Lidov1Gateway.TransactOpts)
}
