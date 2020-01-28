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
from .. import std
from .. import value
from ..annotation import modifies_stack
from . import contract_templates
from ..vm import VM

account_state = std.Struct(
    "account_state",
    [
        ("nonce", value.IntType()),
        ("code", std.byterange.typ),
        ("code_point", value.ValueType()),
        ("code_size", value.IntType()),
        ("code_hash", value.IntType()),
        ("storage", std.keyvalue_int_int.typ),
        ("balance", value.IntType()),
    ],
)

EMPTY_HASH_STRING = "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"


def make_empty_account():
    vm = VM()
    vm.push(0)
    vm.push(eth_utils.to_int(hexstr=EMPTY_HASH_STRING))
    vm.push(0)
    std.keyvalue_int_int.new(vm)
    vm.push(0)
    account_state.new(vm)
    account_state.set_val("nonce")(vm)
    account_state.set_val("storage")(vm)
    account_state.set_val("balance")(vm)
    account_state.set_val("code_hash")(vm)
    account_state.set_val("code_size")(vm)

    return vm.stack[0]


@modifies_stack([account_state.typ], [value.IntType()])
def is_empty(vm):
    vm.dup0()
    account_state.get("nonce")(vm)
    vm.push(0)
    vm.eq()
    vm.swap1()
    vm.dup0()
    account_state.get("balance")(vm)
    vm.push(0)
    vm.eq()
    vm.swap1()
    account_state.get("code_size")(vm)
    vm.push(0)
    vm.eq()
    vm.bitwise_and()
    vm.bitwise_and()


account_store = std.make_keyvalue_type(
    value.IntType(), account_state.typ, make_empty_account()
)


@modifies_stack([account_store.typ, value.IntType()], [account_store.typ])
def create_erc20(vm):
    # accounts contract_id
    vm.dup1()
    vm.dup1()
    account_store.get(vm)
    is_empty(vm)
    vm.ifelse(
        lambda vm: [
            vm.push(contract_templates.ERC20_ADDRESS),
            vm.swap1(),
            clone_contract(vm),
        ],
        lambda vm: [vm.swap1(), vm.pop()],
    )


@modifies_stack([account_store.typ, value.IntType()], [account_store.typ])
def create_erc721(vm):
    # accounts contract_id
    vm.dup1()
    vm.dup1()
    account_store.get(vm)
    is_empty(vm)
    vm.ifelse(
        lambda vm: [
            vm.push(contract_templates.ERC721_ADDRESS),
            vm.swap1(),
            clone_contract(vm),
        ],
        lambda vm: [vm.swap1(), vm.pop()],
    )


@modifies_stack(
    [account_store.typ, value.IntType(), value.IntType()], [account_store.typ]
)
def clone_contract(vm):
    # accounts from_id to_id
    vm.swap1()
    vm.dup1()
    account_store.get(vm)
    vm.push(0)
    vm.swap1()
    account_state.set_val("balance")(vm)
    std.keyvalue_int_int.new(vm)
    vm.swap1()
    account_state.set_val("storage")(vm)
    vm.push(1)
    vm.swap1()
    account_state.set_val("nonce")(vm)
    # new_account accounts to_id
    vm.swap2()
    vm.swap1()
    account_store.set_val(vm)


@modifies_stack(
    [account_store.typ, value.IntType(), value.IntType()],
    [value.IntType(), value.IntType(), account_store.typ],
)
def process_nonce(vm):
    # accounts address nonce
    vm.dup1()
    vm.dup1()
    account_store.get(vm)
    # account accounts address nonce
    vm.swap2()
    vm.auxpush()
    vm.auxpush()
    # account nonce [accounts address]
    vm.dup0()
    vm.auxpush()
    account_state.get("nonce")(vm)
    # old_nonce nonce [account accounts address]
    vm.dup1()
    vm.eq()
    # valid_nonce nonce [account accounts address]
    vm.ifelse(
        lambda vm: [
            vm.push(1),
            vm.add(),
            vm.auxpop(),
            account_state.set_val("nonce")(vm),
            # updated_account [accounts address]
            vm.auxpop(),
            vm.auxpop(),
            vm.swap1(),
            account_store.set_val(vm),
            vm.push(0),
            vm.push(1),
        ],
        lambda vm: [
            # nonce [account accounts address]
            vm.auxpop(),
            vm.pop(),
            vm.auxpop(),
            vm.auxpop(),
            vm.pop(),
            vm.swap1(),
            # nonce accounts
            vm.push(0),
        ],
    )
