// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalAccountCompressor

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

// CreditAccountData is an auto generated low-level Go binding around an user-defined struct.
type CreditAccountData struct {
	CreditAccount     common.Address
	CreditManager     common.Address
	CreditFacade      common.Address
	Underlying        common.Address
	Owner             common.Address
	ExpirationDate    *big.Int
	EnabledTokensMask *big.Int
	Debt              *big.Int
	AccruedInterest   *big.Int
	AccruedFees       *big.Int
	TotalDebtUSD      *big.Int
	TotalValueUSD     *big.Int
	TwvUSD            *big.Int
	TotalValue        *big.Int
	HealthFactor      *big.Int
	Success           bool
	Tokens            []TokenInfo
}

// CreditAccountFilter is an auto generated low-level Go binding around an user-defined struct.
type CreditAccountFilter struct {
	Owner           common.Address
	IncludeZeroDebt bool
	MinHealthFactor *big.Int
	MaxHealthFactor *big.Int
	Reverting       bool
}

// CreditManagerFilter is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerFilter struct {
	Configurators  []common.Address
	CreditManagers []common.Address
	Pools          []common.Address
	Underlying     common.Address
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	Token   common.Address
	Mask    *big.Int
	Balance *big.Int
	Quota   *big.Int
	Success bool
}

// GlobalAccountCompressorMetaData contains all meta data concerning the GlobalAccountCompressor contract.
var GlobalAccountCompressorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"addressProvider_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"countCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"countCreditAccounts\",\"inputs\":[{\"name\":\"cmFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccountData\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountData\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"cmFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"cmFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditManagerFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"creditManagers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// GlobalAccountCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalAccountCompressorMetaData.ABI instead.
var GlobalAccountCompressorABI = GlobalAccountCompressorMetaData.ABI

// GlobalAccountCompressor is an auto generated Go binding around an Ethereum contract.
type GlobalAccountCompressor struct {
	GlobalAccountCompressorCaller     // Read-only binding to the contract
	GlobalAccountCompressorTransactor // Write-only binding to the contract
	GlobalAccountCompressorFilterer   // Log filterer for contract events
}

// GlobalAccountCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalAccountCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccountCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalAccountCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccountCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalAccountCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalAccountCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalAccountCompressorSession struct {
	Contract     *GlobalAccountCompressor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GlobalAccountCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalAccountCompressorCallerSession struct {
	Contract *GlobalAccountCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// GlobalAccountCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalAccountCompressorTransactorSession struct {
	Contract     *GlobalAccountCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// GlobalAccountCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalAccountCompressorRaw struct {
	Contract *GlobalAccountCompressor // Generic contract binding to access the raw methods on
}

// GlobalAccountCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalAccountCompressorCallerRaw struct {
	Contract *GlobalAccountCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalAccountCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalAccountCompressorTransactorRaw struct {
	Contract *GlobalAccountCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalAccountCompressor creates a new instance of GlobalAccountCompressor, bound to a specific deployed contract.
func NewGlobalAccountCompressor(address common.Address, backend bind.ContractBackend) (*GlobalAccountCompressor, error) {
	contract, err := bindGlobalAccountCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalAccountCompressor{GlobalAccountCompressorCaller: GlobalAccountCompressorCaller{contract: contract}, GlobalAccountCompressorTransactor: GlobalAccountCompressorTransactor{contract: contract}, GlobalAccountCompressorFilterer: GlobalAccountCompressorFilterer{contract: contract}}, nil
}

// NewGlobalAccountCompressorCaller creates a new read-only instance of GlobalAccountCompressor, bound to a specific deployed contract.
func NewGlobalAccountCompressorCaller(address common.Address, caller bind.ContractCaller) (*GlobalAccountCompressorCaller, error) {
	contract, err := bindGlobalAccountCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalAccountCompressorCaller{contract: contract}, nil
}

// NewGlobalAccountCompressorTransactor creates a new write-only instance of GlobalAccountCompressor, bound to a specific deployed contract.
func NewGlobalAccountCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalAccountCompressorTransactor, error) {
	contract, err := bindGlobalAccountCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalAccountCompressorTransactor{contract: contract}, nil
}

// NewGlobalAccountCompressorFilterer creates a new log filterer instance of GlobalAccountCompressor, bound to a specific deployed contract.
func NewGlobalAccountCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalAccountCompressorFilterer, error) {
	contract, err := bindGlobalAccountCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalAccountCompressorFilterer{contract: contract}, nil
}

