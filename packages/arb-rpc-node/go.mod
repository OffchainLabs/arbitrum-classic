module github.com/offchainlabs/arbitrum/packages/arb-rpc-node

go 1.13

require (
	github.com/btcsuite/btcd v0.22.1 // indirect
	github.com/c-bata/go-prompt v0.2.6
	github.com/ethereum/go-ethereum v1.10.18
	github.com/ethersphere/bee v1.6.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40 // indirect
	github.com/knadh/koanf v1.4.0 // indirect
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-node-core v0.8.0
	github.com/offchainlabs/arbitrum/packages/arb-util v0.8.0
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/rs/zerolog v1.26.1
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-node-core => ../arb-node-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/ethereum/go-ethereum => ../go-ethereum
