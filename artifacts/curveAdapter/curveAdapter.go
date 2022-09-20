// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curveAdapter

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

// CurveAdapterMetaData contains all meta data concerning the CurveAdapter contract.
var CurveAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metapoolBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_previous_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_price_cumulative_last\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_first_balances\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_last_balances\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_time_elapsed\",\"type\":\"uint256\"}],\"name\":\"get_twap_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metapoolBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balances\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_previous_balances\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_price_cumulative_last\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_first_balances\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_last_balances\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"_time_elapsed\",\"type\":\"uint256\"}],\"name\":\"get_twap_balances\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[3]\",\"name\":\"\",\"type\":\"uint256[3]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metapoolBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balances\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_previous_balances\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_price_cumulative_last\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_first_balances\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_last_balances\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"_time_elapsed\",\"type\":\"uint256\"}],\"name\":\"get_twap_balances\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"amounts\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curvePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metapoolBase\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nCoins\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curveDeposit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nCoins\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creditManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_curveStETHPoolGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectIndexException\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TokenIsNotInAllowedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterType\",\"outputs\":[{\"internalType\":\"enumAdapterType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_gearboxAdapterVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"add_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_add_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditFacade\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditManager\",\"outputs\":[{\"internalType\":\"contractICreditManagerV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"rateMinRAY\",\"type\":\"uint256\"}],\"name\":\"exchange_all_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_previous_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_price_cumulative_last\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_first_balances\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_last_balances\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_time_elapsed\",\"type\":\"uint256\"}],\"name\":\"get_twap_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metapoolBase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nCoins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minRateRAY\",\"type\":\"uint256\"}],\"name\":\"remove_all_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_steth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotImplementedException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressException\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"A_precise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"admin_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"block_timestamp_last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"bool\",\"name\":\"_is_deposit\",\"type\":\"bool\"}],\"name\":\"calc_token_amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burn_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"calc_withdraw_one_coin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min_dy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchange_underlying\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"get_dy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"get_dy_underlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_previous_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_price_cumulative_last\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"get_twap_balances\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_virtual_price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initial_A_time\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lp_token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"min_amounts\",\"type\":\"uint256[2]\"}],\"name\":\"remove_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"amounts\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"max_burn_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_imbalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token_amount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"min_amount\",\"type\":\"uint256\"}],\"name\":\"remove_liquidity_one_coin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"underlying_coins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CurveAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveAdapterMetaData.ABI instead.
var CurveAdapterABI = CurveAdapterMetaData.ABI

// CurveAdapter is an auto generated Go binding around an Ethereum contract.
type CurveAdapter struct {
	CurveAdapterCaller     // Read-only binding to the contract
	CurveAdapterTransactor // Write-only binding to the contract
	CurveAdapterFilterer   // Log filterer for contract events
}

// CurveAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveAdapterSession struct {
	Contract     *CurveAdapter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveAdapterCallerSession struct {
	Contract *CurveAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CurveAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveAdapterTransactorSession struct {
	Contract     *CurveAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CurveAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveAdapterRaw struct {
	Contract *CurveAdapter // Generic contract binding to access the raw methods on
}

// CurveAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveAdapterCallerRaw struct {
	Contract *CurveAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// CurveAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveAdapterTransactorRaw struct {
	Contract *CurveAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveAdapter creates a new instance of CurveAdapter, bound to a specific deployed contract.
func NewCurveAdapter(address common.Address, backend bind.ContractBackend) (*CurveAdapter, error) {
	contract, err := bindCurveAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveAdapter{CurveAdapterCaller: CurveAdapterCaller{contract: contract}, CurveAdapterTransactor: CurveAdapterTransactor{contract: contract}, CurveAdapterFilterer: CurveAdapterFilterer{contract: contract}}, nil
}

// NewCurveAdapterCaller creates a new read-only instance of CurveAdapter, bound to a specific deployed contract.
func NewCurveAdapterCaller(address common.Address, caller bind.ContractCaller) (*CurveAdapterCaller, error) {
	contract, err := bindCurveAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAdapterCaller{contract: contract}, nil
}

// NewCurveAdapterTransactor creates a new write-only instance of CurveAdapter, bound to a specific deployed contract.
func NewCurveAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveAdapterTransactor, error) {
	contract, err := bindCurveAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveAdapterTransactor{contract: contract}, nil
}

// NewCurveAdapterFilterer creates a new log filterer instance of CurveAdapter, bound to a specific deployed contract.
func NewCurveAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveAdapterFilterer, error) {
	contract, err := bindCurveAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveAdapterFilterer{contract: contract}, nil
}

// bindCurveAdapter binds a generic wrapper to an already deployed contract.
func bindCurveAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAdapter *CurveAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAdapter.Contract.CurveAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAdapter *CurveAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAdapter.Contract.CurveAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAdapter *CurveAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAdapter.Contract.CurveAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveAdapter *CurveAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveAdapter *CurveAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveAdapter *CurveAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveAdapter.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A() (*big.Int, error) {
	return _CurveAdapter.Contract.A(&_CurveAdapter.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A() (*big.Int, error) {
	return _CurveAdapter.Contract.A(&_CurveAdapter.CallOpts)
}

// A0 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A0 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A0() (*big.Int, error) {
	return _CurveAdapter.Contract.A0(&_CurveAdapter.CallOpts)
}

// A0 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A0() (*big.Int, error) {
	return _CurveAdapter.Contract.A0(&_CurveAdapter.CallOpts)
}

// A1 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A1 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A1() (*big.Int, error) {
	return _CurveAdapter.Contract.A1(&_CurveAdapter.CallOpts)
}

// A1 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A1() (*big.Int, error) {
	return _CurveAdapter.Contract.A1(&_CurveAdapter.CallOpts)
}

// A2 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A2 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A2() (*big.Int, error) {
	return _CurveAdapter.Contract.A2(&_CurveAdapter.CallOpts)
}

// A2 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A2() (*big.Int, error) {
	return _CurveAdapter.Contract.A2(&_CurveAdapter.CallOpts)
}

// A3 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A3 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A3() (*big.Int, error) {
	return _CurveAdapter.Contract.A3(&_CurveAdapter.CallOpts)
}

// A3 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A3() (*big.Int, error) {
	return _CurveAdapter.Contract.A3(&_CurveAdapter.CallOpts)
}

// A4 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A4 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A4() (*big.Int, error) {
	return _CurveAdapter.Contract.A4(&_CurveAdapter.CallOpts)
}

// A4 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A4() (*big.Int, error) {
	return _CurveAdapter.Contract.A4(&_CurveAdapter.CallOpts)
}

// A5 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) A5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A5 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) A5() (*big.Int, error) {
	return _CurveAdapter.Contract.A5(&_CurveAdapter.CallOpts)
}

// A5 is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) A5() (*big.Int, error) {
	return _CurveAdapter.Contract.A5(&_CurveAdapter.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise(&_CurveAdapter.CallOpts)
}

// APrecise is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise(&_CurveAdapter.CallOpts)
}

// APrecise0 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise0 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise0() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise0(&_CurveAdapter.CallOpts)
}

// APrecise0 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise0() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise0(&_CurveAdapter.CallOpts)
}

// APrecise1 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise1 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise1() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise1(&_CurveAdapter.CallOpts)
}

// APrecise1 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise1() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise1(&_CurveAdapter.CallOpts)
}

// APrecise2 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise2 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise2() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise2(&_CurveAdapter.CallOpts)
}

// APrecise2 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise2() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise2(&_CurveAdapter.CallOpts)
}

// APrecise3 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise3 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise3() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise3(&_CurveAdapter.CallOpts)
}

// APrecise3 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise3() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise3(&_CurveAdapter.CallOpts)
}

// APrecise4 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise4 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise4() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise4(&_CurveAdapter.CallOpts)
}

// APrecise4 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise4() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise4(&_CurveAdapter.CallOpts)
}

// APrecise5 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) APrecise5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "A_precise5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APrecise5 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) APrecise5() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise5(&_CurveAdapter.CallOpts)
}

// APrecise5 is a free data retrieval call binding the contract method 0x76a2f0f0.
//
// Solidity: function A_precise() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) APrecise5() (*big.Int, error) {
	return _CurveAdapter.Contract.APrecise5(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType0(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType0")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType0() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType0(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType0 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType0() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType0(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType1(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType1")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType1() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType1(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType1 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType1() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType1(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType2 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType2(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType2")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType2 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType2() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType2(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType2 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType2() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType2(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType3 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType3(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType3")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType3 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType3() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType3(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType3 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() view returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType3() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType3(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType4 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterType4(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterType4")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GearboxAdapterType4 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterType4() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType4(&_CurveAdapter.CallOpts)
}

// GearboxAdapterType4 is a free data retrieval call binding the contract method 0xce30bbdb.
//
// Solidity: function _gearboxAdapterType() pure returns(uint8)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterType4() (uint8, error) {
	return _CurveAdapter.Contract.GearboxAdapterType4(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion0(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion0")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion0() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion0(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion0 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion0() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion0(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion1(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion1")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion1() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion1(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion1 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion1() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion1(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion2 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion2(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion2")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion2 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion2() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion2(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion2 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion2() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion2(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion3 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion3(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion3")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion3 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion3() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion3(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion3 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion3() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion3(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion4 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCaller) GearboxAdapterVersion4(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "_gearboxAdapterVersion4")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GearboxAdapterVersion4 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterSession) GearboxAdapterVersion4() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion4(&_CurveAdapter.CallOpts)
}

// GearboxAdapterVersion4 is a free data retrieval call binding the contract method 0x78aa73a4.
//
// Solidity: function _gearboxAdapterVersion() view returns(uint16)
func (_CurveAdapter *CurveAdapterCallerSession) GearboxAdapterVersion4() (uint16, error) {
	return _CurveAdapter.Contract.GearboxAdapterVersion4(&_CurveAdapter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin() (common.Address, error) {
	return _CurveAdapter.Contract.Admin(&_CurveAdapter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin() (common.Address, error) {
	return _CurveAdapter.Contract.Admin(&_CurveAdapter.CallOpts)
}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin0() (common.Address, error) {
	return _CurveAdapter.Contract.Admin0(&_CurveAdapter.CallOpts)
}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin0() (common.Address, error) {
	return _CurveAdapter.Contract.Admin0(&_CurveAdapter.CallOpts)
}

// Admin1 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin1 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin1() (common.Address, error) {
	return _CurveAdapter.Contract.Admin1(&_CurveAdapter.CallOpts)
}

// Admin1 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin1() (common.Address, error) {
	return _CurveAdapter.Contract.Admin1(&_CurveAdapter.CallOpts)
}

// Admin2 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin2 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin2() (common.Address, error) {
	return _CurveAdapter.Contract.Admin2(&_CurveAdapter.CallOpts)
}

// Admin2 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin2() (common.Address, error) {
	return _CurveAdapter.Contract.Admin2(&_CurveAdapter.CallOpts)
}

// Admin3 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin3 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin3() (common.Address, error) {
	return _CurveAdapter.Contract.Admin3(&_CurveAdapter.CallOpts)
}

// Admin3 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin3() (common.Address, error) {
	return _CurveAdapter.Contract.Admin3(&_CurveAdapter.CallOpts)
}

// Admin4 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin4 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin4() (common.Address, error) {
	return _CurveAdapter.Contract.Admin4(&_CurveAdapter.CallOpts)
}

// Admin4 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin4() (common.Address, error) {
	return _CurveAdapter.Contract.Admin4(&_CurveAdapter.CallOpts)
}

// Admin5 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Admin5(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin5")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin5 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Admin5() (common.Address, error) {
	return _CurveAdapter.Contract.Admin5(&_CurveAdapter.CallOpts)
}

// Admin5 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Admin5() (common.Address, error) {
	return _CurveAdapter.Contract.Admin5(&_CurveAdapter.CallOpts)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances(&_CurveAdapter.CallOpts, i)
}

// AdminBalances is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances(&_CurveAdapter.CallOpts, i)
}

