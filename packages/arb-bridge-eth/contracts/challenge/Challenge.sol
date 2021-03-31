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
    event ContinuedExecutionProven();

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
    bytes32 private constant UNREACHABLE_ASSERTION = bytes32(uint256(0));

    IOneStepProof[] public executors;
    ISequencerInbox public sequencerBridge;
    IBridge public delayedBridge;

    IRollup internal resultReceiver;

    uint256 maxMessageCount;
    uint256 maxSeqBatchCount;

    address public override asserter;
    address public override challenger;

    uint256 public override lastMoveBlock;
    uint256 public asserterTimeLeft;
    uint256 public challengerTimeLeft;

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

    function initializeChallenge(
        IOneStepProof[] calldata _executors,
        address _resultReceiver,
        bytes32 _executionHash,
        uint256[2] calldata _maxMessageAndBatchCounts,
        address _asserter,
        address _challenger,
        uint256 _asserterTimeLeft,
        uint256 _challengerTimeLeft,
        ISequencerInbox _sequencerBridge,
        IBridge _delayedBridge
    ) external override {
        require(turn == Turn.NoChallenge, CHAL_INIT_STATE);

        executors = _executors;

        resultReceiver = IRollup(_resultReceiver);

        maxMessageCount = _maxMessageAndBatchCounts[0];
        maxSeqBatchCount = _maxMessageAndBatchCounts[1];

        asserter = _asserter;
        challenger = _challenger;
        asserterTimeLeft = _asserterTimeLeft;
        challengerTimeLeft = _challengerTimeLeft;

        turn = Turn.Challenger;

        challengeState = _executionHash;

        lastMoveBlock = block.number;
        sequencerBridge = _sequencerBridge;
        delayedBridge = _delayedBridge;

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
    ) external onlyOnTurn {
        if (_chainHashes[_chainHashes.length - 1] != UNREACHABLE_ASSERTION) {
            require(_challengedSegmentLength > 1, "TOO_SHORT");
        }
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
        require(_chainHashes[0] != UNREACHABLE_ASSERTION, "UNREACHABLE_START");

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

        updateBisectionRoot(_chainHashes, _challengedSegmentStart, _challengedSegmentLength);

        emit Bisected(
            challengeState,
            _challengedSegmentStart,
            _challengedSegmentLength,
            _chainHashes
        );
    }

    function proveContinuedExecution(
        bytes32[] calldata _merkleNodes,
        uint256 _merkleRoute,
        uint256 _challengedSegmentStart,
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        uint256 _gasUsedBefore,
        bytes32 _assertionRest
    ) external onlyOnTurn {
        bytes32 beforeChainHash = ChallengeLib.assertionHash(_gasUsedBefore, _assertionRest);

        bytes32 bisectionHash =
            ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                beforeChainHash,
                _oldEndHash
            );
        verifySegmentProof(bisectionHash, _merkleNodes, _merkleRoute);

        require(
            _gasUsedBefore >= _challengedSegmentStart.add(_challengedSegmentLength),
            "NOT_CONT"
        );
        require(beforeChainHash != _oldEndHash, "WRONG_END");
        emit ContinuedExecutionProven();
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
        uint256 _challengedSegmentLength,
        bytes32 _oldEndHash,
        uint256[2] calldata _initialMessagesAndBatchesRead,
        bytes32[2] calldata _initialAccs,
        uint256[3] memory _initialState,
        bytes memory _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) public onlyOnTurn {
        bytes32 rootHash;
        {
            (
                uint64 gasUsed,
                uint256[2] memory totalMessagesAndBatchesRead,
                bytes32[4] memory proofFields
            ) =
                executors[prover].executeStep(
                    sequencerBridge,
                    delayedBridge,
                    _initialMessagesAndBatchesRead,
                    _initialAccs,
                    _executionProof,
                    _bufferProof
                );

            require(totalMessagesAndBatchesRead[0] <= maxMessageCount, "TOO_MANY_MESSAGES");
            require(totalMessagesAndBatchesRead[1] <= maxSeqBatchCount, "TOO_MANY_BATCHES");

            require(
                // if false, this segment must be proven with proveContinuedExecution
                _initialState[0] < _challengedSegmentStart.add(_challengedSegmentLength),
                "OSP_CONT"
            );
            require(
                _initialState[0].add(gasUsed) >=
                    _challengedSegmentStart.add(_challengedSegmentLength),
                "OSP_SHORT"
            );

            require(
                _oldEndHash !=
                    oneStepProofExecutionAfter(
                        _initialAccs[0],
                        _initialAccs[1],
                        _initialState,
                        gasUsed,
                        totalMessagesAndBatchesRead,
                        proofFields
                    ),
                "WRONG_END"
            );

            rootHash = ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                oneStepProofExecutionBefore(
                    _initialMessagesAndBatchesRead,
                    _initialAccs[0],
                    _initialAccs[1],
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

    function timeout() external override {
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

    function currentResponderTimeLeft() public view override returns (uint256) {
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

    // proofFields
    //  initialMachineHash
    //  afterMachineHash
    //  afterInboxAcc
    //  afterMessagesHash
    //  afterLogsHash
    function oneStepProofExecutionBefore(
        uint256[2] calldata _initialMessagesAndBatchesRead,
        bytes32 _initialSendAcc,
        bytes32 _initialLogAcc,
        uint256[3] memory _initialState,
        bytes32[4] memory proofFields
    ) private pure returns (bytes32) {
        return
            ChallengeLib.assertionHash(
                _initialState[0],
                ChallengeLib.assertionRestHash(
                    _initialMessagesAndBatchesRead[0],
                    _initialMessagesAndBatchesRead[1],
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
        uint256[2] memory totalMessagesAndBatchesRead,
        bytes32[4] memory proofFields
    ) private pure returns (bytes32) {
        uint256 newSendCount = _initialState[1].add((_initialSendAcc == proofFields[2] ? 0 : 1));
        uint256 newLogCount = _initialState[2].add((_initialLogAcc == proofFields[3] ? 0 : 1));
        // The one step proof already guarantees us that firstMessage and lastMessage
        // are either one or 0 messages apart and the same is true for logs. Therefore
        // we can infer the message count and log count based on whether the fields
        // are equal or not
        return
            ChallengeLib.assertionHash(
                _initialState[0].add(gasUsed),
                ChallengeLib.assertionRestHash(
                    totalMessagesAndBatchesRead[0],
                    totalMessagesAndBatchesRead[1],
                    proofFields[1],
                    proofFields[2],
                    newSendCount,
                    proofFields[3],
                    newLogCount
                )
            );
    }
}
