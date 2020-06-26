module github.com/offchainlabs/arbitrum/packages/arb-validator-core

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.15
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.5
	github.com/pkg/errors v0.9.1
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
	google.golang.org/protobuf v1.25.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
