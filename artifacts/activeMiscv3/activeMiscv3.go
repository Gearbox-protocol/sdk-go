// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activeMiscv3

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

// IPActionMiscV3Call3 is an auto generated low-level Go binding around an user-defined struct.
type IPActionMiscV3Call3 struct {
	AllowFailure bool
	CallData     []byte
}

// IPActionMiscV3Result is an auto generated low-level Go binding around an user-defined struct.
type IPActionMiscV3Result struct {
	Success    bool
	ReturnData []byte
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

// RedeemYtIncomeToTokenStruct is an auto generated low-level Go binding around an user-defined struct.
type RedeemYtIncomeToTokenStruct struct {
	Yt                common.Address
	DoRedeemInterest  bool
	DoRedeemRewards   bool
	TokenRedeemSy     common.Address
	MinTokenRedeemOut *big.Int
}

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	SwapType    uint8
	ExtRouter   common.Address
	ExtCalldata []byte
	NeedScale   bool
}

// SwapDataExtra is an auto generated low-level Go binding around an user-defined struct.
type SwapDataExtra struct {
	TokenIn  common.Address
	TokenOut common.Address
	MinOut   *big.Int
	SwapData SwapData
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

// ActiveMiscv3MetaData contains all meta data concerning the ActiveMiscv3 contract.
var ActiveMiscv3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleSyKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netYtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyMintPy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"AddLiquiditySingleTokenKeepYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPostExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ExitPreExpToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"}],\"name\":\"MintPyFromSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"MintPyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"MintSyFromToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RedeemPyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"}],\"name\":\"RedeemSyToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualSyAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityDualTokenAndPt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netPtOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySinglePt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netLpToRemove\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquiditySingleToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"name\":\"SelectorToFacetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapPtAndToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netPtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"}],\"name\":\"SwapPtAndYt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netSyToAccount\",\"type\":\"int256\"}],\"name\":\"SwapYtAndSy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netYtToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"netTokenToAccount\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"name\":\"SwapYtAndToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"boostMarkets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"reflector\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"selfCall1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"selfCall2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reflectCall\",\"type\":\"bytes\"}],\"name\":\"callAndReflect\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"selfRes1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"selfRes2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reflectRes\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"}],\"name\":\"exitPostExpToSy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"}],\"name\":\"exitPostExpToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"internalType\":\"structExitPostExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"exitPreExpToSy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"limitRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epsSkipMarket\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"normalFills\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumIPLimitOrderType.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lnImpliedRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"failSafeRate\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structFillOrderParams[]\",\"name\":\"flashFills\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"optData\",\"type\":\"bytes\"}],\"internalType\":\"structLimitOrderData\",\"name\":\"limit\",\"type\":\"tuple\"}],\"name\":\"exitPreExpToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"netPtFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPyRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromRedeem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netYtSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFromSwap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSyOut\",\"type\":\"uint256\"}],\"internalType\":\"structExitPreExpReturnParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPyOut\",\"type\":\"uint256\"}],\"name\":\"mintPyFromSy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minPyOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"mintPyFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netPyOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"mintSyFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structIPActionMiscV3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structIPActionMiscV3.Result[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"sys\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"yts\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"markets\",\"type\":\"address[]\"}],\"name\":\"redeemDueInterestAndRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStandardizedYield[]\",\"name\":\"SYs\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"contractIPYieldToken\",\"name\":\"yt\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"doRedeemInterest\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"doRedeemRewards\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenRedeemOut\",\"type\":\"uint256\"}],\"internalType\":\"structRedeemYtIncomeToTokenStruct[]\",\"name\":\"YTs\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIPMarket[]\",\"name\":\"markets\",\"type\":\"address[]\"},{\"internalType\":\"contractIPSwapAggregator\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structSwapDataExtra[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"}],\"name\":\"redeemDueInterestAndRewardsV2\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"netOutFromSwaps\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"netInterests\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSyOut\",\"type\":\"uint256\"}],\"name\":\"redeemPyToSy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netSyOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"YT\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netPyIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"}],\"name\":\"redeemPyToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netSyIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenOutput\",\"name\":\"output\",\"type\":\"tuple\"}],\"name\":\"redeemSyToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"inp\",\"type\":\"tuple\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"SY\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenMintSy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structTokenInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenRedeemSy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minTokenOut\",\"type\":\"uint256\"}],\"name\":\"swapTokenToTokenViaSy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"netTokenOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netSyInterm\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPSwapAggregator\",\"name\":\"pendleSwap\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"extRouter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extCalldata\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"needScale\",\"type\":\"bool\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"internalType\":\"structSwapDataExtra[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"netSwaps\",\"type\":\"uint256[]\"}],\"name\":\"swapTokensToTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"netOutFromSwaps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ActiveMiscv3ABI is the input ABI used to generate the binding from.
// Deprecated: Use ActiveMiscv3MetaData.ABI instead.
var ActiveMiscv3ABI = ActiveMiscv3MetaData.ABI

// ActiveMiscv3 is an auto generated Go binding around an Ethereum contract.
type ActiveMiscv3 struct {
	ActiveMiscv3Caller     // Read-only binding to the contract
	ActiveMiscv3Transactor // Write-only binding to the contract
	ActiveMiscv3Filterer   // Log filterer for contract events
}

// ActiveMiscv3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ActiveMiscv3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveMiscv3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ActiveMiscv3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveMiscv3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActiveMiscv3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActiveMiscv3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActiveMiscv3Session struct {
	Contract     *ActiveMiscv3     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActiveMiscv3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActiveMiscv3CallerSession struct {
	Contract *ActiveMiscv3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ActiveMiscv3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActiveMiscv3TransactorSession struct {
	Contract     *ActiveMiscv3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ActiveMiscv3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ActiveMiscv3Raw struct {
	Contract *ActiveMiscv3 // Generic contract binding to access the raw methods on
}

// ActiveMiscv3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActiveMiscv3CallerRaw struct {
	Contract *ActiveMiscv3Caller // Generic read-only contract binding to access the raw methods on
}

// ActiveMiscv3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActiveMiscv3TransactorRaw struct {
	Contract *ActiveMiscv3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewActiveMiscv3 creates a new instance of ActiveMiscv3, bound to a specific deployed contract.
func NewActiveMiscv3(address common.Address, backend bind.ContractBackend) (*ActiveMiscv3, error) {
	contract, err := bindActiveMiscv3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3{ActiveMiscv3Caller: ActiveMiscv3Caller{contract: contract}, ActiveMiscv3Transactor: ActiveMiscv3Transactor{contract: contract}, ActiveMiscv3Filterer: ActiveMiscv3Filterer{contract: contract}}, nil
}

// NewActiveMiscv3Caller creates a new read-only instance of ActiveMiscv3, bound to a specific deployed contract.
func NewActiveMiscv3Caller(address common.Address, caller bind.ContractCaller) (*ActiveMiscv3Caller, error) {
	contract, err := bindActiveMiscv3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3Caller{contract: contract}, nil
}

// NewActiveMiscv3Transactor creates a new write-only instance of ActiveMiscv3, bound to a specific deployed contract.
func NewActiveMiscv3Transactor(address common.Address, transactor bind.ContractTransactor) (*ActiveMiscv3Transactor, error) {
	contract, err := bindActiveMiscv3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3Transactor{contract: contract}, nil
}

// NewActiveMiscv3Filterer creates a new log filterer instance of ActiveMiscv3, bound to a specific deployed contract.
func NewActiveMiscv3Filterer(address common.Address, filterer bind.ContractFilterer) (*ActiveMiscv3Filterer, error) {
	contract, err := bindActiveMiscv3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3Filterer{contract: contract}, nil
}

// bindActiveMiscv3 binds a generic wrapper to an already deployed contract.
func bindActiveMiscv3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ActiveMiscv3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActiveMiscv3 *ActiveMiscv3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActiveMiscv3.Contract.ActiveMiscv3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActiveMiscv3 *ActiveMiscv3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ActiveMiscv3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActiveMiscv3 *ActiveMiscv3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ActiveMiscv3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActiveMiscv3 *ActiveMiscv3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ActiveMiscv3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActiveMiscv3 *ActiveMiscv3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActiveMiscv3 *ActiveMiscv3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.contract.Transact(opts, method, params...)
}

// BoostMarkets is a paid mutator transaction binding the contract method 0x2d8f9d8d.
//
// Solidity: function boostMarkets(address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3Transactor) BoostMarkets(opts *bind.TransactOpts, markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "boostMarkets", markets)
}

// BoostMarkets is a paid mutator transaction binding the contract method 0x2d8f9d8d.
//
// Solidity: function boostMarkets(address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3Session) BoostMarkets(markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.BoostMarkets(&_ActiveMiscv3.TransactOpts, markets)
}

