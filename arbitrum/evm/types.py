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
