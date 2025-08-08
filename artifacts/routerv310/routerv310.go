// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerv310

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

// Edge is an auto generated low-level Go binding around an user-defined struct.
type Edge struct {
	Id               *big.Int
	TokenIn          common.Address
	TokenOut         common.Address
	Adapter          common.Address
	Worker           common.Address
	ExtraData        []byte
	AmountInTotal    *big.Int
	AmountOutTotal   *big.Int
	AmountInCurrent  *big.Int
	AmountOutCurrent *big.Int
}

// Graph is an auto generated low-level Go binding around an user-defined struct.
type Graph struct {
	Vertices []Vertex
	Edges    []Edge
}

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// RouterResult is an auto generated low-level Go binding around an user-defined struct.
type RouterResult struct {
	Amount    *big.Int
	MinAmount *big.Int
	Calls     []MultiCall
}

// StrategyPathTask is an auto generated low-level Go binding around an user-defined struct.
type StrategyPathTask struct {
	CreditAccount     common.Address
	Graph             Graph
	Target            common.Address
	Adapters          []common.Address
	Slippage          *big.Int
	Force             bool
	InitTargetBalance *big.Int
	Calls             []MultiCall
}

// TokenData is an auto generated low-level Go binding around an user-defined struct.
type TokenData struct {
	Token           common.Address
	Balance         *big.Int
	LeftoverBalance *big.Int
	NumSplits       *big.Int
	ClaimRewards    bool
}

// Vertex is an auto generated low-level Go binding around an user-defined struct.
type Vertex struct {
	Token              common.Address
	Balance            *big.Int
	LeftoverBalance    *big.Int
	NumSplits          *big.Int
	CurrentOptimalEdge *big.Int
}

// Routerv310MetaData contains all meta data concerning the Routerv310 contract.
var Routerv310MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UnsupportedRouterComponent\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"RouterComponentUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetFutureRouter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"componentAddressByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICreditManagerV3\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"name\":\"createOpenStrategyPathTask\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentOptimalEdge\",\"type\":\"uint256\"}],\"internalType\":\"structVertex[]\",\"name\":\"vertices\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountInTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutCurrent\",\"type\":\"uint256\"}],\"internalType\":\"structEdge[]\",\"name\":\"edges\",\"type\":\"tuple[]\"}],\"internalType\":\"structGraph\",\"name\":\"graph\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"initTargetBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structStrategyPathTask\",\"name\":\"task\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"name\":\"createStrategyPathTask\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentOptimalEdge\",\"type\":\"uint256\"}],\"internalType\":\"structVertex[]\",\"name\":\"vertices\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountInTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInCurrent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutCurrent\",\"type\":\"uint256\"}],\"internalType\":\"structEdge[]\",\"name\":\"edges\",\"type\":\"tuple[]\"}],\"internalType\":\"structGraph\",\"name\":\"graph\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"initTargetBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structStrategyPathTask\",\"name\":\"task\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICreditManagerV3\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getAdapters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"name\":\"getDefaultTokenData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isRouterConfigurator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"knownComponentTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"knownTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevRouter\",\"type\":\"address\"}],\"name\":\"migrateRouterComponents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"name\":\"processClaims\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"name\":\"routeManyToOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"}],\"name\":\"routeOneToOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"}],\"name\":\"routeOneToOneDiff\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numSplits\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimRewards\",\"type\":\"bool\"}],\"internalType\":\"structTokenData[]\",\"name\":\"tData\",\"type\":\"tuple[]\"}],\"name\":\"routeOpenManyToOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routingManager\",\"outputs\":[{\"internalType\":\"contractIRoutingManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_futureRouter\",\"type\":\"address\"}],\"name\":\"setFutureRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"componentAddress\",\"type\":\"address\"}],\"name\":\"setPathComponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"componentAddresses\",\"type\":\"address[]\"}],\"name\":\"setPathComponentBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Routerv310ABI is the input ABI used to generate the binding from.
// Deprecated: Use Routerv310MetaData.ABI instead.
var Routerv310ABI = Routerv310MetaData.ABI

// Routerv310 is an auto generated Go binding around an Ethereum contract.
type Routerv310 struct {
	Routerv310Caller     // Read-only binding to the contract
	Routerv310Transactor // Write-only binding to the contract
	Routerv310Filterer   // Log filterer for contract events
}

