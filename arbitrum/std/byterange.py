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
from . import tup
from . import bigtuple_int
from . import bitwise
from .. import value
from ..annotation import modifies_stack
from ..value import IntType
from .struct import Struct
from .closure import make_closure

TT256M1 = 2 ** 256 - 1

byterange = Struct("byterange", [
    ("bigtuple", bigtuple_int.typ)
])

typ = byterange.typ


# [break_point, number, old number]
@modifies_stack([IntType(), IntType(), IntType()], [IntType()])
def _set_first_half(vm):
    vm.swap2()
    vm.dup2()
    bitwise.n_highest_mask(vm)
    vm.bitwise_and()
    # [old number masked, number, break_point]
    vm.swap2()
    vm.swap1()
    # [number, break_point, old number]
    bitwise.shift_right(vm)
    vm.bitwise_or()


# [break_point, number, old number] -> [new number]
@modifies_stack([IntType(), IntType(), IntType()], [IntType()])
def _set_second_half(vm):
    vm.push(256)
    vm.sub()
    vm.swap2()
    # [old number, number, 256 - break_point]
    vm.dup2()
    bitwise.n_lowest_mask(vm)
    vm.bitwise_and()

    vm.swap2()
    vm.swap1()
    # [number, 256 - break_point, old number]
    bitwise.shift_left(vm)
    vm.bitwise_or()


# [number, break_point]
@modifies_stack([IntType(), IntType()], [IntType()])
def _get_first_half(vm):
    bitwise.shift_left(vm)


# [number, break_point]
@modifies_stack([IntType(), IntType()], [IntType()])
def _get_second_half(vm):
    vm.swap1()
    vm.push(256)
    vm.sub()
    vm.swap1()
    bitwise.shift_right(vm)


def make():
    return bigtuple_int.make()


# []
@modifies_stack([], [typ])
def new(vm):
    vm.push(make())
    vm.cast(typ)


# [tuple, index, value]
@modifies_stack([typ, IntType(), IntType()], [typ])
def set_val(vm):
    vm.dup1()
    vm.push(32)
    vm.swap1()
    vm.mod()
    vm.push(0)
    vm.eq()
    vm.ifelse(lambda vm: [
        vm.swap1(),
        vm.push(32),
        vm.swap1(),
        vm.div(),
        vm.swap1(),
        byterange.get("bigtuple")(vm),
        bigtuple_int.set_val(vm),
        byterange.set_val("bigtuple")(vm)
    ], lambda vm: [
        _set_impl(vm)
    ])


# [index] -> [index / 32, (index % 32) * 8]
@modifies_stack([IntType()], [IntType(), IntType()])
def _mod_div_impl(vm):
    vm.dup0()
    vm.push(32)
    vm.swap1()
    vm.mod()
    vm.push(8)
    vm.mul()

    vm.swap1()
    vm.push(32)
    vm.swap1()
    vm.div()


# # [tuple, index, value]
@modifies_stack([typ, IntType(), IntType()], [typ])
def _set_impl(vm):
    first_half_closure = make_closure(_set_first_half, 2)
    second_half_closure = make_closure(_set_second_half, 2)

    byterange.get("bigtuple")(vm)
    vm.auxpush()
    # [index, value]
    _mod_div_impl(vm)
    # [index / 32, (index % 32) * 8, value]
    vm.swap2()
    vm.dup1()
    vm.dup1()
    first_half_closure.new(vm)
    vm.auxpush()
    second_half_closure.new(vm)
    # [closure2, index / 32]
    vm.swap1()
    vm.auxpop()
    # [closure1, index / 32, closure2]
    vm.dup1()
    vm.auxpop()
    # [tuple, index / 32, closure1, index / 32, closure2]
    bigtuple_int.read_modify_write(first_half_closure)(vm)
    # [tuple, index / 32, closure2]
    vm.swap1()
    vm.push(1)
    vm.add()
    vm.swap1()
    bigtuple_int.read_modify_write(second_half_closure)(vm)
    byterange.set_val("bigtuple")(vm)


