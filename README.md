# Arbitrum

### How to use:

Install depdencies using:
```
git checkout myym/ganache-deployment
git submodule update --init --recursive
yarn 
yarn build
yarn install:validator
```

Update `.env` file and perform deployment using:
```
yarn deploy // deploys a new pair of contracts on Layer 1 specified in .env file

yarn initialize // creates a new rollup on the chain specified in .env

yarn run:rpc // to run the rpc node

yarn run:validator // to run the validator node
```