// AdminBalances0 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances0(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances0", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances0 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances0(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances0(&_CurveAdapter.CallOpts, i)
}

// AdminBalances0 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances0(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances0(&_CurveAdapter.CallOpts, i)
}

// AdminBalances1 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances1(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances1", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances1 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances1(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances1(&_CurveAdapter.CallOpts, i)
}

// AdminBalances1 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances1(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances1(&_CurveAdapter.CallOpts, i)
}

// AdminBalances2 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances2(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances2", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances2 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances2(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances2(&_CurveAdapter.CallOpts, i)
}

// AdminBalances2 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances2(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances2(&_CurveAdapter.CallOpts, i)
}

// AdminBalances3 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances3(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances3", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances3 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances3(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances3(&_CurveAdapter.CallOpts, i)
}

// AdminBalances3 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances3(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances3(&_CurveAdapter.CallOpts, i)
}

// AdminBalances4 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances4(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances4", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances4 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances4(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances4(&_CurveAdapter.CallOpts, i)
}

// AdminBalances4 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances4(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances4(&_CurveAdapter.CallOpts, i)
}

// AdminBalances5 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminBalances5(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_balances5", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminBalances5 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminBalances5(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances5(&_CurveAdapter.CallOpts, i)
}

// AdminBalances5 is a free data retrieval call binding the contract method 0xe2e7d264.
//
// Solidity: function admin_balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminBalances5(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.AdminBalances5(&_CurveAdapter.CallOpts, i)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee(&_CurveAdapter.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee(&_CurveAdapter.CallOpts)
}

// AdminFee0 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee0 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee0() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee0(&_CurveAdapter.CallOpts)
}

// AdminFee0 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee0() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee0(&_CurveAdapter.CallOpts)
}

// AdminFee1 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee1 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee1() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee1(&_CurveAdapter.CallOpts)
}

// AdminFee1 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee1() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee1(&_CurveAdapter.CallOpts)
}

// AdminFee2 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee2 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee2() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee2(&_CurveAdapter.CallOpts)
}

// AdminFee2 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee2() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee2(&_CurveAdapter.CallOpts)
}

// AdminFee3 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee3 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee3() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee3(&_CurveAdapter.CallOpts)
}

// AdminFee3 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee3() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee3(&_CurveAdapter.CallOpts)
}

// AdminFee4 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee4 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee4() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee4(&_CurveAdapter.CallOpts)
}

// AdminFee4 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee4() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee4(&_CurveAdapter.CallOpts)
}

// AdminFee5 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) AdminFee5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "admin_fee5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee5 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) AdminFee5() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee5(&_CurveAdapter.CallOpts)
}

// AdminFee5 is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) AdminFee5() (*big.Int, error) {
	return _CurveAdapter.Contract.AdminFee5(&_CurveAdapter.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance0(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance0", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance0(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance0(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance0 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance0(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance0(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance1(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance1", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance1(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance1(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance1 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance1(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance1(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance2 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance2(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance2", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance2 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance2(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance2(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance2 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance2(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance2(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance3 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance3(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance3", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance3 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance3(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance3(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance3 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance3(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance3(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance4 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance4(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance4", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance4 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance4(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance4(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance4 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance4(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance4(&_CurveAdapter.CallOpts, owner, spender)
}

// Allowance5 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Allowance5(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "allowance5", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance5 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Allowance5(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance5(&_CurveAdapter.CallOpts, arg0, arg1)
}

// Allowance5 is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Allowance5(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.Allowance5(&_CurveAdapter.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf(&_CurveAdapter.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf(&_CurveAdapter.CallOpts, account)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf0(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf0(&_CurveAdapter.CallOpts, account)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf0(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf0(&_CurveAdapter.CallOpts, account)
}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf1(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf1", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf1(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf1(&_CurveAdapter.CallOpts, account)
}

// BalanceOf1 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf1(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf1(&_CurveAdapter.CallOpts, account)
}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf2(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf2", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf2(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf2(&_CurveAdapter.CallOpts, account)
}

// BalanceOf2 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf2(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf2(&_CurveAdapter.CallOpts, account)
}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf3(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf3", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf3(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf3(&_CurveAdapter.CallOpts, account)
}

// BalanceOf3 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf3(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf3(&_CurveAdapter.CallOpts, account)
}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf4(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf4", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf4(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf4(&_CurveAdapter.CallOpts, account)
}

// BalanceOf4 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf4(account common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf4(&_CurveAdapter.CallOpts, account)
}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BalanceOf5(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balanceOf5", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BalanceOf5(arg0 common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf5(&_CurveAdapter.CallOpts, arg0)
}

// BalanceOf5 is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BalanceOf5(arg0 common.Address) (*big.Int, error) {
	return _CurveAdapter.Contract.BalanceOf5(&_CurveAdapter.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances(&_CurveAdapter.CallOpts, i)
}

// Balances is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances(&_CurveAdapter.CallOpts, i)
}

// Balances0 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances0(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances0", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances0 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances0(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances0(&_CurveAdapter.CallOpts, i)
}

// Balances0 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances0(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances0(&_CurveAdapter.CallOpts, i)
}

// Balances1 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances1(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances1", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances1 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances1(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances1(&_CurveAdapter.CallOpts, i)
}

// Balances1 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances1(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances1(&_CurveAdapter.CallOpts, i)
}

// Balances10 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances10(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances10", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances10 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances10(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances10(&_CurveAdapter.CallOpts, i)
}

// Balances10 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances10(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances10(&_CurveAdapter.CallOpts, i)
}

// Balances11 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances11(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances11", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances11 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances11(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances11(&_CurveAdapter.CallOpts, i)
}

// Balances11 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances11(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances11(&_CurveAdapter.CallOpts, i)
}

// Balances12 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances12(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances12", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances12 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances12(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances12(&_CurveAdapter.CallOpts, i)
}

// Balances12 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances12(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances12(&_CurveAdapter.CallOpts, i)
}

// Balances2 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances2(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances2", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances2 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances2(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances2(&_CurveAdapter.CallOpts, i)
}

// Balances2 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances2(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances2(&_CurveAdapter.CallOpts, i)
}

// Balances3 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances3(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances3", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances3 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances3(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances3(&_CurveAdapter.CallOpts, i)
}

// Balances3 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances3(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances3(&_CurveAdapter.CallOpts, i)
}

// Balances4 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances4(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances4", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances4 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances4(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances4(&_CurveAdapter.CallOpts, i)
}

// Balances4 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances4(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances4(&_CurveAdapter.CallOpts, i)
}

// Balances5 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances5(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances5", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances5 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances5(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances5(&_CurveAdapter.CallOpts, i)
}

// Balances5 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances5(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances5(&_CurveAdapter.CallOpts, i)
}

// Balances6 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances6(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances6", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances6 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances6(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances6(&_CurveAdapter.CallOpts, i)
}

// Balances6 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances6(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances6(&_CurveAdapter.CallOpts, i)
}

// Balances7 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances7(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances7", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances7 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances7(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances7(&_CurveAdapter.CallOpts, i)
}

// Balances7 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances7(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances7(&_CurveAdapter.CallOpts, i)
}

// Balances8 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances8(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances8", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances8 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances8(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances8(&_CurveAdapter.CallOpts, i)
}

// Balances8 is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances8(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances8(&_CurveAdapter.CallOpts, i)
}

// Balances9 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Balances9(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "balances9", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances9 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Balances9(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances9(&_CurveAdapter.CallOpts, i)
}

// Balances9 is a free data retrieval call binding the contract method 0x065a80d8.
//
// Solidity: function balances(int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Balances9(i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.Balances9(&_CurveAdapter.CallOpts, i)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast0 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast0 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast0() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast0(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast0 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast0() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast0(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast1 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast1 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast1() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast1(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast1 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast1() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast1(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast2 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast2 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast2() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast2(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast2 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast2() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast2(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast3 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast3 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast3() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast3(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast3 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast3() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast3(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast4 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast4 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast4() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast4(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast4 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast4() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast4(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast5 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) BlockTimestampLast5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "block_timestamp_last5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTimestampLast5 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) BlockTimestampLast5() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast5(&_CurveAdapter.CallOpts)
}

// BlockTimestampLast5 is a free data retrieval call binding the contract method 0x63543f06.
//
// Solidity: function block_timestamp_last() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) BlockTimestampLast5() (*big.Int, error) {
	return _CurveAdapter.Contract.BlockTimestampLast5(&_CurveAdapter.CallOpts)
}

// CalcAddOneCoin is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin0 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin0(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin0", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin0 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin0(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin0(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin0 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin0(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin0(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin1 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin1(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin1", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin1 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin1(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin1(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin1 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin1(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin1(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin2 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin2(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin2", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin2 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin2(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin2(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin2 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin2(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin2(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin3 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin3(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin3", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin3 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin3(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin3(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin3 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin3(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin3(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin4 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcAddOneCoin4(opts *bind.CallOpts, amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_add_one_coin4", amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcAddOneCoin4 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcAddOneCoin4(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin4(&_CurveAdapter.CallOpts, amount, i)
}

// CalcAddOneCoin4 is a free data retrieval call binding the contract method 0x2c5788d2.
//
// Solidity: function calc_add_one_coin(uint256 amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcAddOneCoin4(amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcAddOneCoin4(&_CurveAdapter.CallOpts, amount, i)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcTokenAmount(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_token_amount", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcTokenAmount(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcTokenAmount0(opts *bind.CallOpts, _amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_token_amount0", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcTokenAmount0(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount0(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount0 is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcTokenAmount0(_amounts [3]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount0(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcTokenAmount1(opts *bind.CallOpts, _amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_token_amount1", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcTokenAmount1(_amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount1(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount1 is a free data retrieval call binding the contract method 0xcf701ff7.
//
// Solidity: function calc_token_amount(uint256[4] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcTokenAmount1(_amounts [4]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount1(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcTokenAmount2(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_token_amount2", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcTokenAmount2(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount2(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount2 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcTokenAmount2(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount2(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount3 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcTokenAmount3(opts *bind.CallOpts, _amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_token_amount3", _amounts, _is_deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount3 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcTokenAmount3(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount3(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcTokenAmount3 is a free data retrieval call binding the contract method 0xed8e84f3.
//
// Solidity: function calc_token_amount(uint256[2] _amounts, bool _is_deposit) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcTokenAmount3(_amounts [2]*big.Int, _is_deposit bool) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcTokenAmount3(&_CurveAdapter.CallOpts, _amounts, _is_deposit)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin0 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin0(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin0", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin0 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin0(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin0(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin0 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin0(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin0(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin1 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin1(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin1", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin1 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin1(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin1(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin1 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin1(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin1(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin2 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin2(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin2", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin2 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin2(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin2(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin2 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin2(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin2(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin3 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin3(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin3", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin3 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin3(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin3(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin3 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin3(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin3(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin4 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin4(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin4", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin4 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin4(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin4(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin4 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin4(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin4(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin5 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) CalcWithdrawOneCoin5(opts *bind.CallOpts, _burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "calc_withdraw_one_coin5", _burn_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin5 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) CalcWithdrawOneCoin5(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin5(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// CalcWithdrawOneCoin5 is a free data retrieval call binding the contract method 0xcc2b27d7.
//
// Solidity: function calc_withdraw_one_coin(uint256 _burn_amount, int128 i) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) CalcWithdrawOneCoin5(_burn_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.CalcWithdrawOneCoin5(&_CurveAdapter.CallOpts, _burn_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins(&_CurveAdapter.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins(&_CurveAdapter.CallOpts, i)
}

