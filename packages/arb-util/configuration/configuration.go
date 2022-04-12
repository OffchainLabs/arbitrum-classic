/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package configuration

import (
	"context"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
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
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/providers/s3"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"

	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

const PASSWORD_NOT_SET = "PASSWORD_NOT_SET"

var logger = arblog.Logger.With().Str("component", "configuration").Logger()

type Conf struct {
	Dump      bool   `koanf:"dump"`
	EnvPrefix string `koanf:"env-prefix"`
	File      string `koanf:"file"`
	S3        S3     `koanf:"s3"`
	String    string `koanf:"string"`
}

type Database struct {
	Compact       bool          `koanf:"compact"`
	ExitAfter     bool          `koanf:"exit-after"`
	Metadata      bool          `koanf:"metadata"`
	L0Files       int           `koanf:"l0-files"`
	SaveInterval  time.Duration `koanf:"save-interval"`
	SaveOnStartup bool          `koanf:"save-on-startup"`
	SavePath      string        `koanf:"save-path"`
	Threads       int           `koanf:"threads"`
}

type Core struct {
	AddMessagesMaxFailureCount int           `koanf:"add-messages-max-failure-count"`
	Cache                      CoreCache     `koanf:"cache"`
	CheckpointGasFrequency     int           `koanf:"checkpoint-gas-frequency"`
	CheckpointLoadGasCost      int           `koanf:"checkpoint-load-gas-cost"`
	CheckpointLoadGasFactor    int           `koanf:"checkpoint-load-gas-factor"`
	CheckpointMaxExecutionGas  int           `koanf:"checkpoint-max-execution-gas"`
	CheckpointMaxToPrune       int           `koanf:"checkpoint-max-to-prune"`
	CheckpointPruningMode      string        `koanf:"checkpoint-pruning-mode"`
	CheckpointPruneOnStartup   bool          `koanf:"checkpoint-prune-on-startup"`
	Database                   Database      `koanf:"database"`
	Debug                      bool          `koanf:"debug"`
	DebugTiming                bool          `koanf:"debug-timing"`
	IdleSleep                  time.Duration `koanf:"idle-sleep"`
	LazyLoadCoreMachine        bool          `koanf:"lazy-load-core-machine"`
	LazyLoadArchiveQueries     bool          `koanf:"lazy-load-archive-queries"`
	MessageProcessCount        int           `koanf:"message-process-count"`
	Test                       CoreTest      `koanf:"test"`
}

type CoreCache struct {
	BasicInterval int           `koanf:"basic-interval"`
	BasicSize     int           `koanf:"basic-size"`
	Disable       bool          `koanf:"disable"`
	Last          bool          `koanf:"last"`
	LRUSize       int           `koanf:"lru-size"`
	SeedOnStartup bool          `koanf:"seed-on-startup"`
	TimedExpire   time.Duration `koanf:"timed-expire"`
}

type CoreTest struct {
	LoadCount           int64       `koanf:"load-count"`
	ReorgTo             TestReorgTo `koanf:"reorg-to"`
	ResetAllExceptInbox bool        `koanf:"reset-all-except-inbox"`
	RunUntil            int64       `koanf:"run-until"`
}

type TestReorgTo struct {
	L1Block int64 `koanf:"l1-block"`
	L2Block int64 `koanf:"l2-block"`
	Log     int64 `koanf:"log"`
	Message int64 `koanf:"message"`
}

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

type Tracing struct {
	Enable    bool   `koanf:"enable"`
	Namespace string `koanf:"namespace"`
}

type RPC struct {
	Addr              string  `koanf:"addr"`
	Port              string  `koanf:"port"`
	Path              string  `koanf:"path"`
	EnableL1Calls     bool    `koanf:"enable-l1-calls"`
	Tracing           Tracing `koanf:"tracing"`
	MaxCallGas        uint64  `koanf:"max-call-gas"`
	EnableDevopsStubs bool    `koanf:"enable-devops-stubs"`
}

type S3 struct {
	AccessKey string `koanf:"access-key"`
	Bucket    string `koanf:"bucket"`
	ObjectKey string `koanf:"object-key"`
	Region    string `koanf:"region"`
	SecretKey string `koanf:"secret-key"`
}

type L1PostingStrategy struct {
	HighGasThreshold   float64 `koanf:"high-gas-threshold"`
	HighGasDelayBlocks int64   `koanf:"high-gas-delay-blocks"`
}

type SequencerDangerous struct {
	ReorgOutHugeMessages            bool `koanf:"reorg-out-huge-messages" json:"reorg-out-huge-messages"`
	PublishBatchesWithoutLockout    bool `koanf:"publish-batches-without-lockout" json:"publish-batches-without-lockout"`
	RewriteSequencerAddress         bool `koanf:"rewrite-sequencer-address" json:"rewrite-sequencer-address"`
	DisableBatchPosting             bool `koanf:"disable-batch-posting" json:"disable-batch-posting"`
	DisableDelayedMessageSequencing bool `koanf:"disable-delayed-message-sequencing" json:"disable-delayed-message-sequencing"`
}

type Sequencer struct {
	CreateBatchBlockInterval          int64              `koanf:"create-batch-block-interval"`
	ContinueBatchPostingBlockInterval int64              `koanf:"continue-batch-posting-block-interval"`
	DelayedMessagesTargetDelay        int64              `koanf:"delayed-messages-target-delay"`
	Lockout                           Lockout            `koanf:"lockout"`
	L1PostingStrategy                 L1PostingStrategy  `koanf:"l1-posting-strategy"`
	MaxBatchGasCost                   int64              `koanf:"max-batch-gas-cost"`
	GasRefunderAddress                string             `koanf:"gas-refunder-address"`
	GasRefunderExtraGas               uint64             `koanf:"gas-refunder-extra-gas"`
	Dangerous                         SequencerDangerous `koanf:"dangerous"`
	DebugTiming                       bool               `koanf:"debug-timing"`
}

type WS struct {
	Addr string `koanf:"addr"`
	Port string `koanf:"port"`
	Path string `koanf:"path"`
}

type Forwarder struct {
	Target    string `koanf:"target"`
	Submitter string `koanf:"submitter-address"`
	RpcMode   string `koanf:"rpc-mode"`
}

type InboxReader struct {
	DelayBlocks              int64         `koanf:"delay-blocks"`
	Paranoid                 bool          `koanf:"paranoid"`
	SequencerSignatureExpiry time.Duration `koanf:"sequencer-signature-expiry"`
}

type Node struct {
	Aggregator      Aggregator    `koanf:"aggregator"`
	Cache           NodeCache     `koanf:"cache"`
	ChainID         uint64        `koanf:"chain-id"`
	Forwarder       Forwarder     `koanf:"forwarder"`
	InboxReader     InboxReader   `koanf:"inbox-reader"`
	LogProcessCount int           `koanf:"log-process-count"`
	LogIdleSleep    time.Duration `koanf:"log-idle-sleep"`
	RPC             RPC           `koanf:"rpc"`
	Sequencer       Sequencer     `koanf:"sequencer"`
	TypeImpl        string        `koanf:"type"`
	WS              WS            `koanf:"ws"`
}

type NodeType uint8

const (
	UnknownNodeType NodeType = iota
	ForwarderNodeType
	AggregatorNodeType
	SequencerNodeType
	ValidatorNodeType
)

func (c *Node) Type() NodeType {
	if strings.EqualFold(c.TypeImpl, "forwarder") {
		return ForwarderNodeType
	} else if strings.EqualFold(c.TypeImpl, "aggregator") {
		return AggregatorNodeType
	} else if strings.EqualFold(c.TypeImpl, "sequencer") {
		return SequencerNodeType
	} else if strings.EqualFold(c.TypeImpl, "validator") {
		return ValidatorNodeType
	} else {
		return UnknownNodeType
	}
}

type NodeCache struct {
	AllowSlowLookup  bool          `koanf:"allow-slow-lookup"`
	LRUSize          int           `koanf:"lru-size"`
	BlockInfoLRUSize int           `koanf:"block-info-lru-size"`
	TimedInitialSize int           `koanf:"timed-initial-size"`
	TimedExpire      time.Duration `koanf:"timed-expire"`
}

type Persistent struct {
	Chain        string `koanf:"chain"`
	GlobalConfig string `koanf:"global-config"`
}

type Rollup struct {
	Address   string `koanf:"address"`
	FromBlock int64  `koanf:"from-block"`
	Machine   struct {
		Filename string `koanf:"filename"`
		URL      string `koanf:"url"`
	} `koanf:"machine"`
}

type Validator struct {
	StrategyImpl         string            `koanf:"strategy"`
	UtilsAddress         string            `koanf:"utils-address"`
	StakerDelay          time.Duration     `koanf:"staker-delay"`
	WalletFactoryAddress string            `koanf:"wallet-factory-address"`
	L1PostingStrategy    L1PostingStrategy `koanf:"l1-posting-strategy"`
	DontChallenge        bool              `koanf:"dont-challenge"`
	WithdrawDestination  string            `koanf:"withdraw-destination"`
}

type ValidatorStrategy uint8

const (
	UnknownStrategy ValidatorStrategy = iota
	WatchtowerStrategy
	DefensiveStrategy
	StakeLatestStrategy
	MakeNodesStrategy
)

func (s ValidatorStrategy) IsActive() bool {
	if s == StakeLatestStrategy || s == MakeNodesStrategy {
		return true
	}

	return false
}

func (c *Validator) Strategy() ValidatorStrategy {
	if strings.EqualFold(c.StrategyImpl, "Watchtower") {
		return WatchtowerStrategy
	} else if strings.EqualFold(c.StrategyImpl, "Defensive") {
		return DefensiveStrategy
	} else if strings.EqualFold(c.StrategyImpl, "StakeLatest") {
		return StakeLatestStrategy
	} else if strings.EqualFold(c.StrategyImpl, "MakeNodes") {
		return MakeNodesStrategy
	} else {
		return UnknownStrategy
	}
}

type Wallet struct {
	Fireblocks WalletFireblocks `koanf:"fireblocks"`
	Local      WalletLocal      `koanf:"local"`
}

type WalletFireblocks struct {
	APIKey               string     `koanf:"api-key,omitempty"`
	AssetId              string     `koanf:"asset-id,omitempty"`
	BaseURL              string     `koanf:"base-url,omitempty"`
	DisableHandlePending bool       `koanf:"disable-handle-pending"`
	ExternalWallets      string     `koanf:"external-wallets"`
	FeedSigner           FeedSigner `koanf:"feed-signer"`
	InternalWallets      string     `koanf:"internal-wallets"`
	SourceAddress        string     `koanf:"source-address,omitempty"`
	SourceId             string     `koanf:"source-id,omitempty"`
	SourceType           string     `koanf:"source-type,omitempty"`
	SSLKey               string     `koanf:"ssl-key,omitempty"`
	SSLKeyPassword       string     `koanf:"ssl-key-password,omitempty"`
	UseFireblocksFees    bool       `koanf:"use-fireblocks-fees"`
}

type FeedSigner struct {
	Pathname     string `koanf:"pathname"`
	PasswordImpl string `koanf:"password"`
	PrivateKey   string `koanf:"private-key"`
}

func (f *FeedSigner) Password() *string {
	if f.PasswordImpl == PASSWORD_NOT_SET {
		return nil
	}
	return &f.PasswordImpl
}

type WalletLocal struct {
	OnlyCreateKey bool   `koanf:"only-create-key"`
	Pathname      string `koanf:"pathname"`
	PasswordImpl  string `koanf:"password"`
	PrivateKey    string `koanf:"private-key"`
}

func (w WalletLocal) Password() *string {
	if w.PasswordImpl == PASSWORD_NOT_SET {
		return nil
	}
	return &w.PasswordImpl
}

type Log struct {
	RPC  string `koanf:"rpc"`
	Core string `koanf:"core"`
}

type Metrics struct {
	Addr string `koanf:"addr"`
	Port string `koanf:"port"`
}

type Config struct {
	BridgeUtilsAddress string      `koanf:"bridge-utils-address"`
	Conf               Conf        `koanf:"conf"`
	Core               Core        `koanf:"core"`
	Feed               Feed        `koanf:"feed"`
	GasPrice           float64     `koanf:"gas-price"`
	Healthcheck        Healthcheck `koanf:"healthcheck"`
	L1                 struct {
		ChainID int    `koanf:"chain-id"`
		URL     string `koanf:"url"`
	} `koanf:"l1"`
	Log           Log        `koanf:"log"`
	Node          Node       `koanf:"node"`
	Persistent    Persistent `koanf:"persistent"`
	PProfEnable   bool       `koanf:"pprof-enable"`
	Rollup        Rollup     `koanf:"rollup"`
	Validator     Validator  `koanf:"validator"`
	WaitToCatchUp bool       `koanf:"wait-to-catch-up"`
	Wallet        Wallet     `koanf:"wallet"`

	// The following field needs to be top level for compatibility with the underlying go-ethereum lib
	Metrics       bool    `koanf:"metrics"`
	MetricsServer Metrics `koanf:"metrics-server"`
}

// DefaultCoreSettingsNoMaxExecution is useful in unit tests
func DefaultCoreSettingsNoMaxExecution() *Core {
	return &Core{
		Cache: CoreCache{
			BasicInterval: 100,
			BasicSize:     1000,
			LRUSize:       1000,
			TimedExpire:   20 * time.Minute,
		},
		CheckpointGasFrequency:    1_000_000,
		CheckpointLoadGasCost:     1_000_000,
		CheckpointLoadGasFactor:   4,
		CheckpointMaxExecutionGas: 0,
		CheckpointPruningMode:     "default",
		MessageProcessCount:       10,
	}
}

// DefaultCoreSettingsMaxExecution is useful in unit tests
func DefaultCoreSettingsMaxExecution() *Core {
	return &Core{
		Cache: CoreCache{
			BasicInterval: 100,
			BasicSize:     1000,
			LRUSize:       1000,
			TimedExpire:   20 * time.Minute,
		},
		CheckpointGasFrequency:    1_000_000,
		CheckpointLoadGasCost:     1_000_000,
		CheckpointLoadGasFactor:   4,
		CheckpointMaxExecutionGas: 1_000_000_000,
		CheckpointPruningMode:     "default",
		MessageProcessCount:       10,
	}
}

// DefaultNodeSettings is useful in unit tests
func DefaultNodeSettings() *Node {
	return &Node{
		Cache: NodeCache{
			AllowSlowLookup: true,
			LRUSize:         1000,
			TimedExpire:     20 * time.Minute,
		},
		InboxReader: InboxReader{
			DelayBlocks: 4,
			Paranoid:    false,
		},
		LogProcessCount: 100,
		LogIdleSleep:    10 * time.Millisecond, // 10 for dev, 100 for server
		Sequencer: Sequencer{
			CreateBatchBlockInterval:   40,
			DelayedMessagesTargetDelay: 1,
			MaxBatchGasCost:            2_000_000,
			GasRefunderAddress:         "",
		},
	}
}

func (c *Config) GetDatabasePath() string {
	return path.Join(c.Persistent.Chain, "db")
}

func ParseCLI(ctx context.Context) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	AddForwarderTarget(f)

	return ParseNonRelay(ctx, f, "cli-wallet", 0)
}

