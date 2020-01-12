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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"

	"github.com/gogo/protobuf/proto"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type RollupCheckpointer interface {
	RestoreLatestState(ethcommon.Address, *structures.ChainParams, bool) (blockHeight *common.TimeBlocks, content *ChainObserverBuf, resCtx structures.RestoreContext)
	GetInitialMachine() (machine.Machine, error)
	AsyncSaveCheckpoint(
		blockHeight *common.TimeBlocks,
		blockHeaderHash [32]byte,
		contents []byte,
		cpCtx structures.CheckpointContext,
		closeWhenDone chan interface{},
	)
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

func (dcp *DummyCheckpointer) RestoreLatestState(
	contractAddr ethcommon.Address,
	params *structures.ChainParams,
	beOpinionated bool,
) (*common.TimeBlocks, *ChainObserverBuf, structures.RestoreContext) {
	blockHeight := common.NewTimeBlocks(big.NewInt(0))
	cob := &ChainObserverBuf{}
	resCtx := structures.NewSimpleRestoreContext()
	resCtx.AddMachine(dcp.initialMachine)
	return blockHeight, cob, resCtx
}

func (dcp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return dcp.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(
	blockNum *common.TimeBlocks,
	blockHeaderHash [32]byte,
	contents []byte,
	cpCtx structures.CheckpointContext,
	doneChan chan interface{},
) {
	if doneChan != nil {
		doneChan <- struct{}{}
	}
}

type ProductionCheckpointer struct {
	maxReorgDepth *big.Int
	cp            checkpointerWithMetadata
	asyncWriter   *AsyncCheckpointWriter
}

const checkpointDatabasePathBase = "/tmp/arb-validator-checkpoint-"

func makeCheckpointDatabasePath(rollupAddr ethcommon.Address) string {
	return checkpointDatabasePathBase + rollupAddr.Hex()[2:]
}

func NewProductionCheckpointer(
	ctx context.Context,
	rollupAddr ethcommon.Address,
	arbitrumCodeFilePath string,
	maxReorgDepth *big.Int,
	forceFreshStart bool, // this should be false in production use
) RollupCheckpointer {
	databasePath := makeCheckpointDatabasePath(rollupAddr)
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
	blockHeight *big.Int,
	blockHeaderHash [32]byte,
	contents []byte,
	checkpointCtx structures.CheckpointContext,
) error {
	// read in metadata
	var metadataBuf *structures.CheckpointMetadata
	var newestInCp *structures.BlockIdBuf
	rawMetadata := rcp.cp.RestoreMetadata()

	// read in metadata, or create it if it doesn't already exist
	if rawMetadata == nil || len(rawMetadata) == 0 {
		idBuf := &structures.BlockIdBuf{
			Height:     common.MarshalBigInt(blockHeight),
			HeaderHash: common.MarshalHash(blockHeaderHash),
		}
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
	newestInCp = metadataBuf.Newest

	// save all of the data for this checkpoint
	blockId := &structures.BlockIdBuf{
		Height:     common.MarshalBigInt(blockHeight),
		HeaderHash: common.MarshalHash(blockHeaderHash),
	}
	rcp.cp.SaveCheckpoint(
		blockId,
		newestInCp,
		contents,
		checkpointCtx.Manifest(),
		checkpointCtx.Values(),
		checkpointCtx.Machines(),
	)

	// update the metadata to include this checkpoint
	newId := &structures.BlockIdBuf{
		Height:     common.MarshalBigInt(blockHeight),
		HeaderHash: common.MarshalHash(blockHeaderHash),
	}
	metadataBuf.Newest = newId
	buf, err := proto.Marshal(metadataBuf)
	if err != nil {
		return err
	}
	rcp.cp.SaveMetadata(buf)

	return nil
}

func (rcp *ProductionCheckpointer) RestoreLatestState(
	contractAddr ethcommon.Address,
	params *structures.ChainParams,
	beOpinionated bool,
) (*common.TimeBlocks, *ChainObserverBuf, structures.RestoreContext) {
	metadataBytes := rcp.cp.RestoreMetadata()
	if metadataBytes == nil || len(metadataBytes) == 0 {
		initMachine, err := rcp.GetInitialMachine()
		if err != nil {
			return nil, nil, nil
		}
		blockHeight := common.NewTimeBlocks(big.NewInt(0))
		cob := MakeInitialChainObserverBuf(contractAddr, initMachine.Hash(), params, beOpinionated)
		resCtx := structures.NewSimpleRestoreContext()
		resCtx.AddMachine(initMachine)
		return blockHeight, cob, resCtx
	}
	metadata := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadata); err != nil {
		return nil, nil, nil
	}
	newestId := metadata.Newest
	cobBytes, resCtx, err := rcp.RestoreCheckpoint(newestId)
	if err != nil {
		return nil, nil, nil
	}
	cob := &ChainObserverBuf{}
	if err := proto.Unmarshal(cobBytes, cob); err != nil {
		return nil, nil, nil
	}
	return common.NewTimeBlocks(common.UnmarshalBigInt(newestId.Height)), cob, resCtx
}

func (rcp *ProductionCheckpointer) RestoreCheckpoint(blockId *structures.BlockIdBuf) ([]byte, structures.RestoreContext, error) {
	var metadataBuf *structures.CheckpointMetadata
	var oldestHeightInCp *big.Int
	var newestHeightInCp *big.Int
	rawMetadata := rcp.cp.RestoreMetadata()
	if rawMetadata == nil {
		return nil, nil, nil
	}

	metadataBuf = &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
		return nil, nil, err
	}
	oldestHeightInCp = common.UnmarshalBigInt(metadataBuf.Oldest.Height)
	newestHeightInCp = common.UnmarshalBigInt(metadataBuf.Newest.Height)

	blockHeight := common.UnmarshalBigInt(blockId.Height)
	if blockHeight.Cmp(oldestHeightInCp) < 0 || blockHeight.Cmp(newestHeightInCp) > 0 {
		return nil, nil, nil
	}

	buf, checkpointCtx := rcp.cp.RestoreCheckpoint(blockId)
	return buf, checkpointCtx, nil
}

func (cp *ProductionCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.cp.GetInitialMachine()
}

func (cp *ProductionCheckpointer) AsyncSaveCheckpoint(
	blocknum *common.TimeBlocks,
	blockHeaderHash [32]byte,
	buf []byte,
	cpCtx structures.CheckpointContext,
	doneChan chan interface{},
) {
	cp.asyncWriter.SubmitJob(
		func() {
			cp._saveCheckpoint(blocknum.AsInt(), blockHeaderHash, buf, cpCtx)
		},
		doneChan,
	)
}

type AsyncCheckpointWriter struct {
	*sync.Mutex
	checkpointer *ProductionCheckpointer
	notifyChan   chan interface{}
	nextJob      func()
	doneChans    []chan interface{}
}

func NewAsyncCheckpointWriter(ctx context.Context, cp *ProductionCheckpointer) *AsyncCheckpointWriter {
	ret := &AsyncCheckpointWriter{&sync.Mutex{}, cp, make(chan interface{}, 1), nil, nil}
	go func() {
		for {
			select {
			case <-ret.notifyChan:
				ret.Lock()
				job := ret.nextJob
				if job != nil {
					ret.nextJob = nil
				}
				doneChansCopy := append([]chan interface{}{}, ret.doneChans...)
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
			case <-ctx.Done():
				return
			}
		}
	}()
	return ret
}

func (acw *AsyncCheckpointWriter) SubmitJob(job func(), doneChan chan interface{}) {
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
		blockId *structures.BlockIdBuf,
		prevBlockId *structures.BlockIdBuf,
		contents []byte,
		manifest *structures.CheckpointManifest,
		values map[[32]byte]value.Value,
		machines map[[32]byte]machine.Machine,
	)
	RestoreCheckpoint(blockId *structures.BlockIdBuf) ([]byte, structures.RestoreContext) // returns nil, nil if no data at blockHeight
	DeleteOldCheckpoints(earliestRollbackPoint *big.Int)

	GetInitialMachine() (machine.Machine, error)
}

