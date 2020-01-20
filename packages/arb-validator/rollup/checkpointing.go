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

package rollup

import (
	"context"
	"log"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"

	"github.com/gogo/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type RollupCheckpointer interface {
	RestoreLatestState(context.Context, arbbridge.ArbClient, common.Address, bool) (blockId *structures.BlockId, content *ChainObserverBuf, resCtx structures.RestoreContext, err error)
	GetInitialMachine() (machine.Machine, error)
	AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{})
}

type DummyCheckpointer struct {
	initialMachine machine.Machine
}

func NewDummyCheckpointer(arbitrumCodefilePath string) RollupCheckpointer {
	theMachine, err := loader.LoadMachineFromFile(arbitrumCodefilePath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", arbitrumCodefilePath)
	}
	return &DummyCheckpointer{theMachine}
}

func (dcp *DummyCheckpointer) RestoreLatestState(ctx context.Context, client arbbridge.ArbClient, contractAddr common.Address, beOpinionated bool) (*structures.BlockId, *ChainObserverBuf, structures.RestoreContext, error) {
	blockId := &structures.BlockId{common.NewTimeBlocks(big.NewInt(0)), common.Hash{}}
	cob := &ChainObserverBuf{}
	resCtx := structures.NewSimpleRestoreContext()
	resCtx.AddMachine(dcp.initialMachine)
	return blockId, cob, resCtx, nil
}

func (dcp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return dcp.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{}) {
	if closeWhenDone != nil {
		closeWhenDone <- struct{}{}
	}
}

type ProductionCheckpointer struct {
	maxReorgDepth *big.Int
	cp            checkpointerWithMetadata
	asyncWriter   *AsyncCheckpointWriter
}

const checkpointDatabasePathBase = "/tmp/arb-validator-checkpoint-"

func makeCheckpointDatabasePath(rollupAddr common.Address, dbPrefix string) string {
	return checkpointDatabasePathBase + dbPrefix + rollupAddr.Hex()[2:]
}

func NewProductionCheckpointer(
	ctx context.Context,
	rollupAddr common.Address,
	arbitrumCodeFilePath string,
	maxReorgDepth *big.Int,
	dbPrefix string,
	forceFreshStart bool, // this should be false in production use
) *ProductionCheckpointer {
	databasePath := makeCheckpointDatabasePath(rollupAddr, dbPrefix)
	if forceFreshStart {
		// for testing only -- use production checkpointer but delete old database first
		if err := os.RemoveAll(databasePath); err != nil {
			log.Fatal(err)
		}
	}
	ret := &ProductionCheckpointer{
		maxReorgDepth: maxReorgDepth,
		cp:            newProductionCheckpointer(databasePath, arbitrumCodeFilePath, maxReorgDepth),
	}
	ret.asyncWriter = NewAsyncCheckpointWriter(ctx, ret)
	return ret
}

func (rcp *ProductionCheckpointer) _saveCheckpoint(
	id *structures.BlockId,
	contents []byte,
	checkpointCtx structures.CheckpointContext,
) error {
	// read in metadata
	var metadataBuf *structures.CheckpointMetadata
	var newestInCp *structures.BlockId
	rawMetadata := rcp.cp.RestoreMetadata()

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
		rcp.cp.SaveMetadata(buf)
	} else {
		metadataBuf = &structures.CheckpointMetadata{}
		if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
			return err
		}
	}
	newestInCp = metadataBuf.Newest.Unmarshal()
	// save all of the data for this checkpoint
	rcp.cp.SaveCheckpoint(
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
	rcp.cp.SaveMetadata(buf)

	return nil
}

func (rcp *ProductionCheckpointer) RestoreLatestState(
	ctx context.Context,
	client arbbridge.ArbClient,
	contractAddr common.Address,
	beOpinionated bool,
) (*structures.BlockId, *ChainObserverBuf, structures.RestoreContext, error) {
	rcp.cp.QueueReorgedCheckpointsForDeletion(ctx, client)

	metadataBytes := rcp.cp.RestoreMetadata()
	if metadataBytes == nil || len(metadataBytes) == 0 {
		rollupWatcher, err := client.NewRollupWatcher(contractAddr)
		if err != nil {
			return nil, nil, nil, err
		}
		params, err := rollupWatcher.GetParams(ctx)
		if err != nil {
			return nil, nil, nil, err
		}
		blockId, err := rollupWatcher.GetCreationHeight(ctx)
		if err != nil {
			return nil, nil, nil, err
		}

		initMachine, err := rcp.GetInitialMachine()
		if err != nil {
			return nil, nil, nil, err
		}
		cob := MakeInitialChainObserverBuf(contractAddr, initMachine.Hash(), &params, beOpinionated)
		resCtx := structures.NewSimpleRestoreContext()
		resCtx.AddMachine(initMachine)
		return blockId, cob, resCtx, nil
	}
	metadata := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadata); err != nil {
		return nil, nil, nil, err
	}
	newestId := metadata.Newest.Unmarshal()
	cobBytes, resCtx, err := rcp.RestoreCheckpoint(newestId)
	if err != nil {
		return nil, nil, nil, err
	}
	cob := &ChainObserverBuf{}
	if err := proto.Unmarshal(cobBytes, cob); err != nil {
		return nil, nil, nil, err
	}
	return newestId, cob, resCtx, nil
}

