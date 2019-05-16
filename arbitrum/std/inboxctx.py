from ..annotation import modifies_stack
from . import tup
from . import queue_tup
from .stack_manip import swap_n
from .. import value
from .types import inbox_ctx

# inboxctx = [queue nodeAlreadySeen]

typ = inbox_ctx.typ

@modifies_stack(0, [typ])
def new(vm):
    # -> ctx
    vm.tnewn(0)
    queue_tup.new(vm)
    inbox_ctx.new(vm)
    inbox_ctx.set_val("queue")(vm)
    inbox_ctx.set_val("nodeAlreadySeen")(vm)


@modifies_stack([typ], [value.IntType()])
def isempty(vm):
    # ctx -> isempty
    vm.dup0()
    inbox_ctx.get("queue")(vm)
    queue_tup.isempty(vm)
    vm.ifelse(lambda vm: [
        # ctx
        inbox_ctx.get("nodeAlreadySeen")(vm),
        vm.tnewn(8),
        vm.inbox(),
        vm.eq(),
    ], lambda vm: [
        # ctx
        vm.pop(),
        vm.push(0),
    ])

@modifies_stack([typ], [value.ValueType(), typ])
def getmsg(vm):
    # ctx -> msg updatedctx
    vm.dup0()
    inbox_ctx.get("queue")(vm)
    queue_tup.isempty(vm)
    vm.ifelse(lambda vm: [
        vm.dup0(),
        inbox_ctx.get("queue")(vm),
        vm.swap1(),
        inbox_ctx.get("nodeAlreadySeen")(vm),
        vm.dup0(),
        vm.inbox(),
        vm.swap2(),
        vm.swap1(),
        vm.dup2(),
        # inbox nodeAlreadySeen queue inbox
        _inhale2(vm),
        # nodeAlreadySeen updatedq inbox
        vm.pop(),
        # updatedq inbox
        queue_tup.get(vm),
        # msg updatedq inbox
        vm.swap2(),
        inbox_ctx.new(vm),
        inbox_ctx.set_val("nodeAlreadySeen")(vm),
        inbox_ctx.set_val("queue")(vm)
        # ctx msg
    ], lambda vm: [
        # ctx
        vm.dup0(),
        inbox_ctx.get("queue")(vm),
        # queue ctx
        queue_tup.get(vm),
        # msg updatedq ctx
        vm.swap2(),
        # ctx updatedq msg
        inbox_ctx.set_val("queue")(vm)
    ])
    # updatedctx msg
    vm.swap1()

# @modifies_stack([typ], [value.ValueType(), typ])
# def getmsg(vm):
#     # assume inboxctx is not empty
#     # ctx -> msg updatedctx
#     inhale(vm)
#     # ctx
#     vm.dup0()
#     inbox_ctx.get("queue")(vm)
#     # queue ctx
#     queue_tup.get(vm)
#     # msg updatedq ctx
#     vm.swap2()
#     # ctx updatedq msg
#     inbox_ctx.set_val("queue")(vm)
#     # updatedctx msg
#     vm.swap1()


@modifies_stack([typ], [typ])
def inhale(vm):
    # ctx -> updatedctx
    vm.dup0()
    inbox_ctx.get("queue")(vm)
    vm.swap1()
    inbox_ctx.get("nodeAlreadySeen")(vm)
    vm.tnewn(8)
    vm.inbox()
    vm.swap2()
    vm.swap1()
    vm.dup2()
    # inbox nodeAlreadySeen queue inbox
    _inhale2(vm)
    # nodeAlreadySeen updatedq inbox
    vm.pop()
    inbox_ctx.new(vm)
    inbox_ctx.set_val("queue")(vm)
    inbox_ctx.set_val("nodeAlreadySeen")(vm)


@modifies_stack([
    value.TupleType(),
    value.TupleType(),
    queue_tup.typ
], [
    value.TupleType(),
    queue_tup.typ
])
def _inhale2(vm):
    # inbox nodeAlreadySeen queue -> nodeAlreadySeen updatedq
    vm.dup1()
    vm.dup1()
    vm.eq()
    vm.dup1()
    vm.tnewn(0)
    vm.eq()
    vm.bitwise_or()
    # inbox==nodeAlreadySeen inbox nodeAlreadySeen queue
    vm.ifelse(lambda vm: [
        vm.pop()
    ], lambda vm: [
        vm.cast(value.TupleType([value.IntType(), value.TupleType(), value.TupleType()])),
        # inbox nodeAlreadySeen queue
        tup.tbreak(3)(vm),
        vm.push(0),
        vm.eq(),
        vm.ifelse(_inhale_message, _inhale_messages),
    ])

@modifies_stack([
    value.TupleType(),
    value.TupleType(),
    value.TupleType(),
    queue_tup.typ
], [
    value.TupleType(),
    queue_tup.typ
])
def _inhale_message(vm):
    # inbox msg nodeAlreadySeen queue
    vm.swap1()
    vm.auxpush()
    # inbox nodeAlreadySeen queue
    _inhale2(vm)
    # nodeAlreadySeen queue
    vm.swap1()
    vm.auxpop()
    vm.swap1()
    # queue msg nodeAlreadySeen
    queue_tup.put(vm)
    vm.swap1()


@modifies_stack([
    value.TupleType(),
    value.TupleType(),
    value.TupleType(),
    queue_tup.typ
], [
    value.TupleType(),
    queue_tup.typ
])
def _inhale_messages(vm):
    # inboxA inboxB nodeAlreadySeen queue
    vm.swap1()
    vm.auxpush()
    # inboxA nodeAlreadySeen queue
    _inhale2(vm)
    # nodeAlreadySeen queue
    vm.auxpop()
    _inhale2(vm)
