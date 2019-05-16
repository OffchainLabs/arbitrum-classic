from ..annotation import modifies_stack


# [bst key]
# BUG: Handle return value 0 for value not found
@modifies_stack(2, 1)
def bst_find(vm):
    vm.while_loop(lambda vm: [
        vm.dup0(),
        vm.tgetn(0),
        vm.dup2(),
        vm.eq(),
        vm.dup1(),
        vm.tnewn(1),
        vm.eq(),
        vm.bitwise_or(),
        vm.iszero()
    ], lambda vm: [
        vm.dup0(),
        vm.tgetn(0),
        vm.dup2(),
        vm.lt(),
        vm.ifelse(
            lambda vm: vm.tgetn(2),
            lambda vm: vm.tgetn(3)
        )
    ])
    vm.tgetn(1)
    vm.swap1()
    vm.pop()
