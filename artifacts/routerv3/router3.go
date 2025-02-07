// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerv3

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

// Balance is an auto generated low-level Go binding around an user-defined struct.
type Balance struct {
	Token   common.Address
	Balance *big.Int
}

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// PathOption is an auto generated low-level Go binding around an user-defined struct.
type PathOption struct {
	Target       common.Address
	Option       uint8
	TotalOptions uint8
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
	Balances          []Balance
	LeftoverBalances  []Balance
	Target            common.Address
	Connectors        []common.Address
	Adapters          []common.Address
	Slippage          *big.Int
	Force             bool
	TargetType        uint8
	FoundAdapters     []TokenAdapters
	InitTargetBalance *big.Int
	Calls             []MultiCall
}

// SwapTask is an auto generated low-level Go binding around an user-defined struct.
type SwapTask struct {
	SwapOperation  uint8
	CreditAccount  common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	Connectors     []common.Address
	Amount         *big.Int
	LeftoverAmount *big.Int
}

// TokenAdapters is an auto generated low-level Go binding around an user-defined struct.
type TokenAdapters struct {
	Token           common.Address
	DepositAdapter  common.Address
	WithdrawAdapter common.Address
}

// TokenToTokenType is an auto generated low-level Go binding around an user-defined struct.
type TokenToTokenType struct {
	Token     common.Address
	TokenType uint8
}

// TokenTypeToResolver is an auto generated low-level Go binding around an user-defined struct.
type TokenTypeToResolver struct {
	TokenType0 uint8
	TokenType1 uint8
	Resolver   uint8
}

// Routerv3MetaData contains all meta data concerning the Routerv3 contract.
var Routerv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NoSpaceForSlippageCallException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UnknownToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UnsupportedRouterComponent\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ttIn\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ttOut\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"rc\",\"type\":\"uint8\"}],\"name\":\"ResolverUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"RouterComponentUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SetFutureRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"tt\",\"type\":\"uint8\"}],\"name\":\"TokenTypeUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"componentAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICreditManagerV3\",\"name\":\"creditManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"leftoverBalances\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"createOpenStrategyPathTask\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"leftoverBalances\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"targetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawAdapter\",\"type\":\"address\"}],\"internalType\":\"structTokenAdapters[]\",\"name\":\"foundAdapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"initTargetBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structStrategyPathTask\",\"name\":\"task\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"createStrategyPathTask\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"leftoverBalances\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"targetType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawAdapter\",\"type\":\"address\"}],\"internalType\":\"structTokenAdapters[]\",\"name\":\"foundAdapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"initTargetBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structStrategyPathTask\",\"name\":\"task\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumSwapOperation\",\"name\":\"swapOperation\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leftoverAmount\",\"type\":\"uint256\"}],\"internalType\":\"structSwapTask\",\"name\":\"swapTask\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"findAllSwaps\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"expectedBalances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"leftoverBalances\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"totalOptions\",\"type\":\"uint8\"}],\"internalType\":\"structPathOption[]\",\"name\":\"pathOptions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"loops\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"force\",\"type\":\"bool\"}],\"name\":\"findBestClosePath\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"findOneTokenPath\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"leftoverBalances\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"slippage\",\"type\":\"uint256\"}],\"name\":\"findOpenStrategyPath\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"internalType\":\"structRouterResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICreditManagerV3\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getAdapters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ttIn\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"ttOut\",\"type\":\"uint8\"}],\"name\":\"getResolver\",\"outputs\":[{\"internalType\":\"contractIPathResolver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isRouterConfigurator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevRouter\",\"type\":\"address\"}],\"name\":\"migrateRouterComponents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"resolvers\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_futureRouter\",\"type\":\"address\"}],\"name\":\"setFutureRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"componentAddress\",\"type\":\"address\"}],\"name\":\"setPathComponent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"componentAddresses\",\"type\":\"address[]\"}],\"name\":\"setPathComponentBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"tokenType0\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenType1\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"resolver\",\"type\":\"uint8\"}],\"internalType\":\"structTokenTypeToResolver[]\",\"name\":\"tokenTypeToResolvers\",\"type\":\"tuple[]\"}],\"name\":\"setResolversBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenToTokenType[]\",\"name\":\"tokensToTokenTypes\",\"type\":\"tuple[]\"}],\"name\":\"setTokenTypesBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenTypes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Routerv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Routerv3MetaData.ABI instead.
var Routerv3ABI = Routerv3MetaData.ABI

// Routerv3 is an auto generated Go binding around an Ethereum contract.
type Routerv3 struct {
	Routerv3Caller     // Read-only binding to the contract
	Routerv3Transactor // Write-only binding to the contract
	Routerv3Filterer   // Log filterer for contract events
}

