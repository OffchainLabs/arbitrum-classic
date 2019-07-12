package cppavm

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lavm -lstdc++
#include <cmachine.h>
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"github.com/offchainlabs/arb-util/machine"
	"github.com/offchainlabs/arb-util/protocol"
	"github.com/offchainlabs/arb-util/value"
	"runtime"
	"unsafe"
	//"github.com/offchainlabs/arb-avm/loader"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/offchainlabs/arb-util/evm"
	//"github.com/offchainlabs/arb-avm/loader"
	//"github.com/offchainlabs/arb-util/protocol"
	//"github.com/offchainlabs/arb-util/value"
	//"log"
	//"math/big"
	//"os"
)

type Machine struct {
	c unsafe.Pointer
}

func New(codeFile string) *Machine {
	fmt.Println("CCreartVM codeFile=", codeFile)
	//****************
	// C stuff
	cFilename := C.CString(codeFile)

	cMachine := C.machineCreate(cFilename)
	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	C.free(unsafe.Pointer(cFilename))
	return ret
}

func cdestroyVM(cMachine *Machine) {
	fmt.Println("Calling C.machine_destroy")
	C.machineDestroy(cMachine.c)
}

func (m *Machine) Hash() (ret [32]byte) {
	C.machineHash(m.c, unsafe.Pointer(&ret[0]))
	return
}


func (m *Machine) Clone() machine.Machine {
	return &Machine{C.machineClone(m.c)}
}

func (m *Machine) InboxHash() (value.HashOnlyValue) {
	var hash [32]byte
	C.machineInboxHash(m.c, unsafe.Pointer(&hash[0]))
	return value.NewHashOnlyValue(hash, 0)
}

func (m *Machine) HasPendingMessages() bool {
	return C.machineHasPendingMessages(m.c) != 0
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

func (m *Machine) SendOffchainMessages(msgs[]protocol.Message) {
	var buf bytes.Buffer
	for _, msg := range msgs {
		err := value.MarshalValue(msg.AsValue(), &buf)
		if err != nil {
			panic(err)
		}
	}
	msgsData := buf.Bytes()
	C.machineSendOffchainMessages(m.c, unsafe.Pointer(&msgsData[0]), C.int(len(msgsData)))
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	assertion := C.machineExecuteAssertion(
		m.c,
		C.ulonglong(maxSteps),
		C.ulonglong(timeBounds[0]),
		C.ulonglong(timeBounds[1]),
	)

	fmt.Println("Finished raw assertion", assertion.outMessageLength, assertion.logLength)

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
	vals := []value.Value{}
	for i := 0; i < valCount; i++ {
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return vals
}