---
id: Public_Chains
title: Public Arbitrum Chains
sidebar_label: Public Arbitrum Chains
---

The following is a comprehensive list of all of the currently live Arbitrum chains:



| Name                                     | Chain ID | Type    | Underlying L1 | Current Tech Stack | RPC Url(s)                                                                                                                                     | Explorer(s)                                                        | Native Currency | Retryable Dashboard                          |
| ---------------------------------------- | -------- | ------- | ------------- | ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ | --------------- | -------------------------------------------- |
| Arbitrum One                             | 42161    | Mainnet | Ethereum      | Classic Rollup     | https://arb1.arbitrum.io/rpc <br> https://arbitrum-mainnet.infura.io/v3/YOUR-PROJECT-ID <br> https://arb-mainnet.g.alchemy.com/v2/your-api-key | https://arbiscan.io/ https://explorer.arbitrum.io/                 | ETH             | https://retryable-dashboard.arbitrum.io/     |
| Arbitrum Nova                            | 42170    | Mainnet | Ethereum      | Nitro AnyTrust     | [coming soon!]                                                                                                                                 | https://a4ba-explorer.arbitrum.io/                                 | ETH             | http://retryable-tx-panel-nitro.arbitrum.io/ |
| Arbitrum Rinkeby (aka "RinkArby")        | 421611   | Testnet | Rinkeby       | Classic Rollup     | https://rinkeby.arbitrum.io/rpc                                                                                                                | https://testnet.arbiscan.io/ https://rinkeby-explorer.arbitrum.io/ | RinkebyETH      | https://retryable-dashboard.arbitrum.io/     |
| Arbitrum Nitro Goerli Rollup Testnet     | 421613   | Testnet | Goerli        | Nitro Rollup       | https://goerli-rollup.arbitrum.io/rpc                                                                                                          | https://goerli-rollup-explorer.arbitrum.io/                        | GoerliETH       | http://retryable-tx-panel-nitro.arbitrum.io/ |
| [Deprecated Soon!] Arbitrum Nitro Devnet | 421612   | Testnet | Goerli        | Nitro Rollup       | https://nitro-devnet.arbitrum.io/rpc                                                                                                           | https://nitro-devnet-explorer.arbitrum.io/                         | GoerliETH       | http://retryable-tx-panel-nitro.arbitrum.io/ |


For a list of useful contract addresses, look at the list [here](Useful_Addresses.md).

## Arbitrum Chains Summary
**Arbitrum One**: Arbitrum One is the flagship Arbitrum mainnet chain; it is an Optimistic Rollup chain running on top of Ethereum Mainnet, and is open to all users. In an upcoming upgrade, the Arbitrum One chain will be upgraded to use the [Nitro](https://medium.com/offchainlabs/its-nitro-time-86944693bf29) tech stack, maintaining the same state. (Stay tuned for updates!)

**Arbitrum Nova**: Arbitrum Nova is the first  mainnet [AnyTrust](AnyTrust.md) chain; it is currently open for [developer access](https://medium.com/offchainlabs/introducing-nova-arbitrum-anytrust-mainnet-is-open-for-developers-9a54692f345e).

**RinkArby**: RinkArby is the longest running Arbitrum testnet. It runs on the classic stack, and will soon be upgraded to use the Nitro stack. Rinkarby will be deprecated [when Rinkeby itself gets deprecated](https://blog.ethereum.org/2022/06/21/testnet-deprecation/); plan accordingly! 

**Nitro Goerli Rollup Testnet**: This testnet (421613) uses the Nitro rollup tech stack; it is expected to be the primary, stable Arbitrum testnet moving forward.

**Arbitrum Nitro Devnet**: The devnet chain (421612) will soon be deprecated in favor of 421613; for last minute bridging needs, use https://nitro-devnet-bridge.arbitrum.io.

## Using Arbitrum 

_**Note: before interacting with a mainnet chain, users should familiarize themselves with the risks; see [Mainnet Beta](Mainnet.md)**_.

#### Connect Your Wallet
Connect [your wallet](https://portal.arbitrum.one/#wallets) to an Arbitrum chain, adding the chain's RPC endpoint if required. 

#### Get Some Native Currency
You'll need a chain's native currency to transact. You can either acquire funds directly on an Arbitrum chain, or get funds on a chain's underlying L1 and bridge it across. You can get testnet Ether from the following faucets: 
- [Goerli](https://goerlifaucet.com/)
- [Rinkeby](https://faucet.rinkeby.io/)
- [Nitro Goerli Rollup](https://twitter.com/intent/tweet?text=ok%20I%20need%20@arbitrum%20to%20give%20me%20Nitro%20testnet%20gas.%20like%20VERY%20SOON.%20I%20cant%20take%20this,%20I%E2%80%99ve%20been%20waiting%20for%20@nitro_devnet%20release.%20I%20just%20want%20to%20start%20developing.%20but%20I%20need%20the%20gas%20IN%20MY%20WALLET%20NOW.%20can%20devs%20DO%20SOMETHING??%20%20SEND%20HERE:%200xAddA0B73Fe69a6E3e7c1072Bb9523105753e08f8)

[Supported centralized exchanges](https://portal.arbitrum.one/#centralizedexchanges) allow you purchase (mainnet) Ether and withdraw it directly onto Arbitrum one.


### Deposit And Withdraw
To move your Ether and Tokens between Arbitrum and Ethereum chains at [bridge.arbitrum.io](https://bridge.arbitrum.io/).

### Use L2 Dapps!
Interacting with Arbitrum chains will feel very similar to using Ethereum, just cheaper and faster! To get a sense of what's out there, you can check out our [portal page](https://portal.arbitrum.one/), where we showcase some of the dApps, wallets, and infrastructure currently live on Arbitrum One.

### Build on Arbitrum
Dapp developers can build on Arbitrum seamlessly using their favorite Ethereum tooling; see  [here](Contract_Deployment.md) for contract deployment and [here](Frontend_Integration.md) for frontend integration.

### What Next

The team working on Arbitrum is always interested and looking forward to engage with its users.  
Why not follow us on [Twitter](https://twitter.com/arbitrum) or join our community on [Discord](https://discord.gg/5KE54JwyTs)?
