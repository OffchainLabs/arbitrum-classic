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

package monitor

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type InboxSynced struct {
	Metrics *Metrics
	MaxDiff int64
}

func NewInboxSyncedCheck(metrics *Metrics, config configuration.Healthcheck) InboxSynced {
	return InboxSynced{
		Metrics: metrics,
		MaxDiff: config.MaxInboxSyncDiff,
	}
}

func (c InboxSynced) Execute(context.Context) (interface{}, error) {
	if c.Metrics.Inbox.Initialized.Value() != 1 {
		return nil, errors.New("inbox reader not initialized")
	}
	chainCount := c.Metrics.Inbox.L1MessageCount.Value()
	localCount := c.Metrics.Core.MessageCount.Value()
	details := fmt.Sprintf("inbox messages synced = %v/%v", localCount, chainCount)
	if chainCount-localCount > c.MaxDiff {
		return details, errors.New("inbox reader behind")
	}
	return details, nil
}

func (c InboxSynced) Name() string {
	return "inbox-reader"
}
