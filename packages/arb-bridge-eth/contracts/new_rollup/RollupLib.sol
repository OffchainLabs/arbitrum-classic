// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.5.17;

library RollupLib {
    function nodeStateHash(
        bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount,
        uint256 messageCount,
        uint256 logCount
    ) internal pure returns (bytes32) {
        return
            keccak256(abi.encodePacked(machineHash, inboxTop, inboxCount, messageCount, logCount));
    }
}
