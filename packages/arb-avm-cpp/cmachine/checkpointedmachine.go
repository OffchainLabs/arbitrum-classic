/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package cmachine

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lsecp256k1 -lff -lgmp -lkeccak -ldl
#include "../cavm/ccheckpointedmachine.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"github.com/pkg/errors"
	"runtime"
	"time"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type CheckpointedMachine struct {
	c unsafe.Pointer
}

func NewCheckpointedMachine(codeFile string, dbPath string) (*CheckpointedMachine, error) {
	cFilename := C.CString(codeFile)
	defer C.free(unsafe.Pointer(cFilename))
	cDbPath := C.CString(dbPath)
	defer C.free(unsafe.Pointer(cDbPath))
	cCheckpointedMachine := C.checkpointedMachineCreate(cFilename, cDbPath)
	if cCheckpointedMachine == nil {
		return nil, errors.Errorf("error creating checkpointed machine from file %s and path %s", codeFile, dbPath)
	}
	ret := &CheckpointedMachine{cCheckpointedMachine}
	runtime.SetFinalizer(ret, cdestroyCheckpointedMachine)
	return ret, nil
}

func cdestroyCheckpointedMachine(cCheckpointedMachine *CheckpointedMachine) {
	C.checkpointedMachineDestroy(cCheckpointedMachine.c)
}

func (cm *CheckpointedMachine) Hash() (ret common.Hash) {
	C.machineHash(cm.c, unsafe.Pointer(&ret[0]))
	return
}

func (cm *CheckpointedMachine) Clone() machine.Machine {
	cMachine := C.machineClone(cm.c)
	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret
}

func (cm *CheckpointedMachine) CurrentStatus() machine.Status {
	cStatus := C.machineCurrentStatus(cm.c)
	switch cStatus {
	case C.STATUS_EXTENSIVE:
		return machine.Extensive
	case C.STATUS_ERROR_STOP:
		return machine.ErrorStop

	case C.STATUS_HALT:
		return machine.Halt
	default:
		panic("Unknown status")
	}
}

func (cm *CheckpointedMachine) IsBlocked(newMessages bool) machine.BlockReason {
	newMessagesInt := 0
	if newMessages {
		newMessagesInt = 1
	}
	cBlockReason := C.machineIsBlocked(cm.c, C.int(newMessagesInt))
	switch cBlockReason.blockType {
	case C.BLOCK_TYPE_NOT_BLOCKED:
		return nil
	case C.BLOCK_TYPE_HALT:
		return machine.HaltBlocked{}
	case C.BLOCK_TYPE_ERROR:
		return machine.ErrorBlocked{}
	case C.BLOCK_TYPE_BREAKPOINT:
		return machine.BreakpointBlocked{}
	case C.BLOCK_TYPE_INBOX:
		return machine.InboxBlocked{}
	default:
	}
	return nil
}

func (cm *CheckpointedMachine) ExecuteAssertion(
	maxSteps uint64,
	inboxMessages []inbox.InboxMessage,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	msgDataC := C.CBytes(encodeInboxMessages(inboxMessages))
	defer C.free(msgDataC)

	beforeHash := cm.Hash()
	assertion := C.executeAssertion(
		cm.c,
		C.uint64_t(maxSteps),
		msgDataC,
		C.uint64_t(len(inboxMessages)),
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)

	return makeExecutionAssertion(assertion, beforeHash, cm.Hash())
}

func (cm *CheckpointedMachine) ExecuteCallServerAssertion(
	maxSteps uint64,
	inboxMessages []inbox.InboxMessage,
	fakeInboxPeekValue value.Value,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	msgDataC := C.CBytes(encodeInboxMessages(inboxMessages))
	defer C.free(msgDataC)

	inboxPeekDataC := C.CBytes(encodeValue(fakeInboxPeekValue))
	defer C.free(inboxPeekDataC)

	beforeHash := cm.Hash()
	assertion := C.executeCallServerAssertion(
		cm.c,
		C.uint64_t(maxSteps),
		msgDataC,
		C.uint64_t(len(inboxMessages)),
		inboxPeekDataC,
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)

	return makeExecutionAssertion(assertion, beforeHash, cm.Hash())
}

