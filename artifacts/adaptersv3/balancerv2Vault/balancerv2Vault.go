// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerv2Vault

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

// BatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type BatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// ExitPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type ExitPoolRequest struct {
	Assets            []common.Address
	MinAmountsOut     []*big.Int
	UserData          []byte
	ToInternalBalance bool
}

// FundManagement is an auto generated low-level Go binding around an user-defined struct.
type FundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// JoinPoolRequest is an auto generated low-level Go binding around an user-defined struct.
type JoinPoolRequest struct {
	Assets              []common.Address
	MaxAmountsIn        []*big.Int
	UserData            []byte
	FromInternalBalance bool
}

// SingleSwap is an auto generated low-level Go binding around an user-defined struct.
type SingleSwap struct {
	PoolId   [32]byte
	Kind     uint8
	AssetIn  common.Address
	AssetOut common.Address
	Amount   *big.Int
	UserData []byte
}

// SingleSwapDiff is an auto generated low-level Go binding around an user-defined struct.
type SingleSwapDiff struct {
	PoolId         [32]byte
	LeftoverAmount *big.Int
	AssetIn        common.Address
	AssetOut       common.Address
	UserData       []byte
}

// Balancerv2VaultMetaData contains all meta data concerning the Balancerv2Vault contract.
var Balancerv2VaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumAdapterType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_gearboxAdapterVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acl\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSwap\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumSwapKind\"},{\"name\":\"swaps\",\"type\":\"tuple[]\",\"internalType\":\"structBatchSwapStep[]\",\"components\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetInIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetOutIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFundManagement\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"toInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"limits\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"creditManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exitPool\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structExitPoolRequest\",\"components\":[{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"minAmountsOut\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitPoolSingleAsset\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitPoolSingleAssetDiff\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minRateRAY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinPool\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structJoinPoolRequest\",\"components\":[{\"name\":\"assets\",\"type\":\"address[]\",\"internalType\":\"contractIAsset[]\"},{\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fromInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinPoolSingleAsset\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetIn\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"joinPoolSingleAssetDiff\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assetIn\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minRateRAY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"poolStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumPoolStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPoolStatus\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"internalType\":\"enumPoolStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"singleSwap\",\"type\":\"tuple\",\"internalType\":\"structSingleSwap\",\"components\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumSwapKind\"},{\"name\":\"assetIn\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFundManagement\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"toInternalBalance\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapDiff\",\"inputs\":[{\"name\":\"singleSwapDiff\",\"type\":\"tuple\",\"internalType\":\"structSingleSwapDiff\",\"components\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"leftoverAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"assetIn\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"assetOut\",\"type\":\"address\",\"internalType\":\"contractIAsset\"},{\"name\":\"userData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"limitRateRAY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokensToEnable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensToDisable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SetPoolStatus\",\"inputs\":[{\"name\":\"poolId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumPoolStatus\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotConfiguratorException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotCreditFacadeException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolNotSupportedException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// Balancerv2VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use Balancerv2VaultMetaData.ABI instead.
var Balancerv2VaultABI = Balancerv2VaultMetaData.ABI

// Balancerv2Vault is an auto generated Go binding around an Ethereum contract.
type Balancerv2Vault struct {
	Balancerv2VaultCaller     // Read-only binding to the contract
	Balancerv2VaultTransactor // Write-only binding to the contract
	Balancerv2VaultFilterer   // Log filterer for contract events
}

// Balancerv2VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type Balancerv2VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Balancerv2VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Balancerv2VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Balancerv2VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Balancerv2VaultSession struct {
	Contract     *Balancerv2Vault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Balancerv2VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Balancerv2VaultCallerSession struct {
	Contract *Balancerv2VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Balancerv2VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Balancerv2VaultTransactorSession struct {
	Contract     *Balancerv2VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Balancerv2VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type Balancerv2VaultRaw struct {
	Contract *Balancerv2Vault // Generic contract binding to access the raw methods on
}

// Balancerv2VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Balancerv2VaultCallerRaw struct {
	Contract *Balancerv2VaultCaller // Generic read-only contract binding to access the raw methods on
}

// Balancerv2VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Balancerv2VaultTransactorRaw struct {
	Contract *Balancerv2VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerv2Vault creates a new instance of Balancerv2Vault, bound to a specific deployed contract.
func NewBalancerv2Vault(address common.Address, backend bind.ContractBackend) (*Balancerv2Vault, error) {
	contract, err := bindBalancerv2Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerv2Vault{Balancerv2VaultCaller: Balancerv2VaultCaller{contract: contract}, Balancerv2VaultTransactor: Balancerv2VaultTransactor{contract: contract}, Balancerv2VaultFilterer: Balancerv2VaultFilterer{contract: contract}}, nil
}

// NewBalancerv2VaultCaller creates a new read-only instance of Balancerv2Vault, bound to a specific deployed contract.
func NewBalancerv2VaultCaller(address common.Address, caller bind.ContractCaller) (*Balancerv2VaultCaller, error) {
	contract, err := bindBalancerv2Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2VaultCaller{contract: contract}, nil
}

// NewBalancerv2VaultTransactor creates a new write-only instance of Balancerv2Vault, bound to a specific deployed contract.
func NewBalancerv2VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*Balancerv2VaultTransactor, error) {
	contract, err := bindBalancerv2Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Balancerv2VaultTransactor{contract: contract}, nil
}

