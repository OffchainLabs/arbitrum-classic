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

from . import instructions
from . import value

PUSH_WEIGHT = 5


class ASTNode:
    def __init__(self, path):
        if path is None:
            path = []
        self.path = path

    def add_node(self, typ):
        self.path.append(typ)


def add_label_to_ast(node, label):
    node.add_node(label)
    return node
    # def impl(op):
    #     op.add_node(label)
    #     return op
    # return node.modify_ast(impl)


class BlockStatement(ASTNode):
    def __init__(self, code, path=None):
        super(BlockStatement, self).__init__(path)
        assert isinstance(code, list)
        self.code = code

        self.stack_mod_saved = None

    def clone(self):
        return BlockStatement(
            [op.clone() for op in self.code],
            list(self.path)
        )

    def __repr__(self):
        res = "["
        for i, op in enumerate(self.code):
            res += str(op)
            if i < len(self.code) - 1:
                res += ", "
        res += "]"
        return res

    def __len__(self):
        return sum(len(op) for op in self.code)

    def stack_mod(self):
        if self.stack_mod_saved is None:
            expectations = []
            net_stacks = [0]
            net = 0
            for op in self.code:
                mods, expect = op.stack_mod()
                expectations += expect
                net -= mods["pop"]
                net_stacks.append(net)
                net += mods["push"]
            pop_count = -min(net_stacks)
            push_count = max(net + pop_count, 0)
            self.stack_mod_saved = (
                {
                    "pop": pop_count,
                    "push": push_count
                }, expectations
            )
        return self.stack_mod_saved

    def typecheck(self, stack):
        for op in self.code:
            op.typecheck(stack)

    def traverse_ast(self, func):
        func(self)
        for op in self.code:
            op.traverse_ast(func)

    def modify_ast(self, func):
        self.code = [op.modify_ast(func) for op in self.code]
        return func(self)


class IfElseStatement(ASTNode):
    def __init__(self, true_code, false_code, path=None):
        super(IfElseStatement, self).__init__(path)
        self.true_code = true_code
        self.false_code = false_code

    def clone(self):
        return IfElseStatement(
            self.true_code.clone(),
            self.false_code.clone(),
            list(self.path)
        )

    def __repr__(self):
        return f"IfElse({self.true_code}, {self.false_code})"

    def __len__(self):
        return len(self.true_code) + len(self.false_code) + 2 + 2*PUSH_WEIGHT

    def stack_mod(self):
        true_mods, true_expectations = self.true_code.stack_mod()
        false_mods, false_expectations = self.false_code.stack_mod()
        expectations = true_expectations + false_expectations
        expectations.append((
            'eq',
            true_mods["push"] - true_mods["pop"],
            false_mods["push"] - false_mods["pop"]
        ))
        return {
            "pop": max(true_mods["pop"], false_mods["pop"]) + 1,
            "push": max(true_mods["push"], false_mods["push"])
        }, expectations

    def typecheck(self, stack):
        stack.pop(value.IntType())
        temp = stack.clone()
        self.true_code.typecheck(stack)
        self.false_code.typecheck(temp)
        stack.merge(temp)

    def traverse_ast(self, func):
        func(self)
        self.true_code.traverse_ast(func)
        self.false_code.traverse_ast(func)

    def modify_ast(self, func):
        self.true_code = self.true_code.modify_ast(func)
        self.false_code = self.false_code.modify_ast(func)
        return func(self)


class CastStatement(ASTNode):
    def __init__(self, typ, path=None):
        super(CastStatement, self).__init__(path)
        self.typ = typ

    def clone(self):
        return CastStatement(self.typ, list(self.path))

    def __repr__(self):
        return f"CastStatement({self.typ})"

    def __len__(self):
        return 0

    def stack_mod(self):
        return {
            "pop": 1,
            "push": 1
        }, []

    def typecheck(self, stack):
        stack.pop()
        stack.push(self.typ)

    def traverse_ast(self, func):
        pass

    def modify_ast(self, func):
        return func(self)


