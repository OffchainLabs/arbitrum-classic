from ..annotation import modifies_stack
from .struct import Struct
from .queue import queue


def make_closure(func, capture_count):
    bound_types = func.pops[:capture_count]
    param_types = func.pops[capture_count:]
    struct = Struct(
        f"closure[{func.__module__}_{func.__name__}_{capture_count}]",
        [
            ("capture", queue.typ),
        ]
    )
    typ = struct.typ

    class Closure:
        @staticmethod
        @modifies_stack(bound_types, [typ], typ.name)
        def new(vm):
            queue.new(vm)
            for _ in range(capture_count):
                queue.put(vm)
            struct.set_val("capture")(vm)

        @staticmethod
        @modifies_stack([typ] + param_types, func.pushes, typ.name)
        def call(vm):
            struct.get("capture")(vm)
            for typ in bound_types[::-1]:
                queue.get(vm)
                vm.cast(typ)
                vm.swap1()
            vm.pop()
            func(vm)
    Closure.typ = typ
    return Closure
