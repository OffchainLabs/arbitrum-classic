package healthcheck

import (
	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/ethereum/go-ethereum/metrics"
)

// MetricsListener reports metrics on each check registration, start and completion event (as gosundheit.CheckListener)
// This listener all reports metrics for the entire service health (as gosundheit.HealthListener)
type MetricsListener struct {
	registry metrics.Registry
}

func NewMetricsListener(registry metrics.Registry) *MetricsListener {
	return &MetricsListener{
		registry: registry,
	}
}

func (c *MetricsListener) OnResultsUpdated(results map[string]gosundheit.Result) {
	for name, res := range results {
		statusGauge := metrics.GetOrRegisterGauge(name+"/healthy", c.registry)
		contiguousFailureGauge := metrics.GetOrRegisterGauge(name+"/contiguous_failures", c.registry)
		if res.IsHealthy() {
			healthyCounter := metrics.GetOrRegisterCounter(name+"/success_total", c.registry)
			statusGauge.Update(1)
			healthyCounter.Inc(1)
		} else {
			unhealthyCounter := metrics.GetOrRegisterCounter(name+"/fail_total", c.registry)
			statusGauge.Update(0)
			unhealthyCounter.Inc(1)
		}
		contiguousFailureGauge.Update(res.ContiguousFailures)
	}
}
