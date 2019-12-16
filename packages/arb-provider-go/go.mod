module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.8
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/rpc v1.2.0
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.2.0
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.2.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
