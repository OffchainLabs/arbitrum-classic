# arb-avm-cpp

Arbitrum technologies are patent pending. This repository is offered under the Apache 2.0 license. See LICENSE for details.

## Bulding, testing, and running

```bash
mkdir release
cd release
conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite
conan install ..
cmake .. -DCMAKE_BUILD_TYPE=Release
cmake --build .
ctest .
./app/avm_runner contract.ao inbox.dat
```

## Formatting code

Format depends on clang-format which can be installed with `brew install clang-format` on the mac. Formatting can be run with:

```bash
make format
```
