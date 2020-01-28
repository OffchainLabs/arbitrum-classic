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
from .. import ast
from .. import value
from .accounts import account_state, account_store
from .types import (
    local_exec_state,
    eth_transfer_message,
    ethbridge_message,
    token_transfer_message,
    message,
)

WITHDRAW_ETH_TYPECODE = 1
WITHDRAW_ERC20_TYPECODE = 2
WITHDRAW_ERC721_TYPECODE = 3


@modifies_stack(0, 0)
def save_up_call_frame(vm):
    os.get_call_frame(vm)
    vm.dup0()
    call_frame.call_frame.get("parent_frame")
    call_frame.merge(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


@modifies_stack(0, 0)
def clear_up_call_frame(vm):
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


@modifies_stack([value.IntType()], 0)
def setup_initial_call_frame(vm):
    # caller
    os.get_chain_state(vm)
    os.chain_state.get("accounts")(vm)
    call_frame.new_fresh(vm)
    call_frame.call_frame.set_val("contractID")(vm)
    call_frame.setup_state(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("call_frame")(vm)
    os.set_chain_state(vm)


@modifies_stack([call_frame.call_frame.typ], [std.sized_byterange.sized_byterange.typ])
def commit_call_frame(vm):
    # frame
    vm.dup0()
    call_frame.call_frame.get("accounts")(vm)
    os.get_chain_state(vm)
    os.chain_state.set_val("accounts")(vm)
    os.set_chain_state(vm)
    # frame
    call_frame.call_frame.get("parent_frame")(vm)
    call_frame.call_frame.get("return_data")(vm)


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
        lambda vm: [vm.pop(), _perform_precompile_call(vm)],
        lambda vm: [_perform_real_call(vm, call_num)],
    )


def _perform_precompile_call(vm):
    # local_exec_state
    vm.dup0()
    local_exec_state.get("data")(vm)
    vm.push(0)
    vm.swap1()
    std.sized_byterange.get(vm)
    vm.push(224)
    vm.swap1()
    std.bitwise.shift_right(vm)
    vm.dup0()
    vm.push(0x1B9A91A4)
    vm.eq()
    vm.ifelse(
        lambda vm: [vm.pop(), withdraw_eth_interrupt(vm)],
        lambda vm: [
            vm.dup0(),
            vm.push(0xA1DB9782),
            vm.eq(),
            vm.ifelse(
                lambda vm: [vm.pop(), withdraw_erc20_interrupt(vm)],
                lambda vm: [
                    vm.dup0(),
                    vm.push(0xF3E414F8),
                    vm.eq(),
                    vm.ifelse(
                        lambda vm: [vm.pop(), withdraw_erc721_interrupt(vm)],
                        lambda vm: [
                            vm.dup0(),
                            vm.push(0xBDE19776),
                            vm.eq(),
                            vm.ifelse(
                                lambda vm: [vm.pop(), arbsys_time_upper_bound(vm)],
                                lambda vm: [
                                    vm.dup0(),
                                    vm.push(0x44F50653),
                                    vm.eq(),
                                    vm.ifelse(
                                        lambda vm: [
                                            vm.pop(),
                                            arbsys_current_message_time(vm),
                                        ],
                                        lambda vm: [vm.pop(), vm.push(0)],
                                    ),
                                ],
                            ),
                        ],
                    ),
                ],
            ),
        ],
    )


def parse_withdraw_call(vm):
    # local_exec_state
    vm.dup0()
    local_exec_state.get("caller")(vm)
    vm.swap1()
    local_exec_state.get("data")(vm)
    vm.dup0()
    vm.push(4)
    vm.swap1()
    std.sized_byterange.get(vm)
    # dest data sender
    vm.swap1()
    vm.push(36)
    vm.swap1()
    std.sized_byterange.get(vm)
    # amount dest sender


def withdraw_eth_interrupt(vm):
    # local_exec_state
    parse_withdraw_call(vm)
    # amount dest sender
    vm.dup0()
    os.process_eth_withdraw(vm)
    vm.ifelse(
        lambda vm: [
            vm.push(eth_transfer_message.make()),
            vm.cast(eth_transfer_message.typ),
            eth_transfer_message.set_val("amount")(vm),
            eth_transfer_message.set_val("dest")(vm),
            # token_transfer_message sender
            vm.push(WITHDRAW_ETH_TYPECODE),
            vm.push(message.make()),
            vm.cast(message.typ),
            message.set_val("type")(vm),
            message.set_val("message")(vm),
            message.set_val("sender")(vm),
            vm.send(),
            vm.push(3),
        ],
        lambda vm: [vm.pop(), vm.pop(), vm.pop(), vm.push(0)],
    )


def withdraw_token_interrupt(vm, token_type):
    # local_exec_state
    parse_withdraw_call(vm)
    # amount dest token_address
    vm.push(token_transfer_message.make())
    vm.cast(token_transfer_message.typ)
    token_transfer_message.set_val("amount")(vm)
    token_transfer_message.set_val("dest")(vm)
    token_transfer_message.set_val("token_address")(vm)
    os.message_caller(vm)
    vm.push(token_type)
    vm.push(message.make())
    vm.cast(message.typ)
    message.set_val("type")(vm)
    message.set_val("sender")(vm)
    message.set_val("message")(vm)
    vm.send()
    vm.push(3)


def withdraw_erc20_interrupt(vm):
    # local_exec_state
    withdraw_token_interrupt(vm, WITHDRAW_ERC20_TYPECODE)


def withdraw_erc721_interrupt(vm):
    # local_exec_state
    withdraw_token_interrupt(vm, WITHDRAW_ERC721_TYPECODE)


def return_one_uint_to_solidity_caller(vm):
    vm.push(0)
    std.byterange.new(vm)
    std.byterange.set_val(vm)
    vm.push(32)
    vm.swap1()
    std.tup.make(2)(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.get("parent_frame")(vm)
    os.call_frame.call_frame.set_val("return_data")(vm)
    os.get_call_frame(vm)
    call_frame.call_frame.set_val("parent_frame")(vm)
    os.set_call_frame(vm)
    vm.push(2)


def arbsys_time_upper_bound(vm):
    vm.gettime()
    vm.tgetn(1)
    return_one_uint_to_solidity_caller(vm)


def arbsys_current_message_time(vm):
    os.get_chain_state(vm)
    os.chain_state.get("global_exec_state")(vm)
    os.global_exec_state.get("current_msg")(vm)
    ethbridge_message.get("block_number")(vm)

    return_one_uint_to_solidity_caller(vm)


def _perform_real_call(vm, call_num):
    os.get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)
    # balance dest local_exec_state
    vm.dup2()
    local_exec_state.get("value")(vm)
    std.comparison.lte(vm)
    vm.ifelse(lambda vm: [_execute_call(vm, call_num)], lambda vm: [vm.push(0)])


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

    _complete_call(vm, call_num)


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


@modifies_stack(
    [value.IntType(), value.TupleType([value.IntType()] * 7)], [value.IntType()]
)
def _mutable_call_ret(vm):
    # ret_type calltup
    translate_ret_type(vm)
    # return_val calltup
    vm.ifelse(
        lambda vm: [save_up_call_frame(vm), vm.push(1)],
        lambda vm: [clear_up_call_frame(vm), vm.push(0)],
    )
    vm.swap1()
    os.copy_return_data(vm)


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
    vm.set_exception_handler(invalid_tx)
    vm.dup0()
    setup_initial_call_frame(vm)
    os.tx_call_to_local_exec_state(vm)
    _perform_call(vm, label)
    # ret_code
    vm.clear_exception_handler()
    vm.auxpush()

    save_up_call_frame(vm)

    os.get_call_frame(vm)
    call_frame.save_state(vm)
    commit_call_frame(vm)
    vm.auxpop()
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
    _mutable_call_ret(vm)


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
    _mutable_call_ret(vm)


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
    _mutable_call_ret(vm)


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
