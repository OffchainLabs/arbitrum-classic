### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Geth with the EthBridge deployed
### Exports bridge_eth_addresses.json
### --------------------------------------------------------------------

FROM ethereum/client-go:stable

RUN apk add --no-cache nodejs npm git && \
    addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
ENV PATH="/home/user/.local/bin:/home/user/.npm-global/bin:${PATH}"
RUN mkdir -p /home/user/.npm-global && \
    npm config set prefix "/home/user/.npm-global" && \
    npm install -g truffle@5.0.30 yarn@1.17.3
COPY package.json ./
RUN yarn --frozen-lockfile --non-interactive
COPY contracts ./contracts
COPY migrations ./migrations
COPY test ./test
COPY installed_contracts ./installed_contracts
COPY truffle-config.js ./
RUN truffle compile

COPY geth ./geth

RUN echo arbitrum > password.txt && \
    geth --datadir data init geth/ethbridge.json && \
    cp geth/keystore/* data/keystore && \
    geth --datadir data --rpc --rpcaddr 'localhost' --rpcport 7545 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --allow-insecure-unlock --unlock 0x81183c9c61bdf79db7330bbcda47be30c0a85064 --password ~/password.txt --mine & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for geth on localhost:7545..." && \
    truffle migrate --network parity -q && [ -f bridge_eth_addresses.json ]

FROM ethereum/client-go:stable

RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
COPY --from=0 --chown=user /home/user/data/ /home/user/data/
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json ./
RUN echo arbitrum > password.txt
ENTRYPOINT ["/usr/local/bin/geth", "--datadir", "data", "--allow-insecure-unlock", "--unlock", "0x81183c9c61bdf79db7330bbcda47be30c0a85064", "--password", "/home/user/password.txt", "--mine"]
