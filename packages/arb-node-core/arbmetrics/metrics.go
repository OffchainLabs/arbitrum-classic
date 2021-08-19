package arbmetrics

import (
	"time"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/exp"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type Registerable interface {
	Register(metrics.Registry) error
}

// Config Metrics configuration struct
type Config struct {
	// Prometheus Registry to handle the exposed service and register metrics
	Registry metrics.Registry
}

func NewMetricsConfig(config configuration.MetricsServer) *Config {
	if config.Enable {
		exp.Setup(config.Endpoint.Addr + ":" + config.Endpoint.Port)
	}

	registry := metrics.DefaultRegistry
	if config.Prefix != "" {
		registry = metrics.NewPrefixedChildRegistry(registry, config.Prefix)
	}
	registry = metrics.NewPrefixedChildRegistry(registry, "arbitrum/")

	// Registers metrics orthogonal to Arbitrum (Go subsystem metrics, Process InboxMetrics, etc.).
	go metrics.CollectProcessMetrics(3 * time.Second)

	return &Config{
		Registry: registry,
	}
}

func (c *Config) Register(r Registerable) error {
	return r.Register(c.Registry)
}

func (c *Config) RegisterWithPrefix(r Registerable, prefix string) error {
	m := metrics.NewPrefixedChildRegistry(c.Registry, prefix)
	return r.Register(m)
}
