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

from collections import Counter

import networkx as nx

from . import ast, instructions, value
from .std.bigstruct import BigStruct
from .vm import VM


def build_bst(items):
    if not items:
        return None
    mid = len(items) // 2
    tree = [
        items[mid][0],
        items[mid][1],
        build_bst(items[:mid]),
        build_bst(items[mid + 1 :]),
    ]
    if not tree[2]:
        tree[2] = value.Tuple([0])
    if not tree[3]:
        tree[3] = value.Tuple([0])
    return value.Tuple(tree)


class TempCodePoint:
    def __init__(self, pc):
        self.pc = pc


def replace_code_points(x, block, i):
    if isinstance(x, TempCodePoint):
        return block[x.pc]

    if isinstance(x, ast.AVMLabeledPos):
        if x.pc <= i:
            raise Exception("Immediate codepoint value can only jump forwards")
        return block[x.pc]

    if isinstance(x, value.Tuple):
        return value.Tuple([replace_code_points(v, block, i) for v in x])

    return x


def generate_code_pointers(insns):
    code_points = [None] * len(insns)
    next_hash = b""
    imm_types = Counter()
    for i in range(len(insns) - 1, -1, -1):
        if isinstance(insns[i], ast.BasicOp):
            code_point = value.AVMCodePoint(i, insns[i], next_hash, insns[i].path)
        elif isinstance(insns[i], ast.ImmediateOp):
            immediate = replace_code_points(insns[i].val, code_points, i)
            imm_types[type(immediate)] += 1
            code_point = value.AVMCodePoint(
                i,
                ast.ImmediateOp(insns[i].op, immediate, insns[i].path),
                next_hash,
                insns[i].path,
            )
        else:
            raise Exception(
                "Can't generate code pointer at {} from unexpected value {}".format(
                    i, insns[i]
                )
            )
        next_hash = value.value_hash(code_point)
        code_points[i] = code_point
    return code_points


def check_compiled(insns):
    for i, insn in enumerate(insns):
        if not isinstance(insn, (ast.BasicOp, ast.AVMLabel, ast.AVMUniqueLabel)):
            raise Exception("Found not basic op {} at position {}".format(insn, i))


def build_call_graph(compiled_funcs):
    digraph = nx.DiGraph()
    for func in compiled_funcs:
        counter = CallCounter()
        compiled_funcs[func].modify_ast(counter)
        for dep in counter.call_counts:
            digraph.add_edge(func, dep)
    return digraph


def get_non_recursive(compiled_funcs):
    digraph = build_call_graph(compiled_funcs)

    while True:
        cycles = [cycle for cycle in nx.simple_cycles(digraph)]
        bad_nodes = [x for cycle in cycles for x in cycle]
        if not bad_nodes:
            break
        digraph.remove_nodes_from(bad_nodes)
    return set(x for x in digraph.nodes)


def topological_sort(compiled_funcs):
    digraph = build_call_graph(compiled_funcs)
    cycle_nodes = []
    while True:
        cycles = [cycle for cycle in nx.simple_cycles(digraph)]
        bad_nodes = [x for cycle in cycles for x in cycle]
        if not bad_nodes:
            break
        cycle_nodes += bad_nodes
        digraph.remove_nodes_from(bad_nodes)
    cycle_nodes = sorted(list(set(cycle_nodes)))
    sorted_nodes = list(nx.lexicographical_topological_sort(digraph))
    return cycle_nodes + sorted_nodes


