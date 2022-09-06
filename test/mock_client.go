package test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TestClient struct {
	// Blocks map[int64]BlockInput
	blockNums []int64
	// blockNum
	events   map[int64]map[string][]types.Log
	prices   map[string]map[int64]*big.Int
	masks    map[int64]map[string]*big.Int
	state    *StateManager
	USDCAddr string
	WETHAddr string
	token    map[string]*schemas.Token
}

func NewTestClient() *TestClient {
	return &TestClient{
		events: make(map[int64]map[string][]types.Log),
		token:  make(map[string]*schemas.Token),
		state:  NewStateManager(),
	}
}

func (t *TestClient) ChainID(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1337), nil
}
func (t *TestClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return types.NewBlock(&types.Header{Time: uint64(number.Int64()) * 86400},
		[]*types.Transaction{},
		[]*types.Header{},
		[]*types.Receipt{}, nil), nil
}
func (t *TestClient) BlockNumber(ctx context.Context) (uint64, error) {
	if len(t.blockNums) == 0 {
		return 1, nil
	}
	return uint64(t.blockNums[len(t.blockNums)-1]), nil
}
func topic(v string) common.Hash {
	return crypto.Keccak256Hash([]byte(v))
}
func ContainsHash(list []common.Hash, v common.Hash) bool {
	for _, hash := range list {
		if hash == v {
			return true
		}
	}
	return false
}
func (t *TestClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	toBlock := query.ToBlock.Int64()
	txLogs := []types.Log{}
	for i := query.FromBlock.Int64(); i <= toBlock; i++ {
		for _, address := range query.Addresses {
			if t.events[i] != nil {
				if len(query.Topics) > 0 { // if topic is present  for transfer
					if query.Topics[0][0] == topic("Transfer(address,address,uint256)") {
						for _, txLog := range t.events[i][address.Hex()] {
							if ContainsHash(query.Topics[2], txLog.Topics[2]) {
								txLogs = append(txLogs, txLog)
							}
						}
					} else { // for other tokens
						for _, txLog := range t.events[i][address.Hex()] {
							if ContainsHash(query.Topics[0], txLog.Topics[0]) {
								txLogs = append(txLogs, txLog)
							}
						}
					}
				} else { // no topic
					txLogs = append(txLogs, t.events[i][address.Hex()]...)
				}
			}
		}
	}
	txLogList := TxLogList(txLogs)
	sort.Sort(txLogList)
	return ([]types.Log)(txLogList), nil
}

