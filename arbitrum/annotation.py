import functools
from .value import ValueType


def modifies_stack(pop_count, push_count, func_suffix=None):
    def decorator_modifies_stack(func):
        if isinstance(pop_count, int):
            pops = [ValueType() for _ in range(pop_count)]
        else:
            pops = pop_count

        if isinstance(push_count, int):
            pushes = [ValueType() for _ in range(push_count)]
        else:
            pushes = push_count

        if func_suffix:
            func.__name__ = f"{func.__name__}_{func_suffix}"

        @functools.wraps(func)
        def wrapper_modifies_stack(vm, *args):
            if not args:
                real_func = func
            else:
                def real_func(vm):
                    func(vm, *args)
                real_func.__name__ = f"{func.__name__}_{'_'.join(args)}"

            real_func.pops = pops
            real_func.pushes = pushes
            real_func.can_call = True
            real_func.typecheck = True
            return vm.call(real_func)

        wrapper_modifies_stack.pops = pops
        wrapper_modifies_stack.pushes = pushes

        return wrapper_modifies_stack
    return decorator_modifies_stack

def modifies_stack_unchecked(pop_count, push_count, func_suffix=None):
    def decorator_modifies_stack(func):
        if isinstance(pop_count, int):
            pops = [ValueType() for _ in range(pop_count)]
        else:
            pops = pop_count

        if isinstance(push_count, int):
            pushes = [ValueType() for _ in range(push_count)]
        else:
            pushes = push_count

        if func_suffix:
            func.__name__ = f"{func.__name__}_{func_suffix}"

        @functools.wraps(func)
        def wrapper_modifies_stack(vm, *args):
            if not args:
                real_func = func
            else:
                def real_func(vm):
                    func(vm, *args)
                real_func.__name__ = f"{func.__name__}_{'_'.join(args)}"

            real_func.pops = pops
            real_func.pushes = pushes
            real_func.can_call = True
            real_func.typecheck = False
            return vm.call(real_func)

        return wrapper_modifies_stack
    return decorator_modifies_stack


def uncountable_stack(func):
    @functools.wraps(func)
    def wrapper_modifies_stack(vm):
        return vm.call(func)
    func.uncountable = True
    func.can_call = True
    return wrapper_modifies_stack


def noreturn(func):
    func.uncountable = True
    func.can_call = False
    return func