// Routerv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Routerv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Routerv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Routerv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Routerv3Session struct {
	Contract     *Routerv3         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Routerv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Routerv3CallerSession struct {
	Contract *Routerv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Routerv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Routerv3TransactorSession struct {
	Contract     *Routerv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Routerv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Routerv3Raw struct {
	Contract *Routerv3 // Generic contract binding to access the raw methods on
}

// Routerv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Routerv3CallerRaw struct {
	Contract *Routerv3Caller // Generic read-only contract binding to access the raw methods on
}

// Routerv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Routerv3TransactorRaw struct {
	Contract *Routerv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterv3 creates a new instance of Routerv3, bound to a specific deployed contract.
func NewRouterv3(address common.Address, backend bind.ContractBackend) (*Routerv3, error) {
	contract, err := bindRouterv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Routerv3{Routerv3Caller: Routerv3Caller{contract: contract}, Routerv3Transactor: Routerv3Transactor{contract: contract}, Routerv3Filterer: Routerv3Filterer{contract: contract}}, nil
}

// NewRouterv3Caller creates a new read-only instance of Routerv3, bound to a specific deployed contract.
func NewRouterv3Caller(address common.Address, caller bind.ContractCaller) (*Routerv3Caller, error) {
	contract, err := bindRouterv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv3Caller{contract: contract}, nil
}

// NewRouterv3Transactor creates a new write-only instance of Routerv3, bound to a specific deployed contract.
func NewRouterv3Transactor(address common.Address, transactor bind.ContractTransactor) (*Routerv3Transactor, error) {
	contract, err := bindRouterv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv3Transactor{contract: contract}, nil
}

// NewRouterv3Filterer creates a new log filterer instance of Routerv3, bound to a specific deployed contract.
func NewRouterv3Filterer(address common.Address, filterer bind.ContractFilterer) (*Routerv3Filterer, error) {
	contract, err := bindRouterv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Routerv3Filterer{contract: contract}, nil
}

// bindRouterv3 binds a generic wrapper to an already deployed contract.
func bindRouterv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Routerv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv3 *Routerv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv3.Contract.Routerv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv3 *Routerv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv3.Contract.Routerv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv3 *Routerv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv3.Contract.Routerv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv3 *Routerv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv3 *Routerv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv3 *Routerv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv3.Contract.contract.Transact(opts, method, params...)
}

// ComponentAddressById is a free data retrieval call binding the contract method 0x77d3e647.
//
// Solidity: function componentAddressById(uint8 ) view returns(address)
func (_Routerv3 *Routerv3Caller) ComponentAddressById(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "componentAddressById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComponentAddressById is a free data retrieval call binding the contract method 0x77d3e647.
//
// Solidity: function componentAddressById(uint8 ) view returns(address)
func (_Routerv3 *Routerv3Session) ComponentAddressById(arg0 uint8) (common.Address, error) {
	return _Routerv3.Contract.ComponentAddressById(&_Routerv3.CallOpts, arg0)
}

// ComponentAddressById is a free data retrieval call binding the contract method 0x77d3e647.
//
// Solidity: function componentAddressById(uint8 ) view returns(address)
func (_Routerv3 *Routerv3CallerSession) ComponentAddressById(arg0 uint8) (common.Address, error) {
	return _Routerv3.Contract.ComponentAddressById(&_Routerv3.CallOpts, arg0)
}

// CreateOpenStrategyPathTask is a free data retrieval call binding the contract method 0xe654d4ec.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3Caller) CreateOpenStrategyPathTask(opts *bind.CallOpts, creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (StrategyPathTask, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "createOpenStrategyPathTask", creditManager, balances, leftoverBalances, target, connectors, slippage)

	if err != nil {
		return *new(StrategyPathTask), err
	}

	out0 := *abi.ConvertType(out[0], new(StrategyPathTask)).(*StrategyPathTask)

	return out0, err

}

// CreateOpenStrategyPathTask is a free data retrieval call binding the contract method 0xe654d4ec.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3Session) CreateOpenStrategyPathTask(creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (StrategyPathTask, error) {
	return _Routerv3.Contract.CreateOpenStrategyPathTask(&_Routerv3.CallOpts, creditManager, balances, leftoverBalances, target, connectors, slippage)
}

// CreateOpenStrategyPathTask is a free data retrieval call binding the contract method 0xe654d4ec.
//
// Solidity: function createOpenStrategyPathTask(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3CallerSession) CreateOpenStrategyPathTask(creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (StrategyPathTask, error) {
	return _Routerv3.Contract.CreateOpenStrategyPathTask(&_Routerv3.CallOpts, creditManager, balances, leftoverBalances, target, connectors, slippage)
}

// CreateStrategyPathTask is a free data retrieval call binding the contract method 0x5168e775.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, address[] connectors, uint256 slippage, bool force) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3Caller) CreateStrategyPathTask(opts *bind.CallOpts, creditAccount common.Address, target common.Address, connectors []common.Address, slippage *big.Int, force bool) (StrategyPathTask, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "createStrategyPathTask", creditAccount, target, connectors, slippage, force)

	if err != nil {
		return *new(StrategyPathTask), err
	}

	out0 := *abi.ConvertType(out[0], new(StrategyPathTask)).(*StrategyPathTask)

	return out0, err

}

// CreateStrategyPathTask is a free data retrieval call binding the contract method 0x5168e775.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, address[] connectors, uint256 slippage, bool force) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3Session) CreateStrategyPathTask(creditAccount common.Address, target common.Address, connectors []common.Address, slippage *big.Int, force bool) (StrategyPathTask, error) {
	return _Routerv3.Contract.CreateStrategyPathTask(&_Routerv3.CallOpts, creditAccount, target, connectors, slippage, force)
}

