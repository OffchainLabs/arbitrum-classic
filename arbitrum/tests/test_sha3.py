import hashlib
import eth_utils
import random
from unittest import TestCase

from arbitrum.std import sha3, byterange
from arbitrum import VM


class TestSha3(TestCase):
    def test_vec0(self):
        vm = VM()
        sha3.ctx_new(vm)
        sha3.ctx_finish(vm)
        real_hash = int.from_bytes(hashlib.sha3_256().digest(), byteorder="big")
        self.assertEqual(vm.stack[0], real_hash)

    def test_vec1(self):
        vm = VM()
        sha3.ctx_new(vm)
        hasher = hashlib.sha3_256()
        for i in range(200):
            vm.push(0xa3)
            vm.swap1()
            sha3.ctx_pushbyte(vm)
            hasher.update(bytes([0xa3]))
        real_hash = int.from_bytes(hasher.digest(), byteorder="big")
        sha3.ctx_finish(vm)
        self.assertEqual(vm.stack[0], real_hash)

    def test_random_sha3(self):
        data = bytearray(random.getrandbits(8) for _ in range(64))
        vm = VM()
        sha3.ctx_new(vm)
        hasher = hashlib.sha3_256()
        hasher.update(data)
        for v in data:
            vm.push(v)
            vm.swap1()
            sha3.ctx_pushbyte(vm)
        real_hash = int.from_bytes(hasher.digest(), byteorder="big")
        sha3.ctx_finish(vm)
        self.assertEqual(vm.stack[0], real_hash)

    def test_random_keccak256(self):
        data = bytearray(random.getrandbits(8) for _ in range(64))
        vm = VM()
        sha3.ctx_new(vm)
        for v in data:
            vm.push(v)
            vm.swap1()
            sha3.ctx_pushbyte(vm)
        real_hash = int.from_bytes(eth_utils.crypto.keccak(data), byteorder="big")
        sha3.keccak_ctx_finish(vm)
        self.assertEqual(real_hash, vm.stack[0])

    def test_byterange_keccack256(self):
        data = bytearray(random.getrandbits(8) for _ in range(200))
        vm = VM()
        vm.push(len(data))
        vm.push(byterange.frombytes(data))
        sha3.hash_byterange(vm)
        real_hash = int.from_bytes(eth_utils.crypto.keccak(data), byteorder="big")
        self.assertEqual(real_hash, vm.stack[0])

    def test_failing_keccak256(self):
        data = bytearray.fromhex("00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000")
        vm = VM()
        sha3.ctx_new(vm)
        for v in data:
            vm.push(v)
            vm.swap1()
            sha3.ctx_pushbyte(vm)
        real_hash = int.from_bytes(eth_utils.crypto.keccak(data), byteorder="big")
        sha3.keccak_ctx_finish(vm)
        print(real_hash)
        self.assertEqual(real_hash, vm.stack[0])
