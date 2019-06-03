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

from unittest import TestCase

from arbitrum.std import sized_byterange
from arbitrum import VM


class TestSizedByteRange(TestCase):
    def test_getset(self):
        indexes = [1, 10, 5, 6, 20, 100, 7, 5]
        vm = VM()
        sized_byterange.new(vm)
        for val in indexes:
            vm.push(val * 58)
            vm.push(val + 100)
            vm.swap2()
            sized_byterange.set_val(vm)

        for val in indexes:
            vm.push(val * 58)
            vm.dup1()
            sized_byterange.get(vm)
            self.assertEqual(val + 100, vm.stack.items[0])
            vm.pop()

    def test_static_get(self):
        vm = VM()
        sized_byterange.new(vm)
        for val in range(200):
            vm.push(val * 58)
            vm.push(val + 100)
            vm.swap2()
            sized_byterange.set_val(vm)

        tup = vm.stack.items[0]
        for val in range(200):
            vm.push(val * 58)
            vm.dup1()
            sized_byterange.get(vm)
            self.assertEqual(sized_byterange.get_static(tup, val * 58), vm.stack.items[0])
            vm.pop()

    def test_static_set(self):
        br = sized_byterange.make()
        for val in range(200):
            br = sized_byterange.set_static(br, val * 58, val + 100)

        for val in range(200):
            self.assertEqual(sized_byterange.get_static(br, val * 58), val + 100)
