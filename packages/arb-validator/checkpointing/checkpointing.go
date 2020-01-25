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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"log"
	"math/big"
	"os"

	"github.com/gogo/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type RollupCheckpointerFactory interface {
	New(ctx context.Context) RollupCheckpointer
}

type RollupCheckpointer interface {
	HasCheckpointedState() bool
	RestoreLatestState(context.Context, arbbridge.ArbClient, common.Address, bool) (content []byte, resCtx structures.RestoreContext, err error)
	GetInitialMachine() (machine.Machine, error)
	AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{})
}

type RollupCheckpointerImplFactory struct {
	rollupAddr      common.Address
	arbCodeFilePath string
	databasePath    string
	maxReorgDepth   *big.Int
	forceFreshStart bool
}

func NewRollupCheckpointerImplFactory(
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	databasePath string,
	maxReorgDepth *big.Int,
	forceFreshStart bool,
) *RollupCheckpointerImplFactory {
	if databasePath == "" {
		databasePath = makeCheckpointDatabasePath(rollupAddr)
	}
	return &RollupCheckpointerImplFactory{
		rollupAddr,
		arbitrumCodeFilePath,
		databasePath,
		maxReorgDepth,
		forceFreshStart,
	}
}

type RollupCheckpointerImpl struct {
	st            machine.CheckpointStorage
	maxReorgDepth *big.Int
	asyncWriter   *asyncCheckpointWriter
}

const checkpointDatabasePathBase = "/tmp/arb-validator-checkpoint-"

func makeCheckpointDatabasePath(rollupAddr common.Address) string {
	return checkpointDatabasePathBase + rollupAddr.Hex()[2:]
}

func (fac *RollupCheckpointerImplFactory) New(ctx context.Context) RollupCheckpointer {
	if fac.forceFreshStart {
		// for testing only -- use production checkpointer but delete old database first
		if err := os.RemoveAll(fac.databasePath); err != nil {
			log.Fatal(err)
		}
		fac.forceFreshStart = false
	}
	cCheckpointer, err := cmachine.NewCheckpoint(fac.databasePath, fac.arbCodeFilePath)
	if err != nil {
		log.Fatal(err)
	}
	ret := &RollupCheckpointerImpl{
		maxReorgDepth: fac.maxReorgDepth,
		st:            cCheckpointer,
	}
	ret.asyncWriter = NewAsyncCheckpointWriter(ctx, ret)
	return ret
}

func (rcp *RollupCheckpointerImpl) _saveCheckpoint(
	id *structures.BlockId,
	contents []byte,
	checkpointCtx structures.CheckpointContext,
) error {
	// read in metadata
	var metadataBuf *structures.CheckpointMetadata
	var newestInCp *structures.BlockId
	rawMetadata := rcp.RestoreMetadata()

	// read in metadata, or create it if it doesn't already exist
	if rawMetadata == nil || len(rawMetadata) == 0 {
		idBuf := id.MarshalToBuf()
		metadataBuf = &structures.CheckpointMetadata{
			FormatVersion: 1,
			Oldest:        idBuf,
			Newest:        idBuf,
		}
		buf, err := proto.Marshal(metadataBuf)
		if err != nil {
			return err
		}
		rcp.SaveMetadata(buf)
	} else {
		metadataBuf = &structures.CheckpointMetadata{}
		if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
			return err
		}
	}
	newestInCp = metadataBuf.Newest.Unmarshal()
	// save all of the data for this checkpoint
	rcp.SaveCheckpoint(
		id,
		newestInCp,
		contents,
		checkpointCtx.Manifest(),
		checkpointCtx.Values(),
		checkpointCtx.Machines(),
	)

	// update the metadata to include this checkpoint
	metadataBuf.Newest = id.MarshalToBuf()
	buf, err := proto.Marshal(metadataBuf)
	if err != nil {
		return err
	}
	rcp.SaveMetadata(buf)

	return nil
}

func (rcp *RollupCheckpointerImpl) HasCheckpointedState() bool {
	metadataBytes := rcp.RestoreMetadata()
	return metadataBytes != nil && len(metadataBytes) > 0
}

func (rcp *RollupCheckpointerImpl) RestoreLatestState(
	ctx context.Context,
	client arbbridge.ArbClient,
	contractAddr common.Address,
	beOpinionated bool,
) ([]byte, structures.RestoreContext, error) {
	rcp.QueueReorgedCheckpointsForDeletion(ctx, client)

	metadataBytes := rcp.RestoreMetadata()
	if !rcp.HasCheckpointedState() {
		return nil, nil, errors.New("no checkpoints in database")
	}
	metadata := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadata); err != nil {
		return nil, nil, err
	}
	newestId := metadata.Newest.Unmarshal()
	cobBytes, resCtx, err := rcp.RestoreCheckpoint(newestId)
	if err != nil {
		return nil, nil, err
	}
	return cobBytes, resCtx, nil
}

