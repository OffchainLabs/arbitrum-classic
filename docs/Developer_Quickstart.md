---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Developer_Quickstart.md
---

Get started with Arbitrum by installing the Arbitrum compiler,
`arbc-truffle`, and its dependencies. Next,
[build and run the demo app](#hello-arbitrum) or
[port your own dapp](#porting-to-arbitrum).

**Want to learn more? Join the team on [Discord](https://discord.gg/ZpZuw7p) and
read the [white paper](https://offchainlabs.com/arbitrum.pdf)!**

## Install System Dependencies

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

### 1. Install python3 and docker:

#### MacOS

Using [Homebrew](https://brew.sh/):

```bash
brew install python3 docker docker-compose parity rocksdb
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
bash <(curl https://get.parity.io -L)
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
-   [parity](https://www.parity.io/ethereum)
-   [yarn](https://yarnpkg.com/en/)

> Requires `node -v` version 8, 10 or 12

> Requires`python3 --version` 3.6 or greater

## Install Arbitrum

Download the Arbitrum Monorepo from source:

```bash
git clone -b v0.2.1 --depth=1 -c advice.detachedHead=false https://github.com/offchainlabs/arbitrum.git
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

## Setup Blockchain

In the current alpha, Arbitrum is setup to run against a local test blockchain rather than a public blockchain.

To build a docker image hosting a the local test blockchain docker image with Arbitrum smart contracts already deployed, run:

```bash
yarn docker:build:geth
```

To start the local blockchain inside the Arbitrum monorepo, run:

```bash
yarn docker:geth
```

The local test blockchain should be running for all steps inside this tutorial. Note that
stopping and restarting the client will lose all blockchain state.

## Hello, Arbitrum

Now you'll compile and run a demo dApp on Arbitrum. The dApp is based on
a simple Pet Shop dApp that is used in a Truffle tutorial.

Inside the Arbitrum monorepo, open the `pet-shop` demo dApp:

```bash
cd demos/pet-shop
```

### Build and Run

You'll need to do these steps every time you make a change to the Solidity. For
this dApp, you do not need to change any Solidity files.

1. Compile Solidity to Arbitrum:

    Truffle will output the compiled contract as `contract.ao` :

    ```bash
    truffle migrate --network arbitrum
    ```

2. Create a rollup chain and initialize 3 validators with the given `contract.ao`. This command will initialize your rollup chain and create a `validator-states` folder with configuration information on how to run your chain.

    ```bash
    ../../scripts/setup_rollup.py contract.ao 3
    ```

    > Note: this step may take about 10 minutes the very first time. Subsequent
    > builds are much faster. You can also use the `--up` flag to skip builds
    > if one has completed successfully before.

3. Run the 3 validators

    ```bash
    ../../scripts/arb_deploy.py validator-states
    ```

4. Examine the output from the previous step

    When pet-shop is finished being deployed, start seeing blocks of text like:

    ```txt
    arb-validator1_1  | 2020/02/04 16:25:43
    arb-validator1_1  | == nodes:
    arb-validator1_1  | ==   0:3097fd
    arb-validator1_1  | ==     3:e01d08 leaf latestConfirmed stake:dec077 stake:bcaf2d
    arb-validator1_1  | == stakers:
    arb-validator1_1  | ==   depth:1 addr:dec077 created:5733000 loc:e01d08
    arb-validator1_1  | ==   depth:1 addr:bcaf2d created:5733000 loc:e01d08
    ```

### Use the DApp

1. Install [Metamask](https://metamask.io/)

    > Once Metamask is installed, open it and select
    > `Import Account` and enter one of the following pre-funded private keys
    >
    > ```
    > 0x979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76
    > 0xd26a199ae5b6bed1992439d1840f7cb400d0a55a0c9f796fa67d7c571fbb180e
    > 0xaf5c2984cb1e2f668ae3fd5bbfe0471f68417efd012493538dcd42692299155b
    > 0x9af1e691e3db692cc9cad4e87b6490e099eb291e3b434a0d3f014dfd2bb747cc
    > 0x27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa
    > 0xe4b33c0bb790b88f2463facaf86ae7c17cbdab41187e69ddde8cc1c1fda7c9ab
    > ```

2) Select local network test in Metamask

    - Go back to Metamask or click the extension icon
    - Select `Main Ethereum Network` top right hand side
    - Choose `Custom RPC`
    - Enter `Local Test` as the network name
    - Enter `http://127.0.0.1:7545` as the RPC url
    - Press the save button
    - Metamask should now have an Local Test account holding ETH

3) Launch the front-end

    In another session navigate to `demos/pet-shop` and run:

    ```bash
    yarn start
    ```

    The browser will open to [localhost:8080](http://localhost:8080)

    In the popup window that appears, select `Connect`

4) Adopt some pets

    The pet shop dapp should now be running in your browser. Choose a pet or two
    and click the adopt button to adopt your new animal friend(s).

### Summary

If you want to try another dapp run, first run the following to compile the
Solidity contract and deploy validators:

```bash
cd demos/election
truffle migrate --network arbitrum
../../scripts/setup_rollup.py contract.ao 3
../../scripts/arb_deploy.py validator-states
```

Then open a new command line session, navigate to `demos/election`, and run:

```bash
yarn start
```

The next step is to learn how to port an Ethereum dapp to Arbitrum.

## Porting to Arbitrum

### Prerequisites

The dApp must:

    - Be a Truffle-based project
    - Use web3.js or ethers.js
    - Use webpack or a similar build system

### Overview

Here are the steps needed to port your dApp to Arbitrum:

1. Make sure your dApp compiles and runs correctly on Ethereum or a local testnet
2. Configure the Truffle project to use the Arbitrum Truffle provider (arb-provider-truffle)
3. Add the Arbitrum front-end provider (arb-provider-web3 or arb-provider-ethers)
4. Compile your Truffle project to Arbitrum bytecode (output as contract.ao)
5. Launch a set of Arbitrum Validators with the bytecode
6. Launch the front-end of your dApp

### Setup your workspace

Move your project folder into `arbitrum/workspace/projectname` in order to pick up local versions of arbitrum packages.

### Configure Truffle

1.  Add the `arb-provider-truffle` to your project:

    ```bash
    yarn add arb-provider-truffle
    ```

2.  Edit the `truffle-config.js`:

    -   Import `arb-provider-truffle` and set the mnemonic at the top of the file. This mnemonic is used to set the caller of your contract's constructor when migrating.:

    ```js
    const ArbProvider = require("arb-provider-truffle");
    const mnemonic =
        "jar deny prosper gasp flush glass core corn alarm treat leg smart";
    ```

    -   Add the `arbitrum` network to `module.exports` and `solc` version `0.5.3`:

    ```js
    module.exports = {
        networks: {
            arbitrum: {
                provider: function() {
                    if (!this.provider.prov) {
                        this.provider.prov = ArbProvider.provider(
                            __dirname,
                            "build/contracts",
                            {
                                mnemonic: mnemonic
                            }
                        );
                    }
                    return this.provider.prov;
                },
                network_id: "*"
            }
        }
    };
    ```

3.  Modify your dapp to use the Arbitrum provider

    Find the place in your code where you initialize the ethers (or web3) provider.
    For example, this might be in `src/index.js` or `src/app.js` or somewhere else.

    At the top of the file, find where `ethers` is imported (or `web3`):

    ```js
    const ethers = require("ethers"); // or   const Web3 = require('web3');
    ```

    Right below the ethers import, require the Arbitrum provider:

    ```js
    const ArbProvider = require("arb-provider-ethers");
    ```

    or

    ```js
    const ArbProvider = require("arb-provider-web3");
    ```

    Next, find where web3 (or ethers) is initialized in your code. For example:

    ```js
    let standardProvider = null;
    if (window.ethereum) {
        standardProvider = ethereum;
        try {
            await ethereum.enable();
        } catch (error) {
            console.log("User denied account access");
        }
    } else if (window.web3) {
        standardProvider = web3.currentProvider;
    }
    this.web3 = new Web3(standardProvider);
    ```

    Then set the provider to use Arbitrum instead. For example for web3.js replace the last line with:

    ```js
    this.web3 = new Web3(
        ArbProvider("http://localhost:1235", standardProvider)
    );
    ```

    Or for for ethers.js use

    ```js
    let provider = new ArbProvider(
        "http://localhost:1235",
        new ethers.providers.Web3Provider(standardProvider)
    );
    ```

### Compile to Arbitrum bytecode

Now that the Arbitrum provider is setup correctly, the next step is to compile
the Truffle project into Arbitrum bytecode:

Run the following command to generate `contract.ao`:

```bash
truffle migrate --reset --network arbitrum
```

### Run the Validators

```bash
../../scripts/setup_rollup.py contract.ao 3
../../scripts/arb_deploy.py validator-states
```

### Run the front-end

For example the command might be:

```bash
yarn start
```
