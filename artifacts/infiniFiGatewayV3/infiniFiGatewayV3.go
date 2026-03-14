// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package infiniFiGatewayV3

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

// AllocationVotingAllocationVote is an auto generated low-level Go binding around an user-defined struct.
type AllocationVotingAllocationVote struct {
	Farm   common.Address
	Weight *big.Int
}

// CoreControlledCall is an auto generated low-level Go binding around an user-defined struct.
type CoreControlledCall struct {
	Target   common.Address
	Value    *big.Int
	CallData []byte
}

// InfiniFiGatewayV3MetaData contains all meta data concerning the InfiniFiGatewayV3 contract.
var InfiniFiGatewayV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZapFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZapRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"MinAssetsOutError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingLossesUnapplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"UnderlyingCallReverted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldCore\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCore\",\"type\":\"address\"}],\"name\":\"CoreUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SetEnabledRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zapFee\",\"type\":\"uint256\"}],\"name\":\"ZapFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiptTokens\",\"type\":\"uint256\"}],\"name\":\"ZapIn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unwindingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_newUnwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"cancelInfiWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unwindingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_newUnwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"cancelUnwinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRedemption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unwindingTimestamp\",\"type\":\"uint256\"}],\"name\":\"completeInfiWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"internalType\":\"contractInfiniFiCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"createPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCoreControlled.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"emergencyAction\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabledRouters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_oldUnwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_newUnwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"increaseInfiWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_oldUnwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_newUnwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"increaseUnwindingEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_core\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"lockInfi\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_minIusdReceived\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"mintAndLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintAndStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_assets\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"farm\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"weight\",\"type\":\"uint96\"}],\"internalType\":\"structAllocationVoting.AllocationVote[][]\",\"name\":\"_liquidVotes\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"farm\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"weight\",\"type\":\"uint96\"}],\"internalType\":\"structAllocationVoting.AllocationVote[][]\",\"name\":\"_illiquidVotes\",\"type\":\"tuple[][]\"}],\"name\":\"multiVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAssetsOut\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCore\",\"type\":\"address\"}],\"name\":\"setCore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setEnabledRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_zapFee\",\"type\":\"uint256\"}],\"name\":\"setZapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_receiptTokens\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"startInfiWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"startUnwinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakedTokens\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"}],\"name\":\"unstakeAndLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"farm\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"weight\",\"type\":\"uint96\"}],\"internalType\":\"structAllocationVoting.AllocationVote[]\",\"name\":\"_liquidVotes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"farm\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"weight\",\"type\":\"uint96\"}],\"internalType\":\"structAllocationVoting.AllocationVote[]\",\"name\":\"_illiquidVotes\",\"type\":\"tuple[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unwindingTimestamp\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unwindingTimestamp\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"zapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_unwindingEpochs\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"zapInAndLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"zapInAndStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// InfiniFiGatewayV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use InfiniFiGatewayV3MetaData.ABI instead.
var InfiniFiGatewayV3ABI = InfiniFiGatewayV3MetaData.ABI

// InfiniFiGatewayV3 is an auto generated Go binding around an Ethereum contract.
type InfiniFiGatewayV3 struct {
	InfiniFiGatewayV3Caller     // Read-only binding to the contract
	InfiniFiGatewayV3Transactor // Write-only binding to the contract
	InfiniFiGatewayV3Filterer   // Log filterer for contract events
}

// InfiniFiGatewayV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type InfiniFiGatewayV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniFiGatewayV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type InfiniFiGatewayV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniFiGatewayV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfiniFiGatewayV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfiniFiGatewayV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfiniFiGatewayV3Session struct {
	Contract     *InfiniFiGatewayV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InfiniFiGatewayV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfiniFiGatewayV3CallerSession struct {
	Contract *InfiniFiGatewayV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InfiniFiGatewayV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfiniFiGatewayV3TransactorSession struct {
	Contract     *InfiniFiGatewayV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InfiniFiGatewayV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type InfiniFiGatewayV3Raw struct {
	Contract *InfiniFiGatewayV3 // Generic contract binding to access the raw methods on
}

// InfiniFiGatewayV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfiniFiGatewayV3CallerRaw struct {
	Contract *InfiniFiGatewayV3Caller // Generic read-only contract binding to access the raw methods on
}

// InfiniFiGatewayV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfiniFiGatewayV3TransactorRaw struct {
	Contract *InfiniFiGatewayV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewInfiniFiGatewayV3 creates a new instance of InfiniFiGatewayV3, bound to a specific deployed contract.
func NewInfiniFiGatewayV3(address common.Address, backend bind.ContractBackend) (*InfiniFiGatewayV3, error) {
	contract, err := bindInfiniFiGatewayV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3{InfiniFiGatewayV3Caller: InfiniFiGatewayV3Caller{contract: contract}, InfiniFiGatewayV3Transactor: InfiniFiGatewayV3Transactor{contract: contract}, InfiniFiGatewayV3Filterer: InfiniFiGatewayV3Filterer{contract: contract}}, nil
}

// NewInfiniFiGatewayV3Caller creates a new read-only instance of InfiniFiGatewayV3, bound to a specific deployed contract.
func NewInfiniFiGatewayV3Caller(address common.Address, caller bind.ContractCaller) (*InfiniFiGatewayV3Caller, error) {
	contract, err := bindInfiniFiGatewayV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3Caller{contract: contract}, nil
}

// NewInfiniFiGatewayV3Transactor creates a new write-only instance of InfiniFiGatewayV3, bound to a specific deployed contract.
func NewInfiniFiGatewayV3Transactor(address common.Address, transactor bind.ContractTransactor) (*InfiniFiGatewayV3Transactor, error) {
	contract, err := bindInfiniFiGatewayV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3Transactor{contract: contract}, nil
}

// NewInfiniFiGatewayV3Filterer creates a new log filterer instance of InfiniFiGatewayV3, bound to a specific deployed contract.
func NewInfiniFiGatewayV3Filterer(address common.Address, filterer bind.ContractFilterer) (*InfiniFiGatewayV3Filterer, error) {
	contract, err := bindInfiniFiGatewayV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3Filterer{contract: contract}, nil
}

// bindInfiniFiGatewayV3 binds a generic wrapper to an already deployed contract.
func bindInfiniFiGatewayV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfiniFiGatewayV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfiniFiGatewayV3.Contract.InfiniFiGatewayV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.InfiniFiGatewayV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.InfiniFiGatewayV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfiniFiGatewayV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) Addresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Addresses(arg0 [32]byte) (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.Addresses(&_InfiniFiGatewayV3.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0x699f200f.
//
// Solidity: function addresses(bytes32 ) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) Addresses(arg0 [32]byte) (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.Addresses(&_InfiniFiGatewayV3.CallOpts, arg0)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) Core(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "core")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Core() (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.Core(&_InfiniFiGatewayV3.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) Core() (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.Core(&_InfiniFiGatewayV3.CallOpts)
}

// EnabledRouters is a free data retrieval call binding the contract method 0xc23544ad.
//
// Solidity: function enabledRouters(address ) view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) EnabledRouters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "enabledRouters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnabledRouters is a free data retrieval call binding the contract method 0xc23544ad.
//
// Solidity: function enabledRouters(address ) view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) EnabledRouters(arg0 common.Address) (bool, error) {
	return _InfiniFiGatewayV3.Contract.EnabledRouters(&_InfiniFiGatewayV3.CallOpts, arg0)
}

