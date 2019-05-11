---
id: start
title: Arbitrum Developer Quickstart
---

## Step 1: Install

### Dependencies

Python3, NodeJS, Truffle, VirtualBox, and Docker.

#### MacOS

```
brew install python3 node docker docker-machine virtualbox
npm install -g truffle
```

#### Ubuntu 18.04

```
sudo apt-get update
sudo apt-get install -y python3 python3-pip nodejs npm docker virtualbox
sudo npm install -g truffle
```

### Arbitrum Compiler: `arbc-truffle-compile`

``` bash
git clone --depth=1 https://github.com/OffchainLabs/arbc-solidity.git
cd arbc-solidity
# pip3 install virtualenv
# python3 -m venv venv
# source venv/bin/activate
pip3 install -r requirements.txt
sudo python3 setup.py install
# deactivate
```

## Step 2: Add truffle provider and arb-web3 providers to project

```
yarn add https://github.com/OffchainLabs/arb-truffle-provider.git
yarn add https://github.com/OffchainLabs/arb-web3-provider.git
```

## Step 3: Compile

```
truffle migrate --network arbitrum
arbc-truffle-compile compiled.json contract.ao
```

## Step 4: Launch 3 validators

```
```