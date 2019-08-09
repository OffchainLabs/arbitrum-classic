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

### MacOS

1. Install python3, nodejs, & docker using [Homebrew](https://brew.sh/):

    ``` bash
    brew install python3 node@8 docker docker-compose
    brew unlink node
    brew link --force --overwrite node@8
    brew cask install docker
    open -a Docker
    ```

    Once the Docker app appears in the menu bar, wait until the yellow light turns
    green (no need to log into Docker). Also check that node version 8 is installed
    correctly by running `node -v`.

2. Change npm's default directory:

    If you have not installed any npm global packages before,
    [change npm's default directory](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally)
    with the following commands:

    ``` bash
    mkdir ~/.npm-global
    npm config set prefix '~/.npm-global'
    echo $'# npm\nexport PATH="~/.npm-global/bin:$PATH"' >> ~/.bash_profile
    source ~/.bash_profile
    ```

3. Install truffle and yarn

    ``` bash
    npm install -g truffle ganache-cli yarn
    ```

### Ubuntu 18.04

Install python3, nodejs, docker, truffle, and yarn:

``` bash
sudo apt-get update
sudo apt-get install -y python3 python3-pip nodejs npm docker docker-compose
sudo npm install -g truffle ganache-cli yarn
```

> Docker [can be used without sudo](https://docs.docker.com/install/linux/linux-postinstall/)
> to give permissions "equivalent to the `root` user". See [the security warning](https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface).

### Full List

Here are the important dependencies in case you are not running on a supported OS:

- [docker](https://github.com/docker/docker-ce/releases) and
  [docker-compose](https://github.com/docker/compose/releases)
- [node and npm](https://nodejs.org/en/)
- [python3 and pip3](https://www.python.org/downloads/)
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
- [ganache](https://www.npmjs.com/package/ganache-cli)
- [yarn](https://yarnpkg.com/en/)

> Requires `node -v` version 8 or 10

> Requires`python3 --version` 3.6 or greater

> Requires `solc` 0.5.10 or less in `truffle version`

## Install Arbitrum

Download the Arbitrum Monorepo from source:

``` bash
git clone -b v0.2.0 --depth=1 -c advice.detachedHead=false https://github.com/offchainlabs/arbitrum.git
cd arbitrum
yarn
yarn install:deps
```

Check `arbc-truffle` was installed:

```
which arbc-truffle
```

Expected output:

> /usr/local/bin/arbc-truffle

## Hello, Arbitrum

Now you'll compile and run a demo dApp on Arbitrum. The dApp is based on
a simple Pet Shop dApp that is used in a Truffle tutorial.

Inside the Arbitrum monorepo, open the `pet-shop` demo dApp:

```
cd demos/pet-shop
```

### Build and Run

You'll need to do these steps every time you make a change to the Solidity. For
this dApp, you do not need to change any Solidity files.

1. Compile Solidity to Arbitrum:

    Truffle will output the compiled contract as `contract.ao` as well as a
    `compiled.json` file needed for the frontend:

    ``` bash
    truffle migrate --reset --compile-all --network arbitrum
    ```

    > Note: if `--compile-all` if not recognized then try `--all`

    Return to the root of the Monorepo:

    ```
    cd ../..
    ```

2. Deploy `contract.ao` to 3 Validators

    ``` bash
    ./scripts/arb_deploy.py demos/pet-shop/contract.ao 3
    ```

    > Note: this step may take about 10 minutes the very first time. Subsequent
    > builds are much faster. You can also use the `--up` flag to skip builds
    > if one has completed successfully before.

3. Examine the `yarn deploy contract.ao 3` output

    When pet-shop is finished being deployed, you should see output similar to this:

    ``` txt
    arb-validator-coordinator_1  | Finished waiting for arb-bridge-eth:7545...
    arb-validator-coordinator_1  | 2019/08/09 23:49:12 Coordinator is trying to create the VM
    arb-validator-coordinator_1  | 2019/08/09 23:49:13 http: TLS handshake error from 192.168.208.4:32963: EOF
    arb-validator-coordinator_1  | 2019/08/09 23:49:13 http: TLS handshake error from 192.168.208.5:37875: EOF
    arb-validator2_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator1_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | 2019/08/09 23:49:15 Coordinator connected with follower 0x85794eceb590b9b53554bc6d28c964be00aaa893
    arb-validator2_1             | 2019/08/09 23:49:15 Validator formed connection with coordinator
    arb-validator-coordinator_1  | 2019/08/09 23:49:15 Coordinator connected with follower 0xfbb0fc9161f9c824cb5ff5222166b7ea247e85ca
    arb-validator-coordinator_1  | 2019/08/09 23:49:15 Coordinator gathering signatures
    arb-validator1_1             | 2019/08/09 23:49:15 Validator formed connection with coordinator
    arb-validator-coordinator_1  | 2019/08/09 23:49:16 Coordinator created VM
    ```

### Use the DApp

1. Install [Metamask](https://metamask.io/)

    If you don't have Metamask already, download the extension and add it to
    your browser and create a new account by importing the following seed phrase:

    ``` text
    jar deny prosper gasp flush glass core corn alarm treat leg smart
    ```

    > If you already have Metamask installed, then open it and select
    > `Import Account` and enter one of the following private keys
    > derived from the mnemonic listed above:
    > ```
    > 0x41a9550a0ae23fd52f3b99acab194db2e4474262db64dfd46807bca9e061e211
    > 0x77500b500284eab4d5201d230ca015b82c32752e42c79dc3d6ff3668ada9d340
    > 0x54f4370ee20fd563acaac3ea63eef5cc62d3e0cb11f7f03e70180e538c882bc8
    > 0xa36dd563650acd8305d222a68abcaa4b3db69f28cc40d0abba391ec58ac12fba
    > 0x2090bf383976cdcb04fc776585f5e65f71929be0e36d53ffc8eb066ef8ec2d18
    > 0x1b153b674c13af2974acbb66027fa4386b85b31cb27d159276d05e9542359f3f
    > ```

    This mnemonic is the default used by `yarn deploy` and these accounts will
    be pre-funded.

2. Select Ganache local network in Metamask

    - Go back to Metamask or click the extension icon
    - Select `Main Ethereum Network` top right hand side
    - Choose `Custom RPC`
    - Enter `Ganache` as the network name
    - Enter `http://127.0.0.1:7545` as the RPC url
    - Press the save button
    - Metamask should now have an Ganache testnet account holding ETH

4. Launch the frontend

    In another session navigate to `demo-dapp-pet-shop` and run:

    ``` bash
    yarn start
    ```

    The browser will open to [localhost:8080](http://localhost:8080)

    In the popup window that appears, select `Connect`

5. Adopt some pets

    The pet shop dapp should now be running in your browser. Choose a pet or two
    and click the adopt button to adopt your new animal friend(s).

### Summary

If you want to try another dapp run these steps once:

``` bash
git clone -b v0.2.0 --depth=1 -c advice.detachedHead=false https://github.com/offchainlabs/demo-dapp-election.git
cd demo-dapp-election
yarn
```

and run these steps every time a change is made to the Solidity contract:

``` bash
truffle migrate --reset --compile--all --network arbitrum
yarn deploy contract.ao 3
```

Open a new command line session, navigate to `demo-dapp-election`, and run:

``` bash
yarn start
```

The next step is to learn how to port an Ethereum dapp to Arbitrum.

## Porting to Arbitrum

### Prerequisites

The dApp must:

- Be a Truffle-based project
- Use web3 or ethers
- Use webpack or a similar build system

### Overview

Here are the steps needed to port your dApp to Arbitrum:

   1. Make sure your dApp compiles and runs correctly on Ethereum or ganache
   2. Configure the Truffle project to use the Arbitrum Truffle provider (arb-provider-truffle)
   3. Add the Arbitrum frontend provider (arb-provider-web3 or arb-provider-ethers)
   4. Compile your Truffle project to Arbitrum bytecode (output as contract.ao)
   5. Launch a set of Arbitrum Validators with the bytecode
   6. Launch the frontend of your dApp

### Configure Truffle

First, identify if your project uses web3 or truffle. The following examples
will use truffle. If you are using web3 just replace arb-provider-truffle with
arb-provider-web3 for these steps:

1. Add the `arb-provider-truffle` to your project:

    ``` bash
    yarn add https://github.com/offchainlabs/arb-provider-truffle#v0.2.0
    ```

2. Edit the `truffle-config.js`:

    - Import `arb-provider-truffle` and set the mnemonic at the top of the file:

        ``` js
        const ArbProvider = require("arb-provider-truffle");
        const mnemonic = "jar deny prosper gasp flush glass core corn alarm treat leg smart";
        ```

    - Add the `arbitrum` network to `module.exports` and `solc` version `0.5.3`:

        ``` js
        module.exports = {
          networks: {
            arbitrum: {
              provider: function() {
                if(typeof this.provider.prov == 'undefined') {
                    this.provider.prov = ArbProvider.provider(
                      __dirname,
                      'build/contracts',
                      {
                        'mnemonic': mnemonic,
                      }
                    );
                }
                return this.provider.prov
              },
              network_id: "*",
            },
          },
        },
        compilers: {
          solc: {
            version: "0.5.3",
            docker: true,
            settings: {
              optimizer: {
                enabled: true,
                runs: 200
              }
            }
          }
        }
        ```

        > Requires solc version `0.5.10` or below

3. Modify your dapp to use the Arbitrum provider

    Find the place in your code where you initialize the ethers (or web3) provider.
    For example, this might be in `src/index.js` or `src/app.js` or somewhere else.

    At the top of the file, find where `ethers` is imported (or `web3`):

    ``` js
    const ethers = require('ethers'); // or   const Web3 = require('web3');
    ```

    Right below the ethers import, require the Arbitrum provider:

    ``` js
    const ArbProvider = require('arb-provider-ethers')
    ```

    Next, find where the ethers (or web3) provider is initialized in your code,
    and set the provider to use Arbitrum instead. For example for ethers:

    ``` js
    var standardProvider = null;
    if (window.ethereum) {
      standardProvider = ethereum;
      try {
        // Request account access if needed
        await ethereum.enable();
      } catch (error) {
        console.log("User denied account access")
      }
    } else if (window.web3) {
      // Legacy dapp browsers...
      standardProvider = web3.currentProvider;
    } else {
      // Non-dapp browsers...
      console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
    }
    let standardProvider = null;
    let contracts = require('../compiled.json');
    let App.provider = new ArbProvider(
      'http://localhost:1235',
      contracts,
      new ethers.providers.Web3Provider(standardProvider)
    );
    ```

    If you are using web3 instead, then the code would look like this:

    ``` js
    const contracts = require('../compiled.json');
    App.provider = await ArbProvider(
      'http://localhost:1235',
      contracts,
      new Web3.providers.HttpProvider('http://localhost:7545')
    );
    ```

    > Note: make the path to `compiled.json` correspond to the root directory of the project

### Compile to Arbitrum bytecode

Now that the Arbitrum provider is setup correctly, the next step is to compile
the Truffle project into Arbitrum bytecode:

Run the following command to generate `compiled.json` and `contract.ao`:

``` bash
truffle migrate --reset --compile-all --network arbitrum
```

We do not need to copy the `compiled.json` file into the frontend folder because
it has already been setup to correctly to retrieve the json in the previous part.

### Run the Validators

``` bash
yarn deploy contract.ao 3
```

### Run the frontend

For example the command might be:

``` bash
yarn
start
```