func ParseDBTool() (*Config, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	AddPersistent(f)
	AddCore(f, 0)

	k, err := beginCommonParse(f)
	if err != nil {
		return nil, err
	}

	out, wallet, err := endCommonParse(k)
	if err != nil {
		return nil, err
	}

	err = resolveDirectoryNames(out, wallet)

	return out, nil
}

func AddL1PostingStrategyOptions(f *flag.FlagSet, prefix string) {
	f.Float64(prefix+"l1-posting-strategy.high-gas-threshold", 150, "gwei threshold at which to consider gas price high and delay batch posting")
	f.Int64(prefix+"l1-posting-strategy.high-gas-delay-blocks", 270, "wait up to this many more blocks when gas costs are high")
}

func ParseNode(ctx context.Context) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f := flag.NewFlagSet("", flag.ContinueOnError)

	AddFeedOutputOptions(f)
	AddForwarderTarget(f)
	AddL1PostingStrategyOptions(f, "node.sequencer.")
	AddL1PostingStrategyOptions(f, "validator.")

	f.String("validator.strategy", "StakeLatest", "strategy for validator to use")
	f.String("validator.utils-address", "", "validator utilities address")
	f.Duration("validator.staker-delay", 60*time.Second, "delay between updating stake")
	f.String("validator.wallet-factory-address", "", "strategy for validator to use")
	f.Bool("validator.dont-challenge", false, "don't challenge any other validators' assertions")
	f.String("validator.withdraw-destination", "", "the address to withdraw funds to (defaults to the wallet address)")

	f.String("node.aggregator.inbox-address", "", "address of the inbox contract")
	f.Int("node.aggregator.max-batch-time", 10, "max-batch-time=NumSeconds")
	f.Bool("node.aggregator.stateful", false, "enable pending state tracking")

	f.Bool("node.cache.allow-slow-lookup", false, "load L2 block from disk if not in memory cache")
	f.Int("node.cache.lru-size", 1000, "number of recently used L2 blocks to hold in lru memory cache")
	f.Int("node.cache.block-info-lru-size", 100_000, "number of recently used L2 block info to hold in lru memory cache")
	f.Duration("node.cache.timed-expire", 20*time.Minute, "length of time to hold L2 blocks in timed memory cache")

	f.Uint64("node.chain-id", 42161, "chain id of the arbitrum chain")

	f.String("node.forwarder.submitter-address", "", "address of the node that will submit your transaction to the chain")
	f.String("node.forwarder.rpc-mode", "full", "RPC mode: either full, non-mutating (no eth_sendRawTransaction), or forwarding-only (only requests forwarded upstream are permitted)")

	f.Int64("node.inbox-reader.delay-blocks", 4, "number of L1 blocks to wait for confirmation before updating L2 state")
	f.Bool("node.inbox-reader.paranoid", false, "if enabled, check for reorgs before searching for messages")
	f.Duration("node.inbox-reader.sequencer-signature-expiry", 10*time.Minute, "length of time between verifying sequencer feed signing address on-chain")

	f.Duration("node.log-idle-sleep", 100*time.Millisecond, "milliseconds for log reader to sleep between reading logs")
	f.Int("node.log-process-count", 100, "maximum number of logs to process at a time")

	f.String("node.rpc.addr", "0.0.0.0", "RPC address")
	f.Int("node.rpc.port", 8547, "RPC port")
	f.String("node.rpc.path", "/", "RPC path")
	f.Bool("node.rpc.enable-l1-calls", false, "If RPC calls which query the L1 node indirectly should be allowed")
	f.Bool("node.rpc.tracing.enable", false, "enable tracing api")
	f.String("node.rpc.tracing.namespace", "arbtrace", "rpc namespace for tracing api")
	f.Uint64("node.rpc.max-call-gas", 5000000, "Max computational arbgas limit when processing eth_call and eth_estimateGas")
	f.Bool("node.rpc.enable-devops-stubs", false, "Enable fake versions of eth_syncing and eth_netPeers")

	f.Int64("node.sequencer.create-batch-block-interval", 270, "block interval at which to create new batches")
	f.Int64("node.sequencer.continue-batch-posting-block-interval", 2, "block interval to post the next batch after posting a partial one")
	f.Int64("node.sequencer.delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	f.String("node.sequencer.lockout.redis", "", "sequencer lockout redis instance URL")
	f.String("node.sequencer.lockout.self-rpc-url", "", "own RPC URL for other sequencers to failover to")
	f.Int64("node.sequencer.max-batch-gas-cost", 2_000_000, "max L1 batch gas cost to post before splitting it up into multiple batches")
	f.String("node.sequencer.gas-refunder-address", "", "address of the L1 gas refunder contract (optional)")
	f.Uint64("node.sequencer.gas-refunder-extra-gas", 50_000, "amount of extra gas to supply for the gas refunder operation")
	f.Bool("node.sequencer.dangerous.reorg-out-huge-messages", false, "erase any huge messages in database that cannot be published (DANGEROUS)")
	f.Bool("node.sequencer.dangerous.publish-batches-without-lockout", false, "continue publishing batches (but not sequencing) without the lockout (DANGEROUS)")
	f.Bool("node.sequencer.dangerous.rewrite-sequencer-address", false, "reorganize to rewrite the sequencer address if it's not the loaded wallet (DANGEROUS)")
	f.Bool("node.sequencer.dangerous.disable-batch-posting", false, "disable posting batches to L1 (DANGEROUS)")
	f.Bool("node.sequencer.dangerous.disable-delayed-message-sequencing", false, "disable sequencing delayed messages (DANGEROUS)")
	f.Bool("node.sequencer.debug-timing", false, "log elapsed time throughout core sequencing loop")

	f.String("node.type", "forwarder", "forwarder, aggregator or sequencer")

	f.String("node.ws.addr", "0.0.0.0", "websocket address")
	f.Int("node.ws.port", 8548, "websocket port")
	f.String("node.ws.path", "/", "websocket path")

	return ParseNonRelay(ctx, f, "rpc-wallet", 250_000_000)
}

