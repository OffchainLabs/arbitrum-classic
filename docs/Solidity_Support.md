---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum Rollup allows you to deploy a set of Solidity contracts as a trustless layer 2 sidechain. To accomplish this, Arbitrum compiles the contracts to a customized virtual machine architecture, optimized for off-chain execution.

# Compilation

A set of solidity contracts can be compiled into an Arbitrum virtual machine.

Currently the only supported deployment mechanism for Arbitrum is through a truffle plugin. Add the following block to your `truffle-config.js` file to add an Arbitrum deployment configuration to your project.

```js
arbitrum: {
  provider: function() {
    if (typeof this.provider.prov == "undefined") {
      this.provider.prov = ArbProvider.provider(
        __dirname,
        "build/contracts",
        {
          mnemonic: mnemonic
        }
      );
    }
    return this.provider.prov;
  },
  network_id: "*"
}
```

After configuring truffle, compile your contracts into an Arbitrum VM using:

`truffle migrate --network arbitrum`

Running that command will produce a compiled Arbitrum VM binary, `contract.ao`.

# Restrictions

Although we support most solidity code, there are a number of restrictions that currently exist.

-   Unsupported Solidity Features:

    -   `blockhash(uint blockNumber) returns (bytes32)`
    -   `block.coinbase`
    -   `block.difficulty`
    -   `block.gaslimit`
    -   `block.timestamp`
    -   `gasleft() returns (uint256)`
    -   `now`
    -   General contract creation

# Workarounds

-   Contract cloning
    -   Despite the fact that we don't support standard contract creation, we do support the cloning of already deployed (producing a similar result as [EIP-1167](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1167.md)).
    -   Cloning an existing contract is done through `ArbSys(100).cloneContract(contractAddress)`
