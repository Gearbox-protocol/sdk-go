// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewardPool

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

// RewardPoolABI is the input ABI used to generate the binding from.
const RewardPoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"deposit_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"op_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"contractIDeposit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queuedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_historicalRewards\",\"type\":\"uint256\"}],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardPool is an auto generated Go binding around an Ethereum contract.
type RewardPool struct {
	RewardPoolCaller     // Read-only binding to the contract
	RewardPoolTransactor // Write-only binding to the contract
	RewardPoolFilterer   // Log filterer for contract events
}

// RewardPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardPoolSession struct {
	Contract     *RewardPool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardPoolCallerSession struct {
	Contract *RewardPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RewardPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardPoolTransactorSession struct {
	Contract     *RewardPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RewardPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardPoolRaw struct {
	Contract *RewardPool // Generic contract binding to access the raw methods on
}

// RewardPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardPoolCallerRaw struct {
	Contract *RewardPoolCaller // Generic read-only contract binding to access the raw methods on
}

// RewardPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardPoolTransactorRaw struct {
	Contract *RewardPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardPool creates a new instance of RewardPool, bound to a specific deployed contract.
func NewRewardPool(address common.Address, backend bind.ContractBackend) (*RewardPool, error) {
	contract, err := bindRewardPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardPool{RewardPoolCaller: RewardPoolCaller{contract: contract}, RewardPoolTransactor: RewardPoolTransactor{contract: contract}, RewardPoolFilterer: RewardPoolFilterer{contract: contract}}, nil
}

// NewRewardPoolCaller creates a new read-only instance of RewardPool, bound to a specific deployed contract.
func NewRewardPoolCaller(address common.Address, caller bind.ContractCaller) (*RewardPoolCaller, error) {
	contract, err := bindRewardPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardPoolCaller{contract: contract}, nil
}

// NewRewardPoolTransactor creates a new write-only instance of RewardPool, bound to a specific deployed contract.
func NewRewardPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardPoolTransactor, error) {
	contract, err := bindRewardPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardPoolTransactor{contract: contract}, nil
}

// NewRewardPoolFilterer creates a new log filterer instance of RewardPool, bound to a specific deployed contract.
func NewRewardPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardPoolFilterer, error) {
	contract, err := bindRewardPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardPoolFilterer{contract: contract}, nil
}