// bindGlobalAccountCompressor binds a generic wrapper to an already deployed contract.
func bindGlobalAccountCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlobalAccountCompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalAccountCompressor *GlobalAccountCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalAccountCompressor.Contract.GlobalAccountCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalAccountCompressor *GlobalAccountCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalAccountCompressor.Contract.GlobalAccountCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalAccountCompressor *GlobalAccountCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalAccountCompressor.Contract.GlobalAccountCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalAccountCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalAccountCompressor *GlobalAccountCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalAccountCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalAccountCompressor *GlobalAccountCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalAccountCompressor.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) AddressProvider() (common.Address, error) {
	return _GlobalAccountCompressor.Contract.AddressProvider(&_GlobalAccountCompressor.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) AddressProvider() (common.Address, error) {
	return _GlobalAccountCompressor.Contract.AddressProvider(&_GlobalAccountCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) ContractType() ([32]byte, error) {
	return _GlobalAccountCompressor.Contract.ContractType(&_GlobalAccountCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) ContractType() ([32]byte, error) {
	return _GlobalAccountCompressor.Contract.ContractType(&_GlobalAccountCompressor.CallOpts)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x548ca6d0.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) CountCreditAccounts(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "countCreditAccounts", creditManager, caFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x548ca6d0.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) CountCreditAccounts(creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.CountCreditAccounts(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x548ca6d0.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) CountCreditAccounts(creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.CountCreditAccounts(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter)
}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0xa6ecd4d5.
//
// Solidity: function countCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) CountCreditAccounts0(opts *bind.CallOpts, cmFilter CreditManagerFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "countCreditAccounts0", cmFilter, caFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0xa6ecd4d5.
//
// Solidity: function countCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) CountCreditAccounts0(cmFilter CreditManagerFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.CountCreditAccounts0(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter)
}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0xa6ecd4d5.
//
// Solidity: function countCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter) view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) CountCreditAccounts0(cmFilter CreditManagerFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.CountCreditAccounts0(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[]))
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) GetCreditAccountData(opts *bind.CallOpts, creditAccount common.Address) (CreditAccountData, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "getCreditAccountData", creditAccount)

	if err != nil {
		return *new(CreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditAccountData)).(*CreditAccountData)

	return out0, err

}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[]))
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) GetCreditAccountData(creditAccount common.Address) (CreditAccountData, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccountData(&_GlobalAccountCompressor.CallOpts, creditAccount)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[]))
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) GetCreditAccountData(creditAccount common.Address) (CreditAccountData, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccountData(&_GlobalAccountCompressor.CallOpts, creditAccount)
}

// GetCreditAccounts is a free data retrieval call binding the contract method 0x8b59b911.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) GetCreditAccounts(opts *bind.CallOpts, cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "getCreditAccounts", cmFilter, caFilter, offset, limit)

	outstruct := new(struct {
		Data       []CreditAccountData
		NextOffset *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]CreditAccountData)).(*[]CreditAccountData)
	outstruct.NextOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccounts is a free data retrieval call binding the contract method 0x8b59b911.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) GetCreditAccounts(cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter, offset, limit)
}

// GetCreditAccounts is a free data retrieval call binding the contract method 0x8b59b911.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) GetCreditAccounts(cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter, offset, limit)
}

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0xc4ab5af7.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) GetCreditAccounts0(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "getCreditAccounts0", creditManager, caFilter, offset)

	outstruct := new(struct {
		Data       []CreditAccountData
		NextOffset *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]CreditAccountData)).(*[]CreditAccountData)
	outstruct.NextOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0xc4ab5af7.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) GetCreditAccounts0(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts0(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter, offset)
}

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0xc4ab5af7.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) GetCreditAccounts0(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts0(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter, offset)
}

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xcd5a836f.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) GetCreditAccounts1(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "getCreditAccounts1", creditManager, caFilter, offset, limit)

	outstruct := new(struct {
		Data       []CreditAccountData
		NextOffset *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]CreditAccountData)).(*[]CreditAccountData)
	outstruct.NextOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xcd5a836f.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) GetCreditAccounts1(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts1(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter, offset, limit)
}

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xcd5a836f.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint256,uint256,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) GetCreditAccounts1(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts1(&_GlobalAccountCompressor.CallOpts, creditManager, caFilter, offset, limit)
}

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xf43bdb34.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) GetCreditAccounts2(opts *bind.CallOpts, cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "getCreditAccounts2", cmFilter, caFilter, offset)

	outstruct := new(struct {
		Data       []CreditAccountData
		NextOffset *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]CreditAccountData)).(*[]CreditAccountData)
	outstruct.NextOffset = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xf43bdb34.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) GetCreditAccounts2(cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts2(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter, offset)
}

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xf43bdb34.
//
// Solidity: function getCreditAccounts((address[],address[],address[],address) cmFilter, (address,bool,uint256,uint256,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) GetCreditAccounts2(cmFilter CreditManagerFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _GlobalAccountCompressor.Contract.GetCreditAccounts2(&_GlobalAccountCompressor.CallOpts, cmFilter, caFilter, offset)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlobalAccountCompressor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorSession) Version() (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.Version(&_GlobalAccountCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_GlobalAccountCompressor *GlobalAccountCompressorCallerSession) Version() (*big.Int, error) {
	return _GlobalAccountCompressor.Contract.Version(&_GlobalAccountCompressor.CallOpts)
}
