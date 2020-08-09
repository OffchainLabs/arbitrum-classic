### --------------------------------------------------------------------
### Dockerfile
### arb-bridge-eth
### Runs Geth with the EthBridge deployed
### Exports bridge_eth_addresses.json
### --------------------------------------------------------------------

FROM ethereum/client-go:stable
RUN addgroup -g 1000 -S user && \
    adduser -u 1000 -S user -G user -s /bin/ash -h /home/user
USER user
WORKDIR /home/user/
COPY geth ./geth
RUN echo arbitrum > password.txt && \
    mkdir -p data/keystore && \
    cp geth/keystore/* data/keystore && \
    geth --datadir data init geth/ethbridge.json
ENTRYPOINT ["/usr/local/bin/geth", "--syncmode", "full", "--datadir", "data", "--allow-insecure-unlock", "--unlock", "0x81183c9c61bdf79db7330bbcda47be30c0a85064", "--password", "/home/user/password.txt", "--mine"]
