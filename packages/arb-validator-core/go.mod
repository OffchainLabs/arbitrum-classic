module github.com/offchainlabs/arbitrum/packages/arb-validator-core

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.10
	github.com/golang/protobuf v1.3.2
	github.com/offchainlabs/arbitrum/packages/arb-util v0.4.3
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.23.1
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
