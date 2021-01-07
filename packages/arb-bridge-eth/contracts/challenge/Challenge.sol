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
import "./ChallengeLib.sol";

import "../rollup/Messages.sol";
import "../rollup/IRollup.sol";
import "../arch/IOneStepProof.sol";

import "../libraries/Cloneable.sol";
import "../libraries/MerkleLib.sol";

contract Challenge is Cloneable, IChallenge {
    enum Kind { Uninitialized, InboxConsistency, InboxDelta, Execution, StoppedShort }

    enum State { NoChallenge, AsserterTurn, ChallengerTurn }

    event InitiatedChallenge(uint256 deadlineBlock);
    event Bisected(
        uint256 segmentIndex,
        bytes32[] chainHashes,
        uint256 totalLength,
        uint256 deadlineBlock
    );
    event AsserterTimedOut();
    event ChallengerTimedOut();
    event OneStepProofCompleted();

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

    address internal rollupAddress;

    bytes32 inboxConsistencyHash;
    bytes32 inboxDeltaHash;
    bytes32 executionHash;

    address payable internal asserter;
    address payable internal challenger;
    uint256 private challengePeriodBlocks;
    uint256 private executionCheckTimeBlocks;

    uint256 private gasBefore;

    Kind kind;

    // The current deadline at which the challenge timeouts and a winner is
    // declared. This deadline resets at each step in the challenge
    uint256 internal deadlineBlock;
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
        require(block.number <= deadlineBlock, BIS_DEADLINE);
        _;
    }

    modifier inboxConsistencyChallenge {
        verifyAndSetup(Kind.InboxConsistency, inboxConsistencyHash);
        _;
    }

    modifier inboxDeltaChallenge {
        verifyAndSetup(Kind.InboxConsistency, inboxDeltaHash);
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
        address _rollupAddress,
        bytes32 _inboxConsistencyHash,
        bytes32 _inboxDeltaHash,
        bytes32 _executionHash,
        uint256 _executionCheckTimeBlocks,
        address payable _asserter,
        address payable _challenger,
        uint256 _challengePeriodBlocks
    ) external override {
        require(state == State.NoChallenge, CHAL_INIT_STATE);

        executor = IOneStepProof(_executionOneStepProofCon);
        executor2 = IOneStepProof2(_executionOneStepProof2Con);

        rollupAddress = _rollupAddress;

        inboxConsistencyHash = _inboxConsistencyHash;
        inboxDeltaHash = _inboxDeltaHash;
        executionHash = _executionHash;

        asserter = _asserter;
        challenger = _challenger;
        challengePeriodBlocks = _challengePeriodBlocks;
        executionCheckTimeBlocks = _executionCheckTimeBlocks;

        kind = Kind.Uninitialized;

        deadlineBlock = block.number + _challengePeriodBlocks + _executionCheckTimeBlocks;
        state = State.ChallengerTurn;

        challengeState = 0;

        emit InitiatedChallenge(deadlineBlock);
    }

    function bisectInboxConsistency(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength
    ) external inboxConsistencyChallenge onlyOnTurn {
        require(_chainLength > 1, "bisection too short");
        require(_chainHashes.length == bisectionDegree(_chainLength));
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash);

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(_chainLength, _chainHashes[0], _oldEndHash);

        verifySegmentProof(_proof, bisectionHash, _segmentToChallenge);

        challengeState = calculateBisectionRoot(_chainHashes, _chainLength);

        responded(1);
        emit Bisected(_segmentToChallenge, _chainHashes, _chainLength, deadlineBlock);
    }

    function oneStepProveInboxConsistency(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _lowerHash,
        bytes32 _value
    ) external inboxConsistencyChallenge onlyOnTurn {
        bytes32 prevHash =
            ChallengeLib.bisectionChunkHash(
                1,
                _lowerHash,
                Messages.addMessageToInbox(_lowerHash, _value)
            );

        verifySegmentProof(_proof, prevHash, _segmentToChallenge);

        emit OneStepProofCompleted();
        _currentWin();
    }

    function bisectInboxDelta(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldInboxAcc,
        bytes32 _oldInboxDelta,
        bytes32 _newInboxDelta,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength
    ) external inboxDeltaChallenge onlyOnTurn {
        require(_chainLength > 1, "bisection too short");
        require(_chainHashes.length == bisectionDegree(_chainLength));
        require(_newInboxDelta != _oldInboxDelta);
        bytes32 oldInboxDeltaHash = ChallengeLib.inboxDeltaHash(_oldInboxAcc, _oldInboxDelta);
        bytes32 newInboxDeltaHash = ChallengeLib.inboxDeltaHash(_oldInboxAcc, _newInboxDelta);
        require(_chainHashes[_chainHashes.length - 1] == newInboxDeltaHash);

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(_chainLength, _chainHashes[0], oldInboxDeltaHash);

        verifySegmentProof(_proof, bisectionHash, _segmentToChallenge);

        challengeState = calculateBisectionRoot(_chainHashes, _chainLength);

        responded(1);
        emit Bisected(_segmentToChallenge, _chainHashes, _chainLength, deadlineBlock);
    }

    function oneStepProveInboxDelta(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _prevInboxAcc,
        bytes32 _prevInboxDelta,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes calldata _msgData
    ) external inboxDeltaChallenge onlyOnTurn {
        verifySegmentProof(
            _proof,
            ChallengeLib.bisectionChunkHash(
                1,
                ChallengeLib.inboxDeltaHash(_prevInboxAcc, _prevInboxDelta),
                oneStepProofInboxDeltaAfter(
                    _prevInboxAcc,
                    _prevInboxDelta,
                    _kind,
                    _blockNumber,
                    _timestamp,
                    _sender,
                    _inboxSeqNum,
                    _msgData
                )
            ),
            _segmentToChallenge
        );

        emit OneStepProofCompleted();
        _currentWin();
    }

    function _verifyBisection(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength,
        bytes32 _bisectionHash,
    ) external executionChallenge onlyOnTurn {
        require(_chainLength > 1, "TOO_SHORT");
        require(_chainHashes.length == bisectionDegree(_chainLength), "BISECT_DEGREE");
        require(_chainHashes[_chainHashes.length - 1] != _oldEndHash, "SAME_END");

        verifySegmentProof(_proof, _bisectionHash, _segmentToChallenge);
    }

    function bisectExecution(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32[] calldata _chainHashes,
        uint256 _chainLength,
        uint256[4] segmentPreFields,
    ) external executionChallenge onlyOnTurn {
        require(_chainLength > 1, "TOO SHORT");

        require(_chainHashes[0] == ChallengeLib.assertionHash(
            _segmentPreFields[0],
            _segmentPreFields[1],
            _segmentPreFields[2],
            _segmentPreFields[3]
        ), "segment pre-fields");

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(_chainLength, _chainHashes[0], _oldEndHash);
        _verifyBisection(_segmentToChallenge, _proof, _oldEndHash, _chainHashes, _chainLength, bisectionHash);

        uint256 segmentStartGas = segmentPreFields[1];

        require(_chainLength + gasBefore > segmentStartGas, "invalid segment length");
        challengeState = calculateBisectionRoot(_chainHashes, _chainLength + gasBefore - segmentStartGas);
        gasBefore = segmentStartGas;

        responded(executionCheckTimeBlocks);
        emit Bisected(_segmentToChallenge, _chainHashes, _chainLength, deadlineBlock);
    }

    function constraintWinExecution(
        uint256 _segmentToChallenge,
        bytes calldata _proof,
        bytes32 _oldEndHash,
        bytes32 _beforeChainHash,
        bytes32 _afterChainHash,
        uint256[4] _beforeFields,
        uint256 _chainLength,
    ) external executionChallenge onlyOnTurn {
        require(_chainLength > 1, "TOO SHORT");

        require(_beforeChainHash == ChallengeLib.assertionHash(
            _beforeFields[0],
            _beforeFields[1],
            _beforeFields[2],
            _beforeFields[3]
        ), "segment pre-fields");

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(_chainLength, _beforeChainHash, _afterChainHash);
        verifySegmentProof(_proof, bisectionHash, _segmentToChallenge);

        require(_beforeFields[1] >= gasBefore + _chainLength);
        require(_beforeChainHash != _afterChainHash);
        emit ConstraintWin();
        _currentWin();
    }

    // machineFields
    //  initialInbox
    //  initialMessage
    //  initialLog
    function oneStepProveExecution(
        uint256 _segmentToChallenge,
        bytes memory _proof,
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

        bytes32 rootHash =
            ChallengeLib.bisectionChunkHash(
                gasUsed,
                oneStepProofExecutionBefore(
                    _machineFields,
                    _initialGasUsed,
                    _initialMessageCount,
                    _initialLogCount,
                    proofFields
                ),
                oneStepProofExecutionAfter(
                    _machineFields,
                    _initialGasUsed,
                    _initialMessageCount,
                    _initialLogCount,
                    gasUsed,
                    proofFields
                )
            );

        _verifyBisection(_segmentToChallenge, _proof, _oldEndHash, _chainHashes, _chainLength, rootHash);

        emit OneStepProofCompleted();
        _currentWin();
    }

    // Can only do a stopped short bisection as a first move
    function bisectExecutionStoppedShort(
        uint256 _prevStepsExecuted,
        bytes32 _startAssertionHash,
        bytes32 _prevEndAssertionHash,
        bytes32[] calldata _chainHashes,
        uint256 _newStepsExecuted
    ) external onlyOnTurn {
        require(kind == Kind.Uninitialized);
        // Unlike the other bisections, it's safe for the number of steps executed to be 0 or 1
        require(_chainHashes.length == bisectionDegree(_newStepsExecuted));
        require(_newStepsExecuted < _prevStepsExecuted);

        require(
            ChallengeLib.bisectionChunkHash(
                _prevStepsExecuted,
                _startAssertionHash,
                _prevEndAssertionHash
            ) == executionHash
        );

        bytes32 bisectionRoot;
        bytes32 endAssertion;
        if (_newStepsExecuted > 0) {
            bisectionRoot = calculateBisectionRoot(_chainHashes, _newStepsExecuted);
            endAssertion = _chainHashes[_chainHashes.length - 1];
        } else {
            bisectionRoot = 0;
            endAssertion = _startAssertionHash;
        }

        initializeKind(Kind.StoppedShort, bisectionRoot);

        // Reuse the executionHash variable to store last assertion
        executionHash = endAssertion;
        responded(executionCheckTimeBlocks);
    }

    function oneStepProveStoppedShortCanRun(
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

        if (state == State.AsserterTurn) {
            emit AsserterTimedOut();
            _challengerWin();
        } else {
            emit ChallengerTimedOut();
            _asserterWin();
        }
    }

    function verifyAndSetup(Kind _kind, bytes32 initialState) private {
        if (kind == Kind.Uninitialized) {
            initializeKind(_kind, initialState);
        } else {
            require(kind == _kind);
        }
    }

    function initializeKind(Kind _kind, bytes32 initialState) private {
        challengeState = initialState;
        kind = _kind;
        // Free no longer needed storage
        inboxConsistencyHash = 0;
        inboxDeltaHash = 0;
        executionHash = 0;
    }

    function responded(uint256 additionalTimeBlocks) private {
        if (state == State.ChallengerTurn) {
            state = State.AsserterTurn;
        } else {
            state = State.ChallengerTurn;
        }
        deadlineBlock = block.number + challengePeriodBlocks + additionalTimeBlocks;
    }

    function _currentWin() private {
        if (state == State.AsserterTurn) {
            _asserterWin();
        } else {
            _challengerWin();
        }
    }

    function _asserterWin() private {
        IRollup(rollupAddress).completeChallenge(asserter, challenger);
        safeSelfDestruct(msg.sender);
    }

    function _challengerWin() private {
        IRollup(rollupAddress).completeChallenge(challenger, asserter);
        safeSelfDestruct(msg.sender);
    }

    function verifySegmentProof(
        bytes memory _proof,
        bytes32 item,
        uint256 _segmentToChallenge
    ) private view {
        (bytes32 calcRoot, ) = MerkleLib.verifyMerkleProof(_proof, item, _segmentToChallenge + 1);
        require(challengeState == calcRoot, BIS_PREV);
    }

    function bisectionDegree(uint256 _chainLength) private pure returns (uint256) {
        if (_chainLength < BISECTION_DEGREE) {
            return _chainLength;
        } else {
            return BISECTION_DEGREE;
        }
    }

    function calculateBisectionRoot(bytes32[] memory _chainHashes, uint256 _chainLength)
        private
        pure
        returns (bytes32)
    {
        uint256 bisectionCount = _chainHashes.length - 1;
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

        (bytes32 root, ) = MerkleLib.generateMerkleRoot(hashes);
        return root;
    }

    function oneStepProofInboxDeltaAfter(
        bytes32 _prevInboxAcc,
        bytes32 _prevInboxDelta,
        uint8 _kind,
        uint256 _blockNumber,
        uint256 _timestamp,
        address _sender,
        uint256 _inboxSeqNum,
        bytes memory _msgData
    ) private pure returns (bytes32) {
        return
            ChallengeLib.inboxDeltaHash(
                Messages.addMessageToInbox(
                    _prevInboxAcc,
                    Messages.messageHash(
                        _kind,
                        _sender,
                        _blockNumber,
                        _timestamp,
                        _inboxSeqNum,
                        keccak256(_msgData)
                    )
                ),
                Messages.addMessageToInbox(
                    _prevInboxDelta,
                    Hashing.hash(
                        Messages.messageValue(
                            _kind,
                            _blockNumber,
                            _timestamp,
                            _sender,
                            _inboxSeqNum,
                            _msgData
                        )
                    )
                )
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
        bytes32 a1OutputAccHash =
            ChallengeLib.outputAccHash(
                _machineFields[1],
                _initialMessageCount,
                _machineFields[2],
                _initialLogCount
            );
        return
            ChallengeLib.assertionHash(
                _machineFields[0],
                _initialGasUsed,
                a1OutputAccHash,
                proofFields[0]
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
        bytes32 a2OutputAccHash =
            ChallengeLib.outputAccHash(
                proofFields[3],
                _initialMessageCount + (_machineFields[1] == proofFields[3] ? 0 : 1),
                proofFields[4],
                _initialLogCount + (_machineFields[2] == proofFields[4] ? 0 : 1)
            );
        return
            ChallengeLib.assertionHash(
                proofFields[2],
                _initialGasUsed + gasUsed,
                a2OutputAccHash,
                proofFields[1]
            );
    }
}
