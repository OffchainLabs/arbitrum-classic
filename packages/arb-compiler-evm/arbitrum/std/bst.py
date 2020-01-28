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


# [bst key]
# BUG: Handle return value 0 for value not found
@modifies_stack(2, 1)
def bst_find(vm):
    vm.while_loop(
        lambda vm: [
            vm.dup0(),
            vm.tgetn(0),
            vm.dup2(),
            vm.eq(),
            vm.dup1(),
            vm.tnewn(1),
            vm.eq(),
            vm.bitwise_or(),
            vm.iszero(),
        ],
        lambda vm: [
            vm.dup0(),
            vm.tgetn(0),
            vm.dup2(),
            vm.lt(),
            vm.ifelse(lambda vm: vm.tgetn(2), lambda vm: vm.tgetn(3)),
        ],
    )
    vm.tgetn(1)
    vm.swap1()
    vm.pop()


def make_static_lookup(items, default_val=None):
    if not default_val:
        default_val = value.Tuple([])
    return _make_static_lookup([(x, items[x]) for x in sorted(items)], default_val)


def _make_static_lookup(items, default_val):
    if len(items) >= 3:
        mid = len(items) // 2
        left = items[:mid]
        right = items[mid + 1 :]
        pivot = items[mid][0]

        def impl_n(vm):
            # index
            vm.push(pivot)
            vm.dup1()
            vm.lt()
            # index < pivot, index
            vm.ifelse(
                lambda vm: [
                    # index < pivot
                    _make_static_lookup(left, default_val)(vm)
                ],
                lambda vm: [
                    vm.push(pivot),
                    vm.dup1(),
                    vm.gt(),
                    vm.ifelse(
                        lambda vm: [
                            # index > pivot
                            _make_static_lookup(right, default_val)(vm)
                        ],
                        lambda vm: [
                            # index == pivot
                            vm.pop(),
                            vm.push(items[mid][1]),
                        ],
                    ),
                ],
            )

        return impl_n

    if len(items) == 2:

        def impl_2(vm):
            # index
            vm.dup0()
            vm.push(items[0][0])
            vm.eq()
            vm.ifelse(
                lambda vm: [vm.pop(), vm.push(items[0][1])],
                lambda vm: [
                    vm.push(items[1][0]),
                    vm.eq(),
                    vm.ifelse(
                        lambda vm: [vm.push(items[1][1])],
                        lambda vm: [vm.push(default_val)],
                    ),
                ],
            )

        return impl_2

    if len(items) == 1:

        def impl_1(vm):
            vm.push(items[0][0])
            vm.eq()
            vm.ifelse(
                lambda vm: [vm.push(items[0][1])], lambda vm: [vm.push(default_val)]
            )

        return impl_1

    def impl_0(vm):
        vm.push(default_val)

    return impl_0
