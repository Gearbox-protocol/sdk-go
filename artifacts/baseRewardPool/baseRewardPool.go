// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baseRewardPool

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

// BaseRewardPoolMetaData contains all meta data concerning the BaseRewardPool contract.
var BaseRewardPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reward\",\"type\":\"address\"}],\"name\":\"addExtraReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearExtraRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extraRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraRewardsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimExtras\",\"type\":\"bool\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queuedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_historicalRewards\",\"type\":\"uint256\"}],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAllAndUnwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAndUnwrap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BaseRewardPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseRewardPoolMetaData.ABI instead.
var BaseRewardPoolABI = BaseRewardPoolMetaData.ABI

// BaseRewardPool is an auto generated Go binding around an Ethereum contract.
type BaseRewardPool struct {
	BaseRewardPoolCaller     // Read-only binding to the contract
	BaseRewardPoolTransactor // Write-only binding to the contract
	BaseRewardPoolFilterer   // Log filterer for contract events
}

// BaseRewardPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseRewardPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRewardPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseRewardPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRewardPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseRewardPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRewardPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseRewardPoolSession struct {
	Contract     *BaseRewardPool   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRewardPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseRewardPoolCallerSession struct {
	Contract *BaseRewardPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BaseRewardPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseRewardPoolTransactorSession struct {
	Contract     *BaseRewardPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BaseRewardPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRewardPoolRaw struct {
	Contract *BaseRewardPool // Generic contract binding to access the raw methods on
}

// BaseRewardPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseRewardPoolCallerRaw struct {
	Contract *BaseRewardPoolCaller // Generic read-only contract binding to access the raw methods on
}

// BaseRewardPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseRewardPoolTransactorRaw struct {
	Contract *BaseRewardPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseRewardPool creates a new instance of BaseRewardPool, bound to a specific deployed contract.
func NewBaseRewardPool(address common.Address, backend bind.ContractBackend) (*BaseRewardPool, error) {
	contract, err := bindBaseRewardPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPool{BaseRewardPoolCaller: BaseRewardPoolCaller{contract: contract}, BaseRewardPoolTransactor: BaseRewardPoolTransactor{contract: contract}, BaseRewardPoolFilterer: BaseRewardPoolFilterer{contract: contract}}, nil
}

// NewBaseRewardPoolCaller creates a new read-only instance of BaseRewardPool, bound to a specific deployed contract.
func NewBaseRewardPoolCaller(address common.Address, caller bind.ContractCaller) (*BaseRewardPoolCaller, error) {
	contract, err := bindBaseRewardPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolCaller{contract: contract}, nil
}

// NewBaseRewardPoolTransactor creates a new write-only instance of BaseRewardPool, bound to a specific deployed contract.
func NewBaseRewardPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseRewardPoolTransactor, error) {
	contract, err := bindBaseRewardPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolTransactor{contract: contract}, nil
}

// NewBaseRewardPoolFilterer creates a new log filterer instance of BaseRewardPool, bound to a specific deployed contract.
func NewBaseRewardPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseRewardPoolFilterer, error) {
	contract, err := bindBaseRewardPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolFilterer{contract: contract}, nil
}

