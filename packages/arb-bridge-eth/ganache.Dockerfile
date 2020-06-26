### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Ganache with the EthBridge deployed
### Exports bridge_eth_addresses.json and keys.json
### --------------------------------------------------------------------

# Global build args (for multistage build)
ARG MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"
ARG NUM_WALLETS=15

FROM alpine:3.10

# Alpine dependencies and Non-root user
# Check dependencies
RUN apk add --no-cache g++ git make \
    nodejs npm python2 && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
ENV PATH="/home/user/.npm-global/bin:${PATH}"
WORKDIR "/home/user/"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g ganache-cli@6.5.0
COPY package.json ./
RUN npm install --only=prod --ignore-scripts --no-package-lock

# Source code
COPY deploy ./deploy
COPY contracts ./contracts
COPY buidler.config.ts .
COPY tsconfig.docker.json tsconfig.json
COPY parity ./parity

# Global arguments
ARG MNEMONIC
ARG NUM_WALLETS
ENV MNEMONIC=$MNEMONIC \
    NUM_WALLETS=$NUM_WALLETS

# Generate bridge_eth_addresses.json for export
RUN mkdir db && ganache-cli --db db -e 100000 \
        -p 7545 -a "${NUM_WALLETS}" -m "${MNEMONIC}" & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for ganache on localhost:${PORT}..." && \
    DOCKER=true npx buidler deploy --network parity && [ -f bridge_eth_addresses.json ]


# Minimize image
FROM alpine:3.10

RUN apk add --no-cache nodejs

# Non-root user
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR "/home/user"

# Addresses
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json ./

# ganache-cli and truffle (placed in /bin and /lib) and build folder
COPY --from=0 --chown=user /home/user/.npm-global /
COPY --from=0 --chown=user /home/user/build /home/user/build
COPY --from=0 --chown=user /home/user/db /home/user/db

# Source files

ARG MNEMONIC
ARG NUM_WALLETS
ENV MNEMONIC=$MNEMONIC \
    NUM_WALLETS=$NUM_WALLETS

# Start ganache-cli using --db to use the EthBridge contract
ENTRYPOINT ["/bin/ganache-cli", "-h", "0.0.0.0", "--db", "db", "-p", "7545"]
