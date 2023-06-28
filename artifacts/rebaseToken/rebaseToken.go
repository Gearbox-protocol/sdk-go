// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rebaseToken

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

// RebaseTokenMetaData contains all meta data concerning the RebaseToken contract.
var RebaseTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDepositedValidators\",\"type\":\"uint256\"}],\"name\":\"unsafeChangeDepositedValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferSharesFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveWithdrawals\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEIP712StETH\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDepositsCount\",\"type\":\"uint256\"},{\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_timeElapsed\",\"type\":\"uint256\"},{\"name\":\"_clValidators\",\"type\":\"uint256\"},{\"name\":\"_clBalance\",\"type\":\"uint256\"},{\"name\":\"_withdrawalVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_elRewardsVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_sharesRequestedToBurn\",\"type\":\"uint256\"},{\"name\":\"_withdrawalFinalizationBatches\",\"type\":\"uint256[]\"},{\"name\":\"_simulatedShareRate\",\"type\":\"uint256\"}],\"name\":\"handleOracleReport\",\"outputs\":[{\"name\":\"postRebaseAmounts\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLidoLocator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositableEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLValidators\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLValidators\",\"type\":\"uint256\"}],\"name\":\"CLValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositedValidators\",\"type\":\"uint256\"}],\"name\":\"DepositedValidatorsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"executionLayerRewardsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postBufferedEther\",\"type\":\"uint256\"}],\"name\":\"ETHDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeElapsed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"TokenRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lidoLocator\",\"type\":\"address\"}],\"name\":\"LidoLocatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eip712StETH\",\"type\":\"address\"}],\"name\":\"EIP712StETHInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"}]",
}

// RebaseTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RebaseTokenMetaData.ABI instead.
var RebaseTokenABI = RebaseTokenMetaData.ABI

// RebaseToken is an auto generated Go binding around an Ethereum contract.
type RebaseToken struct {
	RebaseTokenCaller     // Read-only binding to the contract
	RebaseTokenTransactor // Write-only binding to the contract
	RebaseTokenFilterer   // Log filterer for contract events
}

// RebaseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RebaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebaseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RebaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebaseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RebaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RebaseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RebaseTokenSession struct {
	Contract     *RebaseToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RebaseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RebaseTokenCallerSession struct {
	Contract *RebaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RebaseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RebaseTokenTransactorSession struct {
	Contract     *RebaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RebaseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RebaseTokenRaw struct {
	Contract *RebaseToken // Generic contract binding to access the raw methods on
}

// RebaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RebaseTokenCallerRaw struct {
	Contract *RebaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RebaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RebaseTokenTransactorRaw struct {
	Contract *RebaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRebaseToken creates a new instance of RebaseToken, bound to a specific deployed contract.
func NewRebaseToken(address common.Address, backend bind.ContractBackend) (*RebaseToken, error) {
	contract, err := bindRebaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RebaseToken{RebaseTokenCaller: RebaseTokenCaller{contract: contract}, RebaseTokenTransactor: RebaseTokenTransactor{contract: contract}, RebaseTokenFilterer: RebaseTokenFilterer{contract: contract}}, nil
}

// NewRebaseTokenCaller creates a new read-only instance of RebaseToken, bound to a specific deployed contract.
func NewRebaseTokenCaller(address common.Address, caller bind.ContractCaller) (*RebaseTokenCaller, error) {
	contract, err := bindRebaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenCaller{contract: contract}, nil
}

// NewRebaseTokenTransactor creates a new write-only instance of RebaseToken, bound to a specific deployed contract.
func NewRebaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RebaseTokenTransactor, error) {
	contract, err := bindRebaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenTransactor{contract: contract}, nil
}

// NewRebaseTokenFilterer creates a new log filterer instance of RebaseToken, bound to a specific deployed contract.
func NewRebaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RebaseTokenFilterer, error) {
	contract, err := bindRebaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenFilterer{contract: contract}, nil
}

// bindRebaseToken binds a generic wrapper to an already deployed contract.
func bindRebaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RebaseTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RebaseToken *RebaseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RebaseToken.Contract.RebaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RebaseToken *RebaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.Contract.RebaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RebaseToken *RebaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RebaseToken.Contract.RebaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RebaseToken *RebaseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RebaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RebaseToken *RebaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RebaseToken *RebaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RebaseToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _RebaseToken.Contract.DOMAINSEPARATOR(&_RebaseToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _RebaseToken.Contract.DOMAINSEPARATOR(&_RebaseToken.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) PAUSEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.PAUSEROLE(&_RebaseToken.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) PAUSEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.PAUSEROLE(&_RebaseToken.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) RESUMEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.RESUMEROLE(&_RebaseToken.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) RESUMEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.RESUMEROLE(&_RebaseToken.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) STAKINGCONTROLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _RebaseToken.Contract.STAKINGCONTROLROLE(&_RebaseToken.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _RebaseToken.Contract.STAKINGCONTROLROLE(&_RebaseToken.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) STAKINGPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.STAKINGPAUSEROLE(&_RebaseToken.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _RebaseToken.Contract.STAKINGPAUSEROLE(&_RebaseToken.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) UNSAFECHANGEDEPOSITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _RebaseToken.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_RebaseToken.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _RebaseToken.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_RebaseToken.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_RebaseToken *RebaseTokenCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_RebaseToken *RebaseTokenSession) AllowRecoverability(token common.Address) (bool, error) {
	return _RebaseToken.Contract.AllowRecoverability(&_RebaseToken.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _RebaseToken.Contract.AllowRecoverability(&_RebaseToken.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.Allowance(&_RebaseToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.Allowance(&_RebaseToken.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) AppId() ([32]byte, error) {
	return _RebaseToken.Contract.AppId(&_RebaseToken.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) AppId() ([32]byte, error) {
	return _RebaseToken.Contract.AppId(&_RebaseToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.BalanceOf(&_RebaseToken.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.BalanceOf(&_RebaseToken.CallOpts, _account)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_RebaseToken *RebaseTokenCaller) CanDeposit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "canDeposit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_RebaseToken *RebaseTokenSession) CanDeposit() (bool, error) {
	return _RebaseToken.Contract.CanDeposit(&_RebaseToken.CallOpts)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) CanDeposit() (bool, error) {
	return _RebaseToken.Contract.CanDeposit(&_RebaseToken.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_RebaseToken *RebaseTokenCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_RebaseToken *RebaseTokenSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _RebaseToken.Contract.CanPerform(&_RebaseToken.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _RebaseToken.Contract.CanPerform(&_RebaseToken.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RebaseToken *RebaseTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RebaseToken *RebaseTokenSession) Decimals() (uint8, error) {
	return _RebaseToken.Contract.Decimals(&_RebaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_RebaseToken *RebaseTokenCallerSession) Decimals() (uint8, error) {
	return _RebaseToken.Contract.Decimals(&_RebaseToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_RebaseToken *RebaseTokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_RebaseToken *RebaseTokenSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _RebaseToken.Contract.Eip712Domain(&_RebaseToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_RebaseToken *RebaseTokenCallerSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _RebaseToken.Contract.Eip712Domain(&_RebaseToken.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_RebaseToken *RebaseTokenCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_RebaseToken *RebaseTokenSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _RebaseToken.Contract.GetBeaconStat(&_RebaseToken.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_RebaseToken *RebaseTokenCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _RebaseToken.Contract.GetBeaconStat(&_RebaseToken.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetBufferedEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetBufferedEther(&_RebaseToken.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetBufferedEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetBufferedEther(&_RebaseToken.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetContractVersion() (*big.Int, error) {
	return _RebaseToken.Contract.GetContractVersion(&_RebaseToken.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetContractVersion() (*big.Int, error) {
	return _RebaseToken.Contract.GetContractVersion(&_RebaseToken.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _RebaseToken.Contract.GetCurrentStakeLimit(&_RebaseToken.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _RebaseToken.Contract.GetCurrentStakeLimit(&_RebaseToken.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetDepositableEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getDepositableEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetDepositableEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetDepositableEther(&_RebaseToken.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetDepositableEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetDepositableEther(&_RebaseToken.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetEIP712StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getEIP712StETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetEIP712StETH() (common.Address, error) {
	return _RebaseToken.Contract.GetEIP712StETH(&_RebaseToken.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetEIP712StETH() (common.Address, error) {
	return _RebaseToken.Contract.GetEIP712StETH(&_RebaseToken.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_RebaseToken *RebaseTokenSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _RebaseToken.Contract.GetEVMScriptExecutor(&_RebaseToken.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _RebaseToken.Contract.GetEVMScriptExecutor(&_RebaseToken.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetEVMScriptRegistry() (common.Address, error) {
	return _RebaseToken.Contract.GetEVMScriptRegistry(&_RebaseToken.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _RebaseToken.Contract.GetEVMScriptRegistry(&_RebaseToken.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_RebaseToken *RebaseTokenCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_RebaseToken *RebaseTokenSession) GetFee() (uint16, error) {
	return _RebaseToken.Contract.GetFee(&_RebaseToken.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_RebaseToken *RebaseTokenCallerSession) GetFee() (uint16, error) {
	return _RebaseToken.Contract.GetFee(&_RebaseToken.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_RebaseToken *RebaseTokenCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_RebaseToken *RebaseTokenSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _RebaseToken.Contract.GetFeeDistribution(&_RebaseToken.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_RebaseToken *RebaseTokenCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _RebaseToken.Contract.GetFeeDistribution(&_RebaseToken.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetInitializationBlock() (*big.Int, error) {
	return _RebaseToken.Contract.GetInitializationBlock(&_RebaseToken.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _RebaseToken.Contract.GetInitializationBlock(&_RebaseToken.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetLidoLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getLidoLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetLidoLocator() (common.Address, error) {
	return _RebaseToken.Contract.GetLidoLocator(&_RebaseToken.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetLidoLocator() (common.Address, error) {
	return _RebaseToken.Contract.GetLidoLocator(&_RebaseToken.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetOracle() (common.Address, error) {
	return _RebaseToken.Contract.GetOracle(&_RebaseToken.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetOracle() (common.Address, error) {
	return _RebaseToken.Contract.GetOracle(&_RebaseToken.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _RebaseToken.Contract.GetPooledEthByShares(&_RebaseToken.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _RebaseToken.Contract.GetPooledEthByShares(&_RebaseToken.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetRecoveryVault() (common.Address, error) {
	return _RebaseToken.Contract.GetRecoveryVault(&_RebaseToken.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetRecoveryVault() (common.Address, error) {
	return _RebaseToken.Contract.GetRecoveryVault(&_RebaseToken.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _RebaseToken.Contract.GetSharesByPooledEth(&_RebaseToken.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _RebaseToken.Contract.GetSharesByPooledEth(&_RebaseToken.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_RebaseToken *RebaseTokenCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_RebaseToken *RebaseTokenSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _RebaseToken.Contract.GetStakeLimitFullInfo(&_RebaseToken.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_RebaseToken *RebaseTokenCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _RebaseToken.Contract.GetStakeLimitFullInfo(&_RebaseToken.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalELRewardsCollected(&_RebaseToken.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalELRewardsCollected(&_RebaseToken.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetTotalPooledEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalPooledEther(&_RebaseToken.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalPooledEther(&_RebaseToken.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) GetTotalShares() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalShares(&_RebaseToken.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) GetTotalShares() (*big.Int, error) {
	return _RebaseToken.Contract.GetTotalShares(&_RebaseToken.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_RebaseToken *RebaseTokenCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_RebaseToken *RebaseTokenSession) GetTreasury() (common.Address, error) {
	return _RebaseToken.Contract.GetTreasury(&_RebaseToken.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) GetTreasury() (common.Address, error) {
	return _RebaseToken.Contract.GetTreasury(&_RebaseToken.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_RebaseToken *RebaseTokenCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_RebaseToken *RebaseTokenSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _RebaseToken.Contract.GetWithdrawalCredentials(&_RebaseToken.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_RebaseToken *RebaseTokenCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _RebaseToken.Contract.GetWithdrawalCredentials(&_RebaseToken.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_RebaseToken *RebaseTokenCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_RebaseToken *RebaseTokenSession) HasInitialized() (bool, error) {
	return _RebaseToken.Contract.HasInitialized(&_RebaseToken.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) HasInitialized() (bool, error) {
	return _RebaseToken.Contract.HasInitialized(&_RebaseToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RebaseToken *RebaseTokenCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RebaseToken *RebaseTokenSession) Implementation() (common.Address, error) {
	return _RebaseToken.Contract.Implementation(&_RebaseToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) Implementation() (common.Address, error) {
	return _RebaseToken.Contract.Implementation(&_RebaseToken.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_RebaseToken *RebaseTokenCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_RebaseToken *RebaseTokenSession) IsPetrified() (bool, error) {
	return _RebaseToken.Contract.IsPetrified(&_RebaseToken.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) IsPetrified() (bool, error) {
	return _RebaseToken.Contract.IsPetrified(&_RebaseToken.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_RebaseToken *RebaseTokenCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_RebaseToken *RebaseTokenSession) IsStakingPaused() (bool, error) {
	return _RebaseToken.Contract.IsStakingPaused(&_RebaseToken.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) IsStakingPaused() (bool, error) {
	return _RebaseToken.Contract.IsStakingPaused(&_RebaseToken.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_RebaseToken *RebaseTokenCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_RebaseToken *RebaseTokenSession) IsStopped() (bool, error) {
	return _RebaseToken.Contract.IsStopped(&_RebaseToken.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_RebaseToken *RebaseTokenCallerSession) IsStopped() (bool, error) {
	return _RebaseToken.Contract.IsStopped(&_RebaseToken.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_RebaseToken *RebaseTokenCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_RebaseToken *RebaseTokenSession) Kernel() (common.Address, error) {
	return _RebaseToken.Contract.Kernel(&_RebaseToken.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_RebaseToken *RebaseTokenCallerSession) Kernel() (common.Address, error) {
	return _RebaseToken.Contract.Kernel(&_RebaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_RebaseToken *RebaseTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_RebaseToken *RebaseTokenSession) Name() (string, error) {
	return _RebaseToken.Contract.Name(&_RebaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_RebaseToken *RebaseTokenCallerSession) Name() (string, error) {
	return _RebaseToken.Contract.Name(&_RebaseToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.Nonces(&_RebaseToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.Nonces(&_RebaseToken.CallOpts, owner)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.SharesOf(&_RebaseToken.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _RebaseToken.Contract.SharesOf(&_RebaseToken.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_RebaseToken *RebaseTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_RebaseToken *RebaseTokenSession) Symbol() (string, error) {
	return _RebaseToken.Contract.Symbol(&_RebaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_RebaseToken *RebaseTokenCallerSession) Symbol() (string, error) {
	return _RebaseToken.Contract.Symbol(&_RebaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RebaseToken *RebaseTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RebaseToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RebaseToken *RebaseTokenSession) TotalSupply() (*big.Int, error) {
	return _RebaseToken.Contract.TotalSupply(&_RebaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RebaseToken *RebaseTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RebaseToken.Contract.TotalSupply(&_RebaseToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.Approve(&_RebaseToken.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.Approve(&_RebaseToken.TransactOpts, _spender, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RebaseToken *RebaseTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RebaseToken *RebaseTokenSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.DecreaseAllowance(&_RebaseToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_RebaseToken *RebaseTokenTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.DecreaseAllowance(&_RebaseToken.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_RebaseToken *RebaseTokenTransactor) Deposit(opts *bind.TransactOpts, _maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "deposit", _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_RebaseToken *RebaseTokenSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Deposit(&_RebaseToken.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_RebaseToken *RebaseTokenTransactorSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Deposit(&_RebaseToken.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_RebaseToken *RebaseTokenTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "finalizeUpgrade_v2", _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_RebaseToken *RebaseTokenSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.FinalizeUpgradeV2(&_RebaseToken.TransactOpts, _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_RebaseToken *RebaseTokenTransactorSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.FinalizeUpgradeV2(&_RebaseToken.TransactOpts, _lidoLocator, _eip712StETH)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_RebaseToken *RebaseTokenTransactor) HandleOracleReport(opts *bind.TransactOpts, _reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "handleOracleReport", _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_RebaseToken *RebaseTokenSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.HandleOracleReport(&_RebaseToken.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_RebaseToken *RebaseTokenTransactorSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.HandleOracleReport(&_RebaseToken.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RebaseToken *RebaseTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RebaseToken *RebaseTokenSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.IncreaseAllowance(&_RebaseToken.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_RebaseToken *RebaseTokenTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.IncreaseAllowance(&_RebaseToken.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_RebaseToken *RebaseTokenTransactor) Initialize(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "initialize", _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_RebaseToken *RebaseTokenSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.Initialize(&_RebaseToken.TransactOpts, _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_RebaseToken *RebaseTokenTransactorSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.Initialize(&_RebaseToken.TransactOpts, _lidoLocator, _eip712StETH)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_RebaseToken *RebaseTokenTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_RebaseToken *RebaseTokenSession) PauseStaking() (*types.Transaction, error) {
	return _RebaseToken.Contract.PauseStaking(&_RebaseToken.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_RebaseToken *RebaseTokenTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _RebaseToken.Contract.PauseStaking(&_RebaseToken.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_RebaseToken *RebaseTokenTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_RebaseToken *RebaseTokenSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Permit(&_RebaseToken.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_RebaseToken *RebaseTokenTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Permit(&_RebaseToken.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_RebaseToken *RebaseTokenTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_RebaseToken *RebaseTokenSession) ReceiveELRewards() (*types.Transaction, error) {
	return _RebaseToken.Contract.ReceiveELRewards(&_RebaseToken.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_RebaseToken *RebaseTokenTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _RebaseToken.Contract.ReceiveELRewards(&_RebaseToken.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_RebaseToken *RebaseTokenTransactor) ReceiveWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "receiveWithdrawals")
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_RebaseToken *RebaseTokenSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _RebaseToken.Contract.ReceiveWithdrawals(&_RebaseToken.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_RebaseToken *RebaseTokenTransactorSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _RebaseToken.Contract.ReceiveWithdrawals(&_RebaseToken.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_RebaseToken *RebaseTokenTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_RebaseToken *RebaseTokenSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _RebaseToken.Contract.RemoveStakingLimit(&_RebaseToken.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_RebaseToken *RebaseTokenTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _RebaseToken.Contract.RemoveStakingLimit(&_RebaseToken.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RebaseToken *RebaseTokenTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RebaseToken *RebaseTokenSession) Resume() (*types.Transaction, error) {
	return _RebaseToken.Contract.Resume(&_RebaseToken.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RebaseToken *RebaseTokenTransactorSession) Resume() (*types.Transaction, error) {
	return _RebaseToken.Contract.Resume(&_RebaseToken.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_RebaseToken *RebaseTokenTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_RebaseToken *RebaseTokenSession) ResumeStaking() (*types.Transaction, error) {
	return _RebaseToken.Contract.ResumeStaking(&_RebaseToken.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_RebaseToken *RebaseTokenTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _RebaseToken.Contract.ResumeStaking(&_RebaseToken.TransactOpts)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_RebaseToken *RebaseTokenTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_RebaseToken *RebaseTokenSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.SetStakingLimit(&_RebaseToken.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_RebaseToken *RebaseTokenTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.SetStakingLimit(&_RebaseToken.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_RebaseToken *RebaseTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_RebaseToken *RebaseTokenSession) Stop() (*types.Transaction, error) {
	return _RebaseToken.Contract.Stop(&_RebaseToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_RebaseToken *RebaseTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _RebaseToken.Contract.Stop(&_RebaseToken.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_RebaseToken *RebaseTokenTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_RebaseToken *RebaseTokenSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.Submit(&_RebaseToken.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_RebaseToken *RebaseTokenTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.Submit(&_RebaseToken.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.Transfer(&_RebaseToken.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.Transfer(&_RebaseToken.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferFrom(&_RebaseToken.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RebaseToken *RebaseTokenTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferFrom(&_RebaseToken.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferShares(&_RebaseToken.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferShares(&_RebaseToken.TransactOpts, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenTransactor) TransferSharesFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "transferSharesFrom", _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferSharesFrom(&_RebaseToken.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_RebaseToken *RebaseTokenTransactorSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferSharesFrom(&_RebaseToken.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_RebaseToken *RebaseTokenTransactor) TransferToVault(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "transferToVault", arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_RebaseToken *RebaseTokenSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferToVault(&_RebaseToken.TransactOpts, arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_RebaseToken *RebaseTokenTransactorSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _RebaseToken.Contract.TransferToVault(&_RebaseToken.TransactOpts, arg0)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_RebaseToken *RebaseTokenTransactor) UnsafeChangeDepositedValidators(opts *bind.TransactOpts, _newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _RebaseToken.contract.Transact(opts, "unsafeChangeDepositedValidators", _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_RebaseToken *RebaseTokenSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.UnsafeChangeDepositedValidators(&_RebaseToken.TransactOpts, _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_RebaseToken *RebaseTokenTransactorSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _RebaseToken.Contract.UnsafeChangeDepositedValidators(&_RebaseToken.TransactOpts, _newDepositedValidators)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RebaseToken *RebaseTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RebaseToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RebaseToken *RebaseTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Fallback(&_RebaseToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RebaseToken *RebaseTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RebaseToken.Contract.Fallback(&_RebaseToken.TransactOpts, calldata)
}

// RebaseTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RebaseToken contract.
type RebaseTokenApprovalIterator struct {
	Event *RebaseTokenApproval // Event containing the contract specifics and raw log

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
func (it *RebaseTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenApproval)
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
		it.Event = new(RebaseTokenApproval)
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
func (it *RebaseTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenApproval represents a Approval event raised by the RebaseToken contract.
type RebaseTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RebaseToken *RebaseTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RebaseTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenApprovalIterator{contract: _RebaseToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RebaseToken *RebaseTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RebaseTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenApproval)
				if err := _RebaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RebaseToken *RebaseTokenFilterer) ParseApproval(log types.Log) (*RebaseTokenApproval, error) {
	event := new(RebaseTokenApproval)
	if err := _RebaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenCLValidatorsUpdatedIterator is returned from FilterCLValidatorsUpdated and is used to iterate over the raw logs and unpacked data for CLValidatorsUpdated events raised by the RebaseToken contract.
type RebaseTokenCLValidatorsUpdatedIterator struct {
	Event *RebaseTokenCLValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *RebaseTokenCLValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenCLValidatorsUpdated)
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
		it.Event = new(RebaseTokenCLValidatorsUpdated)
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
func (it *RebaseTokenCLValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenCLValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenCLValidatorsUpdated represents a CLValidatorsUpdated event raised by the RebaseToken contract.
type RebaseTokenCLValidatorsUpdated struct {
	ReportTimestamp  *big.Int
	PreCLValidators  *big.Int
	PostCLValidators *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCLValidatorsUpdated is a free log retrieval operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_RebaseToken *RebaseTokenFilterer) FilterCLValidatorsUpdated(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*RebaseTokenCLValidatorsUpdatedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenCLValidatorsUpdatedIterator{contract: _RebaseToken.contract, event: "CLValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchCLValidatorsUpdated is a free log subscription operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_RebaseToken *RebaseTokenFilterer) WatchCLValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *RebaseTokenCLValidatorsUpdated, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenCLValidatorsUpdated)
				if err := _RebaseToken.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
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

// ParseCLValidatorsUpdated is a log parse operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_RebaseToken *RebaseTokenFilterer) ParseCLValidatorsUpdated(log types.Log) (*RebaseTokenCLValidatorsUpdated, error) {
	event := new(RebaseTokenCLValidatorsUpdated)
	if err := _RebaseToken.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the RebaseToken contract.
type RebaseTokenContractVersionSetIterator struct {
	Event *RebaseTokenContractVersionSet // Event containing the contract specifics and raw log

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
func (it *RebaseTokenContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenContractVersionSet)
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
		it.Event = new(RebaseTokenContractVersionSet)
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
func (it *RebaseTokenContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenContractVersionSet represents a ContractVersionSet event raised by the RebaseToken contract.
type RebaseTokenContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_RebaseToken *RebaseTokenFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*RebaseTokenContractVersionSetIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenContractVersionSetIterator{contract: _RebaseToken.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_RebaseToken *RebaseTokenFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *RebaseTokenContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenContractVersionSet)
				if err := _RebaseToken.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_RebaseToken *RebaseTokenFilterer) ParseContractVersionSet(log types.Log) (*RebaseTokenContractVersionSet, error) {
	event := new(RebaseTokenContractVersionSet)
	if err := _RebaseToken.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenDepositedValidatorsChangedIterator is returned from FilterDepositedValidatorsChanged and is used to iterate over the raw logs and unpacked data for DepositedValidatorsChanged events raised by the RebaseToken contract.
type RebaseTokenDepositedValidatorsChangedIterator struct {
	Event *RebaseTokenDepositedValidatorsChanged // Event containing the contract specifics and raw log

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
func (it *RebaseTokenDepositedValidatorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenDepositedValidatorsChanged)
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
		it.Event = new(RebaseTokenDepositedValidatorsChanged)
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
func (it *RebaseTokenDepositedValidatorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenDepositedValidatorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenDepositedValidatorsChanged represents a DepositedValidatorsChanged event raised by the RebaseToken contract.
type RebaseTokenDepositedValidatorsChanged struct {
	DepositedValidators *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositedValidatorsChanged is a free log retrieval operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_RebaseToken *RebaseTokenFilterer) FilterDepositedValidatorsChanged(opts *bind.FilterOpts) (*RebaseTokenDepositedValidatorsChangedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenDepositedValidatorsChangedIterator{contract: _RebaseToken.contract, event: "DepositedValidatorsChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedValidatorsChanged is a free log subscription operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_RebaseToken *RebaseTokenFilterer) WatchDepositedValidatorsChanged(opts *bind.WatchOpts, sink chan<- *RebaseTokenDepositedValidatorsChanged) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenDepositedValidatorsChanged)
				if err := _RebaseToken.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
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

// ParseDepositedValidatorsChanged is a log parse operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_RebaseToken *RebaseTokenFilterer) ParseDepositedValidatorsChanged(log types.Log) (*RebaseTokenDepositedValidatorsChanged, error) {
	event := new(RebaseTokenDepositedValidatorsChanged)
	if err := _RebaseToken.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenEIP712StETHInitializedIterator is returned from FilterEIP712StETHInitialized and is used to iterate over the raw logs and unpacked data for EIP712StETHInitialized events raised by the RebaseToken contract.
type RebaseTokenEIP712StETHInitializedIterator struct {
	Event *RebaseTokenEIP712StETHInitialized // Event containing the contract specifics and raw log

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
func (it *RebaseTokenEIP712StETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenEIP712StETHInitialized)
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
		it.Event = new(RebaseTokenEIP712StETHInitialized)
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
func (it *RebaseTokenEIP712StETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenEIP712StETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenEIP712StETHInitialized represents a EIP712StETHInitialized event raised by the RebaseToken contract.
type RebaseTokenEIP712StETHInitialized struct {
	Eip712StETH common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEIP712StETHInitialized is a free log retrieval operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_RebaseToken *RebaseTokenFilterer) FilterEIP712StETHInitialized(opts *bind.FilterOpts) (*RebaseTokenEIP712StETHInitializedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenEIP712StETHInitializedIterator{contract: _RebaseToken.contract, event: "EIP712StETHInitialized", logs: logs, sub: sub}, nil
}

// WatchEIP712StETHInitialized is a free log subscription operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_RebaseToken *RebaseTokenFilterer) WatchEIP712StETHInitialized(opts *bind.WatchOpts, sink chan<- *RebaseTokenEIP712StETHInitialized) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenEIP712StETHInitialized)
				if err := _RebaseToken.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
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

// ParseEIP712StETHInitialized is a log parse operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_RebaseToken *RebaseTokenFilterer) ParseEIP712StETHInitialized(log types.Log) (*RebaseTokenEIP712StETHInitialized, error) {
	event := new(RebaseTokenEIP712StETHInitialized)
	if err := _RebaseToken.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the RebaseToken contract.
type RebaseTokenELRewardsReceivedIterator struct {
	Event *RebaseTokenELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *RebaseTokenELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenELRewardsReceived)
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
		it.Event = new(RebaseTokenELRewardsReceived)
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
func (it *RebaseTokenELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenELRewardsReceived represents a ELRewardsReceived event raised by the RebaseToken contract.
type RebaseTokenELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*RebaseTokenELRewardsReceivedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenELRewardsReceivedIterator{contract: _RebaseToken.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *RebaseTokenELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenELRewardsReceived)
				if err := _RebaseToken.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) ParseELRewardsReceived(log types.Log) (*RebaseTokenELRewardsReceived, error) {
	event := new(RebaseTokenELRewardsReceived)
	if err := _RebaseToken.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenETHDistributedIterator is returned from FilterETHDistributed and is used to iterate over the raw logs and unpacked data for ETHDistributed events raised by the RebaseToken contract.
type RebaseTokenETHDistributedIterator struct {
	Event *RebaseTokenETHDistributed // Event containing the contract specifics and raw log

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
func (it *RebaseTokenETHDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenETHDistributed)
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
		it.Event = new(RebaseTokenETHDistributed)
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
func (it *RebaseTokenETHDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenETHDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenETHDistributed represents a ETHDistributed event raised by the RebaseToken contract.
type RebaseTokenETHDistributed struct {
	ReportTimestamp                *big.Int
	PreCLBalance                   *big.Int
	PostCLBalance                  *big.Int
	WithdrawalsWithdrawn           *big.Int
	ExecutionLayerRewardsWithdrawn *big.Int
	PostBufferedEther              *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterETHDistributed is a free log retrieval operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_RebaseToken *RebaseTokenFilterer) FilterETHDistributed(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*RebaseTokenETHDistributedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenETHDistributedIterator{contract: _RebaseToken.contract, event: "ETHDistributed", logs: logs, sub: sub}, nil
}

// WatchETHDistributed is a free log subscription operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_RebaseToken *RebaseTokenFilterer) WatchETHDistributed(opts *bind.WatchOpts, sink chan<- *RebaseTokenETHDistributed, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenETHDistributed)
				if err := _RebaseToken.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
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

// ParseETHDistributed is a log parse operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_RebaseToken *RebaseTokenFilterer) ParseETHDistributed(log types.Log) (*RebaseTokenETHDistributed, error) {
	event := new(RebaseTokenETHDistributed)
	if err := _RebaseToken.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenLidoLocatorSetIterator is returned from FilterLidoLocatorSet and is used to iterate over the raw logs and unpacked data for LidoLocatorSet events raised by the RebaseToken contract.
type RebaseTokenLidoLocatorSetIterator struct {
	Event *RebaseTokenLidoLocatorSet // Event containing the contract specifics and raw log

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
func (it *RebaseTokenLidoLocatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenLidoLocatorSet)
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
		it.Event = new(RebaseTokenLidoLocatorSet)
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
func (it *RebaseTokenLidoLocatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenLidoLocatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenLidoLocatorSet represents a LidoLocatorSet event raised by the RebaseToken contract.
type RebaseTokenLidoLocatorSet struct {
	LidoLocator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLidoLocatorSet is a free log retrieval operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_RebaseToken *RebaseTokenFilterer) FilterLidoLocatorSet(opts *bind.FilterOpts) (*RebaseTokenLidoLocatorSetIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenLidoLocatorSetIterator{contract: _RebaseToken.contract, event: "LidoLocatorSet", logs: logs, sub: sub}, nil
}

// WatchLidoLocatorSet is a free log subscription operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_RebaseToken *RebaseTokenFilterer) WatchLidoLocatorSet(opts *bind.WatchOpts, sink chan<- *RebaseTokenLidoLocatorSet) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenLidoLocatorSet)
				if err := _RebaseToken.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
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

// ParseLidoLocatorSet is a log parse operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_RebaseToken *RebaseTokenFilterer) ParseLidoLocatorSet(log types.Log) (*RebaseTokenLidoLocatorSet, error) {
	event := new(RebaseTokenLidoLocatorSet)
	if err := _RebaseToken.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the RebaseToken contract.
type RebaseTokenRecoverToVaultIterator struct {
	Event *RebaseTokenRecoverToVault // Event containing the contract specifics and raw log

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
func (it *RebaseTokenRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenRecoverToVault)
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
		it.Event = new(RebaseTokenRecoverToVault)
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
func (it *RebaseTokenRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenRecoverToVault represents a RecoverToVault event raised by the RebaseToken contract.
type RebaseTokenRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*RebaseTokenRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenRecoverToVaultIterator{contract: _RebaseToken.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *RebaseTokenRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenRecoverToVault)
				if err := _RebaseToken.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) ParseRecoverToVault(log types.Log) (*RebaseTokenRecoverToVault, error) {
	event := new(RebaseTokenRecoverToVault)
	if err := _RebaseToken.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the RebaseToken contract.
type RebaseTokenResumedIterator struct {
	Event *RebaseTokenResumed // Event containing the contract specifics and raw log

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
func (it *RebaseTokenResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenResumed)
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
		it.Event = new(RebaseTokenResumed)
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
func (it *RebaseTokenResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenResumed represents a Resumed event raised by the RebaseToken contract.
type RebaseTokenResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_RebaseToken *RebaseTokenFilterer) FilterResumed(opts *bind.FilterOpts) (*RebaseTokenResumedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenResumedIterator{contract: _RebaseToken.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_RebaseToken *RebaseTokenFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *RebaseTokenResumed) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenResumed)
				if err := _RebaseToken.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_RebaseToken *RebaseTokenFilterer) ParseResumed(log types.Log) (*RebaseTokenResumed, error) {
	event := new(RebaseTokenResumed)
	if err := _RebaseToken.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the RebaseToken contract.
type RebaseTokenScriptResultIterator struct {
	Event *RebaseTokenScriptResult // Event containing the contract specifics and raw log

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
func (it *RebaseTokenScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenScriptResult)
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
		it.Event = new(RebaseTokenScriptResult)
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
func (it *RebaseTokenScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenScriptResult represents a ScriptResult event raised by the RebaseToken contract.
type RebaseTokenScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_RebaseToken *RebaseTokenFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*RebaseTokenScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenScriptResultIterator{contract: _RebaseToken.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_RebaseToken *RebaseTokenFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *RebaseTokenScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenScriptResult)
				if err := _RebaseToken.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_RebaseToken *RebaseTokenFilterer) ParseScriptResult(log types.Log) (*RebaseTokenScriptResult, error) {
	event := new(RebaseTokenScriptResult)
	if err := _RebaseToken.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the RebaseToken contract.
type RebaseTokenSharesBurntIterator struct {
	Event *RebaseTokenSharesBurnt // Event containing the contract specifics and raw log

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
func (it *RebaseTokenSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenSharesBurnt)
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
		it.Event = new(RebaseTokenSharesBurnt)
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
func (it *RebaseTokenSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenSharesBurnt represents a SharesBurnt event raised by the RebaseToken contract.
type RebaseTokenSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_RebaseToken *RebaseTokenFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*RebaseTokenSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenSharesBurntIterator{contract: _RebaseToken.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_RebaseToken *RebaseTokenFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *RebaseTokenSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenSharesBurnt)
				if err := _RebaseToken.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_RebaseToken *RebaseTokenFilterer) ParseSharesBurnt(log types.Log) (*RebaseTokenSharesBurnt, error) {
	event := new(RebaseTokenSharesBurnt)
	if err := _RebaseToken.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the RebaseToken contract.
type RebaseTokenStakingLimitRemovedIterator struct {
	Event *RebaseTokenStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *RebaseTokenStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenStakingLimitRemoved)
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
		it.Event = new(RebaseTokenStakingLimitRemoved)
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
func (it *RebaseTokenStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenStakingLimitRemoved represents a StakingLimitRemoved event raised by the RebaseToken contract.
type RebaseTokenStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_RebaseToken *RebaseTokenFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*RebaseTokenStakingLimitRemovedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenStakingLimitRemovedIterator{contract: _RebaseToken.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_RebaseToken *RebaseTokenFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *RebaseTokenStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenStakingLimitRemoved)
				if err := _RebaseToken.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_RebaseToken *RebaseTokenFilterer) ParseStakingLimitRemoved(log types.Log) (*RebaseTokenStakingLimitRemoved, error) {
	event := new(RebaseTokenStakingLimitRemoved)
	if err := _RebaseToken.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the RebaseToken contract.
type RebaseTokenStakingLimitSetIterator struct {
	Event *RebaseTokenStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *RebaseTokenStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenStakingLimitSet)
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
		it.Event = new(RebaseTokenStakingLimitSet)
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
func (it *RebaseTokenStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenStakingLimitSet represents a StakingLimitSet event raised by the RebaseToken contract.
type RebaseTokenStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_RebaseToken *RebaseTokenFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*RebaseTokenStakingLimitSetIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenStakingLimitSetIterator{contract: _RebaseToken.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_RebaseToken *RebaseTokenFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *RebaseTokenStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenStakingLimitSet)
				if err := _RebaseToken.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_RebaseToken *RebaseTokenFilterer) ParseStakingLimitSet(log types.Log) (*RebaseTokenStakingLimitSet, error) {
	event := new(RebaseTokenStakingLimitSet)
	if err := _RebaseToken.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the RebaseToken contract.
type RebaseTokenStakingPausedIterator struct {
	Event *RebaseTokenStakingPaused // Event containing the contract specifics and raw log

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
func (it *RebaseTokenStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenStakingPaused)
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
		it.Event = new(RebaseTokenStakingPaused)
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
func (it *RebaseTokenStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenStakingPaused represents a StakingPaused event raised by the RebaseToken contract.
type RebaseTokenStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_RebaseToken *RebaseTokenFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*RebaseTokenStakingPausedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenStakingPausedIterator{contract: _RebaseToken.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_RebaseToken *RebaseTokenFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *RebaseTokenStakingPaused) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenStakingPaused)
				if err := _RebaseToken.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_RebaseToken *RebaseTokenFilterer) ParseStakingPaused(log types.Log) (*RebaseTokenStakingPaused, error) {
	event := new(RebaseTokenStakingPaused)
	if err := _RebaseToken.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the RebaseToken contract.
type RebaseTokenStakingResumedIterator struct {
	Event *RebaseTokenStakingResumed // Event containing the contract specifics and raw log

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
func (it *RebaseTokenStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenStakingResumed)
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
		it.Event = new(RebaseTokenStakingResumed)
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
func (it *RebaseTokenStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenStakingResumed represents a StakingResumed event raised by the RebaseToken contract.
type RebaseTokenStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_RebaseToken *RebaseTokenFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*RebaseTokenStakingResumedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenStakingResumedIterator{contract: _RebaseToken.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_RebaseToken *RebaseTokenFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *RebaseTokenStakingResumed) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenStakingResumed)
				if err := _RebaseToken.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_RebaseToken *RebaseTokenFilterer) ParseStakingResumed(log types.Log) (*RebaseTokenStakingResumed, error) {
	event := new(RebaseTokenStakingResumed)
	if err := _RebaseToken.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the RebaseToken contract.
type RebaseTokenStoppedIterator struct {
	Event *RebaseTokenStopped // Event containing the contract specifics and raw log

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
func (it *RebaseTokenStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenStopped)
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
		it.Event = new(RebaseTokenStopped)
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
func (it *RebaseTokenStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenStopped represents a Stopped event raised by the RebaseToken contract.
type RebaseTokenStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_RebaseToken *RebaseTokenFilterer) FilterStopped(opts *bind.FilterOpts) (*RebaseTokenStoppedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenStoppedIterator{contract: _RebaseToken.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_RebaseToken *RebaseTokenFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *RebaseTokenStopped) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenStopped)
				if err := _RebaseToken.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_RebaseToken *RebaseTokenFilterer) ParseStopped(log types.Log) (*RebaseTokenStopped, error) {
	event := new(RebaseTokenStopped)
	if err := _RebaseToken.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the RebaseToken contract.
type RebaseTokenSubmittedIterator struct {
	Event *RebaseTokenSubmitted // Event containing the contract specifics and raw log

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
func (it *RebaseTokenSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenSubmitted)
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
		it.Event = new(RebaseTokenSubmitted)
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
func (it *RebaseTokenSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenSubmitted represents a Submitted event raised by the RebaseToken contract.
type RebaseTokenSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_RebaseToken *RebaseTokenFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*RebaseTokenSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenSubmittedIterator{contract: _RebaseToken.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_RebaseToken *RebaseTokenFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *RebaseTokenSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenSubmitted)
				if err := _RebaseToken.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_RebaseToken *RebaseTokenFilterer) ParseSubmitted(log types.Log) (*RebaseTokenSubmitted, error) {
	event := new(RebaseTokenSubmitted)
	if err := _RebaseToken.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenTokenRebasedIterator is returned from FilterTokenRebased and is used to iterate over the raw logs and unpacked data for TokenRebased events raised by the RebaseToken contract.
type RebaseTokenTokenRebasedIterator struct {
	Event *RebaseTokenTokenRebased // Event containing the contract specifics and raw log

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
func (it *RebaseTokenTokenRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenTokenRebased)
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
		it.Event = new(RebaseTokenTokenRebased)
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
func (it *RebaseTokenTokenRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenTokenRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenTokenRebased represents a TokenRebased event raised by the RebaseToken contract.
type RebaseTokenTokenRebased struct {
	ReportTimestamp    *big.Int
	TimeElapsed        *big.Int
	PreTotalShares     *big.Int
	PreTotalEther      *big.Int
	PostTotalShares    *big.Int
	PostTotalEther     *big.Int
	SharesMintedAsFees *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenRebased is a free log retrieval operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_RebaseToken *RebaseTokenFilterer) FilterTokenRebased(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*RebaseTokenTokenRebasedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenTokenRebasedIterator{contract: _RebaseToken.contract, event: "TokenRebased", logs: logs, sub: sub}, nil
}

// WatchTokenRebased is a free log subscription operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_RebaseToken *RebaseTokenFilterer) WatchTokenRebased(opts *bind.WatchOpts, sink chan<- *RebaseTokenTokenRebased, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenTokenRebased)
				if err := _RebaseToken.contract.UnpackLog(event, "TokenRebased", log); err != nil {
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

// ParseTokenRebased is a log parse operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_RebaseToken *RebaseTokenFilterer) ParseTokenRebased(log types.Log) (*RebaseTokenTokenRebased, error) {
	event := new(RebaseTokenTokenRebased)
	if err := _RebaseToken.contract.UnpackLog(event, "TokenRebased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RebaseToken contract.
type RebaseTokenTransferIterator struct {
	Event *RebaseTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RebaseTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenTransfer)
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
		it.Event = new(RebaseTokenTransfer)
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
func (it *RebaseTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenTransfer represents a Transfer event raised by the RebaseToken contract.
type RebaseTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RebaseToken *RebaseTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RebaseTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenTransferIterator{contract: _RebaseToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RebaseToken *RebaseTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RebaseTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenTransfer)
				if err := _RebaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RebaseToken *RebaseTokenFilterer) ParseTransfer(log types.Log) (*RebaseTokenTransfer, error) {
	event := new(RebaseTokenTransfer)
	if err := _RebaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the RebaseToken contract.
type RebaseTokenTransferSharesIterator struct {
	Event *RebaseTokenTransferShares // Event containing the contract specifics and raw log

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
func (it *RebaseTokenTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenTransferShares)
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
		it.Event = new(RebaseTokenTransferShares)
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
func (it *RebaseTokenTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenTransferShares represents a TransferShares event raised by the RebaseToken contract.
type RebaseTokenTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_RebaseToken *RebaseTokenFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RebaseTokenTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RebaseTokenTransferSharesIterator{contract: _RebaseToken.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_RebaseToken *RebaseTokenFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *RebaseTokenTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenTransferShares)
				if err := _RebaseToken.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_RebaseToken *RebaseTokenFilterer) ParseTransferShares(log types.Log) (*RebaseTokenTransferShares, error) {
	event := new(RebaseTokenTransferShares)
	if err := _RebaseToken.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the RebaseToken contract.
type RebaseTokenUnbufferedIterator struct {
	Event *RebaseTokenUnbuffered // Event containing the contract specifics and raw log

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
func (it *RebaseTokenUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenUnbuffered)
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
		it.Event = new(RebaseTokenUnbuffered)
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
func (it *RebaseTokenUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenUnbuffered represents a Unbuffered event raised by the RebaseToken contract.
type RebaseTokenUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*RebaseTokenUnbufferedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenUnbufferedIterator{contract: _RebaseToken.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *RebaseTokenUnbuffered) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenUnbuffered)
				if err := _RebaseToken.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) ParseUnbuffered(log types.Log) (*RebaseTokenUnbuffered, error) {
	event := new(RebaseTokenUnbuffered)
	if err := _RebaseToken.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RebaseTokenWithdrawalsReceivedIterator is returned from FilterWithdrawalsReceived and is used to iterate over the raw logs and unpacked data for WithdrawalsReceived events raised by the RebaseToken contract.
type RebaseTokenWithdrawalsReceivedIterator struct {
	Event *RebaseTokenWithdrawalsReceived // Event containing the contract specifics and raw log

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
func (it *RebaseTokenWithdrawalsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RebaseTokenWithdrawalsReceived)
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
		it.Event = new(RebaseTokenWithdrawalsReceived)
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
func (it *RebaseTokenWithdrawalsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RebaseTokenWithdrawalsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RebaseTokenWithdrawalsReceived represents a WithdrawalsReceived event raised by the RebaseToken contract.
type RebaseTokenWithdrawalsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsReceived is a free log retrieval operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) FilterWithdrawalsReceived(opts *bind.FilterOpts) (*RebaseTokenWithdrawalsReceivedIterator, error) {

	logs, sub, err := _RebaseToken.contract.FilterLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return &RebaseTokenWithdrawalsReceivedIterator{contract: _RebaseToken.contract, event: "WithdrawalsReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsReceived is a free log subscription operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) WatchWithdrawalsReceived(opts *bind.WatchOpts, sink chan<- *RebaseTokenWithdrawalsReceived) (event.Subscription, error) {

	logs, sub, err := _RebaseToken.contract.WatchLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RebaseTokenWithdrawalsReceived)
				if err := _RebaseToken.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
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

// ParseWithdrawalsReceived is a log parse operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_RebaseToken *RebaseTokenFilterer) ParseWithdrawalsReceived(log types.Log) (*RebaseTokenWithdrawalsReceived, error) {
	event := new(RebaseTokenWithdrawalsReceived)
	if err := _RebaseToken.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
