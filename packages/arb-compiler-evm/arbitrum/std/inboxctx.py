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
from . import queue_tup
from .. import value

typ = queue_tup.typ


@modifies_stack(0, [typ])
def new(vm):
    # -> q
    queue_tup.new(vm)


@modifies_stack([typ], [value.IntType()])
def isempty(vm):
    # q -> isempty
    queue_tup.isempty(vm)


@modifies_stack([typ], [value.ValueType(), typ])
def getmsg(vm):
    # q -> msg updatedq
    vm.dup0()
    queue_tup.isempty(vm)
    vm.ifelse(
        lambda vm: [
            # q
            vm.inbox(),
            # inbox q
            _inhale2(vm),
            # updatedq
        ]
    )
    # nonemptyq
    queue_tup.get(vm)


@modifies_stack([value.TupleType(), typ], [typ])
def _inhale2(vm):
    # inbox q -> updatedq
    vm.dup0()
    vm.tnewn(0)
    vm.eq()
    vm.ifelse(
        lambda vm: [vm.pop()],
        lambda vm: [
            # inbox q
            vm.cast(value.TupleType([value.TupleType(), value.TupleType()])),
            vm.dup0(),
            vm.tgetn(1),
            vm.auxpush(),
            vm.tgetn(0),
            # inbox q
            _inhale2(vm),
            # q
            vm.auxpop(),
            vm.swap1(),
            # q msg
            queue_tup.put(vm),
        ],
    )
