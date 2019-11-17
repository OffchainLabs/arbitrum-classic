### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Ganache with the EthBridge deployed
### Exports bridge_eth_addresses.json and keys.json
### --------------------------------------------------------------------

# Global build args (for multistage build)
ARG MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"
ARG NUM_WALLETS=110
ARG NETWORK=dev.json

FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && apt-get update
RUN apt-get install -y curl sudo
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get install -y nodejs
RUN apt-get install -y git
RUN apt-get install -y netcat
RUN apt-get install -y psmisc
RUN curl https://get.parity.io -L | bash /dev/stdin -r stable
RUN useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g truffle@5.0.30 yarn@1.17.3
    # snap install parity && \
    # rm -rf /var/lib/apt/lists/* && \
    # useradd -ms /bin/bash user
COPY package.json ./
RUN yarn --production --frozen-lockfile --non-interactive
COPY contracts ./contracts
COPY migrations ./migrations
COPY test ./test
COPY truffle-config.js ./

RUN truffle compile

# # Global arguments
ARG MNEMONIC
ARG NUM_WALLETS
ARG NETWORK
ENV MNEMONIC=$MNEMONIC NUM_WALLETS=$NUM_WALLETS NETWORK=$NETWORK
# Generate bridge_eth_addresses.json for export
RUN echo arbitrum > password.txt
COPY parity ./parity
COPY --chown=user parity/config.toml /home/user/.local/share/io.parity.ethereum/config.toml
RUN parity --chain=parity/$NETWORK account import parity/keystore/
RUN parity --chain=parity/$NETWORK account list
RUN parity --chain parity/$NETWORK --unlock 0xe83f8ae25F873b1e17e05bda065ABEAc2FbD2E82 --password ~/password.txt & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for parity on localhost:8545..." && \
    truffle migrate --network parity -q && [ -f bridge_eth_addresses.json ]

FROM debian:buster-slim

ARG MNEMONIC
ARG NUM_WALLETS
ARG NETWORK
ENV MNEMONIC=$MNEMONIC NUM_WALLETS=$NUM_WALLETS NETWORK=$NETWORK

RUN useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
COPY --from=0 --chown=user /home/user/parity/ /home/user/parity/
COPY --from=0 --chown=user /usr/bin/parity /usr/bin/parity
COPY --from=0 --chown=user /home/user/.local/share/io.parity.ethereum/ /home/user/.local/share/io.parity.ethereum/
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json ./
RUN echo arbitrum > password.txt
ENTRYPOINT ["/usr/bin/parity", "--unlock", "0xe83f8ae25F873b1e17e05bda065ABEAc2FbD2E82", "--password", "/home/user/password.txt"]
# --chain parity/dev.json --jsonrpc-port=7545

# ENTRYPOINT ["/bin/bash"]

# Alpine dependencies and Non-root user
# Check dependencies
# RUN apk add --no-cache bash=5.0.11-r0 g++=9.2.0-r2 git=2.24.0-r0 make=4.2.1-r2 \
#     nodejs=12.13.0-r1 npm=12.13.0-r1 python2=2.7.16-r3 rhash=1.3.8-r0 \
#     libstdc++ eudev-libs libgcc && \
#     addgroup -g 1000 -S user && \
#     adduser -u 1000 -S user -G user -s /bin/bash -h /home/user
# USER user
# ENV PATH="/home/user/.npm-global/bin:${PATH}"
# WORKDIR "/home/user/"
# COPY --chown=user:user --from=parity/parity:stable /bin/parity /bin/parity
# RUN mkdir -p /home/user/.local/share/io.parity.ethereum/
# RUN mkdir -p /home/user/.npm-global && \
#     npm config set prefix "/home/user/.npm-global" && \
#     npm install -g truffle@5.0.30 yarn@1.17.3
# COPY package.json ./
# RUN yarn --production --frozen-lockfile --non-interactive

# # Source code
# COPY . ./
# RUN truffle compile




# # Minimize image
# FROM alpine:3.9

# RUN apk add --no-cache nodejs=10.14.2-r0

# # Non-root user
# RUN addgroup -g 1000 -S user && \
#     adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
# USER user
# WORKDIR "/home/user"

# # Addresses and keys
# COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json \
#     /home/user/keys.json ./

# # ganache-cli and truffle (placed in /bin and /lib) and build folder
# COPY --from=0 --chown=user /home/user/.npm-global /
# COPY --from=0 --chown=user /home/user/build /home/user/build
# COPY --from=0 --chown=user /home/user/ganachedb /home/user/ganachedb

# # Source files
# COPY --chown=user . ./

# # Global arguments
# ARG MNEMONIC
# ARG NUM_WALLETS

# ENV GAS_LIMIT=6721975 \
#     VERBOSE="-q" \
#     GAS_PER_WALLET=100 \
#     BLOCK_TIME=0 \
#     PORT=7545 \
#     MNEMONIC=$MNEMONIC \
#     NUM_WALLETS=$NUM_WALLETS \
#     DOCKER=true
# # DOCKER=true makes ganache run on host 0.0.0.0

# # Start ganache-cli using --db to use the EthBridge contract
# CMD ganache-cli --db ganachedb -p $PORT -l $GAS_LIMIT -e $GAS_PER_WALLET \
#     -b $BLOCK_TIME -a $NUM_WALLETS -m "${MNEMONIC}" $VERBOSE
# EXPOSE ${PORT}