// EnabledRouters is a free data retrieval call binding the contract method 0xc23544ad.
//
// Solidity: function enabledRouters(address ) view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) EnabledRouters(arg0 common.Address) (bool, error) {
	return _InfiniFiGatewayV3.Contract.EnabledRouters(&_InfiniFiGatewayV3.CallOpts, arg0)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) GetAddress(opts *bind.CallOpts, _name string) (common.Address, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "getAddress", _name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) GetAddress(_name string) (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.GetAddress(&_InfiniFiGatewayV3.CallOpts, _name)
}

// GetAddress is a free data retrieval call binding the contract method 0xbf40fac1.
//
// Solidity: function getAddress(string _name) view returns(address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) GetAddress(_name string) (common.Address, error) {
	return _InfiniFiGatewayV3.Contract.GetAddress(&_InfiniFiGatewayV3.CallOpts, _name)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Paused() (bool, error) {
	return _InfiniFiGatewayV3.Contract.Paused(&_InfiniFiGatewayV3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) Paused() (bool, error) {
	return _InfiniFiGatewayV3.Contract.Paused(&_InfiniFiGatewayV3.CallOpts)
}

// ZapFee is a free data retrieval call binding the contract method 0xbf0673d2.
//
// Solidity: function zapFee() view returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Caller) ZapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfiniFiGatewayV3.contract.Call(opts, &out, "zapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZapFee is a free data retrieval call binding the contract method 0xbf0673d2.
//
// Solidity: function zapFee() view returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) ZapFee() (*big.Int, error) {
	return _InfiniFiGatewayV3.Contract.ZapFee(&_InfiniFiGatewayV3.CallOpts)
}

// ZapFee is a free data retrieval call binding the contract method 0xbf0673d2.
//
// Solidity: function zapFee() view returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3CallerSession) ZapFee() (*big.Int, error) {
	return _InfiniFiGatewayV3.Contract.ZapFee(&_InfiniFiGatewayV3.CallOpts)
}

// CancelInfiWithdrawal is a paid mutator transaction binding the contract method 0x8f5cba2b.
//
// Solidity: function cancelInfiWithdrawal(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) CancelInfiWithdrawal(opts *bind.TransactOpts, _unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "cancelInfiWithdrawal", _unwindingTimestamp, _newUnwindingEpochs)
}

// CancelInfiWithdrawal is a paid mutator transaction binding the contract method 0x8f5cba2b.
//
// Solidity: function cancelInfiWithdrawal(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) CancelInfiWithdrawal(_unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CancelInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp, _newUnwindingEpochs)
}

// CancelInfiWithdrawal is a paid mutator transaction binding the contract method 0x8f5cba2b.
//
// Solidity: function cancelInfiWithdrawal(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) CancelInfiWithdrawal(_unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CancelInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp, _newUnwindingEpochs)
}

// CancelUnwinding is a paid mutator transaction binding the contract method 0x1a11dc4c.
//
// Solidity: function cancelUnwinding(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) CancelUnwinding(opts *bind.TransactOpts, _unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "cancelUnwinding", _unwindingTimestamp, _newUnwindingEpochs)
}

// CancelUnwinding is a paid mutator transaction binding the contract method 0x1a11dc4c.
//
// Solidity: function cancelUnwinding(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) CancelUnwinding(_unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CancelUnwinding(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp, _newUnwindingEpochs)
}

// CancelUnwinding is a paid mutator transaction binding the contract method 0x1a11dc4c.
//
// Solidity: function cancelUnwinding(uint256 _unwindingTimestamp, uint32 _newUnwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) CancelUnwinding(_unwindingTimestamp *big.Int, _newUnwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CancelUnwinding(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp, _newUnwindingEpochs)
}

// ClaimRedemption is a paid mutator transaction binding the contract method 0xa16ba4e0.
//
// Solidity: function claimRedemption() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) ClaimRedemption(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "claimRedemption")
}

// ClaimRedemption is a paid mutator transaction binding the contract method 0xa16ba4e0.
//
// Solidity: function claimRedemption() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) ClaimRedemption() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ClaimRedemption(&_InfiniFiGatewayV3.TransactOpts)
}

