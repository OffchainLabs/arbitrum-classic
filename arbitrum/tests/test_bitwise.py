from unittest import TestCase

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
