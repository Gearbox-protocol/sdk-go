// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapConnectorChecker

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
)

// UniswapConnectorCheckerMetaData contains all meta data concerning the UniswapConnectorChecker contract.
var UniswapConnectorCheckerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"connectorToken0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken4\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken5\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken6\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken7\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken8\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectorToken9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConnectors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"connectors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isConnector\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConnectors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapConnectorCheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapConnectorCheckerMetaData.ABI instead.
var UniswapConnectorCheckerABI = UniswapConnectorCheckerMetaData.ABI

// UniswapConnectorChecker is an auto generated Go binding around an Ethereum contract.
type UniswapConnectorChecker struct {
	UniswapConnectorCheckerCaller     // Read-only binding to the contract
	UniswapConnectorCheckerTransactor // Write-only binding to the contract
	UniswapConnectorCheckerFilterer   // Log filterer for contract events
}

// UniswapConnectorCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapConnectorCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapConnectorCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapConnectorCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapConnectorCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapConnectorCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapConnectorCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapConnectorCheckerSession struct {
	Contract     *UniswapConnectorChecker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniswapConnectorCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapConnectorCheckerCallerSession struct {
	Contract *UniswapConnectorCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// UniswapConnectorCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapConnectorCheckerTransactorSession struct {
	Contract     *UniswapConnectorCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// UniswapConnectorCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapConnectorCheckerRaw struct {
	Contract *UniswapConnectorChecker // Generic contract binding to access the raw methods on
}

// UniswapConnectorCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapConnectorCheckerCallerRaw struct {
	Contract *UniswapConnectorCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapConnectorCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapConnectorCheckerTransactorRaw struct {
	Contract *UniswapConnectorCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapConnectorChecker creates a new instance of UniswapConnectorChecker, bound to a specific deployed contract.
func NewUniswapConnectorChecker(address common.Address, backend bind.ContractBackend) (*UniswapConnectorChecker, error) {
	contract, err := bindUniswapConnectorChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapConnectorChecker{UniswapConnectorCheckerCaller: UniswapConnectorCheckerCaller{contract: contract}, UniswapConnectorCheckerTransactor: UniswapConnectorCheckerTransactor{contract: contract}, UniswapConnectorCheckerFilterer: UniswapConnectorCheckerFilterer{contract: contract}}, nil
}

// NewUniswapConnectorCheckerCaller creates a new read-only instance of UniswapConnectorChecker, bound to a specific deployed contract.
func NewUniswapConnectorCheckerCaller(address common.Address, caller bind.ContractCaller) (*UniswapConnectorCheckerCaller, error) {
	contract, err := bindUniswapConnectorChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapConnectorCheckerCaller{contract: contract}, nil
}

// NewUniswapConnectorCheckerTransactor creates a new write-only instance of UniswapConnectorChecker, bound to a specific deployed contract.
func NewUniswapConnectorCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapConnectorCheckerTransactor, error) {
	contract, err := bindUniswapConnectorChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapConnectorCheckerTransactor{contract: contract}, nil
}

// NewUniswapConnectorCheckerFilterer creates a new log filterer instance of UniswapConnectorChecker, bound to a specific deployed contract.
func NewUniswapConnectorCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapConnectorCheckerFilterer, error) {
	contract, err := bindUniswapConnectorChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapConnectorCheckerFilterer{contract: contract}, nil
}

// bindUniswapConnectorChecker binds a generic wrapper to an already deployed contract.
func bindUniswapConnectorChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapConnectorCheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapConnectorChecker *UniswapConnectorCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapConnectorChecker.Contract.UniswapConnectorCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapConnectorChecker *UniswapConnectorCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapConnectorChecker.Contract.UniswapConnectorCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapConnectorChecker *UniswapConnectorCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapConnectorChecker.Contract.UniswapConnectorCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapConnectorChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapConnectorChecker *UniswapConnectorCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapConnectorChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapConnectorChecker *UniswapConnectorCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapConnectorChecker.Contract.contract.Transact(opts, method, params...)
}

