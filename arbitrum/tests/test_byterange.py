from unittest import TestCase
import random
import eth_utils

from arbitrum.std import byterange
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
        vm = VM()
        vm.push(byterange.frombytes(data))
        for i in range(100):
            vm.push(i)
            vm.dup1()
            byterange.get8(vm)
            self.assertEqual(data[i], vm.stack[0])
            vm.pop()

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
