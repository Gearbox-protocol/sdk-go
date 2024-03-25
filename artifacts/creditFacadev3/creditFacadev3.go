// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creditFacadev3

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

// MultiCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCall struct {
	Target   common.Address
	CallData []byte
}

// CreditFacadev3MetaData contains all meta data concerning the CreditFacadev3 contract.
var CreditFacadev3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_degenNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_expirable\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceLessThanExpectedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowAmountOutOfLimitsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowedBlockLimitException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotConfiguratorException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotCreditAccountOwnerException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotPausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotUnpausableAdminException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CloseAccountWithEnabledTokensException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreditAccountNotLiquidatableException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CustomHealthFactorTooLowException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedBalancesAlreadySetException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedBalancesNotSetException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForbiddenInWhitelistedModeException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForbiddenTokenBalanceIncreasedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForbiddenTokenEnabledException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForbiddenTokensException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectParameterException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateralHintException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permission\",\"type\":\"uint256\"}],\"name\":\"NoPermissionException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAfterExpirationException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedWhenNotExpirableException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedBotException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceFeedDoesNotExistException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingTokenBalanceIncreasedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetContractNotAllowedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedPermissionsException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownMethodException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"CloseCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"}],\"name\":\"Execute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"FinishMultiCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingFunds\",\"type\":\"uint256\"}],\"name\":\"LiquidateCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"NewController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"OpenCreditAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"StartMultiCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"botList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"botMulticall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canLiquidateWhilePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"closeCreditAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debtLimits\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"minDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxDebt\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"degenNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expirationDate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forbiddenTokenMask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"liquidateCreditAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lossParams\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"currentCumulativeLoss\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxCumulativeLoss\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDebtPerBlockMultiplier\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxQuotaMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"referralCode\",\"type\":\"uint256\"}],\"name\":\"openCreditAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBotList\",\"type\":\"address\"}],\"name\":\"setBotList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creditAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"permissions\",\"type\":\"uint192\"}],\"name\":\"setBotPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newController\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newMaxCumulativeLoss\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"resetCumulativeLoss\",\"type\":\"bool\"}],\"name\":\"setCumulativeLossParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newMinDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"newMaxDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"newMaxDebtPerBlockMultiplier\",\"type\":\"uint8\"}],\"name\":\"setDebtLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"enumAllowanceAction\",\"name\":\"allowance\",\"type\":\"uint8\"}],\"name\":\"setEmergencyLiquidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"newExpirationDate\",\"type\":\"uint40\"}],\"name\":\"setExpirationDate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumAllowanceAction\",\"name\":\"allowance\",\"type\":\"uint8\"}],\"name\":\"setTokenAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CreditFacadev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditFacadev3MetaData.ABI instead.
var CreditFacadev3ABI = CreditFacadev3MetaData.ABI

// CreditFacadev3 is an auto generated Go binding around an Ethereum contract.
type CreditFacadev3 struct {
	CreditFacadev3Caller     // Read-only binding to the contract
	CreditFacadev3Transactor // Write-only binding to the contract
	CreditFacadev3Filterer   // Log filterer for contract events
}

// CreditFacadev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type CreditFacadev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditFacadev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFacadev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFacadev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditFacadev3Session struct {
	Contract     *CreditFacadev3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditFacadev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditFacadev3CallerSession struct {
	Contract *CreditFacadev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CreditFacadev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditFacadev3TransactorSession struct {
	Contract     *CreditFacadev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CreditFacadev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type CreditFacadev3Raw struct {
	Contract *CreditFacadev3 // Generic contract binding to access the raw methods on
}

// CreditFacadev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditFacadev3CallerRaw struct {
	Contract *CreditFacadev3Caller // Generic read-only contract binding to access the raw methods on
}

// CreditFacadev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditFacadev3TransactorRaw struct {
	Contract *CreditFacadev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditFacadev3 creates a new instance of CreditFacadev3, bound to a specific deployed contract.
func NewCreditFacadev3(address common.Address, backend bind.ContractBackend) (*CreditFacadev3, error) {
	contract, err := bindCreditFacadev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3{CreditFacadev3Caller: CreditFacadev3Caller{contract: contract}, CreditFacadev3Transactor: CreditFacadev3Transactor{contract: contract}, CreditFacadev3Filterer: CreditFacadev3Filterer{contract: contract}}, nil
}

// NewCreditFacadev3Caller creates a new read-only instance of CreditFacadev3, bound to a specific deployed contract.
func NewCreditFacadev3Caller(address common.Address, caller bind.ContractCaller) (*CreditFacadev3Caller, error) {
	contract, err := bindCreditFacadev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3Caller{contract: contract}, nil
}

// NewCreditFacadev3Transactor creates a new write-only instance of CreditFacadev3, bound to a specific deployed contract.
func NewCreditFacadev3Transactor(address common.Address, transactor bind.ContractTransactor) (*CreditFacadev3Transactor, error) {
	contract, err := bindCreditFacadev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3Transactor{contract: contract}, nil
}

// NewCreditFacadev3Filterer creates a new log filterer instance of CreditFacadev3, bound to a specific deployed contract.
func NewCreditFacadev3Filterer(address common.Address, filterer bind.ContractFilterer) (*CreditFacadev3Filterer, error) {
	contract, err := bindCreditFacadev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3Filterer{contract: contract}, nil
}

// bindCreditFacadev3 binds a generic wrapper to an already deployed contract.
func bindCreditFacadev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditFacadev3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadev3 *CreditFacadev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadev3.Contract.CreditFacadev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadev3 *CreditFacadev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.CreditFacadev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadev3 *CreditFacadev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.CreditFacadev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditFacadev3 *CreditFacadev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditFacadev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditFacadev3 *CreditFacadev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditFacadev3 *CreditFacadev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.contract.Transact(opts, method, params...)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) Acl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "acl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) Acl() (common.Address, error) {
	return _CreditFacadev3.Contract.Acl(&_CreditFacadev3.CallOpts)
}

// Acl is a free data retrieval call binding the contract method 0xde287359.
//
// Solidity: function acl() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Acl() (common.Address, error) {
	return _CreditFacadev3.Contract.Acl(&_CreditFacadev3.CallOpts)
}

// BotList is a free data retrieval call binding the contract method 0xf6722f62.
//
// Solidity: function botList() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) BotList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "botList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BotList is a free data retrieval call binding the contract method 0xf6722f62.
//
// Solidity: function botList() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) BotList() (common.Address, error) {
	return _CreditFacadev3.Contract.BotList(&_CreditFacadev3.CallOpts)
}

