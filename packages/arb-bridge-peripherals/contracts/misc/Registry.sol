// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

contract Registry {
    event SignalEmitted(address indexed from, address indexed signal);

    function signal(address signalAddr) external {
        emit SignalEmitted(msg.sender, signalAddr);
    }
}
