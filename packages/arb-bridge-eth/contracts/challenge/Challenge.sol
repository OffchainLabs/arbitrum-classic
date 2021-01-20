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
import "../bridge/Messages.sol";
import "../arch/Marshaling.sol";
import "../libraries/MerkleLib.sol";

contract Challenge is Cloneable, IChallenge {
    using SafeMath for uint256;

    enum Kind { Uninitialized, InboxConsistency, InboxDelta, Execution, StoppedShort }

    enum Turn { NoChallenge, Asserter, Challenger }

    event InitiatedChallenge();
    event Bisected(
        bytes32 indexed challengeRoot,
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength,
        bytes32[] chainHashes
    );
    event BisectedInboxDelta(
        bytes32 indexed challengeRoot,
        uint256 challengedSegmentStart,
        uint256 challengedSegmentLength,
        bytes32[] inboxAccHashes,
        bytes32[] inboxDeltaHashes
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

    IRollup internal resultReceiver;

    bytes32 inboxConsistencyHash;
    bytes32 inboxDeltaHash;
    bytes32 executionHash;

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

    modifier inboxConsistencyChallenge {
        verifyAndSetup(Kind.InboxConsistency, inboxConsistencyHash);
        _;
    }

    modifier inboxDeltaChallenge {
        verifyAndSetup(Kind.InboxDelta, inboxDeltaHash);
        _;
    }

    modifier executionChallenge {
        // If we're in a stopped short challenge and the next step is an execution challenge, that means the asserter has decided to challenge the bisection
        if (kind == Kind.StoppedShort) {
            kind = Kind.Execution;
            executionHash = 0;
        }
        verifyAndSetup(Kind.Execution, executionHash);
        _;
    }

    function initializeChallenge(
        address _executionOneStepProofCon,
        address _executionOneStepProof2Con,
        address _resultReceiver,
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash,
        address _asserter,
        address _challenger,
        uint256 _asserterTimeLeft,
        uint256 _challengerTimeLeft
    ) external override {
        require(turn == Turn.NoChallenge, CHAL_INIT_STATE);

        executor = IOneStepProof(_executionOneStepProofCon);
        executor2 = IOneStepProof2(_executionOneStepProof2Con);

        resultReceiver = IRollup(_resultReceiver);

        inboxConsistencyHash = _inboxConsistencyHash;
        inboxDeltaHash = _inboxDeltaHash;
        executionHash = _executionHash;

        asserter = _asserter;
        challenger = _challenger;
        asserterTimeLeft = _asserterTimeLeft;
        challengerTimeLeft = _challengerTimeLeft;

        kind = Kind.Uninitialized;
        turn = Turn.Challenger;

        challengeState = 0;

        lastMoveBlock = block.number;

        emit InitiatedChallenge();
    }

    /**
     * @notice Initiate the next round in the bisection by objecting to inbox consistency with a bisection
     * of an inbox acculator segment with the same length but a different endpoint.  This is either the
     * initial move or follows another inbox consistency objection
     *
     * @param _merkleNodes List of hashes of stubs in the merkle root of segments left by the previous round
     * @param _merkleRoute Bitmap marking whether the path went left or right at each height
     * @param _challengedSegmentStart Offset of the challenged segment into the original challenged segment
     * @param _challengedSegmentLength Number of messages in the challenged segment
     * @param _oldEndHash End of the challenged segment. This must be different than the new end since the challenger is disagreeing
     * @param _chainHashes Array of intermediate hashes of the challenged segment
     */
    function bisectInboxConsistency(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes
    ) external inboxConsistencyChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "TOO_SHORT");
        require(
            _chainHashes.length ==
                bisectionDegree(_challengedSegmentLength, INBOX_CONSISTENCY_BISECTION_DEGREE) + 1,
            "CUT_COUNT"
        );
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash, "END_HASH");

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                _chainHashes[0],
                _oldEndHash
            );

        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        updateBisectionRoot(_chainHashes, _challengedSegmentStart, _challengedSegmentLength);

        emit Bisected(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _chainHashes
        );
    }

    /**
     * @notice Prove the correctness of a single inbox consistency step by proving that you know a preimage
     * for the start of the length 1 segment that leads to a different after inbox accumulator value.
     * This is either the initial move or follows another inbox consistency objection
     *
     * @param _merkleNodes List of hashes of stubs in the merkle root of segments left by the previous round
     * @param _merkleRoute Bitmap marking whether the path went left or right at each height
     * @param _challengedSegmentStart Offset of the challenged segment into the original challenged segment
     * @param _oldEndHash End of the challenged segment. This must be different than the new end since the challenger is disagreeing
     * @param _lowerHash Correct inbox hash after removing the value from the start point of the segment
     * @param _value Hash of the inbox value at in the segment
     */
    function oneStepProveInboxConsistency(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndHash,
        bytes32 _lowerHash,
        bytes32 _value
    ) external inboxConsistencyChallenge onlyOnTurn {
        require(_lowerHash != _oldEndHash, "SAME_END");
        bytes32 upperHash = Messages.addMessageToInbox(_lowerHash, _value);
        bytes32 prevHash =
            ChallengeLib.bisectionChunkHash(_challengedSegmentStart, 1, upperHash, _oldEndHash);

        verifySegmentProof(prevHash, _merkleNodes, _merkleRoute);

        emit OneStepProofCompleted();
        _currentWin();
    }

    /**
     * @notice Initiate the next round in the bisection by objecting to inbox delta correctness with a bisection
     * of an inbox delta segment with the same length but a different endpoint. This is either the initial move
     * or follows another inbox delta objection
     *
     * @param _merkleNodes List of hashes of stubs in the merkle root of segments left by the previous round
     * @param _merkleRoute Bitmap marking whether the path went left or right at each height
     * @param _challengedSegmentStart Offset of the challenged segment into the original challenged segment
     * @param _challengedSegmentLength Number of messages in the challenged segment
     * @param _oldEndInboxDelta Inbox delta hash of the end of the challenged segment.
     * This must be different than the new inbox delta end since the challenger is disagreeing
     * @param _inboxAccHashes Array of intermediate inbox accumulator hashes of the challenged segment
     * @param _inboxDeltaHashes Array of intermediate inbox delta hashes of the challenged segment
     */
    function bisectInboxDelta(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndInboxDelta,
        bytes32[] calldata _inboxAccHashes,
        bytes32[] calldata _inboxDeltaHashes
    ) external inboxDeltaChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "TOO_SHORT");

        uint256 newSegmentCount = _inboxAccHashes.length;
        require(_inboxDeltaHashes.length == newSegmentCount, "WRONG_COUNT");
        require(
            newSegmentCount ==
                bisectionDegree(_challengedSegmentLength, INBOX_DELTA_BISECTION_DEGREE) + 1,
            "CUT_COUNT"
        );
        require(_inboxDeltaHashes[newSegmentCount - 1] != _oldEndInboxDelta, "WRONG_END");

        bytes32[] memory chainHashes = new bytes32[](newSegmentCount);
        for (uint256 i = 0; i < newSegmentCount; i++) {
            chainHashes[i] = ChallengeLib.inboxDeltaHash(_inboxAccHashes[i], _inboxDeltaHashes[i]);
        }
        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                chainHashes[0],
                ChallengeLib.inboxDeltaHash(_inboxAccHashes[newSegmentCount - 1], _oldEndInboxDelta)
            );

        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        updateBisectionRoot(chainHashes, _challengedSegmentStart, _challengedSegmentLength);

        emit BisectedInboxDelta(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _inboxAccHashes,
            _inboxDeltaHashes
        );
    }

    /**
     * @notice Prove the correctness of a single inbox delta step by proving that you know a preimage
     * for the inbox acc at the start of the length 1 segment that leads to a different inbox delta at the end.
     * This is either the initial move or follows another inbox delta objection
     *
     * @param _merkleNodes List of hashes of stubs in the merkle root of segments left by the previous round
     * @param _merkleRoute Bitmap marking whether the path went left or right at each height
     * @param _challengedSegmentStart Offset of the challenged segment into the original challenged segment
     * @param _oldEndInboxDelta Inbox delta hash of the end of the challenged segment.
     * This must be different than the new inbox delta end since the challenger is disagreeing
     * @param _prevInboxDelta Inbox delta of the beginning of the segment
     * @param _nextInboxAcc Inbox accumulator of the end of the segment
     * @param _kind Message kind of the message in the segment
     * @param _blockNumber Block number of the message in the segment
     * @param _timestamp Timestamp of the message in the segment
     * @param _sender Sender of the message in the segment
     * @param _inboxSeqNum Sequence number of the message in the segment
     * @param _msgData Data of the message in the segment
     */
    function oneStepProveInboxDelta(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndInboxDelta,
        bytes32 _prevInboxDelta,
        bytes32 _nextInboxAcc,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _msgData
    ) public inboxDeltaChallenge onlyOnTurn {
        bytes32 chunkHash =
            oneStepProveInboxDeltaOldChunkHash(
                _challengedSegmentStart,
                _oldEndInboxDelta,
                _prevInboxDelta,
                _nextInboxAcc,
                Messages.messageHash(
                    _kind,
                    _sender,
                    _blockNumber,
                    _timestamp,
                    _inboxSeqNum,
                    keccak256(_msgData)
                ),
                messageValueHash(_kind, _blockNumber, _timestamp, _sender, _inboxSeqNum, _msgData)
            );

        verifySegmentProof(chunkHash, _merkleNodes, _merkleRoute);

        emit OneStepProofCompleted();
        _currentWin();
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
    //  _initialMessageCount
    //  _initialLogCount
    function oneStepProveExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndHash,
        bytes32[3] memory _machineFields,
        uint256[3] memory _initialState,
        bytes memory _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) public executionChallenge onlyOnTurn {
        (uint64 gasUsed, bytes32[5] memory proofFields) =
            executeMachineStep(prover, _machineFields, _executionProof, _bufferProof);

        require(
            _oldEndHash !=
                oneStepProofExecutionAfter(_machineFields, _initialState, gasUsed, proofFields),
            "WRONG_END"
        );

        bytes32 rootHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                gasUsed,
                oneStepProofExecutionBefore(_machineFields, _initialState, proofFields),
                _oldEndHash
            );

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
        // Free no longer needed storage
        inboxConsistencyHash = 0;
        inboxDeltaHash = 0;

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
        // Free no longer needed storage
        inboxConsistencyHash = 0;
        inboxDeltaHash = 0;
        executionHash = 0;
    }

    // initialState
    //  _initialGasUsed
    //  _initialMessageCount
    //  _initialLogCount
    function oneStepProveStoppedShort(
        bytes32[3] calldata _machineFields,
        uint256[3] calldata _initialState,
        bytes calldata _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) external onlyOnTurn {
        require(kind == Kind.StoppedShort, "WRONG_KIND");

        // If this doesn't revert, we were able to successfully execute the machine
        (, bytes32[5] memory proofFields) =
            executeMachineStep(prover, _machineFields, _executionProof, _bufferProof);

        // Check that the before state is the end of the stopped short bisection which was stored in executionHash
        require(
            oneStepProofExecutionBefore(_machineFields, _initialState, proofFields) ==
                executionHash,
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

    function verifyAndSetup(Kind _kind, bytes32 initialState) private {
        if (kind == Kind.Uninitialized) {
            challengeState = initialState;
            kind = _kind;
            // Free no longer needed storage
            inboxConsistencyHash = 0;
            inboxDeltaHash = 0;
            executionHash = 0;
        } else {
            require(kind == _kind, "WRONG_KIND");
        }
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

    function oneStepProveInboxDeltaOldChunkHash(
        uint256 _challengedSegmentStart,
        bytes32 _oldEndInboxDelta,
        bytes32 _prevInboxDelta,
        bytes32 _nextInboxAcc,
        bytes32 _messageHash,
        bytes32 _messageValueHash
    ) private pure returns (bytes32) {
        require(
            _oldEndInboxDelta != Messages.addMessageToInbox(_prevInboxDelta, _messageValueHash),
            "WRONG_END"
        );
        bytes32 prevInboxAcc = Messages.addMessageToInbox(_nextInboxAcc, _messageHash);
        return
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                1,
                ChallengeLib.inboxDeltaHash(prevInboxAcc, _prevInboxDelta),
                ChallengeLib.inboxDeltaHash(_nextInboxAcc, _oldEndInboxDelta)
            );
    }

    function executeMachineStep(
        uint8 prover,
        bytes32[3] memory _machineFields,
        bytes memory _executionProof,
        bytes memory _bufferProof
    ) private view returns (uint64 gas, bytes32[5] memory fields) {
        if (prover == 0) {
            return executor.executeStep(_machineFields, _executionProof);
        } else if (prover == 1) {
            return executor2.executeStep(_machineFields, _executionProof, _bufferProof);
        } else {
            require(false, "INVALID_PROVER");
        }
    }

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    // proofFields
    //  initialMachineHash
    //  afterMachineHash
    //  afterInboxHash
    //  afterMessagesHash
    //  afterLogsHash
    function oneStepProofExecutionBefore(
        bytes32[3] memory _machineFields,
        uint256[3] memory _initialState,
        bytes32[5] memory proofFields
    ) private pure returns (bytes32) {
        return
            ChallengeLib.assertionHash(
                _initialState[0],
                ChallengeLib.assertionRestHash(
                    _machineFields[0],
                    proofFields[0],
                    _machineFields[1],
                    _initialState[1],
                    _machineFields[2],
                    _initialState[2]
                )
            );
    }

    function oneStepProofExecutionAfter(
        bytes32[3] memory _machineFields,
        uint256[3] memory _initialState,
        uint64 gasUsed,
        bytes32[5] memory proofFields
    ) private pure returns (bytes32) {
        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        return
            ChallengeLib.assertionHash(
                _initialState[0].add(gasUsed),
                ChallengeLib.assertionRestHash(
                    proofFields[2],
                    proofFields[1],
                    proofFields[3],
                    _initialState[1].add((_machineFields[1] == proofFields[3] ? 0 : 1)),
                    proofFields[4],
                    _initialState[2].add((_machineFields[2] == proofFields[4] ? 0 : 1))
                )
            );
    }

    function messageValueHash(
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _messageData
    ) internal pure returns (bytes32) {
        bytes32 messageBufHash = Hashing.bytesToBufferHash(_messageData, 0, _messageData.length);
        Value.Data[] memory tupData = new Value.Data[](7);
        tupData[0] = Value.newInt(uint256(_kind));
        tupData[1] = Value.newInt(_blockNumber);
        tupData[2] = Value.newInt(_timestamp);
        tupData[3] = Value.newInt(uint256(_sender));
        tupData[4] = Value.newInt(_inboxSeqNum);
        tupData[5] = Value.newInt(_messageData.length);
        tupData[6] = Value.newHashedValue(messageBufHash, 1);

        return Hashing.hash(Value.newTuple(tupData));
    }
}
