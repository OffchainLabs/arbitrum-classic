/*
* Copyright 2022, Offchain Labs, Inc.
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

package dev

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
	"math/big"
)

func NewForkNode(
	ctx context.Context,
	dir string,
	chainId *big.Int,
	agg common.Address,
	reorgMessage int64,
) (*Backend, *txdb.TxDB, func(), <-chan error, error) {
	nodeConfig := configuration.DefaultNodeSettings()
	coreConfig := configuration.DefaultCoreSettingsMaxExecution()
	coreConfig.LazyLoadCoreMachine = true
	coreConfig.Cache.Last = true
	coreConfig.CheckpointPruneOnStartup = false
	coreConfig.CheckpointPruningMode = "off"
	coreConfig.LazyLoadArchiveQueries = true
	coreConfig.Cache.SeedOnStartup = false

	mon, err := monitor.NewMonitor(dir, coreConfig)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error opening monitor")
	}
	if err := mon.InitializeExisting(); err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error opening monitor")
	}
	if err := mon.Start(); err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error opening monitor")
	}

	db, errChan, err := txdb.New(ctx, mon.Core, mon.Storage.GetNodeStore(), nodeConfig)
	if err != nil {
		mon.Close()
		return nil, nil, nil, nil, errors.Wrap(err, "error opening txdb")
	}

	if err := core.ReorgAndWait(mon.Core, new(big.Int).SetInt64(reorgMessage)); err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error reorging")
	}

	latestBlock, err := db.LatestBlock()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	l2Block, err := db.GetL2Block(latestBlock)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	initialL1Height := l2Block.L1BlockNum.Uint64() + 1

	backendCore, err := NewBackendCore(ctx, mon.Core, chainId)
	if err != nil {
		mon.Close()
		return nil, nil, nil, nil, err
	}

	cancel := func() {
		db.Close()
		mon.Close()
	}
	signer := types.NewEIP155Signer(chainId)
	l1 := NewL1Emulator(initialL1Height)
	backend := NewBackend(ctx, backendCore, db, l1, signer, agg, big.NewInt(100000000000))

	return backend, db, cancel, errChan, nil
}

func GetMessageCount(dir string) (uint64, error) {
	coreConfig := configuration.DefaultCoreSettingsMaxExecution()
	mon, err := monitor.NewMonitor(dir, coreConfig)
	if err != nil {
		return 0, errors.Wrap(err, "error opening monitor")
	}
	msgCount, err := mon.Core.GetMessageCount()
	mon.Close()
	return msgCount.Uint64(), err
}
