#!/bin/bash
set -eu
cd "$(dirname "$0")"

if [ $# -gt 0 ]; then
    cores="$1"
else
    # former works on linux, latter works on macos
    cores="$(nproc --all || sysctl -n hw.logicalcpu)"
fi

mkdir -p corpus
cp ../../../arb-avm-cpp/tests/machine-cases/* corpus/
cp ../../../arb-os/arb_os/arbos.mexe corpus/
honggfuzz -i corpus -z -t 10 -P -n "$cores" -- ./machine-fuzz
