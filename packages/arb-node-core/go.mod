module github.com/offchainlabs/arbitrum/packages/arb-node-core

go 1.13

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.10.1
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-rpc-node v0.8.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/zerolog v1.21.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
