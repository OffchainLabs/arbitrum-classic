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
	"errors"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var errNoCheckpoint = errors.New("cannot restore because no checkpoint exists")
var errNoMatchingCheckpoint = errors.New("cannot restore because no matching checkpoint exists")

type IndexedCheckpointer struct {
	*sync.Mutex
	db                    machine.CheckpointStorage
	nextCheckpointToWrite *writableCheckpoint
}

func NewIndexedCheckpointerFactory(
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) RollupCheckpointerFactory {
	ret, err := newIndexedCheckpointerFactory(
		rollupAddr,
		arbitrumCodeFilePath,
		databasePath,
		forceFreshStart,
	)

	if err != nil {
		log.Fatal(err)
	}

	go ret.writeDaemon()
	go cleanupDaemon(ret.db, maxReorgHeight)
	return ret
}

// newIndexedCheckpointerFactory creates the checkpoint factory, but doesn't
// launch it's reading and writing threads. This is useful for deterministic
// testing
func newIndexedCheckpointerFactory(
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	databasePath string,
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
	cCheckpointer, err := cmachine.NewCheckpoint(databasePath, arbitrumCodeFilePath)
	if err != nil {
		return nil, err
	}

	return &IndexedCheckpointer{
		new(sync.Mutex),
		cCheckpointer,
		nil,
	}, nil
}

// The checkpointer interface uses a factory pattern. The idea is that the rollup manager makes a factory, then
// uses that factory to make a first checkpointer. On every reorg it kills the old checkpointer and calls the factory
// to make a new checkpointer. But IndexedCheckpointer is reorg-aware, so it doesn't need to die and get
// re-instantiated after each reorg.  To comply with the factory-based interface, it just makes a single object
// which acts as both the factory and the checkpointer. So when the factory's New is called, it just returns itself.
func (cp *IndexedCheckpointer) New(_ context.Context) RollupCheckpointer {
	return cp
}

func (cp *IndexedCheckpointer) HasCheckpointedState() bool {
	return !cp.db.IsBlockStoreEmpty()
}

func (cp *IndexedCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.db.GetInitialMachine()
}

type writableCheckpoint struct {
	blockId  *common.BlockId
	contents []byte
	ckpCtx   *CheckpointContext
}

func (cp *IndexedCheckpointer) AsyncSaveCheckpoint(
	blockId *common.BlockId,
	contents []byte,
	cpCtx *CheckpointContext,
) {
	cp.Lock()
	defer cp.Unlock()

	cp.nextCheckpointToWrite = &writableCheckpoint{
		blockId:  blockId,
		contents: contents,
		ckpCtx:   cpCtx,
	}
}

func (cp *IndexedCheckpointer) RestoreLatestState(ctx context.Context, clnt arbbridge.ChainTimeGetter, unmarshalFunc func([]byte, RestoreContext) error) error {
	return restoreLatestState(ctx, cp.db, clnt, unmarshalFunc)
}

func restoreLatestState(ctx context.Context, db machine.CheckpointStorage, clnt arbbridge.ChainTimeGetter, unmarshalFunc func([]byte, RestoreContext) error) error {
	if db.IsBlockStoreEmpty() {
		return errNoCheckpoint
	}

	startHeight := db.MaxBlockStoreHeight()
	lowestHeight := db.MinBlockStoreHeight()

	for height := startHeight; height.Cmp(lowestHeight) >= 0; height = common.NewTimeBlocks(new(big.Int).Sub(height.AsInt(), big.NewInt(1))) {
		onchainId, err := clnt.BlockIdForHeight(ctx, height)
		if err != nil {
			return err
		}
		blockData, err := db.GetBlock(onchainId)
		if err != nil {
			// If no record was found, try the next block
			continue
		}
		ckpWithMan := &CheckpointWithManifest{}
		if err := proto.Unmarshal(blockData, ckpWithMan); err != nil {
			// If something went wrong, try the next block
			continue
		}
		return unmarshalFunc(ckpWithMan.Contents, &restoreContextLocked{db})

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
			err := writeCheckpoint(cp.db, checkpoint)
			if err != nil {
				log.Println("Error writing checkpoint: {}", err)
			}
		}
	}
}

func writeCheckpoint(db machine.CheckpointStorage, wc *writableCheckpoint) error {
	// save values and machines
	for _, val := range wc.ckpCtx.Values() {
		if ok := db.SaveValue(val); !ok {
			return errors.New("failed to write value to checkpoint db")
		}
	}
	for _, mach := range wc.ckpCtx.Machines() {
		if ok := mach.Checkpoint(db); !ok {
			return errors.New("failed to write machine to checkpoint db")
		}
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
	if err := db.PutBlock(wc.blockId, bytesBuf); err != nil {
		return errors.New("failed to write checkpoint to checkpoint db")
	}

	return nil
}

func cleanupDaemon(db machine.CheckpointStorage, maxReorgHeight *big.Int) {
	ticker := time.NewTicker(common.NewTimeBlocksInt(25).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cleanup(db, maxReorgHeight)
	}
}

func cleanup(db machine.CheckpointStorage, maxReorgHeight *big.Int) {
	currentMin := db.MinBlockStoreHeight()
	currentMax := db.MaxBlockStoreHeight()
	height := common.NewTimeBlocks(new(big.Int).Sub(currentMin.AsInt(), big.NewInt(1)))
	heightLimit := common.NewTimeBlocks(new(big.Int).Sub(currentMax.AsInt(), maxReorgHeight))
	var prevIds []*common.BlockId
	for height.Cmp(heightLimit) < 0 {
		blockIds := db.BlocksAtHeight(height)
		if len(blockIds) > 0 {
			for _, id := range prevIds {
				_ = deleteCheckpointForKey(db, id)
			}
			prevIds = blockIds
		}
		height = common.NewTimeBlocks(new(big.Int).Add(height.AsInt(), big.NewInt(1)))
	}
}

func deleteCheckpointForKey(db machine.CheckpointStorage, id *common.BlockId) error {
	val, err := db.GetBlock(id)
	if err != nil {
		return err
	}
	ckp := &CheckpointWithManifest{}
	if err := proto.Unmarshal(val, ckp); err != nil {
		return err
	}
	_ = db.DeleteBlock(id) // ignore error
	if ckp.Manifest != nil {
		for _, hbuf := range ckp.Manifest.Values {
			h := hbuf.Unmarshal()
			_ = db.DeleteValue(h) // ignore error
		}
		for _, hbuf := range ckp.Manifest.Machines {
			h := hbuf.Unmarshal()
			_ = db.DeleteCheckpoint(h) // ignore error
		}
	}
	return nil
}

type restoreContextLocked struct {
	db machine.CheckpointStorage
}

func (rcl *restoreContextLocked) GetValue(h common.Hash) value.Value {
	return rcl.db.GetValue(h)
}

func (rcl *restoreContextLocked) GetMachine(h common.Hash) machine.Machine {
	ret, err := rcl.db.GetMachine(h)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
