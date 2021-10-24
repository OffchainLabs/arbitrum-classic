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

package batcher

import (
	"context"

	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type ForwarderHealth struct {
	Url          string
	ReadyUrl     string
	MaxBlockDiff int64
	TxDBMetrics  *txdb.Metrics
}

func NewForwarderCheck(url string, txDBMetrics *txdb.Metrics, config configuration.Healthcheck) ForwarderHealth {
	return ForwarderHealth{
		Url:          url,
		ReadyUrl:     config.ForwarderReadyURL,
		MaxBlockDiff: config.MaxL2BlockDiff,
		TxDBMetrics:  txDBMetrics,
	}
}

func (c ForwarderHealth) Execute(ctx context.Context) (interface{}, error) {
	// If we're checking readyness, first check that
	if c.ReadyUrl != "" {
		check, err := checks.NewHTTPCheck(checks.HTTPCheckConfig{
			CheckName: "",
			URL:       c.Url,
			Timeout:   0,
		})
		if err != nil {
			return nil, err
		}
		httpDetails, err := check.Execute(ctx)
		if err != nil {
			return httpDetails, err
		}
	}

	client, err := ethclient.DialContext(ctx, c.Url)
	if err != nil {
		return "failed dialing node", err
	}
	targetBlockNum, err := client.BlockNumber(ctx)
	if err != nil {
		return "failed getting target block", err
	}
	blockDiff := c.TxDBMetrics.LatestBlock.Value() - int64(targetBlockNum)
	if blockDiff > c.MaxBlockDiff {
		return nil, errors.Errorf("Target is %v blocks behind with target max %v", blockDiff, c.MaxBlockDiff)
	}
	if blockDiff < 0 && -blockDiff > c.MaxBlockDiff {
		return nil, errors.Errorf("Target is %v blocks ahead with target max %v", -blockDiff, c.MaxBlockDiff)
	}
	return nil, nil
}

func (c ForwarderHealth) Name() string {
	return "transaction-submission"
}

type SequencerHealth struct {
}

func (c SequencerHealth) Execute(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (c SequencerHealth) Name() string {
	return "transaction-submission"
}

type LockoutSequencerHealth struct {
}

func (c LockoutSequencerHealth) Execute(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (c LockoutSequencerHealth) Name() string {
	return "transaction-submission"
}

type BatcherHealth struct {
}

func (c BatcherHealth) Execute(context.Context) (interface{}, error) {
	return nil, nil
}

func (c BatcherHealth) Name() string {
	return "transaction-submission"
}
