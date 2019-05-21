
# Credit to https://github.com/ethereum/pyethereum/blob/master/ethereum/vm.py
# for EVM-like implementation details

from eth_utils import big_endian_to_int
from . import instructions
from .ast import AVMLabeledCodePoint
from . import value

TT256 = 2 ** 256
TT256M1 = 2 ** 256 - 1
TT255 = 2 ** 255


def to_signed(i):
    return i if i < TT255 else i - TT256


class VMBlocked(Exception):
    """VM tried to run opcode that blocks"""
    pass


class VMBlockedAdvance(Exception):
    """VM tried to run opcode that blocks"""
    pass


class VMEnv:
    def __init__(self):
        self.messages = value.Tuple([])
        self.pending_messages = value.Tuple([])
        self.time_bounds = value.Tuple([0, 100000000])

    def send_message(self, message):
        self.pending_messages = value.Tuple([
            0,
            self.pending_messages,
            value.Tuple(message)
        ])

    def deliver_pending(self):
        self.messages = value.Tuple([
            1,
            self.messages,
            self.pending_messages
        ])
        self.pending_messages = value.Tuple([])


class Stack:
    def __init__(self, items=None):
        if items is None:
            items = []
        self.items = items

    def __len__(self):
        return len(self.items)

    def __getitem__(self, *args):
        return self.items.__getitem__(*args)

    def push(self, val):
        self.items.insert(0, val)

    def pop(self, typehint=None):
        val = self.items[0]
        if typehint and not typehint.accepts(val):
            raise Exception(f"Pop expected {typehint}, but got {val}")
        del self.items[0]
        return val

    def peak(self):
        return self.items[0]