// ClaimRedemption is a paid mutator transaction binding the contract method 0xa16ba4e0.
//
// Solidity: function claimRedemption() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) ClaimRedemption() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ClaimRedemption(&_InfiniFiGatewayV3.TransactOpts)
}

// CompleteInfiWithdrawal is a paid mutator transaction binding the contract method 0xcaab435f.
//
// Solidity: function completeInfiWithdrawal(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) CompleteInfiWithdrawal(opts *bind.TransactOpts, _user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "completeInfiWithdrawal", _user, _unwindingTimestamp)
}

// CompleteInfiWithdrawal is a paid mutator transaction binding the contract method 0xcaab435f.
//
// Solidity: function completeInfiWithdrawal(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) CompleteInfiWithdrawal(_user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CompleteInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _user, _unwindingTimestamp)
}

// CompleteInfiWithdrawal is a paid mutator transaction binding the contract method 0xcaab435f.
//
// Solidity: function completeInfiWithdrawal(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) CompleteInfiWithdrawal(_user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CompleteInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _user, _unwindingTimestamp)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xed4254ec.
//
// Solidity: function createPosition(uint256 _amount, uint32 _unwindingEpochs, address _recipient) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) CreatePosition(opts *bind.TransactOpts, _amount *big.Int, _unwindingEpochs uint32, _recipient common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "createPosition", _amount, _unwindingEpochs, _recipient)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xed4254ec.
//
// Solidity: function createPosition(uint256 _amount, uint32 _unwindingEpochs, address _recipient) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) CreatePosition(_amount *big.Int, _unwindingEpochs uint32, _recipient common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CreatePosition(&_InfiniFiGatewayV3.TransactOpts, _amount, _unwindingEpochs, _recipient)
}

// CreatePosition is a paid mutator transaction binding the contract method 0xed4254ec.
//
// Solidity: function createPosition(uint256 _amount, uint32 _unwindingEpochs, address _recipient) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) CreatePosition(_amount *big.Int, _unwindingEpochs uint32, _recipient common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.CreatePosition(&_InfiniFiGatewayV3.TransactOpts, _amount, _unwindingEpochs, _recipient)
}

// EmergencyAction is a paid mutator transaction binding the contract method 0x7df3927e.
//
// Solidity: function emergencyAction((address,uint256,bytes)[] calls) payable returns(bytes[] returnData)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) EmergencyAction(opts *bind.TransactOpts, calls []CoreControlledCall) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "emergencyAction", calls)
}

// EmergencyAction is a paid mutator transaction binding the contract method 0x7df3927e.
//
// Solidity: function emergencyAction((address,uint256,bytes)[] calls) payable returns(bytes[] returnData)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) EmergencyAction(calls []CoreControlledCall) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.EmergencyAction(&_InfiniFiGatewayV3.TransactOpts, calls)
}

// EmergencyAction is a paid mutator transaction binding the contract method 0x7df3927e.
//
// Solidity: function emergencyAction((address,uint256,bytes)[] calls) payable returns(bytes[] returnData)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) EmergencyAction(calls []CoreControlledCall) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.EmergencyAction(&_InfiniFiGatewayV3.TransactOpts, calls)
}

// IncreaseInfiWithdrawalPeriod is a paid mutator transaction binding the contract method 0xfa60e9ba.
//
// Solidity: function increaseInfiWithdrawalPeriod(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) IncreaseInfiWithdrawalPeriod(opts *bind.TransactOpts, _oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "increaseInfiWithdrawalPeriod", _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// IncreaseInfiWithdrawalPeriod is a paid mutator transaction binding the contract method 0xfa60e9ba.
//
// Solidity: function increaseInfiWithdrawalPeriod(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) IncreaseInfiWithdrawalPeriod(_oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.IncreaseInfiWithdrawalPeriod(&_InfiniFiGatewayV3.TransactOpts, _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// IncreaseInfiWithdrawalPeriod is a paid mutator transaction binding the contract method 0xfa60e9ba.
//
// Solidity: function increaseInfiWithdrawalPeriod(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) IncreaseInfiWithdrawalPeriod(_oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.IncreaseInfiWithdrawalPeriod(&_InfiniFiGatewayV3.TransactOpts, _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// IncreaseUnwindingEpochs is a paid mutator transaction binding the contract method 0xd5915096.
//
// Solidity: function increaseUnwindingEpochs(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) IncreaseUnwindingEpochs(opts *bind.TransactOpts, _oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "increaseUnwindingEpochs", _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// IncreaseUnwindingEpochs is a paid mutator transaction binding the contract method 0xd5915096.
//
// Solidity: function increaseUnwindingEpochs(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) IncreaseUnwindingEpochs(_oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.IncreaseUnwindingEpochs(&_InfiniFiGatewayV3.TransactOpts, _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// IncreaseUnwindingEpochs is a paid mutator transaction binding the contract method 0xd5915096.
//
// Solidity: function increaseUnwindingEpochs(uint32 _oldUnwindingEpochs, uint32 _newUnwindingEpochs, uint256 _shares) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) IncreaseUnwindingEpochs(_oldUnwindingEpochs uint32, _newUnwindingEpochs uint32, _shares *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.IncreaseUnwindingEpochs(&_InfiniFiGatewayV3.TransactOpts, _oldUnwindingEpochs, _newUnwindingEpochs, _shares)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _core) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Init(opts *bind.TransactOpts, _core common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "init", _core)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _core) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Init(_core common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Init(&_InfiniFiGatewayV3.TransactOpts, _core)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _core) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Init(_core common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Init(&_InfiniFiGatewayV3.TransactOpts, _core)
}

// LockInfi is a paid mutator transaction binding the contract method 0x7edd1a7e.
//
// Solidity: function lockInfi(uint256 _amount, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) LockInfi(opts *bind.TransactOpts, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "lockInfi", _amount, _unwindingEpochs)
}

