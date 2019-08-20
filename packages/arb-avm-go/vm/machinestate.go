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

package vm

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/code"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm/stack"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	// implements Machinestate
	stack       stack.Stack
	auxstack    stack.Stack
	register    *MachineValue
	static      *MachineValue
	pc          *MachinePC
	errHandler  value.CodePointValue
	context     machine.Context
	status      machine.Status
	blockReason machine.BlockReason

	inbox   *protocol.Inbox
	balance *protocol.BalanceTracker

	sizeLimit     int64
	sizeException bool

	warnHandler WarningHandler
}

func Equal(x, y *Machine) (bool, string) {
	if ok, err := x.stack.Equal(y.stack); !ok {
		tmp := "stack error: "
		tmp += err
		return false, tmp
	}
	if ok, err := x.auxstack.Equal(y.auxstack); !ok {
		tmp := "auxstack error: "
		tmp += err
		return false, tmp
	}
	if ok, err := x.register.Equal(y.register); !ok {
		tmp := "register error: "
		tmp += err
		return false, tmp
	}
	if ok, err := x.static.Equal(y.static); !ok {
		tmp := "static error: "
		tmp += err
		return false, tmp
	}
	if ok := x.errHandler.Equal(y.errHandler); !ok {
		return false, "err handlers not equal"
	}
	return true, ""
}

func NewMachine(opCodes []value.Operation, staticVal value.Value, warn bool, sizeLimit int64) *Machine {
	datastack := stack.NewEmptyFlat()
	auxstack := stack.NewEmptyFlat()
	// stack := NewTuple(value.NewEmptyTuple())
	// auxstack := NewTuple(value.NewEmptyTuple())
	register := NewMachineValue(value.NewEmptyTuple())
	static := NewMachineValue(staticVal)
	errHandler := value.ErrorCodePoint
	inbox := protocol.NewEmptyInbox()
	balance := protocol.NewBalanceTracker()
	var wh WarningHandler
	if warn {
		wh = NewVerboseWarningHandler(nil)
	} else {
		wh = NewSilentWarningHandler()
	}
	pc := NewMachinePC(opCodes, wh)
	wh.SwitchMachinePC(pc)
	ret := &Machine{
		datastack,
		auxstack,
		register,
		static,
		pc,
		errHandler,
		&machine.NoContext{},
		machine.Extensive,
		nil,
		inbox,
		balance,
		sizeLimit,
		false,
		wh,
	}
	ret.checkSize()
	return ret
}

// func RestoreMachine(opCodes []value.Operation, stackVal, auxStackVal, registerVal, staticVal, pcVal value.Value, errHandlerVal value.CodePointValue, sizeLimit int64) *Machine {
//	datastack := stack.FlatFromTupleChain(stackVal)
//	auxStack := stack.FlatFromTupleChain(auxStackVal)
//	register := NewMachineValue(registerVal)
//	static := NewMachineValue(staticVal)
//	wh := NewSilentWarningHandler()
//	pc := NewMachinePC(opCodes, wh)
//	wh.SwitchMachinePC(pc)
//	pc.SetPCForced(pcVal)
//	return &Machine{datastack, auxStack, register, static, pc, errHandlerVal, &machine.NoContext{}, Extensive, sizeLimit, false, wh}
//}

func (m *Machine) Stack() stack.Stack {
	return m.stack
}

func (m *Machine) AuxStack() stack.Stack {
	return m.auxstack
}

func (m *Machine) Register() *MachineValue {
	return m.register
}

func (m *Machine) Static() *MachineValue {
	return m.static
}

func (m *Machine) SetContext(mc machine.Context) {
	m.context = mc
}

func (m *Machine) ReadInbox() value.Value {
	return m.inbox.Receive()
}

func (m *Machine) CanSpend(tokenType protocol.TokenType, currency *big.Int) bool {
	return m.balance.CanSpend(tokenType, currency)
}

func (m *Machine) GetTimeBounds() value.Value {
	return m.context.GetTimeBounds()
}

func (m *Machine) IncrPC() {
	if !m.HaveSizeException() {
		err := m.pc.IncrPC()
		if err != nil {
			m.status = machine.ErrorStop
		}
	}
}

func (m *Machine) GetPC() value.Value {
	return m.pc.GetPC()
}

func (m *Machine) GetErrHandler() value.CodePointValue {
	return m.errHandler
}

func (m *Machine) GetOperation() value.Operation {
	return m.pc.GetCurrentInsn()
}

func (m *Machine) GetAllOperations() []value.Operation {
	ret := make([]value.Operation, len(m.pc.flat)) // be cautious; copy the slice
	copy(ret, m.pc.flat)
	return ret
}

func (m *Machine) SetPC(iv value.Value) error {
	if !m.HaveSizeException() && !m.IsHalted() {
		return m.pc.SetPCForced(iv)
	}
	return nil
}

func (m *Machine) Halt() {
	m.status = machine.Halt
}

func (m *Machine) ErrorStop() {
	m.status = machine.ErrorStop
}

func (m *Machine) IsHalted() bool {
	return m.status == machine.Halt
}

func (m *Machine) IsErrored() bool {
	return m.status == machine.ErrorStop
}

func (m *Machine) HaveSizeException() bool {
	return m.sizeException
}

func (m *Machine) checkSize() {
	if !m.IsHalted() && !m.HaveSizeException() {
		if m.stack.Size()+m.register.Size()+m.static.Size() >= m.sizeLimit {
			m.sizeException = true
		}
	}
}

func (m *Machine) GetSizeLimit() int64 {
	return m.sizeLimit
}

func (m *Machine) CurrentStatus() machine.Status {
	return m.status
}

func (m *Machine) LastBlockReason() machine.BlockReason {
	return m.blockReason
}

