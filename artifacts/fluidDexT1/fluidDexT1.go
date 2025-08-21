// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fluidDexT1

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

// StructsCollateralReserves is an auto generated low-level Go binding around an user-defined struct.
type StructsCollateralReserves struct {
	Token0RealReserves      *big.Int
	Token1RealReserves      *big.Int
	Token0ImaginaryReserves *big.Int
	Token1ImaginaryReserves *big.Int
}

// StructsConstantViews is an auto generated low-level Go binding around an user-defined struct.
type StructsConstantViews struct {
	DexId                   *big.Int
	Liquidity               common.Address
	Factory                 common.Address
	Implementations         StructsImplementations
	DeployerContract        common.Address
	Token0                  common.Address
	Token1                  common.Address
	SupplyToken0Slot        [32]byte
	BorrowToken0Slot        [32]byte
	SupplyToken1Slot        [32]byte
	BorrowToken1Slot        [32]byte
	ExchangePriceToken0Slot [32]byte
	ExchangePriceToken1Slot [32]byte
	OracleMapping           *big.Int
}

// StructsConstantViews2 is an auto generated low-level Go binding around an user-defined struct.
type StructsConstantViews2 struct {
	Token0NumeratorPrecision   *big.Int
	Token0DenominatorPrecision *big.Int
	Token1NumeratorPrecision   *big.Int
	Token1DenominatorPrecision *big.Int
}

// StructsDebtReserves is an auto generated low-level Go binding around an user-defined struct.
type StructsDebtReserves struct {
	Token0Debt              *big.Int
	Token1Debt              *big.Int
	Token0RealReserves      *big.Int
	Token1RealReserves      *big.Int
	Token0ImaginaryReserves *big.Int
	Token1ImaginaryReserves *big.Int
}

// StructsImplementations is an auto generated low-level Go binding around an user-defined struct.
type StructsImplementations struct {
	Shift                       common.Address
	Admin                       common.Address
	ColOperations               common.Address
	DebtOperations              common.Address
	PerfectOperationsAndSwapOut common.Address
}

// StructsOracle is an auto generated low-level Go binding around an user-defined struct.
type StructsOracle struct {
	Twap1by0         *big.Int
	LowestPrice1by0  *big.Int
	HighestPrice1by0 *big.Int
	Twap0by1         *big.Int
	LowestPrice0by1  *big.Int
	HighestPrice0by1 *big.Int
}

// FluidDexT1MetaData contains all meta data concerning the FluidDexT1 contract.
var FluidDexT1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dexId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"liquidity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"shift\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"colOperations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtOperations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perfectOperationsAndSwapOut\",\"type\":\"address\"}],\"internalType\":\"structStructs.Implementations\",\"name\":\"implementations\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"deployerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"supplyToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"borrowToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"supplyToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"borrowToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"exchangePriceToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"exchangePriceToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"oracleMapping\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.ConstantViews\",\"name\":\"constantViews_\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorId_\",\"type\":\"uint256\"}],\"name\":\"FluidDexError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorId\",\"type\":\"uint256\"}],\"name\":\"FluidDexFactoryError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"FluidDexLiquidityOutput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"FluidDexPerfectLiquidityOutput\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lastStoredPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"centerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowerRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"geometricMean\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyToken0ExchangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowToken0ExchangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyToken1ExchangePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowToken1ExchangePrice\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.PricesAndExchangePrice\",\"name\":\"pex_\",\"type\":\"tuple\"}],\"name\":\"FluidDexPricesAndExchangeRates\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmt\",\"type\":\"uint256\"}],\"name\":\"FluidDexSingleTokenOutput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"FluidDexSwapResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorId_\",\"type\":\"uint256\"}],\"name\":\"FluidLiquidityCalcsError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"errorId_\",\"type\":\"uint256\"}],\"name\":\"FluidSafeTransferError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"routing\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amtOut\",\"type\":\"uint256\"}],\"name\":\"LogArbitrage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"LogBorrowDebtLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogBorrowPerfectDebtLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"LogDepositColLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogDepositPerfectColLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogPaybackDebtInOneToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"LogPaybackDebtLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogPaybackPerfectDebtLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawColInOneToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawColLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Amt\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawPerfectColLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swap0to1\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEX_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSharesAmt_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken0Borrow_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken1Borrow_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"borrowPerfect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constantsView\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dexId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"liquidity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"shift\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"colOperations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtOperations\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"perfectOperationsAndSwapOut\",\"type\":\"address\"}],\"internalType\":\"structStructs.Implementations\",\"name\":\"implementations\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"deployerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"supplyToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"borrowToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"supplyToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"borrowToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"exchangePriceToken0Slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"exchangePriceToken1Slot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"oracleMapping\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.ConstantViews\",\"name\":\"constantsView_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constantsView2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token0NumeratorPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0DenominatorPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1NumeratorPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1DenominatorPrecision\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.ConstantViews2\",\"name\":\"constantsView2_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSharesAmt_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"estimate_\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken0Deposit_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken1Deposit_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"estimate_\",\"type\":\"bool\"}],\"name\":\"depositPerfect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"geometricMean_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperRange_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowerRange_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0SupplyExchangePrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1SupplyExchangePrice_\",\"type\":\"uint256\"}],\"name\":\"getCollateralReserves\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token0RealReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1RealReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0ImaginaryReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1ImaginaryReserves\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.CollateralReserves\",\"name\":\"c_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"geometricMean_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperRange_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowerRange_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0BorrowExchangePrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1BorrowExchangePrice_\",\"type\":\"uint256\"}],\"name\":\"getDebtReserves\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token0Debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0RealReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1RealReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token0ImaginaryReserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1ImaginaryReserves\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.DebtReserves\",\"name\":\"d_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricesAndExchangePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"liquidityCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"secondsAgos_\",\"type\":\"uint256[]\"}],\"name\":\"oraclePrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"twap1by0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowestPrice1by0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestPrice1by0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"twap0by1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lowestPrice0by1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestPrice0by1\",\"type\":\"uint256\"}],\"internalType\":\"structStructs.Oracle[]\",\"name\":\"twaps_\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"currentPrice_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSharesAmt_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"estimate_\",\"type\":\"bool\"}],\"name\":\"payback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken0Payback_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken1Payback_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"estimate_\",\"type\":\"bool\"}],\"name\":\"paybackPerfect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken0_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxToken1_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"estimate_\",\"type\":\"bool\"}],\"name\":\"paybackPerfectInOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paybackAmt_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot_\",\"type\":\"bytes32\"}],\"name\":\"readFromStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swap0to1_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"swapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swap0to1_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"swapInWithCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swap0to1_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"swapOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swap0to1_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"swapOutWithCallback\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSharesAmt_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken0Withdraw_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken1Withdraw_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"withdrawPerfect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Amt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Amt_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken0_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minToken1_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"}],\"name\":\"withdrawPerfectInOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawAmt_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FluidDexT1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FluidDexT1MetaData.ABI instead.
var FluidDexT1ABI = FluidDexT1MetaData.ABI

// FluidDexT1 is an auto generated Go binding around an Ethereum contract.
type FluidDexT1 struct {
	FluidDexT1Caller     // Read-only binding to the contract
	FluidDexT1Transactor // Write-only binding to the contract
	FluidDexT1Filterer   // Log filterer for contract events
}