// NewBalancerv2VaultFilterer creates a new log filterer instance of Balancerv2Vault, bound to a specific deployed contract.
func NewBalancerv2VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*Balancerv2VaultFilterer, error) {
	contract, err := bindBalancerv2Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Balancerv2VaultFilterer{contract: contract}, nil
}

// bindBalancerv2Vault binds a generic wrapper to an already deployed contract.
func bindBalancerv2Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Balancerv2VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2Vault *Balancerv2VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2Vault.Contract.Balancerv2VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2Vault *Balancerv2VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.Balancerv2VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2Vault *Balancerv2VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.Balancerv2VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerv2Vault *Balancerv2VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerv2Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerv2Vault *Balancerv2VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerv2Vault *Balancerv2VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.contract.Transact(opts, method, params...)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultSession) GearboxAdapterType() (uint8, error) {
	return _Balancerv2Vault.Contract.GearboxAdapterType(&_Balancerv2Vault.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) GearboxAdapterType() (uint8, error) {
	return _Balancerv2Vault.Contract.GearboxAdapterType(&_Balancerv2Vault.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Balancerv2Vault *Balancerv2VaultCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Balancerv2Vault *Balancerv2VaultSession) GearboxAdapterVersion() (uint16, error) {
	return _Balancerv2Vault.Contract.GearboxAdapterVersion(&_Balancerv2Vault.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _Balancerv2Vault.Contract.GearboxAdapterVersion(&_Balancerv2Vault.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCaller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultSession) Acl() (common.Address, error) {
	return _Balancerv2Vault.Contract.Acl(&_Balancerv2Vault.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) Acl() (common.Address, error) {
	return _Balancerv2Vault.Contract.Acl(&_Balancerv2Vault.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultSession) AddressProvider() (common.Address, error) {
	return _Balancerv2Vault.Contract.AddressProvider(&_Balancerv2Vault.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) AddressProvider() (common.Address, error) {
	return _Balancerv2Vault.Contract.AddressProvider(&_Balancerv2Vault.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultSession) CreditManager() (common.Address, error) {
	return _Balancerv2Vault.Contract.CreditManager(&_Balancerv2Vault.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) CreditManager() (common.Address, error) {
	return _Balancerv2Vault.Contract.CreditManager(&_Balancerv2Vault.CallOpts)
}

// PoolStatus is a free data retrieval call binding the contract method 0x066ada90.
//
// Solidity: function poolStatus(bytes32 ) view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultCaller) PoolStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "poolStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolStatus is a free data retrieval call binding the contract method 0x066ada90.
//
// Solidity: function poolStatus(bytes32 ) view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultSession) PoolStatus(arg0 [32]byte) (uint8, error) {
	return _Balancerv2Vault.Contract.PoolStatus(&_Balancerv2Vault.CallOpts, arg0)
}

// PoolStatus is a free data retrieval call binding the contract method 0x066ada90.
//
// Solidity: function poolStatus(bytes32 ) view returns(uint8)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) PoolStatus(arg0 [32]byte) (uint8, error) {
	return _Balancerv2Vault.Contract.PoolStatus(&_Balancerv2Vault.CallOpts, arg0)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerv2Vault.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultSession) TargetContract() (common.Address, error) {
	return _Balancerv2Vault.Contract.TargetContract(&_Balancerv2Vault.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_Balancerv2Vault *Balancerv2VaultCallerSession) TargetContract() (common.Address, error) {
	return _Balancerv2Vault.Contract.TargetContract(&_Balancerv2Vault.CallOpts)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) , int256[] limits, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) BatchSwap(opts *bind.TransactOpts, kind uint8, swaps []BatchSwapStep, assets []common.Address, arg3 FundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "batchSwap", kind, swaps, assets, arg3, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) , int256[] limits, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) BatchSwap(kind uint8, swaps []BatchSwapStep, assets []common.Address, arg3 FundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.BatchSwap(&_Balancerv2Vault.TransactOpts, kind, swaps, assets, arg3, limits, deadline)
}

// BatchSwap is a paid mutator transaction binding the contract method 0x945bcec9.
//
// Solidity: function batchSwap(uint8 kind, (bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] assets, (address,bool,address,bool) , int256[] limits, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) BatchSwap(kind uint8, swaps []BatchSwapStep, assets []common.Address, arg3 FundManagement, limits []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.BatchSwap(&_Balancerv2Vault.TransactOpts, kind, swaps, assets, arg3, limits, deadline)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) ExitPool(opts *bind.TransactOpts, poolId [32]byte, arg1 common.Address, arg2 common.Address, request ExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "exitPool", poolId, arg1, arg2, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) ExitPool(poolId [32]byte, arg1 common.Address, arg2 common.Address, request ExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPool(&_Balancerv2Vault.TransactOpts, poolId, arg1, arg2, request)
}

// ExitPool is a paid mutator transaction binding the contract method 0x8bdb3913.
//
// Solidity: function exitPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) ExitPool(poolId [32]byte, arg1 common.Address, arg2 common.Address, request ExitPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPool(&_Balancerv2Vault.TransactOpts, poolId, arg1, arg2, request)
}

// ExitPoolSingleAsset is a paid mutator transaction binding the contract method 0xbc5a07df.
//
// Solidity: function exitPoolSingleAsset(bytes32 poolId, address assetOut, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) ExitPoolSingleAsset(opts *bind.TransactOpts, poolId [32]byte, assetOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "exitPoolSingleAsset", poolId, assetOut, amountIn, minAmountOut)
}

// ExitPoolSingleAsset is a paid mutator transaction binding the contract method 0xbc5a07df.
//
// Solidity: function exitPoolSingleAsset(bytes32 poolId, address assetOut, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) ExitPoolSingleAsset(poolId [32]byte, assetOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPoolSingleAsset(&_Balancerv2Vault.TransactOpts, poolId, assetOut, amountIn, minAmountOut)
}

// ExitPoolSingleAsset is a paid mutator transaction binding the contract method 0xbc5a07df.
//
// Solidity: function exitPoolSingleAsset(bytes32 poolId, address assetOut, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) ExitPoolSingleAsset(poolId [32]byte, assetOut common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPoolSingleAsset(&_Balancerv2Vault.TransactOpts, poolId, assetOut, amountIn, minAmountOut)
}

// ExitPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1a970270.
//
// Solidity: function exitPoolSingleAssetDiff(bytes32 poolId, address assetOut, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) ExitPoolSingleAssetDiff(opts *bind.TransactOpts, poolId [32]byte, assetOut common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "exitPoolSingleAssetDiff", poolId, assetOut, leftoverAmount, minRateRAY)
}

// ExitPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1a970270.
//
// Solidity: function exitPoolSingleAssetDiff(bytes32 poolId, address assetOut, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) ExitPoolSingleAssetDiff(poolId [32]byte, assetOut common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPoolSingleAssetDiff(&_Balancerv2Vault.TransactOpts, poolId, assetOut, leftoverAmount, minRateRAY)
}

// ExitPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1a970270.
//
// Solidity: function exitPoolSingleAssetDiff(bytes32 poolId, address assetOut, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) ExitPoolSingleAssetDiff(poolId [32]byte, assetOut common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.ExitPoolSingleAssetDiff(&_Balancerv2Vault.TransactOpts, poolId, assetOut, leftoverAmount, minRateRAY)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) JoinPool(opts *bind.TransactOpts, poolId [32]byte, arg1 common.Address, arg2 common.Address, request JoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "joinPool", poolId, arg1, arg2, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) JoinPool(poolId [32]byte, arg1 common.Address, arg2 common.Address, request JoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPool(&_Balancerv2Vault.TransactOpts, poolId, arg1, arg2, request)
}

// JoinPool is a paid mutator transaction binding the contract method 0xb95cac28.
//
// Solidity: function joinPool(bytes32 poolId, address , address , (address[],uint256[],bytes,bool) request) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) JoinPool(poolId [32]byte, arg1 common.Address, arg2 common.Address, request JoinPoolRequest) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPool(&_Balancerv2Vault.TransactOpts, poolId, arg1, arg2, request)
}

