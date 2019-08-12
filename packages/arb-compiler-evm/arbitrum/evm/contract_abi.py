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
import eth_utils

from .contract import Contract
from . import log
from ..std import sized_byterange
from .. import value


def generate_func(func_id, func_abi, address):
    def impl(self, seq, *args):
        if len(args) != len(func_abi["inputs"]):
            raise Exception(
                "Function with abi {} passed not matching {} args".format(
                    func_abi, list(args)
                )
            )
        encoded_input = func_id + eth_abi.encode_abi(
            [inp["type"] for inp in func_abi["inputs"]], list(args)
        )
        msg_data = sized_byterange.frombytes(encoded_input)
        return value.Tuple([msg_data, address, seq])

    return impl


def generate_func2(func_id, func_abi):
    def impl(self, seq, *args):
        encoded_input = eth_abi.encode_abi(
            [inp["type"] for inp in func_abi["inputs"]], list(args)
        )
        return (func_id + encoded_input).hex()

    return impl


class ContractABI(Contract):
    def __init__(self, contractInfo):
        super().__init__(contractInfo)
        self.name = contractInfo["name"]

        abi = contractInfo["abi"]
        func_abis = [func_abi for func_abi in abi if func_abi["type"] == "function"]
        event_abis = [event_abi for event_abi in abi if event_abi["type"] == "event"]

        self.funcs = {}
        for func_abi in func_abis:
            id_bytes = eth_utils.function_abi_to_4byte_selector(func_abi)
            self.funcs[id_bytes] = func_abi

        for func_id, func_abi in self.funcs.items():
            setattr(
                ContractABI,
                func_abi["name"],
                generate_func(func_id, func_abi, self.address),
            )
            setattr(
                ContractABI, "_" + func_abi["name"], generate_func2(func_id, func_abi)
            )

        self.events = {}
        for event_abi in event_abis:
            id_bytes = eth_utils.event_abi_to_log_topic(event_abi)
            event_id = eth_utils.big_endian_to_int(id_bytes)
            self.events[event_id] = event_abi

    def __repr__(self):
        return "ContractABI({})".format(self.name)


def create_output_handler(contracts):
    events = {}
    functions = {}
    for contract in contracts:
        for event_id, abi in contract.events.items():
            events[(contract.address, event_id)] = abi
        for func_id, abi in contract.funcs.items():
            functions[(contract.address, func_id.hex())] = abi

    def output_handler(output):
        output = log.parse(output)
        output.decode(functions, events)
        return output

    return output_handler