// FluidDexT1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FluidDexT1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluidDexT1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FluidDexT1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluidDexT1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FluidDexT1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FluidDexT1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FluidDexT1Session struct {
	Contract     *FluidDexT1       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FluidDexT1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FluidDexT1CallerSession struct {
	Contract *FluidDexT1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FluidDexT1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FluidDexT1TransactorSession struct {
	Contract     *FluidDexT1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FluidDexT1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FluidDexT1Raw struct {
	Contract *FluidDexT1 // Generic contract binding to access the raw methods on
}

// FluidDexT1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FluidDexT1CallerRaw struct {
	Contract *FluidDexT1Caller // Generic read-only contract binding to access the raw methods on
}

// FluidDexT1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FluidDexT1TransactorRaw struct {
	Contract *FluidDexT1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFluidDexT1 creates a new instance of FluidDexT1, bound to a specific deployed contract.
func NewFluidDexT1(address common.Address, backend bind.ContractBackend) (*FluidDexT1, error) {
	contract, err := bindFluidDexT1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FluidDexT1{FluidDexT1Caller: FluidDexT1Caller{contract: contract}, FluidDexT1Transactor: FluidDexT1Transactor{contract: contract}, FluidDexT1Filterer: FluidDexT1Filterer{contract: contract}}, nil
}

// NewFluidDexT1Caller creates a new read-only instance of FluidDexT1, bound to a specific deployed contract.
func NewFluidDexT1Caller(address common.Address, caller bind.ContractCaller) (*FluidDexT1Caller, error) {
	contract, err := bindFluidDexT1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FluidDexT1Caller{contract: contract}, nil
}

// NewFluidDexT1Transactor creates a new write-only instance of FluidDexT1, bound to a specific deployed contract.
func NewFluidDexT1Transactor(address common.Address, transactor bind.ContractTransactor) (*FluidDexT1Transactor, error) {
	contract, err := bindFluidDexT1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FluidDexT1Transactor{contract: contract}, nil
}

// NewFluidDexT1Filterer creates a new log filterer instance of FluidDexT1, bound to a specific deployed contract.
func NewFluidDexT1Filterer(address common.Address, filterer bind.ContractFilterer) (*FluidDexT1Filterer, error) {
	contract, err := bindFluidDexT1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FluidDexT1Filterer{contract: contract}, nil
}

// bindFluidDexT1 binds a generic wrapper to an already deployed contract.
func bindFluidDexT1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FluidDexT1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FluidDexT1 *FluidDexT1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FluidDexT1.Contract.FluidDexT1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FluidDexT1 *FluidDexT1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluidDexT1.Contract.FluidDexT1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FluidDexT1 *FluidDexT1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FluidDexT1.Contract.FluidDexT1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FluidDexT1 *FluidDexT1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FluidDexT1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FluidDexT1 *FluidDexT1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluidDexT1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FluidDexT1 *FluidDexT1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FluidDexT1.Contract.contract.Transact(opts, method, params...)
}

// DEXID is a free data retrieval call binding the contract method 0xf4b9a3fb.
//
// Solidity: function DEX_ID() view returns(uint256)
func (_FluidDexT1 *FluidDexT1Caller) DEXID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "DEX_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEXID is a free data retrieval call binding the contract method 0xf4b9a3fb.
//
// Solidity: function DEX_ID() view returns(uint256)
func (_FluidDexT1 *FluidDexT1Session) DEXID() (*big.Int, error) {
	return _FluidDexT1.Contract.DEXID(&_FluidDexT1.CallOpts)
}

// DEXID is a free data retrieval call binding the contract method 0xf4b9a3fb.
//
// Solidity: function DEX_ID() view returns(uint256)
func (_FluidDexT1 *FluidDexT1CallerSession) DEXID() (*big.Int, error) {
	return _FluidDexT1.Contract.DEXID(&_FluidDexT1.CallOpts)
}

// ConstantsView is a free data retrieval call binding the contract method 0xb7791bf2.
//
// Solidity: function constantsView() view returns((uint256,address,address,(address,address,address,address,address),address,address,address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256) constantsView_)
func (_FluidDexT1 *FluidDexT1Caller) ConstantsView(opts *bind.CallOpts) (StructsConstantViews, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "constantsView")

	if err != nil {
		return *new(StructsConstantViews), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsConstantViews)).(*StructsConstantViews)

	return out0, err

}

// ConstantsView is a free data retrieval call binding the contract method 0xb7791bf2.
//
// Solidity: function constantsView() view returns((uint256,address,address,(address,address,address,address,address),address,address,address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256) constantsView_)
func (_FluidDexT1 *FluidDexT1Session) ConstantsView() (StructsConstantViews, error) {
	return _FluidDexT1.Contract.ConstantsView(&_FluidDexT1.CallOpts)
}

// ConstantsView is a free data retrieval call binding the contract method 0xb7791bf2.
//
// Solidity: function constantsView() view returns((uint256,address,address,(address,address,address,address,address),address,address,address,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,uint256) constantsView_)
func (_FluidDexT1 *FluidDexT1CallerSession) ConstantsView() (StructsConstantViews, error) {
	return _FluidDexT1.Contract.ConstantsView(&_FluidDexT1.CallOpts)
}

// ConstantsView2 is a free data retrieval call binding the contract method 0x1595cbd3.
//
// Solidity: function constantsView2() view returns((uint256,uint256,uint256,uint256) constantsView2_)
func (_FluidDexT1 *FluidDexT1Caller) ConstantsView2(opts *bind.CallOpts) (StructsConstantViews2, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "constantsView2")

	if err != nil {
		return *new(StructsConstantViews2), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsConstantViews2)).(*StructsConstantViews2)

	return out0, err

}

// ConstantsView2 is a free data retrieval call binding the contract method 0x1595cbd3.
//
// Solidity: function constantsView2() view returns((uint256,uint256,uint256,uint256) constantsView2_)
func (_FluidDexT1 *FluidDexT1Session) ConstantsView2() (StructsConstantViews2, error) {
	return _FluidDexT1.Contract.ConstantsView2(&_FluidDexT1.CallOpts)
}

// ConstantsView2 is a free data retrieval call binding the contract method 0x1595cbd3.
//
// Solidity: function constantsView2() view returns((uint256,uint256,uint256,uint256) constantsView2_)
func (_FluidDexT1 *FluidDexT1CallerSession) ConstantsView2() (StructsConstantViews2, error) {
	return _FluidDexT1.Contract.ConstantsView2(&_FluidDexT1.CallOpts)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x6560abaa.
//
// Solidity: function getCollateralReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0SupplyExchangePrice_, uint256 token1SupplyExchangePrice_) view returns((uint256,uint256,uint256,uint256) c_)
func (_FluidDexT1 *FluidDexT1Caller) GetCollateralReserves(opts *bind.CallOpts, geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0SupplyExchangePrice_ *big.Int, token1SupplyExchangePrice_ *big.Int) (StructsCollateralReserves, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "getCollateralReserves", geometricMean_, upperRange_, lowerRange_, token0SupplyExchangePrice_, token1SupplyExchangePrice_)

	if err != nil {
		return *new(StructsCollateralReserves), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsCollateralReserves)).(*StructsCollateralReserves)

	return out0, err

}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x6560abaa.
//
// Solidity: function getCollateralReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0SupplyExchangePrice_, uint256 token1SupplyExchangePrice_) view returns((uint256,uint256,uint256,uint256) c_)
func (_FluidDexT1 *FluidDexT1Session) GetCollateralReserves(geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0SupplyExchangePrice_ *big.Int, token1SupplyExchangePrice_ *big.Int) (StructsCollateralReserves, error) {
	return _FluidDexT1.Contract.GetCollateralReserves(&_FluidDexT1.CallOpts, geometricMean_, upperRange_, lowerRange_, token0SupplyExchangePrice_, token1SupplyExchangePrice_)
}

// GetCollateralReserves is a free data retrieval call binding the contract method 0x6560abaa.
//
// Solidity: function getCollateralReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0SupplyExchangePrice_, uint256 token1SupplyExchangePrice_) view returns((uint256,uint256,uint256,uint256) c_)
func (_FluidDexT1 *FluidDexT1CallerSession) GetCollateralReserves(geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0SupplyExchangePrice_ *big.Int, token1SupplyExchangePrice_ *big.Int) (StructsCollateralReserves, error) {
	return _FluidDexT1.Contract.GetCollateralReserves(&_FluidDexT1.CallOpts, geometricMean_, upperRange_, lowerRange_, token0SupplyExchangePrice_, token1SupplyExchangePrice_)
}

