### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: depends on mounting `/home/user/contract.ao` as a volume
### --------------------------------------------------------------------

FROM golang:1.12-alpine

# Alpine dependencies
RUN apk add --no-cache gcc git libc-dev

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash
USER user
RUN mkdir -p /home/user/arb-validator/arb-avm
WORKDIR "/home/user/arb-validator"

# Dependencies
COPY --chown=user arb-avm/go.mod arb-avm/go.sum /home/user/arb-avm/
COPY --chown=user go.mod go.sum /home/user/arb-validator/
RUN go mod download

# Source code
COPY --chown=user ./ /home/user/arb-validator
RUN mv arb-avm/* ../arb-avm && rm -rf arb-avm && \
    export GOOS=linux GOARCH=amd64 && \
    cd cmd/followerServer && go build -a -v -ldflags "-w -s" && \
    cd ../coordinatorServer && go build -a -v -ldflags "-w -s" && \
    mv coordinatorServer ../followerServer/followerServer /go/bin/

# Minimize image
FROM alpine:3.9

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/state"

# coordinatorServer, followerServer, server.key, server.crt
COPY --chown=user --from=0 /go/bin/*Server /usr/bin/
COPY --chown=user --from=0 /home/user/arb-validator/server* ../
# Get EthBridge addresses and Validator private keys and addresses
COPY --chown=user --from=arb-ethbridge      \
    /home/node/ethbridge_addresses.json     \
    /home/node/validator_private_keys.txt   \
    /home/node/validator_addresses.txt ../

# Arguments
ARG WAIT_FOR
ARG ID
ARG ETH_URL
ARG COORDINATOR_URL
ENV WAIT_FOR=$WAIT_FOR ID=$ID ETH_URL=$ETH_URL \
    COORDINATOR_URL=$COORDINATOR_URL

# 1) Waits for host:port if $WAIT_FOR is set
# 2) Copies address files from ../ to ./ (state volume)
# 3) Launches follower if $COORDINATOR_URL else launches coordinator
CMD if [[ ! -z ${WAIT_FOR} ]]; then \
sleep 2 && while ! nc -z ${WAIT_FOR//:/ }; do sleep 2; done && sleep 2;\
echo "Finished waiting for ${WAIT_FOR}..."; else echo "Starting..."; fi\
&& cp ../ethbridge_addresses.json ../validator_addresses.txt \
    ../server.* ../contract.ao ./ && touch contract.ao && \
sed -n $((${ID}+1))p ../validator_private_keys.txt > private_key.txt &&\
T=follower; if [[ -z ${COORDINATOR_URL} ]]; then T=coordinator; fi; \
${T}Server contract.ao private_key.txt validator_addresses.txt \
    ethbridge_addresses.json ${ETH_URL} ${COORDINATOR_URL}
EXPOSE 1235 1236
