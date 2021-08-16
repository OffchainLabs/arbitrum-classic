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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

var logger = log.With().Caller().Stack().Str("component", "monitor").Logger()

type Monitor struct {
	Storage machine.ArbStorage
	Core    core.ArbCore
	Reader  *InboxReader
}

func NewMonitor(dbDir string, contractFile string, coreConfig *configuration.Core) (*Monitor, error) {
	if coreConfig.Profile.ResetAllExceptInbox {
		err := cmachine.ResetAllExceptInbox(dbDir, contractFile)
		if err != nil {
			return nil, err
		}

		if coreConfig.Profile.RunUntil == 0 {
			return nil, errors.New("database reset except inbox, nothing else to do")
		}
	}

	storage, err := cmachine.NewArbStorage(dbDir, coreConfig)
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
	reader, err := NewInboxReader(ctx, delayedBridgeWatcher, sequencerInboxWatcher, bridgeUtils, m.Core, healthChan, sequencerFeed)
	if err != nil {
		return nil, err
	}
	reader.Start(ctx)
	m.Reader = reader
	return reader, nil
}
