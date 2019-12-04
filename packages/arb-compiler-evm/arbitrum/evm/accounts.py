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
from ..annotation import modifies_stack

account_state = std.Struct(
    "account_state",
    [
        ("code_point", value.CodePointType()),
        ("storage", std.keyvalue_int_int.typ),
        ("wallet", std.currency_store.typ),
    ],
)

account_store = std.make_keyvalue_type(value.IntType(), account_state.typ)


@modifies_stack([value.CodePointType()], [account_state.typ])
def create_account(vm):
    std.currency_store.new(vm)
    std.keyvalue.new(vm)

    account_state.new(vm)
    account_state.set_val("storage")(vm)
    account_state.set_val("wallet")(vm)
    account_state.set_val("code_point")(vm)


@modifies_stack(
    [account_store.typ, value.IntType(), value.IntType()], [account_store.typ]
)
def clone_contract(vm):
    # accounts from_id to_id
    vm.swap1()
    vm.dup1()
    account_store.get(vm)
    account_state.get("code_point")(vm)
    create_account(vm)
    # new_account accounts to_id
    vm.swap2()
    vm.swap1()
    account_store.set_val(vm)
