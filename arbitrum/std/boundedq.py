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
from .struct import Struct

from . import make_bigtuple_type

# bounded queue: [bigtuple nextGetIndex nextPutIndex capacity]
# if nextGetIndex == nextPutIndex, queue is empty
# Note that the max number of items in queue is one less than "capacity"
#             (could probably fix this, but it's not worth the
#              trouble at this point)

def make_boundedq_type(typ):
    bigtuple = make_bigtuple_type(typ, typ.empty_val())
    boundedq_type = Struct(f"boundedq[{typ}]", [
        ("bigtuple", bigtuple.typ),
        ("nextGetIndex", value.IntType()),
        ("nextPutIndex", value.IntType()),
        ("capacity", value.IntType())
    ])

    class BoundedQueue:
        @staticmethod
        def make(capacity):
            return value.Tuple([value.Tuple([]), 0, 0, capacity])


        @staticmethod
        @modifies_stack([value.IntType()], [boundedq_type.typ])
        def new(vm):
            # capacity -> bq
            vm.push(BoundedQueue.make(0))
            vm.tsetn(3)
            vm.cast(boundedq_type.typ)


        @staticmethod
        @modifies_stack([boundedq_type.typ], [value.IntType()])
        def isempty(vm):
            # bq -> isempty
            vm.dup0()
            boundedq_type.get("nextGetIndex")(vm)
            vm.swap1()
            boundedq_type.get("nextPutIndex")(vm)
            # nextPutIndex nextGetIndex
            vm.eq()

        @staticmethod
        def _incrmod_field(vm, field_name):
            # bq -> updatedbq
            vm.dup0()
            boundedq_type.get("capacity")(vm)
            # capacity bq
            vm.dup1()
            boundedq_type.get(field_name)(vm)
            # field capacity bq
            vm.push(1)
            vm.add()
            vm.mod()
            # (field+1)%capacity bq
            vm.swap1()
            boundedq_type.set_val(field_name)(vm)

        @staticmethod
        @modifies_stack([boundedq_type.typ], [value.IntType()])
        def isfull(vm):
            # bq -> isfull
            vm.dup0()
            boundedq_type.get("capacity")(vm)
            # capacity bq

            vm.dup1()
            boundedq_type.get("nextPutIndex")(vm)
            vm.push(1)
            vm.add()

            # nextPutIndex+1 capacity bq
            vm.mod()

            # (npi+1)%cap bq
            vm.swap1()
            boundedq_type.get("nextGetIndex")(vm)
            vm.eq()

        @staticmethod
        @modifies_stack([boundedq_type.typ], [typ, boundedq_type.typ])
        def get(vm):
            # assume queue is non-empty
            # bq -> item updatedbq
            vm.dup0()
            boundedq_type.get("nextGetIndex")(vm)
            # nextGetIndex bq
            vm.dup1()
            boundedq_type.get("bigtuple")(vm)
            # bigtuple nextGetIndex bq
            bigtuple.get(vm)
            # result bq
            vm.swap1()
            BoundedQueue._incrmod_field(vm, "nextGetIndex")
            vm.swap1()

        @staticmethod
        @modifies_stack([boundedq_type.typ, typ], [boundedq_type.typ])
        def put(vm):
            # assume queue is non-full
            # bq item -> updatedbq
            vm.swap1()
            vm.dup1()
            boundedq_type.get("nextPutIndex")(vm)
            # nextPutIndex item bq
            vm.dup2()
            boundedq_type.get("bigtuple")(vm)
            # bigtuple nextPutIndex item bq
            bigtuple.set_val(vm)
            # updatedbigtuple bq
            vm.swap1()
            boundedq_type.set_val("bigtuple")(vm)
            # bq
            BoundedQueue._incrmod_field(vm, "nextPutIndex")
            # updatedbq
    BoundedQueue.typ = boundedq_type.typ
    BoundedQueue.struct = boundedq_type
    return BoundedQueue


boundedq = make_boundedq_type(value.ValueType())
