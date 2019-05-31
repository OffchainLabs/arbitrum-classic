from ..annotation import modifies_stack
from .. import value


@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def max(vm):
    vm.dup1()
    vm.dup1()
    vm.lt()
    vm.ifelse(
        lambda vm: [vm.pop()],
        lambda vm: [vm.swap1(), vm.pop()]
    )


@modifies_stack([value.IntType(), value.IntType()], [value.IntType()])
def min(vm):
    vm.dup1()
    vm.dup1()
    vm.lt()
    vm.ifelse(
        lambda vm: [vm.swap1(), vm.pop()],
        lambda vm: [vm.pop()]
    )
