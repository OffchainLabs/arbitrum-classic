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

from arbitrum.std import bigstruct
from arbitrum import VM


class TestBigStruct(TestCase):
    def test_getset(self):
        vm = VM()
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        bs.initialize(vm)
        vm.dup0()
        bs.get('13', vm)
        self.assertEqual(vm.stack[0], 13)
        vm.pop()
        vm.push(666)
        vm.swap1()
        bs.set_val('25', vm)
        vm.pop()

    def test_contains(self):
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        self.assertTrue('50' in bs)
        self.assertFalse('105' in bs)

    def test_get_static(self):
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        for i in range(101):
            self.assertEqual(bs[str(i + 1)], i + 1)

    def test_set_static(self):
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        for i in range(101):
            bs.set_static(str(i + 1), i + 100)
        for i in range(101):
            self.assertEqual(bs[str(i + 1)], i + 100)

    def test_get(self):
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        vm = VM()
        bs.initialize(vm)
        for i in range(101):
            vm.dup0()
            bs.get(str(i + 1), vm)
            self.assertEqual(vm.stack[0], i + 1)
            vm.pop()

    def test_set(self):
        bs = bigstruct.BigStruct([(i+1, str(i+1), i+1) for i in range(101)])
        vm = VM()
        bs.initialize(vm)
        for i in range(101):
            vm.push(i + 100)
            vm.swap1()
            bs.set_val(str(i + 1), vm)

        for i in range(101):
            vm.dup0()
            bs.get(str(i + 1), vm)
            self.assertEqual(vm.stack[0], i + 100)
            vm.pop()
