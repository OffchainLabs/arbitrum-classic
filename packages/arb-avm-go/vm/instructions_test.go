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
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/code"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

// This is to test that a machine can be built and run
// It creates a macine with 4 steps and runs it
// There is no automated test check so pass/fail must be verified visually
func TestMachineAdd(t *testing.T) {
	insns := []value.Operation{
		value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(2)},
		value.ImmediateOperation{Op: code.ADD, Val: value.NewInt64Value(4)},
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	m.ExecuteAssertion(80000, protocol.NewTimeBounds(0, 100000))
}

func runInstOpNoFault(m *Machine, oper value.Operation) (bool, string) {
	if _, blockReason := RunInstruction(m, oper); blockReason != nil {
		return false, fmt.Sprintf("RunInstruction blocked: %#v", blockReason)
	}

	if m.status != machine.Extensive {
		return false, fmt.Sprintf("RunInstruction should have succeeded, but had bad status: %v", m.status)
	}

	return true, ""
}

func runInstNoFault(m *Machine, oper value.Opcode) (bool, string) {
	return runInstOpNoFault(m, value.BasicOperation{Op: oper})
}

func runInstWithError(m *Machine, oper value.Opcode) (bool, string) {
	if _, blockReason := RunInstruction(m, value.BasicOperation{Op: value.Opcode(oper)}); blockReason != nil {
		return false, fmt.Sprintf("RunInstruction blocked: %#v", blockReason)
	}

	if m.status != machine.ErrorStop {
		return false, fmt.Sprintf("RunInstruction should have errored, but had status: %v", m.status)
	}

	return true, ""
}

// base operation tests for one, two, or three operands
// Push the required number of operands (from passed in values)
// Run the given instruction
// Push the expected result to the stack of a second machine
// Compare the two machines
func naryValueOpTest(vals []value.Value, expected value.Value, oper value.Opcode) (bool, string) {
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}
	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	for _, val := range vals {
		m.Stack().Push(val)
	}

	if succeeded, reason := runInstNoFault(m, oper); !succeeded {
		return succeeded, reason
	}

	knownMachine.Stack().Push(expected)
	if ok, err := Equal(knownMachine, m); !ok {
		tmp := "machines not equal: "
		tmp += err
		return false, tmp
	}

	return true, ""
}

func unaryIntOpTest(x, expected *big.Int, oper value.Opcode) (bool, string) {
	return naryValueOpTest([]value.Value{value.NewIntValue(x)}, value.NewIntValue(expected), oper)
}

func binaryIntOpTest(x, y, expected *big.Int, oper value.Opcode) (bool, string) {
	return naryValueOpTest([]value.Value{value.NewIntValue(y), value.NewIntValue(x)}, value.NewIntValue(expected), oper)
}

func binaryValueOpTest(x, y value.Value, expected *big.Int, oper value.Opcode) (bool, string) {
	return naryValueOpTest([]value.Value{y, x}, value.NewIntValue(expected), oper)
}

func tertiaryIntOpTest(x, y, z, expected *big.Int, oper value.Opcode) (bool, string) {
	return naryValueOpTest(
		[]value.Value{value.NewIntValue(z), value.NewIntValue(y), value.NewIntValue(x)},
		value.NewIntValue(expected),
		oper,
	)
}

// This test is to test an operation missing the second value
func TestAddMissingValue(t *testing.T) {
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(1))

	if failed, reason := runInstWithError(m, code.ADD); !failed {
		// Tried to pop empty stack
		t.Error(reason)
	}
	knownMachine.Stack().Push(value.NewInt64Value(2))
	if ok, _ := Equal(knownMachine, m); ok {
		tmp := "machines equal expected different"
		t.Error(tmp)
	}
}

//*************************************
// These are the tests for each opcode
//*************************************
func TestAdd(t *testing.T) {
	// test 3+4=7
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(4), big.NewInt(7), code.ADD)
	if !res {
		t.Error(err)
	}
	// test 0+0=0
	res, err = binaryIntOpTest(big.NewInt(0), big.NewInt(0), big.NewInt(0), code.ADD)
	if !res {
		t.Error(err)
	}
	// test (2**256-1)+4=3
	res, err = binaryIntOpTest(math.U256(big.NewInt(-1)), big.NewInt(4), big.NewInt(3), code.ADD)
	if !res {
		t.Error(err)
	}
	// test (2**256-2)+1=0xffff...ff
	res, err = binaryIntOpTest(math.U256(big.NewInt(-2)), big.NewInt(1), math.U256(big.NewInt(-1)), code.ADD)
	if !res {
		t.Error(err)
	}
	// test 3 + tuple = error
	tup := value.NewEmptyTuple()
	res, err = binaryValueOpTest(value.NewInt64Value(3), tup, big.NewInt(7), code.ADD)
	if res {
		t.Error("expected error")
	}
}

