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