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

from ..annotation import noreturn, modifies_stack
from .. import std
from . import call_frame
from . import os
from .. import ast
from .. import value
from .types import local_exec_state


@noreturn
def _get_call_location(vm, dispatch_func):
    # contract_id
    dispatch_func(vm)
    vm.dup0()
    vm.tnewn(0)
    vm.eq()
    vm.ifelse(lambda vm: vm.error())

@noreturn
def _perform_call(vm, call_num):
    # destCodePoint destId message
    os.get_chain_state(vm)
    os.chain_state.set_val("scratch")(vm)
    os.set_chain_state(vm)
    vm.push(ast.AVMLabel("evm_call_{}".format(call_num)))
    vm.swap2()
    vm.swap1()
    # contractID message ret_pc destCodePoint

    # setup call frame
    os.get_call_frame(vm)
    call_frame.spawn(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    os.add_message_to_wallet(vm)

    std.stack_manip.compress(vm)
    std.stack_manip.compress_aux(vm)
    # compressed_stack compressed_aux_stack
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.set_val("saved_aux_stack")(vm)
    call_frame.call_frame.set_val("saved_stack")(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.set_val("parent_frame")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)

    # Enter call frame
    os.get_chain_state(vm)
    os.chain_state.get("scratch")(vm)
    vm.jump()

    vm.set_label(ast.AVMLabel("evm_call_{}".format(call_num)))
    vm.auxpush()

    std.stack_manip.kill(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.get("saved_stack")(vm)
    std.stack_manip.uncompress(vm)
    vm.auxpop()
    std.stack_manip.kill_aux(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.get("saved_aux_stack")(vm)
    std.stack_manip.uncompress_aux(vm)


@noreturn
def setup_initial_call(vm, dispatch_func):
    # contractID message
    vm.set_exception_handler(invalid_tx)
    os.get_chain_state(vm)
    os.chain_state.get("contracts")(vm)
    call_frame.new_fresh(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)

    vm.dup0()
    _get_call_location(vm, dispatch_func)

    _perform_call(vm, "initial")

    vm.clear_exception_handler()
    vm.auxpush()

    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")
    call_frame.merge(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)

    os.get_call_frame(vm)
    call_frame.call_frame.get("contracts")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("contracts")(vm)
    os.set_chain_state(vm)

    os.get_call_frame(vm)
    call_frame.call_frame.get("sent_queue")(vm)
    os.send_all_in_sent_queue(vm)

    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.get("return_data")(vm)
    vm.auxpop()
    os.log_func_result(vm)


# [[gas, dest, value, arg offset, arg length, ret offset, ret length]]
@noreturn
def call(vm, dispatch_func, call_num, contract_id):
    std.tup.make(7)(vm)
    vm.dup0()
    os.is_simple_send(vm)
    vm.ifelse(lambda vm: [
        # call sends no gas, has no arguments, and gets no return
        os.evm_call_to_send(vm),
        os.add_send_to_queue(vm)
    ], lambda vm: [
        vm.dup0(),
        vm.tgetn(1),
        vm.push(1),
        vm.eq(),
        vm.ifelse(lambda vm: [
            os.evm_call_to_send(vm),
            vm.dup0(),
            local_exec_state.get("data")(vm),
            vm.push(0),
            vm.swap1(),
            std.sized_byterange.get(vm),
            vm.push(224),
            vm.swap1(),
            std.bitwise.shift_right(vm),
            vm.dup0(),
            vm.push(0xda795ea3),
            vm.eq(),
            vm.ifelse(lambda vm: [
                vm.pop(),
                send_erc20_interupt(vm)
            ], lambda vm: [
                vm.push(0x8e725ee1),
                vm.eq(),
                vm.ifelse(send_erc721_interupt, lambda vm: vm.error())
            ])
        ], lambda vm: [
            _inner_call(vm, dispatch_func, call_num, contract_id)
        ])
    ])

def _send_interupt(vm):
    def get_dest(vm):
        vm.push(4)
        vm.swap1()
        std.sized_byterange.get(vm)

    def get_token_type(vm):
        vm.push(36)
        vm.swap1()
        std.sized_byterange.get(vm)

    def get_token_val(vm):
        vm.push(68)
        vm.swap1()
        std.sized_byterange.get(vm)
    vm.tgetn(0)
    # data
    vm.dup0()
    get_token_val(vm)
    # val data
    vm.dup1()
    get_token_type(vm)
    # type val data
    vm.swap2()
    get_dest(vm)
    # dest val type
    vm.push(local_exec_state.make())
    vm.cast(local_exec_state.typ)
    local_exec_state.set_val("sender")(vm)
    local_exec_state.set_val("amount")(vm)
    local_exec_state.set_val("type")(vm)

def send_erc20_interupt(vm):
    _send_interupt(vm)
    vm.dup0()
    local_exec_state.get("type")(vm)
    vm.push(256)
    vm.mul()
    vm.swap1()
    local_exec_state.set_val("type")(vm)
    os.add_send_to_queue(vm)

def send_erc721_interupt(vm):
    _send_interupt(vm)
    vm.dup0()
    local_exec_state.get("type")(vm)
    vm.push(256)
    vm.mul()
    vm.push(1)
    vm.add()
    vm.swap1()
    local_exec_state.set_val("type")(vm)
    os.add_send_to_queue(vm)


@modifies_stack([value.IntType(), value.TupleType([value.IntType()]*7)], [value.IntType()])
def _mutable_call_ret(vm):
    translate_ret_type(vm)
    # return_val calltup
    vm.ifelse(lambda vm: [
        os.get_call_frame(vm),
        vm.dup0(),
        call_frame.call_frame.get("parent_frame")(vm),
        call_frame.merge(vm),
        # parent_frame
        os.get_chain_state(vm),
        os.chain_state.set_val("call_frame")(vm),
        os.set_chain_state(vm),
        vm.push(1),
    ], lambda vm: [
        os.get_call_frame(vm),
        call_frame.call_frame.get("parent_frame")(vm),
        os.get_chain_state(vm),
        os.chain_state.set_val("call_frame")(vm),
        os.set_chain_state(vm),
        vm.push(0)
    ])
    vm.swap1()
    os.copy_return_data(vm)


@modifies_stack([], [])
def _save_call_frame(vm):
    os.get_call_frame(vm)
    call_frame.save_state(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


# [[gas, dest, value, arg offset, arg length, ret offset, ret length]]
@noreturn
def _inner_call(vm, dispatch_func, call_num, contract_id):
    vm.dup0()
    os.evm_call_to_send(vm)
    # msg
    vm.dup0()
    vm.tgetn(1)
    # contractId msg
    vm.swap1()
    vm.push(contract_id)
    vm.swap1()
    vm.tsetn(1)
    vm.swap1()

    # destID msg
    _save_call_frame(vm)

    vm.dup0()
    _get_call_location(vm, dispatch_func)

    _perform_call(vm, call_num)
    _mutable_call_ret(vm)

# [gas, dest, value, arg offset, arg length, ret offset, ret length]
@noreturn
def callcode(vm, dispatch_func, call_num, contract_id):
    std.tup.make(7)(vm)
    # calltup
    vm.dup0()
    os.evm_call_to_send(vm)
    # msg calltup
    vm.dup0()
    vm.tgetn(1)
    _get_call_location(vm, dispatch_func)
    # destCodePoint msg calltup
    vm.swap1()
    vm.push(contract_id)
    vm.swap1()
    vm.tsetn(1)
    # msg destCodePoint calltup

    vm.push(contract_id)
    vm.swap1()
    vm.swap2()

    # destCodePoint destId message calltup
    _save_call_frame(vm)
    _perform_call(vm, call_num)
    _mutable_call_ret(vm)

# [[gas, dest, arg offset, arg length, ret offset, ret length]]
@noreturn
def delegatecall(vm, dispatch_func, call_num, contract_id):
    os.message_value(vm)
    # value, gas, dest
    vm.swap2()
    vm.swap1()
    # gas, dest, value
    std.tup.make(7)(vm)
    # calltup
    vm.dup0()
    os.evm_call_to_send(vm)
    # msg calltup
    vm.dup0()
    vm.tgetn(1)
    _get_call_location(vm, dispatch_func)
    # destCodePoint msg calltup
    vm.swap1()
    os.message_caller(vm)
    vm.swap1()
    vm.tsetn(1)
    # msg destCodePoint calltup

    vm.push(contract_id)
    vm.swap1()
    vm.swap2()

    # destCodePoint destId message calltup
    _save_call_frame(vm)
    _perform_call(vm, call_num)
    _mutable_call_ret(vm)


# [[gas, dest, value, arg offset, arg length, ret offset, ret length]]
@noreturn
def staticcall(vm, dispatch_func, call_num, contract_id):
    vm.push(0)
    # value, gas, dest
    vm.swap2()
    vm.swap1()
    # gas, dest, value
    std.tup.make(7)(vm)

    # calltup
    vm.dup0()
    os.evm_call_to_send(vm)
    # msg calltup
    vm.dup0()
    vm.tgetn(1)
    # contractId msg calltup
    vm.swap1()
    vm.push(contract_id)
    vm.swap1()
    vm.tsetn(1)  # update the sender to this contract
    vm.swap1()
    # contractId msg calltup
    os.get_call_frame(vm)
    call_frame.save_state(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)

    # contractId msg calltup
    vm.dup0()
    _get_call_location(vm, dispatch_func)
    _perform_call(vm, "static_{}".format(call_num))
    translate_ret_type(vm)
    # ret calltup
    vm.swap1()

    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    # calltup ret
    os.copy_return_data(vm)


@noreturn
def selfdestruct(vm):
    os.get_call_frame(vm)
    call_frame.call_frame.get("sent_queue")(vm)
    # send waiting messages
    os.send_all_in_sent_queue(vm)
    vm.pop()  # address to transfer all funds to
    vm.halt()


# [offset, length]
@noreturn
def ret(vm):
    vm.dup1()
    vm.swap1()
    os.get_mem_segment(vm)
    std.tup.make(2)(vm)
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
    vm.dup0()
    call_frame.call_frame.get("return_location")(vm)
    vm.swap1()
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(2)
    vm.swap1()
    vm.jump()


@noreturn
def stop(vm):
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
    vm.push(3)
    vm.swap1()
    vm.jump()


# [memory offset, memory length]
@noreturn
def revert(vm):
    vm.debug()
    vm.dup1()
    vm.swap1()
    os.get_mem_segment(vm)
    std.tup.make(2)(vm)
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
    vm.dup0()
    call_frame.call_frame.get("return_location")(vm)
    vm.swap1()
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)
    vm.push(0)
    vm.swap1()
    vm.jump()


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


@modifies_stack([value.IntType()], [value.IntType()])
def translate_ret_type(vm):
    vm.push(1)
    vm.lt()
