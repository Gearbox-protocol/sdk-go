// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aaveFlTaker

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

// AaveFlTakerMetaData contains all meta data concerning the AaveFlTaker contract.
var AaveFlTakerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_aavePool\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aavePool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowedFLReceiver\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFLReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"takeFlashLoan\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAllowedFLReceiver\",\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotAllowedReceiverException\",\"inputs\":[]}]",
}

// AaveFlTakerABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveFlTakerMetaData.ABI instead.
var AaveFlTakerABI = AaveFlTakerMetaData.ABI

// AaveFlTaker is an auto generated Go binding around an Ethereum contract.
type AaveFlTaker struct {
	AaveFlTakerCaller     // Read-only binding to the contract
	AaveFlTakerTransactor // Write-only binding to the contract
	AaveFlTakerFilterer   // Log filterer for contract events
}

// AaveFlTakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveFlTakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFlTakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveFlTakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFlTakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFlTakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFlTakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveFlTakerSession struct {
	Contract     *AaveFlTaker      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveFlTakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveFlTakerCallerSession struct {
	Contract *AaveFlTakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AaveFlTakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveFlTakerTransactorSession struct {
	Contract     *AaveFlTakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AaveFlTakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveFlTakerRaw struct {
	Contract *AaveFlTaker // Generic contract binding to access the raw methods on
}

// AaveFlTakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveFlTakerCallerRaw struct {
	Contract *AaveFlTakerCaller // Generic read-only contract binding to access the raw methods on
}

// AaveFlTakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveFlTakerTransactorRaw struct {
	Contract *AaveFlTakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveFlTaker creates a new instance of AaveFlTaker, bound to a specific deployed contract.
func NewAaveFlTaker(address common.Address, backend bind.ContractBackend) (*AaveFlTaker, error) {
	contract, err := bindAaveFlTaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveFlTaker{AaveFlTakerCaller: AaveFlTakerCaller{contract: contract}, AaveFlTakerTransactor: AaveFlTakerTransactor{contract: contract}, AaveFlTakerFilterer: AaveFlTakerFilterer{contract: contract}}, nil
}

// NewAaveFlTakerCaller creates a new read-only instance of AaveFlTaker, bound to a specific deployed contract.
func NewAaveFlTakerCaller(address common.Address, caller bind.ContractCaller) (*AaveFlTakerCaller, error) {
	contract, err := bindAaveFlTaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveFlTakerCaller{contract: contract}, nil
}

// NewAaveFlTakerTransactor creates a new write-only instance of AaveFlTaker, bound to a specific deployed contract.
func NewAaveFlTakerTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveFlTakerTransactor, error) {
	contract, err := bindAaveFlTaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveFlTakerTransactor{contract: contract}, nil
}

// NewAaveFlTakerFilterer creates a new log filterer instance of AaveFlTaker, bound to a specific deployed contract.
func NewAaveFlTakerFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFlTakerFilterer, error) {
	contract, err := bindAaveFlTaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFlTakerFilterer{contract: contract}, nil
}

// bindAaveFlTaker binds a generic wrapper to an already deployed contract.
func bindAaveFlTaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveFlTakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveFlTaker *AaveFlTakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveFlTaker.Contract.AaveFlTakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveFlTaker *AaveFlTakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.AaveFlTakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveFlTaker *AaveFlTakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.AaveFlTakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveFlTaker *AaveFlTakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveFlTaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveFlTaker *AaveFlTakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveFlTaker *AaveFlTakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.contract.Transact(opts, method, params...)
}

