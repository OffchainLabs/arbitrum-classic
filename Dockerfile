### --------------------------------------------------------------------
### Dockerfile
### arb-ethbridge
### --------------------------------------------------------------------

# Global build args (for multistage build)
ARG MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"
ARG NUM_VALIDATORS=10
ARG NUM_WALLETS=110

FROM alpine:3.9

# Alpine dependencies and Non-root user
# Check dependencies
RUN apk add --no-cache g++ git make nodejs npm python && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
ENV PATH="/home/user/.npm-global/bin:${PATH}"
WORKDIR "/home/user/"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g ganache-cli truffle yarn
COPY package.json yarn.lock ./
RUN yarn --production --frozen-lockfile && yarn cache clean && \
    npm uninstall -g yarn && npm cache clean --force

# Source code
COPY . ./
RUN truffle compile

# Global arguments
ARG MNEMONIC
ARG NUM_VALIDATORS
ARG NUM_WALLETS
ENV DOCKER=true MNEMONIC=$MNEMONIC NUM_VALIDATORS=$NUM_VALIDATORS \
    NUM_WALLETS=$NUM_WALLETS

# Generate ethbrigde_addresses.json for export
RUN PORT=$(awk '/port: / {print $2}' truffle-config.js | sed 's/,//g');\
    mkdir db && ganache-cli --db db --acctKeys keys.json \
        -p "${PORT}" -a "${NUM_WALLETS}" -m "${MNEMONIC}" & \
    while ! nc -z localhost ${PORT}; do sleep 2; done; \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    truffle migrate --reset && [ -f ethbridge_addresses.json ] && \
    node -e "                                               \
    const data = require('./keys.json')['addresses'];       \
    const addresses = new Array(0);                         \
    const privates = new Array(0);                          \
    let start = ${NUM_WALLETS} - ${NUM_VALIDATORS};         \
    for (const address of Object.keys(data).slice(start)) { \
        addresses.push(address.slice(2));                   \
        let pk = data[address]['secretKey']['data'].reduce( \
        (a, B) => a+(B).toString(16).padStart(2, '0'), ''); \
        privates.push(pk);                                  \
    }                                                       \
    console.log(addresses.join('\n'));                      \
    console.error(privates.join('\n'));                     \
    " > validator_addresses.txt 2> validator_private_keys.txt


# Minimize image
FROM alpine:3.9

RUN apk add --no-cache nodejs

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"

# Addresses and keys
COPY --from=0 --chown=user /home/user/ethbridge_addresses.json \
    /home/user/validator_*.txt ./

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

# Wait for ganache-cli to launch and then deploy the EthBridge contract
CMD sed -i "s/port: [0-9]*,/port: ${PORT},/p" truffle-config.js && \
    ganache-cli --db db -p $PORT -l $GAS_LIMIT -e $GAS_PER_WALLET \
    -b $BLOCK_TIME -a $NUM_WALLETS -m "${MNEMONIC}" $VERBOSE
EXPOSE ${PORT}
