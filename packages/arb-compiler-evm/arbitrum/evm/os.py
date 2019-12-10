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

from .. import std
from ..std.struct import Struct

from ..annotation import modifies_stack
from ..vm import VM
from .. import value
from .types import ethbridge_message, message
from .types import tx_message, tx_call_data, local_exec_state
from .types import token_transfer_message, eth_transfer_message
from . import call_frame
from . import tokens
from . import accounts
from .accounts import account_store, account_state

# Blockchain Simulator
# inbox - global
# storage - global per contract
# backup_storage - global per contract
# balance - global per contract

# memory - per call frame
# message - per call frame
# return_data - per call frame
# logs - per call frame

global_exec_state = Struct(
    "global_execution_state",
    [
        ("origin", value.IntType()),
        ("block_number", value.IntType()),
        ("timestamp", value.IntType()),
        ("txhash", value.IntType()),
        ("current_msg", ethbridge_message.typ),
    ],
)

chain_state = Struct(
    "chain_state",
    [
        ("accounts", account_store.typ),
        ("inbox", std.inboxctx.typ),
        ("call_frame", call_frame.typ),
        ("sender_seq", std.keyvalue_int_int.typ),
        ("global_exec_state", global_exec_state.typ),
        "scratch",
    ],
)


ERC20_TRANSFER_TYPE = 1
ERC721_TRANSFER_TYPE = 2


def make_global_exec_state():
    vm = VM()
    vm.push(0)
    vm.push(0)
    vm.push(0)
    vm.push(0)
    tx_message.new(vm)

    global_exec_state.new(vm)
    global_exec_state.set_val("current_msg")(vm)
    global_exec_state.set_val("origin")(vm)
    global_exec_state.set_val("block_number")(vm)
    global_exec_state.set_val("timestamp")(vm)
    global_exec_state.set_val("txhash")(vm)
    return vm.stack.items[0]


@modifies_stack([ethbridge_message.typ, global_exec_state.typ], [global_exec_state.typ])
def update_execution_state(vm):
    # msg exec_state
    vm.dup0()
    vm.auxpush()

    vm.swap1()
    vm.dup0()
    global_exec_state.get("timestamp")(vm)
    vm.swap1()
    global_exec_state.get("block_number")(vm)
    vm.swap2()
    # msg old_timestamp old_block_number

    vm.dup0()
    ethbridge_message.get("timestamp")(vm)
    # timestamp msg old_timestamp old_block_number

    vm.swap1()
    vm.swap2()
    std.arith.max(vm)
    vm.swap2()
    # old_block_number msg timestamp

    vm.dup1()
    ethbridge_message.get("block_number")(vm)
    std.arith.max(vm)
    vm.swap1()
    # msg block_number timestamp
    ethbridge_message.get("txhash")(vm)
    vm.auxpop()
    vm.dup0()
    ethbridge_message.get("message")(vm)
    message.get("sender")(vm)

    # origin msg txhash block_number timestamp
    vm.push(global_exec_state.make())
    vm.cast(global_exec_state.typ)
    global_exec_state.set_val("origin")(vm)
    global_exec_state.set_val("current_msg")(vm)
    global_exec_state.set_val("txhash")(vm)
    global_exec_state.set_val("block_number")(vm)
    global_exec_state.set_val("timestamp")(vm)


@modifies_stack([], [chain_state.typ])
def get_chain_state(vm):
    vm.rpush()
    vm.cast(chain_state.typ)


@modifies_stack([chain_state.typ], [])
def set_chain_state(vm):
    vm.rset()


@modifies_stack([], [value.IntType()])
def message_origin(vm):
    get_chain_state(vm)
    chain_state.get("global_exec_state")(vm)
    global_exec_state.get("origin")(vm)


@modifies_stack(1, 0)
def set_scratch(vm):
    get_chain_state(vm)
    chain_state.set_val("scratch")(vm)
    set_chain_state(vm)


@modifies_stack(0, 1)
def get_scratch(vm):
    get_chain_state(vm)
    chain_state.get("scratch")(vm)


