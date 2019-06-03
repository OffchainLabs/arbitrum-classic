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

from arbitrum.std import stack_manip, stack
from arbitrum import VM


class TestStackManip(TestCase):

    def test_compress(self):
        vm = VM()
        vm.stack.items = list(range(93))
        stack_manip.compress(vm)

        vm2 = VM()
        stack.new(vm2)
        for i in range(93):
            vm2.push(i)
            vm2.swap1()
            stack.push(vm2)

        self.assertEqual(vm.stack[:], vm2.stack[:])

    def test_uncompress(self):
        for i in [0, 97]:
            with self.subTest():
                vm = VM()
                vm.stack.items = list(range(i))
                stack_manip.compress(vm)
                stack_manip.uncompress(vm)
                self.assertEqual(vm.aux_stack[:], [])
                self.assertEqual(vm.stack[:], list(range(i)))

    def test_compress_aux(self):
        vm = VM()
        vm.aux_stack.items = list(range(93))
        stack_manip.compress_aux(vm)

        vm2 = VM()
        stack.new(vm2)
        for i in range(93):
            vm2.push(i)
            stack.push(vm2)

        self.assertEqual(vm.stack[:], vm2.stack[:])

    def test_uncompress_aux(self):
        for i in [0, 97]:
            with self.subTest():
                vm = VM()
                vm.aux_stack.items = list(range(i))
                stack_manip.compress_aux(vm)
                stack_manip.uncompress_aux(vm)
                self.assertEqual(vm.stack[:], [])
                self.assertEqual(vm.aux_stack[:], list(range(i)))

    def test_dup_n(self):
        for i in range(100):
            with self.subTest():
                vm = VM()
                vm.stack.items = list(range(100))
                stack_manip.dup_n(i)(vm)
                self.assertEqual(vm.stack[0], i)
                self.assertEqual(vm.stack[1:], list(range(100)))


    def test_swap_n(self):
        for i in range(1, 100):
            with self.subTest():
                vm = VM()
                orig_list = list(range(100))
                vm.stack.items = list(orig_list)
                stack_manip.swap_n(i)(vm)
                orig_list[0], orig_list[i] = orig_list[i], orig_list[0]
                self.assertEqual(vm.stack[:], orig_list)

    def test_take_n(self):
        for i in range(1, 100):
            with self.subTest():
                vm = VM()
                orig_list = list(range(100))
                vm.stack.items = list(orig_list)
                stack_manip.take_n(i)(vm)
                new_list = [orig_list[i]] + orig_list[:i] + orig_list[i + 1:]
                self.assertEqual(vm.stack[:], new_list)

    def test_push_to_n(self):
        for i in range(1, 100):
            with self.subTest():
                vm = VM()
                orig_list = list(range(100))
                vm.stack.items = list(orig_list)
                stack_manip.push_to_n(i)(vm)
                new_list = orig_list[1:i + 1] + [orig_list[0]] + orig_list[i + 1:]
                self.assertEqual(vm.stack[:], new_list)
