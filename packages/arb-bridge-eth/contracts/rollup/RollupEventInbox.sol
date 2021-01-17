// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

import "../bridge/interfaces/IBridge.sol";

contract RollupEventInbox {
    uint8 internal constant ROLLUP_PROTOCOL_EVENT_TYPE = 8;

    uint8 internal constant CREATE_NODE_EVENT = 0;
    uint8 internal constant CONFIRM_NODE_EVENT = 0;
    uint8 internal constant REJECT_NODE_EVENT = 0;
    uint8 internal constant STAKE_CREATED_EVENT = 0;

    IBridge bridge;

    function nodeCreated(
        uint256 nodeNum,
        uint256 prev,
        uint256 deadline,
        address asserter
    ) public {
        bytes32 messageHash =
            keccak256(
                abi.encodePacked(
                    CREATE_NODE_EVENT,
                    nodeNum,
                    prev,
                    block.number,
                    deadline,
                    uint256(uint160(bytes20(asserter)))
                )
            );
        deliverToBridge(messageHash);
    }

    function nodeConfirmed(uint256 nodeNum) public {
        bytes32 messageHash = keccak256(abi.encodePacked(CONFIRM_NODE_EVENT, nodeNum));
        deliverToBridge(messageHash);
    }

    function nodeRejected(uint256 nodeNum) public {
        bytes32 messageHash = keccak256(abi.encodePacked(REJECT_NODE_EVENT, nodeNum));
        deliverToBridge(messageHash);
    }

    function stakeCreated(address staker, uint256 nodeNum) public {
        bytes32 messageHash =
            keccak256(
                abi.encodePacked(
                    STAKE_CREATED_EVENT,
                    uint256(uint160(bytes20(staker))),
                    nodeNum,
                    block.number
                )
            );
        deliverToBridge(messageHash);
    }

    function deliverToBridge(bytes32 messageHash) private returns (uint256) {
        return bridge.deliverMessageToInbox(ROLLUP_PROTOCOL_EVENT_TYPE, msg.sender, messageHash);
    }
}