func (t *TestClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}
func (t *TestClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (t *TestClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

// for otherCalls in call of blocks
//
func addrAndData(input []string) (addr string, data []string) {
	isAddr := false
	for len(input) > 0 {
		entry := input[len(input)-1]
		if entry == "" {
			isAddr = true
		} else if isAddr {
			addr = entry
		} else {
			data = append(data, entry)
		}
		input = input[:len(input)-1]
	}
	return
}

func (t *TestClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	sig := hex.EncodeToString(call.Data[:4])
	var blockNum int64
	if blockNumber != nil {
		blockNum = blockNumber.Int64()
	}
	if data := t.state.GetOtherCall(blockNum, sig, *call.To); data != "" {
		return common.HexToHash(data).Bytes(), nil
	}
	if sig == "95d89b41" { // symbol
		sym := t.token[call.To.Hex()].Symbol
		zeroBytes := make([]byte, 96)
		zeroBytes[31] = 32
		zeroBytes[63] = byte(len(sym))
		for i, r := range sym {
			zeroBytes[64+i] = byte(r)
		}
		return zeroBytes, nil
	} else if sig == "313ce567" { // decimals
		decimals := t.token[call.To.Hex()].Decimals
		zeroBytes := make([]byte, 32)
		zeroBytes[31] = byte(decimals)
		return zeroBytes, nil
	} else if sig == "d6d19b27" { // convert with credit account on priceoraclev2
		// only skip the creditaccount bytes, not for the sig
		// as convertPrice will itself skip 4 bytes for sig and the value of these first 4 bytes of data doesn't matter
		return common.HexToHash(fmt.Sprintf("%x", t.convertPrice(blockNum, call.Data[32:]))).Bytes(), nil
		// getprice on priceoracle with nft support taking credit account is arg
	} else if sig == "ac41865a" {
		// credit account and then token
		tokenAddr := common.BytesToAddress(call.Data[4+32:]).Hex()
		return common.HexToHash(fmt.Sprintf("%x", t.getPrice(blockNum, tokenAddr))).Bytes(), nil
	} else if sig == "41976e09" {
		tokenAddr := common.BytesToAddress(call.Data[4:]).Hex()
		return common.HexToHash(fmt.Sprintf("%x", t.getPrice(blockNum, tokenAddr))).Bytes(), nil
		// convert on priceOracle
	} else if sig == "b66102df" {
		return common.HexToHash(fmt.Sprintf("%x", t.convertPrice(blockNum, call.Data))).Bytes(), nil

		// enabledmask on creditfilter for account
		// v1 or v2 get mask
	} else if sig == "b451cecc" || sig == "8991b2f1" {
		s := 4
		account := common.BytesToAddress(call.Data[s : s+32]).Hex()
		mask := t.masks[blockNum][account]
		return common.HexToHash(fmt.Sprintf("%x", mask)).Bytes(), nil
		// phaseId
	} else if sig == "58303b10" {
		oracle := *call.To
		index := t.state.OracleMgr.GetIndex(oracle, blockNum)
		return common.HexToHash(fmt.Sprintf("%x", index)).Bytes(), nil
		// current phase aggregator
	} else if sig == "c1597304" {
		s := 4
		index, ok := new(big.Int).SetString(hex.EncodeToString(call.Data[s:s+32]), 16)
		if !ok {
			log.Fatal("oracle:%s data: %s", call.To, call.Data)
		}
		oracle := *call.To
		feed := t.state.OracleMgr.GetState(oracle, int(index.Int64()))
		return common.HexToHash(feed.Feed).Bytes(), nil
	} else if sig == "bce38bd7" { // multicall sig
		obj := map[string]interface{}{}
		parser := core.GetAbi("MultiCall")
		method, err := parser.MethodById(call.Data[:4])
		log.CheckFatal(err)
		method.Inputs.UnpackIntoMap(obj, call.Data[4:])
		calls := *abi.ConvertType(obj["calls"], new([]multicall.Multicall2Call)).(*[]multicall.Multicall2Call)
		resultArray := []multicall.Multicall2Result{}
		for _, call := range calls {
			internalCallSig := hex.EncodeToString(call.CallData[:4])
			switch internalCallSig {
			case "f875365d": // observer token
				resultArray = append(resultArray, multicall.Multicall2Result{
					Success: false,
				})
			case "b66102df": // convert on priceoracle v1
				price := t.convertPrice(blockNum, call.CallData)
				resultArray = append(resultArray, multicall.Multicall2Result{
					Success:    true,
					ReturnData: common.HexToHash(fmt.Sprintf("%x", price)).Bytes(),
				})
			case "ac41865a": // get price  on priceoracle v2 with credit account
				tokenAddr := common.BytesToAddress(call.CallData[4+32:]).Hex()
				price := t.getPrice(blockNum, tokenAddr)
				resultArray = append(resultArray, multicall.Multicall2Result{
					Success:    true,
					ReturnData: common.HexToHash(fmt.Sprintf("%x", price)).Bytes(),
				})
			case "41976e09": // getprice  on priceoracle v2
				tokenAddr := common.BytesToAddress(call.CallData[4:]).Hex()
				price := t.getPrice(blockNum, tokenAddr)
				resultArray = append(resultArray, multicall.Multicall2Result{
					Success:    true,
					ReturnData: common.HexToHash(fmt.Sprintf("%x", price)).Bytes(),
				})
			case "f93f515b", // creditFilter
				"f9aa028a", //creditConfigurator
				"570a7af2", //poolService
				"2495a599", //underlyingToken
				"6f307dc3", //underlying
				"2f7a1881", //creditFacade
				"313ce567", // decimals
				"70a08231", // balanceOf
				"feaf968c", // latestRoundData
				"95d89b41": // symbol
				if data := t.state.GetOtherCall(blockNum, internalCallSig, call.Target); data == "" {
					resultArray = append(resultArray, multicall.Multicall2Result{
						Success:    false,
						ReturnData: []byte{},
					})
				} else {
					splits := strings.Split(data, ":")
					var tmpBytes []byte
					if len(splits) == 2 {
						switch splits[0] {
						case "bigint":
							arg, ok := new(big.Int).SetString(splits[1], 10)
							if !ok {
								log.Fatalf("bigint parsing failed for %s", data)
							}
							tmpBytes = common.BytesToHash(arg.Bytes()).Bytes()
						}
					} else {
						tmpBytes = common.HexToHash(data).Bytes()
					}
					var ansBytes []byte
					switch internalCallSig {
					case "feaf968c": // latestRoundData
						ansBytes = make([]byte, 32*5)
						copy(ansBytes[32:], tmpBytes)
					case "95d89b41": // symbol
						zeroBytes := make([]byte, 96)
						zeroBytes[31] = 32
						zeroBytes[63] = byte(len(data))
						for i, r := range data {
							zeroBytes[64+i] = byte(r)
						}
						ansBytes = zeroBytes
					default:
						ansBytes = tmpBytes
					}
					resultArray = append(resultArray, multicall.Multicall2Result{
						Success:    true,
						ReturnData: ansBytes,
					})
				}
			default:
				panic(fmt.Sprintf("sig %s and target %s", hex.EncodeToString(call.CallData[:4]), call.Target))
			}
		}
		outputData, err := method.Outputs.Pack(resultArray)
		log.CheckFatal(err)
		return outputData, nil
	}
	return nil, nil
}

func (t *TestClient) convertPrice(blockNum int64, data []byte) *big.Int {
	s := 4
	amount, ok := new(big.Int).SetString(hex.EncodeToString(data[s:s+32]), 16)
	if !ok {
		log.Fatal("failed in parsing int")
	}
	s += 32
	token0 := common.BytesToAddress(data[s : s+32]).Hex()
	decimalT0 := t.token[token0].Decimals
	s += 32
	token1 := common.BytesToAddress(data[s : s+32]).Hex()
	decimalT1 := t.token[token1].Decimals
	price0 := t.getPrice(blockNum, token0)
	price1 := t.getPrice(blockNum, token1)
	newAmount := new(big.Int).Mul(amount, price0)
	newAmount = utils.GetInt64(newAmount, decimalT0-decimalT1)
	newAmount = new(big.Int).Quo(newAmount, price1)
	return newAmount
}
func (t *TestClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return nil, nil
}
func (t *TestClient) getPrice(blockNum int64, tokenAddr string) *big.Int {
	if t.prices[tokenAddr] != nil {
		var ansBlockNum int64
		for currentNum := range t.prices[tokenAddr] {
			if currentNum <= blockNum && ansBlockNum < currentNum {
				ansBlockNum = currentNum
			}
		}
		return t.prices[tokenAddr][ansBlockNum]
	} else if tokenAddr == t.WETHAddr { // only for v1
		value, _ := new(big.Int).SetString("1000000000000000000", 10)
		return value
	} else {
		panic(fmt.Sprintf("token(%s) price not present", tokenAddr))
	}
}
func (t *TestClient) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	return nil, nil
}
func (t *TestClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return nil, nil
}
func (t *TestClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}
func (t *TestClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (t *TestClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (t *TestClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 0, nil
}
func (t *TestClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (t *TestClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}
