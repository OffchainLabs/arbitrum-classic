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
from .. import value
from .. import std


@modifies_stack([value.IntType(), value.IntType()], [std.sized_byterange.typ])
def make_token_mint_message(vm):
    _setup_message(vm, "e58306f9")


@modifies_stack([value.IntType(), value.IntType()], [std.sized_byterange.typ])
def make_token_burn_message(vm):
    _setup_message(vm, "06dd0419")


def _setup_message(vm, func_code):
    # dest value
    vm.push(4)
    vm.push(std.byterange.frombytes(bytes.fromhex(func_code)))
    vm.cast(std.byterange.typ)
    std.byterange.set_val(vm)

    # br value
    vm.push(36)
    vm.swap1()
    std.byterange.set_val(vm)

    # br
    std.sized_byterange.new(vm)
    std.sized_byterange.sized_byterange.set_val("data")(vm)

    vm.push(68)
    vm.swap1()
    std.sized_byterange.sized_byterange.set_val("size")(vm)
    # sbr
