FROM offchainlabs/build-base:0.1.0

COPY --chown=user . /home/user/arbitrum/

WORKDIR /home/user/arbitrum
RUN yarn
RUN yarn install:validator
RUN yarn install:ci

WORKDIR /home/user/arbitrum/packages/arb-avm-cpp/ci
RUN ctest . && \
    lcov --capture --directory . --output-file coverage.info && \
    lcov --remove coverage.info --output-file coverage.info '/usr/*' '*/tests/*' '*/external/*' && \
    lcov --list coverage.info && \
    codecov --file coverage.info --root /home/user/arbitrum --flags arb-avm-cpp

WORKDIR /home/user/arbitrum/packages/arb-avm-go
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./... && \
    codecov --root /home/user/arbitrum --flags arb-avm-go

WORKDIR /home/user/arbitrum/packages/arb-bridge-eth
RUN truffle test || true && \
    codecov --root /home/user/arbitrum --flags arb-bridge-eth

WORKDIR /home/user/arbitrum/packages/arb-compiler-evm
RUN coverage run --source=arbitrum/ setup.py test && \
    cd tests/sol-syscall && \
    truffle migrate --reset --compile-all --network arbitrum && \
    coverage run --source=../../arbitrum/ truffle_runner.py compiled.json && \
    cd ../.. && \
    coverage combine .coverage tests/sol-syscall/.coverage && \
    codecov --root /home/user/arbitrum --flags arb-compiler-evm

WORKDIR /home/user/arbitrum/packages/arb-provider-ethers
RUN yarn jest --coverage && \
    codecov --root /home/user/arbitrum --flags arb-provider-ethers

WORKDIR /home/user/arbitrum/packages/arb-provider-go
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./... && \
    codecov --root /home/user/arbitrum --flags arb-provider-go

WORKDIR /home/user/arbitrum/packages/arb-provider-truffle
RUN yarn jest --coverage --pass-with-no-tests && \
    codecov --root /home/user/arbitrum --flags arb-provider-truffle

WORKDIR /home/user/arbitrum/packages/arb-provider-web3
RUN yarn jest --coverage --pass-with-no-tests && \
    codecov --root /home/user/arbitrum --flags arb-provider-web3

WORKDIR /home/user/arbitrum/packages/arb-util
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./... && \
    codecov --root /home/user/arbitrum --flags arb-util

WORKDIR /home/user/arbitrum/packages/arb-validator
RUN go test -race -coverprofile=coverage.txt -covermode=atomic ./... && \
    codecov --root /home/user/arbitrum --flags arb-validator
