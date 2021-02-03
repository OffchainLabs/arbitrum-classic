---
id: Ethereum_Interoperability
title: Ethereum / Arbitrum Interoperability
sidebar_label: Ethereum Interoperability
---

## Payments and tokens in Arbitrum Chains

An Arbitrum chain functions as an independent blockchain with its security built on the Ethereum blockchain. Moving assets from Ethereum to an Arbitrum chain and back is done through a trustless bridge.

Currently Eth as well as any token confirming to ether the ERC-20 or ERC-721 standard can be transferred over to an Arbitrum chain.

All standard methods of transferring Eth or tokens, if executed within an Arbitrum chain, will only adjust their balance inside that chain. Withdrawing funds from an Arbitrum chain is a separate operation that you can do at any time, if you own those assets within the Arbitrum chain.

### Ethereum to Arbitrum

To move assets into an Arbitrum chain, you execute a deposit transaction on Arbitrum's global EthBridge. This transfers funds to the EthBridge on the Ethereum side, and credits the same funds to you inside the Arbitrum chain you specified. In the case of an ERC-20 or ERC-721 transfer, the Arbitrum chain will spawn a token contract of the correct type within the Arbitrum chain the first time a given token is deposited, at the same token contract address as the token has on Ethereum.

As far as Ethereum knows, all deposited funds are held by Arbitrum's global EthBridge contract.

In order to programmatically trigger transfers, call one of the following methods in the [`GlobalInbox`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/inbox/GlobalInbox.sol) contract with the chain address equal to the address of the Arbitrum Rollup chain.

```solidity
function depositEthMessage(address chain, address to) external payable;
function depositERC20Message(
        address chain,
        address to,
        address erc20,
        uint256 value
    )
        external;
function depositERC721Message(
        address chain,
        address to,
        address erc721,
        uint256 id
    )
        external;
```

### Withdrawing funds from Arbitrum to Ethereum

- Eth: The ArbSys library can be used to withdraw Eth, `ArbSys(100).withdrawEth(destAddress)`
- ERC-20 and ERC-721: The system generated token contracts in Arbitrum contain a withdraw method:
  ```solidity
  function withdrawERC20(address dest, uint256 amount) external;
  function withdrawERC721(address dest, uint256 id) external;
  ```

In all cases, withdrawing is similar to a transfer, except the balance is burned on the Arbitrum side, and eventually those funds become available to their destination on the Ethereum side.

When your withdraw transaction is fully confirmed, the withdrawn funds will be put into your "lockbox" in the EthBridge.
At any time you can call the EthBridge to recover the funds in your lockbox.

## Transaction calls on Arbitrum

### Transaction calls from clients

All client-generated transaction calls on the Arbitrum chain are sent through the EthBridge using:

```solidity
function sendL2Message(address chain, bytes calldata messageData) external;
```

Generally calls will come in batches from an aggregator as described in Transaction Lifestyle. However, Arbitrum supports a number number of different message types described in [ArbOS Formats](ArbOS_Formats.md).

## Transaction calls from Arbitrum to Ethereum

In our initial Rollup release, we are not supporting transaction calls from contracts in an Arbitrum chain to Ethereum contracts. In the future, support for this functionality will be added.
