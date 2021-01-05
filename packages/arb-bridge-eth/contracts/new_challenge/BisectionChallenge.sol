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

import "../libraries/Cloneable.sol";
import "../libraries/MerkleLib.sol";
import "../libraries/RollupTime.sol";
import "../rollup/IStaking.sol";
import "./ChallengeLib.sol";

contract BisectionChallenge is Cloneable {
    enum State { NoChallenge, AsserterTurn, ChallengerTurn }

    event InitiatedChallenge(uint256 deadlineTicks);
    event Bisected(
        uint256 segmentIndex,
        bytes32[] chainHashes,
        uint256 totalLength,
        uint256 deadlineTicks
    );
    event AsserterTimedOut();
    event ChallengerTimedOut();

    // Can online initialize once
    string private constant CHAL_INIT_STATE = "CHAL_INIT_STATE";
    // Can only bisect assertion in response to a challenge
    string private constant BIS_STATE = "BIS_STATE";
    // deadline expired
    string private constant BIS_DEADLINE = "BIS_DEADLINE";
    // Only original asserter can continue bisect
    string private constant BIS_SENDER = "BIS_SENDER";
    // Incorrect previous state
    string private constant BIS_PREV = "BIS_PREV";
    // Invalid assertion selected
    string private constant CON_PROOF = "CON_PROOF";
    // Can't timeout before deadline
    string private constant TIMEOUT_DEADLINE = "TIMEOUT_DEADLINE";

    address internal rollupAddress;
    address payable internal asserter;
    address payable internal challenger;

    uint256 internal deadlineTicks;

    // The current deadline at which the challenge timeouts and a winner is
    // declared. This deadline resets at each step in the challenge
    uint256 private challengePeriodTicks;

    State private state;

    // This is the root of a merkle tree with nodes like (prev, next, steps)
    bytes32 internal challengeState;

    modifier onlyOnTurn {
        if (state == State.AsserterTurn) {
            require(msg.sender == asserter, BIS_SENDER);
        } else if (state == State.ChallengerTurn) {
            require(msg.sender == challenger, BIS_SENDER);
        } else {
            require(false, BIS_STATE);
        }
        require(RollupTime.blocksToTicks(block.number) <= deadlineTicks, BIS_DEADLINE);
        _;
    }

    function initializeChallenge(
        address _rollupAddress,
        address payable _asserter,
        address payable _challenger,
        uint256 _challengePeriodTicks,
        bytes32 _challengeState
    ) external {
        require(state == State.NoChallenge, CHAL_INIT_STATE);

        rollupAddress = _rollupAddress;
        asserter = _asserter;
        challenger = _challenger;
        challengePeriodTicks = _challengePeriodTicks;
        challengeState = _challengeState;
        state = State.AsserterTurn;
        updateDeadline();

        emit InitiatedChallenge(deadlineTicks);
    }

    function timeoutChallenge() public {
        require(RollupTime.blocksToTicks(block.number) > deadlineTicks, TIMEOUT_DEADLINE);

        if (state == State.AsserterTurn) {
            emit AsserterTimedOut();
            _challengerWin();
        } else {
            emit ChallengerTimedOut();
            _asserterWin();
        }
    }

    function bisect(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength
    ) external onlyOnTurn {
        uint256 bisectionCount = _chainHashes.length - 1;

        require(_chainHashes[bisectionCount] != _oldEndHash);
        require(_chainLength > 1, "bisection too short");

        bytes32 bisectionHash = ChallengeLib.bisectionChunkHash(
            _chainLength,
            _chainHashes[0],
            _oldEndHash
        );

        require(
            MerkleLib.verifyProof(_proof, challengeState, bisectionHash, _segmentToChallenge + 1),
            CON_PROOF
        );

        bytes32[] memory hashes = new bytes32[](bisectionCount);
        hashes[0] = ChallengeLib.bisectionChunkHash(
            ChallengeLib.firstSegmentSize(_chainLength, bisectionCount),
            _chainHashes[0],
            _chainHashes[1]
        );
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeLib.bisectionChunkHash(
                ChallengeLib.otherSegmentSize(_chainLength, bisectionCount),
                _chainHashes[i],
                _chainHashes[i + 1]
            );
        }
        challengeState = MerkleLib.generateRoot(hashes);
        responded();
        emit Bisected(_segmentToChallenge, _chainHashes, _chainLength, deadlineTicks);
    }

    function updateDeadline() internal {
        deadlineTicks = RollupTime.blocksToTicks(block.number) + challengePeriodTicks;
    }

    function responded() internal {
        if (state == State.ChallengerTurn) {
            state = State.AsserterTurn;
        } else {
            state = State.ChallengerTurn;
        }
        updateDeadline();
    }

    function _asserterWin() internal {
        IStaking(rollupAddress).resolveChallenge(asserter, challenger);
        safeSelfDestruct(msg.sender);
    }

    function _challengerWin() internal {
        IStaking(rollupAddress).resolveChallenge(challenger, asserter);
        safeSelfDestruct(msg.sender);
    }

    function commitToSegment(bytes32[] memory hashes) internal {
        challengeState = MerkleLib.generateRoot(hashes);
    }

    function requireMatchesPrevState(bytes32 _challengeState) internal view {
        require(_challengeState == challengeState, BIS_PREV);
    }
}