// Routerv310Caller is an auto generated read-only Go binding around an Ethereum contract.
type Routerv310Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv310Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Routerv310Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv310Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Routerv310Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv310Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Routerv310Session struct {
	Contract     *Routerv310       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Routerv310CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Routerv310CallerSession struct {
	Contract *Routerv310Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Routerv310TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Routerv310TransactorSession struct {
	Contract     *Routerv310Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Routerv310Raw is an auto generated low-level Go binding around an Ethereum contract.
type Routerv310Raw struct {
	Contract *Routerv310 // Generic contract binding to access the raw methods on
}

// Routerv310CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Routerv310CallerRaw struct {
	Contract *Routerv310Caller // Generic read-only contract binding to access the raw methods on
}

// Routerv310TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Routerv310TransactorRaw struct {
	Contract *Routerv310Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterv310 creates a new instance of Routerv310, bound to a specific deployed contract.
func NewRouterv310(address common.Address, backend bind.ContractBackend) (*Routerv310, error) {
	contract, err := bindRouterv310(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Routerv310{Routerv310Caller: Routerv310Caller{contract: contract}, Routerv310Transactor: Routerv310Transactor{contract: contract}, Routerv310Filterer: Routerv310Filterer{contract: contract}}, nil
}

// NewRouterv310Caller creates a new read-only instance of Routerv310, bound to a specific deployed contract.
func NewRouterv310Caller(address common.Address, caller bind.ContractCaller) (*Routerv310Caller, error) {
	contract, err := bindRouterv310(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv310Caller{contract: contract}, nil
}

// NewRouterv310Transactor creates a new write-only instance of Routerv310, bound to a specific deployed contract.
func NewRouterv310Transactor(address common.Address, transactor bind.ContractTransactor) (*Routerv310Transactor, error) {
	contract, err := bindRouterv310(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv310Transactor{contract: contract}, nil
}

// NewRouterv310Filterer creates a new log filterer instance of Routerv310, bound to a specific deployed contract.
func NewRouterv310Filterer(address common.Address, filterer bind.ContractFilterer) (*Routerv310Filterer, error) {
	contract, err := bindRouterv310(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Routerv310Filterer{contract: contract}, nil
}

// bindRouterv310 binds a generic wrapper to an already deployed contract.
func bindRouterv310(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Routerv310MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv310 *Routerv310Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv310.Contract.Routerv310Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv310 *Routerv310Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv310.Contract.Routerv310Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv310 *Routerv310Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv310.Contract.Routerv310Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv310 *Routerv310CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv310.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv310 *Routerv310TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv310.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv310 *Routerv310TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv310.Contract.contract.Transact(opts, method, params...)
}

// ComponentAddressByType is a free data retrieval call binding the contract method 0x0eb1bdec.
//
// Solidity: function componentAddressByType(bytes32 ) view returns(address)
func (_Routerv310 *Routerv310Caller) ComponentAddressByType(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "componentAddressByType", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComponentAddressByType is a free data retrieval call binding the contract method 0x0eb1bdec.
//
// Solidity: function componentAddressByType(bytes32 ) view returns(address)
func (_Routerv310 *Routerv310Session) ComponentAddressByType(arg0 [32]byte) (common.Address, error) {
	return _Routerv310.Contract.ComponentAddressByType(&_Routerv310.CallOpts, arg0)
}

// ComponentAddressByType is a free data retrieval call binding the contract method 0x0eb1bdec.
//
// Solidity: function componentAddressByType(bytes32 ) view returns(address)
func (_Routerv310 *Routerv310CallerSession) ComponentAddressByType(arg0 [32]byte) (common.Address, error) {
	return _Routerv310.Contract.ComponentAddressByType(&_Routerv310.CallOpts, arg0)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Routerv310 *Routerv310Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Routerv310 *Routerv310Session) ContractType() ([32]byte, error) {
	return _Routerv310.Contract.ContractType(&_Routerv310.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_Routerv310 *Routerv310CallerSession) ContractType() ([32]byte, error) {
	return _Routerv310.Contract.ContractType(&_Routerv310.CallOpts)
}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv310 *Routerv310Caller) FutureRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "futureRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv310 *Routerv310Session) FutureRouter() (common.Address, error) {
	return _Routerv310.Contract.FutureRouter(&_Routerv310.CallOpts)
}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv310 *Routerv310CallerSession) FutureRouter() (common.Address, error) {
	return _Routerv310.Contract.FutureRouter(&_Routerv310.CallOpts)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv310 *Routerv310Caller) GetAdapters(opts *bind.CallOpts, creditManager common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "getAdapters", creditManager)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv310 *Routerv310Session) GetAdapters(creditManager common.Address) ([]common.Address, error) {
	return _Routerv310.Contract.GetAdapters(&_Routerv310.CallOpts, creditManager)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv310 *Routerv310CallerSession) GetAdapters(creditManager common.Address) ([]common.Address, error) {
	return _Routerv310.Contract.GetAdapters(&_Routerv310.CallOpts, creditManager)
}

// GetDefaultTokenData is a free data retrieval call binding the contract method 0x71552920.
//
// Solidity: function getDefaultTokenData(address creditAccount) view returns((address,uint256,uint256,uint256,bool)[] tData)
func (_Routerv310 *Routerv310Caller) GetDefaultTokenData(opts *bind.CallOpts, creditAccount common.Address) ([]TokenData, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "getDefaultTokenData", creditAccount)

	if err != nil {
		return *new([]TokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenData)).(*[]TokenData)

	return out0, err

}

// GetDefaultTokenData is a free data retrieval call binding the contract method 0x71552920.
//
// Solidity: function getDefaultTokenData(address creditAccount) view returns((address,uint256,uint256,uint256,bool)[] tData)
func (_Routerv310 *Routerv310Session) GetDefaultTokenData(creditAccount common.Address) ([]TokenData, error) {
	return _Routerv310.Contract.GetDefaultTokenData(&_Routerv310.CallOpts, creditAccount)
}

// GetDefaultTokenData is a free data retrieval call binding the contract method 0x71552920.
//
// Solidity: function getDefaultTokenData(address creditAccount) view returns((address,uint256,uint256,uint256,bool)[] tData)
func (_Routerv310 *Routerv310CallerSession) GetDefaultTokenData(creditAccount common.Address) ([]TokenData, error) {
	return _Routerv310.Contract.GetDefaultTokenData(&_Routerv310.CallOpts, creditAccount)
}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv310 *Routerv310Caller) IsRouterConfigurator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "isRouterConfigurator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv310 *Routerv310Session) IsRouterConfigurator(account common.Address) (bool, error) {
	return _Routerv310.Contract.IsRouterConfigurator(&_Routerv310.CallOpts, account)
}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv310 *Routerv310CallerSession) IsRouterConfigurator(account common.Address) (bool, error) {
	return _Routerv310.Contract.IsRouterConfigurator(&_Routerv310.CallOpts, account)
}

