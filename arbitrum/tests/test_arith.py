from unittest import TestCase

from arbitrum.std import arith
from arbitrum import VM


class TestArish(TestCase):
    def test_max(self):
        vm = VM()
        vm.push(10)
        vm.push(15)
        arith.max(vm)
        self.assertEqual(vm.stack[0], 15)

    def test_min(self):
        vm = VM()
        vm.push(10)
        vm.push(15)
        arith.min(vm)
        self.assertEqual(vm.stack[0], 10)