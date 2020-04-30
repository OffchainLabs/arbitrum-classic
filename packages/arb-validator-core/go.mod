module github.com/offchainlabs/arbitrum/packages/arb-validator-core

go 1.12

require (
	github.com/ethereum/go-ethereum v1.9.13
	github.com/golang/protobuf v1.3.2
	github.com/offchainlabs/arbitrum/packages/arb-util v0.4.3
	github.com/pkg/errors v0.9.1
	github.com/robertkrimen/otto v0.0.0-20170205013659-6a77b7cbc37d // indirect
	google.golang.org/grpc v1.23.1
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
