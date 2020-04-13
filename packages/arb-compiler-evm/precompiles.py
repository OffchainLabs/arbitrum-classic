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

import arbitrum as arb


def infiniteCounterLoop(vm):
    vm.push(0)
    vm.while_loop(lambda v: [v.push(1)], lambda v: [v.push(1), v.ecdsa()])


def makeAoFile(func, filepath):
    code = arb.compile_block(func)
    vm = arb.compile_program(arb.ast.BlockStatement([]), code)
    vm.static = 4
    with open(filepath, "wb") as f:
        arb.marshall.marshall_vm(vm, f)


makeAoFile(infiniteCounterLoop, "precompiles.ao")
