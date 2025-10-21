// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountMigrationBot

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

// MigratedCollateral is an auto generated low-level Go binding around an user-defined struct.
type MigratedCollateral struct {
	Collateral          common.Address
	Amount              *big.Int
	TargetQuotaIncrease *big.Int
	UnderlyingInSource  bool
	UnderlyingInTarget  bool
	PhantomTokenParams  PhantomTokenParams
}

// MigrationParams is an auto generated low-level Go binding around an user-defined struct.
type MigrationParams struct {
	AccountOwner            common.Address
	SourceCreditAccount     common.Address
	TargetCreditManager     common.Address
	MigratedCollaterals     []MigratedCollateral
	TargetBorrowAmount      *big.Int
	UnderlyingSwapCalls     []MultiCall
	ExtraOpeningCalls       []MultiCall
	UniqueTransferredTokens []common.Address
	NumAddCollateralCalls   *big.Int
	NumRemoveQuotasCalls    *big.Int
	NumIncreaseQuotaCalls   *big.Int
	NumPhantomTokenCalls    *big.Int
}

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// PhantomTokenOverride is an auto generated low-level Go binding around an user-defined struct.
type PhantomTokenOverride struct {
	NewToken           common.Address
	Underlying         common.Address
	WithdrawalCallData []byte
}

// PhantomTokenParams is an auto generated low-level Go binding around an user-defined struct.
type PhantomTokenParams struct {
	IsPhantomToken   bool
	Underlying       common.Address
	UnderlyingAmount *big.Int
}

// PriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type PriceUpdate struct {
	PriceFeed common.Address
	Data      []byte
}

// AccountMigrationBotMetaData contains all meta data concerning the AccountMigrationBot contract.
var AccountMigrationBotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mcFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ioProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractsRegisterOld\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ForceApproveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsRegisterOld\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mcFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceCreditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetCreditManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"targetQuotaIncrease\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"underlyingInSource\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"underlyingInTarget\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPhantomToken\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPhantomTokenParams\",\"name\":\"phantomTokenParams\",\"type\":\"tuple\"}],\"internalType\":\"structMigratedCollateral[]\",\"name\":\"migratedCollaterals\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"targetBorrowAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"underlyingSwapCalls\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"extraOpeningCalls\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"uniqueTransferredTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"numAddCollateralCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numRemoveQuotasCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numIncreaseQuotaCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPhantomTokenCalls\",\"type\":\"uint256\"}],\"internalType\":\"structMigrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceCreditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetCreditManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"targetQuotaIncrease\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"underlyingInSource\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"underlyingInTarget\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isPhantomToken\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPhantomTokenParams\",\"name\":\"phantomTokenParams\",\"type\":\"tuple\"}],\"internalType\":\"structMigratedCollateral[]\",\"name\":\"migratedCollaterals\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"targetBorrowAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"underlyingSwapCalls\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"extraOpeningCalls\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"uniqueTransferredTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"numAddCollateralCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numRemoveQuotasCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numIncreaseQuotaCalls\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numPhantomTokenCalls\",\"type\":\"uint256\"}],\"internalType\":\"structMigrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structPriceUpdate[]\",\"name\":\"priceUpdates\",\"type\":\"tuple[]\"}],\"name\":\"migrateCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"phantomToken\",\"type\":\"address\"}],\"name\":\"phantomTokenOverrides\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"withdrawalCallData\",\"type\":\"bytes\"}],\"internalType\":\"structPhantomTokenOverride\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredPermissions\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serialize\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"phantomToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"withdrawalCallData\",\"type\":\"bytes\"}],\"name\":\"setPhantomTokenOverride\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AccountMigrationBotABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMigrationBotMetaData.ABI instead.
var AccountMigrationBotABI = AccountMigrationBotMetaData.ABI

// AccountMigrationBot is an auto generated Go binding around an Ethereum contract.
type AccountMigrationBot struct {
	AccountMigrationBotCaller     // Read-only binding to the contract
	AccountMigrationBotTransactor // Write-only binding to the contract
	AccountMigrationBotFilterer   // Log filterer for contract events
}

// AccountMigrationBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountMigrationBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMigrationBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountMigrationBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMigrationBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountMigrationBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMigrationBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountMigrationBotSession struct {
	Contract     *AccountMigrationBot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AccountMigrationBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountMigrationBotCallerSession struct {
	Contract *AccountMigrationBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AccountMigrationBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountMigrationBotTransactorSession struct {
	Contract     *AccountMigrationBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AccountMigrationBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountMigrationBotRaw struct {
	Contract *AccountMigrationBot // Generic contract binding to access the raw methods on
}

// AccountMigrationBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountMigrationBotCallerRaw struct {
	Contract *AccountMigrationBotCaller // Generic read-only contract binding to access the raw methods on
}

// AccountMigrationBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountMigrationBotTransactorRaw struct {
	Contract *AccountMigrationBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountMigrationBot creates a new instance of AccountMigrationBot, bound to a specific deployed contract.
func NewAccountMigrationBot(address common.Address, backend bind.ContractBackend) (*AccountMigrationBot, error) {
	contract, err := bindAccountMigrationBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountMigrationBot{AccountMigrationBotCaller: AccountMigrationBotCaller{contract: contract}, AccountMigrationBotTransactor: AccountMigrationBotTransactor{contract: contract}, AccountMigrationBotFilterer: AccountMigrationBotFilterer{contract: contract}}, nil
}

// NewAccountMigrationBotCaller creates a new read-only instance of AccountMigrationBot, bound to a specific deployed contract.
func NewAccountMigrationBotCaller(address common.Address, caller bind.ContractCaller) (*AccountMigrationBotCaller, error) {
	contract, err := bindAccountMigrationBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMigrationBotCaller{contract: contract}, nil
}

// NewAccountMigrationBotTransactor creates a new write-only instance of AccountMigrationBot, bound to a specific deployed contract.
func NewAccountMigrationBotTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountMigrationBotTransactor, error) {
	contract, err := bindAccountMigrationBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMigrationBotTransactor{contract: contract}, nil
}

// NewAccountMigrationBotFilterer creates a new log filterer instance of AccountMigrationBot, bound to a specific deployed contract.
func NewAccountMigrationBotFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountMigrationBotFilterer, error) {
	contract, err := bindAccountMigrationBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountMigrationBotFilterer{contract: contract}, nil
}

// bindAccountMigrationBot binds a generic wrapper to an already deployed contract.
func bindAccountMigrationBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountMigrationBotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMigrationBot *AccountMigrationBotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountMigrationBot.Contract.AccountMigrationBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMigrationBot *AccountMigrationBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.AccountMigrationBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMigrationBot *AccountMigrationBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.AccountMigrationBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMigrationBot *AccountMigrationBotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountMigrationBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMigrationBot *AccountMigrationBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMigrationBot *AccountMigrationBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AccountMigrationBot *AccountMigrationBotCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AccountMigrationBot *AccountMigrationBotSession) ContractType() ([32]byte, error) {
	return _AccountMigrationBot.Contract.ContractType(&_AccountMigrationBot.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) ContractType() ([32]byte, error) {
	return _AccountMigrationBot.Contract.ContractType(&_AccountMigrationBot.CallOpts)
}

// ContractsRegisterOld is a free data retrieval call binding the contract method 0x47293039.
//
// Solidity: function contractsRegisterOld() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCaller) ContractsRegisterOld(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "contractsRegisterOld")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegisterOld is a free data retrieval call binding the contract method 0x47293039.
//
// Solidity: function contractsRegisterOld() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotSession) ContractsRegisterOld() (common.Address, error) {
	return _AccountMigrationBot.Contract.ContractsRegisterOld(&_AccountMigrationBot.CallOpts)
}

// ContractsRegisterOld is a free data retrieval call binding the contract method 0x47293039.
//
// Solidity: function contractsRegisterOld() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) ContractsRegisterOld() (common.Address, error) {
	return _AccountMigrationBot.Contract.ContractsRegisterOld(&_AccountMigrationBot.CallOpts)
}

// McFactory is a free data retrieval call binding the contract method 0xb7328aef.
//
// Solidity: function mcFactory() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCaller) McFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "mcFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// McFactory is a free data retrieval call binding the contract method 0xb7328aef.
//
// Solidity: function mcFactory() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotSession) McFactory() (common.Address, error) {
	return _AccountMigrationBot.Contract.McFactory(&_AccountMigrationBot.CallOpts)
}

// McFactory is a free data retrieval call binding the contract method 0xb7328aef.
//
// Solidity: function mcFactory() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) McFactory() (common.Address, error) {
	return _AccountMigrationBot.Contract.McFactory(&_AccountMigrationBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotSession) Owner() (common.Address, error) {
	return _AccountMigrationBot.Contract.Owner(&_AccountMigrationBot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) Owner() (common.Address, error) {
	return _AccountMigrationBot.Contract.Owner(&_AccountMigrationBot.CallOpts)
}

// PhantomTokenOverrides is a free data retrieval call binding the contract method 0xbfd00a9c.
//
// Solidity: function phantomTokenOverrides(address phantomToken) view returns((address,address,bytes))
func (_AccountMigrationBot *AccountMigrationBotCaller) PhantomTokenOverrides(opts *bind.CallOpts, phantomToken common.Address) (PhantomTokenOverride, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "phantomTokenOverrides", phantomToken)

	if err != nil {
		return *new(PhantomTokenOverride), err
	}

	out0 := *abi.ConvertType(out[0], new(PhantomTokenOverride)).(*PhantomTokenOverride)

	return out0, err

}