func (rcp *ProductionCheckpointer) RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext, error) {
	var metadataBuf *structures.CheckpointMetadata
	var oldestHeightInCp *common.TimeBlocks
	var newestHeightInCp *common.TimeBlocks
	rawMetadata := rcp.cp.RestoreMetadata()
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

	buf, checkpointCtx := rcp.cp.RestoreCheckpoint(blockId)
	return buf, checkpointCtx, nil
}

func (cp *ProductionCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.cp.GetInitialMachine()
}

func (cp *ProductionCheckpointer) AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{}) {
	cp.asyncWriter.SubmitJob(
		func() {
			cp._saveCheckpoint(blockId, contents, cpCtx)
		},
		closeWhenDone,
	)
}

func (cp *ProductionCheckpointer) Close() {
	cp.cp.Close()
}

type AsyncCheckpointWriter struct {
	*sync.Mutex
	checkpointer *ProductionCheckpointer
	notifyChan   chan interface{}
	nextJob      func()
	doneChans    []chan struct{}
}

func NewAsyncCheckpointWriter(ctx context.Context, cp *ProductionCheckpointer) *AsyncCheckpointWriter {
	ret := &AsyncCheckpointWriter{&sync.Mutex{}, cp, make(chan interface{}, 1), nil, nil}
	go func() {
		deleteTicker := time.NewTicker(time.Minute)
		defer deleteTicker.Stop()
		for {
			select {
			case <-ret.notifyChan:
				ret.Lock()
				job := ret.nextJob
				if job != nil {
					ret.nextJob = nil
				}
				doneChansCopy := append([]chan struct{}{}, ret.doneChans...)
				ret.Unlock()
				if job != nil {
					job()
				}
				ret.Lock()
				for _, dc := range doneChansCopy {
					if dc != nil {
						close(dc)
					}
				}
				ret.Unlock()
			case <-deleteTicker.C:
				ret.Lock()
				ret.checkpointer.cp.(*productionCheckpointer).deleteSomeOldCheckpoints()
				ret.Unlock()
			case <-ctx.Done():
				ret.checkpointer.Close() //BUGBUG: must ensure this finishes before allowing db to be reopened
				return
			}
		}
	}()
	return ret
}

func (acw *AsyncCheckpointWriter) SubmitJob(job func(), doneChan chan struct{}) {
	acw.Lock()
	defer acw.Unlock()
	acw.nextJob = job
	acw.doneChans = append(acw.doneChans, doneChan)
	select {
	case acw.notifyChan <- nil: // do nothing; only purpose was to send on the channel
	default: // no need to do anything, because channel already has something in it
	}
}

type checkpointerWithMetadata interface {
	SaveMetadata([]byte)
	RestoreMetadata() []byte
	SaveCheckpoint(
		blockId *structures.BlockId,
		prevBlockId *structures.BlockId,
		contents []byte,
		manifest *structures.CheckpointManifest,
		values map[common.Hash]value.Value,
		machines map[common.Hash]machine.Machine,
	)
	RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext) // returns nil, nil if no data at blockHeight
	QueueOldCheckpointsForDeletion(earliestRollbackPoint *common.TimeBlocks)
	QueueReorgedCheckpointsForDeletion(ctx context.Context, client arbbridge.ArbClient)
	QueueCheckpointForDeletion(blockId *structures.BlockId)

	GetInitialMachine() (machine.Machine, error)

	Close()
}

type dummyCheckpointer struct {
	metadata       []byte
	cp             map[*structures.BlockId]*dummyCheckpoint
	initialMachine machine.Machine
}

func newDummyCheckpointer(contractPath string) *dummyCheckpointer {
	theMachine, err := loader.LoadMachineFromFile(contractPath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", contractPath)
	}
	return &dummyCheckpointer{
		nil,
		make(map[*structures.BlockId]*dummyCheckpoint),
		theMachine,
	}
}

