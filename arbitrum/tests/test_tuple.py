from unittest import TestCase

from arbitrum.std import tup
from arbitrum import VM


class TestTuple(TestCase):
    def test_pack(self):
        for i in range(1, 100):
            with self.subTest():
                data = list(range(100))
                vm = VM()
                vm.stack.items = list(data)
                tup.pack(i)(vm)
                self.assertEqual(vm.stack[1:], data[i:])

    def test_pack_unpack(self):
        for i in range(1, 100):
            with self.subTest():
                data = list(range(100))
                vm = VM()
                vm.stack.items = list(data)
                tup.pack(i)(vm)
                tup.unpack(i)(vm)
                self.assertEqual(vm.stack[:], data[:])
