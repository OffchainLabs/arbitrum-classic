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
	"github.com/pkg/errors"
	"math/big"
	"os"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var initialEntryBlockId = &common.BlockId{
	Height:     common.NewTimeBlocksInt(10),
	HeaderHash: common.Hash{20},
}

var initialEntryBlockId2 = &common.BlockId{
	Height:     common.NewTimeBlocksInt(10),
	HeaderHash: common.Hash{30},
}

var laterEntryBlockId = &common.BlockId{
	Height:     common.NewTimeBlocksInt(15),
	HeaderHash: common.Hash{21},
}

var laterEntryBlockId2 = &common.BlockId{
	Height:     common.NewTimeBlocksInt(15),
	HeaderHash: common.Hash{31},
}

var distantEntryBlockId = &common.BlockId{
	Height:     common.NewTimeBlocksInt(200),
	HeaderHash: common.Hash{41},
}

var checkpointData = []byte{5, 3, 2}
var checkpointData2 = []byte{5, 3, 4}
var checkpointData3 = []byte{5, 3, 5}

var dbPath = "./testdb"
var maxReorgHeight = big.NewInt(100)

type TimeGetterMock struct {
	blockIdFunc func(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error)
}

func (m *TimeGetterMock) BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	return m.blockIdFunc(ctx, height)
}

func (m *TimeGetterMock) TimestampForBlockHash(context.Context, common.Hash) (*big.Int, error) {
	return nil, errors.New("unsupported method")
}

func TestMain(m *testing.M) {
	code := m.Run()
	if err := os.RemoveAll(dbPath); err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}
	os.Exit(code)
}

func TestEmpty(t *testing.T) {
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	if cp.HasCheckpointedState() {
		t.Error("checkpoint should start empty")
	}
}

func TestWriteCheckpoint(t *testing.T) {
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	blockData, err := cp.bs.GetBlock(initialEntryBlockId)
	if err == nil {
		t.Error("block shouldn't exist before writing")
	}
	_ = blockData

	/* TODO
	checkpointContext := ckptcontext.NewCheckpointContext()
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  initialEntryBlockId,
		contents: checkpointData,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}

	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  laterEntryBlockId,
		contents: checkpointData2,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}


	blockData, err = cp.bs.GetBlock(initialEntryBlockId)
	if err != nil {
		t.Error(err)
	}

	ckpWithMan := &CheckpointWithManifest{}
	if err := proto.Unmarshal(blockData, ckpWithMan); err != nil {
		t.Error(err)
	}

	if !bytes.Equal(ckpWithMan.Contents, checkpointData) {
		t.Error("block data didn't match. Got:", blockData, "wanted:", checkpointData)
	}

	blockData2, err := cp.bs.GetBlock(laterEntryBlockId)
	if err != nil {
		t.Error(err)
	}

	ckpWithMan2 := &CheckpointWithManifest{}
	if err := proto.Unmarshal(blockData2, ckpWithMan2); err != nil {
		t.Error(err)
	}

	if !bytes.Equal(ckpWithMan2.Contents, checkpointData2) {
		t.Error("block data didn't match. Got:", blockData2, "wanted:", checkpointData2)
	}
	*/
}

func TestRestoreEmpty(t *testing.T) {
	/* TODO
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	if err = cp.RestoreLatestState(context.Background(), &TimeGetterMock{}, func(bytes []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		return nil
	}); err != errNoCheckpoint {
		t.Error(err)
	}
	*/
}

func TestRestoreSingleCheckpoint(t *testing.T) {
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	/* TODO
	ctx := context.Background()

	checkpointContext := ckptcontext.NewCheckpointContext()
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  initialEntryBlockId,
		contents: checkpointData,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}

	tgm := &TimeGetterMock{
		func(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
			if height.Cmp(initialEntryBlockId.Height) != 0 {
				t.Error("incorrect initial time query")
			}
			return initialEntryBlockId2, nil
		},
	}
	// Should fail restore if checkpoint has changed
	if err = cp.RestoreLatestState(ctx, tgm, func(bytes []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		t.Error("unmarshal func called")
		return nil
	}); err != errNoMatchingCheckpoint {
		t.Error(err)
	}

	tgm = &TimeGetterMock{
		func(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
			if height.Cmp(initialEntryBlockId.Height) != 0 {
				t.Error("incorrect initial time query")
			}
			return initialEntryBlockId, nil
		},
	}
	// Should succeed restore if checkpoint hasn't changed
	if err = cp.RestoreLatestState(ctx, tgm, func(data []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		if !bytes.Equal(data, checkpointData) {
			t.Error("incorrect checkpoint data restored")
		}
		return nil
	}); err != nil {
		t.Error(err)
	}

	*/
}