// Coins0 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins0(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins0", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins0 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins0(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins0(&_CurveAdapter.CallOpts, i)
}

// Coins0 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins0(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins0(&_CurveAdapter.CallOpts, i)
}

// Coins1 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins1(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins1", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins1 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins1(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins1(&_CurveAdapter.CallOpts, i)
}

// Coins1 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins1(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins1(&_CurveAdapter.CallOpts, i)
}

// Coins10 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins10(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins10", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins10 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins10(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins10(&_CurveAdapter.CallOpts, i)
}

// Coins10 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins10(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins10(&_CurveAdapter.CallOpts, i)
}

// Coins11 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins11(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins11", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins11 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins11(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins11(&_CurveAdapter.CallOpts, i)
}

// Coins11 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins11(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins11(&_CurveAdapter.CallOpts, i)
}

// Coins12 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins12(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins12", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins12 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins12(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins12(&_CurveAdapter.CallOpts, i)
}

// Coins12 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins12(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins12(&_CurveAdapter.CallOpts, i)
}

// Coins2 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins2(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins2", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins2 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins2(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins2(&_CurveAdapter.CallOpts, i)
}

// Coins2 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins2(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins2(&_CurveAdapter.CallOpts, i)
}

// Coins3 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins3(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins3", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins3 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins3(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins3(&_CurveAdapter.CallOpts, i)
}

// Coins3 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins3(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins3(&_CurveAdapter.CallOpts, i)
}

// Coins4 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins4(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins4", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins4 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins4(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins4(&_CurveAdapter.CallOpts, i)
}

// Coins4 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins4(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins4(&_CurveAdapter.CallOpts, i)
}

// Coins5 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins5(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins5", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins5 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins5(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins5(&_CurveAdapter.CallOpts, i)
}

// Coins5 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins5(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins5(&_CurveAdapter.CallOpts, i)
}

// Coins6 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins6(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins6", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins6 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins6(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins6(&_CurveAdapter.CallOpts, i)
}

// Coins6 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins6(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins6(&_CurveAdapter.CallOpts, i)
}

// Coins7 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins7(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins7", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins7 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins7(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins7(&_CurveAdapter.CallOpts, i)
}

// Coins7 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins7(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins7(&_CurveAdapter.CallOpts, i)
}

// Coins8 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins8(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins8", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins8 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins8(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins8(&_CurveAdapter.CallOpts, i)
}

// Coins8 is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins8(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins8(&_CurveAdapter.CallOpts, i)
}

// Coins9 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Coins9(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "coins9", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins9 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) Coins9(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins9(&_CurveAdapter.CallOpts, i)
}

// Coins9 is a free data retrieval call binding the contract method 0x23746eb8.
//
// Solidity: function coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Coins9(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.Coins9(&_CurveAdapter.CallOpts, i)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade(&_CurveAdapter.CallOpts)
}

// CreditFacade is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade(&_CurveAdapter.CallOpts)
}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade0() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade0(&_CurveAdapter.CallOpts)
}

// CreditFacade0 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade0() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade0(&_CurveAdapter.CallOpts)
}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade1() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade1(&_CurveAdapter.CallOpts)
}

// CreditFacade1 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade1() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade1(&_CurveAdapter.CallOpts)
}

// CreditFacade2 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade2 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade2() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade2(&_CurveAdapter.CallOpts)
}

// CreditFacade2 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade2() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade2(&_CurveAdapter.CallOpts)
}

// CreditFacade3 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade3 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade3() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade3(&_CurveAdapter.CallOpts)
}

// CreditFacade3 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade3() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade3(&_CurveAdapter.CallOpts)
}

// CreditFacade4 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditFacade4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditFacade4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditFacade4 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditFacade4() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade4(&_CurveAdapter.CallOpts)
}

// CreditFacade4 is a free data retrieval call binding the contract method 0x2f7a1881.
//
// Solidity: function creditFacade() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditFacade4() (common.Address, error) {
	return _CurveAdapter.Contract.CreditFacade4(&_CurveAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager(&_CurveAdapter.CallOpts)
}

// CreditManager is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager(&_CurveAdapter.CallOpts)
}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager0() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager0(&_CurveAdapter.CallOpts)
}

// CreditManager0 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager0() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager0(&_CurveAdapter.CallOpts)
}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager1() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager1(&_CurveAdapter.CallOpts)
}

// CreditManager1 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager1() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager1(&_CurveAdapter.CallOpts)
}

// CreditManager2 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager2 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager2() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager2(&_CurveAdapter.CallOpts)
}

// CreditManager2 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager2() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager2(&_CurveAdapter.CallOpts)
}

// CreditManager3 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager3 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager3() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager3(&_CurveAdapter.CallOpts)
}

// CreditManager3 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager3() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager3(&_CurveAdapter.CallOpts)
}

// CreditManager4 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) CreditManager4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "creditManager4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditManager4 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterSession) CreditManager4() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager4(&_CurveAdapter.CallOpts)
}

// CreditManager4 is a free data retrieval call binding the contract method 0xc12c21c0.
//
// Solidity: function creditManager() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) CreditManager4() (common.Address, error) {
	return _CurveAdapter.Contract.CreditManager4(&_CurveAdapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals(&_CurveAdapter.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals(&_CurveAdapter.CallOpts)
}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals0() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals0(&_CurveAdapter.CallOpts)
}

// Decimals0 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals0() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals0(&_CurveAdapter.CallOpts)
}

// Decimals1 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals1 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals1() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals1(&_CurveAdapter.CallOpts)
}

// Decimals1 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals1() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals1(&_CurveAdapter.CallOpts)
}

// Decimals2 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals2 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals2() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals2(&_CurveAdapter.CallOpts)
}

// Decimals2 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals2() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals2(&_CurveAdapter.CallOpts)
}

// Decimals3 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals3 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals3() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals3(&_CurveAdapter.CallOpts)
}

// Decimals3 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals3() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals3(&_CurveAdapter.CallOpts)
}

// Decimals4 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals4 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals4() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals4(&_CurveAdapter.CallOpts)
}

// Decimals4 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals4() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals4(&_CurveAdapter.CallOpts)
}

// Decimals5 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Decimals5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "decimals5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals5 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Decimals5() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals5(&_CurveAdapter.CallOpts)
}

// Decimals5 is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Decimals5() (*big.Int, error) {
	return _CurveAdapter.Contract.Decimals5(&_CurveAdapter.CallOpts)
}

// ExchangeUnderlying5 is a free data retrieval call binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 , int128 , uint256 , uint256 ) pure returns()
func (_CurveAdapter *CurveAdapterCaller) ExchangeUnderlying5(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int) error {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "exchange_underlying5", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// ExchangeUnderlying5 is a free data retrieval call binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 , int128 , uint256 , uint256 ) pure returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying5(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int) error {
	return _CurveAdapter.Contract.ExchangeUnderlying5(&_CurveAdapter.CallOpts, arg0, arg1, arg2, arg3)
}

// ExchangeUnderlying5 is a free data retrieval call binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 , int128 , uint256 , uint256 ) pure returns()
func (_CurveAdapter *CurveAdapterCallerSession) ExchangeUnderlying5(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int) error {
	return _CurveAdapter.Contract.ExchangeUnderlying5(&_CurveAdapter.CallOpts, arg0, arg1, arg2, arg3)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee(&_CurveAdapter.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee(&_CurveAdapter.CallOpts)
}

// Fee0 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee0 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee0() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee0(&_CurveAdapter.CallOpts)
}

// Fee0 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee0() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee0(&_CurveAdapter.CallOpts)
}

// Fee1 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee1 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee1() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee1(&_CurveAdapter.CallOpts)
}

// Fee1 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee1() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee1(&_CurveAdapter.CallOpts)
}

// Fee2 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee2 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee2() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee2(&_CurveAdapter.CallOpts)
}

// Fee2 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee2() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee2(&_CurveAdapter.CallOpts)
}

// Fee3 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee3 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee3() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee3(&_CurveAdapter.CallOpts)
}

// Fee3 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee3() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee3(&_CurveAdapter.CallOpts)
}

// Fee4 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee4 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee4() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee4(&_CurveAdapter.CallOpts)
}

// Fee4 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee4() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee4(&_CurveAdapter.CallOpts)
}

// Fee5 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) Fee5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "fee5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee5 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) Fee5() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee5(&_CurveAdapter.CallOpts)
}

// Fee5 is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) Fee5() (*big.Int, error) {
	return _CurveAdapter.Contract.Fee5(&_CurveAdapter.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA(&_CurveAdapter.CallOpts)
}

// FutureA is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA(&_CurveAdapter.CallOpts)
}

// FutureA0 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA0 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA0() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA0(&_CurveAdapter.CallOpts)
}

// FutureA0 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA0() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA0(&_CurveAdapter.CallOpts)
}

// FutureA1 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA1 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA1() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA1(&_CurveAdapter.CallOpts)
}

// FutureA1 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA1() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA1(&_CurveAdapter.CallOpts)
}

// FutureA2 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA2 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA2() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA2(&_CurveAdapter.CallOpts)
}

// FutureA2 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA2() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA2(&_CurveAdapter.CallOpts)
}

// FutureA3 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA3 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA3() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA3(&_CurveAdapter.CallOpts)
}

// FutureA3 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA3() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA3(&_CurveAdapter.CallOpts)
}

// FutureA4 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA4 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA4() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA4(&_CurveAdapter.CallOpts)
}

// FutureA4 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA4() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA4(&_CurveAdapter.CallOpts)
}

// FutureA5 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureA5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureA5 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureA5() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA5(&_CurveAdapter.CallOpts)
}

// FutureA5 is a free data retrieval call binding the contract method 0xb4b577ad.
//
// Solidity: function future_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureA5() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureA5(&_CurveAdapter.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime(&_CurveAdapter.CallOpts)
}

// FutureATime is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime(&_CurveAdapter.CallOpts)
}

// FutureATime0 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime0 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime0() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime0(&_CurveAdapter.CallOpts)
}

// FutureATime0 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime0() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime0(&_CurveAdapter.CallOpts)
}

