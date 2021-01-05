// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.5.11;

import "./Messages.sol";
import "../libraries/MerkleLib.sol";

contract Outbox {
    OutboxEntry[] outboxes;

    function addOutbox() internal {}

    function executeTransaction(
        uint256 outboxIndex,
        bytes calldata _proof,
        uint256 _index,
        address destAddr,
        uint256 amount,
        bytes calldata calldataForL1
    ) external {
        bytes32 userTx = keccak256(
            abi.encodePacked(uint256(uint160(bytes20(destAddr))), amount, calldataForL1)
        );

        outboxes[outboxIndex].spendOutput(_proof, _index, userTx);

        (bool success, ) = destAddr.call.value(amount)(calldataForL1);
        require(success);
    }
}

contract OutboxEntry {
    bytes32 outputRoot;
    mapping(uint256 => bool) spentOutput;

    constructor(bytes32 root) public {
        outputRoot = root;
    }

    function spendOutput(
        bytes calldata proof,
        uint256 index,
        bytes32 item
    ) external {
        // TODO: Verify that this is actually a leaf and not an intermediate node
        // One way to do this would be to include the tree depth in the original output message
        (bytes32 calcRoot, ) = MerkleLib.verifyMerkleProof(proof, item, index + 1);
        require(calcRoot == outputRoot);
        spentOutput[index] = true;
    }
}
