---
id: Withdrawals
title: L1 Finality and Fast Withdrawals
sidebar_label: Withdrawals
---

As discussed in the section on [finality](Finality.md), because Arbitrum's execution happens optimistically, the Ethereum blockchain cannot immediately confirm the correct state and must wait for the challenge window to expire (or until all challenges are resolved).

When it comes to execution within the Rollup this does not pose any problem or add any delay. The [Arbitrum Rollup protocol](Rollup_Protocol.md) manages a tree of assertions, and allows validators to pipeline execution by continuing to build the tree even before all nodes are confirmed. This means that an honest validator can continue to advance the state of the machine with confidence (and the ability to enforce) that eventually Ethereum will recognize the honest branch as the correct and valid one. And although it will take some time for Ethereum to recognize which branch is correct, anyone that is validating the chain will immediately know this. Anyone that is validating the chain will immediately know which branch is correct and therefore which branch will eventually be accepted by the protocol.

The one part of the protocol that is affected by the confirmation delay are [L2 to L1 messages](L1_L2_Messages.md), notably, withdrawals. Since Arbitrum cannot undo a withdrawal once it has released funds from the L2, the system cannot allow funds to be withdrawn from the ArbChain until it has been confirmed on the Ethereum chain that the withdrawal is valid.

## Liquidity Exits ("Fast Withdrawals")

When withdrawing Ether or an ERC20 token, one can take advantage of the asset's fungibility to sidestep the confirmation delay entirely. When using these fast,"liquidity exit" techniques, a user with funds on Arbitrum simply pays some un-trusted third party to pay them back directly on L1 (presumably for some small fee). Broadly speaking, this exchange can be carried out two different ways: by withdrawing via a bridge and trading the in-flight exit, or by atomic-swapping assets across the L1 and L2 chains (without any direct "withdrawal" actually taking place).

### Tradeable Bridge Exits

To carry out a tradable exit, a user first initializes a withdrawal; a third party — the liquidity provider — can immediately verify that the withdrawal is valid (i.e., that it will finalize _eventually_) by validating the Arbitrum chain. The liquidity provider then offers to buy the exit by paying the user on L1.

Our [token bridge](Bridging_Assets.md) includes an implementation of this technical

```sol /**
     * @notice Allows a user to redirect their right to claim a withdrawal to a liquidityProvider, in exchange for a fee.
     * @dev This method expects the liquidityProvider to verify the liquidityProof, but it ensures the withdrawer's balance
     * is appropriately updated. It is otherwise agnostic to the details of IExitLiquidityProvider.requestLiquidity.
     * @param liquidityProvider address of an IExitLiquidityProvider
     * @param liquidityProof encoded data required by the liquidityProvider in order to validate a fast withdrawal.
     * @param initialDestination address the L2 withdrawal call initially set as the destination.
     * @param erc20 L1 token address
     * @param amount token amount (should match amount in previously-initiated withdrawal)
     * @param exitNum Sequentially increasing exit counter determined by the L2 bridge
     * @param maxFee max mount of erc20 token user will pay for fast exit
     */
    function fastWithdrawalFromL2(
        address liquidityProvider,
        bytes memory liquidityProof,
        address initialDestination,
        address erc20,
        uint256 amount,
        uint256 exitNum,
        uint256 maxFee
    ) external override
```

#### Atomic Swaps / HTLCs

To carry out fast withdrawal via an atomic swap, an Arbitrum user who wants to "withdraw" onto L1 pays a liquidity provider directly on L2, who in turn transfers funds to the user's address on L1. [Hashed time locked contracts (HTLCs)](https://www.investopedia.com/terms/h/hashed-timelock-contract.asp) are used to ensure that these two operations are ultimately atomic; i.e., either both take place or neither of them do, preserving trustlessness.

Variants of both of these approaches can also be extended to provide fast transfers between multiple L2 chains (i.e., 2 different Arbitrum chains.)

## NFTs and Messages

Liquidity exits work well for fungible tokens. However for non-fungible tokens for which a liquidity provider cannot lend an equivalent substitute, withdrawing will still incur the delay for system confirmation. Similarly, when one wants the ArbChain to post a message to the L1 (e.g. that will be processed as an asynchronous call by another contract), they will also incur the confirmation delay.
