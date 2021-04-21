---
id: Bridging_Assets
title: Token Bridging
sidebar_label: Token Bridging
---

The Arbitrum protocol's ability to [pass messages between L1 and L2](L1_L2_Messages.md) can be leveraged to trustlessly move assets from Ethereum to an Arbitrum chain and back. Any asset / asset type can in principle be bridged, including Ether, ERC20 tokens, ERC-721 tokens, etc.

## Depositing And Withdrawing Ether

To move Ether from Ethereum onto the Arbitrum chain, you execute a deposit transaction via `Inbox.depositEth`. This transfers funds to the Bridge contract on the L1 and credits the same funds to you inside the Arbitrum chain at the specified address.

```sol
function depositEth(address destAddr) external payable override returns (uint256)
```

As far as Ethereum knows, all deposited funds are held by Arbitrum's Bridge contract.

Withdrawing ether can be done using the [ArbSys](ArbSys.md) withdrawEth method:

```sol
ArbSys(100).withdrawEth{ value: 2300000 }(destAddress)
```

Upon withdrawing, the Ether balance is burnt on the Arbitrum side, and will later be made available on the Ethereum side.

`ArbSys.withdrawEth` is actually a convenience function is which is equivalent to calling `ArbSys.sendTxToL1` with empty calldataForL1. Like any other `sendTxToL1` call, it will require an additional call to `Outbox.executeTransaction` on L1 after the dispute period elapses for the user to finalize claiming their funds on L1 (see ["L2 to L1 Messages Lifecycle && API"](L1_L2_Messages.md)). Once the withdrawal is executed from the Outbox, the user's Ether balance will be credited on L1.

## Bridging ERC20 Tokens

### Overview

The Arbitrum protocol itself technically has no native notion of any token standards, and gives no built-in advantage or special recognition to any particular token bridge. Described here is the "Canonical Bridge," which we at Offchain Labs implemented, and which should be the primary bridge most users and applications use; it is (effectively) a DApp with contracts on both Ethereum and Arbitrum that leverages Arbitrum's cross-chain message passing system to achieve basic desired token-bridging functionality. We recommend that you use it!

"Basic desired token bridging functionality" for most ERC20 tokens is the following: a token contract on Ethereum is associated with a "paired" token contract on Arbitrum. Depositing a token entails escrowing some amount of the token in an L1 bridge contract, and minting the same amount at the paired token contract on L2. On L2, the paired contract behaves much like a normal ERC20 token contract. Withdrawing entails burning some amount of the token in the L2 contract, which then can later be claimed from the L1 bridge contract.

### Canonical Token Bridge Implementation

The two main bridging contracts are EthERC20Bridge.sol (L1 side) and ArbTokenBridge.sol (L2 side). The system offers two options for creating an L2 pairing for a token contract: the standard ERC20 option and the "custom token" option. Custom tokens should only be used if you have a good reason to.

#### Standard Arb-ERC20 Bridging

Any ERC20 token on Ethereum can be bridged onto Arbitrum "the standard way" by simply calling `EthERC20Bridge.deposit` with the token's L1 address:

```sol
    /**
     * @notice Deposit standard or custom ERC20 token. If L2 side hasn't been deployed yet, includes name/symbol/decimals data for initial L2 deploy.
     * @param erc20 L1 address of ERC20
     * @param destination account to be credited with the tokens in the L2 (can be the user's L2 account or a contract)
     * @param amount Token Amount
     * @param maxSubmissionCost Max gas deducted from user's L2 balance to cover base submission fee
     * @param maxGas Max gas deducted from user's L2 balance to cover L2 execution
     * @param gasPriceBid Gas price for L2 execution
     * @param callHookData optional data for external call upon minting
     * @return ticket ID used to redeem the retryable transaction in the L2
     */
    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        bytes calldata callHookData
    ) external payable override returns (uint256)
```

If no paired-L2 has yet been deployed (i.e., if this is the first deposit), a [StandardArbERC20](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-bridge-peripherals/contracts/tokenbridge/arbitrum/StandardArbERC20.sol) will be deployed on Arbitrum at a deterministically generated address. StandardArbErc20 implements the ERC20 standard with additional bridging-related methods (i.e., `bridgeMint, withdraw`, etc.) as well as extensions to improve UX (ERC1363 & ERC2612). The L1 contracts name, symbol, and decimals will also be pushed directly into the new StandardArbERC20.

