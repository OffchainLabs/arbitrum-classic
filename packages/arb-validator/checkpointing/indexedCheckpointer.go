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

	"github.com/golang/protobuf/proto"
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

type ckpHeightBounds struct {
	lo *common.TimeBlocks
	hi *common.TimeBlocks
}

func (db *ckpHeightBounds) Marshal() *HeightBoundsBuf {
	return &HeightBoundsBuf{
		Lo: db.lo.Marshal(),
		Hi: db.hi.Marshal(),
	}
}

func (dbuf *HeightBoundsBuf) Unmarshal() *ckpHeightBounds {
	return &ckpHeightBounds{
		lo: dbuf.Lo.Unmarshal(),
		hi: dbuf.Hi.Unmarshal(),
	}
}

func (cp *IndexedCheckpointer) HasCheckpointedState() bool {
	cp.Lock()
	defer cp.Unlock()

	bounds, err := cp.getHeightBounds()
	if err != nil {
		log.Fatal(err)
	}
	return bounds != nil
}

func (cp *IndexedCheckpointer) getHeightBounds() (*ckpHeightBounds, error) {
	key := []byte{0}
	valBytes := cp.db.GetData(key)
	if valBytes == nil {
		return nil, nil
	}
	heightBounds := &HeightBoundsBuf{}
	if err := proto.Unmarshal(valBytes, heightBounds); err != nil {
		return nil, err
	}
	return heightBounds.Unmarshal(), nil
}

func (cp *IndexedCheckpointer) setHeightBounds(bounds *ckpHeightBounds) error {
	key := []byte{0}
	buf := bounds.Marshal()
	val, err := proto.Marshal(buf)
	if err != nil {
		return err
	}
	ok := cp.db.SaveData(key, val)
	if !ok {
		return errors.New("db write error in checkpointer.setHeightBounds")
	}
	return nil
}

func (cp *IndexedCheckpointer) updateHeightUpperBound(height *common.TimeBlocks) error {
	bounds, err := cp.getHeightBounds()
	if err != nil {
		return err
	}
	if bounds == nil {
		bounds = &ckpHeightBounds{
			lo: height.Clone(),
			hi: height.Clone(),
		}
		return cp.setHeightBounds(bounds)
	} else if bounds.hi.Cmp(height) < 0 {
		bounds.hi = height.Clone()
		return cp.setHeightBounds(bounds)
	} else {
		return nil
	}
}

func (cp *IndexedCheckpointer) getIdsAtHeight(height *common.TimeBlocks) ([]*common.BlockId, error) {
	key := append([]byte{1}, height.AsInt().Bytes()...)
	valBytes := cp.db.GetData(key)
	if valBytes == nil {
		return nil, nil
	}
	idBufList := &BlockIdBufList{}
	if err := proto.Unmarshal(valBytes, idBufList); err != nil {
		return nil, err
	}
	ret := []*common.BlockId{}
	for _, idBuf := range idBufList.Bufs {
		ret = append(ret, idBuf.Unmarshal())
	}
	return ret, nil
}

func (cp *IndexedCheckpointer) setIdsAtHeight(height *common.TimeBlocks, ids []*common.BlockId) error {
	key := append([]byte{1}, height.AsInt().Bytes()...)
	idBufs := []*common.BlockIdBuf{}
	for _, id := range ids {
		idBufs = append(idBufs, id.MarshalToBuf())
	}
	idBufList := &BlockIdBufList{
		Bufs: idBufs,
	}
	valBytes, err := proto.Marshal(idBufList)
	if err != nil {
		return err
	}
	ok := cp.db.SaveData(key, valBytes)
	if !ok {
		return errors.New("db write error in checkpointer.setIdsAtHeight")
	}
	return nil
}

func (cp *IndexedCheckpointer) deleteIdsAtHeight(height *common.TimeBlocks) error {
	key := append([]byte{1}, height.AsInt().Bytes()...)
	if ok := cp.db.DeleteData(key); !ok {
		return errors.New("Checkpointer failed to delete idsAtHeight")
	}
	return nil
}

func (cp *IndexedCheckpointer) recordIdAsCheckpointed(newId *common.BlockId) error {
	ids, err := cp.getIdsAtHeight(newId.Height)
	if err != nil {
		return err
	}
	ids = append(ids, newId)
	return cp.setIdsAtHeight(newId.Height, ids)
}

func (cp *IndexedCheckpointer) GetInitialMachine() (machine.Machine, error) {
	cp.Lock()
	defer cp.Unlock()

	return cp.db.GetInitialMachine()
}

type writableCheckpoint struct {
	blockId  *common.BlockId
	contents []byte
	ckpCtx   CheckpointContext
}

