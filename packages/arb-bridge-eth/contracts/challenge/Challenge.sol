// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2020, Offchain Labs, Inc.
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

import "./IChallenge.sol";
import "../rollup/IRollup.sol";
import "./ChallengeLib.sol";

import "../bridge/Messages.sol";
import "../rollup/IRollup.sol";
import "../arch/IOneStepProof.sol";
import "../arch/Marshaling.sol";

import "../libraries/Cloneable.sol";
import "../libraries/MerkleLib.sol";

contract Challenge is Cloneable, IChallenge {
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

    uint256 private constant BISECTION_DEGREE = 20;

    IOneStepProof private executor;
    IOneStepProof2 private executor2;

    IRollup internal resultReceiver;

    bytes32 inboxConsistencyHash;
    bytes32 inboxDeltaHash;
    bytes32 executionHash;

    address public asserter;
    address public challenger;
    uint256 private challengePeriodBlocks;
    uint256 private executionCheckTimeBlocks;

    Kind public kind;

    // The current deadline at which the challenge timeouts and a winner is
    // declared. This deadline resets at each step in the challenge
    uint256 public deadlineBlock;
    Turn public turn;
    // This is the root of a merkle tree with nodes like (prev, next, steps)
    bytes32 public challengeState;

    modifier onlyOnTurn {
        require(msg.sender == currentResponder(), BIS_SENDER);
        require(block.number <= deadlineBlock, BIS_DEADLINE);
        _;
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
        uint256 _executionCheckTimeBlocks,
        address _asserter,
        address _challenger,
        uint256 _challengePeriodBlocks
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
        challengePeriodBlocks = _challengePeriodBlocks;
        executionCheckTimeBlocks = _executionCheckTimeBlocks;

        kind = Kind.Uninitialized;

        deadlineBlock = block.number + _challengePeriodBlocks + _executionCheckTimeBlocks;
        turn = Turn.Challenger;

        challengeState = 0;

        emit InitiatedChallenge();
    }

    function bisectInboxConsistency(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes
    ) external inboxConsistencyChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "bisection too short");
        require(_chainHashes.length == bisectionDegree(_challengedSegmentLength) + 1);
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash);

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                _chainHashes[0],
                _oldEndHash
            );

        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        updateBisectionRoot(_chainHashes, _challengedSegmentStart, _challengedSegmentLength);

        responded(1);
        emit Bisected(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _chainHashes
        );
    }

    function oneStepProveInboxConsistency(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndHash,
        bytes32 _lowerHash,
        bytes32 _value
    ) external inboxConsistencyChallenge onlyOnTurn {
        require(_lowerHash != _oldEndHash);
        bytes32 upperHash = Messages.addMessageToInbox(_lowerHash, _value);
        bytes32 prevHash =
            ChallengeLib.bisectionChunkHash(_challengedSegmentStart, 1, upperHash, _oldEndHash);

        verifySegmentProof(prevHash, _merkleNodes, _merkleRoute);

        emit OneStepProofCompleted();
        _currentWin();
    }

    function bisectInboxDelta(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldInboxDelta,
        bytes32[] calldata _inboxAccHashes,
        bytes32[] calldata _inboxDeltaHashes
    ) external inboxDeltaChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "bisection too short");

        uint256 newSegmentCount = _inboxAccHashes.length;
        require(_inboxDeltaHashes.length == newSegmentCount, "WRONG_COUNT");
        require(newSegmentCount == bisectionDegree(_challengedSegmentLength) + 1);
        require(_inboxDeltaHashes[newSegmentCount - 1] != _oldInboxDelta);

        bytes32[] memory chainHashes = new bytes32[](newSegmentCount);
        for (uint256 i = 0; i < newSegmentCount; i++) {
            chainHashes[i] = ChallengeLib.inboxDeltaHash(_inboxAccHashes[i], _inboxDeltaHashes[i]);
        }
        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                chainHashes[0],
                ChallengeLib.inboxDeltaHash(_inboxAccHashes[newSegmentCount - 1], _oldInboxDelta)
            );

        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        updateBisectionRoot(chainHashes, _challengedSegmentStart, _challengedSegmentLength);

        responded(1);
        emit BisectedInboxDelta(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _inboxAccHashes,
            _inboxDeltaHashes
        );
    }

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

    function bisectExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _gasUsedBefore,
        bytes32 _assertionRest
    ) external executionChallenge onlyOnTurn {
        require(_challengedSegmentLength > 1, "TOO_SHORT");
        require(
            _chainHashes.length == bisectionDegree(_challengedSegmentLength) + 1,
            "BISECT_DEGREE"
        );
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash, "SAME_END");

        require(
            _chainHashes[0] == ChallengeLib.assertionHash(_gasUsedBefore, _assertionRest),
            "segment pre-fields"
        );

        require(
            _gasUsedBefore < _challengedSegmentStart + _challengedSegmentLength,
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
            _challengedSegmentStart + _challengedSegmentLength - _gasUsedBefore
        );

        responded(executionCheckTimeBlocks);
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

        require(_gasUsedBefore >= _challengedSegmentStart + _challengedSegmentLength);
        require(beforeChainHash != _oldEndHash);
        emit ConstraintWin();
        _currentWin();
    }

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    function oneStepProveExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        bytes32 _oldEndHash,
        bytes32[3] memory _machineFields,
        uint64 _initialGasUsed,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        bytes memory _executionProof,
        bytes memory _bufferProof
    ) public executionChallenge onlyOnTurn {
        uint64 gasUsed;
        bytes32[5] memory proofFields;

        if (_bufferProof.length == 0) {
            (gasUsed, proofFields) = executor.executeStep(_machineFields, _executionProof);
        } else {
            (gasUsed, proofFields) = executor2.executeStep(
                _machineFields,
                _executionProof,
                _bufferProof
            );
        }

        require(
            _oldEndHash !=
                oneStepProofExecutionAfter(
                    _machineFields,
                    _initialGasUsed,
                    _initialMessageCount,
                    _initialLogCount,
                    gasUsed,
                    proofFields
                )
        );

        bytes32 rootHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                gasUsed,
                oneStepProofExecutionBefore(
                    _machineFields,
                    _initialGasUsed,
                    _initialMessageCount,
                    _initialLogCount,
                    proofFields
                ),
                _oldEndHash
            );

        verifySegmentProof(rootHash, _merkleNodes, _merkleRoute);

        emit OneStepProofCompleted();
        _currentWin();
    }

    // Can only do a stopped short bisection as a first move
    function bisectExecutionStoppedShort(
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _newSegmentLength,
        bytes32 _startAssertionHash
    ) external onlyOnTurn {
        require(kind == Kind.Uninitialized);
        // Unlike the other bisections, it's safe for the number of steps executed to be 1
        require(_newSegmentLength > 0);
        require(_chainHashes.length == bisectionDegree(_newSegmentLength) + 1);
        require(_newSegmentLength < _challengedSegmentLength);

        require(
            ChallengeLib.bisectionChunkHash(
                0,
                _challengedSegmentLength,
                _startAssertionHash,
                _oldEndHash
            ) == executionHash
        );

        // Reuse the executionHash variable to store last assertion
        updateBisectionRoot(_chainHashes, 0, _newSegmentLength);
        executionHash = _chainHashes[_chainHashes.length - 1];

        kind = Kind.StoppedShort;
        // Free no longer needed storage
        inboxConsistencyHash = 0;
        inboxDeltaHash = 0;
        executionHash = 0;

        responded(executionCheckTimeBlocks);

        emit Bisected(challengeState, 0, _challengedSegmentLength, _chainHashes);
    }

    // Can only do a stopped short bisection as a first move
    function executionCantRun(
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        bytes32 _startAssertionHash
    ) external onlyOnTurn {
        require(kind == Kind.Uninitialized);

        require(
            ChallengeLib.bisectionChunkHash(
                0,
                _challengedSegmentLength,
                _startAssertionHash,
                _oldEndHash
            ) == executionHash
        );
        executionHash = _startAssertionHash;
        kind = Kind.StoppedShort;
        // Free no longer needed storage
        inboxConsistencyHash = 0;
        inboxDeltaHash = 0;
        executionHash = 0;

        responded(executionCheckTimeBlocks);
    }

    function oneStepProveStoppedShort(
        bytes32[3] calldata _machineFields,
        uint64 _initialGasUsed,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        bytes calldata _executionProof
    ) external onlyOnTurn {
        require(kind == Kind.StoppedShort);

        // If this doesn't revert, we were able to successfully execute the machine
        (, bytes32[5] memory proofFields) = executor.executeStep(_machineFields, _executionProof);

        // Check that the before state is the end of the stopped short bisection which was stored in executionHash
        require(
            oneStepProofExecutionBefore(
                _machineFields,
                _initialGasUsed,
                _initialMessageCount,
                _initialLogCount,
                proofFields
            ) == executionHash
        );

        emit OneStepProofCompleted();
        _currentWin();
    }

    function timeout() external {
        require(block.number > deadlineBlock, TIMEOUT_DEADLINE);

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
        segmentStart += chunkSize;
        chunkSize = ChallengeLib.otherSegmentSize(_challengedSegmentLength, bisectionCount);
        for (uint256 i = 1; i < bisectionCount; i++) {
            hashes[i] = ChallengeLib.bisectionChunkHash(
                segmentStart,
                chunkSize,
                _chainHashes[i],
                _chainHashes[i + 1]
            );
            segmentStart += chunkSize;
        }
        (bytes32 root, ) = MerkleLib.generateMerkleRoot(hashes);
        challengeState = root;
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
            require(kind == _kind);
        }
    }

    function responded(uint256 additionalTimeBlocks) private {
        if (turn == Turn.Challenger) {
            turn = Turn.Asserter;
        } else {
            turn = Turn.Challenger;
        }
        deadlineBlock = block.number + challengePeriodBlocks + additionalTimeBlocks;
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

    function bisectionDegree(uint256 _chainLength) private pure returns (uint256) {
        if (_chainLength < BISECTION_DEGREE) {
            return _chainLength;
        } else {
            return BISECTION_DEGREE;
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
            _oldEndInboxDelta != Messages.addMessageToInbox(_prevInboxDelta, _messageValueHash)
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
        uint64 _initialGasUsed,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        bytes32[5] memory proofFields
    ) private pure returns (bytes32) {
        return
            ChallengeLib.assertionHash(
                _initialGasUsed,
                ChallengeLib.assertionRestHash(
                    _machineFields[0],
                    proofFields[0],
                    _machineFields[1],
                    _initialMessageCount,
                    _machineFields[2],
                    _initialLogCount
                )
            );
    }

    function oneStepProofExecutionAfter(
        bytes32[3] memory _machineFields,
        uint64 _initialGasUsed,
        uint256 _initialMessageCount,
        uint256 _initialLogCount,
        uint64 gasUsed,
        bytes32[5] memory proofFields
    ) private pure returns (bytes32) {
        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        return
            ChallengeLib.assertionHash(
                _initialGasUsed + gasUsed,
                ChallengeLib.assertionRestHash(
                    proofFields[2],
                    proofFields[1],
                    proofFields[3],
                    _initialMessageCount + (_machineFields[1] == proofFields[3] ? 0 : 1),
                    proofFields[4],
                    _initialLogCount + (_machineFields[2] == proofFields[4] ? 0 : 1)
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
        Value.Data[] memory tupData = new Value.Data[](6);
        tupData[0] = Value.newInt(uint256(_kind));
        tupData[1] = Value.newInt(_blockNumber);
        tupData[2] = Value.newInt(_timestamp);
        tupData[3] = Value.newInt(uint256(_sender));
        tupData[4] = Value.newInt(_inboxSeqNum);
        tupData[5] = Marshaling.bytesToBytestack(_messageData, 0, _messageData.length);
        return Hashing.hash(Value.newTuple(tupData));
    }
}
