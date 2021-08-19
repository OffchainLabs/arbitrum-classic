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

package core

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type MessagesSynced struct {
	Metrics *Metrics
	MaxDiff int64
}

func (c MessagesSynced) Execute(context.Context) (interface{}, error) {
	totalMessages := c.Metrics.MessageCount.Value()
	processedMessages := c.Metrics.MachineMessagesReadCount.Value()
	details := fmt.Sprintf("inbox messages processed = %v/%v", processedMessages, totalMessages)
	if totalMessages-processedMessages > c.MaxDiff {
		return details, errors.New("core thread behind")
	}
	return details, nil
}

func (c MessagesSynced) Name() string {
	return "core-messages-syncing"
}

type LogsSynced struct {
	Metrics *Metrics
	MaxDiff int64
}

func (c LogsSynced) Execute(context.Context) (interface{}, error) {
	totalLogs := c.Metrics.LogCount.Value()
	processedLogs := c.Metrics.LogsCursorPosition.Value()
	details := fmt.Sprintf("inbox messages processed = %v/%v", processedLogs, totalLogs)
	if totalLogs-processedLogs > c.MaxDiff {
		return details, errors.New("log reader behind")
	}
	return details, nil
}

func (c LogsSynced) Name() string {
	return "core-logs-syncing"
}
