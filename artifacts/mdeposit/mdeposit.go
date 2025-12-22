// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mdeposit

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

// InstantInitParams is an auto generated low-level Go binding around an user-defined struct.
type InstantInitParams struct {
	InstantFee        *big.Int
	InstantDailyLimit *big.Int
}

// MTokenInitParams is an auto generated low-level Go binding around an user-defined struct.
type MTokenInitParams struct {
	MToken         common.Address
	MTokenDataFeed common.Address
}

// ReceiversInitParams is an auto generated low-level Go binding around an user-defined struct.
type ReceiversInitParams struct {
	TokensReceiver common.Address
	FeeReceiver    common.Address
}

// MdepositMetaData contains all meta data concerning the Mdeposit contract.
var MdepositMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataFeed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"AddPaymentToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AddWaivedFeeAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOutRate\",\"type\":\"uint256\"}],\"name\":\"ApproveRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"name\":\"ChangeTokenAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ChangeTokenFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"DepositInstant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"DepositInstantWithCustomRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenOutRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"DepositRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenOutRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"DepositRequestWithCustomRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"FreeFromMinAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"FreeFromMinDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"fn\",\"type\":\"bytes4\"}],\"name\":\"PauseFn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RejectRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RemovePaymentToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RemoveWaivedFeeAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newOutRate\",\"type\":\"uint256\"}],\"name\":\"SafeApproveRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"name\":\"SetFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"SetGreenlistEnable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"SetInstantDailyLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetInstantFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetMaxSupplyCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"SetMinAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetMinMTokenAmountForFirstDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSanctionsList\",\"type\":\"address\"}],\"name\":\"SetSanctionsList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reciever\",\"type\":\"address\"}],\"name\":\"SetTokensReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTolerance\",\"type\":\"uint256\"}],\"name\":\"SetVariationTolerance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"fn\",\"type\":\"bytes4\"}],\"name\":\"UnpauseFn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLACKLISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLACKLIST_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GREENLISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GREENLIST_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANUAL_FULLFILMENT_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_EDGE_CUSTOM_AGGREGATOR_FEED_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_EDGE_DEPOSIT_VAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"M_EDGE_REDEMPTION_VAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ONE_HUNDRED_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STABLECOIN_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessControl\",\"outputs\":[{\"internalType\":\"contractMidasAccessControl\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dataFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWaivedFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newOutRate\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"}],\"name\":\"changeTokenAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"changeTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReceiveAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositInstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReceiveAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"depositInstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"}],\"name\":\"depositRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"referrerId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"depositRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"fnPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"freeFromMinAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaymentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"greenlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"greenlistTogglerRole\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"greenlistedRole\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ac\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"mToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mTokenDataFeed\",\"type\":\"address\"}],\"internalType\":\"structMTokenInitParams\",\"name\":\"_mTokenInitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokensReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structReceiversInitParams\",\"name\":\"_receiversInitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instantFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instantDailyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInstantInitParams\",\"name\":\"_instantInitParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sanctionsList\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_variationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minMTokenAmountForFirstDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupplyCap\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ac\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"mToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mTokenDataFeed\",\"type\":\"address\"}],\"internalType\":\"structMTokenInitParams\",\"name\":\"_mTokenInitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokensReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structReceiversInitParams\",\"name\":\"_receiversInitParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"instantFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instantDailyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInstantInitParams\",\"name\":\"_instantInitParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sanctionsList\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_variationTolerance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minMTokenAmountForFirstDeposit\",\"type\":\"uint256\"}],\"name\":\"initializeV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupplyCap\",\"type\":\"uint256\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantDailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"instantFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isFreeFromMinAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mToken\",\"outputs\":[{\"internalType\":\"contractIMToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mTokenDataFeed\",\"outputs\":[{\"internalType\":\"contractIDataFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupplyCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minMTokenAmountForFirstDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"enumRequestStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"depositedUsdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdAmountWithoutFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOutRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAdminRole\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"fn\",\"type\":\"bytes4\"}],\"name\":\"pauseFn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"rejectRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removePaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWaivedFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newOutRate\",\"type\":\"uint256\"}],\"name\":\"safeApproveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"safeBulkApproveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"newOutRate\",\"type\":\"uint256\"}],\"name\":\"safeBulkApproveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"safeBulkApproveRequestAtSavedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sanctionsList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sanctionsListAdminRole\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setGreenlistEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInstantDailyLimit\",\"type\":\"uint256\"}],\"name\":\"setInstantDailyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInstantFee\",\"type\":\"uint256\"}],\"name\":\"setInstantFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMaxSupplyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinMTokenAmountForFirstDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSanctionsList\",\"type\":\"address\"}],\"name\":\"setSanctionsList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setTokensReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tolerance\",\"type\":\"uint256\"}],\"name\":\"setVariationTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"dataFeed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"fn\",\"type\":\"bytes4\"}],\"name\":\"unpauseFn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"variationTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultRole\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"waivedFeeRestriction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"withdrawTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MdepositABI is the input ABI used to generate the binding from.
// Deprecated: Use MdepositMetaData.ABI instead.
var MdepositABI = MdepositMetaData.ABI

// Mdeposit is an auto generated Go binding around an Ethereum contract.
type Mdeposit struct {
	MdepositCaller     // Read-only binding to the contract
	MdepositTransactor // Write-only binding to the contract
	MdepositFilterer   // Log filterer for contract events
}

// MdepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type MdepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MdepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MdepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MdepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MdepositSession struct {
	Contract     *Mdeposit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MdepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MdepositCallerSession struct {
	Contract *MdepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MdepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MdepositTransactorSession struct {
	Contract     *MdepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MdepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type MdepositRaw struct {
	Contract *Mdeposit // Generic contract binding to access the raw methods on
}

// MdepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MdepositCallerRaw struct {
	Contract *MdepositCaller // Generic read-only contract binding to access the raw methods on
}

// MdepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MdepositTransactorRaw struct {
	Contract *MdepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMdeposit creates a new instance of Mdeposit, bound to a specific deployed contract.
func NewMdeposit(address common.Address, backend bind.ContractBackend) (*Mdeposit, error) {
	contract, err := bindMdeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mdeposit{MdepositCaller: MdepositCaller{contract: contract}, MdepositTransactor: MdepositTransactor{contract: contract}, MdepositFilterer: MdepositFilterer{contract: contract}}, nil
}

// NewMdepositCaller creates a new read-only instance of Mdeposit, bound to a specific deployed contract.
func NewMdepositCaller(address common.Address, caller bind.ContractCaller) (*MdepositCaller, error) {
	contract, err := bindMdeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MdepositCaller{contract: contract}, nil
}

// NewMdepositTransactor creates a new write-only instance of Mdeposit, bound to a specific deployed contract.
func NewMdepositTransactor(address common.Address, transactor bind.ContractTransactor) (*MdepositTransactor, error) {
	contract, err := bindMdeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MdepositTransactor{contract: contract}, nil
}

// NewMdepositFilterer creates a new log filterer instance of Mdeposit, bound to a specific deployed contract.
func NewMdepositFilterer(address common.Address, filterer bind.ContractFilterer) (*MdepositFilterer, error) {
	contract, err := bindMdeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MdepositFilterer{contract: contract}, nil
}

// bindMdeposit binds a generic wrapper to an already deployed contract.
func bindMdeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MdepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mdeposit *MdepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mdeposit.Contract.MdepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mdeposit *MdepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mdeposit.Contract.MdepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mdeposit *MdepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mdeposit.Contract.MdepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mdeposit *MdepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mdeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mdeposit *MdepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mdeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mdeposit *MdepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mdeposit.Contract.contract.Transact(opts, method, params...)
}

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) BLACKLISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "BLACKLISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) BLACKLISTEDROLE() ([32]byte, error) {
	return _Mdeposit.Contract.BLACKLISTEDROLE(&_Mdeposit.CallOpts)
}

// BLACKLISTEDROLE is a free data retrieval call binding the contract method 0x62b199c5.
//
// Solidity: function BLACKLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) BLACKLISTEDROLE() ([32]byte, error) {
	return _Mdeposit.Contract.BLACKLISTEDROLE(&_Mdeposit.CallOpts)
}

// BLACKLISTOPERATORROLE is a free data retrieval call binding the contract method 0x5300b4ba.
//
// Solidity: function BLACKLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) BLACKLISTOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "BLACKLIST_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLACKLISTOPERATORROLE is a free data retrieval call binding the contract method 0x5300b4ba.
//
// Solidity: function BLACKLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) BLACKLISTOPERATORROLE() ([32]byte, error) {
	return _Mdeposit.Contract.BLACKLISTOPERATORROLE(&_Mdeposit.CallOpts)
}

// BLACKLISTOPERATORROLE is a free data retrieval call binding the contract method 0x5300b4ba.
//
// Solidity: function BLACKLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) BLACKLISTOPERATORROLE() ([32]byte, error) {
	return _Mdeposit.Contract.BLACKLISTOPERATORROLE(&_Mdeposit.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.DEFAULTADMINROLE(&_Mdeposit.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.DEFAULTADMINROLE(&_Mdeposit.CallOpts)
}

// GREENLISTEDROLE is a free data retrieval call binding the contract method 0x1ed41163.
//
// Solidity: function GREENLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) GREENLISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "GREENLISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GREENLISTEDROLE is a free data retrieval call binding the contract method 0x1ed41163.
//
// Solidity: function GREENLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) GREENLISTEDROLE() ([32]byte, error) {
	return _Mdeposit.Contract.GREENLISTEDROLE(&_Mdeposit.CallOpts)
}

// GREENLISTEDROLE is a free data retrieval call binding the contract method 0x1ed41163.
//
// Solidity: function GREENLISTED_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) GREENLISTEDROLE() ([32]byte, error) {
	return _Mdeposit.Contract.GREENLISTEDROLE(&_Mdeposit.CallOpts)
}

// GREENLISTOPERATORROLE is a free data retrieval call binding the contract method 0x15b9598a.
//
// Solidity: function GREENLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) GREENLISTOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "GREENLIST_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GREENLISTOPERATORROLE is a free data retrieval call binding the contract method 0x15b9598a.
//
// Solidity: function GREENLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) GREENLISTOPERATORROLE() ([32]byte, error) {
	return _Mdeposit.Contract.GREENLISTOPERATORROLE(&_Mdeposit.CallOpts)
}

// GREENLISTOPERATORROLE is a free data retrieval call binding the contract method 0x15b9598a.
//
// Solidity: function GREENLIST_OPERATOR_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) GREENLISTOPERATORROLE() ([32]byte, error) {
	return _Mdeposit.Contract.GREENLISTOPERATORROLE(&_Mdeposit.CallOpts)
}

// MANUALFULLFILMENTTOKEN is a free data retrieval call binding the contract method 0x9af40265.
//
// Solidity: function MANUAL_FULLFILMENT_TOKEN() view returns(address)
func (_Mdeposit *MdepositCaller) MANUALFULLFILMENTTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "MANUAL_FULLFILMENT_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MANUALFULLFILMENTTOKEN is a free data retrieval call binding the contract method 0x9af40265.
//
// Solidity: function MANUAL_FULLFILMENT_TOKEN() view returns(address)
func (_Mdeposit *MdepositSession) MANUALFULLFILMENTTOKEN() (common.Address, error) {
	return _Mdeposit.Contract.MANUALFULLFILMENTTOKEN(&_Mdeposit.CallOpts)
}

// MANUALFULLFILMENTTOKEN is a free data retrieval call binding the contract method 0x9af40265.
//
// Solidity: function MANUAL_FULLFILMENT_TOKEN() view returns(address)
func (_Mdeposit *MdepositCallerSession) MANUALFULLFILMENTTOKEN() (common.Address, error) {
	return _Mdeposit.Contract.MANUALFULLFILMENTTOKEN(&_Mdeposit.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Mdeposit *MdepositCaller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Mdeposit *MdepositSession) MAXUINT() (*big.Int, error) {
	return _Mdeposit.Contract.MAXUINT(&_Mdeposit.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) MAXUINT() (*big.Int, error) {
	return _Mdeposit.Contract.MAXUINT(&_Mdeposit.CallOpts)
}

// MEDGECUSTOMAGGREGATORFEEDADMINROLE is a free data retrieval call binding the contract method 0x8294344e.
//
// Solidity: function M_EDGE_CUSTOM_AGGREGATOR_FEED_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) MEDGECUSTOMAGGREGATORFEEDADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "M_EDGE_CUSTOM_AGGREGATOR_FEED_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MEDGECUSTOMAGGREGATORFEEDADMINROLE is a free data retrieval call binding the contract method 0x8294344e.
//
// Solidity: function M_EDGE_CUSTOM_AGGREGATOR_FEED_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) MEDGECUSTOMAGGREGATORFEEDADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGECUSTOMAGGREGATORFEEDADMINROLE(&_Mdeposit.CallOpts)
}

// MEDGECUSTOMAGGREGATORFEEDADMINROLE is a free data retrieval call binding the contract method 0x8294344e.
//
// Solidity: function M_EDGE_CUSTOM_AGGREGATOR_FEED_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) MEDGECUSTOMAGGREGATORFEEDADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGECUSTOMAGGREGATORFEEDADMINROLE(&_Mdeposit.CallOpts)
}

// MEDGEDEPOSITVAULTADMINROLE is a free data retrieval call binding the contract method 0xdf3ded45.
//
// Solidity: function M_EDGE_DEPOSIT_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) MEDGEDEPOSITVAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "M_EDGE_DEPOSIT_VAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MEDGEDEPOSITVAULTADMINROLE is a free data retrieval call binding the contract method 0xdf3ded45.
//
// Solidity: function M_EDGE_DEPOSIT_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) MEDGEDEPOSITVAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGEDEPOSITVAULTADMINROLE(&_Mdeposit.CallOpts)
}

// MEDGEDEPOSITVAULTADMINROLE is a free data retrieval call binding the contract method 0xdf3ded45.
//
// Solidity: function M_EDGE_DEPOSIT_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) MEDGEDEPOSITVAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGEDEPOSITVAULTADMINROLE(&_Mdeposit.CallOpts)
}

// MEDGEREDEMPTIONVAULTADMINROLE is a free data retrieval call binding the contract method 0xd9778397.
//
// Solidity: function M_EDGE_REDEMPTION_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCaller) MEDGEREDEMPTIONVAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "M_EDGE_REDEMPTION_VAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MEDGEREDEMPTIONVAULTADMINROLE is a free data retrieval call binding the contract method 0xd9778397.
//
// Solidity: function M_EDGE_REDEMPTION_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositSession) MEDGEREDEMPTIONVAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGEREDEMPTIONVAULTADMINROLE(&_Mdeposit.CallOpts)
}

// MEDGEREDEMPTIONVAULTADMINROLE is a free data retrieval call binding the contract method 0xd9778397.
//
// Solidity: function M_EDGE_REDEMPTION_VAULT_ADMIN_ROLE() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) MEDGEREDEMPTIONVAULTADMINROLE() ([32]byte, error) {
	return _Mdeposit.Contract.MEDGEREDEMPTIONVAULTADMINROLE(&_Mdeposit.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Mdeposit *MdepositCaller) ONEHUNDREDPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "ONE_HUNDRED_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Mdeposit *MdepositSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Mdeposit.Contract.ONEHUNDREDPERCENT(&_Mdeposit.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _Mdeposit.Contract.ONEHUNDREDPERCENT(&_Mdeposit.CallOpts)
}

// STABLECOINRATE is a free data retrieval call binding the contract method 0x978ff560.
//
// Solidity: function STABLECOIN_RATE() view returns(uint256)
func (_Mdeposit *MdepositCaller) STABLECOINRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "STABLECOIN_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STABLECOINRATE is a free data retrieval call binding the contract method 0x978ff560.
//
// Solidity: function STABLECOIN_RATE() view returns(uint256)
func (_Mdeposit *MdepositSession) STABLECOINRATE() (*big.Int, error) {
	return _Mdeposit.Contract.STABLECOINRATE(&_Mdeposit.CallOpts)
}

