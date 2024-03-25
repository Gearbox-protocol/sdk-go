// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundv2CToken

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

// Compoundv2CTokenMetaData contains all meta data concerning the Compoundv2CToken contract.
var Compoundv2CTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cTokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemUnderlying\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"CTokenError\",\"inputs\":[{\"name\":\"errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Compoundv2CTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Compoundv2CTokenMetaData.ABI instead.
var Compoundv2CTokenABI = Compoundv2CTokenMetaData.ABI

// Compoundv2CToken is an auto generated Go binding around an Ethereum contract.
type Compoundv2CToken struct {
	Compoundv2CTokenCaller     // Read-only binding to the contract
	Compoundv2CTokenTransactor // Write-only binding to the contract
	Compoundv2CTokenFilterer   // Log filterer for contract events
}

// Compoundv2CTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Compoundv2CTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Compoundv2CTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Compoundv2CTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Compoundv2CTokenSession struct {
	Contract     *Compoundv2CToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Compoundv2CTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Compoundv2CTokenCallerSession struct {
	Contract *Compoundv2CTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Compoundv2CTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Compoundv2CTokenTransactorSession struct {
	Contract     *Compoundv2CTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Compoundv2CTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Compoundv2CTokenRaw struct {
	Contract *Compoundv2CToken // Generic contract binding to access the raw methods on
}

// Compoundv2CTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Compoundv2CTokenCallerRaw struct {
	Contract *Compoundv2CTokenCaller // Generic read-only contract binding to access the raw methods on
}

// Compoundv2CTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Compoundv2CTokenTransactorRaw struct {
	Contract *Compoundv2CTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundv2CToken creates a new instance of Compoundv2CToken, bound to a specific deployed contract.
func NewCompoundv2CToken(address common.Address, backend bind.ContractBackend) (*Compoundv2CToken, error) {
	contract, err := bindCompoundv2CToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CToken{Compoundv2CTokenCaller: Compoundv2CTokenCaller{contract: contract}, Compoundv2CTokenTransactor: Compoundv2CTokenTransactor{contract: contract}, Compoundv2CTokenFilterer: Compoundv2CTokenFilterer{contract: contract}}, nil
}

// NewCompoundv2CTokenCaller creates a new read-only instance of Compoundv2CToken, bound to a specific deployed contract.
func NewCompoundv2CTokenCaller(address common.Address, caller bind.ContractCaller) (*Compoundv2CTokenCaller, error) {
	contract, err := bindCompoundv2CToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CTokenCaller{contract: contract}, nil
}

// NewCompoundv2CTokenTransactor creates a new write-only instance of Compoundv2CToken, bound to a specific deployed contract.
func NewCompoundv2CTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Compoundv2CTokenTransactor, error) {
	contract, err := bindCompoundv2CToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CTokenTransactor{contract: contract}, nil
}

// NewCompoundv2CTokenFilterer creates a new log filterer instance of Compoundv2CToken, bound to a specific deployed contract.
func NewCompoundv2CTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Compoundv2CTokenFilterer, error) {
	contract, err := bindCompoundv2CToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CTokenFilterer{contract: contract}, nil
}

// bindCompoundv2CToken binds a generic wrapper to an already deployed contract.
func bindCompoundv2CToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Compoundv2CTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2CToken *Compoundv2CTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2CToken.Contract.Compoundv2CTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2CToken *Compoundv2CTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Compoundv2CTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2CToken *Compoundv2CTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Compoundv2CTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2CToken *Compoundv2CTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2CToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2CToken *Compoundv2CTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2CToken *Compoundv2CTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CToken *Compoundv2CTokenCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CToken *Compoundv2CTokenSession) GearboxAdapterType() (uint8, error) {
	return _Compoundv2CToken.Contract.GearboxAdapterType(&_Compoundv2CToken.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) GearboxAdapterType() (uint8, error) {
	return _Compoundv2CToken.Contract.GearboxAdapterType(&_Compoundv2CToken.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CToken *Compoundv2CTokenCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CToken *Compoundv2CTokenSession) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2CToken.Contract.GearboxAdapterVersion(&_Compoundv2CToken.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2CToken.Contract.GearboxAdapterVersion(&_Compoundv2CToken.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) Acl() (common.Address, error) {
	return _Compoundv2CToken.Contract.Acl(&_Compoundv2CToken.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) Acl() (common.Address, error) {
	return _Compoundv2CToken.Contract.Acl(&_Compoundv2CToken.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) AddressProvider() (common.Address, error) {
	return _Compoundv2CToken.Contract.AddressProvider(&_Compoundv2CToken.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) AddressProvider() (common.Address, error) {
	return _Compoundv2CToken.Contract.AddressProvider(&_Compoundv2CToken.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) CToken() (common.Address, error) {
	return _Compoundv2CToken.Contract.CToken(&_Compoundv2CToken.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) CToken() (common.Address, error) {
	return _Compoundv2CToken.Contract.CToken(&_Compoundv2CToken.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenCaller) CTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "cTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenSession) CTokenMask() (*big.Int, error) {
	return _Compoundv2CToken.Contract.CTokenMask(&_Compoundv2CToken.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) CTokenMask() (*big.Int, error) {
	return _Compoundv2CToken.Contract.CTokenMask(&_Compoundv2CToken.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) CreditManager() (common.Address, error) {
	return _Compoundv2CToken.Contract.CreditManager(&_Compoundv2CToken.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) CreditManager() (common.Address, error) {
	return _Compoundv2CToken.Contract.CreditManager(&_Compoundv2CToken.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) TargetContract() (common.Address, error) {
	return _Compoundv2CToken.Contract.TargetContract(&_Compoundv2CToken.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) TargetContract() (common.Address, error) {
	return _Compoundv2CToken.Contract.TargetContract(&_Compoundv2CToken.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenCaller) TokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "tokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenSession) TokenMask() (*big.Int, error) {
	return _Compoundv2CToken.Contract.TokenMask(&_Compoundv2CToken.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) TokenMask() (*big.Int, error) {
	return _Compoundv2CToken.Contract.TokenMask(&_Compoundv2CToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenSession) Underlying() (common.Address, error) {
	return _Compoundv2CToken.Contract.Underlying(&_Compoundv2CToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CToken *Compoundv2CTokenCallerSession) Underlying() (common.Address, error) {
	return _Compoundv2CToken.Contract.Underlying(&_Compoundv2CToken.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Mint(&_Compoundv2CToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Mint(&_Compoundv2CToken.TransactOpts, amount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactor) MintDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.contract.Transact(opts, "mintDiff", leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenSession) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.MintDiff(&_Compoundv2CToken.TransactOpts, leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactorSession) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.MintDiff(&_Compoundv2CToken.TransactOpts, leftoverAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Redeem(&_Compoundv2CToken.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.Redeem(&_Compoundv2CToken.TransactOpts, amount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactor) RedeemDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.contract.Transact(opts, "redeemDiff", leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.RedeemDiff(&_Compoundv2CToken.TransactOpts, leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactorSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.RedeemDiff(&_Compoundv2CToken.TransactOpts, leftoverAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.contract.Transact(opts, "redeemUnderlying", amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenSession) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.RedeemUnderlying(&_Compoundv2CToken.TransactOpts, amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CToken *Compoundv2CTokenTransactorSession) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CToken.Contract.RedeemUnderlying(&_Compoundv2CToken.TransactOpts, amount)
}
