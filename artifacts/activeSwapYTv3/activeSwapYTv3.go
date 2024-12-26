// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activeSwapYTv3

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

// ApproxParams is an auto generated low-level Go binding around an user-defined struct.
type ApproxParams struct {
	GuessMin      *big.Int
	GuessMax      *big.Int
	GuessOffchain *big.Int
	MaxIteration  *big.Int
	Eps           *big.Int
}

// ExitPostExpReturnParams is an auto generated low-level Go binding around an user-defined struct.
type ExitPostExpReturnParams struct {
	NetPtFromRemove *big.Int
	NetSyFromRemove *big.Int
	NetPtRedeem     *big.Int
	NetSyFromRedeem *big.Int
	TotalSyOut      *big.Int
}

// ExitPreExpReturnParams is an auto generated low-level Go binding around an user-defined struct.
type ExitPreExpReturnParams struct {
	NetPtFromRemove *big.Int
	NetSyFromRemove *big.Int
	NetPyRedeem     *big.Int
	NetSyFromRedeem *big.Int
	NetPtSwap       *big.Int
	NetYtSwap       *big.Int
	NetSyFromSwap   *big.Int
	NetSyFee        *big.Int
	TotalSyOut      *big.Int
}

// FillOrderParams is an auto generated low-level Go binding around an user-defined struct.
type FillOrderParams struct {
	Order        Order
	Signature    []byte
	MakingAmount *big.Int
}

// LimitOrderData is an auto generated low-level Go binding around an user-defined struct.
type LimitOrderData struct {
	LimitRouter   common.Address
	EpsSkipMarket *big.Int
	NormalFills   []FillOrderParams
	FlashFills    []FillOrderParams
	OptData       []byte
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Expiry        *big.Int
	Nonce         *big.Int
	OrderType     uint8
	Token         common.Address
	YT            common.Address
	Maker         common.Address
	Receiver      common.Address
	MakingAmount  *big.Int
	LnImpliedRate *big.Int
	FailSafeRate  *big.Int
	Permit        []byte
}

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	SwapType    uint8
	ExtRouter   common.Address
	ExtCalldata []byte
	NeedScale   bool
}

// TokenInput is an auto generated low-level Go binding around an user-defined struct.
type TokenInput struct {
	TokenIn     common.Address
	NetTokenIn  *big.Int
	TokenMintSy common.Address
	PendleSwap  common.Address
	SwapData    SwapData
}

// TokenOutput is an auto generated low-level Go binding around an user-defined struct.
type TokenOutput struct {
	TokenOut      common.Address
	MinTokenOut   *big.Int
	TokenRedeemSy common.Address
	PendleSwap    common.Address
	SwapData      SwapData
}

// ActiveSwapYTv3MetaData contains all meta data concerning the ActiveSwapYTv3 contract.
var ActiveSwapYTv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"}],\"name\":\"MarketExchangeRateBelowOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketProportionMustNotEqualOne\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"proportion\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maxProportion\",\"type\":\"int256\"}],\"name\":\"MarketProportionTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"rateScalar\",\"type\":\"int256\"}],\"name\":\"MarketRateScalarBelowZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"totalPt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalAsset\",\"type\":\"int256\"}],\"name\":\"MarketZeroTotalPtOrTotalAsset\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSyKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleTokenKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"}],\"name\":\"MintPyFromSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"MintPyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"MintSyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"}],\"name\":\"RedeemSyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"name\":\"SelectorToFacetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapPtAndToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapYtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapYtAndToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessTotalPtToSwap\",\"type\":\"tuple\"}],\"name\":\"swapExactPtForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactSyIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessYtOut\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactSyForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessYtOut\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactTokenForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessTotalPtFromSwap\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForPt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForSy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ActiveSwapYTv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ActiveSwapYTv3MetaData.ABI instead.
var ActiveSwapYTv3ABI = ActiveSwapYTv3MetaData.ABI

// ActiveSwapYTv3 is an auto generated Go binding around an Ethereum contract.
type ActiveSwapYTv3 struct {
	ActiveSwapYTv3Caller     // Read-only binding to the contract
	ActiveSwapYTv3Transactor // Write-only binding to the contract
	ActiveSwapYTv3Filterer   // Log filterer for contract events
}

