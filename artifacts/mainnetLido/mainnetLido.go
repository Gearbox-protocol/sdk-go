// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainnetLido

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

// MainnetLidoABI is the input ABI used to generate the binding from.
const MainnetLidoABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPooledEtherSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSharesSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainnetLido is an auto generated Go binding around an Ethereum contract.
type MainnetLido struct {
	MainnetLidoCaller     // Read-only binding to the contract
	MainnetLidoTransactor // Write-only binding to the contract
	MainnetLidoFilterer   // Log filterer for contract events
}

// MainnetLidoCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnetLidoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnetLidoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnetLidoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnetLidoSession struct {
	Contract     *MainnetLido      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainnetLidoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnetLidoCallerSession struct {
	Contract *MainnetLidoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MainnetLidoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnetLidoTransactorSession struct {
	Contract     *MainnetLidoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MainnetLidoRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnetLidoRaw struct {
	Contract *MainnetLido // Generic contract binding to access the raw methods on
}

// MainnetLidoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnetLidoCallerRaw struct {
	Contract *MainnetLidoCaller // Generic read-only contract binding to access the raw methods on
}

// MainnetLidoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnetLidoTransactorRaw struct {
	Contract *MainnetLidoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnetLido creates a new instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLido(address common.Address, backend bind.ContractBackend) (*MainnetLido, error) {
	contract, err := bindMainnetLido(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainnetLido{MainnetLidoCaller: MainnetLidoCaller{contract: contract}, MainnetLidoTransactor: MainnetLidoTransactor{contract: contract}, MainnetLidoFilterer: MainnetLidoFilterer{contract: contract}}, nil
}

// NewMainnetLidoCaller creates a new read-only instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoCaller(address common.Address, caller bind.ContractCaller) (*MainnetLidoCaller, error) {
	contract, err := bindMainnetLido(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoCaller{contract: contract}, nil
}

// NewMainnetLidoTransactor creates a new write-only instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnetLidoTransactor, error) {
	contract, err := bindMainnetLido(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoTransactor{contract: contract}, nil
}

// NewMainnetLidoFilterer creates a new log filterer instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnetLidoFilterer, error) {
	contract, err := bindMainnetLido(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoFilterer{contract: contract}, nil
}

// bindMainnetLido binds a generic wrapper to an already deployed contract.
func bindMainnetLido(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetLidoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetLido *MainnetLidoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainnetLido.Contract.MainnetLidoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetLido *MainnetLidoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetLido.Contract.MainnetLidoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetLido *MainnetLidoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetLido.Contract.MainnetLidoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetLido *MainnetLidoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainnetLido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetLido *MainnetLidoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetLido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetLido *MainnetLidoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetLido.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MainnetLido *MainnetLidoSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.Allowance(&_MainnetLido.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.Allowance(&_MainnetLido.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.BalanceOf(&_MainnetLido.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.BalanceOf(&_MainnetLido.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MainnetLido *MainnetLidoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MainnetLido *MainnetLidoSession) Decimals() (uint8, error) {
	return _MainnetLido.Contract.Decimals(&_MainnetLido.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MainnetLido *MainnetLidoCallerSession) Decimals() (uint8, error) {
	return _MainnetLido.Contract.Decimals(&_MainnetLido.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _MainnetLido.Contract.GetPooledEthByShares(&_MainnetLido.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _MainnetLido.Contract.GetPooledEthByShares(&_MainnetLido.CallOpts, _sharesAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _MainnetLido.Contract.GetSharesByPooledEth(&_MainnetLido.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _MainnetLido.Contract.GetSharesByPooledEth(&_MainnetLido.CallOpts, _ethAmount)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_MainnetLido *MainnetLidoSession) GetTotalPooledEther() (*big.Int, error) {
	return _MainnetLido.Contract.GetTotalPooledEther(&_MainnetLido.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _MainnetLido.Contract.GetTotalPooledEther(&_MainnetLido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_MainnetLido *MainnetLidoSession) GetTotalShares() (*big.Int, error) {
	return _MainnetLido.Contract.GetTotalShares(&_MainnetLido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) GetTotalShares() (*big.Int, error) {
	return _MainnetLido.Contract.GetTotalShares(&_MainnetLido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MainnetLido *MainnetLidoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MainnetLido *MainnetLidoSession) Name() (string, error) {
	return _MainnetLido.Contract.Name(&_MainnetLido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_MainnetLido *MainnetLidoCallerSession) Name() (string, error) {
	return _MainnetLido.Contract.Name(&_MainnetLido.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.SharesOf(&_MainnetLido.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _MainnetLido.Contract.SharesOf(&_MainnetLido.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MainnetLido *MainnetLidoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MainnetLido *MainnetLidoSession) Symbol() (string, error) {
	return _MainnetLido.Contract.Symbol(&_MainnetLido.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_MainnetLido *MainnetLidoCallerSession) Symbol() (string, error) {
	return _MainnetLido.Contract.Symbol(&_MainnetLido.CallOpts)
}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) TotalPooledEtherSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "totalPooledEtherSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoSession) TotalPooledEtherSynced() (*big.Int, error) {
	return _MainnetLido.Contract.TotalPooledEtherSynced(&_MainnetLido.CallOpts)
}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) TotalPooledEtherSynced() (*big.Int, error) {
	return _MainnetLido.Contract.TotalPooledEtherSynced(&_MainnetLido.CallOpts)
}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) TotalSharesSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "totalSharesSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoSession) TotalSharesSynced() (*big.Int, error) {
	return _MainnetLido.Contract.TotalSharesSynced(&_MainnetLido.CallOpts)
}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) TotalSharesSynced() (*big.Int, error) {
	return _MainnetLido.Contract.TotalSharesSynced(&_MainnetLido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MainnetLido *MainnetLidoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MainnetLido *MainnetLidoSession) TotalSupply() (*big.Int, error) {
	return _MainnetLido.Contract.TotalSupply(&_MainnetLido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MainnetLido *MainnetLidoCallerSession) TotalSupply() (*big.Int, error) {
	return _MainnetLido.Contract.TotalSupply(&_MainnetLido.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.Approve(&_MainnetLido.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.Approve(&_MainnetLido.TransactOpts, _spender, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_MainnetLido *MainnetLidoTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_MainnetLido *MainnetLidoSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.DecreaseAllowance(&_MainnetLido.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_MainnetLido *MainnetLidoTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.DecreaseAllowance(&_MainnetLido.TransactOpts, _spender, _subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_MainnetLido *MainnetLidoTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_MainnetLido *MainnetLidoSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.IncreaseAllowance(&_MainnetLido.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_MainnetLido *MainnetLidoTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.IncreaseAllowance(&_MainnetLido.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.Transfer(&_MainnetLido.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.Transfer(&_MainnetLido.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.TransferFrom(&_MainnetLido.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MainnetLido *MainnetLidoTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MainnetLido.Contract.TransferFrom(&_MainnetLido.TransactOpts, _sender, _recipient, _amount)
}

// MainnetLidoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MainnetLido contract.
type MainnetLidoApprovalIterator struct {
	Event *MainnetLidoApproval // Event containing the contract specifics and raw log

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
func (it *MainnetLidoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetLidoApproval)
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
		it.Event = new(MainnetLidoApproval)
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
func (it *MainnetLidoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetLidoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetLidoApproval represents a Approval event raised by the MainnetLido contract.
type MainnetLidoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MainnetLidoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MainnetLido.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoApprovalIterator{contract: _MainnetLido.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainnetLidoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MainnetLido.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetLidoApproval)
				if err := _MainnetLido.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) ParseApproval(log types.Log) (*MainnetLidoApproval, error) {
	event := new(MainnetLidoApproval)
	if err := _MainnetLido.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MainnetLidoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MainnetLido contract.
type MainnetLidoTransferIterator struct {
	Event *MainnetLidoTransfer // Event containing the contract specifics and raw log

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
func (it *MainnetLidoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetLidoTransfer)
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
		it.Event = new(MainnetLidoTransfer)
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
func (it *MainnetLidoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetLidoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetLidoTransfer represents a Transfer event raised by the MainnetLido contract.
type MainnetLidoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MainnetLidoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MainnetLido.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoTransferIterator{contract: _MainnetLido.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainnetLidoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MainnetLido.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetLidoTransfer)
				if err := _MainnetLido.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) ParseTransfer(log types.Log) (*MainnetLidoTransfer, error) {
	event := new(MainnetLidoTransfer)
	if err := _MainnetLido.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