class StaticTracker:
    def __init__(self, code):

        counter = PushCounter()
        code.modify_ast(counter)

        items = []
        for item in counter.push_counts:
            items.append((counter.push_counts[item], item, item))
        items = sorted(items, key=lambda x: (x[0], x[1]), reverse=True)

        self.immediate_pushes = {}
        for item in counter.immediate_push_counts:
            self.immediate_pushes[item] = item
        self.big_struct = BigStruct(items)

    def __getitem__(self, field_name):
        if field_name in self.big_struct:
            return self.big_struct[field_name]

        val = self.immediate_pushes[field_name]
        if val is None:
            raise Exception("Static {} has no value".format(field_name))
        return val

    def __contains__(self, field_name):
        return field_name in self.big_struct or field_name in self.immediate_pushes

    def get_arb_value(self):
        return self.big_struct.initial_val

    def load_static_func(self, val):
        def impl(vm):
            self.big_struct.get(val, vm)

        return impl

    def resolve_label(self, old_val, new_val):
        if old_val in self.big_struct:
            cur_val = self.big_struct[old_val]
            if cur_val != old_val:
                raise Exception(
                    "Tried to resolve label {} with set value {} and new value {}".format(
                        old_val, cur_val, new_val
                    )
                )
            self.big_struct.set_static(old_val, new_val)
        elif old_val in self.immediate_pushes:
            cur_val = self.immediate_pushes[old_val]
            self.immediate_pushes[old_val] = new_val
        else:
            raise Exception(
                "Unhandled label update with old {} and new {}".format(old_val, new_val)
            )


class ASTTransformer:
    def __call__(self, op):
        if isinstance(op, ast.IfElseStatement):
            return self.transform_ifelse(op)
        if isinstance(op, ast.IfStatement):
            return self.transform_if(op)
        if isinstance(op, ast.WhileStatement):
            return self.transform_while(op)
        if isinstance(op, ast.BlockStatement):
            return self.transform_block(op)
        if isinstance(op, ast.CallStatement):
            return self.transform_call(op)
        if isinstance(op, ast.ImmediateOp):
            return self.transform_immediate(op)
        if isinstance(op, ast.IndirectPushStatement):
            return self.transform_indirect_push(op)
        if isinstance(op, ast.BasicOp):
            return self.transform_basic(op)
        if isinstance(op, ast.AVMLabel):
            return self.transform_label(op)
        if isinstance(op, ast.AVMUniqueLabel):
            return self.transform_unique_label(op)
        if isinstance(op, ast.FuncDefinition):
            return self.transform_func_definition(op)
        if isinstance(op, ast.CastStatement):
            return self.transform_cast(op)
        if isinstance(op, ast.SetErrorHandlerFunctionStatement):
            return self.transform_set_error_handler(op)

        raise Exception("Unhandled AST Type {}".format(type(op)))

    def transform_ifelse(self, op):
        return op

    def transform_if(self, op):
        return op

    def transform_while(self, op):
        return op

    def transform_block(self, op):
        return op

    def transform_call(self, op):
        return op

    def transform_indirect_push(self, op):
        return op

    def transform_immediate(self, op):
        return op

    def transform_basic(self, op):
        return op

    def transform_label(self, op):
        return op

    def transform_unique_label(self, op):
        return op

    def transform_func_definition(self, op):
        return op

    def transform_cast(self, op):
        return op

    def transform_set_error_handler(self, op):
        return op


class FlowControlTransformer(ASTTransformer):
    def __init__(self, label_gen):
        super(FlowControlTransformer, self).__init__()
        self.label_gen = label_gen

    def transform_ifelse(self, op):
        def impl(vm):
            mid_label = self.label_gen.generate_unique_label("ifelse_mid")
            end_label = self.label_gen.generate_unique_label("ifelse_end")
            vm.push(mid_label)
            vm.cjump()
            vm.block.append(ast.add_label_to_ast(op.false_code, "IfElseStatementFalse"))
            vm.jump_direct(end_label)
            vm.set_label(mid_label)
            vm.block.append(ast.add_label_to_ast(op.true_code, "IfElseStatementTrue"))
            vm.set_label(end_label)

        return ast.add_label_to_ast(compile_block(impl), "IfElseStatement")

    def transform_if(self, op):
        def impl(vm):
            end_label = self.label_gen.generate_unique_label("if_end")
            vm.iszero()
            vm.push(end_label)
            vm.cjump()
            vm.block.append(ast.add_label_to_ast(op.true_code, "IfStatementTrue"))
            vm.set_label(end_label)

        return ast.add_label_to_ast(compile_block(impl), "IfStatement")

    def transform_while(self, op):
        def impl(vm):
            bottom_label = self.label_gen.generate_unique_label("while_bottom")
            vm.pcpush()
            vm.auxpush()
            vm.block.append(ast.add_label_to_ast(op.cond_code, "WhileStatementCond"))
            vm.iszero()
            vm.push(bottom_label)
            vm.cjump()
            vm.block.append(ast.add_label_to_ast(op.body_code, "WhileStatementBody"))
            vm.auxpop()
            vm.jump()
            vm.set_label(bottom_label)
            vm.auxpop()
            vm.pop()

        return ast.add_label_to_ast(compile_block(impl), "WhileStatement")


