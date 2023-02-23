package priceFetcher

import (
	"os"
	"sync"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSpotPriceStore(t *testing.T) {
	if err := os.Chdir("../../"); err != nil {
		t.Fatal(err)
	}
	client, err := ethclient.Dial(os.Getenv("ETH_PROVIDER"))
	if err != nil {
		t.Fatal(err)
	}
	//
	var blockNumber int64 = 16112061
	//
	expectedTokenPrices := map[common.Address]string{}
	utils.ReadJsonAndSetInterface("inputs/spot_price_test.json", &expectedTokenPrices)
	tStore := getDecimalStore(client, expectedTokenPrices, blockNumber, t)
	t.Log(utils.ToJson(tStore.decimals))
	inchOracle := common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
	//
	store := New1InchOracle(client, 1, inchOracle, tStore)
	calls := store.GetCalls()
	//
	results := core.MakeMultiCall(client, blockNumber, false, calls)
	require.JSONEq(t, utils.ToJson(expectedTokenPrices), utils.ToJson(store.GetPrices(results, blockNumber)))
}

func getDecimalStore(client core.ClientI, tokenPrices map[common.Address]string, blockNum int64, t *testing.T) *TokensStore {
	erc20ABI := core.GetAbi("Token")
	decimalsData, err := erc20ABI.Pack("decimals")
	if err != nil {
		t.Fatal("Crv", err)
	}

	calls := make([]multicall.Multicall2Call, 0, len(tokenPrices))
	tokenInOrder := make([]common.Address, 0, len(tokenPrices))
	for token := range tokenPrices {
		calls = append(calls, multicall.Multicall2Call{
			Target:   token,
			CallData: decimalsData,
		})
		tokenInOrder = append(tokenInOrder, token)
	}

	tStore := map[common.Address]schemas.Token{}

	results := core.MakeMultiCall(client, blockNum, false, calls)
	for ind, entry := range results {
		token := tokenInOrder[ind]
		if entry.Success {
			values, err := erc20ABI.Unpack("decimals", entry.ReturnData)
			if err != nil {
				t.Fatal(err)
			}
			tStore[token] = schemas.Token{
				Decimals: int8(values[0].(uint8)),
				Address:  token.Hex(),
			}
		}
	}

	return &TokensStore{decimals: tStore, mu: &sync.RWMutex{}}
}
