---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
---

## Install

### Dependencies

The following dependencies are needed to get up and running with Arbitrum on Ubuntu or MacOS:

- [python3 and pip3](https://www.python.org/downloads/)
    - MacOS: `brew install python3`
    - Ubuntu 18.04: `sudo apt-get install -y python3 pip3`
- [node and npm](https://nodejs.org/en/)
    - MacOS: `brew install node`
    - Ubuntu 18.04: `sudo apt-get install -y nodejs npm`
- [virtualbox](https://www.virtualbox.org/wiki/Downloads)
    - MacOS: [download](https://www.virtualbox.org/wiki/Downloads)
    - Ubuntu 18.04: `sudo apt-get install virtualbox`
- [docker](https://github.com/docker/docker-ce/releases) and [docker-compose](https://github.com/docker/compose/releases)
    - MacOS: `brew install docker docker-machine docker-compose`
    - Ubuntu 18.04: `sudo apt-get install -y docker docker-compose`
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation) with the following command:
    - `npm install -g truffle`

### Arbitrum Compiler

Install the Arbitrum compiler, `arbc-truffle-compile`, by building it from source:

``` bash
git clone --depth=1 https://github.com/OffchainLabs/arbc-solidity.git
cd arbc-solidity
pip3 install virtualenv
python3 -m venv venv
source venv/bin/activate
pip3 install -r requirements.txt
deactivate
sudo python3 setup.py install
```

## Run the Demo App

``` bash
git clone --depth=1 https://github.com/OffchainLabs/demo-app.git
cd demo-app
./arb.py
```

And from another prompt in the same directory run:
```
yarn start
```

## Porting a solidity project

Note this section is TODO. Stop reading here.

### Add truffle provider and arb-web3 providers to project

```
yarn add https://github.com/OffchainLabs/arb-truffle-provider.git
yarn add https://github.com/OffchainLabs/arb-web3-provider.git
```

### Compile

```
truffle migrate --network arbitrum
arbc-truffle-compile compiled.json contract.ao
sudo docker build -t arb-app -f arb-app.Dockerfile .
sudo docker-compose build
sudo docker
```

## Step 4: Launch 3 validators

```
```
