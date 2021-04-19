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
#cgo LDFLAGS: -L. -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lsecp256k1 -lff -lgmp -lkeccak -ldl
#cgo linux LDFLAGS: -latomic
#include "../cavm/cmachine.h"
#include "../cavm/carbstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var logger = log.With().Caller().Stack().Str("component", "cmachine").Logger()

type Machine struct {
	c unsafe.Pointer
}

func New(codeFile string) (*Machine, error) {
	cFilename := C.CString(codeFile)
	defer C.free(unsafe.Pointer(cFilename))
	cMachine := C.machineCreate(cFilename)
	if cMachine == nil {

		return nil, errors.Errorf("error creating machine from file %s", codeFile)

	}

	return WrapCMachine(cMachine), nil
}

func cdestroyVM(cMachine *Machine) {

	C.machineDestroy(cMachine.c)

}

func WrapCMachine(cMachine unsafe.Pointer) *Machine {
	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret
}

func (m *Machine) Hash() (ret common.Hash, err error) {
	success := C.machineHash(m.c, unsafe.Pointer(&ret[0]))
	if success == 0 {
		err = errors.New("Cannot get machine hash")
	}

	return
}

func (m *Machine) CodePointHash() (ret common.Hash) {
	C.machineCodePointHash(m.c, unsafe.Pointer(&ret[0]))
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

func (m *Machine) String() string {
	cStr := C.machineInfo(m.c)
	defer C.free(unsafe.Pointer(cStr))
	return C.GoString(cStr)
}

func makeExecutionAssertion(assertion C.RawAssertion) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	sendsRaw := receiveByteSlice(assertion.sends)
	sendAcc := receive32Bytes(assertion.sendAcc)
	logsRaw := receiveByteSlice(assertion.logs)
	logAcc := receive32Bytes(assertion.logAcc)
	debugPrints := protocol.BytesArrayToVals(receiveByteSlice(assertion.debugPrints), uint64(assertion.debugPrintCount))
	return protocol.NewExecutionAssertion(
		uint64(assertion.numGas),
		uint64(assertion.inbox_messages_consumed),
		sendsRaw,
		uint64(assertion.sendCount),
		sendAcc,
		logsRaw,
		uint64(assertion.logCount),
		logAcc,
	), debugPrints, uint64(assertion.numSteps)
}

func (m *Machine) ExecuteAssertion(
	maxGas uint64,
	goOverGas bool,
	messages []inbox.InboxMessage,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	return m.ExecuteAssertionAdvanced(
		maxGas,
		goOverGas,
		messages,
		nil,
		false,
		common.Hash{},
		common.Hash{},
	)
}

func bytesArrayToByteSliceArray(bytes [][]byte) C.struct_ByteSliceArrayStruct {
	byteSlices := encodeByteSliceList(bytes)
	sliceArrayData := C.malloc(C.size_t(C.sizeof_struct_ByteSliceStruct * len(byteSlices)))
	sliceArray := (*[1 << 30]C.struct_ByteSliceStruct)(sliceArrayData)[:len(byteSlices):len(byteSlices)]
	for i, data := range byteSlices {
		sliceArray[i] = data
	}
	return C.struct_ByteSliceArrayStruct{slices: sliceArrayData, count: C.int(len(byteSlices))}
}

func inboxMessagesToByteSliceArray(messages []inbox.InboxMessage) C.struct_ByteSliceArrayStruct {
	return bytesArrayToByteSliceArray(encodeInboxMessages(messages))
}

func (m *Machine) ExecuteAssertionAdvanced(
	maxGas uint64,
	goOverGas bool,
	messages []inbox.InboxMessage,
	sideloads []inbox.InboxMessage,
	stopOnSideload bool,
	beforeSendAcc common.Hash,
	beforeLogAcc common.Hash,
) (*protocol.ExecutionAssertion, []value.Value, uint64) {
	conf := C.machineExecutionConfigCreate()

	goOverGasInt := C.int(0)
	if goOverGas {
		goOverGasInt = 1
	}
	C.machineExecutionConfigSetMaxGas(conf, C.uint64_t(maxGas), goOverGasInt)

	msgData := bytesArrayToByteSliceArray(encodeMachineInboxMessages(messages))
	defer C.free(msgData.slices)
	C.machineExecutionConfigSetInboxMessages(conf, msgData)

	C.machineExecutionConfigSetInboxMessages(conf, msgData)

	sideloadsData := bytesArrayToByteSliceArray(encodeInboxMessages(sideloads))
	defer C.free(sideloadsData.slices)
	C.machineExecutionConfigSetSideloads(conf, sideloadsData)

	stopOnSideloadInt := C.int(0)
	if stopOnSideload {
		stopOnSideloadInt = 1
	}
	C.machineExecutionConfigSetStopOnSideload(conf, stopOnSideloadInt)

	assertion := C.executeAssertion(
		m.c,
		conf,
		unsafeDataPointer(beforeSendAcc.Bytes()),
		unsafeDataPointer(beforeLogAcc.Bytes()),
	)

	return makeExecutionAssertion(assertion)
}

func (m *Machine) MarshalForProof() ([]byte, []byte, error) {
	rawProof := C.machineMarshallForProof(m.c)
	return receiveByteSlice(rawProof.standard_proof), receiveByteSlice(rawProof.buffer_proof), nil
}

func (m *Machine) MarshalState() ([]byte, error) {
	stateData := C.machineMarshallState(m.c)
	return receiveByteSlice(stateData), nil
}