// LockInfi is a paid mutator transaction binding the contract method 0x7edd1a7e.
//
// Solidity: function lockInfi(uint256 _amount, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) LockInfi(_amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.LockInfi(&_InfiniFiGatewayV3.TransactOpts, _amount, _unwindingEpochs)
}

// LockInfi is a paid mutator transaction binding the contract method 0x7edd1a7e.
//
// Solidity: function lockInfi(uint256 _amount, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) LockInfi(_amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.LockInfi(&_InfiniFiGatewayV3.TransactOpts, _amount, _unwindingEpochs)
}

// Migrate is a paid mutator transaction binding the contract method 0xf5d954c0.
//
// Solidity: function migrate(address _farm, address _token, uint256 _amount, uint32 _unwindingEpochs, uint256 _minIusdReceived) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Migrate(opts *bind.TransactOpts, _farm common.Address, _token common.Address, _amount *big.Int, _unwindingEpochs uint32, _minIusdReceived *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "migrate", _farm, _token, _amount, _unwindingEpochs, _minIusdReceived)
}

// Migrate is a paid mutator transaction binding the contract method 0xf5d954c0.
//
// Solidity: function migrate(address _farm, address _token, uint256 _amount, uint32 _unwindingEpochs, uint256 _minIusdReceived) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Migrate(_farm common.Address, _token common.Address, _amount *big.Int, _unwindingEpochs uint32, _minIusdReceived *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Migrate(&_InfiniFiGatewayV3.TransactOpts, _farm, _token, _amount, _unwindingEpochs, _minIusdReceived)
}

// Migrate is a paid mutator transaction binding the contract method 0xf5d954c0.
//
// Solidity: function migrate(address _farm, address _token, uint256 _amount, uint32 _unwindingEpochs, uint256 _minIusdReceived) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Migrate(_farm common.Address, _token common.Address, _amount *big.Int, _unwindingEpochs uint32, _minIusdReceived *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Migrate(&_InfiniFiGatewayV3.TransactOpts, _farm, _token, _amount, _unwindingEpochs, _minIusdReceived)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Mint(&_InfiniFiGatewayV3.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Mint(&_InfiniFiGatewayV3.TransactOpts, _to, _amount)
}

// MintAndLock is a paid mutator transaction binding the contract method 0x26cf8338.
//
// Solidity: function mintAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) MintAndLock(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "mintAndLock", _to, _amount, _unwindingEpochs)
}

// MintAndLock is a paid mutator transaction binding the contract method 0x26cf8338.
//
// Solidity: function mintAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) MintAndLock(_to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MintAndLock(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _unwindingEpochs)
}

// MintAndLock is a paid mutator transaction binding the contract method 0x26cf8338.
//
// Solidity: function mintAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) MintAndLock(_to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MintAndLock(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _unwindingEpochs)
}

// MintAndStake is a paid mutator transaction binding the contract method 0x230ae89c.
//
// Solidity: function mintAndStake(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) MintAndStake(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "mintAndStake", _to, _amount)
}

// MintAndStake is a paid mutator transaction binding the contract method 0x230ae89c.
//
// Solidity: function mintAndStake(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) MintAndStake(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MintAndStake(&_InfiniFiGatewayV3.TransactOpts, _to, _amount)
}

// MintAndStake is a paid mutator transaction binding the contract method 0x230ae89c.
//
// Solidity: function mintAndStake(address _to, uint256 _amount) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) MintAndStake(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MintAndStake(&_InfiniFiGatewayV3.TransactOpts, _to, _amount)
}

// MultiVote is a paid mutator transaction binding the contract method 0x4ec78381.
//
// Solidity: function multiVote(address[] _assets, uint32[] _unwindingEpochs, (address,uint96)[][] _liquidVotes, (address,uint96)[][] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) MultiVote(opts *bind.TransactOpts, _assets []common.Address, _unwindingEpochs []uint32, _liquidVotes [][]AllocationVotingAllocationVote, _illiquidVotes [][]AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "multiVote", _assets, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// MultiVote is a paid mutator transaction binding the contract method 0x4ec78381.
//
// Solidity: function multiVote(address[] _assets, uint32[] _unwindingEpochs, (address,uint96)[][] _liquidVotes, (address,uint96)[][] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) MultiVote(_assets []common.Address, _unwindingEpochs []uint32, _liquidVotes [][]AllocationVotingAllocationVote, _illiquidVotes [][]AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MultiVote(&_InfiniFiGatewayV3.TransactOpts, _assets, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// MultiVote is a paid mutator transaction binding the contract method 0x4ec78381.
//
// Solidity: function multiVote(address[] _assets, uint32[] _unwindingEpochs, (address,uint96)[][] _liquidVotes, (address,uint96)[][] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) MultiVote(_assets []common.Address, _unwindingEpochs []uint32, _liquidVotes [][]AllocationVotingAllocationVote, _illiquidVotes [][]AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.MultiVote(&_InfiniFiGatewayV3.TransactOpts, _assets, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Pause() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Pause(&_InfiniFiGatewayV3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Pause() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Pause(&_InfiniFiGatewayV3.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address _to, uint256 _amount, uint256 _minAssetsOut) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Redeem(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _minAssetsOut *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "redeem", _to, _amount, _minAssetsOut)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address _to, uint256 _amount, uint256 _minAssetsOut) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Redeem(_to common.Address, _amount *big.Int, _minAssetsOut *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Redeem(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _minAssetsOut)
}

// Redeem is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address _to, uint256 _amount, uint256 _minAssetsOut) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Redeem(_to common.Address, _amount *big.Int, _minAssetsOut *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Redeem(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _minAssetsOut)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) SetAddress(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "setAddress", _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetAddress(&_InfiniFiGatewayV3.TransactOpts, _name, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0x9b2ea4bd.
//
// Solidity: function setAddress(string _name, address _address) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) SetAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetAddress(&_InfiniFiGatewayV3.TransactOpts, _name, _address)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address newCore) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) SetCore(opts *bind.TransactOpts, newCore common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "setCore", newCore)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address newCore) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) SetCore(newCore common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetCore(&_InfiniFiGatewayV3.TransactOpts, newCore)
}

