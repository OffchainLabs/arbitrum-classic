### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.3.0 as arb-avm-cpp

# Copy external dependencies
COPY --chown=user arb-avm-cpp/CMakeLists.txt .
COPY --chown=user arb-avm-cpp/external ./external
COPY --chown=user arb-avm-cpp/cmake ./cmake
COPY --chown=user arb-avm-cpp/scripts ./scripts
# Build arb-avm-cpp
RUN mkdir -p build && cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/avm_values ./avm_values
RUN cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/avm ./avm
RUN cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/data_storage ./data_storage
RUN cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/cavm ./cavm
RUN cd build && \
    cmake .. -DCMAKE_BUILD_TYPE=Release -DBUILD_TESTING=0 && \
    cmake --build . -j $(nproc) && \
    cd ../ && \
    ./scripts/install-cmachine-build

FROM offchainlabs/backend-base:0.3.0 as arb-validator-builder

# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-node-core/go.* /home/user/arb-node-core/
COPY --chown=user arb-rpc-node/go.* /home/user/arb-rpc-node/
COPY --chown=user arb-evm/go.* /home/user/arb-evm/
RUN cd arb-rpc-node && go mod download

# Copy source code
COPY --chown=user arb-util/ /home/user/arb-util/
RUN cd arb-util && go build -v ./...

COPY --chown=user arb-evm/ /home/user/arb-evm/
RUN cd arb-evm && go build -v ./...


RUN cd arb-node-core && go build -v ./...

COPY --chown=user arb-avm-cpp/ /home/user/arb-avm-cpp/
COPY --chown=user arb-node-core/ /home/user/arb-node-core/
COPY --chown=user arb-rpc-node/ /home/user/arb-rpc-node/

COPY --from=arb-avm-cpp /home/user/cmachine /home/user/arb-avm-cpp/cmachine/

# Build arb-validator
RUN cd arb-node-core && go install -v ./cmd/arb-validator && \
    cd ../arb-rpc-node && go install -v ./cmd/arb-node && go install -v ./cmd/arb-dev-node

FROM offchainlabs/cpp-base:0.3.0 as arb-validator
# Export binary

COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 8547 8548
