// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pauseMulticall

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PauseMulticallABI is the input ABI used to generate the binding from.
const PauseMulticallABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"contractACL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cr\",\"outputs\":[{\"internalType\":\"contractContractsRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAllContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAllCreditManagers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAllPools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PauseMulticall is an auto generated Go binding around an Ethereum contract.
type PauseMulticall struct {
	PauseMulticallCaller     // Read-only binding to the contract
	PauseMulticallTransactor // Write-only binding to the contract
	PauseMulticallFilterer   // Log filterer for contract events
}

// PauseMulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauseMulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseMulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauseMulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseMulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauseMulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseMulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauseMulticallSession struct {
	Contract     *PauseMulticall   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauseMulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauseMulticallCallerSession struct {
	Contract *PauseMulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PauseMulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauseMulticallTransactorSession struct {
	Contract     *PauseMulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PauseMulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauseMulticallRaw struct {
	Contract *PauseMulticall // Generic contract binding to access the raw methods on
}

// PauseMulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauseMulticallCallerRaw struct {
	Contract *PauseMulticallCaller // Generic read-only contract binding to access the raw methods on
}

// PauseMulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauseMulticallTransactorRaw struct {
	Contract *PauseMulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauseMulticall creates a new instance of PauseMulticall, bound to a specific deployed contract.
func NewPauseMulticall(address common.Address, backend bind.ContractBackend) (*PauseMulticall, error) {
	contract, err := bindPauseMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauseMulticall{PauseMulticallCaller: PauseMulticallCaller{contract: contract}, PauseMulticallTransactor: PauseMulticallTransactor{contract: contract}, PauseMulticallFilterer: PauseMulticallFilterer{contract: contract}}, nil
}

// NewPauseMulticallCaller creates a new read-only instance of PauseMulticall, bound to a specific deployed contract.
func NewPauseMulticallCaller(address common.Address, caller bind.ContractCaller) (*PauseMulticallCaller, error) {
	contract, err := bindPauseMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauseMulticallCaller{contract: contract}, nil
}

// NewPauseMulticallTransactor creates a new write-only instance of PauseMulticall, bound to a specific deployed contract.
func NewPauseMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*PauseMulticallTransactor, error) {
	contract, err := bindPauseMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauseMulticallTransactor{contract: contract}, nil
}

// NewPauseMulticallFilterer creates a new log filterer instance of PauseMulticall, bound to a specific deployed contract.
func NewPauseMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*PauseMulticallFilterer, error) {
	contract, err := bindPauseMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauseMulticallFilterer{contract: contract}, nil
}

// bindPauseMulticall binds a generic wrapper to an already deployed contract.
func bindPauseMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PauseMulticallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseMulticall *PauseMulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseMulticall.Contract.PauseMulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseMulticall *PauseMulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseMulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseMulticall *PauseMulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseMulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseMulticall *PauseMulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseMulticall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseMulticall *PauseMulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseMulticall *PauseMulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseMulticall.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PauseMulticall *PauseMulticallCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PauseMulticall.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PauseMulticall *PauseMulticallSession) Acl() (common.Address, error) {
	return _PauseMulticall.Contract.Acl(&_PauseMulticall.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_PauseMulticall *PauseMulticallCallerSession) Acl() (common.Address, error) {
	return _PauseMulticall.Contract.Acl(&_PauseMulticall.CallOpts)
}

// Cr is a free data retrieval call binding the contract method 0x62eb33e3.
//
// Solidity: function cr() view returns(address)
func (_PauseMulticall *PauseMulticallCaller) Cr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PauseMulticall.contract.Call(opts, &out, "cr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cr is a free data retrieval call binding the contract method 0x62eb33e3.
//
// Solidity: function cr() view returns(address)
func (_PauseMulticall *PauseMulticallSession) Cr() (common.Address, error) {
	return _PauseMulticall.Contract.Cr(&_PauseMulticall.CallOpts)
}

// Cr is a free data retrieval call binding the contract method 0x62eb33e3.
//
// Solidity: function cr() view returns(address)
func (_PauseMulticall *PauseMulticallCallerSession) Cr() (common.Address, error) {
	return _PauseMulticall.Contract.Cr(&_PauseMulticall.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PauseMulticall *PauseMulticallCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PauseMulticall.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PauseMulticall *PauseMulticallSession) Paused() (bool, error) {
	return _PauseMulticall.Contract.Paused(&_PauseMulticall.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PauseMulticall *PauseMulticallCallerSession) Paused() (bool, error) {
	return _PauseMulticall.Contract.Paused(&_PauseMulticall.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PauseMulticall *PauseMulticallTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PauseMulticall *PauseMulticallSession) Pause() (*types.Transaction, error) {
	return _PauseMulticall.Contract.Pause(&_PauseMulticall.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PauseMulticall *PauseMulticallTransactorSession) Pause() (*types.Transaction, error) {
	return _PauseMulticall.Contract.Pause(&_PauseMulticall.TransactOpts)
}

// PauseAllContracts is a paid mutator transaction binding the contract method 0x7c205cc4.
//
// Solidity: function pauseAllContracts() returns()
func (_PauseMulticall *PauseMulticallTransactor) PauseAllContracts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.contract.Transact(opts, "pauseAllContracts")
}

// PauseAllContracts is a paid mutator transaction binding the contract method 0x7c205cc4.
//
// Solidity: function pauseAllContracts() returns()
func (_PauseMulticall *PauseMulticallSession) PauseAllContracts() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllContracts(&_PauseMulticall.TransactOpts)
}

// PauseAllContracts is a paid mutator transaction binding the contract method 0x7c205cc4.
//
// Solidity: function pauseAllContracts() returns()
func (_PauseMulticall *PauseMulticallTransactorSession) PauseAllContracts() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllContracts(&_PauseMulticall.TransactOpts)
}

// PauseAllCreditManagers is a paid mutator transaction binding the contract method 0x6f074436.
//
// Solidity: function pauseAllCreditManagers() returns()
func (_PauseMulticall *PauseMulticallTransactor) PauseAllCreditManagers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.contract.Transact(opts, "pauseAllCreditManagers")
}

// PauseAllCreditManagers is a paid mutator transaction binding the contract method 0x6f074436.
//
// Solidity: function pauseAllCreditManagers() returns()
func (_PauseMulticall *PauseMulticallSession) PauseAllCreditManagers() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllCreditManagers(&_PauseMulticall.TransactOpts)
}

// PauseAllCreditManagers is a paid mutator transaction binding the contract method 0x6f074436.
//
// Solidity: function pauseAllCreditManagers() returns()
func (_PauseMulticall *PauseMulticallTransactorSession) PauseAllCreditManagers() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllCreditManagers(&_PauseMulticall.TransactOpts)
}

