---
id: Ethereum_Interoperability
title: Ethereum / Arbitrum Interoperability
sidebar_label: Ethereum Interoperability
---

## Payments and tokens in Arbitrum Chains

An Arbitrum chain functions as an independent blockchain with its security built on the Ethereum blockchain. Moving assets from Ethereum to an Arbitrum chain and back is done through a trustless bridge.

Currently Eth as well as any token confirming to ether the ERC-20 or ERC-721 standard can be transferred over to an Arbitrum chain.

All standard methods of transferring Eth or tokens, if executed within an Arbitrum chain, will only adjust their balance inside that chain. Withdrawing funds from an Arbitrum chain is a separate operation that you can do at any time, if you own those assets within the Arbitrum chain.

Visit the [Arbitrum Token Bridge](https://developer.offchainlabs.com/tools/tokenbridge/local/) to easily move funds back and forth using the facilities described in this document.

### Ethereum to Arbitrum

To move assets into an Arbitrum chain, you execute a deposit transaction on Arbitrum's global EthBridge. This transfers funds to the EthBridge on the Ethereum side, and credits the same funds to you inside the Arbitrum chain you specified. In the case of an ERC-20 or ERC-721 transfer, the Arbitrum chain will spawn a token contract of the correct type within the Arbitrum chain, at the same token contract address as the token has on Ethereum.

As far as Ethereum knows, all deposited funds are held by Arbitrum's global EthBridge contract.

In order to programmatically trigger transfers, call one of the following methods in the [`GlobalInbox`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/GlobalInbox.sol) contract with the chain address equal to the address of the Arbitrum Rollup chain.

```
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

- Eth: The ArbSys library can be used to withdraw Eth, `ArbSys(100).withdrawEth(destAddress, amount)`
- ERC-20 and ERC-721: The system generated token contracts in Arbitrum contain a withdraw method
  ```
  function withdraw(address account, uint256 amount) public; // ERC-20
  function withdraw(address account, uint256 tokenId) public; // ERC-721
  ```

In all cases, withdrawing is similar to a transfer, except the balance is burned on the Arbitrum side, and eventually those funds away become available to their destination on the Ethereum side.

When your withdraw transaction is fully confirmed, the withdrawn funds will be put into your "lockbox" in the EthBridge.
At any time you can call the EthBridge to recover the funds in your lockbox.

## Transaction calls on Arbitrum

### Transaction calls from clients

All client-generated transaction calls on the Arbitrum chain are sent through the EthBridge using:

```
function sendTransactionMessage(
        address chain,
        address to,
        uint256 seqNumber,
        uint256 value,
        bytes calldata _data
    )
        external;
```

The sequence number is similar to the account nonce in Ethereum. The rollup chain tracks a sequence number for each account which is the number of transactions that have been accepted from that sender. If the sequence number supplied in the transaction doesn't match the internal sequence number, the transaction is rejected by the Arbitrum chain.

The value and data fields specify the amount of Eth to transfer with the call and the calldata associated with the call respectively.

### Transaction calls from contracts

A smart contract could also use this interface, but formatting calldata and tracking sequence numbers would be complex to do on-chain. In order to simplify Ethereum contract calls to Arbitrum smart contracts, we provide an Arbitrum contract proxy interface.

The proxy interface is a smart contract which implements all methods of the Arbitrum contract which do not have return values. To find a proxy contract address if it already exists, look it up in the `ArbRollup` smart contract.

```
mapping(address => address) public supportedContracts;
```

Given `arbContractAddress`, the address of a contract on an Arbitrum Chain, the address of the corresponding proxy contract
with be `supportedContracts[arbContractAddress]`.

If no such proxy contract already exists, a proxy can be launched using:

```
function spawnCallProxy(address _arbContract) external;
```

Any Ethereum contract can easily and safely make calls to Arbitrum contracts using this interface.

## Transaction calls from Arbitrum to Ethereum

In our initial Rollup release, we are not supporting transaction calls from contracts in a Arbitrum chain to Ethereum contracts. In the future, support for this functionality will be added.
