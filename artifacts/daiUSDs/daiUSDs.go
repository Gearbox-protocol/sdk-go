// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package daiUSDs

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

// DaiUSDsMetaData contains all meta data concerning the DaiUSDs contract.
var DaiUSDsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdsJoin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"DaiToUsds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"UsdsToDai\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiToUsds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usds\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdsJoin\",\"outputs\":[{\"internalType\":\"contractUsdsJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"usdsToDai\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DaiUSDsABI is the input ABI used to generate the binding from.
// Deprecated: Use DaiUSDsMetaData.ABI instead.
var DaiUSDsABI = DaiUSDsMetaData.ABI

// DaiUSDs is an auto generated Go binding around an Ethereum contract.
type DaiUSDs struct {
	DaiUSDsCaller     // Read-only binding to the contract
	DaiUSDsTransactor // Write-only binding to the contract
	DaiUSDsFilterer   // Log filterer for contract events
}

// DaiUSDsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaiUSDsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiUSDsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaiUSDsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiUSDsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaiUSDsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiUSDsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaiUSDsSession struct {
	Contract     *DaiUSDs          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaiUSDsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaiUSDsCallerSession struct {
	Contract *DaiUSDsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DaiUSDsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaiUSDsTransactorSession struct {
	Contract     *DaiUSDsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DaiUSDsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaiUSDsRaw struct {
	Contract *DaiUSDs // Generic contract binding to access the raw methods on
}

// DaiUSDsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaiUSDsCallerRaw struct {
	Contract *DaiUSDsCaller // Generic read-only contract binding to access the raw methods on
}

// DaiUSDsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaiUSDsTransactorRaw struct {
	Contract *DaiUSDsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDaiUSDs creates a new instance of DaiUSDs, bound to a specific deployed contract.
func NewDaiUSDs(address common.Address, backend bind.ContractBackend) (*DaiUSDs, error) {
	contract, err := bindDaiUSDs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DaiUSDs{DaiUSDsCaller: DaiUSDsCaller{contract: contract}, DaiUSDsTransactor: DaiUSDsTransactor{contract: contract}, DaiUSDsFilterer: DaiUSDsFilterer{contract: contract}}, nil
}

// NewDaiUSDsCaller creates a new read-only instance of DaiUSDs, bound to a specific deployed contract.
func NewDaiUSDsCaller(address common.Address, caller bind.ContractCaller) (*DaiUSDsCaller, error) {
	contract, err := bindDaiUSDs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaiUSDsCaller{contract: contract}, nil
}

// NewDaiUSDsTransactor creates a new write-only instance of DaiUSDs, bound to a specific deployed contract.
func NewDaiUSDsTransactor(address common.Address, transactor bind.ContractTransactor) (*DaiUSDsTransactor, error) {
	contract, err := bindDaiUSDs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaiUSDsTransactor{contract: contract}, nil
}

// NewDaiUSDsFilterer creates a new log filterer instance of DaiUSDs, bound to a specific deployed contract.
func NewDaiUSDsFilterer(address common.Address, filterer bind.ContractFilterer) (*DaiUSDsFilterer, error) {
	contract, err := bindDaiUSDs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaiUSDsFilterer{contract: contract}, nil
}

// bindDaiUSDs binds a generic wrapper to an already deployed contract.
func bindDaiUSDs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DaiUSDsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaiUSDs *DaiUSDsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaiUSDs.Contract.DaiUSDsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaiUSDs *DaiUSDsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaiUSDs.Contract.DaiUSDsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaiUSDs *DaiUSDsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaiUSDs.Contract.DaiUSDsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DaiUSDs *DaiUSDsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DaiUSDs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DaiUSDs *DaiUSDsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DaiUSDs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DaiUSDs *DaiUSDsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DaiUSDs.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DaiUSDs *DaiUSDsCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaiUSDs.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DaiUSDs *DaiUSDsSession) Dai() (common.Address, error) {
	return _DaiUSDs.Contract.Dai(&_DaiUSDs.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_DaiUSDs *DaiUSDsCallerSession) Dai() (common.Address, error) {
	return _DaiUSDs.Contract.Dai(&_DaiUSDs.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DaiUSDs *DaiUSDsCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaiUSDs.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DaiUSDs *DaiUSDsSession) DaiJoin() (common.Address, error) {
	return _DaiUSDs.Contract.DaiJoin(&_DaiUSDs.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_DaiUSDs *DaiUSDsCallerSession) DaiJoin() (common.Address, error) {
	return _DaiUSDs.Contract.DaiJoin(&_DaiUSDs.CallOpts)
}

// Usds is a free data retrieval call binding the contract method 0x4cf282fb.
//
// Solidity: function usds() view returns(address)
func (_DaiUSDs *DaiUSDsCaller) Usds(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaiUSDs.contract.Call(opts, &out, "usds")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usds is a free data retrieval call binding the contract method 0x4cf282fb.
//
// Solidity: function usds() view returns(address)
func (_DaiUSDs *DaiUSDsSession) Usds() (common.Address, error) {
	return _DaiUSDs.Contract.Usds(&_DaiUSDs.CallOpts)
}

// Usds is a free data retrieval call binding the contract method 0x4cf282fb.
//
// Solidity: function usds() view returns(address)
func (_DaiUSDs *DaiUSDsCallerSession) Usds() (common.Address, error) {
	return _DaiUSDs.Contract.Usds(&_DaiUSDs.CallOpts)
}

// UsdsJoin is a free data retrieval call binding the contract method 0xfa1e2e86.
//
// Solidity: function usdsJoin() view returns(address)
func (_DaiUSDs *DaiUSDsCaller) UsdsJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DaiUSDs.contract.Call(opts, &out, "usdsJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdsJoin is a free data retrieval call binding the contract method 0xfa1e2e86.
//
// Solidity: function usdsJoin() view returns(address)
func (_DaiUSDs *DaiUSDsSession) UsdsJoin() (common.Address, error) {
	return _DaiUSDs.Contract.UsdsJoin(&_DaiUSDs.CallOpts)
}

// UsdsJoin is a free data retrieval call binding the contract method 0xfa1e2e86.
//
// Solidity: function usdsJoin() view returns(address)
func (_DaiUSDs *DaiUSDsCallerSession) UsdsJoin() (common.Address, error) {
	return _DaiUSDs.Contract.UsdsJoin(&_DaiUSDs.CallOpts)
}

// DaiToUsds is a paid mutator transaction binding the contract method 0xf2c07aae.
//
// Solidity: function daiToUsds(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsTransactor) DaiToUsds(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.contract.Transact(opts, "daiToUsds", usr, wad)
}

// DaiToUsds is a paid mutator transaction binding the contract method 0xf2c07aae.
//
// Solidity: function daiToUsds(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsSession) DaiToUsds(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.Contract.DaiToUsds(&_DaiUSDs.TransactOpts, usr, wad)
}

// DaiToUsds is a paid mutator transaction binding the contract method 0xf2c07aae.
//
// Solidity: function daiToUsds(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsTransactorSession) DaiToUsds(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.Contract.DaiToUsds(&_DaiUSDs.TransactOpts, usr, wad)
}

// UsdsToDai is a paid mutator transaction binding the contract method 0x68f30150.
//
// Solidity: function usdsToDai(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsTransactor) UsdsToDai(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.contract.Transact(opts, "usdsToDai", usr, wad)
}

// UsdsToDai is a paid mutator transaction binding the contract method 0x68f30150.
//
// Solidity: function usdsToDai(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsSession) UsdsToDai(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.Contract.UsdsToDai(&_DaiUSDs.TransactOpts, usr, wad)
}

