/*
 * Copyright 2019, Offchain Labs, Inc.
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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -lstdc++ -lm -lrocksdb
#include "../cavm/cmachine.h"
#include "../cavm/ccheckpointstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"fmt"
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

func (m *Machine) InboxHash() value.HashOnlyValue {
	var hash [32]byte
	C.machineInboxHash(m.c, unsafe.Pointer(&hash[0]))
	return value.NewHashOnlyValue(hash, 0)
}

func (m *Machine) PendingMessageCount() uint64 {
	return uint64(C.machinePendingMessageCount(m.c))
}

func (m *Machine) PrintState() {
	C.machinePrint(m.c)
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	var buf bytes.Buffer
	err := value.MarshalValue(msg.AsValue(), &buf)
	if err != nil {
		panic(err)
	}
	msgData := buf.Bytes()
	C.machineSendOnchainMessage(m.c, unsafe.Pointer(&msgData[0]))
}

func (m *Machine) DeliverOnchainMessage() {
	C.machineDeliverOnchainMessages(m.c)
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	var buf bytes.Buffer
	for _, msg := range msgs {
		err := value.MarshalValue(msg.AsValue(), &buf)
		if err != nil {
			panic(err)
		}
	}
	msgsData := buf.Bytes()
	if len(msgsData) > 0 {
		C.machineSendOffchainMessages(m.c, unsafe.Pointer(&msgsData[0]), C.int(len(msgs)))
	}
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds *protocol.TimeBounds) *protocol.Assertion {
	assertion := C.machineExecuteAssertion(
		m.c,
		C.uint64_t(maxSteps),
		C.uint64_t(timeBounds.StartTime),
		C.uint64_t(timeBounds.EndTime),
	)

	outMessagesRaw := C.GoBytes(unsafe.Pointer(assertion.outMessageData), assertion.outMessageLength)
	logsRaw := C.GoBytes(unsafe.Pointer(assertion.logData), assertion.logLength)
	outMessageVals := bytesArrayToVals(outMessagesRaw, int(assertion.outMessageCount))
	logVals := bytesArrayToVals(logsRaw, int(assertion.logCount))

	return protocol.NewAssertion(
		m.Hash(),
		uint32(assertion.numSteps),
		uint64(assertion.numGas),
		outMessageVals,
		logVals,
	)
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

func (m *Machine) RestoreCheckpoint(storage machine.CheckpointStorage, checkpointName string) bool {
	cCheckpointName := C.CString(checkpointName)
	cCheckpointStorage, ok := storage.(*CheckpointStorage)

	if ok {
		success := C.restoreMachine(m.c, cCheckpointStorage.c, cCheckpointName)
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