// KnownComponentTypes is a free data retrieval call binding the contract method 0x47a3ccf4.
//
// Solidity: function knownComponentTypes() view returns(bytes32[] knownTypes)
func (_Routerv310 *Routerv310Caller) KnownComponentTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "knownComponentTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// KnownComponentTypes is a free data retrieval call binding the contract method 0x47a3ccf4.
//
// Solidity: function knownComponentTypes() view returns(bytes32[] knownTypes)
func (_Routerv310 *Routerv310Session) KnownComponentTypes() ([][32]byte, error) {
	return _Routerv310.Contract.KnownComponentTypes(&_Routerv310.CallOpts)
}

// KnownComponentTypes is a free data retrieval call binding the contract method 0x47a3ccf4.
//
// Solidity: function knownComponentTypes() view returns(bytes32[] knownTypes)
func (_Routerv310 *Routerv310CallerSession) KnownComponentTypes() ([][32]byte, error) {
	return _Routerv310.Contract.KnownComponentTypes(&_Routerv310.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv310 *Routerv310Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv310 *Routerv310Session) Owner() (common.Address, error) {
	return _Routerv310.Contract.Owner(&_Routerv310.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv310 *Routerv310CallerSession) Owner() (common.Address, error) {
	return _Routerv310.Contract.Owner(&_Routerv310.CallOpts)
}

// RoutingManager is a free data retrieval call binding the contract method 0x04bbab14.
//
// Solidity: function routingManager() view returns(address)
func (_Routerv310 *Routerv310Caller) RoutingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "routingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoutingManager is a free data retrieval call binding the contract method 0x04bbab14.
//
// Solidity: function routingManager() view returns(address)
func (_Routerv310 *Routerv310Session) RoutingManager() (common.Address, error) {
	return _Routerv310.Contract.RoutingManager(&_Routerv310.CallOpts)
}

// RoutingManager is a free data retrieval call binding the contract method 0x04bbab14.
//
// Solidity: function routingManager() view returns(address)
func (_Routerv310 *Routerv310CallerSession) RoutingManager() (common.Address, error) {
	return _Routerv310.Contract.RoutingManager(&_Routerv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv310 *Routerv310Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Routerv310.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv310 *Routerv310Session) Version() (*big.Int, error) {
	return _Routerv310.Contract.Version(&_Routerv310.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv310 *Routerv310CallerSession) Version() (*big.Int, error) {
	return _Routerv310.Contract.Version(&_Routerv310.CallOpts)
}

// CreateOpenStrategyPathTask is a paid mutator transaction binding the contract method 0x37d631df.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310Transactor) CreateOpenStrategyPathTask(opts *bind.TransactOpts, creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "createOpenStrategyPathTask", creditManager, target, slippage, tData)
}

// CreateOpenStrategyPathTask is a paid mutator transaction binding the contract method 0x37d631df.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310Session) CreateOpenStrategyPathTask(creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.CreateOpenStrategyPathTask(&_Routerv310.TransactOpts, creditManager, target, slippage, tData)
}

// CreateOpenStrategyPathTask is a paid mutator transaction binding the contract method 0x37d631df.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310TransactorSession) CreateOpenStrategyPathTask(creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.CreateOpenStrategyPathTask(&_Routerv310.TransactOpts, creditManager, target, slippage, tData)
}

// CreateStrategyPathTask is a paid mutator transaction binding the contract method 0x3dfbc906.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, uint256 slippage, bool force, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310Transactor) CreateStrategyPathTask(opts *bind.TransactOpts, creditAccount common.Address, target common.Address, slippage *big.Int, force bool, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "createStrategyPathTask", creditAccount, target, slippage, force, tData)
}

