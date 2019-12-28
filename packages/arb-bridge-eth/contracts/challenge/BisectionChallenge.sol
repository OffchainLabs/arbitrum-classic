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

import "./Challenge.sol";

import "../libraries/MerkleLib.sol";


contract BisectionChallenge is Challenge {

    event Continued (
        uint segmentIndex,
        uint64 deadline
    );

    // Incorrect previous state
    string constant BIS_PREV = "BIS_PREV";

    // Incorrect previous state
    string constant CON_PREV = "CON_PREV";
    // Invalid assertion selected
    string constant CON_PROOF = "CON_PROOF";
    // Incorrect previous state

    // After bisection this is an array of all sub-assertions
    // After a challenge, the first assertion is the main assertion
    bytes32 private challengeState;

    function initializeBisection(
        address _vmAddress,
        address payable _asserter,
        address payable _challenger,
        uint32 _challengePeriod,
        bytes32 _challengeState
    )
        internal
    {
        Challenge.initializeChallenge(
            _vmAddress,
            _asserter,
            _challenger,
            _challengePeriod
        );
        challengeState = _challengeState;
    }

    function chooseSegment(
        uint _segmentToChallenge,
        bytes memory _proof,
        bytes32 _bisectionRoot,
        bytes32 _bisectionHash
    )
        public
        challengerAction
    {
        require(_bisectionRoot == challengeState, CON_PREV);
        require(
            MerkleLib.verifyProof(
                _proof,
                _bisectionRoot,
                _bisectionHash,
                _segmentToChallenge + 1
            ),
            CON_PROOF
        );

        challengeState = _bisectionHash;

        challengerResponded();
        emit Continued(_segmentToChallenge, deadline);
    }

    function commitToSegment(bytes32[] memory hashes) internal {
        challengeState = MerkleLib.generateRoot(hashes);
    }

    function requireMatchesPrevState(bytes32 _challengeState) internal view {
        require(_challengeState == challengeState, BIS_PREV);
    }

    function firstSegmentSize(uint32 totalCount, uint bisectionCount) internal pure returns(uint32) {
        return totalCount / uint32(bisectionCount) + totalCount%uint32(bisectionCount);
    }

    function otherSegmentSize(uint32 totalCount, uint bisectionCount) internal pure returns(uint32) {
        return totalCount / uint32(bisectionCount);
    }
}
