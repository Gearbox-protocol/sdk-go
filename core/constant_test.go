package core

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestGetMarketConfigurators(t *testing.T) {
	// Test with environment variable set - this tests the core parsing logic
	t.Run("WithEnvironmentVariable", func(t *testing.T) {

		result := GetMarketConfigurators(1) // Mainnet

		// Should have at least the addresses from environment variable
		assert.True(t, len(result) >= 2)

		expected := []common.Address{
			common.HexToAddress("0x354fe9f450F60b8547f88BE042E4A45b46128a06"),
			common.HexToAddress("0x4d427D418342d8CE89a7634c3a402851978B680A"),
			common.HexToAddress("0x3b56538833fc02f4f0e75609390f26ded0c32e42"),
			common.HexToAddress("0xc168343C791D56dD1Da4b4B8B0cc1C1EC1A16E6B"),
			common.HexToAddress("0x7a133fbd01736fd076158307c9476cc3877f1af5"),
			common.HexToAddress("0x09d8305F49374AEA6A78aF6C996df2913e8f3b19"),
		}

		// Check that the first two addresses match our environment variable
		for ind, exp := range expected {
			assert.Equal(t, exp, result[ind])
		}
	})
	// Test with single address in environment variable
}