// BotList is a free data retrieval call binding the contract method 0xf6722f62.
//
// Solidity: function botList() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) BotList() (common.Address, error) {
	return _CreditFacadev3.Contract.BotList(&_CreditFacadev3.CallOpts)
}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Caller) CanLiquidateWhilePaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "canLiquidateWhilePaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Session) CanLiquidateWhilePaused(arg0 common.Address) (bool, error) {
	return _CreditFacadev3.Contract.CanLiquidateWhilePaused(&_CreditFacadev3.CallOpts, arg0)
}

// CanLiquidateWhilePaused is a free data retrieval call binding the contract method 0x38975bc4.
//
// Solidity: function canLiquidateWhilePaused(address ) view returns(bool)
func (_CreditFacadev3 *CreditFacadev3CallerSession) CanLiquidateWhilePaused(arg0 common.Address) (bool, error) {
	return _CreditFacadev3.Contract.CanLiquidateWhilePaused(&_CreditFacadev3.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) Controller() (common.Address, error) {
	return _CreditFacadev3.Contract.Controller(&_CreditFacadev3.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Controller() (common.Address, error) {
	return _CreditFacadev3.Contract.Controller(&_CreditFacadev3.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) CreditManager() (common.Address, error) {
	return _CreditFacadev3.Contract.CreditManager(&_CreditFacadev3.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) CreditManager() (common.Address, error) {
	return _CreditFacadev3.Contract.CreditManager(&_CreditFacadev3.CallOpts)
}

// DebtLimits is a free data retrieval call binding the contract method 0x166bf9d9.
//
// Solidity: function debtLimits() view returns(uint128 minDebt, uint128 maxDebt)
func (_CreditFacadev3 *CreditFacadev3Caller) DebtLimits(opts *bind.CallOpts) (struct {
	MinDebt *big.Int
	MaxDebt *big.Int
}, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "debtLimits")

	outstruct := new(struct {
		MinDebt *big.Int
		MaxDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinDebt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DebtLimits is a free data retrieval call binding the contract method 0x166bf9d9.
//
// Solidity: function debtLimits() view returns(uint128 minDebt, uint128 maxDebt)
func (_CreditFacadev3 *CreditFacadev3Session) DebtLimits() (struct {
	MinDebt *big.Int
	MaxDebt *big.Int
}, error) {
	return _CreditFacadev3.Contract.DebtLimits(&_CreditFacadev3.CallOpts)
}

// DebtLimits is a free data retrieval call binding the contract method 0x166bf9d9.
//
// Solidity: function debtLimits() view returns(uint128 minDebt, uint128 maxDebt)
func (_CreditFacadev3 *CreditFacadev3CallerSession) DebtLimits() (struct {
	MinDebt *big.Int
	MaxDebt *big.Int
}, error) {
	return _CreditFacadev3.Contract.DebtLimits(&_CreditFacadev3.CallOpts)
}

// DegenNFT is a free data retrieval call binding the contract method 0x9408b63f.
//
// Solidity: function degenNFT() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) DegenNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "degenNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DegenNFT is a free data retrieval call binding the contract method 0x9408b63f.
//
// Solidity: function degenNFT() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) DegenNFT() (common.Address, error) {
	return _CreditFacadev3.Contract.DegenNFT(&_CreditFacadev3.CallOpts)
}

// DegenNFT is a free data retrieval call binding the contract method 0x9408b63f.
//
// Solidity: function degenNFT() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) DegenNFT() (common.Address, error) {
	return _CreditFacadev3.Contract.DegenNFT(&_CreditFacadev3.CallOpts)
}

// Expirable is a free data retrieval call binding the contract method 0xb1836d32.
//
// Solidity: function expirable() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Caller) Expirable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "expirable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Expirable is a free data retrieval call binding the contract method 0xb1836d32.
//
// Solidity: function expirable() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Session) Expirable() (bool, error) {
	return _CreditFacadev3.Contract.Expirable(&_CreditFacadev3.CallOpts)
}

// Expirable is a free data retrieval call binding the contract method 0xb1836d32.
//
// Solidity: function expirable() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Expirable() (bool, error) {
	return _CreditFacadev3.Contract.Expirable(&_CreditFacadev3.CallOpts)
}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(uint40)
func (_CreditFacadev3 *CreditFacadev3Caller) ExpirationDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "expirationDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(uint40)
func (_CreditFacadev3 *CreditFacadev3Session) ExpirationDate() (*big.Int, error) {
	return _CreditFacadev3.Contract.ExpirationDate(&_CreditFacadev3.CallOpts)
}

// ExpirationDate is a free data retrieval call binding the contract method 0x8f620487.
//
// Solidity: function expirationDate() view returns(uint40)
func (_CreditFacadev3 *CreditFacadev3CallerSession) ExpirationDate() (*big.Int, error) {
	return _CreditFacadev3.Contract.ExpirationDate(&_CreditFacadev3.CallOpts)
}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Caller) ForbiddenTokenMask(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "forbiddenTokenMask")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Session) ForbiddenTokenMask() (*big.Int, error) {
	return _CreditFacadev3.Contract.ForbiddenTokenMask(&_CreditFacadev3.CallOpts)
}

// ForbiddenTokenMask is a free data retrieval call binding the contract method 0x9fd12b77.
//
// Solidity: function forbiddenTokenMask() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3CallerSession) ForbiddenTokenMask() (*big.Int, error) {
	return _CreditFacadev3.Contract.ForbiddenTokenMask(&_CreditFacadev3.CallOpts)
}

