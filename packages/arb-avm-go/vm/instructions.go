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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/code"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Instruction struct {
	code value.Opcode
	impl func(*Machine) (StackMods, error)
}

var allInsns = []Instruction{ // code, not necessarily in order
	{code.ADD, insnAdd},
	{code.MUL, insnMul},
	{code.SUB, insnSub},
	{code.DIV, insnDiv},
	{code.SDIV, insnSdiv},
	{code.MOD, insnMod},
	{code.SMOD, insnSmod},
	{code.ADDMOD, insnAddmod},
	{code.MULMOD, insnMulmod},
	{code.EXP, insnExp},

	{code.LT, insnLt},
	{code.GT, insnGt},
	{code.SLT, insnSlt},
	{code.SGT, insnSgt},
	{code.EQ, insnEq},
	{code.ISZERO, insnIszero},
	{code.AND, insnAnd},
	{code.OR, insnOr},
	{code.XOR, insnXor},
	{code.NOT, insnNot},
	{code.BYTE, insnByte},
	{code.SIGNEXTEND, insnSignextend},

	{code.SHA3, insnHash},
	{code.TYPE, insnType},

	{code.POP, insnPop},
	{code.SPUSH, insnSpush},
	{code.RPUSH, insnRpush},
	{code.RSET, insnRset},
	{code.JUMP, insnJump},
	{code.CJUMP, insnCjump},
	{code.STACKEMPTY, insnStackempty},
	{code.PCPUSH, insnPcpush},
	{code.AUXPUSH, insnAuxpush},
	{code.AUXPOP, insnAuxpop},
	{code.AUXSTACKEMPTY, insnAuxStackempty},
	{code.NOP, insnNop},
	{code.ERRPUSH, insnErrPush},
	{code.ERRSET, insnErrSet},

	{code.DUP0, insnDup0},
	{code.DUP1, insnDup1},
	{code.DUP2, insnDup2},
	{code.SWAP1, insnSwap1},
	{code.SWAP2, insnSwap2},

	{code.TGET, insnTget},
	{code.TSET, insnTset},
	{code.TLEN, insnTlen},

	{code.BREAKPOINT, insnBreakpoint},
	{code.LOG, insnLog},

	{code.SEND, insnSend},
	{code.NBSEND, insnNBSend},
	{code.GETTIME, insnGettime},
	{code.INBOX, insnInbox},
	{code.ERROR, insnError},
	{code.HALT, insnHalt},
	{code.DEBUG, insnDebug},
}

var (
	bigZero = new(big.Int)
	tt255   = math.BigPow(2, 255)
	tt256   = math.BigPow(2, 256)
	tt256m1 = new(big.Int).Sub(tt256, big.NewInt(1))
)

var Instructions = []Instruction(nil)

func init() {
	Instructions = make([]Instruction, code.MaxOpcode)
	for _, ins := range allInsns {
		Instructions[ins.code] = ins
	}
}

type BlockedError struct {
	reason machine.BlockReason
}

func (w BlockedError) Error() string {
	return "VMBlockederror"
}

func RunInstruction(m *Machine, op value.Operation) (StackMods, machine.BlockReason) {
	if m.IsHalted() {
		return NewStackMods(0, 0), machine.HaltBlocked{}
	}
	if m.IsErrored() {
		return NewStackMods(0, 0), machine.ErrorBlocked{}
	}
	if m.HaveSizeException() {
		return NewStackMods(0, 0), machine.ErrorBlocked{}
	}
	m.context.NotifyStep()
	mods, err := func() (StackMods, error) {
		if _, ok := code.InstructionNames[op.GetOp()]; !ok {
			return StackMods{}, errors.New("invalid opcode")
		}

		if immediate, ok := op.(value.ImmediateOperation); ok {
			m.stack.Push(immediate.Val)
		}

		return Instructions[op.GetOp()].impl(m)
	}()

	if err == nil {
		return mods, nil
	}

	if blocked, isBlocked := err.(BlockedError); isBlocked {
		return mods, blocked.reason
	}

	//fmt.Printf("error running instruction %v: %v\n", code.InstructionNames[op.GetOp()], err)

	// in case of any errors from operation
	// pop remaining stack values and set
	// PC to errHandler
	m.Warn(err.Error())
	for mods.popsRemaining > 0 {
		var poperr error
		_, mods, poperr = PopStackBox(m, mods)
		if poperr != nil {
			break
		}
	}

	// Clear the error by jumping to the error handler
	if !m.errHandler.Equal(value.ErrorCodePoint) {
		err = m.pc.SetPCForced(m.errHandler)
	}
	if err != nil {
		m.ErrorStop()
	}
	return mods, nil
}

