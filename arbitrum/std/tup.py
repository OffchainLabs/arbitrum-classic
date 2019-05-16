from ..annotation import modifies_stack


def make(size):
    def make(vm):
        vm.tnewn(size)
        for i in range(size):
            vm.tsetn(i)

    return make


def tbreak(size):
    def tbreak(vm):
        for i in range(size - 1, 0, -1):
            vm.dup0()
            vm.tgetn(i)
            vm.swap1()
        if size > 0:
            vm.tgetn(0)
    return tbreak


def _get_sizes(size, vals=None):
    if not vals:
        vals = []

    if size == 0:
        return vals

    if size > 8:
        size += 1
        return _get_sizes(size - 8, vals + [8])

    return vals + [size]


def pack(size):
    @modifies_stack(size, 1, size)
    def pack(vm):
        for i in _get_sizes(size):
            make(i)(vm)

    return pack


def unpack(size):
    @modifies_stack(1, size, size)
    def unpack(vm):
        for i in _get_sizes(size)[::-1]:
            tbreak(i)(vm)

    return unpack
