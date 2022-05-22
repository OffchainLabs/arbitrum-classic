#!/bin/sh
cd "$(dirname "$0")"

cd ../../../arb-os
cargo run -- compile ../arb-avm-cpp/tests/wild-segments/src/main.mini -o ../arb-avm-cpp/tests/wild-segments/main.mexe