// bindBaseRewardPool binds a generic wrapper to an already deployed contract.
func bindBaseRewardPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseRewardPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRewardPool *BaseRewardPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRewardPool.Contract.BaseRewardPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRewardPool *BaseRewardPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.BaseRewardPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRewardPool *BaseRewardPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.BaseRewardPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRewardPool *BaseRewardPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRewardPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRewardPool *BaseRewardPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRewardPool *BaseRewardPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.BalanceOf(&_BaseRewardPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.BalanceOf(&_BaseRewardPool.CallOpts, account)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) CurrentRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "currentRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) CurrentRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.CurrentRewards(&_BaseRewardPool.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) CurrentRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.CurrentRewards(&_BaseRewardPool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) Duration() (*big.Int, error) {
	return _BaseRewardPool.Contract.Duration(&_BaseRewardPool.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) Duration() (*big.Int, error) {
	return _BaseRewardPool.Contract.Duration(&_BaseRewardPool.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) Earned(account common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.Earned(&_BaseRewardPool.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.Earned(&_BaseRewardPool.CallOpts, account)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_BaseRewardPool *BaseRewardPoolCaller) ExtraRewards(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "extraRewards", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_BaseRewardPool *BaseRewardPoolSession) ExtraRewards(arg0 *big.Int) (common.Address, error) {
	return _BaseRewardPool.Contract.ExtraRewards(&_BaseRewardPool.CallOpts, arg0)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 ) view returns(address)
func (_BaseRewardPool *BaseRewardPoolCallerSession) ExtraRewards(arg0 *big.Int) (common.Address, error) {
	return _BaseRewardPool.Contract.ExtraRewards(&_BaseRewardPool.CallOpts, arg0)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) ExtraRewardsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "extraRewardsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) ExtraRewardsLength() (*big.Int, error) {
	return _BaseRewardPool.Contract.ExtraRewardsLength(&_BaseRewardPool.CallOpts)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) ExtraRewardsLength() (*big.Int, error) {
	return _BaseRewardPool.Contract.ExtraRewardsLength(&_BaseRewardPool.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) HistoricalRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.HistoricalRewards(&_BaseRewardPool.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) HistoricalRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.HistoricalRewards(&_BaseRewardPool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _BaseRewardPool.Contract.LastTimeRewardApplicable(&_BaseRewardPool.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _BaseRewardPool.Contract.LastTimeRewardApplicable(&_BaseRewardPool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) LastUpdateTime() (*big.Int, error) {
	return _BaseRewardPool.Contract.LastUpdateTime(&_BaseRewardPool.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) LastUpdateTime() (*big.Int, error) {
	return _BaseRewardPool.Contract.LastUpdateTime(&_BaseRewardPool.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) NewRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "newRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) NewRewardRatio() (*big.Int, error) {
	return _BaseRewardPool.Contract.NewRewardRatio(&_BaseRewardPool.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) NewRewardRatio() (*big.Int, error) {
	return _BaseRewardPool.Contract.NewRewardRatio(&_BaseRewardPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BaseRewardPool *BaseRewardPoolSession) Operator() (common.Address, error) {
	return _BaseRewardPool.Contract.Operator(&_BaseRewardPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCallerSession) Operator() (common.Address, error) {
	return _BaseRewardPool.Contract.Operator(&_BaseRewardPool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) PeriodFinish() (*big.Int, error) {
	return _BaseRewardPool.Contract.PeriodFinish(&_BaseRewardPool.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) PeriodFinish() (*big.Int, error) {
	return _BaseRewardPool.Contract.PeriodFinish(&_BaseRewardPool.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) Pid() (*big.Int, error) {
	return _BaseRewardPool.Contract.Pid(&_BaseRewardPool.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) Pid() (*big.Int, error) {
	return _BaseRewardPool.Contract.Pid(&_BaseRewardPool.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) QueuedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "queuedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) QueuedRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.QueuedRewards(&_BaseRewardPool.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) QueuedRewards() (*big.Int, error) {
	return _BaseRewardPool.Contract.QueuedRewards(&_BaseRewardPool.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BaseRewardPool *BaseRewardPoolSession) RewardManager() (common.Address, error) {
	return _BaseRewardPool.Contract.RewardManager(&_BaseRewardPool.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCallerSession) RewardManager() (common.Address, error) {
	return _BaseRewardPool.Contract.RewardManager(&_BaseRewardPool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) RewardPerToken() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardPerToken(&_BaseRewardPool.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) RewardPerToken() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardPerToken(&_BaseRewardPool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) RewardPerTokenStored() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardPerTokenStored(&_BaseRewardPool.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardPerTokenStored(&_BaseRewardPool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) RewardRate() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardRate(&_BaseRewardPool.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) RewardRate() (*big.Int, error) {
	return _BaseRewardPool.Contract.RewardRate(&_BaseRewardPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolSession) RewardToken() (common.Address, error) {
	return _BaseRewardPool.Contract.RewardToken(&_BaseRewardPool.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCallerSession) RewardToken() (common.Address, error) {
	return _BaseRewardPool.Contract.RewardToken(&_BaseRewardPool.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.Rewards(&_BaseRewardPool.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.Rewards(&_BaseRewardPool.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolSession) StakingToken() (common.Address, error) {
	return _BaseRewardPool.Contract.StakingToken(&_BaseRewardPool.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_BaseRewardPool *BaseRewardPoolCallerSession) StakingToken() (common.Address, error) {
	return _BaseRewardPool.Contract.StakingToken(&_BaseRewardPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) TotalSupply() (*big.Int, error) {
	return _BaseRewardPool.Contract.TotalSupply(&_BaseRewardPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BaseRewardPool.Contract.TotalSupply(&_BaseRewardPool.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseRewardPool.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.UserRewardPerTokenPaid(&_BaseRewardPool.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_BaseRewardPool *BaseRewardPoolCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _BaseRewardPool.Contract.UserRewardPerTokenPaid(&_BaseRewardPool.CallOpts, arg0)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) AddExtraReward(opts *bind.TransactOpts, _reward common.Address) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "addExtraReward", _reward)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) AddExtraReward(_reward common.Address) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.AddExtraReward(&_BaseRewardPool.TransactOpts, _reward)
}

// AddExtraReward is a paid mutator transaction binding the contract method 0x5e43c47b.
//
// Solidity: function addExtraReward(address _reward) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) AddExtraReward(_reward common.Address) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.AddExtraReward(&_BaseRewardPool.TransactOpts, _reward)
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_BaseRewardPool *BaseRewardPoolTransactor) ClearExtraRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "clearExtraRewards")
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_BaseRewardPool *BaseRewardPoolSession) ClearExtraRewards() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.ClearExtraRewards(&_BaseRewardPool.TransactOpts)
}

// ClearExtraRewards is a paid mutator transaction binding the contract method 0x0569d388.
//
// Solidity: function clearExtraRewards() returns()
func (_BaseRewardPool *BaseRewardPoolTransactorSession) ClearExtraRewards() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.ClearExtraRewards(&_BaseRewardPool.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) Donate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "donate", _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Donate(&_BaseRewardPool.TransactOpts, _amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) Donate(_amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Donate(&_BaseRewardPool.TransactOpts, _amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) GetReward() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.GetReward(&_BaseRewardPool.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) GetReward() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.GetReward(&_BaseRewardPool.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "getReward0", _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.GetReward0(&_BaseRewardPool.TransactOpts, _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.GetReward0(&_BaseRewardPool.TransactOpts, _account, _claimExtras)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) QueueNewRewards(opts *bind.TransactOpts, _rewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "queueNewRewards", _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.QueueNewRewards(&_BaseRewardPool.TransactOpts, _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _rewards) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) QueueNewRewards(_rewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.QueueNewRewards(&_BaseRewardPool.TransactOpts, _rewards)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Stake(&_BaseRewardPool.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Stake(&_BaseRewardPool.TransactOpts, _amount)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) StakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "stakeAll")
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) StakeAll() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.StakeAll(&_BaseRewardPool.TransactOpts)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) StakeAll() (*types.Transaction, error) {
	return _BaseRewardPool.Contract.StakeAll(&_BaseRewardPool.TransactOpts)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) StakeFor(opts *bind.TransactOpts, _for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "stakeFor", _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.StakeFor(&_BaseRewardPool.TransactOpts, _for, _amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address _for, uint256 _amount) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) StakeFor(_for common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.StakeFor(&_BaseRewardPool.TransactOpts, _for, _amount)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_BaseRewardPool *BaseRewardPoolTransactor) Sync(opts *bind.TransactOpts, _periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "sync", _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_BaseRewardPool *BaseRewardPoolSession) Sync(_periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Sync(&_BaseRewardPool.TransactOpts, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Sync is a paid mutator transaction binding the contract method 0x39782886.
//
// Solidity: function sync(uint256 _periodFinish, uint256 _rewardRate, uint256 _lastUpdateTime, uint256 _rewardPerTokenStored, uint256 _queuedRewards, uint256 _currentRewards, uint256 _historicalRewards) returns()
func (_BaseRewardPool *BaseRewardPoolTransactorSession) Sync(_periodFinish *big.Int, _rewardRate *big.Int, _lastUpdateTime *big.Int, _rewardPerTokenStored *big.Int, _queuedRewards *big.Int, _currentRewards *big.Int, _historicalRewards *big.Int) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Sync(&_BaseRewardPool.TransactOpts, _periodFinish, _rewardRate, _lastUpdateTime, _rewardPerTokenStored, _queuedRewards, _currentRewards, _historicalRewards)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "withdraw", amount, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) Withdraw(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Withdraw(&_BaseRewardPool.TransactOpts, amount, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) Withdraw(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.Withdraw(&_BaseRewardPool.TransactOpts, amount, claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolTransactor) WithdrawAll(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "withdrawAll", claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAll(&_BaseRewardPool.TransactOpts, claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolTransactorSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAll(&_BaseRewardPool.TransactOpts, claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolTransactor) WithdrawAllAndUnwrap(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "withdrawAllAndUnwrap", claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAllAndUnwrap(&_BaseRewardPool.TransactOpts, claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_BaseRewardPool *BaseRewardPoolTransactorSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAllAndUnwrap(&_BaseRewardPool.TransactOpts, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactor) WithdrawAndUnwrap(opts *bind.TransactOpts, amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.contract.Transact(opts, "withdrawAndUnwrap", amount, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolSession) WithdrawAndUnwrap(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAndUnwrap(&_BaseRewardPool.TransactOpts, amount, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 amount, bool claim) returns(bool)
func (_BaseRewardPool *BaseRewardPoolTransactorSession) WithdrawAndUnwrap(amount *big.Int, claim bool) (*types.Transaction, error) {
	return _BaseRewardPool.Contract.WithdrawAndUnwrap(&_BaseRewardPool.TransactOpts, amount, claim)
}

// BaseRewardPoolRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the BaseRewardPool contract.
type BaseRewardPoolRewardAddedIterator struct {
	Event *BaseRewardPoolRewardAdded // Event containing the contract specifics and raw log

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
func (it *BaseRewardPoolRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseRewardPoolRewardAdded)
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
		it.Event = new(BaseRewardPoolRewardAdded)
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
func (it *BaseRewardPoolRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseRewardPoolRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseRewardPoolRewardAdded represents a RewardAdded event raised by the BaseRewardPool contract.
type BaseRewardPoolRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_BaseRewardPool *BaseRewardPoolFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*BaseRewardPoolRewardAddedIterator, error) {

	logs, sub, err := _BaseRewardPool.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolRewardAddedIterator{contract: _BaseRewardPool.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_BaseRewardPool *BaseRewardPoolFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *BaseRewardPoolRewardAdded) (event.Subscription, error) {

	logs, sub, err := _BaseRewardPool.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseRewardPoolRewardAdded)
				if err := _BaseRewardPool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_BaseRewardPool *BaseRewardPoolFilterer) ParseRewardAdded(log types.Log) (*BaseRewardPoolRewardAdded, error) {
	event := new(BaseRewardPoolRewardAdded)
	if err := _BaseRewardPool.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseRewardPoolRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the BaseRewardPool contract.
type BaseRewardPoolRewardPaidIterator struct {
	Event *BaseRewardPoolRewardPaid // Event containing the contract specifics and raw log

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
func (it *BaseRewardPoolRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseRewardPoolRewardPaid)
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
		it.Event = new(BaseRewardPoolRewardPaid)
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
func (it *BaseRewardPoolRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseRewardPoolRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseRewardPoolRewardPaid represents a RewardPaid event raised by the BaseRewardPool contract.
type BaseRewardPoolRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_BaseRewardPool *BaseRewardPoolFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*BaseRewardPoolRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolRewardPaidIterator{contract: _BaseRewardPool.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_BaseRewardPool *BaseRewardPoolFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *BaseRewardPoolRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseRewardPoolRewardPaid)
				if err := _BaseRewardPool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_BaseRewardPool *BaseRewardPoolFilterer) ParseRewardPaid(log types.Log) (*BaseRewardPoolRewardPaid, error) {
	event := new(BaseRewardPoolRewardPaid)
	if err := _BaseRewardPool.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseRewardPoolStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the BaseRewardPool contract.
type BaseRewardPoolStakedIterator struct {
	Event *BaseRewardPoolStaked // Event containing the contract specifics and raw log

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
func (it *BaseRewardPoolStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseRewardPoolStaked)
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
		it.Event = new(BaseRewardPoolStaked)
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
func (it *BaseRewardPoolStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseRewardPoolStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseRewardPoolStaked represents a Staked event raised by the BaseRewardPool contract.
type BaseRewardPoolStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_BaseRewardPool *BaseRewardPoolFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*BaseRewardPoolStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolStakedIterator{contract: _BaseRewardPool.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_BaseRewardPool *BaseRewardPoolFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *BaseRewardPoolStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseRewardPoolStaked)
				if err := _BaseRewardPool.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_BaseRewardPool *BaseRewardPoolFilterer) ParseStaked(log types.Log) (*BaseRewardPoolStaked, error) {
	event := new(BaseRewardPoolStaked)
	if err := _BaseRewardPool.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseRewardPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BaseRewardPool contract.
type BaseRewardPoolWithdrawnIterator struct {
	Event *BaseRewardPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *BaseRewardPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseRewardPoolWithdrawn)
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
		it.Event = new(BaseRewardPoolWithdrawn)
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
func (it *BaseRewardPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseRewardPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseRewardPoolWithdrawn represents a Withdrawn event raised by the BaseRewardPool contract.
type BaseRewardPoolWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_BaseRewardPool *BaseRewardPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*BaseRewardPoolWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &BaseRewardPoolWithdrawnIterator{contract: _BaseRewardPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_BaseRewardPool *BaseRewardPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BaseRewardPoolWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BaseRewardPool.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseRewardPoolWithdrawn)
				if err := _BaseRewardPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_BaseRewardPool *BaseRewardPoolFilterer) ParseWithdrawn(log types.Log) (*BaseRewardPoolWithdrawn, error) {
	event := new(BaseRewardPoolWithdrawn)
	if err := _BaseRewardPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
