module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.14
	github.com/gorilla/rpc v1.2.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.5.0
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.5.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
