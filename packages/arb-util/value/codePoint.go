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

package value

import (
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"io"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Opcode uint8

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

type Operation interface {
	GetOp() Opcode
	TypeCode() uint8
	Equals(other Operation) bool
}

type BasicOperation struct {
	Op Opcode
}

func NewBasicOperationFromReader(rd io.Reader) (BasicOperation, error) {
	op, err := NewOpcodeFromReader(rd)
	return BasicOperation{op}, err
}

func (op BasicOperation) TypeCode() uint8 {
	return 0
}

func (op BasicOperation) GetOp() Opcode {
	return op.Op
}

func (op BasicOperation) Equals(other Operation) bool {
	o, ok := other.(BasicOperation)
	if !ok {
		return false
	}
	return op.Op == o.Op
}

func (op BasicOperation) String() string {
	return fmt.Sprintf("0x%x", op.GetOp())
}

type ImmediateOperation struct {
	Op  Opcode
	Val Value
}

func NewImmediateOperationFromReader(rd io.Reader) (ImmediateOperation, error) {
	op, err := NewOpcodeFromReader(rd)
	if err != nil {
		return ImmediateOperation{}, err
	}
	val, err := UnmarshalValue(rd)
	return ImmediateOperation{op, val}, err
}

func (op ImmediateOperation) TypeCode() uint8 {
	return 1
}

func (op ImmediateOperation) String() string {
	return fmt.Sprintf("0x%x Imd(%v)", op.GetOp(), op.Val)
}

func (op ImmediateOperation) GetOp() Opcode {
	return op.Op
}

func (op ImmediateOperation) Equals(other Operation) bool {
	o, ok := other.(ImmediateOperation)
	if !ok {
		return false
	}
	return op.Op == o.Op && Eq(op.Val, o.Val)
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

type CodePointValue struct {
	Op       Operation
	NextHash common.Hash
}

func NewCodePointValueFromReader(rd io.Reader) (CodePointValue, error) {
	var op Operation
	op, err := NewOperationFromReader(rd)
	if err != nil {
		return CodePointValue{}, err
	}
	var nextHash common.Hash
	_, err = io.ReadFull(rd, nextHash[:])
	return CodePointValue{op, nextHash}, err
}

func (cv CodePointValue) TypeCode() uint8 {
	return TypeCodeCodePoint
}

func (cv CodePointValue) Equal(val Value) bool {
	o, ok := val.(CodePointValue)
	if !ok {
		return false
	}
	return cv.NextHash == o.NextHash && cv.Op.Equals(o.Op)
}

func (cv CodePointValue) Size() int64 {
	return 1
}

func (cv CodePointValue) String() string {
	return fmt.Sprintf("CodePoint(%v)", cv.Op)
}
