module github.com/offchainlabs/arbitrum/packages/arb-avm-cpp

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.24
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/pkg/errors v0.9.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
