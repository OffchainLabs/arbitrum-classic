module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.12

require (
	github.com/aristanetworks/goarista v0.0.0-20190912214011-b54698eaaca6 // indirect
	github.com/elastic/gosigar v0.10.5 // indirect
	github.com/ethereum/go-ethereum v1.9.2
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/karalabe/usb v0.0.0-20190819132248-550797b1cad8 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.2.0
	github.com/offchainlabs/arbitrum/packages/arb-avm-go v0.2.0
	github.com/offchainlabs/arbitrum/packages/arb-provider-go v0.2.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.2.0
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.7.0 // indirect
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
