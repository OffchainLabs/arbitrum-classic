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

from . import contract_templates
from ..std import stack, bytestack_tohex

REVERT_CODE = 0
INVALID_CODE = 1
RETURN_CODE = 2
STOP_CODE = 3
INVALID_SEQUENCE_CODE = 4
INSUFFICIENT_BALANCE = 5
ETH_DEPOSIT = 6


class EVMLog:
    def __init__(self, val):
        self.decoded = False
        self.name = ""
        self.args = []
        self.contract_id = val[0]
        print("val[1]", val[1])
        self.data = eth_utils.to_bytes(hexstr=bytestack_tohex(val[1]))
        self.event_id = val[2]
        self.topics = []
        for topic in val[3:]:
            raw_bytes = eth_utils.int_to_big_endian(topic)
            raw_bytes = (32 - len(raw_bytes)) * b"\x00" + raw_bytes
            self.topics.append(raw_bytes)

    def decode(self, events):
        abi = events[self.contract_id][self.event_id]
        ret = {}
        topics = [inp for inp in abi["inputs"] if inp["indexed"]]
        for (topic, topic_data) in zip(topics, self.topics):
            ret[topic["name"]] = eth_abi.decode_single(topic["type"], topic_data)
        other_inputs = [inp for inp in abi["inputs"] if not inp["indexed"]]
        arg_type = "(" + ",".join([inp["type"] for inp in other_inputs]) + ")"
        decoded = eth_abi.decode_single(arg_type, self.data)

        for (inp, val) in zip(other_inputs, decoded):
            ret[inp["name"]] = val

        self.decoded = True
        self.name = abi["name"]
        self.args = ret

    def __repr__(self):
        if self.decoded:
            return "Log({}, {})".format(self.name, self.args)
        else:
            return "Log({})".format(self.topics)

    def __str__(self):
        if self.decoded:
            return "{}({})".format(self.name, self.args)
        else:
            return "Log({})".format(self.topics)


class LogMessage:
    def __init__(self, value):
        self.block_number = value[0]
        self.tx_hash = value[1]
        wrapped_data = value[2]
        self.message_type = wrapped_data[0]
        self.caller = wrapped_data[1]
        if self.message_type == 0:
            tx_message = wrapped_data[2]
            self.contract_id = tx_message[0]
            self.sequence_num = tx_message[1]
            self.value = tx_message[2]
            self.data = bytestack_tohex(tx_message[3])
        elif self.message_type in [2, 3]:
            token_message = wrapped_data[2]
            self.token_address = token_message[0]
            self.dest = token_message[1]
            self.amount = token_message[2]

    def func_id(self):
        return self.data[2:10]

    def get_abi(self, functions):
        funcs = functions.get(self.contract_id, None)
        if funcs:
            return funcs.get(self.func_id(), None)
        else:
            return None

    def raw_func_name(self):
        if self.message_type == 0:
            return self.func_id()
        elif self.message_type == 1:
            return "ETH_DEPOSIT"
        elif self.message_type == 2:
            return "ERC20_DEPOSIT"
        elif self.message_type == 3:
            return "ERC721_DEPOSIT"
        elif self.message_type == 4:
            return "ETH_WITHDRAWAL"
        elif self.message_type == 5:
            return "ERC20_WITHDRAWAL"
        elif self.message_type == 6:
            return "ERC721_WITHDRAWAL"
        else:
            raise Exception("Unknown function type")


class EVMOutput:
    def __init__(self, val):
        self.orig_message = LogMessage(val[0])
        self.decoded = False
        self.abi = {}
        self.name = self.orig_message.raw_func_name()

    def decode(self, functions, events):
        if self.orig_message.message_type == 0:
            self.abi = self.orig_message.get_abi(functions)
            if self.abi:
                self.name = self.abi["name"]

        if self.orig_message.message_type == 2:
            events[self.orig_message.token_address] = events[
                contract_templates.ERC20_ADDRESS
            ]
            functions[self.orig_message.token_address] = functions[
                contract_templates.ERC20_ADDRESS
            ]
        elif self.orig_message.message_type == 3:
            events[self.orig_message.token_address] = events[
                contract_templates.ERC721_ADDRESS
            ]
            functions[self.orig_message.token_address] = functions[
                contract_templates.ERC721_ADDRESS
            ]

        self.decoded = True


