// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditAccountCompressor

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
	HealthFactor      uint16
	Success           bool
	Tokens            []TokenInfo
}

// CreditAccountFilter is an auto generated low-level Go binding around an user-defined struct.
type CreditAccountFilter struct {
	Owner           common.Address
	IncludeZeroDebt bool
	MinHealthFactor uint16
	MaxHealthFactor uint16
	Reverting       bool
}

// MarketFilter is an auto generated low-level Go binding around an user-defined struct.
type MarketFilter struct {
	Configurators []common.Address
	Pools         []common.Address
	Underlying    common.Address
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	Token   common.Address
	Mask    *big.Int
	Balance *big.Int
	Quota   *big.Int
	Success bool
}

// CreditAccountCompressorMetaData contains all meta data concerning the CreditAccountCompressor contract.
var CreditAccountCompressorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"addressProvider_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addressProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"countCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"countCreditAccounts\",\"inputs\":[{\"name\":\"marketFilter\",\"type\":\"tuple\",\"internalType\":\"structMarketFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccountData\",\"inputs\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountData\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"marketFilter\",\"type\":\"tuple\",\"internalType\":\"structMarketFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCreditAccounts\",\"inputs\":[{\"name\":\"marketFilter\",\"type\":\"tuple\",\"internalType\":\"structMarketFilter\",\"components\":[{\"name\":\"configurators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"pools\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"caFilter\",\"type\":\"tuple\",\"internalType\":\"structCreditAccountFilter\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"includeZeroDebt\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxHealthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reverting\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"tuple[]\",\"internalType\":\"structCreditAccountData[]\",\"components\":[{\"name\":\"creditAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"creditFacade\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"underlying\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expirationDate\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"enabledTokensMask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedInterest\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accruedFees\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalDebtUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValueUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"twvUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structTokenInfo[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mask\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quota\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"nextOffset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketConfiguratorFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidAddressProviderException\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressException\",\"inputs\":[]}]",
}

// CreditAccountCompressorABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditAccountCompressorMetaData.ABI instead.
var CreditAccountCompressorABI = CreditAccountCompressorMetaData.ABI

// CreditAccountCompressor is an auto generated Go binding around an Ethereum contract.
type CreditAccountCompressor struct {
	CreditAccountCompressorCaller     // Read-only binding to the contract
	CreditAccountCompressorTransactor // Write-only binding to the contract
	CreditAccountCompressorFilterer   // Log filterer for contract events
}

// CreditAccountCompressorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditAccountCompressorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountCompressorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditAccountCompressorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountCompressorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditAccountCompressorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditAccountCompressorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditAccountCompressorSession struct {
	Contract     *CreditAccountCompressor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CreditAccountCompressorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditAccountCompressorCallerSession struct {
	Contract *CreditAccountCompressorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// CreditAccountCompressorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditAccountCompressorTransactorSession struct {
	Contract     *CreditAccountCompressorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// CreditAccountCompressorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditAccountCompressorRaw struct {
	Contract *CreditAccountCompressor // Generic contract binding to access the raw methods on
}

// CreditAccountCompressorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditAccountCompressorCallerRaw struct {
	Contract *CreditAccountCompressorCaller // Generic read-only contract binding to access the raw methods on
}

// CreditAccountCompressorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditAccountCompressorTransactorRaw struct {
	Contract *CreditAccountCompressorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditAccountCompressor creates a new instance of CreditAccountCompressor, bound to a specific deployed contract.
func NewCreditAccountCompressor(address common.Address, backend bind.ContractBackend) (*CreditAccountCompressor, error) {
	contract, err := bindCreditAccountCompressor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditAccountCompressor{CreditAccountCompressorCaller: CreditAccountCompressorCaller{contract: contract}, CreditAccountCompressorTransactor: CreditAccountCompressorTransactor{contract: contract}, CreditAccountCompressorFilterer: CreditAccountCompressorFilterer{contract: contract}}, nil
}

// NewCreditAccountCompressorCaller creates a new read-only instance of CreditAccountCompressor, bound to a specific deployed contract.
func NewCreditAccountCompressorCaller(address common.Address, caller bind.ContractCaller) (*CreditAccountCompressorCaller, error) {
	contract, err := bindCreditAccountCompressor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditAccountCompressorCaller{contract: contract}, nil
}

