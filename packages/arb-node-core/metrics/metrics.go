package metrics

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/prometheus/client_golang/prometheus"
)

//Metrics configuration struct
type MetricsConfig struct {

	// Holds the method call counter used by multiple RPCs
	MethodCallCounter *prometheus.CounterVec

	// Prometheus Registerer to register histograms on
	Registerer prometheus.Registerer
	// Prometheus Registery to handle the exposed service
	Registry *prometheus.Registry
}

func NewMetricsConfig(prefix *string) *MetricsConfig {
	registry := prometheus.NewRegistry()
	var registerer prometheus.Registerer = registry
	if prefix != nil && *prefix != "" {
		registerer = prometheus.WrapRegistererWithPrefix(*prefix, registry)
	}
	methodCallCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "arbitrum",
			Subsystem: "rpc",
			Name:      "call",
		},
		[]string{"method", "success"},
	)
	return &MetricsConfig{
		MethodCallCounter: methodCallCounter,
		Registry:          registry,
		Registerer:        registerer,
	}
}

func (m *MetricsConfig) RegisterMetrics(collectors ...prometheus.Collector) {
	m.Registerer.MustRegister(collectors...)
}

/// Register metrics orthagonal to Arbitrum (Go subsystem metrics, Process Metrics, etc).
func (m *MetricsConfig) RegisterSystemMetrics() {
	m.Registerer.MustRegister(
		prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
	)
}

/// Register metrics that are stored statically, i.e. counters and "push" gauges
func (m *MetricsConfig) RegisterStaticMetrics() {
	m.Registerer.MustRegister(
		m.MethodCallCounter,
		cmachine.GasCounter, cmachine.StepsCounter,
		monitor.BatchesCounter, monitor.EthHeightGauge, monitor.DelayedCounter, monitor.MessageGauge)
}

func RegisterNodeStoreMetrics(nodeStore machine.NodeStore, metrics *MetricsConfig) {
	metrics.RegisterMetrics(
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Namespace: "arbitrum",
				Subsystem: "avm",
				Name:      "block_height",
				Help:      "Current height of the Arbitrum chain",
			},
			func() float64 {
				count, _ := nodeStore.BlockCount()
				return float64(count)
			}))
}
