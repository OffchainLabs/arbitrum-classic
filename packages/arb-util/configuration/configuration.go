package configuration

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/mitchellh/mapstructure"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

var logger = log.With().Caller().Stack().Str("component", "configuration").Logger()

type FeedInput struct {
	Timeout time.Duration `koanf:"timeout"`
	URLs    []string      `koanf:"url"`
}

type FeedOutput struct {
	Addr          string        `koanf:"addr"`
	IOTimeout     time.Duration `koanf:"io-timeout"`
	Port          string        `koanf:"port"`
	Ping          time.Duration `koanf:"ping"`
	ClientTimeout time.Duration `koanf:"client-timeout"`
	Queue         int           `koanf:"queue"`
	Workers       int           `koanf:"workers"`
}

type Feed struct {
	Input  FeedInput  `koanf:"input"`
	Output FeedOutput `koanf:"output"`
}

type Healthcheck struct {
	Addr          string `koanf:"addr"`
	Enable        bool   `koanf:"enable"`
	L1Node        bool   `koanf:"l1-node"`
	Metrics       bool   `koanf:"metrics"`
	MetricsPrefix string `koanf:"metrics-prefix"`
	Port          string `koanf:"port"`
	Sequencer     bool   `koanf:"sequencer"`
}

type Lockout struct {
	Redis         string        `koanf:"redis"`
	SelfRPCURL    string        `koanf:"self-rpc-url"`
	Timeout       time.Duration `koanf:"timeout"`
	MaxLatency    time.Duration `koanf:"max-latency"`
	SeqNumTimeout time.Duration `koanf:"seq-num-timeout"`
}

type Aggregator struct {
	InboxAddress string `koanf:"inbox-address"`
	MaxBatchTime int64  `koanf:"max-batch-time"`
	Stateful     bool   `koanf:"stateful"`
}

type RPC struct {
	Addr string `koanf:"addr"`
	Port string `koanf:"port"`
}

type Sequencer struct {
	CreateBatchBlockInterval   int64   `koanf:"create-batch-block-interval"`
	DelayedMessagesTargetDelay int64   `koanf:"delayed-messages-target-delay"`
	Lockout                    Lockout `koanf:"lockout"`
}

type WS struct {
	Addr string `koanf:"addr"`
	Port string `koanf:"port"`
}

type Node struct {
	Aggregator Aggregator `koanf:"aggregator"`
	Forwarder  struct {
		Target string `koanf:"target"`
	} `koanf:"forwarder"`
	RPC       RPC       `koanf:"rpc"`
	Sequencer Sequencer `koanf:"sequencer"`
	Type      string    `koanf:"type"`
	WS        WS        `koanf:"ws"`
}

type Persistent struct {
	Chain        string `koanf:"chain"`
	GlobalConfig string `koanf:"global-config"`
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
	Strategy             string `koanf:"strategy"`
	UtilsAddress         string `koanf:"utils-address"`
	WalletFactoryAddress string `koanf:"wallet-factory-address"`
}

type Wallet struct {
	Password string `koanf:"password"`
}

type Log struct {
	RPC  string `koanf:"rpc"`
	Core string `koanf:"core"`
}

type Config struct {
	BridgeUtilsAddress string      `koanf:"bridge-utils-address"`
	Conf               string      `koanf:"conf"`
	DumpConf           bool        `koanf:"dump-conf"`
	EnvPrefix          string      `koanf:"env-prefix"`
	Feed               Feed        `koanf:"feed"`
	GasPrice           float64     `koanf:"gas-price"`
	GasPriceUrl        string      `koanf:"gas-price-url"`
	Healthcheck        Healthcheck `koanf:"healthcheck"`
	L1                 struct {
		URL string `koanf:"url"`
	} `koanf:"l1"`
	Log           Log        `koanf:"log"`
	Node          Node       `koanf:"node"`
	Persistent    Persistent `koanf:"persistent"`
	PProfEnable   bool       `koanf:"pprof-enable"`
	Rollup        Rollup     `koanf:"rollup"`
	Validator     Validator  `koanf:"validator"`
	WaitToCatchUp bool       `koanf:"wait-to-catch-up"`
	Wallet        Wallet     `koanf:"wallet"`
}

func (c *Config) GetNodeDatabasePath() string {
	return path.Join(c.Persistent.Chain, "db")
}

func (c *Config) GetValidatorDatabasePath() string {
	return path.Join(c.Persistent.Chain, "validator_db")
}

func ParseNode(ctx context.Context) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f := flag.NewFlagSet("", flag.ContinueOnError)

	AddFeedOutputOptions(f)

	f.String("node.aggregator.inbox-address", "", "address of the inbox contract")
	f.Int("node.aggregator.max-batch-time", 10, "max-batch-time=NumSeconds")
	f.Bool("node.aggregator.stateful", false, "enable pending state tracking")
	f.String("node.forwarder.target", "", "url of another node to send transactions through")
	f.String("node.rpc.addr", "0.0.0.0", "RPC address")
	f.Int("node.rpc.port", 8547, "RPC port")
	f.Int64("node.sequencer.create-batch-block-interval", 270, "block interval at which to create new batches")
	f.Int64("node.sequencer.delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	f.String("node.sequencer.lockout.redis", "", "sequencer lockout redis instance URL")
	f.String("node.sequencer.lockout.self-rpc-url", "", "own RPC URL for other sequencers to failover to")
	f.String("node.type", "forwarder", "forwarder, aggregator or sequencer")
	f.String("node.ws.addr", "0.0.0.0", "websocket address")
	f.Int("node.ws.port", 8548, "websocket port")

	return ParseNonRelay(ctx, f)
}

