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

    uint8 internal constant SendType_sendTxToL1 = 3;

    address rollup;
    IBridge bridge;

    ICloneable outboxEntryTemplate;
    OutboxEntry[] public outboxes;

    // Note, these variables are set and then wiped during a single transaction.
    // Therefore their values don't need to be maintained, and their slots will
    // be empty outside of transactions
    address private _sender;
    uint128 private _l2Block;
    uint128 private _l1Block;
    uint128 private _timestamp;

    constructor(address _rollup, IBridge _bridge) public {
        rollup = _rollup;
        bridge = _bridge;
        outboxEntryTemplate = ICloneable(new OutboxEntry());
    }

    /// @notice When l2ToL1Sender returns a nonzero address, the message was originated by an L2 account
    /// When the return value is zero, that means this is a system message
    function l2ToL1Sender() external view override returns (address) {
        return _sender;
    }

    function l2ToL1Block() external view override returns (uint256) {
        return uint256(_l2Block);
    }

    function l2ToL1EthBlock() external view override returns (uint256) {
        return uint256(_l1Block);
    }

    function l2ToL1Timestamp() external view override returns (uint256) {
        return uint256(_timestamp);
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
            require(data.length == 97, "BAD_LENGTH");
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
        uint256 l1Block,
        uint256 l2Timestamp,
        uint256 amount,
        bytes calldata calldataForL1
    ) external {
        bytes32 userTx = calculateItemHash(
            l2Sender,
            destAddr,
            l2Block,
            l1Block,
            l2Timestamp,
            amount,
            calldataForL1
        );

        spendOutput(outboxIndex, proof, index, userTx);

        address currentSender = _sender;
        uint128 currentL2Block = _l2Block;
        uint128 currentL1Block = _l1Block;
        uint128 currentTimestamp = _timestamp;

        _sender = l2Sender;
        _l2Block = uint128(l2Block);
        _l1Block = uint128(l1Block);
        _timestamp = uint128(l2Timestamp);

        executeBridgeCall(destAddr, amount, calldataForL1);

        _sender = currentSender;
        _l2Block = currentL2Block;
        _l1Block = currentL1Block;
        _timestamp = currentTimestamp;
    }

    function spendOutput(
        uint256 outboxIndex,
        bytes32[] memory proof,
        uint256 path,
        bytes32 item
    ) private {
        require(proof.length <= 256, "PROOF_TOO_LONG");
        require(path < 2**proof.length, "PATH_NOT_MINIMAL");

        // Hash the leaf an extra time to prove it's a leaf
        bytes32 calcRoot = calculateMerkleRoot(proof, path, item);
        OutboxEntry outbox = outboxes[outboxIndex];
        require(address(outbox) != address(0), "NO_OUTBOX");

        // With a minimal path, the pair of path and proof length should always identify
        // a unique leaf. The path itself is not enough since the path length to different
        // leaves could potentially be different
        bytes32 uniqueKey = keccak256(abi.encodePacked(path, proof.length));

        executeBridgeSystemCall(
            address(outbox),
            0,
            abi.encodeWithSelector(OutboxEntry.spendOutput.selector, calcRoot, uniqueKey)
        );

        if (outbox.numRemaining() == 0) {
            executeBridgeSystemCall(
                address(outbox),
                0,
                abi.encodeWithSelector(OutboxEntry.destroy.selector)
            );
            outboxes[outboxIndex] = OutboxEntry(address(0));
        }
    }

    function executeBridgeSystemCall(
        address destAddr,
        uint256 amount,
        bytes memory data
    ) private {
        address currentSender = _sender;
        _sender = address(0);
        executeBridgeCall(destAddr, amount, data);
        _sender = currentSender;
    }

    function executeBridgeCall(
        address destAddr,
        uint256 amount,
        bytes memory data
    ) private {
        (bool success, bytes memory returndata) = bridge.executeCall(destAddr, amount, data);
        if (!success) {
            if (returndata.length > 0) {
                // solhint-disable-next-line no-inline-assembly
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert("BRIDGE_CALL_FAILED");
            }
        }
    }

    function calculateItemHash(
        address l2Sender,
        address destAddr,
        uint256 l2Block,
        uint256 l1Block,
        uint256 l2Timestamp,
        uint256 amount,
        bytes calldata calldataForL1
    ) public pure returns (bytes32) {
        return keccak256(
            abi.encodePacked(
                SendType_sendTxToL1,
                uint256(uint160(bytes20(l2Sender))),
                uint256(uint160(bytes20(destAddr))),
                l2Block,
                l1Block,
                l2Timestamp,
                amount,
                calldataForL1
            )
        );
    }

    function calculateMerkleRoot(
        bytes32[] memory proof,
        uint256 path,
        bytes32 item
    ) public pure returns (bytes32) {
        return MerkleLib.calculateRoot(proof, path, keccak256(abi.encodePacked(item)));
    }

    function outboxesLength() public view returns (uint256) {
        return outboxes.length;
    }
}
