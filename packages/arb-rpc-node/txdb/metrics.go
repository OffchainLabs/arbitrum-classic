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

package txdb

import (
	"github.com/ethereum/go-ethereum/metrics"
)

type Metrics struct {
	LogsAddedTotal metrics.Gauge
	LogsAdded      metrics.Counter
	LogsDeleted    metrics.Counter
	LatestBlock    metrics.Gauge
	BlocksAdded    metrics.Counter
}

func NewMetrics() *Metrics {
	return &Metrics{
		LogsAddedTotal: metrics.NewGauge(),
		LogsAdded:      metrics.NewCounter(),
		LogsDeleted:    metrics.NewCounter(),
		LatestBlock:    metrics.NewGauge(),
		BlocksAdded:    metrics.NewCounter(),
	}
}

func (m *Metrics) Register(r metrics.Registry) error {
	if err := r.Register("logs_added_total", m.LogsAddedTotal); err != nil {
		return err
	}
	if err := r.Register("logs_added", m.LogsAdded); err != nil {
		return err
	}
	if err := r.Register("logs_deleted", m.LogsDeleted); err != nil {
		return err
	}
	if err := r.Register("latest_block", m.LatestBlock); err != nil {
		return err
	}
	if err := r.Register("blocks_added", m.BlocksAdded); err != nil {
		return err
	}
	return nil
}
