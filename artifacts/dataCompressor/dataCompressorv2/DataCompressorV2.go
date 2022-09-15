// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataCompressorv2

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

// ContractAdapter is an auto generated low-level Go binding around an user-defined struct.
type ContractAdapter struct {
	AllowedContract common.Address
	Adapter         common.Address
}

// CreditAccountData is an auto generated low-level Go binding around an user-defined struct.
type CreditAccountData struct {
	Addr                              common.Address
	Borrower                          common.Address
	InUse                             bool
	CreditManager                     common.Address
	Underlying                        common.Address
	BorrowedAmountPlusInterest        *big.Int
	BorrowedAmountPlusInterestAndFees *big.Int
	TotalValue                        *big.Int
	HealthFactor                      *big.Int
	BorrowRate                        *big.Int
	Balances                          []TokenBalance
	RepayAmount                       *big.Int
	LiquidationAmount                 *big.Int
	CanBeClosed                       bool
	BorrowedAmount                    *big.Int
	CumulativeIndexAtOpen             *big.Int
	Since                             *big.Int
	Version                           uint8
	EnabledTokenMask                  *big.Int
}

// CreditManagerData is an auto generated low-level Go binding around an user-defined struct.
type CreditManagerData struct {
	Addr                       common.Address
	Underlying                 common.Address
	Pool                       common.Address
	IsWETH                     bool
	CanBorrow                  bool
	BorrowRate                 *big.Int
	MinAmount                  *big.Int
	MaxAmount                  *big.Int
	MaxLeverageFactor          *big.Int
	AvailableLiquidity         *big.Int
	CollateralTokens           []common.Address
	Adapters                   []ContractAdapter
	LiquidationThresholds      []*big.Int
	Version                    uint8
	CreditFacade               common.Address
	CreditConfigurator         common.Address
	IsDegenMode                bool
	DegenNFT                   common.Address
	IsIncreaseDebtForbidden    bool
	ForbiddenTokenMask         *big.Int
	MaxEnabledTokensLength     uint8
	FeeInterest                uint16
	FeeLiquidation             uint16
	LiquidationDiscount        uint16
	FeeLiquidationExpired      uint16
	LiquidationDiscountExpired uint16
}

// PoolData is an auto generated low-level Go binding around an user-defined struct.
type PoolData struct {
	Addr                   common.Address
	IsWETH                 bool
	Underlying             common.Address
	DieselToken            common.Address
	LinearCumulativeIndex  *big.Int
	AvailableLiquidity     *big.Int
	ExpectedLiquidity      *big.Int
	ExpectedLiquidityLimit *big.Int
	TotalBorrowed          *big.Int
	DepositAPYRAY          *big.Int
	BorrowAPYRAY           *big.Int
	DieselRateRAY          *big.Int
	WithdrawFee            *big.Int
	CumulativeIndexRAY     *big.Int
	TimestampLU            *big.Int
	Version                uint8
}

// TokenBalance is an auto generated low-level Go binding around an user-defined struct.
type TokenBalance struct {
	Token     common.Address
	Balance   *big.Int
	IsAllowed bool
	IsEnabled bool
}

// DataCompressorv2MetaData contains all meta data concerning the DataCompressorv2 contract.
var DataCompressorv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotCreditManagerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPoolException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractsRegister\",\"outputs\":[{\"internalType\":\"contractContractsRegister\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_allowedContract\",\"type\":\"address\"}],\"name\":\"getAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inUse\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterestAndFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokenMask\",\"type\":\"uint256\"}],\"internalType\":\"structCreditAccountData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getCreditAccountList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"inUse\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmountPlusInterestAndFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canBeClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndexAtOpen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"enabledTokenMask\",\"type\":\"uint256\"}],\"internalType\":\"structCreditAccountData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"}],\"name\":\"getCreditManagerData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"collateralTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDegenMode\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isIncreaseDebtForbidden\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreditManagersList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canBorrow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLeverageFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"collateralTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"allowedContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"internalType\":\"structContractAdapter[]\",\"name\":\"adapters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidationThresholds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creditFacade\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creditConfigurator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDegenMode\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"degenNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isIncreaseDebtForbidden\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"forbiddenTokenMask\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxEnabledTokensLength\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeInterest\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidation\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscount\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeLiquidationExpired\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationDiscountExpired\",\"type\":\"uint16\"}],\"internalType\":\"structCreditManagerData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"linearCumulativeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidityLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampLU\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"internalType\":\"structPoolData\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolsList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWETH\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dieselToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"linearCumulativeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLiquidityLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dieselRate_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeIndex_RAY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestampLU\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"internalType\":\"structPoolData[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"hasOpenedCreditAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DataCompressorv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use DataCompressorv2MetaData.ABI instead.
var DataCompressorv2ABI = DataCompressorv2MetaData.ABI

