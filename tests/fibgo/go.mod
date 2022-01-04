module github.com/offchainlabs/arbitrum/tests/fibgo

go 1.13

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/consensys/gurvy v0.3.8 // indirect
	github.com/ethereum/go-ethereum v1.10.10
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/knadh/koanf v1.3.2 // indirect
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.7.3
	github.com/offchainlabs/go-solidity-sha3 v0.1.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/zerolog v1.26.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../../packages/arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../../packages/arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../../packages/arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../../packages/arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../../packages/arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator => ../../packages/arb-tx-aggregator
