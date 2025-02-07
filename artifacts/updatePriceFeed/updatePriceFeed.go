// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package updatePriceFeed

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

// UpdatePriceFeedMetaData contains all meta data concerning the UpdatePriceFeed contract.
var UpdatePriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"skipPriceCheck\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatable\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updatePrice\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"UpdatePrice\",\"inputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// UpdatePriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use UpdatePriceFeedMetaData.ABI instead.
var UpdatePriceFeedABI = UpdatePriceFeedMetaData.ABI

// UpdatePriceFeed is an auto generated Go binding around an Ethereum contract.
type UpdatePriceFeed struct {
	UpdatePriceFeedCaller     // Read-only binding to the contract
	UpdatePriceFeedTransactor // Write-only binding to the contract
	UpdatePriceFeedFilterer   // Log filterer for contract events
}

// UpdatePriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpdatePriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdatePriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpdatePriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdatePriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpdatePriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpdatePriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpdatePriceFeedSession struct {
	Contract     *UpdatePriceFeed  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpdatePriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpdatePriceFeedCallerSession struct {
	Contract *UpdatePriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UpdatePriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpdatePriceFeedTransactorSession struct {
	Contract     *UpdatePriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UpdatePriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpdatePriceFeedRaw struct {
	Contract *UpdatePriceFeed // Generic contract binding to access the raw methods on
}

// UpdatePriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpdatePriceFeedCallerRaw struct {
	Contract *UpdatePriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// UpdatePriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpdatePriceFeedTransactorRaw struct {
	Contract *UpdatePriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpdatePriceFeed creates a new instance of UpdatePriceFeed, bound to a specific deployed contract.
func NewUpdatePriceFeed(address common.Address, backend bind.ContractBackend) (*UpdatePriceFeed, error) {
	contract, err := bindUpdatePriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpdatePriceFeed{UpdatePriceFeedCaller: UpdatePriceFeedCaller{contract: contract}, UpdatePriceFeedTransactor: UpdatePriceFeedTransactor{contract: contract}, UpdatePriceFeedFilterer: UpdatePriceFeedFilterer{contract: contract}}, nil
}

// NewUpdatePriceFeedCaller creates a new read-only instance of UpdatePriceFeed, bound to a specific deployed contract.
func NewUpdatePriceFeedCaller(address common.Address, caller bind.ContractCaller) (*UpdatePriceFeedCaller, error) {
	contract, err := bindUpdatePriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpdatePriceFeedCaller{contract: contract}, nil
}

// NewUpdatePriceFeedTransactor creates a new write-only instance of UpdatePriceFeed, bound to a specific deployed contract.
func NewUpdatePriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*UpdatePriceFeedTransactor, error) {
	contract, err := bindUpdatePriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpdatePriceFeedTransactor{contract: contract}, nil
}

// NewUpdatePriceFeedFilterer creates a new log filterer instance of UpdatePriceFeed, bound to a specific deployed contract.
func NewUpdatePriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*UpdatePriceFeedFilterer, error) {
	contract, err := bindUpdatePriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpdatePriceFeedFilterer{contract: contract}, nil
}

