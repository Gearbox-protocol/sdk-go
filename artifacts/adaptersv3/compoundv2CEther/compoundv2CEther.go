// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundv2CEther

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

// Compoundv2CEtherMetaData contains all meta data concerning the Compoundv2CEther contract.
var Compoundv2CEtherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_cethGateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cTokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemDiff\",\"inputs\":[{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemUnderlying\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMask\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlying\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"CTokenError\",\"inputs\":[{\"name\":\"errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Compoundv2CEtherABI is the input ABI used to generate the binding from.
// Deprecated: Use Compoundv2CEtherMetaData.ABI instead.
var Compoundv2CEtherABI = Compoundv2CEtherMetaData.ABI

// Compoundv2CEther is an auto generated Go binding around an Ethereum contract.
type Compoundv2CEther struct {
	Compoundv2CEtherCaller     // Read-only binding to the contract
	Compoundv2CEtherTransactor // Write-only binding to the contract
	Compoundv2CEtherFilterer   // Log filterer for contract events
}

// Compoundv2CEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type Compoundv2CEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Compoundv2CEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Compoundv2CEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Compoundv2CEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Compoundv2CEtherSession struct {
	Contract     *Compoundv2CEther // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Compoundv2CEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Compoundv2CEtherCallerSession struct {
	Contract *Compoundv2CEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Compoundv2CEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Compoundv2CEtherTransactorSession struct {
	Contract     *Compoundv2CEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Compoundv2CEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type Compoundv2CEtherRaw struct {
	Contract *Compoundv2CEther // Generic contract binding to access the raw methods on
}

// Compoundv2CEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Compoundv2CEtherCallerRaw struct {
	Contract *Compoundv2CEtherCaller // Generic read-only contract binding to access the raw methods on
}

// Compoundv2CEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Compoundv2CEtherTransactorRaw struct {
	Contract *Compoundv2CEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundv2CEther creates a new instance of Compoundv2CEther, bound to a specific deployed contract.
func NewCompoundv2CEther(address common.Address, backend bind.ContractBackend) (*Compoundv2CEther, error) {
	contract, err := bindCompoundv2CEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CEther{Compoundv2CEtherCaller: Compoundv2CEtherCaller{contract: contract}, Compoundv2CEtherTransactor: Compoundv2CEtherTransactor{contract: contract}, Compoundv2CEtherFilterer: Compoundv2CEtherFilterer{contract: contract}}, nil
}

// NewCompoundv2CEtherCaller creates a new read-only instance of Compoundv2CEther, bound to a specific deployed contract.
func NewCompoundv2CEtherCaller(address common.Address, caller bind.ContractCaller) (*Compoundv2CEtherCaller, error) {
	contract, err := bindCompoundv2CEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CEtherCaller{contract: contract}, nil
}

// NewCompoundv2CEtherTransactor creates a new write-only instance of Compoundv2CEther, bound to a specific deployed contract.
func NewCompoundv2CEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*Compoundv2CEtherTransactor, error) {
	contract, err := bindCompoundv2CEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CEtherTransactor{contract: contract}, nil
}

// NewCompoundv2CEtherFilterer creates a new log filterer instance of Compoundv2CEther, bound to a specific deployed contract.
func NewCompoundv2CEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*Compoundv2CEtherFilterer, error) {
	contract, err := bindCompoundv2CEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Compoundv2CEtherFilterer{contract: contract}, nil
}

