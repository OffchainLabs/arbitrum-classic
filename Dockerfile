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
RUN apk add --no-cache g++ git make python nodejs npm && \
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
    ganache-cli -p "${PORT}" -a "${NUM_WALLETS}" -m "${MNEMONIC}" & \
    while ! nc -z localhost ${PORT}; do sleep 2; done; \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    truffle migrate --reset && [ -f ethbridge_addresses.json ]


# Calculate validator_addresses.txt and validator_private_keys.txt
FROM alpine:3.9 as keys

# Alpine dependencies and Non-root user
RUN apk add --no-cache g++ git make python nodejs npm && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"
RUN echo "{}" > package.json && npm install ethers

# Global arguments
ARG MNEMONIC
ARG NUM_VALIDATORS
ARG NUM_WALLETS
ENV MNEMONIC=$MNEMONIC NUM_VALIDATORS=$NUM_VALIDATORS \
    NUM_WALLETS=$NUM_WALLETS

# Generates validator_addresses.txt and validator_private_keys.txt
RUN node -e "                                               \
    const ethers = require('ethers');                       \
    let m = '${MNEMONIC}';                                  \
    let addrs = new Array();                                \
    let privs = new Array();                                \
    let p = \"m/44'/60'/0'/0/\";                            \
    let base = ${NUM_WALLETS} - ${NUM_VALIDATORS};          \
    for (let i = base; i < base + ${NUM_VALIDATORS}; i++) { \
        let a = ethers.Wallet.fromMnemonic(m, path=(p+i));  \
        addrs.push(a.address);                              \
        privs.push(a.privateKey);                           \
    }                                                       \
    console.log(addrs.join('\n'));                          \
    console.error(privs.join('\n'));                        \
    " > validator_addresses.txt 2> validator_private_keys.txt


# Minimize image
FROM alpine:3.9

RUN apk add --no-cache nodejs

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"

# Source files
COPY --chown=user . ./
# Build files
COPY --chown=user --from=0 /home/user/build /home/user/build
# addresses
COPY --chown=user --from=0 /home/user/ethbridge_addresses.json ./
COPY --chown=user --from=keys /home/user/validator_*.txt ./
# ganache-cli and truffle (placed in /bin and /lib)
COPY --chown=user --from=0 /home/user/.npm-global /

# Global arguments
ARG MNEMONIC
ARG NUM_VALIDATORS
ARG NUM_WALLETS
# Final arguments
ARG GAS_LIMIT=6721975
ARG VERBOSE="-q"
ARG GAS_PER_WALLET=100
ARG PORT=7545
ARG CANARY_PORT=17545
# DOCKER=true makes ganache run on host 0.0.0.0
ENV DOCKER=true MNEMONIC=$MNEMONIC NUM_VALIDATORS=$NUM_VALIDATORS \
    NUM_WALLETS=$NUM_WALLETS GL=$GAS_LIMIT V=$VERBOSE \
    GPW=$GAS_PER_WALLET P=$PORT CP=$CANARY_PORT

# Wait for ganache-cli to launch and then deploy the EthBridge contract
CMD sed -i "s/port: [0-9]*,/port: ${P},/p" truffle-config.js && \
    (while ! nc -z localhost ${P}; do sleep 2; done &&           \
    echo "Finished waiting for ganache on localhost:${P}..." &&  \
    truffle migrate && nc -lvp ${CP} -w 360) & \
    ganache-cli -p $P -l $GL -e $GPW -a $NUM_WALLETS -m "${MNEMONIC}" $V
EXPOSE ${P}
