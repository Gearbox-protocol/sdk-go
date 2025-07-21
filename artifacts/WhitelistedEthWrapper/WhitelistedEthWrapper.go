// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WhitelistedEthWrapper

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

// WhitelistedEthWrapperMetaData contains all meta data concerning the WhitelistedEthWrapper contract.
var WhitelistedEthWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETH_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wstETH_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stETH_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_DEPOSITOR_WHITELIST_STATUS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_DEPOSIT_WHITELIST_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"depositWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isDepositWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"depositWhitelist_\",\"type\":\"bool\"}],\"name\":\"setDepositWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setDepositorWhitelistStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wstETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WhitelistedEthWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistedEthWrapperMetaData.ABI instead.
var WhitelistedEthWrapperABI = WhitelistedEthWrapperMetaData.ABI

// WhitelistedEthWrapper is an auto generated Go binding around an Ethereum contract.
type WhitelistedEthWrapper struct {
	WhitelistedEthWrapperCaller     // Read-only binding to the contract
	WhitelistedEthWrapperTransactor // Write-only binding to the contract
	WhitelistedEthWrapperFilterer   // Log filterer for contract events
}

// WhitelistedEthWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistedEthWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedEthWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistedEthWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedEthWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistedEthWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistedEthWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistedEthWrapperSession struct {
	Contract     *WhitelistedEthWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WhitelistedEthWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistedEthWrapperCallerSession struct {
	Contract *WhitelistedEthWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WhitelistedEthWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistedEthWrapperTransactorSession struct {
	Contract     *WhitelistedEthWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WhitelistedEthWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistedEthWrapperRaw struct {
	Contract *WhitelistedEthWrapper // Generic contract binding to access the raw methods on
}

// WhitelistedEthWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistedEthWrapperCallerRaw struct {
	Contract *WhitelistedEthWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistedEthWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistedEthWrapperTransactorRaw struct {
	Contract *WhitelistedEthWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistedEthWrapper creates a new instance of WhitelistedEthWrapper, bound to a specific deployed contract.
func NewWhitelistedEthWrapper(address common.Address, backend bind.ContractBackend) (*WhitelistedEthWrapper, error) {
	contract, err := bindWhitelistedEthWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistedEthWrapper{WhitelistedEthWrapperCaller: WhitelistedEthWrapperCaller{contract: contract}, WhitelistedEthWrapperTransactor: WhitelistedEthWrapperTransactor{contract: contract}, WhitelistedEthWrapperFilterer: WhitelistedEthWrapperFilterer{contract: contract}}, nil
}

// NewWhitelistedEthWrapperCaller creates a new read-only instance of WhitelistedEthWrapper, bound to a specific deployed contract.
func NewWhitelistedEthWrapperCaller(address common.Address, caller bind.ContractCaller) (*WhitelistedEthWrapperCaller, error) {
	contract, err := bindWhitelistedEthWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedEthWrapperCaller{contract: contract}, nil
}

// NewWhitelistedEthWrapperTransactor creates a new write-only instance of WhitelistedEthWrapper, bound to a specific deployed contract.
func NewWhitelistedEthWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistedEthWrapperTransactor, error) {
	contract, err := bindWhitelistedEthWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistedEthWrapperTransactor{contract: contract}, nil
}

// NewWhitelistedEthWrapperFilterer creates a new log filterer instance of WhitelistedEthWrapper, bound to a specific deployed contract.
func NewWhitelistedEthWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistedEthWrapperFilterer, error) {
	contract, err := bindWhitelistedEthWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistedEthWrapperFilterer{contract: contract}, nil
}

// bindWhitelistedEthWrapper binds a generic wrapper to an already deployed contract.
func bindWhitelistedEthWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WhitelistedEthWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitelistedEthWrapper.Contract.WhitelistedEthWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.WhitelistedEthWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.WhitelistedEthWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitelistedEthWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.contract.Transact(opts, method, params...)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) ETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "ETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) ETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.ETH(&_WhitelistedEthWrapper.CallOpts)
}

// ETH is a free data retrieval call binding the contract method 0x8322fff2.
//
// Solidity: function ETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) ETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.ETH(&_WhitelistedEthWrapper.CallOpts)
}