// UsdsToDai is a paid mutator transaction binding the contract method 0x68f30150.
//
// Solidity: function usdsToDai(address usr, uint256 wad) returns()
func (_DaiUSDs *DaiUSDsTransactorSession) UsdsToDai(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DaiUSDs.Contract.UsdsToDai(&_DaiUSDs.TransactOpts, usr, wad)
}

// DaiUSDsDaiToUsdsIterator is returned from FilterDaiToUsds and is used to iterate over the raw logs and unpacked data for DaiToUsds events raised by the DaiUSDs contract.
type DaiUSDsDaiToUsdsIterator struct {
	Event *DaiUSDsDaiToUsds // Event containing the contract specifics and raw log

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
func (it *DaiUSDsDaiToUsdsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiUSDsDaiToUsds)
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
		it.Event = new(DaiUSDsDaiToUsds)
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
func (it *DaiUSDsDaiToUsdsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiUSDsDaiToUsdsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiUSDsDaiToUsds represents a DaiToUsds event raised by the DaiUSDs contract.
type DaiUSDsDaiToUsds struct {
	Caller common.Address
	Usr    common.Address
	Wad    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaiToUsds is a free log retrieval operation binding the contract event 0x23889db6b244d239344deac88ec7781d810b3873fe323939d1ce0e8ac9564235.
//
// Solidity: event DaiToUsds(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) FilterDaiToUsds(opts *bind.FilterOpts, caller []common.Address, usr []common.Address) (*DaiUSDsDaiToUsdsIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _DaiUSDs.contract.FilterLogs(opts, "DaiToUsds", callerRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &DaiUSDsDaiToUsdsIterator{contract: _DaiUSDs.contract, event: "DaiToUsds", logs: logs, sub: sub}, nil
}

// WatchDaiToUsds is a free log subscription operation binding the contract event 0x23889db6b244d239344deac88ec7781d810b3873fe323939d1ce0e8ac9564235.
//
// Solidity: event DaiToUsds(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) WatchDaiToUsds(opts *bind.WatchOpts, sink chan<- *DaiUSDsDaiToUsds, caller []common.Address, usr []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _DaiUSDs.contract.WatchLogs(opts, "DaiToUsds", callerRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiUSDsDaiToUsds)
				if err := _DaiUSDs.contract.UnpackLog(event, "DaiToUsds", log); err != nil {
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

// ParseDaiToUsds is a log parse operation binding the contract event 0x23889db6b244d239344deac88ec7781d810b3873fe323939d1ce0e8ac9564235.
//
// Solidity: event DaiToUsds(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) ParseDaiToUsds(log types.Log) (*DaiUSDsDaiToUsds, error) {
	event := new(DaiUSDsDaiToUsds)
	if err := _DaiUSDs.contract.UnpackLog(event, "DaiToUsds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DaiUSDsUsdsToDaiIterator is returned from FilterUsdsToDai and is used to iterate over the raw logs and unpacked data for UsdsToDai events raised by the DaiUSDs contract.
type DaiUSDsUsdsToDaiIterator struct {
	Event *DaiUSDsUsdsToDai // Event containing the contract specifics and raw log

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
func (it *DaiUSDsUsdsToDaiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiUSDsUsdsToDai)
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
		it.Event = new(DaiUSDsUsdsToDai)
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
func (it *DaiUSDsUsdsToDaiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiUSDsUsdsToDaiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiUSDsUsdsToDai represents a UsdsToDai event raised by the DaiUSDs contract.
type DaiUSDsUsdsToDai struct {
	Caller common.Address
	Usr    common.Address
	Wad    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUsdsToDai is a free log retrieval operation binding the contract event 0x9c7bd3b60b7e46fc1d7a38d88b7dd4be0e028b80b706848fc3c065eabc8e6a9b.
//
// Solidity: event UsdsToDai(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) FilterUsdsToDai(opts *bind.FilterOpts, caller []common.Address, usr []common.Address) (*DaiUSDsUsdsToDaiIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _DaiUSDs.contract.FilterLogs(opts, "UsdsToDai", callerRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &DaiUSDsUsdsToDaiIterator{contract: _DaiUSDs.contract, event: "UsdsToDai", logs: logs, sub: sub}, nil
}

// WatchUsdsToDai is a free log subscription operation binding the contract event 0x9c7bd3b60b7e46fc1d7a38d88b7dd4be0e028b80b706848fc3c065eabc8e6a9b.
//
// Solidity: event UsdsToDai(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) WatchUsdsToDai(opts *bind.WatchOpts, sink chan<- *DaiUSDsUsdsToDai, caller []common.Address, usr []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _DaiUSDs.contract.WatchLogs(opts, "UsdsToDai", callerRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiUSDsUsdsToDai)
				if err := _DaiUSDs.contract.UnpackLog(event, "UsdsToDai", log); err != nil {
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

// ParseUsdsToDai is a log parse operation binding the contract event 0x9c7bd3b60b7e46fc1d7a38d88b7dd4be0e028b80b706848fc3c065eabc8e6a9b.
//
// Solidity: event UsdsToDai(address indexed caller, address indexed usr, uint256 wad)
func (_DaiUSDs *DaiUSDsFilterer) ParseUsdsToDai(log types.Log) (*DaiUSDsUsdsToDai, error) {
	event := new(DaiUSDsUsdsToDai)
	if err := _DaiUSDs.contract.UnpackLog(event, "UsdsToDai", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
