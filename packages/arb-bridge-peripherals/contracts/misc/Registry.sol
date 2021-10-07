// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

contract Registry {
    event SignalEmitted(
        address sender,
        address indexed l1Address,
        address indexed l2Address,
        address indexed l2OwnerAddress
    );

    function signal(
        address l1Address,
        address l2Address,
        address l2OwnerAddress
    ) external {
        emit SignalEmitted(msg.sender, l1Address, l2Address, l2OwnerAddress);
    }
}