// CreateStrategyPathTask is a paid mutator transaction binding the contract method 0x3dfbc906.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, uint256 slippage, bool force, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310Session) CreateStrategyPathTask(creditAccount common.Address, target common.Address, slippage *big.Int, force bool, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.CreateStrategyPathTask(&_Routerv310.TransactOpts, creditAccount, target, slippage, force, tData)
}

// CreateStrategyPathTask is a paid mutator transaction binding the contract method 0x3dfbc906.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, uint256 slippage, bool force, (address,uint256,uint256,uint256,bool)[] tData) returns((address,((address,uint256,uint256,uint256,uint256)[],(uint256,address,address,address,address,bytes,uint256,uint256,uint256,uint256)[]),address,address[],uint256,bool,uint256,(address,bytes)[]) task)
func (_Routerv310 *Routerv310TransactorSession) CreateStrategyPathTask(creditAccount common.Address, target common.Address, slippage *big.Int, force bool, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.CreateStrategyPathTask(&_Routerv310.TransactOpts, creditAccount, target, slippage, force, tData)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv310 *Routerv310Transactor) MigrateRouterComponents(opts *bind.TransactOpts, _prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "migrateRouterComponents", _prevRouter)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv310 *Routerv310Session) MigrateRouterComponents(_prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.MigrateRouterComponents(&_Routerv310.TransactOpts, _prevRouter)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv310 *Routerv310TransactorSession) MigrateRouterComponents(_prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.MigrateRouterComponents(&_Routerv310.TransactOpts, _prevRouter)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0xe2796e56.
//
// Solidity: function processClaims(address creditAccount, (address,uint256,uint256,uint256,bool)[] tData) returns((address,bytes)[] calls)
func (_Routerv310 *Routerv310Transactor) ProcessClaims(opts *bind.TransactOpts, creditAccount common.Address, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "processClaims", creditAccount, tData)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0xe2796e56.
//
// Solidity: function processClaims(address creditAccount, (address,uint256,uint256,uint256,bool)[] tData) returns((address,bytes)[] calls)
func (_Routerv310 *Routerv310Session) ProcessClaims(creditAccount common.Address, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.ProcessClaims(&_Routerv310.TransactOpts, creditAccount, tData)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0xe2796e56.
//
// Solidity: function processClaims(address creditAccount, (address,uint256,uint256,uint256,bool)[] tData) returns((address,bytes)[] calls)
func (_Routerv310 *Routerv310TransactorSession) ProcessClaims(creditAccount common.Address, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.ProcessClaims(&_Routerv310.TransactOpts, creditAccount, tData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv310 *Routerv310Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv310 *Routerv310Session) RenounceOwnership() (*types.Transaction, error) {
	return _Routerv310.Contract.RenounceOwnership(&_Routerv310.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv310 *Routerv310TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Routerv310.Contract.RenounceOwnership(&_Routerv310.TransactOpts)
}

// RouteManyToOne is a paid mutator transaction binding the contract method 0x1bc63d56.
//
// // Solidity: function routeManyToOne(address creditAccount, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
// func (_Routerv310 *Routerv310Transactor) RouteManyToOne(opts *bind.TransactOpts, creditAccount common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
// 	return _Routerv310.contract.Transact(opts, "routeManyToOne", creditAccount, target, slippage, tData)
// }

// // RouteManyToOne is a paid mutator transaction binding the contract method 0x1bc63d56.
// //
// // Solidity: function routeManyToOne(address creditAccount, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
// func (_Routerv310 *Routerv310Session) RouteManyToOne(creditAccount common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
// 	return _Routerv310.Contract.RouteManyToOne(&_Routerv310.TransactOpts, creditAccount, target, slippage, tData)
// }

// // RouteManyToOne is a paid mutator transaction binding the contract method 0x1bc63d56.
// //
// // Solidity: function routeManyToOne(address creditAccount, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
// func (_Routerv310 *Routerv310TransactorSession) RouteManyToOne(creditAccount common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
// 	return _Routerv310.Contract.RouteManyToOne(&_Routerv310.TransactOpts, creditAccount, target, slippage, tData)
// }

// RouteOneToOne is a paid mutator transaction binding the contract method 0x095ec1a6.
//
// Solidity: function routeOneToOne(address creditAccount, address tokenIn, uint256 amount, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Transactor) RouteOneToOne(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, amount *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "routeOneToOne", creditAccount, tokenIn, amount, target, slippage, numSplits)
}

// RouteOneToOne is a paid mutator transaction binding the contract method 0x095ec1a6.
//
// Solidity: function routeOneToOne(address creditAccount, address tokenIn, uint256 amount, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Session) RouteOneToOne(creditAccount common.Address, tokenIn common.Address, amount *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOneToOne(&_Routerv310.TransactOpts, creditAccount, tokenIn, amount, target, slippage, numSplits)
}

// RouteOneToOne is a paid mutator transaction binding the contract method 0x095ec1a6.
//
// Solidity: function routeOneToOne(address creditAccount, address tokenIn, uint256 amount, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310TransactorSession) RouteOneToOne(creditAccount common.Address, tokenIn common.Address, amount *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOneToOne(&_Routerv310.TransactOpts, creditAccount, tokenIn, amount, target, slippage, numSplits)
}

// RouteOneToOneDiff is a paid mutator transaction binding the contract method 0xdc5a5ef8.
//
// Solidity: function routeOneToOneDiff(address creditAccount, address tokenIn, uint256 balance, uint256 leftoverBalance, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Transactor) RouteOneToOneDiff(opts *bind.TransactOpts, creditAccount common.Address, tokenIn common.Address, balance *big.Int, leftoverBalance *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "routeOneToOneDiff", creditAccount, tokenIn, balance, leftoverBalance, target, slippage, numSplits)
}