// DataCompressorv2 is an auto generated Go binding around an Ethereum contract.
type DataCompressorv2 struct {
	DataCompressorv2Caller     // Read-only binding to the contract
	DataCompressorv2Transactor // Write-only binding to the contract
	DataCompressorv2Filterer   // Log filterer for contract events
}

// DataCompressorv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type DataCompressorv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DataCompressorv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataCompressorv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataCompressorv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataCompressorv2Session struct {
	Contract     *DataCompressorv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCompressorv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCompressorv2CallerSession struct {
	Contract *DataCompressorv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DataCompressorv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataCompressorv2TransactorSession struct {
	Contract     *DataCompressorv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DataCompressorv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type DataCompressorv2Raw struct {
	Contract *DataCompressorv2 // Generic contract binding to access the raw methods on
}

// DataCompressorv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCompressorv2CallerRaw struct {
	Contract *DataCompressorv2Caller // Generic read-only contract binding to access the raw methods on
}

// DataCompressorv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataCompressorv2TransactorRaw struct {
	Contract *DataCompressorv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDataCompressorv2 creates a new instance of DataCompressorv2, bound to a specific deployed contract.
func NewDataCompressorv2(address common.Address, backend bind.ContractBackend) (*DataCompressorv2, error) {
	contract, err := bindDataCompressorv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv2{DataCompressorv2Caller: DataCompressorv2Caller{contract: contract}, DataCompressorv2Transactor: DataCompressorv2Transactor{contract: contract}, DataCompressorv2Filterer: DataCompressorv2Filterer{contract: contract}}, nil
}

// NewDataCompressorv2Caller creates a new read-only instance of DataCompressorv2, bound to a specific deployed contract.
func NewDataCompressorv2Caller(address common.Address, caller bind.ContractCaller) (*DataCompressorv2Caller, error) {
	contract, err := bindDataCompressorv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv2Caller{contract: contract}, nil
}

// NewDataCompressorv2Transactor creates a new write-only instance of DataCompressorv2, bound to a specific deployed contract.
func NewDataCompressorv2Transactor(address common.Address, transactor bind.ContractTransactor) (*DataCompressorv2Transactor, error) {
	contract, err := bindDataCompressorv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv2Transactor{contract: contract}, nil
}

// NewDataCompressorv2Filterer creates a new log filterer instance of DataCompressorv2, bound to a specific deployed contract.
func NewDataCompressorv2Filterer(address common.Address, filterer bind.ContractFilterer) (*DataCompressorv2Filterer, error) {
	contract, err := bindDataCompressorv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataCompressorv2Filterer{contract: contract}, nil
}

// bindDataCompressorv2 binds a generic wrapper to an already deployed contract.
func bindDataCompressorv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataCompressorv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressorv2 *DataCompressorv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressorv2.Contract.DataCompressorv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressorv2 *DataCompressorv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressorv2.Contract.DataCompressorv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressorv2 *DataCompressorv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressorv2.Contract.DataCompressorv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataCompressorv2 *DataCompressorv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataCompressorv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataCompressorv2 *DataCompressorv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataCompressorv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataCompressorv2 *DataCompressorv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataCompressorv2.Contract.contract.Transact(opts, method, params...)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Caller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Session) WETHToken() (common.Address, error) {
	return _DataCompressorv2.Contract.WETHToken(&_DataCompressorv2.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_DataCompressorv2 *DataCompressorv2CallerSession) WETHToken() (common.Address, error) {
	return _DataCompressorv2.Contract.WETHToken(&_DataCompressorv2.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Caller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Session) AddressProvider() (common.Address, error) {
	return _DataCompressorv2.Contract.AddressProvider(&_DataCompressorv2.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_DataCompressorv2 *DataCompressorv2CallerSession) AddressProvider() (common.Address, error) {
	return _DataCompressorv2.Contract.AddressProvider(&_DataCompressorv2.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Caller) ContractsRegister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "contractsRegister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv2 *DataCompressorv2Session) ContractsRegister() (common.Address, error) {
	return _DataCompressorv2.Contract.ContractsRegister(&_DataCompressorv2.CallOpts)
}

// ContractsRegister is a free data retrieval call binding the contract method 0x7a0c7b21.
//
// Solidity: function contractsRegister() view returns(address)
func (_DataCompressorv2 *DataCompressorv2CallerSession) ContractsRegister() (common.Address, error) {
	return _DataCompressorv2.Contract.ContractsRegister(&_DataCompressorv2.CallOpts)
}

// GetAdapter is a free data retrieval call binding the contract method 0x4c472fc9.
//
// Solidity: function getAdapter(address _creditManager, address _allowedContract) view returns(address adapter)
func (_DataCompressorv2 *DataCompressorv2Caller) GetAdapter(opts *bind.CallOpts, _creditManager common.Address, _allowedContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getAdapter", _creditManager, _allowedContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdapter is a free data retrieval call binding the contract method 0x4c472fc9.
//
// Solidity: function getAdapter(address _creditManager, address _allowedContract) view returns(address adapter)
func (_DataCompressorv2 *DataCompressorv2Session) GetAdapter(_creditManager common.Address, _allowedContract common.Address) (common.Address, error) {
	return _DataCompressorv2.Contract.GetAdapter(&_DataCompressorv2.CallOpts, _creditManager, _allowedContract)
}

// GetAdapter is a free data retrieval call binding the contract method 0x4c472fc9.
//
// Solidity: function getAdapter(address _creditManager, address _allowedContract) view returns(address adapter)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetAdapter(_creditManager common.Address, _allowedContract common.Address) (common.Address, error) {
	return _DataCompressorv2.Contract.GetAdapter(&_DataCompressorv2.CallOpts, _creditManager, _allowedContract)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256) result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetCreditAccountData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (CreditAccountData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getCreditAccountData", _creditManager, borrower)

	if err != nil {
		return *new(CreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditAccountData)).(*CreditAccountData)

	return out0, err

}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256) result)
func (_DataCompressorv2 *DataCompressorv2Session) GetCreditAccountData(_creditManager common.Address, borrower common.Address) (CreditAccountData, error) {
	return _DataCompressorv2.Contract.GetCreditAccountData(&_DataCompressorv2.CallOpts, _creditManager, borrower)
}

// GetCreditAccountData is a free data retrieval call binding the contract method 0x0dbd616d.
//
// Solidity: function getCreditAccountData(address _creditManager, address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256) result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetCreditAccountData(_creditManager common.Address, borrower common.Address) (CreditAccountData, error) {
	return _DataCompressorv2.Contract.GetCreditAccountData(&_DataCompressorv2.CallOpts, _creditManager, borrower)
}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256)[] result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetCreditAccountList(opts *bind.CallOpts, borrower common.Address) ([]CreditAccountData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getCreditAccountList", borrower)

	if err != nil {
		return *new([]CreditAccountData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditAccountData)).(*[]CreditAccountData)

	return out0, err

}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256)[] result)
func (_DataCompressorv2 *DataCompressorv2Session) GetCreditAccountList(borrower common.Address) ([]CreditAccountData, error) {
	return _DataCompressorv2.Contract.GetCreditAccountList(&_DataCompressorv2.CallOpts, borrower)
}

