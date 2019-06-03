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

from .value import AVMCodePoint
from . import value
from .basic_vm import BasicVM
from .instructions import OP_CODES


class AVMOp:
    def __init__(self, name):
        self.name = name

    def __repr__(self):
        return f"AVMOp({self.name})"


class VM(BasicVM):
    def __init__(self, code=None, output_handler=None):
        super(VM, self).__init__()
        self.code = code
        self.output_handler = output_handler

        self.ops = {}
        for (op_name, op_code, pop_count, push_count) in OP_CODES:
            self.ops[op_code] = getattr(self, op_name)
        if code:
            self.pc = code[0]
        else:
            self.pc = AVMCodePoint(0, 0, b'')

    def debug_print(self):
        print(
            "debug_print:",
            "\nstack:",
            self.stack,
            "\nmessage:",
            self.register[5],
            "\npc:",
            self.pc.pc
        )

    def ifelse(self, true_block, false_block=None):
        val = self.stack.pop()
        if val:
            true_block(self)
        else:
            if false_block:
                false_block(self)

    def while_loop(self, cond_block, body_block):
        while True:
            self.push(999999)
            self.auxpush()
            cond_block(self)
            val = self.stack.pop()
            if not val:
                self.auxpop()
                self.pop()
                return
            body_block(self)
            self.auxpop()
            self.pop()

    def call(self, func):
        assert func.can_call
        func(self)

    def tnewn(self, size):
        self.push(value.Tuple([value.Tuple([]) for i in range(size)]))

    def tgetn(self, val):
        self.push(val)
        self.tget()

    def tsetn(self, val):
        self.push(val)
        self.tset()
