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

from arbitrum.std import queue
from arbitrum import VM


class TestQueue(TestCase):
    def test_simple(self):
        vm = VM()
        queue.new(vm)
        for val in range(200):
            vm.push(val)
            vm.swap1()
            queue.put(vm)

        for val in range(200):
            queue.get(vm)
            self.assertEqual(vm.stack[0], val)
            vm.pop()

    def test_empty(self):
        vm = VM()
        vm.push(10)
        queue.new(vm)
        vm.dup0()
        queue.isempty(vm)
        self.assertTrue(vm.stack[0])
        vm.pop()
        queue.put(vm)
        queue.isempty(vm)
        self.assertFalse(vm.stack[0])
