FROM offchainlabs/build-base:0.1.1

COPY --chown=user . /home/user/arbitrum/

WORKDIR /home/user/arbitrum
RUN mv /home/user/vconan /home/user/arbitrum/vconan
RUN yarn
RUN yarn install:validator
RUN yarn install:ci

WORKDIR /home/user/arbitrum/packages/arb-avm-cpp/ci
RUN ctest . && \
    lcov --capture --directory . --output-file coverage.info && \
    lcov --remove coverage.info --output-file coverage.info '*/.conan/*' '/usr/*' '*/tests/*' '*/external/*' && \
    lcov --list coverage.info

WORKDIR /home/user/arbitrum/packages/arb-avm-go
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./...

WORKDIR /home/user/arbitrum/packages/arb-bridge-eth
RUN truffle test || true

WORKDIR /home/user/arbitrum/packages/arb-compiler-evm
RUN coverage run --source=arbitrum/ setup.py test && \
    cd tests/sol-syscall && \
    truffle migrate --reset --compile-all --network arbitrum && \
    coverage run --source=../../arbitrum/ truffle_runner.py compiled.json && \
    cd ../..

WORKDIR /home/user/arbitrum/packages/arb-provider-ethers
RUN yarn jest --coverage

WORKDIR /home/user/arbitrum/packages/arb-provider-go
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./...

WORKDIR /home/user/arbitrum/packages/arb-provider-truffle
RUN yarn jest --coverage --pass-with-no-tests

WORKDIR /home/user/arbitrum/packages/arb-provider-web3
RUN yarn jest --coverage --pass-with-no-tests

WORKDIR /home/user/arbitrum/packages/arb-util
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./...

WORKDIR /home/user/arbitrum/packages/arb-validator
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./...