func TestRestoreReorg(t *testing.T) {
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	/* TOOD
	checkpointContext := ckptcontext.NewCheckpointContext()
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  initialEntryBlockId,
		contents: checkpointData,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}

	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  laterEntryBlockId,
		contents: checkpointData2,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}

	gap := new(big.Int).Sub(laterEntryBlockId.Height.AsInt(), initialEntryBlockId.Height.AsInt()).Uint64()
	generateReorgTimeGetterMock := func(
		firstResult *common.BlockId,
		earlierResult *common.BlockId,
	) func(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
		runCount := uint64(0)
		return func(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
			if new(big.Int).Sub(laterEntryBlockId.Height.AsInt(), height.AsInt()).Uint64() != runCount {
				t.Error("should query each height once in descending order", runCount, height)
			}
			runCount++
			if runCount == 1 {
				return firstResult, nil
			} else if runCount == gap {
				return earlierResult, nil
			} else {
				return &common.BlockId{
					Height:     height,
					HeaderHash: common.Hash{9},
				}, nil
			}
		}
	}

	ctx := context.Background()
	tgm := &TimeGetterMock{generateReorgTimeGetterMock(laterEntryBlockId, initialEntryBlockId)}
	// Should restore to latest without reorg
	if err = cp.RestoreLatestState(ctx, tgm, func(data []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		if !bytes.Equal(data, checkpointData2) {
			t.Error("incorrect checkpoint data restored")
		}
		return nil
	}); err != nil {
		t.Error(err)
	}

	tgm = &TimeGetterMock{generateReorgTimeGetterMock(laterEntryBlockId2, initialEntryBlockId)}
	// Should restore older after reorg
	if err = cp.RestoreLatestState(ctx, tgm, func(data []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		if !bytes.Equal(data, checkpointData) {
			t.Error("incorrect checkpoint data restored")
		}
		return nil
	}); err != nil {
		t.Error(err)
	}

	tgm = &TimeGetterMock{generateReorgTimeGetterMock(laterEntryBlockId2, initialEntryBlockId2)}
	// Should fail to restore if everything is reorged out
	if err = cp.RestoreLatestState(ctx, tgm, func(data []byte, restoreContext ckptcontext.RestoreContext, _ *common.BlockId) error {
		t.Error("shouldn't be able to restore")
		return nil
	}); err != errNoMatchingCheckpoint {
		t.Error(err)
	}
	*/
}

func TestCleanup(t *testing.T) {
	var rollupAddr common.Address
	cp, err := newIndexedCheckpointer(rollupAddr, dbPath, maxReorgHeight, true)
	if err != nil {
		t.Fatal(err)
	}
	defer cp.db.CloseCheckpointStorage()

	/* TODO
	checkpointContext := ckptcontext.NewCheckpointContext()
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  initialEntryBlockId,
		contents: checkpointData,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  laterEntryBlockId,
		contents: checkpointData2,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}
	if err = writeCheckpoint(cp.bs, cp.db, &writableCheckpoint{
		blockId:  distantEntryBlockId,
		contents: checkpointData3,
		ckpCtx:   checkpointContext,
	}); err != nil {
		t.Error(err)
	}

	if cp.bs.MinBlockStoreHeight().Cmp(initialEntryBlockId.Height) != 0 {
		t.Error("minimum height incorrect")
	}

	cleanup(cp.bs, cp.db, big.NewInt(100))

	if cp.bs.MinBlockStoreHeight().Cmp(laterEntryBlockId.Height) != 0 {
		t.Error("minimum height incorrect after cleanup", cp.bs.MinBlockStoreHeight())
	}

	_, err = cp.bs.GetBlock(laterEntryBlockId)
	if err != nil {
		t.Error(err)
	}

	_, err = cp.bs.GetBlock(distantEntryBlockId)
	if err != nil {
		t.Error(err)
	}
	*/
}