// CreateStrategyPathTask is a free data retrieval call binding the contract method 0x5168e775.
//
// Solidity: function createStrategyPathTask(address creditAccount, address target, address[] connectors, uint256 slippage, bool force) view returns((address,(address,uint256)[],(address,uint256)[],address,address[],address[],uint256,bool,uint8,(address,address,address)[],uint256,(address,bytes)[]) task)
func (_Routerv3 *Routerv3CallerSession) CreateStrategyPathTask(creditAccount common.Address, target common.Address, connectors []common.Address, slippage *big.Int, force bool) (StrategyPathTask, error) {
	return _Routerv3.Contract.CreateStrategyPathTask(&_Routerv3.CallOpts, creditAccount, target, connectors, slippage, force)
}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv3 *Routerv3Caller) FutureRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "futureRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv3 *Routerv3Session) FutureRouter() (common.Address, error) {
	return _Routerv3.Contract.FutureRouter(&_Routerv3.CallOpts)
}

// FutureRouter is a free data retrieval call binding the contract method 0x108033a5.
//
// Solidity: function futureRouter() view returns(address)
func (_Routerv3 *Routerv3CallerSession) FutureRouter() (common.Address, error) {
	return _Routerv3.Contract.FutureRouter(&_Routerv3.CallOpts)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv3 *Routerv3Caller) GetAdapters(opts *bind.CallOpts, creditManager common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "getAdapters", creditManager)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv3 *Routerv3Session) GetAdapters(creditManager common.Address) ([]common.Address, error) {
	return _Routerv3.Contract.GetAdapters(&_Routerv3.CallOpts, creditManager)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(address[] result)
func (_Routerv3 *Routerv3CallerSession) GetAdapters(creditManager common.Address) ([]common.Address, error) {
	return _Routerv3.Contract.GetAdapters(&_Routerv3.CallOpts, creditManager)
}

// GetResolver is a free data retrieval call binding the contract method 0x620416e6.
//
// Solidity: function getResolver(uint8 ttIn, uint8 ttOut) view returns(address)
func (_Routerv3 *Routerv3Caller) GetResolver(opts *bind.CallOpts, ttIn uint8, ttOut uint8) (common.Address, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "getResolver", ttIn, ttOut)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetResolver is a free data retrieval call binding the contract method 0x620416e6.
//
// Solidity: function getResolver(uint8 ttIn, uint8 ttOut) view returns(address)
func (_Routerv3 *Routerv3Session) GetResolver(ttIn uint8, ttOut uint8) (common.Address, error) {
	return _Routerv3.Contract.GetResolver(&_Routerv3.CallOpts, ttIn, ttOut)
}

// GetResolver is a free data retrieval call binding the contract method 0x620416e6.
//
// Solidity: function getResolver(uint8 ttIn, uint8 ttOut) view returns(address)
func (_Routerv3 *Routerv3CallerSession) GetResolver(ttIn uint8, ttOut uint8) (common.Address, error) {
	return _Routerv3.Contract.GetResolver(&_Routerv3.CallOpts, ttIn, ttOut)
}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv3 *Routerv3Caller) IsRouterConfigurator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "isRouterConfigurator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv3 *Routerv3Session) IsRouterConfigurator(account common.Address) (bool, error) {
	return _Routerv3.Contract.IsRouterConfigurator(&_Routerv3.CallOpts, account)
}

