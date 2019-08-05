[![Build Status](https://travis-ci.com/OffchainLabs/arb-avm-cpp.svg?branch=master)](https://travis-ci.com/OffchainLabs/arb-avm-cpp) [![codecov](https://codecov.io/gh/OffchainLabs/arb-avm-cpp/branch/master/graph/badge.svg)](https://codecov.io/gh/OffchainLabs/arb-avm-cpp)
# arb-avm-cpp

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

## Docker

Build with:

```
docker build -t arb-avm-cpp .
```

To enable caching, remove the `##DEV_` prefixes. Note that there must be at least one
successful `arb-avm-cpp` image build before using the cache will work. Enable caching with:

```
awk '{gsub(/##DEV_/, "")}1' Dockerfile > .tmp.Dockerfile && mv .tmp.Dockerfile Dockerfile
```

Exported file locations:

```
/build/cmachine.h
/build/avm/libavm.a
```
