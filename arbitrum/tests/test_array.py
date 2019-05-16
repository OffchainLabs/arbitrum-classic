from unittest import TestCase

from arbitrum.std import Array
from arbitrum import VM


class TestArray(TestCase):
    def test_get(self):
        vm = VM()
        arr = Array(100)
        vm.push(Array.from_list(list(range(100))))
        for i in range(100):
            vm.dup0()
            arr.get(i)(vm)
            self.assertEqual(vm.stack[0], i)
            vm.pop()

    def test_set(self):
        vm = VM()
        arr = Array(100)
        vm.push(Array.from_list(list(range(100))))
        for i in range(100):
            vm.push(i + 100)
            vm.swap1()
            arr.set_val(i)(vm)

        result = Array.from_list([x + 100 for x in range(100)])
        self.assertEqual(vm.stack[0], result)
