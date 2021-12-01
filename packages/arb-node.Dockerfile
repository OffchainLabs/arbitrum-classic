### --------------------------------------------------------------------
### Dockerfile
### arb-node
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.4.1 as arb-avm-cpp

# Copy external dependencies
COPY --chown=user arb-avm-cpp/CMakeLists.txt /home/user/arb-avm-cpp/
COPY --chown=user arb-avm-cpp/external /home/user/arb-avm-cpp/external
COPY --chown=user arb-avm-cpp/cmake /home/user/arb-avm-cpp/cmake
# Build arb-avm-cpp
RUN mkdir -p arb-avm-cpp/build && cd arb-avm-cpp/build && \
    cmake .. -DCMAKE_BUILD_TYPE=RelWithDebInfo -DBUILD_TESTING=0 -DENABLE_TCMALLOC=true && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/avm_values /home/user/arb-avm-cpp/avm_values
RUN cd arb-avm-cpp/build && \
    cmake .. && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/avm /home/user/arb-avm-cpp/avm
RUN cd arb-avm-cpp/build && \
    cmake .. && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/data_storage /home/user/arb-avm-cpp/data_storage
RUN cd arb-avm-cpp/build && \
    cmake .. && \
    cmake --build . -j $(nproc)

COPY --chown=user arb-avm-cpp/cavm /home/user/arb-avm-cpp/cavm
COPY --chown=user arb-avm-cpp/cmachine/flags.go.in /home/user/arb-avm-cpp/cmachine/
RUN cd arb-avm-cpp/build && \
    cmake .. && \
    cmake --build . -j $(nproc)

FROM offchainlabs/backend-base:0.4.1 as arb-node-builder

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

COPY --from=arb-avm-cpp /home/user/arb-avm-cpp/build/lib /home/user/arb-avm-cpp/build/lib
COPY --from=arb-avm-cpp /home/user/arb-avm-cpp/cmachine/flags.go /home/user/arb-avm-cpp/cmachine/
COPY --from=arb-avm-cpp /home/user/.hunter /home/user/.hunter

# Build arb-node
RUN cd arb-node-core && go install -v ./cmd/arb-validator && go install -v ./cmd/arb-relay && \
    cd ../arb-rpc-node && go install -v ./cmd/arb-node && go install -v ./cmd/arb-dev-node

FROM offchainlabs/dist-base:0.4.1 as arb-node
# Export binary

COPY --chown=user --from=arb-node-builder /home/user/go/bin /home/user/go/bin
COPY --chown=user arb-os/arb_os/arbos.mexe /home/user/arb-os/arb_os/
RUN mkdir -p /home/user/.arbitrum && \
    chown 1000:1000 /home/user/.arbitrum && \
    curl https://raw.githubusercontent.com/OffchainLabs/arb-os/48bdb999a703575d26a856499e6eb3e17691e99d/arb_os/arbos.mexe --output /home/user/.arbitrum/mainnet.arb1.mexe && \
    curl https://raw.githubusercontent.com/OffchainLabs/arb-os/26ab8d7c818681c4ee40792aeb12981a8f2c3dfa/arb_os/arbos.mexe --output /home/user/.arbitrum/testnet.rinkeby.mexe

ENTRYPOINT ["/home/user/go/bin/arb-node"]
EXPOSE 8547 8548
