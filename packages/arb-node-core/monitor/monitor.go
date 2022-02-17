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
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"
)

var logger = arblog.Logger.With().Str("component", "monitor").Logger()

type Monitor struct {
	Storage machine.ArbStorage
	Core    core.ArbCore
	Reader  *InboxReader
}

func NewInitializedMonitor(dbDir string, contractFile string, coreConfig *configuration.Core) (*Monitor, error) {
	m, err := NewMonitor(dbDir, coreConfig)
	if err != nil {
		return nil, err
	}
	if err := m.Initialize(contractFile); err != nil {
		return nil, err
	}
	if err := m.Start(); err != nil {
		return nil, err
	}
	return m, nil
}

func NewMonitor(dbDir string, coreConfig *configuration.Core) (*Monitor, error) {
	storage, err := cmachine.NewArbStorage(dbDir, coreConfig)
	if err != nil {
		return nil, err
	}
	logger.Info().Str("directory", dbDir).Msg("database opened")
	return &Monitor{
		Storage: storage,
		Core:    storage.GetArbCore(),
	}, nil
}

func (m *Monitor) Initialize(contractFile string) error {
	err := m.Storage.Initialize(contractFile)
	if err != nil {
		return err
	}
	logger.Info().Msg("storage initialized")
	return nil
}

func (m *Monitor) InitializeExisting() error {
	err := m.Storage.InitializeExisting()
	if err != nil {
		return err
	}
	logger.Info().Msg("storage initialized")
	return nil
}

func (m *Monitor) Start() error {
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
	healthChan chan nodehealth.Log,
	sequencerFeed chan broadcaster.BroadcastFeedMessage,
	inboxReaderConfig configuration.InboxReader,
) (*InboxReader, error) {
	rollup, err := ethbridge.NewRollupWatcher(rollupAddress.ToEthAddress(), fromBlock, ethClient, bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	creationEvent, err := rollup.LookupCreation(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error checking initial chain state")
	}
	initialExecutionCursor, err := m.Core.GetExecutionCursor(big.NewInt(0), true)
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
	reader, err := NewInboxReader(
		ctx,
		delayedBridgeWatcher,
		sequencerInboxWatcher,
		bridgeUtils,
		m.Core,
		healthChan,
		sequencerFeed,
		inboxReaderConfig,
	)
	if err != nil {
		return nil, err
	}
	reader.Start(ctx, inboxReaderConfig.DelayBlocks)
	m.Reader = reader
	m.listenForSignal(ctx)
	return reader, nil
}

func (m *Monitor) listenForSignal(ctx context.Context) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGUSR1)
	go func() {
		defer close(signalChan)
		for {
			select {
			case <-ctx.Done():
				break
			case sig := <-signalChan:
				switch sig {
				case syscall.SIGUSR1:
					logger.Info().Msg("triggering save of rocksdb checkpoint")
					m.Core.SaveRocksdbCheckpoint()
				default:
					logger.Info().Str("signal", sig.String()).Msg("caught unexpected signal")
				}
			}
		}
	}()
}