// BoostMarkets is a paid mutator transaction binding the contract method 0x2d8f9d8d.
//
// Solidity: function boostMarkets(address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) BoostMarkets(markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.BoostMarkets(&_ActiveMiscv3.TransactOpts, markets)
}

// CallAndReflect is a paid mutator transaction binding the contract method 0x9fa02c86.
//
// Solidity: function callAndReflect(address reflector, bytes selfCall1, bytes selfCall2, bytes reflectCall) payable returns(bytes selfRes1, bytes selfRes2, bytes reflectRes)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) CallAndReflect(opts *bind.TransactOpts, reflector common.Address, selfCall1 []byte, selfCall2 []byte, reflectCall []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "callAndReflect", reflector, selfCall1, selfCall2, reflectCall)
}

// CallAndReflect is a paid mutator transaction binding the contract method 0x9fa02c86.
//
// Solidity: function callAndReflect(address reflector, bytes selfCall1, bytes selfCall2, bytes reflectCall) payable returns(bytes selfRes1, bytes selfRes2, bytes reflectRes)
func (_ActiveMiscv3 *ActiveMiscv3Session) CallAndReflect(reflector common.Address, selfCall1 []byte, selfCall2 []byte, reflectCall []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.CallAndReflect(&_ActiveMiscv3.TransactOpts, reflector, selfCall1, selfCall2, reflectCall)
}

// CallAndReflect is a paid mutator transaction binding the contract method 0x9fa02c86.
//
// Solidity: function callAndReflect(address reflector, bytes selfCall1, bytes selfCall2, bytes reflectCall) payable returns(bytes selfRes1, bytes selfRes2, bytes reflectRes)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) CallAndReflect(reflector common.Address, selfCall1 []byte, selfCall2 []byte, reflectCall []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.CallAndReflect(&_ActiveMiscv3.TransactOpts, reflector, selfCall1, selfCall2, reflectCall)
}

// ExitPostExpToSy is a paid mutator transaction binding the contract method 0xc2d6d65d.
//
// Solidity: function exitPostExpToSy(address receiver, address market, uint256 netPtIn, uint256 netLpIn, uint256 minSyOut) returns((uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) ExitPostExpToSy(opts *bind.TransactOpts, receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "exitPostExpToSy", receiver, market, netPtIn, netLpIn, minSyOut)
}

// ExitPostExpToSy is a paid mutator transaction binding the contract method 0xc2d6d65d.
//
// Solidity: function exitPostExpToSy(address receiver, address market, uint256 netPtIn, uint256 netLpIn, uint256 minSyOut) returns((uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Session) ExitPostExpToSy(receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPostExpToSy(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netLpIn, minSyOut)
}

// ExitPostExpToSy is a paid mutator transaction binding the contract method 0xc2d6d65d.
//
// Solidity: function exitPostExpToSy(address receiver, address market, uint256 netPtIn, uint256 netLpIn, uint256 minSyOut) returns((uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) ExitPostExpToSy(receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPostExpToSy(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netLpIn, minSyOut)
}

// ExitPostExpToToken is a paid mutator transaction binding the contract method 0xf06a07a0.
//
// Solidity: function exitPostExpToToken(address receiver, address market, uint256 netPtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) ExitPostExpToToken(opts *bind.TransactOpts, receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "exitPostExpToToken", receiver, market, netPtIn, netLpIn, output)
}

// ExitPostExpToToken is a paid mutator transaction binding the contract method 0xf06a07a0.
//
// Solidity: function exitPostExpToToken(address receiver, address market, uint256 netPtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Session) ExitPostExpToToken(receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPostExpToToken(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netLpIn, output)
}

// ExitPostExpToToken is a paid mutator transaction binding the contract method 0xf06a07a0.
//
// Solidity: function exitPostExpToToken(address receiver, address market, uint256 netPtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) ExitPostExpToToken(receiver common.Address, market common.Address, netPtIn *big.Int, netLpIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPostExpToToken(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netLpIn, output)
}

// ExitPreExpToSy is a paid mutator transaction binding the contract method 0x8354a5e5.
//
// Solidity: function exitPreExpToSy(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) ExitPreExpToSy(opts *bind.TransactOpts, receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "exitPreExpToSy", receiver, market, netPtIn, netYtIn, netLpIn, minSyOut, limit)
}

// ExitPreExpToSy is a paid mutator transaction binding the contract method 0x8354a5e5.
//
// Solidity: function exitPreExpToSy(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Session) ExitPreExpToSy(receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPreExpToSy(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netYtIn, netLpIn, minSyOut, limit)
}

// ExitPreExpToSy is a paid mutator transaction binding the contract method 0x8354a5e5.
//
// Solidity: function exitPreExpToSy(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, uint256 minSyOut, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) ExitPreExpToSy(receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, minSyOut *big.Int, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPreExpToSy(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netYtIn, netLpIn, minSyOut, limit)
}

// ExitPreExpToToken is a paid mutator transaction binding the contract method 0x7036e052.
//
// Solidity: function exitPreExpToToken(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) ExitPreExpToToken(opts *bind.TransactOpts, receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "exitPreExpToToken", receiver, market, netPtIn, netYtIn, netLpIn, output, limit)
}

// ExitPreExpToToken is a paid mutator transaction binding the contract method 0x7036e052.
//
// Solidity: function exitPreExpToToken(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Session) ExitPreExpToToken(receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPreExpToToken(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netYtIn, netLpIn, output, limit)
}

// ExitPreExpToToken is a paid mutator transaction binding the contract method 0x7036e052.
//
// Solidity: function exitPreExpToToken(address receiver, address market, uint256 netPtIn, uint256 netYtIn, uint256 netLpIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output, (address,uint256,((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],((uint256,uint256,uint256,uint8,address,address,address,address,uint256,uint256,uint256,bytes),bytes,uint256)[],bytes) limit) returns(uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) ExitPreExpToToken(receiver common.Address, market common.Address, netPtIn *big.Int, netYtIn *big.Int, netLpIn *big.Int, output TokenOutput, limit LimitOrderData) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.ExitPreExpToToken(&_ActiveMiscv3.TransactOpts, receiver, market, netPtIn, netYtIn, netLpIn, output, limit)
}

// MintPyFromSy is a paid mutator transaction binding the contract method 0x1a8631b2.
//
// Solidity: function mintPyFromSy(address receiver, address YT, uint256 netSyIn, uint256 minPyOut) returns(uint256 netPyOut)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) MintPyFromSy(opts *bind.TransactOpts, receiver common.Address, YT common.Address, netSyIn *big.Int, minPyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "mintPyFromSy", receiver, YT, netSyIn, minPyOut)
}

// MintPyFromSy is a paid mutator transaction binding the contract method 0x1a8631b2.
//
// Solidity: function mintPyFromSy(address receiver, address YT, uint256 netSyIn, uint256 minPyOut) returns(uint256 netPyOut)
func (_ActiveMiscv3 *ActiveMiscv3Session) MintPyFromSy(receiver common.Address, YT common.Address, netSyIn *big.Int, minPyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintPyFromSy(&_ActiveMiscv3.TransactOpts, receiver, YT, netSyIn, minPyOut)
}

// MintPyFromSy is a paid mutator transaction binding the contract method 0x1a8631b2.
//
// Solidity: function mintPyFromSy(address receiver, address YT, uint256 netSyIn, uint256 minPyOut) returns(uint256 netPyOut)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) MintPyFromSy(receiver common.Address, YT common.Address, netSyIn *big.Int, minPyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintPyFromSy(&_ActiveMiscv3.TransactOpts, receiver, YT, netSyIn, minPyOut)
}

// MintPyFromToken is a paid mutator transaction binding the contract method 0xd0f42385.
//
// Solidity: function mintPyFromToken(address receiver, address YT, uint256 minPyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netPyOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) MintPyFromToken(opts *bind.TransactOpts, receiver common.Address, YT common.Address, minPyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "mintPyFromToken", receiver, YT, minPyOut, input)
}

// MintPyFromToken is a paid mutator transaction binding the contract method 0xd0f42385.
//
// Solidity: function mintPyFromToken(address receiver, address YT, uint256 minPyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netPyOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Session) MintPyFromToken(receiver common.Address, YT common.Address, minPyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintPyFromToken(&_ActiveMiscv3.TransactOpts, receiver, YT, minPyOut, input)
}

// MintPyFromToken is a paid mutator transaction binding the contract method 0xd0f42385.
//
// Solidity: function mintPyFromToken(address receiver, address YT, uint256 minPyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netPyOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) MintPyFromToken(receiver common.Address, YT common.Address, minPyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintPyFromToken(&_ActiveMiscv3.TransactOpts, receiver, YT, minPyOut, input)
}

