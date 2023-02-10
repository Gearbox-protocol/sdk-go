package core

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/utils"
)

func TestGetSnapshotsInRange(t *testing.T) {
	ans := MainnetPoolRewards.getSnapshotsInRange(0, 15977000)
	if len(ans) != 2 {
		t.Fatal(utils.ToJson(ans))
	}
	ans = MainnetPoolRewards.getSnapshotsInRange(0, 15820000)
	if len(ans) != 1 {
		t.Fatal(utils.ToJson(ans))
	}
	ans = MainnetPoolRewards.getSnapshotsInRange(0, 15819999)
	if len(ans) != 0 {
		t.Fatal(utils.ToJson(ans))
	}
	ans = MainnetPoolRewards.getSnapshotsInRange(15977000, 15977010)
	if len(ans) != 1 {
		t.Fatal(utils.ToJson(ans))
	}
}