class InlineCallTransformer(ASTTransformer):
    def __init__(self, definition):
        super(InlineCallTransformer, self).__init__()
        self.definition = definition

    def transform_call(self, op):
        if op.func_name == self.definition.name:
            return self.definition.code.clone()
        return op


class PushCounter(ASTTransformer):
    def __init__(self):
        super(PushCounter, self).__init__()
        self.push_counts = Counter()
        self.immediate_push_counts = Counter()

    def transform_indirect_push(self, op):
        self.push_counts[op.val] += 1
        return op

    def transform_immediate(self, op):
        self.immediate_push_counts[op.val] += 1
        return op


class CallCounter(ASTTransformer):
    def __init__(self):
        super(CallCounter, self).__init__()
        self.call_counts = Counter()

    def transform_call(self, op):
        self.call_counts[op.func_name] += 1
        return op


class CastRemover(ASTTransformer):
    def __init__(self):
        super(CastRemover, self).__init__()

    def transform_cast(self, op):
        return ast.BlockStatement([])


class ForwardImmediateTransformer(ASTTransformer):
    def __init__(self):
        super(ForwardImmediateTransformer, self).__init__()
        self.seen_labels = set()

    def transform_label(self, op):
        self.seen_labels.add(op)
        return op

    def transform_unique_label(self, op):
        self.seen_labels.add(op)
        return op

    def transform_immediate(self, op):
        def contains_matching_label(val):
            if isinstance(val, value.Tuple):
                return any(contains_matching_label(val) for val in val.val)
            return val in self.seen_labels

        if isinstance(op.val, str) or contains_matching_label(op.val):
            return ast.BlockStatement([ast.IndirectPushStatement(op.val), op.op])
        return op


class PushTransformer(ASTTransformer):
    def __init__(self, static_tracker):
        super(PushTransformer, self).__init__()
        self.static_tracker = static_tracker

    def transform_indirect_push(self, op):
        def replace_push(val):
            def impl(vm):
                vm.spush()
                self.static_tracker.load_static_func(val)(vm)

            return impl

        return ast.add_label_to_ast(
            compile_block(replace_push(op.val)), "Push({})".format(op.val)
        )


class CallTransformer(ASTTransformer):
    def __init__(self, label_gen):
        super(CallTransformer, self).__init__()
        self.label_gen = label_gen

    def transform_call(self, op):
        def impl(vm):
            return_label = self.label_gen.generate_unique_label(
                op.func_name + "_return"
            )
            vm.push(return_label)
            vm.auxpush()
            vm.jump_direct(ast.AVMLabel(op.func_name))
            vm.set_label(return_label)

        return ast.add_label_to_ast(compile_block(impl), op)

    def transform_set_error_handler(self, op):
        def impl(vm):
            vm.push(ast.AVMLabel(op.func_name))
            vm.errset()

        return ast.add_label_to_ast(
            compile_block(impl), "SetErrorHandlerFunction({})".format(op.func_name)
        )


