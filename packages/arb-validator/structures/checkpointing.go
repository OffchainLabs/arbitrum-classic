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

package structures

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"log"
	"math/big"
	"os"
)

type CheckpointContext interface {
	AddValue(value.Value)
	AddMachine(machine.Machine)
	Manifest() *CheckpointManifest
	Values() map[[32]byte]value.Value
	Machines() map[[32]byte]machine.Machine
}

type CheckpointContextImpl struct {
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

type RestoreContext interface {
	GetValue([32]byte) value.Value
	GetMachine([32]byte) machine.Machine
}

func NewCheckpointContextImpl() *CheckpointContextImpl {
	return &CheckpointContextImpl{
		values:   make(map[[32]byte]value.Value),
		machines: make(map[[32]byte]machine.Machine),
	}
}

func (ctx *CheckpointContextImpl) AddValue(val value.Value) {
	ctx.values[val.Hash()] = val
}

func (ctx *CheckpointContextImpl) AddMachine(mach machine.Machine) {
	if ctx.machines[mach.Hash()] == nil {
		ctx.machines[mach.Hash()] = mach.Clone()
	}
}

func (ctx *CheckpointContextImpl) Manifest() *CheckpointManifest {
	vals := []*value.HashBuf{}
	for h, _ := range ctx.values {
		vals = append(vals, utils.MarshalHash(h))
	}
	machines := []*value.HashBuf{}
	for h, _ := range ctx.machines {
		machines = append(machines, utils.MarshalHash(h))
	}
	return &CheckpointManifest{Values: vals, Machines: machines}
}

func (ctx *CheckpointContextImpl) Values() map[[32]byte]value.Value {
	return ctx.values
}

func (ctx *CheckpointContextImpl) Machines() map[[32]byte]machine.Machine {
	return ctx.machines
}

func (ctx *CheckpointContextImpl) GetValue(h [32]byte) value.Value {
	return ctx.values[h]
}

func (ctx *CheckpointContextImpl) GetMachine(h [32]byte) machine.Machine {
	return ctx.machines[h]
}

type RollupCheckpointer struct {
	versionsToKeep *big.Int
	cp             CheckpointerWithMetadata
}

const checkpointDatabasePathBase = "/tmp/arb-validator-checkpoint-"

func makeCheckpointDatabasePath(rollupAddr common.Address) string {
	return checkpointDatabasePathBase + rollupAddr.Hex()[2:]
}

func NewRollupCheckpointer(
	rollupAddr common.Address,
	kind string,
	versionsToKeep int64,
	contractPath string,
) *RollupCheckpointer {
	databasePath := makeCheckpointDatabasePath(rollupAddr)
	switch kind {
	case "dummy": // inefficient in-memory checkpointer, for testing
		return &RollupCheckpointer{big.NewInt(versionsToKeep), NewDummyCheckpointer(contractPath)}
	case "fresh_rocksdb": // for testing only -- use rocksdb but delete old database first
		if err := os.RemoveAll(databasePath); err != nil {
			log.Fatal(err)
		}
		fallthrough
	case "": // empty string gives you what you want for production
		fallthrough
	case "rocksdb": // store in rocksdb database, keyed to rollupAddr -- use this for production
		return &RollupCheckpointer{
			big.NewInt(versionsToKeep),
			NewCstoreCheckpointer(databasePath, "contract.ao"),
		}
	default:
		return nil
	}
}

func (rcp *RollupCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	contents []byte,
	ctx CheckpointContext,
) error {
	var metadataBuf *CheckpointMetadata
	var oldestInCp *big.Int
	var newestInCp *big.Int
	rawMetadata := rcp.cp.RestoreMetadata()
	if rawMetadata == nil || len(rawMetadata) == 0 {
		oldestInCp = blockHeight
		newestInCp = blockHeight
		metadataBuf = &CheckpointMetadata{
			FormatVersion:     1,
			OldestBlockHeight: utils.MarshalBigInt(oldestInCp),
			NewestBlockHeight: utils.MarshalBigInt(newestInCp),
		}
		buf, err := proto.Marshal(metadataBuf)
		if err != nil {
			return err
		}
		rcp.cp.SaveMetadata(buf)
	} else {
		metadataBuf = &CheckpointMetadata{}
		if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
			return err
		}
		oldestInCp = utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
		newestInCp = utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)
		if blockHeight.Cmp(newestInCp) > 0 {
			metadataBuf.NewestBlockHeight = utils.MarshalBigInt(blockHeight)
			buf, err := proto.Marshal(metadataBuf)
			if err != nil {
				return err
			}
			rcp.cp.SaveMetadata(buf)
		}
	}
	rcp.cp.SaveCheckpoint(blockHeight, contents, ctx.Manifest(), ctx.Values(), ctx.Machines())

	if oldestInCp.Cmp(new(big.Int).Sub(newestInCp, rcp.versionsToKeep)) < 0 {
		go func() {
			//TODO: clean up old versions
		}()
	}
	return nil
}

