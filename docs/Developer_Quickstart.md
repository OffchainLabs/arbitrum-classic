---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
custom_edit_url: https://github.com/OffchainLabs/developer-website/edit/staging/docs/Developer_Quickstart.md
---

Get started with Arbitrum by installing the Arbitrum compiler,
`arbc-truffle-compile`, and its dependencies. Next,
[build and run the demo app](#hello-arbitrum) or
[port your own dapp](#porting-to-arbitrum).

**Want to learn more? Join the team on [Discord](https://discord.gg/ZpZuw7p) and
read the [white paper](https://OffchainLabs.com/arbitrum.pdf)!**

## Install Dependencies

Follow the instructions for supported operating systems or use the comprehensive
list of dependencies.

> Requires `node -v` version less than 12 for the web3.js frontend in the
> demo app. You can use [nvm](https://github.com/nvm-sh/nvm) to switch
> between multiple node versions.

> Requires python3 to be at least version 3.6

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
    correctly by running `node --version`.

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
git clone -b v0.1.0 --depth=1 -c advice.detachedHead=false https://github.com/OffchainLabs/arbc-solidity.git
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
git clone -b v0.1.1 --depth=1 -c advice.detachedHead=false https://github.com/OffchainLabs/demo-dapp-pet-shop.git
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
    truffle migrate --reset --network arbitrum
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
    arb-ethbridge_1              | > Final cost:          0.229981374766015443 ETH
    arb-ethbridge_1              |
    arb-ethbridge_1              | listening on [::]:17545 ...
    arb-ethbridge_1              | connect to [::ffff:192.168.112.2]:17545 from demo-dapp-pet-shop_arb-validator-coordinator_1.demo-dapp-pet-shop_default:34891 ([::ffff:192.168.112.3]:34891)
    arb-validator-coordinator_1  | Finished waiting for arb-ethbridge:17545...
    arb-validator-coordinator_1  | 2019/06/04 07:10:18 Coordinator is trying to create the VM
    arb-validator-coordinator_1  | 2019/06/04 07:10:19 http: TLS handshake error from 192.168.112.4:41587: EOF
    arb-validator-coordinator_1  | 2019/06/04 07:10:19 http: TLS handshake error from 192.168.112.5:45429: EOF
    arb-validator2_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator1_1             | Finished waiting for arb-validator-coordinator:1236...
    arb-validator-coordinator_1  | 2019/06/04 07:10:21 Coordinator connected with follower 0x38299d74a169e68df4da85fb12c6fd22246add9f
    arb-validator2_1             | 2019/06/04 07:10:21 Validator formed connected with coordinator
    arb-validator-coordinator_1  | 2019/06/04 07:10:21 Coordinator connected with follower 0xc7711f36b2c13e00821ffd9ec54b04a60aefbd1b
    arb-validator-coordinator_1  | 2019/06/04 07:10:21 Coordinator gathering signatures
    arb-validator1_1             | 2019/06/04 07:10:21 Validator formed connected with coordinator
    arb-validator-coordinator_1  | 2019/06/04 07:10:22 Coordinator created VM
    ```

### Use the DApp

1. Install [Metamask](https://metamask.io/)

    If you don't have Metamask already, download the extension and add it to
    your browser and create a new account.

2. Select Ganache local network in Metamask

    - Go back to Metamask or click the extension icon
    - Select `Main Ethereum Network` top right hand side
    - Choose `Custom RPC`
    - Enter `Ganache` as the network name
    - Enter `http://127.0.0.1:7545` as the RPC url
    - Press the save button

3. Add pre-funded accounts to metamask
    - Go back to Metamask or click the extension icon
    - Select the circle icon on the top right hand side
    - Select `Import Account`
    - Enter any of the following private keys
        - `0x41a9550a0ae23fd52f3b99acab194db2e4474262db64dfd46807bca9e061e211`
        - `0x77500b500284eab4d5201d230ca015b82c32752e42c79dc3d6ff3668ada9d340`
        - `0x54f4370ee20fd563acaac3ea63eef5cc62d3e0cb11f7f03e70180e538c882bc8`
        - `0xa36dd563650acd8305d222a68abcaa4b3db69f28cc40d0abba391ec58ac12fba`
        - `0x2090bf383976cdcb04fc776585f5e65f71929be0e36d53ffc8eb066ef8ec2d18`
        - `0x1b153b674c13af2974acbb66027fa4386b85b31cb27d159276d05e9542359f3f`
    - Metamask should now have an Ganache testnet account holding ETH

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
    yarn add https://github.com/OffchainLabs/arb-truffle-provider#v0.1.0
    ```

2. Edit the `truffle-config.js`:

    - Import `arb-truffle-provider` and set the mnemonic at the top of the file:

        ``` js
        const ArbProvider = require("arb-truffle-provider");
        const mnemonic = "jar deny prosper gasp flush glass core corn alarm treat leg smart";
        ```

    - Add the `arbitrum` network to `module.exports`:

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
        ```

### Compile Solidity to Arbitrum bytecode

Now that the Arbitrum provider is setup correctly, the next step is to compile
the Solidity files in the `contracts` folder:

1. Generate Arbitrum bytecode

    Run the following command to generate `compiled.json` and `contract.ao`:

    ``` bash
    truffle migrate --reset --network arbitrum
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

1. Add the Arbitrum `web3.js` Provider

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-web3-provider#v0.1.0
    ```

    Note: there is alternatively an `ethers.js` provider:

    ``` bash
    yarn add https://github.com/OffchainLabs/arb-ethers-provider#v0.1.0
    ```

2. Import the provider in the javascript app

    First, find the existing provider in the project. It could look something
    like this:

    ``` js
    const provider = new Web3.providers.HttpProvider('http://localhost:7545');
    ```

    Replace the existing provider with the Arbitrum Web3 provider:

    ``` js
    const contracts = await new Promise(function(resolve, reject) {
        $.getJSON('compiled.json', function(data) {
            resolve(data)
        }).fail(function (jqhr, textStatus, error) {
            reject("Failed to load compiled.json. " + textStatus + ": " + error)
        });
    });

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