func TestMul(t *testing.T) {
	// test 3*4=12
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(4), big.NewInt(12), code.MUL)
	if !res {
		t.Error(err)
	}
	// test 3*0=0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(0), big.NewInt(0), code.MUL)
	if !res {
		t.Error(err)
	}
	// test (2**256-1)*1=(2**256-1)
	res, err = binaryIntOpTest(math.U256(big.NewInt(-1)), big.NewInt(1), math.U256(big.NewInt(-1)), code.MUL)
	if !res {
		t.Error(err)
	}
	// test (2**256-1)*2=(2**256-2)
	res, err = binaryIntOpTest(math.U256(big.NewInt(-1)), big.NewInt(2), math.U256(big.NewInt(-2)), code.MUL)
	if !res {
		t.Error(err)
	}
}

func TestSub(t *testing.T) {
	// test 4-3=1
	res, err := binaryIntOpTest(big.NewInt(4), big.NewInt(3), big.NewInt(1), code.SUB)
	if !res {
		t.Error(err)
	}
	// test 0-1=0xffff...ff
	res, err = binaryIntOpTest(big.NewInt(0), big.NewInt(1), math.U256(big.NewInt(-1)), code.SUB)
	if !res {
		t.Error(err)
	}
}

func TestDiv(t *testing.T) {
	// test 6/2=3
	res, err := binaryIntOpTest(big.NewInt(6), big.NewInt(2), big.NewInt(3), code.DIV)
	if !res {
		t.Error(err)
	}
	// test -6/2!=-3 (should be unsigned division)
	res, err = binaryIntOpTest(math.U256(big.NewInt(-6)), big.NewInt(2), math.U256(big.NewInt(-3)), code.DIV)
	if res {
		t.Error("should not be -3")
	}
	// test 6/0=0
	res, err = binaryIntOpTest(big.NewInt(6), big.NewInt(0), big.NewInt(0), code.DIV)
	if res {
		t.Error("Divide by 0 expected")
	}
}

func TestSDiv(t *testing.T) {
	// test -6/3=-2
	res, err := binaryIntOpTest(math.U256(big.NewInt(-6)), math.U256(big.NewInt(3)), math.U256(big.NewInt(-2)), code.SDIV)
	if !res {
		t.Error(err)
	}
	// test 6/-3=-2
	res, err = binaryIntOpTest(math.U256(big.NewInt(6)), math.U256(big.NewInt(-3)), math.U256(big.NewInt(-2)), code.SDIV)
	if !res {
		t.Error(err)
	}
	// test -6/-3=2
	res, err = binaryIntOpTest(math.U256(big.NewInt(-6)), math.U256(big.NewInt(-3)), math.U256(big.NewInt(2)), code.SDIV)
	if !res {
		t.Error(err)
	}
	// test 6/3=2
	res, err = binaryIntOpTest(math.U256(big.NewInt(6)), math.U256(big.NewInt(3)), math.U256(big.NewInt(2)), code.SDIV)
	if !res {
		t.Error(err)
	}
	// test 6/0=0
	res, err = binaryIntOpTest(big.NewInt(6), big.NewInt(0), big.NewInt(0), code.SDIV)
	if res {
		t.Error("Divide by 0 expected")
	}
}

func TestMod(t *testing.T) {
	// test 8%3=2
	res, err := binaryIntOpTest(big.NewInt(8), big.NewInt(3), big.NewInt(2), code.MOD)
	if !res {
		t.Error(err)
	}
	// test 0%3=0
	res, err = binaryIntOpTest(big.NewInt(0), big.NewInt(3), big.NewInt(0), code.MOD)
	if !res {
		t.Error(err)
	}
	// test 3%8=0 - divide by zero error
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(0), big.NewInt(0), code.MOD)
	if res {
		t.Error("Divide by 0 expected")
	}
}

func TestSMod(t *testing.T) {
	// test -8%3=-2
	res, err := binaryIntOpTest(math.U256(big.NewInt(-8)), math.U256(big.NewInt(3)), math.U256(big.NewInt(-2)), code.SMOD)
	if !res {
		t.Error(err)
	}
	// test -8%-3=-2 (per spec modulo by a negative number takes the sign of the dividend
	res, err = binaryIntOpTest(math.U256(big.NewInt(-8)), math.U256(big.NewInt(-3)), math.U256(big.NewInt(-2)), code.SMOD)
	if !res {
		t.Error(err)
	}
	// test 8%3=2
	res, err = binaryIntOpTest(math.U256(big.NewInt(8)), math.U256(big.NewInt(3)), math.U256(big.NewInt(2)), code.SMOD)
	if !res {
		t.Error(err)
	}
	// test -8%0=0
	res, err = binaryIntOpTest(math.U256(big.NewInt(-8)), math.U256(big.NewInt(0)), math.U256(big.NewInt(0)), code.SMOD)
	if res {
		t.Error("Divide by 0 expected")
	}
}

func TestAddMod(t *testing.T) {
	// test (8+5)%3=1
	res, err := tertiaryIntOpTest(big.NewInt(8), big.NewInt(5), big.NewInt(3), big.NewInt(1), code.ADDMOD)
	if !res {
		t.Error(err)
	}
	// test ((2**256-1)+1)%7=2 - shows that internal addition does not 256 bit truncate
	res, err = tertiaryIntOpTest(math.U256(big.NewInt(-1)), big.NewInt(1), big.NewInt(7), big.NewInt(2), code.ADDMOD)
	if !res {
		t.Error(err)
	}
	// test (0+0)%7=0
	res, err = tertiaryIntOpTest(math.U256(big.NewInt(0)), big.NewInt(0), big.NewInt(7), big.NewInt(0), code.ADDMOD)
	if !res {
		t.Error(err)
	}
}