// bindCompoundv2CEther binds a generic wrapper to an already deployed contract.
func bindCompoundv2CEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Compoundv2CEtherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2CEther *Compoundv2CEtherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2CEther.Contract.Compoundv2CEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2CEther *Compoundv2CEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Compoundv2CEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2CEther *Compoundv2CEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Compoundv2CEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compoundv2CEther *Compoundv2CEtherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compoundv2CEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compoundv2CEther *Compoundv2CEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compoundv2CEther *Compoundv2CEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CEther *Compoundv2CEtherCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CEther *Compoundv2CEtherSession) GearboxAdapterType() (uint8, error) {
	return _Compoundv2CEther.Contract.GearboxAdapterType(&_Compoundv2CEther.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) GearboxAdapterType() (uint8, error) {
	return _Compoundv2CEther.Contract.GearboxAdapterType(&_Compoundv2CEther.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CEther *Compoundv2CEtherCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CEther *Compoundv2CEtherSession) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2CEther.Contract.GearboxAdapterVersion(&_Compoundv2CEther.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Compoundv2CEther.Contract.GearboxAdapterVersion(&_Compoundv2CEther.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) Acl() (common.Address, error) {
	return _Compoundv2CEther.Contract.Acl(&_Compoundv2CEther.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) Acl() (common.Address, error) {
	return _Compoundv2CEther.Contract.Acl(&_Compoundv2CEther.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) AddressProvider() (common.Address, error) {
	return _Compoundv2CEther.Contract.AddressProvider(&_Compoundv2CEther.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) AddressProvider() (common.Address, error) {
	return _Compoundv2CEther.Contract.AddressProvider(&_Compoundv2CEther.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) CToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "cToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) CToken() (common.Address, error) {
	return _Compoundv2CEther.Contract.CToken(&_Compoundv2CEther.CallOpts)
}

// CToken is a free data retrieval call binding the contract method 0x69e527da.
//
// Solidity: function cToken() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) CToken() (common.Address, error) {
	return _Compoundv2CEther.Contract.CToken(&_Compoundv2CEther.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherCaller) CTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "cTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherSession) CTokenMask() (*big.Int, error) {
	return _Compoundv2CEther.Contract.CTokenMask(&_Compoundv2CEther.CallOpts)
}

// CTokenMask is a free data retrieval call binding the contract method 0xd990c5a6.
//
// Solidity: function cTokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) CTokenMask() (*big.Int, error) {
	return _Compoundv2CEther.Contract.CTokenMask(&_Compoundv2CEther.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) CreditManager() (common.Address, error) {
	return _Compoundv2CEther.Contract.CreditManager(&_Compoundv2CEther.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) CreditManager() (common.Address, error) {
	return _Compoundv2CEther.Contract.CreditManager(&_Compoundv2CEther.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) TargetContract() (common.Address, error) {
	return _Compoundv2CEther.Contract.TargetContract(&_Compoundv2CEther.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) TargetContract() (common.Address, error) {
	return _Compoundv2CEther.Contract.TargetContract(&_Compoundv2CEther.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherCaller) TokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "tokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherSession) TokenMask() (*big.Int, error) {
	return _Compoundv2CEther.Contract.TokenMask(&_Compoundv2CEther.CallOpts)
}

// TokenMask is a free data retrieval call binding the contract method 0x027bd577.
//
// Solidity: function tokenMask() view returns(uint256)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) TokenMask() (*big.Int, error) {
	return _Compoundv2CEther.Contract.TokenMask(&_Compoundv2CEther.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compoundv2CEther.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherSession) Underlying() (common.Address, error) {
	return _Compoundv2CEther.Contract.Underlying(&_Compoundv2CEther.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Compoundv2CEther *Compoundv2CEtherCallerSession) Underlying() (common.Address, error) {
	return _Compoundv2CEther.Contract.Underlying(&_Compoundv2CEther.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Mint(&_Compoundv2CEther.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Mint(&_Compoundv2CEther.TransactOpts, amount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactor) MintDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.contract.Transact(opts, "mintDiff", leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherSession) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.MintDiff(&_Compoundv2CEther.TransactOpts, leftoverAmount)
}

// MintDiff is a paid mutator transaction binding the contract method 0x9f2d31c5.
//
// Solidity: function mintDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactorSession) MintDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.MintDiff(&_Compoundv2CEther.TransactOpts, leftoverAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Redeem(&_Compoundv2CEther.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.Redeem(&_Compoundv2CEther.TransactOpts, amount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactor) RedeemDiff(opts *bind.TransactOpts, leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.contract.Transact(opts, "redeemDiff", leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.RedeemDiff(&_Compoundv2CEther.TransactOpts, leftoverAmount)
}

// RedeemDiff is a paid mutator transaction binding the contract method 0x0acb3202.
//
// Solidity: function redeemDiff(uint256 leftoverAmount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactorSession) RedeemDiff(leftoverAmount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.RedeemDiff(&_Compoundv2CEther.TransactOpts, leftoverAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactor) RedeemUnderlying(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.contract.Transact(opts, "redeemUnderlying", amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherSession) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.RedeemUnderlying(&_Compoundv2CEther.TransactOpts, amount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 amount) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Compoundv2CEther *Compoundv2CEtherTransactorSession) RedeemUnderlying(amount *big.Int) (*types.Transaction, error) {
	return _Compoundv2CEther.Contract.RedeemUnderlying(&_Compoundv2CEther.TransactOpts, amount)
}
