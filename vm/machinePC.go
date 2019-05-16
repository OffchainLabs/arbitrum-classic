package vm

import (
	"crypto/sha256"
	"fmt"
	"github.com/offchainlabs/arb-avm/code"
	"github.com/offchainlabs/arb-avm/value"

	"errors"
)

const CodeSaveFrequency = 2

var HashOfLastInstruction [32]byte

func init() {
	HashOfLastInstruction = sha256.Sum256([]byte("This is the hash of the last instruction"))
}

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

func (s *MachinePC) Equal(y *MachinePC) (bool, string) {
	for i := range s.flat {
		if s.flat[i].GetOp() != y.flat[i].GetOp() {
			return false, "MachinePC flat Op different"
		}
		if s.flat[i].TypeCode() != y.flat[i].TypeCode() {
			return false, "MachinePC flat TypeCode different"
		}
	}
	for i := range s.savedValues {
		if !(value.Eq(s.savedValues[i], y.savedValues[i])) {
			return false, "Flat stack savedValues different"
		}
	}
	if s.pc != y.pc {
		return false, "Flat stack PC different"
	}

	return true, ""
}

func (m MachinePC) GetCurrentInsn() value.Operation {
	return m.flat[m.pc]
}

func (m MachinePC) GetCurrentInsnName() string {
	if m.pc >= 0 {
		return code.InstructionNames[m.flat[m.pc].GetOp()]
	} else if m.pc == -1 {
		return "HaltState"
	} else if m.pc == -2 {
		return "ErrorState"
	} else {
		panic("Bad pc")
	}
}

func (m MachinePC) GetPC() value.CodePointValue {
	//fmt.Println("Getting PC", m.pc)
	if m.pc == -1 {
		return value.HaltCodePoint
	} else if m.pc == -2 {
		return value.ErrorCodePoint
	}
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
	} else {
		for i := curPC - 1; i >= m.pc; i-- {
			codePoint = value.CodePointValue{InsnNum: i, Op: m.flat[i], NextHash: codePoint.Hash()}
		}
		return codePoint
	}
}

func (m MachinePC) GetCurrentCodePointHash() [32]byte {
	if m.pc == -1 {
		return HashOfLastInstruction
	} else {
		return m.GetPC().Hash()
	}
}

func (m *MachinePC) IncrPC() {
	if !m.IsHalted() {
		m.pc = 1 + m.pc
		if m.pc >= int64(len(m.flat)) {
			m.warn.Warn("IncrPC: PC reached end and halted")
			m.Halt()
		}
	}
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

func (m *MachinePC) Halt() {
	m.pc = -1
}

func (m MachinePC) IsHalted() bool {
	return m.pc == -1
}

func (m MachinePC) IsErrored() bool {
	return m.pc == -2
}
