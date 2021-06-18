package configuration

import (
	"context"
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/mitchellh/mapstructure"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
	"math/big"
	"os"
	"path"
	"time"
)

var logger = log.With().Caller().Stack().Str("component", "configuration").Logger()

type Bridge struct {
	Utils struct {
		Address string `koanf:"address"`
	} `koanf:"utils"`
}

type Feed struct {
	Input struct {
		URL string `koanf:"url"`
	} `koanf:"input"`
	Output struct {
		Addr    string `koanf:"addr"`
		Port    string `koanf:"port"`
		Ping    string `koanf:"ping"`
		Timeout string `koanf:"timeout"`
	} `koanf:"output"`
}

type Healthcheck struct {
	Addr   string `koanf:"addr"`
	Enable bool   `koanf:"enable"`
	L1Node struct {
		Enable bool `koanf:"enable"`
	} `koanf:"l1-node"`
	Metrics struct {
		Enable bool   `koanf:"enable"`
		Prefix string `koanf:"prefix"`
	} `koanf:"metrics"`
	Port      string `koanf:"port"`
	Sequencer struct {
		Enable bool `koanf:"enable"`
	} `koanf:"sequencer"`
}

type Node struct {
	Aggregator struct {
		Inbox struct {
			Address string `koanf:"address"`
		} `koanf:"inbox"`
		MaxBatchTime int64 `koanf:"max-batch-time"`
		Stateful     bool  `koanf:"stateful"`
	} `koanf:"aggregator"`
	Forward struct {
		URL string `koanf:"url"`
	} `koanf:"forward"`
	Sequencer struct {
		CreateBatchBlockInterval   int64 `koanf:"create-batch-block-interval"`
		DelayedMessagesTargetDelay int64 `koanf:"delayed-messages-target-delay"`
		Enable                     bool  `koanf:"enable"`
	} `koanf:"sequencer"`
}

type Persistent struct {
	Database struct {
		Path string `koanf:"path"`
	} `koanf:"database"`
	Storage struct {
		Path string `koanf:"path"`
	} `koanf:"storage"`
}

type Rollup struct {
	Address   string `koanf:"address"`
	ChainID   uint64 `koanf:"chain-id"`
	FromBlock int64  `koanf:"from-block"`
	Machine   struct {
		Filename string `koanf:"filename"`
		//URL string `koanf:"url"`
	} `koanf:"machine"`
}

type Validator struct {
	Strategy string `koanf:"strategy"`
	Utils    struct {
		Address string `koanf:"address"`
	} `koanf:"utils"`
	WalletFactory struct {
		Address string `koanf:"address"`
	} `koanf:"wallet-factory"`
}

type Wallet struct {
	Password string  `koanf:"password"`
	GasPrice float64 `koanf:"gas-price"`
}

type Config struct {
	Bridge Bridge `koanf:"bridge"`
	Conf   string `koanf:"conf"`
	Dump   struct {
		Conf bool `koanf:"conf"`
	} `koanf:"dump"`
	Feed        Feed        `koanf:"feed"`
	GasPriceUrl string      `koanf:"gas-price-url"`
	Healthcheck Healthcheck `koanf:"healthcheck"`
	L1          struct {
		URL string `koanf:"url"`
	} `koanf:"l1"`
	Log struct {
		RPC  string `koanf:"rpc"`
		Core string `koanf:"core"`
	} `koanf:"log"`
	Node       Node       `kaonf:"sequencer"`
	Persistent Persistent `koanf:"persistent"`
	PProf      struct {
		Enable bool `koanf:"enable"`
	} `koanf:"pprof"`
	Rollup Rollup `koanf:"rollup"`
	RPC    struct {
		Addr string `koanf:"addr"`
		Port string `koanf:"port"`
	} `koanf:"rpc"`
	Validator     Validator `koanf:"validator"`
	WaitToCatchUp bool      `koanf:"wait-to-catch-up"`
	Wallet        Wallet    `koanf:"wallet"`
	WS            struct {
		Addr string `koanf:"addr"`
		Port string `koanf:"port"`
	} `koanf:"ws"`
}

