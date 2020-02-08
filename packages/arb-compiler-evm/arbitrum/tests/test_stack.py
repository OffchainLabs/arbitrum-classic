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

import random
from unittest import TestCase

from arbitrum.std import stack, bytestack_frombytes, bytestack_tohex
from arbitrum import VM


class TestStack(TestCase):
    def test_simple(self):
        vm = VM()
        stack.new(vm)
        for val in range(200):
            vm.push(val)
            vm.swap1()
            stack.push(vm)

        for val in range(199, -1, -1):
            stack.pop(vm)
            self.assertEqual(vm.stack[0], val)
            vm.pop()

    def test_empty(self):
        vm = VM()
        vm.push(10)
        stack.new(vm)
        vm.dup0()
        stack.isempty(vm)
        self.assertTrue(vm.stack[0])
        vm.pop()
        stack.push(vm)
        stack.isempty(vm)
        self.assertFalse(vm.stack[0])

    def test_bytestack(self):
        data = bytearray(random.getrandbits(8) for _ in range(60))
        bs = bytestack_frombytes(data)
        print(bs)

        print(data.hex())
        print(bytestack_tohex(bs))
