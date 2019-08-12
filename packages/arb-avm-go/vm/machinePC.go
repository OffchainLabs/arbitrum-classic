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
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/code"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

const CodeSaveFrequency = 2

var HashOfLastInstruction [32]byte

type MachinePC struct {
	// implements Machinestate
	warn        WarningHandler
	flat        []value.Operation
	savedValues []value.CodePointValue
	pc          int64 // -1 if machine has halted, otherwise index into code
}

func NewMachinePC(insns []value.Operation, handler WarningHandler) *MachinePC {
	flat := make([]value.Operation, len(insns))
	copy(flat, insns)
	savedValues := make([]value.CodePointValue, len(insns)/CodeSaveFrequency+1)

	prevHash := HashOfLastInstruction
	var codePoint value.CodePointValue
	for i := int64(len(insns) - 1); i >= 0; i-- {
		codePoint = value.CodePointValue{InsnNum: i, Op: insns[i], NextHash: prevHash}
		prevHash = codePoint.Hash()
		if i%CodeSaveFrequency == 0 {
			savedValues[i/CodeSaveFrequency] = codePoint
		}
	}
	return &MachinePC{handler, flat, savedValues, 0}
}

func (m *MachinePC) Equal(y *MachinePC) (bool, string) {
	for i := range m.flat {
		if m.flat[i].GetOp() != y.flat[i].GetOp() {
			return false, "MachinePC flat Op different"
		}
		if m.flat[i].TypeCode() != y.flat[i].TypeCode() {
			return false, "MachinePC flat TypeCode different"
		}
	}
	for i := range m.savedValues {
		if !(value.Eq(m.savedValues[i], y.savedValues[i])) {
			return false, "Flat stack savedValues different"
		}
	}
	if m.pc != y.pc {
		return false, "Flat stack PC different"
	}

	return true, ""
}

func (m MachinePC) GetCurrentInsn() value.Operation {
	return m.flat[m.pc]
}

func (m MachinePC) GetCurrentInsnName() string {
	if m.pc < 0 {
		panic("Bad pc")
	}
	return code.InstructionNames[m.flat[m.pc].GetOp()]
}

func (m MachinePC) GetPC() value.CodePointValue {
	if m.pc >= int64(len(m.flat)) || m.pc < 0 {
		panic(fmt.Sprintf("Invalid pc: %v", m.pc))
	}
	if m.pc == int64(len(m.flat))-1 {
		return value.CodePointValue{InsnNum: m.pc, Op: m.flat[m.pc], NextHash: HashOfLastInstruction}
	}

	lookupPoint := (m.pc + CodeSaveFrequency - 1) / CodeSaveFrequency
	codePoint := m.savedValues[lookupPoint]
	curPC := lookupPoint * CodeSaveFrequency

	if curPC == m.pc {
		return codePoint
	}
	for i := curPC - 1; i >= m.pc; i-- {
		codePoint = value.CodePointValue{InsnNum: i, Op: m.flat[i], NextHash: codePoint.Hash()}
	}
	return codePoint
}

func (m MachinePC) GetCurrentCodePointHash() [32]byte {
	if m.pc == -1 {
		return HashOfLastInstruction
	}
	return m.GetPC().Hash()
}

func (m *MachinePC) IncrPC() error {
	m.pc = 1 + m.pc
	if m.pc >= int64(len(m.flat)) {
		return errors.New("IncrPC: PC reached end and halted")
	}
	return nil
}

func (m *MachinePC) SetPCForced(iv value.Value) error {
	codePointVal, ok := iv.(value.CodePointValue)
	if !ok {
		return errors.New("SetPC: tried to set PC to unknown value. Cannot set PC")
	}
	iv64 := codePointVal.InsnNum
	if !(iv64 >= -2 && iv64 < int64(len(m.flat))) {
		m.warn.Warn("SetPC: set PC to invalid value")
	}
	m.pc = iv64
	return nil
}
