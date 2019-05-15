---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
custom_edit_url: https://github.com/OffchainLabs/developer-website/edit/master/docs/Developer_Quickstart.md
---

Get started with Arbitrum by installing the dependencies, running the demo app,
and porting your Solidity project.

## Install Dependencies

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

### MacOS

1. Install python3, nodejs, & docker using [brew](https://brew.sh/):

    ``` bash
    brew install python3 node docker docker-machine docker-compose
    brew link docker
    brew link docker-compose
    ```

2. Change npm's default directory:

    If you have not installed any npm global packages before,
    [change npm's default directory](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally)
    with the following commands:

    ``` bash
    mkdir ~/.npm-global
    npm config set prefix '~/.npm-global'
    echo $'# npm\nexport PATH=~/.npm-global/bin:$PATH' >> ~/.bash_profile
    source ~/.bash_profile
    ```

3. Install truffle and yarn

    ``` bash
    npm install -g truffle yarn
    ```

### Ubuntu 18.04

1. Install python3, nodejs, docker, truffle, and yarn:

    ``` bash
    sudo apt-get update
    sudo apt-get install -y python3 python3-pip nodejs npm virtualbox docker docker-compose
    sudo npm install -g truffle yarn
    ```

2. Use docker without sudo

    Docker [can be installed](https://docs.docker.com/install/linux/linux-postinstall/)
    to give permissions "equivalent to the `roor` user", but without the `root`
    user group.

    > Warning: "The docker group grants privileges equivalent to the `root`
    > user. For details on how this impacts security in your system, see
    > [Docker Daemon Attack Surface](https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface)".

    Run the following command to use Docker without `sudo`:

    ``` bash
    sudo usermod -aG docker $USER
    ```

    If the `docker` group does not already exist you can create it with:
    `sudo groupadd docker`.

    Finally: log out and log back in before using `docker` without `sudo`.

    Note: if you skip this step you will need to add `sudo` in front of all
    `docker` and `docker-compose` commands. You will also need to build the
    app [manually](#build-manually).

### Full List

Here are the important dependencies in case you are not running on a supported OS:

- [docker](https://github.com/docker/docker-ce/releases) and
  [docker-compose](https://github.com/docker/compose/releases)
- [node and npm](https://nodejs.org/en/)
- [python3 and pip3](https://www.python.org/downloads/)
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
- [virtualbox](https://www.virtualbox.org/wiki/Downloads)
- [yarn](https://yarnpkg.com/en/)

## Install the Arbitrum Compiler

Install the Arbitrum compiler `arbc-truffle-compile` by building it from source:

### Build from source

``` bash
git clone --depth=1 https://github.com/OffchainLabs/arbc-solidity.git
cd arbc-solidity
pip3 install -r requirements.txt
python3 setup.py install
cd ..
```

Note: you may need to run `sudo python3 setup.py install` if running on Ubuntu
or `python3 setup.py install --user` to install without root.

### Check installation

Verify the installation was successful. You may need to open a new shell if the
`arbc-truffle-compile` command is not found.

``` bash
which arbc-truffle-compile
```

The expected output is:

> `/usr/local/bin/arbc-truffle-compile`

## Hello, Arbitrum

There are two ways to get started running your first code on Arbitrum.
Either use the [build script](#build-script) or build the demo app
[manually](#build-manually).

### Build Script

1. Build everything and launch the validators using `arb.py`

    > Note: the `node -v` version must be before 12 since web3.js has a conflict
    > with version 12. Otherwise `yarn` will produce a compile error. We
    > recommend using [nvm](https://github.com/nvm-sh/nvm) to switch versions.

    ``` bash
    git clone --depth=1 https://github.com/OffchainLabs/demo-app.git
    cd demo-app
    yarn
    python3 arb.py
    ```

2. Start the frontend

    Open another shell and go to `demo-app` and run:

    ``` bash
    yarn start
    ```

    The browser will open to [localhost:8080](http://localhost:8080)

### Build Manually

1. Download the demo app

    ``` bash
    git clone --depth=1 https://github.com/OffchainLabs/demo-app.git
    cd demo-app
    ```

    And it's dependencies:

    ``` bash
    mkdir compose
    git clone https://github.com/OffchainLabs/arb-ethbridge.git ./compose/arb-ethbridge
    git clone https://github.com/OffchainLabs/arb-validator.git ./compose/arb-validator
    git clone https://github.com/OffchainLabs/arb-avm.git ./compose/arb-validator/arb-avm
    ```

2. Compile the Fibonacci contract:

    ``` bash
    truffle migrate --network arbitrum
    arbc-truffle-compile compiled.json contract.ao
    docker build -t arb-app -f .arb-app.Dockerfile .
    ```

3. Build

    Note: This step may take about ten minutes the first time. Subsuquent builds
    are much, much faster because intermediate build images are cached.

    ``` bash
    docker-compose build
    ```

    Next build the frontend:

    ``` bash
    yarn
    ```

4. Run the Validators:

    Note: this step will run simultaneously with step 5, so you will need to
    open another bash prompt and continue to step 5 in parallel.

    ``` bash
    docker-compose up --build
    ```

5. Run the Web3 frontend

    The browser will open up [localhost:8080](http://localhost:8080) after
    running the following command:

    ``` bash
    yarn start
    ```

### Use the App

1. Check the validators are running

    You should see the following output from docker-compose at the very end of
    the log:

    ``` txt
    arb-validator-coordinator_1  | Leader is creating VM
    arb-validator-coordinator_1  | Got wait request
    arb-validator-coordinator_1  | 2019/05/13 03:18:46 http: TLS handshake error from 172.19.0.3:40065: EOF
    arb-validator-coordinator_1  | 2019/05/13 03:18:47 http: TLS handshake error from 172.19.0.4:38229: EOF
    arb-validator2_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | Leader serving client
    arb-validator-coordinator_1  | Leader upgraded client <nil>
    arb-validator-coordinator_1  | Client registered
    arb-ethbridge_1              | eth_getTransactionCount
    arb-validator2_1             | Follower connected to leader <nil>
    arb-ethbridge_1              | eth_subscribe
    ...
    arb-ethbridge_1              | eth_subscribe
    arb-validator1_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | Leader serving client
    arb-validator-coordinator_1  | Leader upgraded client <nil>
    arb-validator-coordinator_1  | Client registered
    arb-validator-coordinator_1  | Getting PC 0
    arb-ethbridge_1              | eth_getTransactionCount
    arb-validator1_1             | Follower connected to leader <nil>
    arb-ethbridge_1              | eth_subscribe
    arb-ethbridge_1              | eth_sendRawTransaction
    arb-ethbridge_1              | eth_subscribe
    arb-ethbridge_1              |
    arb-ethbridge_1              |   Transaction: 0x980498aa01a2fb89932dfa14df09fbe0d8c7cc460efc219a00b97fdf1b323887
    arb-ethbridge_1              |   Gas usage: 292476
    arb-ethbridge_1              |   Block Number: 18
    arb-ethbridge_1              |   Block Time: Mon May 13 2019 03:18:52 GMT+0000 (UTC)
    arb-ethbridge_1              |
    arb-ethbridge_1              | eth_subscribe
    ...
    arb-ethbridge_1              | eth_subscribe
    ```

2. Enter number of Fibonacci numbers to generate under "Arbitrum" and "Generate numbers"

    For example, 50. Then you should see:

    > Executing transaction

    Followed by:

    > Successfully generated numbers

    The Arbitrum contract, compiled from the Fibonacci.sol contract, calculated
    the first 50 fibonacci numbers. These can be looked up at indices 0 to 49.
    Generating again would lookup N fibonacci numbers from 1 to N. This next
    sequence can be looked up at indices 50 to N-1. This is because the contract
    stacks the return results in an accumulating fashion instead of overwriting
    them at index 0 to N.

3. Lookup a Fibonacci number

    For exmaple enter 49 under "Lookup numbers". This is the 50th fibonacci
    number and should return the correct result:

    > 12586269025

The next step is porting your own solidity code to an Arbitrum app.

## Porting to Arbitrum

To get a project up and running on Arbitrum, configure the Truffle project to
use the Arbitrum provider, compile the Solidity contracts to Arbitrum bytecode
in `contract.ao`, launch validators, and use the Arbitrum Web3 provider to hook
into the frontend.

### Configure Truffle

If you are not already using a Truffle project, go ahead and run `truffle init`
and place your `*.sol` files in the `contracts` folder that is generated.

1. Add the `arb-truffle-provider`:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-truffle-provider.git
    ```

2. Edit the `truffle-config.js`:

    - Import `arb-truffle-provider` and set a test mnemonic:

        ``` js
        const ArbProvider = require("arb-truffle-provider");
        const path = require("path");
        const mnemonic = "jar deny prosper gasp flush glass core corn alarm treat leg smart";
        ```

    - Add the Arbitrum provider to `module.exports`:

        ``` js
        module.exports = {
          networks: {
            arbitrum: {
              provider: ArbProvider.provider(
                __dirname,
                null,
                {
                  'mnemonic': mnemonic,
                }
              ),
              network_id: "*",
            },
        },
        ```

### Compile Solidity to Arbitrum bytecode

Now that the Arbitrum provider is setup correctly, we are ready to compile the
Solidity files in the `contracts` folder.

1. Generate Arbitrum bytecode

    Run the following commands to generate `compiled.json` and `contract.ao`,
    respectively:

    ``` bash
    truffle migrate --network arbitrum
    arbc-truffle-compile compiled.json contract.ao
    ```

    The `contract.ao` is the Arbitrum bytecode compiled from the EVM bytecode in
    `compiled.json`.

2. Export `contract.ao` as Docker image `arb-contract`

    Next, the compiled Arbitrum bytecode is passed to the validators by creating
    the image `arb-contract` with the single `contract.ao` file:

    ``` bash
    echo "FROM scratch" > .arb-contract.Dockerfile
    echo "COPY contract.ao ./" >>.arb-contract.Dockerfile
    docker build -t arb-contract -f .arb-contract.Dockerfile .
    ```

3. Add the `docker-compose.yml`

    Copy the `docker-compose.yml` file from the [demo-app](https://github.com/OffchainLabs/demo-app/blob/master/docker-compose.yml)
    into the root directory and optionally edit the mnemonic:

    ``` bash
            args:
                MNEMONIC: "jar deny prosper gasp flush glass core corn alarm treat leg smart"
    ```

    This must be the same mnemonic used in the `truffle-config.js` file.

4. Install arbitrum packages

    Clone Arbitrum packages needed as dependencies:

    ``` bash
    mkdir compose
    git clone https://github.com/OffchainLabs/arb-ethbridge.git ./compose/arb-ethbridge
    git clone https://github.com/OffchainLabs/arb-validator.git ./compose/arb-validator
    git clone https://github.com/OffchainLabs/arb-avm.git ./compose/arb-validator/arb-avm
    ```
4. Run multiple Validators:

    Now that the `contract.ao` is exported, we can launch three validators:

    ``` bash
    docker-compose build
    docker-compose up
    ```

### Create a Web3 frontend

1. Add the Arbitrum Web3 Provider:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-web3-provider.git
    ```

2. Import the provider in javascript:

    This example uses a Solidity contract "Fibonacci" and the `compiled.json`
    output from running `truffle migrate --network arbitrum`. The `compiled.json`
    is in the top level folder of the project.

    ``` js
    const ArbProvider = require('arb-web3-provider');

    const contracts = require('../compiled.json');

    let provider = new ArbProvider(
        'http://localhost:1235',
        contracts,
        new ethers.providers.JsonRpcProvider(url)
    );

    // Find "Fibonacci" contract
    var contract = null
    for (var c of contracts) {
        if (c.name == "Fibonacci") {
            contract = c
            break
        }
    }

    let wallet = provider.getSigner(0);
    let ContractRaw = new ethers.Contract(contract.address, contract.abi, provider);
    let myContract = myContractRaw.connect(wallet);
    ```

    This `myContract` can now be used to as an interface to call functions in
    javascript with the same name as in the Solidity file.

    For more information see the demo [Fibonacci contract](https://github.com/OffchainLabs/demo-app/blob/master/contracts/Fibonacci.sol)
    and its [Web3 interface](https://github.com/OffchainLabs/demo-app/blob/master/src/index.js).

3. Run the frontend:

    ``` bash
    yarn start
    ```
