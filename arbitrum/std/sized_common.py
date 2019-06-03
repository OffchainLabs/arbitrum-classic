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

from .. import value


def make(make_func):
    return value.Tuple([make_func(), 0])


def new(vm, new_func):
    new_func(vm)
    vm.push(value.Tuple([0, 0]))
    vm.tsetn(0)

def get(vm, struct, get_func):
    struct.get("data")(vm)
    get_func(vm)


# [sized_bigtuple, index, value]
def set_val(vm, struct, set_func, unit_size):
    vm.dup0()
    struct.get("size")(vm)
#   [old_size, sized_bigtuple, index, value]
    vm.dup2()
#   [index, old_size, sized_bigtuple, index, value]
    vm.push(unit_size)
    vm.add()
    vm.gt()
    vm.ifelse(
        lambda vm: [
            # [sized_bigtuple, index, value]
            vm.dup1(),
            vm.push(unit_size),
            vm.add(),
            vm.swap1(),
            struct.set_val("size")(vm)
        ]
    )
    vm.swap2()
    vm.swap1()
    vm.dup2()
    struct.get("data")(vm)
    # [bigtuple, index, value, sized_bigtuple]
    set_func(vm)
    # [bigtuple, sized_bigtuple]
    vm.swap1()
    struct.set_val("data")(vm)


def get_static(sized, index, get_func):
    return get_func(sized[0], index)


def set_static(sized, index, val, set_func, unit_size):
    if index + unit_size > sized[1]:
        sized = sized.set_tup_val(1, index + unit_size)

    return sized.set_tup_val(0, set_func(sized[0], index, val))
