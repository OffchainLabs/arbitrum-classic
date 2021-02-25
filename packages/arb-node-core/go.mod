module github.com/offchainlabs/arbitrum/packages/arb-node-core

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.24
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.7.3 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
