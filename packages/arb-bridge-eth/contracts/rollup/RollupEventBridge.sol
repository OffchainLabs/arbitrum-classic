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

import "./Rollup.sol";

import "../bridge/interfaces/IBridge.sol";
import "./INode.sol";

contract RollupEventBridge {
    uint8 internal constant INITIALIZATION_MSG_TYPE = 4;
    uint8 internal constant ROLLUP_PROTOCOL_EVENT_TYPE = 8;

    uint8 internal constant CREATE_NODE_EVENT = 0;
    uint8 internal constant CONFIRM_NODE_EVENT = 1;
    uint8 internal constant REJECT_NODE_EVENT = 2;
    uint8 internal constant STAKE_CREATED_EVENT = 3;
    uint8 internal constant CLAIM_NODE_EVENT = 4;

    IBridge bridge;
    address rollup;

    modifier onlyRollup {
        require(msg.sender == rollup, "ONLY_ROLLUP");
        _;
    }

    constructor(address _bridge, address _rollup) public {
        bridge = IBridge(_bridge);
        rollup = _rollup;
    }

    function rollupInitialized(
        uint256 confirmPeriodBlocks,
        uint256 extraChallengeTimeBlocks,
        uint256 arbGasSpeedLimitPerBlock,
        uint256 baseStake,
        address stakeToken,
        address owner,
        bytes calldata extraConfig
    ) external onlyRollup {
        bytes32 initMsgHash =
            keccak256(
                abi.encodePacked(
                    confirmPeriodBlocks,
                    extraChallengeTimeBlocks,
                    arbGasSpeedLimitPerBlock,
                    baseStake,
                    uint256(uint160(bytes20(stakeToken))),
                    uint256(uint160(bytes20(owner))),
                    extraConfig
                )
            );
        bridge.deliverMessageToInbox(INITIALIZATION_MSG_TYPE, msg.sender, initMsgHash);
    }

    function nodeCreated(
        uint256 nodeNum,
        uint256 prev,
        uint256 deadline,
        address asserter
    ) external onlyRollup {
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

    function nodeConfirmed(uint256 nodeNum) external onlyRollup {
        bytes32 messageHash = keccak256(abi.encodePacked(CONFIRM_NODE_EVENT, nodeNum));
        deliverToBridge(messageHash);
    }

    function nodeRejected(uint256 nodeNum) external onlyRollup {
        bytes32 messageHash = keccak256(abi.encodePacked(REJECT_NODE_EVENT, nodeNum));
        deliverToBridge(messageHash);
    }

    function stakeCreated(address staker, uint256 nodeNum) external onlyRollup {
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

    function claimNode(uint256 nodeNum, address staker) external onlyRollup {
        Rollup r = Rollup(rollup);
        INode node = r.getNode(nodeNum);
        require(node.stakers(staker), "NOT_STAKED");
        r.requireUnresolved(nodeNum);

        bytes32 messageHash =
            keccak256(
                abi.encodePacked(CLAIM_NODE_EVENT, nodeNum, uint256(uint160(bytes20(staker))))
            );
        deliverToBridge(messageHash);
    }

    function deliverToBridge(bytes32 messageHash) private returns (uint256) {
        return bridge.deliverMessageToInbox(ROLLUP_PROTOCOL_EVENT_TYPE, msg.sender, messageHash);
    }
}
