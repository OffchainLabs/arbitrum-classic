# Copyright 2019, Offchain Labs, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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

        if func_suffix is not None:
            func.__name__ = "{}_{}".format(func.__name__, func_suffix)

        @functools.wraps(func)
        def wrapper_modifies_stack(vm, *args):
            if not args:
                real_func = func
            else:
                def real_func(vm):
                    func(vm, *args)
                real_func.__name__ = "{}_{}".format(func.__name__, '_'.join(args))

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

        if func_suffix is not None:
            func.__name__ = "{}_{}".format(func.__name__, func_suffix)

        @functools.wraps(func)
        def wrapper_modifies_stack(vm, *args):
            if not args:
                real_func = func
            else:
                def real_func(vm):
                    func(vm, *args)
                real_func.__name__ = "{}_{}".format(func.__name__, '_'.join(args))

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
