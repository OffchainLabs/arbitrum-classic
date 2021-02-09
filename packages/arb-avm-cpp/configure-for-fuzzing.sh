#!/bin/bash
set -eu

cd "$(dirname "$0")"

# former works on linux, latter works on macos
cores="$(nproc --all || sysctl -n hw.logicalcpu)"

mkdir -p build
cd build
cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_C_COMPILER=hfuzz-clang -DCMAKE_CXX_COMPILER=hfuzz-clang++ ..
make -j"$cores"
cd ..
./scripts/install-cmachine-build