func TestMulMod(t *testing.T) {
	// test (8*2)%3=1
	res, err := tertiaryIntOpTest(big.NewInt(8), big.NewInt(2), big.NewInt(3), big.NewInt(1), code.MULMOD)
	if !res {
		t.Error(err)
	}
	// test ((2**256-1)*2)%7=2 - shows that internal addition does not 256 bit truncate
	res, err = tertiaryIntOpTest(math.U256(big.NewInt(-1)), big.NewInt(2), big.NewInt(7), big.NewInt(2), code.MULMOD)
	if !res {
		t.Error(err)
	}
}

func TestExp(t *testing.T) {
	// test 3^2=9
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(2), big.NewInt(9), code.EXP)
	if !res {
		t.Error(err)
	}
	// test 2 exp 256 = 0 - test wrap
	res, err = binaryIntOpTest(big.NewInt(2), big.NewInt(256), big.NewInt(0), code.EXP)
	if !res {
		t.Error(err)
	}
}

func TestSignextend(t *testing.T) {
	// test
	res, err := binaryIntOpTest(big.NewInt(-1), big.NewInt(0), math.U256(big.NewInt(-1)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(1), big.NewInt(0), math.U256(big.NewInt(1)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(127), big.NewInt(0), math.U256(big.NewInt(127)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(128), big.NewInt(0), math.U256(big.NewInt(-128)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(254), big.NewInt(0), math.U256(big.NewInt(-2)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(257), big.NewInt(0), math.U256(big.NewInt(1)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(65534), big.NewInt(1), math.U256(big.NewInt(-2)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(65537), big.NewInt(1), math.U256(big.NewInt(1)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
	// test
	res, err = binaryIntOpTest(big.NewInt(65537), big.NewInt(2), math.U256(big.NewInt(65537)), code.SIGNEXTEND)
	if !res {
		t.Error(err)
	}
}

func TestLt(t *testing.T) {
	// test 3<9 res 1
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(1), code.LT)
	if !res {
		t.Error(err)
	}
	// test 9<3 res 0
	res, err = binaryIntOpTest(big.NewInt(9), big.NewInt(3), big.NewInt(0), code.LT)
	if !res {
		t.Error(err)
	}
	// test 3<3 res 0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.LT)
	if !res {
		t.Error(err)
	}
	// test 0xfffffffffffffffffffffffffffffffc((2**256)-4)<9 res 0
	res, err = binaryIntOpTest(math.U256(big.NewInt(-4)), big.NewInt(9), big.NewInt(0), code.LT)
	if !res {
		t.Error(err)
	}
	// test 9< tuple res 0
	res, err = binaryValueOpTest(value.NewInt64Value(9), value.NewEmptyTuple(), big.NewInt(0), code.LT)
	if res {
		t.Error("expected error")
	}
}

func TestGt(t *testing.T) {
	// test 3>9 res 0
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(0), code.GT)
	if !res {
		t.Error(err)
	}
	// test 9>3 res 1
	res, err = binaryIntOpTest(big.NewInt(9), big.NewInt(3), big.NewInt(1), code.GT)
	if !res {
		t.Error(err)
	}
	// test 3>3 res 0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.GT)
	if !res {
		t.Error(err)
	}
	// test 0xfffffffffffffffffffffffffffffffc(-4)>9 res 1
	res, err = binaryIntOpTest(math.U256(big.NewInt(-4)), big.NewInt(9), big.NewInt(1), code.GT)
	if !res {
		t.Error(err)
	}
}

func TestSlt(t *testing.T) {
	// test 3 < 9 = 1
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(1), code.SLT)
	if !res {
		t.Error(err)
	}
	// test 9 < 3 = 0
	res, err = binaryIntOpTest(big.NewInt(9), big.NewInt(3), big.NewInt(0), code.SLT)
	if !res {
		t.Error(err)
	}
	// test 3 < 3 = 0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.SLT)
	if !res {
		t.Error(err)
	}
	// test -3 < 3 = 1
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), big.NewInt(3), big.NewInt(1), code.SLT)
	if !res {
		t.Error(err)
	}
	// test -3 < -4 = 0
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), math.U256(big.NewInt(-4)), big.NewInt(0), code.SLT)
	if !res {
		t.Error(err)
	}
	// test -3 < -2 = 1
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), math.U256(big.NewInt(-2)), big.NewInt(1), code.SLT)
	if !res {
		t.Error(err)
	}
}

func TestSgt(t *testing.T) {
	// test 3 > 9 = 0
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(0), code.SGT)
	if !res {
		t.Error(err)
	}
	// test 9 > 3 = 1
	res, err = binaryIntOpTest(big.NewInt(9), big.NewInt(3), big.NewInt(1), code.SGT)
	if !res {
		t.Error(err)
	}
	// test 3 > 3 = 0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.SGT)
	if !res {
		t.Error(err)
	}
	// test -3 > 3 = 0
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), big.NewInt(3), big.NewInt(0), code.SGT)
	if !res {
		t.Error(err)
	}
	// test -3 > -4 = 1
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), math.U256(big.NewInt(-4)), big.NewInt(1), code.SGT)
	if !res {
		t.Error(err)
	}
	// test -3 > -2 = 0
	res, err = binaryIntOpTest(math.U256(big.NewInt(-3)), math.U256(big.NewInt(-2)), big.NewInt(0), code.SGT)
	if !res {
		t.Error(err)
	}
}

