# Copyright 2019-2020, Offchain Labs, Inc.
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

from ..annotation import noreturn, modifies_stack
from .. import std
from . import call_frame
from . import os
from .. import value

# [offset, length]
@noreturn
def ret(vm):
    ret_impl(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("return_location")(vm)
    vm.jump()


# [offset, length]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def ret_impl(vm):
    vm.dup1()
    vm.swap1()
    os.get_mem_segment(vm)
    std.tup.make(2)(vm)
    vm.cast(std.sized_byterange.typ)
    # return_data
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")(vm)
    # parent_frame current_frame return_data
    vm.swap1()
    vm.swap2()
    # return_data parent_frame current_frame
    vm.swap1()
    call_frame.call_frame.set_val("return_data")(vm)
    # parent_frame current_frame
    vm.swap1()
    call_frame.call_frame.set_val("parent_frame")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(2)


@noreturn
def stop(vm):
    stop_impl(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("return_location")(vm)
    vm.jump()


@modifies_stack([], [value.IntType()])
def stop_impl(vm):
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")(vm)
    # parent_frame current_frame
    std.sized_byterange.new(vm)
    vm.swap1()
    call_frame.call_frame.set_val("return_data")(vm)
    vm.swap1()
    call_frame.call_frame.set_val("parent_frame")(vm)
    # call_frame
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(3)


# [memory offset, memory length]
@noreturn
def revert(vm):
    revert_impl(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("return_location")(vm)
    vm.jump()


# [memory offset, memory length]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def revert_impl(vm):
    vm.dup1()
    vm.swap1()
    os.get_mem_segment(vm)
    std.tup.make(2)(vm)
    vm.cast(std.sized_byterange.typ)
    # return_data
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")(vm)
    # parent_frame current_frame return_data
    vm.swap1()
    vm.swap2()
    # return_data parent_frame current_frame
    vm.swap1()
    call_frame.call_frame.set_val("return_data")(vm)
    # parent_frame current_frame
    vm.swap1()
    call_frame.call_frame.set_val("parent_frame")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(0)


# []
@noreturn
def invalid_tx(vm):
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")(vm)
    # parent_frame current_frame
    std.sized_byterange.new(vm)
    vm.swap1()
    call_frame.call_frame.set_val("return_data")(vm)
    vm.swap1()
    call_frame.call_frame.set_val("parent_frame")(vm)
    # call_frame
    vm.dup0()
    call_frame.call_frame.get("return_location")(vm)
    # return_location call_frame
    vm.swap1()
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(1)
    vm.swap1()
    vm.jump()
