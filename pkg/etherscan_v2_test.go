package pkg

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func TestGetEtherscanLogs(t *testing.T) {
	addr := []common.Address{common.HexToAddress("0x05eF8Da767687c0137aE143EB23883FcaD235ce7")}
	tops := [][]common.Hash{{
		core.Topic("SetQuotaIncreaseFee(address,uint16)"),
		core.Topic("SetTokenLimit(address,uint96)"),
		core.Topic("UpdateTokenQuotaRate(address,uint16)"),
		core.Topic("AddQuotaToken(address)"),
	}}
	var toBlock int64 = 22550825
	scanlogs, err := core.GetEtherscanLogs(1, addr, toBlock,
		tops)

	client, err := ethclient.Dial(utils.GetEnvOrDefault("ETH_PROVIDER", ""))
	log.CheckFatal(err)
	txLogs, err := Node{Client: client}.GetLogs(0, toBlock, addr, tops)
	log.CheckFatal(err)
	if len(scanlogs) != len(txLogs) {
		t.Errorf("Expected %d logs, got %d", len(txLogs), len(scanlogs))
	} else {
		t.Logf("Got %d logs", len(scanlogs))
	}
}