type dummyCheckpoint struct {
	contents []byte
	manifest *structures.CheckpointManifest
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func (dcp *dummyCheckpoint) GetValue(h common.Hash) value.Value {
	return dcp.values[h]
}

func (dcp *dummyCheckpoint) GetMachine(h common.Hash) machine.Machine {
	return dcp.machines[h]
}

func (cp *dummyCheckpointer) SaveMetadata(data []byte) {
	cp.metadata = append([]byte{}, data...)
}

func (cp *dummyCheckpointer) RestoreMetadata() []byte {
	return append([]byte{}, cp.metadata...)
}

func (cp *dummyCheckpointer) SaveCheckpoint(
	id *structures.BlockId,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[common.Hash]value.Value,
	machines map[common.Hash]machine.Machine,
) {
	cp.cp[id] = &dummyCheckpoint{contents, manifest, values, machines}
}

func (cp *dummyCheckpointer) RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext) {
	dcp := cp.cp[blockId]
	if dcp == nil {
		return nil, nil
	} else {
		return dcp.contents, dcp
	}
}

func (cp *dummyCheckpointer) DeleteOldCheckpoints(earliestRollbackPoint *big.Int) {
	// ignore cleanup requests; we're being simple and inefficient for debugging
}

func (cp *dummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.initialMachine.Clone(), nil
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

type productionCheckpointer struct {
	st            machine.CheckpointStorage
	maxReorgDepth *big.Int
}

func newProductionCheckpointer(dbpath, contractpath string, maxReorgDepth *big.Int) *productionCheckpointer {
	checkpoint, err := cmachine.NewCheckpoint(dbpath, contractpath)
	if err != nil {
		log.Fatal(err)
	}
	return &productionCheckpointer{checkpoint, maxReorgDepth}
}

func (csc *productionCheckpointer) SaveMetadata(data []byte) {
	ok := csc.st.SaveData([]byte("metadata"), data)
	if !ok {
		log.Fatal("metadata checkpointing failure")
	}
}

func (csc *productionCheckpointer) RestoreMetadata() []byte {
	return csc.st.GetData([]byte("metadata"))
}

func (csc *productionCheckpointer) SaveCheckpoint(
	blockId *structures.BlockId,
	prevBlockId *structures.BlockId,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[common.Hash]value.Value,
	machines map[common.Hash]machine.Machine,
) {
	for _, val := range values {
		csc.st.SaveValue(val)
	}

	for _, mach := range machines {
		savedMachine := mach.Checkpoint(csc.st)
		if !savedMachine {
			log.Fatalln("Failed to checkpoint machine")
		}
	}

	manifestBuf, err := proto.Marshal(manifest)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(getManifestKey(blockId), manifestBuf)

	csc.st.SaveData(getContentsKey(blockId), contents)

	csc._updateNextPointer(prevBlockId, blockId)
	csc._setBothPointers(blockId, prevBlockId, blockId)
}

func (csc *productionCheckpointer) _setBothPointers(id, prev, next *structures.BlockId) {
	links := &structures.CheckpointLinks{
		Prev: prev.MarshalToBuf(),
		Next: next.MarshalToBuf(),
	}
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(getLinksKey(id), linksBuf)
}

func (csc *productionCheckpointer) _updatePrevPointer(id, prev *structures.BlockId) {
	key := getLinksKey(id)
	linksBuf := csc.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Prev = prev.MarshalToBuf()
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(key, linksBuf)
}

func (csc *productionCheckpointer) _updateNextPointer(id, next *structures.BlockId) {
	key := getLinksKey(id)
	linksBuf := csc.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Next = next.MarshalToBuf()
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(key, linksBuf)
}

func (csc *productionCheckpointer) RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext) { // returns nil, nil if no data at blockHeight
	// check for consistency with metadata
	metadataBytes := csc.RestoreMetadata()
	metadataBuf := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
		log.Fatal(err)
	}
	oldestHeight := metadataBuf.Oldest.Height.Unmarshal()
	newestHeight := metadataBuf.Newest.Height.Unmarshal()
	blockHeight := blockId.Height
	if blockHeight.Cmp(oldestHeight) < 0 || blockHeight.Cmp(newestHeight) > 0 {
		return nil, nil
	}

	// read contents
	contentsKey := getContentsKey(blockId)
	contentBytes := csc.st.GetData(contentsKey)

	return contentBytes, csc
}

