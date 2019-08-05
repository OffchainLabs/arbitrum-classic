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

from arbitrum.std import bigtuple
from arbitrum import VM
from arbitrum import value


class TestBigTuple(TestCase):
    def test_static_get(self):
        vm = VM()
        bigtuple.new(vm)
        for val in range(200):
            vm.push(val)
            vm.push(val + 100)
            vm.swap2()
            bigtuple.set_val(vm)

        tup = vm.stack.items[0]
        for val in range(200):
            self.assertEqual(bigtuple.get_static(tup, val), val + 100)

    def test_static_set(self):
        tup = value.Tuple([])
        for val in range(200):
            tup = bigtuple.set_static(tup, val, val + 100)

        for val in range(200):
            self.assertEqual(bigtuple.get_static(tup, val), val + 100)

    def test_getset(self):
        indexes = [1, 10, 5, 6, 20, 100, 7, 5]
        vm = VM()
        bigtuple.new(vm)
        for val in indexes:
            vm.push(val)
            vm.push(val + 100)
            vm.swap2()
            bigtuple.set_val(vm)

        for val in indexes:
            vm.push(val)
            vm.dup1()
            bigtuple.get(vm)
            self.assertEqual(val + 100, vm.stack.items[0])
            vm.pop()

    def test_subset(self):
        for (start, stop) in [(0, 32), (0, 16), (0, 6), (100, 200), (33, 107)]:
            with self.subTest(start=start, stop=stop):
                data = list(range(200))
                vm = VM()
                vm.push(stop)
                vm.push(start)
                vm.push(bigtuple.fromints(data))
                bigtuple.get_subset(vm)
                self.assertEqual(bigtuple.fromints(data[start:stop]), vm.stack[0])

    def test_copy(self):
        indexes = [(0, 32, 0), (0, 16, 0), (0, 32, 32), (0, 6, 0), (37, 108, 42)]
        for (source_start, source_end, dest_start) in indexes:
            with self.subTest(source_start=source_start, source_end=source_end, dest_start=dest_start):
                source = list(range(200))
                dest = [x + 1000 for x in range(200)]
                size = source_end - source_start
                vm = VM()
                vm.push(dest_start)
                vm.push(bigtuple.fromints(dest))
                vm.push(source_end)
                vm.push(source_start)
                vm.push(bigtuple.fromints(source))
                bigtuple.copy(vm)

                result = dest[:dest_start] + source[source_start:source_end] + dest[dest_start + size:]
                self.assertEqual(bigtuple.fromints(result), vm.stack[0])
