# Copyright 2019, Offchain Labs, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import arbitrum as arb

count = 0


def runUnaryOp(vm, arg1, op):
    global count
    vm.push(arg1)
    op()


def runBinaryOp(vm, arg1, arg2, op):
    global count
    vm.push(arg2)
    vm.push(arg1)
    op()


def runTertiaryOp(vm, arg1, arg2, arg3, op):
    global count
    vm.push(arg3)
    vm.push(arg2)
    vm.push(arg1)
    op()


def testUnaryOp(vm, arg1, res, op):
    runUnaryOp(vm, arg1, op)
    cmpEqual(vm, res)


def testBinaryOp(vm, arg1, arg2, res, op):
    runBinaryOp(vm, arg1, arg2, op)
    cmpEqual(vm, res)


def testTertiaryOp(vm, arg1, arg2, arg3, res, op):
    runTertiaryOp(vm, arg1, arg2, arg3, op)
    cmpEqual(vm, res)


def cmpEqual(vm, res):
    global count
    vm.push(res)
    vm.eq()
    vm.push(arb.ast.AVMLabel("next" + str(count)))
    vm.cjump()
    vm.error()
    vm.set_label(arb.ast.AVMLabel("next" + str(count)))
    count += 1


def cmpNotEqual(vm, res):
    global count
    vm.push(res)
    vm.eq()
    vm.iszero()
    vm.push(arb.ast.AVMLabel("next" + str(count)))
    vm.cjump()
    vm.error()
    vm.set_label(arb.ast.AVMLabel("next" + str(count)))
    count += 1


