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

import eth_utils
from . import byterange, sized_common, bytestack, tup, stack_int, bigtuple_int
from ..annotation import modifies_stack
from ..vm import VM
from .. import value
from .struct import Struct

sized_byterange = Struct(
    "sized_byterange", [("data", byterange.typ), ("size", value.IntType())]
)

typ = sized_byterange.typ


def make():
    return sized_common.make(byterange.make)


# [] -> [tuple]
@modifies_stack(0, [typ])
def new(vm):
    sized_common.new(vm, byterange.new)
    vm.cast(typ)


@modifies_stack([typ], [value.IntType()])
def length(vm):
    sized_byterange.get("size")(vm)


# [tuple, index, value] -> [tuple]
@modifies_stack([typ, value.IntType(), value.IntType()], [typ])
def set_val(vm):
    sized_common.set_val(vm, sized_byterange, byterange.set_val, 32)


@modifies_stack([typ, value.IntType(), value.IntType()], [typ])
def set_val8(vm):
    sized_common.set_val(vm, sized_byterange, byterange.set_val8, 1)


# [tuple, index] -> [value]
@modifies_stack([typ, value.IntType()], [value.IntType()])
def get(vm):
    sized_common.get(vm, sized_byterange, byterange.get)


# [bytestack] -> [sized_byterange]
@modifies_stack([bytestack.typ], [typ])
def from_bytestack(vm):
    vm.dup0()
    bytestack.get("size")(vm)
    vm.swap1()
    bytestack.get("stack")(vm)

    vm.dup1()
    vm.push(31)
    vm.add()
    vm.push(32)
    vm.swap1()
    vm.div()
    vm.push(1)
    vm.swap1()
    vm.sub()
    # index stack size

    bigtuple_int.new(vm)
    # bigtuple index stack size
    vm.swap1()
    vm.swap2()
    # stack bigtuple index size
    vm.while_loop(
        lambda vm: [vm.dup0(), stack_int.isempty(vm), vm.iszero()],
        lambda vm: [
            stack_int.pop(vm),
            vm.swap1(),
            vm.auxpush(),
            # next_val bigtuple index
            vm.swap1(),
            vm.dup2(),
            vm.swap1(),
            # bigtuple index next_val index
            bigtuple_int.set_val(vm),
            # bigtuple index
            vm.swap1(),
            vm.push(1),
            vm.swap1(),
            vm.sub(),
            vm.swap1(),
            vm.auxpop()
            # stack bigtuple index size
        ],
    )
    vm.pop()
    vm.swap1()
    vm.pop()
    vm.cast(byterange.typ)

    new(vm)
    sized_byterange.set_val("data")(vm)
    sized_byterange.set_val("size")(vm)


# [sized_byterange] -> [bytestack]
@modifies_stack([typ], [bytestack.typ])
def to_bytestack(vm):
    vm.dup0()
    sized_byterange.get("size")(vm)
    vm.swap1()
    vm.push(0)
    vm.swap1()
    sized_byterange.get("data")(vm)
    stack_int.new(vm)
    tup.make(4)(vm)
    # [stack data index size]
    vm.while_loop(
        lambda vm: [vm.dup0(), vm.tgetn(3), vm.dup1(), vm.tgetn(2), vm.lt()],
        lambda vm: [
            # [stack data index size]
            vm.dup0(),
            vm.tgetn(2),
            vm.dup1(),
            vm.tgetn(1),
            byterange.get(vm),
            # cell [stack data index size]
            vm.dup1(),
            vm.tgetn(0),
            stack_int.push(vm),
            vm.swap1(),
            vm.tsetn(0),
            vm.dup0(),
            vm.tgetn(2),
            vm.push(32),
            vm.add(),
            vm.swap1(),
            vm.tsetn(2),
        ],
    )
    vm.dup0()
    vm.tgetn(3)
    vm.swap1()
    vm.tgetn(0)
    vm.tnewn(2)
    vm.cast(bytestack.typ)
    bytestack.set_val("stack")(vm)
    bytestack.set_val("size")(vm)


def get_static(val, index):
    return sized_common.get_static(val, index, byterange.get_static)


def set_static(byterange_val, index, val):
    return sized_common.set_static(byterange_val, index, val, byterange.set_static, 32)


def frombytes(data):
    return value.Tuple([byterange.frombytes(data), len(data)])


def tohex(byterange_val):
    tot = ""
    for i in range(0, byterange_val[1], 32):
        segment = eth_utils.to_hex(byterange.get_static(byterange_val[0], i))[2:]
        segment = (64 - len(segment)) * "0" + segment
        tot += segment
    return "0x" + tot[: byterange_val[1] * 2]


def create_sized_bytearray(data):
    vm = VM()
    new(vm)
    for item in data:
        vm.push(item[0])
        vm.push(item[1])
        vm.swap2()
        set(vm)
    return vm.stack[0]
