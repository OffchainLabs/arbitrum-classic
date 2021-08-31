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
import "@openzeppelin/contracts/math/SafeMath.sol";

import "./IChallenge.sol";
import "../rollup/facets/RollupUser.sol";
import "../arch/IOneStepProof.sol";

import "./ChallengeLib.sol";

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
    // deadline expired
    string private constant BIS_DEADLINE = "BIS_DEADLINE";
    // Only original asserter can continue bisect
    string private constant BIS_SENDER = "BIS_SENDER";
    // Incorrect previous state
    string private constant BIS_PREV = "BIS_PREV";
    // Can't timeout before deadline
    string private constant TIMEOUT_DEADLINE = "TIMEOUT_DEADLINE";

    bytes32 private constant UNREACHABLE_ASSERTION = bytes32(uint256(0));

    IOneStepProof[] public executors;
    address[2] public bridges;

    RollupUserFacet internal resultReceiver;

    uint256 maxMessageCount;

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
        uint256 _maxMessageCount,
        address _asserter,
        address _challenger,
        uint256 _asserterTimeLeft,
        uint256 _challengerTimeLeft,
        ISequencerInbox _sequencerBridge,
        IBridge _delayedBridge
    ) external override {
        require(turn == Turn.NoChallenge, CHAL_INIT_STATE);

        executors = _executors;

        resultReceiver = RollupUserFacet(_resultReceiver);

        maxMessageCount = _maxMessageCount;

        asserter = _asserter;
        challenger = _challenger;
        asserterTimeLeft = _asserterTimeLeft;
        challengerTimeLeft = _challengerTimeLeft;

        turn = Turn.Challenger;

        challengeState = _executionHash;

        lastMoveBlock = block.number;
        bridges = [address(_sequencerBridge), address(_delayedBridge)];

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
        uint256 challengeExecutionBisectionDegree =
            resultReceiver.challengeExecutionBisectionDegree();
        require(
            _chainHashes.length ==
                bisectionDegree(_challengedSegmentLength, challengeExecutionBisectionDegree) + 1,
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
        require(
            ChallengeLib.verifySegmentProof(
                challengeState,
                bisectionHash,
                _merkleNodes,
                _merkleRoute
            ),
            BIS_PREV
        );

        bytes32 newChallengeState =
            ChallengeLib.updatedBisectionRoot(
                _chainHashes,
                _challengedSegmentStart,
                _challengedSegmentLength
            );
        challengeState = newChallengeState;

        emit Bisected(
            newChallengeState,
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
        require(
            ChallengeLib.verifySegmentProof(
                challengeState,
                bisectionHash,
                _merkleNodes,
                _merkleRoute
            ),
            BIS_PREV
        );

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
        uint256 _initialMessagesRead,
        bytes32[2] calldata _initialAccs,
        uint256[3] memory _initialState,
        bytes memory _executionProof,
        bytes memory _bufferProof,
        uint8 prover
    ) external onlyOnTurn {
        bytes32 rootHash;
        {
            (uint64 gasUsed, uint256 totalMessagesRead, bytes32[4] memory proofFields) =
                executors[prover].executeStep(
                    bridges,
                    _initialMessagesRead,
                    _initialAccs,
                    _executionProof,
                    _bufferProof
                );

            require(totalMessagesRead <= maxMessageCount, "TOO_MANY_MESSAGES");

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
                        totalMessagesRead,
                        proofFields
                    ),
                "WRONG_END"
            );

            rootHash = ChallengeLib.bisectionChunkHash(
                _challengedSegmentStart,
                _challengedSegmentLength,
                oneStepProofExecutionBefore(
                    _initialMessagesRead,
                    _initialAccs[0],
                    _initialAccs[1],
                    _initialState,
                    proofFields
                ),
                _oldEndHash
            );
        }

        require(
            ChallengeLib.verifySegmentProof(challengeState, rootHash, _merkleNodes, _merkleRoute),
            BIS_PREV
        );

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

    function clearChallenge() external override {
        require(msg.sender == address(resultReceiver), "NOT_RES_RECEIVER");
        safeSelfDestruct(msg.sender);
    }

    function _currentWin() private {
        // As a safety measure, challenges can only be resolved by timeouts during mainnet beta.
        // As state is 0, no move is possible. The other party will lose via timeout
        challengeState = bytes32(0);

        // if (turn == Turn.Asserter) {
        //     _asserterWin();
        // } else {
        //     _challengerWin();
        // }
    }

    function _asserterWin() private {
        resultReceiver.completeChallenge(asserter, challenger);
        safeSelfDestruct(msg.sender);
    }

    function _challengerWin() private {
        resultReceiver.completeChallenge(challenger, asserter);
        safeSelfDestruct(msg.sender);
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
                    totalMessagesRead,
                    proofFields[1],
                    proofFields[2],
                    newSendCount,
                    proofFields[3],
                    newLogCount
                )
            );
    }
}
