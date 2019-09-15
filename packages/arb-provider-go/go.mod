module github.com/offchainlabs/arbitrum/packages/arb-provider-go

go 1.12

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.9.2
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.0
	github.com/gorilla/rpc v1.2.0
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.2.0
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.2.0
	github.com/olekukonko/tablewriter v0.0.1 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48 // indirect
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/wsddn/go-ecdh v0.0.0-20161211032359-48726bab9208 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-avm-go => ../arb-avm-go

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