// ExecuteAssertion runs the machine up to maxSteps steps, stoping earlier if halted, errored or blocked
func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	assCtx := NewMachineAssertionContext(
		m,
		timeBounds,
	)
	m.blockReason = nil
	for assCtx.StepCount() < uint32(maxSteps) {
		_, blocked := RunInstruction(m, m.pc.GetCurrentInsn())
		if blocked != nil {
			m.blockReason = blocked
			break
		}
	}
	return assCtx.Finalize(m)
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	m.inbox.SendMessage(msg)
	m.balance.Add(msg.TokenType, msg.Currency)
}

func (m *Machine) DeliverOnchainMessage() {
	m.inbox.DeliverMessages()
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	m.inbox.InsertMessageGroup(msgs)
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	return value.NewHashOnlyValueFromValue(m.inbox.Receive())
}

func (m *Machine) PendingMessageCount() uint64 {
	return m.inbox.PendingQueue.MessageCount()
}

func (m *Machine) Send(message protocol.Message) error {
	err := m.balance.Spend(message.TokenType, message.Currency)
	if err != nil {
		return err
	}
	m.context.Send(message)
	return nil
}

func (m *Machine) Warn(str string) {
	m.warnHandler.Warn(str)
}

func (m *Machine) Log(val value.Value) {
	m.context.LoggedValue(val)
}

func (m *Machine) Hash() [32]byte {
	switch m.status {
	case machine.Extensive:
		ret := [32]byte{}
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(m.pc.GetCurrentCodePointHash()),
			solsha3.Bytes32(m.stack.StateValue().Hash()),
			solsha3.Bytes32(m.auxstack.StateValue().Hash()),
			solsha3.Bytes32(m.register.StateValue().Hash()),
			solsha3.Bytes32(m.static.StateValue().Hash()),
			solsha3.Bytes32(m.errHandler.Hash()),
		))
		return ret
	case machine.ErrorStop:
		return value.NewInt64Value(1).ToBytes()
	case machine.Halt:
		return value.NewInt64Value(0).ToBytes()
	}
	panic("Machine::Hash: invalid machine status")
}

func (m *Machine) PrintState() {
	codePointHash := m.pc.GetPC().Hash()
	stackHash := m.stack.StateValue().Hash()
	auxStackHash := m.auxstack.StateValue().Hash()
	registerHash := m.register.StateValue().Hash()
	staticHash := m.static.StateValue().Hash()
	errHandlerHash := m.errHandler.Hash()
	fmt.Println("codePointHash", hexutil.Encode(codePointHash[:]))
	fmt.Println("stackHash", hexutil.Encode(stackHash[:]))
	fmt.Println("auxStackHash", hexutil.Encode(auxStackHash[:]))
	fmt.Println("registerHash", hexutil.Encode(registerHash[:]))
	fmt.Println("staticHash", hexutil.Encode(staticHash[:]))
	fmt.Println("errHandlerHash", hexutil.Encode(errHandlerHash[:]))
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := m.marshalForProof(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m *Machine) marshalForProof(wr io.Writer) error {
	codePoint := m.pc.GetPC()

	stackPops := code.InstructionStackPops[codePoint.Op.GetOp()]
	if _, ok := codePoint.Op.(value.ImmediateOperation); ok && len(stackPops) > 0 {
		stackPops = stackPops[1:]
	}
	auxStackPops := code.InstructionAuxStackPops[codePoint.Op.GetOp()]

	baseStackVal, stackVals := m.stack.SolidityProofValue(stackPops)
	baseStackValHash := baseStackVal.Hash()
	baseAuxStackVal, auxStackVals := m.auxstack.SolidityProofValue(auxStackPops)
	baseAuxStackValHash := baseAuxStackVal.Hash()
	registerHash := m.register.ProofValue().Hash()
	staticHash := m.static.ProofValue().Hash()
	errHandlerHash := m.errHandler.Hash()

	fmt.Printf("Proof of %v has %d stack vals and %d aux stack vals s\n", codePoint, len(stackVals), len(auxStackVals))

	if _, err := wr.Write(codePoint.NextHash[:]); err != nil {
		return err
	}
	if _, err := wr.Write(baseStackValHash[:]); err != nil {
		return err
	}
	if _, err := wr.Write(baseAuxStackValHash[:]); err != nil {
		return err
	}
	if _, err := wr.Write(registerHash[:]); err != nil {
		return err
	}
	if _, err := wr.Write(staticHash[:]); err != nil {
		return err
	}
	if _, err := wr.Write(errHandlerHash[:]); err != nil {
		return err
	}
	if err := value.MarshalOperationProof(codePoint.Op, wr, true); err != nil {
		return err
	}
	for _, val := range stackVals {
		if err := value.MarshalValue(val, wr); err != nil {
			return err
		}
	}
	for _, val := range auxStackVals {
		if err := value.MarshalValue(val, wr); err != nil {
			return err
		}
	}
	return nil
}

func (m *Machine) Clone() machine.Machine { // clone machine state--new machine wll NOT be in proving mode
	newWarnHandler := m.warnHandler.Clone()
	newPc := *m.pc
	newPc.warn = newWarnHandler
	newPcPointer := &newPc
	newWarnHandler.SwitchMachinePC(newPcPointer)
	ret := &Machine{
		m.stack.Clone(),
		m.auxstack.Clone(),
		m.register.Clone(),
		m.static.Clone(),
		newPcPointer,
		m.errHandler,
		&machine.NoContext{},
		m.status,
		m.blockReason,
		m.inbox.Clone(),
		m.balance.Clone(),
		m.sizeLimit,
		m.sizeException,
		newWarnHandler,
	}
	// WARNING: risk of bug here, because of shallow copy of stack, callstack
	return ret
}