// STABLECOINRATE is a free data retrieval call binding the contract method 0x978ff560.
//
// Solidity: function STABLECOIN_RATE() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) STABLECOINRATE() (*big.Int, error) {
	return _Mdeposit.Contract.STABLECOINRATE(&_Mdeposit.CallOpts)
}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_Mdeposit *MdepositCaller) AccessControl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "accessControl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_Mdeposit *MdepositSession) AccessControl() (common.Address, error) {
	return _Mdeposit.Contract.AccessControl(&_Mdeposit.CallOpts)
}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_Mdeposit *MdepositCallerSession) AccessControl() (common.Address, error) {
	return _Mdeposit.Contract.AccessControl(&_Mdeposit.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(uint256 _value)
func (_Mdeposit *MdepositCaller) CurrentRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "currentRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(uint256 _value)
func (_Mdeposit *MdepositSession) CurrentRequestId() (*big.Int, error) {
	return _Mdeposit.Contract.CurrentRequestId(&_Mdeposit.CallOpts)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5ae2bfdb.
//
// Solidity: function currentRequestId() view returns(uint256 _value)
func (_Mdeposit *MdepositCallerSession) CurrentRequestId() (*big.Int, error) {
	return _Mdeposit.Contract.CurrentRequestId(&_Mdeposit.CallOpts)
}

// DailyLimits is a free data retrieval call binding the contract method 0x6dc69e03.
//
// Solidity: function dailyLimits(uint256 ) view returns(uint256)
func (_Mdeposit *MdepositCaller) DailyLimits(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "dailyLimits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyLimits is a free data retrieval call binding the contract method 0x6dc69e03.
//
// Solidity: function dailyLimits(uint256 ) view returns(uint256)
func (_Mdeposit *MdepositSession) DailyLimits(arg0 *big.Int) (*big.Int, error) {
	return _Mdeposit.Contract.DailyLimits(&_Mdeposit.CallOpts, arg0)
}

// DailyLimits is a free data retrieval call binding the contract method 0x6dc69e03.
//
// Solidity: function dailyLimits(uint256 ) view returns(uint256)
func (_Mdeposit *MdepositCallerSession) DailyLimits(arg0 *big.Int) (*big.Int, error) {
	return _Mdeposit.Contract.DailyLimits(&_Mdeposit.CallOpts, arg0)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Mdeposit *MdepositCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Mdeposit *MdepositSession) FeeReceiver() (common.Address, error) {
	return _Mdeposit.Contract.FeeReceiver(&_Mdeposit.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Mdeposit *MdepositCallerSession) FeeReceiver() (common.Address, error) {
	return _Mdeposit.Contract.FeeReceiver(&_Mdeposit.CallOpts)
}

// FnPaused is a free data retrieval call binding the contract method 0x0b5a57bd.
//
// Solidity: function fnPaused(bytes4 ) view returns(bool)
func (_Mdeposit *MdepositCaller) FnPaused(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "fnPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FnPaused is a free data retrieval call binding the contract method 0x0b5a57bd.
//
// Solidity: function fnPaused(bytes4 ) view returns(bool)
func (_Mdeposit *MdepositSession) FnPaused(arg0 [4]byte) (bool, error) {
	return _Mdeposit.Contract.FnPaused(&_Mdeposit.CallOpts, arg0)
}

// FnPaused is a free data retrieval call binding the contract method 0x0b5a57bd.
//
// Solidity: function fnPaused(bytes4 ) view returns(bool)
func (_Mdeposit *MdepositCallerSession) FnPaused(arg0 [4]byte) (bool, error) {
	return _Mdeposit.Contract.FnPaused(&_Mdeposit.CallOpts, arg0)
}

// GetPaymentTokens is a free data retrieval call binding the contract method 0xca5e553e.
//
// Solidity: function getPaymentTokens() view returns(address[])
func (_Mdeposit *MdepositCaller) GetPaymentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "getPaymentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPaymentTokens is a free data retrieval call binding the contract method 0xca5e553e.
//
// Solidity: function getPaymentTokens() view returns(address[])
func (_Mdeposit *MdepositSession) GetPaymentTokens() ([]common.Address, error) {
	return _Mdeposit.Contract.GetPaymentTokens(&_Mdeposit.CallOpts)
}

// GetPaymentTokens is a free data retrieval call binding the contract method 0xca5e553e.
//
// Solidity: function getPaymentTokens() view returns(address[])
func (_Mdeposit *MdepositCallerSession) GetPaymentTokens() ([]common.Address, error) {
	return _Mdeposit.Contract.GetPaymentTokens(&_Mdeposit.CallOpts)
}

// GreenlistEnabled is a free data retrieval call binding the contract method 0x105ed2b2.
//
// Solidity: function greenlistEnabled() view returns(bool)
func (_Mdeposit *MdepositCaller) GreenlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "greenlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GreenlistEnabled is a free data retrieval call binding the contract method 0x105ed2b2.
//
// Solidity: function greenlistEnabled() view returns(bool)
func (_Mdeposit *MdepositSession) GreenlistEnabled() (bool, error) {
	return _Mdeposit.Contract.GreenlistEnabled(&_Mdeposit.CallOpts)
}

// GreenlistEnabled is a free data retrieval call binding the contract method 0x105ed2b2.
//
// Solidity: function greenlistEnabled() view returns(bool)
func (_Mdeposit *MdepositCallerSession) GreenlistEnabled() (bool, error) {
	return _Mdeposit.Contract.GreenlistEnabled(&_Mdeposit.CallOpts)
}

// GreenlistTogglerRole is a free data retrieval call binding the contract method 0x32b30cce.
//
// Solidity: function greenlistTogglerRole() view returns(bytes32)
func (_Mdeposit *MdepositCaller) GreenlistTogglerRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "greenlistTogglerRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GreenlistTogglerRole is a free data retrieval call binding the contract method 0x32b30cce.
//
// Solidity: function greenlistTogglerRole() view returns(bytes32)
func (_Mdeposit *MdepositSession) GreenlistTogglerRole() ([32]byte, error) {
	return _Mdeposit.Contract.GreenlistTogglerRole(&_Mdeposit.CallOpts)
}

// GreenlistTogglerRole is a free data retrieval call binding the contract method 0x32b30cce.
//
// Solidity: function greenlistTogglerRole() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) GreenlistTogglerRole() ([32]byte, error) {
	return _Mdeposit.Contract.GreenlistTogglerRole(&_Mdeposit.CallOpts)
}

// GreenlistedRole is a free data retrieval call binding the contract method 0xbbae4086.
//
// Solidity: function greenlistedRole() view returns(bytes32)
func (_Mdeposit *MdepositCaller) GreenlistedRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "greenlistedRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GreenlistedRole is a free data retrieval call binding the contract method 0xbbae4086.
//
// Solidity: function greenlistedRole() view returns(bytes32)
func (_Mdeposit *MdepositSession) GreenlistedRole() ([32]byte, error) {
	return _Mdeposit.Contract.GreenlistedRole(&_Mdeposit.CallOpts)
}

// GreenlistedRole is a free data retrieval call binding the contract method 0xbbae4086.
//
// Solidity: function greenlistedRole() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) GreenlistedRole() ([32]byte, error) {
	return _Mdeposit.Contract.GreenlistedRole(&_Mdeposit.CallOpts)
}

// InstantDailyLimit is a free data retrieval call binding the contract method 0x3972183c.
//
// Solidity: function instantDailyLimit() view returns(uint256)
func (_Mdeposit *MdepositCaller) InstantDailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "instantDailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantDailyLimit is a free data retrieval call binding the contract method 0x3972183c.
//
// Solidity: function instantDailyLimit() view returns(uint256)
func (_Mdeposit *MdepositSession) InstantDailyLimit() (*big.Int, error) {
	return _Mdeposit.Contract.InstantDailyLimit(&_Mdeposit.CallOpts)
}

// InstantDailyLimit is a free data retrieval call binding the contract method 0x3972183c.
//
// Solidity: function instantDailyLimit() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) InstantDailyLimit() (*big.Int, error) {
	return _Mdeposit.Contract.InstantDailyLimit(&_Mdeposit.CallOpts)
}

// InstantFee is a free data retrieval call binding the contract method 0xc47d51be.
//
// Solidity: function instantFee() view returns(uint256)
func (_Mdeposit *MdepositCaller) InstantFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "instantFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InstantFee is a free data retrieval call binding the contract method 0xc47d51be.
//
// Solidity: function instantFee() view returns(uint256)
func (_Mdeposit *MdepositSession) InstantFee() (*big.Int, error) {
	return _Mdeposit.Contract.InstantFee(&_Mdeposit.CallOpts)
}

// InstantFee is a free data retrieval call binding the contract method 0xc47d51be.
//
// Solidity: function instantFee() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) InstantFee() (*big.Int, error) {
	return _Mdeposit.Contract.InstantFee(&_Mdeposit.CallOpts)
}

// IsFreeFromMinAmount is a free data retrieval call binding the contract method 0xd7fd2bae.
//
// Solidity: function isFreeFromMinAmount(address ) view returns(bool)
func (_Mdeposit *MdepositCaller) IsFreeFromMinAmount(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "isFreeFromMinAmount", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFreeFromMinAmount is a free data retrieval call binding the contract method 0xd7fd2bae.
//
// Solidity: function isFreeFromMinAmount(address ) view returns(bool)
func (_Mdeposit *MdepositSession) IsFreeFromMinAmount(arg0 common.Address) (bool, error) {
	return _Mdeposit.Contract.IsFreeFromMinAmount(&_Mdeposit.CallOpts, arg0)
}

// IsFreeFromMinAmount is a free data retrieval call binding the contract method 0xd7fd2bae.
//
// Solidity: function isFreeFromMinAmount(address ) view returns(bool)
func (_Mdeposit *MdepositCallerSession) IsFreeFromMinAmount(arg0 common.Address) (bool, error) {
	return _Mdeposit.Contract.IsFreeFromMinAmount(&_Mdeposit.CallOpts, arg0)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Mdeposit *MdepositCaller) MToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "mToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Mdeposit *MdepositSession) MToken() (common.Address, error) {
	return _Mdeposit.Contract.MToken(&_Mdeposit.CallOpts)
}

// MToken is a free data retrieval call binding the contract method 0xc3b6f939.
//
// Solidity: function mToken() view returns(address)
func (_Mdeposit *MdepositCallerSession) MToken() (common.Address, error) {
	return _Mdeposit.Contract.MToken(&_Mdeposit.CallOpts)
}

// MTokenDataFeed is a free data retrieval call binding the contract method 0x6254afb6.
//
// Solidity: function mTokenDataFeed() view returns(address)
func (_Mdeposit *MdepositCaller) MTokenDataFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "mTokenDataFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MTokenDataFeed is a free data retrieval call binding the contract method 0x6254afb6.
//
// Solidity: function mTokenDataFeed() view returns(address)
func (_Mdeposit *MdepositSession) MTokenDataFeed() (common.Address, error) {
	return _Mdeposit.Contract.MTokenDataFeed(&_Mdeposit.CallOpts)
}

// MTokenDataFeed is a free data retrieval call binding the contract method 0x6254afb6.
//
// Solidity: function mTokenDataFeed() view returns(address)
func (_Mdeposit *MdepositCallerSession) MTokenDataFeed() (common.Address, error) {
	return _Mdeposit.Contract.MTokenDataFeed(&_Mdeposit.CallOpts)
}

// MaxSupplyCap is a free data retrieval call binding the contract method 0x9d70902f.
//
// Solidity: function maxSupplyCap() view returns(uint256)
func (_Mdeposit *MdepositCaller) MaxSupplyCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "maxSupplyCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupplyCap is a free data retrieval call binding the contract method 0x9d70902f.
//
// Solidity: function maxSupplyCap() view returns(uint256)
func (_Mdeposit *MdepositSession) MaxSupplyCap() (*big.Int, error) {
	return _Mdeposit.Contract.MaxSupplyCap(&_Mdeposit.CallOpts)
}

// MaxSupplyCap is a free data retrieval call binding the contract method 0x9d70902f.
//
// Solidity: function maxSupplyCap() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) MaxSupplyCap() (*big.Int, error) {
	return _Mdeposit.Contract.MaxSupplyCap(&_Mdeposit.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Mdeposit *MdepositCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Mdeposit *MdepositSession) MinAmount() (*big.Int, error) {
	return _Mdeposit.Contract.MinAmount(&_Mdeposit.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) MinAmount() (*big.Int, error) {
	return _Mdeposit.Contract.MinAmount(&_Mdeposit.CallOpts)
}

// MinMTokenAmountForFirstDeposit is a free data retrieval call binding the contract method 0xd63567a5.
//
// Solidity: function minMTokenAmountForFirstDeposit() view returns(uint256)
func (_Mdeposit *MdepositCaller) MinMTokenAmountForFirstDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "minMTokenAmountForFirstDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinMTokenAmountForFirstDeposit is a free data retrieval call binding the contract method 0xd63567a5.
//
// Solidity: function minMTokenAmountForFirstDeposit() view returns(uint256)
func (_Mdeposit *MdepositSession) MinMTokenAmountForFirstDeposit() (*big.Int, error) {
	return _Mdeposit.Contract.MinMTokenAmountForFirstDeposit(&_Mdeposit.CallOpts)
}

// MinMTokenAmountForFirstDeposit is a free data retrieval call binding the contract method 0xd63567a5.
//
// Solidity: function minMTokenAmountForFirstDeposit() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) MinMTokenAmountForFirstDeposit() (*big.Int, error) {
	return _Mdeposit.Contract.MinMTokenAmountForFirstDeposit(&_Mdeposit.CallOpts)
}

// MintRequests is a free data retrieval call binding the contract method 0x424e6575.
//
// Solidity: function mintRequests(uint256 ) view returns(address sender, address tokenIn, uint8 status, uint256 depositedUsdAmount, uint256 usdAmountWithoutFees, uint256 tokenOutRate)
func (_Mdeposit *MdepositCaller) MintRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender               common.Address
	TokenIn              common.Address
	Status               uint8
	DepositedUsdAmount   *big.Int
	UsdAmountWithoutFees *big.Int
	TokenOutRate         *big.Int
}, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "mintRequests", arg0)

	outstruct := new(struct {
		Sender               common.Address
		TokenIn              common.Address
		Status               uint8
		DepositedUsdAmount   *big.Int
		UsdAmountWithoutFees *big.Int
		TokenOutRate         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenIn = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.DepositedUsdAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UsdAmountWithoutFees = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TokenOutRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MintRequests is a free data retrieval call binding the contract method 0x424e6575.
//
// Solidity: function mintRequests(uint256 ) view returns(address sender, address tokenIn, uint8 status, uint256 depositedUsdAmount, uint256 usdAmountWithoutFees, uint256 tokenOutRate)
func (_Mdeposit *MdepositSession) MintRequests(arg0 *big.Int) (struct {
	Sender               common.Address
	TokenIn              common.Address
	Status               uint8
	DepositedUsdAmount   *big.Int
	UsdAmountWithoutFees *big.Int
	TokenOutRate         *big.Int
}, error) {
	return _Mdeposit.Contract.MintRequests(&_Mdeposit.CallOpts, arg0)
}

// MintRequests is a free data retrieval call binding the contract method 0x424e6575.
//
// Solidity: function mintRequests(uint256 ) view returns(address sender, address tokenIn, uint8 status, uint256 depositedUsdAmount, uint256 usdAmountWithoutFees, uint256 tokenOutRate)
func (_Mdeposit *MdepositCallerSession) MintRequests(arg0 *big.Int) (struct {
	Sender               common.Address
	TokenIn              common.Address
	Status               uint8
	DepositedUsdAmount   *big.Int
	UsdAmountWithoutFees *big.Int
	TokenOutRate         *big.Int
}, error) {
	return _Mdeposit.Contract.MintRequests(&_Mdeposit.CallOpts, arg0)
}

// PauseAdminRole is a free data retrieval call binding the contract method 0xcabccc7f.
//
// Solidity: function pauseAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositCaller) PauseAdminRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "pauseAdminRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PauseAdminRole is a free data retrieval call binding the contract method 0xcabccc7f.
//
// Solidity: function pauseAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositSession) PauseAdminRole() ([32]byte, error) {
	return _Mdeposit.Contract.PauseAdminRole(&_Mdeposit.CallOpts)
}

// PauseAdminRole is a free data retrieval call binding the contract method 0xcabccc7f.
//
// Solidity: function pauseAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) PauseAdminRole() ([32]byte, error) {
	return _Mdeposit.Contract.PauseAdminRole(&_Mdeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mdeposit *MdepositCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mdeposit *MdepositSession) Paused() (bool, error) {
	return _Mdeposit.Contract.Paused(&_Mdeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mdeposit *MdepositCallerSession) Paused() (bool, error) {
	return _Mdeposit.Contract.Paused(&_Mdeposit.CallOpts)
}

// SanctionsList is a free data retrieval call binding the contract method 0xec571c6a.
//
// Solidity: function sanctionsList() view returns(address)
func (_Mdeposit *MdepositCaller) SanctionsList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "sanctionsList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SanctionsList is a free data retrieval call binding the contract method 0xec571c6a.
//
// Solidity: function sanctionsList() view returns(address)
func (_Mdeposit *MdepositSession) SanctionsList() (common.Address, error) {
	return _Mdeposit.Contract.SanctionsList(&_Mdeposit.CallOpts)
}

// SanctionsList is a free data retrieval call binding the contract method 0xec571c6a.
//
// Solidity: function sanctionsList() view returns(address)
func (_Mdeposit *MdepositCallerSession) SanctionsList() (common.Address, error) {
	return _Mdeposit.Contract.SanctionsList(&_Mdeposit.CallOpts)
}

// SanctionsListAdminRole is a free data retrieval call binding the contract method 0xdaddcb16.
//
// Solidity: function sanctionsListAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositCaller) SanctionsListAdminRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "sanctionsListAdminRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SanctionsListAdminRole is a free data retrieval call binding the contract method 0xdaddcb16.
//
// Solidity: function sanctionsListAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositSession) SanctionsListAdminRole() ([32]byte, error) {
	return _Mdeposit.Contract.SanctionsListAdminRole(&_Mdeposit.CallOpts)
}

// SanctionsListAdminRole is a free data retrieval call binding the contract method 0xdaddcb16.
//
// Solidity: function sanctionsListAdminRole() view returns(bytes32)
func (_Mdeposit *MdepositCallerSession) SanctionsListAdminRole() ([32]byte, error) {
	return _Mdeposit.Contract.SanctionsListAdminRole(&_Mdeposit.CallOpts)
}

// TokensConfig is a free data retrieval call binding the contract method 0xbc979af6.
//
// Solidity: function tokensConfig(address ) view returns(address dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositCaller) TokensConfig(opts *bind.CallOpts, arg0 common.Address) (struct {
	DataFeed  common.Address
	Fee       *big.Int
	Allowance *big.Int
	Stable    bool
}, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "tokensConfig", arg0)

	outstruct := new(struct {
		DataFeed  common.Address
		Fee       *big.Int
		Allowance *big.Int
		Stable    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataFeed = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Allowance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// TokensConfig is a free data retrieval call binding the contract method 0xbc979af6.
//
// Solidity: function tokensConfig(address ) view returns(address dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositSession) TokensConfig(arg0 common.Address) (struct {
	DataFeed  common.Address
	Fee       *big.Int
	Allowance *big.Int
	Stable    bool
}, error) {
	return _Mdeposit.Contract.TokensConfig(&_Mdeposit.CallOpts, arg0)
}