class EVMCall(EVMOutput):
    def __init__(self, val):
        super().__init__(val)
        self.output_bytes = eth_utils.to_bytes(hexstr=bytestack_tohex(val[2]))
        self.output_values = []
        self.logs = [EVMLog(logVal) for logVal in stack.to_list(val[1])]

    def __repr__(self):
        if self.decoded:
            return "EVMCall({}, {}, {})".format(
                self.name, self.output_values, self.logs
            )
        return "EVMCall({}, {}, {})".format(
            self.orig_message.raw_func_name(), self.output_bytes, self.logs
        )

    def __str__(self):
        if self.decoded:
            ret = "{} returned {}".format(self.name, self.output_values)
            for log in self.logs:
                ret += "\n{} logged event {}".format(self.name, log)
            return ret
        else:
            return repr(self)

    def decode(self, functions, events):
        super().decode(functions, events)
        self.output_values = eth_abi.decode_abi(
            [out["type"] for out in self.abi["outputs"]], self.output_bytes
        )
        for log in self.logs:
            log.decode(events)


class EVMStop(EVMOutput):
    def __init__(self, val):
        super().__init__(val)
        self.logs = [EVMLog(logVal) for logVal in stack.to_list(val[1])]

    def __repr__(self):
        if self.decoded:
            return "EVMStop({}, {})".format(self.name, self.logs)
        return "EVMStop({}, {})".format(self.orig_message.raw_func_name(), self.logs)

    def __str__(self):
        if self.decoded:
            ret = "{} completed successfully".format(self.name)
            for log in self.logs:
                ret += "\n{} logged event {}".format(self.name, log)
            return ret
        else:
            return repr(self)

    def decode(self, functions, events):
        super().decode(functions, events)
        for log in self.logs:
            log.decode(events)


class EVMRevert(EVMOutput):
    def __init__(self, val):
        super().__init__(val)
        self.output_bytes = eth_utils.to_bytes(hexstr=bytestack_tohex(val[2]))

    def __repr__(self):
        return "EVMRevert({}, {})".format(self.name, self.output_bytes)

    def decode(self, functions, events):
        super().decode(functions, events)


class EVMInvalid(EVMOutput):
    def __init__(self, val):
        super().__init__(val)

    def __repr__(self):
        return "EVMInvalid()"


class EVMInvalidSequence(EVMOutput):
    def __init__(self, val):
        super().__init__(val)

    def __repr__(self):
        return "EVMInvalidSequence()"


class EVMInsufficientBalance(EVMOutput):
    def __init__(self, val):
        super().__init__(val)

    def __repr__(self):
        return "EVMInsufficientBalance()"


class EVMDeposit(EVMOutput):
    def __init__(self, val):
        super().__init__(val)

    def __repr__(self):
        return "EVMDeposit()"


class EVMUnknownResponseError(EVMOutput):
    def __init__(self, val):
        super().__init__(val)
        self.val = val

    def __repr__(self):
        return "EVMUnknownResponseError({})".format(self.val[3])


EVM_OUTPUT_TYPES = {
    RETURN_CODE: EVMCall,
    REVERT_CODE: EVMRevert,
    INVALID_CODE: EVMInvalid,
    INVALID_SEQUENCE_CODE: EVMInvalidSequence,
    INSUFFICIENT_BALANCE: EVMInsufficientBalance,
    STOP_CODE: EVMStop,
    ETH_DEPOSIT: EVMDeposit,
}


def parse(val):
    return EVM_OUTPUT_TYPES.get(val[3], EVMUnknownResponseError)(val)
