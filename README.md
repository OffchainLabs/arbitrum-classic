# Arbitrum

### How to use:

Install depdencies using:
```
git checkout myym/ganache-deployment
git submodule update --init --recursive
yarn 
yarn build
```

Update `.env` file and perform deployment using:
```
yarn deploy // deploys a new pair of contracts on Layer 1 specified in .env file

yarn initialize // creates a new rollup on the chain specified in .env
```