func (insn Instruction) GetName() string {
	return code.InstructionNames[insn.code]
}

func (insn Instruction) GetCode() value.Opcode {
	return insn.code
}

const MaxStackPops = 3
const MaxAuxStackPops = 1

type StackMods struct {
	popsRemaining         int
	pushesRemaining       int
	stackPopsPerformed    int
	auxStackPopsPerformed int
	stackPopTypes         [MaxStackPops]byte
	auxStackPopTypes      [MaxAuxStackPops]byte
}

func NewStackMods(pops, pushes int) StackMods {
	return StackMods{
		pops,
		pushes,
		0,
		0,
		[MaxStackPops]byte{},
		[MaxAuxStackPops]byte{},
	}
}

func (m *StackMods) removePop() {
	for i := 1; i < MaxStackPops; i++ {
		m.stackPopTypes[i-1] = m.stackPopTypes[i]
	}
	m.stackPopsPerformed--
}

func (m StackMods) stackPopInfo() []byte {
	return m.stackPopTypes[:m.stackPopsPerformed]
}

func (m StackMods) auxStackPopInfo() []byte {
	return m.auxStackPopTypes[:m.auxStackPopsPerformed]
}

func (m StackMods) check() {
	if m.popsRemaining != 0 || m.pushesRemaining != 0 {
		panic("Instruction left stack with nonzero StackMods")
	}
}

func (m *StackMods) poppedValue() {
	m.popsRemaining--
	m.stackPopTypes[m.stackPopsPerformed] = 1
	m.stackPopsPerformed++
}

func PushStackBox(m *Machine, mods StackMods, b value.Value) StackMods {
	mods.pushesRemaining--
	m.Stack().Push(b)
	return mods
}

func PushStackInt(m *Machine, mods StackMods, v value.IntValue) StackMods {
	mods.pushesRemaining--
	m.Stack().PushInt(v)
	return mods
}

func PushStackTuple(m *Machine, mods StackMods, v value.TupleValue) StackMods {
	mods.pushesRemaining--
	m.Stack().PushTuple(v)
	return mods
}

func PushStackCodePoint(m *Machine, mods StackMods, v value.CodePointValue) StackMods {
	mods.pushesRemaining--
	m.Stack().PushCodePoint(v)
	return mods
}

func PopAuxStack(m *Machine, mods StackMods) (value.Value, StackMods, error) {
	if m.AuxStack().IsEmpty() {
		return value.NewEmptyTuple(), mods, nil
	}

	mods.auxStackPopTypes[mods.auxStackPopsPerformed] = 0
	mods.auxStackPopsPerformed++
	b, err := m.AuxStack().Pop()
	return b, mods, err
}

type PopTypeWarning struct {
	msg  string
	mods StackMods
}

func (w PopTypeWarning) Error() string {
	return w.msg
}

type EmptyStackError struct {
	mods StackMods
}

func (w EmptyStackError) Error() string {
	return "Tried to pop empty stack"
}

func PopStackBox(m *Machine, mods StackMods) (value.Value, StackMods, error) {
	if m.Stack().IsEmpty() {
		return value.NewEmptyTuple(), mods, EmptyStackError{mods}
	}
	mods.popsRemaining--
	mods.stackPopTypes[mods.stackPopsPerformed] = 0
	mods.stackPopsPerformed++
	b, _ := m.Stack().Pop()
	return b, mods, nil
}

