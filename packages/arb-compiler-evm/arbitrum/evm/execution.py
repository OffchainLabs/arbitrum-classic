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
from . import os, arbsys, call_finish
from .. import ast
from .. import value
from .accounts import account_state, account_store
from .types import local_exec_state


@modifies_stack(0, 0)
def save_up_call_frame(vm):
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.merge(vm)
    os.set_call_frame(vm)


@modifies_stack(0, 0)
def clear_up_call_frame(vm):
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    os.set_call_frame(vm)


def setup_initial_call_frame(vm, label):
    # caller
    vm.push(ast.AVMLabel("evm_call_{}".format(label)))
    os.get_chain_state(vm)
    os.chain_state.get("accounts")(vm)
    call_frame.new_fresh(vm)
    call_frame.call_frame.set_val("return_location")(vm)
    call_frame.call_frame.set_val("contractID")(vm)
    call_frame.setup_state(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


@noreturn
def save_stacks(vm):
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


@noreturn
def _perform_call(vm, call_num):
    # dest local_exec_state
    vm.dup0()
    vm.push(100)
    vm.eq()
    vm.ifelse(
        lambda vm: [vm.pop(), arbsys.perform_precompile_call(vm)],
        lambda vm: [_perform_real_call(vm, call_num)],
    )


def _perform_real_call(vm, call_num):
    os.get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)
    # balance dest local_exec_state
    vm.dup2()
    local_exec_state.get("value")(vm)
    std.comparison.lte(vm)
    vm.ifelse(
        lambda vm: [_execute_call(vm, call_num)],
        lambda vm: [
            os.get_call_frame(vm),
            call_frame.spawn_child(vm),
            os.set_call_frame(vm),
            vm.push(0),
        ],
    )
    _complete_call(vm, call_num)


@noreturn
def _enter_exec(vm, call_num):
    vm.dup0()
    vm.tnewn(0)
    vm.eq()
    vm.ifelse(
        lambda vm: [
            vm.pop(),
            vm.push(3),
            vm.push(ast.AVMLabel("evm_call_{}".format(call_num))),
            vm.jump(),
        ],
        lambda vm: [vm.jump()],
    )


@noreturn
def _execute_call(vm, call_num):
    vm.push(ast.AVMLabel("evm_call_{}".format(call_num)))
    vm.swap2()
    # local_exec_state destId ret_pc

    # setup call frame
    os.get_call_frame(vm)
    call_frame.spawn(vm)
    os.set_call_frame(vm)

    save_stacks(vm)

    # Enter call frame
    os.get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("code_point")(vm)
    _enter_exec(vm, call_num)


@noreturn
def _perform_callcode(vm, call_num):
    # dest local_exec_state
    vm.dup0()
    os.get_call_frame(vm)
    call_frame.call_frame.get("accounts")(vm)
    account_store.get(vm)
    account_state.get("code_point")(vm)
    os.set_scratch(vm)

    vm.push(ast.AVMLabel("evm_call_{}".format(call_num)))
    vm.swap2()
    # local_exec_state destId ret_pc

    # setup call frame
    os.get_call_frame(vm)
    call_frame.spawn_callcode(vm)
    os.set_call_frame(vm)

    save_stacks(vm)

    # Enter call frame
    os.get_scratch(vm)
    _enter_exec(vm, call_num)

    _complete_call(vm, call_num)


@noreturn
def _perform_delegatecall(vm, call_num):
    # dest local_exec_state
    os.get_call_frame(vm)
    call_frame.call_frame.get("accounts")(vm)
    account_store.get(vm)
    account_state.get("code_point")(vm)
    os.set_scratch(vm)

    # local_exec_state
    vm.push(ast.AVMLabel("evm_call_{}".format(call_num)))
    vm.swap1()
    # local_exec_state ret_pc

    # setup call frame
    os.get_call_frame(vm)
    call_frame.spawn_delegatecall(vm)
    os.set_call_frame(vm)

    save_stacks(vm)

    # Enter call frame
    os.get_scratch(vm)
    _enter_exec(vm, call_num)

    _complete_call(vm, call_num)


@noreturn
def _complete_call(vm, call_num):
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


@modifies_stack([value.IntType()], [])
def _handle_mutable_call_return(vm):
    # ret_type calltup
    translate_ret_type(vm)
    # return_val calltup
    vm.ifelse(lambda vm: [save_up_call_frame(vm)], lambda vm: [clear_up_call_frame(vm)])


