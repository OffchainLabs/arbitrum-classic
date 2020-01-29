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


def ntimes(vm, times, op):
    ret = []
    for i in range(times):
        ret.append(op(vm))
    return ret


def speedtestUnaryOp(vm, arg, op):
    vm.push(arg)
    vm.while_loop(lambda v: [v.push(1)], lambda v: ntimes(v, 1000, op))


def speedtestUnaryOp_lis(vm, argLis, op):
    if argLis == []:
        vm.while_loop(lambda v: [v.push(1)], lambda v: ntimes(v, 1000, op))
    else:
        vm.push(argLis[-1])
        speedtestUnaryOp_lis(vm, argLis[:-1], op)


def speedtestBinaryOp_Dup(vm, arg, op):
    speedtestUnaryOp(vm, arg, lambda v: [v.dup0(), op(v)])


def speedtestBinaryOp_Pushes(vm, arg1, arg2, op):
    speedtestUnaryOp(vm, 0, lambda v: [v.push(arg2), v.push(arg1), op(v), v.pop()])


def speedtestTernaryOp_Pushes(vm, arg1, arg2, arg3, op):
    speedtestUnaryOp(
        vm, 0, lambda v: [v.push(arg3), v.push(arg2), v.push(arg1), op(v), v.pop()]
    )


def makeAoFile(func, filepath):
    code = arb.compile_block(func)
    vm = arb.compile_program(arb.ast.BlockStatement([]), code)
    vm.static = 4
    with open(filepath, "wb") as f:
        arb.marshall.marshall_vm(vm, f)


