package metrics

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/exp"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = arblog.Logger.With().Str("component", "metrics").Logger()

// Config Metrics configuration struct
type Config struct {
	// Prometheus Registry to handle the exposed service and register metrics
	Registry metrics.Registry
}

func NewMetricsConfig(config configuration.Metrics, prefix *string) *Config {
	if metrics.Enabled {
		exp.Setup(config.Addr + ":" + config.Port)
	}

	registry := metrics.DefaultRegistry
	if prefix != nil && *prefix != "" {
		registry = metrics.NewPrefixedChildRegistry(registry, *prefix)
	}

	// Registers metrics orthogonal to Arbitrum (Go subsystem metrics, Process Metrics, etc.).
	go metrics.CollectProcessMetrics(3 * time.Second)

	return &Config{
		Registry: registry,
	}
}

func (m *Config) RegisterNodeStoreMetrics(nodeStore machine.NodeStore) {
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/avm/block_height",
		m.Registry,
		func() int64 {
			count, _ := nodeStore.BlockCount()
			return int64(count)
		},
	)
}

func (m *Config) RegisterArbCoreMetrics(arbCore core.ArbCore) {
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/message_count",
		m.Registry,
		func() int64 {
			messageCount, err := arbCore.GetMessageCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting message count")
				return 0
			}
			return messageCount.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/delayed_message_count",
		m.Registry,
		func() int64 {
			messageCount, err := arbCore.GetDelayedMessageCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting delayed message count")
				return 0
			}
			return messageCount.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/sequenced_delayed_message_count",
		m.Registry,
		func() int64 {
			messageCount, err := arbCore.GetTotalDelayedMessagesSequenced()
			if err != nil {
				logger.Error().Err(err).Msg("error getting sequenced delayed message count")
				return 0
			}
			return messageCount.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/messages_read_count",
		m.Registry,
		func() int64 {
			return arbCore.MachineMessagesRead().Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/log_count",
		m.Registry,
		func() int64 {
			logCount, err := arbCore.GetLogCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting log count")
				return 0
			}
			return logCount.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/send_count",
		m.Registry,
		func() int64 {
			sendCount, err := arbCore.GetSendCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting send count")
				return 0
			}
			return sendCount.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/total_gas_processed",
		m.Registry,
		func() int64 {
			gas, err := arbCore.GetLastMachineTotalGas()
			if err != nil {
				logger.Error().Err(err).Msg("error getting gas used")
				return 0
			}
			return gas.Int64()
		},
	)
	metrics.NewRegisteredFunctionalGauge(
		"arbitrum/core/logs_cursor_position",
		m.Registry,
		func() int64 {
			pos, err := arbCore.LogsCursorPosition(big.NewInt(0))
			if err != nil {
				logger.Error().Err(err).Msg("error getting logs cursor position")
				return 0
			}
			return pos.Int64()
		},
	)
}
