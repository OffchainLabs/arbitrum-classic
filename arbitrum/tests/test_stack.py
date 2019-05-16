from unittest import TestCase

from arbitrum.std import stack
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
