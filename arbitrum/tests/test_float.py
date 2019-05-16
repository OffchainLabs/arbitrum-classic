from unittest import TestCase
import struct

from arbitrum.std import floatlib
from arbitrum import VM


def python_add(i, j):
    if i == _nan or j == _nan:
        return _nan
    if (i == _inf and j == _neg) or (i == _neg and j == _inf):
        return _nan
    if (i == _inf or i == _neg):
        return i
    if (j == _inf or j == _neg):
        return j
    fi = reformat_int_to_float(i)
    fj = reformat_int_to_float(j)
    return reformat_float_to_int(fi+fj)


def python_multiply(i, j):
    if i == _nan or j == _nan:
        return _nan
    if (i == _inf and j == _inf) or (i == _neg and j == _neg):
        return _inf
    if (i == _inf and j == _neg) or (j == _neg and i == _inf):
        return _neg
    if (i in [_inf, _neg] and j & ((1 << 53)-1) == 0):
        return _nan
    if (j in [_inf, _neg] and i & ((1 << 53)-1) == 0):
        return _nan
    if i == _inf:
        if reformat_int_to_float(j) > 0.0:
            return _inf
        else:
            return _neg
    if i == _neg:
        if reformat_int_to_float(j) > 0.0:
            return _neg
        else:
            return _inf
    if j == _inf:
        if reformat_int_to_float(i) > 0.0:
            return _inf
        else:
            return _neg
    if j == _neg:
        if reformat_int_to_float(i) > 0.0:
            return _neg
        else:
            return _inf 
    fi = reformat_int_to_float(i)
    fj = reformat_int_to_float(j)
    return reformat_float_to_int(fi*fj)


def python_divide(i, j):
    if i == _nan or j == _nan:
        return _nan
    if i in [_inf, _neg] and j in [_inf, _neg]:
        return _nan
    if (i & ((1 << 53)-1) == 0 and j & ((1 << 53)-1) == 0):
        return _nan
    if i == _inf:
        fj = reformat_int_to_float(j)
        if fj > 0:
            return _inf
        else:
            return _neg
    if i == _neg:
        fj = reformat_int_to_float(j)
        if fj > 0:
            return _neg
        else:
            return _inf
    fi = reformat_int_to_float(i)
    fj = reformat_int_to_float(j)
    if fj == 0.0:
        if fi > 0.0:
            return _inf
        else:
            return _neg
    else:
        return reformat_float_to_int(fi / fj)


def reformat_int_to_float(i):
    packed = struct.pack('Q', i)
    return struct.unpack('d', packed)[0]


def reformat_float_to_int(f):
    packed = struct.pack('d', f)
    return struct.unpack('Q', packed)[0]


_nan = floatlib.getnan()
_inf = floatlib.getinfinity()
_neg = floatlib.getneginfinity()


class TestBignum(TestCase):
    def setUp(self):
        self.vals = [0, _nan, _inf, _neg]

    def test_addsubtract(self):
        for i in self.vals:
            for j in self.vals:
                vm = VM()
                pyth_res = python_add(i, j)
                vm.push(i)
                vm.push(j)
                floatlib.add(vm)
                self.assertEqual(vm.stack[0], pyth_res)

    def test_multiply(self):
        for i in self.vals:
            for j in self.vals:
                vm = VM()
                pyth_res = python_multiply(i, j)
                vm.push(i)
                vm.push(j)
                floatlib.mul(vm)
                self.assertEqual(vm.stack[0], pyth_res)

    def test_divide(self):
        for i in self.vals:
            for j in self.vals:
                vm = VM()
                pyth_res = python_divide(i, j)
                vm.push(j)
                vm.push(i)
                floatlib.div(vm)
                self.assertEqual(vm.stack[0], pyth_res)
