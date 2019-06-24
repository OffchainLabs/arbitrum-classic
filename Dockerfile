### --------------------------------------------------------------------
### Dockerfile
### arb-validator
### Note: depends on mounting `/home/user/contract.ao` as a volume
### When arb-deploy runs in --dev-mode a separate dev-mode.Dockerfile is
### created with the commented DEV_COPYs changed to COPY
### --------------------------------------------------------------------

FROM alpine:3.9 as dev

# Alpine dependencies
RUN apk add --no-cache git go libc-dev

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user/"

# Dependencies
COPY --chown=user go.mod go.sum /home/user/
##DEV_COPY --chown=user vendor ./vendor
RUN if [[ -d vendor ]]; then \
    rm -rf ./vendor/github.com/offchainlabs/arb-avm && \
    go build -v ./vendor/... ; fi

# arb-avm
##DEV_COPY --chown=user arb-avm /home/user/arb-avm
RUN if [[ -d arb-avm ]]; then cd arb-avm && go build -v ./... ; fi

# arb-validator
# if dev-mode build quickly else make the smallest binary possible
COPY --chown=user ./ /home/user/
RUN if [[ -d arb-avm ]]; then \
    echo "replace github.com/offchainlabs/arb-avm => ./arb-avm" >> go.mod && \
    go build -v ./cmd/followerServer ./cmd/coordinatorServer && \
    go install ./cmd/followerServer ./cmd/coordinatorServer && \
    go clean -cache -modcache -r ./cmd/followerServer ./cmd/coordinatorServer; \
    else \
    export GOOS=linux GOARCH=amd64 && \
    cd cmd/followerServer && go build -a -v -ldflags "-w -s" && \
    cd ../coordinatorServer && go build -a -v -ldflags "-w -s" && \
    rm -rf /home/user/go && mkdir -p /home/user/go/bin && \
    mv coordinatorServer ../followerServer/followerServer /home/user/go/bin/; fi


# Minimize
FROM alpine:3.9
# Non root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
# Note: state will be mounted as a volume and initially overwritten
RUN mkdir -p /home/user/state
WORKDIR "/home/user/"
COPY --chown=user --from=0 /home/user/go /home/user/go

# Get EthBridge addresses and Validator private keys and addresses
COPY --chown=user --from=arb-ethbridge      \
    /home/user/ethbridge_addresses.json     \
    /home/user/validator_private_keys.txt   \
    /home/user/validator_addresses.txt ./
COPY --chown=user server.crt server.key ./

# Arguments
ARG WAIT_FOR
ARG ID
ARG ETH_URL
ARG COORDINATOR_URL
ENV WAIT_FOR=$WAIT_FOR ID=$ID ETH_URL=$ETH_URL \
    COORDINATOR_URL=$COORDINATOR_URL \
    PATH="/home/user/go/bin:${PATH}"

# 1) Waits for host:port if $WAIT_FOR is set
# 2) Copies address files from ../ to ./ (state volume)
# 3) Launches follower if $COORDINATOR_URL else launches coordinator
CMD if [[ ! -z ${WAIT_FOR} ]]; then \
sleep 2 && while ! nc -z ${WAIT_FOR//:/ }; do sleep 2; done && sleep 2; \
echo "Finished waiting for ${WAIT_FOR}..."; else echo "Starting..."; fi \
&& cp ethbridge_addresses.json validator_addresses.txt \
    server.* contract.ao ./state/ && touch ./state/contract.ao && \
sed -n $((${ID}+1))p validator_private_keys.txt > ./state/private_key.txt && \
T=follower; if [[ -z ${COORDINATOR_URL} ]]; then T=coordinator; fi; cd state &&\
${T}Server contract.ao private_key.txt validator_addresses.txt \
    ethbridge_addresses.json ${ETH_URL} ${COORDINATOR_URL}
EXPOSE 1235 1236
