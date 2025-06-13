// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalCreditSuite

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

// AdapterState is an auto generated low-level Go binding around an user-defined struct.
type AdapterState struct {
	BaseParams     BaseParams
	TargetContract common.Address
}

// BaseParams is an auto generated low-level Go binding around an user-defined struct.
type BaseParams struct {
	Addr             common.Address
	Version          *big.Int
	ContractType     [32]byte
	SerializedParams []byte
}

// BaseState is an auto generated low-level Go binding around an user-defined struct.
type BaseState struct {
	BaseParams BaseParams
}

// CollateralToken is an auto generated low-level Go binding around an user-defined struct.
type CollateralToken struct {
	Token                common.Address
	LiquidationThreshold uint16
}

// CreditFacadeState is an auto generated low-level Go binding around an user-defined struct.
type CreditFacadeState struct {
	BaseParams                BaseParams
	DegenNFT                  common.Address
	BotList                   common.Address
	Expirable                 bool
	ExpirationDate            *big.Int
	MaxDebtPerBlockMultiplier uint8
	MinDebt                   *big.Int
	MaxDebt                   *big.Int
	ForbiddenTokensMask       *big.Int
	IsPaused                  bool
}

// CreditManagerFilter is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerFilter struct {
	Configurators  []common.Address
	CreditManagers []common.Address
	Pools          []common.Address
	Underlying     common.Address
}

// CreditManagerState is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerState struct {
	BaseParams                 BaseParams
	Name                       string
	AccountFactory             common.Address
	Underlying                 common.Address
	Pool                       common.Address
	CreditFacade               common.Address
	CreditConfigurator         common.Address
	MaxEnabledTokens           uint8
	CollateralTokens           []CollateralToken
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}

// CreditSuiteData is an auto generated low-level Go binding around an user-defined struct.
type CreditSuiteData struct {
	CreditFacade       CreditFacadeState
	CreditManager      CreditManagerState
	CreditConfigurator BaseState
	AccountFactory     BaseState
	Adapters           []AdapterState
}

// GlobalCreditSuiteMetaData contains all meta data concerning the GlobalCreditSuite contract.
var GlobalCreditSuiteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressProvider_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"}],\"name\":\"getAccountFactoryState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getAdapters\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"internalType\":\"structAdapterState[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"}],\"name\":\"getCreditConfiguratorState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"}],\"name\":\"getCreditFacadeState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditFacadeState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getCreditManagerState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getCreditSuiteData\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditFacadeState\",\"name\":\"creditFacade\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerState\",\"name\":\"creditManager\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"creditConfigurator\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"accountFactory\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"internalType\":\"structAdapterState[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structCreditSuiteData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"configurators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"creditManagers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"internalType\":\"structCreditManagerFilter\",\"name\":\"filter\",\"type\":\"tuple\"}],\"name\":\"getCreditSuites\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokensMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditFacadeState\",\"name\":\"creditFacade\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerState\",\"name\":\"creditManager\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"creditConfigurator\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"accountFactory\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"internalType\":\"structAdapterState[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structCreditSuiteData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GlobalCreditSuiteABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalCreditSuiteMetaData.ABI instead.
var GlobalCreditSuiteABI = GlobalCreditSuiteMetaData.ABI

// GlobalCreditSuite is an auto generated Go binding around an Ethereum contract.
type GlobalCreditSuite struct {
	GlobalCreditSuiteCaller     // Read-only binding to the contract
	GlobalCreditSuiteTransactor // Write-only binding to the contract
	GlobalCreditSuiteFilterer   // Log filterer for contract events
}

// GlobalCreditSuiteCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalCreditSuiteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalCreditSuiteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalCreditSuiteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalCreditSuiteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalCreditSuiteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalCreditSuiteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalCreditSuiteSession struct {
	Contract     *GlobalCreditSuite // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GlobalCreditSuiteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalCreditSuiteCallerSession struct {
	Contract *GlobalCreditSuiteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GlobalCreditSuiteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalCreditSuiteTransactorSession struct {
	Contract     *GlobalCreditSuiteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GlobalCreditSuiteRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalCreditSuiteRaw struct {
	Contract *GlobalCreditSuite // Generic contract binding to access the raw methods on
}

// GlobalCreditSuiteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalCreditSuiteCallerRaw struct {
	Contract *GlobalCreditSuiteCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalCreditSuiteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalCreditSuiteTransactorRaw struct {
	Contract *GlobalCreditSuiteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalCreditSuite creates a new instance of GlobalCreditSuite, bound to a specific deployed contract.
func NewGlobalCreditSuite(address common.Address, backend bind.ContractBackend) (*GlobalCreditSuite, error) {
	contract, err := bindGlobalCreditSuite(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalCreditSuite{GlobalCreditSuiteCaller: GlobalCreditSuiteCaller{contract: contract}, GlobalCreditSuiteTransactor: GlobalCreditSuiteTransactor{contract: contract}, GlobalCreditSuiteFilterer: GlobalCreditSuiteFilterer{contract: contract}}, nil
}

// NewGlobalCreditSuiteCaller creates a new read-only instance of GlobalCreditSuite, bound to a specific deployed contract.
func NewGlobalCreditSuiteCaller(address common.Address, caller bind.ContractCaller) (*GlobalCreditSuiteCaller, error) {
	contract, err := bindGlobalCreditSuite(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalCreditSuiteCaller{contract: contract}, nil
}

// NewGlobalCreditSuiteTransactor creates a new write-only instance of GlobalCreditSuite, bound to a specific deployed contract.
func NewGlobalCreditSuiteTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalCreditSuiteTransactor, error) {
	contract, err := bindGlobalCreditSuite(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalCreditSuiteTransactor{contract: contract}, nil
}

// NewGlobalCreditSuiteFilterer creates a new log filterer instance of GlobalCreditSuite, bound to a specific deployed contract.
func NewGlobalCreditSuiteFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalCreditSuiteFilterer, error) {
	contract, err := bindGlobalCreditSuite(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalCreditSuiteFilterer{contract: contract}, nil
}

// bindGlobalCreditSuite binds a generic wrapper to an already deployed contract.
func bindGlobalCreditSuite(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlobalCreditSuiteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalCreditSuite *GlobalCreditSuiteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalCreditSuite.Contract.GlobalCreditSuiteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalCreditSuite *GlobalCreditSuiteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalCreditSuite.Contract.GlobalCreditSuiteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalCreditSuite *GlobalCreditSuiteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalCreditSuite.Contract.GlobalCreditSuiteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalCreditSuite *GlobalCreditSuiteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalCreditSuite.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalCreditSuite *GlobalCreditSuiteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalCreditSuite.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalCreditSuite *GlobalCreditSuiteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalCreditSuite.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) AddressProvider() (common.Address, error) {
	return _GlobalCreditSuite.Contract.AddressProvider(&_GlobalCreditSuite.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) AddressProvider() (common.Address, error) {
	return _GlobalCreditSuite.Contract.AddressProvider(&_GlobalCreditSuite.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) ContractType() ([32]byte, error) {
	return _GlobalCreditSuite.Contract.ContractType(&_GlobalCreditSuite.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) ContractType() ([32]byte, error) {
	return _GlobalCreditSuite.Contract.ContractType(&_GlobalCreditSuite.CallOpts)
}

// GetAccountFactoryState is a free data retrieval call binding the contract method 0x0d278d96.
//
// Solidity: function getAccountFactoryState(address accountFactory) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetAccountFactoryState(opts *bind.CallOpts, accountFactory common.Address) (BaseState, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getAccountFactoryState", accountFactory)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetAccountFactoryState is a free data retrieval call binding the contract method 0x0d278d96.
//
// Solidity: function getAccountFactoryState(address accountFactory) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetAccountFactoryState(accountFactory common.Address) (BaseState, error) {
	return _GlobalCreditSuite.Contract.GetAccountFactoryState(&_GlobalCreditSuite.CallOpts, accountFactory)
}

// GetAccountFactoryState is a free data retrieval call binding the contract method 0x0d278d96.
//
// Solidity: function getAccountFactoryState(address accountFactory) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetAccountFactoryState(accountFactory common.Address) (BaseState, error) {
	return _GlobalCreditSuite.Contract.GetAccountFactoryState(&_GlobalCreditSuite.CallOpts, accountFactory)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(((address,uint256,bytes32,bytes),address)[] adapters)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetAdapters(opts *bind.CallOpts, creditManager common.Address) ([]AdapterState, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getAdapters", creditManager)

	if err != nil {
		return *new([]AdapterState), err
	}

	out0 := *abi.ConvertType(out[0], new([]AdapterState)).(*[]AdapterState)

	return out0, err

}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(((address,uint256,bytes32,bytes),address)[] adapters)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetAdapters(creditManager common.Address) ([]AdapterState, error) {
	return _GlobalCreditSuite.Contract.GetAdapters(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetAdapters is a free data retrieval call binding the contract method 0x89e0dd3e.
//
// Solidity: function getAdapters(address creditManager) view returns(((address,uint256,bytes32,bytes),address)[] adapters)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetAdapters(creditManager common.Address) ([]AdapterState, error) {
	return _GlobalCreditSuite.Contract.GetAdapters(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetCreditConfiguratorState is a free data retrieval call binding the contract method 0xa35fe295.
//
// Solidity: function getCreditConfiguratorState(address creditConfigurator) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetCreditConfiguratorState(opts *bind.CallOpts, creditConfigurator common.Address) (BaseState, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getCreditConfiguratorState", creditConfigurator)

	if err != nil {
		return *new(BaseState), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseState)).(*BaseState)

	return out0, err

}

// GetCreditConfiguratorState is a free data retrieval call binding the contract method 0xa35fe295.
//
// Solidity: function getCreditConfiguratorState(address creditConfigurator) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetCreditConfiguratorState(creditConfigurator common.Address) (BaseState, error) {
	return _GlobalCreditSuite.Contract.GetCreditConfiguratorState(&_GlobalCreditSuite.CallOpts, creditConfigurator)
}

// GetCreditConfiguratorState is a free data retrieval call binding the contract method 0xa35fe295.
//
// Solidity: function getCreditConfiguratorState(address creditConfigurator) view returns(((address,uint256,bytes32,bytes)))
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetCreditConfiguratorState(creditConfigurator common.Address) (BaseState, error) {
	return _GlobalCreditSuite.Contract.GetCreditConfiguratorState(&_GlobalCreditSuite.CallOpts, creditConfigurator)
}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address creditFacade) view returns(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetCreditFacadeState(opts *bind.CallOpts, creditFacade common.Address) (CreditFacadeState, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getCreditFacadeState", creditFacade)

	if err != nil {
		return *new(CreditFacadeState), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditFacadeState)).(*CreditFacadeState)

	return out0, err

}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address creditFacade) view returns(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetCreditFacadeState(creditFacade common.Address) (CreditFacadeState, error) {
	return _GlobalCreditSuite.Contract.GetCreditFacadeState(&_GlobalCreditSuite.CallOpts, creditFacade)
}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address creditFacade) view returns(((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetCreditFacadeState(creditFacade common.Address) (CreditFacadeState, error) {
	return _GlobalCreditSuite.Contract.GetCreditFacadeState(&_GlobalCreditSuite.CallOpts, creditFacade)
}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address creditManager) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetCreditManagerState(opts *bind.CallOpts, creditManager common.Address) (CreditManagerState, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getCreditManagerState", creditManager)

	if err != nil {
		return *new(CreditManagerState), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditManagerState)).(*CreditManagerState)

	return out0, err

}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address creditManager) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetCreditManagerState(creditManager common.Address) (CreditManagerState, error) {
	return _GlobalCreditSuite.Contract.GetCreditManagerState(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address creditManager) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetCreditManagerState(creditManager common.Address) (CreditManagerState, error) {
	return _GlobalCreditSuite.Contract.GetCreditManagerState(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetCreditSuiteData(opts *bind.CallOpts, creditManager common.Address) (CreditSuiteData, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getCreditSuiteData", creditManager)

	if err != nil {
		return *new(CreditSuiteData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditSuiteData)).(*CreditSuiteData)

	return out0, err

}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetCreditSuiteData(creditManager common.Address) (CreditSuiteData, error) {
	return _GlobalCreditSuite.Contract.GetCreditSuiteData(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetCreditSuiteData(creditManager common.Address) (CreditSuiteData, error) {
	return _GlobalCreditSuite.Contract.GetCreditSuiteData(&_GlobalCreditSuite.CallOpts, creditManager)
}

// GetCreditSuites is a free data retrieval call binding the contract method 0x08740cb2.
//
// Solidity: function getCreditSuites((address[],address[],address[],address) filter) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[] result)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) GetCreditSuites(opts *bind.CallOpts, filter CreditManagerFilter) ([]CreditSuiteData, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "getCreditSuites", filter)

	if err != nil {
		return *new([]CreditSuiteData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditSuiteData)).(*[]CreditSuiteData)

	return out0, err

}

// GetCreditSuites is a free data retrieval call binding the contract method 0x08740cb2.
//
// Solidity: function getCreditSuites((address[],address[],address[],address) filter) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[] result)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) GetCreditSuites(filter CreditManagerFilter) ([]CreditSuiteData, error) {
	return _GlobalCreditSuite.Contract.GetCreditSuites(&_GlobalCreditSuite.CallOpts, filter)
}

// GetCreditSuites is a free data retrieval call binding the contract method 0x08740cb2.
//
// Solidity: function getCreditSuites((address[],address[],address[],address) filter) view returns((((address,uint256,bytes32,bytes),address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[])[] result)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) GetCreditSuites(filter CreditManagerFilter) ([]CreditSuiteData, error) {
	return _GlobalCreditSuite.Contract.GetCreditSuites(&_GlobalCreditSuite.CallOpts, filter)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalCreditSuite *GlobalCreditSuiteCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlobalCreditSuite.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalCreditSuite *GlobalCreditSuiteSession) Version() (*big.Int, error) {
	return _GlobalCreditSuite.Contract.Version(&_GlobalCreditSuite.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalCreditSuite *GlobalCreditSuiteCallerSession) Version() (*big.Int, error) {
	return _GlobalCreditSuite.Contract.Version(&_GlobalCreditSuite.CallOpts)
}
