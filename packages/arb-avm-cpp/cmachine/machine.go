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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lsecp256k1 -lff
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
	"time"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

func (m *Machine) Hash() (ret common.Hash) {
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

func (m *Machine) IsBlocked(currentTime *common.TimeBlocks, newMessages bool) machine.BlockReason {
	var currentTimeBuf bytes.Buffer
	err := value.NewIntValue(currentTime.AsInt()).Marshal(&currentTimeBuf)
	if err != nil {
		log.Fatal(err)
	}
	newMessagesInt := 0
	if newMessages {
		newMessagesInt = 1
	}
	currentTimeDataC := C.CBytes(currentTimeBuf.Bytes())
	cBlockReason := C.machineIsBlocked(m.c, currentTimeDataC, C.int(newMessagesInt))
	C.free(currentTimeDataC)
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
		rawTimeoutBytes := C.GoBytes(unsafe.Pointer(cBlockReason.val.data), cBlockReason.val.length)
		timeout, err := value.UnmarshalValue(bytes.NewReader(rawTimeoutBytes[:]))
		if err != nil {
			panic(err)
		}
		timeoutInt, ok := timeout.(value.IntValue)
		if !ok {
			panic("Inbox hash must be an int")
		}
		C.free(cBlockReason.val.data)
		return machine.InboxBlocked{Timeout: timeoutInt}
	default:
	}
	return nil
}

func (m *Machine) PrintState() {
	C.machinePrint(m.c)
}

func (m *Machine) ExecuteAssertion(
	maxSteps uint64,
	timeBounds *protocol.TimeBounds,
	inbox value.TupleValue,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, uint64) {
	lowerBoundBlock := timeBounds.LowerBoundBlock.AsInt()
	upperBoundBlock := timeBounds.UpperBoundBlock.AsInt()
	lowerBoundTimestamp := timeBounds.LowerBoundTimestamp
	upperBoundTimestamp := timeBounds.UpperBoundTimestamp

	var lowerBoundBlockBuf bytes.Buffer
	err := value.NewIntValue(lowerBoundBlock).Marshal(&lowerBoundBlockBuf)
	if err != nil {
		log.Fatal(err)
	}

	var upperBoundBlockBuf bytes.Buffer
	err = value.NewIntValue(upperBoundBlock).Marshal(&upperBoundBlockBuf)
	if err != nil {
		log.Fatal(err)
	}

	var lowerBoundTimestampBuf bytes.Buffer
	err = value.NewIntValue(lowerBoundTimestamp).Marshal(&lowerBoundTimestampBuf)
	if err != nil {
		log.Fatal(err)
	}

	var upperBoundTimestampBuf bytes.Buffer
	err = value.NewIntValue(upperBoundTimestamp).Marshal(&upperBoundTimestampBuf)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = value.MarshalValue(inbox, &buf)
	if err != nil {
		log.Fatal(err)
	}

	lowerBoundBlockData := lowerBoundBlockBuf.Bytes()
	upperBoundBlockData := upperBoundBlockBuf.Bytes()
	lowerBoundTimestampData := lowerBoundTimestampBuf.Bytes()
	upperBoundTimestampData := upperBoundTimestampBuf.Bytes()
	msgData := buf.Bytes()
	lowerBoundBlockDataC := C.CBytes(lowerBoundBlockData)
	upperBoundBlockDataC := C.CBytes(upperBoundBlockData)
	lowerBoundTimestampDataC := C.CBytes(lowerBoundTimestampData)
	upperBoundTimestampDataC := C.CBytes(upperBoundTimestampData)
	msgDataC := C.CBytes(msgData)
	assertion := C.machineExecuteAssertion(
		m.c,
		C.uint64_t(maxSteps),
		lowerBoundBlockDataC,
		upperBoundBlockDataC,
		lowerBoundTimestampDataC,
		upperBoundTimestampDataC,
		msgDataC,
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)
	C.free(lowerBoundBlockDataC)
	C.free(upperBoundBlockDataC)
	C.free(lowerBoundTimestampDataC)
	C.free(upperBoundTimestampDataC)
	C.free(msgDataC)

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
	), uint64(assertion.numSteps)
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
