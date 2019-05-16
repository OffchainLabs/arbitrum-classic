from unittest import TestCase

from arbitrum.std import sized_bigtuple
from arbitrum import VM
from arbitrum import value


class TestSizedBigTuple(TestCase):
    def test_static_get(self):
        vm = VM()
        sized_bigtuple.new(vm)
        for val in range(200):
            vm.push(val)
            vm.push(val + 100)
            vm.swap2()
            sized_bigtuple.set_val(vm)

        tup = vm.stack.items[0]
        for val in range(200):
            self.assertEqual(sized_bigtuple.get_static(tup, val), val + 100)

    def test_static_set(self):
        tup = sized_bigtuple.make()
        for val in range(200):
            tup = sized_bigtuple.set_static(tup, val, val + 100)

        for val in range(200):
            self.assertEqual(sized_bigtuple.get_static(tup, val), val + 100)

    def test_getset(self):
        indexes = [1, 10, 5, 6, 20, 100, 7, 5]
        vm = VM()
        sized_bigtuple.new(vm)
        for val in indexes:
            vm.push(val)
            vm.push(val + 100)
            vm.swap2()
            sized_bigtuple.set_val(vm)

        for val in indexes:
            vm.push(val)
            vm.dup1()
            sized_bigtuple.get(vm)
            self.assertEqual(val + 100, vm.stack.items[0])
            vm.pop()
