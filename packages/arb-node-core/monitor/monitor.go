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
	"math/big"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = log.With().Caller().Stack().Str("component", "monitor").Logger()

type Metrics struct {
	Inbox *InboxMetrics
	Core  *core.Metrics
}

func (m *Metrics) Register(r metrics.Registry) error {
	coreRegistry := metrics.NewPrefixedChildRegistry(r, "core/")
	if err := m.Core.Register(coreRegistry); err != nil {
		return err
	}
	inboxRegistry := metrics.NewPrefixedChildRegistry(r, "inbox/")
	return m.Inbox.Register(inboxRegistry)
}

func (m *Metrics) RegisterSyncChecks(config configuration.Healthcheck, health gosundheit.Health) error {
	if err := health.RegisterCheck(NewInboxSyncedCheck(m, config)); err != nil {
		return err
	}
	if err := health.RegisterCheck(core.NewMessagesSyncedCheck(m.Core, config)); err != nil {
		return err
	}
	return nil
}

type Monitor struct {
	Storage machine.ArbStorage
	Core    core.ArbCore
	Reader  *InboxReader
	Metrics *Metrics
}

func NewMonitor(dbDir string, coreConfig *configuration.Core) (*Monitor, error) {
	storage, err := cmachine.NewArbStorage(dbDir, coreConfig)
	if err != nil {
		return nil, err
	}

	arbCore := storage.GetArbCore()

	return &Monitor{
		Storage: storage,
		Core:    arbCore,
		Metrics: &Metrics{
			Core:  core.NewArbCoreMetrics(arbCore),
			Inbox: NewInboxMetrics(),
		},
	}, nil
}

func NewStartedMonitor(dbDir, contractFile string, coreConfig *configuration.Core) (*Monitor, error) {
	m, err := NewMonitor(dbDir, coreConfig)
	if err != nil {
		return nil, err
	}
	return m, m.StartCore(contractFile)
}

func (m *Monitor) StartCore(contractFile string) error {
	err := m.Storage.Initialize(contractFile)
	if err != nil {
		return err
	}

	started := m.Core.StartThread()
	if !started {
		return errors.New("error starting ArbCore thread")
	}
	return nil
}

func (m *Monitor) Close() {
	if m.Reader != nil {
		m.Reader.Stop()
	}
	m.Storage.CloseArbStorage()
	logger.Info().Msg("Database closed")
}

func (m *Monitor) StartInboxReader(
	ctx context.Context,
	ethClient ethutils.EthClient,
	rollupAddress common.Address,
	fromBlock int64,
	bridgeUtilsAddress common.Address,
	sequencerFeed chan broadcaster.BroadcastFeedMessage,
) (*InboxReader, error) {
	rollup, err := ethbridge.NewRollupWatcher(rollupAddress.ToEthAddress(), fromBlock, ethClient, bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	creationEvent, err := rollup.LookupCreation(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error checking initial chain state")
	}
	initialExecutionCursor, err := m.Core.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		return nil, errors.Wrap(err, "error loading initial ArbCore machine")
	}
	initialMachineHash := initialExecutionCursor.MachineHash()
	if initialMachineHash != creationEvent.MachineHash {
		return nil, errors.Errorf("Initial machine hash loaded from arbos.mexe doesn't match chain's initial machine hash: chain %v, arbCore %v", hexutil.Encode(creationEvent.MachineHash[:]), initialMachineHash)
	}

	delayedBridgeAddress, err := rollup.DelayedBridge(ctx)
	if err != nil {
		return nil, err
	}
	delayedBridgeWatcher, err := ethbridge.NewDelayedBridgeWatcher(delayedBridgeAddress.ToEthAddress(), fromBlock, ethClient)
	if err != nil {
		return nil, err
	}
	sequencerAddress, err := rollup.SequencerBridge(ctx)
	if err != nil {
		return nil, err
	}
	sequencerInboxWatcher, err := ethbridge.NewSequencerInboxWatcher(sequencerAddress.ToEthAddress(), ethClient)
	if err != nil {
		return nil, err
	}
	bridgeUtils, err := ethbridge.NewBridgeUtils(bridgeUtilsAddress.ToEthAddress(), ethClient, delayedBridgeWatcher, sequencerInboxWatcher)
	if err != nil {
		return nil, err
	}
	reader, err := NewInboxReader(ctx, delayedBridgeWatcher, sequencerInboxWatcher, bridgeUtils, m.Core, sequencerFeed, m.Metrics.Inbox)
	if err != nil {
		return nil, err
	}
	reader.Start(ctx)
	m.Reader = reader
	return reader, nil
}

func (m *Monitor) TryStartInboxReaderLoop(ctx context.Context, l1URL string, sequencerFeed chan broadcaster.BroadcastFeedMessage, config *configuration.Config) (*InboxReader, error) {
	tryCreate := func() (*InboxReader, error) {
		l1Client, err := ethutils.NewRPCEthClient(l1URL)
		if err != nil {
			return nil, err
		}
		return m.StartInboxReader(ctx, l1Client, common.HexToAddress(config.Rollup.Address), config.Rollup.FromBlock, common.HexToAddress(config.BridgeUtilsAddress), sequencerFeed)
	}
	for {
		inboxReader, err := tryCreate()
		if err == nil {
			return inboxReader, nil
		}
		logger.Warn().Err(err).
			Str("url", config.L1.URL).
			Str("rollup", config.Rollup.Address).
			Str("bridgeUtils", config.BridgeUtilsAddress).
			Msg("failed to start inbox reader, waiting and retrying")

		select {
		case <-ctx.Done():
			return nil, errors.New("ctx cancelled StartInboxReader retry loop")
		case <-time.After(5 * time.Second):
		}
	}
}
