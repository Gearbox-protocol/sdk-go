// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package convexAdapter

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

// IBoosterPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type IBoosterPoolInfo struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}

// ConvexAdapterMetaData contains all meta data concerning the ConvexAdapter contract.
var ConvexAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseRewardPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedPhantomToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotImplementedException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenIsNotAddedToCreditManagerException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveLPtoken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraReward1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraReward2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"extraRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraRewardsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claimExtras\",\"type\":\"bool\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedPhantomToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAllAndUnwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claim\",\"type\":\"bool\"}],\"name\":\"withdrawAndUnwrap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_booster\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnPausableAdminException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_acl\",\"outputs\":[{\"internalType\":\"contractIACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"depositAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockFees\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pidToPhantomToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crvRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stash\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"}],\"internalType\":\"structIBooster.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateStakedPhantomTokensMap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimZap\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"rewardContracts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"extraRewardContracts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenRewardContracts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenRewardTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConvexAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ConvexAdapterMetaData.ABI instead.
var ConvexAdapterABI = ConvexAdapterMetaData.ABI

// ConvexAdapter is an auto generated Go binding around an Ethereum contract.
type ConvexAdapter struct {
	ConvexAdapterCaller     // Read-only binding to the contract
	ConvexAdapterTransactor // Write-only binding to the contract
	ConvexAdapterFilterer   // Log filterer for contract events
}

// ConvexAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvexAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvexAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvexAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvexAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvexAdapterSession struct {
	Contract     *ConvexAdapter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvexAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvexAdapterCallerSession struct {
	Contract *ConvexAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConvexAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvexAdapterTransactorSession struct {
	Contract     *ConvexAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConvexAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvexAdapterRaw struct {
	Contract *ConvexAdapter // Generic contract binding to access the raw methods on
}

// ConvexAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvexAdapterCallerRaw struct {
	Contract *ConvexAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// ConvexAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvexAdapterTransactorRaw struct {
	Contract *ConvexAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvexAdapter creates a new instance of ConvexAdapter, bound to a specific deployed contract.
func NewConvexAdapter(address common.Address, backend bind.ContractBackend) (*ConvexAdapter, error) {
	contract, err := bindConvexAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvexAdapter{ConvexAdapterCaller: ConvexAdapterCaller{contract: contract}, ConvexAdapterTransactor: ConvexAdapterTransactor{contract: contract}, ConvexAdapterFilterer: ConvexAdapterFilterer{contract: contract}}, nil
}

// NewConvexAdapterCaller creates a new read-only instance of ConvexAdapter, bound to a specific deployed contract.
func NewConvexAdapterCaller(address common.Address, caller bind.ContractCaller) (*ConvexAdapterCaller, error) {
	contract, err := bindConvexAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexAdapterCaller{contract: contract}, nil
}

// NewConvexAdapterTransactor creates a new write-only instance of ConvexAdapter, bound to a specific deployed contract.
func NewConvexAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvexAdapterTransactor, error) {
	contract, err := bindConvexAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvexAdapterTransactor{contract: contract}, nil
}

// NewConvexAdapterFilterer creates a new log filterer instance of ConvexAdapter, bound to a specific deployed contract.
func NewConvexAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvexAdapterFilterer, error) {
	contract, err := bindConvexAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvexAdapterFilterer{contract: contract}, nil
}

