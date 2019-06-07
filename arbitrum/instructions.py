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

from .value import ValueType, IntType, TupleType, CodePointType, Tuple

OP_CODES = [
    # Arithmetic
    ('add', 0x01, [IntType(), IntType()], [IntType()]),
    ('mul', 0x02, [IntType(), IntType()], [IntType()]),
    ('sub', 0x03, [IntType(), IntType()], [IntType()]),
    ('div', 0x04, [IntType(), IntType()], [IntType()]),
    ('sdiv', 0x05, [IntType(), IntType()], [IntType()]),
    ('mod', 0x06, [IntType(), IntType()], [IntType()]),
    ('smod', 0x07, [IntType(), IntType()], [IntType()]),
    ('addmod', 0x08, [IntType(), IntType(), IntType()], [IntType()]),
    ('mulmod', 0x09, [IntType(), IntType(), IntType()], [IntType()]),
    ('exp', 0x0a, [IntType(), IntType()], [IntType()]),

    # Comparison & Bitwise Logic Operations
    ('lt', 0x10, [IntType(), IntType()], [IntType()]),
    ('gt', 0x11, [IntType(), IntType()], [IntType()]),
    ('slt', 0x12, [IntType(), IntType()], [IntType()]),
    ('sgt', 0x13, [IntType(), IntType()], [IntType()]),
    ('eq', 0x14, [ValueType(), ValueType()], [IntType()]),
    ('iszero', 0x15, [IntType()], [IntType()]),
    ('bitwise_and', 0x16, [IntType(), IntType()], [IntType()]),
    ('bitwise_or', 0x17, [IntType(), IntType()], [IntType()]),
    ('bitwise_xor', 0x18, [IntType(), IntType()], [IntType()]),
    ('bitwise_not', 0x19, [IntType()], [IntType()]),
    ('byte', 0x1a, [IntType(), IntType()], [IntType()]),
    ('signextend', 0x1b, [IntType(), IntType()], [IntType()]),

    # SHA3
    ('hash', 0x20, [ValueType()], [IntType()]),
    ('type', 0x21, [ValueType()], [IntType()]),

    # Stack, Memory, Storage and Flow Operations
    ('pop', 0x30, [ValueType()], []),
    ('spush', 0x31, [], [ValueType()]),
    ('rpush', 0x32, [], [ValueType()]),
    ('rset', 0x33, [ValueType()], []),
    ('jump', 0x34, [CodePointType()], []),
    ('cjump', 0x35, [CodePointType(), IntType()], []),
    ('stackempty', 0x36, [], [IntType()]),
    ('pcpush', 0x37, [], [CodePointType()]),
    ('auxpush', 0x38, [ValueType()], []),
    ('auxpop', 0x39, [], [ValueType()]),
    ('auxstackempty', 0x3a, [], [IntType()]),
    ('nop', 0x3b, [], []),
    ('errpush', 0x3c, [ValueType()], []),
    ('errset', 0x3d, [ValueType()], []),

    # Duplication and Exchange Operations
    ('dup0', 0x40, [ValueType()], [ValueType(), ValueType()]),
    (
        'dup1',
        0x41,
        [ValueType(), ValueType()],
        [ValueType(), ValueType(), ValueType()]
    ),
    (
        'dup2',
        0x42,
        [ValueType(), ValueType(), ValueType()],
        [ValueType(), ValueType(), ValueType(), ValueType()]
    ),
    (
        'swap1',
        0x43,
        [ValueType(), ValueType()],
        [ValueType(), ValueType()]
    ),
    (
        'swap2',
        0x44,
        [ValueType(), ValueType(), ValueType()],
        [ValueType(), ValueType(), ValueType()]
    ),

    # Tuple Operations
    ('tget', 0x50, [IntType(), TupleType()], [ValueType()]),
    ('tset', 0x51, [IntType(), TupleType(), ValueType()], [TupleType()]),
    ('tlen', 0x52, [TupleType()], [IntType()]),

    # Logging Operations
    ('breakpoint', 0x60, [ValueType()], []),
    ('log', 0x61, [ValueType()], []),

    # System operations
    ('send', 0x70, [TupleType()], []),
    ('nbsend', 0x71, [TupleType()], [IntType()]),
    ('gettime', 0x72, [], [TupleType()]),
    ('inbox', 0x73, [TupleType()], [TupleType()]),
    ('error', 0x74, [], []),
    ('halt', 0x75, [], []),
    ('debug', 0x76, [], [])
]


def swap1(stack):
    item0 = stack.pop()
    item1 = stack.pop()
    stack.push(item0)
    stack.push(item1)


def swap2(stack):
    item0 = stack.pop()
    item1 = stack.pop()
    item2 = stack.pop()
    stack.push(item0)
    stack.push(item1)
    stack.push(item2)


def dup0(stack):
    item0 = stack.pop()
    stack.push(item0)
    stack.push(item0)


def dup1(stack):
    item0 = stack.pop()
    item1 = stack.pop()
    stack.push(item1)
    stack.push(item0)
    stack.push(item1)


def dup2(stack):
    item0 = stack.pop()
    item1 = stack.pop()
    item2 = stack.pop()
    stack.push(item2)
    stack.push(item1)
    stack.push(item0)
    stack.push(item2)


def tlen(stack):
    tup = stack.pop(TupleType())
    stack.push(tup.size())


def tnew(stack):
    size = stack.pop(IntType())
    stack.push(Tuple([Tuple([]) for i in range(size)]))


def tget(stack):
    index = stack.pop(IntType())
    tup = stack.pop(TupleType())
    if not tup.has_member_at_index(index):
        raise Exception("Tried to get index {} from tuple {}".format(index, tup))
    stack.push(tup.get_tup(index))


def tgetn(stack, index):
    tup = stack.pop(TupleType())
    if not tup.has_member_at_index(index):
        raise Exception("Tried to get index {} from tuple {}".format(index, tup))
    stack.push(tup.get_tup(index))


def tset(stack):
    index = stack.pop(IntType())
    tup = stack.pop(TupleType())
    val = stack.pop(ValueType())
    stack.push(tup.set_tup_val(index, val))


def tsetn(stack, index):
    tup = stack.pop(TupleType())
    val = stack.pop(ValueType())
    stack.push(tup.set_tup_val(index, val))


OPS = {}
OP_NAMES = {}
OF_INFO = {}
for (op_name, op_code, pops, pushes) in OP_CODES:
    OPS[op_name] = op_code
    OP_NAMES[op_code] = op_name
    OF_INFO[op_code] = {"pop": pops, "push": pushes}