// SetCore is a paid mutator transaction binding the contract method 0x80009630.
//
// Solidity: function setCore(address newCore) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) SetCore(newCore common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetCore(&_InfiniFiGatewayV3.TransactOpts, newCore)
}

// SetEnabledRouter is a paid mutator transaction binding the contract method 0x7b96eef1.
//
// Solidity: function setEnabledRouter(address _router, bool _enabled) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) SetEnabledRouter(opts *bind.TransactOpts, _router common.Address, _enabled bool) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "setEnabledRouter", _router, _enabled)
}

// SetEnabledRouter is a paid mutator transaction binding the contract method 0x7b96eef1.
//
// Solidity: function setEnabledRouter(address _router, bool _enabled) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) SetEnabledRouter(_router common.Address, _enabled bool) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetEnabledRouter(&_InfiniFiGatewayV3.TransactOpts, _router, _enabled)
}

// SetEnabledRouter is a paid mutator transaction binding the contract method 0x7b96eef1.
//
// Solidity: function setEnabledRouter(address _router, bool _enabled) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) SetEnabledRouter(_router common.Address, _enabled bool) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetEnabledRouter(&_InfiniFiGatewayV3.TransactOpts, _router, _enabled)
}

// SetZapFee is a paid mutator transaction binding the contract method 0xa5a85b8d.
//
// Solidity: function setZapFee(uint256 _zapFee) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) SetZapFee(opts *bind.TransactOpts, _zapFee *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "setZapFee", _zapFee)
}

// SetZapFee is a paid mutator transaction binding the contract method 0xa5a85b8d.
//
// Solidity: function setZapFee(uint256 _zapFee) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) SetZapFee(_zapFee *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetZapFee(&_InfiniFiGatewayV3.TransactOpts, _zapFee)
}

// SetZapFee is a paid mutator transaction binding the contract method 0xa5a85b8d.
//
// Solidity: function setZapFee(uint256 _zapFee) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) SetZapFee(_zapFee *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.SetZapFee(&_InfiniFiGatewayV3.TransactOpts, _zapFee)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _to, uint256 _receiptTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Stake(opts *bind.TransactOpts, _to common.Address, _receiptTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "stake", _to, _receiptTokens)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _to, uint256 _receiptTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Stake(_to common.Address, _receiptTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Stake(&_InfiniFiGatewayV3.TransactOpts, _to, _receiptTokens)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _to, uint256 _receiptTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Stake(_to common.Address, _receiptTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Stake(&_InfiniFiGatewayV3.TransactOpts, _to, _receiptTokens)
}

// StartInfiWithdrawal is a paid mutator transaction binding the contract method 0xaa030aff.
//
// Solidity: function startInfiWithdrawal(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) StartInfiWithdrawal(opts *bind.TransactOpts, _shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "startInfiWithdrawal", _shares, _unwindingEpochs)
}

// StartInfiWithdrawal is a paid mutator transaction binding the contract method 0xaa030aff.
//
// Solidity: function startInfiWithdrawal(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) StartInfiWithdrawal(_shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.StartInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _shares, _unwindingEpochs)
}

// StartInfiWithdrawal is a paid mutator transaction binding the contract method 0xaa030aff.
//
// Solidity: function startInfiWithdrawal(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) StartInfiWithdrawal(_shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.StartInfiWithdrawal(&_InfiniFiGatewayV3.TransactOpts, _shares, _unwindingEpochs)
}

// StartUnwinding is a paid mutator transaction binding the contract method 0xc00716ea.
//
// Solidity: function startUnwinding(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) StartUnwinding(opts *bind.TransactOpts, _shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "startUnwinding", _shares, _unwindingEpochs)
}

// StartUnwinding is a paid mutator transaction binding the contract method 0xc00716ea.
//
// Solidity: function startUnwinding(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) StartUnwinding(_shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.StartUnwinding(&_InfiniFiGatewayV3.TransactOpts, _shares, _unwindingEpochs)
}

// StartUnwinding is a paid mutator transaction binding the contract method 0xc00716ea.
//
// Solidity: function startUnwinding(uint256 _shares, uint32 _unwindingEpochs) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) StartUnwinding(_shares *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.StartUnwinding(&_InfiniFiGatewayV3.TransactOpts, _shares, _unwindingEpochs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Unpause() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Unpause(&_InfiniFiGatewayV3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Unpause() (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Unpause(&_InfiniFiGatewayV3.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _to, uint256 _stakedTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Unstake(opts *bind.TransactOpts, _to common.Address, _stakedTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "unstake", _to, _stakedTokens)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _to, uint256 _stakedTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Unstake(_to common.Address, _stakedTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Unstake(&_InfiniFiGatewayV3.TransactOpts, _to, _stakedTokens)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _to, uint256 _stakedTokens) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Unstake(_to common.Address, _stakedTokens *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Unstake(&_InfiniFiGatewayV3.TransactOpts, _to, _stakedTokens)
}

// UnstakeAndLock is a paid mutator transaction binding the contract method 0xd95ce53b.
//
// Solidity: function unstakeAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) UnstakeAndLock(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "unstakeAndLock", _to, _amount, _unwindingEpochs)
}

// UnstakeAndLock is a paid mutator transaction binding the contract method 0xd95ce53b.
//
// Solidity: function unstakeAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) UnstakeAndLock(_to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.UnstakeAndLock(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _unwindingEpochs)
}