// FutureATime1 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime1 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime1() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime1(&_CurveAdapter.CallOpts)
}

// FutureATime1 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime1() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime1(&_CurveAdapter.CallOpts)
}

// FutureATime2 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime2 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime2() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime2(&_CurveAdapter.CallOpts)
}

// FutureATime2 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime2() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime2(&_CurveAdapter.CallOpts)
}

// FutureATime3 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime3 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime3() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime3(&_CurveAdapter.CallOpts)
}

// FutureATime3 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime3() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime3(&_CurveAdapter.CallOpts)
}

// FutureATime4 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime4 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime4() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime4(&_CurveAdapter.CallOpts)
}

// FutureATime4 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime4() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime4(&_CurveAdapter.CallOpts)
}

// FutureATime5 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) FutureATime5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "future_A_time5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureATime5 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) FutureATime5() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime5(&_CurveAdapter.CallOpts)
}

// FutureATime5 is a free data retrieval call binding the contract method 0x14052288.
//
// Solidity: function future_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) FutureATime5() (*big.Int, error) {
	return _CurveAdapter.Contract.FutureATime5(&_CurveAdapter.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetBalances(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_balances")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetBalances() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances(&_CurveAdapter.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetBalances() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances(&_CurveAdapter.CallOpts)
}

// GetBalances0 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCaller) GetBalances0(opts *bind.CallOpts) ([3]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_balances0")

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetBalances0 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterSession) GetBalances0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances0(&_CurveAdapter.CallOpts)
}

// GetBalances0 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCallerSession) GetBalances0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances0(&_CurveAdapter.CallOpts)
}

// GetBalances1 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCaller) GetBalances1(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_balances1")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBalances1 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterSession) GetBalances1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances1(&_CurveAdapter.CallOpts)
}

// GetBalances1 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCallerSession) GetBalances1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances1(&_CurveAdapter.CallOpts)
}

// GetBalances2 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetBalances2(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_balances2")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances2 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetBalances2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances2(&_CurveAdapter.CallOpts)
}

// GetBalances2 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetBalances2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances2(&_CurveAdapter.CallOpts)
}

// GetBalances3 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetBalances3(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_balances3")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances3 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetBalances3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances3(&_CurveAdapter.CallOpts)
}

// GetBalances3 is a free data retrieval call binding the contract method 0x14f05979.
//
// Solidity: function get_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetBalances3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetBalances3(&_CurveAdapter.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy0 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy0(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy0", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy0 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy0(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy0(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy0 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy0(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy0(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy1 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy1(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy1", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy1 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy1(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy1(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy1 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy1(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy1(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy2 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy2(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy2", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy2 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy2(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy2(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy2 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy2(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy2(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy3 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy3(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy3", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy3 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy3(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy3(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy3 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy3(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy3(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy4 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy4(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy4", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy4 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy4(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy4(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy4 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy4(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy4(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy5 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDy5(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy5", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy5 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDy5(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy5(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDy5 is a free data retrieval call binding the contract method 0x5e0d443f.
//
// Solidity: function get_dy(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDy5(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDy5(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying0 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying0(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying0", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying0 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying0(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying0(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying0 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying0(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying0(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying1 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying1(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying1", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying1 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying1(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying1(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying1 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying1(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying1(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying2 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying2(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying2", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying2 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying2(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying2(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying2 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying2(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying2(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying3 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying3(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying3", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying3 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying3(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying3(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying3 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying3(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying3(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying4 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying4(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying4", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying4 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying4(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying4(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying4 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 i, int128 j, uint256 dx) view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying4(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying4(&_CurveAdapter.CallOpts, i, j, dx)
}

// GetDyUnderlying5 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 , int128 , uint256 ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetDyUnderlying5(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_dy_underlying5", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDyUnderlying5 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 , int128 , uint256 ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetDyUnderlying5(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying5(&_CurveAdapter.CallOpts, arg0, arg1, arg2)
}

// GetDyUnderlying5 is a free data retrieval call binding the contract method 0x07211ef7.
//
// Solidity: function get_dy_underlying(int128 , int128 , uint256 ) pure returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetDyUnderlying5(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (*big.Int, error) {
	return _CurveAdapter.Contract.GetDyUnderlying5(&_CurveAdapter.CallOpts, arg0, arg1, arg2)
}

// GetPreviousBalances is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPreviousBalances(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_previous_balances")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPreviousBalances is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPreviousBalances() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPreviousBalances() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances0 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCaller) GetPreviousBalances0(opts *bind.CallOpts) ([3]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_previous_balances0")

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetPreviousBalances0 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterSession) GetPreviousBalances0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances0(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances0 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCallerSession) GetPreviousBalances0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances0(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances1 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCaller) GetPreviousBalances1(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_previous_balances1")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetPreviousBalances1 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterSession) GetPreviousBalances1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances1(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances1 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCallerSession) GetPreviousBalances1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances1(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances2 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPreviousBalances2(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_previous_balances2")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPreviousBalances2 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPreviousBalances2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances2(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances2 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPreviousBalances2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances2(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances3 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPreviousBalances3(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_previous_balances3")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPreviousBalances3 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPreviousBalances3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances3(&_CurveAdapter.CallOpts)
}

// GetPreviousBalances3 is a free data retrieval call binding the contract method 0xd96c7fce.
//
// Solidity: function get_previous_balances() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPreviousBalances3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPreviousBalances3(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPriceCumulativeLast(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_price_cumulative_last")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPriceCumulativeLast is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPriceCumulativeLast() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPriceCumulativeLast() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast0 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCaller) GetPriceCumulativeLast0(opts *bind.CallOpts) ([3]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_price_cumulative_last0")

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetPriceCumulativeLast0 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterSession) GetPriceCumulativeLast0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast0(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast0 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCallerSession) GetPriceCumulativeLast0() ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast0(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast1 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCaller) GetPriceCumulativeLast1(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_price_cumulative_last1")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetPriceCumulativeLast1 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterSession) GetPriceCumulativeLast1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast1(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast1 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCallerSession) GetPriceCumulativeLast1() ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast1(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast2 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPriceCumulativeLast2(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_price_cumulative_last2")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPriceCumulativeLast2 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPriceCumulativeLast2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast2(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast2 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPriceCumulativeLast2() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast2(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast3 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetPriceCumulativeLast3(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_price_cumulative_last3")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPriceCumulativeLast3 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetPriceCumulativeLast3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast3(&_CurveAdapter.CallOpts)
}

// GetPriceCumulativeLast3 is a free data retrieval call binding the contract method 0x4469e30e.
//
// Solidity: function get_price_cumulative_last() pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetPriceCumulativeLast3() ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetPriceCumulativeLast3(&_CurveAdapter.CallOpts)
}

// GetTwapBalances is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetTwapBalances(opts *bind.CallOpts, _first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_twap_balances", _first_balances, _last_balances, _time_elapsed)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetTwapBalances is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetTwapBalances(_first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetTwapBalances(_first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances0 is a free data retrieval call binding the contract method 0x85ca3c6f.
//
// Solidity: function get_twap_balances(uint256[3] _first_balances, uint256[3] _last_balances, uint256 _time_elapsed) view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCaller) GetTwapBalances0(opts *bind.CallOpts, _first_balances [3]*big.Int, _last_balances [3]*big.Int, _time_elapsed *big.Int) ([3]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_twap_balances0", _first_balances, _last_balances, _time_elapsed)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// GetTwapBalances0 is a free data retrieval call binding the contract method 0x85ca3c6f.
//
// Solidity: function get_twap_balances(uint256[3] _first_balances, uint256[3] _last_balances, uint256 _time_elapsed) view returns(uint256[3])
func (_CurveAdapter *CurveAdapterSession) GetTwapBalances0(_first_balances [3]*big.Int, _last_balances [3]*big.Int, _time_elapsed *big.Int) ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances0(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances0 is a free data retrieval call binding the contract method 0x85ca3c6f.
//
// Solidity: function get_twap_balances(uint256[3] _first_balances, uint256[3] _last_balances, uint256 _time_elapsed) view returns(uint256[3])
func (_CurveAdapter *CurveAdapterCallerSession) GetTwapBalances0(_first_balances [3]*big.Int, _last_balances [3]*big.Int, _time_elapsed *big.Int) ([3]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances0(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances1 is a free data retrieval call binding the contract method 0x53b546cd.
//
// Solidity: function get_twap_balances(uint256[4] _first_balances, uint256[4] _last_balances, uint256 _time_elapsed) view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCaller) GetTwapBalances1(opts *bind.CallOpts, _first_balances [4]*big.Int, _last_balances [4]*big.Int, _time_elapsed *big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_twap_balances1", _first_balances, _last_balances, _time_elapsed)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetTwapBalances1 is a free data retrieval call binding the contract method 0x53b546cd.
//
// Solidity: function get_twap_balances(uint256[4] _first_balances, uint256[4] _last_balances, uint256 _time_elapsed) view returns(uint256[4])
func (_CurveAdapter *CurveAdapterSession) GetTwapBalances1(_first_balances [4]*big.Int, _last_balances [4]*big.Int, _time_elapsed *big.Int) ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances1(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances1 is a free data retrieval call binding the contract method 0x53b546cd.
//
// Solidity: function get_twap_balances(uint256[4] _first_balances, uint256[4] _last_balances, uint256 _time_elapsed) view returns(uint256[4])
func (_CurveAdapter *CurveAdapterCallerSession) GetTwapBalances1(_first_balances [4]*big.Int, _last_balances [4]*big.Int, _time_elapsed *big.Int) ([4]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances1(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances2 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetTwapBalances2(opts *bind.CallOpts, _first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_twap_balances2", _first_balances, _last_balances, _time_elapsed)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetTwapBalances2 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetTwapBalances2(_first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances2(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances2 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] _first_balances, uint256[2] _last_balances, uint256 _time_elapsed) view returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetTwapBalances2(_first_balances [2]*big.Int, _last_balances [2]*big.Int, _time_elapsed *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances2(&_CurveAdapter.CallOpts, _first_balances, _last_balances, _time_elapsed)
}

// GetTwapBalances3 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] , uint256[2] , uint256 ) pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCaller) GetTwapBalances3(opts *bind.CallOpts, arg0 [2]*big.Int, arg1 [2]*big.Int, arg2 *big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_twap_balances3", arg0, arg1, arg2)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetTwapBalances3 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] , uint256[2] , uint256 ) pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterSession) GetTwapBalances3(arg0 [2]*big.Int, arg1 [2]*big.Int, arg2 *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances3(&_CurveAdapter.CallOpts, arg0, arg1, arg2)
}

// GetTwapBalances3 is a free data retrieval call binding the contract method 0x0f6ba8e3.
//
// Solidity: function get_twap_balances(uint256[2] , uint256[2] , uint256 ) pure returns(uint256[2])
func (_CurveAdapter *CurveAdapterCallerSession) GetTwapBalances3(arg0 [2]*big.Int, arg1 [2]*big.Int, arg2 *big.Int) ([2]*big.Int, error) {
	return _CurveAdapter.Contract.GetTwapBalances3(&_CurveAdapter.CallOpts, arg0, arg1, arg2)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice0 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice0 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice0() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice0(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice0 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice0() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice0(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice1 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice1 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice1() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice1(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice1 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice1() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice1(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice2 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice2 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice2() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice2(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice2 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice2() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice2(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice3 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice3 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice3() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice3(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice3 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice3() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice3(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice4 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice4 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice4() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice4(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice4 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice4() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice4(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice5 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) GetVirtualPrice5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "get_virtual_price5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice5 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) GetVirtualPrice5() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice5(&_CurveAdapter.CallOpts)
}

// GetVirtualPrice5 is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) GetVirtualPrice5() (*big.Int, error) {
	return _CurveAdapter.Contract.GetVirtualPrice5(&_CurveAdapter.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA(&_CurveAdapter.CallOpts)
}

// InitialA is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA(&_CurveAdapter.CallOpts)
}

// InitialA0 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA0 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA0() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA0(&_CurveAdapter.CallOpts)
}

// InitialA0 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA0() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA0(&_CurveAdapter.CallOpts)
}

// InitialA1 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA1 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA1() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA1(&_CurveAdapter.CallOpts)
}

// InitialA1 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA1() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA1(&_CurveAdapter.CallOpts)
}

// InitialA2 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA2 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA2() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA2(&_CurveAdapter.CallOpts)
}

// InitialA2 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA2() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA2(&_CurveAdapter.CallOpts)
}

// InitialA3 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA3 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA3() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA3(&_CurveAdapter.CallOpts)
}

// InitialA3 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA3() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA3(&_CurveAdapter.CallOpts)
}

// InitialA4 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA4 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA4() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA4(&_CurveAdapter.CallOpts)
}

// InitialA4 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA4() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA4(&_CurveAdapter.CallOpts)
}

// InitialA5 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialA5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialA5 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialA5() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA5(&_CurveAdapter.CallOpts)
}

// InitialA5 is a free data retrieval call binding the contract method 0x5409491a.
//
// Solidity: function initial_A() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialA5() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialA5(&_CurveAdapter.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime(&_CurveAdapter.CallOpts)
}

// InitialATime is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime(&_CurveAdapter.CallOpts)
}

// InitialATime0 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime0 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime0() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime0(&_CurveAdapter.CallOpts)
}

// InitialATime0 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime0() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime0(&_CurveAdapter.CallOpts)
}