// LossParams is a free data retrieval call binding the contract method 0x50393109.
//
// Solidity: function lossParams() view returns(uint128 currentCumulativeLoss, uint128 maxCumulativeLoss)
func (_CreditFacadev3 *CreditFacadev3Caller) LossParams(opts *bind.CallOpts) (struct {
	CurrentCumulativeLoss *big.Int
	MaxCumulativeLoss     *big.Int
}, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "lossParams")

	outstruct := new(struct {
		CurrentCumulativeLoss *big.Int
		MaxCumulativeLoss     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentCumulativeLoss = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxCumulativeLoss = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LossParams is a free data retrieval call binding the contract method 0x50393109.
//
// Solidity: function lossParams() view returns(uint128 currentCumulativeLoss, uint128 maxCumulativeLoss)
func (_CreditFacadev3 *CreditFacadev3Session) LossParams() (struct {
	CurrentCumulativeLoss *big.Int
	MaxCumulativeLoss     *big.Int
}, error) {
	return _CreditFacadev3.Contract.LossParams(&_CreditFacadev3.CallOpts)
}

// LossParams is a free data retrieval call binding the contract method 0x50393109.
//
// Solidity: function lossParams() view returns(uint128 currentCumulativeLoss, uint128 maxCumulativeLoss)
func (_CreditFacadev3 *CreditFacadev3CallerSession) LossParams() (struct {
	CurrentCumulativeLoss *big.Int
	MaxCumulativeLoss     *big.Int
}, error) {
	return _CreditFacadev3.Contract.LossParams(&_CreditFacadev3.CallOpts)
}

// MaxDebtPerBlockMultiplier is a free data retrieval call binding the contract method 0x478ade36.
//
// Solidity: function maxDebtPerBlockMultiplier() view returns(uint8)
func (_CreditFacadev3 *CreditFacadev3Caller) MaxDebtPerBlockMultiplier(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "maxDebtPerBlockMultiplier")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MaxDebtPerBlockMultiplier is a free data retrieval call binding the contract method 0x478ade36.
//
// Solidity: function maxDebtPerBlockMultiplier() view returns(uint8)
func (_CreditFacadev3 *CreditFacadev3Session) MaxDebtPerBlockMultiplier() (uint8, error) {
	return _CreditFacadev3.Contract.MaxDebtPerBlockMultiplier(&_CreditFacadev3.CallOpts)
}

// MaxDebtPerBlockMultiplier is a free data retrieval call binding the contract method 0x478ade36.
//
// Solidity: function maxDebtPerBlockMultiplier() view returns(uint8)
func (_CreditFacadev3 *CreditFacadev3CallerSession) MaxDebtPerBlockMultiplier() (uint8, error) {
	return _CreditFacadev3.Contract.MaxDebtPerBlockMultiplier(&_CreditFacadev3.CallOpts)
}

// MaxQuotaMultiplier is a free data retrieval call binding the contract method 0x874b2e02.
//
// Solidity: function maxQuotaMultiplier() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Caller) MaxQuotaMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "maxQuotaMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxQuotaMultiplier is a free data retrieval call binding the contract method 0x874b2e02.
//
// Solidity: function maxQuotaMultiplier() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Session) MaxQuotaMultiplier() (*big.Int, error) {
	return _CreditFacadev3.Contract.MaxQuotaMultiplier(&_CreditFacadev3.CallOpts)
}

// MaxQuotaMultiplier is a free data retrieval call binding the contract method 0x874b2e02.
//
// Solidity: function maxQuotaMultiplier() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3CallerSession) MaxQuotaMultiplier() (*big.Int, error) {
	return _CreditFacadev3.Contract.MaxQuotaMultiplier(&_CreditFacadev3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3Session) Paused() (bool, error) {
	return _CreditFacadev3.Contract.Paused(&_CreditFacadev3.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Paused() (bool, error) {
	return _CreditFacadev3.Contract.Paused(&_CreditFacadev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3Session) Version() (*big.Int, error) {
	return _CreditFacadev3.Contract.Version(&_CreditFacadev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Version() (*big.Int, error) {
	return _CreditFacadev3.Contract.Version(&_CreditFacadev3.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditFacadev3.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CreditFacadev3 *CreditFacadev3Session) Weth() (common.Address, error) {
	return _CreditFacadev3.Contract.Weth(&_CreditFacadev3.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_CreditFacadev3 *CreditFacadev3CallerSession) Weth() (common.Address, error) {
	return _CreditFacadev3.Contract.Weth(&_CreditFacadev3.CallOpts)
}

// BotMulticall is a paid mutator transaction binding the contract method 0x7e2ca9db.
//
// Solidity: function botMulticall(address creditAccount, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) BotMulticall(opts *bind.TransactOpts, creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "botMulticall", creditAccount, calls)
}

// BotMulticall is a paid mutator transaction binding the contract method 0x7e2ca9db.
//
// Solidity: function botMulticall(address creditAccount, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3Session) BotMulticall(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.BotMulticall(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// BotMulticall is a paid mutator transaction binding the contract method 0x7e2ca9db.
//
// Solidity: function botMulticall(address creditAccount, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) BotMulticall(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.BotMulticall(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x36b2ced3.
//
// Solidity: function closeCreditAccount(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) CloseCreditAccount(opts *bind.TransactOpts, creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "closeCreditAccount", creditAccount, calls)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x36b2ced3.
//
// Solidity: function closeCreditAccount(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3Session) CloseCreditAccount(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.CloseCreditAccount(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// CloseCreditAccount is a paid mutator transaction binding the contract method 0x36b2ced3.
//
// Solidity: function closeCreditAccount(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) CloseCreditAccount(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.CloseCreditAccount(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe3f46b26.
//
// Solidity: function liquidateCreditAccount(address creditAccount, address to, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) LiquidateCreditAccount(opts *bind.TransactOpts, creditAccount common.Address, to common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "liquidateCreditAccount", creditAccount, to, calls)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe3f46b26.
//
// Solidity: function liquidateCreditAccount(address creditAccount, address to, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3Session) LiquidateCreditAccount(creditAccount common.Address, to common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.LiquidateCreditAccount(&_CreditFacadev3.TransactOpts, creditAccount, to, calls)
}

// LiquidateCreditAccount is a paid mutator transaction binding the contract method 0xe3f46b26.
//
// Solidity: function liquidateCreditAccount(address creditAccount, address to, (address,bytes)[] calls) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) LiquidateCreditAccount(creditAccount common.Address, to common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.LiquidateCreditAccount(&_CreditFacadev3.TransactOpts, creditAccount, to, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xebe4107c.
//
// Solidity: function multicall(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) Multicall(opts *bind.TransactOpts, creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "multicall", creditAccount, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xebe4107c.
//
// Solidity: function multicall(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3Session) Multicall(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Multicall(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0xebe4107c.
//
// Solidity: function multicall(address creditAccount, (address,bytes)[] calls) payable returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) Multicall(creditAccount common.Address, calls []MultiCall) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Multicall(&_CreditFacadev3.TransactOpts, creditAccount, calls)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x92beab1d.
//
// Solidity: function openCreditAccount(address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns(address creditAccount)
func (_CreditFacadev3 *CreditFacadev3Transactor) OpenCreditAccount(opts *bind.TransactOpts, onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "openCreditAccount", onBehalfOf, calls, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x92beab1d.
//
// Solidity: function openCreditAccount(address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns(address creditAccount)
func (_CreditFacadev3 *CreditFacadev3Session) OpenCreditAccount(onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.OpenCreditAccount(&_CreditFacadev3.TransactOpts, onBehalfOf, calls, referralCode)
}

// OpenCreditAccount is a paid mutator transaction binding the contract method 0x92beab1d.
//
// Solidity: function openCreditAccount(address onBehalfOf, (address,bytes)[] calls, uint256 referralCode) payable returns(address creditAccount)
func (_CreditFacadev3 *CreditFacadev3TransactorSession) OpenCreditAccount(onBehalfOf common.Address, calls []MultiCall, referralCode *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.OpenCreditAccount(&_CreditFacadev3.TransactOpts, onBehalfOf, calls, referralCode)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFacadev3 *CreditFacadev3Session) Pause() (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Pause(&_CreditFacadev3.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) Pause() (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Pause(&_CreditFacadev3.TransactOpts)
}

// SetBotList is a paid mutator transaction binding the contract method 0x8ad1386e.
//
// Solidity: function setBotList(address newBotList) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetBotList(opts *bind.TransactOpts, newBotList common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setBotList", newBotList)
}

// SetBotList is a paid mutator transaction binding the contract method 0x8ad1386e.
//
// Solidity: function setBotList(address newBotList) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetBotList(newBotList common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetBotList(&_CreditFacadev3.TransactOpts, newBotList)
}

// SetBotList is a paid mutator transaction binding the contract method 0x8ad1386e.
//
// Solidity: function setBotList(address newBotList) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetBotList(newBotList common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetBotList(&_CreditFacadev3.TransactOpts, newBotList)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0xc5b73ed0.
//
// Solidity: function setBotPermissions(address creditAccount, address bot, uint192 permissions) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetBotPermissions(opts *bind.TransactOpts, creditAccount common.Address, bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setBotPermissions", creditAccount, bot, permissions)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0xc5b73ed0.
//
// Solidity: function setBotPermissions(address creditAccount, address bot, uint192 permissions) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetBotPermissions(creditAccount common.Address, bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetBotPermissions(&_CreditFacadev3.TransactOpts, creditAccount, bot, permissions)
}

// SetBotPermissions is a paid mutator transaction binding the contract method 0xc5b73ed0.
//
// Solidity: function setBotPermissions(address creditAccount, address bot, uint192 permissions) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetBotPermissions(creditAccount common.Address, bot common.Address, permissions *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetBotPermissions(&_CreditFacadev3.TransactOpts, creditAccount, bot, permissions)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetController(opts *bind.TransactOpts, newController common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setController", newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetController(&_CreditFacadev3.TransactOpts, newController)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address newController) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetController(newController common.Address) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetController(&_CreditFacadev3.TransactOpts, newController)
}

// SetCumulativeLossParams is a paid mutator transaction binding the contract method 0x0103dc6b.
//
// Solidity: function setCumulativeLossParams(uint128 newMaxCumulativeLoss, bool resetCumulativeLoss) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetCumulativeLossParams(opts *bind.TransactOpts, newMaxCumulativeLoss *big.Int, resetCumulativeLoss bool) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setCumulativeLossParams", newMaxCumulativeLoss, resetCumulativeLoss)
}

// SetCumulativeLossParams is a paid mutator transaction binding the contract method 0x0103dc6b.
//
// Solidity: function setCumulativeLossParams(uint128 newMaxCumulativeLoss, bool resetCumulativeLoss) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetCumulativeLossParams(newMaxCumulativeLoss *big.Int, resetCumulativeLoss bool) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetCumulativeLossParams(&_CreditFacadev3.TransactOpts, newMaxCumulativeLoss, resetCumulativeLoss)
}

// SetCumulativeLossParams is a paid mutator transaction binding the contract method 0x0103dc6b.
//
// Solidity: function setCumulativeLossParams(uint128 newMaxCumulativeLoss, bool resetCumulativeLoss) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetCumulativeLossParams(newMaxCumulativeLoss *big.Int, resetCumulativeLoss bool) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetCumulativeLossParams(&_CreditFacadev3.TransactOpts, newMaxCumulativeLoss, resetCumulativeLoss)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0x1656af9d.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt, uint8 newMaxDebtPerBlockMultiplier) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetDebtLimits(opts *bind.TransactOpts, newMinDebt *big.Int, newMaxDebt *big.Int, newMaxDebtPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setDebtLimits", newMinDebt, newMaxDebt, newMaxDebtPerBlockMultiplier)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0x1656af9d.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt, uint8 newMaxDebtPerBlockMultiplier) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetDebtLimits(newMinDebt *big.Int, newMaxDebt *big.Int, newMaxDebtPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetDebtLimits(&_CreditFacadev3.TransactOpts, newMinDebt, newMaxDebt, newMaxDebtPerBlockMultiplier)
}

// SetDebtLimits is a paid mutator transaction binding the contract method 0x1656af9d.
//
// Solidity: function setDebtLimits(uint128 newMinDebt, uint128 newMaxDebt, uint8 newMaxDebtPerBlockMultiplier) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetDebtLimits(newMinDebt *big.Int, newMaxDebt *big.Int, newMaxDebtPerBlockMultiplier uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetDebtLimits(&_CreditFacadev3.TransactOpts, newMinDebt, newMaxDebt, newMaxDebtPerBlockMultiplier)
}