// JoinPoolSingleAsset is a paid mutator transaction binding the contract method 0x72a802ce.
//
// Solidity: function joinPoolSingleAsset(bytes32 poolId, address assetIn, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) JoinPoolSingleAsset(opts *bind.TransactOpts, poolId [32]byte, assetIn common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "joinPoolSingleAsset", poolId, assetIn, amountIn, minAmountOut)
}

// JoinPoolSingleAsset is a paid mutator transaction binding the contract method 0x72a802ce.
//
// Solidity: function joinPoolSingleAsset(bytes32 poolId, address assetIn, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) JoinPoolSingleAsset(poolId [32]byte, assetIn common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPoolSingleAsset(&_Balancerv2Vault.TransactOpts, poolId, assetIn, amountIn, minAmountOut)
}

// JoinPoolSingleAsset is a paid mutator transaction binding the contract method 0x72a802ce.
//
// Solidity: function joinPoolSingleAsset(bytes32 poolId, address assetIn, uint256 amountIn, uint256 minAmountOut) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) JoinPoolSingleAsset(poolId [32]byte, assetIn common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPoolSingleAsset(&_Balancerv2Vault.TransactOpts, poolId, assetIn, amountIn, minAmountOut)
}

// JoinPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1cc04018.
//
// Solidity: function joinPoolSingleAssetDiff(bytes32 poolId, address assetIn, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) JoinPoolSingleAssetDiff(opts *bind.TransactOpts, poolId [32]byte, assetIn common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "joinPoolSingleAssetDiff", poolId, assetIn, leftoverAmount, minRateRAY)
}

// JoinPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1cc04018.
//
// Solidity: function joinPoolSingleAssetDiff(bytes32 poolId, address assetIn, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) JoinPoolSingleAssetDiff(poolId [32]byte, assetIn common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPoolSingleAssetDiff(&_Balancerv2Vault.TransactOpts, poolId, assetIn, leftoverAmount, minRateRAY)
}

// JoinPoolSingleAssetDiff is a paid mutator transaction binding the contract method 0x1cc04018.
//
// Solidity: function joinPoolSingleAssetDiff(bytes32 poolId, address assetIn, uint256 leftoverAmount, uint256 minRateRAY) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) JoinPoolSingleAssetDiff(poolId [32]byte, assetIn common.Address, leftoverAmount *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.JoinPoolSingleAssetDiff(&_Balancerv2Vault.TransactOpts, poolId, assetIn, leftoverAmount, minRateRAY)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x8f4c6ec5.
//
// Solidity: function setPoolStatus(bytes32 poolId, uint8 newStatus) returns()
func (_Balancerv2Vault *Balancerv2VaultTransactor) SetPoolStatus(opts *bind.TransactOpts, poolId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "setPoolStatus", poolId, newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x8f4c6ec5.
//
// Solidity: function setPoolStatus(bytes32 poolId, uint8 newStatus) returns()
func (_Balancerv2Vault *Balancerv2VaultSession) SetPoolStatus(poolId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.SetPoolStatus(&_Balancerv2Vault.TransactOpts, poolId, newStatus)
}

// SetPoolStatus is a paid mutator transaction binding the contract method 0x8f4c6ec5.
//
// Solidity: function setPoolStatus(bytes32 poolId, uint8 newStatus) returns()
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) SetPoolStatus(poolId [32]byte, newStatus uint8) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.SetPoolStatus(&_Balancerv2Vault.TransactOpts, poolId, newStatus)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) , uint256 limit, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) Swap(opts *bind.TransactOpts, singleSwap SingleSwap, arg1 FundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "swap", singleSwap, arg1, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) , uint256 limit, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) Swap(singleSwap SingleSwap, arg1 FundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.Swap(&_Balancerv2Vault.TransactOpts, singleSwap, arg1, limit, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x52bbbe29.
//
// Solidity: function swap((bytes32,uint8,address,address,uint256,bytes) singleSwap, (address,bool,address,bool) , uint256 limit, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) Swap(singleSwap SingleSwap, arg1 FundManagement, limit *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.Swap(&_Balancerv2Vault.TransactOpts, singleSwap, arg1, limit, deadline)
}

// SwapDiff is a paid mutator transaction binding the contract method 0xd1b5797e.
//
// Solidity: function swapDiff((bytes32,uint256,address,address,bytes) singleSwapDiff, uint256 limitRateRAY, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactor) SwapDiff(opts *bind.TransactOpts, singleSwapDiff SingleSwapDiff, limitRateRAY *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.contract.Transact(opts, "swapDiff", singleSwapDiff, limitRateRAY, deadline)
}

