// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolQuotaKeeperv3

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

// PoolQuotaKeeperv3MetaData contains all meta data concerning the PoolQuotaKeeperv3 contract.
var PoolQuotaKeeperv3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int96\",\"name\":\"quotaChange\",\"type\":\"int96\"}],\"name\":\"UpdateQuota\",\"type\":\"event\"}]",
}

// PoolQuotaKeeperv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolQuotaKeeperv3MetaData.ABI instead.
var PoolQuotaKeeperv3ABI = PoolQuotaKeeperv3MetaData.ABI

// PoolQuotaKeeperv3 is an auto generated Go binding around an Ethereum contract.
type PoolQuotaKeeperv3 struct {
	PoolQuotaKeeperv3Caller     // Read-only binding to the contract
	PoolQuotaKeeperv3Transactor // Write-only binding to the contract
	PoolQuotaKeeperv3Filterer   // Log filterer for contract events
}

// PoolQuotaKeeperv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolQuotaKeeperv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolQuotaKeeperv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolQuotaKeeperv3Session struct {
	Contract     *PoolQuotaKeeperv3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PoolQuotaKeeperv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolQuotaKeeperv3CallerSession struct {
	Contract *PoolQuotaKeeperv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PoolQuotaKeeperv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolQuotaKeeperv3TransactorSession struct {
	Contract     *PoolQuotaKeeperv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PoolQuotaKeeperv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type PoolQuotaKeeperv3Raw struct {
	Contract *PoolQuotaKeeperv3 // Generic contract binding to access the raw methods on
}

// PoolQuotaKeeperv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3CallerRaw struct {
	Contract *PoolQuotaKeeperv3Caller // Generic read-only contract binding to access the raw methods on
}

// PoolQuotaKeeperv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolQuotaKeeperv3TransactorRaw struct {
	Contract *PoolQuotaKeeperv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolQuotaKeeperv3 creates a new instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3(address common.Address, backend bind.ContractBackend) (*PoolQuotaKeeperv3, error) {
	contract, err := bindPoolQuotaKeeperv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3{PoolQuotaKeeperv3Caller: PoolQuotaKeeperv3Caller{contract: contract}, PoolQuotaKeeperv3Transactor: PoolQuotaKeeperv3Transactor{contract: contract}, PoolQuotaKeeperv3Filterer: PoolQuotaKeeperv3Filterer{contract: contract}}, nil
}

// NewPoolQuotaKeeperv3Caller creates a new read-only instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Caller(address common.Address, caller bind.ContractCaller) (*PoolQuotaKeeperv3Caller, error) {
	contract, err := bindPoolQuotaKeeperv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Caller{contract: contract}, nil
}

// NewPoolQuotaKeeperv3Transactor creates a new write-only instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Transactor(address common.Address, transactor bind.ContractTransactor) (*PoolQuotaKeeperv3Transactor, error) {
	contract, err := bindPoolQuotaKeeperv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Transactor{contract: contract}, nil
}

// NewPoolQuotaKeeperv3Filterer creates a new log filterer instance of PoolQuotaKeeperv3, bound to a specific deployed contract.
func NewPoolQuotaKeeperv3Filterer(address common.Address, filterer bind.ContractFilterer) (*PoolQuotaKeeperv3Filterer, error) {
	contract, err := bindPoolQuotaKeeperv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3Filterer{contract: contract}, nil
}

// bindPoolQuotaKeeperv3 binds a generic wrapper to an already deployed contract.
func bindPoolQuotaKeeperv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolQuotaKeeperv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.PoolQuotaKeeperv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolQuotaKeeperv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolQuotaKeeperv3.Contract.contract.Transact(opts, method, params...)
}

// PoolQuotaKeeperv3UpdateQuotaIterator is returned from FilterUpdateQuota and is used to iterate over the raw logs and unpacked data for UpdateQuota events raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateQuotaIterator struct {
	Event *PoolQuotaKeeperv3UpdateQuota // Event containing the contract specifics and raw log

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
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolQuotaKeeperv3UpdateQuota)
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
		it.Event = new(PoolQuotaKeeperv3UpdateQuota)
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
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolQuotaKeeperv3UpdateQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolQuotaKeeperv3UpdateQuota represents a UpdateQuota event raised by the PoolQuotaKeeperv3 contract.
type PoolQuotaKeeperv3UpdateQuota struct {
	CreditAccount common.Address
	Token         common.Address
	QuotaChange   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuota is a free log retrieval operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) FilterUpdateQuota(opts *bind.FilterOpts, creditAccount []common.Address, token []common.Address) (*PoolQuotaKeeperv3UpdateQuotaIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.FilterLogs(opts, "UpdateQuota", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PoolQuotaKeeperv3UpdateQuotaIterator{contract: _PoolQuotaKeeperv3.contract, event: "UpdateQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateQuota is a free log subscription operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) WatchUpdateQuota(opts *bind.WatchOpts, sink chan<- *PoolQuotaKeeperv3UpdateQuota, creditAccount []common.Address, token []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PoolQuotaKeeperv3.contract.WatchLogs(opts, "UpdateQuota", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolQuotaKeeperv3UpdateQuota)
				if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateQuota", log); err != nil {
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

// ParseUpdateQuota is a log parse operation binding the contract event 0x22cce666192befd41ad1b89f8592d35a7ce7c6960853f89ada56db03bb61b096.
//
// Solidity: event UpdateQuota(address indexed creditAccount, address indexed token, int96 quotaChange)
func (_PoolQuotaKeeperv3 *PoolQuotaKeeperv3Filterer) ParseUpdateQuota(log types.Log) (*PoolQuotaKeeperv3UpdateQuota, error) {
	event := new(PoolQuotaKeeperv3UpdateQuota)
	if err := _PoolQuotaKeeperv3.contract.UnpackLog(event, "UpdateQuota", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
