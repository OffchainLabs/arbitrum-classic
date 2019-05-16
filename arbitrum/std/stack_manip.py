from .stack import stack
from ..annotation import modifies_stack, uncountable_stack, noreturn


@uncountable_stack
def kill(vm):
    vm.while_loop(
        lambda vm: [vm.stackempty(), vm.iszero()],
        lambda vm: vm.pop()
    )

@noreturn
def kill_aux(vm):
    vm.while_loop(
        lambda vm: [
            vm.auxpop(),
            vm.auxstackempty(),
            vm.iszero(),
            vm.swap1(),
            vm.auxpush()
        ],
        lambda vm: [
            vm.auxpop(),
            vm.auxpop(),
            vm.pop(),
            vm.auxpush(),
        ]
    )


@uncountable_stack
def compress(vm):
    stack.new(vm)
    vm.while_loop(
        lambda vm: [
            vm.auxpush(),
            vm.stackempty(),
            vm.iszero(),
            vm.auxpop(),
            vm.swap1()
        ], lambda vm: [
            # stack item
            stack.push(vm)
        ]
    )


@uncountable_stack
def uncompress(vm):
    # compressed_stack
    vm.while_loop(
        lambda vm: [
            vm.dup0(),
            stack.isempty(vm),
            vm.iszero()
        ], lambda vm: [
            stack.pop(vm),
            vm.swap1()
        ]
    )
    vm.pop()


@noreturn
def compress_aux(vm):
    stack.new(vm)
    # store
    vm.while_loop(
        lambda vm: [
            vm.auxpop(),
            vm.auxstackempty(),
            vm.iszero(),
            vm.swap1(),
            vm.auxpush(),
        ], lambda vm: [
            vm.auxpop(),
            vm.swap1(),
            vm.auxpop(),
            # stack item
            stack.push(vm),
            vm.swap1(),
            vm.auxpush(),
        ]
    )

@noreturn
def uncompress_aux(vm):
    # compressed_stack
    vm.while_loop(
        lambda vm: [
            vm.dup0(),
            stack.isempty(vm),
            vm.iszero()
        ], lambda vm: [
            vm.auxpop(),
            vm.swap1(),
            # auxc ret
            stack.pop(vm),
            vm.swap1(),
            # val auxc ret
            vm.auxpush(),
            vm.swap1(),
            # ret auxc
            vm.auxpush()
        ]
    )
    vm.pop()


def dup_n(index):
    # @modifies_stack(index + 1, index + 2, index)
    def dup(vm):
        if index == 0:
            vm.dup0()
        elif index == 1:
            vm.dup1()
        else:
            for _ in range(index - 2):
                vm.auxpush()
            vm.dup2()
            for _ in range(index - 2):
                vm.auxpop()
                vm.swap1()
    return dup


def swap_n(index):
    # @modifies_stack(index + 1, index + 1, index)
    def swap(vm):
        if index == 1:
            vm.swap1()
        else:
            for _ in range(index - 2):
                vm.swap1()
                vm.auxpush()
            vm.swap2()
            for _ in range(index - 2):
                vm.auxpop()
                vm.swap1()
    return swap


def take_n(index):
    @modifies_stack(index, index, index)
    def take(vm):
        if index == 1:
            vm.swap1()
        else:
            for _ in range(index - 2):
                vm.auxpush()
            vm.swap1()
            vm.swap2()
            for _ in range(index - 2):
                vm.auxpop()
                vm.swap1()

    return take


def push_to_n(index):
    @modifies_stack(index, index, index)
    def push_to(vm):
        if index == 1:
            vm.swap1()
        else:
            for _ in range(index - 2):
                vm.swap1()
                vm.auxpush()
            vm.swap2()
            vm.swap1()
            for _ in range(index - 2):
                vm.auxpop()
    return push_to
