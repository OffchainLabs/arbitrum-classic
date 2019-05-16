from .. import value


def make(make_func):
    return value.Tuple([make_func(), 0])


def new(vm, new_func):
    new_func(vm)
    vm.push(value.Tuple([0, 0]))
    vm.tsetn(0)

def get(vm, struct, get_func):
    struct.get("data")(vm)
    get_func(vm)


# [sized_bigtuple, index, value]
def set_val(vm, struct, set_func, unit_size):
    vm.dup0()
    struct.get("size")(vm)
#   [old_size, sized_bigtuple, index, value]
    vm.dup2()
#   [index, old_size, sized_bigtuple, index, value]
    vm.push(unit_size)
    vm.add()
    vm.gt()
    vm.ifelse(
        lambda vm: [
            # [sized_bigtuple, index, value]
            vm.dup1(),
            vm.push(unit_size),
            vm.add(),
            vm.swap1(),
            struct.set_val("size")(vm)
        ]
    )
    vm.swap2()
    vm.swap1()
    vm.dup2()
    struct.get("data")(vm)
    # [bigtuple, index, value, sized_bigtuple]
    set_func(vm)
    # [bigtuple, sized_bigtuple]
    vm.swap1()
    struct.set_val("data")(vm)


def get_static(sized, index, get_func):
    return get_func(sized[0], index)


def set_static(sized, index, val, set_func, unit_size):
    if index + unit_size > sized[1]:
        sized = sized.set_tup_val(1, index + unit_size)

    return sized.set_tup_val(0, set_func(sized[0], index, val))
