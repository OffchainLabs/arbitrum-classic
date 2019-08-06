### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: `##DEV_` commands are run in dev-mode. They are not comments
### Note: run depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM alpine:3.9

# Alpine dependencies
RUN apk add --no-cache build-base git go libc-dev linux-headers

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"

# Dependencies
COPY --chown=user go.mod go.sum /home/user/
##DEV_COPY --chown=user arb-avm-go/go.mod arb-avm-go/go.sum /home/user/arb-avm-go/
##DEV_COPY --chown=user arb-util/go.mod arb-util/go.sum /home/user/arb-util/
##DEV_COPY --chown=user arb-avm-cpp/go.mod arb-avm-cpp/go.sum /home/user/arb-avm-cpp/
RUN if [[ -d arb-avm-go ]]; then                                                        \
    echo "replace github.com/offchainlabs/arb-avm-go => ./arb-avm-go" >> go.mod &&      \
    echo "replace github.com/offchainlabs/arb-util => ./arb-util" >> go.mod &&          \
    echo "replace github.com/offchainlabs/arb-avm-cpp => ./arb-avm-cpp" >> go.mod &&    \
    echo "replace github.com/offchainlabs/arb-util => ../arb-util" >> arb-avm-go/go.mod;fi;\
    go mod download

##DEV_COPY --chown=user arb-util /home/user/arb-util
##DEV_COPY --chown=user arb-avm-go /home/user/arb-avm-go
COPY --from=arb-avm-cpp --chown=user /arb-avm-cpp /home/user/arb-avm-cpp/
COPY --chown=user ./ /home/user/

# Build cache
##DEV_COPY --from=arb-validator --chown=user /build /home/user/.cache/go-build

# Build arb-validator
RUN if [[ -d arb-avm-go ]]; then \
    echo "replace github.com/offchainlabs/arb-avm-go => ./arb-avm-go" >> go.mod &&      \
    echo "replace github.com/offchainlabs/arb-util => ./arb-util" >> go.mod &&          \
    echo "replace github.com/offchainlabs/arb-avm-cpp => ./arb-avm-cpp" >> go.mod &&    \
    echo "replace github.com/offchainlabs/arb-util => ../arb-util" >> arb-avm-go/go.mod;fi;\
    go build -v ./cmd/followerServer ./cmd/coordinatorServer && \
    go install ./cmd/followerServer ./cmd/coordinatorServer


# Minimize
FROM alpine:3.9

# Alpine dependencies
RUN apk add --no-cache libstdc++ libgcc

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"

# Compiled code
COPY --chown=user --from=0 /home/user/go/bin /home/user/go/bin

# Get EthBridge addresses and Validator private keys and addresses
COPY --chown=user --from=arb-bridge-eth     \
    /home/user/bridge_eth_addresses.json    \
    /home/user/validator_private_keys.txt   \
    /home/user/validator_addresses.txt ./
COPY --chown=user server.crt server.key ./

ENV ID=0 \
    WAIT_FOR="arb-bridge-eth:7545" \
    ETH_URL="ws://arb-bridge-eth:7545" \
    COORDINATOR_URL="" \
    AVM="cpp" \
    PATH="/home/user/go/bin:${PATH}"

# Build cache
COPY --chown=user --from=0 /home/user/.cache/go-build /build

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
