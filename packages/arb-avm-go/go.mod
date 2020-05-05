module github.com/offchainlabs/arbitrum/packages/arb-avm-go

go 1.12

require (
	github.com/dgraph-io/badger v1.6.1
	github.com/ethereum/go-ethereum v1.9.13
	github.com/offchainlabs/arbitrum/packages/arb-util v0.5.0
	github.com/robertkrimen/otto v0.0.0-20170205013659-6a77b7cbc37d // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
