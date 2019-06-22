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
from . import keyvalue_int_int
from .stack_manip import dup_n
from .struct import Struct
from .. import value, std

# currencyStore keep track of balances of currencies
# implemented as a keyvalue store, currencyid to balance

non_fung_keyvalue = std.make_keyvalue_type(value.IntType(), keyvalue_int_int)

currency_store = Struct("currency_store", [
    ("fung", keyvalue_int_int.typ),
    ("non_fung", non_fung_keyvalue.typ)
])

typ = currency_store.typ


def make():
    return value.Tuple([keyvalue_int_int.make(), non_fung_keyvalue.make()])


@modifies_stack(0, [typ])
def new(vm):
    vm.push(make())
    vm.cast(typ)


@modifies_stack([typ, value.IntType()], [value.IntType()])
def get(vm):
    # cstore currId -> balance
    currency_store.get("fung")(vm)
    keyvalue_int_int.get(vm)
    # value


@modifies_stack(
    [typ, value.IntType(), value.IntType()],
    [typ]
)
def add(vm):
    # cstore currId delta -> updatedcstore
    vm.swap1()
    vm.swap2()
    # delta cstore currId
    vm.dup2()
    vm.dup2()
    get(vm)
    # oldval delta cstore currId
    vm.add()
    # newval cstore currId
    vm.swap2()
    vm.swap1()
    # cstore currId newval
    vm.swap2()
    vm.swap1()
    vm.dup2()
    currency_store.get("fung")(vm)
    keyvalue_int_int.set_val(vm)
    vm.swap1()
    currency_store.set_val("fung")(vm)
    # updatedcstore


@modifies_stack([
    typ,
    value.IntType(),
    value.IntType()
], [
    value.IntType(),
    typ
])
def deduct(vm):
    # cstore currId delta -> success updatedcstore
    vm.swap1()
    vm.swap2()
    vm.dup2()
    vm.dup2()
    get(vm)
    vm.swap1()
    # oldval delta cstore currId
    vm.dup1()
    vm.dup1()
    vm.gt()
    vm.iszero()
    vm.ifelse(lambda vm: [
        vm.add(),
        # newval cstore currId
        vm.swap2(),
        vm.swap1(),
        # cstore currId newval
        vm.swap2(),
        vm.swap1(),
        vm.dup2(),
        currency_store.get("fung")(vm),
        keyvalue_int_int.set_val(vm),
        vm.swap1(),
        currency_store.set_val("fung")(vm),
        # updatedcstore
        vm.push(1),
    ], lambda vm: [
        # oldval delta cstore currId
        vm.pop(),
        vm.pop(),
        vm.swap1(),
        vm.pop(),
        vm.push(0)
    ])
