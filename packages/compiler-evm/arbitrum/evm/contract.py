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

import web3
import eth_utils
import eth_abi

from .compile import generate_evm_code
from .. import value, compile_program
from ..std import sized_byterange, stack


def generate_func(func_name, interface, address):
    def impl(self, seq, *args):
        func = getattr(interface.functions, func_name)
        msg_data = sized_byterange.frombytes(
            bytes.fromhex(func(*args)._encode_transaction_data()[2:])
        )
        return value.Tuple([msg_data, eth_utils.to_int(hexstr=address), seq])

    return impl

def generate_func2(func_name, interface):
    def impl(self, *args):
        func = getattr(interface.functions, func_name)
        return func(*args)._encode_transaction_data()

    return impl


class ArbContract:
    def __init__(self, contractInfo):
        self.w3 = web3.Web3()
        self.interface = self.w3.eth.contract(
            address=contractInfo["address"],
            abi=contractInfo["abi"]
        )
        self.address_string = contractInfo["address"]
        self.address = eth_utils.to_int(hexstr=self.address_string)
        self.abi = contractInfo["abi"]
        self.name = contractInfo["name"]
        self.funcs = {}
        self.functions = []
        self.code = bytes.fromhex(contractInfo["code"][2:])
        self.storage = {}
        if "storage" in contractInfo:
            raw_storage = contractInfo["storage"]
            for item in raw_storage:
                key = eth_utils.to_int(hexstr=item)
                self.storage[key] = eth_utils.to_int(hexstr=raw_storage[item])
        self.address = eth_utils.to_int(hexstr=contractInfo["address"])
        for func_interface in self.interface.abi:
            if func_interface["type"] == "function":
                id_bytes = eth_utils.function_abi_to_4byte_selector(
                    func_interface
                )
                func_id = eth_utils.big_endian_to_int(id_bytes)
                self.funcs[func_id] = func_interface
            elif func_interface["type"] == "event":
                id_bytes = eth_utils.event_abi_to_log_topic(func_interface)
                func_id = eth_utils.big_endian_to_int(id_bytes)
                self.funcs[func_id] = func_interface
        funcs = [
            x for x in dir(self.interface.functions)
            if x[0] != '_' and x != "abi"
        ]
        for func in funcs:
            setattr(
                ArbContract,
                func,
                generate_func(
                    func,
                    self.interface,
                    contractInfo["address"]
                )
            )
            setattr(
                ArbContract,
                "_" + func,
                generate_func2(
                    func,
                    self.interface
                )
            )
            self.functions.append(func)

    def __repr__(self):
        return "ArbContract({})".format(self.name)


def get_return_abi(func_info):
    output_types = [param['type'] for param in func_info["outputs"]]
    return '(' + ','.join(output_types) + ')'


def convert_log_raw(logVal):
    topics = []
    for topic in logVal[3:]:
        raw_bytes = eth_utils.int_to_big_endian(topic)
        raw_bytes = (32 - len(raw_bytes)) * b'\x00' + raw_bytes
        topics.append(raw_bytes)

    output_byte_str = sized_byterange.tohex(logVal[1])
    output_bytes = eth_utils.to_bytes(hexstr=output_byte_str)

    return {
        "contract": logVal[0],
        "id": logVal[2],
        "data": output_bytes,
        "topics": topics
    }


def decode_log(logVal, abis):
    event_interface = abis[logVal["contract"]].funcs[logVal["id"]]
    ret = {}
    topics = [inp for inp in event_interface["inputs"] if inp["indexed"]]
    for (topic, topic_data) in zip(topics, logVal["topics"]):
        ret[topic["name"]] = eth_abi.decode_single(topic["type"], topic_data)

    other_inputs = [
        inp for inp in event_interface["inputs"]
        if not inp["indexed"]
    ]
    arg_type = '(' + ','.join([inp["type"] for inp in other_inputs]) + ')'
    decoded = eth_abi.decode_single(arg_type, logVal["data"])

    for (inp, val) in zip(other_inputs, decoded):
        ret[inp["name"]] = val

    return {
        "name": event_interface['name'],
        "args": ret
    }


