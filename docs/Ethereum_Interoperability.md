# Ethereum Interoperability

## Payments and tokens in Arbitrum Chains

An Arbitrum chain functions as an independent blockchain with it's security sourced in the Ethereum blockchain. Moving assets from Ethereum to an Arbitrum chain and back is done through a trustless bridge.

Currently Eth as well as any token confirming to ether the ERC-20 or ERC-721 standard can be transferred over to an Arbitrum chain.

By default all standard methods of transferring Eth or tokens will only adjust their balance inside the chain that the asset is currently on.

### Ethereum to Arbitrum

To move assets into an Arbitrum chain, you execute a deposit transaction on Arbitrum's global EthBridge. This transfers funds to the EthBridge on the Ethereum side, and credits the same funds to you inside the Arbitrum chain you specified. In the case of an ERC-20 or ERC-721 transfer, the Arbitrum chain will spawn a token contract of the correct type at the same token contract address as the token has on Ethereum.

As far as Ethereum knows, all deposited funds are held by Arbitrum's global EthBridge contract.

In order to programmatically trigger transfers, call one of the following methods in the [`GlobalPendingInbox`](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-eth/contracts/GlobalPendingInbox.sol) contract with the chain address equal to the address of the Rollup chain.

```javascript
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

### Arbitrum to Ethereum

-   Eth: The ArbSys library can be used to withdraw Eth, `ArbSys(100).withdrawEth(dest, amount)`
-   ERC-20 and ERC-721: The system generated token contracts in Arbitrum contain a withdraw method
    ```js
    function withdraw(address account, uint256 amount) public; // ERC-20
    function withdraw(address account, uint256 tokenId) public; // ERC-721
    ```

In all cases, withdrawing is similar to a transfer, except the balance is burned on the Arbitrum side, and eventually transfers those funds away from the Arbitrum chain to their destination on the Ethereum side.

When your withdraw transaction is fully confirmed, the withdrawn funds will be put into your "lockbox" in the EthBridge.
At any time you can call the EthBridge to recover the funds in your lockbox.

## Transaction calls on Arbitrum

All transaction calls on the Arbitrum chain are sent through the EthBridge using:

```js
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

## Transaction calls from Arbitrum to Ethereum

In our initial Rollup release, we are not supporting transaction calls from contracts in a Arbitrum chain to Ethereum contracts. In the future, support for this functionality will be added.