func ParseNonRelay(ctx context.Context, f *flag.FlagSet, defaultWalletPathname string, maxExecutionGas int) (*Config, *Wallet, *ethutils.RPCEthClient, *big.Int, error) {
	f.String("bridge-utils-address", "", "bridgeutils contract address")

	f.Float64("gas-price", 0, "float of gas price to use in gwei (0 = use L1 node's recommended value)")

	f.String("l1.url", "", "layer 1 ethereum node RPC URL")
	f.Uint64("l1.chain-id", 0, "if set other than 0, will be used to validate database and L1 connection")

	f.String("rollup.address", "", "layer 2 rollup contract address")
	f.String("rollup.machine.filename", "", "file to load machine from")

	f.Bool("wallet.local.only-create-key", false, "create new wallet and exit")
	f.String("wallet.local.pathname", defaultWalletPathname, "path to store wallet in")
	f.String("wallet.local.password", PASSWORD_NOT_SET, "password for wallet")
	f.String("wallet.local.private-key", "", "wallet private key string")

	f.String("wallet.fireblocks.feed-signer.pathname", "feed-signer-wallet", "path to store feed-signer wallet in")
	f.String("wallet.fireblocks.feed-signer.password", PASSWORD_NOT_SET, "password for feed-signer wallet")
	f.String("wallet.fireblocks.feed-signer.private-key", "", "wallet feed-signer private key string")

	f.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")

	AddCore(f, maxExecutionGas)
	AddHealthcheckOptions(f)
	AddPersistent(f)

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
		return nil, nil, nil, nil, errors.Wrapf(err, "error connecting to ethereum L1 node: %s", l1URL)
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
	logger.Info().Str("l1url", l1URL).Str("l1chainid", l1ChainId.String()).Msg("connected to l1 chain")

	rollupAddress := k.String("rollup.address")
	if len(rollupAddress) != 0 {
		logger.Info().Str("rollup", rollupAddress).Msg("using custom rollup address")
	} else {
		if l1ChainId.Cmp(big.NewInt(1)) == 0 {
			err := k.Load(confmap.Provider(map[string]interface{}{
				"bridge-utils-address":             "0x84efa170dc6d521495d7942e372b8e4b2fb918ec",
				"feed.input.url":                   []string{"wss://arb1.arbitrum.io/feed"},
				"node.aggregator.inbox-address":    "0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f",
				"node.chain-id":                    "42161",
				"node.forwarder.target":            "https://arb1.arbitrum.io/rpc",
				"persistent.chain":                 "mainnet",
				"rollup.address":                   "0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A",
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
				"node.aggregator.inbox-address":    "0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e",
				"node.chain-id":                    "421611",
				"node.forwarder.target":            "https://rinkeby.arbitrum.io/rpc",
				"persistent.chain":                 "rinkeby",
				"rollup.address":                   "0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c",
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

	if err := applyOverrides(f, k); err != nil {
		return nil, nil, nil, nil, err
	}

	out, wallet, err := endCommonParse(k)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Fixup directories
	err = resolveDirectoryNames(out, wallet)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if len(out.Rollup.Machine.Filename) == 0 {
		// Machine not provided, so use default
		out.Rollup.Machine.Filename = path.Join(out.Persistent.Chain, "arbos.mexe")
	}

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

	if out.L1.ChainID != 0 && l1ChainId.Int64() != int64(out.L1.ChainID) {
		logger.
			Error().
			Int("expected-chainid", out.L1.ChainID).
			Int64("l1-chainid", l1ChainId.Int64()).
			Msg("unexpected chain id")
		return nil, nil, nil, nil, fmt.Errorf("expected chain id %v but l1 node has chain id %v", out.L1.ChainID, l1ChainId)
	}

	if out.Node.Cache.AllowSlowLookup {
		// Force unlimited execution
		out.Core.CheckpointMaxExecutionGas = 0

		// Never prune checkpoints
		if out.Core.CheckpointPruningMode != "off" {
			logger.Warn().Msg("disabling checkpoint pruning because allow-slow-lookup enabled")
		}
		out.Core.CheckpointPruningMode = "off"
	}

	if (out.Core.CheckpointPruningMode != "on") &&
		(out.Core.CheckpointPruningMode != "default") {
		return nil, nil, nil, nil,
			fmt.Errorf("value '%v' for core.checkpoint-pruning-mode is not 'on', 'off', or 'default'", out.Core.CheckpointPruningMode)
	}

	if out.Node.Type() == SequencerNodeType && !out.Core.Cache.Last {
		logger.Info().Msg("enabling last machine cache for sequencer")
		out.Core.Cache.Last = true
	}

	return out, wallet, l1Client, l1ChainId, nil
}

func resolveDirectoryNames(out *Config, wallet *Wallet) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "Unable to read users home directory")
	}

	// Make persistent storage directory relative to home directory if not already absolute
	if !filepath.IsAbs(out.Persistent.GlobalConfig) {
		out.Persistent.GlobalConfig = path.Join(homeDir, out.Persistent.GlobalConfig)
	}
	err = os.MkdirAll(out.Persistent.GlobalConfig, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Unable to create global configuration directory")
	}

	// Make chain directory relative to persistent storage directory if not already absolute
	if !filepath.IsAbs(out.Persistent.Chain) {
		out.Persistent.Chain = path.Join(out.Persistent.GlobalConfig, out.Persistent.Chain)
	}
	err = os.MkdirAll(out.Persistent.Chain, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "Unable to create chain directory")
	}
	if DatabaseInDirectory(out.Persistent.Chain) {
		return errors.New("Database in --persistent.chain directory, try specifying parent directory")
	}

	// Make rocksdb backup directory relative to persistent storage directory if not already absolute
	if !filepath.IsAbs(out.Core.Database.SavePath) {
		out.Core.Database.SavePath = path.Join(out.Persistent.Chain, out.Core.Database.SavePath)
	}

	// Make machine relative to storage directory if not already absolute
	if !filepath.IsAbs(out.Rollup.Machine.Filename) {
		out.Rollup.Machine.Filename = path.Join(out.Persistent.GlobalConfig, out.Rollup.Machine.Filename)
	}

	// Make wallet directories relative to chain directory if not already absolute
	if !filepath.IsAbs(wallet.Local.Pathname) {
		wallet.Local.Pathname = path.Join(out.Persistent.Chain, wallet.Local.Pathname)
	}
	if !filepath.IsAbs(wallet.Fireblocks.FeedSigner.Pathname) {
		wallet.Fireblocks.FeedSigner.Pathname = path.Join(out.Persistent.Chain, wallet.Fireblocks.FeedSigner.Pathname)
	}

	return nil
}

func ParseRelay() (*Config, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	AddFeedOutputOptions(f)

	k, err := beginCommonParse(f)
	if err != nil {
		return nil, err
	}

	out, _, err := endCommonParse(k)
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

func AddForwarderTarget(f *flag.FlagSet) {
	f.String("node.forwarder.target", "", "url of another node to send transactions through")
}

func AddCore(f *flag.FlagSet, maxExecutionGas int) {
	f.Bool("core.cache.last", false, "whether to always cache the machine from last block")
	f.Int("core.cache.basic-interval", 100_000_000, "amount of gas to wait between saving to basic cache")
	f.Int("core.cache.basic-size", 100, "number of basic cache entries to save")
	f.Bool("core.cache.disable", false, "disable saving to cache while in core thread")
	f.Int("core.cache.lru-size", 1000, "number of recently used L2 blocks to hold in lru memory cache")
	f.Bool("core.cache.seed-on-startup", false, "seed cache on startup by re-executing timed-expire worth of history")
	f.Duration("core.cache.timed-expire", 20*time.Minute, "length of time to hold L2 blocks in arbcore timed memory cache")

	f.Int("core.add-messages-max-failure-count", 10, "number of add messages failures before exiting program")
	f.Int("core.checkpoint-gas-frequency", 1_000_000_000, "amount of gas between saving checkpoints")
	f.Int("core.checkpoint-load-gas-cost", 250_000_000, "running machine for given gas takes same amount of time as loading database entry")
	f.Int("core.checkpoint-load-gas-factor", 4, "factor to weight difference in database checkpoint vs cache checkpoint")
	f.Int("core.checkpoint-max-execution-gas", maxExecutionGas, "maximum amount of gas any given checkpoint is allowed to execute")
	f.Int("core.checkpoint-max-to-prune", 2, "number of checkpoints to delete at a time, 0 for no limit")
	f.Bool("core.checkpoint-prune-on-startup", false, "perform full database pruning on startup")
	f.String("core.checkpoint-pruning-mode", "default", "Prune old checkpoints: 'on', 'off', or 'default'")

	f.Bool("core.database.compact", false, "perform database compaction")
	f.Bool("core.database.exit-after", false, "exit after loading or manipulating database")
	f.Bool("core.database.metadata", false, "just print database metadata and exit")
	f.Duration("core.database.save-interval", 0, "duration between saving database backups, 0 to disable")
	f.Bool("core.database.save-on-startup", false, "save database backup on start")
	f.String("core.database.save-path", "db_checkpoints", "path to save database backups in")

	f.Bool("core.debug", false, "print extra debug messages in arbcore")
	f.Bool("core.debug-timing", false, "print extra debug timing messages in arbcore")

	f.Duration("core.idle-sleep", 5*time.Millisecond, "how long core thread should sleep when idle")

	f.Bool("core.lazy-load-core-machine", false, "if the core machine should be loaded as it's run")
	f.Bool("core.lazy-load-archive-queries", true, "if the archive queries should be loaded as they're run")

	f.Int("core.message-process-count", 100, "maximum number of messages to process at a time")

	f.Int("core.test.load-count", 0, "number of snapshots to load from database for profile test, zero to disable")
	f.Int("core.test.reorg-to.l1-block", 0, "reorg to snapshot with given L1 block or before, zero to disable")
	f.Int("core.test.reorg-to.l2-block", 0, "reorg to snapshot with given L2 block or before, zero to disable")
	f.Int("core.test.reorg-to.log", 0, "reorg to snapshot with given log or before, zero to disable")
	f.Int("core.test.reorg-to.message", 0, "reorg to snapshot with given message or before, zero to disable")
	f.Bool("core.test.reset-all-except-inbox", false, "remove all database info except for inbox")
	f.Int("core.test.run-until", 0, "run until gas is reached for profile test, zero to disable")

}

func AddPersistent(f *flag.FlagSet) {
	f.String("persistent.global-config", ".arbitrum", "location global configuration is located")
	f.String("persistent.chain", "", "path that chain specific state is located")
}

func AddHealthcheckOptions(f *flag.FlagSet) {
	f.Bool("healthcheck.enable", false, "enable healthcheck endpoint")
	f.Bool("healthcheck.sequencer", false, "enable checking the health of the sequencer")
	f.Bool("healthcheck.l1-node", false, "enable checking the health of the L1 node")
	f.Bool("healthcheck.metrics", false, "enable prometheus endpoint")
	f.String("healthcheck.metrics-prefix", "", "prepend the specified prefix to the exported metrics names")
	f.String("healthcheck.addr", "", "address to bind the healthcheck endpoint to")
	f.Int("healthcheck.port", 0, "port to bind the healthcheck endpoint to")
}

func beginCommonParse(f *flag.FlagSet) (*koanf.Koanf, error) {
	f.Bool("conf.dump", false, "print out currently active configuration file")
	f.String("conf.env-prefix", "", "environment variables with given prefix will be loaded as configuration values")
	f.String("conf.file", "", "name of configuration file")
	f.String("conf.s3.access-key", "", "S3 access key")
	f.String("conf.s3.secret-key", "", "S3 secret key")
	f.String("conf.s3.region", "", "S3 region")
	f.String("conf.s3.bucket", "", "S3 bucket")
	f.String("conf.s3.object-key", "", "S3 object key")
	f.String("conf.string", "", "configuration as JSON string")

	f.Duration("feed.input.timeout", 20*time.Second, "duration to wait before timing out connection to server")
	f.StringSlice("feed.input.url", []string{}, "URL of sequencer feed source")

	f.Bool("metrics", false, "enable metrics")
	f.String("metrics-server.addr", "127.0.0.1", "metrics server address")
	f.String("metrics-server.port", "6070", "metrics server address")

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

	// Initial application of command line parameters and environment variables so other methods can be applied
	if err := applyOverrides(f, k); err != nil {
		return nil, err
	}

	return k, nil
}

func applyOverrides(f *flag.FlagSet, k *koanf.Koanf) error {
	// Apply command line options and environment variables
	if err := applyOverrideOverrides(f, k); err != nil {
		return err
	}

	// Load configuration file from S3 if setup
	if len(k.String("conf.s3.secret-key")) != 0 {
		if err := loadS3Variables(k); err != nil {
			return errors.Wrap(err, "error loading S3 settings")
		}

		if err := applyOverrideOverrides(f, k); err != nil {
			return err
		}
	}

	// Local config file overrides S3 config file
	configFile := k.String("conf.file")
	if len(configFile) > 0 {
		if err := k.Load(file.Provider(configFile), json.Parser()); err != nil {
			return errors.Wrap(err, "error loading local config file")
		}

		if err := applyOverrideOverrides(f, k); err != nil {
			return err
		}
	}

	return nil
}

// applyOverrideOverrides for configuration values that need to be re-applied for each configuration item applied
func applyOverrideOverrides(f *flag.FlagSet, k *koanf.Koanf) error {
	// Command line overrides config file or config string
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return errors.Wrap(err, "error loading command line config")
	}

	// Config string overrides any config file
	configString := k.String("conf.string")
	if len(configString) > 0 {
		if err := k.Load(rawbytes.Provider([]byte(configString)), json.Parser()); err != nil {
			return errors.Wrap(err, "error loading config string config")
		}

		// Command line overrides config file or config string
		if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
			return errors.Wrap(err, "error loading command line config")
		}
	}

	// Environment variables overrides config files or command line options
	if err := loadEnvironmentVariables(k); err != nil {
		return errors.Wrap(err, "error loading environment variables")
	}

	return nil
}

