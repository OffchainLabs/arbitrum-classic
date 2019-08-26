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
#cgo LDFLAGS: -L. -lcavm -lavm -lstdc++
#include "../cavm/cmachine.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"math/big"
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
		rawInboxHash := C.GoBytes(unsafe.Pointer(cBlockReason.val1.data), cBlockReason.val1.length)
		inboxHash, err := value.UnmarshalValue(bytes.NewReader(rawInboxHash[:]))
		if err != nil {
			panic(err)
		}
		inboxHashInt, ok := inboxHash.(value.IntValue)
		if !ok {
			panic("Inbox hash must be an int")
		}
		C.free(cBlockReason.val1.data)
		return machine.InboxBlocked{Inbox: value.NewHashOnlyValue(inboxHashInt.ToBytes(), 0)}
	case C.BLOCK_TYPE_SEND:
		rawCurrency := C.GoBytes(unsafe.Pointer(cBlockReason.val1.data), cBlockReason.val1.length)
		currency, err := value.UnmarshalValue(bytes.NewReader(rawCurrency[:]))
		if err != nil {
			panic(err)
		}
		currencyInt, ok := currency.(value.IntValue)
		if !ok {
			panic("Inbox hash must be an int")
		}
		C.free(cBlockReason.val1.data)

		rawTokenType := C.GoBytes(unsafe.Pointer(cBlockReason.val2.data), cBlockReason.val2.length)
		var tokType protocol.TokenType
		copy(tokType[:], rawTokenType)
		C.free(cBlockReason.val2.data)
		return machine.SendBlocked{
			Currency:  currencyInt.BigInt(),
			TokenType: tokType,
		}
	default:
	}
	return nil
}

func (m *Machine) CanSpend(tokenType protocol.TokenType, currency *big.Int) bool {
	var currencyBuf bytes.Buffer
	_ = value.NewIntValue(currency).Marshal(&currencyBuf)
	currencyData := currencyBuf.Bytes()
	canSpend := C.machineCanSpend(m.c, (*C.char)(unsafe.Pointer(&tokenType[0])), (*C.char)(unsafe.Pointer(&currencyData[0])))
	return int(canSpend) != 0
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	var hash [32]byte
	C.machineInboxHash(m.c, unsafe.Pointer(&hash[0]))
	return value.NewHashOnlyValue(hash, 0)
}

func (m *Machine) PendingMessageCount() uint64 {
	return uint64(C.machinePendingMessageCount(m.c))
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

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	assertion := C.machineExecuteAssertion(
		m.c,
		C.uint64_t(maxSteps),
		C.uint64_t(timeBounds[0]),
		C.uint64_t(timeBounds[1]),
	)

	outMessagesRaw := C.GoBytes(unsafe.Pointer(assertion.outMessageData), assertion.outMessageLength)
	logsRaw := C.GoBytes(unsafe.Pointer(assertion.logData), assertion.logLength)

	outMessageVals := bytesArrayToVals(outMessagesRaw, int(assertion.outMessageCount))
	outMessages := make([]protocol.Message, 0, len(outMessageVals))
	for _, msgVal := range outMessageVals {
		msg, err := protocol.NewMessageFromValue(msgVal)
		if err != nil {
			panic(err)
		}
		outMessages = append(outMessages, msg)
	}

	logVals := bytesArrayToVals(logsRaw, int(assertion.logCount))

	return protocol.NewAssertion(
		m.Hash(),
		uint32(assertion.numSteps),
		outMessages,
		logVals,
	)
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	rawProof := C.machineMarshallForProof(m.c)
	return C.GoBytes(unsafe.Pointer(rawProof.data), rawProof.length), nil
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
