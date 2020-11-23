// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

import "../challenge/Challenge.sol";
import "../challenge/ChallengeUtils.sol";

import "../inbox/Messages.sol";

import "../libraries/MerkleLib.sol";

contract InboxTopChallenge is Challenge {
    event Bisected(bytes32[] chainHashes, uint256 totalLength, uint256 deadlineTicks);

    event OneStepProofCompleted();

    // Incorrect previous state
    string private constant BIS_PREV = "BIS_PREV";

    // Proof was incorrect
    string private constant HC_OSP_PROOF = "HC_OSP_PROOF";

    // Invalid assertion selected
    string private constant CON_PROOF = "CON_PROOF";

    bytes32 private challengeState;

    function bisect(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength
    ) external asserterOrChallengerAction {
        uint256 bisectionCount = _chainHashes.length - 1;

        require(_chainHashes[bisectionCount] != _oldEndHash);
        require(_chainLength > 1, "bisection too short");

        bytes32 bisectionHash = ChallengeUtils.inboxTopHash(
            _chainHashes[0],
            _oldEndHash,
            _chainLength
        );

        require(
            MerkleLib.verifyProof(_proof, challengeState, bisectionHash, _segmentToChallenge + 1),
            CON_PROOF
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeUtils.inboxTopHash(
            _chainHashes[0],
            _chainHashes[1],
            ChallengeUtils.firstSegmentSize(_chainLength, bisectionCount)
        );
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeUtils.inboxTopHash(
                _chainHashes[i],
                _chainHashes[i + 1],
                ChallengeUtils.otherSegmentSize(_chainLength, bisectionCount)
            );
        }
        challengeState = MerkleLib.generateRoot(hashes);
        responded();
        emit Bisected(_chainHashes, _chainLength, deadlineTicks);
    }

    function oneStepProof(bytes32 _lowerHash, bytes32 _value) external asserterAction {
        bytes32 prevHash = ChallengeUtils.inboxTopHash(
            _lowerHash,
            Messages.addMessageToInbox(_lowerHash, _value),
            1
        );
        require(challengeState == prevHash, BIS_PREV);

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
