module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.23
	github.com/golang/protobuf v1.4.2
	github.com/kr/pretty v0.2.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.7.2
	github.com/offchainlabs/arbitrum/packages/arb-checkpointer v0.7.2
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.7.1 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator v0.7.1 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.2
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.7.2
	github.com/c-bata/go-prompt v0.2.3 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/term v0.0.0-20200520122047-c3ffed290a03 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-checkpointer => ../arb-checkpointer
