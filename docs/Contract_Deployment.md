---
id: Executable_Creation
title: Executable Creation
sidebar_label: Executable Creation
---

Arbitrum supports standard EVM contract deployment. This allows standard Solidity smart contracts to be deployed on Arbitrum Chains using existing developer tools.

To deploy your contracts, you need to set your deployment tool to deploy on an Arbitrum rollup chain instead of Ethereum. While this should be straightforward, we include instructions for some build systems here, and we will add to the list over time. If you're using a build system that's not listed here and having trouble configuring it, please reach out to us on [Discord](https://discord.gg/ZpZuw7p).

## Truffle 

To port an existing truffle configuration:

1.  First add the `arb-ethers-web3-bridge`, `ethers`, and `arb-provider-ethers` to your project:

    ```bash
    yarn add --dev arb-ethers-web3-bridge ethers@^4.0.44 arb-provider-ethers
    ```

2.  Edit the `truffle-config.js`:

    - Import `arb-ethers-web3-bridge` and `ethers`, and `arb-provider-ethers`, and set the mnemonic at the top of the file:

    ```js
    const ethers = require('ethers')
    const ArbEth = require('arb-provider-ethers')
    const ProviderBridge = require('arb-ethers-web3-bridge')

    const mnemonic =
      'jar deny prosper gasp flush glass core corn alarm treat leg smart'
    ```

    - Add the `arbitrum` network to `module.exports`:

    ```js
    module.exports = {
        networks: {
            arbitrum: {
                provider: function () {
                    // Provider to the L1 chain that the rollup is deployed on
                    const provider = new ethers.providers.JsonRpcProvider(
                      'http://localhost:7545'
                    )
                    const arbProvider = new ArbEth.ArbProvider(
                      'http://localhost:1235', // Url to an Arbitrum validator with an open rpc interface
                      provider
                    )
                    const wallet = new ethers.Wallet.fromMnemonic(mnemonic).connect(
                      provider
                    )
                    return new ProviderBridge(
                      arbProvider,
                      new ArbEth.ArbWallet(wallet, arbProvider)
                    )
                }
                network_id: "*",
                gasPrice: 0
            }
        }
    };
    ```

Now that the truffle project is set up correctly, just run migrate to deploy your contracts

```bash
truffle migrate --reset --network arbitrum
```
