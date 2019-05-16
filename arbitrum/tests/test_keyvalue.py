from unittest import TestCase

from arbitrum.std import keyvalue, keyvalue_int_int
from arbitrum import value
from arbitrum import VM


class TestKeyValue(TestCase):
    def test_getset(self):
        vm = VM()
        keyvalue.new(vm)
        for val in range(200):
            vm.push(val)
            vm.push(val + 1000)
            vm.swap2()
            keyvalue.set_val(vm)

        for val in range(200):
            vm.push(val)
            vm.dup1()
            keyvalue.get(vm)
            self.assertEqual(val + 1000, vm.stack.items[0])
            vm.pop()

    def test_keyvalue_int(self):
        vm = VM()
        keyvalue_int_int.new(vm)
        for val in range(200):
            vm.push(val)
            vm.push(val + 1000)
            vm.swap2()
            keyvalue_int_int.set_val(vm)

        vm.push(10000)
        vm.dup1()
        keyvalue_int_int.get(vm)
        self.assertEqual(0, vm.stack.items[0])
        vm.pop()

        vm.push(100)
        vm.dup1()
        keyvalue_int_int.get(vm)
        self.assertEqual(1100, vm.stack.items[0])

    def test_get_static(self):
        vm = VM()
        keyvalue.new(vm)
        for val in range(200):
            vm.push(val)
            vm.push(val + 1000)
            vm.swap2()
            keyvalue.set_val(vm)

        kvs = vm.stack.items[0]
        for val in range(200):
            vm.push(val)
            vm.dup1()
            keyvalue.get(vm)
            self.assertEqual(keyvalue.get_static(kvs, val), vm.stack.items[0])
            vm.pop()

        self.assertEqual(keyvalue.get_static(kvs, 100000), value.Tuple([]))

    def test_set_static(self):
        kvs = keyvalue.make()
        for val in range(200):
            kvs = keyvalue.set_static(kvs, val, val + 1000)

        vm = VM()
        vm.push(kvs)
        for val in range(200):
            vm.push(val)
            vm.dup1()
            keyvalue.get(vm)
            self.assertEqual(vm.stack[0], val + 1000)
            vm.pop()

        kvs = keyvalue.set_static(kvs, 100, 2100)
        self.assertEqual(keyvalue.get_static(kvs, 100), 2100)