func (csc *productionCheckpointer) QueueCheckpointForDeletion(blockId *structures.BlockId) {
	// make a best effort to delete an old checkpoint, but ignore any errors
	// errors might cause some harmless extra info to remain in the database

	queueBytes := csc.st.GetData([]byte("deadqueue"))
	queue := &structures.BlockIdBufList{}
	if err := proto.Unmarshal(queueBytes, queue); err != nil {
		return
	}

	queue.Bufs = append(queue.Bufs, blockId.MarshalToBuf())

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return
	}
	csc.st.SaveData([]byte("deadqueue"), queueBytes)
}

func (csc *productionCheckpointer) QueueReorgedCheckpointsForDeletion(ctx context.Context, client arbbridge.ArbClient) {
	metadataBuf := csc.RestoreMetadata()
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
		linksBytes := csc.st.GetData(getLinksKey(newestId))
		linksBuf := &structures.CheckpointLinks{}
		if err := proto.Unmarshal(linksBytes, linksBuf); err != nil {
			return
		}
		metadata.Newest = linksBuf.Prev
		metadataBuf, err = proto.Marshal(metadata)
		if err != nil {
			return
		}
		csc.SaveMetadata(metadataBuf)
		csc.QueueCheckpointForDeletion(newestId)
		newestId = metadata.Newest.Unmarshal()
	}

	// now only a single checkpoint remains
	onChainId, err := client.BlockIdForHeight(ctx, newestId.Height)
	if err != nil {
		return
	}
	if !onChainId.HeaderHash.Equals(newestId.HeaderHash) {
		csc.DeleteMetadata()
		csc.QueueCheckpointForDeletion(newestId)
	}
}

func (csc *productionCheckpointer) QueueOldCheckpointsForDeletion(earliestRollbackPoint *common.TimeBlocks) {
	for {
		metadataBytes := csc.RestoreMetadata()
		metadataBuf := &structures.CheckpointMetadata{}
		if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
			return
		}
		candidateId := metadataBuf.Oldest.Unmarshal()

		linksBuf := csc.st.GetData(getLinksKey(candidateId))
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

		csc.QueueCheckpointForDeletion(candidateId)
	}
}

func (csc *productionCheckpointer) deleteSomeOldCheckpoints() {
	queueBytes := csc.st.GetData([]byte("deadqueue"))
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
		csc.DeleteOneOldCheckpoint(blockId)
		queue.Bufs = queue.Bufs[1:]
	}

	queueBytes, err := proto.Marshal(queue)
	if err != nil {
		return
	}
	csc.st.SaveData([]byte("deadqueue"), queueBytes)
}

func (csc *productionCheckpointer) DeleteOneOldCheckpoint(blockId *structures.BlockId) {
	// assume metadata has already been updated to reflect deletion
	manifestBytes := csc.st.GetData(getManifestKey(blockId))
	if manifestBytes == nil {
		return
	}
	manifestBuf := &structures.CheckpointManifest{}
	if err := proto.Unmarshal(manifestBytes, manifestBuf); err != nil {
		return
	}
	csc.st.DeleteData(getManifestKey(blockId))
	for _, vbuf := range manifestBuf.Values {
		valhash := vbuf.Unmarshal()
		csc.st.DeleteValue(valhash)
	}
	for _, mbuf := range manifestBuf.Machines {
		machhash := mbuf.Unmarshal()
		csc.st.DeleteCheckpoint(machhash)
	}
	csc.st.DeleteData(getContentsKey(blockId))
	csc.st.DeleteData(getLinksKey(blockId))
}

func (csc *productionCheckpointer) GetValue(h common.Hash) value.Value {
	return csc.st.GetValue(h)
}

func (csc *productionCheckpointer) GetMachine(h common.Hash) machine.Machine {
	ret, err := csc.st.GetInitialMachine()
	if err != nil {
		log.Fatal(err)
	}
	if ret.Hash() == h {
		return ret
	}
	restored := ret.RestoreCheckpoint(csc.st, h)
	if !restored {
		log.Fatalln("Failed to restore machine", h, "from checkpoint")
	}
	if ret.Hash() != h {
		log.Fatalln("Restore machine", h, "from checkpoint with wrong hash", ret.Hash())
	}
	return ret
}

func (csc *productionCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return csc.st.GetInitialMachine()
}

func (csc *productionCheckpointer) Close() {
	csc.st.CloseCheckpointStorage()
}

func (csc *productionCheckpointer) DeleteMetadata() {
	csc.st.DeleteData([]byte("metadata"))
}
