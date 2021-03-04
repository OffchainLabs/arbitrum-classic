module github.com/offchainlabs/arbitrum/packages/arb-validator

go 1.13

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.10.0
	github.com/golang/protobuf v1.4.3
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-runewidth v0.0.6 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-avm-cpp v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-checkpointer v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-util v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.7.3
	github.com/pkg/errors v0.9.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/zerolog v1.20.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-checkpointer => ../arb-checkpointer