// IsRouterConfigurator is a free data retrieval call binding the contract method 0x429b2de9.
//
// Solidity: function isRouterConfigurator(address account) view returns(bool)
func (_Routerv3 *Routerv3CallerSession) IsRouterConfigurator(account common.Address) (bool, error) {
	return _Routerv3.Contract.IsRouterConfigurator(&_Routerv3.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv3 *Routerv3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv3 *Routerv3Session) Owner() (common.Address, error) {
	return _Routerv3.Contract.Owner(&_Routerv3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Routerv3 *Routerv3CallerSession) Owner() (common.Address, error) {
	return _Routerv3.Contract.Owner(&_Routerv3.CallOpts)
}

// Resolvers is a free data retrieval call binding the contract method 0xee47c965.
//
// Solidity: function resolvers(uint8 , uint8 ) view returns(uint8)
func (_Routerv3 *Routerv3Caller) Resolvers(opts *bind.CallOpts, arg0 uint8, arg1 uint8) (uint8, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "resolvers", arg0, arg1)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Resolvers is a free data retrieval call binding the contract method 0xee47c965.
//
// Solidity: function resolvers(uint8 , uint8 ) view returns(uint8)
func (_Routerv3 *Routerv3Session) Resolvers(arg0 uint8, arg1 uint8) (uint8, error) {
	return _Routerv3.Contract.Resolvers(&_Routerv3.CallOpts, arg0, arg1)
}

// Resolvers is a free data retrieval call binding the contract method 0xee47c965.
//
// Solidity: function resolvers(uint8 , uint8 ) view returns(uint8)
func (_Routerv3 *Routerv3CallerSession) Resolvers(arg0 uint8, arg1 uint8) (uint8, error) {
	return _Routerv3.Contract.Resolvers(&_Routerv3.CallOpts, arg0, arg1)
}

// TokenTypes is a free data retrieval call binding the contract method 0xb39f252d.
//
// Solidity: function tokenTypes(address ) view returns(uint8)
func (_Routerv3 *Routerv3Caller) TokenTypes(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "tokenTypes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TokenTypes is a free data retrieval call binding the contract method 0xb39f252d.
//
// Solidity: function tokenTypes(address ) view returns(uint8)
func (_Routerv3 *Routerv3Session) TokenTypes(arg0 common.Address) (uint8, error) {
	return _Routerv3.Contract.TokenTypes(&_Routerv3.CallOpts, arg0)
}

// TokenTypes is a free data retrieval call binding the contract method 0xb39f252d.
//
// Solidity: function tokenTypes(address ) view returns(uint8)
func (_Routerv3 *Routerv3CallerSession) TokenTypes(arg0 common.Address) (uint8, error) {
	return _Routerv3.Contract.TokenTypes(&_Routerv3.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv3 *Routerv3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Routerv3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv3 *Routerv3Session) Version() (*big.Int, error) {
	return _Routerv3.Contract.Version(&_Routerv3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Routerv3 *Routerv3CallerSession) Version() (*big.Int, error) {
	return _Routerv3.Contract.Version(&_Routerv3.CallOpts)
}

// FindAllSwaps is a paid mutator transaction binding the contract method 0x8d3fd28b.
//
// Solidity: function findAllSwaps((uint8,address,address,address,address[],uint256,uint256) swapTask, uint256 slippage) returns((uint256,uint256,(address,bytes)[])[] result)
func (_Routerv3 *Routerv3Transactor) FindAllSwaps(opts *bind.TransactOpts, swapTask SwapTask, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "findAllSwaps", swapTask, slippage)
}

// FindAllSwaps is a paid mutator transaction binding the contract method 0x8d3fd28b.
//
// Solidity: function findAllSwaps((uint8,address,address,address,address[],uint256,uint256) swapTask, uint256 slippage) returns((uint256,uint256,(address,bytes)[])[] result)
func (_Routerv3 *Routerv3Session) FindAllSwaps(swapTask SwapTask, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindAllSwaps(&_Routerv3.TransactOpts, swapTask, slippage)
}

// FindAllSwaps is a paid mutator transaction binding the contract method 0x8d3fd28b.
//
// Solidity: function findAllSwaps((uint8,address,address,address,address[],uint256,uint256) swapTask, uint256 slippage) returns((uint256,uint256,(address,bytes)[])[] result)
func (_Routerv3 *Routerv3TransactorSession) FindAllSwaps(swapTask SwapTask, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindAllSwaps(&_Routerv3.TransactOpts, swapTask, slippage)
}

// FindBestClosePath is a paid mutator transaction binding the contract method 0xf0a29424.
//
// // Solidity: function findBestClosePath(address creditAccount, (address,uint256)[] expectedBalances, (address,uint256)[] leftoverBalances, address[] connectors, uint256 slippage, (address,uint8,uint8)[] pathOptions, uint256 loops, bool force) returns((uint256,uint256,(address,bytes)[]) result)
// func (_Routerv3 *Routerv3Transactor) FindBestClosePath(opts *bind.TransactOpts, creditAccount common.Address, expectedBalances []Balance, leftoverBalances []Balance, connectors []common.Address, slippage *big.Int, pathOptions []PathOption, loops *big.Int, force bool) (*types.Transaction, error) {
// 	return _Routerv3.contract.Transact(opts, "findBestClosePath", creditAccount, expectedBalances, leftoverBalances, connectors, slippage, pathOptions, loops, force)
// }

// // FindBestClosePath is a paid mutator transaction binding the contract method 0xf0a29424.
// //
// // Solidity: function findBestClosePath(address creditAccount, (address,uint256)[] expectedBalances, (address,uint256)[] leftoverBalances, address[] connectors, uint256 slippage, (address,uint8,uint8)[] pathOptions, uint256 loops, bool force) returns((uint256,uint256,(address,bytes)[]) result)
// func (_Routerv3 *Routerv3Session) FindBestClosePath(creditAccount common.Address, expectedBalances []Balance, leftoverBalances []Balance, connectors []common.Address, slippage *big.Int, pathOptions []PathOption, loops *big.Int, force bool) (*types.Transaction, error) {
// 	return _Routerv3.Contract.FindBestClosePath(&_Routerv3.TransactOpts, creditAccount, expectedBalances, leftoverBalances, connectors, slippage, pathOptions, loops, force)
// }

// // FindBestClosePath is a paid mutator transaction binding the contract method 0xf0a29424.
// //
// // Solidity: function findBestClosePath(address creditAccount, (address,uint256)[] expectedBalances, (address,uint256)[] leftoverBalances, address[] connectors, uint256 slippage, (address,uint8,uint8)[] pathOptions, uint256 loops, bool force) returns((uint256,uint256,(address,bytes)[]) result)
// func (_Routerv3 *Routerv3TransactorSession) FindBestClosePath(creditAccount common.Address, expectedBalances []Balance, leftoverBalances []Balance, connectors []common.Address, slippage *big.Int, pathOptions []PathOption, loops *big.Int, force bool) (*types.Transaction, error) {
// 	return _Routerv3.Contract.FindBestClosePath(&_Routerv3.TransactOpts, creditAccount, expectedBalances, leftoverBalances, connectors, slippage, pathOptions, loops, force)
// }

// FindOneTokenPath is a paid mutator transaction binding the contract method 0x0cf7f659.
//
// Solidity: function findOneTokenPath(address tokenIn, uint256 amount, address tokenOut, address creditAccount, address[] connectors, uint256 slippage) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3Transactor) FindOneTokenPath(opts *bind.TransactOpts, tokenIn common.Address, amount *big.Int, tokenOut common.Address, creditAccount common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "findOneTokenPath", tokenIn, amount, tokenOut, creditAccount, connectors, slippage)
}

// FindOneTokenPath is a paid mutator transaction binding the contract method 0x0cf7f659.
//
// Solidity: function findOneTokenPath(address tokenIn, uint256 amount, address tokenOut, address creditAccount, address[] connectors, uint256 slippage) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3Session) FindOneTokenPath(tokenIn common.Address, amount *big.Int, tokenOut common.Address, creditAccount common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindOneTokenPath(&_Routerv3.TransactOpts, tokenIn, amount, tokenOut, creditAccount, connectors, slippage)
}

// FindOneTokenPath is a paid mutator transaction binding the contract method 0x0cf7f659.
//
// Solidity: function findOneTokenPath(address tokenIn, uint256 amount, address tokenOut, address creditAccount, address[] connectors, uint256 slippage) returns((uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3TransactorSession) FindOneTokenPath(tokenIn common.Address, amount *big.Int, tokenOut common.Address, creditAccount common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindOneTokenPath(&_Routerv3.TransactOpts, tokenIn, amount, tokenOut, creditAccount, connectors, slippage)
}

// FindOpenStrategyPath is a paid mutator transaction binding the contract method 0x09c60c61.
//
// Solidity: function findOpenStrategyPath(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) returns((address,uint256)[], (uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3Transactor) FindOpenStrategyPath(opts *bind.TransactOpts, creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "findOpenStrategyPath", creditManager, balances, leftoverBalances, target, connectors, slippage)
}

