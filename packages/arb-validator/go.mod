module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.14
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/rpc v1.2.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.5.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-go v0.5.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.5.0
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.5.0
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
	google.golang.org/protobuf v1.22.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core
