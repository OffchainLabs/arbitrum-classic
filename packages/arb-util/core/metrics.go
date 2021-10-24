package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/metrics"
)

type Metrics struct {
	MessageCount                 metrics.Gauge
	DelayedMessageCount          metrics.Gauge
	SequencedDelayedMessageCount metrics.Gauge
	MachineMessagesReadCount     metrics.Gauge
	LogCount                     metrics.Gauge
	SendCount                    metrics.Gauge
	TotalGasProcessed            metrics.Gauge
	LogsCursorPosition           metrics.Gauge
}

func NewArbCoreMetrics(arbCore ArbCore) *Metrics {
	return &Metrics{
		MessageCount: metrics.NewFunctionalGauge(func() int64 {
			messageCount, err := arbCore.GetMessageCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting message count")
				return 0
			}
			return messageCount.Int64()
		}),
		DelayedMessageCount: metrics.NewFunctionalGauge(func() int64 {
			messageCount, err := arbCore.GetDelayedMessageCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting delayed message count")
				return 0
			}
			return messageCount.Int64()
		}),
		SequencedDelayedMessageCount: metrics.NewFunctionalGauge(func() int64 {
			messageCount, err := arbCore.GetTotalDelayedMessagesSequenced()
			if err != nil {
				logger.Error().Err(err).Msg("error getting sequenced delayed message count")
				return 0
			}
			return messageCount.Int64()
		}),
		MachineMessagesReadCount: metrics.NewFunctionalGauge(func() int64 {
			return arbCore.MachineMessagesRead().Int64()
		}),
		LogCount: metrics.NewFunctionalGauge(func() int64 {
			logCount, err := arbCore.GetLogCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting log count")
				return 0
			}
			return logCount.Int64()
		}),
		SendCount: metrics.NewFunctionalGauge(func() int64 {
			sendCount, err := arbCore.GetSendCount()
			if err != nil {
				logger.Error().Err(err).Msg("error getting send count")
				return 0
			}
			return sendCount.Int64()
		}),
		TotalGasProcessed: metrics.NewFunctionalGauge(func() int64 {
			gas, err := arbCore.GetLastMachineTotalGas()
			if err != nil {
				logger.Error().Err(err).Msg("error getting gas used")
				return 0
			}
			return gas.Int64()
		}),
		LogsCursorPosition: metrics.NewFunctionalGauge(func() int64 {
			pos, err := arbCore.LogsCursorPosition(big.NewInt(0))
			if err != nil {
				logger.Error().Err(err).Msg("error getting logs cursor position")
				return 0
			}
			return pos.Int64()
		}),
	}
}

func (m *Metrics) Register(r metrics.Registry) error {
	if err := r.Register("message_count", m.MessageCount); err != nil {
		return err
	}
	if err := r.Register("delayed_message_count", m.DelayedMessageCount); err != nil {
		return err
	}
	if err := r.Register("sequenced_delayed_message_count", m.SequencedDelayedMessageCount); err != nil {
		return err
	}
	if err := r.Register("messages_read_count", m.MachineMessagesReadCount); err != nil {
		return err
	}
	if err := r.Register("log_count", m.LogCount); err != nil {
		return err
	}
	if err := r.Register("send_count", m.SendCount); err != nil {
		return err
	}
	if err := r.Register("total_gas_processed", m.TotalGasProcessed); err != nil {
		return err
	}
	if err := r.Register("logs_cursor_position", m.LogsCursorPosition); err != nil {
		return err
	}
	return nil
}
