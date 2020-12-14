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
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var errNoCheckpoint = errors.New("cannot restore because no checkpoint exists")
var errNoMatchingCheckpoint = errors.New("cannot restore because no matching checkpoint exists")

type CheckpointedMachine struct {
	*sync.Mutex
	cm                    *cmachine.CheckpointedMachine
	nextCheckpointToWrite *writableCheckpoint
	maxReorgHeight        *big.Int
}

func NewCheckpointedMachine(
	rollupAddr common.Address,
	executablePath string,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) (*CheckpointedMachine, error) {
	ret, err := newCheckpointedMachine(
		rollupAddr,
		executablePath,
		databasePath,
		new(big.Int).Set(maxReorgHeight),
		forceFreshStart,
	)

	if err != nil {
		return nil, err
	}

	go ret.writeDaemon()
	// TODO
	//go cleanupDaemon(ret.bs, ret.db, maxReorgHeight)
	return ret, nil
}

// newCheckpointedMachine creates the checkpointedmachine, but doesn't
// launch it's reading and writing threads. This is useful for deterministic
// testing
func newCheckpointedMachine(
	rollupAddr common.Address,
	executablePath string,
	databasePath string,
	maxReorgHeight *big.Int,
	forceFreshStart bool,
) (*CheckpointedMachine, error) {
	if databasePath == "" {
		databasePath = MakeCheckpointDatabasePath(rollupAddr)
	}
	if forceFreshStart {
		// for testing only --  delete old database to get fresh start
		if err := os.RemoveAll(databasePath); err != nil {
			return nil, err
		}
	}
	cCheckpointedMachine, err := cmachine.NewCheckpointedMachine(executablePath, databasePath)
	if err != nil {
		return nil, err
	}

	return &CheckpointedMachine{
		new(sync.Mutex),
		cCheckpointedMachine,
		nil,
		maxReorgHeight,
	}, nil
}

func (cp *CheckpointedMachine) Initialize(arbitrumCodeFilePath string) error {
	return cp.cm.Initialize(arbitrumCodeFilePath)
}

func (cp *CheckpointedMachine) Initialized() bool {
	return cp.cm.Initialized()
}

func (cp *CheckpointedMachine) MaxReorgHeight() *big.Int {
	return new(big.Int).Set(cp.maxReorgHeight)
}

func (cp *CheckpointedMachine) GetAggregatorStore() *cmachine.AggregatorStore {
	return cp.cm.GetAggregatorStore()
}

// HasCheckpointedState checks whether the block store is empty, which is the table
// which contains the checkpoints recorded by the CheckpointedMachine
func (cp *CheckpointedMachine) HasCheckpointedState() bool {
	// TODO
	return true
}

func (cp *CheckpointedMachine) GetInitialMachine(vc machine.ValueCache) (machine.Machine, error) {
	return cp.cm.GetInitialMachine(vc)
}

func (cp *CheckpointedMachine) AsyncSaveCheckpoint(
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

func (cp *CheckpointedMachine) RestoreLatestState(ctx context.Context, clnt arbbridge.ChainTimeGetter, unmarshalFunc func([]byte, ckptcontext.RestoreContext, *common.BlockId) error) error {
	return restoreLatestState(ctx, cp.cm, clnt, unmarshalFunc)
}

func restoreLatestState(
	ctx context.Context,
	cm *cmachine.CheckpointedMachine,
	clnt arbbridge.ChainTimeGetter,
	unmarshalFunc func([]byte, ckptcontext.RestoreContext, *common.BlockId) error,
) error {
	// TODO
	return errNoMatchingCheckpoint
}

func (cp *CheckpointedMachine) writeDaemon() {
	ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
	defer ticker.Stop()
	for {
		<-ticker.C
		cp.Lock()
		checkpoint := cp.nextCheckpointToWrite
		cp.nextCheckpointToWrite = nil
		cp.Unlock()
		if checkpoint != nil {
			err := writeCheckpoint(cp.cm, checkpoint)
			if err != nil {
				logger.Error().Stack().Err(err).Msg("Error writing checkpoint")
			}
			checkpoint.errChan <- err
			close(checkpoint.errChan)
		}
	}
}

func writeCheckpoint(cm *cmachine.CheckpointedMachine, wc *writableCheckpoint) error {

	// TODO

	return nil
}

func deleteCheckpointForKey(bs machine.BlockStore, db machine.CheckpointStorage, id *common.BlockId) error {
	val, err := bs.GetBlock(id)
	if err != nil {
		return err
	}
	ckp := &CheckpointWithManifest{}
	if err := proto.Unmarshal(val, ckp); err != nil {
		return err
	}
	_ = bs.DeleteBlock(id) // ignore error
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
	db         machine.CheckpointStorage
	values     map[common.Hash]value.Value
	machines   map[common.Hash]machine.Machine
	valueCache machine.ValueCache
}

func newRestoreContextLocked(db machine.CheckpointStorage, manifest *ckptcontext.CheckpointManifest) (*restoreContextLocked, error) {
	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return nil, err
	}

	rcl := restoreContextLocked{db, map[common.Hash]value.Value{}, map[common.Hash]machine.Machine{}, valueCache}

	for _, valHash := range manifest.GetValues() {
		hash := valHash.Unmarshal()
		val, err := rcl.db.GetValue(hash, rcl.valueCache)
		if err != nil {
			return nil, err
		}

		rcl.values[hash] = val
	}

	for _, machHash := range manifest.GetMachines() {
		hash := machHash.Unmarshal()
		mach, err := rcl.db.GetMachine(hash, valueCache)
		if err != nil {
			return nil, err
		}

		rcl.machines[hash] = mach
	}

	return &rcl, nil
}

func (rcl *restoreContextLocked) GetValue(h common.Hash) (value.Value, error) {
	if val, ok := rcl.values[h]; ok {
		return val, nil
	}

	return rcl.db.GetValue(h, rcl.valueCache)
}

func (rcl *restoreContextLocked) GetMachine(h common.Hash) (machine.Machine, error) {
	if mach, ok := rcl.machines[h]; ok {
		return mach, nil
	}

	return rcl.db.GetMachine(h, rcl.valueCache)
}
