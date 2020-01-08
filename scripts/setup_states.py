#!/usr/bin/env python3

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

import argparse
import json
import os
import shutil
from web3 import Web3
from eth_account import Account

from support.run import run

NAME = "setup_states"
DESCRIPTION = ""
VALIDATOR_STATES = "validator-states"
VALIDATOR_STATE = "validator%s"
ROOT_DIR = os.path.abspath(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


# Retrieve bridge_eth_addresses.json
# arb-bridge-eth must be have been built first
def setup_validator_states_docker(contract, n_validators, node_type, sudo=False):
    ethaddrs = "bridge_eth_addresses.json"

    if node_type == "parity":
        image_name = "arb-bridge-eth"
    elif node_type == "ganache":
        image_name = "arb-bridge-eth-ganache"
    else:
        print(node_type, "is bad")
        raise Exception("Unknown node type " + node_type)

    layer = run(
        "docker create %s" % image_name, capture_stdout=True, quiet=True, sudo=sudo
    ).strip()
    if layer == "":
        print("Docker image %s does not exist" % image_name)
        return
    run(
        "docker cp %s:/home/user/bridge_eth_addresses.json %s" % (layer, ethaddrs),
        sudo=sudo,
    )
    run("docker rm %s" % layer, quiet=True, sudo=sudo)

    setup_validator_states(
        contract,
        n_validators,
        ethaddrs,
        "0x81183C9C61bdf79DB7330BBcda47Be30c0a85064",
        "7545",
    )

    os.remove(ethaddrs)


def setup_validator_states(contract, n_validators, ethaddrs, source_address, port):
    ARB_VALIDATOR = os.path.join(ROOT_DIR, "packages", "arb-validator")

    # Check for validator_states in cwd
    if os.path.isdir(VALIDATOR_STATES):
        exit("Error: " + VALIDATOR_STATES + " exists in the current working directory")

    # Extract keys from acct_keys
    accounts = [Account.create() for _ in range(n_validators)]
    addresses = [account.address for account in accounts]
    privates = [account.key.hex()[2:] for account in accounts]

    web3 = Web3(Web3.HTTPProvider("http://localhost:" + str(port)))
    for dest in addresses:
        web3.eth.sendTransaction(
            {"to": dest, "from": source_address, "value": 100000000000000000000}
        )

    # Create VALIDATOR_STATES
    os.mkdir(VALIDATOR_STATES)
    for i in range(n_validators):
        state = os.path.join(VALIDATOR_STATES, VALIDATOR_STATE % i)
        os.mkdir(state)
        # contract.ao
        shutil.copyfile(contract, os.path.join(state, "contract.ao"))
        # bridge_eth_addresses.json
        shutil.copyfile(ethaddrs, os.path.join(state, "bridge_eth_addresses.json"))
        # server.crt and server.key
        shutil.copy(os.path.join(ARB_VALIDATOR, "server.crt"), state)
        shutil.copy(os.path.join(ARB_VALIDATOR, "server.key"), state)
        # validator_addresses.txt
        with open(os.path.join(state, "validator_addresses.txt"), "w") as f:
            f.write("\n".join(addresses))
        # private_key.txt
        with open(os.path.join(state, "private_key.txt"), "w") as f:
            f.write(privates[i])


def check_file(name):
    if not os.path.isfile(name):
        raise argparse.ArgumentTypeError("%s is not a valid file" % name)
    return name


def check_json(name):
    if not os.path.isfile(name):
        raise argparse.ArgumentTypeError("%s is not a valid file" % name)
    try:
        with open(name, "r") as f:
            json.loads(f.read())
    except ValueError:
        raise argparse.ArgumentTypeError("%s is not valid json" % name)
    return name


def main():
    parser = argparse.ArgumentParser(prog=NAME, description=DESCRIPTION)
    parser.add_argument(
        "contract", type=check_file, help="The Arbitrum bytecode contract to deploy"
    )
    parser.add_argument(
        "n_validators",
        choices=range(2, 101),
        metavar="[2-100]",
        type=int,
        help="The number of validators to deploy",
    )
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument(
        "--ganache-docker",
        action="store_true",
        dest="is_ganache",
        help="Generate states based on arb-bridge-eth docker images",
    )
    group.add_argument(
        "--parity-docker",
        action="store_true",
        dest="is_parity",
        help="Generate states based on arb-bridge-eth docker images",
    )
    group.add_argument(
        "--local",
        action="store_true",
        dest="is_local",
        help="Generate states based on local inputs",
    )
    parser.add_argument(
        "-a",
        "--funder_key",
        required=False,
        help="Unlocked key holding ETH to fund validators",
    )
    parser.add_argument(
        "-b",
        "--bridge_eth_addresses",
        required=False,
        type=check_json,
        help="EthBridge contract addresses",
    )
    parser.add_argument(
        "-p",
        "--port",
        required=False,
        type=int,
        default=7545,
        help="Port number to search for local node on",
    )

    args = parser.parse_args()

    if args.is_local:
        setup_validator_states(
            args.contract,
            args.n_validators,
            args.funder_key,
            args.bridge_eth_addresses,
            args.port,
        )
    elif args.is_parity:
        setup_validator_states_docker(args.contract, args.n_validators, "parity")
    elif args.is_ganache:
        setup_validator_states_docker(args.contract, args.n_validators, "ganache")


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        exit(1)