func (rcp *RollupCheckpointerImpl) RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext, error) {
	var metadataBuf *structures.CheckpointMetadata
	var oldestHeightInCp *common.TimeBlocks
	var newestHeightInCp *common.TimeBlocks
	rawMetadata := rcp.RestoreMetadata()
	if rawMetadata == nil {
		return nil, nil, nil
	}

	metadataBuf = &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
		return nil, nil, err
	}
	oldestHeightInCp = metadataBuf.Oldest.Height.Unmarshal()
	newestHeightInCp = metadataBuf.Newest.Height.Unmarshal()

	blockHeight := blockId.Height
	if blockHeight.Cmp(oldestHeightInCp) < 0 || blockHeight.Cmp(newestHeightInCp) > 0 {
		return nil, nil, nil
	}

	// read contents
	contentsKey := getContentsKey(blockId)
	contentBytes := rcp.st.GetData(contentsKey)

	return contentBytes, rcp, nil
}

func (rcp *RollupCheckpointerImpl) GetInitialMachine() (machine.Machine, error) {
	return rcp.st.GetInitialMachine()
}

func (rcp *RollupCheckpointerImpl) AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{}) {
	rcp.asyncWriter.SubmitJob(
		func() {
			rcp._saveCheckpoint(blockId, contents, cpCtx)
		},
		closeWhenDone,
	)
}

func (rcp *RollupCheckpointerImpl) Close() {
	rcp.st.CloseCheckpointStorage()
}

func getKeyForId(prefix []byte, id *structures.BlockId) []byte {
	idBytes, err := proto.Marshal(id.MarshalToBuf())
	if err != nil {
		log.Fatal(err)
	}
	return append(prefix, idBytes...)
}

func getManifestKey(blockId *structures.BlockId) []byte {
	return getKeyForId([]byte("manifest:"), blockId)
}

func getContentsKey(blockId *structures.BlockId) []byte {
	return getKeyForId([]byte("contents:"), blockId)
}

func getLinksKey(blockId *structures.BlockId) []byte {
	return getKeyForId([]byte("links:"), blockId)
}

func (rcp *RollupCheckpointerImpl) SaveMetadata(data []byte) {
	ok := rcp.st.SaveData([]byte("metadata"), data)
	if !ok {
		log.Fatal("metadata checkpointing failure")
	}
}

func (rcp *RollupCheckpointerImpl) RestoreMetadata() []byte {
	return rcp.st.GetData([]byte("metadata"))
}

func (rcp *RollupCheckpointerImpl) SaveCheckpoint(
	blockId *structures.BlockId,
	prevBlockId *structures.BlockId,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[common.Hash]value.Value,
	machines map[common.Hash]machine.Machine,
) {
	for _, val := range values {
		rcp.st.SaveValue(val)
	}

	for _, mach := range machines {
		savedMachine := mach.Checkpoint(rcp.st)
		if !savedMachine {
			log.Fatalln("Failed to checkpoint machine")
		}
	}

	manifestBuf, err := proto.Marshal(manifest)
	if err != nil {
		log.Fatal(err)
	}
	rcp.st.SaveData(getManifestKey(blockId), manifestBuf)

	rcp.st.SaveData(getContentsKey(blockId), contents)

	rcp._updateNextPointer(prevBlockId, blockId)
	rcp._setBothPointers(blockId, prevBlockId, blockId)
}

func (rcp *RollupCheckpointerImpl) _setBothPointers(id, prev, next *structures.BlockId) {
	links := &structures.CheckpointLinks{
		Prev: prev.MarshalToBuf(),
		Next: next.MarshalToBuf(),
	}
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	rcp.st.SaveData(getLinksKey(id), linksBuf)
}

func (rcp *RollupCheckpointerImpl) _updatePrevPointer(id, prev *structures.BlockId) {
	key := getLinksKey(id)
	linksBuf := rcp.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Prev = prev.MarshalToBuf()
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	rcp.st.SaveData(key, linksBuf)
}

func (rcp *RollupCheckpointerImpl) _updateNextPointer(id, next *structures.BlockId) {
	key := getLinksKey(id)
	linksBuf := rcp.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Next = next.MarshalToBuf()
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	rcp.st.SaveData(key, linksBuf)
}

func (rcp *RollupCheckpointerImpl) QueueCheckpointForDeletion(blockId *structures.BlockId) {
	// make a best effort to delete an old checkpoint, but ignore any errors
	// errors might cause some harmless extra info to remain in the database

	queueBytes := rcp.st.GetData([]byte("deadqueue"))
	queue := &structures.BlockIdBufList{}
	if err := proto.Unmarshal(queueBytes, queue); err != nil {
		return
	}

	queue.Bufs = append(queue.Bufs, blockId.MarshalToBuf())

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return
	}
	rcp.st.SaveData([]byte("deadqueue"), queueBytes)
}

