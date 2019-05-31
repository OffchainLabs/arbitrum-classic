from unittest import TestCase
import random

from arbitrum.std import bitwise
from arbitrum import VM


class TestArray(TestCase):
    def test_flip_endianness(self):
        hexstr = bytearray.fromhex("ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d")
        littleInt = int.from_bytes(hexstr, byteorder="little")
        bigInt = int.from_bytes(hexstr, byteorder="big")

        vm = VM()
        vm.push(littleInt)
        bitwise.flip_endianness(vm)
        self.assertEqual(vm.stack[0], bigInt)

    def test_set_byte(self):
        origstring = bytearray.fromhex("ada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d")
        bigInt = int.from_bytes(origstring, byteorder="big")
        for i in range(32):
            new_val = random.getrandbits(8)
            vm = VM()
            vm.push(new_val)
            vm.push(i)
            vm.push(bigInt)
            bitwise.set_byte(vm)
            finalstring = bytearray(origstring)
            finalstring[i] = new_val
            self.assertEqual(vm.stack[0], int.from_bytes(finalstring, byteorder="big"))
        
