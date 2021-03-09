module github.com/offchainlabs/arbitrum/packages/arb-validator-core

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.1
	github.com/golang/protobuf v1.4.3
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/protobuf v1.25.0
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