// RouteOneToOneDiff is a paid mutator transaction binding the contract method 0xdc5a5ef8.
//
// Solidity: function routeOneToOneDiff(address creditAccount, address tokenIn, uint256 balance, uint256 leftoverBalance, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Session) RouteOneToOneDiff(creditAccount common.Address, tokenIn common.Address, balance *big.Int, leftoverBalance *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOneToOneDiff(&_Routerv310.TransactOpts, creditAccount, tokenIn, balance, leftoverBalance, target, slippage, numSplits)
}

// RouteOneToOneDiff is a paid mutator transaction binding the contract method 0xdc5a5ef8.
//
// Solidity: function routeOneToOneDiff(address creditAccount, address tokenIn, uint256 balance, uint256 leftoverBalance, address target, uint256 slippage, uint256 numSplits) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310TransactorSession) RouteOneToOneDiff(creditAccount common.Address, tokenIn common.Address, balance *big.Int, leftoverBalance *big.Int, target common.Address, slippage *big.Int, numSplits *big.Int) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOneToOneDiff(&_Routerv310.TransactOpts, creditAccount, tokenIn, balance, leftoverBalance, target, slippage, numSplits)
}

// RouteOpenManyToOne is a paid mutator transaction binding the contract method 0x33b6562b.
//
// Solidity: function routeOpenManyToOne(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Transactor) RouteOpenManyToOne(opts *bind.TransactOpts, creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "routeOpenManyToOne", creditManager, target, slippage, tData)
}

