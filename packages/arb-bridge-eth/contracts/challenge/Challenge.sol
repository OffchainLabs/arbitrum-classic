// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

import "../libraries/Cloneable.sol";
import "../libraries/SafeMath.sol";

import "./IChallenge.sol";
import "../rollup/IRollup.sol";
import "../arch/IOneStepProof.sol";

import "./ChallengeLib.sol";
import "../libraries/MerkleLib.sol";

contract Challenge is Cloneable, IChallenge {
    using SafeMath for uint256;

    enum Kind { Uninitialized, Execution, StoppedShort }

    enum Turn { NoChallenge, Asserter, Challenger }

    event InitiatedChallenge();
    event Bisected(
        bytes32 indexed challengeRoot,
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength,
        bytes32[] chainHashes
    );
    event AsserterTimedOut();
    event ChallengerTimedOut();
    event OneStepProofCompleted();
    event ConstraintWin();

    // Can only initialize once
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

    uint256 private constant INBOX_CONSISTENCY_BISECTION_DEGREE = 400;
    uint256 private constant INBOX_DELTA_BISECTION_DEGREE = 250;
    uint256 private constant EXECUTION_BISECTION_DEGREE = 400;

    IOneStepProof public executor;
    IOneStepProof2 public executor2;
    IBridge public bridge;

    IRollup internal resultReceiver;

    bytes32 executionHash;
    uint256 maxMessageCount;

    address public asserter;
    address public challenger;

    uint256 public lastMoveBlock;
    uint256 public asserterTimeLeft;
    uint256 public challengerTimeLeft;

    Kind public kind;
    Turn public turn;

    // This is the root of a merkle tree with nodes like (prev, next, steps)
    bytes32 public challengeState;

    modifier onlyOnTurn {
        require(msg.sender == currentResponder(), BIS_SENDER);
        require(block.number.sub(lastMoveBlock) <= currentResponderTimeLeft(), BIS_DEADLINE);

        _;

        if (turn == Turn.Challenger) {
            challengerTimeLeft = challengerTimeLeft.sub(block.number.sub(lastMoveBlock));
            turn = Turn.Asserter;
        } else {
            asserterTimeLeft = asserterTimeLeft.sub(block.number.sub(lastMoveBlock));
            turn = Turn.Challenger;
        }
        lastMoveBlock = block.number;
    }

    modifier executionChallenge {
        // If we're in a stopped short challenge and the next step is an execution challenge, that means the asserter has decided to challenge the bisection
        if (kind == Kind.StoppedShort) {
            kind = Kind.Execution;
            executionHash = 0;
        }
        if (kind == Kind.Uninitialized) {
            challengeState = executionHash;
            kind = Kind.Execution;
            // Free no longer needed storage
            executionHash = 0;
        } else {
            require(kind == Kind.Execution, "WRONG_KIND");
        }
        _;
    }

    function initializeChallenge(
        address _executionOneStepProofCon,
        address _executionOneStepProof2Con,
        address _resultReceiver,
        bytes32 _executionHash,
        uint256 _maxMessageCount,
        address _asserter,
        address _challenger,
        uint256 _asserterTimeLeft,
        uint256 _challengerTimeLeft,
        IBridge _bridge
    ) external override {
        require(turn == Turn.NoChallenge, CHAL_INIT_STATE);

        executor = IOneStepProof(_executionOneStepProofCon);
        executor2 = IOneStepProof2(_executionOneStepProof2Con);

        resultReceiver = IRollup(_resultReceiver);

        executionHash = _executionHash;

        maxMessageCount = _maxMessageCount;

        asserter = _asserter;
        challenger = _challenger;
        asserterTimeLeft = _asserterTimeLeft;
        challengerTimeLeft = _challengerTimeLeft;

        kind = Kind.Uninitialized;
        turn = Turn.Challenger;

        challengeState = 0;

        lastMoveBlock = block.number;
        bridge = _bridge;

        emit InitiatedChallenge();
    }

    /**
     * @notice Initiate the next round in the bisection by objecting to execution correctness with a bisection
     * of an execution segment with the same length but a different endpoint. This is either the initial move
     * or follows another execution objection
     *
     * @param _merkleNodes List of hashes of stubs in the merkle root of segments left by the previous round
     * @param _merkleRoute Bitmap marking whether the path went left or right at each height
     * @param _challengedSegmentStart Offset of the challenged segment into the original challenged segment
     * @param _challengedSegmentLength Number of messages in the challenged segment
     * @param _oldEndHash Hash of the end of the challenged segment. This must be different than the new end since the challenger is disagreeing
     * @param _gasUsedBefore Amount of gas used at the beginning of the challenged segment
     * @param _assertionRest Hash of the rest of the assertion at the beginning of the challenged segment
     * @param _chainHashes Array of intermediate hashes of the challenged segment
     */
    function bisectExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        uint256 _gasUsedBefore,
        bytes32 _assertionRest,
        bytes32[] calldata _chainHashes
    ) external executionChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "TOO_SHORT");
        require(
            _chainHashes.length ==
                bisectionDegree(_challengedSegmentLength, EXECUTION_BISECTION_DEGREE) + 1,
            "CUT_COUNT"
        );
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash, "SAME_END");

        require(
            _chainHashes[0] == ChallengeLib.assertionHash(_gasUsedBefore, _assertionRest),
            "segment pre-fields"
        );

        require(
            _gasUsedBefore < _challengedSegmentStart.add(_challengedSegmentLength),
            "invalid segment length"
        );

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                _chainHashes[0],
                _oldEndHash
            );
        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        updateBisectionRoot(
            _chainHashes,
            _challengedSegmentStart,
            _challengedSegmentStart.add(_challengedSegmentLength).sub(_gasUsedBefore)
        );

        emit Bisected(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _chainHashes
        );
    }

    function constraintWinExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        uint256 _gasUsedBefore,
        bytes32 _assertionRest
    ) external executionChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "TOO SHORT");

        bytes32 beforeChainHash = ChallengeLib.assertionHash(_gasUsedBefore, _assertionRest);

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                beforeChainHash,
                _oldEndHash
            );
        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        require(_gasUsedBefore >= _challengedSegmentStart.add(_challengedSegmentLength), "BAD_GAS");
        require(beforeChainHash != _oldEndHash, "WRONG_END");
        emit ConstraintWin();
        _currentWin();
    }

    // machineFields
    //  initialInbox
    //  initialMessageAcc
    //  initialLogAcc
    // initialState
    //  _initialGasUsed
    //  _initialSendCount
    //  _initialLogCount
    function oneStepProveExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndHash,
        uint256 _initialMessagesRead,
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        uint256[3] memory _initialState,
        bytes memory _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) public executionChallenge onlyOnTurn {
        bytes32 rootHash;
        {
            (uint64 gasUsed, uint256 totalMessagesRead, bytes32[4] memory proofFields) =
                executeMachineStep(
                    prover,
                    _initialMessagesRead,
                    _initialSendAcc,
                    _initialLogAcc,
                    _executionProof,
                    _bufferProof
                );

            require(
                _oldEndHash !=
                    oneStepProofExecutionAfter(
                        _initialSendAcc,
                        _initialLogAcc,
                        _initialState,
                        gasUsed,
                        totalMessagesRead,
                        proofFields
                    ),
                "WRONG_END"
            );

            rootHash = ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                gasUsed,
                oneStepProofExecutionBefore(
                    _initialMessagesRead,
                    _initialSendAcc,
                    _initialLogAcc,
                    _initialState,
                    proofFields
                ),
                _oldEndHash
            );
        }

        verifySegmentProof(rootHash, _merkleNodes, _merkleRoute);

        emit OneStepProofCompleted();
        _currentWin();
    }

    /**
     * @notice Object that the machine should have blocked after less gas used than was claimed and provide
     * bisection of that smaller chunk. This can only occur as the first move in a challenge
     *
     * @dev Can only do a stopped short bisection as a first move
     * @param _challengedSegmentLength Number of messages in the challenged segment
     * @param _oldEndHash Hash of the end of the challenged segment
     * @param _chainHashes Array of intermediate hashes of the challenged segment
     * @param _newSegmentLength New segment length that's shorter than the challenged segment
     * @param _startAssertionHash Hash of the assertion at the beginning of the challenged segment
     */
    function bisectExecutionStoppedShort(
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _newSegmentLength,
        bytes32 _startAssertionHash
    ) external onlyOnTurn {
        require(kind == Kind.Uninitialized, "BAD_KIND");
        // Unlike the other bisections, it's safe for the number of steps executed to be 1
        require(_newSegmentLength > 0, "BAD_LENGTH");
        require(
            _chainHashes.length ==
                bisectionDegree(_newSegmentLength, EXECUTION_BISECTION_DEGREE) + 1,
            "CUT_COUNT"
        );
        require(_newSegmentLength < _challengedSegmentLength, "TOO_LONG");

        require(
            ChallengeLib.bisectionChunkHash(
                0,
                _challengedSegmentLength,
                _startAssertionHash,
                _oldEndHash
            ) == executionHash,
            "END_HASH"
        );

        // Reuse the executionHash variable to store last assertion
        updateBisectionRoot(_chainHashes, 0, _newSegmentLength);
        executionHash = _chainHashes[_chainHashes.length - 1];
        kind = Kind.StoppedShort;

        emit Bisected(challengeState, 0, _challengedSegmentLength, _chainHashes);
    }

    // Can only do a stopped short bisection as a first move
    function executionCantRun(
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32 _startAssertionHash
    ) external onlyOnTurn {
        require(kind == Kind.Uninitialized, "WRONG_KIND");

        require(
            ChallengeLib.bisectionChunkHash(
                0,
                _challengedSegmentLength,
                _startAssertionHash,
                _oldEndHash
            ) == executionHash,
            "WRONG_KIND"
        );
        executionHash = _startAssertionHash;
        kind = Kind.StoppedShort;
    }

    // initialState
    //  _initialGasUsed
    //  _initialSendCount
    //  _initialLogCount
    function oneStepProveStoppedShort(
        uint256 _initialMessagesRead,
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        uint256[3] calldata _initialState,
        bytes calldata _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) external onlyOnTurn {
        require(kind == Kind.StoppedShort, "WRONG_KIND");

        // If this doesn't revert, we were able to successfully execute the machine
        (, uint256 totalMessagesRead, bytes32[4] memory proofFields) =
            executeMachineStep(
                prover,
                _initialMessagesRead,
                _initialSendAcc,
                _initialLogAcc,
                _executionProof,
                _bufferProof
            );
        require(totalMessagesRead <= maxMessageCount, "TOO_MANY_MESSAGES");

        // Check that the before state is the end of the stopped short bisection which was stored in executionHash
        require(
            oneStepProofExecutionBefore(
                _initialMessagesRead,
                _initialSendAcc,
                _initialLogAcc,
                _initialState,
                proofFields
            ) == executionHash,
            "WRONG_END"
        );

        emit OneStepProofCompleted();
        _currentWin();
    }

    function timeout() external {
        uint256 timeSinceLastMove = block.number.sub(lastMoveBlock);
        require(timeSinceLastMove > currentResponderTimeLeft(), TIMEOUT_DEADLINE);

        if (turn == Turn.Asserter) {
            emit AsserterTimedOut();
            _challengerWin();
        } else {
            emit ChallengerTimedOut();
            _asserterWin();
        }
    }

    function currentResponder() public view returns (address) {
        if (turn == Turn.Asserter) {
            return asserter;
        } else if (turn == Turn.Challenger) {
            return challenger;
        } else {
            require(false, "NO_TURN");
        }
    }

    function currentResponderTimeLeft() public view returns (uint256) {
        if (turn == Turn.Asserter) {
            return asserterTimeLeft;
        } else if (turn == Turn.Challenger) {
            return challengerTimeLeft;
        } else {
            require(false, "NO_TURN");
        }
    }

    function updateBisectionRoot(
        bytes32[] memory _chainHashes,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength
    ) private returns (bytes32) {
        uint256 bisectionCount = _chainHashes.length - 1;
        bytes32[] memory hashes = new bytes32[](bisectionCount);
        uint256 chunkSize = ChallengeLib.firstSegmentSize(_challengedSegmentLength, bisectionCount);
        uint256 segmentStart = _challengedSegmentStart;
        hashes[0] = ChallengeLib.bisectionChunkHash(
            segmentStart,
            chunkSize,
            _chainHashes[0],
            _chainHashes[1]
        );
        segmentStart = segmentStart.add(chunkSize);
        chunkSize = ChallengeLib.otherSegmentSize(_challengedSegmentLength, bisectionCount);
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeLib.bisectionChunkHash(
                segmentStart,
                chunkSize,
                _chainHashes[i],
                _chainHashes[i + 1]
            );
            segmentStart = segmentStart.add(chunkSize);
        }
        challengeState = MerkleLib.generateRoot(hashes);
    }

    function _currentWin() private {
        if (turn == Turn.Asserter) {
            _asserterWin();
        } else {
            _challengerWin();
        }
    }

    function _asserterWin() private {
        resultReceiver.completeChallenge(asserter, challenger);
        safeSelfDestruct(msg.sender);
    }

    function _challengerWin() private {
        resultReceiver.completeChallenge(challenger, asserter);
        safeSelfDestruct(msg.sender);
    }

    function verifySegmentProof(
        bytes32 item,
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute
    ) private view {
        require(
            challengeState == MerkleLib.calculateRoot(_merkleNodes, _merkleRoute, item),
            BIS_PREV
        );
    }

    function bisectionDegree(uint256 _chainLength, uint256 targetDegree)
        private
        pure
        returns (uint256)
    {
        if (_chainLength < targetDegree) {
            return _chainLength;
        } else {
            return targetDegree;
        }
    }

    function executeMachineStep(
        uint8 prover,
        uint256 _initialMessagesRead,
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        bytes memory _executionProof,
        bytes memory _bufferProof
    )
        private
        view
        returns (
            uint64 gas,
            uint256 totalMessagesRead,
            bytes32[4] memory fields
        )
    {
        if (prover == 0) {
            return
                executor.executeStep(
                    bridge,
                    _initialMessagesRead,
                    _initialSendAcc,
                    _initialLogAcc,
                    _executionProof
                );
        } else if (prover == 1) {
            return
                executor2.executeStep(
                    _initialMessagesRead,
                    _initialSendAcc,
                    _initialLogAcc,
                    _executionProof,
                    _bufferProof
                );
        } else {
            require(false, "INVALID_PROVER");
        }
    }

    // proofFields
    //  initialMachineHash
    //  afterMachineHash
    //  afterInboxHash
    //  afterMessagesHash
    //  afterLogsHash
    function oneStepProofExecutionBefore(
        uint256 _initialMessagesRead,
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        uint256[3] memory _initialState,
        bytes32[4] memory proofFields
    ) private pure returns (bytes32) {
        return
            ChallengeLib.assertionHash(
                _initialState[0],
                ChallengeLib.assertionRestHash(
                    _initialMessagesRead,
                    proofFields[0],
                    _initialSendAcc,
                    _initialState[1],
                    _initialLogAcc,
                    _initialState[2]
                )
            );
    }

    function oneStepProofExecutionAfter(
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        uint256[3] memory _initialState,
        uint64 gasUsed,
        uint256 totalMessagesRead,
        bytes32[4] memory proofFields
    ) private pure returns (bytes32) {
        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        return
            ChallengeLib.assertionHash(
                _initialState[0].add(gasUsed),
                ChallengeLib.assertionRestHash(
                    totalMessagesRead,
                    proofFields[1],
                    proofFields[2],
                    _initialState[1].add((_initialSendAcc == proofFields[2] ? 0 : 1)),
                    proofFields[3],
                    _initialState[2].add((_initialLogAcc == proofFields[3] ? 0 : 1))
                )
            );
    }
}
