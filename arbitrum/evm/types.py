from .. import std
from .. import value

contract_state = std.Struct("contract_state", [
    ('storage', std.keyvalue_int_int.typ),
    ('wallet', std.currency_store.typ)
])

message = std.Struct("message", [
    "data",
    ("sender", value.IntType()),
    ("amount", value.IntType()),
    ("type", value.IntType()),
    ("timestamp", value.IntType())
])

contract_store = std.make_keyvalue_type(value.IntType(), contract_state.typ)
