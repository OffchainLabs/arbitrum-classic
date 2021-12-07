module github.com/offchainlabs/arbitrum/packages/arb-evm

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.13-0.20211112145008-abc74a5ffeb7
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.26.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-node-core => ../arb-node-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