// PauseAllPools is a paid mutator transaction binding the contract method 0x488bf5b4.
//
// Solidity: function pauseAllPools() returns()
func (_PauseMulticall *PauseMulticallTransactor) PauseAllPools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.contract.Transact(opts, "pauseAllPools")
}

// PauseAllPools is a paid mutator transaction binding the contract method 0x488bf5b4.
//
// Solidity: function pauseAllPools() returns()
func (_PauseMulticall *PauseMulticallSession) PauseAllPools() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllPools(&_PauseMulticall.TransactOpts)
}

// PauseAllPools is a paid mutator transaction binding the contract method 0x488bf5b4.
//
// Solidity: function pauseAllPools() returns()
func (_PauseMulticall *PauseMulticallTransactorSession) PauseAllPools() (*types.Transaction, error) {
	return _PauseMulticall.Contract.PauseAllPools(&_PauseMulticall.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PauseMulticall *PauseMulticallTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseMulticall.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PauseMulticall *PauseMulticallSession) Unpause() (*types.Transaction, error) {
	return _PauseMulticall.Contract.Unpause(&_PauseMulticall.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PauseMulticall *PauseMulticallTransactorSession) Unpause() (*types.Transaction, error) {
	return _PauseMulticall.Contract.Unpause(&_PauseMulticall.TransactOpts)
}

// PauseMulticallPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PauseMulticall contract.
type PauseMulticallPausedIterator struct {
	Event *PauseMulticallPaused // Event containing the contract specifics and raw log

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
func (it *PauseMulticallPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseMulticallPaused)
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
		it.Event = new(PauseMulticallPaused)
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
func (it *PauseMulticallPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseMulticallPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseMulticallPaused represents a Paused event raised by the PauseMulticall contract.
type PauseMulticallPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PauseMulticall *PauseMulticallFilterer) FilterPaused(opts *bind.FilterOpts) (*PauseMulticallPausedIterator, error) {

	logs, sub, err := _PauseMulticall.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PauseMulticallPausedIterator{contract: _PauseMulticall.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PauseMulticall *PauseMulticallFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PauseMulticallPaused) (event.Subscription, error) {

	logs, sub, err := _PauseMulticall.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseMulticallPaused)
				if err := _PauseMulticall.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PauseMulticall *PauseMulticallFilterer) ParsePaused(log types.Log) (*PauseMulticallPaused, error) {
	event := new(PauseMulticallPaused)
	if err := _PauseMulticall.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PauseMulticallUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PauseMulticall contract.
type PauseMulticallUnpausedIterator struct {
	Event *PauseMulticallUnpaused // Event containing the contract specifics and raw log

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
func (it *PauseMulticallUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseMulticallUnpaused)
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
		it.Event = new(PauseMulticallUnpaused)
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
func (it *PauseMulticallUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseMulticallUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseMulticallUnpaused represents a Unpaused event raised by the PauseMulticall contract.
type PauseMulticallUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PauseMulticall *PauseMulticallFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PauseMulticallUnpausedIterator, error) {

	logs, sub, err := _PauseMulticall.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PauseMulticallUnpausedIterator{contract: _PauseMulticall.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PauseMulticall *PauseMulticallFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PauseMulticallUnpaused) (event.Subscription, error) {

	logs, sub, err := _PauseMulticall.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseMulticallUnpaused)
				if err := _PauseMulticall.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PauseMulticall *PauseMulticallFilterer) ParseUnpaused(log types.Log) (*PauseMulticallUnpaused, error) {
	event := new(PauseMulticallUnpaused)
	if err := _PauseMulticall.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