// SetEmergencyLiquidator is a paid mutator transaction binding the contract method 0xc5d7ca39.
//
// Solidity: function setEmergencyLiquidator(address liquidator, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetEmergencyLiquidator(opts *bind.TransactOpts, liquidator common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setEmergencyLiquidator", liquidator, allowance)
}

// SetEmergencyLiquidator is a paid mutator transaction binding the contract method 0xc5d7ca39.
//
// Solidity: function setEmergencyLiquidator(address liquidator, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetEmergencyLiquidator(liquidator common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetEmergencyLiquidator(&_CreditFacadev3.TransactOpts, liquidator, allowance)
}

// SetEmergencyLiquidator is a paid mutator transaction binding the contract method 0xc5d7ca39.
//
// Solidity: function setEmergencyLiquidator(address liquidator, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetEmergencyLiquidator(liquidator common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetEmergencyLiquidator(&_CreditFacadev3.TransactOpts, liquidator, allowance)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetExpirationDate(opts *bind.TransactOpts, newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setExpirationDate", newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetExpirationDate(&_CreditFacadev3.TransactOpts, newExpirationDate)
}

// SetExpirationDate is a paid mutator transaction binding the contract method 0xeb9606df.
//
// Solidity: function setExpirationDate(uint40 newExpirationDate) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetExpirationDate(newExpirationDate *big.Int) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetExpirationDate(&_CreditFacadev3.TransactOpts, newExpirationDate)
}

// SetTokenAllowance is a paid mutator transaction binding the contract method 0x26f30692.
//
// Solidity: function setTokenAllowance(address token, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) SetTokenAllowance(opts *bind.TransactOpts, token common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "setTokenAllowance", token, allowance)
}

