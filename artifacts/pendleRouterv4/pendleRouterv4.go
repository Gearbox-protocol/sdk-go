// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pendleRouterv4

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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Order        Struct1
	Signature    []byte
	MakingAmount *big.Int
}

// Struct3 is an auto generated low-level Go binding around an user-defined struct.
type Struct3 struct {
	LimitRouter   common.Address
	EpsSkipMarket *big.Int
	NormalFills   []Struct2
	FlashFills    []Struct2
	OptData       []byte
}

// Struct5 is an auto generated low-level Go binding around an user-defined struct.
type Struct5 struct {
	TokenIn     common.Address
	NetTokenIn  *big.Int
	TokenMintSy common.Address
	PendleSwap  common.Address
	SwapData    Struct4
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	GuessMin      *big.Int
	GuessMax      *big.Int
	GuessOffchain *big.Int
	MaxIteration  *big.Int
	Eps           *big.Int
}

// Struct6 is an auto generated low-level Go binding around an user-defined struct.
type Struct6 struct {
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

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
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

// Struct4 is an auto generated low-level Go binding around an user-defined struct.
type Struct4 struct {
	SwapType    uint8
	ExtRouter   common.Address
	ExtCalldata []byte
	NeedScale   bool
}

// PendleRouterv4MetaData contains all meta data concerning the PendleRouterv4 contract.
var PendleRouterv4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"exchangeRate\",\"type\":\"int256\"}],\"name\":\"MarketExchangeRateBelowOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketProportionMustNotEqualOne\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"proportion\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"maxProportion\",\"type\":\"int256\"}],\"name\":\"MarketProportionTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"rateScalar\",\"type\":\"int256\"}],\"name\":\"MarketRateScalarBelowZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"totalPt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"totalAsset\",\"type\":\"int256\"}],\"name\":\"MarketZeroTotalPtOrTotalAsset\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSyKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleTokenKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"}],\"name\":\"MintPyFromSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"MintPyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"MintSyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"}],\"name\":\"RedeemSyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"name\":\"SelectorToFacetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapPtAndToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapYtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapYtAndToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessTotalPtToSwap\",\"type\":\"tuple\"}],\"name\":\"swapExactPtForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactSyIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessYtOut\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactSyForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessYtOut\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactTokenForYt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minYtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessYtOut\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactTokenForPt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPtOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"guessMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guessOffchain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIteration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eps\",\"type\":\"uint256\"}],\"internalType\":\"structApproxParams\",\"name\":\"guessTotalPtFromSwap\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForPt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForSy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactTokenForPt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactPtForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactYtIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"swapExactYtForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PendleRouterv4ABI is the input ABI used to generate the binding from.
// Deprecated: Use PendleRouterv4MetaData.ABI instead.
var PendleRouterv4ABI = PendleRouterv4MetaData.ABI

// PendleRouterv4 is an auto generated Go binding around an Ethereum contract.
type PendleRouterv4 struct {
	PendleRouterv4Caller     // Read-only binding to the contract
	PendleRouterv4Transactor // Write-only binding to the contract
	PendleRouterv4Filterer   // Log filterer for contract events
}

// PendleRouterv4Caller is an auto generated read-only Go binding around an Ethereum contract.
type PendleRouterv4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterv4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleRouterv4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterv4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleRouterv4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleRouterv4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleRouterv4Session struct {
	Contract     *PendleRouterv4   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleRouterv4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleRouterv4CallerSession struct {
	Contract *PendleRouterv4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PendleRouterv4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleRouterv4TransactorSession struct {
	Contract     *PendleRouterv4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PendleRouterv4Raw is an auto generated low-level Go binding around an Ethereum contract.
type PendleRouterv4Raw struct {
	Contract *PendleRouterv4 // Generic contract binding to access the raw methods on
}

// PendleRouterv4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleRouterv4CallerRaw struct {
	Contract *PendleRouterv4Caller // Generic read-only contract binding to access the raw methods on
}

// PendleRouterv4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleRouterv4TransactorRaw struct {
	Contract *PendleRouterv4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleRouterv4 creates a new instance of PendleRouterv4, bound to a specific deployed contract.
func NewPendleRouterv4(address common.Address, backend bind.ContractBackend) (*PendleRouterv4, error) {
	contract, err := bindPendleRouterv4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4{PendleRouterv4Caller: PendleRouterv4Caller{contract: contract}, PendleRouterv4Transactor: PendleRouterv4Transactor{contract: contract}, PendleRouterv4Filterer: PendleRouterv4Filterer{contract: contract}}, nil
}

// NewPendleRouterv4Caller creates a new read-only instance of PendleRouterv4, bound to a specific deployed contract.
func NewPendleRouterv4Caller(address common.Address, caller bind.ContractCaller) (*PendleRouterv4Caller, error) {
	contract, err := bindPendleRouterv4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4Caller{contract: contract}, nil
}

// NewPendleRouterv4Transactor creates a new write-only instance of PendleRouterv4, bound to a specific deployed contract.
func NewPendleRouterv4Transactor(address common.Address, transactor bind.ContractTransactor) (*PendleRouterv4Transactor, error) {
	contract, err := bindPendleRouterv4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4Transactor{contract: contract}, nil
}

// NewPendleRouterv4Filterer creates a new log filterer instance of PendleRouterv4, bound to a specific deployed contract.
func NewPendleRouterv4Filterer(address common.Address, filterer bind.ContractFilterer) (*PendleRouterv4Filterer, error) {
	contract, err := bindPendleRouterv4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4Filterer{contract: contract}, nil
}

// bindPendleRouterv4 binds a generic wrapper to an already deployed contract.
func bindPendleRouterv4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PendleRouterv4MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouterv4 *PendleRouterv4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouterv4.Contract.PendleRouterv4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouterv4 *PendleRouterv4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.PendleRouterv4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouterv4 *PendleRouterv4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.PendleRouterv4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleRouterv4 *PendleRouterv4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleRouterv4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleRouterv4 *PendleRouterv4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleRouterv4 *PendleRouterv4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.contract.Transact(opts, method, params...)
}

// SwapExactPtForToken is a paid mutator transaction binding the contract method 0x594a88cc.
//
// Solidity: function swapExactPtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactPtForToken(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactPtForToken", receiver, market, exactYtIn, output, limit)
}