func Parse(ctx context.Context) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f := flag.NewFlagSet("config", flag.ContinueOnError)

	f.String("bridge.utils.address", "", "bridgeutils contract address")

	f.String("conf", "", "name of configuration file")

	f.Bool("dump.conf", false, "print out currently active configuration file")

	f.String("gas-price-url", "", "gas price rpc url (etherscan compatible)")

	f.String("node.aggregator.inbox.address", "", "address of the inbox contract")
	f.Int64("node.aggregator.max-batch-time", 10, "maxBatchTime=NumSeconds")
	f.Bool("node.aggregator.stateful", false, "enable pending state tracking")
	f.String("node.forward.url", "", "url of another node to send transactions through")
	f.Int64("node.sequencer.create-batch-block-interval", 1, "block interval at which to create new batches")
	f.Int64("node.sequencer.delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	f.Bool("node.sequencer.enable", false, "act as sequencer")

	f.String("rollup.address", "", "layer 2 rollup contract address")
	f.Uint64("rollup.chain-id", 42161, "chain id of the arbitrum chain")
	f.String("rollup.machine.filename", "", "file to load machine from")

	f.String("feed.input.url", "", "URL of sequencer feed source")
	f.String("feed.output.addr", "0.0.0.0", "address to bind the relay feed output to")
	f.String("feed.output.port", "9642", "port to bind the relay feed output to")
	f.Duration("feed.output.ping", 5*time.Second, "number of seconds for ping interval")
	f.Duration("feed.output.timeout", 15*time.Second, "number of seconds for timeout")

	f.Bool("healthcheck.enable", false, "enable healthcheck endpoint")
	f.Bool("healthcheck.sequencer.enable", false, "enable checking the health of the sequencer")
	f.Bool("healthcheck.l1-node.enable", false, "enable checking the health of the L1 node")
	f.Bool("healthcheck.metrics.enable", false, "enable prometheus endpoint")
	f.String("healthcheck.metrics.prefix", "", "prepend the specified prefix to the exported metrics names")
	f.String("healthcheck.addr", "", "address to bind the healthcheck endpoint to")
	f.String("healthcheck.port", "", "port to bind the healthcheck endpoint to")

	f.String("l1.url", "", "layer 1 ethereum node RPC URL")

	f.String("log.rpc", "info", "log level for rpc")
	f.String("log.core", "info", "log level for general arb node logging")

	f.Bool("pprof.enable", false, "enable profiling server")

	f.String("persistent.storage.path", "state", "location persistent storage is located")

	f.String("rpc.addr", "0.0.0.0", "RPC address")
	f.String("rpc.port", "8547", "RPC port")

	f.String("validator.strategy", "", "strategy for validator to use")

	f.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")

	f.String("wallet.password", "", "password for wallet")
	f.Float64("wallet.gas-price", 4.5, "wallet.gasprice=FloatInGwei")

	f.String("ws.addr", "0.0.0.0", "websocket address")
	f.String("ws.port", "8548", "websocket port")

	err := f.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if f.NArg() != 0 {
		// Unexpected number of parameters
		return nil, nil, nil, nil, errors.New("unexpected number of parameters")
	}

	var k = koanf.New(".")
	// Load configuration file if provided

	configFile, _ := f.GetString("conf")
	if len(configFile) > 0 {
		if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "error loading config file")
		}
	}

	rollupAddress, _ := f.GetString("rollup.address")

	l1URL, err := f.GetString("l1.url")
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error getting --l1.url")
	}

	l1Client, err := ethutils.NewRPCEthClient(l1URL)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error running NewRPcEthClient")
	}

	var l1ChainId *big.Int
	for {
		l1ChainId, err = l1Client.ChainID(ctx)
		if err == nil {
			break
		}
		logger.Warn().Err(err).Msg("Error getting chain ID")

		select {
		case <-ctx.Done():
			return nil, nil, nil, nil, errors.New("ctx cancelled getting chain ID")
		case <-time.After(5 * time.Second):
		}
	}
	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	if len(rollupAddress) == 0 {
		if l1ChainId.Cmp(big.NewInt(1)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"rollup.address":          "0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A",
				"rollup.chain-id":         "42161",
				"rollup.from-block":       "12525700",
				"rollup.machine.filename": "mainnet.arb1.mexe",
				"bridge.utils.address":    "0x84efa170dc6d521495d7942e372b8e4b2fb918ec",
				"feed.input.url":          "wss://arb1.arbitrum.io/feed",
				"node.forward.url":        "https://arb1.arbitrum.io/rpc",
			}, "."), nil)

			if err != nil {
				return nil, nil, nil, nil, errors.Wrap(err, "error setting mainnet.arb1 rollup parameters")
			}
		} else if l1ChainId.Cmp(big.NewInt(4)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"rollup.address":          "0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c",
				"rollup.chain-id":         "421611",
				"rollup.from-block":       "8700589",
				"rollup.machine.filename": "testnet.rinkeby.mexe",
				"bridge.utils.address":    "0xA556F0eF1A0E37a7837ceec5527aFC7771Bf9a67",
				"feed.input.url":          "wss://rinkeby.arbitrum.io/feed",
				"node.forward.url":        "https://rinkeby.arbitrum.io/rpc",
			}, "."), nil)

			if err != nil {
				return nil, nil, nil, nil, errors.Wrap(err, "error setting testnet.rinkeby rollup parameters")
			}
		} else {
			return nil, nil, nil, nil, fmt.Errorf("connected to unrecognized ethereum network with chain ID: %v", l1ChainId)
		}
	}

	// Any settings provided on command line override items in configuration file
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error loading config")
	}

	var out Config
	decoderConfig := mapstructure.DecoderConfig{
		ErrorUnused: true,

		// Default values
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc()),
		Metadata:         nil,
		Result:           &out,
		WeaklyTypedInput: true,
	}
	err = k.UnmarshalWithConf("", &out, koanf.UnmarshalConf{DecoderConfig: &decoderConfig})
	if err != nil {

		return nil, nil, nil, nil, err
	}

	// Fixup directories
	if len(out.Persistent.Storage.Path) == 0 {
		// Error message will be output by caller
		return &out, nil, nil, nil, nil
	}

	if len(out.Persistent.Database.Path) == 0 {
		out.Persistent.Database.Path = path.Join(out.Persistent.Storage.Path, "arbStorage")
	}

	if len(out.Rollup.Machine.Filename) == 0 {
		// Nothing provided, so use default
		out.Rollup.Machine.Filename = path.Join(out.Persistent.Database.Path, "arbos.mexe")
	}

	if out.Dump.Conf {
		// Print out current configuration

		// Don't keep printing configuration file and don't print wallet password
		err := k.Load(confmap.Provider(map[string]interface{}{
			"dump.conf":       false,
			"wallet.password": "",
		}, "."), nil)

		c, err := k.Marshal(json.Parser())
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "unable to marshal config file to JSON")
		}

		fmt.Println(string(c))
		os.Exit(1)
	}

	// Don't pass around password with normal configuration
	wallet := out.Wallet
	out.Wallet.Password = ""

	return &out, &wallet, l1Client, l1ChainId, nil
}
