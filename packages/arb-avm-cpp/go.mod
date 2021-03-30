module github.com/offchainlabs/arbitrum/packages/arb-avm-cpp

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.1
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.21.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
