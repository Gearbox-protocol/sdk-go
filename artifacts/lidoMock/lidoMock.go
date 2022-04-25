// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lidoMock

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

// LidoMockABI is the input ABI used to generate the binding from.
const LidoMockABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"syncer_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"burnShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTotalShares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalPooledEther\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"syncExchangeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncer\",\"outputs\":[{\"internalType\":\"contractSyncer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPooledEtherSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSharesSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// LidoMock is an auto generated Go binding around an Ethereum contract.
type LidoMock struct {
	LidoMockCaller     // Read-only binding to the contract
	LidoMockTransactor // Write-only binding to the contract
	LidoMockFilterer   // Log filterer for contract events
}

// LidoMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoMockSession struct {
	Contract     *LidoMock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoMockCallerSession struct {
	Contract *LidoMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LidoMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoMockTransactorSession struct {
	Contract     *LidoMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LidoMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoMockRaw struct {
	Contract *LidoMock // Generic contract binding to access the raw methods on
}

// LidoMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoMockCallerRaw struct {
	Contract *LidoMockCaller // Generic read-only contract binding to access the raw methods on
}

// LidoMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoMockTransactorRaw struct {
	Contract *LidoMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidoMock creates a new instance of LidoMock, bound to a specific deployed contract.
func NewLidoMock(address common.Address, backend bind.ContractBackend) (*LidoMock, error) {
	contract, err := bindLidoMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LidoMock{LidoMockCaller: LidoMockCaller{contract: contract}, LidoMockTransactor: LidoMockTransactor{contract: contract}, LidoMockFilterer: LidoMockFilterer{contract: contract}}, nil
}

// NewLidoMockCaller creates a new read-only instance of LidoMock, bound to a specific deployed contract.
func NewLidoMockCaller(address common.Address, caller bind.ContractCaller) (*LidoMockCaller, error) {
	contract, err := bindLidoMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoMockCaller{contract: contract}, nil
}

// NewLidoMockTransactor creates a new write-only instance of LidoMock, bound to a specific deployed contract.
func NewLidoMockTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoMockTransactor, error) {
	contract, err := bindLidoMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoMockTransactor{contract: contract}, nil
}

// NewLidoMockFilterer creates a new log filterer instance of LidoMock, bound to a specific deployed contract.
func NewLidoMockFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoMockFilterer, error) {
	contract, err := bindLidoMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoMockFilterer{contract: contract}, nil
}