func TestEq(t *testing.T) {
	// test 3==9 = 0
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(0), code.EQ)
	if !res {
		t.Error(err)
	}
	// test 3==3 = 1
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(1), code.EQ)
	if !res {
		t.Error(err)
	}

	var vals [8]value.Value
	vals[0] = value.NewInt64Value(1)
	vals[1] = value.NewInt64Value(2)
	vals[2] = value.NewInt64Value(3)
	vals[3] = value.NewInt64Value(4)
	tup, _ := value.NewTupleOfSizeWithContents(vals, 4)

	vals[0] = value.NewInt64Value(1)
	vals[1] = value.NewInt64Value(2)
	vals[2] = value.NewInt64Value(3)
	vals[3] = value.NewInt64Value(4)
	tup2, _ := value.NewTupleOfSizeWithContents(vals, 4)

	vals[0] = value.NewInt64Value(1)
	vals[1] = value.NewInt64Value(2)
	vals[2] = value.NewInt64Value(7) // one value is different
	vals[3] = value.NewInt64Value(4)
	tup3, _ := value.NewTupleOfSizeWithContents(vals, 4)
	// test matching tuples
	res, err = binaryValueOpTest(tup, tup2, big.NewInt(1), code.EQ)
	if !res {
		t.Error(err)
	}
	// test different tuples
	res, err = binaryValueOpTest(tup, tup3, big.NewInt(0), code.EQ)
	if !res {
		t.Error(err)
	}
	// test different types
	res, err = binaryValueOpTest(tup, value.NewInt64Value(1), big.NewInt(0), code.EQ)
	if !res {
		t.Error("expected fail")
	}
}

func TestIszero(t *testing.T) {
	// test 0 isZero = 1
	res, err := unaryIntOpTest(big.NewInt(0), big.NewInt(1), code.ISZERO)
	if !res {
		t.Error(err)
	}
	// test 2 isZero = 0
	res, err = unaryIntOpTest(big.NewInt(3), big.NewInt(0), code.ISZERO)
	if !res {
		t.Error(err)
	}
}

func TestAnd(t *testing.T) {
	// test 0x03 and 0x09 = 0x01
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(1), code.AND)
	if !res {
		t.Error(err)
	}
	// test 0x03 and 0x03 = 0x03
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(3), code.AND)
	if !res {
		t.Error(err)
	}
}

func TestOr(t *testing.T) {
	// test 0x03 or 0x09 = 0x0b
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(11), code.OR)
	if !res {
		t.Error(err)
	}
	// test 0x03 or 0x03 = 0x03
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(3), code.OR)
	if !res {
		t.Error(err)
	}
}

func TestXor(t *testing.T) {
	// test 0x03 xor 0x09 = 0x0a
	res, err := binaryIntOpTest(big.NewInt(3), big.NewInt(9), big.NewInt(10), code.XOR)
	if !res {
		t.Error(err)
	}
	// test 0x03 xor 0x03 = 0x00
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.XOR)
	if !res {
		t.Error(err)
	}
}

func TestNot(t *testing.T) {
	// test !0x00 = 0xffffffffffffffffffffffffffffffff(-1)
	res, err := unaryIntOpTest(big.NewInt(0), math.U256(big.NewInt(-1)), code.NOT)
	if !res {
		t.Error(err)
	}
	// test !0x03 = 0xfffffffffffffffffffffffffffffffc(-4)
	res, err = unaryIntOpTest(big.NewInt(3), math.U256(big.NewInt(-4)), code.NOT)
	if !res {
		t.Error(err)
	}
	// test !0xfffffffffffffffffffffffffffffffc(-4) = 0x03(3)
	res, err = unaryIntOpTest(math.U256(big.NewInt(-4)), math.U256(big.NewInt(3)), code.NOT)
	if !res {
		t.Error(err)
	}
}

func TestByte(t *testing.T) {
	// test 31st byte of 16 = 16
	res, err := binaryIntOpTest(big.NewInt(16), big.NewInt(31), big.NewInt(16), code.BYTE)
	if !res {
		t.Error(err)
	}
	// test 3rd byte of 3 = 0
	res, err = binaryIntOpTest(big.NewInt(3), big.NewInt(3), big.NewInt(0), code.BYTE)
	if !res {
		t.Error(err)
	}
}

func TestSha3(t *testing.T) {
	// test
	hash, _ := new(big.Int).SetString("80084422859880547211683076133703299733277748156566366325829078699459944778998", 10)
	res, err := unaryIntOpTest(big.NewInt(1), hash, code.SHA3)
	if !res {
		t.Error(err)
	}
}