func (cm *CheckpointedMachine) ExecuteSideloadedAssertion(
	maxSteps uint64,
	inboxMessages []inbox.InboxMessage,
	sideloadValue *value.TupleValue,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	msgDataC := C.CBytes(encodeInboxMessages(inboxMessages))
	defer C.free(msgDataC)

	sideloadDataC := C.CBytes(encodeValue(sideloadValue))
	defer C.free(sideloadDataC)

	beforeHash := cm.Hash()
	assertion := C.executeSideloadedAssertion(
		cm.c,
		C.uint64_t(maxSteps),
		msgDataC,
		C.uint64_t(len(inboxMessages)),
		sideloadDataC,
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)

	return makeExecutionAssertion(assertion, beforeHash, cm.Hash())
}

func (cm *CheckpointedMachine) Initialize(contractPath string) error {
	cContractPath := C.CString(contractPath)
	defer C.free(unsafe.Pointer(cContractPath))
	success := C.initializeCheckpointedMachine(cm.c, cContractPath)

	if success == 0 {
		return errors.New("failed to initialize storage")
	}
	return nil
}

func (cm *CheckpointedMachine) Initialized() bool {
	return C.checkpointStorageInitialized(cm.c) == 1
}

func (cm *CheckpointedMachine) CloseCheckpointedMachine() bool {
	return C.closeCheckpointedMachine(cm.c) == 1
}

func (cm *CheckpointedMachine) GetInitialMachine(valueCache machine.ValueCache) (machine.Machine, error) {
	cMachine := C.getInitialMachine(cm.c, valueCache.(*ValueCache).c)

	if cMachine == nil {
		return nil, errors.Errorf("error getting initial machine from checkpointstorage")
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (cm *CheckpointedMachine) GetMachine(machineHash common.Hash, valueCache machine.ValueCache) (machine.Machine, error) {
	cMachine := C.getMachine(cm.c, unsafe.Pointer(&machineHash[0]), valueCache.(*ValueCache).c)

	if cMachine == nil {
		return nil, &machine.MachineNotFoundError{HashValue: machineHash}
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (cm *CheckpointedMachine) DeleteCheckpoint(machineHash common.Hash) bool {
	success := C.deleteCheckpoint(cm.c, unsafe.Pointer(&machineHash[0]))

	return success == 1
}

func (cm *CheckpointedMachine) SaveValue(val value.Value) bool {
	var buf bytes.Buffer

	err := value.MarshalValue(val, &buf)
	if err != nil {
		panic(err)
	}

	valData := buf.Bytes()
	success := C.saveValue(cm.c, unsafe.Pointer(&valData[0]))

	return success == 1
}

func (cm *CheckpointedMachine) GetValue(hashValue common.Hash, valueCache machine.ValueCache) (value.Value, error) {
	cData := C.getValue(cm.c, unsafe.Pointer(&hashValue[0]), valueCache.(*ValueCache).c)
	if cData.data == nil {
		return nil, &machine.ValueNotFoundError{HashValue: hashValue}
	}

	dataBuff := toByteSlice(cData)

	val, err := value.UnmarshalValue(bytes.NewReader(dataBuff[:]))
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (cm *CheckpointedMachine) DeleteValue(hashValue common.Hash) bool {
	success := C.deleteValue(cm.c, unsafe.Pointer(&hashValue[0]))

	return success == 1
}

func (cm *CheckpointedMachine) SaveData(key []byte, data []byte) bool {
	if len(key) == 0 {
		return false
	}

	if len(data) == 0 {
		success := C.saveData(cm.c,
			unsafe.Pointer(&key[0]),
			C.int(len(key)),
			unsafe.Pointer(nil),
			C.int(0),
		)
		return success == 1
	}

	success := C.saveData(cm.c,
		unsafe.Pointer(&key[0]),
		C.int(len(key)),
		unsafe.Pointer(&data[0]),
		C.int(len(data)))

	return success == 1
}

func (cm *CheckpointedMachine) GetData(key []byte) ([]byte, error) {
	cData := C.getData(cm.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	if cData.found == 0 {
		return nil, &machine.DataNotFoundError{Key: key}
	}

	return toByteSlice(cData.slice), nil
}

func (cm *CheckpointedMachine) DeleteData(key []byte) bool {
	success := C.deleteData(cm.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	return success == 1
}

func (cm *CheckpointedMachine) GetBlockStore() machine.BlockStore {
	bs := C.createBlockStore(cm.c)

	return NewBlockStore(bs)
}

func (cm *CheckpointedMachine) GetAggregatorStore() *AggregatorStore {
	bs := C.createAggregatorStore(cm.c)

	return NewAggregatorStore(bs)
}