func PopStackValue(m *Machine, mods StackMods) (value.Value, StackMods, error) {
	if m.Stack().IsEmpty() {
		return value.NewEmptyTuple(), mods, EmptyStackError{mods}
	}
	v, err := m.Stack().Pop()
	mods.poppedValue()
	return v, mods, err
}

func PopStackInt(m *Machine, mods StackMods) (value.IntValue, StackMods, error) {
	v, err := m.Stack().PopInt()
	mods.poppedValue()
	return v, mods, err
}

func PopStackTuple(m *Machine, mods StackMods) (value.TupleValue, StackMods, error) {
	v, err := m.Stack().PopTuple()
	mods.poppedValue()
	return v, mods, err
}

func PopStackCodePoint(m *Machine, mods StackMods) (value.CodePointValue, StackMods, error) {
	v, err := m.Stack().PopCodePoint()
	mods.poppedValue()
	return v, mods, err
}

func unaryIntOp(state *Machine, intOp func(value.IntValue) (value.IntValue, error)) (StackMods, error) {
	mods := NewStackMods(1, 1)
	arg1, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	r, err := intOp(arg1)
	if err == nil {
		mods = PushStackInt(state, mods, r)
		state.IncrPC()
	}
	return mods, err
}

func binaryIntOp(state *Machine, intOp func(value.IntValue, value.IntValue) (value.IntValue, error)) (StackMods, error) {
	mods := NewStackMods(2, 1)
	arg1, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	arg2, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	r, err := intOp(arg1, arg2)
	if err == nil {
		mods = PushStackInt(state, mods, r)
		state.IncrPC()
	}
	return mods, err
}

func trinaryIntOp(state *Machine, intOp func(value.IntValue, value.IntValue, value.IntValue) (value.IntValue, error)) (StackMods, error) {
	mods := NewStackMods(3, 1)
	arg1, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	arg2, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	arg3, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}
	r, err := intOp(arg1, arg2, arg3)
	if err == nil {
		mods = PushStackInt(state, mods, r)
		state.IncrPC()
	}
	return mods, err
}

func insnHalt(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 0)
	state.Halt()
	return mods, nil
}

// BEGIN STUB OPS

func insnAdd(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Add(x.BigInt(), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnMul(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Mul(x.BigInt(), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnSub(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Sub(new(big.Int).Add(x.BigInt(), tt256), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

type DivideByZeroError struct {
}

func (w DivideByZeroError) Error() string {
	return "Tried to divide or modulo by zero"
}

func insnDiv(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			yBig := y.BigInt()
			if yBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := new(big.Int).Div(x.BigInt(), yBig)
			return value.NewIntValue(ret), nil
		})
}

func insnSdiv(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			yBig := y.BigInt()
			if yBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := math.U256(new(big.Int).Div(math.S256(x.BigInt()), math.S256(yBig)))
			return value.NewIntValue(ret), nil
		})
}

func insnMod(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			yBig := y.BigInt()
			if yBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := new(big.Int).Mod(x.BigInt(), yBig)
			return value.NewIntValue(ret), nil
		})
}

func insnSmod(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			xBig := math.S256(x.BigInt())
			yBig := math.S256(y.BigInt())
			if yBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := new(big.Int).Mul(big.NewInt(int64(xBig.Sign())), new(big.Int).Mod(new(big.Int).Abs(xBig), new(big.Int).Abs(yBig)))
			return value.NewIntValue(math.U256(ret)), nil
		})
}

func insnAddmod(state *Machine) (StackMods, error) {
	return trinaryIntOp(state,
		func(x, y, z value.IntValue) (value.IntValue, error) {
			zBig := z.BigInt()
			if zBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := math.U256(new(big.Int).Mod(new(big.Int).Add(x.BigInt(), y.BigInt()), zBig))
			return value.NewIntValue(ret), nil
		})
}

func insnMulmod(state *Machine) (StackMods, error) {
	return trinaryIntOp(state,
		func(x, y, z value.IntValue) (value.IntValue, error) {
			zBig := z.BigInt()
			if zBig.Sign() == 0 {
				return value.IntegerZero, DivideByZeroError{}
			}
			ret := math.U256(new(big.Int).Mod(new(big.Int).Mul(x.BigInt(), y.BigInt()), zBig))
			return value.NewIntValue(ret), nil
		})
}