class IfStatement(ASTNode):
    def __init__(self, true_code, path=None):
        super(IfStatement, self).__init__(path)
        self.true_code = true_code

    def clone(self):
        return IfStatement(self.true_code.clone(), list(self.path))

    def __repr__(self):
        return f"If({self.true_code})"

    def __len__(self):
        return len(self.true_code) + 2 + PUSH_WEIGHT

    def stack_mod(self):
        true_mods, true_expectations = self.true_code.stack_mod()
        expectations = list(true_expectations)
        expectations.append(('eq', true_mods["push"] - true_mods["pop"], 0))
        return {
            "pop": true_mods["pop"] + 1,
            "push": true_mods["push"]
        }, expectations

    def typecheck(self, stack):
        stack.pop(value.IntType())
        temp = stack.clone()
        self.true_code.typecheck(stack)
        stack.merge(temp)

    def traverse_ast(self, func):
        func(self)
        self.true_code.traverse_ast(func)

    def modify_ast(self, func):
        self.true_code = self.true_code.modify_ast(func)
        return func(self)


class WhileStatement(ASTNode):
    def __init__(self, cond_code, body_code, path=None):
        super(WhileStatement, self).__init__(path)
        self.cond_code = cond_code
        self.body_code = body_code

    def clone(self):
        return WhileStatement(
            self.cond_code.clone(),
            self.body_code.clone(),
            list(self.path)
        )

    def __repr__(self):
        return f"WhileStatement({self.cond_code}, {self.body_code})"

    def __len__(self):
        return len(self.cond_code) + len(self.body_code) + 3 + 2*PUSH_WEIGHT

    def stack_mod(self):
        cmod, cond_expectations = self.cond_code.stack_mod()
        bmod, body_expectation = self.body_code.stack_mod()
        expectations = cond_expectations + body_expectation
        expectations.append((
            'eq',
            cmod["push"] - cmod["pop"] - 1 + bmod["push"] - bmod["pop"],
            0,
            f"while_loop({cmod['pop']}, {cmod['push']}, {bmod['pop']}, {bmod['push']})"
        ))
        pop_count = max(
            cmod["pop"],
            bmod["pop"] + 1 - cmod["push"] + cmod["pop"]
        )
        mods = {
            "push": pop_count + cmod["push"] - cmod["pop"] - 1,
            "pop": pop_count
        }
        return mods, expectations

    def typecheck(self, stack):
        temp = stack.clone()
        self.cond_code.typecheck(stack)
        stack.pop(value.IntType())
        self.body_code.typecheck(stack)
        temp.merge(stack)

    def traverse_ast(self, func):
        func(self)
        self.cond_code.traverse_ast(func)
        self.body_code.traverse_ast(func)

    def modify_ast(self, func):
        self.cond_code = self.cond_code.modify_ast(func)
        self.body_code = self.body_code.modify_ast(func)
        return func(self)


class FuncDefinition(ASTNode):
    def __init__(self, name, func, code, is_callable, path=None):
        super(FuncDefinition, self).__init__(path)
        self.name = name
        self.func = func
        self.code = code
        self.is_callable = is_callable

        if not isinstance(code, ASTNode):
            raise Exception("Func definition expects ASTNode for code")

    def clone(self):
        raise Exception("Func definitions aren't clonable")

    def __repr__(self):
        return f"FuncDefinition({self.code})"

    def __len__(self):
        return len(self.code)

    def can_typecheck(self):
        return hasattr(self.func, "pops") and hasattr(self.func, "pushes") and self.func.typecheck

    def typecheck(self):
        stack = value.TypeStack()
        for typ in self.func.pops[::-1]:
            stack.push(typ)
        try:
            self.code.typecheck(stack)
        except Exception as err:
            raise Exception(f"Error typechecking {self.name} body: {err}")
        
        try:
            for typ in self.func.pushes:
                stack.pop(typ)
        except Exception as err:
            raise Exception(f"Error typechecking {self.name} return vals: {err}")

    def traverse_ast(self, func):
        func(self)
        self.code.traverse_ast(func)

    def modify_ast(self, func):
        self.code = self.code.modify_ast(func)
        return func(self)