func TestPop(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}
	m := NewMachine(insns, value.NewInt64Value(1), false, 100)

	m.Stack().Push(value.NewInt64Value(1))
	a := m.Stack().Count()
	if a != 1 {
		tmp := "PUSH failed stack size = "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}

	if succeeded, reason := runInstNoFault(m, code.POP); !succeeded {
		t.Error(reason)
	}

	a = m.Stack().Count()
	if a != 0 {
		tmp := "POP stack size check failed"
		t.Error(tmp)
	}
}

func TestSpush(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)

	if succeeded, reason := runInstNoFault(m, code.SPUSH); !succeeded {
		t.Error(reason)
	}
	a := m.Stack().Count()
	if a != 1 {
		tmp := "SPUSH stack size check failed"
		t.Error(tmp)
	}
}

func TestRpush(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)

	if succeeded, reason := runInstNoFault(m, code.RPUSH); !succeeded {
		t.Error(reason)
	}
	a := m.Stack().Count()
	if a != 1 {
		tmp := "RPUSH stack size check failed"
		t.Error(tmp)
	}
}

func TestRset(t *testing.T) {
	//test:
	// 1. push value
	// 2. run RSET
	// 3. verify machines different
	// 4. push value to known
	// 5. run RSET on known
	// 6. verify machines match
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(5))

	if succeeded, reason := runInstNoFault(m, code.RSET); !succeeded {
		t.Error(reason)
	}
	a := m.Stack().Count()
	if a != 0 {
		t.Error("RSET stack size check failed")
	}
	if ok, _ := Equal(knownMachine, m); ok {
		t.Error("machines equal expected different")
	}

	knownMachine.Stack().Push(value.NewInt64Value(5))
	if succeeded, reason := runInstNoFault(knownMachine, code.RSET); !succeeded {
		t.Error(reason)
	}
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestInbox(t *testing.T) {
	//test:
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)
	knowninbox := protocol.NewEmptyInbox()

	var tok protocol.TokenType
	tok[0] = 15
	tok[20] = 1

	msg := protocol.NewMessage(
		value.NewInt64Value(1),
		tok,
		big.NewInt(3),
		value.NewInt64Value(7).ToBytes(),
	)
	m.SendOnchainMessage(msg)
	m.DeliverOnchainMessage()

	knowninbox.SendMessage(msg)
	knowninbox.DeliverMessages()

	NewMachineAssertionContext(m, [2]uint64{0, 10000})

	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		tok.ToIntValue(),
		value.NewInt64Value(3),
		value.NewInt64Value(4),
	})
	m.Stack().Push(tup)
	if succeeded, reason := runInstNoFault(m, code.INBOX); !succeeded {
		t.Error(reason)
	}
	knownMachine.Stack().Push(knowninbox.Receive())
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestJump(t *testing.T) {
	//test:
	insns := []value.Operation{
		value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(1)},
		value.ImmediateOperation{Op: code.ADD, Val: value.NewInt64Value(4)},
		value.ImmediateOperation{Op: code.SUB, Val: value.NewInt64Value(5)},
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// run NOP to push value 1
	if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
		t.Error(reason)
	}
	// push 2 to set jump point
	var nextHash [32]byte
	codept := value.CodePointValue{InsnNum: 2, Op: value.BasicOperation{Op: code.SUB}, NextHash: nextHash}
	m.Stack().Push(codept)
	// JUMP
	if succeeded, reason := runInstNoFault(m, code.JUMP); !succeeded {
		t.Error(reason)
	}
	// PC should now be 2 - immediate operation that pushes 5 and subtracts
	if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
		t.Error(reason)
	}
	// verify sub was executed
	knownMachine.Stack().Push(value.NewInt64Value(4))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestCJump(t *testing.T) {
	//test:
	insns := []value.Operation{
		value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(1)},
		value.ImmediateOperation{Op: code.ADD, Val: value.NewInt64Value(4)},
		value.ImmediateOperation{Op: code.SUB, Val: value.NewInt64Value(5)},
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		// run NOP to push value 1
		if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
			t.Error(reason)
		}
		// push 0 for conditional
		m.Stack().Push(value.NewInt64Value(0))
		// push 2 to set jump point
		codept := value.CodePointValue{InsnNum: 2, Op: value.BasicOperation{Op: code.SUB}, NextHash: [32]byte{}}
		m.Stack().Push(codept)
		// CJUMP
		if succeeded, reason := runInstNoFault(m, code.CJUMP); !succeeded {
			t.Error(reason)
		}
		// PC should now be 2 - immediate operation that pushes 5 and subtracts
		if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
			t.Error(reason)
		}
		// verify sub was executed
		knownMachine.Stack().Push(value.NewInt64Value(4))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		// repeat test with conditional set to 1
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)
		// run NOP to push value 1
		if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
			t.Error(reason)
		}
		// push 1 for conditional
		m.Stack().Push(value.NewInt64Value(1))
		// push 2 to set jump point
		codept := value.CodePointValue{InsnNum: 2, Op: value.BasicOperation{Op: code.SUB}, NextHash: [32]byte{}}
		m.Stack().Push(codept)
		// CJUMP
		if succeeded, reason := runInstNoFault(m, code.CJUMP); !succeeded {
			t.Error(reason)
		}
		// PC should now be 2 - immediate operation that pushes 5 and subtracts
		if succeeded, reason := runInstOpNoFault(m, m.GetOperation()); !succeeded {
			t.Error(reason)
		}
		// verify sub was executed
		knownMachine.Stack().Push(value.NewInt64Value(4))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestStackempty(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		if succeeded, reason := runInstNoFault(m, code.STACKEMPTY); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match one item value = 1
		knownMachine.Stack().Push(value.NewInt64Value(1))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		m.Stack().Push(value.NewInt64Value(1))
		knownMachine := m.Clone().(*Machine)

		if succeeded, reason := runInstNoFault(m, code.STACKEMPTY); !succeeded {
			t.Error(reason)
		}
		// push 1 as matching value and 0 to knownMachine as result of STACKEMPTY call
		knownMachine.Stack().Push(value.NewInt64Value(0))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestPcpush(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	if succeeded, reason := runInstNoFault(m, code.PCPUSH); !succeeded {
		t.Error(reason)
	}
	// stack should have one item - current codepoint
	a := m.Stack().Count()
	if a != 1 {
		tmp := "PCPUSH stack size check failed expected 3 found "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}
	// verify known and unknown match one item value = 1
	var nextHash [32]byte
	codept := value.CodePointValue{Op: value.BasicOperation{Op: code.HALT}, NextHash: nextHash}
	knownMachine.Stack().Push(codept)
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestAuxpush(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(4))
	if succeeded, reason := runInstNoFault(m, code.AUXPUSH); !succeeded {
		t.Error(reason)
	}
	// auxstack should have one item - value popped from stack
	a := m.AuxStack().Count()
	if a != 1 {
		tmp := "AUXPUSH stack size check failed expected 1 found "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}
	// stack should be empty
	a = m.Stack().Count()
	if a != 0 {
		tmp := "AUXPUSH stack size check failed expected 3 found "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}
	// verify known and unknown match one item value = 4
	knownMachine.AuxStack().Push(value.NewInt64Value(4))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestAuxpop(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.AuxStack().Push(value.NewInt64Value(5))
	if succeeded, reason := runInstNoFault(m, code.AUXPOP); !succeeded {
		t.Error(reason)
	}
	// auxstack should be empty
	a := m.AuxStack().Count()
	if a != 0 {
		tmp := "AUXPOP stack size check failed expected 1 found "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}
	// stack should have one item - value popped from auxstack
	a = m.Stack().Count()
	if a != 1 {
		tmp := "AUXPOP stack size check failed expected 3 found "
		tmp += strconv.FormatInt(a, 10)
		t.Error(tmp)
	}
	// verify known and unknown match one item value = 5
	knownMachine.Stack().Push(value.NewInt64Value(5))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestAuxstackempty(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		// auxstack should be empty
		a := m.AuxStack().Count()
		if a != 0 {
			t.Errorf("AUXPOP stack size check failed expected 1 found %v", a)
		}
		// check aux stack empty and push results on data stack
		if succeeded, reason := runInstNoFault(m, code.AUXSTACKEMPTY); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match one item value = 1
		knownMachine.Stack().Push(value.NewInt64Value(1))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)
		m.AuxStack().Push(value.NewInt64Value(5))
		// auxstack should not be empty
		a := m.AuxStack().Count()
		if a != 1 {
			t.Errorf("AUXSTACKEMPTY stack size check failed expected 3 found %v", a)
		}
		if succeeded, reason := runInstNoFault(m, code.AUXSTACKEMPTY); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match
		knownMachine.AuxStack().Push(value.NewInt64Value(5))
		knownMachine.Stack().Push(value.NewInt64Value(0))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestNop(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}
	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		// verify known and unknown match
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
		// check NOP does nothing
		if succeeded, reason := runInstNoFault(m, code.NOP); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		// check NOP does nothing
		// immediate operation pushes value then does nothing
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)
		if succeeded, reason := runInstOpNoFault(m, value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(1)}); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match
		knownMachine.Stack().Push(value.NewInt64Value(1))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

}

