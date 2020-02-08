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
	"github.com/golang/protobuf/proto"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"log"
	"math/big"
	"os"
	"sync"
	"time"
)

type IndexedCheckpointer struct {
	*sync.Mutex
	db            machine.CheckpointStorage
	maxReorgDepth *big.Int

	nextCheckpointToWrite *writableCheckpoint
}

func NewIndexedCheckpoinerFactory(
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	databasePath string,
	maxReorgDepth *big.Int,
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
		new(big.Int).Set(maxReorgDepth),
		nil,
	}
	go ret.writeDaemon()
	go ret.cleanupDaemon()
	return ret
}

func (cp *IndexedCheckpointer) New(_ context.Context) RollupCheckpointer {
	return cp
}

type ckpDepthBounds struct {
	lo *common.TimeBlocks
	hi *common.TimeBlocks
}

func (db *ckpDepthBounds) Marshal() *DepthBoundsBuf {
	return &DepthBoundsBuf{
		Lo: db.lo.Marshal(),
		Hi: db.hi.Marshal(),
	}
}

func (dbuf *DepthBoundsBuf) Unmarshal() *ckpDepthBounds {
	return &ckpDepthBounds{
		lo: dbuf.Lo.Unmarshal(),
		hi: dbuf.Hi.Unmarshal(),
	}
}

func (cp *IndexedCheckpointer) HasCheckpointedState() bool {
	cp.Lock()
	defer cp.Unlock()

	bounds, err := cp.getDepthBounds()
	if err != nil {
		log.Fatal(err)
	}
	return bounds != nil
}

func (cp *IndexedCheckpointer) getDepthBounds() (*ckpDepthBounds, error) {
	key := []byte{0}
	valBytes := cp.db.GetData(key)
	depthBounds := &DepthBoundsBuf{}
	if err := proto.Unmarshal(valBytes, depthBounds); err != nil {
		return nil, err
	}
	return depthBounds.Unmarshal(), nil
}

func (cp *IndexedCheckpointer) setDepthBounds(bounds *ckpDepthBounds) error {
	key := []byte{0}
	buf := bounds.Marshal()
	val, err := proto.Marshal(buf)
	if err != nil {
		return err
	}
	ok := cp.db.SaveData(key, val)
	if !ok {
		return errors.New("db write error in checkpointer.setDepthBounds")
	}
	return nil
}

func (cp *IndexedCheckpointer) updateDepthUpperBound(depth *common.TimeBlocks) error {
	bounds, err := cp.getDepthBounds()
	if err != nil {
		return err
	}
	if bounds.hi.Cmp(depth) < 0 {
		bounds.hi = depth
		return cp.setDepthBounds(bounds)
	} else {
		return nil
	}
}

func (cp *IndexedCheckpointer) getIdsAtDepth(depth *common.TimeBlocks) ([]*common.BlockId, error) {
	key := append([]byte{1}, depth.AsInt().Bytes()...)
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

func (cp *IndexedCheckpointer) setIdsAtDepth(depth *common.TimeBlocks, ids []*common.BlockId) error {
	key := append([]byte{1}, depth.AsInt().Bytes()...)
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
		return errors.New("db write error in checkpointer.setIdsAtDepth")
	}
	return nil
}

func (cp *IndexedCheckpointer) recordIdAsCheckpointed(newId *common.BlockId) error {
	ids, err := cp.getIdsAtDepth(newId.Height)
	if err != nil {
		return err
	}
	ids = append(ids, newId)
	return cp.setIdsAtDepth(newId.Height, ids)
}

func (cp *IndexedCheckpointer) GetInitialMachine() (machine.Machine, error) {
	cp.Lock()
	defer cp.Unlock()

	return cp.db.GetInitialMachine()
}

type writableCheckpoint struct {
	blockId       *common.BlockId
	contents      []byte
	ckpCtx        CheckpointContext
	closeWhenDone chan struct{}
}

func (cp *IndexedCheckpointer) AsyncSaveCheckpoint(
	blockId *common.BlockId,
	contents []byte,
	cpCtx CheckpointContext,
	closeWhenDone chan struct{},
) {
	cp.Lock()
	defer cp.Unlock()

	if cp.nextCheckpointToWrite != nil {
		close(cp.nextCheckpointToWrite.closeWhenDone)
	}
	cp.nextCheckpointToWrite = &writableCheckpoint{
		blockId:       blockId,
		contents:      contents,
		ckpCtx:        cpCtx,
		closeWhenDone: closeWhenDone,
	}
}

func (_ *IndexedCheckpointer) makeContentsKey(id *common.BlockId) []byte {
	bidBuf := &common.BlockIdBuf{
		Height:     id.Height.Marshal(),
		HeaderHash: id.HeaderHash.MarshalToBuf(),
	}
	bytesBuf, err := proto.Marshal(bidBuf)
	if err != nil {
		log.Fatal(err)
	}
	return append([]byte{2}, bytesBuf...)
}

func (cp *IndexedCheckpointer) RestoreLatestState(
	ctx context.Context,
	clnt arbbridge.ArbClient,
	_ common.Address,
	_ bool,
) (
	[]byte,
	RestoreContext,
	error,
) {
	cp.Lock()
	defer cp.Unlock()

	depthBounds, err := cp.getDepthBounds()
	if err != nil {
		return nil, nil, err
	}
	for depth := depthBounds.hi; depth.Cmp(depthBounds.lo) >= 0; depth = common.NewTimeBlocks(new(big.Int).Sub(depth.AsInt(), big.NewInt(1))) {
		ids, err := cp.getIdsAtDepth(depth)
		if err != nil {
			return nil, nil, err
		}
		onchainId, err := clnt.BlockIdForHeight(ctx, depth)
		if err != nil {
			return nil, nil, err
		}
		for _, id := range ids {
			if id.Equals(onchainId) {
				key := cp.makeContentsKey(id)
				val := cp.db.GetData(key)
				ckpWithMan := &CheckpointWithManifest{}
				if err := proto.Unmarshal(val, ckpWithMan); err != nil {
					return nil, nil, err
				}
				return ckpWithMan.Contents, cp, nil
			}
		}
	}
	return nil, nil, nil
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
			close(cp.nextCheckpointToWrite.closeWhenDone)
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

	// update depth bounds if needed
	return cp.updateDepthUpperBound(wc.blockId.Height)
}

func (cp *IndexedCheckpointer) cleanupDaemon() {
	//TODO
	log.Println("Checkpoint cleanup daemon not yet implemented")
}

func (cp *IndexedCheckpointer) GetValue(h common.Hash) value.Value {
	cp.Lock()
	defer cp.Unlock()

	return cp.db.GetValue(h)
}

func (cp *IndexedCheckpointer) GetMachine(h common.Hash) machine.Machine {
	cp.Lock()
	defer cp.Unlock()

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