func insnExp(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(base, exponent value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Exp(base.BigInt(), exponent.BigInt(), tt256))
			return value.NewIntValue(ret), nil
		})
}

func insnSignextend(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(num, back value.IntValue) (value.IntValue, error) {
			bBig := back.BigInt()
			if !bBig.IsInt64() {
				return num, nil
			}
			b64 := bBig.Int64()
			if b64 > 31 {
				return num, nil
			}
			t := 248 - 8*b64
			numBi := num.BigInt()
			signBit := numBi.Bit(int(255 - t))
			mask := new(big.Int).Sub(math.BigPow(2, 255-t), big.NewInt(1))
			var ret *big.Int
			if signBit == 0 {
				ret = new(big.Int).And(num.BigInt(), mask)
			} else {
				mask = new(big.Int).Xor(tt256m1, mask)
				ret = new(big.Int).Or(num.BigInt(), mask)
			}
			return value.NewIntValue(math.U256(ret)), nil
		})
}

func insnLt(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			return value.NewBooleanValue(x.BigInt().Cmp(y.BigInt()) == -1), nil
		})
}

func insnGt(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			return value.NewBooleanValue(x.BigInt().Cmp(y.BigInt()) == 1), nil
		})
}

func insnSlt(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			return value.NewBooleanValue(math.S256(x.BigInt()).Cmp(math.S256(y.BigInt())) == -1), nil
		})
}

func insnSgt(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			return value.NewBooleanValue(math.S256(x.BigInt()).Cmp(math.S256(y.BigInt())) == 1), nil
		})
}

func insnIszero(state *Machine) (StackMods, error) {
	return unaryIntOp(state,
		func(x value.IntValue) (value.IntValue, error) {
			return value.NewBooleanValue(x.BigInt().Sign() == 0), nil
		})
}

func insnAnd(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).And(x.BigInt(), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnOr(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Or(x.BigInt(), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnXor(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(x, y value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Xor(x.BigInt(), y.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnNot(state *Machine) (StackMods, error) {
	return unaryIntOp(state,
		func(x value.IntValue) (value.IntValue, error) {
			ret := math.U256(new(big.Int).Not(x.BigInt()))
			return value.NewIntValue(ret), nil
		})
}

func insnByte(state *Machine) (StackMods, error) {
	return binaryIntOp(state,
		func(val, th value.IntValue) (value.IntValue, error) {
			thBig := th.BigInt()
			if !thBig.IsUint64() {
				return value.IntegerZero, nil
			}
			th64 := thBig.Uint64()
			if th64 >= 32 {
				return value.IntegerZero, nil
			}
			ret := math.Byte(val.BigInt(), value.BytesPerInt, int(th64))
			return value.NewInt64Value(int64(ret)), nil
		})
}

// END STUB OPS

func insnEq(state *Machine) (StackMods, error) {
	mods := NewStackMods(2, 1)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	y, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackInt(state, mods, value.NewBooleanValue(value.Eq(x, y)))
	state.IncrPC()
	return mods, nil
}

func insnHash(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 1)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	hashVal := x.Hash()
	intVal := big.NewInt(0)
	intVal.SetBytes(hashVal[:])
	mods = PushStackInt(state, mods, value.NewIntValue(intVal))
	state.IncrPC()
	return mods, nil
}

func insnPop(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)
	_, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	state.IncrPC()
	return mods, nil
}

func insnSpush(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackBox(state, mods, state.Static().Get())
	state.IncrPC()
	return mods, nil
}

func insnRpush(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackBox(state, mods, state.Register().Get())
	state.IncrPC()
	return mods, nil
}

func insnRset(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	state.Register().Set(x)
	state.IncrPC()
	return mods, nil
}

func insnInbox(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 1)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	inboxVal := state.ReadInbox()
	mods = PushStackBox(state, mods, inboxVal)
	if value.Eq(x, inboxVal) {
		return mods, BlockedError{machine.InboxBlocked{
			Inbox: value.NewHashOnlyValueFromValue(inboxVal),
		}}
	}
	state.IncrPC()
	return mods, nil
}

func insnErrPush(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackCodePoint(state, mods, state.errHandler)
	state.IncrPC()
	return mods, nil
}

func insnErrSet(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)

	rawTarget, mods, err := PopStackCodePoint(state, mods)
	if err != nil {
		return mods, err
	}
	state.errHandler = rawTarget
	state.IncrPC()
	return mods, err
}

type ErrorInstructionError struct {
}

func (w ErrorInstructionError) Error() string {
	return "Executed error instruction"
}

func insnError(*Machine) (StackMods, error) {
	mods := NewStackMods(0, 0)
	return mods, ErrorInstructionError{}
}

func insnJump(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)

	rawTarget, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	err = state.SetPC(rawTarget)
	return mods, err
}

func insnCjump(state *Machine) (StackMods, error) {
	mods := NewStackMods(2, 0)

	rawTarget, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}

	cond, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}

	if cond.BigInt().Cmp(big.NewInt(0)) != 0 {
		err := state.SetPC(rawTarget)
		return mods, err
	}
	state.IncrPC()
	return mods, nil
}

