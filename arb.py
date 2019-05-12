#!/usr/bin/env python
# arb.py
# Compiles truffle projects with the arbitrum provider into Arbitrum bytecode
# `contract.ao` and starts docker-compose with the `contract.ao` bytecode.

from argparse import ArgumentParser
import os
import subprocess
import sys

# contract.ao export constants
EXPORT_IMAGE="arb-app"
EXPORT_FILENAME=(".%s.Dockerfile" % EXPORT_IMAGE)
EXPORT_DOCKERFILE="FROM scratch\nCOPY contract.ao ./"

"""
Compiles Solidity contracts into Arbitrum bytecode file `contract.ao` and
exports as the docker image EXPORT_IMAGE. Finally runs `docker-compose build`.
"""
def build():
    # Compile contract.ao
    run("truffle migrate --network arbitrum")
    run("arbc-truffle-compile compiled.json contract.ao")

    # Export contract.ao
    with open(EXPORT_FILENAME, 'w') as d:
        d.write(EXPORT_DOCKERFILE)
    run("sudo docker build -t %s -f %s ." % (EXPORT_IMAGE, EXPORT_FILENAME))
    os.remove(EXPORT_FILENAME)

    # Build arb-ethbridge and arb-validators and run
    run("sudo docker-compose build")
    run("sudo docker-compose up")

"""
Pretty print and run commands
"""
def run(command):
    print("\033[1m$ %s\n\033[0m" % command)
    subprocess.call(command.split())

"""
Command line interface
"""
def main():
    # Arguments
    parser = ArgumentParser()

    parser.add_argument('cwd', nargs="?", default=".",
        help="Choose working directory")

    #parser.add_argument("-v", "--verbose", dest="verbose", default=False,
    #    help="Print all messages from `docker-compose up`")

    args = parser.parse_args()
    
    # Set current working directory to args.cwd
    if os.access(args.cwd, os.W_OK):
        os.chdir(args.cwd)
    else:
        sys.exit("Argument error: cannot write to directory %s" % args.cwd)
    
    # Build and run
    build()

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(1)
