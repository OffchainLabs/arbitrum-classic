from .. import std
from .. import value

contract_state = std.Struct("contract_state", [
    ('storage', std.keyvalue_int_int.typ),
    ('wallet', std.currency_store.typ)
])

message = std.Struct("message", [
    ("data", value.ValueType()),
    ("sender", value.IntType()),
    ("amount", value.IntType()),
    ("type", value.IntType()),
])

message_blockchain_data = std.Struct("message_blockchain_data", [
    ("data", value.ValueType()),
    ("timestamp", value.IntType()),
    ("block_number", value.IntType()),
    ("txhash", value.IntType())
])

message_data = std.Struct("message_data", [
    ("data", value.ValueType()),
    ("contract_id", value.IntType()),
    ("sequence_num", value.IntType())
])

contract_store = std.make_keyvalue_type(value.IntType(), contract_state.typ)

local_exec_state = std.Struct("local_exec_state", [
    ("data", value.ValueType()),
    ("sender", value.IntType()),
    ("amount", value.IntType()),
    ("type", value.IntType())
])