def test(vm):
    # uncomment push, jump and set_label and move set_label if we want to skip some tests
    vm.push(arb.ast.AVMLabel("jump_to_test"))
    vm.jump()
    # ADD
    testBinaryOp(vm, 4, 3, 7, vm.add)
    #    testBinaryOp(vm,4,3,6,vm.add)
    testBinaryOp(vm, 0, 0, 0, vm.add)
    testBinaryOp(vm, 2 ** 256 - 1, 4, 3, vm.add)
    testBinaryOp(vm, 2 ** 256 - 2, 1, 2 ** 256 - 1, vm.add)
    # MUL
    testBinaryOp(vm, 4, 3, 12, vm.mul)
    testBinaryOp(vm, 3, 0, 0, vm.mul)
    testBinaryOp(vm, 2 ** 256 - 1, 1, 2 ** 256 - 1, vm.mul)
    testBinaryOp(vm, 2 ** 256 - 2, 1, 2 ** 256 - 2, vm.mul)
    # SUB
    testBinaryOp(vm, 4, 3, 1, vm.sub)
    testBinaryOp(vm, 3, 4, 2 ** 256 - 1, vm.sub)
    # DIV
    testBinaryOp(vm, 12, 3, 4, vm.div)
    runBinaryOp(vm, 2 ** 256 - 6, 3, vm.div)
    cmpNotEqual(vm, 4)
    # divide by 0
    vm.push(arb.ast.AVMLabel("DIV_divide_by_0_expected"))
    vm.errset()
    runBinaryOp(vm, 12, 0, vm.div)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("DIV_divide_by_0_expected"))
    # SDIV
    testBinaryOp(vm, 12, 3, 4, vm.sdiv)
    testBinaryOp(vm, 12, 2 ** 256 - 3, 2 ** 256 - 4, vm.sdiv)
    testBinaryOp(vm, 2 ** 256 - 12, 3, 2 ** 256 - 4, vm.sdiv)
    testBinaryOp(vm, 2 ** 256 - 12, 2 ** 256 - 3, 4, vm.sdiv)
    # sdivide by 0
    vm.push(arb.ast.AVMLabel("SDIV_divide_by_0_expected"))
    vm.errset()
    runBinaryOp(vm, 3, 0, vm.sdiv)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("SDIV_divide_by_0_expected"))
    # MOD
    testBinaryOp(vm, 8, 3, 2, vm.mod)
    testBinaryOp(vm, 8, 2 ** 256 - 3, 8, vm.mod)
    testBinaryOp(vm, 0, 3, 0, vm.mod)
    # mod by 0
    vm.push(arb.ast.AVMLabel("MOD_by_0_expected"))
    vm.errset()
    runBinaryOp(vm, 3, 0, vm.mod)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("MOD_by_0_expected"))
    # SMOD
    testBinaryOp(vm, 8, 3, 2, vm.smod)
    testBinaryOp(vm, 8, 2 ** 256 - 3, 2, vm.smod)
    testBinaryOp(vm, 2 ** 256 - 8, 3, 2 ** 256 - 2, vm.smod)
    testBinaryOp(vm, 2 ** 256 - 8, 2 ** 256 - 3, 2 ** 256 - 2, vm.smod)
    # smod by 0
    vm.push(arb.ast.AVMLabel("SMOD_by_0_expected"))
    vm.errset()
    runBinaryOp(vm, 3, 0, vm.smod)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("SMOD_by_0_expected"))
    # ADDMOD
    testTertiaryOp(vm, 8, 5, 3, 1, vm.addmod)
    testTertiaryOp(vm, 2 ** 256 - 1, 1, 7, 2, vm.addmod)
    testTertiaryOp(vm, 0, 0, 7, 0, vm.addmod)
    # addmod by 0
    vm.push(arb.ast.AVMLabel("ADDMOD_by_0_expected"))
    vm.errset()
    runTertiaryOp(vm, 8, 3, 0, vm.addmod)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("ADDMOD_by_0_expected"))
    # MULMOD
    testTertiaryOp(vm, 8, 2, 3, 1, vm.mulmod)
    testTertiaryOp(vm, 2 ** 256 - 1, 2, 7, 2, vm.mulmod)
    testTertiaryOp(vm, 0, 0, 7, 0, vm.mulmod)
    # mulmod by 0
    vm.push(arb.ast.AVMLabel("MULMOD_by_0_expected"))
    vm.errset()
    runTertiaryOp(vm, 8, 3, 0, vm.mulmod)
    vm.error()
    vm.set_label(arb.ast.AVMLabel("MULMOD_by_0_expected"))
    # EXP
    testBinaryOp(vm, 3, 2, 9, vm.exp)
    testBinaryOp(vm, 2, 256, 0, vm.exp)
    # LT
    testBinaryOp(vm, 3, 9, 1, vm.lt)
    testBinaryOp(vm, 9, 3, 0, vm.lt)
    testBinaryOp(vm, 3, 3, 0, vm.lt)
    testBinaryOp(vm, 2 ** 256 - 3, 9, 0, vm.lt)
    # GT
    testBinaryOp(vm, 3, 9, 0, vm.gt)
    testBinaryOp(vm, 9, 3, 1, vm.gt)
    testBinaryOp(vm, 3, 3, 0, vm.gt)
    testBinaryOp(vm, 2 ** 256 - 3, 9, 1, vm.gt)
    # SLT
    testBinaryOp(vm, 7, 3, 0, vm.slt)
    testBinaryOp(vm, 3, 7, 1, vm.slt)
    testBinaryOp(vm, 2 ** 256 - 3, 2 ** 256 - 7, 0, vm.slt)
    testBinaryOp(vm, 2 ** 256 - 7, 2 ** 256 - 3, 1, vm.slt)
    testBinaryOp(vm, 3, 2 ** 256 - 7, 0, vm.slt)
    testBinaryOp(vm, 2 ** 256 - 3, 7, 1, vm.slt)
    testBinaryOp(vm, 3, 3, 0, vm.slt)
    # SGT
    testBinaryOp(vm, 7, 3, 1, vm.sgt)
    testBinaryOp(vm, 3, 7, 0, vm.sgt)
    testBinaryOp(vm, 2 ** 256 - 3, 2 ** 256 - 7, 1, vm.sgt)
    testBinaryOp(vm, 2 ** 256 - 7, 2 ** 256 - 3, 0, vm.sgt)
    testBinaryOp(vm, 3, 2 ** 256 - 7, 1, vm.sgt)
    testBinaryOp(vm, 2 ** 256 - 7, 3, 0, vm.sgt)
    testBinaryOp(vm, 3, 3, 0, vm.sgt)
    # EQ
    testBinaryOp(vm, 7, 3, 0, vm.eq)
    testBinaryOp(vm, 3, 3, 1, vm.eq)
    testBinaryOp(
        vm,
        arb.std.bigtuple.fromints([1, 2]),
        arb.std.bigtuple.fromints([1, 2]),
        1,
        vm.eq,
    )
    testBinaryOp(
        vm,
        arb.std.bigtuple.fromints([1, 2]),
        arb.std.bigtuple.fromints([1, 3]),
        0,
        vm.eq,
    )
    # ISZERO
    testUnaryOp(vm, 0, 1, vm.iszero)
    testUnaryOp(vm, 2, 0, vm.iszero)
    # AND
    testBinaryOp(vm, 3, 9, 1, vm.bitwise_and)
    testBinaryOp(vm, 3, 3, 3, vm.bitwise_and)
    # OR
    testBinaryOp(vm, 3, 9, 11, vm.bitwise_or)
    testBinaryOp(vm, 3, 3, 3, vm.bitwise_or)
    # XOR
    testBinaryOp(vm, 3, 9, 10, vm.bitwise_xor)
    testBinaryOp(vm, 3, 3, 0, vm.bitwise_xor)
    # NOT
    testUnaryOp(vm, 0, 2 ** 256 - 1, vm.bitwise_not)
    testUnaryOp(vm, 3, 2 ** 256 - 4, vm.bitwise_not)
    testUnaryOp(vm, 2 ** 256 - 4, 3, vm.bitwise_not)
    # BYTE
    testBinaryOp(vm, 16, 31, 16, vm.byte)
    testBinaryOp(vm, 3, 3, 0, vm.byte)
    # SIGNEXTEND
    testBinaryOp(vm, 2 ** 256 - 1, 0, 2 ** 256 - 1, vm.signextend)
    testBinaryOp(vm, 1, 0, 1, vm.signextend)
    testBinaryOp(vm, 127, 0, 127, vm.signextend)
    testBinaryOp(vm, 128, 0, 2 ** 256 - 128, vm.signextend)
    testBinaryOp(vm, 254, 0, 2 ** 256 - 2, vm.signextend)
    testBinaryOp(vm, 257, 0, 1, vm.signextend)
    testBinaryOp(vm, 65534, 1, 2 ** 256 - 2, vm.signextend)
    testBinaryOp(vm, 65537, 1, 1, vm.signextend)
    testBinaryOp(vm, 65537, 2, 65537, vm.signextend)
    # HASH
    testUnaryOp(
        vm,
        10,
        int("c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8", 16),
        vm.hash,
    )
    # TYPE
    testUnaryOp(vm, 3, 0, vm.type)
    testUnaryOp(vm, arb.value.AVMCodePoint(0, 0, b"\0" * 32), 1, vm.type)
    testUnaryOp(vm, arb.std.bigtuple.fromints([1, 2]), 3, vm.type)
    # POP
    vm.push(3)
    vm.push(5)
    vm.pop()
    cmpEqual(vm, 3)
    # SPUSH
    #    vm.static = 4
    #    vm.spush()
    #    cmpEqual(vm, 4)
    # RPUSH
    #    vm.register = 5
    #    vm.rpush()
    #    cmpEqual(vm, 5)
    # RSET
    vm.push(7)
    vm.rset()
    # JUMP
    vm.push(arb.ast.AVMLabel("jump_opcode_test"))
    vm.jump()
    vm.nop()
    vm.set_label(arb.ast.AVMLabel("jump_opcode_test"))
    # CJUMP
    vm.push(3)
    vm.push(0)
    vm.push(arb.ast.AVMLabel("cjump_opcode_test1"))
    vm.cjump()
    vm.push(4)
    vm.set_label(arb.ast.AVMLabel("cjump_opcode_test1"))
    cmpEqual(vm, 4)
    vm.pop()
    # CJUMP false
    vm.push(3)
    vm.push(1)
    vm.push(arb.ast.AVMLabel("cjump_opcode_test2"))
    vm.cjump()
    vm.push(4)
    vm.set_label(arb.ast.AVMLabel("cjump_opcode_test2"))
    cmpEqual(vm, 3)
    # STACKEMPTY
    vm.stackempty()
    cmpEqual(vm, 1)
    vm.push(3)
    vm.stackempty()
    cmpEqual(vm, 0)
    vm.pop()
    # PCPUSH
    # vm.set_label(arb.ast.AVMLabel("pcpush_opcode_test"))
    # vm.pcpush()
    # cmpEqual(vm, arb.ast.AVMLabel("pcpush_opcode_test"))
    # AUXPUSH/AUXPOP/AUXSTACKEMPTY
    vm.push(5)
    vm.auxpush()
    vm.auxstackempty()
    cmpEqual(vm, 0)
    vm.auxpop()
    vm.auxstackempty()
    cmpEqual(vm, 1)
    vm.pop()
    # NOP
    vm.nop()
    # ERRPUSH
    # ERRSET
    # DUP0
    vm.set_label(arb.ast.AVMLabel("jump_to_test"))
    vm.push(6)
    vm.dup0()
    cmpEqual(vm, 6)
    cmpEqual(vm, 6)
    # DUP1
    vm.push(6)
    vm.push(7)
    vm.dup1()
    cmpEqual(vm, 6)
    cmpEqual(vm, 7)
    cmpEqual(vm, 6)
    # DUP2
    vm.set_label(arb.ast.AVMLabel("jump_to_test"))
    vm.push(6)
    vm.push(7)
    vm.push(8)
    vm.dup2()
    cmpEqual(vm, 6)
    cmpEqual(vm, 8)
    cmpEqual(vm, 7)
    cmpEqual(vm, 6)
    # SWAP1
    vm.push(6)
    vm.push(7)
    vm.swap1()
    cmpEqual(vm, 6)
    cmpEqual(vm, 7)
    # SWAP2
    vm.push(6)
    vm.push(7)
    vm.push(8)
    vm.swap1()
    cmpEqual(vm, 6)
    cmpEqual(vm, 7)
    cmpEqual(vm, 8)
    # TGET
    # TSET
    # TLEN
    # BREAKPOINT
    # LOG
    # SEND
    # NBSEND
    # GETTIME
    # INBOX
    # ERROR
    # HALT
    #
    vm.nop()
    vm.nop()
    vm.nop()
    vm.halt()


code = arb.compile_block(test)
vm = arb.compile_program(arb.ast.BlockStatement([]), code)
# print(vm.code)
with open("../arb-validator/test/opcodetest.ao", "wb") as f:
    arb.marshall.marshall_vm(vm, f)