// bindRewardPool binds a generic wrapper to an already deployed contract.
func bindRewardPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardPool *RewardPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardPool.Contract.RewardPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardPool *RewardPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardPool.Contract.RewardPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardPool *RewardPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardPool.Contract.RewardPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardPool *RewardPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardPool *RewardPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardPool *RewardPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardPool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RewardPool *RewardPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RewardPool *RewardPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RewardPool.Contract.BalanceOf(&_RewardPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RewardPool.Contract.BalanceOf(&_RewardPool.CallOpts, account)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_RewardPool *RewardPoolCaller) CurrentRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "currentRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_RewardPool *RewardPoolSession) CurrentRewards() (*big.Int, error) {
	return _RewardPool.Contract.CurrentRewards(&_RewardPool.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) CurrentRewards() (*big.Int, error) {
	return _RewardPool.Contract.CurrentRewards(&_RewardPool.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(address)
func (_RewardPool *RewardPoolCaller) Deposits(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "deposits")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(address)
func (_RewardPool *RewardPoolSession) Deposits() (common.Address, error) {
	return _RewardPool.Contract.Deposits(&_RewardPool.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(address)
func (_RewardPool *RewardPoolCallerSession) Deposits() (common.Address, error) {
	return _RewardPool.Contract.Deposits(&_RewardPool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_RewardPool *RewardPoolCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_RewardPool *RewardPoolSession) Duration() (*big.Int, error) {
	return _RewardPool.Contract.Duration(&_RewardPool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) Duration() (*big.Int, error) {
	return _RewardPool.Contract.Duration(&_RewardPool.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_RewardPool *RewardPoolCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_RewardPool *RewardPoolSession) Earned(account common.Address) (*big.Int, error) {
	return _RewardPool.Contract.Earned(&_RewardPool.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _RewardPool.Contract.Earned(&_RewardPool.CallOpts, account)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_RewardPool *RewardPoolCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_RewardPool *RewardPoolSession) HistoricalRewards() (*big.Int, error) {
	return _RewardPool.Contract.HistoricalRewards(&_RewardPool.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) HistoricalRewards() (*big.Int, error) {
	return _RewardPool.Contract.HistoricalRewards(&_RewardPool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_RewardPool *RewardPoolCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_RewardPool *RewardPoolSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _RewardPool.Contract.LastTimeRewardApplicable(&_RewardPool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _RewardPool.Contract.LastTimeRewardApplicable(&_RewardPool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_RewardPool *RewardPoolCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_RewardPool *RewardPoolSession) LastUpdateTime() (*big.Int, error) {
	return _RewardPool.Contract.LastUpdateTime(&_RewardPool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) LastUpdateTime() (*big.Int, error) {
	return _RewardPool.Contract.LastUpdateTime(&_RewardPool.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_RewardPool *RewardPoolCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_RewardPool *RewardPoolSession) Manager() (common.Address, error) {
	return _RewardPool.Contract.Manager(&_RewardPool.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_RewardPool *RewardPoolCallerSession) Manager() (common.Address, error) {
	return _RewardPool.Contract.Manager(&_RewardPool.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_RewardPool *RewardPoolCaller) NewRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "newRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_RewardPool *RewardPoolSession) NewRewardRatio() (*big.Int, error) {
	return _RewardPool.Contract.NewRewardRatio(&_RewardPool.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) NewRewardRatio() (*big.Int, error) {
	return _RewardPool.Contract.NewRewardRatio(&_RewardPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewardPool *RewardPoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewardPool *RewardPoolSession) Operator() (common.Address, error) {
	return _RewardPool.Contract.Operator(&_RewardPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_RewardPool *RewardPoolCallerSession) Operator() (common.Address, error) {
	return _RewardPool.Contract.Operator(&_RewardPool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_RewardPool *RewardPoolCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_RewardPool *RewardPoolSession) PeriodFinish() (*big.Int, error) {
	return _RewardPool.Contract.PeriodFinish(&_RewardPool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) PeriodFinish() (*big.Int, error) {
	return _RewardPool.Contract.PeriodFinish(&_RewardPool.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_RewardPool *RewardPoolCaller) QueuedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "queuedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_RewardPool *RewardPoolSession) QueuedRewards() (*big.Int, error) {
	return _RewardPool.Contract.QueuedRewards(&_RewardPool.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) QueuedRewards() (*big.Int, error) {
	return _RewardPool.Contract.QueuedRewards(&_RewardPool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_RewardPool *RewardPoolCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_RewardPool *RewardPoolSession) RewardPerToken() (*big.Int, error) {
	return _RewardPool.Contract.RewardPerToken(&_RewardPool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) RewardPerToken() (*big.Int, error) {
	return _RewardPool.Contract.RewardPerToken(&_RewardPool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_RewardPool *RewardPoolCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_RewardPool *RewardPoolSession) RewardPerTokenStored() (*big.Int, error) {
	return _RewardPool.Contract.RewardPerTokenStored(&_RewardPool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _RewardPool.Contract.RewardPerTokenStored(&_RewardPool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_RewardPool *RewardPoolCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_RewardPool *RewardPoolSession) RewardRate() (*big.Int, error) {
	return _RewardPool.Contract.RewardRate(&_RewardPool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) RewardRate() (*big.Int, error) {
	return _RewardPool.Contract.RewardRate(&_RewardPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardPool *RewardPoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardPool *RewardPoolSession) RewardToken() (common.Address, error) {
	return _RewardPool.Contract.RewardToken(&_RewardPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardPool *RewardPoolCallerSession) RewardToken() (common.Address, error) {
	return _RewardPool.Contract.RewardToken(&_RewardPool.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_RewardPool *RewardPoolCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_RewardPool *RewardPoolSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _RewardPool.Contract.Rewards(&_RewardPool.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _RewardPool.Contract.Rewards(&_RewardPool.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardPool *RewardPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardPool *RewardPoolSession) TotalSupply() (*big.Int, error) {
	return _RewardPool.Contract.TotalSupply(&_RewardPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _RewardPool.Contract.TotalSupply(&_RewardPool.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_RewardPool *RewardPoolCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardPool.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_RewardPool *RewardPoolSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _RewardPool.Contract.UserRewardPerTokenPaid(&_RewardPool.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_RewardPool *RewardPoolCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _RewardPool.Contract.UserRewardPerTokenPaid(&_RewardPool.CallOpts, arg0)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_RewardPool *RewardPoolTransactor) Donate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "donate", _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_RewardPool *RewardPoolSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Donate(&_RewardPool.TransactOpts, _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_RewardPool *RewardPoolTransactorSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Donate(&_RewardPool.TransactOpts, _amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_RewardPool *RewardPoolTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_RewardPool *RewardPoolSession) GetReward() (*types.Transaction, error) {
	return _RewardPool.Contract.GetReward(&_RewardPool.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_RewardPool *RewardPoolTransactorSession) GetReward() (*types.Transaction, error) {
	return _RewardPool.Contract.GetReward(&_RewardPool.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_RewardPool *RewardPoolTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "getReward0", _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_RewardPool *RewardPoolSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _RewardPool.Contract.GetReward0(&_RewardPool.TransactOpts, _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_RewardPool *RewardPoolTransactorSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _RewardPool.Contract.GetReward0(&_RewardPool.TransactOpts, _account)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_RewardPool *RewardPoolTransactor) QueueNewRewards(opts *bind.TransactOpts, _rewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "queueNewRewards", _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_RewardPool *RewardPoolSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.QueueNewRewards(&_RewardPool.TransactOpts, _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_RewardPool *RewardPoolTransactorSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.QueueNewRewards(&_RewardPool.TransactOpts, _rewards)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolTransactor) Stake(opts *bind.TransactOpts, _account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "stake", _account, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolSession) Stake(_account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Stake(&_RewardPool.TransactOpts, _account, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolTransactorSession) Stake(_account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Stake(&_RewardPool.TransactOpts, _account, amount)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_RewardPool *RewardPoolTransactor) Sync(opts *bind.TransactOpts, _periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "sync", _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_RewardPool *RewardPoolSession) Sync(_periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Sync(&_RewardPool.TransactOpts, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_RewardPool *RewardPoolTransactorSession) Sync(_periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Sync(&_RewardPool.TransactOpts, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolTransactor) Withdraw(opts *bind.TransactOpts, _account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.contract.Transact(opts, "withdraw", _account, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolSession) Withdraw(_account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Withdraw(&_RewardPool.TransactOpts, _account, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _account, uint256 amount) returns()
func (_RewardPool *RewardPoolTransactorSession) Withdraw(_account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardPool.Contract.Withdraw(&_RewardPool.TransactOpts, _account, amount)
}

// RewardPoolRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the RewardPool contract.
type RewardPoolRewardAddedIterator struct {
	Event *RewardPoolRewardAdded // Event containing the contract specifics and raw log

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
func (it *RewardPoolRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardPoolRewardAdded)
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
		it.Event = new(RewardPoolRewardAdded)
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
func (it *RewardPoolRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardPoolRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardPoolRewardAdded represents a RewardAdded event raised by the RewardPool contract.
type RewardPoolRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_RewardPool *RewardPoolFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*RewardPoolRewardAddedIterator, error) {

	logs, sub, err := _RewardPool.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &RewardPoolRewardAddedIterator{contract: _RewardPool.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_RewardPool *RewardPoolFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *RewardPoolRewardAdded) (event.Subscription, error) {

	logs, sub, err := _RewardPool.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardPoolRewardAdded)
				if err := _RewardPool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_RewardPool *RewardPoolFilterer) ParseRewardAdded(log types.Log) (*RewardPoolRewardAdded, error) {
	event := new(RewardPoolRewardAdded)
	if err := _RewardPool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardPoolRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the RewardPool contract.
type RewardPoolRewardPaidIterator struct {
	Event *RewardPoolRewardPaid // Event containing the contract specifics and raw log

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
func (it *RewardPoolRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardPoolRewardPaid)
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
		it.Event = new(RewardPoolRewardPaid)
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
func (it *RewardPoolRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardPoolRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardPoolRewardPaid represents a RewardPaid event raised by the RewardPool contract.
type RewardPoolRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_RewardPool *RewardPoolFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*RewardPoolRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardPoolRewardPaidIterator{contract: _RewardPool.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_RewardPool *RewardPoolFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *RewardPoolRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardPoolRewardPaid)
				if err := _RewardPool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_RewardPool *RewardPoolFilterer) ParseRewardPaid(log types.Log) (*RewardPoolRewardPaid, error) {
	event := new(RewardPoolRewardPaid)
	if err := _RewardPool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardPoolStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the RewardPool contract.
type RewardPoolStakedIterator struct {
	Event *RewardPoolStaked // Event containing the contract specifics and raw log

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
func (it *RewardPoolStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardPoolStaked)
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
		it.Event = new(RewardPoolStaked)
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
func (it *RewardPoolStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardPoolStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardPoolStaked represents a Staked event raised by the RewardPool contract.
type RewardPoolStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*RewardPoolStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardPoolStakedIterator{contract: _RewardPool.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *RewardPoolStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardPoolStaked)
				if err := _RewardPool.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) ParseStaked(log types.Log) (*RewardPoolStaked, error) {
	event := new(RewardPoolStaked)
	if err := _RewardPool.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RewardPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RewardPool contract.
type RewardPoolWithdrawnIterator struct {
	Event *RewardPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *RewardPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardPoolWithdrawn)
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
		it.Event = new(RewardPoolWithdrawn)
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
func (it *RewardPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardPoolWithdrawn represents a Withdrawn event raised by the RewardPool contract.
type RewardPoolWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*RewardPoolWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardPoolWithdrawnIterator{contract: _RewardPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RewardPoolWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RewardPool.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardPoolWithdrawn)
				if err := _RewardPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_RewardPool *RewardPoolFilterer) ParseWithdrawn(log types.Log) (*RewardPoolWithdrawn, error) {
	event := new(RewardPoolWithdrawn)
	if err := _RewardPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
