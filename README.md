# Arbitrum

### How to use:

Initialize submodules and install required depenencies by running:
```
chmod +x ./scripts/preq.sh && ./scripts/preq.sh
```
To install the validator node run: `yarn install:validator`


Update `.env` file and perform deployment using:
```
yarn deploy // deploys a new pair of contracts on Layer 1 specified in .env file

yarn initialize // creates a new rollup on the chain specified in .env

yarn run:rpc // to run the rpc node

yarn run:validator // to run the validator node
```

#### FAQ:

- If getting storage initialization erorr when runing rpc/validator node copy arbose.mexe to arbitrum directory by:
`mv /home/ubuntu/arbitrum/rollups/local_ganache/validator0/arbos.mexe ./home/ubuntu/arbitrum`