// MintSyFromToken is a paid mutator transaction binding the contract method 0x2e071dc6.
//
// Solidity: function mintSyFromToken(address receiver, address SY, uint256 minSyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) MintSyFromToken(opts *bind.TransactOpts, receiver common.Address, SY common.Address, minSyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "mintSyFromToken", receiver, SY, minSyOut, input)
}

// MintSyFromToken is a paid mutator transaction binding the contract method 0x2e071dc6.
//
// Solidity: function mintSyFromToken(address receiver, address SY, uint256 minSyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Session) MintSyFromToken(receiver common.Address, SY common.Address, minSyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintSyFromToken(&_ActiveMiscv3.TransactOpts, receiver, SY, minSyOut, input)
}

// MintSyFromToken is a paid mutator transaction binding the contract method 0x2e071dc6.
//
// Solidity: function mintSyFromToken(address receiver, address SY, uint256 minSyOut, (address,uint256,address,address,(uint8,address,bytes,bool)) input) payable returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) MintSyFromToken(receiver common.Address, SY common.Address, minSyOut *big.Int, input TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.MintSyFromToken(&_ActiveMiscv3.TransactOpts, receiver, SY, minSyOut, input)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) payable returns((bool,bytes)[] res)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) Multicall(opts *bind.TransactOpts, calls []IPActionMiscV3Call3) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) payable returns((bool,bytes)[] res)
func (_ActiveMiscv3 *ActiveMiscv3Session) Multicall(calls []IPActionMiscV3Call3) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.Multicall(&_ActiveMiscv3.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) payable returns((bool,bytes)[] res)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) Multicall(calls []IPActionMiscV3Call3) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.Multicall(&_ActiveMiscv3.TransactOpts, calls)
}

// RedeemDueInterestAndRewards is a paid mutator transaction binding the contract method 0xf7e375e8.
//
// Solidity: function redeemDueInterestAndRewards(address user, address[] sys, address[] yts, address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3Transactor) RedeemDueInterestAndRewards(opts *bind.TransactOpts, user common.Address, sys []common.Address, yts []common.Address, markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "redeemDueInterestAndRewards", user, sys, yts, markets)
}

// RedeemDueInterestAndRewards is a paid mutator transaction binding the contract method 0xf7e375e8.
//
// Solidity: function redeemDueInterestAndRewards(address user, address[] sys, address[] yts, address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3Session) RedeemDueInterestAndRewards(user common.Address, sys []common.Address, yts []common.Address, markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemDueInterestAndRewards(&_ActiveMiscv3.TransactOpts, user, sys, yts, markets)
}

// RedeemDueInterestAndRewards is a paid mutator transaction binding the contract method 0xf7e375e8.
//
// Solidity: function redeemDueInterestAndRewards(address user, address[] sys, address[] yts, address[] markets) returns()
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) RedeemDueInterestAndRewards(user common.Address, sys []common.Address, yts []common.Address, markets []common.Address) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemDueInterestAndRewards(&_ActiveMiscv3.TransactOpts, user, sys, yts, markets)
}

// RedeemDueInterestAndRewardsV2 is a paid mutator transaction binding the contract method 0x0741a803.
//
// Solidity: function redeemDueInterestAndRewardsV2(address[] SYs, (address,bool,bool,address,uint256)[] YTs, address[] markets, address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps) returns(uint256[] netOutFromSwaps, uint256[] netInterests)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) RedeemDueInterestAndRewardsV2(opts *bind.TransactOpts, SYs []common.Address, YTs []RedeemYtIncomeToTokenStruct, markets []common.Address, pendleSwap common.Address, swaps []SwapDataExtra) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "redeemDueInterestAndRewardsV2", SYs, YTs, markets, pendleSwap, swaps)
}

// RedeemDueInterestAndRewardsV2 is a paid mutator transaction binding the contract method 0x0741a803.
//
// Solidity: function redeemDueInterestAndRewardsV2(address[] SYs, (address,bool,bool,address,uint256)[] YTs, address[] markets, address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps) returns(uint256[] netOutFromSwaps, uint256[] netInterests)
func (_ActiveMiscv3 *ActiveMiscv3Session) RedeemDueInterestAndRewardsV2(SYs []common.Address, YTs []RedeemYtIncomeToTokenStruct, markets []common.Address, pendleSwap common.Address, swaps []SwapDataExtra) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemDueInterestAndRewardsV2(&_ActiveMiscv3.TransactOpts, SYs, YTs, markets, pendleSwap, swaps)
}

// RedeemDueInterestAndRewardsV2 is a paid mutator transaction binding the contract method 0x0741a803.
//
// Solidity: function redeemDueInterestAndRewardsV2(address[] SYs, (address,bool,bool,address,uint256)[] YTs, address[] markets, address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps) returns(uint256[] netOutFromSwaps, uint256[] netInterests)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) RedeemDueInterestAndRewardsV2(SYs []common.Address, YTs []RedeemYtIncomeToTokenStruct, markets []common.Address, pendleSwap common.Address, swaps []SwapDataExtra) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemDueInterestAndRewardsV2(&_ActiveMiscv3.TransactOpts, SYs, YTs, markets, pendleSwap, swaps)
}

// RedeemPyToSy is a paid mutator transaction binding the contract method 0x339748cb.
//
// Solidity: function redeemPyToSy(address receiver, address YT, uint256 netPyIn, uint256 minSyOut) returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) RedeemPyToSy(opts *bind.TransactOpts, receiver common.Address, YT common.Address, netPyIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "redeemPyToSy", receiver, YT, netPyIn, minSyOut)
}

// RedeemPyToSy is a paid mutator transaction binding the contract method 0x339748cb.
//
// Solidity: function redeemPyToSy(address receiver, address YT, uint256 netPyIn, uint256 minSyOut) returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Session) RedeemPyToSy(receiver common.Address, YT common.Address, netPyIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemPyToSy(&_ActiveMiscv3.TransactOpts, receiver, YT, netPyIn, minSyOut)
}

// RedeemPyToSy is a paid mutator transaction binding the contract method 0x339748cb.
//
// Solidity: function redeemPyToSy(address receiver, address YT, uint256 netPyIn, uint256 minSyOut) returns(uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) RedeemPyToSy(receiver common.Address, YT common.Address, netPyIn *big.Int, minSyOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemPyToSy(&_ActiveMiscv3.TransactOpts, receiver, YT, netPyIn, minSyOut)
}

// RedeemPyToToken is a paid mutator transaction binding the contract method 0x47f1de22.
//
// Solidity: function redeemPyToToken(address receiver, address YT, uint256 netPyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) RedeemPyToToken(opts *bind.TransactOpts, receiver common.Address, YT common.Address, netPyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "redeemPyToToken", receiver, YT, netPyIn, output)
}

// RedeemPyToToken is a paid mutator transaction binding the contract method 0x47f1de22.
//
// Solidity: function redeemPyToToken(address receiver, address YT, uint256 netPyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Session) RedeemPyToToken(receiver common.Address, YT common.Address, netPyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemPyToToken(&_ActiveMiscv3.TransactOpts, receiver, YT, netPyIn, output)
}

// RedeemPyToToken is a paid mutator transaction binding the contract method 0x47f1de22.
//
// Solidity: function redeemPyToToken(address receiver, address YT, uint256 netPyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) RedeemPyToToken(receiver common.Address, YT common.Address, netPyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemPyToToken(&_ActiveMiscv3.TransactOpts, receiver, YT, netPyIn, output)
}

// RedeemSyToToken is a paid mutator transaction binding the contract method 0x339a5572.
//
// Solidity: function redeemSyToToken(address receiver, address SY, uint256 netSyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) RedeemSyToToken(opts *bind.TransactOpts, receiver common.Address, SY common.Address, netSyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "redeemSyToToken", receiver, SY, netSyIn, output)
}

// RedeemSyToToken is a paid mutator transaction binding the contract method 0x339a5572.
//
// Solidity: function redeemSyToToken(address receiver, address SY, uint256 netSyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3Session) RedeemSyToToken(receiver common.Address, SY common.Address, netSyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemSyToToken(&_ActiveMiscv3.TransactOpts, receiver, SY, netSyIn, output)
}