// bindUpdatePriceFeed binds a generic wrapper to an already deployed contract.
func bindUpdatePriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpdatePriceFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdatePriceFeed *UpdatePriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdatePriceFeed.Contract.UpdatePriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdatePriceFeed *UpdatePriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.UpdatePriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdatePriceFeed *UpdatePriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.UpdatePriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpdatePriceFeed *UpdatePriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpdatePriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpdatePriceFeed *UpdatePriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpdatePriceFeed *UpdatePriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.contract.Transact(opts, method, params...)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_UpdatePriceFeed *UpdatePriceFeedSession) ContractType() ([32]byte, error) {
	return _UpdatePriceFeed.Contract.ContractType(&_UpdatePriceFeed.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) ContractType() ([32]byte, error) {
	return _UpdatePriceFeed.Contract.ContractType(&_UpdatePriceFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpdatePriceFeed *UpdatePriceFeedSession) Decimals() (uint8, error) {
	return _UpdatePriceFeed.Contract.Decimals(&_UpdatePriceFeed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) Decimals() (uint8, error) {
	return _UpdatePriceFeed.Contract.Decimals(&_UpdatePriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_UpdatePriceFeed *UpdatePriceFeedSession) Description() (string, error) {
	return _UpdatePriceFeed.Contract.Description(&_UpdatePriceFeed.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) Description() (string, error) {
	return _UpdatePriceFeed.Contract.Description(&_UpdatePriceFeed.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256 updatedAt, uint80)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) LatestRoundData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "latestRoundData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256 updatedAt, uint80)
func (_UpdatePriceFeed *UpdatePriceFeedSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _UpdatePriceFeed.Contract.LatestRoundData(&_UpdatePriceFeed.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80, int256 answer, uint256, uint256 updatedAt, uint80)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) LatestRoundData() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _UpdatePriceFeed.Contract.LatestRoundData(&_UpdatePriceFeed.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) SkipPriceCheck(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "skipPriceCheck")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedSession) SkipPriceCheck() (bool, error) {
	return _UpdatePriceFeed.Contract.SkipPriceCheck(&_UpdatePriceFeed.CallOpts)
}

// SkipPriceCheck is a free data retrieval call binding the contract method 0xd62ada11.
//
// Solidity: function skipPriceCheck() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) SkipPriceCheck() (bool, error) {
	return _UpdatePriceFeed.Contract.SkipPriceCheck(&_UpdatePriceFeed.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) Updatable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "updatable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedSession) Updatable() (bool, error) {
	return _UpdatePriceFeed.Contract.Updatable(&_UpdatePriceFeed.CallOpts)
}

// Updatable is a free data retrieval call binding the contract method 0xe75aeec8.
//
// Solidity: function updatable() view returns(bool)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) Updatable() (bool, error) {
	return _UpdatePriceFeed.Contract.Updatable(&_UpdatePriceFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_UpdatePriceFeed *UpdatePriceFeedCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpdatePriceFeed.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_UpdatePriceFeed *UpdatePriceFeedSession) Version() (*big.Int, error) {
	return _UpdatePriceFeed.Contract.Version(&_UpdatePriceFeed.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_UpdatePriceFeed *UpdatePriceFeedCallerSession) Version() (*big.Int, error) {
	return _UpdatePriceFeed.Contract.Version(&_UpdatePriceFeed.CallOpts)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_UpdatePriceFeed *UpdatePriceFeedTransactor) UpdatePrice(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _UpdatePriceFeed.contract.Transact(opts, "updatePrice", data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_UpdatePriceFeed *UpdatePriceFeedSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.UpdatePrice(&_UpdatePriceFeed.TransactOpts, data)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8736ec47.
//
// Solidity: function updatePrice(bytes data) returns()
func (_UpdatePriceFeed *UpdatePriceFeedTransactorSession) UpdatePrice(data []byte) (*types.Transaction, error) {
	return _UpdatePriceFeed.Contract.UpdatePrice(&_UpdatePriceFeed.TransactOpts, data)
}

// UpdatePriceFeedUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the UpdatePriceFeed contract.
type UpdatePriceFeedUpdatePriceIterator struct {
	Event *UpdatePriceFeedUpdatePrice // Event containing the contract specifics and raw log

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
func (it *UpdatePriceFeedUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpdatePriceFeedUpdatePrice)
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
		it.Event = new(UpdatePriceFeedUpdatePrice)
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
func (it *UpdatePriceFeedUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpdatePriceFeedUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpdatePriceFeedUpdatePrice represents a UpdatePrice event raised by the UpdatePriceFeed contract.
type UpdatePriceFeedUpdatePrice struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_UpdatePriceFeed *UpdatePriceFeedFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*UpdatePriceFeedUpdatePriceIterator, error) {

	logs, sub, err := _UpdatePriceFeed.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &UpdatePriceFeedUpdatePriceIterator{contract: _UpdatePriceFeed.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_UpdatePriceFeed *UpdatePriceFeedFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *UpdatePriceFeedUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _UpdatePriceFeed.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpdatePriceFeedUpdatePrice)
				if err := _UpdatePriceFeed.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0x1a15ab7124a4e1ce00837351261771caf1691cd7d85ed3a0ac3157a1ee1a3805.
//
// Solidity: event UpdatePrice(uint256 price)
func (_UpdatePriceFeed *UpdatePriceFeedFilterer) ParseUpdatePrice(log types.Log) (*UpdatePriceFeedUpdatePrice, error) {
	event := new(UpdatePriceFeedUpdatePrice)
	if err := _UpdatePriceFeed.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
