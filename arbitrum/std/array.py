from .. import value
from ..annotation import modifies_stack

def build_array(arr):
    arr_length = len(arr)
    if len(arr) > 8:
        base_chunk_size = arr_length // 8
        ret = []
        offset = 0
        for chunk in range(8):
            size = base_chunk_size
            if (8 - chunk) <= arr_length % 8:
                size += 1
            ret.append(build_array(arr[offset:offset + size]))
            offset += size
        return value.Tuple(ret)

    if arr_length == 1:
        return arr[0]

    return value.Tuple(arr)

def build_array_type(arr):
    arr_length = len(arr)
    if len(arr) > 8:
        base_chunk_size = arr_length // 8
        ret = []
        offset = 0
        for chunk in range(8):
            size = base_chunk_size
            if (8 - chunk) <= arr_length % 8:
                size += 1
            ret.append(build_array_type(arr[offset:offset + size]))
            offset += size
        return value.TupleType(ret)

    if arr_length == 1:
        return arr[0]

    return value.TupleType(arr)


def array_path(arr_size, index):
    if arr_size > 8:
        base_chunk_size = arr_size // 8
        offset = 0
        for chunk in range(8):
            size = base_chunk_size
            if (8 - chunk) <= arr_size % 8:
                size += 1
            if offset <= index < offset + size:
                return [chunk] + array_path(size, index - offset)
            offset += size
        assert False
        return []
    if arr_size == 1:
        return []

    return [index]


class Array:

    def __init__(self, types):
        if isinstance(types, int):
            types = [value.ValueType()]*types
        self.typ = build_array_type(types)
        self.types = types
        self.length = len(types)

    def make(self):
        return build_array([value.Tuple([])] * self.length)

    def update_type(self, index, typ):
        self.types[index] = typ
        self.typ = build_array_type(self.types)

    @staticmethod
    def from_list(vals):
        return build_array(vals)

    def new(self, vm):
        vm.push(self.make())

    def build(self, vm):
        vm.push(self.make())
        for i in range(len(self.types)):
            self.set_val(i)(vm)

    def get(self, i):
        if self.length == 1:
            @modifies_stack(
                [],
                [],
                f"{self.length}_0"
            )
            def get(vm):
                pass
            return get

        def binder(index):
            @modifies_stack(
                [self.typ],
                [self.types[index]],
                f"{self.length}_{index}"
            )
            def get(vm):
                path = array_path(self.length, index)
                for i in path:
                    vm.tgetn(i)
            return get
        return binder(i)

    def set_val(self, i):
        if self.length == 1:
            @modifies_stack(
                [],
                [],
                f"{self.length}_0"
            )
            def set_val(vm):
                pass
            return set_val

        def binder(index):
            @modifies_stack(
                [self.typ, self.types[index]],
                [self.typ],
                f"{self.length}_{index}"
            )
            def set_val(vm):
                # [array, val]
                path = array_path(self.length, index)
                if len(path) == 1:
                    vm.tsetn(path[0])
                elif len(path) == 2:
                    vm.swap1()
                    # [val, array]
                    vm.dup1()
                    vm.tgetn(path[0])
                    vm.tsetn(path[1])
                    vm.swap1()
                    vm.tsetn(path[0])
                else:
                    vm.swap1()
                    vm.auxpush()
                    # [array]
                    for i in path[:-1]:
                        vm.dup0()
                        vm.tgetn(i)
                    vm.auxpop()
                    for i in path[::-1]:
                        vm.swap1()
                        vm.tsetn(i)
            return set_val
        return binder(i)
