// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditSuiteCompressor

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
	CreditManager             common.Address
	DegenNFT                  common.Address
	BotList                   common.Address
	Expirable                 bool
	ExpirationDate            *big.Int
	MaxDebtPerBlockMultiplier uint8
	MinDebt                   *big.Int
	MaxDebt                   *big.Int
	ForbiddenTokenMask        *big.Int
	IsPaused                  bool
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
	Adapters           []AdapterState
}

// CreditSuiteCompressorMetaData contains all meta data concerning the CreditSuiteCompressor contract.
var CreditSuiteCompressorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cf\",\"type\":\"address\"}],\"name\":\"getCreditFacadeState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditFacadeState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cm\",\"type\":\"address\"}],\"name\":\"getCreditManagerState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"}],\"name\":\"getCreditSuiteData\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"botList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"uint40\",\"name\":\"expirationDate\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"maxDebtPerBlockMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"internalType\":\"structCreditFacadeState\",\"name\":\"creditFacade\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokens\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structCollateralToken[]\",\"name\":\"collateralTokens\",\"type\":\"tuple[]\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerState\",\"name\":\"creditManager\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"internalType\":\"structBaseState\",\"name\":\"creditConfigurator\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"serializedParams\",\"type\":\"bytes\"}],\"internalType\":\"structBaseParams\",\"name\":\"baseParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"internalType\":\"structAdapterState[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"}],\"internalType\":\"structCreditSuiteData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditSuiteCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditSuiteCompressorMetaData.ABI instead.
var CreditSuiteCompressorABI = CreditSuiteCompressorMetaData.ABI

// CreditSuiteCompressor is an auto generated Go binding around an Ethereum contract.
type CreditSuiteCompressor struct {
	CreditSuiteCompressorCaller     // Read-only binding to the contract
	CreditSuiteCompressorTransactor // Write-only binding to the contract
	CreditSuiteCompressorFilterer   // Log filterer for contract events
}

// CreditSuiteCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditSuiteCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSuiteCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditSuiteCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSuiteCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditSuiteCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSuiteCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditSuiteCompressorSession struct {
	Contract     *CreditSuiteCompressor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CreditSuiteCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditSuiteCompressorCallerSession struct {
	Contract *CreditSuiteCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CreditSuiteCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditSuiteCompressorTransactorSession struct {
	Contract     *CreditSuiteCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CreditSuiteCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditSuiteCompressorRaw struct {
	Contract *CreditSuiteCompressor // Generic contract binding to access the raw methods on
}

// CreditSuiteCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditSuiteCompressorCallerRaw struct {
	Contract *CreditSuiteCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// CreditSuiteCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditSuiteCompressorTransactorRaw struct {
	Contract *CreditSuiteCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditSuiteCompressor creates a new instance of CreditSuiteCompressor, bound to a specific deployed contract.
func NewCreditSuiteCompressor(address common.Address, backend bind.ContractBackend) (*CreditSuiteCompressor, error) {
	contract, err := bindCreditSuiteCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditSuiteCompressor{CreditSuiteCompressorCaller: CreditSuiteCompressorCaller{contract: contract}, CreditSuiteCompressorTransactor: CreditSuiteCompressorTransactor{contract: contract}, CreditSuiteCompressorFilterer: CreditSuiteCompressorFilterer{contract: contract}}, nil
}

// NewCreditSuiteCompressorCaller creates a new read-only instance of CreditSuiteCompressor, bound to a specific deployed contract.
func NewCreditSuiteCompressorCaller(address common.Address, caller bind.ContractCaller) (*CreditSuiteCompressorCaller, error) {
	contract, err := bindCreditSuiteCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditSuiteCompressorCaller{contract: contract}, nil
}

// NewCreditSuiteCompressorTransactor creates a new write-only instance of CreditSuiteCompressor, bound to a specific deployed contract.
func NewCreditSuiteCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditSuiteCompressorTransactor, error) {
	contract, err := bindCreditSuiteCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditSuiteCompressorTransactor{contract: contract}, nil
}

// NewCreditSuiteCompressorFilterer creates a new log filterer instance of CreditSuiteCompressor, bound to a specific deployed contract.
func NewCreditSuiteCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditSuiteCompressorFilterer, error) {
	contract, err := bindCreditSuiteCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditSuiteCompressorFilterer{contract: contract}, nil
}

