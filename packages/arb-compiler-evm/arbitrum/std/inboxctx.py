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
from . import tup
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
            vm.push((1 << 256) - 1),
            vm.inbox(),
            # inbox q
            _inhale2(vm),
            # updatedq
        ]
    )
    # nonemptyq
    queue_tup.get(vm)


# @modifies_stack([typ], [value.ValueType(), typ])
# def getmsg(vm):
#     # assume inboxctx is not empty
#     # ctx -> msg updatedctx
#     inhale(vm)
#     # ctx
#     vm.dup0()
#     inbox_ctx.get("queue")(vm)
#     # queue ctx
#     queue_tup.get(vm)
#     # msg updatedq ctx
#     vm.swap2()
#     # ctx updatedq msg
#     inbox_ctx.set_val("queue")(vm)
#     # updatedctx msg
#     vm.swap1()


# @modifies_stack([typ], [typ])
# def inhale(vm):
# q -> updatedq
#   vm.dup0()
#   inbox_ctx.get("queue")(vm)
#   vm.swap1()
#   inbox_ctx.get("nodeAlreadySeen")(vm)
#   vm.tnewn(8)
#   vm.inbox()
#   vm.swap2()
#    vm.swap1()
#   vm.dup2()
# inbox nodeAlreadySeen queue inbox
#    _inhale2(vm)
# nodeAlreadySeen updatedq inbox
#    vm.pop()
#    inbox_ctx.new(vm)
#    inbox_ctx.set_val("queue")(vm)
#   inbox_ctx.set_val("nodeAlreadySeen")(vm)


@modifies_stack([value.TupleType(), typ], [typ])
def _inhale2(vm):
    # inbox q -> updatedq
    vm.dup0()
    vm.tnewn(0)
    vm.eq()
    vm.ifelse(
        lambda vm: [vm.pop()],
        lambda vm: [
            vm.cast(
                value.TupleType([value.IntType(), value.TupleType(), value.TupleType()])
            ),
            # inbox q
            tup.tbreak(3)(vm),
            vm.push(0),
            vm.eq(),
            vm.ifelse(_inhale_message, _inhale_messages),
        ],
    )


@modifies_stack([value.TupleType(), value.TupleType(), typ], [typ])
def _inhale_message(vm):
    # inbox msg q
    vm.swap1()
    vm.auxpush()
    # inbox q
    _inhale2(vm)
    # q
    vm.auxpop()
    vm.swap1()
    # q msg
    queue_tup.put(vm)


@modifies_stack([value.TupleType(), value.TupleType(), typ], [typ])
def _inhale_messages(vm):
    # inboxA inboxB q
    vm.swap1()
    vm.auxpush()
    # inboxA q
    _inhale2(vm)
    # q
    vm.auxpop()
    _inhale2(vm)
