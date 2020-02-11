---
id: Developer_Quickstart
title: Arbitrum Developer Quickstart
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Developer_Quickstart.md
---

Get started with Arbitrum by [installing](Installation.md) the Arbitrum compiler,
`arbc-truffle`, and its dependencies. Next,
[build and run the demo app](#hello-arbitrum) or
[port your own dapp](#porting-to-arbitrum).

Arbitrum has three modes: channels, AnyTrust sidechains, and rollup. Channels and sidechains provide the AnyTrust Guarantee which ensures that the code will run correctly as long as any validator is honest.

The following documention describes how to use Arbitrum Rollup.

**Want to learn more? Join the team on [Discord](https://discord.gg/ZpZuw7p) and
read about [how Arbitrum Rollup works](https://medium.com/offchainlabs/how-arbitrum-rollup-works-39788e1ed73f)!**

## Setup Blockchain

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

1. Compile Solidity to Arbitrum:

    Truffle will output the compiled contract as `contract.ao` :

    ```bash
    truffle migrate --network arbitrum
    ```

2. Create a rollup chain with the given `contract.ao` and prepare 3 validators to validate it.

    ```bash
    ../../scripts/setup_rollup.py contract.ao 3
    ```

    > Note: this step may take about 10 minutes the very first time. Subsequent
    > builds are much faster. You can also use the `--up` flag to skip builds
    > if one has completed successfully before.

3. Run the 3 validators given a path to the `validator-states` folder created by the previous command

    ```bash
    ../../scripts/arb_deploy.py validator-states
    ```

4. Examine the output from the previous step

    When pet-shop is finished being deployed, you should start seeing blocks of text like:

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