func TestErrpush(t *testing.T) {
	// test
	insns := []value.Operation{
		value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(1)},
		value.ImmediateOperation{Op: code.ADD, Val: value.NewInt64Value(4)},
		value.ImmediateOperation{Op: code.SUB, Val: value.NewInt64Value(5)},
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// push codepoint onto stack
	var nextHash [32]byte
	codept := value.CodePointValue{InsnNum: 4, Op: value.BasicOperation{Op: code.HALT}, NextHash: nextHash}
	m.Stack().Push(codept)
	knownMachine.Stack().Push(codept)
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// run errset to set the error handler
	if succeeded, reason := runInstNoFault(m, code.ERRSET); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown different
	if ok, _ := Equal(knownMachine, m); ok {
		tmp := "machines equal expected different"
		t.Error(tmp)
	}
	// set known to match
	if succeeded, reason := runInstNoFault(knownMachine, code.ERRSET); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// run errpush to push error handler to data stack
	if succeeded, reason := runInstNoFault(m, code.ERRPUSH); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown different
	if ok, _ := Equal(knownMachine, m); ok {
		tmp := "machines equal expected different"
		t.Error(tmp)
	}
	// push error handler code point to known data stack
	knownMachine.Stack().Push(codept)
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestErrset(t *testing.T) {
	// test
	insns := []value.Operation{
		value.ImmediateOperation{Op: code.NOP, Val: value.NewInt64Value(1)},
		value.ImmediateOperation{Op: code.ADD, Val: value.NewInt64Value(4)},
		value.ImmediateOperation{Op: code.SUB, Val: value.NewInt64Value(5)},
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// push codepoint onto stack
	codept := value.CodePointValue{InsnNum: 4, Op: value.BasicOperation{Op: code.HALT}, NextHash: [32]byte{}}
	m.Stack().Push(codept)
	knownMachine.Stack().Push(codept)
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// run errset to set the error handler
	if succeeded, reason := runInstNoFault(m, code.ERRSET); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown different
	if ok, _ := Equal(knownMachine, m); ok {
		tmp := "machines equal expected different"
		t.Error(tmp)
	}
	// set known to match
	if succeeded, reason := runInstNoFault(knownMachine, code.ERRSET); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestError(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// check NOP does nothing
	if failed, reason := runInstWithError(m, code.ERROR); !failed {
		// ERROR failed - should have generated error
		t.Error(reason)
	}
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestDup0(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(1))
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if succeeded, reason := runInstNoFault(m, code.DUP0); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestDup1(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(1))
	m.Stack().Push(value.NewInt64Value(2))
	knownMachine.Stack().Push(value.NewInt64Value(1))
	knownMachine.Stack().Push(value.NewInt64Value(2))
	if succeeded, reason := runInstNoFault(m, code.DUP1); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestDup2(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(1))
	m.Stack().Push(value.NewInt64Value(2))
	m.Stack().Push(value.NewInt64Value(3))
	knownMachine.Stack().Push(value.NewInt64Value(1))
	knownMachine.Stack().Push(value.NewInt64Value(2))
	knownMachine.Stack().Push(value.NewInt64Value(3))
	if succeeded, reason := runInstNoFault(m, code.DUP2); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestSwap2(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewInt64Value(1))
	m.Stack().Push(value.NewInt64Value(2))
	m.Stack().Push(value.NewInt64Value(3))
	if succeeded, reason := runInstNoFault(m, code.SWAP2); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(3))
	knownMachine.Stack().Push(value.NewInt64Value(2))
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestTget(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		tup := value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2))

		m.Stack().Push(tup)
		m.Stack().Push(value.NewInt64Value(1))
		if succeeded, reason := runInstNoFault(m, code.TGET); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match one item value = 1
		knownMachine.Stack().Push(value.NewInt64Value(2))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		// test with only int on stack
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)
		m.Stack().Push(value.NewInt64Value(1))
		if failed, reason := runInstWithError(m, code.TGET); !failed {
			t.Error(reason)
		}
		// verify known and unknown match expect empty stack
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		// test A out of range
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)
		m.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)))
		m.Stack().Push(value.NewInt64Value(3))
		if failed, reason := runInstWithError(m, code.TGET); !failed {
			t.Error(reason)
		}
		// verify known and unknown match expect empty stack
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestTset(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		m.Stack().Push(value.NewInt64Value(3))
		m.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)))
		m.Stack().Push(value.NewInt64Value(1))
		if succeeded, reason := runInstNoFault(m, code.TSET); !succeeded {
			t.Error(reason)
		}
		// verify known and unknown match
		knownMachine.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(3)))
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		m.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)))
		// test with only tuple on stack
		if failed, reason := runInstWithError(m, code.TSET); !failed {
			t.Error(reason)
		}
		// verify known and unknown match expect empty stack
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}

	{
		// test incorrect A
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		m.Stack().Push(value.NewInt64Value(3))
		m.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)))
		m.Stack().Push(value.NewInt64Value(4))
		if failed, reason := runInstWithError(m, code.TSET); !failed {
			t.Error(reason)
		}
		// verify known and unknown match
		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestTlen(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.Stack().Push(value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)))
	if succeeded, reason := runInstNoFault(m, code.TLEN); !succeeded {
		t.Error(reason)
	}
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(2))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// test A not a tuple
	if failed, reason := runInstWithError(m, code.TLEN); !failed {
		t.Error(reason)
	}
	// verify known and unknown match expect empty stack
	_, err := knownMachine.Stack().Pop()
	if err != nil {
		t.Error(err)
	}
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestType(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	testValues := []value.Value{
		value.NewEmptyTuple(),
		value.NewTuple2(value.NewInt64Value(1), value.NewInt64Value(2)),
		value.ErrorCodePoint,
		value.NewInt64Value(100),
	}

	resultValues := []value.Value{
		value.NewInt64Value(3),
		value.NewInt64Value(3),
		value.NewInt64Value(1),
		value.NewInt64Value(0),
	}

	for i := range testValues {
		m := NewMachine(insns, value.NewInt64Value(1), false, 100)
		knownMachine := m.Clone().(*Machine)

		m.Stack().Push(testValues[i])
		if succeeded, reason := runInstNoFault(m, code.TYPE); !succeeded {
			t.Error(reason)
		}

		knownMachine.Stack().Push(resultValues[i])

		if ok, err := Equal(knownMachine, m); !ok {
			t.Error(err)
		}
	}
}

