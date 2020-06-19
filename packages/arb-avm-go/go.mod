module github.com/offchainlabs/arbitrum/packages/arb-avm-go

go 1.12

require (
	github.com/dgraph-io/badger v1.6.1
	github.com/ethereum/go-ethereum v1.9.13
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.5
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
