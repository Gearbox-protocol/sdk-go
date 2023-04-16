package ethclient

import (
	"fmt"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient(t *testing.T) {
	t.Skip("skip")
	log.SetTestLogging(t)
	client, err := Dial("")
	log.CheckFatal(err)
	ans, err := getDataViaRetry(client, func(c *ethclient.Client) (int, error) {
		return 0, fmt.Errorf("as")
	})
	log.Info(ans, err)
	ans, err = getDataViaRetry(client, func(c *ethclient.Client) (int, error) {
		return 0, fmt.Errorf("as")
	})
	log.Info(ans, err)
}