// FindOpenStrategyPath is a paid mutator transaction binding the contract method 0x09c60c61.
//
// Solidity: function findOpenStrategyPath(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) returns((address,uint256)[], (uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3Session) FindOpenStrategyPath(creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindOpenStrategyPath(&_Routerv3.TransactOpts, creditManager, balances, leftoverBalances, target, connectors, slippage)
}

// FindOpenStrategyPath is a paid mutator transaction binding the contract method 0x09c60c61.
//
// Solidity: function findOpenStrategyPath(address creditManager, (address,uint256)[] balances, (address,uint256)[] leftoverBalances, address target, address[] connectors, uint256 slippage) returns((address,uint256)[], (uint256,uint256,(address,bytes)[]))
func (_Routerv3 *Routerv3TransactorSession) FindOpenStrategyPath(creditManager common.Address, balances []Balance, leftoverBalances []Balance, target common.Address, connectors []common.Address, slippage *big.Int) (*types.Transaction, error) {
	return _Routerv3.Contract.FindOpenStrategyPath(&_Routerv3.TransactOpts, creditManager, balances, leftoverBalances, target, connectors, slippage)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv3 *Routerv3Transactor) MigrateRouterComponents(opts *bind.TransactOpts, _prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "migrateRouterComponents", _prevRouter)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv3 *Routerv3Session) MigrateRouterComponents(_prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.MigrateRouterComponents(&_Routerv3.TransactOpts, _prevRouter)
}

// MigrateRouterComponents is a paid mutator transaction binding the contract method 0xc17ea347.
//
// Solidity: function migrateRouterComponents(address _prevRouter) returns()
func (_Routerv3 *Routerv3TransactorSession) MigrateRouterComponents(_prevRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.MigrateRouterComponents(&_Routerv3.TransactOpts, _prevRouter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv3 *Routerv3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv3 *Routerv3Session) RenounceOwnership() (*types.Transaction, error) {
	return _Routerv3.Contract.RenounceOwnership(&_Routerv3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Routerv3 *Routerv3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Routerv3.Contract.RenounceOwnership(&_Routerv3.TransactOpts)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv3 *Routerv3Transactor) SetFutureRouter(opts *bind.TransactOpts, _futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "setFutureRouter", _futureRouter)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv3 *Routerv3Session) SetFutureRouter(_futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetFutureRouter(&_Routerv3.TransactOpts, _futureRouter)
}

// SetFutureRouter is a paid mutator transaction binding the contract method 0x8a1acdc0.
//
// Solidity: function setFutureRouter(address _futureRouter) returns()
func (_Routerv3 *Routerv3TransactorSession) SetFutureRouter(_futureRouter common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetFutureRouter(&_Routerv3.TransactOpts, _futureRouter)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv3 *Routerv3Transactor) SetPathComponent(opts *bind.TransactOpts, componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "setPathComponent", componentAddress)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv3 *Routerv3Session) SetPathComponent(componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetPathComponent(&_Routerv3.TransactOpts, componentAddress)
}

// SetPathComponent is a paid mutator transaction binding the contract method 0x7e07e68a.
//
// Solidity: function setPathComponent(address componentAddress) returns()
func (_Routerv3 *Routerv3TransactorSession) SetPathComponent(componentAddress common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetPathComponent(&_Routerv3.TransactOpts, componentAddress)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv3 *Routerv3Transactor) SetPathComponentBatch(opts *bind.TransactOpts, componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "setPathComponentBatch", componentAddresses)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv3 *Routerv3Session) SetPathComponentBatch(componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetPathComponentBatch(&_Routerv3.TransactOpts, componentAddresses)
}

// SetPathComponentBatch is a paid mutator transaction binding the contract method 0x3686d3b4.
//
// Solidity: function setPathComponentBatch(address[] componentAddresses) returns()
func (_Routerv3 *Routerv3TransactorSession) SetPathComponentBatch(componentAddresses []common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.SetPathComponentBatch(&_Routerv3.TransactOpts, componentAddresses)
}

