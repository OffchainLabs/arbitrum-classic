---
id: Porting
title: Porting Existing Dapps to Arbitrum
sidebar_label: Porting
---

## Porting to Arbitrum

Existing Ethereum applications can be ported to Arbitrum as long as they only use [supported](Solidity_Support.md) parts of Solidity.

### Prerequisites

The dApp must:

    - Be a Truffle-based project
    - Use web3.js or ethers.js
    - Use webpack or a similar build system

### Overview

Here are the steps needed to port your dApp to Arbitrum:

1. Make sure your dApp compiles and runs correctly on Ethereum or a local testnet
2. Launch a set of Arbitrum Validators on a [Local Testnet](Local_Blockchain.md) or on [Rinkeby](Rinkeby.md)
3. Configure the Truffle project to use the Arbitrum web3 provider and [deploy your contracts to your Arbitrum chain](Contract_Deployment.md)
4. Add the Arbitrum [front-end provider](Frontend_Integration.md)
5. Launch the front-end of your dApp