func loadEnvironmentVariables(k *koanf.Koanf) error {
	envPrefix := k.String("conf.env-prefix")
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

func loadS3Variables(k *koanf.Koanf) error {
	return k.Load(s3.Provider(s3.Config{
		AccessKey: k.String("conf.s3.access-key"),
		SecretKey: k.String("conf.s3.secret-key"),
		Region:    k.String("conf.s3.region"),
		Bucket:    k.String("conf.s3.bucket"),
		ObjectKey: k.String("conf.s3.object-key"),
	}), nil)
}

func endCommonParse(k *koanf.Koanf) (*Config, *Wallet, error) {
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

	if len(out.Wallet.Fireblocks.SSLKey) != 0 {
		if len(out.Wallet.Fireblocks.APIKey) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.api-key")
		}
		if len(out.Wallet.Fireblocks.BaseURL) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.base-url")
		}
		if len(out.Wallet.Fireblocks.SSLKey) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.ssl-key")
		}
		if len(out.Wallet.Fireblocks.SourceAddress) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.source-address")
		}
		if len(out.Wallet.Fireblocks.SourceId) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.source-id")
		}
		if len(out.Wallet.Fireblocks.SourceType) == 0 {
			return nil, nil, errors.New("fireblocks configured but missing fireblocks.source-type")
		}

		out.Wallet.Fireblocks.SSLKey = strings.Replace(out.Wallet.Fireblocks.SSLKey, "\\n", "\n", -1)
	}

	if out.Conf.Dump {
		// Print out current configuration

		// Don't keep printing configuration file and don't print wallet passwords
		err := k.Load(confmap.Provider(map[string]interface{}{
			"conf.dump":                                 false,
			"wallet.fireblocks.feed-signer.password":    "",
			"wallet.fireblocks.feed-signer.private-key": "",
			"wallet.fireblocks.ssl-key":                 "",
			"wallet.fireblocks.ssl-key-password":        "",
			"wallet.local.password":                     "",
			"wallet.local.private-key":                  "",
		}, "."), nil)

		c, err := k.Marshal(json.Parser())
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to marshal config file to JSON")
		}

		fmt.Println(string(c))
		os.Exit(1)
	}

	// Don't pass around wallet contents with normal configuration
	wallet := out.Wallet
	out.Wallet = Wallet{}

	return &out, &wallet, nil
}

func UnmarshalMap(marshalled string) map[string]string {
	unmarshalled := make(map[string]string)
	if len(marshalled) == 0 {
		return unmarshalled
	}
	items := strings.Split(marshalled, ",")
	for _, pair := range items {
		item := strings.Split(pair, ":")
		unmarshalled[item[0]] = item[1]
	}

	return unmarshalled
}

func DatabaseInDirectory(path string) bool {
	// Consider database present if file `CURRENT` in directory
	_, err := os.Stat(path + "/CURRENT")

	return err == nil
}
