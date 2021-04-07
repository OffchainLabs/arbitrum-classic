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

package staker

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = log.With().Caller().Stack().Str("component", "staker").Logger()

type Monitor struct {
	Storage machine.ArbStorage
	Core    core.ArbCore
}

func NewMonitor(dbDir string, contractFile string) (*Monitor, error) {
	storage, err := cmachine.NewArbStorage(dbDir)
	if err != nil {
		return nil, err
	}

	err = storage.Initialize(contractFile)
	if err != nil {
		return nil, err
	}

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		return nil, errors.New("error starting ArbCore thread")
	}

	return &Monitor{
		Storage: storage,
		Core:    arbCore,
	}, nil
}

func (m *Monitor) Close() {
	m.Storage.CloseArbStorage()
}

func (m *Monitor) StartInboxReader(ctx context.Context, ethurl string, rollupAddress common.Address) (*InboxReader, error) {
	ethClient, err := ethutils.NewRPCEthClient(ethurl)
	if err != nil {
		return nil, err
	}
	rollup, err := ethbridge.NewRollupWatcher(rollupAddress.ToEthAddress(), ethClient)
	if err != nil {
		return nil, err
	}
	panic("TODO: redo inbox reader")
	bridgeAddress, err := rollup.DelayedBridge(context.Background())
	if err != nil {
		return nil, err
	}
	bridgeWatcher, err := ethbridge.NewBridgeWatcher(bridgeAddress.ToEthAddress(), ethClient)
	if err != nil {
		return nil, err
	}
	reader, err := NewInboxReader(ctx, bridgeWatcher, m.Core)
	if err != nil {
		return nil, err
	}
	reader.Start(ctx)
	return reader, nil
}
