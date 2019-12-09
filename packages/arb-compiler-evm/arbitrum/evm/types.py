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


message = std.Struct(
    "message",
    [
        ("type", value.IntType()),
        ("sender", value.IntType()),
        ("message", value.ValueType()),
    ],
)

token_transfer_message = std.Struct(
    "token_transfer_message",
    [
        ("token_address", value.IntType()),
        ("dest", value.IntType()),
        ("amount", value.IntType()),
    ],
)

eth_transfer_message = std.Struct(
    "eth_transfer_message", [("dest", value.IntType()), ("amount", value.IntType())]
)

tx_message = std.Struct(
    "tx_message",
    [
        ("to", value.IntType()),
        ("sequence_num", value.IntType()),
        ("value", value.IntType()),
        ("data", value.ValueType()),
    ],
)

tx_call_data = std.Struct(
    "tx_call_data",
    [
        ("dest", value.IntType()),
        ("value", value.IntType()),
        ("data", value.ValueType()),
    ],
)

local_exec_state = std.Struct(
    "local_exec_state",
    [
        ("data", value.ValueType()),
        ("caller", value.IntType()),
        ("value", value.IntType()),
    ],
)

ethbridge_message = std.Struct(
    "ethbridge_message",
    [
        ("timestamp", value.IntType()),
        ("block_number", value.IntType()),
        ("txhash", value.IntType()),
        ("message", message.typ),
    ],
)
