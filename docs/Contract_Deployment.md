---
id: Executable_Creation
title: Executable Creation
sidebar_label: Executable Creation
---

Arbitrum support standard EVM contract deployment. This allows standard solidity smart to be deployed on Arbitrum Chains.

To deploy your contracts, you need to set your deployment tool to deploy on an Arbitrum rollup chain instead of Ethereum.

To port an existing truffle configureation:

1.  First add the `arb-provider-web3` and a modern web3 version `web3@^1.2.6` to your project:

    ```bash
    yarn add --dev arb-provider-web3 web3@^1.2.6
    ```

2.  Edit the `truffle-config.js`:

    - Import `arb-provider-web3` and `web3`:

    ```js
    const Web3 = require('web3')
    const ArbProvider = require('arb-provider-web3')
    ```

    - Add the `arbitrum` network to `module.exports`:

    ```js
    module.exports = {
        networks: {
            arbitrum: {
                provider: function () {
                    return ArbProvider(
                        'http://localhost:1235', // Url to an Arbitrum validator with an open rpc interface
                        new Web3.providers.HttpProvider('http://localhost:7545') // Provider to the L1 chain that the rollup is deployed on
                    )
                }
                network_id: "*",
                gasPrice: 0
            }
        }
    };
    ```

Now that the truffle project is setup correctly, just run migrate to deploy your contracts

```bash
truffle migrate --reset --network arbitrum
```