// ConnectorToken0 is a free data retrieval call binding the contract method 0x04711e2c.
//
// Solidity: function connectorToken0() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken0 is a free data retrieval call binding the contract method 0x04711e2c.
//
// Solidity: function connectorToken0() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken0() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken0(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken0 is a free data retrieval call binding the contract method 0x04711e2c.
//
// Solidity: function connectorToken0() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken0() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken0(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken1 is a free data retrieval call binding the contract method 0xf236ea0f.
//
// Solidity: function connectorToken1() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken1 is a free data retrieval call binding the contract method 0xf236ea0f.
//
// Solidity: function connectorToken1() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken1() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken1(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken1 is a free data retrieval call binding the contract method 0xf236ea0f.
//
// Solidity: function connectorToken1() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken1() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken1(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken2 is a free data retrieval call binding the contract method 0xb96bc265.
//
// Solidity: function connectorToken2() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken2 is a free data retrieval call binding the contract method 0xb96bc265.
//
// Solidity: function connectorToken2() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken2() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken2(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken2 is a free data retrieval call binding the contract method 0xb96bc265.
//
// Solidity: function connectorToken2() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken2() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken2(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken3 is a free data retrieval call binding the contract method 0x3f01665c.
//
// Solidity: function connectorToken3() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken3 is a free data retrieval call binding the contract method 0x3f01665c.
//
// Solidity: function connectorToken3() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken3() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken3(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken3 is a free data retrieval call binding the contract method 0x3f01665c.
//
// Solidity: function connectorToken3() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken3() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken3(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken4 is a free data retrieval call binding the contract method 0x427b608d.
//
// Solidity: function connectorToken4() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken4 is a free data retrieval call binding the contract method 0x427b608d.
//
// Solidity: function connectorToken4() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken4() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken4(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken4 is a free data retrieval call binding the contract method 0x427b608d.
//
// Solidity: function connectorToken4() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken4() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken4(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken5 is a free data retrieval call binding the contract method 0x62fe4283.
//
// Solidity: function connectorToken5() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken5(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken5")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken5 is a free data retrieval call binding the contract method 0x62fe4283.
//
// Solidity: function connectorToken5() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken5() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken5(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken5 is a free data retrieval call binding the contract method 0x62fe4283.
//
// Solidity: function connectorToken5() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken5() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken5(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken6 is a free data retrieval call binding the contract method 0x57dcfca4.
//
// Solidity: function connectorToken6() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken6(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken6")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken6 is a free data retrieval call binding the contract method 0x57dcfca4.
//
// Solidity: function connectorToken6() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken6() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken6(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken6 is a free data retrieval call binding the contract method 0x57dcfca4.
//
// Solidity: function connectorToken6() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken6() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken6(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken7 is a free data retrieval call binding the contract method 0xc4ea156d.
//
// Solidity: function connectorToken7() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken7(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken7")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken7 is a free data retrieval call binding the contract method 0xc4ea156d.
//
// Solidity: function connectorToken7() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken7() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken7(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken7 is a free data retrieval call binding the contract method 0xc4ea156d.
//
// Solidity: function connectorToken7() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken7() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken7(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken8 is a free data retrieval call binding the contract method 0xd2ceed8d.
//
// Solidity: function connectorToken8() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken8(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken8")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken8 is a free data retrieval call binding the contract method 0xd2ceed8d.
//
// Solidity: function connectorToken8() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken8() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken8(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken8 is a free data retrieval call binding the contract method 0xd2ceed8d.
//
// Solidity: function connectorToken8() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken8() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken8(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken9 is a free data retrieval call binding the contract method 0x017eab50.
//
// Solidity: function connectorToken9() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) ConnectorToken9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "connectorToken9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorToken9 is a free data retrieval call binding the contract method 0x017eab50.
//
// Solidity: function connectorToken9() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) ConnectorToken9() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken9(&_UniswapConnectorChecker.CallOpts)
}

// ConnectorToken9 is a free data retrieval call binding the contract method 0x017eab50.
//
// Solidity: function connectorToken9() view returns(address)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) ConnectorToken9() (common.Address, error) {
	return _UniswapConnectorChecker.Contract.ConnectorToken9(&_UniswapConnectorChecker.CallOpts)
}

// GetConnectors is a free data retrieval call binding the contract method 0xc74eaa11.
//
// Solidity: function getConnectors() view returns(address[] connectors)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) GetConnectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "getConnectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConnectors is a free data retrieval call binding the contract method 0xc74eaa11.
//
// Solidity: function getConnectors() view returns(address[] connectors)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) GetConnectors() ([]common.Address, error) {
	return _UniswapConnectorChecker.Contract.GetConnectors(&_UniswapConnectorChecker.CallOpts)
}

// GetConnectors is a free data retrieval call binding the contract method 0xc74eaa11.
//
// Solidity: function getConnectors() view returns(address[] connectors)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) GetConnectors() ([]common.Address, error) {
	return _UniswapConnectorChecker.Contract.GetConnectors(&_UniswapConnectorChecker.CallOpts)
}

// IsConnector is a free data retrieval call binding the contract method 0x47f37d1d.
//
// Solidity: function isConnector(address token) view returns(bool)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) IsConnector(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "isConnector", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConnector is a free data retrieval call binding the contract method 0x47f37d1d.
//
// Solidity: function isConnector(address token) view returns(bool)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) IsConnector(token common.Address) (bool, error) {
	return _UniswapConnectorChecker.Contract.IsConnector(&_UniswapConnectorChecker.CallOpts, token)
}

// IsConnector is a free data retrieval call binding the contract method 0x47f37d1d.
//
// Solidity: function isConnector(address token) view returns(bool)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) IsConnector(token common.Address) (bool, error) {
	return _UniswapConnectorChecker.Contract.IsConnector(&_UniswapConnectorChecker.CallOpts, token)
}

// NumConnectors is a free data retrieval call binding the contract method 0x67015b4c.
//
// Solidity: function numConnectors() view returns(uint256)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCaller) NumConnectors(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapConnectorChecker.contract.Call(opts, &out, "numConnectors")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumConnectors is a free data retrieval call binding the contract method 0x67015b4c.
//
// Solidity: function numConnectors() view returns(uint256)
func (_UniswapConnectorChecker *UniswapConnectorCheckerSession) NumConnectors() (*big.Int, error) {
	return _UniswapConnectorChecker.Contract.NumConnectors(&_UniswapConnectorChecker.CallOpts)
}

// NumConnectors is a free data retrieval call binding the contract method 0x67015b4c.
//
// Solidity: function numConnectors() view returns(uint256)
func (_UniswapConnectorChecker *UniswapConnectorCheckerCallerSession) NumConnectors() (*big.Int, error) {
	return _UniswapConnectorChecker.Contract.NumConnectors(&_UniswapConnectorChecker.CallOpts)
}