// GetCreditAccountList is a free data retrieval call binding the contract method 0xa80deda3.
//
// Solidity: function getCreditAccountList(address borrower) view returns((address,address,bool,address,address,uint256,uint256,uint256,uint256,uint256,(address,uint256,bool,bool)[],uint256,uint256,bool,uint256,uint256,uint256,uint8,uint256)[] result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetCreditAccountList(borrower common.Address) ([]CreditAccountData, error) {
	return _DataCompressorv2.Contract.GetCreditAccountList(&_DataCompressorv2.CallOpts, borrower)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _creditManager) view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16) result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address) (CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getCreditManagerData", _creditManager)

	if err != nil {
		return *new(CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new(CreditManagerData)).(*CreditManagerData)

	return out0, err

}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _creditManager) view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16) result)
func (_DataCompressorv2 *DataCompressorv2Session) GetCreditManagerData(_creditManager common.Address) (CreditManagerData, error) {
	return _DataCompressorv2.Contract.GetCreditManagerData(&_DataCompressorv2.CallOpts, _creditManager)
}

// GetCreditManagerData is a free data retrieval call binding the contract method 0xae093f3f.
//
// Solidity: function getCreditManagerData(address _creditManager) view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16) result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetCreditManagerData(_creditManager common.Address) (CreditManagerData, error) {
	return _DataCompressorv2.Contract.GetCreditManagerData(&_DataCompressorv2.CallOpts, _creditManager)
}