// ActiveSwapYTv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ActiveSwapYTv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveSwapYTv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ActiveSwapYTv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveSwapYTv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActiveSwapYTv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveSwapYTv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActiveSwapYTv3Session struct {
	Contract     *ActiveSwapYTv3   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActiveSwapYTv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActiveSwapYTv3CallerSession struct {
	Contract *ActiveSwapYTv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ActiveSwapYTv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActiveSwapYTv3TransactorSession struct {
	Contract     *ActiveSwapYTv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ActiveSwapYTv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ActiveSwapYTv3Raw struct {
	Contract *ActiveSwapYTv3 // Generic contract binding to access the raw methods on
}

// ActiveSwapYTv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActiveSwapYTv3CallerRaw struct {
	Contract *ActiveSwapYTv3Caller // Generic read-only contract binding to access the raw methods on
}

// ActiveSwapYTv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActiveSwapYTv3TransactorRaw struct {
	Contract *ActiveSwapYTv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewActiveSwapYTv3 creates a new instance of ActiveSwapYTv3, bound to a specific deployed contract.
func NewActiveSwapYTv3(address common.Address, backend bind.ContractBackend) (*ActiveSwapYTv3, error) {
	contract, err := bindActiveSwapYTv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3{ActiveSwapYTv3Caller: ActiveSwapYTv3Caller{contract: contract}, ActiveSwapYTv3Transactor: ActiveSwapYTv3Transactor{contract: contract}, ActiveSwapYTv3Filterer: ActiveSwapYTv3Filterer{contract: contract}}, nil
}

// NewActiveSwapYTv3Caller creates a new read-only instance of ActiveSwapYTv3, bound to a specific deployed contract.
func NewActiveSwapYTv3Caller(address common.Address, caller bind.ContractCaller) (*ActiveSwapYTv3Caller, error) {
	contract, err := bindActiveSwapYTv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3Caller{contract: contract}, nil
}

// NewActiveSwapYTv3Transactor creates a new write-only instance of ActiveSwapYTv3, bound to a specific deployed contract.
func NewActiveSwapYTv3Transactor(address common.Address, transactor bind.ContractTransactor) (*ActiveSwapYTv3Transactor, error) {
	contract, err := bindActiveSwapYTv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3Transactor{contract: contract}, nil
}

// NewActiveSwapYTv3Filterer creates a new log filterer instance of ActiveSwapYTv3, bound to a specific deployed contract.
func NewActiveSwapYTv3Filterer(address common.Address, filterer bind.ContractFilterer) (*ActiveSwapYTv3Filterer, error) {
	contract, err := bindActiveSwapYTv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3Filterer{contract: contract}, nil
}

// bindActiveSwapYTv3 binds a generic wrapper to an already deployed contract.
func bindActiveSwapYTv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ActiveSwapYTv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActiveSwapYTv3 *ActiveSwapYTv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActiveSwapYTv3.Contract.ActiveSwapYTv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActiveSwapYTv3 *ActiveSwapYTv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.ActiveSwapYTv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActiveSwapYTv3 *ActiveSwapYTv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.ActiveSwapYTv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActiveSwapYTv3 *ActiveSwapYTv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActiveSwapYTv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.contract.Transact(opts, method, params...)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactPtForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactPtForYt", receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactPtForYt(receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactPtForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactPtForYt(receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactPtForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactSyForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut ApproxParams, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactSyForYt", receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactSyForYt(receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut ApproxParams, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactSyForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactSyForYt(receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut ApproxParams, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactSyForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactTokenForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut ApproxParams, input TokenInput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactTokenForYt", receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactTokenForYt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut ApproxParams, input TokenInput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactTokenForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactTokenForYt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut ApproxParams, input TokenInput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactTokenForYt(&_ActiveSwapYTv3.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactYtForPt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactYtForPt", receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactYtForPt(receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForPt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactYtForPt(receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap ApproxParams) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForPt(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactYtForSy(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactYtForSy", receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactYtForSy(receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForSy(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactYtForSy(receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForSy(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Transactor) SwapExactYtForToken(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.contract.Transact(opts, "swapExactYtForToken", receiver, market, exactYtIn, output, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Session) SwapExactYtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForToken(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3TransactorSession) SwapExactYtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveSwapYTv3.Contract.SwapExactYtForToken(&_ActiveSwapYTv3.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// ActiveSwapYTv3AddLiquidityDualSyAndPtIterator is returned from FilterAddLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualSyAndPt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquidityDualSyAndPtIterator struct {
	Event *ActiveSwapYTv3AddLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquidityDualSyAndPt)
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
		it.Event = new(ActiveSwapYTv3AddLiquidityDualSyAndPt)
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
func (it *ActiveSwapYTv3AddLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquidityDualSyAndPt represents a AddLiquidityDualSyAndPt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquidityDualSyAndPt struct {
	Caller    common.Address
	Market    common.Address
	Receiver  common.Address
	NetSyUsed *big.Int
	NetPtUsed *big.Int
	NetLpOut  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidityDualSyAndPt is a free log retrieval operation binding the contract event 0x9334566f6358cd68e33d423fb1c9119c6837e3a2a7a8affaaa5020ed06aec069.
//
// Solidity: event AddLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyUsed, uint256 netPtUsed, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3AddLiquidityDualSyAndPtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquidityDualSyAndPtIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0x9334566f6358cd68e33d423fb1c9119c6837e3a2a7a8affaaa5020ed06aec069.
//
// Solidity: event AddLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyUsed, uint256 netPtUsed, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquidityDualSyAndPt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
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

// ParseAddLiquidityDualSyAndPt is a log parse operation binding the contract event 0x9334566f6358cd68e33d423fb1c9119c6837e3a2a7a8affaaa5020ed06aec069.
//
// Solidity: event AddLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyUsed, uint256 netPtUsed, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquidityDualSyAndPt(log types.Log) (*ActiveSwapYTv3AddLiquidityDualSyAndPt, error) {
	event := new(ActiveSwapYTv3AddLiquidityDualSyAndPt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator is returned from FilterAddLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualTokenAndPt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator struct {
	Event *ActiveSwapYTv3AddLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquidityDualTokenAndPt)
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
		it.Event = new(ActiveSwapYTv3AddLiquidityDualTokenAndPt)
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
func (it *ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquidityDualTokenAndPt represents a AddLiquidityDualTokenAndPt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquidityDualTokenAndPt struct {
	Caller       common.Address
	Market       common.Address
	TokenIn      common.Address
	Receiver     common.Address
	NetTokenUsed *big.Int
	NetPtUsed    *big.Int
	NetLpOut     *big.Int
	NetSyInterm  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidityDualTokenAndPt is a free log retrieval operation binding the contract event 0x8969c3e485cb9f3b23622228064bc63e7350f6cf343dd86ab86169041a91eaae.
//
// Solidity: event AddLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenIn, address receiver, uint256 netTokenUsed, uint256 netPtUsed, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenIn []common.Address) (*ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquidityDualTokenAndPtIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x8969c3e485cb9f3b23622228064bc63e7350f6cf343dd86ab86169041a91eaae.
//
// Solidity: event AddLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenIn, address receiver, uint256 netTokenUsed, uint256 netPtUsed, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquidityDualTokenAndPt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
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

// ParseAddLiquidityDualTokenAndPt is a log parse operation binding the contract event 0x8969c3e485cb9f3b23622228064bc63e7350f6cf343dd86ab86169041a91eaae.
//
// Solidity: event AddLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenIn, address receiver, uint256 netTokenUsed, uint256 netPtUsed, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquidityDualTokenAndPt(log types.Log) (*ActiveSwapYTv3AddLiquidityDualTokenAndPt, error) {
	event := new(ActiveSwapYTv3AddLiquidityDualTokenAndPt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquiditySinglePtIterator is returned from FilterAddLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for AddLiquiditySinglePt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySinglePtIterator struct {
	Event *ActiveSwapYTv3AddLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquiditySinglePt)
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
		it.Event = new(ActiveSwapYTv3AddLiquiditySinglePt)
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
func (it *ActiveSwapYTv3AddLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquiditySinglePt represents a AddLiquiditySinglePt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySinglePt struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetPtIn  *big.Int
	NetLpOut *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddLiquiditySinglePt is a free log retrieval operation binding the contract event 0xc87b85efc5055ef177e0092af0d4e624fff4b1d57db748857f65e4b7e4a28a36.
//
// Solidity: event AddLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netPtIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3AddLiquiditySinglePtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquiditySinglePtIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySinglePt is a free log subscription operation binding the contract event 0xc87b85efc5055ef177e0092af0d4e624fff4b1d57db748857f65e4b7e4a28a36.
//
// Solidity: event AddLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netPtIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquiditySinglePt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
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

// ParseAddLiquiditySinglePt is a log parse operation binding the contract event 0xc87b85efc5055ef177e0092af0d4e624fff4b1d57db748857f65e4b7e4a28a36.
//
// Solidity: event AddLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netPtIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquiditySinglePt(log types.Log) (*ActiveSwapYTv3AddLiquiditySinglePt, error) {
	event := new(ActiveSwapYTv3AddLiquiditySinglePt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquiditySingleSyIterator is returned from FilterAddLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleSyIterator struct {
	Event *ActiveSwapYTv3AddLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquiditySingleSy)
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
		it.Event = new(ActiveSwapYTv3AddLiquiditySingleSy)
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
func (it *ActiveSwapYTv3AddLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquiditySingleSy represents a AddLiquiditySingleSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleSy struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetSyIn  *big.Int
	NetLpOut *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddLiquiditySingleSy is a free log retrieval operation binding the contract event 0xb51950711c9b21dc7888d41f68a19540231ffb5f0d19d8f75cbccaf90ffa7fa5.
//
// Solidity: event AddLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3AddLiquiditySingleSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquiditySingleSyIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSy is a free log subscription operation binding the contract event 0xb51950711c9b21dc7888d41f68a19540231ffb5f0d19d8f75cbccaf90ffa7fa5.
//
// Solidity: event AddLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquiditySingleSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
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

// ParseAddLiquiditySingleSy is a log parse operation binding the contract event 0xb51950711c9b21dc7888d41f68a19540231ffb5f0d19d8f75cbccaf90ffa7fa5.
//
// Solidity: event AddLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netLpOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquiditySingleSy(log types.Log) (*ActiveSwapYTv3AddLiquiditySingleSy, error) {
	event := new(ActiveSwapYTv3AddLiquiditySingleSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator is returned from FilterAddLiquiditySingleSyKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSyKeepYt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator struct {
	Event *ActiveSwapYTv3AddLiquiditySingleSyKeepYt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquiditySingleSyKeepYt)
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
		it.Event = new(ActiveSwapYTv3AddLiquiditySingleSyKeepYt)
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
func (it *ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquiditySingleSyKeepYt represents a AddLiquiditySingleSyKeepYt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleSyKeepYt struct {
	Caller      common.Address
	Market      common.Address
	Receiver    common.Address
	NetSyIn     *big.Int
	NetSyMintPy *big.Int
	NetLpOut    *big.Int
	NetYtOut    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquiditySingleSyKeepYt is a free log retrieval operation binding the contract event 0x890839d8cbce575f9d1ee3d55bc4d466623de60742c7ad665958f8a9916a54a5.
//
// Solidity: event AddLiquiditySingleSyKeepYt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netSyMintPy, uint256 netLpOut, uint256 netYtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquiditySingleSyKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquiditySingleSyKeepYtIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquiditySingleSyKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSyKeepYt is a free log subscription operation binding the contract event 0x890839d8cbce575f9d1ee3d55bc4d466623de60742c7ad665958f8a9916a54a5.
//
// Solidity: event AddLiquiditySingleSyKeepYt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netSyMintPy, uint256 netLpOut, uint256 netYtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquiditySingleSyKeepYt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquiditySingleSyKeepYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquiditySingleSyKeepYt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
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

// ParseAddLiquiditySingleSyKeepYt is a log parse operation binding the contract event 0x890839d8cbce575f9d1ee3d55bc4d466623de60742c7ad665958f8a9916a54a5.
//
// Solidity: event AddLiquiditySingleSyKeepYt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netSyMintPy, uint256 netLpOut, uint256 netYtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquiditySingleSyKeepYt(log types.Log) (*ActiveSwapYTv3AddLiquiditySingleSyKeepYt, error) {
	event := new(ActiveSwapYTv3AddLiquiditySingleSyKeepYt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquiditySingleTokenIterator is returned from FilterAddLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleTokenIterator struct {
	Event *ActiveSwapYTv3AddLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquiditySingleToken)
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
		it.Event = new(ActiveSwapYTv3AddLiquiditySingleToken)
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
func (it *ActiveSwapYTv3AddLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquiditySingleToken represents a AddLiquiditySingleToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleToken struct {
	Caller      common.Address
	Market      common.Address
	Token       common.Address
	Receiver    common.Address
	NetTokenIn  *big.Int
	NetLpOut    *big.Int
	NetSyInterm *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquiditySingleToken is a free log retrieval operation binding the contract event 0x387bf301bf673df0120e2d57e639f0e05e5e03d5336577c4cd83c1bff96e8dee.
//
// Solidity: event AddLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3AddLiquiditySingleTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquiditySingleTokenIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleToken is a free log subscription operation binding the contract event 0x387bf301bf673df0120e2d57e639f0e05e5e03d5336577c4cd83c1bff96e8dee.
//
// Solidity: event AddLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquiditySingleToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
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

// ParseAddLiquiditySingleToken is a log parse operation binding the contract event 0x387bf301bf673df0120e2d57e639f0e05e5e03d5336577c4cd83c1bff96e8dee.
//
// Solidity: event AddLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquiditySingleToken(log types.Log) (*ActiveSwapYTv3AddLiquiditySingleToken, error) {
	event := new(ActiveSwapYTv3AddLiquiditySingleToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator is returned from FilterAddLiquiditySingleTokenKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleTokenKeepYt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator struct {
	Event *ActiveSwapYTv3AddLiquiditySingleTokenKeepYt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3AddLiquiditySingleTokenKeepYt)
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
		it.Event = new(ActiveSwapYTv3AddLiquiditySingleTokenKeepYt)
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
func (it *ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3AddLiquiditySingleTokenKeepYt represents a AddLiquiditySingleTokenKeepYt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3AddLiquiditySingleTokenKeepYt struct {
	Caller      common.Address
	Market      common.Address
	Token       common.Address
	Receiver    common.Address
	NetTokenIn  *big.Int
	NetLpOut    *big.Int
	NetYtOut    *big.Int
	NetSyMintPy *big.Int
	NetSyInterm *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquiditySingleTokenKeepYt is a free log retrieval operation binding the contract event 0xa9b749795680682fbc5a34071e19282bbb23496a8cf9bbd99bf941359bbe65bf.
//
// Solidity: event AddLiquiditySingleTokenKeepYt(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netYtOut, uint256 netSyMintPy, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterAddLiquiditySingleTokenKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3AddLiquiditySingleTokenKeepYtIterator{contract: _ActiveSwapYTv3.contract, event: "AddLiquiditySingleTokenKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleTokenKeepYt is a free log subscription operation binding the contract event 0xa9b749795680682fbc5a34071e19282bbb23496a8cf9bbd99bf941359bbe65bf.
//
// Solidity: event AddLiquiditySingleTokenKeepYt(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netYtOut, uint256 netSyMintPy, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchAddLiquiditySingleTokenKeepYt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3AddLiquiditySingleTokenKeepYt, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3AddLiquiditySingleTokenKeepYt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
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

// ParseAddLiquiditySingleTokenKeepYt is a log parse operation binding the contract event 0xa9b749795680682fbc5a34071e19282bbb23496a8cf9bbd99bf941359bbe65bf.
//
// Solidity: event AddLiquiditySingleTokenKeepYt(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netYtOut, uint256 netSyMintPy, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseAddLiquiditySingleTokenKeepYt(log types.Log) (*ActiveSwapYTv3AddLiquiditySingleTokenKeepYt, error) {
	event := new(ActiveSwapYTv3AddLiquiditySingleTokenKeepYt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3ExitPostExpToSyIterator is returned from FilterExitPostExpToSy and is used to iterate over the raw logs and unpacked data for ExitPostExpToSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPostExpToSyIterator struct {
	Event *ActiveSwapYTv3ExitPostExpToSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3ExitPostExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3ExitPostExpToSy)
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
		it.Event = new(ActiveSwapYTv3ExitPostExpToSy)
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
func (it *ActiveSwapYTv3ExitPostExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3ExitPostExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3ExitPostExpToSy represents a ExitPostExpToSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPostExpToSy struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetLpIn  *big.Int
	Params   ExitPostExpReturnParams
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitPostExpToSy is a free log retrieval operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterExitPostExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3ExitPostExpToSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3ExitPostExpToSyIterator{contract: _ActiveSwapYTv3.contract, event: "ExitPostExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToSy is a free log subscription operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchExitPostExpToSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3ExitPostExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3ExitPostExpToSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
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

// ParseExitPostExpToSy is a log parse operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseExitPostExpToSy(log types.Log) (*ActiveSwapYTv3ExitPostExpToSy, error) {
	event := new(ActiveSwapYTv3ExitPostExpToSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3ExitPostExpToTokenIterator is returned from FilterExitPostExpToToken and is used to iterate over the raw logs and unpacked data for ExitPostExpToToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPostExpToTokenIterator struct {
	Event *ActiveSwapYTv3ExitPostExpToToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3ExitPostExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3ExitPostExpToToken)
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
		it.Event = new(ActiveSwapYTv3ExitPostExpToToken)
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
func (it *ActiveSwapYTv3ExitPostExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3ExitPostExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3ExitPostExpToToken represents a ExitPostExpToToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPostExpToToken struct {
	Caller        common.Address
	Market        common.Address
	Token         common.Address
	Receiver      common.Address
	NetLpIn       *big.Int
	TotalTokenOut *big.Int
	Params        ExitPostExpReturnParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitPostExpToToken is a free log retrieval operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterExitPostExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3ExitPostExpToTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3ExitPostExpToTokenIterator{contract: _ActiveSwapYTv3.contract, event: "ExitPostExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToToken is a free log subscription operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchExitPostExpToToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3ExitPostExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3ExitPostExpToToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
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

// ParseExitPostExpToToken is a log parse operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseExitPostExpToToken(log types.Log) (*ActiveSwapYTv3ExitPostExpToToken, error) {
	event := new(ActiveSwapYTv3ExitPostExpToToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3ExitPreExpToSyIterator is returned from FilterExitPreExpToSy and is used to iterate over the raw logs and unpacked data for ExitPreExpToSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPreExpToSyIterator struct {
	Event *ActiveSwapYTv3ExitPreExpToSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3ExitPreExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3ExitPreExpToSy)
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
		it.Event = new(ActiveSwapYTv3ExitPreExpToSy)
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
func (it *ActiveSwapYTv3ExitPreExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3ExitPreExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3ExitPreExpToSy represents a ExitPreExpToSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPreExpToSy struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetLpIn  *big.Int
	Params   ExitPreExpReturnParams
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitPreExpToSy is a free log retrieval operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterExitPreExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3ExitPreExpToSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3ExitPreExpToSyIterator{contract: _ActiveSwapYTv3.contract, event: "ExitPreExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToSy is a free log subscription operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchExitPreExpToSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3ExitPreExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3ExitPreExpToSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
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

// ParseExitPreExpToSy is a log parse operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseExitPreExpToSy(log types.Log) (*ActiveSwapYTv3ExitPreExpToSy, error) {
	event := new(ActiveSwapYTv3ExitPreExpToSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3ExitPreExpToTokenIterator is returned from FilterExitPreExpToToken and is used to iterate over the raw logs and unpacked data for ExitPreExpToToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPreExpToTokenIterator struct {
	Event *ActiveSwapYTv3ExitPreExpToToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3ExitPreExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3ExitPreExpToToken)
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
		it.Event = new(ActiveSwapYTv3ExitPreExpToToken)
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
func (it *ActiveSwapYTv3ExitPreExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3ExitPreExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3ExitPreExpToToken represents a ExitPreExpToToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3ExitPreExpToToken struct {
	Caller        common.Address
	Market        common.Address
	Token         common.Address
	Receiver      common.Address
	NetLpIn       *big.Int
	TotalTokenOut *big.Int
	Params        ExitPreExpReturnParams
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitPreExpToToken is a free log retrieval operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterExitPreExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3ExitPreExpToTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3ExitPreExpToTokenIterator{contract: _ActiveSwapYTv3.contract, event: "ExitPreExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToToken is a free log subscription operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchExitPreExpToToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3ExitPreExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3ExitPreExpToToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
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

// ParseExitPreExpToToken is a log parse operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseExitPreExpToToken(log types.Log) (*ActiveSwapYTv3ExitPreExpToToken, error) {
	event := new(ActiveSwapYTv3ExitPreExpToToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3MintPyFromSyIterator is returned from FilterMintPyFromSy and is used to iterate over the raw logs and unpacked data for MintPyFromSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintPyFromSyIterator struct {
	Event *ActiveSwapYTv3MintPyFromSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3MintPyFromSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3MintPyFromSy)
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
		it.Event = new(ActiveSwapYTv3MintPyFromSy)
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
func (it *ActiveSwapYTv3MintPyFromSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3MintPyFromSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3MintPyFromSy represents a MintPyFromSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintPyFromSy struct {
	Caller   common.Address
	Receiver common.Address
	YT       common.Address
	NetSyIn  *big.Int
	NetPyOut *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMintPyFromSy is a free log retrieval operation binding the contract event 0x52e05e4badd3463bad837f42fe3ba58c739d1b3081cff9bb6eb02a24034d455d.
//
// Solidity: event MintPyFromSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netSyIn, uint256 netPyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterMintPyFromSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*ActiveSwapYTv3MintPyFromSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3MintPyFromSyIterator{contract: _ActiveSwapYTv3.contract, event: "MintPyFromSy", logs: logs, sub: sub}, nil
}

// WatchMintPyFromSy is a free log subscription operation binding the contract event 0x52e05e4badd3463bad837f42fe3ba58c739d1b3081cff9bb6eb02a24034d455d.
//
// Solidity: event MintPyFromSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netSyIn, uint256 netPyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchMintPyFromSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3MintPyFromSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3MintPyFromSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
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

// ParseMintPyFromSy is a log parse operation binding the contract event 0x52e05e4badd3463bad837f42fe3ba58c739d1b3081cff9bb6eb02a24034d455d.
//
// Solidity: event MintPyFromSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netSyIn, uint256 netPyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseMintPyFromSy(log types.Log) (*ActiveSwapYTv3MintPyFromSy, error) {
	event := new(ActiveSwapYTv3MintPyFromSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3MintPyFromTokenIterator is returned from FilterMintPyFromToken and is used to iterate over the raw logs and unpacked data for MintPyFromToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintPyFromTokenIterator struct {
	Event *ActiveSwapYTv3MintPyFromToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3MintPyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3MintPyFromToken)
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
		it.Event = new(ActiveSwapYTv3MintPyFromToken)
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
func (it *ActiveSwapYTv3MintPyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3MintPyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3MintPyFromToken represents a MintPyFromToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintPyFromToken struct {
	Caller      common.Address
	TokenIn     common.Address
	YT          common.Address
	Receiver    common.Address
	NetTokenIn  *big.Int
	NetPyOut    *big.Int
	NetSyInterm *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintPyFromToken is a free log retrieval operation binding the contract event 0x3193c546cf854c6a4c63afa03b04d35e4242c2761af34a4093fc5daa88dd5308.
//
// Solidity: event MintPyFromToken(address indexed caller, address indexed tokenIn, address indexed YT, address receiver, uint256 netTokenIn, uint256 netPyOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterMintPyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, YT []common.Address) (*ActiveSwapYTv3MintPyFromTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3MintPyFromTokenIterator{contract: _ActiveSwapYTv3.contract, event: "MintPyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintPyFromToken is a free log subscription operation binding the contract event 0x3193c546cf854c6a4c63afa03b04d35e4242c2761af34a4093fc5daa88dd5308.
//
// Solidity: event MintPyFromToken(address indexed caller, address indexed tokenIn, address indexed YT, address receiver, uint256 netTokenIn, uint256 netPyOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchMintPyFromToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3MintPyFromToken, caller []common.Address, tokenIn []common.Address, YT []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3MintPyFromToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
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

// ParseMintPyFromToken is a log parse operation binding the contract event 0x3193c546cf854c6a4c63afa03b04d35e4242c2761af34a4093fc5daa88dd5308.
//
// Solidity: event MintPyFromToken(address indexed caller, address indexed tokenIn, address indexed YT, address receiver, uint256 netTokenIn, uint256 netPyOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseMintPyFromToken(log types.Log) (*ActiveSwapYTv3MintPyFromToken, error) {
	event := new(ActiveSwapYTv3MintPyFromToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3MintSyFromTokenIterator is returned from FilterMintSyFromToken and is used to iterate over the raw logs and unpacked data for MintSyFromToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintSyFromTokenIterator struct {
	Event *ActiveSwapYTv3MintSyFromToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3MintSyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3MintSyFromToken)
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
		it.Event = new(ActiveSwapYTv3MintSyFromToken)
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
func (it *ActiveSwapYTv3MintSyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3MintSyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3MintSyFromToken represents a MintSyFromToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3MintSyFromToken struct {
	Caller     common.Address
	TokenIn    common.Address
	SY         common.Address
	Receiver   common.Address
	NetTokenIn *big.Int
	NetSyOut   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintSyFromToken is a free log retrieval operation binding the contract event 0x71c7a44161eb32e4640f6c8f0586db5f1d2e03306e2c63bb2e0f7cd0a8fc690c.
//
// Solidity: event MintSyFromToken(address indexed caller, address indexed tokenIn, address indexed SY, address receiver, uint256 netTokenIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterMintSyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, SY []common.Address) (*ActiveSwapYTv3MintSyFromTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var SYRule []interface{}
	for _, SYItem := range SY {
		SYRule = append(SYRule, SYItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3MintSyFromTokenIterator{contract: _ActiveSwapYTv3.contract, event: "MintSyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintSyFromToken is a free log subscription operation binding the contract event 0x71c7a44161eb32e4640f6c8f0586db5f1d2e03306e2c63bb2e0f7cd0a8fc690c.
//
// Solidity: event MintSyFromToken(address indexed caller, address indexed tokenIn, address indexed SY, address receiver, uint256 netTokenIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchMintSyFromToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3MintSyFromToken, caller []common.Address, tokenIn []common.Address, SY []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var SYRule []interface{}
	for _, SYItem := range SY {
		SYRule = append(SYRule, SYItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3MintSyFromToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
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

// ParseMintSyFromToken is a log parse operation binding the contract event 0x71c7a44161eb32e4640f6c8f0586db5f1d2e03306e2c63bb2e0f7cd0a8fc690c.
//
// Solidity: event MintSyFromToken(address indexed caller, address indexed tokenIn, address indexed SY, address receiver, uint256 netTokenIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseMintSyFromToken(log types.Log) (*ActiveSwapYTv3MintSyFromToken, error) {
	event := new(ActiveSwapYTv3MintSyFromToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3OwnershipTransferredIterator struct {
	Event *ActiveSwapYTv3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3OwnershipTransferred)
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
		it.Event = new(ActiveSwapYTv3OwnershipTransferred)
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
func (it *ActiveSwapYTv3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3OwnershipTransferred represents a OwnershipTransferred event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActiveSwapYTv3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3OwnershipTransferredIterator{contract: _ActiveSwapYTv3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3OwnershipTransferred)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseOwnershipTransferred(log types.Log) (*ActiveSwapYTv3OwnershipTransferred, error) {
	event := new(ActiveSwapYTv3OwnershipTransferred)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RedeemPyToSyIterator is returned from FilterRedeemPyToSy and is used to iterate over the raw logs and unpacked data for RedeemPyToSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemPyToSyIterator struct {
	Event *ActiveSwapYTv3RedeemPyToSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RedeemPyToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RedeemPyToSy)
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
		it.Event = new(ActiveSwapYTv3RedeemPyToSy)
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
func (it *ActiveSwapYTv3RedeemPyToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RedeemPyToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RedeemPyToSy represents a RedeemPyToSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemPyToSy struct {
	Caller   common.Address
	Receiver common.Address
	YT       common.Address
	NetPyIn  *big.Int
	NetSyOut *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemPyToSy is a free log retrieval operation binding the contract event 0x31af33f80f4b396e3d4e42b38ecd3e022883a9bf689fd63f47afbe1d389cb6e7.
//
// Solidity: event RedeemPyToSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netPyIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRedeemPyToSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*ActiveSwapYTv3RedeemPyToSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RedeemPyToSyIterator{contract: _ActiveSwapYTv3.contract, event: "RedeemPyToSy", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToSy is a free log subscription operation binding the contract event 0x31af33f80f4b396e3d4e42b38ecd3e022883a9bf689fd63f47afbe1d389cb6e7.
//
// Solidity: event RedeemPyToSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netPyIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRedeemPyToSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RedeemPyToSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RedeemPyToSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
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

// ParseRedeemPyToSy is a log parse operation binding the contract event 0x31af33f80f4b396e3d4e42b38ecd3e022883a9bf689fd63f47afbe1d389cb6e7.
//
// Solidity: event RedeemPyToSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netPyIn, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRedeemPyToSy(log types.Log) (*ActiveSwapYTv3RedeemPyToSy, error) {
	event := new(ActiveSwapYTv3RedeemPyToSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RedeemPyToTokenIterator is returned from FilterRedeemPyToToken and is used to iterate over the raw logs and unpacked data for RedeemPyToToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemPyToTokenIterator struct {
	Event *ActiveSwapYTv3RedeemPyToToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RedeemPyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RedeemPyToToken)
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
		it.Event = new(ActiveSwapYTv3RedeemPyToToken)
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
func (it *ActiveSwapYTv3RedeemPyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RedeemPyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RedeemPyToToken represents a RedeemPyToToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemPyToToken struct {
	Caller      common.Address
	TokenOut    common.Address
	YT          common.Address
	Receiver    common.Address
	NetPyIn     *big.Int
	NetTokenOut *big.Int
	NetSyInterm *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRedeemPyToToken is a free log retrieval operation binding the contract event 0x5f2e0499a3b6a21fd5e1fac44ac47c9aa7c3afa39076d67162a4993411d496da.
//
// Solidity: event RedeemPyToToken(address indexed caller, address indexed tokenOut, address indexed YT, address receiver, uint256 netPyIn, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRedeemPyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, YT []common.Address) (*ActiveSwapYTv3RedeemPyToTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RedeemPyToTokenIterator{contract: _ActiveSwapYTv3.contract, event: "RedeemPyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToToken is a free log subscription operation binding the contract event 0x5f2e0499a3b6a21fd5e1fac44ac47c9aa7c3afa39076d67162a4993411d496da.
//
// Solidity: event RedeemPyToToken(address indexed caller, address indexed tokenOut, address indexed YT, address receiver, uint256 netPyIn, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRedeemPyToToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RedeemPyToToken, caller []common.Address, tokenOut []common.Address, YT []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var YTRule []interface{}
	for _, YTItem := range YT {
		YTRule = append(YTRule, YTItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RedeemPyToToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
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

// ParseRedeemPyToToken is a log parse operation binding the contract event 0x5f2e0499a3b6a21fd5e1fac44ac47c9aa7c3afa39076d67162a4993411d496da.
//
// Solidity: event RedeemPyToToken(address indexed caller, address indexed tokenOut, address indexed YT, address receiver, uint256 netPyIn, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRedeemPyToToken(log types.Log) (*ActiveSwapYTv3RedeemPyToToken, error) {
	event := new(ActiveSwapYTv3RedeemPyToToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RedeemSyToTokenIterator is returned from FilterRedeemSyToToken and is used to iterate over the raw logs and unpacked data for RedeemSyToToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemSyToTokenIterator struct {
	Event *ActiveSwapYTv3RedeemSyToToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RedeemSyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RedeemSyToToken)
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
		it.Event = new(ActiveSwapYTv3RedeemSyToToken)
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
func (it *ActiveSwapYTv3RedeemSyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RedeemSyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RedeemSyToToken represents a RedeemSyToToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RedeemSyToToken struct {
	Caller      common.Address
	TokenOut    common.Address
	SY          common.Address
	Receiver    common.Address
	NetSyIn     *big.Int
	NetTokenOut *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRedeemSyToToken is a free log retrieval operation binding the contract event 0xcd34b6ac7e4b72ab30845649aef2f4fd41945ae2dc08f625be69738bbd0f9aa9.
//
// Solidity: event RedeemSyToToken(address indexed caller, address indexed tokenOut, address indexed SY, address receiver, uint256 netSyIn, uint256 netTokenOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRedeemSyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, SY []common.Address) (*ActiveSwapYTv3RedeemSyToTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var SYRule []interface{}
	for _, SYItem := range SY {
		SYRule = append(SYRule, SYItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RedeemSyToTokenIterator{contract: _ActiveSwapYTv3.contract, event: "RedeemSyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemSyToToken is a free log subscription operation binding the contract event 0xcd34b6ac7e4b72ab30845649aef2f4fd41945ae2dc08f625be69738bbd0f9aa9.
//
// Solidity: event RedeemSyToToken(address indexed caller, address indexed tokenOut, address indexed SY, address receiver, uint256 netSyIn, uint256 netTokenOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRedeemSyToToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RedeemSyToToken, caller []common.Address, tokenOut []common.Address, SY []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}
	var SYRule []interface{}
	for _, SYItem := range SY {
		SYRule = append(SYRule, SYItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RedeemSyToToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
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

// ParseRedeemSyToToken is a log parse operation binding the contract event 0xcd34b6ac7e4b72ab30845649aef2f4fd41945ae2dc08f625be69738bbd0f9aa9.
//
// Solidity: event RedeemSyToToken(address indexed caller, address indexed tokenOut, address indexed SY, address receiver, uint256 netSyIn, uint256 netTokenOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRedeemSyToToken(log types.Log) (*ActiveSwapYTv3RedeemSyToToken, error) {
	event := new(ActiveSwapYTv3RedeemSyToToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator is returned from FilterRemoveLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualSyAndPt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator struct {
	Event *ActiveSwapYTv3RemoveLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RemoveLiquidityDualSyAndPt)
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
		it.Event = new(ActiveSwapYTv3RemoveLiquidityDualSyAndPt)
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
func (it *ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RemoveLiquidityDualSyAndPt represents a RemoveLiquidityDualSyAndPt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquidityDualSyAndPt struct {
	Caller        common.Address
	Market        common.Address
	Receiver      common.Address
	NetLpToRemove *big.Int
	NetPtOut      *big.Int
	NetSyOut      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityDualSyAndPt is a free log retrieval operation binding the contract event 0xd9f35a37b64d95edfd8f26adf130ce45f3e9ddf3c7ab8c1fb7224727a339a98e.
//
// Solidity: event RemoveLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRemoveLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RemoveLiquidityDualSyAndPtIterator{contract: _ActiveSwapYTv3.contract, event: "RemoveLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0xd9f35a37b64d95edfd8f26adf130ce45f3e9ddf3c7ab8c1fb7224727a339a98e.
//
// Solidity: event RemoveLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRemoveLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RemoveLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RemoveLiquidityDualSyAndPt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
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

// ParseRemoveLiquidityDualSyAndPt is a log parse operation binding the contract event 0xd9f35a37b64d95edfd8f26adf130ce45f3e9ddf3c7ab8c1fb7224727a339a98e.
//
// Solidity: event RemoveLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRemoveLiquidityDualSyAndPt(log types.Log) (*ActiveSwapYTv3RemoveLiquidityDualSyAndPt, error) {
	event := new(ActiveSwapYTv3RemoveLiquidityDualSyAndPt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator is returned from FilterRemoveLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualTokenAndPt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator struct {
	Event *ActiveSwapYTv3RemoveLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RemoveLiquidityDualTokenAndPt)
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
		it.Event = new(ActiveSwapYTv3RemoveLiquidityDualTokenAndPt)
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
func (it *ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RemoveLiquidityDualTokenAndPt represents a RemoveLiquidityDualTokenAndPt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquidityDualTokenAndPt struct {
	Caller        common.Address
	Market        common.Address
	TokenOut      common.Address
	Receiver      common.Address
	NetLpToRemove *big.Int
	NetPtOut      *big.Int
	NetTokenOut   *big.Int
	NetSyInterm   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityDualTokenAndPt is a free log retrieval operation binding the contract event 0x5349e52482e38bcf6018163f5f871bbae5e00e667aa8e7c531b74c07d5397f92.
//
// Solidity: event RemoveLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenOut, address receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRemoveLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenOut []common.Address) (*ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RemoveLiquidityDualTokenAndPtIterator{contract: _ActiveSwapYTv3.contract, event: "RemoveLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x5349e52482e38bcf6018163f5f871bbae5e00e667aa8e7c531b74c07d5397f92.
//
// Solidity: event RemoveLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenOut, address receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRemoveLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RemoveLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RemoveLiquidityDualTokenAndPt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
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

// ParseRemoveLiquidityDualTokenAndPt is a log parse operation binding the contract event 0x5349e52482e38bcf6018163f5f871bbae5e00e667aa8e7c531b74c07d5397f92.
//
// Solidity: event RemoveLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenOut, address receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRemoveLiquidityDualTokenAndPt(log types.Log) (*ActiveSwapYTv3RemoveLiquidityDualTokenAndPt, error) {
	event := new(ActiveSwapYTv3RemoveLiquidityDualTokenAndPt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RemoveLiquiditySinglePtIterator is returned from FilterRemoveLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySinglePt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySinglePtIterator struct {
	Event *ActiveSwapYTv3RemoveLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RemoveLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RemoveLiquiditySinglePt)
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
		it.Event = new(ActiveSwapYTv3RemoveLiquiditySinglePt)
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
func (it *ActiveSwapYTv3RemoveLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RemoveLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RemoveLiquiditySinglePt represents a RemoveLiquiditySinglePt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySinglePt struct {
	Caller        common.Address
	Market        common.Address
	Receiver      common.Address
	NetLpToRemove *big.Int
	NetPtOut      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquiditySinglePt is a free log retrieval operation binding the contract event 0xac97b87f5422fa3beec99bff8f336310d8ebc7d33d909b7d534cd7c72f61e871.
//
// Solidity: event RemoveLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRemoveLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3RemoveLiquiditySinglePtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RemoveLiquiditySinglePtIterator{contract: _ActiveSwapYTv3.contract, event: "RemoveLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySinglePt is a free log subscription operation binding the contract event 0xac97b87f5422fa3beec99bff8f336310d8ebc7d33d909b7d534cd7c72f61e871.
//
// Solidity: event RemoveLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRemoveLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RemoveLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RemoveLiquiditySinglePt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
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

// ParseRemoveLiquiditySinglePt is a log parse operation binding the contract event 0xac97b87f5422fa3beec99bff8f336310d8ebc7d33d909b7d534cd7c72f61e871.
//
// Solidity: event RemoveLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRemoveLiquiditySinglePt(log types.Log) (*ActiveSwapYTv3RemoveLiquiditySinglePt, error) {
	event := new(ActiveSwapYTv3RemoveLiquiditySinglePt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RemoveLiquiditySingleSyIterator is returned from FilterRemoveLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySingleSyIterator struct {
	Event *ActiveSwapYTv3RemoveLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RemoveLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RemoveLiquiditySingleSy)
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
		it.Event = new(ActiveSwapYTv3RemoveLiquiditySingleSy)
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
func (it *ActiveSwapYTv3RemoveLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RemoveLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RemoveLiquiditySingleSy represents a RemoveLiquiditySingleSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySingleSy struct {
	Caller        common.Address
	Market        common.Address
	Receiver      common.Address
	NetLpToRemove *big.Int
	NetSyOut      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquiditySingleSy is a free log retrieval operation binding the contract event 0xd31f02c44717b409d13b92ec9d2eaf1427fb4e63f85f4777f1458fb8d9387761.
//
// Solidity: event RemoveLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRemoveLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3RemoveLiquiditySingleSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RemoveLiquiditySingleSyIterator{contract: _ActiveSwapYTv3.contract, event: "RemoveLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleSy is a free log subscription operation binding the contract event 0xd31f02c44717b409d13b92ec9d2eaf1427fb4e63f85f4777f1458fb8d9387761.
//
// Solidity: event RemoveLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRemoveLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RemoveLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RemoveLiquiditySingleSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
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

// ParseRemoveLiquiditySingleSy is a log parse operation binding the contract event 0xd31f02c44717b409d13b92ec9d2eaf1427fb4e63f85f4777f1458fb8d9387761.
//
// Solidity: event RemoveLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netSyOut)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRemoveLiquiditySingleSy(log types.Log) (*ActiveSwapYTv3RemoveLiquiditySingleSy, error) {
	event := new(ActiveSwapYTv3RemoveLiquiditySingleSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3RemoveLiquiditySingleTokenIterator is returned from FilterRemoveLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySingleTokenIterator struct {
	Event *ActiveSwapYTv3RemoveLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3RemoveLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3RemoveLiquiditySingleToken)
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
		it.Event = new(ActiveSwapYTv3RemoveLiquiditySingleToken)
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
func (it *ActiveSwapYTv3RemoveLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3RemoveLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3RemoveLiquiditySingleToken represents a RemoveLiquiditySingleToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3RemoveLiquiditySingleToken struct {
	Caller        common.Address
	Market        common.Address
	Token         common.Address
	Receiver      common.Address
	NetLpToRemove *big.Int
	NetTokenOut   *big.Int
	NetSyInterm   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquiditySingleToken is a free log retrieval operation binding the contract event 0x5258a3c624debb1cc84b0f5f66c73eef048832eeebe7178e63e95a53cf28dc94.
//
// Solidity: event RemoveLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpToRemove, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterRemoveLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3RemoveLiquiditySingleTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3RemoveLiquiditySingleTokenIterator{contract: _ActiveSwapYTv3.contract, event: "RemoveLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleToken is a free log subscription operation binding the contract event 0x5258a3c624debb1cc84b0f5f66c73eef048832eeebe7178e63e95a53cf28dc94.
//
// Solidity: event RemoveLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpToRemove, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchRemoveLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3RemoveLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3RemoveLiquiditySingleToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
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

// ParseRemoveLiquiditySingleToken is a log parse operation binding the contract event 0x5258a3c624debb1cc84b0f5f66c73eef048832eeebe7178e63e95a53cf28dc94.
//
// Solidity: event RemoveLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpToRemove, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseRemoveLiquiditySingleToken(log types.Log) (*ActiveSwapYTv3RemoveLiquiditySingleToken, error) {
	event := new(ActiveSwapYTv3RemoveLiquiditySingleToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SelectorToFacetSetIterator is returned from FilterSelectorToFacetSet and is used to iterate over the raw logs and unpacked data for SelectorToFacetSet events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SelectorToFacetSetIterator struct {
	Event *ActiveSwapYTv3SelectorToFacetSet // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SelectorToFacetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SelectorToFacetSet)
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
		it.Event = new(ActiveSwapYTv3SelectorToFacetSet)
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
func (it *ActiveSwapYTv3SelectorToFacetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SelectorToFacetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SelectorToFacetSet represents a SelectorToFacetSet event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SelectorToFacetSet struct {
	Selector [4]byte
	Facet    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSelectorToFacetSet is a free log retrieval operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSelectorToFacetSet(opts *bind.FilterOpts, selector [][4]byte, facet []common.Address) (*ActiveSwapYTv3SelectorToFacetSetIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SelectorToFacetSetIterator{contract: _ActiveSwapYTv3.contract, event: "SelectorToFacetSet", logs: logs, sub: sub}, nil
}

// WatchSelectorToFacetSet is a free log subscription operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSelectorToFacetSet(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SelectorToFacetSet, selector [][4]byte, facet []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SelectorToFacetSet)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
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

// ParseSelectorToFacetSet is a log parse operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSelectorToFacetSet(log types.Log) (*ActiveSwapYTv3SelectorToFacetSet, error) {
	event := new(ActiveSwapYTv3SelectorToFacetSet)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SwapPtAndSyIterator is returned from FilterSwapPtAndSy and is used to iterate over the raw logs and unpacked data for SwapPtAndSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndSyIterator struct {
	Event *ActiveSwapYTv3SwapPtAndSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SwapPtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SwapPtAndSy)
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
		it.Event = new(ActiveSwapYTv3SwapPtAndSy)
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
func (it *ActiveSwapYTv3SwapPtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SwapPtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SwapPtAndSy represents a SwapPtAndSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndSy struct {
	Caller         common.Address
	Market         common.Address
	Receiver       common.Address
	NetPtToAccount *big.Int
	NetSyToAccount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwapPtAndSy is a free log retrieval operation binding the contract event 0x3f5e2944826baeaed8eb77f0f74e6088a154a0fc1317f062fd984585607b4739.
//
// Solidity: event SwapPtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSwapPtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3SwapPtAndSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SwapPtAndSyIterator{contract: _ActiveSwapYTv3.contract, event: "SwapPtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndSy is a free log subscription operation binding the contract event 0x3f5e2944826baeaed8eb77f0f74e6088a154a0fc1317f062fd984585607b4739.
//
// Solidity: event SwapPtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSwapPtAndSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SwapPtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SwapPtAndSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
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

// ParseSwapPtAndSy is a log parse operation binding the contract event 0x3f5e2944826baeaed8eb77f0f74e6088a154a0fc1317f062fd984585607b4739.
//
// Solidity: event SwapPtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSwapPtAndSy(log types.Log) (*ActiveSwapYTv3SwapPtAndSy, error) {
	event := new(ActiveSwapYTv3SwapPtAndSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SwapPtAndTokenIterator is returned from FilterSwapPtAndToken and is used to iterate over the raw logs and unpacked data for SwapPtAndToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndTokenIterator struct {
	Event *ActiveSwapYTv3SwapPtAndToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SwapPtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SwapPtAndToken)
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
		it.Event = new(ActiveSwapYTv3SwapPtAndToken)
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
func (it *ActiveSwapYTv3SwapPtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SwapPtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SwapPtAndToken represents a SwapPtAndToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndToken struct {
	Caller            common.Address
	Market            common.Address
	Token             common.Address
	Receiver          common.Address
	NetPtToAccount    *big.Int
	NetTokenToAccount *big.Int
	NetSyInterm       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapPtAndToken is a free log retrieval operation binding the contract event 0xd3c1d9b397236779b29ee5b5b150c1110fc8221b6b6ec0be49c9f4860ceb2036.
//
// Solidity: event SwapPtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netPtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSwapPtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3SwapPtAndTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SwapPtAndTokenIterator{contract: _ActiveSwapYTv3.contract, event: "SwapPtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndToken is a free log subscription operation binding the contract event 0xd3c1d9b397236779b29ee5b5b150c1110fc8221b6b6ec0be49c9f4860ceb2036.
//
// Solidity: event SwapPtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netPtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSwapPtAndToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SwapPtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SwapPtAndToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
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

// ParseSwapPtAndToken is a log parse operation binding the contract event 0xd3c1d9b397236779b29ee5b5b150c1110fc8221b6b6ec0be49c9f4860ceb2036.
//
// Solidity: event SwapPtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netPtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSwapPtAndToken(log types.Log) (*ActiveSwapYTv3SwapPtAndToken, error) {
	event := new(ActiveSwapYTv3SwapPtAndToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SwapPtAndYtIterator is returned from FilterSwapPtAndYt and is used to iterate over the raw logs and unpacked data for SwapPtAndYt events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndYtIterator struct {
	Event *ActiveSwapYTv3SwapPtAndYt // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SwapPtAndYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SwapPtAndYt)
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
		it.Event = new(ActiveSwapYTv3SwapPtAndYt)
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
func (it *ActiveSwapYTv3SwapPtAndYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SwapPtAndYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SwapPtAndYt represents a SwapPtAndYt event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapPtAndYt struct {
	Caller         common.Address
	Market         common.Address
	Receiver       common.Address
	NetPtToAccount *big.Int
	NetYtToAccount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwapPtAndYt is a free log retrieval operation binding the contract event 0xa4519acd6c251a3f533922c5aaf3fbae71546aad6c01f8241e23143519a67ac8.
//
// Solidity: event SwapPtAndYt(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netYtToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSwapPtAndYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3SwapPtAndYtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SwapPtAndYtIterator{contract: _ActiveSwapYTv3.contract, event: "SwapPtAndYt", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndYt is a free log subscription operation binding the contract event 0xa4519acd6c251a3f533922c5aaf3fbae71546aad6c01f8241e23143519a67ac8.
//
// Solidity: event SwapPtAndYt(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netYtToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSwapPtAndYt(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SwapPtAndYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SwapPtAndYt)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
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

// ParseSwapPtAndYt is a log parse operation binding the contract event 0xa4519acd6c251a3f533922c5aaf3fbae71546aad6c01f8241e23143519a67ac8.
//
// Solidity: event SwapPtAndYt(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netYtToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSwapPtAndYt(log types.Log) (*ActiveSwapYTv3SwapPtAndYt, error) {
	event := new(ActiveSwapYTv3SwapPtAndYt)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SwapYtAndSyIterator is returned from FilterSwapYtAndSy and is used to iterate over the raw logs and unpacked data for SwapYtAndSy events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapYtAndSyIterator struct {
	Event *ActiveSwapYTv3SwapYtAndSy // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SwapYtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SwapYtAndSy)
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
		it.Event = new(ActiveSwapYTv3SwapYtAndSy)
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
func (it *ActiveSwapYTv3SwapYtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SwapYtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SwapYtAndSy represents a SwapYtAndSy event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapYtAndSy struct {
	Caller         common.Address
	Market         common.Address
	Receiver       common.Address
	NetYtToAccount *big.Int
	NetSyToAccount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwapYtAndSy is a free log retrieval operation binding the contract event 0x05499aba408f669fb848399c146fad5bd604d50b15566bdc19e81c40922fab8d.
//
// Solidity: event SwapYtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netYtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSwapYtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveSwapYTv3SwapYtAndSyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SwapYtAndSyIterator{contract: _ActiveSwapYTv3.contract, event: "SwapYtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndSy is a free log subscription operation binding the contract event 0x05499aba408f669fb848399c146fad5bd604d50b15566bdc19e81c40922fab8d.
//
// Solidity: event SwapYtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netYtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSwapYtAndSy(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SwapYtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SwapYtAndSy)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
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

// ParseSwapYtAndSy is a log parse operation binding the contract event 0x05499aba408f669fb848399c146fad5bd604d50b15566bdc19e81c40922fab8d.
//
// Solidity: event SwapYtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netYtToAccount, int256 netSyToAccount)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSwapYtAndSy(log types.Log) (*ActiveSwapYTv3SwapYtAndSy, error) {
	event := new(ActiveSwapYTv3SwapYtAndSy)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveSwapYTv3SwapYtAndTokenIterator is returned from FilterSwapYtAndToken and is used to iterate over the raw logs and unpacked data for SwapYtAndToken events raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapYtAndTokenIterator struct {
	Event *ActiveSwapYTv3SwapYtAndToken // Event containing the contract specifics and raw log

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
func (it *ActiveSwapYTv3SwapYtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveSwapYTv3SwapYtAndToken)
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
		it.Event = new(ActiveSwapYTv3SwapYtAndToken)
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
func (it *ActiveSwapYTv3SwapYtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveSwapYTv3SwapYtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveSwapYTv3SwapYtAndToken represents a SwapYtAndToken event raised by the ActiveSwapYTv3 contract.
type ActiveSwapYTv3SwapYtAndToken struct {
	Caller            common.Address
	Market            common.Address
	Token             common.Address
	Receiver          common.Address
	NetYtToAccount    *big.Int
	NetTokenToAccount *big.Int
	NetSyInterm       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapYtAndToken is a free log retrieval operation binding the contract event 0xa3a2846538c60e47775faa60c6ae79b67dee6d97bb70e386ebbaf4c3a38e8b81.
//
// Solidity: event SwapYtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netYtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) FilterSwapYtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveSwapYTv3SwapYtAndTokenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.FilterLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveSwapYTv3SwapYtAndTokenIterator{contract: _ActiveSwapYTv3.contract, event: "SwapYtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndToken is a free log subscription operation binding the contract event 0xa3a2846538c60e47775faa60c6ae79b67dee6d97bb70e386ebbaf4c3a38e8b81.
//
// Solidity: event SwapYtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netYtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) WatchSwapYtAndToken(opts *bind.WatchOpts, sink chan<- *ActiveSwapYTv3SwapYtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ActiveSwapYTv3.contract.WatchLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveSwapYTv3SwapYtAndToken)
				if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
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

// ParseSwapYtAndToken is a log parse operation binding the contract event 0xa3a2846538c60e47775faa60c6ae79b67dee6d97bb70e386ebbaf4c3a38e8b81.
//
// Solidity: event SwapYtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netYtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveSwapYTv3 *ActiveSwapYTv3Filterer) ParseSwapYtAndToken(log types.Log) (*ActiveSwapYTv3SwapYtAndToken, error) {
	event := new(ActiveSwapYTv3SwapYtAndToken)
	if err := _ActiveSwapYTv3.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