// SwapExactPtForToken is a paid mutator transaction binding the contract method 0x594a88cc.
//
// Solidity: function swapExactPtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactPtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactPtForToken(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactPtForToken is a paid mutator transaction binding the contract method 0x594a88cc.
//
// Solidity: function swapExactPtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactPtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactPtForToken(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactPtForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactPtForYt", receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactPtForYt(receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactPtForYt(&_PendleRouterv4.TransactOpts, receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactPtForYt is a paid mutator transaction binding the contract method 0xc861a898.
//
// Solidity: function swapExactPtForYt(address receiver, address market, uint256 exactPtIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtToSwap) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactPtForYt(receiver common.Address, market common.Address, exactPtIn *big.Int, minYtOut *big.Int, guessTotalPtToSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactPtForYt(&_PendleRouterv4.TransactOpts, receiver, market, exactPtIn, minYtOut, guessTotalPtToSwap)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactSyForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut Struct0, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactSyForYt", receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactSyForYt(receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut Struct0, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactSyForYt(&_PendleRouterv4.TransactOpts, receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactSyForYt is a paid mutator transaction binding the contract method 0x7b8b4b95.
//
// Solidity: function swapExactSyForYt(address receiver, address market, uint256 exactSyIn, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netYtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactSyForYt(receiver common.Address, market common.Address, exactSyIn *big.Int, minYtOut *big.Int, guessYtOut Struct0, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactSyForYt(&_PendleRouterv4.TransactOpts, receiver, market, exactSyIn, minYtOut, guessYtOut, limit)
}

// SwapExactTokenForPt is a paid mutator transaction binding the contract method 0xc81f847a.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactTokenForPt(opts *bind.TransactOpts, receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactTokenForPt", receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForPt is a paid mutator transaction binding the contract method 0xc81f847a.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactTokenForPt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForPt(&_PendleRouterv4.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForPt is a paid mutator transaction binding the contract method 0xc81f847a.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactTokenForPt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForPt(&_PendleRouterv4.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForPt0 is a paid mutator transaction binding the contract method 0x04db06ae.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactTokenForPt0(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactTokenForPt0", receiver, market, exactYtIn, output, limit)
}

// SwapExactTokenForPt0 is a paid mutator transaction binding the contract method 0x04db06ae.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactTokenForPt0(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForPt0(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactTokenForPt0 is a paid mutator transaction binding the contract method 0x04db06ae.
//
// Solidity: function swapExactTokenForPt(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactTokenForPt0(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForPt0(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactTokenForYt(opts *bind.TransactOpts, receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactTokenForYt", receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactTokenForYt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForYt(&_PendleRouterv4.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactTokenForYt is a paid mutator transaction binding the contract method 0xed48907e.
//
// Solidity: function swapExactTokenForYt(address receiver, address market, uint256 minYtOut, (uint256,uint256,uint256,uint256,uint256) guessYtOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) payable returns(uint256 netYtOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactTokenForYt(receiver common.Address, market common.Address, minYtOut *big.Int, guessYtOut Struct0, input Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactTokenForYt(&_PendleRouterv4.TransactOpts, receiver, market, minYtOut, guessYtOut, input, limit)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactYtForPt(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactYtForPt", receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactYtForPt(receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForPt(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForPt is a paid mutator transaction binding the contract method 0x448b9b95.
//
// Solidity: function swapExactYtForPt(address receiver, address market, uint256 exactYtIn, uint256 minPtOut, (uint256,uint256,uint256,uint256,uint256) guessTotalPtFromSwap) returns(uint256 netPtOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactYtForPt(receiver common.Address, market common.Address, exactYtIn *big.Int, minPtOut *big.Int, guessTotalPtFromSwap Struct0) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForPt(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, minPtOut, guessTotalPtFromSwap)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactYtForSy(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactYtForSy", receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactYtForSy(receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForSy(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForSy is a paid mutator transaction binding the contract method 0x80c4d566.
//
// Solidity: function swapExactYtForSy(address receiver, address market, uint256 exactYtIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netSyOut, uint256 netSyFee)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactYtForSy(receiver common.Address, market common.Address, exactYtIn *big.Int, minSyOut *big.Int, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForSy(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, minSyOut, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Transactor) SwapExactYtForToken(opts *bind.TransactOpts, receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.contract.Transact(opts, "swapExactYtForToken", receiver, market, exactYtIn, output, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Session) SwapExactYtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForToken(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// SwapExactYtForToken is a paid mutator transaction binding the contract method 0x05eb5327.
//
// Solidity: function swapExactYtForToken(address receiver, address market, uint256 exactYtIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 netTokenOut, uint256 netSyFee, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4TransactorSession) SwapExactYtForToken(receiver common.Address, market common.Address, exactYtIn *big.Int, output Struct5, limit Struct3) (*types.Transaction, error) {
	return _PendleRouterv4.Contract.SwapExactYtForToken(&_PendleRouterv4.TransactOpts, receiver, market, exactYtIn, output, limit)
}

