module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.8
	github.com/gorilla/rpc v1.2.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.4.0
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.4.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
