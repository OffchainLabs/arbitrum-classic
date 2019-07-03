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
import random
import eth_utils

from arbitrum.std import byterange, sized_byterange
from arbitrum import VM


class TestByteRange(TestCase):
    def test_getset(self):
        indexes = [1, 10, 5, 6, 20, 100, 7, 5]
        vm = VM()
        byterange.new(vm)
        for val in indexes:
            vm.push(val * 58)
            vm.push(val + 100)
            vm.swap2()
            byterange.set_val(vm)

        for val in indexes:
            vm.push(val * 58)
            vm.dup1()
            byterange.get(vm)
            self.assertEqual(val + 100, vm.stack.items[0])
            vm.pop()

    def test_static_get(self):
        vm = VM()
        byterange.new(vm)
        for val in range(200):
            vm.push(val * 58)
            vm.push(val + 100)
            vm.swap2()
            byterange.set_val(vm)

        tup = vm.stack.items[0]
        for val in range(200):
            vm.push(val * 58)
            vm.dup1()
            byterange.get(vm)
            self.assertEqual(byterange.get_static(tup, val * 58), vm.stack.items[0])
            vm.pop()

    def test_static_set(self):
        br = byterange.make()
        for val in range(200):
            br = byterange.set_static(br, val * 58, val + 100)

        for val in range(200):
            self.assertEqual(byterange.get_static(br, val * 58), val + 100)

    def test_get8(self):
        data = bytearray(random.getrandbits(8) for _ in range(100))
        br = byterange.frombytes(data)
        for i in range(100):
            with self.subTest(index=i):
                vm = VM()
                vm.push(br)
                vm.push(i)
                vm.dup1()
                byterange.get8(vm)
                self.assertEqual(data[i], vm.stack[0])

    def test_set8(self):
        data = bytearray(random.getrandbits(8) for _ in range(100))
        br = byterange.frombytes(data)
        update_bytes = random.getrandbits(8)
        for i in range(100):
            with self.subTest(index=i):
                vm = VM()
                vm.push(update_bytes)
                vm.push(i)
                vm.push(br)
                byterange.set_val8(vm)

                solution = bytearray(data)
                solution[i] = update_bytes
                self.assertEqual(solution.hex(), sized_byterange.tohex([vm.stack[0], 100])[2:])

    def test_frombytes(self):
        data = bytearray(random.getrandbits(8) for _ in range(500))
        data2 = bytearray(data)
        if len(data2) % 32 != 0:
            data2 = data2 + b'\0'*(32 - (len(data2) % 32))
        chunks = [
            eth_utils.big_endian_to_int(data2[i: i + 32])
            for i in range(0, len(data2), 32)
        ]
        vm = VM()
        byterange.new(vm)
        for i, chunk in enumerate(chunks):
            vm.push(i * 32)
            vm.push(chunk)
            vm.swap2()
            byterange.set_val(vm)
        self.assertEqual(byterange.frombytes(data), vm.stack[0])

    def test_subset(self):
        for (start, stop) in [(0, 32), (0, 16), (0, 6), (100, 200), (33, 107)]:
            with self.subTest(start=start, stop=stop):
                data = bytearray(random.getrandbits(8) for _ in range(500))
                vm = VM()
                vm.push(stop)
                vm.push(start)
                vm.push(byterange.frombytes(data))
                byterange.get_subset(vm)
                self.assertEqual(byterange.frombytes(data[start:stop]), vm.stack[0])

    def test_copy(self):
        indexes = [(0, 32, 0), (0, 16, 0), (0, 32, 32), (0, 6, 0), (37, 108, 42)]
        for (source_start, source_end, dest_start) in indexes:
            with self.subTest(source_start=source_start, source_end=source_end, dest_start=dest_start):
                source = bytearray(random.getrandbits(8) for _ in range(500))
                dest = bytearray(random.getrandbits(8) for _ in range(500))
                size = source_end - source_start
                vm = VM()
                vm.push(dest_start)
                vm.push(byterange.frombytes(dest))
                vm.push(source_end)
                vm.push(source_start)
                vm.push(byterange.frombytes(source))
                byterange.copy(vm)
                result = dest[:dest_start] + source[source_start:source_end] + dest[dest_start + size:]
                self.assertEqual(byterange.frombytes(result), vm.stack[0])
