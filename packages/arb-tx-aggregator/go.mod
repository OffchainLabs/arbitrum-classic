module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.13
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/rpc v1.2.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.4
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.6.4
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core