// UnstakeAndLock is a paid mutator transaction binding the contract method 0xd95ce53b.
//
// Solidity: function unstakeAndLock(address _to, uint256 _amount, uint32 _unwindingEpochs) returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) UnstakeAndLock(_to common.Address, _amount *big.Int, _unwindingEpochs uint32) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.UnstakeAndLock(&_InfiniFiGatewayV3.TransactOpts, _to, _amount, _unwindingEpochs)
}

// Vote is a paid mutator transaction binding the contract method 0xc1d9fe3e.
//
// Solidity: function vote(address _asset, uint32 _unwindingEpochs, (address,uint96)[] _liquidVotes, (address,uint96)[] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Vote(opts *bind.TransactOpts, _asset common.Address, _unwindingEpochs uint32, _liquidVotes []AllocationVotingAllocationVote, _illiquidVotes []AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "vote", _asset, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// Vote is a paid mutator transaction binding the contract method 0xc1d9fe3e.
//
// Solidity: function vote(address _asset, uint32 _unwindingEpochs, (address,uint96)[] _liquidVotes, (address,uint96)[] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Vote(_asset common.Address, _unwindingEpochs uint32, _liquidVotes []AllocationVotingAllocationVote, _illiquidVotes []AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Vote(&_InfiniFiGatewayV3.TransactOpts, _asset, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// Vote is a paid mutator transaction binding the contract method 0xc1d9fe3e.
//
// Solidity: function vote(address _asset, uint32 _unwindingEpochs, (address,uint96)[] _liquidVotes, (address,uint96)[] _illiquidVotes) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Vote(_asset common.Address, _unwindingEpochs uint32, _liquidVotes []AllocationVotingAllocationVote, _illiquidVotes []AllocationVotingAllocationVote) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Vote(&_InfiniFiGatewayV3.TransactOpts, _asset, _unwindingEpochs, _liquidVotes, _illiquidVotes)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Withdraw(opts *bind.TransactOpts, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "withdraw", _unwindingTimestamp)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Withdraw(_unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Withdraw(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Withdraw(_unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Withdraw(&_InfiniFiGatewayV3.TransactOpts, _unwindingTimestamp)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) Withdraw0(opts *bind.TransactOpts, _user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "withdraw0", _user, _unwindingTimestamp)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) Withdraw0(_user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Withdraw0(&_InfiniFiGatewayV3.TransactOpts, _user, _unwindingTimestamp)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address _user, uint256 _unwindingTimestamp) returns()
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) Withdraw0(_user common.Address, _unwindingTimestamp *big.Int) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.Withdraw0(&_InfiniFiGatewayV3.TransactOpts, _user, _unwindingTimestamp)
}

// ZapIn is a paid mutator transaction binding the contract method 0x7b4aedcf.
//
// Solidity: function zapIn(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) ZapIn(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "zapIn", _token, _amount, _router, _routerData, _to)
}

// ZapIn is a paid mutator transaction binding the contract method 0x7b4aedcf.
//
// Solidity: function zapIn(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) ZapIn(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapIn(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _to)
}

// ZapIn is a paid mutator transaction binding the contract method 0x7b4aedcf.
//
// Solidity: function zapIn(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) ZapIn(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapIn(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _to)
}

// ZapInAndLock is a paid mutator transaction binding the contract method 0x4a6bdbbe.
//
// Solidity: function zapInAndLock(address _token, uint256 _amount, address _router, bytes _routerData, uint32 _unwindingEpochs, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) ZapInAndLock(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _unwindingEpochs uint32, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "zapInAndLock", _token, _amount, _router, _routerData, _unwindingEpochs, _to)
}

// ZapInAndLock is a paid mutator transaction binding the contract method 0x4a6bdbbe.
//
// Solidity: function zapInAndLock(address _token, uint256 _amount, address _router, bytes _routerData, uint32 _unwindingEpochs, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) ZapInAndLock(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _unwindingEpochs uint32, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapInAndLock(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _unwindingEpochs, _to)
}

// ZapInAndLock is a paid mutator transaction binding the contract method 0x4a6bdbbe.
//
// Solidity: function zapInAndLock(address _token, uint256 _amount, address _router, bytes _routerData, uint32 _unwindingEpochs, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) ZapInAndLock(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _unwindingEpochs uint32, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapInAndLock(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _unwindingEpochs, _to)
}

// ZapInAndStake is a paid mutator transaction binding the contract method 0x365f1c00.
//
// Solidity: function zapInAndStake(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Transactor) ZapInAndStake(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.contract.Transact(opts, "zapInAndStake", _token, _amount, _router, _routerData, _to)
}

// ZapInAndStake is a paid mutator transaction binding the contract method 0x365f1c00.
//
// Solidity: function zapInAndStake(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Session) ZapInAndStake(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapInAndStake(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _to)
}

// ZapInAndStake is a paid mutator transaction binding the contract method 0x365f1c00.
//
// Solidity: function zapInAndStake(address _token, uint256 _amount, address _router, bytes _routerData, address _to) payable returns(uint256)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3TransactorSession) ZapInAndStake(_token common.Address, _amount *big.Int, _router common.Address, _routerData []byte, _to common.Address) (*types.Transaction, error) {
	return _InfiniFiGatewayV3.Contract.ZapInAndStake(&_InfiniFiGatewayV3.TransactOpts, _token, _amount, _router, _routerData, _to)
}

