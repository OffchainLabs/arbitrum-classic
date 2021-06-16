package configuration

import (
	"github.com/knadh/koanf/providers/posflag"
	"os"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

var logger = log.With().Caller().Stack().Str("component", "message").Logger()

type Wallet struct {
	Password string  `koanf:"password"`
	GasPrice float64 `koanf:"gasprice"`
}

type Config struct {
	Conf                       string `koanf:"conf"`
	Pending                    bool   `koanf:"pending"`
	Sequencer                  bool   `koanf:"pending"`
	WaitToCatchUp              bool   `koanf:"wait-to-catch-up"`
	DelayedMessagesTargetDelay int64  `koanf:"delayed-messages-target-delay"`
	CreateBatchBlockInterval   int64  `koanf:"create-batch-block-interval"`
	GasPriceUrl                string `koanf:"gas-price-url"`
	ChainID                    uint64 `koanf:"chainid"`
	Healthcheck                struct {
		Enabled   bool `koanf:"healthcheck.enabled"`
		Sequencer struct {
			Enabled bool `koanf:"healthcheck.sequencer.enabled"`
		}
		L1Node struct {
			Enabled bool `koanf:"healthcheck.l1-node.enabled"`
		}
		Metrics struct {
			Enabled bool   `koanf:"healthcheck.metrics.enabled"`
			Prefix  string `koanf:"healthcheck.metrics.prefix"`
		}
		Addr string `koanf:"healthcheck.addr"`
		Port string `koanf:"healthcheck.port"`
	}

	MaxBatchTime int64 `koanf:"max-batch-time"`
	Inbox        struct {
		Address string `koanf:"inbox.address"`
	}
	Forward struct {
		URL string `koanf:"forward.url"`
	}
	Feed struct {
		URL    string `koanf:"feed.url"`
		Output struct {
			Addr    string `koanf:"feed.output.addr"`
			Port    string `koanf:"feed.output.port"`
			Ping    string `koanf:"feed.output.ping"`
			Timeout string `koanf:"feed.output.timeout"`
		}
	}
	RPC struct {
		LogLevel string `koanf:"rcp.loglevel"`
		Addr     string `koanf:"rpc.addr"`
		Port     string `koanf:"rpc.port"`
	}
	WS struct {
		Addr string `koanf:"ws.addr"`
		Port string `koanf:"ws.port"`
	}
	Pprof struct {
		Enabled bool `koanf:"pprof.enabled"`
	}
	LogLevel string `koanf:"loglevel"`

	Database struct {
		Path string `koanf:"database.path"`
	}
	L1 struct {
		URL string `koanf:"l1.url"`
	}
	Rollup struct {
		Address string `koanf:"rollup.address"`
		Machine struct {
			Filename string `koanf:"rollup.machine.filename"`
			//URL string `koanf:"rollup.machine.url"`
		}
	}
	Bridge struct {
		Utils struct {
			Address string `koanf:"bridge.utils.address"`
		}
	}
	Validator struct {
		Utils struct {
			Address string `koanf:"validator.utils.address"`
		}
		WalletFactory struct {
			Address string `koanf:"validator.walletfactory.address"`
		}
		Strategy string `koanf:"validator.strategy"`
	}
	Mainnet struct {
		Arb1 bool `koanf:"mainnet.arb1"`
	}
	Testnet struct {
		rinkeby bool `koanf:"testnet.rinkeby"`
	}
}

func Parse() (*Config, *Wallet, error) {
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.String("conf", "", "name of configuration file")
	f.String("wallet.password", "", "password=pass")
	f.Float64("wallet.gasprice", 4.5, "gasprice=FloatInGwei")
	f.Bool("pending", false, "enable pending state tracking")
	f.Bool("sequencer", false, "act as sequencer")
	f.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")
	f.Int64("delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	f.Int64("create-batch-block-interval", 1, "block interval at which to create new batches")
	f.String("gas-price-url", "", "gas price rpc url (etherscan compatible)")
	f.String("validator.strategy", "", "strategy for validator to use")

	f.Uint64("chainid", 42161, "chain id of the arbitrum chain")

	// Healthcheck Config
	f.Bool("healthcheck.enabled", false, "enable healthcheck endpoint")
	f.Bool("healthcheck.sequencer.enabled", false, "enable checking the health of the sequencer")
	f.Bool("healthcheck.l1-node.enabled", false, "enable checking the health of the L1 node")
	f.Bool("healthcheck.metrics.enabled", false, "enable prometheus endpoint")
	f.String("healthcheck.metrics.prefix", "", "prepend the specified prefix to the exported metrics names")
	f.String("healthcheck.addr", "", "address to bind the healthcheck endpoint to")
	f.String("healthcheck.port", "", "port to bind the healthcheck endpoint to")

	f.Int64("maxBatchTime", 10, "maxBatchTime=NumSeconds")

	f.String("inbox.address", "", "address of the inbox contract")
	f.String("forward.url", "", "url of another node to send transactions through")
	f.String("feed.url", "", "URL of sequencer feed source")
	f.String("rpc.addr", "0.0.0.0", "RPC address")
	f.String("rpc.port", "8547", "RPC port")
	f.String("ws.addr", "0.0.0.0", "websocket address")
	f.String("ws.port", "8548", "websocket port")
	f.String("feedoutput.addr", "0.0.0.0", "address to bind the relay feed output to")
	f.String("feedoutput.port", "9642", "port to bind the relay feed output to")
	f.Duration("feedoutput.ping", 5*time.Second, "number of seconds for ping interval")
	f.Duration("feedoutput.timeout", 15*time.Second, "number of seconds for timeout")
	f.Bool("pprof.enabled", false, "enable profiling server")
	f.String("rpc.loglevel", "info", "log level for rpc")
	f.String("arbloglevel", "info", "log level for general arb node logging")

	f.String("database.path", "state", "location database is located")
	f.String("l1.url", "", "layer 1 ethereum node RPC URL")
	f.String("rollup.address", "", "layer 2 rollup contract address")
	f.String("bridgeutils.address", "", "bridgeutils contract address")

	f.Bool("mainnet.arb1", false, "connect to arb1 mainnet")
	f.Bool("testnet.rinkeby", false, "connect to rinkeby testnet")

	err := f.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, errors.Wrap(err, "error parsing arguments")
	}

	var k = koanf.New(".")
	// Load configuration file if provided
	configFile, _ := f.GetString("conf")
	if len(configFile) > 0 {
		if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
			return nil, nil, errors.Wrap(err, "error loading config file")
		}
	}

	if useArb1, _ := f.GetBool("mainnet.arb1"); useArb1 {
		err := k.Load(confmap.Provider(map[string]interface{}{
			"rollup.address":      "0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A",
			"rollup.startblock":   "12640751",
			"bridgeutils.address": "0x84efa170dc6d521495d7942e372b8e4b2fb918ec",
			"feed.url":            "wss://arb1.arbitrum.io/feed",
			"forward.url":         "https://arb1.arbitrum.io/rpc",
			"chainid":             "42161",
		}, "."), nil)

		if err != nil {
			return nil, nil, errors.Wrap(err, "error setting mainnet.arb1 rollup parameters")
		}
	}

	if useRinkeby, _ := f.GetBool("testnet.rinkeby"); useRinkeby {
		err := k.Load(confmap.Provider(map[string]interface{}{
			"rollup.address":      "0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c",
			"rollup.startblock":   "8700874",
			"bridgeutils.address": "0xA556F0eF1A0E37a7837ceec5527aFC7771Bf9a67",
			"feed.url":            "wss://rinkeby.arbitrum.io/feed",
			"forward.url":         "https://rinkeby.arbitrum.io/rpc",
			"chainid":             "421611",
		}, "."), nil)

		if err != nil {
			return nil, nil, errors.Wrap(err, "error setting testnet.rinkeby rollup parameters")
		}
	}

	if f.NArg() == 4 {
		// Support legacy parameters
		validatorFolder := f.Arg(0)
		ethURL := f.Arg(1)
		addressString := f.Arg(2)
		bridgeUtilsAddressString := f.Arg(3)

		err := k.Load(confmap.Provider(map[string]interface{}{
			"database.path":       validatorFolder,
			"l1.url":              ethURL,
			"rollup.address":      addressString,
			"bridgeutils.address": bridgeUtilsAddressString,
		}, "."), nil)

		if err != nil {
			return nil, nil, errors.Wrap(err, "error parsing rollup parameters")
		}
	} else if f.NArg() == 6 {
		// Support legacy parameters
		validatorFolder := f.Arg(0)
		ethURL := f.Arg(1)
		addressString := f.Arg(2)
		bridgeUtilsAddressString := f.Arg(3)
		validatorUtilsAddressString := f.Arg(4)
		validatorWalletFactoryAddressString := f.Arg(5)

		err := k.Load(confmap.Provider(map[string]interface{}{
			"database.path":                   validatorFolder,
			"l1.url":                          ethURL,
			"rollup.address":                  addressString,
			"bridgeutils.address":             bridgeUtilsAddressString,
			"validator.utils.address":         validatorUtilsAddressString,
			"validator.walletfactory.address": validatorWalletFactoryAddressString,
		}, "."), nil)

		if err != nil {
			return nil, nil, errors.Wrap(err, "error parsing rollup parameters")
		}
	} else if f.NArg() != 0 {
		// Unexpected number of parameters
		return nil, nil, errors.Wrap(err, "unexpected number of parameters")
	}

	// Any settings provided on command line override items in configuration file
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, nil, errors.Wrap(err, "error loading config")
	}

	var out Config
	err = k.Unmarshal("", &out)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error unmarshalling configuration")
	}

	var wallet Wallet
	err = k.Unmarshal("wallet", &wallet)

	return &out, &wallet, nil
}
