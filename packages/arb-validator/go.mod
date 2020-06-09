module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.13
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/rpc v1.2.0
	github.com/hashicorp/golang-lru v0.0.0-20160813221303-0a025b7e63ad
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.6.4
	github.com/offchainlabs/arbitrum/packages/arb-avm-go v0.6.4
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.4
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.6.4
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.24.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core