// NewCreditAccountCompressorTransactor creates a new write-only instance of CreditAccountCompressor, bound to a specific deployed contract.
func NewCreditAccountCompressorTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditAccountCompressorTransactor, error) {
	contract, err := bindCreditAccountCompressor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditAccountCompressorTransactor{contract: contract}, nil
}

// NewCreditAccountCompressorFilterer creates a new log filterer instance of CreditAccountCompressor, bound to a specific deployed contract.
func NewCreditAccountCompressorFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditAccountCompressorFilterer, error) {
	contract, err := bindCreditAccountCompressor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditAccountCompressorFilterer{contract: contract}, nil
}

// bindCreditAccountCompressor binds a generic wrapper to an already deployed contract.
func bindCreditAccountCompressor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditAccountCompressorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditAccountCompressor *CreditAccountCompressorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditAccountCompressor.Contract.CreditAccountCompressorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditAccountCompressor *CreditAccountCompressorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditAccountCompressor.Contract.CreditAccountCompressorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditAccountCompressor *CreditAccountCompressorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditAccountCompressor.Contract.CreditAccountCompressorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditAccountCompressor *CreditAccountCompressorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditAccountCompressor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditAccountCompressor *CreditAccountCompressorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditAccountCompressor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditAccountCompressor *CreditAccountCompressorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditAccountCompressor.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorSession) AddressProvider() (common.Address, error) {
	return _CreditAccountCompressor.Contract.AddressProvider(&_CreditAccountCompressor.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) AddressProvider() (common.Address, error) {
	return _CreditAccountCompressor.Contract.AddressProvider(&_CreditAccountCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditAccountCompressor *CreditAccountCompressorSession) ContractType() ([32]byte, error) {
	return _CreditAccountCompressor.Contract.ContractType(&_CreditAccountCompressor.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() view returns(bytes32)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) ContractType() ([32]byte, error) {
	return _CreditAccountCompressor.Contract.ContractType(&_CreditAccountCompressor.CallOpts)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x7849a81d.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) CountCreditAccounts(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "countCreditAccounts", creditManager, caFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x7849a81d.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorSession) CountCreditAccounts(creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	return _CreditAccountCompressor.Contract.CountCreditAccounts(&_CreditAccountCompressor.CallOpts, creditManager, caFilter)
}

// CountCreditAccounts is a free data retrieval call binding the contract method 0x7849a81d.
//
// Solidity: function countCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) CountCreditAccounts(creditManager common.Address, caFilter CreditAccountFilter) (*big.Int, error) {
	return _CreditAccountCompressor.Contract.CountCreditAccounts(&_CreditAccountCompressor.CallOpts, creditManager, caFilter)
}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0x9e62d5ef.
//
// Solidity: function countCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) CountCreditAccounts0(opts *bind.CallOpts, marketFilter MarketFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "countCreditAccounts0", marketFilter, caFilter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0x9e62d5ef.
//
// Solidity: function countCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorSession) CountCreditAccounts0(marketFilter MarketFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	return _CreditAccountCompressor.Contract.CountCreditAccounts0(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter)
}

// CountCreditAccounts0 is a free data retrieval call binding the contract method 0x9e62d5ef.
//
// Solidity: function countCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter) view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) CountCreditAccounts0(marketFilter MarketFilter, caFilter CreditAccountFilter) (*big.Int, error) {
	return _CreditAccountCompressor.Contract.CountCreditAccounts0(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[]))
func (_CreditAccountCompressor *CreditAccountCompressorCaller) GetCreditAccountData(opts *bind.CallOpts, creditAccount common.Address) (CreditAccountData, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "getCreditAccountData", creditAccount)

	if err != nil {
		return *new(CreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditAccountData)).(*CreditAccountData)

	return out0, err

}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[]))
func (_CreditAccountCompressor *CreditAccountCompressorSession) GetCreditAccountData(creditAccount common.Address) (CreditAccountData, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccountData(&_CreditAccountCompressor.CallOpts, creditAccount)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0xa595f79e.
//
// Solidity: function getCreditAccountData(address creditAccount) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[]))
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) GetCreditAccountData(creditAccount common.Address) (CreditAccountData, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccountData(&_CreditAccountCompressor.CallOpts, creditAccount)
}

