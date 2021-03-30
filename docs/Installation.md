---
id: Installation
title: Installation
sidebar_label: Installation
---

## Setup instructions

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

### 1. Install python3 and docker:

#### MacOS

Using [Homebrew](https://brew.sh/):

```bash
brew install python3 docker docker-compose
brew cask install docker
open -a Docker
```

Once the Docker app appears in the menu bar, wait until the yellow light turns
green (no need to log into Docker).

#### Ubuntu 18.04

Using apt:

```bash
sudo apt update
sudo apt install -y curl python3 python3-pip
```

Then setup docker using the [official instructions](https://docs.docker.com/engine/install/ubuntu/)

Also setup docker compose using the [official instructions](https://docs.docker.com/compose/install/)

### 2. Install node, yarn and truffle

```bash
touch ~/.bashrc
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.35.3/install.sh | bash
. ~/.bashrc
nvm install 10.16.3

curl -o- -L https://yarnpkg.com/install.sh | bash
. ~/.bashrc

yarn global add truffle
```

### Full list

Here are the important dependencies in case you are not running on a supported OS:

- [docker](https://github.com/docker/docker-ce/releases) and
  [docker-compose](https://github.com/docker/compose/releases)
- [node](https://nodejs.org/en/)
- [python3 and pip3](https://www.python.org/downloads/)
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
- [yarn](https://yarnpkg.com/en/)

> Requires `node -v` version 8, 10 or 12

> Requires`python3 --version` 3.6 or greater

## Download Arbitrum

Download the Arbitrum Monorepo from source:

```bash
git clone -b v0.8.0 https://github.com/offchainlabs/arbitrum.git
cd arbitrum
git submodule update --init --recursive
yarn
yarn build
```

# Native setup instructions

For most users we recommend that you run Arbitrum through our easy docker setup. However if you want to install Arbitrum natively, additionally follow this instructions:

#### MacOS

```bash
brew install autoconf automake boost cmake gmp go libtool rocksdb
```

#### Ubuntu 18.04

```bash
sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y autoconf automake cmake libboost-dev libboost-filesystem-dev libgmp-dev librocksdb-dev libssl-dev libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev liblz4-dev libzstd-dev libtool golang-go clang-format cmake-format

git clone -b v6.11.4 https://github.com/facebook/rocksdb
cd rocksdb
make shared_lib
make install
```

## Install Arbitrum

With the dependencies installed, you can now install Arbitrum locally with

```bash
cd arbitrum
yarn install:validator
```
