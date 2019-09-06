---
id: Solidity_Support
title: Solidity Support
custom_edit_url: https://github.com/OffchainLabs/arbitrum/edit/master/docs/Solidity_Support.md
---

Arbitrum provides the ability to take arbitrary solidity contracts and run them offchain in a trustless layer 2 context. To accomplish this, Arbitrum compiles the contracts to a customized virtual machine architecture, optimized for offchain execution.

# Restrictions

Although we support most solidity code, there are a number of restrictions that currently exist.

-   No support for general calls to external contracts
    -   In the future, we'll provide mechanisms for Arbitrum contracts to interoperate with standard Ethereum contracts. However this is quite difficult when operating offchain.
    -   We do support deploying multiple Solidity contracts as a single unit (VM) in Arbitrum. If you're using truffle, you can deploy multiple contracts to Arbitrum and they will all be able to call each other naturally.
-   No support for a running Arbitrum contract dynamically creating new contracts
    -   Currently we don't support spawning new contracts from your deployed Arbitrum contract. This support will be added in the future.

# Workarounds

-   Address support
    -   Arbitrum API functions use 32 byte addresses instead of 20 bytes addresses. To convert a normal Ethereum address into input for an on-chain Arbitrum API call, use `bytes32(bytes20(address))`
-   Token Support
    -   Arbitrum VMs support sending and receiving Eth, ERC20, and ERC721 tokens through the on-chain EthBridge contract. Tokens can be deposited into an Arbitrum wallet through the [`ArbBalanceTracker`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/ArbBalanceTracker.sol) contract by calling `depositEth`, `depositERC20`, or `depositERC721` respectively. Then tokens can be sent to an Arbitrum contract using the [`VMTracker`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/VMTracker.sol) `sendMessage` function.
    -   Tokens can be sent from an Arbitrum contract by use of the [`ArbSys`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-compiler-evm/ArbSys.sol) interface. For example use
        `ArbSys(address(0x01)).sendERC20( bytes32(bytes20(address)), tokenAddress, amount );`
        to send tokens from the Arbitrum contract to a particular Ethereum address.
    -   When Eth and tokens are sent to a non-Arbitrum contract address, they are placed into the `ArbBalanceTracker` wallet f. A call to the `ArbBalanceTracker` is necessary to withdraw them back to your standard Ethereum wallet.
