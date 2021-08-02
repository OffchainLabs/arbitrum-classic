---
id: Contract_Deployment
title: Contract Deployment
sidebar_label: Contract Deployment
---

Arbitrum supports standard EVM contract deployment. This allows standard Solidity smart contracts to be deployed on Arbitrum Chains using existing developer tools.

To deploy your contracts, you need to set your deployment tool to deploy on an Arbitrum rollup chain instead of Ethereum. While this should be straightforward, we include instructions for some build systems here, and we will add to the list over time. If you're using a build system that's not listed here and having trouble configuring it, please reach out to us on [Discord](https://discord.gg/ZpZuw7p).

## Hardhat

To port an existing hardhat configuring, simply include the Arbitrum RPC url in `hardhat.config.ts`:

```ts
module.exports = {
  solidity: '0.7.3',
  networks: {
    arbitrum: {
      url: 'https://rinkeby.arbitrum.io/rpc',
      gasPrice: 0,
    },
  },
}
```

See [Pet Shop Demo](https://github.com/OffchainLabs/arbitrum-tutorials/packages/demo-dapp-pet-shop).

## Truffle

To port an existing truffle configuration:

1.  First add the `arb-ethers-web3-bridge` to your project:

    ```bash
    yarn add --dev arb-ethers-web3-bridge
    ```

2.  Edit the `truffle-config.js`:

    - Set the mnemonic and the url to an Arbitrum aggregator at the top of the file.

    ```js
    const HDWalletProvider = require('@truffle/hdwallet-provider')

    const mnemonic =
      'jar deny prosper gasp flush glass core corn alarm treat leg smart'
    const arbProviderUrl = 'http://localhost:8547/'
    ```

    - Add the `arbitrum` network to `module.exports`:

    ```js
    module.exports = {
        arbitrum: {
          provider: function () {
              return new HDWalletProvider(mnemonic, arbProviderUrl)
            )
          },
          network_id: '*',
          gasPrice: 0,
        },
      },
    }
    ```

Now that the truffle project is set up correctly, just run migrate to deploy your contracts

```bash
truffle migrate --reset --network arbitrum
```

For older versions of truffle (< 0.5.x), do the following:

- Import `wrapProvider` from `arb-ethers-web3-bridge` at the top of `truffle-config.js`:

```bash
const wrapProvider = require('arb-ethers-web3-bridge').wrapProvider
```

- return the wrapped provider here:
  ```js
    module.exports = {
        arbitrum: {
          provider: function () {
            // return wrapped provider:
            return wrapProvider(
              new HDWalletProvider(mnemonic, arbProviderUrl)
            )
          },
          network_id: '*',
          gasPrice: 0,
        },
      },
    }
  ```