// bindLidoMock binds a generic wrapper to an already deployed contract.
func bindLidoMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LidoMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoMock *LidoMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoMock.Contract.LidoMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoMock *LidoMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoMock.Contract.LidoMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoMock *LidoMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoMock.Contract.LidoMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoMock *LidoMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoMock *LidoMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoMock *LidoMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LidoMock *LidoMockCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LidoMock *LidoMockSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LidoMock.Contract.Allowance(&_LidoMock.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LidoMock *LidoMockCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LidoMock.Contract.Allowance(&_LidoMock.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _LidoMock.Contract.BalanceOf(&_LidoMock.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _LidoMock.Contract.BalanceOf(&_LidoMock.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_LidoMock *LidoMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_LidoMock *LidoMockSession) Decimals() (uint8, error) {
	return _LidoMock.Contract.Decimals(&_LidoMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_LidoMock *LidoMockCallerSession) Decimals() (uint8, error) {
	return _LidoMock.Contract.Decimals(&_LidoMock.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_LidoMock *LidoMockCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_LidoMock *LidoMockSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _LidoMock.Contract.GetPooledEthByShares(&_LidoMock.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_LidoMock *LidoMockCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _LidoMock.Contract.GetPooledEthByShares(&_LidoMock.CallOpts, _sharesAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_LidoMock *LidoMockCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_LidoMock *LidoMockSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _LidoMock.Contract.GetSharesByPooledEth(&_LidoMock.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_LidoMock *LidoMockCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _LidoMock.Contract.GetSharesByPooledEth(&_LidoMock.CallOpts, _ethAmount)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_LidoMock *LidoMockCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_LidoMock *LidoMockSession) GetTotalPooledEther() (*big.Int, error) {
	return _LidoMock.Contract.GetTotalPooledEther(&_LidoMock.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_LidoMock *LidoMockCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _LidoMock.Contract.GetTotalPooledEther(&_LidoMock.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_LidoMock *LidoMockCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_LidoMock *LidoMockSession) GetTotalShares() (*big.Int, error) {
	return _LidoMock.Contract.GetTotalShares(&_LidoMock.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_LidoMock *LidoMockCallerSession) GetTotalShares() (*big.Int, error) {
	return _LidoMock.Contract.GetTotalShares(&_LidoMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LidoMock *LidoMockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LidoMock *LidoMockSession) Name() (string, error) {
	return _LidoMock.Contract.Name(&_LidoMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LidoMock *LidoMockCallerSession) Name() (string, error) {
	return _LidoMock.Contract.Name(&_LidoMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LidoMock *LidoMockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LidoMock *LidoMockSession) Owner() (common.Address, error) {
	return _LidoMock.Contract.Owner(&_LidoMock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LidoMock *LidoMockCallerSession) Owner() (common.Address, error) {
	return _LidoMock.Contract.Owner(&_LidoMock.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _LidoMock.Contract.SharesOf(&_LidoMock.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_LidoMock *LidoMockCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _LidoMock.Contract.SharesOf(&_LidoMock.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LidoMock *LidoMockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LidoMock *LidoMockSession) Symbol() (string, error) {
	return _LidoMock.Contract.Symbol(&_LidoMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LidoMock *LidoMockCallerSession) Symbol() (string, error) {
	return _LidoMock.Contract.Symbol(&_LidoMock.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoMock *LidoMockCaller) Syncer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "syncer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoMock *LidoMockSession) Syncer() (common.Address, error) {
	return _LidoMock.Contract.Syncer(&_LidoMock.CallOpts)
}

// Syncer is a free data retrieval call binding the contract method 0xa81cd300.
//
// Solidity: function syncer() view returns(address)
func (_LidoMock *LidoMockCallerSession) Syncer() (common.Address, error) {
	return _LidoMock.Contract.Syncer(&_LidoMock.CallOpts)
}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_LidoMock *LidoMockCaller) TotalPooledEtherSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "totalPooledEtherSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_LidoMock *LidoMockSession) TotalPooledEtherSynced() (*big.Int, error) {
	return _LidoMock.Contract.TotalPooledEtherSynced(&_LidoMock.CallOpts)
}

// TotalPooledEtherSynced is a free data retrieval call binding the contract method 0xbdd17e1c.
//
// Solidity: function totalPooledEtherSynced() view returns(uint256)
func (_LidoMock *LidoMockCallerSession) TotalPooledEtherSynced() (*big.Int, error) {
	return _LidoMock.Contract.TotalPooledEtherSynced(&_LidoMock.CallOpts)
}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_LidoMock *LidoMockCaller) TotalSharesSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "totalSharesSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_LidoMock *LidoMockSession) TotalSharesSynced() (*big.Int, error) {
	return _LidoMock.Contract.TotalSharesSynced(&_LidoMock.CallOpts)
}

// TotalSharesSynced is a free data retrieval call binding the contract method 0xcdb9a21c.
//
// Solidity: function totalSharesSynced() view returns(uint256)
func (_LidoMock *LidoMockCallerSession) TotalSharesSynced() (*big.Int, error) {
	return _LidoMock.Contract.TotalSharesSynced(&_LidoMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoMock *LidoMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoMock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoMock *LidoMockSession) TotalSupply() (*big.Int, error) {
	return _LidoMock.Contract.TotalSupply(&_LidoMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoMock *LidoMockCallerSession) TotalSupply() (*big.Int, error) {
	return _LidoMock.Contract.TotalSupply(&_LidoMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.Approve(&_LidoMock.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.Approve(&_LidoMock.TransactOpts, _spender, _amount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_LidoMock *LidoMockTransactor) BurnShares(opts *bind.TransactOpts, _account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "burnShares", _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_LidoMock *LidoMockSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.BurnShares(&_LidoMock.TransactOpts, _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_LidoMock *LidoMockTransactorSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.BurnShares(&_LidoMock.TransactOpts, _account, _sharesAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_LidoMock *LidoMockTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_LidoMock *LidoMockSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.DecreaseAllowance(&_LidoMock.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_LidoMock *LidoMockTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.DecreaseAllowance(&_LidoMock.TransactOpts, _spender, _subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_LidoMock *LidoMockTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_LidoMock *LidoMockSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.IncreaseAllowance(&_LidoMock.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_LidoMock *LidoMockTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.IncreaseAllowance(&_LidoMock.TransactOpts, _spender, _addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LidoMock *LidoMockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LidoMock *LidoMockSession) RenounceOwnership() (*types.Transaction, error) {
	return _LidoMock.Contract.RenounceOwnership(&_LidoMock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LidoMock *LidoMockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LidoMock.Contract.RenounceOwnership(&_LidoMock.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_LidoMock *LidoMockTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_LidoMock *LidoMockSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _LidoMock.Contract.Submit(&_LidoMock.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_LidoMock *LidoMockTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _LidoMock.Contract.Submit(&_LidoMock.TransactOpts, _referral)
}

// SyncExchangeRate is a paid mutator transaction binding the contract method 0x22301f6e.
//
// Solidity: function syncExchangeRate(uint256 totalPooledEther, uint256 totalShares) returns()
func (_LidoMock *LidoMockTransactor) SyncExchangeRate(opts *bind.TransactOpts, totalPooledEther *big.Int, totalShares *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "syncExchangeRate", totalPooledEther, totalShares)
}

// SyncExchangeRate is a paid mutator transaction binding the contract method 0x22301f6e.
//
// Solidity: function syncExchangeRate(uint256 totalPooledEther, uint256 totalShares) returns()
func (_LidoMock *LidoMockSession) SyncExchangeRate(totalPooledEther *big.Int, totalShares *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.SyncExchangeRate(&_LidoMock.TransactOpts, totalPooledEther, totalShares)
}

// SyncExchangeRate is a paid mutator transaction binding the contract method 0x22301f6e.
//
// Solidity: function syncExchangeRate(uint256 totalPooledEther, uint256 totalShares) returns()
func (_LidoMock *LidoMockTransactorSession) SyncExchangeRate(totalPooledEther *big.Int, totalShares *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.SyncExchangeRate(&_LidoMock.TransactOpts, totalPooledEther, totalShares)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.Transfer(&_LidoMock.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.Transfer(&_LidoMock.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.TransferFrom(&_LidoMock.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_LidoMock *LidoMockTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LidoMock.Contract.TransferFrom(&_LidoMock.TransactOpts, _sender, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LidoMock *LidoMockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LidoMock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LidoMock *LidoMockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LidoMock.Contract.TransferOwnership(&_LidoMock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LidoMock *LidoMockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LidoMock.Contract.TransferOwnership(&_LidoMock.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoMock *LidoMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoMock *LidoMockSession) Receive() (*types.Transaction, error) {
	return _LidoMock.Contract.Receive(&_LidoMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoMock *LidoMockTransactorSession) Receive() (*types.Transaction, error) {
	return _LidoMock.Contract.Receive(&_LidoMock.TransactOpts)
}

// LidoMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LidoMock contract.
type LidoMockApprovalIterator struct {
	Event *LidoMockApproval // Event containing the contract specifics and raw log

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
func (it *LidoMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoMockApproval)
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
		it.Event = new(LidoMockApproval)
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
func (it *LidoMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoMockApproval represents a Approval event raised by the LidoMock contract.
type LidoMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoMock *LidoMockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LidoMockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoMock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LidoMockApprovalIterator{contract: _LidoMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoMock *LidoMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LidoMockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoMock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoMockApproval)
				if err := _LidoMock.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LidoMock *LidoMockFilterer) ParseApproval(log types.Log) (*LidoMockApproval, error) {
	event := new(LidoMockApproval)
	if err := _LidoMock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LidoMockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LidoMock contract.
type LidoMockOwnershipTransferredIterator struct {
	Event *LidoMockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LidoMockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoMockOwnershipTransferred)
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
		it.Event = new(LidoMockOwnershipTransferred)
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
func (it *LidoMockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoMockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoMockOwnershipTransferred represents a OwnershipTransferred event raised by the LidoMock contract.
type LidoMockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LidoMock *LidoMockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LidoMockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LidoMock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LidoMockOwnershipTransferredIterator{contract: _LidoMock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LidoMock *LidoMockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LidoMockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LidoMock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoMockOwnershipTransferred)
				if err := _LidoMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LidoMock *LidoMockFilterer) ParseOwnershipTransferred(log types.Log) (*LidoMockOwnershipTransferred, error) {
	event := new(LidoMockOwnershipTransferred)
	if err := _LidoMock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LidoMockSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the LidoMock contract.
type LidoMockSubmittedIterator struct {
	Event *LidoMockSubmitted // Event containing the contract specifics and raw log

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
func (it *LidoMockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoMockSubmitted)
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
		it.Event = new(LidoMockSubmitted)
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
func (it *LidoMockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoMockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoMockSubmitted represents a Submitted event raised by the LidoMock contract.
type LidoMockSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_LidoMock *LidoMockFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*LidoMockSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoMock.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoMockSubmittedIterator{contract: _LidoMock.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_LidoMock *LidoMockFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *LidoMockSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoMock.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoMockSubmitted)
				if err := _LidoMock.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_LidoMock *LidoMockFilterer) ParseSubmitted(log types.Log) (*LidoMockSubmitted, error) {
	event := new(LidoMockSubmitted)
	if err := _LidoMock.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LidoMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LidoMock contract.
type LidoMockTransferIterator struct {
	Event *LidoMockTransfer // Event containing the contract specifics and raw log

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
func (it *LidoMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoMockTransfer)
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
		it.Event = new(LidoMockTransfer)
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
func (it *LidoMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoMockTransfer represents a Transfer event raised by the LidoMock contract.
type LidoMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoMock *LidoMockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoMockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoMock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoMockTransferIterator{contract: _LidoMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoMock *LidoMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LidoMockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoMock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoMockTransfer)
				if err := _LidoMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LidoMock *LidoMockFilterer) ParseTransfer(log types.Log) (*LidoMockTransfer, error) {
	event := new(LidoMockTransfer)
	if err := _LidoMock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