// bindCreditSuiteCompressor binds a generic wrapper to an already deployed contract.
func bindCreditSuiteCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditSuiteCompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditSuiteCompressor *CreditSuiteCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditSuiteCompressor.Contract.CreditSuiteCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditSuiteCompressor *CreditSuiteCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditSuiteCompressor.Contract.CreditSuiteCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditSuiteCompressor *CreditSuiteCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditSuiteCompressor.Contract.CreditSuiteCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditSuiteCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditSuiteCompressor *CreditSuiteCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditSuiteCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditSuiteCompressor *CreditSuiteCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditSuiteCompressor.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditSuiteCompressor *CreditSuiteCompressorCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditSuiteCompressor.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditSuiteCompressor *CreditSuiteCompressorSession) ContractType() ([32]byte, error) {
	return _CreditSuiteCompressor.Contract.ContractType(&_CreditSuiteCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerSession) ContractType() ([32]byte, error) {
	return _CreditSuiteCompressor.Contract.ContractType(&_CreditSuiteCompressor.CallOpts)
}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address _cf) view returns(((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCaller) GetCreditFacadeState(opts *bind.CallOpts, _cf common.Address) (CreditFacadeState, error) {
	var out []interface{}
	err := _CreditSuiteCompressor.contract.Call(opts, &out, "getCreditFacadeState", _cf)

	if err != nil {
		return *new(CreditFacadeState), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditFacadeState)).(*CreditFacadeState)

	return out0, err

}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address _cf) view returns(((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorSession) GetCreditFacadeState(_cf common.Address) (CreditFacadeState, error) {
	return _CreditSuiteCompressor.Contract.GetCreditFacadeState(&_CreditSuiteCompressor.CallOpts, _cf)
}

// GetCreditFacadeState is a free data retrieval call binding the contract method 0x012f16eb.
//
// Solidity: function getCreditFacadeState(address _cf) view returns(((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerSession) GetCreditFacadeState(_cf common.Address) (CreditFacadeState, error) {
	return _CreditSuiteCompressor.Contract.GetCreditFacadeState(&_CreditSuiteCompressor.CallOpts, _cf)
}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address _cm) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCaller) GetCreditManagerState(opts *bind.CallOpts, _cm common.Address) (CreditManagerState, error) {
	var out []interface{}
	err := _CreditSuiteCompressor.contract.Call(opts, &out, "getCreditManagerState", _cm)

	if err != nil {
		return *new(CreditManagerState), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditManagerState)).(*CreditManagerState)

	return out0, err

}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address _cm) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorSession) GetCreditManagerState(_cm common.Address) (CreditManagerState, error) {
	return _CreditSuiteCompressor.Contract.GetCreditManagerState(&_CreditSuiteCompressor.CallOpts, _cm)
}

// GetCreditManagerState is a free data retrieval call binding the contract method 0xcff8e936.
//
// Solidity: function getCreditManagerState(address _cm) view returns(((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerSession) GetCreditManagerState(_cm common.Address) (CreditManagerState, error) {
	return _CreditSuiteCompressor.Contract.GetCreditManagerState(&_CreditSuiteCompressor.CallOpts, _cm)
}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCaller) GetCreditSuiteData(opts *bind.CallOpts, creditManager common.Address) (CreditSuiteData, error) {
	var out []interface{}
	err := _CreditSuiteCompressor.contract.Call(opts, &out, "getCreditSuiteData", creditManager)

	if err != nil {
		return *new(CreditSuiteData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditSuiteData)).(*CreditSuiteData)

	return out0, err

}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorSession) GetCreditSuiteData(creditManager common.Address) (CreditSuiteData, error) {
	return _CreditSuiteCompressor.Contract.GetCreditSuiteData(&_CreditSuiteCompressor.CallOpts, creditManager)
}

// GetCreditSuiteData is a free data retrieval call binding the contract method 0xdc91ae00.
//
// Solidity: function getCreditSuiteData(address creditManager) view returns((((address,uint256,bytes32,bytes),address,address,address,bool,uint40,uint8,uint256,uint256,uint256,bool),((address,uint256,bytes32,bytes),string,address,address,address,address,address,uint8,(address,uint16)[],uint16,uint16,uint16,uint16,uint16),((address,uint256,bytes32,bytes)),((address,uint256,bytes32,bytes),address)[]) result)
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerSession) GetCreditSuiteData(creditManager common.Address) (CreditSuiteData, error) {
	return _CreditSuiteCompressor.Contract.GetCreditSuiteData(&_CreditSuiteCompressor.CallOpts, creditManager)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditSuiteCompressor *CreditSuiteCompressorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditSuiteCompressor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditSuiteCompressor *CreditSuiteCompressorSession) Version() (*big.Int, error) {
	return _CreditSuiteCompressor.Contract.Version(&_CreditSuiteCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditSuiteCompressor *CreditSuiteCompressorCallerSession) Version() (*big.Int, error) {
	return _CreditSuiteCompressor.Contract.Version(&_CreditSuiteCompressor.CallOpts)
}