type dummyCheckpointer struct {
	metadata       []byte
	cp             map[*big.Int]*dummyCheckpoint
	initialMachine machine.Machine
}

func newDummyCheckpointer(contractPath string) *dummyCheckpointer {
	theMachine, err := loader.LoadMachineFromFile(contractPath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", contractPath)
	}
	return &dummyCheckpointer{
		nil,
		make(map[*big.Int]*dummyCheckpoint),
		theMachine,
	}
}

type dummyCheckpoint struct {
	contents []byte
	manifest *structures.CheckpointManifest
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

func (dcp *dummyCheckpoint) GetValue(h [32]byte) value.Value {
	return dcp.values[h]
}

func (dcp *dummyCheckpoint) GetMachine(h [32]byte) machine.Machine {
	return dcp.machines[h]
}

func (cp *dummyCheckpointer) SaveMetadata(data []byte) {
	cp.metadata = append([]byte{}, data...)
}

func (cp *dummyCheckpointer) RestoreMetadata() []byte {
	return append([]byte{}, cp.metadata...)
}

func (cp *dummyCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	prevBlockHeight *big.Int,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[[32]byte]value.Value,
	machines map[[32]byte]machine.Machine,
) {
	cp.cp[blockHeight] = &dummyCheckpoint{contents, manifest, values, machines}
}

func (cp *dummyCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, structures.RestoreContext) {
	dcp := cp.cp[blockHeight]
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

func getKeyForId(prefix []byte, id *structures.BlockIdBuf) []byte {
	idBytes, err := proto.Marshal(id)
	if err != nil {
		log.Fatal(err)
	}
	return append(prefix, idBytes...)
}

func getManifestKey(blockId *structures.BlockIdBuf) []byte {
	return getKeyForId([]byte("manifest:"), blockId)
}

func getContentsKey(blockId *structures.BlockIdBuf) []byte {
	return getKeyForId([]byte("contents:"), blockId)
}

func getLinksKey(blockId *structures.BlockIdBuf) []byte {
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
	blockId *structures.BlockIdBuf,
	prevBlockId *structures.BlockIdBuf,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[[32]byte]value.Value,
	machines map[[32]byte]machine.Machine,
) {
	for _, val := range values {
		csc.st.SaveValue(val)
	}

	for _, mach := range machines {
		mach.Checkpoint(csc.st)
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

func (csc *productionCheckpointer) _setBothPointers(id, prev, next *structures.BlockIdBuf) {
	links := &structures.CheckpointLinks{
		Prev: prev,
		Next: next,
	}
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(getLinksKey(id), linksBuf)
}

func (csc *productionCheckpointer) _updatePrevPointer(id, prev *structures.BlockIdBuf) {
	key := getLinksKey(id)
	linksBuf := csc.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Prev = prev
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(key, linksBuf)
}

func (csc *productionCheckpointer) _updateNextPointer(id, next *structures.BlockIdBuf) {
	key := getLinksKey(id)
	linksBuf := csc.st.GetData(key)
	links := &structures.CheckpointLinks{}
	if err := proto.Unmarshal(linksBuf, links); err != nil {
		log.Fatal(err)
	}
	links.Next = next
	linksBuf, err := proto.Marshal(links)
	if err != nil {
		log.Fatal(err)
	}
	csc.st.SaveData(key, linksBuf)
}

func (csc *productionCheckpointer) RestoreCheckpoint(blockId *structures.BlockIdBuf) ([]byte, structures.RestoreContext) { // returns nil, nil if no data at blockHeight
	// check for consistency with metadata
	metadataBytes := csc.RestoreMetadata()
	metadataBuf := &structures.CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
		log.Fatal(err)
	}
	oldestHeight := common.UnmarshalBigInt(metadataBuf.Oldest.Height)
	newestHeight := common.UnmarshalBigInt(metadataBuf.Newest.Height)
	blockHeight := common.UnmarshalBigInt(blockId.Height)
	if blockHeight.Cmp(oldestHeight) < 0 || blockHeight.Cmp(newestHeight) > 0 {
		return nil, nil
	}

	// read contents
	contentsKey := getContentsKey(blockId)
	contentBytes := csc.st.GetData(contentsKey)

	return contentBytes, csc
}

func (csc *productionCheckpointer) DeleteOldCheckpoints(earliestRollbackPoint *big.Int) {
	// make a best effort to delete an old checkpoint, but ignore any errors
	// errors might cause some harmless extra info to remain in the database

	for {
		metadataBytes := csc.RestoreMetadata()
		metadataBuf := &structures.CheckpointMetadata{}
		if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
			return
		}

		linksBuf := csc.st.GetData(getLinksKey(metadataBuf.Oldest))
		links := &structures.CheckpointLinks{}
		if err := proto.Unmarshal(linksBuf, links); err != nil {
			return
		}

		nextHeight := common.UnmarshalBigInt(links.Next.Height)
		if nextHeight.Cmp(earliestRollbackPoint) >= 0 {
			return
		}

		metadataBuf.Newest.Height = common.MarshalBigInt(nextHeight)
		metadataBytes, err := proto.Marshal(metadataBuf)
		if err != nil {
			return
		}

		csc.DeleteOneOldCheckpoint(metadataBuf.Oldest)
	}
}

func (csc *productionCheckpointer) DeleteOneOldCheckpoint(blockId *structures.BlockIdBuf) {
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
		valhash := common.UnmarshalHash(vbuf)
		csc.st.DeleteValue(valhash)
	}
	for _, mbuf := range manifestBuf.Machines {
		machhash := common.UnmarshalHash(mbuf)
		csc.st.DeleteCheckpoint(machhash)
	}
	csc.st.DeleteData(getContentsKey(blockId))
}

func (csc *productionCheckpointer) GetValue(h [32]byte) value.Value {
	return csc.st.GetValue(h)
}

func (csc *productionCheckpointer) GetMachine(h [32]byte) machine.Machine {
	ret, err := csc.st.GetInitialMachine()
	if err != nil {
		log.Fatal(err)
	}
	ret.RestoreCheckpoint(csc.st, h)
	return ret
}

func (csc *productionCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return csc.st.GetInitialMachine()
}