class CallStatement(ASTNode):
    def __init__(self, func, path=None):
        super(CallStatement, self).__init__(path)
        self.func = func
        self.func_name = f"{func.__module__}.{func.__name__}"
        self.is_callable = True
        if hasattr(self.func, "uncountable"):
            self.mods = {"pop": 0, "push": 0}, [('invalid',)]

        elif (
                not hasattr(self.func, "pushes") or
                not hasattr(self.func, "pops")
        ):
            raise Exception(f"Call {self.func_name} has unknown stack mods")
        else:
            self.mods = {
                "pop": len(self.func.pops),
                "push": len(self.func.pushes)
            }, []
            self.pops = self.func.pops
            self.pushes = self.func.pushes

    def clone(self):
        return CallStatement(self.func, list(self.path))

    def __repr__(self):
        return f"Call({self.func_name})"

    def __len__(self):
        # Approximation
        return 12

    def stack_mod(self):
        return self.mods[0], self.mods[1]

    def typecheck(self, stack):
        try:
            for typ in self.func.pops:
                stack.pop(typ)
            for typ in self.func.pushes[::-1]:
                stack.push(typ)
        except Exception as err:
            raise Exception(f"Type error calling func {self.func_name}: {err}")

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)


class SetErrorHandlerFunctionStatement(ASTNode):
    def __init__(self, func, path=None):
        super(SetErrorHandlerFunctionStatement, self).__init__(path)
        self.func = func
        self.func_name = f"{func.__module__}.{func.__name__}"
        self.is_callable = False

    def clone(self):
        return SetErrorHandlerFunctionStatement(self.func, list(self.path))

    def __repr__(self):
        return f"SetErrorHandlerFunction({self.func_name})"

    def __len__(self):
        # Approximation
        return 1

    def stack_mod(self):
        return {
            "pop": 0,
            "push": 0
        }

    def typecheck(self, stack):
        pass

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)



class IndirectPushStatement(ASTNode):
    def __init__(self, val, path=None):
        super(IndirectPushStatement, self).__init__(path)
        self.val = val

    def clone(self):
        return IndirectPushStatement(self.val, list(self.path))

    def __len__(self):
        # Approximation
        return 6

    def stack_mod(self):
        return {"pop": 0, "push": 1}, []

    def typecheck(self, stack):
        if isinstance(self.val, (AVMLabel,)):
            typ = value.CodePointType()
        else:
            typ = self.val
        stack.push(typ)

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)

    def __repr__(self):
        return f"Push({self.val})"


class AVMLabel(ASTNode):
    def __init__(self, name, path=None):
        super(AVMLabel, self).__init__(path)
        self.name = name
        # print("Label", name)

    def clone(self):
        raise Exception(f"You can't clone a label '{self.name}'")

    def __len__(self):
        return 0

    def stack_mod(self):
        return {"pop": 0, "push": 0}, []

    def typecheck(self, stack):
        pass

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)

    def __repr__(self):
        return f"AVMLabel({self.name})"

    def __eq__(self, other):
        if not isinstance(other, AVMLabel):
            return False
        return self.name == other.name

    def __ne__(self, other):
        return self.name != other.name

    def __hash__(self):
        assert isinstance(self.name, str)
        return self.name.__hash__()


class AVMUniqueLabel(ASTNode):
    def __init__(self, name, path=None):
        super(AVMUniqueLabel, self).__init__(path)
        self.name = name

    def clone(self):
        raise Exception(f"You can't clone a label '{self.name}'")

    def __len__(self):
        return 0

    def stack_mod(self):
        return {"pop": 0, "push": 0}, []

    def typecheck(self, stack):
        pass

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)

    def __repr__(self):
        return f"AVMUniqueLabel({self.name})"

    def __eq__(self, other):
        if not isinstance(other, AVMUniqueLabel):
            return False

        return self.name == other.name

    def __ne__(self, other):
        return self.name != other.name

    def __hash__(self):
        assert isinstance(self.name, str)
        return self.name.__hash__()


class AVMLabeledPos(ASTNode):
    def __init__(self, name, pc, path=None):
        super(AVMLabeledPos, self).__init__(path)
        self.name = name
        self.pc = pc

    def clone(self):
        return AVMLabeledPos(self.name, self.pc, list(self.path))

    def __len__(self):
        return 0

    def stack_mod(self):
        return {"pop": 0, "push": 0}, []

    def typecheck(self, stack):
        pass

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)

    def __repr__(self):
        return f"AVMLabeledPos({self.name, self.pc})"

    def __eq__(self, other):
        if not isinstance(other, AVMLabeledPos):
            return False
        return self.pc == other.pc


