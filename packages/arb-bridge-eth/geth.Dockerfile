### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Geth with the EthBridge deployed
### Exports bridge_eth_addresses.json
### --------------------------------------------------------------------

FROM ethereum/client-go:stable
RUN apk add --no-cache \
    git \
    npm
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
COPY package.json .
RUN npm install --only=prod --ignore-scripts --no-package-lock
COPY deploy ./deploy
COPY contracts ./contracts
COPY buidler.config.ts .
COPY tsconfig.docker.json tsconfig.json
COPY geth ./geth
RUN echo arbitrum > password.txt && \
    geth --datadir data init geth/ethbridge.json && \
    cp geth/keystore/* data/keystore && \
    geth --datadir data --rpc --rpcaddr 'localhost' --rpcport 7545 --rpcapi 'personal,db,eth,net,web3,txpool,miner' --allow-insecure-unlock --unlock 0x81183c9c61bdf79db7330bbcda47be30c0a85064 --password ~/password.txt --mine & \
    while ! nc -z localhost 7545; do sleep 2; done; \
    echo "Finished waiting for geth on localhost:7545..." && \
    CI=true npx buidler deploy --network parity && [ -f bridge_eth_addresses.json ]

FROM ethereum/client-go:stable

RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
COPY --from=0 --chown=user /home/user/data/ /home/user/data/
COPY --from=0 --chown=user /home/user/bridge_eth_addresses.json ./
RUN echo arbitrum > password.txt
ENTRYPOINT ["/usr/local/bin/geth", "--datadir", "data", "--allow-insecure-unlock", "--unlock", "0x81183c9c61bdf79db7330bbcda47be30c0a85064", "--password", "/home/user/password.txt", "--mine"]