# [tuple, index]
@modifies_stack([typ, IntType()], [IntType()])
def get(vm):
    vm.dup1()
    vm.push(32)
    vm.swap1()
    vm.mod()
    vm.push(0)
    vm.eq()
    vm.ifelse(lambda vm: [
        vm.swap1(),
        vm.push(32),
        vm.swap1(),
        vm.div(),
        vm.swap1(),
        byterange.get("bigtuple")(vm),
        bigtuple_int.get(vm),
    ], lambda vm: [
        _get_impl(vm)
    ])


# [tuple, index]
@modifies_stack([typ, IntType()], [IntType()])
def _get_impl(vm):
    vm.swap1()

    _mod_div_impl(vm)

    vm.swap1()
    vm.swap2()
    vm.dup2()
    vm.dup2()

    vm.dup2()
    byterange.get("bigtuple")(vm)
    bigtuple_int.get(vm)
    _get_first_half(vm)

    vm.swap2()
    vm.push(1)
    vm.add()
    vm.swap1()
    byterange.get("bigtuple")(vm)
    bigtuple_int.get(vm)

    vm.swap1()
    vm.swap2()
    vm.swap1()
    _get_second_half(vm)
    vm.bitwise_or()


copy_types = [
    typ,
    value.IntType(),
    value.IntType(),
    typ,
    value.IntType()
]


# [source bytearray, start offset, end offset, dest bytearray, dest offset]
@modifies_stack(copy_types, [typ])
def copy(vm):
    tup.make(5)(vm)
    _copy_impl(vm)
    vm.tgetn(3)


@modifies_stack([value.TupleType(copy_types)], [value.TupleType(copy_types)])
def _copy_impl(vm):

    def get_source(vm):
        vm.dup0()
        vm.tgetn(1)
        vm.dup1()
        vm.tgetn(0)
        get(vm)

    def set_dest(vm):
        vm.dup1()
        vm.tgetn(4)
        vm.dup2()
        vm.tgetn(3)
        set_val(vm)
        vm.swap1()
        vm.tsetn(3)

    def get_dest(vm):
        vm.dup0()
        vm.tgetn(4)
        vm.dup1()
        vm.tgetn(3)
        get(vm)

    # [[
    #    source bytearray,
    #    start offset,
    #    end offset,
    #    dest bytearray,
    #    dest offset
    # ]]
    vm.while_loop(lambda vm: [
        vm.dup0(),
        vm.tgetn(2),
        vm.dup1(),
        vm.tgetn(1),
        vm.push(32),
        vm.add(),
        vm.lt()
    ], lambda vm: [
        get_source(vm),
        set_dest(vm),

        # increment destination offset
        vm.dup0(),
        vm.tgetn(4),
        vm.push(32),
        vm.add(),
        vm.swap1(),
        vm.tsetn(4),

        # increment source offset
        vm.dup0(),
        vm.tgetn(1),
        vm.push(32),
        vm.add(),
        vm.swap1(),
        vm.tsetn(1)
    ])

    vm.dup0()
    vm.tgetn(1)
    vm.dup1()
    vm.tgetn(2)
    vm.sub()

    vm.dup0()
    vm.push(0)
    vm.eq()
    vm.ifelse(lambda vm: [
        vm.pop()
    ], lambda vm: [
        vm.push(32),
        vm.eq(),
        vm.ifelse(lambda vm: [
            get_source(vm),
            set_dest(vm),
        ], lambda vm: [
            # [
            #   source bytearray,
            #   start offset,
            #   end offset,
            #   dest bytearray,
            #   dest offset
            # ]

            # save remaining byte size into end offset slot
            vm.dup0(),
            vm.tgetn(1),
            # start offset [...]
            vm.dup1(),
            vm.tgetn(2),
            # end_offset start_offset [...]
            vm.sub(),
            # (end_offset- start_offset) [...]
            vm.push(8),
            vm.mul(),
            vm.swap1(),
            get_source(vm),
            # value [...] index
            vm.swap1(),
            vm.swap2(),
            vm.swap1(),
            # value index [...]
            vm.dup2(),
            get_dest(vm),
            vm.swap1(),
            vm.pop(),
            # old_num, value, index, [...]
            _merge_numbers(vm),
            set_dest(vm)
        ])
    ])


