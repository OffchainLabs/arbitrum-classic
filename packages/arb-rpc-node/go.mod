module github.com/offchainlabs/arbitrum/packages/arb-rpc-node

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.0
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-runewidth v0.0.6 // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.0.0-20200123000308-a60dcd172b4c
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-evm v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-node-core v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.20.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-node-core => ../arb-node-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
