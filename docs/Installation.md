---
id: Installation
title: Installation
sidebar_label: Installation
---

## Install System Dependencies

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

### 1. Install python3 and docker:

#### MacOS

Using [Homebrew](https://brew.sh/):

```bash
brew install python3 docker docker-compose rocksdb
brew cask install docker
open -a Docker
```

Once the Docker app appears in the menu bar, wait until the yellow light turns
green (no need to log into Docker). Also check that node version 10 is installed
correctly by running `node -v`.

#### Ubuntu 18.04

Using apt:

```bash
sudo apt update
sudo apt install -y python3 python3-pip docker docker-compose
```

> Docker [can be used without sudo](https://docs.docker.com/install/linux/linux-postinstall/)
> to give permissions "equivalent to the `root` user". See [the security warning](https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface).

### 2. Install yarn and truffle

```bash
touch ~/.bashrc
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.34.0/install.sh | bash
curl -o- -L https://yarnpkg.com/install.sh | bash
nvm install 10.16.3
. ~/.bashrc
yarn global add truffle
```

### Full List

Here are the important dependencies in case you are not running on a supported OS:

-   [docker](https://github.com/docker/docker-ce/releases) and
    [docker-compose](https://github.com/docker/compose/releases)
-   [node](https://nodejs.org/en/)
-   [python3 and pip3](https://www.python.org/downloads/)
-   [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
-   [yarn](https://yarnpkg.com/en/)

> Requires `node -v` version 8, 10 or 12

> Requires`python3 --version` 3.6 or greater

## Install Arbitrum

Download the Arbitrum Monorepo from source:

```bash
git clone -b v0.3.0 --depth=1 -c advice.detachedHead=false https://github.com/offchainlabs/arbitrum.git
cd arbitrum
yarn
yarn build
yarn install:deps
```

Check `arbc-truffle` was installed:

```bash
which arbc-truffle
```

Expected output:

> /usr/local/bin/arbc-truffle
