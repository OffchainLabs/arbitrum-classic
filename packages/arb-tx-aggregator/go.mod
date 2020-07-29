module github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.17
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/rpc v1.2.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.6.5
	github.com/offchainlabs/arbitrum/packages/arb-checkpointer v0.0.0-00010101000000-000000000000
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.0.0-00010101000000-000000000000
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.6.5
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.5
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.6.5
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-provider-go => ../arb-provider-go

replace github.com/offchainlabs/arbitrum/packages/arb-checkpointer => ../arb-checkpointer
