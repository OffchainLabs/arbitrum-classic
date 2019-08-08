### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:3.9 as arb-avm-cpp
# Alpine dependencies
RUN apk add --no-cache boost-dev=1.67.0-r2 cmake=3.13.0-r0 g++=8.3.0-r0 \
    make=4.2.1-r2 musl-dev=1.1.20-r5 python3-dev=3.6.8-r2 && \
    pip3 install conan==1.18.1 && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"
# Build dependencies
COPY --chown=user arb-avm-cpp/conanfile.txt ./
RUN mkdir -p build && cd build && \
    conan profile new default --detect && \
    conan profile update settings.compiler.libcxx=libstdc++11 default && \
    conan remote add nonstd-lite https://api.bintray.com/conan/martinmoene/nonstd-lite && \
    conan install ..
# Copy source code
COPY --chown=user arb-avm-cpp/ ./
# Copy build cache
COPY --from=arb-validator --chown=user /cpp-build build/
# Build arb-avm-cpp
RUN cd build && conan install .. && \
    cmake .. -DCMAKE_BUILD_TYPE=Release && \
    cmake --build . -j $(nproc) && \
    cp lib/* ../cmachine


FROM alpine:3.9 as arb-validator-builder
# Alpine dependencies
RUN apk add --no-cache build-base=0.5-r1 git=2.20.1-r0 go=1.11.5-r0 \
    libc-dev=0.7.1-r0 linux-headers=4.18.13-r1 && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/arb-validator"
# Build dependencies
COPY --chown=user arb-avm-cpp/go.* /home/user/arb-avm-cpp/
COPY --chown=user arb-avm-go/go.* /home/user/arb-avm-go/
COPY --chown=user arb-util/go.* /home/user/arb-util/
COPY --chown=user arb-validator/go.* /home/user/arb-validator/
RUN go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp=../arb-avm-cpp && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-avm-go=../arb-avm-go && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-util=../arb-util && \
    cd ../arb-avm-go && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-util=../arb-util && \
    cd ../arb-validator && \
    go mod download
# Copy source code
COPY --from=arb-avm-cpp /home/user/go.mod /home/user/go.sum /home/user/arb-avm-cpp/
COPY --from=arb-avm-cpp /home/user/cavm/cmachine.h /home/user/arb-avm-cpp/cavm/cmachine.h
COPY --from=arb-avm-cpp /home/user/cmachine /home/user/arb-avm-cpp/cmachine/
COPY --chown=user arb-avm-go/ /home/user/arb-avm-go/
COPY --chown=user arb-util/ /home/user/arb-util/
COPY --chown=user arb-validator/ /home/user/arb-validator/
# Copy build cache
COPY --from=arb-validator --chown=user /build /home/user/.cache/go-build
# Build arb-validator
RUN go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-avm-cpp=../arb-avm-cpp && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-avm-go=../arb-avm-go && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-util=../arb-util && \
    cd ../arb-avm-go && \
    go mod edit -replace github.com/offchainlabs/arbitrum/packages/arb-util=../arb-util && \
    cd ../arb-validator && \
    go install -v ./cmd/followerServer ./cmd/coordinatorServer


FROM alpine:3.9 as arb-validator
# Export binary
RUN apk add --no-cache libstdc++=8.3.0-r0 libgcc=8.3.0-r0 && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"
COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin
COPY --chown=user arb-validator/server.crt arb-validator/server.key ./

ENV ID=0 \
    WAIT_FOR="arb-bridge-eth:7545" \
    ETH_URL="ws://arb-bridge-eth:7545" \
    COORDINATOR_URL="" \
    AVM="cpp" \
    PATH="/home/user/go/bin:${PATH}"

# Build cache
COPY --chown=user --from=arb-validator-builder /home/user/.cache/go-build /build
COPY --from=arb-avm-cpp /home/user/build /cpp-build

# 1) Waits for host:port if $WAIT_FOR is set
# 2) Copies address files from ../ to ./ (state volume)
# 3) Launches follower if $COORDINATOR_URL else launches coordinator
CMD if [[ ! -z ${WAIT_FOR} ]]; then \
sleep 2 && while ! nc -z ${WAIT_FOR//:/ }; do sleep 2; done && sleep 2; \
echo "Finished waiting for ${WAIT_FOR}..."; else echo "Starting..."; fi && \
T=follower; if [[ -z ${COORDINATOR_URL} ]]; then T=coordinator; fi; cd state &&\
${T}Server --avm=${AVM} contract.ao private_key.txt validator_addresses.txt \
    bridge_eth_addresses.json ${ETH_URL} ${COORDINATOR_URL}
EXPOSE 1235 1236
