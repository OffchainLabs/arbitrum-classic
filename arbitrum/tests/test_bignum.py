from unittest import TestCase
import math
from itertools import product

from arbitrum.std import bignum
from arbitrum.std import random
from arbitrum import VM


class TestBignum(TestCase):
    def setUp(self):
        self.vals = [
            0,
            35,
            73,
            -73,
            144,
            389324890895428428935890459813458123059823590301713946278546193487563194781639487123649281736918237462398756387463985736597823948736298736578324632897461289374619,
            8795078423587904278942359784523870324870245428792457842879418907425784278934108972458904289741389734187934287923089723489024894198074789056781234987412784,
            -8795078423587904278942359784523870324870245428792457842879418907425784278934108972458904289741389734187934287923089723489024894198074789056781234987412784
        ]
        self.bignumVals = [bignum.make_from_int(x) for x in self.vals]

    def iter_vals(self):
        for ((aBig, bBig), (a, b)) in zip(
                product(self.bignumVals, self.bignumVals),
                product(self.vals, self.vals)
        ):
            yield (aBig, bBig, a, b)

    def test_conversion(self):
        for val in self.vals:
            with self.subTest():
                vm = VM()
                vm.push(bignum.make_from_int(val))
                x = bignum.to_python_int(vm.stack.items[0])
                vm.pop()
                self.assertEqual(x, val)

    def test_negation(self):
        for val in self.vals:
            with self.subTest():
                vm = VM()
                vm.push(bignum.make_from_int(val))
                bignum.negate(vm)
                x = bignum.to_python_int(vm.stack.items[0])
                vm.pop()
                self.assertEqual(x, -val)

                vm.push(bignum.make_from_int(-val))
                bignum.negate(vm)
                x = bignum.to_python_int(vm.stack.items[0])
                vm.pop()
                self.assertEqual(x, val)

    def test_addition(self):
        for (aBig, bBig, a, b) in self.iter_vals():
            with self.subTest():
                vm = VM()
                vm.push(aBig)
                vm.push(bBig)
                bignum.add(vm)
                x = bignum.to_python_int(vm.stack.items[0])
                self.assertEqual(x, b + a)

    def test_subtraction(self):
        for (aBig, bBig, a, b) in self.iter_vals():
            with self.subTest():
                vm = VM()
                vm.push(aBig)
                vm.push(bBig)
                bignum.subtract(vm)
                x = bignum.to_python_int(vm.stack.items[0])
                self.assertEqual(x, b - a)

    def test_multiplication(self):
        for (aBig, bBig, a, b) in self.iter_vals():
            with self.subTest():
                vm = VM()
                vm.push(aBig)
                vm.push(bBig)
                bignum.multiply(vm)
                x = bignum.to_python_int(vm.stack.items[0])
                self.assertEqual(x, b * a)

    def test_integer_multiplication(self):
        for (aBig, _, a, b) in self.iter_vals():
            if ((b > -(2**126)) and (b < (2**126))):
                with self.subTest():
                    vm = VM()
                    vm.push(aBig)
                    vm.push(b)
                    bignum.intmultiply(vm)
                    x = bignum.to_python_int(vm.stack.items[0])
                    self.assertEqual(x, b * a)

    def test_divmod(self):
        for (aBig, bBig, a, b) in self.iter_vals():
            if a > 0 and b > 0:
                with self.subTest():
                    vm = VM()
                    vm.push(aBig)
                    vm.push(bBig)
                    bignum.divmodallpositive(vm)
                    q = bignum.to_python_int(vm.stack.items[0])
                    r = bignum.to_python_int(vm.stack.items[1])
                    self.assertEqual(q, b // a)
                    self.assertEqual(r, b % a)

    def test_modpow(self):
        for ((aBig, bBig, cBig), (a, b, c)) in zip(
                product(self.bignumVals, self.bignumVals, self.bignumVals),
                product(self.vals, self.vals, self.vals)
        ):
            if a > 0 and b > 0 and b < 5000 and c > 0 and c < 100000:
                with self.subTest():
                    vm = VM()
                    vm.push(cBig)
                    vm.push(bBig)
                    vm.push(aBig)
                    bignum.modpow(vm)
                    x = bignum.to_python_int(vm.stack.items[0])
                    self.assertEqual(x, pow(a, b, c))

    def test_modinv(self):
        for (aBig, mBig, a, m) in self.iter_vals():
            if a > 0 and a < m and m < 2**500 and math.gcd(a, m) == 1:
                with self.subTest():
                    vm = VM()
                    vm.push(mBig)
                    vm.push(aBig)
                    bignum.modinv(vm)
                    vm.push(mBig)
                    vm.swap1()
                    vm.push(aBig)
                    bignum.modmul(vm)
                    x = bignum.to_python_int(vm.stack.items[0])
                    self.assertEqual(x, 1)

    def test_lessthan(self):
        for (aBig, bBig, a, b) in self.iter_vals():
            with self.subTest():
                vm = VM()
                vm.push(aBig)
                vm.push(bBig)
                bignum.lt(vm)
                if b < a:
                    self.assertEqual(vm.stack[0], 1)
                else:
                    self.assertEqual(vm.stack[0], 0)

    def test_prime(self):
        # # test primality testing
        for val in [5147, 10709, 40423, 84499, 104107]:
            print('testing prime', val)
            vm = VM()
            vm.push(64)
            vm.push(0)
            random.new(vm)
            vm.push(bignum.make_from_int(val))
            bignum.isprime(vm)
            self.assertEqual(vm.stack[0], 1)
        for val in [5148, 10710, 40425, 84500, 104115, 84499*104107]:
            print('testing notprime', val)
            vm = VM()
            vm.push(64)
            vm.push(0)
            random.new(vm)
            vm.push(bignum.make_from_int(val))
            bignum.isprime(vm)
            self.assertEqual(vm.stack[0], 0)
