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

# ----------------------------------------------------------------------------
# arb-deploy
# ----------------------------------------------------------------------------

import argparse
import os
import sys
import json
import shutil

import setup_states
import build_validator_docker
from support.run import run

# package configuration
NAME = "arb-deploy"
DESCRIPTION = "Manage Arbitrum dockerized deployments"
ROOT_DIR = os.path.abspath(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

# filename constants
DOCKER_COMPOSE_FILENAME = "docker-compose.yml"
VALIDATOR_STATE_DIRNAME = "validator-states/validator"


### ----------------------------------------------------------------------------
### Deploy
### ----------------------------------------------------------------------------


# Compile contracts to `contract.ao` and export to Docker and run validators
def deploy(args, sudo_flag=False):

    if not build_validator_docker.is_built(sudo_flag):
        print("arb-validator is not build, building now")
        build_validator_docker.build_validator(sudo_flag)
        return

    if os.path.isdir(setup_states.VALIDATOR_STATES):
        shutil.rmtree(setup_states.VALIDATOR_STATES)

    if args.is_parity:
        image_name = "arb-bridge-eth"
        ws_port = 7546
    elif args.is_ganache:
        image_name = "arb-bridge-eth-ganache"
        ws_port = 7545
    elif args.is_geth:
        image_name = "arb-bridge-eth-geth"
        ws_port = 7546
    else:
        raise Exception("Must select either parity or ganache")

    setup_states.setup_validator_states_docker(
        args.contract, args.n_validators, image_name, args.is_geth, sudo_flag
    )

    with open(
        os.path.join(
            setup_states.VALIDATOR_STATES, "validator0/bridge_eth_addresses.json"
        )
    ) as json_file:
        data = json.load(json_file)
        factory_address = data["ArbFactory"]

    rollup_creation_cmd = (
        "docker run -it --network=arb-network -v %s:/home/user/state arb-validator create state/contract.ao state/private_key.txt ws://%s:%s %s"
        % (
            os.path.abspath("validator-states/validator0"),
            image_name,
            ws_port,
            factory_address,
        )
    )
    rollup_address = run(rollup_creation_cmd, capture_stdout=True, quiet=False)
    print("rollup_address", rollup_address)

    with open("validator-states/config.json", "w") as outfile:
        json.dump(
            {
                "rollup_address": rollup_address.strip(),
                "websocket_port": ws_port,
                "validator_count": args.n_validators,
            },
            outfile,
        )


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
        "--ganache",
        action="store_true",
        dest="is_ganache",
        help="Generate states based on arb-bridge-eth docker images",
    )
    group.add_argument(
        "--geth",
        action="store_true",
        dest="is_geth",
        help="Generate states based on arb-bridge-eth docker images",
    )
    group.add_argument(
        "--parity",
        action="store_true",
        dest="is_parity",
        help="Generate states based on arb-bridge-eth docker images",
    )

    args = parser.parse_args()
    deploy(args)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(1)
