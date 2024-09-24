// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mellowR7

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

// IVaultProcessWithdrawalsStack is an auto generated low-level Go binding around an user-defined struct.
type IVaultProcessWithdrawalsStack struct {
	Tokens         []common.Address
	RatiosX96      []*big.Int
	Erc20Balances  []*big.Int
	TotalSupply    *big.Int
	TotalValue     *big.Int
	RatiosX96Value *big.Int
	Timestamp      *big.Int
	FeeD9          *big.Int
	TokensHash     [32]byte
}

// IVaultWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type IVaultWithdrawalRequest struct {
	To         common.Address
	LpAmount   *big.Int
	TokensHash [32]byte
	MinAmounts []*big.Int
	Deadline   *big.Int
	Timestamp  *big.Int
}

// MellowR7MetaData contains all meta data concerning the MellowR7 contract.
var MellowR7MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Deadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Forbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLpAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"name\":\"DelegateCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"DepositCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIVault.WithdrawalRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"EmergencyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"name\":\"ExternalCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TvlModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TvlModuleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"WithdrawCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"WithdrawalRequestCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIVault.WithdrawalRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"statuses\",\"type\":\"bool[]\"}],\"name\":\"WithdrawalsProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_DELEGATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"D9\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Q96\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTvlModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"ratiosX96\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256[]\",\"name\":\"erc20Balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratiosX96Value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeD9\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIVault.ProcessWithdrawalsStack\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.WithdrawalRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"analyzeRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"expectedAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTvl\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateStack\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"ratiosX96\",\"type\":\"uint128[]\"},{\"internalType\":\"uint256[]\",\"name\":\"erc20Balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratiosX96Value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeD9\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIVault.ProcessWithdrawalsStack\",\"name\":\"s\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelWithdrawalRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurator\",\"outputs\":[{\"internalType\":\"contractIVaultConfigurator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minLpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"actualAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"actualAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"externalCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isUnderlyingToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isUnderlying\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"pendingWithdrawers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingWithdrawers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingWithdrawersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"processWithdrawals\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"statuses\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestDeadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"closePrevious\",\"type\":\"bool\"}],\"name\":\"registerWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTvlModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"requireAdmin\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"requireAtLeastOperator\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tvlModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTvl\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"withdrawalRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"tokensHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIVault.WithdrawalRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MellowR7ABI is the input ABI used to generate the binding from.
// Deprecated: Use MellowR7MetaData.ABI instead.
var MellowR7ABI = MellowR7MetaData.ABI

// MellowR7 is an auto generated Go binding around an Ethereum contract.
type MellowR7 struct {
	MellowR7Caller     // Read-only binding to the contract
	MellowR7Transactor // Write-only binding to the contract
	MellowR7Filterer   // Log filterer for contract events
}