// TokensConfig is a free data retrieval call binding the contract method 0xbc979af6.
//
// Solidity: function tokensConfig(address ) view returns(address dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositCallerSession) TokensConfig(arg0 common.Address) (struct {
	DataFeed  common.Address
	Fee       *big.Int
	Allowance *big.Int
	Stable    bool
}, error) {
	return _Mdeposit.Contract.TokensConfig(&_Mdeposit.CallOpts, arg0)
}

// TokensReceiver is a free data retrieval call binding the contract method 0x1fa1e8d4.
//
// Solidity: function tokensReceiver() view returns(address)
func (_Mdeposit *MdepositCaller) TokensReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "tokensReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokensReceiver is a free data retrieval call binding the contract method 0x1fa1e8d4.
//
// Solidity: function tokensReceiver() view returns(address)
func (_Mdeposit *MdepositSession) TokensReceiver() (common.Address, error) {
	return _Mdeposit.Contract.TokensReceiver(&_Mdeposit.CallOpts)
}

// TokensReceiver is a free data retrieval call binding the contract method 0x1fa1e8d4.
//
// Solidity: function tokensReceiver() view returns(address)
func (_Mdeposit *MdepositCallerSession) TokensReceiver() (common.Address, error) {
	return _Mdeposit.Contract.TokensReceiver(&_Mdeposit.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_Mdeposit *MdepositCaller) TotalMinted(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "totalMinted", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_Mdeposit *MdepositSession) TotalMinted(arg0 common.Address) (*big.Int, error) {
	return _Mdeposit.Contract.TotalMinted(&_Mdeposit.CallOpts, arg0)
}

// TotalMinted is a free data retrieval call binding the contract method 0x003d4790.
//
// Solidity: function totalMinted(address ) view returns(uint256)
func (_Mdeposit *MdepositCallerSession) TotalMinted(arg0 common.Address) (*big.Int, error) {
	return _Mdeposit.Contract.TotalMinted(&_Mdeposit.CallOpts, arg0)
}

// VariationTolerance is a free data retrieval call binding the contract method 0x7192de4b.
//
// Solidity: function variationTolerance() view returns(uint256)
func (_Mdeposit *MdepositCaller) VariationTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "variationTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VariationTolerance is a free data retrieval call binding the contract method 0x7192de4b.
//
// Solidity: function variationTolerance() view returns(uint256)
func (_Mdeposit *MdepositSession) VariationTolerance() (*big.Int, error) {
	return _Mdeposit.Contract.VariationTolerance(&_Mdeposit.CallOpts)
}

// VariationTolerance is a free data retrieval call binding the contract method 0x7192de4b.
//
// Solidity: function variationTolerance() view returns(uint256)
func (_Mdeposit *MdepositCallerSession) VariationTolerance() (*big.Int, error) {
	return _Mdeposit.Contract.VariationTolerance(&_Mdeposit.CallOpts)
}

// VaultRole is a free data retrieval call binding the contract method 0xeaf896fd.
//
// Solidity: function vaultRole() pure returns(bytes32)
func (_Mdeposit *MdepositCaller) VaultRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "vaultRole")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VaultRole is a free data retrieval call binding the contract method 0xeaf896fd.
//
// Solidity: function vaultRole() pure returns(bytes32)
func (_Mdeposit *MdepositSession) VaultRole() ([32]byte, error) {
	return _Mdeposit.Contract.VaultRole(&_Mdeposit.CallOpts)
}

// VaultRole is a free data retrieval call binding the contract method 0xeaf896fd.
//
// Solidity: function vaultRole() pure returns(bytes32)
func (_Mdeposit *MdepositCallerSession) VaultRole() ([32]byte, error) {
	return _Mdeposit.Contract.VaultRole(&_Mdeposit.CallOpts)
}

// WaivedFeeRestriction is a free data retrieval call binding the contract method 0x042da5ee.
//
// Solidity: function waivedFeeRestriction(address ) view returns(bool)
func (_Mdeposit *MdepositCaller) WaivedFeeRestriction(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Mdeposit.contract.Call(opts, &out, "waivedFeeRestriction", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WaivedFeeRestriction is a free data retrieval call binding the contract method 0x042da5ee.
//
// Solidity: function waivedFeeRestriction(address ) view returns(bool)
func (_Mdeposit *MdepositSession) WaivedFeeRestriction(arg0 common.Address) (bool, error) {
	return _Mdeposit.Contract.WaivedFeeRestriction(&_Mdeposit.CallOpts, arg0)
}

// WaivedFeeRestriction is a free data retrieval call binding the contract method 0x042da5ee.
//
// Solidity: function waivedFeeRestriction(address ) view returns(bool)
func (_Mdeposit *MdepositCallerSession) WaivedFeeRestriction(arg0 common.Address) (bool, error) {
	return _Mdeposit.Contract.WaivedFeeRestriction(&_Mdeposit.CallOpts, arg0)
}

// AddPaymentToken is a paid mutator transaction binding the contract method 0xf41759e7.
//
// Solidity: function addPaymentToken(address token, address dataFeed, uint256 tokenFee, uint256 allowance, bool stable) returns()
func (_Mdeposit *MdepositTransactor) AddPaymentToken(opts *bind.TransactOpts, token common.Address, dataFeed common.Address, tokenFee *big.Int, allowance *big.Int, stable bool) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "addPaymentToken", token, dataFeed, tokenFee, allowance, stable)
}

// AddPaymentToken is a paid mutator transaction binding the contract method 0xf41759e7.
//
// Solidity: function addPaymentToken(address token, address dataFeed, uint256 tokenFee, uint256 allowance, bool stable) returns()
func (_Mdeposit *MdepositSession) AddPaymentToken(token common.Address, dataFeed common.Address, tokenFee *big.Int, allowance *big.Int, stable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.AddPaymentToken(&_Mdeposit.TransactOpts, token, dataFeed, tokenFee, allowance, stable)
}

// AddPaymentToken is a paid mutator transaction binding the contract method 0xf41759e7.
//
// Solidity: function addPaymentToken(address token, address dataFeed, uint256 tokenFee, uint256 allowance, bool stable) returns()
func (_Mdeposit *MdepositTransactorSession) AddPaymentToken(token common.Address, dataFeed common.Address, tokenFee *big.Int, allowance *big.Int, stable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.AddPaymentToken(&_Mdeposit.TransactOpts, token, dataFeed, tokenFee, allowance, stable)
}

// AddWaivedFeeAccount is a paid mutator transaction binding the contract method 0xe428877e.
//
// Solidity: function addWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositTransactor) AddWaivedFeeAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "addWaivedFeeAccount", account)
}

// AddWaivedFeeAccount is a paid mutator transaction binding the contract method 0xe428877e.
//
// Solidity: function addWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositSession) AddWaivedFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.AddWaivedFeeAccount(&_Mdeposit.TransactOpts, account)
}

// AddWaivedFeeAccount is a paid mutator transaction binding the contract method 0xe428877e.
//
// Solidity: function addWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositTransactorSession) AddWaivedFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.AddWaivedFeeAccount(&_Mdeposit.TransactOpts, account)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x2c0a90a9.
//
// Solidity: function approveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactor) ApproveRequest(opts *bind.TransactOpts, requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "approveRequest", requestId, newOutRate)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x2c0a90a9.
//
// Solidity: function approveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositSession) ApproveRequest(requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ApproveRequest(&_Mdeposit.TransactOpts, requestId, newOutRate)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0x2c0a90a9.
//
// Solidity: function approveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactorSession) ApproveRequest(requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ApproveRequest(&_Mdeposit.TransactOpts, requestId, newOutRate)
}

// ChangeTokenAllowance is a paid mutator transaction binding the contract method 0x8a0ae615.
//
// Solidity: function changeTokenAllowance(address token, uint256 allowance) returns()
func (_Mdeposit *MdepositTransactor) ChangeTokenAllowance(opts *bind.TransactOpts, token common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "changeTokenAllowance", token, allowance)
}

// ChangeTokenAllowance is a paid mutator transaction binding the contract method 0x8a0ae615.
//
// Solidity: function changeTokenAllowance(address token, uint256 allowance) returns()
func (_Mdeposit *MdepositSession) ChangeTokenAllowance(token common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ChangeTokenAllowance(&_Mdeposit.TransactOpts, token, allowance)
}

// ChangeTokenAllowance is a paid mutator transaction binding the contract method 0x8a0ae615.
//
// Solidity: function changeTokenAllowance(address token, uint256 allowance) returns()
func (_Mdeposit *MdepositTransactorSession) ChangeTokenAllowance(token common.Address, allowance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ChangeTokenAllowance(&_Mdeposit.TransactOpts, token, allowance)
}

// ChangeTokenFee is a paid mutator transaction binding the contract method 0xdb74d8b5.
//
// Solidity: function changeTokenFee(address token, uint256 fee) returns()
func (_Mdeposit *MdepositTransactor) ChangeTokenFee(opts *bind.TransactOpts, token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "changeTokenFee", token, fee)
}

// ChangeTokenFee is a paid mutator transaction binding the contract method 0xdb74d8b5.
//
// Solidity: function changeTokenFee(address token, uint256 fee) returns()
func (_Mdeposit *MdepositSession) ChangeTokenFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ChangeTokenFee(&_Mdeposit.TransactOpts, token, fee)
}

// ChangeTokenFee is a paid mutator transaction binding the contract method 0xdb74d8b5.
//
// Solidity: function changeTokenFee(address token, uint256 fee) returns()
func (_Mdeposit *MdepositTransactorSession) ChangeTokenFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.ChangeTokenFee(&_Mdeposit.TransactOpts, token, fee)
}

// DepositInstant is a paid mutator transaction binding the contract method 0x42e8866b.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId, address recipient) returns()
func (_Mdeposit *MdepositTransactor) DepositInstant(opts *bind.TransactOpts, tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "depositInstant", tokenIn, amountToken, minReceiveAmount, referrerId, recipient)
}

// DepositInstant is a paid mutator transaction binding the contract method 0x42e8866b.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId, address recipient) returns()
func (_Mdeposit *MdepositSession) DepositInstant(tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositInstant(&_Mdeposit.TransactOpts, tokenIn, amountToken, minReceiveAmount, referrerId, recipient)
}

// DepositInstant is a paid mutator transaction binding the contract method 0x42e8866b.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId, address recipient) returns()
func (_Mdeposit *MdepositTransactorSession) DepositInstant(tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositInstant(&_Mdeposit.TransactOpts, tokenIn, amountToken, minReceiveAmount, referrerId, recipient)
}

// DepositInstant0 is a paid mutator transaction binding the contract method 0xc02dd27a.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId) returns()
func (_Mdeposit *MdepositTransactor) DepositInstant0(opts *bind.TransactOpts, tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "depositInstant0", tokenIn, amountToken, minReceiveAmount, referrerId)
}

// DepositInstant0 is a paid mutator transaction binding the contract method 0xc02dd27a.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId) returns()
func (_Mdeposit *MdepositSession) DepositInstant0(tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositInstant0(&_Mdeposit.TransactOpts, tokenIn, amountToken, minReceiveAmount, referrerId)
}

// DepositInstant0 is a paid mutator transaction binding the contract method 0xc02dd27a.
//
// Solidity: function depositInstant(address tokenIn, uint256 amountToken, uint256 minReceiveAmount, bytes32 referrerId) returns()
func (_Mdeposit *MdepositTransactorSession) DepositInstant0(tokenIn common.Address, amountToken *big.Int, minReceiveAmount *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositInstant0(&_Mdeposit.TransactOpts, tokenIn, amountToken, minReceiveAmount, referrerId)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x6e26b9f8.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId) returns(uint256)
func (_Mdeposit *MdepositTransactor) DepositRequest(opts *bind.TransactOpts, tokenIn common.Address, amountToken *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "depositRequest", tokenIn, amountToken, referrerId)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x6e26b9f8.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId) returns(uint256)
func (_Mdeposit *MdepositSession) DepositRequest(tokenIn common.Address, amountToken *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositRequest(&_Mdeposit.TransactOpts, tokenIn, amountToken, referrerId)
}

// DepositRequest is a paid mutator transaction binding the contract method 0x6e26b9f8.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId) returns(uint256)
func (_Mdeposit *MdepositTransactorSession) DepositRequest(tokenIn common.Address, amountToken *big.Int, referrerId [32]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositRequest(&_Mdeposit.TransactOpts, tokenIn, amountToken, referrerId)
}

// DepositRequest0 is a paid mutator transaction binding the contract method 0xe50e3dbb.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId, address recipient) returns(uint256)
func (_Mdeposit *MdepositTransactor) DepositRequest0(opts *bind.TransactOpts, tokenIn common.Address, amountToken *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "depositRequest0", tokenIn, amountToken, referrerId, recipient)
}

// DepositRequest0 is a paid mutator transaction binding the contract method 0xe50e3dbb.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId, address recipient) returns(uint256)
func (_Mdeposit *MdepositSession) DepositRequest0(tokenIn common.Address, amountToken *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositRequest0(&_Mdeposit.TransactOpts, tokenIn, amountToken, referrerId, recipient)
}

// DepositRequest0 is a paid mutator transaction binding the contract method 0xe50e3dbb.
//
// Solidity: function depositRequest(address tokenIn, uint256 amountToken, bytes32 referrerId, address recipient) returns(uint256)
func (_Mdeposit *MdepositTransactorSession) DepositRequest0(tokenIn common.Address, amountToken *big.Int, referrerId [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.DepositRequest0(&_Mdeposit.TransactOpts, tokenIn, amountToken, referrerId, recipient)
}

// FreeFromMinAmount is a paid mutator transaction binding the contract method 0x39dac34d.
//
// Solidity: function freeFromMinAmount(address user, bool enable) returns()
func (_Mdeposit *MdepositTransactor) FreeFromMinAmount(opts *bind.TransactOpts, user common.Address, enable bool) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "freeFromMinAmount", user, enable)
}

// FreeFromMinAmount is a paid mutator transaction binding the contract method 0x39dac34d.
//
// Solidity: function freeFromMinAmount(address user, bool enable) returns()
func (_Mdeposit *MdepositSession) FreeFromMinAmount(user common.Address, enable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.FreeFromMinAmount(&_Mdeposit.TransactOpts, user, enable)
}

// FreeFromMinAmount is a paid mutator transaction binding the contract method 0x39dac34d.
//
// Solidity: function freeFromMinAmount(address user, bool enable) returns()
func (_Mdeposit *MdepositTransactorSession) FreeFromMinAmount(user common.Address, enable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.FreeFromMinAmount(&_Mdeposit.TransactOpts, user, enable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc14c07d6.
//
// Solidity: function initialize(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit, uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositTransactor) Initialize(opts *bind.TransactOpts, _ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int, _maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "initialize", _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit, _maxSupplyCap)
}

// Initialize is a paid mutator transaction binding the contract method 0xc14c07d6.
//
// Solidity: function initialize(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit, uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositSession) Initialize(_ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int, _maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.Initialize(&_Mdeposit.TransactOpts, _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit, _maxSupplyCap)
}

// Initialize is a paid mutator transaction binding the contract method 0xc14c07d6.
//
// Solidity: function initialize(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit, uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositTransactorSession) Initialize(_ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int, _maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.Initialize(&_Mdeposit.TransactOpts, _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit, _maxSupplyCap)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x96f5c6e1.
//
// Solidity: function initializeV1(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit) returns()
func (_Mdeposit *MdepositTransactor) InitializeV1(opts *bind.TransactOpts, _ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "initializeV1", _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x96f5c6e1.
//
// Solidity: function initializeV1(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit) returns()
func (_Mdeposit *MdepositSession) InitializeV1(_ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.InitializeV1(&_Mdeposit.TransactOpts, _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x96f5c6e1.
//
// Solidity: function initializeV1(address _ac, (address,address) _mTokenInitParams, (address,address) _receiversInitParams, (uint256,uint256) _instantInitParams, address _sanctionsList, uint256 _variationTolerance, uint256 _minAmount, uint256 _minMTokenAmountForFirstDeposit) returns()
func (_Mdeposit *MdepositTransactorSession) InitializeV1(_ac common.Address, _mTokenInitParams MTokenInitParams, _receiversInitParams ReceiversInitParams, _instantInitParams InstantInitParams, _sanctionsList common.Address, _variationTolerance *big.Int, _minAmount *big.Int, _minMTokenAmountForFirstDeposit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.InitializeV1(&_Mdeposit.TransactOpts, _ac, _mTokenInitParams, _receiversInitParams, _instantInitParams, _sanctionsList, _variationTolerance, _minAmount, _minMTokenAmountForFirstDeposit)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositTransactor) InitializeV2(opts *bind.TransactOpts, _maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "initializeV2", _maxSupplyCap)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositSession) InitializeV2(_maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.InitializeV2(&_Mdeposit.TransactOpts, _maxSupplyCap)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5b4e128c.
//
// Solidity: function initializeV2(uint256 _maxSupplyCap) returns()
func (_Mdeposit *MdepositTransactorSession) InitializeV2(_maxSupplyCap *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.InitializeV2(&_Mdeposit.TransactOpts, _maxSupplyCap)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mdeposit *MdepositTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mdeposit *MdepositSession) Pause() (*types.Transaction, error) {
	return _Mdeposit.Contract.Pause(&_Mdeposit.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mdeposit *MdepositTransactorSession) Pause() (*types.Transaction, error) {
	return _Mdeposit.Contract.Pause(&_Mdeposit.TransactOpts)
}

// PauseFn is a paid mutator transaction binding the contract method 0x3733337d.
//
// Solidity: function pauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositTransactor) PauseFn(opts *bind.TransactOpts, fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "pauseFn", fn)
}

// PauseFn is a paid mutator transaction binding the contract method 0x3733337d.
//
// Solidity: function pauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositSession) PauseFn(fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.PauseFn(&_Mdeposit.TransactOpts, fn)
}

// PauseFn is a paid mutator transaction binding the contract method 0x3733337d.
//
// Solidity: function pauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositTransactorSession) PauseFn(fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.PauseFn(&_Mdeposit.TransactOpts, fn)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId) returns()
func (_Mdeposit *MdepositTransactor) RejectRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "rejectRequest", requestId)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId) returns()
func (_Mdeposit *MdepositSession) RejectRequest(requestId *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.RejectRequest(&_Mdeposit.TransactOpts, requestId)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId) returns()
func (_Mdeposit *MdepositTransactorSession) RejectRequest(requestId *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.RejectRequest(&_Mdeposit.TransactOpts, requestId)
}

