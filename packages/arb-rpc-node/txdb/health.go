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
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

type Synced struct {
	CoreMetrics *core.Metrics
	Metrics     *Metrics
	MaxDiff     int64
}

func NewSyncedCheck(coreMetrics *core.Metrics, metrics *Metrics, config configuration.Healthcheck) Synced {
	return Synced{
		CoreMetrics: coreMetrics,
		Metrics:     metrics,
		MaxDiff:     config.MaxLogsProcessedSyncDiff,
	}
}

func (c Synced) Execute(context.Context) (interface{}, error) {
	logsOutput := c.CoreMetrics.LogCount.Value()
	logsProcessed := c.Metrics.LogsAddedTotal.Value()
	details := fmt.Sprintf("txdb processed %v/%v logs", logsProcessed, logsOutput)
	if logsOutput-logsProcessed > c.MaxDiff {
		return details, errors.New("txdb not synced")
	}
	return details, nil
}

func (c Synced) Name() string {
	return "txdb-synced"
}
