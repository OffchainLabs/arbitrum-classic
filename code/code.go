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

package code

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Opcode uint8

func NewOpcodeFromReader(rd io.Reader) (Opcode, error) {
	var ret Opcode
	if err := binary.Read(rd, binary.LittleEndian, &ret); err!=nil {
		return 0, err
	}
	if _,ok := InstructionNames[ret]; !ok {
		return ret, fmt.Errorf("Invalid opcode %v from reader", ret)
	}
	return ret, nil
}

func (o Opcode) Marshal(wr io.Writer) error {
	return binary.Write(wr, binary.LittleEndian, &o)
}

// 0x0 range - arithmetic ops.
const (
	HALT Opcode = iota
	ADD
	MUL
	SUB
	DIV
	SDIV
	MOD
	SMOD
	ADDMOD
	MULMOD
	EXP
	SIGNEXTEND
)

// 0x10 range - comparison ops.
const (
	LT Opcode = iota + 0x10
	GT
	SLT
	SGT
	EQ
	ISZERO
	AND
	OR
	XOR
	NOT
	BYTE

	SHA3 = 0x20
)

// 0x50 range - 'storage' and execution.
const (
	POP Opcode = 0x30 + iota
	SPUSH
	RPUSH
	RSET
	INBOX
	JUMP
	CJUMP
	STACKEMPTY
	PCPUSH
	AUXPUSH
	AUXPOP
	AUXSTACKEMPTY
	NOP
	ERRPUSH
	ERRSET
	ERROR
)

// 0x40 range.
const (
	DUP0 Opcode = 0x40 + iota
	DUP1
	DUP2
	SWAP1
	SWAP2
)

// 0x50 range.
const (
	TGET Opcode = 0x50 + iota
	TSET
	TLEN
	ISTUPLE
)

// 0xa0 range.
const (
	BREAKPOINT Opcode = 0x60 + iota
	LOG
)

// 0xf0 range.
const (
	SEND Opcode = 0x70 + iota
	NBSEND
	GETTIME
	DEBUG
)

const MaxOpcode = 0x7f

var InstructionNames = map[Opcode]string{
	HALT:       "halt",
	ADD:        "add",
	MUL:        "mul",
	SUB:        "sub",
	DIV:        "div",
	SDIV:       "sdiv",
	MOD:        "mod",
	SMOD:       "smod",
	ADDMOD:     "addmod",
	MULMOD:     "mulmod",
	EXP:        "exp",
	SIGNEXTEND: "signextend",

	LT:     "lt",
	GT:     "gt",
	SLT:    "slt",
	SGT:    "sgt",
	EQ:     "eq",
	ISZERO: "iszero",
	AND:    "and",
	OR:     "or",
	XOR:    "xor",
	NOT:    "not",
	BYTE:   "byte",

	SHA3: "hash",

	POP:        "pop",
	SPUSH:      "spush",
	RPUSH:      "rpush",
	RSET:       "rset",
	INBOX:      "inbox",
	JUMP:       "jump",
	CJUMP:      "cjump",
	STACKEMPTY: "stackempty",
	PCPUSH:     "pcpush",
	AUXPUSH:    "auxpush",
	AUXPOP:     "auxpop",
	AUXSTACKEMPTY: "auxstackempty",
	NOP:        "nop",
	ERRPUSH:    "errpush",
	ERRSET:     "errset",
	ERROR:	    "error",

	DUP0:  "dup0",
	DUP1:  "dup1",
	DUP2:  "dup2",
	SWAP1: "swap1",
	SWAP2: "swap2",

	TGET:    "tget",
	TSET:    "tset",
	TLEN:    "tlen",
	ISTUPLE: "istuple",

	BREAKPOINT:  "breakpoint",
	LOG: "log",

	SEND:      "send",
	NBSEND:    "nbsend",
	GETTIME:   "gettime",
	DEBUG:     "debug",
}

var InstructionStackPops = map[Opcode][]byte{
	HALT:       {},
	ADD:        {1, 1},
	MUL:        {1, 1},
	SUB:        {1, 1},
	DIV:        {1, 1},
	SDIV:       {1, 1},
	MOD:        {1, 1},
	SMOD:       {1, 1},
	ADDMOD:     {1, 1, 1},
	MULMOD:     {1, 1, 1},
	EXP:        {1, 1},
	SIGNEXTEND: {1, 1},

	LT:     {1, 1},
	GT:     {1, 1},
	SLT:    {1, 1},
	SGT:    {1, 1},
	EQ:     {0, 0},
	ISZERO: {1},
	AND:    {1, 1},
	OR:     {1, 1},
	XOR:    {1, 1},
	NOT:    {1},
	BYTE:   {1, 1},

	SHA3: {0},

	POP:        {0},
	SPUSH:      {},
	RPUSH:      {},
	RSET:       {0},
	INBOX:      {0},
	JUMP:       {0},
	CJUMP:      {0, 1},
	STACKEMPTY: {},
	PCPUSH:     {},
	AUXPUSH:    {0},
	AUXPOP:     {},
	AUXSTACKEMPTY: {},
	NOP:        {},
	ERRPUSH:    {},
	ERRSET:     {1},
	ERROR:	    {},

	DUP0:  {0},
	DUP1:  {0, 0},
	DUP2:  {0, 0, 0},
	SWAP1: {0, 0},
	SWAP2: {0, 0, 0},

	TGET:    {1, 1},
	TSET:    {1, 1, 0},
	TLEN:    {1},
	ISTUPLE: {1},

	BREAKPOINT:  {},
	LOG: {0},

	SEND:      {1},
	NBSEND:      {1},
	GETTIME:   {},
	DEBUG:     {},
}

var InstructionAuxStackPops = map[Opcode][]byte{
	HALT:       {},
	ADD:        {},
	MUL:        {},
	SUB:        {},
	DIV:        {},
	SDIV:       {},
	MOD:        {},
	SMOD:       {},
	ADDMOD:     {},
	MULMOD:     {},
	EXP:        {},
	SIGNEXTEND: {},

	LT:     {},
	GT:     {},
	SLT:    {},
	SGT:    {},
	EQ:     {},
	ISZERO: {},
	AND:    {},
	OR:     {},
	XOR:    {},
	NOT:    {},
	BYTE:   {},

	SHA3: {},

	POP:        {},
	SPUSH:      {},
	RPUSH:      {},
	RSET:       {},
	INBOX:      {},
	JUMP:       {},
	CJUMP:      {},
	STACKEMPTY: {},
	PCPUSH:     {},
	AUXPUSH:    {},
	AUXPOP:     {0},
	AUXSTACKEMPTY: {},
	NOP:        {},
	ERRPUSH:    {},
	ERRSET:     {},
	ERROR:	    {},

	DUP0:  {},
	DUP1:  {},
	DUP2:  {},
	SWAP1: {},
	SWAP2: {},

	TGET:    {},
	TSET:    {},
	TLEN:    {},
	ISTUPLE: {},

	BREAKPOINT:  {},
	LOG:   {},

	SEND:      {},
	NBSEND:    {},
	GETTIME:   {},
	DEBUG:     {},
}