# (new & n_highest_mask(48)) | (old & n_lowest_mask(208))
# [old, new, break_point]
@modifies_stack([IntType(), IntType(), IntType()], [IntType()])
def _merge_numbers(vm):
    vm.swap1()
    vm.dup2()
    bitwise.n_highest_mask(vm)
    vm.bitwise_and()
    # new_masked old break_point
    vm.swap2()
    vm.push(256)
    vm.sub()
    bitwise.n_lowest_mask(vm)
    vm.bitwise_and()

    vm.bitwise_or()


# [tuple, index]
@modifies_stack([typ, value.IntType()], [value.IntType()])
def get8(vm):
    vm.swap1()
    _mod_div_impl(vm)
    # [index / 32, index % 32, tuple]
    vm.swap1()
    vm.push(8)
    vm.add()
    # [index % 32 + 8, index / 32, tuple]
    vm.swap2()
    byterange.get("bigtuple")(vm)
    bigtuple_int.get(vm)
    # [val, index % 32 + 8]
    _get_second_half(vm)
    vm.push(8)
    bitwise.n_lowest_mask(vm)
    vm.bitwise_and()

# [index, byte, old number] -> [new number]
@modifies_stack([IntType(), IntType(), IntType()], [IntType()])
def _update_byte(vm):
    vm.swap1()
    vm.swap2()
    # [int, index, byte]
    bitwise.set_byte(vm)


# [tuple, index, value]
@modifies_stack([typ, IntType(), IntType()], [typ])
def set_val8(vm):
    update_byte_closure = make_closure(_update_byte, 2)

    vm.swap1()
    _mod_div_impl(vm)
    # [index / 32, index % 32, tuple, value]
    vm.auxpush()
    # [index % 32, tuple, value]
    vm.swap1()
    vm.swap2()
    vm.swap1()
    # [index % 32, value, tuple]
    update_byte_closure.new(vm)
    # [closure, tuple]
    vm.swap1()
    byterange.get("bigtuple")(vm)
    # [tuple, closure]
    vm.auxpop()
    vm.swap1()
    bigtuple_int.read_modify_write(update_byte_closure)(vm)
    byterange.set_val("bigtuple")(vm)


# [source_tuple, start offset, end offset]
@modifies_stack([typ, value.IntType(), value.IntType()], [typ])
def get_subset(vm):
    tup.make(3)(vm)
    new(vm)
    vm.push(0)
    vm.swap2()
    tup.tbreak(3)(vm)
    copy(vm)


def frombytes(data):
    if len(data) % 32 != 0:
        data = data + b'\0'*(32 - (len(data) % 32))
    chunks = [
        eth_utils.big_endian_to_int(data[i: i + 32])
        for i in range(0, len(data), 32)
    ]
    return bigtuple_int.fromints(chunks)


# [tuple, index]
def get_static(tup_val, index):
    if index % 32 == 0:
        return bigtuple_int.get_static(tup_val, index // 32)

    first_half = bigtuple_int.get_static(
        tup_val,
        index // 32
    ) << ((index % 32) * 8)
    second_half = bigtuple_int.get_static(
        tup_val,
        index // 32 + 1
    ) >> (256 - ((index % 32) * 8))
    return (first_half | second_half) & TT256M1


def set_static(tup_val, index, val):
    if index % 32 == 0:
        return bigtuple_int.set_static(tup_val, index // 32, val)

    first_half = bigtuple_int.get_static(tup_val, index // 32)
    first_half &= bitwise.n_highest_mask_static((index % 32) * 8)
    first_half |= val >> (index % 32) * 8
    tup_val = bigtuple_int.set_static(tup_val, index // 32, first_half)

    second_half = bigtuple_int.get_static(tup_val, index // 32 + 1)
    second_half &= bitwise.n_lowest_mask_static(256 - (index % 32) * 8)
    second_half |= val << (256 - (index % 32) * 8)
    tup_val = bigtuple_int.set_static(tup_val, index // 32 + 1, second_half)

    return tup_val