func (rcp *RollupCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext, error) {
	var metadataBuf *CheckpointMetadata
	var oldestInCp *big.Int
	var newestInCp *big.Int
	rawMetadata := rcp.cp.RestoreMetadata()
	if rawMetadata == nil {
		return nil, nil, nil
	}

	metadataBuf = &CheckpointMetadata{}
	if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
		return nil, nil, err
	}
	oldestInCp = utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
	newestInCp = utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)

	if blockHeight.Cmp(oldestInCp) < 0 || blockHeight.Cmp(newestInCp) > 0 {
		return nil, nil, nil
	}

	buf, ctx := rcp.cp.RestoreCheckpoint(blockHeight)
	return buf, ctx, nil
}

func (cp *RollupCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.cp.GetInitialMachine()
}

type CheckpointerWithMetadata interface {
	SaveMetadata([]byte)
	RestoreMetadata() []byte
	SaveCheckpoint(
		blockHeight *big.Int,
		contents []byte,
		manifest *CheckpointManifest,
		values map[[32]byte]value.Value,
		machines map[[32]byte]machine.Machine,
	)
	RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext) // returns nil, nil if no data at blockHeight
	DeleteCheckpoint(blockHeight *big.Int)

	GetInitialMachine() (machine.Machine, error)
}

type DummyCheckpointer struct {
	metadata       []byte
	cp             map[*big.Int]*dummyCheckpoint
	initialMachine machine.Machine
}

func NewDummyCheckpointer(contractPath string) *DummyCheckpointer {
	theMachine, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		log.Fatal("NewDummyCheckpointer: error loading ", contractPath)
	}
	return &DummyCheckpointer{
		nil,
		make(map[*big.Int]*dummyCheckpoint),
		theMachine,
	}
}