func ParseValidator(ctx context.Context) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f := flag.NewFlagSet("", flag.ContinueOnError)

	AddFeedOutputOptions(f)

	f.String("validator.strategy", "StakeLatest", "strategy for validator to use")
	f.String("validator.utils-address", "", "strategy for validator to use")
	f.String("validator.wallet-factory-address", "", "strategy for validator to use")

	return ParseNonRelay(ctx, f)
}

func ParseNonRelay(ctx context.Context, f *flag.FlagSet) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f.String("bridge-utils-address", "", "bridgeutils contract address")

	f.Float64("gas-price", 4.5, "gasprice=FloatInGwei")
	f.String("gas-price-url", "", "gas price rpc url (etherscan compatible)")

	f.String("rollup.address", "", "layer 2 rollup contract address")
	f.Uint64("rollup.chain-id", 42161, "chain id of the arbitrum chain")
	f.String("rollup.machine.filename", "", "file to load machine from")

	f.String("l1.url", "", "layer 1 ethereum node RPC URL")

	f.String("persistent.global-config", ".arbitrum", "location global configuration is located")
	f.String("persistent.chain", "", "path that chain specific state is located")

	f.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")

	f.String("wallet.password", "", "password for wallet")

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
				"bridge-utils-address":             "0x84efa170dc6d521495d7942e372b8e4b2fb918ec",
				"feed.input.url":                   []string{"wss://arb1.arbitrum.io/feed"},
				"node.forwarder.target":            "https://arb1.arbitrum.io/rpc",
				"persistent.chain":                 "mainnet",
				"rollup.address":                   "0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A",
				"rollup.chain-id":                  "42161",
				"rollup.from-block":                "12525700",
				"rollup.machine.filename":          "mainnet.arb1.mexe",
				"rollup.machine.url":               "https://raw.githubusercontent.com/OffchainLabs/arb-os/48bdb999a703575d26a856499e6eb3e17691e99d/arb_os/arbos.mexe",
				"validator.utils-address":          "0x2B36F23ce0bAbD57553b26Da4C7a0585bac65DC1",
				"validator.wallet-factory-address": "0xe17d8Fa6BC62590f840C5Dd35f300F77D55CC178",
			}, "."), nil)

			if err != nil {
				return nil, nil, nil, nil, errors.Wrap(err, "error setting mainnet.arb1 rollup parameters")
			}
		} else if l1ChainId.Cmp(big.NewInt(4)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"bridge-utils-address":             "0xA556F0eF1A0E37a7837ceec5527aFC7771Bf9a67",
				"feed.input.url":                   []string{"wss://rinkeby.arbitrum.io/feed"},
				"node.forwarder.target":            "https://rinkeby.arbitrum.io/rpc",
				"persistent.chain":                 "rinkeby",
				"rollup.address":                   "0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c",
				"rollup.chain-id":                  "421611",
				"rollup.from-block":                "8700589",
				"rollup.machine.filename":          "testnet.rinkeby.mexe",
				"rollup.machine.url":               "https://raw.githubusercontent.com/OffchainLabs/arb-os/26ab8d7c818681c4ee40792aeb12981a8f2c3dfa/arb_os/arbos.mexe",
				"validator.utils-address":          "0xbb14D9837f6E596167638Ba0963B9Ba8351F68CD",
				"validator.wallet-factory-address": "0x5533D1578a39690B6aC692673F771b3fc668f0a3",
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
	if !filepath.IsAbs(out.Persistent.GlobalConfig) {
		out.Persistent.GlobalConfig = path.Join(homeDir, out.Persistent.GlobalConfig)
	}
	err = os.MkdirAll(out.Persistent.GlobalConfig, os.ModePerm)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to create global configuration directory")
	}

	// Make chain directory relative to persistent storage directory if not already absolute
	if !filepath.IsAbs(out.Persistent.Chain) {
		out.Persistent.Chain = path.Join(out.Persistent.GlobalConfig, out.Persistent.Chain)
	}
	err = os.MkdirAll(out.Persistent.Chain, os.ModePerm)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "Unable to create chain directory")
	}

	if len(out.Rollup.Machine.Filename) == 0 {
		// Machine not provided, so use default
		out.Rollup.Machine.Filename = path.Join(out.Persistent.Chain, "arbos.mexe")
	}

	// Make machine relative to storage directory if not already absolute
	out.Rollup.Machine.Filename = path.Join(out.Persistent.GlobalConfig, out.Rollup.Machine.Filename)

	_, err = os.Stat(out.Rollup.Machine.Filename)
	if os.IsNotExist(err) && len(out.Rollup.Machine.URL) != 0 {
		// Machine does not exist, so load it from provided URL
		logger.Debug().Str("URL", out.Rollup.Machine.URL).Msg("downloading machine")

		resp, err := http.Get(out.Rollup.Machine.URL)
		if err != nil {
			return nil, nil, nil, nil, errors.Wrapf(err, "unable to get machine from: %s", out.Rollup.Machine.URL)
		}
		if resp.StatusCode != 200 {
			return nil, nil, nil, nil, fmt.Errorf("HTTP status '%v' when trying to get machine from: %s", resp.Status, out.Rollup.Machine.URL)
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

func ParseRelay() (*Config, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	AddFeedOutputOptions(f)

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

func AddFeedOutputOptions(f *flag.FlagSet) {
	f.String("feed.output.addr", "0.0.0.0", "address to bind the relay feed output to")
	f.Duration("feed.output.io-timeout", 5*time.Second, "duration to wait before timing out HTTP to WS upgrade")
	f.Int("feed.output.port", 9642, "port to bind the relay feed output to")
	f.Duration("feed.output.ping", 5*time.Second, "duration for ping interval")
	f.Duration("feed.output.client-timeout", 15*time.Second, "duraction to wait before timing out connections to client")
	f.Int("feed.output.workers", 100, "Number of threads to reserve for HTTP to WS upgrade")
}

func beginCommonParse(f *flag.FlagSet) (*koanf.Koanf, error) {
	f.String("conf", "", "name of configuration file")

	f.Bool("dump-conf", false, "print out currently active configuration file")

	f.String("env-prefix", "", "environment variables with given prefix will be loaded as configuration values")

	f.Duration("feed.input.timeout", 20*time.Second, "duration to wait before timing out connection to server")
	f.StringSlice("feed.input.url", []string{}, "URL of sequencer feed source")

	f.Bool("healthcheck.enable", false, "enable healthcheck endpoint")
	f.Bool("healthcheck.sequencer", false, "enable checking the health of the sequencer")
	f.Bool("healthcheck.l1-node", false, "enable checking the health of the L1 node")
	f.Bool("healthcheck.metrics", false, "enable prometheus endpoint")
	f.String("healthcheck.metrics-prefix", "", "prepend the specified prefix to the exported metrics names")
	f.String("healthcheck.addr", "", "address to bind the healthcheck endpoint to")
	f.Int("healthcheck.port", 0, "port to bind the healthcheck endpoint to")

	f.String("log.rpc", "info", "log level for rpc")
	f.String("log.core", "info", "log level for general arb node logging")

	f.Bool("pprof-enable", false, "enable profiling server")

	err := f.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	if f.NArg() != 0 {
		// Unexpected number of parameters
		return nil, errors.New("unexpected number of parameters")
	}

	var k = koanf.New(".")

	// Load defaults that are not specified on command line
	err = k.Load(confmap.Provider(map[string]interface{}{
		"feed.output.queue":                      100,
		"node.sequencer.lockout.timeout":         30 * time.Second,
		"node.sequencer.lockout.max-latency":     10 * time.Second,
		"node.sequencer.lockout.seq-num-timeout": 5 * time.Minute,
	}, "."), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error applying default values")
	}

	// Load configuration file if provided
	configFile, _ := f.GetString("conf")
	if len(configFile) > 0 {
		if err = k.Load(file.Provider(configFile), json.Parser()); err != nil {
			return nil, errors.Wrap(err, "error loading config file")
		}
	}

	// Any settings provided on command line override items in configuration file
	// Command line parameters will be applied again later
	if err = k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, errors.Wrap(err, "error loading config")
	}

	// Any settings provided through environment variables override all other custom settings
	// Environment variable parameters will be applied again later
	err = loadEnvironmentVariables(k)
	if err != nil {
		return nil, errors.Wrap(err, "error loading environment variables")
	}

	return k, nil
}

func loadEnvironmentVariables(k *koanf.Koanf) error {
	envPrefix := k.String("env-prefix")
	if len(envPrefix) != 0 {
		return k.Load(env.Provider(envPrefix+"_", ".", func(s string) string {
			// FOO__BAR -> foo-bar to handle dash in config names
			s = strings.Replace(strings.ToLower(
				strings.TrimPrefix(s, envPrefix+"_")), "__", "-", -1)
			return strings.Replace(s, "_", ".", -1)
		}), nil)
	}

	return nil
}

func endCommonParse(f *flag.FlagSet, k *koanf.Koanf) (*Config, *Wallet, error) {
	// Any settings provided on command line override any custom settings
	// Second time command line parameters are applied so auto chain parameters can be overridden
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, nil, errors.Wrap(err, "error loading config")
	}

	// Any settings provided through environment variables override all other custom settings
	// Second time environment variables are applied so auto chain parameters can be overridden
	err := loadEnvironmentVariables(k)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error loading environment variables")
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

		return nil, nil, err
	}

	if out.DumpConf {
		// Print out current configuration

		// Don't keep printing configuration file and don't print wallet password
		err := k.Load(confmap.Provider(map[string]interface{}{
			"dump-conf":       false,
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