func insnPcpush(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackBox(state, mods, state.GetPC())
	state.IncrPC()
	return mods, nil
}

func insnAuxpush(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)
	val, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	state.AuxStack().Push(val)
	state.IncrPC()
	return mods, nil
}

func insnAuxpop(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	val, mods, err := PopAuxStack(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, val)
	state.IncrPC()
	return mods, nil
}

func insnStackempty(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackInt(state, mods, value.NewBooleanValue(state.stack.IsEmpty()))
	state.IncrPC()
	return mods, nil
}

func insnAuxStackempty(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackInt(state, mods, value.NewBooleanValue(state.auxstack.IsEmpty()))
	state.IncrPC()
	return mods, nil
}

func insnNop(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 0)
	state.IncrPC()
	return mods, nil
}

func insnDup0(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 2)
	// if s, ok := state.Stack().(*FlatStack); ok {
	//	mods.popsRemaining = 0
	//	mods.pushesRemaining = 0
	//	mods.callStackPopsPerformed = 0
	//	mods.stackPopsPerformed = 1
	//	mods.stackPopTypes = [3]byte{0}
	//	s.duplicate()
	//}
	var x value.Value
	var err error
	x, mods, err = PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, x)
	mods = PushStackBox(state, mods, x)
	state.IncrPC()
	return mods, nil
}

func insnDup1(state *Machine) (StackMods, error) {
	mods := NewStackMods(2, 3)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	y, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, y)
	mods = PushStackBox(state, mods, x)
	mods = PushStackBox(state, mods, y)
	state.IncrPC()
	return mods, nil
}

func insnDup2(state *Machine) (StackMods, error) {
	mods := NewStackMods(3, 4)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	y, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	z, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, z)
	mods = PushStackBox(state, mods, y)
	mods = PushStackBox(state, mods, x)
	mods = PushStackBox(state, mods, z)
	state.IncrPC()
	return mods, nil
}

func insnSwap1(state *Machine) (StackMods, error) {
	mods := NewStackMods(2, 2)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	y, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, x)
	mods = PushStackBox(state, mods, y)
	state.IncrPC()
	return mods, nil
}

func insnSwap2(state *Machine) (StackMods, error) {
	mods := NewStackMods(3, 3)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	y, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	z, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackBox(state, mods, x)
	mods = PushStackBox(state, mods, y)
	mods = PushStackBox(state, mods, z)
	state.IncrPC()
	return mods, nil
}

func insnTget(state *Machine) (StackMods, error) {
	mods := NewStackMods(2, 1)

	index, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}

	tuple, mods, err := PopStackTuple(state, mods)
	if err != nil {
		return mods, err
	}

	val, err := tuple.Get(index)
	if err != nil {
		// index out of range
		fmt.Println(state.stack)
		fmt.Println("pc = ", state.pc.GetPC())
		return mods, fmt.Errorf("insn_tget: index %v out of range %v", index.BigInt(), tuple.Len())
	}

	mods = PushStackBox(state, mods, val)
	state.IncrPC()
	return mods, nil
}