// bindConvexAdapter binds a generic wrapper to an already deployed contract.
func bindConvexAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvexAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexAdapter *ConvexAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexAdapter.Contract.ConvexAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexAdapter *ConvexAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.ConvexAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexAdapter *ConvexAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.ConvexAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvexAdapter *ConvexAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvexAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvexAdapter *ConvexAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvexAdapter *ConvexAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Acl() (common.Address, error) {
	return _ConvexAdapter.Contract.Acl(&_ConvexAdapter.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xa50cf2c8.
//
// Solidity: function _acl() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Acl() (common.Address, error) {
	return _ConvexAdapter.Contract.Acl(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterType() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterType0(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterType0")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterType0() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType0(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterType0() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType0(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterType1(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterType1")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterType1() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType1(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterType1() (uint8, error) {
	return _ConvexAdapter.Contract.GearboxAdapterType1(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterVersion0(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion0")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterVersion0() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion0(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterVersion0() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion0(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCaller) GearboxAdapterVersion1(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion1")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterSession) GearboxAdapterVersion1() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion1(&_ConvexAdapter.CallOpts)
}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_ConvexAdapter *ConvexAdapterCallerSession) GearboxAdapterVersion1() (uint16, error) {
	return _ConvexAdapter.Contract.GearboxAdapterVersion1(&_ConvexAdapter.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.BalanceOf(&_ConvexAdapter.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.BalanceOf(&_ConvexAdapter.CallOpts, account)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditFacade() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade(&_ConvexAdapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditFacade() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade(&_ConvexAdapter.CallOpts)
}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditFacade0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditFacade0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditFacade0() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade0(&_ConvexAdapter.CallOpts)
}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditFacade0() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade0(&_ConvexAdapter.CallOpts)
}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditFacade1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditFacade1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditFacade1() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade1(&_ConvexAdapter.CallOpts)
}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditFacade1() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditFacade1(&_ConvexAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditManager() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager(&_ConvexAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditManager() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager(&_ConvexAdapter.CallOpts)
}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditManager0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditManager0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditManager0() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager0(&_ConvexAdapter.CallOpts)
}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditManager0() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager0(&_ConvexAdapter.CallOpts)
}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CreditManager1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "creditManager1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CreditManager1() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager1(&_ConvexAdapter.CallOpts)
}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CreditManager1() (common.Address, error) {
	return _ConvexAdapter.Contract.CreditManager1(&_ConvexAdapter.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Crv() (common.Address, error) {
	return _ConvexAdapter.Contract.Crv(&_ConvexAdapter.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Crv() (common.Address, error) {
	return _ConvexAdapter.Contract.Crv(&_ConvexAdapter.CallOpts)
}

// Crv0 is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Crv0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "crv0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv0 is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Crv0() (common.Address, error) {
	return _ConvexAdapter.Contract.Crv0(&_ConvexAdapter.CallOpts)
}

// Crv0 is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Crv0() (common.Address, error) {
	return _ConvexAdapter.Contract.Crv0(&_ConvexAdapter.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) CurrentRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "currentRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) CurrentRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.CurrentRewards(&_ConvexAdapter.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) CurrentRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.CurrentRewards(&_ConvexAdapter.CallOpts)
}

// CurveLPtoken is a free data retrieval call binding the contract method 0x927188d9.
//
// Solidity: function curveLPtoken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) CurveLPtoken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "curveLPtoken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveLPtoken is a free data retrieval call binding the contract method 0x927188d9.
//
// Solidity: function curveLPtoken() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) CurveLPtoken() (common.Address, error) {
	return _ConvexAdapter.Contract.CurveLPtoken(&_ConvexAdapter.CallOpts)
}

// CurveLPtoken is a free data retrieval call binding the contract method 0x927188d9.
//
// Solidity: function curveLPtoken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) CurveLPtoken() (common.Address, error) {
	return _ConvexAdapter.Contract.CurveLPtoken(&_ConvexAdapter.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Cvx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "cvx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Cvx() (common.Address, error) {
	return _ConvexAdapter.Contract.Cvx(&_ConvexAdapter.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Cvx() (common.Address, error) {
	return _ConvexAdapter.Contract.Cvx(&_ConvexAdapter.CallOpts)
}

// Cvx0 is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Cvx0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "cvx0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cvx0 is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Cvx0() (common.Address, error) {
	return _ConvexAdapter.Contract.Cvx0(&_ConvexAdapter.CallOpts)
}

// Cvx0 is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Cvx0() (common.Address, error) {
	return _ConvexAdapter.Contract.Cvx0(&_ConvexAdapter.CallOpts)
}

// Donate is a free data retrieval call binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterCaller) Donate(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "donate", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Donate is a free data retrieval call binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Donate(arg0 *big.Int) (bool, error) {
	return _ConvexAdapter.Contract.Donate(&_ConvexAdapter.CallOpts, arg0)
}

// Donate is a free data retrieval call binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterCallerSession) Donate(arg0 *big.Int) (bool, error) {
	return _ConvexAdapter.Contract.Donate(&_ConvexAdapter.CallOpts, arg0)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) Duration() (*big.Int, error) {
	return _ConvexAdapter.Contract.Duration(&_ConvexAdapter.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) Duration() (*big.Int, error) {
	return _ConvexAdapter.Contract.Duration(&_ConvexAdapter.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) Earned(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.Earned(&_ConvexAdapter.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.Earned(&_ConvexAdapter.CallOpts, account)
}

// ExtraReward1 is a free data retrieval call binding the contract method 0xda5b383f.
//
// Solidity: function extraReward1() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) ExtraReward1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "extraReward1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraReward1 is a free data retrieval call binding the contract method 0xda5b383f.
//
// Solidity: function extraReward1() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) ExtraReward1() (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraReward1(&_ConvexAdapter.CallOpts)
}

// ExtraReward1 is a free data retrieval call binding the contract method 0xda5b383f.
//
// Solidity: function extraReward1() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) ExtraReward1() (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraReward1(&_ConvexAdapter.CallOpts)
}

// ExtraReward2 is a free data retrieval call binding the contract method 0x97c3413b.
//
// Solidity: function extraReward2() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) ExtraReward2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "extraReward2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraReward2 is a free data retrieval call binding the contract method 0x97c3413b.
//
// Solidity: function extraReward2() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) ExtraReward2() (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraReward2(&_ConvexAdapter.CallOpts)
}

