// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

contract Registry {
    event ArbitrumSignalEmitted(
        address sender,
        address indexed l1Token,
        address indexed l2Token,
        address indexed l2OwnerAddress
    );

    function signalTokenAddressInArbitrum(
        address l1Token,
        address l2Token,
        address l2OwnerAddress
    ) external {
        emit ArbitrumSignalEmitted(msg.sender, l1Token, l2Token, l2OwnerAddress);
    }
}