func TestBreakpoint(t *testing.T) {
	// test
	insns := []value.Operation{value.BasicOperation{Op: code.NOP}, value.BasicOperation{Op: code.HALT}}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	_, blocked := RunInstruction(m, value.BasicOperation{Op: code.BREAKPOINT})
	if _, ok := blocked.(machine.BreakpointBlocked); !ok {
		t.Error("Failed to produce breakpoint block")
	}
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}

func TestLog(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.LOG},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)
	m.Stack().Push(value.NewInt64Value(5))
	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// verify out message
	logs := ad.Logs
	if len(logs) != 1 {
		t.Error("No log value generated")
	}
	if !logs[0].Equal(value.NewInt64Value(5)) {
		t.Error("log value incorrect")
	}
}

func TestSendFungible(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.SEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// fungible value=10
	var tok protocol.TokenType
	tok[0] = 15
	tok[20] = 0
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(7),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(10), [32]byte{}))

	// send token 15 value=7 to dest 4
	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	// verify out message
	if len(ad.OutMsgs) != 1 {
		t.Error("No out message generated")
	}

	dest := [32]byte{}
	dest[31] = 4
	knownmessage := protocol.NewMessage(value.NewInt64Value(1), tok, big.NewInt(7), dest)
	if !ad.OutMsgs[0].Equals(knownmessage) {
		t.Error("Out message incorrect")
	}
}

