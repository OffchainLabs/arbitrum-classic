/*
* Copyright 2020, Offchain Labs, Inc.
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

package checkpointing

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/pkg/errors"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type IndexedCheckpointer struct {
	*sync.Mutex
	db                    *cmachine.ArbStorage
	bs                    machine.BlockStore
	nextCheckpointToWrite *writableCheckpoint
	maxReorgHeight        *big.Int
}

func NewIndexedCheckpointer(
	rollupAddr common.Address,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) (*IndexedCheckpointer, error) {
	ret, err := newIndexedCheckpointer(
		rollupAddr,
		databasePath,
		new(big.Int).Set(maxReorgHeight),
		forceFreshStart,
	)

	if err != nil {
		return nil, err
	}

	go cleanupDaemon(ret.bs, ret.db, maxReorgHeight)
	return ret, nil
}

// newIndexedCheckpointerFactory creates the checkpointer, but doesn't
// launch it's reading and writing threads. This is useful for deterministic
// testing
func newIndexedCheckpointer(
	rollupAddr common.Address,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) (*IndexedCheckpointer, error) {
	if databasePath == "" {
		databasePath = MakeCheckpointDatabasePath(rollupAddr)
	}
	if forceFreshStart {
		// for testing only --  delete old database to get fresh start
		if err := os.RemoveAll(databasePath); err != nil {
			return nil, err
		}
	}
	cCheckpointer, err := cmachine.NewCheckpoint(databasePath)
	if err != nil {
		return nil, err
	}

	return &IndexedCheckpointer{
		new(sync.Mutex),
		cCheckpointer,
		cCheckpointer.GetBlockStore(),
		nil,
		maxReorgHeight,
	}, nil
}

func (cp *IndexedCheckpointer) Initialize(arbitrumCodeFilePath string) error {
	return cp.db.Initialize(arbitrumCodeFilePath)
}

func (cp *IndexedCheckpointer) Initialized() bool {
	return cp.db.Initialized()
}

func (cp *IndexedCheckpointer) MaxReorgHeight() *big.Int {
	return new(big.Int).Set(cp.maxReorgHeight)
}

func (cp *IndexedCheckpointer) GetAggregatorStore() *cmachine.AggregatorStore {
	return cp.db.GetAggregatorStore()
}

// HasCheckpointedState checks whether the block store is empty, which is the table
// which contains the checkpoints recorded by the IndexedCheckpointer
func (cp *IndexedCheckpointer) HasCheckpointedState() bool {
	return false
}

func (cp *IndexedCheckpointer) GetInitialMachine(vc machine.ValueCache) (machine.Machine, error) {
	return cp.db.GetInitialMachine(vc)
}

type writableCheckpoint struct {
	blockId  *common.BlockId
	contents []byte
	ckpCtx   *ckptcontext.CheckpointContext
	errChan  chan<- error
}

func (cp *IndexedCheckpointer) AsyncSaveCheckpoint(
	blockId *common.BlockId,
	contents []byte,
	cpCtx *ckptcontext.CheckpointContext,
) <-chan error {
	cp.Lock()
	defer cp.Unlock()

	errChan := make(chan error, 1)

	if cp.nextCheckpointToWrite != nil {
		cp.nextCheckpointToWrite.errChan <- errors.New("replaced by newer checkpoint")
		close(cp.nextCheckpointToWrite.errChan)
	}

	cp.nextCheckpointToWrite = &writableCheckpoint{
		blockId:  blockId,
		contents: contents,
		ckpCtx:   cpCtx,
		errChan:  errChan,
	}

	return errChan
}

func (cp *IndexedCheckpointer) RestoreLatestState(ctx context.Context, clnt arbbridge.ChainTimeGetter, unmarshalFunc func([]byte, ckptcontext.RestoreContext, *common.BlockId) error) error {
	return restoreLatestState(ctx, cp.bs, cp.db, clnt, unmarshalFunc)
}

func restoreLatestState(
	ctx context.Context,
	bs machine.BlockStore,
	db machine.ArbStorage,
	clnt arbbridge.ChainTimeGetter,
	unmarshalFunc func([]byte, ckptcontext.RestoreContext, *common.BlockId) error,
) error {
	if bs.IsBlockStoreEmpty() {
		return errNoCheckpoint
	}

	startHeight := bs.MaxBlockStoreHeight()
	lowestHeight := bs.MinBlockStoreHeight()

	logger.Info().
		Str("max", bs.MaxBlockStoreHeight().String()).
		Str("min", bs.MinBlockStoreHeight().String()).
		Msg("Restoring latest state")
	for height := startHeight; height.Cmp(lowestHeight) >= 0; height = common.NewTimeBlocks(new(big.Int).Sub(height.AsInt(), big.NewInt(1))) {
		logger := logger.With().Str("height", height.String()).Logger()
		onchainId, err := clnt.BlockIdForHeight(ctx, height)
		if err != nil {
			if err == ethereum.NotFound {
				continue
			}
			return err
		}

		blockData, err := bs.GetBlock(onchainId)
		if err != nil {
			// If no record was found, try the next block
			continue
		}
		ckpWithMan := &CheckpointWithManifest{}
		if err := proto.Unmarshal(blockData, ckpWithMan); err != nil {
			// If something went wrong, try the next block
			logger.Warn().Err(err).Msg("Invalid block found when restoring state")
			continue
		}

		rcl, err := newRestoreContextLocked(db, ckpWithMan.Manifest)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Failed load manifest data")
			continue
		}
		if err := unmarshalFunc(ckpWithMan.Contents, rcl, onchainId); err != nil {
			logger.Error().Stack().Err(err).Msg("Failed load checkpoint")
			continue
		}
		logger.Info().Msg("Restored state")
		return nil
	}
	return errNoMatchingCheckpoint
}

func (cp *IndexedCheckpointer) writeDaemon() {
	ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cp.Lock()
		checkpoint := cp.nextCheckpointToWrite
		cp.nextCheckpointToWrite = nil
		cp.Unlock()
		if checkpoint != nil {
			err := writeCheckpoint(cp.bs, cp.db, checkpoint)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("Error writing checkpoint")
			}
			checkpoint.errChan <- err
			close(checkpoint.errChan)
		}
	}
}

func writeCheckpoint(bs machine.BlockStore, db machine.ArbStorage, wc *writableCheckpoint) error {
	// save values and machines
	if err := ckptcontext.SaveCheckpointContext(db, wc.ckpCtx); err != nil {
		return err
	}

	// save main checkpoint data
	ckpWithMan := &CheckpointWithManifest{
		Contents: wc.contents,
		Manifest: wc.ckpCtx.Manifest(),
	}
	bytesBuf, err := proto.Marshal(ckpWithMan)
	if err != nil {
		return err
	}
	if err := bs.PutBlock(wc.blockId, bytesBuf); err != nil {
		return errors.New("failed to write checkpoint to checkpoint db")
	}

	return nil
}

func cleanupDaemon(bs machine.BlockStore, db machine.ArbStorage, maxReorgHeight *big.Int) {
	ticker := time.NewTicker(common.NewTimeBlocksInt(25).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cleanup(bs, db, maxReorgHeight)
	}
}

func cleanup(bs machine.BlockStore, db machine.ArbStorage, maxReorgHeight *big.Int) {
	currentMin := bs.MinBlockStoreHeight()
	currentMax := bs.MaxBlockStoreHeight()
	height := common.NewTimeBlocks(new(big.Int).Sub(currentMin.AsInt(), big.NewInt(1)))
	heightLimit := common.NewTimeBlocks(new(big.Int).Sub(currentMax.AsInt(), maxReorgHeight))
	var prevIds []*common.BlockId
	for height.Cmp(heightLimit) < 0 {
		blockIds := bs.BlocksAtHeight(height)
		if len(blockIds) > 0 {
			for _, id := range prevIds {
				err := deleteCheckpointForKey(bs, db, id)
				if err != nil {
					// Can still continue if error
					logger.Warn().Stack().Err(err).Object("id", id).Msg("Error deleting checkpoint")
				}
			}
			prevIds = blockIds
		}
		height = common.NewTimeBlocks(new(big.Int).Add(height.AsInt(), big.NewInt(1)))
	}
}
