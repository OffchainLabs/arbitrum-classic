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

count = 0


def binaryOp(vm, arg1, arg2, res, op):
    global count
    vm.push(arg1)
    vm.push(arg2)
    op()
    vm.push(res)
    vm.eq()
    vm.push(arb.ast.AVMLabel("next" + str(count)))
    vm.cjump()
    vm.error()
    vm.set_label(arb.ast.AVMLabel("next" + str(count)))
    count += 1


def test(vm):
    binaryOp(vm, 4, 3, 7, vm.add)
    #    binaryOp(vm,4,3,6,vm.add)
    binaryOp(vm, 0, 0, 0, vm.add)
    #    binaryOp(vm,neg1,4,vm.add)
    #    binaryOp(vm,-2,1,vm.add)
    binaryOp(vm, 4, 3, 12, vm.mul)
    binaryOp(vm, 3, 0, 0, vm.mul)
    vm.halt()


code = arb.compile_block(test)
vm = arb.compile_program(arb.ast.BlockStatement([]), code)
# print(vm.code)
with open("../arb-validator/test/opcodetest.ao", "wb") as f:
    arb.marshall.marshall_vm(vm, f)