// InitialATime1 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime1 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime1() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime1(&_CurveAdapter.CallOpts)
}

// InitialATime1 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime1() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime1(&_CurveAdapter.CallOpts)
}

// InitialATime2 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime2 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime2() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime2(&_CurveAdapter.CallOpts)
}

// InitialATime2 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime2() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime2(&_CurveAdapter.CallOpts)
}

// InitialATime3 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime3 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime3() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime3(&_CurveAdapter.CallOpts)
}

// InitialATime3 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime3() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime3(&_CurveAdapter.CallOpts)
}

// InitialATime4 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime4 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime4() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime4(&_CurveAdapter.CallOpts)
}

// InitialATime4 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime4() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime4(&_CurveAdapter.CallOpts)
}

// InitialATime5 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) InitialATime5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "initial_A_time5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialATime5 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) InitialATime5() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime5(&_CurveAdapter.CallOpts)
}

// InitialATime5 is a free data retrieval call binding the contract method 0x2081066c.
//
// Solidity: function initial_A_time() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) InitialATime5() (*big.Int, error) {
	return _CurveAdapter.Contract.InitialATime5(&_CurveAdapter.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken(&_CurveAdapter.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken(&_CurveAdapter.CallOpts)
}

// LpToken0 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken0 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken0() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken0(&_CurveAdapter.CallOpts)
}

// LpToken0 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken0() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken0(&_CurveAdapter.CallOpts)
}

// LpToken1 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken1 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken1() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken1(&_CurveAdapter.CallOpts)
}

// LpToken1 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken1() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken1(&_CurveAdapter.CallOpts)
}

// LpToken2 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken2 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken2() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken2(&_CurveAdapter.CallOpts)
}

// LpToken2 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken2() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken2(&_CurveAdapter.CallOpts)
}

// LpToken3 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken3 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken3() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken3(&_CurveAdapter.CallOpts)
}

// LpToken3 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken3() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken3(&_CurveAdapter.CallOpts)
}

// LpToken4 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken4 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken4() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken4(&_CurveAdapter.CallOpts)
}

// LpToken4 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken4() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken4(&_CurveAdapter.CallOpts)
}

// LpToken5 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) LpToken5(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "lp_token5")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken5 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) LpToken5() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken5(&_CurveAdapter.CallOpts)
}

// LpToken5 is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) LpToken5() (common.Address, error) {
	return _CurveAdapter.Contract.LpToken5(&_CurveAdapter.CallOpts)
}

// MetapoolBase is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase(&_CurveAdapter.CallOpts)
}

// MetapoolBase is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase(&_CurveAdapter.CallOpts)
}

// MetapoolBase0 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase0 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase0() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase0(&_CurveAdapter.CallOpts)
}

// MetapoolBase0 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase0() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase0(&_CurveAdapter.CallOpts)
}

// MetapoolBase1 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase1 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase1() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase1(&_CurveAdapter.CallOpts)
}

// MetapoolBase1 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase1() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase1(&_CurveAdapter.CallOpts)
}

// MetapoolBase2 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase2 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase2() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase2(&_CurveAdapter.CallOpts)
}

// MetapoolBase2 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase2() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase2(&_CurveAdapter.CallOpts)
}

// MetapoolBase3 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase3 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase3() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase3(&_CurveAdapter.CallOpts)
}

// MetapoolBase3 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase3() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase3(&_CurveAdapter.CallOpts)
}

// MetapoolBase4 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) MetapoolBase4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "metapoolBase4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetapoolBase4 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterSession) MetapoolBase4() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase4(&_CurveAdapter.CallOpts)
}

// MetapoolBase4 is a free data retrieval call binding the contract method 0x64a89bca.
//
// Solidity: function metapoolBase() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) MetapoolBase4() (common.Address, error) {
	return _CurveAdapter.Contract.MetapoolBase4(&_CurveAdapter.CallOpts)
}

// NCoins is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins(&_CurveAdapter.CallOpts)
}

// NCoins is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins(&_CurveAdapter.CallOpts)
}

// NCoins0 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins0 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins0() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins0(&_CurveAdapter.CallOpts)
}

// NCoins0 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins0() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins0(&_CurveAdapter.CallOpts)
}

// NCoins1 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins1 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins1() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins1(&_CurveAdapter.CallOpts)
}

// NCoins1 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins1() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins1(&_CurveAdapter.CallOpts)
}

// NCoins2 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins2 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins2() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins2(&_CurveAdapter.CallOpts)
}

// NCoins2 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins2() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins2(&_CurveAdapter.CallOpts)
}

// NCoins3 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins3 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins3() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins3(&_CurveAdapter.CallOpts)
}

// NCoins3 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins3() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins3(&_CurveAdapter.CallOpts)
}

// NCoins4 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) NCoins4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "nCoins4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NCoins4 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) NCoins4() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins4(&_CurveAdapter.CallOpts)
}

// NCoins4 is a free data retrieval call binding the contract method 0xc21ee162.
//
// Solidity: function nCoins() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) NCoins4() (*big.Int, error) {
	return _CurveAdapter.Contract.NCoins4(&_CurveAdapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name() (string, error) {
	return _CurveAdapter.Contract.Name(&_CurveAdapter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name() (string, error) {
	return _CurveAdapter.Contract.Name(&_CurveAdapter.CallOpts)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name0() (string, error) {
	return _CurveAdapter.Contract.Name0(&_CurveAdapter.CallOpts)
}

// Name0 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name0() (string, error) {
	return _CurveAdapter.Contract.Name0(&_CurveAdapter.CallOpts)
}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name1() (string, error) {
	return _CurveAdapter.Contract.Name1(&_CurveAdapter.CallOpts)
}

// Name1 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name1() (string, error) {
	return _CurveAdapter.Contract.Name1(&_CurveAdapter.CallOpts)
}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name2() (string, error) {
	return _CurveAdapter.Contract.Name2(&_CurveAdapter.CallOpts)
}

// Name2 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name2() (string, error) {
	return _CurveAdapter.Contract.Name2(&_CurveAdapter.CallOpts)
}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name3(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name3")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name3() (string, error) {
	return _CurveAdapter.Contract.Name3(&_CurveAdapter.CallOpts)
}

// Name3 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name3() (string, error) {
	return _CurveAdapter.Contract.Name3(&_CurveAdapter.CallOpts)
}

// Name4 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name4(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name4")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name4 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Name4() (string, error) {
	return _CurveAdapter.Contract.Name4(&_CurveAdapter.CallOpts)
}

// Name4 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name4() (string, error) {
	return _CurveAdapter.Contract.Name4(&_CurveAdapter.CallOpts)
}

// Name5 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CurveAdapter *CurveAdapterCaller) Name5(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "name5")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name5 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CurveAdapter *CurveAdapterSession) Name5() (string, error) {
	return _CurveAdapter.Contract.Name5(&_CurveAdapter.CallOpts)
}

// Name5 is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Name5() (string, error) {
	return _CurveAdapter.Contract.Name5(&_CurveAdapter.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Pool() (common.Address, error) {
	return _CurveAdapter.Contract.Pool(&_CurveAdapter.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Pool() (common.Address, error) {
	return _CurveAdapter.Contract.Pool(&_CurveAdapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol() (string, error) {
	return _CurveAdapter.Contract.Symbol(&_CurveAdapter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol() (string, error) {
	return _CurveAdapter.Contract.Symbol(&_CurveAdapter.CallOpts)
}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol0() (string, error) {
	return _CurveAdapter.Contract.Symbol0(&_CurveAdapter.CallOpts)
}

// Symbol0 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol0() (string, error) {
	return _CurveAdapter.Contract.Symbol0(&_CurveAdapter.CallOpts)
}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol1() (string, error) {
	return _CurveAdapter.Contract.Symbol1(&_CurveAdapter.CallOpts)
}

// Symbol1 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol1() (string, error) {
	return _CurveAdapter.Contract.Symbol1(&_CurveAdapter.CallOpts)
}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol2() (string, error) {
	return _CurveAdapter.Contract.Symbol2(&_CurveAdapter.CallOpts)
}

// Symbol2 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol2() (string, error) {
	return _CurveAdapter.Contract.Symbol2(&_CurveAdapter.CallOpts)
}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol3(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol3")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol3() (string, error) {
	return _CurveAdapter.Contract.Symbol3(&_CurveAdapter.CallOpts)
}

// Symbol3 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol3() (string, error) {
	return _CurveAdapter.Contract.Symbol3(&_CurveAdapter.CallOpts)
}

// Symbol4 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol4(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol4")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol4 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol4() (string, error) {
	return _CurveAdapter.Contract.Symbol4(&_CurveAdapter.CallOpts)
}

// Symbol4 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol4() (string, error) {
	return _CurveAdapter.Contract.Symbol4(&_CurveAdapter.CallOpts)
}

// Symbol5 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CurveAdapter *CurveAdapterCaller) Symbol5(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "symbol5")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol5 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CurveAdapter *CurveAdapterSession) Symbol5() (string, error) {
	return _CurveAdapter.Contract.Symbol5(&_CurveAdapter.CallOpts)
}

