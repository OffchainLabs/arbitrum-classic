from ..annotation import modifies_stack
from .. import value
from .struct import Struct

def make_stack_type(typ):
    stack_type = Struct(f"stack[{typ}]", [
        "prev",
        ("val", typ)
    ])

    class Stack:
        # [] -> [stack]
        @staticmethod
        @modifies_stack([], [stack_type.typ])
        def new(vm):
            vm.tnewn(0)
            vm.cast(stack_type.typ)

        @staticmethod
        @modifies_stack([stack_type.typ], [value.IntType()])
        def isempty(vm):
            vm.tnewn(0)
            vm.eq()

        # [stack, value] -> [stack]
        @staticmethod
        @modifies_stack([stack_type.typ, typ], [stack_type.typ])
        def push(vm):
            stack_type.new(vm)
            stack_type.set_val("prev")(vm)
            stack_type.set_val("val")(vm)

        # [stack] -> [value, stack]
        @staticmethod
        @modifies_stack([stack_type.typ], [typ, stack_type.typ])
        def pop(vm):
            vm.dup0()
            stack_type.get("prev")(vm)
            vm.cast(stack_type.typ)
            vm.swap1()
            stack_type.get("val")(vm)

        @staticmethod
        def to_list(stack):
            items = []
            while stack != value.Tuple([]):
                items.append(stack[1])
                stack = stack[0]
            return items

    Stack.typ = stack_type.typ
    return Stack

stack = make_stack_type(value.ValueType())
stack_tup = make_stack_type(value.TupleType())
stack_code = make_stack_type(value.CodePointType())