class FuncTransformer(ASTTransformer):
    def __init__(self):
        super(FuncTransformer, self).__init__()

    def transform_func_definition(self, op):
        if op.is_callable:

            def impl(vm):
                vm.set_label(ast.AVMLabel(op.name))
                vm.block.append(op.code)
                vm.auxpop()
                vm.jump()

        else:

            def impl(vm):
                vm.set_label(ast.AVMLabel(op.name))
                vm.block.append(op.code)

        return ast.add_label_to_ast(
            compile_block(impl), "FuncDefinition({})".format(op.name)
        )


class LabelGenerator:
    def __init__(self):
        self.label_id = 0

    def generate_unique_label(self, name=""):
        if name == "":
            label = ast.AVMUniqueLabel("custom_label_" + str(self.label_id))
        else:
            label = ast.AVMUniqueLabel(name + "_" + str(self.label_id))
        self.label_id += 1
        return label


def compile_block(func):
    compiler = VMCompiler()
    func(compiler)
    return ast.BlockStatement(compiler.block)


def transform_code_block(code, code_pass, op_count=1):
    i = 0
    while i + op_count - 1 < len(code):
        ops = code[i : i + op_count]
        if op_count > 1:
            new_code = code_pass(ops, i)
        else:
            new_code = code_pass(ops[0], i)
        code[i : i + op_count] = new_code
        if new_code == ops:
            i += 1


def remove_nop_swaps(ops, i):
    if ops[0] == ops[1] and (
        ops[0].op_code == instructions.OPS["swap1"]
        or ops[0].op_code == instructions.OPS["swap2"]
    ):
        return []

    return ops


def compress_pushes(ops, i):
    if (
        isinstance(ops[0], ast.ImmediateOp)
        and ops[0].get_op() == instructions.OPS["nop"]
        and isinstance(ops[1], ast.BasicOp)
        and ops[1].get_op() != instructions.OPS["pcpush"]
    ):
        return [ast.ImmediateOp(ops[1], ops[0].val, ops[1].path)]
    else:
        return ops


def resolve_labels(static_tracker):
    def impl(op, i):
        if not isinstance(op, (ast.AVMLabel, ast.AVMUniqueLabel)):
            return [op]
        static_tracker.resolve_label(op, ast.AVMLabeledPos(op.name, i))
        return []

    return impl


def resolve_immediate_ops(static_tracker):
    def contains_label(val):
        if isinstance(val, value.Tuple):
            return any(contains_label(val) for val in val.val)
        else:
            return isinstance(val, (ast.AVMLabel, ast.AVMUniqueLabel))

    def transform_val(val):
        if isinstance(val, value.Tuple):
            return value.Tuple([transform_val(v) for v in val.val])
        elif val in static_tracker:
            new_val = static_tracker[val]
            if isinstance(new_val, (ast.AVMLabel, ast.AVMUniqueLabel)):
                raise Exception("Can't resolve {}, got {}".format(val, new_val))
            return new_val
        else:
            return val

    def impl(op, i):
        if isinstance(op, ast.ImmediateOp):
            if contains_label(op.val):
                return [ast.ImmediateOp(op.op, transform_val(op.val), op.path)]
            else:
                return [op]
        else:
            return [op]

    return impl


