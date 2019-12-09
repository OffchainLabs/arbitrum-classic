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

from ..annotation import modifies_stack
from .. import value
from ..ast import AVMLabel

ERC20_ADDRESS_STRING = "0xfffffffffffffffffffffffffffffffffffffffe"
ER721_ADDRESS_STRING = "0xfffffffffffffffffffffffffffffffffffffffd"
ERC20_ADDRESS = eth_utils.to_int(hexstr=ERC20_ADDRESS_STRING)
ERC721_ADDRESS = eth_utils.to_int(hexstr=ER721_ADDRESS_STRING)


@modifies_stack([], [value.CodePointType()])
def erc20_codepoint(vm):
    vm.push(AVMLabel("contract_entry_" + str(ERC20_ADDRESS)))


@modifies_stack([], [value.CodePointType()])
def erc721_codepoint(vm):
    vm.push(AVMLabel("contract_entry_" + str(ERC721_ADDRESS)))