@modifies_stack([], [call_frame.typ])
def get_call_frame(vm):
    get_chain_state(vm)
    chain_state.get("call_frame")(vm)


@modifies_stack([call_frame.typ], [])
def set_call_frame(vm):
    get_chain_state(vm)
    chain_state.set_val("call_frame")(vm)
    set_chain_state(vm)


@modifies_stack([], [value.IntType()])
def get_timestamp(vm):
    get_chain_state(vm)
    chain_state.get("global_exec_state")(vm)
    global_exec_state.get("timestamp")(vm)


@modifies_stack([], [value.IntType()])
def get_block_number(vm):
    get_chain_state(vm)
    chain_state.get("global_exec_state")(vm)
    global_exec_state.get("block_number")(vm)


def create_initial_evm_state(contracts):
    vm = VM()

    std.keyvalue.new(vm)
    for contract in contracts:
        vm.push(contract["contractID"])

        for storage_item in contract["storage"]:
            vm.push(contract["storage"][storage_item])
            vm.push(storage_item)

        std.keyvalue.new(vm)
        for storage_item in contract["storage"]:
            std.keyvalue.set_val(vm)

        vm.push(contract["code_hash"])
        vm.push(contract["code_size"])
        vm.push(contract["code_point"])
        vm.push(contract["code"])
        vm.push(1)
        vm.push(accounts.make_empty_account())
        account_state.set_val("nonce")(vm)
        account_state.set_val("code")(vm)
        account_state.set_val("code_point")(vm)
        account_state.set_val("code_size")(vm)
        account_state.set_val("code_hash")(vm)
        account_state.set_val("storage")(vm)
        vm.swap2()
        std.keyvalue.set_val(vm)

    std.inboxctx.new(vm)
    std.keyvalue_int_int.new(vm)
    vm.push(make_global_exec_state())
    chain_state.new(vm)
    chain_state.set_val("global_exec_state")(vm)
    chain_state.set_val("sender_seq")(vm)
    chain_state.set_val("inbox")(vm)
    chain_state.set_val("accounts")(vm)
    return vm.stack.items[0]


def initialize(vm, contracts):
    vm.push(create_initial_evm_state(contracts))
    vm.rset()


def _set_call_frame_member_impl(vm, field):
    # memory
    get_chain_state(vm)
    chain_state.get("call_frame")(vm)
    call_frame.call_frame.set_val(field)(vm)
    get_chain_state(vm)
    chain_state.set_val("call_frame")(vm)
    set_chain_state(vm)