REVERT_CODE = 0
INVALID_CODE = 1
RETURN_CODE = 2
STOP_CODE = 3
INVALID_SEQUENCE_CODE = 4


class EVMCall:
    def __init__(self, abis, func_interface, val):
        output_byte_str = sized_byterange.tohex(val[2])
        output_bytes = eth_utils.to_bytes(hexstr=output_byte_str)
        self.interface = func_interface
        self.output_values = eth_abi.decode_single(
            get_return_abi(func_interface),
            output_bytes
        )
        self.logs = [
            decode_log(convert_log_raw(logVal), abis)
            for logVal in stack.to_list(val[1])
        ]

    def __str__(self):
        ret = "{} returned {}".format(self.interface["name"], self.output_values)
        for log in self.logs:
            ret += "\n{} logged event {}{}".format(
                self.interface['name'],
                log['name'],
                log['args']
            )
        return ret


class EVMRevert:
    def __init__(self, func_interface, val):
        self.interface = func_interface
        self.output_byte_str = sized_byterange.tohex(val[2])

    def __str__(self):
        return "{} failed with revert returning {}".format(
            self.interface['name'],
            self.output_byte_str
        )


class EVMInvalid:
    def __init__(self, func_interface):
        self.interface = func_interface

    def __str__(self):
        return "{} failed with invalid op".format(self.interface['name'])


class EVMInvalidSequence:
    def __init__(self, func_interface):
        self.interface = func_interface

    def __str__(self):
        return "{} failed with invalid sequence".format(self.interface['name'])


class EVMUnknownResponseError:
    def __init__(self, func_interface, val):
        self.interface = func_interface
        self.val = val

    def __str__(self):
        return "{} had unknown error: {}".format(self.func_interface['name'], self.val)


class EVMStop:
    def __init__(self, abis, func_interface, val):
        self.interface = func_interface
        self.logs = [
            decode_log(convert_log_raw(logVal), abis)
            for logVal in stack.to_list(val[1])
        ]

    def __str__(self):
        ret = "{} completed successfully".format(self.interface['name'])
        for log in self.logs:
            ret += "\n{} logged event {}{}".format(
                self.interface['name'],
                log['name'],
                log['args']
            )
        return ret


def parse_output(val, abis):
    return_code = val[3]

    contract_num = val[0][0][0][1]
    input_hex = sized_byterange.tohex(val[0][0][0][0])
    func_id = int(input_hex[2:10], 16)

    try:
        func_interface = abis[contract_num].funcs[func_id]
    except KeyError:
        print("Unknown function returned {}".format(return_code))
        return True

    if return_code == RETURN_CODE:
        return EVMCall(abis, func_interface, val)
    elif return_code == REVERT_CODE:
        return EVMRevert(func_interface, val)
    elif return_code == INVALID_CODE:
        return EVMInvalid(func_interface)
    elif return_code == INVALID_SEQUENCE_CODE:
        return EVMInvalidSequence(func_interface)
    elif return_code == STOP_CODE:
        return EVMStop(abis, func_interface, val)
    else:
        return EVMUnknownResponseError(func_interface, val)

# [logs, contract_num, func_code, return_val, return_code]
def create_output_handler(contracts):
    abis = {}
    for contract in contracts:
        abis[contract.address] = contract

    def output_handler(val):
        return parse_output(val, abis)

    return output_handler


def create_evm_vm(contracts, should_optimize=True):
    code = {}
    storage = {}
    for contract in contracts:
        code[contract.address] = contract.code
        storage[contract.address] = contract.storage

    initial_block, code = generate_evm_code(code, storage)
    vm = compile_program(initial_block, code, should_optimize)
    vm.output_handler = create_output_handler(contracts)

    return vm
