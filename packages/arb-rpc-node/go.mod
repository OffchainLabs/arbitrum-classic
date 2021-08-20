module github.com/offchainlabs/arbitrum/packages/arb-rpc-node

go 1.13

require (
	github.com/c-bata/go-prompt v0.2.2
	github.com/ethereum/go-ethereum v1.10.4
	github.com/ethersphere/bee v0.6.2
	github.com/go-redis/redis/v8 v8.10.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/miguelmota/go-ethereum-hdwallet v0.1.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-node-core v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/rs/zerolog v1.23.0
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gotest.tools/gotestsum v1.7.0 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-node-core => ../arb-node-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