func (rcp *RollupCheckpointerImpl) QueueReorgedCheckpointsForDeletion(ctx context.Context, client arbbridge.ArbClient) {
	metadataBuf := rcp.RestoreMetadata()
	if len(metadataBuf) == 0 {
		return
	}
	metadata := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBuf, metadata); err != nil {
		return
	}

	oldestId := metadata.Oldest.Unmarshal()
	newestId := metadata.Newest.Unmarshal()
	for oldestId.Height.Cmp(newestId.Height) < 0 {
		onChainId, err := client.BlockIdForHeight(ctx, newestId.Height)
		if err != nil {
			return
		}
		if onChainId.HeaderHash.Equals(newestId.HeaderHash) {
			// success
			return
		}
		linksBytes := rcp.st.GetData(getLinksKey(newestId))
		linksBuf := &structures.CheckpointLinks{}
		if err := proto.Unmarshal(linksBytes, linksBuf); err != nil {
			return
		}
		metadata.Newest = linksBuf.Prev
		metadataBuf, err = proto.Marshal(metadata)
		if err != nil {
			return
		}
		rcp.SaveMetadata(metadataBuf)
		rcp.QueueCheckpointForDeletion(newestId)
		newestId = metadata.Newest.Unmarshal()
	}

	// now only a single checkpoint remains
	onChainId, err := client.BlockIdForHeight(ctx, newestId.Height)
	if err != nil {
		return
	}
	if !onChainId.HeaderHash.Equals(newestId.HeaderHash) {
		rcp.DeleteMetadata()
		rcp.QueueCheckpointForDeletion(newestId)
	}
}

func (rcp *RollupCheckpointerImpl) QueueOldCheckpointsForDeletion(earliestRollbackPoint *common.TimeBlocks) {
	for {
		metadataBytes := rcp.RestoreMetadata()
		metadataBuf := &structures.CheckpointMetadata{}
		if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
			return
		}
		candidateId := metadataBuf.Oldest.Unmarshal()

		linksBuf := rcp.st.GetData(getLinksKey(candidateId))
		links := &structures.CheckpointLinks{}
		if err := proto.Unmarshal(linksBuf, links); err != nil {
			return
		}

		nextHeight := links.Next.Height.Unmarshal()
		if nextHeight.Cmp(earliestRollbackPoint) >= 0 {
			return
		}

		metadataBuf.Oldest = links.Next
		metadataBytes, err := proto.Marshal(metadataBuf)
		if err != nil {
			return
		}

		rcp.QueueCheckpointForDeletion(candidateId)
	}
}

func (rcp *RollupCheckpointerImpl) deleteSomeOldCheckpoints() {
	queueBytes := rcp.st.GetData([]byte("deadqueue"))
	queue := &structures.BlockIdBufList{}
	if err := proto.Unmarshal(queueBytes, queue); err != nil {
		return
	}
	numInQueue := len(queue.Bufs)
	numToDelete := numInQueue / 10
	if numToDelete == 0 && numInQueue > 0 {
		numToDelete = 1
	}

	for i := 0; i < numToDelete; i++ {
		blockId := queue.Bufs[0].Unmarshal()
		rcp.DeleteOneOldCheckpoint(blockId)
		queue.Bufs = queue.Bufs[1:]
	}

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return
	}
	rcp.st.SaveData([]byte("deadqueue"), queueBytes)
}

func (rcp *RollupCheckpointerImpl) DeleteOneOldCheckpoint(blockId *structures.BlockId) {
	// assume metadata has already been updated to reflect deletion
	manifestBytes := rcp.st.GetData(getManifestKey(blockId))
	if manifestBytes == nil {
		return
	}
	manifestBuf := &structures.CheckpointManifest{}
	if err := proto.Unmarshal(manifestBytes, manifestBuf); err != nil {
		return
	}
	rcp.st.DeleteData(getManifestKey(blockId))
	for _, vbuf := range manifestBuf.Values {
		valhash := vbuf.Unmarshal()
		rcp.st.DeleteValue(valhash)
	}
	for _, mbuf := range manifestBuf.Machines {
		machhash := mbuf.Unmarshal()
		rcp.st.DeleteCheckpoint(machhash)
	}
	rcp.st.DeleteData(getContentsKey(blockId))
	rcp.st.DeleteData(getLinksKey(blockId))
}

func (rcp *RollupCheckpointerImpl) GetValue(h common.Hash) value.Value {
	return rcp.st.GetValue(h)
}

func (rcp *RollupCheckpointerImpl) GetMachine(h common.Hash) machine.Machine {
	ret, err := rcp.st.GetInitialMachine()
	if err != nil {
		log.Fatal(err)
	}
	if ret.Hash() == h {
		return ret
	}
	restored := ret.RestoreCheckpoint(rcp.st, h)
	if !restored {
		log.Fatalln("Failed to restore machine", h, "from checkpoint")
	}
	if ret.Hash() != h {
		log.Fatalln("Restore machine", h, "from checkpoint with wrong hash", ret.Hash())
	}
	return ret
}

func (rcp *RollupCheckpointerImpl) DeleteMetadata() {
	rcp.st.DeleteData([]byte("metadata"))
}