// GetCreditAccounts is a free data retrieval call binding the contract method 0x21a54acc.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) GetCreditAccounts(opts *bind.CallOpts, marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "getCreditAccounts", marketFilter, caFilter, offset)

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

// GetCreditAccounts is a free data retrieval call binding the contract method 0x21a54acc.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorSession) GetCreditAccounts(marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter, offset)
}

// GetCreditAccounts is a free data retrieval call binding the contract method 0x21a54acc.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) GetCreditAccounts(marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter, offset)
}

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0x9243fed4.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) GetCreditAccounts0(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "getCreditAccounts0", creditManager, caFilter, offset)

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

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0x9243fed4.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorSession) GetCreditAccounts0(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts0(&_CreditAccountCompressor.CallOpts, creditManager, caFilter, offset)
}

// GetCreditAccounts0 is a free data retrieval call binding the contract method 0x9243fed4.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) GetCreditAccounts0(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts0(&_CreditAccountCompressor.CallOpts, creditManager, caFilter, offset)
}

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xbf515471.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) GetCreditAccounts1(opts *bind.CallOpts, creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "getCreditAccounts1", creditManager, caFilter, offset, limit)

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

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xbf515471.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorSession) GetCreditAccounts1(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts1(&_CreditAccountCompressor.CallOpts, creditManager, caFilter, offset, limit)
}

// GetCreditAccounts1 is a free data retrieval call binding the contract method 0xbf515471.
//
// Solidity: function getCreditAccounts(address creditManager, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) GetCreditAccounts1(creditManager common.Address, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts1(&_CreditAccountCompressor.CallOpts, creditManager, caFilter, offset, limit)
}

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xef9ac04c.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) GetCreditAccounts2(opts *bind.CallOpts, marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "getCreditAccounts2", marketFilter, caFilter, offset, limit)

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

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xef9ac04c.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorSession) GetCreditAccounts2(marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts2(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter, offset, limit)
}

// GetCreditAccounts2 is a free data retrieval call binding the contract method 0xef9ac04c.
//
// Solidity: function getCreditAccounts((address[],address[],address) marketFilter, (address,bool,uint16,uint16,bool) caFilter, uint256 offset, uint256 limit) view returns((address,address,address,address,address,uint40,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint16,bool,(address,uint256,uint256,uint256,bool)[])[] data, uint256 nextOffset)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) GetCreditAccounts2(marketFilter MarketFilter, caFilter CreditAccountFilter, offset *big.Int, limit *big.Int) (struct {
	Data       []CreditAccountData
	NextOffset *big.Int
}, error) {
	return _CreditAccountCompressor.Contract.GetCreditAccounts2(&_CreditAccountCompressor.CallOpts, marketFilter, caFilter, offset, limit)
}

// MarketConfiguratorFactory is a free data retrieval call binding the contract method 0x3ecc6e7d.
//
// Solidity: function marketConfiguratorFactory() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) MarketConfiguratorFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "marketConfiguratorFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketConfiguratorFactory is a free data retrieval call binding the contract method 0x3ecc6e7d.
//
// Solidity: function marketConfiguratorFactory() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorSession) MarketConfiguratorFactory() (common.Address, error) {
	return _CreditAccountCompressor.Contract.MarketConfiguratorFactory(&_CreditAccountCompressor.CallOpts)
}

// MarketConfiguratorFactory is a free data retrieval call binding the contract method 0x3ecc6e7d.
//
// Solidity: function marketConfiguratorFactory() view returns(address)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) MarketConfiguratorFactory() (common.Address, error) {
	return _CreditAccountCompressor.Contract.MarketConfiguratorFactory(&_CreditAccountCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditAccountCompressor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorSession) Version() (*big.Int, error) {
	return _CreditAccountCompressor.Contract.Version(&_CreditAccountCompressor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditAccountCompressor *CreditAccountCompressorCallerSession) Version() (*big.Int, error) {
	return _CreditAccountCompressor.Contract.Version(&_CreditAccountCompressor.CallOpts)
}
