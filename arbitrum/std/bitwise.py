from ..annotation import modifies_stack
from .. import value


# [number, bits]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def shift_left(vm):
    vm.swap1()
    vm.push(2)
    vm.exp()
    vm.mul()


# [number, bits]
@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def shift_right(vm):
    vm.swap1()
    vm.push(2)
    vm.exp()
    vm.swap1()
    vm.div()


# [bits]
# 2 ** bits - 1
@modifies_stack([value.IntType()], [value.IntType()])
def n_lowest_mask(vm):
    vm.push(2)
    vm.exp()
    vm.push(1)
    vm.swap1()
    vm.sub()


# [bits]
# 2 ** bits - 1 << (256 - bits)
@modifies_stack([value.IntType()], [value.IntType()])
def n_highest_mask(vm):
    vm.dup0()
    n_lowest_mask(vm)
    vm.swap1()
    vm.push(256)
    vm.sub()
    vm.swap1()
    shift_left(vm)


def n_lowest_mask_static(bits):
    return 2**bits - 1


def n_highest_mask_static(bits):
    return n_lowest_mask_static(bits) << (256 - bits)


@modifies_stack([value.IntType()], [value.IntType()])
def flip_endianness(vm):
    flip_endianness_impl(vm, 32)

def flip_endianness_impl(vm, numBytes):
    if numBytes>1:
        nb2 = numBytes//2
        mod = 1<<(8*nb2)
        vm.push(mod)
        vm.dup1()
        # x mod x
        vm.div()
        # x//mod x
        flip_endianness_impl(vm, nb2)
        # flipped(x//mod) x
        vm.swap1()
        vm.push(mod)
        vm.swap1()
        # x mod flipped(x//mod)
        vm.mod()
        flip_endianness_impl(vm, nb2)
        # flipped(x%mod) flipped(x//mod)
        vm.push(mod)
        vm.mul()
        vm.bitwise_or()