// SETDEPOSITORWHITELISTSTATUSROLE is a free data retrieval call binding the contract method 0x5447d02c.
//
// Solidity: function SET_DEPOSITOR_WHITELIST_STATUS_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) SETDEPOSITORWHITELISTSTATUSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "SET_DEPOSITOR_WHITELIST_STATUS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETDEPOSITORWHITELISTSTATUSROLE is a free data retrieval call binding the contract method 0x5447d02c.
//
// Solidity: function SET_DEPOSITOR_WHITELIST_STATUS_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) SETDEPOSITORWHITELISTSTATUSROLE() ([32]byte, error) {
	return _WhitelistedEthWrapper.Contract.SETDEPOSITORWHITELISTSTATUSROLE(&_WhitelistedEthWrapper.CallOpts)
}

// SETDEPOSITORWHITELISTSTATUSROLE is a free data retrieval call binding the contract method 0x5447d02c.
//
// Solidity: function SET_DEPOSITOR_WHITELIST_STATUS_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) SETDEPOSITORWHITELISTSTATUSROLE() ([32]byte, error) {
	return _WhitelistedEthWrapper.Contract.SETDEPOSITORWHITELISTSTATUSROLE(&_WhitelistedEthWrapper.CallOpts)
}

// SETDEPOSITWHITELISTROLE is a free data retrieval call binding the contract method 0x70d2a35a.
//
// Solidity: function SET_DEPOSIT_WHITELIST_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) SETDEPOSITWHITELISTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "SET_DEPOSIT_WHITELIST_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETDEPOSITWHITELISTROLE is a free data retrieval call binding the contract method 0x70d2a35a.
//
// Solidity: function SET_DEPOSIT_WHITELIST_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) SETDEPOSITWHITELISTROLE() ([32]byte, error) {
	return _WhitelistedEthWrapper.Contract.SETDEPOSITWHITELISTROLE(&_WhitelistedEthWrapper.CallOpts)
}

// SETDEPOSITWHITELISTROLE is a free data retrieval call binding the contract method 0x70d2a35a.
//
// Solidity: function SET_DEPOSIT_WHITELIST_ROLE() view returns(bytes32)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) SETDEPOSITWHITELISTROLE() ([32]byte, error) {
	return _WhitelistedEthWrapper.Contract.SETDEPOSITWHITELISTROLE(&_WhitelistedEthWrapper.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) WETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.WETH(&_WhitelistedEthWrapper.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) WETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.WETH(&_WhitelistedEthWrapper.CallOpts)
}

// DepositWhitelist is a free data retrieval call binding the contract method 0xd1bc1d73.
//
// Solidity: function depositWhitelist(address vault) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) DepositWhitelist(opts *bind.CallOpts, vault common.Address) (bool, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "depositWhitelist", vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositWhitelist is a free data retrieval call binding the contract method 0xd1bc1d73.
//
// Solidity: function depositWhitelist(address vault) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) DepositWhitelist(vault common.Address) (bool, error) {
	return _WhitelistedEthWrapper.Contract.DepositWhitelist(&_WhitelistedEthWrapper.CallOpts, vault)
}

// DepositWhitelist is a free data retrieval call binding the contract method 0xd1bc1d73.
//
// Solidity: function depositWhitelist(address vault) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) DepositWhitelist(vault common.Address) (bool, error) {
	return _WhitelistedEthWrapper.Contract.DepositWhitelist(&_WhitelistedEthWrapper.CallOpts, vault)
}

// IsDepositWhitelist is a free data retrieval call binding the contract method 0xb09d93fc.
//
// Solidity: function isDepositWhitelist(address vault, address account) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) IsDepositWhitelist(opts *bind.CallOpts, vault common.Address, account common.Address) (bool, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "isDepositWhitelist", vault, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositWhitelist is a free data retrieval call binding the contract method 0xb09d93fc.
//
// Solidity: function isDepositWhitelist(address vault, address account) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) IsDepositWhitelist(vault common.Address, account common.Address) (bool, error) {
	return _WhitelistedEthWrapper.Contract.IsDepositWhitelist(&_WhitelistedEthWrapper.CallOpts, vault, account)
}

// IsDepositWhitelist is a free data retrieval call binding the contract method 0xb09d93fc.
//
// Solidity: function isDepositWhitelist(address vault, address account) view returns(bool)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) IsDepositWhitelist(vault common.Address, account common.Address) (bool, error) {
	return _WhitelistedEthWrapper.Contract.IsDepositWhitelist(&_WhitelistedEthWrapper.CallOpts, vault, account)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) StETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.StETH(&_WhitelistedEthWrapper.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) StETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.StETH(&_WhitelistedEthWrapper.CallOpts)
}

