# SDK

Components:
- `logging`: a custom wrapper over log package which supports different debugging levels INFO/ERRROR/WARN/FATAL and sending fatal/error messages to configured amqp server.
- `ethclient`: wrapper over geth ethclient library, with support for error handling for infura/alchemy RPCs, parallel request processing using multiple clients internally. 
- `core`: Custom data structure for storing data in db using gorm(bigint/hstore/json/transfer). Other module for endpoints over eth node for getting base gas/making multicalls/getting latest block number etc.
- `core/schemas`: third-eye db schemas. 
- `utils`: utilities for time/string/bigint.
- `artifacts`: go files generated using abigen from gearbox contracts.
