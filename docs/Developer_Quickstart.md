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
> to give permissions "equivalent to the `root` user", but without the `root`
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

3. Examine the Validator logs

    Go back to the Validators, started with `./arb-deploy contract.ao 3`, and
    examine the output of the logs. You should see the output similar to this:

    ``` txt
    arb-ethbridge_1              | Summary
    arb-ethbridge_1              | =======
    arb-ethbridge_1              | > Total deployments:   10
    arb-ethbridge_1              | > Final cost:          0.229704442061650614 ETH
    arb-ethbridge_1              |
    arb-ethbridge_1              | listening on [::]:17545 ...
    arb-ethbridge_1              | connect to [::ffff:172.23.0.2]:17545 from [TRUNCATED]
    arb-validator-coordinator_1  | Finished waiting for arb-ethbridge:17545...
    arb-validator-coordinator_1  | 0x81183C9C61bdf79DB7330BBcda47Be30c0a85064
    arb-validator-coordinator_1  | Coordinator is creating VM
    arb-validator-coordinator_1  | Got wait request
    arb-validator-coordinator_1  | [DATE] [TIMESTAMP] http: TLS handshake error from 172.23.0.3:39233: EOF
    arb-validator-coordinator_1  | [DATE] [TIMESTAMP] http: TLS handshake error from 172.23.0.4:40329: EOF
    arb-validator2_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | Coordinator serving client
    arb-validator-coordinator_1  | Coordinator upgraded client <nil>
    arb-validator-coordinator_1  | Coordinator serving client
    arb-validator2_1             | Follower connected to coordinator <nil>
    arb-validator1_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | Coordinator serving client
    arb-validator-coordinator_1  | Coordinator upgraded client <nil>
    arb-validator-coordinator_1  | Coordinator serving client
    arb-validator1_1             | Follower connected to coordinator <nil>
    ```

### Use the DApp

1. Install [Metamask](https://metamask.io/)

    If you don't have Metamask already, download the extension and add it to
    your browser following these steps:

    - Select `Get Started`
    - Select `Import Wallet` and you will asked if you want to share diagnostics
    - Wallet Seed: `jar deny prosper gasp flush glass core corn alarm treat leg smart`
    - Note: `password` happens to be 8 characters. Don't use a secure password
      on this dummy seed.
    - `All Done`
    - Keep this browser window open for later

    If you already have Metamask installed, then just import the wallet seed:
    `jar deny prosper gasp flush glass core corn alarm treat leg smart`.

2. Add ETH to Metamask wallet

    You can do this by start Ganache on port 8545 with the mnemonic:

    ``` bash
    MNEMONIC="jar deny prosper gasp flush glass core corn alarm treat leg smart"
    ganache-cli -p 8545 -m "$MNEMONIC"
    ```

3. Select Ganache local network in Metamask

    - Go back to Metamask or click the extension icon
    - Select `Main Ethereum Network` top right hand side
    - Choose `Localhost 8545` (if using another port will need to select `Custom RPC`)
    - You should see `100 ETH` in `Account 1` with address `0x81183C9C61bdf79DB7330BBcda47Be30c0a85064`

4. Launch the frontend

    In another session navigate to `demo-dapp-pet-shop` and run:

    ``` bash
    yarn
    yarn dev
    ```

    The browser will open to [localhost:3000](http://localhost:3000)

    In the popup window that appears, select `Connect`

5. Adopt some pets

    The pet shop dapp should now be running in your browser. Choose a pet or two
    and click the adopt button to adopt your new animal friend(s).

The next step is to learn how to port an Ethereum dapp to Arbitrum.

## Porting to Arbitrum

To get a project up and running on Arbitrum, configure the Truffle project to
use the Arbitrum provider, compile the Solidity contracts to Arbitrum bytecode
in `contract.ao`, launch validators, and use the arbitrum web3 provider or
arbitrum ethers provider to hook into the frontend depending on which provider
approach is used in the existing dapp.

> Note: remember to successfully run your Ethereum dapp locally before
> attempting to port it to Arbitrum.

### Configure Truffle

If you are not already using a Truffle project, go ahead and run `truffle init`
and place your `*.sol` files in the `contracts` folder that is generated.

1. Add the `arb-truffle-provider`:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-truffle-provider.git
    ```

2. Edit the `truffle-config.js`:

    - Import `arb-truffle-provider` and set the mnemonic at the top of the file:

        ``` js
        const ArbProvider = require("arb-truffle-provider");
        const path = require("path");
        const mnemonic = "jar deny prosper gasp flush glass core corn alarm treat leg smart";
        ```

    - Add the `arbitrum` network to `module.exports`:

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

Now that the Arbitrum provider is setup correctly, the next step is to compile
the Solidity files in the `contracts` folder:

1. Generate Arbitrum bytecode

    Run the following command to generate `compiled.json` and `contract.ao`:

    ``` bash
    truffle migrate --network arbitrum
    ```

    Copy the `compiled.json` file into your frontend (i.e. a `src` folder).
    The frontend will need it to use with the arbitrum providers.

2. Run the Validators

    Make sure to copy in the [arb-deploy script](https://github.com/OffchainLabs/demo-dapp-pet-shop/blob/master/arb-deploy) from the demo app.

    ``` bash
    MNEMONIC="jar deny prosper gasp flush glass core corn alarm treat leg smart"
    ./arb-deploy contract.ao 3 -m "$MNEMONIC"
    ```

### Connect the Web3 frontend

1. Add the Arbitrum Web3 Provider

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-ethers-provider.git
    ```

    Note: there is alternatively an `ethers` provider:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-ethers-provider.git
    ```

2. Import the provider in the javascript app

    First, find the existing provider in the project. It could look something
    like this:

    ``` js
    const provider = new Web3.providers.HttpProvider('http://localhost:7545');
    ```

    Replace the existing provider with the Arbitrum Web3 provider:

    ``` js
    const contracts_promise = new Promise(function(resolve, reject) {
        $.getJSON('compiled.json', function(data) {
        resolve(data)
        })
        .fail(function (jqhr, textStatus, error) {
        reject("Failed to load compiled.json. " + textStatus + ": " + error)
        });
    });

    const contracts = await contracts_promise;
    const provider = new ArbProvider(
        'http://localhost:1235',
        contracts,
        new Web3.providers.HttpProvider('http://localhost:7545')
        // // ether provider would be used as:
        //new ethers.providers.JsonRpcProvider('http://localhost:7545')
    );
    ```

    Note: this code must either be in an asynchronous function or the const
    `contents` must be set synchronously. One way to do this is to paste the
    contents of `compiled.json` into the javascript file and set `contents` to
    have the json value.

    Finally, import `arb-web3-provider.js` in the `index.html` as a script. For
    example:

    ``` html
    <script src="js/arb-web3-provider.js"></script>
    ```

3. Run the frontend

    For example the command might be:

    ``` bash
    yarn start
    ```