// WstETH is a free data retrieval call binding the contract method 0x4aa07e64.
//
// Solidity: function wstETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCaller) WstETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistedEthWrapper.contract.Call(opts, &out, "wstETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WstETH is a free data retrieval call binding the contract method 0x4aa07e64.
//
// Solidity: function wstETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) WstETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.WstETH(&_WhitelistedEthWrapper.CallOpts)
}

// WstETH is a free data retrieval call binding the contract method 0x4aa07e64.
//
// Solidity: function wstETH() view returns(address)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperCallerSession) WstETH() (common.Address, error) {
	return _WhitelistedEthWrapper.Contract.WstETH(&_WhitelistedEthWrapper.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0bb9f5e1.
//
// Solidity: function deposit(address depositToken, uint256 amount, address vault, address receiver, address referral) payable returns(uint256 shares)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactor) Deposit(opts *bind.TransactOpts, depositToken common.Address, amount *big.Int, vault common.Address, receiver common.Address, referral common.Address) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.contract.Transact(opts, "deposit", depositToken, amount, vault, receiver, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0x0bb9f5e1.
//
// Solidity: function deposit(address depositToken, uint256 amount, address vault, address receiver, address referral) payable returns(uint256 shares)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) Deposit(depositToken common.Address, amount *big.Int, vault common.Address, receiver common.Address, referral common.Address) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.Deposit(&_WhitelistedEthWrapper.TransactOpts, depositToken, amount, vault, receiver, referral)
}

// Deposit is a paid mutator transaction binding the contract method 0x0bb9f5e1.
//
// Solidity: function deposit(address depositToken, uint256 amount, address vault, address receiver, address referral) payable returns(uint256 shares)
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorSession) Deposit(depositToken common.Address, amount *big.Int, vault common.Address, receiver common.Address, referral common.Address) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.Deposit(&_WhitelistedEthWrapper.TransactOpts, depositToken, amount, vault, receiver, referral)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x227a1d99.
//
// Solidity: function setDepositWhitelist(address vault, bool depositWhitelist_) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactor) SetDepositWhitelist(opts *bind.TransactOpts, vault common.Address, depositWhitelist_ bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.contract.Transact(opts, "setDepositWhitelist", vault, depositWhitelist_)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x227a1d99.
//
// Solidity: function setDepositWhitelist(address vault, bool depositWhitelist_) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) SetDepositWhitelist(vault common.Address, depositWhitelist_ bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.SetDepositWhitelist(&_WhitelistedEthWrapper.TransactOpts, vault, depositWhitelist_)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x227a1d99.
//
// Solidity: function setDepositWhitelist(address vault, bool depositWhitelist_) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorSession) SetDepositWhitelist(vault common.Address, depositWhitelist_ bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.SetDepositWhitelist(&_WhitelistedEthWrapper.TransactOpts, vault, depositWhitelist_)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xf58a0336.
//
// Solidity: function setDepositorWhitelistStatus(address vault, address account, bool status) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactor) SetDepositorWhitelistStatus(opts *bind.TransactOpts, vault common.Address, account common.Address, status bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.contract.Transact(opts, "setDepositorWhitelistStatus", vault, account, status)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xf58a0336.
//
// Solidity: function setDepositorWhitelistStatus(address vault, address account, bool status) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) SetDepositorWhitelistStatus(vault common.Address, account common.Address, status bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.SetDepositorWhitelistStatus(&_WhitelistedEthWrapper.TransactOpts, vault, account, status)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xf58a0336.
//
// Solidity: function setDepositorWhitelistStatus(address vault, address account, bool status) returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorSession) SetDepositorWhitelistStatus(vault common.Address, account common.Address, status bool) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.SetDepositorWhitelistStatus(&_WhitelistedEthWrapper.TransactOpts, vault, account, status)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistedEthWrapper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperSession) Receive() (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.Receive(&_WhitelistedEthWrapper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WhitelistedEthWrapper *WhitelistedEthWrapperTransactorSession) Receive() (*types.Transaction, error) {
	return _WhitelistedEthWrapper.Contract.Receive(&_WhitelistedEthWrapper.TransactOpts)
}
