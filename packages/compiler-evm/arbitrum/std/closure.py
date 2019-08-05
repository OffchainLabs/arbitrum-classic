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
from .struct import Struct
from .queue import queue


def make_closure(func, capture_count):
    bound_types = func.pops[:capture_count]
    param_types = func.pops[capture_count:]
    struct = Struct(
        "closure[{}_{}_{}]".format(func.__module__, func.__name__, capture_count),
        [
            ("capture", queue.typ),
        ]
    )
    typ = struct.typ

    class Closure:
        @staticmethod
        @modifies_stack(bound_types, [typ], typ.name)
        def new(vm):
            queue.new(vm)
            for _ in range(capture_count):
                queue.put(vm)
            struct.set_val("capture")(vm)

        @staticmethod
        @modifies_stack([typ] + param_types, func.pushes, typ.name)
        def call(vm):
            struct.get("capture")(vm)
            for typ in bound_types[::-1]:
                queue.get(vm)
                vm.cast(typ)
                vm.swap1()
            vm.pop()
            func(vm)
    Closure.typ = typ
    return Closure
