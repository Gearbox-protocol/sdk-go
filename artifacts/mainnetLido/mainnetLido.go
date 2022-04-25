// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainnetLido

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

// MainnetLidoABI is the input ABI used to generate the binding from.
const MainnetLidoABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"proxyType\",\"outputs\":[{\"name\":\"proxyTypeId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isDepositable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kernel\",\"type\":\"address\"},{\"name\":\"_appId\",\"type\":\"bytes32\"},{\"name\":\"_initializePayload\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ProxyDeposit\",\"type\":\"event\"}]"

// MainnetLido is an auto generated Go binding around an Ethereum contract.
type MainnetLido struct {
	MainnetLidoCaller     // Read-only binding to the contract
	MainnetLidoTransactor // Write-only binding to the contract
	MainnetLidoFilterer   // Log filterer for contract events
}

// MainnetLidoCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnetLidoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnetLidoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnetLidoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetLidoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnetLidoSession struct {
	Contract     *MainnetLido      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainnetLidoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnetLidoCallerSession struct {
	Contract *MainnetLidoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MainnetLidoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnetLidoTransactorSession struct {
	Contract     *MainnetLidoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MainnetLidoRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnetLidoRaw struct {
	Contract *MainnetLido // Generic contract binding to access the raw methods on
}

// MainnetLidoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnetLidoCallerRaw struct {
	Contract *MainnetLidoCaller // Generic read-only contract binding to access the raw methods on
}

// MainnetLidoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnetLidoTransactorRaw struct {
	Contract *MainnetLidoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnetLido creates a new instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLido(address common.Address, backend bind.ContractBackend) (*MainnetLido, error) {
	contract, err := bindMainnetLido(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainnetLido{MainnetLidoCaller: MainnetLidoCaller{contract: contract}, MainnetLidoTransactor: MainnetLidoTransactor{contract: contract}, MainnetLidoFilterer: MainnetLidoFilterer{contract: contract}}, nil
}

// NewMainnetLidoCaller creates a new read-only instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoCaller(address common.Address, caller bind.ContractCaller) (*MainnetLidoCaller, error) {
	contract, err := bindMainnetLido(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoCaller{contract: contract}, nil
}

// NewMainnetLidoTransactor creates a new write-only instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnetLidoTransactor, error) {
	contract, err := bindMainnetLido(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoTransactor{contract: contract}, nil
}

// NewMainnetLidoFilterer creates a new log filterer instance of MainnetLido, bound to a specific deployed contract.
func NewMainnetLidoFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnetLidoFilterer, error) {
	contract, err := bindMainnetLido(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnetLidoFilterer{contract: contract}, nil
}

// bindMainnetLido binds a generic wrapper to an already deployed contract.
func bindMainnetLido(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetLidoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetLido *MainnetLidoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainnetLido.Contract.MainnetLidoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetLido *MainnetLidoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetLido.Contract.MainnetLidoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetLido *MainnetLidoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetLido.Contract.MainnetLidoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetLido *MainnetLidoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainnetLido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetLido *MainnetLidoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetLido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetLido *MainnetLidoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetLido.Contract.contract.Transact(opts, method, params...)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_MainnetLido *MainnetLidoCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_MainnetLido *MainnetLidoSession) AppId() ([32]byte, error) {
	return _MainnetLido.Contract.AppId(&_MainnetLido.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_MainnetLido *MainnetLidoCallerSession) AppId() ([32]byte, error) {
	return _MainnetLido.Contract.AppId(&_MainnetLido.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MainnetLido *MainnetLidoCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MainnetLido *MainnetLidoSession) Implementation() (common.Address, error) {
	return _MainnetLido.Contract.Implementation(&_MainnetLido.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_MainnetLido *MainnetLidoCallerSession) Implementation() (common.Address, error) {
	return _MainnetLido.Contract.Implementation(&_MainnetLido.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_MainnetLido *MainnetLidoCaller) IsDepositable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "isDepositable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_MainnetLido *MainnetLidoSession) IsDepositable() (bool, error) {
	return _MainnetLido.Contract.IsDepositable(&_MainnetLido.CallOpts)
}

// IsDepositable is a free data retrieval call binding the contract method 0x48a0c8dd.
//
// Solidity: function isDepositable() view returns(bool)
func (_MainnetLido *MainnetLidoCallerSession) IsDepositable() (bool, error) {
	return _MainnetLido.Contract.IsDepositable(&_MainnetLido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_MainnetLido *MainnetLidoCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_MainnetLido *MainnetLidoSession) Kernel() (common.Address, error) {
	return _MainnetLido.Contract.Kernel(&_MainnetLido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_MainnetLido *MainnetLidoCallerSession) Kernel() (common.Address, error) {
	return _MainnetLido.Contract.Kernel(&_MainnetLido.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_MainnetLido *MainnetLidoCaller) ProxyType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainnetLido.contract.Call(opts, &out, "proxyType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_MainnetLido *MainnetLidoSession) ProxyType() (*big.Int, error) {
	return _MainnetLido.Contract.ProxyType(&_MainnetLido.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256 proxyTypeId)
func (_MainnetLido *MainnetLidoCallerSession) ProxyType() (*big.Int, error) {
	return _MainnetLido.Contract.ProxyType(&_MainnetLido.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MainnetLido *MainnetLidoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MainnetLido.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MainnetLido *MainnetLidoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MainnetLido.Contract.Fallback(&_MainnetLido.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MainnetLido *MainnetLidoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MainnetLido.Contract.Fallback(&_MainnetLido.TransactOpts, calldata)
}

// MainnetLidoProxyDepositIterator is returned from FilterProxyDeposit and is used to iterate over the raw logs and unpacked data for ProxyDeposit events raised by the MainnetLido contract.
type MainnetLidoProxyDepositIterator struct {
	Event *MainnetLidoProxyDeposit // Event containing the contract specifics and raw log

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
func (it *MainnetLidoProxyDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetLidoProxyDeposit)
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
		it.Event = new(MainnetLidoProxyDeposit)
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
func (it *MainnetLidoProxyDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetLidoProxyDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetLidoProxyDeposit represents a ProxyDeposit event raised by the MainnetLido contract.
type MainnetLidoProxyDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProxyDeposit is a free log retrieval operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) FilterProxyDeposit(opts *bind.FilterOpts) (*MainnetLidoProxyDepositIterator, error) {

	logs, sub, err := _MainnetLido.contract.FilterLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return &MainnetLidoProxyDepositIterator{contract: _MainnetLido.contract, event: "ProxyDeposit", logs: logs, sub: sub}, nil
}

// WatchProxyDeposit is a free log subscription operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) WatchProxyDeposit(opts *bind.WatchOpts, sink chan<- *MainnetLidoProxyDeposit) (event.Subscription, error) {

	logs, sub, err := _MainnetLido.contract.WatchLogs(opts, "ProxyDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetLidoProxyDeposit)
				if err := _MainnetLido.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
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

// ParseProxyDeposit is a log parse operation binding the contract event 0x15eeaa57c7bd188c1388020bcadc2c436ec60d647d36ef5b9eb3c742217ddee1.
//
// Solidity: event ProxyDeposit(address sender, uint256 value)
func (_MainnetLido *MainnetLidoFilterer) ParseProxyDeposit(log types.Log) (*MainnetLidoProxyDeposit, error) {
	event := new(MainnetLidoProxyDeposit)
	if err := _MainnetLido.contract.UnpackLog(event, "ProxyDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}
