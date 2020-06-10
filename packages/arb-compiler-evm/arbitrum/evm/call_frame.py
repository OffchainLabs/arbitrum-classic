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
from .. import value
from .types import local_exec_state
from .accounts import account_state
from . import accounts
from ..vm import VM
from ..annotation import modifies_stack

call_frame = std.struct.Struct(
    "call_frame",
    [
        ("contractID", value.IntType()),  # transient
        ("memory", std.sized_byterange.sized_byterange.typ),  # transient
        ("account_state", accounts.account_state.typ),  # record
        ("accounts", accounts.account_store.typ),  # record
        ("local_exec_state", local_exec_state.typ),  # transient
        ("return_data", std.sized_byterange.sized_byterange.typ),  # transient
        ("logs", std.stack_tup.typ),  # record
        "parent_frame",
        ("return_location", value.CodePointType()),
        ("saved_stack", std.stack.typ),
        ("saved_aux_stack", std.stack.typ),
    ],
)

typ = call_frame.typ

call_frame.update_type("parent_frame", call_frame.typ)


def make_empty():
    vm = VM()
    std.sized_byterange.new(vm)
    std.stack.new(vm)
    std.sized_byterange.new(vm)
    vm.push(value.Tuple())
    call_frame.new(vm)
    call_frame.set_val("parent_frame")(vm)
    call_frame.set_val("memory")(vm)
    call_frame.set_val("logs")(vm)
    call_frame.set_val("return_data")(vm)
    return vm.stack.items[0]


@modifies_stack([call_frame.typ], [accounts.account_state.typ])
def lookup_current_state(vm):
    vm.dup0()
    call_frame.get("contractID")(vm)
    vm.swap1()
    call_frame.get("accounts")(vm)
    accounts.account_store.get(vm)


@modifies_stack([call_frame.typ], [call_frame.typ])
def save_state(vm):
    # frame
    vm.dup0()
    call_frame.get("account_state")(vm)
    vm.dup1()
    call_frame.get("contractID")(vm)
    vm.dup2()
    call_frame.get("accounts")(vm)
    accounts.account_store.set_val(vm)
    vm.swap1()
    call_frame.set_val("accounts")(vm)


@modifies_stack([call_frame.typ], [call_frame.typ])
def setup_state(vm):
    # frame
    vm.dup0()
    lookup_current_state(vm)
    vm.swap1()
    call_frame.set_val("account_state")(vm)


@modifies_stack([call_frame.typ, call_frame.typ], [call_frame.typ])
def merge(vm):
    # parent_frame current_frame
    vm.swap1()
    save_state(vm)
    vm.swap1()
    # parent_frame current_frame
    vm.dup1()
    call_frame.get("accounts")(vm)
    vm.swap1()
    call_frame.set_val("accounts")(vm)
    # parent_frame current_frame
    vm.dup1()
    call_frame.get("logs")(vm)
    vm.swap1()
    call_frame.set_val("logs")(vm)
    # parent_frame current_frame
    vm.swap1()
    vm.pop()
    # parent_frame


@modifies_stack([call_frame.typ], [call_frame.typ])
def spawn_child(vm):
    vm.dup0()
    call_frame.set_val("parent_frame")(vm)
    std.sized_byterange.new(vm)
    vm.swap1()
    call_frame.set_val("memory")(vm)


# update:
#   contractID
#   message
#   memory
#   storage
#   wallet
# maintain:
#   accounts
#   logs
@modifies_stack(
    [call_frame.typ, local_exec_state.typ, value.IntType(), value.CodePointType()],
    [call_frame.typ],
)
def spawn(vm):
    # parent_frame local_exec_state contractID ret_pc
    spawn_child(vm)

    # subtract sent funds from balance
    # frame local_exec_state contractID ret_pc
    vm.dup1()
    local_exec_state.get("value")(vm)
    # value frame local_exec_state contractID ret_pc
    vm.dup1()
    call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)
    vm.sub()
    # new_balance frame local_exec_state contractID ret_pc
    update_frame_balance(vm)
    save_state(vm)

    # frame local_exec_state contractID ret_pc
    call_frame.set_val("local_exec_state")(vm)
    call_frame.set_val("contractID")(vm)
    call_frame.set_val("return_location")(vm)
    setup_state(vm)
    # frame

    # add received funds to balance
    vm.dup0()
    call_frame.get("local_exec_state")(vm)
    local_exec_state.get("value")(vm)
    # value frame
    vm.dup1()
    call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)
    vm.add()
    # new_balance frame
    update_frame_balance(vm)
    # frame


# update:
#   contractID
#   message
#   memory
#   storage
#   wallet
# maintain:
#   accounts
#   logs
@modifies_stack(
    [call_frame.typ, local_exec_state.typ, value.IntType(), value.CodePointType()],
    [call_frame.typ],
)
def spawn_callcode(vm):
    # parent_frame local_exec_state contractID ret_pc
    spawn_child(vm)

    # subtract sent funds from balance
    # frame local_exec_state contractID ret_pc
    vm.dup1()
    local_exec_state.get("value")(vm)
    # value frame local_exec_state contractID ret_pc
    vm.dup1()
    call_frame.get("account_state")(vm)
    account_state.get("balance")(vm)
    vm.sub()
    # new_balance frame local_exec_state contractID ret_pc
    update_frame_balance(vm)
    save_state(vm)

    # frame local_exec_state contractID ret_pc
    call_frame.set_val("local_exec_state")(vm)
    vm.swap1()
    vm.swap2()
    vm.swap1()
    call_frame.set_val("return_location")(vm)
    # frame contractID
    vm.dup1()
    vm.dup1()
    call_frame.get("accounts")(vm)
    accounts.account_store.get(vm)
    # account frame contractID
    vm.dup1()
    call_frame.get("local_exec_state")(vm)
    local_exec_state.get("value")(vm)
    # message_val account frame contractID
    vm.dup1()
    account_state.get("balance")(vm)
    vm.add()
    # new_balance account frame contractID
    vm.swap1()
    account_state.set_val("balance")(vm)
    # new_account frame contractID
    vm.swap1()
    vm.swap2()
    vm.dup2()
    # frame contractID new_account frame
    call_frame.get("accounts")(vm)
    accounts.account_store.set_val(vm)
    # new_accounts frame
    vm.swap1()
    call_frame.set_val("accounts")(vm)
    setup_state(vm)
    # frame


# update:
#   message
#   memory
# maintain:
#   contractID
#   storage
#   balance
#   accounts
#   logs
@modifies_stack(
    [call_frame.typ, local_exec_state.typ, value.CodePointType()], [call_frame.typ]
)
def spawn_delegatecall(vm):
    # parent_frame local_exec_state ret_pc
    spawn_child(vm)
    # frame local_exec_state ret_pc

    # frame local_exec_state ret_pc
    call_frame.set_val("local_exec_state")(vm)
    call_frame.set_val("return_location")(vm)
    # frame


@modifies_stack([value.IntType(), call_frame.typ], [call_frame.typ])
def update_frame_balance(vm):
    vm.dup1()
    call_frame.get("account_state")(vm)
    account_state.set_val("balance")(vm)
    vm.swap1()
    call_frame.set_val("account_state")(vm)


@modifies_stack([accounts.account_store.typ], [call_frame.typ])
def new_fresh(vm):
    # accounts
    vm.push(make_empty())
    vm.cast(call_frame.typ)
    call_frame.set_val("accounts")(vm)
