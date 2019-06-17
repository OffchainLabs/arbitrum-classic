### ---------------------------------------------------------------------
### Dockerfile
### arb-ethbridge
### ---------------------------------------------------------------------

FROM node:carbon-alpine

# Alpine dependencies
RUN apk add --no-cache g++ git make python

# Non-root user
USER node
ENV HOME="/home/node" PATH="/home/node/.npm-global/bin:${PATH}"
WORKDIR "${HOME}"
RUN mkdir "${HOME}/.npm-global" && \
    npm config set prefix "${HOME}/.npm-global" && \
    npm install -g ganache-cli truffle yarn bip39 ethereumjs-wallet
COPY package.json yarn.lock ./
RUN yarn --prod --frozen-lockfile && yarn cache clean

# Source code
COPY . ./
RUN truffle compile

# Allow host 0.0.0.0
ENV DOCKER=true

# Environment argument defaults
ARG MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"
ARG NUM_VALIDATORS=10
ARG NUM_WALLETS=10
ARG GAS_PER_WALLET=100
ARG PORT=7545
ARG CANARY_PORT=17545
ENV MNEMONIC=$MNEMONIC NUM_VALIDATORS=$NUM_VALIDATORS NUM_WALLETS=$NUM_WALLETS \
    GAS_PER_WALLET=$GAS_PER_WALLET PORT=$PORT CANARY_PORT=$CANARY_PORT \
    GANACHE_FLAGS="0 -e ${GAS_PER_WALLET} -a ${NUM_WALLETS} -p ${PORT}"

# Generate ethbrigde_addresses.json for export
RUN ganache-cli ${GANACHE_FLAGS} -m "${MNEMONIC}" & \
    while ! nc -z localhost ${PORT}; do sleep 2; done; \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    truffle migrate --reset && [ -f ethbridge_addresses.json ]

# Generate validator_addresses.txt and mnemonic_private_keys.txt
RUN node -e "                                               \
    const bip39 = require('bip39');                         \
    const hdkey = require('ethereumjs-wallet/hdkey');       \
    var m = '${MNEMONIC}';                                  \
    var w = hdkey.fromMasterSeed(bip39.mnemonicToSeed(m));  \
    var addrs = new Array();                                \
    var privs = new Array();                                \
    var path = \"m/44'/60'/0'/0/\";                         \
    for (var i = 100; i < 100 + ${NUM_VALIDATORS}; i++) {   \
        var acc = w.derivePath(path + i).getWallet();       \
        addrs.push(acc.getAddress().toString('hex'));       \
        privs.push(acc.getPrivateKey().toString('hex'));    \
    }                                                       \
    console.log(addrs.join('\n'));                          \
    console.error(privs.join('\n'));                        \
    " > validator_addresses.txt 2> validator_private_keys.txt

# Auxillary flags (do not require re-build to change)
ARG GAS_LIMIT=6721975
ARG VERBOSE="-q"
ENV GAS_LIMIT=$GAS_LIMIT VERBOSE=$VERBOSE \
    GANACHE_AUX_FLAGS="-l ${GAS_LIMIT} ${VERBOSE}"

# Run
CMD (while ! nc -z localhost ${PORT}; do sleep 2; done && \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    truffle migrate --reset && nc -lvp ${CANARY_PORT} -w 30) & \
    ganache-cli ${GANACHE_AUX_FLAGS} ${GANACHE_FLAGS} -m "${MNEMONIC}"
EXPOSE ${PORT}
