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

// The below code has been ported from the healthcheck library. The original license is included here
// Copyright 2017 by the contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodehealth

import (
	"net/http"

	"github.com/ethereum/go-ethereum/metrics"
	"github.com/heptiolabs/healthcheck"
)

type metricsHandler struct {
	handler   healthcheck.Handler
	registry  metrics.Registry
	namespace string
}

// NewMetricsHandler returns a healthcheck Handler that also exposes metrics
// into the provided Prometheus registry.
func NewMetricsHandler(registry metrics.Registry, namespace string) healthcheck.Handler {
	return &metricsHandler{
		handler:   healthcheck.NewHandler(),
		registry:  registry,
		namespace: namespace,
	}
}

func (h *metricsHandler) AddLivenessCheck(name string, check healthcheck.Check) {
	h.handler.AddLivenessCheck(name, h.wrap(name, check))
}

func (h *metricsHandler) AddReadinessCheck(name string, check healthcheck.Check) {
	h.handler.AddReadinessCheck(name, h.wrap(name, check))
}

func (h *metricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func (h *metricsHandler) LiveEndpoint(w http.ResponseWriter, r *http.Request) {
	h.handler.LiveEndpoint(w, r)
}

func (h *metricsHandler) ReadyEndpoint(w http.ResponseWriter, r *http.Request) {
	h.handler.ReadyEndpoint(w, r)
}

func (h *metricsHandler) wrap(name string, check healthcheck.Check) healthcheck.Check {
	gauge := metrics.NewFunctionalGauge(
		func() int64 {
			if check() == nil {
				return 0
			}
			return 1
		},
	)

	fullName := h.namespace + "/healthcheck/" + name
	// Unregister and replace the existing check if there is one
	h.registry.Unregister(fullName)
	if err := h.registry.Register(fullName, gauge); err != nil {
		panic(err)
	}
	return check
}
