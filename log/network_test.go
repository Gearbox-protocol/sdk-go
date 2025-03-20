package log

import "testing"

func TestTestnet(t *testing.T) {
	if ANVIL != GetTestNet(1) {
		t.Error("Test failed")
	}
	if ARBTEST != GetTestNet(42161) {
		t.Error("Test failed")
	}
	if OPTTEST != GetTestNet(10) {
		t.Error("Test failed")
	}
	if SONICTEST != GetTestNet(146) {
		t.Error("Test failed")
	}
}

func TestBaseNet(t *testing.T) {
	if MAINNET != GetBaseNet(7878) {
		t.Error("Base failed")
	}
	if MAINNET != GetBaseNet(1337) {
		t.Error("Base failed")
	}
	if ARBITRUM != GetBaseNet(7880) {
		t.Error("Base failed")
	}
	if OPTIMISM != GetBaseNet(7879) {
		t.Error("Base failed")
	}
	if SONIC != GetBaseNet(7882) {
		t.Error("Base failed")
	}
}
func TestNet(t *testing.T) {
	if ANVIL != GetNetworkName(7878) {
		t.Error("net name failed")
	}
	if MAINNET != GetNetworkName(1) {
		t.Error("net name failed")
	}
	if ARBTEST != GetNetworkName(7880) {
		t.Error("net name failed")
	}
	if OPTTEST != GetNetworkName(7879) {
		t.Error("net name failed")
	}
	if SONICTEST != GetNetworkName(7882) {
		t.Error("net name failed")
	}
}