class BasicVM:
    def __init__(self):
        self.pc = value.AVMCodePoint(0, 0, b'')
        self.stack = Stack()
        self.aux_stack = Stack()
        self.register = value.Tuple([])
        self.static = value.Tuple([])
        self.err_handler = None
        self.atomic_count = 0

        self.env = VMEnv()
        self.halted = False
        self.sent_messages = []
        self.logs = []

    def log(self):
        self.logs.append(self.stack.pop())

    def breakpoint(self):
        raise VMBlockedAdvance()

    def debug(self):
        print("Debug:\nStack:", self.stack[:], "\nAux:", self.aux_stack[:])

    def push(self, val):
        if isinstance(val, list):
            self.stack.push(value.Tuple(val))
        else:
            self.stack.push(val)

    def inbox(self):
        if self.stack.peak() == self.env.messages:
            raise VMBlocked()

        self.stack.pop()
        self.stack.push(self.env.messages)

    def send(self):
        msg = self.stack.pop()
        self.sent_messages.append(msg)

    def nbsend(self):
        msg = self.stack.pop()
        self.sent_messages.append(msg)
        self.stack.push(1)

    def auxpush(self):
        item = self.stack.pop()
        self.aux_stack.push(item)

    def auxpop(self):
        item = self.aux_stack.pop()
        self.stack.push(item)

    def jump(self):
        dest = self.stack.pop()
        # if isinstance(dest, AVMLabeledCodePoint):
        #     print("Jumping to", dest)
        if isinstance(dest, value.AVMCodePoint):
            pc = dest
        elif isinstance(dest, AVMLabeledCodePoint):
            pc = dest.pc
        else:
            raise Exception("Jump insn requires codepoint but recieved " + str(dest))
        self.pc = pc

    def errpush(self):
        self.stack.push(self.err_handler)

    def errset(self):
        err_handler = self.stack.pop()
        if isinstance(err_handler, value.AVMCodePoint):
            pc = err_handler
        elif isinstance(err_handler, AVMLabeledCodePoint):
            pc = err_handler.pc

        self.err_handler = pc

    def cjump(self):
        dest = self.stack.pop()
        cond = self.stack.pop()

        if isinstance(dest, value.AVMCodePoint):
            pass
        elif isinstance(dest, AVMLabeledCodePoint):
            dest = dest.pc
        else:
            print(f"Conditional jumping to {dest.name}")
            raise Exception("Cjump insn requires codepoint but recieved " + str(dest))

        if cond != 0:
            self.pc = dest

    def swap1(self):
        instructions.swap1(self.stack)

    def swap2(self):
        instructions.swap2(self.stack)

    def dup0(self):
        instructions.dup0(self.stack)

    def dup1(self):
        instructions.dup1(self.stack)

    def dup2(self):
        instructions.dup2(self.stack)

    def tlen(self):
        instructions.tlen(self.stack)

    def tset(self):
        instructions.tset(self.stack)

    def tget(self):
        instructions.tget(self.stack)

    def istuple(self):
        item = self.stack.pop()
        self.stack.push(int(isinstance(item, value.Tuple)))

    def spush(self):
        self.stack.push(self.static)

    def rpush(self):
        self.stack.push(self.register)

    def rset(self):
        self.register = self.stack.pop()

    def add(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push((op1 + op2) & TT256M1)

    def sub(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push((op1 - op2) & TT256M1)

    def mul(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push((op1 * op2) & TT256M1)

    def div(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        if op2 != 0:
            try:
                self.push(op1 // op2)
            except:
                print(op1)
                print(op2)
                raise
        else:
            raise Exception("Can't divide by zero")

    def sdiv(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        s0, s1 = to_signed(op1), to_signed(op2)
        if s1 != 0:
            self.stack.push((abs(s0) // abs(s1) * (-1 if s0 * s1 < 0 else 1)) & TT256M1)
        else:
            raise Exception("Can't divide by zero")

    def mod(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        if op2 != 0:
            self.stack.push(op1 % op2)
        else:
            raise Exception("Can't mod by zero")

    def smod(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        s0, s1 = to_signed(op1), to_signed(op2)
        if s1 == 0:
            self.stack.push((abs(s0) % abs(s1) * (-1 if s0 < 0 else 1)) & TT256M1)
        else:
            raise Exception("Can't mod by zero")

    def exp(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(pow(op1, op2, TT256))

    def addmod(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        op3 = self.stack.pop()
        self.stack.push((op1 + op2) % op3 if op3 else 0)

    def mulmod(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        op3 = self.stack.pop()
        self.stack.push((op1 * op2) % op3 if op3 else 0)

    def signextend(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        if op1 <= 31:
            testbit = op1 * 8 + 7
            if op2 & (1 << testbit):
                self.stack.push(op2 | (TT256 - (1 << testbit)))
            else:
                self.stack.push(op2 & ((1 << testbit) - 1))
        else:
            self.stack.push(op2)

    def stackempty(self):
        self.stack.push(int(len(self.stack) == 0))

    def auxstackempty(self):
        self.stack.push(int(len(self.aux_stack) == 0))

    def nop(self):
        pass

    def pop(self):
        self.stack.pop()

    def pcpush(self):
        self.stack.push(self.pc)

    def halt(self):
        raise Exception("Machine halted improperly")
        self.halted = True

    def eq(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(int(op1 == op2))

    def iszero(self):
        op = self.stack.pop()
        self.stack.push(int(op == 0))

    def lt(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(int(op1 < op2))

    def gt(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(int(op1 > op2))

    def slt(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        s0, s1 = to_signed(op1), to_signed(op2)
        self.stack.push(1 if s0 < s1 else 0)

    def sgt(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        s0, s1 = to_signed(op1), to_signed(op2)
        self.stack.push(1 if s0 > s1 else 0)

    def hash(self):
        op = self.stack.pop()
        self.stack.push(big_endian_to_int(value.value_hash(op)))

    def gettime(self):
        self.stack.push(self.env.time_bounds)

    def bitwise_and(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(op1 & op2)

    def bitwise_or(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(op1 | op2)

    def bitwise_xor(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        self.stack.push(op1 ^ op2)

    def bitwise_not(self):
        op1 = self.stack.pop()
        self.stack.push(TT256M1 - op1)

    def byte(self):
        op1 = self.stack.pop()
        op2 = self.stack.pop()
        if op1 >= 32:
            self.stack.push(0)
        else:
            self.stack.push((op2 // 256 ** (31 - op1)) % 256)

    def incatomic(self):
        self.atomic_count += 1

    def decatomic(self):
        self.atomic_count -= 1

    def cast(self, typ):
        pass