// AavePool is a free data retrieval call binding the contract method 0xa03e4bc3.
//
// Solidity: function aavePool() view returns(address)
func (_AaveFlTaker *AaveFlTakerCaller) AavePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveFlTaker.contract.Call(opts, &out, "aavePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AavePool is a free data retrieval call binding the contract method 0xa03e4bc3.
//
// Solidity: function aavePool() view returns(address)
func (_AaveFlTaker *AaveFlTakerSession) AavePool() (common.Address, error) {
	return _AaveFlTaker.Contract.AavePool(&_AaveFlTaker.CallOpts)
}

// AavePool is a free data retrieval call binding the contract method 0xa03e4bc3.
//
// Solidity: function aavePool() view returns(address)
func (_AaveFlTaker *AaveFlTakerCallerSession) AavePool() (common.Address, error) {
	return _AaveFlTaker.Contract.AavePool(&_AaveFlTaker.CallOpts)
}

// AllowedFLReceiver is a free data retrieval call binding the contract method 0x119b90a2.
//
// Solidity: function allowedFLReceiver(address ) view returns(bool)
func (_AaveFlTaker *AaveFlTakerCaller) AllowedFLReceiver(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AaveFlTaker.contract.Call(opts, &out, "allowedFLReceiver", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedFLReceiver is a free data retrieval call binding the contract method 0x119b90a2.
//
// Solidity: function allowedFLReceiver(address ) view returns(bool)
func (_AaveFlTaker *AaveFlTakerSession) AllowedFLReceiver(arg0 common.Address) (bool, error) {
	return _AaveFlTaker.Contract.AllowedFLReceiver(&_AaveFlTaker.CallOpts, arg0)
}

// AllowedFLReceiver is a free data retrieval call binding the contract method 0x119b90a2.
//
// Solidity: function allowedFLReceiver(address ) view returns(bool)
func (_AaveFlTaker *AaveFlTakerCallerSession) AllowedFLReceiver(arg0 common.Address) (bool, error) {
	return _AaveFlTaker.Contract.AllowedFLReceiver(&_AaveFlTaker.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveFlTaker *AaveFlTakerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveFlTaker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveFlTaker *AaveFlTakerSession) Owner() (common.Address, error) {
	return _AaveFlTaker.Contract.Owner(&_AaveFlTaker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveFlTaker *AaveFlTakerCallerSession) Owner() (common.Address, error) {
	return _AaveFlTaker.Contract.Owner(&_AaveFlTaker.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveFlTaker *AaveFlTakerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveFlTaker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveFlTaker *AaveFlTakerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveFlTaker.Contract.RenounceOwnership(&_AaveFlTaker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveFlTaker *AaveFlTakerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveFlTaker.Contract.RenounceOwnership(&_AaveFlTaker.TransactOpts)
}

// SetAllowedFLReceiver is a paid mutator transaction binding the contract method 0xd0f6b875.
//
// Solidity: function setAllowedFLReceiver(address receiver, bool status) returns()
func (_AaveFlTaker *AaveFlTakerTransactor) SetAllowedFLReceiver(opts *bind.TransactOpts, receiver common.Address, status bool) (*types.Transaction, error) {
	return _AaveFlTaker.contract.Transact(opts, "setAllowedFLReceiver", receiver, status)
}

// SetAllowedFLReceiver is a paid mutator transaction binding the contract method 0xd0f6b875.
//
// Solidity: function setAllowedFLReceiver(address receiver, bool status) returns()
func (_AaveFlTaker *AaveFlTakerSession) SetAllowedFLReceiver(receiver common.Address, status bool) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.SetAllowedFLReceiver(&_AaveFlTaker.TransactOpts, receiver, status)
}

// SetAllowedFLReceiver is a paid mutator transaction binding the contract method 0xd0f6b875.
//
// Solidity: function setAllowedFLReceiver(address receiver, bool status) returns()
func (_AaveFlTaker *AaveFlTakerTransactorSession) SetAllowedFLReceiver(receiver common.Address, status bool) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.SetAllowedFLReceiver(&_AaveFlTaker.TransactOpts, receiver, status)
}

// TakeFlashLoan is a paid mutator transaction binding the contract method 0xb9c83453.
//
// Solidity: function takeFlashLoan(address asset, uint256 amount, bytes data) returns()
func (_AaveFlTaker *AaveFlTakerTransactor) TakeFlashLoan(opts *bind.TransactOpts, asset common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AaveFlTaker.contract.Transact(opts, "takeFlashLoan", asset, amount, data)
}

// TakeFlashLoan is a paid mutator transaction binding the contract method 0xb9c83453.
//
// Solidity: function takeFlashLoan(address asset, uint256 amount, bytes data) returns()
func (_AaveFlTaker *AaveFlTakerSession) TakeFlashLoan(asset common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.TakeFlashLoan(&_AaveFlTaker.TransactOpts, asset, amount, data)
}

// TakeFlashLoan is a paid mutator transaction binding the contract method 0xb9c83453.
//
// Solidity: function takeFlashLoan(address asset, uint256 amount, bytes data) returns()
func (_AaveFlTaker *AaveFlTakerTransactorSession) TakeFlashLoan(asset common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.TakeFlashLoan(&_AaveFlTaker.TransactOpts, asset, amount, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveFlTaker *AaveFlTakerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveFlTaker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveFlTaker *AaveFlTakerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.TransferOwnership(&_AaveFlTaker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveFlTaker *AaveFlTakerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveFlTaker.Contract.TransferOwnership(&_AaveFlTaker.TransactOpts, newOwner)
}

// AaveFlTakerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveFlTaker contract.
type AaveFlTakerOwnershipTransferredIterator struct {
	Event *AaveFlTakerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveFlTakerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveFlTakerOwnershipTransferred)
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
		it.Event = new(AaveFlTakerOwnershipTransferred)
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
func (it *AaveFlTakerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveFlTakerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveFlTakerOwnershipTransferred represents a OwnershipTransferred event raised by the AaveFlTaker contract.
type AaveFlTakerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveFlTaker *AaveFlTakerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveFlTakerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveFlTaker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveFlTakerOwnershipTransferredIterator{contract: _AaveFlTaker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveFlTaker *AaveFlTakerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveFlTakerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveFlTaker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveFlTakerOwnershipTransferred)
				if err := _AaveFlTaker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AaveFlTaker *AaveFlTakerFilterer) ParseOwnershipTransferred(log types.Log) (*AaveFlTakerOwnershipTransferred, error) {
	event := new(AaveFlTakerOwnershipTransferred)
	if err := _AaveFlTaker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveFlTakerSetAllowedFLReceiverIterator is returned from FilterSetAllowedFLReceiver and is used to iterate over the raw logs and unpacked data for SetAllowedFLReceiver events raised by the AaveFlTaker contract.
type AaveFlTakerSetAllowedFLReceiverIterator struct {
	Event *AaveFlTakerSetAllowedFLReceiver // Event containing the contract specifics and raw log

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
func (it *AaveFlTakerSetAllowedFLReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveFlTakerSetAllowedFLReceiver)
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
		it.Event = new(AaveFlTakerSetAllowedFLReceiver)
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
func (it *AaveFlTakerSetAllowedFLReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveFlTakerSetAllowedFLReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveFlTakerSetAllowedFLReceiver represents a SetAllowedFLReceiver event raised by the AaveFlTaker contract.
type AaveFlTakerSetAllowedFLReceiver struct {
	Consumer common.Address
	Status   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetAllowedFLReceiver is a free log retrieval operation binding the contract event 0x5c40d220dd4a83f4a7c95d6a06dd865715fc5a1dd4c7f85f69da051e00eeeb19.
//
// Solidity: event SetAllowedFLReceiver(address indexed consumer, bool status)
func (_AaveFlTaker *AaveFlTakerFilterer) FilterSetAllowedFLReceiver(opts *bind.FilterOpts, consumer []common.Address) (*AaveFlTakerSetAllowedFLReceiverIterator, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _AaveFlTaker.contract.FilterLogs(opts, "SetAllowedFLReceiver", consumerRule)
	if err != nil {
		return nil, err
	}
	return &AaveFlTakerSetAllowedFLReceiverIterator{contract: _AaveFlTaker.contract, event: "SetAllowedFLReceiver", logs: logs, sub: sub}, nil
}

// WatchSetAllowedFLReceiver is a free log subscription operation binding the contract event 0x5c40d220dd4a83f4a7c95d6a06dd865715fc5a1dd4c7f85f69da051e00eeeb19.
//
// Solidity: event SetAllowedFLReceiver(address indexed consumer, bool status)
func (_AaveFlTaker *AaveFlTakerFilterer) WatchSetAllowedFLReceiver(opts *bind.WatchOpts, sink chan<- *AaveFlTakerSetAllowedFLReceiver, consumer []common.Address) (event.Subscription, error) {

	var consumerRule []interface{}
	for _, consumerItem := range consumer {
		consumerRule = append(consumerRule, consumerItem)
	}

	logs, sub, err := _AaveFlTaker.contract.WatchLogs(opts, "SetAllowedFLReceiver", consumerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveFlTakerSetAllowedFLReceiver)
				if err := _AaveFlTaker.contract.UnpackLog(event, "SetAllowedFLReceiver", log); err != nil {
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

// ParseSetAllowedFLReceiver is a log parse operation binding the contract event 0x5c40d220dd4a83f4a7c95d6a06dd865715fc5a1dd4c7f85f69da051e00eeeb19.
//
// Solidity: event SetAllowedFLReceiver(address indexed consumer, bool status)
func (_AaveFlTaker *AaveFlTakerFilterer) ParseSetAllowedFLReceiver(log types.Log) (*AaveFlTakerSetAllowedFLReceiver, error) {
	event := new(AaveFlTakerSetAllowedFLReceiver)
	if err := _AaveFlTaker.contract.UnpackLog(event, "SetAllowedFLReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