// RedeemSyToToken is a paid mutator transaction binding the contract method 0x339a5572.
//
// Solidity: function redeemSyToToken(address receiver, address SY, uint256 netSyIn, (address,uint256,address,address,(uint8,address,bytes,bool)) output) returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) RedeemSyToToken(receiver common.Address, SY common.Address, netSyIn *big.Int, output TokenOutput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.RedeemSyToToken(&_ActiveMiscv3.TransactOpts, receiver, SY, netSyIn, output)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) payable returns()
func (_ActiveMiscv3 *ActiveMiscv3Transactor) Simulate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "simulate", target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) payable returns()
func (_ActiveMiscv3 *ActiveMiscv3Session) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.Simulate(&_ActiveMiscv3.TransactOpts, target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) payable returns()
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.Simulate(&_ActiveMiscv3.TransactOpts, target, data)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x5d3e105c.
//
// Solidity: function swapTokenToToken(address receiver, uint256 minTokenOut, (address,uint256,address,address,(uint8,address,bytes,bool)) inp) payable returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) SwapTokenToToken(opts *bind.TransactOpts, receiver common.Address, minTokenOut *big.Int, inp TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "swapTokenToToken", receiver, minTokenOut, inp)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x5d3e105c.
//
// Solidity: function swapTokenToToken(address receiver, uint256 minTokenOut, (address,uint256,address,address,(uint8,address,bytes,bool)) inp) payable returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3Session) SwapTokenToToken(receiver common.Address, minTokenOut *big.Int, inp TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokenToToken(&_ActiveMiscv3.TransactOpts, receiver, minTokenOut, inp)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x5d3e105c.
//
// Solidity: function swapTokenToToken(address receiver, uint256 minTokenOut, (address,uint256,address,address,(uint8,address,bytes,bool)) inp) payable returns(uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) SwapTokenToToken(receiver common.Address, minTokenOut *big.Int, inp TokenInput) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokenToToken(&_ActiveMiscv3.TransactOpts, receiver, minTokenOut, inp)
}

// SwapTokenToTokenViaSy is a paid mutator transaction binding the contract method 0xa89eba4a.
//
// Solidity: function swapTokenToTokenViaSy(address receiver, address SY, (address,uint256,address,address,(uint8,address,bytes,bool)) input, address tokenRedeemSy, uint256 minTokenOut) payable returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) SwapTokenToTokenViaSy(opts *bind.TransactOpts, receiver common.Address, SY common.Address, input TokenInput, tokenRedeemSy common.Address, minTokenOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "swapTokenToTokenViaSy", receiver, SY, input, tokenRedeemSy, minTokenOut)
}

// SwapTokenToTokenViaSy is a paid mutator transaction binding the contract method 0xa89eba4a.
//
// Solidity: function swapTokenToTokenViaSy(address receiver, address SY, (address,uint256,address,address,(uint8,address,bytes,bool)) input, address tokenRedeemSy, uint256 minTokenOut) payable returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Session) SwapTokenToTokenViaSy(receiver common.Address, SY common.Address, input TokenInput, tokenRedeemSy common.Address, minTokenOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokenToTokenViaSy(&_ActiveMiscv3.TransactOpts, receiver, SY, input, tokenRedeemSy, minTokenOut)
}

// SwapTokenToTokenViaSy is a paid mutator transaction binding the contract method 0xa89eba4a.
//
// Solidity: function swapTokenToTokenViaSy(address receiver, address SY, (address,uint256,address,address,(uint8,address,bytes,bool)) input, address tokenRedeemSy, uint256 minTokenOut) payable returns(uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) SwapTokenToTokenViaSy(receiver common.Address, SY common.Address, input TokenInput, tokenRedeemSy common.Address, minTokenOut *big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokenToTokenViaSy(&_ActiveMiscv3.TransactOpts, receiver, SY, input, tokenRedeemSy, minTokenOut)
}

// SwapTokensToTokens is a paid mutator transaction binding the contract method 0xa373cf1a.
//
// Solidity: function swapTokensToTokens(address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps, uint256[] netSwaps) payable returns(uint256[] netOutFromSwaps)
func (_ActiveMiscv3 *ActiveMiscv3Transactor) SwapTokensToTokens(opts *bind.TransactOpts, pendleSwap common.Address, swaps []SwapDataExtra, netSwaps []*big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.contract.Transact(opts, "swapTokensToTokens", pendleSwap, swaps, netSwaps)
}

// SwapTokensToTokens is a paid mutator transaction binding the contract method 0xa373cf1a.
//
// Solidity: function swapTokensToTokens(address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps, uint256[] netSwaps) payable returns(uint256[] netOutFromSwaps)
func (_ActiveMiscv3 *ActiveMiscv3Session) SwapTokensToTokens(pendleSwap common.Address, swaps []SwapDataExtra, netSwaps []*big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokensToTokens(&_ActiveMiscv3.TransactOpts, pendleSwap, swaps, netSwaps)
}

// SwapTokensToTokens is a paid mutator transaction binding the contract method 0xa373cf1a.
//
// Solidity: function swapTokensToTokens(address pendleSwap, (address,address,uint256,(uint8,address,bytes,bool))[] swaps, uint256[] netSwaps) payable returns(uint256[] netOutFromSwaps)
func (_ActiveMiscv3 *ActiveMiscv3TransactorSession) SwapTokensToTokens(pendleSwap common.Address, swaps []SwapDataExtra, netSwaps []*big.Int) (*types.Transaction, error) {
	return _ActiveMiscv3.Contract.SwapTokensToTokens(&_ActiveMiscv3.TransactOpts, pendleSwap, swaps, netSwaps)
}