type dummyCheckpoint struct {
	contents []byte
	manifest *CheckpointManifest
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

func (dcp *dummyCheckpoint) GetValue(h [32]byte) value.Value {
	return dcp.values[h]
}

func (dcp *dummyCheckpoint) GetMachine(h [32]byte) machine.Machine {
	return dcp.machines[h]
}

func (cp *DummyCheckpointer) SaveMetadata(data []byte) {
	cp.metadata = append([]byte{}, data...)
}

func (cp *DummyCheckpointer) RestoreMetadata() []byte {
	return append([]byte{}, cp.metadata...)
}

func (cp *DummyCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	contents []byte,
	manifest *CheckpointManifest,
	values map[[32]byte]value.Value,
	machines map[[32]byte]machine.Machine,
) {
	cp.cp[blockHeight] = &dummyCheckpoint{contents, manifest, values, machines}
}

func (cp *DummyCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext) {
	dcp := cp.cp[blockHeight]
	if dcp == nil {
		return nil, nil
	} else {
		return dcp.contents, dcp
	}
}

func (cp *DummyCheckpointer) DeleteCheckpoint(blockHeight *big.Int) {
	delete(cp.cp, blockHeight)
}

func (cp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.initialMachine.Clone(), nil
}

var metadataKey []byte
var contentsKey []byte

func init() {
	metadataKey = []byte("metadata")
	contentsKey = []byte("contents")
}

func manifestKey(blockHeight *big.Int) []byte {
	bhBytes := blockHeight.Bytes()
	return append([]byte("manifest:"), bhBytes...)
}

type CStoreCheckpointer struct {
	st machine.CheckpointStorage
}

func NewCstoreCheckpointer(dbpath, contractpath string) *CStoreCheckpointer {
	checkpoint, err := cmachine.NewCheckpoint(dbpath, contractpath)
	if err != nil {
		log.Fatal(err)
	}
	return &CStoreCheckpointer{checkpoint}
}

func (csc *CStoreCheckpointer) SaveMetadata(data []byte) {
	ok := csc.st.SaveData(metadataKey, data)
	if !ok {
		log.Fatal("metadata checkpointing failure")
	}
}

func (csc *CStoreCheckpointer) RestoreMetadata() []byte {
	return csc.st.GetData(metadataKey)
}

func (csc *CStoreCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	contents []byte,
	manifest *CheckpointManifest,
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
	csc.st.SaveData(manifestKey(blockHeight), manifestBuf)

	csc.st.SaveData(contentsKey, contents)

	metadataBytes := csc.RestoreMetadata()
	metadataBuf := &CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
		log.Fatal(err)
	}
	newestHeight := utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)
	if blockHeight.Cmp(newestHeight) > 0 {
		metadataBuf.NewestBlockHeight = utils.MarshalBigInt(blockHeight)
		var err error
		metadataBytes, err = proto.Marshal(metadataBuf)
		if err != nil {
			log.Fatal(err)
		}
		csc.SaveMetadata(metadataBytes)
	}
}

func (csc *CStoreCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext) { // returns nil, nil if no data at blockHeight
	// check for consistency with metadata
	metadataBytes := csc.RestoreMetadata()
	metadataBuf := &CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
		log.Fatal(err)
	}
	oldestHeight := utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
	newestHeight := utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)
	if blockHeight.Cmp(oldestHeight) < 0 || blockHeight.Cmp(newestHeight) > 0 {
		return nil, nil
	}

	// read contents
	contentBytes := csc.st.GetData(contentsKey)

	return contentBytes, csc
}

func (csc *CStoreCheckpointer) DeleteCheckpoint(blockHeight *big.Int) {
	// update metadata
	metadataBytes := csc.RestoreMetadata()
	metadataBuf := &CheckpointMetadata{}
	if err := proto.Unmarshal(metadataBytes, metadataBuf); err != nil {
		log.Fatal(err)
	}
	oldestHeight := utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
	newestHeight := utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)
	if blockHeight.Cmp(newestHeight) >= 0 {
		// deleted the last item, so null the metadata
		csc.SaveMetadata([]byte{})
	} else if blockHeight.Cmp(oldestHeight) > 0 {
		metadataBuf.OldestBlockHeight = utils.MarshalBigInt(blockHeight)
		var err error
		metadataBytes, err = proto.Marshal(metadataBuf)
		if err != nil {
			log.Fatal(err)
		}
		csc.SaveMetadata(metadataBytes)
	}

	//TODO: need to clean up no-longer-needed data
	// read manifest
	// delete manifest from DB
	// use manifest to delete values and machines
	// delete contents
}

func (csc *CStoreCheckpointer) GetValue(h [32]byte) value.Value {
	return csc.st.GetValue(h)
}

func (csc *CStoreCheckpointer) GetMachine(h [32]byte) machine.Machine {
	ret, err := csc.st.GetInitialMachine()
	if err != nil {
		log.Fatal(err)
	}
	ret.RestoreCheckpoint(csc.st, h)
	return ret
}

func (csc *CStoreCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return csc.st.GetInitialMachine()
}
