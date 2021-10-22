module github.com/offchainlabs/arbitrum/tests/fibgo

go 1.13

require (
	github.com/VictoriaMetrics/fastcache v1.6.0 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v1.1.1 // indirect
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/cloudflare/cloudflare-go v0.14.0 // indirect
	github.com/consensys/gnark-crypto v0.4.1-0.20210426202927-39ac3d4b3f1f // indirect
	github.com/consensys/gurvy v0.3.8 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.9.25
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4 // indirect
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.5 // indirect
	github.com/graph-gophers/graphql-go v0.0.0-20201113091052-beb923fada29 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.0 // indirect
	github.com/huin/goupnp v1.0.2 // indirect
	github.com/influxdata/influxdb v1.8.3 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.4.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097 // indirect
	github.com/knadh/koanf v1.2.2 // indirect
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f // indirect
	github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-validator v0.7.3
	github.com/offchainlabs/arbitrum/packages/arb-validator-core v0.7.3
	github.com/offchainlabs/go-solidity-sha3 v0.1.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/rs/zerolog v1.24.0
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954 // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/sys v0.0.0-20210816183151-1e6c022a8912 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-validator => ../../packages/arb-validator

replace github.com/offchainlabs/arbitrum/packages/arb-validator-core => ../../packages/arb-validator-core

replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp => ../../packages/arb-avm-cpp

replace github.com/offchainlabs/arbitrum/packages/arb-evm => ../../packages/arb-evm

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../../packages/arb-util

replace github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator => ../../packages/arb-tx-aggregator
