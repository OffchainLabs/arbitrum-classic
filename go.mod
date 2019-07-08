module github.com/offchainlabs/arb-avm

go 1.12

require (
	github.com/dgraph-io/badger v1.6.0
	github.com/ethereum/go-ethereum v1.8.23
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/offchainlabs/arb-util v0.1.1-0.20190620180800-aea2007ae532
)

replace github.com/offchainlabs/arb-util => ../arb-util/