// SetResolversBatch is a paid mutator transaction binding the contract method 0x08c98417.
//
// Solidity: function setResolversBatch((uint8,uint8,uint8)[] tokenTypeToResolvers) returns()
func (_Routerv3 *Routerv3Transactor) SetResolversBatch(opts *bind.TransactOpts, tokenTypeToResolvers []TokenTypeToResolver) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "setResolversBatch", tokenTypeToResolvers)
}

// SetResolversBatch is a paid mutator transaction binding the contract method 0x08c98417.
//
// Solidity: function setResolversBatch((uint8,uint8,uint8)[] tokenTypeToResolvers) returns()
func (_Routerv3 *Routerv3Session) SetResolversBatch(tokenTypeToResolvers []TokenTypeToResolver) (*types.Transaction, error) {
	return _Routerv3.Contract.SetResolversBatch(&_Routerv3.TransactOpts, tokenTypeToResolvers)
}

// SetResolversBatch is a paid mutator transaction binding the contract method 0x08c98417.
//
// Solidity: function setResolversBatch((uint8,uint8,uint8)[] tokenTypeToResolvers) returns()
func (_Routerv3 *Routerv3TransactorSession) SetResolversBatch(tokenTypeToResolvers []TokenTypeToResolver) (*types.Transaction, error) {
	return _Routerv3.Contract.SetResolversBatch(&_Routerv3.TransactOpts, tokenTypeToResolvers)
}

// SetTokenTypesBatch is a paid mutator transaction binding the contract method 0x51541d04.
//
// Solidity: function setTokenTypesBatch((address,uint8)[] tokensToTokenTypes) returns()
func (_Routerv3 *Routerv3Transactor) SetTokenTypesBatch(opts *bind.TransactOpts, tokensToTokenTypes []TokenToTokenType) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "setTokenTypesBatch", tokensToTokenTypes)
}

// SetTokenTypesBatch is a paid mutator transaction binding the contract method 0x51541d04.
//
// Solidity: function setTokenTypesBatch((address,uint8)[] tokensToTokenTypes) returns()
func (_Routerv3 *Routerv3Session) SetTokenTypesBatch(tokensToTokenTypes []TokenToTokenType) (*types.Transaction, error) {
	return _Routerv3.Contract.SetTokenTypesBatch(&_Routerv3.TransactOpts, tokensToTokenTypes)
}