// GetDebtReserves is a free data retrieval call binding the contract method 0x05d455a9.
//
// Solidity: function getDebtReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0BorrowExchangePrice_, uint256 token1BorrowExchangePrice_) view returns((uint256,uint256,uint256,uint256,uint256,uint256) d_)
func (_FluidDexT1 *FluidDexT1Caller) GetDebtReserves(opts *bind.CallOpts, geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0BorrowExchangePrice_ *big.Int, token1BorrowExchangePrice_ *big.Int) (StructsDebtReserves, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "getDebtReserves", geometricMean_, upperRange_, lowerRange_, token0BorrowExchangePrice_, token1BorrowExchangePrice_)

	if err != nil {
		return *new(StructsDebtReserves), err
	}

	out0 := *abi.ConvertType(out[0], new(StructsDebtReserves)).(*StructsDebtReserves)

	return out0, err

}

// GetDebtReserves is a free data retrieval call binding the contract method 0x05d455a9.
//
// Solidity: function getDebtReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0BorrowExchangePrice_, uint256 token1BorrowExchangePrice_) view returns((uint256,uint256,uint256,uint256,uint256,uint256) d_)
func (_FluidDexT1 *FluidDexT1Session) GetDebtReserves(geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0BorrowExchangePrice_ *big.Int, token1BorrowExchangePrice_ *big.Int) (StructsDebtReserves, error) {
	return _FluidDexT1.Contract.GetDebtReserves(&_FluidDexT1.CallOpts, geometricMean_, upperRange_, lowerRange_, token0BorrowExchangePrice_, token1BorrowExchangePrice_)
}

// GetDebtReserves is a free data retrieval call binding the contract method 0x05d455a9.
//
// Solidity: function getDebtReserves(uint256 geometricMean_, uint256 upperRange_, uint256 lowerRange_, uint256 token0BorrowExchangePrice_, uint256 token1BorrowExchangePrice_) view returns((uint256,uint256,uint256,uint256,uint256,uint256) d_)
func (_FluidDexT1 *FluidDexT1CallerSession) GetDebtReserves(geometricMean_ *big.Int, upperRange_ *big.Int, lowerRange_ *big.Int, token0BorrowExchangePrice_ *big.Int, token1BorrowExchangePrice_ *big.Int) (StructsDebtReserves, error) {
	return _FluidDexT1.Contract.GetDebtReserves(&_FluidDexT1.CallOpts, geometricMean_, upperRange_, lowerRange_, token0BorrowExchangePrice_, token1BorrowExchangePrice_)
}

// OraclePrice is a free data retrieval call binding the contract method 0xd811b2ce.
//
// Solidity: function oraclePrice(uint256[] secondsAgos_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] twaps_, uint256 currentPrice_)
func (_FluidDexT1 *FluidDexT1Caller) OraclePrice(opts *bind.CallOpts, secondsAgos_ []*big.Int) (struct {
	Twaps        []StructsOracle
	CurrentPrice *big.Int
}, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "oraclePrice", secondsAgos_)

	outstruct := new(struct {
		Twaps        []StructsOracle
		CurrentPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Twaps = *abi.ConvertType(out[0], new([]StructsOracle)).(*[]StructsOracle)
	outstruct.CurrentPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OraclePrice is a free data retrieval call binding the contract method 0xd811b2ce.
//
// Solidity: function oraclePrice(uint256[] secondsAgos_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] twaps_, uint256 currentPrice_)
func (_FluidDexT1 *FluidDexT1Session) OraclePrice(secondsAgos_ []*big.Int) (struct {
	Twaps        []StructsOracle
	CurrentPrice *big.Int
}, error) {
	return _FluidDexT1.Contract.OraclePrice(&_FluidDexT1.CallOpts, secondsAgos_)
}

// OraclePrice is a free data retrieval call binding the contract method 0xd811b2ce.
//
// Solidity: function oraclePrice(uint256[] secondsAgos_) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] twaps_, uint256 currentPrice_)
func (_FluidDexT1 *FluidDexT1CallerSession) OraclePrice(secondsAgos_ []*big.Int) (struct {
	Twaps        []StructsOracle
	CurrentPrice *big.Int
}, error) {
	return _FluidDexT1.Contract.OraclePrice(&_FluidDexT1.CallOpts, secondsAgos_)
}

// ReadFromStorage is a free data retrieval call binding the contract method 0xb5c736e4.
//
// Solidity: function readFromStorage(bytes32 slot_) view returns(uint256 result_)
func (_FluidDexT1 *FluidDexT1Caller) ReadFromStorage(opts *bind.CallOpts, slot_ [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FluidDexT1.contract.Call(opts, &out, "readFromStorage", slot_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadFromStorage is a free data retrieval call binding the contract method 0xb5c736e4.
//
// Solidity: function readFromStorage(bytes32 slot_) view returns(uint256 result_)
func (_FluidDexT1 *FluidDexT1Session) ReadFromStorage(slot_ [32]byte) (*big.Int, error) {
	return _FluidDexT1.Contract.ReadFromStorage(&_FluidDexT1.CallOpts, slot_)
}

// ReadFromStorage is a free data retrieval call binding the contract method 0xb5c736e4.
//
// Solidity: function readFromStorage(bytes32 slot_) view returns(uint256 result_)
func (_FluidDexT1 *FluidDexT1CallerSession) ReadFromStorage(slot_ [32]byte) (*big.Int, error) {
	return _FluidDexT1.Contract.ReadFromStorage(&_FluidDexT1.CallOpts, slot_)
}

// Borrow is a paid mutator transaction binding the contract method 0x242011d5.
//
// Solidity: function borrow(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Transactor) Borrow(opts *bind.TransactOpts, token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "borrow", token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// Borrow is a paid mutator transaction binding the contract method 0x242011d5.
//
// Solidity: function borrow(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Session) Borrow(token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Borrow(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// Borrow is a paid mutator transaction binding the contract method 0x242011d5.
//
// Solidity: function borrow(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1TransactorSession) Borrow(token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Borrow(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// BorrowPerfect is a paid mutator transaction binding the contract method 0xe27203cd.
//
// Solidity: function borrowPerfect(uint256 shares_, uint256 minToken0Borrow_, uint256 minToken1Borrow_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Transactor) BorrowPerfect(opts *bind.TransactOpts, shares_ *big.Int, minToken0Borrow_ *big.Int, minToken1Borrow_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "borrowPerfect", shares_, minToken0Borrow_, minToken1Borrow_, to_)
}

// BorrowPerfect is a paid mutator transaction binding the contract method 0xe27203cd.
//
// Solidity: function borrowPerfect(uint256 shares_, uint256 minToken0Borrow_, uint256 minToken1Borrow_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Session) BorrowPerfect(shares_ *big.Int, minToken0Borrow_ *big.Int, minToken1Borrow_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.BorrowPerfect(&_FluidDexT1.TransactOpts, shares_, minToken0Borrow_, minToken1Borrow_, to_)
}

// BorrowPerfect is a paid mutator transaction binding the contract method 0xe27203cd.
//
// Solidity: function borrowPerfect(uint256 shares_, uint256 minToken0Borrow_, uint256 minToken1Borrow_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) BorrowPerfect(shares_ *big.Int, minToken0Borrow_ *big.Int, minToken1Borrow_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.BorrowPerfect(&_FluidDexT1.TransactOpts, shares_, minToken0Borrow_, minToken1Borrow_, to_)
}

// Deposit is a paid mutator transaction binding the contract method 0xe980e1eb.
//
// Solidity: function deposit(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Transactor) Deposit(opts *bind.TransactOpts, token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "deposit", token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// Deposit is a paid mutator transaction binding the contract method 0xe980e1eb.
//
// Solidity: function deposit(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Session) Deposit(token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Deposit(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// Deposit is a paid mutator transaction binding the contract method 0xe980e1eb.
//
// Solidity: function deposit(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1TransactorSession) Deposit(token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Deposit(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// DepositPerfect is a paid mutator transaction binding the contract method 0x4d9036de.
//
// Solidity: function depositPerfect(uint256 shares_, uint256 maxToken0Deposit_, uint256 maxToken1Deposit_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Transactor) DepositPerfect(opts *bind.TransactOpts, shares_ *big.Int, maxToken0Deposit_ *big.Int, maxToken1Deposit_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "depositPerfect", shares_, maxToken0Deposit_, maxToken1Deposit_, estimate_)
}

// DepositPerfect is a paid mutator transaction binding the contract method 0x4d9036de.
//
// Solidity: function depositPerfect(uint256 shares_, uint256 maxToken0Deposit_, uint256 maxToken1Deposit_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Session) DepositPerfect(shares_ *big.Int, maxToken0Deposit_ *big.Int, maxToken1Deposit_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.DepositPerfect(&_FluidDexT1.TransactOpts, shares_, maxToken0Deposit_, maxToken1Deposit_, estimate_)
}

