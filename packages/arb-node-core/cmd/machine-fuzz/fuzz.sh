#!/bin/bash
set -eu
cd "$(dirname "$0")"

# former works on linux, latter works on macos
cores="$(nproc --all || sysctl -n hw.logicalcpu)"

mkdir -p corpus
cp ../../../arb-avm-cpp/tests/machine-cases/* corpus/
cp ../../../arb-os/arb_os/arbos.mexe corpus/
honggfuzz -i corpus -z -t 10 -P -n "$cores" -- ./machine-fuzz
