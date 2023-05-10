// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditConfigurator210

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

// CreditConfigurator210MetaData contains all meta data concerning the CreditConfigurator210 contract.
var CreditConfigurator210MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\",\"indexed\":false}],\"type\":\"event\",\"name\":\"NewEmergencyLiquidationDiscount\",\"anonymous\":false},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"NewMaxCumulativeLoss\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CumulativeLossReset\",\"type\":\"event\"}]",
}

// CreditConfigurator210ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditConfigurator210MetaData.ABI instead.
var CreditConfigurator210ABI = CreditConfigurator210MetaData.ABI

// CreditConfigurator210 is an auto generated Go binding around an Ethereum contract.
type CreditConfigurator210 struct {
	CreditConfigurator210Caller     // Read-only binding to the contract
	CreditConfigurator210Transactor // Write-only binding to the contract
	CreditConfigurator210Filterer   // Log filterer for contract events
}

// CreditConfigurator210Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditConfigurator210Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfigurator210Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditConfigurator210Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfigurator210Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditConfigurator210Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditConfigurator210Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditConfigurator210Session struct {
	Contract     *CreditConfigurator210 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CreditConfigurator210CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditConfigurator210CallerSession struct {
	Contract *CreditConfigurator210Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CreditConfigurator210TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditConfigurator210TransactorSession struct {
	Contract     *CreditConfigurator210Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CreditConfigurator210Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditConfigurator210Raw struct {
	Contract *CreditConfigurator210 // Generic contract binding to access the raw methods on
}

// CreditConfigurator210CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditConfigurator210CallerRaw struct {
	Contract *CreditConfigurator210Caller // Generic read-only contract binding to access the raw methods on
}

// CreditConfigurator210TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditConfigurator210TransactorRaw struct {
	Contract *CreditConfigurator210Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditConfigurator210 creates a new instance of CreditConfigurator210, bound to a specific deployed contract.
func NewCreditConfigurator210(address common.Address, backend bind.ContractBackend) (*CreditConfigurator210, error) {
	contract, err := bindCreditConfigurator210(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210{CreditConfigurator210Caller: CreditConfigurator210Caller{contract: contract}, CreditConfigurator210Transactor: CreditConfigurator210Transactor{contract: contract}, CreditConfigurator210Filterer: CreditConfigurator210Filterer{contract: contract}}, nil
}

// NewCreditConfigurator210Caller creates a new read-only instance of CreditConfigurator210, bound to a specific deployed contract.
func NewCreditConfigurator210Caller(address common.Address, caller bind.ContractCaller) (*CreditConfigurator210Caller, error) {
	contract, err := bindCreditConfigurator210(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210Caller{contract: contract}, nil
}

// NewCreditConfigurator210Transactor creates a new write-only instance of CreditConfigurator210, bound to a specific deployed contract.
func NewCreditConfigurator210Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditConfigurator210Transactor, error) {
	contract, err := bindCreditConfigurator210(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210Transactor{contract: contract}, nil
}

// NewCreditConfigurator210Filterer creates a new log filterer instance of CreditConfigurator210, bound to a specific deployed contract.
func NewCreditConfigurator210Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditConfigurator210Filterer, error) {
	contract, err := bindCreditConfigurator210(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210Filterer{contract: contract}, nil
}

// bindCreditConfigurator210 binds a generic wrapper to an already deployed contract.
func bindCreditConfigurator210(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditConfigurator210MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator210 *CreditConfigurator210Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator210.Contract.CreditConfigurator210Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator210 *CreditConfigurator210Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator210.Contract.CreditConfigurator210Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator210 *CreditConfigurator210Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator210.Contract.CreditConfigurator210Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditConfigurator210 *CreditConfigurator210CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditConfigurator210.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditConfigurator210 *CreditConfigurator210TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditConfigurator210.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditConfigurator210 *CreditConfigurator210TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditConfigurator210.Contract.contract.Transact(opts, method, params...)
}

// CreditConfigurator210CumulativeLossResetIterator is returned from FilterCumulativeLossReset and is used to iterate over the raw logs and unpacked data for CumulativeLossReset events raised by the CreditConfigurator210 contract.
type CreditConfigurator210CumulativeLossResetIterator struct {
	Event *CreditConfigurator210CumulativeLossReset // Event containing the contract specifics and raw log

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
func (it *CreditConfigurator210CumulativeLossResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfigurator210CumulativeLossReset)
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
		it.Event = new(CreditConfigurator210CumulativeLossReset)
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
func (it *CreditConfigurator210CumulativeLossResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfigurator210CumulativeLossResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfigurator210CumulativeLossReset represents a CumulativeLossReset event raised by the CreditConfigurator210 contract.
type CreditConfigurator210CumulativeLossReset struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCumulativeLossReset is a free log retrieval operation binding the contract event 0x76867738b1ea031dba288b212a346e09748d59c0dd4a24a79d73b2ff4825445c.
//
// Solidity: event CumulativeLossReset()
func (_CreditConfigurator210 *CreditConfigurator210Filterer) FilterCumulativeLossReset(opts *bind.FilterOpts) (*CreditConfigurator210CumulativeLossResetIterator, error) {

	logs, sub, err := _CreditConfigurator210.contract.FilterLogs(opts, "CumulativeLossReset")
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210CumulativeLossResetIterator{contract: _CreditConfigurator210.contract, event: "CumulativeLossReset", logs: logs, sub: sub}, nil
}

// WatchCumulativeLossReset is a free log subscription operation binding the contract event 0x76867738b1ea031dba288b212a346e09748d59c0dd4a24a79d73b2ff4825445c.
//
// Solidity: event CumulativeLossReset()
func (_CreditConfigurator210 *CreditConfigurator210Filterer) WatchCumulativeLossReset(opts *bind.WatchOpts, sink chan<- *CreditConfigurator210CumulativeLossReset) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator210.contract.WatchLogs(opts, "CumulativeLossReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfigurator210CumulativeLossReset)
				if err := _CreditConfigurator210.contract.UnpackLog(event, "CumulativeLossReset", log); err != nil {
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

// ParseCumulativeLossReset is a log parse operation binding the contract event 0x76867738b1ea031dba288b212a346e09748d59c0dd4a24a79d73b2ff4825445c.
//
// Solidity: event CumulativeLossReset()
func (_CreditConfigurator210 *CreditConfigurator210Filterer) ParseCumulativeLossReset(log types.Log) (*CreditConfigurator210CumulativeLossReset, error) {
	event := new(CreditConfigurator210CumulativeLossReset)
	if err := _CreditConfigurator210.contract.UnpackLog(event, "CumulativeLossReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfigurator210NewEmergencyLiquidationDiscountIterator is returned from FilterNewEmergencyLiquidationDiscount and is used to iterate over the raw logs and unpacked data for NewEmergencyLiquidationDiscount events raised by the CreditConfigurator210 contract.
type CreditConfigurator210NewEmergencyLiquidationDiscountIterator struct {
	Event *CreditConfigurator210NewEmergencyLiquidationDiscount // Event containing the contract specifics and raw log

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
func (it *CreditConfigurator210NewEmergencyLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfigurator210NewEmergencyLiquidationDiscount)
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
		it.Event = new(CreditConfigurator210NewEmergencyLiquidationDiscount)
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
func (it *CreditConfigurator210NewEmergencyLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfigurator210NewEmergencyLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfigurator210NewEmergencyLiquidationDiscount represents a NewEmergencyLiquidationDiscount event raised by the CreditConfigurator210 contract.
type CreditConfigurator210NewEmergencyLiquidationDiscount struct {
	Arg0 uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewEmergencyLiquidationDiscount is a free log retrieval operation binding the contract event 0xf6752a535d833e0a40003935cfd7e73e6865d523e7bbcf4e5526a2afc5e18e5c.
//
// Solidity: event NewEmergencyLiquidationDiscount(uint16 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) FilterNewEmergencyLiquidationDiscount(opts *bind.FilterOpts) (*CreditConfigurator210NewEmergencyLiquidationDiscountIterator, error) {

	logs, sub, err := _CreditConfigurator210.contract.FilterLogs(opts, "NewEmergencyLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210NewEmergencyLiquidationDiscountIterator{contract: _CreditConfigurator210.contract, event: "NewEmergencyLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchNewEmergencyLiquidationDiscount is a free log subscription operation binding the contract event 0xf6752a535d833e0a40003935cfd7e73e6865d523e7bbcf4e5526a2afc5e18e5c.
//
// Solidity: event NewEmergencyLiquidationDiscount(uint16 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) WatchNewEmergencyLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *CreditConfigurator210NewEmergencyLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator210.contract.WatchLogs(opts, "NewEmergencyLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfigurator210NewEmergencyLiquidationDiscount)
				if err := _CreditConfigurator210.contract.UnpackLog(event, "NewEmergencyLiquidationDiscount", log); err != nil {
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

// ParseNewEmergencyLiquidationDiscount is a log parse operation binding the contract event 0xf6752a535d833e0a40003935cfd7e73e6865d523e7bbcf4e5526a2afc5e18e5c.
//
// Solidity: event NewEmergencyLiquidationDiscount(uint16 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) ParseNewEmergencyLiquidationDiscount(log types.Log) (*CreditConfigurator210NewEmergencyLiquidationDiscount, error) {
	event := new(CreditConfigurator210NewEmergencyLiquidationDiscount)
	if err := _CreditConfigurator210.contract.UnpackLog(event, "NewEmergencyLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditConfigurator210NewMaxCumulativeLossIterator is returned from FilterNewMaxCumulativeLoss and is used to iterate over the raw logs and unpacked data for NewMaxCumulativeLoss events raised by the CreditConfigurator210 contract.
type CreditConfigurator210NewMaxCumulativeLossIterator struct {
	Event *CreditConfigurator210NewMaxCumulativeLoss // Event containing the contract specifics and raw log

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
func (it *CreditConfigurator210NewMaxCumulativeLossIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditConfigurator210NewMaxCumulativeLoss)
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
		it.Event = new(CreditConfigurator210NewMaxCumulativeLoss)
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
func (it *CreditConfigurator210NewMaxCumulativeLossIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditConfigurator210NewMaxCumulativeLossIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditConfigurator210NewMaxCumulativeLoss represents a NewMaxCumulativeLoss event raised by the CreditConfigurator210 contract.
type CreditConfigurator210NewMaxCumulativeLoss struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewMaxCumulativeLoss is a free log retrieval operation binding the contract event 0x1cb1645cf44744a79c6a2c5b59b369aeb21904c0348c3b5952fd588fb861cc46.
//
// Solidity: event NewMaxCumulativeLoss(uint128 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) FilterNewMaxCumulativeLoss(opts *bind.FilterOpts) (*CreditConfigurator210NewMaxCumulativeLossIterator, error) {

	logs, sub, err := _CreditConfigurator210.contract.FilterLogs(opts, "NewMaxCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return &CreditConfigurator210NewMaxCumulativeLossIterator{contract: _CreditConfigurator210.contract, event: "NewMaxCumulativeLoss", logs: logs, sub: sub}, nil
}

// WatchNewMaxCumulativeLoss is a free log subscription operation binding the contract event 0x1cb1645cf44744a79c6a2c5b59b369aeb21904c0348c3b5952fd588fb861cc46.
//
// Solidity: event NewMaxCumulativeLoss(uint128 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) WatchNewMaxCumulativeLoss(opts *bind.WatchOpts, sink chan<- *CreditConfigurator210NewMaxCumulativeLoss) (event.Subscription, error) {

	logs, sub, err := _CreditConfigurator210.contract.WatchLogs(opts, "NewMaxCumulativeLoss")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditConfigurator210NewMaxCumulativeLoss)
				if err := _CreditConfigurator210.contract.UnpackLog(event, "NewMaxCumulativeLoss", log); err != nil {
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

// ParseNewMaxCumulativeLoss is a log parse operation binding the contract event 0x1cb1645cf44744a79c6a2c5b59b369aeb21904c0348c3b5952fd588fb861cc46.
//
// Solidity: event NewMaxCumulativeLoss(uint128 arg0)
func (_CreditConfigurator210 *CreditConfigurator210Filterer) ParseNewMaxCumulativeLoss(log types.Log) (*CreditConfigurator210NewMaxCumulativeLoss, error) {
	event := new(CreditConfigurator210NewMaxCumulativeLoss)
	if err := _CreditConfigurator210.contract.UnpackLog(event, "NewMaxCumulativeLoss", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