@modifies_stack([], [])
def _save_call_frame(vm):
    os.get_call_frame(vm)
    call_frame.save_state(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


@noreturn
def initial_call(vm, label):
    # sender tx_call_data
    vm.set_exception_handler(call_finish.invalid_tx)
    vm.dup0()
    setup_initial_call_frame(vm, label)
    os.tx_call_to_local_exec_state(vm)
    _perform_call(vm, label)
    # ret_code
    vm.dup0()
    os.get_call_frame(vm)
    call_frame.save_state(vm)
    os.set_call_frame(vm)
    _handle_mutable_call_return(vm)
    vm.auxpush()
    vm.clear_exception_handler()
    os.get_call_frame(vm)
    # frame
    vm.dup0()
    call_frame.call_frame.get("accounts")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("accounts")(vm)
    os.set_chain_state(vm)
    # frame
    call_frame.call_frame.get("return_data")(vm)

    vm.auxpop()
    # ret_code data
    os.log_func_result(vm)


@noreturn
def initial_static_call(vm, label):
    # sender tx_call_data
    vm.set_exception_handler(call_finish.invalid_tx)
    vm.dup0()
    setup_initial_call_frame(vm, label)
    os.tx_call_to_local_exec_state(vm)
    _perform_call(vm, label)
    # ret_code
    vm.clear_exception_handler()

    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.get("return_data")(vm)
    vm.swap1()

    # ret_code data
    vm.swap1()
    vm.dup1()
    os.log_func_result(vm)
    # ret_code


# [[gas, dest, value, arg offset, arg length, ret offset, ret length]]
@noreturn
def call(vm, call_num):
    std.tup.make(7)(vm)
    vm.dup0()
    os.evm_call_to_tx_call_data(vm)
    # tx_call calltup
    os.get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    os.tx_call_to_local_exec_state(vm)
    # dest local_exec_state calltup

    _save_call_frame(vm)
    _perform_call(vm, call_num)
    # ret calltup
    vm.dup0()
    _handle_mutable_call_return(vm)
    translate_ret_type(vm)
    # return_val calltup
    vm.swap1()
    os.copy_return_data(vm)


# [gas, dest, value, arg offset, arg length, ret offset, ret length]
@noreturn
def callcode(vm, call_num):
    std.tup.make(7)(vm)
    # calltup
    vm.dup0()
    os.evm_call_to_tx_call_data(vm)
    # msg calltup
    os.get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    os.tx_call_to_local_exec_state(vm)
    # dest local_exec_state calltup

    _save_call_frame(vm)
    _perform_callcode(vm, call_num)
    vm.dup0()
    _handle_mutable_call_return(vm)
    translate_ret_type(vm)
    # return_val calltup
    vm.swap1()
    os.copy_return_data(vm)


# [gas, dest, arg offset, arg length, ret offset, ret length]
@noreturn
def delegatecall(vm, call_num):
    os.message_value(vm)
    # value, gas, dest
    vm.swap2()
    vm.swap1()
    # gas, dest, value
    std.tup.make(7)(vm)
    # calltup
    vm.dup0()
    os.evm_call_to_tx_call_data(vm)
    # msg calltup
    os.message_caller(vm)
    os.tx_call_to_local_exec_state(vm)

    # dest destId message calltup
    _save_call_frame(vm)
    _perform_delegatecall(vm, call_num)
    vm.dup0()
    _handle_mutable_call_return(vm)
    translate_ret_type(vm)
    # return_val calltup
    vm.swap1()
    os.copy_return_data(vm)


# [[gas, dest, arg offset, arg length, ret offset, ret length]]
@noreturn
def staticcall(vm, call_num):
    vm.push(0)
    # value, gas, dest
    vm.swap2()
    vm.swap1()
    # gas, dest, value
    std.tup.make(7)(vm)

    # calltup
    vm.dup0()
    os.evm_call_to_tx_call_data(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    os.tx_call_to_local_exec_state(vm)
    # dest msg calltup
    _save_call_frame(vm)

    # dest msg calltup
    _perform_call(vm, "static_{}".format(call_num))
    translate_ret_type(vm)
    # ret calltup
    vm.swap1()

    clear_up_call_frame(vm)
    # calltup ret
    os.copy_return_data(vm)


# TODO: IMPLEMENT EVM SELFDESTRUCT
@noreturn
def selfdestruct(vm):
    vm.pop()  # address to transfer all funds to
    vm.halt()


@modifies_stack([value.IntType()], [value.IntType()])
def translate_ret_type(vm):
    vm.push(1)
    vm.lt()