// PhantomTokenOverrides is a free data retrieval call binding the contract method 0xbfd00a9c.
//
// Solidity: function phantomTokenOverrides(address phantomToken) view returns((address,address,bytes))
func (_AccountMigrationBot *AccountMigrationBotSession) PhantomTokenOverrides(phantomToken common.Address) (PhantomTokenOverride, error) {
	return _AccountMigrationBot.Contract.PhantomTokenOverrides(&_AccountMigrationBot.CallOpts, phantomToken)
}

// PhantomTokenOverrides is a free data retrieval call binding the contract method 0xbfd00a9c.
//
// Solidity: function phantomTokenOverrides(address phantomToken) view returns((address,address,bytes))
func (_AccountMigrationBot *AccountMigrationBotCallerSession) PhantomTokenOverrides(phantomToken common.Address) (PhantomTokenOverride, error) {
	return _AccountMigrationBot.Contract.PhantomTokenOverrides(&_AccountMigrationBot.CallOpts, phantomToken)
}

// RequiredPermissions is a free data retrieval call binding the contract method 0x2e7ad41f.
//
// Solidity: function requiredPermissions() view returns(uint192)
func (_AccountMigrationBot *AccountMigrationBotCaller) RequiredPermissions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "requiredPermissions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredPermissions is a free data retrieval call binding the contract method 0x2e7ad41f.
//
// Solidity: function requiredPermissions() view returns(uint192)
func (_AccountMigrationBot *AccountMigrationBotSession) RequiredPermissions() (*big.Int, error) {
	return _AccountMigrationBot.Contract.RequiredPermissions(&_AccountMigrationBot.CallOpts)
}

// RequiredPermissions is a free data retrieval call binding the contract method 0x2e7ad41f.
//
// Solidity: function requiredPermissions() view returns(uint192)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) RequiredPermissions() (*big.Int, error) {
	return _AccountMigrationBot.Contract.RequiredPermissions(&_AccountMigrationBot.CallOpts)
}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_AccountMigrationBot *AccountMigrationBotCaller) Serialize(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "serialize")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_AccountMigrationBot *AccountMigrationBotSession) Serialize() ([]byte, error) {
	return _AccountMigrationBot.Contract.Serialize(&_AccountMigrationBot.CallOpts)
}

// Serialize is a free data retrieval call binding the contract method 0xbc8018b1.
//
// Solidity: function serialize() view returns(bytes)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) Serialize() ([]byte, error) {
	return _AccountMigrationBot.Contract.Serialize(&_AccountMigrationBot.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccountMigrationBot *AccountMigrationBotCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccountMigrationBot.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccountMigrationBot *AccountMigrationBotSession) Version() (*big.Int, error) {
	return _AccountMigrationBot.Contract.Version(&_AccountMigrationBot.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AccountMigrationBot *AccountMigrationBotCallerSession) Version() (*big.Int, error) {
	return _AccountMigrationBot.Contract.Version(&_AccountMigrationBot.CallOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x6b9c1e6b.
//
// Solidity: function migrate((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactor) Migrate(opts *bind.TransactOpts, params MigrationParams) (*types.Transaction, error) {
	return _AccountMigrationBot.contract.Transact(opts, "migrate", params)
}

// Migrate is a paid mutator transaction binding the contract method 0x6b9c1e6b.
//
// Solidity: function migrate((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params) returns()
func (_AccountMigrationBot *AccountMigrationBotSession) Migrate(params MigrationParams) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.Migrate(&_AccountMigrationBot.TransactOpts, params)
}

// Migrate is a paid mutator transaction binding the contract method 0x6b9c1e6b.
//
// Solidity: function migrate((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactorSession) Migrate(params MigrationParams) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.Migrate(&_AccountMigrationBot.TransactOpts, params)
}

// MigrateCreditAccount is a paid mutator transaction binding the contract method 0x029f21bc.
//
// Solidity: function migrateCreditAccount((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params, (address,bytes)[] priceUpdates) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactor) MigrateCreditAccount(opts *bind.TransactOpts, params MigrationParams, priceUpdates []PriceUpdate) (*types.Transaction, error) {
	return _AccountMigrationBot.contract.Transact(opts, "migrateCreditAccount", params, priceUpdates)
}