func TestSendNonFungible(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.SEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// test send of non fungible
	var tok protocol.TokenType
	tok[0] = 16
	tok[20] = 1
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(7),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(7), [32]byte{}))

	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 1 {
		t.Error("No out message generated")
	}

	dest := [32]byte{}
	dest[31] = 4
	knownmessage := protocol.NewMessage(value.NewInt64Value(1), tok, big.NewInt(7), dest)
	if !msgs[0].Equals(knownmessage) {
		t.Error("Out message incorrect")
	}
}

func TestSendLowBalance(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.SEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// test send with insufficient funds
	var tok protocol.TokenType
	tok[0] = 17
	tok[20] = 0
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(17),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(10), [32]byte{}))

	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	knownMachine.Stack().Push(tup)
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 0 {
		t.Error("No out message generated")
	}
}

func TestNbsend1(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.NBSEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	var tok protocol.TokenType
	tok[0] = 15
	tok[20] = 1
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(10),
		tok.ToIntValue(),
	})

	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(10), [32]byte{}))

	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))

	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}

	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 1 {
		t.Error("No out message generated")
	}
}

func TestNBSendFungible(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.NBSEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// fungible value=10
	var tok protocol.TokenType
	tok[0] = 15
	tok[20] = 0
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(7),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(10), [32]byte{}))

	// send token 15 value=7 to dest 4
	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 1 {
		t.Error("No out message generated")
	}

	dest := [32]byte{}
	dest[31] = 4
	knownmessage := protocol.NewMessage(value.NewInt64Value(1), tok, big.NewInt(7), dest)
	if !msgs[0].Equals(knownmessage) {
		t.Error("Out message incorrect")
	}
}

func TestNBSendNonFungible(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.NBSEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// test send of non fungible
	var tok protocol.TokenType
	tok[0] = 16
	tok[20] = 1
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(7),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(7), [32]byte{}))

	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(1))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 1 {
		t.Error("No out message generated")
	}

	dest := [32]byte{}
	dest[31] = 4
	knownmessage := protocol.NewMessage(value.NewInt64Value(1), tok, big.NewInt(7), dest)
	if !msgs[0].Equals(knownmessage) {
		t.Error("Out message incorrect")
	}
}

func TestNBSendLowBalance(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.NBSEND},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	// test send with insufficient funds
	var tok protocol.TokenType
	tok[0] = 17
	tok[20] = 0
	tup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewInt64Value(1),
		value.NewInt64Value(4),
		value.NewInt64Value(17),
		tok.ToIntValue(),
	})
	m.Stack().Push(tup)

	// add tokens to balanceTracker
	m.SendOnchainMessage(protocol.NewMessage(value.NewEmptyTuple(), tok, big.NewInt(10), [32]byte{}))

	ad := m.ExecuteAssertion(10, protocol.NewTimeBounds(0, 1000))
	// verify known and unknown match
	knownMachine.Stack().Push(value.NewInt64Value(0))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
	msgs := ad.OutMsgs
	// verify out message
	if len(msgs) != 0 {
		t.Error("No out message generated")
	}
}

func TestGettime(t *testing.T) {
	// test
	insns := []value.Operation{
		value.BasicOperation{Op: code.GETTIME},
		value.BasicOperation{Op: code.HALT},
	}

	m := NewMachine(insns, value.NewInt64Value(1), false, 100)
	knownMachine := m.Clone().(*Machine)

	m.ExecuteAssertion(10, [2]uint64{5, 10})

	// verify known and unknown match
	knownMachine.Stack().Push(value.NewTuple2(value.NewInt64Value(5), value.NewInt64Value(10)))
	if ok, err := Equal(knownMachine, m); !ok {
		t.Error(err)
	}
}
