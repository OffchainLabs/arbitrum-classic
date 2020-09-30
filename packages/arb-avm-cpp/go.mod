module github.com/offchainlabs/arbitrum/packages/arb-avm-cpp

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.22
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
