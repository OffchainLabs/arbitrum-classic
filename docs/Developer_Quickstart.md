---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
custom_edit_url: https://github.com/OffchainLabs/developer-website/edit/master/docs/Developer_Quickstart.md
---

Get started with Arbitrum by installing the Arbitrum compiler,
`arbc-truffle-compile`, and its dependencies. Next,
[build and run the demo app](#hello-arbitrum) or
[port your own dapp](#porting-to-arbitrum).

## Install Dependencies

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

### MacOS

1. Install python3, nodejs, & docker using [Homebrew](https://brew.sh/):

    > Requires `node -v` version less than 12 for the web3.js frontend in the
    > demo app. You can use [nvm](https://github.com/nvm-sh/nvm) to switch
    > between multiple node versions.

    ``` bash
    brew install python3 node@8 docker docker-machine docker-compose
    brew link docker
    brew link docker-compose
    brew unlink node
    brew link node@8
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

Install python3, nodejs, docker, truffle, and yarn:

``` bash
sudo apt-get update
sudo apt-get install -y python3 python3-pip nodejs npm docker docker-compose
sudo npm install -g truffle yarn
```

> Docker [can be used without sudo](https://docs.docker.com/install/linux/linux-postinstall/)
> to give permissions "equivalent to the `roor` user", but without the `root`
> user group. See: [Docker Daemon Attack Surface](https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface).

### Full List

Here are the important dependencies in case you are not running on a supported OS:

- [docker](https://github.com/docker/docker-ce/releases) and
  [docker-compose](https://github.com/docker/compose/releases)
- [node and npm](https://nodejs.org/en/)
- [python3 and pip3](https://www.python.org/downloads/)
- [truffle](https://truffleframework.com/docs/truffle/getting-started/installation)
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

Note: you may need to run `sudo python3 setup.py install` if running on Ubuntu.

### Check installation

Verify the installation was successful:

``` bash
which arbc-truffle-compile
```

The expected output is:

> `/usr/local/bin/arbc-truffle-compile`

## Hello, Arbitrum

Now you'll compile and run a demo dapp on Arbitrum. The dapp is based on
a simple Pet Shop dapp that is used in a Truffle tutorial.

### Install

You only need to run these commands once:

``` bash
git clone --depth=1 https://github.com/OffchainLabs/demo-dapp-pet-shop.git
cd demo-dapp-pet-shop
yarn
```

### Build and Run

You'll need to do these steps every time you make a change to the Solidity. For
this dapp, you do not need to change any Solidity files.

1. Compile Solidity to Arbitrum:

    Truffle will output the compiled contract as `contract.ao` as well as a
    `compiled.json` file needed for the frontend:

    ``` bash
    truffle migrate --network arbitrum
    ```

    Move the `compiled.json` folder into the frontend:

    ``` bash
    mv compiled.json src
    ```

2. Deploy `contract.ao` to 3 Validators

    > Note: this step may take about 10 minutes the very first time. Subsequent
    > builds are much, much faster because they are cached as Docker images.

    ``` bash
    ./arb-deploy contract.ao 3
    ```

3. Start the frontend

    Open another shell and go to `demo-app` and run:

    ``` bash
    yarn start
    ```

    The browser will open to [localhost:8080](http://localhost:8080)

### Run

1. Examine the Validator logs

    You should see the following output at the end of the `arb-deploy` output:

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

2. Adopt some pets

    The pet shop dapp should now be running in your browser. Choose a pet or two and click the adopt
    button to adopt your new animal friend(s).

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
                'build/contracts',
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

    Run the following command to generate `compiled.json` and `contract.ao`:

    ``` bash
    truffle migrate --network arbitrum
    ```

2. Run the Validators

    Make sure to copy in the [arb-deploy script](https://github.com/OffchainLabs/demo-app/blob/master/arb-deploy)
    from the demo app.

    ``` bash
    MNEMONIC="jar deny prosper gasp flush glass core corn alarm treat leg smart"
    ./arb-deploy contract.ao 3 -m "$MNEMONIC"
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
