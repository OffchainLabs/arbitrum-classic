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

import json
import arbitrum as arb
import eth_utils
from collections import Counter
from arbitrum.instructions import OPS
import sys
from arbitrum.evm.contract import ArbContract, create_evm_vm


def run_until_halt(vm):
    log = []
    i = 0
    push_counts = Counter()
    while True:
        # print(vm.pc, repr(vm.stack[:1])[:50])
        try:
            if vm.pc.op.get_op() == OPS["spush"]:
                push_counts[vm.pc.path[-1][5:-1]] += 1
            run = arb.run_vm_once(vm)
            if not run:
                print("Hit blocked insn")
                break
            i += 1
        except Exception as err:
            print("Error at", vm.pc.pc - 1, vm.code[vm.pc.pc - 1])
            print("Context", vm.code[vm.pc.pc - 6: vm.pc.pc + 4])
            raise err
        if vm.halted:
            break
    for log in vm.logs:
        vm.output_handler(log)
    vm.logs = []
    print("Ran VM for {} steps".format(i))
    print("VM sent messages {}".format(vm.sent_messages))
    # print(push_counts)
    return log


def make_msg_val(calldata):
    return arb.value.Tuple([calldata, 0, 0, 0])


if __name__ == '__main__':
    if len(sys.argv) != 2:
        raise Exception("Call as truffle_runner.py [compiled.json]")

    with open(sys.argv[1]) as json_file:
        raw_contracts = json.load(json_file)

    contracts = [ArbContract(contract) for contract in raw_contracts]
    vm = create_evm_vm(contracts, False)
    with open("code.txt", "w") as f:
        for instr in vm.code:
            f.write("{} {}".format(instr, instr.path))
            f.write("\n")
    contract = contracts[0]

    person_a = '0x1111111122222222000000000000000000000000'
    person_b = '0x2222222222222222222222222222222222222222'
    person_a_int = eth_utils.to_int(hexstr=person_a)
    person_b_int = eth_utils.to_int(hexstr=person_b)
    print("person_a_int", person_a_int)
    print("person_b_int", person_b_int)
    erc20_address = '0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359'
    erc721_address = '0x06012c8cf97BEaD5deAe237070F9587f8E7A266d'

    vm.env.send_message([
        make_msg_val(contract.deposit(10)),
        person_a_int,
        10000000,
        eth_utils.to_int(hexstr=erc20_address + '00')
    ])

    vm.env.send_message([
        make_msg_val(contract.sendERC20(12, erc20_address, 5432)),
        person_a_int,
        0,
        0
    ])

    vm.env.send_message([
        make_msg_val(contract.deposit(10)),
        person_a_int,
        10000000,
        eth_utils.to_int(hexstr=erc721_address + '01')
    ])

    vm.env.send_message([
        make_msg_val(contract.sendERC721(12, erc721_address, 10000000)),
        person_a_int,
        0,
        0
    ])

    vm.env.deliver_pending()
    run_until_halt(vm)