// SetTokenTypesBatch is a paid mutator transaction binding the contract method 0x51541d04.
//
// Solidity: function setTokenTypesBatch((address,uint8)[] tokensToTokenTypes) returns()
func (_Routerv3 *Routerv3TransactorSession) SetTokenTypesBatch(tokensToTokenTypes []TokenToTokenType) (*types.Transaction, error) {
	return _Routerv3.Contract.SetTokenTypesBatch(&_Routerv3.TransactOpts, tokensToTokenTypes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv3 *Routerv3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Routerv3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv3 *Routerv3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.TransferOwnership(&_Routerv3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Routerv3 *Routerv3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Routerv3.Contract.TransferOwnership(&_Routerv3.TransactOpts, newOwner)
}

// Routerv3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Routerv3 contract.
type Routerv3OwnershipTransferredIterator struct {
	Event *Routerv3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Routerv3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv3OwnershipTransferred)
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
		it.Event = new(Routerv3OwnershipTransferred)
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
func (it *Routerv3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv3OwnershipTransferred represents a OwnershipTransferred event raised by the Routerv3 contract.
type Routerv3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Routerv3 *Routerv3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Routerv3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Routerv3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Routerv3OwnershipTransferredIterator{contract: _Routerv3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Routerv3 *Routerv3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Routerv3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Routerv3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv3OwnershipTransferred)
				if err := _Routerv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Routerv3 *Routerv3Filterer) ParseOwnershipTransferred(log types.Log) (*Routerv3OwnershipTransferred, error) {
	event := new(Routerv3OwnershipTransferred)
	if err := _Routerv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv3ResolverUpdateIterator is returned from FilterResolverUpdate and is used to iterate over the raw logs and unpacked data for ResolverUpdate events raised by the Routerv3 contract.
type Routerv3ResolverUpdateIterator struct {
	Event *Routerv3ResolverUpdate // Event containing the contract specifics and raw log

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
func (it *Routerv3ResolverUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv3ResolverUpdate)
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
		it.Event = new(Routerv3ResolverUpdate)
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
func (it *Routerv3ResolverUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv3ResolverUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv3ResolverUpdate represents a ResolverUpdate event raised by the Routerv3 contract.
type Routerv3ResolverUpdate struct {
	TtIn  uint8
	TtOut uint8
	Rc    uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterResolverUpdate is a free log retrieval operation binding the contract event 0x9081d4c394cb22ad7d47837e87c89a58d1051fdf26813daa3f1c928c3ca16b30.
//
// Solidity: event ResolverUpdate(uint8 indexed ttIn, uint8 indexed ttOut, uint8 indexed rc)
func (_Routerv3 *Routerv3Filterer) FilterResolverUpdate(opts *bind.FilterOpts, ttIn []uint8, ttOut []uint8, rc []uint8) (*Routerv3ResolverUpdateIterator, error) {

	var ttInRule []interface{}
	for _, ttInItem := range ttIn {
		ttInRule = append(ttInRule, ttInItem)
	}
	var ttOutRule []interface{}
	for _, ttOutItem := range ttOut {
		ttOutRule = append(ttOutRule, ttOutItem)
	}
	var rcRule []interface{}
	for _, rcItem := range rc {
		rcRule = append(rcRule, rcItem)
	}

	logs, sub, err := _Routerv3.contract.FilterLogs(opts, "ResolverUpdate", ttInRule, ttOutRule, rcRule)
	if err != nil {
		return nil, err
	}
	return &Routerv3ResolverUpdateIterator{contract: _Routerv3.contract, event: "ResolverUpdate", logs: logs, sub: sub}, nil
}

// WatchResolverUpdate is a free log subscription operation binding the contract event 0x9081d4c394cb22ad7d47837e87c89a58d1051fdf26813daa3f1c928c3ca16b30.
//
// Solidity: event ResolverUpdate(uint8 indexed ttIn, uint8 indexed ttOut, uint8 indexed rc)
func (_Routerv3 *Routerv3Filterer) WatchResolverUpdate(opts *bind.WatchOpts, sink chan<- *Routerv3ResolverUpdate, ttIn []uint8, ttOut []uint8, rc []uint8) (event.Subscription, error) {

	var ttInRule []interface{}
	for _, ttInItem := range ttIn {
		ttInRule = append(ttInRule, ttInItem)
	}
	var ttOutRule []interface{}
	for _, ttOutItem := range ttOut {
		ttOutRule = append(ttOutRule, ttOutItem)
	}
	var rcRule []interface{}
	for _, rcItem := range rc {
		rcRule = append(rcRule, rcItem)
	}

	logs, sub, err := _Routerv3.contract.WatchLogs(opts, "ResolverUpdate", ttInRule, ttOutRule, rcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv3ResolverUpdate)
				if err := _Routerv3.contract.UnpackLog(event, "ResolverUpdate", log); err != nil {
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

// ParseResolverUpdate is a log parse operation binding the contract event 0x9081d4c394cb22ad7d47837e87c89a58d1051fdf26813daa3f1c928c3ca16b30.
//
// Solidity: event ResolverUpdate(uint8 indexed ttIn, uint8 indexed ttOut, uint8 indexed rc)
func (_Routerv3 *Routerv3Filterer) ParseResolverUpdate(log types.Log) (*Routerv3ResolverUpdate, error) {
	event := new(Routerv3ResolverUpdate)
	if err := _Routerv3.contract.UnpackLog(event, "ResolverUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv3RouterComponentUpdateIterator is returned from FilterRouterComponentUpdate and is used to iterate over the raw logs and unpacked data for RouterComponentUpdate events raised by the Routerv3 contract.
type Routerv3RouterComponentUpdateIterator struct {
	Event *Routerv3RouterComponentUpdate // Event containing the contract specifics and raw log

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
func (it *Routerv3RouterComponentUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv3RouterComponentUpdate)
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
		it.Event = new(Routerv3RouterComponentUpdate)
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
func (it *Routerv3RouterComponentUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv3RouterComponentUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv3RouterComponentUpdate represents a RouterComponentUpdate event raised by the Routerv3 contract.
type Routerv3RouterComponentUpdate struct {
	Arg0    uint8
	Arg1    common.Address
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRouterComponentUpdate is a free log retrieval operation binding the contract event 0x154abede065a65ed31601ad4a386b9743487d106f61887037d935f6ce1c0f144.
//
// Solidity: event RouterComponentUpdate(uint8 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv3 *Routerv3Filterer) FilterRouterComponentUpdate(opts *bind.FilterOpts, arg0 []uint8, arg1 []common.Address) (*Routerv3RouterComponentUpdateIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Routerv3.contract.FilterLogs(opts, "RouterComponentUpdate", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return &Routerv3RouterComponentUpdateIterator{contract: _Routerv3.contract, event: "RouterComponentUpdate", logs: logs, sub: sub}, nil
}

// WatchRouterComponentUpdate is a free log subscription operation binding the contract event 0x154abede065a65ed31601ad4a386b9743487d106f61887037d935f6ce1c0f144.
//
// Solidity: event RouterComponentUpdate(uint8 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv3 *Routerv3Filterer) WatchRouterComponentUpdate(opts *bind.WatchOpts, sink chan<- *Routerv3RouterComponentUpdate, arg0 []uint8, arg1 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}
	var arg1Rule []interface{}
	for _, arg1Item := range arg1 {
		arg1Rule = append(arg1Rule, arg1Item)
	}

	logs, sub, err := _Routerv3.contract.WatchLogs(opts, "RouterComponentUpdate", arg0Rule, arg1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv3RouterComponentUpdate)
				if err := _Routerv3.contract.UnpackLog(event, "RouterComponentUpdate", log); err != nil {
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

// ParseRouterComponentUpdate is a log parse operation binding the contract event 0x154abede065a65ed31601ad4a386b9743487d106f61887037d935f6ce1c0f144.
//
// Solidity: event RouterComponentUpdate(uint8 indexed arg0, address indexed arg1, uint256 version)
func (_Routerv3 *Routerv3Filterer) ParseRouterComponentUpdate(log types.Log) (*Routerv3RouterComponentUpdate, error) {
	event := new(Routerv3RouterComponentUpdate)
	if err := _Routerv3.contract.UnpackLog(event, "RouterComponentUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv3SetFutureRouterIterator is returned from FilterSetFutureRouter and is used to iterate over the raw logs and unpacked data for SetFutureRouter events raised by the Routerv3 contract.
type Routerv3SetFutureRouterIterator struct {
	Event *Routerv3SetFutureRouter // Event containing the contract specifics and raw log

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
func (it *Routerv3SetFutureRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv3SetFutureRouter)
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
		it.Event = new(Routerv3SetFutureRouter)
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
func (it *Routerv3SetFutureRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv3SetFutureRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv3SetFutureRouter represents a SetFutureRouter event raised by the Routerv3 contract.
type Routerv3SetFutureRouter struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetFutureRouter is a free log retrieval operation binding the contract event 0x38ee79447c54ed2235ae0312a2a622f96c8dcb3ba5b20ceb62cd62edeb19ee03.
//
// Solidity: event SetFutureRouter(address indexed arg0)
func (_Routerv3 *Routerv3Filterer) FilterSetFutureRouter(opts *bind.FilterOpts, arg0 []common.Address) (*Routerv3SetFutureRouterIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Routerv3.contract.FilterLogs(opts, "SetFutureRouter", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &Routerv3SetFutureRouterIterator{contract: _Routerv3.contract, event: "SetFutureRouter", logs: logs, sub: sub}, nil
}

// WatchSetFutureRouter is a free log subscription operation binding the contract event 0x38ee79447c54ed2235ae0312a2a622f96c8dcb3ba5b20ceb62cd62edeb19ee03.
//
// Solidity: event SetFutureRouter(address indexed arg0)
func (_Routerv3 *Routerv3Filterer) WatchSetFutureRouter(opts *bind.WatchOpts, sink chan<- *Routerv3SetFutureRouter, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Routerv3.contract.WatchLogs(opts, "SetFutureRouter", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv3SetFutureRouter)
				if err := _Routerv3.contract.UnpackLog(event, "SetFutureRouter", log); err != nil {
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
func (_Routerv3 *Routerv3Filterer) ParseSetFutureRouter(log types.Log) (*Routerv3SetFutureRouter, error) {
	event := new(Routerv3SetFutureRouter)
	if err := _Routerv3.contract.UnpackLog(event, "SetFutureRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Routerv3TokenTypeUpdateIterator is returned from FilterTokenTypeUpdate and is used to iterate over the raw logs and unpacked data for TokenTypeUpdate events raised by the Routerv3 contract.
type Routerv3TokenTypeUpdateIterator struct {
	Event *Routerv3TokenTypeUpdate // Event containing the contract specifics and raw log

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
func (it *Routerv3TokenTypeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv3TokenTypeUpdate)
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
		it.Event = new(Routerv3TokenTypeUpdate)
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
func (it *Routerv3TokenTypeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv3TokenTypeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv3TokenTypeUpdate represents a TokenTypeUpdate event raised by the Routerv3 contract.
type Routerv3TokenTypeUpdate struct {
	TokenAddress common.Address
	Tt           uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenTypeUpdate is a free log retrieval operation binding the contract event 0xfcf5808e948444a9f9a56552e9f4d2436bf3f547440c412c27d3556296b308cf.
//
// Solidity: event TokenTypeUpdate(address indexed tokenAddress, uint8 indexed tt)
func (_Routerv3 *Routerv3Filterer) FilterTokenTypeUpdate(opts *bind.FilterOpts, tokenAddress []common.Address, tt []uint8) (*Routerv3TokenTypeUpdateIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var ttRule []interface{}
	for _, ttItem := range tt {
		ttRule = append(ttRule, ttItem)
	}

	logs, sub, err := _Routerv3.contract.FilterLogs(opts, "TokenTypeUpdate", tokenAddressRule, ttRule)
	if err != nil {
		return nil, err
	}
	return &Routerv3TokenTypeUpdateIterator{contract: _Routerv3.contract, event: "TokenTypeUpdate", logs: logs, sub: sub}, nil
}

// WatchTokenTypeUpdate is a free log subscription operation binding the contract event 0xfcf5808e948444a9f9a56552e9f4d2436bf3f547440c412c27d3556296b308cf.
//
// Solidity: event TokenTypeUpdate(address indexed tokenAddress, uint8 indexed tt)
func (_Routerv3 *Routerv3Filterer) WatchTokenTypeUpdate(opts *bind.WatchOpts, sink chan<- *Routerv3TokenTypeUpdate, tokenAddress []common.Address, tt []uint8) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var ttRule []interface{}
	for _, ttItem := range tt {
		ttRule = append(ttRule, ttItem)
	}

	logs, sub, err := _Routerv3.contract.WatchLogs(opts, "TokenTypeUpdate", tokenAddressRule, ttRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv3TokenTypeUpdate)
				if err := _Routerv3.contract.UnpackLog(event, "TokenTypeUpdate", log); err != nil {
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

// ParseTokenTypeUpdate is a log parse operation binding the contract event 0xfcf5808e948444a9f9a56552e9f4d2436bf3f547440c412c27d3556296b308cf.
//
// Solidity: event TokenTypeUpdate(address indexed tokenAddress, uint8 indexed tt)
func (_Routerv3 *Routerv3Filterer) ParseTokenTypeUpdate(log types.Log) (*Routerv3TokenTypeUpdate, error) {
	event := new(Routerv3TokenTypeUpdate)
	if err := _Routerv3.contract.UnpackLog(event, "TokenTypeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
