module github.com/offchainlabs/arbitrum/packages/arb-test

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.2
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.0
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-go v0.3.0
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.0.0-00010101000000-000000000000
	github.com/offchainlabs/arbitrum/packages/arb-util v0.3.0
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.2.0
	github.com/offchainlabs/poker-dapp v0.0.0-20190814222605-11eab5eaf9ff
	github.com/pkg/errors v0.8.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-provider-go => ../arb-provider-go

replace github.com/offchainlabs/arbitrum/packages/arb-test => ../arb-test
