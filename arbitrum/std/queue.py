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
from . import make_boundedq_type
from .. import value
from .struct import Struct


# a queue is just a boundedq; when the boundedq gets full
# we generate a new, bigger one

def make_queue_type(typ):
    boundedq = make_boundedq_type(typ)
    queue_type = Struct(f"queue[{typ}]", [
        ("boundedq", boundedq.typ)
    ])
    # stack_type.fields[]

    class Queue:
        @staticmethod
        def make():
            return boundedq.make(8)

        @staticmethod
        @modifies_stack([], [queue_type.typ])
        def new(vm):
            vm.push(Queue.make())
            vm.cast(queue_type.typ)

        @staticmethod
        @modifies_stack([queue_type.typ], [value.IntType()])
        def isempty(vm):
            # q -> isempty
            queue_type.get("boundedq")(vm)
            boundedq.isempty(vm)

        @staticmethod
        @modifies_stack([queue_type.typ, typ], [queue_type.typ])
        def put(vm):
            # q item -> updatedq
            queue_type.get("boundedq")(vm)
            vm.dup0()
            boundedq.isfull(vm)
            # bqisfull bq item
            vm.ifelse(lambda vm: [
                # bq item
                vm.dup0(),
                boundedq.struct.get("nextPutIndex")(vm),
                vm.push(2),
                vm.mul(),
                boundedq.new(vm),
                vm.swap1(),
                # oldbq newbq item
                vm.while_loop(lambda vm: [
                    vm.dup0(),
                    boundedq.isempty(vm),
                    vm.push(0),
                    vm.eq()
                    # (oldbq is nonempty) oldbq newbq item
                ], lambda vm: [
                    # oldbq newbq item
                    boundedq.get(vm),
                    # moveitem oldbq newbq item
                    vm.swap1(),
                    vm.swap2(),
                    # newbq moveitem oldbq item
                    boundedq.put(vm),
                    # newbq oldbq item
                    vm.swap1(),
                ]),
                # oldbq newbq item
                vm.pop()
            ])
            # bq item
            boundedq.put(vm)
            queue_type.set_val("boundedq")(vm)

        @staticmethod
        @modifies_stack([queue_type.typ], [typ, queue_type.typ])
        def get(vm):
            # assume queue is non-empty
            # q -> item q
            queue_type.get("boundedq")(vm)
            boundedq.get(vm)
            vm.swap1()
            queue_type.set_val("boundedq")(vm)
            vm.swap1()
    Queue.typ = queue_type.typ
    return Queue

queue = make_queue_type(value.ValueType())
queue_tup = make_queue_type(value.TupleType())
