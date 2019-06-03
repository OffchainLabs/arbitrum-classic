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

from ..annotation import modifies_stack
from .. import value


# [number, bits]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def shift_left(vm):
    vm.swap1()
    vm.push(2)
    vm.exp()
    vm.mul()


# [number, bits]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def shift_right(vm):
    vm.swap1()
    vm.push(2)
    vm.exp()
    vm.swap1()
    vm.div()


# [bits]
# 2 ** bits - 1
@modifies_stack([value.IntType()], [value.IntType()])
def n_lowest_mask(vm):
    vm.push(2)
    vm.exp()
    vm.push(1)
    vm.swap1()
    vm.sub()


# [bits]
# 2 ** bits - 1 << (256 - bits)
@modifies_stack([value.IntType()], [value.IntType()])
def n_highest_mask(vm):
    vm.dup0()
    n_lowest_mask(vm)
    vm.swap1()
    vm.push(256)
    vm.sub()
    vm.swap1()
    shift_left(vm)


def n_lowest_mask_static(bits):
    return 2**bits - 1


def n_highest_mask_static(bits):
    return n_lowest_mask_static(bits) << (256 - bits)


@modifies_stack([value.IntType()], [value.IntType()])
def flip_endianness(vm):
    flip_endianness_impl(vm, 32)

def flip_endianness_impl(vm, numBytes):
    if numBytes>1:
        nb2 = numBytes//2
        mod = 1<<(8*nb2)
        vm.push(mod)
        vm.dup1()
        # x mod x
        vm.div()
        # x//mod x
        flip_endianness_impl(vm, nb2)
        # flipped(x//mod) x
        vm.swap1()
        vm.push(mod)
        vm.swap1()
        # x mod flipped(x//mod)
        vm.mod()
        flip_endianness_impl(vm, nb2)
        # flipped(x%mod) flipped(x//mod)
        vm.push(mod)
        vm.mul()
        vm.bitwise_or()

# [int, index, byte]
@modifies_stack([value.IntType(), value.IntType(), value.IntType()], [value.IntType()])
def set_byte(vm):
    vm.dup0()
    vm.dup2()
    # [index, int, int, index, byte]
    vm.byte()
    # # [orig, int, index, byte]
    vm.auxpush()
    vm.swap2()
    vm.push(0xff)
    vm.bitwise_and()
    vm.auxpop()
    # # [orig, byte, index, int]
    vm.bitwise_xor()
    vm.swap1()
    vm.push(8)
    vm.mul()
    vm.push(248)
    vm.sub()
    # # [bit index, updated_byte, int]
    vm.swap1()
    shift_left(vm)
    vm.bitwise_xor()