// GetCreditManagersList is a free data retrieval call binding the contract method 0x663b8fdb.
//
// Solidity: function getCreditManagersList() view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16)[] result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetCreditManagersList(opts *bind.CallOpts) ([]CreditManagerData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getCreditManagersList")

	if err != nil {
		return *new([]CreditManagerData), err
	}

	out0 := *abi.ConvertType(out[0], new([]CreditManagerData)).(*[]CreditManagerData)

	return out0, err

}

// GetCreditManagersList is a free data retrieval call binding the contract method 0x663b8fdb.
//
// Solidity: function getCreditManagersList() view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16)[] result)
func (_DataCompressorv2 *DataCompressorv2Session) GetCreditManagersList() ([]CreditManagerData, error) {
	return _DataCompressorv2.Contract.GetCreditManagersList(&_DataCompressorv2.CallOpts)
}

// GetCreditManagersList is a free data retrieval call binding the contract method 0x663b8fdb.
//
// Solidity: function getCreditManagersList() view returns((address,address,address,bool,bool,uint256,uint256,uint256,uint256,uint256,address[],(address,address)[],uint256[],uint8,address,address,bool,address,bool,uint256,uint8,uint16,uint16,uint16,uint16,uint16)[] result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetCreditManagersList() ([]CreditManagerData, error) {
	return _DataCompressorv2.Contract.GetCreditManagersList(&_DataCompressorv2.CallOpts)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8) result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetPoolData(opts *bind.CallOpts, _pool common.Address) (PoolData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getPoolData", _pool)

	if err != nil {
		return *new(PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolData)).(*PoolData)

	return out0, err

}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8) result)
func (_DataCompressorv2 *DataCompressorv2Session) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressorv2.Contract.GetPoolData(&_DataCompressorv2.CallOpts, _pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address _pool) view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8) result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetPoolData(_pool common.Address) (PoolData, error) {
	return _DataCompressorv2.Contract.GetPoolData(&_DataCompressorv2.CallOpts, _pool)
}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8)[] result)
func (_DataCompressorv2 *DataCompressorv2Caller) GetPoolsList(opts *bind.CallOpts) ([]PoolData, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "getPoolsList")

	if err != nil {
		return *new([]PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolData)).(*[]PoolData)

	return out0, err

}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8)[] result)
func (_DataCompressorv2 *DataCompressorv2Session) GetPoolsList() ([]PoolData, error) {
	return _DataCompressorv2.Contract.GetPoolsList(&_DataCompressorv2.CallOpts)
}

// GetPoolsList is a free data retrieval call binding the contract method 0x1bcd8fc0.
//
// Solidity: function getPoolsList() view returns((address,bool,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint8)[] result)
func (_DataCompressorv2 *DataCompressorv2CallerSession) GetPoolsList() ([]PoolData, error) {
	return _DataCompressorv2.Contract.GetPoolsList(&_DataCompressorv2.CallOpts)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address _creditManager, address borrower) view returns(bool)
func (_DataCompressorv2 *DataCompressorv2Caller) HasOpenedCreditAccount(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "hasOpenedCreditAccount", _creditManager, borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address _creditManager, address borrower) view returns(bool)
func (_DataCompressorv2 *DataCompressorv2Session) HasOpenedCreditAccount(_creditManager common.Address, borrower common.Address) (bool, error) {
	return _DataCompressorv2.Contract.HasOpenedCreditAccount(&_DataCompressorv2.CallOpts, _creditManager, borrower)
}

// HasOpenedCreditAccount is a free data retrieval call binding the contract method 0xfc9914cb.
//
// Solidity: function hasOpenedCreditAccount(address _creditManager, address borrower) view returns(bool)
func (_DataCompressorv2 *DataCompressorv2CallerSession) HasOpenedCreditAccount(_creditManager common.Address, borrower common.Address) (bool, error) {
	return _DataCompressorv2.Contract.HasOpenedCreditAccount(&_DataCompressorv2.CallOpts, _creditManager, borrower)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv2 *DataCompressorv2Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DataCompressorv2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv2 *DataCompressorv2Session) Version() (*big.Int, error) {
	return _DataCompressorv2.Contract.Version(&_DataCompressorv2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_DataCompressorv2 *DataCompressorv2CallerSession) Version() (*big.Int, error) {
	return _DataCompressorv2.Contract.Version(&_DataCompressorv2.CallOpts)
}
