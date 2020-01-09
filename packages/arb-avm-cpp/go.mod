module github.com/offchainlabs/arbitrum/packages/arb-avm-cpp

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.8
	github.com/offchainlabs/arbitrum/packages/arb-util v0.2.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
