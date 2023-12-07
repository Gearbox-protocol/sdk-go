// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundv2ERC20

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

// Compoundv2ERC20MetaData contains all meta data concerning the Compoundv2ERC20 contract.
var Compoundv2ERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_cToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cTokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemUnderlying\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"CTokenError\",\"inputs\":[{\"name\":\"errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Compoundv2ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Compoundv2ERC20MetaData.ABI instead.
var Compoundv2ERC20ABI = Compoundv2ERC20MetaData.ABI

// Compoundv2ERC20 is an auto generated Go binding around an Ethereum contract.
type Compoundv2ERC20 struct {
	Compoundv2ERC20Caller     // Read-only binding to the contract
	Compoundv2ERC20Transactor // Write-only binding to the contract
	Compoundv2ERC20Filterer   // Log filterer for contract events
}

// Compoundv2ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Compoundv2ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Compoundv2ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Compoundv2ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Compoundv2ERC20Session struct {
	Contract     *Compoundv2ERC20  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Compoundv2ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Compoundv2ERC20CallerSession struct {
	Contract *Compoundv2ERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Compoundv2ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Compoundv2ERC20TransactorSession struct {
	Contract     *Compoundv2ERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Compoundv2ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Compoundv2ERC20Raw struct {
	Contract *Compoundv2ERC20 // Generic contract binding to access the raw methods on
}

// Compoundv2ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Compoundv2ERC20CallerRaw struct {
	Contract *Compoundv2ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// Compoundv2ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Compoundv2ERC20TransactorRaw struct {
	Contract *Compoundv2ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundv2ERC20 creates a new instance of Compoundv2ERC20, bound to a specific deployed contract.
func NewCompoundv2ERC20(address common.Address, backend bind.ContractBackend) (*Compoundv2ERC20, error) {
	contract, err := bindCompoundv2ERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compoundv2ERC20{Compoundv2ERC20Caller: Compoundv2ERC20Caller{contract: contract}, Compoundv2ERC20Transactor: Compoundv2ERC20Transactor{contract: contract}, Compoundv2ERC20Filterer: Compoundv2ERC20Filterer{contract: contract}}, nil
}

// NewCompoundv2ERC20Caller creates a new read-only instance of Compoundv2ERC20, bound to a specific deployed contract.
func NewCompoundv2ERC20Caller(address common.Address, caller bind.ContractCaller) (*Compoundv2ERC20Caller, error) {
	contract, err := bindCompoundv2ERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2ERC20Caller{contract: contract}, nil
}

// NewCompoundv2ERC20Transactor creates a new write-only instance of Compoundv2ERC20, bound to a specific deployed contract.
func NewCompoundv2ERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*Compoundv2ERC20Transactor, error) {
	contract, err := bindCompoundv2ERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2ERC20Transactor{contract: contract}, nil
}

// NewCompoundv2ERC20Filterer creates a new log filterer instance of Compoundv2ERC20, bound to a specific deployed contract.
func NewCompoundv2ERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*Compoundv2ERC20Filterer, error) {
	contract, err := bindCompoundv2ERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Compoundv2ERC20Filterer{contract: contract}, nil
}

// bindCompoundv2ERC20 binds a generic wrapper to an already deployed contract.
func bindCompoundv2ERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Compoundv2ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2ERC20 *Compoundv2ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2ERC20.Contract.Compoundv2ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2ERC20 *Compoundv2ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Compoundv2ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2ERC20 *Compoundv2ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Compoundv2ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2ERC20 *Compoundv2ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) GearboxAdapterType() (uint8, error) {
	return _Compoundv2ERC20.Contract.GearboxAdapterType(&_Compoundv2ERC20.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) GearboxAdapterType() (uint8, error) {
	return _Compoundv2ERC20.Contract.GearboxAdapterType(&_Compoundv2ERC20.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2ERC20.Contract.GearboxAdapterVersion(&_Compoundv2ERC20.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2ERC20.Contract.GearboxAdapterVersion(&_Compoundv2ERC20.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) Acl() (common.Address, error) {
	return _Compoundv2ERC20.Contract.Acl(&_Compoundv2ERC20.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) Acl() (common.Address, error) {
	return _Compoundv2ERC20.Contract.Acl(&_Compoundv2ERC20.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) AddressProvider() (common.Address, error) {
	return _Compoundv2ERC20.Contract.AddressProvider(&_Compoundv2ERC20.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) AddressProvider() (common.Address, error) {
	return _Compoundv2ERC20.Contract.AddressProvider(&_Compoundv2ERC20.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) CToken() (common.Address, error) {
	return _Compoundv2ERC20.Contract.CToken(&_Compoundv2ERC20.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) CToken() (common.Address, error) {
	return _Compoundv2ERC20.Contract.CToken(&_Compoundv2ERC20.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) CTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "cTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) CTokenMask() (*big.Int, error) {
	return _Compoundv2ERC20.Contract.CTokenMask(&_Compoundv2ERC20.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) CTokenMask() (*big.Int, error) {
	return _Compoundv2ERC20.Contract.CTokenMask(&_Compoundv2ERC20.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) CreditManager() (common.Address, error) {
	return _Compoundv2ERC20.Contract.CreditManager(&_Compoundv2ERC20.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) CreditManager() (common.Address, error) {
	return _Compoundv2ERC20.Contract.CreditManager(&_Compoundv2ERC20.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) TargetContract() (common.Address, error) {
	return _Compoundv2ERC20.Contract.TargetContract(&_Compoundv2ERC20.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) TargetContract() (common.Address, error) {
	return _Compoundv2ERC20.Contract.TargetContract(&_Compoundv2ERC20.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) TokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "tokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) TokenMask() (*big.Int, error) {
	return _Compoundv2ERC20.Contract.TokenMask(&_Compoundv2ERC20.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) TokenMask() (*big.Int, error) {
	return _Compoundv2ERC20.Contract.TokenMask(&_Compoundv2ERC20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2ERC20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) Underlying() (common.Address, error) {
	return _Compoundv2ERC20.Contract.Underlying(&_Compoundv2ERC20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2ERC20 *Compoundv2ERC20CallerSession) Underlying() (common.Address, error) {
	return _Compoundv2ERC20.Contract.Underlying(&_Compoundv2ERC20.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Mint(&_Compoundv2ERC20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Mint(&_Compoundv2ERC20.TransactOpts, amount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Transactor) MintDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.contract.Transact(opts, "mintDiff", leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.MintDiff(&_Compoundv2ERC20.TransactOpts, leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorSession) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.MintDiff(&_Compoundv2ERC20.TransactOpts, leftoverAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Transactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Redeem(&_Compoundv2ERC20.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.Redeem(&_Compoundv2ERC20.TransactOpts, amount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Transactor) RedeemDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.contract.Transact(opts, "redeemDiff", leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.RedeemDiff(&_Compoundv2ERC20.TransactOpts, leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.RedeemDiff(&_Compoundv2ERC20.TransactOpts, leftoverAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Transactor) RedeemUnderlying(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.contract.Transact(opts, "redeemUnderlying", amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20Session) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.RedeemUnderlying(&_Compoundv2ERC20.TransactOpts, amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2ERC20 *Compoundv2ERC20TransactorSession) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2ERC20.Contract.RedeemUnderlying(&_Compoundv2ERC20.TransactOpts, amount)
}
