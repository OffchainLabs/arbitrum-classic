module github.com/offchainlabs/arbitrum/packages/arb-evm

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.21.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