// RemovePaymentToken is a paid mutator transaction binding the contract method 0xa5125421.
//
// Solidity: function removePaymentToken(address token) returns()
func (_Mdeposit *MdepositTransactor) RemovePaymentToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "removePaymentToken", token)
}

// RemovePaymentToken is a paid mutator transaction binding the contract method 0xa5125421.
//
// Solidity: function removePaymentToken(address token) returns()
func (_Mdeposit *MdepositSession) RemovePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.RemovePaymentToken(&_Mdeposit.TransactOpts, token)
}

// RemovePaymentToken is a paid mutator transaction binding the contract method 0xa5125421.
//
// Solidity: function removePaymentToken(address token) returns()
func (_Mdeposit *MdepositTransactorSession) RemovePaymentToken(token common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.RemovePaymentToken(&_Mdeposit.TransactOpts, token)
}

// RemoveWaivedFeeAccount is a paid mutator transaction binding the contract method 0x16683aa5.
//
// Solidity: function removeWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositTransactor) RemoveWaivedFeeAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "removeWaivedFeeAccount", account)
}

// RemoveWaivedFeeAccount is a paid mutator transaction binding the contract method 0x16683aa5.
//
// Solidity: function removeWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositSession) RemoveWaivedFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.RemoveWaivedFeeAccount(&_Mdeposit.TransactOpts, account)
}

// RemoveWaivedFeeAccount is a paid mutator transaction binding the contract method 0x16683aa5.
//
// Solidity: function removeWaivedFeeAccount(address account) returns()
func (_Mdeposit *MdepositTransactorSession) RemoveWaivedFeeAccount(account common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.RemoveWaivedFeeAccount(&_Mdeposit.TransactOpts, account)
}

// SafeApproveRequest is a paid mutator transaction binding the contract method 0x88a6de68.
//
// Solidity: function safeApproveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactor) SafeApproveRequest(opts *bind.TransactOpts, requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "safeApproveRequest", requestId, newOutRate)
}

// SafeApproveRequest is a paid mutator transaction binding the contract method 0x88a6de68.
//
// Solidity: function safeApproveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositSession) SafeApproveRequest(requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeApproveRequest(&_Mdeposit.TransactOpts, requestId, newOutRate)
}

// SafeApproveRequest is a paid mutator transaction binding the contract method 0x88a6de68.
//
// Solidity: function safeApproveRequest(uint256 requestId, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactorSession) SafeApproveRequest(requestId *big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeApproveRequest(&_Mdeposit.TransactOpts, requestId, newOutRate)
}

// SafeBulkApproveRequest is a paid mutator transaction binding the contract method 0xa0c74afc.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds) returns()
func (_Mdeposit *MdepositTransactor) SafeBulkApproveRequest(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "safeBulkApproveRequest", requestIds)
}

// SafeBulkApproveRequest is a paid mutator transaction binding the contract method 0xa0c74afc.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds) returns()
func (_Mdeposit *MdepositSession) SafeBulkApproveRequest(requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequest(&_Mdeposit.TransactOpts, requestIds)
}

// SafeBulkApproveRequest is a paid mutator transaction binding the contract method 0xa0c74afc.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds) returns()
func (_Mdeposit *MdepositTransactorSession) SafeBulkApproveRequest(requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequest(&_Mdeposit.TransactOpts, requestIds)
}

// SafeBulkApproveRequest0 is a paid mutator transaction binding the contract method 0xf5d46c51.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactor) SafeBulkApproveRequest0(opts *bind.TransactOpts, requestIds []*big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "safeBulkApproveRequest0", requestIds, newOutRate)
}

// SafeBulkApproveRequest0 is a paid mutator transaction binding the contract method 0xf5d46c51.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds, uint256 newOutRate) returns()
func (_Mdeposit *MdepositSession) SafeBulkApproveRequest0(requestIds []*big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequest0(&_Mdeposit.TransactOpts, requestIds, newOutRate)
}

// SafeBulkApproveRequest0 is a paid mutator transaction binding the contract method 0xf5d46c51.
//
// Solidity: function safeBulkApproveRequest(uint256[] requestIds, uint256 newOutRate) returns()
func (_Mdeposit *MdepositTransactorSession) SafeBulkApproveRequest0(requestIds []*big.Int, newOutRate *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequest0(&_Mdeposit.TransactOpts, requestIds, newOutRate)
}

// SafeBulkApproveRequestAtSavedRate is a paid mutator transaction binding the contract method 0xdb6c68e3.
//
// Solidity: function safeBulkApproveRequestAtSavedRate(uint256[] requestIds) returns()
func (_Mdeposit *MdepositTransactor) SafeBulkApproveRequestAtSavedRate(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "safeBulkApproveRequestAtSavedRate", requestIds)
}

// SafeBulkApproveRequestAtSavedRate is a paid mutator transaction binding the contract method 0xdb6c68e3.
//
// Solidity: function safeBulkApproveRequestAtSavedRate(uint256[] requestIds) returns()
func (_Mdeposit *MdepositSession) SafeBulkApproveRequestAtSavedRate(requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequestAtSavedRate(&_Mdeposit.TransactOpts, requestIds)
}

// SafeBulkApproveRequestAtSavedRate is a paid mutator transaction binding the contract method 0xdb6c68e3.
//
// Solidity: function safeBulkApproveRequestAtSavedRate(uint256[] requestIds) returns()
func (_Mdeposit *MdepositTransactorSession) SafeBulkApproveRequestAtSavedRate(requestIds []*big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SafeBulkApproveRequestAtSavedRate(&_Mdeposit.TransactOpts, requestIds)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_Mdeposit *MdepositTransactor) SetFeeReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setFeeReceiver", receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_Mdeposit *MdepositSession) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetFeeReceiver(&_Mdeposit.TransactOpts, receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_Mdeposit *MdepositTransactorSession) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetFeeReceiver(&_Mdeposit.TransactOpts, receiver)
}

// SetGreenlistEnable is a paid mutator transaction binding the contract method 0x27abf518.
//
// Solidity: function setGreenlistEnable(bool enable) returns()
func (_Mdeposit *MdepositTransactor) SetGreenlistEnable(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setGreenlistEnable", enable)
}

// SetGreenlistEnable is a paid mutator transaction binding the contract method 0x27abf518.
//
// Solidity: function setGreenlistEnable(bool enable) returns()
func (_Mdeposit *MdepositSession) SetGreenlistEnable(enable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetGreenlistEnable(&_Mdeposit.TransactOpts, enable)
}

// SetGreenlistEnable is a paid mutator transaction binding the contract method 0x27abf518.
//
// Solidity: function setGreenlistEnable(bool enable) returns()
func (_Mdeposit *MdepositTransactorSession) SetGreenlistEnable(enable bool) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetGreenlistEnable(&_Mdeposit.TransactOpts, enable)
}

// SetInstantDailyLimit is a paid mutator transaction binding the contract method 0x6957463a.
//
// Solidity: function setInstantDailyLimit(uint256 newInstantDailyLimit) returns()
func (_Mdeposit *MdepositTransactor) SetInstantDailyLimit(opts *bind.TransactOpts, newInstantDailyLimit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setInstantDailyLimit", newInstantDailyLimit)
}

// SetInstantDailyLimit is a paid mutator transaction binding the contract method 0x6957463a.
//
// Solidity: function setInstantDailyLimit(uint256 newInstantDailyLimit) returns()
func (_Mdeposit *MdepositSession) SetInstantDailyLimit(newInstantDailyLimit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetInstantDailyLimit(&_Mdeposit.TransactOpts, newInstantDailyLimit)
}

// SetInstantDailyLimit is a paid mutator transaction binding the contract method 0x6957463a.
//
// Solidity: function setInstantDailyLimit(uint256 newInstantDailyLimit) returns()
func (_Mdeposit *MdepositTransactorSession) SetInstantDailyLimit(newInstantDailyLimit *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetInstantDailyLimit(&_Mdeposit.TransactOpts, newInstantDailyLimit)
}

// SetInstantFee is a paid mutator transaction binding the contract method 0xad9e5649.
//
// Solidity: function setInstantFee(uint256 newInstantFee) returns()
func (_Mdeposit *MdepositTransactor) SetInstantFee(opts *bind.TransactOpts, newInstantFee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setInstantFee", newInstantFee)
}

// SetInstantFee is a paid mutator transaction binding the contract method 0xad9e5649.
//
// Solidity: function setInstantFee(uint256 newInstantFee) returns()
func (_Mdeposit *MdepositSession) SetInstantFee(newInstantFee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetInstantFee(&_Mdeposit.TransactOpts, newInstantFee)
}

// SetInstantFee is a paid mutator transaction binding the contract method 0xad9e5649.
//
// Solidity: function setInstantFee(uint256 newInstantFee) returns()
func (_Mdeposit *MdepositTransactorSession) SetInstantFee(newInstantFee *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetInstantFee(&_Mdeposit.TransactOpts, newInstantFee)
}

// SetMaxSupplyCap is a paid mutator transaction binding the contract method 0xc626b980.
//
// Solidity: function setMaxSupplyCap(uint256 newValue) returns()
func (_Mdeposit *MdepositTransactor) SetMaxSupplyCap(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setMaxSupplyCap", newValue)
}

// SetMaxSupplyCap is a paid mutator transaction binding the contract method 0xc626b980.
//
// Solidity: function setMaxSupplyCap(uint256 newValue) returns()
func (_Mdeposit *MdepositSession) SetMaxSupplyCap(newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMaxSupplyCap(&_Mdeposit.TransactOpts, newValue)
}

// SetMaxSupplyCap is a paid mutator transaction binding the contract method 0xc626b980.
//
// Solidity: function setMaxSupplyCap(uint256 newValue) returns()
func (_Mdeposit *MdepositTransactorSession) SetMaxSupplyCap(newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMaxSupplyCap(&_Mdeposit.TransactOpts, newValue)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 newAmount) returns()
func (_Mdeposit *MdepositTransactor) SetMinAmount(opts *bind.TransactOpts, newAmount *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setMinAmount", newAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 newAmount) returns()
func (_Mdeposit *MdepositSession) SetMinAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMinAmount(&_Mdeposit.TransactOpts, newAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 newAmount) returns()
func (_Mdeposit *MdepositTransactorSession) SetMinAmount(newAmount *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMinAmount(&_Mdeposit.TransactOpts, newAmount)
}

// SetMinMTokenAmountForFirstDeposit is a paid mutator transaction binding the contract method 0x1e022f4c.
//
// Solidity: function setMinMTokenAmountForFirstDeposit(uint256 newValue) returns()
func (_Mdeposit *MdepositTransactor) SetMinMTokenAmountForFirstDeposit(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setMinMTokenAmountForFirstDeposit", newValue)
}

// SetMinMTokenAmountForFirstDeposit is a paid mutator transaction binding the contract method 0x1e022f4c.
//
// Solidity: function setMinMTokenAmountForFirstDeposit(uint256 newValue) returns()
func (_Mdeposit *MdepositSession) SetMinMTokenAmountForFirstDeposit(newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMinMTokenAmountForFirstDeposit(&_Mdeposit.TransactOpts, newValue)
}

// SetMinMTokenAmountForFirstDeposit is a paid mutator transaction binding the contract method 0x1e022f4c.
//
// Solidity: function setMinMTokenAmountForFirstDeposit(uint256 newValue) returns()
func (_Mdeposit *MdepositTransactorSession) SetMinMTokenAmountForFirstDeposit(newValue *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetMinMTokenAmountForFirstDeposit(&_Mdeposit.TransactOpts, newValue)
}

// SetSanctionsList is a paid mutator transaction binding the contract method 0x49dc5e8d.
//
// Solidity: function setSanctionsList(address newSanctionsList) returns()
func (_Mdeposit *MdepositTransactor) SetSanctionsList(opts *bind.TransactOpts, newSanctionsList common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setSanctionsList", newSanctionsList)
}

// SetSanctionsList is a paid mutator transaction binding the contract method 0x49dc5e8d.
//
// Solidity: function setSanctionsList(address newSanctionsList) returns()
func (_Mdeposit *MdepositSession) SetSanctionsList(newSanctionsList common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetSanctionsList(&_Mdeposit.TransactOpts, newSanctionsList)
}

// SetSanctionsList is a paid mutator transaction binding the contract method 0x49dc5e8d.
//
// Solidity: function setSanctionsList(address newSanctionsList) returns()
func (_Mdeposit *MdepositTransactorSession) SetSanctionsList(newSanctionsList common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetSanctionsList(&_Mdeposit.TransactOpts, newSanctionsList)
}

// SetTokensReceiver is a paid mutator transaction binding the contract method 0x476abc76.
//
// Solidity: function setTokensReceiver(address receiver) returns()
func (_Mdeposit *MdepositTransactor) SetTokensReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setTokensReceiver", receiver)
}

// SetTokensReceiver is a paid mutator transaction binding the contract method 0x476abc76.
//
// Solidity: function setTokensReceiver(address receiver) returns()
func (_Mdeposit *MdepositSession) SetTokensReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetTokensReceiver(&_Mdeposit.TransactOpts, receiver)
}

// SetTokensReceiver is a paid mutator transaction binding the contract method 0x476abc76.
//
// Solidity: function setTokensReceiver(address receiver) returns()
func (_Mdeposit *MdepositTransactorSession) SetTokensReceiver(receiver common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetTokensReceiver(&_Mdeposit.TransactOpts, receiver)
}

// SetVariationTolerance is a paid mutator transaction binding the contract method 0x34c24489.
//
// Solidity: function setVariationTolerance(uint256 tolerance) returns()
func (_Mdeposit *MdepositTransactor) SetVariationTolerance(opts *bind.TransactOpts, tolerance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "setVariationTolerance", tolerance)
}

// SetVariationTolerance is a paid mutator transaction binding the contract method 0x34c24489.
//
// Solidity: function setVariationTolerance(uint256 tolerance) returns()
func (_Mdeposit *MdepositSession) SetVariationTolerance(tolerance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetVariationTolerance(&_Mdeposit.TransactOpts, tolerance)
}

// SetVariationTolerance is a paid mutator transaction binding the contract method 0x34c24489.
//
// Solidity: function setVariationTolerance(uint256 tolerance) returns()
func (_Mdeposit *MdepositTransactorSession) SetVariationTolerance(tolerance *big.Int) (*types.Transaction, error) {
	return _Mdeposit.Contract.SetVariationTolerance(&_Mdeposit.TransactOpts, tolerance)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mdeposit *MdepositTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mdeposit *MdepositSession) Unpause() (*types.Transaction, error) {
	return _Mdeposit.Contract.Unpause(&_Mdeposit.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mdeposit *MdepositTransactorSession) Unpause() (*types.Transaction, error) {
	return _Mdeposit.Contract.Unpause(&_Mdeposit.TransactOpts)
}

// UnpauseFn is a paid mutator transaction binding the contract method 0x3807be7d.
//
// Solidity: function unpauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositTransactor) UnpauseFn(opts *bind.TransactOpts, fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "unpauseFn", fn)
}

// UnpauseFn is a paid mutator transaction binding the contract method 0x3807be7d.
//
// Solidity: function unpauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositSession) UnpauseFn(fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.UnpauseFn(&_Mdeposit.TransactOpts, fn)
}

// UnpauseFn is a paid mutator transaction binding the contract method 0x3807be7d.
//
// Solidity: function unpauseFn(bytes4 fn) returns()
func (_Mdeposit *MdepositTransactorSession) UnpauseFn(fn [4]byte) (*types.Transaction, error) {
	return _Mdeposit.Contract.UnpauseFn(&_Mdeposit.TransactOpts, fn)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address withdrawTo) returns()
