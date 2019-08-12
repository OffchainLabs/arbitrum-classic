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
import sys

import arbitrum as arb
from arbitrum.evm.contract import create_evm_vm
from arbitrum.evm.contract_abi import ContractABI, create_output_handler


def run_until_halt(vm):
    i = 0
    while True:
        try:
            run = arb.run_vm_once(vm)
            if not run:
                print("Hit blocked insn")
                break
            i += 1
        except Exception as err:
            print("Error at", vm.pc.pc - 1, vm.code[vm.pc.pc - 1])
            print("Context", vm.code[vm.pc.pc - 6 : vm.pc.pc + 4])
            raise err
        if vm.halted:
            break
    logs = vm.logs
    vm.logs = []
    print("Ran VM for {} steps".format(i))
    return logs


def run_n_steps(vm, steps):
    log = []
    i = 0
    while i < steps:
        log.append((vm.pc.pc, vm.stack[:]))
        try:
            # print(vm.pc, vm.stack[:])
            arb.run_vm_once(vm)
            i += 1
        except Exception as err:
            print("Error at", vm.pc.pc - 1, vm.code[vm.pc.pc - 1])
            print("Context", vm.code[vm.pc.pc - 6 : vm.pc.pc + 4])
            raise err
        if vm.halted:
            break
    print("Ran VM for {} steps".format(i))
    return log


def make_msg_val(calldata):
    return arb.value.Tuple([calldata, 0, 0, 0])


if __name__ == "__main__":
    if len(sys.argv) != 2:
        raise Exception("Call as truffle_runner.py [compiled.json]")

    with open(sys.argv[1]) as json_file:
        raw_contracts = json.load(json_file)

    contracts = [ContractABI(contract) for contract in raw_contracts]
    vm = create_evm_vm(contracts)
    output_handler = create_output_handler(contracts)
    with open("code.txt", "w") as f:
        for instr in vm.code:
            f.write("{} {}".format(instr, instr.path))
            f.write("\n")

    elections = contracts[0]
    # vm.env.send_message([elections.candidatesCount(), 2345, 0, 0, 0])
    vm.env.send_message([make_msg_val(elections.candidates(14, 1)), 2345, 0, 0])
    vm.env.send_message([make_msg_val(elections.vote(16, 1)), 2345, 0, 0])
    vm.env.send_message([make_msg_val(elections.candidates(18, 1)), 2345, 0, 0])
    # vm.env.send_message([elections.candidates(2), 2345, 0, 0, 0])
    vm.env.deliver_pending()
    logs = run_until_halt(vm)
    for log in logs:
        print(output_handler(log))
