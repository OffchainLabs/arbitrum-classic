#!/usr/bin/env python
# arb.py

import os
import subprocess
import sys

# Constants
CONTRACT_IMAGE="arb-contract"
CONTRACT_DOCKERFILE=".arb-contract.Dockerfile"

# Compile contracts to `contract.ao` and export to Docker and run validators
def build():
    # Check for compose folder and get dependencies
    if not os.path.isdir("./compose"):
        run("mkdir compose")
        run("git clone https://github.com/OffchainLabs/arb-ethbridge.git ./compose/arb-ethbridge")
        run("git clone https://github.com/OffchainLabs/arb-validator.git ./compose/arb-validator")
        run("git clone https://github.com/OffchainLabs/arb-avm.git ./compose/arb-validator/arb-avm")

    run("truffle migrate --network arbitrum")
    run("arbc-truffle-compile compiled.json contract.ao")
    run("sudo docker build -t %s -f %s ." % (CONTRACT_IMAGE, CONTRACT_DOCKERFILE))
    run("sudo docker-compose up --build")

# Run commands in shell
def run(command):
    print("\033[1m$ %s\n\033[0m" % command)
    subprocess.call(command.split())

if __name__ == "__main__":
    try:
        build()
    except KeyboardInterrupt:
        sys.exit(1)
