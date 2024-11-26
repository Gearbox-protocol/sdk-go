package core

// https://github.com/Gearbox-protocol/integrations-v2/tree/faa9cfd4921c62165782dcdc196ff5a0c0e6075d/contracts/oracles
// https://github.com/Gearbox-protocol/oracles-v3/tree/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles
const (
	V3_CHAINLINK_ORACLE = iota
	V3_YEARN_ORACLE
	V3_CURVE_2LP_ORACLE
	V3_CURVE_3LP_ORACLE
	V3_CURVE_4LP_ORACLE
	V3_ZERO_ORACLE
	V3_WSTETH_ORACLE
	V3_BOUNDED_ORACLE
	V3_COMPOSITE_ORACLE
	V3_WRAPPED_AAVE_V2_ORACLE
	V3_COMPOUND_V2_ORACLE
	V3_BALANCER_STABLE_LP_ORACLE
	V3_BALANCER_WEIGHTED_LP_ORACLE
	V3_CURVE_CRYPTO_ORACLE
	V3_THE_SAME_AS
	V3_REDSTONE_ORACLE
	V3_ERC4626_VAULT_ORACLE
	V3_NETWORK_DEPENDENT
	V3_CURVE_USD_ORACLE
	V3_PYTH_ORACLE
	V3_MELLOW_LRT_ORACLE
	V3_PENDLE_PT_TWAP_ORACLE
	//
	V3_BACKEND_COMPOSITE_REDSTONE_ORACLE = 100
	V3_BACKEND_GENERAL_ORACLE            = 101
	V3_PULL_UNDERLYING_ORACLE            = 102
)