func (cp *IndexedCheckpointer) AsyncSaveCheckpoint(
	blockId *common.BlockId,
	contents []byte,
	cpCtx CheckpointContext,
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

func (_ *IndexedCheckpointer) makeContentsKey(id *common.BlockId) []byte {
	bidBuf := id.MarshalToBuf()
	bytesBuf, err := proto.Marshal(bidBuf)
	if err != nil {
		log.Fatal(err)
	}
	return append([]byte{2}, bytesBuf...)
}

func (cp *IndexedCheckpointer) RestoreLatestState(ctx context.Context, clnt arbbridge.ArbClient, unmarshalFunc func([]byte, RestoreContext) error) error {
	cp.Lock()
	defer cp.Unlock()

	heightBounds, err := cp.getHeightBounds()
	if err != nil {
		return err
	}
	if heightBounds == nil {
		return errors.New("Cannot restore because no checkpoint exists")
	}
	for height := heightBounds.hi; height.Cmp(heightBounds.lo) >= 0; height = common.NewTimeBlocks(new(big.Int).Sub(height.AsInt(), big.NewInt(1))) {
		ids, err := cp.getIdsAtHeight(height)
		if err != nil {
			return err
		}
		onchainId, err := clnt.BlockIdForHeight(ctx, height)
		if err != nil {
			return err
		}
		for _, id := range ids {
			if id.Equals(onchainId) {
				key := cp.makeContentsKey(id)
				val := cp.db.GetData(key)
				ckpWithMan := &CheckpointWithManifest{}
				if err := proto.Unmarshal(val, ckpWithMan); err != nil {
					return err
				}
				return unmarshalFunc(ckpWithMan.Contents, cp.newRestoreContextLocked())
			}
		}
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
			return errors.New("Failed to write value to checkpoint db")
		}
	}
	for _, mach := range wc.ckpCtx.Machines() {
		if ok := mach.Checkpoint(cp.db); !ok {
			return errors.New("Failed to write machine to checkpoint db")
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
	key := cp.makeContentsKey(wc.blockId)
	if ok := cp.db.SaveData(key, bytesBuf); !ok {
		return errors.New("Failed to write checkpoint to checkpoint db")
	}

	// record this blockId as checkpointed
	if err := cp.recordIdAsCheckpointed(wc.blockId); err != nil {
		return nil
	}

	// update height bounds if needed
	return cp.updateHeightUpperBound(wc.blockId.Height)
}

func (cp *IndexedCheckpointer) cleanupDaemon() {
	ticker := time.NewTicker(common.NewTimeBlocksInt(25).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		func() {
			cp.Lock()
			bounds, err := cp.getHeightBounds()
			cp.Unlock()
			if err != nil {
				return
			}
			height := common.NewTimeBlocks(new(big.Int).Sub(bounds.lo.AsInt(), big.NewInt(1)))
			heightLimit := common.NewTimeBlocks(new(big.Int).Sub(bounds.hi.AsInt(), cp.maxReorgHeight))
			var prevHeight *common.TimeBlocks = nil
			prevIds := []*common.BlockId{}
			for height.Cmp(heightLimit) < 0 {
				cp.Lock()
				ids, err := cp.getIdsAtHeight(height)
				cp.Unlock()
				if err != nil {
					return
				}
				if len(ids) > 0 {
					for _, id := range prevIds {
						_ = cp.deleteCheckpointForId(id) // OK to call without lock; callee will acquire lock
					}
					cp.Lock()
					if prevHeight != nil {
						if err := cp.deleteIdsAtHeight(prevHeight); err != nil {
							cp.Unlock()
							return
						}
					}
					bounds, err := cp.getHeightBounds()
					if err != nil {
						cp.Unlock()
						return
					}
					bounds.lo = height
					if err := cp.setHeightBounds(bounds); err != nil {
						cp.Unlock()
						return
					}
					cp.Unlock()
					prevHeight = height
					prevIds = ids
				}
				height = common.NewTimeBlocks(new(big.Int).Add(height.AsInt(), big.NewInt(1)))
			}
		}()
	}
}

func (cp *IndexedCheckpointer) deleteCheckpointForId(id *common.BlockId) error {
	cp.Lock()
	defer cp.Unlock()

	key := cp.makeContentsKey(id)
	val := cp.db.GetData(key)
	ckp := &CheckpointWithManifest{}
	if err := proto.Unmarshal(val, ckp); err != nil {
		return err
	}
	_ = cp.db.DeleteData(key) // ignore error
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

type restoreContextLocked struct {
	cp *IndexedCheckpointer
}

func (cp *IndexedCheckpointer) newRestoreContextLocked() RestoreContext {
	return &restoreContextLocked{cp}
}

func (rcl *restoreContextLocked) GetValue(h common.Hash) value.Value {
	return rcl.cp.getValue_locked(h)
}

func (rcl *restoreContextLocked) GetMachine(h common.Hash) machine.Machine {
	return rcl.cp.getMachine_locked(h)
}

func (cp *IndexedCheckpointer) GetValue(h common.Hash) value.Value {
	cp.Lock()
	defer cp.Unlock()

	return cp.getValue_locked(h)
}

func (cp *IndexedCheckpointer) getValue_locked(h common.Hash) value.Value {
	return cp.db.GetValue(h)
}

func (cp *IndexedCheckpointer) GetMachine(h common.Hash) machine.Machine {
	cp.Lock()
	defer cp.Unlock()

	return cp.getMachine_locked(h)
}

func (cp *IndexedCheckpointer) getMachine_locked(h common.Hash) machine.Machine {
	ret, err := cp.db.GetInitialMachine()
	if err != nil {
		log.Fatal(err)
	}
	if ret.Hash() == h {
		return ret
	}
	restored := ret.RestoreCheckpoint(cp.db, h)
	if !restored {
		log.Fatalln("Failed to restore machine", h, "from checkpoint")
	}
	if ret.Hash() != h {
		log.Fatalln("Restore machine", h, "from checkpoint with wrong hash", ret.Hash())
	}
	return ret
}
