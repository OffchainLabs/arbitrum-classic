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

from support.run import run

NAME = "setup_states"
DESCRIPTION = ""
VALIDATOR_STATES = "validator-states"
VALIDATOR_STATE = "validator%s"
ROOT_DIR = os.path.abspath(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


# Retrieve bridge_eth_addresses.json and keys.json
# arb-bridge-eth must be have been built first
def setup_validator_states_ethbridge(contract, n_validators, sudo=False):
    keys = "keys.json"
    ethaddrs = "bridge_eth_addresses.json"

    layer = run(
        "docker create arb-bridge-eth", capture_stdout=True, quiet=True, sudo=sudo
    ).strip()
    if layer == "":
        print("Docker image arb-bridge-eth does not exist")
        return
    run(
        "docker cp %s:/home/user/bridge_eth_addresses.json %s" % (layer, ethaddrs),
        sudo=sudo,
    )
    run("docker cp %s:/home/user/keys.json %s" % (layer, keys), sudo=sudo)
    run("docker rm %s" % layer, quiet=True, sudo=sudo)

    setup_validator_states(contract, n_validators, keys, ethaddrs)

    os.remove(keys)
    os.remove(ethaddrs)


def setup_validator_states(contract, n_validators, acct_keys, ethaddrs):
    ARB_VALIDATOR = os.path.join(ROOT_DIR, "packages", "arb-validator")

    # Check for validator_states in cwd
    if os.path.isdir(VALIDATOR_STATES):
        exit("Error:", VALIDATOR_STATES, "exists in the current working directory")

    # Extract keys from acct_keys
    with open(acct_keys, "r") as f:
        data = json.loads(f.read())["addresses"]
    addresses = [addr for addr in list(data.keys())][-n_validators:]
    privates = []
    for key in [data[addr]["secretKey"]["data"] for addr in list(data.keys())]:
        privates.append("".join([hex(byte)[2:].zfill(2) for byte in key]))
    privates = privates[-n_validators:]

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
        "--docker",
        action="store_true",
        dest="is_ethbridge",
        help="Generate states based on arb-bridge-eth docker images",
    )
    group.add_argument(
        "--local",
        action="store_false",
        dest="is_ethbridge",
        help="Generate states based on local inputs",
    )
    parser.add_argument(
        "-a",
        "--acctKeys",
        type=check_json,
        required=False,
        help='Generate with: ganache-cli --acctKeys keys.json -m "$MNEMONIC" -a "$NUM_WALLETS"',
    )
    parser.add_argument(
        "-b",
        "--bridge_eth_addresses",
        type=check_json,
        help="EthBridge contract addresses",
    )
    parser.add_argument(
        "-p",
        "--port",
        type=int,
        default=7545,
        help="Port number to search for Ganache on",
    )

    args = parser.parse_args()

    print("is_ethbridge", args.is_ethbridge)

    if args.is_ethbridge:
        setup_validator_states_ethbridge(args.contract, args.n_validators)
    else:
        setup_validator_states(
            args.contract, args.n_validators, args.acctKeys, args.bridge_eth_addresses
        )


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        exit(1)