// DepositPerfect is a paid mutator transaction binding the contract method 0x4d9036de.
//
// Solidity: function depositPerfect(uint256 shares_, uint256 maxToken0Deposit_, uint256 maxToken1Deposit_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) DepositPerfect(shares_ *big.Int, maxToken0Deposit_ *big.Int, maxToken1Deposit_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.DepositPerfect(&_FluidDexT1.TransactOpts, shares_, maxToken0Deposit_, maxToken1Deposit_, estimate_)
}

// GetPricesAndExchangePrices is a paid mutator transaction binding the contract method 0x916cef4e.
//
// Solidity: function getPricesAndExchangePrices() returns()
func (_FluidDexT1 *FluidDexT1Transactor) GetPricesAndExchangePrices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "getPricesAndExchangePrices")
}

// GetPricesAndExchangePrices is a paid mutator transaction binding the contract method 0x916cef4e.
//
// Solidity: function getPricesAndExchangePrices() returns()
func (_FluidDexT1 *FluidDexT1Session) GetPricesAndExchangePrices() (*types.Transaction, error) {
	return _FluidDexT1.Contract.GetPricesAndExchangePrices(&_FluidDexT1.TransactOpts)
}

// GetPricesAndExchangePrices is a paid mutator transaction binding the contract method 0x916cef4e.
//
// Solidity: function getPricesAndExchangePrices() returns()
func (_FluidDexT1 *FluidDexT1TransactorSession) GetPricesAndExchangePrices() (*types.Transaction, error) {
	return _FluidDexT1.Contract.GetPricesAndExchangePrices(&_FluidDexT1.TransactOpts)
}

// LiquidityCallback is a paid mutator transaction binding the contract method 0xad207501.
//
// Solidity: function liquidityCallback(address token_, uint256 amount_, bytes data_) returns()
func (_FluidDexT1 *FluidDexT1Transactor) LiquidityCallback(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "liquidityCallback", token_, amount_, data_)
}

// LiquidityCallback is a paid mutator transaction binding the contract method 0xad207501.
//
// Solidity: function liquidityCallback(address token_, uint256 amount_, bytes data_) returns()
func (_FluidDexT1 *FluidDexT1Session) LiquidityCallback(token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _FluidDexT1.Contract.LiquidityCallback(&_FluidDexT1.TransactOpts, token_, amount_, data_)
}

// LiquidityCallback is a paid mutator transaction binding the contract method 0xad207501.
//
// Solidity: function liquidityCallback(address token_, uint256 amount_, bytes data_) returns()
func (_FluidDexT1 *FluidDexT1TransactorSession) LiquidityCallback(token_ common.Address, amount_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _FluidDexT1.Contract.LiquidityCallback(&_FluidDexT1.TransactOpts, token_, amount_, data_)
}

// Payback is a paid mutator transaction binding the contract method 0x68766981.
//
// Solidity: function payback(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Transactor) Payback(opts *bind.TransactOpts, token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "payback", token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// Payback is a paid mutator transaction binding the contract method 0x68766981.
//
// Solidity: function payback(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Session) Payback(token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Payback(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// Payback is a paid mutator transaction binding the contract method 0x68766981.
//
// Solidity: function payback(uint256 token0Amt_, uint256 token1Amt_, uint256 minSharesAmt_, bool estimate_) payable returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1TransactorSession) Payback(token0Amt_ *big.Int, token1Amt_ *big.Int, minSharesAmt_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Payback(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, minSharesAmt_, estimate_)
}

// PaybackPerfect is a paid mutator transaction binding the contract method 0x5b3d38d7.
//
// Solidity: function paybackPerfect(uint256 shares_, uint256 maxToken0Payback_, uint256 maxToken1Payback_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Transactor) PaybackPerfect(opts *bind.TransactOpts, shares_ *big.Int, maxToken0Payback_ *big.Int, maxToken1Payback_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "paybackPerfect", shares_, maxToken0Payback_, maxToken1Payback_, estimate_)
}

// PaybackPerfect is a paid mutator transaction binding the contract method 0x5b3d38d7.
//
// Solidity: function paybackPerfect(uint256 shares_, uint256 maxToken0Payback_, uint256 maxToken1Payback_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Session) PaybackPerfect(shares_ *big.Int, maxToken0Payback_ *big.Int, maxToken1Payback_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.PaybackPerfect(&_FluidDexT1.TransactOpts, shares_, maxToken0Payback_, maxToken1Payback_, estimate_)
}

// PaybackPerfect is a paid mutator transaction binding the contract method 0x5b3d38d7.
//
// Solidity: function paybackPerfect(uint256 shares_, uint256 maxToken0Payback_, uint256 maxToken1Payback_, bool estimate_) payable returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) PaybackPerfect(shares_ *big.Int, maxToken0Payback_ *big.Int, maxToken1Payback_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.PaybackPerfect(&_FluidDexT1.TransactOpts, shares_, maxToken0Payback_, maxToken1Payback_, estimate_)
}

// PaybackPerfectInOneToken is a paid mutator transaction binding the contract method 0x30acd6fd.
//
// Solidity: function paybackPerfectInOneToken(uint256 shares_, uint256 maxToken0_, uint256 maxToken1_, bool estimate_) payable returns(uint256 paybackAmt_)
func (_FluidDexT1 *FluidDexT1Transactor) PaybackPerfectInOneToken(opts *bind.TransactOpts, shares_ *big.Int, maxToken0_ *big.Int, maxToken1_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "paybackPerfectInOneToken", shares_, maxToken0_, maxToken1_, estimate_)
}

// PaybackPerfectInOneToken is a paid mutator transaction binding the contract method 0x30acd6fd.
//
// Solidity: function paybackPerfectInOneToken(uint256 shares_, uint256 maxToken0_, uint256 maxToken1_, bool estimate_) payable returns(uint256 paybackAmt_)
func (_FluidDexT1 *FluidDexT1Session) PaybackPerfectInOneToken(shares_ *big.Int, maxToken0_ *big.Int, maxToken1_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.PaybackPerfectInOneToken(&_FluidDexT1.TransactOpts, shares_, maxToken0_, maxToken1_, estimate_)
}

// PaybackPerfectInOneToken is a paid mutator transaction binding the contract method 0x30acd6fd.
//
// Solidity: function paybackPerfectInOneToken(uint256 shares_, uint256 maxToken0_, uint256 maxToken1_, bool estimate_) payable returns(uint256 paybackAmt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) PaybackPerfectInOneToken(shares_ *big.Int, maxToken0_ *big.Int, maxToken1_ *big.Int, estimate_ bool) (*types.Transaction, error) {
	return _FluidDexT1.Contract.PaybackPerfectInOneToken(&_FluidDexT1.TransactOpts, shares_, maxToken0_, maxToken1_, estimate_)
}

