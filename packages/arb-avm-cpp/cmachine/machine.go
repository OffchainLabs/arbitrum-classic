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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb
#include "../cavm/cmachine.h"
#include "../cavm/ccheckpointstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	c unsafe.Pointer
}

func New(codeFile string) (*Machine, error) {
	cFilename := C.CString(codeFile)

	cMachine := C.machineCreate(cFilename)
	if cMachine == nil {
		return nil, fmt.Errorf("error loading machine %v", codeFile)
	}
	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	C.free(unsafe.Pointer(cFilename))
	return ret, nil
}

func cdestroyVM(cMachine *Machine) {
	C.machineDestroy(cMachine.c)
}

func (m *Machine) Hash() (ret [32]byte) {
	C.machineHash(m.c, unsafe.Pointer(&ret[0]))
	return
}

func (m *Machine) Clone() machine.Machine {
	cMachine := C.machineClone(m.c)
	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret
}

func (m *Machine) CurrentStatus() machine.Status {
	cStatus := C.machineCurrentStatus(m.c)
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

func (m *Machine) LastBlockReason() machine.BlockReason {
	cBlockReason := C.machineLastBlockReason(m.c)
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
		rawTimeoutBytes := C.GoBytes(unsafe.Pointer(cBlockReason.val1.data), cBlockReason.val1.length)
		timeout, err := value.UnmarshalValue(bytes.NewReader(rawTimeoutBytes[:]))
		if err != nil {
			panic(err)
		}
		timeoutInt, ok := timeout.(value.IntValue)
		if !ok {
			panic("Inbox hash must be an int")
		}
		C.free(cBlockReason.val1.data)
		return machine.InboxBlocked{Timeout: timeoutInt}
	default:
	}
	return nil
}

func (m *Machine) PrintState() {
	C.machinePrint(m.c)
}

func (m *Machine) ExecuteAssertion(
	maxSteps uint32,
	timeBounds *protocol.TimeBoundsBlocks,
	inbox value.TupleValue,
) (*protocol.ExecutionAssertion, uint32) {
	startTime := timeBounds.Start.AsInt()
	endTime := timeBounds.End.AsInt()

	var startTimeBuf bytes.Buffer
	err := value.NewIntValue(startTime).Marshal(&startTimeBuf)
	if err != nil {
		log.Fatal(err)
	}

	var endTimeBuf bytes.Buffer
	err = value.NewIntValue(endTime).Marshal(&endTimeBuf)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = value.MarshalValue(inbox, &buf)
	if err != nil {
		log.Fatal(err)
	}

	startTimeData := startTimeBuf.Bytes()
	endTimeData := endTimeBuf.Bytes()
	msgData := buf.Bytes()
	assertion := C.machineExecuteAssertion(
		m.c,
		C.uint64_t(maxSteps),
		unsafe.Pointer(&startTimeData[0]),
		unsafe.Pointer(&endTimeData[0]),
		unsafe.Pointer(&msgData[0]),
	)

	outMessagesRaw := C.GoBytes(unsafe.Pointer(assertion.outMessageData), assertion.outMessageLength)
	logsRaw := C.GoBytes(unsafe.Pointer(assertion.logData), assertion.logLength)
	outMessageVals := bytesArrayToVals(outMessagesRaw, int(assertion.outMessageCount))
	logVals := bytesArrayToVals(logsRaw, int(assertion.logCount))

	return protocol.NewExecutionAssertion(
		m.Hash(),
		int(assertion.didInboxInsn) != 0,
		uint64(assertion.numGas),
		outMessageVals,
		logVals,
	), uint32(assertion.numSteps)
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	rawProof := C.machineMarshallForProof(m.c)
	return C.GoBytes(unsafe.Pointer(rawProof.data), rawProof.length), nil
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	cCheckpointStorage := storage.(*CheckpointStorage)
	success := C.checkpointMachine(m.c, cCheckpointStorage.c)

	return success == 1
}

func (m *Machine) RestoreCheckpoint(storage machine.CheckpointStorage, machineHash [32]byte) bool {
	cCheckpointStorage, ok := storage.(*CheckpointStorage)

	if ok {
		success := C.restoreMachine(m.c, cCheckpointStorage.c, unsafe.Pointer(&machineHash[0]))

		return success == 1
	} else {
		return false
	}
}

func bytesArrayToVals(data []byte, valCount int) []value.Value {
	rd := bytes.NewReader(data)
	vals := make([]value.Value, 0, valCount)
	for i := 0; i < valCount; i++ {
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return vals
}