// Symbol5 is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_CurveAdapter *CurveAdapterCallerSession) Symbol5() (string, error) {
	return _CurveAdapter.Contract.Symbol5(&_CurveAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract(&_CurveAdapter.CallOpts)
}

// TargetContract is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract(&_CurveAdapter.CallOpts)
}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract0() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract0(&_CurveAdapter.CallOpts)
}

// TargetContract0 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract0() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract0(&_CurveAdapter.CallOpts)
}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract1() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract1(&_CurveAdapter.CallOpts)
}

// TargetContract1 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract1() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract1(&_CurveAdapter.CallOpts)
}

// TargetContract2 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract2 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract2() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract2(&_CurveAdapter.CallOpts)
}

// TargetContract2 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract2() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract2(&_CurveAdapter.CallOpts)
}

// TargetContract3 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract3 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract3() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract3(&_CurveAdapter.CallOpts)
}

// TargetContract3 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract3() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract3(&_CurveAdapter.CallOpts)
}

// TargetContract4 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) TargetContract4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "targetContract4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TargetContract4 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterSession) TargetContract4() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract4(&_CurveAdapter.CallOpts)
}

// TargetContract4 is a free data retrieval call binding the contract method 0xbd90df70.
//
// Solidity: function targetContract() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) TargetContract4() (common.Address, error) {
	return _CurveAdapter.Contract.TargetContract4(&_CurveAdapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token() (common.Address, error) {
	return _CurveAdapter.Contract.Token(&_CurveAdapter.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token() (common.Address, error) {
	return _CurveAdapter.Contract.Token(&_CurveAdapter.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token0() (common.Address, error) {
	return _CurveAdapter.Contract.Token0(&_CurveAdapter.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token0() (common.Address, error) {
	return _CurveAdapter.Contract.Token0(&_CurveAdapter.CallOpts)
}

// Token00 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token00(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token00")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token00 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token00() (common.Address, error) {
	return _CurveAdapter.Contract.Token00(&_CurveAdapter.CallOpts)
}

// Token00 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token00() (common.Address, error) {
	return _CurveAdapter.Contract.Token00(&_CurveAdapter.CallOpts)
}

// Token01 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token01(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token01")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token01 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token01() (common.Address, error) {
	return _CurveAdapter.Contract.Token01(&_CurveAdapter.CallOpts)
}

// Token01 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token01() (common.Address, error) {
	return _CurveAdapter.Contract.Token01(&_CurveAdapter.CallOpts)
}

// Token02 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token02 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token02() (common.Address, error) {
	return _CurveAdapter.Contract.Token02(&_CurveAdapter.CallOpts)
}

// Token02 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token02() (common.Address, error) {
	return _CurveAdapter.Contract.Token02(&_CurveAdapter.CallOpts)
}

// Token03 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token03(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token03")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token03 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token03() (common.Address, error) {
	return _CurveAdapter.Contract.Token03(&_CurveAdapter.CallOpts)
}

// Token03 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token03() (common.Address, error) {
	return _CurveAdapter.Contract.Token03(&_CurveAdapter.CallOpts)
}

// Token04 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token04(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token04")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token04 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token04() (common.Address, error) {
	return _CurveAdapter.Contract.Token04(&_CurveAdapter.CallOpts)
}

// Token04 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token04() (common.Address, error) {
	return _CurveAdapter.Contract.Token04(&_CurveAdapter.CallOpts)
}

// Token05 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token05(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token05")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token05 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token05() (common.Address, error) {
	return _CurveAdapter.Contract.Token05(&_CurveAdapter.CallOpts)
}

// Token05 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token05() (common.Address, error) {
	return _CurveAdapter.Contract.Token05(&_CurveAdapter.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token1() (common.Address, error) {
	return _CurveAdapter.Contract.Token1(&_CurveAdapter.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token1() (common.Address, error) {
	return _CurveAdapter.Contract.Token1(&_CurveAdapter.CallOpts)
}

// Token10 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token10(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token10")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token10 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token10() (common.Address, error) {
	return _CurveAdapter.Contract.Token10(&_CurveAdapter.CallOpts)
}

// Token10 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token10() (common.Address, error) {
	return _CurveAdapter.Contract.Token10(&_CurveAdapter.CallOpts)
}

// Token11 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token11(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token11")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token11 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token11() (common.Address, error) {
	return _CurveAdapter.Contract.Token11(&_CurveAdapter.CallOpts)
}

// Token11 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token11() (common.Address, error) {
	return _CurveAdapter.Contract.Token11(&_CurveAdapter.CallOpts)
}

// Token12 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token12(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token12")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token12 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token12() (common.Address, error) {
	return _CurveAdapter.Contract.Token12(&_CurveAdapter.CallOpts)
}

// Token12 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token12() (common.Address, error) {
	return _CurveAdapter.Contract.Token12(&_CurveAdapter.CallOpts)
}

// Token13 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token13(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token13")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token13 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token13() (common.Address, error) {
	return _CurveAdapter.Contract.Token13(&_CurveAdapter.CallOpts)
}

// Token13 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token13() (common.Address, error) {
	return _CurveAdapter.Contract.Token13(&_CurveAdapter.CallOpts)
}

// Token14 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token14(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token14")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token14 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token14() (common.Address, error) {
	return _CurveAdapter.Contract.Token14(&_CurveAdapter.CallOpts)
}

// Token14 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token14() (common.Address, error) {
	return _CurveAdapter.Contract.Token14(&_CurveAdapter.CallOpts)
}

// Token15 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token15(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token15")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token15 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token15() (common.Address, error) {
	return _CurveAdapter.Contract.Token15(&_CurveAdapter.CallOpts)
}

// Token15 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token15() (common.Address, error) {
	return _CurveAdapter.Contract.Token15(&_CurveAdapter.CallOpts)
}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token2() (common.Address, error) {
	return _CurveAdapter.Contract.Token2(&_CurveAdapter.CallOpts)
}

// Token2 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token2() (common.Address, error) {
	return _CurveAdapter.Contract.Token2(&_CurveAdapter.CallOpts)
}

// Token20 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token20 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token20() (common.Address, error) {
	return _CurveAdapter.Contract.Token20(&_CurveAdapter.CallOpts)
}

// Token20 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token20() (common.Address, error) {
	return _CurveAdapter.Contract.Token20(&_CurveAdapter.CallOpts)
}

// Token21 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token21(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token21")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token21 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token21() (common.Address, error) {
	return _CurveAdapter.Contract.Token21(&_CurveAdapter.CallOpts)
}

// Token21 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token21() (common.Address, error) {
	return _CurveAdapter.Contract.Token21(&_CurveAdapter.CallOpts)
}

// Token22 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token22(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token22")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token22 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token22() (common.Address, error) {
	return _CurveAdapter.Contract.Token22(&_CurveAdapter.CallOpts)
}

// Token22 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token22() (common.Address, error) {
	return _CurveAdapter.Contract.Token22(&_CurveAdapter.CallOpts)
}

// Token23 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token23(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token23")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token23 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token23() (common.Address, error) {
	return _CurveAdapter.Contract.Token23(&_CurveAdapter.CallOpts)
}

// Token23 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token23() (common.Address, error) {
	return _CurveAdapter.Contract.Token23(&_CurveAdapter.CallOpts)
}

// Token24 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token24 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token24() (common.Address, error) {
	return _CurveAdapter.Contract.Token24(&_CurveAdapter.CallOpts)
}

// Token24 is a free data retrieval call binding the contract method 0x25be124e.
//
// Solidity: function token2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token24() (common.Address, error) {
	return _CurveAdapter.Contract.Token24(&_CurveAdapter.CallOpts)
}

// Token3 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token3 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token3() (common.Address, error) {
	return _CurveAdapter.Contract.Token3(&_CurveAdapter.CallOpts)
}

// Token3 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token3() (common.Address, error) {
	return _CurveAdapter.Contract.Token3(&_CurveAdapter.CallOpts)
}

// Token30 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token30(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token30")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token30 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token30() (common.Address, error) {
	return _CurveAdapter.Contract.Token30(&_CurveAdapter.CallOpts)
}

// Token30 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token30() (common.Address, error) {
	return _CurveAdapter.Contract.Token30(&_CurveAdapter.CallOpts)
}

// Token31 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token31(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token31")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token31 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token31() (common.Address, error) {
	return _CurveAdapter.Contract.Token31(&_CurveAdapter.CallOpts)
}

// Token31 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token31() (common.Address, error) {
	return _CurveAdapter.Contract.Token31(&_CurveAdapter.CallOpts)
}

// Token32 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token32(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token32")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token32 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token32() (common.Address, error) {
	return _CurveAdapter.Contract.Token32(&_CurveAdapter.CallOpts)
}

// Token32 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token32() (common.Address, error) {
	return _CurveAdapter.Contract.Token32(&_CurveAdapter.CallOpts)
}

// Token33 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token33(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token33")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token33 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token33() (common.Address, error) {
	return _CurveAdapter.Contract.Token33(&_CurveAdapter.CallOpts)
}

// Token33 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token33() (common.Address, error) {
	return _CurveAdapter.Contract.Token33(&_CurveAdapter.CallOpts)
}

// Token34 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token34(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token34")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token34 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token34() (common.Address, error) {
	return _CurveAdapter.Contract.Token34(&_CurveAdapter.CallOpts)
}

// Token34 is a free data retrieval call binding the contract method 0xef14101e.
//
// Solidity: function token3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token34() (common.Address, error) {
	return _CurveAdapter.Contract.Token34(&_CurveAdapter.CallOpts)
}

// Token4 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token4(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token4")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token4 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token4() (common.Address, error) {
	return _CurveAdapter.Contract.Token4(&_CurveAdapter.CallOpts)
}

// Token4 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token4() (common.Address, error) {
	return _CurveAdapter.Contract.Token4(&_CurveAdapter.CallOpts)
}

// Token5 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token5(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token5")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token5 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token5() (common.Address, error) {
	return _CurveAdapter.Contract.Token5(&_CurveAdapter.CallOpts)
}

// Token5 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token5() (common.Address, error) {
	return _CurveAdapter.Contract.Token5(&_CurveAdapter.CallOpts)
}

// Token6 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token6(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token6")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token6 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token6() (common.Address, error) {
	return _CurveAdapter.Contract.Token6(&_CurveAdapter.CallOpts)
}

// Token6 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token6() (common.Address, error) {
	return _CurveAdapter.Contract.Token6(&_CurveAdapter.CallOpts)
}

// Token7 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token7(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token7")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token7 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token7() (common.Address, error) {
	return _CurveAdapter.Contract.Token7(&_CurveAdapter.CallOpts)
}

// Token7 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token7() (common.Address, error) {
	return _CurveAdapter.Contract.Token7(&_CurveAdapter.CallOpts)
}

// Token8 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token8(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token8")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token8 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token8() (common.Address, error) {
	return _CurveAdapter.Contract.Token8(&_CurveAdapter.CallOpts)
}

// Token8 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token8() (common.Address, error) {
	return _CurveAdapter.Contract.Token8(&_CurveAdapter.CallOpts)
}

// Token9 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Token9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "token9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token9 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Token9() (common.Address, error) {
	return _CurveAdapter.Contract.Token9(&_CurveAdapter.CallOpts)
}

