---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
---

Get started with Arbitrum by installing the dependencies, running the demo app,
and porting your Solidity project.

## Install

### Dependencies

Here is the complete list of dependencies; you may already have them all
installed. If not, there are instructions for each supported operating system:

- [docker](https://github.com/docker/docker-ce/releases) and
  [docker-compose](https://github.com/docker/compose/releases)
- [node and npm](https://nodejs.org/en/)
- [python3 and pip3](https://www.python.org/downloads/)
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
- [virtualbox](https://www.virtualbox.org/wiki/Downloads)
- [yarn](https://yarnpkg.com/en/)

#### MacOS

``` bash
brew install python3 node docker docker-machine docker-compose
npm install -g truffle yarn
```

#### Ubuntu 18.04

``` bash
sudo apt-get install -y python3 pip3 nodejs npm virtualbox docker docker-compose
npm install -g truffle yarn
```

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

## Hello, Arbitrum

Download and build the demo app to run your first app using Arbitrum. You can
follow these instructions or just run `python3 arb.py` after downloading the
demo to do these steps for you.

1. Download the demo app

    ``` bash
    git clone --depth=1 https://github.com/OffchainLabs/demo-app.git
    cd demo-app
    ```

    And it's dependencies:

    ```
    mkdir compose
    git clone https://github.com/OffchainLabs/arb-ethbridge.git ./compose/arb-ethbridge
    git clone https://github.com/OffchainLabs/arb-validator.git ./compose/arb-validator
    git clone https://github.com/OffchainLabs/arb-avm.git ./compose/arb-validator/arb-avm
    ```

2. Build

    Note: This step takes about seven minutes the first time. Subsuquent builds
    are a matter of seconds.

    ``` bash
    sudo docker-compose build
    ```

    And build the frontend:

    ``` bash
    yarn
    ```

3. Compile the Fibonacci contract:

    ``` bash
    truffle migrate --network arbitrum
    arbc-truffle-compile compiled.json contract.ao
    docker build -t arb-app -f .arb-app.Dockerfile .
    ```

4. Export the contract and build and run 3 Validators:

    ``` bash
    sudo docker-compose up --build
    ```

5. Run the Web3 frontend

    Make sure to open another bash prompt since step 4 is running the validators

    ``` bash
    yarn start
    ```

    The browser will open up [localhost:8080](http://localhost:8080).

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
                path.resolve(__dirname, 'compiled.json'),
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

2. Export `contract.ao` as Docker image `arb-app`

    Next, the compiled Arbitrum bytecode is passed to the validators by creating
    the image `arb-app` with the single `contract.ao` file:

    ``` bash
    echo $"FROM scratch\nCOPY contract.ao ./" > .arb-app.Dockerfile
    sudo docker build -t arb-app -f arb-app.Dockerfile .
    ```

3. Run multiple Validators:

    Now that the `contract.ao` is exported, we can launch three validators:

    ``` bash
    sudo docker-compose build
    sudo docker-compose up
    ```

### Create a Web3 frontend

1. Add the Arbitrum Web3 Provider:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-web3-provider.git
    ```

2. Import the provider in javascript:

    This example uses a Solditiy contract "Fibonacci" and the `compiled.json`
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
