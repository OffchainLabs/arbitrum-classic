module github.com/offchainlabs/arbitrum/packages/arb-util

go 1.13

require (
	github.com/ethereum/go-ethereum v1.10.1
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f
	github.com/offchainlabs/arbitrum/packages/arb-node-core v0.8.0
	github.com/offchainlabs/go-solidity-sha3 v0.1.2
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.21.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace github.com/offchainlabs/arbitrum/packages/arb-node-core => ../arb-node-core

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../arb-avm-cpp
