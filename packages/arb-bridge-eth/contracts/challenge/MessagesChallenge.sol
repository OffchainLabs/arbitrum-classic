// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/Value.sol";
import "../inbox/Messages.sol";

contract MessagesChallenge is BisectionChallenge {
    using Hashing for Value.Data;

    event Bisected(
        bytes32[] chainHashes,
        bytes32[] segmentHashes,
        uint256 totalLength,
        uint256 deadlineTicks
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string private constant HS_BIS_INPLEN = "HS_BIS_INPLEN";

    function bisect(
        bytes32[] calldata _chainHashes,
        bytes32[] calldata _segmentHashes,
        uint256 _chainLength
    ) external asserterAction {
        uint256 bisectionCount = _chainHashes.length - 1;
        require(bisectionCount + 1 == _segmentHashes.length, HS_BIS_INPLEN);

        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                _chainHashes[0],
                _chainHashes[bisectionCount],
                _segmentHashes[0],
                _segmentHashes[bisectionCount],
                _chainLength
            )
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.messagesHash(
            _chainHashes[0],
            _chainHashes[1],
            _segmentHashes[0],
            _segmentHashes[1],
            firstSegmentSize(_chainLength, bisectionCount)
        );
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeUtils.messagesHash(
                _chainHashes[i],
                _chainHashes[i + 1],
                _segmentHashes[i],
                _segmentHashes[i + 1],
                otherSegmentSize(_chainLength, bisectionCount)
            );
        }

        commitToSegment(hashes);
        asserterResponded();
        emit Bisected(
            _chainHashes,
            _segmentHashes,
            _chainLength,
            deadlineTicks
        );
    }

    function oneStepProof(
        bytes32 _beforeInbox,
        bytes32 _preImageBHash,
        uint256 _preImageBSize,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes calldata _msgData
    ) external asserterAction {
        bytes32 afterInbox = Messages.addMessageToInbox(
            _beforeInbox,
            Messages.messageHash(
                _kind,
                _sender,
                _blockNumber,
                _timestamp,
                _inboxSeqNum,
                keccak256(_msgData)
            )
        );
        Value.Data memory messageValue = Messages.messageValue(
            _kind,
            _blockNumber,
            _timestamp,
            _sender,
            _inboxSeqNum,
            _msgData
        );
        Value.Data memory beforeVMInbox = Value.newTuplePreImage(
            _preImageBHash,
            _preImageBSize
        );
        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                _beforeInbox,
                afterInbox,
                beforeVMInbox.hash(),
                Messages
                    .addMessageToVMInbox(beforeVMInbox, messageValue)
                    .hash(),
                1
            )
        );

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
