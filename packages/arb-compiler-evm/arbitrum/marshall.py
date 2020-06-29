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

import eth_abi
from .ast import AVMLabeledCodePoint, BasicOp, ImmediateOp
from . import value

INT_TYPE_CODE = 0
CODE_POINT_TYPE_CODE = 1
TUPLE_TYPE_CODE = 3


def marshall_int(val, file):
    file.write(eth_abi.encode_single("uint256", val))


def marshall_op(val, file):
    if isinstance(val, BasicOp):
        op_type = 0
        file.write(op_type.to_bytes(1, byteorder="big", signed=False))
        file.write(val.op_code.to_bytes(1, byteorder="big", signed=False))
    elif isinstance(val, int):
        op_type = 0
        file.write(op_type.to_bytes(1, byteorder="big", signed=False))
        file.write(val.to_bytes(1, byteorder="big", signed=False))
    elif isinstance(val, ImmediateOp):
        op_type = 1
        file.write(op_type.to_bytes(1, byteorder="big", signed=False))
        file.write(val.get_op().to_bytes(1, byteorder="big", signed=False))
        marshall_value(val.val, file)
    else:
        raise Exception("Tried to marshall bad operation type {}".format(val))


def marshall_codepoint(val, file):
    marshall_op(val.op, file)
    next_hash = b"\0" * (32 - len(val.next_hash)) + val.next_hash
    file.write(next_hash)


def marshall_value(val, file):
    if isinstance(val, value.Tuple):
        file.write(
            (TUPLE_TYPE_CODE + len(val)).to_bytes(1, byteorder="big", signed=False)
        )
        for item in val:
            marshall_value(item, file)
    elif isinstance(val, int):
        file.write(INT_TYPE_CODE.to_bytes(1, byteorder="big", signed=False))
        marshall_int(val, file)
    elif isinstance(val, value.AVMCodePoint):
        file.write(CODE_POINT_TYPE_CODE.to_bytes(1, byteorder="big", signed=False))
        marshall_codepoint(val, file)
    elif isinstance(val, AVMLabeledCodePoint):
        file.write(CODE_POINT_TYPE_CODE.to_bytes(1, byteorder="big", signed=False))
        marshall_codepoint(val.pc, file)
    else:
        raise Exception("Can't marshall unexcepted value {}".format(val))


def marshall_value_json(val):
    if isinstance(val, value.Tuple):
        return {"Tuple": [marshall_value_json(item) for item in val]}
    if isinstance(val, int):
        return {"Int": str(val)}
    if isinstance(val, value.AVMCodePoint):
        return {"CodePoint": {"Internal": val.pc}}
    if isinstance(val, AVMLabeledCodePoint):
        return {"CodePoint": {"Internal": val.pc.pc}}
    raise Exception("Can't marshall unexcepted value {}".format(val))


def marshall_op_json(val):
    if isinstance(val, BasicOp):
        return {"opcode": val.op_code, "immediate": None}
    if isinstance(val, int):
        return {"opcode": val, "immediate": None}
    if isinstance(val, ImmediateOp):
        return {"opcode": val.get_op(), "immediate": marshall_value_json(val.val)}
    raise Exception("Tried to marshall bad operation type {}".format(val))


AO_VERSION = 2


def marshall_vm_json(vm):
    return {
        "version": AO_VERSION,
        "code": [marshall_op_json(instr.op) for instr in vm.code],
        "static_val": marshall_value_json(vm.static),
        "extensions": [],
    }
