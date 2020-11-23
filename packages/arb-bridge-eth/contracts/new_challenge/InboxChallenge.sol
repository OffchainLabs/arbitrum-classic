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

import "./BisectionChallenge.sol";
import "../challenge/ChallengeUtils.sol";

import "../inbox/Messages.sol";

import "../libraries/MerkleLib.sol";

contract InboxTopChallenge is BisectionChallenge {
    event Bisected(bytes32[] chainHashes, uint256 totalLength, uint256 deadlineTicks);

    event OneStepProofCompleted();

    // Incorrect previous state
    string private constant BIS_PREV = "BIS_PREV";

    // Proof was incorrect
    string private constant HC_OSP_PROOF = "HC_OSP_PROOF";

    // Invalid assertion selected
    string private constant CON_PROOF = "CON_PROOF";

    function oneStepProof(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _lowerHash,
        bytes32 _value
    ) external onlyOnTurn {
        bytes32 prevHash = ChallengeUtils.inboxTopHash(
            _lowerHash,
            Messages.addMessageToInbox(_lowerHash, _value),
            1
        );
        require(
            MerkleLib.verifyProof(_proof, challengeState, prevHash, _segmentToChallenge + 1),
            CON_PROOF
        );

        emit OneStepProofCompleted();
        _asserterWin();
    }
}
