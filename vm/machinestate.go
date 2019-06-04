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
	"errors"
	"github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/code"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm/stack"

	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
	"time"
)

var HashOfHaltedMachine [32]byte
var HashOfSizeExceptionMachine [32]byte

func init() {
	HashOfHaltedMachine = sha256.Sum256([]byte("This is the hash of a halted Arbitrum VM"))
	HashOfSizeExceptionMachine = sha256.Sum256([]byte("This is the hash of an Arbitrum VM with a size exception"))
}

type MachineContext interface {
	CanSpend(tokenType value.IntValue, currency value.IntValue) bool
	Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error
	ReadInbox() value.Value
	GetTimeBounds() value.Value
	NotifyStep()
	LoggedValue(value.Value) error

	OutMessageCount() int
}

type MachineNoContext struct{}

func (m *MachineNoContext) LoggedValue(data value.Value) error {
	return errors.New("can't log values outside of assertion mode")
}

func (m *MachineNoContext) CanSpend(tokenType value.IntValue, currency value.IntValue) bool {
	return false
}

func (m *MachineNoContext) Send(data value.Value, tokenType value.IntValue, currency value.IntValue, dest value.IntValue) error {
	return errors.New("can't send message outside of assertion mode")
}

func (m *MachineNoContext) OutMessageCount() int {
	return 0
}

func (m *MachineNoContext) ReadInbox() value.Value {
	return value.NewEmptyTuple()
}

func (m *MachineNoContext) GetTimeBounds() value.Value {
	return value.NewEmptyTuple()
}

func (m *MachineNoContext) NotifyStep() {

}

type MachineStatus int
const (
	MACHINE_EXTENSIVE MachineStatus = iota
	MACHINE_ERRORSTOP
	MACHINE_HALT
)

type Machine struct {
	// implements Machinestate
	stack      stack.Stack
	auxstack   stack.Stack
	register   *MachineValue
	static     *MachineValue
	pc         *MachinePC
	errHandler value.CodePointValue
	context    MachineContext
	status     MachineStatus

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
	//stack := NewTuple(value.NewEmptyTuple())
	//auxstack := NewTuple(value.NewEmptyTuple())
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
	ret := &Machine{datastack, auxstack, register, static, pc, errHandler, &MachineNoContext{}, MACHINE_EXTENSIVE,sizeLimit, false, wh}
	ret.checkSize()
	return ret
}

func RestoreMachine(opCodes []value.Operation, stackVal, auxStackVal, registerVal, staticVal, pcVal, errHandlerVal value.CodePointValue, sizeLimit int64) *Machine {
	datastack := stack.FlatFromTupleChain(stackVal)
	auxStack := stack.FlatFromTupleChain(auxStackVal)
	register := NewMachineValue(registerVal)
	static := NewMachineValue(staticVal)
	wh := NewSilentWarningHandler()
	pc := NewMachinePC(opCodes, wh)
	wh.SwitchMachinePC(pc)
	pc.SetPCForced(pcVal)
	return &Machine{datastack, auxStack, register, static, pc, errHandlerVal, &MachineNoContext{}, MACHINE_EXTENSIVE, sizeLimit, false, wh}
}

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

func (m *Machine) Context() MachineContext {
	return m.context
}

func (m *Machine) SetContext(mc MachineContext) {
	m.context = mc
}

func (m *Machine) IncrPC() {
	if !m.HaveSizeException() {
		err := m.pc.IncrPC()
		if err != nil {
			m.status = MACHINE_ERRORSTOP
		}
	}
}

func (m *Machine) GetPC() value.Value {
	return m.pc.GetPC()
}

func (m *Machine) GetErrHandler() value.CodePointValue {
     return m.errHandler;
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
	m.status = MACHINE_HALT
}

func (m *Machine) ErrorStop() {
	m.status = MACHINE_ERRORSTOP
}

func (m *Machine) IsHalted() bool {
	return m.status == MACHINE_HALT
}

func (m *Machine) IsErrored() bool {
	return m.status == MACHINE_ERRORSTOP
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

func (m *Machine) run() (bool, bool, string) {
	//fmt.Println("BEFORE", m.pc.GetPC().Op, m.stack.(*stack.Flat))
	if m.IsHalted() || m.IsErrored() || m.HaveSizeException() {
		return false, false, "Can't run"
	}
	insnName := m.pc.GetCurrentInsnName()
	_, err := RunInstruction(m, m.pc.GetCurrentInsn())
	if _, blocked := err.(VMBlockedError); blocked {
		return false, false, "Blocked"
	}
	m.context.NotifyStep()
	if err != nil {
		fmt.Printf("error running instruction %v: %v\n", insnName, err)
		return false, false, "Error"
	}
	if m.IsHalted() {
		return true, false, "Halted"
	}
	if m.IsErrored() {
		return true, false, "ErrorStopped"
	}
	if m.HaveSizeException() {
		return true, false, "SizeException"
	}
	//fmt.Println("AFTER", m.pc.GetPC().Op, m.stack.(*stack.Flat))
	return true, true, "Success"
}

func (m *Machine) RunUntilStop() {
	//start := time.Now()
	i := 0
	continueRun := true
	timings := make(map[code.Opcode][]time.Duration)
	var ran bool
	for continueRun {
		s1 := time.Now()
		ran, continueRun, _ = m.run()
		if ran {
			i++
			timings[m.pc.GetCurrentInsn().GetOp()] = append(timings[m.pc.GetCurrentInsn().GetOp()], time.Since(s1))
		}

	}
	//elapsed := time.Since(start)
	//fmt.Printf("Execution took %s\n", elapsed)
	//for key, vals := range timings {
	//	var total time.Duration
	//	for _, val := range vals {
	//		total += val
	//	}
	//	average := total / time.Duration(int64(len(vals)))
	//	if total > 1*time.Millisecond {
	//		fmt.Printf("%v: (count: %v, total: %v, average: %v)\n", code.InstructionNames[key], len(vals), total, average)
	//	}
	//}
}

// run up to maxSteps steps, stop earlier if halt or advise instruction, return number of insns actually run, and advise string or ""
func (m *Machine) Run(maxSteps int32) (int32) {
	i := int32(0)
	continueRun := true
	var ran bool
	for continueRun && i < maxSteps {
		ran, continueRun, _ = m.run()
		if ran {
			i++
		}
	}
	return i
}

func (m *Machine) Warn(str string) {
	m.warnHandler.Warn(str)
}

func (m *Machine) Log(val value.Value) {
	m.context.LoggedValue(val)
}

func (m *Machine) Hash() [32]byte {
	switch m.status {
	case MACHINE_EXTENSIVE:
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
	case MACHINE_ERRORSTOP:
		return value.NewInt64Value(1).ToBytes()
	case MACHINE_HALT:
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

func (m *Machine) MarshalForProof(wr io.Writer) error {
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

func (m *Machine) Clone() *Machine { // clone machine state--new machine wll NOT be in proving mode
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
		&MachineNoContext{},
		m.status,
		m.sizeLimit,
		m.sizeException,
		newWarnHandler,
	}
	//WARNING: risk of bug here, because of shallow copy of stack, callstack
	return ret
}
