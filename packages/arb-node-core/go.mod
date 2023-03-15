module github.com/offchainlabs/arbitrum/packages/arb-node-core

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.18
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/rs/zerolog v1.26.1
	golang.org/x/crypto v0.1.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/ethereum/go-ethereum => ../go-ethereum
