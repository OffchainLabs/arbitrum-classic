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

package value

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type Opcode uint8

type Operation interface {
	GetOp() Opcode
	TypeCode() uint8
	Marshal(wr io.Writer) error
	MarshalProof(wr io.Writer, includeVal bool) error
}

type BasicOperation struct {
	Op Opcode
}

type ImmediateOperation struct {
	Op  Opcode
	Val Value
}

func NewOpcodeFromReader(rd io.Reader) (Opcode, error) {
	var ret Opcode
	if err := binary.Read(rd, binary.LittleEndian, &ret); err != nil {
		return 0, err
	}
	return ret, nil
}

func (o Opcode) Marshal(wr io.Writer) error {
	return binary.Write(wr, binary.LittleEndian, &o)
}

func NewBasicOperationFromReader(rd io.Reader) (BasicOperation, error) {
	op, err := NewOpcodeFromReader(rd)
	return BasicOperation{op}, err
}

func NewImmediateOperationFromReader(rd io.Reader) (ImmediateOperation, error) {
	op, err := NewOpcodeFromReader(rd)
	if err != nil {
		return ImmediateOperation{}, err
	}
	val, err := UnmarshalValue(rd)
	return ImmediateOperation{op, val}, err
}

func (op BasicOperation) Marshal(wr io.Writer) error {
	return op.Op.Marshal(wr)
}

func (op BasicOperation) MarshalProof(wr io.Writer, includeVal bool) error {
	return op.Op.Marshal(wr)
}

func (op ImmediateOperation) Marshal(wr io.Writer) error {
	if err := op.Op.Marshal(wr); err != nil {
		return err
	}
	return MarshalValue(op.Val, wr)
}

func (op ImmediateOperation) MarshalProof(wr io.Writer, includeVal bool) error {
	if err := op.Op.Marshal(wr); err != nil {
		return err
	}
	if includeVal {
		return MarshalValue(op.Val.CloneShallow(), wr)
	}
	return MarshalValue(NewHashOnlyValueFromValue(op.Val), wr)
}

func (op BasicOperation) TypeCode() uint8 {
	return 0
}

func (op ImmediateOperation) TypeCode() uint8 {
	return 1
}

func (op BasicOperation) GetOp() Opcode {
	return op.Op
}

// func (op BasicOperation) String() string {
//	return fmt.Sprintf("Basic(%v)", code.InstructionNames[op.Op])
//}
//
// func (op ImmediateOperation) String() string {
//	return fmt.Sprintf("Immediate(%v, %v)", code.InstructionNames[op.Op], op.Val)
//}

func (op ImmediateOperation) GetOp() Opcode {
	return op.Op
}

func NewOperationFromReader(rd io.Reader) (Operation, error) {
	var immediateCount uint8
	err := binary.Read(rd, binary.BigEndian, &immediateCount)
	if err != nil {
		return nil, err
	}
	if immediateCount == 0 {
		return NewBasicOperationFromReader(rd)
	} else if immediateCount == 1 {
		return NewImmediateOperationFromReader(rd)
	} else {
		return nil, errors.New("immediate count must be 0 or 1")
	}
}

func MarshalOperation(op Operation, wr io.Writer) error {
	typ := op.TypeCode()
	if err := binary.Write(wr, binary.BigEndian, &typ); err != nil {
		return err
	}
	return op.Marshal(wr)
}

func MarshalOperationProof(op Operation, wr io.Writer, includeVal bool) error {
	typ := op.TypeCode()
	if err := binary.Write(wr, binary.BigEndian, &typ); err != nil {
		return err
	}
	return op.MarshalProof(wr, includeVal)
}

func NewCodePointForProofFromReader(rd io.Reader) (CodePointValue, error) {
	var op Operation
	op, err := NewOperationFromReader(rd)
	if err != nil {
		return CodePointValue{}, err
	}
	var nextHash [32]byte
	_, err = io.ReadFull(rd, nextHash[:])
	return CodePointValue{0, op, nextHash}, err
}

type CodePointValue struct {
	InsnNum  int64
	Op       Operation
	NextHash [32]byte
}

func NewCodePointValueFromReader(rd io.Reader) (CodePointValue, error) {
	var insnNum int64
	if err := binary.Read(rd, binary.BigEndian, &insnNum); err != nil {
		return CodePointValue{}, err
	}
	var op Operation
	op, err := NewOperationFromReader(rd)
	if err != nil {
		return CodePointValue{}, err
	}
	var nextHash [32]byte
	_, err = io.ReadFull(rd, nextHash[:])
	return CodePointValue{insnNum, op, nextHash}, err
}

func (cv CodePointValue) TypeCode() uint8 {
	return TypeCodeCodePoint
}

func (cv CodePointValue) InternalTypeCode() uint8 {
	return TypeCodeCodePoint
}

func (cv CodePointValue) Clone() Value {
	return CodePointValue{cv.InsnNum, cv.Op, cv.NextHash}
}

func (cv CodePointValue) CloneShallow() Value {
	return CodePointValue{cv.InsnNum, cv.Op, cv.NextHash}
}

func (cv CodePointValue) Equal(val Value) bool {
	if val.TypeCode() == TypeCodeHashOnly {
		return cv.Hash() == val.Hash()
	} else if val.TypeCode() != TypeCodeCodePoint {
		return false
	} else {
		if cv.InsnNum != val.(CodePointValue).InsnNum {
			return false
		}
		// for now only check InsnNum
		// if cv.Op != val.(CodePointValue).Op {
		//	return false
		//}
		// if cv.NextHash != val.(CodePointValue).NextHash {
		//	return false
		//}
		return true
	}
}

func (cv CodePointValue) Size() int64 {
	return 1
}

var ErrorCodePoint CodePointValue

func init() {
	ErrorCodePoint = CodePointValue{0, BasicOperation{0}, [32]byte{}}
}

func (cv CodePointValue) Hash() [32]byte {
	switch op := cv.Op.(type) {
	case ImmediateOperation:
		hash := [32]byte{}
		copy(hash[:], solsha3.SoliditySHA3(
			solsha3.Uint8(TypeCodeCodePoint),
			solsha3.Uint8(byte(op.Op)),
			solsha3.Bytes32(op.Val.Hash()),
			solsha3.Bytes32(cv.NextHash),
		))
		return hash
	case BasicOperation:
		hash := [32]byte{}
		copy(hash[:], solsha3.SoliditySHA3(
			solsha3.Uint8(TypeCodeCodePoint),
			solsha3.Uint8(byte(op.Op)),
			solsha3.Bytes32(cv.NextHash),
		))
		return hash
	default:
		panic(fmt.Sprintf("Bad operation type: %T in with pc %d", op, cv.InsnNum))
	}
}

func (cv CodePointValue) Marshal(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, &cv.InsnNum); err != nil {
		return err
	}
	if err := cv.Op.Marshal(w); err != nil {
		return err
	}
	_, err := w.Write(cv.NextHash[:])
	return err
}

func (cv CodePointValue) MarshalForProof(w io.Writer) error {
	if err := cv.Op.Marshal(w); err != nil {
		return err
	}
	_, err := w.Write(cv.NextHash[:])
	return err
}

func (cv CodePointValue) String() string {
	return fmt.Sprintf("CodePoint(%v, %v)", cv.InsnNum, cv.Op)
}
