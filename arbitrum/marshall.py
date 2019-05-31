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
        file.write(op_type.to_bytes(1, byteorder='little', signed=False))
        file.write(val.op_code.to_bytes(1, byteorder='little', signed=False))
    elif isinstance(val, int):
        op_type = 0
        file.write(op_type.to_bytes(1, byteorder='little', signed=False))
        file.write(val.to_bytes(1, byteorder='little', signed=False))
    elif isinstance(val, ImmediateOp):
        op_type = 1
        file.write(op_type.to_bytes(1, byteorder='little', signed=False))
        file.write(val.get_op().to_bytes(1, byteorder='little', signed=False))
        marshall_value(val.val, file)
    else:
        raise Exception(f"Tried to marshall bad operation type {val}")


def marshall_codepoint(val, file):
    file.write(val.pc.to_bytes(8, byteorder='little', signed=True))
    marshall_op(val.op, file)
    val.next_hash = b'0' * (32 - len(val.next_hash)) + val.next_hash
    file.write(val.next_hash)


def marshall_tuple(val, file):
    for item in val:
        marshall_value(item, file)


def marshall_value(val, file):
    if isinstance(val, value.Tuple):
        file.write((TUPLE_TYPE_CODE + len(val)).to_bytes(
            1,
            byteorder='little',
            signed=False
        ))
        marshall_tuple(val, file)
    elif isinstance(val, int):
        file.write(INT_TYPE_CODE.to_bytes(
            1,
            byteorder='little',
            signed=False
        ))
        marshall_int(val, file)
    elif isinstance(val, value.AVMCodePoint):
        file.write(CODE_POINT_TYPE_CODE.to_bytes(
            1,
            byteorder='little',
            signed=False
        ))
        marshall_codepoint(val, file)
    elif isinstance(val, AVMLabeledCodePoint):
        file.write(CODE_POINT_TYPE_CODE.to_bytes(
            1,
            byteorder='little',
            signed=False
        ))
        marshall_codepoint(val.pc, file)
    else:
        raise Exception(f"Can't marshall unexcepted value {val}")


def marshall_vm(vm, file):
    file.write(len(vm.code).to_bytes(8, byteorder='little', signed=True))
    for instr in vm.code:
        marshall_op(instr.op, file)
    marshall_value(vm.static, file)
