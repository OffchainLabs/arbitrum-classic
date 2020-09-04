module github.com/offchainlabs/arbitrum/tests/fibgo

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.20
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.7.1
	github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator v0.7.1
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.1
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.7.1
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.7.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-provider-go => ../../packages/arb-provider-go

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../../packages/arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../../packages/arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../../packages/arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../../packages/arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../../packages/arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-checkpointer => ../../packages/arb-checkpointer

replace github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator => ../../packages/arb-tx-aggregator