// ExtraReward2 is a free data retrieval call binding the contract method 0x97c3413b.
//
// Solidity: function extraReward2() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) ExtraReward2() (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraReward2(&_ConvexAdapter.CallOpts)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 i) view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) ExtraRewards(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "extraRewards", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 i) view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) ExtraRewards(i *big.Int) (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraRewards(&_ConvexAdapter.CallOpts, i)
}

// ExtraRewards is a free data retrieval call binding the contract method 0x40c35446.
//
// Solidity: function extraRewards(uint256 i) view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) ExtraRewards(i *big.Int) (common.Address, error) {
	return _ConvexAdapter.Contract.ExtraRewards(&_ConvexAdapter.CallOpts, i)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) ExtraRewardsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "extraRewardsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) ExtraRewardsLength() (*big.Int, error) {
	return _ConvexAdapter.Contract.ExtraRewardsLength(&_ConvexAdapter.CallOpts)
}

// ExtraRewardsLength is a free data retrieval call binding the contract method 0xd55a23f4.
//
// Solidity: function extraRewardsLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) ExtraRewardsLength() (*big.Int, error) {
	return _ConvexAdapter.Contract.ExtraRewardsLength(&_ConvexAdapter.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) HistoricalRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.HistoricalRewards(&_ConvexAdapter.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) HistoricalRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.HistoricalRewards(&_ConvexAdapter.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _ConvexAdapter.Contract.LastTimeRewardApplicable(&_ConvexAdapter.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _ConvexAdapter.Contract.LastTimeRewardApplicable(&_ConvexAdapter.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) LastUpdateTime() (*big.Int, error) {
	return _ConvexAdapter.Contract.LastUpdateTime(&_ConvexAdapter.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) LastUpdateTime() (*big.Int, error) {
	return _ConvexAdapter.Contract.LastUpdateTime(&_ConvexAdapter.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) LockFees(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "lockFees")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) LockFees() (common.Address, error) {
	return _ConvexAdapter.Contract.LockFees(&_ConvexAdapter.CallOpts)
}

// LockFees is a free data retrieval call binding the contract method 0xab366292.
//
// Solidity: function lockFees() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) LockFees() (common.Address, error) {
	return _ConvexAdapter.Contract.LockFees(&_ConvexAdapter.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) LockRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "lockRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) LockRewards() (common.Address, error) {
	return _ConvexAdapter.Contract.LockRewards(&_ConvexAdapter.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) LockRewards() (common.Address, error) {
	return _ConvexAdapter.Contract.LockRewards(&_ConvexAdapter.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Minter() (common.Address, error) {
	return _ConvexAdapter.Contract.Minter(&_ConvexAdapter.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Minter() (common.Address, error) {
	return _ConvexAdapter.Contract.Minter(&_ConvexAdapter.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) NewRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "newRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) NewRewardRatio() (*big.Int, error) {
	return _ConvexAdapter.Contract.NewRewardRatio(&_ConvexAdapter.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) NewRewardRatio() (*big.Int, error) {
	return _ConvexAdapter.Contract.NewRewardRatio(&_ConvexAdapter.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Operator() (common.Address, error) {
	return _ConvexAdapter.Contract.Operator(&_ConvexAdapter.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Operator() (common.Address, error) {
	return _ConvexAdapter.Contract.Operator(&_ConvexAdapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConvexAdapter *ConvexAdapterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Paused() (bool, error) {
	return _ConvexAdapter.Contract.Paused(&_ConvexAdapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConvexAdapter *ConvexAdapterCallerSession) Paused() (bool, error) {
	return _ConvexAdapter.Contract.Paused(&_ConvexAdapter.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) PeriodFinish() (*big.Int, error) {
	return _ConvexAdapter.Contract.PeriodFinish(&_ConvexAdapter.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) PeriodFinish() (*big.Int, error) {
	return _ConvexAdapter.Contract.PeriodFinish(&_ConvexAdapter.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) Pid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "pid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) Pid() (*big.Int, error) {
	return _ConvexAdapter.Contract.Pid(&_ConvexAdapter.CallOpts)
}

// Pid is a free data retrieval call binding the contract method 0xf1068454.
//
// Solidity: function pid() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) Pid() (*big.Int, error) {
	return _ConvexAdapter.Contract.Pid(&_ConvexAdapter.CallOpts)
}

// PidToPhantomToken is a free data retrieval call binding the contract method 0x251d48c0.
//
// Solidity: function pidToPhantomToken(uint256 ) view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) PidToPhantomToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "pidToPhantomToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PidToPhantomToken is a free data retrieval call binding the contract method 0x251d48c0.
//
// Solidity: function pidToPhantomToken(uint256 ) view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) PidToPhantomToken(arg0 *big.Int) (common.Address, error) {
	return _ConvexAdapter.Contract.PidToPhantomToken(&_ConvexAdapter.CallOpts, arg0)
}

// PidToPhantomToken is a free data retrieval call binding the contract method 0x251d48c0.
//
// Solidity: function pidToPhantomToken(uint256 ) view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) PidToPhantomToken(arg0 *big.Int) (common.Address, error) {
	return _ConvexAdapter.Contract.PidToPhantomToken(&_ConvexAdapter.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 i) view returns((address,address,address,address,address,bool))
func (_ConvexAdapter *ConvexAdapterCaller) PoolInfo(opts *bind.CallOpts, i *big.Int) (IBoosterPoolInfo, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "poolInfo", i)

	if err != nil {
		return *new(IBoosterPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBoosterPoolInfo)).(*IBoosterPoolInfo)

	return out0, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 i) view returns((address,address,address,address,address,bool))
func (_ConvexAdapter *ConvexAdapterSession) PoolInfo(i *big.Int) (IBoosterPoolInfo, error) {
	return _ConvexAdapter.Contract.PoolInfo(&_ConvexAdapter.CallOpts, i)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 i) view returns((address,address,address,address,address,bool))
func (_ConvexAdapter *ConvexAdapterCallerSession) PoolInfo(i *big.Int) (IBoosterPoolInfo, error) {
	return _ConvexAdapter.Contract.PoolInfo(&_ConvexAdapter.CallOpts, i)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) PoolLength() (*big.Int, error) {
	return _ConvexAdapter.Contract.PoolLength(&_ConvexAdapter.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) PoolLength() (*big.Int, error) {
	return _ConvexAdapter.Contract.PoolLength(&_ConvexAdapter.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) QueuedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "queuedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) QueuedRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.QueuedRewards(&_ConvexAdapter.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) QueuedRewards() (*big.Int, error) {
	return _ConvexAdapter.Contract.QueuedRewards(&_ConvexAdapter.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Registry() (common.Address, error) {
	return _ConvexAdapter.Contract.Registry(&_ConvexAdapter.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Registry() (common.Address, error) {
	return _ConvexAdapter.Contract.Registry(&_ConvexAdapter.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) RewardManager() (common.Address, error) {
	return _ConvexAdapter.Contract.RewardManager(&_ConvexAdapter.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) RewardManager() (common.Address, error) {
	return _ConvexAdapter.Contract.RewardManager(&_ConvexAdapter.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) RewardPerToken() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardPerToken(&_ConvexAdapter.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) RewardPerToken() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardPerToken(&_ConvexAdapter.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) RewardPerTokenStored() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardPerTokenStored(&_ConvexAdapter.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardPerTokenStored(&_ConvexAdapter.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) RewardRate() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardRate(&_ConvexAdapter.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) RewardRate() (*big.Int, error) {
	return _ConvexAdapter.Contract.RewardRate(&_ConvexAdapter.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) RewardToken() (common.Address, error) {
	return _ConvexAdapter.Contract.RewardToken(&_ConvexAdapter.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) RewardToken() (common.Address, error) {
	return _ConvexAdapter.Contract.RewardToken(&_ConvexAdapter.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) Rewards(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "rewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) Rewards(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.Rewards(&_ConvexAdapter.CallOpts, account)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) Rewards(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.Rewards(&_ConvexAdapter.CallOpts, account)
}

// StakeFor is a free data retrieval call binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterCaller) StakeFor(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "stakeFor", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakeFor is a free data retrieval call binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) StakeFor(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ConvexAdapter.Contract.StakeFor(&_ConvexAdapter.CallOpts, arg0, arg1)
}

// StakeFor is a free data retrieval call binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address , uint256 ) pure returns(bool)
func (_ConvexAdapter *ConvexAdapterCallerSession) StakeFor(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ConvexAdapter.Contract.StakeFor(&_ConvexAdapter.CallOpts, arg0, arg1)
}

// StakedPhantomToken is a free data retrieval call binding the contract method 0x20b2c151.
//
// Solidity: function stakedPhantomToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) StakedPhantomToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "stakedPhantomToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedPhantomToken is a free data retrieval call binding the contract method 0x20b2c151.
//
// Solidity: function stakedPhantomToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) StakedPhantomToken() (common.Address, error) {
	return _ConvexAdapter.Contract.StakedPhantomToken(&_ConvexAdapter.CallOpts)
}

// StakedPhantomToken is a free data retrieval call binding the contract method 0x20b2c151.
//
// Solidity: function stakedPhantomToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) StakedPhantomToken() (common.Address, error) {
	return _ConvexAdapter.Contract.StakedPhantomToken(&_ConvexAdapter.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "staker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) Staker() (common.Address, error) {
	return _ConvexAdapter.Contract.Staker(&_ConvexAdapter.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) Staker() (common.Address, error) {
	return _ConvexAdapter.Contract.Staker(&_ConvexAdapter.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) StakerRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "stakerRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) StakerRewards() (common.Address, error) {
	return _ConvexAdapter.Contract.StakerRewards(&_ConvexAdapter.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) StakerRewards() (common.Address, error) {
	return _ConvexAdapter.Contract.StakerRewards(&_ConvexAdapter.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) StakingToken() (common.Address, error) {
	return _ConvexAdapter.Contract.StakingToken(&_ConvexAdapter.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) StakingToken() (common.Address, error) {
	return _ConvexAdapter.Contract.StakingToken(&_ConvexAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) TargetContract() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract(&_ConvexAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) TargetContract() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract(&_ConvexAdapter.CallOpts)
}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) TargetContract0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "targetContract0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) TargetContract0() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract0(&_ConvexAdapter.CallOpts)
}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) TargetContract0() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract0(&_ConvexAdapter.CallOpts)
}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCaller) TargetContract1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "targetContract1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterSession) TargetContract1() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract1(&_ConvexAdapter.CallOpts)
}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_ConvexAdapter *ConvexAdapterCallerSession) TargetContract1() (common.Address, error) {
	return _ConvexAdapter.Contract.TargetContract1(&_ConvexAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) TotalSupply() (*big.Int, error) {
	return _ConvexAdapter.Contract.TotalSupply(&_ConvexAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _ConvexAdapter.Contract.TotalSupply(&_ConvexAdapter.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvexAdapter.contract.Call(opts, &out, "userRewardPerTokenPaid", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterSession) UserRewardPerTokenPaid(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.UserRewardPerTokenPaid(&_ConvexAdapter.CallOpts, account)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address account) view returns(uint256)
func (_ConvexAdapter *ConvexAdapterCallerSession) UserRewardPerTokenPaid(account common.Address) (*big.Int, error) {
	return _ConvexAdapter.Contract.UserRewardPerTokenPaid(&_ConvexAdapter.CallOpts, account)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5a7b87f2.
//
// Solidity: function claimRewards(address[] rewardContracts, address[] extraRewardContracts, address[] tokenRewardContracts, address[] tokenRewardTokens, uint256 , uint256 , uint256 , uint256 , uint256 ) returns()
func (_ConvexAdapter *ConvexAdapterTransactor) ClaimRewards(opts *bind.TransactOpts, rewardContracts []common.Address, extraRewardContracts []common.Address, tokenRewardContracts []common.Address, tokenRewardTokens []common.Address, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 *big.Int, arg8 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "claimRewards", rewardContracts, extraRewardContracts, tokenRewardContracts, tokenRewardTokens, arg4, arg5, arg6, arg7, arg8)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5a7b87f2.
//
// Solidity: function claimRewards(address[] rewardContracts, address[] extraRewardContracts, address[] tokenRewardContracts, address[] tokenRewardTokens, uint256 , uint256 , uint256 , uint256 , uint256 ) returns()
func (_ConvexAdapter *ConvexAdapterSession) ClaimRewards(rewardContracts []common.Address, extraRewardContracts []common.Address, tokenRewardContracts []common.Address, tokenRewardTokens []common.Address, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 *big.Int, arg8 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.ClaimRewards(&_ConvexAdapter.TransactOpts, rewardContracts, extraRewardContracts, tokenRewardContracts, tokenRewardTokens, arg4, arg5, arg6, arg7, arg8)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5a7b87f2.
//
// Solidity: function claimRewards(address[] rewardContracts, address[] extraRewardContracts, address[] tokenRewardContracts, address[] tokenRewardTokens, uint256 , uint256 , uint256 , uint256 , uint256 ) returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) ClaimRewards(rewardContracts []common.Address, extraRewardContracts []common.Address, tokenRewardContracts []common.Address, tokenRewardTokens []common.Address, arg4 *big.Int, arg5 *big.Int, arg6 *big.Int, arg7 *big.Int, arg8 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.ClaimRewards(&_ConvexAdapter.TransactOpts, rewardContracts, extraRewardContracts, tokenRewardContracts, tokenRewardTokens, arg4, arg5, arg6, arg7, arg8)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 , bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, arg1 *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "deposit", _pid, arg1, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 , bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Deposit(_pid *big.Int, arg1 *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Deposit(&_ConvexAdapter.TransactOpts, _pid, arg1, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 , bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) Deposit(_pid *big.Int, arg1 *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Deposit(&_ConvexAdapter.TransactOpts, _pid, arg1, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) DepositAll(opts *bind.TransactOpts, _pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "depositAll", _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.DepositAll(&_ConvexAdapter.TransactOpts, _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.DepositAll(&_ConvexAdapter.TransactOpts, _pid, _stake)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) GetReward() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.GetReward(&_ConvexAdapter.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) GetReward() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.GetReward(&_ConvexAdapter.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "getReward0", _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.GetReward0(&_ConvexAdapter.TransactOpts, _account, _claimExtras)
}

// GetReward0 is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _claimExtras) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) GetReward0(_account common.Address, _claimExtras bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.GetReward0(&_ConvexAdapter.TransactOpts, _account, _claimExtras)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConvexAdapter *ConvexAdapterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConvexAdapter *ConvexAdapterSession) Pause() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Pause(&_ConvexAdapter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) Pause() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Pause(&_ConvexAdapter.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) Stake(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "stake", arg0)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Stake(arg0 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Stake(&_ConvexAdapter.TransactOpts, arg0)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) Stake(arg0 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Stake(&_ConvexAdapter.TransactOpts, arg0)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) StakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "stakeAll")
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) StakeAll() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.StakeAll(&_ConvexAdapter.TransactOpts)
}

// StakeAll is a paid mutator transaction binding the contract method 0x8dcb4061.
//
// Solidity: function stakeAll() returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) StakeAll() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.StakeAll(&_ConvexAdapter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConvexAdapter *ConvexAdapterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConvexAdapter *ConvexAdapterSession) Unpause() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Unpause(&_ConvexAdapter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) Unpause() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Unpause(&_ConvexAdapter.TransactOpts)
}

// UpdateStakedPhantomTokensMap is a paid mutator transaction binding the contract method 0x9b51ecd3.
//
// Solidity: function updateStakedPhantomTokensMap() returns()
func (_ConvexAdapter *ConvexAdapterTransactor) UpdateStakedPhantomTokensMap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "updateStakedPhantomTokensMap")
}

// UpdateStakedPhantomTokensMap is a paid mutator transaction binding the contract method 0x9b51ecd3.
//
// Solidity: function updateStakedPhantomTokensMap() returns()
func (_ConvexAdapter *ConvexAdapterSession) UpdateStakedPhantomTokensMap() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.UpdateStakedPhantomTokensMap(&_ConvexAdapter.TransactOpts)
}

// UpdateStakedPhantomTokensMap is a paid mutator transaction binding the contract method 0x9b51ecd3.
//
// Solidity: function updateStakedPhantomTokensMap() returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) UpdateStakedPhantomTokensMap() (*types.Transaction, error) {
	return _ConvexAdapter.Contract.UpdateStakedPhantomTokensMap(&_ConvexAdapter.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdraw", arg0, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Withdraw(arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Withdraw(&_ConvexAdapter.TransactOpts, arg0, claim)
}

// Withdraw is a paid mutator transaction binding the contract method 0x38d07436.
//
// Solidity: function withdraw(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) Withdraw(arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Withdraw(&_ConvexAdapter.TransactOpts, arg0, claim)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) Withdraw0(opts *bind.TransactOpts, _pid *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdraw0", _pid, arg1)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) Withdraw0(_pid *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Withdraw0(&_ConvexAdapter.TransactOpts, _pid, arg1)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 ) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) Withdraw0(_pid *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.Withdraw0(&_ConvexAdapter.TransactOpts, _pid, arg1)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterTransactor) WithdrawAll(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdrawAll", claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAll(&_ConvexAdapter.TransactOpts, claim)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x1c1c6fe5.
//
// Solidity: function withdrawAll(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) WithdrawAll(claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAll(&_ConvexAdapter.TransactOpts, claim)
}

// WithdrawAll0 is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) WithdrawAll0(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdrawAll0", _pid)
}

// WithdrawAll0 is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) WithdrawAll0(_pid *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAll0(&_ConvexAdapter.TransactOpts, _pid)
}

// WithdrawAll0 is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) WithdrawAll0(_pid *big.Int) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAll0(&_ConvexAdapter.TransactOpts, _pid)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterTransactor) WithdrawAllAndUnwrap(opts *bind.TransactOpts, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdrawAllAndUnwrap", claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAllAndUnwrap(&_ConvexAdapter.TransactOpts, claim)
}

// WithdrawAllAndUnwrap is a paid mutator transaction binding the contract method 0x49f039a2.
//
// Solidity: function withdrawAllAndUnwrap(bool claim) returns()
func (_ConvexAdapter *ConvexAdapterTransactorSession) WithdrawAllAndUnwrap(claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAllAndUnwrap(&_ConvexAdapter.TransactOpts, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactor) WithdrawAndUnwrap(opts *bind.TransactOpts, arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.contract.Transact(opts, "withdrawAndUnwrap", arg0, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterSession) WithdrawAndUnwrap(arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAndUnwrap(&_ConvexAdapter.TransactOpts, arg0, claim)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0xc32e7202.
//
// Solidity: function withdrawAndUnwrap(uint256 , bool claim) returns(bool)
func (_ConvexAdapter *ConvexAdapterTransactorSession) WithdrawAndUnwrap(arg0 *big.Int, claim bool) (*types.Transaction, error) {
	return _ConvexAdapter.Contract.WithdrawAndUnwrap(&_ConvexAdapter.TransactOpts, arg0, claim)
}

// ConvexAdapterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ConvexAdapter contract.
type ConvexAdapterPausedIterator struct {
	Event *ConvexAdapterPaused // Event containing the contract specifics and raw log

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
func (it *ConvexAdapterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexAdapterPaused)
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
		it.Event = new(ConvexAdapterPaused)
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
func (it *ConvexAdapterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexAdapterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexAdapterPaused represents a Paused event raised by the ConvexAdapter contract.
type ConvexAdapterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) FilterPaused(opts *bind.FilterOpts) (*ConvexAdapterPausedIterator, error) {

	logs, sub, err := _ConvexAdapter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ConvexAdapterPausedIterator{contract: _ConvexAdapter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ConvexAdapterPaused) (event.Subscription, error) {

	logs, sub, err := _ConvexAdapter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexAdapterPaused)
				if err := _ConvexAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) ParsePaused(log types.Log) (*ConvexAdapterPaused, error) {
	event := new(ConvexAdapterPaused)
	if err := _ConvexAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConvexAdapterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ConvexAdapter contract.
type ConvexAdapterUnpausedIterator struct {
	Event *ConvexAdapterUnpaused // Event containing the contract specifics and raw log

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
func (it *ConvexAdapterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvexAdapterUnpaused)
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
		it.Event = new(ConvexAdapterUnpaused)
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
func (it *ConvexAdapterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvexAdapterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvexAdapterUnpaused represents a Unpaused event raised by the ConvexAdapter contract.
type ConvexAdapterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ConvexAdapterUnpausedIterator, error) {

	logs, sub, err := _ConvexAdapter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ConvexAdapterUnpausedIterator{contract: _ConvexAdapter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ConvexAdapterUnpaused) (event.Subscription, error) {

	logs, sub, err := _ConvexAdapter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvexAdapterUnpaused)
				if err := _ConvexAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConvexAdapter *ConvexAdapterFilterer) ParseUnpaused(log types.Log) (*ConvexAdapterUnpaused, error) {
	event := new(ConvexAdapterUnpaused)
	if err := _ConvexAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