func (_Mdeposit *MdepositTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Mdeposit.contract.Transact(opts, "withdrawToken", token, amount, withdrawTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address withdrawTo) returns()
func (_Mdeposit *MdepositSession) WithdrawToken(token common.Address, amount *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.WithdrawToken(&_Mdeposit.TransactOpts, token, amount, withdrawTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(address token, uint256 amount, address withdrawTo) returns()
func (_Mdeposit *MdepositTransactorSession) WithdrawToken(token common.Address, amount *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Mdeposit.Contract.WithdrawToken(&_Mdeposit.TransactOpts, token, amount, withdrawTo)
}

// MdepositAddPaymentTokenIterator is returned from FilterAddPaymentToken and is used to iterate over the raw logs and unpacked data for AddPaymentToken events raised by the Mdeposit contract.
type MdepositAddPaymentTokenIterator struct {
	Event *MdepositAddPaymentToken // Event containing the contract specifics and raw log

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
func (it *MdepositAddPaymentTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositAddPaymentToken)
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
		it.Event = new(MdepositAddPaymentToken)
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
func (it *MdepositAddPaymentTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositAddPaymentTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositAddPaymentToken represents a AddPaymentToken event raised by the Mdeposit contract.
type MdepositAddPaymentToken struct {
	Caller    common.Address
	Token     common.Address
	DataFeed  common.Address
	Fee       *big.Int
	Allowance *big.Int
	Stable    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddPaymentToken is a free log retrieval operation binding the contract event 0x049000a9db89588d7bfb162bc0f7e4299ee8762430a468131c2caf0824f1f995.
//
// Solidity: event AddPaymentToken(address indexed caller, address indexed token, address indexed dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositFilterer) FilterAddPaymentToken(opts *bind.FilterOpts, caller []common.Address, token []common.Address, dataFeed []common.Address) (*MdepositAddPaymentTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var dataFeedRule []interface{}
	for _, dataFeedItem := range dataFeed {
		dataFeedRule = append(dataFeedRule, dataFeedItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "AddPaymentToken", callerRule, tokenRule, dataFeedRule)
	if err != nil {
		return nil, err
	}
	return &MdepositAddPaymentTokenIterator{contract: _Mdeposit.contract, event: "AddPaymentToken", logs: logs, sub: sub}, nil
}

// WatchAddPaymentToken is a free log subscription operation binding the contract event 0x049000a9db89588d7bfb162bc0f7e4299ee8762430a468131c2caf0824f1f995.
//
// Solidity: event AddPaymentToken(address indexed caller, address indexed token, address indexed dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositFilterer) WatchAddPaymentToken(opts *bind.WatchOpts, sink chan<- *MdepositAddPaymentToken, caller []common.Address, token []common.Address, dataFeed []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var dataFeedRule []interface{}
	for _, dataFeedItem := range dataFeed {
		dataFeedRule = append(dataFeedRule, dataFeedItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "AddPaymentToken", callerRule, tokenRule, dataFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositAddPaymentToken)
				if err := _Mdeposit.contract.UnpackLog(event, "AddPaymentToken", log); err != nil {
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

// ParseAddPaymentToken is a log parse operation binding the contract event 0x049000a9db89588d7bfb162bc0f7e4299ee8762430a468131c2caf0824f1f995.
//
// Solidity: event AddPaymentToken(address indexed caller, address indexed token, address indexed dataFeed, uint256 fee, uint256 allowance, bool stable)
func (_Mdeposit *MdepositFilterer) ParseAddPaymentToken(log types.Log) (*MdepositAddPaymentToken, error) {
	event := new(MdepositAddPaymentToken)
	if err := _Mdeposit.contract.UnpackLog(event, "AddPaymentToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositAddWaivedFeeAccountIterator is returned from FilterAddWaivedFeeAccount and is used to iterate over the raw logs and unpacked data for AddWaivedFeeAccount events raised by the Mdeposit contract.
type MdepositAddWaivedFeeAccountIterator struct {
	Event *MdepositAddWaivedFeeAccount // Event containing the contract specifics and raw log

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
func (it *MdepositAddWaivedFeeAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositAddWaivedFeeAccount)
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
		it.Event = new(MdepositAddWaivedFeeAccount)
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
func (it *MdepositAddWaivedFeeAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositAddWaivedFeeAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositAddWaivedFeeAccount represents a AddWaivedFeeAccount event raised by the Mdeposit contract.
type MdepositAddWaivedFeeAccount struct {
	Account common.Address
	Caller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddWaivedFeeAccount is a free log retrieval operation binding the contract event 0x221f04b37331150bcfd05e2de362f50785c29ee4ab14f26d4495a51f3c029060.
//
// Solidity: event AddWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) FilterAddWaivedFeeAccount(opts *bind.FilterOpts, account []common.Address, caller []common.Address) (*MdepositAddWaivedFeeAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "AddWaivedFeeAccount", accountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositAddWaivedFeeAccountIterator{contract: _Mdeposit.contract, event: "AddWaivedFeeAccount", logs: logs, sub: sub}, nil
}

// WatchAddWaivedFeeAccount is a free log subscription operation binding the contract event 0x221f04b37331150bcfd05e2de362f50785c29ee4ab14f26d4495a51f3c029060.
//
// Solidity: event AddWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) WatchAddWaivedFeeAccount(opts *bind.WatchOpts, sink chan<- *MdepositAddWaivedFeeAccount, account []common.Address, caller []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "AddWaivedFeeAccount", accountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositAddWaivedFeeAccount)
				if err := _Mdeposit.contract.UnpackLog(event, "AddWaivedFeeAccount", log); err != nil {
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

// ParseAddWaivedFeeAccount is a log parse operation binding the contract event 0x221f04b37331150bcfd05e2de362f50785c29ee4ab14f26d4495a51f3c029060.
//
// Solidity: event AddWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) ParseAddWaivedFeeAccount(log types.Log) (*MdepositAddWaivedFeeAccount, error) {
	event := new(MdepositAddWaivedFeeAccount)
	if err := _Mdeposit.contract.UnpackLog(event, "AddWaivedFeeAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositApproveRequestIterator is returned from FilterApproveRequest and is used to iterate over the raw logs and unpacked data for ApproveRequest events raised by the Mdeposit contract.
type MdepositApproveRequestIterator struct {
	Event *MdepositApproveRequest // Event containing the contract specifics and raw log

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
func (it *MdepositApproveRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositApproveRequest)
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
		it.Event = new(MdepositApproveRequest)
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
func (it *MdepositApproveRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositApproveRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositApproveRequest represents a ApproveRequest event raised by the Mdeposit contract.
type MdepositApproveRequest struct {
	RequestId  *big.Int
	NewOutRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproveRequest is a free log retrieval operation binding the contract event 0xf7d1fde87f32720fc30ce6847e0aae77e640b59bfac41b11b270358ccfa7a0ac.
//
// Solidity: event ApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) FilterApproveRequest(opts *bind.FilterOpts, requestId []*big.Int) (*MdepositApproveRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "ApproveRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &MdepositApproveRequestIterator{contract: _Mdeposit.contract, event: "ApproveRequest", logs: logs, sub: sub}, nil
}

// WatchApproveRequest is a free log subscription operation binding the contract event 0xf7d1fde87f32720fc30ce6847e0aae77e640b59bfac41b11b270358ccfa7a0ac.
//
// Solidity: event ApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) WatchApproveRequest(opts *bind.WatchOpts, sink chan<- *MdepositApproveRequest, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "ApproveRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositApproveRequest)
				if err := _Mdeposit.contract.UnpackLog(event, "ApproveRequest", log); err != nil {
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

// ParseApproveRequest is a log parse operation binding the contract event 0xf7d1fde87f32720fc30ce6847e0aae77e640b59bfac41b11b270358ccfa7a0ac.
//
// Solidity: event ApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) ParseApproveRequest(log types.Log) (*MdepositApproveRequest, error) {
	event := new(MdepositApproveRequest)
	if err := _Mdeposit.contract.UnpackLog(event, "ApproveRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositChangeTokenAllowanceIterator is returned from FilterChangeTokenAllowance and is used to iterate over the raw logs and unpacked data for ChangeTokenAllowance events raised by the Mdeposit contract.
type MdepositChangeTokenAllowanceIterator struct {
	Event *MdepositChangeTokenAllowance // Event containing the contract specifics and raw log

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
func (it *MdepositChangeTokenAllowanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositChangeTokenAllowance)
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
		it.Event = new(MdepositChangeTokenAllowance)
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
func (it *MdepositChangeTokenAllowanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositChangeTokenAllowanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositChangeTokenAllowance represents a ChangeTokenAllowance event raised by the Mdeposit contract.
type MdepositChangeTokenAllowance struct {
	Token     common.Address
	Caller    common.Address
	Allowance *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeTokenAllowance is a free log retrieval operation binding the contract event 0xf7273742887a46d8b97d83d1d12b6d8d8e6d21d814072369e2f4b355690221d7.
//
// Solidity: event ChangeTokenAllowance(address indexed token, address indexed caller, uint256 allowance)
func (_Mdeposit *MdepositFilterer) FilterChangeTokenAllowance(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*MdepositChangeTokenAllowanceIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "ChangeTokenAllowance", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositChangeTokenAllowanceIterator{contract: _Mdeposit.contract, event: "ChangeTokenAllowance", logs: logs, sub: sub}, nil
}

// WatchChangeTokenAllowance is a free log subscription operation binding the contract event 0xf7273742887a46d8b97d83d1d12b6d8d8e6d21d814072369e2f4b355690221d7.
//
// Solidity: event ChangeTokenAllowance(address indexed token, address indexed caller, uint256 allowance)
func (_Mdeposit *MdepositFilterer) WatchChangeTokenAllowance(opts *bind.WatchOpts, sink chan<- *MdepositChangeTokenAllowance, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "ChangeTokenAllowance", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositChangeTokenAllowance)
				if err := _Mdeposit.contract.UnpackLog(event, "ChangeTokenAllowance", log); err != nil {
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

// ParseChangeTokenAllowance is a log parse operation binding the contract event 0xf7273742887a46d8b97d83d1d12b6d8d8e6d21d814072369e2f4b355690221d7.
//
// Solidity: event ChangeTokenAllowance(address indexed token, address indexed caller, uint256 allowance)
func (_Mdeposit *MdepositFilterer) ParseChangeTokenAllowance(log types.Log) (*MdepositChangeTokenAllowance, error) {
	event := new(MdepositChangeTokenAllowance)
	if err := _Mdeposit.contract.UnpackLog(event, "ChangeTokenAllowance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositChangeTokenFeeIterator is returned from FilterChangeTokenFee and is used to iterate over the raw logs and unpacked data for ChangeTokenFee events raised by the Mdeposit contract.
type MdepositChangeTokenFeeIterator struct {
	Event *MdepositChangeTokenFee // Event containing the contract specifics and raw log

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
func (it *MdepositChangeTokenFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositChangeTokenFee)
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
		it.Event = new(MdepositChangeTokenFee)
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
func (it *MdepositChangeTokenFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositChangeTokenFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositChangeTokenFee represents a ChangeTokenFee event raised by the Mdeposit contract.
type MdepositChangeTokenFee struct {
	Token  common.Address
	Caller common.Address
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeTokenFee is a free log retrieval operation binding the contract event 0x1582567d288d96695cf3fe7280c630a4f1c82fc7e665e1db58468f2960fef869.
//
// Solidity: event ChangeTokenFee(address indexed token, address indexed caller, uint256 fee)
func (_Mdeposit *MdepositFilterer) FilterChangeTokenFee(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*MdepositChangeTokenFeeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "ChangeTokenFee", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositChangeTokenFeeIterator{contract: _Mdeposit.contract, event: "ChangeTokenFee", logs: logs, sub: sub}, nil
}

// WatchChangeTokenFee is a free log subscription operation binding the contract event 0x1582567d288d96695cf3fe7280c630a4f1c82fc7e665e1db58468f2960fef869.
//
// Solidity: event ChangeTokenFee(address indexed token, address indexed caller, uint256 fee)
func (_Mdeposit *MdepositFilterer) WatchChangeTokenFee(opts *bind.WatchOpts, sink chan<- *MdepositChangeTokenFee, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "ChangeTokenFee", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositChangeTokenFee)
				if err := _Mdeposit.contract.UnpackLog(event, "ChangeTokenFee", log); err != nil {
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

// ParseChangeTokenFee is a log parse operation binding the contract event 0x1582567d288d96695cf3fe7280c630a4f1c82fc7e665e1db58468f2960fef869.
//
// Solidity: event ChangeTokenFee(address indexed token, address indexed caller, uint256 fee)
func (_Mdeposit *MdepositFilterer) ParseChangeTokenFee(log types.Log) (*MdepositChangeTokenFee, error) {
	event := new(MdepositChangeTokenFee)
	if err := _Mdeposit.contract.UnpackLog(event, "ChangeTokenFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositDepositInstantIterator is returned from FilterDepositInstant and is used to iterate over the raw logs and unpacked data for DepositInstant events raised by the Mdeposit contract.
type MdepositDepositInstantIterator struct {
	Event *MdepositDepositInstant // Event containing the contract specifics and raw log

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
func (it *MdepositDepositInstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositDepositInstant)
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
		it.Event = new(MdepositDepositInstant)
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
func (it *MdepositDepositInstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositDepositInstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositDepositInstant represents a DepositInstant event raised by the Mdeposit contract.
type MdepositDepositInstant struct {
	User        common.Address
	TokenIn     common.Address
	AmountUsd   *big.Int
	AmountToken *big.Int
	Fee         *big.Int
	Minted      *big.Int
	ReferrerId  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositInstant is a free log retrieval operation binding the contract event 0xdd6865ec496cf9bdd5cb1661ab84cf4e86edc877208a54cbf642f69d744530c5.
//
// Solidity: event DepositInstant(address indexed user, address indexed tokenIn, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) FilterDepositInstant(opts *bind.FilterOpts, user []common.Address, tokenIn []common.Address) (*MdepositDepositInstantIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "DepositInstant", userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &MdepositDepositInstantIterator{contract: _Mdeposit.contract, event: "DepositInstant", logs: logs, sub: sub}, nil
}

// WatchDepositInstant is a free log subscription operation binding the contract event 0xdd6865ec496cf9bdd5cb1661ab84cf4e86edc877208a54cbf642f69d744530c5.
//
// Solidity: event DepositInstant(address indexed user, address indexed tokenIn, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) WatchDepositInstant(opts *bind.WatchOpts, sink chan<- *MdepositDepositInstant, user []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "DepositInstant", userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositDepositInstant)
				if err := _Mdeposit.contract.UnpackLog(event, "DepositInstant", log); err != nil {
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

// ParseDepositInstant is a log parse operation binding the contract event 0xdd6865ec496cf9bdd5cb1661ab84cf4e86edc877208a54cbf642f69d744530c5.
//
// Solidity: event DepositInstant(address indexed user, address indexed tokenIn, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) ParseDepositInstant(log types.Log) (*MdepositDepositInstant, error) {
	event := new(MdepositDepositInstant)
	if err := _Mdeposit.contract.UnpackLog(event, "DepositInstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositDepositInstantWithCustomRecipientIterator is returned from FilterDepositInstantWithCustomRecipient and is used to iterate over the raw logs and unpacked data for DepositInstantWithCustomRecipient events raised by the Mdeposit contract.
type MdepositDepositInstantWithCustomRecipientIterator struct {
	Event *MdepositDepositInstantWithCustomRecipient // Event containing the contract specifics and raw log

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
func (it *MdepositDepositInstantWithCustomRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositDepositInstantWithCustomRecipient)
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
		it.Event = new(MdepositDepositInstantWithCustomRecipient)
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
func (it *MdepositDepositInstantWithCustomRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositDepositInstantWithCustomRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositDepositInstantWithCustomRecipient represents a DepositInstantWithCustomRecipient event raised by the Mdeposit contract.
type MdepositDepositInstantWithCustomRecipient struct {
	User        common.Address
	TokenIn     common.Address
	Recipient   common.Address
	AmountUsd   *big.Int
	AmountToken *big.Int
	Fee         *big.Int
	Minted      *big.Int
	ReferrerId  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositInstantWithCustomRecipient is a free log retrieval operation binding the contract event 0xe8bfe7b6cdaff26f82915adfad787fe8cc232bf312d39f4eab839d013e65da5a.
//
// Solidity: event DepositInstantWithCustomRecipient(address indexed user, address indexed tokenIn, address recipient, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) FilterDepositInstantWithCustomRecipient(opts *bind.FilterOpts, user []common.Address, tokenIn []common.Address) (*MdepositDepositInstantWithCustomRecipientIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "DepositInstantWithCustomRecipient", userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &MdepositDepositInstantWithCustomRecipientIterator{contract: _Mdeposit.contract, event: "DepositInstantWithCustomRecipient", logs: logs, sub: sub}, nil
}

// WatchDepositInstantWithCustomRecipient is a free log subscription operation binding the contract event 0xe8bfe7b6cdaff26f82915adfad787fe8cc232bf312d39f4eab839d013e65da5a.
//
// Solidity: event DepositInstantWithCustomRecipient(address indexed user, address indexed tokenIn, address recipient, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) WatchDepositInstantWithCustomRecipient(opts *bind.WatchOpts, sink chan<- *MdepositDepositInstantWithCustomRecipient, user []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "DepositInstantWithCustomRecipient", userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositDepositInstantWithCustomRecipient)
				if err := _Mdeposit.contract.UnpackLog(event, "DepositInstantWithCustomRecipient", log); err != nil {
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

// ParseDepositInstantWithCustomRecipient is a log parse operation binding the contract event 0xe8bfe7b6cdaff26f82915adfad787fe8cc232bf312d39f4eab839d013e65da5a.
//
// Solidity: event DepositInstantWithCustomRecipient(address indexed user, address indexed tokenIn, address recipient, uint256 amountUsd, uint256 amountToken, uint256 fee, uint256 minted, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) ParseDepositInstantWithCustomRecipient(log types.Log) (*MdepositDepositInstantWithCustomRecipient, error) {
	event := new(MdepositDepositInstantWithCustomRecipient)
	if err := _Mdeposit.contract.UnpackLog(event, "DepositInstantWithCustomRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositDepositRequestIterator is returned from FilterDepositRequest and is used to iterate over the raw logs and unpacked data for DepositRequest events raised by the Mdeposit contract.
type MdepositDepositRequestIterator struct {
	Event *MdepositDepositRequest // Event containing the contract specifics and raw log

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
func (it *MdepositDepositRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositDepositRequest)
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
		it.Event = new(MdepositDepositRequest)
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
func (it *MdepositDepositRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositDepositRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositDepositRequest represents a DepositRequest event raised by the Mdeposit contract.
type MdepositDepositRequest struct {
	RequestId    *big.Int
	User         common.Address
	TokenIn      common.Address
	AmountToken  *big.Int
	AmountUsd    *big.Int
	Fee          *big.Int
	TokenOutRate *big.Int
	ReferrerId   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositRequest is a free log retrieval operation binding the contract event 0x3704c9b13a68ac43d7f8a85f2700f0b4f89a11ed9e2bcac5324f0d228d409009.
//
// Solidity: event DepositRequest(uint256 indexed requestId, address indexed user, address indexed tokenIn, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) FilterDepositRequest(opts *bind.FilterOpts, requestId []*big.Int, user []common.Address, tokenIn []common.Address) (*MdepositDepositRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "DepositRequest", requestIdRule, userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &MdepositDepositRequestIterator{contract: _Mdeposit.contract, event: "DepositRequest", logs: logs, sub: sub}, nil
}

// WatchDepositRequest is a free log subscription operation binding the contract event 0x3704c9b13a68ac43d7f8a85f2700f0b4f89a11ed9e2bcac5324f0d228d409009.
//
// Solidity: event DepositRequest(uint256 indexed requestId, address indexed user, address indexed tokenIn, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) WatchDepositRequest(opts *bind.WatchOpts, sink chan<- *MdepositDepositRequest, requestId []*big.Int, user []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "DepositRequest", requestIdRule, userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositDepositRequest)
				if err := _Mdeposit.contract.UnpackLog(event, "DepositRequest", log); err != nil {
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

// ParseDepositRequest is a log parse operation binding the contract event 0x3704c9b13a68ac43d7f8a85f2700f0b4f89a11ed9e2bcac5324f0d228d409009.
//
// Solidity: event DepositRequest(uint256 indexed requestId, address indexed user, address indexed tokenIn, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) ParseDepositRequest(log types.Log) (*MdepositDepositRequest, error) {
	event := new(MdepositDepositRequest)
	if err := _Mdeposit.contract.UnpackLog(event, "DepositRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositDepositRequestWithCustomRecipientIterator is returned from FilterDepositRequestWithCustomRecipient and is used to iterate over the raw logs and unpacked data for DepositRequestWithCustomRecipient events raised by the Mdeposit contract.
type MdepositDepositRequestWithCustomRecipientIterator struct {
	Event *MdepositDepositRequestWithCustomRecipient // Event containing the contract specifics and raw log

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
func (it *MdepositDepositRequestWithCustomRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositDepositRequestWithCustomRecipient)
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
		it.Event = new(MdepositDepositRequestWithCustomRecipient)
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
func (it *MdepositDepositRequestWithCustomRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositDepositRequestWithCustomRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositDepositRequestWithCustomRecipient represents a DepositRequestWithCustomRecipient event raised by the Mdeposit contract.
type MdepositDepositRequestWithCustomRecipient struct {
	RequestId    *big.Int
	User         common.Address
	TokenIn      common.Address
	Recipient    common.Address
	AmountToken  *big.Int
	AmountUsd    *big.Int
	Fee          *big.Int
	TokenOutRate *big.Int
	ReferrerId   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositRequestWithCustomRecipient is a free log retrieval operation binding the contract event 0xd21eaf3019cc16da5c82b2c14e3df524c0599086f690f48357de2c74f1bbdfd6.
//
// Solidity: event DepositRequestWithCustomRecipient(uint256 indexed requestId, address indexed user, address indexed tokenIn, address recipient, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) FilterDepositRequestWithCustomRecipient(opts *bind.FilterOpts, requestId []*big.Int, user []common.Address, tokenIn []common.Address) (*MdepositDepositRequestWithCustomRecipientIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "DepositRequestWithCustomRecipient", requestIdRule, userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &MdepositDepositRequestWithCustomRecipientIterator{contract: _Mdeposit.contract, event: "DepositRequestWithCustomRecipient", logs: logs, sub: sub}, nil
}

// WatchDepositRequestWithCustomRecipient is a free log subscription operation binding the contract event 0xd21eaf3019cc16da5c82b2c14e3df524c0599086f690f48357de2c74f1bbdfd6.
//
// Solidity: event DepositRequestWithCustomRecipient(uint256 indexed requestId, address indexed user, address indexed tokenIn, address recipient, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) WatchDepositRequestWithCustomRecipient(opts *bind.WatchOpts, sink chan<- *MdepositDepositRequestWithCustomRecipient, requestId []*big.Int, user []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "DepositRequestWithCustomRecipient", requestIdRule, userRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositDepositRequestWithCustomRecipient)
				if err := _Mdeposit.contract.UnpackLog(event, "DepositRequestWithCustomRecipient", log); err != nil {
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

// ParseDepositRequestWithCustomRecipient is a log parse operation binding the contract event 0xd21eaf3019cc16da5c82b2c14e3df524c0599086f690f48357de2c74f1bbdfd6.
//
// Solidity: event DepositRequestWithCustomRecipient(uint256 indexed requestId, address indexed user, address indexed tokenIn, address recipient, uint256 amountToken, uint256 amountUsd, uint256 fee, uint256 tokenOutRate, bytes32 referrerId)
func (_Mdeposit *MdepositFilterer) ParseDepositRequestWithCustomRecipient(log types.Log) (*MdepositDepositRequestWithCustomRecipient, error) {
	event := new(MdepositDepositRequestWithCustomRecipient)
	if err := _Mdeposit.contract.UnpackLog(event, "DepositRequestWithCustomRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositFreeFromMinAmountIterator is returned from FilterFreeFromMinAmount and is used to iterate over the raw logs and unpacked data for FreeFromMinAmount events raised by the Mdeposit contract.
type MdepositFreeFromMinAmountIterator struct {
	Event *MdepositFreeFromMinAmount // Event containing the contract specifics and raw log

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
func (it *MdepositFreeFromMinAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositFreeFromMinAmount)
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
		it.Event = new(MdepositFreeFromMinAmount)
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
func (it *MdepositFreeFromMinAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositFreeFromMinAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositFreeFromMinAmount represents a FreeFromMinAmount event raised by the Mdeposit contract.
type MdepositFreeFromMinAmount struct {
	User   common.Address
	Enable bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFreeFromMinAmount is a free log retrieval operation binding the contract event 0x80f6f2f8801c6ac8fc60bf218b44fde97744d8709f69281972ec5557c10226cc.
//
// Solidity: event FreeFromMinAmount(address indexed user, bool enable)
func (_Mdeposit *MdepositFilterer) FilterFreeFromMinAmount(opts *bind.FilterOpts, user []common.Address) (*MdepositFreeFromMinAmountIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "FreeFromMinAmount", userRule)
	if err != nil {
		return nil, err
	}
	return &MdepositFreeFromMinAmountIterator{contract: _Mdeposit.contract, event: "FreeFromMinAmount", logs: logs, sub: sub}, nil
}

// WatchFreeFromMinAmount is a free log subscription operation binding the contract event 0x80f6f2f8801c6ac8fc60bf218b44fde97744d8709f69281972ec5557c10226cc.
//
// Solidity: event FreeFromMinAmount(address indexed user, bool enable)
func (_Mdeposit *MdepositFilterer) WatchFreeFromMinAmount(opts *bind.WatchOpts, sink chan<- *MdepositFreeFromMinAmount, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "FreeFromMinAmount", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositFreeFromMinAmount)
				if err := _Mdeposit.contract.UnpackLog(event, "FreeFromMinAmount", log); err != nil {
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

// ParseFreeFromMinAmount is a log parse operation binding the contract event 0x80f6f2f8801c6ac8fc60bf218b44fde97744d8709f69281972ec5557c10226cc.
//
// Solidity: event FreeFromMinAmount(address indexed user, bool enable)
func (_Mdeposit *MdepositFilterer) ParseFreeFromMinAmount(log types.Log) (*MdepositFreeFromMinAmount, error) {
	event := new(MdepositFreeFromMinAmount)
	if err := _Mdeposit.contract.UnpackLog(event, "FreeFromMinAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositFreeFromMinDepositIterator is returned from FilterFreeFromMinDeposit and is used to iterate over the raw logs and unpacked data for FreeFromMinDeposit events raised by the Mdeposit contract.
type MdepositFreeFromMinDepositIterator struct {
	Event *MdepositFreeFromMinDeposit // Event containing the contract specifics and raw log

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
func (it *MdepositFreeFromMinDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositFreeFromMinDeposit)
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
		it.Event = new(MdepositFreeFromMinDeposit)
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
func (it *MdepositFreeFromMinDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositFreeFromMinDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositFreeFromMinDeposit represents a FreeFromMinDeposit event raised by the Mdeposit contract.
type MdepositFreeFromMinDeposit struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFreeFromMinDeposit is a free log retrieval operation binding the contract event 0xcbf8984cbb78b0361a17992ef91f83544b26a3d8b93c5cce6d7d81bf42472c5d.
//
// Solidity: event FreeFromMinDeposit(address indexed user)
func (_Mdeposit *MdepositFilterer) FilterFreeFromMinDeposit(opts *bind.FilterOpts, user []common.Address) (*MdepositFreeFromMinDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "FreeFromMinDeposit", userRule)
	if err != nil {
		return nil, err
	}
	return &MdepositFreeFromMinDepositIterator{contract: _Mdeposit.contract, event: "FreeFromMinDeposit", logs: logs, sub: sub}, nil
}

// WatchFreeFromMinDeposit is a free log subscription operation binding the contract event 0xcbf8984cbb78b0361a17992ef91f83544b26a3d8b93c5cce6d7d81bf42472c5d.
//
// Solidity: event FreeFromMinDeposit(address indexed user)
func (_Mdeposit *MdepositFilterer) WatchFreeFromMinDeposit(opts *bind.WatchOpts, sink chan<- *MdepositFreeFromMinDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "FreeFromMinDeposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositFreeFromMinDeposit)
				if err := _Mdeposit.contract.UnpackLog(event, "FreeFromMinDeposit", log); err != nil {
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

// ParseFreeFromMinDeposit is a log parse operation binding the contract event 0xcbf8984cbb78b0361a17992ef91f83544b26a3d8b93c5cce6d7d81bf42472c5d.
//
// Solidity: event FreeFromMinDeposit(address indexed user)
func (_Mdeposit *MdepositFilterer) ParseFreeFromMinDeposit(log types.Log) (*MdepositFreeFromMinDeposit, error) {
	event := new(MdepositFreeFromMinDeposit)
	if err := _Mdeposit.contract.UnpackLog(event, "FreeFromMinDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mdeposit contract.
type MdepositInitializedIterator struct {
	Event *MdepositInitialized // Event containing the contract specifics and raw log

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
func (it *MdepositInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositInitialized)
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
		it.Event = new(MdepositInitialized)
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
func (it *MdepositInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositInitialized represents a Initialized event raised by the Mdeposit contract.
type MdepositInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mdeposit *MdepositFilterer) FilterInitialized(opts *bind.FilterOpts) (*MdepositInitializedIterator, error) {

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MdepositInitializedIterator{contract: _Mdeposit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mdeposit *MdepositFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MdepositInitialized) (event.Subscription, error) {

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositInitialized)
				if err := _Mdeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mdeposit *MdepositFilterer) ParseInitialized(log types.Log) (*MdepositInitialized, error) {
	event := new(MdepositInitialized)
	if err := _Mdeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositPauseFnIterator is returned from FilterPauseFn and is used to iterate over the raw logs and unpacked data for PauseFn events raised by the Mdeposit contract.
type MdepositPauseFnIterator struct {
	Event *MdepositPauseFn // Event containing the contract specifics and raw log

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
func (it *MdepositPauseFnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositPauseFn)
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
		it.Event = new(MdepositPauseFn)
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
func (it *MdepositPauseFnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositPauseFnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositPauseFn represents a PauseFn event raised by the Mdeposit contract.
type MdepositPauseFn struct {
	Caller common.Address
	Fn     [4]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPauseFn is a free log retrieval operation binding the contract event 0x2278e547293e53a66144c1743877f8388ac3101bd21cfd7c7f4ce8c15c14f5c1.
//
// Solidity: event PauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) FilterPauseFn(opts *bind.FilterOpts, caller []common.Address) (*MdepositPauseFnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "PauseFn", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositPauseFnIterator{contract: _Mdeposit.contract, event: "PauseFn", logs: logs, sub: sub}, nil
}

// WatchPauseFn is a free log subscription operation binding the contract event 0x2278e547293e53a66144c1743877f8388ac3101bd21cfd7c7f4ce8c15c14f5c1.
//
// Solidity: event PauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) WatchPauseFn(opts *bind.WatchOpts, sink chan<- *MdepositPauseFn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "PauseFn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositPauseFn)
				if err := _Mdeposit.contract.UnpackLog(event, "PauseFn", log); err != nil {
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

// ParsePauseFn is a log parse operation binding the contract event 0x2278e547293e53a66144c1743877f8388ac3101bd21cfd7c7f4ce8c15c14f5c1.
//
// Solidity: event PauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) ParsePauseFn(log types.Log) (*MdepositPauseFn, error) {
	event := new(MdepositPauseFn)
	if err := _Mdeposit.contract.UnpackLog(event, "PauseFn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Mdeposit contract.
type MdepositPausedIterator struct {
	Event *MdepositPaused // Event containing the contract specifics and raw log

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
func (it *MdepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositPaused)
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
		it.Event = new(MdepositPaused)
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
func (it *MdepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositPaused represents a Paused event raised by the Mdeposit contract.
type MdepositPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Mdeposit *MdepositFilterer) FilterPaused(opts *bind.FilterOpts) (*MdepositPausedIterator, error) {

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MdepositPausedIterator{contract: _Mdeposit.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Mdeposit *MdepositFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MdepositPaused) (event.Subscription, error) {

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositPaused)
				if err := _Mdeposit.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Mdeposit *MdepositFilterer) ParsePaused(log types.Log) (*MdepositPaused, error) {
	event := new(MdepositPaused)
	if err := _Mdeposit.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositRejectRequestIterator is returned from FilterRejectRequest and is used to iterate over the raw logs and unpacked data for RejectRequest events raised by the Mdeposit contract.
type MdepositRejectRequestIterator struct {
	Event *MdepositRejectRequest // Event containing the contract specifics and raw log

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
func (it *MdepositRejectRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositRejectRequest)
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
		it.Event = new(MdepositRejectRequest)
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
func (it *MdepositRejectRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositRejectRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositRejectRequest represents a RejectRequest event raised by the Mdeposit contract.
type MdepositRejectRequest struct {
	RequestId *big.Int
	User      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRejectRequest is a free log retrieval operation binding the contract event 0x00ce63cc55966b103e4f4cb39f3426cb91718ad4f8eb4ad08c14a7ee749d8157.
//
// Solidity: event RejectRequest(uint256 indexed requestId, address indexed user)
func (_Mdeposit *MdepositFilterer) FilterRejectRequest(opts *bind.FilterOpts, requestId []*big.Int, user []common.Address) (*MdepositRejectRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "RejectRequest", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &MdepositRejectRequestIterator{contract: _Mdeposit.contract, event: "RejectRequest", logs: logs, sub: sub}, nil
}

// WatchRejectRequest is a free log subscription operation binding the contract event 0x00ce63cc55966b103e4f4cb39f3426cb91718ad4f8eb4ad08c14a7ee749d8157.
//
// Solidity: event RejectRequest(uint256 indexed requestId, address indexed user)
func (_Mdeposit *MdepositFilterer) WatchRejectRequest(opts *bind.WatchOpts, sink chan<- *MdepositRejectRequest, requestId []*big.Int, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "RejectRequest", requestIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositRejectRequest)
				if err := _Mdeposit.contract.UnpackLog(event, "RejectRequest", log); err != nil {
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

// ParseRejectRequest is a log parse operation binding the contract event 0x00ce63cc55966b103e4f4cb39f3426cb91718ad4f8eb4ad08c14a7ee749d8157.
//
// Solidity: event RejectRequest(uint256 indexed requestId, address indexed user)
func (_Mdeposit *MdepositFilterer) ParseRejectRequest(log types.Log) (*MdepositRejectRequest, error) {
	event := new(MdepositRejectRequest)
	if err := _Mdeposit.contract.UnpackLog(event, "RejectRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositRemovePaymentTokenIterator is returned from FilterRemovePaymentToken and is used to iterate over the raw logs and unpacked data for RemovePaymentToken events raised by the Mdeposit contract.
type MdepositRemovePaymentTokenIterator struct {
	Event *MdepositRemovePaymentToken // Event containing the contract specifics and raw log

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
func (it *MdepositRemovePaymentTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositRemovePaymentToken)
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
		it.Event = new(MdepositRemovePaymentToken)
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
func (it *MdepositRemovePaymentTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositRemovePaymentTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositRemovePaymentToken represents a RemovePaymentToken event raised by the Mdeposit contract.
type MdepositRemovePaymentToken struct {
	Token  common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovePaymentToken is a free log retrieval operation binding the contract event 0x652fa2f5d587d3f1c189df0081b7bf3121f47d51d5471bf58d7d2c8a084894c3.
//
// Solidity: event RemovePaymentToken(address indexed token, address indexed caller)
func (_Mdeposit *MdepositFilterer) FilterRemovePaymentToken(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*MdepositRemovePaymentTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "RemovePaymentToken", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositRemovePaymentTokenIterator{contract: _Mdeposit.contract, event: "RemovePaymentToken", logs: logs, sub: sub}, nil
}

// WatchRemovePaymentToken is a free log subscription operation binding the contract event 0x652fa2f5d587d3f1c189df0081b7bf3121f47d51d5471bf58d7d2c8a084894c3.
//
// Solidity: event RemovePaymentToken(address indexed token, address indexed caller)
func (_Mdeposit *MdepositFilterer) WatchRemovePaymentToken(opts *bind.WatchOpts, sink chan<- *MdepositRemovePaymentToken, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "RemovePaymentToken", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositRemovePaymentToken)
				if err := _Mdeposit.contract.UnpackLog(event, "RemovePaymentToken", log); err != nil {
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

// ParseRemovePaymentToken is a log parse operation binding the contract event 0x652fa2f5d587d3f1c189df0081b7bf3121f47d51d5471bf58d7d2c8a084894c3.
//
// Solidity: event RemovePaymentToken(address indexed token, address indexed caller)
func (_Mdeposit *MdepositFilterer) ParseRemovePaymentToken(log types.Log) (*MdepositRemovePaymentToken, error) {
	event := new(MdepositRemovePaymentToken)
	if err := _Mdeposit.contract.UnpackLog(event, "RemovePaymentToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositRemoveWaivedFeeAccountIterator is returned from FilterRemoveWaivedFeeAccount and is used to iterate over the raw logs and unpacked data for RemoveWaivedFeeAccount events raised by the Mdeposit contract.
type MdepositRemoveWaivedFeeAccountIterator struct {
	Event *MdepositRemoveWaivedFeeAccount // Event containing the contract specifics and raw log

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
func (it *MdepositRemoveWaivedFeeAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositRemoveWaivedFeeAccount)
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
		it.Event = new(MdepositRemoveWaivedFeeAccount)
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
func (it *MdepositRemoveWaivedFeeAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositRemoveWaivedFeeAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositRemoveWaivedFeeAccount represents a RemoveWaivedFeeAccount event raised by the Mdeposit contract.
type MdepositRemoveWaivedFeeAccount struct {
	Account common.Address
	Caller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveWaivedFeeAccount is a free log retrieval operation binding the contract event 0x57c4a95f59c12f0d4d846443c2d54c7d97f1505080199522fca2819e65213ca2.
//
// Solidity: event RemoveWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) FilterRemoveWaivedFeeAccount(opts *bind.FilterOpts, account []common.Address, caller []common.Address) (*MdepositRemoveWaivedFeeAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "RemoveWaivedFeeAccount", accountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositRemoveWaivedFeeAccountIterator{contract: _Mdeposit.contract, event: "RemoveWaivedFeeAccount", logs: logs, sub: sub}, nil
}

// WatchRemoveWaivedFeeAccount is a free log subscription operation binding the contract event 0x57c4a95f59c12f0d4d846443c2d54c7d97f1505080199522fca2819e65213ca2.
//
// Solidity: event RemoveWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) WatchRemoveWaivedFeeAccount(opts *bind.WatchOpts, sink chan<- *MdepositRemoveWaivedFeeAccount, account []common.Address, caller []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "RemoveWaivedFeeAccount", accountRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositRemoveWaivedFeeAccount)
				if err := _Mdeposit.contract.UnpackLog(event, "RemoveWaivedFeeAccount", log); err != nil {
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

// ParseRemoveWaivedFeeAccount is a log parse operation binding the contract event 0x57c4a95f59c12f0d4d846443c2d54c7d97f1505080199522fca2819e65213ca2.
//
// Solidity: event RemoveWaivedFeeAccount(address indexed account, address indexed caller)
func (_Mdeposit *MdepositFilterer) ParseRemoveWaivedFeeAccount(log types.Log) (*MdepositRemoveWaivedFeeAccount, error) {
	event := new(MdepositRemoveWaivedFeeAccount)
	if err := _Mdeposit.contract.UnpackLog(event, "RemoveWaivedFeeAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSafeApproveRequestIterator is returned from FilterSafeApproveRequest and is used to iterate over the raw logs and unpacked data for SafeApproveRequest events raised by the Mdeposit contract.
type MdepositSafeApproveRequestIterator struct {
	Event *MdepositSafeApproveRequest // Event containing the contract specifics and raw log

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
func (it *MdepositSafeApproveRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSafeApproveRequest)
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
		it.Event = new(MdepositSafeApproveRequest)
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
func (it *MdepositSafeApproveRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSafeApproveRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSafeApproveRequest represents a SafeApproveRequest event raised by the Mdeposit contract.
type MdepositSafeApproveRequest struct {
	RequestId  *big.Int
	NewOutRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSafeApproveRequest is a free log retrieval operation binding the contract event 0x03ea09e71742c9c754c9746b3e671ecb27fc372e3d29c31bac0192458ffd9d4b.
//
// Solidity: event SafeApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) FilterSafeApproveRequest(opts *bind.FilterOpts, requestId []*big.Int) (*MdepositSafeApproveRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SafeApproveRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSafeApproveRequestIterator{contract: _Mdeposit.contract, event: "SafeApproveRequest", logs: logs, sub: sub}, nil
}

// WatchSafeApproveRequest is a free log subscription operation binding the contract event 0x03ea09e71742c9c754c9746b3e671ecb27fc372e3d29c31bac0192458ffd9d4b.
//
// Solidity: event SafeApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) WatchSafeApproveRequest(opts *bind.WatchOpts, sink chan<- *MdepositSafeApproveRequest, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SafeApproveRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSafeApproveRequest)
				if err := _Mdeposit.contract.UnpackLog(event, "SafeApproveRequest", log); err != nil {
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

// ParseSafeApproveRequest is a log parse operation binding the contract event 0x03ea09e71742c9c754c9746b3e671ecb27fc372e3d29c31bac0192458ffd9d4b.
//
// Solidity: event SafeApproveRequest(uint256 indexed requestId, uint256 newOutRate)
func (_Mdeposit *MdepositFilterer) ParseSafeApproveRequest(log types.Log) (*MdepositSafeApproveRequest, error) {
	event := new(MdepositSafeApproveRequest)
	if err := _Mdeposit.contract.UnpackLog(event, "SafeApproveRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetFeeReceiverIterator is returned from FilterSetFeeReceiver and is used to iterate over the raw logs and unpacked data for SetFeeReceiver events raised by the Mdeposit contract.
type MdepositSetFeeReceiverIterator struct {
	Event *MdepositSetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *MdepositSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetFeeReceiver)
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
		it.Event = new(MdepositSetFeeReceiver)
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
func (it *MdepositSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetFeeReceiver represents a SetFeeReceiver event raised by the Mdeposit contract.
type MdepositSetFeeReceiver struct {
	Caller   common.Address
	Reciever common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetFeeReceiver is a free log retrieval operation binding the contract event 0x1b092cca381ac00a07e1226c164f47c475d212f5e55699475a7f411811f77dd4.
//
// Solidity: event SetFeeReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) FilterSetFeeReceiver(opts *bind.FilterOpts, caller []common.Address, reciever []common.Address) (*MdepositSetFeeReceiverIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetFeeReceiver", callerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetFeeReceiverIterator{contract: _Mdeposit.contract, event: "SetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchSetFeeReceiver is a free log subscription operation binding the contract event 0x1b092cca381ac00a07e1226c164f47c475d212f5e55699475a7f411811f77dd4.
//
// Solidity: event SetFeeReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) WatchSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *MdepositSetFeeReceiver, caller []common.Address, reciever []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetFeeReceiver", callerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetFeeReceiver)
				if err := _Mdeposit.contract.UnpackLog(event, "SetFeeReceiver", log); err != nil {
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

// ParseSetFeeReceiver is a log parse operation binding the contract event 0x1b092cca381ac00a07e1226c164f47c475d212f5e55699475a7f411811f77dd4.
//
// Solidity: event SetFeeReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) ParseSetFeeReceiver(log types.Log) (*MdepositSetFeeReceiver, error) {
	event := new(MdepositSetFeeReceiver)
	if err := _Mdeposit.contract.UnpackLog(event, "SetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetGreenlistEnableIterator is returned from FilterSetGreenlistEnable and is used to iterate over the raw logs and unpacked data for SetGreenlistEnable events raised by the Mdeposit contract.
type MdepositSetGreenlistEnableIterator struct {
	Event *MdepositSetGreenlistEnable // Event containing the contract specifics and raw log

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
func (it *MdepositSetGreenlistEnableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetGreenlistEnable)
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
		it.Event = new(MdepositSetGreenlistEnable)
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
func (it *MdepositSetGreenlistEnableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetGreenlistEnableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetGreenlistEnable represents a SetGreenlistEnable event raised by the Mdeposit contract.
type MdepositSetGreenlistEnable struct {
	Sender common.Address
	Enable bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetGreenlistEnable is a free log retrieval operation binding the contract event 0xa8434267b880129bc4ba30249aa4a2ac349e8997c699282a9f70562f0f152f54.
//
// Solidity: event SetGreenlistEnable(address indexed sender, bool enable)
func (_Mdeposit *MdepositFilterer) FilterSetGreenlistEnable(opts *bind.FilterOpts, sender []common.Address) (*MdepositSetGreenlistEnableIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetGreenlistEnable", senderRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetGreenlistEnableIterator{contract: _Mdeposit.contract, event: "SetGreenlistEnable", logs: logs, sub: sub}, nil
}

// WatchSetGreenlistEnable is a free log subscription operation binding the contract event 0xa8434267b880129bc4ba30249aa4a2ac349e8997c699282a9f70562f0f152f54.
//
// Solidity: event SetGreenlistEnable(address indexed sender, bool enable)
func (_Mdeposit *MdepositFilterer) WatchSetGreenlistEnable(opts *bind.WatchOpts, sink chan<- *MdepositSetGreenlistEnable, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetGreenlistEnable", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetGreenlistEnable)
				if err := _Mdeposit.contract.UnpackLog(event, "SetGreenlistEnable", log); err != nil {
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

// ParseSetGreenlistEnable is a log parse operation binding the contract event 0xa8434267b880129bc4ba30249aa4a2ac349e8997c699282a9f70562f0f152f54.
//
// Solidity: event SetGreenlistEnable(address indexed sender, bool enable)
func (_Mdeposit *MdepositFilterer) ParseSetGreenlistEnable(log types.Log) (*MdepositSetGreenlistEnable, error) {
	event := new(MdepositSetGreenlistEnable)
	if err := _Mdeposit.contract.UnpackLog(event, "SetGreenlistEnable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetInstantDailyLimitIterator is returned from FilterSetInstantDailyLimit and is used to iterate over the raw logs and unpacked data for SetInstantDailyLimit events raised by the Mdeposit contract.
type MdepositSetInstantDailyLimitIterator struct {
	Event *MdepositSetInstantDailyLimit // Event containing the contract specifics and raw log

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
func (it *MdepositSetInstantDailyLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetInstantDailyLimit)
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
		it.Event = new(MdepositSetInstantDailyLimit)
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
func (it *MdepositSetInstantDailyLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetInstantDailyLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetInstantDailyLimit represents a SetInstantDailyLimit event raised by the Mdeposit contract.
type MdepositSetInstantDailyLimit struct {
	Caller   common.Address
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetInstantDailyLimit is a free log retrieval operation binding the contract event 0x5e8309fc6b2360e7438bc53790b00913395fffa870f39043fe63ddc8a438a9b2.
//
// Solidity: event SetInstantDailyLimit(address indexed caller, uint256 newLimit)
func (_Mdeposit *MdepositFilterer) FilterSetInstantDailyLimit(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetInstantDailyLimitIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetInstantDailyLimit", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetInstantDailyLimitIterator{contract: _Mdeposit.contract, event: "SetInstantDailyLimit", logs: logs, sub: sub}, nil
}

// WatchSetInstantDailyLimit is a free log subscription operation binding the contract event 0x5e8309fc6b2360e7438bc53790b00913395fffa870f39043fe63ddc8a438a9b2.
//
// Solidity: event SetInstantDailyLimit(address indexed caller, uint256 newLimit)
func (_Mdeposit *MdepositFilterer) WatchSetInstantDailyLimit(opts *bind.WatchOpts, sink chan<- *MdepositSetInstantDailyLimit, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetInstantDailyLimit", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetInstantDailyLimit)
				if err := _Mdeposit.contract.UnpackLog(event, "SetInstantDailyLimit", log); err != nil {
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

// ParseSetInstantDailyLimit is a log parse operation binding the contract event 0x5e8309fc6b2360e7438bc53790b00913395fffa870f39043fe63ddc8a438a9b2.
//
// Solidity: event SetInstantDailyLimit(address indexed caller, uint256 newLimit)
func (_Mdeposit *MdepositFilterer) ParseSetInstantDailyLimit(log types.Log) (*MdepositSetInstantDailyLimit, error) {
	event := new(MdepositSetInstantDailyLimit)
	if err := _Mdeposit.contract.UnpackLog(event, "SetInstantDailyLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetInstantFeeIterator is returned from FilterSetInstantFee and is used to iterate over the raw logs and unpacked data for SetInstantFee events raised by the Mdeposit contract.
type MdepositSetInstantFeeIterator struct {
	Event *MdepositSetInstantFee // Event containing the contract specifics and raw log

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
func (it *MdepositSetInstantFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetInstantFee)
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
		it.Event = new(MdepositSetInstantFee)
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
func (it *MdepositSetInstantFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetInstantFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetInstantFee represents a SetInstantFee event raised by the Mdeposit contract.
type MdepositSetInstantFee struct {
	Caller common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetInstantFee is a free log retrieval operation binding the contract event 0x45acc8bd6ebd6fbb59ce049b682c124aeccc93c468fcf60fecf61340e86e79d3.
//
// Solidity: event SetInstantFee(address indexed caller, uint256 newFee)
func (_Mdeposit *MdepositFilterer) FilterSetInstantFee(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetInstantFeeIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetInstantFee", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetInstantFeeIterator{contract: _Mdeposit.contract, event: "SetInstantFee", logs: logs, sub: sub}, nil
}

// WatchSetInstantFee is a free log subscription operation binding the contract event 0x45acc8bd6ebd6fbb59ce049b682c124aeccc93c468fcf60fecf61340e86e79d3.
//
// Solidity: event SetInstantFee(address indexed caller, uint256 newFee)
func (_Mdeposit *MdepositFilterer) WatchSetInstantFee(opts *bind.WatchOpts, sink chan<- *MdepositSetInstantFee, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetInstantFee", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetInstantFee)
				if err := _Mdeposit.contract.UnpackLog(event, "SetInstantFee", log); err != nil {
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

// ParseSetInstantFee is a log parse operation binding the contract event 0x45acc8bd6ebd6fbb59ce049b682c124aeccc93c468fcf60fecf61340e86e79d3.
//
// Solidity: event SetInstantFee(address indexed caller, uint256 newFee)
func (_Mdeposit *MdepositFilterer) ParseSetInstantFee(log types.Log) (*MdepositSetInstantFee, error) {
	event := new(MdepositSetInstantFee)
	if err := _Mdeposit.contract.UnpackLog(event, "SetInstantFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetMaxSupplyCapIterator is returned from FilterSetMaxSupplyCap and is used to iterate over the raw logs and unpacked data for SetMaxSupplyCap events raised by the Mdeposit contract.
type MdepositSetMaxSupplyCapIterator struct {
	Event *MdepositSetMaxSupplyCap // Event containing the contract specifics and raw log

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
func (it *MdepositSetMaxSupplyCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetMaxSupplyCap)
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
		it.Event = new(MdepositSetMaxSupplyCap)
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
func (it *MdepositSetMaxSupplyCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetMaxSupplyCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetMaxSupplyCap represents a SetMaxSupplyCap event raised by the Mdeposit contract.
type MdepositSetMaxSupplyCap struct {
	Caller   common.Address
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMaxSupplyCap is a free log retrieval operation binding the contract event 0x27fdd491319f206fb1fd2dc08831db9085ca1cb91b5795f59c52fee3cd7cb22f.
//
// Solidity: event SetMaxSupplyCap(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) FilterSetMaxSupplyCap(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetMaxSupplyCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetMaxSupplyCap", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetMaxSupplyCapIterator{contract: _Mdeposit.contract, event: "SetMaxSupplyCap", logs: logs, sub: sub}, nil
}

// WatchSetMaxSupplyCap is a free log subscription operation binding the contract event 0x27fdd491319f206fb1fd2dc08831db9085ca1cb91b5795f59c52fee3cd7cb22f.
//
// Solidity: event SetMaxSupplyCap(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) WatchSetMaxSupplyCap(opts *bind.WatchOpts, sink chan<- *MdepositSetMaxSupplyCap, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetMaxSupplyCap", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetMaxSupplyCap)
				if err := _Mdeposit.contract.UnpackLog(event, "SetMaxSupplyCap", log); err != nil {
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

// ParseSetMaxSupplyCap is a log parse operation binding the contract event 0x27fdd491319f206fb1fd2dc08831db9085ca1cb91b5795f59c52fee3cd7cb22f.
//
// Solidity: event SetMaxSupplyCap(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) ParseSetMaxSupplyCap(log types.Log) (*MdepositSetMaxSupplyCap, error) {
	event := new(MdepositSetMaxSupplyCap)
	if err := _Mdeposit.contract.UnpackLog(event, "SetMaxSupplyCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetMinAmountIterator is returned from FilterSetMinAmount and is used to iterate over the raw logs and unpacked data for SetMinAmount events raised by the Mdeposit contract.
type MdepositSetMinAmountIterator struct {
	Event *MdepositSetMinAmount // Event containing the contract specifics and raw log

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
func (it *MdepositSetMinAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetMinAmount)
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
		it.Event = new(MdepositSetMinAmount)
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
func (it *MdepositSetMinAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetMinAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetMinAmount represents a SetMinAmount event raised by the Mdeposit contract.
type MdepositSetMinAmount struct {
	Caller    common.Address
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetMinAmount is a free log retrieval operation binding the contract event 0x57e764c1fef224e74706b109734513889970db6f1dde107b1bda66e10d80ca9b.
//
// Solidity: event SetMinAmount(address indexed caller, uint256 newAmount)
func (_Mdeposit *MdepositFilterer) FilterSetMinAmount(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetMinAmountIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetMinAmount", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetMinAmountIterator{contract: _Mdeposit.contract, event: "SetMinAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinAmount is a free log subscription operation binding the contract event 0x57e764c1fef224e74706b109734513889970db6f1dde107b1bda66e10d80ca9b.
//
// Solidity: event SetMinAmount(address indexed caller, uint256 newAmount)
func (_Mdeposit *MdepositFilterer) WatchSetMinAmount(opts *bind.WatchOpts, sink chan<- *MdepositSetMinAmount, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetMinAmount", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetMinAmount)
				if err := _Mdeposit.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
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

// ParseSetMinAmount is a log parse operation binding the contract event 0x57e764c1fef224e74706b109734513889970db6f1dde107b1bda66e10d80ca9b.
//
// Solidity: event SetMinAmount(address indexed caller, uint256 newAmount)
func (_Mdeposit *MdepositFilterer) ParseSetMinAmount(log types.Log) (*MdepositSetMinAmount, error) {
	event := new(MdepositSetMinAmount)
	if err := _Mdeposit.contract.UnpackLog(event, "SetMinAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetMinMTokenAmountForFirstDepositIterator is returned from FilterSetMinMTokenAmountForFirstDeposit and is used to iterate over the raw logs and unpacked data for SetMinMTokenAmountForFirstDeposit events raised by the Mdeposit contract.
type MdepositSetMinMTokenAmountForFirstDepositIterator struct {
	Event *MdepositSetMinMTokenAmountForFirstDeposit // Event containing the contract specifics and raw log

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
func (it *MdepositSetMinMTokenAmountForFirstDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetMinMTokenAmountForFirstDeposit)
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
		it.Event = new(MdepositSetMinMTokenAmountForFirstDeposit)
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
func (it *MdepositSetMinMTokenAmountForFirstDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetMinMTokenAmountForFirstDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetMinMTokenAmountForFirstDeposit represents a SetMinMTokenAmountForFirstDeposit event raised by the Mdeposit contract.
type MdepositSetMinMTokenAmountForFirstDeposit struct {
	Caller   common.Address
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinMTokenAmountForFirstDeposit is a free log retrieval operation binding the contract event 0xf0af3ac3dc311b130ec783d7ff5582ccf0923fa13c4688c5da387d4cc57d852d.
//
// Solidity: event SetMinMTokenAmountForFirstDeposit(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) FilterSetMinMTokenAmountForFirstDeposit(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetMinMTokenAmountForFirstDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetMinMTokenAmountForFirstDeposit", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetMinMTokenAmountForFirstDepositIterator{contract: _Mdeposit.contract, event: "SetMinMTokenAmountForFirstDeposit", logs: logs, sub: sub}, nil
}

// WatchSetMinMTokenAmountForFirstDeposit is a free log subscription operation binding the contract event 0xf0af3ac3dc311b130ec783d7ff5582ccf0923fa13c4688c5da387d4cc57d852d.
//
// Solidity: event SetMinMTokenAmountForFirstDeposit(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) WatchSetMinMTokenAmountForFirstDeposit(opts *bind.WatchOpts, sink chan<- *MdepositSetMinMTokenAmountForFirstDeposit, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetMinMTokenAmountForFirstDeposit", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetMinMTokenAmountForFirstDeposit)
				if err := _Mdeposit.contract.UnpackLog(event, "SetMinMTokenAmountForFirstDeposit", log); err != nil {
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

// ParseSetMinMTokenAmountForFirstDeposit is a log parse operation binding the contract event 0xf0af3ac3dc311b130ec783d7ff5582ccf0923fa13c4688c5da387d4cc57d852d.
//
// Solidity: event SetMinMTokenAmountForFirstDeposit(address indexed caller, uint256 newValue)
func (_Mdeposit *MdepositFilterer) ParseSetMinMTokenAmountForFirstDeposit(log types.Log) (*MdepositSetMinMTokenAmountForFirstDeposit, error) {
	event := new(MdepositSetMinMTokenAmountForFirstDeposit)
	if err := _Mdeposit.contract.UnpackLog(event, "SetMinMTokenAmountForFirstDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetSanctionsListIterator is returned from FilterSetSanctionsList and is used to iterate over the raw logs and unpacked data for SetSanctionsList events raised by the Mdeposit contract.
type MdepositSetSanctionsListIterator struct {
	Event *MdepositSetSanctionsList // Event containing the contract specifics and raw log

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
func (it *MdepositSetSanctionsListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetSanctionsList)
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
		it.Event = new(MdepositSetSanctionsList)
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
func (it *MdepositSetSanctionsListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetSanctionsListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetSanctionsList represents a SetSanctionsList event raised by the Mdeposit contract.
type MdepositSetSanctionsList struct {
	Caller           common.Address
	NewSanctionsList common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetSanctionsList is a free log retrieval operation binding the contract event 0x7f0c791852a03e270d4c2b78bbd4b959bca234de8d1ccf27eee03afaeafe63c4.
//
// Solidity: event SetSanctionsList(address indexed caller, address indexed newSanctionsList)
func (_Mdeposit *MdepositFilterer) FilterSetSanctionsList(opts *bind.FilterOpts, caller []common.Address, newSanctionsList []common.Address) (*MdepositSetSanctionsListIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var newSanctionsListRule []interface{}
	for _, newSanctionsListItem := range newSanctionsList {
		newSanctionsListRule = append(newSanctionsListRule, newSanctionsListItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetSanctionsList", callerRule, newSanctionsListRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetSanctionsListIterator{contract: _Mdeposit.contract, event: "SetSanctionsList", logs: logs, sub: sub}, nil
}

// WatchSetSanctionsList is a free log subscription operation binding the contract event 0x7f0c791852a03e270d4c2b78bbd4b959bca234de8d1ccf27eee03afaeafe63c4.
//
// Solidity: event SetSanctionsList(address indexed caller, address indexed newSanctionsList)
func (_Mdeposit *MdepositFilterer) WatchSetSanctionsList(opts *bind.WatchOpts, sink chan<- *MdepositSetSanctionsList, caller []common.Address, newSanctionsList []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var newSanctionsListRule []interface{}
	for _, newSanctionsListItem := range newSanctionsList {
		newSanctionsListRule = append(newSanctionsListRule, newSanctionsListItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetSanctionsList", callerRule, newSanctionsListRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetSanctionsList)
				if err := _Mdeposit.contract.UnpackLog(event, "SetSanctionsList", log); err != nil {
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

// ParseSetSanctionsList is a log parse operation binding the contract event 0x7f0c791852a03e270d4c2b78bbd4b959bca234de8d1ccf27eee03afaeafe63c4.
//
// Solidity: event SetSanctionsList(address indexed caller, address indexed newSanctionsList)
func (_Mdeposit *MdepositFilterer) ParseSetSanctionsList(log types.Log) (*MdepositSetSanctionsList, error) {
	event := new(MdepositSetSanctionsList)
	if err := _Mdeposit.contract.UnpackLog(event, "SetSanctionsList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetTokensReceiverIterator is returned from FilterSetTokensReceiver and is used to iterate over the raw logs and unpacked data for SetTokensReceiver events raised by the Mdeposit contract.
type MdepositSetTokensReceiverIterator struct {
	Event *MdepositSetTokensReceiver // Event containing the contract specifics and raw log

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
func (it *MdepositSetTokensReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetTokensReceiver)
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
		it.Event = new(MdepositSetTokensReceiver)
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
func (it *MdepositSetTokensReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetTokensReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetTokensReceiver represents a SetTokensReceiver event raised by the Mdeposit contract.
type MdepositSetTokensReceiver struct {
	Caller   common.Address
	Reciever common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetTokensReceiver is a free log retrieval operation binding the contract event 0xdb5a411e1a379f981ff6bc5284aa2c2522a9b8fd33a9db9ca19b34006cefbe9c.
//
// Solidity: event SetTokensReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) FilterSetTokensReceiver(opts *bind.FilterOpts, caller []common.Address, reciever []common.Address) (*MdepositSetTokensReceiverIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetTokensReceiver", callerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetTokensReceiverIterator{contract: _Mdeposit.contract, event: "SetTokensReceiver", logs: logs, sub: sub}, nil
}

// WatchSetTokensReceiver is a free log subscription operation binding the contract event 0xdb5a411e1a379f981ff6bc5284aa2c2522a9b8fd33a9db9ca19b34006cefbe9c.
//
// Solidity: event SetTokensReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) WatchSetTokensReceiver(opts *bind.WatchOpts, sink chan<- *MdepositSetTokensReceiver, caller []common.Address, reciever []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetTokensReceiver", callerRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetTokensReceiver)
				if err := _Mdeposit.contract.UnpackLog(event, "SetTokensReceiver", log); err != nil {
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

// ParseSetTokensReceiver is a log parse operation binding the contract event 0xdb5a411e1a379f981ff6bc5284aa2c2522a9b8fd33a9db9ca19b34006cefbe9c.
//
// Solidity: event SetTokensReceiver(address indexed caller, address indexed reciever)
func (_Mdeposit *MdepositFilterer) ParseSetTokensReceiver(log types.Log) (*MdepositSetTokensReceiver, error) {
	event := new(MdepositSetTokensReceiver)
	if err := _Mdeposit.contract.UnpackLog(event, "SetTokensReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositSetVariationToleranceIterator is returned from FilterSetVariationTolerance and is used to iterate over the raw logs and unpacked data for SetVariationTolerance events raised by the Mdeposit contract.
type MdepositSetVariationToleranceIterator struct {
	Event *MdepositSetVariationTolerance // Event containing the contract specifics and raw log

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
func (it *MdepositSetVariationToleranceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositSetVariationTolerance)
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
		it.Event = new(MdepositSetVariationTolerance)
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
func (it *MdepositSetVariationToleranceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositSetVariationToleranceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositSetVariationTolerance represents a SetVariationTolerance event raised by the Mdeposit contract.
type MdepositSetVariationTolerance struct {
	Caller       common.Address
	NewTolerance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetVariationTolerance is a free log retrieval operation binding the contract event 0x018be394ba93a0dbca235443cfdc7173b2479180ad766083ce05199fbf3fc624.
//
// Solidity: event SetVariationTolerance(address indexed caller, uint256 newTolerance)
func (_Mdeposit *MdepositFilterer) FilterSetVariationTolerance(opts *bind.FilterOpts, caller []common.Address) (*MdepositSetVariationToleranceIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "SetVariationTolerance", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositSetVariationToleranceIterator{contract: _Mdeposit.contract, event: "SetVariationTolerance", logs: logs, sub: sub}, nil
}

// WatchSetVariationTolerance is a free log subscription operation binding the contract event 0x018be394ba93a0dbca235443cfdc7173b2479180ad766083ce05199fbf3fc624.
//
// Solidity: event SetVariationTolerance(address indexed caller, uint256 newTolerance)
func (_Mdeposit *MdepositFilterer) WatchSetVariationTolerance(opts *bind.WatchOpts, sink chan<- *MdepositSetVariationTolerance, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "SetVariationTolerance", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositSetVariationTolerance)
				if err := _Mdeposit.contract.UnpackLog(event, "SetVariationTolerance", log); err != nil {
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

// ParseSetVariationTolerance is a log parse operation binding the contract event 0x018be394ba93a0dbca235443cfdc7173b2479180ad766083ce05199fbf3fc624.
//
// Solidity: event SetVariationTolerance(address indexed caller, uint256 newTolerance)
func (_Mdeposit *MdepositFilterer) ParseSetVariationTolerance(log types.Log) (*MdepositSetVariationTolerance, error) {
	event := new(MdepositSetVariationTolerance)
	if err := _Mdeposit.contract.UnpackLog(event, "SetVariationTolerance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositUnpauseFnIterator is returned from FilterUnpauseFn and is used to iterate over the raw logs and unpacked data for UnpauseFn events raised by the Mdeposit contract.
type MdepositUnpauseFnIterator struct {
	Event *MdepositUnpauseFn // Event containing the contract specifics and raw log

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
func (it *MdepositUnpauseFnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositUnpauseFn)
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
		it.Event = new(MdepositUnpauseFn)
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
func (it *MdepositUnpauseFnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositUnpauseFnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositUnpauseFn represents a UnpauseFn event raised by the Mdeposit contract.
type MdepositUnpauseFn struct {
	Caller common.Address
	Fn     [4]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpauseFn is a free log retrieval operation binding the contract event 0x929135cc6324f958693bb5f24a4dbc226a83c721523fc2785545019a3423b2d7.
//
// Solidity: event UnpauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) FilterUnpauseFn(opts *bind.FilterOpts, caller []common.Address) (*MdepositUnpauseFnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "UnpauseFn", callerRule)
	if err != nil {
		return nil, err
	}
	return &MdepositUnpauseFnIterator{contract: _Mdeposit.contract, event: "UnpauseFn", logs: logs, sub: sub}, nil
}

// WatchUnpauseFn is a free log subscription operation binding the contract event 0x929135cc6324f958693bb5f24a4dbc226a83c721523fc2785545019a3423b2d7.
//
// Solidity: event UnpauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) WatchUnpauseFn(opts *bind.WatchOpts, sink chan<- *MdepositUnpauseFn, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "UnpauseFn", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositUnpauseFn)
				if err := _Mdeposit.contract.UnpackLog(event, "UnpauseFn", log); err != nil {
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

// ParseUnpauseFn is a log parse operation binding the contract event 0x929135cc6324f958693bb5f24a4dbc226a83c721523fc2785545019a3423b2d7.
//
// Solidity: event UnpauseFn(address indexed caller, bytes4 fn)
func (_Mdeposit *MdepositFilterer) ParseUnpauseFn(log types.Log) (*MdepositUnpauseFn, error) {
	event := new(MdepositUnpauseFn)
	if err := _Mdeposit.contract.UnpackLog(event, "UnpauseFn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Mdeposit contract.
type MdepositUnpausedIterator struct {
	Event *MdepositUnpaused // Event containing the contract specifics and raw log

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
func (it *MdepositUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositUnpaused)
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
		it.Event = new(MdepositUnpaused)
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
func (it *MdepositUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositUnpaused represents a Unpaused event raised by the Mdeposit contract.
type MdepositUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Mdeposit *MdepositFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MdepositUnpausedIterator, error) {

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MdepositUnpausedIterator{contract: _Mdeposit.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Mdeposit *MdepositFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MdepositUnpaused) (event.Subscription, error) {

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositUnpaused)
				if err := _Mdeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Mdeposit *MdepositFilterer) ParseUnpaused(log types.Log) (*MdepositUnpaused, error) {
	event := new(MdepositUnpaused)
	if err := _Mdeposit.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MdepositWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the Mdeposit contract.
type MdepositWithdrawTokenIterator struct {
	Event *MdepositWithdrawToken // Event containing the contract specifics and raw log

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
func (it *MdepositWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MdepositWithdrawToken)
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
		it.Event = new(MdepositWithdrawToken)
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
func (it *MdepositWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MdepositWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MdepositWithdrawToken represents a WithdrawToken event raised by the Mdeposit contract.
type MdepositWithdrawToken struct {
	Caller     common.Address
	Token      common.Address
	WithdrawTo common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed caller, address indexed token, address indexed withdrawTo, uint256 amount)
func (_Mdeposit *MdepositFilterer) FilterWithdrawToken(opts *bind.FilterOpts, caller []common.Address, token []common.Address, withdrawTo []common.Address) (*MdepositWithdrawTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawToRule []interface{}
	for _, withdrawToItem := range withdrawTo {
		withdrawToRule = append(withdrawToRule, withdrawToItem)
	}

	logs, sub, err := _Mdeposit.contract.FilterLogs(opts, "WithdrawToken", callerRule, tokenRule, withdrawToRule)
	if err != nil {
		return nil, err
	}
	return &MdepositWithdrawTokenIterator{contract: _Mdeposit.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed caller, address indexed token, address indexed withdrawTo, uint256 amount)
func (_Mdeposit *MdepositFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *MdepositWithdrawToken, caller []common.Address, token []common.Address, withdrawTo []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawToRule []interface{}
	for _, withdrawToItem := range withdrawTo {
		withdrawToRule = append(withdrawToRule, withdrawToItem)
	}

	logs, sub, err := _Mdeposit.contract.WatchLogs(opts, "WithdrawToken", callerRule, tokenRule, withdrawToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MdepositWithdrawToken)
				if err := _Mdeposit.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed caller, address indexed token, address indexed withdrawTo, uint256 amount)
func (_Mdeposit *MdepositFilterer) ParseWithdrawToken(log types.Log) (*MdepositWithdrawToken, error) {
	event := new(MdepositWithdrawToken)
	if err := _Mdeposit.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
