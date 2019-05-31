from .. import std
from .. import value
from . import types
from ..vm import VM
from ..annotation import modifies_stack

call_frame = std.struct.Struct("call_frame", [
    ("contractID", value.IntType()),       # transient
    ("memory", std.sized_byterange.sized_byterange.typ),           # transient
    ('contract_state', types.contract_state.typ),   # record
    ("contracts", types.contract_store.typ),        # record
    ("local_exec_state", types.local_exec_state.typ),  # transient
    ("return_data", std.sized_byterange.sized_byterange.typ),  # transient
    ("sent_queue", std.queue_tup.typ),       # record
    ("logs", std.stack_tup.typ),             # record
    "parent_frame",
    ("return_location", value.CodePointType()),
    ("saved_stack", std.stack.typ),
    ("saved_aux_stack", std.stack.typ),
])

typ = call_frame.typ

call_frame.update_type("parent_frame", call_frame.typ)

def make_empty():
    vm = VM()
    std.queue.new(vm)
    std.sized_byterange.new(vm)
    std.stack.new(vm)
    std.sized_byterange.new(vm)
    call_frame.new(vm)
    call_frame.set_val("memory")(vm)
    call_frame.set_val("logs")(vm)
    call_frame.set_val("return_data")(vm)
    call_frame.set_val("sent_queue")(vm)
    return vm.stack.items[0]

@modifies_stack([call_frame.typ], [types.contract_state.typ])
def lookup_current_state(vm):
    vm.dup0()
    call_frame.get("contractID")(vm)
    vm.swap1()
    call_frame.get("contracts")(vm)
    types.contract_store.get(vm)

@modifies_stack([call_frame.typ], [call_frame.typ])
def save_state(vm):
    # frame
    vm.dup0()
    call_frame.get("contract_state")(vm)
    vm.dup1()
    call_frame.get("contractID")(vm)
    vm.dup2()
    call_frame.get("contracts")(vm)
    types.contract_store.set_val(vm)
    vm.swap1()
    call_frame.set_val("contracts")(vm)

@modifies_stack([call_frame.typ], [call_frame.typ])
def setup_state(vm):
    # frame
    vm.dup0()
    lookup_current_state(vm)
    vm.swap1()
    call_frame.set_val("contract_state")(vm)

@modifies_stack([call_frame.typ, call_frame.typ], [value.CodePointType(), call_frame.typ])
def merge(vm):
    # parent_frame current_frame
    vm.swap1()
    save_state(vm)
    vm.swap1()
    # parent_frame current_frame
    vm.dup1()
    call_frame.get("contracts")(vm)
    vm.swap1()
    call_frame.set_val("contracts")(vm)
    # parent_frame current_frame
    vm.dup1()
    call_frame.get("sent_queue")(vm)
    vm.swap1()
    call_frame.set_val("sent_queue")(vm)
    # parent_frame current_frame
    vm.dup1()
    call_frame.get("logs")(vm)
    vm.swap1()
    call_frame.set_val("logs")(vm)
    # parent_frame current_frame
    vm.swap1()
    call_frame.get("return_location")(vm)
    # return_location parent_frame

# update:
#   contractID
#   message
#   memory
#   storage
#   wallet
# maintain:
#   contracts
#   sent_queue
#   logs
# unhandled (BUG):
#   return_data
@modifies_stack(
    [
        call_frame.typ,
        value.IntType(),
        types.local_exec_state.typ,
        value.CodePointType()
    ], [call_frame.typ]
)
def spawn(vm):
    # frame contractID message return_location
    vm.dup0()
    vm.swap2()
    # contractID frame frame message return_location
    vm.swap1()
    call_frame.set_val("contractID")(vm)
    setup_state(vm)
    # updated_frame parent_frame message return_location
    call_frame.set_val("parent_frame")(vm)
    # updated_frame message return_location
    call_frame.set_val("local_exec_state")(vm)
    # frame return_location
    call_frame.set_val("return_location")(vm)
    # frame
    std.sized_byterange.new(vm)
    vm.swap1()
    call_frame.set_val("memory")(vm)


@modifies_stack([types.contract_store.typ], [call_frame.typ])
def new_fresh(vm):
    # chain_state
    vm.push(make_empty())
    vm.cast(call_frame.typ)
    call_frame.set_val("contracts")(vm)
