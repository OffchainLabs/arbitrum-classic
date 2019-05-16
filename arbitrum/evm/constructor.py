import web3
from ethereum import vm as EthVM
from ethereum.pow.chain import Chain
from ethereum.messages import Transaction, VMExt
from ethereum import messages
from ethereum import utils


class VmExtBase():
    def __init__(self):
        self.storage = {}
        self.block_prevhash = 0
        self.block_coinbase = 0
        self.block_timestamp = 0
        self.block_number = 0
        self.block_difficulty = 0
        self.block_gas_limit = 0

        self.tx_origin = b'0' * 40
        self.tx_gasprice = 0

        self.log_storage = lambda addr: 0
        self.add_suicide = lambda addr: 0
        self.add_refund = lambda x: 0
        self.post_anti_dos_hardfork = lambda: True
        self.post_spurious_dragon_hardfork = lambda: True

    def create(self, msg):
        print("Creating", msg)
        return 0, 0, 0

    def call(self, msg):
        print("Call")
        return 0, 0, 0

    def sendmsg(self, msg):
        print("sendmsg")
        return 0, 0, 0

    def get_code(self, addr):
        print("Get code")
        return b''

    def log(self, addr, topics, data):
        print("log")
        return 0

    def get_balance(self, addr):
        print("get balance")
        return 0

    def set_balance(self, addr, balance):
        print("set balance")
        return 0

    def set_storage_data(self, addr, key, value):
        if addr not in self.storage:
            self.storage[addr] = {}
        self.storage[addr][key] = value

    def get_storage_data(self, addr, key):
        if addr not in self.storage:
            self.storage[addr] = {}
        if key not in self.storage[addr]:
            self.storage[addr][key] = 0
        return self.storage[addr][key]


class ConstructorMSG:
    def __init__(self, contract_num):
        self.gas = 100000000
        self.value = 0
        self.static = False
        self.to = contract_num


def construct_contract(contract, *args):
    w3 = web3.Web3()
    interface = w3.eth.contract(
        address=contract["address"],
        abi=contract["abi"],
        bytecode=contract["code"]
    )
    return {
        "name": contract["name"],
        "address": contract["address"],
        "abi": contract["abi"],
        "code": interface.constructor(*args).data_in_transaction
    }


def generate_code(contracts):
    ext = VmExtBase()
    new_contracts = {}
    for contract in contracts:
        raw_code = bytes.fromhex(contract["code"][2:])
        contructor_output = EthVM.vm_execute(
            ext,
            ConstructorMSG(contract["address"]),
            raw_code
        )
        new_contracts[contract["address"]] = {
            "name": contract["name"],
            "address": contract["address"],
            "abi": contract["abi"],
            "code": '0x' + contructor_output[2].hex(),
            "storage": {}
        }
    for contract in ext.storage:
        storage = {}
        for key in ext.storage[contract]:
            storage[hex(key)] = hex(ext.storage[contract][key])
        new_contracts[contract]["storage"] = storage
    return list(new_contracts.values())
