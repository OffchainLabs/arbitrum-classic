### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM offchainlabs/cpp-base:0.3.2 as arb-avm-cpp

# Copy external dependencies
COPY --chown=user arb-avm-cpp/CMakeLists.txt /home/user/arb-avm-cpp/
COPY --chown=user arb-avm-cpp/external /home/user/arb-avm-cpp/external
COPY --chown=user arb-avm-cpp/cmake /home/user/arb-avm-cpp/cmake
# Build arb-avm-cpp
RUN mkdir -p arb-avm-cpp/build && cd arb-avm-cpp/build && \
    cmake -DCMAKE_EXPORT_COMPILE_COMMANDS=ON .. -DCMAKE_BUILD_TYPE=RelWithDebInfo -DBUILD_TESTING=0 -DENABLE_JEMALLOC=true && \
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

env GOPROXY https://goproxy.io,direct

FROM offchainlabs/backend-base:0.3.3 as arb-validator-builder

# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-node-core/go.* /home/user/arb-node-core/
COPY --chown=user arb-rpc-node/go.* /home/user/arb-rpc-node/
COPY --chown=user arb-evm/go.* /home/user/arb-evm/
RUN cd arb-rpc-node && GOPROXY=https://goproxy.io,direct go mod download

# Copy source code
COPY --chown=user arb-util/ /home/user/arb-util/
RUN cd arb-util && GOPROXY=https://goproxy.io,direct go build -v ./...

COPY --chown=user arb-evm/ /home/user/arb-evm/
RUN cd arb-evm && GOPROXY=https://goproxy.io,direct go build -v ./...

RUN cd arb-node-core && GOPROXY=https://goproxy.io,direct go build -v ./...

COPY --chown=user arb-avm-cpp/ /home/user/arb-avm-cpp/
COPY --chown=user arb-node-core/ /home/user/arb-node-core/
COPY --chown=user arb-rpc-node/ /home/user/arb-rpc-node/

COPY --from=arb-avm-cpp /home/user/arb-avm-cpp/build/lib /home/user/arb-avm-cpp/build/lib
COPY --from=arb-avm-cpp /home/user/arb-avm-cpp/cmachine/flags.go /home/user/arb-avm-cpp/cmachine/
COPY --from=arb-avm-cpp /home/user/arb-avm-cpp/external/teesdk /home/user/arb-avm-cpp/external/teesdk
COPY --from=arb-avm-cpp /home/user/.hunter /home/user/.hunter

# Build arb-validator
RUN cd arb-node-core && go install -v ./cmd/arb-validator && go install -v ./cmd/arb-relay && \
    cd ../arb-rpc-node && go install -v ./cmd/arb-node && go install -v ./cmd/arb-dev-node

FROM offchainlabs/cpp-base:0.3.2 as arb-validator
# Export binary
##########################################################################################################
# Set these to use TEESDK, further better way maybe provided
ENV LD_LIBRARY_PATH /home/user/arb-avm-cpp/external/teesdk
ENV TEESDK_PUB /home/user/arb-avm-cpp/external/teesdk/examples/auditors/godzilla/godzilla.public.der
ENV TEESDK_PRI /home/user/arb-avm-cpp/external/teesdk/examples/auditors/godzilla/godzilla.sign.sha256
ENV TEESDK_CONF /home/user/arb-avm-cpp/external/teesdk/examples/enclave_info.toml
ENV TEESDK_METHOD echo
ENV TEESDK_ARGS "Hello Eigen"
ENV TEESDK_UID "uid";
ENV TEESDK_TOKEN "token"
##########################################################################################################

COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin
COPY --chown=user --from=arb-validator-builder /home/user/arb-avm-cpp/external/teesdk /home/user/arb-avm-cpp/external/teesdk

ENTRYPOINT ["/home/user/go/bin/arb-validator"]
EXPOSE 8547 8548
