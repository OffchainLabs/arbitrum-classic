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

type IndexedCheckpointer struct {
	*sync.Mutex
	db                    machine.CheckpointStorage
	maxReorgHeight        *big.Int
	nextCheckpointToWrite *writableCheckpoint
	chansToClose          []chan struct{}
}

func NewIndexedCheckpointerFactory(
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) RollupCheckpointerFactory {
	if databasePath == "" {
		databasePath = MakeCheckpointDatabasePath(rollupAddr)
	}
	if forceFreshStart {
		// for testing only --  delete old database to get fresh start
		if err := os.RemoveAll(databasePath); err != nil {
			log.Fatal(err)
		}
	}
	cCheckpointer, err := cmachine.NewCheckpoint(databasePath, arbitrumCodeFilePath)
	if err != nil {
		log.Fatal(err)
	}

	ret := &IndexedCheckpointer{
		new(sync.Mutex),
		cCheckpointer,
		new(big.Int).Set(maxReorgHeight),
		nil,
		nil,
	}
	go ret.writeDaemon()
	go ret.cleanupDaemon()
	return ret
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
	cp.Lock()
	defer cp.Unlock()

	return !cp.db.IsBlockStoreEmpty()
}

func (cp *IndexedCheckpointer) GetInitialMachine() (machine.Machine, error) {
	cp.Lock()
	defer cp.Unlock()

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
	closeWhenDone chan struct{},
) {
	cp.Lock()
	defer cp.Unlock()

	if closeWhenDone != nil {
		cp.chansToClose = append(cp.chansToClose, closeWhenDone)
	}
	cp.nextCheckpointToWrite = &writableCheckpoint{
		blockId:  blockId,
		contents: contents,
		ckpCtx:   cpCtx,
	}
}

func (cp *IndexedCheckpointer) RestoreLatestState(ctx context.Context, clnt arbbridge.ArbClient, unmarshalFunc func([]byte, RestoreContext) error) error {
	cp.Lock()
	defer cp.Unlock()

	if !cp.HasCheckpointedState() {
		return errors.New("cannot restore because no checkpoint exists")
	}

	startHeight := cp.db.MaxBlockStoreHeight()
	lowestHeight := cp.db.MinBlockStoreHeight()

	for height := startHeight; height.Cmp(lowestHeight) >= 0; height = common.NewTimeBlocks(new(big.Int).Sub(height.AsInt(), big.NewInt(1))) {
		onchainId, err := clnt.BlockIdForHeight(ctx, height)
		if err != nil {
			return err
		}
		blockData, err := cp.db.GetBlock(onchainId)
		if err != nil {
			// If no record was found, try the next block
			continue
		}
		ckpWithMan := &CheckpointWithManifest{}
		if err := proto.Unmarshal(blockData, ckpWithMan); err != nil {
			// If something went wrong, try the next block
			continue
		}
		return unmarshalFunc(ckpWithMan.Contents, cp.newRestoreContextLocked())

	}
	log.Fatal("Called RestoreLatestState on checkpointer that has no stored checkpoints")
	return nil // can't reach this but need to make the compiler happy
}

func (cp *IndexedCheckpointer) writeDaemon() {
	ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cp.Lock()
		if cp.nextCheckpointToWrite != nil {
			err := cp.writeCheckpoint(cp.nextCheckpointToWrite)
			if err != nil {
				log.Println("Error writing checkpoint: {}", err)
			}
			for _, c := range cp.chansToClose {
				close(c)
			}
			cp.chansToClose = nil
			cp.nextCheckpointToWrite = nil
		}
		cp.Unlock()
	}
}

func (cp *IndexedCheckpointer) writeCheckpoint(wc *writableCheckpoint) error {
	// save values and machines
	for _, val := range wc.ckpCtx.Values() {
		if ok := cp.db.SaveValue(val); !ok {
			return errors.New("failed to write value to checkpoint db")
		}
	}
	for _, mach := range wc.ckpCtx.Machines() {
		if ok := mach.Checkpoint(cp.db); !ok {
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
	if err := cp.db.PutBlock(wc.blockId, bytesBuf); err != nil {
		return errors.New("failed to write checkpoint to checkpoint db")
	}

	return nil
}

func (cp *IndexedCheckpointer) cleanupDaemon() {
	ticker := time.NewTicker(common.NewTimeBlocksInt(25).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cp.Lock()
		cp.cleanup()
		cp.Unlock()
	}
}

func (cp *IndexedCheckpointer) cleanup() {
	height := common.NewTimeBlocks(new(big.Int).Sub(cp.db.MinBlockStoreHeight().AsInt(), big.NewInt(1)))
	heightLimit := common.NewTimeBlocks(new(big.Int).Sub(cp.db.MaxBlockStoreHeight().AsInt(), cp.maxReorgHeight))
	var prevIds []*common.BlockId
	for height.Cmp(heightLimit) < 0 {
		blockIds := cp.db.BlocksAtHeight(height)
		if len(blockIds) > 0 {
			for _, id := range prevIds {
				_ = cp.deleteCheckpointForKey(id)
			}
			prevIds = blockIds
		}
		height = common.NewTimeBlocks(new(big.Int).Add(height.AsInt(), big.NewInt(1)))
	}
}

func (cp *IndexedCheckpointer) deleteCheckpointForKey(id *common.BlockId) error {
	val, err := cp.db.GetBlock(id)
	if err != nil {
		return err
	}
	ckp := &CheckpointWithManifest{}
	if err := proto.Unmarshal(val, ckp); err != nil {
		return err
	}
	_ = cp.db.DeleteBlock(id) // ignore error
	if ckp.Manifest != nil {
		for _, hbuf := range ckp.Manifest.Values {
			h := hbuf.Unmarshal()
			_ = cp.db.DeleteValue(h) // ignore error
		}
		for _, hbuf := range ckp.Manifest.Machines {
			h := hbuf.Unmarshal()
			_ = cp.db.DeleteCheckpoint(h) // ignore error
		}
	}
	return nil
}

func (cp *IndexedCheckpointer) getValueLocked(h common.Hash) value.Value {
	return cp.db.GetValue(h)
}

func (cp *IndexedCheckpointer) getMachineLocked(h common.Hash) machine.Machine {
	ret, err := cp.db.GetMachine(h)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func (cp *IndexedCheckpointer) newRestoreContextLocked() RestoreContext {
	return &restoreContextLocked{cp}
}

type restoreContextLocked struct {
	cp *IndexedCheckpointer
}

func (rcl *restoreContextLocked) GetValue(h common.Hash) value.Value {
	return rcl.cp.getValueLocked(h)
}

func (rcl *restoreContextLocked) GetMachine(h common.Hash) machine.Machine {
	return rcl.cp.getMachineLocked(h)
}