// Token9 is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Token9() (common.Address, error) {
	return _CurveAdapter.Contract.Token9(&_CurveAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply(&_CurveAdapter.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply(&_CurveAdapter.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply0() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply0(&_CurveAdapter.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply0() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply0(&_CurveAdapter.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply1() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply1(&_CurveAdapter.CallOpts)
}

// TotalSupply1 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply1() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply1(&_CurveAdapter.CallOpts)
}

// TotalSupply2 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply2 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply2() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply2(&_CurveAdapter.CallOpts)
}

// TotalSupply2 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply2() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply2(&_CurveAdapter.CallOpts)
}

// TotalSupply3 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply3 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply3() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply3(&_CurveAdapter.CallOpts)
}

// TotalSupply3 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply3() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply3(&_CurveAdapter.CallOpts)
}

// TotalSupply4 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply4 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply4() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply4(&_CurveAdapter.CallOpts)
}

// TotalSupply4 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply4() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply4(&_CurveAdapter.CallOpts)
}

// TotalSupply5 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CurveAdapter *CurveAdapterCaller) TotalSupply5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "totalSupply5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply5 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CurveAdapter *CurveAdapterSession) TotalSupply5() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply5(&_CurveAdapter.CallOpts)
}

// TotalSupply5 is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() pure returns(uint256)
func (_CurveAdapter *CurveAdapterCallerSession) TotalSupply5() (*big.Int, error) {
	return _CurveAdapter.Contract.TotalSupply5(&_CurveAdapter.CallOpts)
}

// Underlying0 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying0 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying0() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying0(&_CurveAdapter.CallOpts)
}

// Underlying0 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying0() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying0(&_CurveAdapter.CallOpts)
}

// Underlying00 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying00(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying00")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying00 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying00() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying00(&_CurveAdapter.CallOpts)
}

// Underlying00 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying00() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying00(&_CurveAdapter.CallOpts)
}

// Underlying01 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying01(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying01")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying01 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying01() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying01(&_CurveAdapter.CallOpts)
}

// Underlying01 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying01() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying01(&_CurveAdapter.CallOpts)
}

// Underlying02 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying02 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying02() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying02(&_CurveAdapter.CallOpts)
}

// Underlying02 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying02() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying02(&_CurveAdapter.CallOpts)
}

// Underlying03 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying03(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying03")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying03 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying03() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying03(&_CurveAdapter.CallOpts)
}

// Underlying03 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying03() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying03(&_CurveAdapter.CallOpts)
}

// Underlying04 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying04(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying04")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying04 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying04() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying04(&_CurveAdapter.CallOpts)
}

// Underlying04 is a free data retrieval call binding the contract method 0x8ba51dfc.
//
// Solidity: function underlying0() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying04() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying04(&_CurveAdapter.CallOpts)
}

// Underlying1 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying1 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying1() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying1(&_CurveAdapter.CallOpts)
}

// Underlying1 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying1() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying1(&_CurveAdapter.CallOpts)
}

// Underlying10 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying10(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying10")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying10 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying10() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying10(&_CurveAdapter.CallOpts)
}

// Underlying10 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying10() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying10(&_CurveAdapter.CallOpts)
}

// Underlying11 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying11(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying11")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying11 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying11() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying11(&_CurveAdapter.CallOpts)
}

// Underlying11 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying11() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying11(&_CurveAdapter.CallOpts)
}

// Underlying12 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying12(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying12")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying12 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying12() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying12(&_CurveAdapter.CallOpts)
}

// Underlying12 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying12() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying12(&_CurveAdapter.CallOpts)
}

// Underlying13 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying13(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying13")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying13 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying13() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying13(&_CurveAdapter.CallOpts)
}

// Underlying13 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying13() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying13(&_CurveAdapter.CallOpts)
}

// Underlying14 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying14(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying14")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying14 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying14() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying14(&_CurveAdapter.CallOpts)
}

// Underlying14 is a free data retrieval call binding the contract method 0x6e1d8271.
//
// Solidity: function underlying1() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying14() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying14(&_CurveAdapter.CallOpts)
}

// Underlying2 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying2 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying2() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying2(&_CurveAdapter.CallOpts)
}

// Underlying2 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying2() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying2(&_CurveAdapter.CallOpts)
}

// Underlying20 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying20 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying20() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying20(&_CurveAdapter.CallOpts)
}

// Underlying20 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying20() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying20(&_CurveAdapter.CallOpts)
}

// Underlying21 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying21(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying21")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying21 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying21() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying21(&_CurveAdapter.CallOpts)
}

// Underlying21 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying21() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying21(&_CurveAdapter.CallOpts)
}

// Underlying22 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying22(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying22")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying22 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying22() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying22(&_CurveAdapter.CallOpts)
}

// Underlying22 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying22() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying22(&_CurveAdapter.CallOpts)
}

// Underlying23 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying23(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying23")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying23 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying23() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying23(&_CurveAdapter.CallOpts)
}

// Underlying23 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying23() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying23(&_CurveAdapter.CallOpts)
}

// Underlying24 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying24(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying24")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying24 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying24() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying24(&_CurveAdapter.CallOpts)
}

// Underlying24 is a free data retrieval call binding the contract method 0x06871163.
//
// Solidity: function underlying2() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying24() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying24(&_CurveAdapter.CallOpts)
}

// Underlying3 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying3 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying3() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying3(&_CurveAdapter.CallOpts)
}

// Underlying3 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying3() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying3(&_CurveAdapter.CallOpts)
}

// Underlying30 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying30(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying30")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying30 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying30() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying30(&_CurveAdapter.CallOpts)
}

// Underlying30 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying30() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying30(&_CurveAdapter.CallOpts)
}

// Underlying31 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying31(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying31")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying31 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying31() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying31(&_CurveAdapter.CallOpts)
}

// Underlying31 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying31() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying31(&_CurveAdapter.CallOpts)
}

// Underlying32 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying32(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying32")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying32 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying32() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying32(&_CurveAdapter.CallOpts)
}

// Underlying32 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying32() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying32(&_CurveAdapter.CallOpts)
}

// Underlying33 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying33(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying33")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying33 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying33() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying33(&_CurveAdapter.CallOpts)
}

// Underlying33 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying33() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying33(&_CurveAdapter.CallOpts)
}

// Underlying34 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCaller) Underlying34(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying34")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying34 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterSession) Underlying34() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying34(&_CurveAdapter.CallOpts)
}

// Underlying34 is a free data retrieval call binding the contract method 0x57d78875.
//
// Solidity: function underlying3() view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) Underlying34() (common.Address, error) {
	return _CurveAdapter.Contract.Underlying34(&_CurveAdapter.CallOpts)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins0 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins0(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins0", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins0 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins0(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins0(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins0 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins0(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins0(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins1 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins1(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins1", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins1 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins1(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins1(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins1 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins1(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins1(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins10 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins10(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins10", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins10 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins10(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins10(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins10 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins10(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins10(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins11 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 ) pure returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins11(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins11", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins11 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 ) pure returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins11(arg0 *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins11(&_CurveAdapter.CallOpts, arg0)
}

// UnderlyingCoins11 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 ) pure returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins11(arg0 *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins11(&_CurveAdapter.CallOpts, arg0)
}

// UnderlyingCoins12 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 ) pure returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins12(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins12", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins12 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 ) pure returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins12(arg0 *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins12(&_CurveAdapter.CallOpts, arg0)
}

// UnderlyingCoins12 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 ) pure returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins12(arg0 *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins12(&_CurveAdapter.CallOpts, arg0)
}

// UnderlyingCoins2 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins2(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins2", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins2 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins2(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins2(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins2 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins2(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins2(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins3 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins3(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins3", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins3 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins3(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins3(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins3 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins3(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins3(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins4 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins4(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins4", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins4 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins4(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins4(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins4 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins4(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins4(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins5 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins5(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins5", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins5 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins5(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins5(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins5 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins5(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins5(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins6 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins6(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins6", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins6 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins6(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins6(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins6 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins6(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins6(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins7 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins7(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins7", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins7 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins7(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins7(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins7 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins7(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins7(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins8 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins8(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins8", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins8 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins8(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins8(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins8 is a free data retrieval call binding the contract method 0xb9947eb0.
//
// Solidity: function underlying_coins(uint256 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins8(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins8(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins9 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCaller) UnderlyingCoins9(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveAdapter.contract.Call(opts, &out, "underlying_coins9", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingCoins9 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterSession) UnderlyingCoins9(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins9(&_CurveAdapter.CallOpts, i)
}

// UnderlyingCoins9 is a free data retrieval call binding the contract method 0xb739953e.
//
// Solidity: function underlying_coins(int128 i) view returns(address)
func (_CurveAdapter *CurveAdapterCallerSession) UnderlyingCoins9(i *big.Int) (common.Address, error) {
	return _CurveAdapter.Contract.UnderlyingCoins9(&_CurveAdapter.CallOpts, i)
}

// AddAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin", i, rateMinRAY)
}

// AddAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin0(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin0", i, rateMinRAY)
}

// AddAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin0(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin0(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin0(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin0(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin1(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin1", i, rateMinRAY)
}

// AddAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin1(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin1(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin1(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin1(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin2(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin2", i, rateMinRAY)
}

// AddAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin2(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin2(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin2(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin2(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin3(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin3", i, rateMinRAY)
}

// AddAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin3(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin3(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin3(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin3(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddAllLiquidityOneCoin4(opts *bind.TransactOpts, i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_all_liquidity_one_coin4", i, rateMinRAY)
}

// AddAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) AddAllLiquidityOneCoin4(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin4(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xec026ca7.
//
// Solidity: function add_all_liquidity_one_coin(int128 i, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddAllLiquidityOneCoin4(i *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddAllLiquidityOneCoin4(&_CurveAdapter.TransactOpts, i, rateMinRAY)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity", amounts, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidity(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidity(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidity0(opts *bind.TransactOpts, amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity0", amounts, arg1)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidity0(amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity0(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity0 is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidity0(amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity0(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidity1(opts *bind.TransactOpts, amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity1", amounts, arg1)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidity1(amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity1(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity1 is a paid mutator transaction binding the contract method 0x029b2f34.
//
// Solidity: function add_liquidity(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidity1(amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity1(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidity2(opts *bind.TransactOpts, amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity2", amounts, arg1)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidity2(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity2(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidity2(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity2(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidity3(opts *bind.TransactOpts, amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity3", amounts, min_mint_amount)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidity3(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity3(&_CurveAdapter.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity3 is a paid mutator transaction binding the contract method 0x0b4c7e4d.
//
// Solidity: function add_liquidity(uint256[2] amounts, uint256 min_mint_amount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidity3(amounts [2]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidity3(&_CurveAdapter.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin", amount, i, minAmount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin0(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin0", amount, i, minAmount)
}

// AddLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin0(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin0(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin0(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin0(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin1(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin1", amount, i, minAmount)
}

// AddLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin1(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin1(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin1(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin1(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin2(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin2", amount, i, minAmount)
}

// AddLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin2(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin2(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin2(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin2(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin3(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin3", amount, i, minAmount)
}

// AddLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin3(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin3(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin3(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin3(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactor) AddLiquidityOneCoin4(opts *bind.TransactOpts, amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "add_liquidity_one_coin4", amount, i, minAmount)
}

// AddLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterSession) AddLiquidityOneCoin4(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin4(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// AddLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0xcf023dd0.
//
// Solidity: function add_liquidity_one_coin(uint256 amount, int128 i, uint256 minAmount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) AddLiquidityOneCoin4(amount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.AddLiquidityOneCoin4(&_CurveAdapter.TransactOpts, amount, i, minAmount)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange", i, j, arg2, arg3)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange0", i, j, arg2, arg3)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange0(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange0(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange0(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange0(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange1 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange1(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange1", i, j, arg2, arg3)
}

// Exchange1 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange1(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange1(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange1 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange1(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange1(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange2 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange2(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange2", i, j, arg2, arg3)
}

// Exchange2 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange2(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange2(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange2 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange2(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange2(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange3 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange3(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange3", i, j, arg2, arg3)
}

// Exchange3 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange3(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange3(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange3 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange3(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange3(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange4 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange4(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange4", i, j, arg2, arg3)
}

// Exchange4 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange4(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange4(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange4 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange4(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange4(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// Exchange5 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveAdapter *CurveAdapterTransactor) Exchange5(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange5", i, j, dx, min_dy)
}

// Exchange5 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveAdapter *CurveAdapterSession) Exchange5(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange5(&_CurveAdapter.TransactOpts, i, j, dx, min_dy)
}

// Exchange5 is a paid mutator transaction binding the contract method 0x3df02124.
//
// Solidity: function exchange(int128 i, int128 j, uint256 dx, uint256 min_dy) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Exchange5(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.Exchange5(&_CurveAdapter.TransactOpts, i, j, dx, min_dy)
}

// ExchangeAll is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all", i, j, rateMinRAY)
}

// ExchangeAll is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll0 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll0(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all0", i, j, rateMinRAY)
}

// ExchangeAll0 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll0(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll0(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll0 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll0(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll0(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll1 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll1(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all1", i, j, rateMinRAY)
}

// ExchangeAll1 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll1(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll1(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll1 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll1(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll1(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll2 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll2(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all2", i, j, rateMinRAY)
}

// ExchangeAll2 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll2(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll2(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll2 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll2(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll2(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll3 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll3(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all3", i, j, rateMinRAY)
}

// ExchangeAll3 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll3(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll3(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll3 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll3(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll3(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll4 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAll4(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all4", i, j, rateMinRAY)
}

// ExchangeAll4 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAll4(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll4(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAll4 is a paid mutator transaction binding the contract method 0x79bea664.
//
// Solidity: function exchange_all(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAll4(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAll4(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying", i, j, rateMinRAY)
}

// ExchangeAllUnderlying is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying0 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying0", i, j, rateMinRAY)
}

// ExchangeAllUnderlying0 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying0(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying0(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying0 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying0(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying0(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying1 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying1(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying1", i, j, rateMinRAY)
}

// ExchangeAllUnderlying1 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying1(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying1(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying1 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying1(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying1(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying2 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying2(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying2", i, j, rateMinRAY)
}

// ExchangeAllUnderlying2 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying2(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying2(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying2 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying2(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying2(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying3 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying3(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying3", i, j, rateMinRAY)
}

// ExchangeAllUnderlying3 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying3(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying3(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying3 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying3(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying3(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying4 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeAllUnderlying4(opts *bind.TransactOpts, i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_all_underlying4", i, j, rateMinRAY)
}

// ExchangeAllUnderlying4 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeAllUnderlying4(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying4(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeAllUnderlying4 is a paid mutator transaction binding the contract method 0x1af4de83.
//
// Solidity: function exchange_all_underlying(int128 i, int128 j, uint256 rateMinRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeAllUnderlying4(i *big.Int, j *big.Int, rateMinRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeAllUnderlying4(&_CurveAdapter.TransactOpts, i, j, rateMinRAY)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying", i, j, arg2, arg3)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying0(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying0", i, j, arg2, arg3)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying0(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying0(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying0 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying0(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying0(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying1 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying1(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying1", i, j, arg2, arg3)
}

// ExchangeUnderlying1 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying1(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying1(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying1 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying1(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying1(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying2 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying2(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying2", i, j, arg2, arg3)
}

// ExchangeUnderlying2 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying2(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying2(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying2 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying2(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying2(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying3 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying3(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying3", i, j, arg2, arg3)
}

// ExchangeUnderlying3 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying3(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying3(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying3 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying3(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying3(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying4 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) ExchangeUnderlying4(opts *bind.TransactOpts, i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "exchange_underlying4", i, j, arg2, arg3)
}

// ExchangeUnderlying4 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) ExchangeUnderlying4(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying4(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// ExchangeUnderlying4 is a paid mutator transaction binding the contract method 0xa6417ed6.
//
// Solidity: function exchange_underlying(int128 i, int128 j, uint256 , uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) ExchangeUnderlying4(i *big.Int, j *big.Int, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.ExchangeUnderlying4(&_CurveAdapter.TransactOpts, i, j, arg2, arg3)
}

// RemoveAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin0(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin0", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin0(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin0(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin0(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin0(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin1(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin1", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin1(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin1(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin1(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin1(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin2(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin2", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin2(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin2(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin2(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin2(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin3(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin3", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin3(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin3(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin3(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin3(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveAllLiquidityOneCoin4(opts *bind.TransactOpts, i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_all_liquidity_one_coin4", i, minRateRAY)
}

// RemoveAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveAllLiquidityOneCoin4(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin4(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveAllLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x33d2ebf2.
//
// Solidity: function remove_all_liquidity_one_coin(int128 i, uint256 minRateRAY) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveAllLiquidityOneCoin4(i *big.Int, minRateRAY *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveAllLiquidityOneCoin4(&_CurveAdapter.TransactOpts, i, minRateRAY)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidity(opts *bind.TransactOpts, arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity", arg0, arg1)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidity(arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidity(arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 , uint256[3] ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidity0(opts *bind.TransactOpts, arg0 *big.Int, arg1 [3]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity0", arg0, arg1)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 , uint256[3] ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidity0(arg0 *big.Int, arg1 [3]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity0(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity0 is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 , uint256[3] ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidity0(arg0 *big.Int, arg1 [3]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity0(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 , uint256[4] ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidity1(opts *bind.TransactOpts, arg0 *big.Int, arg1 [4]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity1", arg0, arg1)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 , uint256[4] ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidity1(arg0 *big.Int, arg1 [4]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity1(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity1 is a paid mutator transaction binding the contract method 0x7d49d875.
//
// Solidity: function remove_liquidity(uint256 , uint256[4] ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidity1(arg0 *big.Int, arg1 [4]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity1(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidity2(opts *bind.TransactOpts, arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity2", arg0, arg1)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidity2(arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity2(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity2 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 , uint256[2] ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidity2(arg0 *big.Int, arg1 [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity2(&_CurveAdapter.TransactOpts, arg0, arg1)
}

// RemoveLiquidity3 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 amount, uint256[2] min_amounts) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidity3(opts *bind.TransactOpts, amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity3", amount, min_amounts)
}

// RemoveLiquidity3 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 amount, uint256[2] min_amounts) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidity3(amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity3(&_CurveAdapter.TransactOpts, amount, min_amounts)
}

// RemoveLiquidity3 is a paid mutator transaction binding the contract method 0x5b36389c.
//
// Solidity: function remove_liquidity(uint256 amount, uint256[2] min_amounts) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidity3(amount *big.Int, min_amounts [2]*big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidity3(&_CurveAdapter.TransactOpts, amount, min_amounts)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_imbalance", amounts, arg1)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityImbalance(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityImbalance(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityImbalance0(opts *bind.TransactOpts, amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_imbalance0", amounts, arg1)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityImbalance0(amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance0(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance0 is a paid mutator transaction binding the contract method 0x9fdaea0c.
//
// Solidity: function remove_liquidity_imbalance(uint256[3] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityImbalance0(amounts [3]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance0(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance1 is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityImbalance1(opts *bind.TransactOpts, amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_imbalance1", amounts, arg1)
}

// RemoveLiquidityImbalance1 is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityImbalance1(amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance1(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance1 is a paid mutator transaction binding the contract method 0x18a7bd76.
//
// Solidity: function remove_liquidity_imbalance(uint256[4] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityImbalance1(amounts [4]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance1(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance2 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityImbalance2(opts *bind.TransactOpts, amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_imbalance2", amounts, arg1)
}

// RemoveLiquidityImbalance2 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityImbalance2(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance2(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance2 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityImbalance2(amounts [2]*big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance2(&_CurveAdapter.TransactOpts, amounts, arg1)
}

// RemoveLiquidityImbalance3 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityImbalance3(opts *bind.TransactOpts, amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_imbalance3", amounts, max_burn_amount)
}

// RemoveLiquidityImbalance3 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityImbalance3(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance3(&_CurveAdapter.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityImbalance3 is a paid mutator transaction binding the contract method 0xe3103273.
//
// Solidity: function remove_liquidity_imbalance(uint256[2] amounts, uint256 max_burn_amount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityImbalance3(amounts [2]*big.Int, max_burn_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityImbalance3(&_CurveAdapter.TransactOpts, amounts, max_burn_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin", arg0, i, arg2)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin0(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin0", arg0, i, arg2)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin0(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin0(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin0 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin0(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin0(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin1(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin1", arg0, i, arg2)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin1(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin1(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin1 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin1(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin1(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin2(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin2", arg0, i, arg2)
}

// RemoveLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin2(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin2(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin2 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin2(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin2(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin3(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin3", arg0, i, arg2)
}

// RemoveLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin3(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin3(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin3 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin3(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin3(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin4(opts *bind.TransactOpts, arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin4", arg0, i, arg2)
}

// RemoveLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin4(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin4(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin4 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 , int128 i, uint256 ) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin4(arg0 *big.Int, i *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin4(&_CurveAdapter.TransactOpts, arg0, i, arg2)
}

// RemoveLiquidityOneCoin5 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurveAdapter *CurveAdapterTransactor) RemoveLiquidityOneCoin5(opts *bind.TransactOpts, _token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.contract.Transact(opts, "remove_liquidity_one_coin5", _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin5 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurveAdapter *CurveAdapterSession) RemoveLiquidityOneCoin5(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin5(&_CurveAdapter.TransactOpts, _token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin5 is a paid mutator transaction binding the contract method 0x1a4d01d2.
//
// Solidity: function remove_liquidity_one_coin(uint256 _token_amount, int128 i, uint256 min_amount) returns()
func (_CurveAdapter *CurveAdapterTransactorSession) RemoveLiquidityOneCoin5(_token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveAdapter.Contract.RemoveLiquidityOneCoin5(&_CurveAdapter.TransactOpts, _token_amount, i, min_amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CurveAdapter *CurveAdapterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveAdapter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CurveAdapter *CurveAdapterSession) Receive() (*types.Transaction, error) {
	return _CurveAdapter.Contract.Receive(&_CurveAdapter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CurveAdapter *CurveAdapterTransactorSession) Receive() (*types.Transaction, error) {
	return _CurveAdapter.Contract.Receive(&_CurveAdapter.TransactOpts)
}
