### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Ganache with the EthBridge deployed
### Exports bridge_eth_addresses.json and keys.json
### --------------------------------------------------------------------

# Global build args (for multistage build)
ARG NETWORK=poa.json

FROM debian:buster-slim

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y \
    curl \
    sudo && \
    curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
    apt-get install -y \
    git \
    netcat \
    nodejs \
    psmisc &&\
    rm -rf /var/lib/apt/lists/* && \
    curl https://get.parity.io -L | bash /dev/stdin -r stable && \
    useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g truffle@5.0.30 yarn@1.17.3
COPY package.json ./
RUN yarn --production --frozen-lockfile --non-interactive
COPY contracts ./contracts
COPY migrations ./migrations
COPY test ./test
COPY installed_contracts ./installed_contracts
COPY truffle-config.js ./

RUN truffle compile

# # Global arguments
ARG NETWORK
ENV NETWORK=$NETWORK
# Generate bridge_eth_addresses.json for export
COPY parity ./parity
COPY --chown=user parity/config.toml /home/user/.local/share/io.parity.ethereum/config.toml
RUN echo arbitrum > password.txt && \
    parity --chain=parity/$NETWORK account import parity/keystore/ && \
    parity --chain parity/$NETWORK --unlock 0xe83f8ae25F873b1e17e05bda065ABEAc2FbD2E82 --password ~/password.txt & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for parity on localhost:8545..." && \
    truffle migrate --network parity -q && [ -f bridge_eth_addresses.json ]

FROM debian:buster-slim

ARG NETWORK
ENV NETWORK=$NETWORK

RUN useradd -ms /bin/bash user
USER user
WORKDIR /home/user/
COPY --from=0 --chown=user /home/user/parity/ /home/user/parity/
COPY --from=0 --chown=user /usr/bin/parity /usr/bin/parity
COPY --from=0 --chown=user /home/user/.local/share/io.parity.ethereum/ /home/user/.local/share/io.parity.ethereum/
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json ./
RUN echo arbitrum > password.txt
ENTRYPOINT ["/usr/bin/parity", "--unlock", "0xe83f8ae25F873b1e17e05bda065ABEAc2FbD2E82", "--password", "/home/user/password.txt"]
