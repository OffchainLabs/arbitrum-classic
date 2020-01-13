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

import eth_utils
from importlib_resources import read_text
import json

from ..annotation import modifies_stack
from .. import value
from ..ast import AVMLabel

ERC20_ADDRESS_STRING = "0xfffffffffffffffffffffffffffffffffffffffe"
ERC721_ADDRESS_STRING = "0xfffffffffffffffffffffffffffffffffffffffd"
ARBSYS_ADDRESS_STRING = "0x0000000000000000000000000000000000000064"
ERC20_ADDRESS = eth_utils.to_int(hexstr=ERC20_ADDRESS_STRING)
ERC721_ADDRESS = eth_utils.to_int(hexstr=ERC721_ADDRESS_STRING)


def get_templates():
    raw_contract_templates_data = read_text("arbitrum.evm", "contract-templates.json")
    raw_contract_templates = json.loads(raw_contract_templates_data)
    token_templates = {}
    for raw_contract in raw_contract_templates:
        token_templates[raw_contract["name"]] = raw_contract
    return token_templates


def get_arbsys():
    arbsys_data = read_text("arbitrum.evm", "ArbSys.json")
    arbsys = json.loads(arbsys_data)
    arbsys["address"] = ARBSYS_ADDRESS_STRING
    arbsys["code"] = "0x"
    arbsys["name"] = "ArbSys"
    return arbsys


def get_erc20_contract():
    erc20 = get_templates()["ArbERC20"]
    erc20["address"] = ERC20_ADDRESS_STRING
    return erc20


def get_erc721_contract():
    erc721 = get_templates()["ArbERC721"]
    erc721["address"] = ERC721_ADDRESS_STRING
    return erc721


@modifies_stack([], [value.CodePointType()])
def erc20_codepoint(vm):
    vm.push(AVMLabel("contract_entry_" + str(ERC20_ADDRESS)))


@modifies_stack([], [value.CodePointType()])
def erc721_codepoint(vm):
    vm.push(AVMLabel("contract_entry_" + str(ERC721_ADDRESS)))
