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
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
	"io"
	"math/big"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

var logger = log.With().Caller().Stack().Str("component", "configuration").Logger()

type Bridge struct {
	Utils struct {
		Address string `koanf:"address"`
	} `koanf:"utils"`
}

type FeedInput struct {
	Timeout time.Duration `koanf:"timeout"`
	URL     string        `koanf:"url"`
}

type FeedOutput struct {
	Addr string `koanf:"addr"`
	HTTP struct {
		Timeout time.Duration `koanf:"timeout"`
	} `koanf:"http"`
	Port    string        `koanf:"port"`
	Ping    time.Duration `koanf:"ping"`
	Timeout time.Duration `koanf:"timeout"`
	Queue   int           `koanf:"queue"`
	Workers int           `koanf:"workers"`
}

type Feed struct {
	Input  FeedInput  `koanf:"input"`
	Output FeedOutput `koanf:"output"`
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
	Chain struct {
		Path string `koanf:"path"`
	} `koanf:"chain"`
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
		URL      string `koanf:"url"`
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
	f := flag.NewFlagSet("", flag.ContinueOnError)

	f.String("bridge.utils.address", "", "bridgeutils contract address")

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

	f.String("l1.url", "", "layer 1 ethereum node RPC URL")

	f.String("persistent.storage.path", ".arbitrum", "location persistent storage is located")

	f.String("rpc.addr", "0.0.0.0", "RPC address")
	f.String("rpc.port", "8547", "RPC port")

	f.String("validator.strategy", "", "strategy for validator to use")

	f.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")

	f.String("wallet.password", "", "password for wallet")
	f.Float64("wallet.gas-price", 4.5, "wallet.gasprice=FloatInGwei")

	f.String("ws.addr", "0.0.0.0", "websocket address")
	f.String("ws.port", "8548", "websocket port")

	k, err := beginCommonParse(f)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	l1URL := k.String("l1.url")
	if len(l1URL) == 0 {
		return nil, nil, nil, nil, errors.New("required parameter --l1.url is missing")
	}

	l1Client, err := ethutils.NewRPCEthClient(l1URL)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error running NewRPCEthClient")
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

	if len(k.String("rollup.address")) == 0 {
		if l1ChainId.Cmp(big.NewInt(1)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"bridge.utils.address":     "0x84efa170dc6d521495d7942e372b8e4b2fb918ec",
				"feed.input.url":           "wss://arb1.arbitrum.io/feed",
				"node.forward.url":         "https://arb1.arbitrum.io/rpc",
				"persistent.chain.path":    "mainnet",
				"persistent.database.path": "db",
				"rollup.address":           "0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A",
				"rollup.chain-id":          "42161",
				"rollup.from-block":        "12525700",
				"rollup.machine.filename":  "mainnet.arb1.mexe",
				"rollup.machine.url":       "https://raw.githubusercontent.com/OffchainLabs/arb-os/48bdb999a703575d26a856499e6eb3e17691e99d/arb_os/arbos.mexe",
			}, "."), nil)

			if err != nil {
				return nil, nil, nil, nil, errors.Wrap(err, "error setting mainnet.arb1 rollup parameters")
			}
		} else if l1ChainId.Cmp(big.NewInt(4)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"bridge.utils.address":     "0xA556F0eF1A0E37a7837ceec5527aFC7771Bf9a67",
				"feed.input.url":           "wss://rinkeby.arbitrum.io/feed",
				"node.forward.url":         "https://rinkeby.arbitrum.io/rpc",
				"persistent.chain.path":    "rinkeby",
				"persistent.database.path": "db",
				"rollup.address":           "0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c",
				"rollup.chain-id":          "421611",
				"rollup.from-block":        "8700589",
				"rollup.machine.filename":  "testnet.rinkeby.mexe",
				"rollup.machine.url":       "https://raw.githubusercontent.com/OffchainLabs/arb-os/26ab8d7c818681c4ee40792aeb12981a8f2c3dfa/arb_os/arbos.mexe --output /home/user/state/arbos.mexe",
			}, "."), nil)

			if err != nil {
				return nil, nil, nil, nil, errors.Wrap(err, "error setting testnet.rinkeby rollup parameters")
			}
		} else {
			return nil, nil, nil, nil, fmt.Errorf("connected to unrecognized ethereum network with chain ID: %v", l1ChainId)
		}
	}

	out, wallet, err := endCommonParse(f, k)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Fixup directories
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to read users home directory")
	}

	// Make persistent storage directory relative to home directory if not already absolute
	if !filepath.IsAbs(out.Persistent.Storage.Path) {
		out.Persistent.Storage.Path = path.Join(homeDir, out.Persistent.Storage.Path)
	}
	err = os.MkdirAll(out.Persistent.Storage.Path, os.ModePerm)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to create storage directory")
	}

	// Make chain directory relative to persistent storage directory if not already absolute
	if !filepath.IsAbs(out.Persistent.Chain.Path) {
		out.Persistent.Chain.Path = path.Join(out.Persistent.Storage.Path, out.Persistent.Chain.Path)
	}
	err = os.MkdirAll(out.Persistent.Chain.Path, os.ModePerm)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to create chain directory")
	}

	// Make db directory relative to chain directory if not already absolute
	if !filepath.IsAbs(out.Persistent.Database.Path) {
		out.Persistent.Database.Path = path.Join(out.Persistent.Chain.Path, out.Persistent.Database.Path)
	}
	err = os.MkdirAll(out.Persistent.Database.Path, os.ModePerm)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to create database directory")
	}

	if len(out.Rollup.Machine.Filename) == 0 {
		// Machine not provided, so use default
		out.Rollup.Machine.Filename = path.Join(out.Persistent.Storage.Path, "arbos.mexe")
	}

	// Make machine relative to storage directory if not already absolute
	out.Rollup.Machine.Filename = path.Join(out.Persistent.Storage.Path, out.Rollup.Machine.Filename)

	_, err = os.Stat(out.Rollup.Machine.Filename)
	if os.IsNotExist(err) && len(out.Rollup.Machine.URL) != 0 {
		// Machine does not exist, so load it from provided URL
		logger.Debug().Str("URL", out.Rollup.Machine.URL).Msg("downloading machine")

		resp, err := http.Get(out.Rollup.Machine.URL)
		if err != nil {
			return nil, nil, nil, nil, errors.Wrapf(err, "unable to get machine from: %s", out.Rollup.Machine.URL)
		}

		fileOut, err := os.Create(out.Rollup.Machine.Filename)
		if err != nil {
			return nil, nil, nil, nil, errors.Wrapf(err, "unable to open file '%s' for machine", out.Rollup.Machine.Filename)
		}

		_, err = io.Copy(fileOut, resp.Body)
		if err != nil {
			return nil, nil, nil, nil, errors.Wrapf(err, "unable to output machine to: %s", out.Rollup.Machine.Filename)
		}
	}

	return out, wallet, l1Client, l1ChainId, nil
}

