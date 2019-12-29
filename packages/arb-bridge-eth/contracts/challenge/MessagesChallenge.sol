/*
 * Copyright 2019, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./BisectionChallenge.sol";
import "./ChallengeUtils.sol";

import "../arch/Protocol.sol";


contract MessagesChallenge is BisectionChallenge {

    event Bisected(
        bytes32[] chainHashes,
        uint32 totalLength,
        uint64 deadline
    );

    event OneStepProofCompleted();

    // Incorrect previous state
    string constant HS_BIS_INPLEN = "HS_BIS_INPLEN";
    // Proof was incorrect
    string constant HS_OSP_PROOF = "HS_OSP_PROOF";

    function bisect(
        bytes32[] memory _chainHashes,
        bytes32[] memory _segmentHashes,
        uint32 _chainLength
    )
        public
        asserterAction
    {
        uint bisectionCount = _chainHashes.length - 1;
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
        for (uint i = 1; i < bisectionCount; i++) {
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
            _chainLength,
            deadline
        );
    }

    function oneStepProof(
        bytes32 _lowerHashA,
        bytes32 _topHashA,
        bytes32 _lowerHashB,
        bytes32 _topHashB,
        bytes32 _value
    )
        public
        asserterAction
    {
        requireMatchesPrevState(
            ChallengeUtils.messagesHash(
                _lowerHashA,
                _topHashA,
                _lowerHashB,
                _topHashB,
                uint32(1)
            )
        );

        require(Protocol.addMessageToPending(_lowerHashA, _value) == _topHashA, HS_OSP_PROOF);
        require(Protocol.addMessageToPending(_lowerHashB, _value) == _topHashB, HS_OSP_PROOF);

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
