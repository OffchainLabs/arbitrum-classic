### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Geth with the EthBridge deployed
### Exports bridge_eth_addresses.json
### --------------------------------------------------------------------
ARG NETWORK=poa.json

FROM parity/parity:latest
USER root
RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -y \
    ca-certificates && \
    curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
    apt-get install -y \
    git \
    netcat \
    nodejs
USER parity
WORKDIR /home/parity/
COPY package.json .
ENV PATH="/home/parity/.local/bin:/home/parity/.npm-global/bin:${PATH}"
ENTRYPOINT []
RUN mkdir -p /home/parity/.npm-global && \
    npm config set prefix "/home/parity/.npm-global" && \
    npm install --only=prod --ignore-scripts --no-package-lock
COPY deploy ./deploy
COPY contracts ./contracts
COPY buidler.config.ts .
COPY tsconfig.docker.json tsconfig.json
COPY parity ./parity
COPY --chown=parity parity/config.toml /home/parity/.local/share/io.parity.ethereum/config.toml
ARG NETWORK
ENV NETWORK=$NETWORK
RUN echo arbitrum > password.txt && \
    parity --chain=parity/$NETWORK account import parity/keystore/ && \
    parity --chain parity/$NETWORK --unlock 0x81183c9c61bdf79db7330bbcda47be30c0a85064 --password ~/password.txt & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for parity on localhost:7545..." && \
    DOCKER=true npx buidler deploy --network parity && [ -f bridge_eth_addresses.json ]

FROM parity/parity:latest

ARG NETWORK
ENV NETWORK=$NETWORK

COPY --from=0 --chown=parity /home/parity/parity/ /home/parity/parity/
COPY --from=0 --chown=parity /home/parity/.local/share/io.parity.ethereum/ /home/parity/.local/share/io.parity.ethereum/
COPY --from=0 --chown=parity /home/parity/bridge_eth_addresses.json ./
RUN echo arbitrum > password.txt
ENTRYPOINT ["/usr/bin/parity", "--unlock", "0x81183c9c61bdf79db7330bbcda47be30c0a85064", "--password", "/home/parity/password.txt"]