// MellowR7Caller is an auto generated read-only Go binding around an Ethereum contract.
type MellowR7Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MellowR7Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MellowR7Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MellowR7Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MellowR7Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MellowR7Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MellowR7Session struct {
	Contract     *MellowR7         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MellowR7CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MellowR7CallerSession struct {
	Contract *MellowR7Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MellowR7TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MellowR7TransactorSession struct {
	Contract     *MellowR7Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MellowR7Raw is an auto generated low-level Go binding around an Ethereum contract.
type MellowR7Raw struct {
	Contract *MellowR7 // Generic contract binding to access the raw methods on
}

// MellowR7CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MellowR7CallerRaw struct {
	Contract *MellowR7Caller // Generic read-only contract binding to access the raw methods on
}

// MellowR7TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MellowR7TransactorRaw struct {
	Contract *MellowR7Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMellowR7 creates a new instance of MellowR7, bound to a specific deployed contract.
func NewMellowR7(address common.Address, backend bind.ContractBackend) (*MellowR7, error) {
	contract, err := bindMellowR7(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MellowR7{MellowR7Caller: MellowR7Caller{contract: contract}, MellowR7Transactor: MellowR7Transactor{contract: contract}, MellowR7Filterer: MellowR7Filterer{contract: contract}}, nil
}

// NewMellowR7Caller creates a new read-only instance of MellowR7, bound to a specific deployed contract.
func NewMellowR7Caller(address common.Address, caller bind.ContractCaller) (*MellowR7Caller, error) {
	contract, err := bindMellowR7(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MellowR7Caller{contract: contract}, nil
}

// NewMellowR7Transactor creates a new write-only instance of MellowR7, bound to a specific deployed contract.
func NewMellowR7Transactor(address common.Address, transactor bind.ContractTransactor) (*MellowR7Transactor, error) {
	contract, err := bindMellowR7(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MellowR7Transactor{contract: contract}, nil
}

// NewMellowR7Filterer creates a new log filterer instance of MellowR7, bound to a specific deployed contract.
func NewMellowR7Filterer(address common.Address, filterer bind.ContractFilterer) (*MellowR7Filterer, error) {
	contract, err := bindMellowR7(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MellowR7Filterer{contract: contract}, nil
}

// bindMellowR7 binds a generic wrapper to an already deployed contract.
func bindMellowR7(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MellowR7MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MellowR7 *MellowR7Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MellowR7.Contract.MellowR7Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MellowR7 *MellowR7Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MellowR7.Contract.MellowR7Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MellowR7 *MellowR7Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MellowR7.Contract.MellowR7Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MellowR7 *MellowR7CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MellowR7.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MellowR7 *MellowR7TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MellowR7.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MellowR7 *MellowR7TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MellowR7.Contract.contract.Transact(opts, method, params...)
}

// ADMINDELEGATEROLE is a free data retrieval call binding the contract method 0x0952ff54.
//
// Solidity: function ADMIN_DELEGATE_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Caller) ADMINDELEGATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "ADMIN_DELEGATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINDELEGATEROLE is a free data retrieval call binding the contract method 0x0952ff54.
//
// Solidity: function ADMIN_DELEGATE_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Session) ADMINDELEGATEROLE() ([32]byte, error) {
	return _MellowR7.Contract.ADMINDELEGATEROLE(&_MellowR7.CallOpts)
}

// ADMINDELEGATEROLE is a free data retrieval call binding the contract method 0x0952ff54.
//
// Solidity: function ADMIN_DELEGATE_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7CallerSession) ADMINDELEGATEROLE() ([32]byte, error) {
	return _MellowR7.Contract.ADMINDELEGATEROLE(&_MellowR7.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Caller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Session) ADMINROLE() ([32]byte, error) {
	return _MellowR7.Contract.ADMINROLE(&_MellowR7.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7CallerSession) ADMINROLE() ([32]byte, error) {
	return _MellowR7.Contract.ADMINROLE(&_MellowR7.CallOpts)
}

// D9 is a free data retrieval call binding the contract method 0x551530ae.
//
// Solidity: function D9() view returns(uint256)
func (_MellowR7 *MellowR7Caller) D9(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "D9")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D9 is a free data retrieval call binding the contract method 0x551530ae.
//
// Solidity: function D9() view returns(uint256)
func (_MellowR7 *MellowR7Session) D9() (*big.Int, error) {
	return _MellowR7.Contract.D9(&_MellowR7.CallOpts)
}

// D9 is a free data retrieval call binding the contract method 0x551530ae.
//
// Solidity: function D9() view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) D9() (*big.Int, error) {
	return _MellowR7.Contract.D9(&_MellowR7.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _MellowR7.Contract.DEFAULTADMINROLE(&_MellowR7.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MellowR7 *MellowR7CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MellowR7.Contract.DEFAULTADMINROLE(&_MellowR7.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_MellowR7 *MellowR7Caller) OPERATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "OPERATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_MellowR7 *MellowR7Session) OPERATOR() ([32]byte, error) {
	return _MellowR7.Contract.OPERATOR(&_MellowR7.CallOpts)
}

// OPERATOR is a free data retrieval call binding the contract method 0x983d2737.
//
// Solidity: function OPERATOR() view returns(bytes32)
func (_MellowR7 *MellowR7CallerSession) OPERATOR() ([32]byte, error) {
	return _MellowR7.Contract.OPERATOR(&_MellowR7.CallOpts)
}

// Q96 is a free data retrieval call binding the contract method 0x40792465.
//
// Solidity: function Q96() view returns(uint256)
func (_MellowR7 *MellowR7Caller) Q96(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "Q96")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Q96 is a free data retrieval call binding the contract method 0x40792465.
//
// Solidity: function Q96() view returns(uint256)
func (_MellowR7 *MellowR7Session) Q96() (*big.Int, error) {
	return _MellowR7.Contract.Q96(&_MellowR7.CallOpts)
}

// Q96 is a free data retrieval call binding the contract method 0x40792465.
//
// Solidity: function Q96() view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) Q96() (*big.Int, error) {
	return _MellowR7.Contract.Q96(&_MellowR7.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MellowR7 *MellowR7Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MellowR7 *MellowR7Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MellowR7.Contract.Allowance(&_MellowR7.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MellowR7.Contract.Allowance(&_MellowR7.CallOpts, owner, spender)
}

// AnalyzeRequest is a free data retrieval call binding the contract method 0xb5f828a2.
//
// Solidity: function analyzeRequest((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s, (address,uint256,bytes32,uint256[],uint256,uint256) request) pure returns(bool, bool, uint256[] expectedAmounts)
func (_MellowR7 *MellowR7Caller) AnalyzeRequest(opts *bind.CallOpts, s IVaultProcessWithdrawalsStack, request IVaultWithdrawalRequest) (bool, bool, []*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "analyzeRequest", s, request)

	if err != nil {
		return *new(bool), *new(bool), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// AnalyzeRequest is a free data retrieval call binding the contract method 0xb5f828a2.
//
// Solidity: function analyzeRequest((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s, (address,uint256,bytes32,uint256[],uint256,uint256) request) pure returns(bool, bool, uint256[] expectedAmounts)
func (_MellowR7 *MellowR7Session) AnalyzeRequest(s IVaultProcessWithdrawalsStack, request IVaultWithdrawalRequest) (bool, bool, []*big.Int, error) {
	return _MellowR7.Contract.AnalyzeRequest(&_MellowR7.CallOpts, s, request)
}

// AnalyzeRequest is a free data retrieval call binding the contract method 0xb5f828a2.
//
// Solidity: function analyzeRequest((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s, (address,uint256,bytes32,uint256[],uint256,uint256) request) pure returns(bool, bool, uint256[] expectedAmounts)
func (_MellowR7 *MellowR7CallerSession) AnalyzeRequest(s IVaultProcessWithdrawalsStack, request IVaultWithdrawalRequest) (bool, bool, []*big.Int, error) {
	return _MellowR7.Contract.AnalyzeRequest(&_MellowR7.CallOpts, s, request)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MellowR7 *MellowR7Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MellowR7 *MellowR7Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _MellowR7.Contract.BalanceOf(&_MellowR7.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MellowR7.Contract.BalanceOf(&_MellowR7.CallOpts, account)
}

// BaseTvl is a free data retrieval call binding the contract method 0xb07f08f8.
//
// Solidity: function baseTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7Caller) BaseTvl(opts *bind.CallOpts) (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "baseTvl")

	outstruct := new(struct {
		Tokens  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// BaseTvl is a free data retrieval call binding the contract method 0xb07f08f8.
//
// Solidity: function baseTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7Session) BaseTvl() (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _MellowR7.Contract.BaseTvl(&_MellowR7.CallOpts)
}

// BaseTvl is a free data retrieval call binding the contract method 0xb07f08f8.
//
// Solidity: function baseTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7CallerSession) BaseTvl() (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _MellowR7.Contract.BaseTvl(&_MellowR7.CallOpts)
}

// CalculateStack is a free data retrieval call binding the contract method 0x36cbf6bc.
//
// Solidity: function calculateStack() view returns((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s)
func (_MellowR7 *MellowR7Caller) CalculateStack(opts *bind.CallOpts) (IVaultProcessWithdrawalsStack, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "calculateStack")

	if err != nil {
		return *new(IVaultProcessWithdrawalsStack), err
	}

	out0 := *abi.ConvertType(out[0], new(IVaultProcessWithdrawalsStack)).(*IVaultProcessWithdrawalsStack)

	return out0, err

}

// CalculateStack is a free data retrieval call binding the contract method 0x36cbf6bc.
//
// Solidity: function calculateStack() view returns((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s)
func (_MellowR7 *MellowR7Session) CalculateStack() (IVaultProcessWithdrawalsStack, error) {
	return _MellowR7.Contract.CalculateStack(&_MellowR7.CallOpts)
}

// CalculateStack is a free data retrieval call binding the contract method 0x36cbf6bc.
//
// Solidity: function calculateStack() view returns((address[],uint128[],uint256[],uint256,uint256,uint256,uint256,uint256,bytes32) s)
func (_MellowR7 *MellowR7CallerSession) CalculateStack() (IVaultProcessWithdrawalsStack, error) {
	return _MellowR7.Contract.CalculateStack(&_MellowR7.CallOpts)
}

// Configurator is a free data retrieval call binding the contract method 0x2b507df8.
//
// Solidity: function configurator() view returns(address)
func (_MellowR7 *MellowR7Caller) Configurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "configurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Configurator is a free data retrieval call binding the contract method 0x2b507df8.
//
// Solidity: function configurator() view returns(address)
func (_MellowR7 *MellowR7Session) Configurator() (common.Address, error) {
	return _MellowR7.Contract.Configurator(&_MellowR7.CallOpts)
}

// Configurator is a free data retrieval call binding the contract method 0x2b507df8.
//
// Solidity: function configurator() view returns(address)
func (_MellowR7 *MellowR7CallerSession) Configurator() (common.Address, error) {
	return _MellowR7.Contract.Configurator(&_MellowR7.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MellowR7 *MellowR7Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MellowR7 *MellowR7Session) Decimals() (uint8, error) {
	return _MellowR7.Contract.Decimals(&_MellowR7.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MellowR7 *MellowR7CallerSession) Decimals() (uint8, error) {
	return _MellowR7.Contract.Decimals(&_MellowR7.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MellowR7 *MellowR7Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MellowR7 *MellowR7Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MellowR7.Contract.GetRoleAdmin(&_MellowR7.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MellowR7 *MellowR7CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MellowR7.Contract.GetRoleAdmin(&_MellowR7.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MellowR7 *MellowR7Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MellowR7 *MellowR7Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MellowR7.Contract.GetRoleMember(&_MellowR7.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MellowR7 *MellowR7CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MellowR7.Contract.GetRoleMember(&_MellowR7.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MellowR7 *MellowR7Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MellowR7 *MellowR7Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MellowR7.Contract.GetRoleMemberCount(&_MellowR7.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MellowR7.Contract.GetRoleMemberCount(&_MellowR7.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MellowR7 *MellowR7Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MellowR7 *MellowR7Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MellowR7.Contract.HasRole(&_MellowR7.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MellowR7 *MellowR7CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MellowR7.Contract.HasRole(&_MellowR7.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address sender) view returns(bool)
func (_MellowR7 *MellowR7Caller) IsAdmin(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "isAdmin", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address sender) view returns(bool)
func (_MellowR7 *MellowR7Session) IsAdmin(sender common.Address) (bool, error) {
	return _MellowR7.Contract.IsAdmin(&_MellowR7.CallOpts, sender)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address sender) view returns(bool)
func (_MellowR7 *MellowR7CallerSession) IsAdmin(sender common.Address) (bool, error) {
	return _MellowR7.Contract.IsAdmin(&_MellowR7.CallOpts, sender)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address sender) view returns(bool)
func (_MellowR7 *MellowR7Caller) IsOperator(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "isOperator", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address sender) view returns(bool)
func (_MellowR7 *MellowR7Session) IsOperator(sender common.Address) (bool, error) {
	return _MellowR7.Contract.IsOperator(&_MellowR7.CallOpts, sender)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address sender) view returns(bool)
func (_MellowR7 *MellowR7CallerSession) IsOperator(sender common.Address) (bool, error) {
	return _MellowR7.Contract.IsOperator(&_MellowR7.CallOpts, sender)
}

// IsUnderlyingToken is a free data retrieval call binding the contract method 0xa8bad207.
//
// Solidity: function isUnderlyingToken(address token) view returns(bool isUnderlying)
func (_MellowR7 *MellowR7Caller) IsUnderlyingToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "isUnderlyingToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnderlyingToken is a free data retrieval call binding the contract method 0xa8bad207.
//
// Solidity: function isUnderlyingToken(address token) view returns(bool isUnderlying)
func (_MellowR7 *MellowR7Session) IsUnderlyingToken(token common.Address) (bool, error) {
	return _MellowR7.Contract.IsUnderlyingToken(&_MellowR7.CallOpts, token)
}

// IsUnderlyingToken is a free data retrieval call binding the contract method 0xa8bad207.
//
// Solidity: function isUnderlyingToken(address token) view returns(bool isUnderlying)
func (_MellowR7 *MellowR7CallerSession) IsUnderlyingToken(token common.Address) (bool, error) {
	return _MellowR7.Contract.IsUnderlyingToken(&_MellowR7.CallOpts, token)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MellowR7 *MellowR7Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MellowR7 *MellowR7Session) Name() (string, error) {
	return _MellowR7.Contract.Name(&_MellowR7.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MellowR7 *MellowR7CallerSession) Name() (string, error) {
	return _MellowR7.Contract.Name(&_MellowR7.CallOpts)
}

// PendingWithdrawers is a free data retrieval call binding the contract method 0x07395b69.
//
// Solidity: function pendingWithdrawers(uint256 limit, uint256 offset) view returns(address[] result)
func (_MellowR7 *MellowR7Caller) PendingWithdrawers(opts *bind.CallOpts, limit *big.Int, offset *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "pendingWithdrawers", limit, offset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PendingWithdrawers is a free data retrieval call binding the contract method 0x07395b69.
//
// Solidity: function pendingWithdrawers(uint256 limit, uint256 offset) view returns(address[] result)
func (_MellowR7 *MellowR7Session) PendingWithdrawers(limit *big.Int, offset *big.Int) ([]common.Address, error) {
	return _MellowR7.Contract.PendingWithdrawers(&_MellowR7.CallOpts, limit, offset)
}

// PendingWithdrawers is a free data retrieval call binding the contract method 0x07395b69.
//
// Solidity: function pendingWithdrawers(uint256 limit, uint256 offset) view returns(address[] result)
func (_MellowR7 *MellowR7CallerSession) PendingWithdrawers(limit *big.Int, offset *big.Int) ([]common.Address, error) {
	return _MellowR7.Contract.PendingWithdrawers(&_MellowR7.CallOpts, limit, offset)
}

// PendingWithdrawers0 is a free data retrieval call binding the contract method 0xf9e0cd28.
//
// Solidity: function pendingWithdrawers() view returns(address[])
func (_MellowR7 *MellowR7Caller) PendingWithdrawers0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "pendingWithdrawers0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PendingWithdrawers0 is a free data retrieval call binding the contract method 0xf9e0cd28.
//
// Solidity: function pendingWithdrawers() view returns(address[])
func (_MellowR7 *MellowR7Session) PendingWithdrawers0() ([]common.Address, error) {
	return _MellowR7.Contract.PendingWithdrawers0(&_MellowR7.CallOpts)
}

// PendingWithdrawers0 is a free data retrieval call binding the contract method 0xf9e0cd28.
//
// Solidity: function pendingWithdrawers() view returns(address[])
func (_MellowR7 *MellowR7CallerSession) PendingWithdrawers0() ([]common.Address, error) {
	return _MellowR7.Contract.PendingWithdrawers0(&_MellowR7.CallOpts)
}

// PendingWithdrawersCount is a free data retrieval call binding the contract method 0x429bef10.
//
// Solidity: function pendingWithdrawersCount() view returns(uint256)
func (_MellowR7 *MellowR7Caller) PendingWithdrawersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "pendingWithdrawersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawersCount is a free data retrieval call binding the contract method 0x429bef10.
//
// Solidity: function pendingWithdrawersCount() view returns(uint256)
func (_MellowR7 *MellowR7Session) PendingWithdrawersCount() (*big.Int, error) {
	return _MellowR7.Contract.PendingWithdrawersCount(&_MellowR7.CallOpts)
}

// PendingWithdrawersCount is a free data retrieval call binding the contract method 0x429bef10.
//
// Solidity: function pendingWithdrawersCount() view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) PendingWithdrawersCount() (*big.Int, error) {
	return _MellowR7.Contract.PendingWithdrawersCount(&_MellowR7.CallOpts)
}

// RequireAdmin is a free data retrieval call binding the contract method 0x5577210a.
//
// Solidity: function requireAdmin(address sender) view returns()
func (_MellowR7 *MellowR7Caller) RequireAdmin(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "requireAdmin", sender)

	if err != nil {
		return err
	}

	return err

}

// RequireAdmin is a free data retrieval call binding the contract method 0x5577210a.
//
// Solidity: function requireAdmin(address sender) view returns()
func (_MellowR7 *MellowR7Session) RequireAdmin(sender common.Address) error {
	return _MellowR7.Contract.RequireAdmin(&_MellowR7.CallOpts, sender)
}

// RequireAdmin is a free data retrieval call binding the contract method 0x5577210a.
//
// Solidity: function requireAdmin(address sender) view returns()
func (_MellowR7 *MellowR7CallerSession) RequireAdmin(sender common.Address) error {
	return _MellowR7.Contract.RequireAdmin(&_MellowR7.CallOpts, sender)
}

// RequireAtLeastOperator is a free data retrieval call binding the contract method 0x0a2a8aea.
//
// Solidity: function requireAtLeastOperator(address sender) view returns()
func (_MellowR7 *MellowR7Caller) RequireAtLeastOperator(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "requireAtLeastOperator", sender)

	if err != nil {
		return err
	}

	return err

}

// RequireAtLeastOperator is a free data retrieval call binding the contract method 0x0a2a8aea.
//
// Solidity: function requireAtLeastOperator(address sender) view returns()
func (_MellowR7 *MellowR7Session) RequireAtLeastOperator(sender common.Address) error {
	return _MellowR7.Contract.RequireAtLeastOperator(&_MellowR7.CallOpts, sender)
}

// RequireAtLeastOperator is a free data retrieval call binding the contract method 0x0a2a8aea.
//
// Solidity: function requireAtLeastOperator(address sender) view returns()
func (_MellowR7 *MellowR7CallerSession) RequireAtLeastOperator(sender common.Address) error {
	return _MellowR7.Contract.RequireAtLeastOperator(&_MellowR7.CallOpts, sender)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MellowR7 *MellowR7Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MellowR7 *MellowR7Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MellowR7.Contract.SupportsInterface(&_MellowR7.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MellowR7 *MellowR7CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MellowR7.Contract.SupportsInterface(&_MellowR7.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MellowR7 *MellowR7Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MellowR7 *MellowR7Session) Symbol() (string, error) {
	return _MellowR7.Contract.Symbol(&_MellowR7.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MellowR7 *MellowR7CallerSession) Symbol() (string, error) {
	return _MellowR7.Contract.Symbol(&_MellowR7.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MellowR7 *MellowR7Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MellowR7 *MellowR7Session) TotalSupply() (*big.Int, error) {
	return _MellowR7.Contract.TotalSupply(&_MellowR7.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MellowR7 *MellowR7CallerSession) TotalSupply() (*big.Int, error) {
	return _MellowR7.Contract.TotalSupply(&_MellowR7.CallOpts)
}

// TvlModules is a free data retrieval call binding the contract method 0xb46dde2a.
//
// Solidity: function tvlModules() view returns(address[])
func (_MellowR7 *MellowR7Caller) TvlModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "tvlModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TvlModules is a free data retrieval call binding the contract method 0xb46dde2a.
//
// Solidity: function tvlModules() view returns(address[])
func (_MellowR7 *MellowR7Session) TvlModules() ([]common.Address, error) {
	return _MellowR7.Contract.TvlModules(&_MellowR7.CallOpts)
}

// TvlModules is a free data retrieval call binding the contract method 0xb46dde2a.
//
// Solidity: function tvlModules() view returns(address[])
func (_MellowR7 *MellowR7CallerSession) TvlModules() ([]common.Address, error) {
	return _MellowR7.Contract.TvlModules(&_MellowR7.CallOpts)
}

// UnderlyingTokens is a free data retrieval call binding the contract method 0xbd27dc9f.
//
// Solidity: function underlyingTokens() view returns(address[])
func (_MellowR7 *MellowR7Caller) UnderlyingTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "underlyingTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// UnderlyingTokens is a free data retrieval call binding the contract method 0xbd27dc9f.
//
// Solidity: function underlyingTokens() view returns(address[])
func (_MellowR7 *MellowR7Session) UnderlyingTokens() ([]common.Address, error) {
	return _MellowR7.Contract.UnderlyingTokens(&_MellowR7.CallOpts)
}

// UnderlyingTokens is a free data retrieval call binding the contract method 0xbd27dc9f.
//
// Solidity: function underlyingTokens() view returns(address[])
func (_MellowR7 *MellowR7CallerSession) UnderlyingTokens() ([]common.Address, error) {
	return _MellowR7.Contract.UnderlyingTokens(&_MellowR7.CallOpts)
}

// UnderlyingTvl is a free data retrieval call binding the contract method 0x079c3b88.
//
// Solidity: function underlyingTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7Caller) UnderlyingTvl(opts *bind.CallOpts) (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "underlyingTvl")

	outstruct := new(struct {
		Tokens  []common.Address
		Amounts []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Amounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// UnderlyingTvl is a free data retrieval call binding the contract method 0x079c3b88.
//
// Solidity: function underlyingTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7Session) UnderlyingTvl() (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _MellowR7.Contract.UnderlyingTvl(&_MellowR7.CallOpts)
}

// UnderlyingTvl is a free data retrieval call binding the contract method 0x079c3b88.
//
// Solidity: function underlyingTvl() view returns(address[] tokens, uint256[] amounts)
func (_MellowR7 *MellowR7CallerSession) UnderlyingTvl() (struct {
	Tokens  []common.Address
	Amounts []*big.Int
}, error) {
	return _MellowR7.Contract.UnderlyingTvl(&_MellowR7.CallOpts)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x1ef44e68.
//
// Solidity: function withdrawalRequest(address user) view returns((address,uint256,bytes32,uint256[],uint256,uint256))
func (_MellowR7 *MellowR7Caller) WithdrawalRequest(opts *bind.CallOpts, user common.Address) (IVaultWithdrawalRequest, error) {
	var out []interface{}
	err := _MellowR7.contract.Call(opts, &out, "withdrawalRequest", user)

	if err != nil {
		return *new(IVaultWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IVaultWithdrawalRequest)).(*IVaultWithdrawalRequest)

	return out0, err

}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x1ef44e68.
//
// Solidity: function withdrawalRequest(address user) view returns((address,uint256,bytes32,uint256[],uint256,uint256))
func (_MellowR7 *MellowR7Session) WithdrawalRequest(user common.Address) (IVaultWithdrawalRequest, error) {
	return _MellowR7.Contract.WithdrawalRequest(&_MellowR7.CallOpts, user)
}

// WithdrawalRequest is a free data retrieval call binding the contract method 0x1ef44e68.
//
// Solidity: function withdrawalRequest(address user) view returns((address,uint256,bytes32,uint256[],uint256,uint256))
func (_MellowR7 *MellowR7CallerSession) WithdrawalRequest(user common.Address) (IVaultWithdrawalRequest, error) {
	return _MellowR7.Contract.WithdrawalRequest(&_MellowR7.CallOpts, user)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_MellowR7 *MellowR7Transactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_MellowR7 *MellowR7Session) AddToken(token common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.AddToken(&_MellowR7.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns()
func (_MellowR7 *MellowR7TransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.AddToken(&_MellowR7.TransactOpts, token)
}

// AddTvlModule is a paid mutator transaction binding the contract method 0x15a196f8.
//
// Solidity: function addTvlModule(address module) returns()
func (_MellowR7 *MellowR7Transactor) AddTvlModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "addTvlModule", module)
}

// AddTvlModule is a paid mutator transaction binding the contract method 0x15a196f8.
//
// Solidity: function addTvlModule(address module) returns()
func (_MellowR7 *MellowR7Session) AddTvlModule(module common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.AddTvlModule(&_MellowR7.TransactOpts, module)
}

// AddTvlModule is a paid mutator transaction binding the contract method 0x15a196f8.
//
// Solidity: function addTvlModule(address module) returns()
func (_MellowR7 *MellowR7TransactorSession) AddTvlModule(module common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.AddTvlModule(&_MellowR7.TransactOpts, module)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Approve(&_MellowR7.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MellowR7 *MellowR7TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Approve(&_MellowR7.TransactOpts, spender, value)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0xe714a028.
//
// Solidity: function cancelWithdrawalRequest() returns()
func (_MellowR7 *MellowR7Transactor) CancelWithdrawalRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "cancelWithdrawalRequest")
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0xe714a028.
//
// Solidity: function cancelWithdrawalRequest() returns()
func (_MellowR7 *MellowR7Session) CancelWithdrawalRequest() (*types.Transaction, error) {
	return _MellowR7.Contract.CancelWithdrawalRequest(&_MellowR7.TransactOpts)
}

// CancelWithdrawalRequest is a paid mutator transaction binding the contract method 0xe714a028.
//
// Solidity: function cancelWithdrawalRequest() returns()
func (_MellowR7 *MellowR7TransactorSession) CancelWithdrawalRequest() (*types.Transaction, error) {
	return _MellowR7.Contract.CancelWithdrawalRequest(&_MellowR7.TransactOpts)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x56e7b7aa.
//
// Solidity: function delegateCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7Transactor) DelegateCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "delegateCall", to, data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x56e7b7aa.
//
// Solidity: function delegateCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7Session) DelegateCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.Contract.DelegateCall(&_MellowR7.TransactOpts, to, data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x56e7b7aa.
//
// Solidity: function delegateCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7TransactorSession) DelegateCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.Contract.DelegateCall(&_MellowR7.TransactOpts, to, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x5f097d40.
//
// Solidity: function deposit(address to, uint256[] amounts, uint256 minLpAmount, uint256 deadline) returns(uint256[] actualAmounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Transactor) Deposit(opts *bind.TransactOpts, to common.Address, amounts []*big.Int, minLpAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "deposit", to, amounts, minLpAmount, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x5f097d40.
//
// Solidity: function deposit(address to, uint256[] amounts, uint256 minLpAmount, uint256 deadline) returns(uint256[] actualAmounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Session) Deposit(to common.Address, amounts []*big.Int, minLpAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Deposit(&_MellowR7.TransactOpts, to, amounts, minLpAmount, deadline)
}

// Deposit is a paid mutator transaction binding the contract method 0x5f097d40.
//
// Solidity: function deposit(address to, uint256[] amounts, uint256 minLpAmount, uint256 deadline) returns(uint256[] actualAmounts, uint256 lpAmount)
func (_MellowR7 *MellowR7TransactorSession) Deposit(to common.Address, amounts []*big.Int, minLpAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Deposit(&_MellowR7.TransactOpts, to, amounts, minLpAmount, deadline)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x66e04b85.
//
// Solidity: function emergencyWithdraw(uint256[] minAmounts, uint256 deadline) returns(uint256[] actualAmounts)
func (_MellowR7 *MellowR7Transactor) EmergencyWithdraw(opts *bind.TransactOpts, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "emergencyWithdraw", minAmounts, deadline)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x66e04b85.
//
// Solidity: function emergencyWithdraw(uint256[] minAmounts, uint256 deadline) returns(uint256[] actualAmounts)
func (_MellowR7 *MellowR7Session) EmergencyWithdraw(minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.EmergencyWithdraw(&_MellowR7.TransactOpts, minAmounts, deadline)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x66e04b85.
//
// Solidity: function emergencyWithdraw(uint256[] minAmounts, uint256 deadline) returns(uint256[] actualAmounts)
func (_MellowR7 *MellowR7TransactorSession) EmergencyWithdraw(minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.EmergencyWithdraw(&_MellowR7.TransactOpts, minAmounts, deadline)
}

// ExternalCall is a paid mutator transaction binding the contract method 0x654c9bdb.
//
// Solidity: function externalCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7Transactor) ExternalCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "externalCall", to, data)
}

// ExternalCall is a paid mutator transaction binding the contract method 0x654c9bdb.
//
// Solidity: function externalCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7Session) ExternalCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.Contract.ExternalCall(&_MellowR7.TransactOpts, to, data)
}

// ExternalCall is a paid mutator transaction binding the contract method 0x654c9bdb.
//
// Solidity: function externalCall(address to, bytes data) returns(bool success, bytes response)
func (_MellowR7 *MellowR7TransactorSession) ExternalCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _MellowR7.Contract.ExternalCall(&_MellowR7.TransactOpts, to, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.GrantRole(&_MellowR7.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.GrantRole(&_MellowR7.TransactOpts, role, account)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0xaade917c.
//
// Solidity: function processWithdrawals(address[] users) returns(bool[] statuses)
func (_MellowR7 *MellowR7Transactor) ProcessWithdrawals(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "processWithdrawals", users)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0xaade917c.
//
// Solidity: function processWithdrawals(address[] users) returns(bool[] statuses)
func (_MellowR7 *MellowR7Session) ProcessWithdrawals(users []common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.ProcessWithdrawals(&_MellowR7.TransactOpts, users)
}

// ProcessWithdrawals is a paid mutator transaction binding the contract method 0xaade917c.
//
// Solidity: function processWithdrawals(address[] users) returns(bool[] statuses)
func (_MellowR7 *MellowR7TransactorSession) ProcessWithdrawals(users []common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.ProcessWithdrawals(&_MellowR7.TransactOpts, users)
}

// RegisterWithdrawal is a paid mutator transaction binding the contract method 0x2e3d46dc.
//
// Solidity: function registerWithdrawal(address to, uint256 lpAmount, uint256[] minAmounts, uint256 deadline, uint256 requestDeadline, bool closePrevious) returns()
func (_MellowR7 *MellowR7Transactor) RegisterWithdrawal(opts *bind.TransactOpts, to common.Address, lpAmount *big.Int, minAmounts []*big.Int, deadline *big.Int, requestDeadline *big.Int, closePrevious bool) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "registerWithdrawal", to, lpAmount, minAmounts, deadline, requestDeadline, closePrevious)
}

// RegisterWithdrawal is a paid mutator transaction binding the contract method 0x2e3d46dc.
//
// Solidity: function registerWithdrawal(address to, uint256 lpAmount, uint256[] minAmounts, uint256 deadline, uint256 requestDeadline, bool closePrevious) returns()
func (_MellowR7 *MellowR7Session) RegisterWithdrawal(to common.Address, lpAmount *big.Int, minAmounts []*big.Int, deadline *big.Int, requestDeadline *big.Int, closePrevious bool) (*types.Transaction, error) {
	return _MellowR7.Contract.RegisterWithdrawal(&_MellowR7.TransactOpts, to, lpAmount, minAmounts, deadline, requestDeadline, closePrevious)
}

// RegisterWithdrawal is a paid mutator transaction binding the contract method 0x2e3d46dc.
//
// Solidity: function registerWithdrawal(address to, uint256 lpAmount, uint256[] minAmounts, uint256 deadline, uint256 requestDeadline, bool closePrevious) returns()
func (_MellowR7 *MellowR7TransactorSession) RegisterWithdrawal(to common.Address, lpAmount *big.Int, minAmounts []*big.Int, deadline *big.Int, requestDeadline *big.Int, closePrevious bool) (*types.Transaction, error) {
	return _MellowR7.Contract.RegisterWithdrawal(&_MellowR7.TransactOpts, to, lpAmount, minAmounts, deadline, requestDeadline, closePrevious)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_MellowR7 *MellowR7Transactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_MellowR7 *MellowR7Session) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RemoveToken(&_MellowR7.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns()
func (_MellowR7 *MellowR7TransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RemoveToken(&_MellowR7.TransactOpts, token)
}

// RemoveTvlModule is a paid mutator transaction binding the contract method 0x7c367609.
//
// Solidity: function removeTvlModule(address module) returns()
func (_MellowR7 *MellowR7Transactor) RemoveTvlModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "removeTvlModule", module)
}

// RemoveTvlModule is a paid mutator transaction binding the contract method 0x7c367609.
//
// Solidity: function removeTvlModule(address module) returns()
func (_MellowR7 *MellowR7Session) RemoveTvlModule(module common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RemoveTvlModule(&_MellowR7.TransactOpts, module)
}

// RemoveTvlModule is a paid mutator transaction binding the contract method 0x7c367609.
//
// Solidity: function removeTvlModule(address module) returns()
func (_MellowR7 *MellowR7TransactorSession) RemoveTvlModule(module common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RemoveTvlModule(&_MellowR7.TransactOpts, module)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MellowR7 *MellowR7Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MellowR7 *MellowR7Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RenounceRole(&_MellowR7.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MellowR7 *MellowR7TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RenounceRole(&_MellowR7.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RevokeRole(&_MellowR7.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MellowR7 *MellowR7TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MellowR7.Contract.RevokeRole(&_MellowR7.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Transfer(&_MellowR7.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.Transfer(&_MellowR7.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.TransferFrom(&_MellowR7.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MellowR7 *MellowR7TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MellowR7.Contract.TransferFrom(&_MellowR7.TransactOpts, from, to, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MellowR7 *MellowR7Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MellowR7.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MellowR7 *MellowR7Session) Receive() (*types.Transaction, error) {
	return _MellowR7.Contract.Receive(&_MellowR7.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MellowR7 *MellowR7TransactorSession) Receive() (*types.Transaction, error) {
	return _MellowR7.Contract.Receive(&_MellowR7.TransactOpts)
}

// MellowR7ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MellowR7 contract.
type MellowR7ApprovalIterator struct {
	Event *MellowR7Approval // Event containing the contract specifics and raw log

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
func (it *MellowR7ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7Approval)
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
		it.Event = new(MellowR7Approval)
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
func (it *MellowR7ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7Approval represents a Approval event raised by the MellowR7 contract.
type MellowR7Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MellowR7 *MellowR7Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MellowR7ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7ApprovalIterator{contract: _MellowR7.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MellowR7 *MellowR7Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MellowR7Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7Approval)
				if err := _MellowR7.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MellowR7 *MellowR7Filterer) ParseApproval(log types.Log) (*MellowR7Approval, error) {
	event := new(MellowR7Approval)
	if err := _MellowR7.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7DelegateCallIterator is returned from FilterDelegateCall and is used to iterate over the raw logs and unpacked data for DelegateCall events raised by the MellowR7 contract.
type MellowR7DelegateCallIterator struct {
	Event *MellowR7DelegateCall // Event containing the contract specifics and raw log

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
func (it *MellowR7DelegateCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7DelegateCall)
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
		it.Event = new(MellowR7DelegateCall)
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
func (it *MellowR7DelegateCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7DelegateCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7DelegateCall represents a DelegateCall event raised by the MellowR7 contract.
type MellowR7DelegateCall struct {
	To       common.Address
	Data     []byte
	Success  bool
	Response []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateCall is a free log retrieval operation binding the contract event 0x6cb85343e8c77cb180eb29d20c41302a16fb818828fe672c47916295caf80d52.
//
// Solidity: event DelegateCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) FilterDelegateCall(opts *bind.FilterOpts, to []common.Address) (*MellowR7DelegateCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "DelegateCall", toRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7DelegateCallIterator{contract: _MellowR7.contract, event: "DelegateCall", logs: logs, sub: sub}, nil
}

// WatchDelegateCall is a free log subscription operation binding the contract event 0x6cb85343e8c77cb180eb29d20c41302a16fb818828fe672c47916295caf80d52.
//
// Solidity: event DelegateCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) WatchDelegateCall(opts *bind.WatchOpts, sink chan<- *MellowR7DelegateCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "DelegateCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7DelegateCall)
				if err := _MellowR7.contract.UnpackLog(event, "DelegateCall", log); err != nil {
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

// ParseDelegateCall is a log parse operation binding the contract event 0x6cb85343e8c77cb180eb29d20c41302a16fb818828fe672c47916295caf80d52.
//
// Solidity: event DelegateCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) ParseDelegateCall(log types.Log) (*MellowR7DelegateCall, error) {
	event := new(MellowR7DelegateCall)
	if err := _MellowR7.contract.UnpackLog(event, "DelegateCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MellowR7 contract.
type MellowR7DepositIterator struct {
	Event *MellowR7Deposit // Event containing the contract specifics and raw log

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
func (it *MellowR7DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7Deposit)
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
		it.Event = new(MellowR7Deposit)
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
func (it *MellowR7DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7Deposit represents a Deposit event raised by the MellowR7 contract.
type MellowR7Deposit struct {
	To       common.Address
	Amounts  []*big.Int
	LpAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xff195810018e2867a43eaac646e6b3fc71bc32d776175995704b6bc10d7fada8.
//
// Solidity: event Deposit(address indexed to, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) FilterDeposit(opts *bind.FilterOpts, to []common.Address) (*MellowR7DepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7DepositIterator{contract: _MellowR7.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xff195810018e2867a43eaac646e6b3fc71bc32d776175995704b6bc10d7fada8.
//
// Solidity: event Deposit(address indexed to, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MellowR7Deposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7Deposit)
				if err := _MellowR7.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xff195810018e2867a43eaac646e6b3fc71bc32d776175995704b6bc10d7fada8.
//
// Solidity: event Deposit(address indexed to, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) ParseDeposit(log types.Log) (*MellowR7Deposit, error) {
	event := new(MellowR7Deposit)
	if err := _MellowR7.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7DepositCallbackIterator is returned from FilterDepositCallback and is used to iterate over the raw logs and unpacked data for DepositCallback events raised by the MellowR7 contract.
type MellowR7DepositCallbackIterator struct {
	Event *MellowR7DepositCallback // Event containing the contract specifics and raw log

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
func (it *MellowR7DepositCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7DepositCallback)
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
		it.Event = new(MellowR7DepositCallback)
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
func (it *MellowR7DepositCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7DepositCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7DepositCallback represents a DepositCallback event raised by the MellowR7 contract.
type MellowR7DepositCallback struct {
	Callback common.Address
	Amounts  []*big.Int
	LpAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDepositCallback is a free log retrieval operation binding the contract event 0xf9d4830984aa759ba52886b6e857f5127160f28040c406c6fdafdc86f9f7ab9d.
//
// Solidity: event DepositCallback(address indexed callback, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) FilterDepositCallback(opts *bind.FilterOpts, callback []common.Address) (*MellowR7DepositCallbackIterator, error) {

	var callbackRule []interface{}
	for _, callbackItem := range callback {
		callbackRule = append(callbackRule, callbackItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "DepositCallback", callbackRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7DepositCallbackIterator{contract: _MellowR7.contract, event: "DepositCallback", logs: logs, sub: sub}, nil
}

// WatchDepositCallback is a free log subscription operation binding the contract event 0xf9d4830984aa759ba52886b6e857f5127160f28040c406c6fdafdc86f9f7ab9d.
//
// Solidity: event DepositCallback(address indexed callback, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) WatchDepositCallback(opts *bind.WatchOpts, sink chan<- *MellowR7DepositCallback, callback []common.Address) (event.Subscription, error) {

	var callbackRule []interface{}
	for _, callbackItem := range callback {
		callbackRule = append(callbackRule, callbackItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "DepositCallback", callbackRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7DepositCallback)
				if err := _MellowR7.contract.UnpackLog(event, "DepositCallback", log); err != nil {
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

// ParseDepositCallback is a log parse operation binding the contract event 0xf9d4830984aa759ba52886b6e857f5127160f28040c406c6fdafdc86f9f7ab9d.
//
// Solidity: event DepositCallback(address indexed callback, uint256[] amounts, uint256 lpAmount)
func (_MellowR7 *MellowR7Filterer) ParseDepositCallback(log types.Log) (*MellowR7DepositCallback, error) {
	event := new(MellowR7DepositCallback)
	if err := _MellowR7.contract.UnpackLog(event, "DepositCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7EmergencyWithdrawalIterator is returned from FilterEmergencyWithdrawal and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawal events raised by the MellowR7 contract.
type MellowR7EmergencyWithdrawalIterator struct {
	Event *MellowR7EmergencyWithdrawal // Event containing the contract specifics and raw log

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
func (it *MellowR7EmergencyWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7EmergencyWithdrawal)
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
		it.Event = new(MellowR7EmergencyWithdrawal)
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
func (it *MellowR7EmergencyWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7EmergencyWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7EmergencyWithdrawal represents a EmergencyWithdrawal event raised by the MellowR7 contract.
type MellowR7EmergencyWithdrawal struct {
	From    common.Address
	Request IVaultWithdrawalRequest
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawal is a free log retrieval operation binding the contract event 0x73d5d80561fc62718447463fe9794e6fb7d472229521cce725e75bef7ddd62e8.
//
// Solidity: event EmergencyWithdrawal(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request, uint256[] amounts)
func (_MellowR7 *MellowR7Filterer) FilterEmergencyWithdrawal(opts *bind.FilterOpts, from []common.Address) (*MellowR7EmergencyWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "EmergencyWithdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7EmergencyWithdrawalIterator{contract: _MellowR7.contract, event: "EmergencyWithdrawal", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawal is a free log subscription operation binding the contract event 0x73d5d80561fc62718447463fe9794e6fb7d472229521cce725e75bef7ddd62e8.
//
// Solidity: event EmergencyWithdrawal(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request, uint256[] amounts)
func (_MellowR7 *MellowR7Filterer) WatchEmergencyWithdrawal(opts *bind.WatchOpts, sink chan<- *MellowR7EmergencyWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "EmergencyWithdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7EmergencyWithdrawal)
				if err := _MellowR7.contract.UnpackLog(event, "EmergencyWithdrawal", log); err != nil {
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

// ParseEmergencyWithdrawal is a log parse operation binding the contract event 0x73d5d80561fc62718447463fe9794e6fb7d472229521cce725e75bef7ddd62e8.
//
// Solidity: event EmergencyWithdrawal(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request, uint256[] amounts)
func (_MellowR7 *MellowR7Filterer) ParseEmergencyWithdrawal(log types.Log) (*MellowR7EmergencyWithdrawal, error) {
	event := new(MellowR7EmergencyWithdrawal)
	if err := _MellowR7.contract.UnpackLog(event, "EmergencyWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7ExternalCallIterator is returned from FilterExternalCall and is used to iterate over the raw logs and unpacked data for ExternalCall events raised by the MellowR7 contract.
type MellowR7ExternalCallIterator struct {
	Event *MellowR7ExternalCall // Event containing the contract specifics and raw log

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
func (it *MellowR7ExternalCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7ExternalCall)
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
		it.Event = new(MellowR7ExternalCall)
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
func (it *MellowR7ExternalCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7ExternalCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7ExternalCall represents a ExternalCall event raised by the MellowR7 contract.
type MellowR7ExternalCall struct {
	To       common.Address
	Data     []byte
	Success  bool
	Response []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExternalCall is a free log retrieval operation binding the contract event 0x2f4d74076cbab7a79d7579b5ba2c106b1ac26462ac617d3a2079ca16a3bc7351.
//
// Solidity: event ExternalCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) FilterExternalCall(opts *bind.FilterOpts, to []common.Address) (*MellowR7ExternalCallIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "ExternalCall", toRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7ExternalCallIterator{contract: _MellowR7.contract, event: "ExternalCall", logs: logs, sub: sub}, nil
}

// WatchExternalCall is a free log subscription operation binding the contract event 0x2f4d74076cbab7a79d7579b5ba2c106b1ac26462ac617d3a2079ca16a3bc7351.
//
// Solidity: event ExternalCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) WatchExternalCall(opts *bind.WatchOpts, sink chan<- *MellowR7ExternalCall, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "ExternalCall", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7ExternalCall)
				if err := _MellowR7.contract.UnpackLog(event, "ExternalCall", log); err != nil {
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

// ParseExternalCall is a log parse operation binding the contract event 0x2f4d74076cbab7a79d7579b5ba2c106b1ac26462ac617d3a2079ca16a3bc7351.
//
// Solidity: event ExternalCall(address indexed to, bytes data, bool success, bytes response)
func (_MellowR7 *MellowR7Filterer) ParseExternalCall(log types.Log) (*MellowR7ExternalCall, error) {
	event := new(MellowR7ExternalCall)
	if err := _MellowR7.contract.UnpackLog(event, "ExternalCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MellowR7 contract.
type MellowR7RoleAdminChangedIterator struct {
	Event *MellowR7RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MellowR7RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7RoleAdminChanged)
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
		it.Event = new(MellowR7RoleAdminChanged)
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
func (it *MellowR7RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7RoleAdminChanged represents a RoleAdminChanged event raised by the MellowR7 contract.
type MellowR7RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MellowR7 *MellowR7Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MellowR7RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7RoleAdminChangedIterator{contract: _MellowR7.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MellowR7 *MellowR7Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MellowR7RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7RoleAdminChanged)
				if err := _MellowR7.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MellowR7 *MellowR7Filterer) ParseRoleAdminChanged(log types.Log) (*MellowR7RoleAdminChanged, error) {
	event := new(MellowR7RoleAdminChanged)
	if err := _MellowR7.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MellowR7 contract.
type MellowR7RoleGrantedIterator struct {
	Event *MellowR7RoleGranted // Event containing the contract specifics and raw log

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
func (it *MellowR7RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7RoleGranted)
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
		it.Event = new(MellowR7RoleGranted)
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
func (it *MellowR7RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7RoleGranted represents a RoleGranted event raised by the MellowR7 contract.
type MellowR7RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MellowR7RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7RoleGrantedIterator{contract: _MellowR7.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MellowR7RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7RoleGranted)
				if err := _MellowR7.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) ParseRoleGranted(log types.Log) (*MellowR7RoleGranted, error) {
	event := new(MellowR7RoleGranted)
	if err := _MellowR7.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MellowR7 contract.
type MellowR7RoleRevokedIterator struct {
	Event *MellowR7RoleRevoked // Event containing the contract specifics and raw log

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
func (it *MellowR7RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7RoleRevoked)
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
		it.Event = new(MellowR7RoleRevoked)
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
func (it *MellowR7RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7RoleRevoked represents a RoleRevoked event raised by the MellowR7 contract.
type MellowR7RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MellowR7RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7RoleRevokedIterator{contract: _MellowR7.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MellowR7RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7RoleRevoked)
				if err := _MellowR7.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MellowR7 *MellowR7Filterer) ParseRoleRevoked(log types.Log) (*MellowR7RoleRevoked, error) {
	event := new(MellowR7RoleRevoked)
	if err := _MellowR7.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7TokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the MellowR7 contract.
type MellowR7TokenAddedIterator struct {
	Event *MellowR7TokenAdded // Event containing the contract specifics and raw log

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
func (it *MellowR7TokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7TokenAdded)
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
		it.Event = new(MellowR7TokenAdded)
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
func (it *MellowR7TokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7TokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7TokenAdded represents a TokenAdded event raised by the MellowR7 contract.
type MellowR7TokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_MellowR7 *MellowR7Filterer) FilterTokenAdded(opts *bind.FilterOpts) (*MellowR7TokenAddedIterator, error) {

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &MellowR7TokenAddedIterator{contract: _MellowR7.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_MellowR7 *MellowR7Filterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *MellowR7TokenAdded) (event.Subscription, error) {

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7TokenAdded)
				if err := _MellowR7.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address token)
func (_MellowR7 *MellowR7Filterer) ParseTokenAdded(log types.Log) (*MellowR7TokenAdded, error) {
	event := new(MellowR7TokenAdded)
	if err := _MellowR7.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7TokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the MellowR7 contract.
type MellowR7TokenRemovedIterator struct {
	Event *MellowR7TokenRemoved // Event containing the contract specifics and raw log

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
func (it *MellowR7TokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7TokenRemoved)
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
		it.Event = new(MellowR7TokenRemoved)
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
func (it *MellowR7TokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7TokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7TokenRemoved represents a TokenRemoved event raised by the MellowR7 contract.
type MellowR7TokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_MellowR7 *MellowR7Filterer) FilterTokenRemoved(opts *bind.FilterOpts) (*MellowR7TokenRemovedIterator, error) {

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return &MellowR7TokenRemovedIterator{contract: _MellowR7.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_MellowR7 *MellowR7Filterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *MellowR7TokenRemoved) (event.Subscription, error) {

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "TokenRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7TokenRemoved)
				if err := _MellowR7.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address token)
func (_MellowR7 *MellowR7Filterer) ParseTokenRemoved(log types.Log) (*MellowR7TokenRemoved, error) {
	event := new(MellowR7TokenRemoved)
	if err := _MellowR7.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MellowR7 contract.
type MellowR7TransferIterator struct {
	Event *MellowR7Transfer // Event containing the contract specifics and raw log

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
func (it *MellowR7TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7Transfer)
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
		it.Event = new(MellowR7Transfer)
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
func (it *MellowR7TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7Transfer represents a Transfer event raised by the MellowR7 contract.
type MellowR7Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MellowR7 *MellowR7Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MellowR7TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7TransferIterator{contract: _MellowR7.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MellowR7 *MellowR7Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MellowR7Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7Transfer)
				if err := _MellowR7.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MellowR7 *MellowR7Filterer) ParseTransfer(log types.Log) (*MellowR7Transfer, error) {
	event := new(MellowR7Transfer)
	if err := _MellowR7.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7TvlModuleAddedIterator is returned from FilterTvlModuleAdded and is used to iterate over the raw logs and unpacked data for TvlModuleAdded events raised by the MellowR7 contract.
type MellowR7TvlModuleAddedIterator struct {
	Event *MellowR7TvlModuleAdded // Event containing the contract specifics and raw log

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
func (it *MellowR7TvlModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7TvlModuleAdded)
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
		it.Event = new(MellowR7TvlModuleAdded)
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
func (it *MellowR7TvlModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7TvlModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7TvlModuleAdded represents a TvlModuleAdded event raised by the MellowR7 contract.
type MellowR7TvlModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTvlModuleAdded is a free log retrieval operation binding the contract event 0xd1d220e49bdf761bd526b4ba7434f65d7d1498ad06053831a219e3e496328c62.
//
// Solidity: event TvlModuleAdded(address module)
func (_MellowR7 *MellowR7Filterer) FilterTvlModuleAdded(opts *bind.FilterOpts) (*MellowR7TvlModuleAddedIterator, error) {

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "TvlModuleAdded")
	if err != nil {
		return nil, err
	}
	return &MellowR7TvlModuleAddedIterator{contract: _MellowR7.contract, event: "TvlModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTvlModuleAdded is a free log subscription operation binding the contract event 0xd1d220e49bdf761bd526b4ba7434f65d7d1498ad06053831a219e3e496328c62.
//
// Solidity: event TvlModuleAdded(address module)
func (_MellowR7 *MellowR7Filterer) WatchTvlModuleAdded(opts *bind.WatchOpts, sink chan<- *MellowR7TvlModuleAdded) (event.Subscription, error) {

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "TvlModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7TvlModuleAdded)
				if err := _MellowR7.contract.UnpackLog(event, "TvlModuleAdded", log); err != nil {
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

// ParseTvlModuleAdded is a log parse operation binding the contract event 0xd1d220e49bdf761bd526b4ba7434f65d7d1498ad06053831a219e3e496328c62.
//
// Solidity: event TvlModuleAdded(address module)
func (_MellowR7 *MellowR7Filterer) ParseTvlModuleAdded(log types.Log) (*MellowR7TvlModuleAdded, error) {
	event := new(MellowR7TvlModuleAdded)
	if err := _MellowR7.contract.UnpackLog(event, "TvlModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7TvlModuleRemovedIterator is returned from FilterTvlModuleRemoved and is used to iterate over the raw logs and unpacked data for TvlModuleRemoved events raised by the MellowR7 contract.
type MellowR7TvlModuleRemovedIterator struct {
	Event *MellowR7TvlModuleRemoved // Event containing the contract specifics and raw log

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
func (it *MellowR7TvlModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7TvlModuleRemoved)
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
		it.Event = new(MellowR7TvlModuleRemoved)
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
func (it *MellowR7TvlModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7TvlModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7TvlModuleRemoved represents a TvlModuleRemoved event raised by the MellowR7 contract.
type MellowR7TvlModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTvlModuleRemoved is a free log retrieval operation binding the contract event 0x0a4ec75dfa617891caee29ac64c80bdd4e2d2f7b67dbf1ba8288bf0ca5622d89.
//
// Solidity: event TvlModuleRemoved(address module)
func (_MellowR7 *MellowR7Filterer) FilterTvlModuleRemoved(opts *bind.FilterOpts) (*MellowR7TvlModuleRemovedIterator, error) {

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "TvlModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &MellowR7TvlModuleRemovedIterator{contract: _MellowR7.contract, event: "TvlModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTvlModuleRemoved is a free log subscription operation binding the contract event 0x0a4ec75dfa617891caee29ac64c80bdd4e2d2f7b67dbf1ba8288bf0ca5622d89.
//
// Solidity: event TvlModuleRemoved(address module)
func (_MellowR7 *MellowR7Filterer) WatchTvlModuleRemoved(opts *bind.WatchOpts, sink chan<- *MellowR7TvlModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "TvlModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7TvlModuleRemoved)
				if err := _MellowR7.contract.UnpackLog(event, "TvlModuleRemoved", log); err != nil {
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

// ParseTvlModuleRemoved is a log parse operation binding the contract event 0x0a4ec75dfa617891caee29ac64c80bdd4e2d2f7b67dbf1ba8288bf0ca5622d89.
//
// Solidity: event TvlModuleRemoved(address module)
func (_MellowR7 *MellowR7Filterer) ParseTvlModuleRemoved(log types.Log) (*MellowR7TvlModuleRemoved, error) {
	event := new(MellowR7TvlModuleRemoved)
	if err := _MellowR7.contract.UnpackLog(event, "TvlModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7WithdrawCallbackIterator is returned from FilterWithdrawCallback and is used to iterate over the raw logs and unpacked data for WithdrawCallback events raised by the MellowR7 contract.
type MellowR7WithdrawCallbackIterator struct {
	Event *MellowR7WithdrawCallback // Event containing the contract specifics and raw log

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
func (it *MellowR7WithdrawCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7WithdrawCallback)
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
		it.Event = new(MellowR7WithdrawCallback)
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
func (it *MellowR7WithdrawCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7WithdrawCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7WithdrawCallback represents a WithdrawCallback event raised by the MellowR7 contract.
type MellowR7WithdrawCallback struct {
	Callback common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCallback is a free log retrieval operation binding the contract event 0xdbdc898ae09e91ee802f5ebdced25d8398a1bef800c6210cb8b397761437ef90.
//
// Solidity: event WithdrawCallback(address indexed callback)
func (_MellowR7 *MellowR7Filterer) FilterWithdrawCallback(opts *bind.FilterOpts, callback []common.Address) (*MellowR7WithdrawCallbackIterator, error) {

	var callbackRule []interface{}
	for _, callbackItem := range callback {
		callbackRule = append(callbackRule, callbackItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "WithdrawCallback", callbackRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7WithdrawCallbackIterator{contract: _MellowR7.contract, event: "WithdrawCallback", logs: logs, sub: sub}, nil
}

// WatchWithdrawCallback is a free log subscription operation binding the contract event 0xdbdc898ae09e91ee802f5ebdced25d8398a1bef800c6210cb8b397761437ef90.
//
// Solidity: event WithdrawCallback(address indexed callback)
func (_MellowR7 *MellowR7Filterer) WatchWithdrawCallback(opts *bind.WatchOpts, sink chan<- *MellowR7WithdrawCallback, callback []common.Address) (event.Subscription, error) {

	var callbackRule []interface{}
	for _, callbackItem := range callback {
		callbackRule = append(callbackRule, callbackItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "WithdrawCallback", callbackRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7WithdrawCallback)
				if err := _MellowR7.contract.UnpackLog(event, "WithdrawCallback", log); err != nil {
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

// ParseWithdrawCallback is a log parse operation binding the contract event 0xdbdc898ae09e91ee802f5ebdced25d8398a1bef800c6210cb8b397761437ef90.
//
// Solidity: event WithdrawCallback(address indexed callback)
func (_MellowR7 *MellowR7Filterer) ParseWithdrawCallback(log types.Log) (*MellowR7WithdrawCallback, error) {
	event := new(MellowR7WithdrawCallback)
	if err := _MellowR7.contract.UnpackLog(event, "WithdrawCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7WithdrawalRequestCanceledIterator is returned from FilterWithdrawalRequestCanceled and is used to iterate over the raw logs and unpacked data for WithdrawalRequestCanceled events raised by the MellowR7 contract.
type MellowR7WithdrawalRequestCanceledIterator struct {
	Event *MellowR7WithdrawalRequestCanceled // Event containing the contract specifics and raw log

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
func (it *MellowR7WithdrawalRequestCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7WithdrawalRequestCanceled)
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
		it.Event = new(MellowR7WithdrawalRequestCanceled)
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
func (it *MellowR7WithdrawalRequestCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7WithdrawalRequestCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7WithdrawalRequestCanceled represents a WithdrawalRequestCanceled event raised by the MellowR7 contract.
type MellowR7WithdrawalRequestCanceled struct {
	User   common.Address
	Origin common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequestCanceled is a free log retrieval operation binding the contract event 0xced4ad6a94ad9f17d9e89d24158930a071a5de6c1d6cddf799bac950b04fc09b.
//
// Solidity: event WithdrawalRequestCanceled(address indexed user, address origin)
func (_MellowR7 *MellowR7Filterer) FilterWithdrawalRequestCanceled(opts *bind.FilterOpts, user []common.Address) (*MellowR7WithdrawalRequestCanceledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "WithdrawalRequestCanceled", userRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7WithdrawalRequestCanceledIterator{contract: _MellowR7.contract, event: "WithdrawalRequestCanceled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequestCanceled is a free log subscription operation binding the contract event 0xced4ad6a94ad9f17d9e89d24158930a071a5de6c1d6cddf799bac950b04fc09b.
//
// Solidity: event WithdrawalRequestCanceled(address indexed user, address origin)
func (_MellowR7 *MellowR7Filterer) WatchWithdrawalRequestCanceled(opts *bind.WatchOpts, sink chan<- *MellowR7WithdrawalRequestCanceled, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "WithdrawalRequestCanceled", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7WithdrawalRequestCanceled)
				if err := _MellowR7.contract.UnpackLog(event, "WithdrawalRequestCanceled", log); err != nil {
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

// ParseWithdrawalRequestCanceled is a log parse operation binding the contract event 0xced4ad6a94ad9f17d9e89d24158930a071a5de6c1d6cddf799bac950b04fc09b.
//
// Solidity: event WithdrawalRequestCanceled(address indexed user, address origin)
func (_MellowR7 *MellowR7Filterer) ParseWithdrawalRequestCanceled(log types.Log) (*MellowR7WithdrawalRequestCanceled, error) {
	event := new(MellowR7WithdrawalRequestCanceled)
	if err := _MellowR7.contract.UnpackLog(event, "WithdrawalRequestCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7WithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the MellowR7 contract.
type MellowR7WithdrawalRequestedIterator struct {
	Event *MellowR7WithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *MellowR7WithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7WithdrawalRequested)
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
		it.Event = new(MellowR7WithdrawalRequested)
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
func (it *MellowR7WithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7WithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7WithdrawalRequested represents a WithdrawalRequested event raised by the MellowR7 contract.
type MellowR7WithdrawalRequested struct {
	From    common.Address
	Request IVaultWithdrawalRequest
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0x390acd8a6485f0c0b379ba564699e9a4fd3ed26cd9d957a50d56320b4bf2309b.
//
// Solidity: event WithdrawalRequested(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request)
func (_MellowR7 *MellowR7Filterer) FilterWithdrawalRequested(opts *bind.FilterOpts, from []common.Address) (*MellowR7WithdrawalRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "WithdrawalRequested", fromRule)
	if err != nil {
		return nil, err
	}
	return &MellowR7WithdrawalRequestedIterator{contract: _MellowR7.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0x390acd8a6485f0c0b379ba564699e9a4fd3ed26cd9d957a50d56320b4bf2309b.
//
// Solidity: event WithdrawalRequested(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request)
func (_MellowR7 *MellowR7Filterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *MellowR7WithdrawalRequested, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "WithdrawalRequested", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7WithdrawalRequested)
				if err := _MellowR7.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0x390acd8a6485f0c0b379ba564699e9a4fd3ed26cd9d957a50d56320b4bf2309b.
//
// Solidity: event WithdrawalRequested(address indexed from, (address,uint256,bytes32,uint256[],uint256,uint256) request)
func (_MellowR7 *MellowR7Filterer) ParseWithdrawalRequested(log types.Log) (*MellowR7WithdrawalRequested, error) {
	event := new(MellowR7WithdrawalRequested)
	if err := _MellowR7.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MellowR7WithdrawalsProcessedIterator is returned from FilterWithdrawalsProcessed and is used to iterate over the raw logs and unpacked data for WithdrawalsProcessed events raised by the MellowR7 contract.
type MellowR7WithdrawalsProcessedIterator struct {
	Event *MellowR7WithdrawalsProcessed // Event containing the contract specifics and raw log

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
func (it *MellowR7WithdrawalsProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MellowR7WithdrawalsProcessed)
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
		it.Event = new(MellowR7WithdrawalsProcessed)
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
func (it *MellowR7WithdrawalsProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MellowR7WithdrawalsProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MellowR7WithdrawalsProcessed represents a WithdrawalsProcessed event raised by the MellowR7 contract.
type MellowR7WithdrawalsProcessed struct {
	Users    []common.Address
	Statuses []bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsProcessed is a free log retrieval operation binding the contract event 0x9dd911706834ae347234c3bcfc7c9e4b74ba1f489af4c39c4481c4b527d4a63f.
//
// Solidity: event WithdrawalsProcessed(address[] users, bool[] statuses)
func (_MellowR7 *MellowR7Filterer) FilterWithdrawalsProcessed(opts *bind.FilterOpts) (*MellowR7WithdrawalsProcessedIterator, error) {

	logs, sub, err := _MellowR7.contract.FilterLogs(opts, "WithdrawalsProcessed")
	if err != nil {
		return nil, err
	}
	return &MellowR7WithdrawalsProcessedIterator{contract: _MellowR7.contract, event: "WithdrawalsProcessed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsProcessed is a free log subscription operation binding the contract event 0x9dd911706834ae347234c3bcfc7c9e4b74ba1f489af4c39c4481c4b527d4a63f.
//
// Solidity: event WithdrawalsProcessed(address[] users, bool[] statuses)
func (_MellowR7 *MellowR7Filterer) WatchWithdrawalsProcessed(opts *bind.WatchOpts, sink chan<- *MellowR7WithdrawalsProcessed) (event.Subscription, error) {

	logs, sub, err := _MellowR7.contract.WatchLogs(opts, "WithdrawalsProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MellowR7WithdrawalsProcessed)
				if err := _MellowR7.contract.UnpackLog(event, "WithdrawalsProcessed", log); err != nil {
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

// ParseWithdrawalsProcessed is a log parse operation binding the contract event 0x9dd911706834ae347234c3bcfc7c9e4b74ba1f489af4c39c4481c4b527d4a63f.
//
// Solidity: event WithdrawalsProcessed(address[] users, bool[] statuses)
func (_MellowR7 *MellowR7Filterer) ParseWithdrawalsProcessed(log types.Log) (*MellowR7WithdrawalsProcessed, error) {
	event := new(MellowR7WithdrawalsProcessed)
	if err := _MellowR7.contract.UnpackLog(event, "WithdrawalsProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
