### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.2.5 as arb-avm-cpp

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

FROM offchainlabs/backend-base:0.2.8 as arb-validator-builder

# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-validator-core/go.* /home/user/arb-validator-core/
COPY --chown=user arb-evm/go.* /home/user/arb-evm/
COPY --chown=user arb-tx-aggregator/go.* /home/user/arb-tx-aggregator/
RUN cd arb-validator && go mod download && cd ../arb-tx-aggregator && go mod download

# Copy source code
COPY --chown=user arb-util/ /home/user/arb-util/
RUN cd arb-util && go build -v ./...

COPY --chown=user arb-validator-core/ /home/user/arb-validator-core/
RUN cd arb-validator-core && go build -v ./...

COPY --chown=user arb-evm/ /home/user/arb-evm/
RUN cd arb-evm && go build -v ./...


COPY --chown=user arb-avm-cpp/ /home/user/arb-avm-cpp/
COPY --chown=user arb-validator/ /home/user/arb-validator/
COPY --chown=user arb-tx-aggregator/ /home/user/arb-tx-aggregator/

COPY --from=arb-avm-cpp /home/user/cmachine /home/user/arb-avm-cpp/cmachine/

# Build arb-validator
RUN cd arb-validator && go install -v ./cmd/arb-validator && \
    cd ../arb-tx-aggregator && go install -v ./cmd/arb-tx-aggregator

FROM offchainlabs/cpp-base:0.2.5 as arb-validator
# Export binary

COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 8547 8548