// InfiniFiGatewayV3AddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3AddressSetIterator struct {
	Event *InfiniFiGatewayV3AddressSet // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3AddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3AddressSet)
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
		it.Event = new(InfiniFiGatewayV3AddressSet)
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
func (it *InfiniFiGatewayV3AddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3AddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3AddressSet represents a AddressSet event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3AddressSet struct {
	Timestamp *big.Int
	Name      common.Hash
	Address   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x34765092fb4826a2f48dae7ee370790341f5fe0147f013fca4108e6ea65a68ed.
//
// Solidity: event AddressSet(uint256 timestamp, string indexed name, address _address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterAddressSet(opts *bind.FilterOpts, name []string) (*InfiniFiGatewayV3AddressSetIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "AddressSet", nameRule)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3AddressSetIterator{contract: _InfiniFiGatewayV3.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x34765092fb4826a2f48dae7ee370790341f5fe0147f013fca4108e6ea65a68ed.
//
// Solidity: event AddressSet(uint256 timestamp, string indexed name, address _address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3AddressSet, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "AddressSet", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3AddressSet)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0x34765092fb4826a2f48dae7ee370790341f5fe0147f013fca4108e6ea65a68ed.
//
// Solidity: event AddressSet(uint256 timestamp, string indexed name, address _address)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseAddressSet(log types.Log) (*InfiniFiGatewayV3AddressSet, error) {
	event := new(InfiniFiGatewayV3AddressSet)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3CoreUpdateIterator is returned from FilterCoreUpdate and is used to iterate over the raw logs and unpacked data for CoreUpdate events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3CoreUpdateIterator struct {
	Event *InfiniFiGatewayV3CoreUpdate // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3CoreUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3CoreUpdate)
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
		it.Event = new(InfiniFiGatewayV3CoreUpdate)
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
func (it *InfiniFiGatewayV3CoreUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3CoreUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3CoreUpdate represents a CoreUpdate event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3CoreUpdate struct {
	OldCore common.Address
	NewCore common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCoreUpdate is a free log retrieval operation binding the contract event 0x9209b7c8c06dcfd261686a663e7c55989337b18d59da5433c6f2835fb6970920.
//
// Solidity: event CoreUpdate(address indexed oldCore, address indexed newCore)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterCoreUpdate(opts *bind.FilterOpts, oldCore []common.Address, newCore []common.Address) (*InfiniFiGatewayV3CoreUpdateIterator, error) {

	var oldCoreRule []interface{}
	for _, oldCoreItem := range oldCore {
		oldCoreRule = append(oldCoreRule, oldCoreItem)
	}
	var newCoreRule []interface{}
	for _, newCoreItem := range newCore {
		newCoreRule = append(newCoreRule, newCoreItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "CoreUpdate", oldCoreRule, newCoreRule)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3CoreUpdateIterator{contract: _InfiniFiGatewayV3.contract, event: "CoreUpdate", logs: logs, sub: sub}, nil
}

// WatchCoreUpdate is a free log subscription operation binding the contract event 0x9209b7c8c06dcfd261686a663e7c55989337b18d59da5433c6f2835fb6970920.
//
// Solidity: event CoreUpdate(address indexed oldCore, address indexed newCore)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchCoreUpdate(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3CoreUpdate, oldCore []common.Address, newCore []common.Address) (event.Subscription, error) {

	var oldCoreRule []interface{}
	for _, oldCoreItem := range oldCore {
		oldCoreRule = append(oldCoreRule, oldCoreItem)
	}
	var newCoreRule []interface{}
	for _, newCoreItem := range newCore {
		newCoreRule = append(newCoreRule, newCoreItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "CoreUpdate", oldCoreRule, newCoreRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3CoreUpdate)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "CoreUpdate", log); err != nil {
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

// ParseCoreUpdate is a log parse operation binding the contract event 0x9209b7c8c06dcfd261686a663e7c55989337b18d59da5433c6f2835fb6970920.
//
// Solidity: event CoreUpdate(address indexed oldCore, address indexed newCore)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseCoreUpdate(log types.Log) (*InfiniFiGatewayV3CoreUpdate, error) {
	event := new(InfiniFiGatewayV3CoreUpdate)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "CoreUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3PausedIterator struct {
	Event *InfiniFiGatewayV3Paused // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3Paused)
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
		it.Event = new(InfiniFiGatewayV3Paused)
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
func (it *InfiniFiGatewayV3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3Paused represents a Paused event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterPaused(opts *bind.FilterOpts) (*InfiniFiGatewayV3PausedIterator, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3PausedIterator{contract: _InfiniFiGatewayV3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3Paused) (event.Subscription, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3Paused)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParsePaused(log types.Log) (*InfiniFiGatewayV3Paused, error) {
	event := new(InfiniFiGatewayV3Paused)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3SetEnabledRouterIterator is returned from FilterSetEnabledRouter and is used to iterate over the raw logs and unpacked data for SetEnabledRouter events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3SetEnabledRouterIterator struct {
	Event *InfiniFiGatewayV3SetEnabledRouter // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3SetEnabledRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3SetEnabledRouter)
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
		it.Event = new(InfiniFiGatewayV3SetEnabledRouter)
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
func (it *InfiniFiGatewayV3SetEnabledRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3SetEnabledRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3SetEnabledRouter represents a SetEnabledRouter event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3SetEnabledRouter struct {
	Timestamp *big.Int
	Router    common.Address
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetEnabledRouter is a free log retrieval operation binding the contract event 0xaf11e799c0dfe70b6ee7eeeb1e99889edf753033f557a3a5568838bdc770a08b.
//
// Solidity: event SetEnabledRouter(uint256 timestamp, address router, bool enabled)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterSetEnabledRouter(opts *bind.FilterOpts) (*InfiniFiGatewayV3SetEnabledRouterIterator, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "SetEnabledRouter")
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3SetEnabledRouterIterator{contract: _InfiniFiGatewayV3.contract, event: "SetEnabledRouter", logs: logs, sub: sub}, nil
}

// WatchSetEnabledRouter is a free log subscription operation binding the contract event 0xaf11e799c0dfe70b6ee7eeeb1e99889edf753033f557a3a5568838bdc770a08b.
//
// Solidity: event SetEnabledRouter(uint256 timestamp, address router, bool enabled)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchSetEnabledRouter(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3SetEnabledRouter) (event.Subscription, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "SetEnabledRouter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3SetEnabledRouter)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "SetEnabledRouter", log); err != nil {
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

// ParseSetEnabledRouter is a log parse operation binding the contract event 0xaf11e799c0dfe70b6ee7eeeb1e99889edf753033f557a3a5568838bdc770a08b.
//
// Solidity: event SetEnabledRouter(uint256 timestamp, address router, bool enabled)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseSetEnabledRouter(log types.Log) (*InfiniFiGatewayV3SetEnabledRouter, error) {
	event := new(InfiniFiGatewayV3SetEnabledRouter)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "SetEnabledRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3UnpausedIterator struct {
	Event *InfiniFiGatewayV3Unpaused // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3Unpaused)
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
		it.Event = new(InfiniFiGatewayV3Unpaused)
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
func (it *InfiniFiGatewayV3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3Unpaused represents a Unpaused event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*InfiniFiGatewayV3UnpausedIterator, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3UnpausedIterator{contract: _InfiniFiGatewayV3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3Unpaused) (event.Subscription, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3Unpaused)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseUnpaused(log types.Log) (*InfiniFiGatewayV3Unpaused, error) {
	event := new(InfiniFiGatewayV3Unpaused)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3ZapFeeSetIterator is returned from FilterZapFeeSet and is used to iterate over the raw logs and unpacked data for ZapFeeSet events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3ZapFeeSetIterator struct {
	Event *InfiniFiGatewayV3ZapFeeSet // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3ZapFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3ZapFeeSet)
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
		it.Event = new(InfiniFiGatewayV3ZapFeeSet)
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
func (it *InfiniFiGatewayV3ZapFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3ZapFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3ZapFeeSet represents a ZapFeeSet event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3ZapFeeSet struct {
	Timestamp *big.Int
	ZapFee    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZapFeeSet is a free log retrieval operation binding the contract event 0x675607ea61b02df8e5abbbe4772c0b60bb371fa80782206f2efa1b31e37311eb.
//
// Solidity: event ZapFeeSet(uint256 timestamp, uint256 zapFee)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterZapFeeSet(opts *bind.FilterOpts) (*InfiniFiGatewayV3ZapFeeSetIterator, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "ZapFeeSet")
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3ZapFeeSetIterator{contract: _InfiniFiGatewayV3.contract, event: "ZapFeeSet", logs: logs, sub: sub}, nil
}

// WatchZapFeeSet is a free log subscription operation binding the contract event 0x675607ea61b02df8e5abbbe4772c0b60bb371fa80782206f2efa1b31e37311eb.
//
// Solidity: event ZapFeeSet(uint256 timestamp, uint256 zapFee)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchZapFeeSet(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3ZapFeeSet) (event.Subscription, error) {

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "ZapFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3ZapFeeSet)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "ZapFeeSet", log); err != nil {
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

// ParseZapFeeSet is a log parse operation binding the contract event 0x675607ea61b02df8e5abbbe4772c0b60bb371fa80782206f2efa1b31e37311eb.
//
// Solidity: event ZapFeeSet(uint256 timestamp, uint256 zapFee)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseZapFeeSet(log types.Log) (*InfiniFiGatewayV3ZapFeeSet, error) {
	event := new(InfiniFiGatewayV3ZapFeeSet)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "ZapFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfiniFiGatewayV3ZapInIterator is returned from FilterZapIn and is used to iterate over the raw logs and unpacked data for ZapIn events raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3ZapInIterator struct {
	Event *InfiniFiGatewayV3ZapIn // Event containing the contract specifics and raw log

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
func (it *InfiniFiGatewayV3ZapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfiniFiGatewayV3ZapIn)
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
		it.Event = new(InfiniFiGatewayV3ZapIn)
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
func (it *InfiniFiGatewayV3ZapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfiniFiGatewayV3ZapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfiniFiGatewayV3ZapIn represents a ZapIn event raised by the InfiniFiGatewayV3 contract.
type InfiniFiGatewayV3ZapIn struct {
	Timestamp     *big.Int
	User          common.Address
	Token         common.Address
	Amount        *big.Int
	ReceiptTokens *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZapIn is a free log retrieval operation binding the contract event 0x864d8445691e89cf27c9d9cc6155ad638a9833f28365127ad89aa779a6d3cf83.
//
// Solidity: event ZapIn(uint256 timestamp, address indexed user, address indexed token, uint256 amount, uint256 receiptTokens)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) FilterZapIn(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*InfiniFiGatewayV3ZapInIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.FilterLogs(opts, "ZapIn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfiniFiGatewayV3ZapInIterator{contract: _InfiniFiGatewayV3.contract, event: "ZapIn", logs: logs, sub: sub}, nil
}

// WatchZapIn is a free log subscription operation binding the contract event 0x864d8445691e89cf27c9d9cc6155ad638a9833f28365127ad89aa779a6d3cf83.
//
// Solidity: event ZapIn(uint256 timestamp, address indexed user, address indexed token, uint256 amount, uint256 receiptTokens)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) WatchZapIn(opts *bind.WatchOpts, sink chan<- *InfiniFiGatewayV3ZapIn, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _InfiniFiGatewayV3.contract.WatchLogs(opts, "ZapIn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfiniFiGatewayV3ZapIn)
				if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "ZapIn", log); err != nil {
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

// ParseZapIn is a log parse operation binding the contract event 0x864d8445691e89cf27c9d9cc6155ad638a9833f28365127ad89aa779a6d3cf83.
//
// Solidity: event ZapIn(uint256 timestamp, address indexed user, address indexed token, uint256 amount, uint256 receiptTokens)
func (_InfiniFiGatewayV3 *InfiniFiGatewayV3Filterer) ParseZapIn(log types.Log) (*InfiniFiGatewayV3ZapIn, error) {
	event := new(InfiniFiGatewayV3ZapIn)
	if err := _InfiniFiGatewayV3.contract.UnpackLog(event, "ZapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
