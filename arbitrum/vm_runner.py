from .annotation import modifies_stack
from .ast import ImmediateOp, AVMLabeledCodePoint
from . import value
from .basic_vm import VMBlocked, VMBlockedAdvance
import traceback


# [val]
@modifies_stack(1, 1)
def is_zero(vm):
    vm.push(0)
    vm.eq()


def run_vm_once(vm):
    if vm.halted:
        raise Exception("Can't run VM since it is halted")
    if vm.pc.pc == -2:
        raise Exception("VM hit unhandled error")
    instr = vm.pc.op
    old_pc = vm.pc

    try:
        if isinstance(instr, ImmediateOp):
            vm.push(instr.val)
            vm.ops[instr.op.op_code]()
        else:
            vm.ops[instr.op_code]()
    except VMBlocked:
        return False
    except VMBlockedAdvance:
        vm.pc = vm.code[vm.pc.pc + 1]
        return False
    except Exception as err:
        print("Hit exception", err, vm.err_handler, traceback.print_tb(err.__traceback__))
        if isinstance(vm.err_handler, value.CodePointType):
            vm.pc = vm.err_handler
        elif isinstance(vm.err_handler, AVMLabeledCodePoint):
            vm.pc = vm.err_handler.pc
        elif isinstance(vm.err_handler, value.AVMCodePoint):
            vm.pc = vm.err_handler
        else:
            print("Error handler", vm.err_handler)
            raise

    if vm.pc.pc == old_pc.pc:
        vm.pc = vm.code[vm.pc.pc + 1]

    return True
