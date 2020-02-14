module github.com/offchainlabs/arbitrum/packages/arb-avm-go

go 1.12

require (
	github.com/dgraph-io/badger v1.6.0
	github.com/ethereum/go-ethereum v1.9.8
	github.com/offchainlabs/arbitrum/packages/arb-util v0.4.2
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