// SwapIn is a paid mutator transaction binding the contract method 0x2668dfaa.
//
// Solidity: function swapIn(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1Transactor) SwapIn(opts *bind.TransactOpts, swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "swapIn", swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapIn is a paid mutator transaction binding the contract method 0x2668dfaa.
//
// Solidity: function swapIn(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1Session) SwapIn(swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapIn(&_FluidDexT1.TransactOpts, swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapIn is a paid mutator transaction binding the contract method 0x2668dfaa.
//
// Solidity: function swapIn(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1TransactorSession) SwapIn(swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapIn(&_FluidDexT1.TransactOpts, swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapInWithCallback is a paid mutator transaction binding the contract method 0xbe17c79c.
//
// Solidity: function swapInWithCallback(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1Transactor) SwapInWithCallback(opts *bind.TransactOpts, swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "swapInWithCallback", swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapInWithCallback is a paid mutator transaction binding the contract method 0xbe17c79c.
//
// Solidity: function swapInWithCallback(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1Session) SwapInWithCallback(swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapInWithCallback(&_FluidDexT1.TransactOpts, swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapInWithCallback is a paid mutator transaction binding the contract method 0xbe17c79c.
//
// Solidity: function swapInWithCallback(bool swap0to1_, uint256 amountIn_, uint256 amountOutMin_, address to_) payable returns(uint256 amountOut_)
func (_FluidDexT1 *FluidDexT1TransactorSession) SwapInWithCallback(swap0to1_ bool, amountIn_ *big.Int, amountOutMin_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapInWithCallback(&_FluidDexT1.TransactOpts, swap0to1_, amountIn_, amountOutMin_, to_)
}

// SwapOut is a paid mutator transaction binding the contract method 0x286f0e61.
//
// Solidity: function swapOut(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1Transactor) SwapOut(opts *bind.TransactOpts, swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "swapOut", swap0to1_, amountOut_, amountInMax_, to_)
}

// SwapOut is a paid mutator transaction binding the contract method 0x286f0e61.
//
// Solidity: function swapOut(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1Session) SwapOut(swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapOut(&_FluidDexT1.TransactOpts, swap0to1_, amountOut_, amountInMax_, to_)
}

// SwapOut is a paid mutator transaction binding the contract method 0x286f0e61.
//
// Solidity: function swapOut(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1TransactorSession) SwapOut(swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapOut(&_FluidDexT1.TransactOpts, swap0to1_, amountOut_, amountInMax_, to_)
}

// SwapOutWithCallback is a paid mutator transaction binding the contract method 0x653295aa.
//
// Solidity: function swapOutWithCallback(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1Transactor) SwapOutWithCallback(opts *bind.TransactOpts, swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "swapOutWithCallback", swap0to1_, amountOut_, amountInMax_, to_)
}

// SwapOutWithCallback is a paid mutator transaction binding the contract method 0x653295aa.
//
// Solidity: function swapOutWithCallback(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1Session) SwapOutWithCallback(swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapOutWithCallback(&_FluidDexT1.TransactOpts, swap0to1_, amountOut_, amountInMax_, to_)
}

// SwapOutWithCallback is a paid mutator transaction binding the contract method 0x653295aa.
//
// Solidity: function swapOutWithCallback(bool swap0to1_, uint256 amountOut_, uint256 amountInMax_, address to_) payable returns(uint256 amountIn_)
func (_FluidDexT1 *FluidDexT1TransactorSession) SwapOutWithCallback(swap0to1_ bool, amountOut_ *big.Int, amountInMax_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.SwapOutWithCallback(&_FluidDexT1.TransactOpts, swap0to1_, amountOut_, amountInMax_, to_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Transactor) Withdraw(opts *bind.TransactOpts, token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "withdraw", token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1Session) Withdraw(token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Withdraw(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 token0Amt_, uint256 token1Amt_, uint256 maxSharesAmt_, address to_) returns(uint256 shares_)
func (_FluidDexT1 *FluidDexT1TransactorSession) Withdraw(token0Amt_ *big.Int, token1Amt_ *big.Int, maxSharesAmt_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Withdraw(&_FluidDexT1.TransactOpts, token0Amt_, token1Amt_, maxSharesAmt_, to_)
}

// WithdrawPerfect is a paid mutator transaction binding the contract method 0x35f0df98.
//
// Solidity: function withdrawPerfect(uint256 shares_, uint256 minToken0Withdraw_, uint256 minToken1Withdraw_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Transactor) WithdrawPerfect(opts *bind.TransactOpts, shares_ *big.Int, minToken0Withdraw_ *big.Int, minToken1Withdraw_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "withdrawPerfect", shares_, minToken0Withdraw_, minToken1Withdraw_, to_)
}

// WithdrawPerfect is a paid mutator transaction binding the contract method 0x35f0df98.
//
// Solidity: function withdrawPerfect(uint256 shares_, uint256 minToken0Withdraw_, uint256 minToken1Withdraw_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1Session) WithdrawPerfect(shares_ *big.Int, minToken0Withdraw_ *big.Int, minToken1Withdraw_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.WithdrawPerfect(&_FluidDexT1.TransactOpts, shares_, minToken0Withdraw_, minToken1Withdraw_, to_)
}

// WithdrawPerfect is a paid mutator transaction binding the contract method 0x35f0df98.
//
// Solidity: function withdrawPerfect(uint256 shares_, uint256 minToken0Withdraw_, uint256 minToken1Withdraw_, address to_) returns(uint256 token0Amt_, uint256 token1Amt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) WithdrawPerfect(shares_ *big.Int, minToken0Withdraw_ *big.Int, minToken1Withdraw_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.WithdrawPerfect(&_FluidDexT1.TransactOpts, shares_, minToken0Withdraw_, minToken1Withdraw_, to_)
}

// WithdrawPerfectInOneToken is a paid mutator transaction binding the contract method 0x4c89bfd4.
//
// Solidity: function withdrawPerfectInOneToken(uint256 shares_, uint256 minToken0_, uint256 minToken1_, address to_) returns(uint256 withdrawAmt_)
func (_FluidDexT1 *FluidDexT1Transactor) WithdrawPerfectInOneToken(opts *bind.TransactOpts, shares_ *big.Int, minToken0_ *big.Int, minToken1_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.contract.Transact(opts, "withdrawPerfectInOneToken", shares_, minToken0_, minToken1_, to_)
}

// WithdrawPerfectInOneToken is a paid mutator transaction binding the contract method 0x4c89bfd4.
//
// Solidity: function withdrawPerfectInOneToken(uint256 shares_, uint256 minToken0_, uint256 minToken1_, address to_) returns(uint256 withdrawAmt_)
func (_FluidDexT1 *FluidDexT1Session) WithdrawPerfectInOneToken(shares_ *big.Int, minToken0_ *big.Int, minToken1_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.WithdrawPerfectInOneToken(&_FluidDexT1.TransactOpts, shares_, minToken0_, minToken1_, to_)
}