class VMCompiler:
    def __init__(self):
        def create_op(op_name, op_code):
            def impl(self):
                self.block.append(ast.BasicOp(op_code))

            return impl

        self.block = []
        for (op_name, op_code, pop_count, push_count) in instructions.OP_CODES:
            if op_name != "push":
                setattr(VMCompiler, op_name, create_op(op_name, op_code))

    def set_label(self, val):
        self.block.append(val)

    def call(self, func):
        if not hasattr(func, "can_call") or not func.can_call:
            raise Exception(
                "Tried to call uncallable function {}".format(func.__name__)
            )
        self.block.append(ast.CallStatement(func))

    def set_exception_handler(self, func):
        self.block.append(ast.SetErrorHandlerFunctionStatement(func))

    def clear_exception_handler(self):
        self.push(value.ERROR_CODE_POINT)
        self.errset()

    def while_loop(self, cond_block, body_block):
        self.block.append(
            ast.WhileStatement(compile_block(cond_block), compile_block(body_block))
        )

    def ifelse(self, true_block, false_block=None):
        if false_block:
            self.block.append(
                ast.IfElseStatement(
                    compile_block(true_block), compile_block(false_block)
                )
            )
        else:
            self.block.append(ast.IfStatement(compile_block(true_block)))

    def cast(self, typ):
        self.block.append(ast.CastStatement(typ))

    def push(self, val):
        self.block.append(ast.ImmediateOp(ast.BasicOp(instructions.OPS["nop"]), val))

    def tnewn(self, size):
        self.push(value.Tuple([value.Tuple([]) for i in range(size)]))

    def tsetn(self, val):
        self.block.append(ast.ImmediateOp(ast.BasicOp(instructions.OPS["tset"]), val))

    def tgetn(self, val):
        self.block.append(ast.ImmediateOp(ast.BasicOp(instructions.OPS["tget"]), val))

    def jump_direct(self, location):
        self.block.append(
            ast.ImmediateOp(ast.BasicOp(instructions.OPS["jump"]), location)
        )


def expectation_dependencies(expectations):
    expectation_dependencies = set()
    for x in expectations:
        expectation_dependencies |= x.free_symbols
    return expectation_dependencies


def flatten_block(op):
    if not isinstance(op, ast.BlockStatement):
        return [op]
    ret = []
    for sub_op in op.code:
        ret += flatten_block(sub_op)
    for sub_op in ret:
        sub_op.path = list(op.path) + sub_op.path
    return ret


def optimize_program(compiled_funcs):
    # use cycle checking to figure out which functions are safe to inline
    non_recursive = get_non_recursive(compiled_funcs)
    non_recursive = [x for x in non_recursive if compiled_funcs[x].is_callable]
    # IMPORTANT: Inling requires code cloning which only currently works
    #            if the ast in the code includes no labels.
    # # count how many times each function is called
    counter = CallCounter()
    for func in compiled_funcs:
        compiled_funcs[func].modify_ast(counter)

    # inline non-recursive functions that are called a single time
    single_call = [
        x
        for x in counter.call_counts
        if counter.call_counts[x] == 1 and x in non_recursive
    ]
    for single_func in single_call:
        func_to_inline = compiled_funcs[single_func]
        for func in compiled_funcs:
            compiled_funcs[func] = compiled_funcs[func].modify_ast(
                InlineCallTransformer(func_to_inline)
            )
        non_recursive.remove(single_func)
        del compiled_funcs[single_func]

    # inline short non-recursive functions
    while True:
        if not non_recursive:
            break
        shortest_non_recursive = min(
            non_recursive, key=lambda func: len(compiled_funcs[func])
        )
        if len(compiled_funcs[shortest_non_recursive]) >= 150:
            break
        func_to_inline = compiled_funcs[shortest_non_recursive]
        del compiled_funcs[shortest_non_recursive]
        for func in compiled_funcs:
            compiled_funcs[func] = compiled_funcs[func].modify_ast(
                InlineCallTransformer(func_to_inline)
            )
        non_recursive.remove(shortest_non_recursive)


def verify_stackmods(compiled_funcs):
    # Verify manual stack count labeling
    for func in compiled_funcs:
        if func == "MAIN_FUNC":
            continue
        if hasattr(compiled_funcs[func].func, "uncountable"):
            continue
        if not compiled_funcs[func].is_callable:
            continue
        mods, expects = compiled_funcs[func].code.stack_mod()
        uncountable = False
        incorrect = False
        for x in expects:
            if x[0] == "eq":
                if x[1] != x[2]:
                    incorrect = True
                    break
            elif x[0] == "invalid":
                uncountable = True
                break
            else:
                raise Exception("Unhandled expectation type")

        if uncountable:
            raise Exception(
                "{} calculated as uncountable, but isn't labeled that way".format(func)
            )
        elif incorrect:
            raise Exception(
                "Function '{}'' violates constraints {}".format(func, expects)
            )
        else:
            if not hasattr(compiled_funcs[func].func, "pops"):
                raise Exception(
                    "{} calculated {} but wasn't labeled with count".format(
                        func, mods["pop"]
                    )
                )
            if mods["pop"] != len(compiled_funcs[func].func.pops):
                raise Exception(
                    "{} calculated {} pops but was labeled with {}".format(
                        func, mods["pop"], len(compiled_funcs[func].func.pops)
                    )
                )
            if mods["push"] != len(compiled_funcs[func].func.pushes):
                raise Exception(
                    "{} calculated {} pushes but was labeled with {}".format(
                        func, mods["push"], len(compiled_funcs[func].func.pushes)
                    )
                )


