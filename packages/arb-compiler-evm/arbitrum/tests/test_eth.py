from unittest import TestCase

from arbitrum import run_vm_once, value
from arbitrum.evm.contract import ArbContract, create_evm_vm, EVMCall, EVMInvalid

from pyevmasm import instruction_tables, assemble_hex, assemble_one, disassemble_one
import eth_utils
import random


def make_msg_val(calldata):
    return value.Tuple([calldata, 0, 0, 0])


def run_until_block(vm, test):
    while True:
        try:
            run = run_vm_once(vm)
            if not run:
                break
        except Exception as err:
            test.fail("VM run hit error " + err)
        if vm.halted:
            test.fail("VM unintentionally halted")

def make_evm_ext_code(op, address):
    instruction_table = instruction_tables['byzantium']
    return [
        assemble_one("PUSH20 " + address),
        op,
        assemble_one("PUSH1 0x00"),
        instruction_table['MSTORE'],
        assemble_one("PUSH1 0x20"),
        assemble_one("PUSH1 0x00"),
        instruction_table['RETURN'],
    ]


def make_evm_codecopy_code(offset, length, address):
    return_data_size = ((length + 31) // 32) * 32 + 64
    instruction_table = instruction_tables['byzantium']
    return [
        assemble_one("PUSH32 0x" + length.to_bytes(32, byteorder="big").hex()),  # length
        assemble_one("PUSH32 0x" + offset.to_bytes(32, byteorder="big").hex()),  # offset
        assemble_one("PUSH32 0x40"),  # destOffset
        assemble_one("PUSH20 " + address),
        instruction_table["EXTCODECOPY"],
        assemble_one("PUSH32 0x20"),
        assemble_one("PUSH1 0x00"),
        instruction_table['MSTORE'],
        assemble_one("PUSH32 0x" + length.to_bytes(32, byteorder="big").hex()),
        assemble_one("PUSH1 0x20"),
        instruction_table['MSTORE'],
        assemble_one("PUSH32 0x" + return_data_size.to_bytes(32, byteorder="big").hex()),
        assemble_one("PUSH1 0x00"),
        instruction_table['RETURN'],
    ]


def make_contract(evm_code, return_type):
    return ArbContract({
        "address": "0x895521964D724c8362A36608AAf09A3D7d0A0445",
        "abi": [{
            "constant": False,
            "inputs": [],
            "name": "testMethod",
            "outputs": [{
                "name": "",
                "type": return_type
            }],
            "payable": False,
            "stateMutability": "view",
            "type": "function"
        }],
        "name": "TestContract",
        "code": assemble_hex(evm_code),
        "storage": {}
    })


def create_many_contract_vm(contract_a):
    contracts = [contract_a]
    for i in range(10):
        contracts.append(ArbContract({
            "address": eth_utils.to_checksum_address(
                random.getrandbits(8*20).to_bytes(20, byteorder="big").hex()
            ),
            "abi": [],
            "name": "TestContract",
            "code": "0x00",
            "storage": {}
        }))
    return create_evm_vm(contracts)


class TestEVM(TestCase):
    def test_codesize_succeed(self):
        instruction_table = instruction_tables['byzantium']
        evm_code = make_evm_ext_code(
            instruction_table["EXTCODESIZE"],
            "0x895521964D724c8362A36608AAf09A3D7d0A0445"
        )
        code_size = len(evm_code) + sum(op.operand_size for op in evm_code)
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMCall)
        self.assertEqual(parsed_out.output_values[0], code_size)

    def test_codesize_fail(self):
        instruction_table = instruction_tables['byzantium']
        evm_code = make_evm_ext_code(
            instruction_table["EXTCODESIZE"],
            "0x9999"
        )
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMInvalid)

    def test_codehash_succeed(self):
        evm_code = make_evm_ext_code(
            disassemble_one(bytes.fromhex('3f')),
            "0x895521964D724c8362A36608AAf09A3D7d0A0445"
        )
        hex_code = assemble_hex(evm_code)
        code_hash = int.from_bytes(
            eth_utils.crypto.keccak(hexstr=hex_code),
            byteorder="big"
        )
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMCall)
        self.assertEqual(parsed_out.output_values[0], code_hash)

    def test_codehash_fail(self):
        evm_code = make_evm_ext_code(
            disassemble_one(bytes.fromhex('3f')),
            "0x9999"
        )
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMInvalid)

    def test_codecopy_succeed(self):
        offset = 30
        length = 80
        evm_code = make_evm_codecopy_code(
            offset,
            length,
            "0x895521964D724c8362A36608AAf09A3D7d0A0445"
        )
        hex_code = assemble_hex(evm_code)
        contract_a = make_contract(evm_code, "bytes")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMCall)
        self.assertEqual(
            parsed_out.output_values[0].hex(),
            hex_code[2 + offset*2:2 + offset*2 + length*2]
        )

    def test_codecopy_fail(self):
        offset = 30
        length = 80
        evm_code = make_evm_codecopy_code(
            offset,
            length,
            "0x9999"
        )
        contract_a = make_contract(evm_code, "bytes")
        vm = create_evm_vm([contract_a])
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 0, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMInvalid)

    def test_balance_succeed(self):
        instruction_table = instruction_tables['byzantium']
        evm_code = make_evm_ext_code(
            instruction_table["BALANCE"],
            "0x895521964D724c8362A36608AAf09A3D7d0A0445"
        )
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 62244, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMCall)
        self.assertEqual(parsed_out.output_values[0], 62244)

    def test_balance_fail(self):
        instruction_table = instruction_tables['byzantium']
        evm_code = make_evm_ext_code(
            instruction_table["BALANCE"],
            "0x9999"
        )
        contract_a = make_contract(evm_code, "uint256")
        vm = create_many_contract_vm(contract_a)
        vm.env.send_message([make_msg_val(contract_a.testMethod(4)), 2345, 62244, 0])
        vm.env.deliver_pending()
        run_until_block(vm, self)
        self.assertEqual(len(vm.logs), 1)
        val = vm.logs[0]
        parsed_out = vm.output_handler(val)
        self.assertIsInstance(parsed_out, EVMInvalid)
