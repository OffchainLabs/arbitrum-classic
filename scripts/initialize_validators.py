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
import shutil

import setup_states

# package configuration
NAME = "initialize_rollup_validators_rinkeby"
DESCRIPTION = "Manage Arbitrum dockerized deployments"

# filename constants
VALIDATOR_STATE_DIRNAME = "validator-states/validator"


### ----------------------------------------------------------------------------
### Deploy
### ----------------------------------------------------------------------------


# Compile contracts to `contract.ao` and export to Docker and run validators
def deploy(args, sudo_flag=False):

    if os.path.isdir(setup_states.VALIDATOR_STATES):
        shutil.rmtree(setup_states.VALIDATOR_STATES)

    setup_states.setup_validator_states_folder(args.contract, args.n_validators)

    config = {
        "rollup_address": args.rollup_address.strip(),
        "eth_url": args.eth_url,
        "blocktime": 13,
    }
    setup_states.setup_validator_configs(config, args.n_validators)


def check_file(name):
    if not os.path.isfile(name):
        raise argparse.ArgumentTypeError("%s is not a valid file" % name)
    return name


def main():
    parser = argparse.ArgumentParser(prog=NAME, description=DESCRIPTION)
    parser.add_argument(
        "contract", type=check_file, help="The Arbitrum bytecode contract to deploy"
    )
    parser.add_argument(
        "n_validators",
        choices=range(1, 101),
        metavar="[1-100]",
        type=int,
        help="The number of validators to deploy",
    )
    parser.add_argument(
        "rollup_address", help="The address of a deployed arbitrum rollup contract"
    )

    parser.add_argument("eth_url", help="RPC or Websocket url for Ethereum node")

    args = parser.parse_args()

    deploy(args)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(1)