// WithdrawPerfectInOneToken is a paid mutator transaction binding the contract method 0x4c89bfd4.
//
// Solidity: function withdrawPerfectInOneToken(uint256 shares_, uint256 minToken0_, uint256 minToken1_, address to_) returns(uint256 withdrawAmt_)
func (_FluidDexT1 *FluidDexT1TransactorSession) WithdrawPerfectInOneToken(shares_ *big.Int, minToken0_ *big.Int, minToken1_ *big.Int, to_ common.Address) (*types.Transaction, error) {
	return _FluidDexT1.Contract.WithdrawPerfectInOneToken(&_FluidDexT1.TransactOpts, shares_, minToken0_, minToken1_, to_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FluidDexT1 *FluidDexT1Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FluidDexT1.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FluidDexT1 *FluidDexT1Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Fallback(&_FluidDexT1.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FluidDexT1 *FluidDexT1TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FluidDexT1.Contract.Fallback(&_FluidDexT1.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FluidDexT1 *FluidDexT1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FluidDexT1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FluidDexT1 *FluidDexT1Session) Receive() (*types.Transaction, error) {
	return _FluidDexT1.Contract.Receive(&_FluidDexT1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FluidDexT1 *FluidDexT1TransactorSession) Receive() (*types.Transaction, error) {
	return _FluidDexT1.Contract.Receive(&_FluidDexT1.TransactOpts)
}

// FluidDexT1LogArbitrageIterator is returned from FilterLogArbitrage and is used to iterate over the raw logs and unpacked data for LogArbitrage events raised by the FluidDexT1 contract.
type FluidDexT1LogArbitrageIterator struct {
	Event *FluidDexT1LogArbitrage // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogArbitrageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogArbitrage)
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
		it.Event = new(FluidDexT1LogArbitrage)
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
func (it *FluidDexT1LogArbitrageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogArbitrageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogArbitrage represents a LogArbitrage event raised by the FluidDexT1 contract.
type FluidDexT1LogArbitrage struct {
	Routing *big.Int
	AmtOut  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogArbitrage is a free log retrieval operation binding the contract event 0x063def03d41a2957d43156b97c271f3e4adea600722defb2cf6ebf9a27650056.
//
// Solidity: event LogArbitrage(int256 routing, uint256 amtOut)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogArbitrage(opts *bind.FilterOpts) (*FluidDexT1LogArbitrageIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogArbitrage")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogArbitrageIterator{contract: _FluidDexT1.contract, event: "LogArbitrage", logs: logs, sub: sub}, nil
}

// WatchLogArbitrage is a free log subscription operation binding the contract event 0x063def03d41a2957d43156b97c271f3e4adea600722defb2cf6ebf9a27650056.
//
// Solidity: event LogArbitrage(int256 routing, uint256 amtOut)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogArbitrage(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogArbitrage) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogArbitrage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogArbitrage)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogArbitrage", log); err != nil {
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

// ParseLogArbitrage is a log parse operation binding the contract event 0x063def03d41a2957d43156b97c271f3e4adea600722defb2cf6ebf9a27650056.
//
// Solidity: event LogArbitrage(int256 routing, uint256 amtOut)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogArbitrage(log types.Log) (*FluidDexT1LogArbitrage, error) {
	event := new(FluidDexT1LogArbitrage)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogArbitrage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogBorrowDebtLiquidityIterator is returned from FilterLogBorrowDebtLiquidity and is used to iterate over the raw logs and unpacked data for LogBorrowDebtLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogBorrowDebtLiquidityIterator struct {
	Event *FluidDexT1LogBorrowDebtLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogBorrowDebtLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogBorrowDebtLiquidity)
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
		it.Event = new(FluidDexT1LogBorrowDebtLiquidity)
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
func (it *FluidDexT1LogBorrowDebtLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogBorrowDebtLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogBorrowDebtLiquidity represents a LogBorrowDebtLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogBorrowDebtLiquidity struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogBorrowDebtLiquidity is a free log retrieval operation binding the contract event 0x7f81427bed699dc7e687c5ddae6061932938818f79fc0e68903d55ef75ca4561.
//
// Solidity: event LogBorrowDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogBorrowDebtLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogBorrowDebtLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogBorrowDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogBorrowDebtLiquidityIterator{contract: _FluidDexT1.contract, event: "LogBorrowDebtLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogBorrowDebtLiquidity is a free log subscription operation binding the contract event 0x7f81427bed699dc7e687c5ddae6061932938818f79fc0e68903d55ef75ca4561.
//
// Solidity: event LogBorrowDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogBorrowDebtLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogBorrowDebtLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogBorrowDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogBorrowDebtLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogBorrowDebtLiquidity", log); err != nil {
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

// ParseLogBorrowDebtLiquidity is a log parse operation binding the contract event 0x7f81427bed699dc7e687c5ddae6061932938818f79fc0e68903d55ef75ca4561.
//
// Solidity: event LogBorrowDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogBorrowDebtLiquidity(log types.Log) (*FluidDexT1LogBorrowDebtLiquidity, error) {
	event := new(FluidDexT1LogBorrowDebtLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogBorrowDebtLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogBorrowPerfectDebtLiquidityIterator is returned from FilterLogBorrowPerfectDebtLiquidity and is used to iterate over the raw logs and unpacked data for LogBorrowPerfectDebtLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogBorrowPerfectDebtLiquidityIterator struct {
	Event *FluidDexT1LogBorrowPerfectDebtLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogBorrowPerfectDebtLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogBorrowPerfectDebtLiquidity)
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
		it.Event = new(FluidDexT1LogBorrowPerfectDebtLiquidity)
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
func (it *FluidDexT1LogBorrowPerfectDebtLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogBorrowPerfectDebtLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogBorrowPerfectDebtLiquidity represents a LogBorrowPerfectDebtLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogBorrowPerfectDebtLiquidity struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogBorrowPerfectDebtLiquidity is a free log retrieval operation binding the contract event 0x486d991947a88580130ff5acd9ec54dc37fb8da4bf6ab78871d5cd6fa5816df7.
//
// Solidity: event LogBorrowPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogBorrowPerfectDebtLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogBorrowPerfectDebtLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogBorrowPerfectDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogBorrowPerfectDebtLiquidityIterator{contract: _FluidDexT1.contract, event: "LogBorrowPerfectDebtLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogBorrowPerfectDebtLiquidity is a free log subscription operation binding the contract event 0x486d991947a88580130ff5acd9ec54dc37fb8da4bf6ab78871d5cd6fa5816df7.
//
// Solidity: event LogBorrowPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogBorrowPerfectDebtLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogBorrowPerfectDebtLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogBorrowPerfectDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogBorrowPerfectDebtLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogBorrowPerfectDebtLiquidity", log); err != nil {
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

// ParseLogBorrowPerfectDebtLiquidity is a log parse operation binding the contract event 0x486d991947a88580130ff5acd9ec54dc37fb8da4bf6ab78871d5cd6fa5816df7.
//
// Solidity: event LogBorrowPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogBorrowPerfectDebtLiquidity(log types.Log) (*FluidDexT1LogBorrowPerfectDebtLiquidity, error) {
	event := new(FluidDexT1LogBorrowPerfectDebtLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogBorrowPerfectDebtLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogDepositColLiquidityIterator is returned from FilterLogDepositColLiquidity and is used to iterate over the raw logs and unpacked data for LogDepositColLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogDepositColLiquidityIterator struct {
	Event *FluidDexT1LogDepositColLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogDepositColLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogDepositColLiquidity)
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
		it.Event = new(FluidDexT1LogDepositColLiquidity)
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
func (it *FluidDexT1LogDepositColLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogDepositColLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogDepositColLiquidity represents a LogDepositColLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogDepositColLiquidity struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogDepositColLiquidity is a free log retrieval operation binding the contract event 0xbfea92097a2487d6a5ccf7b7adc36b6002238f3106568ba4359770f4b67365a4.
//
// Solidity: event LogDepositColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogDepositColLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogDepositColLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogDepositColLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogDepositColLiquidityIterator{contract: _FluidDexT1.contract, event: "LogDepositColLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogDepositColLiquidity is a free log subscription operation binding the contract event 0xbfea92097a2487d6a5ccf7b7adc36b6002238f3106568ba4359770f4b67365a4.
//
// Solidity: event LogDepositColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogDepositColLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogDepositColLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogDepositColLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogDepositColLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogDepositColLiquidity", log); err != nil {
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

// ParseLogDepositColLiquidity is a log parse operation binding the contract event 0xbfea92097a2487d6a5ccf7b7adc36b6002238f3106568ba4359770f4b67365a4.
//
// Solidity: event LogDepositColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogDepositColLiquidity(log types.Log) (*FluidDexT1LogDepositColLiquidity, error) {
	event := new(FluidDexT1LogDepositColLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogDepositColLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogDepositPerfectColLiquidityIterator is returned from FilterLogDepositPerfectColLiquidity and is used to iterate over the raw logs and unpacked data for LogDepositPerfectColLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogDepositPerfectColLiquidityIterator struct {
	Event *FluidDexT1LogDepositPerfectColLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogDepositPerfectColLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogDepositPerfectColLiquidity)
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
		it.Event = new(FluidDexT1LogDepositPerfectColLiquidity)
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
func (it *FluidDexT1LogDepositPerfectColLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogDepositPerfectColLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogDepositPerfectColLiquidity represents a LogDepositPerfectColLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogDepositPerfectColLiquidity struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogDepositPerfectColLiquidity is a free log retrieval operation binding the contract event 0x255672effa3d8ba46e409fc964ae332b84d3107ba3a5096b22734606519528a3.
//
// Solidity: event LogDepositPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogDepositPerfectColLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogDepositPerfectColLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogDepositPerfectColLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogDepositPerfectColLiquidityIterator{contract: _FluidDexT1.contract, event: "LogDepositPerfectColLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogDepositPerfectColLiquidity is a free log subscription operation binding the contract event 0x255672effa3d8ba46e409fc964ae332b84d3107ba3a5096b22734606519528a3.
//
// Solidity: event LogDepositPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogDepositPerfectColLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogDepositPerfectColLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogDepositPerfectColLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogDepositPerfectColLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogDepositPerfectColLiquidity", log); err != nil {
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

// ParseLogDepositPerfectColLiquidity is a log parse operation binding the contract event 0x255672effa3d8ba46e409fc964ae332b84d3107ba3a5096b22734606519528a3.
//
// Solidity: event LogDepositPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogDepositPerfectColLiquidity(log types.Log) (*FluidDexT1LogDepositPerfectColLiquidity, error) {
	event := new(FluidDexT1LogDepositPerfectColLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogDepositPerfectColLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogPaybackDebtInOneTokenIterator is returned from FilterLogPaybackDebtInOneToken and is used to iterate over the raw logs and unpacked data for LogPaybackDebtInOneToken events raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackDebtInOneTokenIterator struct {
	Event *FluidDexT1LogPaybackDebtInOneToken // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogPaybackDebtInOneTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogPaybackDebtInOneToken)
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
		it.Event = new(FluidDexT1LogPaybackDebtInOneToken)
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
func (it *FluidDexT1LogPaybackDebtInOneTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogPaybackDebtInOneTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogPaybackDebtInOneToken represents a LogPaybackDebtInOneToken event raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackDebtInOneToken struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogPaybackDebtInOneToken is a free log retrieval operation binding the contract event 0x97dfa84cbffcf65b8d034f057439472bc93868a66cc0e728c2faffb00f8b4923.
//
// Solidity: event LogPaybackDebtInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogPaybackDebtInOneToken(opts *bind.FilterOpts) (*FluidDexT1LogPaybackDebtInOneTokenIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogPaybackDebtInOneToken")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogPaybackDebtInOneTokenIterator{contract: _FluidDexT1.contract, event: "LogPaybackDebtInOneToken", logs: logs, sub: sub}, nil
}

// WatchLogPaybackDebtInOneToken is a free log subscription operation binding the contract event 0x97dfa84cbffcf65b8d034f057439472bc93868a66cc0e728c2faffb00f8b4923.
//
// Solidity: event LogPaybackDebtInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogPaybackDebtInOneToken(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogPaybackDebtInOneToken) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogPaybackDebtInOneToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogPaybackDebtInOneToken)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackDebtInOneToken", log); err != nil {
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

// ParseLogPaybackDebtInOneToken is a log parse operation binding the contract event 0x97dfa84cbffcf65b8d034f057439472bc93868a66cc0e728c2faffb00f8b4923.
//
// Solidity: event LogPaybackDebtInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogPaybackDebtInOneToken(log types.Log) (*FluidDexT1LogPaybackDebtInOneToken, error) {
	event := new(FluidDexT1LogPaybackDebtInOneToken)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackDebtInOneToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogPaybackDebtLiquidityIterator is returned from FilterLogPaybackDebtLiquidity and is used to iterate over the raw logs and unpacked data for LogPaybackDebtLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackDebtLiquidityIterator struct {
	Event *FluidDexT1LogPaybackDebtLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogPaybackDebtLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogPaybackDebtLiquidity)
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
		it.Event = new(FluidDexT1LogPaybackDebtLiquidity)
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
func (it *FluidDexT1LogPaybackDebtLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogPaybackDebtLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogPaybackDebtLiquidity represents a LogPaybackDebtLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackDebtLiquidity struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogPaybackDebtLiquidity is a free log retrieval operation binding the contract event 0xb69f152a70520703fe7ab4872a0cb3928386b68cf3c6c83c5d1fc08d196991e8.
//
// Solidity: event LogPaybackDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogPaybackDebtLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogPaybackDebtLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogPaybackDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogPaybackDebtLiquidityIterator{contract: _FluidDexT1.contract, event: "LogPaybackDebtLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogPaybackDebtLiquidity is a free log subscription operation binding the contract event 0xb69f152a70520703fe7ab4872a0cb3928386b68cf3c6c83c5d1fc08d196991e8.
//
// Solidity: event LogPaybackDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogPaybackDebtLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogPaybackDebtLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogPaybackDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogPaybackDebtLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackDebtLiquidity", log); err != nil {
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

// ParseLogPaybackDebtLiquidity is a log parse operation binding the contract event 0xb69f152a70520703fe7ab4872a0cb3928386b68cf3c6c83c5d1fc08d196991e8.
//
// Solidity: event LogPaybackDebtLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogPaybackDebtLiquidity(log types.Log) (*FluidDexT1LogPaybackDebtLiquidity, error) {
	event := new(FluidDexT1LogPaybackDebtLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackDebtLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogPaybackPerfectDebtLiquidityIterator is returned from FilterLogPaybackPerfectDebtLiquidity and is used to iterate over the raw logs and unpacked data for LogPaybackPerfectDebtLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackPerfectDebtLiquidityIterator struct {
	Event *FluidDexT1LogPaybackPerfectDebtLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogPaybackPerfectDebtLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogPaybackPerfectDebtLiquidity)
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
		it.Event = new(FluidDexT1LogPaybackPerfectDebtLiquidity)
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
func (it *FluidDexT1LogPaybackPerfectDebtLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogPaybackPerfectDebtLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogPaybackPerfectDebtLiquidity represents a LogPaybackPerfectDebtLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogPaybackPerfectDebtLiquidity struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogPaybackPerfectDebtLiquidity is a free log retrieval operation binding the contract event 0x03b77b44c2fe8816d55a2e4f90a87538c48d4148e35224c07049fd2304fa3a30.
//
// Solidity: event LogPaybackPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogPaybackPerfectDebtLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogPaybackPerfectDebtLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogPaybackPerfectDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogPaybackPerfectDebtLiquidityIterator{contract: _FluidDexT1.contract, event: "LogPaybackPerfectDebtLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogPaybackPerfectDebtLiquidity is a free log subscription operation binding the contract event 0x03b77b44c2fe8816d55a2e4f90a87538c48d4148e35224c07049fd2304fa3a30.
//
// Solidity: event LogPaybackPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogPaybackPerfectDebtLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogPaybackPerfectDebtLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogPaybackPerfectDebtLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogPaybackPerfectDebtLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackPerfectDebtLiquidity", log); err != nil {
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

// ParseLogPaybackPerfectDebtLiquidity is a log parse operation binding the contract event 0x03b77b44c2fe8816d55a2e4f90a87538c48d4148e35224c07049fd2304fa3a30.
//
// Solidity: event LogPaybackPerfectDebtLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogPaybackPerfectDebtLiquidity(log types.Log) (*FluidDexT1LogPaybackPerfectDebtLiquidity, error) {
	event := new(FluidDexT1LogPaybackPerfectDebtLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogPaybackPerfectDebtLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogWithdrawColInOneTokenIterator is returned from FilterLogWithdrawColInOneToken and is used to iterate over the raw logs and unpacked data for LogWithdrawColInOneToken events raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawColInOneTokenIterator struct {
	Event *FluidDexT1LogWithdrawColInOneToken // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogWithdrawColInOneTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogWithdrawColInOneToken)
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
		it.Event = new(FluidDexT1LogWithdrawColInOneToken)
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
func (it *FluidDexT1LogWithdrawColInOneTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogWithdrawColInOneTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogWithdrawColInOneToken represents a LogWithdrawColInOneToken event raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawColInOneToken struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawColInOneToken is a free log retrieval operation binding the contract event 0xc98f37914e06db36c18654484db85c4bb864575a1b9f8181133ff33dea2d34f3.
//
// Solidity: event LogWithdrawColInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogWithdrawColInOneToken(opts *bind.FilterOpts) (*FluidDexT1LogWithdrawColInOneTokenIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogWithdrawColInOneToken")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogWithdrawColInOneTokenIterator{contract: _FluidDexT1.contract, event: "LogWithdrawColInOneToken", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawColInOneToken is a free log subscription operation binding the contract event 0xc98f37914e06db36c18654484db85c4bb864575a1b9f8181133ff33dea2d34f3.
//
// Solidity: event LogWithdrawColInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogWithdrawColInOneToken(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogWithdrawColInOneToken) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogWithdrawColInOneToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogWithdrawColInOneToken)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawColInOneToken", log); err != nil {
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

// ParseLogWithdrawColInOneToken is a log parse operation binding the contract event 0xc98f37914e06db36c18654484db85c4bb864575a1b9f8181133ff33dea2d34f3.
//
// Solidity: event LogWithdrawColInOneToken(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogWithdrawColInOneToken(log types.Log) (*FluidDexT1LogWithdrawColInOneToken, error) {
	event := new(FluidDexT1LogWithdrawColInOneToken)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawColInOneToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogWithdrawColLiquidityIterator is returned from FilterLogWithdrawColLiquidity and is used to iterate over the raw logs and unpacked data for LogWithdrawColLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawColLiquidityIterator struct {
	Event *FluidDexT1LogWithdrawColLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogWithdrawColLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogWithdrawColLiquidity)
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
		it.Event = new(FluidDexT1LogWithdrawColLiquidity)
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
func (it *FluidDexT1LogWithdrawColLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogWithdrawColLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogWithdrawColLiquidity represents a LogWithdrawColLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawColLiquidity struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawColLiquidity is a free log retrieval operation binding the contract event 0xb61c7f3b23fe9335cc6c6a6e7036457758470877e61a19a5b4924e1ff8289624.
//
// Solidity: event LogWithdrawColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogWithdrawColLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogWithdrawColLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogWithdrawColLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogWithdrawColLiquidityIterator{contract: _FluidDexT1.contract, event: "LogWithdrawColLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawColLiquidity is a free log subscription operation binding the contract event 0xb61c7f3b23fe9335cc6c6a6e7036457758470877e61a19a5b4924e1ff8289624.
//
// Solidity: event LogWithdrawColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogWithdrawColLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogWithdrawColLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogWithdrawColLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogWithdrawColLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawColLiquidity", log); err != nil {
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

// ParseLogWithdrawColLiquidity is a log parse operation binding the contract event 0xb61c7f3b23fe9335cc6c6a6e7036457758470877e61a19a5b4924e1ff8289624.
//
// Solidity: event LogWithdrawColLiquidity(uint256 amount0, uint256 amount1, uint256 shares)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogWithdrawColLiquidity(log types.Log) (*FluidDexT1LogWithdrawColLiquidity, error) {
	event := new(FluidDexT1LogWithdrawColLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawColLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1LogWithdrawPerfectColLiquidityIterator is returned from FilterLogWithdrawPerfectColLiquidity and is used to iterate over the raw logs and unpacked data for LogWithdrawPerfectColLiquidity events raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawPerfectColLiquidityIterator struct {
	Event *FluidDexT1LogWithdrawPerfectColLiquidity // Event containing the contract specifics and raw log

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
func (it *FluidDexT1LogWithdrawPerfectColLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1LogWithdrawPerfectColLiquidity)
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
		it.Event = new(FluidDexT1LogWithdrawPerfectColLiquidity)
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
func (it *FluidDexT1LogWithdrawPerfectColLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1LogWithdrawPerfectColLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1LogWithdrawPerfectColLiquidity represents a LogWithdrawPerfectColLiquidity event raised by the FluidDexT1 contract.
type FluidDexT1LogWithdrawPerfectColLiquidity struct {
	Shares    *big.Int
	Token0Amt *big.Int
	Token1Amt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawPerfectColLiquidity is a free log retrieval operation binding the contract event 0x6f837572c1ef6e010a841ff938d593ec054984fefe29df2a0634bbf01f4db35b.
//
// Solidity: event LogWithdrawPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) FilterLogWithdrawPerfectColLiquidity(opts *bind.FilterOpts) (*FluidDexT1LogWithdrawPerfectColLiquidityIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "LogWithdrawPerfectColLiquidity")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1LogWithdrawPerfectColLiquidityIterator{contract: _FluidDexT1.contract, event: "LogWithdrawPerfectColLiquidity", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawPerfectColLiquidity is a free log subscription operation binding the contract event 0x6f837572c1ef6e010a841ff938d593ec054984fefe29df2a0634bbf01f4db35b.
//
// Solidity: event LogWithdrawPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) WatchLogWithdrawPerfectColLiquidity(opts *bind.WatchOpts, sink chan<- *FluidDexT1LogWithdrawPerfectColLiquidity) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "LogWithdrawPerfectColLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1LogWithdrawPerfectColLiquidity)
				if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawPerfectColLiquidity", log); err != nil {
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

// ParseLogWithdrawPerfectColLiquidity is a log parse operation binding the contract event 0x6f837572c1ef6e010a841ff938d593ec054984fefe29df2a0634bbf01f4db35b.
//
// Solidity: event LogWithdrawPerfectColLiquidity(uint256 shares, uint256 token0Amt, uint256 token1Amt)
func (_FluidDexT1 *FluidDexT1Filterer) ParseLogWithdrawPerfectColLiquidity(log types.Log) (*FluidDexT1LogWithdrawPerfectColLiquidity, error) {
	event := new(FluidDexT1LogWithdrawPerfectColLiquidity)
	if err := _FluidDexT1.contract.UnpackLog(event, "LogWithdrawPerfectColLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FluidDexT1SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the FluidDexT1 contract.
type FluidDexT1SwapIterator struct {
	Event *FluidDexT1Swap // Event containing the contract specifics and raw log

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
func (it *FluidDexT1SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FluidDexT1Swap)
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
		it.Event = new(FluidDexT1Swap)
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
func (it *FluidDexT1SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FluidDexT1SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FluidDexT1Swap represents a Swap event raised by the FluidDexT1 contract.
type FluidDexT1Swap struct {
	Swap0to1  bool
	AmountIn  *big.Int
	AmountOut *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xdc004dbca4ef9c966218431ee5d9133d337ad018dd5b5c5493722803f75c64f7.
//
// Solidity: event Swap(bool swap0to1, uint256 amountIn, uint256 amountOut, address to)
func (_FluidDexT1 *FluidDexT1Filterer) FilterSwap(opts *bind.FilterOpts) (*FluidDexT1SwapIterator, error) {

	logs, sub, err := _FluidDexT1.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &FluidDexT1SwapIterator{contract: _FluidDexT1.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xdc004dbca4ef9c966218431ee5d9133d337ad018dd5b5c5493722803f75c64f7.
//
// Solidity: event Swap(bool swap0to1, uint256 amountIn, uint256 amountOut, address to)
func (_FluidDexT1 *FluidDexT1Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *FluidDexT1Swap) (event.Subscription, error) {

	logs, sub, err := _FluidDexT1.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FluidDexT1Swap)
				if err := _FluidDexT1.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xdc004dbca4ef9c966218431ee5d9133d337ad018dd5b5c5493722803f75c64f7.
//
// Solidity: event Swap(bool swap0to1, uint256 amountIn, uint256 amountOut, address to)
func (_FluidDexT1 *FluidDexT1Filterer) ParseSwap(log types.Log) (*FluidDexT1Swap, error) {
	event := new(FluidDexT1Swap)
	if err := _FluidDexT1.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
