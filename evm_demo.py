import sys
from collections import Counter
import hashlib

import web3
import solcx
import eth_utils

import arbitrum as arb
from arbitrum.instructions import OPS
from arbitrum.evm import constructor
from arbitrum.evm.contract import ArbContract, create_evm_vm


def run_until_halt(vm):
    log = []
    i = 0
    push_counts = Counter()
    while True:
        # if (
        #         vm.pc.op.get_op() == OPS["auxpush"] or
        #         vm.pc.op.get_op() == OPS["auxpop"]
        # ):
        #     print(vm.pc, vm.aux_stack[:])
        # print(vm.pc, vm.stack[:])
        # log.append((vm.pc, vm.stack[:]))
        # print(vm.pc, vm.stack.items[:5])
        try:
            if vm.pc.op.get_op() == OPS["spush"]:
                push_counts[vm.pc.path[-1][5:-1]] += 1
            keep_going = arb.run_vm_once(vm)
            if not keep_going:
                break
            i += 1
        except Exception as err:
            print("Error at", vm.pc.pc - 1, vm.code[vm.pc.pc - 1])
            print("Context", vm.code[vm.pc.pc - 6: vm.pc.pc + 4])
            raise err
        if vm.halted:
            break
    print(f"Ran VM for {i} steps")
    # print(push_counts)
    return log


def run_n_steps(vm, steps):
    log = []
    i = 0
    while i < steps:
        log.append((vm.pc.pc, vm.stack[:]))
        try:
            # print(vm.pc.pc, vm.stack)
            arb.run_vm_once(vm)
            i += 1
        except Exception as err:
            print("Error at", vm.pc.pc - 1, vm.code[vm.pc.pc - 1])
            print("Context", vm.code[vm.pc.pc - 6: vm.pc.pc + 4])
            raise err
        if vm.halted:
            break
    print(f"Ran VM for {i} steps")
    return log


def prepare_contracts(solidity_files):
    compiled = solcx.compile_files(
        solidity_files,
        optimize=True,
        allow_paths="/Users/hkalodner/Documents/OffchainLabs/Github/arbitrum-python/contracts"
    )
    contracts = {}
    for contractName in compiled:
        name = contractName.split(":")[-1]
        contract = compiled[contractName]

        address_bytes = hashlib.sha224(bytes.fromhex(
            contract['bin-runtime'][2:]
        )).hexdigest()[:40]
        address_string = web3.Web3.toChecksumAddress(address_bytes)

        contracts[name] = {
            'address': address_string,
            'abi': contract['abi'],
            'name': name,
            'code': contract['bin'],
        }
    return contracts


def test_fibonacci(total_count, query_num):
    contracts_code = prepare_contracts(["Fibonacci.sol"])
    fib = constructor.construct_contract(contracts_code["Fibonacci"])
    raw_contracts = constructor.generate_code([fib])
    contracts = [ArbContract(raw) for raw in raw_contracts]
    fib = contracts[0]
    vm = create_evm_vm(contracts)

    vm.env.send_message([fib.generateFib(total_count), 1234, 100000000, 0, 0])
    vm.env.deliver_pending()
    vm.env.deliver_pending()
    vm.env.send_message([fib.getFib(query_num), 2345, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)

    print("Contract sent messages:", vm.sent_messages)


def test_fibonacci_and_payments(total_count, query_num):
    contracts_code = prepare_contracts(["payment_channel.sol"])
    fib = constructor.construct_contract(contracts_code["Fibonacci"])
    channel = constructor.construct_contract(
        contracts_code["Channel"],
        fib["address"]
    )
    raw_contracts = constructor.generate_code([fib, channel])
    contracts = [ArbContract(raw) for raw in raw_contracts]
    fib = contracts[0]
    channel = contracts[1]
    vm = create_evm_vm(contracts)

    vm.env.send_message([fib.generateFib(total_count), 1234, 100000000, 0, 0])
    vm.env.deliver_pending()
    vm.env.deliver_pending()
    vm.env.send_message([fib.getFib(query_num), 2345, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)

    print("Contract sent messages:", vm.sent_messages)

    person_a = '0x1000000000000000000000000000000000000000'
    person_b = '0x2222222222222222222222222222222222222222'
    person_a_int = eth_utils.to_int(hexstr=person_a)
    person_b_int = eth_utils.to_int(hexstr=person_b)

    vm.env.send_message([channel.deposit(), person_a_int, 10000, 0, 0])
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    vm.env.send_message(
        [channel.transfer(person_b, 4000), person_a_int, 0, 0, 0]
    )
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    vm.env.send_message([channel.withdraw(3000), person_b_int, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    print("Contract sent messages:", vm.sent_messages)


def test_payment_channel():
    contracts_code = prepare_contracts(["PaymentChannel.sol"])
    # print(contracts_code)
    fib = constructor.construct_contract(contracts_code["Fibonacci"])
    channel = constructor.construct_contract(
        contracts_code["PaymentChannel"],
        fib["address"]
    )
    raw_contracts = constructor.generate_code([fib, channel])
    contracts = [ArbContract(raw) for raw in raw_contracts]
    fib = contracts[0]
    channel = contracts[1]
    vm = create_evm_vm(contracts)

    person_a = '0x1000000000000000000000000000000000000000'
    person_b = '0x2222222222222222222222222222222222222222'
    person_a_int = eth_utils.to_int(hexstr=person_a)
    person_b_int = eth_utils.to_int(hexstr=person_b)

    vm.env.send_message([channel.deposit(), person_a_int, 10000, 0, 0])
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    vm.env.send_message(
        [channel.transferFib(person_b, 14), person_a_int, 0, 0, 0]
    )
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    vm.env.send_message([channel.withdraw(10), person_b_int, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_a), 0, 0, 0, 0])
    vm.env.send_message([channel.getBalance(person_b), 0, 0, 0, 0])
    vm.env.deliver_pending()
    run_until_halt(vm)
    print("Contract sent messages:", vm.sent_messages)


if __name__ == '__main__':
    if len(sys.argv) != 3:
        raise Exception("Call as evm_demo.py [generate_count] [read_index]")

    if int(sys.argv[2]) >= int(sys.argv[1]):
        raise Exception("read_index must be less than the generate_count")

    # print("Fibonacci test")
    # test_fibonacci_and_payments(int(sys.argv[1]), int(sys.argv[2]))
    # print()
    print("Payment channel test")
    test_payment_channel()

    # test_erc20()
