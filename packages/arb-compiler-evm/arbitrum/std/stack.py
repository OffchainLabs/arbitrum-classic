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

from ..annotation import modifies_stack
from .. import value
from .struct import Struct


def make_stack_type(typ):
    stack_type = Struct("stack[{}]".format(typ), ["prev", ("val", typ)])

    class Stack:
        # [] -> [stack]
        @staticmethod
        @modifies_stack([], [stack_type.typ])
        def new(vm):
            vm.tnewn(0)
            vm.cast(stack_type.typ)

        @staticmethod
        @modifies_stack([stack_type.typ], [value.IntType()])
        def isempty(vm):
            vm.tnewn(0)
            vm.eq()

        # [stack, value] -> [stack]
        @staticmethod
        @modifies_stack([stack_type.typ, typ], [stack_type.typ])
        def push(vm):
            stack_type.new(vm)
            stack_type.set_val("prev")(vm)
            stack_type.set_val("val")(vm)

        # [stack] -> [value, stack]
        @staticmethod
        @modifies_stack([stack_type.typ], [typ, stack_type.typ])
        def pop(vm):
            vm.dup0()
            stack_type.get("prev")(vm)
            vm.cast(stack_type.typ)
            vm.swap1()
            stack_type.get("val")(vm)

        @staticmethod
        def to_list(stack):
            items = []
            while stack != value.Tuple([]):
                items.append(stack[1])
                stack = stack[0]
            return items

    Stack.typ = stack_type.typ
    return Stack


stack = make_stack_type(value.ValueType())
stack_int = make_stack_type(value.IntType())
stack_tup = make_stack_type(value.TupleType())
stack_code = make_stack_type(value.CodePointType())

bytestack = Struct("bytestack", [("size", value.IntType()), ("stack", stack_int.typ)])


def bytestack_frombytes(data):
    data_length = len(data)
    if len(data) % 32 != 0:
        data = data + b"\0" * (32 - (len(data) % 32))
    chunks = [
        eth_utils.big_endian_to_int(data[i : i + 32]) for i in range(0, len(data), 32)
    ]
    st = value.Tuple([])
    for chunk in chunks:
        st = value.Tuple([st, chunk])
    return value.Tuple([data_length, st])


def bytestack_tohex(st):
    data = ""
    data_length = st[0]
    st = st[1]
    while len(st) != 0:
        segment = eth_utils.to_hex(st[1])[2:]
        segment = (64 - len(segment)) * "0" + segment
        data = segment + data
        st = st[0]
    return "0x" + data[: data_length * 2]
