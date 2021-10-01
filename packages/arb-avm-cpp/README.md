# arb-avm-cpp

This repository is offered under the Apache 2.0 license. See LICENSE for details.

## Bulding and testing

```bash
mkdir release
cd release
cmake .. -DCMAKE_BUILD_TYPE=Release
cmake --build .
ctest .
```

## Formatting code

Format depends on clang-format which can be installed with `brew install clang-format` on the mac. Formatting can be run with:

```bash
make format
```