aos = (
    ("hash_0_0", lambda vm: speedtestUnaryOp(vm, 0, lambda v: v.hash())),
    (
        "push_pop_0_0",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.push(0), v.pop()]),
    ),
    ("add_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 992, lambda v: v.add())),
    ("mul_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 992, lambda v: v.mul())),
    ("div_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 9, lambda v: v.div())),
    ("sdiv_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 9, lambda v: v.sdiv())),
    ("mod_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 9, lambda v: v.mod())),
    ("smod_2_1", lambda vm: speedtestBinaryOp_Pushes(vm, 571, 9, lambda v: v.smod())),
    (
        "addmod_3_1",
        lambda vm: speedtestTernaryOp_Pushes(
            vm, 57112352, 8386423523, 6312, lambda v: v.addmod()
        ),
    ),
    (
        "mulmod_3_1",
        lambda vm: speedtestTernaryOp_Pushes(
            vm, 57112352, 8386423523, 6312, lambda v: v.mulmod()
        ),
    ),
    (
        "exp_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.exp()
        ),
    ),
    (
        "lt_2_1",
        lambda vm: speedtestBinaryOp_Pushes(vm, 57112352, 8386423523, lambda v: v.lt()),
    ),
    (
        "gt_2_1",
        lambda vm: speedtestBinaryOp_Pushes(vm, 57112352, 8386423523, lambda v: v.gt()),
    ),
    (
        "slt_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.slt()
        ),
    ),
    (
        "sgt_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.sgt()
        ),
    ),
    (
        "eq_2_1",
        lambda vm: speedtestBinaryOp_Pushes(vm, 57112352, 8386423523, lambda v: v.eq()),
    ),
    ("iszero_0_0", lambda vm: speedtestUnaryOp(vm, 0, lambda v: v.iszero())),
    (
        "and_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.bitwise_and()
        ),
    ),
    (
        "or_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.bitwise_or()
        ),
    ),
    (
        "xor_2_1",
        lambda vm: speedtestBinaryOp_Pushes(
            vm, 57112352, 8386423523, lambda v: v.bitwise_xor()
        ),
    ),
    ("not_0_0", lambda vm: speedtestUnaryOp(vm, 57112352, lambda v: v.bitwise_not())),
    (
        "byte_2_1",
        lambda vm: speedtestBinaryOp_Pushes(vm, 57112352, 5, lambda v: v.byte()),
    ),
    (
        "signextend_2_1",
        lambda vm: speedtestBinaryOp_Pushes(vm, 57112352, 5, lambda v: v.signextend()),
    ),
    (
        "stackempty_pop_stackempty_0_0",
        lambda vm: speedtestUnaryOp(
            vm, 1, lambda v: [v.stackempty(), v.pop(), v.stackempty()]
        ),
    ),
    ("pcpush_0_1", lambda vm: speedtestUnaryOp(vm, 1, lambda v: [v.pcpush(), v.pop()])),
    (
        "auxpush_auxpop_0_0",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.auxpush(), v.auxpop()]),
    ),
    ("nop_0_0", lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.nop()])),
    (
        "errpush_0_1",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.errpush(), v.pop()]),
    ),
    (
        "pcpush_errset_0_0",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.pcpush(), v.errset()]),
    ),
    ("dup0_0_1", lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.dup0(), v.pop()])),
    (
        "dup1_0_1",
        lambda vm: speedtestUnaryOp_lis(vm, [0, 1], lambda v: [v.dup1(), v.pop()]),
    ),
    (
        "dup2_0_1",
        lambda vm: speedtestUnaryOp_lis(vm, [0, 1, 2], lambda v: [v.dup2(), v.pop()]),
    ),
    ("swap1_0_0", lambda vm: speedtestUnaryOp_lis(vm, [0, 1], lambda v: [v.swap1()])),
    (
        "swap2_0_0",
        lambda vm: speedtestUnaryOp_lis(vm, [0, 1, 2], lambda v: [v.swap2()]),
    ),
    (
        "dup0_tlen_0_1",
        lambda vm: speedtestUnaryOp(
            vm, arb.value.Tuple([1, 2, 3, 4]), lambda v: [v.dup0(), v.tlen(), v.pop()]
        ),
    ),
    (
        "dup1_dup1_tget_0_1",
        lambda vm: speedtestUnaryOp_lis(
            vm,
            [2, arb.value.Tuple([1, 2, 3, 4])],
            lambda v: [v.dup1(), v.dup1(), v.tget(), v.pop()],
        ),
    ),
    (
        "dup2_dup2_dup2_tset_0_1",
        lambda vm: speedtestUnaryOp_lis(
            vm,
            [2, arb.value.Tuple([1, 2, 3, 4]), 5],
            lambda v: [v.dup2(), v.dup2(), v.dup2(), v.tset(), v.pop()],
        ),
    ),
    (
        "dup2_dup2_dup2_tset_hash_0_1",
        lambda vm: speedtestUnaryOp_lis(
            vm,
            [2, arb.value.Tuple([1, 2, 3, 4]), 5],
            lambda v: [v.dup2(), v.dup2(), v.dup2(), v.tset(), v.hash(), v.pop()],
        ),
    ),
    (
        "tset_push_tset_hash_4_1",
        lambda vm: speedtestUnaryOp(
            vm,
            0,
            lambda v: [
                v.push(6),
                v.push(5),
                v.push(arb.value.Tuple([1, 2, 3, 4])),
                v.push(1),
                v.tset(),
                v.push(2),
                v.tset(),
                v.hash(),
                v.pop(),
            ],
        ),
    ),
    (
        "dup0_ethhash2_0_0",
        lambda vm: speedtestUnaryOp(vm, 2, lambda v: [v.dup0(), v.ethhash2()]),
    ),
    (
        "gettime_0_1",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.gettime(), v.pop()]),
    ),
    (
        "inbox_1_1",
        lambda vm: speedtestUnaryOp(vm, 0, lambda v: [v.push(0), v.inbox(), v.pop()]),
    ),
)

for ao in aos:
    makeAoFile(ao[1], "../arb-avm-cpp/speedtest/aos/" + ao[0] + ".ao")

# code = arb.compile_block(test_arithmetic)
# vm = arb.compile_program(arb.ast.BlockStatement([]), code)
# vm.static = 4
# print("math ", len(vm.code), " codepoints")
# print(vm.code)
# with open("../arb-validator/test/opcodetestmath.ao", "wb") as f:
#    arb.marshall.marshall_vm(vm, f)