// SwapDiff is a paid mutator transaction binding the contract method 0xd1b5797e.
//
// Solidity: function swapDiff((bytes32,uint256,address,address,bytes) singleSwapDiff, uint256 limitRateRAY, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultSession) SwapDiff(singleSwapDiff SingleSwapDiff, limitRateRAY *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.SwapDiff(&_Balancerv2Vault.TransactOpts, singleSwapDiff, limitRateRAY, deadline)
}

// SwapDiff is a paid mutator transaction binding the contract method 0xd1b5797e.
//
// Solidity: function swapDiff((bytes32,uint256,address,address,bytes) singleSwapDiff, uint256 limitRateRAY, uint256 deadline) returns(uint256 tokensToEnable, uint256 tokensToDisable)
func (_Balancerv2Vault *Balancerv2VaultTransactorSession) SwapDiff(singleSwapDiff SingleSwapDiff, limitRateRAY *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Balancerv2Vault.Contract.SwapDiff(&_Balancerv2Vault.TransactOpts, singleSwapDiff, limitRateRAY, deadline)
}

// Balancerv2VaultSetPoolStatusIterator is returned from FilterSetPoolStatus and is used to iterate over the raw logs and unpacked data for SetPoolStatus events raised by the Balancerv2Vault contract.
type Balancerv2VaultSetPoolStatusIterator struct {
	Event *Balancerv2VaultSetPoolStatus // Event containing the contract specifics and raw log

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
func (it *Balancerv2VaultSetPoolStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Balancerv2VaultSetPoolStatus)
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
		it.Event = new(Balancerv2VaultSetPoolStatus)
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
func (it *Balancerv2VaultSetPoolStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Balancerv2VaultSetPoolStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Balancerv2VaultSetPoolStatus represents a SetPoolStatus event raised by the Balancerv2Vault contract.
type Balancerv2VaultSetPoolStatus struct {
	PoolId    [32]byte
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetPoolStatus is a free log retrieval operation binding the contract event 0xcb31053d9df846999fdad8189e022e9af1c2069b37a00077a5e0ac41ea25bbea.
//
// Solidity: event SetPoolStatus(bytes32 indexed poolId, uint8 newStatus)
func (_Balancerv2Vault *Balancerv2VaultFilterer) FilterSetPoolStatus(opts *bind.FilterOpts, poolId [][32]byte) (*Balancerv2VaultSetPoolStatusIterator, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2Vault.contract.FilterLogs(opts, "SetPoolStatus", poolIdRule)
	if err != nil {
		return nil, err
	}
	return &Balancerv2VaultSetPoolStatusIterator{contract: _Balancerv2Vault.contract, event: "SetPoolStatus", logs: logs, sub: sub}, nil
}

// WatchSetPoolStatus is a free log subscription operation binding the contract event 0xcb31053d9df846999fdad8189e022e9af1c2069b37a00077a5e0ac41ea25bbea.
//
// Solidity: event SetPoolStatus(bytes32 indexed poolId, uint8 newStatus)
func (_Balancerv2Vault *Balancerv2VaultFilterer) WatchSetPoolStatus(opts *bind.WatchOpts, sink chan<- *Balancerv2VaultSetPoolStatus, poolId [][32]byte) (event.Subscription, error) {

	var poolIdRule []interface{}
	for _, poolIdItem := range poolId {
		poolIdRule = append(poolIdRule, poolIdItem)
	}

	logs, sub, err := _Balancerv2Vault.contract.WatchLogs(opts, "SetPoolStatus", poolIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Balancerv2VaultSetPoolStatus)
				if err := _Balancerv2Vault.contract.UnpackLog(event, "SetPoolStatus", log); err != nil {
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

// ParseSetPoolStatus is a log parse operation binding the contract event 0xcb31053d9df846999fdad8189e022e9af1c2069b37a00077a5e0ac41ea25bbea.
//
// Solidity: event SetPoolStatus(bytes32 indexed poolId, uint8 newStatus)
func (_Balancerv2Vault *Balancerv2VaultFilterer) ParseSetPoolStatus(log types.Log) (*Balancerv2VaultSetPoolStatus, error) {
	event := new(Balancerv2VaultSetPoolStatus)
	if err := _Balancerv2Vault.contract.UnpackLog(event, "SetPoolStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