def compile_program(initialization, body, should_optimize=True):
    compiled_funcs = {}

    # Iteratively resolve all function calls
    seen_funcs = set()
    funcs_to_search = [ast.FuncDefinition("MAIN_FUNC", None, body, False)]
    label_gen = LabelGenerator()
    while funcs_to_search:
        new_funcs = []

        def find_calls(op):
            if isinstance(
                op, (ast.CallStatement, ast.SetErrorHandlerFunctionStatement)
            ):
                if op.func_name not in seen_funcs:
                    new_funcs.append(
                        ast.FuncDefinition(
                            op.func_name,
                            op.func,
                            compile_block(op.func),
                            op.is_callable,
                        )
                    )
                    seen_funcs.add(op.func_name)

        for func in funcs_to_search:
            func.traverse_ast(find_calls)
            compiled_funcs[func.name] = func

        funcs_to_search = new_funcs

    # print(list(compiled_funcs))
    verify_stackmods(compiled_funcs)

    for func in compiled_funcs:
        if func == "MAIN_FUNC":
            continue
        if compiled_funcs[func].can_typecheck():
            compiled_funcs[func].typecheck()

    for func in compiled_funcs:
        compiled_funcs[func] = compiled_funcs[func].modify_ast(CastRemover())

    if should_optimize:
        optimize_program(compiled_funcs)

    for func in compiled_funcs:
        compiled_funcs[func] = compiled_funcs[func].modify_ast(
            FlowControlTransformer(label_gen)
        )

    function_order = topological_sort(compiled_funcs)
    other_funcs = sorted([x for x in compiled_funcs if x not in function_order])
    function_order += other_funcs

    # transform remaining calls into jumps
    for func in compiled_funcs:
        compiled_funcs[func] = compiled_funcs[func].modify_ast(
            CallTransformer(label_gen)
        )

    # merge all functions into a single code block
    main_code = compiled_funcs["MAIN_FUNC"].code.code
    del compiled_funcs["MAIN_FUNC"]

    full_code = initialization.code
    for func in compiled_funcs:
        compiled_funcs[func] = compiled_funcs[func].modify_ast(FuncTransformer())

    for func in [x for x in function_order if x != "MAIN_FUNC"]:
        full_code.append(compiled_funcs[func])
    full_code += main_code
    full_code = ast.BlockStatement(full_code)
    # replace pushes with accesses to the static
    full_code = full_code.modify_ast(ForwardImmediateTransformer())
    static_tracker = StaticTracker(full_code)
    full_code = full_code.modify_ast(PushTransformer(static_tracker))
    full_code = flatten_block(full_code)

    if should_optimize:
        transform_code_block(full_code, remove_nop_swaps, 2)
        transform_code_block(full_code, compress_pushes, 2)

    # replace all labels with code points
    # Warning: After this pass the number of instructions can't change
    transform_code_block(full_code, resolve_labels(static_tracker))
    transform_code_block(full_code, resolve_immediate_ops(static_tracker))
    code_pointers = generate_code_pointers(full_code)
    vm = VM(code_pointers)
    vm.static = static_tracker.get_arb_value()
    vm.static = replace_code_points(vm.static, code_pointers, 0)
    return vm
