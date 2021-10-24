package healthcheck

import (
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/ethereum/go-ethereum/metrics"
)

type Health struct {
	impl gosundheit.Health
}

func New(registry metrics.Registry) gosundheit.Health {
	return gosundheit.New(gosundheit.WithHealthListeners(NewMetricsListener(registry)), gosundheit.ExecutionPeriod(10*time.Second))
}

func (h *Health) RegisterCheck(check gosundheit.Check, opts ...gosundheit.CheckOption) error {
	if metrics.Enabled {
		allOpts := opts
		allOpts = append(allOpts)
		return h.impl.RegisterCheck(check, allOpts...)
	}
	return nil
}

func (h *Health) Deregister(name string) {
	if metrics.Enabled {
		h.impl.Deregister(name)
	}
}

func (h *Health) Results() (results map[string]gosundheit.Result, healthy bool) {
	if metrics.Enabled {
		return h.impl.Results()
	}
	return make(map[string]gosundheit.Result), true
}

func (h *Health) IsHealthy() bool {
	if metrics.Enabled {
		return h.impl.IsHealthy()
	}
	return true
}

func (h *Health) DeregisterAll() {
	if metrics.Enabled {
		h.impl.DeregisterAll()
	}
}
