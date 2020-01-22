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

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/code"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm/stack"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Machine struct {
	// implements Machinestate
	stack      stack.Stack
	auxstack   stack.Stack
	register   *MachineValue
	static     *MachineValue
	pc         *MachinePC
	errHandler value.CodePointValue
	context    Context
	status     machine.Status

	sizeLimit     int64
	sizeException bool

	warnHandler WarningHandler
}

func (m *Machine) Checkpoint(storage machine.CheckpointStorage) bool {
	panic("implement me")
}

func (m *Machine) RestoreCheckpoint(storage machine.CheckpointStorage, machineHash common.Hash) bool {
	panic("implement me")
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
		&NoContext{},
		machine.Extensive,
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

func (m *Machine) SetContext(mc Context) {
	m.context = mc
}

func (m *Machine) GetInbox() value.TupleValue {
	return m.context.GetInbox()
}

func (m *Machine) GetStartTime() value.IntValue {
	return m.context.GetStartTime()
}

func (m *Machine) GetEndTime() value.IntValue {
	return m.context.GetEndTime()
}

func (m *Machine) IncrPC() {
	if !m.HaveSizeException() {
		err := m.pc.IncrPC()
		if err != nil {
			m.status = machine.ErrorStop
		}
	}
}

func (m *Machine) GetPC() value.CodePointValue {
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

func (m *Machine) IsBlocked(currentTime *common.TimeBlocks, newMessages bool) machine.BlockReason {
	if m.status == machine.ErrorStop {
		return machine.ErrorBlocked{}
	}
	if m.status == machine.Halt {
		return machine.HaltBlocked{}
	}
	op := m.GetOperation()
	if op.GetOp() == code.INBOX {
		if newMessages {
			return nil
		}
		var param value.Value
		if immediate, ok := op.(value.ImmediateOperation); ok {
			param = immediate.Val
		} else {
			param, _ = m.stack.Pop()
			m.stack.Push(param)
		}
		paramInt, ok := param.(value.IntValue)
		if !ok {
			return nil
		}
		if currentTime.AsInt().Cmp(paramInt.BigInt()) < 0 {
			return machine.InboxBlocked{Timeout: paramInt}
		}
		return nil
	}
	return nil
}

// ExecuteAssertion runs the machine up to maxSteps steps, stoping earlier if halted, errored or blocked
func (m *Machine) ExecuteAssertion(maxSteps uint32, timeBounds *protocol.TimeBoundsBlocks, inbox value.TupleValue) (*protocol.ExecutionAssertion, uint32) {
	assCtx := NewMachineAssertionContext(
		m,
		timeBounds,
		inbox,
	)
	for assCtx.StepCount() < maxSteps {
		_, blocked := RunInstruction(m, m.pc.GetCurrentInsn())
		if blocked != nil {
			break
		}
	}
	return assCtx.Finalize(m)
}

func (m *Machine) Send(message value.Value) {
	m.context.Send(message)
}

func (m *Machine) Warn(str string) {
	m.warnHandler.Warn(str)
}

func (m *Machine) Log(val value.Value) {
	m.context.LoggedValue(val)
}

func (m *Machine) Hash() common.Hash {
	switch m.status {
	case machine.Extensive:
		return hashing.SoliditySHA3(
			hashing.Bytes32(m.pc.GetCurrentCodePointHash()),
			hashing.Bytes32(m.stack.StateValue().Hash()),
			hashing.Bytes32(m.auxstack.StateValue().Hash()),
			hashing.Bytes32(m.register.StateValue().Hash()),
			hashing.Bytes32(m.static.StateValue().Hash()),
			hashing.Bytes32(m.errHandler.Hash()),
		)
	case machine.ErrorStop:
		return value.NewInt64Value(1).ToBytes()
	case machine.Halt:
		return value.NewInt64Value(0).ToBytes()
	}
	panic("Machine::Hash: invalid machine status")
}

func (m *Machine) PrintState() {
	codePointHash := m.pc.GetCurrentCodePointHash()
	stackHash := m.stack.StateValue().Hash()
	auxStackHash := m.auxstack.StateValue().Hash()
	registerHash := m.register.StateValue().Hash()
	staticHash := m.static.StateValue().Hash()
	errHandlerHash := m.errHandler.Hash()
	fmt.Println("machine state", m.status)
	fmt.Println("codePointHash", codePointHash)
	fmt.Println("stackHash", stackHash[:])
	fmt.Println("auxStackHash", auxStackHash)
	fmt.Println("registerHash", registerHash)
	fmt.Println("staticHash", staticHash)
	fmt.Println("errHandlerHash", errHandlerHash)
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
	includeImmediateVal := false
	if _, ok := codePoint.Op.(value.ImmediateOperation); ok && len(stackPops) > 0 {
		if stackPops[0] == 1 {
			includeImmediateVal = true
		}
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
	if err := value.MarshalOperationProof(codePoint.Op, wr, includeImmediateVal); err != nil {
		return err
	}
	for _, val := range stackVals {
		if err := value.MarshalValueForProof(val, wr); err != nil {
			return err
		}
	}
	for _, val := range auxStackVals {
		if err := value.MarshalValueForProof(val, wr); err != nil {
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
		&NoContext{},
		m.status,
		m.sizeLimit,
		m.sizeException,
		newWarnHandler,
	}
	// WARNING: risk of bug here, because of shallow copy of stack, callstack
	return ret
}