// RouteOpenManyToOne is a paid mutator transaction binding the contract method 0x33b6562b.
//
// Solidity: function routeOpenManyToOne(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310Session) RouteOpenManyToOne(creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOpenManyToOne(&_Routerv310.TransactOpts, creditManager, target, slippage, tData)
}

// RouteOpenManyToOne is a paid mutator transaction binding the contract method 0x33b6562b.
//
// Solidity: function routeOpenManyToOne(address creditManager, address target, uint256 slippage, (address,uint256,uint256,uint256,bool)[] tData) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv310 *Routerv310TransactorSession) RouteOpenManyToOne(creditManager common.Address, target common.Address, slippage *big.Int, tData []TokenData) (*types.Transaction, error) {
	return _Routerv310.Contract.RouteOpenManyToOne(&_Routerv310.TransactOpts, creditManager, target, slippage, tData)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv310 *Routerv310Transactor) SetFutureRouter(opts *bind.TransactOpts, _futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "setFutureRouter", _futureRouter)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv310 *Routerv310Session) SetFutureRouter(_futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetFutureRouter(&_Routerv310.TransactOpts, _futureRouter)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv310 *Routerv310TransactorSession) SetFutureRouter(_futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetFutureRouter(&_Routerv310.TransactOpts, _futureRouter)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv310 *Routerv310Transactor) SetPathComponent(opts *bind.TransactOpts, componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "setPathComponent", componentAddress)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv310 *Routerv310Session) SetPathComponent(componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetPathComponent(&_Routerv310.TransactOpts, componentAddress)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv310 *Routerv310TransactorSession) SetPathComponent(componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetPathComponent(&_Routerv310.TransactOpts, componentAddress)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv310 *Routerv310Transactor) SetPathComponentBatch(opts *bind.TransactOpts, componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "setPathComponentBatch", componentAddresses)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv310 *Routerv310Session) SetPathComponentBatch(componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetPathComponentBatch(&_Routerv310.TransactOpts, componentAddresses)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv310 *Routerv310TransactorSession) SetPathComponentBatch(componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.SetPathComponentBatch(&_Routerv310.TransactOpts, componentAddresses)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv310 *Routerv310Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Routerv310.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv310 *Routerv310Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.TransferOwnership(&_Routerv310.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv310 *Routerv310TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Routerv310.Contract.TransferOwnership(&_Routerv310.TransactOpts, newOwner)
}

// Routerv310OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Routerv310 contract.
type Routerv310OwnershipTransferredIterator struct {
	Event *Routerv310OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Routerv310OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv310OwnershipTransferred)
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
		it.Event = new(Routerv310OwnershipTransferred)
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
func (it *Routerv310OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv310OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv310OwnershipTransferred represents a OwnershipTransferred event raised by the Routerv310 contract.
type Routerv310OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Routerv310 *Routerv310Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Routerv310OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Routerv310.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Routerv310OwnershipTransferredIterator{contract: _Routerv310.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Routerv310 *Routerv310Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Routerv310OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Routerv310.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv310OwnershipTransferred)
				if err := _Routerv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Routerv310 *Routerv310Filterer) ParseOwnershipTransferred(log types.Log) (*Routerv310OwnershipTransferred, error) {
	event := new(Routerv310OwnershipTransferred)
	if err := _Routerv310.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv310RouterComponentUpdateIterator is returned from FilterRouterComponentUpdate and is used to iterate over the raw logs and unpacked data for RouterComponentUpdate events raised by the Routerv310 contract.
type Routerv310RouterComponentUpdateIterator struct {
	Event *Routerv310RouterComponentUpdate // Event containing the contract specifics and raw log

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
func (it *Routerv310RouterComponentUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv310RouterComponentUpdate)
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
		it.Event = new(Routerv310RouterComponentUpdate)
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
func (it *Routerv310RouterComponentUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv310RouterComponentUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv310RouterComponentUpdate represents a RouterComponentUpdate event raised by the Routerv310 contract.
type Routerv310RouterComponentUpdate struct {
	Arg0    [32]byte
	Arg1    common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRouterComponentUpdate is a free log retrieval operation binding the contract event 0x8f3c1cfaa7d59dfb81bd10faf7d41340b1d6e8a7ad6de89c55c22928158173f5.
//
// Solidity: event RouterComponentUpdate(bytes32 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv310 *Routerv310Filterer) FilterRouterComponentUpdate(opts *bind.FilterOpts, arg0 [][32]byte, arg1 []common.Address) (*Routerv310RouterComponentUpdateIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Routerv310.contract.FilterLogs(opts, "RouterComponentUpdate", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &Routerv310RouterComponentUpdateIterator{contract: _Routerv310.contract, event: "RouterComponentUpdate", logs: logs, sub: sub}, nil
}

// WatchRouterComponentUpdate is a free log subscription operation binding the contract event 0x8f3c1cfaa7d59dfb81bd10faf7d41340b1d6e8a7ad6de89c55c22928158173f5.
//
// Solidity: event RouterComponentUpdate(bytes32 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv310 *Routerv310Filterer) WatchRouterComponentUpdate(opts *bind.WatchOpts, sink chan<- *Routerv310RouterComponentUpdate, arg0 [][32]byte, arg1 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Routerv310.contract.WatchLogs(opts, "RouterComponentUpdate", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv310RouterComponentUpdate)
				if err := _Routerv310.contract.UnpackLog(event, "RouterComponentUpdate", log); err != nil {
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

// ParseRouterComponentUpdate is a log parse operation binding the contract event 0x8f3c1cfaa7d59dfb81bd10faf7d41340b1d6e8a7ad6de89c55c22928158173f5.
//
// Solidity: event RouterComponentUpdate(bytes32 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv310 *Routerv310Filterer) ParseRouterComponentUpdate(log types.Log) (*Routerv310RouterComponentUpdate, error) {
	event := new(Routerv310RouterComponentUpdate)
	if err := _Routerv310.contract.UnpackLog(event, "RouterComponentUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv310SetFutureRouterIterator is returned from FilterSetFutureRouter and is used to iterate over the raw logs and unpacked data for SetFutureRouter events raised by the Routerv310 contract.
type Routerv310SetFutureRouterIterator struct {
	Event *Routerv310SetFutureRouter // Event containing the contract specifics and raw log

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
func (it *Routerv310SetFutureRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv310SetFutureRouter)
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
		it.Event = new(Routerv310SetFutureRouter)
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
func (it *Routerv310SetFutureRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv310SetFutureRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv310SetFutureRouter represents a SetFutureRouter event raised by the Routerv310 contract.
type Routerv310SetFutureRouter struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetFutureRouter is a free log retrieval operation binding the contract event 0x38ee79447c54ed2235ae0312a2a622f96c8dcb3ba5b20ceb62cd62edeb19ee03.
//
// Solidity: event SetFutureRouter(address indexed arg0)
func (_Routerv310 *Routerv310Filterer) FilterSetFutureRouter(opts *bind.FilterOpts, arg0 []common.Address) (*Routerv310SetFutureRouterIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Routerv310.contract.FilterLogs(opts, "SetFutureRouter", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &Routerv310SetFutureRouterIterator{contract: _Routerv310.contract, event: "SetFutureRouter", logs: logs, sub: sub}, nil
}

// WatchSetFutureRouter is a free log subscription operation binding the contract event 0x38ee79447c54ed2235ae0312a2a622f96c8dcb3ba5b20ceb62cd62edeb19ee03.
//
// Solidity: event SetFutureRouter(address indexed arg0)
func (_Routerv310 *Routerv310Filterer) WatchSetFutureRouter(opts *bind.WatchOpts, sink chan<- *Routerv310SetFutureRouter, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Routerv310.contract.WatchLogs(opts, "SetFutureRouter", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv310SetFutureRouter)
				if err := _Routerv310.contract.UnpackLog(event, "SetFutureRouter", log); err != nil {
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

// ParseSetFutureRouter is a log parse operation binding the contract event 0x38ee79447c54ed2235ae0312a2a622f96c8dcb3ba5b20ceb62cd62edeb19ee03.
//
// Solidity: event SetFutureRouter(address indexed arg0)
func (_Routerv310 *Routerv310Filterer) ParseSetFutureRouter(log types.Log) (*Routerv310SetFutureRouter, error) {
	event := new(Routerv310SetFutureRouter)
	if err := _Routerv310.contract.UnpackLog(event, "SetFutureRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