def _set_account_state_member_impl(vm, field):
    # val
    get_chain_state(vm)
    chain_state.get("call_frame")(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.set_val(field)(vm)
    get_chain_state(vm)
    chain_state.get("call_frame")(vm)
    call_frame.call_frame.set_val("account_state")(vm)
    get_chain_state(vm)
    chain_state.set_val("call_frame")(vm)
    set_chain_state(vm)


@modifies_stack([std.sized_byterange.sized_byterange.typ], 0)
def set_current_memory(vm):
    _set_call_frame_member_impl(vm, "memory")


@modifies_stack(std.stack.typ, 0)
def set_current_saved_stack(vm):
    _set_call_frame_member_impl(vm, "saved_stack")


@modifies_stack([std.stack_tup.typ], 0)
def set_current_logs(vm):
    _set_call_frame_member_impl(vm, "logs")


@modifies_stack([std.keyvalue_int_int.typ], 0)
def set_current_storage(vm):
    _set_account_state_member_impl(vm, "storage")


@modifies_stack([value.IntType()], 0)
def set_current_balance(vm):
    _set_account_state_member_impl(vm, "balance")


@modifies_stack(0, [std.byterange.typ])
def get_current_return_data_raw(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("return_data")(vm)
    std.sized_byterange.sized_byterange.get("data")(vm)


@modifies_stack(0, [value.IntType()])
def balance_get(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)


@modifies_stack(0, [value.IntType()])
def codesize_get(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("code_size")(vm)


@modifies_stack(0, [std.byterange.typ])
def code_get(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("code")(vm)


@modifies_stack([value.IntType()], [account_state.typ])
def get_ext_account(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    vm.dup1()
    vm.eq()
    vm.ifelse(
        lambda vm: [
            vm.pop(),
            get_call_frame(vm),
            call_frame.call_frame.get("account_state")(vm),
        ],
        lambda vm: [
            get_call_frame(vm),
            call_frame.call_frame.get("accounts")(vm),
            account_store.get(vm),
        ],
    )


@modifies_stack([value.IntType()], [value.IntType()])
def ext_balance(vm):
    # address
    get_ext_account(vm)
    account_state.get("balance")(vm)


@modifies_stack([], [std.byterange.typ])
def get_code_scratch(vm):
    get_scratch(vm)
    vm.cast(std.byterange.typ)


@modifies_stack([value.IntType()] * 4, [])
def ext_codecopy(vm):
    # address
    get_ext_account(vm)
    account_state.get("code")(vm)
    set_scratch(vm)
    evm_copy_to_memory(vm, get_code_scratch)


@modifies_stack([value.IntType()], [value.IntType()])
def ext_codesize(vm):
    # address
    get_ext_account(vm)
    account_state.get("code_size")(vm)


@modifies_stack([value.IntType()], [value.IntType()])
def ext_codehash(vm):
    # address
    get_ext_account(vm)
    vm.dup0()
    accounts.is_empty(vm)
    vm.ifelse(
        lambda vm: [vm.pop(), vm.push(0)],
        lambda vm: [account_state.get("code_hash")(vm)],
    )


@modifies_stack([], [std.sized_byterange.sized_byterange.typ])
def call_message_data(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("local_exec_state")(vm)
    local_exec_state.get("data")(vm)
    vm.cast(std.sized_byterange.sized_byterange.typ)


@modifies_stack(0, 1)
def message_data_size(vm):
    call_message_data(vm)
    std.sized_byterange.length(vm)


@modifies_stack([value.IntType()], [value.IntType()])
def message_data_load(vm):
    call_message_data(vm)
    std.sized_byterange.get(vm)


@modifies_stack(0, [std.byterange.typ])
def message_data_raw(vm):
    call_message_data(vm)
    std.sized_byterange.sized_byterange.get("data")(vm)


@modifies_stack([value.IntType(), value.IntType(), value.IntType()], 0)
def message_data_copy(vm):
    evm_copy_to_memory(vm, message_data_raw)


@modifies_stack(0, 1)
def message_value(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("local_exec_state")(vm)
    local_exec_state.get("value")(vm)


@modifies_stack(0, 1)
def message_caller(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("local_exec_state")(vm)
    local_exec_state.get("caller")(vm)


# [index]
@modifies_stack([value.IntType()], [value.IntType()])
def memory_load(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.get(vm)


# [] -> [int]
@modifies_stack(0, 1)
def memory_length(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.length(vm)


# [index, value]
@modifies_stack([value.IntType(), value.IntType()], [])
def memory_store(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.set_val(vm)
    set_current_memory(vm)


# [index, value]
@modifies_stack([value.IntType(), value.IntType()], [])
def memory_store8(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.set_val8(vm)
    set_current_memory(vm)


# # [index, value]
# @modifies_stack([value.IntType(), value.IntType()], [])
# def memory_store(vm):
#     get_call_frame(vm)
#     call_frame.call_frame.get("memory")(vm)
#     std.sized_byterange.set_val(vm)
#     set_current_memory(vm)

# [index]
@modifies_stack([value.IntType()], [value.IntType()])
def storage_load(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("storage")(vm)
    std.keyvalue_int_int.get(vm)


# [index, value]
@modifies_stack([value.IntType(), value.IntType()], [])
def storage_store(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("account_state")(vm)
    account_state.get("storage")(vm)
    std.keyvalue_int_int.set_val(vm)
    set_current_storage(vm)


# [destOffset, offset, length]
def evm_copy_to_memory(vm, source):
    vm.swap2()
    # [length, offset, destOffset]
    vm.dup1()
    vm.add()
    # [end offset, start offset, destOffset]
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.sized_byterange.get("data")(vm)
    # [memory, end offset, start offset, destOffset]
    vm.swap2()
    # [start offset, end offset, memory, destOffset]
    source(vm)
    # [code bytearray, start offset, end offset, memory, destOffset]
    std.byterange.copy(vm)
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.sized_byterange.get("size")(vm)
    vm.swap1()
    std.sized_byterange.new(vm)
    std.sized_byterange.sized_byterange.set_val("data")(vm)
    std.sized_byterange.sized_byterange.set_val("size")(vm)
    set_current_memory(vm)


@modifies_stack(0, 1)
def _emv_return_data(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("return_data")(vm)


@modifies_stack(0, [value.IntType()])
def return_data_size(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("return_data")(vm)
    std.sized_byterange.length(vm)


@modifies_stack([value.IntType(), value.IntType(), value.IntType()], 0)
def return_data_copy(vm):
    evm_copy_to_memory(vm, get_current_return_data_raw)


# [offset, length]
@modifies_stack([value.IntType(), value.IntType()], [std.byterange.typ])
def get_mem_segment(vm):
    vm.swap1()
    vm.dup1()
    vm.add()
    vm.swap1()
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.sized_byterange.get("data")(vm)
    std.byterange.get_subset(vm)


@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def evm_sha3(vm):
    vm.dup1()
    vm.push(32)
    vm.eq()
    vm.ifelse(
        lambda vm: [
            vm.swap1(),
            vm.pop(),
            get_call_frame(vm),
            call_frame.call_frame.get("memory")(vm),
            std.sized_byterange.sized_byterange.get("data")(vm),
            std.byterange.get(vm),
            vm.hash(),
        ],
        lambda vm: [
            # [pos, length]
            vm.dup1(),
            # [length, pos, length]
            vm.swap1(),
            get_mem_segment(vm),
            std.sha3.hash_byterange(vm),
        ],
    )


# @modifies_stack(2, 1)
# def evm_sha3(vm):
#     vm.dup1()
#     vm.swap1()
#     get_mem_segment(vm)
#     # bytearray length
#     std.sha3.hash_byterange(vm)


@modifies_stack([value.TupleType()], [])
def add_log(vm):
    get_call_frame(vm)
    call_frame.call_frame.get("logs")(vm)
    std.stack_tup.push(vm)
    set_current_logs(vm)


# [offset, length, topic0]
@modifies_stack([value.IntType()] * 3, 0)
def evm_log1(vm):
    vm.dup1()
    vm.swap1()
    get_mem_segment(vm)
    std.tup.make(2)(vm)
    get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    std.tup.make(3)(vm)
    add_log(vm)


# [offset, length, topic0, topic1]
@modifies_stack([value.IntType()] * 4, 0)
def evm_log2(vm):
    vm.dup1()
    vm.swap1()
    get_mem_segment(vm)
    std.tup.make(2)(vm)
    get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    std.tup.make(4)(vm)
    add_log(vm)


# [offset, length, topic0, topic1, topic2]
@modifies_stack([value.IntType()] * 5, 0)
def evm_log3(vm):
    vm.dup1()
    vm.swap1()
    get_mem_segment(vm)
    std.tup.make(2)(vm)
    get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    std.tup.make(5)(vm)
    add_log(vm)


# [offset, length, topic0, topic1, topic2]
@modifies_stack([value.IntType()] * 6, 0)
def evm_log4(vm):
    vm.dup1()
    vm.swap1()
    get_mem_segment(vm)
    std.tup.make(2)(vm)
    get_call_frame(vm)
    call_frame.call_frame.get("contractID")(vm)
    std.tup.make(6)(vm)
    add_log(vm)


# [sequence_num, sender] -> # [approved]
@modifies_stack([value.IntType()] * 2, [value.IntType()])
def check_message_sequence(vm):
    vm.dup1()
    get_chain_state(vm)
    chain_state.get("sender_seq")(vm)
    std.keyvalue_int_int.get(vm)
    # [current_seq, seq, sender]
    vm.swap1()
    vm.push(2)
    vm.dup1()
    vm.mod()
    # [seq % 2, seq, current_seq, sender]
    vm.swap1()
    vm.push(2)
    vm.swap1()
    vm.div()
    # [seq / 2, seq % 2, current_seq, sender]
    vm.swap2()
    vm.swap1()
    # [seq % 2, current_seq, seq / 2, sender]
    vm.ifelse(
        lambda vm: [
            # sequence must be incremented
            # [current_seq, seq / 2, sender]
            vm.push(1),
            vm.add(),
            vm.dup1(),
            vm.eq()
            # [seq / 2 == current_seq + 1, seq / 2, sender]
        ],
        lambda vm: [
            # sequence must be greater
            # [current_seq, seq / 2, sender]
            vm.dup1(),
            vm.gt()
            # [seq / 2 > current_seq, seq / 2, sender]
        ],
    )

    # [seq_should_update, seq / 2, sender]
    vm.ifelse(
        lambda vm: [
            # [seq / 2, sender]
            get_chain_state(vm),
            chain_state.get("sender_seq")(vm),
            std.keyvalue_int_int.set_val(vm),
            get_chain_state(vm),
            chain_state.set_val("sender_seq")(vm),
            set_chain_state(vm),
            vm.push(1),
        ],
        lambda vm: [
            std.sized_byterange.new(vm),
            vm.push(4),
            log_func_result(vm),
            vm.pop(),
            vm.pop(),
            vm.push(0),
        ],
    )


@modifies_stack(0, [message.typ])
def get_next_message(vm):
    get_chain_state(vm)
    chain_state.get("inbox")(vm)
    std.inboxctx.getmsg(vm)
    # msg updatedctx
    vm.swap1()
    get_chain_state(vm)
    chain_state.set_val("inbox")(vm)
    set_chain_state(vm)
    vm.cast(ethbridge_message.typ)

    # ethbridge_message
    get_chain_state(vm)
    chain_state.get("global_exec_state")(vm)
    vm.dup1()
    update_execution_state(vm)
    get_chain_state(vm)
    chain_state.set_val("global_exec_state")(vm)
    set_chain_state(vm)
    ethbridge_message.get("message")(vm)


@modifies_stack(
    [value.IntType(), tx_call_data.typ], [value.IntType(), local_exec_state.typ]
)
def tx_call_to_local_exec_state(vm):
    # caller tx_call
    vm.swap1()
    vm.dup0()
    tx_call_data.get("dest")(vm)
    # dest tx_call caller
    vm.swap2()
    vm.swap1()
    vm.dup0()
    tx_call_data.get("value")(vm)
    vm.swap1()
    tx_call_data.get("data")(vm)
    # data value caller dest
    vm.push(local_exec_state.make())
    vm.cast(local_exec_state.typ)
    local_exec_state.set_val("data")(vm)
    local_exec_state.set_val("value")(vm)
    local_exec_state.set_val("caller")(vm)
    vm.swap1()
    # dest local_exec_state


@modifies_stack([message.typ], [value.TupleType()])
def process_tx_message(vm):
    # msg
    vm.dup0()
    message.get("sender")(vm)
    vm.dup1()
    message.get("message")(vm)
    vm.cast(tx_message.typ)
    tx_message.get("sequence_num")(vm)
    # sequence_num sender message
    check_message_sequence(vm)
    # valid_seq msg
    vm.ifelse(
        lambda vm: [
            vm.dup0(),
            message.get("message")(vm),
            vm.cast(tx_message.typ),
            process_valid_tx_message(vm),
            vm.swap1(),
            message.get("sender")(vm),
            std.tup.make(2)(vm),
            # sender tx_call_data
        ],
        lambda vm: [vm.pop(), vm.push(value.Tuple([]))],
    )


@modifies_stack([tx_message.typ], [tx_call_data.typ])
def process_valid_tx_message(vm):
    # msg
    vm.dup0()
    tx_message.get("to")(vm)
    # dest message
    vm.dup1()
    tx_message.get("value")(vm)
    # value dest message
    vm.swap2()
    tx_message.get("data")(vm)
    # data dest value

    vm.push(tx_call_data.make())
    vm.cast(tx_call_data.typ)
    tx_call_data.set_val("data")(vm)
    tx_call_data.set_val("dest")(vm)
    tx_call_data.set_val("value")(vm)


@modifies_stack([message.typ], [])
def process_deposit_eth_message(vm):
    message.get("message")(vm)
    vm.cast(eth_transfer_message.typ)

    vm.dup0()
    eth_transfer_message.get("amount")(vm)
    vm.dup1()
    eth_transfer_message.get("dest")(vm)

    get_chain_state(vm)
    chain_state.get("accounts")(vm)
    account_store.get(vm)
    # account amount eth_transfer_message
    vm.swap1()
    vm.dup1()
    account_state.get("balance")(vm)
    vm.add()
    # new_balance account eth_transfer_message
    vm.swap1()
    account_state.set_val("balance")(vm)
    # account eth_transfer_message

    vm.swap1()
    eth_transfer_message.get("dest")(vm)
    get_chain_state(vm)
    chain_state.get("accounts")(vm)
    account_store.set_val(vm)
    get_chain_state(vm)
    chain_state.set_val("accounts")(vm)
    set_chain_state(vm)

    std.sized_byterange.new(vm)
    vm.push(6)
    log_func_result(vm)


@modifies_stack([message.typ], [value.IntType()])
def process_withdraw_eth_message(vm):
    message.get("message")(vm)
    vm.cast(eth_transfer_message.typ)

    vm.dup0()
    eth_transfer_message.get("amount")(vm)
    vm.dup1()
    eth_transfer_message.get("dest")(vm)

    get_chain_state(vm)
    chain_state.get("accounts")(vm)
    account_store.get(vm)
    # account amount eth_transfer_message
    vm.swap1()
    vm.dup1()
    account_state.get("balance")(vm)
    # balance amount eth_transfer_message
    vm.dup1()
    vm.dup1()
    std.comparison.gte(vm)
    vm.ifelse(
        lambda vm: [
            vm.sub(),
            # new_balance account eth_transfer_message
            vm.swap1(),
            account_state.set_val("balance")(vm),
            # account eth_transfer_message
            vm.swap1(),
            eth_transfer_message.get("dest")(vm),
            get_chain_state(vm),
            chain_state.get("accounts")(vm),
            account_store.set_val(vm),
            get_chain_state(vm),
            chain_state.set_val("accounts")(vm),
            set_chain_state(vm),
            std.sized_byterange.new(vm),
            vm.push(7),
            log_func_result(vm),
            vm.push(1),
        ],
        lambda vm: [
            vm.pop(),
            vm.pop(),
            vm.pop(),
            vm.pop(),
            std.sized_byterange.new(vm),
            vm.push(5),
            log_func_result(vm),
            vm.push(0),
        ],
    )


@modifies_stack([message.typ], [tx_call_data.typ, value.IntType()])
def process_deposit_erc20_message(vm):
    process_deposit_token_message(vm, accounts.create_erc20)


@modifies_stack([message.typ], [tx_call_data.typ, value.IntType()])
def process_deposit_erc721_message(vm):
    process_deposit_token_message(vm, accounts.create_erc721)


def process_deposit_token_message(vm, account_create_func):
    # message
    vm.dup0()
    message.get("message")(vm)
    vm.cast(token_transfer_message.typ)
    token_transfer_message.get("token_address")(vm)
    get_chain_state(vm)
    chain_state.get("accounts")(vm)
    account_create_func(vm)
    get_chain_state(vm)
    chain_state.set_val("accounts")(vm)
    set_chain_state(vm)

    # message
    message.get("message")(vm)
    vm.cast(token_transfer_message.typ)
    # token_transfer_message
    vm.dup0()
    token_transfer_message.get("amount")(vm)
    # amount token_transfer_message
    vm.dup1()
    token_transfer_message.get("dest")(vm)
    # address amount token_transfer_message
    tokens.make_token_mint_message(vm)
    # data token_transfer_message
    vm.swap1()
    token_transfer_message.get("token_address")(vm)
    # token_address data
    vm.push(0)
    vm.push(tx_call_data.make())
    vm.cast(tx_call_data.typ)
    tx_call_data.set_val("value")(vm)
    tx_call_data.set_val("dest")(vm)
    tx_call_data.set_val("data")(vm)

    # sender is 1 for admin call
    vm.push(1)
    vm.swap1()
    # tx_call_data sender


@modifies_stack([message.typ], [tx_call_data.typ, value.IntType()])
def process_withdraw_erc20_message(vm):
    process_withdraw_token_message(vm, accounts.create_erc20)


@modifies_stack([message.typ], [tx_call_data.typ, value.IntType()])
def process_withdraw_erc721_message(vm):
    process_withdraw_token_message(vm, accounts.create_erc721)


def process_withdraw_token_message(vm, account_create_func):
    # message
    vm.dup0()
    message.get("message")(vm)
    vm.cast(token_transfer_message.typ)
    token_transfer_message.get("token_address")(vm)
    get_chain_state(vm)
    chain_state.get("accounts")(vm)
    account_create_func(vm)
    get_chain_state(vm)
    chain_state.set_val("accounts")(vm)
    set_chain_state(vm)

    # message
    vm.dup0()
    message.get("sender")(vm)
    vm.swap1()
    message.get("message")(vm)
    vm.cast(token_transfer_message.typ)
    # token_transfer_message sender
    vm.dup0()
    token_transfer_message.get("amount")(vm)
    # amount token_transfer_message sender
    vm.swap1()
    vm.swap2()
    tokens.make_token_burn_message(vm)
    # data token_transfer_message

    vm.swap1()
    token_transfer_message.get("token_address")(vm)

    # token_address data
    vm.push(0)
    vm.push(tx_call_data.make())
    vm.cast(tx_call_data.typ)
    tx_call_data.set_val("value")(vm)
    tx_call_data.set_val("dest")(vm)
    tx_call_data.set_val("data")(vm)
    # tx_call_data
    vm.push(1)
    vm.swap1()
    # tx_call_data sender


# [code, data]
@modifies_stack(2, 0)
def log_func_result(vm):
    vm.swap1()
    # [data, code]
    get_call_frame(vm)
    call_frame.call_frame.get("logs")(vm)
    std.stack_tup.new(vm)
    set_current_logs(vm)

    get_chain_state(vm)
    chain_state.get("global_exec_state")(vm)
    global_exec_state.get("current_msg")(vm)

    # [msg, logs, data, code]
    std.tup.make(4)(vm)
    vm.log()


@modifies_stack([value.TupleType([value.IntType()] * 7)], 0)
def copy_return_data(vm):
    vm.dup0()
    vm.tgetn(6)
    vm.swap1()
    vm.tgetn(5)
    # [ret_offset, ret_length]
    vm.push(0)
    vm.swap1()
    # [destOffset, offset, length]
    evm_copy_to_memory(vm, get_current_return_data_raw)


# [[gas, dest, value, arg offset, arg length, ret offset, ret length]]
# send tuple is [data, dest, value, kind]
@modifies_stack([value.TupleType([value.IntType()] * 7)], [tx_call_data.typ])
def evm_call_to_tx_call_data(vm):
    vm.dup0()
    vm.tgetn(2)
    vm.swap1()
    vm.dup0()
    vm.tgetn(1)
    vm.swap1()
    # [tup]
    vm.dup0()
    vm.tgetn(3)
    vm.swap1()
    vm.dup0()
    vm.tgetn(4)
    # [arg length, tup, arg offset]
    vm.dup2()
    vm.add()
    vm.swap1()
    vm.swap2()
    # [arg start, arg end, tup]
    get_call_frame(vm)
    call_frame.call_frame.get("memory")(vm)
    std.sized_byterange.sized_byterange.get("data")(vm)
    std.byterange.get_subset(vm)
    # [ba, tup]
    vm.swap1()
    vm.tgetn(4)
    # [length, ba]
    vm.swap1()
    std.tup.make(2)(vm)
    # [sized byte array, tup]
    vm.push(tx_call_data.make())
    vm.cast(tx_call_data.typ)
    tx_call_data.set_val("data")(vm)
    tx_call_data.set_val("dest")(vm)
    tx_call_data.set_val("value")(vm)
