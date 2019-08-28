module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.12

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.9.2
	github.com/gorilla/rpc v1.2.0
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.2.0
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.2.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
