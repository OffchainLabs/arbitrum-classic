module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.10
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/rpc v1.2.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.4.3
	github.com/offchainlabs/arbitrum/packages/arb-avm-go v0.4.3
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.4.3 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.4.3
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.4.3
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core