func ParseFeed(ctx context.Context) (*Config, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	k, err := beginCommonParse(f)
	if err != nil {
		return nil, err
	}

	out, _, err := endCommonParse(f, k)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func beginCommonParse(f *flag.FlagSet) (*koanf.Koanf, error) {
	f.String("conf", "", "name of configuration file")

	f.Bool("dump.conf", false, "print out currently active configuration file")

	f.Duration("feed.input.timeout", 20*time.Second, "duration to wait before timing out connection to server")
	f.String("feed.input.url", "", "URL of sequencer feed source")
	f.String("feed.output.addr", "0.0.0.0", "address to bind the relay feed output to")
	f.Duration("feed.output.http.timeout", 5*time.Second, "duration to wait before timing out HTTP to WS upgrade")
	f.String("feed.output.port", "9642", "port to bind the relay feed output to")
	f.Duration("feed.output.ping", 5*time.Second, "duration for ping interval")
	f.Duration("feed.output.timeout", 15*time.Second, "duraction to wait before timing out connections to client")
	f.Int("feed.output.workers", 100, "Number of threads to reserve for HTTP to WS upgrade")

	f.Bool("healthcheck.enable", false, "enable healthcheck endpoint")
	f.Bool("healthcheck.sequencer.enable", false, "enable checking the health of the sequencer")
	f.Bool("healthcheck.l1-node.enable", false, "enable checking the health of the L1 node")
	f.Bool("healthcheck.metrics.enable", false, "enable prometheus endpoint")
	f.String("healthcheck.metrics.prefix", "", "prepend the specified prefix to the exported metrics names")
	f.String("healthcheck.addr", "", "address to bind the healthcheck endpoint to")
	f.String("healthcheck.port", "", "port to bind the healthcheck endpoint to")

	f.String("log.rpc", "info", "log level for rpc")
	f.String("log.core", "info", "log level for general arb node logging")

	f.Bool("pprof.enable", false, "enable profiling server")

	err := f.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	if f.NArg() != 0 {
		// Unexpected number of parameters
		return nil, errors.New("unexpected number of parameters")
	}

	var k = koanf.New(".")
	// Load configuration file if provided

	configFile, _ := f.GetString("conf")
	if len(configFile) > 0 {
		if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
			return nil, errors.Wrap(err, "error loading config file")
		}
	}

	return k, nil
}

func endCommonParse(f *flag.FlagSet, k *koanf.Koanf) (*Config, *Wallet, error) {
	// Any settings provided on command line override items in configuration file
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, nil, errors.Wrap(err, "error loading config")
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
	err := k.UnmarshalWithConf("", &out, koanf.UnmarshalConf{DecoderConfig: &decoderConfig})
	if err != nil {

		return nil, nil, err
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
			return nil, nil, errors.Wrap(err, "unable to marshal config file to JSON")
		}

		fmt.Println(string(c))
		os.Exit(1)
	}

	// Don't pass around password with normal configuration
	wallet := out.Wallet
	out.Wallet.Password = ""

	return &out, &wallet, nil
}
