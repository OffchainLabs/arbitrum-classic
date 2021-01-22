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

pragma solidity ^0.6.11;

import "../libraries/CloneFactory.sol";
import "./OutboxEntry.sol";

import "./interfaces/IOutbox.sol";
import "./interfaces/IBridge.sol";

import "./Messages.sol";
import "../libraries/MerkleLib.sol";
import "../libraries/BytesLib.sol";

contract Outbox is CloneFactory, IOutbox {
    using BytesLib for bytes;

    bytes1 internal constant MSG_ROOT = 0;

    uint256 internal constant SendType_sendTxToL1 = 0;
    uint256 internal constant SendType_buddyContractResult = 5;

    address rollup;
    IBridge bridge;

    ICloneable outboxEntryTemplate;
    OutboxEntry[] outboxes;

    // Note, these variables are set and then wiped during a single transaction.
    // Therefore their values don't need to be maintained, and their slots will
    // be empty outside of transactions
    address private _l2ToL1Sender;
    uint128 private _l2ToL1Block;
    uint128 private _l2ToL1Timestamp;

    constructor(address _rollup, IBridge _bridge) public {
        rollup = _rollup;
        bridge = _bridge;
        outboxEntryTemplate = ICloneable(new OutboxEntry());
    }

    /// @notice When l2ToL1Sender returns a nonzero address, the message was originated by an L2 account
    /// When the return value is zero, that means this is a system message
    function l2ToL1Sender() external view override returns (address) {
        return _l2ToL1Sender;
    }

    function l2ToL1Block() external view override returns (uint256) {
        return uint256(_l2ToL1Sender);
    }

    function l2ToL1Timestamp() external view override returns (uint256) {
        return uint256(_l2ToL1Sender);
    }

    function processOutgoingMessages(bytes calldata sendsData, uint256[] calldata sendLengths)
        external
        override
    {
        require(msg.sender == rollup, "ONLY_ROLLUP");
        // If we've reached here, we've already confirmed that sum(sendLengths) == sendsData.length
        uint256 messageCount = sendLengths.length;
        uint256 offset = 0;
        for (uint256 i = 0; i < messageCount; i++) {
            handleOutgoingMessage(bytes(sendsData[offset:offset + sendLengths[i]]));
            offset += sendLengths[i];
        }
    }

    function handleOutgoingMessage(bytes memory data) private {
        // Otherwise we have an unsupported message type and we skip the message
        if (data[0] == MSG_ROOT) {
            uint256 batchNum = data.toUint(1);
            uint256 numInBatch = data.toUint(33);
            bytes32 outputRoot = data.toBytes32(65);

            address clone = createClone(outboxEntryTemplate);
            OutboxEntry(clone).initialize(bridge, outputRoot, numInBatch);
            uint256 outboxIndex = outboxes.length;
            outboxes.push(OutboxEntry(clone));
            emit OutboxEntryCreated(batchNum, outboxIndex, outputRoot, numInBatch);
        }
    }

    function executeTransaction(
        uint256 outboxIndex,
        bytes32[] calldata proof,
        uint256 index,
        address l2Sender,
        address destAddr,
        uint256 l2Block,
        uint256 l2Timestamp,
        uint256 amount,
        bytes calldata calldataForL1
    ) external {
        bytes32 userTx =
            keccak256(
                abi.encodePacked(
                    SendType_sendTxToL1,
                    uint256(uint160(bytes20(l2Sender))),
                    uint256(uint160(bytes20(destAddr))),
                    l2Block,
                    l2Timestamp,
                    amount,
                    calldataForL1
                )
            );

        spendOutput(outboxIndex, proof, index, userTx);

        address currentL2Sender = _l2ToL1Sender;
        uint128 currentL2Block = _l2ToL1Block;
        uint128 currentL2Timestamp = _l2ToL1Timestamp;

        _l2ToL1Sender = l2Sender;
        _l2ToL1Block = uint128(l2Block);
        _l2ToL1Timestamp = uint128(l2Timestamp);

        (bool success, ) = bridge.executeCall(destAddr, amount, calldataForL1);
        require(success, "CALL_FAILED");

        _l2ToL1Sender = currentL2Sender;
        _l2ToL1Block = currentL2Block;
        _l2ToL1Timestamp = currentL2Timestamp;
    }

    function executeBuddyContractReceipt(
        uint256 outboxIndex,
        bytes32[] calldata proof,
        uint256 index,
        address l2Contract,
        bool createdSuccessfully
    ) external {
        bytes32 userTx =
            keccak256(
                abi.encodePacked(
                    SendType_buddyContractResult,
                    uint256(uint160(bytes20(l2Contract))),
                    createdSuccessfully
                )
            );

        spendOutput(outboxIndex, proof, index, userTx);

        address currentL2Sender = _l2ToL1Sender;
        _l2ToL1Sender = address(0);

        (bool success, ) =
            bridge.executeCall(
                l2Contract,
                0,
                abi.encodeWithSignature("buddyContractResult(bool)", createdSuccessfully)
            );
        require(success, "CALL_FAILED");

        _l2ToL1Sender = currentL2Sender;
    }

    function spendOutput(
        uint256 outboxIndex,
        bytes32[] memory proof,
        uint256 path,
        bytes32 item
    ) private {
        // Hash the leaf an extra time to prove it's a leaf
        bytes32 calcRoot = MerkleLib.calculateRoot(proof, path, keccak256(abi.encodePacked(item)));
        address currentL2Sender = _l2ToL1Sender;
        _l2ToL1Sender = address(0);
        (bool success, ) =
            bridge.executeCall(
                address(outboxes[outboxIndex]),
                0,
                abi.encodeWithSignature("spendOutput(bytes32,uint256)", calcRoot, path)
            );
        _l2ToL1Sender = currentL2Sender;
        require(success, "CANT_SPEND");
    }
}