func insnTset(state *Machine) (StackMods, error) {
	mods := NewStackMods(3, 1)

	index, mods, err := PopStackInt(state, mods)
	if err != nil {
		return mods, err
	}

	tuple, mods, err := PopStackTuple(state, mods)
	if err != nil {
		return mods, err
	}

	newVal, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}

	newTup, err := tuple.Set(index, newVal)
	if err != nil {
		return mods, fmt.Errorf("insn_tset: index %v out of range of tuple %v", index, tuple)
	}

	mods = PushStackTuple(state, mods, newTup)
	state.IncrPC()
	return mods, nil
}

func insnTlen(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 1)
	tup, mods, err := PopStackTuple(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackInt(state, mods, value.NewInt64Value(tup.Len()))
	state.IncrPC()
	return mods, nil
}

func insnType(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 1)
	val, mods, err := PopStackValue(state, mods)
	if err != nil {
		return mods, err
	}
	mods = PushStackInt(state, mods, value.NewInt64Value(int64(val.TypeCode())))
	state.IncrPC()
	return mods, nil
}

func insnBreakpoint(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 0)
	state.IncrPC()
	return mods, BlockedError{machine.BreakpointBlocked{}}
}

func insnLog(state *Machine) (StackMods, error) {
	mods := NewStackMods(1, 0)
	x, mods, err := PopStackBox(state, mods)
	if err != nil {
		return mods, err
	}
	state.Log(x)
	state.IncrPC()
	return mods, nil
}

func sendImpl(state *Machine) (value.TupleValue, protocol.Message, StackMods, error) {
	mods := NewStackMods(1, 0)
	sendData, mods, err := PopStackTuple(state, mods)
	if err != nil {
		return sendData, protocol.Message{}, mods, err
	}

	if sendData.Len() != 4 {
		return sendData, protocol.Message{}, mods, err
	}

	data, _ := sendData.GetByInt64(0)
	val2, _ := sendData.GetByInt64(1)
	val3, _ := sendData.GetByInt64(2)
	val4, _ := sendData.GetByInt64(3)

	destination, ok2 := val2.(value.IntValue)
	amount, ok3 := val3.(value.IntValue)
	tokenType, ok4 := val4.(value.IntValue)

	if !ok2 || !ok3 || !ok4 {
		// mods, err := handlePopError(state, mods, PopTypeWarning{"Inbox pop tuple wrong", mods})
		return sendData, protocol.Message{}, mods, err
	}
	return sendData, protocol.NewMessage(data, protocol.TokenTypeFromIntValue(tokenType), amount.BigInt(), destination.ToBytes()), mods, nil
}

func insnSend(state *Machine) (StackMods, error) {
	sendData, msg, mods, err := sendImpl(state)
	if err != nil {
		return mods, err
	}

	err = state.Send(msg)
	if err != nil {
		state.stack.PushTuple(sendData)
		return mods, BlockedError{machine.SendBlocked{
			Currency:  msg.Currency,
			TokenType: msg.TokenType,
		}}
	}

	state.IncrPC()
	return mods, nil
}

func insnNBSend(state *Machine) (StackMods, error) {
	_, msg, mods, err := sendImpl(state)
	if err != nil {
		return mods, err
	}

	if err := state.Send(msg); err != nil {
		state.Warn(err.Error())
		mods = PushStackInt(state, mods, value.NewInt64Value(0))
	} else {
		mods = PushStackInt(state, mods, value.NewInt64Value(1))
	}

	state.IncrPC()
	return mods, nil
}

func insnGettime(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 1)
	mods = PushStackBox(state, mods, state.GetTimeBounds())
	state.IncrPC()
	return mods, nil
}

func insnDebug(state *Machine) (StackMods, error) {
	mods := NewStackMods(0, 0)
	state.IncrPC()
	return mods, nil
}