// SetTokenAllowance is a paid mutator transaction binding the contract method 0x26f30692.
//
// Solidity: function setTokenAllowance(address token, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3Session) SetTokenAllowance(token common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetTokenAllowance(&_CreditFacadev3.TransactOpts, token, allowance)
}

// SetTokenAllowance is a paid mutator transaction binding the contract method 0x26f30692.
//
// Solidity: function setTokenAllowance(address token, uint8 allowance) returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) SetTokenAllowance(token common.Address, allowance uint8) (*types.Transaction, error) {
	return _CreditFacadev3.Contract.SetTokenAllowance(&_CreditFacadev3.TransactOpts, token, allowance)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFacadev3 *CreditFacadev3Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditFacadev3.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFacadev3 *CreditFacadev3Session) Unpause() (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Unpause(&_CreditFacadev3.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditFacadev3 *CreditFacadev3TransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditFacadev3.Contract.Unpause(&_CreditFacadev3.TransactOpts)
}

// CreditFacadev3AddCollateralIterator is returned from FilterAddCollateral and is used to iterate over the raw logs and unpacked data for AddCollateral events raised by the CreditFacadev3 contract.
type CreditFacadev3AddCollateralIterator struct {
	Event *CreditFacadev3AddCollateral // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3AddCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3AddCollateral)
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
		it.Event = new(CreditFacadev3AddCollateral)
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
func (it *CreditFacadev3AddCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3AddCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3AddCollateral represents a AddCollateral event raised by the CreditFacadev3 contract.
type CreditFacadev3AddCollateral struct {
	CreditAccount common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddCollateral is a free log retrieval operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed creditAccount, address indexed token, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterAddCollateral(opts *bind.FilterOpts, creditAccount []common.Address, token []common.Address) (*CreditFacadev3AddCollateralIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "AddCollateral", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3AddCollateralIterator{contract: _CreditFacadev3.contract, event: "AddCollateral", logs: logs, sub: sub}, nil
}

// WatchAddCollateral is a free log subscription operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed creditAccount, address indexed token, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchAddCollateral(opts *bind.WatchOpts, sink chan<- *CreditFacadev3AddCollateral, creditAccount []common.Address, token []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "AddCollateral", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3AddCollateral)
				if err := _CreditFacadev3.contract.UnpackLog(event, "AddCollateral", log); err != nil {
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

// ParseAddCollateral is a log parse operation binding the contract event 0xa32435755c235de2976ed44a75a2f85cb01faf0c894f639fe0c32bb9455fea8f.
//
// Solidity: event AddCollateral(address indexed creditAccount, address indexed token, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseAddCollateral(log types.Log) (*CreditFacadev3AddCollateral, error) {
	event := new(CreditFacadev3AddCollateral)
	if err := _CreditFacadev3.contract.UnpackLog(event, "AddCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3CloseCreditAccountIterator is returned from FilterCloseCreditAccount and is used to iterate over the raw logs and unpacked data for CloseCreditAccount events raised by the CreditFacadev3 contract.
type CreditFacadev3CloseCreditAccountIterator struct {
	Event *CreditFacadev3CloseCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3CloseCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3CloseCreditAccount)
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
		it.Event = new(CreditFacadev3CloseCreditAccount)
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
func (it *CreditFacadev3CloseCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3CloseCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3CloseCreditAccount represents a CloseCreditAccount event raised by the CreditFacadev3 contract.
type CreditFacadev3CloseCreditAccount struct {
	CreditAccount common.Address
	Borrower      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCloseCreditAccount is a free log retrieval operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed creditAccount, address indexed borrower)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterCloseCreditAccount(opts *bind.FilterOpts, creditAccount []common.Address, borrower []common.Address) (*CreditFacadev3CloseCreditAccountIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "CloseCreditAccount", creditAccountRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3CloseCreditAccountIterator{contract: _CreditFacadev3.contract, event: "CloseCreditAccount", logs: logs, sub: sub}, nil
}

// WatchCloseCreditAccount is a free log subscription operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed creditAccount, address indexed borrower)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchCloseCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadev3CloseCreditAccount, creditAccount []common.Address, borrower []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "CloseCreditAccount", creditAccountRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3CloseCreditAccount)
				if err := _CreditFacadev3.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
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

// ParseCloseCreditAccount is a log parse operation binding the contract event 0x460ad03b1cf79b1d64d3aefa28475f110ab66e84649c52bb41ed796b9b391981.
//
// Solidity: event CloseCreditAccount(address indexed creditAccount, address indexed borrower)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseCloseCreditAccount(log types.Log) (*CreditFacadev3CloseCreditAccount, error) {
	event := new(CreditFacadev3CloseCreditAccount)
	if err := _CreditFacadev3.contract.UnpackLog(event, "CloseCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3DecreaseDebtIterator is returned from FilterDecreaseDebt and is used to iterate over the raw logs and unpacked data for DecreaseDebt events raised by the CreditFacadev3 contract.
type CreditFacadev3DecreaseDebtIterator struct {
	Event *CreditFacadev3DecreaseDebt // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3DecreaseDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3DecreaseDebt)
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
		it.Event = new(CreditFacadev3DecreaseDebt)
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
func (it *CreditFacadev3DecreaseDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3DecreaseDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3DecreaseDebt represents a DecreaseDebt event raised by the CreditFacadev3 contract.
type CreditFacadev3DecreaseDebt struct {
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDecreaseDebt is a free log retrieval operation binding the contract event 0x9ddbe9792bb03a06eab016daae23f04ec25454c24b836da2361fc703037f7762.
//
// Solidity: event DecreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterDecreaseDebt(opts *bind.FilterOpts, creditAccount []common.Address) (*CreditFacadev3DecreaseDebtIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "DecreaseDebt", creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3DecreaseDebtIterator{contract: _CreditFacadev3.contract, event: "DecreaseDebt", logs: logs, sub: sub}, nil
}

// WatchDecreaseDebt is a free log subscription operation binding the contract event 0x9ddbe9792bb03a06eab016daae23f04ec25454c24b836da2361fc703037f7762.
//
// Solidity: event DecreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchDecreaseDebt(opts *bind.WatchOpts, sink chan<- *CreditFacadev3DecreaseDebt, creditAccount []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "DecreaseDebt", creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3DecreaseDebt)
				if err := _CreditFacadev3.contract.UnpackLog(event, "DecreaseDebt", log); err != nil {
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

// ParseDecreaseDebt is a log parse operation binding the contract event 0x9ddbe9792bb03a06eab016daae23f04ec25454c24b836da2361fc703037f7762.
//
// Solidity: event DecreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseDecreaseDebt(log types.Log) (*CreditFacadev3DecreaseDebt, error) {
	event := new(CreditFacadev3DecreaseDebt)
	if err := _CreditFacadev3.contract.UnpackLog(event, "DecreaseDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3ExecuteIterator is returned from FilterExecute and is used to iterate over the raw logs and unpacked data for Execute events raised by the CreditFacadev3 contract.
type CreditFacadev3ExecuteIterator struct {
	Event *CreditFacadev3Execute // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3ExecuteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3Execute)
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
		it.Event = new(CreditFacadev3Execute)
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
func (it *CreditFacadev3ExecuteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3ExecuteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3Execute represents a Execute event raised by the CreditFacadev3 contract.
type CreditFacadev3Execute struct {
	CreditAccount  common.Address
	TargetContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecute is a free log retrieval operation binding the contract event 0x1b835de7d84f000a333cdc5822ae62eb63b38d4c622ef96ac50f27db56d7c768.
//
// Solidity: event Execute(address indexed creditAccount, address indexed targetContract)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterExecute(opts *bind.FilterOpts, creditAccount []common.Address, targetContract []common.Address) (*CreditFacadev3ExecuteIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "Execute", creditAccountRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3ExecuteIterator{contract: _CreditFacadev3.contract, event: "Execute", logs: logs, sub: sub}, nil
}

// WatchExecute is a free log subscription operation binding the contract event 0x1b835de7d84f000a333cdc5822ae62eb63b38d4c622ef96ac50f27db56d7c768.
//
// Solidity: event Execute(address indexed creditAccount, address indexed targetContract)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchExecute(opts *bind.WatchOpts, sink chan<- *CreditFacadev3Execute, creditAccount []common.Address, targetContract []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var targetContractRule []interface{}
	for _, targetContractItem := range targetContract {
		targetContractRule = append(targetContractRule, targetContractItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "Execute", creditAccountRule, targetContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3Execute)
				if err := _CreditFacadev3.contract.UnpackLog(event, "Execute", log); err != nil {
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

// ParseExecute is a log parse operation binding the contract event 0x1b835de7d84f000a333cdc5822ae62eb63b38d4c622ef96ac50f27db56d7c768.
//
// Solidity: event Execute(address indexed creditAccount, address indexed targetContract)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseExecute(log types.Log) (*CreditFacadev3Execute, error) {
	event := new(CreditFacadev3Execute)
	if err := _CreditFacadev3.contract.UnpackLog(event, "Execute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3FinishMultiCallIterator is returned from FilterFinishMultiCall and is used to iterate over the raw logs and unpacked data for FinishMultiCall events raised by the CreditFacadev3 contract.
type CreditFacadev3FinishMultiCallIterator struct {
	Event *CreditFacadev3FinishMultiCall // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3FinishMultiCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3FinishMultiCall)
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
		it.Event = new(CreditFacadev3FinishMultiCall)
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
func (it *CreditFacadev3FinishMultiCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3FinishMultiCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3FinishMultiCall represents a FinishMultiCall event raised by the CreditFacadev3 contract.
type CreditFacadev3FinishMultiCall struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFinishMultiCall is a free log retrieval operation binding the contract event 0x9fe19f2060e67aed557c7d1bc297d4bd2d8a8b952e3545c658ec4bc00be7d6c4.
//
// Solidity: event FinishMultiCall()
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterFinishMultiCall(opts *bind.FilterOpts) (*CreditFacadev3FinishMultiCallIterator, error) {

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "FinishMultiCall")
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3FinishMultiCallIterator{contract: _CreditFacadev3.contract, event: "FinishMultiCall", logs: logs, sub: sub}, nil
}

// WatchFinishMultiCall is a free log subscription operation binding the contract event 0x9fe19f2060e67aed557c7d1bc297d4bd2d8a8b952e3545c658ec4bc00be7d6c4.
//
// Solidity: event FinishMultiCall()
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchFinishMultiCall(opts *bind.WatchOpts, sink chan<- *CreditFacadev3FinishMultiCall) (event.Subscription, error) {

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "FinishMultiCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3FinishMultiCall)
				if err := _CreditFacadev3.contract.UnpackLog(event, "FinishMultiCall", log); err != nil {
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

// ParseFinishMultiCall is a log parse operation binding the contract event 0x9fe19f2060e67aed557c7d1bc297d4bd2d8a8b952e3545c658ec4bc00be7d6c4.
//
// Solidity: event FinishMultiCall()
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseFinishMultiCall(log types.Log) (*CreditFacadev3FinishMultiCall, error) {
	event := new(CreditFacadev3FinishMultiCall)
	if err := _CreditFacadev3.contract.UnpackLog(event, "FinishMultiCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3IncreaseDebtIterator is returned from FilterIncreaseDebt and is used to iterate over the raw logs and unpacked data for IncreaseDebt events raised by the CreditFacadev3 contract.
type CreditFacadev3IncreaseDebtIterator struct {
	Event *CreditFacadev3IncreaseDebt // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3IncreaseDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3IncreaseDebt)
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
		it.Event = new(CreditFacadev3IncreaseDebt)
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
func (it *CreditFacadev3IncreaseDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3IncreaseDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3IncreaseDebt represents a IncreaseDebt event raised by the CreditFacadev3 contract.
type CreditFacadev3IncreaseDebt struct {
	CreditAccount common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseDebt is a free log retrieval operation binding the contract event 0xcb6767c6e25552f1ab37156882f03b9ba603ebf3814463bf5eb9b0d2bd8a19b5.
//
// Solidity: event IncreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterIncreaseDebt(opts *bind.FilterOpts, creditAccount []common.Address) (*CreditFacadev3IncreaseDebtIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "IncreaseDebt", creditAccountRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3IncreaseDebtIterator{contract: _CreditFacadev3.contract, event: "IncreaseDebt", logs: logs, sub: sub}, nil
}

// WatchIncreaseDebt is a free log subscription operation binding the contract event 0xcb6767c6e25552f1ab37156882f03b9ba603ebf3814463bf5eb9b0d2bd8a19b5.
//
// Solidity: event IncreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchIncreaseDebt(opts *bind.WatchOpts, sink chan<- *CreditFacadev3IncreaseDebt, creditAccount []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "IncreaseDebt", creditAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3IncreaseDebt)
				if err := _CreditFacadev3.contract.UnpackLog(event, "IncreaseDebt", log); err != nil {
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

// ParseIncreaseDebt is a log parse operation binding the contract event 0xcb6767c6e25552f1ab37156882f03b9ba603ebf3814463bf5eb9b0d2bd8a19b5.
//
// Solidity: event IncreaseDebt(address indexed creditAccount, uint256 amount)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseIncreaseDebt(log types.Log) (*CreditFacadev3IncreaseDebt, error) {
	event := new(CreditFacadev3IncreaseDebt)
	if err := _CreditFacadev3.contract.UnpackLog(event, "IncreaseDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3LiquidateCreditAccountIterator is returned from FilterLiquidateCreditAccount and is used to iterate over the raw logs and unpacked data for LiquidateCreditAccount events raised by the CreditFacadev3 contract.
type CreditFacadev3LiquidateCreditAccountIterator struct {
	Event *CreditFacadev3LiquidateCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3LiquidateCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3LiquidateCreditAccount)
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
		it.Event = new(CreditFacadev3LiquidateCreditAccount)
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
func (it *CreditFacadev3LiquidateCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3LiquidateCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3LiquidateCreditAccount represents a LiquidateCreditAccount event raised by the CreditFacadev3 contract.
type CreditFacadev3LiquidateCreditAccount struct {
	CreditAccount  common.Address
	Liquidator     common.Address
	To             common.Address
	RemainingFunds *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLiquidateCreditAccount is a free log retrieval operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed creditAccount, address indexed liquidator, address to, uint256 remainingFunds)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterLiquidateCreditAccount(opts *bind.FilterOpts, creditAccount []common.Address, liquidator []common.Address) (*CreditFacadev3LiquidateCreditAccountIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "LiquidateCreditAccount", creditAccountRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3LiquidateCreditAccountIterator{contract: _CreditFacadev3.contract, event: "LiquidateCreditAccount", logs: logs, sub: sub}, nil
}

// WatchLiquidateCreditAccount is a free log subscription operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed creditAccount, address indexed liquidator, address to, uint256 remainingFunds)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchLiquidateCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadev3LiquidateCreditAccount, creditAccount []common.Address, liquidator []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "LiquidateCreditAccount", creditAccountRule, liquidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3LiquidateCreditAccount)
				if err := _CreditFacadev3.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
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

// ParseLiquidateCreditAccount is a log parse operation binding the contract event 0x7dfecd8419723a9d3954585a30c2a270165d70aafa146c11c1e1b88ae1439064.
//
// Solidity: event LiquidateCreditAccount(address indexed creditAccount, address indexed liquidator, address to, uint256 remainingFunds)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseLiquidateCreditAccount(log types.Log) (*CreditFacadev3LiquidateCreditAccount, error) {
	event := new(CreditFacadev3LiquidateCreditAccount)
	if err := _CreditFacadev3.contract.UnpackLog(event, "LiquidateCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3NewControllerIterator is returned from FilterNewController and is used to iterate over the raw logs and unpacked data for NewController events raised by the CreditFacadev3 contract.
type CreditFacadev3NewControllerIterator struct {
	Event *CreditFacadev3NewController // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3NewControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3NewController)
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
		it.Event = new(CreditFacadev3NewController)
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
func (it *CreditFacadev3NewControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3NewControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3NewController represents a NewController event raised by the CreditFacadev3 contract.
type CreditFacadev3NewController struct {
	NewController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewController is a free log retrieval operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterNewController(opts *bind.FilterOpts, newController []common.Address) (*CreditFacadev3NewControllerIterator, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3NewControllerIterator{contract: _CreditFacadev3.contract, event: "NewController", logs: logs, sub: sub}, nil
}

// WatchNewController is a free log subscription operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchNewController(opts *bind.WatchOpts, sink chan<- *CreditFacadev3NewController, newController []common.Address) (event.Subscription, error) {

	var newControllerRule []interface{}
	for _, newControllerItem := range newController {
		newControllerRule = append(newControllerRule, newControllerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "NewController", newControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3NewController)
				if err := _CreditFacadev3.contract.UnpackLog(event, "NewController", log); err != nil {
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

// ParseNewController is a log parse operation binding the contract event 0xe253457d9ad994ca9682fc3bbc38c890dca73a2d5ecee3809e548bac8b00d7c6.
//
// Solidity: event NewController(address indexed newController)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseNewController(log types.Log) (*CreditFacadev3NewController, error) {
	event := new(CreditFacadev3NewController)
	if err := _CreditFacadev3.contract.UnpackLog(event, "NewController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3OpenCreditAccountIterator is returned from FilterOpenCreditAccount and is used to iterate over the raw logs and unpacked data for OpenCreditAccount events raised by the CreditFacadev3 contract.
type CreditFacadev3OpenCreditAccountIterator struct {
	Event *CreditFacadev3OpenCreditAccount // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3OpenCreditAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3OpenCreditAccount)
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
		it.Event = new(CreditFacadev3OpenCreditAccount)
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
func (it *CreditFacadev3OpenCreditAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3OpenCreditAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3OpenCreditAccount represents a OpenCreditAccount event raised by the CreditFacadev3 contract.
type CreditFacadev3OpenCreditAccount struct {
	CreditAccount common.Address
	OnBehalfOf    common.Address
	Caller        common.Address
	ReferralCode  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenCreditAccount is a free log retrieval operation binding the contract event 0x6e4927aac3383b13ffc5b6f44447693caf351f2f7ca800c9b4463b76997911b0.
//
// Solidity: event OpenCreditAccount(address indexed creditAccount, address indexed onBehalfOf, address indexed caller, uint256 referralCode)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterOpenCreditAccount(opts *bind.FilterOpts, creditAccount []common.Address, onBehalfOf []common.Address, caller []common.Address) (*CreditFacadev3OpenCreditAccountIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "OpenCreditAccount", creditAccountRule, onBehalfOfRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3OpenCreditAccountIterator{contract: _CreditFacadev3.contract, event: "OpenCreditAccount", logs: logs, sub: sub}, nil
}

// WatchOpenCreditAccount is a free log subscription operation binding the contract event 0x6e4927aac3383b13ffc5b6f44447693caf351f2f7ca800c9b4463b76997911b0.
//
// Solidity: event OpenCreditAccount(address indexed creditAccount, address indexed onBehalfOf, address indexed caller, uint256 referralCode)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchOpenCreditAccount(opts *bind.WatchOpts, sink chan<- *CreditFacadev3OpenCreditAccount, creditAccount []common.Address, onBehalfOf []common.Address, caller []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "OpenCreditAccount", creditAccountRule, onBehalfOfRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3OpenCreditAccount)
				if err := _CreditFacadev3.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
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

// ParseOpenCreditAccount is a log parse operation binding the contract event 0x6e4927aac3383b13ffc5b6f44447693caf351f2f7ca800c9b4463b76997911b0.
//
// Solidity: event OpenCreditAccount(address indexed creditAccount, address indexed onBehalfOf, address indexed caller, uint256 referralCode)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseOpenCreditAccount(log types.Log) (*CreditFacadev3OpenCreditAccount, error) {
	event := new(CreditFacadev3OpenCreditAccount)
	if err := _CreditFacadev3.contract.UnpackLog(event, "OpenCreditAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditFacadev3 contract.
type CreditFacadev3PausedIterator struct {
	Event *CreditFacadev3Paused // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3Paused)
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
		it.Event = new(CreditFacadev3Paused)
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
func (it *CreditFacadev3PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3Paused represents a Paused event raised by the CreditFacadev3 contract.
type CreditFacadev3Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterPaused(opts *bind.FilterOpts) (*CreditFacadev3PausedIterator, error) {

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3PausedIterator{contract: _CreditFacadev3.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditFacadev3Paused) (event.Subscription, error) {

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3Paused)
				if err := _CreditFacadev3.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParsePaused(log types.Log) (*CreditFacadev3Paused, error) {
	event := new(CreditFacadev3Paused)
	if err := _CreditFacadev3.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3StartMultiCallIterator is returned from FilterStartMultiCall and is used to iterate over the raw logs and unpacked data for StartMultiCall events raised by the CreditFacadev3 contract.
type CreditFacadev3StartMultiCallIterator struct {
	Event *CreditFacadev3StartMultiCall // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3StartMultiCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3StartMultiCall)
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
		it.Event = new(CreditFacadev3StartMultiCall)
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
func (it *CreditFacadev3StartMultiCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3StartMultiCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3StartMultiCall represents a StartMultiCall event raised by the CreditFacadev3 contract.
type CreditFacadev3StartMultiCall struct {
	CreditAccount common.Address
	Caller        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStartMultiCall is a free log retrieval operation binding the contract event 0x6637691e02875fb5c598316278034ab86d133a75ab6d76491287290e03979284.
//
// Solidity: event StartMultiCall(address indexed creditAccount, address indexed caller)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterStartMultiCall(opts *bind.FilterOpts, creditAccount []common.Address, caller []common.Address) (*CreditFacadev3StartMultiCallIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "StartMultiCall", creditAccountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3StartMultiCallIterator{contract: _CreditFacadev3.contract, event: "StartMultiCall", logs: logs, sub: sub}, nil
}

// WatchStartMultiCall is a free log subscription operation binding the contract event 0x6637691e02875fb5c598316278034ab86d133a75ab6d76491287290e03979284.
//
// Solidity: event StartMultiCall(address indexed creditAccount, address indexed caller)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchStartMultiCall(opts *bind.WatchOpts, sink chan<- *CreditFacadev3StartMultiCall, creditAccount []common.Address, caller []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "StartMultiCall", creditAccountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3StartMultiCall)
				if err := _CreditFacadev3.contract.UnpackLog(event, "StartMultiCall", log); err != nil {
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

// ParseStartMultiCall is a log parse operation binding the contract event 0x6637691e02875fb5c598316278034ab86d133a75ab6d76491287290e03979284.
//
// Solidity: event StartMultiCall(address indexed creditAccount, address indexed caller)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseStartMultiCall(log types.Log) (*CreditFacadev3StartMultiCall, error) {
	event := new(CreditFacadev3StartMultiCall)
	if err := _CreditFacadev3.contract.UnpackLog(event, "StartMultiCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditFacadev3 contract.
type CreditFacadev3UnpausedIterator struct {
	Event *CreditFacadev3Unpaused // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3Unpaused)
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
		it.Event = new(CreditFacadev3Unpaused)
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
func (it *CreditFacadev3UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3Unpaused represents a Unpaused event raised by the CreditFacadev3 contract.
type CreditFacadev3Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditFacadev3UnpausedIterator, error) {

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3UnpausedIterator{contract: _CreditFacadev3.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditFacadev3Unpaused) (event.Subscription, error) {

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3Unpaused)
				if err := _CreditFacadev3.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseUnpaused(log types.Log) (*CreditFacadev3Unpaused, error) {
	event := new(CreditFacadev3Unpaused)
	if err := _CreditFacadev3.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditFacadev3WithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the CreditFacadev3 contract.
type CreditFacadev3WithdrawCollateralIterator struct {
	Event *CreditFacadev3WithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *CreditFacadev3WithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditFacadev3WithdrawCollateral)
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
		it.Event = new(CreditFacadev3WithdrawCollateral)
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
func (it *CreditFacadev3WithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditFacadev3WithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditFacadev3WithdrawCollateral represents a WithdrawCollateral event raised by the CreditFacadev3 contract.
type CreditFacadev3WithdrawCollateral struct {
	CreditAccount common.Address
	Token         common.Address
	Amount        *big.Int
	To            common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0xe7655dfddd0226889710c711da4e725dd44525fb5717b2321017a97d32793ab8.
//
// Solidity: event WithdrawCollateral(address indexed creditAccount, address indexed token, uint256 amount, address to)
func (_CreditFacadev3 *CreditFacadev3Filterer) FilterWithdrawCollateral(opts *bind.FilterOpts, creditAccount []common.Address, token []common.Address) (*CreditFacadev3WithdrawCollateralIterator, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacadev3.contract.FilterLogs(opts, "WithdrawCollateral", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &CreditFacadev3WithdrawCollateralIterator{contract: _CreditFacadev3.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0xe7655dfddd0226889710c711da4e725dd44525fb5717b2321017a97d32793ab8.
//
// Solidity: event WithdrawCollateral(address indexed creditAccount, address indexed token, uint256 amount, address to)
func (_CreditFacadev3 *CreditFacadev3Filterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *CreditFacadev3WithdrawCollateral, creditAccount []common.Address, token []common.Address) (event.Subscription, error) {

	var creditAccountRule []interface{}
	for _, creditAccountItem := range creditAccount {
		creditAccountRule = append(creditAccountRule, creditAccountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CreditFacadev3.contract.WatchLogs(opts, "WithdrawCollateral", creditAccountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditFacadev3WithdrawCollateral)
				if err := _CreditFacadev3.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0xe7655dfddd0226889710c711da4e725dd44525fb5717b2321017a97d32793ab8.
//
// Solidity: event WithdrawCollateral(address indexed creditAccount, address indexed token, uint256 amount, address to)
func (_CreditFacadev3 *CreditFacadev3Filterer) ParseWithdrawCollateral(log types.Log) (*CreditFacadev3WithdrawCollateral, error) {
	event := new(CreditFacadev3WithdrawCollateral)
	if err := _CreditFacadev3.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
