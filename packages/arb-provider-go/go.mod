module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.20
	github.com/gorilla/rpc v1.2.0
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.7.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.0
	github.com/pkg/errors v0.9.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