To withdraw from a standard ERC20 from Arbitrum , call `SomeStandardArbERC20.withdraw`

```sol
    /**
     * @notice Initiates a token withdrawal
     * @param account destination address
     * @param amount amount of tokens withdrawn
     */
    function withdraw(address account, uint256 amount) external override {
```

This initiates the withdrawal process (i.e., tokens can later be claimed via the Outbox.)

#### Custom Token Bridging

The Canonical Token Bridge also allows for pairing an ERC20 on L1 with a custom token contract on Arbitrum.

**It is important to note**: The Canonical Token Bridge can't account for every possible token contract; if your custom token logic is [sufficiently weird](https://quoteinvestigator.com/2018/12/25/universe/#:~:text=Professor%20J.%20B.%20S.%20Haldane%20once%20shrewdly,the%20ultimate%20queerness%20of%20time.), (i.e., rebasing stablecoins, passively interest accruing tokens, etc.) it may require it's own, custom, tailor-made bridge as well. The implications of a particular custom token should be thought through carefully; feel free to [reach out to us](https://discord.gg/ZpZuw7p). With that said, the steps for creating a custom token pairing in the canonical bridge are as follows.

#### Setting up a custom token pairing

**1. Deploy L2 custom token contract**

Start by deploying the L2 token contract [directly onto Arbitrum](Contract_Deployment.md). A custom L2 contract should conform to the minimal IArbToken interface:

```sol
/**
 * @title Minimum expected interface for L2 token that interacts with the L2 token bridge (this is the interface necessary
 * for a custom token that interacts with the bridge, see TestArbCustomToken.sol for an example implementation).
 */
interface IArbToken {
    /**
     * @notice should increase token supply by amount, and should (probably) only be callable by the L1 bridge.
     */
    function bridgeMint(address account, uint256 amount) external;

    /**
     * @notice should decrease token supply by amount, and should (probably) only be callable by the L1 bridge.
     */
    function bridgeBurn(address account, uint256 amount) external;

    /**
     * @notice withdraw user tokens from L2 to the L1
     */
    function withdraw(address account, uint256 amount) external;
}
```

**2. Deploy L1 custom token contract**

The L1 contract should conform to the minimal `ICustomToken` interface:

```sol
/**
 * @title Minimum expected interface for L1 custom token (see TestCustomTokenL1.sol for an example implementation)
 */
interface ICustomToken {
    /**
     * @notice Should make an external call to EthERC20Bridge.registerCustomL2Token
     */
    function registerTokenOnL2(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) external virtual;

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external virtual returns (bool);

    function balanceOf(address account) external view virtual returns (uint256);
}
```

**3. Register L1/L2 custom pairing**

Calling `MyCustomL1Token.registerTokenOnL2` should make an external call to `EthERC20Bridge.registerCustomL2Token` with the custom L2 token's address (see [TestCustomTokenL1](https://github.com/OffchainLabs/arbitrum/blob/develop/packages/arb-bridge-peripherals/contracts/tokenbridge/test/TestCustomTokenL1.sol) for an example implementation). Once registered, calling `EthERC20Bridge.deposit` with a custom token's L1 address with deposit into the corresponding custom L2 pairing.

It should be noted that if no token contract is actually registered at the L2 address given, a temporary standard token contract will be deployed, ensuring any deposits are safely recoverable. However, we highly recommend you simply save yourself such hassle and simply follow the steps outlined here 😉.

Note that [arb-ts](https://github.com/OffchainLabs/arbitrum/tree/master/packages/arb-ts) provides client side convenience methods for the functionality listed above, and more.

See [integration tests](https://github.com/OffchainLabs/arbitrum/blob/master/packages/arb-ts/integration_test/arb-bridge.test.ts) for example usage.

## Arbitrum-Native ERC20 Tokens

It is (of course) possible to deploy an ERC20 token contract directly to Arbitrum, i.e., with no layer 1 counterpart. Such a token functions normally within Arbitrum, but simply can't be withdrawn to layer 1. (In principle, enabling L2 native tokens to be withdrawn to an L1 contract could be possible via a similar mechanism our Canonical Token Bridge uses with the layers flipped, i.e., an "anti-bridge," but such functionality isn't currently supported.)
