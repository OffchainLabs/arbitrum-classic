### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:3.9 as arb-avm-cpp-builder
# Alpine dependencies
RUN apk add --no-cache boost-dev cmake g++ gcc make musl-dev python3 python3-dev && \
    pip3 install conan && \
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
COPY --from=arb-avm-cpp --chown=user /build build/
# Build arb-avm-cpp
RUN cd build && conan install .. && \
    cmake .. -DCMAKE_BUILD_TYPE=Release && \
    cmake --build . -j $(nproc) && \
    cp lib/* ../cmachine
FROM scratch as arb-avm-cpp
# Export library binary and header
COPY --from=arb-avm-cpp-builder /home/user/go.mod /home/user/go.sum arb-avm-cpp/
COPY --from=arb-avm-cpp-builder /home/user/cavm/cmachine.h arb-avm-cpp/cavm/cmachine.h
COPY --from=arb-avm-cpp-builder /home/user/cmachine arb-avm-cpp/cmachine/
COPY --from=arb-avm-cpp-builder /home/user/build build/


FROM alpine:3.9 as arb-validator-builder
# Alpine dependencies
RUN apk add --no-cache build-base git go libc-dev linux-headers && \
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
COPY --from=arb-avm-cpp --chown=user /arb-avm-cpp/ /home/user/arb-avm-cpp/
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
RUN apk add --no-cache libstdc++ libgcc && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"
COPY --chown=user --from=arb-validator-builder /home/user/go/bin /home/user/go/bin
COPY --chown=user --from=arb-bridge-eth     \
    /home/user/bridge_eth_addresses.json    \
    /home/user/validator_private_keys.txt   \
    /home/user/validator_addresses.txt ./
COPY --chown=user arb-validator/server.crt arb-validator/server.key ./

ENV ID=0 \
    WAIT_FOR="arb-bridge-eth:7545" \
    ETH_URL="ws://arb-bridge-eth:7545" \
    COORDINATOR_URL="" \
    AVM="cpp" \
    PATH="/home/user/go/bin:${PATH}"

# Build cache
COPY --chown=user --from=arb-validator-builder /home/user/.cache/go-build /build

# 1) Waits for host:port if $WAIT_FOR is set
# 2) Copies address files from ../ to ./ (state volume)
# 3) Launches follower if $COORDINATOR_URL else launches coordinator
CMD if [[ ! -z ${WAIT_FOR} ]]; then \
sleep 2 && while ! nc -z ${WAIT_FOR//:/ }; do sleep 2; done && sleep 2; \
echo "Finished waiting for ${WAIT_FOR}..."; else echo "Starting..."; fi \
&& cp bridge_eth_addresses.json validator_addresses.txt \
    server.* contract.ao ./state/ && touch ./state/contract.ao && \
sed -n $((${ID}+1))p validator_private_keys.txt > ./state/private_key.txt && \
T=follower; if [[ -z ${COORDINATOR_URL} ]]; then T=coordinator; fi; cd state &&\
${T}Server --avm=${AVM} contract.ao private_key.txt validator_addresses.txt \
    bridge_eth_addresses.json ${ETH_URL} ${COORDINATOR_URL}
EXPOSE 1235 1236
