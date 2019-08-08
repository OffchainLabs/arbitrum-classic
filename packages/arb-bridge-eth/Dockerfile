### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Ganache with the EthBridge deployed
### Exports bridge_eth_addresses.json and keys.json
### --------------------------------------------------------------------

# Global build args (for multistage build)
ARG MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"
ARG NUM_VALIDATORS=10
ARG NUM_WALLETS=110

FROM alpine:3.9

# Alpine dependencies and Non-root user
# Check dependencies
RUN apk add --no-cache g++=8.3.0-r0 git=2.20.1-r0 make=4.2.1-r2 \
    nodejs=10.14.2-r0 npm=10.14.2-r0 python2=2.7.16-r1 && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
ENV PATH="/home/user/.npm-global/bin:${PATH}"
WORKDIR "/home/user/"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g ganache-cli@6.5.0 truffle@5.0.30 yarn@1.17.3
COPY package.json ./
RUN yarn --production --frozen-lockfile --non-interactive

# Source code
COPY . ./
RUN truffle compile

# Global arguments
ARG MNEMONIC
ARG NUM_VALIDATORS
ARG NUM_WALLETS
ENV DOCKER=true MNEMONIC=$MNEMONIC NUM_VALIDATORS=$NUM_VALIDATORS \
    NUM_WALLETS=$NUM_WALLETS

# Generate bridge_eth_addresses.json for export
RUN PORT=$(awk '/port: / {print $2}' truffle-config.js | sed 's/,//g');\
    mkdir db && ganache-cli --db db --acctKeys keys.json \
        -p "${PORT}" -a "${NUM_WALLETS}" -m "${MNEMONIC}" & \
    while ! nc -z localhost ${PORT}; do sleep 2; done; \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    truffle migrate --reset --compile-all -q && [ -f bridge_eth_addresses.json ]


# Minimize image
FROM alpine:3.9

RUN apk add --no-cache nodejs=10.14.2-r0

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"

# Addresses and keys
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json \
    /home/user/keys.json ./

# ganache-cli and truffle (placed in /bin and /lib) and build folder
COPY --from=0 --chown=user /home/user/.npm-global /
COPY --from=0 --chown=user /home/user/build /home/user/build
COPY --from=0 --chown=user /home/user/db /home/user/db

# Source files
COPY --chown=user . ./

# Global arguments
ARG MNEMONIC
ARG NUM_WALLETS
ARG NUM_VALIDATORS

ENV GAS_LIMIT=6721975 \
    VERBOSE="-q" \
    GAS_PER_WALLET=100 \
    BLOCK_TIME=0 \
    PORT=7545 \
    MNEMONIC=$MNEMONIC \
    NUM_WALLETS=$NUM_WALLETS \
    NUM_VALIDATORS=$NUM_VALIDATORS \
    DOCKER=true
# DOCKER=true makes ganache run on host 0.0.0.0

# Start ganache-cli using --db to use the EthBridge contract
CMD ganache-cli --db db -p $PORT -l $GAS_LIMIT -e $GAS_PER_WALLET \
    -b $BLOCK_TIME -a $NUM_WALLETS -m "${MNEMONIC}" $VERBOSE
EXPOSE ${PORT}
