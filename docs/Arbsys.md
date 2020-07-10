---
id: Arbsys
title: The ArbSys Precompiled Contract
sidebar_label: ArbSys Precompiled Contract
---

ArbSys is a precompiled contract that exists in every Arbitrum Chain.
As its name would imply, ArbSys provides systems functionality useful to some Arbitrum contracts.
Any contract running on an Arbitrum Chain can call the chain's ArbSys.

ArbSys lives at address 100 on every Arbitrum chain.
To call it, write something like this:

    uint256 txCount = ArbSys(address(100)).getTransactionCount();

Here is the interface offered by ArbSys:

    interface ArbSys {
        // Send given amount of ERC-20 tokens to dest with token contract sender.
        // This is safe to freely call since the sender is authenticated and thus
        // you can only send fake tokens, not steal real ones
        function withdrawERC20(address dest, uint256 amount) external;

        // Send given ERC-721 token to dest with token contract sender.
        // This is safe by the above arguement
        function withdrawERC721(address dest, uint256 id) external;

        // Send given amount of Eth to dest with from sender.
        function withdrawEth(address dest, uint256 amount) external;

        // Return the number of transactions issued by the given external account
        // or the account sequence number of the given contract
        function getTransactionCount(address account) external view returns(uint256);

        // Generate a new contract with the same code as the given contract
        // This function returns the address of the new contract
        // This is currently the only way to create new contracts in a compiled rollup instance
        function cloneContract(address account) external returns(address);
    }