// PendleRouterv4AddLiquidityDualSyAndPtIterator is returned from FilterAddLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualSyAndPt events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquidityDualSyAndPtIterator struct {
	Event *PendleRouterv4AddLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquidityDualSyAndPt)
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
		it.Event = new(PendleRouterv4AddLiquidityDualSyAndPt)
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
func (it *PendleRouterv4AddLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquidityDualSyAndPt represents a AddLiquidityDualSyAndPt event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquidityDualSyAndPt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4AddLiquidityDualSyAndPtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquidityDualSyAndPtIterator{contract: _PendleRouterv4.contract, event: "AddLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0x9334566f6358cd68e33d423fb1c9119c6837e3a2a7a8affaaa5020ed06aec069.
//
// Solidity: event AddLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyUsed, uint256 netPtUsed, uint256 netLpOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquidityDualSyAndPt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquidityDualSyAndPt(log types.Log) (*PendleRouterv4AddLiquidityDualSyAndPt, error) {
	event := new(PendleRouterv4AddLiquidityDualSyAndPt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquidityDualTokenAndPtIterator is returned from FilterAddLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualTokenAndPt events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquidityDualTokenAndPtIterator struct {
	Event *PendleRouterv4AddLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquidityDualTokenAndPt)
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
		it.Event = new(PendleRouterv4AddLiquidityDualTokenAndPt)
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
func (it *PendleRouterv4AddLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquidityDualTokenAndPt represents a AddLiquidityDualTokenAndPt event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquidityDualTokenAndPt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenIn []common.Address) (*PendleRouterv4AddLiquidityDualTokenAndPtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquidityDualTokenAndPtIterator{contract: _PendleRouterv4.contract, event: "AddLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x8969c3e485cb9f3b23622228064bc63e7350f6cf343dd86ab86169041a91eaae.
//
// Solidity: event AddLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenIn, address receiver, uint256 netTokenUsed, uint256 netPtUsed, uint256 netLpOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenIn []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquidityDualTokenAndPt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquidityDualTokenAndPt(log types.Log) (*PendleRouterv4AddLiquidityDualTokenAndPt, error) {
	event := new(PendleRouterv4AddLiquidityDualTokenAndPt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquiditySinglePtIterator is returned from FilterAddLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for AddLiquiditySinglePt events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySinglePtIterator struct {
	Event *PendleRouterv4AddLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquiditySinglePt)
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
		it.Event = new(PendleRouterv4AddLiquiditySinglePt)
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
func (it *PendleRouterv4AddLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquiditySinglePt represents a AddLiquiditySinglePt event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySinglePt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4AddLiquiditySinglePtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquiditySinglePtIterator{contract: _PendleRouterv4.contract, event: "AddLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySinglePt is a free log subscription operation binding the contract event 0xc87b85efc5055ef177e0092af0d4e624fff4b1d57db748857f65e4b7e4a28a36.
//
// Solidity: event AddLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netPtIn, uint256 netLpOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquiditySinglePt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquiditySinglePt(log types.Log) (*PendleRouterv4AddLiquiditySinglePt, error) {
	event := new(PendleRouterv4AddLiquiditySinglePt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquiditySingleSyIterator is returned from FilterAddLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSy events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleSyIterator struct {
	Event *PendleRouterv4AddLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquiditySingleSy)
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
		it.Event = new(PendleRouterv4AddLiquiditySingleSy)
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
func (it *PendleRouterv4AddLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquiditySingleSy represents a AddLiquiditySingleSy event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4AddLiquiditySingleSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquiditySingleSyIterator{contract: _PendleRouterv4.contract, event: "AddLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSy is a free log subscription operation binding the contract event 0xb51950711c9b21dc7888d41f68a19540231ffb5f0d19d8f75cbccaf90ffa7fa5.
//
// Solidity: event AddLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netLpOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquiditySingleSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquiditySingleSy(log types.Log) (*PendleRouterv4AddLiquiditySingleSy, error) {
	event := new(PendleRouterv4AddLiquiditySingleSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquiditySingleSyKeepYtIterator is returned from FilterAddLiquiditySingleSyKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSyKeepYt events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleSyKeepYtIterator struct {
	Event *PendleRouterv4AddLiquiditySingleSyKeepYt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquiditySingleSyKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquiditySingleSyKeepYt)
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
		it.Event = new(PendleRouterv4AddLiquiditySingleSyKeepYt)
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
func (it *PendleRouterv4AddLiquiditySingleSyKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquiditySingleSyKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquiditySingleSyKeepYt represents a AddLiquiditySingleSyKeepYt event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleSyKeepYt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquiditySingleSyKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4AddLiquiditySingleSyKeepYtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquiditySingleSyKeepYtIterator{contract: _PendleRouterv4.contract, event: "AddLiquiditySingleSyKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSyKeepYt is a free log subscription operation binding the contract event 0x890839d8cbce575f9d1ee3d55bc4d466623de60742c7ad665958f8a9916a54a5.
//
// Solidity: event AddLiquiditySingleSyKeepYt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netSyMintPy, uint256 netLpOut, uint256 netYtOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquiditySingleSyKeepYt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquiditySingleSyKeepYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquiditySingleSyKeepYt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquiditySingleSyKeepYt(log types.Log) (*PendleRouterv4AddLiquiditySingleSyKeepYt, error) {
	event := new(PendleRouterv4AddLiquiditySingleSyKeepYt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquiditySingleTokenIterator is returned from FilterAddLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleToken events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleTokenIterator struct {
	Event *PendleRouterv4AddLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquiditySingleToken)
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
		it.Event = new(PendleRouterv4AddLiquiditySingleToken)
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
func (it *PendleRouterv4AddLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquiditySingleToken represents a AddLiquiditySingleToken event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4AddLiquiditySingleTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquiditySingleTokenIterator{contract: _PendleRouterv4.contract, event: "AddLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleToken is a free log subscription operation binding the contract event 0x387bf301bf673df0120e2d57e639f0e05e5e03d5336577c4cd83c1bff96e8dee.
//
// Solidity: event AddLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquiditySingleToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquiditySingleToken(log types.Log) (*PendleRouterv4AddLiquiditySingleToken, error) {
	event := new(PendleRouterv4AddLiquiditySingleToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4AddLiquiditySingleTokenKeepYtIterator is returned from FilterAddLiquiditySingleTokenKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleTokenKeepYt events raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleTokenKeepYtIterator struct {
	Event *PendleRouterv4AddLiquiditySingleTokenKeepYt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4AddLiquiditySingleTokenKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4AddLiquiditySingleTokenKeepYt)
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
		it.Event = new(PendleRouterv4AddLiquiditySingleTokenKeepYt)
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
func (it *PendleRouterv4AddLiquiditySingleTokenKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4AddLiquiditySingleTokenKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4AddLiquiditySingleTokenKeepYt represents a AddLiquiditySingleTokenKeepYt event raised by the PendleRouterv4 contract.
type PendleRouterv4AddLiquiditySingleTokenKeepYt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterAddLiquiditySingleTokenKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4AddLiquiditySingleTokenKeepYtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4AddLiquiditySingleTokenKeepYtIterator{contract: _PendleRouterv4.contract, event: "AddLiquiditySingleTokenKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleTokenKeepYt is a free log subscription operation binding the contract event 0xa9b749795680682fbc5a34071e19282bbb23496a8cf9bbd99bf941359bbe65bf.
//
// Solidity: event AddLiquiditySingleTokenKeepYt(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netYtOut, uint256 netSyMintPy, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchAddLiquiditySingleTokenKeepYt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4AddLiquiditySingleTokenKeepYt, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4AddLiquiditySingleTokenKeepYt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseAddLiquiditySingleTokenKeepYt(log types.Log) (*PendleRouterv4AddLiquiditySingleTokenKeepYt, error) {
	event := new(PendleRouterv4AddLiquiditySingleTokenKeepYt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4ExitPostExpToSyIterator is returned from FilterExitPostExpToSy and is used to iterate over the raw logs and unpacked data for ExitPostExpToSy events raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPostExpToSyIterator struct {
	Event *PendleRouterv4ExitPostExpToSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4ExitPostExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4ExitPostExpToSy)
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
		it.Event = new(PendleRouterv4ExitPostExpToSy)
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
func (it *PendleRouterv4ExitPostExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4ExitPostExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4ExitPostExpToSy represents a ExitPostExpToSy event raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPostExpToSy struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetLpIn  *big.Int
	Params   Struct0
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitPostExpToSy is a free log retrieval operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterExitPostExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4ExitPostExpToSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4ExitPostExpToSyIterator{contract: _PendleRouterv4.contract, event: "ExitPostExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToSy is a free log subscription operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchExitPostExpToSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4ExitPostExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4ExitPostExpToSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseExitPostExpToSy(log types.Log) (*PendleRouterv4ExitPostExpToSy, error) {
	event := new(PendleRouterv4ExitPostExpToSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4ExitPostExpToTokenIterator is returned from FilterExitPostExpToToken and is used to iterate over the raw logs and unpacked data for ExitPostExpToToken events raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPostExpToTokenIterator struct {
	Event *PendleRouterv4ExitPostExpToToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4ExitPostExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4ExitPostExpToToken)
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
		it.Event = new(PendleRouterv4ExitPostExpToToken)
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
func (it *PendleRouterv4ExitPostExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4ExitPostExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4ExitPostExpToToken represents a ExitPostExpToToken event raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPostExpToToken struct {
	Caller        common.Address
	Market        common.Address
	Token         common.Address
	Receiver      common.Address
	NetLpIn       *big.Int
	TotalTokenOut *big.Int
	Params        Struct0
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitPostExpToToken is a free log retrieval operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterExitPostExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4ExitPostExpToTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4ExitPostExpToTokenIterator{contract: _PendleRouterv4.contract, event: "ExitPostExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToToken is a free log subscription operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchExitPostExpToToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4ExitPostExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4ExitPostExpToToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseExitPostExpToToken(log types.Log) (*PendleRouterv4ExitPostExpToToken, error) {
	event := new(PendleRouterv4ExitPostExpToToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4ExitPreExpToSyIterator is returned from FilterExitPreExpToSy and is used to iterate over the raw logs and unpacked data for ExitPreExpToSy events raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPreExpToSyIterator struct {
	Event *PendleRouterv4ExitPreExpToSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4ExitPreExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4ExitPreExpToSy)
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
		it.Event = new(PendleRouterv4ExitPreExpToSy)
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
func (it *PendleRouterv4ExitPreExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4ExitPreExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4ExitPreExpToSy represents a ExitPreExpToSy event raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPreExpToSy struct {
	Caller   common.Address
	Market   common.Address
	Receiver common.Address
	NetLpIn  *big.Int
	Params   Struct6
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitPreExpToSy is a free log retrieval operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterExitPreExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4ExitPreExpToSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4ExitPreExpToSyIterator{contract: _PendleRouterv4.contract, event: "ExitPreExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToSy is a free log subscription operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchExitPreExpToSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4ExitPreExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4ExitPreExpToSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseExitPreExpToSy(log types.Log) (*PendleRouterv4ExitPreExpToSy, error) {
	event := new(PendleRouterv4ExitPreExpToSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4ExitPreExpToTokenIterator is returned from FilterExitPreExpToToken and is used to iterate over the raw logs and unpacked data for ExitPreExpToToken events raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPreExpToTokenIterator struct {
	Event *PendleRouterv4ExitPreExpToToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4ExitPreExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4ExitPreExpToToken)
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
		it.Event = new(PendleRouterv4ExitPreExpToToken)
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
func (it *PendleRouterv4ExitPreExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4ExitPreExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4ExitPreExpToToken represents a ExitPreExpToToken event raised by the PendleRouterv4 contract.
type PendleRouterv4ExitPreExpToToken struct {
	Caller        common.Address
	Market        common.Address
	Token         common.Address
	Receiver      common.Address
	NetLpIn       *big.Int
	TotalTokenOut *big.Int
	Params        Struct6
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExitPreExpToToken is a free log retrieval operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterExitPreExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4ExitPreExpToTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4ExitPreExpToTokenIterator{contract: _PendleRouterv4.contract, event: "ExitPreExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToToken is a free log subscription operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchExitPreExpToToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4ExitPreExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4ExitPreExpToToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseExitPreExpToToken(log types.Log) (*PendleRouterv4ExitPreExpToToken, error) {
	event := new(PendleRouterv4ExitPreExpToToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4MintPyFromSyIterator is returned from FilterMintPyFromSy and is used to iterate over the raw logs and unpacked data for MintPyFromSy events raised by the PendleRouterv4 contract.
type PendleRouterv4MintPyFromSyIterator struct {
	Event *PendleRouterv4MintPyFromSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4MintPyFromSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4MintPyFromSy)
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
		it.Event = new(PendleRouterv4MintPyFromSy)
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
func (it *PendleRouterv4MintPyFromSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4MintPyFromSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4MintPyFromSy represents a MintPyFromSy event raised by the PendleRouterv4 contract.
type PendleRouterv4MintPyFromSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterMintPyFromSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*PendleRouterv4MintPyFromSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4MintPyFromSyIterator{contract: _PendleRouterv4.contract, event: "MintPyFromSy", logs: logs, sub: sub}, nil
}

// WatchMintPyFromSy is a free log subscription operation binding the contract event 0x52e05e4badd3463bad837f42fe3ba58c739d1b3081cff9bb6eb02a24034d455d.
//
// Solidity: event MintPyFromSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netSyIn, uint256 netPyOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchMintPyFromSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4MintPyFromSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4MintPyFromSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseMintPyFromSy(log types.Log) (*PendleRouterv4MintPyFromSy, error) {
	event := new(PendleRouterv4MintPyFromSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4MintPyFromTokenIterator is returned from FilterMintPyFromToken and is used to iterate over the raw logs and unpacked data for MintPyFromToken events raised by the PendleRouterv4 contract.
type PendleRouterv4MintPyFromTokenIterator struct {
	Event *PendleRouterv4MintPyFromToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4MintPyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4MintPyFromToken)
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
		it.Event = new(PendleRouterv4MintPyFromToken)
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
func (it *PendleRouterv4MintPyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4MintPyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4MintPyFromToken represents a MintPyFromToken event raised by the PendleRouterv4 contract.
type PendleRouterv4MintPyFromToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterMintPyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, YT []common.Address) (*PendleRouterv4MintPyFromTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4MintPyFromTokenIterator{contract: _PendleRouterv4.contract, event: "MintPyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintPyFromToken is a free log subscription operation binding the contract event 0x3193c546cf854c6a4c63afa03b04d35e4242c2761af34a4093fc5daa88dd5308.
//
// Solidity: event MintPyFromToken(address indexed caller, address indexed tokenIn, address indexed YT, address receiver, uint256 netTokenIn, uint256 netPyOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchMintPyFromToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4MintPyFromToken, caller []common.Address, tokenIn []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4MintPyFromToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseMintPyFromToken(log types.Log) (*PendleRouterv4MintPyFromToken, error) {
	event := new(PendleRouterv4MintPyFromToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4MintSyFromTokenIterator is returned from FilterMintSyFromToken and is used to iterate over the raw logs and unpacked data for MintSyFromToken events raised by the PendleRouterv4 contract.
type PendleRouterv4MintSyFromTokenIterator struct {
	Event *PendleRouterv4MintSyFromToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4MintSyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4MintSyFromToken)
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
		it.Event = new(PendleRouterv4MintSyFromToken)
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
func (it *PendleRouterv4MintSyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4MintSyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4MintSyFromToken represents a MintSyFromToken event raised by the PendleRouterv4 contract.
type PendleRouterv4MintSyFromToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterMintSyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, SY []common.Address) (*PendleRouterv4MintSyFromTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4MintSyFromTokenIterator{contract: _PendleRouterv4.contract, event: "MintSyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintSyFromToken is a free log subscription operation binding the contract event 0x71c7a44161eb32e4640f6c8f0586db5f1d2e03306e2c63bb2e0f7cd0a8fc690c.
//
// Solidity: event MintSyFromToken(address indexed caller, address indexed tokenIn, address indexed SY, address receiver, uint256 netTokenIn, uint256 netSyOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchMintSyFromToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4MintSyFromToken, caller []common.Address, tokenIn []common.Address, SY []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4MintSyFromToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseMintSyFromToken(log types.Log) (*PendleRouterv4MintSyFromToken, error) {
	event := new(PendleRouterv4MintSyFromToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PendleRouterv4 contract.
type PendleRouterv4OwnershipTransferredIterator struct {
	Event *PendleRouterv4OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4OwnershipTransferred)
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
		it.Event = new(PendleRouterv4OwnershipTransferred)
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
func (it *PendleRouterv4OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4OwnershipTransferred represents a OwnershipTransferred event raised by the PendleRouterv4 contract.
type PendleRouterv4OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PendleRouterv4OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4OwnershipTransferredIterator{contract: _PendleRouterv4.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PendleRouterv4OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4OwnershipTransferred)
				if err := _PendleRouterv4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseOwnershipTransferred(log types.Log) (*PendleRouterv4OwnershipTransferred, error) {
	event := new(PendleRouterv4OwnershipTransferred)
	if err := _PendleRouterv4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RedeemPyToSyIterator is returned from FilterRedeemPyToSy and is used to iterate over the raw logs and unpacked data for RedeemPyToSy events raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemPyToSyIterator struct {
	Event *PendleRouterv4RedeemPyToSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RedeemPyToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RedeemPyToSy)
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
		it.Event = new(PendleRouterv4RedeemPyToSy)
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
func (it *PendleRouterv4RedeemPyToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RedeemPyToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RedeemPyToSy represents a RedeemPyToSy event raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemPyToSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRedeemPyToSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*PendleRouterv4RedeemPyToSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RedeemPyToSyIterator{contract: _PendleRouterv4.contract, event: "RedeemPyToSy", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToSy is a free log subscription operation binding the contract event 0x31af33f80f4b396e3d4e42b38ecd3e022883a9bf689fd63f47afbe1d389cb6e7.
//
// Solidity: event RedeemPyToSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netPyIn, uint256 netSyOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRedeemPyToSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RedeemPyToSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RedeemPyToSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRedeemPyToSy(log types.Log) (*PendleRouterv4RedeemPyToSy, error) {
	event := new(PendleRouterv4RedeemPyToSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RedeemPyToTokenIterator is returned from FilterRedeemPyToToken and is used to iterate over the raw logs and unpacked data for RedeemPyToToken events raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemPyToTokenIterator struct {
	Event *PendleRouterv4RedeemPyToToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RedeemPyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RedeemPyToToken)
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
		it.Event = new(PendleRouterv4RedeemPyToToken)
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
func (it *PendleRouterv4RedeemPyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RedeemPyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RedeemPyToToken represents a RedeemPyToToken event raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemPyToToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRedeemPyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, YT []common.Address) (*PendleRouterv4RedeemPyToTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RedeemPyToTokenIterator{contract: _PendleRouterv4.contract, event: "RedeemPyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToToken is a free log subscription operation binding the contract event 0x5f2e0499a3b6a21fd5e1fac44ac47c9aa7c3afa39076d67162a4993411d496da.
//
// Solidity: event RedeemPyToToken(address indexed caller, address indexed tokenOut, address indexed YT, address receiver, uint256 netPyIn, uint256 netTokenOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRedeemPyToToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RedeemPyToToken, caller []common.Address, tokenOut []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RedeemPyToToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRedeemPyToToken(log types.Log) (*PendleRouterv4RedeemPyToToken, error) {
	event := new(PendleRouterv4RedeemPyToToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RedeemSyToTokenIterator is returned from FilterRedeemSyToToken and is used to iterate over the raw logs and unpacked data for RedeemSyToToken events raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemSyToTokenIterator struct {
	Event *PendleRouterv4RedeemSyToToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RedeemSyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RedeemSyToToken)
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
		it.Event = new(PendleRouterv4RedeemSyToToken)
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
func (it *PendleRouterv4RedeemSyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RedeemSyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RedeemSyToToken represents a RedeemSyToToken event raised by the PendleRouterv4 contract.
type PendleRouterv4RedeemSyToToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRedeemSyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, SY []common.Address) (*PendleRouterv4RedeemSyToTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RedeemSyToTokenIterator{contract: _PendleRouterv4.contract, event: "RedeemSyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemSyToToken is a free log subscription operation binding the contract event 0xcd34b6ac7e4b72ab30845649aef2f4fd41945ae2dc08f625be69738bbd0f9aa9.
//
// Solidity: event RedeemSyToToken(address indexed caller, address indexed tokenOut, address indexed SY, address receiver, uint256 netSyIn, uint256 netTokenOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRedeemSyToToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RedeemSyToToken, caller []common.Address, tokenOut []common.Address, SY []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RedeemSyToToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRedeemSyToToken(log types.Log) (*PendleRouterv4RedeemSyToToken, error) {
	event := new(PendleRouterv4RedeemSyToToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RemoveLiquidityDualSyAndPtIterator is returned from FilterRemoveLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualSyAndPt events raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquidityDualSyAndPtIterator struct {
	Event *PendleRouterv4RemoveLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RemoveLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RemoveLiquidityDualSyAndPt)
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
		it.Event = new(PendleRouterv4RemoveLiquidityDualSyAndPt)
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
func (it *PendleRouterv4RemoveLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RemoveLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RemoveLiquidityDualSyAndPt represents a RemoveLiquidityDualSyAndPt event raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquidityDualSyAndPt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRemoveLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4RemoveLiquidityDualSyAndPtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RemoveLiquidityDualSyAndPtIterator{contract: _PendleRouterv4.contract, event: "RemoveLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0xd9f35a37b64d95edfd8f26adf130ce45f3e9ddf3c7ab8c1fb7224727a339a98e.
//
// Solidity: event RemoveLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netSyOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRemoveLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RemoveLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RemoveLiquidityDualSyAndPt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRemoveLiquidityDualSyAndPt(log types.Log) (*PendleRouterv4RemoveLiquidityDualSyAndPt, error) {
	event := new(PendleRouterv4RemoveLiquidityDualSyAndPt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RemoveLiquidityDualTokenAndPtIterator is returned from FilterRemoveLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualTokenAndPt events raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquidityDualTokenAndPtIterator struct {
	Event *PendleRouterv4RemoveLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RemoveLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RemoveLiquidityDualTokenAndPt)
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
		it.Event = new(PendleRouterv4RemoveLiquidityDualTokenAndPt)
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
func (it *PendleRouterv4RemoveLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RemoveLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RemoveLiquidityDualTokenAndPt represents a RemoveLiquidityDualTokenAndPt event raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquidityDualTokenAndPt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRemoveLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenOut []common.Address) (*PendleRouterv4RemoveLiquidityDualTokenAndPtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RemoveLiquidityDualTokenAndPtIterator{contract: _PendleRouterv4.contract, event: "RemoveLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x5349e52482e38bcf6018163f5f871bbae5e00e667aa8e7c531b74c07d5397f92.
//
// Solidity: event RemoveLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenOut, address receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netTokenOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRemoveLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RemoveLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RemoveLiquidityDualTokenAndPt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRemoveLiquidityDualTokenAndPt(log types.Log) (*PendleRouterv4RemoveLiquidityDualTokenAndPt, error) {
	event := new(PendleRouterv4RemoveLiquidityDualTokenAndPt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RemoveLiquiditySinglePtIterator is returned from FilterRemoveLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySinglePt events raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySinglePtIterator struct {
	Event *PendleRouterv4RemoveLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RemoveLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RemoveLiquiditySinglePt)
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
		it.Event = new(PendleRouterv4RemoveLiquiditySinglePt)
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
func (it *PendleRouterv4RemoveLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RemoveLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RemoveLiquiditySinglePt represents a RemoveLiquiditySinglePt event raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySinglePt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRemoveLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4RemoveLiquiditySinglePtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RemoveLiquiditySinglePtIterator{contract: _PendleRouterv4.contract, event: "RemoveLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySinglePt is a free log subscription operation binding the contract event 0xac97b87f5422fa3beec99bff8f336310d8ebc7d33d909b7d534cd7c72f61e871.
//
// Solidity: event RemoveLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRemoveLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RemoveLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RemoveLiquiditySinglePt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRemoveLiquiditySinglePt(log types.Log) (*PendleRouterv4RemoveLiquiditySinglePt, error) {
	event := new(PendleRouterv4RemoveLiquiditySinglePt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RemoveLiquiditySingleSyIterator is returned from FilterRemoveLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleSy events raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySingleSyIterator struct {
	Event *PendleRouterv4RemoveLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RemoveLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RemoveLiquiditySingleSy)
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
		it.Event = new(PendleRouterv4RemoveLiquiditySingleSy)
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
func (it *PendleRouterv4RemoveLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RemoveLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RemoveLiquiditySingleSy represents a RemoveLiquiditySingleSy event raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySingleSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRemoveLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4RemoveLiquiditySingleSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RemoveLiquiditySingleSyIterator{contract: _PendleRouterv4.contract, event: "RemoveLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleSy is a free log subscription operation binding the contract event 0xd31f02c44717b409d13b92ec9d2eaf1427fb4e63f85f4777f1458fb8d9387761.
//
// Solidity: event RemoveLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netSyOut)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRemoveLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RemoveLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RemoveLiquiditySingleSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRemoveLiquiditySingleSy(log types.Log) (*PendleRouterv4RemoveLiquiditySingleSy, error) {
	event := new(PendleRouterv4RemoveLiquiditySingleSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4RemoveLiquiditySingleTokenIterator is returned from FilterRemoveLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleToken events raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySingleTokenIterator struct {
	Event *PendleRouterv4RemoveLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4RemoveLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4RemoveLiquiditySingleToken)
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
		it.Event = new(PendleRouterv4RemoveLiquiditySingleToken)
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
func (it *PendleRouterv4RemoveLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4RemoveLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4RemoveLiquiditySingleToken represents a RemoveLiquiditySingleToken event raised by the PendleRouterv4 contract.
type PendleRouterv4RemoveLiquiditySingleToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterRemoveLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4RemoveLiquiditySingleTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4RemoveLiquiditySingleTokenIterator{contract: _PendleRouterv4.contract, event: "RemoveLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleToken is a free log subscription operation binding the contract event 0x5258a3c624debb1cc84b0f5f66c73eef048832eeebe7178e63e95a53cf28dc94.
//
// Solidity: event RemoveLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpToRemove, uint256 netTokenOut, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchRemoveLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4RemoveLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4RemoveLiquiditySingleToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseRemoveLiquiditySingleToken(log types.Log) (*PendleRouterv4RemoveLiquiditySingleToken, error) {
	event := new(PendleRouterv4RemoveLiquiditySingleToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SelectorToFacetSetIterator is returned from FilterSelectorToFacetSet and is used to iterate over the raw logs and unpacked data for SelectorToFacetSet events raised by the PendleRouterv4 contract.
type PendleRouterv4SelectorToFacetSetIterator struct {
	Event *PendleRouterv4SelectorToFacetSet // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SelectorToFacetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SelectorToFacetSet)
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
		it.Event = new(PendleRouterv4SelectorToFacetSet)
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
func (it *PendleRouterv4SelectorToFacetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SelectorToFacetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SelectorToFacetSet represents a SelectorToFacetSet event raised by the PendleRouterv4 contract.
type PendleRouterv4SelectorToFacetSet struct {
	Selector [4]byte
	Facet    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSelectorToFacetSet is a free log retrieval operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSelectorToFacetSet(opts *bind.FilterOpts, selector [][4]byte, facet []common.Address) (*PendleRouterv4SelectorToFacetSetIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SelectorToFacetSetIterator{contract: _PendleRouterv4.contract, event: "SelectorToFacetSet", logs: logs, sub: sub}, nil
}

// WatchSelectorToFacetSet is a free log subscription operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSelectorToFacetSet(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SelectorToFacetSet, selector [][4]byte, facet []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SelectorToFacetSet)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSelectorToFacetSet(log types.Log) (*PendleRouterv4SelectorToFacetSet, error) {
	event := new(PendleRouterv4SelectorToFacetSet)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SwapPtAndSyIterator is returned from FilterSwapPtAndSy and is used to iterate over the raw logs and unpacked data for SwapPtAndSy events raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndSyIterator struct {
	Event *PendleRouterv4SwapPtAndSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SwapPtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SwapPtAndSy)
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
		it.Event = new(PendleRouterv4SwapPtAndSy)
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
func (it *PendleRouterv4SwapPtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SwapPtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SwapPtAndSy represents a SwapPtAndSy event raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSwapPtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4SwapPtAndSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SwapPtAndSyIterator{contract: _PendleRouterv4.contract, event: "SwapPtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndSy is a free log subscription operation binding the contract event 0x3f5e2944826baeaed8eb77f0f74e6088a154a0fc1317f062fd984585607b4739.
//
// Solidity: event SwapPtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netSyToAccount)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSwapPtAndSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SwapPtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SwapPtAndSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSwapPtAndSy(log types.Log) (*PendleRouterv4SwapPtAndSy, error) {
	event := new(PendleRouterv4SwapPtAndSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SwapPtAndTokenIterator is returned from FilterSwapPtAndToken and is used to iterate over the raw logs and unpacked data for SwapPtAndToken events raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndTokenIterator struct {
	Event *PendleRouterv4SwapPtAndToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SwapPtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SwapPtAndToken)
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
		it.Event = new(PendleRouterv4SwapPtAndToken)
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
func (it *PendleRouterv4SwapPtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SwapPtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SwapPtAndToken represents a SwapPtAndToken event raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSwapPtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4SwapPtAndTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SwapPtAndTokenIterator{contract: _PendleRouterv4.contract, event: "SwapPtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndToken is a free log subscription operation binding the contract event 0xd3c1d9b397236779b29ee5b5b150c1110fc8221b6b6ec0be49c9f4860ceb2036.
//
// Solidity: event SwapPtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netPtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSwapPtAndToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SwapPtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SwapPtAndToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSwapPtAndToken(log types.Log) (*PendleRouterv4SwapPtAndToken, error) {
	event := new(PendleRouterv4SwapPtAndToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SwapPtAndYtIterator is returned from FilterSwapPtAndYt and is used to iterate over the raw logs and unpacked data for SwapPtAndYt events raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndYtIterator struct {
	Event *PendleRouterv4SwapPtAndYt // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SwapPtAndYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SwapPtAndYt)
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
		it.Event = new(PendleRouterv4SwapPtAndYt)
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
func (it *PendleRouterv4SwapPtAndYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SwapPtAndYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SwapPtAndYt represents a SwapPtAndYt event raised by the PendleRouterv4 contract.
type PendleRouterv4SwapPtAndYt struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSwapPtAndYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4SwapPtAndYtIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SwapPtAndYtIterator{contract: _PendleRouterv4.contract, event: "SwapPtAndYt", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndYt is a free log subscription operation binding the contract event 0xa4519acd6c251a3f533922c5aaf3fbae71546aad6c01f8241e23143519a67ac8.
//
// Solidity: event SwapPtAndYt(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netYtToAccount)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSwapPtAndYt(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SwapPtAndYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SwapPtAndYt)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSwapPtAndYt(log types.Log) (*PendleRouterv4SwapPtAndYt, error) {
	event := new(PendleRouterv4SwapPtAndYt)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SwapYtAndSyIterator is returned from FilterSwapYtAndSy and is used to iterate over the raw logs and unpacked data for SwapYtAndSy events raised by the PendleRouterv4 contract.
type PendleRouterv4SwapYtAndSyIterator struct {
	Event *PendleRouterv4SwapYtAndSy // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SwapYtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SwapYtAndSy)
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
		it.Event = new(PendleRouterv4SwapYtAndSy)
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
func (it *PendleRouterv4SwapYtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SwapYtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SwapYtAndSy represents a SwapYtAndSy event raised by the PendleRouterv4 contract.
type PendleRouterv4SwapYtAndSy struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSwapYtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*PendleRouterv4SwapYtAndSyIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SwapYtAndSyIterator{contract: _PendleRouterv4.contract, event: "SwapYtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndSy is a free log subscription operation binding the contract event 0x05499aba408f669fb848399c146fad5bd604d50b15566bdc19e81c40922fab8d.
//
// Solidity: event SwapYtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netYtToAccount, int256 netSyToAccount)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSwapYtAndSy(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SwapYtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SwapYtAndSy)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSwapYtAndSy(log types.Log) (*PendleRouterv4SwapYtAndSy, error) {
	event := new(PendleRouterv4SwapYtAndSy)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleRouterv4SwapYtAndTokenIterator is returned from FilterSwapYtAndToken and is used to iterate over the raw logs and unpacked data for SwapYtAndToken events raised by the PendleRouterv4 contract.
type PendleRouterv4SwapYtAndTokenIterator struct {
	Event *PendleRouterv4SwapYtAndToken // Event containing the contract specifics and raw log

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
func (it *PendleRouterv4SwapYtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleRouterv4SwapYtAndToken)
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
		it.Event = new(PendleRouterv4SwapYtAndToken)
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
func (it *PendleRouterv4SwapYtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleRouterv4SwapYtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleRouterv4SwapYtAndToken represents a SwapYtAndToken event raised by the PendleRouterv4 contract.
type PendleRouterv4SwapYtAndToken struct {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) FilterSwapYtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*PendleRouterv4SwapYtAndTokenIterator, error) {

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

	logs, sub, err := _PendleRouterv4.contract.FilterLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PendleRouterv4SwapYtAndTokenIterator{contract: _PendleRouterv4.contract, event: "SwapYtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndToken is a free log subscription operation binding the contract event 0xa3a2846538c60e47775faa60c6ae79b67dee6d97bb70e386ebbaf4c3a38e8b81.
//
// Solidity: event SwapYtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netYtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_PendleRouterv4 *PendleRouterv4Filterer) WatchSwapYtAndToken(opts *bind.WatchOpts, sink chan<- *PendleRouterv4SwapYtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PendleRouterv4.contract.WatchLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleRouterv4SwapYtAndToken)
				if err := _PendleRouterv4.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
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
func (_PendleRouterv4 *PendleRouterv4Filterer) ParseSwapYtAndToken(log types.Log) (*PendleRouterv4SwapYtAndToken, error) {
	event := new(PendleRouterv4SwapYtAndToken)
	if err := _PendleRouterv4.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