// MigrateCreditAccount is a paid mutator transaction binding the contract method 0x029f21bc.
//
// Solidity: function migrateCreditAccount((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params, (address,bytes)[] priceUpdates) returns()
func (_AccountMigrationBot *AccountMigrationBotSession) MigrateCreditAccount(params MigrationParams, priceUpdates []PriceUpdate) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.MigrateCreditAccount(&_AccountMigrationBot.TransactOpts, params, priceUpdates)
}

// MigrateCreditAccount is a paid mutator transaction binding the contract method 0x029f21bc.
//
// Solidity: function migrateCreditAccount((address,address,address,(address,uint256,uint96,bool,bool,(bool,address,uint256))[],uint256,(address,bytes)[],(address,bytes)[],address[],uint256,uint256,uint256,uint256) params, (address,bytes)[] priceUpdates) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactorSession) MigrateCreditAccount(params MigrationParams, priceUpdates []PriceUpdate) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.MigrateCreditAccount(&_AccountMigrationBot.TransactOpts, params, priceUpdates)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountMigrationBot *AccountMigrationBotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMigrationBot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountMigrationBot *AccountMigrationBotSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.RenounceOwnership(&_AccountMigrationBot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccountMigrationBot *AccountMigrationBotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.RenounceOwnership(&_AccountMigrationBot.TransactOpts)
}

// SetPhantomTokenOverride is a paid mutator transaction binding the contract method 0x5f4c9872.
//
// Solidity: function setPhantomTokenOverride(address phantomToken, address newToken, address underlying, bytes withdrawalCallData) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactor) SetPhantomTokenOverride(opts *bind.TransactOpts, phantomToken common.Address, newToken common.Address, underlying common.Address, withdrawalCallData []byte) (*types.Transaction, error) {
	return _AccountMigrationBot.contract.Transact(opts, "setPhantomTokenOverride", phantomToken, newToken, underlying, withdrawalCallData)
}

// SetPhantomTokenOverride is a paid mutator transaction binding the contract method 0x5f4c9872.
//
// Solidity: function setPhantomTokenOverride(address phantomToken, address newToken, address underlying, bytes withdrawalCallData) returns()
func (_AccountMigrationBot *AccountMigrationBotSession) SetPhantomTokenOverride(phantomToken common.Address, newToken common.Address, underlying common.Address, withdrawalCallData []byte) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.SetPhantomTokenOverride(&_AccountMigrationBot.TransactOpts, phantomToken, newToken, underlying, withdrawalCallData)
}

// SetPhantomTokenOverride is a paid mutator transaction binding the contract method 0x5f4c9872.
//
// Solidity: function setPhantomTokenOverride(address phantomToken, address newToken, address underlying, bytes withdrawalCallData) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactorSession) SetPhantomTokenOverride(phantomToken common.Address, newToken common.Address, underlying common.Address, withdrawalCallData []byte) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.SetPhantomTokenOverride(&_AccountMigrationBot.TransactOpts, phantomToken, newToken, underlying, withdrawalCallData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountMigrationBot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountMigrationBot *AccountMigrationBotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.TransferOwnership(&_AccountMigrationBot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccountMigrationBot *AccountMigrationBotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountMigrationBot.Contract.TransferOwnership(&_AccountMigrationBot.TransactOpts, newOwner)
}

// AccountMigrationBotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountMigrationBot contract.
type AccountMigrationBotOwnershipTransferredIterator struct {
	Event *AccountMigrationBotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountMigrationBotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountMigrationBotOwnershipTransferred)
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
		it.Event = new(AccountMigrationBotOwnershipTransferred)
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
func (it *AccountMigrationBotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountMigrationBotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountMigrationBotOwnershipTransferred represents a OwnershipTransferred event raised by the AccountMigrationBot contract.
type AccountMigrationBotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountMigrationBot *AccountMigrationBotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountMigrationBotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountMigrationBot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountMigrationBotOwnershipTransferredIterator{contract: _AccountMigrationBot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccountMigrationBot *AccountMigrationBotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountMigrationBotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccountMigrationBot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountMigrationBotOwnershipTransferred)
				if err := _AccountMigrationBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AccountMigrationBot *AccountMigrationBotFilterer) ParseOwnershipTransferred(log types.Log) (*AccountMigrationBotOwnershipTransferred, error) {
	event := new(AccountMigrationBotOwnershipTransferred)
	if err := _AccountMigrationBot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
