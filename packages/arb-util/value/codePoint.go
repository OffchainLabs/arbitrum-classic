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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
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
		return MarshalValueForProof(op.Val.CloneShallow(), wr)
	}

	return MarshalValueForProof(NewHashOnlyValueFromValue(op.Val), wr)
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

func (op BasicOperation) String() string {
	return fmt.Sprintf("0x%x", op.GetOp())
}

func (op ImmediateOperation) String() string {
	return fmt.Sprintf("0x%x Imd(%v)", op.GetOp(), op.Val)
}

func (op ImmediateOperation) GetOp() Opcode {
	return op.Op
}

func NewOperationFromReader(rd io.Reader) (Operation, error) {
	var immediateCount uint8
	err := binary.Read(rd, binary.BigEndian, &immediateCount)
	if err != nil {
		return nil, err
	}
	switch immediateCount {
	case 0:
		return NewBasicOperationFromReader(rd)
	case 1:
		return NewImmediateOperationFromReader(rd)
	default:
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
	var nextHash common.Hash
	_, err = io.ReadFull(rd, nextHash[:])

	return CodePointValue{0, op, nextHash}, err
}

type CodePointValue struct {
	InsnNum  int64
	Op       Operation
	NextHash common.Hash
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
	var nextHash common.Hash
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
	switch val := val.(type) {
	case HashOnlyValue:
		return cv.Hash() == val.Hash()
	case CodePointValue:
		return cv.InsnNum == val.InsnNum
	default:
		return false
	}
}

func (cv CodePointValue) Size() int64 {
	return 1
}

var ErrorCodePoint CodePointValue

func init() {
	ErrorCodePoint = CodePointValue{0, BasicOperation{0}, common.Hash{}}
}

func (cv CodePointValue) Hash() common.Hash {
	switch op := cv.Op.(type) {
	case ImmediateOperation:
		return hashing.SoliditySHA3(
			hashing.Uint8(TypeCodeCodePoint),
			hashing.Uint8(byte(op.Op)),
			hashing.Bytes32(op.Val.Hash()),
			hashing.Bytes32(cv.NextHash),
		)
	case BasicOperation:
		return hashing.SoliditySHA3(
			hashing.Uint8(TypeCodeCodePoint),
			hashing.Uint8(byte(op.Op)),
			hashing.Bytes32(cv.NextHash),
		)
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
	if err := MarshalOperationProof(cv.Op, w, false); err != nil {
		return err
	}
	_, err := w.Write(cv.NextHash[:])
	return err
}

func (cv CodePointValue) String() string {
	return fmt.Sprintf("CodePoint(%v, %v)", cv.InsnNum, cv.Op)
}
