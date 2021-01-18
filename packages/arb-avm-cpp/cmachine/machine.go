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
#include "../cavm/cmachine.h"
#include "../cavm/ccheckpointstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"runtime"
	"time"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var logger = log.With().Caller().Str("component", "cmachine").Logger()

type Machine struct {
	c unsafe.Pointer
}

func New(codeFile string) (*Machine, error) {
	cFilename := C.CString(codeFile)
	cMachine := C.machineCreate(cFilename)
	if cMachine == nil {
		return nil, errors.Errorf("error creating machine from file %s", codeFile)
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

func (m *Machine) IsBlocked(newMessages bool) machine.BlockReason {
	newMessagesInt := 0
	if newMessages {
		newMessagesInt = 1
	}
	cBlockReason := C.machineIsBlocked(m.c, C.int(newMessagesInt))
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

func (m *Machine) PrintState() {
	C.machinePrint(m.c)
}

func makeExecutionAssertion(
	assertion C.RawAssertion,
	beforeMachineHash common.Hash,
	afterMachineHash common.Hash,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	outMessagesRaw := toByteSlice(assertion.outMessages)
	logsRaw := toByteSlice(assertion.logs)
	debugPrints := protocol.BytesArrayToVals(toByteSlice(assertion.debugPrints), uint64(assertion.debugPrintCount))
	return protocol.NewExecutionAssertion(
		beforeMachineHash,
		afterMachineHash,
		uint64(assertion.numGas),
		uint64(assertion.inbox_messages_consumed),
		outMessagesRaw,
		uint64(assertion.outMessageCount),
		logsRaw,
		uint64(assertion.logCount),
	), debugPrints, uint64(assertion.numSteps)
}

func encodeInboxMessages(inboxMessages []inbox.InboxMessage) []byte {
	var buf bytes.Buffer
	for _, msg := range inboxMessages {
		// Error just occurs on write, and bytes.Buffer is safe
		_ = value.MarshalValue(msg.AsValue(), &buf)
	}
	return buf.Bytes()
}

func encodeValue(val value.Value) []byte {
	var buf bytes.Buffer

	// Error just occurs on write, and bytes.Buffer is safe
	_ = value.MarshalValue(val, &buf)
	return buf.Bytes()
}

func (m *Machine) ExecuteAssertion(
	maxSteps uint64,
	messages []inbox.InboxMessage,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	msgDataC := C.CBytes(encodeInboxMessages(messages))
	defer C.free(msgDataC)

	beforeHash := m.Hash()
	assertion := C.executeAssertion(
		m.c,
		C.uint64_t(maxSteps),
		msgDataC,
		C.uint64_t(len(messages)),
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)

	return makeExecutionAssertion(assertion, beforeHash, m.Hash())
}

func (m *Machine) ExecuteCallServerAssertion(
	maxSteps uint64,
	inboxMessages []inbox.InboxMessage,
	fakeInboxPeekValue value.Value,
	maxWallTime time.Duration,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	msgDataC := C.CBytes(encodeInboxMessages(inboxMessages))
	defer C.free(msgDataC)

	inboxPeekDataC := C.CBytes(encodeValue(fakeInboxPeekValue))
	defer C.free(inboxPeekDataC)

	beforeHash := m.Hash()
	assertion := C.executeCallServerAssertion(
		m.c,
		C.uint64_t(maxSteps),
		msgDataC,
		C.uint64_t(len(inboxMessages)),
		inboxPeekDataC,
		C.uint64_t(uint64(maxWallTime.Seconds())),
	)

	return makeExecutionAssertion(assertion, beforeHash, m.Hash())
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	rawProof := C.machineMarshallForProof(m.c)
	return C.GoBytes(unsafe.Pointer(rawProof.data), rawProof.length), nil
}

func (m *Machine) MarshalBufferProof() ([]byte, error) {
	rawProof := C.machineMarshallBufferProof(m.c)
	return C.GoBytes(unsafe.Pointer(rawProof.data), rawProof.length), nil
}

func (m *Machine) MarshalState() ([]byte, error) {
	stateData := C.machineMarshallState(m.c)
	return C.GoBytes(unsafe.Pointer(stateData.data), stateData.length), nil
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	cCheckpointStorage := storage.(*CheckpointStorage)
	success := C.checkpointMachine(m.c, cCheckpointStorage.c)

	return success == 1
}
