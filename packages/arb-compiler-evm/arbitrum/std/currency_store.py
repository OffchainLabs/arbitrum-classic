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
from .struct import Struct
from .. import value, std

# currencyStore keep track of balances of currencies
# implemented as a keyvalue store, currencyid to balance

non_fung_keyvalue = std.make_keyvalue_type(value.IntType(), keyvalue_int_int.typ)

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
def get_fung(vm):
    # cstore currId -> balance
    currency_store.get("fung")(vm)
    keyvalue_int_int.get(vm)
    # value

@modifies_stack([typ, value.IntType()], [keyvalue_int_int.typ])
def get_non_fung(vm):
    # cstore currId -> balance
    currency_store.get("non_fung")(vm)
    non_fung_keyvalue.get(vm)
    # value

@modifies_stack(
    [typ, value.IntType(), value.IntType()],
    [typ]
)
def add(vm):
    # cstore currID amount
    vm.dup1()
    vm.push(31)
    vm.byte()
    vm.push(0)
    vm.eq()
    vm.ifelse(lambda vm: [
        add_erc20(vm)
    ], lambda vm: [
        add_erc721(vm)
    ])

@modifies_stack(
    [typ, value.IntType(), value.IntType()],
    [typ]
)
def add_erc20(vm):
    # cstore currId delta -> updatedcstore
    vm.swap1()
    vm.swap2()
    # delta cstore currId
    vm.dup2()
    vm.dup2()
    get_fung(vm)
    # oldval delta cstore currId
    vm.add()
    # newval cstore currId
    vm.swap1()
    vm.swap2()
    # currId newval cstore
    vm.dup2()
    currency_store.get("fung")(vm)
    keyvalue_int_int.set_val(vm)
    vm.swap1()
    currency_store.set_val("fung")(vm)
    # updatedcstore


@modifies_stack(
    [typ, value.IntType(), value.IntType()],
    [typ]
)
def add_erc721(vm):
    # cstore currId coinid -> updatedcstore
    vm.swap1()
    vm.swap2()
    # coinid cstore currId
    vm.dup2()
    vm.dup2()
    get_non_fung(vm)
    # nfkeyval coinid cstore currId
    vm.push(1)
    vm.swap2()
    vm.swap1()
    keyvalue_int_int.set_val(vm)
    # nfkeyval cstore currId
    vm.swap1()
    vm.swap2()
    # currId nfkeyval cstore
    vm.dup2()
    # cstore currId nfkeyval cstore
    currency_store.get("non_fung")(vm)
    non_fung_keyvalue.set_val(vm)
    # non_fung cstore
    vm.swap1()
    currency_store.set_val("non_fung")(vm)
    # updatedcstore

@modifies_stack(
    [typ, value.IntType(), value.IntType()],
    [value.IntType(), typ]
)
def deduct(vm):
    vm.dup1()
    vm.push(31)
    vm.byte()
    vm.push(0)
    vm.eq()
    vm.ifelse(lambda vm: [
        deduct_erc20(vm)
    ], lambda vm: [
        deduct_erc721(vm)
    ])

@modifies_stack([
    typ,
    value.IntType(),
    value.IntType()
], [
    value.IntType(),
    typ
])
def deduct_erc20(vm):
    # cstore currId delta -> success updatedcstore
    vm.swap1()
    vm.swap2()
    vm.dup2()
    vm.dup2()
    get_fung(vm)
    # balance delta cstore currId
    vm.dup1()
    vm.dup1()
    vm.lt()
    vm.ifelse(lambda vm: [
        # balance delta cstore currId
        vm.pop(),
        vm.pop(),
        vm.swap1(),
        vm.pop(),
        vm.push(0)
    ], lambda vm: [
        # balance delta cstore currId
        vm.sub(),
        # balance cstore currId
        vm.swap1(),
        vm.swap2(),
        vm.dup2(),
        # cstore currId balance cstore
        currency_store.get("fung")(vm),
        keyvalue_int_int.set_val(vm),
        vm.swap1(),
        currency_store.set_val("fung")(vm),
        # updatedcstore
        vm.push(1)
    ])


@modifies_stack([
    typ,
    value.IntType(),
    value.IntType()
], [
    value.IntType(),
    typ
])
def deduct_erc721(vm):
    # cstore currId coinid -> success updatedcstore
    vm.swap1()
    vm.swap2()
    vm.dup2()
    vm.dup2()
    get_non_fung(vm)
    # nfkeyval coinid cstore currId
    vm.dup1()
    vm.dup1()
    # nfkeyval coinid nfkeyval coinid cstore currId
    keyvalue_int_int.get(vm)
    # has_coin nfkeyval coinid cstore currId
    vm.ifelse(lambda vm: [
        # nfkeyval coinid cstore currId
        vm.push(0),
        vm.swap2(),
        vm.swap1(),
        # nfkeyval coinid 0 cstore currId
        keyvalue_int_int.set_val(vm),
        # nfkeyval cstore currId
        vm.swap1(),
        vm.swap2(),
        vm.dup2(),
        # cstore currId nfkeyval cstore
        currency_store.get("non_fung")(vm),
        non_fung_keyvalue.set_val(vm),
        vm.swap1(),
        currency_store.set_val("non_fung")(vm),
        vm.push(1)
    ], lambda vm: [
        # nfkeyval coinid cstore currId
        vm.pop(),
        vm.pop(),
        vm.swap1(),
        vm.pop(),
        vm.push(0)
    ])