// ActiveMiscv3AddLiquidityDualSyAndPtIterator is returned from FilterAddLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualSyAndPt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquidityDualSyAndPtIterator struct {
	Event *ActiveMiscv3AddLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquidityDualSyAndPt)
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
		it.Event = new(ActiveMiscv3AddLiquidityDualSyAndPt)
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
func (it *ActiveMiscv3AddLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquidityDualSyAndPt represents a AddLiquidityDualSyAndPt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquidityDualSyAndPt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3AddLiquidityDualSyAndPtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquidityDualSyAndPtIterator{contract: _ActiveMiscv3.contract, event: "AddLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0x9334566f6358cd68e33d423fb1c9119c6837e3a2a7a8affaaa5020ed06aec069.
//
// Solidity: event AddLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyUsed, uint256 netPtUsed, uint256 netLpOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquidityDualSyAndPt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquidityDualSyAndPt(log types.Log) (*ActiveMiscv3AddLiquidityDualSyAndPt, error) {
	event := new(ActiveMiscv3AddLiquidityDualSyAndPt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquidityDualTokenAndPtIterator is returned from FilterAddLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for AddLiquidityDualTokenAndPt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquidityDualTokenAndPtIterator struct {
	Event *ActiveMiscv3AddLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquidityDualTokenAndPt)
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
		it.Event = new(ActiveMiscv3AddLiquidityDualTokenAndPt)
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
func (it *ActiveMiscv3AddLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquidityDualTokenAndPt represents a AddLiquidityDualTokenAndPt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquidityDualTokenAndPt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenIn []common.Address) (*ActiveMiscv3AddLiquidityDualTokenAndPtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquidityDualTokenAndPtIterator{contract: _ActiveMiscv3.contract, event: "AddLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchAddLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x8969c3e485cb9f3b23622228064bc63e7350f6cf343dd86ab86169041a91eaae.
//
// Solidity: event AddLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenIn, address receiver, uint256 netTokenUsed, uint256 netPtUsed, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenIn []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquidityDualTokenAndPt", callerRule, marketRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquidityDualTokenAndPt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquidityDualTokenAndPt(log types.Log) (*ActiveMiscv3AddLiquidityDualTokenAndPt, error) {
	event := new(ActiveMiscv3AddLiquidityDualTokenAndPt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquiditySinglePtIterator is returned from FilterAddLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for AddLiquiditySinglePt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySinglePtIterator struct {
	Event *ActiveMiscv3AddLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquiditySinglePt)
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
		it.Event = new(ActiveMiscv3AddLiquiditySinglePt)
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
func (it *ActiveMiscv3AddLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquiditySinglePt represents a AddLiquiditySinglePt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySinglePt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3AddLiquiditySinglePtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquiditySinglePtIterator{contract: _ActiveMiscv3.contract, event: "AddLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySinglePt is a free log subscription operation binding the contract event 0xc87b85efc5055ef177e0092af0d4e624fff4b1d57db748857f65e4b7e4a28a36.
//
// Solidity: event AddLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netPtIn, uint256 netLpOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquiditySinglePt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquiditySinglePt(log types.Log) (*ActiveMiscv3AddLiquiditySinglePt, error) {
	event := new(ActiveMiscv3AddLiquiditySinglePt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquiditySingleSyIterator is returned from FilterAddLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleSyIterator struct {
	Event *ActiveMiscv3AddLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquiditySingleSy)
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
		it.Event = new(ActiveMiscv3AddLiquiditySingleSy)
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
func (it *ActiveMiscv3AddLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquiditySingleSy represents a AddLiquiditySingleSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3AddLiquiditySingleSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquiditySingleSyIterator{contract: _ActiveMiscv3.contract, event: "AddLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSy is a free log subscription operation binding the contract event 0xb51950711c9b21dc7888d41f68a19540231ffb5f0d19d8f75cbccaf90ffa7fa5.
//
// Solidity: event AddLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netLpOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquiditySingleSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquiditySingleSy(log types.Log) (*ActiveMiscv3AddLiquiditySingleSy, error) {
	event := new(ActiveMiscv3AddLiquiditySingleSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquiditySingleSyKeepYtIterator is returned from FilterAddLiquiditySingleSyKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleSyKeepYt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleSyKeepYtIterator struct {
	Event *ActiveMiscv3AddLiquiditySingleSyKeepYt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquiditySingleSyKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquiditySingleSyKeepYt)
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
		it.Event = new(ActiveMiscv3AddLiquiditySingleSyKeepYt)
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
func (it *ActiveMiscv3AddLiquiditySingleSyKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquiditySingleSyKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquiditySingleSyKeepYt represents a AddLiquiditySingleSyKeepYt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleSyKeepYt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquiditySingleSyKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3AddLiquiditySingleSyKeepYtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquiditySingleSyKeepYtIterator{contract: _ActiveMiscv3.contract, event: "AddLiquiditySingleSyKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleSyKeepYt is a free log subscription operation binding the contract event 0x890839d8cbce575f9d1ee3d55bc4d466623de60742c7ad665958f8a9916a54a5.
//
// Solidity: event AddLiquiditySingleSyKeepYt(address indexed caller, address indexed market, address indexed receiver, uint256 netSyIn, uint256 netSyMintPy, uint256 netLpOut, uint256 netYtOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquiditySingleSyKeepYt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquiditySingleSyKeepYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquiditySingleSyKeepYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquiditySingleSyKeepYt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquiditySingleSyKeepYt(log types.Log) (*ActiveMiscv3AddLiquiditySingleSyKeepYt, error) {
	event := new(ActiveMiscv3AddLiquiditySingleSyKeepYt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleSyKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquiditySingleTokenIterator is returned from FilterAddLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleTokenIterator struct {
	Event *ActiveMiscv3AddLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquiditySingleToken)
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
		it.Event = new(ActiveMiscv3AddLiquiditySingleToken)
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
func (it *ActiveMiscv3AddLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquiditySingleToken represents a AddLiquiditySingleToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3AddLiquiditySingleTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquiditySingleTokenIterator{contract: _ActiveMiscv3.contract, event: "AddLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleToken is a free log subscription operation binding the contract event 0x387bf301bf673df0120e2d57e639f0e05e5e03d5336577c4cd83c1bff96e8dee.
//
// Solidity: event AddLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquiditySingleToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquiditySingleToken(log types.Log) (*ActiveMiscv3AddLiquiditySingleToken, error) {
	event := new(ActiveMiscv3AddLiquiditySingleToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator is returned from FilterAddLiquiditySingleTokenKeepYt and is used to iterate over the raw logs and unpacked data for AddLiquiditySingleTokenKeepYt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator struct {
	Event *ActiveMiscv3AddLiquiditySingleTokenKeepYt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3AddLiquiditySingleTokenKeepYt)
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
		it.Event = new(ActiveMiscv3AddLiquiditySingleTokenKeepYt)
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
func (it *ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3AddLiquiditySingleTokenKeepYt represents a AddLiquiditySingleTokenKeepYt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3AddLiquiditySingleTokenKeepYt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterAddLiquiditySingleTokenKeepYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3AddLiquiditySingleTokenKeepYtIterator{contract: _ActiveMiscv3.contract, event: "AddLiquiditySingleTokenKeepYt", logs: logs, sub: sub}, nil
}

// WatchAddLiquiditySingleTokenKeepYt is a free log subscription operation binding the contract event 0xa9b749795680682fbc5a34071e19282bbb23496a8cf9bbd99bf941359bbe65bf.
//
// Solidity: event AddLiquiditySingleTokenKeepYt(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netTokenIn, uint256 netLpOut, uint256 netYtOut, uint256 netSyMintPy, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchAddLiquiditySingleTokenKeepYt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3AddLiquiditySingleTokenKeepYt, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "AddLiquiditySingleTokenKeepYt", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3AddLiquiditySingleTokenKeepYt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseAddLiquiditySingleTokenKeepYt(log types.Log) (*ActiveMiscv3AddLiquiditySingleTokenKeepYt, error) {
	event := new(ActiveMiscv3AddLiquiditySingleTokenKeepYt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "AddLiquiditySingleTokenKeepYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3ExitPostExpToSyIterator is returned from FilterExitPostExpToSy and is used to iterate over the raw logs and unpacked data for ExitPostExpToSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPostExpToSyIterator struct {
	Event *ActiveMiscv3ExitPostExpToSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3ExitPostExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3ExitPostExpToSy)
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
		it.Event = new(ActiveMiscv3ExitPostExpToSy)
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
func (it *ActiveMiscv3ExitPostExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3ExitPostExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3ExitPostExpToSy represents a ExitPostExpToSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPostExpToSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterExitPostExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3ExitPostExpToSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3ExitPostExpToSyIterator{contract: _ActiveMiscv3.contract, event: "ExitPostExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToSy is a free log subscription operation binding the contract event 0x19ded113217988ae612547683a5154518a667a51ba409520760f26303a8a6f58.
//
// Solidity: event ExitPostExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchExitPostExpToSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3ExitPostExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "ExitPostExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3ExitPostExpToSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseExitPostExpToSy(log types.Log) (*ActiveMiscv3ExitPostExpToSy, error) {
	event := new(ActiveMiscv3ExitPostExpToSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPostExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3ExitPostExpToTokenIterator is returned from FilterExitPostExpToToken and is used to iterate over the raw logs and unpacked data for ExitPostExpToToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPostExpToTokenIterator struct {
	Event *ActiveMiscv3ExitPostExpToToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3ExitPostExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3ExitPostExpToToken)
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
		it.Event = new(ActiveMiscv3ExitPostExpToToken)
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
func (it *ActiveMiscv3ExitPostExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3ExitPostExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3ExitPostExpToToken represents a ExitPostExpToToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPostExpToToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterExitPostExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3ExitPostExpToTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3ExitPostExpToTokenIterator{contract: _ActiveMiscv3.contract, event: "ExitPostExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPostExpToToken is a free log subscription operation binding the contract event 0x6a5433209d35fd4b489a9e43d2bc02e9d1a24430d39be6fff13b4bb52a72a7e0.
//
// Solidity: event ExitPostExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchExitPostExpToToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3ExitPostExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "ExitPostExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3ExitPostExpToToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseExitPostExpToToken(log types.Log) (*ActiveMiscv3ExitPostExpToToken, error) {
	event := new(ActiveMiscv3ExitPostExpToToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPostExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3ExitPreExpToSyIterator is returned from FilterExitPreExpToSy and is used to iterate over the raw logs and unpacked data for ExitPreExpToSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPreExpToSyIterator struct {
	Event *ActiveMiscv3ExitPreExpToSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3ExitPreExpToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3ExitPreExpToSy)
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
		it.Event = new(ActiveMiscv3ExitPreExpToSy)
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
func (it *ActiveMiscv3ExitPreExpToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3ExitPreExpToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3ExitPreExpToSy represents a ExitPreExpToSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPreExpToSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterExitPreExpToSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3ExitPreExpToSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3ExitPreExpToSyIterator{contract: _ActiveMiscv3.contract, event: "ExitPreExpToSy", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToSy is a free log subscription operation binding the contract event 0x5d98132a999dd75863cdd284a57d3eb44c9b14d38240d22576dea4f09a73626e.
//
// Solidity: event ExitPreExpToSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpIn, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchExitPreExpToSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3ExitPreExpToSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "ExitPreExpToSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3ExitPreExpToSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseExitPreExpToSy(log types.Log) (*ActiveMiscv3ExitPreExpToSy, error) {
	event := new(ActiveMiscv3ExitPreExpToSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPreExpToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3ExitPreExpToTokenIterator is returned from FilterExitPreExpToToken and is used to iterate over the raw logs and unpacked data for ExitPreExpToToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPreExpToTokenIterator struct {
	Event *ActiveMiscv3ExitPreExpToToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3ExitPreExpToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3ExitPreExpToToken)
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
		it.Event = new(ActiveMiscv3ExitPreExpToToken)
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
func (it *ActiveMiscv3ExitPreExpToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3ExitPreExpToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3ExitPreExpToToken represents a ExitPreExpToToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3ExitPreExpToToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterExitPreExpToToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3ExitPreExpToTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3ExitPreExpToTokenIterator{contract: _ActiveMiscv3.contract, event: "ExitPreExpToToken", logs: logs, sub: sub}, nil
}

// WatchExitPreExpToToken is a free log subscription operation binding the contract event 0xe2e505a9d93e4a8a524a95c07024bbe068fa9972f10bb08f51fd0d0c4e11834a.
//
// Solidity: event ExitPreExpToToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpIn, uint256 totalTokenOut, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchExitPreExpToToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3ExitPreExpToToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "ExitPreExpToToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3ExitPreExpToToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseExitPreExpToToken(log types.Log) (*ActiveMiscv3ExitPreExpToToken, error) {
	event := new(ActiveMiscv3ExitPreExpToToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "ExitPreExpToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3MintPyFromSyIterator is returned from FilterMintPyFromSy and is used to iterate over the raw logs and unpacked data for MintPyFromSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintPyFromSyIterator struct {
	Event *ActiveMiscv3MintPyFromSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3MintPyFromSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3MintPyFromSy)
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
		it.Event = new(ActiveMiscv3MintPyFromSy)
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
func (it *ActiveMiscv3MintPyFromSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3MintPyFromSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3MintPyFromSy represents a MintPyFromSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintPyFromSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterMintPyFromSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*ActiveMiscv3MintPyFromSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3MintPyFromSyIterator{contract: _ActiveMiscv3.contract, event: "MintPyFromSy", logs: logs, sub: sub}, nil
}

// WatchMintPyFromSy is a free log subscription operation binding the contract event 0x52e05e4badd3463bad837f42fe3ba58c739d1b3081cff9bb6eb02a24034d455d.
//
// Solidity: event MintPyFromSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netSyIn, uint256 netPyOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchMintPyFromSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3MintPyFromSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "MintPyFromSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3MintPyFromSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseMintPyFromSy(log types.Log) (*ActiveMiscv3MintPyFromSy, error) {
	event := new(ActiveMiscv3MintPyFromSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "MintPyFromSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3MintPyFromTokenIterator is returned from FilterMintPyFromToken and is used to iterate over the raw logs and unpacked data for MintPyFromToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintPyFromTokenIterator struct {
	Event *ActiveMiscv3MintPyFromToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3MintPyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3MintPyFromToken)
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
		it.Event = new(ActiveMiscv3MintPyFromToken)
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
func (it *ActiveMiscv3MintPyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3MintPyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3MintPyFromToken represents a MintPyFromToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintPyFromToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterMintPyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, YT []common.Address) (*ActiveMiscv3MintPyFromTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3MintPyFromTokenIterator{contract: _ActiveMiscv3.contract, event: "MintPyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintPyFromToken is a free log subscription operation binding the contract event 0x3193c546cf854c6a4c63afa03b04d35e4242c2761af34a4093fc5daa88dd5308.
//
// Solidity: event MintPyFromToken(address indexed caller, address indexed tokenIn, address indexed YT, address receiver, uint256 netTokenIn, uint256 netPyOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchMintPyFromToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3MintPyFromToken, caller []common.Address, tokenIn []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "MintPyFromToken", callerRule, tokenInRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3MintPyFromToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseMintPyFromToken(log types.Log) (*ActiveMiscv3MintPyFromToken, error) {
	event := new(ActiveMiscv3MintPyFromToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "MintPyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3MintSyFromTokenIterator is returned from FilterMintSyFromToken and is used to iterate over the raw logs and unpacked data for MintSyFromToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintSyFromTokenIterator struct {
	Event *ActiveMiscv3MintSyFromToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3MintSyFromTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3MintSyFromToken)
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
		it.Event = new(ActiveMiscv3MintSyFromToken)
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
func (it *ActiveMiscv3MintSyFromTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3MintSyFromTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3MintSyFromToken represents a MintSyFromToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3MintSyFromToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterMintSyFromToken(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, SY []common.Address) (*ActiveMiscv3MintSyFromTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3MintSyFromTokenIterator{contract: _ActiveMiscv3.contract, event: "MintSyFromToken", logs: logs, sub: sub}, nil
}

// WatchMintSyFromToken is a free log subscription operation binding the contract event 0x71c7a44161eb32e4640f6c8f0586db5f1d2e03306e2c63bb2e0f7cd0a8fc690c.
//
// Solidity: event MintSyFromToken(address indexed caller, address indexed tokenIn, address indexed SY, address receiver, uint256 netTokenIn, uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchMintSyFromToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3MintSyFromToken, caller []common.Address, tokenIn []common.Address, SY []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "MintSyFromToken", callerRule, tokenInRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3MintSyFromToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseMintSyFromToken(log types.Log) (*ActiveMiscv3MintSyFromToken, error) {
	event := new(ActiveMiscv3MintSyFromToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "MintSyFromToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ActiveMiscv3 contract.
type ActiveMiscv3OwnershipTransferredIterator struct {
	Event *ActiveMiscv3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3OwnershipTransferred)
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
		it.Event = new(ActiveMiscv3OwnershipTransferred)
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
func (it *ActiveMiscv3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3OwnershipTransferred represents a OwnershipTransferred event raised by the ActiveMiscv3 contract.
type ActiveMiscv3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ActiveMiscv3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3OwnershipTransferredIterator{contract: _ActiveMiscv3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3OwnershipTransferred)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseOwnershipTransferred(log types.Log) (*ActiveMiscv3OwnershipTransferred, error) {
	event := new(ActiveMiscv3OwnershipTransferred)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RedeemPyToSyIterator is returned from FilterRedeemPyToSy and is used to iterate over the raw logs and unpacked data for RedeemPyToSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemPyToSyIterator struct {
	Event *ActiveMiscv3RedeemPyToSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RedeemPyToSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RedeemPyToSy)
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
		it.Event = new(ActiveMiscv3RedeemPyToSy)
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
func (it *ActiveMiscv3RedeemPyToSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RedeemPyToSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RedeemPyToSy represents a RedeemPyToSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemPyToSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRedeemPyToSy(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, YT []common.Address) (*ActiveMiscv3RedeemPyToSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RedeemPyToSyIterator{contract: _ActiveMiscv3.contract, event: "RedeemPyToSy", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToSy is a free log subscription operation binding the contract event 0x31af33f80f4b396e3d4e42b38ecd3e022883a9bf689fd63f47afbe1d389cb6e7.
//
// Solidity: event RedeemPyToSy(address indexed caller, address indexed receiver, address indexed YT, uint256 netPyIn, uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRedeemPyToSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RedeemPyToSy, caller []common.Address, receiver []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RedeemPyToSy", callerRule, receiverRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RedeemPyToSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRedeemPyToSy(log types.Log) (*ActiveMiscv3RedeemPyToSy, error) {
	event := new(ActiveMiscv3RedeemPyToSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemPyToSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RedeemPyToTokenIterator is returned from FilterRedeemPyToToken and is used to iterate over the raw logs and unpacked data for RedeemPyToToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemPyToTokenIterator struct {
	Event *ActiveMiscv3RedeemPyToToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RedeemPyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RedeemPyToToken)
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
		it.Event = new(ActiveMiscv3RedeemPyToToken)
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
func (it *ActiveMiscv3RedeemPyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RedeemPyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RedeemPyToToken represents a RedeemPyToToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemPyToToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRedeemPyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, YT []common.Address) (*ActiveMiscv3RedeemPyToTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RedeemPyToTokenIterator{contract: _ActiveMiscv3.contract, event: "RedeemPyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemPyToToken is a free log subscription operation binding the contract event 0x5f2e0499a3b6a21fd5e1fac44ac47c9aa7c3afa39076d67162a4993411d496da.
//
// Solidity: event RedeemPyToToken(address indexed caller, address indexed tokenOut, address indexed YT, address receiver, uint256 netPyIn, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRedeemPyToToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RedeemPyToToken, caller []common.Address, tokenOut []common.Address, YT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RedeemPyToToken", callerRule, tokenOutRule, YTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RedeemPyToToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRedeemPyToToken(log types.Log) (*ActiveMiscv3RedeemPyToToken, error) {
	event := new(ActiveMiscv3RedeemPyToToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemPyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RedeemSyToTokenIterator is returned from FilterRedeemSyToToken and is used to iterate over the raw logs and unpacked data for RedeemSyToToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemSyToTokenIterator struct {
	Event *ActiveMiscv3RedeemSyToToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RedeemSyToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RedeemSyToToken)
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
		it.Event = new(ActiveMiscv3RedeemSyToToken)
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
func (it *ActiveMiscv3RedeemSyToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RedeemSyToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RedeemSyToToken represents a RedeemSyToToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RedeemSyToToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRedeemSyToToken(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address, SY []common.Address) (*ActiveMiscv3RedeemSyToTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RedeemSyToTokenIterator{contract: _ActiveMiscv3.contract, event: "RedeemSyToToken", logs: logs, sub: sub}, nil
}

// WatchRedeemSyToToken is a free log subscription operation binding the contract event 0xcd34b6ac7e4b72ab30845649aef2f4fd41945ae2dc08f625be69738bbd0f9aa9.
//
// Solidity: event RedeemSyToToken(address indexed caller, address indexed tokenOut, address indexed SY, address receiver, uint256 netSyIn, uint256 netTokenOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRedeemSyToToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RedeemSyToToken, caller []common.Address, tokenOut []common.Address, SY []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RedeemSyToToken", callerRule, tokenOutRule, SYRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RedeemSyToToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRedeemSyToToken(log types.Log) (*ActiveMiscv3RedeemSyToToken, error) {
	event := new(ActiveMiscv3RedeemSyToToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RedeemSyToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RemoveLiquidityDualSyAndPtIterator is returned from FilterRemoveLiquidityDualSyAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualSyAndPt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquidityDualSyAndPtIterator struct {
	Event *ActiveMiscv3RemoveLiquidityDualSyAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RemoveLiquidityDualSyAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RemoveLiquidityDualSyAndPt)
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
		it.Event = new(ActiveMiscv3RemoveLiquidityDualSyAndPt)
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
func (it *ActiveMiscv3RemoveLiquidityDualSyAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RemoveLiquidityDualSyAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RemoveLiquidityDualSyAndPt represents a RemoveLiquidityDualSyAndPt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquidityDualSyAndPt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRemoveLiquidityDualSyAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3RemoveLiquidityDualSyAndPtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RemoveLiquidityDualSyAndPtIterator{contract: _ActiveMiscv3.contract, event: "RemoveLiquidityDualSyAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualSyAndPt is a free log subscription operation binding the contract event 0xd9f35a37b64d95edfd8f26adf130ce45f3e9ddf3c7ab8c1fb7224727a339a98e.
//
// Solidity: event RemoveLiquidityDualSyAndPt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRemoveLiquidityDualSyAndPt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RemoveLiquidityDualSyAndPt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RemoveLiquidityDualSyAndPt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RemoveLiquidityDualSyAndPt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRemoveLiquidityDualSyAndPt(log types.Log) (*ActiveMiscv3RemoveLiquidityDualSyAndPt, error) {
	event := new(ActiveMiscv3RemoveLiquidityDualSyAndPt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquidityDualSyAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator is returned from FilterRemoveLiquidityDualTokenAndPt and is used to iterate over the raw logs and unpacked data for RemoveLiquidityDualTokenAndPt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator struct {
	Event *ActiveMiscv3RemoveLiquidityDualTokenAndPt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RemoveLiquidityDualTokenAndPt)
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
		it.Event = new(ActiveMiscv3RemoveLiquidityDualTokenAndPt)
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
func (it *ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RemoveLiquidityDualTokenAndPt represents a RemoveLiquidityDualTokenAndPt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquidityDualTokenAndPt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRemoveLiquidityDualTokenAndPt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, tokenOut []common.Address) (*ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RemoveLiquidityDualTokenAndPtIterator{contract: _ActiveMiscv3.contract, event: "RemoveLiquidityDualTokenAndPt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityDualTokenAndPt is a free log subscription operation binding the contract event 0x5349e52482e38bcf6018163f5f871bbae5e00e667aa8e7c531b74c07d5397f92.
//
// Solidity: event RemoveLiquidityDualTokenAndPt(address indexed caller, address indexed market, address indexed tokenOut, address receiver, uint256 netLpToRemove, uint256 netPtOut, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRemoveLiquidityDualTokenAndPt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RemoveLiquidityDualTokenAndPt, caller []common.Address, market []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RemoveLiquidityDualTokenAndPt", callerRule, marketRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RemoveLiquidityDualTokenAndPt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRemoveLiquidityDualTokenAndPt(log types.Log) (*ActiveMiscv3RemoveLiquidityDualTokenAndPt, error) {
	event := new(ActiveMiscv3RemoveLiquidityDualTokenAndPt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquidityDualTokenAndPt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RemoveLiquiditySinglePtIterator is returned from FilterRemoveLiquiditySinglePt and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySinglePt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySinglePtIterator struct {
	Event *ActiveMiscv3RemoveLiquiditySinglePt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RemoveLiquiditySinglePtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RemoveLiquiditySinglePt)
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
		it.Event = new(ActiveMiscv3RemoveLiquiditySinglePt)
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
func (it *ActiveMiscv3RemoveLiquiditySinglePtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RemoveLiquiditySinglePtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RemoveLiquiditySinglePt represents a RemoveLiquiditySinglePt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySinglePt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRemoveLiquiditySinglePt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3RemoveLiquiditySinglePtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RemoveLiquiditySinglePtIterator{contract: _ActiveMiscv3.contract, event: "RemoveLiquiditySinglePt", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySinglePt is a free log subscription operation binding the contract event 0xac97b87f5422fa3beec99bff8f336310d8ebc7d33d909b7d534cd7c72f61e871.
//
// Solidity: event RemoveLiquiditySinglePt(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netPtOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRemoveLiquiditySinglePt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RemoveLiquiditySinglePt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RemoveLiquiditySinglePt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RemoveLiquiditySinglePt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRemoveLiquiditySinglePt(log types.Log) (*ActiveMiscv3RemoveLiquiditySinglePt, error) {
	event := new(ActiveMiscv3RemoveLiquiditySinglePt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySinglePt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RemoveLiquiditySingleSyIterator is returned from FilterRemoveLiquiditySingleSy and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySingleSyIterator struct {
	Event *ActiveMiscv3RemoveLiquiditySingleSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RemoveLiquiditySingleSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RemoveLiquiditySingleSy)
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
		it.Event = new(ActiveMiscv3RemoveLiquiditySingleSy)
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
func (it *ActiveMiscv3RemoveLiquiditySingleSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RemoveLiquiditySingleSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RemoveLiquiditySingleSy represents a RemoveLiquiditySingleSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySingleSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRemoveLiquiditySingleSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3RemoveLiquiditySingleSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RemoveLiquiditySingleSyIterator{contract: _ActiveMiscv3.contract, event: "RemoveLiquiditySingleSy", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleSy is a free log subscription operation binding the contract event 0xd31f02c44717b409d13b92ec9d2eaf1427fb4e63f85f4777f1458fb8d9387761.
//
// Solidity: event RemoveLiquiditySingleSy(address indexed caller, address indexed market, address indexed receiver, uint256 netLpToRemove, uint256 netSyOut)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRemoveLiquiditySingleSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RemoveLiquiditySingleSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RemoveLiquiditySingleSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RemoveLiquiditySingleSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRemoveLiquiditySingleSy(log types.Log) (*ActiveMiscv3RemoveLiquiditySingleSy, error) {
	event := new(ActiveMiscv3RemoveLiquiditySingleSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySingleSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3RemoveLiquiditySingleTokenIterator is returned from FilterRemoveLiquiditySingleToken and is used to iterate over the raw logs and unpacked data for RemoveLiquiditySingleToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySingleTokenIterator struct {
	Event *ActiveMiscv3RemoveLiquiditySingleToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3RemoveLiquiditySingleTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3RemoveLiquiditySingleToken)
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
		it.Event = new(ActiveMiscv3RemoveLiquiditySingleToken)
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
func (it *ActiveMiscv3RemoveLiquiditySingleTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3RemoveLiquiditySingleTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3RemoveLiquiditySingleToken represents a RemoveLiquiditySingleToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3RemoveLiquiditySingleToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterRemoveLiquiditySingleToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3RemoveLiquiditySingleTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3RemoveLiquiditySingleTokenIterator{contract: _ActiveMiscv3.contract, event: "RemoveLiquiditySingleToken", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquiditySingleToken is a free log subscription operation binding the contract event 0x5258a3c624debb1cc84b0f5f66c73eef048832eeebe7178e63e95a53cf28dc94.
//
// Solidity: event RemoveLiquiditySingleToken(address indexed caller, address indexed market, address indexed token, address receiver, uint256 netLpToRemove, uint256 netTokenOut, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchRemoveLiquiditySingleToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3RemoveLiquiditySingleToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "RemoveLiquiditySingleToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3RemoveLiquiditySingleToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseRemoveLiquiditySingleToken(log types.Log) (*ActiveMiscv3RemoveLiquiditySingleToken, error) {
	event := new(ActiveMiscv3RemoveLiquiditySingleToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "RemoveLiquiditySingleToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SelectorToFacetSetIterator is returned from FilterSelectorToFacetSet and is used to iterate over the raw logs and unpacked data for SelectorToFacetSet events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SelectorToFacetSetIterator struct {
	Event *ActiveMiscv3SelectorToFacetSet // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SelectorToFacetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SelectorToFacetSet)
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
		it.Event = new(ActiveMiscv3SelectorToFacetSet)
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
func (it *ActiveMiscv3SelectorToFacetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SelectorToFacetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SelectorToFacetSet represents a SelectorToFacetSet event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SelectorToFacetSet struct {
	Selector [4]byte
	Facet    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSelectorToFacetSet is a free log retrieval operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSelectorToFacetSet(opts *bind.FilterOpts, selector [][4]byte, facet []common.Address) (*ActiveMiscv3SelectorToFacetSetIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SelectorToFacetSetIterator{contract: _ActiveMiscv3.contract, event: "SelectorToFacetSet", logs: logs, sub: sub}, nil
}

// WatchSelectorToFacetSet is a free log subscription operation binding the contract event 0x0038aaccfeca40ea50135fcc37980765ee22033e320eaf575624561bdbbe9300.
//
// Solidity: event SelectorToFacetSet(bytes4 indexed selector, address indexed facet)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSelectorToFacetSet(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SelectorToFacetSet, selector [][4]byte, facet []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var facetRule []interface{}
	for _, facetItem := range facet {
		facetRule = append(facetRule, facetItem)
	}

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SelectorToFacetSet", selectorRule, facetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SelectorToFacetSet)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSelectorToFacetSet(log types.Log) (*ActiveMiscv3SelectorToFacetSet, error) {
	event := new(ActiveMiscv3SelectorToFacetSet)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SelectorToFacetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SwapPtAndSyIterator is returned from FilterSwapPtAndSy and is used to iterate over the raw logs and unpacked data for SwapPtAndSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndSyIterator struct {
	Event *ActiveMiscv3SwapPtAndSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SwapPtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SwapPtAndSy)
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
		it.Event = new(ActiveMiscv3SwapPtAndSy)
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
func (it *ActiveMiscv3SwapPtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SwapPtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SwapPtAndSy represents a SwapPtAndSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSwapPtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3SwapPtAndSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SwapPtAndSyIterator{contract: _ActiveMiscv3.contract, event: "SwapPtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndSy is a free log subscription operation binding the contract event 0x3f5e2944826baeaed8eb77f0f74e6088a154a0fc1317f062fd984585607b4739.
//
// Solidity: event SwapPtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netSyToAccount)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSwapPtAndSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SwapPtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SwapPtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SwapPtAndSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSwapPtAndSy(log types.Log) (*ActiveMiscv3SwapPtAndSy, error) {
	event := new(ActiveMiscv3SwapPtAndSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SwapPtAndTokenIterator is returned from FilterSwapPtAndToken and is used to iterate over the raw logs and unpacked data for SwapPtAndToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndTokenIterator struct {
	Event *ActiveMiscv3SwapPtAndToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SwapPtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SwapPtAndToken)
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
		it.Event = new(ActiveMiscv3SwapPtAndToken)
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
func (it *ActiveMiscv3SwapPtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SwapPtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SwapPtAndToken represents a SwapPtAndToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSwapPtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3SwapPtAndTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SwapPtAndTokenIterator{contract: _ActiveMiscv3.contract, event: "SwapPtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndToken is a free log subscription operation binding the contract event 0xd3c1d9b397236779b29ee5b5b150c1110fc8221b6b6ec0be49c9f4860ceb2036.
//
// Solidity: event SwapPtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netPtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSwapPtAndToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SwapPtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SwapPtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SwapPtAndToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSwapPtAndToken(log types.Log) (*ActiveMiscv3SwapPtAndToken, error) {
	event := new(ActiveMiscv3SwapPtAndToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SwapPtAndYtIterator is returned from FilterSwapPtAndYt and is used to iterate over the raw logs and unpacked data for SwapPtAndYt events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndYtIterator struct {
	Event *ActiveMiscv3SwapPtAndYt // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SwapPtAndYtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SwapPtAndYt)
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
		it.Event = new(ActiveMiscv3SwapPtAndYt)
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
func (it *ActiveMiscv3SwapPtAndYtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SwapPtAndYtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SwapPtAndYt represents a SwapPtAndYt event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapPtAndYt struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSwapPtAndYt(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3SwapPtAndYtIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SwapPtAndYtIterator{contract: _ActiveMiscv3.contract, event: "SwapPtAndYt", logs: logs, sub: sub}, nil
}

// WatchSwapPtAndYt is a free log subscription operation binding the contract event 0xa4519acd6c251a3f533922c5aaf3fbae71546aad6c01f8241e23143519a67ac8.
//
// Solidity: event SwapPtAndYt(address indexed caller, address indexed market, address indexed receiver, int256 netPtToAccount, int256 netYtToAccount)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSwapPtAndYt(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SwapPtAndYt, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SwapPtAndYt", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SwapPtAndYt)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSwapPtAndYt(log types.Log) (*ActiveMiscv3SwapPtAndYt, error) {
	event := new(ActiveMiscv3SwapPtAndYt)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapPtAndYt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SwapYtAndSyIterator is returned from FilterSwapYtAndSy and is used to iterate over the raw logs and unpacked data for SwapYtAndSy events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapYtAndSyIterator struct {
	Event *ActiveMiscv3SwapYtAndSy // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SwapYtAndSyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SwapYtAndSy)
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
		it.Event = new(ActiveMiscv3SwapYtAndSy)
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
func (it *ActiveMiscv3SwapYtAndSyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SwapYtAndSyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SwapYtAndSy represents a SwapYtAndSy event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapYtAndSy struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSwapYtAndSy(opts *bind.FilterOpts, caller []common.Address, market []common.Address, receiver []common.Address) (*ActiveMiscv3SwapYtAndSyIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SwapYtAndSyIterator{contract: _ActiveMiscv3.contract, event: "SwapYtAndSy", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndSy is a free log subscription operation binding the contract event 0x05499aba408f669fb848399c146fad5bd604d50b15566bdc19e81c40922fab8d.
//
// Solidity: event SwapYtAndSy(address indexed caller, address indexed market, address indexed receiver, int256 netYtToAccount, int256 netSyToAccount)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSwapYtAndSy(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SwapYtAndSy, caller []common.Address, market []common.Address, receiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SwapYtAndSy", callerRule, marketRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SwapYtAndSy)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSwapYtAndSy(log types.Log) (*ActiveMiscv3SwapYtAndSy, error) {
	event := new(ActiveMiscv3SwapYtAndSy)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapYtAndSy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ActiveMiscv3SwapYtAndTokenIterator is returned from FilterSwapYtAndToken and is used to iterate over the raw logs and unpacked data for SwapYtAndToken events raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapYtAndTokenIterator struct {
	Event *ActiveMiscv3SwapYtAndToken // Event containing the contract specifics and raw log

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
func (it *ActiveMiscv3SwapYtAndTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActiveMiscv3SwapYtAndToken)
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
		it.Event = new(ActiveMiscv3SwapYtAndToken)
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
func (it *ActiveMiscv3SwapYtAndTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActiveMiscv3SwapYtAndTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActiveMiscv3SwapYtAndToken represents a SwapYtAndToken event raised by the ActiveMiscv3 contract.
type ActiveMiscv3SwapYtAndToken struct {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) FilterSwapYtAndToken(opts *bind.FilterOpts, caller []common.Address, market []common.Address, token []common.Address) (*ActiveMiscv3SwapYtAndTokenIterator, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.FilterLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ActiveMiscv3SwapYtAndTokenIterator{contract: _ActiveMiscv3.contract, event: "SwapYtAndToken", logs: logs, sub: sub}, nil
}

// WatchSwapYtAndToken is a free log subscription operation binding the contract event 0xa3a2846538c60e47775faa60c6ae79b67dee6d97bb70e386ebbaf4c3a38e8b81.
//
// Solidity: event SwapYtAndToken(address indexed caller, address indexed market, address indexed token, address receiver, int256 netYtToAccount, int256 netTokenToAccount, uint256 netSyInterm)
func (_ActiveMiscv3 *ActiveMiscv3Filterer) WatchSwapYtAndToken(opts *bind.WatchOpts, sink chan<- *ActiveMiscv3SwapYtAndToken, caller []common.Address, market []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ActiveMiscv3.contract.WatchLogs(opts, "SwapYtAndToken", callerRule, marketRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActiveMiscv3SwapYtAndToken)
				if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
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
func (_ActiveMiscv3 *ActiveMiscv3Filterer) ParseSwapYtAndToken(log types.Log) (*ActiveMiscv3SwapYtAndToken, error) {
	event := new(ActiveMiscv3SwapYtAndToken)
	if err := _ActiveMiscv3.contract.UnpackLog(event, "SwapYtAndToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