class AVMLabeledCodePoint(ASTNode):
    def __init__(self, name, pc, path=None):
        super(AVMLabeledCodePoint, self).__init__(path)
        self.name = name
        self.pc = pc

    def clone(self):
        return AVMLabeledCodePoint(self.name, self.pc, list(self.path))

    def __len__(self):
        return 0

    def stack_mod(self):
        return {"pop": 0, "push": 0}, []

    def typecheck(self, stack):
        pass

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)

    def __repr__(self):
        return f"AVMLabeledCodePoint({self.name, self.pc})"

    def __eq__(self, other):
        if not isinstance(other, AVMLabeledCodePoint):
            return False
        return self.pc == other.pc


class BasicOp(ASTNode):
    def __init__(self, op_code, path=None):
        super(BasicOp, self).__init__(path)
        self.op_code = op_code

    def clone(self):
        return BasicOp(self.op_code, list(self.path))

    def __len__(self):
        return 1

    def __repr__(self):
        return instructions.OP_NAMES[self.op_code]

    def get_op(self):
        return self.op_code

    def stack_mod(self):
        info = instructions.OF_INFO[self.op_code]
        mod = {
            "pop": len(info["pop"]),
            "push": len(info["push"])
        }
        if (
                instructions.OP_NAMES[self.op_code] == "jump" or
                instructions.OP_NAMES[self.op_code] == "cjump"
        ):
            return mod, [('invalid',)]
        if instructions.OP_NAMES[self.op_code] == "halt":
            return mod, [('invalid',)]
        return mod, []

    def typecheck(self, stack):
        try:
            name = instructions.OP_NAMES[self.op_code]
            if name == "auxpush":
                stack.push_aux(stack.pop())
            elif name == "auxpop":
                stack.push(stack.pop_aux())
            elif name == "dup0":
                instructions.dup0(stack)
            elif name == "dup1":
                instructions.dup1(stack)
            elif name == "dup2":
                instructions.dup2(stack)
            elif name == "swap1":
                instructions.swap1(stack)
            elif name == "swap2":
                instructions.swap2(stack)
            elif name == "tnew":
                size = stack.pop(value.IntType())
                if isinstance(size, int):
                    stack.push(value.Tuple([value.Tuple([]) for i in range(size)]))
                else:
                    stack.push(value.TupleType())
            elif name == "tget":
                index = stack.pop(value.IntType())
                tup = stack.pop(value.TupleType())
                if isinstance(index, int) and not tup.has_member_at_index(index):
                    raise Exception(f"Tried to get index {index} from tuple {tup}")
                stack.push(tup.get_tup(index))
            elif name[:4] == "tget":
                instructions.tgetn(stack, int(name[4:]))
            elif name == "tset":
                instructions.tset(stack)
            elif name[:4] == "tset":
                instructions.tsetn(stack, int(name[4:]))
            else:
                info = instructions.OF_INFO[self.op_code]
                for pop in info["pop"]:
                    stack.pop(pop)
                for push in info["push"]:
                    stack.push(push)
        except Exception as err:
            raise Exception(
                f"Exception typechecking {instructions.OP_NAMES[self.op_code]}: {err}"
            )

    def __eq__(self, other):
        if not isinstance(other, BasicOp):
            return False
        return self.op_code == other.op_code

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)


class ImmediateOp(ASTNode):
    def __init__(self, op, val, path=None):
        super(ImmediateOp, self).__init__(path)
        self.op = op
        self.val = val

    def clone(self):
        return ImmediateOp(self.op, self.val, list(self.path))

    def __len__(self):
        # Approximation
        return 1

    def __repr__(self):
        return f"Immediate({self.op, self.val})"

    def get_op(self):
        return self.op.get_op()

    def stack_mod(self):
        op_mod, constraints = self.op.stack_mod()
        if op_mod["pop"] > 0:
            op_mod["pop"] -= 1
        else:
            op_mod["push"] += 1
        return op_mod, constraints

    def typecheck(self, stack):
        if isinstance(self.val, (AVMLabel,)):
            typ = value.CodePointType()
        else:
            typ = self.val
        stack.push(typ)
        self.op.typecheck(stack)

    def traverse_ast(self, func):
        func(self)

    def modify_ast(self, func):
        return func(self)
